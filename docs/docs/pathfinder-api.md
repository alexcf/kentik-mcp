# Pathfinder API

Source: https://kb.kentik.com/docs/pathfinder-api

---

This article covers how to get started with the Pathfinder API:

> **Notes**:
>
> * The **API Tester** in the v4 Portal enables you to securely call the methods of these APIs using your organization's Kentik data.
> * For information about Cloud Pathfinder, start with **Cloud Pathfinder**.
> * Protobuf and OpenAPI specifications for Kentik’s v6 APIs are available in our **api-schema-public** repository.

## Pathfinder Usage

The Pathfinder API provides programmatic access to Kentik's Cloud Pathfinder, enabling you to check and monitor forward and return network paths between two points in your cloud environment. It supports entities like instances, subnets, and network interfaces over a specified time range.

### Supported Cloud Providers

The API's functionality and configuration parameter vary by cloud provider, as follows:

* Amazon Web Services (AWS), see **Kentik for AWS**
* Microsoft Azure, see **Kentik for Azure**

### Supported Entity Types

For AWS, the API supports checking connectivity between:

* Subnets
* Instances
* Network Interfaces

For Azure, the API supports checking connectivity between:

* Subnets

### Requirements

To successfully query connections between entity resources with Kentik's Pathfinder API, you must first configure a Kentik cloud export (see Cloud Exports and Devices) for a supported cloud provider (AWS or Azure). Once configured, Kentik will regularly pull metadata from the specified cloud account. Pathfinder then uses this metadata to:

* Visualize the path for connections between any two points in your AWS network.
* Check routing and NACL/SG configurations to quickly highlight blocked connections.
* Provide a direct link to misconfigured gateways or security policies in the cloud dashboard for quick issue resolution.

### Frequency of Available Data

Since Pathfinder relies on discovered metadata, query results are limited to the timeframe of the latest metadata retrieval fetch, currently 15-minutes intervals.

## Pathfinder RPCs

The tables below provide a quick reference to key information about each Remote Procedure Call in these APIs. Use the links in the tables to see schemas for the request and/or response body (if any).

> **Note:** To test methods using your own Kentik data, use the portal's API Tester.

### CreatePathfinderReport

The endpoints are as follows:

| REST Method | REST Endpoint | Description | Parameters |
| --- | --- | --- | --- |
| **POST** | /pathfinder/v202505beta1/create | Create a pathfinder report based on configuration provided in the request. | None |

* Request Body Schema: **CreatePathfinderReportRequest**
* Response Body Schema: **CreatePathfinderReportResponse**

> **Note**: To test methods using your own Kentik data, use the portal's **API Tester**.

## Pathfinder Schemas

This API uses the following schemas.

### CreatePathfinderReportRequest

The schema is as follows:

| Field Name | Type | Description | Required |
| --- | --- | --- | --- |
| cloudProvider | Object | Reference to v202505beta1CloudProvider schema. | Yes |
| src | String | Source | Yes |
| dst | String | Destination | Yes |
| dstPort | String | Destination Port | Yes |
| protocol | String | Protocol | Yes |
| srcType | Object | Reference to v202505beta1EntityType schema. | No |
| dstType | Object | Reference to v202505beta1EntityType schema. | No |
| startTime | String | Start Timestamp (UTC) which defines the time range for the report. | Yes |
| endTime | String | End timestamp (UTC) which defines the time range for the report. | Yes |

### CreatePathfinderReportResponse

The schema is as follows:

| Field Name | Type | Description | Required |
| --- | --- | --- | --- |
| reachable | String | Attribute controlling whether the instance is active | Yes |
| returnReachable | String | Attribute controlling whether the instance is active | Yes |
| queryStatus | String | Status of the query. | Yes |
| reportUrl | String | URL to view the pathfinder report in the Kentik platform. | Yes |
| paths | Array | The connection paths in your network. (Items are $ref to v202505beta1PathElement schema) | Yes |
| returnPaths | Array | The return connection paths in your network. (Items are $ref to v202505beta1PathElement schema) | Yes |
| lastMetadataFetch | String | The timestamp indicating when the metadata used to build the pathfinder report was last fetched. | Yes |
| summary | String | The pathfinder report summary | Yes |
