package helpers

import (
	"encoding/json"
	"net/http"
)

func ResponseOk(w http.ResponseWriter, a interface{}) {
	res := ResponseApi{
		Response: ResponseCode{
			Code: 0,
			Msg:  "success",
		},
		Data: a,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func ResponseError(w http.ResponseWriter, code int64, msg string) {
	res := ResponseApi{
		Response: ResponseCode{
			Code: code,
			Msg:  msg,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(res)
}

type ResponseApi struct {
	Response ResponseCode `json:"response"`
	Data     interface{}  `json:"data"`
}

type ResponseCode struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}
