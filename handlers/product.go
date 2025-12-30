// Package classification of Product API
//
// # Documentation for Product API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/jon
//
// Produces:
// - application/json
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/xsynch/microserviceapp/data"
)

// A list of products returns in the response
// swagger:response productResponse
type productResponse struct {
	// All products in the system
	// in: body
	Body []data.Product
}

// swagger:parameters deleteProduct
type productIDParameterWrapper struct {
	// The id of the product to delete from the information store
	// in: path
	// required: true
	ID int `json:"id"`
}

//swagger:response noContent
type productsNoContent struct {

}


type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products{
	return &Products{l}
}





type KeyProduct struct {}

func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func (rw http.ResponseWriter, req *http.Request){
		prod := data.Product{}

		err := prod.FromJSON(req.Body)
		if err != nil {
			p.l.Println("Error decoding json",err.Error())
			http.Error(rw,"Error unmarshaling JSON",http.StatusBadRequest)
			return 
		}
		err = prod.Validate()
		if err != nil {
			p.l.Println("Error Validating Product",err.Error())
			http.Error(rw,fmt.Sprintf("Error validating product:  %s",err.Error()),http.StatusBadRequest)
			return 			
		}
		
		ctx := context.WithValue(req.Context(),KeyProduct{},prod)
		r := req.WithContext(ctx)


		next.ServeHTTP(rw,r)
	})
}