package main

import(
"fmt"
"time"
)


//to create post request

type Request struct{

FirstName string `json : "first_name""`
LastName string  `json : "last_name""`

}


func say(s string){

for i :=0; i < 5; i++{
time.Sleep(100*time.Millisecond)
fmt.Println(s)
}

}


func main(){
   go say("Hello World")
   say("hello")
}
