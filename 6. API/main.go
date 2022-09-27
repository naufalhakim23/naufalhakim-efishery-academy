package main

// http stateful
//  sedangkan lgin tuh stateless
// HTTPS merupakan HTTP yang menggunakan enkripsi/sertifikat/SSL untuk mengamankan data yang dikirimkan untuk men-encrypt data yang dikirimkan ke server
//
// JSON merupakan salah satu alat percakapan/komunikasi antara server dan client
import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware" // middleware
)

type Product struct {
	ID          int    `json:"id"`
	Product     string `json:"product"`
	Description string `json:"description"`
}

var (
	product = map[int]*Product{}
	nomor   = 1
)

func CreateProduct(c echo.Context) error {
	p := &Product{
		ID: nomor,
	}
	if err := c.Bind(p); err != nil {
		return err
	}

	product[p.ID] = p
	nomor++
	return c.JSON(http.StatusCreated, p)

}

func GetProducts(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, product[id])
}

func GetAllProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, product)
}

func UpdateProduct(c echo.Context) error {
	p := new(Product)
	if err := c.Bind(p); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	product[id].Product = p.Product
	product[id].Description = p.Description
	return c.JSON(http.StatusOK, product[id])
}

func UpdateProductPatch(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	p := product[id]
	if err := c.Bind(p); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, p)
}
func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(product, id)
	return c.NoContent(http.StatusNoContent)
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use((middleware.Logger()))
	e.Use(middleware.Recover())

	// Routes
	e.POST("/api/products", CreateProduct)
	e.GET("/api/products", GetAllProducts)
	e.GET("/api/products/:id", GetProducts)
	e.DELETE("/api/products/:id", DeleteUser)
	e.PUT("/api/products/:id", UpdateProduct)
	e.PATCH("/api/products/:id", UpdateProductPatch)

	// Start Server
	e.Logger.Fatal(e.Start(":1323"))

}
