package main

import (
    "fmt"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

func openDB() (*sql.DB){
    database, err := sql.Open("sqlite3", "./xgm.db")
    if(err != nil) {
        panic(err)
    }

    fmt.Println("> Database opened.")

    statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS maps (mapname TEXT)")
    if(err != nil) {
        panic(err)
    }
    statement.Exec()
    statement.Close()

    statement2, err := database.Prepare("CREATE TABLE IF NOT EXISTS entries (mapname TEXT, discorduser TEXT, channelid TEXT)")
    if(err != nil) {
        panic(err)
    }
    statement2.Exec()
    statement2.Close()

    return database
}

func existsMap(mapName string) (bool){
    rows, err := mapDB.Query("SELECT * from maps WHERE mapname = (?)", mapName)
    if(err != nil) {
        panic(err)
    }
    exists := rows.Next()

    return exists
}

func addMap(mapName string) {
    statement, err := mapDB.Prepare("INSERT INTO maps (mapname) VALUES (?)")
    if(err != nil) {
        panic(err)
    }
    statement.Exec(mapName)
    statement.Close()
}

func addUserToMap(mapName string, discordUser string, channelID string) {
    statement, err := mapDB.Prepare("INSERT INTO entries (mapname, discorduser, channelid) VALUES (?, ?, ?)")
    if(err != nil) {
        panic(err)
    }
    statement.Exec(mapName,discordUser,channelID)
    statement.Close()
}

func removeUserFromMap(mapName string, channelID string) {
    statement, err := mapDB.Prepare("DELETE FROM entries WHERE mapname = (?) AND channelid = (?)")
    if(err != nil) {
        panic(err)
    }
    statement.Exec(mapName,channelID)
    statement.Close()
}

func purgeUser(channelID string){
    statement, err := mapDB.Prepare("DELETE FROM entries WHERE channelid = (?)")
    if(err != nil) {
        panic(err)
    }
    statement.Exec(channelID)
    statement.Close()
}

func findUsersByMap(mapName string) ([]string){
    rows, err := mapDB.Query("SELECT * from entries WHERE mapname = (?)", mapName)
    if(err != nil) {
        panic(err)
    }
    channels := make([]string, 0)
    for rows.Next() {
        var curMap string
        var discordUser string
        var channelID string
        rows.Scan(&curMap, &discordUser, &channelID)
        channels = append(channels,channelID)
    }
    rows.Close()

    return channels
}
