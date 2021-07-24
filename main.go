package main

import (
	mysqladapter "github.com/AlejandroWaiz/novels-box/internal/adapters/drivens/MySqlAdapter"
	muxadapter "github.com/AlejandroWaiz/novels-box/internal/adapters/drivers/MuxAdapter"
	"github.com/AlejandroWaiz/novels-box/internal/domain"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {

		panic("Cant load enviroment variables.")

	}

	dbAdapter, err := mysqladapter.CreateMySqlAdapter()

	if err != nil {

		panic("DB connection failed.")

	}

	domain := domain.CreateNewNovelsDomain(dbAdapter)

	HttpAdapter := muxadapter.CreateMuxAdapter(domain)

	HttpAdapter.Run()

}
