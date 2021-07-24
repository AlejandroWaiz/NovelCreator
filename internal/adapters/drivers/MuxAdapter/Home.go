package muxadapter

import (
	"fmt"
	"net/http"
)

func (ma *MuxAdapter) Home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello, this is the main page! To add a novel try: /Add ")
	fmt.Fprintf(w, "To delete a novel by his name try: /DeleteByName")
	fmt.Fprintf(w, "To get a novel by his name try: /GetByName")
	fmt.Fprintf(w, "Good luck out there, and always remember: Use Condom/Pills!")

}
