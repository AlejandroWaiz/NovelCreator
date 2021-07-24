package muxadapter

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	web "github.com/AlejandroWaiz/novels-box/infrastructure/HttpResponseMock"
	"github.com/AlejandroWaiz/novels-box/internal/domain/structs"
)

func (ma *MuxAdapter) AddNovelByName(w http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {

		web.ErrBadRequest.Send(w)

		return

	}

	var novels structs.Novel

	err = json.Unmarshal(reqBody, &novels)

	if err != nil {

		log.Println(err)

		web.ErrInvalidJSON.Send(w)

		return

	}

	err = ma.domainport.Add(novels)

	if err != nil {

		web.InternalError.Send(w)

		return

	}

	web.Success("Omedetto! Your novel has been created", 200).Send(w)

}
