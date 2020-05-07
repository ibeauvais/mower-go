package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	inputFile := os.Args[1]

	configFile, err := ioutil.ReadFile(inputFile)

	if err != nil {
		panic(err)
	}

	lawn, mowersAndCommands, err := ParseConfigFile(configFile)
	if err != nil {
		panic(err)
	}

	for _, mowerAndCommands := range *mowersAndCommands {
		mower := mowerAndCommands.mower
		mowerMoved := mower.move(mowerAndCommands.commands, *lawn)

		fmt.Println(mowerMoved.positionAsString())
	}

}
