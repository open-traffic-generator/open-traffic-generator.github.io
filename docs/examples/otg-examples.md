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

| Lab                                                                                                             | OTG Tool    | DUT  | Client     | Infrastructure | CI  | Description                          |
| --------------------------------------------------------------------------------------------------------------- | ----------- | ---- | ---------- | -------------- | --- | ------------------------------------ |
| [**`dp-b2b`**](https://github.com/open-traffic-generator/otg-examples/blob/main/docker-compose/b2b)             | Ixia-c TE   | B2B  | `otgen`    | Compose        | yes | Ixia-c traffic engine back-to-back   |
| [**`dp-b2b-3pair`**](https://github.com/open-traffic-generator/otg-examples/blob/main/docker-compose/b2b-3pair) | KENG TE     | B2B  | `otgen`    | Compose        | no  | KENG 3 back-to-back pairs            |
| [**`cpdp-b2b`**](https://github.com/open-traffic-generator/otg-examples/blob/main/docker-compose/cpdp-b2b)      | KENG PE+TE  | B2B  | `gosnappi` | Compose        | yes | KENG back-to-back BGP and traffic    |
| [**`dp-b2b`** ](https://github.com/open-traffic-generator/otg-examples/blob/main/clab/ixia-c-b2b)               | Ixia-c-one  | B2B  | `snappi`   | Containerlab   | no  | Hello, snappi! Welcome to the Clab!  |
| [**`dp-frr`**](https://github.com/open-traffic-generator/otg-examples/blob/main/clab/ixia-c-te-frr)             | Ixia-c FE   | FRR  | `otgen`    | Containerlab   | no  | Ixia-c Traffic Engine and FRR        |
| [**`rtbh`**](https://github.com/open-traffic-generator/otg-examples/blob/main/clab/rtbh)                        | Ixia-c-one  | FRR  | `gosnappi` | Containerlab   | yes | Remote Triggered Black Hole Lab      |
