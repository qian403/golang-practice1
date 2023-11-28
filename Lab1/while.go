package main

import (
    "fmt"
    "os"
    "os/exec"
    "runtime"

)

func main() {



    fmt.Print("輸入一個數字：")
    var x int
    fmt.Scanln(&x)
    fmt.Println(x)
    if runtime.GOOS  == "windows" {

        cmd := exec.Command("cmd", "/c", "pause")
        cmd.Stdin = os.Stdin
        cmd.Stdout = os.Stdout
        cmd.Run()
    }


    fmt.Println("owo")
}