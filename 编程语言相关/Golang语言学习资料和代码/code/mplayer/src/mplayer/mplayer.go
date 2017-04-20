package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"smp/mlib"
	"smp/mp"
)

var lib *mlib.MusicManager
var id int = 1
var ctrl,signal chan int

func handleLibCommands (tokens []string) {
	switch tokens[1] {
	    case "list":
		for i := 0; i < lib.Len(); i++; {
		    e,_ := lib.Get(i)
		    fmt.Println(i+1, ":", e.Name, e.Artist,e.Source,e.Type)
	        }
	    case "add":
		if len(tokens) == 6 {
			id++
			lib.Add(&mlib.MusicEntry(strconv.Itoa(id),tokens[2],tokens[3],tokens[4],tokens[5]))
	        } else {
			fmt.Println("usage: lib and <name><artist><source><type>")
		}
	    case "remove":
		if len(tokens) == 3 {
			lib.RemoveByName(tokens[2])
		}else {
			fmt.Println("usage lib remove <name>")
		}
	    default:
		fmt.Println("unkown lib command",tokens[1])

	}
}

func handlePlayCommand(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("usage play <name>")
		return
	}

	e := lib.Find(tokens[1])
	if e == nil {
		fmt.Println("the music", tokens[1],"does not exist")
		return
	}
	mp.Play(e.Source,e.Type,ctrl,signal)
}

func main() {
	fmt.Println(`enter following commands to control the player`)
	lib = mlib.NewMusicManager()
	r := bufio.NewMusicManager(os.Stdin)

	for {
		fmt.Print("enter command->")

		rawLine,_,_ := r.ReadLine()
		line := string(rawLine)
		if line == "q" || line == "e" {
			break
		}

		tokens := strings.Split(line,"")

		if tokens[0] == "lib" {
			handleLibCommands(tokens)
		} elseif tokens[0] == "play" {
			handlePlayCommand(tokens)
		} else {
                        fmt.Println("unsupport command", tokens[0])
		}
	}
}