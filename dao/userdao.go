package dao

import "goland/global"

type user struct {
	ID int `json:"id"`

	Name string `json:"name"`

	Sex string `json:"sex"`

	Reamrk string `json:"reamrk"`
}

func GetAllUsers() []user {

	rows, err := global.Mysql.Query("select * from member")

	if err != nil {

		return nil

	}

	var persons = make([]user, 0)

	for rows.Next() {

		var a user

		err := rows.Scan(&a.ID, &a.Name, &a.Sex, &a.Reamrk)

		if err != nil {

			return nil

		}

		persons = append(persons, a)

	}

	return persons

}
