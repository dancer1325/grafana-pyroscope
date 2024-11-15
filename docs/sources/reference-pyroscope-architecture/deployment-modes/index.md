---
title: "Pyroscope deployment modes"
menuTitle: "Deployment modes"
description: "You can deploy Pyroscope in either monolithic mode or microservices mode."
weight: 20
---

# Pyroscope deployment modes

* == ways to deploy Pyroscope
  * Monolithic mode
    * ALL components run | 1! process
    * uses
      * you ONLY need 1 pyroscope instance
  * Microservices mode
    * components are deployed | distinct processes 
      * -> `-target=componentName` / EACH process
        * _Example:_ `-target=ingester` or `-target=distributor`
    * allows
      * scaling out the # of instance
        * -> 1! backend / shared for
          * storage
          * querying
      * scaling / component
* determined -- via -- `-target` parameter /
  * ways to set
    * CLI flag
    * YAML configuration

## Monolithic mode

* 👀== ALL components run simultaneously | 1! process 👀
  * the simplest way to deploy Pyroscope
  * if you want to see the list of components ->
    ```bash
    # use `-modules`
    ./pyroscope -modules
    ```
* default mode of operation
* way to specify it
  * `-target=all`
* uses
  * get started quickly
  * work with Pyroscope | development environment

![Pyroscope's monolithic mode](monolithic-mode.svg)

[//]: # "Diagram source at https://docs.google.com/presentation/d/1C1fl0pH8wmKZe8gXo-VwmUuLvGiPmADfvey15FSkWpE/edit#slide=id.g11694eaa76e_0_0"

* if you want to horizontally scaled out -> deploy MULTIPLE Pyroscope binaries -- via -- `-target=all`
  * vs [FULL microservices deployment](#microservices-mode)
    * NO complex configuration 

    ![Pyroscope's horizontally scaled monolithic mode](scaled-monolithic-mode.svg)

[//]: # "Diagram source at https://docs.google.com/presentation/d/1C1fl0pH8wmKZe8gXo-VwmUuLvGiPmADfvey15FSkWpE/edit#slide=id.g11658e7e4c6_1_20"

## Microservices mode

* MOST complex
* scaling / component ->
  * greater flexibility | scaling
  * MORE granular failure domains
* uses
  * production deployment
* requirements to get a working Pyroscope instance
  * 👀deploy EVERY required component👀
    * see [components](../components)
* recommendations
  * use [Kubernetes](https://kubernetes.io/)

[//]: # "Diagram source at https://docs.google.com/presentation/d/1C1fl0pH8wmKZe8gXo-VwmUuLvGiPmADfvey15FSkWpE/edit#slide=id.g11658e7e4c6_1_53"

![Pyroscope's microservices mode](microservices-mode.svg)
