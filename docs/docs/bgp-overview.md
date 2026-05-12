# BGP Overview

Source: https://kb.kentik.com/docs/bgp-overview

---

This article provides an overview of Kentik’s BGP implementation.

## About BGP

Border Gateway Protocol (BGP) is a Layer 4 protocol that facilitates the interconnection of autonomous systems (AS) to route traffic across the Internet. It uses an AS Path attribute to provide routing information. A path is a list of AS Numbers (ASNs) that define a route to a destination network. BGP-enabled routers exchange routing information over TCP (port 179) when configured to peer. Routers prevent looping by not importing routes that contain their own ASN.

> **Note:** The current standard for BGP is defined in **RFC 4271**.

### Kentik BGP Collection

Kentik collects BGP routing information and associates it with individual flows as they are ingested into the Kentik Data Engine (KDE) datastore. This enables queries that encompass flow routing.

As KDE receives flow data (e.g., NetFlow), it correlates src/dst IP flow data to BGP prefix information to extract source and destination AS Path and community information per flow.

This information is stored in the KDE:

* For each KDE row, the `src_as` and `dst_as` values are calculated from the routing table of the router sending the flow.

  > **IMPORTANT**: If local device routing data is unavailable, these values are taken from a Kentik-provided global (generic) routing table.
* For each KDE row, the source and destination `_bgp_aspath`, `_bgp_community`, and `_nexthop_as` fields are derived from the route associated with the src and dst IPs.

### Kentik BGP Use Cases

Kentik customers interested in these use cases can benefit from the correlation of flow and BGP information:

* Creating tags based on BGP routing
* Filtering query results for more customer-specific insights
* BGP-correlated traffic analysis
* Sophisticated peering analytics (see **Discover Peers**)

## BGP in the Portal

The following sections cover the BGP features in the Kentik portal.

### BGP in Device List

The **Device List** on the portal's Devices page provides BGP activity information for each listed device:

* **BGP Enabled**: Indicates whether BGP Peering is enabled (checkmark icon) or not (no icon) in the device settings.
* **BGP Routes last 5m**: Number of routes in the BGP routing table (last 5 minutes).
* **BGP updates last 24h**: Number of routing table updates in last 24 hours.

### BGP Tagging

Kentik collects BGP data and assigns tags to incoming flows that match communities or AS Paths (see **Tag BGP Matching**). Tags are applied separately for source and destination when a flow is ingested into the KDE (see **Flow Overview**). You can use tags to narrow query results by applying them using src and dst filter functions in the Data Explorer and Query Builder pages of the Kentik portal.

> **Note**: A tag is applied only to flows arriving after it was created.

For AS Path tags and BGP community tags, matches are made on substrings:

* **BGP AS Path Tags**: A subset of standard regex is supported, meaning that "\_10\_" matches paths that include ASN 10. Tags where `as-path` is specified as "\_10 100\_" are also allowed.
* **BGP Community Tags**: Also support periods, allowing you to specify, for example, "2000:1...." to find flows with community “2000:1xxxx”.

The following table defines the regex special characters supported by Kentik for matching BGP AS Paths and Communities during the tagging process:

| Special Character | Matches… |
| --- | --- |
| \_  (underscore) | * Start of string * End of string * "" (space) |
| . | Any single character, including white space |
| [ ] | Characters, or ranges of characters, separated by a hyphen and contained in square brackets. |
| ' | Character or null strings at the beginning of an input string. |
| ? | Zero or one occurrence of the pattern containing the question mark. |
| $ | End of string |
| \* | Zero or more sequences of the preceding character. Also acts as a wildcard for matching any number of characters. |
| + | One or more sequences of the preceding character. |
| () | (Used for nesting expressions) |

> **Note:** Spaces at the start and end of the input field, and before and after each comma, are removed for BGP community and AS path tags.

### BGP Grouping and Filtering

BGP data can be used for grouping and filtering in the Kentik Portal’s **Data Explorer**:

* **BGP Grouping**: Enable BGP grouping (SQL `GROUP BY` clauses) with the BGP-related options on the **Group** **by Metric** dropdown.
* **BGP Filtering**: Enable BGP filtering (SQL `WHERE` clauses) in ANDed or ORed groups with the **Filters** sidebar.

Kentik supports grouping and filtering by the following BGP dimensions (source and destination path):

* **Prefix**: Route Prefix, Route LEN
* **Pathing**: BGP AS\_PATH, BGP Community
* **Hops**: Next Hop IP/CIDR, Next Hop AS (Number/Name), 2nd & 3rd BGP\_Hop AS (Number/Name)

## BGP Analysis

Kentik’s BGP data can be used for sophisticated analytics in the Kentik Portal, including peering analysis. This analysis shows diagrams and tables of the Autonomous Systems that traffic leaving your network passes through on the way to its destination (see **Discover Peers**).
