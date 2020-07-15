package main

import "testing"

func TestBot(t *testing.T) {

    // Test for invalid IP handling
    matchFound, _, _ := pollPGMServerByIP("not.a.real.ip")
    if matchFound {
        t.Errorf("pollPGMServerByIP failed, expected matchFound == false, received matchFound == true")
    } else {
        t.Logf("pollPGMServerByIP handled an invalid IP correctly")
    }

    // Test command with no arguments
    command, args := parseCommandArgs("stop")
    if command != "stop" || len(args) != 0 {
        t.Errorf("parseCommandArgs failed, expected command == 'stop' and len(args) == 0")
    } else {
        t.Logf("parseCommandArgs parsed a command with 0 arguments correctly")
    }

    // Test command with 1 argument
    command, args = parseCommandArgs("load map1")
    if command != "load" || len(args) != 1 || args[0] != "map1" {
        t.Errorf("parseCommandArgs failed, expected command == 'stop' and len(args) == 1 and correct arguments in args")
    } else {
        t.Logf("parseCommandArgs parsed a command with 1 argument correctly")
    }

    // Test command with 2 arguments
    command, args = parseCommandArgs("load map1 map2")
    if command != "load" || len(args) != 2 || args[0] != "map1" || args[1] != "map2" {
        t.Errorf("parseCommandArgs failed, expected command == 'stop' and len(args) == 2 and correct arguments in args")
    } else {
        t.Logf("parseCommandArgs parsed a command with 2 arguments correctly")
    }
}
