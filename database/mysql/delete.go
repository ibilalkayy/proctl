package mysql

import (
	"errors"
	"fmt"

	"github.com/ibilalkayy/proctl/middleware"
)

func DeleteWorkspace(value [2]string) {
	db := Connect()                                          // Connect to the MySQL database
	q := "DELETE FROM Workspaces WHERE emails=? AND names=?" // SQL query to delete a workspace
	delete, err := db.Prepare(q)                             // Prepare the SQL statement
	middleware.HandleError(err)                              // Handle any errors that occur

	defer delete.Close() // Close the prepared statement after the function ends

	if len(value[0]) != 0 && len(value[1]) != 0 {
		_, err = delete.Exec(value[0], value[1]) // Execute the prepared statement with the provided values
		middleware.HandleError(err)              // Handle any errors that occur during execution
	} else {
		fmt.Println(errors.New("Flags are required to delete the workspace")) // Print an error message if required values are missing
	}
}

func DeleteBoard(email, board string) {
	db := Connect()                                       // Connect to the MySQL database
	q := "DELETE FROM Boards WHERE emails=? AND boards=?" // SQL query to delete a board
	delete, err := db.Prepare(q)                          // Prepare the SQL statement
	middleware.HandleError(err)                           // Handle any errors that occur

	defer delete.Close() // Close the prepared statement after the function ends

	if len(email) != 0 && len(board) != 0 {
		_, err = delete.Exec(email, board) // Execute the prepared statement with the provided values
		middleware.HandleError(err)        // Handle any errors that occur during execution
	} else {
		fmt.Println(errors.New("Flags are required to delete the board")) // Print an error message if required values are missing
	}
}

func DeleteProject(email, board, project string) {
	db := Connect()                                                        // Connect to the MySQL database
	q := "DELETE FROM Projects WHERE emails=? AND boards=? AND projects=?" // SQL query to delete a board
	delete, err := db.Prepare(q)                                           // Prepare the SQL statement
	middleware.HandleError(err)                                            // Handle any errors that occur

	defer delete.Close() // Close the prepared statement after the function ends

	if len(email) != 0 && len(board) != 0 && len(project) != 0 {
		_, err = delete.Exec(email, board, project) // Execute the prepared statement with the provided values
		middleware.HandleError(err)                 // Handle any errors that occur during execution
	} else {
		fmt.Println(errors.New("Flags are required to delete the project")) // Print an error message if required values are missing
	}
}
