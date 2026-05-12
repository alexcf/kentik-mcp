# Connectivity Cost APIs

Source: https://kb.kentik.com/docs/connectivity-cost-apis

---

This article covers how to get started with the Connectivity Cost APIs.

> **Notes:**
>
> * The **API Tester** in the v4 Portal enables you to securely call the methods of these APIs using your organization's Kentik data.
> * For information about connectivity costs, start with **Connectivity Costs**.
> * Protobuf and OpenAPI specifications for Kentik’s v6 APIs are available in our **api-schema-public** repository.

## Connectivity Cost Usage

The Connectivity Cost Configuration API provides programmatic access to configuration of Connectivity Costs.

## Connectivity Cost RPCs

The tables below provide a quick reference to key information about each Remote Procedure Call in these APIs. Use the links in the tables to see schemas for the request and/or response body (if any).

> **Note:** To test methods using your own Kentik data, use the portal's **API Tester**.

### ListCostProviders

**API: CostService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| GET | /cost/v202308 /cost/providers | Returns list of configured cost providers. |
| |  |  | | --- | --- | | **Request body**: None **Parameters**: None | **Response body**:  v202308ListCostProvidersResponse | | | |

### ListCostProviderSummaries

**API: CostService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| GET | /cost/v202308 /cost/summary | Returns list of summaries of configured cost providers. |
| |  |  | | --- | --- | | **Request body**: None | **Response body**:  v202308ListCostProviderSummariesResponse |    **Parameters:**  | Name | Description | Required | Type | | --- | --- | --- | --- | | date | Date of the cost provider summary (YYYY-MM) | false | string | | | |

### GetCostProviderSummary

**API: CostService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| GET | /cost/v202308 /cost/summary/{id} | Returns summary of configured cost provider. |
| |  |  | | --- | --- | | **Request body:** None | **Response body**:  v202308GetCostProviderSummaryResponse |  **Parameters:**  | Name | Description | Required | Type | | --- | --- | --- | --- | | id | ID of the cost provider (can be found using ListCostProviders RPC) | true | string | | date | Date of the cost provider summary (YYYY-MM) | false | string | | | |

## Connectivity Cost Schemas

This API uses the following schemas.

#### Status

|  |  |
| --- | --- |
| **Schema:** costv202308Status | **Type:** string |
| **Attributes:**  | Key | Value | | --- | --- | | enum | STATUS\_UNSPECIFIED, STATUS\_INCOMPLETE, STATUS\_COMPLETE | | default | STATUS\_UNSPECIFIED | | |

#### googlerpcStatus

|  |  |
| --- | --- |
| **Schema:** googlerpcStatus | **Type:** object |
| **Properties:**  | Name | Value | | --- | --- | | code | type: integer  format: int32 | | message | type: string | | details | type: array  items: $ref: protobufAny | | |

#### protobufAny

|  |  |
| --- | --- |
| **Schema:** protobufAny | **Type:** object |
| **Properties:**  | Name | Value | | --- | --- | | typeUrl | type: string | | value | type: string  format: byte | | |

#### CostProviderConcise

|  |  |
| --- | --- |
| **Schema:** v202308CostProviderConcise | **Type:** object |
| **Properties (\*  =  required)**  | Name | Value | | --- | --- | | id | type: string  description: ID of the cost provider  readOnly: true | | name \* | type: string  description: Name of the cost provider | | |

#### CostProviderSummary

|  |  |
| --- | --- |
| **Schema:** v202308CostProviderSummary | **Type:** object |
| **Properties:**  | Name | Value | | --- | --- | | date | type: string  description: Date of the cost provider summary (YYYY-MM) | | status | $ref: costv202308Status | | totalCost | type: number  format: float  description: Total cost of the cost provider | | totalCostVariation | type: string  description: Total cost percent variation of the cost provider (percentage changed vs the previous month) | | totalCostGroupAdditionalCost | type: number  format: float  description: Total cost group additional cost of the cost provider | | totalCostGroupAdditionalInterfaceCost | type: number  format: float  description: Total cost group additional interface cost of the cost provider | | currency | type: string  description: Currency (ISO 4217) used for cost values of the cost provider | | costPerMbps | type: number  format: float  description: Cost per mbps (million bits per second) of the cost provider | | costPerMbpsVariation | type: string  description: Cost per mbps percent variation of the cost provider (percentage changed vs the previous month) | | providerName | type: string  description: Provider name of the cost provider | | costGroupName | type: string  description: Cost group name of the cost provider | | costGroupConnType | type: string  description: Cost group connection type of the cost provider | | siteName | type: string  description: Site name of the cost provider | | siteMarket | type: string  description: Site market of the cost provider | | ingressTrafficMbps | type: number  format: float  description: Ingress traffic mbps (million bits per second) of the cost provider | | ingressTrafficVariation | type: string  description: Ingress traffic percent variation of the cost provider (percentage changed vs the previous month) | | egressTrafficMbps | type: number  format: float  description: Egress traffic mbps (million bits per second) of the cost provider | | egressTrafficVariation | type: string  description: Egress traffic percent variation of the cost provider (percentage changed vs the previous month) | | |

#### GetCostProviderSummaryResponse

|  |  |
| --- | --- |
| **Schema:** v202308GetCostProviderSummaryResponse | **Type:** object |
| **Properties:**  | Name | Value | | --- | --- | | providers | type: array  items: $ref: v202308CostProviderSummary  description: List of summaries of requested cost providers | | invalidCount | type: integer  format: int64  description: Number of invalid entries encountered while collecting data | | |

#### ListCostProviderSummariesResponse

|  |  |
| --- | --- |
| **Schema:** v202308ListCostProviderSummariesResponse | **Type:** object |
| **Properties:**  | Name | Value | | --- | --- | | providers | type: array  items: $ref: v202308CostProviderSummary  description: List of summaries of requested cost providers | | invalidCount | type: integer  format: int64  description: Number of invalid entries encountered while collecting data | | |

#### ListCostProvidersResponse

|  |  |
| --- | --- |
| **Schema:** v202308ListCostProvidersResponse | **Type:** object |
| **Properties:**  | Name | Value | | --- | --- | | providers | type: array  items: $ref: v202308CostProviderConcise  description: List of configurations of requested cost providers | | invalidCount | type: integer  format: int64  description: Number of invalid entries encountered while collecting data | | |
