package models

import (
	"log"

	db "../database"
)

type Person struct {
	Id        int    `json:"id" form:"id"`
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
}

// 添加用户

func (p *Person) AddPerson() (id int64, err error) {

	rs, err := db.SqlDB.Exec("INSERT INTO person(first_name,last_name) VALUES (?,?)", p.FirstName, p.LastName)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}

//  获取所有的用户

func (p *Person) GetPersons() (persons []Person, err error) {
	persons = make([]Person, 0)
	rows, err := db.SqlDB.Query("SELECT id,first_name,last_name FROM person")

	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var person Person
		rows.Scan(&person.Id, &person.FirstName, &person.LastName)
		persons = append(persons, person)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}

// 获取单个用户的信息

func (p *Person) GetPerson() (person Person, err error) {
	err = db.SqlDB.QueryRow("SELECT id,first_name,last_name FROM person WHERE id=?", p.Id).Scan(
		&person.Id, &person.FirstName, &person.LastName,
	)
	return
}

// 修改单个用户的信息

func (p *Person) ModPerson() (ra int64, err error) {
	stmt, err := db.SqlDB.Prepare("UPDATE person SET first_name=?,last_name=? WHERE id=?")
	defer stmt.Close()
	if err != nil {
		return
	}
	rs, err := stmt.Exec(p.FirstName, p.LastName, p.Id)
	if err != nil {
		return
	}
	ra, err = rs.RowsAffected()
	return
}

// 删除单个用户的信息

func (p *Person) DelPerson() (ra int64, err error) {
	rs, err := db.SqlDB.Exec("DELETE FROM person WHERE id=?", p.Id)
	if err != nil {
		log.Fatalln(err)
	}
	ra, err = rs.RowsAffected()
	return
}
