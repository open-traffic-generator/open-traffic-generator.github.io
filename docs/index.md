# Home

## Overview

Open Traffic Generator (OTG) APIs and Data Models are APIs to test Layer 2 & 3 routers and switches. They include support for testing the forwarding plane as well as testing the control plane.  They are:

* Open
* Vendor Neutral
* Intent Based
* Declarative

The specification is available on GitHub [here](https://github.com/open-traffic-generator/models/blob/master/artifacts/openapi.yaml) (viewer friendly version at reDocly [here](https://redocly.github.io/redoc/?url=https://raw.githubusercontent.com/open-traffic-generator/models/master/artifacts/openapi.yaml)).  Specifications support both REST or gRPC interfaces.

## Features

OTG is an actively developing specification, with contributions coming from real [use cases](/examples/#use-cases). On high level, the model allows to express the following building blocks of a traffic generator configuration:
 
* *Test Ports* with:
	- LLDP
	- LAG
	- LACP
* *Traffic Flows* associated with:
	- Test Ports
	- Emulated Devices
* *Emulated Devices* with:
	- BGP
	- IS-IS


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

