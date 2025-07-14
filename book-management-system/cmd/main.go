package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm/dialects/mysql"
	"github.com/FranciscoGontijo/book-management-system/pkg/routes"
	"fmt"
)

func main(){
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	fmt.Printf("Starting server at port 9010\n")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}