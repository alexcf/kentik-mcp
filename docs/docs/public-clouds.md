# Public Clouds

Source: https://kb.kentik.com/docs/public-clouds

---

This article covers the **Public Clouds** page in the Kentik portal.

> **Notes:**
>
> * For an overview of public clouds in Kentik, see **Cloud Overview**.
> * For step-by-step procedures by cloud provider, see the articles in **Cloud Provider Configuration**.

![Graph showing public cloud utilization and detailed cloud export configurations.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/PC-main.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A06Z&se=2026-05-12T09%3A42%3A06Z&sr=c&sp=r&sig=pqm8Mc2ZeXKWJpo3QnFUY3iztdOO7CombEQ1nXnnvnM%3D)

*Manage your public cloud resources in the Kentik portal using the Public Clouds page.*

## Public Clouds Page

The Public Clouds page (Settings » **Public Clouds**) is where you manage "cloud exports", which are collections of data such as flow logs, metadata, and metrics gathered by Kentik from your public cloud resources.

> **IMPORTANT:** Administrator-level permissions (see **User Level**) are required to add, edit, or remove cloud exports.

### Public Clouds Page UI

The Public Clouds page features the following UI elements:

* **Create Cloud Export** (button): Click to configure a new cloud export using the **Cloud Export Configuration Page**.
* **Utilization** (pane): Displays the cloud provider flow logs per second (FPS) ingested by Kentik as compared to the Max FPS allowed by your CloudPak license (see **Licenses**). Yellow/red ranges indicate Provider FPS values exceeding the limit, suggesting Kentik may be downsampling to comply.

  + Select a time window, either 1 (default), 7, or 30 days, using the dropdown.
  + Click a tab to select from the following charts:

    - **Raw Traffic**: The total volume of flow data before processing.
    - **Received Traffic**: The volume of traffic successfully ingested by Kentik.
    - **Stored Traffic**: The volume of traffic data committed to long-term storage.

  ![Graph showing traffic utilization over seven days with highlighted maximum flow levels.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/PC-main-utilization-graphs.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A06Z&se=2026-05-12T09%3A42%3A06Z&sr=c&sp=r&sig=pqm8Mc2ZeXKWJpo3QnFUY3iztdOO7CombEQ1nXnnvnM%3D)

  > **Note**: Filtering the Cloud Exports list by Provider also auto-updates the Utilization chart (see **Filters** below).
* **Cloud Config Status** (pane): Lists the health status (e.g., warning, critical, etc.) of resources in your cloud providers. For each provider:

  + Expand a provider to view status by region.
  + Click **View Details** to go to the **Configuration Status Page** for a provider.
* **Show/Hide Filters** (filter icon): A button to expand/collapse the Filters pane.
* **Search** (field): Enter text to filter the **Cloud Exports List** by Provider, Name, Properties, or Devices.

  + Remove filters by clicking the **X** in the lozenge.
  + Filters added via the **Filters** pane also appear in this field.
* **Cloud Exports** (table): Lists your organization’s cloud exports (see **Cloud Exports List**).
* **Filters** (pane): Checkboxes to filter the Cloud Exports list and the Utilization chart by provider. Expand/collapse the pane with the **Show/Hide Filters** button.

### Cloud Exports List

The Cloud Exports list provides details about your organization's cloud exports (see **Cloud Exports and Devices**). Click on any column heading to sort the list.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(574).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A06Z&se=2026-05-12T09%3A42%3A06Z&sr=c&sp=r&sig=pqm8Mc2ZeXKWJpo3QnFUY3iztdOO7CombEQ1nXnnvnM%3D)

*The Cloud Exports list provides basic information about each cloud export and access to actions and additional details.*

The columns include:

* **Provider**: The cloud provider from which Kentik pulls flow logs (AWS, Azure, etc.).
* **Status**: The cloud export status (e.g., Active, Error, see **Cloud Export Status Icons**).
* **Name**: The specified name for the cloud export. Click to access the **Cloud Details Page**.
* **Properties**: A summary of settings, varying by provider:

  + **AWS**: Path and Role
  + **GCP**: Project and Subscription
  + **Azure**: Resource Group and Storage Account
  + **OCI**: Tenancy ID, Compartment ID, User ID, and Default Region
* **Devices**: Auto-derived devices by Kentik, varying by provider (see **Cloud Exports and Devices**).
* **FPS Sampled**: The peak flow records ingested by Kentik over the last 6 hours.
* **Edit**: Opens the configuration dialog for the cloud export (see **Cloud Export Settings**).
* **Refresh**: Retrieves the latest flow logs for the cloud export.
* **Remove**: Opens a confirmation dialog to remove the cloud export.

#### Cloud Export Status Icons

These icons indicate the status of a cloud export:

* **Start** (blue circle): Kentik is setting up the cloud export based on provided information.
* **Pending** (orange clock): Setup is in progress based on provided information.
* **Error** (red exclamation): A configuration or connection issue is preventing access to flow logs.
* **OK** (green checkmark): Flow logs are successfully being ingested.
* **Disabled** (cloud with strikethrough): Export is disabled in the settings.
* **Halt** (red hand): Export is halted. Contact **Customer Care** for assistance.

## Cloud Details Page

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(575).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A06Z&se=2026-05-12T09%3A42%3A06Z&sr=c&sp=r&sig=pqm8Mc2ZeXKWJpo3QnFUY3iztdOO7CombEQ1nXnnvnM%3D)

*The Cloud Details page offers access to the properties and settings of a specific cloud export.*

The Cloud Details page includes the following UI elements:

* **Cloud provider** (logo): Displays the logo of the provider hosting the cloud export.
* **Cloud name**: The name of the cloud export.
* **Configure**: A button to open the configuration dialog for the cloud export (see **Cloud Export Settings**).
* **Cloud Summary**: Displays information about the cloud export, including the assigned plan and provider-specific details (see **Cloud Summary Information**).
* **Devices list**: A table listing the devices associated with the cloud export (see **Cloud Devices List**).

### Cloud Summary Information

Cloud export summary provides the following information, varying by provider:

* **Export ID**: A unique ID assigned to the cloud export by Kentik.
* **Billing Plan**: The name of the Kentik plan to which the cloud is assigned (see **About Plans**).
* **Plan ID**: The ID of the Kentik plan associated with the cloud export.
* **S3 Bucket Region (AWS)**: The AWS region of the S3 bucket for flow logs.
* **S3 Bucket Name (AWS)**: The name of the S3 bucket (not the full ARN).
* **IAM Role ARN (AWS)**: The full ARN of the role allowing Kentik access to AWS resources.
* **Delete After Read (AWS)**: Setting to auto-delete flow logs in AWS post-ingestion by Kentik.
* **Subscription ID (Azure)**: The ID of the Azure instance for exporting flow logs via Kentik's NSG Flow Exporter.
* **Resource Group (Azure)**: The group of Azure resources for which flow logs are collected.
* **Location (Azure)**: The location of the Azure resources for flow logs collection.
* **Storage Account (Azure)**: The name of the Azure storage account for exporting logs.
* **Project (GCP)**: The name of the GCP project with the Pub/Sub topic for flow log export.
* **Subscription (GCP)**: The name of the subscription for Kentik’s access to the Pub/Sub topic.
* **Description (AWS, GCP, OCI)**: The description provided during the creation or editing of the cloud export, if available.

### Cloud Devices List

The Devices list is a sortable table with these columns:

* **Status**: The device’s flow status:

  + Ok (green checkmark): Flow logs successfully ingested
  + Error (orange exclamation): A configuration/connection issue is preventing access to flow logs
  + Pending (orange clock): Device setup is in progress
* **Name**: The device name with ID in parentheses.
* **FPS Sampled**: The peak flow records ingested by Kentik in the last 6 hours.

## Cloud Export Configuration Page

The Cloud Export Configuration Page, accessible via Settings » Public Clouds » **Create Cloud Export**, allows you to configure new cloud exports.

> **Notes**:
>
> * For step-by-step instructions on creating cloud exports by provider, see **Kentik for AWS**, **Kentik for GCP**, **Kentik for Azure**, and **Kentik for OCI**.
> * For descriptions of all cloud export settings for all providers, see **Cloud Export Settings**.

The Cloud Export Configuration page is a multi-step setup wizard with the following UI elements:

* **Cancel (button)**: Cancels the new cloud export.
* **Save (button)**: Saves the new cloud export. Active only after all required fields in all wizard steps are completed.
* **Summary Pane**: Lists the wizard steps to complete based on cloud provider selection:

  + **Navigation**: Click a step to jump to it (active only once you've started that step).
  + **Warnings**: Hover on a step for warnings about missing required information.
* **Settings Pane**: Contains the detailed settings for each wizard step, e.g., "Provider and Features".
* **Forward/Backward** **(buttons)**: Arrow buttons to go forward/backward in the wizard. Forward buttons become active as you complete required fields on a particular wizard screen.

### Wizard Summary

The steps of the Cloud Export Configuration wizard are summarized below:

* **Provider and Features**: Choose a provider and select observability features to configure.
* **Kentik Export Configuration**: Configure roles, regions, project IDs, etc. for the provider.
* **Azure Metadata Enrichment Scope** (Azure only): Optionally specify Azure subscription IDs and resource groups.
* **Final Step**: Choose a cloud export name, description, and billing plan.

## Configuration Status Page

The Configuration Status pages for cloud providers are accessed via the **View Details** links in the Cloud Config Status pane.

> **Note:** Pages are similar across providers, with differences noted in the topics below.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(576).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A06Z&se=2026-05-12T09%3A42%3A06Z&sr=c&sp=r&sig=pqm8Mc2ZeXKWJpo3QnFUY3iztdOO7CombEQ1nXnnvnM%3D)

*The Configuration Status page for each provider shows cloud exports with issues.*

These status pages assist in cloud troubleshooting by identifying issues configuration with your organization's cloud exports. Each cloud provider has a dedicated page showing results from Kentik’s cloud configuration checker, which evaluates:

* Access to required API endpoints
* Receipt of flow logs

Results are displayed in tables for each region (AWS, Azure, and OCI) or project (GCP) where you have cloud resources.

> **Note**: The checker looks at each AWS account (ARN), Azure account (Resource ID), GCP account (Resource ID), and OCI account (Tenancy ID) for which your organization has a cloud export.

### Cloud Config Status UI

The Configuration Status page includes the following UI elements:

* **Refresh**: Click to refresh the statuses shown on the page.
* **Page Title**: Shows you which provider for which you’re viewing status, e.g., “AWS Configuration Status”.
* **Filter**: Enter text to filter rows by account/subscription, export, or entity ID.
* **Configuration Status**: Tables for each region (AWS, Azure, and OCI) or project (GCP) with issues reported by the configuration checker (see **Configuration Status List**).

### Configuration Status List

This table lists cloud export statuses by region (AWS, Azure, and OCI) or project (GCP). Click a row to view the **Config Status Details** sidebar for the account, subscription, or project.

The columns include:

* **Account ID (AWS)**: The AWS account ID.
* **Subscription ID (Azure, GCP)**: The Azure Subscription ID or GCP Project Number.
* **Compartment ID (OCI)**: The OCI Compartment ID.
* **API**: A status icon for API endpoint access.
* **Flow**: A status icon for flow log receipt.
* **Exports**: A status icon for operational issues on Kentik’s side.
* **Export IDs**: The unique IDs for cloud exports. Click to view the **Cloud Details Page** for the export.

> **Note**: Hover over non-green checkmarks for status descriptions.

### Config Status Details

The **Config Status Details** sidebar appears when you click on a row in the **Configuration Status List**, providing details on the operations needed for the cloud export’s success.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(577).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A06Z&se=2026-05-12T09%3A42%3A06Z&sr=c&sp=r&sig=pqm8Mc2ZeXKWJpo3QnFUY3iztdOO7CombEQ1nXnnvnM%3D)The drawer includes the following elements, varying by provider:

* **Owner ID (AWS only)**: The AWS account ID.
* **Subscription ID** **(Azure, GCP)**: The Azure subscription ID or GCP project account.
* **Compartment ID** **(OCI)**: The OCI account ID.
* **Region (AWS, Azure, OCI)**: The cloud provider region.
* **GCP Project (GCP)**: The name of the GCP project.
* **Subnets with Sampling (GCP)**: The number of subnets with sampling.
* **Average Sampling Rate (GCP)**: The average sampling rate.
* **API Access Status**: A table with:

  + **Endpoint**: API endpoints, grouped into collapsible sections by type, with a status summary for each type
  + **Can Access**: Access status icon.
* **Flow Access Status (AWS)**: A table showing status for individual VPCs (see **Flow Access Status**).
* **Cloud Export Status**: A table with:

  + **Status**: Export status icon (hover for details if not a green checkmark).
  + **Export ID**: The export ID.
  + **Name**: The export name (click to view the **Cloud Details Page**).

#### Flow Access Status

The Flow Access Status table includes:

* **Entity ID**: The IDs of the AWS VPCs in the account.

  + **Has Flow (True):** Green checkmark; bucket count shown in blue. Hover to see bucket details.
  + **Has Flow (False):** Red icon; "Flow logs not configured." in orange.
* **Has Flow**: An icon showing if Kentik receives logs from the VPC.
* **Open in Data Explorer**: Opens **Data Explorer** in a new tab, filtering results for this VPC.

## Cloud Export Settings

To edit a cloud export in the Kentik portal, use the Configure Your Cloud dialog.

> **Note:** The Configure Your Cloud dialog is for modifying existing cloud exports. For new exports, click **Create Cloud Export**.

This dialog allows modification of cloud export information, and can be accessed from:

* **Public Clouds Page**: Click the edit icon in the **Cloud Exports List**.
* **Cloud Details Page**: Click an export name to open the cloud details page, then use the **Configure** button (see **Cloud Details Page UI**).

Fields in the dialog vary by provider, with some common to all (see **Common Cloud Settings**).

### Cloud Export Settings UI

The **Cloud Export Settings** dialog for all providers includes:

* **Close** (button): Click the **X** in the upper right corner to close the dialog without saving changes.
* **Settings** (fields): The common and provider-specific fields containing configuration settings for the cloud export (see topics below).
* **Cancel** (button): Cancel and exit the dialog without saving changes.
* **Save** (button): Save changes to Cloud settings and exit the dialog.

### Common Cloud Settings

The Cloud Export Settings dialog includes the following fields for all providers:

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/PC-Settings_dialog_common-252h418w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A06Z&se=2026-05-12T09%3A42%3A06Z&sr=c&sp=r&sig=pqm8Mc2ZeXKWJpo3QnFUY3iztdOO7CombEQ1nXnnvnM%3D)

* **Name** (required): Enter a descriptive name for the cloud export.
* **Description** (optional): Provide a brief description of the export.
* **Billing Plan** (required): Select the Kentik billing plan for the cloud export (see **About Plans**).
* **Enabled**: Toggle the switch to enable/disable the export.

  > **Note:** Disabling a cloud export does not affect the cloud provider. Flow logs will continue until stopped on the provider’s side.

### AWS Provider Settings

Kentik ingests flow logs from AWS VPCs, subnets, or interfaces via an S3 bucket (see **AWS Logging Setup Overview**). To set up this connection, provide the following:

* ![Edit dialog settings for AWS cloud export](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/PC-edit-AWS-dialog(1).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A06Z&se=2026-05-12T09%3A42%3A06Z&sr=c&sp=r&sig=pqm8Mc2ZeXKWJpo3QnFUY3iztdOO7CombEQ1nXnnvnM%3D)

  **S3 Bucket Region** (required): Select the region of your S3 bucket (see **Create an S3 Bucket**).
* **Metadata only**: Enable this to only collect metadata and not flow logs (hides S3 Bucket settings).
* **S3 Bucket Name** (required): Enter the bucket name, not the full ARN.
* **Optional: Additional S3 bucket Options**: Expand this section to specify additional options for the S3 bucket.

  + **S3 Bucket Prefix**: Enter a prefix to refine the selection of objects in the bucket (e.g., folder/subfolder name)
  + **Delete logs**: Enable this to delete logs from the bucket after export.
* **AWS Role** (required): Enter the full ARN of the IAM role allowing Kentik to access the S3 bucket.
* **Organization Role**: Check the box if you are using an AWS Organizations role, allowing Kentik access to all accounts under the organization.
* **Cloud Metrics with NMS (Beta)**: Enable this to collect metrics from AWS CloudWatch to support NMS policy alerts and historical telemetry analysis.
* **Delete After Read**: Enable this to automatically delete flow logs from your S3 bucket after they have been processed by Kentik; disable to manage deletion yourself.
* **Additional Metadata Roles**: Expand this section to configure additional roles/accounts/regions for collecting metadata.

  + **Secondary AWS Accounts**: Enter a comma-separated list of account IDs where you want to collect metadata
  + **Regions Used Across Secondary Accounts**: Select the regions from which you want to collect metadata for the specified secondary accounts.
  + **Role Suffix**: Enter a suffix to append to the primary ARN to generate ARNs for the specified secondary accounts and regions.
  + **Verify**: Click to validate the secondary roles allow Kentik access to the metadata
* **Sampling**: Set the sampling rate for the cloud export (see **Cloud Export Sampling**).

#### Cloud Export Sampling

These settings determine the sampling rate for a cloud export:

* **Sampling Type**:

  + **Legacy** (AWS only): Randomly samples up to 10k flows per file
  + **Sampling Rate**: Uses the specified rate in the **Sampling Rate** field
  + **Unsampled**: No sampling (all flows are processed)
* **Sampling Rate** (applicable when **Sampling Type** is “Sampling Rate”):

  + Enter a number from 2 to 2000, e.g., a sampling rate of “2” means that every other record will be sampled.

### Azure Provider Settings

Kentik uses an Azure "storage account" to pull flow logs for cloud exports. The Kentik enterprise application or a custom application registration (Service Principal Name) accesses this account to forward logs to KDE. Configure the following settings:

* ![Edit dialog showing details of an Azure VNET flow log export.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/KFAZ-edit-export-dialog(1).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A06Z&se=2026-05-12T09%3A42%3A06Z&sr=c&sp=r&sig=pqm8Mc2ZeXKWJpo3QnFUY3iztdOO7CombEQ1nXnnvnM%3D)

  **Custom App Registration**: These settings only apply when using a custom app registration. Click the corresponding tab to edit the settings:

  + **Application (client) ID**: Enter the ID for the custom application.
  + **Directory (tenant) ID**: Enter the tenant ID.
  + Select a credential from the **Kentik Credential** dropdown or click **Create New Secret** to create a new secret (see **Credentials Settings Dialogs**).
  + Click **Save App Registration** to finalize the settings changes.
* **Subscription ID** (required): Enter the Subscription ID of the Azure instance for exporting flow logs (see **Find Azure Subscription ID**) via Kentik’s NSG Flow Exporter application.
* **Authorize**: Click to grant the necessary permissions to the NSG Flow Exporter. Ensure your Azure role (e.g., Application Administrator) permits this.
* **Resource Group to Monitor**: Enter the resource group name for generating flow logs (see **Find Resource Group and Location**).
* **Location (Required)**: Select the Azure region where the resources are located.
* **Firewall Log Collection**: Enable this to collect firewall logs from your Azure resources.
* **Cloud Metrics with NMS (Beta)**: Enable this to collect metrics from Azure Monitor, useful for performance monitoring.
* **Flow Log Collection**: Enable this to collect flow logs from your Azure resources.

  **Storage Account Name** (required): Enter a globally unique name for the storage account that will store the exported flow logs.
* **Sampling**: Configure the sampling rate for the cloud export (see **Cloud Export Sampling**).
* **Enrichment Scope**: Click **Edit** to open the Enrichment Scope dialog in which to paste or drag a file of comma-delimited Subscription IDs. This defines which Azure resources will be exported and how their metadata will be enriched.
* **Configure Manually**: Opens a dialog of manual configuration steps to be completed in the Azure portal (see **Choose Configuration Method**).
* **Configure Using PowerShell**: Opens theLogging Configuration Script dialog with a Kentik-generated script for automatically configuring the flow log export (see **Generate PowerShell Script**).
* **Validate Configuration**: Click **Validate** to start the validation process for the flow log export setup:

  + Validation can take up to one hour.
  + The export’s status will show as “Pending” on the Public Clouds Page until Kentik completes registration.

### GCP Provider Settings

Kentik ingests flow logs from GCP by subscribing to a Pub/Sub topic that is publishing from one or more subnets/VPCs (see **GCP Process Overview**). Provide the following information to establish this connection:

* ![Edit dialog settings for GCP cloud export.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/PC-edit-gcp-dialog(1).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A06Z&se=2026-05-12T09%3A42%3A06Z&sr=c&sp=r&sig=pqm8Mc2ZeXKWJpo3QnFUY3iztdOO7CombEQ1nXnnvnM%3D)

  **Project** (required): Enter the ID of the GCP project you want to monitor (see **Create a New Topic**).
* **Metadata only**: Enable this if you only want to collect metadata and not flow logs.
* **Subscription** (required): Enter the name of the Pub/Sub subscription that Kentik will use to pull flow logs (see **Create a Pull Subscription**).
* **Average Sampling Rate**: Displays the average flow log sampling rate configured in the subnets in your project. Configured in GCP and applied individually per subnet.
* **Cloud Metrics with NMS (Beta)**: Enable this to collect metrics from GCP Monitoring, useful for historical analysis and setting up NMS policy alerts.

### OCI Provider Settings

Kentik retrieves flow logs from OCI using the OCI object storage service. This service organizes resources within an OCI tenancy by region and compartment. Flow logs are collected via a service connector and stored in an OCI bucket.

#### OCI Details Pane

![Edit dialog settings for OCI cloud export](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/PC-edit-oci-dialog.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A06Z&se=2026-05-12T09%3A42%3A06Z&sr=c&sp=r&sig=pqm8Mc2ZeXKWJpo3QnFUY3iztdOO7CombEQ1nXnnvnM%3D)To set up an OCI cloud export, gather the following from the OCI console (see **OCI Console Info for Portal**):

* **Tenancy ID** (required): Enter your OCID tenancy ID, found on the OCI Portal Tenancy Page.
* **Compartment ID**: Enter the OCID of the compartment you want to monitor, found on the OCI Portal Compartment Page.

  + **Load Compartments**: Click to automatically load all compartments within your tenancy.
* **User ID** (required): Enter the OCID of the user that Kentik will use to access your tenancy.
* **OCI Default Region** (required): Select the region for your tenancy.
* **Verify**: Click to verify that Kentik can access your tenancy with the provided credentials. A lozenge displays the result:

  + **Valid**: Credentials verified.
  + **Not Authenticated**: Credentials invalid.
* **Flow Log Collection**: Enable this to collect flow logs, revealing additional **Flow Log Collection Controls**.

> **Notes:**
>
> * A detailed setup guide is available in **Create a Kentik Cloud Export**.
> * "OCID" is the unique ID for OCI resources like tenancy, compartment, user, bucket, connector, etc.

#### Flow Log Collection Controls

These fields appear when the Flow Log Collection switch is enabled. Gather the following from the OCI console (see **OCI Console Info for Portal**):

* **Bucket Name (Required)**: Enter the name of the OCI Object Storage bucket for storing flow logs.
* **Bucket Namespace (Required)**: Enter the namespace of the bucket.
* **Service Connector OCID (Required)**: Enter the OCID of the service connector used to export flow logs to the bucket.
* **Flow Object Name Prefix**: Enter a prefix to filter flow log objects in the bucket.
* **Verify Flow Logs Bucket Access**: Click to verify Kentik’s access to the bucket. Results are shown in a notification lozenge.
* **Average Sampling Rate**: Displays the average flow log sampling rate configured for capture filters in your tenancy. This is set in OCI and applied per capture filter.

> **Note:** A detailed setup guide is available in **Flow Log Collection**.

## Manage Cloud Exports

To add or edit a Kentik cloud export, follow these steps:

### Create a Cloud Export

To add a cloud export to Kentik:

1. Go to Settings » **Public Clouds** from the main Kentik menu.
2. On the Public Clouds page, click **Create Cloud Export**.
3. Follow the step-by-step procedure for the cloud provider:

   * **Kentik for AWS**
   * **Kentik for GCP**
   * **Kentik for Azure**
   * **Kentik for OCI**

### Edit a Cloud Export

To edit an existing cloud export:

1. Choose Settings » **Public Clouds** from the main Kentik menu.
2. In the **Cloud Exports** list, click the **Edit** icon for the cloud you wish to modify.
3. Modify fields as needed:

   1. For common fields, see **Common Cloud Settings**.
   2. For provider-specific fields, see **Cloud Export Settings**.
4. Click **Save** to apply changes.

> **Note**: To remove a cloud export, click **Remove** (trash icon) in the **Cloud Exports** list.

The process of reducing the number of flow records by sampling, where only one flow record is exported for every X flows.
