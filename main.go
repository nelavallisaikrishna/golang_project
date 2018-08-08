package main

import(
"strconv"
"github.com/gin-gonic/gin"
)


//to create post request

type Request struct{
FirstName string `json : "first_name""`
LastName string  `json : "last_name""`

}


func main(){
  r := gin.Default()
  r.POST("/add", func(c *gin.Context){
        var request Request
        c.Bind(&request)
        c.JSON(200, gin.H{
        "status": true,
        "value":"helloWorld",
        "name": request.FirstName+ " " + request.LastName,
        })

  })
  r.Run(":5432")
}


func add(a string,b int) int{
    i, err := strconv.Atoi(a)
    if err == nil {
    return i+b
    }else{
    return 0
    }
}
