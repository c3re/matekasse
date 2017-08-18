package matekasse

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func connectDB() error {
	// connect to DB
	var err error
	db, err = sql.Open("sqlite3", dbfile)
	ce(err)
	// Create TABLE scheme
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS matekasse (id INTEGER PRIMARY KEY, name TEXT NOT NULL DEFAULT '', balance INTEGER DEFAULT 0) ")
	ce(err)
	if db != nil {
		return nil
	} else {
		return errors.New("Was not able to connect to DB")
	}
}

func closeDB() {
	if db != nil {
		db.Close()
	}
}

func createUser(id ID) {
	_, err := db.Exec("INSERT INTO matekasse (id) values (?)", id)
	ce(err)
}

func executeBooking(id ID, amount int) {
	if !exists(id) {
		createUser(id)
	}
	_, err := db.Exec("UPDATE matekasse SET balance = balance + ? WHERE id=?", amount, id)
	ce(err)
}

func exists(id ID) bool {
	var returner bool
	err := db.QueryRow("SELECT EXISTS (SELECT id FROM matekasse WHERE id=?)", id).Scan(&returner)
	ce(err)
	return returner
}

func getUser(id ID) account {
	if !exists(id) {
		createUser(id)
	}
	var returner account
	err := db.QueryRow("SELECT id, name, balance FROM matekasse where id=?", id).Scan(&returner.Id, &returner.Name, &returner.Balance)
	ce(err)
	return returner
}

func getAllUsers() []nameIdPair {
	returner := make([]nameIdPair, 0)
	r, err := db.Query("SELECT id, name FROM matekasse WHERE name NOT LIKE ''")
	ce(err)
	defer r.Close()
	for r.Next() {
		var id ID
		var name string
		r.Scan(&id, &name)
		returner = append(returner, nameIdPair{id, name})
	}
	return returner

}

func getSum() int {
	var returner int
	err := db.QueryRow("SELECT SUM(balance) FROM matekasse").Scan(&returner)
	ce(err)
	return returner
}
