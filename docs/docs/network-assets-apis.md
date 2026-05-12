# Network Assets APIs

Source: https://kb.kentik.com/docs/network-assets-apis

---

This article discusses the network assets APIs of the Kentik V5 Admin APIs.

> **Notes:**
>
> * For an overview of all Kentik APIs, see **Kentik APIs**.
> * For information on using APIs with cURL, see **API Access Via cURL**
> * For help using any API version, contact **support@kentik.com**.

## About Network Assets APIs

The Kentik V5 Admin APIs offer "network assets" APIs for programmatically managing settings in the Admin section of the Kentik portal, including the following:

* **Devices and interfaces**: Managed via the **Device API**, offering similar functionality to the portal’s **Device Settings Dialog**.
* **Device labels**: Managed via the **Device Label API**, provides functionality equivalent to the portal’s **Labels Page**.

## Device API

The Device API allows you to manage your organization's devices in Kentik.

> **Notes:**
>
> * Member-level users only have access to the GET methods of this API.
> * For information on adding and managing devices via the Kentik portal, see **Manage Devices**.
> * To make calls to this API using cURL, see **API Access Via cURL**.
> * The V5 API tester was discontinued in January 2025.

### Device JSON

Device API calls return an HTTP response with a JSON device object, or, an array of such objects for the Device List call. The object contains the fields (name/value pairs) shown in the following example:

```
{
"device": {
"id": "9623",
     "company_id": "1289",
     "device_name": "pd_router",
     "device_subtype": "router",
     "device_description": "test of device settings",
     "plan_id": 123,
     "site_id": 394,
     "device_flow_type": "netflow.v5",
     "device_sample_rate": "1024",
     "sending_ips": [
       "142.254.47.216"
    ],
    "device_snmp_ip": "142.254.47.216",
     "device_snmp_community": "",
     "minimize_snmp": false,
     "device_bgp_type": "device",
     "device_bgp_neighbor_ip": "142.254.47.216",
     "device_bgp_neighbor_asn": "5001",
     "device_bgp_password": null,
     "use_bgp_device_id": null,
     "custom_columns": null,
     "created_date": "2016-09-09T17:27:18.080Z",
     "updated_date": "2016-09-09T17:27:18.080Z"
     "device_snmp_v3_conf": {
       "UserName": "eenie",
       "AuthenticationProtocol": "MD5",
       "AuthenticationPassphrase": "meenie",
       "PrivacyProtocol": "DES",
       "PrivacyPassphrase": "minie"
     },
   }
 }
```

Each device object contains fields with information about an individual device registered in Kentik to your organization. These fields are detailed in the following tables.

#### Device Object

| JSON name | Type | Description |
| --- | --- | --- |
| id | number | The system-assigned device ID. |
| company\_id | number | The system-assigned ID of the company (organization). |
| device\_name | string | User-supplied device name. |
| device\_subtype | string | The subtype of the device, see **Device Subtypes**. |
| device\_description | string | User-supplied device description. |
| plan\_id | number | The ID of the plan to which this device is assigned. Available plans can be found via the Plans API.   * Valid value: integer. |
| site\_id | number | The system-assigned ID of the site, if any, to which this device is assigned. |
| device\_flow\_type | string | For routers, the flow type sent to Kentik (NetFlow v5, NetFlow v9, sFlow, or IPFIX). For hosts, it’s IPFIX, the flow type sent from the Kentik nProbe host agent.  **Note:** Kentik auto-detects flow type. |
| device\_sample\_rate | string | Total packets transiting the device for each packet processed for flow data (see **Flow Sampling**). |
| sending\_ips | array of strings | IPs from which the device will send flow to Kentik. |
| device\_snmp\_ip | string | The IP address for polling the device (router). |
| device\_snmp\_community | string | The SNMP community for polling the device (router). |
| minimize\_snmp | boolean | * The SNMP polling interval * If false (standard), interface counter is polled every 5 minutes and interface description every 30 minutes. * If true (minimized), interface counter is not polled, and interface description is polled every 6 hours. |
| device\_bgp\_type | string | Reserved for internal use. |
| device\_bgp\_neighbor\_ip | string | BGP Neighbor IP |
| device\_bgp\_neighbor\_asn | string | BGP Neighbor ASN (16- or 32-bit) |
| device\_bgp\_password | string | BGP Password |
| use\_bgp\_device\_id | string | Use this device's BGP table instead of creating a new BGP session. |
| custom\_columns | N.A. | Internal use only. |
| created\_date | string | Date-time of registration of the device in Kentik creation, in UTC (ISO 8601), e.g., 2015-01-27T01:39:17.186Z |
| updated\_date | string | Date-time of last edit to the device in UTC, e.g., 2015-01-27T01:39:17.186Z |
| device\_snmp\_v3\_conf | object | See **Device\_snmp\_v3\_conf Object** |

> **Note:** For more details about the settings indicated by the name/value pairs, see **Device Settings Dialog**.

#### Device Subtypes

Kentik supports the device subtypes found in the Subtype column of the table in **Supported Device Types**.

#### Device\_snmp\_v3\_conf Object

| JSON name | Type | Description |
| --- | --- | --- |
| UserName | string | The user name for SNMP v3 authentication.  **Note:** Required if AuthenticationProtocol is MD5 or SHA, or if PrivacyProtocol is DES or AES-128. |
| AuthenticationProtocol | string | SNMP v3 authentication protocol.   * Valid values: NoAuth (none), MD5, SHA. |
| AuthenticationPassphrase | string | Password for SNMP V3 authentication.  **Note:** Required if AuthenticationProtocol is MD5 or SHA. |
| PrivacyProtocol | string | The SNMP V3 privacy type:   * Valid values: NoPriv (none), DES (56-bit DES encryption), AES-128. |
| PrivacyPassphrase | string | Password for SNMP V3 privacy.  **Note:** Required if PrivacyProtocol is DES or AES-128. |

  

### Device List

The Device List `GET` method fetches JSON array of your Kentik devices, with each element representing an individual device.

> **Note:** The “id” value in each device object can be used in subsequent calls to retrieve, update, or delete that specific device.

**HTTP Request**

The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| **URL** | https://api.kentik.com/api/v5/devices |
| **Request** | GET /api/v5/devices HTTP/1.1  Host: api.kentik.com  X-CH-Auth-API-Token: user\_api\_token  X-CH-Auth-Email: user@domain.suffix  Content-Type: application/json |

> **Note:** If your organization is registered on Kentik’s EU cluster, use `api.kentik.eu` in place of `api.kentik.com` in the URL above.

**HTTP Response**

A successful response from the Device List method includes the following elements:

* Response headers
* HTTP response code
* Response body with a JSON device array, where each element is a device object containing device information.

> **Note:** For details on the JSON name/value pairs in a device object, see **Device JSON**.

### Device Info

The Device Info `GET` method fetches details of a single device, identified by its ID, from your Kentik devices list.

**HTTP Request**

The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| **URL** | https://api.kentik.com/api/v5/device/device\_id |
| **Request** | GET /api/v5/device/device\_id HTTP/1.1  Host: api.kentik.com  X-CH-Auth-API-Token: user\_api\_token  X-CH-Auth-Email: user@domain.suffix  Content-Type: application/json |

> **Notes:**
>
> * The “device\_id” in the path corresponds to the “id” value in each `device` object from the **Device List** array.
> * If your organization is registered on Kentik’s EU cluster, use `api.kentik.eu` in place of `api.kentik.com` in the URL above.

**HTTP Response**

A successful response from the Device Info method includes the following elements:

* Response headers
* HTTP response code
* A single JSON `device` object with details about the device specified by `device_id`.

> **Note:** For details on the JSON name/value pairs in a `device` object, see **Device JSON**.

### Device Create

The Device Create `POST` method adds a new device to Kentik.

**HTTP Request**

The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| **URL** | https://api.kentik.com/api/v5/device |
| **Request** | POST /api/v5/device HTTP/1.1  Host: api.kentik.com  X-CH-Auth-API-Token: user\_api\_token  X-CH-Auth-Email: user@domain.suffix  Content-Type: application/json |

> **Note:** If your organization is registered on Kentik’s EU cluster, use `api.kentik.eu` in place of `api.kentik.com` in the URL above.

The request body includes the following parameters in a JSON object:

| Parameter | Type | Description |
| --- | --- | --- |
| device\_name | string | Required: The name of the device:   * Valid characters: alphanumeric and underscores. * Length: min=4, max=60. |
| Device\_subtype | string | Required: The type of device:   * Valid values: see **Device Subtypes**. |
| Device\_description | string | A description of this device:   * Valid characters: any. * Length: max=128. |
| Plan\_id | number | Required: The ID of the plan to which this device is assigned. Available plans can be found via the Plans API.   * Valid value: integer. |
| Site\_id | number | The ID of the site, if any, to which this device will be assigned. Site IDs are system-generated upon site creation.   * Valid value: integer. |
| Device\_sample\_rate | string | Required: Total packets transiting the device for each packet processed for flow data. |
| Sending\_ips | array of strings | Required: The IP(s) from which flow is sent.  **Note:** The IP must not already be sending flow to another device registered in Kentik. |
| Device\_snmp\_ip | string | The SNMP IP for polling the device.  **Note:** Ignored unless device\_subtype is set to router. |
| Device\_snmp\_community | string | The SNMP community for polling the device.  **Note:** Ignored unless device\_subtype is set to router. |
| Minimize\_snmp | boolean | Required when device type is router: The interval at which SNMP will be polled:   * If false (standard), interface counter is polled every 5 minutes and interface description every 3 hours. * If true (minimized), interface counter is not polled and interface description is polled every 6 hours. |
| Device\_snmp\_v3\_conf | object | See **Setting device\_snmp\_v3\_conf**.  **Note:** If included, Kentik will poll this router with SNMP V3. |
| Device\_bgp\_type | string | Required: Device BGP type. Valid values:   * “none” (use generic IP/ASN mapping) * “device” (peer with the device itself) * “other\_device” (share routing table of existing peered device) |
| device\_bgp\_neighbor\_ip | string | A valid IPv4 address for peering with this device.  **Notes:**   * Either this or its IPv6 equivalent is required when device\_bgp\_type is set to “device.” * IP must not already be peering with another Kentik device. |
| Device\_bgp\_neighbor\_ip6 | string | A valid IPv6 address for peering with this device.  **Notes:**   * Either this or its IPv4 equivalent is required when device\_bgp\_type is set to “device.” * IP must not already be peering with another Kentik device. |
| Device\_bgp\_neighbor\_asn | string | The valid ASN (16 or 32-bit) of the autonomous system to which this device belongs.  **Note:** Required when device\_bgp\_type is set to “device.” |
| Device\_bgp\_password | string | Optional BGP MD5 password (shared authentication password for BGP peering). Valid characters: alphanumeric. Length: 32.  **Note:** Required when device\_bgp\_type is set to “device.” |
| Use\_bgp\_device\_id | string | The ID of the device whose BGP table will be shared with this device.  **Note:** Required when device\_bgp\_type is set to “other\_device”).   * Valid value: a system-generated device\_id. |

#### Setting device\_snmp\_v3\_conf

The `device_snmp_v3_conf` object is included in request JSON to use SNMP V3 for router polling. If omitted, Kentik will assume SNMP V3 is intended.

| Parameter | Type | Description |
| --- | --- | --- |
| UserName | string | The username for SNMP v3 authentication.  **Note:** Required if AuthenticationProtocol is MD5 or SHA, or if PrivacyProtocol is DES or AES-128. |
| AuthenticationProtocol | string | SNMP v3 authentication protocol.   * Valid values: NoAuth (none), MD5, SHA. |
| AuthenticationPassphrase | string | Password for SNMP V3 authentication.  **Note:** Required if AuthenticationProtocol is MD5 or SHA. |
| PrivacyProtocol | string | The SNMP V3 privacy type:   * Valid values: NoPriv (none), DES (56-bit DES encryption), AES-128. |
| PrivacyPassphrase | string | Password for SNMP V3 privacy.  **Note:** Required if PrivacyProtocol is DES or AES-128. |

> **Note:** SNMP V3 authentication and privacy settings are independent. To disable both, set  AuthenticationProtocol to NoAuth and PrivacyProtocol to NoPriv.

**HTTP Response**

A successful response from the Device Create method includes the following elements:

* Response headers
* HTTP response code
* A single JSON `device` object with details about the newly-added device

> **Note:** For details on the JSON name/value pairs in a `device` object, see **Device JSON**.

### Device Update

The Device Update `PUT` method updates data for a specific device, identified by ID, in your Kentik devices list.

**HTTP Request**

The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| **URL** | https://api.kentik.com/api/v5/device/device\_id |
| **Request** | PUT /api/v5/device/device\_id HTTP/1.1  Host: api.kentik.com  X-CH-Auth-API-Token: user\_api\_token  X-CH-Auth-Email: user@domain.suffix  Content-Type: application/json |

> **Notes:**
>
> * The “device\_id” in the path corresponds to the “id” in each device object from the **Device List** array.
> * If your organization is registered on Kentik’s EU cluster, use `api.kentik.eu` in place of `api.kentik.com` in the URL above.

Pass the parameters to be updated in a JSON `device` object containing only the fields to be changed. Here’s an example for updating a device’s description:

```
{
“device”: {
     “device_description”: “This is a really wonderful device.”
   }
 }
```

**HTTP Response**

A successful response from the Device Update method includes the following elements:

* Response headers
* HTTP response code
* A single JSON `device` object with details about the newly-updated device.

> **Note:** For details on the JSON name/value pairs in a `device` object, see **Device JSON**.

### Device Apply Labels

The Device Apply Labels PUT method removes all existing labels (see **Labels**) from the specified device and applies the specified new labels.

> **Notes:**
>
> * If no labels are specified, this method will still remove all existing labels from the specified device.
> * To create and manage device labels via API, see **Device Label API**.

**HTTP Request**

The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| **URL** | api.kentik.com/api/v5/devices/device\_id/labels |
| **Request** | POST /api/v5/devices/device\_id/labels HTTP/1.1  Host: api.kentik.com  X-CH-Auth-API-Token: user\_api\_token  X-CH-Auth-Email: user@domain.suffix  Content-Type: application/json |

> **Note:**
>
> If your organization is registered on Kentik's EU cluster, use `api.kentik.eu` in place of `api.kentik.com` in the URL above.

The following parameters are passed in a JSON object in the request body:

| Parameter | Type | Description |
| --- | --- | --- |
| labels | array | An array of id objects.  **Note:** Omitting the labels array in the request body will result in the removal of all existing labels without adding new ones. |
| id | integer | The id of a label to apply to the device. |

The example below shows a request body for adding two labels to a device.

```
{
"labels": [
    {
      "id": 1096
    },
    {
      "id": 32
    }
  ]
 }
```

**HTTP Response**

A successful response from the Device Apply Labels method includes the following elements:

* Response headers
* HTTP response code
* A single JSON device object with a labels array showing the labels assigned to the device post-operation. An example of a successful label application is provided below.

```
{
"id": "device_id",
"device_name": "device_name",
"labels": [
    {
      "id": label_id,
      "name": "label_name",
      "edate": "2019-03-07T00:53:51.759Z",
      "cdate": "2019-03-07T00:53:51.759Z",
      "user_id": "#####",
      "company_id": "####",
      "color": "#800000",
    }
  ]
 }
```

> **Note:** For details on the JSON name/value pairs in an individual labels object, see **Device Label Object**.

### Device Delete

The Device Delete method removes one device, identified by ID, from your collection of Kentik devices.

**HTTP Request**

The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| **URL** | https://api.kentik.com/api/v5/device/device\_id |
| **Request** | DELETE /api/v5/device/device\_id HTTP/1.1  Host: api.kentik.com  X-CH-Auth-API-Token: user\_api\_token  X-CH-Auth-Email: user@domain.suffix  Content-Type: application/json |

> **Notes:**
>
> * The “device\_id” in the path corresponds to the “id” in each `device` object from the **Device List** array.
> * If your organization is registered on Kentik's EU cluster, use `api.kentik.eu` in place of `api.kentik.com` in the URL above.

**HTTP Response**

A successful response from the Device Delete method includes the following elements:

* Response headers
* HTTP response code 204 (no content)

### Interface JSON

The interface methods of the Device API return an HTTP response body with an "interface" JSON object (or an array of objects for the Interface List call). The object includes name/value pairs as shown in the following example:

```
{
"id": "9201304925",
"company_id": "1289",
"device_id": "2951",
"snmp_id": "150",
"snmp_speed": "1000",
"snmp_type": 6,
"snmp_type": 53,
"snmp_alias": "irs",
"interface_ip": "198.186.192.33",
"interface_description": "Vlan350",
"interface_kvs": "\"updated\"=>\"1\"",
"interface_tags": "",
"interface_status": "V",
"cdate": "2017-04-19T00:40:01.145Z",
"edate": "2017-05-03T03:20:27.303Z",
"initial_snmp_id": "150",
"initial_snmp_alias": "irs",
"initial_interface_description": "Vlan350",
"initial_snmp_speed": "1000",
"interface_ip_netmask": "255.255.255.224",
"secondary_ips": [
    {
      "address": "198.186.193.51",
      "netmask": "255.255.255.240"
    }
  ],
"connectivity_type": "",
"network_boundary": "",
"initial_connectivity_type": "",
"initial_network_boundary": ""
"top_nexthop_asns": null,
"provider":"",  
"initial_provider": "",
"vrf_id": ####
"vrf": {
     "name" : "test_vrf",
     "description" : "This is still a test",
     "route_distinguisher" : "11.121.111.13:3254",
     "ext_route_distinguisher" : 296507164417724,
     "route_target" : "101:100"
   }
 }
```

Each interface object contains fields with information about an individual interface on a Kentik-registered device. These fields are described in the following tables.

> **Note:** For details on the settings represented by the name/value pairs (see **Interfaces Page**).

#### Interface Object

An `interface` object contains all the information about an individual interface of a device registered with Kentik.

| JSON name | Type | Description |
| --- | --- | --- |
| id | string | The system-assigned ID of the interface. |
| Company\_id | string | The system-assigned ID of the customer. |
| Device\_id | string | The system-assigned ID of the device to which this interface belongs. |
| Snmp\_id | string | The interface index (ifIndex) as defined in the device and retrieved via SNMP. Same as initial\_snmp\_id unless manually overridden when device is created or updated in Kentik. |
| Snmp\_speed | number | The capacity of the interface in Mbps. Same as initial\_snmp\_speed unless manually overridden when device is created or updated in Kentik. |
| Snmp\_type | integer | Reserved for internal use. |
| Snmp\_alias | string | User-specified text used by Kentik as the description of the interface. Same as initial\_interface\_description unless manually overridden when device is created or updated in Kentik. |
| Interface\_ip | string | The IP address assigned to the interface. |
| Interface\_description | string | User-specified text used by Kentik as the name of the interface. Same as initial\_snmp\_alias unless manually overridden when device is created or updated in Kentik. |
| Interface\_kvs | number | Reserved for internal use. |
| Interface\_tags | number | Reserved for internal use. |
| Interface\_status |  | Reserved for internal use. |
| Cdate | string | Date-time of registration of the interface in Kentik creation, in UTC (ISO 8601), e.g. 2015-01-27T01:39:17.186Z |
| edate | string | Date-time of last edit to the interface, in UTC, e.g. 2015-01-27T01:39:17.186Z |
| initial\_snmp\_id | string | Last SNMP-polled SNMP ID of the interface on its device. |
| Initial\_snmp\_alias | string | Last SNMP-polled interface description. |
| Initial\_interface\_description | string | Last SNMP-polled interface name. |
| initial\_snmp\_speed | string | Last SNMP-polled interface capacity (Mbps). |
| Interface\_ip\_netmask | string | The netmask configured for the interface (applies to Ipv4 interfaces only). |
| Secondary\_ips | array | See **secondary\_ips Array**. |
| Connectivity\_type | string | Reserved for internal use. |
| Network\_boundary | string | Reserved for internal use. |
| Initial\_connectivity\_type | string | Same as connectivity\_type. |
| Initial\_network\_boundary | string | Same as network\_boundary. |
| Top\_nexthop\_asns | array | An array of ASNs, see **Top Next Hop ASNs**. |
| Provider | object | The provider, as determined by provider classification, via which traffic from a given externally facing interface reaches the Internet; see **Provider Classification**. |
| initial\_provider | string | Reserved for internal use. |
| vrf\_id | string | If the interface is assigned to a VRF, the ID of that VRF, otherwise null. |
| vrf | string | Present only if the interface is assigned to a VRF; see **VRF Attributes**. |

#### secondary\_ips Array

The `secondary_ips` array contains one or more objects, each specifying a secondary IP address for the interface.

| JSON name | Type | Description |
| --- | --- | --- |
| address | string | A secondary IP for the interface. |
| netmask | string | The netmask configured for this secondary interface\_ip (applies to IPv4 interfaces only). |

#### Top Next Hop ASNs

An array of objects showing ASNs that received the most packets from the interface in the past hour.

| JSON name | Type | Description |
| --- | --- | --- |
| ASN | number | The AS number. |
| packets | number | The number of packets in the hour preceding the call. |

#### VRF Attributes

The `vrf` object, present only if the interface is assigned to a VRF, consists of elements corresponding to each VRF attribute:

| Parameter | Type | Description |
| --- | --- | --- |
| name | string | Name of the VRF entry. |
| description | string | Description of the VRF entry. |
| route\_target | string | Route Target Community of the VRF entry. |
| route\_distinguisher | string | Route Distinguisher of the VRF entry. |
| ext\_route\_distinguisher | integer | External Route Distinguisher of the VRF entry. |

  

### Interface List

The Interface List `GET` method retrieves a JSON array of Kentik interfaces for a specified device, with each element representing an individual interface.

> **Note:** The "id" value in each interface object can be used in subsequent calls to retrieve, update, or delete that specific interface.

**HTTP Request**

The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| URL | https://api.kentik.com/api/v5/device/device\_id/interfaces |
| Request | GET /api/v5/device/device\_id/interfaces HTTP/1.1  Host: api.kentik.com  X-CH-Auth-API-Token: user\_api\_token  X-CH-Auth-Email: user@domain.suffix  Content-Type: application/json |

> **Note:** If your organization is registered on Kentik's EU cluster, use api.kentik.eu in place of api.kentik.com in the URL above.

**HTTP Response**

A successful response from the Interface List method includes:

* Response headers
* HTTP response code
* A JSON response body with an array of interface objects, each representing an interface of the specified device.

> **Note:** For details of the JSON name/value pairs in an interface object, see **Interface JSON**.

### Interface Info

The Interface Info `GET` method fetches details about a single interface, identified by its ID, from the specified device's collection of interfaces.

**HTTP Request**

The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| **URL** | https://api.kentik.com/api/v5/device/device\_id/interface/interface\_id |
| **Request** | GET /api/v5/device/device\_id/interface/interface\_id HTTP/1.1  Host: api.kentik.com  X-CH-Auth-API-Token: user\_api\_token  X-CH-Auth-Email: user@domain.suffix  Content-Type: application/json |

> **Notes:**
>
> * The “interface\_id” in the path corresponds to the “id” in each interface object from the **Interface List** array.
> * If your organization is registered on Kentik's EU cluster, use api.kentik.eu in place of api.kentik.com in the URL above.

**HTTP Response**

A successful response from the `Interface` Info method includes the following elements:

* Response headers
* HTTP response code
* A single JSON `interface` object with details about the interface specified by `interface_id.`

> **Note:** For details on the JSON name/value pairs in an `interface` object, see **Interface JSON**.

### Interface Create

The Interface Create `POST` method adds a new interface to the specified Kentik-registered device’s collection of interfaces.

HTTP Request

The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| **URL** | https://api.kentik.com/api/v5/device/device\_id/interface |
| **Request** | POST /api/v5/device/device\_id/interface HTTP/1.1  Host: api.kentik.com  X-CH-Auth-API-Token: user\_api\_token  X-CH-Auth-Email: user@domain.suffix  Content-Type: application/json |

> **Note:** If your organization is registered on Kentik's EU cluster, use `api.kentik.eu` in place of `api.kentik.com` in the URL above.

The following parameters are passed in a JSON object in the request body:

| Parameter | Type | Description |
| --- | --- | --- |
| device\_id | string | Required: The system-assigned ID of the device to which this interface belongs. |
| snmp\_id | string | Required: The interface index (ifIndex) as defined in the device and retrieved via SNMP. See **Getting an Interface Index**. |
| interface\_description | string | Required: User-specified text used by Kentik as the name of the interface. |
| snmp\_alias | string | User-specified text used by Kentik as the description of the interface. |
| interface\_ip | string | The IP address assigned to the interface. |
| interface\_ip\_netmask | string | The netmask configured for the interface (applies to IPv4 interfaces only). |
| snmp\_speed | number | Required: The capacity of the interface in Mbps. |
| secondary\_ips | array | Array of secondary\_ip objects; see **Secondary IPs Array**. |
| vrf | object | An object specifying the VRF attributes of the interface; see **VRF Object**. |

#### Secondary IPs Array

An array of secondary IP objects, each of which must include the address and netmask parameters.

| Parameter | Type | Description |
| --- | --- | --- |
| address | string | In secondary\_ip object, a secondary ip for the interface. |
| netmask | string | In secondary\_ip object, the netmask configured for this secondary interface\_ip (applies to IPv4 interfaces only) |

#### VRF Object

An object specifying the VRF attributes of the interface:

| Parameter | Type | Description |
| --- | --- | --- |
| name | string | Required: Name of the VRF entry. |
| description | string | Required: Description of the VRF entry. |
| route\_target | string | Required: Route Target Community of the VRF entry. |
| route\_distinguisher | string | Required: Route Distinguisher of the VRF entry. |
| ext\_route\_distinguisher | integer | External Route Distinguisher of the VRF entry. |

**HTTP Response**

A successful response from the Interface Create method includes the following elements:

* Response headers
* HTTP response code
* A single JSON `interface` object with details about the newly-added interface.

> **Note:** For details of the JSON name/value pairs in an interface object, see **Interface JSON**.

### Interface Update

The Interface Update `PUT` method modifies system data for a specific interface, identified by ID, within the specified device's existing interfaces.

HTTP Request

The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| URL | https://api.kentik.com/api/v5/device/device\_id/interface/interface\_id |
| Request | PUT /api/v5/device/device\_id/interface/interface\_id HTTP/1.1  Host: api.kentik.com  X-CH-Auth-API-Token: user\_api\_token  X-CH-Auth-Email: user@domain.suffix  Content-Type: application/json |

> **Notes**:
>
> * The “interface\_id” in the path corresponds to the “id” in each interface object from the **Interface List** array.
> * If your organization is registered on Kentik's EU cluster, use api.kentik.eu in place of api.kentik.com in the URL above

When updating an interface, the JSON object must include the required snmp\_id and any fields to be updated (see table under HTTP Request in **Interface Create**). The following example shows an update to the interface’s description:

```
{
"snmp_id": "150"
"interface_description": "This is a truly great interface."
 }
```

HTTP Response

A successful response from the Interface Update call includes the following elements:

* Response headers
* HTTP response code
* A single JSON `interface` object with details about the newly-updated interface.

> **Note**: For details of the JSON name/value pairs in an `interface` object, see **Interface JSON**.

### Interface Delete

The Interface Delete method removes a specified interface by ID from a Kentik-registered device’s collection of interfaces.

HTTP Request

The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| URL | https://api.kentik.com/api/v5/device/device\_id/interface/interface\_id |
| Request | DELETE /api/v5/device/device\_id/interface/interface\_id HTTP/1.1  Host: api.kentik.com  X-CH-Auth-API-Token: user\_api\_token  X-CH-Auth-Email: user@domain.suffix  Content-Type: application/json |

> **Notes**:
>
> * The “interface\_id” in the path corresponds to the “id” in each `interface` object from the **Interface List**array.
> * If your organization is registered on Kentik's EU cluster, use `api.kentik.eu` in place of `api.kentik.com` in the URL above.

HTTP Response

A successful response from the Interface Delete method includes the following elements:

* Response headers
* The HTTP response code 204 (no content)

## Device Label API

The Device Label API allows for programmatic management of your device labels in Kentik.

> **Notes**:
>
> * Member-level users whose only have access to the GET methods of this API
> * For information on adding and managing device labels via the Kentik portal, see **Labels**.
> * To make calls to this API using cURL, see **API Access Via cURL**.

### Device Label JSON

Calls to the Device Label API return an HTTP response body with a JSON device label object, or an array of such objects for the Device Label List call. The object includes fields (name/value pairs) as shown in the following example:

```
{
"id": 32,
"name": "ISP",
"color": "#f1d5b9",
"user_id": "#####",
"company_id": "####",
"order": 0,
"devices": [
    {
      "id": "#####",
      "device_name": "my_device_1",
      "device_subtype": "router"
    },
  ],
"created_date": "2018-05-16T20:21:10.406Z",
"updated_date": "2018-05-16T20:21:10.406Z"
 }
```

Each device label object contains fields with information about an individual device label. These fields are described in the following tables.

> **Note**: For details of the settings represented by these name/value pairs, see **Labels**.

#### Device Label Object

| JSON name | Type | Description |
| --- | --- | --- |
| id | number | The system-assigned ID of the device label. |
| name | string | User-specified text of the device label (e.g. "ISP"). |
| color | string | The color, in hex, associated with this device label (e.g., "#f1d5b9"). |
| user\_id | string | The ID of the user who created or last edited the label. |
| company\_id | string | The system-assigned ID of the customer in which the label exists. |
| devices | array | See Devices Array for Labels. |
| created\_date | string | Date-time in UTC (ISO 8601) at which the device label was created (e.g., 2015-01-27T01:39:17.186Z). |
| updated\_date | string | Date-time at which the label was most-recently edited. |

#### Devices Array for Labels

An array of `device` objects, each containing a subset of elements from the full `device` object (see **Device JSON**), represents devices associated with labels. The table below lists the included elements.

| JSON name | Type | Description |
| --- | --- | --- |
| id | string | The system-assigned ID of the device. |
| device\_name | string | Name specified by user when device was last edited. |
| device\_subtype | string | See **Device Subtypes**. |

  

### Device Label List

The Device Label List `GET` method retrieves a JSON array of device labels that currently existing in your organization, with each element representing an individual device label.

> **Note**: The "id" value in each device label object can be used in subsequent calls to retrieve, update, or delete that specific device label.

HTTP Request

The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| URL | api.kentik.com/api/v5/deviceLabels |
| Request | GET /api/v5/deviceLabels HTTP/1.1  Host: api.kentik.com  X-CH-Auth-API-Token: user\_api\_token  X-CH-Auth-Email: user@domain.suffix  Content-Type: application/json |

> **Note**: If your organization is registered on Kentik's EU cluster, use `api.kentik.eu` in place of `api.kentik.com` in the URL above.

HTTP Response

A successful response from the Device Label List call includes the following elements:

* Response headers
* HTTP response code
* A response body including a JSON array of device labels. Each element in the array is an object containing information about a device label.

> **Note**: For details of the JSON name/value pairs in a device label object, see **Device Label JSON**.

### Device Label Info

The Device Label Info `GET` method fetches details of a specific device label by ID from your organization's collection.

HTTP Request

The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| URL | api.kentik.com/api/v5/deviceLabels/label\_id |
| Request | GET api/v5/deviceLabels/label\_id HTTP/1.1  Host: api.kentik.com  X-CH-Auth-API-Token: user\_api\_token  X-CH-Auth-Email: user@domain.suffix  Content-Type: application/json |

> **Notes**:
>
> * The "label\_id" in the path corresponds to the “id” in the device\_label object from the **Device Label List** array.
> * If your organization is registered on Kentik's EU cluster, use api.kentik.eu in place of api.kentik.com in the URL above.

HTTP Response

A successful response from the Device Label Info method includes the following elements:

* Response headers
* HTTP response code
* A single JSON device label object with details about the device label specified by label\_id.

> **Note**: For details of the JSON name/value pairs in a device label object, see **Device Label JSON**.

### Device Label Create

The Device Label Create `POST` method adds a new device label to your organization's collection.

HTTP Request

The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| URL | api.kentik.com/api/v5/deviceLabels |
| Request | POST api/v5/deviceLabels HTTP/1.1  Host: api.kentik.com  X-CH-Auth-API-Token: user\_api\_token  X-CH-Auth-Email: user@domain.suffix  Content-Type: application/json |

> **Note**: If your organization is registered on Kentik's EU cluster, use `api.kentik.eu` in place of `api.kentik.com` in the URL above.

The following parameters are passed in a JSON object in the request body:

| Parameter | Type | Description |
| --- | --- | --- |
| name | string | Required: A unique name for the device label.  - Length: max=100. |
| color | string | Required: A color value expressed in hex (e.g. "#a9c5e0"). |

HTTP Response

A successful response from the Device Label Create method includes the following elements:

* Response headers
* HTTP response code
* A single JSON device label object with details about the newly-added device label.

> **Note**: For details of the JSON name/value pairs in a device label object, see **Device Label JSON**.

### Device Label Update

The Device Label Update `PUT` method modifies data for a specified device label by ID in your organization's collection.

HTTP Request

The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| URL | api.kentik.com/api/v5/deviceLabels/label\_id |
| Request | PUT /api/v5/deviceLabels/label\_id HTTP/1.1  Host: api.kentik.com  X-CH-Auth-API-Token: user\_api\_token  X-CH-Auth-Email: user@domain.suffix  Content-Type: application/json |

> **Notes**:
>
> * The "label\_id" in the path corresponds to the “id” in the `device_label` object from the **Device Label List** array.
> * If your organization is registered on Kentik's EU cluster, use `api.kentik.eu` in place of `api.kentik.com` in the URL above.

The parameters to be changed are passed in a JSON device label object:

* To update the label's name, include the `name` field only.
* To update the label's color, include both the `name` and `color` fields.

Example of a device label object to update the color to dark blue:

```
{
"name": "Edge"
"color": "#000080"
 }
```

HTTP Response

A successful response from the Device Label Update method includes the following elements:

* Response headers
* HTTP response code
* A single JSON device label object with details about the newly-updated device label.

> **Note**: For details on the JSON name/value pairs in a device label object, see **Device Label JSON**.

### Device Label Delete

The Device Label Delete method removes a device label by ID from your organization's collection.

HTTP Request

The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| URL | api.kentik.com/api/v5/deviceLabels/label\_id |
| Request | DELETE /api/v5/deviceLabels/label\_id HTTP/1.1  Host: api.kentik.com  X-CH-Auth-API-Token: user\_api\_token  X-CH-Auth-Email: user@domain.suffix  Content-Type: application/json |

> **Notes**:
>
> * The "label\_id" in the path corresponds to the “id” in the `device_label` object from the **Device Label List** array.
> * If your organization is registered on Kentik's EU cluster, use `api.kentik.eu` in place of `api.kentik.com` in the URL above.

HTTP Response

A successful response from the Device Label Delete method includes the following elements:

* Response headers
* HTTP response code
* A single JSON object confirming the operation completed successfully:

```
{
"success": true
 }
```
