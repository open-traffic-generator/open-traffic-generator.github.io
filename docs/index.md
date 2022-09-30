# Home

## Overview

Open Traffic Generator (OTG) APIs and data models are northbound API specifications for modern Traffic Generators and Protocol Emulators designed to test Layer 2-7 network devices. They include support for testing  forwarding as well as control planes. OTG APIs are:

* Open
* Vendor-neutral
* Intent-based
* Declarative

The formal [model specification](https://github.com/open-traffic-generator/models/blob/master/artifacts/openapi.yaml) can be found on GitHub under [Open Traffic Generator](https://github.com/open-traffic-generator) organization. To explore the model, a viewer friendly [**ReDoc rendering**](https://redocly.github.io/redoc/?url=https://raw.githubusercontent.com/open-traffic-generator/models/master/artifacts/openapi.yaml) is available as well. The OTG APIs support both REST and gRPC interfaces.

## Features

OTG is an actively developing specification, with contributions from real [use cases](/examples/#use-cases). The model allows to express the following building blocks of a traffic generator configuration:
<!-- TODO add links from bold items to paragraphs in Model section -->
* **Test Ports** with Layer 1&2 capabilities, including:
	- LLDP, LAG, LACP
* Emulated **OTG Devices** with Layer 2&3 features:
	- IPv4, IPv6 interfaces
	- ARP, IPv6 ND
	- BGP, IS-IS routing protocols
* **Traffic Flows** 
    - associated with either Test Ports, or OTG Devices
	- expressing L2-4 properties like Ethernet, IPv4/IPv6, TCP/UDP
	- stateless or stateful capabilities for transport protocols
	- with implementation-specific application payload
* Run-time **Metrics** and traffic **Capture** capabilities

![Example OTG Diagram](images/otg-example-diagram.svg)
<sub>Fig. 1. Sample diagram of the OTG configuration with emulated BGP routers, traffic flows and a Device Under Test</sub>


## Implementations

Following tools provide OTG compliant implementations.
 
* *[Ixia-c Community Edition](https://ixia-c.dev)*: container-based traffic generator from Keysight. Limited to 4 ports and Traffic Flows
* *[Keysigh Elastic Network Generator](https://www.keysight.com/us/en/products/network-test/protocol-load-test/keysight-elastic-network-generator.html)*: commercial offering of OTG implementation for a family of Keysigh/Ixia Traffic Generators
* *[IxNetwork](https://www.keysight.com/us/en/products/network-test/protocol-load-test/ixnetwork.html)*: [snappi-ixnetwork](https://github.com/open-traffic-generator/snappi-ixnetwork) supports running OTG/snappi scripts against Keysight IxNetwork
* *[Magna](https://github.com/openconfig/magna)*: native open-source OTG implementation from [OpenConfig project](https://openconfig.net/)
* *[TRex](https://trex-tgn.cisco.com/)*: [snappi-trex](https://github.com/open-traffic-generator/snappi-trex) supports running OTG/snappi scripts against TRex. Supports layer 2-3 Traffic Flows

## Clients

There are multiple ways to communicate with OTG-based Traffic Generator:
 
* *OTGen* command-line tool
* *Test Program* written in Python or Go using *snappi* library
* *Raw API calls* with *curl* or similar utilities
* *Custom* clients anyone can develop

[snappi](https://pypi.org/project/snappi/) and [gosnappi](https://pkg.go.dev/github.com/open-traffic-generator/snappi/gosnappi) provide custom built client side API libraries for the OTG specifications for Python and Golang respectively.  For other languages, SDKs can be built using [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator) (for REST API) or [protobuf tools](https://github.com/protocolbuffers/protobuf) (for gRPC).  

## Use Cases
 
Real use case are the basic of OTG development. Notable open-source projects leveraging OTG:
 
* [OpenConfig Feature Profiles](https://github.com/openconfig/featureprofiles)
* [SONiC Testbed](https://github.com/sonic-net/sonic-mgmt)
* [SONiC-DASH CI Pipeline](https://github.com/Azure/DASH)

