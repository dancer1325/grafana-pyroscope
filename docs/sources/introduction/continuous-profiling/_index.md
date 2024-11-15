---
title: When to use continuous profiling
menuTitle: When to use continuous profiling
description: Discover the benefits of continuous profiling and its role in modern application performance analysis.
weight: 200
keywords:
  - pyroscope
  - phlare
  - continuous profiling
  - flame graphs
---

* see [continuous profiling](../profiling.md#continuous-profiling)
* alternative to pyroscope
  * run a benchmark tool locally and get a .pprof | Go or .jrf | Java
    * 👀NOT recommendation | production 👀
* how does it work?
  * collect profiles | production systems
  * stores the profiles | database

## Benefits

* see [continuous profiling](../profiling.md#continuous-profiling)

![3 benefits of continuous profiling](https://grafana.com/static/img/pyroscope/profiling-use-cases-diagram.png)

## Use cases

* see [continuous profiling](../profiling.md#continuous-profiling)
![Infographic illustrating key business benefits](https://grafana.com/static/img/pyroscope/cost-cutting-diagram.png)

### Reduce operational costs

* operational costs -- come from --
  * infrastructure
  * observability,
  * incident management,
  * messaging/queuing,
  * deployment tools

* costs of using Pyroscope
  * minimal overhead | using sampling profilers 
    * [2%, 5%] -- depending on a -- few factors
  * store the data
    * low
      * Reason: 🧠efficient compresses 🧠

### Reduced latency

* Reason: 🧠identifying and addressing performance bottlenecks before 🧠
* -> 
  * faster application response times,
  * improved user experience -> better business outcomes
  * revenue

### Enhanced incident management

* Reason: 🧠immediate, actionable insights shared 🧠
