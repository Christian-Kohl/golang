package main

import (
    "fmt"
)

//different ways to declare variables

//global variable declaration
var (

    pklvl string = "foo"

)

func main() {
    // Separate declaration and assignment
    var j float32
    // Separate assignment
    j = 27
    //Combined declaration and assignment, with inferred type
    i  := 42
    //Combined declaration and assignment
    var k int = 12
    //Variable naming, if it has a small scope of use, then give it a short name
    // If it has a long variable name it should have a larger scope of usage defining the name
    fmt.Println(pklvl, j, i, k)
    //Acronyms should always be all uppercase
    //ie
    var theURL string = "https://www.google.com"
    fmt.Println(theURL)
    //variable type conversion
    ten := 10
    ten = float32(ten)
    fmt.Println(ten)
}
