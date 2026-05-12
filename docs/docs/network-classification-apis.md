# Network Classification APIs

Source: https://kb.kentik.com/docs/network-classification-apis

---

> **Note**: This alpha API is neither supported nor recommended for production use. For additional information, contact Kentik (see **Customer Care**).

This article covers how to get started with the Network Classification APIs.

> **Notes:**
>
> * The **API Tester** in the v4 Portal enables you to securely call the methods of these APIs using your organization's Kentik data.
> * For information about Network Classification, start with **About Network Classification**.
> * Protobuf and OpenAPI specifications for Kentik’s v6 APIs are available in our **api-schema-public** repository.

## Classification RPCs

The tables below provide a quick reference to key information about each Remote Procedure Call in these APIs. Use the links in the tables to see schemas for the request and/or response body (if any).

> **Note:** To test methods using your own Kentik data, use the portal's **API Tester**.

### NetworkClassGet

**API: NetworkClassService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| GET | /network\_class/v202109alpha1 /network\_class | Returns information about a network classification for the company. |
| |  |  | | --- | --- | | **Request body:** None **Parameters:** None | **Response body:**  v202109alpha1GetNetworkClassResponse | | | |

### NetworkClassUpdate

**API: NetworkClassService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| POST | /network\_class/v202109alpha1 /network\_class | Replaces the entire network classification attributes for the company. |
| |  |  | | --- | --- | | **Request body:** v202109alpha1UpdateNetworkClassRequest **Parameters:** None | **Response body:**  v202109alpha1UpdateNetworkClassResponse | | | |

## Classification Schemas

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

#### CloudSubnet

|  |  |
| --- | --- |
| **Schema:** v202109alpha1CloudSubnet | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | type | $ref: v202109alpha1CloudType | | subnets | type: array  items: type: string  title: Subnet masks | | |

#### CloudType

|  |  |
| --- | --- |
| **Schema:** v202109alpha1CloudType | **Type:** string |
| **Attributes**:  | Key | Value | | --- | --- | | enum | CLOUD\_TYPE\_UNSPECIFIED, CLOUD\_TYPE\_AWS, CLOUD\_TYPE\_AZURE, CLOUD\_TYPE\_GCE, CLOUD\_TYPE\_IBM, CLOUD\_TYPE\_OCI | | default | CLOUD\_TYPE\_UNSPECIFIED | | description | • CLOUD\_PROVIDER\_UNSPECIFIED: Invalid value.  • CLOUD\_PROVIDER\_AWS: Amazon Web Services  • CLOUD\_PROVIDER\_AZURE: Microsoft Azure  • CLOUD\_PROVIDER\_GCE: Google Cloud  • CLOUD\_PROVIDER\_IBM: IBM Cloud Deprecated: IBM Cloud exports are no longer supported.  • CLOUD\_PROVIDER\_OCI: OCI | | |

#### GetNetworkClassResponse

|  |  |
| --- | --- |
| **Schema:** v202109alpha1GetNetworkClassResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | networkClass | $ref: v202109alpha1NetworkClass | | |

#### NetworkClass

|  |  |
| --- | --- |
| **Schema:** v202109alpha1NetworkClass | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | internalAsns | type: array  items: type: string  title: Internal ASNs | | internalIps | type: array  items: type: string  title: Internal IPs | | usePrivateAsns | type: boolean  title: Specify to use internal ASNs | | usePrivateSubnets | type: boolean  title: Specify to use internal IPs | | cloudSubnets | type: array  items: $ref: v202109alpha1CloudSubnet  description: Cloud provider and subnet info. | | |

#### UpdateNetworkClassRequest

|  |  |
| --- | --- |
| **Schema:** v202109alpha1UpdateNetworkClassRequest | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | networkClass | $ref: v202109alpha1NetworkClass | | |

#### UpdateNetworkClassResponse

|  |  |
| --- | --- |
| **Schema:** v202109alpha1UpdateNetworkClassResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | networkClass | $ref: v202109alpha1NetworkClass | | |
