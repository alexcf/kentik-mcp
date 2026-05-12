# Interface APIs

Source: https://kb.kentik.com/docs/interface-apis

---

> **Note**: This alpha API is neither supported nor recommended for production use. For additional information, contact Kentik (**see Customer Care**).

This article covers how to get started with the Interface APIs.

> **Notes:**
>
> * The **API Tester** in the v4 Portal enables you to securely call the methods of these APIs using your organization's Kentik data.
> * For information about interfaces, start with **About Interfaces**.
> * Protobuf and OpenAPI specifications for Kentik’s v6 APIs are available in our **api-schema-public** repository.

## Interface Usage

Manage interfaces programmatically, including listing, creating, updating, and deleting.

## Interface RPCs

The tables below provide a quick reference to key information about each Remote Procedure Call in these APIs. Use the links in the tables to see schemas for the request and/or response body (if any).

> **Note:** To test methods using your own Kentik data, use the portal's **API Tester**.

### ListInterface

**API: InterfaceService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| GET | /interface/v202108alpha1 /interfaces | Return list of interfaces matches search critera. |
| |  |  | | --- | --- | | **Request body**:  None | **Response body**:  v202108alpha1ListInterfaceResponse |    **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | filters.text | Search text match in interface description and alias. | false | string | | filters.deviceIds | Search by device ID. | false | array | | filters.connectivityTypes | Search by type of interface connectivity. | false | array | | filters.networkBoundaries | Search by type of network boundary. | false | array | | filters.providers | Search by provider. | false | array | | filters.snmpSpeeds | SNMP speed in Mbps. | false | array | | filters.ipTypes | Search by ip address. | false | array | | | |

### InterfaceCreate

**API: InterfaceService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| POST | /interface/v202108alpha1 /interfaces | Create a interface from request. returns created |
| |  |  | | --- | --- | | **Request body**:  v202108alpha1CreateInterfaceRequest | **Response body:**  v202108alpha1CreateInterfaceResponse | | **Parameters:**  None | | | | |

### InterfaceGet

**API: InterfaceService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| GET | /interface/v202108alpha1 /interfaces/{id} | Returns information about a interface specified with ID. |
| |  |  | | --- | --- | | **Request body**:  None | **Response body**:  v202108alpha1GetInterfaceResponse |  **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | id | ID of interface. | true | string | | | |

### InterfaceDelete

**API: InterfaceService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| DELETE | /interface/v202108alpha1 /interfaces/{id} | Deletes the interface specified with id. |
| |  |  | | --- | --- | | **Request body**:  None | **Response body**:  v202108alpha1DeleteInterfaceResponse |  **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | id | ID of interface. | true | string | | | |

### InterfaceUpdate

**API: InterfaceService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| PUT | /interface/v202108alpha1 /interfaces/{id} | Replaces the entire interface attributes specified with id. |
| |  |  | | --- | --- | | **Request body**:  v202108alpha1UpdateInterfaceRequest | **Response body**:  v202108alpha1UpdateInterfaceResponse |    **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | id | ID of interface. | true | string | | | |

### ManualClassify

**API: InterfaceService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| POST | /interface/v202108alpha1 /manual\_classify | Manually set interface(s) classification. |
| |  |  | | --- | --- | | **Request body**:  v202108alpha1ManualClassifyRequest | **Response body:**  v202108alpha1ManualClassifyResponse | | **Parameters:**  None |  | | | |

## Interface Schemas

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

#### ConnectivityType

|  |  |
| --- | --- |
| **Schema:** v202108alpha1ConnectivityType | **Type:** string |
| **Attributes**:  | Key | Value | | --- | --- | | enum | CONNECTIVITY\_TYPE\_UNSPECIFIED, CONNECTIVITY\_TYPE\_AGGREGATION\_INTERCONNECT, CONNECTIVITY\_TYPE\_AVAILABLE, CONNECTIVITY\_TYPE\_BACKBONE, CONNECTIVITY\_TYPE\_CLOUD\_INTERCONNECT, CONNECTIVITY\_TYPE\_CUSTOMER, CONNECTIVITY\_TYPE\_DATACENTER\_FABRIC, CONNECTIVITY\_TYPE\_DATACENTER\_INTERCONNECT, CONNECTIVITY\_TYPE\_EMBEDDED\_CACHE, CONNECTIVITY\_TYPE\_FREE\_PNI, CONNECTIVITY\_TYPE\_HOST, CONNECTIVITY\_TYPE\_IX, CONNECTIVITY\_TYPE\_OTHER, CONNECTIVITY\_TYPE\_PAID\_PNI, CONNECTIVITY\_TYPE\_RESERVED, CONNECTIVITY\_TYPE\_TRANSIT, CONNECTIVITY\_TYPE\_VIRTUAL\_CROSS\_CONNECT | | default | CONNECTIVITY\_TYPE\_UNSPECIFIED | | |

#### CreateInterfaceRequest

|  |  |
| --- | --- |
| **Schema:** v202108alpha1CreateInterfaceRequest | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | interface | $ref: v202108alpha1Interface | | |

#### CreateInterfaceResponse

|  |  |
| --- | --- |
| **Schema:** v202108alpha1CreateInterfaceResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | interface | $ref: v202108alpha1Interface | | |

#### DeleteInterfaceResponse

|  |  |
| --- | --- |
| **Schema:** v202108alpha1DeleteInterfaceResponse | **Type:** object |
| **Properties**:  None | |

#### GetInterfaceResponse

|  |  |
| --- | --- |
| **Schema:** v202108alpha1GetInterfaceResponse | **Type:** object |
| **Properties:**  | Name | Value | | --- | --- | | interface | $ref: v202108alpha1Interface | | |

#### Interface

|  |  |
| --- | --- |
| **Schema:** v202108alpha1Interface | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | id | type: string  description: ID of interface. | | deviceId | type: string  description: ID of device with this interface. | | snmpId | type: string  description: SNMP ID. | | snmpSpeed | type: integer  format: int32  description: Network speed in Mbps. | | snmpType | type: integer  format: int32  description: Snmp Type. | | snmpAlias | type: string  description: Interface alias. | | interfaceIp | type: string  description: IP of interface. | | interfaceDescription | type: string  description: Readable string description of interface. | | cdate | type: string  format: date-time  description: Create timestamp. | | edate | type: string  format: date-time  description: Update Timestamp. | | interfaceIpNetmask | type: string  description: Subnet mask. | | connectivityType | $ref: v202108alpha1ConnectivityType | | networkBoundary | $ref: v202108alpha1NetworkBoundary | | topNexthopAsns | type: array  items: type: integer  format: int32  description: Top asn hops. | | provider | type: string  description: Network provider. | | |

#### InterfaceFilter

|  |  |
| --- | --- |
| **Schema:** v202108alpha1InterfaceFilter | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | text | type: string  description: Search text match in interface description and alias. | | deviceIds | type: array  items: type: string  description: Search by device ID. | | connectivityTypes | type: array  items: $ref: v202108alpha1ConnectivityType  description: Search by type of interface connectivity. | | networkBoundaries | type: array  items: $ref: v202108alpha1NetworkBoundary  description: Search by type of network boundary. | | providers | type: array  items: type: string  title: Search by provider | | snmpSpeeds | type: array  items: type: integer  format: int32  title: SNMP speed in Mbps | | ipTypes | type: array  items: $ref: v202108alpha1IpFilter  description: Search by ip address. | | |

#### IpFilter

|  |  |
| --- | --- |
| **Schema:** v202108alpha1IpFilter | **Type:** string |
| **Attributes**:  | Name | Value | | --- | --- | | enum | IP\_FILTER\_UNSPECIFIED, IP\_FILTER\_PRIVATE, IP\_FILTER\_PUBLIC, IP\_FILTER\_UNSET | | default | IP\_FILTER\_UNSPECIFIED | | |

#### ListInterfaceResponse

|  |  |
| --- | --- |
| **Schema:** v202108alpha1ListInterfaceResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | interfaces | type: array  items: $ref: v202108alpha1Interface | | totalCount | type: integer  format: int32 | | invalidCount | type: integer  format: int32 | | |

#### ManualClassifyRequest

|  |  |
| --- | --- |
| **Schema:** v202108alpha1ManualClassifyRequest | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | interfaceIds | type: array  items: type: string  description: Interface ids to set the properties. | | connectivityType | $ref: v202108alpha1ConnectivityType | | networkBoundary | $ref: v202108alpha1NetworkBoundary | | provider | type: string  description: Network provider. | | |

#### ManualClassifyResponse

|  |  |
| --- | --- |
| **Schema:** v202108alpha1ManualClassifyResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | deviceIds | type: array  items: type: string | | |

#### NetworkBoundary

|  |  |
| --- | --- |
| **Schema:** v202108alpha1NetworkBoundary | **Type:** string |
| **Attributes:**  | Key | Value | | --- | --- | | enum | NETWORK\_BOUNDARY\_UNSPECIFIED, NETWORK\_BOUNDARY\_AUTO, NETWORK\_BOUNDARY\_INTERNAL, NETWORK\_BOUNDARY\_EXTERNAL, NETWORK\_BOUNDARY\_NONE | | default | NETWORK\_BOUNDARY\_UNSPECIFIED | | |

#### UpdateInterfaceRequest

|  |  |
| --- | --- |
| **Schema:** v202108alpha1UpdateInterfaceRequest | **Type:** object |
| **Properties:**  | Name | Value | | --- | --- | | interface | $ref: v202108alpha1Interface | | |

#### UpdateInterfaceResponse

|  |  |
| --- | --- |
| **Schema:** v202108alpha1UpdateInterfaceResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | interface | $ref: v202108alpha1Interface | | |
