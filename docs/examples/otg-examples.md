# OTG Examples

## Overview 

[Open Traffic Generator examples](https://github.com/open-traffic-generator/otg-examples) repository is a great way to get started. It features a collection of software-only network labs ranging from very simple to more complex. To setup network labs in software we use containerized or virtualized NOS images.

## Device Under Test

Many network vendors provide versions of their Network Operating Systems as a CNF or VNF. To make OTG Examples available for a widest range of users, our labs use open-source or freely available NOSes like FRR. Replacing FRR with a container from a different vendor is a matter of modifying one of the lab examples.

Some examples don't have any DUT and use back-2-back connections between Test Ports. These are quite useful to make sure the Traffic Generator part works just fine by itself, before introducing a DUT.

## Infrastructure

To manage deployment of the example labs, we use one of the following declarative tools:

* [Docker Compose](https://docs.docker.com/compose/) - general-purpose tool for defining and running multi-container Docker applications
* [Containerlab](https://containerlab.dev/) - simple yet powerful specialized tool for orchestrating and managing container-based networking labs

## CI with Github Actions

Some of the lab examples include Github Action workflow for executing OTG tests on any changes to the lab code. This could serve as a template for your CI workflow.

## Reference

| Lab                                                                                                                     | OTG Tool    | DUT  | Client     | Infrastructure | CI  |
| ----------------------------------------------------------------------------------------------------------------------- | ----------- | ---- | ---------- | -------------- | --- |
| [Ixia-c traffic engine](https://github.com/open-traffic-generator/otg-examples/blob/main/docker-compose/b2b)            | Ixia-c TE   | B2B  | `otgen`    | Compose        | yes |
| [KENG 3 pairs](https://github.com/open-traffic-generator/otg-examples/blob/main/docker-compose/b2b-3pair)               | KENG TE     | B2B  | `otgen`    | Compose        | no  |
| [KENG BGP and traffic](https://github.com/open-traffic-generator/otg-examples/blob/main/docker-compose/cpdp-b2b)        | KENG PE+TE  | B2B  | `gosnappi` | Compose        | yes |
| [Hello, snappi! Welcome to the Clab!](https://github.com/open-traffic-generator/otg-examples/blob/main/clab/ixia-c-b2b) | Ixia-c-one  | B2B  | `snappi`   | Containerlab   | no  |
| [Ixia-c-one and FRR](https://github.com/open-traffic-generator/otg-examples/blob/main/clab/ixia-c-te-frr)               | Ixia-c FE   | FRR  | `otgen`    | Containerlab   | no  |
| [Remote Triggered Black Hole](https://github.com/open-traffic-generator/otg-examples/blob/main/clab/rtbh)               | Ixia-c-one  | FRR  | `gosnappi` | Containerlab   | yes |