# Routing Dimensions

Source: https://kb.kentik.com/docs/routing-dimensions

---

This article covers Kentik’s Data Explorer **dimensions** related to traffic routing.

> **Note:** As of May 1st, 2025, the **Query SQL Method** has been deprecated and is no longer supported.

## IP & BGP Routing

The **IP & BGP Routing** dimensions category is used to filter or group-by on IP addresses (IPv4 or IPv6), protocol (e.g., TCP or UDP), TCP flags, and ToS, as well as routing information including source and destination AS, AS path, AS names, community, prefixes, and hops.

### IP Info Dimensions

Dimensions related to IP properties:

| Dimension name (portal) | Description | Type: value column | Direction:  KDE name(s) |
| --- | --- | --- | --- |
| IP/CIDR | The source/destination IP address, either IPv4 or IPv6, of the flow. | string Native | Src/Dst: `inet_src_addr` `inet_dst_addr` |
| Site by IP | The site name for the source/destination IP, e.g., “acme\_dfw1”. | string  Native | Src/Dst:  `kt_src_site_title`  `kt_dst_site_title` |
| Site Type by IP | The site type for the source/destination IP, e.g., “Cloud”. | string  Native | Src/Dst:  `kt_src_site_type`  `kt_dst_site_type` |
| Port Number | Layer 4 source/destination port (e.g., 80, 443).  **Note:** If Protocol is ICMP, this value equals the result of this formula: (ICMP Type \* 256) + ICMP code. | integer Virtual | Src/Dst: `l4_src_port` `l4_dst_port` |
| Protocol | The number of the protocol. See **Wikipedia's List of IP Protocol Numbers**. | integer Native | Non-directional: `protocol` |
| Protocol Name | The name of the protocol followed by the corresponding protocol number in parentheses, e.g., TCP (6). In SQL, supports case-insensitive equality and IN matching. | string Virtual | Non-directional: `i_protocol_name` |
| INET Family | The address family of the flow, either 4 (IPv4) or 6 (IPv6). | integer Native | Non-directional: `inet_family` |
| DSCP | A DSCP (differentiated services code point) value from the DS field in a packet's IP header, which classifies the packet's contents to enable differentiated QoS. | integer Native | Non-directional: `dscp` |
| ToS | An 8-bit value, typically made up of a six-bit Differentiated Services Code Point (DSCP) field and a two-bit Explicit Congestion Notification (ECN) field. | integer Native | Non-directional: `tos` |
| IPv6 Flow Label | Allows network operators to analyze and categorize IPv6 traffic based on a 20-bit field in the IPv6 header (see **IPv6 Flow Label Dimension**) | unsigned32 | Non-directional: `flowLabelIPv6` |

### BGP Dimensions

Dimensions related to BGP properties (see **About Kentik BGP**):

| Dimension name (portal) | Description | Type: value column | Direction  KDE name(s) |
| --- | --- | --- | --- |
| Route Prefix + LEN | The BGP table prefix concatenated with the prefix length for the source/destination IP of the flow. | string  Native | Src/Dst:  `src_route_prefix_len`  `dst_route_prefix_len` |
| Route Prefix | The BGP table prefix, either IPv4 or IPv6, that contains the source/destination IP of the flow. | string Native | Src/Dst: `inet_src_route_prefix` `inet_dst_route_prefix` |
| Route LEN | The BGP prefix length for the source/destination IP of the flow. | integer Native | Src/Dst: `src_route_length` `dst_route_length` |
| AS Number | The origin ASN associated with the source/destination IP of the flow. | bigint Native | Src/Dst: `src_as` `dst_as` |
| AS Name | The name associated with AS Number. | string Virtual | Src/Dst: `i_src_as_name i_dst_as_name` |
| AS Group | A label assigned to a collection of ASes (see **About AS Groups**). | string  Virtual | Src/Dst: `kt_src_as_group` `kt_dst_as_group` |
| Next Hop IP/CIDR | The BGP next-hop IP address, either IPv4 or IPv6, for the source/destination IP of the flow (see **About BGP**). | string Native | Src/Dst: `inet_src_next_hop` `inet_dst_next_hop` |
| Next Hop AS Number | The ASN in the first position of the AS\_PATH for the source IP of the flow (see **About BGP**). | integer Native | Src/Dst: `src_nexthop_as` `dst_nexthop_as` |
| Next Hop AS Name | Name of Next Hop AS Number | string Virtual | Src/Dst: `i_src_nexthop_as_name` `i_dst_nexthop_as_name` |
| 2nd Hop AS Number | The ASN in the second position of the AS\_PATH for the source/destination IP of the flow (see **About BGP**). | integer Native | Src/Dst: `src_second_asn` `dst_second_asn` |
| 2nd Hop AS Name | Name of 2nd Hop AS Number. | string Virtual | Src/Dst: `i_src_second_asn_name` `i_dst_second_asn_name` |
| 3rd Hop AS Number | The ASN in the third position of the AS\_PATH for the source/destination IP of the flow (see **About BGP**). | integer Native | Src/Dst: `src_third_asn` `dst_third_asn` |
| 3rd Hop AS Name | Name of 3rd Hop AS Number. | string Virtual | Src/Dst: `i_src_third_asn_name` `i_dst_third_asn_name` |
| AS Path | The BGP ASPATH for the flow’s source/destination IP (see **About BGP**). | string Native | Src/Dst: `src_bgp_aspath` `dst_bgp_aspath` |
| BGP Community | The set of BGP communities associated with the flow’s source/destination IP (see **About BGP**). | string Native | Src/Dst: `src_bgp_community` `dst_bgp_community` |
| BGP Extended Community | A set of BGP communities used to extend the standard BGP Community attribute, often to carry more specific routing information like specific route targets or encapsulation details (e.g., “RT:65000:1000”) | string  Native | Src/Dst:  `src_bgp_ext_community`  `dst_bgp_ext_community` |
| BGP Large Community | A set of BGP communities designed to provide a larger format than the standard community, typically used to carry 96-bit values for Autonomous System (AS) numbers and local values. | string  Native | Src/Dst:  `src_bgp_large_community`  `dst_bgp_large_community` |
| BGP Origin | The BGP path attribute that indicates the origin of the route, i.e. how the route was injected into BGP (e.g., “IGP” ). | string Native | Src/Dst:  `src_bgp_origin`  `dst_bgp_origin` |
| RPKI Validation Status | The RPKI (Resource Public Key Infrastructure; see **RPKI Documentation**) status of a prefix in a BGP-advertised route, which indicates whether the route would be used or dropped if the router were configured to enforce strict route validation. | string Virtual | Dst: `i_dst_rpki_name` |
| RPKI Quick Status | Provides a simplified view of RPKI status, enabling easier determination of the action to take on the prefix. | string Virtual | Dst: `i_dst_rpki_min_name` |

### VRF Dimensions

Dimensions related to VRF properties:

| Dimension name (portal) | Description | Type: value column | Direction KDE name(s) |
| --- | --- | --- | --- |
| VRF Name | The locally significant name of the VRF via which this flow was routed (input or output).  **Note:** VRF names may vary in different contexts. | string Virtual | Src/Dst: `i_input_vrf` `i_output_vrf` |
| VRF Route Distinguisher | Uniquely identifies the VRF via which this flow was routed (input or output). | string Virtual | Src/Dst: `i_input_vrf_rd` `i_output_vrf_rd` |
| VRF Route Target | Uniquely identifies a shared route (used by multiple VRFs) via which this flow was routed (input or output). | string Virtual | Src/Dst: `i_input_vrf_rt` `i_output_vrf_rt` |
| VRF Extended Route Distinguisher | An encoding of the VRF route distinguisher (for Kentik internal use only). | integer Native | Src/Dst: `input_vrf` `output_vrf` |

### Per-flow Metrics

These metrics are available as dimensions that can be used to filter or group-by based on stats related to the bytes and packets of the flow:

| Dimension name (portal) | Description | Type: value column | Direction: KDE name(s) |
| --- | --- | --- | --- |
| Packet Size | Packet size of flow (bytes/packet). | integer Native | Non-directional:  `sampledpktsize` |
| Packet Size (nearest 100) | Packet size of flow (bytes/packets) rounded down to the nearest multiple of 100. | integer Native | Non-directional: `sampledpktsize_100` |
| Sampling Rate \* 100 | The rate at which traffic was sampled when flow was collected (see **Flow Sampling**). | integer Native | Non-directional: `sample_rate` |

### TCP Performance Dimensions

Dimensions related to TCP retransmission and windowing, which indicate network and application performance issues:

| Dimension name  (portal) | Description | Type:  value  column | Direction:  KDE name(s) |
| --- | --- | --- | --- |
| TCP Retransmits | The number of outbound TCP packets that were retransmitted by the source/sender. | string  Native | Non-directional: `retransmitted_out_pkts` |
| Repeated TCP Retransmits | The number of times the same TCP packet was retransmitted multiple times (i.e., multiple retransmissions for a single sequence number). | string  Native | Non-directional: `repeated_retransmits` |
| TCP Receive Window | The size of the TCP receive buffer advertised by the flow's source/destination endpoint. | string  Native | Non-directional: `receive_window` |
| TCP Zero Windows | The number of times the source/destination endpoint advertised a receive window of zero, which forces the sender to stop transmitting data. | string  Native | Non-directional: `zero_windows` |
| TCP Client Latency (ms) | The network latency (Round Trip Time, RTT) in milliseconds as measured from the client's perspective (e.g., based on the initial TCP handshake). | integer  Native | Non-directional: `client_nw_latency_ms` |
| TCP Server Latency (ms) | The network latency (Round Trip Time, RTT) in milliseconds as measured from the server's perspective. | integer  Native | Non-directional: `server_nw_latency_ms` |

## MPLS Router Dimensions

The following table shows the MPLS-router-related dimensions that may be used for filtering or group-by in queries.

> **Did You Know?** Multi-protocol Label Switching (MPLS) is a routing scheme for network data that enables network operators to define label-switched paths that let routers move packets within the network without consulting a routing table at each hop (see **Using MPLS**).

| Dimension name (portal) | Description | Type: value column | Direction |
| --- | --- | --- | --- |
| MPLS Forwarding Type | This field denotes the label distribution scheme used by the router to forward the MPLS traffic.  Currently only Cisco ASR devices export this field. | integer UDR | Non-directional |
| MPLS Forwarding Address | The IP address of the destination PE (provider edge) router where the flow will exit the MPLS domain before re-entering standard IP forwarding domains. | integer UDR | Non-directional |
| MPLS Forwarding Address Prefix Length | The prefix length for the destination Forwarding Address of the flow. | integer UDR | Non-directional |
| MPLS Label 1 | The value of the top label assigned to the flow. | integer UDR | Non-directional |
| MPLS Label 1 EXP | The value of the experimental bits assigned to the flow. Typically, this is used to map IP-based Quality of Service (QoS) markings into MPLS domains so that routers can apply appropriate forwarding policies to MPLS flows. | integer UDR | Non-directional |
| MPLS Label 2 | The value of the second label assigned to the flow. | integer  UDR | Non-directional |
| MPLS Label 2 EXP | The value of the experimental bits assigned to the flow. Typically, this is used to map IP-based Quality of Service (QoS) markings into MPLS domains so that routers can apply appropriate forwarding policies to MPLS flows. | integer UDR | Non-directional |
| Forwarding Status | The two-bit forwarding status of the flow and associated six-bit reason code or flag. This dimension represents IPFIX entity 89 (IANA). | integer Native | Non-directional |
| MPLS Label 3 | The value of the third label assigned to the flow. | integer UDR | Non-directional |
| MPLS Label 3 EXP | The value of the experimental bits assigned to the flow. Typically, this is used to map IP-based Quality of Service (QoS) markings into MPLS domains so that routers can apply appropriate forwarding policies to MPLS flows. | integer UDR | Non-directional |

## IPv6 Flow Label Dimension

Kentik can ingest, store, and display the IPv6 Flow Label field from IPFIX flow records. This allows network operators to analyze and categorize IPv6 traffic based on a 20-bit field in the IPv6 header. When `kflow` receives an IPFIX template from a device containing Information Element (IE) 31, it extracts the value and forwards it to the Kentik Data Engine (KDE). The value is stored in a dedicated column, making it available as a first-class dimension in the Kentik portal.

### Device Configuration

Ingestion is enabled by default, and no configuration is required within the Kentik portal. Your network device (e.g., router, switch, etc.) must be configured to include the `flowLabelIPv6` (IE 31) in its active IPFIX export template. The specific commands vary by vendor and platform.

> **Note**: Refer to your device vendor’s documentation for instructions on adding information Element 31 to your IPFIX template configuration. Once the device begins exporting this field, Kentik will automatically detect and ingest it.

### Querying IPv6 Flow Label in Data Explorer

In **Data Explorer**, the dimension is available as IPv6 Flow Label, and it can be used in queries like any other dimension.

* **Group By**: Select IPv6 Flow Label as a Group By Dimension to view a breakdown of traffic volume per label.

  ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(835).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A23Z&se=2026-05-12T09%3A37%3A23Z&sr=c&sp=r&sig=kbyblH9m8VdvKK5Vk5WEIJ33Wi57JHhpBWwx84lKrf0%3D)
* **Filter**: Add a filter where IPv6 Flow Label equals a specific value to isolate all traffic related to a particular application flow or project.

### IPv6 Flow Label Use Cases

The IPv6 Flow Label dimension in Kentik allows several advanced monitoring and analysis workflows.

* **Research and High-Performance Computing (HPC)**: Large-scale scientific projects require moving massive datasets between institutions. The IPv6 Flow Label dimension can be used to tag traffic belonging to specific experiments or domains. Network managers at a transit provider can group by IPv6 Flow Label to quantify the bandwidth consumed by each scientific domain to help justify funding and plan capacity.
* **Differentiated Quality of Service (QoS)**: IPv6 Flow Label can be used to segregate traffic from a single application into different QoS classes to verify a network’s QoS policies (e.g., DSCP marking and queueing) are being applied correctly. Filter the server’s IP address and group by IPv6 Flow Label and DSCP to validate that high-priority flows are correctly marked with Expedited Forwarding (EF) and that best-effort flows are marked with Best Effort (BE).
* **ECMP Load Balancing and Flow Integrity**: In networks using Equal-Cost Multi-Path (ECMP) routing, routers hash packet headers to distribute traffic across available links. Including the IPv6 Flow Label in the hash provides greater entropy and enables an application to maintain a related sequence of packets on the same path, which prevents out-of-order delivery. You can use the IPv6 Flow Label dimension to investigate issues that might correlate with path changes.
* **Cloud and Multi-Tenant Traffic Identification**: Cloud providers can use the IPv6 Flow Label to distinguish between different types/tenants (e.g., live video streaming, API gateway, large file downloads) of traffic egressing from a shared infrastructure component that uses a single-source IP address. Operators can create dashboards that show traffic breakdowns per service by grouping the IPv6 Flow Label in addition to the source IP.
