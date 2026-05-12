# Credentials Vault APIs

Source: https://kb.kentik.com/docs/credentials-vault-apis

---

> **Note**: This alpha API is neither supported nor recommended for production use. For additional information, please contact **Product Support**.

This article covers how to get started with the Credentials Vault APIs.

> **Notes:**
>
> * The **API Tester** in the v4 Portal enables you to securely call the methods of these APIs using your organization's Kentik data.
> * For an explanation of the capabilities and usage of the Credentials Vault in the Kentik portal, see **Credentials Vault**.
> * Protobuf and OpenAPI specifications for Kentik’s v6 APIs are available in our **api-schema-public** repository.

## Credentials Vault Usage

The Credentials Vault API enables programmatic access to credential information in Kentik Credential Vault.

## Credentials Vault RPCs

The tables below provide a quick reference to key information about each Remote Procedure Call in these APIs. Use the links in the tables to see schemas for the request and/or response body (if any).

> **Note:** To test methods using your own Kentik data, use the portal's **API Tester**.

### ListCredentialGroup

**API: CredentialService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| GET | /credential/v202407alpha1 /group | Returns list of credential group information in Kentik vault. |
| |  |  | | --- | --- | | **Request body**: None  **Parameters**: None | **Response body**:  v202407alpha1ListCredentialGroupResponse | | | |

### GetCredentialGroup

**API: CredentialService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| GET | /credential/v202407alpha1 /group/{id} | Returns specific credential group information in Kentik vault. |
| |  |  | | --- | --- | | **Request body**: None | **Response body**:  v202407alpha1GetCredentialGroupResponse |  **Parameters:**  | Name | Description | Required | Type | | --- | --- | --- | --- | | id | Undefined. | true | string | | | |

## Credentials Vault Schemas

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

#### PermissionEntry

|  |  |
| --- | --- |
| **Schema:** v202211PermissionEntry | **Type:** object |
| **Properties** **(\*  =  required)**  | Name | Value | | --- | --- | | capability \* | type: string  description: String identifying capability that is granted of denied | | allowed \* | type: boolean  description: Flag indicating whether operation is allowed | | |

#### Role

|  |  |
| --- | --- |
| **Schema:** v202211Role | **Type:** string |
| **Attributes:**  | Key | Value | | --- | --- | | enum | ROLE\_UNSPECIFIED, ROLE\_MEMBER, ROLE\_ADMINISTRATOR, ROLE\_SUPER\_ADMINISTRATOR | | default | ROLE\_UNSPECIFIED | | description | • ROLE\_UNSPECIFIED: Invalid value.  • ROLE\_MEMBER: Member  • ROLE\_ADMINISTRATOR: Administrator  • ROLE\_SUPER\_ADMINISTRATOR: Super-administrator | | |

#### User

|  |  |
| --- | --- |
| **Schema:** v202211User | **Type:** object |
| **Properties (\*  =  required)**  | Name | Value | | --- | --- | | id | type: string  description: System generated unique identifier  readOnly: true | | userEmail \* | type: string  description: User e-mail address (serves also as username) | | userFullName \* | type: string  description: Full name | | role | $ref: v202211Role | | permissions | type: array  items: $ref: v202211PermissionEntry  description: Optional list of permissions granted to the user | | filter | type: string  description: Optional JSON string defining filter for objects visible to the user | | lastLogin | type: string  format: date-time  description: UTC Timestamp of user's last login session  readOnly: true | | cdate | type: string  format: date-time  description: Creation timestamp (UTC)  readOnly: true | | edate | type: string  format: date-time  description: Last modification timestamp (UTC)  readOnly: true | | |

#### Secret

|  |  |
| --- | --- |
| **Schema:** v202312alpha1Secret | **Type:** object |
| **Properties** **(\*  =  required)**  | Name | Value | | --- | --- | | name \* | type: string  description: The secret's name | | value \* | type: string  description: The secret's actual value | | version \* | type: integer  format: int64  description: The secret's version | | description | type: string  description: Purpose of this secret | | type | $ref: v202312alpha1SecretType | | id \* | type: string  description: The secret's identifier | | |

#### SecretType

|  |  |
| --- | --- |
| **Schema:** v202312alpha1SecretType | **Type:** string |
| **Attributes:**  | Key | Value | | --- | --- | | enum | SECRET\_TYPE\_UNSPECIFIED, SECRET\_TYPE\_BASIC\_AUTH, SECRET\_TYPE\_SNMP\_V1, SECRET\_TYPE\_SNMP\_V2C, SECRET\_TYPE\_SNMP\_V3, SECRET\_TYPE\_STREAMING\_TELEMETRY, SECRET\_TYPE\_BGP\_MD5, SECRET\_TYPE\_API\_TOKEN, SECRET\_TYPE\_OTHER | | default | SECRET\_TYPE\_UNSPECIFIED | | |

#### CredentialGroup

|  |  |
| --- | --- |
| **Schema:** v202407alpha1CredentialGroup | **Type:** object |
| **Properties (\*  =  required)**  | Name | Value | | --- | --- | | id | type: string | | name | type: string | | version \* | type: integer  format: int64  description: The secret's version | | description | type: string  description: Purpose of this secret | | type | $ref: v202312alpha1SecretType | | cdate | type: string  format: date-time | | edate | type: string  format: date-time | | createdBy | $ref: v202211User | | credentials | type: array  items: $ref: v202312alpha1Secret | | labels | type: array  items: type: string | | |

#### GetCredentialGroupResponse

|  |  |
| --- | --- |
| **Schema:** v202407alpha1GetCredentialGroupResponse | **Type:** object |
| **Properties:**  | Name | Value | | --- | --- | | group | $ref: v202407alpha1CredentialGroup | | |

#### ListCredentialGroupResponse

|  |  |
| --- | --- |
| **Schema:** v202407alpha1ListCredentialGroupResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | groups | type: array  items: $ref: v202407alpha1CredentialGroup | | invalidCount | type: integer  format: int64 | | |
