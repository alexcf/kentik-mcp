# Cloud Export APIs

Source: https://kb.kentik.com/docs/cloud-export-apis

---

This article covers how to get started with the Cloud Export APIs.

> **Notes:**
>
> * The **API Tester** in the v4 Portal enables you to securely call the methods of these APIs using your organization's Kentik data.
> * For information about cloud exports, start with **Cloud Overview**.
> * Protobuf and OpenAPI specifications for Kentik’s v6 APIs are available in our **api-schema-public** repository.

## Cloud Export Usage

The topics below provide important background information for the use of these APIs.

### Overview

In Kentik, a "cloud export" is an object whose properties are the values that Kentik needs to access network flow logs from a given set of resources in a given cloud provider (see **Cloud Exports and Devices**). The Cloud Export API enables programmatic management of cloud exports, serving two primary functions:

* Configuration of the Kentik resources required to export network flow logs and metadata from **Public Clouds**.
* Basic status information on active export processes.

#### Supported Cloud Providers

The functionality and configuration parameters supported by this API differ by cloud provider. The API currently supports the following providers:

* **Amazon Web Services** (AWS)
* **Microsoft Azure**
* **Google Cloud** (GCP)
* **Oracle Cloud Infrastructure** (OCI)

#### Additional Public Resources

The following additional resources are available for working with this API:

* Kentik community **Python** and **Go** SDKs provide language-specific support for using this and other Kentik APIs. These SDKs can be also used as example code for development.
* A **Terraform provider** is available for configuring `cloud_export` instances in Kentik. Terraform modules are available for **AWS**, **Azure**  and **GCP**. These modules support onboarding and ongoing management of all resources (in public clouds and in Kentik) required to export flow logs to Kentik.

### Anatomy of a Cloud Export

Configuration and status of a `cloud_export` instance is represented by the `CloudExport` object, which contains three categories of attributes:

* Common Configuration Attributes
* Cloud Provider Configuration Attributes
* Metadata and Status Attributes

#### Common Configuration Attributes

The configuration attributes in the table below are common to `CloudExport` objects for all cloud providers.

| Attribute | Required | Default |
| --- | --- | --- |
| enabled | No | False |
| name | Yes |  |
| description | No | <empty string> |
| type | Yes |  |
| cloud\_provider | Yes |  |
| plan\_id | Yes |  |

#### Cloud Provider Specific Attributes

The attributes listed in the table below are objects whose parameters are specific to each cloud provider.

| Cloud Provider | Attribute | Object Name |
| --- | --- | --- |
| AWS | aws | AwsProperties |
| Azure | azure | AzureProperties |
| Google Cloud | gce | GceProperties |
| Oracle Cloud Infrastructure | oci | OciProperties |

#### Metadata and Status Attributes

The attributes in the table below provide read-only metadata and status for a `cloud_export` instance.

| Attribute | Purpose |
| --- | --- |
| id | System-generated unique identifier of the instance |
| cdate | Creation timestamp |
| edate | Last-modification timestamp |
| current\_status | Runtime status of the flow and metadata collection process |

## Cloud Export RPCs

The tables below provide a quick reference to key information about each Remote Procedure Call in these APIs. Use the links in the tables to see schemas for the request and/or response body (if any).

> **Note:** To test methods using your own Kentik data, use the portal's **API Tester**.

### ListCloudExports

**API: CloudExportAdminService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| GET | /cloud\_export/v202210 /exports | Returns a list of all cloud exports in the account. |
| |  |  | | --- | --- | | **Request body**:  None  **Parameters**:  None | **Response body**:  v202210ListCloudExportsResponse | | | |

### CreateCloudExport

API: CloudExportAdminService

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| POST | /cloud\_export/v202210 /exports | Create new cloud export based on configuration in the request. |
| |  |  | | --- | --- | | **Request body**:  v202210CreateCloudExportRequest  **Parameters**:  None | **Response body**:  v202210CreateCloudExportResponse | | | |

### UpdateCloudExport

API: CloudExportAdminService

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| PUT | /cloud\_export/v202210 /exports/{export.id} | Replace complete configuration of a cloud export with data in the request. |
| |  |  | | --- | --- | | **Request body**:  v202210UpdateCloudExportRequest | **Response body**:  v202210UpdateCloudExportResponse |  **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | export.id | Unique identifier of the instance | true | string | | | |

### GetCloudExport

**API: CloudExportAdminService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| GET | /cloud\_export/v202210 /exports/{export.id} | Returns configuration and status of cloud export with specified ID. |
| |  |  | | --- | --- | | **Request body**:  None | **Response body**:  v202210GetCloudExportResponse |    **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | export.id | ID of the cloud export to be retrieved. | true | string | | | |

### DeleteCloudExport

**API: CloudExportAdminService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| DELETE | /cloud\_export/v202210 /exports/{export.id} | Delete cloud export with specified ID. |
| |  |  | | --- | --- | | **Request body**:  None | **Response body**:  v202210DeleteCloudExportResponse |    **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | export.id | ID of the cloud export to be deleted. | true | string | | | |

## Cloud Export Schemas

This API uses the following schemas:

#### protobufAny

|  |  |
| --- | --- |
| **Schema:** protobufAny | **Type:** object |
| **Properties:**  | Name | Value | | --- | --- | | typeUrl | type: string | | value | type: string  format: byte | | |

#### rpcStatus

|  |  |
| --- | --- |
| **Schema:** rpcStatus | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | code | type: integer  format: int32 | | message | type: string | | details | type: array  items: $ref: protobufAny | | |

#### AwsProperties

|  |  |
| --- | --- |
| **Schema:** v202210AwsProperties | **Type:** object |
| **Properties** (\*  =  required)  | Name | Value | | --- | --- | | bucket | type: string  description: Name of S3 bucket from which flow logs are to be exported. | | iamRoleArn \* | type: string  description: ARN of the IAM role granted access to the S3 bucket and describe API end-points. | | region \* | type: string  description: Name of AWS region from which to export flow logs. | | deleteAfterRead | type: boolean  description: Delete from logs from the S3 bucket after export (default false). | | metadataOnly | type: boolean  description: Import only metadata without any flows (default false). | | awsIamRoleArnIsOrg | type: boolean  description: Iam role is organization role (default false). | | secondaryAwsAccounts | type: array  items: type: string  description: Accounts ids granted access to the describe API end-points. | | secondaryAwsBlockedAccounts | type: array  items: type: string  description: Accounts ids that should be filtered from organization when running describe API end-points. | | secondaryAwsRegions | type: array  items: type: string  description: Regions that secondary accounts to scrape. | | secondaryAwsSuffix | type: string  description: Used to generate secondary account ARN based on template arn:aws:iam::<>:role/<> | | |

#### AzureProperties

|  |  |
| --- | --- |
| **Schema:** v202210AzureProperties | **Type:** object |
| **Properties** (\*  =  required)  | Name | Value | | --- | --- | | location \* | type: string  description: Azure region/location from which to export flow logs. | | resourceGroup \* | type: string  description: Resource group containing the NSG generating flow logs. | | storageAccount \* | type: string  description: Storage account from which flow logs are to be extracted. | | subscriptionId \* | type: string  description: ID of Azure account from which flows logs are to be exported. | | securityPrincipalEnabled | type: boolean  description: Indication whether security principal for the Kentik flow export application has been authorized. | | |

#### CloudExport

|  |  |
| --- | --- |
| **Schema:** v202210CloudExport | **Type:** object |
| **Properties** (\*  =  required)  | Name | Value | | --- | --- | | id | type: string  description: Unique identifier of the instance  readOnly: true | | type | $ref: v202210CloudExportType | | enabled \* | type: boolean  description: Attribute controlling whether the instance is active | | name \* | type: string  description: User selected name of the instance | | description | type: string  description: Description of the instance | | planId \* | type: string  description: Identifier of the billing plan for the instance | | cloudProvider | $ref: v202210CloudProvider | | aws | $ref: v202210AwsProperties | | azure | $ref: v202210AzureProperties | | gce | $ref: v202210GceProperties | | ibm | $ref: v202210IbmProperties  description: Deprecated: IBM Cloud exports are no longer supported. | | oci | $ref: v202210OciProperties | | currentStatus | $ref: v202210CloudExportStatus | | cdate | type: string  format: date-time  description: Creation timestamp (UTC)  readOnly: true | | edate | type: string  format: date-time  description: Last modification timestamp (UTC)  readOnly: true | | |

#### CloudExportStatus

|  |  |
| --- | --- |
| **Schema:** v202210CloudExportStatus | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | status | type: string  description: Status of the export task.  readOnly: true | | errorMessage | type: string  description: Text of the last error message (empty if status is OK).  readOnly: true | | flowFound | type: boolean  description: Indication whether any flow data were exported.  readOnly: true | | apiAccess | type: boolean  description: Indication whether the export process is able to access cloud API.  readOnly: true | | storageAccountAccess | type: boolean  description: Indication whether the export process is able to access storage account containing flow logs.  readOnly: true | | |

#### CloudExportType

|  |  |
| --- | --- |
| **Schema:** v202210CloudExportType | **Type:** string |
| **Attributes**:  | Key | Value | | --- | --- | | enum | CLOUD\_EXPORT\_TYPE\_UNSPECIFIED, CLOUD\_EXPORT\_TYPE\_KENTIK\_MANAGED, CLOUD\_EXPORT\_TYPE\_CUSTOMER\_MANAGED | | default | CLOUD\_EXPORT\_TYPE\_UNSPECIFIED | | description | • CLOUD\_EXPORT\_TYPE\_UNSPECIFIED: Invalid value.  • CLOUD\_EXPORT\_TYPE\_KENTIK\_MANAGED: Cloud export process is managed by Kentik  • CLOUD\_EXPORT\_TYPE\_CUSTOMER\_MANAGED: Cloud export process is managed by customer | | |

#### CloudProvider

|  |  |
| --- | --- |
| **Schema:** v202210CloudProvider | **Type:** string |
| **Attributes**:  | Key | Value | | --- | --- | | enum | CLOUD\_PROVIDER\_UNSPECIFIED, CLOUD\_PROVIDER\_AWS, CLOUD\_PROVIDER\_AZURE, CLOUD\_PROVIDER\_GCE, CLOUD\_PROVIDER\_IBM, CLOUD\_PROVIDER\_OCI | | default | CLOUD\_PROVIDER\_UNSPECIFIED | | description | •  CLOUD\_PROVIDER\_UNSPECIFIED: Invalid value.  • CLOUD\_PROVIDER\_AWS: Amazon Web Services  • CLOUD\_PROVIDER\_AZURE: Microsoft Azure  • CLOUD\_PROVIDER\_GCE: Google Cloud  • CLOUD\_PROVIDER\_IBM: IBM Cloud   Deprecated: IBM Cloud exports are no longer supported.  • CLOUD\_PROVIDER\_OCI: OCI | | |

#### CreateCloudExportRequest

|  |  |
| --- | --- |
| **Schema:** v202210CreateCloudExportRequest | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | export | $ref: v202210CloudExport | | |

#### CreateCloudExportResponse

|  |  |
| --- | --- |
| **Schema:** v202210CreateCloudExportResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | export | $ref: v202210CloudExport | | |

#### DeleteCloudExportResponse

|  |  |
| --- | --- |
| **Schema:** v202210DeleteCloudExportResponse | **Type:** object |
| **Properties**: None. | |

#### GceProperties

|  |  |
| --- | --- |
| **Schema:** v202210GceProperties | **Type:** object |
| **Properties** (\*  =  required)  | Name | Value | | --- | --- | | project \* | type: string  description: Name of the project from which to export flow logs. | | subscription \* | type: string  description: GCP Pub/Sub subscription providing flow logs. | | |

#### GetCloudExportResponse

|  |  |
| --- | --- |
| **Schema:** v202210GetCloudExportResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | export | $ref: v202210CloudExport | | |

#### IbmProperties

|  |  |
| --- | --- |
| **Schema:** v202210IbmProperties | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | bucket | type: string  description: Storage bucket from which flow logs are to be extracted. | | |

#### ListCloudExportsResponse

|  |  |
| --- | --- |
| **Schema:** v202210ListCloudExportsResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | exports | type: array  items: $ref: v202210CloudExport  description: List of cloud export objects. | | invalidExportsCount | type: integer  format: int64  description: Number of objects with invalid data (which are not returned in the response). | | |

#### OciProperties

|  |  |
| --- | --- |
| **Schema:** v202210OciProperties | **Type:** object |
| **Properties** (\*  =  required)  | Name | Value | | --- | --- | | ociUserId \* | type: string  description: ID of the user created to represent the Kentik cloud export tool. | | ociTenancyId \* | type: string  description: OCI Tenancy details. | | ociCompartmentId | type: array  items: type: string  description: Compartment IDs to scrape metadata from. | | ociDefaultRegion \* | type: string  description: Default Enabled Region to scrape.. | | ociCollectFlowLogs | type: boolean  description: Default Enabled Region to scrape.. | | ociBucketName | type: string  description: Default Enabled Region to scrape.. | | ociBucketNamespaceName | type: string  description: Default Enabled Region to scrape.. | | ociServiceConnectorOcid | type: string  description: Default Enabled Region to scrape.. | | ociFlowObjectNamePrefix | type: string  description: We will look for files in the bucket whose names start with {%prefix%}/{%service\_connector\_ocid%} | | |

#### UpdateCloudExportRequest

|  |  |
| --- | --- |
| **Schema:** v202210UpdateCloudExportRequest | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | export | $ref: v202210CloudExport | | |

#### UpdateCloudExportResponse

|  |  |
| --- | --- |
| **Schema:** v202210UpdateCloudExportResponse | **Type:** object |
| **Properties**:  | Name | Name | | --- | --- | | export | $ref: v202210CloudExport | | |
