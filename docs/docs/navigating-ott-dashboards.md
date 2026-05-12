# Navigating OTT Dashboards

Source: https://kb.kentik.com/docs/navigating-ott-dashboards

---

This article explores the primary views within the **OTT Service Tracking** workflow, guiding you through the top-level dashboards, tables, and controls used to analyze your network's OTT traffic.

## OTT Service Tracking Page

The OTT Service Tracking page includes the following main sections and controls:

* **Share**: Opens the **Share** dialog (see **Sharing via the Share Dialog**) for sharing options.
* **Actions**: **Export** visual reports (PDF) of the page’s content. A notification appears when the PDF is ready to download.
* **Favorite**: Add the page to Favorites in the portal search (see **Portal Search Tabs**).
* **Configure**: Access the **OTT Configuration Page** for **Thresholds** (see **OTT Thresholds Tab**) and **Advanced Filtering** (see **OTT Advanced Filtering Tab**).
* **Time Range**: Select the duration for traffic information display (Last Day, Last Week, Last 2 Weeks, Last 30 Days).
* **Intelligent Classification**: When switched On, shows services identified via algorithmic matching (see **Intelligent Classification**).
* **Display Area Tabs**:

  + **Overview** (default): Shows OTT service traffic categories, providers, and services.
  + **Capacity**: Displays interface utilization by OTT service category.
  + **Pending**: Displays traffic from high-volume hostnames that do not yet match a known OTT service pattern.

### Intelligent Classification

The **Intelligent Classification** (also known as IQC) toggle allows you to control how OTT traffic is displayed across the OTT Service Tracking pages, including the Overview, Capacity, and Pending tabs as well as the Provider Detail, and Service Detail views. The toggle is located in the upper-right corner of these pages.

**Disabled** (Default): The dashboard displays only Fully Classified traffic. This represents traffic deterministically identified via DNS mapping.

**Enabled**: The dashboard view updates to include both Fully Classified traffic and Intelligently Classified (statistically inferred) traffic. When enabled, a persistent UI indicator will remain visible to provide transparency that statistical data is being included in the current view.

> **Note**: The **Capacity** tab on the Service Detail page relies on deterministic data and is incompatible with statistical classification. The Intelligent Classification toggle will appear disabled (grayed out) when viewing the Capacity tab.

## OTT Overview Tab

The **Overview** tab of the OTT Service Tracking page includes:

* **Total OTT Traffic**(pane): Provides a high-level view of OTT services on your network.
* **Top Categories**(list): Ranks OTT service categories by traffic volume over the selected time range.
* **Top Providers**(list): Ranks OTT service providers by traffic volume over the selected time range.
* **Top Services**(table): Lists services by traffic volume associated with incoming network traffic.

### Total OTT Traffic

The **Total OTT Traffic** pane provides a high-level view of OTT services on your network, including:

* **OTT engine**: Details on the current state of the OTT engine, including the number of services, providers, categories, and the last update time.
* **OTT bitrate**: The maximum bitrate of all OTT traffic types during the selected time range.
* **Bar chart**: A visual breakdown of traffic classification (see **OTT Classification Values**), showing percentage and bitrate for each segment.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(251).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A53Z&se=2026-05-12T09%3A35%3A53Z&sr=c&sp=r&sig=%2BiJY8i54xrR0WceB4VYp77P16m%2BpL04FJTyxRuVXoWI%3D)

*Total OTT Traffic shows info about overall bitrate and OTT Classification.*

### Top Categories List

The **Top Categories** list displays your organization's top 8 OTT services categories by traffic bitrate (max reported by SNMP). Each row includes:

* **Ranking**: The position based on traffic bitrate.
* **Expand/Collapse**: Toggle to view the top **X** services in this category, including logo, name, and max SNMP-reported bitrate. Click a name for service details.
* **Category**: The name of the service category (e.g., video, gaming, social). Click to view the view the **Service Category Details**.
* **Traffic Volume**: The max traffic volume for the category in the selected time range.
* **Traffic Trend**: The percent change in traffic volume for the category during the time range.
* **Traffic Chart**: A sparkline showing traffic volume fluctuations for the category during the time

### Top Providers List

The **Top Providers** list displays your organization’s top 8 OTT service providers, ranked by traffic bitrate. Each row includes:

* **Ranking**: The position based on traffic bitrate.
* **Expand/Collapse**: Toggle to view the top X services from the provider, including logo, name, and max SNMP-reported bitrate. Click a name for service details.
* **Service Provider**: The Service Provider name (e.g., Microsoft, Netflix). Click a name for provider details (see **Service Provider Details**).
* **Traffic Volume**: The traffic volume transported by the provider.
* **Traffic Trend**: The percent change in traffic volume over the time range.
* **Traffic Chart**: A sparkline showing traffic fluctuations over the time range.
* **Logo**: The provider's logo.

### Top Services Table

The **Top Services** table lists, in order of max SNMP-reported bitrate, the services associated with the traffic coming into your network. The table may be filtered using the **Classification Types** field.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(252).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A53Z&se=2026-05-12T09%3A35%3A53Z&sr=c&sp=r&sig=%2BiJY8i54xrR0WceB4VYp77P16m%2BpL04FJTyxRuVXoWI%3D)

*A list, ordered by max bitrate, of the OTT services on your network.*

The table includes the following columns:

* **Traffic sparkline**: The visual traffic volume over time, highlighting peak traffic events.
* **OTT Service**: The service name, clickable for details (see **Service Details Pages**). May be unknown based on classification.
* **OTT Service Category**: The service category (i.e., service type), clickable for details (see **Service Category Details**). Displays “unknown” if classified as Provider-only or Pending.
* **OTT Service Provider**: The provider of the service. If the service is classified as Pending, it displays as “unknown”. Click the provider name to view details (see **Service Provider Details**).
* **Classification**: The OTT classification by Kentik (see **OTT Classification Values**).
* **Max Gbits/s**: The highest traffic volume from the CDN during the interval.
* **Max Unique Dst IPs**: The highest count of unique destination IPs during the interval.
* **Max Gbits/s per dst ip**: The highest bitrate per IP during the interval.

> **Notes:**
>
> * Aggregation steps (intervals) vary with the selected time range (see **OTT Service Tracking Page**). Longer ranges mean longer intervals.
> * Click **View in Data Explorer** to open **Data Explorer** with pre-set controls.
> * When the Intelligent Classification toggle is enabled, the traffic volumes and rankings in the **Overview** tab will dynamically recalculate to include statistically inferred traffic alongside fully classified deterministic data.

#### Classification Types

The **Classification Types** field displays a lozenge for each selected OTT service classification types (see **OTT Classification Values**):

* If no types are selected, all classification types are listed.
* Click the field to select the type from the popup.
* Remove a type by clicking the X in its lozenge.

## OTT Capacity Tab

The **Capacity** tab on the OTT Service Tracking page displays the health (utilization vs. capacity) of interfaces carrying OTT-classified traffic (see **OTT Classification Values**). Key elements include:

* **Hide Healthy Services**: A switch at the top right:

  + **On** (default): Shows treemaps only for services with over 25% traffic from over-utilized interfaces.
  + **Off**: Displays treemaps for all OTT services.
* **Service Category Panes**: Each pane corresponds to an OTT service category (see **OTT Service Category Panes**).

  ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(255).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A53Z&se=2026-05-12T09%3A35%3A53Z&sr=c&sp=r&sig=%2BiJY8i54xrR0WceB4VYp77P16m%2BpL04FJTyxRuVXoWI%3D)

  *A list showing the health of interfaces used for OTT traffic in all OTT service categories.*

### OTT Service Category Panes

Each pane includes the following UI elements:

* **Category**: The category of OTT service detailed in this pane (also referred to as service type).

  + Click the category name to go to a Details page for that category (see **Service Category Details**).
* **Health Indicator**: One lozenge per health status in the interfaces carrying traffic for this service category. Each lozenge shows the number of interfaces with the corresponding status: green (Healthy), orange (Warning), or red (Critical).
* **Service Name**: The individual service(s) listed for each category. By default, only Warning or Critical services are shown when **Hide healthy services** is On.

  + Click the service name to go to the Details page for that service (see **Service Details Pages**).
* **Treemap**: One treemap diagram for each service in the category. Each area represents the interfaces used by that service. If a service uses multiple interfaces, the area size represents the relative traffic volume (max bitrate during the selected time range) over those interfaces. The color of each area corresponds to the interface’s health as defined by the **OTT Thresholds Tab**.

## OTT Pending Tab

The **Pending** tab on the OTT Service Tracking page isolates traffic categorized as "Pending Classification". This view displays traffic from high-volume hostnames that do not yet match a known OTT service pattern, allowing you to monitor unmapped traffic that the True Origin engine is actively evaluating for future classification.

![Table displaying pending OTT services with their maximum and last data point values.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/OTT-pending-tab(1).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A53Z&se=2026-05-12T09%3A35%3A53Z&sr=c&sp=r&sig=%2BiJY8i54xrR0WceB4VYp77P16m%2BpL04FJTyxRuVXoWI%3D)

Because the True Origin engine has not yet mapped these specific hostnames to a known service or provider, the data presented in this tab behaves slightly differently than the fully classified views:

* **Hostnames Displayed:** Instead of known OTT Service names, the table lists the raw, high-volume hostnames (e.g., `ns1.nexellent.net`).
* **Unknown Dimensions:** Because the traffic is still pending classification, the **OTT Service** and **OTT Service Category** dimensions will display as "unknown".
* **Traffic Metrics:** Even without a known service mapping, you can still monitor the exact traffic volume for these hostnames, including the Max Mbits/s and Max Unique Destination IPs.

Now that you understand the high-level dashboard views, let's drill down into the specifics in **OTT Service Details & Analytics**.
