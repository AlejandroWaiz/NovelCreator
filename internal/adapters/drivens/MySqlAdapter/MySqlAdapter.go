package mysqladapter

import "database/sql"

type MySqlAdapter struct {
	dbconn *sql.DB
}

func CreateMySqlAdapter() (*MySqlAdapter, error) {

	db, err := startDatabase()

	if err != nil {

		return &MySqlAdapter{db}, err

	}

	return &MySqlAdapter{db}, err

}
