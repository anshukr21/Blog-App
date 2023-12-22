package router

import (
	"github.com/anshukr21/blogapp/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/blogposts", controller.GetAllBlogposts).Methods("GET")
	router.HandleFunc("/api/blogpost", controller.CreateBlogpost).Methods("POST")
	router.HandleFunc("/api/blogpost/{id}", controller.UpdateBlogpost).Methods("PUT")
	router.HandleFunc("/api/blogpost/{id}", controller.DeleteBlogpost).Methods("DELETE")
	router.HandleFunc("/api/deleteallblogpost", controller.DeleteAllBlogposts).Methods("DELETE")

	return router
}
