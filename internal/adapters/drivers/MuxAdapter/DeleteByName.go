package muxadapter

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	web "github.com/AlejandroWaiz/novels-box/infrastructure/HttpResponseMock"
)

func (ma *MuxAdapter) DeleteNovelByName(w http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {

		web.ErrBadRequest.Send(w)

	}

	var nameOfNovel string

	err = json.Unmarshal(reqBody, &nameOfNovel)

	if err != nil {

		web.ErrInvalidJSON.Send(w)

	}

	err = ma.domainport.DeleteByName(nameOfNovel)

	if err != nil {

		web.InternalError.Send(w)

	}

	web.Success("Omedetto, your novel has been deleted Successfully!", 200)

}