package rest

type ResponseBodySuccess struct {
	Ok   bool        `json:"ok"`
	Data interface{} `json:"data"`
	Ts   int64       `json:"ts"`
}

type ResponseBodyError struct {
	Ok  bool   `json:"ok"`
	Err string `json:"err"`
	Msg string `json:"msg"`
	Ts  int64  `json:"ts"`
}
