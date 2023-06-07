package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	// Initialize the database connection
	connectionString := "user:password@tcp(mysql-container:3306)/database"
	var err error
	DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	// Test the database connection
	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to the database!")

	Migrate()
	log.Print("Migrated the database!")
}

func Migrate(){
	_, err := DB.Exec("CREATE TABLE IF NOT EXISTS burritos (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255), size VARCHAR(255), price DECIMAL(10,2))")
	if err != nil {
		log.Fatal(err)
	}
	
	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS orders (id INT AUTO_INCREMENT PRIMARY KEY, total_cost DECIMAL(10,2))")
	if err != nil {
		log.Fatal(err)
	}

	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS order_items (id INT AUTO_INCREMENT PRIMARY KEY, order_id INT, burrito_id INT, quantity INT, FOREIGN KEY (order_id) REFERENCES orders(id), FOREIGN KEY (burrito_id) REFERENCES burritos(id))")
	if err != nil {
		log.Fatal(err)
	}
}

func CloseDB() {
	// Close the database connection
	err := DB.Close()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Disconnected from the database!")
}
