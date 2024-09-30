package main

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Vulnerable to SQL Injection
	userInput := "admin' OR '1'='1"
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	query := fmt.Sprintf("SELECT * FROM users WHERE username = '%s'", userInput)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Error running query:", err)
		return
	}
	defer rows.Close()

	// Insecure Command Execution
	command := "ls " + os.Args[1]
	out, err := exec.Command("sh", "-c", command).Output()
	if err != nil {
		fmt.Println("Error running command:", err)
		return
	}
	fmt.Println(string(out))
}
