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

func ShowResponse() ([]User, error) {
	var arr []User
	rows, err := db.Query(`
		SELECT NAME, EMAIL, REG, APPLICANTTYPE
		FROM USER
	`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var result User
		err = rows.Scan(&result.Name, &result.Email, &result.Reg, &result.ApplicantType)
		if err != nil {
			return arr, err
		} else {
			arr = append(arr, result)
		}
	}
	return arr, nil

}

func ShowByReg(reg string) (*User, error) {
	result := &User{}
	row := db.QueryRow(`
		SELECT NAME, EMAIL, REG, APPLICANTTYPE
		FROM USER
		WHERE REG=$1
	`, reg)
	err := row.Scan(&result.Name, &result.Email, &result.Reg, &result.ApplicantType)
	if err != nil {
		return nil, err
	}
	return result, nil

}
