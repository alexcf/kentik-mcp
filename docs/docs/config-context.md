# Config Context

Source: https://kb.kentik.com/docs/config-context

---

Misconfigurations are a major cause of disruptions and outages. Being able to understand when changes occur dramatically reduces MTTI and MTTR.

The **Config Context** feature in **Kentik NMS** provides a contextual understanding of device configs and changes as they occur. It enables the autonomous scraping of configs from network devices, storing them as secret-redacted backups, comparing changes, and allowing you to talk to a config-informed **AI Advisor**.

> **IMPORTANT:** Ensure your environment meets the Prerequisites (including Universal Agent deployment and TCP/22 reachability) before attempting to configure device synchronization.

## Setup: Configuration Sync & Backup

To enable AI Advisor to "see" your device's state, navigate to the **SSH** tab in **Device Settings**.

![Configuration settings for SSH access and device diagnostics.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/ANMS-edit-device-ssh-tab.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A50Z&se=2026-05-12T09%3A31%3A50Z&sr=c&sp=r&sig=1G4pigCdqWyuDoxFi%2F8wy9ZGvBR2HILYlMcvJkM%2B5eQ%3D)

1. **Enable SSH command access:** Toggle this **On** (the master switch).
2. **Credential:** Select your pre-configured **SSH key or password**.
3. **Collection Agent:** Select the **agent** (e.g., `nomad-0`) with reachability to the device.
4. **Device Platform:** Recommended setting is **Autodetect**.
5. **SSH Hostname (or IP):** The Fully Qualified Domain Name (FQDN) or IP address the agent will use to connect.
6. **Enable device configuration sync and backup:** Toggle this **On** so that Kentik periodically scrapes the running config.
7. **Fetch Interval:** Set the frequency. A interval of **5-minutes** is recommended for high-visibility environments.

> **TIP**: If the connection fails initially, verify that the Universal Agent has accepted the device's SSH host key or that your firewall isn't blocking the initial handshake on TCP/22.

## Example AI Advisor Prompts

* **Collect & Store:** "When was the last successful backup for [Device Name], and were there any errors?"
* **View & Read:** "Show me the current running configuration for [Device Name]."
* **Compare & Diff:** "Summarize the impact of the configuration changes made to [Device Name] during last night's maintenance window."
* **Historical Search:** "Locate the last known working configuration snippet for VLAN 100 on any device in the branch office."

## Example AI Advisor Response

Recent configuration changes for **tor301.our1.kentik.com** (Device ID: 483375) have been retrieved:

**Configuration Revision Summary**

| Revision | Timestamp | Size | Status |
| --- | --- | --- | --- |
| 11 (Latest) | 2025-12-23 19:34:32 UTC | 42,780 bytes | ✅ Success |
| 10 | 2025-12-15 12:40:51 UTC | 42,408 bytes | ✅ Success |
| 9 | 2025-12-15 12:35:51 UTC | 42,780 bytes | ✅ Success |
| 8 | 2025-12-14 15:55:51 UTC | 42,408 bytes | ✅ Success |
| 7 | 2025-12-14 15:50:51 UTC | 42,780 bytes | ✅ Success |

**Recent Activity Analysis**

The most recent change (Revision 11) occurred just moments ago:

* **When:** 2025-12-23 19:34:32 UTC
* **Size Change:** Increased from 42,408 to 42,780 bytes (+372 bytes)
* **Last Commit:** 2025-11-18 23:47:11 UTC by user "alistair"

A pattern in the configuration sizes suggests that there may be configuration changes being applied and then reverted, or periodic updates occurring approximately every 5 minutes during certain time windows.

**Key Configuration Details (from latest revision)**

Device Information:

* **Model:** Juniper QFX5100-48S-6Q
* **Location:** iad1-dc3-5310 (Equinix DC3)
* **Junos Version:** 14.1X53-D46.7
* **Serial Number:** TA3718040593
* **Autonomous System:** AS4200000301

Active Protocols include:

* BGP (IPv4 and IPv6) with multiple peer groups
* OSPF and OSPFv3
* sFlow monitoring
* LLDP/LLDP-MED
* BFD for link monitoring
