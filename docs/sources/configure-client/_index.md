---
aliases:
  - /docs/phlare/latest/operators-guide/configure-agent/
  - /docs/phlare/latest/configure-client/
title: "Configure the client to send profiles"
menuTitle: "Configure the client to send profiles"
description: "Learn how to configure the client to send profiles from your application."
weight: 300
---

# Configure the client to send profiles

* Pyroscope
  * == continuous profiling database /
    * allows
      * analyze the performance of your applications
    * ways to send profiles | Pyroscope
      1. Auto-instrumentation -- via -- Grafana Alloy
      2. SDK instrumentation
      3. SDK instrumentation -- through -- Grafana Alloy

    ![Pyroscope agent server diagram](https://grafana.com/media/docs/pyroscope/pyroscope_client_server_diagram_09_18_2024.png)

## auto-instrumentation -- via -- Grafana Alloy or Agent collectors

* collectors
  * send data from your application -- through them -- to Pyroscope
  * allows
    * 👀if you are using MULTIPLE applications or microservices -> you can centralize the profiling process / WITHOUT changing your application's codebase 👀
* Grafana Alloy & Grafana Agent collectors
  * support profiling / pull mode, with
    * eBPF,
    * Java,
    * Golang
* [Grafana Alloy](https://grafana.com/docs/alloy/latest/)
  * 👀== vendor-neutral distribution of the OpenTelemetry (OTel) Collector 👀
    * configurable -- via -- River file
    * component / -- runs alongside -- your application
  * recommended
  * requirements
    * install & configure the collector | SAME machine or container / run your application
  * how does it work?
    * periodically gathers profiling data -- from -- your application
    * the captured profiles -- are then sent to the -- Pyroscope server | 
      * store
      * analyze 
  * uses
    * collect profiles -- from -- applications / NOT modify their source code
  * eBPF profiling option
    * NOT need pull or push mechanisms

## About instrumentation with Pyroscope SDKs

* Pyroscope SDKs
  * uses
    * MORE precise & customizable profiling
  * requirements to use it
    * application / you are profiling is written | language / -- supported by the -- SDKs (for example, Java, Python, .NET, and others)
  * how to use it?
    * install the relevant Pyroscope SDK
    * instrument your application's code -- via the SDK, to capture the -- necessary profiling data
    * SDK automatically periodically -- pushes the -- captured profiles | Pyroscope server
      * Reason: 🧠for storage and analysis 🧠

## About instrumentation with Pyroscope SDKs through Alloy

* TODO:
Pyroscope SDKs can be configured to send profiles to Grafana Alloy first, which then forwards them to the Pyroscope server. This method combines the flexibility of SDK instrumentation with Alloy's infrastructure benefits.

Here's how it works:
1. Your application is instrumented with Pyroscope SDKs
2. Instead of sending profiles directly to Pyroscope, the SDK sends them to Alloy's `pyroscope.receive_http` component
3. Alloy processes and forwards the profiles to the Pyroscope server

By sending profiles through Alloy, you benefit from lower latency as profiles are sent to a local Alloy instance instead of directly over the internet to Grafana Cloud. Your application code remains focused on instrumentation while infrastructure concerns like authentication and routing are handled by Alloy's configuration. This separation allows for centralized management of metadata, where you can enrich profiles with infrastructure labels such as Kubernetes metadata or environment tags without modifying application code.

## Choose the right profiling method

You can use Grafana Alloy for auto-instrumentation, the Pyroscope instrumentation SDKs directly, or SDKs through Alloy. 
The method you choose depends on your specific use case and requirements.

Here are some factors to consider when making the choice:

- Ease of setup: Grafana Alloy is an ideal choice for a quick and straightforward setup without modifying your application's code. eBPF profiling supports some languages (for example, Golang, Python) better than others. More robust support for Java and other languages is in development. Using SDKs through Alloy adds minimal setup complexity while providing infrastructure benefits.
- Language support: If you want more control over the profiling process and your application is written in a language supported by the Pyroscope SDKs, consider using the SDKs - either directly or through Alloy depending on your infrastructure needs.
- Flexibility: The Pyroscope SDKs offer greater flexibility in terms of customizing the profiling process and capturing specific sections of code with labels. If you have particular profiling needs or want to fine-tune the data collection process, the SDKs would be your best bet. When used with Alloy, you gain additional infrastructure flexibility without compromising SDK capabilities.

To get started, choose one of the integrations below:
<table>
   <tr>
      <td align="center"><a href="https://grafana.com/docs/pyroscope/latest/configure-client/grafana-alloy/go_pull"><img src="/media/docs/alloy/alloy_icon.png" width="100px;" alt=""/><br />
        <b>Grafana Alloy</b></a><br />
          <a href="https://grafana.com/docs/pyroscope/latest/configure-client/grafana-alloy/go_pull/" title="Documentation">Documentation</a><br />
          <a href="https://github.com/grafana/pyroscope/tree/main/examples/grafana-agent-auto-instrumentation" title="examples">Examples</a>
      </td>
      <td align="center"><a href="https://grafana.com/docs/pyroscope/latest/configure-client/language-sdks/go_push/"><img src="https://user-images.githubusercontent.com/23323466/178160549-2d69a325-56ec-4e19-bca7-d460d400b163.png" width="100px;" alt=""/><br />
        <b>Golang</b></a><br />
          <a href="https://grafana.com/docs/pyroscope/latest/configure-client/language-sdks/go_push/" title="Documentation">Documentation</a><br />
          <a href="https://github.com/grafana/pyroscope/tree/main/examples/language-sdk-instrumentation/golang-push" title="golang-examples">Examples</a>
      </td>
      <td align="center"><a href="https://grafana.com/docs/pyroscope/latest/configure-client/language-sdks/java/"><img src="https://user-images.githubusercontent.com/23323466/178160550-2b5a623a-0f4c-4911-923f-2c825784d45d.png" width="100px;" alt=""/><br />
        <b>Java</b></a><br />
          <a href="https://grafana.com/docs/pyroscope/latest/configure-client/language-sdks/java/">Documentation</a><br />
          <a href="https://github.com/grafana/pyroscope/tree/main/examples/language-sdk-instrumentation/java/rideshare" title="java-examples">Examples</a>
      </td>
      <td align="center"><a href="https://grafana.com/docs/pyroscope/latest/configure-client/grafana-alloy/ebpf"><img src="https://user-images.githubusercontent.com/23323466/178160548-e974c080-808d-4c5d-be9b-c983a319b037.png" width="100px;" alt=""/><br />
        <b>eBPF</b></a><br />
          <a href="https://grafana.com/docs/pyroscope/latest/configure-client/grafana-alloy/ebpf" title="Documentation">Documentation</a><br />
          <a href="https://github.com/grafana/pyroscope/tree/main/examples/grafana-agent-auto-instrumentation/ebpf" title="examples">Examples</a>
      </td>
      <td align="center"><a href="https://grafana.com/docs/pyroscope/latest/configure-client/language-sdks/python/"><img src="https://user-images.githubusercontent.com/23323466/178160553-c78b8c15-99b4-43f3-a2a0-252b6c4862b1.png" width="100px;" alt=""/><br />
        <b>Python</b></a><br />
          <a href="https://grafana.com/docs/pyroscope/latest/configure-client/language-sdks/python/" title="Documentation">Documentation</a><br />
          <a href="https://github.com/grafana/pyroscope/tree/main/examples/language-sdk-instrumentation/python" title="python-examples">Examples</a>
      </td>
   </tr>
   <tr>
      <td align="center"><a href="https://grafana.com/docs/pyroscope/latest/configure-client/language-sdks/dotnet/"><img src="https://user-images.githubusercontent.com/23323466/178160544-d2e189c6-a521-482c-a7dc-5375c1985e24.png" width="100px;" alt=""/><br />
        <b>Dotnet</b></a><br />
          <a href="https://grafana.com/docs/pyroscope/latest/configure-client/language-sdks/dotnet/" title="Documentation">Documentation</a><br />
          <a href="https://github.com/grafana/pyroscope/tree/main/examples/language-sdk-instrumentation/dotnet" title="examples">Examples</a>
      </td>
      <td align="center"><a href="https://grafana.com/docs/pyroscope/latest/configure-client/language-sdks/ruby/"><img src="https://user-images.githubusercontent.com/23323466/178160554-b0be2bc5-8574-4881-ac4c-7977c0b2c195.png" width="100px;" alt=""/><br />
        <b>Ruby</b></a><br />
          <a href="https://grafana.com/docs/pyroscope/latest/configure-client/language-sdks/ruby/" title="Documentation">Documentation</a><br />
          <a href="https://github.com/grafana/pyroscope/tree/main/examples/language-sdk-instrumentation/ruby" title="ruby-examples">Examples</a>
      </td>
      <td align="center"><a href="https://grafana.com/docs/pyroscope/latest/configure-client/language-sdks/nodejs/"><img src="https://user-images.githubusercontent.com/23323466/178160551-a79ee6ff-a5d6-419e-89e6-39047cb08126.png" width="100px;" alt=""/><br />
        <b>Node.js</b></a><br />
          <a href="https://grafana.com/docs/pyroscope/latest/configure-client/language-sdks/nodejs/" title="Documentation">Documentation</a><br />
          <a href="https://github.com/grafana/pyroscope/tree/main/examples/language-sdk-instrumentation/nodejs/express" title="examples">Examples</a>
      </td>
      <td align="center"><a href="https://grafana.com/docs/pyroscope/latest/configure-client/language-sdks/rust/"><img src="https://user-images.githubusercontent.com/23323466/178160555-fb6aeee7-5d31-4bcb-9e3e-41e9f2f7d5b4.png" width="100px;" alt=""/><br />
        <b>Rust</b></a><br />
          <a href="https://grafana.com/docs/pyroscope/latest/configure-client/language-sdks/rust/" title="Documentation">Documentation</a><br />
          <a href="https://github.com/grafana/pyroscope/tree/main/examples/language-sdk-instrumentation/rust/rideshare" title="examples">Examples</a>
      </td>
   </tr>
</table>

## Enriching profile data

You can add tags to your profiles to help correlate them with your other telemetry signals.
Commonly used tags include version, region, environment, and request types.
You have the ability to add tags using both the SDK and Alloy.

Valid tag formats may contain ASCII letters and digits, as well as underscores. It must match the regex `[a-zA-Z_][a-zA-Z0-9_]`.
In Pyroscope, a period (`.`) isn't a valid character inside of tags and labels.

## Assistance with Pyroscope

If you have more questions, feel free to reach out in [the community Slack channel](https://grafana.slack.com/) or create an [issue on GitHub](https://github.com/grafana/pyroscope).
