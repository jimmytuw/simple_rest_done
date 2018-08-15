package user

import (
	"log"
	"net/http"
	"simple_rest/api/protocol"
	"simple_rest/database"
	"simple_rest/env"
  "fmt"
	"github.com/gin-gonic/gin"
)

type UserInput struct {
	account     string    `form:"Account" binding:"exists"`
	password  string    `form:"Password" binding:"exists"`
}


func Change(c *gin.Context,) {
  dbS := database.GetConn(env.AccountDB)
  res1 := protocol.Response{}
	var check input
	res1.Result = &check

  //SQL Variables
  var sid string
  var saccount string
  var spassword string

	//Fetching from form
	user := c.PostForm("Account")
	password := c.PostForm("Password")


	// Query
	rows, err := dbS.Query("SELECT id , account , password FROM user")
	checkErr(err)
	for rows.Next() {
		err = rows.Scan(&sid,&saccount,&spassword)
		checkErr(err)
      if user == saccount{
        stmt, err := dbS.Prepare("UPDATE user SET password=? WHERE account=?")
        checkErr(err)
        res, err := stmt.Exec(password,user)
        checkErr(err)
        res.LastInsertId()
        fmt.Println("Update Completed!")
        check.IsOk = true
        break
      }

	}

  if(check.IsOk == false){
    fmt.Println("Update Failed")
  }


	if err := c.Bind(&check); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, res1)
		return
	}


	c.JSON(http.StatusOK, res1)

  dbS.Close()
}
