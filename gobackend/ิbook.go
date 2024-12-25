package main

import (
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Handler functions
// getBooks godoc
// @Summary Get all books
// @Description Get details of all books
// @Tags books
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} Book
// @Router /books [get]
func getBooks(c *fiber.Ctx) error{ // get all books
	return c.JSON(books)
}

func getBook(c *fiber.Ctx) error{ // get a book by id
	bookId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for _, book := range books {
		if book.ID == bookId {
			return c.JSON(book)
		}
	}
	return c.SendStatus(fiber.StatusNotFound)
}

func createBook(c *fiber.Ctx) error{ // create a new book
	book := new(Book)

	// detect request body
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	books = append(books, *book)
	return c.JSON(book)
}

func updateBook(c *fiber.Ctx) error{ // adjust a book by id
	bookId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	bookupdate := new(Book)

	if err := c.BodyParser(bookupdate); err != nil { // detect request body
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for index, book := range books {
		if book.ID == bookId {
			books[index].Title = bookupdate.Title
			books[index].Author = bookupdate.Author
			return c.JSON(books[index])
		}
	}
	return c.SendStatus(fiber.StatusNotFound)
}

func deleteBook(c *fiber.Ctx) error{ // delete a book by id
	bookId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i, book := range books {
		if book.ID == bookId {
			books = append(books[:i],books[i+1:]...)
			return c.JSON(book)
		}
	}
	return c.SendStatus(fiber.StatusNotFound)
}

func uploadFile(c *fiber.Ctx) error{ // upload image file request saving at the server
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = c.SaveFile(file, "./uploads/" + file.Filename)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendString("File upload complete!")
}

func testHTML(c *fiber.Ctx) error{ // update VIEWs with Fiber to HTML template
	return c.Render("index", fiber.Map{
		"Title": "Hello World",
	})
}

func getENV(c *fiber.Ctx) error{ // testing get environment variable from .env
	return c.JSON(fiber.Map{
		"SECRET": os.Getenv("SECRET"),
	})
}