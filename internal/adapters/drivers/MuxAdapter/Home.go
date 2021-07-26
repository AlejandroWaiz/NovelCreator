package muxadapter

import (
	"fmt"
	"net/http"
)

func (ma *MuxAdapter) Home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, " Hello, this is the main page! To add a novel try: /Add\n ")
	fmt.Fprintf(w, "To delete a novel by his title try: /DeleteByTitle\n")
	fmt.Fprintf(w, " To get a novel by his title try: /GetByTitle\n")
	fmt.Fprintf(w, " Good luck out there, and always remember: Use Condom/Pills!\n")

}
