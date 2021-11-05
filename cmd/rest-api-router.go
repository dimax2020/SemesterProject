package cmd

import (
	"SemesterProject/logining"
	"SemesterProject/mystructs"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func signIn(c *gin.Context) {
	var requestJson mystructs.SignInStruct

	if c.BindJSON(&requestJson) == nil {

		if len(requestJson.Login) > 4 {

			if len(requestJson.Password) > 8 {

				user := logining.AddNewUser(requestJson.Login, requestJson.Password)
				c.SetCookie(
					"Authorization",
					user.Cookie,
					int(time.Since(user.ExpDate).Seconds()),
					"/",
					"domain.ru",
					false,
					true,
				)
				c.JSON(http.StatusOK, "Registered and authorized")

			} else {
				c.JSON(http.StatusBadRequest, "Too small password")
			}

		} else {
			c.JSON(http.StatusBadRequest, "Too small login")
		}

	} else {
		c.JSON(http.StatusBadRequest, "Bad json")
	}
}

func logIn(c *gin.Context) {
	var requestJson mystructs.LogInStruct

	if c.BindJSON(&requestJson) == nil {

		flag, cookie, expDate := logining.CheckLoginAndPassword(requestJson.Login, requestJson.Password)

		if flag {
			c.SetCookie(
				"Authorization",
				cookie,
				int(time.Since(expDate).Seconds()),
				"/",
				"domain.ru",
				false,
				true,
			)
			c.JSON(http.StatusOK, "Authorized")

		} else {
			c.JSON(http.StatusBadRequest, "Invalid login or password")
		}

	} else {
		c.JSON(http.StatusBadRequest, "Bad json")
	}
}

func addFile(c *gin.Context) {

}

func renameFile(c *gin.Context) {

}

func getFileByID(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "1")
}

func getMyFiles(c *gin.Context) {
	c.JSON(http.StatusOK, "1, 2, 3, 4, 5")
}

func deleteFileByID(c *gin.Context) {

}

func deleteAllFiles(c *gin.Context) {

}

func StartRestAPI() {
	router := gin.Default()

	router.POST  ("/sign%20in/", signIn        ) // sign in
	router.POST  ("/log%20in/",  logIn         ) // log in
	router.POST  ("/files/",     renameFile    ) // rename
	router.GET   ("/files/:id",  getFileByID   ) // get one file
	router.GET   ("/files/",     getMyFiles    ) // get all files
	router.DELETE("/files/:id",  deleteFileByID) // delete one file
	router.DELETE("/files/",     deleteAllFiles) // delete all files
	router.PUT   ("/files/",     addFile       ) // add file

	fmt.Println("[API][cmd/rest-api-router] Router error:", router.Run("localhost:56655"))
}

// middle ware
// loger
// panic, recover
// interface