package db

import (
	// "database/sql"

	// _ "github.com/mattn/go-sqlite3"
)

// var oldDB *sql.DB

// func oldInitDB() {
// 	var err error
// 	DB, err = sql.Open("sqlite3", "api.db")
// 	if err != nil {
// 		panic("Could not connect to the database")
// 	}
//
// 	DB.SetMaxIdleConns(10)
// 	DB.SetMaxIdleConns(5)
//
// 	createTables()
// }

// func createTables() {
// 	_, err := DB.Exec(`
//     CREATE TABLE IF NOT EXISTS users (
//       id INTEGER PRIMARY KEY AUTOINCREMENT,
//       email TEXT NOT NULL UNIQUE,
//       password TEXT NOT NULL
//     );
//   `)
// 	if err != nil {
// 		panic("Could not create the users table")
// 	}
//
// 	_, err = DB.Exec(`
//     CREATE TABLE IF NOT EXISTS events (
//       id INTEGER PRIMARY KEY AUTOINCREMENT,
//       name TEXT NOT NULL,
//       description TEXT NOT NULL,
//       location TEXT NOT NULL,
//       date_time DATETIME NOT NULL,
//       user_id INTEGER NOT NULL,
//       FOREIGN KEY(user_id) REFERENCES users(id)
//     );
//   `)
// 	if err != nil {
// 		panic("Could not create the events table")
// 	}
//
// 	_, err = DB.Exec(`
// 	  CREATE TABLE IF NOT EXISTS registrations (
// 	    id INTEGER PRIMARY KEY AUTOINCREMENT,
// 	    event_id INTEGER NOT NULL,
// 	    user_id INTEGER NOT NULL,
// 	    FOREIGN KEY(event_id) REFERENCES events(id),
// 	    FOREIGN KEY(user_id) REFERENCES users(id)
// 	  )  
// 	`)
// 	if err != nil {
// 		panic("Could not create the registrations table")
// 	}
// }
