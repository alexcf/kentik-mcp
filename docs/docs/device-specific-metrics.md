# Device-Specific Metrics

Source: https://kb.kentik.com/docs/device-specific-metrics

---

This article discusses device-specific metrics in the Kentik Data Engine (KDE).

> **Notes:**
>
> * For general information about metrics in Kentik, see **About Metrics**.
> * Filter and group-by dimensions are covered in **Dimensions Reference** (see **Per-flow Metrics**).
> * Metrics are used for Kentik portal query settings and the **Query API**.

## About Device-Specific Metrics

While Kentik ingests standard flow data (NetFlow, sFlow, IPFIX) into a common schema, many networking devices export unique performance metadata that doesn't fit into traditional fields. To capture this specialized information, Kentik utilizes **Universal Data Records**(UDR).

### Why Use Device-Specific Metrics?

Standard metrics provide a broad overview of network traffic, but device-specific metrics offer deep-tier visibility into vendor-specific features:

* **Bidirectional Correlation:** Firewalls like Cisco ASA or Meraki often export "Initiator" and "Responder" data in a single record. Device-specific metrics allow you to analyze both directions of a session simultaneously.
* **Performance Insight:** Vendors like Silver Peak export application-layer metrics, such as network-to-server or network-to-client delay, which are not available in standard NetFlow.
* **Flexible Data Mapping:** UDR allows Kentik to map these unique vendor elements to flexible columns in the Kentik Data Engine (KDE) without requiring a static schema change for every new device type.

> **Note:** Unlike standard metrics, UDR-based metrics do not have persistent KDE columns. They are dynamically allocated based on the device type and the specific templates being exported.

## Cisco ASA Metrics

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/logo-cisco-063h120w(1).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A17Z&se=2026-05-12T09%3A29%3A17Z&sr=c&sp=r&sig=5oFxeXO%2Fdb8YTYiHnpHWMGraDmrbgy9uySRjPSmZ2sE%3D)These metrics represent traffic volume of flow through Cisco Adaptive Security Appliances (ASA), which can be standalone appliances, blades, or virtual appliances. ASA uses bidirectional flow records in which the "initiator" (source) initiating a request and the "responder" (destination) replying.

> **Notes:**
>
> * These KDE flow fields store a sum for the flow, used to derive the Average, 95th Percentile, and Max numbers returned from queries.
> * For more context on these dimensions, see the Cisco document **ASA NetFlow Implementation Guide**.

| Metric Name (Portal) | Description | Type: value column |
| --- | --- | --- |
| Initiator Bytes | The bytes going from the initiator to the responder. | integer UDR |
| Responder Bytes | The bytes going from the responder to the initiator. | integer UDR |

## Cisco Meraki Metrics

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/logo-Meraki-045h155w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A17Z&se=2026-05-12T09%3A29%3A17Z&sr=c&sp=r&sig=5oFxeXO%2Fdb8YTYiHnpHWMGraDmrbgy9uySRjPSmZ2sE%3D)Like Cisco ASA, Cisco Meraki uses bidirectional flow records. Kentik's Meraki-specific UDR fields, listed below, store the volume of the flow's "responder" (destination) traffic, whose direction is from the Meraki firewall back to the "initiator" (source) of a request. These fields correspond to the `OUT_BYTES` and `OUT_PKTS` data fields in Meraki's **NetFlow Version 9 Templates**.

> **Notes:**
>
> * These KDE flow fields store a sum for the flow, which is used to derive the Average, 95th Percentile, and Max numbers that return from queries (e.g., the columns of the tables returned in Data Explorer).
> * The volume from the initiator to the responder is stored in the standard KDE metrics `fields in_bytes` and `in_packets` (see **General Metrics**).

| Metric Name (Portal) | Description | Type: value column |
| --- | --- | --- |
| Out Bytes | Number of bytes leaving the firewall for this flow. | integer UDR |
| Out Packets | Number of packets leaving the firewall for this flow. | integer UDR |

## Cisco Zone-based Firewall

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/logo-cisco-063h120w(1).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A17Z&se=2026-05-12T09%3A29%3A17Z&sr=c&sp=r&sig=5oFxeXO%2Fdb8YTYiHnpHWMGraDmrbgy9uySRjPSmZ2sE%3D)These metrics represent traffic volume of flow through Zone-based Firewalls (ZFW) on Cisco IOS devices. ZFW uses bidirectional flow records in which the "initiator" (source) is the entity that initiates a request and the "responder" (destination) is the entity that replies with a response.

> **Notes:**
>
> * These KDE flow fields store a sum for the flow, which is used to derive the Average, 95th Percentile, and Max numbers that return from queries (e.g., the columns of the tables returned in Data Explorer).
> * For more context on these dimensions, see the Cisco document **Zone-based Policy Firewall Guide**.

| Metric Name (Portal) | Description | Type: value column |
| --- | --- | --- |
| Initiator Bytes | The bytes going from the initiator to the responder. | integer UDR |
| Responder Bytes | The bytes going from the responder to the initiator. | text UDR |

## Silver Peak Metrics

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/logo-silver_peak-045h196w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A17Z&se=2026-05-12T09%3A29%3A17Z&sr=c&sp=r&sig=5oFxeXO%2Fdb8YTYiHnpHWMGraDmrbgy9uySRjPSmZ2sE%3D)These metrics represent App Performance Information elements available from Silver Peak appliances running VXOA software (version 8.21 or higher), which is described in this **Silver Peak white paper****.**

| Metric Name (Portal) | Description | Type: value column |
| --- | --- | --- |
| Network To Server Delay | The delay, in microseconds, from the network to the server. | integer UDR |
| Network To Client Delay | The delay, in microseconds, from the network to the client. | integer UDR |
| Client To Server Response Delay | The total delay, in microseconds, from client to server. | integer UDR |
