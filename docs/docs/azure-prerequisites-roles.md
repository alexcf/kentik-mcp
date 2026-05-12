# Azure Prerequisites & Roles

Source: https://kb.kentik.com/docs/azure-prerequisites-roles

---

Before **configuring your log exports**, use this guide to ensure you have the necessary Azure administrative roles, permissions, and subscription details ready.

### About Azure Roles

Integrating Kentik with Azure requires understanding two key Azure user roles:

* **Role to Enable Exporter**:

  + Typically, an Application Administrator role is used at the start of the setup process for each tenant.
  + The way you grant these permissions depends on whether you use the pre-integrated Kentik enterprise application or a Custom app registration.

    - **Kentik Enterprise Application**: When you select this option in Kentik, you'll grant permissions to a pre-existing enterprise application. This is typically done through the Microsoft Entra ID » **Enterprise applications** panel, where you search for the "Kentik NSG Flow Exporter" application and assign it the necessary roles to access your resources.
    - **Custom App Registration**: When you select this option, you are responsible for creating the application identity in your Azure tenant. The key difference is that you will grant the application permissions from the resource's perspective rather than from the application's perspective. After creating the app registration, you navigate to the specific Azure resource (e.g., a subscription or a resource group) and use the **Access control (IAM)** panel to assign a role to your custom app's Service Principal.
* **Role to Manage Exports**:

  + After permissions are granted by an Application Administrator, any user role can manage the setup and maintenance of log exports.

> **Notes:**
>
> * Use **Check Azure Role**to verify your current role.
> * For additional roles that may grant permissions for NSG Flow Exporter, see the Azure document **Available roles**.

### Check Azure Role

To verify your role in Azure, follow these steps:

1. On the Home page of your Azure portal, click **Microsoft Entra ID** in the sidebar at left.
2. In the resulting **Default Directory - Overview** page, find the **Manage** list (second sidebar from the left) and click **Roles and administrators**.
3. Your role will be indicated at the top of the main page (above the **Administrative roles** heading).

![Overview of Azure Active Directory roles and their descriptions for administrators.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/Azure-Roles-361h674w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A01Z&se=2026-05-12T09%3A35%3A01Z&sr=c&sp=r&sig=LhC4XjOGYpdw%2FF3qVRFtg6m3bOjmsqcMoBF6ZljYxO0%3D)

*Microsoft Entra ID roles and their descriptions for administrators*

### Role to Enable Exporter

To grant the necessary permissions, follow the steps below based on your selected authentication method (see **About Azure Roles**).

#### **Kentik Enterprise Application**

To grant permissions for Kentik’s NSG Flow Exporter, follow these steps if you are an Application Administrator or need to designate one:

1. In the **Administrative** roles list on the **Roles and administrators** page, click **Application Administrator**.
2. On the **Application administrator - Members** page, click **Add member** above the list of current administrators.

   > **Note*:*** This option will be grayed out if you are not an application administrator.
3. In the **Add Member** popup, find and select the user(s) you want to designate as application administrators.
4. Use the **Select** field to filter the list, then click a user to add it to the **Selected members** list.
5. Click the **Select** button to close the popup and return to the list of application administrators.

> **Note:** For more details on assigning administrator roles, refer to the Azure documentation **View and assign roles**.

#### **Custom Application Registration**

> **IMPORTANT**: **Tenant Type Selection**
>
> * When creating a custom app registration in the Azure portal, you must select **Accounts in this organizational directory only (Single tenant)**.
> * While Kentik is an external service, the authentication model requires that *you* own the service principal within your own tenant. You will then provide the credentials (Secret or Certificate) to Kentik so it can assume that identity to access your resources.

**Configuration Requirements**

To ensure the custom application can successfully export data to Kentik, verify the following settings:

* **App Ownership**: The customer (you) creates and manages the app registration principal.
* **Credential Sharing**: Once created, you must upload the application's secret or certificate into the Kentik portal under **Settings »** **Public Clouds**.
* **Permissions**:

  + **Subscription/Resource Group**: Assign the **Reader** role to your Service Principal Name (SPN). Ensure this Reader access also applies to any additional Subscriptions or Resource Groups included in your Enrichment Scope.
  + **Storage Account**: Assign the **Storage Blob Data Reader** (or "Reader and Data Access") role to the SPN.
  + **Firewall/Network**: If your storage account has firewall restrictions enabled, ensure you allow Kentik’s public IPs: `20.69.189.228` and `20.69.185.115`.

If you chose the custom app registration option, you will grant permissions to your own application's Service Principal. This process is managed from the resource's perspective in the Azure portal.

1. Navigate to the Azure resource you want to monitor (e.g., a specific resource group or your entire subscription).
2. In the resource's menu, select **Access control (IAM)**.
3. Click the **Add** button and then select **Add role assignment**.
4. Choose the required role (e.g., Reader).
5. On the Members tab, select **User, group, or service principal**.
6. In the **Select members** field, search for the name of your custom app registration.
7. Select your application from the list and click **Review + assign**.

### Required Azure Permissions

If you are manually configuring a **Custom App Registration** (Service Principal) instead of using the pre-integrated **Kentik Enterprise Application**, you must assign the following roles and permissions in the Azure Portal:

* **Subscription & Resource Group Access:** Assign the **Reader** role to your Service Principal Name (SPN) at the Subscription or Resource Group level where your flow logs reside.
* **Enrichment Scope Access:** If you plan to use Kentik's data enrichment features (like adding GeoIP or BGP data), you must also assign the **Reader** role to the SPN for *any additional* Subscriptions or Resource Groups included in your Enrichment Scope.
* **Storage Account Access:** Assign the **Reader and Data Access** (or **Storage Blob Data Reader**) role to the SPN for the specific storage account where your logs are being exported.
* **Storage Account Firewall (If applicable):** If your organization's security policy restricts storage account access, ensure your firewall rules are configured to allow Kentik to read from this account.

### Role to Manage Exports

After the Service Principal (either the Kentik NSG Flow Exporter or your Custom app registration) has been granted permissions by an Application Administrator, other users can manage the setup for collecting and storing flow logs and metadata. Here’s how to proceed:

1. Create a "Kentik" role with read-only permissions for metadata related to cloud naming, routing, ACL, and storage locations for flow logs and firewall logs.
2. Assign the Service Principal representing the application to this role.

   > **Note:** For comprehensive traffic path views, it’s common to assign the application at a Tenant level.
3. In the **Add Member** popup, locate the created role using the **Select** field to filter the list.
4. Click the “Kentik NSG Flow Log Exporter” or the name of your custom app registration to add it to the **Selected members** list.
5. Click **Select** to close the popup and return to the list of application administrators, which now includes the selected members.

### Find Azure Subscription ID

To locate the subscription ID for exporting flow logs from Azure:

1. On the Azure portal Home page, click **All Services** in the sidebar at left.
2. Click **General** on the left of the **All Services** page to filter the list.
3. Click on **Subscriptions** in the list of General services.
4. In the table at the bottom of the Subscriptions page, find and copy the 32-digit GUID in the **Subscription ID** row for the relevant subscription.

### Find Resource Group and Location

To identify the Resource Group and Location for creating a Storage Account:

1. On the Azure portal Home page, click **Virtual Machines** in the left sidebar.
2. In the resulting table, find the VM from which you want to export flow logs.
3. Copy and save the values in the **Resource Group** and **Location** columns.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/Azure-VMs-345h674w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A01Z&se=2026-05-12T09%3A35%3A01Z&sr=c&sp=r&sig=LhC4XjOGYpdw%2FF3qVRFtg6m3bOjmsqcMoBF6ZljYxO0%3D)

*List of Azure virtual machines showing their status and resource group*

## Azure Logging Checklist

Before continuing on to **Flow/Firewall Log Collection**, use this checklist to verify that your Azure environment is correctly configured to send telemetry to Kentik:

* **Storage Account****:** A dedicated storage account exists to house all logs.
* **Storage Account Firewall**: If firewall restrictions are enabled on your storage account, ensure public access is allowed for Kentik’s specific public IPs (`20.69.189.228` and `20.69.185.115`).
* **VNet Flow Logs:** Enabled and configured to send data to the storage account (Version 2 recommended).
* **Firewall Diagnostic Settings:** In Azure Firewall > Monitoring > Diagnostic Settings:

  + Created a setting targeting the same storage account used for flow logs.
  + Selected `azfwapplicationrule` and `azfwnetworkrule`.
* **Permissions & Roles**:

  + The **Kentik Service Principal** (Enterprise or Custom) has been created and assigned the necessary **Reader** roles at the Subscription or Resource Group level.
  + The Service Principal is specifically assigned the **Storage Blob Data Reader** role on the storage account containing the logs.
