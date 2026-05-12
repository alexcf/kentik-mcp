# kproxy Agents

Source: https://kb.kentik.com/docs/kproxy-agents

---

This article discusses the **kproxy Agents** page of theKentikportal's Settings.

> **Note:** In Kentik NMS, SNMP and Streaming Telemetry are enabled by our Universal agent rather than by `kproxy` (which will eventually be incorporated into the Universal agent). For more information see:
>
> * **Device SNMP Settings**: Selecting one or both agents to enable SNMP on a device.
> * **Device ST Settings**: Selecting one or both agents to enable Streaming Telemetry on a device.
> * **Kentik NMS Configuration**: Universal agent setup for NMS.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(641).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A31Z&se=2026-05-12T09%3A34%3A31Z&sr=c&sp=r&sig=TyA%2B96J2gjg71CH6OCNCyinfnXrIwIMAW3FIFI8sZTo%3D)

*Manage settings for proxy agents on the kproxy Agents page.*

## About the `kproxy` Agent

Kentik provides a free **NetFlow** proxy agent called `kproxy` that enables data collected from network devices to be sent securely to Kentik. The machine running `kproxy` isn't actually handling traffic directly, but instead collects flow records (NetFlow v5/v9, IPFIX, and sFlow) and SNMP and encrypts it locally before forwarding it to Kentik. The agent also handles rate limiting and resampling as needed to keep maximum flows per second (FPS) within applicable plan limits. A single instance of the `kproxy` executable can redirect flow for multiple routers and switches. Multiple servers across the network can run `kproxy` to distribute traffic and load.

> ***Notes:***
>
> * If a device sends flow to Kentik via a `kproxy` agent then in Kentik UI related to that device the agent forwarding the data will be indicated, e.g., on the **More Info** tab of the device's Details page in Core (see **More Info Tab**).
> * For more detailed information on `kproxy`, including step-by-step procedures for configuration, see **Kentik Proxy Agent**.

## `kproxy` Status

A deployed `kproxy` agent can be in one of the following states, which are indicated in the **Status** column of the kproxy List:

* **Healthy**: The agent was last seen by the backend less than 30 minutes ago.
* **Warning**: The agent was seen by the backend more than 30 minutes ago, but less than 45 minutes ago.
* **Critical**: The agent was seen by the backend more than 45 minutes ago.
* **Pending**: The agent has been installed via **Onboarding** but hasn’t yet contacted the Kentik backend.

## `kproxy` Agents Page

The **kproxy Agents** page is covered in the following topics.

### About the kproxy Agents Page

The **kproxy Agents** page is used to manage your organization’s `kproxy` agents and to see their status. The page includes the following main UI elements:

* **Add kproxy Agent**: A button that opens the Adding a kproxy Agent dialog (see **Add kproxy Agent**).
* **kproxy Agents Map**: A map that shows the location of your deployed `kproxy` agents (see **kproxy Map**).
* **kproxy Agents List**: A table that lists your deployed `kproxy` agents and provides information about each agent (see **kproxy List**).

### `kproxy` Map

The **kproxy Map** shows the physical location of `kproxy` agents throughout the world. Hovering on a shown agent opens a popup with identifying information about that agent. Double-clicking on any point in the map zooms the display into the clicked area.  
  
The **kproxy Map** is filterable using the filter controls above the kproxy List, which enable you to show or hide agents by status. As various agents are shown or hidden the map will zoom automatically to include all of the agents that are still shown.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/KPA-kproxy_map-674w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A31Z&se=2026-05-12T09%3A34%3A31Z&sr=c&sp=r&sig=TyA%2B96J2gjg71CH6OCNCyinfnXrIwIMAW3FIFI8sZTo%3D)

### `kproxy` Controls

The area between the map and the list includes the following fields and controls, which affect the display of agents in the list and on the map:

* **Filter** (magnifying glass icon): Filters the list to show only the rows for agents in which the entered string is contained in one of the column values.
* **Status**: A drop-down from which you can choose a status. Only the rows for agents with that status will be displayed.
* **Auto-refresh**: A checkbox that determines whether or not the page will be regularly refreshed.
* **Refresh**: A button that manually refreshes the content of the page.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(642).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A31Z&se=2026-05-12T09%3A34%3A31Z&sr=c&sp=r&sig=TyA%2B96J2gjg71CH6OCNCyinfnXrIwIMAW3FIFI8sZTo%3D)

### `kproxy` List

The kproxy List provides information about the `kproxy` agents deployed in your organization. Each of which is represented by an expandable row.

#### `kproxy` List Columns

The kproxy List includes the following columns:

* **Status**: The current status of the agent (see **kproxy Status**).
* **Last seen**: The last time the agent passed any flow to Kentik.
* **Agent name**: The name assigned to the agent by your organization.
* **Site**: The physical location where the agent is deployed, typically a data center in your organization's infrastructure (see **Sites**).
* **IP Address**: The IP address from which Kentik receives flow data from this instance of kproxy (i.e. the WAN address used to identify the `kproxy` instance in Kentik).
* **Private IP**: The address to which data sources should send flow records for this instance of `kproxy`. The address (which need not be RFC1918) is assigned to the agent in the **kproxy Agent Dialog**.

  > **Note:** In the snippets that Kentik provides in the **Device Configs Directory**, the placeholder for this IP is `{{kentik_flow_proxy_IP}}`.
* **Version**: The version number of the `kproxy` agent.
* **Edit** (icon): A button that opens the kproxy Agent dialog, where you can change settings for the agent.

  > **Note:** If either the **Site** or **Private IP** settings are indicated as "Undefined" then the **Edit** button is labeled "Update."
* **Remove** (icon): Opens a confirming dialog that enables you to remove the agent from your organization's collection of agents.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(643).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A31Z&se=2026-05-12T09%3A34%3A31Z&sr=c&sp=r&sig=TyA%2B96J2gjg71CH6OCNCyinfnXrIwIMAW3FIFI8sZTo%3D)

*Click a row in the* *kproxy* *list to reveal additional agent details.*

#### `kproxy` Row Details

To see further details about a `kproxy` agent in the kproxy List, click on the agent's row in the list. The row will expand, showing the following additional information:

* **Description**: The description string provided by the user in the **kproxy Agent Dialog**.
* **Devices**: The devices sending flow to Kentik via this instance of `kproxy`.
* **Agent API User**: The value specified for the `--api_email` argument in the `kproxy` CLI when the `kproxy` instance was started (typically the value specified in the instructions provided in the Adding a kproxy Agent dialog).
* **Last User Agent**: The User Agent string that the `kproxy` instance uses to place HTTP POST/GET requests to KDE (Kentik backend).

## `kproxy` Agent Dialog

The kproxy Agent dialog enables you to set properties of a kproxy instance. The dialog is covered in the following topics:

> **Note:** This dialog is accessed via the **Edit** button in the **kproxy List**.

### `kproxy` Agent UI

The kproxy Agent dialog includes the following UI elements:

* **Close button**: Click the **X** in the upper right corner to close the dialog without saving changes to the agent settings.
* **Properties**: The basic properties of the agent (see **kproxy Property Fields**).
* **Cancel button**: Close the dialog without saving changes. All elements will be restored to their values at the time the dialog was opened.
* **Save button**: Save changes to agent settings and exit the dialog.

### `kproxy` Property Fields

The following fields set basic properties of the Private agent:![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(644).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A31Z&se=2026-05-12T09%3A34%3A31Z&sr=c&sp=r&sig=TyA%2B96J2gjg71CH6OCNCyinfnXrIwIMAW3FIFI8sZTo%3D)

* **Name**: Your organization's name for the agent. By default, this is "Auto Generated."
* **Description**: An optional user-provided description of the agent.
* **Private IP**: The address at which data sources should send flow records to this instance of `kproxy`.
* **Site**: Choose a site (see **About Sites**) to which to assign the agent, which could be a data center or a VPC in a cloud provider.
* **Create a New Site**: If a site hasn't yet been defined for the location where the agent host is located, click the button to create a new site (see **Configure Site Fields**).

#### Configure Site Fields

When you click the **Create a New Site**button, the following fields are added to thekproxy Agent dialog to enable you to enter the information required to create a new site:

* **Site Name**: A name for the new site.
* **Street Address**: The physical location of the site given as a street address.
* **Use an Existing Site**: If you decide not to create a new site, click this button to hide the **Configure Site** fields and instead choose a site from the **Site** drop-down.

## Add `kproxy` Agent

The **Add kproxy Agent** button opens the Adding a kproxy Agent dialog, which includes expandable panes that each correspond to a `kproxy` deployment scenario. Click on a pane to expand it and see instructions for one of the following scenarios:

* **Linux packages** (see **kproxy Package Install**):

  + Debian/Ubuntu
  + RPM
* **Docker** (see **kproxy Docker Install**).

![Instructions for deploying a kproxy agent using Docker and environment variables.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/KA-adding-kproxy.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A31Z&se=2026-05-12T09%3A34%3A31Z&sr=c&sp=r&sig=TyA%2B96J2gjg71CH6OCNCyinfnXrIwIMAW3FIFI8sZTo%3D)

> **Note:** The instructions in the Adding a kproxy Agent dialog include authentication values for the kproxy CLI arguments `--api_email` and `--api_token`. Kentik recommends using these organization-wide values rather than using the API email and token values from your own user profile.

Once you've worked through the agent deployment steps at the links above the new kproxy instance will appear in the **kproxy List** with **Agent Name** indicated as "Auto Generated." To complete the settings for the agent:

1. Click the **Edit** icon at the right of the agent's row in the list (the icon will be labeled "Update" until initial settings are complete).
2. In the resultingkproxy Agent dialog, fill in the following **kproxy Property Fields**: **Name**, **Private IP**, and **Site**.
3. Click the **Save** button to save the settings and close the dialog.

> **Note:** `kproxy` instances may also be deployed via the v4 portal's **Onboarding** workflow. The status of an agent added via onboarding will appear in the **kproxy List** as "Pending" until Kentik begins to receive flow from that agent.
