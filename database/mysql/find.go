package mysql

type UserCredentials struct {
	Email    string
	Password string
}

func FindAccount(email, password string) (string, string, bool) {
	db := Connect()
	var uc UserCredentials
	q := "SELECT emails, passwords FROM Signup WHERE emails=? and passwords=?"
	if err := db.QueryRow(q, email, password).Scan(&uc.Email, &uc.Password); err != nil {
		return "", "", false
	}

	return email, password, true
}
