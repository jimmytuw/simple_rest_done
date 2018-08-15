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

type GetUserLogin struct{
	GetAccount string `form:GetAccount`
	GetPassword string `form:GetPassword`
}

func Login(c *gin.Context,) {
  dbS := database.GetConn(env.AccountDB)
  res1 := protocol.Response{}
	var check input
	res1.Result = nil
	i_input := &GetUserLogin{}


	if err := c.Bind(i_input); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, res1)
		return
	}

	var sid string
  var saccount string
  var spassword string

	// Query
	rows, err := dbS.Query("SELECT id , account , password FROM user")
	checkErr(err)
	for rows.Next() {
		err = rows.Scan(&sid,&saccount,&spassword)
		checkErr(err)
      if (i_input.GetAccount == saccount) && (i_input.GetPassword == spassword){
        fmt.Println("Login Successful")
        check.IsOk = true
        break
      }

	}

  if(check.IsOk == false){
		res1.Code  = 2
		res1.Message = "Failed"
    fmt.Println("Login Failed")
  }

	c.JSON(http.StatusOK, res1)
  dbS.Close()
}
