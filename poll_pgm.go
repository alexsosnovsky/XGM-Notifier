package main

import (
    "fmt"
    "strconv"
    "time"
    "io/ioutil"
    "net/http"
    _ "github.com/mattn/go-sqlite3"
    gojsonq "github.com/thedevsaddam/gojsonq/v2"
)

const API_LINK = "https://eu.mc-api.net/v3/server/ping/"

type pgmMatch struct {
    mapName string
    duration int
}

type serverMap map[string]pgmMatch

func managePGMServers(){
    pgmServers := make(serverMap)
    // Future: fill map with all pgm servers specified in config
    curServer := pgmMatch{mapName: "Not Pinged Yet", duration: -1}
    pgmServers["play.oc.tc"] = curServer
    for {
        pgmServers.pollPGMServers()
        // Future: make ping rate configurable
        time.Sleep(30 * time.Second)
    }
}

func (pgmServers serverMap) pollPGMServers() {
    for server, prevMatch := range pgmServers {
        prevMapName := prevMatch.mapName
        prevDuration := prevMatch.duration
        matchFound, curMapName, curDuration := pollPGMServerByIP(server)
        if matchFound {
            if prevMapName != curMapName {
                sendMessages(curMapName)
                if !existsMap(curMapName) {
                    addMap(curMapName)
                    if DEBUG {
                        fmt.Printf("> New map %s added to the database.", curMapName)
                    }
                }
                if DEBUG {
                    if prevMapName != "Not Pinged Yet" {
                        fmt.Printf("\n> IP %s has cycled, last ping for %s was %d.\n",server,prevMapName,prevDuration)
                    } else {
                        fmt.Printf("\n> Pinged %s for the first time.\n",server)
                    }
                }
            }
            if DEBUG {
                fmt.Printf("\n> IP %s has been playing %s for %d seconds.\n",server,curMapName,curDuration)
                printConsoleInput()
            }
            newMatch := pgmMatch{mapName: curMapName, duration: curDuration}
            pgmServers[server] = newMatch
        } else {
            fmt.Printf("Connection failed to %s; the server is offline or the IP is invalid.\n", server)
            printConsoleInput()
        }
    }
}

func pollPGMServerByIP(curIP string) (matchFound bool, mapName string, duration int) {
    res, err := http.Get(API_LINK + curIP)
    defer res.Body.Close()
    if err != nil {
        panic(err)
    }
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        panic(err)
    }

    matchFound = false
    lookForMatch := 0
    mapName = "None Found"
    duration = -1

    for !matchFound {
        matchString := "bukkit_extra.pgm." + strconv.Itoa(lookForMatch)
        curpgmData := gojsonq.New().JSONString(string(body)).Find(matchString)
        if curpgmData == nil {
            lookForMatch += 1
            if lookForMatch > 200 {
                // Normally pgm servers don't even get as high as 100 matches before restarting, so this prevents infinite loops in case no match is detected.
                // Future: make the threshold configurable.
                break
            }
        } else {
            matchFound = true

            mapNameData, err := gojsonq.New().JSONString(string(body)).FindR(matchString+".map.name")
            if err != nil {
                panic(err)
            }
            curMapName, err := mapNameData.String()
            if err != nil {
                panic(err)
            }
            mapName = curMapName

            durationData, err := gojsonq.New().JSONString(string(body)).FindR(matchString+".duration")
            if err != nil {
                panic(err)
            }
            curDuration, err := durationData.Float64()
            if err != nil {
                panic(err)
            }
            duration = int(curDuration)
        }
    }

    // Returns mapName, duration
    return
}
