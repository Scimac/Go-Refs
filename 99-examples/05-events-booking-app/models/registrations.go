package models

import (
	"events-booking/db"
)

type Registration struct {
	ID      int64 `json:"id"`
	UserID  int64 `json:"user_id"`
	EventID int64 `json:"event_id"`
}

func GetAllRegistrations() ([]Registration, error) {
	query := "SELECT * FROM registrations"

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var regs []Registration

	for rows.Next() {
		var r Registration
		err := rows.Scan(&r.ID, &r.UserID, &r.EventID)
		if err != nil {
			return nil, err
		}
		regs = append(regs, r)
	}

	return regs, nil
}
