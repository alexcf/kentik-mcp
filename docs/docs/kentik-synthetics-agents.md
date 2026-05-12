# Synthetics Agent Overview

Source: https://kb.kentik.com/docs/kentik-synthetics-agents

---

This article provides an overview of Kentik's `ksynth-agent` synthetics agent.

> **Notes:**
>
> * For a general explanation of synthetic monitoring in Kentik, see **Synthetics**.
> * For assistance with any aspect of the agent setup process, contact Kentik (see **Customer Care**).

## About Synthetics Agent

Kentik is able to perform continuous synthetic performance testing from and to network locations both within your infrastructure and elsewhere around the Internet. These tests are enabled by Kentik’s `ksynth-agent` software agent for synthetic monitoring.

## Synthetics Agent Types

The `ksynth-agent` software agent for synthetic testing includes the following modes, known as agent types:

| Synthetics Agent Type | Description |
| --- | --- |
| **App** | * Supports all test types (network, page load, transaction) * Runs with **Headless Chromium**, so has a heavier resource footprint |
| **Network** | Supports network tests only (ping, traceroute, HTTP, DNS) |

## Synthetics Agent Deployments

Kentik's synthetic testing agent is used in the following types of deployments:

* **Global Agents**: Available to every customer that has activated our synthetic monitoring services, "global agents" are the agents in the Kentik Global Agent Network, a worldwide network of Kentik-maintained `ksynth` agents. Hosted in cloud offerings, these agents enable performance testing to and from key Internet hubs worldwide.

  > **Note:** In some portal contexts, the subset of global agents that is deployed in the infrastructure of key cloud service providers (AWS, GCP, Azure, OCI, etc.) is referred to separately as "Public Cloud" agents.
* **Private Agents**: Every Kentik customer can deploy as many `ksynth-agent` instances as they care to in their own on-prem and/or cloud infrastructure (no additional license required). These private agents are for the exclusive use of the customer who deploys them (not available to other Kentik customers).

## Deployment Considerations

The following considerations apply when deploying `ksynth-agent`.

### `ksynth-agent` Protocols

`ksynth-agent` uses the following protocols to support testing and monitoring.

#### Communication with the Kentik Platform

This is the control and reporting channel for the agent to communicate with the Kentik backend (US or EU regions).

| Protocol | Port | Description |
| --- | --- | --- |
| TCP | 443 | Used for HTTPS (secure communication) to send test results, status updates, and receive new test instructions. |

#### Network Tests (IPv4 and/or IPv6)

These protocols are used by the agent to actively probe network performance and pathing.

| Test Type | Protocols Used | Configuration Options |
| --- | --- | --- |
| Ping | ICMP, TCP, or "UDP-ICMP" | TCP ping allows specifying a target port. “UDP-ICMP” means that Kentik attempts a UDP connection, while accepting ICMP denials as successful response. |
| Traceroute | ICMP, UDP, or TCP | TCP traceroute allows specifying a target port. |
| General Ping | ICMP Echo ping | Standard ICMP request/reply. |

#### HTTP or API Tests

These protocols are used for application and service availability monitoring.

| Request Method | Protocol | Target Ports |
| --- | --- | --- |
| HTTP GET/POST/PUT/PATCH | HTTP(s) | TCP 80 (HTTP) or TCP 443 (HTTPS) |

### Allowing `ksynth-agent`

When running one or more `ksynth-agent` instances on your network, the domains in the table below should be allowed in your firewall rules to ensure that the agents can communicate with Kentik.

| US domains | EU Domains |
| --- | --- |
| api.kentik.com | api.kentik.eu |
| flow.kentik.com | flow.kentik.eu |
| grpc.api.kentik.com | grpc.api.kentik.eu |
| portal.kentik.com | portal.kentik.eu |
| storage.googleapis.com | storage.googleapis.com |
| whoami.kentiklabs.com | whoami.kentiklabs.com |
| dns.google/resolve | dns.google/resolve |

### NTP Configuration

The server on which `ksynth-agent` is installed must be configured as an NTP client to avoid known issues related to clock skew. If the NTP service is correctly configured, the following command should return successfully:

```
sudo ntpq -p
```

### `ksynth-agent` Security

The `ksynth-agent` agent generates a unique identity and uses this to authenticate with the Kentik platform. The identity is stored in a local file to be used across restarts and upgrades.
