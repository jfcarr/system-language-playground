package main

import (
	"encoding/json"
	"fmt"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
)

// EmailConfigs .. List of email configurations
type EmailConfigs struct {
	EmailConfigs []EmailConfig `json:"EmailConfigs"`
}

// EmailConfig .. Configuration for IMAP connection and account
type EmailConfig struct {
	ServerName   string `json:"ServerName"`
	ServerPort   int    `json:"ServerPort"`
	UserName     string `json:"UserName"`
	UserPassword string `json:"UserPassword"`
	ReportAll    bool   `json:"ReportAll"`
}

func checkEmail(serverName string, serverPort int, userName string, userPassword string, printZeroCount bool) int {
	// Connect to the IMAP server
	c, err := client.DialTLS(fmt.Sprintf("%v:%v", serverName, serverPort), nil)
	checkError(err)

	// Logout when we're done
	defer c.Logout()

	// Login
	if err := c.Login(userName, userPassword); err != nil {
		checkError(err)
	}

	// Select INBOX
	_, err = c.Select("INBOX", false)
	checkError(err)

	// Set search criteria and search for unread messages
	criteria := imap.NewSearchCriteria()
	criteria.WithoutFlags = []string{imap.SeenFlag}
	ids, err := c.Search(criteria)
	checkError(err)

	if len(ids) > 0 {
		seqset := new(imap.SeqSet)
		seqset.AddNum(ids...)

		messages := make(chan *imap.Message, 10)
		done := make(chan error, 1)
		go func() {
			done <- c.Fetch(seqset, []imap.FetchItem{imap.FetchEnvelope}, messages)
		}()

		// Count them and display the results
		messageCount := 0
		for range messages {
			messageCount = messageCount + 1
		}
		if messageCount > 0 || printZeroCount == true {
			fmt.Println(messageCount, ":", userName)
		}

		if err := <-done; err != nil {
			checkError(err)
		}
	} else {
		if printZeroCount == true {
			fmt.Println("0 :", userName)
		}
	}

	return 0
}

func readConfig() {
	// Get the working directory
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	// Full path to the config file
	configFile := path.Join(exPath, "email-util.json")

	jsonFile, err := os.Open(configFile)
	checkError(err)

	defer jsonFile.Close()

	// Read all the contents of the config file
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var emailConfigs EmailConfigs

	// Unmarshal the json-formatted contents of the file into an array of EmailConfig structs
	json.Unmarshal(byteValue, &emailConfigs)

	// Walk through the list of structs and use the info to call the email checker
	for i := 0; i < len(emailConfigs.EmailConfigs); i++ {
		checkEmail(
			emailConfigs.EmailConfigs[i].ServerName,
			emailConfigs.EmailConfigs[i].ServerPort,
			emailConfigs.EmailConfigs[i].UserName,
			emailConfigs.EmailConfigs[i].UserPassword,
			emailConfigs.EmailConfigs[i].ReportAll)
	}
}

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	readConfig()
}
