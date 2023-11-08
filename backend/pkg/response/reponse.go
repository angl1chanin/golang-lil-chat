package response

type Response struct {
	Status string            `json:"status"`
	Error  string            `json:"error,omitempty"`
	Info   map[string]string `json:"info,omitempty"`
}

const (
	StatusOK      = "OK"
	StatusError   = "error"
	StatusSuccess = "success"
	StatusEmpty   = "empty"
)

func Error(msg string) Response {
	return Response{
		Status: StatusError,
		Error:  msg,
	}
}

func Success() Response {
	return Response{
		Status: StatusSuccess,
	}
}

func SuccessWithInfo(info map[string]string) Response {
	return Response{
		Status: StatusSuccess,
		Info:   info,
	}
}

func Empty() Response {
	return Response{
		Status: StatusEmpty,
	}
}
