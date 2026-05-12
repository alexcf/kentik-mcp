# Device Metrics Dimensions

Source: https://kb.kentik.com/docs/device-metrics-dimensions

---

This article discusses dimensions related to devices and interfaces based on data collected via SNMP polling or Streaming Telemetry.

> **Notes:**
>
> * Streaming Telemetry is discussed in this article in relation to Kentik's non-NMS ST implementation. For information about ST for Kentik NMS, see **NMS via Streaming Telemetry**).
> * As of May 1st, 2025, the **Query SQL Method** has been deprecated and is no longer supported.

## SNMP Dimensions

> **Note:** Queries using SNMP dimensions are subject to the following restrictions:
>
> * SNMP dimensions cannot be mixed with dimensions from any other category.
> * Results will return only if the query's metrics are from the SNMP category in the Available Metrics list (see **Metrics Dialog UI**).

If SNMP polling is enabled on a router (see **SNMP OID Polling**), Kentik is able to enrich the KDE with the SNMP-derived dimensions listed below. These dimensions are stored in Kentik using **Universal Data Records** (UDR), allowing flexible data allocation to the columns of the Kentik Data Engine (see **Information SNMP OIDs**).

> **Note:** UDR dimensions have no persistent KDE columns.

### SNMP Interface Dimensions

The dimensions listed in the table below are derived by Kentik from interfaces using the SNMP OIDs covered in **Kentik-polled SNMP OIDs**.

| Dimension name (portal) | Description | Type: value column | Direction |
| --- | --- | --- | --- |
| Interface | The Name and Description of the interface (see **Interface Field Definitions**) stated in the form of "name: description". | string UDR | Non-directional |
| Device | The Kentik-registered name for the device (see **Device General Settings**). | int UDR | Non-directional |
| Site | Name of the site to which the device has been assigned (see **About Sites**). If the device hasn't been assigned to a site, returns an empty string.  **Notes:**   * Supported operators for WHERE clause: case-insensitive equality, LIKE, IN, and regex matching. * Site assignments in the table may lag Admin settings by up to 10 minutes. | string UDR | Non-directional |
| Provider | A string representing the provider via which source/destination traffic over a given interface reaches the Internet (see **About Provider Classification**). | string UDR | Non-directional |
| Connectivity Type | The connectivity type, such as transit, IX peering, etc., of the source/destination interface of this flow (see **Connectivity Type Attribute**). | string UDR | Non-directional |
| Network Boundary | The network boundary value (internal or external) of the source/destination interface of this flow (see **Network Boundary Attribute**). | string UDR | Non-directional |
| Interface Capacity | The speed of the device interface through which flow ingressed/egressed (see **Interface Field Definitions**). | integer UDR |  |

### SNMP Device Dimensions

The dimensions listed in the table below are derived by Kentik from devices using the SNMP OIDs covered in **Kentik-polled SNMP OIDs**.

| Dimension name (portal) | Description | Type: value column | Direction |
| --- | --- | --- | --- |
| Error | An error string associated with failed polls (used for troubleshooting). | string UDR | Non-directional |
| Component | The name of the component that reports CPU or memory utilization. Examples***:*** "Routing Engine," "FPC: QFX5100-48S-6Q @ 0/\*/\*," "Processor" | int UDR | Non-directional |
| Uptime | The time (in hundredths of a second) since the network management portion of the system was last re-initialized. | string UDR | Non-directional |
| Device | The Kentik-registered name for the device (see **Device General Settings**). | string UDR | Non-directional |
| Site | Name of the site to which the device has been assigned (see **About Sites**). If the device hasn't been assigned to a site, returns an empty string.  **Notes:**   * Supported operators for WHERE clause: case-insensitive equality, LIKE, IN, and regex matching. * Site assignments in the table may lag Admin settings by up to 10 minutes. | string UDR | Non-directional |

## Streaming Telemetry Dimensions

> **Note:** Queries using Streaming Telemetry dimensions are subject to the following restrictions:
>
> * Streaming Telemetry dimensions cannot be mixed with dimensions from any other category.
> * Results will return only if the query's metrics are from the Streaming Telemetry category in the Available Metrics list (see **Metrics Dialog UI**).

If streaming telemetry publishing is enabled on a router (see **Streaming Telemetry Device Support**) then you can run queries in the portal that use streaming telemetry interface dimensions. The dimensions that are available for these ST queries have the same names as the dimensions listed in the **SNMP Dimensions** table above, but they appear in the Streaming Telemetry Interface section of the list in the portal's Group By Dimensions dialog (e.g., in Data Explorer).

> **Note:** Streaming Telemetry is discussed above in relation to Kentik's non-NMS ST implementation. For information about ST for Kentik NMS, see **NMS via Streaming Telemetry**).
