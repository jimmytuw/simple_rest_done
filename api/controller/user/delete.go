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




func Delete(c *gin.Context) {
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

  //Getting Informations
  fmt.Printf("Account: ")
  fmt.Scanf("%s",&use_account)


	// Query
	rows, err := dbS.Query("SELECT id , account , password FROM user")
	checkErr(err)
	for rows.Next() {
		err = rows.Scan(&sid,&saccount,&spassword)
		checkErr(err)
      if use_account == saccount{
        stmt, err := dbS.Prepare("DELETE FROM user where account=?")
        checkErr(err)
        res, err := stmt.Exec(use_account)
        checkErr(err)
        id, err := res.LastInsertId()
        checkErr(err)
        fmt.Println("Delete Completed!",id)
        check.IsOk = true
        break
      }

	}

  //If failed
  if(check.IsOk == false){
    fmt.Println("Delete Failed")
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
