package main

import(
"fmt"
)

//Main is public
//main is protected

/*func main(){
//var name string or
name := "Hello world"
fmt.Println(name)
}*/

func main(){
    result := add(5,6)
    fmt.Println(result)
}


func add(a,b int) int{
    return a+b;
}
