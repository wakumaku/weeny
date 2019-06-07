package main

import (
	"bufio"
	"fmt"
	"os"
	"weeny/application"
	"weeny/cache"
	"weeny/hasher"
)

func main() {

	app := application.New(cache.NewInMemory(), &hasher.Md5{})

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter an URL to generate a hash or a hash to get an URL: ")
		text, _ := reader.ReadString('\n')

		if len(text) >= 2 {
			text = text[:len(text)-1]
		} else {
			fmt.Print("Entered string is too short, try again.")
			continue
		}

		if len(text) > 4 && text[0:4] == "http" {
			hash, err := app.Save(text)
			if err != nil {
				fmt.Printf("Error saving url: %s\n", err)
			} else {
				fmt.Printf("Generated hash: %s\n", hash)
			}
		} else {
			url, err := app.Get(text)
			if err != nil {
				fmt.Printf("Error getting url from hash: %s: %s\n", err, text)
			} else {
				fmt.Printf("URL: %s\n", url)
			}
		}

	}
}
