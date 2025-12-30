package handlers

import (
	"net/http"

	"github.com/xsynch/microserviceapp/data"
)

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
