package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/showwin/speedtest-go/speedtest"
)

func ensureConnected() (connected bool) {
	serverList := [3]string{
		"http://clients3.google.com/generate_204",
		"https://google.com",
		"https://example.com",
	}

	errList := []error{}

	for index := range serverList {
		server := serverList[index]

		_, err := http.Get(server)
		if err != nil {
			fmt.Printf("Could not connect to %s\n", server)
			errList = append(errList, err)
		}
	}

	if len(errList) == len(serverList) {
		return false
	}

	return true
}

func main() {
	var (
		speedTestClient = speedtest.New()
		serverID        int
	)

	if !ensureConnected() {
		fmt.Println(
			"You are not connected to network. Try connecting to it first, before checking your speed again.",
		)

		os.Exit(1)
	}

	serverList, err := speedTestClient.FetchServers()
	if err != nil {
		fmt.Printf("An error ocurred while fetching servers: %s", err)
		os.Exit(1)
	}

	for _, s := range serverList {
		fmt.Printf("%s\n", s)
	}

	fmt.Printf(
		"Enter id of the server you want to connect to (or leave 0 for auto): ",
	)
	fmt.Scan(&serverID)

	if serverID == 0 {
		// Store it as a variable for future reference
		var autoServerID string = serverList[0].ID[:5]
		serverFetch, err := speedTestClient.FetchServerByID(autoServerID)
		if err != nil {
			fmt.Printf("Error while fetching the server details: %s", err)
			os.Exit(1)
		}

		fmt.Printf("Connecting to %s\n", serverFetch)
	}

	targets, err := serverList.FindServer([]int{serverID})
	if err != nil {
		fmt.Printf("Could not find the server: %s", err)
		os.Exit(1)
	}

	fmt.Println("Wait while we do the magic...")
	for _, s := range targets {
		s.PingTest(nil)
		s.DownloadTest()
		s.UploadTest()
		fmt.Printf(
			"Latency: %s, Download Speed: %s, Upload Speed: %s\n",
			s.Latency,
			s.DLSpeed,
			s.ULSpeed,
		)

		s.Context.Reset()
	}
}
