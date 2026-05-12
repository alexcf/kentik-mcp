# CDN Analytics

Source: https://kb.kentik.com/docs/cdn-analytics

---

This article describes the **CDN Analytics** page of the Kentik portal.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(181).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A39Z&se=2026-05-12T09%3A50%3A39Z&sr=c&sp=r&sig=gBVJqRpeCLCYUlZxi0HY%2FxMS1TVXQFrvihJZehdjH6E%3D)

*The landing page of the CDN Analytics workflow.*

## About CDN Analytics

|  |  |
| --- | --- |
| **Purpose** | Enable “eyeball” ISPs to track and optimize CDN traffic delivery and performance to subscribers while providing data-driven analytics for interconnection negotiations. |
| **Benefits** | * Improve subscriber support by facilitating investigation of performance issues. * Boost subscriber retention. * Determine the optimal balance between performance and cost. * Explore data without the limitations of pre-aggregated reports. * Real-time results without lengthy waits. * Underlying dataset is available for alerting. |
| **Use Cases** | * Understand precisely how your network connects with each CDN. * See routing changes for any CDN that will impact cost or subscriber performance. * Compare embedded CDN caches to CDN traffic from outside your network. |
| **Relevant Roles** | Network/NOC Engineer, Peering Manager |

### CDN Analytics Overview

Kentik’s CDN Analytics workflow is primarily of interest to “eyeball” ISPs, whose subscribers are the consumers of vast amounts of content transported by Content Delivery Networks (CDNs). To optimize a network for this heavily inbound traffic profile, an ISP needs to understand the nature of CDN-transported traffic (video, gaming, social, etc.) as well as how this traffic is handed over to the network (at which cost and at what performance).  
  
The basic attributes available from traditional flow telemetry are of little help in this regard, and the picture is further complicated by the fact that content providers typically use multiple CDNs, each of which dynamically modifies their methods and routing in response to changing constraints such as capacity, cost, and weather.  
  
By enriching flow records with data derived via non-flow techniques, Kentik provides eyeball ISPs with unsurpassed visibility into their CDN traffic.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(182).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A39Z&se=2026-05-12T09%3A50%3A39Z&sr=c&sp=r&sig=gBVJqRpeCLCYUlZxi0HY%2FxMS1TVXQFrvihJZehdjH6E%3D)

*CDN Analytics enables detailed visibility into traffic that reaches your network via CDN.*

### CDN Attribution Dimensions

CDN Analytics is based on Kentik technology that allows us to determine whether a given flow record originated or terminated with a commercial CDN, and to store that information for each record using the following two dimensions (see **Application Context and Security**):

* **Source CDN**: The commercial name of the CDN derived from the source IP of an ingested flow.
* **Destination CDN**: The commercial name of the CDN derived from the destination IP of an ingested flow.

  > **Note:** This dimension enables you to track "fill traffic" that is pointed toward a CDN server to fill a local cache.

Once stored in the Kentik Data Engine (KDE; see **KDE Tables**), the columns (`src_cdn, dst_cdn`) corresponding to these dimensions can be used in the Kentik queries underlying the CDN Analytics workflow.

### CDN Types

Each Kentik-identified CDN is defined as belonging to the following CDN types (a given CDN may be more than one type):

* **Cloud**: Commercial CDN offerings from public cloud providers, including CloudFront (AWS), Google CDN (GCP), and Azure CDN (Microsoft).
* **Commercial**: Global content delivery services (including Akamai, Level3, Limelight, EdgeCast, Fastly, and Cloudflare) that are sold to content providers.
* **Content**: Purpose-built CDNs that are owned and operated by content providers for their own primary use, including Netflix’s Open Connect (OCA) program, Facebook’s FNA program, and Google’s GGC Program.
* **Telco**: CDNs owned by large telecommunication providers (carriers) that also sell connectivity, colocation, and/or transit, including Level3, Comcast, AT&T, Telefonica, and Tata Communications.

## CDN Analytics Page

The CDN Analytics page is covered in the following topics.

### CDN Analytics Page UI

The CDN Analytics page includes the following main UI elements:

* **Subnav**: A silver horizontal strip below the navbar that includes the controls described in **CDN Analytics Subnav**.
* **Title pane**: The set of controls and indicators below the subnav, which are described in **CDN Analytics Title Pane**.
* **Display area**: The main display area of the page, which contains the indicators, charts, and tables described in **CDN Analytics Display Area**.

#### CDN Analytics Subnav

The silver horizontal strip below the navbar includes the following UI elements:

* **Breadcrumbs**: An indicator of your current location within the Kentik portal. As you drill down deeper, you can click on a breadcrumb to go back to a higher level.
* **Share**: A button that opens the Share dialog so you can share the current view (see **Sharing via the Share Dialog**).
* **Actions**: A drop-down from which you can choose **Export** to download a visual report (PDF) covering the page’s visualizations and tables. A notification appears when the PDF is ready to download.

#### CDN Analytics Title Pane

The page title bar, located under the subnav, includes the following elements:

* **Favorite** (star): An icon that indicates whether this page is currently designated as a favorite (see **Library Bookmarks**) and toggles the designation on/off when clicked.
* **Title**: CDN Analytics (the name of the page).
* **Time Range selector**: Indicates the duration, back from the present, for which the information displayed on this page has been calculated.
* **Configure**: Opens the **CDN Analytics Configuration** page.

#### CDN Analytics Display Area

The main body of the page, located below the title bar, contains the following visualizations:

* **Peak Traffic**: An indicator stating the peak traffic rate that occurred during the selected time range.
* **Traffic chart**: A sparkline showing the overall traffic rate over the selected time range.
* **Did you know pane**: A set of indicators with relevant information about your organization’s CDN traffic (see **Did You Know Pane**).
* **Embedded Caching pane**: A set of indicators, one for each of the primary CDNs using embedded caching, that show the CDN’s cache-sourced traffic as a percentage of their overall traffic (see **Embedded Caching Pane**).
* **CDN Types field**: A field that allows you to filter the **CDNs List** by CDN type (see **CDN Types Filtering**).
* **CDNs List**: A list of the CDNs from which your network is receiving traffic (see **CDNs List**).

### Did You Know Pane

The **Did You Know** pane highlights useful and relevant statistics regarding the CDN traffic on your network. The following indicators are included:

* **Top 5 CDNs**: The percentage of traffic to your network’s subscribers that comes either directly from the top five Kentik-identified CDNs or from caches embedded in your network by those CDNs. A list at the bottom of the indicator provides links to each CDN's **CDN Details Page**.
* **Non-CDN Traffic**: The rate of traffic to your network’s subscribers that does not come from either a Kentik-identified CDN or an embedded cache, in Gbits/s. Also included is the percentage of total subscriber traffic that this non-CDN traffic represents.
* **Connectivity Optimization**: The percentage of your traffic from Kentik-identified CDNs that is delivered via transit, which is typically the costliest and least performant approach.
* **CDNs Mapped**: The current number of Kentik-identified CDNs and the date on which the Kentik CDN engine was last updated.

### Embedded Caching Pane

To improve performance for popular high-bandwidth content, several of the major CDNs (e.g., Facebook Appliance, Google Global Cache, Netflix Open Connect, Akamai Caches, etc.) install cache appliance servers in ISP datacenters, as close as possible to content consumers (ISP subscribers). Content is served from these caches and refreshed via proxy feed during off-peak hours. Using **Interface Classification**, Kentik identifies the interfaces on your network that connect to these caches and assigns a Connectivity Type of “Embedded Cache” to inbound traffic on those interfaces.  
  
The **Embedded Caching** pane includes a panel for each of the primary CDNs that use embedded caching on your network, if any. Each panel shows the following information:

* **CDN logo**: Click the logo to view the **CDN Details Page** for that CDN.
* **Offloaded percentage**: The CDN’s embedded cache traffic as a percentage of its overall traffic.
* **Peak cache performance**: Maximum bits per destination IP going to subscribers from the CDN's embedded cache.
* **View Dashboard**: Click this button to open the Embedded CDN Traffic Analysis dashboard for that CDN.

### CDNs List

The CDNs List is a table showing the CDNs from which your network is receiving traffic. The list is filterable using the **CDN Types** field (see **CDN Types Filtering**). The table can be sorted by clicking the heading of the desired sort column; clicking the same heading again will change the sort direction (ascending vs. descending).  
  
 The CDNs List includes the following columns:

* **Traffic chart**: A sparkline showing the rate of traffic matched to this CDN over the currently selected time range.
* **Source CDN**: A Kentik-identified CDN. Click the name to go to the **CDN Details Page**.
* **Peak Subscriber Traffic from CDN**: The peak bitrate of traffic (in Gbits/s) from the CDN to subscribers during the selected time range.
* **Max Unique Subscribers**: The peak count of subscribers simultaneously receiving content from this CDN during the selected time range (calculated from the maximum unique Destination IPs).
* **Max Mbits/s per Subscriber**: The highest bitrate, averaged across subscribers, for traffic received from the CDN during the selected time range. The value is calculated by comparing the bitrate for each interval in the time range divided by the number of subscribers (unique Destination IPs) receiving data from that CDN during that time.
* **Last Datapoint**: The bitrate (in Gbits/s) of subscriber traffic at the last datapoint in the selected time range.
* **Last Datapoint Unique Subscribers**: The count of subscribers simultaneously receiving content from this CDN at the last datapoint in the selected time range.
* **Last Datapoint Mbits/s per Subscriber**: The bitrate for the last interval in the selected time range divided by the number of subscribers receiving data from that CDN during that time.

> **Notes:**
>
> * The “subscriber traffic” shown in each row of the table is traffic to your network’s subscribers that comes either directly from the CDN or from caches embedded in your network by the CDN.
> * The time slices used for data aggregation vary depending on the duration of the selected time range.

#### CDN Types Filtering

The CDNs List is filtered via the **CDN Types** field, which shows a lozenge for each CDN type that's been selected for inclusion in the list:

* When no types are selected, CDNs of all types appear in the list.
* To select a type, click in the field and choose the type from the resulting popup.
* To remove a type, click the **X** in the type's lozenge.

## CDN Analytics Configuration

The CDN Analytics Configuration page includes the following tabs.

### Landing Page Tab

The **Landing Page** tab allows you to choose the cache-embedding CDNs (if any) for which you want the offload percentage to be displayed in the **Embedded Caching Pane** of the CDN Analytics landing page.  
![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(224).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A39Z&se=2026-05-12T09%3A50%3A39Z&sr=c&sp=r&sig=gBVJqRpeCLCYUlZxi0HY%2FxMS1TVXQFrvihJZehdjH6E%3D)

If any cache-embedding CDNs are configured and enabled, this tab includes a horizontal pane for each CDN with the following elements:

* **CDN logo**: The CDN's logo, displayed on the left.
* **Display offload**: A toggle switch, displayed on the right, that determines whether the offload percentage for this CDN will be included in the **Embedded Caching Pane** on the CDN Analytics landing page:

  + **When On (blue and to the right)**: the percentage will be shown.
  + **When Off (gray and to the left)**: the percentage will not be shown.

> **Note**: Clicking anywhere in a pane toggles the switch for that CDN.

If no cache-embedding CDNs are detected by Kentik (see **Detected Embedded Caches**), nor have any been added manually (see **Additional Embedded Caches**), then this tab displays only the message “No Embedded Caches found.”

### Detected Embedded Caches

The **Detected Embedded Caches** tab lists the embedded caches that we’ve automatically detected based on the rules configured in the portal's **Interface Classification** module. If IC assigns a provider to an interface, the detection engine will attempt a loose match with any known CDNs.

> **Note**: If a cache doesn’t display on the **Detected Embedded Caches List**, you can add it manually on the **Additional Embedded Caches** tab.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(225).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A39Z&se=2026-05-12T09%3A50%3A39Z&sr=c&sp=r&sig=gBVJqRpeCLCYUlZxi0HY%2FxMS1TVXQFrvihJZehdjH6E%3D)

#### Detected Embedded Caches UI

The **Detected Embedded Caches** tab includes the following main UI elements:

* **Interface info**: A red pane that appears at the top of the tab If any of your interfaces are without a CDN provider. The pane's **Only show interfaces without a provider** switch limits the Detected Embedded Caches list to interfaces without a provider.
* **Group By**: A dropdown from which you can choose how to group the rows in the list: None (no grouping, the default), Device, Enabled, or CDN.
* **Search**: A field that filters the Detected Embedded Cache list to show only rows that contain the entered text in the Subnet/Mask, Interface, or CDN Provider columns.
* **Detected Embedded Caches List**: A table listing the embedded caches Kentik has automatically detected (see **Detected Embedded Caches List**).

#### Detected Embedded Caches List

The **Detected Embedded Caches** list provides information about any embedded caches automatically detected by Kentik. The table can be sorted by clicking the heading of the desired sort column; clicking the same heading again will change the sort direction (ascending vs. descending).  
  
 The Detected Embedded Caches list includes the following columns:

* **Subnet/Mask**: The IP range detected for the listed interface.
* **Interface**: The name and description of the detected interface.
* **Device**: The name and location of the detected device.
* **CDN Provider**: The CDN name where your cache is embedded (e.g., AWS or Google) as entered in the Provider/Customer field of the Edit Interface dialog (see below). If that field is blank, the CDN Provider column will read “Undetermined”, and the row will be marked with a red bar at its left border.
* **Enabled**: A toggle switch that determines whether the traffic for this interface will be included along with any other cache traffic for this CDN Provider in the provider's pane on the **Landing Page** tab of the Configuration page:

  + **When On (switched right)**: the traffic will be included if the value in the CDN Provider column is not "Undetermined."
  + **When Off (switched left)**: the traffic will not be included.

    > **Note:** If the provider associated with an interface is not automatically identified via **Interface Classification** it may be manually specified in the Edit Interface dialog.
* **Edit**: Opens the Edit Interface dialog to modify that interface’s settings (see **Interface Field Definitions**).

### Additional Embedded Caches

The **Additional Embedded Caches** tab enables you to manually add or edit any embedded caches that Kentik hasn’t detected automatically. The reasons that a cache might not have been detected include:

* There are missing Interface Classification rules.
* The cache is connected to a router that is not registered with Kentik.

  > **Notes**:
  >
  > + Once a cache is added, its IPs will be classified as Source CDN or Destination CDN.
  > + The situations in which embedded caches may not be detected correctly are explained in **Corner cases for Embedded Cache detection**.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(226).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A39Z&se=2026-05-12T09%3A50%3A39Z&sr=c&sp=r&sig=gBVJqRpeCLCYUlZxi0HY%2FxMS1TVXQFrvihJZehdjH6E%3D)

#### Additional Embedded Caches UI

The **Additional Embedded Caches** tab includes the following main UI elements:

* **Add Cache Subnet**: A button that opens the Add Cache Subnet dialog (see **Cache Subnet Dialogs**) where you can manually add a new embedded cache.
* **Group By**: A dropdown from which you can choose to group the rows in the list by CDN or not at all (None).
* **Search**: A field that filters the **Additional Embedded Caches** list to show only rows that contain the entered text in the Subnet/Mask, Embedding CDN, or Description columns.
* **Additional Embedded Caches List**: A table listing any manually-added embedded caches (see **Additional Embedded Caches List**).

#### Additional Embedded Caches List

The **Additional Embedded Caches** table lists your organization's manually added embedded caches. The table can be sorted by clicking the heading of the desired sort column; clicking the same heading again will change the sort direction (ascending vs. descending).  
  
 The table includes the following columns:

* **Subnet/Mask**: The IP range entered by the user for the cache.
* **Embedding CDN**: The logo of the CDN for which this cache is embedded.
* **Description**: A user-defined description of the cache.
* **Edit** (pencil icon): Opens the Edit Interface dialog in order to modify that interface’s settings (see **Cache Subnet Dialogs**).
* **Delete** (trash icon): Opens a confirming dialog that allows you to remove the cache from the **Additional Embedded Caches** list.

> **Note**: If your organization has no manually added embedded caches, the table displays the message “No Embedded Caches found” and an **Add** button that is identical to the **Add Cache Subnet**button.

#### Cache Subnet Dialogs

The Add Cache Subnet and Edit Cache Subnet dialogs share the same layout and the following common UI elements:

* ![Form for adding a cache subnet with fields for subnet, CDN, and description.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/CDN-add-cache-subnet-dialog.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A39Z&se=2026-05-12T09%3A50%3A39Z&sr=c&sp=r&sig=gBVJqRpeCLCYUlZxi0HY%2FxMS1TVXQFrvihJZehdjH6E%3D)**Close**: Both the **X** in the upper right corner and the **Cancel** button close the dialog and restore all elements to their values at the time the dialog was opened.
* **Subnet/Mask**: The IP prefix/range for the embedded cache.
* **Embedding CDN**: The CDN in which this cache is embedded. Click the field to open the drop-down menu and select the appropriate CDN name.
* **Description**: A description of the cache subnet being added. This text will display under the **Description** column in the **Additional Embedded Caches** list.
* **Add Cache Subnet** (Add Cache Subnet dialog only): Saves the cache subnet to the **Additional Embedded Cache** list and closes the dialog.
* **Save** (Edit Cache Subnet dialog only): Updates the cache subnet and closes the dialog.

### Advanced Filtering

> **Note:** Kentik recommends that customers make no modifications to this tab without first checking with their **Customer Success** representative.

## CDN Details Page

The details pages for individual CDNs are covered in the following topics.

### CDN Details Page Layout

The details pages for individual CDNs include the following main UI elements:

* **Filters**: A selector that enables you to choose one or more dimensions by which to narrow the traffic displayed in the charts and tables of this page (see **Filters Control**).
* **Time Range**: A dropdown to choose the duration, back from the present, for which the information displayed on this page has been calculated (see **Time Range Selector**).
* **Information pane**: Useful information about the CDN, including its type and ASNs (see **CDN Information Pane**).
* **Offload Profile pane**: A breakdown of the types of connections over which your subscribers are receiving content from the CDN (see **Offload Profile**).
* **CDN Transport Breakdown**: A diagram and table — by connectivity type, provider, and site — of the traffic transported to your network by this CDN (see **CDN Transport Breakdown**).
* **Top OTT Services**: A chart and top-X list showing the individual “over-the-top” (OTT) content services via which traffic is coming to your network from the CDN (see **Top OTT Services**).
* **Performance Analysis**: A line chart of the bitrate per subscriber for the selected CDN service over the selected time range (see **Performance Analysis**).

### CDN Information Pane

The **Information** pane is the topmost pane on the page and presents general information about how the CDN operates and what impact it is having on your network. The pane is divided into three parts:

* **CDN Types**: Shows the CDN categories (see **CDN Types**) assigned to this CDN by Kentik.
* **Owned and Operated ASNs**: Lists the Autonomous Systems controlled by the CDN.

  + Click on an ASN in the list to go to that ASN’s Details page (see **Core Details Pages**).
  + Click **Lookup PeeringDB Records** to see that CDN’s records on **PeeringDB.com**.
* **CDN monitoring status**: Shows if a synthetics test exists for this CDN (see **CDN Monitoring Status**).

#### CDN Monitoring Status

The **CDN Monitoring Status** portion of the **CDN Information Pane** shows if a synthetics test exists for this CDN, and if so, how many agents are being used.

* If the CDN is not currently monitored, you'll see a **Start Monitoring** button, which opens the Start Monitoring dialog, which is identical to the **Configure Monitoring Dialog**.
* If the CDN is currently being monitored, you'll see the following buttons:

  + **View Details**: View and edit the test in the **Tests List** of the Test Control Center.
  + **Configure**: Open the **Configure Monitoring Dialog**.
  + **Stop Monitoring**: Open a confirmation dialog in which you can click the Stop button to stop monitoring the CDN and delete the test from the Test Control Center.

    > **Notes**:
    >
    > - A monitoring status of enabled tells you that a test exists for the CDN, but not whether that test is currently active or paused (to find out, click **View Details**).
    > - CDN tests are part of the Autonomous category of Network tests in Kentik Synthetics (see **Network: Autonomous Tests**.

#### Configure Monitoring Dialog

![Configuration options for Akamai monitoring with selected agents and expected credit usage.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/CDN-configure-monitoring-dialog-akamai.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A39Z&se=2026-05-12T09%3A50%3A39Z&sr=c&sp=r&sig=gBVJqRpeCLCYUlZxi0HY%2FxMS1TVXQFrvihJZehdjH6E%3D)The Configure Monitoring and Start Monitoring dialogs allow you to configure the agents from which you monitor a CDN. These dialogs contain the following UI elements:

* **Close**: The **X** in the upper right corner closes the dialog and restores all settings to their values at the time the dialog was opened.
* **Name**: An editable field for the name of the test.
* **Agent List**: A list of agents from which you can select one or more to test (see **Agents List**).
* **Expected monthly credit usage**: A bar graph that shows the impact of the test, as currently configured, on your organization's overall consumption of credits available for the current month (see **About Test Credits**). Also shown is the number of credits the test will use in the current month and the remaining number of credits that will be available in the current month if this test is active.
* **Start Monitoring** (Start Monitoring dialog only): A button that starts the test configured in the dialog. The button becomes active when the current settings are sufficient for Kentik to create a test.
* **Update Test** (Configure Monitoring dialog only): A button that updates the test configured in the dialog. The button becomes active when you make a change to the current settings.

#### Agents List

The **Agents** list shows you the available agents for this test. If the test allows selection of both Private and Global agents, the agents will be grouped into four sections that each start with a heading row (Private Network Agents, Private App Agents, Global Network Agents, and Global App Agents). The list includes the following columns:

* **Select All** (checkbox): A checkbox (in the heading row) for toggling the selection state of all agents in the list between checked and unchecked.
* **Select** (checkbox): Check/uncheck an agent to include/exclude it from this test.
* **Agent**: The name of the agent.
* **Location**: The name of the site or cloud region in which the agent is deployed.
* **IP Version**: The IP version of the agent (v4 + v6 or v4 only).

> **Note**: The projected number of test credits consumed by the agent(s) is immediately visible on the **Expected monthly credit usage** graph (in blue) after checking the agent boxes.

### Offload Profile

The **Offload Profile** pane provides a breakdown of the types of connections over which your subscribers are receiving content from the CDN. The types are determined by the connectivity type of the interfaces through which traffic from the CDN enters your network (see **Understanding Connectivity Types**).  
  
The breakdown shows the percentage of traffic from the CDN that is attributed to various connectivity types. The percentages are displayed numerically in panels across the top of the pane (below the help text), and graphically in the stacked area chart below those panels. The following connectivity types are represented in the breakdown: Caches, Free Private Peering, Transit, IX, Customer, and Other.  
  
The **View in Data Explorer** button at the bottom right of the chart takes you to Data Explorer, where the **Query** sidebar will be set to show a configurable version of the same visualization.

### CDN Transport Breakdown

The **CDN Transport Breakdown** pane shows the traffic transported to your network by this CDN, broken out by connectivity type, provider, and site. The pane includes two main elements, the **CDN Transport Chart** and the **CDN Transport Table**.

#### CDN Transport Chart

The CDN Transport chart is a Sankey diagram in which each path represents a unique combination of the following dimensions:

* **Connectivity type**: The type of connection over which traffic from this CDN enters your network (see **Understanding Connectivity Types**).
* **Provider**: The provider via which traffic from this CDN enters your network (see **Provider Classification**).
* **Site**: The site where the provider receives traffic from the CDN that goes to your network (see **About Sites**).
* **CDN**: The name of this CDN.
* **Source ASN**: The origin ASN associated with the source IP of the traffic.
* **View in Data Explorer**: A link to Data Explorer, where the Query sidebar will be set to show the traffic displayed in the CDN Transport chart.

#### CDN Transport Table

The CDN Transport table lists the unique paths by which traffic from this CDN entered your network during the currently specified time range (see **CDN Analytics Title Pane**). The table can be sorted by clicking the heading of the desired sort column; clicking the same heading again will change the sort direction (ascending vs. descending).  
  
 The table includes the following columns:

* **Traffic sparkline**: A chart showing the rate of traffic matching this CDN over the selected time range.
* **Source Connectivity Type**: The type of connection over which traffic from this CDN enters your network (see **Understanding Connectivity Types**). Click the type to go to a Network Explorer page detailing that individual connectivity type (see **Core Details Views**).
* **Source Provider**: The provider via which traffic from this CDN enters your network (see **Provider Classification**). Click the provider to go to that provider’s Details page (see **Core Details Pages**).
* **Site**: The site where the provider receives traffic from the CDN that goes to your network (see **About Sites**). Click the site to go to that site’s Details page (see **Core Details Pages**).
* **Average Gbits/s**: The average traffic rate from the CDN to your network over this path.
* **Average Unique Dst IPs**: The average number of unique destination IPs over this path.
* **Average Mbits/s per dst ip**: The average traffic rate over this path divided by the average number of unique destination IPs.

### Top OTT Services

The **Top OTT Services** pane shows the individual “over-the-top” (OTT) content services via which traffic is coming to your network from the CDN. The pane includes two main elements, the **Top OTT Services Chart** and the **Top OTT Services Table**.

#### Top OTT Services Chart

This stacked-area chart shows the distribution, over the page's currently specified time range, of the OTT services used by traffic coming to your network from this CDN. Hovering on the chart at a given point in the time range pops up a breakdown, by service, of the traffic at that time.  
  
The **View in Data Explorer** button at the bottom right of the chart takes you to Data Explorer, where the **Query** sidebar will be set to show the same traffic that is displayed in the **Top OTT Services** chart.

#### Top OTT Services Table

This table lists the OTT services via which traffic comes to your network from this CDN. The list can be filtered with the **OTT Classification Types** field to show only services of one or more classification types.  
  
 The OTT Services table includes the following columns:

* **OTT Service**: An individual OTT content service whose hostname is looked up via DNS. Click the service name to go to an OTT Service Tracking page detailing that individual OTT service (see **Service Details Pages**).
* **OTT Service Category**: The nature of the content provided by an OTT content service (for possible values, see **Application Context and Security**). Click the service category to go to an OTT Service Tracking page detailing that individual OTT service category (see **Service Category Details**).
* **OTT Service Provider**: An entity that offers an OTT content service (e.g., Google is the provider for Google Drive, Gmail, Google Maps, etc.). Click the provider’s name to go to an OTT Service Trackingpage detailing that individual OTT service provider (see **Service Provider Details**).
* **Classification**: The extent to which Kentik has been able to classify the OTT service, service category, and service provider (see **OTT Classification Values**).
* **p95th Mbits/s**: The 95th percentile traffic rate from the CDN via this service during the time slice with the highest traffic rate.
* **Max Unique Dst IPs**: The number of unique destination IPs via this service during the time slice with the highest count of destination IPs.
* **Max Mbits/s per Dst IP**: The highest bitrate per destination IP for of all time slices in the time range.

> **Note:** The time slices (aggregation steps) used to calculate the above maximum values vary depending on the **Time Range** setting (see **CDN Analytics Title Pane**), with a longer range resulting in longer slices.

#### OTT Classification Values

The **Classification** column shows the extent to which Kentik has been able to classify the OTT Service Name, OTT Service Category (also referred to as service type), and OTT Provider Name, which are represented in KDE as **Application Context and Security** dimensions. The state of this classification for a given service is indicated with the following values:

* **Full**: The service, service category, and service provider have been identified.
* **Provider-Only**: The service provider has been identified, but service and service category have not. This indicates that the provider's services are all served from the same hosts (i.e. where multiple simultaneous DNS queries towards different services return the same IP address).
* **Pending**: Traffic patterns from this CDN to your network (i.e. a hostname generating high traffic rates) indicate a service that hasn't yet been matched to an existing Kentik-identified OTT service.
* **Unknown**: The service, service category, and service provider have not been identified.

#### OTT Classification Types

The **Classification Types** field shows a lozenge for each OTT service classification types (see **OTT Classification Values**) that's been selected for inclusion in the list:

* When no types are selected, services of all classification types appear in the list.
* To select a type, click in the field and choose the type from the resulting popup.
* To remove a type, click the **X** in the type's lozenge.

### Performance Analysis

The **Performance Analysis** pane shows a line chart plotting the bitrate per subscriber over the page's current time range. This chart helps to compare performance for this CDN by connectivity type, connectivity provider, and site so that you can quickly spot any deficiencies.  
  
The left axis of this chart is “bits per Subscriber,” which corresponds to the average bits per second per unique destination IP. This is usually considered a sufficient proxy for subscriber bitrate, since each unique destination IP can be understood as the shared IP for a household of one or more subscribers.  
  
This pane includes the following controls and displays:

* **Aggregate**: A dropdown from which you select the type of aggregation (Average, 95th Percentile, or Maximum) used to calculate the bitrate for each time slice in the currently selected time range.
* **Show Performance By**: A dropdown from which you select the analysis that will be charted. The resulting chart(s) may be single-level or dual-level (see Charts below).
* **Charts**: One or more line charts (depending on the **Show Performance By**setting) showing the bitrate per subscriber over the selected time range:

  + **Single-level chart**: If **Show Performance By** is set to a single-level value (Site, Provider, or Connectivity Type) there will be a single chart with a plot for each instance of the selected chart subject (e.g., if the setting is Site then a plot for each site).
  + **Dual-level charts**: If **Show Performance By** is set to a dual-level value (e.g., "Site by Provider") then there will be a series of charts, one for each top-level entity (e.g., site), on which there will be a plot for each lower-level entity (e.g., provider).
* **View in Data Explorer**: A link, at the bottom right of each chart, to Data Explorer, where the Query sidebar will be set to show the traffic displayed in the chart.
