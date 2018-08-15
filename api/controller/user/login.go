package user

import (
	"log"
	"net/http"
	"simple_rest/database"
	"simple_rest/env"
  "fmt"
	"github.com/gin-gonic/gin"
)

type Login_Response struct {
	Code    int         `json:"Code"`
	Message string      `json:"Message"`
	Result  interface{} `json:"Result"`
}



func Login(c *gin.Context) {
  dbS := database.GetConn(env.AccountDB)
  res1 := Login_Response{}
	var check input
	res1.Result = nil

  //SQL Variables
  var sid string
  var saccount string
  var spassword string

  //User Variables
  var use_account string
  var use_password string

  fmt.Printf("Account: ")
  fmt.Scanf("%s",&use_account)
  fmt.Printf("New Password: ")
  fmt.Scanf("%s",&use_password)

  //To check if true or false
  var flag int
  flag=0


	// Query
	rows, err := dbS.Query("SELECT id , account , password FROM user")
	checkErr(err)
	for rows.Next() {
		err = rows.Scan(&sid,&saccount,&spassword)
		checkErr(err)
      if (use_account == saccount)&&(use_password == spassword){
          flag = 1
      }

	}
  //Check login validation
  if flag == 1{
    fmt.Println("Success")
    c.JSON(http.StatusOK, res1)
  } else {
    fmt.Println("Login Failed")
    res1.Message = "Fail"
    res1.Code = 2
		c.JSON(http.StatusBadRequest, res1)
		return
  }

	if err := c.Bind(&check); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, res1)
		return
	}


  dbS.Close()
}
