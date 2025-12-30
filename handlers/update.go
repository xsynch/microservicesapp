package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/xsynch/microserviceapp/data"
)


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