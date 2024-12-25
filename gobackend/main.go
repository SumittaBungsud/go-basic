package main

import (
	"log"
	"os"
	"time"

	_ "github.com/SumittaBungsud/gobackend/docs"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v4"
	"github.com/gofiber/swagger"
	"github.com/gofiber/template/html/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

// Maps Book fields from json request/response by Fiber, and table fields of database
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func checkMiddleware(c *fiber.Ctx) error {
	// start := time.Now()

	// fmt.Printf("URL = %s, Mthod = %s, Time = %s\n",c.OriginalURL(), c.Method(),start)
	// return c.Next()
	user := c.Locals("user").(*jwt.Token) // get the locals user that was written by the jwtware to the local context
	claims := user.Claims.(jwt.MapClaims)

	if claims["role"] != "admin" {
		return fiber.ErrUnauthorized
	}

	return c.Next()
}

// @title Book API
// @version 1.0
// @description This is a sample server for a book API
// @host localhost:8080
// @BasePath /
// @schemes http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main(){
	if err := godotenv.Load(); err != nil {
		log.Fatal("Load .env error")
	}

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	books = append(books, Book{ID:1, Title:"Mickey Life", Author:"Mickey Mouse"})
	books = append(books, Book{ID:2, Title:"Donald Life", Author:"Donald Duck"})

	app.Post("/login", login)

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
	}))

	app.Use(checkMiddleware) // check user role if login successfully

	app.Get("/hello", func(c *fiber.Ctx) error{
		return c.SendString("Hello World!")
	})
	app.Get("/books", getBooks)
	app.Get("/books/:id", getBook)
	app.Post("/books", createBook)
	app.Put("/books/:id", updateBook)
	app.Delete("/books/:id", deleteBook)

	app.Post("/upload", uploadFile)
	app.Get("/test-html", testHTML)
	app.Get("/config", getENV)

	app.Listen(":8080")
}

type User struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

var memberUser = User{ // dummy user
	Email: "user@example.com",
	Password: "password123",
}

func login(c *fiber.Ctx) error{
	user := new(User)

	if err := c.BodyParser(user); err != nil { // detect request body
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if user.Email != memberUser.Email || user.Password != memberUser.Password { // mismatch user authorization
		return fiber.ErrUnauthorized
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"email":  "John Doe",
		"role": "admin",
		"exp":   time.Now().Add(time.Hour * 72).Unix(), // expire time
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"message": "Login success!",
		"token": t,
	})
}

// nodemon --watch . --ext go --exec go run . --signal SIGTERM