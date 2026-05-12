# Device Settings

Source: https://kb.kentik.com/docs/device-settings

---

This article discusses the settings for adding a new flow device to Kentik using the **Device** dialog.

![Device configuration interface showing fields for name, description, site, and type selection.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/DS-device-dialog-general.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A49Z&se=2026-05-12T09%3A40%3A49Z&sr=c&sp=r&sig=bP4GsEFVGMQCmFuDKrR43vvOOQsRH2lP%2BGqwhPA2UJU%3D)

*Device settings are made on the tabs of the Device dialog.*

> **Notes:**
>
> * Kentik’s Member role does not allow users to add, edit, or remove devices.
> * For more about the **Devices** page, see **Networking Devices**.
> * Devices may be set up during your organization's initial onboarding with Kentik (see **Getting Started**).
>
>   + If you would like assistance with any aspect of adding or managing a device, contact Kentik (see **Customer Care**).
>   + Devices can also be added and edited with the **Device API**.

Kentik needs certain information to connect with your devices to perform operations such as receiving flow, polling SNMP, and establishing BGP peering. These settings are found in the Settings » **Devices** page (see **Devices**) when adding a new device as follows:

* **Add Device**: Click the **Add Device** dropdown at the top right and select **+ Add Flow Device**.

> **Notes:**
>
> * Available settings vary depending on device type (router vs. host) and whether the device is being added or edited.
> * Each device settings tab is marked with a caution icon until its required settings are complete.
> * The Device dialog also includes information for configuring the devices themselves to connect with Kentik (see **Device Config Info**).

## Common UI Elements

The Device settings dialog includes the following common UI elements:

* **Cancel**: Click **Cancel** at lower right or the **X** at top right to exit the dialog without saving changes.
* **Tabs**: Click a dialog tab to view it, choosing from **General**, **Flow**, **SNMP**, **Streaming Telemetry**, **BGP**, **Integrations**, and **SSH**.
* **Add Device**: Click to add the new device and exit the dialog.

## General Tab

The **General** tab of the Device settings dialog has the shared settings for all data types that Kentik can collect from a device:

| Element | Control Type | Description |
| --- | --- | --- |
| Name | Editable field | User-supplied string. |
| Description | Editable field | User-supplied string. |
| Site | Searchable dropdown | The site to which the device is assigned (see **Sites**). |
| Role | Dropdown | Appears if the selected site’s architecture is defined in its settings (see **Sites**); allows you to choose the role of this device. |
| Create a new Site | Button | Opens the settings covered in **Create Site Settings**, which allow you to create a new site. |
| Flow Device Type | Dropdown | Specify the category of the device (see **Supported Device Types**).  **Note:** Older device types for hosts, e.g., DNS (host-nProbe-dns-www), nHst (host-nProbe-basic), and kproxy — are deprecated. |
| Label(s) | Searchable dropdown | Shows the labels applied to the device; when clicked, opens a dropdown from which you can select **Device Labels**. |
| Add Label | Link | Opens the **New Label Dialog**. The new label is not automatically added to the device. |

#### Create Site Settings

When you click **Create a New Site**, the following settings are shown:

| Element | Control Type | Description |
| --- | --- | --- |
| Site Name | Editable field | User-supplied string. |
| Street Address | Searchable dropdown | The physical location of the site given as a street address. Only an address chosen from the dropdown will be geocoded to derive additional dimensions.  **Note:** The **X** at right clears the field. |
| Use an Existing Site | Button | Exits the creation of a new site, hiding the **Site Name** and **Street** **Address** fields and showing the **Site** field. |

## Flow Tab

The **Flow** tab of the Device settings dialog is covered in the following topics.

### Collection Method

The **Collection Method** pane tells you to which Kentik IP/port your device should send flow records, varying by the selected collection method:

* **Direct:** Displays the **Kentik Ingest IP** and **Kentik Ingest UDP port**.
* **via Flow Proxy**: Once you select a Universal Agent from the dropdown, displays the **Flow Proxy Private IP** and **Flow Proxy port**.
* **via Kproxy:** Once you select a Universal Agent from the dropdown, displays the **kproxy Private IP** and **kproxy port**.

> **TIP**: The **Need help configuring Flows?** link opens the **Router Configuration**KB article for more information.

### Flow Export Configuration

The **Flow Export Configuration** pane enables you to specify the IP(s) from which Kentik should expect flow data for this device:

* **Sending IPs**: Enter the IP address(es) from which the router sends flow to Kentik.

  > **Note:** The IP must be unique except as described in **IP Overloading**.
* **Add Sending IP**: A button that adds a new **Sending IPs** field.

> **TIP**: Click the **X** at the right of a **Sending IPs** field to remove it.

* **Sample Rate**: Enter the ratio of the total flows transiting the device to the flows sent to Kentik (see **Flow Sampling**). For example, entering 1,000 means that one flow record will be sent to Kentik for every 1,000 flows transiting the device.

  > **Notes:**
  >
  > + Kentik may downsample from this nominal sample rate to stay within the Flows Per Second (FPS) limits specified in the device’s plan (see **Kentik Plans**).
  > + For hosts, see **Sample Rate for Hosts**.
* **Licenses**: Select a Flow Plan from the dropdown for this device to use.

### IP Overloading

For devices sending flow data directly to Kentik, the IP address(es) specified in **Sending IPs** must not be used for any other such devices in your organization. However, for devices sending flow data via Kentik's **Universal Agent**, the IP(s) may be shared across devices if the following is true:

* The two devices do not use the same instance of `kproxy`.
* Both instances of `kproxy` specify a valid site ID using the `-site_id` parameter in the `kproxy` command line (see **kproxy Proxy Agent Arguments**).
* The value of `-site_id` for the two instances of `kproxy` is not the same.

### Sample Rate for Hosts

Kentik’s **Universal Agent** generates network traffic data from hosts, and the sampling rate of that data involves two settings:

* The `--sample` CLI parameter (optional).
* The **Sample Rate** setting on the Flow tab (required when a host device is registered in the portal).

The actual sample rate used by Kentik is determined by the following:

* If the `--sample` parameter is included in the command line, the CLI-provided value takes precedence over the **Sample Rate** field value.
* If the `--sample` parameter is not included in the command line, the **Sample Rate** field value is used.

> **Note:** If the value is not set in the command line and the **Sample Rate** field value is reset in the portal, the corresponding Universal Agent instance will exit. If the agent is not run under a supervisor, then it must be restarted manually.

## SNMP Tab

The **SNMP** tab of the Device settings dialog is covered in the following topics.

### Collection Method

The Collection Method pane includes the following options for collecting SNMP data from this device.

* **None. Do not collect metrics via SNMP**: No further settings necessary.
* **Flow Enrichment via Universal Agent**
* **Full monitoring via Universal Agent**: Same settings as **Flow Enrichment via Universal Agent** except the Everything monitoring template is preselected.
* **(Legacy) Flow enrichment via kProxy SNMP**
* **(Legacy) Full monitoring via Universal Agent + kProxy Flow Enrichment**

### Flow Enrichment via Universal Agent

The **For Flow Enrichment via Universal Agent** pane includes the following:

* **Credential**: Choose the credential to be used for the collection agent via the dropdown (see **SNMP NMS Device Credential**).
* **Collection Agent**: Select, via the dropdown, the Kentik agent to use to collect SNMP data.
* **Device SNMP IP**: Enter the SNMP IP address for Kentik to poll.
* **Port**: Enter the port number for Kentik to poll.
* **Timeout**: Enter the time in seconds for Kentik to wait for an SNMP server response.

#### SNMP NMS Device Credential

The **Credential** field allows you to create a new credential or use an existing credential from Kentik’s **Credentials Vault**.

* **To use an existing credential**: Click the **Credential** dropdown and select a credential from the list.
* **To create a new SNMP credential**: Click the **Credential** dropdown and select "New Credential", which opens the **Add SNMP Credential Dialog**.

#### Add SNMP Credential Dialog

The **Add SNMP Credential** dialog is used for creating a new credential for a device, and includes the following UI elements:

* **Cancel**: A button at lower right and an **X** at top right that both exit the dialog without creating the credential.
* **Type**: Choose via radio buttons the credential type to match the SNMP version used to poll the device (v1, v2c, or v3). The remaining fields of the dialog vary depending on this setting.

  > **Note**: All three SNMP types use system templates for creating a credential, so neither their key/value pairs nor their key names can be changed (see **Keys by Credential Type**).
* **Credential Name**: Enter a name for the credential (alpha-numeric characters, dashes, and underscores are allowed).

  > **IMPORTANT**: Once saved, a credential’s name cannot be changed.
* **Labels**: Click the field to open a searchable list of labels. To apply a label, click it in the list, after which it will appear in the field. Repeat to apply multiple labels. When done, click outside the list to close it.

  > **Note:** To remove a label, click the **X** at the right of that label.
* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(557).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A49Z&se=2026-05-12T09%3A40%3A49Z&sr=c&sp=r&sig=bP4GsEFVGMQCmFuDKrR43vvOOQsRH2lP%2BGqwhPA2UJU%3D)**Add/Edit**: Opens the **Labels Page** to create or remove the available labels.
* **Community** (SNMP v1 and SNMP v2c only): Enter a masked value for the `community` key.
* **User Name** (SNMP v3 only): Enter a value for the `username` key.
* **Authentication** (SNMP v3 only):

  + **Authentication protocol**: Select the `authentication` protocol (None, MD5, or SHA) from the dropdown.
  + **Value**: Enter a masked value (passphrase) for the authentication key.
* **Privacy** (SNMP v3 only):

  + **Privacy protocol**: Choose the privacy protocol (None, DES, or AES) from the dropdown.
  + **Value**: Enter a masked value (passphrase) for the `privacy` key.
* **Add Credential**: Click to save the new credential and exit the dialog.

> **Notes**:
>
> * The contents of fields used for masked values can't be seen as text is entered.
> * The behavior of fields used to enter key values for credentials is covered in **Credential Values**.
> * While descriptions can't be specified for a credential in the **Add SNMP Credential** dialog, they can be added by editing in the Credentials Vault (see **Edit a Credential**).

### **Flow Enrichment via kProxy SNMP**

The **For Flow Enrichment via kProxy** pane includes the following fields and controls:

* **Kentik SNMP polling IPs**: A callout giving the IP addresses of the Kentik SNMP Polling IPs.

  > **Note:** Use these IPs when configuring SNMP on the device itself, which is required to enable polling by Kentik.
* **SNMP polling**: A dropdown to choose the polling frequency for SNMP:

  + **Standard**: Interface counter is polled every 5 minutes; interface description is polled every 3 hours.
  + **Minimum**: Interface counter isn’t polled; interface description is polled every 6 hours.
* **Device SNMP IP**: Enter the IP address for Kentik to poll for SNMP.
* **SNMP Community**: Enter the SNMP community for Kentik to use when polling the router via SNMP v1 or v2c.

  > **Notes:**
  >
  > + Not shown when SNMP v3 is enabled.
  > + The entered string is obscured but can be edited.
* **Enable SNMP v3 Authentication**: When switched On, sets SNMP polling to v3 (see **About SNMP V3**), displays the **SNMP v3 Settings Pane**, and hides the **SNMP Community** field

  > **Note:** Available for routers only; overrides SNMP Community setting (above).

  + **SNMP v3 User Name**: Enter the user name for SNMP v3 authentication (required).
  + **Authentication**: Choose the SNMP v3 authentication protocol from the dropdown (None, MD5, or SHA):

    - **Passphrase**: Enter the password for SNMP v3 authentication.
  + **Privacy**: Select the SNMP v3 privacy protocol from the dropdown (DES or AES):

    - **Passphrase**: Enter the password for SNMP v3 privacy.
    > **Note:** Entered strings are obscured but can be edited.

## Streaming Telemetry Tab

The **Streaming Telemetry** tab of the Device settings dialog enables the collection of ST data for two distinct purposes:

* **Dial-In from Universal Agent**: Advanced ST using Kentik NMS.
* **Dial-Out to kProxy**: Basic ST included with the Kentik platform.

To configure ST for the device, select either or both of the **Collection Method** checkboxes to open the corresponding panes described below.

### Dial-In from Universal Agent

The **Dial-In from Universal** **Agent Options** pane enables the Universal Agent to connect to the device via the gNMI protocol, and includes:

* **Credential**: Select, via the dropdown, the credential to use for the collection agent:

  + **To use an existing credential**: Click the **Credential** dropdown and select a credential from the list.
  + **To create a new ST credential**: See **Add ST Device Credential**.
* **Collection Agent**: Displays the Kentik agent to use to collect ST data, same as selected in the **For NMS via Universal Agent** pane.
* **Timeout**: Enter the time in seconds for Kentik to wait for a server response.
* **Port**: Enter the device's listening port number.
* **Use secure connection (SSL)**: When switched On, turns on SSL encryption for the data transmission.

> **TIP**: You can access ST data collected with this method with **Metrics Explorer**.

#### Add ST Credential Dialog

The Add Streaming Telemetry Credential dialog includes the following UI elements:

* **Credential Name**: Enter a credential name (alpha-numeric characters, dashes, and underscores allowed only).

  > **IMPORTANT**: Once saved, a credential’s name cannot be changed.
* **Labels**: Click the field to open a searchable list of labels. To apply a label, click it in the list, after which it will appear in the field. Repeat to apply multiple labels. When done, click outside the list to close it.

  > **Note:** To remove a label, click the **X** at the right of that label.
* **Add/Edit**: Opens the **Labels Page** to create or remove the available labels.
* **User Name**: Enter a value for the username key.
* **Password**: Enter a masked value for the password key.
* **Cancel**: Click **Cancel** at lower right or the **X** at top right to exit the dialog without saving.
* **Add Credential**: Click to save the new credential and exit the dialog.

> **Notes**:
>
> * The contents of fields used for masked values can't be seen as text is entered.
> * While labels and descriptions can't be specified for a credential in the **Add Streaming Telemetry Credential** dialog, they can be added using the Credentials Vault (see **Edit a Credential**).

### Dial-Out to kProxy

With the “dial-out” collection method, devices are pre-configured to send Streaming Telemetry (ST) data to a `kproxy` agent. The **Dial-Out to kProxy Configuration** pane of the includes the following information:

* **Sending IP**: The IP address from which the device sends ST data.
* **Kentik Ingest IP**: The Kentik IP address to which to send ST data.
* **Kentik ST port**: The port on which Kentik will receive ST data.

> **TIP**: ST data collected with this method is accessible through **Data Explorer**.

## BGP Tab

The settings on the **BGP** tab of the Device settings dialog determine how Kentik collects BGP data for enriching a device’s flow records.

> **Notes:**
>
> * The BGP tab’s settings are only available when the selected Billing Plan (see **Flow Tab**) supports BGP (see **About Plans**).
> * The available settings vary depending on the chosen **BGP Type**.

### Common BGP Settings

The following controls are common for all **BGP Type** settings:

* **BGP Type**: Select how BGP will be enabled on this device via the dropdown:

  + **No peer, use generic IP/ASN mapping**: Map the device's IP address to an ASN.
  + **Peer with device**: Kentik will BGP peer with this device (this setting is required for RTBH/Flowspec).
  + **Use table from another peered device**: The BGP table will be obtained from another device that is already set to peer with Kentik.
* **BGP Flowspec Compatible**: Switch this to On if the router supports MP-BGP and is therefore compatible with BGP Flowspec.
* **Segment Routing**: Switch this to On if the router is part of a segment-routing-enabled fabric, allowing the platform to better decode and attribute traffic that might otherwise look like generic MPLS or encapsulated packets.
* **BGP Route Selection**: Choose, via the dropdown, the method for Kentik, for both VRF and non-VRF interfaces, to match each flow’s IP address against the BGP route received via the device's BGP sessions (see **BGP Route Selection**).

### Type-Specific BGP Settings

The following controls are present only for specific **BGP Type** settings:

* **Use AS Numbers from Flow**: When switched On, Kentik will retain the source and destination ASN information from the flows exported by the network device.

  + Present only when **BGP Type** is set to "No peer, use generic IP/ASN mapping”.
* **IPv4 Peering Address**: Enter the IPv4 address of the peering device.

  + Present only when BGP is set to “Peer with device”.
  + RFC1918 addresses are not valid.
  + Cannot be an IP that is already being used to peer with a different Kentik device.
* **IPv6 Peering Address**: Enter the IPv6 address of the peering device.

  + Present only when BGP is set to “Peer with device”.
  + RFC1918 addresses are not valid.
  + Cannot be an IP that is already being used to peer with a different Kentik device.
* **ASN**: Enter the number (16- or 32-bit) of the autonomous system (AS) to which the peering device belongs.

  + Present only when **BGP Type** is set to "Peer with device”.
* **BGP MD5 Password**: Enter an optional shared authentication password (32 alphanumeric characters) for BGP peering.

  + Present only when **BGP Type** is set to “Peer with device".
* **Primary BGP Device**: Choose the device from which the BGP table will be shared with this device.

  + Present only when **BGP Type** is set to “Use table from another peered device”.

### Kentik Ingest Peering IPs

When **BGP Type** is set to "Peer with device," Kentik's Ingest Peering IPs (v4 and v6) will be shown to the right of the Peering Address fields (IPv4 and IPv6) listed above. You'll use these as the IPs that the device should be set to peer with in the BGP configuration on the device itself.

### BGP Route Selection

As flow records from devices are ingested into the **Kentik Data Engine** (KDE), they are enriched with BGP/routing information. This process depends on matching each flow’s IP address against the BGP route received via the device's BGP sessions.

The **BGP Route Selection** dropdown determines how this matching will be performed for both VRF and non-VRF interfaces, as shown in the following table:

| Dropdown Menu Option | VRF Interface | Non-VRF Interface |
| --- | --- | --- |
| VPN table for VRF interface, Unicast table for non-VRF interface (default) | Use only L3VPN routes. | Uses only Unicast routes. |
| VPN table, fallback to Unicast table | Uses L3VPN. If no match, uses Unicast | Uses Unicast. |
| VPN table, fallback to Labeled-Unicast table, fallback to Unicast table | Uses L3VPN. If no match, uses Labeled-Unicast. If no match, uses Unicast. | Use Labeled-Unicast. If no match, uses Unicast. |
| VPN table for VRF, Labeled-Unicast, fallback to Unicast for non-VRF (both directions) | Uses L3VPN. If no match, uses Labeled-Unicast. If no match, uses Unicast.  **Note:** This variation checks routing tables based on interfaces in both directions. | Same as VRF |
| VPN Table for VRF, Unicast for non-VRF (both directions) | Use only L3VPN routes, but check routing tables based on interfaces in both directions. | Same as VRF |

> **Note:** For each of the "both directions" options above, the routing table is first checked based on the interface in the opposite direction, then checked based on the interface in the same direction.

## Integrations Tab

The **Integrations** tab on the Device settings dialog is used to specify settings for device-specific integrations.

> **Note**: This tab is currently only used to specify the **Kentik Firehose URL Endpoint** to which the kflow from this device should be sent. For more information, see **Firehose Data Sources**.

## SSH Tab

The **SSH** tab on the Device settings dialog is used to enable Kentik’s **AI Advisor** to fetch configurations and run diagnostic commands on this device. See **SSH Command Access** for details on this tab’s UI elements.

## Device Config Info

The following table shows which tab of the Device settings dialog to find the required information for configuring devices for Kentik:

| Field | Dialog tab | Location | Description |
| --- | --- | --- | --- |
| Kentik Ingest IP | Flow | Collection method pane, Direct tab | The IP address at Kentik to which your router should be configured to send data. |
| Kentik Ingest UDP port | Flow | Collection method pane, Direct tab | The port at Kentik to which your router should be configured to send data. |
| Kentik SNMP polling IPs | SNMP | For Flow Enrichment pane | The IPs from which your router should be configured to allow SNMP polling using the Community supplied in the router configuration. |
| Peering Addresses | BGP | Peer with device settings | The IPv4 and IPv6 addresses to which the device should be set to peer in the BGP configuration on the device itself. |

FPS stands for flows per second, which is the rate at which flow metadata is sent to Kentik from a given customer. It directly correlates with the resources required to provide that customer with the Kentik service.
