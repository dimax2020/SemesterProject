package logining

import (
	"SemesterProject/bd"
	"SemesterProject/mystructs"
	"time"
)


func GenerateNewID() string {
	return "1"
}

func AddNewUser(login string, password string) mystructs.UserStruct {
	var user mystructs.UserStruct
	user.Cookie, user.ExpDate = GenerateNewCookie()
	user.ID = GenerateNewID()
	user.Password = password
	user.Login = login
	// bd
	return user
}

func CheckLoginAndPassword(login string, password string) (bool, string, time.Time) {
	 if bd.CheckLoginAndPassword(login, password) {
		 cookie, expDate := GenerateNewCookie()
		 return true, cookie, expDate
	 }
	 return false, "nil", time.Now()
}
