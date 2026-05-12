# General Metrics

Source: https://kb.kentik.com/docs/general-metrics

---

This article discusses general (i.e. flow) metrics used in Kentik queries.

> **Notes:**
>
> * For general metric information, see **About Metrics**.
> * For non-flow metrics, see **Non-flow Metrics**.
> * For device-specific metrics, see **Device-specific Metrics**.
> * Filter and group-by dimensions are covered in **Dimensions Reference** (see **Per-flow Metrics**).
> * Metrics are used for Kentik portal query settings and the **Query API**.

## About General Metrics

In Kentik, general metrics are transmitted via flow protocols like NetFlow or sFlow and aren’t tied to specific devices. Their availability in settings like the **Metrics Pane** of the Query sidebar in Data Explorer depends on the device category (Router or Host) and the specific device type (see **Supported Device Types**).

## Metrics from All Devices

When a query includes traffic solely from data sources (see **About Data Sources**) in the "Routers" category, only the following metrics are available:

| Metric in portal | Variations | Calculated as... | Derived from KDE column(s) |
| --- | --- | --- | --- |
| Bits per second | Sampled at:   * Ingress and Egress * Ingress * Egress | * Average * 95th Percentile * 99th Percentile * Max * Total | in\_bytes,  out\_bytes |
| Packets per second | Sampled at:   * Ingress and Egress * Ingress * Egress | * Average * 95th Percentile * 99th Percentile * Max * Total | in\_pkts,  out\_pkts |
| Flows per second | N.A. | * Average * 95th Percentile * 99th Percentile * Max * Total | Based on rows per KDE main table |
| Source IPs | * Unique Count * Bitrate per IP | * Average * 95th Percentile * 99th Percentile * Max * Total | inet\_src\_addr |
| Destination IPs | * Unique Count * Bitrate per IP | * Average * 95th Percentile * 99th Percentile * Max * Total | inet\_dst\_addr |
| Unique Route Prefixes | * Source * Destination | * Average * 95th Percentile * 99th Percentile * Max * Total | inet\_src\_route\_prefix, inet\_dst\_route\_prefix |
| Unique Ports | * Source * Destination | * Average * 95th Percentile * 99th Percentile * Max * Total | inet\_src\_route\_prefix, inet\_dst\_route\_prefix |
| Unique ASNs | * Source * Destination * Next Hop Destination | * Average * 95th Percentile * 99th Percentile * Max * Total | src\_as, dst\_as, i\_dst\_nexthop\_as\_name |
| Unique Countries | * Source * Destination | * Average * 95th Percentile * 99th Percentile * Max * Total | src\_geo, dst\_geo |
| Unique Regions | * Source * Destination | * Average * 95th Percentile * 99th Percentile * Max * Total | src\_geo\_region, dst\_geo\_region |
| Unique Cities | * Source * Destination | * Average * 95th Percentile * 99th Percentile * Max * Total | src\_geo\_city, dst\_geo\_city |
| Sample rate | * Max * Average | * Average * 95th Percentile * 99th Percentile * Max * Total | sample\_rate |

> **Notes:**
>
> * The “Total” metric is available only when the chart type is set to Table or Matrix (see **Chart Visualization Types**).
> * Metrics labeled as "unique" are evaluated across all instances on queried devices within each individual time slice of the query's time range, but not between different time slices (see **Table Time-slicing**).

## Host Traffic Metrics

When querying traffic that includes devices in the "Host" category (using Kentik’s host agent, see **About the Universal Agent**), the data from these hosts is stored in KDE.

### Host Metrics by Protocol

The availability of a given metric varies by protocol:

| Metric in portal | TCP | HTTP | DNS | Calculated as... | Derived from KDE column(s) |
| --- | --- | --- | --- | --- | --- |
| Retransmits | * Per second * Repeated per second * Percent * Percent Repeated | * Repeated Retransmits * Retransmitted Packets Out | N.A. | * Average * Percentile (95th or 98th) * Max * Total | retransmitted\_out\_pkts, repeated\_retransmits, both\_pkts (for %) |
| Out-of-order packets | * Per second * Percent | * Per second (In) | N.A. | * Average * Percentile (95th or 98th) * Max * Total | ooorder\_in\_pkts, ooorder\_out\_pkts, both\_pkts (for %) |
| Fragments | * Per second * Percent | * Per second | * Per second | * Average * Percentile (95th or 98th) * Max * Total | fragments, both\_pkts (for %) |
| Zero Windows | * Count * Percent | * Per second | N.A. | * Average * Percentile (95th or 98th) * Max | zero\_windows |
| Receive Window | N.A. | N.A. | N.A. | * Average * Percentile (95th or 98th) * Max | receive\_window |
| Latency | * Client * Server * Application * First Payload Exchange | * Client * Server * Application * FPEX | * Application | * Average * Percentile (95th or 98th) * Max | client\_nw\_latency\_ms, server\_nw\_latency\_ms, appl\_latency\_ms, fpex\_latency\_ms |

> **Note:** The “Total” metric is available only when the chart type is set to Table or Matrix (see **Chart Visualization Types**).

### Host Metrics Descriptions

The following descriptions apply to the host-only metrics listed in the table above:

* **Retransmits**: Packets re-sent from source to destination, applicable to reliable protocols like TCP.

  + Indicates network delivery issues affecting performance.
  + Measured per second and  as a percentage of all packets sent.

    > **Note:** "Repeated" refers to packets retransmitted 3 or more times.
* **Out of order**: Packets arriving out of sequence.

  + High values suggest variability in delivery paths, impacting latency-sensitive traffic like real-time audio/video.
  + Measured per second and  as a percentage of all packets sent.
* **Fragments/second**: Packets split into smaller packets for network delivery.

  + Fragmentation can increase CPU load and cause retransmits if fragments arrive out-of-order.
  + Measured per second and as a percentage of all packets sent.
* **Zero Windows**: Count of TCP receive windows with a value of zero, indicating a full buffer.
* **Receive Window**: Size of the TCP receive window.
* **Latency**: Various measures include:

  + RTT/2 client latency (derived): One-way latency from the client perspective. High values indicate network or server-side issues
  + RTT/2 server latency (derived): One-way latency from the server perspective. High values indicate network or client-side issues.
  + RTT/2 application latency (derived): Derived from request/response pairs at the application layer, indicating end-user experience latency. Effective for protocols with clear request/response pairings.
  + First Payload Exchange Latency: Measures application response time when the protocol isn't understood or can’t be decoded (e.g., HTTPS, SQL). Excludes TCP setup, starting with the first packet sent and ending with the first packet returned.

### Host Metrics KDE Columns

The following table shows the correspondence between the metrics described above and the columns of the Kentik Data Engine (KDE):

| Metric in portal | Type: value column | Derived from KDE column(s) |
| --- | --- | --- |
| Retransmits | bigint Native | retransmitted\_out\_pkts, repeated\_retransmits, both\_pkts (for %) |
| Out-of-order packets | bigint Native | ooorder\_in\_pkts, ooorder\_out\_pkts, both\_pkts (for %) |
| Fragments | bigint Native | fragments, both\_pkts (for %) |
| Zero Windows | bigint Native | zero\_windows |
| Receive Window | bigint Native | receive\_window |
| Latency | bigint Native | client\_nw\_latency\_ms, server\_nw\_latency\_ms, appl\_latency\_ms, fpex\_latency\_ms |

## Application Decodes Metrics

When using Kentik's software host agent for application decodes (see **About Application Decodes**), both dimensions and metrics are generated and stored in or derived from KDE columns. The metrics available for queries via the metric selector (see **Metrics Dialog UI**) include:

| Metric in portal | Description | Type: value column |
| --- | --- | --- |
| Connection Name | TCP connection ID. | bigint UDR |
| Application Latency | One-way network latency derived by examining request/response pairs at the application layer. | bigint UDR |
| FPEX Latency | Elapsed time from first packet sent to first packet returned. | bigint UDR |

> **Note:** For dimensions related to application decodes, see **Application Decodes Dimensions**.

#### Legacy Application Decodes Metrics

The "legacy" metrics in the table below are from older versions (<v1.3.0) of Kentik’s legacy host software agent, `kprobe`:

| Metric in portal | Description | Type: value column | KDE name(s) |
| --- | --- | --- | --- |
| Connection ID | TCP connection ID.  **Note:** Superseded by Connection Name. | bigint Native | connection\_id |
| Application Latency (ms) | One-way network latency derived by examining request/response pairs at the application layer. | bigint Native | appl\_latency\_ms |
| First Payload Exchange Latency (ms) | Elapsed time from first packet sent to first packet returned.  **Note:** Superseded by FPEX Latency. | bigint Native | fpex\_latency\_ms |
