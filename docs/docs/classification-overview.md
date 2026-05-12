# Classification Overview

Source: https://kb.kentik.com/docs/classification-overview

---

This article provides an overview of Kentik’s interface classification process, which automatically classifies the types of interfaces through which traffic enters and leaves your network.

> **Notes:**
>
> * To manage classification via the Kentik portal's Interface Classification page, see **Interface Classification**.
> * For interface classification use cases, see **Using Interface Classification**.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/AdmIC-UI-465h674w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A27Z&se=2026-05-12T09%3A37%3A27Z&sr=c&sp=r&sig=pHtYwA392OfqEi2TBmCp1xkIcS%2BpcPP9hUm9yy6SWRE%3D)

## About Interface Classification

Interface classification is the process of automatically identifying the role of every interface on your network. By assigning specific attributes to these interfaces, Kentik enables deeper insights into traffic patterns and prevents data inaccuracies like overcounting.

### Core Attributes

Each interface is classified using two primary attributes:

* **Network Boundary**: Defines whether an interface is **Internal** or **External** to your network.
* **Connectivity Type**: Defines the functional role of the interface (e.g., Transit, IX Peering, or Customer).

### Key Benefits

* **Dimensional Analysis**: Use classification values as source or destination dimensions in **Data Explorer**, **Dashboards**, or **Alerting**.
* **Filter Accuracy**: Easily filter traffic by connectivity type (e.g., "Show all Transit traffic") or network boundary.
* **API Integration**: Access all classification data programmatically via the Kentik **V5 Query API**

## How Classification Works

* **Rule Definition**: You create "if-match-then-classify" rules on the **Interface Classification Page** (**Settings »** **Interface Classification**).
* **Evaluation**: The rules engine analyzes **SNMP polling data**, matching interface descriptions or IP addresses against your defined criteria (see **Classification Logic Summary**)
* **Data Augmentation**: Once matched, the **Kentik Data Engine** (KDE) automatically adds these classification values to every flow record associated with that interface.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(24).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A27Z&se=2026-05-12T09%3A37%3A27Z&sr=c&sp=r&sig=pHtYwA392OfqEi2TBmCp1xkIcS%2BpcPP9hUm9yy6SWRE%3D)

*Classification of a given interface is stored in the device database and applied to flows across that interface.*

#### Classification Logic Summary

| Step | Process | Backend Action |
| --- | --- | --- |
| **1. Collection** | Flow records are collected from network interfaces. | Kentik ingests raw flow data (e.g., NetFlow). |
| **2. Lookup** | Interfaces are cross-referenced with the device database. | Kentik identifies the source and destination interfaces for each flow. |
| **3. Augmentation** | Flow records are "tagged" with classification attributes. | Values are added to the **Kentik Data Engine (KDE)**. |
| **4. Storage** | Attributes are saved into specific KDE dimensions | Data is stored in `src_interface_network_boundary`, `dst_interface_connectivity_type` |
| **5. Analysis** | Users query the augmented data. | Metrics are filtered or grouped by these attributes in **Data Explorer**. |

> **Note**: While classification is largely automated, you can perform manual modifications for interfaces that do not match existing rules.

## Network Boundary Attribute

The **Network Boundary** attribute identifies interfaces as **Internal** or **External**. This distinction allows technical decision-makers to accurately view traffic entering or exiting their network withouth the risk of over-counting.

#### Why it Matters: Preventing Over-counting

As shown in the diagram below, a single flow passing through your network may generate multiple flow records (at ingress, internal hops, and egress). By classifying boundaries, Kentik can isolate specific segments:

* **Ingress Counting:** Kentik filters for records from "External" interfaces where traffic is entering the network.
* **Egress Counting:** Kentik filters for records from "External" interfaces where traffic is leaving the network.

![Network boundary classification is designed to prevent overcounting of flows.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/AdmIC-Flows_diagram-497h749w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A27Z&se=2026-05-12T09%3A37%3A27Z&sr=c&sp=r&sig=pHtYwA392OfqEi2TBmCp1xkIcS%2BpcPP9hUm9yy6SWRE%3D)

*Network boundary classification prevents overcounting of flows.*

#### Configuration & Overrides

* **Automatic Assignment:** Boundary values are automatically assigned based on an interface's **Connectivity Type**.
* **Manual Overrides:** You can override default assignments within specific **Rule THEN Settings** or by changing the **IC Default Network Boundaries**.

Here’s an example visualization you can generate in Data Explorer once network boundary classification is implemented on an interface.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/AdmIC-Boundary_viz-674w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A27Z&se=2026-05-12T09%3A37%3A27Z&sr=c&sp=r&sig=pHtYwA392OfqEi2TBmCp1xkIcS%2BpcPP9hUm9yy6SWRE%3D)

## Connectivity Type Attribute

The **Connectivity Type** attribute classifies an interface based on its functional role, such as Transit, IX Peering, or Customer. This classification helps Kentik determine where traffic enters and exits your network, which is essential for optimizing costs and analyzing traffic patterns.

For the full list of connectivity types, see **Understanding Connectivity Types**.

### Discovery Methods

Kentik determines an interface's connectivity type by evaluating **SNMP OID Polling** data. It uses two primary methods for this evaluation:

#### 1. Interface Descriptions

Kentik scans the descriptions retrieved from your network devices via SNMP.

* **Manual Overrides**: If you have manually edited a description within the Kentik portal, the rules engine will prioritize the over the raw SNMP string (see **Edit an Interface**).

#### 2. Interface IP Addresses

Kentik can infer a connectivity type based on your network’s IP addressing policies. Examples include:

* **Customer Ranges**: IP blocks reserved exclusively for transit customers.
* **Infrastructure Ranges**: RFC1918 addresses used for internal CDN servers or management interfaces.
* **Exchange Points**: IP ranges corresponding to a specific Internet Exchange (IX) LAN.

> **Tip**: Using IP-based classification is often more reliable than descriptions if your naming conventions are inconsistent across different hardware vendors.

---

### Visualizing Connectivity Data

Once classified, you can use these attributes as dimensions in **Data Explorer** to generate visualizations like the one below, showing traffic volume broken down by role.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/AdmIC-Connectivity_viz-674w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A27Z&se=2026-05-12T09%3A37%3A27Z&sr=c&sp=r&sig=pHtYwA392OfqEi2TBmCp1xkIcS%2BpcPP9hUm9yy6SWRE%3D)

## Understanding Connectivity Types

Kentik’s supported connectivity types are described in the table below.

> **Note:** For more information, see **Applying Classification Rules**.

| Connectivity Type | Description | Diagram |
| --- | --- | --- |
| Aggregation Interconnect | Commonly found in hierarchical data center switching or routing topologies like Clos or leaf-and-spine, this type applies to bundles of interfaces between devices on separate traffic aggregation levels. |  |
| Available | The names or descriptions may indicate availability to facilitate network planning. Available interfaces might be provisioned with a transceiver depending on the company policy.  **Notes:**   * This type is not currently used and is not reserved for future use. * The network boundary of an available interface indicates if it’s for internal or external use. |  |
| Backbone | The generic connectivity type between two devices within the same network, regardless of their location within the data center. Users with separate backbone and data center teams usually use “datacenter interconnect” and “aggregation interconnect” categories within the data center, and “backbone” outside. |  |
| Cloud Interconnect | Used to connect a physical device to a Public Cloud. The term varies by cloud provider:   * **AWS**: **Direct Connect** * **Azure**: **ExpressRoute** * **GCP**: **Cloud Interconnect** * **OCI**: **FastConnect**   **Note:** If the interface accesses the cloud provider through a partner, it is **Virtual Cross Connect**. | Diagram illustrating cloud interconnects for various services in a network setup. |
| Customer | Typically points toward a downstream customer. The Border Gateway Protocol (BGP) session established through this interface usually announces a “Full view transit” or a “Partial Transit” to the downstream customer. A customer interface is the financial inverse of a transit (Full view) or paid private peering (Partial view) interface. | Diagram illustrating BGP routing options between networks and customer interface. |
| Data Center Fabric | Connects devices within a data center to other devices within the same data center (e.g., devices assigned Clos fabric roles like Super Spine, Spine, Leaf, or Top of Rack). | No Diagram |
| Datacenter Interconnect | Connects devices within a single data center. This is intended for cases where a data center team and a backbone team coexist, with one being a customer of the other. |  |
| DDoS / Cloud Mitigation | These two connectivity types classify interfaces in front of a DDoS mitigation platform (internal appliance-based or external scrubbing DDoS Mitigation Cloud provider). One instance will default Network Boundary to “Internal” and the latter will be “External”. |  |
| Embedded Cache | Connects to a Cache Appliance Server provided by a CDN or content provider (e.g., Facebook Appliance, Google Global Cache, Netflix Open Connect, Akamai Caches). These appliances are embedded in the ISP's last mile network to deliver popular traffic directly, bypassing the CDN or content provider’s network. Content on these appliances is refreshed via proxy feed during off-peak hours. | Diagram illustrating traffic flow between subscriber, network, and content provider server. |
| Free Private Peering | Connects to a non-transit AS using direct private connections to the AS's facing router. BGP peering typically involves routes from the directly connected AS (including a few downstream ASes).  **Note:** If either party charges the other, it becomes paid **Private Peering**. | Diagram illustrating paid peering interface and peer network routes between networks. |
| Host/Access | Indicates a direct connection between a router and a server host. |  |
| IX Peering (Fabric) | Connects to a public peering facility, a switching LAN where multiple ASes offer to peer. Unlike free or paid peering interfaces, IX peering allows multiple BGP sessions to multiple ASes through the same interface, resulting in multiple Next-Hop ASNs. | Diagram illustrating network peering with interfaces and connections between multiple networks. |
| IPX Interconnect | Supports various IP services, including voice (like VoLTE), data, and messaging, with mechanisms for traffic engineering to ensure high-quality and low-latency for real-time applications. It supports bilateral and multilateral interconnection. By default, the interfaces bearing these interconnections are considered “External”. They are the equivalent of an IX Connection in the world of packetized voice. |  |
| Management | Describes the port on a device connected to a Management network (the common network used to administer devices). The default connectivity type is “Internal”, but it can be set to “External” during out-of-band (OOB) monitoring. |  |
| Other | A “catch-all” type for interfaces where no other connectivity type applies. Alternatively, exclude interfaces from Interface Classification if there is no IP address or description. | No Diagram |
| Paid Private Peering | Similar to **Free Private Peering** but involves payment between parties. It supports cost-related business analytics. | Diagram illustrating paid peering interface and peer network routes between two networks. |
| Reserved | The interface is available but allocated for future use. Available interfaces might be provisioned with a transceiver depending on the company policy.  **Notes:**   * The network boundary of an available interface indicates if it’s for internal or external use. |  |
| SIP Interconnect | Also referred to as SIP trunk, a Voice over IP (VoIP) carrying interface. They are the equivalent of a BGP-enabled External interface and connect two networks to exchange VoIP traffic. Since these interfaces interconnect two separate networks, they are considered “External”. |  |
| Transit | Used to connect to upstream providers of Internet connectivity. They typically support BGP full routes. | Diagram illustrating BGP full routes between networks and transit interfaces. |
| Virtual Cross Connect | A third-party Layer 2 interface that enables point-to-point connectivity. Use cases include backbone and point-to-point interfaces, customer extranets, and indirect cloud interconnects. | Diagram illustrating virtual cross connect and peering fabric for network connections. |

## Troubleshooting Classification

If an interface appears as "Unclassified" or displays an incorrect connectivity type, use this checklist to identify and resolve common issues.

* ### Check SNMP Status

Kentik requires active SNMP polling to retrieve interface descriptions and IP addresses.

* **Verification**: Ensure the device is successfully communicating with Kentik via SNMP in **Networking Devices » Configure »** **SNMP Tab**.
* **Resolution**: If SNMP is failing, verify your community strings and ensure firewall rules allow traffic between the device and Kentik.

* ### Validate Rule Order

Classification rules are evaluated in order from top to bottom.

* **Issue**: A generic rule at the top of your list may be matching an interface before a more specific rule further down can be evaluated.
* **Resolution**: Go to the **Interface Classification Page** and move specific rules (like those for specific IP ranges) above more general description-based rules.

* ### Inspect Interface Descriptions

The rules engine relies on the descriptions retrieved from your hardware.

* **Issue**: The interface description on the device may have changed or might not match the regex pattern in your rule.
* **Resolution**: Check the current description in the **Interfaces** list. If the description is missing or incorrect, you can apply a **manual override** directly in the Kentik portal to force a match.

* ### Verify IP Address Coverage

If using IP-based classification, the interface must have an IP address that falls within your defined ranges.

* **Issue**: Interfaces without assigned IP addresses (common in some L2 deployments) cannot be classified using IP-based rules.
* **Resolution**: Switch to description-based rules or **manual classification** for these specific interfaces.

* ### Confirm Rule Application

Newly created rules or manual changes do not apply to historical data instantly.

* **Process**: After updating a rule, you must click **Save and Run**.
* **Wait Time**: It may take a few minutes for the **KDE** to begin reflecting the new classification values in **Data Explorer**.
