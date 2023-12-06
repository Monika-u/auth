package db

import (
	"demo/resources"
	"fmt"
)

func CreateUser(input resources.User) error {

	db := DbClient
	query := "INSERT INTO `users` (`name`,`email_id`, `password`,`role`) VALUES (?,?, ?,?)"
	_, err := db.Exec(query, input.Name, input.EmailId, input.Password, input.Role)
	if err != nil {
		fmt.Println("err", err)
		return err
	}
	return nil

}

func GetUser(input resources.User) (resources.User, error) {

	db := DbClient
	var result resources.User
	query := "select * from users where email_id=?"
	err := db.QueryRow(query, input.EmailId).Scan(&result.UserId, &result.Name, &result.Password, &result.Role, &result.EmailId)
	if err != nil {
		return result, err
	}

	return result, nil

}

func GetUsers() ([]resources.User, error) {

	db := DbClient
	var result []resources.User
	// var userInfo resources.User
	query := "select * from users where role = ?"
	// err := db.QueryRow(query, "user")
	rows, err := db.Query(query, "user")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var userInfo resources.User
		if err := rows.Scan(&userInfo.UserId, &userInfo.Name, &userInfo.Password, &userInfo.Role, &userInfo.EmailId); err != nil {
			return nil, err
		}
		result = append(result, userInfo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil

}

func GetUserByName(username string) (resources.User, error) {

	db := DbClient
	var result resources.User
	query := "select * from users where name=?"
	err := db.QueryRow(query, username).Scan(&result.UserId, &result.Name, &result.Password, &result.Role, &result.EmailId)
	if err != nil {
		return result, err
	}

	return result, nil
}
