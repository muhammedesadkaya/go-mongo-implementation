package model

import "time"

type User struct {
	FullName    string    `json:"fullName" bson:"FullName"`
	Identity    string    `json:"identity" bson:"Identity"`
	PhoneNumber string    `json:"phoneNumber" bson:"PhoneNumber"`
	Email       string    `json:"email" bson:"Email"`
	Gender      string    `json:"gender" bson:"Gender"`
	IsActive    bool      `json:"isActive" bson:"IsActive"`
	IsDelete    bool      `json:"isDelete" bson:"IsDelete"`
	CreateDate  time.Time `json:"createDate" bson:"CreateDate"`
}
