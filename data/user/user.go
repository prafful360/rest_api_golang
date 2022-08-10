package user

import (
	"strconv"
	DB "github.com/rest_api/DB"
	Models "github.com/rest_api/Models"
	"fmt"
)


func Insert(user Models.User) {
	db := DB.Initialize()
	res, err := db.Prepare(`INSERT INTO users(name, email, phone, username) VALUES (?,?,?,?)`)
	if err != nil {
		fmt.Println(err)
	}
	res.Exec(user.Name, user.Email, user.Phone, user.Username)
	defer db.Close()
}

func GetAll() []Models.User {
	db := DB.Initialize()
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}
	users := []Models.User{}
	for rows.Next() {
		var id int
		var name, email, username, phone string
		err := rows.Scan(&id, &name, &email, &phone, &username)
		if err != nil {
			panic(err.Error())
		}
		user := Models.User{ID: id, Name: name, Username: email, Email: username, Phone: phone}
		users = append(users, user)
	}
	defer db.Close()
	return users
}

func Update(id int) string {
	db := DB.Initialize()
	query := "SELECT * FROM users WHERE id=" + strconv.Itoa(id)

	if DB.RowExists(query) {
		_, err := db.Exec(`UPDATE users SET name=?, phone=? WHERE id=?`, "UpdatedName", "00000", id)
		if err != nil {
			defer db.Close()
			return "Something went wrong!"
		}
		defer db.Close()
		return "Updated"
	}
	defer db.Close()
	return `Does not exist`
}


