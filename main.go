package main

import(
"strconv"
"github.com/gin-gonic/gin"
)


func main(){
  r := gin.Default()
  r.GET("/add", func(c *gin.Context){
        c.JSON(200, gin.H{
        "status": true,
        "value":"helloWorld",
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
