package response

type FileResponse struct {
	FileName string `json:"fileName"`
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
}
