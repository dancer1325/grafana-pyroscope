---
headless: true
description: Shared file for available profile types.
---

[//]: # 'This file documents the available profile types in Pyroscope.'
[//]: # 'This shared file is included in these locations:'
[//]: # '/pyroscope/docs/sources/configure-client/profile-types.md'
[//]: # '/pyroscope/docs/sources/introduction/profiling-types.md'
[//]: #
[//]: # 'If you make changes to this file, verify that the meaning and content are not changed in any place where the file is included.'
[//]: # 'Any links should be fully qualified and not relative: /docs/grafana/ instead of ../grafana/.'

* supported profile types -- by -- Pyroscope
  * CPU (CPU time, wall time)
  * Memory (allocation objects, allocation space, heap)
  * In use objects and in-use space
  * Goroutines
  * Mutex count & duration
    * mutex
      * == mutual exclusion
      * == programming construct /
        * uses
          * manage access to shared resources | multithreaded or concurrent environment 
  * Block count & duration
  * Lock count & duration
  * Exceptions