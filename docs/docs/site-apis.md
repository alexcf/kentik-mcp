# Site APIs

Source: https://kb.kentik.com/docs/site-apis

---

This article covers how to get started with the Site APIs.

> **Notes:**
>
> * The **API Tester** in the v4 Portal enables you to securely call the methods of these APIs using your organization's Kentik data.
> * For information about sites, start with **About Sites**.
> * Protobuf and OpenAPI specifications for Kentik’s v6 APIs are available in our **api-schema-public** repository.

## Site Usage

The topics below provide important background information for the use of these APIs.

### Overview

The Site Configuration API provides programmatic access to configuration of Sites and Site Markets:

* **Site**: A site is a specific user-defined physical location (e.g., a data center at a given address) or logical location at which there is hosting of your devices, services, providers, or partner networks. Information that can be associated with a site includes the classification of IP addresses and the site's logical network topology/architecture.
* **Site Market**: Sites with any common characteristics of your choosing (e.g., all PoPs in a particular region) can be logically grouped into a site market.

Both REST endpoint and gRPC RPCs are provided.

> **Notes:**
>
> * Once a site is created in Kentik, you can assign one or more devices to the site via the settings for those devices; to do so programmatically you'd use the v5 Device API (see **Device Create**).
> * Sites can also be managed using the legacy REST-only **Site API**, which provides access to a subset of site configuration attributes.

## Site RPCs

The tables below provide a quick reference to key information about each Remote Procedure Call in these APIs. Use the links in the tables to see schemas for the request and/or response body (if any).

> **Note:** To test methods using your own Kentik data, use the portal's **API Tester**.

### ListSiteMarkets

**API: SiteService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| GET | /site/v202211 /site\_markets | Returns list of configured site markets. |
| |  |  | | --- | --- | | **Request body:** None  **Parameters:** None | **Response body:**  v202211ListSiteMarketsResponse | | | |

### CreateSiteMarket

**API: SiteService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| POST | /site/v202211 /site\_markets | Create configuration for a new site market. Returns the newly created configuration. |
| |  |  | | --- | --- | | **Request body**: v202211CreateSiteMarketRequest  **Parameters**: None | **Response body**:  v202211CreateSiteMarketResponse | | | |

### GetSiteMarket

**API: SiteService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| GET | /site/v202211 /site\_markets/{id} | Returns configuration of a site market specified by ID. |
| |  |  | | --- | --- | | **Request body:** None | **Response body**:  v202211GetSiteMarketResponse |  **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | id | ID of the requested site market | true | string | | | |

### DeleteSiteMarket

**API: SiteService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| DELETE | /site/v202211 /site\_markets/{id} | Deletes configuration of a site market with specific ID. |
| |  |  | | --- | --- | | **Request body**: None | **Response body**:  v202211DeleteSiteMarketResponse |  **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | id | ID of the site market to be deleted | true | string | | | |

### UpdateSiteMarket

**API: SiteService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| PUT | /site/v202211 /site\_markets/{id} | Replaces configuration of a site market with attributes in the request. Returns the updated configuration. |
| |  |  | | --- | --- | | **Request body**:  v202211UpdateSiteMarketRequest | **Response body**:  v202211UpdateSiteMarketResponse |  **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | id | System generated unique identifier | true | string | | | |

### ListSites

**API: SiteService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| GET | /site/v202211 /sites | Returns list of configured sites. |
| |  |  | | --- | --- | | **Request body:** None **Parameters:** None | **Response body:**  v202211ListSitesResponse | | | |

### CreateSite

**API: SiteService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| POST | /site/v202211 /sites | Create configuration for a new site. Returns the newly created configuration. |
| |  |  | | --- | --- | | **Request body**: v202211CreateSiteRequest **Parameters**: None | **Response body**:  v202211CreateSiteResponse | | | |

### GetSite

**API: SiteService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| GET | /site/v202211 /sites/{id} | Returns configuration of a site specified by ID. |
| |  |  | | --- | --- | | **Request body**: None | **Response body**:  v202211GetSiteResponse |  **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | id | ID of the requested site | true | string | | | |

### DeleteSite

**API: SiteService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| DELETE | /site/v202211 /sites/{id} | Deletes configuration of a site with specific ID. |
| |  |  | | --- | --- | | **Request body**: None | **Response body**:  v202211DeleteSiteResponse |  **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | id | ID of the site to be deleted | true | string | | | |

### UpdateSite

**API: SiteService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| PUT | /site/v202211 /sites/{id} | Replaces configuration of a site with attributes in the request. Returns the updated configuration. |
| |  |  | | --- | --- | | **Request body**: v202211UpdateSiteRequest | **Response body**:  v202211UpdateSiteResponse |  **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | id | System generated unique identifier | true | string | | | |

## Site Schemas

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

#### CreateSiteMarketRequest

|  |  |
| --- | --- |
| **Schema:** v202211CreateSiteMarketRequest | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | siteMarket | $ref: v202211SiteMarket | | |

#### CreateSiteMarketResponse

|  |  |
| --- | --- |
| **Schema:** v202211CreateSiteMarketResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | siteMarket | $ref: v202211SiteMarket | | |

#### CreateSiteRequest

|  |  |
| --- | --- |
| **Schema:** v202211CreateSiteRequest | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | site | $ref: v202211Site | | |

#### CreateSiteResponse

|  |  |
| --- | --- |
| **Schema:** v202211CreateSiteResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | site | $ref: v202211Site | | |

#### DeleteSiteMarketResponse

|  |  |
| --- | --- |
| **Schema:** v202211DeleteSiteMarketResponse | **Type:** object |
| **Properties**: None. | |

#### DeleteSiteResponse

|  |  |
| --- | --- |
| **Schema:** v202211DeleteSiteResponse | **Type:** object |
| **Properties**: None. | |

#### GetSiteMarketResponse

|  |  |
| --- | --- |
| **Schema:** v202211GetSiteMarketResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | siteMarket | $ref: v202211SiteMarket | | |

#### GetSiteResponse

|  |  |
| --- | --- |
| **Schema:** v202211GetSiteResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | site | $ref: v202211Site | | |

#### Layer

|  |  |
| --- | --- |
| **Schema:** v202211Layer | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | name | type: string  description: Name of the network layer | | deviceIds | type: array  items: type: string  description: IDs of devices that are deemed to be part of the network layer | | |

#### LayerSet

|  |  |
| --- | --- |
| **Schema:** v202211LayerSet | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | layers | type: array  items: $ref: v202211Layer  description: List of parallel network layers | | |

#### ListSiteMarketsResponse

|  |  |
| --- | --- |
| **Schema:** v202211ListSiteMarketsResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | siteMarkets | type: array  items: $ref: v202211SiteMarket  description: List of configurations of requested site markets | | invalidCount | type: integer  format: int64  description: Number of invalid entries encountered while collecting data | | |

#### ListSitesResponse

|  |  |
| --- | --- |
| **Schema:** v202211ListSitesResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | sites | type: array  items: $ref: v202211Site  description: List of configurations of requested sites | | invalidCount | type: integer  format: int64  description: Number of invalid entries encountered while collecting data | | |

#### PostalAddress

|  |  |
| --- | --- |
| **Schema:** v202211PostalAddress | **Type:** object |
| **Properties (\*  =  required)**  | Name | Value | | --- | --- | | address \* | type: string  description: Street address | | city \* | type: string  description: City (full name) | | region | type: string  description: Geographical region | | postalCode | type: string  description: Country specific postal code | | country \* | type: string  description: Country (full name or country code) | | |

#### Site

|  |  |
| --- | --- |
| **Schema:** v202211Site | **Type:** object |
| **Properties (\*  =  required)**  | Name | Value | | --- | --- | | id | type: string  description: System generated unique identifier  readOnly: true | | title \* | type: string  description: User selected title/name | | lat | type: number  format: double  description: Latitude (signed decimal degrees) | | lon | type: number  format: double  description: Longitude (signed decimal degrees) | | postalAddress | $ref: v202211PostalAddress | | type | $ref: v202211SiteType | | addressClassification | $ref: v202211SiteIpAddressClassification | | architecture | type: array  items: $ref: v202211LayerSet  description: Logical network topology/architecture | | siteMarket | type: string  description: Name of the Site Market this sire belongs to | | |

#### SiteIpAddressClassification

|  |  |
| --- | --- |
| **Schema:** v202211SiteIpAddressClassification | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | infrastructureNetworks | type: array  items: type: string  description List of IP address prefixes (in standard CIDR notation) used in infrastructure networks | | userAccessNetworks | type: array  items: type: string  description: List of IP address prefixes (in standard CIDR notation) used in access networks | | otherNetworks | type: array  items: type: string  description: List of IP address prefixes (in standard CIDR notation) used in other networks | | |

#### SiteMarket

|  |  |
| --- | --- |
| **Schema:** v202211SiteMarket | **Type:** object |
| **Properties (\*  =  required)**  | Name | Value | | --- | --- | | id | type: string  description: System generated unique identifier  readOnly: true | | name \* | type: string  description: User selected unique name | | description | type: string  description: Free-form description | | numberOfSites | type: integer  format: int64  description: Number of sites in this market  readOnly: true | | cdate | type: string  format: date-time  description: Creation timestamp (UTC)  readOnly: true | | edate | type: string  format: date-time  description: Last modification timestamp (UTC)  readOnly: true | | |

#### SiteType

|  |  |
| --- | --- |
| **Schema:** v202211SiteType | **Type:** string |
| **Attributes**:  | Key | Value | | --- | --- | | enum | SITE\_TYPE\_UNSPECIFIED, SITE\_TYPE\_DATA\_CENTER, SITE\_TYPE\_CLOUD, SITE\_TYPE\_BRANCH, SITE\_TYPE\_CONNECTIVITY, SITE\_TYPE\_CUSTOMER, SITE\_TYPE\_OTHER | | default | SITE\_TYPE\_UNSPECIFIED | | description | • SITE\_TYPE\_UNSPECIFIED: Invalid value.  • SITE\_TYPE\_DATA\_CENTER: Data center site type.  • SITE\_TYPE\_CLOUD: Cloud site type.  • SITE\_TYPE\_BRANCH: Branch office site type.  • SITE\_TYPE\_CONNECTIVITY: Connectivity/PoP site type.  • SITE\_TYPE\_CUSTOMER: Customer/partner site type.  • SITE\_TYPE\_OTHER: Other site type. | | |

#### UpdateSiteMarketRequest

|  |  |
| --- | --- |
| **Schema:** v202211UpdateSiteMarketRequest | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | siteMarket | $ref: v202211SiteMarket | | |

#### UpdateSiteMarketResponse

|  |  |
| --- | --- |
| **Schema:** v202211UpdateSiteMarketResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | siteMarket | $ref: v202211SiteMarket | | |

#### UpdateSiteRequest

|  |  |
| --- | --- |
| **Schema:** v202211UpdateSiteRequest | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | site | $ref: v202211Site | | |

#### UpdateSiteResponse

|  |  |
| --- | --- |
| **Schema:** v202211UpdateSiteResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | site | $ref: v202211Site | | |
