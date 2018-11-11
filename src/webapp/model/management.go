package model

type Management struct {
	Q1  string `json:q1`
	Q2  string `json:q2`
	Q3  string `json:q3`
	Q4  string `json:q4`
	Q5  string `json:q5`
	Q6  string `json:q6`
	Q7  string `json:q7`
	Q8  string `json:q8`
	Q9  string `json:q9`
	Q10 string `json:q10`
	Reg string `json:reg,omitempty`
}

func RecordManager(m *Management) error {
	var k string
	row := db.QueryRow(`
		SELECT NAME
		FROM USER
		WHERE REG=$1
	`, m.Reg)
	err := row.Scan(&k)
	if err != nil {
		return err
	}
	_, err = db.Query(`
		INSERT INTO MANAGEMENT
		VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
	`, m.Q1, m.Q2, m.Q3, m.Q4, m.Q5, m.Q6, m.Q7, m.Q8, m.Q9, m.Q10, m.Reg)
	if err != nil {
		return err
	}
	return nil
}

func Showmanager(reg string) (*Management, error) {
	m := &Management{}
	row := db.QueryRow(`
		SELECT Q1, Q2, Q3, Q4, Q5, Q6, Q7, Q8, Q9, Q10, REG
		FROM MANAGEMENT
		WHERE REG=$1
	`, reg)
	err := row.Scan(&m.Q1, &m.Q2, &m.Q3, &m.Q4, &m.Q5, &m.Q6, &m.Q7, &m.Q8, &m.Q9, &m.Q10, &m.Reg)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func ShowAllmanager() ([]Management, error) {
	var arr []Management
	rows, err := db.Query(`
		SELECT Q1, Q2, Q3, Q4, Q5, Q6, Q7, Q8, Q9, Q10, REG
		FROM MANAGEMENT
	`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var m Management
		err = rows.Scan(&m.Q1, &m.Q2, &m.Q3, &m.Q4, &m.Q5, &m.Q6, &m.Q7, &m.Q8, &m.Q9, &m.Q10, &m.Reg)
		if err != nil {
			return nil, err
		}
		arr = append(arr, m)
	}
	return arr, nil
}
