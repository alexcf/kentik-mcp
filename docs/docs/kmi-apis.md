# KMI APIs

Source: https://kb.kentik.com/docs/kmi-apis

---

This article covers how to get started with the KMI APIs.

> **Notes:**
>
> * The **API Tester** in the v4 Portal enables you to securely call the methods of these APIs using your organization's Kentik data.
> * For information about KMI, start with **Kentik Market Intelligence**.
> * Protobuf and OpenAPI specifications for Kentik’s v6 APIs are available in our **api-schema-public** repository.

## KMI Usage

The topics below provide important background information for the use of these APIs.

### Overview

The Kentik Market Intelligence (KMI) API provides programmatic access to information related to KMI rankings, KMI markets, and the customers, providers, and peers of individual Autonomous Systems (ASes). This information is derived from analysis of the global routing table, which enables us to classify the peering and transit relationships between ASes and to identify the providers, peers, and customers of a given AS in any geography (market). KMI estimates the volume of IP space transited by ASes in different geographies and produces rankings based on that volume, thereby enabling users to compare ASes in various markets.

### KMI Ranking Types

The following types of rankings are shown:

* **Customer Base**: Ranked by the size of the overall customer base, estimated by determining how much IP address space a given AS transits relative to other ASes:

  + Retail networks provide services (e.g., originate content) or have end-users that are consumers of services (e.g., ISPs or "eyeball" networks).
  + Wholesale networks connect retail networks to backbone networks.
  + Backbone networks carry high volumes of traffic between wholesale networks.
* **Customer Growth**: Ranked by the change in overall customer base (gain/loss of prefixes) over the last 20 days.
* **Peering**: Ranked by the amount of IP address space sent to the AS over a settlement-free peering session.

Both REST endpoint and gRPC RPCs are provided.

> **Note:** More information about KMI can be found at **Kentik Market Intelligence**.

## KMI RPCs

The tables below provide a quick reference to key information about each Remote Procedure Call in these APIs. Use the links in the tables to see schemas for the request and/or response body (if any).

> **Note:** To test methods using your own Kentik data, use the portal's **API Tester**.

### GetASNDetails

**API: KmiService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| POST | /kmi/v202212 /market/{marketId}/network/{asn}/{type} | Returns metadata and list of customers, providers, and peers for an Autonomous System. |
| |  |  | | --- | --- | | **Request body:**  v202212GetASNDetailsRequest | **Response body:**  v202212GetASNDetailsResponse |  **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | marketId | Unique Geo Market identifier (as provided by the ListMarkets RPC) | true | string | | asn | Autonomous System Number (ASN) | true | string | | type | Type of the requested ASN ('all', 'customer', 'provider', 'peer'). Defaults to 'all'. | true | string | | | |

### GetRankings

**API: KmiService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| POST | /kmi/v202212 /market/{marketId}/rankings/{rankType} | Returns list of KMI rankings. |
| |  |  | | --- | --- | | **Request body:**  v202212GetRankingsRequest | **Response body:**  v202212GetRankingsResponse |  **Parameters:**  | Name | Description | Required | Type | | --- | --- | --- | --- | | marketId | Unique Geo Market identifier (as provided by the ListMarkets RPC) | true | string | | rankType | Type of the requested ranking ('customer\_base', 'customer\_base\_retail', 'customer\_base\_wholesome', 'customer\_base\_backbone', 'customer\_growth', 'peering\_base'). Defaults to 'customer\_base'. | true | string | | | |

### ListMarkets

**API: KmiService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| GET | /kmi/v202212 /markets | Returns list of geo markets for KMI. |
| |  |  | | --- | --- | | **Request body:**  None | **Response body:**  v202212ListMarketsResponse | | **Parameters**:  None | | | | |

## KMI Schemas

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
| **Properties**:  | Name | Value | | --- | --- | | code | type: integer  format: int32 | | message | type: string | | details | type: array  items: $ref: protobufAny | | |

#### ASNDetails

|  |  |
| --- | --- |
| **Schema:** v202212ASNDetails | **Type:** object |
| **Properties:**  | Name | Value | | --- | --- | | asn | type: integer  format: int64  description: Autonomous System Number (ASN)  readOnly: true | | name | type: string  description: Name of the Autonomous System  readOnly: true | | countryName | type: string  description: Country Name of the Autonomous System  readOnly: true | | customers | type: array  items: $ref: v202212CustomerProvider  description: List of Customers  readOnly: true | | providers | type: array  items: $ref: v202212CustomerProvider  description: List of Providers  readOnly: true | | peers | type: array  items: $ref: v202212Peer  description: List of Peers  readOnly: true | | |

#### CustomerProvider

|  |  |
| --- | --- |
| **Schema:** v202212CustomerProvider | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | asn | type: integer  format: int64  description: Autonomous System Number (ASN)  readOnly: true | | name | type: string  description: Name of the Autonomous System  readOnly: true | | score | type: integer  format: int64  description: Score of the Autonomous System  readOnly: true | | singlehomedCustomer | type: boolean  description: Singlehomed customer (only one upstream provider to the internet)  readOnly: true | | mutualCustomer | type: boolean  description: Mutual customer  readOnly: true | | mutualProvider | type: boolean  description: Mutual provider  readOnly: true | | |

#### GetASNDetailsRequest

|  |  |
| --- | --- |
| **Schema:** v202212GetASNDetailsRequest | **Type:** object |
| **Properties** (\*  =  required)  | Name | Value | | --- | --- | | marketId \* | type: string  description: Unique Geo Market identifier (as provided by the ListMarkets RPC) | | asn \* | type: string  description: Autonomous System Number (ASN) | | ip | type: string  description: IP Address Family ('v4' or 'v6'). Defaults to 'v4'. | | type | type: string  description: Type of the requested ASN ('all', 'customer', 'provider', 'peer'). Defaults to 'all'. | | mutualProvider | type: string  description: Filter by mutual provider ('all', 'only', 'exclude'). Defaults to 'all'. | | mutualCustomer | type: string  description: Filter by mutual customer ('all', 'only', 'exclude'). Defaults to 'all'. | | singlehomedCustomer | type: string  description: Filter by singlehomed customer ('all', 'only', 'exclude'). Defaults to 'all'. | | |

#### GetASNDetailsResponse

|  |  |
| --- | --- |
| **Schema:** v202212GetASNDetailsResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | asnDetails | $ref: v202212ASNDetails | | |

#### GetRankingsRequest

|  |  |
| --- | --- |
| **Schema:** v202212GetRankingsRequest | **Type:** object |
| **Properties** (\*  =  required)  | Name | Value | | --- | --- | | marketId \* | type: string  description: Unique Geo Market identifier (as provided by the ListMarkets RPC) | | rankType | type: string  description: Type of the requested ranking ('customer\_base', 'customer\_base\_retail', 'customer\_base\_wholesome', 'customer\_base\_backbone', 'customer\_growth', 'peering\_base'). Defaults to 'customer\_base'. | | ip | type: string  description: IP Address Family ('v4' or 'v6') of requested ranking. Defaults to 'v4'. | | limit | type: integer  format: int64  description: Maximum number of entries returned. (Default: 600). | | |

#### GetRankingsResponse

|  |  |
| --- | --- |
| **Schema:** v202212GetRankingsResponse | **Type:** object |
| **Properties:**  | Name | Value | | --- | --- | | rankings | type: array  items: $ref: v202212Ranking  description: List of rankings | | invalidCount | type: integer  format: int64  description: Number of invalid entries encountered while collecting data | | |

#### ListMarketsResponse

|  |  |
| --- | --- |
| **Schema:** v202212ListMarketsResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | markets | type: array  items: $ref: v202212Market  description: Markets  readOnly: true | | |

#### Market

|  |  |
| --- | --- |
| **Schema:** v202212Market | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | marketId | type: string  description: Unique Geo Market identifier  readOnly: true | | name | type: string  description: Geo Market Name  readOnly: true | | |

#### Peer

|  |  |
| --- | --- |
| **Schema:** v202212Peer | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | asn | type: integer  format: int64  description: Autonomous System Number (ASN)  readOnly: true | | name | type: string  description: Name of the Autonomous System  readOnly: true | | pfxCount | type: integer  format: int64  description: Prefix Count (number of distinct IP address blocks announced by AS)  readOnly: true | | |

#### Ranking

|  |  |
| --- | --- |
| **Schema:** v202212Ranking | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | asn | type: integer  format: int64  description: Autonomous System Number (ASN)  readOnly: true | | name | type: string  description: Name of the Autonomous System  readOnly: true | | rank | type: integer  format: int64  description: Rank of the Autonomous System  readOnly: true | | rankChange | type: integer  format: int64  description: Rank Change of the Autonomous System  readOnly: true | | score | type: integer  format: int64  description: Score of the Autonomous System  readOnly: true | | scoreChange | type: integer  format: int64  description: Score Change of the Autonomous System  readOnly: true | | |
