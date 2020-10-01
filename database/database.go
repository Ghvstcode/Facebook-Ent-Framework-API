package database

import (
	"context"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/GhvstCode/LR-ENT/ent"
)
var EntClient *ent.Client
func init() {
	//fmt.Println("Connected to database successfully")
	Client, err := ent.Open("postgres","host=localhost port=5432 user=postgres dbname=notesdb password=mysecretpassword sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to database successfully")
	//defer Client.Close()
	//return Client
	if err := Client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	EntClient = Client
}

//docker run -d -p 5432:5432 --name postgres_DB -e POSTGRES_PASSWORD=mysecretpassword postgres
//docker exec -it postgres_DB bash