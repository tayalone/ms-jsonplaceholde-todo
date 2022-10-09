package domain

import "time"

/*ToDo Represent Domain Problem Data */
type ToDo struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"userId:"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
	UpdateAt  time.Time `json:"updatedAt"`
}
