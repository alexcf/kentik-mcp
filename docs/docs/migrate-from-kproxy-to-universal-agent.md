# Migrate from Kproxy to Universal Agent

Source: https://kb.kentik.com/docs/migrate-from-kproxy-to-universal-agent

---

This article covers how to migrate from Kentik’s legacy `kproxy` agent to **Universal Agent** for collecting NetFlow & SNMP telemetry from your network infrastructure devices.

## Overview

Still using Kentik’s `kproxy` standalone software agent? You’ll want to deploy Kentik’s **Universal Agent** (UA) and take advantage of its **Flow Proxy** and **SNMP/ST** capabilities as summarized in this migration checklist:

* **Deploy** UA and enable the **Flow Proxy** and/or **SNMP/ST** capabilities.
* **Update** your devices and access lists with UA's IP address.
* **Configure** devices in Kentik for SNMP flow enrichment or full monitoring.
* **Decommission** the legacy `kproxy` agent.

> **Did You Know?**: Flow enrichment is included when the device has a **FlowPak** license. Full monitoring requires a Kentik NMS **DevicePak** license, and includes many more features including deeper metric visibility, health monitoring, syslog, traps, and **SSH access for configs, scrapes and ad-hoc investigation with AI Advisor**.

### Why Migrate?

Kentik’s **Universal Agent** (UA) provides a centralized framework for ingesting network telemetry. Consolidating telemetry collection into the UA architecture offers several operational benefits:

* **Future-Ready Innovation:** The UA is the designated agent platform for all future Kentik feature investments. New capabilities, such as advanced **NetBox** integrations for infrastructure correlation, are developed exclusively for the UA framework.
* Resource Efficiency: The UA serves as a high-performance source for both Flow data and SNMP metadata. This centralized approach optimizes the processing load on network infrastructure by coordinating telemetry requests through a single, efficient framework.
* Architectural Consolidation: The UA replaces disparate installation procedures and security audit surfaces with a single, managed extension of the Kentik SaaS platform.

## Prerequisites

Verify the following environment requirements before beginning the configuration:

* **Host Environment:** A physical or virtual host running a supported Linux distribution or an environment with Docker container support.
* **Agent Registration:** The Kentik UA is installed and registered within the Kentik portal.
* **Access Control:** Administrator permissions for the Kentik portal and appropriate network access (UDP/SNMP) between the host and managed network devices.

Next, let’s take a look at some **Migration Scenarios**.
