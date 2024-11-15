---
title: Pyroscope and profiling in Grafana
menuTitle: Pyroscope in Grafana
description: Learn about how you can use profile data in Grafana.
weight: 400
keywords:
  - Pyroscope
  - Profiling
  - Grafana
---

<!-- This is placeholder page while we get the content written.  -->

# Pyroscope and profiling in Grafana

* Pyroscope
  * -- can be used alongside -- OTHER Grafana tools (Loki, Tempo, Mimir, and k6)
  * 👀if you add [Pyroscope data source plugin](https://grafana.com/docs/grafana/latest/datasources/pyroscope/) -> Pyroscope can be used | Grafana 👀

## Visualize traces and profiles data

![trace-profiler-view](https://grafana.com/static/img/pyroscope/pyroscope-trace-profiler-view-2023-11-30.png)
* **Explore** page 
  * display combined traces + profiles
  * uses
    * granular line-level detail
    * see the exact function & specific request / cause a bottleneck in your application

## Integrate profiles | dashboards

![dashboard](https://grafana.com/static/img/pyroscope/grafana-pyroscope-dashboard-2023-11-30.png)

* memory profiles -- alongside -- panels for logs & metrics
