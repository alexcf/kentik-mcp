# Flow Tag APIs

Source: https://kb.kentik.com/docs/flow-tag-apis

---

> **Note:** This alpha API is neither supported nor recommended for production use. For additional information, please contact **Product Support**.

This article covers how to get started with the Flow Tag APIs.

> **Notes:**
>
> * The **API Tester** in the v4 Portal enables you to securely call the methods of these APIs using your organization's Kentik data.
> * For information about flow tags, start with **Flow Tags**.
> * Protobuf and OpenAPI specifications for Kentik’s v6 APIs are available in our **api-schema-public** repository.

## Flow Tag Usage

The topics below provide important background information for the use of these APIs.

### Overview

The Flow Tag API enables programmatic creation and management of flow tags.

* In addition to tag name, at least one other field must be specified to create a tag.
* All specified tag fields are ANDed for evaluation.
* Except for Tag Name, any field may contain multiple comma-separated values, which will be ORed for evaluation.
* A tag will be applied to a flow only when all of the fields specified for that tag are matched. Source and Destination flows are each evaluated independently for matches.
* Source and Destination flows are each evaluated independently for matches.
* A match results in the addition of the tag name to a delimited list of tags in the src\_flow\_tags and/or dst\_flow\_tags column of each specified devices KDE main table.
* Tags in a **KDE table** can be searched as part of a query. Tag searches are substring-based. Query results vary depending on how tags are named.
* Additional information on the values of individual fields may be found in **Tag Field Definitions** in the Kentik Knowledge Base.

## Flow Tag RPCs

The tables below provide a quick reference to key information about each Remote Procedure Call in these APIs. Use the links in the tables to see schemas for the request and/or response body (if any).

> **Note:** To test methods using your own Kentik data, use the portal's **API Tester**.

### SearchFlowTag

**API: FlowTagService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| GET | /flow\_tag/v202404alpha1 /tag | Returns configuration of flow tag with search parameters. |
| |  |  | | --- | --- | | **Request body**:  None | **Response body**:  v202404alpha1SearchFlowTagResponse |    **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | search.limit | Limit the number of rows to fetch. | false | integer | | search.offset | The number rows to skip before returning. | false | integer | | search.lookupFields | List of lookup fields.   - LOOKUP\_FIELD\_UNSPECIFIED: Invalid value  - LOOKUP\_FIELD\_VALUE: Lookup by name of tag | false | array | | search.lookupValues | List of lookup values. | false | array | | search.fieldLimit | Limit the number of record to return for nested fields. | false | integer | | | |

### CreateFlowTag

**API: FlowTagService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| POST | /flow\_tag/v202404alpha1 /tag | Create a flow tag configuration. |
| |  |  | | --- | --- | | **Request body**:  v202404alpha1CreateFlowTagRequest | **Response body**:  v202404alpha1CreateFlowTagResponse | | **Parameters**:  None | | | | |

### UpdateFlowTag

**API: FlowTagService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| PUT | /flow\_tag/v202404alpha1 /tag/{flowTag.id} | Update a flow tag configuration. |
| |  |  | | --- | --- | | **Request body**:  v202404alpha1UpdateFlowTagRequest | **Response body**:  v202404alpha1UpdateFlowTagResponse |  **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | flowTag.id | Unique system assigned identifier of the flow tag | true | string | | | |

### GetFlowTag

**API: FlowTagService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| GET | /flow\_tag/v202404alpha1 /tag/{flowTag.id} | Returns configuration of flow tag with specified ID. |
| |  |  | | --- | --- | | **Request body**:  None | **Response body**:  v202404alpha1GetFlowTagResponse |    **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | flowTag.id | Undefined. | true | string | | | |

### DeleteFlowTag

**API: FlowTagService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| DELETE | /flow\_tag/v202404alpha1 /tag/{flowTag.id} | Delete a flow tag configuration with id. |
| |  |  | | --- | --- | | **Request body**:  None | **Response body**:  v202404alpha1DeleteFlowTagResponse |  **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | flowTag.id | Undefined. | true | string | | | |

## Flow Tag Schemas

This API uses the following schemas.

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

#### AddressInfo

|  |  |
| --- | --- |
| **Schema:** v202404alpha1AddressInfo | **Type:** object |
| **Properties:**  | Name | Value | | --- | --- | | addresses | type: array  items: type: string  title: List of returning mac or ip address | | totalCount | type: integer  format: int32  title: Total number of addresses available | | |

#### CreateFlowTagRequest

|  |  |
| --- | --- |
| **Schema:** v202404alpha1CreateFlowTagRequest | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | flowTag | $ref: v202404alpha1FlowTag | | |

#### CreateFlowTagResponse

|  |  |
| --- | --- |
| **Schema:** v202404alpha1CreateFlowTagResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | flowTag | $ref: v202404alpha1FlowTag | | |

#### DeleteFlowTagResponse

|  |  |
| --- | --- |
| **Schema:** v202404alpha1DeleteFlowTagResponse | **Type:** object |
| **Properties**:  None | |

#### FlowTag

|  |  |
| --- | --- |
| **Schema:** v202404alpha1FlowTag | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | id | type: string  description: Unique system assigned identifier of the flow tag  readOnly: true | | name | type: string  description: This will appear in places where selecting a tag is necessary  readOnly: true | | editedBy | type: string  description: User who last edited this tag  readOnly: true | | createdBy | type: string  description: User who created this tag  readOnly: true | | cdate | type: string  format: date-time  description: Date and time when this tag was created  readOnly: true | | edate | type: string  format: date-time  description: Date and time when this tag was last updated  readOnly: true | | ip | $ref: v202404alpha1AddressInfo | | port | type: array  items: type: string  description: Port number associated with the flow tag | | tcpFlags | type: integer  format: int64  description: TCP flags associated with the flow tag | | protocol | type: array  items: type: integer  format: int64  description: Protocol numbers associated with the flow tag | | deviceName | type: array  items: type: string  description: Name of the device associated with the flow tag | | deviceType | type: array  items: type: string  description: Type of the device associated with the flow tag | | site | type: array  items: type: string  description: Site where the device associated with the flow tag is located | | interfaceName | type: array  items: type: string  description: Name of the interface associated with the flow tag | | asn | type: array  items: type: string  description: Autonomous System Number (ASN) associated with the flow tag | | lasthopAsName | type: array  items: type: string  description: Name of the last hop's Autonomous System (AS) associated with the flow tag | | nexthopAsn | type: array  items: type: string  description: Autonomous System Number (ASN) of the next hop associated with the flow tag | | nexthopAsName | type: array  items: type: string  description: Name of the next hop's Autonomous System (AS) associated with the flow tag | | nexthop | type: array  items: type: string  description: Next hop associated with the flow tag | | bgpAspath | type: array  items: type: string  description: BGP AS path associated with the flow tag | | bgpCommunity | type: array  items: type: string  description: BGP community associated with the flow tag | | mac | $ref: v202404alpha1AddressInfo | | country | type: array  items: type: string  description: Country associated with the flow tag | | vlans | type: array  items: type: string  description: VLANs associated with the flow tag | | |

#### FlowTagSearch

|  |  |
| --- | --- |
| **Schema:** v202404alpha1FlowTagSearch | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | limit | type: integer  format: int32  title: Limit the number of rows to fetch | | offset | type: integer  format: int32  title: The number rows to skip before returning | | lookupFields | type: array  items: $ref: v202404alpha1LookupField  title: List of lookup fields | | lookupValues | type: array  items: type: string  title: List of lookup values | | orderBy | type: array  items: $ref: v202404alpha1OrderField  title: Sort order | | fieldLimit | type: integer  format: int32  title: Limit the number of record to return for nested fields | | |

#### GetFlowTagResponse

|  |  |
| --- | --- |
| **Schema:** v202404alpha1GetFlowTagResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | flowTag | $ref: v202404alpha1FlowTag | | |

#### LookupField

|  |  |
| --- | --- |
| **Schema:** v202404alpha1LookupField | **Type:** string |
| **Attributes**:  | Key | Value | | --- | --- | | enum | LOOKUP\_FIELD\_UNSPECIFIED, LOOKUP\_FIELD\_VALUE | | default | LOOKUP\_FIELD\_UNSPECIFIED | | description | • LOOKUP\_FIELD\_UNSPECIFIED: Invalid value  • LOOKUP\_FIELD\_VALUE: Lookup by name of tag | | |

#### OrderDirection

|  |  |
| --- | --- |
| **Schema:** v202404alpha1OrderDirection | **Type:** string |
| **Attributes**:  | Key | Value | | --- | --- | | enum | ORDER\_DIRECTION\_UNSPECIFIED, ORDER\_DIRECTION\_ASC, ORDER\_DIRECTION\_DESC | | default | ORDER\_DIRECTION\_UNSPECIFIED | | description | • ORDER\_DIRECTION\_UNSPECIFIED: Invalid value  • ORDER\_DIRECTION\_ASC: Ascending sort order  • ORDER\_DIRECTION\_DESC: Descending sort order | | |

#### OrderField

|  |  |
| --- | --- |
| **Schema:** v202404alpha1OrderField | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | field | $ref: v202404alpha1LookupField | | direction | $ref: v202404alpha1OrderDirection | | |

#### SearchFlowTagResponse

|  |  |
| --- | --- |
| **Schema:** v202404alpha1SearchFlowTagResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | flowTags | type: array  items: $ref: v202404alpha1FlowTag | | totalCount | type: integer  format: int64 | | invalidCount | type: integer  format: int64 | | |

#### UpdateFlowTagRequest

|  |  |
| --- | --- |
| **Schema:** v202404alpha1UpdateFlowTagRequest | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | flowTag | $ref: v202404alpha1FlowTag | | |

#### UpdateFlowTagResponse

|  |  |
| --- | --- |
| **Schema:** v202404alpha1UpdateFlowTagResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | flowTag | $ref: v202404alpha1FlowTag | | |
