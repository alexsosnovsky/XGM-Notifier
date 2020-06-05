package main

import (
    "fmt"
    "strconv"
    "time"
    "io/ioutil"
    "net/http"
    gojsonq "github.com/thedevsaddam/gojsonq/v2"
)

const API_LINK = "https://eu.mc-api.net/v3/server/ping/"

type PGMMatch struct {
    mapName string
    duration int
}

type serverMap map[string]PGMMatch

func managePGMServers(){
    PGMServers := make(serverMap)
    // Future: fill map with all PGM servers specified in config
    ocn := PGMMatch{mapName: "Not Pinged Yet", duration: -1}
    PGMServers["play.oc.tc"] = ocn
    for {
        PGMServers.pollPGMServers()
        time.Sleep(30 * time.Second)
    }
}

func (PGMServers serverMap) pollPGMServers() {
    for server, prevMatch := range PGMServers {
        prevMapName := prevMatch.mapName
        prevDuration := prevMatch.duration
        curMapName, curDuration := pollPGMServerByIP(server)
        if prevMapName != curMapName {
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
        newMatch := PGMMatch{mapName: curMapName, duration: curDuration}
        PGMServers[server] = newMatch
    }
}

func pollPGMServerByIP(curIP string) (mapName string, duration int) {
    res, err := http.Get(API_LINK + curIP)
    defer res.Body.Close()
    if err != nil {
    	panic(err)
    }
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
    	panic(err)
    }

    matchFound := false
    lookForMatch := 0
    mapName = "None Found"
    duration = -1

    for !matchFound {
        matchString := "bukkit_extra.pgm." + strconv.Itoa(lookForMatch)
        curPGMData := gojsonq.New().JSONString(string(body)).Find(matchString)
        if curPGMData == nil {
            lookForMatch += 1
            if lookForMatch > 200 {
                // Normally PGM servers don't even get as high as 100 matches before restarting, so this prevents infinite loops in case no match is detected.
                // Convert to pull from config file.
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
