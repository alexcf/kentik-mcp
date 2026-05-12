# Kentik Market Intelligence

Source: https://kb.kentik.com/docs/kentik-market-intelligence

---

This article discusses the **Kentik Market Intelligence** feature of the Kentik portal.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(230).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A05Z&se=2026-05-12T09%3A58%3A05Z&sr=c&sp=r&sig=F%2BlbDBwGl2FLgte%2B9sjW2Wk93Ts1ccsULONyd6dlzT0%3D)

*KMI provides BGP-derived rankings of ASes in terms of customer base and peering.*

## About KMI

Kentik Market Intelligence (KMI) is a Service Provider workflow that uses the global routing table to classify the peering and transit relationships between Autonomous Systems (ASes) and to identify the providers, peers, and customers for any AS in any geography. KMI estimates the volume of IP space transited by ASes in different geographies and produces rankings based on that volume, thereby enabling users to compare ASes in various markets.  
  
KMI also features **Top Insights**, which are automatic real-time notices about networks and markets. KMI insights are currently presented only within KMI rather than being part of the overall **Kentik Insights** engine. KMI insights identify interesting or anomalous events or trends that can deepen your understanding of the networks involved and the markets in which they operate.  
  
KMI requires no setup or configuration, does not use flow data or synthetic test data, and covers ASes that both are and are not Kentik customers. Kentik draws public RIB data feeds from the University of Oregon Route Views project, computes scores from this information (see **KMI Ranking Methodology**), and publishes the results once daily on the Kentik Market Intelligence page (see **KMI Page**) as a set of tables in which ASes are ranked based on their score. A multi-tabbed details page for each AS (see **AS Details Page**) enables deeper drill-down into specific information for individual networks, including charts and tables in areas such as ranking, markets, customers, providers, and peers.

### KMI Use Cases

In its initial iteration, KMI is oriented toward service providers that sell connectivity (a.k.a. Transit Providers), where it can be a valuable tool for a number of roles:

* Product Managers can use it to assess competition in existing and new markets.
* Marketing departments at leading providers can use it to validate their position in various markets.
* Sales and Product executives can use it to set goals in markets defined by customer type or by territory.
* Sales can use it to identify single-homed ASes to prospect and to determine existing providers for a given prospect.

KMI is also useful to networks such as content providers and large broadband providers, particularly for Peering Managers, who need to know all providers used by a given network:

* KMI can help identify networks with which it would be advantageous to peer.
* KMI can help determine ways to influence inbound path for a specific source network by using community routing to see them from another provider.

### KMI Ranking Type

Kentik Market Intelligence ranks ASes in order of three main types of Kentik-computed KMI rankings (see **KMI Ranking Methodology**). The rankings displayed on the **KMI Page** or an **AS Details Page** take into account the current **KMI Filter Bar** settings on that page. The following types of rankings are shown:

* **Customer Base**: Ranked by the size of the overall customer base, estimated by determining how much IP address space a given AS transits relative to other ASes. Rankings are also available by category (see **Customer Base Type**).
* **Customer Growth**: Ranked by the change in overall customer base (gain/loss of prefixes) over the last 28 days.
* **Peering**: Ranked by the amount of IP address space sent to the AS over a settlement-free peering session.

On AS Details pages, the various ranking types are found on the following tabs:

* **Overview**: Customer Base (including categories).
* **Rankings**: Customer Base (including categories).
* **Markets**: Customer Base (including categories), Customer Growth, Peering.

## KMI Filter Bar

The filter bar is a set of controls that appears near the top of the tabs in the KMI module (on the **KMI Page** or on an individual **AS Details Page**). The controls in the bar allow you to define the scope of the information presented on the tab.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(233).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A05Z&se=2026-05-12T09%3A58%3A05Z&sr=c&sp=r&sig=F%2BlbDBwGl2FLgte%2B9sjW2Wk93Ts1ccsULONyd6dlzT0%3D)

Depending on the tab (see **Filter Controls by Tab**), the filter bar may include the following individual controls:

* **IP Address Family**: A drop-down from which you can choose the IP version (v4 or v6) used for the evaluation of IP address space.

  > **Note:** KMI scores are not directly comparable between IP versions. Kentik currently has no plans to enable an IP version setting of All.
* **Geo Market**: A drop-down from which you can choose the geographical scope of the KMI analysis presented on this page.
* > **Note:** Kentik uses GeoIP to determine which ASes are included when you filter with this control.
* **Ranking** **Type**: A drop-down from which you can select a Rankings type (see **KMI Ranking Type**) by which to narrow the scope of the rankings in the tab's tables and charts. When narrowing by customers, the control allows you to choose a specific **Customer Base Type**.
* **Comparison ASN**: A drop-down from which you can select one of your organization’s ASNs (as entered in your **Network Classification** settings) to compare with the specified AS. This control is only available on AS Details pages (see **Filter Controls by Tab**).

The results on each tab are filtered based on the above settings.

> **Note:** Filter bar settings are retained from page to page, including from the landing page to a details page (and vice versa) as well as from one details page to another.

#### Customer Base Type

On the KMI landing page, and on several tabs of the AS Details pages, customer base rankings may be broken down further by category, which is chosen from the drop-down **Ranking Type** control in the **KMI Filter Bar**. In addition to all customers, the following categories are available:

* **Retail**: Networks that provide services (e.g. originate content) or whose end-users are consumers of services (e.g., ISPs or "eyeball" networks).
* **Wholesale**: Networks that connect retail networks to backbone networks.
* **Backbone**: Networks that carry high volumes of traffic between wholesale networks.

  ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(235).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A05Z&se=2026-05-12T09%3A58%3A05Z&sr=c&sp=r&sig=F%2BlbDBwGl2FLgte%2B9sjW2Wk93Ts1ccsULONyd6dlzT0%3D)

  *Customer base types include Retail, Wholesale, and Backbone.*

#### Filter Controls by Tab

The table below shows how the filter bar controls vary depending on the tab.

| Tab | IP Address Family | Geo Market | Ranking Type | Comparison ASN |
| --- | --- | --- | --- | --- |
| ***KMI page (landing):*** | | | | |
| Customer Base | Yes | Yes | Yes | No |
| Customer Growth | Yes | Yes | Yes | No |
| Peering | Yes | Yes | Yes | No |
| ***AS Details pages:*** | | | | |
| Overview | Yes | Yes | Yes | Yes |
| Rankings | Yes | Yes | Yes | Yes |
| Markets | Yes | No | Yes | No |
| Customers & Providers | Yes | Yes | No | Yes |
| Peers | Yes | Yes | No | Yes |

## KMI Page

The **Kentik Market Intelligence** page includes the following main settings, controls, and information:

* **Export** (on subnav):A button that opens the Market Intelligence CSV Export dialog, enabling you to export the **Latest Rankings** table as CSV data:

  + Start Rank: Enter the highest rank to include in the exported data.
  + End Rank: Enter the lowest rank to include in the exported data.
* **Insights** (on subnav):A button that toggles visibility of the right-side **Top Insights** drawer (see **Top Insights**) which contains market and ranking insights based on your **Top Insights Configuration**.
* **Top insights**: A pane, displayed in a drawer, that contains KMI insights (see **Top Insights**).
* > **Note:** The drawer is open by default; to close it, click the **Insights** button.
* **Search by**: A field into which you can enter an ASN or AS name to return a drop-down list of matching ASes. Click on an AS in the resulting list to go to the **AS Details Page** for that AS.
* **Configuration**: A button that opens the Market Intelligence Configuration dialog (see **KMI Configuration**).
* **Filter Bar**: A set of controls across the top of the chart that enable you to define the scope of the information presented on the page (see **KMI Filter Bar**).
* **Ranking Over Time**: A pane containing a chart and list that show, based on the current **Filter Bar** settings, the rankings of ASes in the category corresponding to the selected tab (see **Ranking Over Time Pane**).
* **Latest Rankings:** A pane containing a table listing top providers and a filter control that lets you narrow, by ranking, the range of providers that are shown (see **Latest Rankings Pane**).

### Ranking Over Time Pane

The **Ranking Over****Time** pane shows a chart and a list, both of which are filtered based on the settings of the **KMI Filter Bar**:

* **Ranking chart**: A time-series chart showing the ranking of providers for each month of the last year.
* **Ranking list**: The current ranking (post-filter) of providers (see **Ranking Over Time List**).

  ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(237).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A05Z&se=2026-05-12T09%3A58%3A05Z&sr=c&sp=r&sig=F%2BlbDBwGl2FLgte%2B9sjW2Wk93Ts1ccsULONyd6dlzT0%3D)

  *Ranking over time shows how the current top-ranked providers have ranked over the last year.*

The providers that are ranked in the chart and list depend on which mode is selected with the controls at the top right (above the list):

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(238).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A05Z&se=2026-05-12T09%3A58%3A05Z&sr=c&sp=r&sig=F%2BlbDBwGl2FLgte%2B9sjW2Wk93Ts1ccsULONyd6dlzT0%3D)

* **Top 10 Providers** (default): The list and chart show the current top-10 of the post-filter providers.
* **Select Provider:** The list and chart show five providers on either side (in terms of ranking) of the provider selected with the **See +/-5** around drop-down. For example, if the ranking of the selected provider is 100, the other displayed providers will be those with rankings from 95-99 and 101-105. If no provider is selected, the displayed providers will be the current top-10 of the post-filter providers.

When you hover over a given location in the chart, a popup appears that shows what the ranking was for the month at that point in the time range.

> **Note:** For information on how rankings are derived, see **KMI Ranking Methodology**.

#### Ranking Over Time List

Each row of the **Ranking Over Time** list includes the following (left to right):

* **Color key**: The color that represents this provider in the **Ranking Over Time** chart.
* **Provider logo**: The provider’s logo (if included).
* **Name**: The name of the provider’s owner/operator.
* **ASN**: A lozenge giving the provider’s ASN. Click the lozenge to open an **ASN Links Popup** that includes the AS’s country of registration and links to the following AS-related destinations in Kentik:

  + **Overview** tab of its **AS Details Page** in KMI
  + **PeeringDB Info Tab** of its Core Details page
  + **Traffic Tab** of its Core Details page

Click any row of the **Ranking Over Time** list to go to the **AS Details Page** for that provider

### Latest Rankings Pane

The **Latest Rankings** pane of the Kentik Market Intelligence page is made up of the Latest Rankings table and its associated range filter, which are covered in the topics below.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(239).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A05Z&se=2026-05-12T09%3A58%3A05Z&sr=c&sp=r&sig=F%2BlbDBwGl2FLgte%2B9sjW2Wk93Ts1ccsULONyd6dlzT0%3D)

#### Latest Rankings Table

The **Latest Rankings** table shows the current ranking of providers, filtered based on the settings of the **KMI Filter Bar**. Click the row for any provider to go to that provider’s **AS Details Page** in KMI. The table includes the following columns (left to right):

* **Ranking**:

  + Rank: The ranking currently assigned to each AS, taking into account the page's current **IP Address** **Family**, **Geo Market**, and **Ranking Type** settings.
  + Change: The change in rank since the previous ranking (currently daily).
* **Provider**:

  + Logo: The provider’s logo.
  + Name: The name of the provider's owner/operator.
  + ASN: A lozenge giving the provider’s ASN. Click the lozenge to open a popup that includes the AS’s country of registration and links to AS-related destinations in Kentik (see ASN in **Ranking Over Time List**).
* **% of top score**:

  + Percentage: This provider's score as a percentage of the score of the top-ranked provider listed in the table.
  + Bar chart: A horizontal bar chart illustrating this provider's score percentage.

> **Note:** For information on how rankings are derived, see **KMI Ranking Methodology**.

#### Latest Rankings Range Filter

The range filter is a blue bar that appears above the **Latest Rankings** table. A series of dots is displayed under the bar, with top-ranked providers represented as individual dots and the remaining providers, whose scores are too close to be represented individually, represented with one large dot. The position of each dot from left to right is proportional to the corresponding provider's **% of top score** as shown in the table. Hovering over a dot opens a popup that shows you the provider(s) and its percentage of the top score.  
  
The range control has two handles, which appear by default at the far left and right ends of the bar. Dragging these handles has the following effect on which providers are listed in the table:

* Sliding the left handle to the right trims the highest-ranked providers from the top of the table.
* Sliding the right handle to the left trims the lowest-ranked providers from the bottom of the table.

### KMI Configuration

KMI is configured in the Market Intelligence Configuration dialog that opens from the Configuration button on the Kentik Market Intelligence page. When you click in the field on the **My ASNs** tab, you'll see a list of the ASNs that you have defined in Settings » **Network Classification** as being part of your network. Select all of these ASNs that you want KMI to consider as part of your network. The ASNs will be indicated as yours in tables and will be included in comparisons involving your network as a filter (see **Filter Controls by Tab**).

> **Note:** To change the ASNs that are available to designate as part of your network, change your **Network Classification** settings    

## AS Details Page

Every Autonomous System (AS) is detailed in KMI on a multi-tabbed page that enables deeper drill-down into specific information for individual networks.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(242).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A05Z&se=2026-05-12T09%3A58%3A05Z&sr=c&sp=r&sig=F%2BlbDBwGl2FLgte%2B9sjW2Wk93Ts1ccsULONyd6dlzT0%3D)

*The multi-tabbed AS Details page enables a deep dive into an individual AS.*

### AS Details UI

The details page for each AS includes the following UI elements:

* **AS title bar**: The logo, name, country, and Autonomous System Number (ASN) of this AS.
* **View in Network Explorer**: A button that takes you to the Traffic tab of this AS’s Details page in Network Explorer (see **Traffic Tab**).
* **View PeeringDB Record**: A button that takes you to the PeeringDB tab of this AS’s Network Explorer Details page (see **PeeringDB Info Tab**).
* **Search by**: A field into which you can enter an ASN or AS name to return a drop-down list of matching ASes. Click on an AS in the resulting list to go to the KMI Details page for that AS.
* **Tabs**: A set of tabs presenting different types of information about the AS (**Overview**, **Rankings**, **Markets**, **Customers & Providers**, and **Peers**). The charts and tables on each tab are covered in the topics below.
* **Filter bar**: A filter bar at the top of each tab enables you to define the scope of the information presented on the tab (see **KMI Filter Bar**). To see which controls are included on which tab, see **Filter Controls by Tab**.

  > **Note:** Filter bar settings are retained from page to page, including from the landing page to a Details page (and vice versa) as well as from one Details page to another.

### AS Overview Details

The **Overview** tab provides a high-level look (visualized in charts) at metrics for this AS, which are each compared (when presented as a bar graph) with the same metric for your own network. The **KMI Filter Bar** includes scope controls for IP Address Family, Geo Market, Ranking Type, and Comparison ASN (see **Filter Controls by Tab**). Unless noted below, each chart will update when you make changes to the filter controls.  
  
 The following panes and charts are included on this tab: ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(245).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A05Z&se=2026-05-12T09%3A58%3A05Z&sr=c&sp=r&sig=F%2BlbDBwGl2FLgte%2B9sjW2Wk93Ts1ccsULONyd6dlzT0%3D)

* **Footprint**: Displays bar charts comparing the geographical distribution of traffic on this AS with that of your own network:

  + Regions: The number of regions from/to which the AS transits traffic.

    > **Note**: This chart is not affected by the Geo Market or Ranking Type filter settings.

    - Countries: The number of countries (identified via GeoIP data for the prefix being transited) from which the AS transits traffic.
* **Providers**: Displays the Provider Count by Month line chart, which shows the AS’s upstream provider count for the last 12 months.

  > **Notes**:
  >
  > + If there are no upstream providers (e.g., the AS is part of the Transit Free Zone), the chart is not displayed.
  > + The provider count on the graph’s Y axis may be truncated and as such, may not start at zero.
  > + This chart is not affected by the Ranking Type filter setting.
* **Customers**: Displays charts showing current and historical data about the AS’s customers (see **Overview Customers Charts**).
* **Prefixes/IPs**: This section shows the originated and transited prefixes and IPs for the AS (see **Overview Prefixes Charts**).
* **Top Insights**: A pane that lists insights related to this AS (see **Top Insights**). The configuration of the pane when you arrive at the AS Details page will be the same as it was on the KMI landing page. If you change the configuration (see **Top Insights Configuration**), the updated configuration remains when you return to the KMI page.

#### Overview Customers Charts

The **Customers** pane of the **Overview** tab includes the following charts:

* **Customers**: The number of ASes identified as customers of this AS compared with your own network.
* **Single-Homed Customers**: The number of ASes identified as single-homed customers of this AS (see Single Homed ASes) compared with your own network.
* **Customer Count by Month**: A stacked area chart that shows both single-homed and multi-home customer counts for the AS over the last 12 months (hover over any point of the graph to see exact values and to see which points are single or multi-home).

  > **Note:** The customer count on the graph’s Y axis may be truncated and as such, may not start at zero.

* **Customer Base Split**: A breakdown of the AS's customers by customer type (retail, wholesale, or backbone).

  > **Note:** This chart is not affected by the Ranking Type filter setting.

#### Overview Prefixes Charts

The **Prefixes/IP** pane of the **Overview** tab includes the following charts:

* **Transited Prefixes**: A bar chart comparing, for this AS and your own network, the number of unique prefixes received from another network and passed to a downstream network.
* **Transited IPs**: A bar chart comparing, for this AS and your own network, the number of unique IPs received from another network and passed to a downstream network.
* **Originated Prefixes**: A bar chart comparing, for this AS and your own network, the number of unique prefixes originated (owned and announced).
* **Originated IPs**: A bar chart comparing, for this AS and your own network, the number of unique IPs originated (owned and announced) by this AS.
* **Transited Prefix Distribution**: A histogram that breaks down the distribution of the transited prefixes for the current month that are shown in the Transited Prefixes graph (above). The darker the color on the histogram, the more prefixes are being represented (hover over the bottom number to see the exact number for that prefix).
* **Originated Prefix Distribution**: A histogram that breaks down the distribution of the originated prefixes for the current month that are shown in the Originated Prefixes graph (above). The darker the color on the histogram, the more prefixes are being represented (hover over the bottom number to see the exact number for that prefix).

> **Note:** Origination data (Originated Prefixes, Originated IPs, and Originated Prefix Distribution visualizations) will not appear when the KMI Ranking Type is Customer Base: Wholesale or Customer Base: Backbone.

#### Single Homed ASes

An AS that is "single-homed" — also referred to as "critically dependent" — has only one upstream provider to the internet. If this single point of failure actually fails, a single-homed network is effectively cut off from the internet.

### AS Rankings Details

The **Rankings** tab provides a high-level look (visualized in charts and tables) at the customer and peering ranking for the selected AS. The **KMI Filter Bar** includes scope controls for IP Address Family, Geo Market, Ranking Type, and Comparison ASN (see **Filter Controls by Tab**).

#### Rankings Details Charts

Each chart on the **Rankings** tab measures the selected ranking (Customer Base or Peering) on the graph’s Y axis, which may be truncated and as such, may not start at zero. The Y axis also starts with lower values at the top of the chart increasing as it moves down to meet the X axis. The following charts are included on this tab:

* **Customer Base Ranking by Month**: A line graph that shows the AS’s Customer Base Ranking for the last 12 months (see **KMI Ranking Type**). Hover over any point of the graph to see exact values.
* **Peering Ranking by Month**: A line graph that shows the AS’s Peering Ranking for the last 12 months (see **KMI Ranking Type**). Hover over any point of the graph to see exact values.

#### Rankings Details Tables

The tables of the **Rankings** tab (Customer Base, Peering, and Customer Growth) each include the highest-ranked AS in the category (labelled as Top), the AS whose details page you’re viewing, the five ASes immediately above and below the selected AS (when possible), and your own ASes (labelled as My Network). If neither the selected AS nor your own AS is ranked, the top 10 ASes for that table will display and both your AS and the featured AS will display at the bottom of the table (labelled as No Ranking). Click on any AS in a table to switch to the **Rankings** tab for that AS.  
  
 The following tables are included on this tab:

* **Customer Base**: This table shows where this AS (outlined in blue) falls in the overall rankings for Customer Base (see **KMI Ranking Type**).
* **Peering**: This table shows where this AS (outlined in blue) falls in the overall rankings for Peering (see **KMI Ranking Type**).
* **Customer Growth**: This table shows where this AS (outlined in blue) falls in the overall rankings for Customer Growth (see **KMI Ranking Type**).

### AS Markets Details

The **Markets** tab shows the strength of this AS by geographical market, based on a count of IP space in each market. This geo markets breakdown is represented with the following UI elements:

* **Continents gauge**: A horizontal bar chart showing this AS's share of customer IP space in each continent.
* **Heatmap**: A world map in which the color applied to each country indicates the IP space announced to this AS as a percent of the total IP space announced to all ASes operating in that country.
* **Countries list**: A list of the countries, grouped by region, in which this AS has customer networks, with a count of the customers in each country.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(246).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A05Z&se=2026-05-12T09%3A58%3A05Z&sr=c&sp=r&sig=F%2BlbDBwGl2FLgte%2B9sjW2Wk93Ts1ccsULONyd6dlzT0%3D)

*The heatmap on the Markets tab shows the strength of an AS by country.*

### AS Customers & Providers

The **Customers & Providers** tab identifies the providers upstream of the AS and the customers downstream and shows you which of those are common to your network. The tab includes the following charts and tables:

* **Export** (in subnav): A button that opens the **Market Intelligence CSV Export** dialog.
* **Providers chart**: A bar chart comparing this AS with your own network in terms of the number of upstream providers.
* **Customers chart:** A bar chart comparing this AS with your own network in terms of the number of downstream customers.
* **Providers list**: A table listing the upstream providers to this AS (see **Providers List**).
* **Customers list**: A table listing the downstream customers of this AS (see **Customers List**).

#### Market Intelligence CSV Export

This dialog enables you to configure the export of data in CSV format from the **Customers & Providers** tab. The following options are provided:

* **Include Providers**: If checked, providers data is included in the export; otherwise only customers data is included.
* **Exclude Mutual Customers**: Exclude from the data any customers that are common to the AS and your own network.
* **Only Single-Homed Customers**: Include in the data only the AS's single-homed customers.

#### Providers List

The Providers list is a table listing the upstream providers to this AS. The list includes the following elements:

* **Number**: An indicator, displayed to the right of the heading, of the number of providers listed, which is the total number of providers unless the list is filtered with the following controls.
* **Mutual selector**: A drop-down with which you choose which providers the list includes:

  + All Including: All providers to this AS.
  + Only: Only providers in common to this AS and your network.
  + Exclude: Only providers not in common.
* **Filter field**: Filters the list to providers whose name or ASN includes the entered text.

> **Note:** If the AS is part of the Transit Free Zone, the Providers list is replaced by a notification that the AS has no upstream providers.

#### Customers List

The Customers list is a table listing the downstream customers of this AS. The list includes the following elements:

* **Number**: An indicator, displayed to the right of the heading, of the number of customers listed, which is the total number of customers unless the list is filtered with the below controls.
* **Mutual selector**: A drop-down with which you choose which providers the list includes:

  + All Including: All customers of this AS.
  + Only: Only customers in common to this AS and your network.
  + Exclude: Only customers not in common.
* **Single Homed selector**: A drop-down with which you choose which providers the list includes:

  + All Including: All customers of this AS.
  + Only: Only customers that are single-homed.
  + Exclude: Only customers that are not single-homed.
* **Filter field**: Filters the list to customers whose name or ASN includes the entered text.

#### Providers and Customers Rows

Each row of both the Providers list and the Customers list includes the following information:

* **Ranking**: The ranking of the AS among the total listed providers or customers. If you filter either list using a selector (**Mutual** or **Single Homed**), the filtered ASes will display their original ranking as it was displayed in the full list of ASes.
* **Name**: The name of the AS.
* **ASN**: A lozenge giving the AS number of the provider or customer. Click the lozenge to open a popup that includes the AS’s country of registration and links to related destinations in Kentik (see ASN in **Ranking Over Time List**).
* **Mutual or Single Homed**: A lozenge that's displayed to the left of the **% of top score** bar chart when the AS has a mutual or single-homed provider or customer.
* **% of top score**: A horizontal bar chart illustrating the percentage of the top score (in blue) for the provider or customer. Hover over the bar to see an exact percentage.

### AS Peers Details

The **Peers** tab shows a chart comparing the AS's peers with your own, and a Selected Peers list. Scope controls are provided for IP Version, Geo Market, and Comparison ASN (see **KMI Filter Bar**). The tab includes the following charts and tables:

* **Total Peers**: A chart comparing the AS's total count of peers with your own organization’s peers. Use the **Comparison ASN** control (see **KMI Filter Bar**) to change which of your organization’s AS’s is compared.
* **Selected Peers**: A table listing the ASes with which this AS peers. The list includes the following controls and information:

  + Filter field: Filters the list to peers whose name or ASN includes the entered text.
  + Name: The name of the AS.
  + ASN: A lozenge giving the AS number of the provider. Click the lozenge to open a popup that includes the AS’s country of registration and links to related destinations in Kentik (see ASN in **Ranking Over Time List**).
  + Prefixes: The number of unique prefixes announced.

    > **Note**: Click on any row of the table to go to the Peers tab on the selected AS’s KMI details page.

## Top Insights

KMI Insights are automatic real-time notices about networks and markets. These insights are currently presented only within KMI rather than being part of the overall **Kentik Insights** engine. KMI insights identify interesting or anomalous events or trends that can deepen your understanding of the networks involved and the markets in which they operate.  
  
 KMI insights are presented in the **Top Insights** pane, which appears in two locations within the KMI module:

* **Kentik Market Intelligence page**: The pane appears in a drawer on the landing page whose visibility is toggled via the **Insights** button on the subnav (see **KMI Page UI**).
* **AS Details page**: The pane appears at the right of the Overview tab of the AS Details page (see **AS Overview Details**). In this setting it contains only insights that pertain to the AS detailed on the page.

  > **Notes:**
  >
  > + The configuration of the pane when you arrive at the AS Details page will be the same as it was on the KMI landing page. If you change the configuration (see **Top Insights Configuration**), the updated configuration remains when you return to the KMI page.
  > + In addition to seeing KMI insights in the KMI module, you can also add the **Top Insights** pane as a visualization on your **Observation Deck**.

### Top Insights UI

The Top Insights pane includes the following main parts: ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(248).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A05Z&se=2026-05-12T09%3A58%3A05Z&sr=c&sp=r&sig=F%2BlbDBwGl2FLgte%2B9sjW2Wk93Ts1ccsULONyd6dlzT0%3D)

* **Geo Market indicator** (shown only if **Use Geo Filter**is checked in configuration): In parentheses to the right of the title, this indicator tells you that the list is filtered to show only insights related to the market set with the page's **Geo Market** drop-down (see **KMI Filter Bar**).
* **Options menu** (ellipsis): A drop-down from which you can select the following options:

  + Configure: Opens the **Configure Insights** pane, replacing the **Top Insights** pane (see **Top Insights Configuration**).
  + Add to Observation Deck: Adds a **Top** **Insights** pane as a visualization in your **Observation Deck**.
* **Top Insights list**: A list of insights, each of which is presented in an **Insights Card**.

#### Insights Card

The information for each insight is presented in a card that includes the following elements:

* **Magnitude**: A five-segment indicator in which the number of segments (starting from the left) colored blue indicates the importance of the insight from least important (1) to most (5). Importance is most often correlated with a change in the prefix score, either plus or minus, when measured between two networks mentioned in the insight (see **KMI Scoring Procedure**).

  > **Note:** Insights are listed in order of magnitude (high importance to low).
* **Insights type**: A lozenge stating the insight's type (see **KMI Insight Types**):

  + If the insight's information is positive (e.g., a customer has been gained), the lozenge is green with an up-trend arrow.
  + If the information is negative (e.g., a customer has been lost), the lozenge is red with a down-trend arrow.
* **Time stamp**: How long ago the insight was surfaced.
* **Description**: A brief sentence explaining what happened, which varies depending on the insight type. The sentence will include a blue ASN lozenge; click it to open a popup that includes the AS’s country of registration and links to related destinations in Kentik (see ASN in **Ranking Over Time List**).

### Top Insights Configuration

The insights in the **Top Insights** pane are configured in the **Configure Insights** ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(249).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A05Z&se=2026-05-12T09%3A58%3A05Z&sr=c&sp=r&sig=F%2BlbDBwGl2FLgte%2B9sjW2Wk93Ts1ccsULONyd6dlzT0%3D)pane, which is opened via the **Options** menu (see **Top Insights UI**). The pane includes the following settings:

* **Use Geo Market Filter**: A checkbox that, when checked, filters insights to those that are relevant to the selected **Geo Market** filter (see **KMI Filter Bar**). If the checkbox is checked, the current Geo Market will be shown in parenthesis at the top of the Top Insights pane. If the box is unchecked, results for all markets are shown.
* **Lookback range**: A set of radio buttons that determines how old insights will be when they are dropped from the **Top Insights** list: 24 hours, 7 days, or 30 days.
* **Minimum Magnitude**: A slider that lets you set the magnitude below which an insight won't be displayed in the list.
* **Insight Types**: A series of checkboxes with which you can specify the types of insights you'd like to include in the list (see **KMI Insight Types**).
* **Save**: A button that saves configuration changes and returns you to the **Top Insights** pane.
* **Cancel**: A button that cancels configuration changes and returns configuration settings to what they were before the **Configuration Insights** pane was opened. You will return to the **Top Insights** pane.

### KMI Insight Types

KMI currently supports the following types of insights:

* **Market Entrance:** An AS is newly transiting prefixes that are geolocated to the market currently set with the Geo Market filter. The insight includes the AS's rank in that market.
* **Market Exit**: An AS that was transiting prefixes that are geolocated to this market is no longer doing so.
* **Ranking Gain**: An AS has increased in rank relative to other ASes transiting prefixes geolocated to this market. This can either be caused by an increase in the address space transited or because another AS that was ranked higher has dropped due to a decrease in the amount of address space from this market. The insight shows how many places it rose, and its new ranking.
* **Ranking Loss**: An AS has decreased in rank relative to other ASes transiting prefixes geolocated to this market. This can either be caused by a decrease in the address space transited or because another AS that was ranked lower has moved up the rankings due to an increase in the amount of address space transited from this market. The insight shows how many places it fell, and its new ranking.
* **Start Prefix Origination:** An AS has begun to originate one or more prefixes in the geolocation as determined by the Geo Market filter.
* **Stop Prefix Origination:** An AS is no longer originating one or more prefixes in the geolocation as determined by the Geo Market filter.
* **Increase Prefix Origination**: An AS has increased the amount of IP space originated in this geolocation. The insight shows all providers that have increased prefix origination and by what percentage.
* **Decrease Prefix Origination**: An AS has decreased the amount of IP space originated in this geolocation. The insight shows all providers that have decreased prefix origination and by what percentage.
* **Customer Gain**: A new AS has been classified as a transit customer of this AS. The insight shows the name of the provider and the name of the customer it gained.
* **Customer Loss**: An AS is no longer classified as a transit customer of this AS. The insight shows the name of the provider and the name of the customer it lost.
* **Customer Routing Gain**: A customer AS has increased the IP space transited by this AS. The insight shows the name of the provider, the name of the customer, and the percentage by which routing was increased.
* **Customer Routing Loss**: A customer AS has decreased the IP space transited by this AS. The insight shows the name of the provider, the name of the customer, and the percentage by which routing was decreased.
* **Provider Gain**: A new AS has been classified as a transit provider of this AS.
* **Provider Loss**: An AS is no longer classified as a transit provider of this AS.
* **Provider Routing Gain:** A provider AS has increased the IP space transited from this AS. The insight shows the name of the AS, the provider to which it increased routing, and by what percentage.
* **Provider Routing Loss**: A provider AS has decreased the IP space transited from this AS. The insight shows the name of the AS, the provider to which it decreased routing, and by what percentage.

> **Note:** In the above descriptions, "this AS" refers to the AS of an **AS Details Page**.      

## KMI Ranking Methodology

How does Kentik assign rankings to the ASes listed in KMI? The key assumption behind our methodology is that the volume of transited internet traffic is correlated to the amount of transited IP address space. In other words, the more transited IP address space, the more potential for transited internet traffic and the higher the transit score.  
  
Depending on how deeply one delves into the details, the process by which we apply the above assumption can be summarized as follows:

* **Short explanation**: Each AS receives a score based on the amount of IP address space it transits for a given geography. ASes are ranked using these scores.
* **Longer explanation**: KMI rankings are based on the relative amount of transited IP space by each AS for a given market. While we cannot directly measure the volume of internet traffic routed by any given AS on the internet, we can estimate the size of the customer base by determining how much IP address space it transits relative to other ASes. We can do this for any AS in any market without the need for a direct peering relationship for route collection.

A step-by-step description of the process is provided in **KMI Scoring Procedure**.

### BGP Data Source

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/KMI-Route_views_logo-150h150w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A05Z&se=2026-05-12T09%3A58%3A05Z&sr=c&sp=r&sig=F%2BlbDBwGl2FLgte%2B9sjW2Wk93Ts1ccsULONyd6dlzT0%3D)The computation covered in **KMI Scoring Procedure** is performed daily (with greater frequency planned in the future). The BGP data we use is collected from the public **Route Views** project at University of Oregon, which collects BGP data from global vantage points around the world.

> **Note:** Future plans for KMI include expanding vantage points and visibility by augmenting the Route Views dataset with data from our own BGP sessions with customers.

### KMI Scoring Procedure

The ranking of ASes for KMI is a multi-step process, which is covered in the topics below.

#### Step 1. Classify AS-AS Relationships

Every AS-AS adjacency is classified as either transit or peering through a machine learning algorithm that accounts for several factors. We start with a known set of ASes that make up the default-free zone (DFZ), which is the handful of companies (e.g., NTT, Arelion, Cogent, etc.) that don’t buy transit from any other providers. With rare exceptions, AS paths exhibit a “valley-free” property that can be used to determine the transit relationships between ASes in the global routing table.  
  
 To illustrate the logic, consider the below example of ASx, which passes a route to DFZ AS1, which then passes the route to DFZ AS2:  
 `... « DFZ AS2 « DFZ AS1 « ASx`  
  
Because DFZ ASes don’t buy transit, ASx can’t be a provider to DFZ AS1. It also can’t be a peer to DFZ AS1 without violating the valley-free rule, because relationships between DFZ ASes must be peering and two successive peering relationships are not allowed. ASx must therefore be a transit customer of DFZ AS1, and any ASes that appear to the right of ASx in an AS path should be classified as customers of ASx solely by virtue of the valley-free rule.  
  
The above pattern can be applied until every AS-AS relationship is classified as either transit or peering. In so doing, a number of heuristics are employed to account for exceptions to the valley-free property. For example, the AS of an IXP can appear in an AS path when it is simply serving as the “glue” between two peers rather than as a direct participant in the relationship. In these cases, the IXP AS is removed from consideration in the evaluation of the relationship of the adjacent ASes.

#### Step 2. Score Routed Prefixes

Every routed prefix is assigned a score based on its prefix length:

* **IP v4**: The weight of one prefix is (1 + 24 - masklen)3 for *masklen* between 8 and 24 inclusive (prefixes outside this range are ignored).
* **IP v6**: The weight of one prefix is currently (1 + 48 - masklen)3 for *masklen* between 24 and 48 inclusive. However v6 prefixes between /16 and /24 get the same score, as well as prefixes between /48 and /64, which get a score of 1.

#### Step 3. Assign Geo and Transit Scores

Each prefix propagated across an AS-AS adjacency that has been classified as transit is assigned both a geolocation and a transit score that determine how the prefix contributes to the transit score of the provider of the adjacency:

* **Geolocation**: The geolocation is typically a country, which results in the prefix score contributing to the provider’s country, subcontinent, continent, and global transit scores. If the prefix is anycasted, however, it contributes only to global scores.
* **Transit score**: Depending on the AS hop count from the prefix's origin in the AS path, prefix scores contribute to the following types of transit scores:

  + Retail score of the origin AS as well as any single-homed upstream
  + Wholesale score of the next transit AS after the origin or single-homed upstream
  + Backbone score of the next transit AS after the wholesale AS
  + Customer base score of any AS classified as a transit provider regardless of location in the AS path

    ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(250).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A05Z&se=2026-05-12T09%3A58%3A05Z&sr=c&sp=r&sig=F%2BlbDBwGl2FLgte%2B9sjW2Wk93Ts1ccsULONyd6dlzT0%3D)

    *How the elements of a simplified IPv4 RIB entry, shown at top against a light gray background, contribute to the score calculation described in steps 2 and 3.*

#### Step 4. Rank ASes by Transit Score

Once the geolocation and transit score contribution to each AS has been calculated for every prefix in the global routing table, Kentik has a complete set of transit scores by geolocation and score type, which means that every AS that is either originating a prefix or transiting one will have some type of transit score for at least one geography. KMI can then rank ASes by transit scores of each type for every geolocation.

#### Step 5. Rank ASes by Peering Score

Peer rankings apply to ASes that have peering relationships. As with transit rankings, these rankings are based on observing IP address space. The greater the amount of address space sent over a settlement-free peering session, the greater the potential for traffic to be “peered off” or sent cost-free, avoiding transit costs. So generally speaking, the more address space sent to an AS via peering, the higher the AS's peering rank will be.  
  
ASes with peering relationships can also be ranked within a specified geography based on their overall peering base score. The peering base score for an AS is the sum of the scores for each of its peer ASes.

---

© 2014-25 Kentik
