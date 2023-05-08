package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoring = 3
const delay = 5

func main() {
	displayMenu()

	for {

		displayIntroduction()
		userChoice := readChoice()
		switch userChoice {
		case 1:
			fmt.Println("Monitoring...")
			startMonitoring()
		case 2:
			fmt.Println("Displaying Logs...")
			displayLogs()
		case 0:
			fmt.Println("Exiting the program")
			os.Exit(0)
		default:
			fmt.Println("Unknown command")
			os.Exit(-1)
		}
	}
}

func displayIntroduction() {
	fmt.Println("1- Start Monitoring")
	fmt.Println("2- Display Logs")
	fmt.Println("0- Exit Program")

}

func displayMenu() {
	name := "Luiz"
	version := 1.3
	fmt.Println("Hello,", name, "!")
	fmt.Println("Software Version:", version)
}

func readChoice() int {
	var choice int
	fmt.Printf("select your operation: ")
	fmt.Scan(&choice)
	return choice
}

func startMonitoring() {
	fmt.Println("Getting Requests: ")

	sites := readSitesFromFile()

	for i := 0; i < monitoring; i++ {
		for i, site := range sites {
			fmt.Println("Testing site", i, ":", site)
			testSite(site)
		}
		time.Sleep(delay * time.Second)
		newLine()
	}

}

func displayLogs() {

	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))

}

func testSite(site string) {
	status, _ := http.Get(site)

	if status.StatusCode == 200 {
		fmt.Println("Site:", site, "loaded successfully!")
		writeLog(site, true)
	} else {
		fmt.Println("Site:", site, "is experiencing connectivity issues. Status Code:", status.StatusCode)
		writeLog(site, false)
	}

}

func newLine() {
	fmt.Println("")
	fmt.Println("")
}

func readSitesFromFile() []string {
	var sites []string

	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("An error has occurred:", err)
	}
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}

func writeLog(site string, status bool) {

	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + "-" + site + "- online:" + strconv.FormatBool(status) + "\n")
	file.Close()
}
