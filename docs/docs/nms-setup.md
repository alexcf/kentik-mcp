# NMS Setup

Source: https://kb.kentik.com/docs/nms-setup

---

This article discusses the initial setup of **NMS** (Network Monitoring System) in the Kentik portal.

> **Note:** For an overview of the initial Kentik setup tasks, see **Getting Started**.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(53).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A32Z&se=2026-05-12T09%3A31%3A32Z&sr=c&sp=r&sig=BVl730kYws89%2BgWWPfnNzXxwzuTm20fr0wBmYIPgfyk%3D)

*NMS Setup discovers devices for NMS monitoring.*

## About NMS Setup

Kentik’s NMS (Network Monitoring System) discovers and monitors devices within your infrastructure. Setup is managed via the **NMS Setup Wizard**, which configures the **Universal Agent** to collect metrics and transmit them to the Kentik SaaS.

### Access the NMS Setup Wizard

**New/Trial Customers**:

* Go to the **Welcome to Network Observability** page during onboarding.
* NMS setup is listed as one of the setup tasks on the left.

**Existing Deployments**:

* From the main portal nav menu, open the NMS Devices page.
* Select the Manage dropdown in the subnav and click **Discoveries** to open the **NMS Discoveries** page.
* Click **Discover Devices** to open the wizard.

### How NMS Polling Works

Unlike non-NMS applications involving SNMP, NMS requires the **Universal Agent** to be deployed within your network. The following default polling intervals apply:

| Metric Type | Polling Interval |
| --- | --- |
| Device Status, CPU, Memory, Interfaces | 1 minute |
| All other metrics | 5 minutes |

> **Note:** These intervals can be overridden by using Streaming Telemetry (ST) for supported devices (see **NMS via Streaming Telemetry**).

## NMS Setup Wizard

The setup process is divided into three main steps: deployment, credentialing, and discovery.

### Step 1: Deploy the Universal Agent

1. Navigate to the **NMS Setup** page.
2. Select your deployment method: Docker or Linux.
3. Copy the provided install command.

   * Optional*:* If using a web proxy, set the `https_proxy` variable as shown in the UI.
4. Run the command in the terminal on the host machine.
5. Once the agent appears in the Select an Agent list, click **Use this Agent**.

> **TIP**: The SNMP/ST capability of the Universal Agent will be installed automatically by the wizard. To install other capabilities for it, or any other agent, see **Install a Capability**.

### Step 2: Add Subnets & Credentials

1. In the**IP Range** field, enter a comma-separated list of IP ranges to scan.
2. From the **Credentials** dropdown, either select existing SNMP credentials or click **New Credential** and specify the following:

   * **Type:** SNMP version (v1, v2c, or v3).
   * **Credential Name**: Any user chosen name (no spaces).
   * **Community:** The SNMP community string configured on your devices.
3. Confirm the **Collection Agent** and click **Start Discovery**.

### Step 3: Discover and Import

The agent scans the provided ranges. Results populate in real-time.

1. Review the Discovered Devices list.

   > **TIP**: Click the customize button to choose which columns of information are shown in the list.
2. Uncheck any devices you do not wish to monitor.
3. Click **Add Devices** to begin data collection and open the **NMS Devices**page.

## Manage NMS Discoveries

Device discovery may take time depending on network size. You can monitor progress via the status indicator in the main navbar or by clicking **Manage Discoveries** on the **NMS Discoveries**page.

### Discovery Status Types

* **In Progress:** Scanning is active; an estimated time remaining is provided.
* **Finished:** All IPs have been scanned.
* **Started/Completed:** Timestamps for the discovery window.

## NMS via Streaming Telemetry

Streaming Telemetry offers a modern alternative to traditional SNMP polling.

> **Notes:**
>
> * Kentik also has limited support for collection via ST of data for the Kentik Data Engine (see **SNMP and ST KDE Support**).
> * For information about differences between SNMP and Streaming Telemetry, see **SNMP vs. Streaming Telemetry**.

### Comparison: SNMP vs. ST

* **Efficiency:** ST uses a subscription model, reducing CPU load on devices.
* **Granularity:** Supports sub-second data collection (defaulting to 30s in Kentik).
* **Accuracy:** Data is timestamped at the source, eliminating transit-time jitter.
* **Data Availability**: Some vendors (e.g., Juniper) are making new data types available only via ST.

> **Note:** For assistance with any obstacles to your adoption of ST, contact Kentik (see **Customer Care**).

### Supported ST Protocols

| Feature | Supported | Not Supported |
| --- | --- | --- |
| Transport | gRPC | TCP, UDP |
| Session | Dial-in (gNMI) | Dial-out |
| Encoding | JSON, Protobuf (GPB) | XML |

### Implementing ST (Two-Phase Process)

To use ST, a device must first be discovered via the **NMS Setup Wizard** to map its SNMP components.

1. **Device Config:** Configure your hardware (Cisco, Juniper, or Arista) for gNMI Dial-in (see Kentik’s **Configuration Repo** for vendor-specific examples).
2. ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(61).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A32Z&se=2026-05-12T09%3A31%3A32Z&sr=c&sp=r&sig=BVl730kYws89%2BgWWPfnNzXxwzuTm20fr0wBmYIPgfyk%3D)**Kentik Portal Config:** Navigate to Device Settings > **Streaming Telemetry** and enable **Dial-In from Universal Agent**.

### What Changes After Enabling ST?

Once enabled, the Universal Agent stops polling SNMP for CPU, memory, and interface stats, switching to ST intervals:

* **Default Metrics:** 30-second collection.
* **CPU Metrics:** 2-second collection (visible as live updates on the **NMS Devices** page).

  > **Note:** SNMP polling **continues** for non-telemetry data such as hardware make/model and BGP monitoring.

### View Results

ST data is normalized into the same schema as SNMP. You can view integrated results in the **NMS Dashboard**, **Metrics Explorer**, and **NMS Devices** pages. To verify if a device is actively using ST, check the **Dial-In from Universal Agent** checkbox status in the device's Settings dialog (see **Device ST Settings**).

## NMS Discoveries Page UI

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(55).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A32Z&se=2026-05-12T09%3A31%3A32Z&sr=c&sp=r&sig=BVl730kYws89%2BgWWPfnNzXxwzuTm20fr0wBmYIPgfyk%3D)

*A list of NMS discoveries.*

The Discoveries page provides a list of NMS discoveries, ongoing or completed, initiated from the **NMS Setup Wizard**, and includes the following:

* **Discover Devices** (button): Opens the **NMS Setup Wizard** to deploy agents, discover devices in specified IP ranges, and turn on monitoring for those devices. Discoveries started in the wizard are be listed on the Discoveries page.
* **Discoveries** (list): Lists a card for each discovery, including the following UI elements:

  + **Discovery status**: The status of the discovery:

    - **Finished**: All IPs specified in the **NMS Setup Wizard** have been scanned.
    - **In progress**: The scan of IPs has not completed. An estimate is provided of the time remaining.
  + **Started**: The date-time at which the scan began.
  + **Completed**: The date-time at which the scan finished (if completed).
  + **Completion percentage**: A horizontal bar and accompanying numeric indicator of the percent of the IPs that have been scanned.
  + **IPs scanned**: The number of IPs scanned.
  + **Devices found**: The number of devices found.
  + **View Details** (button): Opens the **Discovery Details Page** for this discovery.
  + **Dismiss** (button): Removes the discovery from the list.

    > **Note**: Discoveries stay on the list unless manually dismissed.

### Discovery Details Page

The Discovery Details page provides details about a discovery on the **NMS Discoveries** page, and includes:

* **Information Pane**: A set of indicators about the discovery. Click **Dismiss** to close the pane if desired.
* **Device List**: The list of devices discovered so far, including:

  + **Search**: Click to enter search terms by which to filter the devices list.
  + **Monitored device count**: The number of devices selected for monitoring.
  + **Add Devices**: Click to initiate NMS monitoring for the selected devices.
  + **Select all**: A checkbox to toggle selection of all devices on or off.
  + **Select device**: A checkbox to select an individual device for monitoring.
  + **Columns**: Information about each listed device (see **Discovery Details Columns**).
  > **Note:** By default, the Select device checkmarks for all found devices are on.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(58).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A32Z&se=2026-05-12T09%3A31%3A32Z&sr=c&sp=r&sig=BVl730kYws89%2BgWWPfnNzXxwzuTm20fr0wBmYIPgfyk%3D)

*The Details page for an NMS discovery*

#### Discovery Details Columns

The following columns are shown in the Device list for each discovery:

* **Status**: A set of lozenges indicating the status of the device:

  + **Up**: Indicates the device is intended to be operational.
  + **Down**: Indicates the device is intentionally not operational.
* **Name**: The device name specified during registration or editing.
* **IP Address**: The IP address(es) from which the device sends data to Kentik.
* **Object ID**: The device's globally unique identifier used in SNMP (sysObjectID).
* **Vendor**: The device manufacturer.
* **Model**: The device model number.
* **Device ID**: The Kentik-assigned device ID.
* **Monitoring Template**: A dropdown to select a monitoring template for the device:

  + Select one of Kentik’s presets (e.g., Everything, Minimal) or one of your own custom templates.
  + Click **Create NMS Template** to open the dialog to create a new template.
  + Click **Manage Templates** to open the Manage Templates page.
