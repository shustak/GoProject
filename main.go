package main

import (
	"database/sql"
	"github.corp.globant.com/a-muliarchik/GoTraining/bookPackage"
	"github.corp.globant.com/a-muliarchik/GoTraining/util"
	"log"
	"net/http"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	db, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Printf("Error %s when opening DB", err)
	}
	err = db.Ping()
	if err != nil {
		log.Printf("Errors %s pinging DB", err)
	}
	defer db.Close()

	handle := bookPackage.NewBaseHandler(db)

	log.Printf("Connected to DB %s successfully\n", "books_database")
	log.Println("Server is up on 8080 port")

	router := NewRouter(handle)
	log.Fatalln(http.ListenAndServe(":8080", router))
}
