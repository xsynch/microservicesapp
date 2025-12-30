package handlers

import (
	"net/http"
	"strconv"
	
	"github.com/gorilla/mux"
	"github.com/xsynch/microserviceapp/data"

)


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