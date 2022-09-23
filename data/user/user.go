package user

import (
	// "encoding/json"
	"fmt"
	"strconv"

	DB "github.com/rest_api/DB"
	Models "github.com/rest_api/Models"
	Token "github.com/rest_api/http/token"
)

func Insert(user Models.User) bool {
	db := DB.Initialize()
	res, err := db.Prepare(`INSERT INTO users(name, email, password, phone, username) VALUES (?,?,?,?)`)
	if err != nil {
		fmt.Println(err)
		return false
	}
	user_pass, _ := Token.HashPassword(user.Password)
	res.Exec(user.Name, user.Email, user_pass, user.Phone, user.Username)
	defer db.Close()
	return true
}

func GetAll() []interface{} {
	db := DB.Initialize()
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}
	users := []interface{}{}
	for rows.Next() {
		var id int
		var name, email, username, phone string
		err := rows.Scan(&id, &name, &email, &username, &phone)
		if err != nil {
			panic(err.Error())
		}
		user := Models.User{ID: id, Name: name, Username: username, Email: email, Phone: phone}
		user_encoded := map[string]string{"Name": user.Name, "Email": user.Email, "Username": user.Username, "Phone": user.Phone, "ID": strconv.Itoa(user.ID)}
		users = append(users, user_encoded)
	}
	defer db.Close()
	return users
}

func Update(user Models.User) (bool, string) {
	db := DB.Initialize()
	query := "SELECT * FROM users WHERE id=" + strconv.Itoa(user.ID)

	if DB.RowExists(query) {
		_, err := db.Exec(`UPDATE users SET name=?, email=?, phone=? WHERE id=?`, user.Name, user.Email, user.Phone, user.ID)
		if err != nil {
			defer db.Close()
			return false, err.Error()
		}
		defer db.Close()
		return true, "Updated"
	}
	defer db.Close()
	return true, `Does not exist`
}

func Remove(id int) (bool, string) {
	db := DB.Initialize()
	query := "SELECT * FROM users where id=" + strconv.Itoa(id)

	if DB.RowExists(query) {
		_, err := db.Exec(`DELETE FROM users WHERE id=?`, id)
		if err != nil {
			defer db.Close()
			return false, err.Error()
		}
		defer db.Close()
		return true, ""
	}
	defer db.Close()
	return false, "Doesn't Exist!"
}
