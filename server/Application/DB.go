package Application

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func makeConnection() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=gorm port=5432 sslmode=disable TimeZone=Africa/Cairo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error Message for connection:", err.Error())
	}
	return db
}

func returnConnection(db *gorm.DB) *sql.DB {
	connection, err := db.DB()
	if err != nil {
		fmt.Println("Error Message for get DB:", err.Error())
	}
	return connection
}

// connectToDB func
func connectToDB(share ShareResource) {
	switch share.(type) {
	case *Application:
		app := share.(*Application)
		app.DB = makeConnection()
		app.Connection = returnConnection(app.DB)
	case *Request:
		req := share.(*Request)
		req.DB = makeConnection()
		req.Connection = returnConnection(req.DB)
	}
}

// CloseConnection req method to close DB Connection
func CloseConnection(share ShareResource) {
	var err error
	switch share.(type) {
	case *Application:
		app := share.(*Application)
		err = app.Connection.Close()
	case *Request:
		req := share.(*Request)
		err = req.Connection.Close()
	}
	if err != nil {
		fmt.Println("Error Message for close connection:", err.Error())
	} else {
		fmt.Println("Connection was closed successfully!")
	}
}
