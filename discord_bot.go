package main

import (
    "fmt"
    "strings"
    "io/ioutil"
    "github.com/bwmarrin/discordgo"
)

func createBot() (*discordgo.Session){
    token, err := ioutil.ReadFile("bot_token.txt")
    if err != nil {
        panic(err)
    }
    discord, err := discordgo.New("Bot " + string(token))
    if err != nil {
        panic(err)
    }
    discord.AddHandler(messageCreate)
    err = discord.Open()
    if err != nil {
        panic(err)
    }
    fmt.Printf("> Discord session opened.\n")
    return discord
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

    if DEBUG {
        fmt.Printf("> Message received.\n",m.Content)
    }

    // Ignore messages sent by this bot.
    if m.Author.ID == s.State.User.ID {
        return
    }

    // Ignore non-private messages.
    channel, err := s.State.Channel(m.ChannelID)
    if err != nil {
        channel, err = s.Channel(m.ChannelID)
        if err != nil {
            panic(err)
        }
    }
    if channel.Type != discordgo.ChannelTypeDM {
        return
    }


    discordUser := m.Author.Username + "#" + m.Author.Discriminator
    channelID := m.ChannelID

    response := "Message received."
    command, args := parseCommandArgs(m.Content)
    switch command {
        case "add":
            if args == nil {
                response = "Not enough arguments. Specify at least one map to add."
                break
            }
            response = ""
            for index := 0; index < len(args); index++ {
                mapName := strings.Replace(args[index], "_", " ", -1)
                if existsMap(mapName) {
                    addUserToMap(mapName, discordUser, channelID)
                    response += "Added " + mapName + " to your notification list.\n"
                } else {
                    response += mapName + " is not a selectable map.\n"
                }
            }
        case "remove":
            if args == nil {
                response = "Not enough arguments. Specify at least one map to remove."
                break
            }
            for index := 0; index < len(args); index++ {
                mapName := strings.Replace(args[index], "_", " ", -1)
                removeUserFromMap(mapName, channelID)
            }
            response = "Removed all the requested maps from your notification list."
        case "removeall":
            response = "Removed all maps from your notification list."
            purgeUser(channelID)
        default:
            response = command + " is not a recognized command."
    }

    s.ChannelMessageSend(m.ChannelID, response)
}

func sendMessages(mapName string){
    channels := findUsersByMap(mapName)
    response := mapName + " is now playing on `play.oc.tc`."
    for index := 0; index < len(channels); index++ {
        discordSession.ChannelMessageSend(channels[index], response)
    }
}
