---
title: Profiling fundamentals
menuTitle: Profiling fundamentals
description: Discover the benefits of continuous profiling and its role in modern application performance analysis.
weight: 100
keywords:
  - pyroscope
  - continuous profiling
  - flame graphs
---

# Profiling fundamentals

* Profiling
  * ⭐️:= technique / measure & analyze the runtime behavior of a program ⭐️
    * if you profile a program -> identify parts of the program / consume the MOST resources (CPU time, memory, or I/O operations)
  * types
    * traditional
    * continuous

* 👀Pyroscope allows addressing BOTH types of profiling 👀

## Traditional profiling (non-continuous)

* named as 
  * "sample-based" 
    * profiler interrupts the program | regular intervals / capture the program's state each time
      * == take snapshots
  * "instrumentation-based" profiling
    * developers insert additional code | program / records information about its execution
    * cons
      * can alter the program's behavior -- due to the -- added code overhead 
* origin
  * early days of computing
    * goal
      * understand how a program -- utilized the -- limited computational resources available
* benefits
  * **Precision**
    * deep dive specific sections of the code
  * **Control**
    * developers can initiate profiling sessions
  * **Detailed reports**
* use cases
  * **When**
    * development or testing phases
  * **Advantages**
    * detailed insights | specific parts of code
  * **Disadvantages**
    * snapshot | time

## Continuous profiling

* vs traditional profiling
  * | software systems more complex and scale
    * Reason: 🧠 limitations of traditional profiling 🧠
  * issues NOT notice -- through -- limited profiling sessions

* := profiling / data is continuously collected & minimal overhead
  * helping to identify sporadic or long-term performance issues.

* benefits
  * **Consistent monitoring**
    * == expose 
      * immediate issues
      * long-term performance issues
  * **Proactive bottleneck detection**
    * performance bottlenecks are identified and addressed | BEFORE they
  * **Broad performance landscape**
    * available | different
      * platforms,
      * technology stack
      * OS
  * **Bridging the Dev-Prod gap**
    * highlight possible differences between development and production
      * Hardware discrepancies
      * Software inconsistencies
      * Real-world workload challenges
        * == potential pitfalls | real user interactions and loads
  * **Economical advantages**:
    * Resource optimization
    * Rapid problem resolution
  * **Non-intrusive operation**
  * **Real-time response**

* use cases
  * **When**
    * production environments or
    * extended performance tests
  * **Advantages**
    * continuous view of system performance
  * **Disadvantages**
    * less detailed than traditional profiling
      * Reason: 🧠need to minimize impact | running system 🧠
