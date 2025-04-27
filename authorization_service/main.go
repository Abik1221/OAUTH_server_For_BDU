package main

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Client struct {
	ID          string         `gorm:"primaryKey"`
	Name        string         `gorm:"unique"`
	Website     string         `gorm:"unique"`
	Logo        string         `gorm:"not null"`
	RedirectURI string         `gorm:"not null"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("unable to load the .env file")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		panic("DB_URL is not set in the .env file")
	}

	dbURL := "postgresql://postgres:password@localhost:5432/authorization"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}

	migrateing the schema

	err = db.AutoMigrate(&Client{})
	if err != nil {
		panic("failed to migrate the database")
	}

	Insert dummy client data to test the system

	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "website", "logo", "redirect_uri"}),
	}).Create(&Client{
		ID:          "1",
		Name:        "Google",
		Website:     "https://test.com",
		Logo:        "https://google.com/logo.png",
		RedirectURI: "http://localhost:8080/auth/callback",
	})

	R := fiber.New(fiber.Config{
		AppName: "Authorization Service",
	})
	R.Use(logger.New())
	R.Use(recover.New())
	R.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Authorization Service")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	R.Listen(":" + port)
}
