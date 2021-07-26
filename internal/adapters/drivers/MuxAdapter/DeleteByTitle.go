package muxadapter

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	web "github.com/AlejandroWaiz/novels-box/infrastructure/HttpResponseMock"
)

func (ma *MuxAdapter) DeleteNovelByTitle(w http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {

		web.ErrBadRequest.Send(w)
		return

	}

	var titleOfNovel string

	err = json.Unmarshal(reqBody, &titleOfNovel)

	if err != nil {

		web.ErrInvalidJSON.Send(w)
		return

	}

	err = ma.domainport.DeleteByTitle(titleOfNovel)

	if err != nil {

		web.InternalError.Send(w)
		return

	}

	web.Success("Omedetto, your novel has been deleted Successfully!", 200)

}
