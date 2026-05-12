# Using Kentik Firehose

Source: https://kb.kentik.com/docs/using-kentik-firehose

---

This article discusses the use of Kentik Firehose.

## Firehose Overview

Kentik, built on the Kentik Data Engine (KDE), collects flow records (e.g., NetFlow, sFlow) from data sources (network devices, hosts, clouds) and enriches them with additional data (SNMP, GeoIP, BGP, threat feeds) before storing them in a time-series database. Kentik Firehose, and its accompanying software agent `ktranslate`, allows copies of these enriched flow records to be sent to a chosen location, for example, into other analytics systems (directly or via a data lake).

> **Note:** The Kentik portal does not support the Firehose capability for cloud data sources. To use `ktranslate` for flow log data from cloud resources, contact Kentik (see **Customer Care**).

### `ktranslate` Overview

To receive copies of the enriched flow records, you deploy a Kentik-provided software agent (client binary) called `ktranslate` in a public cloud resource or on a server in your data center:

* Data is sent to `ktranslate` in Kentik's normalized `kflow` format, and can be converted by `ktranslate` into a desired output format (see **ktranslate Operations**).

  > **Note:** `ktranslate` does not input HTTPS. Use a proxy to convert to HTTP before passing encrypted `kflow` to `ktranslate` (HTTPS » [proxy] » HTTP » [`ktranslate`]).
* Kentik-registered data sources (e.g., router or host) can send data to one or more endpoints (URLs), each representing a single instance of `ktranslate`. To specify endpoints for a data source, use the **Integrations** tab of the Device dialog (see **Firehose Data Sources**), accessed in the portal at Settings » **Devices**.
* Each `ktranslate` instance can receive flow data from multiple data sources.
* You can deploy and run as many instances of `ktranslate` as you like.
* Kentik recommends one instance per 100k flows per second ingested from data sources to `ktranslate`.

### `ktranslate` Scenarios

`ktranslate` facilitates data transfer between network entities by listening for HTTP traffic from a source, processing it based on CLI arguments (see **ktranslate Command Line**), and forwarding it to one or more destinations.

**Primary Use Case: Business Intelligence**

The most common deployment involves sending enriched `kflow` from the KDE to a third-party analytics system or data lake. In this scenario, `ktranslate` formats the data and directs the resulting stream to a specific "sink" (destination) required by the external system.

### `ktranslate` Operations

> **Notes:**
>
> * The operations/options below are for `ktranslate` version 2.
> * For CLI examples, see **ktranslate CLI Examples**.

`ktranslate` can perform various operations on incoming data before passing it to an output stream or file, categorized as follows:

* **Formats**: Specify the output format for `kflow` input. Options include JSON (default), NetFlow, Apache Avro, InfluxDB line protocol, Splunk, and Prometheus.

  > **Note:** Rollups are not supported with NetFlow format.
* **Compression**: Apply a compression algorithm to the output data. Options include none (default), gzip, snappy, null, and deflate.
* **Sinks**: Set the destinations for the output data streams. Options include stdout (default), Kafka, New Relic, Kentik, net, HTTP, Splunk, Prometheus, and file.

  > **Note**: You can specify multiple sinks for the output of a single `ktranslate` instance.

  + **Rollups**: Group data into dimension and metric sets using rollups. Specify the rollup method, name, metric, and dimensions as a comma-separated string.

    - **Example**: `sum` is the rollup method, `kentik.bytes` is the name, `in_bytes` is the metric, and `dst_addr` is dimension 1 (of 1):  
      `-rollups sum,kentik.bytes,in_bytes,dst_addr`

      > **Notes:**
      >
      > * Supported rollup methods are sum, min, max, mean, median, entropy, percentilerank, percentile, and unique.
      > * Output may contain multiple rollups.
      > * The default time-slice width is 15 seconds, but you can set it with the argument `-rollup_interval`.
      > * Additional rollup variations can be specified with `-rollup_` arguments (see **ktranslate CLI Reference**).
      > * Metrics and dimensions supported for rollups are listed in **ktranslate Metrics and Dimensions**.
* **Filters**: Focus the `ktranslate` output data by specifying filters. A filter is specified as a comma-separated string with four attributes: type, dimension, operator, and value.

  + **Example**:  `-filters string,src_addr,==,12.0.1.2`

    > **Notes:**
    >
    > - Supported filter types are string, int, and addr.
    > - For supported filter dimensions, see **ktranslate Metrics and Dimensions**.

## Firehose Setup Overview

Setting up Firehose is a two-stage process:

* Deploy `ktranslate` in a Docker container on a server in your data center or in a public cloud resource, as detailed in **ktranslate Download and Install**:

  + The server must have a public IP address and meet the requirements given in **ktranslate Requirements**.
  + The `ktranslate` CLI arguments (see **ktranslate CLI Reference**) in the Docker `run` command will determine which operations this instance of `ktranslate` performs and the destinations to which it outputs streams. Some typical deployment options are covered in **ktranslate CLI Examples**.
* In the Kentik portal, set the deployed `ktranslate` instance as the destination of a T-ed stream from KDE that will contain the `kflow` for a given Kentik-registered data source (see **Firehose Data Sources**). Repeat as needed to cover all of the data sources for which you wish to receive `kflow` on this `ktranslate` instance.

> **Tip:** By repeating the process above, you may send `kflow` to multiple instances of `ktranslate`. You may deploy as many instances as needed to route the data for your specific purposes.

## `ktranslate` Deployment

Deployment of `ktranslate` is covered here.

### `ktranslate` Requirements

The following resources must be available (at minimum) to support the use of `ktranslate` running in a container:

* **RAM allocation**: 1GB
* **Core**: One Intel Xeon CPU E5-2676 v3 @ 2.40GHz
* **AWS VM Spec** (EC2 instance type): t2.micro (**https://aws.amazon.com/ec2/instance-types/**)

### `ktranslate` Download and Install

Installation from a downloaded Docker image provides a convenient and easy deployment mechanism for systems that already use Docker-based containerized applications.  
  
 To install `ktranslate` via Docker:

1. Pull down the latest `ktranslate` image from Kentik's Docker hub repository (**https://hub.docker.com/r/kentik/ktranslate**):  
   `# docker pull kentik/ktranslate:v2`
2. Run the Docker image as shown in this example:  
    `# docker run -p 8082:8082 kentik/ktranslate:v2`
3. In the portal, the data sources from which you wish to send `kflow` to this `ktranslate` instance must be provided with the instance's endpoint (see **Firehose Data Sources**).

> **Notes:**
>
> * To install and run the latest version of `ktranslate`, leave `:v2` off of the end of the above examples.
> * In the above `docker run` command (step 2) the `-p` argument maps an external port to a port in the container.
> * For a reference to the `ktranslate` arguments that may be used in the above `run` command, see **ktranslate CLI Reference**.
> * Example command lines for common use cases are provided in **ktranslate CLI Examples**.

## `ktranslate` Command Line

The command line for `ktranslate` is covered here.

### `ktranslate` CLI Examples

The topics below provide examples of the `ktranslate` arguments for common `docker run` commands.

#### `ktranslate` Operation Examples

The following examples illustrate commands that are used for some common `ktranslate` operations:

* **Filtered output**: Translate `kflow` to JSON and output one `stdout` stream that is filtered for traffic on source port 80 only.

  ```
  docker run -p 8082:8082 kentik/ktranslate:v2 -format json -sinks stdout -filters int,l4_src_port,==,80 -listen=0.0.0.0:8082
  ```

  > **Note**: Multiple filters may be used. All filters are ANDed.
* **Output with rollups**: Translate `kflow` to JSON and output two rollups to an `stdout` sink. One rollup aggregates the top 10 unique destination and source IPs every 60 seconds. The second represents the top source ports by bytes, aggregated at the default 15 second interval. The optional `-rollup_and_alpha` flag will result in the specified sinks receiving, in addition to the rollups, a clone of the enriched `kflow` stream that is received by `ktranslate`.

  ```
  docker run -p 8082:8082 kentik/ktranslate:v2 -format json -sinks stdout -rollups unique,top_src_addr_by_count_dst_addr,dst_addr,src_addr -rollup_top_k 10 -rollup_interval 60 -rollups sum,in_bytes+out_bytes,in_bytes,l4_src_port -rollup_and_alpha -listen=0.0.0.0:8082
  ```

  > **Note:** Rollups are not supported when `-format` is NetFlow.

#### `ktranslate` Output Examples

The following examples illustrate commands that are used for common output destinations (the filter and rollup operations illustrated above may be used with any of the outputs below):

* **Output to Elastic**: Translate `kflow` to JSON and send to a local instance of Elastic via HTTP:

  ```
  docker run -p 8082:8082 kentik/ktranslate:v2 -format json -sinks http -http_url=http://127.0.0.1:9200/indexname/typename/optionalUniqueId -listen=0.0.0.0:8082
  ```
* **Output to file**: Translate `kflow` to InfluxDB line protocol for output to a file:

  ```
  docker run kentik/ktranslate -format influx -sinks file -file_on -file_out destination_directory -listen=0.0.0.0:8082
  ```

  > **Note:** Output to file requires that the `-file_on` flag be present.
* **Output to InfluxDB**: Translate `kflow` to InfluxDB line protocol for output via HTTP to a local instance named "kentik":

  ```
  docker run -p 8082:8082 kentik/ktranslate:v2 -format influx -sinks http -http_url=http://localhost:8086/write?db=kentik -listen=0.0.0.0:8082
  ```
* **Output to Kafka**: Translate `kflow` to Apache Avro and output to a Kafka topic named "kentik\_netflow":

  ```
  docker run -p 8082:8082 kentik/ktranslate:v2 -format avro -sinks kafka --kafka_topic=kentik_netflow -listen=0.0.0.0:8082
  ```
* **Output as NetFlow**: Translate `kflow` to NetFlow and send, with a limit of 10 flows per message, to a local collector running on 127.0.0.1:9913:

  ```
  docker run -p 8082:8082 kentik/ktranslate:v2 -format netflow -sinks net -net_server 127.0.0.1:9913 -max_flows_per_message 10 -listen=0.0.0.0:8082
  ```

  > **Notes:**
  >
  > + NetFlow output defaults to IPFIX (version 10). For NetFlow version 9 add this argument: `-netflow_version netflow9`.
  > + Rollups are not supported when `-format` is NetFlow.
* **Output to New Relic**: Translate `kflow` to JSON and to a New Relic account with the specified ID (requires New Relic API key):

  ```
  docker run -e NEW_RELIC_API_KEY=$API_KEY -p 8082:8082 kentik/ktranslate:v2 -format new_relic -sinks new_relic --nr_account_id=ID -listen=0.0.0.0:8082
  ```
* **Output to Prometheus**: Translate `kflow` to Prometheus and output to a Prometheus server listening on port 8084:

  ```
  docker run -p 8082:8082 kentik/ktranslate:v2 -format prometheus -sinks prometheus -prom_listen=:8084 -listen=0.0.0.0:8082
  ```
* **Output to Splunk**: Send `kflow` to a local instance of Splunk:

  ```
  docker run -e KENTIK_API_TOKEN=kentik_api_token -p 8082:8082 kentik/ktranslate:v2 -format splunk -sinks http -http_url=https://your-collector.splunkcloud.com:8088/services/collector/event -http_header 'Authorization: Splunk placeholder-for-your-splunk-token' -kentik_email=name@company.com -compression gzip -listen=0.0.0.0:8082
  ```

> **Note:** In the Splunk example above, the authorization value that you provide in the `-http_header` argument is a Splunk HTTP Event Collector (HEC) token. See Splunk’s docs on the **HTTP Event Collector**.

### `ktranslate` CLI Reference

The table below lists the arguments of the `ktranslate` CLI version 2.

> **TIP**: Use `-h` to return the current argument list for your `ktranslate` version.

| Argument (type) | Description |
| --- | --- |
| `-asn4` (string) | ASN IPv6 mapping file |
| `-asn6` (string) | ASN IPv6 mapping file |
| `-bootstrap.servers` (string) | List of Kafka broker addresses. |
| `-city` (string) | City mapping file |
| `-compression` (string) | Compression algo to use (none|gzip|snappy|deflate|null) (default "none") |
| `-dns` (string) | Resolve IPs at this ip:port |
| `-file_out` (string) | Write flows seen to log to this directory if set (default "./") |
| `-filters` (value) | Any filters to use. Format: type dimension operator value |
| `-format` (string) | Format to convert kflow to: (json|avro|netflow|influx|prometheus) (default "json") |
| `-geo` (string) | Geo mapping file |
| `-healthcheck` (string) | Bind to this interface to allow healthchecks |
| `-http_header` (value) | Any custom http headers to set on outbound requests |
| `-http_url` (string) | URL to post to (default "http://localhost:8086/write?db=kentik") |
| `-info_collector` | Also send stats about this collector |
| `-interfaces` (string) | Interface mapping file |
| `-kafka.debug` (string) | Debug contexts to enable for kafka |
| `-kafka_topic` (string) | Kafka topic to produce on |
| `-kentik_email` (string) | Email to use for sending flow on to Kentik |
| `-kentik_url` (string) | URL to use for sending flow on to Kentik (default "https://flow.kentik.com/chf") |
| `-listen` (string) | The interface and port on which ktranslate should listen for flow data from Kentik. Default is off (no data received from Kentik). |
| `-log_level` (string) | Logging Level (default "debug") |
| `-mapping` (string) | Mapping file to use for enums (default "config.json") |
| `-max_flows_per_message` (int) | Max number of flows to put in each emitted message (default: 10000) |
| `-max_sql_conns` (int) | Max concurrent SQL connections (default: 16) |
| `-measurement` (string) | Measurement to use for rollups. (default: "kflow") |
| `-metalisten` (string) | HTTP port to bind on (default: "localhost:0") |
| `-metrics` (string) | Metrics Configuration. none|syslog|stderr|graphite:127.0.0.1:2003 (default: "syslog") |
| `-net_protocol` (string) | Use this protocol for writing data (udp|tcp|unix) (default: "udp") |
| `-net_server` (string) | Write flows seen to this address (host and port) |
| `-netflow_version` (string) | Version of netflow to produce: (netflow9|ipfix) (default: "ipfix") |
| `-nr_account_id` (string) | If set, sends flow to New Relic |
| `-nr_url` (string) | URL to use to send into NR (default: "https://insights-collector.newrelic.com/v1/accounts/%s/events") |
| `-olly_dataset` (string) | Olly dataset name |
| `-olly_write_key` (string) | Olly dataset name |
| `-prom_listen` (string) | Bind to listen for prometheus requests on. (default: ":8082") |
| `-region` (string) | Region mapping file |
| `-rollup_and_alpha` | Send both rollups and alpha inputs to sinks. The alpha output is a clone of the enriched kflow stream received by ktranslate. |
| `-rollup_interval` (int) | Export timer for rollups in seconds |
| `-rollup_key_join` (string) | Token to use to join dimension keys together (default: "'") |
| `-rollup_top_k` (int) | Export only these top values (default: 10) |
| `-rollups` (value) | Any rollups to use. Structure: method, name, metric, dimensions (1 through n). Available methods are sum, min, max, mean, median, entropy, percentilerank, percentile, unique  **Note:** For available metrics, see **ktranslate Metrics and Dimensions**. |
| `-s3_bucket` (string) | AWS S3 Bucket to write flows to |
| `-s3_flush_sec` (int) | Create a new output file every this many seconds (default: 60) |
| `-s3_prefix` (string) | AWS S3 Object prefix (default: "/kentik") |
| `-sample_rate` (int) | Sampling rate to use. 1 -> 1:1 sampling, 2 -> 1:2 sampling and so on. (default: 1) |
| `-security.protocol` (string) | Protocol used to communicate with brokers. |
| `-service_name` (string) | Service identifier (default: "ktranslate") |
| `-sinks` (string) | List of sinks to send data to. Options: (kafka|stdout|new\_relic|kentik|net|http|splunk|prometheus|file) (default: "stdout") |
| `-ssl.ca.location` (string) | Location of the SSL CA certificate file. |
| `-stdout` | Log to stdout (default: true) |
| `-tag_lookup` (string) | Tag service port to run lookups on |
| `-threads` (int) | Number of threads to run for processing |
| `-v` | Show version and build information |

### `ktranslate` Metrics and Dimensions

The table below lists the allowed values for metrics and dimensions in a `ktranslate` rollup or filter. Usage depends on the type of the value:

* **String**:

  + In a rollup, if the method is `unique`, a value of type `string` can be used as a metric.
  + In all other cases, `string` can only be used as a rollup dimension or filter dimension.
  + In an `addr` filter, the dimension must be one of the below string values ending in "\_addr" (e.g., `dst_addr`, `src_addr`).
* **Int**: Any `int` value can be used as a metric or a dimension.

| Value | Type |
| --- | --- |
| `company_id` | int64 |
| `custom_bigint` | string |
| `custom_int` | string |
| `custom_str` | string |
| `device_id` | int64 |
| `device_name` | string |
| `dst_addr` | string |
| `dst_as` | int64 |
| `dst_bgp_as_path` | string |
| `dst_bgp_comm` | string |
| `dst_eth_mac` | string |
| `dst_flow_tags` | string |
| `dst_geo` | string |
| `dst_geo_city` | string |
| `dst_geo_region` | string |
| `dst_nexthop` | string |
| `dst_nexthop_as` | int64 |
| `dst_route_prefix` | int64 |
| `dst_second_asn` | int64 |
| `dst_third_asn` | int64 |
| `header_len` | int64 |
| `in_bytes` | int64 |
| `in_pkts` | int64 |
| `input_int_alias` | string |
| `input_int_desc` | string |
| `input_port` | int64 |
| `l4_dst_port` | int64 |
| `l4_src_port` | int64 |
| `out_bytes` | int64 |
| `out_pkts` | int64 |
| `output_int_alias` | string |
| `output_int_desc` | string |
| `output_port` | int64 |
| `protocol` | int64 |
| `sample_rate` | int64 |
| `sampled_packet_size` | int64 |
| `src_addr` | string |
| `src_as` | int64 |
| `src_bgp_comm` | string |
| `src_bgp_as_path` | string |
| `src_eth_mac` | string |
| `src_flow_tags` | string |
| `src_geo` | string |
| `src_geo_city` | string |
| `src_geo_region` | string |
| `src_nexthop` | string |
| `src_nexthop_as` | int64 |
| `src_route_prefix` | int64 |
| `src_second_asn` | int64 |
| `src_third_asn` | int64 |
| `tcp_flags` | int64 |
| `tcp_rx` | int64 |
| `timestamp` | int64 |
| `tos` | int64 |
| `vlan_in` | int64 |
| `vlan_out` | int64 |

> **Notes:**
>
> * The `custom_bigint`, `custom_int`, and `custom_string` values are used for dynamic fields such as **Universal Data Records** (UDR) and **Custom Dimensions**.
> * For more information about above table values, see **Metrics and Dimensions**.

## Firehose Data Sources

Once a `ktranslate` instance is deployed in your datacenter or cloud resource, you can send `kflow` to it from KDE or other sources (see **ktranslate Scenarios**). Since multiple instances may exist, indicate the instance to Kentik via the portal for each data source.

![Kentik Firehose URL Endpoint configuration for data integration and analytics systems.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/UK-Firehose.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A10Z&se=2026-05-12T09%3A41%3A10Z&sr=c&sp=r&sig=qMQIZt9SfBPaHX3YeKQjgnmWL%2FQEws5UMtyT5Oxp0r0%3D)

To send `kflow` for a data source from KDE to `ktranslate`:

1. Determine the public IP address and port number of the `ktranslate` instance to which you want to send `kflow` for a device.
2. From the Kentik portal’s main menu, click **Settings** in the column at right to go to the Settings page.
3. In the card at upper left, click the **Network Devices** link to go to the Devices page.
4. Find the device (router or host) for which you want to send enriched flow data to `ktranslate`. Click the **Edit** icon at the right of the device's row in the list to open the Devices dialog.
5. In the **Firehose Endpoint** field in the dialog's **Integrations** tab, enter the public IP address and port number of the `ktranslate` instance, then click **Save** to save changes and exit.

> **Note:** To send `kflow` for this device to multiple `ktranslate` instances, enter a comma-separated list of the corresponding URLs.
