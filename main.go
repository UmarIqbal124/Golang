package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	var name, number string

	db, err := sql.Open("mysql", "root:#465Uia0407@tcp(127.0.0.1)/checkdb")
	if err != nil {
		log.Fatal("Error in opening", err)
	}

	for i := 1; i <= 10; i++ {

		fmt.Print("\nEnter your name: ")
		fmt.Scan(&name)
		fmt.Print("\nEnter your number: ")
		fmt.Scan(&number)
		quary := "INSERT INTO `checkemp` (`empName`, `phoneNumber`) VALUES (?, ?)"
		enter, err := db.ExecContext(context.Background(), quary, name, number)
		if err != nil {
			log.Fatal("Error in entry of data", err)
		}
		id, err := enter.LastInsertId()
		if err != nil {
			log.Fatal("Imposible to retrive last result ", err)
		}

		fmt.Println("\nYour id is : ", id)
	}

	fmt.Print("Check successfully done\n")
	defer db.Close()
}
