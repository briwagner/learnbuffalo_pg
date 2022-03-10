package actions

import (
	"fmt"
	"learnbuffalo_pg/models"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
)

// BooksIndex default implementation.
func BooksIndex(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	books := models.Books{}

	err := tx.All(&books)
	if err != nil {
		c.Flash().Add("warning", "No books found")
		return c.Redirect(307, "/")
	}

	c.Set("books", books)
	return c.Render(http.StatusOK, r.HTML("books/index.html"))
}

// BooksShow default implementation.
func BooksShow(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	book := models.Book{}

	err := tx.Find(&book, c.Param("id"))
	if err != nil {
		c.Flash().Add("warning", "No book found")
		return c.Redirect(307, "/")
	}

	c.Set("book", book)
	return c.Render(http.StatusOK, r.HTML("books/show.html"))
}

func BooksForm(c buffalo.Context) error {
	book := models.Book{}
	c.Set("book", book)
	return c.Render(http.StatusOK, r.HTML("books/create.html"))
}

// BooksCreate default implementation.
func BooksCreate(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	book := &models.Book{}

	c.Bind(book)
	_, err := tx.ValidateAndCreate(book)
	if err != nil {
		fmt.Print(err)
		c.Flash().Add("warning", "Cannot create book")
		return c.Redirect(307, "/")
	}
	c.Set("book", book)
	c.Flash().Add("info", "New book created")
	return c.Redirect(301, fmt.Sprintf("/books/%s", book.ID.String()))
}
