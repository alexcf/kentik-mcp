# Universal Agents

Source: https://kb.kentik.com/docs/universal-agents

---

This article covers the **Universal Agents** page in the Kentik portal.

![Overview of Universal Agents with their status, capabilities, and performance metrics displayed.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/UA-main.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A35Z&se=2026-05-12T09%3A43%3A35Z&sr=c&sp=r&sig=83A8I68FwGVeRJrH6KsSRq88mTfw8%2FUCltx6zD4ebcY%3D)

*Manage your Universal Agents on the Universal Agents settings page.*

## About the Universal Agent

Kentik’s **Universal Agent** collects key metrics on the availability, health, and performance of your network infrastructure. After installing the agent on a host machine on your network (see **Deploy a Universal Agent**), you add features to the agent by installing **Agent Capabilities**.

> **Did You Know?**: All Kentik **Agents** will gradually be replaced with **Agent Capabilities** of the Kentik Universal Agent.

### Agent Capabilities

General features of a Universal Agent capability are:

* Automatically updates itself
* Sends host metrics and capability metrics, usable in Data Explorer and Metrics Explorer
* Agent/capability metrics can be used to create agent alerts

#### **Supported Capabilities**

| **Capability** | **Description** | **Use Cases** | **Related Workflows** |
| --- | --- | --- | --- |
| **SNMP/ST** | Discovers/collects NMS device data (SNMP or Streaming Telemetry) | * Monitor health and status of NMS devices. * Diagnose issues by analyzing device metrics. * Create alerts and notifications based on device metric threshold breaches and state changes. | * **NMS Dashboard** * **NMS Devices** * **NMS Interfaces** * **Metrics Explorer** * **Query Assistant** * **Alerting** |
| **SNMP Trap Receiver** | Captures real-time events from SNMP-enabled devices | * Filter and search trap events by name and OID. * Create alerts and notifications based on SNMP trap events. * Visualize trap events with other telemetry for quicker root cause analysis. | * **Data Explorer** * **Alerting** |
| **Syslog Server** | Captures real-time log messages from devices | * Filter and search syslog events by name, severity, and message content. * Create alerts and notifications based on syslog events. * Visualize syslog events with other telemetry for quicker root cause analysis. | * **Data Explorer** * **Alerting** |
| **DNS OTT Tap** | Provides deep visibility into DNS infrastructure by converting DNS queries into flow records, in addition to tapping into DNS data for OTT services | * Detect and analyze “content events” for network operations guidance. * Evaluate OTT metrics for user categories and delivery methods (e.g., CDNs, interconnection types, or PoPs). * Assess Mbps-per-subscriber for OTT service impact. * Consider zero-rating content provider implications. * Improve customer retention by identifying delivery issues. * Analyze suspicious traffic for legal liability. * Get alerts on OTT service performance issues. * **DNS Security & Attack Detection**: Identify flood attacks, reflection victims, and invalid queries. * **DNS Server Health**: Monitor response failures, record integrity, and performance baselines. * **DNS Server Configuration**: Validate expec * ted values, audit TTLs, and monitor configuration changes in real-time. | * **OTT Service Tracking** |
| **Flow Proxy** | Forwards network flow telemetry (e.g. NetFlow, sFlow,  IPFIX) to the Kentik platform. | * Collecting flow in situations where  deploying a physical or virtual Kentik appliance is not feasible * Performance monitoring * Security monitoring and threat detection * Capacity planning & optimization | **Data Explorer** |
| **Connectivity Test** | Enables on-demand IP connectivity testing via `ping` or `traceroute` | * Test for device connectivity, latency, or packet loss using `ping` . * Identify bottlenecks, points of failure, or routing loops using `traceroute`. | **Synthetics** |

> **Notes:**
>
> * The **DNS OTT Tap** capability replaces the standalone `kprobe` software agent as the method for tapping DNS data, as described in **Enable OTT DNS Collection**.
> * The **Flow Proxy** and **SNMP/ST** capabilities replace the standalone `kproxy` software agent as the method for forwarding flow telemetry data.

### System Requirements

When **deploying Universal Agents** and enabling capabilities, ensure your host machine meets the necessary hardware requirements for optimal performance. Hardware needs scale based on the volume of telemetry data and the specific capabilities enabled.

| Capability / Agent | Recommended Hardware | Additional Notes |
| --- | --- | --- |
| **Flow Proxy (**`kproxy`**)** | 8 CPU cores, 32GB RAM | Recommended for 25,000+ FPS. Add a 20-30% buffer for bursty environments. |
| **Flow Proxy (**`ktraffic`**)** | 1 CPU core, 2GB RAM | Recommended for 20,000+ FPS. (Note: Beta new flow proxy). |
| **SNMP/ST** | 2 CPU cores, 8GB RAM | Recommended for environments with up to 500 Devices. |
| **Syslog Server** (`ksyslog`) | **[Y]** CPU cores, **[Z]** GB RAM | Requirements scale based on log volume: Recommended for **[X]** events/sec. |
| **Synthetics (Network & App)** | 4-8 CPU cores, 8GB RAM | Recommended for enterprise-scale, high-volume testing. Provides necessary headroom for burst test scenarios and multiple simultaneous synthetic tests. |
| **BGP Proxy** | 2 CPU cores, 2GB RAM | Having extra capacity helps maintain stability during BGP table updates. |
| **DNS OTT Tap** | 1 CPU core, 1GB RAM | Required for every 400k DNS lookups/sec. |

### Agent Status

A deployed Universal Agent can be in one of the following states, as indicated in the **Status** column of the Universal Agents list:

* **Up:** The agent was last seen by Kentik within 30 minutes.
* **Down:** The agent was last seen by Kentik over 30 minutes ago.
* **Not Authorized:** The agent is not yet authorized to connect to Kentik.

> **Note:** Authorization is part of the agent installation process (see **Deploy a Universal Agent**).

## Deploy a Universal Agent

To deploy a new instance of the Universal Agent via Docker or Linux, follow these steps:

1. Go to the **Universal Agents** page via **Settings****»** **Universal Agents** in the Kentik portal’s main nav menu.
2. Click **+ Deploy Agent** to open the dialog.
3. Under **Installation**, choose Docker or Linux and click the corresponding tab. The necessary command will appear in the text area below the tabs.
4. Click **Copy** to copy the command from the text area.
5. Open a terminal window on the host machine and paste the command. (To use a web proxy, see **Install Agent via Web Proxy**).
6. Run the command. The new agent will appear in the **Select an Agent** list. The **Host Name** column will vary as follows:

   1. **Docker**: The agent's Docker container ID.
   2. **Linux**: The hostname of the Linux system where the agent was installed.
7. Click **Authorize** to authorize the agent to connect with Kentik. The dialog will refresh with basic information, optional fields, and a **Continue** button.

   ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(649).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A35Z&se=2026-05-12T09%3A43%3A35Z&sr=c&sp=r&sig=83A8I68FwGVeRJrH6KsSRq88mTfw8%2FUCltx6zD4ebcY%3D)
8. Optionally add **Display Name**, **Description,** and **Site Name**. Click **Continue** to proceed.
9. Select the capabilities to enable for this agent by clicking the toggle switch next to each capability.

   ![Deployment options for the Universal Agent, highlighting SNMP and Syslog capabilities.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/UA-deploy-agent-dialog-capability-selection(1).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A35Z&se=2026-05-12T09%3A43%3A35Z&sr=c&sp=r&sig=83A8I68FwGVeRJrH6KsSRq88mTfw8%2FUCltx6zD4ebcY%3D)
10. Click the **Finish** button to complete the deployment process. Once deployed, the agent will appear on the **Universal Agents List**.

> **Tip:** If you close the dialog before clicking Authorize, the agent will appear on the **Universal Agents List** with a “Not Authorized“ status and won't be functional. To resolve this, click **Deploy Agent** to reopen the dialog and click **Authorize**.

#### Install Agent via Web Proxy

Use the `https_proxy` environment variable to set the proxy's URL:

* **Docker**: Provide the environment variable as part of the `docker run` command:

  `--env https_proxy=http://address.of.your.proxy`
* **Linux**: Set the environment variable before running the install command:

  `export https_proxy=http://address.of.your.proxy`

## Edit a Universal Agent

To edit a deployed Universal Agent instance, follow these steps:

1. Navigate to the **Universal Agents** page via **Settings****»** **Universal Agents** in the Kentik portal’s main nav menu.
2. Click the edit button (pencil icon) in the agent's row in the list.
3. Update the Name, Description, Closest Network Device, and/or Site.
4. Click **Save** to save changes and exit or **Cancel** to exit without saving.

## Install a Capability

To install a capability (see **Agent Capabilities**) on a Universal Agent instance, follow these steps:

1. Go to the **Universal Agents** page via **Settings****»** **Universal Agents** in the Kentik portal’s main nav menu.
2. Click the agent's row in the list to open the **Agent Details Drawer**.
3. Under Capabilities » **Available: Compatible and Supported**, click **Install** for any capabilities you wish to add this Universal Agent instance.

## Manage a Capability

To manage an installed capability on a Universal Agent instance, follow these steps:

1. Go to the **Universal Agents** page via **Settings****»** **Universal Agents** in the Kentik portal’s main nav menu.
2. Click the agent's row in the list to open the **Agent Details Drawer**.
3. Under **Capabilities****»** **Installed: Running or Paused on this Agent**, click **Details**.
4. In the **Capability Drawer**, you can take the following actions:

   1. View the capability’s details like Status, Run State, Command Line Args, and metrics visualizations.
   2. When applicable, click **Edit** to reveal additional capability settings like command line arguments/values.

      1. Click **Save** to save changes and exit or **Cancel** to exit without saving.
   3. Click **Back** to return to the Agent Details drawer.
   4. Click **Disable this Capability** to temporarily turn off this capability for this Universal Agent instance.

## Configure the DNS OTT Tap Capability

Telemetry is collected via the **DNS OTT Tap** capability, which supports three primary data ingestion modes:

* **Direct (On-Server)**: Agent installed directly on BIND9 or Linux-based resolvers.
* **Promiscuous (SPAN/Mirror)**: This mode is used when the agent is installed on a separate host and receives a copy of the traffic from a network tap or port mirror. The agent interface must be set to promiscuous mode to capture traffic not specifically destined for the host's MAC address.
* **DNSTAP Receiver (Preferred)**: Agent acts as a listener for real-time DNS logs from appliances like Infoblox. This mode must be explicitly configured to listen on a specific port for the incoming DNSTAP stream.

To edit the **DNS OTT Tap** capability, follow these steps:

1. Go to the **Universal Agents Page** via **Settings****»** **Universal Agents** in the Kentik portal’s main nav menu.
2. Click the agent's row in the list to open the **Agent Details Drawer**.
3. Under **Capabilities****»** **Installed: Running or Paused on this Agent**, click **Details** for DNS OTT Tap.
4. Click **Edit** to reveal the current command line argument/value settings.
5. Make the changes, as follows:

   1. **Update an Existing Argument**: Select a different **Command Line Value** from the dropdown.
   2. **Add a New Argument**: Click **Add** to add a row, then choose the argument and value (see supported combinations below).
6. Click **Save** to save changes and exit or **Cancel** to exit without saving.

Supported argument/value combinations for the DNS OTT Tap capability:

| **Command Line Argument** | **Accepted Values** | **Description** |
| --- | --- | --- |
| Interface | All, lo, eth0, [other network interfaces] | Specifies the interface(s) from which to capture DNS OTT data (e.g., "All" for all available interfaces, "lo" for the loopback interface, "eth0" for the first Ethernet interface). |
| Promiscuous Mode | On, Off | Determines whether the interface should operate in promiscuous mode. When "On," the interface captures all packets on the segment. "Off" means it only captures packets addressed to its MAC address.  **Note**: If DNS traffic is being collected through SPAN/mirror ports, the interface may need to run in promiscuous mode so it can capture mirrored packets. However, most deployments (like Direct or DNSTAP) do not use promiscuous mode. |

## Configure the Flow Proxy Capability

To configure the **Flow Proxy** capability of the Universal Agent, which receives, processes, and forwards your flow telemetry data to Kentik, follow these steps:

> **Note**: The **Flow Proxy** capability packages the standard Kentik Proxy (`kproxy`) binary into a managed component within the Universal Agent framework.

1. Go to the **Universal Agents** page via **Settings****»** **Universal Agents** in the Kentik portal’s main nav menu.
2. Click the agent's row in the list to open the **Agent Details Drawer**.
3. Under **Capabilities****»** **Installed: Running or Paused on this Agent**, click **Details** for Flow Proxy.
4. Click **Edit** to reveal the current command line argument/value settings.
5. Configure the command-line arguments as necessary (see below table), and click **Save**.

   1. The Universal Agent downloads the `kproxy` binary, applies the configuration, and starts the process.
   2. The agent’s **Status** field on the Universal Agents page updates from 'Pending' to 'Up'.
6. Configure your routers, switches, and other network devices to export flow telemetry to the IP address of the Universal Agent host on the UDP port specified in the **Port** field (e.g., 9995).

| **Command Line Argument** | **Description** | **Example Value** | **Required?** |
| --- | --- | --- | --- |
| **Host** | The IP address on which the proxy will listen for incoming flow packets. Use 0.0.0.0 to listen on all available network interfaces on the host. | 0.0.0.0 | No |
| **Port** | The UDP port used to listen for flow packets from your network devices. | 9995 | No |
| **Base Port** | The port used by Flow Proxy to send flow to Kentik. Adjustable in case of conflict on the host machine. | 40010 | No |
| **Sflow Agent Address** | (sFlow only) Specifies the IP address of the sFlow agent. Can be used to override the agent address in sFlow datagrams if its being reported incorrectly by the device. | 192.168.1.1 | No |
| **Site** | Associates all flow data processed by this Flow Proxy with a specific Site configured in your Kentik portal. This is essential for data filtering and organization. | data-center-frankfurt | No |

### Troubleshooting a Flow Proxy

* **Status is Down**: If the capability status displays 'Down' with an error, such as exit status 10, it typically indicates a configuration error. Ensure that all required fields are correctly populated.
* **No Data in Portal:** If the capability status is ‘Up’ but you see no flow data in the Kentik portal, verify the following:

  + Network devices are configured with the correct destination IP and port.
  + There are no firewalls on the agent’s host or in the network path blocking the flow export UDP port.
  + Check the agent logs for connection errors or other warnings.

## Uninstall a Capability

To uninstall a capability from a Universal Agent instance, follow these steps:

1. Go to the **Universal Agents** page via **Settings****»** **Universal Agents** in the Kentik portal’s main nav menu.
2. Click the agent's row in the list to open the **Agent Details Drawer**.
3. Under **Capabilities****»** **Installed: Running or Paused on this Agent**, click **Uninstall** for the capability you wish to remove from this Universal Agent instance. Once the capability has been removed, you’re returned to the Agents list.

> **Tip**: You can reinstall a capability on an agent at any time by clicking the capability’s **Install** button in the **Agent Details Drawer**.

## Back Up & Restore a Universal Agent

Each Universal Agent maintains a unique identity based on a cryptographic keypair. Preserving this identity is crucial for disaster recovery or when migrating the agent to a new host machine (VM). By restoring the agent's identity, you ensure that the Kentik platform recognizes the agent as the same entity, retaining its configuration and history without requiring re-authorization.

There are two methods for backing up the agent:

1. **Full Directory Backup** (Recommended): Backs up the entire binary and configuration state.
2. **Identity-Only Backup** (Advanced): Backs up only the specific key files required to recover the agent's identity.

#### Method 1: Full Directory Backup (Recommended)

This method involves archiving the entire `/opt/kentik` directory and is the simplest approach and is recommended for most users to back up and restore a Universal Agent.

**Backup**

1. Ensure the agent is authorized and functioning.
2. Create a backup of the entire `/opt/kentik` directory. You can use standard Linux tools (e.g., `tar`, `rsync`, or snapshotting tools) to secure this directory.

**Restore**

To restore a failed agent or migrate to a new VM:

1. Restore the `/opt/kentik` directory from your backup to the new host.
2. Execute the agent repair command to fix permissions, recreate `systemd` service files, and start the service:

```
sudo /opt/kentik/bin/kagent repair
```

> **Note**: This process works even if the new host has a different hostname or IP address.

#### Method 2: Identity-Only Backup (Advanced)

This method is useful for "Infrastructure as Code" environments where you prefer to install a fresh Universal Agent binary but inject the previous agent's identity.

**Backup**

Locate and safely store the following two key files, which constitute the agent's identity:

* `/opt/kentik/components/kagent/current/conf/keys/private_key.pem`
* `/opt/kentik/components/kagent/current/conf/keys/public_key.pem`

**Restore**

To recover the agent state on a new or rebuilt host/VM:

1. **Install the Agent**: Perform a fresh installation of the Kentik Universal Agent on the new host using the standard installation script or package.
2. **Stop the Agent**: Ensure the service is stopped before swapping keys.

```
sudo systemctl stop kagent
```

3. **Restore Keys**: Copy your backed-up `private_key.pem` and `public_key.pem` files to the target directory, overwriting any auto-generated keys:

   `/opt/kentik/components/kagent/current/conf/keys/`
4. **Restart the Agent:** Once restarted, the agent will reconnect to the Kentik platform using the restored identity.

```
sudo systemctl restart kagent
```

## Universal Agents Page UI

The Universal Agents page, accessible via **Settings****»** **Universal Agents** in the main nav menu, allows you to manage all the Universal Agents deployed by your organization.

Key UI elements include:

* **Auto-refresh/Refresh:** Automatically or manually updates the agent list.
* **Deploy Agent (Admin-only)**: Opens the **Deploy a Universal Agent** dialog for deploying new agents.
* **Show/Hide Filters:** Toggles the Filters pane (see**Universal Agents Filters**).
* **Search**: Restricts the **Universal Agents List** based on keywords you enter and/or filters you select

  + To clear all entered text, click the **X** in the Search field.
  + To clear a filter, click the **X** in a lozenge.
  > **Note:** For text-based searches, keywords are matched against the **Name** and **Capabilities** columns only.
* **Filters:** Provides controls to narrow the agents list (see **Universal Agents Filters**).
* **Universal Agents List:**A table displaying your deployed Universal Agents.

> **Note:** Member-level users can view Universal Agents, but only Admin-level users can add, edit, or remove them (see **About Users**).

### Universal Agents Filters

The **Filters** pane on the left of the **Universal Agents List** refines the list based on these options:

* **Capability**: Filters by installed capabilities (see **Agent Capabilities**).
* **Status**: Filters by agent status (e.g., “Up”).
* **Version**: Allows adding and removing filters for specific agent versions via a dropdown list.
* **Site**: Allows adding and removing filters for specific sites (see **About Sites**) via a dropdown list.

> **Notes:**
>
> * Add a Version or Site filter by selecting one or more values from the dropdown.
> * To remove any filter, click the **X** in the lozenge in the search field.

### Universal Agents List

The Universal Agents listshows all the deployed Universal Agents for your organization. To change the sort order of the list, click a heading to select a column, then click the resulting blue up or down arrow to choose the sorting direction (ascending or descending).

#### **List Columns**

The Universal Agentslistincludes the following columns of information about each deployed Universal Agent:

* **Status:** Agent's current operational state.
* **Name:** Agent's assigned name.
* **CPU:** Percentage of CPU utilized by the agent.
* **MEM:** Percentage of memory utilized by the agent.
* **Public IP:** Agent's public IP address.
* **Private IP:** Agent's private IP address.
* **Site:** Associated site name.
* **Version:** Agent's software version.
* **Capabilities:** Tags indicating agent functionalities.

#### List Actions

The Universal Agentslist includes the following action buttons per row:

* **Edit (pencil icon)**: Modify agent properties (see **Edit Agent Drawer**).
* **Remove (trash icon)**: Delete the agent instance (requires confirmation).

### Agent Details Drawer

![Agent details for fedora-temp-66, including status, IP addresses, and capabilities.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/UA-agent-details-drawer.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A35Z&se=2026-05-12T09%3A43%3A35Z&sr=c&sp=r&sig=83A8I68FwGVeRJrH6KsSRq88mTfw8%2FUCltx6zD4ebcY%3D)Access the Agent Details drawer by clicking an agent's row in the Universal Agents List.

#### Agent Details UI

The Agent Detailsdrawer includes the following UI elements:

* **Close (button):** Closes the drawer.
* **Name:** Displays the agent's name.
* **Edit (button):** Opens the **Edit Agent Drawer**.
* **Properties (fields):** A set of fields with metadata about the agent (see **Agent Details Fields**).
* **Capabilities (pane):** Lists the agent's capabilities in collapsible sections by **Capability Categories**, each with:

  + **Install (button)**: Installs a capability.
  + **Details (button):** Opens a detailed view of an installed capability (see **Capability Drawer**).
  + **Toggles (switch):** Turns a capability ON (blue) or OFF (grey).
* **Metrics (pane):** Charts showing the agent's CPU and memory utilization for the last hour (all capabilities combined).
* **View in Metrics Explorer (button):** Opens the view in Metrics Explorer.

> **Note:** When a capability section is collapsed, the total count of capabilities for that section is shown in a blue bubble.

#### Capability Categories

Capability categories shown in the Agent Details drawer are:

* **Installed**: Running or Paused on this Agent
* **Available**: Compatible and Supported
* **Available**: Compatibility not guaranteed
* **Upcoming**: Planned future release

#### Agent Details Fields

Agent properties are listed in the drawer and include the following:

* **Status:** The agent status (e.g., "Up").
* **Host Name:** The name of the host where the agent is deployed.
* **Version:** The Kentik-maintained version number for the agent.
* **Public IP:** The public IP address of the agent's host machine.
* **Private IP:** The private IP address of the agent's host machine.
* **Site:** The site to which the agent is associated, if any.
* **Description (only if specified):** A description string specified in the agent's settings
* **Closest Device:** A field with two states:

  + If the **Closest Network Device** field is already specified in the **Edit Agent Drawer**, this field shows that value.
  + If the value isn't already specified, this field is an Edit link which opens the **Edit Agent Drawer**, enabling you to specify the closest device.
* **Host OS:** The operating system and CPU type of the agent’s host machine.
* **Created on:** The date-time when the agent was created.
* **Last seen on:** The date-time when the agent last connected to Kentik.
* **Agent ID:** The ID number (integer) of the agent.
* **Install ID:** A unique ID (string) generated and stored on the host when the agent was installed.
* **Host Machine ID:** The unique ID (string) of the host, collected from the host OS.

### Edit Agent Drawer

Use the Edit Agent drawer, accessible via the **Edit** button in the **Agent Details Drawer**, to view and change details of your installed Universal Agents.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(647).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A35Z&se=2026-05-12T09%3A43%3A35Z&sr=c&sp=r&sig=83A8I68FwGVeRJrH6KsSRq88mTfw8%2FUCltx6zD4ebcY%3D)

The Edit Agent drawer includes the following UI elements:

* **Cancel (button):** Closes the drawer without saving any changes.
* **Save (button):** Saves changes to agent settings and exits the drawer.
* **Name:** Your organization's name for the agent (defaults to host name).
* **Description:** An optional, user-provided description of the agent.
* **Closest Network Device (SNMP/ST only):** A drop-down with which to specify the device in your network that is the fewest hops from the host on which the agent is installed.

  > **Note:** This setting facilitates topology-aware device status reporting in NMS.
* **Site:** A field showing the site, if any, to which the agent is currently assigned, which could be a data center or a VPC in a cloud provider (see **About Sites**). To set or change the site, click in the field to drop down a filterable list of sites, then click on a site in the list to select it.
* **Create a New Site (button):** If a site hasn't yet been defined for the location where the agent host is located, click the button to create a new site (see **Configure Site Fields**).

#### Configure Site Fields

When you click the **Create a New Site** button, the following fields are added to the **Edit Agent** drawer to enable you to enter the information required to create a new site:

* **Site Name:** A name for the new site.
* **Street Address:** The physical location of the new site, given as a street address.
* **Use an Existing Site:** If you decide not to create a new site, click this button to hide the **Configure Site** fields and instead choose a site from the **Site** drop-down.

### Capability Drawer

The **Capability** drawer, which opens by clicking **Details** on an agent capability, has the following UI elements:

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(648).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A35Z&se=2026-05-12T09%3A43%3A35Z&sr=c&sp=r&sig=83A8I68FwGVeRJrH6KsSRq88mTfw8%2FUCltx6zD4ebcY%3D)

* **Back/Cancel (button):** Closes the drawer without saving any changes.
* **Save (button):** Saves changes and exits the drawer.
* **Status:** The capability status (e.g., "Up").
* **Version:** The version number for the capability, as distinct from the agent version, e.g., "v1.109.0".
* **Binary:** The binary name installed for this capability (e.g., "kdns").
* **Run State:** A description of how long the capability has been enabled (e.g., "Running since 2025-05-28 10:00 as pid 762"). If capability is disabled, this will be blank.
* **Interface (DNS OTT Tap only):** Indicates which interfaces are included, e.g., "All", "eth0, "lo", "gre0", or "gre101" (see **Manage a Capability**).
* **Promiscuous Mode (DNS OTT Tap only):** "True" if promiscuous mode is enabled, otherwise "False" (see Edit a Capability).
* **Metrics chart(s):** Varying by capability as follows:

  + **SNMP/ST**: Metrics per second
  + **DNS OTT Tap**: DNS Responses per second, Discards per second
  + **SNMP Trap Receiver**: Traps per second, Receiver errors per second
  + **Syslog Server**: Messages per second, Dropped messages per second
* **View in Metrics Explorer:** Opens the view in Metrics Explorer
* **Devices (SNMP/ST only):** Lists the devices for which a SNMP/ST capability is collecting metrics, if any. Click the device name to open its details page. If no devices are found, "No devices found" is displayed.
* **Enable/Disable this Capability:** Click to enable/disable the capability.

> **Tip:** In addition to using the button in the **Capability** drawer, you can enable/disable capabilities with the toggle switches in the **Agent Details** drawer (see **Agent Details UI**).
