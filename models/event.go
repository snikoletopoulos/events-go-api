package models

import (
	"time"

	"events-rest-api/db"
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      uint
	User        User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" binding:""`
}

func GetAllEvents() ([]Event, error) {
	var events []Event
	result := db.DB.Find(&events)
	if result.Error != nil {
		return nil, result.Error
	}
	return events, nil
}

func FindEventByID(id uint) (*Event, error) {
	var event Event
	result := db.DB.Find(&event, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &event, nil
}

func (event *Event) Save() error {
	result := db.DB.Create(&event)
	return result.Error
}

func (event Event) Update() error {
	result := db.DB.Save(&event)
	return result.Error
}

func (event Event) Delete() error {
	result := db.DB.Delete(&event)
	return result.Error
}

// func (event Event) Register(userID int64) error {
// 	statement, err := db.DB.Prepare(`
// 		INSERT INTO registrations (event_id, user_id) VALUES (?, ?)
// 	`)
// 	if err != nil {
// 		return err
// 	}
// 	defer statement.Close()
//
// 	_, err = statement.Exec(event.ID, userID)
// 	return err
// }
//
// func (event Event) CancelRegistration(userID int64) error {
// 	statement, err := db.DB.Prepare(`
// 		DELETE FROM registrations WHERE event_id = ? AND user_id = ?
// 	`)
// 	if err != nil {
// 		return err
// 	}
// 	defer statement.Close()
//
// 	_, err = statement.Exec(event.ID, userID)
// 	return err
// }
