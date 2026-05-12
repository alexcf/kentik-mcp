# Kentik for GCP

Source: https://kb.kentik.com/docs/kentik-for-gcp

---

This article covers the integration of Kentik with the Google Cloud Platform (GCP).

![Diagram illustrating GCP project network architecture with VPN connections and subnets.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/Kentik-for-GCP.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A54Z&se=2026-05-12T09%3A35%3A54Z&sr=c&sp=r&sig=3yg40LkuxwVt2rLy4XqGvMXXLg8Gs7n%2FxlLOlynaAMw%3D)

*Combining GCP and on-prem resources into a hybrid cloud*

> **Note**: See the **Cloud Overview** for an introduction to Kentik cloud setup.

## Process Overview

Integrating GCP with Kentik involves setting up both the GCP environment and the Kentik portal to collect metadata, flow logs, Cloud Run logs, or metrics from Virtual Private Clouds (VPCs). Here's how the process works:

1. **Logging Setup (GCP)**:

   1. Activate flow logs for desired VPC subnets and set up log export to a cloud Pub/Sub topic.
   2. Create and grant access to a pull subscription for Kentik to request entries from the Pub/Sub topic.
   3. Follow the steps in **Grant Metadata Access (GCP)** to display your GCP resources in the **Kentik Map**.
2. **Create a Kentik Cloud Export**:

   1. Configure a new "cloud export" to ingest data from GCP.

Once the setup is complete, you can use the Kentik portal to:

* Monitor your GCP network traffic.
* Visualize resource utilization.
* Gain insights for optimizing network performance and enhancing security monitoring.

> **Note:** See Google’s Cloud VPC Documentation for more details:
>
> * **Using VPC Flow Logs**
> * **What Is Cloud Pub/Sub**

## Logging Setup (GCP)

To set up flow log export from GCP, follow these topics.

### Process Options

GCP users can set up flow logs and cloud Pub/Sub using:

* **GCP Console**: Manage resources through a UI (see **GCP Console**).
* **Gcloud compute**: A command-line tool for resource management (see **gcloud compute**).
* **Compute Engine API**: A REST API for creating and running VMs (see **Compute Engine API**).

> **Note:** This guide covers the use of the GCP console only.

### Enable VPC Flow Logs

To enable VPC flow logs for existing subnets in GCP:

1. Ensure the current project is selected in the console.
2. Click the menu icon and go to Networking » VPC Network » **VPC Networks**
3. Find and click the subnet (e.g., "default") you want to monitor.
4. On the Subnet details page, click **Edit**.
5. Set the flow logs to “On”.
6. Click **Configure Logs**, then check **Include metadata** to allow Kentik to ingest flow logs
7. Adjust these settings as needed:

   1. Aggregation interval: Any interval is supported
   2. Sample rate: Adjust between 0% and 100%. Default is 50%.

      > **Note:** Check **Estimated logs generated per day** for cost impact
8. Click **Save**.

> **Note:** For more details, see Google’s Cloud VPC Documentation:
>
> * **Networks and subnets**
> * **Enabling VPC flow logging**

### Create a New Topic

To configure logging and create a new Pub/Sub topic in GCP:![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(75).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A54Z&se=2026-05-12T09%3A35%3A54Z&sr=c&sp=r&sig=3yg40LkuxwVt2rLy4XqGvMXXLg8Gs7n%2FxlLOlynaAMw%3D)

1. Ensure you're in the correct project, then click the menu icon and go to Operations » Logging » **Logs Router**
2. Click **Create Sink**, then enter a Name and Description under **Sink details**.
3. In **Sink destination**, select Cloud Pub/Sub topic as the sink service.
4. Click **Create a Topic**, enter a Topic ID (name), and click **Create Topic**.
5. In **Choose logs to include**, enter the filter in **Build inclusion filter**:  
   `resource.type="gce_subnetwork" AND log_id("compute.googleapis.com/vpc_flows")`
6. Click **Preview Logs** to open **Logs Explorer** in a new tab and view the logs included in the sink.
7. If logs are incorrect, adjust the inclusion filter. If correct, click **Next**.
8. Optional: If you need to exclude specific logs, see the Google documentation on **Creating a Sink**.
9. Click **Create Sink** to finish and return to the **Logs Router** page. The new sink will be listed and you can edit it via the **Edit Sink** option.

> **Notes:**
>
> * Consider a more restrictive filter to reduce costs (e.g., specific port, protocol, subnet). See the Google documentation on **Networking queries**.
> * For more details, see Google’s Cloud VPC documentation on **Overview of Logs Export** and **Exporting with the Logs Viewer**.

### Create a Pull Subscription

To set up a “pull” subscription for Kentik’s flow log collection:  
![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(76).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A54Z&se=2026-05-12T09%3A35%3A54Z&sr=c&sp=r&sig=3yg40LkuxwVt2rLy4XqGvMXXLg8Gs7n%2FxlLOlynaAMw%3D)

1. Ensure you're in the correct project where the logging topic was created.
2. Click the menu icon on the main navbar, go to the **Big Data** section and choose Pub/Sub » **Topics.**
3. In the left-hand pane, find your topic and click the **More** icon (vertical dots) next to it.
4. Choose **Create Subscription** from the submenu.
5. Configure the subscription:

   1. Subscription ID: Enter a unique name (no spaces)
   2. Delivery type: Select “Pull”
   3. Configure other properties as needed (refer to Google’s documentation for details)
6. Click **Create**. A confirmation popup will appear, and you'll be directed to the **Subscription Details** page.

> **Note:** For more information, see Google’s Cloud VPC documentation:
>
> * **Cloud Pub/Sub**
> * **Pull subscription**

### Set Permissions

To allow Kentik access to your Google Cloud subscription for flow logs, follow these steps:

1. Ensure you're in the correct project (see **Create a Pull Subscription**).
2. Go to the **Subscriptions** page via the menuBig Data » Pub/Sub » **Subscriptions.**
3. Find and click the checkbox for the topic you created. Permissions will appear on the right.
4. Click **Add Member** and enter `kentik-vpc-flow@kentik-vpc-flow.iam.gserviceaccount.com` in the **New** **members** field.
5. From the Role drop-down, find "Pub/Sub", then choose “Pub/Sub Subscriber.”
6. Click **Add Another Role**, then repeat, this time choosing “Pub/Sub Viewer.”
7. Click **Save** to confirm. A popup will indicate success.

   > **Note:** For more information, see Google’s Cloud VPC documentation on **Permissions and roles****.**

## Grant Metadata Access (GCP)

To display your GCP resources in the **Kentik Map**, grant metadata access to Kentik by following these steps:

1. ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(80).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A54Z&se=2026-05-12T09%3A35%3A54Z&sr=c&sp=r&sig=3yg40LkuxwVt2rLy4XqGvMXXLg8Gs7n%2FxlLOlynaAMw%3D)Ensure you’re in the correct project, see **Logging Setup (GCP)****.**
2. Go to IAM & Admin » **IAM** via the sidebar menu.
3. Click **Grant Access** to open the drawer
4. Enter `kentik-vpc-flow@kentik-vpc-flow.iam.gserviceaccount.com` in the **Add Principals** field.
5. In the **Assign Roles** field, select **Compute Viewer** (not Compute Network Viewer).
6. Click **Save** to confirm and return to the **Permissions** tab.
7. In the **View by Principals** tab, ensure the service account is listed with the “Compute Viewer” role.

> **Notes:**
>
> * Ensure you have `compute.networks.list` and `resourcemanager.projects.setIamPolicy` permissions for the Google project.
> * For projects in a folder structure, add the service account `kentik-vpc-flow@kentik-vpc-flow.iam.gserviceaccount.com` as Principal at the top-level folder and assign the “Compute Viewer” role to allow nested projects to inherit permissions.

## Create a Kentik Cloud Export

To create a GCP cloud export in the Kentik portal, follow these steps.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(81).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A54Z&se=2026-05-12T09%3A35%3A54Z&sr=c&sp=r&sig=3yg40LkuxwVt2rLy4XqGvMXXLg8Gs7n%2FxlLOlynaAMw%3D)

1. Navigate to Settings » **Public Clouds**.
2. Click **Create Cloud Export**.
3. Click **GCP Cloud**.
4. Under **Observability Features**, select the data types to collect:

   1. Metadata collection (Required): Automatically selected.
   2. Flow log collection: Select to collect flow logs.
   3. GCP Cloud Run Log Collection: Select to create a second cloud export dedicated to collecting Cloud Run Logs (see Google's **Cloud Run Logging** docs).
   4. Cloud metrics history: Select to collect GCP metrics.
5. Click the green arrow to proceed.
6. Specify the GCP project ID with the Cloud Pub/Sub topic you created for publishing flow logs from your VPC subnets (see **Create a New Topic**).
7. Provide the subscription name you created for Kentik to subscribe to your Pub/Sub topic (see **Create a Pull Subscription**).
8. Click the green arrow to proceed.
9. Enter a cloud export name and description or accept the default.
10. Choose the appropriate Kentik billing plan for the cloud export from the dropdown.
11. Click **Save** to finalize the cloud export and return to the Public Clouds page, where the new export will be listed.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(82).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A54Z&se=2026-05-12T09%3A35%3A54Z&sr=c&sp=r&sig=3yg40LkuxwVt2rLy4XqGvMXXLg8Gs7n%2FxlLOlynaAMw%3D)

*The Public Clouds page lists the clouds registered with Kentik.*

Check the status of your GCP cloud:

1. From the **Cloud Config Status** pane, click **View Details** to open the GCP Configuration Status page.
2. Check for any configuration errors (red circles).
3. Click an export to open the **Config Status Details** drawer

> **Note:** Flow status for Metadata-only exports always shows as 'red circle' since these exports do not have flow or sampling.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(83).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A54Z&se=2026-05-12T09%3A35%3A54Z&sr=c&sp=r&sig=3yg40LkuxwVt2rLy4XqGvMXXLg8Gs7n%2FxlLOlynaAMw%3D)

*Error reported for a cloud export that's missing permissions for metadata access.*

## Using Your Cloud Export

Once the setup process is complete, you can view and utilize your cloud export in Kentik:

* **Cloud Exports List**:

  + Go to Settings » **Public Clouds** to see the updated list of cloud exports.
  + A new cloud export will be listed, representing the VPC subnets whose logs are pulled from the specified subscription.
* **Devices Column**:

  + Each VPC subnet sending flow logs is listed as a cloud device.
  + Devices are named after their respective VPC subnet.
  + These names can be used as group-by and filter values in Kentik queries using the Device Name dimension.
* **Metadata and Mapping**:

  + The collected metadata, such as routing tables, security groups, and ACLs, enables Kentik to automatically map and visualize the topology of your GCP resources in the **Kentik Map**.

## GCP Endpoints Lists

Kentik needs permission to access selected GCP endpoints on your behalf in order to collect metadata and metrics, as detailed in the following lists.

### GCP Metadata Endpoints

**@google-cloud/compute**:

* `BackendServicesClient`
* `ExternalVpnGatewaysClient`
* `FirewallsClient`
* `ForwardingRulesClient`
* `GlobalForwardingRulesClient`
* `GlobalNetworkEndpointGroupsClient`
* `HealthChecksClient`
* `InstancesClient`
* `InstanceGroupsClient`
* `InterconnectsClient`
* `InterconnectAttachmentsClient`
* `NetworksClient`
* `NetworkEndpointGroupsClient`
* `NetworkFirewallPoliciesClient`
* `ProjectsClient`
* `RegionsClient`
* `RegionBackendServicesClient`
* `RegionHealthChecksClient`
* `RegionHealthCheckServicesClient`
* `RegionInstanceGroupsClient`
* `RegionNetworkEndpointGroupsClient`
* `RegionNetworkFirewallPoliciesClient`
* `RegionSecurityPoliciesClient`
* `RegionTargetHttpProxiesClient`
* `RegionTargetHttpsProxiesClient`
* `RegionTargetTcpProxiesClient`
* `RegionUrlMapsClient`
* `RoutersClient`
* `RoutersClient`
* `RoutersClient`
* `RoutersClient`
* `RoutesClient`
* `SubnetworksClient`
* `TargetGrpcProxiesClient`
* `TargetHttpProxiesClient`
* `TargetHttpsProxiesClient`
* `TargetInstancesClient`
* `TargetSslProxiesClient`
* `TargetTcpProxiesClient`
* `TargetPoolsClient`
* `TargetVpnGatewaysClient`
* `UrlMapsClient`
* `VpnGatewaysClient`
* `VpnTunnelsClient`
* `ZonesClient`
