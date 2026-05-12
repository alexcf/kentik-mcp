# Using MPLS

Source: https://kb.kentik.com/docs/using-mpls

---

This article covers Kentik’s support for Multiprotocol Label Switching (MPLS).

> **Note:** For a list of MPLS-derived filtering and group-by dimensions for Kentik’s Data Explorer, see **MPLS Dimensions**.

## About MPLS

Multiprotocol Label Switching (MPLS) is a high-performance routing scheme designed as an alternative to traditional IP routing. By using short path labels rather than long network addresses, MPLS avoids complex lookups in a router’s local routing table, significantly increasing packet-forwarding speed.

### How it Works

In an MPLS domain, network operators define a **Label-Switched Path** (LSP) across a series routers. The process follows a specific hierarchy of roles:

* **Ingress Router:** The entry point to the MPLS domain. It examines the destination IP, determines the Forwarding Equivalence Class (FEC), and "pushes" an MPLS header (label stack) onto the packet.
* **Transit (Interior) Routers:** Also known as Label Switch Routers (LSRs), these devices swap the top-most label in the stack for a new one based on their Label Forwarding Information Base (LFIB) and forward the packet to the next hop.
* **Egress Router:** The exit point of the domain. It "pops" the final label and forwards the original IP packet to its destination.

![Diagram illustrating MPLS domain with ingress, egress, and core LSR connections.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/UM-mpls-diagram.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A06Z&se=2026-05-12T09%3A33%3A06Z&sr=c&sp=r&sig=8LqTH9d2WyCmTIKzeVWUp%2FvDHo%2FwJI5A7qNTxQ35Ytg%3D)

*Diagram illustrating MPLS domain with ingress, egress, and core LSR connections.*

By eliminating individual path calculations at every hop, MPLS reduces latency and resource consumption. Additionally, the label stack supports quality-of-service (QoS) bits, allowing operators to prioritize sensitive traffic like voice or video over standard data.

## MPLS Ingest

As a packet traverses an MPLS path, each participating router updates its flow cache to reflect the labels it is actively swapping. When these records are exported to Kentik, we ingest the label stack to provide hop-by-hop visibility into the LSP.

### Why IP Data Might be Missing

In a standard IP-over-MPLS environment, you expect to see both labels and IP addresses. However, flow records from transit routers sometimes lack standard IP information (Source/Dest IP) for two primary reasons:

* **Non-IP Traffic:** The router is acting as a "pure" transit device for Layer 2 traffic (such as ATM or Frame Relay) that does not contain an IP header to inspect.
* **Template Limitations:** Many NetFlow/IPFIX templates are not configured to export the MPLS label stack and the IPv4/v6 payload simultaneously. If the template prioritize labels, the IP fields may be left blank.

Kentik solves this by normalizing these varied records into a consistent set of **MPLS Dimensions**, allowing you to group and filter traffic even when the underlying IP data is unavailable.

### MPLS Label Capture

Kentik typically captures up to three labels from the stack, which are mapped to specific dimensions in the portal:

* **MPLS Label 1 (Top):** Usually the transport label used to reach the next hop.
* **MPLS Label 2:** Often a service or VPN label.
* **MPLS Label 3 (Bottom):** Frequently used in complex "carrier-of-carrier" scenarios.

## Applying MPLS in Kentik

Without MPLS support, transit routers in an LSP may report traffic volumes that don't align with SNMP data, or export records missing source/destination IPs. Kentik bridges this gap by allowing you to query "reserved" label values, designated by **IETF RFC7274**, to identify where specific policies are being imposed on your traffic.

### Reserved Label Use Cases

Use the following reserved labels in the **MPLS Label 1** dimension to audit your LSP boundaries:

| Label Value | Definition | Use Case in Kentik |
| --- | --- | --- |
| `0` (IPv4) / `2` (IPv6) | Explicit NULL | Identify Egress Traffic: Indicates the end of an LSP. Use this to find where traffic is leaving the MPLS domain and entering a standard IP domain. |
| `1` | Router Alert | Identify Local Processing: Indicates the packet is being delivered to a local software module. Use this to find traffic dropped, processed locally, or re-inserted into the domain. |

![Filtering options interface showing conditions for MPLS Label 1 selection.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/UK-MPLS.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A06Z&se=2026-05-12T09%3A33%3A06Z&sr=c&sp=r&sig=8LqTH9d2WyCmTIKzeVWUp%2FvDHo%2FwJI5A7qNTxQ35Ytg%3D)

### Auditing QoS and Service Policies

Because MPLS does not directly support the IP DSCP field, operators use the **EXP (Experimental) bits** in the MPLS header to preserve quality-of-service settings across a domain.

* **Troubleshooting Policy Mismatches:** Filter on the `MPLS Label 1 EXP` and `MPLS Label 2 EXP` dimensions.
* **Audit Goal:** Identify areas where service policies (Priority vs. Best Effort) are not being correctly applied as packets traverse the MPLS forwarding domain.

## Kentik Support for 6PE

Kentik supports IPv6 Provider Edge (6PE) technology and the BGP IPv6 Labeled-Unicast (LU) address family. To ingest these records:

1. **Configure BGP:** In **Settings »** **Network Devices**, open the **Device Dialog** for your 6PE device.
2. **Set Route Selection:** On the **BGP** tab, set **BGP Route Selection** to “VPN table, fallback to Labeled-Unicast table, fallback to Unicast table*”.*
3. **Establish Session:** Ensure a BGP session with IPv6 Labeled-unicast AF is active between the device and Kentik.

Once configured, Kentik populates standard BGP dimensions, including **Ultimate Exit**(identified via the IPv4-mapped IPv6 address).

## Further Reading

For additional technical background on MPLS specifications, see the following resources:

* **TechTarget: MPLS Overview**
* **Cisco FAQ: MPLS Basics**
* **Network World: MPLS Explained**
