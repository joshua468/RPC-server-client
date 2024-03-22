package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Item struct {
	Title string
	Body  string
}

func main() {
	var reply Item
	var db []Item

	client, err := rpc.DialHTTP("tcp", "localhost:4040")
	if err != nil {
		log.Fatal("connection error:", err)
	}
	a := Item{"first", "a test item"}
	b := Item{"second", "a second item"}
	c := Item{"third", "a third item"}

	client.Call("API.AddItem", a, &reply)
	client.Call("API.AddItem", b, &reply)
	client.Call("API.AddItem", c, &reply)
	client.Call("API.AddItem", "", &db)
	fmt.Println("Database:", db)
	client.Call("API.EditItem", Item{"Second", "A new second item"}, &reply)

	client.Call("API.DeleteItem", c, &reply)
	client.Call("API.GETDB", "", &db)
	fmt.Println("Database:", db)
}
