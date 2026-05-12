# User API

Source: https://kb.kentik.com/docs/user-api

---

The user API, part of the Kentik V5 Admin APIs, is covered here.

> Notes:
>
> * For an overview of all Kentik APIs, see **Kentik APIs**.
> * For information on using this API with cURL, see **API Access Via cURL**.
> * The V5 API tester was discontinued in January 2025.
> * For help with any API version, contact **support@kentik.com**.

## About the User API

The Kentik V5 Admin APIs offer an API for managing user settings programmatically, similar to the Users page (see **Users**) under the Organization menu of the Kentik portal.

> **Note:** Member-level users only have access to the GET methods of Admin APIs.

## User JSON

User API calls return an HTTP response with a JSON "user" object, or an array of such objects for the User List call. The object contains the fields (name/value pairs) shown in the following example:

```
{
"user": {
     "id": "#####",
     "username": "newuser",
     "user_full_name": "New Username",
     "user_email": "new_user@kentik.com",
     "role": "Member"
     "email_service": true,
     "email_product": true,
     "last_login": "2016-09-22T21:26:43.279Z",
     "created_date": "2016-09-22T21:26:43.279Z",
     "updated_date": "2016-09-22T21:26:43.279Z",
     "company_id": "####",
   }
 }
```

Each user object contains fields with information about an individual user registered in your organization. These fields are described in the following table:

| JSON name | Type | Description |
| --- | --- | --- |
| id | number | The system-assigned user id), e.g., “1111”.  **Note:** This field is the primary key for the table. |
| user\_name | string | The username provided at signup, e.g., "janedoe". Valid characters: alphanumeric plus underscores.  Length: min=3, max=40. |
| user\_full\_name | string | The user's full name (up to 128 characters) as provided at signup (e.g. "Jane Doe"). Valid characters: all except double quotes.  Length: max=50. |
| user\_email | string | The user email address provided at signup, e.g., "name@domain.suffix". |
| role | string | The role of the user: Member (0) or Administrator (1). |
| email\_service | boolean | Opt-in for service emails. |
| email\_product | boolean | Opt-in for product emails. |
| last\_login | string | UTC time of most recent login, e.g., 2015-01-27T00:32:34.559Z. |
| created\_date | string | Date-time of user creation in UTC (ISO 8601), e.g., 2015-01-27T01:39:17.186Z. |
| updated\_date | string | Date-time of last user edit in UTC, e.g., 2015-01-27T01:39:17.186Z. |
| company\_id | number | The system-assigned company id, e.g., “9999”. |

> **Note:** For details on the settings represented by the name/value pairs, see **General Settings**.

## User List

The User List `GET` method fetches a JSON array of your Kentik users, with each element  representing an individual user.

> **Note:** The "id" value in each user object can be used in subsequent calls to retrieve, update, or delete that user.

**HTTP Request**  
 The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| **URL** | https://api.kentik.com/api/v5/users |
| **Request** | GET /api/v5/users HTTP/1.1  Host: api.kentik.com  X-CH-Auth-API-Token: user\_api\_token  X-CH-Auth-Email: user@domain.suffix  Content-Type: application/json |

> **Note:** If your organization is registered on Kentik's EU cluster, use `api.kentik.eu` in place of `api.kentik.com` in the URL above.

**HTTP Response**  
 A successful response from the User List method includes the following elements:

* Response headers
* HTTP response code
* A response body containing a JSON array with `user` objects, each containing information about a user.

> **Note:** For details on the JSON name/value pairs in a `user` object, see **User JSON**.

## User Info

The User Info `GET` method fetches details of a single user by ID from your list of Kentik users.  
  
**HTTP Request**  
The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| **URL** | https://api.kentik.com/api/v5/user/user\_id |
| **Request** | GET /api/v5/user/user\_id HTTP/1.1  Host: api.kentik.com  X-CH-Auth-API-Token: user\_api\_token  X-CH-Auth-Email: user@domain.suffix  Content-Type: application/json |

> **Notes**:
>
> * The "user\_id" value in the path corresponds to the "id" value in each `user` object from the **User List** array.
> * If your organization is registered on Kentik's EU cluster, use `api.kentik.eu` in place of `api.kentik.com` in the URL above.

**HTTP Response**  
A successful response from the User Info method includes the following elements:

* Response headers
* HTTP response code
* A single JSON `user` object with details about the user specified by `user_id`.

> **Note:** For details of the JSON name/value pairs in a `user` object, see **User JSON**.

## User Create

The User Create POST method adds a new user to your organization in Kentik.  
  
**HTTP Request**  
 The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| **URL** | https://api.kentik.com/api/v5/user |
| **Request** | POST /api/v5/user HTTP/1.1  Host: api.kentik.com  X-CH-Auth-API-Token: user\_api\_token  X-CH-Auth-Email: user@domain.suffix  Content-Type: application/json |

> **Note:** If your organization is registered on Kentik's EU cluster, `use api.kentik.eu` in place of `api.kentik.com` in the URL above.

The following parameters are passed in the request body:

| Parameter | Type | Description |
| --- | --- | --- |
| user\_name | string | Required: A user name (no spaces). |
| user\_full\_name | string | Required: A full name for this user. |
| user\_email | string | Required: The user's email address. |
| user\_password | string | Required: The user's password; must be at least seven characters. |
| role | string | Required: The user's role: Member (0) or Administrator (1). |
| email\_service | boolean | Required: Opt-in for service emails. |
| email\_product | boolean | Required: Opt-in for product emails. |

**HTTP Response**  
A successful response from the User Create method includes the following elements:

* Response headers
* HTTP response code
* A single JSON user object with details about the newly-added user.

> **Note:** For details of the JSON name/value pairs in a user object, see **User JSON**.

## User Update

The User Update `PUT` method modifies details of a specific user by ID in your collection of existing users.  
  
**HTTP Request**  
The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| **URL** | https://api.kentik.com/api/v5/user/user\_id |
| **Request** | PUT /api/v5/user/user\_id HTTP/1.1  Host: api.kentik.com  X-CH-Auth-API-Token: user\_api\_token  X-CH-Auth-Email: user@domain.suffix  Content-Type: application/json |

> **Notes:**
>
> * The "user\_id" value in the path corresponds to the "id" value in each user object from the **User List** array.
> * If your organization is registered on Kentik's EU cluster, use api.kentik.eu in place of api.kentik.com in the URL above.

Parameters to be updated are included in a JSON user object with only the fields to be changed. The following example updates the user’s service email preference:

```
{
"user": {
     "email_service": true
   }
 }
```

**HTTP Response**  
 A successful response from the User Update method includes the following elements:

* Response headers
* HTTP response code
* A single JSON `user` object with details about the newly-updated user.

> **Note:** For details of the JSON name/value pairs in a `user` object, see **User JSON**.

## User Delete

The User Delete method removes a specific user by ID from your collection of Kentik users.  
  
**HTTP Request**  
The following table shows the path and HTTP request for this call:

|  |  |
| --- | --- |
| **URL** | https://api.kentik.com/api/v5/user/user\_id |
| **Request** | DELETE /api/v5/user/user\_id HTTP/1.1  Host: api.kentik.com  X-CH-Auth-API-Token: user\_api\_token  X-CH-Auth-Email: user@domain.suffix  Content-Type: application/json |

> **Notes**:
>
> * The "user\_id" value in the path corresponds to the "id" value in each user object from the **User List** array.
> * If your organization is registered on Kentik's EU cluster, use api.kentik.eu in place of api.kentik.com in the URL above.

**HTTP Response**  
 A successful response from the User Delete method includes the following elements:

* Response headers

HTTP response code 204 (no content)
