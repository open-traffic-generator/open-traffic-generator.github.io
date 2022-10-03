# Model

## Formal Model
The formal [model specification](https://github.com/open-traffic-generator/models/blob/master/artifacts/openapi.yaml) can be found on GitHub under [Open Traffic Generator](https://github.com/open-traffic-generator) organization. To explore the model, a viewer friendly [**ReDoc rendering**](https://redocly.github.io/redoc/?url=https://raw.githubusercontent.com/open-traffic-generator/models/master/artifacts/openapi.yaml) is available as well. The OTG APIs support both REST and gRPC interfaces.

## Building Blocks

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
