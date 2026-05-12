# Customization APIs

Source: https://kb.kentik.com/docs/customization-apis

---

The Customization APIs of the Kentik V5 Admin APIs are covered here.

> **Notes:**
>
> * For an overview of all Kentik APIs, see **Kentik APIs**.
> * For information on using APIs with cURL, see **API Access Via cURL**.
> * For assistance using any API version please contact **support@kentik.com**.
> * The V5 API tester was discontinued in January 2025.

## About Customization APIs

The Kentik V5 Admin APIs offer APIs for programmatically managing settings under "Customize" in the Admin section of the Kentik portal. Available APIs include:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(731).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A24Z&se=2026-05-12T10%3A09%3A24Z&sr=c&sp=r&sig=4y%2BRrgYLG6Jm2gD6qRZyDQdzF6OeBmw%2Bju1IJD8U5rk%3D)**Custom Dimensions**: Managed via the **Custom Dimension API**, mirroring the **Custom Dimensions**portal settings.
* **Saved Filters**: Managed via the **Saved Filter API**, mirroring the **Saved Filters** portal settings.
* **Flow Tags**: Managed via the **Tag API**, mirroring the **Flow Tags**portal settings.

> **Note:** Interface Classification and Network Classification are not supported via APIs.

## Tag API

> **Notes:**
>
> * This API is deprecated. Use the v6 **Flow Tag APIs**instead.
> * Member-level users only have access to the GET methods of this API.

The Tag API allows programmatic management of your organization’s tags in Kentik. The methods for managing tags are detailed in the following topics:

### Tag JSON

Calls to the Tag API return an HTTP response with a JSON "tag" object, or an array of such objects for the Tag List call. The object contains fields as shown in the following example:

```
{
"tag": {
     "id": 664,
     "flow_tag": "INTF_NAME_TEST",
     "device_name": "",
     "interface_name": "",
     "addr": "",
     "port": "700",
     "tcp_flags": "",
     "protocol": "",
     "asn": "",
     "nexthop_asn": "",
     "bgp_aspath": "",
     "bgp_community": "",
     "user": "3584",
     "created_date": "2015-07-14T18:02:40.469Z",
     "updated_date": "2015-10-20T00:47:03.509Z",
     "company_id": "1289"
   }
 }
```

> **Note:** The returned object may not include all fields, depending on which fields are defined.

Each tag object contains fields with information about an individual tag registered with your organization. These fields are described in the following table:

| JSON name | Type | Description | Portal field |
| --- | --- | --- | --- |
| id | number | The system-assigned ID of the tag. | Tag ID |
| flow\_tag | string | Required: A name for the tag. When a match for any of the other tag parameters is found in the flow from a given device, this string will be added to the src\_flow\_tags and/or dst\_flow\_tags column in that device's KDE main tabl   * Length: min=2, max=20. * Valid characters: alphanumeric, hyphen, or underscores (no spaces). | Tag name |
| device\_name | string | A match results when the device\_name associated with a device sending flow data contains this string. | Device name |
| device\_type | string | A comma-delimited list of device types or regular expressions. A match results when any specified device type matches the device\_type associated with a device sending flow data. | Device type |
| site | string | A comma-separated list of sites or regular expressions. A match results when any specified site matches the site associated with a device sending flow data. | Site |
| interface\_name | string | A match results when the name or description of an interface sending flow data contains this string. | Interface name |
| addr | string | A range of IP addresses, expressed in CIDR notation (e.g. 38.12.34.0/24). A match results when this value corresponds to a range of IP addresses in incoming flow. | IP address |
| port | string | A port number, either source (SRC Port) or destination DST Port). A match results when this value appears within a port number in incoming flow. | Port |
| tcp\_flags | string | An integer number between 0 and 255 representing an 8-bit binary bit pattern corresponding to TCP flags. A match will result if the value in both the flow bit pattern and the bitmask is 1 at any of the eight places. | TCP flag |
| protocol | string | The protocol of traffic represented by a flow. The protocol of TCP is 6, and of UDP is 17. A match results when this value is the same as the protocol of the traffic represented by the flow. | Protocol name/number |
| asn | string | A comma-delimited list of ASNs (16- or 32-bit). A match results when any specified ASN is the same as the ASN of the last-hop router based on AS-PATH. | Last-hop (origin) ASN |
| lasthop\_as\_name | string | A comma-separated list of AS names or regular expressions. A match results when any specified AS name represents the name corresponding to the last ASN in the path in the routing table for either the source (SRC IP) or destination (DST IP). | Last-hop (origin) AS Name |
| nexthop\_asn | string | A comma-delimited list of ASNs (16- or 32-bit). A match results when any specified ASN is the same as the ASN of the next-hop router based on AS-PATH. | Next-hop ASN |
| nexthop\_as\_name | string | A comma-separated list of AS names or regular expressions. A match results when any specified AS name corresponds to the AS name of the next hop router based on AS path. | Next-hop AS Name |
| nexthop | string | A comma-separated list of IPv4 and/or IPv6 CIDRs. If a CIDR grouping (IPv4 or IPv6) is specified, a match can be on any address within that grouping. If no CIDR grouping is specified a match requires an exact IP.   * CIDRs may be expressed in "short form" (e.g. 1::2/127). | Next-hop IP |
| bgp\_aspath | string | A comma-delimited list of numbers or regular expressions representing BGP AS path. A match results when this value is the same as the BGP AS-PATH in the route.   * Permitted characters []\*:\_'$.0123456789()+?,space * Example: "'3737 1212,\_7801\_,2906$" would look for any of those 3 combinations in the AS path. | BGP AS path |
| bgp\_community | string | A comma-delimited list of numbers or regular expression representing BGP community (i.e. 2096:2212). A match results when this value is the same as the BGP community of the BGP route associated with incoming flow data.   * Permitted characters []\*:\_'$.0123456789()+?,space- | BGP community |
| mac | string | A comma-separated list of MAC Addresses. Results in a match if this value matches source or destination Ethernet (L2) address. | MAC Address |
| country | string | A comma-separated list of two-character country codes. Results in a match if this value includes a two-letter country code associated with the source or destination IP of the flow. | Country |
| vlans | string | A comma-separated list of VLAN IDs. Results in a match if this value includes a VLAN ID associated with the source or destination IP of the flow. | VLAN(s) |
| user | string | The system-assigned ID of the user who created the tag. | N.A. |
| created\_date | string | Date-time of tag creation, in UTC (ISO 8601), e.g. 2015-01-27T01:39:17.186Z | N.A. |
| updated\_date | string | Date-time of most-recent tag edit, in UTC, e.g. 2015-01-27T01:39:17.186Z | N.A. |
| company\_id | number | The system-assigned ID of the customer. | N.A. |

> **Notes:**
>
> * The Portal field column shows the corresponding setting name on the Add Tag or Edit Tag page in the portal (see **Add or Edit Tags**).
> * Commas are only valid as delimiters in comma-delimited lists within field values.
> * For more details on validations for the above field values, see **Tag Field Validation**.

### Tag List

The Tag List `GET` method retrieves a list of your Kentik tags as a JSON array, with each element representing an individual tag.

> **Note:** The "id" value in a tag’s structure can be used in subsequent calls to retrieve, update, or delete that tag.

**HTTP Request**  
The following table shows the path and HTTP request for this call (placeholders in italics):

|  |  |
| --- | --- |
| **URL** | https://api.kentik.com/api/v5/tags |
| **Request** | GET /api/v5/tags HTTP/1.1  Host: api.kentik.com  X-CH-Auth-API-Token: user\_api\_token  X-CH-Auth-Email: user@domain.suffix  Content-Type: application/json |

> **Note:** If your organization is registered on Kentik's EU cluster, use `api.kentik.eu` in place of `api.kentik.com` in the URL above.

**HTTP Response**  
A successful response from the Tag List method includes the following elements:

* Response headers
* HTTP response code
* A response body with a JSON tag array with an element containing information for each of your tags.

> **Note:** For details of the JSON name/value pairs in a `tag` object, see **Tag JSON**.

### Tag Info

The Tag Info `GET` method fetches details of a specific tag by ID from your list of Kentik tags.  
  
**HTTP Request**  
The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| **URL** | https://api.kentik.com/api/v5/tag/tag\_id |
| **Request** | GET /api/v5/tag/tag\_id HTTP/1.1  Host: api.kentik.com  X-CH-Auth-API-Token: user\_api\_token  X-CH-Auth-Email: user@domain.suffix  Content-Type: application/json |

> **Notes:**
>
> * The "tag\_id" in the path corresponds to the "id" value in each `tag` object from the **Tag List** array.
> * If your organization is registered on Kentik's EU cluster, `use api.kentik.eu` in place of `api.kentik.com` in the URL above.

**HTTP Response**  
A successful response from the Tag Info method includes the following elements:

* Response headers
* HTTP response code
* A single JSON `tag` object with details about a tag specified by *tag\_id*.

> **Note:** For details on the JSON name/value pairs in a tag object, see **Tag JSON**.

### Tag Create

The Tag Create `POST` method adds a new tag to your Kentik tags.  
  
**HTTP Request**  
The following table shows the path and HTTP request for this cal:

|  |  |
| --- | --- |
| **URL** | https://api.kentik.com/api/v5/tag |
| **Request** | POST /api/v5/tag HTTP/1.1  Host: api.kentik.com  X-CH-Auth-API-Token: user\_api\_token  X-CH-Auth-Email: user@domain.suffix  Content-Type: application/json |

> **Note:** If your organization is registered on Kentik's EU cluster, use `api.kentik.eu` in place of `api.kentik.com` in the URL above.

To create a tag, specify the tag name using the `flow_tag` parameter and at least one additional tag attribute parameter for matching with incoming flow fields. The following parameters can be included in a JSON object in the request body:

| Parameter | Type | Description |
| --- | --- | --- |
| flow\_tag | string | Required: A name for the tag. When a match for any of the other tag parameters is found in the flow from a given device, this string will be added to the src\_flow\_tags and/or dst\_flow\_tags column in that device's KDE main table. |
| device\_name | string | A match results when the device\_name associated with a device sending flow data contains this string. |
| device\_type | string | A comma-delimited list of device types or regular expressions. A match results when any specified device type matches the device\_type associated with a device sending flow data. |
| site | string | A comma-separated list of sites or regular expressions. A match results when any specified site matches the site associated with a device sending flow data. |
| interface\_name | string | A match results when the name or description of an interface sending flow data contains this string. |
| addr | string | A range of IP addresses, expressed in CIDR notation (e.g. 38.12.34.0/24). A match results when this value corresponds to a range of IP addresses in incoming flow.  **Note**: This field can contain up to 249 IP/CIDR items in a comma-delimited list. |
| port | string | A port number, either source (SRC Port) or destination DST Port). A match results when this value appears within a port number in incoming flow. |
| tcp\_flags | string | An integer number between 0 and 255 representing an 8-bit binary bit pattern corresponding to TCP flags. A match will result if the value in both the flow bit pattern and the bitmask is 1 at any of the eight places. |
| protocol | string | The protocol of traffic represented by a flow. The protocol of TCP is 6, and of UDP is 17. A match results when this value is the same as the protocol of the traffic represented by the flow. |
| asn | string | A comma-separated list of ASNs (16- or 32-bit). A match results when any specified ASN is the same as the ASN of the last-hop router based on AS-PATH. |
| lasthop\_as\_name | string | A comma-separated list of AS names or regular expressions. A match results when any specified AS name represents the name corresponding to the last ASN in the path in the routing table for either the source (SRC IP) or destination (DST IP). |
| nexthop\_asn | string | A comma-separated list of ASNs (16- or 32-bit). A match results when any specified ASN is the same as the ASN of the next-hop router based on AS-PATH. |
| nexthop\_as\_name | string | A comma-separated list of AS names or regular expressions. A match results when any specified AS name corresponds to the AS name of the next hop router based on AS path. |
| nexthop | string | A comma-separated list of IPv4 and/or IPv6 CIDRs. If a CIDR grouping (IPv4 or IPv6) is specified, a match can be on any address within that grouping. If no CIDR grouping is specified a match requires an exact IP.   * CIDRs may be expressed in "short form" (e.g. 1::2/127). |
| bgp\_aspath | string | A comma-separated list of numbers or regular expressions representing BGP AS path. A match results when this value is the same as the BGP AS-PATH in the route.   * Permitted characters []\*:\_'$.0123456789()+?,space- * Example: "'3737 1212,\_7801\_,2906$" would look for any of those 3 combinations in the AS path. |
| bgp\_community | string | A comma-separated list of numbers or regular expression representing BGP community (i.e. 2096:2212). A match results when this value is the same as the BGP community of the BGP route associated with incoming flow data.   * Permitted characters []\*:\_'$.0123456789()+?,space- |
| mac | string | A comma-separated list of MAC Addresses. Results in a match if this value matches source or destination Ethernet (L2) address. |
| country | string | A comma-separated list of two-character country codes. Results in a match if this value includes a two-letter country code associated with the source or destination IP of the flow. |
| VLAN(s) | string | A comma-separated list of VLAN IDs. Results in a match if this value includes a VLAN ID associated with the source or destination IP of the flow. |

**HTTP Response**  
A successful response from the Tag Create method includes the following elements:

* Response headers
* HTTP response code
* A JSON `tag` object detailing the newly-added tag, with elements based on specified optional tag fields.

> **Note:** For details on the JSON name/value pairs in a `tag` object, see **Tag JSON**.

### Tag Update

The Tag Update `PUT` method modifies an existing tag, identified by ID, with new values for one or more fields.  
  
**HTTP Request**  
The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| **URL** | https://api.kentik.com/api/v5/tag/tag\_id |
| **Request** | PUT /api/v5/tag/tag\_id HTTP/1.1  Host: api.kentik.com  X-CH-Auth-API-Token: user\_api\_token  X-CH-Auth-Email: user@domain.suffix  Content-Type: application/json |

> **Notes:**
>
> * The "tag\_id" in the path corresponds to the "id" value in each tag object from the **Tag List** array.
> * If your organization is registered on Kentik's EU cluster, `use api.kentik.eu` in place of `api.kentik.com` in the URL above.

Parameters to be updated are included in a JSON tag object with only the fields to be changed. The following example shows an update to the port field:

```
{
"tag": {
     "port": "700"
   }
 }
```

**HTTP Response**  
A successful response from the Tag Update method includes the following elements:

* Response headers
* HTTP response code
* A single JSON `tag` object with details about the newly-updated tag.

> **Note:** For details of the JSON name/value pairs in a `tag` object, see **Tag JSON**.

### Tag Delete

The Tag Delete method removes a specific tag, identified by its ID, from your collection of Kentik tags.  
  
**HTTP Request**  
The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| **URL** | https://api.kentik.com/api/v5/tag/tag\_id |
| **Request** | DELETE /api/v5/tag/tag\_id HTTP/1.1  Host: api.kentik.com  X-CH-Auth-API-Token: user\_api\_token  X-CH-Auth-Email: user@domain.suffix  Content-Type: application/json |

> **Notes:**
>
> * The "tag\_id" in the path corresponds to the "id" value in each tag object from the **Tag List** array.
> * If your organization is registered on Kentik's EU cluster, use api.kentik.eu in place of api.kentik.com in the URL above.

**HTTP Response**  
A successful response from the Tag Delete method includes the following elements:

* Response headers
* HTTP response code 204 (no content)

## Custom Dimension API

> **Note**: Member-level users only have access to the GET methods of this API.

The Custom Dimension API allows for programmatic management of your organization's custom dimensions in Kentik.

> **Notes:**
>
> * For an overview explanation of custom dimensions and their associated populators, see **Dimensions and Populators**.
> * To make calls to this API using cURL, see **API Access Via cURL**.

### Custom Dimension JSON

Calls to the Custom Dimension API return an HTTP response body with a JSON `customDimension` object, or an array of such objects for the Custom Dimension List call. The object includes fields as shown in the following example:

```
{
  “customDimension”: {
    “id”: 9,
    “display_name”: “Mikes int column”,
    “name”: “c_mh_test”,
    “type”: “uint32”,
    “populators”: [
      {
        “id”: 2600,
        “dimension_id”: 9,
        “value”: “1234”,
        “direction”: “src”,
        “device_name”: “My_device”,
        “user”: “6762”,
        “created_date”: “2016-08-24T11:27:04.452Z”,
        “updated_date”: “2016-08-24T11:27:04.452Z”,
        “company_id”: “1289”
      }
    ],
    “created_date”: “2016-08-24T11:25:57.497Z”,
    “updated_date”: “2016-08-27T00:45:22.725Z”,
    “company_id”: “1289”
  }
}
```

Each `customDimension` object contains fields with information about an individual custom dimension registered with Kentik. These fields are described in the following table:

| JSON name | Type | Description |
| --- | --- | --- |
| id | number | The system-assigned ID of the custom dimension. |
| Name | string | The name of the custom dimension (used by system and in SQL):   * Must start with “c\_”. * Valid characters: alphanumeric, underscores. * Length: min=1, max=20. |
| Type | string | The type of the custom dimension.   * Valid values: “string” or “uint32”. |
| Display\_name | string | User-supplied name string for the custom dimension.   * Valid characters: alphanumeric, spaces, dashes, underscores. * Length: min=2, max=30. |
| Populators | array of strings | An array of the populators currently assigned to the custom dimension.  **Note:** For a description of populator fields, see **Populator JSON**. |
| Created\_date | string | Date-time of registration of the custom dimension in Kentik creation in UTC (ISO 8601), e.g., 2015-01-27T01:39:17.186Z |
| updated\_date | string | Date-time of most-recent edit to the custom dimension in UTC, e.g., 2015-01-27T01:39:17.186Z |
| company\_id | number | The system-assigned ID of the customer. |

> **Note:** When a custom dimension is deleted, the name cannot be reused until the longest data retention period of any of your Kentik plans (see **About Plans**) has passed. For example, if your plan retains the Fast dataseries for one year, the deleted dimension’s name cannot be reused for one year.

### Custom Dimension List

The Custom Dimension List `GET` method retrieves a JSON array of your Kentik custom dimensions, where each element represents an individual custom dimension.

> **Note:** The "id" in the `customDimension` object can be used in subsequent calls to retrieve, update, or delete that custom dimension.

**HTTP Request**  
The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| **URL** | https://api.kentik.com/api/v5/customdimensions |
| **Request** | GET /api/v5/customdimensions HTTP/1.1  Host: api.kentik.com  X-CH-Auth-API-Token: user\_api\_token  X-CH-Auth-Email: user@domain.suffix  Content-Type: application/json |

> **Note:** If your organization is registered on Kentik's EU cluster, use `api.kentik.eu` in place of `api.kentik.com` in the URL above.

**HTTP Response**  
A successful response from the Custom Dimension List method includes the following elements:

* Response headers
* HTTP response code
* A response body with a JSON custom dimension array, with a `customDimension` object containing information about each custom dimension.

> **Note:** For details on the JSON name/value pairs in a `customDimension` object, see **Custom Dimension JSON**.

### Custom Dimension Info

The Custom Dimension Info GET method fetches details of a specific custom dimension by its ID from your list of Kentik custom dimensions.  
  
**HTTP Request**  
The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| **URL** | https://api.kentik.com/api/v5/customdimension/dimension\_id |
| **Request** | GET /api/v5/customdimension/dimension\_id HTTP/1.1  Host: api.kentik.com  X-CH-Auth-API-Token: user\_api\_token  X-CH-Auth-Email: user@domain.suffix  Content-Type: application/json |

> **Notes:**
>
> * The `dimension_id` in the path corresponds to the id value in the `customDimension` object from the **Custom Dimension List** array.
> * If your organization is registered on Kentik's EU cluster, use `api.kentik.eu` in place of `api.kentik.com` in the URL above.

**HTTP Response**  
A successful response from the Custom Dimension Info method includes the following elements:

* Response headers
* HTTP response code
* A single JSON customDimension object containing information about the custom dimension specified by `dimension_id`.

> **Note:** For details of the JSON name/value pairs in a custom dimension object, see **Custom Dimension JSON**.

### Custom Dimension Create

The Custom Dimension Create POST method adds a new custom dimension to your collection of custom dimensions.  
  
**HTTP Request**  
The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| **URL** | https://api.kentik.com/api/v5/customdimension |
| **Request** | POST /api/v5/customdimension HTTP/1.1  Host: api.kentik.com  X-CH-Auth-API-Token: user\_api\_token  X-CH-Auth-Email: user@domain.suffix  Content-Type: application/json |

> **Note:** If your organization is registered on Kentik's EU cluster, use `api.kentik.eu` in place of `api.kentik.com` in the URL above.

The following parameters are passed in a JSON object in the request body:

| JSON name | Type | Description |
| --- | --- | --- |
| name | string | Required: The name of the custom dimension (used by system and in SQL).  - Must start with "c\_".  - Valid characters: alphanumeric, dashes, underscores.  - Length: min=1, max=20. |
| type | string | Required: The type of the custom dimension.  - Valid values: "string" or "uint32". |
| display\_name | string | Required: A name string for the custom dimension (used in portal for group-by dimensions and filters).  - Valid characters: alphanumeric, spaces, dashes, underscores.  - Length: min=2, max=30. |

> **Note:** When a custom dimension is deleted, its name cannot be reused until the longest data retention period of any of your Kentik plans has passed. For example, if your plan includes retention of the Fast dataseries for one year, then that dimension's name cannot be reused for one year.

**HTTP Response**  
A successful response from the Custom Dimension Create method includes the following elements:

* Response headers
* HTTP response code
* A single JSON `customDimension` object containing information about the newly-added custom dimension.

> **Note:** For details of the JSON name/value pairs in a customDimension object, see **Custom Dimension JSON**.

### Custom Dimension Update

The Custom Dimension Update PUT method modifies a specific custom dimension, identified by its ID, within your collection of custom dimensions. The only attribute that can be updated is the display name.  
  
**HTTP Request**  
The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| **URL** | https://api.kentik.com/api/v5/customdimension/dimension\_id |
| **Request** | PUT /api/v5/customdimension/dimension\_id HTTP/1.1  Host: api.kentik.com  X-CH-Auth-API-Token: user\_api\_token  X-CH-Auth-Email: user@domain.suffix  Content-Type: application/json |

> **Notes:**
>
> * The "dimension\_id" in the path corresponds to the "id" in the customDimension object from the **Custom Dimension List** array.
> * If your organization is registered on Kentik's EU cluster, use api.kentik.eu in place of api.kentik.com in the URL above.

The parameters to be updated are included in a JSON customDimension object with only the fields to be changed. The following example shows an update to the custom dimension’s display name:

```
{
"custom dimension": {
     "display_name": "The 4th dimension"
   }
 }
```

> **Note:** Only the display name can be changed when updating a custom dimension.

**HTTP Response**  
A successful response from the Custom Dimension Update method includes the following elements:

* Response headers
* HTTP response code
* A JSON `customDimension` object with details of the updated custom dimension, excluding the `populators` array.

> **Note:** For details on the JSON name/value pairs in a `customDimension` object, see **Custom Dimension JSON**

### Custom Dimension Delete

The Custom Dimension Delete method removes a specific custom dimension, identified by its ID, from a customer's collection.

> **Note:** When a custom dimension is deleted, its name cannot be reused until the longest data retention period of any of your Kentik plans has passed. For example, if your plan retains the Fast dataseries for one year, the deleted dimension's name cannot be reused for one year.

**HTTP Request**  
The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| **URL** | https://api.kentik.com/api/v5/customdimension/dimension\_id |
| **Request** | DELETE /api/v5/customdimension/dimension\_id HTTP/1.1  Host: api.kentik.com  X-CH-Auth-API-Token: user\_api\_token  X-CH-Auth-Email: user@domain.suffix  Content-Type: application/json |

> **Notes:**
>
> * The "custom dimension\_id" in the path corresponds to the "id" in the customDimension object from the **Custom Dimension List** array.
> * If your organization is registered on Kentik's EU cluster, use api.kentik.eu in place of api.kentik.com in the URL above.

**HTTP Response**  
A successful response from the Custom Dimension Delete method includes the following elements:

* Response headers
* HTTP response code 204 (no content)

### Populator JSON

The Custom Dimensions API uses `populator` objects to decide when the KDE main table field (see **KDE Tables**), corresponding to a custom dimension, is populated for a specific row (see **Dimensions and Populators**). These populators are located in the `populators` array of the `customDimension` object, which is returned in an HTTP response from a **Custom Dimension List** or **Custom Dimension Info** call. A `populator` object consists of the fields shown in the following example:

```
{
"populator": {
     "id": 2600,
     "dimension_id": 9,
     "value": "1234",
     "direction": "dst",
     "device_name": "my_device1",
     "interface_name": "a great interface",
     "addr": "192.168.2.1/32",
     "port": "800",
     "tcp_flags": "254",
     "protocol": "7",
     "asn": "1235",
     "nexthop_asn": "5437",
     "bgp_aspath": "1212",
     "bgp_community": "2096:2212",
     "user": "2345",
     "created_date": "2016-08-24T11:27:04.452Z",
     "updated_date": "2016-10-12T22:37:19.251Z",
     "company_id": "1289"
   }
 }
```

> **Note:** The returned `populator` object may not include all fields, as it depends on which fields are defined in the object.

The fields of each populator object provide information about an individual populator assigned to a specific custom dimension. These fields are described in the following table:

| JSON name | Type | Description | Portal field |
| --- | --- | --- | --- |
| id | number | The system-assigned ID of the populator. | Populator ID |
| dimension\_id | number | Required: the ID of the custom dimension to which the populator is assigned. | N.A. |
| value | string | Required: The value placed in the custom column when there’s a match with the ANDed populator parameters below.   * When custom dimension type is "string": * Valid characters: alphanumeric, spaces, dashes, underscores. * Length: min=1, max=128. * When the custom dimension's type is "uint32" * Valid values: min=0, max=4294967295. | Dimension value |
| direction | string | The direction (source, destination, or either) of the flow fields in which to look for a match with the populator parameters below. | Direction |
| device\_name | string | A match results when the device\_name associated with a device sending flow data contains this string. | Device name |
| device\_type | string | A comma-delimited list of device types or regular expressions. A match results when any specified device type matches the device\_type associated with a device sending flow data. | Device type |
| site | string | A comma-separated list of sites or regular expressions. A match results when any specified site matches the site associated with a device sending flow data. | Site |
| interface\_name | string | A match results when the name or description of an interface sending flow data contains this string. | Interface name/description |
| addr | string | A range of IP addresses, expressed in CIDR notation (e.g., 38.12.34.0/24). A match results when this value corresponds to a range of IP addresses in incoming flow. | IP address |
| port | string | A port number, either source (SRC Port) or destination DST Port). A match results when this value appears within a port number in incoming flow. | Port |
| tcp\_flags | string | An integer number between 0 and 255 representing an 8-bit binary bit pattern corresponding to TCP flags. A match will result if the value in both the flow bit pattern and the bitmask is 1 at any of the eight places. | TCP flag |
| protocol | string | The protocol of traffic represented by a flow. The protocol of TCP is 6, and of UDP is 17. A match results when this value is the same as the protocol of the traffic represented by the flow. | Protocol name/number |
| asn | string | A comma-separated list of ASNs (16- or 32-bit). A match results when any specified ASN is the same as the ASN of the last-hop router based on AS-PATH. | Last-hop (origin) ASN |
| lasthop\_as\_name | string | A comma-separated list of AS names or regular expressions. A match results when any specified AS name represents the name corresponding to the last ASN in the path in the routing table for either the source (SRC IP) or destination (DST IP). | Last-hop (origin) AS Name |
| nexthop\_asn | string | A comma-separated list of ASNs (16- or 32-bit). A match results when any specified ASN is the same as the ASN of the next-hop router based on AS-PATH. | Next-hop ASN |
| nexthop\_as\_name | string | A comma-separated list of AS names or regular expressions. A match results when any specified AS name corresponds to the AS name of the next hop router based on AS path. | Next-hop AS Name |
| nexthop | string | A comma-separated list of IPv4 and/or IPv6 CIDRs. If a CIDR grouping (IPv4 or IPv6) is specified, a match can be on any address within that grouping. If no CIDR grouping is specified a match requires an exact IP.   * CIDRs may be expressed in "short form" (e.g., 1::2/127). | Next-hop IP |
| bgp\_aspath | string | A comma-separated list of numbers or regular expressions representing BGP AS path. A match results when this value is the same as the BGP AS-PATH in the route.   * Permitted characters []\*:\_'$.0123456789()+?,space-  - Example: "'3737 1212,\_7801\_,2906$" would look for any of those 3 combinations in the AS path. | BGP AS path |
| bgp\_community | string | A comma-separated list of numbers or regular expression representing BGP community (i.e. 2096:2212). A match results when this value is the same as the BGP community of the BGP route associated with incoming flow data.   * Permitted characters []\*:\_'$.0123456789()+?,space- | BGP community |
| mac | string | A comma-separated list of MAC Addresses. Results in a match if this value matches source or destination Ethernet (L2) address. | MAC Address |
| country | string | A comma-separated list of two-character country codes. Results in a match if this value includes a two-letter country code associated with the source or destination IP of the flow. | Country |
| vlans | string | A comma-separated list of VLAN IDs. Results in a match if this value includes a VLAN ID associated with the source or destination IP of the flow. | VLAN(s) |
| user | number | The system-assigned ID of the user who created the populator. | N.A. |
| created\_date | string | Date-time of populator creation, in UTC (ISO 8601), e.g., 2025-01-27T01:39:17.186Z | N.A. |
| updated\_date | string | Date-time of most-recent populator edit, in UTC, e.g., 2025-01-27T01:39:17.186Z | N.A. |
| company\_id | number | The system-assigned ID of the customer. | N.A. |

> **Notes:**
>
> * The Portal field column shows the corresponding setting name in the **Populator Settings Dialog** in the portal.
> * Commas are used only as delimiters in comma-delimited lists.
> * For more on field value validations, see **Populator Field Validation**.

### Populator Create

The Populator Create POST method adds a new populator to your Kentik populators.

**HTTP Request**  
The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| **URL** | api.kentik.com/api/v5/customdimension/dimension\_id/populator |
| **Request** | POST /api/v5/customdimension/dimension\_id/populator HTTP/1.1  Host: api.kentik.com  X-CH-Auth-API-Token: user\_api\_token  X-CH-Auth-Email: user@domain.suffix  Content-Type: application/json |

> **Notes:**
>
> * The `dimension_id` in the path corresponds to the id value in the `customDimension` object from the **Custom Dimension List** array.
> * If your organization is registered on Kentik's EU cluster, use `api.kentik.eu` in place of `api.kentik.com` in the URL above.

The following parameters are passed in a JSON object in the request body. At least one of the non-required parameters must be specified for the populator to work:

| JSON name | Type | Description |
| --- | --- | --- |
| dimension\_id | number | Required: the ID of the custom dimension to which the populator is assigned. |
| value | string | Required: The value placed in the custom column when there’s a match with the ANDed populator parameters below.   * When custom dimension type is "string": * Valid characters: alphanumeric, spaces, dashes, underscores. * Length: min=1, max=128. * When the custom dimension's type is "uint32" * Valid values: min=0, max=4294967295. |
| direction | string | Required: The direction (src, dst, or either) of the flow fields in which to look for a match with the populator parameters below. |
| device\_name | string | A match results when the device\_name associated with a device sending flow data contains this string. |
| interface\_name | string | A match results when the name or description of an interface sending flow data contains this string. |
| addr | string | A range of IP addresses, expressed in CIDR notation (e.g. 38.12.34.0/24). A match results when this value corresponds to a range of IP addresses in incoming flow. |
| port | string | A port number, either source (SRC Port) or destination DST Port). A match results when this value appears within a port number in incoming flow. |
| tcp\_flags | string | An integer number between 0 and 255 representing an 8-bit binary bit pattern corresponding to TCP flags. A match will result if the value in both the flow bit pattern and the bitmask is 1 at any of the eight places. |
| protocol | string | The protocol of traffic represented by a flow. The protocol of TCP is 6, and of UDP is 17. A match results when this value is the same as the protocol of the traffic represented by the flow. |
| asn | string | A comma-separated list of ASNs (16- or 32-bit). A match results when any specified ASN is the same as the ASN of the last-hop router based on AS-PATH. Both 16- and 32-bit ASNs are valid. |
| lasthop\_as\_name | string | A comma-separated list of AS names or regular expressions. A match results when any specified AS name represents the name corresponding to the last ASN in the path in the routing table for either the source (SRC IP) or destination (DST IP). |
| nexthop\_asn | string | A comma-separated list of ASNs (16- or 32-bit). A match results when any specified ASN is the same as the ASN of the next-hop router based on AS-PATH. Both 16- and 32-bit ASNs are valid. |
| nexthop\_as\_name | string | A comma-separated list of AS names or regular expressions. A match results when any specified AS name corresponds to the AS name of the next hop router based on AS path. |
| nexthop | string | A comma-separated list of IPv4 and/or IPv6 CIDRs. If a CIDR grouping (IPv4 or IPv6) is specified, a match can be on any address within that grouping. If no CIDR grouping is specified a match requires an exact IP.   * CIDRs may be expressed in "short form" (e.g. 1::2/127). |
| bgp\_aspath | string | A comma-separated list of numbers or regular expressions representing BGP AS path. A match results when this value is the same as the BGP AS-PATH in the route.   * Permitted characters []\*:\_'$.0123456789()+?,space- * Example: "'3737 1212,\_7801\_,2906$" would look for any of those 3 combinations in the AS path. |
| bgp\_community | string | A comma-separated list of numbers or regular expression representing BGP community (i.e. 2096:2212). A match results when this value is the same as the BGP community of the BGP route associated with incoming flow data.   * Permitted characters []\*:\_'$.0123456789()+?,space- |
| mac | string | A comma-separated list of MAC Addresses. Results in a match if this value matches source or destination Ethernet (L2) address. |
| country | string | A comma-separated list of two-character country codes. Results in a match if this value includes a two-letter country code associated with the source or destination IP of the flow. |
| vlans | string | A comma-separated list of VLAN IDs. Results in a match if this value includes a VLAN ID associated with the source or destination IP of the flow. |

**HTTP Response**  
A successful response from the Populator Create method includes the following elements:

* Response headers
* HTTP response code
* A JSON `populator` object with details of the newly added populator’s parameters.

> **Note:** For details on the JSON name/value pairs in a `populator` object, see **Populator JSON.**

### Populator Update

The Populator Update `PUT` method updates data for a specific populator, identified by its ID, in your list of existing Kentik populators.

> **Note:** Use the **Custom Dimension Info** call to find the ID of the populator you want to update.

**HTTP Request**  
The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| **URL** | api.kentik.com/api/v5/customdimension/dimension\_id/populator/populator\_id |
| **Request** | PUT /api/v5/customdimension/dimension\_id/populator/populator\_id HTTP/1.1  Host: api.kentik.com  X-CH-Auth-API-Token: user\_api\_token  X-CH-Auth-Email: user@domain.suffix  Content-Type: application/json |

> **Notes:**
>
> * The `dimension_id` in the path corresponds to the `id` value in the `customDimension` object from the **Custom Dimension List** array.
> * If your organization is registered on Kentik's EU cluster, use `api.kentik.eu in` place of `api.kentik.com` in the URL above.

The parameters to be updated are included in a JSON `populator` object with only the fields to be changed. The following example shows an update to the port number:

```
{
"populator": {
     "port": "700"
   }
 }
```

**HTTP Response**  
A successful response from the Populator Update method includes the following elements:

* Response headers
* HTTP response code
* A JSON `populator` object with the updated populator’s parameters.

> **Note:** For details of the JSON name/value pairs in a `populator` object, see **Populator JSON**.

### Populator Delete

This Populator Delete method removes one populator, identified by its ID, from your collection of Kentik populators.

> **Notes:**
>
> * Deleting a populator for a custom dimension makes previously matched KDE main table rows unfilterable by that populator's value.
> * Use the **Custom Dimension Info** call to identify the ID of the populator you wish to delete.

**HTTP Request**  
The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| **URL** | api.kentik.com/api/v5/customdimension/dimension\_id/populator/populator\_id |
| **Request** | DELETE /api/v5/populator/populator\_id HTTP/1.1  Host: api.kentik.com  X-CH-Auth-API-Token: user\_api\_token  X-CH-Auth-Email: user@domain.suffix  Content-Type: application/json |

> **Notes:**
>
> * The `dimension_id` in the path corresponds to the id value in the `customDimension` object from the **Custom Dimension List** array.
> * If your organization is registered on Kentik's EU cluster, use `api.kentik.eu` in place of `api.kentik.com` in the URL above.

**HTTP Response**  
A successful response from the Populator Delete method includes the following elements:

* Response headers
* HTTP response code 204 (no content)

## Batch API

> **Note:** Member-level users do not have access to this API.

The Batch API is covered here.

> ***Notes:***
>
> * A Python integration is available for making calls to this API. See our **Hypertagging API project** on GitHub.
> * To make calls to this API using cURL, see **API Access Via cURL**.
> * For information on adding and managing tags via the Kentik portal, see **Flow Tags**.
> * For information on adding and managing custom dimensions via the portal, see **Dimensions and Populators**.

### About the Batch API

The Batch API allows programmatic management of your organization’s flow tags (see **About Flow Tags**) and custom dimension populators (see **Dimensions and Populators**) in Kentik. It supports:

* Adding a batch of tags or populators
* Modifying existing tags or populators
* Removing existing tags or populators
* Retrieving status updates on batch operation progress

> **Note:** The Batch API complements, but does not replace, our APIs for managing flow tags and custom dimension populators. See **Tag API** and **Custom Dimension API**.

### Using the Batch API

The Batch API differs from other Kentik V5 APIs by using a single `POST` method (see **Batch Request**) for creating, updating, and deleting objects, with actions determined by parameter values. It also provides a separate `GET` method (see **Batch Status Request**) to retrieve status information.  
  
 For the create/update/delete method, multiple requests, known as “parts”, can be included in a single batch operation, with each part having a 4MB payload limit. The multi-part process involves:

* **Step 1**: POST the request JSON for the first part to the appropriate endpoint (tags or populators). Set the `complete` parameter to `false` to indicate more parts will follow.
* **Step 2**:An HTTP response with status 202 (Accepted) confirms receipt, queueing the request for offline processing. The response includes a GUID (Globally Unique Identifier), which is used to:

  + Identify subsequent requests as part of the same batch operation.
  + Request the batch operation status in a **Batch Status Request**.
* **Step 3**:POST the request JSON for additional parts, using the same GUID. In the final request, set the `complete` parameter (see **Batch Request JSON Fields**) to `true`.

> **Note:** The creation or updating of flow tags or custom dimension populators starts with the first request, but deletion of unneeded tags or populators occurs only after all batch operation parts are received.

### Batch Request

The Batch Request `POST` method is used to manage either:

* All your flow tags
* The populators for a single custom dimension

#### Criteria and Value

Managing tags or populators with the Batch Request method requires specifying the following:

* **Criteria**: An array of criteria to evaluate each flow record for a match at ingest, including:

  + Direction: Specifies the direction (source, destination, or either) of fields like SRC IP or SRC port
  + Field: The field name and value to match, e.g., `"addr": ["1.2.3.4"]` for matching the IP address "1.2.3.4.".
* **Value**:

  + Tags: The tag name to be entered into the `src_flow_tags or dst_flow_tags` column if criteria match.
  + Populators: The text to be inserted into the custom dimension field if criteria match (see **Dimensions and Populators**).

The JSON allows multiple `criteria` sets for each `value`, with each unique criteria-value combination defining a tag or populator, depending on the endpoint.

> **Note:** A single set of `criteria` can define multiple values, leading to multiple tags/populators matching the same criteria:
>
> * Tags: The flow record field `(src_flow_tags or dst_flow_tags)` will include all matching tag names.
> * Populators: A random value from the matching criteria will be selected for the custom dimension field.

#### Upserts and Deletes

The Batch API uses the `criteria` and `value` parameters based on context to determine the action taken on the flow tag:

* **Upserts**: An array for update or insert operations:

  + Update: Modify the value if the criterion is already specified for the tag.
  + Insert: Add the criterion and set it to the provided value if not already specified for the tag.

    > **Note:** Users don’t need to know if a criterion is already specified.

**Deletes**: An array for delete operations where each `value` parameter leads to the deletion of the corresponding tag or populator.

> **Notes:**
>
> * Deletions do not require a `criteria` parameter.
> * The API allows updates without needing the ID of each tag or populator, which may sometimes change the ID during an update.

#### Batch Request Paths

The batch request path varies based on whether you’re managing flow tags or populators:

* **Flow Tags**: The following table shows the path and HTTP request for a tag-related request:

|  |  |
| --- | --- |
| **URL** | https://api.kentik.com/api/v5/batch/tags |
| **Request** | POST /api/v5/batch/tags HTTP/1.1  Host: api.kentik.com  X-CH-Auth-API-Token: user\_api\_token  X-CH-Auth-Email: user@domain.suffix  Content-Type: application/json |

* **Populators**: The following table shows the path and HTTP request for a populator-related request:

|  |  |
| --- | --- |
| **URL** | https://api.kentik.com/api/v5/batch/customdimensions/dimension\_name/populators |
| **Request** | POST /api/v5/batch/customdimensions/dimension\_name/populators HTTP/1.1  Host: api.kentik.com  X-CH-Auth-API-Token: user\_api\_token  X-CH-Auth-Email: user@domain.suffix  Content-Type: application/json |

> **Note:** If your organization is registered on Kentik's EU cluster, use `api.kentik.eu` in place of `api.kentik.com` in the URLs above.

#### Sample Batch Request JSON

The following example JSON is used in a single-part request to the Custom Dimension endpoint to:

* Create or update three populators with the value "service\_abc", based on specified direction and IP address `criteria`.
* Delete two populators as specified in the `deletes` array.

```
{
"replace_all": false,
"complete": true,
"upserts": [
    {
      "value": "service_abc",
      "criteria": [
        {
          "direction": "dst",
          "addr": ["91.221.37.160"]
         },
         {
           "direction": "src",
           "addr": ["94.74.74.192"]
        },
        {
          "direction": "dst",
          "addr": ["1.2.3.4"]
         }
       ]
     }
   ],
  "deletes": [
    {
      "value": "324234"
    },
    {
      "value": "test4"
    }
  ]
 }
```

#### Batch Request JSON Fields

Parameters in a JSON request body can be categorized as follows:

* **Batch-related parameters**: Settings and structures for batch operations. See table below.
* **Match fields**: Included in the `criteria` array, either:

  + Tag-related parameters: See the table in **Tag Create**.
  + Populator-related parameters: See the table in **Populator Create**.

The following batch-related parameters may be passed in the request body:

| Parameter | Type | Description |
| --- | --- | --- |
| replace\_all | Boolean | Required:   * If false, all tags or populators that are not included in the upserts array will be left unchanged. * If true, all tags or populators that are not included in the upserts array will be deleted. |
| complete | Boolean | Required:   * If false, this is a multi-part operation and this is not the last part (additional parts will be sent in subsequent requests). * If true, this is the last part of the operation (which may also be the only part). |
| upserts | array | An array of objects that each contain a value (the name of one or more tags or populators) and a criteria array in which each element represents a tag or populator (depending on endpoint). |
| upserts.value | string | Required if there are upserts:   * For tags, the name of the tag. * For populators, the populator's "value" property. |
| upserts.criteria | string | An array of objects that each contain the information needed to update or insert match criteria for one tag or populator. |
| upserts.criteria.direction | string | The direction (src, dst, or either) of the direction-specific fields in the flow record (e.g. SRC IP, SRC port, etc.) that will be evaluated for a match.  **Note:** The default for populators is “either.” For tags this parameter is ignored. |
| deletes | array | An array of objects that each contain the information needed to identify items for deletion. All items (tags or populators, depending on endpoint) matching one of the values specified in the array will be deleted. |
| deletes.value | string | Required if there are deletes:   * For tags, the name of the tag. * For populators, the populator's "value" property. |

  

### Batch Response

A successful response from the Batch Request method includes the following elements:

* Response headers
* HTTP response code 202 (accepted)
* A JSON structure with status information for the operation identified in the **Batch Status Request**

#### Batch Response JSON

The following shows the JSON returned in the batch response:

```
{
"message": "Successfully stored request. Batch queued for processing.",
"guid": "batch_operation_guid"
}
```

The GUID is used to:

* Identify subsequent batches as part of the same operation.
* Request the status of a batch operation via a **Batch Status Request**.

### Batch Status Request

The Batch Status Request `GET` method allows you to check the status of a batch operation initiated with a **Batch Request**. The operation is identified by the GUID from the original batch request and is included in the status request URL as shown in **Batch Status HTTP Request**.

#### Batch Status HTTP Request

The HTTP request for a batch status request is as shown in the following table:

|  |  |
| --- | --- |
| **URL** | https://api.kentik.com/api/v5/batch/batch\_operation\_guid/status |
| **Request** | GET /api/v5/batch/batch\_operation\_guid/status HTTP/1.1  Host: api.kentik.com  X-CH-Auth-API-Token: user\_api\_token  X-CH-Auth-Email: user@domain.suffix  Content-Type: application/json |

> **Note:** If your organization is registered on Kentik's EU cluster, use `api.kentik.eu` in place of `api.kentik.com` in the URL above.

### Batch Status Response

A successful response from the Batch Status Request method includes the following elements:

* Response headers
* HTTP response code 200 (OK)
* A JSON structure with the operation’s status information

#### Sample Batch Status Response JSON

Each call to the Batch Status endpoint returns a JSON object with the fields (name/value pairs) as shown in the following example:

```
{
"custom_dimension": {
     "id": 98,
     "name": "MY_FAVE_DIM"
   },
"guid": "kentik_assigned_batch_id",
"is_multipart": false,
"is_pending": false,
"is_complete": true,
"number_of_parts": 1,
"user": {
     "id": 2233,
     "email": "username@domain.com"
   },
"upserts": {
     "total": 9955,
    "applied": 9898,
    "invalid": 0,
     "unchanged": 0,
     "over_limit": 57
   },
"deletes": {
     "total": 0,
    "applied": 0,
     "unchanged": 0,
     "invalid": 0
   },
"replace_all": {
     "requested": true,
     "deletes_performed": 0,
     "successful": true
   },
"batch_date": "2018-09-25T21:41:18.88816Z"
}
```

> **Note:** The `custom_dimension` object is excluded when the batch operation status requested pertains to tags instead of custom dimension populators.

#### Batch Status Response JSON Fields

The table below describes the elements of the JSON returned with the status response.

| JSON name | Type | Description |
| --- | --- | --- |
| custom\_dimension | object | An object containing ID and name information identifying the custom dimension whose populators are being managed by this batch operation.  **Note:** Included only when the batch operation is for a custom dimension. |
| custom\_dimension.id | number | The unique ID of the custom dimension (Kentik assigned). |
| custom\_dimension.name | string | The name specified when the custom dimension was created (see **Dimension General Properties**). |
| Guid | string | A unique identifier for this batch operation. Used for additional batch requests (if the batch operation is multipart) and batch status requests related to the operation. |
| Is\_multipart | Boolean | Indicates whether this batch operation is made up of multiple batches. |
| Is\_complete | Boolean | * If true, all batches in the batch operation have been received and processed. * If false, remaining batches in the operation have not yet been received or processed. |
| Number\_of\_parts | number | The total number of batches received so far in the batch operation. |
| User | object | Container for information about the user who initiated the batch operation. |
| User.id | number | The unique Kentik ID of the user. |
| User.email | string | The email address associated with the user. |
| Upserts | object | Container for information about upserts processed so far in this batch operation. |
| Upserts.total | number | The sum of the remaining four parameters in the upserts object. |
| Upserts.applied | number | The number of values so far that have been applied. |
| Upserts.invalid | number | The number of values so far that were determined to be invalid. |
| Upserts.unchanged | number | The number of values so far that resulted in no change to an existing value. |
| Upserts.over\_limit | number | The number of values so far that were ignored because updating or inserting them would result in excess populators (over contractual limits). |
| Deletes | object | Container for information about deletes processed so far in this batch operation. |
| Deletes.total | number | The sum of the remaining four parameters in the deletes object. |
| Deletes.applied | number | The number of values so far that have been deleted. |
| Deletes.unchanged | number | The number of values so far that resulted in no deletion. |
| Deletes.invalid | number | The number of values so far for which the deletion request was determined to be invalid. |
| Replace\_all | object | Container for information about changes made so far in this batch operation due to the replace\_all setting in the batch operation request. |
| Replace\_all.requested | Boolean | Indicates whether replace\_all was or was not requested. |
| Replace\_all.deletes\_performed | number | The number of values deleted because they were not included in the request. |
| Replace\_all.successful | Boolean | True unless an issue has been encountered that would prevent deletion of any values that would otherwise be deleted. |
| Batch\_date | string | The date-time at which the batch request with this GUID was initially received by the server. |

  

## Saved Filter API

The Saved Filter API allows programmatic management of custom filters saved in Kentik by your organization’s users (see **Company Saved Filters**).

> **Notes:**
>
> * Member-level users only have access to the GET methods of this API.
> * For information on adding and managing saved filters via the Kentik portal, see **Saved Filters**.
> * To make calls to this API using cURL, see **API Access Via cURL**.

### Saved Filter JSON

Calls to the Saved Filter API return a JSON “saved filter” object, or an array of such objects for the Saved Filter List call. For example, this Saved Filter List call returns an array with two objects:

* The first object (ID 001) was created via the API
* The second object (ID 002) was saved from the Kentik portal and includes additional internal fields not found in API-created filters.

```
[
   {
    “cdate”: “2016-10-26T06:00:08.574Z”,
    “company_id”: “1223”,
    “edate”: “2016-10-26T06:00:08.574Z”,
    “filter_description”: “zzz”,
    “filter_level”: “company”,
    “filter_name”: “zzz”,
    “filters”: {
      “connector”: “any”,
      “custom”: true,
      “filterGroups”: [
        {
          “connector”: “All”,
          “filters”: [
            {
              “filterField”: “dst_as”,
              “filterValue”: “17”,
              “operator”: “=”
            },
            {
              “filterField”: “src_geo”,
              “filterValue”: “us”,
              “operator”: “=”
            }
          ],
           “not”: true
         }
       ]
     },
     “id”: 001
   },
   {
     “cdate”: “2016-12-09T21:03:12.738Z”,
     “company_id”: “1223”,
     “edate”: “2016-12-09T21:03:12.738Z”,
     “filter_description”: “Source IP address is within RFC1918 address space:  10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16”,
     “filter_level”: “company”,
     “filter_name”: “RFC1918_SRC Copy 0.7142274979608774”,
     “filters”: {
       “connector”: “All”,
       “custom”: false,
       “filterGroups”: [
        {
          “connector”: “Any”,
          “filterString”: “(  (  (inet_src_addr ILIKE ‘10.0.0.0/8’) )  OR  (  (inet_src_addr ILIKE ‘172.16.0.0/12’) )  OR  (  (inet_src_addr ILIKE ‘192.168.0.0/16’) )) “,
          “filters”: [
            {
              “filterField”: “inet_src_addr”,
              “filterValue”: “10.0.0.0/8”,
              “id”: “c115”,
              “operator”: “ILIKE”
            },
            {
              “filterField”: “inet_src_addr”,
              “filterValue”: “192.168.0.0/16”,
              “id”: “c117”,
              “operator”: “ILIKE”
            }
          ],
           “id”: “c118”,
           “metric”: null,
           “not”: false
         }
       ],
       “filterString”: “((  (  (inet_src_addr ILIKE ‘10.0.0.0/8’) )  OR  (  (inet_src_addr ILIKE ‘172.16.0.0/12’) )  OR  (  (inet_src_addr ILIKE ‘192.168.0.0/16’) )))”
     },
     “id”: 002
   }
]
```

> **Note:** The returned JSON may not include all fields shown above, depending on the defined filter groups and individual filters.

Each saved filter object provides details on a custom filter saved by a user in your organization. The components of the object are described in the following topics.

#### Saved Filter Object

The highest level in the hierarchy of a saved filter, containing the following elements:

| JSON name | Type | Description |
| --- | --- | --- |
| cdate | string | The system-assigned date-time of filter creation, in UTC (ISO 8601), e.g., 2015-01-27T01:39:17.186Z |
| company\_id | number | The system-assigned ID of the customer. |
| Edate | string | The system-assigned date-time of most-recent modification, in UTC (ISO 8601), e.g., 2015-01-27T01:39:17.186Z |
| filter\_description | string | An optional description of the saved filter. |
| Filter\_level | string | The scope across which the shared filter is available: user (personal only), company (organization-wide), or global (Kentik-provided preset).  ***Note:***User and global saved filters are not yet implemented in the API. |
| Filter\_name | string | A unique name for the saved filter. |
| Filters | object | An object representing a set of filter groups that make up a saved filter.  See **Filters Object**. |
| Id | number | The system-assigned ID of the saved filter. |

#### Filters Object

A set of one or more filter groups, each made up of the name/value pairs shown in the following table:

| JSON name | Type | Description |
| --- | --- | --- |
| connector | string | Sets the conjunctive operator (“and” or “or”) that will join all filter groups in this filter set. |
| Custom | Boolean | Reserved for internal use. |
| filterGroups | array | An array of filter groups that make up this saved filter. See **FilterGroups Array**. |
| filterString | string | Reserved for internal use.  **Note:** Present only if the filter was saved from the Kentik portal. |

#### FilterGroups Array

An array of filter groups, each made up of the name/value pairs shown in the following table:

| JSON name | Type | Description |
| --- | --- | --- |
| connector | string | Sets the conjunctive operator (“and” or “or”) that will join all filters in this filter group. |
| filterString | string | Reserved for internal use.  **Note:** Present only if the filter was saved from the Kentik portal. |
| Filters | array | An array of filters that make up this filter group. See **Filters Array**. |
| Id | string | The system-assigned unique ID identifying an individual filter group.  **Note:** Present only if the filter was saved from the Kentik portal. |
| Metric | string | Reserved for internal use. |
| Not | Boolean | Sets whether traffic matching the filter group is included (false) or excluded (true). |

#### Filters Array

An array of individual filters, each made up of the name/value pairs shown in the following table:

| JSON name | Type | Description |
| --- | --- | --- |
| filterField | string | The dimension to filter on.   * Valid values: src\_geo, src\_geo\_region, src\_geo\_city, src\_as, src\_flow\_tags, l4\_src\_port, vlan\_in, src\_eth\_mac, inet\_src\_addr, input\_port, i\_input\_snmp\_alias, i\_input\_interface\_description, ipv4\_src\_route\_prefix, src\_route\_length, src\_bgp\_aspath, src\_bgp\_community, ipv4\_src\_next\_hop, src\_nexthop\_as, src\_second\_asn, src\_third\_asn, dst\_geo, dst\_geo\_region, dst\_geo\_city, dst\_as, dst\_flow\_tags, l4\_dst\_port, vlan\_out, dst\_eth\_mac, inet\_dst\_addr, output\_port, i\_output\_snmp\_alias, i\_output\_interface\_description, ipv4\_dst\_route\_prefix, dst\_route\_length, dst\_bgp\_aspath, dst\_bgp\_community, ipv4\_dst\_next\_hop, dst\_nexthop\_as, dst\_second\_asn, dst\_third\_asn, tcp\_flags, tcp\_flags\_raw, protocol, i\_device\_name, both\_pkts, tcp\_retransmit, or tos |
| id | string | The system-assigned unique ID identifying an individual filter.  **Note:** Present only if the filter was saved from the Kentik portal. |
| filterValue | string | The filterField value to match. |
| Operator | string | The operator to apply in the filter.  - Valid value: “=”, “<>”, “ILIKE”, “NOT ILIKE”, “˜”, “!˜”, “>”, “<”, “&” |

> **Note:** For more details on how the name/value pairs correspond to custom filters saved via the Kentik portal, see **Filter Groups UI**.

### Saved Filter List

The Saved Filter List GET method fetches JSON array of custom filters saved by users in your organization, with each element representing an individual custom filter.

> **Note:** The "id" for each individual custom filter can be used in subsequent calls to retrieve, update, or delete that custom filter.

**HTTP Request**  
The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| **URL** | https://api.kentik.com/api/v5/saved-filters/custom |
| **Request** | GET /api/v5/saved-filters/custom HTTP/1.1  Host: api.kentik.com  X-CH-Auth-API-Token: user\_api\_token  X-CH-Auth-Email: user@domain.suffix  Content-Type: application/json |

> **Note:** If your organization is registered on Kentik's EU cluster, use `api.kentik.eu` in place of `api.kentik.com` in the URL above.

**HTTP Response**  
A successful response from the Saved Filter List method includes the following elements:

* Response headers
* HTTP response code
* A response body with a JSON array containing elements for each saved custom filter in your organization, with details about each filter.

> **Note:** For details of the JSON name/value pairs in an array element, see **Saved Filter JSON**.

### Saved Filter Info

The Saved Filter Info `GET` method fetches details of a specific custom filter, identified by its ID, from the collection of filters saved in Kentik by your organization’s users.  
  
**HTTP Request**  
The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| **URL** | https://api.kentik.com/api/v5/saved-filter/custom/filter\_id |
| **Request** | GET /api/v5/saved-filter/custom/filter\_id HTTP/1.1  Host: api.kentik.com  X-CH-Auth-API-Token: user\_api\_token  X-CH-Auth-Email: user@domain.suffix  Content-Type: application/json |

> **Notes:**
>
> * The "filter\_id" in the path corresponds to the "id" of an array element returned from **Saved Filter List**.
> * If your organization is registered on Kentik's EU cluster, use `api.kentik.eu` in place of `api.kentik.com` in the URL above.

**HTTP Response**  
A successful response from the Saved Filter Info method includes the following elements:

* Response headers
* HTTP response code
* A JSON saved filter object with details of the custom filter specified by `filter_id`.

> **Note:** For details of the JSON name/value pairs in the returned saved filter object, see **Saved Filter JSON**.

### Saved Filter Create

The Saved Filter Create `POST` method adds a new custom filter to your collection of saved filters.  
  
**HTTP Request**  
The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| **URL** | https://api.kentik.com/api/v5/saved-filter/custom |
| **Request** | POST /api/v5/saved-filter/custom HTTP/1.1  Host: api.kentik.com  X-CH-Auth-API-Token: user\_api\_token  X-CH-Auth-Email: user@domain.suffix  Content-Type: application/json |

> **Note:** If your organization is registered on Kentik's EU cluster, use `api.kentik.eu` in place of `api.kentik.com` in the URL above.

The table below lists the fields with values defined in the saved filter JSON included in the request body:

| Parameter | Structure | Type | Description |
| --- | --- | --- | --- |
| filter\_name | **Saved Filter Object** | string | Required: A unique name for the custom filter. |
| filter\_description | **Saved Filter Object** | string | An optional description for the custom filter. |
| connector | **Filters Object** | string | Required: The conjunctive operator ("and" or "or") that will join all filter groups in this filter set. |
| connector | **FilterGroups Array** | string | Required: The conjunctive operator ("and" or "or") that will join all filters in this filter group. |
| not | **FilterGroups Array** | Boolean | Required: Sets whether traffic matching the filter group is included (false) or excluded (true). |
| filterField | **Filters Array** | string | The dimension to filter on. |
| operator | **Filters Array** | string | The operator to apply in the filter. |
| filterValue | **Filters Array** | string | The filterField value to match. |

> **Note:** For detailed information on a parameter in the table above, refer to the topic discussing the structure to which the parameter belongs.

**HTTP Response**  
A successful response from the Saved Filter Create method includes the following elements:

* Response headers
* HTTP response code
* A single JSON saved filter object containing information about the newly-added custom filter.

> **Note:** For details of the JSON name/value pairs in the returned filter object, see **Saved Filter JSON**.

### Saved Filter Update

The Saved Filter Update `PUT` method modifies a saved custom filter, identified by its ID, by updating one or more of its field values.  
  
**HTTP Request**  
The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| **URL** | https://api.kentik.com/api/v5/saved-filter/custom/filter\_id |
| **Request** | PUT /api/v5/saved-filter/custom/filter\_id HTTP/1.1  Host: api.kentik.com  X-CH-Auth-API-Token: user\_api\_token  X-CH-Auth-Email: user@domain.suffix  Content-Type: application/json |

> **Notes:**
>
> * The "filter\_id" in the path corresponds to the "id" of an array element returned from **Saved Filter List**.
> * If your organization is registered on Kentik's EU cluster, use `api.kentik.eu` in place of `api.kentik.com` in the URL above.

When updating a saved filter, the saved filter object in the request body must include all fields from **Saved Filter Create**, even those you don't wish to change. The specified values in this object will become the updated filter’s values.  
  
**HTTP Response**  
A successful response from the Saved Filter Update method includes the following elements:

* Response headers
* HTTP response code
* A JSON saved filter object containing information about the newly-updated custom filter.

> **Note:** For details of the JSON name/value pairs in the returned filter object, see **Saved Filter JSON**.

### Saved Filter Delete

The Saved Filter Delete method removes a specific custom filter, identified by its ID, from your collection of saved custom filters.  
  
**HTTP Request**  
The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| **URL** | https://api.kentik.com/api/v5/saved-filter/custom/filter\_id |
| **Request** | DELETE /api/v5/saved-filter/custom/filter\_id HTTP/1.1  Host: api.kentik.com  X-CH-Auth-API-Token: user\_api\_token  X-CH-Auth-Email: user@domain.suffix  Content-Type: application/json |

> **Notes:**
>
> * The "filter\_id" in the path corresponds to the "id" in each of the array elements returned from **Saved Filter List**.
> * If your organization is registered on Kentik's EU cluster, use api.kentik.eu in place of api.kentik.com in the URL above.

**HTTP Response**  
A successful response from the Saved Filter Delete method includes the following elements:

* Response headers
* HTTP response code 204 (no content)

## Custom Application API

The Custom Application API allows programmatic management of your organization's custom applications (see **About Applications**).  
  
More information is coming soon.
