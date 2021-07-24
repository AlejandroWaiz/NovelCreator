package mysqladapter

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func startDatabase() (*sql.DB, error) {

	s := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", os.Getenv("db_user"), os.Getenv("db_passwd"), os.Getenv("db_addr"), os.Getenv("db_db"))

	fmt.Println(s)

	db, err := sql.Open("mysql", s)

	if err != nil {

		log.Println(err)

		return nil, err
	}

	if err := db.Ping(); err != nil {

		log.Println(err)

		return nil, err

	}

	fmt.Println("Connnected to the database!")

	return db, nil

}
