# Introduction

Open Traffic Generator (OTG) APIs and Data Models are APIs to test Layer 2 & 3 routers and switches. They include support for testing the forwarding plane as well as testing the control plane.  They are:

* Open
* Vendor Neutral
* Intent Based
* Declarative

The specification is available on GitHub [here](https://github.com/open-traffic-generator/models/blob/master/artifacts/openapi.yaml) (viewer friendly version at reDocly [here](https://redocly.github.io/redoc/?url=https://raw.githubusercontent.com/open-traffic-generator/models/master/artifacts/openapi.yaml)).  Specifications support both REST or gRPC interfaces.

[snappi](https://pypi.org/project/snappi/) and [gosnappi](https://pkg.go.dev/github.com/open-traffic-generator/snappi/gosnappi) provide custom built client side API libraries for the OTG specifications for Python and Golang respectively.  For other languages, SDKs can be built using [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator) (for REST API) or [protobuf tools](https://github.com/protocolbuffers/protobuf) (for gRPC).  

Following tools provide OTG compliant implementations.

* [Ixia-c](https://github.com/open-traffic-generator/ixia-c) : Free container-based traffic generator from Keysight.  Limited to 4 ports, supports REST APIs.  Commercial version includes control plane and gRPC support.
* [IxNetwork](https://www.keysight.com/us/en/products/network-test/protocol-load-test/ixnetwork.html) : [snappi-ixnetwork](https://github.com/open-traffic-generator/snappi-ixnetwork) supports running OTG/snappi scripts against IxNetwork.  
* TRex : [snappi-trex](https://github.com/open-traffic-generator/snappi-trex) supports running OTG/snappi scripts against TRex.  Support is present for layer 2-3 traffic generation.
