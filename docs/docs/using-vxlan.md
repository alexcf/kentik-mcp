# Using VXLAN

Source: https://kb.kentik.com/docs/using-vxlan

---

This article provides a high-level overview to the use of VXLAN in Kentik.

![Kentik enables visibility into VXLAN encapsulated packets in virtual network traffic.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/UV-VXLAN_use_case-617h990w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A26Z&se=2026-05-12T09%3A32%3A26Z&sr=c&sp=r&sig=CzTzHncD%2BrF2wxiciCI1tEdghCUbqrL331UbYENX4lM%3D)

*Kentik enables visibility into VXLAN encapsulated packets in virtual network traffic.*

## About VXLAN

Formalized as **RFC7348**, Virtual eXtensible Local Area Network (VXLAN) is a protocol enabling the establishment of overlay networks using tunnels that carry Layer 3 UDP packets of encapsulated Layer 2 Ethernet frames. Developed primarily for use in data centers with multiple tenants, VXLAN addresses the limitations of VLAN to provide a more scalable and flexible approach to network virtualization.

## VXLAN in Kentik

Kentik supports VXLAN through two distinct device types and both require the device flow format to be **sFlow** (see **Flow Protocols**).

* **VXLAN Dimensions:** Best for high-level visibility into the VXLAN overlay. Data is stored in specialized **VXLAN Dimensions**. To enable this set the device **Type** to “VXLAN”.
* **sFlow Tunnel Decode:** Best for deep visibility into the traffic **inside** the tunnel by promoting encapsulated data to use standard dimensions (see **sFlow Tunnel Decode Dimensions**). To enable this, set the device **Type** to “sFlow Tunnel Decode”.

> **Note**: For more on setting device type, see **Device General Settings**.

## VXLAN Devices

The VXLAN device type provides visibility for traffic on sFlow-supporting network devices that use tunneling protocols, such as VXLAN, to carry traffic over data center overlay networks.

In standard monitoring, valuable information about virtual network traffic is “hidden” inside encapsulation headers. When the device type is set to **VXLAN**, Kentik uncovers this data using the following mechanism:

* **Header Export:** Kentik leverages sFlow’s ability to export the first 128-256 bytes of a packet header.
* **Underlay/Overlay Visibility:** By reading these headers, Kentik extends observability beyond traditional "underlay" traffic to include the "overlay" traffic traversing tunnels between servers or VMs.
* **Dimensional Decoding:** Information extracted from the sFlow headers is decoded into VXLAN-specific dimensions for filtering and grouping.

The diagram below shows how Kentik populates dimensions from relevant data in the packet header for a single VXLAN tunnel:

![For VXLAN devices, Kentik populates dimensions from the header of an encapsulated packet.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/UV-Vxlan_dimensions-1170h1105w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A26Z&se=2026-05-12T09%3A32%3A26Z&sr=c&sp=r&sig=CzTzHncD%2BrF2wxiciCI1tEdghCUbqrL331UbYENX4lM%3D)

*For VXLAN devices, Kentik populates dimensions from the header of an encapsulated packet.*

## sFlow Tunnel Decode Devices

While the standard VXLAN type stores data in custom fields, **sFlow Tunnel Decode** automates the promotion of encapsulated data (like inner IPs and Ports) into standard Kentik dimensions. This allows you to treat tunneled traffic like any other flow while maintaining tunnel context:

* **Standard Dimensions**: Tunneled traffic (inner packet) is represented by the standard traffic dimensions (see **sFlow Tunnel Decode Dimensions**).
* **Tunnel Dimensions**: VXLAN-specific data (VNI/VTEP) remains accessible in VXLAN-related dimensions (see **VXLAN Dimensions**).

The diagram below illustrates the mapping of the header information into Kentik dimensions in the sFlow tunnel decode category.

![Kentik's tunnel decoding uses standard dimensions for everything except VXLAN-specific data.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/UV-sFlow_tunnel_device-757h990w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A26Z&se=2026-05-12T09%3A32%3A26Z&sr=c&sp=r&sig=CzTzHncD%2BrF2wxiciCI1tEdghCUbqrL331UbYENX4lM%3D)

*Kentik's tunnel decoding uses standard dimensions for everything except VXLAN-specific data.*

## Future Implementations

Kentik is committed to expanding data center observability. If you have feedback on these VXLAN approaches or would like to see support for other tunneling protocols—such as **GRE/NVGRE** or **GENEVE**—please reach out via the **Contact Support** form in the portal's main navbar.
