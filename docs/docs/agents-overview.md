# Agents Overview

Source: https://kb.kentik.com/docs/agents-overview

---

This article provides an overview of software agents used with Kentik.

> **Note:** The software agents described here are available for both Debian/Ubuntu and CentOS/RHEL.

## About Kentik Agents

> **IMPORTANT**: The Universal Agent will gradually replace all legacy Kentik agents. Kentik recommends Migrating to the Universal Agent to take advantage of the latest performance optimizations and security updates.

Kentik uses various software agents to gather data and perform tasks supporting the platform. These tasks include:

* **NMS Metrics**: Collecting data via SNMP and Streaming Telemetry (see **Universal Agent**).
* **Flow Encryption**: Securing data in transit (see **Universal Agent**).
* **Host Traffic**: Collecting and sending flow directly from a host (see **Universal Agent**).
* **Synthetic Monitoring**: Performing active network testing (see **Kentik Synthetics Agents**).
* **Kubernetes**: Monitoring container orchestration clusters, (see **Kentik Kubernetes Agents**).
* **Data Export**: Sending enriched records to external systems (see **Using Kentik Firehose**).

The following table provides a quick reference to help you identify the correct agent based on your specific monitoring use case and deployment environment:

| Agent Type | Primary Use Case | Environment |
| --- | --- | --- |
| Universal Agent | SNMP, Flow, Telemetry | Linux (Debian, Ubuntu, CentOS, RHEL) |
| Synthetics Agent | Connectivity & Latency Testing | Global / Private Nodes |
| Kubernetes Agent | Pod & Service Visibility | K8s Clusters |
| Firehose | High-volume Data Export | Cloud / On-Prem |

## Universal Agent

The **Universal Agent** is Kentik’s next-generation observability solution. Unlike legacy specialized agents, the Universal Agent is a single, unified software package that can perform multiple monitoring tasks.

Once installed, you don’t need to download new binaries to expand your monitoring scope; you simply enable or disable specific **Capabilities** via configuration.

### What are Capabilities?

Think of Capabilities as modular features within the Universal Agent. Instead of managing five different services, you manage one agent that can be toggled to handle:

* **Standard Infrastructure:** Collecting Flow, SNMP, and Streaming Telemetry.
* **Host Visibility:** Monitoring local host traffic and performance.
* **Secure Proxy:** Encrypting and tunneling data from your private network to the Kentik platform.

### Key Benefits

* **Simplified Maintenance:** Update one agent to get the latest features across all monitoring types.
* **Reduced Footprint:** Lower CPU and memory overhead compared to running multiple standalone agents.
* **Dynamic Scaling:** Turn on **Agent Capabilities** as your network grows without needing to redeploy infrastructure.
