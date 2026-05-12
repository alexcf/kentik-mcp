# Kentik Data Engine

Source: https://kb.kentik.com/docs/kde-overview

---

This article discusses querying the Kentik Data Engine (KDE), a datastore for storing current and historical flow data.

> **Notes:**
>
> * For information about the KDE main tables, including the columns from which Kentik Detect derives filters and group-by dimensions, see **KDE Tables**.
> * As of May 1st, 2025, the **Query SQL Method** has been deprecated and is no longer supported.

## About Kentik Data Engine

Kentik Data Engine, a custom-built scalable columnar datastore, ingests flow data from customers (e.g., NetFlow or sFlow), correlates it with GeoIP and BGP data, and stores it in flow records. These records can be queried from the Kentik portal or via Kentik APIs (see **Query API**).  
  
KDE maintains separate databases for each customer’s flow records. These databases contain “Main Tables” for flow records and associated data for each device (router or host), as well as supplemental data derived from the flow data. The columns in these Main Tables contain the data queried by Kentik, most of which are dimensions used for filtering and grouping in queries. For a categorized list of these dimensions, see **Dimensions Reference**.

## Subqueries, Slices, and Shards

> **Note:** This topic discusses Kentik's implementation of SQL as it relates to Kentik's querying of the Kentik Data Engine. Querying via the **Query SQL Method** of the V5 Query API and direct user querying of KDE using SQL are both **deprecated**. For additional information please contact Kentik support.

Kentik’s implementation of SQL for querying the Kentik Data Engine uses subqueries, slices, shards, masters, and workers.

Kentik Data Engine presents as a traditional SQL database but uses a scalable distributed cluster architecture.

When customer flow records and data are ingested into KDE, each device’s data is stored in two parallel “main tables”: one for a full resolution dataseries and another for a “fast” dataseries optimized for faster queries covering long timespans (see **KDE Tables**and **Resolution Overview**).

Each KDE main table is divided into logical “slices” of one minute (for Full dataseries) or one hour (for Fast dataseries). Every slice from a given device’s main table corresponds to at least two identical “shards,” each an on-disk physical block in KDE. For high availability, shards for a given slice are stored on different machines (typically in different racks) referred to as “workers.”

The KDE subquery process involves workers and “masters.” Workers handle specific subqueries by having corresponding shards on their machines. Masters split queries into slices, identify workers with access to shards, and assign subqueries to workers. Workers generate result sets from shard data and return them to masters, which reassemble them into final results returned to the PostgreSQL interface as SQL/MED foreign data (see **https://en.wikipedia.org/wiki/SQL/MED**). This approach enables Data Engine to achieve exceptional performance with a SQL interface, even with large datasets.

The fact that queries to Data Engine will be broken into subqueries means that there are some special functions required to properly interact with the underlying architecture. These functions can affect the way that queries are written in the following situations:

* Queries involving aggregation.
* Queries involving CIDR

> **Note:** Kentik Data Engine does not support the use of NOT IN for a WHERE clause.

## KDE Resolution

Kentik creates two independent dataseries at ingest: one at full resolution and another optimized for faster queries spanning long timespans.

* **Full dataseries**: Includes all flow records sent by a customer to Kentik (within service agreement limits).
* **Fast dataseries**: Includes a subset of flow records, enabling faster response to queries spanning 24 hours or more.

Parallel dataseries allows long-timespan queries to return in seconds without compromising detail for shorter timespans. For more information, see **Resolution Overview**.

## Table Time-slicing

Each Main Table and All Devices table in KDE is subdivided into time-slices for distributed processing of subqueries (see **Subqueries, Slices, and Shards**):

* For tables in the Full dataseries (see **KDE Resolution**), ingesting flow data into a main table results in one row per flow record, with time-slices of one, five, or 10 minutes based on query width.
* For Fast dataseries tables, time-slices are one hour (see **About Dataseries Resolution**).

## KDE Query Efficiency

KDE stores network data and supports multiple ways to find answers. Some querying approaches are more efficient than others. Here are tips to ensure efficient queries:

* **Use the device selector**: Use the **Selected Devices Dialog** to choose devices (individually, by site, or by label) from which to return results. This is more efficient than choosing All Devices and filtering. Only use All Devices when you want all devices.
* **Favor native dimensions**: Native dimensions are more efficient for KDE to process than virtual dimensions. Check if a dimension is native or virtual in **Dimensions Reference**.
* **Apply tags for recurring use cases**: KDE columns for **Flow Tags** and **Custom Dimensions** are populated at ingest and are native. Use tags or custom dimensions to narrow results for repeated queries (e.g., dashboards, saved views, alert policies) instead of filters (applied at query time).

## Populating BGP Fields

When a flow is ingested into a device's KDE main table, the way the flow record's BGP fields are populated depends on the following three conditions at time-of-ingest:

1. Is the device enabled for peering with Kentik?
2. Does the BGP routing table obtained by peering with the device include an IP route for the received flow?
3. Is the AS path (list of ASNs) in the IP route empty or not-empty?

The following table shows how the conditions above determine the values written to the BGP-related fields in a given device's main Data Engine table:

| Conditions: | State A | State B | State C | State D |
| --- | --- | --- | --- | --- |
| 1. Peering | Enabled | Enabled | Enabled | Not enabled |
| 2. IP route for flow | Yes | Yes | No | N.A. |
| 3. AS path | Non-empty string | Empty string | N.A. | N.A. |
| **Fields:** | | | | |
| src\_bgp\_aspath dst\_bgp\_aspath | The BGP AS path in the route. | The ASN of the peering device. | Empty string | Empty string |
| src\_as dst\_as | The last ASN in BGP AS path. | The ASN of the peering device. | The ASN of the corresponding (dst or src) IP address. | The ASN of the corresponding (dst or src) IP address. |
| i\_src\_as\_name i\_dst\_as\_name | The AS name corresponding to the last ASN in BGP AS path. | The AS name corresponding to the ASN of the peering device. | The AS name corresponding to the (dst or src) ASN. | The AS name corresponding to the (dst or src) ASN. |
| src\_nexthop\_as dst\_nexthop\_as | The first ASN in BGP AS path. | The ASN of the peering device. | 0 | 0 |
| i\_src\_nexthop\_as\_name i\_dst\_nexthop\_as\_name | The name corresponding to the first ASN in BGP AS path. | The AS name corresponding to the ASN of the peering device. | Empty string | Empty string |
| src\_second\_asn dst\_second\_asn | The ASN in the second position of the BGP AS path. | 0 | 0 | 0 |
| i\_src\_second\_asn\_name i\_dst\_second\_asn\_name | The name corresponding to the ASN in the second position of the BGP AS path. | Empty string | Empty string | Empty string |
| src\_third\_asn dst\_third\_asn | The ASN in the third position of the BGP AS path. | 0 | 0 | 0 |
| i\_src\_third\_asn\_name i\_dst\_third\_asn\_name | The name corresponding to the ASN in the third position of the BGP AS path. | Empty string | Empty string | Empty string |
| src\_bgp\_community dst\_bgp\_community | BGP community of the router. | BGP community of the route. | Empty string | Empty string |

## KDE Access via PostgreSQL

> **Note:** PostgreSQL (psql) access to KDE is now **deprecated**. For additional information please contact Kentik support.
