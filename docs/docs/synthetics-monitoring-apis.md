# Synthetics Monitoring APIs

Source: https://kb.kentik.com/docs/synthetics-monitoring-apis

---

This article covers how to get started with the Synthetics Monitoring APIs.

> **Notes:**
>
> * The **API Tester** in the v4 Portal enables you to securely call the methods of these APIs using your organization's Kentik data.
> * For information about Kentik's synthetic monitoring capabilities, start with **Synthetics**.
> * Protobuf and OpenAPI specifications for Kentik’s v6 APIs are available in our **api-schema-public** repository.

## Synthetics Usage

The topics below provide important background information for the use of these APIs.

### Overview

The Synthetics Monitoring API provides programmatic access to Kentik's **synthetic monitoring system**. The API consists of two endpoints:

| Endpoint | Purpose |
| --- | --- |
| SyntheticsAdminService | CRUD operations for synthetic tests, agents, and offline agent alerts |
| SyntheticsDataService | Retrieval of synthetic test results and network traces |

Both REST endpoint and gRPC RPCs are provided. Note: API version v202309 is the same as v202202 except that the timestamps returned for synthetic test results are closer to when the test was actually run.

#### Known Limitations

The API currently does not support the following **Synthetic Test Types**:

* BGP Monitor tests, which are supported in a **separate API**
* Transaction tests.

#### Additional Public Resources

Kentik community **Python**and **Go**SDKs provide language-specific support for using this and other Kentik APIs. These SDKs can be also used as example code for development. A **Terraform provider** is available for configuring tests and agents for Kentik synthetic monitoring.

### Anatomy of a Synthetic Test

Each `Test` consists of one or more tasks. Tasks are executed by monitoring Agents that send synthetic traffic (probes) over the network. The API currently supports following tasks:

| Task name | Purpose |
| --- | --- |
| ping | Test basic address, and optionally TCP port reachability |
| traceroute (a.k.a. trace) | Discover unidirectional network path |
| http | Perform a simple HTTP/HTTPS request |
| page-load | Use headless Chromium to execute an HTTP/HTTPS request |
| dns | Execute a DNS query |

The set of tasks executed on behalf of a given test depends on the `type` of that test. The following test types are currently supported by the API:

| API type | Portal (UI) equivalent | Tasks |
| --- | --- | --- |
| ip | IP Address | ping, traceroute |
| hostname | Hostname | ping, traceroute |
| network\_grid | Network Grid | ping, traceroute |
| agent | Agent-to-Agent | ping, traceroute |
| network\_mesh | Network Mesh | ping, traceroute |
| flow | Autonomous Tests (5 variants) | ping, traceroute |
| url | HTTP(S) or API | http, ping (optional), traceroute (optional) |
| page\_load | Page Load | page-load, ping (optional), traceroute (optional) |
| dns | DNS Server Monitor | dns |
| dns\_grid | DNS Server Grid | dns |

> **Note:** `ping` and `traceroute` tasks are always run together (never one without the other).

### Test Attributes and Settings

The attributes of the test object enable configuration of test settings, access to test metadata, and access to runtime state information.

#### State and Metadata Attributes

The following table lists the metadata and state attributes:

| Attribute | Access | Purpose |
| --- | --- | --- |
| id | RO | System-generated unique identifier of the test |
| name | RW | User specified name for the test (need not be unique) |
| type | RO (after creation) | Type of the test (set on creation; read-only thereafter) |
| status | RW | Life-cycle status of the test |
| cdate | RO | Creation timestamp |
| edate | RO | Last-modification timestamp |
| created\_by | RO | Identity of the user that created the test |
| last\_updated\_by | RO | Identity of the latest user to modify the test |
| labels | RW | List of names of labels applied to the test |

Test configuration is performed via the test's `settings` attribute. Some settings are common to all tests while others are specific to tests of a given type.

#### Common Test Settings

The following settings are used for tests of all types:

| Attribute | Purpose | Required |
| --- | --- | --- |
| agentIds | IDs of agents to execute tasks for the test | YES |
| period | Test execution interval in seconds | NO (default 60s) |
| family | IP address family. Used only for tests whose type is url or dns. Selects which type of DNS resource is queried for resolving hostname to target address | NO (default IP\_FAMILY\_DUAL) |
| notificationChannels | List of notification channels for the test | NO (default empty list) |
| healthSettings | A HealthSettings object that configures health settings for this test, which includes metric thresholds that define health status (warning and critical) and trigger associated alarms. | YES |
| ping | A TestPingSettings object that configures the ping task of the test | NO (default depends on test type) |
| trace | A TestTraceSettings object that configures the trace task of the test | NO (default depends on test type) |
| tasks | List of names of the tasks that will be executed for this test | YES |

#### Type-Specific Settings

Each test type has its own configuration object that represents the settings for that type. These type-specific objects are referenced by the attributes in `Test.settings`:

| Test type | Settings attribute | Configuration object |
| --- | --- | --- |
| ip | ip | IpTest |
| hostname | hostname | HostnameTest |
| network\_grid | networkGrid | IpTest |
| agent | agent | AgentTest |
| network\_mesh | networkMesh | NetworkMeshTest |
| flow | flow | FlowTest |
| url | url | UrlTest |
| page\_load | pageLoad | PageLoadTest |
| dns | dns | DnsTest |
| dns\_grid | dnsGrid | DnsTest |

### Test Results

Results of synthetic tests are returned as a sequence of `TestResults`  objects. Each such object represents measurements and health evaluation for a single test at specific point in time. Measurements and health evaluation are grouped by agent and by task. Granularity of timestamps in test results depends on the frequency (period) of the test and on the requested time range. The minimum granularity is 1 minute (even when period < 1 minute). The longer the time range, the lower the granularity.

### Network Traces

Synthetic tests that include the `traceroute` task collect the unidirectional network path from the agent to the target for each agent/target pair. The trace data are returned in the `GetTraceForTestResponse` object. The `paths` attribute of this object contains the collected network path for each agent/target pair and the round-trip time (RTT) to each hop. Hops in actual network traces are identified by a `nodeId`. The mapping of node IDs to address, name, location, and other attributes of the hop is provided in a map that is stored in the `nodes` attribute of the `GetTraceForTestResponse`  object.

### Agents

The Kentik synthetic monitoring system recognizes 2 types of agents:

* **Global** (public): Managed by Kentik and available to every Kentik user. All information about global agents in this API is read-only.
* **Private**: Deployed by each customer and available only to that customer. To be visible in this API, a private agent must first associate itself with a customer account by contacting the Kentik system (via private API). Once the agent is associated it can be authorized via the API by changing its `status` to `AGENT_STATUS_OK`. For more information about private agent deployment, see **Synthetic Agent Deployments**.

## Synthetics RPCs

The tables below provide a quick reference to key information about each Remote Procedure Call in these APIs. Use the links in the tables to see schemas for the request and/or response body (if any).

> **Note:** To test methods using your own Kentik data, use the portal's **API Tester**.

### ListAgentAlerts

**API: SyntheticsAdminService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| GET | /synthetics/v202309 /agentAlerts | Lists all agent alert configurations, optionally filtered by a list of agent ids. |
| |  |  | | --- | --- | | **Request body**:  None | **Response body:**  v202309ListAgentAlertsResponse |  **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | agentIds | Undefined. | false | array | | | |

### CreateAgentAlert

**API: SyntheticsAdminService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| POST | /synthetics/v202309 /agentAlerts | Creates a new agent alert configuration. |
| |  |  | | --- | --- | | **Request body:**  v202309CreateAgentAlertRequest | **Response body:**  v202309CreateAgentAlertResponse | | **Parameters**: None | | | | |

### GetAgentAlert

**API: SyntheticsAdminService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| GET | /synthetics/v202309 /agentAlerts/{id} | Retrieves an existing agent alert configuration. |
| |  |  | | --- | --- | | **Request body:**  None | **Response body:**  v202309GetAgentAlertResponse |  **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | id | Undefined. | true | string | | | |

### DeleteAgentAlert

**API: SyntheticsAdminService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| DELETE | /synthetics/v202309 /agentAlerts/{id} | Deletes an existing agent alert configuration. |
| |  |  | | --- | --- | | **Request body:**  None | **Response body:**  v202309DeleteAgentAlertResponse |  **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | id | Undefined. | true | string | | | |

### UpdateAgentAlert

**API: SyntheticsAdminService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| PUT | /synthetics/v202309 /agentAlerts/{id} | Updates an existing agent alert configuration with the time threshold and notification channels provided. |
| |  |  | | --- | --- | | **Request body:**  v202309UpdateAgentAlertRequest | **Response body:**  v202309UpdateAgentAlertResponse |  **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | id | Undefined. | true | string | | | |

### ListAgents

**API: SyntheticsAdminService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| GET | /synthetics/v202309 /agents | Returns list of all synthetic agents available in the account. |
| |  |  | | --- | --- | | **Request body:**  None | **Response body:**  v202309ListAgentsResponse | | **Parameters:**  None |  | | | |

### UpdateAgent

**API: SyntheticsAdminService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| PUT | /synthetics/v202309 /agents/{agent.id} | Update configuration of a synthetic agent. |
| |  |  | | --- | --- | | **Request body:**  v202309UpdateAgentRequest | **Response body:**  v202309UpdateAgentResponse |  **Parameters:**  | Name | Description | Required | Type | | --- | --- | --- | --- | | agent.id | Unique identifier of the agent | true | string | | | |

### GetAgent

**API: SyntheticsAdminService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| GET | /synthetics/v202309 /agents/{agent.id} | Returns information about the requested synthetic agent. |
| |  |  | | --- | --- | | **Request body:**  None | **Response body:**  v202309GetAgentResponse |  **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | agent.id | ID of the requested agent | true | string | | | |

### DeleteAgent

**API: SyntheticsAdminService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| DELETE | /synthetics/v202309 /agents/{agent.id} | Deletes the requested agent. The deleted agent is removed from configuration of all tests. |
| |  |  | | --- | --- | | **Request body:**  None | **Response body**:  v202309DeleteAgentResponse |  **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | agent.id | ID of the agent to be deleted | true | string | | | |

### ListTests

**API: SyntheticsAdminService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| GET | /synthetics/v202309 /tests | Returns a list of all configured active and paused synthetic tests. |
| |  |  | | --- | --- | | **Request body:**  None | **Response body:**  v202309ListTestsResponse | | **Parameters**: None |  | | | |

### CreateTest

**API: SyntheticsAdminService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| POST | /synthetics/v202309 /tests | Create synthetic test based on configuration provided in the request. |
| |  |  | | --- | --- | | **Request body:**  v202309CreateTestRequest | **Response body:**  v202309CreateTestResponse | | **Parameters**:  None |  | | | |

### GetTest

**API: SyntheticsAdminService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| GET | /synthetics/v202309 /tests/{id} | Returns configuration and status for the requested synthetic test. |
| |  |  | | --- | --- | | **Request body:**  None | **Response body**:  v202309GetTestResponse |  **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | id | ID of requested test | true | string | | | |

### DeleteTest

**API: SyntheticsAdminService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| DELETE | /synthetics/v202309 /tests/{id} | Deletes the synthetics test. All accumulated results for the test cease to be accessible. |
| |  |  | | --- | --- | | **Request body:**  None | **Response body:**  v202309DeleteTestResponse |  **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | id | ID of the test to be deleted | true | string | | | |

### UpdateTest

**API: SyntheticsAdminService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| PUT | /synthetics/v202309 /tests/{id} | Updates configuration of a synthetic test. |
| |  |  | | --- | --- | | **Request body:**  v202309UpdateTestRequest | **Response body:**  v202309UpdateTestResponse |  **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | id | Unique ID of the test | true | string | | | |

### SetTestStatus

**API: SyntheticsAdminService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| PUT | /synthetics/v202309 /tests/{id}/status | Update status of a synthetic test |
| |  |  | | --- | --- | | **Request body:**  v202309SetTestStatusRequest | **Response body:**  v202309SetTestStatusResponse |  **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | id | ID of the test which status is to be modified | true | string | | | |

### GetResultsForTests

**API: SyntheticsDataService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| POST | /synthetics/v202309 /results | Returns probe results for a set of tests for specified period of time. |
| |  |  | | --- | --- | | **Request body:**  v202309GetResultsForTestsRequest | **Response body:**  v202309GetResultsForTestsResponse | | **Parameters**:  None |  | | | |

### GetTraceForTest

**API: SyntheticsDataService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| POST | /synthetics/v202309 /trace | Get network trace data for a specific synthetic test. The test must have traceroute task configured. |
| |  |  | | --- | --- | | **Request body:**  v202309GetTraceForTestRequest | **Response body:**  v202309GetTraceForTestResponse | | **Parameters:**  None |  | | | |

## Synthetics Schemas

This API uses the following schemas.

#### AgentMetadataIpValue

|  |  |
| --- | --- |
| **Schema:** AgentMetadataIpValue | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | Value | type: string | | |

#### protobufAny

|  |  |
| --- | --- |
| **Schema:** protobufAny | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | typeUrl | type: string | | value | type: string  format: byte | | |

#### rpcStatus

|  |  |
| --- | --- |
| **Schema:** rpcStatus | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | code | type: integer  format: int32 | | message | type: string | | details | type: array  items: $ref: protobufAny | | |

#### Location

|  |  |
| --- | --- |
| **Schema:** syntheticsv202309Location | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | latitude | type: number  format: double  description: Latitude in signed decimal degrees | | longitude | type: number  format: double  description: Longitude in signed decimal degrees | | country | type: string  description: Country of the location | | region | type: string  description: Geographic region within the country | | city | type: string  description: City of the location | | |

#### UserInfo

|  |  |
| --- | --- |
| **Schema:** v202303UserInfo | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | id | type: string  description: Unique system generated ID  readOnly: true | | email | type: string  description: E-mail address of the user  readOnly: true | | fullName | type: string  description: Full name of the user  readOnly: true | | |

#### ActivationSettings

|  |  |
| --- | --- |
| **Schema:** v202309ActivationSettings | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | gracePeriod | type: string  description: Period of healthy status in minutes within the time window not cancelling alarm activation | | timeUnit | type: string  description: Time unit for specifying time window (m | h) | | timeWindow | type: string  description: Time window for evaluating of test for alarm activation | | times | type: string  description: Number of occurrences of unhealthy test status within the time window triggering alarm activation | | |

#### Agent

|  |  |
| --- | --- |
| **Schema:** v202309Agent | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | id | type: string  description: Unique identifier of the agent  readOnly: true | | siteName | type: string  description: Name of the site where agent is located | | status | $ref: v202309AgentStatus | | alias | type: string  description: User selected descriptive name of the agent | | type | type: string  description: Type of agent (global | private)  readOnly: true | | os | type: string  description: OS version of server/VM hosting the agent  readOnly: true | | ip | type: string  description: Public IP address of the agent (auto-detected)  readOnly: true | | lat | type: number  format: double  description: Latitude of agent's location (signed decimal degrees) | | long | type: number  format: double  description: Longitude of agent's location (signed decimal degrees) | | lastAuthed | type: string  format: date-time  description: Timestamp of the last authorization  readOnly: true | | family | $ref: v202309IPFamily | | asn | type: integer  format: int64  description: ASN of the AS owning agent's public address | | siteId | type: string  description: ID of the site hosting the agent (if configured in Kentik) | | version | type: string  description: Software version of the agent  readOnly: true | | city | type: string  description: City where the agent is located | | region | type: string  description: Geographical region of agent's location | | country | type: string  description: Country of agent's location | | testIds | type: array  items: type: string  description: IDs of user's test running on the agent  readOnly: true | | localIp | type: string  description: Internal IP address of the agent | | cloudRegion | type: string  description: Cloud region (if any) hosting the agent | | cloudProvider | type: string  description: Cloud provider (if any) hosting the agent | | agentImpl | $ref: v202309ImplementType | | labels | type: array  items: type: string  description: List of names of labels associated with the agent | | metadata | $ref: v202309AgentMetadata | | |

#### AgentAlert

|  |  |
| --- | --- |
| **Schema:** v202309AgentAlert | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | id | type: string | | thresholdSeconds | type: integer  format: int64 | | notificationChannelIds | type: array  items: type: string | | agentId | type: string | | agentName | type: string | | |

#### AgentMetadata

|  |  |
| --- | --- |
| **Schema:** v202309AgentMetadata | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | privateIpv4Addresses | type: array  items: $ref: AgentMetadataIpValue  description: List of private IPv4 addresses | | publicIpv4Addresses | type: array  items: $ref: AgentMetadataIpValue  description: List of public IPv4 addresses  readOnly: true | | privateIpv6Addresses | type: array  items: $ref: AgentMetadataIpValue  description: List of private IPv6 addresses | | publicIpv6Addresses | type: array  items: $ref: AgentMetadataIpValue  description: List of public IPv6 addresses  readOnly: true | | |

#### AgentResults

|  |  |
| --- | --- |
| **Schema:** v202309AgentResults | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | agentId | type: string  description: ID of the agent providing results | | health | type: string  description: Overall health status of all task for the test executed by this agent | | tasks | type: array  items: $ref: v202309TaskResults  description: List of results for individual tasks | | |

#### AgentStatus

|  |  |
| --- | --- |
| **Schema:** v202309AgentStatus | **Type:** string |
| **Attributes**:  | Key | Value | | --- | --- | | enum | AGENT\_STATUS\_UNSPECIFIED, AGENT\_STATUS\_OK, AGENT\_STATUS\_WAIT, AGENT\_STATUS\_DELETED | | default | AGENT\_STATUS\_UNSPECIFIED | | description | •  AGENT\_STATUS\_UNSPECIFIED: Invalid value.  • AGENT\_STATUS\_OK: Agent is ready to accept tests  • AGENT\_STATUS\_WAIT: Agent is waiting for authorization  • AGENT\_STATUS\_DELETED: Agent was deleted • not user settable | | |

#### AgentTest

|  |  |
| --- | --- |
| **Schema:** v202309AgentTest | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | target | type: string  description: ID of the target agent | | useLocalIp | type: boolean  description: Boolean value indicating whether to use local (private) IP address of the target agent | | reciprocal | type: boolean  description: Boolean value indicating whether to make the test bidirectional | | |

#### CreateAgentAlertRequest

|  |  |
| --- | --- |
| **Schema:** v202309CreateAgentAlertRequest | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | thresholdSeconds | type: integer  format: int64 | | notificationChannelIds | type: array  items: type: string | | agentId | type: string | | |

#### CreateAgentAlertResponse

|  |  |
| --- | --- |
| **Schema:** v202309CreateAgentAlertResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | agentAlert | $ref: v202309AgentAlert | | |

#### CreateTestRequest

|  |  |
| --- | --- |
| **Schema:** v202309CreateTestRequest | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | test | $ref: v202309Test | | |

#### CreateTestResponse

|  |  |
| --- | --- |
| **Schema:** v202309CreateTestResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | test | $ref: v202309Test | | |

#### DNSRecord

|  |  |
| --- | --- |
| **Schema:** v202309DNSRecord | **Type:** string |
| **Attributes**:  | Key | Value | | --- | --- | | enum | DNS\_RECORD\_UNSPECIFIED, DNS\_RECORD\_A, DNS\_RECORD\_AAAA, DNS\_RECORD\_CNAME, DNS\_RECORD\_DNAME, DNS\_RECORD\_NS, DNS\_RECORD\_MX, DNS\_RECORD\_PTR, DNS\_RECORD\_SOA | | default | DNS\_RECORD\_UNSPECIFIED | | description | •  DNS\_RECORD\_UNSPECIFIED: Invalid value  • DNS\_RECORD\_A: name to IPv4 address(es) mapping  • DNS\_RECORD\_AAAA: name to IPv6 address(es) mapping  • DNS\_RECORD\_CNAME: alternative resource name  • DNS\_RECORD\_DNAME: alternative resource set name  • DNS\_RECORD\_NS: domain to name server mapping  • DNS\_RECORD\_MX: SMTP mail server record  • DNS\_RECORD\_PTR: IPv4/6 address to name mapping  • DNS\_RECORD\_SOA: domain meta-data | | |

#### DNSResponseData

|  |  |
| --- | --- |
| **Schema:** v202309DNSResponseData | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | status | type: integer  format: int64  description: Received DNS status | | data | type: string  description: Text rendering of received DNS resolution | | |

#### DNSResults

|  |  |
| --- | --- |
| **Schema:** v202309DNSResults | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | target | type: string  description: Queried DNS record | | server | type: string  description: DNS server used for the query | | latency | $ref: v202309MetricData | | response | $ref: v202309DNSResponseData | | |

#### DeleteAgentAlertResponse

|  |  |
| --- | --- |
| **Schema:** v202309DeleteAgentAlertResponse | **Type:** object |
| **Properties**:  None | |

#### DeleteAgentResponse

|  |  |
| --- | --- |
| **Schema:** v202309DeleteAgentResponse | **Type:** object |
| **Properties**:  None | |

#### DeleteTestResponse

|  |  |
| --- | --- |
| **Schema:** v202309DeleteTestResponse | **Type:** object |
| **Properties**:  None | |

#### DisabledMetrics

|  |  |
| --- | --- |
| **Schema:** v202309DisabledMetrics | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | pingLatency | type: boolean | | pingJitter | type: boolean | | pingPacketLoss | type: boolean | | httpLatency | type: boolean | | type: boolean | type: boolean | | httpCodes | type: boolean | | httpCertExpiry | type: boolean | | transactionLatency | type: boolean | | dnsLatency | type: boolean | | dnsCodes | type: boolean | | dnsIps | type: boolean | | |

#### DnsTest

|  |  |
| --- | --- |
| **Schema:** v202309DnsTest | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | target | type: string  description: Fully qualified DNS name to query | | timeout | type: integer  format: int64  description: --- Deprecated: value is ignored. --- | | recordType | $ref: v202309DNSRecord | | servers | type: array  items: type: string  description: List of IP addresses of DNS servers | | port | type: integer  format: int64  description: Target DNS server port | | |

#### FlowTest

|  |  |
| --- | --- |
| **Schema:** v202309FlowTest | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | target | type: string  description: Target ASN, CDN, Country, Region of City for autonomous test (type of value depends on flow test sub-type) | | targetRefreshIntervalMillis | type: integer  format: int64  description: Period (in milliseconds) for refreshing list of targets based on available flow data | | maxProviders | type: integer  format: int64  description: Maximum number of IP providers to track autonomously | | maxIpTargets | type: integer  format: int64  description: Maximum number of target IP addresses to select based flow data query | | type | type: string  description: Autonomous test sub-type (asn | cdn | country | region | city) | | inetDirection | type: string  description: Selection of address from flow data (src = source address in inbound flows | dst = destination addresses in outbound flows) | | direction | type: string  description: Direction of flows to match target attribute for extraction of target addresses (src | dst) | | |

#### GetAgentAlertResponse

|  |  |
| --- | --- |
| **Schema:** v202309GetAgentAlertResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | agentAlert | $ref: v202309AgentAlert | | |

#### GetAgentResponse

|  |  |
| --- | --- |
| **Schema:** v202309GetAgentResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | agent | $ref: v202309Agent | | |

#### GetResultsForTestsRequest

|  |  |
| --- | --- |
| **Schema:** v202309GetResultsForTestsRequest | **Type:** object |
| **Properties (\*  =  required)**  | Name | Value | | --- | --- | | ids \* | type: array  items: type: string  description: List of test IDs for which to retrieve results | | startTime \* | type: string  format: date-time  description: Timestamp of the oldest results to include in results | | endTime \* | type: string  format: date-time  description: Timestamp of the newest results to include in results | | agentIds | type: array  items: type: string  description: List of agent IDs from which to return results | | targets | type: array  items: type: string  description: List of targets (test dependent) for which to retrieve results | | aggregate | type: boolean  description: If true, retrieve result aggregated across the requested time period, else return complete time series | | |

#### GetResultsForTestsResponse

|  |  |
| --- | --- |
| **Schema:** v202309GetResultsForTestsResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | results | type: array  items: $ref: v202309TestResults | | |

#### GetTestResponse

|  |  |
| --- | --- |
| **Schema:** v202309GetTestResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | test | $ref: v202309Test | | |

#### GetTraceForTestRequest

|  |  |
| --- | --- |
| **Schema:** v202309GetTraceForTestRequest | **Type:** object |
| **Properties (\*  =  required)**  | Name | Value | | --- | --- | | id | type: string  description: ID of test for which to retrieve network path trace data | | startTime \* | type: string  format: date-time  description: Timestamp of the oldest results to include in results | | endTime \* | type: string  format: date-time  description: Timestamp of the newest results to include in results | | agentIds | type: array  items: type: string  description: List of agent IDs from which to return results | | targetIps | type: array  items: type: string  description: List of target IP addresses for which to retrieve results | | |

#### GetTraceForTestResponse

|  |  |
| --- | --- |
| **Schema:** v202309GetTraceForTestResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | nodes | type: object  additionalProperties: $ref: v202309NetNode  description: Map of network node information keyed by node IDs | | paths | type: array  items: $ref: v202309Path  description: List of retrieved network path data | | |

#### HTTPResponseData

|  |  |
| --- | --- |
| **Schema:** v202309HTTPResponseData | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | status | type: integer  format: int64  description: HTTP status in response | | size | type: integer  format: int64  description: Total size of received response body | | data | type: string  description: Detailed information about transaction timing, connection characteristics and response | | |

#### HTTPResults

|  |  |
| --- | --- |
| **Schema:** v202309HTTPResults | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | target | type: string  description: Target probed URL | | latency | $ref: v202309MetricData | | response | $ref: v202309HTTPResponseData | | dstIp | type: string  description: IP address of probed target server | | |

#### HealthSettings

|  |  |
| --- | --- |
| **Schema:** v202309HealthSettings | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | latencyCritical | type: number  format: float  description: Threshold for ping response latency (in microseconds) to trigger critical alarm | | latencyWarning | type: number  format: float  description: Threshold for ping response latency (in microseconds) to trigger warning alarm | | packetLossCritical | type: number  format: float  description: Threshold for ping packet loss (in %) to trigger critical alarm | | packetLossWarning | type: number  format: float  description: Threshold for ping packet loss (in %) to trigger warning alarm | | jitterCritical | type: number  format: float  description: Threshold for ping jitter (in microseconds) to trigger critical alarm | | jitterWarning | type: number  format: float  description: Threshold for ping jitter (in microseconds) to trigger critical alarm | | httpLatencyCritical | type: number  format: float  description: Threshold for HTTP response latency (in microseconds) to trigger critical alarm | | httpLatencyWarning | type: number  format: float  description: Threshold for HTTP response latency (in microseconds) to trigger warning alarm | | httpValidCodes | type: array  items: type: integer  format: int64  description: List of HTTP status codes indicating healthy state | | dnsValidCodes | type: array  items: type: integer  format: int64  description: List of DNS status codes indicating healthy state | | latencyCriticalStddev | type: number  format: float  description: Threshold for standard deviation (in microseconds) of ping response latency to trigger critical alarm | | latencyWarningStddev | type: number  format: float  description: Threshold for standard deviation (in microseconds) of ping response latency to trigger warning alarm | | jitterCriticalStddev | type: number  format: float  description: Threshold for standard deviation of ping jitter (in microseconds) to trigger critical alarm | | jitterWarningStddev | type: number  format: float  description: Threshold for standard deviation of ping jitter (in microseconds) to trigger warning alarm | | httpLatencyCriticalStddev | type: number  format: float  description: Threshold for standard deviation of HTTP response latency (in microseconds) to trigger critical alarm | | httpLatencyWarningStddev | type: number  format: float  description: Threshold for standard deviation of HTTP response latency (in microseconds) to trigger warning alarm | | unhealthySubtestThreshold | type: integer  format: int64  description: Number of tasks (across all agents) that must report unhealthy status in order for alarm to be triggered | | activation | $ref: v202309ActivationSettings | | certExpiryWarning | type: integer  format: int64  description: Threshold for remaining validity of TLS certificate (in days) to trigger warning alarm | | certExpiryCritical | type: integer  format: int64  description: Threshold for remaining validity of TLS certificate (in days) to trigger critical alarm | | dnsValidIps | type: string  description: Comma separated list of IP addresses expected to be received in response to DNS A or AAAA query | | dnsLatencyCritical | type: number  format: float  description: Threshold for DNS response latency (in microseconds) to trigger critical alarm | | dnsLatencyWarning | type: number  format: float  description: Threshold for DNS response latency (in microseconds) to trigger warning alarm | | dnsLatencyCriticalStddev | type: number  format: float  description: Threshold for standard deviation (in microseconds) of DNS response latency to trigger critical alarm | | dnsLatencyWarningStddev | type: number  format: float  description: Threshold for standard deviation (in microseconds) of DNS response latency to trigger warning alarm | | perAgentAlerting | type: boolean  description: Boolean value indicating whether to use per-agent alerting | | disabledMetrics | $ref: v202309DisabledMetrics | | healthDisabled | type: boolean  description: Disable all health evaluation for this test | | |

#### HostnameTest

|  |  |
| --- | --- |
| **Schema:** v202309HostnameTest | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | target | type: string  description: Fully qualified DNS name of the target host | | |

#### IPFamily

|  |  |
| --- | --- |
| **Schema:** v202309IPFamily | **Type:** string |
| **Attributes**:  | Key | Value | | --- | --- | | enum | IP\_FAMILY\_UNSPECIFIED, IP\_FAMILY\_V4, IP\_FAMILY\_V6, IP\_FAMILY\_DUAL | | default | IP\_FAMILY\_UNSPECIFIED | | description | • IP\_FAMILY\_UNSPECIFIED: Invalid value.  • IP\_FAMILY\_V4: IPv4 only  • IP\_FAMILY\_V6: IPv6 only  • IP\_FAMILY\_DUAL: IPv4 and IPv6 supported | | |

#### ImplementType

|  |  |
| --- | --- |
| **Schema:** v202309ImplementType | **Type:** string |
| **Attributes**:  | Key | Value | | --- | --- | | enum | IMPLEMENT\_TYPE\_UNSPECIFIED, IMPLEMENT\_TYPE\_RUST, IMPLEMENT\_TYPE\_NODE, IMPLEMENT\_TYPE\_NETWORK | | default | IMPLEMENT\_TYPE\_UNSPECIFIED | | description | • IMPLEMENT\_TYPE\_RUST: ksynth, a.k.a network agent (implemented in Rust) capable of running all tasks except for page-load and transaction  • IMPLEMENT\_TYPE\_NODE: ksynth-agent, a.k.a. app agent (implemented in NodeJS) with Chromium enabled capable of running all tasks  • IMPLEMENT\_TYPE\_NETWORK: ksynth-agent, a.k.a. app agent with Chromium disabled, capable of running all tasks except for page-load and transaction | | |

#### IpTest

|  |  |
| --- | --- |
| **Schema:** v202309IpTest | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | targets | type: array  items: type: string  description: List of IP addresses of targets | | useLocalIp | type: boolean  description: Boolean value indicating whether to use local (private) IP address of the target agents | | |

#### ListAgentAlertsResponse

|  |  |
| --- | --- |
| **Schema:** v202309ListAgentAlertsResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | agentAlerts | type: array  items: $ref: v202309AgentAlert | | |

#### ListAgentsResponse

|  |  |
| --- | --- |
| **Schema:** v202309ListAgentsResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | agents | type: array  items: $ref: v202309Agent  description: List of available agents | | invalidCount | type: integer  format: int64  description: Number of invalid entries encountered while collecting data | | |

#### ListTestsResponse

|  |  |
| --- | --- |
| **Schema:** v202309ListTestsResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | tests | type: array  items: $ref: v202309Test  description: List of configured active or paused tests | | invalidCount | type: integer  format: int64  description: Number of invalid entries encountered while collecting data | | |

#### MetricData

|  |  |
| --- | --- |
| **Schema:** v202309MetricData | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | current | type: integer  format: int64  description: Current value of metric | | rollingAvg | type: integer  format: int64  description: Rolling average of metric | | rollingStddev | type: integer  format: int64  description: Rolling average of standard deviation of metric | | health | type: string  description: Health evaluation status for the metric (healthy | warning | critical) | | |

#### NetNode

|  |  |
| --- | --- |
| **Schema:** v202309NetNode | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | ip | type: string  description: IP address of the node in standard textual notation | | asn | type: integer  format: int64  description: AS number owning the address of the node | | asName | type: string  description: Name of the AS owning the address of the node | | location | $ref: syntheticsv202309Location | | dnsName | type: string  description: DNS name of the node (obtained by reverse DNS resolution) | | deviceId | type: string  description: ID of the device corresponding with the node in Kentik configuration | | siteId | type: string  description: ID of the site containing the device corresponding with the node in Kentik configuration | | |

#### NetworkMeshTest

|  |  |
| --- | --- |
| **Schema:** v202309NetworkMeshTest | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | useLocalIp | type: boolean  description: Boolean value indicating whether to use local (private) IP address of the target agents | | |

#### PacketLossData

|  |  |
| --- | --- |
| **Schema:** v202309PacketLossData | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | current | type: number  format: double  description: Current packet loss value | | health | type: string  description: Health evaluation status for the metric (healthy | warning | critical) | | |

#### PageLoadTest

|  |  |
| --- | --- |
| **Schema:** v202309PageLoadTest | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | target | type: string  description: HTTP or HTTPS URL to request | | timeout | type: integer  format: int64  description: HTTP transaction timeout (in milliseconds) | | headers | type: object  additionalProperties: type: string  description: Map of HTTP header values keyed by header names | | ignoreTlsErrors | type: boolean  description: Boolean indicating whether to ignore TLS certificate verification errors | | cssSelectors | type: object  additionalProperties: type: string  description: Map of CSS selector values keyed by selector name | | |

#### Path

|  |  |
| --- | --- |
| **Schema:** v202309Path | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | agentId | type: string  description: ID of the agent generating the path data | | targetIp | type: string  description: IP address of the target of the path | | hopCount | $ref: v202309Stats | | maxAsPathLength | type: integer  format: int32  description: Maximum length of AS path across all traces | | traces | type: array  items: $ref: v202309PathTrace  description: Data for individual traces | | time | type: string  format: date-time  description: Timestamp (UTC) of initiation of the path trace | | |

#### PathTrace

|  |  |
| --- | --- |
| **Schema:** v202309PathTrace | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | asPath | type: array  items: type: integer  format: int32  description: AS path of the network trace | | isComplete | type: boolean  description: Indication whether response from target was received | | hops | type: array  items: $ref: v202309TraceHop  description: List of hops in the trace | | |

#### PingResults

|  |  |
| --- | --- |
| **Schema:** v202309PingResults | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | target | type: string  description: Hostname or address of the probed target | | packetLoss | $ref: v202309PacketLossData | | latency | $ref: v202309MetricData | | jitter | $ref: v202309MetricData | | dstIp | type: string  description: IP address of probed target | | |

#### SetTestStatusRequest

|  |  |
| --- | --- |
| **Schema:** v202309SetTestStatusRequest | **Type:** object |
| **Properties (\*  =  required)**  | Name | Value | | --- | --- | | id \* | type: string  description: ID of the test which status is to be modified | | status | $ref: v202309TestStatus | | |

#### SetTestStatusResponse

|  |  |
| --- | --- |
| **Schema:** v202309SetTestStatusResponse | **Type:** object |
| **Properties**: None. | |

#### Stats

|  |  |
| --- | --- |
| **Schema:** v202309Stats | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | average | type: integer  format: int32  description: Average value | | min | type: integer  format: int32  description: Minimum value | | max | type: integer  format: int32  description: Maximum value | | |

#### TaskResults

|  |  |
| --- | --- |
| **Schema:** v202309TaskResults | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | ping | $ref: v202309PingResults | | http | $ref: v202309HTTPResults | | dns | $ref: v202309DNSResults | | health | type: string  description: Health status of the task | | |

#### Test

|  |  |
| --- | --- |
| **Schema:** v202309Test | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | id | type: string  description: Unique ID of the test  readOnly: true | | name | type: string  description: User selected name of the test | | type | type: string  description: Type of the test | | status | $ref: v202309TestStatus | | settings | $ref: v202309TestSettings | | cdate | type: string  format: date-time  description: Creation timestamp (UTC)  readOnly: true | | edate | type: string  format: date-time  description: Last modification timestamp (UTC)  readOnly: true | | createdBy | $ref: v202303UserInfo | | lastUpdatedBy | $ref: v202303UserInfo | | labels | type: array  items: type: string  description: Set of labels associated with the test | | |

#### TestPingSettings

|  |  |
| --- | --- |
| **Schema:** v202309TestPingSettings | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | count | type: integer  format: int64  description: Number of probe packets to send in one iteration | | protocol | type: string  description: Transport protocol to use (icmp | tcp) | | port | type: integer  format: int64  description: Target port for TCP probes (ignored for ICMP) | | timeout | type: integer  format: int64  description: Timeout in milliseconds for execution of the task | | delay | type: number  format: float  description: Inter-probe delay in milliseconds | | dscp | type: integer  format: int64  description: DSCP code to be set in IP header of probe packets | | |

#### TestResults

|  |  |
| --- | --- |
| **Schema:** v202309TestResults | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | testId | type: string  description: ID of the test for which results are provided | | time | type: string  format: date-time  description: Results timestamp (UTC) | | health | type: string  description: Health status of the test | | agents | type: array  items: $ref: v202309AgentResults  description: List of results from agents executing tasks on behalf of the test | | |

#### TestSettings

|  |  |
| --- | --- |
| **Schema:** v202309TestSettings | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | hostname | $ref: v202309HostnameTest | | ip | $ref: v202309IpTest | | agent | $ref: v202309AgentTest | | flow | $ref: v202309FlowTest | | dns | $ref: v202309DnsTest | | url | $ref: v202309UrlTest | | networkGrid | $ref: v202309IpTest | | pageLoad | $ref: v202309PageLoadTest | | dnsGrid | $ref: v202309DnsTest | | networkMesh | $ref:v202309NetworkMeshTest | | agentIds | type: array  items: type: string  description: IDs of agents assigned to run tasks on behalf of the test | | tasks | type: array  items: type: string  description: List of task names to run for the test | | healthSettings | $ref: v202309HealthSettings | | ping | $ref: v202309TestPingSettings | | trace | $ref: v202309TestTraceSettings | | period | type: integer  format: int64  description: Test evaluation period (in seconds) | | family | $ref: v202309IPFamily | | notificationChannels | type: array  items: type: string  description: List of IDs of notification channels for alarms triggered by the test | | notes | type: string  description: Add a note or comment for this test | | |

#### TestStatus

|  |  |
| --- | --- |
| **Schema:** v202309TestStatus | **Type:** string |
| **Attributes**:  | Name | Value | | --- | --- | | enum | TEST\_STATUS\_UNSPECIFIED, TEST\_STATUS\_ACTIVE, TEST\_STATUS\_PAUSED, TEST\_STATUS\_DELETED, TEST\_STATUS\_PREVIEW | | default | TEST\_STATUS\_UNSPECIFIED | | description | • TEST\_STATUS\_UNSPECIFIED: Invalid value.  • TEST\_STATUS\_ACTIVE: Test is active.  • TEST\_STATUS\_PAUSED: Test is paused.  • TEST\_STATUS\_DELETED: Test is deleted. Not user settable.  • TEST\_STATUS\_PREVIEW: Test is preview | | |

#### TestTraceSettings

|  |  |
| --- | --- |
| **Schema:** v202309TestTraceSettings | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | count | type: integer  format: int64  description: Number of probe packets to send in one iteration | | protocol | type: string  description: Transport protocol to use (icmp | tcp | udp) | | port | type: integer  format: int64  description: Target port for TCP or UDP probes (ignored for ICMP) | | timeout | type: integer  format: int64  description: Timeout in milliseconds for execution of the task | | limit | type: integer  format: int64  description: Maximum number of hops to probe (i.e. maximum TTL) | | delay | type: number  format: float  description: Inter-probe delay in milliseconds | | dscp | type: integer  format: int64  description: DSCP code to be set in IP header of probe packets | | |

#### TraceHop

|  |  |
| --- | --- |
| **Schema:** v202309TraceHop | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | latency | type: integer  format: int32  description: Round-trip packet latency to the node (in microseconds) - 0 if no response was received | | nodeId | type: string  description: ID of the node for this hop in the Nodes map - empty if no response was received | | |

#### UpdateAgentAlertRequest

|  |  |
| --- | --- |
| **Schema:** v202309UpdateAgentAlertRequest | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | id | type: string | | thresholdSeconds | type: integer  format: int64 | | notificationChannelIds | type: array  items: type: string | | |

#### UpdateAgentAlertResponse

|  |  |
| --- | --- |
| **Schema:** v202309UpdateAgentAlertResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | agentAlert | $ref: v202309AgentAlert | | |

#### UpdateAgentRequest

|  |  |
| --- | --- |
| **Schema:** v202309UpdateAgentRequest | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | agent | $ref: v202309Agent | | |

#### UpdateAgentResponse

|  |  |
| --- | --- |
| **Schema:** v202309UpdateAgentResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | agent | $ref: v202309Agent | | |

#### UpdateTestRequest

|  |  |
| --- | --- |
| **Schema:** v202309UpdateTestRequest | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | test | $ref: v202309Test | | |

#### UpdateTestResponse

|  |  |
| --- | --- |
| **Schema:** v202309UpdateTestResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | test | $ref: v202309Test | | |

#### UrlTest

|  |  |
| --- | --- |
| **Schema:** v202309UrlTest | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | target | type: string  description: HTTP or HTTPS URL to request | | timeout | type: integer  format: int64  description: HTTP transaction timeout (in milliseconds) | | method | type: string  description: HTTP method to use (GET | HEAD | PATCH | POST | PUT) | | headers | type: object  additionalProperties: type: string  description: Map of HTTP header values keyed by header names | | body | type: string  description: HTTP request body | | ignoreTlsErrors | type: boolean  description: Boolean indicating whether to ignore TLS certificate verification errors | | |
