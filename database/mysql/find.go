package mysql

type UserCredentials struct {
	Email    string
	Password string
	Status   string
}

func FindAccount(email, password string) (string, string, string, bool) {
	db := Connect()
	var uc UserCredentials
	q := "SELECT emails, passwords, is_active FROM Signup WHERE emails=? and passwords=?"
	if err := db.QueryRow(q, email, password).Scan(&uc.Email, &uc.Password, &uc.Status); err != nil {
		return "", "", "", false
	}

	return uc.Email, uc.Password, uc.Status, true
}
