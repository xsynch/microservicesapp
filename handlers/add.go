package handlers

import (
	"net/http"

	"github.com/xsynch/microserviceapp/data"
)

func (p *Products) AddProduct(rw http.ResponseWriter, req *http.Request) {
	p.l.Println("Handle POST Product")

	prod := req.Context().Value(KeyProduct{}).(data.Product)
	p.l.Printf("Prod: %#v",prod)
	data.AddProduct(&prod)
}