# Admin APIs

Source: https://kb.kentik.com/docs/admin-apis

---

The Kentik V5 Admin APIs enable programmatic management of your organization’s Kentik admin settings.

> **Notes:**
>
> * Member-level users can only access the GET methods of Admin APIs.
> * For an overview of all Kentik APIs, see **Kentik APIs**.
> * For information on using the V5 Admin APIs with cURL, see **API Access Via cURL**.
> * For help using any API version, contact **support@kentik.com**.

## About V5 Admin APIs

The Kentik V5 Admin APIs allow you to manage your organization’s Kentik settings programmatically. They provide methods to read/write backend information related to users, devices, sites, and tags. Information is returned in a JSON structure with key/value pairs within an HTTP response body.  
  
For more details on the data accessible via each API for your organization, refer to the following articles:

* **User API**: Data for each user registered on the Kentik system (see **Users**).
* **Device API**: Data for each device (e.g., router, host) registered on the Kentik system (see **About Devices**).
* **Device Label API**: Manage **Labels** for devices.
* **Plan API**: Access data on “plans”, which are services provided to you by Kentik (see **About Plans**).
* **Site API**: Access data on “sites”, which are groups of devices based on geographic location (see **Site Field Definitions**).
* **Tag API**: Access data on tags (see **Flow Tags**).
* **Custom Dimension API**: Access data on up to 10 custom dimensions and their associated populators (see **Dimensions and Populators**).
* **Saved Filter API**: Access data on custom filters saved by your users, see (**Company Saved Filters**).
* **Alerting APIs**: Manage and use Kentik’s alert-related features (**Manual Mitigation**only).
* **My Kentik API**: Manage the **My Kentik Portal** programmatically.

Kentik V5 Admin APIs can be accessed via:

* **Command line**: Use **cURL** in environments like Terminal (see **API Access Via cURL**).
* **Programmatically**: Send a request body to a V5 API endpoint using any application language that supports HTTP.

> **Note:** The V5 API tester was discontinued in January 2025.

## V5 Admin Methods

The following table lists available methods in the V5 Admin APIs. Click the topic link for more details on each method:

| Method | Endpoint | Description | Topic |
| --- | --- | --- | --- |
| **User API** | | | |
| GET | /users | Get information about all users | **User List** |
| GET | /user/user\_id | Get information about a user | **User Info** |
| POST | /user | Create a new user | **User Create** |
| PUT | /user/user\_id | Update user info | **User Update** |
| DELETE | /user/user\_id | Delete a user | **User Delete** |
| **Device API** | | | |
| GET | /devices | Get information about all devices | **Device List** |
| GET | /device/device\_id | Get information about a device | **Device Info** |
| POST | /device | Create a new device | **Device Create** |
| PUT | /device/device\_id | Update device info | **Device Update** |
| DELETE | /device/device\_id | Delete a device | **Device Delete** |
| GET | device/device\_id/interfaces | Get information about all interfaces on a device | **Interface List** |
| GET | device/device\_id/interface/interface\_id | Get information about an individual interface on a device | **Interface Info** |
| POST | device\_id/interface | Create a new interface on a device | **Interface Create** |
| PUT | device/device\_id/interface/interface\_id | Update interface info | **Interface Update** |
| DELETE | device/device\_id/interface/interface\_id | Delete an interface | **Interface Delete** |
| **Device Label API** | | | |
| GET | /deviceLabels | Get information about all labels | **Device Label List** |
| GET | /deviceLabels/label\_id | Get information about a label | **Device Label Info** |
| POST | /deviceLabels | Create a new label | **Device Label Create** |
| PUT | /deviceLabels/label\_id | Update label info | **Device Label Update** |
| DELETE | /deviceLabels/label\_id | Delete a label | **Device Label Delete** |
| **Plan API** | | | |
| GET | /plans | Get information about plans | **Plan List** |
| **Site API** | | | |
| GET | /sites | Get information about all devices | **Site List** |
| GET | /site/site\_id | Get information about a device | **Site Info** |
| POST | /site | Create a new device | **Site Create** |
| PUT | /site/site\_id | Update device info | **Site Update** |
| DELETE | /site/site\_id | Delete a device | **Site Delete** |
| **Tag API** | | | |
| GET | /tags | Get information about all tags | **Tag List** |
| GET | /tag/tag\_id | Get information about a tag | **Tag Info** |
| POST | /tag | Create a new tag | **Tag Create** |
| PUT | /tag/tag\_id | Update tag info | **Tag Update** |
| DELETE | /tag/tag\_id | Delete a tag | **Tag Delete** |
| **Custom Dimension API** | | | |
| GET | /customdimensions | Get information about all custom dimensions | **Custom Dimension List** |
| GET | /customdimension/dimension\_id | Get information about a custom dimension | **Custom Dimension Info** |
| POST | /customdimension | Create a new custom dimension | **Custom Dimension Create** |
| PUT | /customdimension/dimension\_id | Update custom dimension info | **Custom Dimension Update** |
| DELETE | /customdimension/dimension\_id | Delete a custom dimension | **Custom Dimension Delete** |
| POST | /customdimension/dimension\_id/populator | Create a new populator for a dimension | **Populator Create** |
| PUT | /customdimension/dimension\_id/populator/populator\_id | Update populator info | **Populator Update** |
| DELETE | /customdimension/dimension\_id/populator/populator\_id | Delete a populator | **Populator Delete** |
| **Batch API** | | | |
| POST | /batch/tags  /batch/customdimensions/dimension\_name/populators | Create, update, or delete tags or populators by the batch. | **Batch Request** |
| GET | /batch/batch\_operation\_guid/status | Get status information about a batch operation. | **Batch Status Request** |
| **Saved Filter API** | | | |
| GET | /saved-filters/custom | Get information about all custom saved filters. | **Saved Filter List** |
| GET | /saved-filter/saved-filter/custom/savedfilter\_id | Get information about a custom saved filter. | **Saved Filter Info** |
| POST | /saved-filter/custom | Create a new custom saved filter. | **Saved Filter Create** |
| PUT | /saved-filter/saved-filter/custom/savedfilter\_id | Update custom saved filter info. | **Saved Filter Update** |
| DELETE | /saved-filter/saved-filter/custom/savedfilter\_id | Delete a custom saved filter. | **Saved Filter Delete** |
| **Custom Application API** | | | |
| GET | /customApplications | Get information about all custom applications. | **Custom Application API** |
| POST | /customApplications | Create a new custom application. | **Custom Application API** |
| PUT | /customApplications/application\_id | Update custom application info. | **Custom Application API** |
| DELETE | /customApplications/application\_id | Delete a custom application. | **Custom Application API** |
| **Manual Mitigation API** | | | |
| POST | alerts/manual-mitigate | Start a new manual mitigation | **Manual Mitigation Create** |
| **My Kentik API** | | | |
| GET | /mykentik/tenants | Get information about all tenants. | **Tenant List** |
| GET | /mykentik/tenant/tenant\_id | Get information about an individual tenant. | **Tenant Info** |
| POST | /mykentik/tenant/tenant\_id/user | Create a new tenant user. | **Tenant User Create** |
| DELETE | /mykentik/tenant/tenant\_id/user/user\_id | Delete a tenant  user. | **Tenant User Delete** |

> **Note:** Endpoint URLs are region-specific:
>
> * US: `https://api.kentik.com/api/v5`
> * EU: `https://api.kentik.eu/api/v5`
