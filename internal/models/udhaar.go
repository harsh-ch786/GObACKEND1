package models

import "time"


type Udhaar struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	UserID      string    `json:"userId" bson:"userId"`
	FriendName  string    `json:"friendName" bson:"friendName"`
	Amount      float64   `json:"amount" bson:"amount"`
	Description string    `json:"description" bson:"description"`
	Status      string    `json:"status" bson:"status"`
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
	DueDate     time.Time `json:"dueDate" bson:"dueDate"`
}
