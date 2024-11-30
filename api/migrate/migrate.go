package main

import (
	"fmt"
	"go-rest-api/db"
	"go-rest-api/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(
		&model.User{},
		&model.Tweet{},
		&model.Favorite{},
		&model.Comment{},
		&model.Retweet{},
		&model.Bookmark{},
		&model.Relationship{},
	)
}
