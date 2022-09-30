# Model

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

The specification is available on GitHub [here](https://github.com/open-traffic-generator/models/blob/master/artifacts/openapi.yaml) (viewer friendly version at reDocly [here](https://redocly.github.io/redoc/?url=https://raw.githubusercontent.com/open-traffic-generator/models/master/artifacts/openapi.yaml)).  Specifications support both REST or gRPC interfaces.
