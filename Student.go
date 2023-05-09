package main

//Necessary packages are imported here.
import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// This block of code declares constant variables for the database connection parameters
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "0919"
	dbname   = "mydb"
)

// To check the error.
func CheckError(err error) {
	if err != nil {
		panic(err)
	}

}

// Main function
func main() {
	//To format a string connection.
	psqlconn := fmt.Sprintf("host = %s port= %d user =%s password =%s dbname = %s sslmode = disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	//To close the database.
	defer db.Close()

	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")
	//Inserting the values
	insertStmt := `insert into "Student1"("Rollno", "Name", "DOB", "Age", "Gender") values('3','Danush','2010-05-23','13','Male')`
	_, e := db.Exec(insertStmt)
	CheckError(e)

	insertStmt2 := `insert into "Student1"("Rollno", "Name", "DOB", "Age", "Gender") values($1,$2,$3,$4,$5)`
	_, e = db.Exec(insertStmt2, 4, "Diviya", "2010-05-19", 13, "Female")
	CheckError(e)

	fmt.Println("The values are inserted.")

	// //Update the values
	// updateStmt := `UPDATE "Student" SET "DOB" = '2010-04-20' WHERE "Rollno" ='3'`
	// _, e = db.Exec(updateStmt)
	// CheckError(e)
	// fmt.Println("The value is updated")

	// //Delete a value from the table
	// deleteStmt := `DELETE FROM "Student" WHERE "Rollno" = '3'`
	// _, e = db.Exec(deleteStmt)
	// CheckError(e)
	// fmt.Println("The value is deleted sucessfully")
	// //Read the values from the table
	// readstmt := `select * from "Student"`
	// rows, err := db.Query(readstmt)
	// CheckError(err)

	// defer rows.Close()

	// // Print the rows
	// for rows.Next() {
	// 	var rollno int
	// 	var name string
	// 	var dob string
	// 	var age int
	// 	var gender string

	// 	err = rows.Scan(&rollno, &name, &dob, &age, &gender)
	// 	CheckError(err)

	// 	fmt.Printf("Rollno: %d, Name: %s\t, DOB: %s, Age: %d, Gender: %s\n", rollno, name, dob, age, gender)
	// }

	// err = rows.Err()
	// CheckError(err)

}
