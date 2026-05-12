# Custom DNS

Source: https://kb.kentik.com/docs/custom-dns

---

This article covers the **Custom DNS** settings page of the Kentik portal.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(124).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A00Z&se=2026-05-12T09%3A37%3A00Z&sr=c&sp=r&sig=Ec6DKUm5jWHpuuFYRqdxYyaXeo3R62nXRrCmevOHbz4%3D)

*Custom DNS lets you choose the servers for reverse DNS lookup in Data Explorer queries.*

## About Custom DNS

Custom DNS allows you to specify a DNS server for Kentik to use for reverse DNS (rDNS) lookups during **Data Explorer** queries.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/CDns-Explorer_IP_column-218h359w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A00Z&se=2026-05-12T09%3A37%3A00Z&sr=c&sp=r&sig=Ec6DKUm5jWHpuuFYRqdxYyaXeo3R62nXRrCmevOHbz4%3D)By default, Kentik can only resolve names for public IP addresses. For private IP addresses (RFC 1918), only a hyphen is displayed. By configuring a Custom DNS server, you can resolve these private addresses directly within your query results.

* **Behavior**: When the **Enable Reverse DNS Lookups** switch is toggled on in the Advanced Query Settings, hostnames appear in parentheses after the IP (e.g., 123.45.210.12/32 (www.host\_name.com)).
* **Data Impact**: Custom DNS only affects the query side. It does not change the data stored in the Kentik Data Engine (KDE). The hostname displayed will always reflect the **current** value returned by the DNS server at the time of the query.

> **Notes:**
>
> * To enable parallel lookups, you may have multiple Custom DNS server(s) registered simultaneously.
> * Your Custom DNS server(s) must be available to query from Kentik. To ensure secure access, the IP address used by Kentik to connect to the server is shown on your Custom DNS page in the Kentik portal.
> * Any delay in reaching a Custom DNS server will add to query response time. If the response of a Custom DNS server is sufficiently delayed (e.g., in case of "Internet weather"), timeouts may result.
> * For “permanent” DNS enrichment options, see **Persistent DNS Enrichment via Universal Agent**.

### Persistent DNS Enrichment via Universal Agent

If you need to track how an IP’s hostname changes over time, or if you want hostnames stored permanently within your flow records, you can use the **rDNS flow enrichment** option of the Flow Proxy capability of Kentik’s **Universal Agent**.

This performs rDNS lookups on source and destination IPs during ingest. The resulting names are stored in the Source Hostname and Destination Hostname dimensions, allowing for historical reporting on name changes.

### Comparison of Custom DNS vs. Flow Proxy rDNS

The following table compares the features of using Custom DNS versus persistent DNS enrichment:

| Feature | Custom DNS | Flow Proxy rDNS |
| --- | --- | --- |
| Timing | Query-time (on-the-fly) | Ingest-time (stored) |
| Persistence | Not stored in KDE | Stored in flow records |
| Use Case | Quick lookups for current names | Historical tracking of hostnames |
| Dimensions | IP/CIDR (in parentheses) | Source/Destination Hostname |

## Custom DNS Page

The **Custom DNS** page, accessed via the **Organization Settings** menu, is used to specify the IP address of one or more alternate DNS servers to be used for reverse DNS lookup instead of the default servers used for this purpose by Kentik. While Kentik users whose level is Member can view the content of the Custom DNS page, only Administrators can add, edit, verify, or remove a Custom DNS server.  
  
The Custom DNS page is made up of the following UI elements:

* **Add Custom DNS Server**: A button that opens the **Custom DNS** dialog (see **Custom DNS Dialogs**).
* **Verify Reverse DNS Lookup**: A button that opens the **Verify Reverse DNS Lookup** dialog (see **Verify Reverse DNS Lookup**).
* **Info field** (indicated by info icon): Provides an explanation of Custom DNS as well as the IP address from which Kentik will query the Custom DNS servers added to the **DNS Servers** list. The listed servers must allow reverse DNS lookup from this IP.
* **Filter field**: Enter text to filter the Custom DNS list. The list will show only servers for which there is a match in the **Name** and **IP Address** columns for the string entered in this field.
* **Custom DNS Servers list**: A list of the IP addresses of servers to use for reverse DNS lookups. May contain multiple servers. See **Custom DNS List**.

### Custom DNS List

The Custom DNS list is a table that lists the custom DNS servers configured in Kentik for your organization. By default, the list is ordered alphabetically by name. To change the sort order of the list, click a heading to select a column, and click the resulting blue up or down arrow to choose the sort direction (ascending or descending).

#### Custom DNS Info Columns

The **Custom DNS** list includes the following columns (left to right):

* **Name**: The name of the Custom DNS server.
* **IP Address**: The IP address of the Custom DNS server.

#### Custom DNS Admin Actions

The following actions (far-right columns) are available only to admin-level users:

* **Edit Custom DNS**: Clicking the pencil icon opens the Custom DNS dialog that allows you to review and edit the Custom DNS server’s information. See **Custom DNS Dialogs**.
* **Remove**: Removes the Custom DNS server from the list. There is no confirming dialog box before removal; when you click the trash can icon, the server is removed from the list.

## Custom DNS Dialogs

The custom DNS dialogs (**Add Custom DNS** and **Edit Custom DNS**) are used to create or modify a Custom DNS for your account. The dialogs are made up of the following UI elements:![Interface for adding custom DNS with fields for name and IP address.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/CDNS-add-custom-dns-dialog.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A00Z&se=2026-05-12T09%3A37%3A00Z&sr=c&sp=r&sig=Ec6DKUm5jWHpuuFYRqdxYyaXeo3R62nXRrCmevOHbz4%3D)

* **Close button**: Click the **X** in the upper right corner to close the dialog. All elements will be restored to their values at the time the dialog was opened.
* **Name field**: Enter a DNS server alias to label the DNS Server within the Kentik system.
* **IP Address field**: Enter an IP address upon which to perform a reverse DNS lookup using the Custom DNS servers in the DNS Servers list.
* **Cancel button**: Cancel the add DNS operation and exit the dialog. All elements will be restored to their values at the time the dialog was opened.
* **Add Custom DNS button** (Add dialog only): Add the new DNS Server to the list of Custom DNS servers.
* **Save** (Edit dialog only): Save any changes you made to the existing DNS Server.

> **Note:** For more information on using the dialogs, see **Add a Custom DNS Server** and **Edit a Custom DNS Server**.

## Verify Reverse DNS Lookup

Use the **Verify Reverse DNS Lookup** dialog to verify the IP address for a DNS server and confirm the hostname of the DNS service provider at that address.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/CDns-Verify_DNS_dialog-330h674w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A00Z&se=2026-05-12T09%3A37%3A00Z&sr=c&sp=r&sig=Ec6DKUm5jWHpuuFYRqdxYyaXeo3R62nXRrCmevOHbz4%3D)

The dialog is made up of the following UI elements:

* **Close button**: Click the **X** in the upper right corner to close the dialog.
* **IP Address field**: Enter the IP address you want to verify.
* **Resolve button**: Enabled when an IP address has been entered in the **Add DNS IP** field. Click to resolve the IP address to a host name.
* **Result field**: Appears after a successful lookup.

For more information on using the dialog, see **Verify a Custom DNS Server**.

## Configuring Custom DNS

Configuration of a Custom DNS on the Custom DNS page (Admin » **Custom DNS**) is covered here.

### Add a Custom DNS Server

To add a Custom DNS server:

1. Open the Custom DNS page (Settings » **Custom DNS**).
2. Click the **Add Custom DNS Server** button to open the Add Custom DNS dialog.
3. Enter an alias for the DNS Server in the **Name** field.
4. Enter the IP address of a DNS server into the **IP Address** field.

   1. If valid, the **Add Custom DNS** button is enabled.
   2. If the IP address is not valid, a notification appears under the field asking you to enter a valid IP address.
5. Click the **Add Custom DNS** button. The Custom DNS is added to the list.

### Edit a Custom DNS Server

To edit an existing Custom DNS server:

1. Open the Custom DNS page (Settings » **Custom DNS**).
2. In the DNS Server list, find the row corresponding to the DNS Server you’d like to edit, then click the Edit button at the right of that row.
3. Modify the alias of the DNS Server as needed.
4. Modify the IP Address as needed.

   1. If valid, the **Save** button is enabled.
   2. If the IP address is not valid, a notification appears under the field asking you to enter a valid IP address.
5. Click the **Save** button. The Custom DNS list should reflect the changes immediately.

### Remove a Custom DNS Server

To remove an existing Custom DNS server:

1. Open the Custom DNS page (Settings » **Custom DNS**).
2. In the DNS Server list, find the row corresponding to the DNS Server you’d like to remove, then click the Remove button (trash icon) at the right of that row.
3. In the resulting confirmation dialog, click **Cancel** to abort the server removal or Remove to proceed with removal of the server.
4. The Custom DNS list should reflect the change immediately.

### Verify a Custom DNS Server

To verify that an IP address works for reverse DNS lookup:

1. Open the Custom DNS page (Settings » **Custom DNS**).
2. Click the **Verify Reverse DNS Lookup** button.
3. In the **IP Address** field, enter an IP address. The **Resolve** button at the right of the field is enabled.
4. Click the **Resolve** button:

   1. If reverse DNS lookup is successful a "Success" notification appears in the dialog.
   2. If you enter a valid IP but the reverse DNS lookup is not successful, a "Fail" notification appears in the dialog.
   3. If you enter an invalid IP, an "Invalid IP Address" notification appears at the top of the page.
