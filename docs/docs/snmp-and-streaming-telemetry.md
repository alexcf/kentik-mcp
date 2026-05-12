# SNMP and Streaming Telemetry

Source: https://kb.kentik.com/docs/snmp-and-streaming-telemetry

---

This article covers Kentik’s implementation of Simple Network Management Protocol (SNMP) and Streaming Telemetry within the Kentik Data Engine (KDE).

> **Note:** For information about NMS-specific implementations, see **NMS via Streaming Telemetry**.

![Streaming telemetry data shown in Data Explorer.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/ST-Main-1020w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A10Z&se=2026-05-12T09%3A30%3A10Z&sr=c&sp=r&sig=lqMM%2FGIFv4v0YIoO96Q5CkdPz2hhtXhPgVhFKGqUONA%3D)

*Streaming Telemetry data shown in Data Explorer.*

## SNMP vs. Streaming Telemetry

While both methods provide visibility into network performance, SNMP and ST differ in how data is retrieved and processed:

| Feature | SNMP (Simple Network Management Protocol) | Streaming Telemetry (ST) |
| --- | --- | --- |
| Model | **Pull**: The monitoring system polls devices at intervals. | **Push**: Devices publish a continuous stream of data. |
| Efficiency | Higher overhead; polling intervals can miss spikes. | **High efficiency**: Transmits only changed values (incremental). |
| Standardization | Universal (RFC 1155). | Vendor-specific (evolving standards). |

> **TIP:** For more about Kentik's perspective on Streaming Telemetry, see the blog post **Streaming Telemetry for Network Monitoring and Analytics**.

## Data Support in KDE

The Kentik Data Engine (KDE), a columnar datastore, correlates flow data (NetFlow/sFlow) with SNMP and ST metadata to provide a unified view in **Data Explorer**.

### SNMP Data

If devices are configured for **SNMP OID Polling**, Kentik stores these two types of data:

* **Counter OIDs:** Represented as **SNMP Metrics** (e.g., interface throughput).
* **Information OIDs:** Represented as **SNMP Dimensions** (e.g., device metadata).

### Streaming Telemetry Data

If devices publish ST data to Kentik, it is integrated into KDE flow records:

* **ST Metrics:** Functionality equivalent to traditional SNMP metrics; listed in **Streaming Telemetry Metrics**.
* **ST Dimensions:** Available when querying ST-derived data; listed in **Streaming Telemetry Dimensions**.

## Deployment & Configuration

Kentik supports a **Direct Push** architecture: send ST data from supported devices directly to the Kentik Data Engine (KDE).

### Supported Platforms

| Vendor | Requirement | Implementation |
| --- | --- | --- |
| Juniper Junos | Junos OS 18.4R2.7+ | Apply **Direct Dialout Snippet**. |
| Cisco IOS-XRv 9000 | Version 6.2.3+ | Apply **Direct Dialout Snippet**. |

> **IMPORTANT**:
>
> * To request ST collection for your account, contact **Customer Care**.
> * Ensure your ACLs allow outbound UDP traffic to Kentik.
> * Vendor support for ST is evolving; contact support to inquire about specific device series.
