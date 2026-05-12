# Label APIs

Source: https://kb.kentik.com/docs/label-apis

---

This article covers how to get started with the Label APIs.

> **Notes:**
>
> * The **API Tester** in the v4 Portal enables you to securely call the methods of these APIs using your organization's Kentik data.
> * For information about labels, start with **About Labels**.
> * Protobuf and OpenAPI specifications for Kentik’s v6 APIs are available in our **api-schema-public** repository.

## Label Usage

The topics below provide important background information for the use of these APIs.

### Overview

The Label API enables programmatic creation and management of labels, which are tags that can be applied to objects such as devices, synthetic tests, and ksynth agents in order to create logical groups. While this API is used to manage labels, the application of a label to a given object is done with the API corresponding to the type of that object (see Applying Labels via API).

Both gRPC RPCs and REST endpoints are provided.

### Applying Labels via API

The Label API does not handle the application of labels to configuration objects. Instead, a label is applied to an object using the management API for that object. The table below lists the APIs for the types of objects to which labels may be applied.

| Object type | **API for attaching labels** |
| --- | --- |
| Device | **Device Apply Labels** |
| Synthetic monitoring test | **SyntheticsAdminService API** |
| Synthetic monitoring agent | **SyntheticsAdminService API** |
| BGP monitor | **BgpMonitoringAdminService API** |

## Label RPCs

The tables below provide a quick reference to key information about each Remote Procedure Call in these APIs. Use the links in the tables to see schemas for the request and/or response body (if any).

> **Note:** To test methods using your own Kentik data, use the portal's **API Tester**.

### ListLabels

**API: LabelService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| GET | /label/v202210 /labels | Returns list of all labels configured in the account. |
| |  |  | | --- | --- | | **Request body**: None  **Parameters**: None | **Response body**:  v202210ListLabelsResponse | | | |

### CreateLabel

**API: LabelService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| POST | /label/v202210 /labels | Creates a new label based on data in the request. |
| |  |  | | --- | --- | | **Request body:** v202210CreateLabelRequest **Parameters**: None | **Response body:** v202210CreateLabelResponse | | | |

### DeleteLabel

**API: LabelService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| DELETE | /label/v202210 /labels/{id} | Deletes label with specified with id. |
| |  |  | | --- | --- | | **Request body**: None | **Response body**:  v202210DeleteLabelResponse |  **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | id | ID of the label to be deleted | true | string | | | |

### UpdateLabel

**API: LabelService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| POST | /label/v202210 /labels/{id} | Updates configuration of a label. |
| |  |  | | --- | --- | | **Request body**: v202210UpdateLabelRequest | **Response body:**  v202210UpdateLabelResponse |  **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | id | Unique system assigned identifier of the label | true | string | | | |

## Label Schemas

This API uses the following schemas.

#### Label

|  |  |
| --- | --- |
| **Schema:** labelv202210Label | **Type:** object |
| **Properties** **(\*  =  required)**  | Name | Value | | --- | --- | | id | type: string  description: Unique system assigned identifier of the label  readOnly: true | | name \* | type: string  description: Label text visible in listing of configuration of objects to which it has been applied | | description | type: string  description: Optional description of the label, visible only in API responses | | color | type: string  description: [Hexadecimal code of the color](https://www.color-hex.com/) of the label text background in the portal | | cdate | type: string  format: date-time  description: Creation timestamp (UTC)  readOnly: true | | edate | type: string  format: date-time  description: Last modification timestamp (UTC)  readOnly: true | | |

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

#### CreateLabelRequest

|  |  |
| --- | --- |
| **Schema:** v202210CreateLabelRequest | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | label | $ref: labelv202210Label | | |

#### CreateLabelResponse

|  |  |
| --- | --- |
| **Schema:** v202210CreateLabelResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | label | $ref: labelv202210Label | | |

#### DeleteLabelResponse

|  |  |
| --- | --- |
| **Schema:** v202210DeleteLabelResponse | **Type:** object |
| **Properties**: None. | |

#### ListLabelsResponse

|  |  |
| --- | --- |
| **Schema:** v202210ListLabelsResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | labels | type: array  items: $ref: labelv202210Label  description: List of configured labels  readOnly: true | | invalidCount | type: integer  format: int32  description: Number of invalid entries encountered while collecting data (should be always 0) | | |

#### UpdateLabelRequest

|  |  |
| --- | --- |
| **Schema:** v202210UpdateLabelRequest | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | label | $ref: labelv202210Label | | |

#### UpdateLabelResponse

|  |  |
| --- | --- |
| **Schema:** v202210UpdateLabelResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | label | $ref: labelv202210Label | | |
