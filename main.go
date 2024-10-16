package main

import (
	"github.com/showwin/speedtest-go/speedtest"
  "fmt"
)

func main()  {
  var speedTestClient = speedtest.New()
  var serverID int 

  serverList, _ := speedTestClient.FetchServers()
  for _, s := range serverList {
    fmt.Printf("%s\n", s)
  }

  fmt.Printf("Enter id of the server you want to connect to (or leave 0 for auto): ")
  fmt.Scan(&serverID)

  if serverID == 0 {
    // Store it as a variable for future reference
    var autoServerID string = serverList[0].ID[:5]
    serverFetch, _ := speedTestClient.FetchServerByID(autoServerID)

    fmt.Printf("Connecting to %s", serverFetch)
  }

  targets, _ := serverList.FindServer([]int{serverID})

  fmt.Println("Wait while we do the magic...")
  for _, s := range targets {
    s.PingTest(nil)
		s.DownloadTest()
		s.UploadTest()
		fmt.Printf("Latency: %s, Download Speed: %s, Upload Speed: %s\n", s.Latency, s.DLSpeed, s.ULSpeed)
		s.Context.Reset() 
  }
}

