# Discover Peers

Source: https://kb.kentik.com/docs/discover-peers

---

This article covers the **Discover Peers** workflow in the Kentik portal.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(333).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A27Z&se=2026-05-12T09%3A42%3A27Z&sr=c&sp=r&sig=QLEsfiVm7NWLKccRizgegx19XKMXaL%2BXQebOuKMJ%2FAg%3D)

*Discover Peers shows the ASes via which traffic gets to and from your network.*

## About Discover Peers

The following topics explain peering and the benefits of Kentik's Discover Peers workflow.

### Peering Overview

“Peering” refers to connecting one network with another in order to exchange data. The arrangements governing this interconnection typically fall into one of the following models:

* **Paid peering**: A paid arrangement in which one network pays another for Internet transit or access to a subset of specific routes. Internet transit (a.k.a. DIA, for Direct Internet Access) is how many networks access the Internet and connect with other entities. It offers little control over how traffic is sent or delivered once it has leaves your network, and it can be costly depending on the location and provider of the service.
* **Settlement-free peering**: Networks interconnect without charge, thus allowing access to specific destinations on each other’s networks. Although no money is exchanged, there are still associated costs such as equipment, colocation, circuit costs, and sometimes port fees (if the peering happens over public peering facilities).

Regardless of the peering model, the costs associated with interconnection are an important consideration in choosing peering arrangements. Another critical factor that can be influenced by peering is Internet performance. Every intermediate device or network transited by data between its source and destination introduces latency and potential points of failure. Thus, direct peering between content providers and consumers is considered to be the gold standard for network performance at Internet scale.

### Discover Peers Use Cases

Discover Peers is intended to enable fast discovery of peering opportunities for two main classes of networks:

* **Access ISPs and traditional enterprises**: The operators of these networks are primarily concerned with transporting the content demanded by their users while controlling Internet costs and ensuring the best possible performance.
* **Digital content providers**: The operators of these networks are primarily concerned with sending traffic via routes that are both cost-effective and highly direct.

Operators at both of the above classes of networks will be able to use Discover Peers for the following:

* Discover which networks to peer with and where you can interconnect.
* Predict impacts on your existing transit circuits.
* Schedule regular reports and receive timely insights when action is needed.

The following table shows how Discover Peers use cases apply to various roles within a network organization.

| Role | Scenarios |
| --- | --- |
| Network Strategist | * Find opportunities to save money on transit costs and/or facilitate better performance for customers and applications. * See the physical locations of potential peers to aid estimation of feasibility and cost. * See the potential impact of peering on existing transit and peering links. * See the effective overall cost of peering. * Use baselines of destination or source ASNs to assess the value of interconnection based on historical traffic patterns. |
| Network Architect | * See current and historical utilization trends on peering and transit interfaces, enabling informed decisions on augmenting existing services or building new ones. |
| Network Engineer | * Use information on peers when automating the provisioning of new peering configurations on routers. |

## Discover Peers Page

The Discover Peers page is covered in the following topics.

### Discover Peers Layout

The **Discover Peers** page includes the following main UI elements:

* **Related Insights**: Click the button at upper right to drop down a list of Kentik Insights related to peering.
* **Configure button**: Click the button to go to the **Configure Discover Peers**.
* **Filter Paths pane**: Contains a set of controls for limiting the scope of the Discover Peers analysis (see **Filter Paths Pane**).
* **Peering Chart pane**: Contains the peering chart as well as associated controls (see **Peering Chart Pane**).
* **Potential Peers pane**: Contains the Potential Peers table as well as associated controls (see **Potential Peers Pane**).

### Filter Paths Pane

The **Filter Paths**pane contains a set of controls that enable you to limit the paths that are included in the analysis whose results are displayed in the **Peering Chart** and the **Potential Peers** table. The following table shows the filters in the Filter Paths pane and how their effect varies depending on the **Direction** setting (see **Peering Chart Pane**):

| Filter | Inbound | Outbound |
| --- | --- | --- |
| Sites | Show only paths that enter your network at the specified sites. | Show only paths that exit your network at the specified sites. |
| Countries | Show only paths that originate in the specified countries. | Show only paths that terminate in the specified countries. |
| Site Countries | Show only paths that enter your network at sites within the specified country. | Show only paths that exit your network from sites within the specified country. |
| Peering Policy | Show only paths that originate from ASNs that conform to the selected PeeringDB peering policies. | Show only paths that terminate or pass through ASNs that conform to the selected PeeringDB peering policies. |
| Peering Traffic Ratio | Show only paths that originate from ASNs that conform to the selected PeeringDB Peering Traffic Ratios. | Show only paths that terminate or pass through ASNs that conform to the selected PeeringDB Peering Traffic Ratios. |
| Show Traffic Through | N.A. | Determines the columns shown in the Peering Chart and the attribution of traffic to ASes in the Potential Peers table (see **Show Traffic Through**). |

#### Show Traffic Through

If the **Direction** selector is set to Outbound then the **Show Traffic Through** setting affects both the **Peering Chart** and the **Potential Peers** table. The table below shows how the setting determines which dimensions are represented as Sankey columns in the chart.

| Show Traffic Through | Site | Dest Connectivity Type | Dest Next Hop ASN | Dest 2nd Hop ASN | Dest 3rd Hop ASN | Dest ASN |
| --- | --- | --- | --- | --- | --- | --- |
| Destination ASN | Y | Y | Y | N | N | Y |
| 2nd Hop ASN | Y | Y | Y | Y | N | Y |
| 3rd Hop ASN | Y | Y | Y | N | Y | Y |
| 2nd Hop ASN and 3rd Hop ASN | Y | Y | Y | Y | Y | Y |

In the **Potential Peers** table, if the **Direction** selector is set to Outbound then the traffic attributed to a given AS varies depending on the **Show Traffic Through** setting:

* **Destination ASN** (default): Include only traffic for which this AS is the destination.
* **2nd Hop ASN**: Include traffic for which this AS is either the destination or the 2nd Hop AS.
* **3rd Hop ASN**: Include traffic for which this AS is either the destination or the 3rd Hop AS.
* **2nd Hop ASN and 3rd Hop ASN**: Include traffic for which this AS is either the destination, the 2nd Hop AS, or the 3rd Hop AS.

### Peering Chart Pane

The **Peering Chart** pane displays a Sankey chart showing the paths that haven't been filtered out by the settings in the **Filter Paths Pane**. The pane includes the following controls:

* **Direction selector**: Determines the perspective from which the chart and table are populated:

  + Inbound: Shows paths carrying traffic from external source ASes to your network's sites (see **About Sites**). Designed for networks that receive a lot of traffic from the Internet (e.g. Access ISPs).
  + Outbound: Shows paths carrying traffic from your network's sites to external destination ASes. Designed for networks that send a lot of traffic to the Internet (e.g. content providers).
* **Peering Sankey chart**: A chart showing the paths taken by the traffic within the scope of the peering analysis (see **Peering Sankey Chart**).
* **Group By Boundary ASN** (shown only when Direction is Inbound): A switch that determines the dimensions represented in the columns of the Peering Sankey chart.

#### Peering Sankey Chart

The **Peering Sankey** chart shows the paths taken by the traffic within the scope of the peering analysis:

* Hops are represented by colored vertical bars. Hover to see information about the hop.
* Paths are represented by gray bands. The width of a band between hops is proportional to the quantity of traffic (95th percentile Mbps) taking that path. Upon hover, a band is highlighted from source to destination and a tool tip appears with information about the path.

The columns of the Sankey vary depending on the Direction setting (see **Peering Chart Pane**):

* If **Direction** is Inbound the columns of the Sankey depend on the state of the **Group By Boundary ASN** switch:

  + **If on**: columns (left to right) are Src AS Number, Src Next Hop AS Number, Src Connectivity Type, and Site (where traffic on this path enters your network).
  + **If off**: columns are Src AS Number, Src Provider, Src Connectivity Type, and Site.
* If **Direction** is Outbound, the Sankey columns are Site and Dest Connectivity Type, followed by one or more AS Path steps.

### Potential Peers Pane

The Potential Peers pane displays a list of Autonomous Systems (ASes) with which it may be advantageous for you to peer as determined by Kentik's potential peers analysis based on your filter and direction settings. The pane includes the following UI elements:

* **Time Range selector**: The time range (back from the present) over which traffic is evaluated for the potential peers analysis.

  > **Note:** The time range setting affects only the Potential Peers table (not the Sankey chart).
* **Potential Peers table**: A list of potential peers (see **Potential Peers Table**).

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(335).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A27Z&se=2026-05-12T09%3A42%3A27Z&sr=c&sp=r&sig=QLEsfiVm7NWLKccRizgegx19XKMXaL%2BXQebOuKMJ%2FAg%3D)

*A list of ASes with whom peering might be cost-effective for your organization.*

#### Potential Peers Table

The Potential Peers table contains the following columns:

* **Peer Name**: The name of the potential peer. Click to open the **Potential Peer Page** for this potential peer.
* **AS Name**: The ASN and AS Name of the potential peer. Click to open the details page for this AS (see **Core Details Pages**).
* **Shared Facilities**: Data centers (if any) where your organization and the potential peer both have a point of presence, which would facilitate interconnection.
* **P95 Traffic**: The 95th percentile volume (in Mbps) of traffic in the currently selected direction (Inbound or Outbound; see **Peering Chart Pane**) over the specified time range.

  > **Note:** The traffic included in this calculation varies depending on the **Show Traffic Through** setting.
* **Top sites**: Varies depending on the **Direction** setting:

  + **If inbound**: Sites via which traffic from the potential peer enters your network.
  + **If outbound**: Sites via which traffic to the potential peer leaves your network.

    > **Note:** The traffic included in this calculation varies depending on the **Show Traffic Through** setting.
* **Exclude** (red **X**): Removes the potential peer from the Sankey chart and the Potential Peers list.

  > **Note:** Excluded ASes can be reinstated via the Excluded ASes list on the **Configure Discover Peers**.

## Configure Discover Peers

The **Configure Discover Peers** page, which affects how results are displayed on the Discover Peers page, is opened with the Configure button on that page. This page is covered in the following topics:

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(337).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A27Z&se=2026-05-12T09%3A42%3A27Z&sr=c&sp=r&sig=QLEsfiVm7NWLKccRizgegx19XKMXaL%2BXQebOuKMJ%2FAg%3D)

*The settings of the Defaults pane control what's displayed on the Discover Peers page.*

### Defaults Pane

The Defaults Pane determines various settings applied during the peering analysis to shape the results displayed on the **Discover Peers Page**. This pane includes the following controls:

* **Connectivity Type**: Specify the connectivity types (see **Understanding Connectivity Types**) that will be included in the results, which may be one or more of the following:

  + Transit: Connected to an upstream provider of Internet connectivity.
  + Free Private Peering: Connected to a non-transit AS using one or a bundle of direct private connections to the AS’s facing router.
  + Paid Private Peering: Connected identically to free private peering, but the arrangement involves some form of payment between the parties.
  + IX Peering: Connected to a public peering facility, which is a switching LAN on which multiple ASes offer to peer.
* **Sankey Depth**: The maximum number of ASes that will be displayed in the Sankey chart on the **Discover****Peers** page.
* **Minimum Traffic Volume**: The minimum traffic volume required for an AS to be displayed in the peering Sankey chart and the Potential Peers table. The volume is calculated over the **Time Range** specified on the **Discover Peers** page. Options are None (default), 100 Mbps, 500 Mbps, and 1 Gbps.

### Excluded ASes List

TheExcluded ASes list shows any ASes that have been excluded from the chart and table on the**Discover****Peers** page using the **Exclude** button (red **X**) in the corresponding row of the **Potential Peers Table**. To undo the exclusion of a given AS, restoring it to the chart and table, click the X at the right of the row to remove that AS from theExcluded ASes list.

> **Note:** This list is only visible when one or more ASes have been excluded.    

## Potential Peer Page

The **Potential Peer** page, which provides information about an individual potential peer, is covered in the following topics.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(338).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A27Z&se=2026-05-12T09%3A42%3A27Z&sr=c&sp=r&sig=QLEsfiVm7NWLKccRizgegx19XKMXaL%2BXQebOuKMJ%2FAg%3D)

*Details about an individual AS that is a potential peer.*

### Potential Peer Page Layout

The **Potential Peer** page includes the following main elements:

* **Time Range**: A drop-down for setting the duration (back from the present) covered by the information displayed on this page (charts, tables, etc.).
* **Filter Traffic Pane**: A set of drop-downs for filtering the traffic that is included in the information displayed on this page (see **Filter Traffic Pane**).
* **Information Pane**: Information about the potential peer, taken from the AS’s PeeringDB registration (see **Information Pane**).
* **Peer Traffic chart**: A line chart showing traffic between your network and the potential peer over the currently specified time range(see **Peer Traffic Chart**).
* **Autonomous Systems Behind chart**: This **Horizon Chart** is shown only when there is more than one unique destination ASN in the paths associated with the traffic to the potential peer. The ability to see who will ultimately consume the traffic can help inform your peering strategy:

  + If your goal for peering is primarily to offload traffic from paid links, then you can accomplish that by peering with the intermediate ASN, which is the potential peer.
  + If your goal is to reduce hops between your network and the ultimate destination (as is often the case with content providers) then you may wish to peer directly with the downstream ASNs listed in the chart.
* **Unique IPs Pane**: Statistics on the unique IPs in the traffic between your network and the potential peer during the currently specified time range (see **Unique IPs Pane**).
* **Unique Prefixes Pane**:  Statistics on the unique prefixes in the traffic between your network and the potential peer during the currently specified time range (see **Unique Prefixes Pane**).
* **Potential Peer Traffic tables**: Tables detailing traffic between your network and the potential peer (see **Potential Peer Traffic Tables**).

### Filter Traffic Pane

The **Filter Traffic** pane contains a set of drop-downs that enable you to limit the traffic that is included in the Peer **Traffic Chart**, the **Unique IPs** and **Unique Prefixes** panes, and the traffic tables (by country and by sites). The pane contains the following filters:

* **Sites**: Show only traffic that enters or exits your network at the specified site(s).
* **Countries**: Show only traffic that originates or terminates in the specified country(ies).
* **Regions**: Show only traffic that originates or terminates in the specified region(s).
* **Cities**: Show only traffic that originates or terminates in the specified city(ies).
* **BGP** **Communities**: Show only traffic that is tagged with the specified BGP Community value.

The following table shows the filters in the Filter Paths pane and how their effect varies depending on the Direction setting (see **Potential Peer Traffic Tables**):

| Filter | Inbound | Outbound |
| --- | --- | --- |
| Sites | Show only traffic that enters your network at the specified sites. | Show only traffic that exits your network at the specified sites. |
| Countries | Show only traffic that originates in the specified countries. | Show only traffic that terminates in the specified countries. |
| Regions | Show only traffic that originates in the specified regions. | Show only traffic that terminates in the specified regions. |
| Cities | Show only traffic that originates in the specified cities. | Show only traffic that terminates in the specified cities. |
| BGP Communities | Show only traffic that is tagged with the specified BGP Community value. | Show only traffic that is tagged with the specified BGP Community value. |

### Information Pane

This pane provides information that is sourced from the potential peer's listing in **PeeringDB**. The following table shows the categories of PeeringDB information that appear on each of the **Information** pane’s tabs.

| General Info Tab | Policy Info Tab | Notes |
| --- | --- | --- |
| * Protocols Supported * IRR as-set/route-set * IPv4 Prefixes * IPv6 Prefixes * Network Type * Geographic Scope * Traffic Ratios * Out * In * Traffic Levels | * View Company Website: link * General Policy * Multiple Locations * Ratio Requirement * Contract Requirement | * Notes |

### Peer Traffic Chart

Peer Traffic is a line chart showing traffic between your network and the potential peer over the currently specified time range:

* Traffic to the peer is plotted with the blue line.
* Traffic from the peer is plotted with the green line.
* Hover over either line to open a popup that states the volume of traffic in that direction at that point in the time range.
* The aggregation interval, which is stated at the bottom of the chart, varies from 10 minutes for a one-day time range to 60 minutes for a 30-day time range.

### Unique IPs Pane

This pane shows the number of unique IPs, source (top) and destination (bottom), per aggregation period (which is stated at the bottom of the**Peer Traffic** chart and varies depending on the specified time range) in the traffic between your network and the potential peer during the currently specified time range. The results are calculated using average, 95th percentile, and maximum aggregation methods.

### Unique Prefixes Pane

This pane shows the number of unique prefixes, source (top) and destination (bottom), per aggregation period (which is stated at the bottom of the**Peer Traffic** chart and varies depending on the specified time range) in the traffic between your network and the potential peer during the currently specified time range. The results are calculated using average, 95th percentile, and maximum aggregation methods.

### Potential Peer Traffic Tables

The traffic tables on the **Potential Peer Page** detail traffic between your network and the potential peer. This section of the page contains the following main UI elements:

* **Direction selector**: Determines the perspective from which the tables are populated:

  + Inbound: Show traffic from the potential peer to your network.
  + Outbound: Show traffic from your network to the potential peer.
* **Traffic by Country**: A heatmap providing traffic information broken down by the country traffic comes from or goes to (see **Traffic by Country Heatmap**).
* **External Traffic by Sites**: A table providing traffic information broken down by the site via which the traffic enters or exits your network (see **External Traffic Table**).

#### Traffic by Country Heatmap

The **Traffic by Country** heatmap shows traffic volume by country, either from (if **Direction** is Inbound) or to (if **Direction** is Outbound) the potential peer. Hover over a country to see the volume of traffic attributed to that country.

#### External Traffic Table

This table provides information about traffic between your network and the potential peer, broken down by individual devices, which are grouped in the table by the site in which they exist. Depending on the Direction setting, each row represents a device where the traffic enters your network (Inbound) or exits your network (Outbound). The table includes the following columns:

* **Site**: The name of a site in which you have devices handling traffic from (Inbound) or to (Outbound) the potential peer.
* **Device**: The device in this site via which the potential peer’s traffic enters (Inbound) or exits (Outbound) your network.
* **Connectivity Type**: The source (Inbound) or destination (Outbound) connectivity type  as determined by Interface Classification (see **Understanding Connectivity Types**).
* **Provider**: The source (Inbound) or destination (Outbound) provider as determined by **Provider Classification**.
* **Network Boundary**: The source (Inbound) or destination (Outbound) network boundary as determined by Interface Classification (see **Network Boundary Attribute**).
* **Average Mbits/s**: Average traffic volume via this site over the specified time range.
* **p95th Mbits/s**: 95th percentile traffic volume via this site over the specified time range.
* **Max Mbits/s**:  Maximum traffic volume via this site over the specified time range.

## Using Discover Peers

As described in the following brief walk-throughs, using **Discover Peers** involves both the overall **Discover Peers** page and the individual Potential Peer pages.

#### Using the Discover Peers Page

To build a list of ASes with which it might be beneficial to peer:

1. From the main menu, navigate to the **Discover Peers Page** (Edge » **Discover Peers**).
2. Click the **Configure** button to check the settings on the **Configure Discover Peers** page, then use the breadcrumb to return to the Discover Peers page.
3. Set any desired filters in the **Filter Paths Pane**.
4. Set the Direction selector in the **Peering Chart Pane**. At this point the **Peering Chart** (Sankey) shows paths in the specified direction, and the **Potential Peers Table** shows ASes determined by Kentik to be promising peers for your network.
5. To evaluate a specific potential peer, see **Using the Potential Peer Page**.

#### Using the Potential Peer Page

To explore a specific potential peer in detail:

1. On the **Discover Peers** page, click on a **Peer Name** in the **Potential Peers** table to open that peer's **Potential Peer Page**.
2. Set the **Time Range** for the information displayed on the page (see **Potential Peer Page Layout**).
3. Set any desired filters in the **Filter Traffic Pane**.
4. You'll now be able to see:

   1. Information about the potential peer, taken from the AS’s PeeringDB registration (see **Information Pane**).
   2. Traffic between your network and the potential peer (see **Peer Traffic Chart**).
   3. The ASes behind the potential peer (see **Autonomous Systems Behind** chart in **Potential Peer Page Layout**).
   4. Statistics on the unique IPs (see **Unique IPs Pane**) and unique prefixes (see **Unique Prefixes Pane**) in the traffic between your network and the potential peer.
   5. Traffic between your network and the potential peer broken down by country (see **Traffic by Country Heatmap**) and by sites (see **External Traffic Table**).
