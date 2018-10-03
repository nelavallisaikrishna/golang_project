package main

import(
"fmt"

)


//to create post request

type Request struct{

FirstName string `json : "first_name""`
LastName string  `json : "last_name""`

}

var chan1 chan string

func say(s string){

for i :=0; i < 5; i++{
/*fmt.Println(s)

if(i == 4){
    chan1 <- "done"
}*/
chan1 <- s
}

}


func main(){
   chan1 = make(chan string)
   go say("World")
   go say("hello")

   fmt.Println(<- chan1)
   fmt.Println(<- chan1)
   for{
    select{
        case success := <- chan1:
        fmt.Println(success)
    }

   }


}
