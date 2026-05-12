# Kentik Cloud

Source: https://kb.kentik.com/docs/kentik-cloud

---

This article covers the Kentik Cloud page.

> **Note:** For general information about monitoring your cloud resources with Kentik, start with **Cloud Overview**.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/KC-Main-orig.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A43Z&se=2026-05-12T09%3A27%3A43Z&sr=c&sp=r&sig=RMXC33oqhUJ%2BZ4cnExPS2zNopgJdQf8itLbqWEMcvTs%3D)

*Kentik Cloud is a dashboard for cloud-related visualizations.*

## About Kentik Cloud

The Kentik Cloud page can be accessed by clicking the **Cloud** heading on the main nav menu of the Kentik portal. This page features a curated set of visualizations for a high-level overview of various aspects of your cloud traffic and infrastructure.

### Key Features

* **Customizable Layout**: Drag the title bar of any visualization to rearrange it on the page.
* **Supported Providers**: Visualizations are available for AWS, Azure, GCP, and OCI.

## Kentik Cloud UI

The Kentik Cloud page includes the following main functional elements:

* **Show control**: Located at the right of the page's title bar, this control allows you to show a specific cloud providers' visualizations in the display area or to show all providers.

  > **Notes**:
  >
  > + Only providers for which your organization has cloud resources are shown as options in the Show control.
  > + The Cloud Traffic Overview and Public Cloud Flow visualizations update automatically based on the selected provider (see **Common Visualizations**).
* **Display area**: The section of the page where visualizations are displayed.

### Visualization Title Bar Elements

Each individual visualization includes the following elements in its title bar:

* **Information icon**: Hover over this icon see a brief explanation of the visualization.
* **Action button (vertical ellipsis)**: Click to reveal the View in Data Explorer option, which opens the visualization in **Data Explorer**.

## Kentik Cloud Visualizations

The Kentik Cloud page includes both **Common Visualizations** and **Provider-Specific Visualizations**.

### Common Visualizations

These visualizations are available for all cloud providers (AWS, Azure, GCP, OCI):

* **Cloud Traffic Overview**: Tabs summarizing traffic volume as divided into the following profiles (see **Simple Traffic Profile**). Click the tab for the corresponding visualization:

  + Total: Total traffic volume across all profiles.
  + Internal: Traffic both originating and terminating inside your network.
  + Inbound: Traffic originating outside your network and terminating inside.
  + Outbound: Traffic originating inside your network and terminating outside.
  + Other: Traffic not classified by the above profiles (see **Traffic Classified as Other**).

    > **Note**: For more details on each tab, use the **View in Data Explorer** button.
* **Public Cloud Flow**: Displays your traffic flows per second for each cloud provider. Allows you to view your total flow consumption at a glance.
* **All Regions**: Shows traffic volume across all regions of the provider and a breakdown by traffic profile. Example name: "GCP – All Regions".
* **Insights Overview**: Provides insights by identifying interesting issues with your data and configurations (see **Insights)**.

### Provider-Specific Visualizations

These visualizations are only available if your organization has registered a cloud in Kentik for the mentioned provider:

* **Gateway Traffic (AWS, Azure only)**: Displays increases and decreases in utilization for various types of gateways in your infrastructure. Unexpected increases in outbound traffic can lead to higher data transfer costs. Example name: "AWS Gateway Traffic".
* **Entity Explorer (Azure, AWS, OCI only)**: A searchable list of provider-specific network entities and components related to your cloud resources.
* **AWS InterVPC Traffic by Transit Gateway**: A table providing a top-X breakdown of the traffic routed through your AWS transit gateways.
* **AWS Security Group Reject Actions**: A table listing source and destination AWS VPCs where traffic is being rejected. This is useful for security analysis and connectivity troubleshooting.

  ---
* © 2014-25 Kentik
