package helpers

func ResponseOk(a interface{}) ResponseApi {
	res := ResponseApi{
		Response: ResponseCode{
			Code: 0,
			Msg:  "success",
		},
		Data: a,
	}

	return res
}

type ResponseApi struct {
	Response ResponseCode `json:"response"`
	Data     interface{}  `json:"data"`
}

type ResponseCode struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}
