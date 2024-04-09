package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
)

func main() {
	// Connect to MySQL database
	db, err := sql.Open("mysql", "root:netcore@123@tcp(127.0.0.1:3306)/csv")
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	// Open the CSV file
	file, err := os.Open("data.csv")
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer file.Close()

	// Prepare SQL statement
	stmt, err := db.Prepare("INSERT INTO csvtable (id, name, sername) VALUES (?, ?, ?)")
	if err != nil {
		fmt.Println("Error preparing SQL statement:", err)
		return
	}
	defer stmt.Close()

	// Read and insert CSV data
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	for _, record := range records {
		_, err := stmt.Exec(record[0], record[1], record[2]) // Adjust indexes as per your CSV structure
		if err != nil {
			fmt.Println("Error inserting record:", err)
			return
		}
	}

	fmt.Println("Data inserted successfully")
}
