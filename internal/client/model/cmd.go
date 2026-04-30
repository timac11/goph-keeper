package model

type DeleteArgs struct {
	ID string
}

type ListArgs struct{}

type GetArgs struct {
	ID string
}

type LoginArgs struct {
	Login    string
	Password string
}

type RegisterArgs struct {
	Login    string
	Password string
}

type UploadAuthArgs struct {
	Login      string `json:"login"`
	Password   string `json:"password"`
	Metadata   string `json:"-"`
	UploadName string `json:"-"`
}

type UploadCardArgs struct {
	Name       string `json:"name"`
	Number     string `json:"number"`
	CVV        string `json:"cvv"`
	Metadata   string `json:"-"`
	UploadName string `json:"-"`
}

type UploadFileArgs struct {
	Path     string
	Metadata string
}
