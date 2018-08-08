package main

import(
"fmt"
"strconv"
)

//Main is public
//main is protected

/*func main(){
//var name string or
name := "Hello world"
fmt.Println(name)
}*/

func main(){
    a := "5"
    b := 6
    result := add(a,b)
    fmt.Println(result)
}


func add(a string,b int) int{
    i, err := strconv.Atoi(a)
    if err == nil {
    return i+b
    }else{
    return 0
    }
}
