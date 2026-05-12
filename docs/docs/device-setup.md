# Device Setup

Source: https://kb.kentik.com/docs/device-setup

---

This article explains setting up a device in the Kentik portal for **new/trial** customers.

> **Notes:**
>
> * To add/manage devices on an ongoing basis, use the Settings » **Networking Devices** page from the main nav menu.
> * For an overview of initial setup tasks, see**Getting Started****.**

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(38).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A02Z&se=2026-05-12T09%3A33%3A02Z&sr=c&sp=r&sig=Lz9Anw8tB1jCFknXfBDWlD1klDIHltaj383G5ss16vY%3D)

*Device setup is accessed from the Tasks List on the Welcome page.*

## About Device Setup

New/trial customers go through device setup as part of getting up and running with Kentik (see **Getting Started**). These tasks are listed at the left of the **Welcome to Network Observability** page. The process involves working through a wizard that gathers needed information. Separate from the Kentik portal, you’ll also need to configure the device itself to collect flow data and pass it to Kentik.

#### Starting Device Setup

To begin the device setup task:

1. Click **Connect a Router** in the Setup Tasks list.
2. Click **Set Up a Router** in the main display area to open the Connect a Router wizard.

> **Note:** Information entered on pages with a **Save** button is saved, but not on other pages if you use the **Back** button to go backwards in the wizard.

## Flow Export Method

On the first page of the wizard, choose how you'd like the router's flow data (e.g., NetFlow, sFlow) to get to Kentik. Click **Next**.

Two options are available:

* **Direct Flow Export**: Send flow records directly to Kentik through the public Internet.
* **kproxy Export**: Use the Flow Proxy capability of Kentik's **Universal Agent**, which:

  + Collects flow records and SNMP for multiple routers.
  + Encrypts the data locally before forwarding it to Kentik.

       ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(40).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A02Z&se=2026-05-12T09%3A33%3A02Z&sr=c&sp=r&sig=Lz9Anw8tB1jCFknXfBDWlD1klDIHltaj383G5ss16vY%3D)

## Direct Flow Export

The direct flow export path of the device setup wizard is covered in the following topics.

### Device Setup Steps

If you chose “Direct flow export” or “Host agent export” and completed the steps in **Host Agent Export**,the next stage begins on the **Register your flow source** page.

The page lists steps for gathering information to complete device registration:

* **General Info**: Device name, description, make, model, and location.
* **Flow**: Technical details like IP addresses and flow sampling rate.
* **SNMP**: Information for enabling polling, including SNMP version and IP address.

  + SNMP polling enriches your Kentik-stored flow records with additional interface-specific data points that are queryable as dimensions (from information OIDs) and metrics (from counter OIDs).
* **BGP**: Information for enabling flow data correlation with BGP routing data.

  + Use this to apply customer-specific filters to query results, create routing-based tags, and run AS-oriented traffic analysis (including peering analytics).

Each category is represented in the device settings workflow as pages that gather required information and that provide instructions for configuring the device to send data to Kentik.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(41).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A02Z&se=2026-05-12T09%3A33%3A02Z&sr=c&sp=r&sig=Lz9Anw8tB1jCFknXfBDWlD1klDIHltaj383G5ss16vY%3D)

*The ‘Register your flow source’ page includes a progress list that shows what setup steps remain for this device.*

#### Setup Progress List

As you complete tasks, you'll be redirected to the **Register your flow source** page, where you’re progress is displayed. Completed tasks are indicated by a checkmark in the left list column and a link changing from **Get Started** to **Edit** in the right column.

When all tasks are complete, the **Create Device** button becomes active, allowing you to save settings and finalize device registration.

> **Note:**If you encounter questions or need assistance, click the **Get Help** link at the upper right of any onboarding screen.

### Device General Information

Device general information settings are accessed from the **Register your flow source** page:

1. Click the **Get Started** link in the **General Info** row.
2. Enter a name and description for your device on the Name Your Device page. Click **Next** when finished.
3. On the **Device make and model** page:

   1. Click the card corresponding to your device’s manufacturer.
   2. If there is more than one series or model, choose the appropriate model of your device from the **Model** popup.
   3. If no card is found, click the **Other** card.

      > **Note:** Once you've specified your make and model, the page will advance to the next step.
4. On the **Set device location** page, enter the physical location (e.g., street address of a data center). We'll use this information to create a "site" for location-based queries (see **About Sites**). Click **Finish** to return to the **Register your flow source** page.

### Device Flow Setup

After entering **Device General Information****,** you can proceed to the flow settings:

1. Click the **Get Started** link in the **Flow** row to open the **Sending IP addresses** page.
2. Enter a comma-separated list of IP addresses from which the device will export flow data to Kentik. Click **Next** when finished.
3. Set the rate at which the device generates flow records on the **Flow Sample Rate** page. Click **Next** when finished.

   1. This is the ratio of collected flows to all flows (e.g., “1:1000” means one flow record per thousand flows).
   2. The accuracy of traffic volume calculations depends on this information.
4. Configure the device to send flow data to Kentik using the configuration snippets provided on the **Configure Flow on Device** page based on the device’s make and model.
5. Click **Finish** to return to the **Register your flow source** page.

### Device SNMP Setup

After completing **Device Flow Setup**, move on to the SNMP settings:

1. Click the **Get Started** link in the **SNMP** row.
2. Activate SNMP polling for this device and choose the SNMP version (SNMP v3 is more secure and is recommended for customers concerned about using SNMP v2 over the public Internet). Click **Next** when finished.
3. On the **SNMP Details** page, click **Next** after making the following settings (see **Device SNMP Settings**):

   1. SNMP v2: IP Address and Community.
   2. SNMP v3: IP Address and User Name (both required), and optionally, Authentication and Privacy (see **About SNMP V3**).
4. On the **Configure SNMP on Device** page, you’ll find configuration snippets based on the device’s make and model to allow SNMP polling from Kentik. After configuring the device, click **Finish** to return to the **Register your flow source** page.

### Device BGP Setup

After completing **Device SNMP Setup**, move on to the BGP settings:

1. Click the **Get Started** link in the **BGP** row.
2. Activate BGP for this device and specify the IP type (e.g., IPv4) of the device's peering address. Click **Next** when finished.
3. Set the **Device BGP Settings**and click **Next**.
4. Configure the device for BGP peering with Kentik using the configuration snippets provided on the **Configure BGP on Device**. Click **Finish**.
5. Upon returning to the **Register your flow source** page, you'll see checkmarks at the left of each setup item and an **Edit** link at the right. The **Create Device** button at the bottom of the page will be activated. Click it to save your settings and finalize device registration.

### Configuration Snippets

Configuration snippets are pre-written pieces of configuration or commands that help complete your device setup. They allow you to quickly configure devices to export flow, SNMP, and BGP data. Each snippet is specific to the device selected during initial onboarding (see **Device General Information**).

> **Note:** Kentik’s **public configuration repository** in GitHub contains the snippets. Kentik customers can explore, branch, and submit pull requests for updates.

## Host Agent Export

Exporting device flow to Kentik via a host software agent enables you to collect flow + SNMP for multiple routers with one agent, Kentik’s **Universal Agent**. This involves the following additional steps in the Kentik portal:

1. Choose **Host Agent Export** during device setup wizard (instead of Direct Flow Export)
2. Install the Universal Agent on your host machine
3. Configure Flow Proxy Capability:

   1. Go to Settings » **Universal Agents**
   2. Click your agent » Capabilities » Flow Proxy » Details » Edit
   3. Configure the following settings, then click **Save**:

      1. **Host**: 0.0.0.0 (listen on all interfaces)
      2. **Port**: 9995 (default UDP port)
      3. **Site**: Associate with your site
4. Configure Network Devices:

   1. Point routers/switches to export flow to the Universal Agent host IP on port 9995.
5. After setting up your host agent, you're taken to the **Register your flow source** page (see **Device Setup Steps**), where the setup steps for devices are the same whether you’re using Direct Flow or Host Agent export.

## Setup in Multiple Sessions

Onboarding a device with Kentik doesn’t require a single session. If you log out before your organization's device settings are complete, you’ll pick back up from the following place, depending on where the process was left:

* **Data source not started**: You'll be taken to the initial post-login page for new accounts, entering the new customer Setup workflow (see **Getting Started**).
* **Data source incomplete**: If your first data source is a device and you started onboarding but didn’t finish, you'll be taken directly to the **Register your flow source** page for the partially onboarded device, continuing the setup (see **Device Setup Steps**).
* **Data source complete**: Upon completing registration and settings for at least one data source, you'll be taken to **Network Explorer**. If you initiated registration for multiple data sources and any device settings are incomplete, a callout will notify you, and you can proceed with **Using Device Status**.

> **Note:** The main menu will include a **Setup Tasks** section until all setup tasks are complete. Click the **See all** link to view your organization's remaining **Setup Tasks**.
