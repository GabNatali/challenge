package main

import (
	"context"
	"fmt"
	"log"

	"github.com/codeableorg/weekend-challenge-13-GabNatali/api/cli"
	"github.com/codeableorg/weekend-challenge-13-GabNatali/internal/auth"
	"github.com/codeableorg/weekend-challenge-13-GabNatali/internal/base/crypto"
	"github.com/codeableorg/weekend-challenge-13-GabNatali/internal/base/database"
	"github.com/codeableorg/weekend-challenge-13-GabNatali/internal/entry"
	"github.com/codeableorg/weekend-challenge-13-GabNatali/internal/user"
	"github.com/gin-gonic/gin"
)

func main() {

	ctx := context.Background()

	parser := cli.NewParser()

	config, err := parser.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}

	dbClient := database.NewClient(ctx, config.DatabaseURL)
	err = dbClient.Connect()
	if err != nil {
		log.Fatal(err)
	}

	db := dbClient.DB()

	db.AutoMigrate(&user.UserModel{}, &entry.EntryModel{})
	crypto := crypto.NewCrypto()
	userRepository := user.NewUserRepository(db)
	userUseCase := user.NewUserUseCase(userRepository)

	authSerciceOpts := auth.AuthServiceOpts{
		Crypto:         crypto,
		UserRepository: userRepository,
		Config:         config.AccessTokenSecret,
	}

	authService := auth.NewAuthService(authSerciceOpts)

	entryRepository := entry.NewEntryRepository(db)
	entryUseCases := entry.NewEntryUsesCases(entryRepository)

	router := gin.Default()
	user.AddUserRouter(router, userUseCase)
	auth.AddAuthRouter(router, authService)
	entry.AddEntryRoutes(router, entryUseCases)

	fmt.Printf("API server listening at: %s\n\n", ":3333")
	log.Fatal(router.Run(":3333"))
}
