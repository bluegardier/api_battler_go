package models

import (
	"api_battler/db"
)

func Insert(people People) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 
	}

	defer conn.Close()

	sql := `INSERT INTO people (apelido, nome, nascimento, stack) VALUES ($1, $2, $3, $4) RETURNING id`
	
	err = conn.QueryRow(sql, people.Apelido, people.Nome, people.Nascimento, people.Stack).Scan(&id)

	return
}

func GetAll() (people []People, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 
	}

	defer conn.Close()

	sql := `SELECT * FROM people`
	rows, err := conn.Query(sql)

	if err != nil {
		return 
	}

	for rows.Next() {
		var multiple_people People

		err = rows.Scan(&multiple_people.Apelido, &multiple_people.Nome, &multiple_people.Nascimento, &multiple_people.Stack)
		if err != nil {
			continue
		}

		people = append(people, multiple_people)

	}


	return
}

func Get(id int64) (people People, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 
	}

	defer conn.Close()

	sql := `SELECT * FROM people WHERE id = $1`
	err = conn.QueryRow(sql, id).Scan(&people.Apelido, &people.Nome, &people.Nascimento, &people.Stack)
	return

}

func Update(id int64, people People) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}

	defer conn.Close()

	sql := `UPDATE people SET apelido=$1, nome=$2, nascimento=$3, stack=$4`
	res, err := conn.Exec(sql, people.Apelido, people.Nome, people.Nascimento, people.Stack)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func Delete(id int64) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}

	defer conn.Close()

	sql := `DELETE FROM people WHERE id=$1`
	res, err := conn.Exec(sql, id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
 