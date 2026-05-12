# User APIs

Source: https://kb.kentik.com/docs/user-apis

---

This article covers how to get started with the User APIs.

> **Notes:**
>
> * The **API Tester** in the v4 Portal enables you to securely call the methods of these APIs using your organization's Kentik data.
> * For information about users in Kentik, start with **About Users**.
> * Protobuf and OpenAPI specifications for Kentik’s v6 APIs are available in our **api-schema-public** repository.

## User Usage

The topics below provide important background information for the use of these APIs.

### Overview

The User Management API enables programmatic administration of users within your organization. The API may be used to grant permissions to users at two levels:

* **User role**: Permissions based on the level assigned to the Kentik user, either Member, Administrator, or Super Administrator (see **About Users**).
* **User permissions**: Permissions granting access to specific capabilities of the Kentik platform (see **User Permissions**).

Both REST endpoint and gRPC RPCs are provided.

> **Note:** The legacy REST-only **User API** provides a subset of the functionality of this API.

### User Permission Identifiers

Individual user permissions are represented in this API by string-based identifiers. The table below shows the identifier used in the API for each permission that can be enabled/disabled with the checkboxes on the Permissions tab of the Edit User dialog in the Kentik portal's Users module (Settings » **Users**).

| Portal setting | API identifier |
| --- | --- |
| Cancel Queries in Forensic Viewer | forensic\_viewer.cancel\_queries |
| Synthetics Permissions: Create tests | synthetics.tests.create |
| Synthetics Permissions: Edit tests | synthetics.tests.edit |
| Synthetics Permissions: Delete tests | synthetics.tests.delete |
| Synthetics Permissions: Create agents | synthetics.agents.create |
| Synthetics Permissions: Edit agents | synthetics.agents.edit |
| Synthetics Permissions: Delete agents | synthetics.agents.delete |
| Connectivity Costs Permissions: View Costs | connectivity\_costs.read |
| Connectivity Costs Permissions: Configure Costs | connectivity\_costs.edit |

> **Note:** The API returns an error if the same permission is both granted and denied in the same request.

#### Default Permissions

The permissions above that are enabled/disabled by default depend on the role (level) of the user:

* **Member**: All permissions in the table above are denied by default.
* **Administrator** or **Super Administrator**: All permissions are granted.

### User Filter Attribute

User filters enable Administrators to apply filters that restrict the data that can be returned from queries by a given user (see **User Filters**). In the Kentik portal these filters are set on the Filters tab of the Edit User dialog in the Users module. In this API the filters are set in the `user` object with the `filter` attribute, whose type is string. The value of the attribute is given in JSON notation. The recommended way to construct the JSON string is to configure the desired filter for one user in the Kentik portal and then copy the value of the `user.filter` attribute returned for that user by the `GetUser` RPC.

## User RPCs

The tables below provide a quick reference to key information about each Remote Procedure Call in these APIs. Use the links in the tables to see schemas for the request and/or response body (if any).

> **Note:** To test methods using your own Kentik data, use the portal's **API Tester**.

### ListUsers

**API: UserService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| GET | /user/v202211 /users | Returns a list of all user accounts in the company. |
| |  |  | | --- | --- | | **Request body:** None **Parameters**: None | **Response body:**  v202211ListUsersResponse | | | |

### CreateUser

**API: UserService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| POST | /user/v202211 /users | Creates new user account based on attributes in the request. Returns attributes of the newly created account. |
| |  |  | | --- | --- | | **Request body:** v202211CreateUserRequest  **Parameters**: None | **Response body:**  v202211CreateUserResponse | | | |

### GetUser

**API: UserService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| GET | /user/v202211 /users/{id} | Returns attributes of a user account specified by ID. |
| |  |  | | --- | --- | | **Request body:** None | **Response body:**  v202211GetUserResponse |  **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | id | ID of the requested user | true | string | | | |

### DeleteUser

**API: UserService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| DELETE | /user/v202211 /users/{id} | Deletes user account specified by ID. |
| |  |  | | --- | --- | | **Request body**: None | **Response body**:  v202211DeleteUserResponse |    **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | id | ID of the the user account to be deleted | true | string | | | |

### UpdateUser

**API: UserService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| PUT | /user/v202211 /users/{id} | Replaces all attributes of a user account specified by ID with attributes in the request. Returns updated attributes. |
| |  |  | | --- | --- | | **Request body**: v202211UpdateUserRequest | **Response body**: v202211UpdateUserResponse |  **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | id | System generated unique identifier | true | string | | | |

### ResetActiveSessions

**API: UserService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| PUT | /user/v202211 /users/{id}/reset\_active\_sessions | Resets active sessions for a user specified by ID. |
| |  |  | | --- | --- | | **Request body**: None | **Response body**:  v202211ResetActiveSessionsResponse |  **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | id | ID of the the user whose sessions should be reset | true | string | | | |

### ResetApiToken

**API: UserService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| PUT | /user/v202211 /users/{id}/reset\_api\_token | Resets API token for a user specified by ID. |
| |  |  | | --- | --- | | **Request body**: None | **Response body**:  v202211ResetApiTokenResponse |  **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | id | ID of the the user whose API token should be reset | true | string | | | |

## User Schemas

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

#### CreateUserRequest

|  |  |
| --- | --- |
| **Schema:** v202211CreateUserRequest | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | user | $ref: v202211User | | |

#### CreateUserResponse

|  |  |
| --- | --- |
| **Schema:** v202211CreateUserResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | user | $ref: v202211User | | |

#### DeleteUserResponse

|  |  |
| --- | --- |
| **Schema:** v202211DeleteUserResponse | **Type:** object |
| **Properties**: None. | |

#### GetUserResponse

|  |  |
| --- | --- |
| **Schema:** v202211GetUserResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | user | $ref: v202211User | | |

#### ListUsersResponse

|  |  |
| --- | --- |
| **Schema:** v202211ListUsersResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | users | type: array  items: $ref: v202211User  description: Last of users in the account  readOnly: true | | invalidCount | type: integer  format: int64  description: Number of invalid entries encountered while collecting data | | |

#### PermissionEntry

|  |  |
| --- | --- |
| **Schema:** v202211PermissionEntry | **Type:** object |
| **Properties (\*  =  required)**  | Name | Value | | --- | --- | | capability \* | type: string  description: String identifying capability that is granted of denied | | allowed \* | type: boolean  description: Flag indicating whether operation is allowed | | |

#### ResetActiveSessionsResponse

|  |  |
| --- | --- |
| **Schema:** v202211ResetActiveSessionsResponse | **Type:** object |
| **Properties**: None. | |

#### ResetApiTokenResponse

|  |  |
| --- | --- |
| **Schema:** v202211ResetApiTokenResponse | **Type:** object |
| **Properties**: None. | |

#### Role

|  |  |
| --- | --- |
| **Schema:** v202211Role | **Type:** string |
| **Attributes**:  | Key | Value | | --- | --- | | enum | ROLE\_UNSPECIFIED, ROLE\_MEMBER, ROLE\_ADMINISTRATOR, ROLE\_SUPER\_ADMINISTRATOR | | default | ROLE\_UNSPECIFIED | | description | • ROLE\_UNSPECIFIED: Invalid value.  • ROLE\_MEMBER: Member  • ROLE\_ADMINISTRATOR: Administrator  • ROLE\_SUPER\_ADMINISTRATOR: Super-administrator | | |

#### UpdateUserRequest

|  |  |
| --- | --- |
| **Schema:** v202211UpdateUserRequest | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | user | $ref: v202211User | | |

#### UpdateUserResponse

|  |  |
| --- | --- |
| **Schema:** v202211UpdateUserResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | user | $ref: v202211User | | |

#### User

|  |  |
| --- | --- |
| **Schema:** v202211User | **Type:** object |
| **Properties (\*  =  required)**  | Name | Value | | --- | --- | | id | type: string  description: System generated unique identifier  readOnly: true | | userEmail \* | type: string  description: User e-mail address (serves also as username) | | userFullName \* | type: string  description: Full name | | role | $ref: v202211Role | | permissions | type: array  items: $ref: v202211PermissionEntry  description: Optional list of permissions granted to the user | | filter | type: string  description: Optional JSON string defining filter for objects visible to the user | | lastLogin | type: string  format: date-time  description: UTC Timestamp of user's last login session  readOnly: true | | cdate | type: string  format: date-time  description: Creation timestamp (UTC)  readOnly: true | | edate | type: string  format: date-time  description: Last modification timestamp (UTC)  readOnly: true | | |
