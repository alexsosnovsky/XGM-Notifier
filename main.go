package main

import (
    "fmt"
)

const DEBUG = true

func main() {
    go managePGMServers()

    stopping := make(chan bool)
    go manageUserConsole(stopping)
    <- stopping
    fmt.Println("Bot Stopping")
}

func printConsoleInput(){
    fmt.Printf(">>> ")
}
