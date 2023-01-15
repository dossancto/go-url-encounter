package database

func CreateUrl(url string, code string) error {
	res, err := database.Exec("INSERT into url (original, code) VALUES ($1, $2)", url, code)

	if err != nil {
		return err
	}

	_, err = res.RowsAffected()

	if err != nil {
		return err
	}

	return nil
}

func GetUrl(code string) (string, error) {
	stmt, err := database.Prepare("SELECT original FROM url WHERE code = $1")

	var url string

	if err != nil {
		return "", err
	}

	defer stmt.Close()

	err = stmt.QueryRow(code).Scan(&url)

	if err != nil {
		return "", err
	}

	return url, nil
}
