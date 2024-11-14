---
title: Introduction
menuTitle: Introduction
description: Learn about Pyrsocope and profiling.
weight: 200
keywords:
  - Pyroscope
  - Profiling
---

# Introduction

{{< youtube id="XL2yTCPy2e0" >}}

## Why Pyroscope

* see [main index](../_index.md)

## Core functionality

* Minimal CPU overhead & efficient compression
  * minimal CPU overhead
    * == the process does NOT take up much of the CPU
* Architecture -- aligned with -- Grafana Loki, Mimir, and Tempo ==
    - Horizontally scalable
      - == Grafana Pyroscope can run | MULTIPLE machines
    - Reliable
      - == HIGHLY available
    - Multi-tenancy Support
      - 👀== possible to run 1 database / MULTIPLE independent teams or business units 👀
    - Cost Effective | Scale
      - extensive historical data storage -- via -- object storage / NO significant costs
      - -- compatible with -- MULTIPLE object store implementations
        - _Example:_ AWS S3, Google Cloud Storage, Azure Blob Storage, OpenStack Swift, 
* Advanced Analysis UI
  * allowed adding tag/label
  * ways to differentiate performance
    * tags/labels
    * time intervals
