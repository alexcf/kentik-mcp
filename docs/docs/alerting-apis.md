# Alerting APIs

Source: https://kb.kentik.com/docs/alerting-apis

---

Alerting APIs enable programmatic management and use of the alerting features of Kentik.

> Notes:
>
> * For an overview of all Kentik APIs, see **Kentik APIs**.
> * For information on using APIs with cURL, see **API Access Via cURL**.
> * For assistance using any API version, see **Customer Care**.

## V5 Alerting Methods

The following table shows the methods that are available in the V5 Alerting APIs (click on the topic link for more information about a specific method):

| Method | Endpoint | Description | Topic |
| --- | --- | --- | --- |
| ***Manual Mitigation API*** | | | |
| POST | /alerts/manual-mitigate | Start a new manual mitigation | **Manual Mitigation Create** |

> **Note:** Endpoints shown above are relative to a path that depends on the region for which your organization is registered with Kentik:
>
> * US: `https://api.kentik.com/api/v5`
> * EU: `https://api.kentik.eu/api/v5`

## Manual Mitigation API

The Manual Mitigation API enables you to programmatically apply a manual mitigation, which functions the same as those started from the Alerting page in the Kentik portal (see **Manual Mitigation**).

> **Notes:**
>
> * Member-level users have no access to the methods of this API.
> * To make calls to this API using cURL, see **API Access Via cURL**.
> * The V5 API tester was discontinued in January 2025.

### Manual Mitigation JSON

Calls to the Manual Mitigation API return an HTTP response body with a JSON response object, containing a result field as shown in the following example:

```
{
  "response": {
    "result": "OK",
  }
}
```

### Manual Mitigation Create

The Manual Mitigation Create `POST` method initiates a new manual mitigation for a specified IP/CIDR using a chosen mitigation platform and method.  
  
**HTTP Request**  
The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| **URL** | https://api.kentik.com/api/v5/alerts/manual-mitigate |
| **Request** | POST /api/v5/alerts/manual-mitigate HTTP/1.1  Host: api.kentik.com  X-CH-Auth-API-Token: user\_api\_token  X-CH-Auth-Email: user@domain.suffix  Content-Type: application/json |

> **Note:** If your organization is registered on Kentik’s EU cluster, use `api.kentik.eu` in place of `api.kentik.com` in the URL above.

The following parameters are passed in a JSON object in the request body:

| JSON name | Type | Description |
| --- | --- | --- |
| ipCidr | string | The IP/CIDR to which you want to apply the mitigation. |
| Comment | string | A user-defined string describing the mitigation to others in your organization. (Reserved for future use.) |
| platformID | string | A valid ID for a Kentik mitigation platform existing in your organization. |
| methodId | string | A valid ID for a Kentik mitigation method existing in your organization. |
| minutesBeforeAutoStop | string | A duration after which the mitigation will stop. If specified as “0” the mitigation will continue until stopped manually in the portal (see **Stop a Manual Mitigation**). |

**HTTP Response**  
A successful response to this call includes the following elements:

* Response headers
* HTTP response code
* A JSON response object indicating that the call was successful

> **Note:** For details of the JSON name/value pairs in a response object, see **Manual Mitigation JSON**.
