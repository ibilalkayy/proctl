package mysql

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ibilalkayy/proctl/database/redis"
	"github.com/ibilalkayy/proctl/middleware"
)

// UpdateUser updates user information in the Signup table
func UpdateUser(value [4]string, email, password string) {
	db := Connect()
	defer db.Close()

	// Update AccountFullName in Redis if provided
	if len(value[0]) != 0 {
		redis.SetAccountInfo("AccountFullName", value[0])
	}

	// Update AccountName in Redis if provided
	if len(value[1]) != 0 {
		redis.SetAccountInfo("AccountName", value[1])
	}

	var hashPass string
	if len(value[2]) != 0 {
		redis.SetAccountInfo("AccountPassword", value[2])
		hashPass = middleware.HashPassword([]byte(value[2]))
		redis.DelToken("LoginToken")
		values := [4]string{redis.GetAccountInfo("AccountEmail"), hashPass, redis.GetAccountInfo("AccountFullName"), redis.GetAccountInfo("AccountName")}
		redis.SetUserCredentials(values)
	} else {
		hashPass = redis.GetAccountInfo("AccountPassword")
	}

	// Update user information in the Signup table
	q := "UPDATE Signup SET fullnames=?, accountnames=?, passwords=?, is_active=? WHERE emails=? AND passwords=?"
	_, err := db.Exec(q, redis.GetAccountInfo("AccountFullName"), redis.GetAccountInfo("AccountName"), hashPass, value[3], email, password)
	middleware.HandleError(err)

	redis.DelToken("LoginToken")
}

// UpdateProfile updates profile information in the Profiles table
func UpdateProfile(value [4]string, email string) {
	if len(value) == 0 && len(email) == 0 {
		return
	}

	db := Connect()
	q := "UPDATE Profiles SET "
	var updateValues []interface{}

	for i := 0; i < len(value); i++ {
		if len(value[i]) != 0 {
			switch i {
			case 0:
				q += "titles=?, "
			case 1:
				q += "phones=?, "
			case 2:
				q += "locations=?, "
			case 3:
				q += "working_statuses=?, "
			}
			updateValues = append(updateValues, value[i])
		}
	}

	q = strings.TrimSuffix(q, ", ")
	q += " WHERE emails=?"
	updateValues = append(updateValues, email)

	update, err := db.Prepare(q)
	middleware.HandleError(err)

	defer update.Close()

	_, err = update.Exec(updateValues...)
	middleware.HandleError(err)
}

// UpdateWorkspace updates workspace information in the Workspaces table
func UpdateWorkspace(value [3]string) {
	db := Connect()
	q := "UPDATE Workspaces SET names=? WHERE emails=? AND names=?"
	update, err := db.Prepare(q)
	middleware.HandleError(err)

	defer update.Close()

	if len(value[0]) != 0 && len(value[1]) != 0 && len(value[2]) != 0 {
		_, err = update.Exec(value[0], value[1], value[2])
		middleware.HandleError(err)
	} else {
		fmt.Println(errors.New("More flags are required to update the workspace"))
	}
}

// UpdateMember updates member information in the Members table
func UpdateMember(value [8]string, email, password string, isSet bool) {
	if len(value) == 0 && len(email) == 0 && (len(password) == 0 || isSet) {
		return
	}

	db := Connect()
	q := "UPDATE Members SET "
	var updateValues []interface{}
	var hashPass string

	for i := 0; i < len(value); i++ {
		if len(value[i]) != 0 {
			if isSet {
				switch i {
				case 0:
					q += "passwords=?, "
				case 1:
					q += "fullnames=?, "
				case 2:
					q += "accountnames=?, "
				}
				updateValues = append(updateValues, value[i])
			} else {
				switch i {
				case 0:
					q += "emails=?, "
					redis.SetAccountInfo("MemberAccountEmail", value[0])
					redis.DelToken("MemberLoginToken")
					values := [3]string{value[0], redis.GetAccountInfo("MemberAccountPassword"), redis.GetAccountInfo("MemberAccountName")}
					redis.SetMemberCredentials(values)
				case 1:
					q += "passwords=?, "
					hashPass = middleware.HashPassword([]byte(value[1]))
					redis.SetAccountInfo("MemberAccountPassword", hashPass)
					redis.DelToken("MemberLoginToken")
					values := [3]string{redis.GetAccountInfo("MemberAccountEmail"), hashPass, redis.GetAccountInfo("MemberAccountName")}
					redis.SetMemberCredentials(values)
				case 2:
					q += "fullnames=?, "
				case 3:
					q += "accountnames=?, "
					redis.SetAccountInfo("MemberAccountName", value[3])
					values := [3]string{redis.GetAccountInfo("MemberAccountEmail"), redis.GetAccountInfo("MemberAccountPassword"), value[3]}
					redis.SetMemberCredentials(values)
				case 4:
					q += "titles=?, "
				case 5:
					q += "phones=?, "
				case 6:
					q += "locations=?, "
				case 7:
					q += "working_statuses=?, "
				}
				if i != 1 {
					updateValues = append(updateValues, value[i])
				} else {
					updateValues = append(updateValues, hashPass)
				}
			}
		}
	}

	q += "is_active=? WHERE emails=?"
	if !isSet {
		q += " AND passwords=?"
		updateValues = append(updateValues, "1", email, password)
	} else {
		updateValues = append(updateValues, "1", email)
	}

	update, err := db.Prepare(q)
	middleware.HandleError(err)

	defer update.Close()

	_, err = update.Exec(updateValues...)
	middleware.HandleError(err)
}

// UpdateDepartment updates department information in the Departments table
func UpdateDepartment(email, dep string) {
	db := Connect()
	q := "UPDATE Departments SET departments=? WHERE emails=?"
	update, err := db.Prepare(q)
	middleware.HandleError(err)

	defer update.Close()

	if len(email) != 0 && len(dep) != 0 {
		_, err = update.Exec(dep, email)
		middleware.HandleError(err)
	} else {
		fmt.Println(errors.New("More flags are required to update the department"))
	}
}

// UpdateRole updates role information in the Roles table
func UpdateRole(email, role string) {
	db := Connect()
	q := "UPDATE Roles SET roles=? WHERE emails=?"
	update, err := db.Prepare(q)
	middleware.HandleError(err)

	defer update.Close()

	if len(email) != 0 && len(role) != 0 {
		_, err = update.Exec(role, email)
		middleware.HandleError(err)
	} else {
		fmt.Println(errors.New("More flags are required to update the role"))
	}
}

// UpdateBoard updates board information in the Boards table
func UpdateBoard(email, oldBoard, newBoard string) {
	db := Connect()
	q := "UPDATE Boards SET boards=? WHERE emails=? AND boards=?"
	update, err := db.Prepare(q)
	middleware.HandleError(err)

	defer update.Close()

	if len(email) != 0 && len(oldBoard) != 0 && len(newBoard) != 0 {
		_, err = update.Exec(newBoard, email, oldBoard)
		middleware.HandleError(err)
	} else {
		fmt.Println(errors.New("More flags are required to update the board"))
	}
}

// UpdateProject updates the project information in the Projects table.
func UpdateProject(value [3]string, boardName, projectName string) {
	if len(value) == 0 && len(boardName) == 0 && len(projectName) == 0 {
		return
	}

	db := Connect()
	q := "UPDATE Projects SET "
	var updateValues []interface{}

	for i := 0; i < len(value); i++ {
		if len(value[i]) != 0 {
			switch i {
			case 0:
				q += "projects=?, "
			case 1:
				q += "statuses=?, "
			case 2:
				q += "dates=?, "
			}
			updateValues = append(updateValues, value[i])
		}
	}

	q = strings.TrimSuffix(q, ", ")
	q += " WHERE boards=? AND projects=?"
	updateValues = append(updateValues, boardName, projectName)

	update, err := db.Prepare(q)
	middleware.HandleError(err)

	defer update.Close()

	_, err = update.Exec(updateValues...)
	middleware.HandleError(err)
}
