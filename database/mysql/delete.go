package mysql

import (
	"errors"
	"fmt"

	"github.com/ibilalkayy/proctl/middleware"
)

func DeleteWorkspace(value [2]string) {
	db := Connect()
	q := "DELETE FROM Workspaces WHERE emails=? AND names=?"
	delete, err := db.Prepare(q)
	middleware.HandleError(err)

	defer delete.Close()

	if len(value[0]) != 0 && len(value[1]) != 0 {
		_, err = delete.Exec(value[0], value[1])
		middleware.HandleError(err)
	} else {
		fmt.Println(errors.New("Flags are required to delete the workspace"))
	}
}
