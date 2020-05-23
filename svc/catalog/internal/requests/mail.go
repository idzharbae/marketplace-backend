package requests

type SendMailReq struct {
	Destination string
	Title       string
	Message     string
}
