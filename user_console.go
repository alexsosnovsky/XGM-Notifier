package main

import (
    "fmt"
)

func manageUserConsole(stopping chan bool) {
    var userCommand string
    for {
        fmt.Scanln(&userCommand)
        parseCommand(userCommand, stopping)
        printConsoleInput()
    }
}

func parseCommand(userCommand string, stopping chan bool) {
    switch userCommand {
        case "exit":
            stopBot(stopping)
        case "stop":
            stopBot(stopping)
        case "load":
            fmt.Println("\nNot implemented yet")
        default:
            fmt.Printf("\n%s is not a recognized command.\n",userCommand)
    }
}

func stopBot(stopping chan bool) {
    stopping <- true
}
