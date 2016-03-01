package main

import (
  //"encoding/json"
  //"fmt"
  "os/exec"
  "github.com/gin-gonic/gin"
  "net/http"
  //"bytes"
  //"time"
)

type JsonData struct {
  User string `json:"user" binding:"required"`
  Sql string `json:"sql" binding:"required"`

}

func queryJason(sql string)[]byte{

  bs, err := exec.Command("osqueryi", sql, "--json").Output()

  if err != nil {
    panic(err)
  }

  return bs
}

func main(){
  router := gin.Default()

  router.GET("/", func(c *gin.Context) {
    c.String(http.StatusOK, "Hello World!")
  })

  router.POST("/query", func(c *gin.Context){
    var jsondata JsonData
   c.BindJSON(&jsondata)

   results := queryJason(jsondata.Sql)


   c.String(http.StatusOK, string(results))
   //c.JSON(http.StatusOK,gin.H{"message": string(a)})
 })

  router.Run(":8080")
}
