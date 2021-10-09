package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"

	"github.com/rabindra/mongo-golang/controllers"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.POST("/users", uc.CreateUser)
	r.GET("/user/:id", uc.GetUser)
	http.ListenAndServe("localhost:9000", r)
	

}

func getSession() *mgo.Session{

	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil{
		panic(err)
	}
	return s
}
