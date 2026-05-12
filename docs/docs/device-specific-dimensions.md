# Device-Specific Dimensions

Source: https://kb.kentik.com/docs/device-specific-dimensions

---

This article covers Kentik’s **Data Explorer** dimensions that are unique to a specific data source (e.g., router model).

> **IMPORTANT:** As of May 1st, 2025, the **Query SQL Method** has been deprecated and is no longer supported.

## About Device-Specific Dimensions

Device-specific dimensions in Kentik’s **Data Explorer** originate as flow records that are specific to given types of devices, whether physical or virtual, such as Kubernetes containers, Palo Alto Networks firewalls, or Cisco ASA appliances. These records are ingested into Kentik as **Universal Data Records** (UDR), allowing flexible allocation of flow fields to the columns of the Kentik Data Engine. The resulting dimensions are used for filter or group-by like any other fields in Kentik-ingested flow records.

> **Notes:**
>
> * When a device's type is specified in Kentik as "MPLS Router" (e.g., with the Type setting on the General tab in the **Device General Settings**), data collection will be enabled for **MPLS Dimensions** but not for the device's device-specific dimensions.
> * UDR dimensions have no persistent KDE columns.
> * Kentik also stores and uses certain **Device-Specific Metrics**.

## A10 Thunder CGN Dimensions

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/logo-A10-045h126w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A00Z&se=2026-05-12T09%3A45%3A00Z&sr=c&sp=r&sig=vsuQg87%2FsR%2FhXPywVi1HrcSQi8EgmVT1z%2Bo4jyjKkTM%3D)These dimensions support the NAT NetFlow template that is provided by **A10 Thunder Carrier Grade Networking** devices. For more information on CGN, see **What is Carrier Grade NAT?**

| Dimension Name (Portal) | Description | Type: value column | Direction |
| --- | --- | --- | --- |
| Post-NAT Transport Port | The source/destination port identifier in the transport header, as modified by the firewall during network address port translation after the packet traversed the interface. | integer Virtual | Src/Dst |
| Post-NAT Address | The IPv4 source/destination address in the IP packet header, as modified by the firewall during network address translation after the packet traversed the interface. | string Virtual | Src/Dst |
| NAT Event | A NAT Event Type as defined in the IANA registry (see **IPFIX NAT event types**), such as NAT translation create, NAT translation delete, Threshold Reached, or Threshold Exceeded. | string Native | Non-directional |

## Cisco ASA Dimensions

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/logo-cisco-063h120w(2).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A00Z&se=2026-05-12T09%3A45%3A00Z&sr=c&sp=r&sig=vsuQg87%2FsR%2FhXPywVi1HrcSQi8EgmVT1z%2Bo4jyjKkTM%3D)These dimensions are used to filter or group-by on fields in flow records from Cisco Adaptive Security Appliances (ASA), which run Cisco ASA software to deliver enterprise-class firewall capabilities in a variety of form factors including standalone appliances, blades, and virtual appliances. For more context on these dimensions, see the Cisco document **ASA NetFlow Implementation Guide**.

> **Note:** Syslog from Cisco ASA is ingested into KDE via the Flow Proxy capability of Kentik's **Universal Agent**. For further information, contact Kentik (see **Customer Care**).

| Dimension Name (Portal) | Description | Type: value column | Direction |
| --- | --- | --- | --- |
| Post-NAT Transport Port | The source/destination port identifier in the transport header, as modified by the firewall during network address port translation after the packet traversed the interface. | integer Virtual | Src/Dst |
| Post-NAT Address | The IPv4 source/destination address in the IP packet header, as modified by the firewall during network address translation after the packet traversed the interface. | string Virtual | Src/Dst |
| Flow ID | An identifier of a flow that is unique within an observation domain. You can use this information element to distinguish between different flows if flow keys such as IP addresses and port numbers are not reported or are reported in separate records. The `flowID` corresponds to the session ID field in Traffic and Threat logs. | integer Virtual | Non-directional |
| Firewall Event | Indicates a firewall event:   * 0 = Ignore (invalid)—Not used. * 1 = Flow created—The NetFlow data record is for a new flow. * 2 = Flow deleted—The NetFlow data record is for the end of a flow. * 3 = Flow denied—The NetFlow data record indicates a flow that firewall policy denied. * 4 = Flow alert—Not used. * 5 = Flow update—The NetFlow data record is sent for a long-lasting flow, which is a flow that lasts longer than the Active Timeout period configured in the NetFlow server profile. | integer Virtual | Non-directional |
| Extended Event Code | Provides additional information about an event:   * 1001 = the flow was denied by an ingress ACL. * 1002 = the flow was denied by an egress ACL. * 1003 = the flow was denied because connection to ASA interface was denied, an ICMP packet (v4 or v6) was denied, or for an unspecified reason. * 1004 = the flow denied because the first packet on the TCP was not a TCP SYN packet. * 2001 or greater = the flow was terminated. | integer Virtual | Non-directional |
| AAA Username | The username associated with the ASA instance that generated the flow. | string Virtual | Non-directional |
| Ingress ACL | The ID of the ACL that was applied on the input interface and either permitted or denied the flow. | string Virtual | Non-directional |
| Egress ACL | The ID of the ACL that was applied on the output interface and either permitted or denied the flow. | string Virtual | Non-directional |

> **Note:** See also **Cisco ASA Metrics**.

## Cisco ASA Syslog Dimensions

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/logo-cisco-063h120w(2).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A00Z&se=2026-05-12T09%3A45%3A00Z&sr=c&sp=r&sig=vsuQg87%2FsR%2FhXPywVi1HrcSQi8EgmVT1z%2Bo4jyjKkTM%3D)These dimensions are used to filter or group-by on KDE fields whose values are extracted at ingest from syslog messages generated by Cisco Adaptive Security Appliances (ASA); see **About ASA Syslog Messages**. Syslog data may provide additional details that supplement the data available in Cisco ASA NetFlow (see **Cisco ASA Dimensions**).

> **Note:** Syslog from Cisco ASA is ingested into KDE via the Flow Proxy capability of Kentik's **Universal Agent**. For further information, contact Kentik (see **Customer Care**).

| Dimension Name (Portal) | Description | Type: value column | Direction |
| --- | --- | --- | --- |
| Flow ID | An identifier of a flow that is unique within an observation domain. You can use this information element to distinguish between different flows if flow keys such as IP addresses and port numbers are not reported or are reported in separate records. Flow ID corresponds to the session ID field in Traffic and Threat logs. | integer Native | Non-directional |
| Message | A Cisco ASA Series syslog message. Messages are listed by message ID in **Cisco ASA Series Syslog Messages**. | string Native | Non-directional |
| Severity | The severity level of the message, which varies depending on the cause (see **Messages Listed by Severity Level**). | integer Native | Non-directional |
| Message ID | The Cisco-assigned ID for the message. | integer Native | Non-directional |

## Cisco IOS XE SD-WAN Dimensions

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/logo-cisco-063h120w(2).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A00Z&se=2026-05-12T09%3A45%3A00Z&sr=c&sp=r&sig=vsuQg87%2FsR%2FhXPywVi1HrcSQi8EgmVT1z%2Bo4jyjKkTM%3D)Kentik supports the added SD-WAN Edge functionality in Cisco’s IOS XE devices via the dimensions described in the below table.

> **TIP**: For more information about these devices, refer to the Cisco document **Cisco Catalyst SD-WAN Policies Configuration Guide, Cisco IOS XE Catalyst SD-WAN Release 17.x**.

| Dimension Name (Portal) | Description | Type: value column | Direction |
| --- | --- | --- | --- |
| Application | Application name, from the field 96. | string Native | Non-directional |
| Application Category | Application category, from proprietary field 12232. | string Native | Non-directional |
| Application Subcategory | Application subcategory, from proprietary field 12233. | string Native | Non-directional |
| Application Group | Application group, from proprietary field 12234. | string Native | Non-directional |
| Application Traffic Class | Application traffic class, from proprietary field 12243. | string Native | Non-directional |
| Application Business Relevance | Application business relevance, from proprietary field 12244. | string Native | Non-directional |
| Flow end reason | This information is extracted from IPFIX field 136. Possible values can be found in the IANA specification: **IPFIX Flow End Reason**. | integer Native | Non-directional |
| VPN identifier | Cisco IOS XE Catalyst SD-WAN Device VPN identifier. The device uses the enterprise ID for VIP\_IANA\_ENUM or 41916, and the VPN element ID is 4321. | integer Native | Non-directional |
| Overlay Session ID | A 32-bit identifier for input/output overlay session identifier. From proprietary fields 12432 and 12433. | integer Native | Src/Dst |
| Routing VRF Service | Routing VRF Service, from proprietary field 12434. | integer Native | Non-directional |
| Connection ID | A 64-bit identifier for a connection between client and server. From proprietary field 12441. | integer Native | Non-directional |
| BFD avg latency | Calculation of the Bidirectional Forwarding Detection (BFD) average latency for each tunnel. From proprietary field 45296. | integer Native | Non-directional |
| BFD avg loss | Calculation of the BFD average loss for each tunnel. From proprietary field 45295. | integer Native | Non-directional |
| BFD avg jitter | Calculation of the BFD average jitter for each tunnel. From proprietary field 45297. | integer Native | Non-directional |
| BFD rx cnt | Count of received BFD packets. From proprietary field 45299. | counter | Non-directional |
| BFD tx cnt | Count of transmitted BFD packets, from proprietary field 45300. | counter | Non-directional |
| BFD rx octets | Count of received BFD octets, from proprietary field 45304. | bytes counter | Non-directional |
| BFD tx octets | Count of transmitted BFD octets, from proprietary field 45305. | bytes counter | Non-directional |

## Cisco IOS XR Dimensions

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/logo-cisco-063h120w(2).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A00Z&se=2026-05-12T09%3A45%3A00Z&sr=c&sp=r&sig=vsuQg87%2FsR%2FhXPywVi1HrcSQi8EgmVT1z%2Bo4jyjKkTM%3D)These dimensions are used to filter or group-by on fields in flow records from Cisco products running the IOS XR operating system. These fields contain IPFIX "entity" values as described in **IPFIX Information Elements**.

> **TIP**: For additional information, see the Cisco document **Configure NetFlow on IOS XR**.

| Dimension Name (Portal) | Description | Type: value column | Direction |
| --- | --- | --- | --- |
| Dest ToS | Entity 55: The IPFIX postIpClassOfService value, which is the post-observation value of  ToS (Type of Service) field (IPv4) or Traffic Class field (IPv6) in the packet header. | integer Native | Dst only |
| Minimum TTL | Entity 52: The minimum value observed for the TTL (time to live) field in the IP header of any packet in this flow. | integer Native | Non-directional |
| Maximum TTL | Entity 55: The maximum value observed for the TTL (time to live) field in the IP header of any packet in this flow. | integer Native | Non-directional |
| Forwarding Status | Entity 89: The two-bit forwarding status of the flow and associated six-bit reason code or flag. | integer Native | Non-directional |

## Cisco nvzFlow Dimensions

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/logo-cisco-063h120w(2).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A00Z&se=2026-05-12T09%3A45%3A00Z&sr=c&sp=r&sig=vsuQg87%2FsR%2FhXPywVi1HrcSQi8EgmVT1z%2Bo4jyjKkTM%3D)Kentik dimensions for Cisco nvzFlow are described in the table below.

### nvzFlow Overview

Cisco AnyConnect version 4.2, a Cisco VPN client, introduced support for extended endpoint network communication visibility. This is achieved by exporting network metadata using the IPFIX protocol, Cisco Network Visibility Flow protocol (**nvzFlow**). The protocol provides lightweight network visibility of endpoints with a small set of high-value endpoint context data, enabling IT administrators to better understand the following vectors:

* **User**: Who is accessing enterprise networks and using applications?
* **Device**: Which device is used for access?
* **Client application**: Which application is used for access?
* **Location**: Which location is access being performed from?
* **Destination application**: Which SaaS application has been accessed?

| Dimension Name (Portal) | Description | Type: value column | Direction |
| --- | --- | --- | --- |
| Unique ID | 20-byte Unique ID that identifies the endpoint. From proprietary field 12332. | string Native | Non-directional |
| Logged In User | This is the logged in user on the device, in the form Authority/Principle. From proprietary field 12333. | string Native | Non-directional |
| Process Account | Authority/Principle of the process associated with the flow, e.g., “ACME/JSMITH, <machine>/JSMITH”. From proprietary field 12338. | string Native | Non-directional |
| Process Name | Name of the process associated with the flow, e.g., “firefox.exe”. From proprietary field 12340. | string Native | Non-directional |
| Process Hash | SHA256 hash of the process image on disk associated with the flow. From proprietary field 12341. | string Native | Non-directional |
| Parent Process Account | Authority/Principle of the parent process associated with the flow, e.g., “ACME/JSMITH, <machine>/JSMITH”. From proprietary field 12339. | string Native | Non-directional |
| Parent Process Name | Name of the parent process associated with the flow, e.g., “firefox.exe”. From proprietary field 12342. | string Native | Non-directional |
| Parent Process Hash | SHA256 hash of the process image on disk of the parent process associated with the flow. From proprietary field 12341. | string Native | Non-directional |
| DNS Suffix | Per-interface DNS suffix configured on the adapter associated with the flow for the endpoint. From proprietary field 12344. | string Native | Non-directional |
| Destination Hostname | The FQDN (hostname) that resolved to the destination IP on the endpoint, e.g., “www.google.com”. From proprietary field 12338. | string Native | Non-directional |
| Layer 4 Byte Count in | Total number of incoming bytes on the flow at Layer4 (Transport). Payload only, without L4 headers. From proprietary field 12346. | bytes counter | Non-directional |
| Layer 4 Byte Count Out | Total number of outgoing bytes on the flow at Layer4 (Transport). Payload only, without L4 headers. From proprietary field 12347. | bytes counter | Non-directional |

## Cisco SD-WAN vEdge Dimensions

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/logo-cisco-063h120w(2).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A00Z&se=2026-05-12T09%3A45%3A00Z&sr=c&sp=r&sig=vsuQg87%2FsR%2FhXPywVi1HrcSQi8EgmVT1z%2Bo4jyjKkTM%3D)These dimensions are used to filter or group-by on IPFIX fields (see **IPFIX Information Elements Exported to the Collector**) in `cflowd` records from Cisco SD-WAN routers.

> **TIP**: For more information about these devices, refer to Cisco’s **Cisco SD-WAN vEdge Routers Data Sheet**.

| Dimension Name (Portal) | Description | Type: value column | Direction |
| --- | --- | --- | --- |
| Maximum packet length | Length of the largest packet observed for this flow. | integer Native | Non-directional |
| Minimum packet length | Length of the smallest packet observed for this flow. | integer Native | Non-directional |
| VPN identifier | VEdge VPN identifier. | integer Native | Non-directional |
| Field 4322 | Reserved for internal use. | integer Native | Non-directional |
| Flow end reason | Reason for the flow termination (see **IANA IPFIX Entities**). | integer Native | Non-directional |

## FortiGate Dimensions

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/logo-fortinet-30h256w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A00Z&se=2026-05-12T09%3A45%3A00Z&sr=c&sp=r&sig=vsuQg87%2FsR%2FhXPywVi1HrcSQi8EgmVT1z%2Bo4jyjKkTM%3D)These dimensions are used to filter or group-by on fields from devices whose Type is set in Kentik to "Fortinet FortiGate" (see **Device General Settings**). The dimensions enabled for these devices store various types of information that is specific to the data fields of NetFlow templates supported by FortiOS.

> **TIP**: For more information about Fortinet FortiOS NetFlow templates, see **NetFlow templates**.

| Dimension Name (Portal) | Description | Type: value column | Direction |
| --- | --- | --- | --- |
| Post-NAT Address | The IPv4 or IPv6 source/destination address in the IP packet header, as modified by the firewall during network address translation after the packet traversed the device. Extracted from the following IPFIX fields:   * 225: Post-NAT Source IPv4 Address * 281: Post-NAT Source IPv6 Address * 226: Post-NAT Destination IPv4 Address * 282: Post-NAT Destination IPv6 Address | ip-address Native | Src/Dst |
| Post-NAT Transport Port | The source/destination port identifier in the transport header, as modified by the firewall during network address port translation after the packet traversed the device. Extracted from the following IPFIX fields:   * 227: Post-NAT Source Transport Port * 228: Post-NAT Destination Transport Port | integer Native | Src/Dst |
| Flow Flags | Extracted from IPFIX field 65. | integer Native | Non-directional |
| Forwarding Status | Extracted from IPFIX field 89. For possible values, see **IANA specification**. | Integer/string Native | Non-directional |
| Flow End Reason | Extracted from the IPFIX field 136. For possible values, see **IANA specification**. | integer Native | Non-directional |
| Application Name | Extracted from IPFIX field 96. | string Native | Non-directional |
| Application Category | Extracted from IPFIX field 372. | string Native | Non-directional |

## Juniper PFE Syslog Dimensions

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/logo-juniper-045h159w(1).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A00Z&se=2026-05-12T09%3A45%3A00Z&sr=c&sp=r&sig=vsuQg87%2FsR%2FhXPywVi1HrcSQi8EgmVT1z%2Bo4jyjKkTM%3D)These dimensions represent event-triggered syslog messages from a Juniper switch equipped with a Packet Forwarding Engine (see Juniper’s **Informal Guide to Packet Forwarding**). Here are a few considerations:

* If a given switch has multiple PFEs their messages are grouped as if they were from a single PFE.
* In addition to the dimensions below, the remaining portion of the syslog message may contain information (e.g., MAC address, protocol, IP addresses, and bytes) that is accessible via KDE dimensions that aren't device-specific.

> **Note:** Syslog from Juniper PFE is ingested into KDE via the Flow Proxy capability of Kentik's **Universal Agent**. For further information, contact Kentik (see **Customer Care**).

| Dimension Name (Portal) | Description | Type: value column | Direction |
| --- | --- | --- | --- |
| Message | The first 64 chars of the PFE syslog message. | string Native | Non-directional |
| Subtype | The subtype of the message, e.g., "FW" for firewall. | string Native | Non-directional |
| Interface | The device interface on which the event occurred. | string Native | Non-directional |
| Event | The nature of the event, e.g., "D" for dropped packets. | string Native | Non-directional |

## Nokia Layer 2 Dimensions

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/logo-nokia-36h144w(1).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A00Z&se=2026-05-12T09%3A45%3A00Z&sr=c&sp=r&sig=vsuQg87%2FsR%2FhXPywVi1HrcSQi8EgmVT1z%2Bo4jyjKkTM%3D)These dimensions are used to filter or group-by on fields in IPFIX flow records from Nokia routers connected to a "Dot1Q" (IEEE 802.1Q) VLAN.

| Dimension Name (Portal) | Description | Type: value column | Direction |
| --- | --- | --- | --- |
| Physical Interface | The router's index for the physical interface where the packets of the flow were sent/received. | integer UDR | Src/Dst |
| Dot1q Vlan ID | The value of the 12-bit VLAN Identifier portion of the Tag Control Information field of an Ethernet frame. | integer UDR | Src/Dst |
| Dot1q Customer Vlan ID | The value of the Customer VLAN identifier in the Customer VLAN Tag (C-TAG) Tag Control Information (TCI) field. | integer UDR | Src/Dst |
| SAP | The value of the Service Access Point (SAP) field derived from the vendor MIB. | String virtual | Src/Dst |

## Palo Alto Networks Firewall

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/logo-pan-47h256w(1).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A00Z&se=2026-05-12T09%3A45%3A00Z&sr=c&sp=r&sig=vsuQg87%2FsR%2FhXPywVi1HrcSQi8EgmVT1z%2Bo4jyjKkTM%3D)These dimensions are used to filter or group-by on fields in flow records from Palo Alto Networks firewalls. In addition to the port, IP address, and type of packets, the data identifies the application and includes firewall event information.

> **TIP**: For more context on these dimensions, see the Palo Alto Networks document **NetFlow Templates**.

| Dimension Name (Portal) | Description | Type: value column | Direction |
| --- | --- | --- | --- |
| Post-NAT Transport Port | The source/destination port identifier in the transport header, as modified by the firewall during network address port translation after the packet traversed the interface. | integer Virtual | Src/Dst |
| Post-NAT Address | The IPv4 source/destination address in the IP packet header, as modified by the firewall during network address translation after the packet traversed the interface. | string Virtual | Src/Dst |
| ICMP Type | Internet Control Message Protocol (ICMP) packet type. This is reported as: ICMP Type \* 256 + ICMP code | integer Virtual | Non-directional |
| Flow ID | An identifier of a flow that is unique within an observation domain. You can use this information element to distinguish between different flows if flow keys such as IP addresses and port numbers are not reported or are reported in separate records. The flowID corresponds to the session ID field in Traffic and Threat logs. | integer Virtual | Non-directional |
| Application ID | The name of an application (up to 32 bytes). | string Virtual | Non-directional |
| User ID | A username that User-ID identified. The name can be up to 64 bytes. | string Virtual | Non-directional |
| Firewall Event | Indicates a firewall event:   * 0 = Ignore (invalid)—Not used. * 1 = Flow created—The NetFlow data record is for a new flow. * 2 = Flow deleted—The NetFlow data record is for the end of a flow. * 3 = Flow denied—The NetFlow data record indicates a flow that firewall policy denied. * 4 = Flow alert—Not used. * 5 = Flow update—The NetFlow data record is sent for a long-lasting flow, which is a flow that lasts longer than the Active Timeout period configured in the NetFlow server profile. | integer Virtual | Non-directional |
| Direction | The direction of the flow:   * 0 = ingress * 1 = egress | integer Virtual | Non-directional |

## sFlow Tunnel Decode Dimensions

These dimensions are used to filter or group-by on fields from the headers of VXLAN-encapsulated packets on virtual networks in data centers. The dimensions are on available for traffic on devices that report flow using sFlow and whose Type (see **Device General Settings**) is set to "sFlow Tunnel Decode".

> **Note:** For information about usage, see **Using VXLAN**.

| Dimension Name (Portal) | Description | Type: value column | Direction |
| --- | --- | --- | --- |
| VXLAN VTEP 0/1 IP Address | The VXLAN tunnel endpoint used to map tenants’ end devices to VXLAN segments and to perform VXLAN encapsulation and decapsulation. | string UDR | Src/Dst |
| VXLAN 0/1 VNI | The IP address of an encapsulated packet inside a virtual network (VXLAN tunnel). | integer UDR | Non-directional |

## Silver Peak Dimensions

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/logo-silver_peak-045h196w(1).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A00Z&se=2026-05-12T09%3A45%3A00Z&sr=c&sp=r&sig=vsuQg87%2FsR%2FhXPywVi1HrcSQi8EgmVT1z%2Bo4jyjKkTM%3D)These dimensions are used to filter or group-by on flow records from Silver Peak appliances running VXOA software (version 8.2.1 or higher). Silver Peak analyzes the actual packets as traffic flows through their appliances, identifies the applications (e.g., SaaS service) with which each packet is associated, and prioritizes routing by applying application-specific rules. For more context on these dimensions, see the Silver Peaks page about **EdgeConnect Appliance Logs**.

| Dimension Name (Portal) | Description | Type: value column | Direction |
| --- | --- | --- | --- |
| Application name | The name of an application as identified by a Silver Peak VXOA appliance. | string UDR | Non-directional |
| Business Intent Overlay | The Silver Peak **Business Intent Overlay** | UDR | Non-directional |
| Application Category | The category grouping of the applications identified by Silver Peak Edge Connect | UDR | Non-directional |
| From Zone | The firewall zone the traffic is coming from. | UDR | Src |
| To Zone | The firewall zone the traffic is going to. | UDR | Dst |
| Firewall Event | Indicates a firewall event:   * 0 = Ignore (invalid)—Not used. * 1 = Flow created—The NetFlow data record is for a new flow. * 2 = Flow deleted—The NetFlow data record is for the end of a flow. * 3 = Flow denied—The NetFlow data record indicates a flow that firewall policy denied. * 4 = Flow alert—Not used. * 5 = Flow update—The NetFlow data record is sent for a long-lasting flow, which is a flow that lasts longer than the Active Timeout period configured in the NetFlow server profile. | integer UDR | Non-directional |

## VMware SD-WAN Dimensions

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/logo-vmware-34h214w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A00Z&se=2026-05-12T09%3A45%3A00Z&sr=c&sp=r&sig=vsuQg87%2FsR%2FhXPywVi1HrcSQi8EgmVT1z%2Bo4jyjKkTM%3D)Kentik dimensions for VMware SD-WAN are described in the below table.

### Overview

Kentik ingests and analyzes IPFIX flow data from VMware SD-WAN Edge devices (formerly Velocloud). These devices export enriched telemetry that goes beyond standard traffic metrics, mapping directly to the configurations managed within the VMware SASE Orchestrator.

The dimensions in the table below bridge the gap between raw network traffic and SD-WAN operational logic:

* **Network Segmentation & Policy:** While standard flows use VRFs, VMware exports the Ingress VRF ID, which maps directly to specific “Segments” configured in the Orchestrator. Business logic is further captured by the Business Policy ID and VC Link Policy, which identify the specific steering and remediation rules applied to a flow.
* **Application Intelligence:** The Edge device performs Deep Packet Inspection (DPI) to identify traffic, providing the Application Name and Application Category (IANA fields 96 and 372) directly in the flow record.
* **Path Selection & Performance:** Dimensions like VC Flow Path, VC Route Type, and VC Priority reveal how the SD-WAN overlay is routing traffic—whether it’s traveling via a Gateway, Hub, or Direct-to-Cloud.
* **Flow Aggregation:** Unlike traditional NetFlow, VMware SD-WAN Edge aggregates flows with different source ports if they share the same source/destination IPs and destination port. The Delta Flow Count field is used to track the number of original flows contributing to these aggregated records.
* **Visibility & NAT:** To maintain end-to-end visibility through address translation, the Post-NAT Address fields capture the modified IP headers after the packet traverses the firewall.

| Dimension Name (Portal) | Description | Type: value column | Direction |
| --- | --- | --- | --- |
| Biflow Direction | A description of the direction assignment method used to assign the biflow source and destination. This information element may be present in flow data record or applied to all flows exported from an exporting process or observation domain using IPFIX options. **IANA field 239**. | integer Native | Non-directional |
| Flow ID | An identifier of a Flow that is unique within an Observation Domain. This information element can be used to distinguish between different Flows if Flow Keys such as IP addresses and port numbers are not reported or are reported in separate records. IANA field 148. | string Native | Non-directional |
| Flow end reason | This information is extracted from the IPFIX field 136. Possible values of this field can be found on the IANA’s specification: **IPFIX Flow End Reason**. | integer  Native | Non-directional |
| Ingress VRF ID | A unique identifier of the VRF name where the packets of this flow are being received. This identifier is unique per metering process. This maps to the VMware SASE orchestrator. segments. A segment should be visualized and reported as a separated L3 domain within edge. IANA field 234. | string Native  string Native | Non-directional |
| Application Name | Application Name, IANA field 96. | string Native | Non-directional |
| Application Category | An Attribute that provides a first level categorization for each application ID. IANA field 372. | string Native | Non-directional |
| Destination UUI | Destination node UUID. | string Native | Non-directional |
| VC Priority | This identifies the BizPolicy “Priority” classification applied. Unset should monitored to deduce a warning since it would only occur during overflow:   * 0 = Unset * 1 = Control * 2 = High * 3 = Normal * 4 - Low | integer  Native | Non-directional |
| VC Route Type | This identifies the path type out to internet the flow is taking. Unset should be monitored to deduce a warning since it would only occur during overflow:   * 0 = Unset * 1 = Gateway * 2 = Direct * 3 = Backhaul (using Hub to Internet) | integer  Native | Non-directional |
| VC Link Policy | This value provides the type of link steering and remediation configured for this application under BizPolicy.   * 0 = NA * 1 = Fixed * 2 = Load balance * 3 = Replicate | integer  Native | Non-directional |
| VC Traffic Type | This identifies the BizPolicy “Service Class” classification applied.   * 0 = Realtime * 1 = Transactional * 3 = Bulk | integer  Native | Non-directional |
| VC Flow Path | This identifies the type of “path” the flow is taking.   * 0 = Edge2CloudViaGateway (SaaS optimized) * 1 = Edge2CloudDirect (SaaS not optimized) * 2 = Edge2EdgeViaGateway (spoke2hub2 spoke via VCG) * 3 = Edge2EdgeViaHub (spoke2hub2 spoke via PDC  Hub) * 4 = Edge2EdgeDirect (Edge2Edge dynamic) * 5 = Edge2DataCenterDirect (Edge2PDC using underlay routing) * 6 = Edge2DataCenterViaGateway (Edge2PDC using  NVS) * 7 = Edge2Backhaul (Edge2intern et using PDC Hub) * 8 = Edge2Proxy * 9 = Edge2OPG (PGW) * 10 = Routed (path using underlay routing) * 11 = Edge2CloudViaSecurityService (path using a CASB service to internet) | integer  Native | Non-directional |
| Business Policy ID | The business policy logical ID that this flow matches. This value is a UUID and must be mapped to a BizPolicy via Orchestrator API. | integer  Native | Non-directional |
| Next Hop UUID | Next hop UUID for this flow. This will be populated in case of overlay traffic. This value identifies the device that is in the path between source and destination in the SD-WAN overlay network (not underlay). | integer  Native | Non-directional |
| Post-NAT Address | The IPv4 or IPv6 source/destination address in the IP packet header, as modified by the firewall during network address translation after the packet traversed the device. This information is extracted from the following IPFIX fields:   * Post-NAT Source IPv4 Address (IPFIX field 225) * Post-NAT Source IPv6 Address (IPFIX field 281) * Post-NAT Destination IPv4 Address (IPFIX field 226) * Post-NAT Destination IPv6 Address (IPFIX field 282) | ip-address  Native | Src/Dst |
| Responder Bytes | The total number of layer 4 bytes in a flow from the responder since the previous report. The responder is the device which replies to the initiator and remains the same for the life of the session. | bytes counter | Non-directional |
| Responder Packets | The total number of layer 4 packets in a flow from the responder since the previous report. The responder is the device which replies to the initiator and remains the same for the life of the session. | counter | Non-directional |
| Delta Flow Count | The conservative count of original flows contributing to this aggregated flow; may be distributed via any of the methods expressed by the valueDistributedMethod information element. IANA field 3. | counter | Non-directional |
| Replicated Packets Rx | Count of replicated packets received for the flow. | gauge | Non-directional |
| Replicated Packets Tx | Count of packets replicated for the flow. | gauge | Non-directional |
| Lost Packets Rx | Count of packets lost for the flow at the receive. | gauge | Non-directional |
| Retransmitted Packets Tx | Count of packets retransmitted for the flow. | gauge | Non-directional |
| TCP RTT ms | Maximum RTT observed for a TCP flow. | gauge | Non-directional |
| TCP Retransmits | Count of TCP packets retransmitted for the flow. | gauge | Non-directional |

## VXLAN Dimensions

These dimensions are used to filter or group-by on fields from the headers of VXLAN-encapsulated packets within virtual networks in data centers. The dimensions are only available for traffic on devices that report flow using sFlow and have a Type configured as "VXLAN" (see **Device General Settings**).

> **Note:** For information about usage, see **Using VXLAN**.

| Dimension Name (Portal) | Description | Type: value column | Direction |
| --- | --- | --- | --- |
| VXLAN VNI 0/1 MAC Address | The MAC address of an encapsulated packet inside a virtual network (VXLAN tunnel). | string UDR | Src/Dst |
| VXLAN VNI 0/1 IP Address | The IP address of an encapsulated packet inside a virtual network (VXLAN tunnel). | string UDR | Src/Dst |
| VXLAN VNI 0/1 Port | The Port of an encapsulated packet inside a virtual network (VXLAN tunnel). | integer UDR | Src/Dst |
| IP TTL | The value of the IP TTL (time to live) field in the header of a VXLAN encapsulated packet. | integer UDR | Non-directional |
| VXLAN VNI 0/1 | A virtual network identifier (VNI) that identifies a specific virtual network in the data plane. | integer UDR | Non-directional |
| VXLAN VNI 0/1 DSCP | A DSCP (differentiated services code point) value from the DS field in the header of a VXLAN encapsulated packet. | integer UDR | Non-directional |
| VXLAN VNI 0/1 TCP Flags | TCP flags set in the header of a VXLAN encapsulated packet. | integer UDR | Non-directional |
| VXLAN VNI 0/1 IP TTL | The value of the TTL (time to live) field in the IP header of a VXLAN encapsulated packet. | integer UDR | Non-directional |
| VXLAN VNI 0/1 Protocol | The protocol of an encapsulated packet inside a virtual network (VXLAN tunnel). | integer UDR | Non-directional |
