# Router Configuration

Source: https://kb.kentik.com/docs/router-configuration

---

This article discusses configuring routers, switches, and other network hardware to collect and export data to Kentik.

> **Notes:**
>
> * Model-specific configuration settings for devices sending data to Kentik are in the **Device Configs Directory**.
> * To register routers on the Kentik system, see **Device Settings**.
> * For general information about flow, see **Flow Overview**.
> * In this Knowledge Base, the term “router” also refers to other non-host network devices such as switches.

## Router Configuration Overview

The Kentik Data Engine (KDE) collects and correlates data from various sources, including routers, switches, and other hardware in your Kentik-monitored network. Data includes flow records like NetFlow, IPFIX, sFlow (see **Overview**), as well as other network data such as BGP and SNMP.  
  
To enable Kentik to gather data from a router, configure the device and the Kentik portal using the **Device Settings Dialog** (accessed via Settings » **Devices**). Before starting, decide on the data collection method:

* Direct to KDE ingest servers
* Through a local encryptor/redirector host running Kentik’s **Universal Agent** with the Flow Proxy capability enabled.

The device configuration process varies by manufacturer but typically involves “configuration mode” or a “configuration editor.” Before starting, know the following:

* **IP and port**: The destination IP and Port for flow data:

  + If sending directly to Kentik, this information varies by customer and is found in the **General** tab of the **Device** dialog in the Kentik portal (see **Device Config Info**).
  + If encrypting with Kentik’s **Universal Agent** before sending, these are the IP and Port chosen on your local encryptor/redirector running the agent.
* **Sample rate**: The desired sample rate for flow records (see **Flow Sampling**). The rate on the router should match the Kentik portal (see **Device General Settings**).
* **Ingress or egress**: Examine traffic at ingress (recommended) or egress (see **Ingress and Egress**).

Once you have the information, configure your routers for Kentik. Common networking hardware products’ configurations are in the **Device Configs Directory**.

## Router BGP Considerations

Consider the following when configuring routers for BGP peering with Kentik:

* Kentik establishes a BGP session when it receives a peering request, but session info appears in the portal only after flow data is received.
* Kentik peers as an iBGP `rr-client` (same ASN as the peering router).
* 4-byte ASN compatibility is mandatory.
* Inbound firewall policies (ACLs) must allow inbound BGP sessions from the Kentik peering IP.

> **Note:** Kentik recommends filtering (not propagating) your default route to Kentik. If present, it may override the final destination ASN assignment of all unattributed-route flow records, either with your ASN or your default transit provider’s ASN.

#### BGP Session Stability

Unlike standard peering or transit connections, which are generally point-to-point, BGP sessions with Kentik pass through multiple transit providers that are beyond our customers' direct oversight.

* If your BGP session with Kentik is interrupted:

  + Routing remains unaffected. Our BGP sessions provide only telemetry data, not traffic forwarding.
  + Flow data received by Kentik while a session is down is correlated with the last known good data from your network.
* To maximize BGP session stability, consider:

  + Using longer timers to avoid unwanted resets when the Internet is “stormy”.
  + Contacting Kentik support to establishing a Private Network Interconnect (PNI) for bypassing the Internet.

## Router Troubleshooting

If you've configured a router to send flow to Kentik (using the router-specific configurations listed in the **Device Configs Directory**) and don’t see the flow in the Kentik portal, we need to know if the router can reliably ping our collectors with large packets. To test this, perform the following simple tests:

* Determine that there's no loss between your server and Kentik:

```
ping -c200 -D -s400 flow.kentik.com
```

* Determine if the MTU between you and Kentik is "normal":

```
ping -c100 -D -s1472 flow.kentik.com
```

* Determine if the MTU between you and Kentik is "normal":

```
ping -c100 -s1500 flow.kentik.com
```

> **Note:** If your organization is registered with Kentik in the EU, use `flow.kentik.eu` instead of the above URLs.

The information you gather from these tests help us troubleshoot the issue if you contact **support@kentik.com**.

## SNMP OID Polling

OIDs are identifiers for SNMP objects representing network-connected devices like routers. Each OID is a path to the object it represents, similar to a standard HTTP path. Each segment represents a narrower slice of the networked universe, but each segment is a pre-assigned number for an OID. The base OID for MIB-2 defined SNMP variables is `1.3.6.1.2.1.`

Kentik polls SNMP OIDs in different categories:

* Selected **Counter SNMP OIDs**
* Selected **Information SNMP OIDs**
* Selected **System Information OIDs**
* Selected **LLDP OIDs**
* **Vendor-specific SNMP OIDs**

> **Notes:**
>
> * By default, SNMP is polled only when Kentik receives flow from a device. To poll a device without flow, use the SNMP/ST capability of Kentik’s **Universal Agent**, which enables full device monitoring independently of flow data.
> * The polling timeout is 60 seconds. If a response is not received, polling is skipped until the next interval (see **SNMP Polling Intervals**).

### SNMP Polling Intervals

The polling intervals for a device depend on its SNMP Polling setting, which is set on the SNMP tab in the Add Device or Edit Device dialog (see **Device SNMP Settings**). The following options are supported:

* **Standard**: Polls interface counters and device counters every 5 minutes and interface descriptions every 3 hours.
* **Minimum**: Polls interface descriptions every 6 hours. Interface counters and device counters are not polled.

Choosing Standard polling enables features of Kentik that depend on device/interface metrics:

* **Health status**:

  + Display of health status, including interface utilization, in Kentik Map (see **Kentik Map Health**).
  + Querying based on health in Data Explorer (see **SNMP Interface Metrics**).
* **Connectivity Costs**: Calculation of costs based on SNMP-collected traffic volume data (see **Connectivity Traffic Calculation**).

> **Notes:**
>
> * The SNMP Polling setting doesn’t affect Kentik's polling of System Information OIDs.
> * The Interface List (see **Interfaces Page**) includes indicators to compare SNMP-reported flow volume with flow records from the same device.

### Enabling SNMP Polling

To enable Kentik to poll SNMP on a router:

1. Determine the SNMP version (see **About SNMP V3**).
2. Ensure SNMP is enabled for the router (consult documentation for your router make and model).
3. Permit SNMP polling of the router from Kentik's SNMP polling IPs. The IPs are listed at the top of the **SNMP** tab of the router's **Edit Device** dialog in the portal (accessed by clicking the Edit icon for that router in the **Device List**).
4. Set the router’s community to match the **SNMP Community** string indicated on the router's **SNMP** tab.
5. If the router blocks polling of specific OIDs (see **Kentik-polled SNMP OIDs**), re-enable polling of those OIDs.

#### About SNMP V3

Kentik supports polling via SNMP V3, which is more secure than previous versions. It’s recommended for customers concerned about using SNMP V2 over the public Internet.

Kentik’s SNMP V3 implementation allows independent configuration of:

* **Authentication**: Options include:

  + None
  + MD5
  + SHA
* **Privacy**: The actual encryption of SNMP transactions:

  + None
  + 56-bit DES encryption
  + AES-128

> **Note:** Kentik's SNMP V3 implementation does not support 168-bit 3DES.

#### Using SNMP V3

To use SNMP V3:

1. Configure your router to enable polling via SNMP V3. Consult your router documentation for the correct settings.
2. Enable SNMP V3 using the **SNMP V3 Auth** toggle switch in the **Add Device** or **Edit Device** dialog in the Kentik portal.
3. Fill in the additional fields (see **Edit a Device** and **Device SNMP Settings**).

### Verifying SNMP Polling

After completing the steps in **Enabling SNMP Polling**, verify in the portal that Kentik can poll your router after about 5 minutes (one polling interval):

1. Go to the portal's Settings page and choose **Networking Devices**.
2. Find the router in the **Device** list and confirm the green SNMP indicator in the **Status** column.
3. Click anywhere in the router's row to bring up the **Device Details** pane.
4. Click **View Interfaces** (under Interfaces) to open the router’s **Interfaces** page.
5. Verify that the names and descriptions for the interfaces appear on the Interfaces page.
6. In the **Traffic In** and **Traffic Out** columns, verify that the SNMP value is greater than zero.

> **Note**: If columns are not visible, click the **Customize** button at the top right, and check the boxes for any missing columns (see **Customize Columns Popup**).

### Kentik-polled SNMP OIDs

Kentik polls the following OIDs. To enable Kentik to poll SNMP on a given device, the device must not be configured to block polling of any of the listed OIDs.

> **Notes:**
>
> * Discontinuities in counter values can occur at re-initialization of the management system or at other times indicated by the value of the OID `ifCounterDiscontinuityTime` `(1.3.6.1.2.1.31.1.1.1.19)`.
> * Additional information about the OIDs below can be found in the Global OID reference database at **https://oidref.com**.

#### Counter SNMP OIDs

These counter OIDs are polled every 5 minutes when SNMP polling is standard (see **SNMP Polling Intervals**), and are not polled when polling is minimized:

| OID | Object/variable name  (SNMP\_...) | Portal metric | Streaming Telemetry path | Description |
| --- | --- | --- | --- | --- |
| 1.3.6.1.2.1.1.3.0 | `sysUpTime` | Uptime  (dimension) | N.A. | The time (in hundredths of a second) since the network management portion of the system was last re-initialized. Polled for every device metric collection and stored as Uptime dimension. |
| 1.3.6.1.2.1.31.1.1.1.6 | `ifHCInOctets` | Input Bit Rate | in-octets | The total number of octets received on the interface, including framing characters. |
| 1.3.6.1.2.1.31.1.1.1.10 | `ifHCOutOctets` | Output Bit Rate | out-octets | The total number of octets transmitted out of the interface, including framing characters. |
| 1.3.6.1.2.1.31.1.1.1.7 | `ifHCInUcastPkts` | Input Packets | in-unicast-pkts | The number of packets, delivered by this sub-layer to a higher sub-layer, which were not addressed to a multicast or broadcast address at this sub-layer. |
| 1.3.6.1.2.1.31.1.1.1.11 | `ifHCOutUcastPkts` | Output Packets | out-unicast-pkts | The total number of packets that higher-level protocols requested be transmitted, and which were not addressed to a multicast or broadcast address at this sub-layer, including those that were discarded or not sent. |
| 1.3.6.1.2.1.2.2.1.14 | `ifInErrors` | Input Errors | in-errors | The number of inbound packets that contained errors preventing them from being deliverable to a higher-layer protocol. |
| 1.3.6.1.2.1.2.2.1.20 | `ifOutErrors` | Output Errors | out-errors | The number of outbound packets that could not be transmitted because of errors. |
| 1.3.6.1.2.1.2.2.1.13 | `ifInDiscards` | Input Discards | in-discards | The number of inbound packets which were chosen to be discarded even though no errors had been detected, possibly to free up buffer space. |
| 1.3.6.1.2.1.2.2.1.19 | `ifOutDiscards` | Output Discards | out-discards | The number of outbound packets which were chosen to be discarded even though no errors had been detected, possibly to free up buffer space. |
| 1.3.6.1.2.1.31.1.1.1.8 | `ifHCInMulticastPkts` | Input Multicast Packets | in-multicast-pkts | The number of packets, delivered by this sub-layer to a higher sub-layer, which were addressed to a multicast address at this sub-layer. For a MAC layer protocol, this includes both Group and Functional addresses. |
| 1.3.6.1.2.1.31.1.1.1.12 | `ifHCOutMulticastPkts` | Output Multicast Packets | out-multicast-pkts | The total number of packets that higher-level protocols requested be transmitted, and which were addressed to a multicast address at this sub-layer, including those that were discarded or not sent. For a MAC layer protocol, this includes both Group and Functional addresses. |
| 1.3.6.1.2.1.31.1.1.1.9 | `ifHCInBroadcastPkts` | Input Broadcast Packets | in-broadcast-pkts | The number of packets, delivered by this sub-layer to a higher sub-layer, which were addressed to a broadcast address at this sub-layer. |
| 1.3.6.1.2.1.31.1.1.1.13 | `ifHCOutBroadcastPkts` | Output Broadcast Packets | out-broadcast-pkts | The total number of packets that higher-level protocols requested be transmitted, and which were addressed to a broadcast address at this sub-layer, including those that were discarded or not sent. |

#### Information SNMP OIDs

The information OIDs listed below are polled every 3 hours when SNMP polling is standard (see **SNMP Polling Intervals**), and every 6 hours when polling is minimized.

| OID | Object/variable name  (SNMP\_...) | Portal dimension (filtering) | Description |
| --- | --- | --- | --- |
| 1.3.6.1.2.1.10.166.11.1.2.2.1.3 | `mplsL3VpnVrfDescription` | VRF Name | The human-readable description of this VRF. Default is "" (empty string). |
| 1.3.6.1.2.1.10.166.11.1.2.2.1.4 | `mplsL3VpnVrfRD` | VRF Route Distinguisher | The route distinguisher for this VRF. Default is "" (empty string). |
| 1.3.6.1.2.1.10.166.11.1.2.3.1.4 | `mplsL3VpnVrfRT` | VRF Route Target | The route target distribution policy. Default is "" (empty string). |
| 1.3.6.1.2.1.10.166.11.1.2.1.1.2 | `mplsL3VpnIfVpnClassification` | N.A. (Kentik internal use) | Denotes whether this link participates in a carrier's carrier, enterprise, or inter-provider scenario. Default is "enterprise." |
| 1.3.6.1.2.1.2.2.1.2 | `ifDescr` | Interface Name | A textual string containing information about the interface. Includes manufacturer name, product name, and interface version. |
| 1.3.6.1.2.1.31.1.1.1.18 | `ifAlias` | Interface Name | An 'alias' name for the interface, as specified by a network manager, that provides a non-volatile 'handle' for the interface. |
| 1.3.6.1.2.1.31.1.1.1.15 | `ifHighSpeed` | Interface Capacity | An estimate of the interface's current bandwidth in bits per second. |
| 1.3.6.1.2.1.2.2.1.3 | `ifType` | N.A. | IANA ifType |
| 1.3.6.1.2.1.2.2.1.4 | `ifMtu` | N.A. | Interface MTU size |
| 1.3.6.1.2.1.2.2.1.6 | `ifPhysAddress` | N.A. | Interface Physical (MAC) address |
| 1.3.6.1.2.1.2.2.1.7 | `ifAdminStatus` | N.A. | The configured status of the interface. |
| 1.3.6.1.2.1.2.2.1.8 | `ifOperStatus` | N.A. | The interface operational status. |
| 1.3.6.1.2.1.2.2.1.9 | `ifLastChange` | N.A. | The value of `sysUpTime` at the time the interface entered its current operational state. If the current state was entered prior to the last re-initialization of the local network management subsystem, then this object contains a zero value. |
| 1.3.6.1.2.1.31.1.2.1.3 | `ifStackStatus` | N.A. | The status of the relationship between two sub-layers. |
| 1.3.6.1.2.1.31.1.1.1.17 | `ifConnectorPresent` | N.A. | This object has the value 'true(1)' if the interface sublayer has a physical connector and the value 'false(2)' otherwise. |
| 1.3.6.1.2.1.4.20.1.2 | `ipAdEntIfIndex` | N.A. | An index value that uniquely identifies an interface. Used to derive the IP displayed for the interface in the portal (interface-related pages and dialogs). |
| 1.3.6.1.2.1.4.20.1.3 | `ipAdEntNetMask` | N.A. | The subnet mask associated with the IP address of this entry. Used to derive the IP mask displayed for the interface in the portal (interface-related pages and dialogs). |
| 1.3.6.1.2.1.55.1.8.1.2 | `ipv6AddrPfxLength` | N.A. (Kentik internal use) | The length of the prefix (in bits) associated with the IPv6 address of this entry. |

#### System Information OIDs

The system information OIDs listed below are polled by Kentik only once per client restart and are not affected by the SNMP polling interval setting.

| OID | Object/variable name (SNMP\_...) | Description |
| --- | --- | --- |
| 1.3.6.1.2.1.1.1 | `sysDescr` | A textual description of the entity. Includes the full name and version identification of the system's hardware type, software operating-system, and networking software. |
| 1.3.6.1.2.1.1.2 | `sysObjectID` | The vendor's authoritative identification of the network management subsystem contained in the entity.  **Note**: Used by Kentik to determine the vendor. |
| 1.3.6.1.2.1.1.4 | `sysContact` | The textual identification of the contact person for this managed node, together with information on how to contact this person. If no contact information is known, the value is the zero-length string. |
| 1.3.6.1.2.1.1.5 | `sysName` | An administratively-assigned name for this managed node. By convention, this is the node's fully-qualified domain name. If the name is unknown, the value is a zero-length string. |
| 1.3.6.1.2.1.1.6 | `sysLocation` | The physical location of this node. If the location is unknown, the value is a zero-length string. |
| 1.3.6.1.2.1.1.7 | `sysServices` | A value indicating the set of services potentially offered by this entity. |

#### LLDP OIDs

The LLDP OIDs listed below are polled by Kentik every three hours. Kentik is able to create an L2 topology map based on the LLDP information.

| OID | Object/variable name (SNMP\_...) | Description |
| --- | --- | --- |
| 1.0.8802.1.1.2.1.4.1.1.6 | `lldpRemPortIdSubtype` | The type of port identifier encoding used in the associated `lldpRemPortId` object. |
| 1.0.8802.1.1.2.1.4.1.1.7 | `lldpRemPortId` | The string value used to identify the port component associated with the remote system. |
| 1.0.8802.1.1.2.1.4.1.1.8 | `lldpRemPortIdDesc` | The string value used to identify the description of the given port associated with the remote system. |
| 1.0.8802.1.1.2.1.4.1.1.9 | `lldpRemSysName` | The string value used to identify the system name of the remote system. |

### Vendor-specific SNMP OIDs

Each OID listed below is polled only for devices of a specific vendor, which is determined based on the `sysDescr` OID (see **System Information OIDs**). Data gathered from polling these OIDs is used to enable metrics in the category**SNMP Device Metrics**.

> **Notes:**
>
> * When you configure your device to enable polling of a given OID, that permission includes polling of all children of that OID.
> * While these OIDs contain both counter data (metrics) and non-numeric information, for the purpose of setting the polling interval they are treated as counters (see SNMP Polling Intervals).
> * Additional information about the OIDs below can be found in the Global OID reference database at **http://oidref.com**.

| OID | Object/ variable name | Kentik usage | Description |
| --- | --- | --- | --- |
| **Arista and Palo Alto Networks OIDs** | | | |
| 1.3.6.1.2.1.25.2.3 | `hrStorageTable` | Metric: Memory Utilization | The (conceptual) table of logical storage areas on the host. Used for Memory Utilization metric and Component dimension. |
| 1.3.6.1.2.1.25.3.2 | `hrDeviceTable` | Metadata: Dimension Component | The (conceptual) table of devices contained by the host. Used for Component dimension of the CPU metric. |
| 1.3.6.1.2.1.25.3.3 | `hrProcessorTable` | Metric: CPU | The (conceptual) table of processors contained by the host. Used for CPU Utilization metric. |
| **Cisco OIDs** | | | |
| 1.3.6.1.2.1.47.1.1.1.1.7 | `entPhysicalName` | Metadata: Dimension Component | The textual name of the physical entity. |
| 1.3.6.1.4.1.9.9.109.1.1.1 | `cpmCPUTotalTable` | Metric: CPU | A table of overall CPU statistics. |
| 1.3.6.1.4.1.9.9.48.1.1 | `ciscoMemoryPoolTable` | Metric: Memory Utilization | An entry in the memory pool monitoring table. |
| **F5 BIG-IP OIDs** | | | |
| 1.3.6.1.4.1.3375.2.1.1.2.20.2 | `sysGlobalHostMemTotal` | Component = Global | The total host memory in bytes for the system. |
| 1.3.6.1.4.1.3375.2.1.1.2.20.3 | `sysGlobalHostMemUsed` | Component = Global | The host memory in bytes currently in use for the system. |
| 1.3.6.1.4.1.3375.2.1.1.2.20.37 | `sysGlobalHostCpuUsageRatio5m` | Component = Global | The average usage ratio of CPU for the system in the last five minutes. |
| 1.3.6.1.4.1.3375.2.1.1.2.1.44 | `sysStatMemoryTotal(44)` | Component = TMM | The total memory available in bytes for TMM (Traffic Management Module). |
| 1.3.6.1.4.1.3375.2.1.1.2.1.45 | `sysStatMemoryUsed` | Component = TMM | The memory in use in bytes for TMM (Traffic Management Module). |
| 1.3.6.1.4.1.3375.2.1.1.2.20.44 | `sysGlobalHostOtherMemoryTotal` | Component = Other | The total other non-TMM memory in bytes for the system. |
| 1.3.6.1.4.1.3375.2.1.1.2.20.45 | `sysGlobalHostOtherMemoryUsed` | Component = Other | The other non-TMM memory in bytes currently in use for the system. |
| 1.3.6.1.4.1.3375.2.1.1.2.20.46 | `sysGlobalHostSwapTotal` | Component = Swap | The total swap in bytes for the system. |
| 1.3.6.1.4.1.3375.2.1.1.2.20.47 | `sysGlobalHostSwapUsed` | Component = Swap | The swap in bytes currently in use for the system. |
| **Huawei OIDs** | | | |
| 1.3.6.1.4.1.2011.5.25.110.1.2.1.2 | `hwifNet32BitIndex` | Interface index mapping | NetStream32BitIndex indicates the interface index which is distributed by VRP. |
| 1.3.6.1.4.1.2011.5.25.31.1.1.1.1.5 | `hwEntityCpuUsage` | Metric: CPU | The overall CPU usage of an entity, generally calculated without considering the number of CPUs. |
| 1.3.6.1.4.1.2011.5.25.31.1.1.1.1.7 | `hwEntityMemUsage` | Metric: Memory Utilization | The memory usage of an entity, expressed as the percentage of the memory that has been used. |
| 1.3.6.1.4.1.2011.5.25.31.1.1.1.1.10 | `hwEntityUpTime` | Dimension: Uptime | The total duration when an entity is in the UP state. |
| **Juniper OIDs** | | | |
| 1.3.6.1.4.1.2636.3.1.13 | `jnxOperatingTable` | Dimensions and Metrics | A list of operating status entries. Used for CPU and Memory Utilization metrics and Component dimension. |
| **Nokia OIDs** | | | |
| 1.3.6.1.4.1.6527.3.1.2.3.4.1.4 | `vRtrIfName` | Metadata: VRF | The administrative name assigned this router interface. The interface name must be unique among entries with the same `vRtrID` value. In order for row creation to succeed, a value must also be assigned to `vRtrIfName`. |
| 1.3.6.1.4.1.6527.3.1.2.3.4.1.34 | `vRtrIfDescription` | Metadata: VRF | A user provided description string for this virtual router interface |
| 1.3.6.1.4.1.6527.3.1.2.3.4.1.63 | `vRtrIfGlobalIndex` | Metadata: VRF | Uniquely identifies this interface in the tmnx system. This field provides an identifier for router interfaces similar to the `vRtrIfIndex` value, except that `vRtrIfIndex` is unique within each virtual router. The `vRtrIfGlobalIndex` is unique system wide regardless of the `vRtrID`. |

> **Note:** Kentik will ignore any components which have CPU and Memory utilization of 0.
