package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	docs "github.com/nidumolu/simple-gin-swagger/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

// customer represents data about a registering customer.
type customer struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	PhoneNumber  int    `json:"phonenumber"`
	Country      string `json:"country"`
	CountryCode  int    `json:"countrycode"`
	Verified     string `json:"verified"`
	CustomerType string `json:"customertype"`
}

// customers slice to seed customer data.
var customers = []customer{
	{ID: "1", Name: "Ram Kumar", PhoneNumber: 9025017654, Country: "India", CountryCode: 91, Verified: "false", CustomerType: "REGULAR"},
	{ID: "2", Name: "Shyam Kumar", PhoneNumber: 8075017654, Country: "India", CountryCode: 91, Verified: "false", CustomerType: "AGENT"},
	{ID: "3", Name: "Tom Cruise", PhoneNumber: 6075017666, Country: "India", CountryCode: 91, Verified: "false", CustomerType: "RECRUITER"},
}

// GetCustomers responds with the list of all Customers as JSON.
// GetCustomers             godoc
// @Summary      Get customers array
// @Description  Responds with the list of all customers as JSON.
// @Tags         customers
// @Produce      json
// @Success      200  {array}  customer
// @Router       /registration/customers [get]
func GetCustomers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, customers)
}

// PostCustomer  godoc
// @Summary                    Store a new customer
// @Description                Takes a customer JSON and store in DB. Return saved JSON.
// @Tags                       newcustomer
// @Accept                     json
// @Produce                    json
// @Param                      customer body     customer true "customer Data"
// @Success                    200  {object} customer
// @Router                     /registration/customers [post]
func PostCustomer(c *gin.Context) {
	var customer customer

	// Call BindJSON to bind the received JSON to
	// newCustomer
	if err := c.BindJSON(&customer); err != nil {
		return
	}

	// Add the new customer to the slice.
	customers = append(customers, customer)
	c.IndentedJSON(http.StatusCreated, customer)
}

func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/example")
		{
			eg.GET("/helloworld", Helloworld)
		}

		reg := v1.Group("/registration")
		{
			reg.GET("/customers", GetCustomers)
			reg.POST("/customers", PostCustomer)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")

}
