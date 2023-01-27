package main

import (
	"appDB/package/my_micro_serv"
	"appDB/package/storage/memDB"
	"appDB/package/storage/mongo"
	"appDB/package/storage/postgres"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	//"net/http"
)


func main(){
	_ = memDB.New()
	_ = postgres.New()
	db := mongo.New()
	defer db.Close()     // не зыбываем закрыть подключения к БД
	server := my_micro_serv.New(db)
	ListenAndServe(":8051",server.Router())
}


func ListenAndServe(addr string, router *mux.Router){
	err := http.ListenAndServe(addr, router)
	if err != nil{
		log.Fatal(err)
	}
}

/*

// for Testing...

fmt.Println("Connection successful")
var(
	err error
	result []storage.Post
)

result, err = db.Posts()
if err != nil{
log.Fatal(err)
}
for _, post := range result{
fmt.Println(post)
}
fmt.Println("Hello world")
err = db.AddPost(storage.Post{ID: 2, Title: "Hello", Content: "There is a very interesting story",
AuthorID: 0, AuthorName: "Vladimir", CreatedAt : 33563, PublishedAt : 656768887})
if err != nil{
log.Fatal(err)
}
err = db.UpdatePost(storage.Post{ID: 2, Title: "33333", Content: "333333",
AuthorID: 0, AuthorName: "33333", CreatedAt : 333333, PublishedAt : 333333333})
if err != nil{
log.Fatal(err)
}

result, err = db.Posts()
if err != nil{
log.Fatal(err)
}
for _, post := range result{
fmt.Println(post)
}

fmt.Println("Hello world")

err = db.DeletePost(storage.Post{ID : 2})
if err != nil{
log.Fatal(err)
}



result, err = db.Posts()
if err != nil{
log.Fatal(err)
}
for _, post := range result{
fmt.Println(post)
}*/