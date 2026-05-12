# Using Interface Classification

Source: https://kb.kentik.com/docs/using-interface-classification

---

This article discusses the practical usage of interface classification, illustrated with specific use cases.

> **Notes:**
>
> * For an overall description of interface classification, see **Classification Overview**.
> * For a reference to the **Interface Classification** module in the Kentik portal, see **Interface Classification**.
> * Kentik doesn't aggregate the network flow records stored in KDE, which means that the same flow could potentially be counted multiple times as it transits different devices in your network. For a discussion of how the network boundary dimension is used to avoid double-counting, see **Network Boundary Attribute**.

## Interface Classification Dimensions

Interface classification enables an organization to quickly and efficiently assign both a Connectivity Type and Network Boundary value to every interface in the network. As flow records are ingested into the Kentik Data Engine (KDE) they are augmented with the connectivity and boundary values for the interfaces that the flows have used. Incorporating these values into the flow records makes it possible for them to be represented as dimensions in queries, which enables you to create queries that segment traffic based on specific interface characteristics.  
  
The easiest place to see interface classification in action is the Kentik portal’s **Data Explorer** module. Interface classifications may be accessed here in two forms:

* **Group-by dimensions** (see **Dimension Selectors**): Set the combination of fields that define a set of traffic that can be counted (by metric) and ranked.
* **Filters** (see **Filtering Settings**) Include or exclude traffic records that contain a specified value in the field that you’re filtering on.

The following table shows the columns in which the interface classification values are stored in the records of the KDE and how they are referred to in the Data Explorer UI.

| Portal dimensions and filters | KDE field name |
| --- | --- |
| SOURCE: Connectivity Type | `i_src_connect_type_name` |
| SOURCE: Network Boundary | `i_src_network_bndry_name` |
| DESTINATION: Connectivity Type | `i_dst_connect_type_name` |
| DESTINATION: Network Boundary | `i_dst_network_bndry_name` |

In the next sections we’ll look at how these new dimensions can be easily applied in some very common use cases. Without them, performing the same tasks would require devoting significantly more effort to the construction of queries.

## Internal and External Interfaces

Commercial-scale networks typically deploy interfaces in two functional categories:

* **External**: "Edge" interfaces that link to interfaces in a different Autonomous System (AS).
* **Internal**: Interfaces that link to other interfaces in the same AS.

The Network Boundary dimension, assigned to interfaces during Interface Classification, is what allows you to distinguish between traffic that is moving around inside your network and traffic that is entering or exiting through an edge interface. This allows you to narrow queries to edge traffic of one or both of the following kinds:

* **Inbound**: Flows where the value of Source Network Boundary is “External.” These represent traffic entering your network from elsewhere.
* **Outbound**: Flows where the value of Destination Network Boundary is “External.” these represent traffic exiting your network to elsewhere.

> **Note**: Correct network boundary classification is essential for the functioning of many Kentik modules in areas such as Core, Edge, Protect, and Service Provider.

## Connectivity by Country

One common task for network operators is to regularly review global connectivity, often evaluating it on a country-by-country basis. To intelligently price content destined for a given country, for example, content providers need a pretty good idea how much it actually costs them to deliver traffic to that locale. With costs of bandwidth varying from PNI (Private Network Interconnect) to Transit, it’s typically useful to look at segmentation for traffic — either inbound or outbound, depending on whether you are a content provider, carrier, or ISP — to each country of interest.  
  
 To set up for this type of analysis we would use the following settings in the Query sidebar of Data Explorer:

* **Visualization Type**: Pie Chart (gives us a quick visual overview)
* **Data Sources**: All Routers (under Types). We don’t need to be more specific because the Network Boundary filter will ensure that we only query on traffic as it exits the network.
* **Group-by Dimensions**: [Destination] Connectivity Type
* **Filtering** (ANDed):

  + Destination Network Boundary = External
  + Destination Country = FR (France for this example)

    ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(204).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A14Z&se=2026-05-12T09%3A39%3A14Z&sr=c&sp=r&sig=qVIscOA9%2FVUaHXTHaAhqMci7K9c0Qwtt2MKcB5i6k%2Bs%3D)

With these settings, we can see in the result shown above that a lot of the French traffic goes through transit, which is the costliest form of connectivity. We can go deeper on this first pass of investigation by figuring out which PoPs this traffic is delivered by, and, even more precisely, which devices and towards which ISP networks.  
  
To do this, the Data Explorer settings would be as follows:

* **Visualization Type**: Sankey
* **Data Sources**: All routers
* **Group-by Dimensions**:

  + [Non-Directional/Other] Site (under Network & Traffic Topology)
  + [Destination] Connectivity Type
  + [Non-Directional/Other] Device (under Network & Traffic Topology)
  + [Destination] Next Hop AS Number
  + [Destination] AS Number
* **Metrics** (Customize Metrics):

  + Primary Display & Sort Metric = 95th percentile (smooths out potential spikes)
* **Filtering** (ANDed):

  + Destination Network Boundary = External
  + Destination Country = FR (France for this example)

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(205).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A14Z&se=2026-05-12T09%3A39%3A14Z&sr=c&sp=r&sig=qVIscOA9%2FVUaHXTHaAhqMci7K9c0Qwtt2MKcB5i6k%2Bs%3D)

The resulting diagram shows where (from which of your devices) the traffic with a given destination ASN is leaving your network, as well as the 1st hop ASN for that traffic. By combining that information with your cost information for each of those ASes, you can now develop an overview of your outbound traffic cost structure in the selected country.

## Optimizing by Country

A lot of peering review meetings basically consist of reviewing connectivity in a given market and trying to see what can be optimized for cost and/or performance. Either way, it’s always a good practice to try to shift as much traffic as possible from transit to any form of peering (ideally free PNI, but possibly public peering in an IX fabric). Figuring out how to do that is effectively a two-step process.

#### Step One

We begin by selecting the destination country. How you do this depends on whether you already know what destination country you want to study or whether you want to select that country based on how much transit connectivity it requires. For the latter analysis in Data Explorer, we would use the following settings:

* **Visualization Type**: Stacked Area Chart
* **Data Sources**: All Routers
* **Group-by Dimension**: [Destination] Country
* **Filtering** (ANDed):

  + Destination Network Boundary = External
  + Destination Connectivity Type = Transit

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(206).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A14Z&se=2026-05-12T09%3A39%3A14Z&sr=c&sp=r&sig=qVIscOA9%2FVUaHXTHaAhqMci7K9c0Qwtt2MKcB5i6k%2Bs%3D)

As shown above, we now have a list of destination countries ranked by transit connectivity (at the edge of your network). This tells us where the greatest potential is for saving money by shifting from transit to peering. Let’s focus on Germany (DE) for our next step.

#### Step Two

1. In the table under the graph, click on the **Options** icon (hamburger) at the right of the German (DE) row.![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(207).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A14Z&se=2026-05-12T09%3A39%3A14Z&sr=c&sp=r&sig=qVIscOA9%2FVUaHXTHaAhqMci7K9c0Qwtt2MKcB5i6k%2Bs%3D)
2. In the resulting popup menu, choose **Show By**, which opens the Show By Dimensions dialog.
3. Select **Destination****»** **AS Number**, then click the **Show By Selected Dimensions** button.

After applying the changes, we have an ordered list of Germany’s top ASNs that are currently reached via transit, a.k.a. a list of the traffic destinations with which it would be most beneficial for us to peer. This is valuable business intelligence that it would be far more difficult to obtain without interface classification dimensions.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(208).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A14Z&se=2026-05-12T09%3A39%3A14Z&sr=c&sp=r&sig=qVIscOA9%2FVUaHXTHaAhqMci7K9c0Qwtt2MKcB5i6k%2Bs%3D)

## Adding Ultimate Exit

By now we have a good idea of the power that interface classification (IC) gives us in terms of extracting business value from network data. We can take this even further by marrying IC with another feature of Kentik: Ultimate Exit. As a quick refresher (for more, see **Using Ultimate Exit**), Ultimate Exit augments flow records in KDE with information about the site, device, and interface of the exit point of the underlying flows from which the records are derived.  
  
As shown in the Sankey diagram below, using this Ultimate Exit information allows you to slice traffic by the sites/devices/interfaces toward which it is headed so that you can see how much traffic is flowing between an entry site and an ultimate exit site/device/interface (specified with **Ultimate Exit Dimensions**). You can also see which eyeball networks (the destination ASN) are reached by that traffic, and you can export that information as a CSV file for business analytics (from the **Actions** menu at the top right of the page, choose **Export** **»** **Data Table**). The business utility of this capability is to enable you to look at where transport is the most expensive between your entry site and ultimate exit site so that you can take the appropriate steps to optimize.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(209).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A14Z&se=2026-05-12T09%3A39%3A14Z&sr=c&sp=r&sig=qVIscOA9%2FVUaHXTHaAhqMci7K9c0Qwtt2MKcB5i6k%2Bs%3D)

To understand the benefit of combining Interface Classification with Ultimate Exit, let’s assume that your cost per Mb/s depends on connectivity type, and that the greater the cost, the lower your ROI. As shown in the diagram below, the resulting ROI hierarchy would be as follows (listed from low ROI [1] to high ROI [5]):

1. Transit (ROI = $)
2. Paid peering (ROI = $$)
3. IX peering (ROI = $$$)
4. Routing between devices across your network (ROI = $$$$)
5. Routing within a single device, a.k.a. hot-potato-routing (ROI = $$$$$).

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(210).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A14Z&se=2026-05-12T09%3A39%3A14Z&sr=c&sp=r&sig=qVIscOA9%2FVUaHXTHaAhqMci7K9c0Qwtt2MKcB5i6k%2Bs%3D)

To maximize ROI, you need network business intelligence that tells you where the traffic you carry for a given customer falls into this hierarchy. You can get that information in Data Explorer using two specific dimensions: BGP Ultimate Exit Connectivity Type and BGP Ultimate Exit Network Boundary. Here’s how to set up such a query:

* **Data Sources**: All Routers
* **Display Type**: Sankey
* **Dimensions** (described in **Ultimate Exit Dimensions**):

  + [Source] AS Number
  + [Non-Directional/Other] Site(under Network & Traffic Topology)
  + [Non-Directional/Other] BGP Ultimate Exit Site
  + [Non-Directional/Other] BGP Ultimate Exit Connectivity Type
  + [Destination] AS Number (optional)
* **Filtering**:

  + [Source] Interface Description = "Customer AS1234"

    > **Note**: Assuming that we've included the customer’s name in the description of the interface serving that customer, this filters the results to just those coming from that customer.

> **Note**: Alternatively, you could use the Interface\_Provider attribute from Interface Classification (IC) to set Customer names automatically based on interface descriptions. The filtering to use would then be less brittle: [Source] Interface Provider = `customer_name`.

This query results in a Sankey with some very valuable ROI information related to the traffic coming from Customer AS1234. Focusing on the central part of the diagram, the annotations in the image below use color to show three different types of paths taken by the traffic of the customer in this example, with each representing different levels of ROI:

* Traffic on the pink path (near the bottom of the diagram) is the most profitable. This traffic is routed within the same site (Washington to Washington), so it has neither a metro nor long-haul cost. And since it’s delivered to another customer (Ultimate Exit Connectivity Type is “Customer”), it can be billed to that other customer.
* Traffic on the orange path (top of the diagram) is second in terms of ROI. It’s hot-potato-routed, so it has no transport costs, and it’s handed over via free PNI (Ultimate Exit Connectivity Type is “Free PNI”), so there’s no cost to deliver it. But it does not actually generate revenue like traffic on the pink path.
* Lastly, traffic on the blue path has the lowest ROI. It enters in San Francisco and is delivered in Sao Paulo, Brazil, which involves costly transport between two countries. This cost is probably not offset by the fact that it is delivered to a Free PNI.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(211).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A14Z&se=2026-05-12T09%3A39%3A14Z&sr=c&sp=r&sig=qVIscOA9%2FVUaHXTHaAhqMci7K9c0Qwtt2MKcB5i6k%2Bs%3D)

The fact that Kentik enables us to see different paths like the ones that we’ve colored in the diagram above illustrates that we can use it to precisely establish the different types of traffic that can be compounded into an ROI — weighted by traffic volume — for each customer. A ballpark ROI analysis can be conducted by viewing a Sankey generated from the settings shown above, which can be downloaded as an image to share during a traffic engineering or sales meeting. The underlying data can also be exported from Kentik for further analysis in a financial or BI environment, either as CSV via the portal (see **Export Chart or Table**) or as JSON via our Query API (see **Query Data Method)**.

#### Ultimate Exit Dimensions

The following dimensions are used in the query settings described in **Adding Ultimate Exit**:

* **[Source] AS Number**: The customer ASN, which helps ensure that we’re catching the right ports.
* **[Non-Directional/Other] Site** (under Network & Traffic Topology): The sites where we connect to our customer.
* **[Non-Directional/Other] BGP Ultimate Exit Site**: The site where our customer’s traffic will be handed off to an adjacent network, which enables us to estimate how much transport is involved to deliver it.
* **[Non-Directional/Other] BGP Ultimate Exit Connectivity Type**: The type of connectivity over which this customer’s traffic will be handed off, which can help us calculate the cost of transport between the customer’s traffic entry and exit sites.
* **[Destination] AS Number** (optional): the ASN of the destination of the customer's traffic.

## Taking IC Further

The examples presented in the previous sections can be enhanced by leveraging additional functionalities of Kentik's **Dashboards**.

#### Guided Mode Dashboards

Guided dashboards enable users to specify (without getting deep into filtering settings) the value of a dimension that is applied as a filter to some or all of the dashboard’s panels (see **Guided Mode Pane**). To investigate country-level connectivity you can create a guided dashboard whose guided-mode dimension is a country.

#### Cross-panel Filtering

Another feature of dashboards is the Cross-panel Filtering switch in the settings for individual panels. When enabled, this results in the application of key-based filtering (see **Filtering With the Key**) to all panels on the dashboard. Using the example in **Connectivity by Country**, this capability could be used to filter all of a dashboard's panels to show data for whichever country is clicked on in the pie chart. In the image below, for example, the Sankey diagram is filtered because Canada (CA) was clicked in the chart.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(213).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A14Z&se=2026-05-12T09%3A39%3A14Z&sr=c&sp=r&sig=qVIscOA9%2FVUaHXTHaAhqMci7K9c0Qwtt2MKcB5i6k%2Bs%3D)
