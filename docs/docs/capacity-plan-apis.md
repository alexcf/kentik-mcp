# Capacity Plan APIs

Source: https://kb.kentik.com/docs/capacity-plan-apis

---

This article covers how to get started with the Capacity Plan APIs.

> **Notes:**
>
> * The **API Tester** in the v4 Portal enables you to securely call the methods of these APIs using your organization's Kentik data.
> * For information about capacity planning in Kentik, start with **Capacity Planning**.
> * Protobuf and OpenAPI specifications for Kentik’s v6 APIs are available in our **api-schema-public** repository.

## Capacity Plan Usage

The Capacity Plan API provides read-only access to configured capacity plans.

## Capacity Plan RPCs

The tables below provide a quick reference to key information about each Remote Procedure Call in these APIs. Use the links in the tables to see schemas for the request and/or response body (if any).

> **Note:** To test methods using your own Kentik data, use the portal's **API Tester**.

### ListCapacityPlans

**API: CapacityPlanService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| GET | /capacity\_plan/v202212 /capacity\_plan | Returns list of capacity plans. |
| |  |  | | --- | --- | | **Request body**: None **Parameters**:  None | **Response body**:  v202212ListCapacityPlansResponse | | | |

### ListCapacitySummaries

**API: CapacityPlanService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| GET | /capacity\_plan/v202212 /capacity\_plan/summary | Returns list of capacity summaries. |
| |  |  | | --- | --- | | **Request body:**  None  **Parameters:**  None | **Response body:**  v202212ListCapacitySummariesResponse | | | |

### GetCapacityPlan

**API: CapacityPlanService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| GET | /capacity\_plan/v202212 /capacity\_plan/{id} | Returns capacity plan specified by ID. |
| |  |  | | --- | --- | | **Request body:**  None | **Response body:**  v202212GetCapacityPlanResponse |    **Parameters:**  | Name | Description | Required | Type | | --- | --- | --- | --- | | id | ID of the requested capacity plan | true | string | | | |

### GetCapacitySummary

**API: CapacityPlanService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| GET | /capacity\_plan/v202212 /capacity\_plan/{id}/summary | Returns capacity plan summary specified by ID. |
| |  |  | | --- | --- | | **Request body:**  None | **Response body:**  v202212GetCapacitySummaryResponse |    **Parameters:**  | Name | Description | Required | Type | | --- | --- | --- | --- | | id | ID of the requested capacity plan summary | true | string | | | |

## Capacity Plan Schemas

This API uses the following schemas.

#### InterfaceDetail

|  |  |
| --- | --- |
| **Schema:** CapacityPlanInterfaceDetail | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | deviceName | type: string  description: Device name  readOnly: true | | intfName | type: string  description: Interface name  readOnly: true | | intfDescription | type: string  description: Interface description  readOnly: true | | intfCapacity | type: string  description: Interface capacity  readOnly: true | | networkBoundary | type: string  description: Network boundary  readOnly: true | | connType | type: string  description: Connectivity type  readOnly: true | | provider | type: string  description: Provider  readOnly: true | | utilStatus | type: string  description: Utilization status  readOnly: true | | utilOutMbps | type: string  description: Utilization out mbps  readOnly: true | | utilOutPct | type: string  description: Utilization out percentage  readOnly: true | | utilInMbps | type: string  description: Utilization in mbps  readOnly: true | | utilInPct | type: string  description: Utilization in percentage  readOnly: true | | runoutStatus | type: string  description: Runout status  readOnly: true | | runoutInDate | type: string  description: Runout in date  readOnly: true | | runoutInVariation | type: string  description: Runout in variation  readOnly: true | | runoutOutDate | type: string  description: Runout out date  readOnly: true | | runoutOutVariation | type: string  description: Runout out variation  readOnly: true | | |

#### InterfacesDetail

|  |  |
| --- | --- |
| **Schema:** CapacitySummaryInterfacesDetail | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | totalCount | type: integer  format: int64  description: Total number of interfaces  readOnly: true | | totalCapacityBps | type: string  format: uint64  description: Total capacity in bps  readOnly: true | | healthy | $ref: InterfacesDetail StatusDetail | | warning | $ref: InterfacesDetail StatusDetail | | critical | $ref: InterfacesDetail StatusDetail | | |

#### RunoutConfig

|  |  |
| --- | --- |
| **Schema:** ConfigRunoutConfig | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | strategy | type: string  description: Strategy for runout  readOnly: true | | warnQty | type: integer  format: int64  description: Warning quantity for runout  readOnly: true | | critQty | type: integer  format: int64  description: Critical quantity for runout  readOnly: true | | |

#### UtilConfig

|  |  |
| --- | --- |
| **Schema:** ConfigUtilConfig | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | aggregate | type: string  description: Aggregate for utilization  readOnly: true | | warnPct | type: integer  format: int64  description: Warning percentage for utilization  readOnly: true | | critPct | type: integer  format: int64  description: Critical percentage for utilization  readOnly: true | | |

#### StatusDetail

|  |  |
| --- | --- |
| **Schema:** InterfacesDetailStatusDetail | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | bps | type: string  format: uint64  description: Bandwidth in bps  readOnly: true | | count | type: integer  format: int64  description: Number of interfaces  readOnly: true | | |

#### RunoutStatus

|  |  |
| --- | --- |
| **Schema:** SummaryStatusRunoutStatus | **Type:** object |
| **Properties:**  | Name | Value | | --- | --- | | status | type: string  description: Status of runout  readOnly: true | | earliestDate | type: string  format: uint64  description: Earliest runout date  readOnly: true | | |

#### UtilStatus

|  |  |
| --- | --- |
| **Schema:** SummaryStatusUtilStatus | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | status | type: string  description: Status of utilization  readOnly: true | | highestPct | type: integer  format: int64  description: Highest utilization percentage  readOnly: true | | |

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

#### CapacityPlan

|  |  |
| --- | --- |
| **Schema:** v202212CapacityPlan | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | id | type: string  description: ID of capacity plan  readOnly: true | | name | type: string  description: Name of capacity plan  readOnly: true | | description | type: string  description: Description of capacity plan  readOnly: true | | status | type: string  description: Status of capacity plan  readOnly: true | | interfaces | type: array  items: $ref: CapacityPlan InterfaceDetail  description: List of interfaces  readOnly: true | | config | $ref: v202212Config | | summaryStatus | $ref: v202212SummaryStatus | | |

#### CapacitySummary

|  |  |
| --- | --- |
| **Schema:** v202212CapacitySummary | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | id | type: string  description: ID of capacity plan  readOnly: true | | name | type: string  description: Name of capacity plan  readOnly: true | | description | type: string  description: Description of capacity plan  readOnly: true | | status | type: string  description: Status of capacity plan  readOnly: true | | interfaces | $ref: CapacitySummary InterfacesDetail | | config | $ref: v202212Config | | summaryStatus | $ref: v202212SummaryStatus | | |

#### Config

|  |  |
| --- | --- |
| **Schema:** v202212Config | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | runout | $ref: ConfigRunoutConfig | | utilization | $ref: ConfigUtilConfig | | |

#### GetCapacityPlanResponse

|  |  |
| --- | --- |
| **Schema:** v202212GetCapacityPlanResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | capacity | $ref: v202212CapacityPlan | | |

#### GetCapacitySummaryResponse

|  |  |
| --- | --- |
| **Schema:** v202212GetCapacitySummaryResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | capacity | $ref: v202212CapacitySummary | | |

#### ListCapacityPlansResponse

|  |  |
| --- | --- |
| **Schema:** v202212ListCapacityPlansResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | capacity | type: array  items: $ref: v202212CapacityPlan  description: List of capacity plans | | invalidCount | type: integer  format: int64  description: Number of invalid entries encountered while collecting data | | |

#### ListCapacitySummariesResponse

|  |  |
| --- | --- |
| **Schema:** v202212ListCapacitySummariesResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | capacity | type: array  items: $ref: v202212CapacitySummary  description: List of capacity plan summaries | | invalidCount | type: integer  format: int64  description: Number of invalid entries encountered while collecting data | | |

#### SummaryStatus

|  |  |
| --- | --- |
| **Schema:** v202212SummaryStatus | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | runout | $ref:SummaryStatus RunoutStatus | | utilization | $ref:SummaryStatus RunoutStatus | | |
