package models

import (
	"example.com/evently-rest-api/db"
)

type Registration struct {
	ID       int64
	Event_ID int64
	User_ID  int64
}

func GetAllRegistrations(userId int64) ([]Registration, error) {
	query := `SELECT * FROM registrations
			WHERE user_id = ?`

	result, err := db.DB.Query(query, userId)

	if err != nil {
		return nil, err
	}

	defer result.Close()
	var registerations []Registration
	for result.Next() {
		var registeration Registration
		err := result.Scan(&registeration.ID, &registeration.Event_ID, &registeration.User_ID)

		if err != nil {
			return nil, err
		}
		registerations = append(registerations, registeration)
	}
	return registerations, nil

}
