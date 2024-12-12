package main

import (
	"net/http"

	"github.com/diebbCN/go-mongodb/controllers"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {

	r := httprouter.New() //生成要给httprouter实例,简化操作
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:9001", r)

}

func getSession() *mgo.Session {

	s, err := mgo.Dial(`mongodb://124.71.87.98:27017`)
	if err != nil {
		panic(err)
	}
	return s
}
