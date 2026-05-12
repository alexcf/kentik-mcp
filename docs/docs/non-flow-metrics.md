# Non-Flow Metrics

Source: https://kb.kentik.com/docs/non-flow-metrics

---

This article discusses **non-flow metrics** used in Kentik queries.

Non-flow metrics allow you to monitor health indicators, such as CPU, Memory, and Interface utilization, that traditional flow records do not capture. These metrics are primarily derived from **SNMP polling** and **Streaming Telemetry (ST)**.

> **Notes**:
>
> * For general information about how Kentik handles data, see **About Metrics**.
> * For metrics specific to individual hardware vendors, see **Device-Specific Metrics**.

## SNMP Metrics

When **SNMP OID Polling** is enabled, Kentik enriches the Kentik Data Engine (KDE) with health and performance data. These metrics use **Universal Data Records**(UDR), providing a flexible schema that adapts to different device types.

### Query Restrictions

To ensure data accuracy, SNMP metrics are subject to the following logic in the **Dimension Selector**:

* **No Mixing:** SNMP metrics cannot be mixed with flow-based metrics (e.g., Bits/s) in the same query.
* **Dimension Matching:** Results only return if the query’s dimensions (Group By/Filter) are also from the SNMP category.

## SNMP Interface Metrics

These metrics are collected from network interfaces using standard **Kentik-Polled SNMP OIDs**. They are typically used to monitor port saturation and physical link health.

| Metric in Portal | Variations | OID | Description |
| --- | --- | --- | --- |
| Bit Rate | Input, Output | `ifHCInOctets`,`ifHCOutOctets` | Rate of bits received/transmitted, including framing. |
| Packets | Input, Output | `ifHCInUcastPkts`, `ifHCOutUcastPkts` | Rate of unicast packets (excludes discards). |
| Errors | Input, Output | `ifInErrors`, `ifOutErrors` | Inbound/outbound packets containing delivery-preventing errors. |
| Discards | Input, Output | `ifInDiscards`, `ifOutDiscards` | Packets dropped by the device (e.g., due to buffer exhaustion). |
| Multicast | Input, Output | `ifHCInMulticastPkts`, `ifHCOutMulticastPkts` | Rate of packets addressed to a multicast group. |
| Broadcast | Input, Output | `ifHCInBroadcastPkts`, `ifHCOutBroadcastPkts` | Rate of packets addressed to a broadcast address. |

> **TIP:** All SNMP metrics above can be aggregated in the Portal as **Average**, **95th Percentile**, **Max**, or **Total**.

## SNMP Device Metrics

Device metrics track the internal health of a chassis or its individual components (e.g., line cards). These are identified in the portal using the Component dimension.

| Metric in Portal | Description |
| --- | --- |
| CPU | The percentage of CPU utilization for the component. |
| Memory Total | The total physical memory (in bytes) on the component. |
| Memory Used | The amount of memory (in bytes) currently in use. |
| Memory Free | The amount of memory (in bytes) currently available. |
| Memory Utilization | The percentage of total memory used by the component. |

For a list of supported OIDs by manufacturer, see **Vendor-Specific SNMP OIDs**.

## Streaming Telemetry Metrics

If **Streaming Telemetry** is enabled, you can query interface metrics that are pushed by the device in real-time rather than polled via SNMP.

> **IMPORTANT**: **Which ST are you using?**
>
> * **NMS Implementation**: If you are using Kentik NMS, refer to **NMS via Streaming Telemetry**.
> * **Legacy Implementation**: This section describes the non-NMS Streaming Telemetry workflow used for classic KDE queries.

### Usage in Portal

Streaming Telemetry metrics share the same names and calculation methods as the SNMP Interface Metrics listed above. However, when building a query (e.g., in **Data Explorer**), you must select them from the Streaming Telemetry category in the Metrics dialog.

> **Note**: Just like SNMP, ST metrics cannot be mixed with metrics from other categories in a single query.
