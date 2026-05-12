# Device Configs Directory

Source: https://kb.kentik.com/docs/device-configs-directory

---

This article provides configuration snippets for network infrastructure devices (routers, switches, etc.) to integrate with Kentik. Each section includes the recommended configuration type and a direct link to the full repository on GitHub.

> **Note**: For all devices below, the recommended method is to configure the device to push flow via the Flow Proxy capability of Kentik’s **Universal Agent**.

## Quick Tips

* **Read First**: Review the primary **README**before applying any configurations.
* **Check Compatibility**: Look for a `compatible-platforms.md` file in the specific vendor directory for any confirmed platforms that work with the Kentik production environment.
* **Vendor Notes**: Many directories contain specific `README.md` files with essential background information.

## Arista Configs

**View all Arista snippets on GitHub**

| Type | Description | File |
| --- | --- | --- |
| Flow | sFlow via **Universal Agent** (Recommended) | `sflow-agent.conf` |

## Cisco Configs

**View all Cisco snippets on GitHub**

| Series | GitHub Repository | Recommended File |
| --- | --- | --- |
| 6500-7600 | **Link** | `netflow-9-agent.conf` |
| ASA | **Link** | `netflow-9-agent.conf` |
| Nexus 3000 | **Link** | `sflow-agent.conf` |
| Nexus 6000-7000 | **Link** | `netflow-9-agent.conf` |
| IOS XE | **Link** | `netflow-9-agent.conf` |
| IOS XR | **Link** | `netflow-9-agent.conf` |

## Juniper Configs

**View all Juniper configs on GitHub**

| Series | GitHub Repository | Recommended File |
| --- | --- | --- |
| MX Series | **Link** | `ipfix-agent.conf` |
| QFX/EX Series | **Link** | `sflow-agent.conf` |

## Additional Vendor Configs

| Vendor | GitHub Repository | Recommended File |
| --- | --- | --- |
| Extreme | **Link** | `sflow-agent.conf` |
| Huawei | **Link** | `netflow-9-agent.conf` |
| Nokia | **Link** | `ipfix-agent.conf` |
| Palo Alto | **Link** | `netflow-9-agent.conf` |
| Vyatta | **Link** | `sflow-agent.conf` |
