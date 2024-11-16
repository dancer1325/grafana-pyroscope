---
aliases:
  - /docs/phlare/latest/operators-guide/getting-started/
  - /docs/phlare/latest/operators-guide/get-started/
  - /docs/phlare/latest/get-started/
description: Learn how to get started with Pyroscope.
menuTitle: Get started
title: Get started with Pyroscope
weight: 250
---

# Get started with Pyroscope

* goal
  * start Pyroscope in [monolith mode](../reference-pyroscope-architecture/deployment-modes)

{{< youtube id="XL2yTCPy2e0" >}}

## requirements

* install [Docker](https://docs.docker.com/engine/install/)

## Download and configure Pyroscope

1. download
   1. -- via -- docker
       ```bash
         docker pull grafana/pyroscope:latest
       ```
   2. -- via -- local binary
       ```bash
       # Download Pyroscope v1.0.0 and unpack it to the current folder
       curl -fL https://github.com/grafana/pyroscope/releases/download/v1.0.0/pyroscope_1.0.0_linux_amd64.tar.gz | tar xvz
       ```
2. Run Pyroscope
   1. -- via -- docker
     ```bash
     docker network create pyroscope-demo
     docker run --rm --name pyroscope --network=pyroscope-demo -p 4040:4040 grafana/pyroscope:latest
     # 4040   default port of Pyroscope
     ```
   2. -- via -- local binary
     ```bash
     ./pyroscope
     ```
3. Verify that Pyroscope is ready

      ```bash
      curl localhost:4040/ready
      ```

4. Configure Pyroscope -- to scrape -- profiles
   1. by default, it's -- configured to scrape -- itself
   2. if you want to collect more profiles -> instrument your application with
      1. SDK or
      2. Grafana Alloy

## Add a Pyroscope data source | Grafana & query data

1. run a local Grafana server -- via -- Docker

    ```bash
    docker run --rm --name=grafana \
      --network=pyroscope-demo \
      -p 3000:3000 \
      -e "GF_INSTALL_PLUGINS=grafana-pyroscope-app"\
      -e "GF_AUTH_ANONYMOUS_ENABLED=true" \
      -e "GF_AUTH_ANONYMOUS_ORG_ROLE=Admin" \
      -e "GF_AUTH_DISABLE_LOGIN_FORM=true" \
      grafana/grafana:main
    ```
2. open [http://localhost:3000/datasources](http://localhost:3000/datasources) | browser
   1. == Grafana URL
3. configure a Pyroscope data source / -- query the -- local Pyroscope server

   | Field | Value                                                                |
   | ----- | -------------------------------------------------------------------- |
   | Name  | Pyroscope                                                            |
   | URL   | [http://pyroscope:4040/](http://pyroscope:4040/) OR [http://host.docker.internal:4040/](http://host.docker.internal:4040/) if using Docker  |

* allows
  * querying profiles | [Grafana Explore](/docs/grafana/<GRAFANA_VERSION>/explore/)
  * create dashboard panels -- via the -- newly configured Pyroscope data source
