package main

import (
  "github.com/prometheus-community/pro-bing"
  "fmt"
  "os"
  "os/signal"
)

func OnRecieve(pkt *probing.Packet) {
  fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v\n",
    pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt)
}

func OnDuplicateRecieve(pkt *probing.Packet) {
  fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v ttl=%v (DUP!)\n",
    pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt, pkt.TTL)
}

func OnFinish(stats *probing.Statistics) {
  fmt.Printf("\n--- %s ping statistics ---\n", stats.Addr)
  fmt.Printf("%d packets transmitted, %d packets received, %v%% packet loss\n",
    stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
  fmt.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
    stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
}

func main()  {
  var server string
  fmt.Printf("Enter host name: ")
  fmt.Scan(&server)

  pinger, err := probing.NewPinger(server)
  c := make(chan os.Signal, 1)
  signal.Notify(c, os.Interrupt)

  go func() {
    for _ = range c {
      pinger.Stop()
    }
  }()

  if err != nil {
    panic(err)
  }

  pinger.OnRecv = func(pkt *probing.Packet) {
    OnRecieve(pkt)
  }

  pinger.OnDuplicateRecv = func(pkt *probing.Packet) {
    OnDuplicateRecieve(pkt)
  }

  pinger.OnFinish = func(stats *probing.Statistics) {
    OnFinish(stats)
  }

  fmt.Printf("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())
  err = pinger.Run()
  if err != nil {
    panic(err)
  }
}

