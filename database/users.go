package database

import (
	"fmt"
	typesv1 "github.com/x0y14/jackal/gen/types/v1"
)

func CreateUser(user *typesv1.User) error {
	res, err := database.Exec(`
		insert into users (user_id, display_name)
		select ?, ? where not exists(select * from users where user_id=?)`, user.UserId, user.DisplayName, user.UserId)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return fmt.Errorf("already exists: %v", user.UserId)
	}

	return nil
}

func GetUser(userId string) (*typesv1.User, error) {
	stmt, err := database.Prepare("select user_id, display_name from users where user_id = ?")
	if err != nil {
		return nil, err
	}
	var userId_ string
	var displayName_ string
	err = stmt.QueryRow(userId).Scan(&userId_, &displayName_)
	if err != nil {
		return nil, err
	}

	return &typesv1.User{
		UserId:      userId_,
		DisplayName: displayName_,
	}, nil
}
