package main

import (
	"go-clean/adapters"
	"go-clean/entities"
	"go-clean/usecases"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main(){
	app := fiber.New()
	// Initialize gorm and sqlite connection
	db, err := gorm.Open(sqlite.Open("orders.db"), &gorm.Config{})
	if err != nil{
		log.Fatalf("Failed to connect database %v", err)
	}

	// Automigrate database and table schema with gorm
	db.AutoMigrate(&entities.Order{})

	orderRepo := adapters.NewGormOrderRepository(db) // Secondary adapter
	orderService := usecases.NewOrderService(orderRepo) // Implement service
	orderHandler := adapters.NewHttpOrderHandler(orderService) // Primary adapter

	app.Post("/order", orderHandler.CreateOrder)

	app.Listen(":8000")
}