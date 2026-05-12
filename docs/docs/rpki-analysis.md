# RPKI Analysis

Source: https://kb.kentik.com/docs/rpki-analysis

---

This articles discusses the **RPKI Analysis** dashboard in the Kentik portal.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(386).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A17Z&se=2026-05-12T09%3A36%3A17Z&sr=c&sp=r&sig=gqZ4tKoNhEai8dc4K1yS9TCqcjesSmIJ9ciYMv4mFXc%3D)

*Use RPKI Analysis to check the validity of routes announced by Autonomous Systems.*

## About RPKI Analysis

Resource Public Key Infrastructure (RPKI) is a security framework used to authenticate BGP routing information. RPKI allows the BGP route prefixes announced by an Autonomous System (AS) to be cryptographically signed such that ISPs and other entities can validate that the AS is authorized to announce said routes.  
  
Kentik enriches flow records with their RPKI validation state as they are ingested into the Kentik Data Engine (KDE), putting some powerful analysis tools at your fingertips. One such tool is the **RPKI Analysis** page, under **Protect** on the portal's main menu. Using this page, you can identify sites (see **About Sites**) on your network with traffic that is invalid and/or will be dropped if strict RPKI validation is enforced on the routers.

> **Note:** For more about Kentik's RPKI implementation, see **Using RPKI**.

## RPKI Analysis Page

The RPKI Analysis page is a dashboard of pre-configured visualization panels that show which sites and devices (if any) on your network have had traffic with invalid RPKI status. You can use this page to improve routing security mechanisms and to check traffic that would be dropped if strict RPKI validation was enabled.  
  
By default, the visualizations on this page show invalid traffic over the last day, but you can customize them using the controls in the panes of the **Query** sidebar in the **RPKI Analysis Subnav**.

> **Note**: A site must be specified in the **Query** sidebar's **Select a Site** pane to see any data on this page, but only the "`Site:` detailed analysis" panel at the bottom of the page is directly affected by site selection. The other panels, whose titles are prefixed with ‘[GLOBAL]’, cover all sites/devices.

## RPKI Analysis Page UI

The UI elements of the RPKI Analysis page are based on the standard Kentik **Dashboards** that live in the portal's **Library**. The topics below cover variances between standard dashboards and the RPKI Analysis page.

### RPKI Analysis Subnav

The subnav (silver strip across top of page) contains page-wide controls that are, except for the following buttons, identical to the portal's **Dashboard Subnav Controls**:

* **Share**: A button that opens the Share dialog, enabling you to share the contents of page (see **Portal View Sharing**). For RPKI Analysis, the only **Report File** Type available in the **Email** and **Subscriptions** tabs is PDF.

  > **Notes**:
  >
  > + To download the entire page rather than share it, use the Export option from the Actions drop-down.
  > + To export an individual visualization, select Export from that panel's context menu (see **RPKI Panel Context Menu**).
* **Actions**: A button that opens a drop-down with multiple options:

  + Clone: Opens the Clone Dashboard dialog (see **Dashboard Navigation**).
  + Export: Prepares a visual report (PDF) including all panels of the dashboard. A notification appears when the PDF is ready to download.
  + Subscribe: Opens the Subscribe dialog, enabling you to create a subscription for this page (see **Subscription Tab UI**).
  + Unsubscribe (visible only when subscribed to this page): Opens a dialog enabling you to confirm that you wish to be removed from the subscription.
  + Preview as Tenant: Opens a submenu from which you can choose a MKP tenant, after which you'll be taken to the page as it would appear to that tenant.

### RPKI Analysis Title Bar

The page title bar, located directly under the subnav, includes the following elements:

* **Favorite** (star): An icon indicating whether this dashboard has been designated as a favorite dashboard. Click the star to designate the page as a favorite (see **Library Bookmarks**).
* **Title**: RPKI Analysis (the name of the page).
* **View Description** (comment icon): A button that opens a Dashboard Description dialog containing a short description of the RPKI Analysis page.

## RPKI Analysis Panels

The visualizations and corresponding panels of the RPKI Analysis page are covered in the following topics.

### RPKI Analysis Panel UI

Beneath the RPKI Analysis page title bar is the page's main display are, which contains a set of panels that each feature one visualization. These panels include the common UI elements covered in the topics below.

#### RPKI Panel Title

The title is located at the upper left of each panel. When clicked, the title opens the visualization in Data Explorer, where you can customize it using the full suite of **Query Sidebar Controls**.

#### RPKI Panel Context Menu

The context menu is accessed via the vertical ellipsis button found at the upper right of each panel. The menu includes the following options:

* **Change Visualization**: A submenu from which you can select the view type (see **Chart Visualization Types**).
* **Export**: A submenu from which you can choose a format in which to export the panel's visualization (see **Portal Export Options**).

### RPKI Dropped Traffic Panel

The default visualization in this panel is a gauge predicting the max bitrate of invalid traffic, meaning the traffic that will be dropped if strict RPKI validation is enabled on the routers of the user's organization. The prediction is based on routing information and RPKI ROAs during the lookback time range selected in the Query sidebar (e.g., last 1 day), and applies to an equal time duration going forward from the present.  
  
The gauge is bracketed (see **About Bracketing**), meaning that it's color will indicate the bracket range in which the traffic rate falls. If no invalid traffic exists for any site, "No Results" will be shown instead of the gauge.

### RPKI Invalids Panel

The default visualization in this panel is a Sankey diagram that visually traces the max invalid traffic for all sites over the lookback time range selected in the **Query** sidebar (e.g., last 1 day). The diagram's columns each represent a dimension of the traffic, as follows (left to right):

* **RPKI Quick Status**: The RPKI status assigned by Kentik to the traffic, grouped for the purposes of this visualization into the following buckets:

  + RPKI Invalid - covering Valid/Unknown;
  + RPKI Invalid - Will be dropped.
* **Site**: The name of the site associated with the invalid traffic.
* **Device**: The name of the device through which the traffic was routed.
* **Dest Connectivity Type**: The name of the type of connectivity between the router and the next hop router, e.g., Free Private Peering.
* **Dest Next Hop AS Number**: The ASN of the next hop router.
* **Dest Provider**:The name of the provider via which traffic reaches the Internet, e.g., Akamai.

When multiple values exist for a given dimension, the corresponding column includes a separate color-coded segment for each value. Hovering over a given segment opens a popup showing the corresponding dimension name, value, and traffic bitrate.

If no invalid traffic exists for any site, the text ‘No Results’ will be shown instead of the diagram.

### RPKI Detailed Analysis Panel

The table in this panel lists the details of the invalid traffic data for the site selected in the Query sidebar.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(389).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A17Z&se=2026-05-12T09%3A36%3A17Z&sr=c&sp=r&sig=gqZ4tKoNhEai8dc4K1yS9TCqcjesSmIJ9ciYMv4mFXc%3D)

*The detailed analysis table shows details of invalid traffic through one site.*

The table includes the following columns:

* **Destination Route Prefix/LEN**: The prefix of the destination block of IP addresses followed by the prefix length in bits, e.g., 190.0.0.0/18. Click the link to go to the Core details page for this route prefix.
* **Destination RPKI Validation Status**: The RPKI status assigned by Kentik to the traffic, e.g., "RPKI Invalid: prefix length out of bounds."

  > **Note:** These statuses are more detailed than the "RPKI Quick Status" shown in the RPKI Invalids Sankey.
* **Destination Next Hop AS Number**: The name associated with the destination next hop ASN followed by a lozenge of the ASN. Click the link to go to the Core details page for this ASN. The lozenge shows the ASN and has the following behaviors:

  + On hover: Open a tooltip that shows the name, country code of registration, and ASN.
  + On click: Open a popup showing the registration country name, along with links to these portal pages for the ASN: Kentik Market Intelligence (see **AS Details Page**), PeeringDB Record (see **PeeringDB Info Tab**), and ASN Traffic View (the Core details page for this ASN).
* **Destination AS Path**: Lists the ASNs this traffic traveled through to reach the destination device, separated by spaces. The first ASN is for the next hop router, followed by the 2nd hop and 3rd hop ASNs, as applicable. Click the link to go to the Core details page for this AS Path.
* **As Path Breakdown** (bullet list icon): A button with the following behaviors:

  + On hover: Open a popup that shows the destination AS Path.
  + On click: Open the **As Path Breakdown** popup, which shows a breakdown of the AS Path, i.e. Next Hop, 2nd Hop, 3rd Hop.
* **Device**: The name of the externally facing device through which the traffic was routed. Click the link to go to the Core details page for this Device.
* **Max Mbits/s**: The max rate of invalid traffic observed by Kentik over the last day in megabits per second.
* **Last Datapoint Mbits/s**: The most recent rate of invalid traffic sampled by Kentik over the last day in megabits per second.

The last row of the table is a summary with totals for the Max Mbits/s and Last Datapoint Mbits/s columns. If no invalid traffic exists for the selected site, the text "No Results" will be shown instead of the table.

> **Notes:**
>
> * Clicking a row (outside of any links) will select that row but will not change the data in the page's panels as it often does on other Kentik dashboards.
> * For information on Core details pages, see **Core Details Pages**.
