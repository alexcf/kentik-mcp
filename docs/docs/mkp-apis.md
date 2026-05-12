# MKP APIs

Source: https://kb.kentik.com/docs/mkp-apis

---

> **Note**: This version replaces v202102alpha1. If you are looking for the older version, please refer to the portal's **API Tester**.

This article covers how to get started with the MKP APIs.

> **Notes:**
>
> * The **API Tester** in the v4 Portal enables you to securely call the methods of these APIs using your organization's Kentik data.
> * For information about MKP, start with **My Kentik Portal**.
> * Protobuf and OpenAPI specifications for Kentik’s v6 APIs are available in our **api-schema-public** repository.

## MKP Usage

The topics below provide important background information for the use of these APIs.

### Overview

My Kentik Portal API enables programmatic access to tenants and package templates.

| Endpoint | Purpose |
| --- | --- |
| TenantService | CRUD operations for MKP tenants. |
| PackageService | CRUD operations for MKP templates. |

Both REST endpoint and gRPC RPCs are provided.

### Tenant Attributes and Settings

| Attribute | Access | Purpose |
| --- | --- | --- |
| id | RO | System-generated unique identifier of the tenant |
| company\_id | RO | System-generated unique identifier of the company |
| name | RW | User specified name for the tenant |
| description | RW | User specified description for the tenant |
| type | RO | subtenant |
| enabled | RW | User specified tenant active status |
| cdate | RO | Creation timestamp |
| edate | RO | Last-modification timestamp |
| alerts | RO | Alert policy configurations |
| assets | RW | dashboard, views, and templates associate with tenant |
| asn | RW | ASN data source |
| cidr | RW | CIDR data source |
| custom\_dimensions | RW | Custom dimension data source |
| devices | RW | Devices data source |
| filters | RW | Data source filters |
| interface\_name | RW | Interface source filters |
| snmp\_alias | RW | SNMP source filters |
| packages | RO | Package templates use by the tenant |
| users | RO | Users assign to the tenant |
| template\_id | WO | Update tenant's packages setting with package/template id |

### Package Attributes and Settings

| Attribute | Access | Purpose |
| --- | --- | --- |
| id | RO | System-generated unique identifier of the package |
| company\_id | RO | System-generated unique identifier of the company |
| name | RW | User specified name for the package template |
| description | RW | User specified description for the package template |
| icon | RO | User specified icon |
| color | RW | User specified color of icon |
| alerts | RO | Alert policy configurations |
| assets | RW | dashboard, views, and templates associate with tenant |
| is\_default | RW | Default package template to include for new tenant |
| tenants | RO | Tenants using this package template |

## MKP RPCs

The tables below provide a quick reference to key information about each Remote Procedure Call in these APIs. Use the links in the tables to see schemas for the request and/or response body (if any).

> **Note:** To test methods using your own Kentik data, use the portal's **API Tester**.

### PackageList

**API: PackageService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| GET | /mkp/v202407 /packages | Returns a list of MKP packages. |
| |  |  | | --- | --- | | **Request body:**  None | **Response body:**  v202407ListPackageResponse | | **Parameters**:  None | | | | |

### PackageCreate

**API: PackageService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| POST | /mkp/v202407 /packages | Create package from request. returns created package. |
| |  |  | | --- | --- | | **Request body:**  v202407CreatePackageRequest | **Response body:**  v202407CreatePackageResponse | | **Parameters**:  None |  | | | |

### PackageGet

**API: PackageService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| GET | /mkp/v202407 /packages/{id} | Returns information about package specified with ID. |
| |  |  | | --- | --- | | **Request body:**  None | **Response body**:  v202407GetPackageResponse |    **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | id | Undefined. | true | string | | | |

### PackageDelete

**API: PackageService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| DELETE | /mkp/v202407 /packages/{id} | Deletes the package specified with id. |
| |  |  | | --- | --- | | **Request body:**  None | **Response body:**  v202407DeletePackageResponse |  **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | id | Undefined. | true | string | | | |

### PackageUpdate

**API: PackageService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| PUT | /mkp/v202407 /packages/{id} | Update package attributes specified with id. |
| |  |  | | --- | --- | | **Request body:**  v202407UpdatePackageRequest | **Response body:**  v202407UpdatePackageResponse |  **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | id | Unique system assigned identifier of the package | true | string | | | |

### TenantList

**API: TenantService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| GET | /mkp/v202407 /tenants | Returns a list of MKP tenants. |
| |  |  | | --- | --- | | **Request body:**  None | **Response body:**  v202407ListTenantResponse | | **Parameters**:  None | | | | |

### TenantCreate

**API: TenantService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| POST | /mkp/v202407 /tenants | Create tenant from request. returns created tenant. |
| |  |  | | --- | --- | | **Request body:**  v202407CreateTenantRequest | **Response body:**  v202407CreateTenantResponse | | **Parameters**:  None | | | | |

### TenantGet

**API: TenantService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| GET | /mkp/v202407 /tenants/{id} | Returns information about package specified with ID. |
| |  |  | | --- | --- | | **Request body:**  None | **Response body:**  v202407GetTenantResponse |  **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | id | Undefined. | true | string | | | |

### TenantDelete

**API: TenantService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| DELETE | /mkp/v202407 /tenants/{id} | Deletes the tenant specified with id. |
| |  |  | | --- | --- | | **Request body:**  None | **Response body:**  v202407DeleteTenantResponse |  **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | id | Undefined. | true | string | | | |

### TenantUpdate

**API: TenantService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| PUT | /mkp/v202407 /tenants/{id} | Update tenant attributes specified with id. |
| |  |  | | --- | --- | | **Request body:**  v202407UpdateTenantRequest | **Response body:**  v202407UpdateTenantResponse |  **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | id | Unique system assigned identifier of the tenant | true | string | | | |

## MKP Schemas

This API uses the following schemas.

#### AssetReport

|  |  |
| --- | --- |
| **Schema:** AssetReport | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | id | type: string | | type | type: string | | |

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

#### PermissionEntry

|  |  |
| --- | --- |
| **Schema:** v202211PermissionEntry | **Type:** object |
| **Properties** (\*  =  required)  | Name | Value | | --- | --- | | capability \* | type: string  description: String identifying capability that is granted of denied | | allowed \* | type: boolean  description: Flag indicating whether operation is allowed | | |

#### Role

|  |  |
| --- | --- |
| **Schema:** v202211Role | **Type:** string |
| **Attributes**:  | Key | Value | | --- | --- | | enum | ROLE\_UNSPECIFIED, ROLE\_MEMBER, ROLE\_ADMINISTRATOR, ROLE\_SUPER\_ADMINISTRATOR | | default | ROLE\_UNSPECIFIED | | description | •  ROLE\_UNSPECIFIED: Invalid value.  • ROLE\_MEMBER: Member  • ROLE\_ADMINISTRATOR: Administrator  • ROLE\_SUPER\_ADMINISTRATOR: Super-administrator | | |

#### User

|  |  |
| --- | --- |
| **Schema:** v202211User | **Type:** object |
| **Properties** (\*  =  required)  | Name | Value | | --- | --- | | id | type: string  description: System generated unique identifier  readOnly: true | | userEmail \* | type: string  description: User e-mail address (serves also as username) | | userFullName \* | type: string  description: Full name | | role | $ref: v202211Role | | permissions | type: array  items: $ref: v202211PermissionEntry  description: Optional list of permissions granted to the user | | filter | type: string  description: Optional JSON string defining filter for objects visible to the user | | lastLogin | type: string  format: date-time  description: UTC Timestamp of user's last login session  readOnly: true | | cdate | type: string  format: date-time  description: Creation timestamp (UTC)  readOnly: true | | edate | type: string  format: date-time  description: Last modification timestamp (UTC)  readOnly: true | | |

#### Activate

|  |  |
| --- | --- |
| **Schema:** v202407Activate | **Type:** object |
| **Properties:**  | Name | Value | | --- | --- | | times | type: integer  format: int64 | | operator | type: string | | timeWindowSeconds | type: integer  format: int64 | | gracePeriodSeconds | type: integer  format: int64 | | |

#### Alert

|  |  |
| --- | --- |
| **Schema:** v202407Alert | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | saved | type: boolean | | policyId | type: string | | thresholds | type: array  items: $ref: v202407Threshold | | primaryMetric | type: string | | secondaryMetrics | type: array  items: type: string | | isTemplate | type: boolean | | subpolicyId | type: string | | |

#### Asset

|  |  |
| --- | --- |
| **Schema:** v202407Asset | **Type:** object |
| **Properties:**  | Name | Value | | --- | --- | | reports | type: array  items: $ref: AssetReport | | defaultReport | $ref: AssetReport | | |

#### Condition

|  |  |
| --- | --- |
| **Schema:** v202407Condition | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | type | type: string | | value | type: string | | metric | type: string | | operator | type: string | | valueType | type: string | | valueSelect | type: string | | |

#### CreatePackageRequest

|  |  |
| --- | --- |
| **Schema:** v202407CreatePackageRequest | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | package | $ref: v202407Package | | |

#### CreatePackageResponse

|  |  |
| --- | --- |
| **Schema:** v202407CreatePackageResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | package | ref: v202407Package | | |

#### CreateTenantRequest

|  |  |
| --- | --- |
| **Schema:** v202407CreateTenantRequest | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | tenant | $ref: v202407Tenant | | |

#### CreateTenantResponse

|  |  |
| --- | --- |
| **Schema:** v202407CreateTenantResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | tenant | $ref: v202407Tenant | | |

#### CustomDimension

|  |  |
| --- | --- |
| **Schema:** v202407CustomDimension | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | dimension | type: string | | populator | type: string | | |

#### DeletePackageResponse

|  |  |
| --- | --- |
| **Schema:** v202407DeletePackageResponse | **Type:** object |
| **Properties**:  None | |

#### DeleteTenantResponse

|  |  |
| --- | --- |
| **Schema:** v202407DeleteTenantResponse | **Type:** object |
| **Properties**:  None | |

#### Devices

|  |  |
| --- | --- |
| **Schema:** v202407Devices | **Type:** object |
| **Properties:**  | Name | Value | | --- | --- | | allDevices | type: boolean | | deviceTypes | type: array  items: type: string | | deviceLabels | type: array  items: type: integer  format: int64 | | deviceSites | type: array  items: type: integer  format: int64 | | deviceName | type: array  items: type: string | | |

#### Filter

|  |  |
| --- | --- |
| **Schema:** v202407Filter | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | name | type: string | | named | type: boolean | | connector | type: string | | not | type: boolean | | autoAdded | type: string | | savedFilters | type: array  items:type: string | | filters | type: array  items: $ref: v202407FilterField | | filterGroups | type: array  items: $ref: v202407Filter | | metric | type: array  items: type: string | | |

#### FilterField

|  |  |
| --- | --- |
| **Schema:** v202407FilterField | **Type:** object |
| **Properties:**  | Name | Value | | --- | --- | | filterField | type: string | | operator | type: string | | filterValue | type: string | | |

#### GetPackageResponse

|  |  |
| --- | --- |
| **Schema:** v202407GetPackageResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | package | $ref: v202407Package | | |

#### GetTenantResponse

|  |  |
| --- | --- |
| **Schema:** v202407GetTenantResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | tenant | $ref: v202407Tenant | | |

#### ListPackageResponse

|  |  |
| --- | --- |
| **Schema:** v202407ListPackageResponse | **Type:** object |
| **Properties:**  | Name | Value | | --- | --- | | packages | type: array  items: $ref: v202407Package | | invalidCount | type: integer  format: int64  description: The number of invalid packages, for troubleshooting. Should be zero. | | |

#### ListTenantResponse

|  |  |
| --- | --- |
| **Schema:** v202407ListTenantResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | tenants | type: array  items: $ref: v202407Tenant | | invalidCount | type: integer  format: int64 | | |

#### Mitigation

|  |  |
| --- | --- |
| **Schema:** v202407Mitigation | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | id | type: string | | status | type: string | | companyId | type: string | | pairingId | type: string | | thresholdId | type: string | | isMethodOverridable | type: boolean | | mitigationApplyType | type: string | | mitigationClearType | type: string | | mitigationApplyTimer | type: integer  format: int64 | | mitigationClearTimer | type: integer  format: int64 | | isPlatformOverridable | type: boolean | | |

#### NotificationChannel

|  |  |
| --- | --- |
| **Schema:** v202407NotificationChannel | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | id | type: string | | |

#### Package

|  |  |
| --- | --- |
| **Schema:** v202407Package | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | id | type: string  description: Unique system assigned identifier of the package  readOnly: true | | companyId | type: string  description: Company id of the associated package  readOnly: true | | name | type: string  description: Name of package template  readOnly: true | | description | type: string  description: Description of package template  readOnly: true | | icon | type: string  description: Icon to display  readOnly: true | | color | type: string  description: Color of Icon  readOnly: true | | alerts | type: array  items: $ref: v202407Alert  description: Alert thresholds and policies  readOnly: true | | assets | $ref: v202407Asset | | isDefault | type: boolean  description: Is default template  readOnly: true | | tenants | type: array  items: $ref: v202407TenantLink  description: Tenant link information  readOnly: true | | |

#### Tenant

|  |  |
| --- | --- |
| **Schema:** v202407Tenant | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | id | type: string  description: Unique system assigned identifier of the tenant  readOnly: true | | companyId | type: string  description: Company id of the associated package  readOnly: true | | name | type: string  description: Name of tenant  readOnly: true | | description | type: string  description: Description of tenant  readOnly: true | | type | type: string  description: Subtenant  readOnly: true | | enabled | type: boolean  description: Is tenant actively enable  readOnly: true | | alerts | type: array  items: $ref: v202407Alert  description: Alert thresholds and policies  title: config object  readOnly: true | | assets | $ref: v202407Asset | | asn | type: string  description: ASN data source.  readOnly: true | | cidr | type: string  description: CIDR data source.  readOnly: true | | customDimensions | type: array  items: $ref: v202407CustomDimension  description: Custom dimension data source  readOnly: true | | devices | $ref: v202407Devices | | filters | $ref: v202407Filter | | interfaceName | type: string  description: Interface data source.  readOnly: true | | snmpAlias | type: string  description: SNMP data source.  readOnly: true | | packages | type: array  items: $ref: v202407Package  description: Packages associated with the tenant.  readOnly: true | | users | type: array  items: $ref: v202211User  description: Users associated with the tenant.  readOnly: true | | templateId | type: string  description: Package template ID to assign with tenant.  readOnly: true | | |

#### TenantLink

|  |  |
| --- | --- |
| **Schema:** v202407TenantLink | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | id | type: string | | pivotTemplateId | type: string | | pivotUserGroupId | type: string | | |

#### Threshold

|  |  |
| --- | --- |
| **Schema:** v202407Threshold | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | id | type: string | | activate | $ref: v202407Activate | | severity | type: string | | conditions | type: array  items: $ref: v202407Condition | | mitigations | type: array  items: $ref: v202407Mitigation | | notificationChannels | type: array  items: $ref: v202407NotificationChannel | | thresholdAckRequired | type: boolean | | enableTenantNotifications | type: boolean | | receiveLandlordNotifications | type: boolean | | |

#### UpdatePackageRequest

|  |  |
| --- | --- |
| **Schema:** v202407UpdatePackageRequest | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | package | $ref: v202407Package | | |

#### UpdatePackageResponse

|  |  |
| --- | --- |
| **Schema:** v202407UpdatePackageResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | package | $ref: v202407Package | | |

#### UpdateTenantRequest

|  |  |
| --- | --- |
| **Schema:** v202407UpdateTenantRequest | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | tenant | $ref: v202407Tenant | | |

#### UpdateTenantResponse

|  |  |
| --- | --- |
| **Schema:** v202407UpdateTenantResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | tenant | $ref: v202407Tenant | | |
