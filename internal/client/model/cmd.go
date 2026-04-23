package model

type DeleteArgs struct {
	ID string
}

type ListArgs struct{}

type LoginArgs struct {
	Login    string
	Password string
}

type RegisterArgs struct {
	Login    string
	Password string
}

type UploadAuthArgs struct {
	Login    string
	Password string
}

type UploadCardArgs struct {
	Name   string
	Number string
	CVV    string
}

type UploadFileArgs struct {
	Path string
}
