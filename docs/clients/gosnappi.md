# gosnappi

![snappi](https://github.com/open-traffic-generator/snappi/raw/main/snappi-logo.png)

## Overview

Test scripts written in `gosnappi`, an auto-generated Go module, can be executed against any traffic generator conforming to the Open Traffic Generator API.

## Install on a client

```Shell
mkdir tests && cd tests
go mod init tests
go get github.com/open-traffic-generator/snappi/gosnappi
```

## Start scripting

Add a new file `hello.go` with following snippet:

```Go
package main

import (
    "log"

    "github.com/open-traffic-generator/snappi/gosnappi"
)

func main() {
    // Create a new API handle to make API calls against a traffic generator
    api := gosnappi.NewApi()

    // Set the transport protocol to either HTTP or GRPC
    api.NewHttpTransport().SetLocation("https://localhost")

    // Create a new traffic configuration that will be set on traffic generator
    config := api.NewConfig()

    // Add port locations to the configuration
    p1 := config.Ports().Add().SetName("p1").SetLocation("localhost:5555")

    // Configure the flow and set the endpoints
    flow := config.Flows().Add().SetName("f1")
    flow.TxRx().Port().SetTxName(p1.Name())

    // Configure the size of a packet and the number of packets to transmit
    flow.Size().SetFixed(128)
    flow.Duration().FixedPackets().SetPackets(1000)

    // Configure the header stack
    pkt = flow.Packet()
    eth := pkt.Add().Ethernet()
    ipv4 := pkt.Add().Ipv4()
    tcp := pkt.Add().Tcp()

    eth.Dst().SetValue("00:11:22:33:44:55")
    eth.Src().SetValue("00:11:22:33:44:66")

    ipv4.Src().SetValue("10.1.1.1")
    ipv4.Dst().SetValue("20.1.1.1")

    tcp.SrcPort().SetValue(5000)
    tcp.DstPort().SetValue(6000)

    // Push traffic configuration constructed so far to traffic generator
    log.Println(config.ToYaml())
    if err := api.SetConfig(config); err != nil {
        log.Fatal(err)
    }

    // Start transmitting the configured flows
    ts := api.NewTransmitState()
    ts.SetState(gosnappi.TransmitStateState.START)
    if err := api.SetTransmitState(ts); err != nil {
        log.Fatal(err)
    }

    // Fetch and the port metrics
    req := api.NewMetricsRequest()
    req.Port().SetPortNames([]string{p1.Name()})
    metrics, err := api.GetMetrics(req)
    if err != nil {
        log.Fatal(err)
    }

    log.Println(metrics.ToYaml())
}
```

## Run test

```Shell
go run hello.go
```