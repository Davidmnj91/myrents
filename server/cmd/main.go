package main

import (
	"github.com/Davidmnj91/myrents/pkg/http/rest"
	"github.com/Davidmnj91/myrents/pkg/jwt"
	"github.com/Davidmnj91/myrents/pkg/login"
	"github.com/Davidmnj91/myrents/pkg/logout"
	"github.com/Davidmnj91/myrents/pkg/middleware"
	"github.com/Davidmnj91/myrents/pkg/storage/auth/redis"
	"github.com/Davidmnj91/myrents/pkg/storage/user/mongo"
	"github.com/Davidmnj91/myrents/pkg/user_register"
	"github.com/Davidmnj91/myrents/pkg/user_remove"
	mongoUtil "github.com/Davidmnj91/myrents/pkg/util/db"
	"github.com/Davidmnj91/myrents/pkg/util/env"
	redisUtil "github.com/Davidmnj91/myrents/pkg/util/redis"
	"github.com/gofiber/fiber/v2"
	"log"
)

const (
	defaultJWTSeed      = "MyRents"
	defaultJWTTTLMillis = 3 * 24 * 60 * 60 * 1000 // 3 days

	defaultRedisAddr = "localhost"
	defaultRedisPass = "MyRents"
	defaultRedisDB   = 0

	defaultDBUser   = "myRents"
	defaultDBPass   = "myRents"
	defaultDBHost   = "localhost"
	defaultDBPort   = "27017"
	defaultDBSchema = "myRents"
)

func main() {
	tokenSeed := env.GetEnvAsStringOrFallback("JWT_SEED", defaultJWTSeed)
	tokenTtl, err := env.GetEnvAsIntOrFallback("JWT_TTL", defaultJWTTTLMillis)
	if err != nil {
		log.Println("WARN: Invalid argument provided for JWT_TTL configuration")
	}

	redisAddr := env.GetEnvAsStringOrFallback("REDIS_ADDR", defaultRedisAddr)
	redisPass := env.GetEnvAsStringOrFallback("REDIS_PASS", defaultRedisPass)
	redisDB, err := env.GetEnvAsIntOrFallback("REDIS_DB", defaultRedisDB)
	if err != nil {
		log.Println("WARN: Invalid argument provided for REDIS_DB configuration")
	}

	dbUser := env.GetEnvAsStringOrFallback("DB_USER", defaultDBUser)
	dbPass := env.GetEnvAsStringOrFallback("DB_PASS", defaultDBPass)
	dbHost := env.GetEnvAsStringOrFallback("DB_HOST", defaultDBHost)
	dbPort := env.GetEnvAsStringOrFallback("DB_PORT", defaultDBPort)
	schema := env.GetEnvAsStringOrFallback("DB_SCHEMA", defaultDBSchema)

	redisClient, err := redisUtil.ConnectRedis(redisUtil.NewRedisConfiguration(redisAddr, redisPass, redisDB))
	if err != nil {
		panic(err)
	}

	collection, err := mongoUtil.ConnectMongo(mongoUtil.NewMongoConfiguration(
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

	repo := mongo.NewRepository(collection)
	redisRepo := redis.NewRepository(redisClient, int64(tokenTtl))

	jwtService := jwt.NewService(tokenSeed, tokenTtl)
	authService := middleware.NewService(jwtService, redis.NewRepository(redisClient, int64(tokenTtl)))

	authMiddleware := middleware.NewAuthMiddleware(authService)
	loginHandler := login.NewLogin(repo, redisRepo, jwtService)
	logoutHandler := logout.NewLogout(redisRepo)

	userRegisterHandler := user_register.NewRegister(repo)
	userDeleteHandler := user_remove.NewReMove(repo)

	router := rest.NewRouter(rest.Routes{LoginHandler: loginHandler, LogoutHandler: logoutHandler, UserRegisterHandler: userRegisterHandler, UserDeleteHandler: userDeleteHandler, AuthMiddleware: authMiddleware})

	api := fiber.New()

	router.Serve(api)
}
