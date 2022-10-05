# Clients

## Overview
 
To successfully use an OTG-based Traffic Generator, you need to be able to execute the following tasks over the OTG API:
 
* Prepare a **Configuration** and apply it to a Traffic Generator
* **Control** states of the configured objects like Protocols or Traffic Flows
* Collect and analyze **Metrics** reported by the Traffic Generator
 
It is a job of an **OTG Client** to perform these tasks by communicating with a Traffic Generator via the OTG API. There are different types of such clients, and the choice between them depends on how and where you want to use a Traffic Generator.
 
## Command-line Utilities
 
Command-line utilities is the fastest way to execute an OTG API request, especially in the environment without access to Python or Go development toolchains. They are also the easiest option to use by someone who is just starting with the OTG.
 
The most basic utility for any kind of REST API calls, including OTG, is **`curl`**. But, it leaves all the knowledge of OTG to the user. See the [curl](clients/curl.md) section for examples of using it with OTG.
 
On the other side of the spectrum is **`otgen`**. This command-line utility comes as part of OTG toolkit. It is capable of manipulating a wide range of OTG features while hiding a lot of complexity from a user. `otgen` is an excelent choice for a first time OTG user to become familiar with this ecosystem. It is also a great choice for an advanced user who needs to execute a common OTG scenario in an environment without a development toolchain. Continue reading about this utility in the [otgen](clients/otgen.md) section.
 
## Standalone Test Programs
 
Standalone Test Programs can be developed in Python, Go or other programming languages. They might be useful to validate a particular scenario. For example, to measure a packet loss duration caused by an artificially introduced failure in the network. Typically, such programs would be executed directly when needed, sometimes from a shell script. Often, these programs are a part of a Proof of Concept setup, or a technology demonstration.
 
To make use of the OTG API easier while developing such programs, **snappi** or **gosnappi** libraries should be leveraged. To find more information about these libraries follow to:
 
* [snappi](clients/snappi.md) section for Python
* [gosnappi](clients/gosnappi.md) section for Go
 
## Testing Frameworks
 
Testing frameworks like Pytest are an established choice for implementing CI/CD pipelines that should run automatically without human intervention. When applied to network test, these pipelines provide opportunities such as:
 
* validate NOS upgrades on network equipment to identify any breaking changes early
* certify compatibility between layers in a disaggregated whitebox or SDN solution
* pre-screen network configuration changes to reduce outages
 
This is the most common category of OTG applications as of today. Most, if not all OTG-based CI/CD pipelines make use of **snappi** or **gosnappi** libraries.
 
## Network Test Runners
 
One of the notable OTG-based developments in Network CI is the OpenConfig project and its ONDATRA test runner. Learn more about this in the [ONDATRA section](clients/ondatra.md).