# Kentik APIs

Source: https://kb.kentik.com/docs/apis-overview

---

This article provides an introduction to Kentik APIs, which enable programmatic control of Kentik.

> **Note:** For assistance using a Kentik API, see **Customer Care**.

## About Kentik APIs

Kentik APIs enable programmatic control over Kentik setup (administrative functions), querying of network data stored in Kentik, and alerting/mitigations based on user-defined conditions.

### API Versions

Kentik supports two sets of APIs:

* **V5**: A set of REST APIs for administrative functions, querying, and alerting (see **About the V5 APIs**).
* **V6**: A set of gRPC APIs that do not directly extend V5 but do have some overlap in functionality (e.g., Users and Sites). V6 APIs are more frequently updated and are useful for high throughput operations due to their increased efficiency compared to REST (see **About the V6 APIs**).

> **Notes:**
>
> * The V5 API tester was discontinued in January 2025.
> * We do not recommend using Kentik APIs for full data extraction. Instead, Kentik Firehose is recommended for that use case (see **Using Kentik Firehose**).

## About the V6 APIs

Kentik’s V6 (gRPC) APIs aim to:

* Offer a more performant alternative to the V5 (REST) APIs
* Offer functionality not covered by the V5 APIs
* Serve as a foundation for expanding capabilities, including areas already covered by the V5 APIs (e.g., Users and Sites).

For a list of available v6 APIs, see **V6 APIs**.

> **Note:** KB articles cover each API listed in the v6 API Tester’s landing page. APIs marked in the KB as Alpha or Beta versions are not supported or recommended for production use.

### V6 API Tester

The **V6 API Tester** allows you to test V6 API methods using your Kentik account data.

> **Note**: All V6 API reference files are also available in Kentik’s **api-schema-public**repo on GitHub**.**

### V6 API Client Stubs

Kentik provides API client stubs in the following languages in its **api-schema-public**repo on GitHub:

* **C**
* **C++**
* **Go**

## About the V5 APIs

The Kentik V5 (REST) APIs enable programmatic management of Kentik and are categorized as follows:

* **V5 Admin APIs**: Manage your company's Kentik settings, including users, devices, sites, and tags. These APIs support methods to read/write backend information, returning data in JSON key/value pairs. For more details, refer to the **V5 Admin APIs**documentation and the list of **V5 Admin Methods**.
* **V5 Query API**: Enables querying of your network data stored in Kentik, returning data or visualizations (see the **V5 Query API** documentation and the list of **V5 Query Methods**).
* **V5 Alerting APIs**: Provides notifications of anomalous conditions (DDoS attacks, blocked routes, etc.) on Kentik-monitored network infrastructure, and supports mitigation management. Returns data on current or historical deviations. See the **V5 Alerting API** documentation and the list of **V5 Alerting Methods**.

> **Note:** The V5 API tester was discontinued in January 2025.

## API Rate Limiting

API calls are rate-limited to ensure fair resource distribution among Kentik customers. The limits allow for reasonable activity bursts and prevent abrupt failures. Key points to consider when developing with Kentik APIs:

* **Separate Counting**: Requests from the **V5 Query API** and **V5 Admin APIs** are counted separately, each with distinct limits.
* **Rolling Time Windows**:Rate limits are calculated over a rolling minute and hour, with each request counted for 60 seconds and 3600 seconds, respectively, from receipt.
* **Server Response Delay:** Applied to requests exceeding the per-minute soft limit (see table below).
* **HTTP 429 Error**: Triggered when requests exceed the per-minute hard limit or the hourly limit.
* **Concurrent Request Limit**: Adhering to this ensures the per-minute count does not exceed the hard limit, and the hourly limit is not exceeded for non-query APIs.

The following table shows the limits for query and non-query API calls.

| Per-customer limit | Non-query APIs | Query API | Description |
| --- | --- | --- | --- |
| Maximum concurrent requests | 1 | 4 | A request is deemed concurrent if received before the server responds to the previous request. |
| Soft limit (per minute) | 20 | 30 | Maximum requests per rolling minute without triggering server response delay. |
| Server response delay (seconds) | 1.5 | 1 | Delay in server response for requests exceeding the soft limit per minute count. |
| Hard limit (per minute) | 60 | 100 | Maximum requests per rolling minute without triggering HTTP 429 error. *Note:* Occurs only if concurrent requests exceed maximum. |
| Hourly limit | 3750 | 1500 | Maximum requests per rolling hour without triggering HTTP 429 error. |

## API Access Via cURL

To access Kentik APIs using **cURL**, you can generate HTTP requests to the Kentik API server.

### API V6 cURL Example

```
curl --location --request GET 'https://grpc.api.kentik.com/device/v202504beta2/device/' \
-H 'X-CH-Auth-Email: <user@domain.suffix>' \
-H 'X-CH-Auth-API-Token: <user API token>'
```

### API V5 cURL Structure

Here’s the structure of an HTTP request to Kentik via cURL for the V5 User API:

| Element | Example | Description |
| --- | --- | --- |
| Command declaration | curl | Declares the following string as a cURL command. |
| Method | -X PUT | HTTP method of the call, e.g.,  GET, PUT, POST, DELETE |
| Authorization: API token | -H 'X-CH-Auth-API-Token: user\_api\_token' | Header: passes the user’s Kentik API Token |
| Authorization: User email | -H 'X-CH-Auth-Email: user@domain.suffix' | Header: passes the user’s Kentik email address |
| Content-Type declaration | Content-Type: application/json | Header: specifies the content type of the data. |
| Data | -d '{   "user": {     "email\_service": true,   } }' | JSON: For creating or updating:  - Create (POST): Include all required fields.  - Update (PUT): Include only fields to be updated; others remain unchanged. For example, updating “email\_service” will leave other fields as is. |
| URL (US) | https://api.kentik.com/api/v5/user/user\_id | Endpoint path on the Kentik V5 API query server. |
| Formatter | | python -m json.tool | Optional: appends to cURL command for "pretty" JSON response layout. |

> **Note:** For organizations registered on Kentik's EU cluster, use `api.kentik.eu` in place of `api.kentik.com` in the API v5 URL.

### API V5 cURL Command

The structure elements can be used to form a cURL command, for example to create a new user (the backslashes are for line breaks for readability):

```
curl -X POST \
 -H 'X-CH-Auth-Email: <user@domain.suffix>' \
 -H 'X-CH-Auth-API-Token: <user_api_token>' \
 -H 'Content-Type: application/json' \
 -d '{
"user": {
     "user_name": "New_Username",
     "user_full_name": "New Username",
     "user_email": "new_user@domain.suffix",
     "user_password": "newUser_password",
     "role": "Member",
     "email_service": true,
     "email_product": true
   }
 }' \
 https://api.kentik.com/api/v5/user | python -m json.tool
```

> **Notes:**
>
> * Find a user's API token on the User Information pane of the User Profile page in Kentik (accessible via the username link in the main navbar).
> * For organizations registered on Kentik's EU cluster, use `api.kentik.eu` in place of `api.kentik.com` in the API v5 URL.

Using a single data (`-d`) parameter in cURL to pass a JSON object simplifies passing field values to Kentik:

* **Create Call**: The `-d` parameter includes the entire JSON object for the API (e.g., the user object for the User API).
* **Update Call**: The `-d` parameter includes only the fields to be updated in the JSON structure for the API.

> **Notes:**
>
> * In some contexts, a `-d` parameter with single quotes may cause a call to fail. Use the Unicode equivalent `\u0027` to escape them in cURL.
> * Examples of JSON objects returned by V5 APIs are shown in **V5 Admin APIs** and **V5 Query API**.
> * For details on dimensions and corresponding JSON name/value pairs in V5 APIs, see **Dimensions Reference**.
