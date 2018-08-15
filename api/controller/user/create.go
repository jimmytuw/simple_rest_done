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

type input struct{
  IsOk bool
}


func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}



func Create(c *gin.Context) {
  dbS := database.GetConn(env.AccountDB)
  res1 := protocol.Response{}
	var check input
	res1.Result = &check

  //SQL Variables
  var sid string
  var saccount string
  var spassword string

  //User Variables
  var use_account string
  var use_password string

  //Getting Information
  fmt.Printf("Account: ")
  fmt.Scanf("%s",&use_account)
  fmt.Printf("Password: ")
  fmt.Scanf("%s",&use_password)

  var flag int
  flag = 0

	// Query
	rows, err := dbS.Query("SELECT id , account , password FROM user")
	checkErr(err)
	for rows.Next() {
		err = rows.Scan(&sid,&saccount,&spassword)
		checkErr(err)
      if use_account == saccount{
        flag = 0
        break
      } else {
        flag = 1
      }

	}
  //Check if same account is found in database
  if flag == 1{
    stmt, err := dbS.Prepare("INSERT user SET account=?, password=?")
    checkErr(err)
    res, err := stmt.Exec(use_account,use_password)
    checkErr(err)
    id, err := res.LastInsertId()
    checkErr(err)
    fmt.Println("Insert Completed!",id)

    check.IsOk = true
  } else {
    fmt.Println("Error, account found..")
  }

  // 綁定Input參數至結構中
	if err := c.Bind(&check); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, res1)
		return
	}


	c.JSON(http.StatusOK, res1)

  dbS.Close()
}
