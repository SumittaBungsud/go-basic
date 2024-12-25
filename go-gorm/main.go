package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host = "localhost"
	port = 5432
	dbname = "mydatabase"
	username = "myuser"
	password = "mypassword"
)

func authRequired(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil || !token.Valid {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	claims := token.Claims.(jwt.MapClaims)
	fmt.Println(claims["user_id"])

	return c.Next()
}

func main() {
	// Connection string
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
    host, port, username, password, dbname)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
		  SlowThreshold: time.Second,   // Slow SQL threshold
		  LogLevel: logger.Info, // Log level
		  Colorful: true,          // Disable color
		},
	)

	// Open a connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("Failed to connect database")
	}

	fmt.Println("Successfully connected!")

	// Build the database with auto migration using struct or models
	db.AutoMigrate(&Book{}, &User{}) // AutoMigrate() cannot delete unused column from a table

	app := fiber.New()
	book := app.Group("/books", authRequired) // used middleware in a group of books path
	user := app.Group("/users")

	book.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(getBooks(db))
	})
	book.Get("/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))

		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		book := getBook(db, id)
		return c.JSON(book)
	})
	book.Post("/", func(c *fiber.Ctx) error {
		book := new(Book)
		if err := c.BodyParser(book); err != nil{
			return c.SendStatus(fiber.StatusBadRequest)
		}
		err = createBook(db, book)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		return c.JSON(fiber.Map{
			"message": "Create Book Successful",
		})
	})
	book.Put("/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		book := new(Book)
		if err := c.BodyParser(book); err != nil{
			return c.SendStatus(fiber.StatusBadRequest)
		}

		book.ID = uint(id)

		err = updateBook(db, book)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		return c.JSON(fiber.Map{
			"message": "Update Book Successful",
		})
	})
	book.Delete("/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		err = deleteBook(db, id)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		return c.JSON(fiber.Map{
			"message": "Delete Book Successful",
		})
	})
	user.Post("/register", func(c *fiber.Ctx) error {
		user := new(User)
		if err := c.BodyParser(user); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		err = createUser(db, user)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		return c.JSON(fiber.Map{
			"message": "Register Successful",
		})
	})
	user.Post("/login", func(c *fiber.Ctx) error {
		user := new(User)
		if err := c.BodyParser(user); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		token, err := loginUser(db, user)
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		// Set cookie
		c.Cookie(&fiber.Cookie{
			Name:     "jwt",
			Value:    token,
			Expires:  time.Now().Add(time.Hour * 72),
			HTTPOnly: true,
		})

		return c.JSON(fiber.Map{"message": "success"})
	})

	app.Listen(":8080")
}
