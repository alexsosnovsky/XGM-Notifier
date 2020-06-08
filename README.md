# XGM Notifier
##### by Alex Sosnovsky
This project is a work-in-progress, and all features are subject to change in the future.

## Overview
XGM Notifier is an automated Discord bot, which collects information about [PGM](https://github.com/pgmdev/pgm) servers, and sends users notifications when their favorite games are playing.

## Technical
* The bot is written entirely in Go, and stores information about games and notifications in a local SQLite3 database.
* Game information is obtained by pinging PGM servers for server status data, available in JSON format.

## Dependencies
* go-sqlite3: https://github.com/mattn/go-sqlite3
    * Interfaces between the Go code and your SQLite3 database.
* DiscordGo: https://github.com/bwmarrin/discordgo
    * Interfaces between the Go code and the Discord API.
* GoJSONQ: https://github.com/thedevsaddam/gojsonq
    * Used to parse PGM server data JSON.

## Installation
* Clone (or download) the repository. 
* Install the dependencies according to their instructions (linked above).
* In the main directory, create a `bot_token.txt` file that contains your bot's token, found in the [Discord Developer Portal](https://discord.com/developers/applications/).
* If you are interested in receiving extra information about what is happening behind the scenes, set the `DEBUG` value in [main.go](https://github.com/alexsosnovsky/XGM-Notifier/blob/b43be90200f44726b82de6ce6f0598bac48a95d7/main.go#L9) to `true` prior to compiling.
* Run `go build` to compile the bot. An executable titled `XGMNotifier` should be created in your main directory.

## Usage
* Run `./XGMNotifier` to start the bot. A file named `xgm.db` will be created the first time you run it, and contains the SQLite3 database used to store maps and user preferences.
* While the bot is running, the terminal functions as an administrative console; a full set of commands can be found below.
* The bot repeatedly pings the specified PGM server, and whenever a match cycle is detected, any users subscribed to the new game will be messaged by the bot.
    * If a game that is not present in the database is detected, it is automatically added to the database and subsequently selectable by users.
* Users can message the bot to add or remove maps from their notification lists.
* To stop the bot, input `stop` into the administrative console.

## User Commands 
##### (Messaged to the Bot)
For any commands that accept map names, separate the maps with spaces, and replace all spaces in the maps' names with underscores.
* `add <map1> <map2> etc.` The bot adds the specified maps to the user's notification list, if they are present in the map list.
* `remove <map1> <map2> etc.` The bot deletes the specified maps from the user's notification list.
* `removeall`: The bot deletes all maps from the user's notification list.

## Administrative Commands
##### (Terminal Input)
For any commands that accept map names, separate the maps with spaces, and replace all spaces in the maps' names with underscores.
* `stop` (alias `exit`): The bot closes its Discord session and stops running.
* `load <map1> <map2> etc.`: The maps specified in the arguments are loaded, and subsequently selectable by users.

