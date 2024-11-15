---
title: Understand profiling types and their uses in Pyroscope
menuTitle: Understand profiling types
description: Learn about the different profiling types available in Pyroscope and how to effectively use them in your application performance analysis.
weight: 300
keywords:
  - profiles
  - profiling types
  - application performance
  - flame graphs
---

# Understand profiling types and their uses in Pyroscope

* profiling types
  * see [profiling types](../profiling.md)
  * 👀-- depend on -- different aspects to analyze of your application 👀
    * see [profiling classification](../../shared/available-profile-types.md)

## CPU profiling

* goal
  * measures the amount of CPU time / -- consumed by -- different parts of your application code
* high CPU usage == inefficient code ->
  * poor performance
  * increased operational costs
* uses 
  * about CPU-intensive functions | your application
    * identify
    * optimize 
* Flame graph insight 
  * 👀width of blocks == CPU time consumed / EACH function 👀
* _Example:_ CPU along with the flame graph

    ![example flame graph](https://grafana.com/static/img/pyroscope/pyroscope-ui-single-2023-11-30.png)

<!-- ## FGprof (for go)
[todo add a link to the docs for fgprof]  -->

## Memory allocation profiling

* goal
  * tracks the amount & frequency of memory allocations -- by the -- application
* excessive or inefficient memory allocation -> memory leaks & high garbage collection overhead
* types
  * Alloc Objects,
  * Alloc Space
* Flame graph insight
  * width of blocks == function's memory allocation
* timeline == memory allocations | time
* uses
  * identify memory leaks
    * Reason: 🧠 gradual increase in memory allocations / NEVER goes down | timeline 🧠

    ![memory leak example](https://grafana.com/static/img/pyroscope/pyroscope-memory-leak-2023-11-30.png)

## Goroutine profiling

* [Goroutines](https://go.dev/tour/concurrency/1)
* goal
  * measures the usage and performance of goroutines thread
* if poor management -> issues (_Example:_ deadlocks, excessive resource usage)
* Flame graph insight
  * == view of goroutine distribution and issues

## Mutex profiling

* goal
  * analyze mutex locks
* uses
  * prevent simultaneous access | shared resources
    * -> 
      * delays
      * reduced application throughput
* types
  * Mutex Count,
  * Mutex Duration
* Flame graph insight
  * == frequency & duration of mutex operations

## Block profiling

* block operations
  * == thread is paused or delayed
  * impact
    * slow down application processes 
* goal
  * measures the frequency & duration of blocking operations 
* types
  * Block Count,
  * Block Duration
* Flame graph insight
  * identifies where & how long threads are blocked
