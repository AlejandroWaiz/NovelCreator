package muxadapter

import (
	"net/http"

	web "github.com/AlejandroWaiz/novels-box/infrastructure/HttpResponseMock"
)

func (ma *MuxAdapter) GetAllNovels(w http.ResponseWriter, r *http.Request) {

	novels, err := ma.domainport.GetAllNovels()

	if err != nil {

		web.InternalError.Send(w)
		return

	}

	web.Success(novels, 200).Send(w)

}
