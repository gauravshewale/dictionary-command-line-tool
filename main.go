package main

import (
	"fmt"
	"log"

	"github.com/gauravshewale/dictionary-command-line-tool/config"
	"github.com/gauravshewale/dictionary-command-line-tool/services"
	"github.com/gauravshewale/dictionary-command-line-tool/utils"
)

func main() {
	err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	var searchString string
	fmt.Println("-------------------------------------------------------")
	fmt.Println("Welcome to Easy Dictionary Tool")
	fmt.Println("-------------------------------------------------------")
	fmt.Print("Please enter the word you are looking for: ")
	fmt.Scanln(&searchString)
	results, err := services.SearchDictionary(searchString)
	if err == nil {
		utils.PrintResult(&results)
	}
}
