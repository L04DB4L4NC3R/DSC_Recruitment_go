package model

type User struct {
	Name          string `json:name,omitempty`
	Email         string `json:email,omitempty`
	Reg           string `json:reg,omitempty`
	ApplicantType string `json:applicantType`
}

func RecordResponse(u *User) (*User, error) {
	return u, nil
}

func ShowResponse() (*User, error) {
	return nil, nil
}

func ShowByReg(reg string) (*User, error) {
	return nil, nil
}
