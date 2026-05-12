# BGP Route Viewer

Source: https://kb.kentik.com/docs/bgp-route-viewer

---

This article covers the BGP Route Viewer dashboard in the Kentik portal.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(497).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A03Z&se=2026-05-12T09%3A33%3A03Z&sr=c&sp=r&sig=Hr546AcbEY4uSEnjcxQiQEHqqkdTQRwSswnPo1V4V68%3D)

*BGP route announcements and withdrawals are shown in the timeline and table.*

## About BGP Route Viewer

The **BGP Route Viewer** tab displays BGP events involving one or more Autonomous System Numbers (ASNs) or IP prefixes that your organization has chosen to monitor, enabling you to correlate issues revealed by synthetic testing at the network (ping, trace) and web layers (HTTP, Page Load) with issues at the routing layer (BGP). An event can be either a route announcement or a route withdrawal.  
  
The prefixes and ASNs to monitor are by default the same as those identified as "internal" in your organization's **Network Classification** settings, but the defaults can be modified in the **BGP Route Viewer** tab's View/Update Prefixes and ASN(s) dialog (see **Prefixes and ASNs**). The route announcements and withdrawals involving these monitored IPs and ASNs within the last 12 hours will be shown in a timeline chart (see **BGP Events Timeline**) and a table (see **BGP Events List**).

#### BGP Event Datasets

The monitoring of events involves thousands of BGP vantage points located in various parts of the Internet. This broad coverage is made possible by querying the following datasets:

* **Route Views project**: Publicly available BGP routing data maintained by the University of Oregon’s Advanced Network Technology Center (see **www.routeviews.org**).
* **Kentik private peers**: Curated from the thousands of private BGP peering sessions that Kentik has established with customer devices located around the world.

  > **Note:** In this dataset the left two AS hops for events, which indicate the ASN of the private peer, are obfuscated to protect customer information.

## BGP Route Viewer UI

The BGP Route Viewer tab includes the following main UI elements:

* **View/Update Prefixes and ASN(s)**: A button that opens the View/Update Prefixes and ASN(s) dialog (see **Prefixes and ASNs**).
* **Start Monitoring**: A button that takes you to a **Test Settings Page** for a new BGP Monitor test (see **Synthetics Test Types**), where the prefix and ASN fields will be populated to match the corresponding fields of the BGP Route Viewer's View/Update Prefixes and ASN(s) dialog.
* **BGP Events Timeline**: A vertical bar chart showing BGP events over the last 12 hours involving the ASNs and IP prefixes specified in the View/Update Prefixes and ASN(s) dialog (see **BGP Events Timeline**).
* **BGP Events List**: A table with details about the events shown in the events timeline (see **BGP Events List**).

## Prefixes and ASNs

The View/Update Prefixes and ASN(s) dialog is used to specify the prefixes and ASNs to monitor for route announcements and withdrawals. The dialog is opened by clicking the **View/Update Prefixes and ASN(s)** button.

The dialog includes the following settings and controls:

* ![Input fields for monitoring prefixes and ASN with a query button.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/BRV-view-update-prefixes-asns-dialog.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A03Z&se=2026-05-12T09%3A33%3A03Z&sr=c&sp=r&sig=Hr546AcbEY4uSEnjcxQiQEHqqkdTQRwSswnPo1V4V68%3D)**Prefixes to monitor**: A field into which you can enter a comma-separated list of prefixes in IP/CIDR notation. By default the list will be populated with the list in the **Internal IPs** field of your organization's **Network Classification** settings page.
* **ASN(s)**: A field into which you can enter a comma-separated list of ASNs. By default the list will be populated with the list in the **Internal ASNs** field in your organization's Network Classification settings.
* **Include covered prefixes**: A switch that, when on (default), includes in the timeline and table results events that are associated with the sub-prefixes of the query's current prefixes.
* **Run Query**: A button that reruns the query whose results are displayed in the timeline and table of the BGP Events pane. If the values in the **Prefixes to monitor or ASN(s)** fields differ from the internal IPs or ASNs specified in your Network Classification settings then the **Update Network Classification** dialog will open after the updated query results return.

#### Update Network Classification

When you click the **Run Query** button, Kentik evaluates whether the values in the **Prefixes to monitor or ASN(s)** fields exactly match the internal IPs or ASNs specified in your Network Classification settings. If not, the Update Network Classification dialog will open after the updated query results return, offering you the opportunity to use some or all of the non-matching IPs or ASNs for your organization's Network Classification.

The dialog includes the following fields and controls:

* ![Dialog box prompting to update network classification with IPs and ASNs.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/BRV-update-network-classification-dialog.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A03Z&se=2026-05-12T09%3A33%3A03Z&sr=c&sp=r&sig=Hr546AcbEY4uSEnjcxQiQEHqqkdTQRwSswnPo1V4V68%3D)**Close**: Click the **X** in the upper right corner to exit the dialog and keep the IPs and ASN settings the same as they were when the dialog was opened.
* **Internal IPs**: A field in which to enter a comma-separated list of IP addresses and/or subnets that will be added to the IPs classified as internal by **Network Classification**. By default the list will be populated with the IPs from the View/Update Prefixes and ASN(s) dialog that don't match your existing Network Classification IPs.
* **Internal ASN**s: A field in which to enter a comma-separated list of ASNs that will be added to the ASNs classified as internal by Network Classification. By default the list will be populated with the ASNs from the View/Update Prefixes and ASN(s) dialog that don't match your existing Network Classification ASNs.
* **Do not show me this dialog again**: Click to prevent Kentik from asking about Network Classification each time that you enter non-matching IPs or ASNs in the View/Update Prefixes and ASN(s) dialog.
* **Cancel**: A button that closes the dialog without saving changes.
* **Save**: A button that saves the current IPs and ASN settings and exits the dialog.

> **Notes:**
>
> * Kentik recommends backing up your existing Network Classification settings before making changes with the Update Network Classification dialog.
> * If you don't add the non-matching IPs/ASNs to your Network Classification settings then when you next visit the **BGP Route Viewer** tab the query will cover only the internal IPs/ASNs that are specified in the Network Classification settings.

## BGP Events Timeline

The events timeline is a vertical bar chart providing a visual summary of the BGP events observed over the most recent 12 hour period that are related to the ASNs and IP prefixes specified in the View/Update Prefixes and ASN(s) dialog (see **Prefixes and ASNs**). Each bar represents the events during a 10-minute period, with positive axis bars (green) representing route announcements and negative axis bars (purple) representing route withdrawals. The height of the bars represents the number of announcements or withdrawals during the period (the taller the bar the more events). Click on a bar to filter the **BGP Events List** to show only events that occurred during that 10-minute period (click on that bar again to restore the table to its default state).

## BGP Events List

The events table provides a detailed view of the information displayed in the events timeline. By default the table shows all events during the last 12 hours. To show only events during an individual 10-minute period, click on that period's bar in the **BGP Events Timeline** (click on that bar again to restore the table to its default state).  
  
 The table includes the following columns:

* **Type**: Indicates the type of this event, either Announcement or Withdrawal.
* **Prefix**: The prefix associated (announced or withdrawn) with the event.
* **Origin AS**: The name and number of the Autonomous System (AS) from which the event originated.
* **AS Path**: The path that the event information took in getting from the **origin AS** to the vantage point, indicated by the ASNs (space-separated) traversed hop-by-hop along the way.

  > **Note:** Because the origin AS indicated in the Origin AS column is the AS from which the event originated, it is found at the right of the AS path.
* **AS Path Length**: A count of the ASNs from and including the origin AS to the AS containing the vantage point.
* **Total**: The total number of events with the same type, prefix, and AS that were observed over the time period (the last 12 hours unless a 10-minute period is currently selected in the timeline).
* **RPKI Status**: One of the following values:

  + Valid: ROA (see **Route Origin Authorization**) is present, and the prefix length and ASN match.
  + Invalid: ROA is present, but the prefix length or AS number doesn't match.
  + Not-found or Unknown: ROA is not present.
* **Dataset**: The dataset from which the event was returned (see **BGP Event Datasets**):

  + Route Views: Indicated as “route-views."
  + Kentik private peers: Indicated as “kentik.”

> **Note:** A BGP event that was observed for both a prefix and an ASN will be shown only once in the results.

---

© 2014-25 Kentik
