package main

import (
    "fmt"
    "database/sql"
    "github.com/bwmarrin/discordgo"
)

const DEBUG = false
var mapDB *sql.DB
var discordSession *discordgo.Session

func main() {
    mapDB = openDB()
    discordSession = createBot()

    go managePGMServers()

    stopping := make(chan bool)
    go manageUserConsole(stopping)

    <- stopping
    discordSession.Close()
    mapDB.Close()
    fmt.Println("Bot Stopping")
}
