# Network Classification

Source: https://kb.kentik.com/docs/network-classification

---

This article discusses **Network Classification** page in the Kentik portal.

> **Note:** If your Kentik-monitored traffic is exclusively to/from one or more cloud providers, you need not fill in the fields of the Network Classificationpage (Kentik classifies the flow logs for cloud traffic automatically at ingest).

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(629).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A32Z&se=2026-05-12T09%3A36%3A32Z&sr=c&sp=r&sig=T6wexZ41CyK5ZbW6vd5R7rewSXd4V5p1tF%2BcknB5AAk%3D)

*Network classification settings enable Kentik to determine the direction of traffic with respect to your network.*

## About Network Classification

Network classification uses source and destination information related to IPs and ASes to determine the direction of network traffic with respect to your network. The following types of classification are supported:

* **Network Directionality**: Used to look for traffic based on the direction from which it enters the network and to which it leaves.
* **Host Directionality**: Used to look for host traffic captured by Kentik’s software agent based on the direction it is flowing.

### Network Classification Dimensions

The two categories of Network Classification directionality listed above are supported by several **Network Classification Dimensions** that can be applied to each flow as it is ingested into the main tables of KDE:

* **Traffic Origination:** This dimension indicates whether the source for a given flow is inside or outside of your network.
* **Traffic Termination:** This dimension indicates whether the destination for a given flow record is inside or outside of the network.
* **Host Direction:** When the flow record has been generated on a host, this dimension indicates whether the direction of traffic is into or out of that host.
* **Traffic Profile**: Derived from Traffic Origination and Traffic Termination, this dimension categorizes traffic into the directionalities described in **Traffic Profile**.
* **Simple Traffic Profile**: A less granular version of the traffic profile, primarily used in Kentik v4 (e.g., Network Explorer, Insights & Alerting), with the possible values listed in **Simple Traffic Profile**.

The dimensions described above are available throughout Kentik as:

* Group-by Dimensions
* Filter match criteria
* Alert keys (dimensions of an Alert Policy)
* Alert filter match criteria

> **Note**: The use of Network Classification in Alerting enables you to monitor traffic that comes from outside the network separately from traffic that is internal to the network.

#### Traffic Profile

This dimension categorizes traffic into the following specific directionalities:

* **Internal**: Traffic originating and terminating inside the network.
* **Through**: Traffic coming from outside the network and terminating outside the network.
* **From outside, terminated inside** (ingress): Traffic coming from outside the network and terminating inside the network.
* **Originated inside, to outside** (egress): Traffic coming from inside the network and terminating outside the network.
* **Internal Cloud**: Traffic originating and terminating inside the same public cloud.
* **From outside, to cloud**: Traffic coming from outside the network and terminating in one of your organization's resources in a public cloud.
* **From cloud to outside**: Traffic coming from one of your organization's resources in a public cloud and terminating outside your network.
* **From cloud to inside**: Traffic coming from one of your organization's resources in a public cloud and terminating inside your network.
* **From inside to cloud**: Traffic coming from inside your network and terminating in one of your organization's resources in a public cloud.
* **Multi-cloud**: Traffic coming from one of your organization's resources in a public cloud and terminating in one of your organization's resources in a different public cloud.

#### Simple Traffic Profile

This dimension, primarily used in the Kentik portal (e.g., Network Explorer, Insights & Alerting, etc.) categorizes traffic into the following general directionalities, which are illustrated in the diagram below:

* **Inbound**: Traffic that originates outside your network and terminates inside your network.
* **Outbound**: Traffic that originates inside your network and terminates outside your network.
* **Internal**: Traffic that both originates and terminates inside your network.
* **Through**: Traffic that both originates and terminates outside your network.
* **Other**: Traffic that is not classified by the above profiles (see **Traffic Classified as Other**).

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(630).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A32Z&se=2026-05-12T09%3A36%3A32Z&sr=c&sp=r&sig=T6wexZ41CyK5ZbW6vd5R7rewSXd4V5p1tF%2BcknB5AAk%3D)

The key to accurately assigning a Simple Traffic Profile value is to thoroughly apply both Network Classification and **Interface Classification**. The table below shows how the values of other dimensions in the same flow record are evaluated to determine the Simple Traffic Profile value for a given flow.

| If... | ...and... | ...then Simple Traffic Profile is: |
| --- | --- | --- |
| Traffic Profile is not Internal Cloud | * Network Boundary is External; and * Ultimate Exit Network Boundary is Internal | Inbound |
| Traffic Profile is not Internal Cloud | * Network Boundary is External; and * Ultimate Exit Network Boundary is None; and * Traffic Termination is Inside | Inbound |
| Traffic Profile is not Internal Cloud | * Traffic Termination is AWS; and * Cloud Provider is not null | Inbound |
| Traffic Profile is not Internal Cloud | * Traffic Termination is Azure; and * Cloud Provider is not null | Inbound |
| Traffic Profile is not Internal Cloud | * Traffic Termination is OCI; and * Cloud Provider is not null | Inbound |
| Traffic Profile is not Internal Cloud | * Traffic Termination is GCP; and * Cloud Provider is not null | Inbound |
| Traffic Profile is not Internal Cloud | * Traffic Origination is Inside; and * Destination Network Boundary is External | Outbound |
| Traffic Profile is not Internal Cloud | * Traffic Origination is AWS; and * Cloud Provider is not null | Outbound |
| Traffic Profile is not Internal Cloud | * Traffic Origination is Azure; and * Cloud Provider is not null | Outbound |
| Traffic Profile is not Internal Cloud | * Traffic Origination is OCI; and * Cloud Provider is not null | Outbound |
| Traffic Profile is not Internal Cloud | * Traffic Origination is GCP; and * Cloud Provider is not null | Outbound |
| Network Boundary is External | Ultimate Exit Network Boundary is External | Through |
| Traffic Termination is Outside | * Network Boundary is External; and * Ultimate Exit Network Boundary is None | Through |
| Traffic Profile is Internal | * Network Boundary is not External; and * Destination Network Boundary is not External | Internal |
| Traffic Profile is Internal Cloud | * Network Boundary is not External; and * Destination Network Boundary is not External | Internal |
| None of the above | None of the above | Other (see **Traffic Classified as Other**) |

#### Traffic Classified as Other

Ideally most traffic flows are assigned a Simple Traffic Profile of Inbound, Outbound, Internal, or Through. The most common reasons for traffic to be assigned a profile of Other are as follows:

* **Multiple internal hops**: If the path of Inbound, Outbound, or Through traffic includes multiple hops within the customer’s network. This may be backbone traffic (Simple Traffic Profile does not currently include a special designation for backbone).
* **Interface classification is incomplete or incorrect**: If **Interface Classification** hasn't yet been run or the Network Boundary values assigned by a given rule are incorrect for the matching interfaces.

  + **Network classification is incomplete or incorrect**: If internal IPs and ASNs haven't yet been specified on the **Network Classification Page** or have been entered incorrectly.
* **Classification mismatch**: If the settings for Interface Classification and Network Classification are inconsistent with each other, resulting in contradictory classifications.

### Network Directionality Use Case

One application of Network Classification is to use the Network Directionality dimensions to investigate spikes in traffic. Suppose, for example, that we used Data Explorer to run a query for top-X customers, and the resulting graph revealed big spikes in flows to a customer called Pear, Inc. (as shown below).

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(631).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A32Z&se=2026-05-12T09%3A36%3A32Z&sr=c&sp=r&sig=T6wexZ41CyK5ZbW6vd5R7rewSXd4V5p1tF%2BcknB5AAk%3D)

*Two spikes in traffic to the Customer Pear, Inc.*

To dig deeper into this anomaly, start in the table beneath the graph in the Explorer display area.

1. Click the **Action** menu at the right of the row corresponding to Pear, Inc., and choose **Show By** to open the Show By Dimensions dialog.
2. Choose one of our Network Classification dimensions, Traffic Origination (listed under Source).
3. Close the dialog by clicking the **Show By Selected Dimensions** button, and re-run the query.

In the resulting graph (below), the spike is made up of traffic that originated outside of our network (having a Traffic Origination value of “outside”). To continue digging further, use **Show By** again, this time looking at source ASN or IP address.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(632).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A32Z&se=2026-05-12T09%3A36%3A32Z&sr=c&sp=r&sig=T6wexZ41CyK5ZbW6vd5R7rewSXd4V5p1tF%2BcknB5AAk%3D)

*Traffic to Pear, Inc. ranked by origination (inside vs. outside).*

### Host Directionality Use Case

Another use case for Network Classification is specific to host traffic captured by Kentik’s software host agent. Since most hosts have only a single interface through which traffic can pass, Kentik captures both inbound and outbound traffic. Host directionality enables you to separate traffic that was coming in from traffic that was leaving. To try this:

1. Set the **Data Sources** pane of the Data Explorer sidebar to include the hosts that you want to check.
2. Run a query with Host Direction (in the **Non-Directional** category) as the group-by dimension.

As shown in the graph below, you can now see separate flows in and out of your hosts.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(633).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A32Z&se=2026-05-12T09%3A36%3A32Z&sr=c&sp=r&sig=T6wexZ41CyK5ZbW6vd5R7rewSXd4V5p1tF%2BcknB5AAk%3D)

*See the proportion of host traffic by direction (in vs. out).*

## Network Classification Page

Before using Network Classification you must first enable Kentik Detect to determine what is inside and what is outside of your network. Network classification is configured on the Network Classification page, which is accessed by selecting **Settings** from the main menu, and then clicking **Network Classification** in the Network Metadata section.

> **Note:** If your Kentik-monitored traffic is exclusively to/from one or more cloud providers, you need not fill in the fields of theNetwork Classification page (Kentik classifies the flow logs for cloud traffic automatically at ingest).

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(634).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A32Z&se=2026-05-12T09%3A36%3A32Z&sr=c&sp=r&sig=T6wexZ41CyK5ZbW6vd5R7rewSXd4V5p1tF%2BcknB5AAk%3D)

* **Internal IPs**: An input field used to enter, with a comma-separated list, the IP CIDR blocks used inside the network.
* **Automatically include RFC1918 IP space**: A checkbox used to specify whether the RFC1918 IP space and other common private ranges are included along with the user-defined Internal IPs list. This option is checked by default.
* **Internal ASNs field**: An input field used to enter, with a comma-separated list, the ASNs used inside the network.
* **Automatically include private ASNs**: A checkbox used to specify whether private 16- and 32-bit ASNs are included along with the user-defined Internal ASNs list. This option is checked by default.
* **Save button**: Click to save any changes that have been made since arriving on the page.

## Configuring Network Classification

To set up network classification:

1. Go to the Network Classification page (Settings » **Network Classification**).
2. In the **Internal IPs** field, enter a comma-separated list of the IPs that are internal to your network.
3. Using the checkbox below the field, choose whether to automatically include the RFC1918 IP space.
4. In the **Internal ASNs** field, enter a comma-separated list of the ASNs that are internal to your network.
5. Using the checkbox below the field, choose whether to automatically include private ASNs.
6. Click the **Save** button. You are now ready to start using network classification dimensions for group-by and filtering in queries and alerting.

---

© 2014-25 Kentik
