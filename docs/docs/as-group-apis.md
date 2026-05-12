# AS Group APIs

Source: https://kb.kentik.com/docs/as-group-apis

---

This article covers how to get started with the AS Group APIs.

> **Notes:**
>
> * The **API Tester**in the v4 Portal enables you to securely call the methods of these APIs using your organization's Kentik data.
> * AS Groups are supported in the Kentik portal via the **Settings »** **AS Groups** page. For more information, see **AS Groups**.
> * Protobuf specifications for Kentik’s v6 APIs are available in our **api-schema-public** repository.

## AS Group Usage

The AS Group API provides programmatic access to the configuration of AS Groups, allowing the traffic of a set of ASes to be summed so it can be evaluated together for the purpose of filtering and top-X grouping in queries.

## AS Group RPCs

The tables below provide a quick reference to key information about each Remote Procedure Call in these APIs. Use the links in the tables to see schemas for the request and/or response body (if any).

> **Note:** To test methods using your own Kentik data, use the portal's **API Tester**.

### ListASGroups

**API: ASGroupService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| GET | /as\_group/v202212 /as\_group | Returns list of configured AS groups. |
| |  |  | | --- | --- | | **Request body:**  None  **Parameters:** None | **Response body**:  v202212ListASGroupsResponse | | | |

### CreateASGroup

**API: ASGroupService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| POST | /as\_group/v202212 /as\_group | Create configuration for a new AS group. Returns the newly created configuration. |
| |  |  | | --- | --- | | **Request body:**  v202212CreateASGroupRequest  **Parameters:**  None | **Response body**:  v202212CreateASGroupResponse | | | |

### UpdateASGroup

**API: ASGroupService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| PUT | /as\_group/v202212 /as\_group/{asGroup.id} | Replaces configuration of an AS group with attributes in the request. Returns the updated configuration. |
| |  |  | | --- | --- | | **Request body:** v202212UpdateASGroupRequest | **Response body:**  v202212UpdateASGroupResponse |    **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | asGroup.id | System generated unique identifier | true | string | | | |

### GetASGroup

**API: ASGroupService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| GET | /as\_group/v202212 /as\_group/{asGroup.id} | Returns configuration of a AS group specified by ID. |
| |  |  | | --- | --- | | **Request body**:  None | **Response body:**  v202212GetASGroupResponse |  **Parameters:**  | Name | Description | Required | Type | | --- | --- | --- | --- | | asGroup.id | ID of the requested AS group | true | string | | | |

### DeleteASGroup

**API: ASGroupService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| DELETE | /as\_group/v202212 /as\_group/{asGroup.id} | Deletes configuration of an AS group with specific ID. |
| |  |  | | --- | --- | | **Request body:**  None | **Response body:**  v202212DeleteASGroupResponse |  **Parameters:**  | Name | Description | Required | Type | | --- | --- | --- | --- | | asGroup.id | ID of the AS group to be deleted | true | string | | | |

## AS Group Schemas

This API uses the following schemas.

#### protobufAny

|  |  |
| --- | --- |
| **Schema:** protobufAny | **Type:** object |
| **Properties:**  | Name | Value | | --- | --- | | typeUrl | type: string | | value | type: string  format: byte | | |

#### rpcStatus

|  |  |
| --- | --- |
| **Schema:** rpcStatus | **Type:** object |
| **Properties:**  | Name | Value | | --- | --- | | code | type: integer  format: int32 | | message | type: string | | details | type: array  items: $ref: protobufAny | | |

#### ASGroupConcise

|  |  |
| --- | --- |
| **Schema:** v202212ASGroupConcise | **Type:** object |
| **Properties** **(\*  =  required)**  | Name | Value | | --- | --- | | id | type: string  description: System generated unique identifier  readOnly: true | | name \* | type: string  description: User selected unique name | | asn | type: array  items: type: string  description: List of ASNs | | createdDate | type: string  format: date-time  description: Creation timestamp (UTC)  readOnly: true | | updatedDate | type: string  format: date-time  description: Last modification timestamp (UTC)  readOnly: true | | |

#### ASGroupDetailed

|  |  |
| --- | --- |
| **Schema:** v202212ASGroupDetailed | **Type:** object |
| **Properties (\*  =  required)**  | Name | Value | | --- | --- | | id | type: string  description: System generated unique identifier  readOnly: true | | name \* | type: string  description: User selected unique name | | asn | type: array  items: $ref: v202212AutonomousSystem  description: List of ASNs | | createdDate | type: string  format: date-time  description: Creation timestamp (UTC)  readOnly: true | | updatedDate | type: string  format: date-time  description: Last modification timestamp (UTC)  readOnly: true | | |

#### AutonomousSystem

|  |  |
| --- | --- |
| **Schema:** v202212AutonomousSystem | **Type:** object |
| **Properties:**  | Name | Value | | --- | --- | | asn | type: integer  format: int64 | | name | type: string | | |

#### CreateASGroupRequest

|  |  |
| --- | --- |
| **Schema:** v202212CreateASGroupRequest | **Type:** object |
| **Properties:**  | Name | Value | | --- | --- | | asGroup | $ref: v202212ASGroupConcise | | |

#### CreateASGroupResponse

|  |  |
| --- | --- |
| **Schema:** v202212CreateASGroupResponse | **Type:** object |
| **Properties:**  | Name | Value | | --- | --- | | asGroup | $ref: v202212ASGroupDetailed | | |

#### DeleteASGroupResponse

|  |  |
| --- | --- |
| **Schema:** v202212DeleteASGroupResponse | **Type:** object |
| **Properties:** None. | |

#### GetASGroupResponse

|  |  |
| --- | --- |
| **Schema:** v202212GetASGroupResponse | **Type:** object |
| **Properties:**  | Name | Value | | --- | --- | | asGroup | $ref: v202212ASGroupDetailed | | |

#### ListASGroupsResponse

|  |  |
| --- | --- |
| **Schema:** v202212ListASGroupsResponse | **Type:** object |
| **Properties:**  | Name | Value | | --- | --- | | asGroups | type: array  items: $ref: v202212ASGroupDetailed  description: List of configurations of requested AS groups | | invalidCount | type: integer  format: int64  description: Number of invalid entries encountered while collecting data | | |

#### UpdateASGroupRequest

|  |  |
| --- | --- |
| **Schema:** v202212UpdateASGroupRequest | **Type:** object |
| **Properties:**  | Name | Value | | --- | --- | | asGroup | $ref:v202212ASGroupConcise | | |

#### UpdateASGroupResponse

|  |  |
| --- | --- |
| **Schema:** v202212UpdateASGroupResponse | **Type:** object |
| **Properties:**  | Name | Value | | --- | --- | | asGroup | $ref: v202212ASGroupDetailed | | |
