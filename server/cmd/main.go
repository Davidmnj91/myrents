package main

import (
	"fmt"
	"github.com/Davidmnj91/myrents/pkg/auth"
	"github.com/Davidmnj91/myrents/pkg/http/rest"
	"github.com/Davidmnj91/myrents/pkg/real_state"
	"github.com/Davidmnj91/myrents/pkg/storage/auth/redis"
	realStateRepository "github.com/Davidmnj91/myrents/pkg/storage/real_state/mongo"
	userRepository "github.com/Davidmnj91/myrents/pkg/storage/user/mongo"
	"github.com/Davidmnj91/myrents/pkg/user"
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
	))

	if err != nil {
		panic(err)
	}

	validator, err := validation.NewValidator()
	if err != nil {
		panic(err)
	}

	userRepo := userRepository.NewRepository(dbClient)
	realStateRepo := realStateRepository.NewRepository(dbClient)
	redisRepo := redis.NewRepository(redisClient, int64(tokenTtl))

	authModule := auth.NewAuthModule(tokenSeed, int64(tokenTtl), redisRepo, userRepo, validator)
	userModule := user.NewUserModule(userRepo, validator)
	realStateModule := real_state.NewRealStateModule(realStateRepo, validator)

	router := rest.NewRouter(rest.Routes{
		LoginHandler:             authModule.LoginHandler,
		LogoutHandler:            authModule.LogoutHandler,
		UserRegisterHandler:      userModule.RegisterHandler,
		UserDeleteHandler:        userModule.DeleteHandler,
		UserProfileHandler:       userModule.ProfileHandler,
		RealStateListerHandler:   realStateModule.ListerHandler,
		RealStateRegisterHandler: realStateModule.RegisterHandler,
		RealStateUpdaterHandler:  realStateModule.UpdaterHandler,
		RealStateRemoverHandler:  realStateModule.RemoverHandler,
		AuthMiddleware:           authModule.AuthMiddleware,
	})

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
