# Kentik for OCI

Source: https://kb.kentik.com/docs/kentik-for-oci

---

This article discusses the integration of Kentik with Oracle Cloud Infrastructure (OCI).

![An example OCI tenancy and region with three domains and one VCN.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/OCI-Diagram-627h1120w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A07Z&se=2026-05-12T09%3A46%3A07Z&sr=c&sp=r&sig=sOP0ssNBeruVGElCWMlWD4coq4aP2JktGbBe6i32mEg%3D)

*An example OCI tenancy and region with three domains and one VCN.*

> **Note**: See the **Cloud Overview** for an introduction to Kentik cloud setup.

## Prerequisites

Before beginning the integration process, ensure you have the following in place:

* **An Active OCI Tenancy:** You must have an active Oracle Cloud Infrastructure account.
* **An Existing VCN:** Ensure you have at least one Virtual Cloud Network (VCN) already created and running in OCI.
* **Administrator Permissions:** You must have the necessary OCI administrator privileges to create policies, connectors, log groups, and storage buckets.

## Process Overview

Integrating Oracle Cloud Infrastructure (OCI) with Kentik involves setting up both the OCI environment and the Kentik portal to collect data from Virtual Cloud Networks (VCNs), subnets, Virtual Network Interface Cards (VNICs), and other sources. Here's how the process works:

1. **Logging Setup (OCI)**

   1. **Configure an OCI Bucket**
   2. **Enable Flow Logs**
   3. **Configure an OCI Connector**
   4. **Create an OCI Policy**
2. **Create a Kentik Cloud Export**

   1. Configure a new "cloud export" in the Kentik portal to ingest data from OCI.

Once the setup is complete, you can use the Kentik portal to:

* Monitor your OCI network traffic.
* Visualize resource utilization.
* Gain insights for optimizing network performance and enhancing security monitoring.

## About OCI Flow Logs

In OCI, flow logging is managed through OCI Logging, which captures, stores, indexes, and monitors log data. Some key points about using OCI flow logs with Kentik include:

* **VCN Flow Logs**

  + Used for monitoring and diagnosing network traffic within OCI Virtual Cloud Networks (VCNs).
  + Can be enabled for an entire VCN or selectively for specific subnets, VNICs, or resources like instances or load balancers.
* **Log Format and Storage**

  + VCN flow logs are in JSON format and contain detailed network flow information.
  + They are sent to an OCI Object Storage bucket, aggregating logs from VCNs within specific compartments or across multiple compartments.
* **Kentik Integration**

  + Kentik accesses VCN flow logs via OCI APIs.
  + Logs are forwarded to the **Kentik Data Engine** (KDE), where they are normalized, enriched with OCI-specific metadata, and stored for analysis.
  + In Kentik, logs appear as a single "cloud export" associated with a "cloud device” on the **Public Clouds Page**.

For more information, see these OCI docs:

* **What is a Tenancy?**
* **Logging Overview**
* **VCN Flow Logs**
* **Details for VCN Flow Logs**
* **Object Storage Buckets**

## Logging Setup (OCI)

Setting up flow log export in the **OCI dashboard**is covered here.

> **Note:** Use the default compartment for flow log export unless specified otherwise (see Oracle's **What is a Compartment**).

### Configure an OCI Bucket

To set up an OCI bucket for Kentik to access flow logs, follow these steps.

#### Create an OCI Bucket

1. ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(710).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A07Z&se=2026-05-12T09%3A46%3A07Z&sr=c&sp=r&sig=sOP0ssNBeruVGElCWMlWD4coq4aP2JktGbBe6i32mEg%3D)

   Log in to the **OCI dashboard** for your tenancy.
2. Click the menu icon at the top left of any page.
3. In the menu sidebar, select **Storage**, then choose Object Storage & Archive Storage» **Buckets** on the main menu.
4. On the **Buckets** tab, click **Create Bucket** to open the **Create Bucket** drawer.
5. Enter a name for the new bucket in **Bucket Name**.

   > **Note:** Additional settings in the **Create Bucket** drawer are not necessary for a Kentik Cloud Export (see **Creating an Object Storage Bucket**).
6. Click **Create** to establish the bucket and return to the **Buckets** list.

#### Create Policy for Bucket

To enable access to the new bucket for the Kentik cloud export tool, follow these steps:

1. Log in to the **OCI dashboard** for your tenancy.
2. Navigate to **Identity & Security** **»** **Identity** **»** **Policies**.
3. Click **Create Policy**. Give the policy a **Name** and **Description**.
4. Turn on the **Show manual editor** switch to open a text input field.
5. Enter the policy statements below into the input field, replacing placeholders with actual values:

   ```
   Define group groupRef as "groupId"
   Allow group groupRef to READ buckets in tenancy WHERE target.bucket.name=<bucketName>
   Allow group groupRef to READ objects in tenancy WHERE target.bucket.name=<bucketName>
   ```
6. Click **Create** to save the policy and go to the new policy’s page.

### Create a Log Group

To capture traffic information for your VCN in OCI, you must first create a log group.![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(88).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A07Z&se=2026-05-12T09%3A46%3A07Z&sr=c&sp=r&sig=sOP0ssNBeruVGElCWMlWD4coq4aP2JktGbBe6i32mEg%3D)

1. Log in to the **OCI dashboard** for your tenancy.
2. Click the menu icon at the top left, and navigate to **Observability & Management** » **Log Groups** (under Logging).
3. Click **Create Log Group**.
4. Enter a unique **Name** for the log group.
5. Ensure that the **Compartment** drop-down is set to your root compartment, e.g., `your_tenancy_name (root)`.
6. Enter a **Description** (optional).
7. Click **Create** to finalize the log group.

### Create a Capture Filter

Next, create a capture filter for enabling flow logs in OCI:

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(89).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A07Z&se=2026-05-12T09%3A46%3A07Z&sr=c&sp=r&sig=sOP0ssNBeruVGElCWMlWD4coq4aP2JktGbBe6i32mEg%3D)

1. Click the menu icon at the top left, select **Networking**, and navigate to **Capture Filters** (under Network Command Center).
2. Click **Create Capture Filter**.
3. Enter a unique **Name** for the capture filter.
4. Ensure that the **Compartment** drop-down is set to your root compartment, e.g., `your_tenancy_name (root)`.
5. Choose a preferred sampling rate from the dropdown, which determines the percentage of flows for which OCI will generate flow logs.

   > **Note:** Additional settings in this drawer are optional for a Kentik cloud export (see Oracle’s **Creating a Capture Filter**).
6. Click **Create capture filter**.

> **TIP:** Capture filters can be configured to capture flows based on criteria other than sampling rate (see the Oracle article **Capture Filters**).

### Enable Flow Logs

With your log group and capture filter created, you can now enable flow logs:

1. Click the menu icon at the top left, select **Networking**, and choose **Flow logs** (under Network Command Center).
2. Click **Enable flow logs** to open the wizard.
3. Enter a **File name prefix** for the flow logs.
4. In the **Flow log destination** pane, click the drop-down and select the **Log Group** you created in the previous step.
5. In the **Capture filter** pane, click the drop-down and select the **Capture Filter** you created in the previous step.
6. Click **Next** to proceed to the **Enablement points** step.

#### Add Enablement Point

To add an enablement point for flow logs in OCI, follow these steps:

1. In the **Enable flow logs** wizard, click **Next** to reach the **Enablement points** step.
2. Click **Add enablement points** to open the dialog.
3. Choose **Virtual Cloud Network**.
4. Click **Continue** to open the **Add virtual cloud network enablement points** drawer.
5. Use the dropdown to select a VCN. Add more VCNs with the **+ Another Enablement point** button if needed.
6. Click **Add enablement points** to close the drawer and return to the wizard.
7. Click **Next** to continue to the **Review and create** step.
8. Click **Enable flow logs** to finalize the configuration and view the new flow log setup page.

### Configure an OCI Connector

An OCI Connector is a managed service offered by Oracle that allows for the integration and automation of data flows across various OCI services. The service connector facilitates the flow of logs or metrics to external destinations such as Kentik. The topics below outline the steps to create an OCI Connector.

#### Create a Connector

To set up an OCI connector, follow these steps:

1. ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(90).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A07Z&se=2026-05-12T09%3A46%3A07Z&sr=c&sp=r&sig=sOP0ssNBeruVGElCWMlWD4coq4aP2JktGbBe6i32mEg%3D)Log in to the **OCI dashboard** for your tenancy.
2. Click the menu icon at the top left of any page, and from the menu sidebar, select **Analytics & AI**.
3. Choose **Messaging****»** **Connector Hub** on the main menu.
4. Click **Create Connector** to open theCreate connector drawer.
5. Enter a **Connector name** and **Description** for the connector.
6. Ensure the **Resource compartment** dropdown is set to your root compartment, e.g., `your_tenancy_name (root)`.
7. Select **Logging** from the **Source** drop-down to display the Configure source pane.
8. Complete the **Configure source** pane as described in **Configure Source Connection**.
9. Select **Object Storage** from the **Target** dropdown to display the **Configure target** pane.
10. Fill out the **Configure target** pane as described in **Configure Target Connection**.

    > **Note:** Additional settings within the **Create connector** page are optional for a Kentik cloud export (see **Creating a Connector**).
11. After configuring the target, click **Create** in the callout that appears to create a policy. The callout will update to "Policy created…" with a link to the new policy.
12. Click the overall **Create** button to close the drawer and access the new connector page.

> **Note:** The **Enable logs** switch in the **Create connector** drawer is for connector logging and can be left **Off**.

#### Configure Source Connection

To set up the source connection in the **Create connector** drawer:

1. Ensure the **Compartment name** dropdown is set to your root compartment, e.g., `your_tenancy_name (root)`.
2. Click the **Log group** dropdown and choose the log group created in **Create a Log Group**.
3. Click the **Logs** dropdown and select the flow log created from **Enable Flow Logs**.

> **Note:** Additional settings in the **Configure source** pane are optional for a Kentik cloud export (see**Create a Connector**).

#### Configure Target Connection

To set up the target connection in the **Create connector** drawer:

1. Ensure the **Compartment name** dropdown is set to your root compartment, e.g., `your_tenancy_name (root)`.
2. Choose the bucket you created in **Create an OCI Bucket** from the **Bucket** dropdown.
3. The connector’s target will be an object in the bucket.

   1. (Optional) Specify a prefix (e.g., `flow-logs-bucket`) in the **Object Name Prefix** field to optimize object location.

> **Note:** The default batch size is 100MB, and the default batch time is 7 seconds. To modify these defaults, click **Show additional options** at the bottom of the pane.

### Create an OCI Policy

There are two options for creating a policy for authorizing Kentik access to your OCI environment:

* **Cross-Tenancy Policy**: Create a policy for Kentik’s tenant to access certain resources in your tenant.
* **Custom User Policy:** Create a policy for a custom user/group you create for Kentik to access certain resources in your tenant (see **Create an OCI User**).

> **Notes**:
>
> * Kentik recommends the **Create a Cross-Tenancy Policy** option for simplicity and ease of setup.
> * See the **OCI documentation**for more on cross-tenancy policies.

#### Cross-Tenancy Policy

To create a cross-tenancy access policy in your OCI tenant, follow these steps.

1. Open a new browser tab, log in to the Kentik portal, and follow the steps in **Create a Kentik Cloud Export** (later in this article) to generate and copy your **Autogenerated Policy**. This policy automatically includes the necessary Kentik IDs and permission statements. Once copied, return to your OCI tab to continue these steps.
2. Log in to the **OCI dashboard** for your tenancy.
3. Navigate to **Identity & Security** **»** **Identity** **»** **Policies**.
4. Click **Create Policy**. Give the policy a **Name** and **Description**.
5. Turn on the **Show manual editor** switch to open a text input field.
6. Paste in the auto-generated policy from the Kentik portal.
7. Click **Create** to save and go to the new policy’s page.

### OCI Console Info for Portal

> **TIP**: Copy and paste these values into a blank text document as you go. You will need them all at once when **configuring the export** in the Kentik portal.

| Field Name | Console Page | Notes |
| --- | --- | --- |
| **Tenancy OCID** | Tenancy details | — |
| **User OCID** | User details | Required only if you created an OCI user for Kentik. |
| **Home Region** | Domain details | Use the city name in parentheses. |
| **Bucket Name** | Bucket details | Bucket where flow logs are directed by the connector. |
| **Bucket Namespace** | Bucket details | Unique namespace name assigned when your OCI tenancy (account) was created. |
| **Service Connector OCID** | Connector details | — |
| **Flow Object Name Prefix** | Connector details | — |
| **Compartment ID** | Compartments page | Same as Tenancy OCID unless using a custom compartment for logging. |

## Create a Kentik Cloud Export

Configuring an OCI cloud export in the Kentik portal is covered in the following topics.

![Create an OCI cloud export in the Kentik portal and select from the available observability features.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/KFO-oci-provider-features.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A07Z&se=2026-05-12T09%3A46%3A07Z&sr=c&sp=r&sig=sOP0ssNBeruVGElCWMlWD4coq4aP2JktGbBe6i32mEg%3D)

*Creating a new OCI cloud export in the Kentik portal, while selecting from available observability features.*

### Metadata-Only Export

To set up a new OCI metadata-only cloud export in the Kentik portal, follow these steps:

1. Go to **Settings »** **Public Clouds** in the main menu.
2. Click **Create Cloud Export** to start the configuration wizard.
3. Choose **OCI Cloud** under **Provider and Features**.
4. Under **Observability Features**, accept the default **Metadata collection** (automatically selected) and click the green arrow to proceed.

   > **Note**: Leave **Flow log collection** unselected when configuring a metadata-only export.
5. ![Kentik OCI export configuration showing user ID, tenancy ID, and autogenerated policy details.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/KFO-metadata-only-export-custom-user.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A07Z&se=2026-05-12T09%3A46%3A07Z&sr=c&sp=r&sig=sOP0ssNBeruVGElCWMlWD4coq4aP2JktGbBe6i32mEg%3D)Under **API Access**, select the **Cross Tenancy Policy** or **Custom User** tab and enter the required OCI Account Info:

   1. **Tenancy ID**: Your OCI tenancy ID.
   2. **OCI Default Region**: Accept the default region selection or select from the dropdown (e.g., `us-ashburn-1`).
   3. **User ID** (Custom User tab only): The OCI user ID you created for Kentik in **Create an OCI User**.
6. Click **Download Public Key** for use in **Create an OCI User**(Custom User policy only)**.**
7. Click **Validate Permissions**

   1. Checks that Kentik can access the necessary OCI API endpoints to gather your data.
   2. Shows a summary panel of the permissions status per endpoint, grouped by Compartment ID and Region.
8. Optional: Enter one or more Compartment IDs.

   1. The default Compartment ID is the Tenancy ID
   2. Click **Load Compartments** to select from a list of additional compartments
9. Click the button to copy the **Autogenerated Policy**. Savefor later use in the OCI console (see **Create an OCI User**).
10. Click the green arrow to proceed to the final step.
11. Enter a cloud export name and description or accept the defaults.
12. Select the appropriate Kentik **Billing Plan** for the cloud export from the dropdown.
13. Click **Save** to finalize the cloud export and return to the **Public Clouds Page**, where the new export will be listed.

> **Note**: If you encounter errors:
>
> * Check that the user created in **Create an OCI User** is assigned to a group (see **Create an OCI User Group**).
> * Check for errors (e.g., placeholders instead of actual values) in the policy creation statements used in **Custom User Policy****.**

### Flow Logs and Metadata Export

To set up a new OCI flow logs and metadata export, follow these steps:

1. Complete the first 4 steps of **Metadata-Only Export** setup, while also selecting **Flow log collection** under **Observability Features**.
2. ![Kentik OCI export configuration settings including API access and compartment ID details.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/KFO-flow-and-metadata-cross-tenancy.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A07Z&se=2026-05-12T09%3A46%3A07Z&sr=c&sp=r&sig=sOP0ssNBeruVGElCWMlWD4coq4aP2JktGbBe6i32mEg%3D)Under **API Access**, select the **Cross Tenancy Policy** or **Custom User** tab and enter the required OCI Account Info:

   1. **Tenancy ID**: Your OCI tenancy ID.
   2. **OCI Default Region**: Accept the default region selection or select from the dropdown (e.g., `us-ashburn-1`).
   3. **User ID** (Custom User tab only): The OCI user ID you created for Kentik in **Create an OCI User**.
3. Click **Download Public Key** for use in **Create an OCI User**(Custom User policy only)**.**
4. Click **Validate Permissions**

   1. Checks that Kentik can access the necessary OCI API endpoints to gather your data.
   2. Shows a summary panel of the permissions status per endpoint, grouped by Compartment ID and Region.
5. Optional: Enter one or more Compartment IDs.

   1. The default Compartment ID is the Tenancy ID
   2. Click **Load Compartments** to select from a list of additional compartments
6. In the **Flowlogs Bucket Configuration** section, enter the following values:

   1. **Bucket Name** (Required): The name of the OCI bucket you assigned in **Configure an OCI Bucket**.
   2. **Bucket Namespace** (Required): The unique namespace name of your OCI tenancy (e.g., `idovcl4rlc88`).
   3. **Service Connector OCID** (Required): The OCID of the service connector you set up in **Configure an OCI Connector**.
   4. **Flow Object Name Prefix**: The Object Name Prefix you assigned in **Configure Target Connection**, if any.
7. Click the button to copy the **Autogenerated Policy**.Save for later use in the OCI console (see **Create an OCI User**).
8. Click the green arrow to proceed.
9. Enter a cloud export name and description or accept the defaults.
10. Select the appropriate Kentik **Billing Plan** for the cloud export from the dropdown.
11. Click **Save** to finalize the cloud export and return to the **Public Clouds Page**, where the new export will be listed.

### Autogenerated Policy

![Checkbox indicating assistance in generating policy content for users.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/Screenshot 2025-08-20 at 10.16.12 AM.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A07Z&se=2026-05-12T09%3A46%3A07Z&sr=c&sp=r&sig=sOP0ssNBeruVGElCWMlWD4coq4aP2JktGbBe6i32mEg%3D)When you select **Help me generate policy** in the Create Cloud Export Configuration wizard, in the next step you’ll get an autogenerated policy based on the chosen configuration settings. You can use this in step 7 of **Custom User Policy** in place of the example statements.

![Autogenerated policy defining tenancy and group permissions for KentikMetadataGroup.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/KFO-autogenerated-policy.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A07Z&se=2026-05-12T09%3A46%3A07Z&sr=c&sp=r&sig=sOP0ssNBeruVGElCWMlWD4coq4aP2JktGbBe6i32mEg%3D)

*Autogenerated policy defining tenancy and group permissions for KentikMetadataGroup.*

## Using Your Cloud Export

Once the setup process is complete, you can view and utilize your cloud export in the Kentik portal:

* **Cloud Exports List**:

  + Go to **Settings »** **Public Clouds** to see the updated list of cloud exports.
  + A new cloud export will be listed, representing the VPC subnets whose logs are pulled from the specified subscription.
* **Devices Column**:

  + Each VPC subnet sending flow logs is listed as a cloud device.
  + Devices are named after their respective VPC subnet.
  + These names can be used as group-by and filter values in Kentik queries using the Device Name dimension.
* **Metadata and Mapping**:

  + The collected metadata, such as routing tables, security groups, and ACLs, enables Kentik to automatically map and visualize the topology of your OCI resources in the Kentik Map.

> **Note**: In some cases (e.g., high volume of flow records) Kentik may optimize the ingest of flow records by creating multiple cloud devices within a single cloud export.

## OCI Policy Statements

The policy created in the above workflow gives the Kentik cloud export tool broad access across your entire OCI tenancy. If preferable, given your organization's security policies, you can instead limit Kentik to the narrowest subset of OCI API calls needed to access your metadata and resources for flow export.

The policy statement below, which grants read-only access to these specific calls, can be used in step 7 of **Custom User Policy** in place of the example statement.

```
Define group groupRef as "<group OCID goes here>"
Allow group groupRef to INSPECT tenancies in tenancy
Allow group groupRef to READ vcns in tenancy
Allow group groupRef to READ capture-filters in tenancy
Allow group groupRef to READ cpes in tenancy
Allow group groupRef to READ nat-gateways in tenancy
Allow group groupRef to READ drg-object in tenancy
Allow group groupRef to READ cross-connects in tenancy
Allow group groupRef to READ route-tables in tenancy
Allow group groupRef to READ virtual-circuits in tenancy
Allow group groupRef to READ local-peering-gateways in tenancy
Allow group groupRef to READ network-security-groups in tenancy
Allow group groupRef to READ drg-attachments in tenancy
Allow group groupRef to READ drg-route-distributions in tenancy
Allow group groupRef to READ drg-route-tables in tenancy
Allow group groupRef to READ subnets in tenancy
Allow group groupRef to READ security-lists in tenancy
Allow group groupRef to READ ipsec-connections in tenancy
Allow group groupRef to READ internet-gateways in tenancy
Allow group groupRef to READ metrics in tenancy
Allow group groupRef to INSPECT metrics in tenancy
```

> **Note:** The group OCID is found on the group's page in the OCI Console.

## OCI Endpoints List

The Kentik cloud export tool uses the API operations below.

| Service | Operation |
| --- | --- |
| `computeClient` | `listInstances` |
| `computeClient` | `listVnicAttachments` |
| `IdentityClient` | `listRegionSubscriptions` |
| `ObjectStorage` | `listObjects` |
| `ObjectStorage` | `getObject` |
| `ObjectStorage` | `listBuckets` |
| `VirtualNetworkClient` | `listVcns` |
| `VirtualNetworkClient` | `listSubnets` |
| `VirtualNetworkClient` | `listRouteTables` |
| `VirtualNetworkClient` | `listSecurityLists` |
| `VirtualNetworkClient` | `listNetworkSecurityGroups` |
| `VirtualNetworkClient` | `listCrossConnects` |
| `VirtualNetworkClient` | `listIPSecConnections` |
| `VirtualNetworkClient` | `listVirtualCircuits` |
| `VirtualNetworkClient` | `listLocalPeeringGateways` |
| `VirtualNetworkClient` | `listNatGateways` |
| `VirtualNetworkClient` | `listInternetGateways` |
| `VirtualNetworkClient` | `listDrgs` |
| `VirtualNetworkClient` | `listDrgsRouteTables` |
| `VirtualNetworkClient` | `listInternalDrgs` |
| `VirtualNetworkClient` | `listCpes` |
| `VirtualNetworkClient` | `listDrgAttachments` |

## Alternative Setup: Custom User

### **Custom User Policy**

To set up an OCI access policy for a custom user/group you create for Kentik, follow these steps:

1. ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(87).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A07Z&se=2026-05-12T09%3A46%3A07Z&sr=c&sp=r&sig=sOP0ssNBeruVGElCWMlWD4coq4aP2JktGbBe6i32mEg%3D)Follow all steps in **Create an OCI User**.
2. Log in to the **OCI dashboard** for your tenancy.
3. Navigate to **Identity & Security** **»** **Identity** **»** **Policies**.
4. Click **Create Policy**. Give the policy a **Name** and **Description**.
5. Choose the compartment where the policy will apply (default is the root compartment).
6. Turn on the **Show manual editor** switch to open a text input field.
7. Enter the policy statements below, replacing the placeholder with the group OCID for the user you created:  
   `Define group groupRef as "groupId"  
    Allow group groupRef to READ all-resources in tenancy`
8. Click **Create** to save the policy and go to the new policy’s page.

> **Note**: The above policy statement grants read-only access to all tenancy metadata. Adjust permissions as needed (see **OCI Policy Statements**).

### Create an OCI User

To set up an OCI user in your tenancy on behalf of Kentik, follow these steps.

> **Note**: Avoid these steps by opting for a **Cross-Tenancy Policy**.

#### **Navigate to Create User**

1. Log in to the **OCI dashboard** for your tenancy.
2. Click the menu icon on any OCI console page, and from the left sidebar, click **Identity & Security**.
3. Choose **Identity** » **Domains** in the main menu
4. Ensure the **Compartment** drop-down (in the left sidebar under **List Scope**) is set to your root compartment, e.g.,  `your_tenancy_name (root)`.
5. Choose the default domain or create a custom domain if preferred (see **Creating an Identity Domain**).
6. On the domain’s **Overview** page, select **Users** from the left sidebar to view existing users.

#### **Create Kentik Export User**

1. ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(84).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A07Z&se=2026-05-12T09%3A46%3A07Z&sr=c&sp=r&sig=sOP0ssNBeruVGElCWMlWD4coq4aP2JktGbBe6i32mEg%3D)On the **Users** page, click **Create user** at the top of the list.
2. Enter the following user information:

   1. **First name**: Enter a first name for the cloud export user.
   2. **Last Name**: Provide a last name for the cloud export user.
   3. **Username / Email**: Enter an email address. If you prefer, uncheck the **Use the email address as the username** checkbox to enter both a username and an email.

      > **Note:** Additional settings in the **Create user** drawer are not necessary for creating a Kentik cloud export (see **Adding Users**).
3. Click **Create** to finalize the user setup and go to the new user’s page.

#### **Create an OCI User Group**

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(85).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A07Z&se=2026-05-12T09%3A46%3A07Z&sr=c&sp=r&sig=sOP0ssNBeruVGElCWMlWD4coq4aP2JktGbBe6i32mEg%3D)To assign your new cloud export user to a user group, follow these steps:

1. On the new user's page, click the breadcrumb segment for your chosen domain (e.g., "Default domain").
2. On the domain's page, use the left sidebar to go to the **Groups** page, where existing groups are listed.
3. Click **Create group** to open the **Create group** drawer.
4. Enter the group details:

   1. **Name**: Provide a unique name for the group.
   2. **Description**: Enter a brief description explaining the group’s purpose.
5. In the user list, check the box for the userto assign it to the group.

   > **Note:** Additional settings in the **Create Group** drawer are optional for creating a Kentik cloud export **Adding Users**.
6. Click **Create** to establish the group and go to the new group’s page.

#### **Configure an API Key**

To add the Kentik public key to your Kentik cloud export user, follow these steps.

1. **Navigate to API Keys**

   1. ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(86).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A07Z&se=2026-05-12T09%3A46%3A07Z&sr=c&sp=r&sig=sOP0ssNBeruVGElCWMlWD4coq4aP2JktGbBe6i32mEg%3D)On the new group's page, click the breadcrumb for your chosen domain (e.g., "Default domain").
   2. Use the left sidebar to navigate to the Users page.
   3. Click the link for the user created in **Create Kentik Export User**to open the user’s page.
   4. Click **API Keys** in the left sidebar to show the **API Keys** list.
2. **Add API Key**

   1. Click **Add API Key** at the top left of the list.
   2. Select **Paste a public key** and enter the public key downloaded from the Kentik portal in **Create a Cloud Export**.
   3. Click **Add** to save the key, which opens the **Configuration file preview** dialog.
   4. Ensure API key fingerprint is `d0:b4:75:ac:39:8a:90:b0:cf:ee:3e:ee:b9:0c:07:ff`
   5. Click **Copy** to save the **Configuration file preview** to your clipboard for later use.
   6. Close the dialog to return to the **API Keys** page for your Kentik cloud export user.

#### **Set up an Access Policy**

To set up an access policy for the Kentik cloud export tool, follow these steps:

1. ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(87).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A07Z&se=2026-05-12T09%3A46%3A07Z&sr=c&sp=r&sig=sOP0ssNBeruVGElCWMlWD4coq4aP2JktGbBe6i32mEg%3D)Click the menu icon at the top left of any OCI console page.
2. In the menu sidebar, select **Identity & Security**, then choose **Identity »** **Policies** on the main menu.
3. On the **Policies** page, click **Create Policy** to open the Create Policy page.
4. Enter a **Name** and **Description** for the policy.
5. Choose the compartment where the policy will apply. By default, this is the root compartment unless specified otherwise.
6. Turn on the **Show manual editor** switch to open a text input field.
7. Enter the policy statements below, replacing the placeholder with the group OCID:  
   `Define group groupRef as "groupId"  
    Allow group groupRef to READ all-resources in tenancy`
8. Click **Create** to save the policy and go to the new policy’s page.

> **Note**: This grants read-only access to all tenancy metadata. Adjust permissions as needed (see **OCI Policy Statements**).
