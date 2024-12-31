package main

import (
	"go-hexa/adapters"
	"go-hexa/core"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main(){
	app := fiber.New()

	// Initialize gorm and sqlite connection
	db, err := gorm.Open(sqlite.Open("orders.db"), &gorm.Config{})
	if err != nil{
		panic("Failed to connect database")
	}

	// Automigrate database and table schema with gorm
	db.AutoMigrate(&core.Order{})

	orderRepo := adapters.NewGormOrderRepository(db) // Secondary adapter
	orderService := core.NewOrderService(orderRepo) // Implement service
	orderHandler := adapters.NewHttpOrderHandler(orderService) // Primary adapter

	app.Post("/order", orderHandler.CreateOrder)

	app.Listen(":8000")
}