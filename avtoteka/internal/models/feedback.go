package models

type Feedback struct {
	ID         int    `json:"id"`
	UserID     int    `json:"userId"`
	CarID      int    `json:"cardId"`
	CreateTime int64  `json:"createTime"`
	Rating     int    `json:"rating"`
	Comment    string `json:"comment"`
}
