package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

//add jsondata
type JsonData struct{
  User string `json:"user" binding:"required"`
  Sql string `json:"sql" binding:"required"`
}

func main(){
  router := gin.Default()
  router.LoadHTMLGlob("templates/*")

  //Test with Hello
  router.GET("/", func(c *gin.Context) {
    c.String(http.StatusOK, "Hello World!")
  })

  //Get the table listed, it has different table in different OS/Distrubition
  router.GET("/tables", func(c *gin.Context){

    lines := getAlltables()
    var mesg []string
    for _, line := range lines{
      mesg = append(mesg, "/"+line)
    }

    c.HTML(http.StatusOK,"tables.tmpl", gin.H{
      "mesg" : mesg,
    })
  })

  //Get the table listed in the tables page
  router.GET("/tables/:table", func(c *gin.Context){
    table := c.Param("table")

    for _, line := range lines{
      if line == table {
        result := getOnetable(table)
        c.String(http.StatusOK, result)
      }
    }

  })

  //Accept from POST, use it if you are familiar with sql
  /*
    {
    "user" : "leitu"
    "sql" : "select * from kernel_info"
  }
  */

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
