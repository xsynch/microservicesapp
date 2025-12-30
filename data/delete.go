package data 

import (
	"slices"
	"fmt"
)


func DeleteProduct(id int  ) error {
	_,pos,err := findProduct(id)
	if err != nil {
		return err
	}
	productList = slices.Delete(productList,pos,pos+1)
	fmt.Println("found",pos)
	return nil 
	
}