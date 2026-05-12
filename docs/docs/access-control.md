# Access Control

Source: https://kb.kentik.com/docs/access-control

---

The **Access Control** settings page helps you prevent unauthorized access to your Kentik portal.

![Access control settings with warnings about security for App, Agent, and API modes.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/ACL-main.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A09Z&se=2026-05-12T09%3A33%3A09Z&sr=c&sp=r&sig=sa%2B3hu1dMqqAFfmTJdP%2BVI9GYjsZwZdT1I6I%2F9g5sMs%3D)

*The Access Control page enables you to choose the IPs from which it's possible to access various Kentik functions.*

## About Access Control

Access control enhances security by enabling you to specify IPs and subnets that are granted access to Kentik. Access is denied to all other IPs and subnets that attempt to connect. For greater flexibility, access settings may be set individually for each of the following Kentik subsystems:

* **App**: Access via the Kentik portal.
* **Agent**: Access via the Flow Proxy capability of Kentik's **Universal Agent**.
* **API**: Access via Kentik APIs (see **Kentik APIs**).

  > **Note:** Database (psql) access to KDE is deprecated. For additional information, contact Kentik support.

The access control settings for all four subsystems are set on the **Access Control Page**.  

## About Exception Groups

Exception groups enable you to specify, for any of the subsystems listed on the **Access Control Page** (App, Agent, API, and Database), the set of IPs and subnets from which access to Kentik is allowed. These groups allow you to group certain subnets and IPs into functional, easily identifiable, and more usable groups. These groups are particularly useful when allowing access for your remote employees.

> **Note:** When you create an exception group, only the IP address(es) or subnet(s) listed will be able to access that subsystem. Be sure to add an exception for your own role if you need one.

## Access Control Page

To open the Access Control page, choose Access Control from the **Organization Settings** menu.

> **Note:** Member-level Kentik users can open the Access Control page, but only admin-level users can modify the settings.

### Access Control UI

The Access Control page is home to the access control settings for the four Kentik subsystems listed in **About Access Control**.  
  
Each subsystem is represented by a pane that contains the following UI elements:

* **Access setting**: A set of radio buttons that enable you to choose one of the following access control options for the subsystem:

  + **Allow All**: No restriction; the subsystem can be accessed from any IP or subnet.
  + **Deny All except**: The subsystem can be accessed only from the IP addresses and subnets listed in the IP Addresses/Subnets field.
* **Exception Group**: A set of fields and controls for one Exception Group. See **Exception Group Card**.
* **Add Exception Group** (shown only if access setting is **Deny All except**): A button that allows you to create a new group of IP addresses and/or subnets with access to the Kentik subsystem (see **Add an Exception Group**).
* **Save button** (shown only if access setting is **Deny All except**): A button that saves changes to access control settings for this subsystem.

  ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(103).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A09Z&se=2026-05-12T09%3A33%3A09Z&sr=c&sp=r&sig=sa%2B3hu1dMqqAFfmTJdP%2BVI9GYjsZwZdT1I6I%2F9g5sMs%3D)

#### Exception Group Card

Each exception group is defined by the settings in a card (gray rectangle). The card for one exception group, named "Migrated" by default, appears when you choose the **Deny All except** radio button. An additional card, named "Exception" by default, will open each time you click the **Add Exception Group** button.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(102).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A09Z&se=2026-05-12T09%3A33%3A09Z&sr=c&sp=r&sig=sa%2B3hu1dMqqAFfmTJdP%2BVI9GYjsZwZdT1I6I%2F9g5sMs%3D)

Each card contains the following fields and controls:

* **Exception Name**: A field in which you can enter a name for your exception group, overwriting the default name.
* **IP Addresses/Subnets**: A field that allows you to add a comma-separated list of IP addresses and/or subnets that are granted access rights to this Kentik subsystem. Access from all other IPs and subsystems will be denied.

  > **Note:** Access Control is based on the public IP from which you attempt to connect to a given Kentik subsystem, even if your client is on an RFC1918 private LAN or behind NAT.
* **Remove**: A button (trash icon) that opens a confirmation dialog from which you can remove the exception group (see **Remove an Exception Group**).

### Access Control Defaults

The default access control setting varies depending on the Kentik subsystem:

* **App**: Default is **Allow All**. Switch to **Deny All except** is recommended.

  > **Note:** If you choose to switch to **Deny All except**, be sure to enter in the **IP Addresses/Subnets** field the IP address from which you are connected or else your access will be denied when you click **Save**. Users who try to log in from an IP/subnet that is not allowed will see an "Unable to authenticate" notification on the Kentik login page.
* **Agent**: Default is **Allow All**. Switch to **Deny All except** is recommended.
* **API**: Default depends on when the organization subscribed to Kentik:

  + **Subscribed before April 19, 2017**: Default is **Allow All**.
  + **Subscribed on or after April 19, 2017**: Default is **Deny All except**. Before the API can be used, access must be enabled by allowing the IPs/subnets from which the API will be accessed.

To determine the IP address to allow when you connect to any of the above Kentik subsystems:

* **For App**: Go to `https://ipinfo.io/ip`.
* **For Agent or API**: Run the following command on the server that hosts the agent or code that will be connecting to Kentik:

  ```
  curl https://ipinfo.io/ip
  ```

## Managing Exception Groups

Exception groups are covered in the following topics.

### Add an Exception Group

To add an exception group:

1. On the portal's main navbar, open the **Organization Settings** menu and choose **Access Control**.
2. In the subsystem where you'd like to create an exception group (App, Agent, or API), click the **Deny All Except** radio button.
3. In the resulting **Exception Group Card**, change the default exception group name to something more descriptive of the group’s function.
4. Enter a comma-separated list of IP addresses and/or subnets that should be allowed to access this subsystem.
5. Click **Save**.

### Edit an Exception Group

To edit an existing exception group:

1. On the portal's main navbar, open the **Organization Settings** menu and choose **Access Control**.
2. In an existing **Exception Group Card**, edit the name or the IP address fields.
3. Click **Save**.

### Remove an Exception Group

To remove an exception group:

1. On the portal's main navbar, open the **Organization Settings** menu and choose **Access Control**.
2. In the **Exception Group Card** of the group you'd like to remove, click the **Remove** button (red trash icon).
3. Click **Remove** in the confirmation dialog.
