# OTT Service Details & Analytics

Source: https://kb.kentik.com/docs/ott-service-details-analytics

---

This article explains how to drill down into specific **OTT Service Tracking** categories, providers, and services to analyze detailed traffic paths, subscribership, and capacity.

## Service Category Details

Access the details page for a service category from the **OTT Service Tracking** page by clicking a category in the **Top Categories List** or the **OTT Service Category** column in the **Top Services Table**. This page provides detailed traffic information for that category coming to your network.

![Graph showing total traffic for observability services, highlighting Kentik's performance metrics.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/OTT-category-details.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A33Z&se=2026-05-12T09%3A36%3A33Z&sr=c&sp=r&sig=AxjM9OdhWzOKqkqS3rxT0yjanNXT73AOqupnH7Hk6P8%3D)

*The details page for a category of OTT service providers.*

The service category details page includes these UI elements:

* **Share**: Opens the **Share** dialog to share the current view (see **Sharing via the Share Dialog**).
* **Actions**: **Export** a PDF report of visualizations and tables. A notification appears when the PDF is ready to download.
* **Favorite**: Add the page to **Favorites** in portal search (see **Portal Search Tabs**).
* **Category Name**: Displays the current category.
* **Time Range**: Select the duration for traffic data display (last day, last week, last 2 weeks, last 30 days).
* **Services by Total Traffic**: Shows top services' traffic info (see **Top Category Services**).
* **Services by Max bits/s**: Lists services by traffic bitrate (see **Category Services by Bitrate**).

#### Top Category Services

The **Top Services by Total Traffic** pane includes the following elements:

* **Total traffic**: Displays the combined max bitrate of all OTT services in this category for selected time range.
* **Traffic by service**: A bar chart illustrating the max bitrate for each service in this category.
* **View in Data Explorer**: A button linking to **Data Explorer** with the same traffic data pre-set.

> **Note**: If the **Intelligent Classification** toggle is enabled, the total traffic and max bitrate values displayed here will dynamically recalculate to include statistically inferred traffic alongside deterministic data.

#### Category Services by Bitrate

The **Services by Max bits/s** table lists services in this category transporting traffic to your network, with these columns:

* **Sparkline**: Visual traffic volume over the selected time range.
* **OTT Service**: The service name, linked to its details page (see **Service Details Pages**).
* **OTT Provider**: The provider name, linked to its details page (see **Service Provider Details**).
* **Max**: Maximum traffic bitrate for the service.
* **Traffic by Connectivity Type**: A breakdown of the max bitrate by connectivity type (e.g., Embedded, IX, Free Peering. See **Understanding Connectivity Types**).

## Service Provider Details

Access a service provider's details page from the OTT Service Tracking landing page by clicking a provider in the **Top Categories List**or the **OTT Service Provider** column in the **Top Services Table**. This page offers detailed traffic information from that provider to your network.

![Traffic analytics for eBay showing connectivity types and maximum bits per second.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/OTT-provider-details.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A33Z&se=2026-05-12T09%3A36%3A33Z&sr=c&sp=r&sig=AxjM9OdhWzOKqkqS3rxT0yjanNXT73AOqupnH7Hk6P8%3D)

*The details page for an OTT service provider.*

The service provider details page includes these UI elements:

* **Share**: Opens the **Share** dialog to share the current view (see **Sharing via the Share Dialog**).
* **Actions**: **Export** the page as a PDF report with visualizations and tables. A notification appears when the PDF is ready to download.
* **Favorite**: Star icon to add the page to **Favorites** in portal search (see **Portal Search Tabs**).
* **Provider Name**: Displays the provider's name.
* **Time Range**: A selector for traffic data duration (last day, last week, last 2 weeks, last 30 days).
* **Top Services by Total Traffic**:

  + Total traffic: Shows combined max bitrate for the provider's OTT services.
  + Traffic by service: Bar chart of max bitrate per service.
* **Traffic by Connectivity Type**: A bar chart with segments for each connectivity type (see **Understanding Connectivity Types**), showing percentage and volume.
* **Traffic analytics**: A Sankey diagram of traffic paths. When the **Intelligent Classification** toggle is enabled, a **Classification** column appears on the far left to visually separate Fully Classified traffic from Intelligently Classified traffic. Includes a button to view in **Data Explorer** with the same traffic data pre-set.
* **Provider Services by Max bits/s**: Lists the services from this provider that transport traffic to your network, sorted by maximum bitrate (see **Provider Services by Bitrate**).

> **Note**: If the **Intelligent Classification** toggle is enabled, the total traffic and max bitrate values displayed here will dynamically recalculate to include statistically inferred traffic alongside deterministic data.

#### Provider Services by Bitrate

The **Provider Services by Max bits/s** table lists the services from this provider that are transporting traffic to your network. The table includes the following columns:

* **Sparkline**: Visual traffic volume over the selected time range.
* **OTT Service**: The service name, linked to its details page (see **Service Details Pages**).
* **OTT Service Category**: The category name, linked to its details page (see **Service Category Details**).
* **Max**: Maximum traffic bitrate for the service.
* **Traffic by Connectivity Type**: Individual columns for each connectivity type, e.g., Embedded, IX, Free Peering, displaying the maximum traffic bitrate for each (see **Understanding Connectivity Types**).

## Service Details Pages

To access the details page for an individual OTT service from the **OTT Service Tracking** page, click the service name in any of these locations:

* OTT Service column of the **Top Services Table**
* Under an expanded provider in the **Top Providers List**
* Under an expanded category in the **Top Categories List**

### Service Details UI

The Details page for a service provides information for traffic entering your network from that service. Key UI elements include:

* **Breadcrumbs**: Navigate your location within the Kentik portal.
* **Share**: Opens the **Share** dialog (see **Sharing via the Share Dialog**).
* **Actions**: **Export** a PDF report of the page's visualizations and tables. A notification appears when the PDF is ready to download.
* **Favorite**: Star icon to add the page to Favorites in portal search (see **Portal Search Tabs**).
* **Time Range**: Select the duration for traffic data display (last day, last week, last 2 weeks, last 30 days).
* **Service Details tabs**: Four tabs with service information (see **Service Detail Tabs**).

#### Service Detail Tabs

Service information is organized into four tabs:

* **Overview**: Basic traffic info and service comparison context.
* **Connectivity**: How traffic reaches your network.
* **Subscribers**: Unique destination IPs consuming traffic.
* **Capacity**: Interface utilization and health.

### Service Details Overview

The **Overview** tab of the details page for individual services includes:

* **Category**: Indicates service type (e.g., video, gaming, social).
* **Provider**: Shows the OTT service provider's name.
* **Logo**: Displays the service or provider's logo.
* **Total Inbound + Embedded**: The max bitrate (SNMP) for the selected time range.
* **Offload**: The percentage of traffic from embedded caches at peak volume.
* **Traffic by Connectivity Type**: A bar chart showing traffic breakdown by connectivity type at peak volume (see **Understanding Connectivity Types**).
* **Rankings**: Compares peak traffic volume:

  + **Service by volume**: Rank among all OTT services.
  + **Service in category**: Rank within the same category.
  + **Service delivered by provider**: Rank among services from the same provider.

![Roblox service overview showing traffic statistics and rankings in gaming category.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/OTT-service-details.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A33Z&se=2026-05-12T09%3A36%3A33Z&sr=c&sp=r&sig=AxjM9OdhWzOKqkqS3rxT0yjanNXT73AOqupnH7Hk6P8%3D)

*The Overview tab of the details page for an individual service.*

> **Note**: If the **Intelligent Classification** toggle is enabled, the total traffic and max bitrate values displayed here will dynamically recalculate to include statistically inferred traffic alongside deterministic data.

#### Service Connectivity Details

The **Connectivity** tab of a service details page includes:

* **Traffic by Connectivity Type**: A bar chart showing connectivity types (see**Understanding Connectivity Types**), each labeled with type, traffic percentage, and volume at peak traffic time.
* **Breakdown**: A tabbed chart displaying traffic from various perspectives (see **Service Traffic Breakdown**).
* **Digital Supply Chain**: A Sankey diagram illustrating the traffic path from source AS to your network. If the **Intelligent Classification** toggle is enabled, the diagram includes a **Classification** column on the far left.

#### Service Traffic Breakdown

The tabbed **Breakdown** chart displays service traffic from these perspectives:

* **Source CDN**: The commercial CDN name from the source IP (see **About Provider Classification**).
* **Connectivity**: The type of traffic connectivity (see **Connectivity Type Attribute**).
* **Providers**: The internet connection provider (see **About Provider Classification**).
* **Sites**: Network sites with devices/interfaces handling the traffic (see **Sites**).
* Use the **View in Data Explorer** button to explore the same traffic in **Data Explorer**.

![Graph showing CDN performance metrics over time with multiple data sources and total bandwidth.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(268).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A33Z&se=2026-05-12T09%3A36%3A33Z&sr=c&sp=r&sig=AxjM9OdhWzOKqkqS3rxT0yjanNXT73AOqupnH7Hk6P8%3D)

*The Connectivity tab includes the tabbed Breakdown chart.*

### Service Subscribers Details

The **Subscribers** tab on the service details page includes the **Overall Subscribership** and **Performance Analysis** panes. When measuring subscribership for an OTT service, Kentik's True Origin engine uses the Max Unique Count of Destination IPs as a proxy for subscriber IPs (see **Metrics from All Devices**).

![Overview of Xbox Live subscribers and performance analysis over a week period.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(269).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A33Z&se=2026-05-12T09%3A36%3A33Z&sr=c&sp=r&sig=AxjM9OdhWzOKqkqS3rxT0yjanNXT73AOqupnH7Hk6P8%3D)

*The Subscribers tab shows a count of subscribers for the service and the bitrate for each.*

### Overall Subscribership

This pane provides details on unique destination IPs consuming traffic for this service, featuring:

* **Compare with Category leader**: A checkbox to include (default) or exclude traffic from the category's top service. Not shown if this service is the leader.
* **Show total Category traffic**: A checkbox used to either include (default) or exclude a plot of total category traffic.
* **Daily Subscribers**: The average daily unique destination IPs receiving content, with percent change.
* **Avg Performance Per Subscriber**: The average bitrate divided by unique destination IPs over the selected time range.
* **Subscribers chart**: Displays unique destination IP counts over time. May include category leader and total category counts based on checkbox settings. Hover for specific counts.
* **Ranking list**: The service's rank by subscriber counts within its category. Click to view service details.

#### Performance Analysis

This pane displays bitrate breakdowns per subscriber for an OTT service over a selected time range, featuring:

* **Breakdown by**: Choose analysis type (single-level or dual-level) for charting.
* **Swap levels**: Swap chart levels (e.g., "Site then Provider" to "Provider then Site") for dual-level charts.
* **Aggregate on**: Select aggregation method (Average, 95th Percentile, or Maximum) for bitrate calculation.
* **Charts**:

  + **Single-level**: One chart per selected subject (e.g., Site, Provider).
  + **Dual-level**: Individual charts for top-level entities with plots for lower-level entities.
* **View in Data Explorer**: Opens **Data Explorer** with the same traffic settings.

  > **Notes:**
  >
  > + Bitrates are calculated per subscriber WAN IP, counting multiple subscribers in the same household as one.
  > + The Avg/95p/Max bitrates in charts reflect those in Data Explorer for each unique destination IP, excluding IPs for embedded caches, with traffic filtered to Network Boundary = External OR Connectivity Type = Embedded Cache.

### Service Capacity Details

The **Capacity** tab on service details page includes:

* **Capacity Treemap**: Displays a treemap where each area represents an interface used by the service.
* **Per-Device Capacity Analysis**: Shows panes for each device, detailing interface utilization for the service.

#### Capacity Treemap

The treemap illustrates interfaces through which users consume service-related traffic. For services using multiple interfaces, each block's size indicates its traffic volume share (max bitrate for the selected time range). Block colors reflect interface health based on thresholds set in the **OTT Thresholds Tab**.

![Treemap showing bandwidth usage for Amazon Prime Video over the last week.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(270).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A33Z&se=2026-05-12T09%3A36%3A33Z&sr=c&sp=r&sig=AxjM9OdhWzOKqkqS3rxT0yjanNXT73AOqupnH7Hk6P8%3D)

*The treemap shows the relative OTT traffic volume on the interfaces used for a given service.*

#### Treemap Popup

Hovering over a treemap block reveals a popup with:

* **Device name**: The name of the device. Click to view Core Details (see **Core Details Pages**).
* **Site**: The network site of the device (see **Sites**).
* **Bitrate**: The max traffic bitrate for the selected time range.
* **Interface name**: The name of the interface. Click to view Core Details (see **Core Details Pages**).
* **Interface description**: SNMP-based description, labeled "Inherited from physical" if derived.
* **Connectivity type**: The traffic connectivity type (see **Connectivity Type Attribute**).
* **Provider**: The internet connection provider (see **About Provider Classification**).
* **Utilization chart**: A peak utilization meter as a percent of capacity.

### Per-Device Capacity Analysis

The **Per-Device Capacity Analysis** pane includes a section for each device handling traffic for the service, featuring:

* **Device name**: The name of the device. Click to expand and reveal the device's interfaces.
* **Bar chart**: Displays interfaces grouped by health status (Healthy = green, Warning = orange, Critical = red), based on overall utilization. Segment length reflects the traffic volume for the service.
* **Traffic volume**: Shows the peak bitrate for the service during the selected time range.
* **Interfaces**: An expandable list of interfaces carrying traffic for the service (see **Service Interfaces List**).

![Analysis of Amazon Prime Video traffic across different interfaces and their capacities.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(274).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A33Z&se=2026-05-12T09%3A36%3A33Z&sr=c&sp=r&sig=AxjM9OdhWzOKqkqS3rxT0yjanNXT73AOqupnH7Hk6P8%3D)

*OTT traffic volume over the interfaces of a given device.*

### Service Interfaces List

The table lists interfaces carryingservice traffic, with each expandable row containing:

* **Expand/Collapse**: Toggle to show/hide **Interface Traffic Charts**.
* **Bitrate**: The max bitrate for the service for the selected time range.
* **Interface name**: The name of the interface.
* **Interface description**: The SNMP-based description, or "Inherited from physical" if derived.
* **Connectivity type**: The type of the service traffic connectivity (see **Understanding Connectivity Types**).
* **Provider**: The internet connection provider (see **About Provider Classification**).
* **Utilization chart**: A meter showing peak utilization, traffic volume, and capacity.

### Interface Traffic Charts

In the **Service Interfaces List**, use the expand/collapse toggle to show/hide line charts for each interface. These charts display:

* **Interface Traffic & Capacity**: Total traffic on the interface for the selected time range.
* **Subscribers**: The number of unique destination IPs for the service on this interface.
* **AVG Mbps / unique subscriber**: The bitrate of service-related traffic divided by the number of unique destination IPs during the time range.

![Graphs showing GitHub traffic, unique subscribers, and average Mbps over time.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(275).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A33Z&se=2026-05-12T09%3A36%3A33Z&sr=c&sp=r&sig=AxjM9OdhWzOKqkqS3rxT0yjanNXT73AOqupnH7Hk6P8%3D)

*Traffic over an interface used for an POTT service is detailed in three charts.*

Now that you've explored the detailed dashboard views, learn how to build custom reports and filter by exact classification methods in **Querying OTT Classification Data**.
