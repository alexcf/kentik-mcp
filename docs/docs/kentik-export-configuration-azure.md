# Kentik Export Configuration

Source: https://kb.kentik.com/docs/kentik-export-configuration-azure

---

Once your **Azure environment is prepared**, set up a cloud export in the Kentik portal to authorize access and begin ingesting your network telemetry.

![Configuration settings for a Kentik cloud export, highlighting Azure observability features.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/KFAZ-cloud-export-config-provider(1).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A42Z&se=2026-05-12T09%3A31%3A42Z&sr=c&sp=r&sig=uBqDftcUXdHagv3wnLtIeixl%2BQClgvbtNMLupxTYyVI%3D)

*Configuration settings for a Kentik cloud export, highlighting Azure observability features.*

## **Configure a New Azure Cloud Export**

1. Go to Settings » **Public Clouds** in the main nav menu.
2. Click **Create Cloud Export** to start the configuration wizard.
3. Choose **Azure Cloud** under **Provider and Features**.
4. Under **Observability Features**, select the data types to collect:

   1. **Metadata collection** (required): Automatically selected.
   2. **Flow log collection**: Select to collect flow logs.
   3. **Help me configure my provider via Powershell**: Choose this to receive a Kentik-generated Powershell script for automatic configuration in Azure Cloud Shell, see **Automated Configuration**.
   4. **Firewall Collection**: Select to collect Azure firewall logs.
   5. **Cloud metrics history**: Select to collect Azure metrics with Kentik’s NMS.
5. Click the green arrow to proceed to the **Kentik Export Azure Configuration** step.
6. Select an **API Access** method by selecting either:

   1. ![Kentik Export Azure configuration interface showing API access and credential fields.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/KFAZ-custom-app-registration.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A42Z&se=2026-05-12T09%3A31%3A42Z&sr=c&sp=r&sig=uBqDftcUXdHagv3wnLtIeixl%2BQClgvbtNMLupxTYyVI%3D)**Kentik enterprise application****:** No additional information is collected in this section.
   2. **Custom app registration**: Enter the following information or choose from the **Saved App Registrations** (if any):

      1. **Application (client) ID**: Enter the ID for the custom application.
      2. **Directory (tenant) ID**: Enter the tenant ID.
      3. Select a credential from the **Kentik Credential** dropdown or click **Create New Secret** to create a new secret (see **Credentials Settings Dialogs**).
      4. Click **Save App Registration**.
7. ![Azure configuration settings including location, subscription ID, and storage account name.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/KFA-azure-new-export.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A42Z&se=2026-05-12T09%3A31%3A42Z&sr=c&sp=r&sig=uBqDftcUXdHagv3wnLtIeixl%2BQClgvbtNMLupxTYyVI%3D)Under **Define Azure Default Fields**, choose the location from the **Location** drop-down, as gathered in **Find Resource Group and Location**.
8. Enter the **Subscription ID** associated with the Azure directory containing the assets (see **Find Azure Subscription ID**).

   1. Click **Validate Service Principal** to authorize the Azure portal to create a Service Principal for Kentik's VNet Flow Exporter. Ensure your Azure role allows granting access to enterprise applications.

      > **IMPORTANT:** The **Validate Service Principal** step in the wizard is **only** required if you are authenticating via the pre-integrated **Kentik Enterprise Application**. If you are using a **Custom App Registration**, skip this button, as you have already manually authorized your Service Principal and assigned permissions in the Azure Portal.
9. In the **Resource Group to Monitor** field, enter the resource group gathered in **Find Resource Group and Location**.

   1. Click **Verify Access** to ensure permission has been granted to the resource group for all required APIs.
10. In the **Define Storage Account** section, enter a unique **Storage Account Name** for the storage account where logs will be exported.

    > **Notes**:
    >
    > * Kentik must access your storage account from the following public Azure IPs: `20.69.189.228` and `20.69.185.115`.
    > * If you have a storage account in the **WestUS2** region, use these IPs instead: `51.8.214.254` and `172.171.46.16` and contact your account team.
11. Under **Sampling**, choose from:

    1. **Sampling Rate**: After selecting this option, enter a sampling rate in the Sampling Rate field. The value must be between 2 and 2000.
    2. **Unsampled**: Select this option if you want all flow logs to be sent without sampling.
12. Click the green arrow to proceed.
13. Optionally, configure the **Azure Metadata Enrichment Scope**, see **Define Enrichment Scope**.
14. Click the green arrow to proceed
15. Enter the cloud export name/description or accept the defaults.
16. Select the appropriate Kentik billing plan for the cloud export from the **Billing Plan** dropdown.
17. Click **Save** to finalize the cloud export and return to the Public Clouds page, where the new export will be listed.

### Define Enrichment ScopeInstructions for defining network scope for Azure metadata enrichment with subscription IDs.

To optionally configure the network scope for data enrichment, allowing Kentik to enhance your Azure flow data with additional information such as GeoIP and BGP, follow these steps:

1. In the **Subscription IDs** box, paste or drag a file containing comma-delimited subscription IDs. This will allow Kentik to view the resource groups associated with these subscriptions.

   > **Note:** Ensure that the Subscription ID is not in use by any other company.
2. A list of the entered **Subscription IDs** will populate the page, each with an **All Resource Groups** drop-down listing all resource groups within that subscription. Select the desired resource groups for enrichment.

> **Notes:**
>
> * The All Resource Groups dropdown is disabled if the subscription is in use by any other user.
> * The Subscription IDs box will indicate the number of valid (green checkmark) and invalid (red exclamation) subscriptions.
> * You can repeat the process to add more subscriptions to the list.
> * Click Remove to delete subscriptions from the list.

## Using Your Cloud Export

Once the setup process is complete, you can view and utilize your cloud export in Kentik:

* **Cloud Exports List**:

  + Go to Settings » **Public Clouds** to see the updated list of cloud exports.
  + A new cloud export will be listed, representing the VNets or NSGs whose logs are pulled from the specified subscription.
* **Devices Column**:

  + Each VNet/NSG sending flow logs is listed as a cloud device.
  + Devices are named after their respective VPC subnet.
  + These names can be used as group-by and filter values in Kentik queries using the Device Name dimension.
* **Metadata and Mapping**:

  + The collected metadata, for example for subnets and gateways, enables Kentik to automatically map and visualize the topology of your Azure resources in the **Kentik Map**.

### Historical Metrics Collection

Kentik allows for the collection of historical metrics for Azure. Along with real-time metrics, this data can provide a comprehensive view of performance trends and patterns over time.

The collection of historical metrics must be enabled in the configuration of the cloud export. Once enabled, you can view the metrics with the following steps:

1. From the Kentik portal main menu, go to Settings » Network Monitoring System » **Metrics Explorer**.
2. In the **Measurement** pane of the **Metrics Explorer Query Sidebar**, select a measurement starting with "/cloud/Azure/", then select from the available metrics using the dropdown. Click **Run Query** to execute the query.

## Edit an Existing Cloud Export

![Azure Cloud configuration with subscription ID and resource group details.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/KFA-azure-edit-export.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A42Z&se=2026-05-12T09%3A31%3A42Z&sr=c&sp=r&sig=uBqDftcUXdHagv3wnLtIeixl%2BQClgvbtNMLupxTYyVI%3D)You can modify the settings of an existing Azure cloud export (such as enabling firewall logs, adjusting sampling rates, or changing the enrichment scope) at any time by navigating to **Settings »** **Public Clouds** and clicking on your export.

When editing, please note the following regarding the **Authorize** button:

* **No Need to Re-Authorize:** If your export is already successfully collecting data, you **do not** need to click the Authorize button again to save your changes. Simply make your adjustments and click **Save** at the bottom of the dialog.
* **Custom App Registrations:** If you originally configured your export using a Custom App Registration, the UI may default to showing the "Kentik enterprise application" tab when you reopen it. **Do not** click the **Authorize** button, as it does not apply to your custom configuration.
