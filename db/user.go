package db

import (
	"demo/resources"
)

func CreateUser(input resources.User) error {

	db := DbClient
	query := "INSERT INTO `users` (`name`,`email_id`, `password`) VALUES (?,?, ?)"
	_, err := db.Exec(query, input.Name, input.EmailId, input.Password)
	if err != nil {
		return err
	}
	return nil

}

func GetUser(input resources.User) (resources.User, error) {

	db := DbClient
	var result resources.User
	query := "select * from users where email_id=?"
	err := db.QueryRow(query, input.EmailId).Scan(&result.UserId, &result.Name, &result.EmailId, &result.Password)
	if err != nil {
		return result, err
	}

	return result, nil

}
