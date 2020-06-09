package main

import (
    "fmt"
    "strings"
    "bufio"
    "os"
)

func manageUserConsole(stopping chan bool) {
    scanner := bufio.NewScanner(os.Stdin)
    printConsoleInput()
    for scanner.Scan() {
        userCommand := scanner.Text()
        parseCommand(userCommand, stopping)
        printConsoleInput()
    }
}

func parseCommand(userCommand string, stopping chan bool) {
    command, args := parseCommandArgs(userCommand)
    switch command {
        case "exit":
            stopBot(stopping)
        case "stop":
            stopBot(stopping)
        case "load":
            if args == nil {
                fmt.Printf("\n> Not enough arguments. Specify at least one map to load.\n")
            } else {
                for index := 0; index < len(args); index++ {
                    mapName := strings.Replace(args[index], "_", " ", -1)
                    addMap(mapName)
                    fmt.Printf("\n> Manually added %s to the list of selectable maps.\n",mapName)
                }
            }
        default:
            fmt.Printf("\n%s is not a recognized command.\n",command)
    }
}

func parseCommandArgs(userCommand string) (string, []string){
    spaceSep := strings.Split(strings.Trim(userCommand," "), " ")
    command := strings.Trim(spaceSep[0]," ")
    if len(spaceSep) <= 1 {
        if DEBUG {
            fmt.Printf("\n> Command received with no arguments: %s\n",command)
        }
        return command, nil
    } else {
        args := make([]string, 0)
        for index := 1; index < len(spaceSep); index++ {
            args = append(args,strings.Trim(spaceSep[index], " "))
        }
        if DEBUG {
            fmt.Printf("\n> Command received with %d arguments: %s\n",len(args),command)
        }
        return command, args
    }
}

func printConsoleInput(){
    fmt.Printf(">>> ")
}

func stopBot(stopping chan bool) {
    stopping <- true
}
