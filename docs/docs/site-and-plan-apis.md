# Site and Plan APIs

Source: https://kb.kentik.com/docs/site-and-plan-apis

---

The site and plan APIs of the Kentik V5 Admin APIs are covered here.

> **Notes:**
>
> * For an overview of all Kentik APIs, see **Kentik APIs**.
> * For information on using APIs with cURL, see **API Access Via cURL**.
> * For assistance using any API version, contact support@kentik.com.

## About Site and Plan APIs

The Kentik V5 Admin APIs allow programmatic management of settings in the Admin section of the Kentik portal, includes the following:

* **Sites**: Managed via the **Site API**, offering functionality similar to the portal’s **Site Settings**.
* **Plans**: Managed via the **Plan API**, providing access to information on the portal's **Licenses Page**.

## Site API

> **Note:** Member-level users only have access to the GET methods of this API.

The Site API allows management of your company's sites in Kentik (see **About Sites**).

> **Notes:**
>
> * For information on adding and managing sites via the Kentik portal, see **Add or Edit a Site**.
> * To make calls to this API using cURL, see **API Access Via cURL**.
> * The V5 API tester was discontinued in January 2025.

### Site JSON

Calls to the Site API return an HTTP response body with a site JSON object, or an array of such objects for the Site List call. The object consists of the fields (name/value pairs) shown in the following example:

Each site object contains fields with information about an individual site registered to your organization in Kentik. These fields are described in the following table:

```
{
"site": {
     "id": 2,
     "site_name": "LAX",
     "lat": 33.7700504,
     "lon": -118.1937395,
     "company_id": 1289
   }
 }
```

| JSON name | Type | Description |
| --- | --- | --- |
| id | number | The system-assigned ID of the site. |
| site\_name | string | User-supplied name string for the site.   * Valid characters: alphanumeric, spaces, and underscores. * Length: min=3, max=40. |
| lat | number | The latitude of the site.   * Valid values: min=-90, max=90. |
| lon | number | The longitude of the site.   * Valid values: min=-180, max=180. |
| company\_id | number | The system-assigned ID of the customer. |

> **Note:** In the portal, the settings represented by the name/value pairs above are configured in the **Site Settings**.

### Site List

The Site List `GET` method retrieves a list of your sites in Kentik, presented as a JSON array where each element corresponds to an individual site.

> **Note:** The "id" in the `site` object can be used in subsequent calls to retrieve, update, or delete that site.

**HTTP Request**  
 The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| **URL** | https://api.kentik.com/api/v5/sites |
| **Request** | GET /api/v5/sites HTTP/1.1 Host: api.kentik.com X-CH-Auth-API-Token: user\_api\_token X-CH-Auth-Email: user@domain.suffix Content-Type: application/json |

> **Note:** If your organization is registered on Kentik's EU cluster, use `api.kentik.eu` in place of `api.kentik.com` in the URL above.

**HTTP Response**  
 A successful response from the Site List method includes the following elements:

* Response headers
* HTTP response code
* A response body with a JSON `site` array, where each element is a site object containing information about one of your organization’s sites.

> **Note:** For details of the JSON name/value pairs in a site object, see **Site JSON**.

### Site Info

The Site Info `GET` method retrieves information about a single site, identified by its ID, from your list of sites in Kentik.  
  
**HTTP Request**  
The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| **URL** | https://api.kentik.com/api/v5/site/site\_id |
| **Request** | GET /api/v5/site/site\_id HTTP/1.1 Host: api.kentik.com X-CH-Auth-API-Token: user\_api\_token X-CH-Auth-Email: user@domain.suffix Content-Type: application/json |

> **Notes:**
>
> * The "site\_id" in the path corresponds to the "id" in the site object from the **Site List** array.
> * If your organization is registered on Kentik's EU cluster, use `api.kentik.eu` in place of `api.kentik.com` in the URL above.

**HTTP Response**  
 A successful response from the Site Info method includes the following elements:

* Response headers
* HTTP response code
* A JSON `site` object containing information about the site specified by site\_id.

> **Note:** For details of the JSON name/value pairs in a site object, see **Site JSON**.

### Site Create

The Site Create `POST` method adds a new site to your sites in Kentik.  
  
**HTTP Request**  
The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| **URL** | https://api.kentik.com/api/v5/site |
| **Request** | POST /api/v5/site HTTP/1.1 Host: api.kentik.com X-CH-Auth-API-Token: user\_api\_token X-CH-Auth-Email: user@domain.suffix Content-Type: application/json |

> **Note:** If your organization is registered on Kentik's EU cluster, use `api.kentik.eu` in place of `api.kentik.com` in the URL above.

The following parameters are passed in a JSON object in the request body:

| JSON name | Type | Description |
| --- | --- | --- |
| site\_name | string | Required: User-supplied name string for the site.   * Valid characters: alphanumeric, spaces, and underscores. * Length: min=3, max=40. |
| lat | string | Required: The latitude of the site.   * Valid values: min=-90, max=90. |
| lon | string | Required: The longitude of the site.   * Valid values: min=-180, max=180. |

**HTTP Response**  
 A successful response from the Site Create method includes the following elements:

* Response headers
* HTTP response code
* A JSON `site` object containing information about the newly-added site.

> **Note:** For details of the JSON name/value pairs in a site object, see **Site JSON**.

### Site Update

The Site Update `PUT` method modifies data for a specific site, identified by its ID, from your list of existing Kentik sites.  
  
**HTTP Request**  
The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| **URL** | https://api.kentik.com/api/v5/site/site\_id |
| **Request** | PUT /api/v5/site/site\_id HTTP/1.1 Host: api.kentik.com X-CH-Auth-API-Token: user\_api\_token X-CH-Auth-Email: user@domain.suffix Content-Type: application/json |

> **Notes:**
>
> * The "site\_id" in the path corresponds to the "id" in the `site` object from the **Site List** array.
> * If your organization is registered on Kentik's EU cluster, use `api.kentik.eu` in place of `api.kentik.com` in the URL above.

Parameters to be updated are included in a JSON site object with only the fields to be changed. The following example shows an update to a site’s name:

```
{
"site": {
     "site_name": "SFO"
   }
 }
```

**HTTP Response**  
 A successful response from the Site Update method includes the following elements:

* Response headers
* HTTP response code
* A JSON `site` object containing information about the newly-updated site.

> **Note:** For details of the JSON name/value pairs in a `site` object, see **Site JSON**.

### Site Delete

The Site Delete method removes a specific site, identified by its ID, from your collection of Kentik sites.  
  
**HTTP Request**  
The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| **URL** | https://api.kentik.com/api/v5/site/site\_id |
| **Request** | DELETE /api/v5/site/site\_id HTTP/1.1 Host: api.kentik.com X-CH-Auth-API-Token: user\_api\_token X-CH-Auth-Email: user@domain.suffix Content-Type: application/json |

> **Notes:**
>
> * The "site\_id" in the path corresponds to the "id" in the site object from the **Site List** array.
> * If your organization is registered on Kentik's EU cluster, use `api.kentik.eu` in place of `api.kentik.com` in the URL above.

**HTTP Response**  
 A successful response from the Site Delete method includes the following elements:

* Response headers
* HTTP response code 204 (no content)

## Plan API

The Plan API allows for programmatic management of your organization's Kentik plans (see **About Licenses**).

> Notes:
>
> * To view information on your organization's plans, go to the Licenses Page (Menu » Licenses) of the Kentik portal.
> * To make calls to this API using cURL, see **API Access Via cURL**.

Calls to the Plan API return an HTTP response body with a "plans" JSON array containing all your current plans. The following example shows a single plan ("Default") that supports three device types and includes three assigned devices:

### Plan JSON

```
{
"plans": [
    {
      "id": 550,
      "company_id": 1289,
      "name": "Default"
      "description": "This is the default plan to which your devices should be added.",
      "active": true,
      "max_devices": 75,
      "max_fps": 4000,
      "bgp_enabled": true,
      "fast_retention": 120,
      "full_retention": 120,
      "cdate": "2016-10-27T00:35:42.652Z",
      "edate": null,
      "max_bigdata_fps": 30,
      "deviceTypes": [
        {
          "device_subtype": "router"
        },
        {
          "device_subtype": "host-nprobe-basic"
        },
        {
          "device_subtype": "host-nprobe-dns-www"
        }
      ],
      "devices": [
        {
          "device_name": "superx4",
          "device_subtype": "router",
          "id": "1206"
        },
        {
          "device_name": "superx2_readnews_com",
          "device_subtype": "router",
          "id": "4806"
        },
        {
          "device_name": "host_nprobe_basic",
          "device_subtype": "host-nprobe-basic",
          "id": "8937"
        }
      ],
     }
   ]
 }
```

> **Note:** The returned JSON may not include all fields shown above, depending on the defined filter groups and individual filters in a plan.

Each plan object provides information on an individual plan. The structure of this object is described in the following topics.

#### Plan Object

The highest level in the hierarchy of a plan, containing the following elements:

| JSON name | Type | Description |
| --- | --- | --- |
| id | number | The system-assigned unique ID of the plan. |
| company\_id | number | The system-assigned unique ID of your organization. |
| name | string | A name for the plan. Every Kentik customer is initially provided with a plan called "Default."  **Note:** Plan names are assigned by the Kentik sales team in coordination with customers and are not necessarily unique. If multiple plan objects return with the same name, check the value of the "active" field to see which plan is active. |
| description | string | An optional description of the plan. |
| active | boolean | Indicates if the plan is currently activated. |
| max\_devices | number | The maximum number of devices that can send flow records to Kentik under this plan. |
| max\_fps | number | Per device limit on flow records per second that can be sent to Kentik (excess FPS may trigger rate-limiting). |
| bgp\_enabled | boolean | Indicates whether or not devices on this plan may be peered to enable Kentik to collect BGP routing data. |
| fast\_retention | number | The number of days that data will be stored in the Fast dataseries. |
| full\_retention | number | The number of days that data will be stored in the Full dataseries. |
| cdate | string | The system-assigned date-time of filter creation, in UTC (ISO 8601), e.g. 2015-01-27T01:39:17.186Z |
| edate | string | The system-assigned date-time of most-recent modification, in UTC (ISO 8601), e.g., 2015-01-27T01:39:17.186Z |
| max\_bigdata\_fps | object | Max FPS applied to fast data rollups. See **About Dataseries Resolution**. |
| deviceTypes | array | The types of allowed devices. See **DeviceTypes Array**. |
| devices | array | The devices currently assigned to this plan. See **Devices Array for Plans**. |

#### DeviceTypes Array

| JSON name | Type | Description |
| --- | --- | --- |
| device\_subtype | string | A type of device that sends flow records to Kentik.   * Valid values: see **Device Subtypes**. |

#### Devices Array for Plans

An array of `device` objects is present, each containing a subset of elements from the full `device` object (see **Device JSON**). Each device in the array represents a `device` assigned to this plan. The table below lists the included elements.

| JSON name | Type | Description |
| --- | --- | --- |
| device\_name | string | The user-assigned name of a device associated with this plan. |
| device\_subtype | string | A type of device that sends flow records to Kentik.   * Valid values: see **Device Subtypes**. |
| id | number | The system-assigned unique ID of a device associated with this plan. |

  

### Plan List

The Plan List `GET` method retrieves a JSON array of plans saved by users in your organization, with each element representing an individual plan.  
  
**HTTP Request**  
The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| **URL** | https://api.kentik.com/api/v5/plans |
| **Request** | GET /api/v5/plans HTTP/1.1 Host: api.kentik.com X-CH-Auth-API-Token: user\_api\_token X-CH-Auth-Email: user@domain.suffix Content-Type: application/json |

> **Note:** If your organization is registered on Kentik's EU cluster, use `api.kentik.eu` in place of `api.kentik.com` in the URL above.

**HTTP Response**  
 A successful response from the Plan List method includes the following elements:

* Response headers
* HTTP response code
* A response body with a JSON array with an element containing information for each saved plan in your organization.

> **Note:** For details of the JSON name/value pairs in each element of the array, see **Plan JSON**.

---

© 2014-25 Kentik
