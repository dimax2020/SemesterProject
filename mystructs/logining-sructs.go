package mystructs

import "time"

type UserStruct struct {
	ID        string      `json:"user_id"`
	Login     string      `json:"login"`
	Password  string      `json:"password"`
	Cookie    string      `json:"cookie"`
	ExpDate   time.Time
}