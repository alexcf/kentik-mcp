# Metrics & Dimensions

Source: https://kb.kentik.com/docs/metrics-dimensions

---

Metrics and dimensions are collected or derived from network data (flow, SNMP, BGP, etc.), stored in the Kentik Data Engine (KDE), and used in the portal (and Kentik APIs) as group-by and filter parameters in Kentik queries.

## About Metrics

A metric is a combination of a unit (e.g., a bit) with a method of calculation (e.g., average per second) to create a quantifiable measurement (average bits/second). In Kentik, metrics represent measurements made on flows, which are used for counts, rankings (e.g., in a top-X list), and thresholds (e.g., in alerting). Metrics are available in several portal locations (see **Metrics in the Portal**).

> **Notes:**
>
> * For metrics that aren't transmitted via a flow protocol, see **Non-Flow Metrics**.
> * For metrics that are device-specific, see **Device-Specific Metrics**.
> * Some metrics that are stored as part of the flow records in KDE may be used as dimensions, see **Per-Flow Metrics**.
> * In addition to being used for query settings in the Kentik portal, metrics and dimensions are also used for the **Query API**.

## Metrics in the Portal

Here are some places in the Kentik portal where metrics are specified and used:

| Portal Location | Primary Metric | Secondary Metric |
| --- | --- | --- |
| Library » Dashboards: **Add View Panel** or **Edit** **View Panel** dialog | Query tab » Metrics pane, either:   * **Metrics** dropdown * Customize Metrics button » **Metrics** dialog (see **Metrics Pane**) | Query tab » Metrics pane » Customize Metrics » **Metrics** dialog |
| Data Explorer: Query sidebar, **Metrics** pane | Query sidebar » Metrics pane, either:   * **Metrics** dropdown * Customize Metrics button » **Metrics** dialog (see **Metrics Pane**) | Query sidebar » Metrics pane » Customize Metrics » **Metrics** dialog |
| Alerting » Policies: **Add Policy** or **Edit Policy** dialog | Dataset tab » Data Funneling pane » **Primary Metric** drop-down | Dataset tab » Data Funneling pane » **Secondary Metric** field |

## Metrics Reference

This section details metrics for Kentik queries. Metrics generally apply to traffic from all device types (e.g., routers, hosts, see **Supported Device Types**), with some specific to host agents like Universal Agent (see **About the Universal Agent**).

#### Metric Categories

Metrics in **KDE Tables** fall into these categories:

| Category | Description | Agent |
| --- | --- | --- |
| **Metrics from All Devices** | Available from routers and hosts (see **Supported Device Types**). | None or Universal |
| **Host Traffic Metrics** | Available only from hosts (see **About the Universal Agent**). | Universal |
| **Application Decodes Metrics** | From application decodes like DNS and HTTP (see **About Application Decodes**). | Universal |
| **SNMP Metrics** | From SNMP polling (see **SNMP OID Polling**). | None or Universal |
| **Streaming Telemetry Metrics** | From Streaming Telemetry (see **Streaming Telemetry Device Support**). | None or Universal |
| **Device-Specific Metrics** | For certain device types only (physical or virtual). | None or Universal |

> **Notes*:*** Use the Category links above for specific metrics lists.

## About Dimensions

In Kentik, dimensions represent specific flow data (see **Flow Overview**), sourced directly from flow records (e.g., NetFlow, sFlow), correlated sources (e.g., GeoIP, threat feeds), or derived by Kentik. Each dimension corresponds to an actual or derived column in the **KDE Tables**.

## Dimensions in the Portal

Dimensions in the Kentik portal and the **Query API** are used for:

* **Group-by dimensions:** Chosen via **Dimension Selectors** in the **Dimensions** pane (e.g., Query sidebar in Data Explorer).
* **Filters:** Selected in the **Filters Dialog** in the **Filters** pane (e.g., Query sidebar in Data Explorer).

#### Dimension Locations

The following table shows more specifically the various locations in the Kentik portal where dimensions are specified and used:

| Portal Location | Group-By | Filters |
| --- | --- | --- |
| **Dashboards** | Add View Panel or Edit View Panel dialog » Query tab » **Dimensions** pane (see **Panel Dialogs**) | Query sidebar » Filters pane » Filters dialog (via Edit Filters button) » **Add Ad-Hoc Filter** (see **Filter Groups UI**). |
| **Data Explorer** | Query sidebar » Dimensions pane » **Group-by Dimensions** dialog (see **Dimension Selectors**) | Query sidebar » Filters pane » Filters dialog (via Edit Filters button) » Add Ad-Hoc Filter (see **Filter Groups UI**) |
| **Alerting » Policies** | Add Policy or Edit Policy dialog » Dataset tab » Data Funneling pane » **Dimensions** | Add Policy or Edit Policy dialog » Dataset tab » Data Funneling pane » **Filters** |
| **Admin » User** | N.A. | Add User or Edit User dialog » **User Specific Filters** pane (see **User Settings Dialogs**) |
| **Admin » Saved Filters** | – | Add Saved Filter or Edit Saved Filter dialog » **Ad-Hoc Filter Groups** pane (see **Saved Filter Admin Dialogs**) |

## Dimensions Reference

This section details the dimensions available for group-by and filtering in Kentik queries. Generally, these dimensions can be used in queries involving traffic from all device types, e.g., routers and hosts (see **Supported Device Types**). However, some dimensions are specific to traffic from host agents like Kentik’s **Universal Agents**.

#### Dimension Categories

Dimensions, which each represent an actual or derived column in the tables of the KDE (see **KDE Tables**) fall into the following functional categories:

| Category | Description | Requires Host Agent |
| --- | --- | --- |
| **Network and Traffic Topology** | Filter/group-by device info like interface names, descriptions, port IDs. | No |
| **IP and BGP Routing** | Filter/group-by IPs, protocols, TCP flags, ToS, routing info (AS, paths, names, etc.), and per-flow metrics. | No |
| **Cloud Dimensions** | Filter/group-by VPC flow log fields from cloud providers. | No |
| **Geolocation Dimensions** | Filter/group-by physical location properties (country codes, city names, etc.). | No |
| **Application Context and Security** | Filter/group-by context factors (e.g., CDN origin/termination, service type) and security threats. | No |
| **Application Decodes** | Data on DNS lookups and HTTP (domain name, referrer, status).  **Note**: DNS decodes require Kentik’s Universal Agent with **DNS OTT Tap** capability. | **Yes** |
| **Container Networking Dimensions** | Related to Kubernetes.  **Note**: Use of Kubernetes with Kentik requires a special software agent; contact Customer Success for further information. | – |
| **Device Metrics Dimensions** | Filter/top-X evaluations based on device metrics (e.g., SNMP, Streaming Telemetry). | – |
| **MPLS Dimensions** | Related to Multiprotocol Label Switching. | – |
| **Device-Specific Dimensions** | Filter/group-by fields in flow records from specific devices (e.g., Palo Alto Networks firewalls or Cisco ASA). | No |

> **Notes:**
>
> * Use the Category links above for specific dimensions lists.
> * For more about dimensions requiring a host agent, see **Host Traffic Dimensions**.

## Host Traffic Dimensions

Certain dimensions are exclusive to traffic from the Universal Agent, Kentik's software host agent (see **About the Universal Agent**):

> **Note**: The following DNS dimensions specifically require the Universal Agent to be running with DNS OTT Tap capability.

* **DNS Query**: Translates domain names to IPv4 or IPv6 addresses.
* **DNS Query Type**: Specifies the resource record type requested (see **List of DNS Record Types**).
* **DNS Response**: Includes resource records (RRs) like:

  + **A**: IPv4 address for given host.
  + **AAAA**: IPv6 address for given host.
  + **CNAME**: Domain name used to resolve the original DNS query.
  + **PTR**: Look up a domain name based on an IP address.
  + **MX**: Mail exchange server for a DNS domain name.
  + **NS**: Authoritative name server for given host.
  + **TXT**: Non-formatted text string typically used by Sender Policy Framework (SPF) to prevent the sending of emails using a fake identity.
* **DNS Return Code**: Status code from a DNS query (see **DNS Parameters**).
* **HTTP Host Header**: Indicates the server's domain name.
* **HTTP Referrer**: Indicates the source address of a webpage request.
* **HTTP Return Code**: HTTP status codes (see **HTTP Status Codes**).
* **HTTP User Agent**: Identifies the client making the request.
* **HTTP URL**: Filename portion of a web resource path, including query string if present.

> **Note:** Kentik supports substring matching in certain host-sourced group-by dimensions (see **DNS/WWW Extract Function**).
