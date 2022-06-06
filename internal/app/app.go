package app

import (
	"context"
	"fmt"
	"github.com/iSezam/go-delivery/internal/loader"
	"github.com/iSezam/go-delivery/internal/repository"
	"os"

	"github.com/iSezam/go-delivery/package/common/storage"
)

func Run() {
	db := storage.NewPostgresDB(context.Background(),
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			os.Getenv("DB_SERVER"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_DATABASE"),
			os.Getenv("DB_SSL")))

	ukrposhtaRepo := repository.NewUkrposhtaRepository(db)

	up := loader.NewUkrposhtaLoader(ukrposhtaRepo)
	up.Run()

	onShutdown(func() {

	})
}
