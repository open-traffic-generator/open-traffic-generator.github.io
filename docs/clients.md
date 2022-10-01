# Clients

There are multiple ways to communicate with a Traffic Generator via the OTG API:
 <!-- TODO add links from bold items to paragraphs in Clients section -->
* **otgen** command-line tool is an easy way to start with
* **snappi** library to accelerate development of *Test Programs* written in Python or Go
* **direct REST or gRPC calls** as an alternative to using *snappi*
* **custom** OTG client applications

[snappi](https://pypi.org/project/snappi/) and [gosnappi](https://pkg.go.dev/github.com/open-traffic-generator/snappi/gosnappi) provide client side API libraries for the OTG specifications for Python and Golang respectively.  For other languages, SDKs can be built using [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator) (for REST API) or [protobuf tools](https://github.com/protocolbuffers/protobuf) (for gRPC).  