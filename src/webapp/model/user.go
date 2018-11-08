package model

import (
	"database/sql"
	"fmt"
)

type User struct {
	Name          string `json:name,omitempty`
	Email         string `json:email,omitempty`
	Reg           string `json:reg,omitempty`
	ApplicantType string `json:applicantType`
}

func RecordResponse(u *User) error {

	// check to see if record exists or not
	var data string
	row := db.QueryRow(`
		SELECT NAME
		FROM USER
		WHERE REG=$1
	`, u.Reg)
	err := row.Scan(&data)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if len(data) > 0 {
		return fmt.Errorf("Already exists")
	}
	// if doesnt exist then proceed
	_, err = db.Exec(`
		INSERT INTO USER(NAME, EMAIL, REG, APPLICANTTYPE)
		VALUES($1,$2,$3,$4)
	`, u.Name, u.Email, u.Reg, u.ApplicantType)
	if err != nil {
		return err
	}
	return nil
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

func ShowTypeResponse(param string) ([]User, error) {
	var arr []User
	rows, err := db.Query(`
	
		SELECT NAME, EMAIL, REG, APPLICANTTYPE
		FROM USER
		WHERE APPLICANTTYPE=$1
	`, param)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var result User
		err = rows.Scan(&result.Name, &result.Email, &result.Reg, &result.ApplicantType)
		switch {
		case err == sql.ErrNoRows:
			return nil, fmt.Errorf("No rows found")
		case err != nil:
			return nil, err
		default:
			arr = append(arr, result)
			break
		}
	}
	return arr, nil
}
