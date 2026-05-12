# SSH Command Access for Kentik NMS

Source: https://kb.kentik.com/docs/ssh-command-access-for-kentik-nms

---

This article covers how to configure and use **Kentik NMS** with **AI Advisor** to assist with NetOps.

![AI Advisor performing an NMS device configuration analysis, highlighting changes and their impacts.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/ANMS-device-details-AIA-overlay.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A25Z&se=2026-05-12T09%3A30%3A25Z&sr=c&sp=r&sig=wT86BGPnuW0HsE2e%2BhklwyUNNNGdXmQJeYXN9GiNPDI%3D)

*AI Advisor performing an NMS device configuration analysis, highlighting changes and their impacts.*

## Overview

**Kentik NMS** has two SSH-powered features to supercharge **NetOps** practitioners with better network intelligence:

* **Command Access**(On-Demand Device Diagnostics)
* **Config Context** (Backups and Diffs)

With these features, **AI Advisor** becomes more powerful: It can access, analyze, and interpret device configurations and, with user permission, pull real-time operational data from devices during live investigations using read-only show commands via SSH.

`"Show me why my network performance dropped after last night's edge router change."`

> **IMPORTANT**: **Early Access Feature**
>
> * This feature is in open Early Access. Functionality may change.
> * Requires an active NMS device license (see **Licenses**).
> * AI Advisor is **read-only**. It cannot push changes or execute `write memory`/`commit` commands (see **SSH Access & Security Policy**).

## Prerequisites

Before configuring AI Advisor for network device analysis, ensure your environment meets the following requirements:

* **Licensing**: An active NMS Device **License** is required for each device you intend to monitor.
* **Kentik Universal Agent**: You must have the latest version of the **Universal Agent** deployed within your infrastructure with network reachability to the target devices.
* **Connectivity**: Port TCP/22 (SSH) must be open between the Universal Agent and the device management IP.

  > **Note**: The Universal Agent must have a direct route to the device's Management IP; it does not scrape configuration data via the data plane.
* **Supported Platforms**: Ensure your device OS is supported for configuration scraping:

  | Platform | Support Level | Recommended Role |
  | --- | --- | --- |
  | **Juniper Junos** | Full | `read-only` class |
  | **Cisco NX-OS** | Full | `network-operator` |
  | **Arista EOS** | Full | `network-operator` |
  | **Cisco IOS-XE** | Full | `Parser View` (Custom) |

* **Credentials**: A read-only SSH user account must be configured on the device (see **Configuration Examples** for platform-specific templates).

## SSH Access & Security Policy

Kentik brings context to your observability data by securely pulling state data via SSH.

* **Read-Only**: Kentik does not request, nor want, write access to your infrastructure.
* **Zero Configuration Changes**: Kentik’s features are designed to scrape config diffs and execute ad-hoc troubleshooting commands. Kentik will never execute `configure terminal` or `commit` changes.
* **Audit Trail**: All activity is initiated via your local **Collection Agent**, ensuring every command is logged in your local AAA (TACACS+/RADIUS) systems.
