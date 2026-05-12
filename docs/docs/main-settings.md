# Main Settings

Source: https://kb.kentik.com/docs/main-settings

---

This article covers the main settings for the Kentik portal.

> **IMPORTANT:** New/trial Kentik customers should refer to the initial Setup Tasks covered in **Getting Started**.

![Dashboard showing device configuration status and alert policies for Kentik settings.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/MS-main.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A42Z&se=2026-05-12T09%3A31%3A42Z&sr=c&sp=r&sig=pQgTYJxVZJgVcxiIXR5unprqH1y8dhaefmahiApHUgs%3D)

*The Settings page provides access to the individual settings modules for the Kentik portal.*

## About Settings

Settings are the administrative controls that enable you to tailor your Kentik setup to make it work most effectively for your organization. These settings are accessed via the following locations in the Kentik portal:

* **Main settings**: The primary settings for daily use of the portal are accessed from the portal's Settings page, which is opened from the link in the left column of the main navbar. These settings are covered in the remaining topics of this article.
* **Organization and User settings**: These settings are accessed from the **Organization Settings** menu and the **User** menu, both of which drop down from the portal's main navbar. For information on these settings, see **Org & User Settings**.

> **Note**: While member-level users can access most of the pages linked to from the **Settings** page, most will only allow admin-level users to manage Kentik settings.

## Settings Categories

The portal’s **Settings** page is laid out as panes that each contain a set of cards related to a category of settings. Each card contains a link to an individual settings page.

### Data Sources

A data source is any network entity that generates data that is ingested into the Kentik Data Engine (KDE), where it is monitored and analyzed. This pane includes links to the admin pages for two main categories of data sources:

* **Networking Devices**: A link to the **Devices** page for device settings and status. Devices are the routers, hosts, switches, firewalls, etc. that make up the physical infrastructure deployed in your sites such as data centers and branch offices. For Kentik to receive data (flow records, metrics, SNMP, BGP, etc.) from a given device, that device must be:

  + Configured (on the device) to connect with Kentik;
  + Registered with Kentik in the portal or via API.
* **Public Clouds**: A links to the **Public Clouds** page, where you can:

  + Create a "cloud export" to Kentik from your resources (e.g., VPCs and subnets) that are hosted by public cloud providers (e.g., AWS, GCP, Azure, or OCI).
  + Manage the settings for existing cloud exports.

> **Notes:**
>
> * An indicator for each data source type shows how many are configured.
> * If any of your data sources are incomplete, missing flow, or require additional configuration, the **Data** **Sources** pane will display a notification that includes a **Review your devices** button. Click the button to go to the **Device Status Page**.
> * For assistance with any aspect of configuring or registering a data source, contact **Customer Support**.

### Alerting

Kentik’s alerting system is implemented via alert policies which are essentially a set of comparative evaluations that, when matched, can trigger an alert, which results in an action such as a notification and/or a mitigation. This pane includes links to pages covering the following aspects of alerting:

* **Alert Policies**: A link to the **Policies** page, where you can add, edit, enable, disable, clone, debug, or remove alert policies, and also access the **Policy Templates** page.
* **Mitigations**: A link to the **Manage Mitigations** page, where you can configure mitigations to prevent undesirable traffic (e.g., a DDoS attack) from disrupting network availability.
* **Notifications**: A link to the **Notification Channels** page, where you can specify how users are notified for important events such as the triggering of an alert (see **Threshold Notifications**). Each notification channel represents a notification mode (e.g., email) and one or more notification targets (e.g., a set of email addresses).
* **Alert Suppressions & Silences**: A link to the **Alert Suppressions & Silences** page, which lists “patterns” that each represent a set of conditions (dimension/value pairs). Alerts will not be generated on traffic that matches those patterns.
* **Auto-acknowledgements**: A link to the **Auto-acknowledgements** page, which lists all active and expired alert auto-acknowledgements and enables you to manage them (edit or remove comments, change expiry dates, or delete them entirely).

  > **TIP**: Auto-acknowledgements are created from the Auto Acknowledge dialog on the Alerting page and cannot be created on the Auto-Acknowledgements page.

### Agents

Kentik uses software agents to gather data and perform tasks to support operation of the Kentik platform (see **About Kentik Agents**). Some agents are configurable only in their own CLIs, but others may be managed via portal pages in the Settings menu:

* **Universal Agents**: Manage your organization’s agents and view their statuses.
* **Synthetics Agents**: Manage your organization’s private instances of `ksynth`, Kentik’s agent for synthetic testing (see **Kentik Synthetics Agents**).

> **Note:** The Universal Agent will gradually replace all other Kentik agents.

### Network Metadata

Network metadata enables Kentik to help you better understand the structure of your network and how your traffic moves through it, providing better context and value. This pane includes links to administrative pages for the following types of network metadata:

* **Interfaces**: A link to the **Interfaces** page, where admin-level users can add and remove interfaces on your organization's Kentik-registered devices and edit the settings of those interfaces.
* **Sites**: A link to the **Sites** page, where you can get information about the physical location (for devices) or cloud region (for public clouds) of your data sources. An indicator at right shows how many of your organization's data sources are already assigned to a site.
* **Site Markets**: A link to the **Site Markets Page**, where you can add and edit site markets (logical groupings of sites with common characteristics). An indicator at right shows how many site markets are already configured.
* **Interface Classification**: A link to the **Interface Classification** page, where you can establish network boundaries and categorize connectivity based on interface metadata. An indicator at right shows the percentage of your organization's interfaces that have already been classified.

  > **Note**: Only admin-level users can access the Interface Classification page.
* **Network Classification**: A link to the **Network Classification** page, where you can define your IP space and ASNs, allowing us to directionally classify traffic in relation to your network.
* **Labels**: A link to the **Labels** page, where you can create and edit labels. You can use labels to treat any combination of devices, dashboards, saved views, policies, tests, or agents as a logical group across Kentik modules.

  > **Note**: Labels created on the Labels page are applied to a given network entity (e.g., device) on the settings page for that entity type (e.g., **Devices** page).
* **Saved Filters**: A link to the **Saved Filters** page, where you can create and manage a saved set of filter conditions that you can store and apply later.

### Data Enrichment

Data enrichment enables your organization's flow data to be correlated with additional information that provides a richer picture of your network and a better understanding of traffic patterns:

* **AS Groups**: A link to the **AS Groups** page, where you can assign Autonomous Systems (ASes) to a group whose traffic will be summed for the purpose of top-X evaluations and filtering in queries.
* **Custom Dimensions**: A link to the **Custom Dimensions** page, where you can add business context to your organization's network data by supplementing Kentik's built-in dimensions (see **Dimensions Reference**) with dimensions that you define for your own specific needs.
* **Custom Applications**: A link to the **Custom Applications** page, where you can define the characteristics — specific values for protocol, port number, IP address, and/or ASN — of an application that will be associated with all traffic sharing those characteristics.
* **Custom Geos**: A link to the **Custom Geos** page, where you can map your traffic with business context by assigning countries to custom groups for querying via group-by or filter dimensions from the Custom Geo dimension family.
* **Flow Tags**: A link to the **Flow Tags** page, where you can create tags that become part of the flow record for traffic that matches tag-defined criteria, enabling you to query for traffic matching a given tag.

  > **Note**: Only admin-level users can access the Flow Tags page.
