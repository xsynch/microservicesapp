package handlers

import (
	"io"
	"log"
	"net/http"
	"fmt"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, req *http.Request){
			
		d, err := io.ReadAll(req.Body)
		if err != nil {
			h.l.Printf("There was an error: %v\n",err.Error())
			
			
			http.Error(rw,"Error reading your data",http.StatusBadRequest)
			return 
		}
		h.l.Printf("Data: %s",d)
		h.l.Println("From handler function in the root")

		fmt.Fprintf(rw,"This is what you sent: %s\n",d)
}