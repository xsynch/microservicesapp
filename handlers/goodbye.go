package handlers

import (
	"fmt"
	"log"
	"net/http"
	"io"
)


type  Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, req *http.Request){
		g.l.Println("Goodbye Endpoint Called")
		d, err := io.ReadAll(req.Body)
		if err != nil {
			g.l.Printf("There was an error: %v\n",err.Error())
			
			
			http.Error(rw,"Error reading your data",http.StatusBadRequest)
			return 
		}

		fmt.Fprintf(rw,"This is goodbye: %s\n",d)
}