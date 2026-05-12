# BGP Monitoring APIs

Source: https://kb.kentik.com/docs/bgp-monitoring-apis

---

This article covers how to get started with the BGP Monitoring APIs.

> **Notes:**
>
> * The **API Tester** in the v4 Portal enables you to securely call the methods of these APIs using your organization's Kentik data.
> * BGP monitoring is supported in the Kentik portal via BGP Monitor tests, which report on BGP reachability, path changes, and events. For more information, see **BGP Monitor Details Page****.**
> * Protobuf and OpenAPI specifications for Kentik’s v6 APIs are available in our **api-schema-public** repository.

## BGP Monitoring Usage

The topics below provide important background information for the use of these APIs.

### Overview

The BGP Monitoring API provides programmatic access to Kentik's BGP monitoring system. The APIs handle two distinct functions:

* **Admin service**: manages monitoring of advertised BGP prefixes as collected from global vantage points.
* **Data service**: handles retrieval of two classes of data for any IP prefix:

  + metrics for reachability and path changes, which provide a time series for a specified interval;
  + routes, which provide a snapshot of routing information for particular timestamp.

The services described above use the following endpoints:

| Endpoint | Purpose |
| --- | --- |
| BgpMonitoringAdminService | CRUD operations for BGP monitors |
| BgpMonitoringDataService | Retrieval of BGP monitoring data |

Both REST endpoints and gRPC RPCs are provided.

#### Known Limitations

The API currently does not support retrieval of BGP event history.

### BGP Monitoring Data

Kentik's BGP monitoring API provides the following information about network prefixes:

* **Reachability**: expressed as the percentage of vantage points reporting the prefix
* **Path changes**: the number of changes in a given time range
* **Routes**: AS Path, next hop

#### Data Retrieval Limitations

The following considerations apply to data retrieval using these APIs:

* BGP metrics and routes can be retrieved for any prefix with prefix length greater than 14 and any time interval. Prefixes with smaller length must be split into subnets (sub-prefixes) and requested individually.
* Requests for BGP metrics can can specify one or more metric types to be retrieved for a single prefix. Observed data are returned as a sequence of messages, one for each metric type and observation time.
* Route requests retrieve snapshot of information for specified prefix and time. Responses contain also mapping of ASNs to AS names for convenience.

### Working with BGP Monitoring

The topics below cover important information related to using Kentik's BGP monitoring APIs.

#### BGP Prefixes

These APIs represent BGP prefixes using NLRI (Network Layer Reachability Information) objects. While NLRI objects allow representation of any type of network address, the system currently supports only IPv4 and IPv6 unicast in standard textual representation:

* **IPv4**: dotted-decimal notation
* **IPv6**: RFC 5952 syntax

#### BGP Monitor Tests

BGP Monitors enable the generation of alerts for group network prefixes (IPv4 or IPv6) based on the following factors:

* Visibility/reachability of prefixes
* Mismatch of originating ASNs with a specified list of valid originators
* Validity of advertisements with respect to RPKI (Resource Public Key Infrastructure)

#### Configuration Limitations

The following considerations apply when configuring a BGP monitor test:

* Each BGP monitor instance can monitor up to 10 prefixes.
* The length of each monitored (IPv4 or IPv6) prefix must be greater than 14. Prefixes with smaller length must be split into subnets (sub-prefixes).

## BGP Monitoring RPCs

The tables below provide a quick reference to key information about each Remote Procedure Call in these APIs. Use the links in the tables to see schemas for the request and/or response body (if any).

> **Note:** To test methods using your own Kentik data, use the portal's **API Tester**.

### ListMonitors

**API: BgpMonitoringAdminService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| GET | /bgp\_monitoring/v202210 /monitors | Returns list of all BGP monitors present in the account. |
| |  |  | | --- | --- | | **Request body:** None  **Parameters:** None | **Response body:**  v202210ListMonitorsResponse | | | |

### CreateMonitor

**API: BgpMonitoringAdminService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| POST | /bgp\_monitoring/v202210 /monitors | Creates new BGP Monitor and if successful returns its configuration. |
| |  |  | | --- | --- | | **Request body:** v202210CreateMonitorRequest  **Parameters:**  None | **Response body:**  v202210CreateMonitorResponse | | | |

### GetMonitor

**API: BgpMonitoringAdminService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| GET | /bgp\_monitoring/v202210 /monitors/{id} | Returns configuration of existing BGP monitor with specific ID. |
| |  |  | | --- | --- | | **Request body:** None | **Response body:**  v202210GetMonitorResponse |  **Parameters:**  | Name | Description | Required | Type | | --- | --- | --- | --- | | id | ID of the BGP monitor to be retrieved | true | string | | | |

### DeleteMonitor

**API: BgpMonitoringAdminService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| DELETE | /bgp\_monitoring/v202210 /monitors/{id} | Delete BGP monitor with with specific ID. |
| |  |  | | --- | --- | | **Request body:** None | **Response body:**  v202210DeleteMonitorResponse |  **Parameters:**  | Name | Description | Required | Type | | --- | --- | --- | --- | | id | ID of the BGP monitor to be deleted | true | string | | | |

### UpdateMonitor

**API: BgpMonitoringAdminService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| PUT | /bgp\_monitoring/v202210 /monitors/{id} | Updates configuration of BGP monitor with specific ID and returns updated configuration. |
| |  |  | | --- | --- | | **Request body**: v202210UpdateMonitorRequest | **Response body:** v202210UpdateMonitorResponse |  **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | id | System generated unique identifier | true | string | | | |

### SetMonitorStatus

**API: BgpMonitoringAdminService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| PUT | /bgp\_monitoring/v202210 /monitors/{id}/status | Sets administrative status of BGP monitor with specific ID. |
| |  |  | | --- | --- | | **Request body:** v202210SetMonitorStatusRequest | **Response body**: v202210SetMonitorStatusResponse |  **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | id | ID of the BGP monitor whose status is to be modified | true | string | | | |

### GetMetricsForTarget

**API: BgpMonitoringDataService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| POST | /bgp\_monitoring/v202210 /metrics | Retrieve metric data for single BGP prefix and time interval. |
| |  |  | | --- | --- | | **Request body**: v202210GetMetricsForTargetRequest  **Parameters**: None | **Response body**:  v202210GetMetricsForTargetResponse | | | |

### GetRoutesForTarget

**API: BgpMonitoringDataService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| POST | /bgp\_monitoring/v202210 /routes | Retrieve snapshot of route information for single BGP prefix at specific time. |
| |  |  | | --- | --- | | **Request body**: v202210GetRoutesForTargetRequest  **Parameters**: None | **Response body**:  v202210GetRoutesForTargetResponse | | | |

## BGP Monitoring Schemas

This API uses the following schemas.

#### protobufAny

|  |  |
| --- | --- |
| **Schema**: protobufAny | **Type**: object |
| **Properties**:  | Name | Value | | --- | --- | | typeUrl | type: string | | value | type: string  format: byte | | |

#### rpcStatus

|  |  |
| --- | --- |
| **Schema:** rpcStatus | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | code | type: integer  format: int32 | | message | type: string | | details | type: array  items: $ref: protobufAny | | |

#### Afi

|  |  |
| --- | --- |
| **Schema:** v202104Afi | **Type:** string |
| **Attributes**:  | Key | Value | | --- | --- | | enum | AFI\_UNSPECIFIED, AFI\_IP4, AFI\_IP6 | | default | AFI\_UNSPECIFIED | | |

#### RpkiStatus

|  |  |
| --- | --- |
| **Schema:** v202104RpkiStatus | **Type:** string |
| **Attributes**:  | Key | Value | | --- | --- | | enum | RPKI\_STATUS\_UNSPECIFIED, RPKI\_STATUS\_NOT\_FOUND, RPKI\_STATUS\_EXPLICIT\_INVALID, RPKI\_STATUS\_INVALID\_PREFIX, RPKI\_STATUS\_INVALID, RPKI\_STATUS\_VALID, RPKI\_STATUS\_ERROR | | default | RPKI\_STATUS\_UNSPECIFIED | | description | * RPKI\_STATUS\_UNSPECIFIED: Invalid value. * RPKI\_STATUS\_NOT\_FOUND: No matching ROAs found for the prefix. * RPKI\_STATUS\_EXPLICIT\_INVALID: Explicitly invalid prefix matching ROA * RPKI\_STATUS\_INVALID\_PREFIX: Prefix format not suitable for validation * RPKI\_STATUS\_INVALID: Prefix origin does not match any matching ROA * RPKI\_STATUS\_VALID: Prefix origin matches an ROA * RPKI\_STATUS\_ERROR: Error during validation | | |

#### Safi

|  |  |
| --- | --- |
| **Schema:** v202104Safi | **Type:** string |
| | Key | Value | | --- | --- | | enum | SAFI\_UNSPECIFIED, SAFI\_UNICAST, SAFI\_MPLS, SAFI\_L3VPN | | default | SAFI\_UNSPECIFIED | | |

#### VantagePoint

|  |  |
| --- | --- |
| **Schema:** v202104VantagePoint | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | dataset | type: string  description: Name of the collector data set  readOnly: true | | collector | type: string  description: Name of the collector  readOnly: true | | peerAsn | type: integer  format: int64  description: ASN of the peer from which the collector receives BGP data  readOnly: true | | peerIp | type: string  description: IP address of the peer from which the collector receives BGP data  readOnly: true | | |

#### UserInfo

|  |  |
| --- | --- |
| **Schema:** v202202UserInfo | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | id | type: string  description: Unique ID of the user  readOnly: true | | email | type: string  description: E-mail address of the user  readOnly: true | | fullName | type: string  description: Full name of the user  readOnly: true | | |

#### BgpHealthSettings

|  |  |
| --- | --- |
| **Schema:** v202210BgpHealthSettings | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | reachabilityWarning | type: number  format: float  description: Threshold (in percents) for triggering warning level alert | | reachabilityCritical | type: number  format: float  description: Threshold (in percents) for triggering critical level alert | | |

#### BgpMetric

|  |  |
| --- | --- |
| **Schema:** v202210BgpMetric | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | timestamp | type: string  format: date-time  description: UTC time of the observation  readOnly: true | | nlri | $ref: v202210Nlri | | reachability | type: number  format: float  description: Percentage of vantage points reporting the prefix reachable  readOnly: true | | pathChanges | type: integer  format: int64  description: Number of AS path changes for the prefix across all vantage points in the queried time interval  readOnly: true | | |

#### BgpMetricType

|  |  |
| --- | --- |
| **Schema:** v202210BgpMetricType | **Type:** string |
| **Attributes**:  | Key | Value | | --- | --- | | enum | BGP\_METRIC\_TYPE\_UNSPECIFIED, BGP\_METRIC\_TYPE\_REACHABILITY, BGP\_METRIC\_TYPE\_PATH\_CHANGES | | default | BGP\_METRIC\_TYPE\_UNSPECIFIED | | description | * BGP\_METRIC\_TYPE\_UNSPECIFIED: Invalid value. * BGP\_METRIC\_TYPE\_REACHABILITY: Reachability metric in percents * BGP\_METRIC\_TYPE\_PATH\_CHANGES: Number of path changes over time interval | | |

#### BgpMonitor

|  |  |
| --- | --- |
| **Schema:** v202210BgpMonitor | **Type:** object |
| **Properties (\* = required)**  | Name | Value | | --- | --- | | id | type: string  description: System generated unique identifier  readOnly: true | | name \* | type: string  description: User selected name of the monitor | | status | $ref: v202210BgpMonitorStatus | | settings | $ref: v202210BgpMonitorSettings | | cdate | type: string  format: date-time  description: Creation timestamp (UTC)  readOnly: true | | edate | type: string  format: date-time  description: Last modification timestamp (UTC)  readOnly: true | | createdBy | $ref: v202202UserInfo | | lastUpdatedBy | $ref: v202202UserInfo | | labels | type: array  items: type: string | | |

#### BgpMonitorSettings

|  |  |
| --- | --- |
| **Schema:** v202210BgpMonitorSettings | **Type:** object |
| **Properties (\* = required)**  | Name | Value | | --- | --- | | allowedAsns | type: array  items: type: integer  format: int64  description: List of ASNs that are valid originators of monitored prefixes | | targets \* | type: array  items: $ref: v202210Nlri  description: List of prefixes to monitor | | checkRpki | type: boolean  description: Enable verification of validity of advertisements of monitored prefixes with respect to RPKI | | includeCoveredPrefixes | type: boolean  description: Include advertised subnets (sub-prefixes) of monitored prefixes | | healthSettings | $ref: v202210BgpHealthSettings | | notificationChannels | type: array  items: type: string  description: List of IDs of notification channels for delivery of alerts | | notes | type: string  description: Free form notes documenting the monitor | | allowedUpstreams | type: array  items: type: integer  format: int64  description: List of ASNs that are expected to propagate monitored prefixes | | |

#### BgpMonitorStatus

|  |  |
| --- | --- |
| **Schema:** v202210BgpMonitorStatus | **Type:** string |
| **Attributes**:  | Key | Value | | --- | --- | | enum | BGP\_MONITOR\_STATUS\_UNSPECIFIED, BGP\_MONITOR\_STATUS\_ACTIVE, BGP\_MONITOR\_STATUS\_PAUSED, BGP\_MONITOR\_STATUS\_DELETED | | default | BGP\_MONITOR\_STATUS\_UNSPECIFIED | | description | • BGP\_MONITOR\_STATUS\_UNSPECIFIED: Invalid value.  • BGP\_MONITOR\_STATUS\_ACTIVE: Monitor is active.  • BGP\_MONITOR\_STATUS\_PAUSED: Monitor is paused.  • BGP\_MONITOR\_STATUS\_DELETED: Monitor is deleted. Not user settable | | |

#### CreateMonitorRequest

|  |  |
| --- | --- |
| **Schema:** v202210CreateMonitorRequest | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | monitor | $ref: v202210BgpMonitor | | |

#### CreateMonitorResponse

|  |  |
| --- | --- |
| **Schema:** v202210CreateMonitorResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | monitor | $ref: v202210BgpMonitor | | |

#### DeleteMonitorResponse

|  |  |
| --- | --- |
| **Schema:** v202210DeleteMonitorResponse | **Type:** object |
| **Properties**: None. | |

#### GetMetricsForTargetRequest

|  |  |
| --- | --- |
| **Schema:** v202210GetMetricsForTargetRequest | **Type:** object |
| **Properties (\* = required)**  | Name | Value | | --- | --- | | startTime \* | type: string  format: date-time  description: UTC timestamp of the beginning of queried interval | | endTime \* | type: string  format: date-time  description: UTC timestamp of the end of queried interval | | target | $ref: v202210Nlri | | includeCovered | type: boolean  description: Return metrics for subnets (sub-prefixes) of the target prefix (default: false) | | metrics \* | type: array  items: $ref: v202210BgpMetricType  description: List of one or more BGP metric types to return | | |

#### GetMetricsForTargetResponse

|  |  |
| --- | --- |
| **Schema:** v202210GetMetricsForTargetResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | metrics | type: array  items: $ref: v202210BgpMetric  description: List of observations (metric values)  readOnly: true | | |

#### GetMonitorResponse

|  |  |
| --- | --- |
| **Schema:** v202210GetMonitorResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | monitor | $ref: v202210BgpMonitor | | |

#### GetRoutesForTargetRequest

|  |  |
| --- | --- |
| **Schema:** v202210GetRoutesForTargetRequest | **Type:** object |
| **Properties** **(\* = required)**  | Name | Value | | --- | --- | | time \* | type: string  format: date-time  description: Reference UTC time for the route information snapshot | | target | $ref: v202210Nlri | | includeCovered | type: boolean  description: Return routes for subnets (sub-prefixes) of the target prefix (default: false) | | checkRpki | type: boolean  description: Return information about validity of prefix advertisements with respect to RPKI (default: false) | | |

#### GetRoutesForTargetResponse

|  |  |
| --- | --- |
| **Schema:** v202210GetRoutesForTargetResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | routes | type: array  items: $ref: v202210RouteInfo  description: List of routes  readOnly: true | | asNames | type: object  additionalProperties: type: string  description: look-aside map AS names keyed by ASNs  readOnly: true | | |

#### ListMonitorsResponse

|  |  |
| --- | --- |
| **Schema:** v202210ListMonitorsResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | monitors | type: array  items: $ref: v202210BgpMonitor  description: List of BGP monitors configured in the account  readOnly: true | | invalidCount | type: integer  format: int64  description: Number of invalid objects (should be always zero)  readOnly: true | | |

#### Nlri

|  |  |
| --- | --- |
| **Schema:** v202210Nlri | **Type:** object |
| **Properties (\* = required)**  | Name | Value | | --- | --- | | afi | $ref: v202104Afi | | safi | $ref: v202104Safi | | prefix \* | type: string  description: Actual prefix data (format depends on AFI) | | |

#### RouteInfo

|  |  |
| --- | --- |
| **Schema:** v202210RouteInfo | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | nlri | $ref: v202210Nlri | | originAsn | type: integer  format: int64  description: The autonomous system number originating the NLRI  readOnly: true | | asPath | type: array  items: type: string  description: AS path observed at the vantage point for the NLRI | | vantagePoint | $ref: v202104VantagePoint | | rpkiStatus | $ref: v202104RpkiStatus | | nexthop | type: string  description: IP address of the first route hop from the vantage point toward the target  readOnly: true | | |

#### SetMonitorStatusRequest

|  |  |
| --- | --- |
| **Schema:** v202210SetMonitorStatusRequest | **Type:** object |
| **Properties (\* = required)**  | Name | Value | | --- | --- | | id \* | type: string  description: ID of the BGP monitor whose status is to be modified | | status | $ref: v202210BgpMonitorStatus | | |

#### SetMonitorStatusResponse

|  |  |
| --- | --- |
| **Schema:** v202210SetMonitorStatusResponse | **Type:** object |
| **Properties**: None. | |

#### UpdateMonitorRequest

|  |  |
| --- | --- |
| **Schema:** v202210UpdateMonitorRequest | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | monitor | $ref: v202210BgpMonitor | | |

#### UpdateMonitorResponse

|  |  |
| --- | --- |
| **Schema:** v202210UpdateMonitorResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | monitor | $ref: v202210BgpMonitor | | |
