package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)

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
	configFile := path.Join(exPath, "email-util.config")

	if _, err := os.Stat(configFile); err == nil {
		// Open the file
		file, err := os.Open(configFile)
		checkError(err)
		defer file.Close()

		scanner := bufio.NewScanner(file)

		// Loop through the file entries
		for scanner.Scan() {
			s := strings.Split(scanner.Text(), "|")

			serverName, serverPort, userName, userPassword, printZeroCount := s[0], s[1], s[2], s[3], s[4]

			serverPortI, err := strconv.Atoi(serverPort)
			checkError(err)

			printZeroCountB, err := strconv.ParseBool(printZeroCount)
			checkError(err)

			checkEmail(serverName, serverPortI, userName, userPassword, printZeroCountB)
		}

		if err := scanner.Err(); err != nil {
			checkError(err)
		}
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
