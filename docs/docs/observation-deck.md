# Observation Deck

Source: https://kb.kentik.com/docs/observation-deck

---

This article discusses the **Observation Deck** in the Kentik portal.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(292).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A26Z&se=2026-05-12T09%3A34%3A26Z&sr=c&sp=r&sig=DUsXjPMrq%2BKp%2B9NBn68a2Il%2BA9cbPGFoT4yv13vIIgg%3D)

*The Observation Deck is a high-level view of your network that's customizable for individual users.*

## About Observation Deck

Observation Deck is a user-configurable high-level view of your organization's network. The core of Observation Deck is the display area where visualizations (a.k.a. "widgets") are displayed. Each visualization focuses on a specific aspect of your traffic volume, performance, costs, connectivity, or utilization. You can customize the display in terms of both content and layout to provide yourself with the information you most want to see when you come to the portal, and you can save your changes so the displayed information is specific to your individual needs. All changes are unique to you as a user (as distinct from the others in your organization).

## Observation Deck Title Bar

The title bar of Observation Deck runs across the top of the page and includes the following controls:

* **Add Visualizations**: A button that opens a drawer from the bottom of the screen from which you can choose additional visualizations to display (see **Add Visualizations Drawer**).
* **Save Changes**: (present only if visualizations have been added, moved, or removed since last being saved): A button that saves the current collection of visualizations, and their current layout, as the visualizations that you see when you next return to Observation Deck.
* **Setup Tasks** (present only if setup tasks remain): A pane showing your organization's progress with initial setup tasks (see **Setup Tasks Pane**).

> **Note:** Changes to the default set and layout of visualizations are user-specific; they don't affect what's seen by other users in your organization.

#### Setup Tasks Pane

The setup tasks pane shows your organization's progress with the setup tasks for which Kentik provides wizards for new customers/trial users (see **Getting Started**). Once all setup tasks are complete the pane will no longer show and all subsequent management of your portal settings will be handled via the portal's Settings section (see **Main Settings**), which you access via the main menu.  
  
The **Setup Tasks** pane includes the following controls:

* **Setup Tasks Chart**: An indicator of how far your organization has progressed with its initial setup tasks.
* **View Setup Tasks**: A button that takes you to the Welcome to Network Observability page in the setup tasks wizard (see **Setup Tasks**).

## Visualization Display Area

The main display area of Observation Deck includes a set of visualizations (a.k.a. "widgets") that show different aspects of your organization's network activity.

> **Note:** For new Kentik trials/customers, the initial selection of visualizations on the Observation Deck will be tailored based on the interests expressed on the How Can We Help page of the initial login sequence covered in **Initial Setup Login**.

### Add or Remove Visualizations

To change the visualizations that are displayed in the display area:

* **Add a visualization**: Click the **Add Visualizations** button, which opens the **Add Visualizations Drawer**.
* **Remove a Visualization**: Click the menu icon (hamburger) at the upper right of any visualization, then choose **Remove**.

### Change Visualization Layout

The following visualization-level controls are available to change the appearance of visualizations in the display area (the other visualizations will rearrange to accommodate the change):

* **Reposition**: To move a visualization, click and drag on its repositioning handle, which is just to the left of its title.
* **Resize**: To resize a visualization, click and drag from the handle at the bottom right.

### Default Visualizations

Existing Kentik users who visit the Observation Deck for the first time will see the following visualizations:

* **Traffic Overview**: Visualizations, each on a separate tab, of total traffic volume and the volume of traffic matching each of the following traffic profiles (see **Simple Traffic Profile**):

  + **Internal**: Traffic that both originates and terminates inside your network.
  + **Inbound**: Traffic that originates outside your network and terminates inside your network.
  + **Outbound**: Traffic that originates inside your network and terminates outside your network.
  + **Through**: Traffic that both originates and terminates outside your network.
  + **Other**: Traffic that is not classified by the above profiles (see Traffic Classified as Other).
* **Traffic By**: A table giving a breakdown of the traffic on top X entities in your network such as IPs, Sites, Devices, Interfaces (default) ASNs, etc. The columns of the table vary depending on which type of entity is selected from the drop-down list.
* **Connectivity Costs**:

  + **Estimated cost**: The estimated cost you are paying per billing cycle (monthly) for connectivity.
  + **Ingress**: The sum of the Ingress rates stated for every interface that has been associated with any of your cost groups.
  + **Egress**: The sum of the Egress rates stated for every interface that has been associated with any of your cost groups.
  + **Connectivity Costs**: A button that links to the **Connectivity Costs** landing page.
* **Peering**:

  + **Inbound/Outbound traffic**: A table showing the inbound and outbound rates, based on the aggregate, metric, and time range settings in the **Parameter Controls**, of traffic whose connectivity type is IX, Free Private Peering, or Paid Private Peering.

    > **Note:** If there's no traffic for a given connectivity type, it will not be listed in the table.
  + Discover New Peers: A button that links to the **Discover Peers** landing page.
* **Capacity Planning**:

  + **Status**: A count of interfaces on your network whose utilization is in each of the various capacity planning states (healthy, warning, critical).
  + **Capacity Planning**: A button that links to the **Capacity Planning** landing page.

## Add Visualizations Drawer

The Add Visualizations drawer opens a drawer from the bottom of the screen from which you can choose additional visualizations to display. The available visualizations are grouped into the following tabs:

* **Alerting**: Visualizations related to alerts generated by alert policies.

  > **Note:** For each policy whose alerts are to be covered by a **Health Map by Sites** visualization:
  >
  > + Site must be specified as a Policy Dimension (see Dimensions in **Data Funneling**).
  > + Each device covered by the policy must be assigned to a site (see **Device General Settings**).
* **Cloud**: Visualizations related to traffic to and from your resources that are housed in public cloud providers.
* **Core**: Visualizations drawn from the **Core** section of the Kentik portal, including widgets based on individual dimensions (IPs, ASNs, etc.) or broken down by data sources (devices, clouds. interfaces).
* **Edge**: Visualizations derived from Edge modules including **Connectivity Costs** and **Discover Peers**.
* **Insights**: Insights from Kentik's **Insights & Alerting** system.
* **NMS**: Visualizations from Kentik's Network Monitoring System (see **Network Monitoring**).
* **Protect**: Information, derived from the **DDoS Defense** module, on DDoS attacks and mitigations.
* **Service Provider**: Visualizations derived from Service Provider modules including **CDN Analytics**, **OTT Service Tracking**, and KMI Insights (see **Top Insights**).
* **Synthetics**: Visualizations related to synthetic testing configured in the **Test Control Center** using global and privately-deployed Kentik `ksynth` agents.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(706).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A26Z&se=2026-05-12T09%3A34%3A26Z&sr=c&sp=r&sig=DUsXjPMrq%2BKp%2B9NBn68a2Il%2BA9cbPGFoT4yv13vIIgg%3D)

*Use the Add Visualizations drawer to choose additional visualizations for display.*
