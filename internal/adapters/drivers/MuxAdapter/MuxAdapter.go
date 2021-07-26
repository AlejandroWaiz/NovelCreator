package muxadapter

import (
	"fmt"
	"log"
	"net/http"
	"os"

	domain "github.com/AlejandroWaiz/novels-box/internal/domain"
	"github.com/gorilla/mux"
)

type MuxAdapter struct {
	domainport domain.Port
}

func CreateMuxAdapter(domain domain.Port) *MuxAdapter {
	return &MuxAdapter{domainport: domain}
}

func (ma *MuxAdapter) Run() {

	port, router := createRouter()

	router.HandleFunc("/", ma.Home)
	router.HandleFunc("/Add", ma.AddNovelByTitle).Methods("POST")
	router.HandleFunc("/DeleteByTitle", ma.DeleteNovelByTitle).Methods("DELETE")
	router.HandleFunc("/GetByTitle", ma.GetNovelByTitle).Methods("GET")
	router.HandleFunc("/GetAll", ma.GetAllNovels).Methods("GET")

	fmt.Println("Router started!")

	log.Fatal(http.ListenAndServe(port, router))

}

func createRouter() (string, *mux.Router) {
	port := fmt.Sprintf(":%s", os.Getenv("MuxPort"))
	router := mux.NewRouter().StrictSlash(true)

	fmt.Println("Starting...")

	return port, router

}
