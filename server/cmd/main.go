package main

import (
	"fmt"
	"github.com/Davidmnj91/myrents/pkg/http/rest"
	"github.com/Davidmnj91/myrents/pkg/jwt"
	"github.com/Davidmnj91/myrents/pkg/login"
	"github.com/Davidmnj91/myrents/pkg/logout"
	"github.com/Davidmnj91/myrents/pkg/middleware"
	"github.com/Davidmnj91/myrents/pkg/storage/auth/redis"
	user "github.com/Davidmnj91/myrents/pkg/storage/user/mongo"
	"github.com/Davidmnj91/myrents/pkg/user_register"
	"github.com/Davidmnj91/myrents/pkg/user_remove"
	mongoUtil "github.com/Davidmnj91/myrents/pkg/util/db"
	"github.com/Davidmnj91/myrents/pkg/util/env"
	redisUtil "github.com/Davidmnj91/myrents/pkg/util/redis"
	"github.com/Davidmnj91/myrents/pkg/util/validation"
	"github.com/gofiber/fiber/v2"
	"log"
)

const (
	defaultJWTSeed      = "MyRents"
	defaultJWTTTLMillis = 3 * 24 * 60 * 60 * 1000 // 3 days

	defaultRedisHost = "localhost"
	defaultRedisPort = 6379
	defaultRedisPass = "myRents"
	defaultRedisDB   = 0

	defaultDBUser   = "myRents"
	defaultDBPass   = "myRents"
	defaultDBHost   = "localhost"
	defaultDBPort   = "27017"
	defaultDBSchema = "myRents"

	defaultApiPort = 5000
)

func main() {
	tokenSeed := env.GetEnvAsStringOrFallback("JWT_SEED", defaultJWTSeed)
	tokenTtl, err := env.GetEnvAsIntOrFallback("JWT_TTL", defaultJWTTTLMillis)
	if err != nil {
		log.Println("WARN: Invalid argument provided for JWT_TTL configuration")
	}

	redisHost := env.GetEnvAsStringOrFallback("REDIS_HOST", defaultRedisHost)
	redisPort, err := env.GetEnvAsIntOrFallback("REDIS_PORT", defaultRedisPort)
	if err != nil {
		log.Println("WARN: Invalid argument provided for REDIS_PORT configuration")
	}
	redisPass := env.GetEnvAsStringOrFallback("REDIS_PASS", defaultRedisPass)
	redisDB, err := env.GetEnvAsIntOrFallback("REDIS_DB", defaultRedisDB)
	if err != nil {
		log.Println("WARN: Invalid argument provided for REDIS_DB configuration")
	}

	redisClient, err := redisUtil.ConnectRedis(redisUtil.NewRedisConfiguration(redisHost, redisPort, redisPass, redisDB))
	if err != nil {
		panic(err)
	}

	dbUser := env.GetEnvAsStringOrFallback("DB_USER", defaultDBUser)
	dbPass := env.GetEnvAsStringOrFallback("DB_PASS", defaultDBPass)
	dbHost := env.GetEnvAsStringOrFallback("DB_HOST", defaultDBHost)
	dbPort := env.GetEnvAsStringOrFallback("DB_PORT", defaultDBPort)
	schema := env.GetEnvAsStringOrFallback("DB_SCHEMA", defaultDBSchema)

	dbClient, err := mongoUtil.ConnectMongo(mongoUtil.NewMongoConfiguration(
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		schema,
		"users",
	))

	if err != nil {
		panic(err)
	}

	validator, err := validation.NewValidator()
	if err != nil {
		panic(err)
	}

	repo := user.NewRepository(dbClient)
	redisRepo := redis.NewRepository(redisClient, int64(tokenTtl))

	jwtService := jwt.NewService(tokenSeed, tokenTtl)
	authService := middleware.NewService(jwtService, redis.NewRepository(redisClient, int64(tokenTtl)))

	authMiddleware := middleware.NewAuthMiddleware(authService)
	loginHandler := login.NewLogin(validator, repo, redisRepo, jwtService)
	logoutHandler := logout.NewLogout(redisRepo)

	userRegisterHandler := user_register.NewRegister(repo, validator)
	userDeleteHandler := user_remove.NewReMove(repo)

	router := rest.NewRouter(rest.Routes{LoginHandler: loginHandler, LogoutHandler: logoutHandler, UserRegisterHandler: userRegisterHandler, UserDeleteHandler: userDeleteHandler, AuthMiddleware: authMiddleware})

	apiPort, err := env.GetEnvAsIntOrFallback("API_PORT", defaultApiPort)
	if err != nil {
		log.Println("WARN: Invalid argument provided for API_PORT configuration")
	}

	app := fiber.New()

	api := app.Group("/api")
	router.Serve(api)

	err = app.Listen(fmt.Sprintf(":%d", apiPort))
	if err != nil {
		panic(err)
	}
}
