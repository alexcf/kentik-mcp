# My Kentik API

Source: https://kb.kentik.com/docs/my-kentik-api

---

The My Kentik API enables programmatic management of the My Kentik Portal.

> **Notes:**
>
> * Member-level users only have access to the GET methods of the My Kentik API.
> * For an overview of all Kentik APIs, see **Kentik APIs**.
> * For information on using this API with cURL, see **API Access Via cURL**.
> * For assistance using any API version, please contact support@kentik.com.

## About My Kentik API

The Kentik V5 Admin APIs include an API for managing tenant settings (see **About Tenancy**) in the My Kentik Portal. The API offers most of the functionality of the My Kentik Portal page, which is listed under Customize in the Admin sidebar of the portal.

> **Notes:**
>
> * To add, update, or delete a tenant, you must use the portal UI (see **My Kentik Portal Page**).
> * The V5 API tester was discontinued in January 2025.

## Tenant JSON

Calls to the My Kentik API return an HTTP response body containing a JSON `tenant` object, or an array of such objects for the Tenant List call. The object contains the fields (name/value pairs) shown in the example below (placeholders in italics).

```
{
  "id": ###,
  "name": "tenant_name",
  "description": "Everything you might wish to know about this tenant.",
  "users": [
    {
      "id": "user_id",
      "user_email": "user@domain.suffix",
      "last_login": "2019-03-01T17:48:51.113Z",
      "tenant_id": "###",
      "company_id": "####"
    }
  ],
  "created_date": "2018-10-11T22:52:05.833Z",
  "updated_date": "2019-01-11T16:23:17.267Z"
}
```

Each `tenant` object contains fields with information about a single tenant in your organization's collection. These fields are described in the tables below.

> **Note:** For details on the settings represented by the name/value pairs below, see **About Tenant Settings**.

#### Tenant Object

The tenant object is the top-level object returned from calls to the My Kentik API.

| JSON name | Type | Description |
| --- | --- | --- |
| id | number | Kentik-assigned ID of the tenant. |
| company\_id | string | Kentik-assigned ID of the customer with which the tenant is associated. |
| name | string | Customer-assigned tenant name. |
| description | string | Description of the tenant. |
| cdate | string | Date-time of tenant creation, in UTC (ISO 8601), e.g., 2025-01-27T01:39:17.186Z |
| edate | string | Date-time tenant was most recently edited. |
| users | array | See **Users Array**. |

#### Users Array

The `users` array consists of objects, each representing a user assigned to this tenant.

| JSON name | Type | Description |
| --- | --- | --- |
| id | string | The system-assigned id of the user. |
| company\_id | string | The system-assigned id of the company. |
| user\_email | string | The user’s email address as provided at signup (e.g. “user@domain.suffix”). |
| user\_name | string | The name provided for the user by admin at signup (e.g. “janedoe”). Valid characters: alphanumeric plus underscores.- Length: min=3, max=40. |
| user\_full\_name | string | The user’s full name (up to 128 characters) as provided at signup (e.g. “Jane Doe”).- Valid characters: all except double quotes.- Length: max=50. |
| tenant\_id | string | Kentik-assigned ID of the tenant. |
| last\_login | string | Date-time of most recent tenant login. |

## Tenant List

The Tenant List `GET` method retrieves a JSON array of `tenant` objects, each representing an individual tenant in your organization.

> **Note:** The `id` value in each `tenant` object can be used in subsequent calls to reference that tenant.

**HTTP Request**  
The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| **URL** | api.kentik.com/api/v5/mykentik/tenants |
| **Request** | GET /api/v5/mykentik/tenants HTTP/1.1  Host: api.kentik.com  X-CH-Auth-API-Token: user\_api\_token  X-CH-Auth-Email: user@domain.suffix  Content-Type: application/json |

> **Note:** If your organization is registered on Kentik's EU cluster, use `api.kentik.eu` in place of `api.kentik.com` in the URL above.

**HTTP Response**  
A successful response from the Tenant List method includes the following elements:

* Response headers
* HTTP response code
* A response body with a JSON `tenants` array, where each element is a `tenant` object containing information about one of your organization’s tenants.

> **Note:** For details of the JSON name/value pairs in a `tenant` object, see **Tenant JSON**.

## Tenant Info

The Tenant Info `GET` method retrieves information about a single tenant, identified by ID, in your organization's collection of tenants.  
  
**HTTP Request**  
The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| **URL** | api.kentik.com/api/v5/mykentik/tenant/tenant\_id |
| **Request** | GET /api/v5/mykentik/tenant/tenant\_id HTTP/1.1 Host: api.kentik.com X-CH-Auth-API-Token: user\_api\_token X-CH-Auth-Email: user@domain.suffix Content-Type: application/json |

> **Notes**:
>
> * The "tenant\_id" in the path corresponds to the `id` of a `tenant` object in the **Tenant List** array.
> * If your organization is registered on Kentik's EU cluster, use `api.kentik.eu` in place of `api.kentik.com` in the URL above.

**HTTP Response**

A successful response from the Tenant Info method includes the following elements:

* Response headers
* HTTP response code
* A JSON `tenant` object containing information about the tenant specified by `tenant_id`.

> **Note:** For details of the JSON name/value pairs in a `tenant` object, see **Tenant JSON**.

## Tenant User Create

The Tenant User Create `POST` method assigns a new user, by email address, to one of your organization's tenants. An email is sent to the provided address, allowing the new user to verify the email, sign in, create a password, and activate the account.

> **Note:** To update a tenant user’s properties, first delete the user and then recreate the user with the corrected properties.

**HTTP Request**  
The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| **URL** | api.kentik.com/api/v5/mykentik/tenant/tenant\_id/user |
| **Request** | POST /api/v5/mykentik/tenant/tenant\_id/user HTTP/1.1 Host: api.kentik.com X-CH-Auth-API-Token: user\_api\_token X-CH-Auth-Email: user@domain.suffix Content-Type: application/json |

> **Note:** If your organization is registered on Kentik's EU cluster, use `api.kentik.eu` in place of `api.kentik.com` in the URL above.

The following parameters are passed in the request body:

| Parameter | Type | Description |
| --- | --- | --- |
| user\_email | string | Required: A valid email address for the new tenant user. |

**HTTP Response**  
A successful response from the Tenant User Create method includes the following elements:

* Response headers
* HTTP response code
* A JSON `user` object containing information about the newly-added user.

> **Note:** For details of the JSON name/value pairs in a `user` object, see **Users Array**.

## Tenant User Delete

The Tenant User Delete method removes a user, identified by ID, from the collection of users assigned to a tenant.  
  
**HTTP Request**  
The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| **URL** | api.kentik.com/api/v5/mykentik/tenant/tenant\_id/user/user\_id |
| **Request** | DELETE /api/v5/user/user\_id HTTP/1.1 Host: api.kentik.com X-CH-Auth-API-Token: user\_api\_token X-CH-Auth-Email: user@domain.suffix Content-Type: application/json |

> **Notes:**
>
> * The "tenant\_id" in the path corresponds to the `id` in a `tenant` object (see **Tenant Info** or **Tenant List**).
> * The "user\_id" in the path corresponds to the `id` from a `user` object in the **Users Array**.
> * If your organization is registered on Kentik's EU cluster, use `api.kentik.eu` in place of `api.kentik.com` in the URL above.

**HTTP Response**  
A successful response from the Tenant User Delete method includes the following elements:

* Response headers
* HTTP response code 204 (no content)

---

© 2014-25 Kentik
