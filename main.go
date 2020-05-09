package main

import (
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

func main() {
	var inputFile = flag.String("inputFile", "input.txt", "Input file")
	configFile, err := ioutil.ReadFile(*inputFile)

	if err != nil {
		panic(err)
	}

	lawn, mowersAndCommands, err := ParseConfigFile(configFile)
	if err != nil {
		log.Errorf("Error while parsing : %v", err)
		return
	}

	for _, mowerAndCommands := range *mowersAndCommands {
		mower := mowerAndCommands.mower
		mowerMoved := mower.move(mowerAndCommands.commands, *lawn)

		fmt.Println(mowerMoved.positionAsString())
	}

}
