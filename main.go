package main

import (
  //"encoding/json"
  //"fmt"
  "os/exec"
  "github.com/gin-gonic/gin"
  "net/http"
  //"regexp"
  //"bytes"
  //"time"
)

type JsonData struct {
  User string `json:"user" binding:"required"`
  Sql string `json:"sql" binding:"required"`

}

func queryJason(sql string)string{

  bs, err := exec.Command("osqueryi", sql, "--json").Output()

  if err != nil {
    panic(err)
  }

  return string(bs)
}

func main(){
  router := gin.Default()
  router.LoadHTMLGlob("templates/*")

  router.GET("/", func(c *gin.Context) {
    c.String(http.StatusOK, "Hello World!")
  })

  router.GET("/tables", func(c *gin.Context){

    lines := getAlltables()
    var mesg []string
    for _, line := range lines{
      mesg = append(mesg, "/"+line)
    }
    //mesg  := []string{"1","2","3","4","5"}

    c.HTML(http.StatusOK,"tables.tmpl", gin.H{
      "mesg" : mesg,
    })
  })

  router.POST("/query", func(c *gin.Context){
    var jsondata JsonData
    c.BindJSON(&jsondata)

    results := queryJason(jsondata.Sql)


    //print with string way, Json way will get "/"
    c.String(http.StatusOK, results)
    //c.JSON(http.StatusOK, results)
 })

  router.Run(":8080")
}
