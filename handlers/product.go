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
	"strconv"

	"github.com/gorilla/mux"
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

// func (p* Products) ServeHTTP(rw http.ResponseWriter, req *http.Request){

// 	if req.Method == http.MethodGet{
// 		p.getProducts(rw,req)
// 		return 
// 	} 

// 	if req.Method == http.MethodPost{
// 		p.addProduct(rw,req)
// 		return 
// 	}
	
// 	if req.Method == http.MethodPut{
// 		r := regexp.MustCompile(`/([0-9]+)`)
// 		g := r.FindAllStringSubmatch(req.URL.Path,-1)

// 		if len(g) != 1 {
// 			p.l.Println("Invalid URI, more than one id")
// 			http.Error(rw,"Invalid URI",http.StatusBadRequest)
// 			return 
// 		}
// 		if len(g[0]) != 2{
// 			p.l.Println("More than one id specified")
// 			http.Error(rw,"Invalid URI",http.StatusBadRequest)
// 			return 
// 		}
// 		idString := g[0][1]
// 		id,err := strconv.Atoi(idString)
// 		if err != nil {
// 			http.Error(rw,"Error converting to integer",http.StatusBadRequest)
// 			return 
// 		}
// 		p.updateProducts(id, rw,req)
// 		return 
		
// 	}
	
// 	rw.WriteHeader(http.StatusMethodNotAllowed)
	

// }

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
// 200: productsResponse

// GetProducts returns the products from the data store
func (p *Products) GetProducts(rw http.ResponseWriter, req *http.Request){
	listProducts := data.GetProducts()
	// data, err := json.Marshal(listProducts)
	err := listProducts.ToJSON(rw)
	if err != nil {
		http.Error(rw,"unable to marshal data and return product data",http.StatusInternalServerError)
	}

}

func (p *Products) AddProduct(rw http.ResponseWriter, req *http.Request) {
	p.l.Println("Handle POST Product")
	// prod := &data.Product{}
	// err := prod.FromJSON(req.Body)
	// if err != nil {
	// 	http.Error(rw,"Error reading the sent data",http.StatusBadRequest)
	// }
	prod := req.Context().Value(KeyProduct{}).(data.Product)
	p.l.Printf("Prod: %#v",prod)
	data.AddProduct(&prod)
}

func (p Products) UpdateProducts(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id,err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw,"Unable convert id",http.StatusBadRequest)
		return 
	}

	p.l.Println("Handle PUT Product",id)
	// prod := &data.Product{}
	// err = prod.FromJSON(req.Body)
	// if err != nil {
	// 	http.Error(rw,"Error reading the sent data",http.StatusBadRequest)
	// }

	prod := req.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id,&prod)
	if err == data.ErrorProductNotFound{
		http.Error(rw,"Product not found",http.StatusNotFound)
		return 
	}
	if err != nil {
		http.Error(rw,err.Error(),http.StatusInternalServerError)
		return 
	}
	

}

// swagger:route DELETE /products/{id} products deleteProducts
// Returns a list of products
// responses:
// 201: noContent

// DeleteProducts deletes a product from the information store

func (p Products) DeleteProduct(rw http.ResponseWriter, req *http.Request){
	p.l.Println("Handling a delete Request")
	vars := mux.Vars(req)
	id,err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw,"Unable convert id",http.StatusBadRequest)
		return 
	}	
	err = data.DeleteProduct(id)
	if err == data.ErrorProductNotFound{
		http.Error(rw,"Product not found",http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw,err.Error(),http.StatusBadRequest)
		return 
	}
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