package main

import (
	"database/sql"
	"fmt"

	_ "github.com/godror/godror"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("godror", `user="electricity_bill" password="" connectString="10.0.0.116:1521/pdb1.gerege.local"`)
	if err != nil {
		fmt.Println("... Oracle DB Setup Failed")
		fmt.Println(err)
		return
	}

	defer db.Close()

	fmt.Println("... Opening Oracle Database Connection")
	if err = db.Ping(); err != nil {
		fmt.Println("Error connecting to the oracle database: ", err)
		return
	}
	fmt.Println("... Connected to Oracle Database")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", "127.0.0.1", 5432, "root", "secret", "p2p_loan")
	db1, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db1.Close()

	err = db1.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("... Postgres Database Successfully connected!\t")

	// rows, err := db.Query(` SELECT
	// 							u.id,
	// 							u.reg_no,
	// 							u.phone,
	// 							u.email,
	// 							u.address,
	// 							u.picture,
	// 							u.is_xyp_sync,
	// 							u.expire_date,
	// 							u.issue_date,
	// 							u.aimag_code,
	// 							u.sum_code,
	// 							u.bag_code,
	// 							u.civil_id,
	// 							u.first_name,
	// 							u.last_name,
	// 							u.family_name,
	// 							u.birth_date,
	// 							u.gender,
	// 							p.password
	// 						FROM
	// 							iam_users          u,
	// 							iam_user_passwords p
	// 						WHERE
	// 								u.id = p.user_id
	// 							AND is_phone_activated = 1
	// 						ORDER BY
	// 							id ASC
	// 						FETCH FIRST 1000 ROWS ONLY`)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()

	// type User struct {
	// 	id          string
	// 	reg_no      string
	// 	phone       int
	// 	email       string
	// 	address     string
	// 	picture     string
	// 	is_xyp_sync int
	// 	expire_date pq.NullTime
	// 	issue_date  pq.NullTime
	// 	aimag_code  int
	// 	sum_code    int
	// 	bag_code    sql.NullInt32
	// 	civil_id    int
	// 	first_name  string
	// 	last_name   string
	// 	family_name string
	// 	birth_date  pq.NullTime
	// 	gender      int
	// 	password    string
	// }
	// var u User
	// for rows.Next() {

	// 	if err := rows.Scan(&u.id, &u.reg_no, &u.phone, &u.email, &u.address, &u.picture, &u.is_xyp_sync, &u.expire_date, &u.issue_date, &u.aimag_code, &u.sum_code, &u.bag_code,
	// 		&u.civil_id, &u.first_name, &u.last_name, &u.family_name, &u.birth_date, &u.gender, &u.password); err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	sqlStatement := `INSERT INTO gerege_app.old_users( id, reg_no, phone, email, address, picture, is_xyp_sync, expire_date, issue_date, aimag_code, sum_code, bag_code,
	// 		password, civil_id, first_name, last_name, family_name, birth_date, gender)
	// 		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9,$10, $11, $12, $13, $14, $15, $16, $17, $18, $19)`
	// 	_, err = db1.Exec(sqlStatement, u.id, u.reg_no, u.phone, u.email, u.address, u.picture, u.is_xyp_sync, u.expire_date, u.issue_date, u.aimag_code, u.sum_code, u.bag_code,
	// 		u.password, u.civil_id, u.first_name, u.last_name, u.family_name, u.birth_date, u.gender)
	// 	if err != nil {
	// 		panic(err)
	// 	} else {

	// 	}

	// 	// fmt.Println(id, reg_no, phone, email, address, picture, is_xyp_sync, expire_date, issue_date, aimag_code, sum_code, bag_code, civil_id, first_name, last_name, family_name, birth_date, gender, password)
	// }
	// fmt.Println("\nRow inserted successfully!")
	// if err := rows.Err(); err != nil {
	// 	log.Fatal(err)
	// }

}
