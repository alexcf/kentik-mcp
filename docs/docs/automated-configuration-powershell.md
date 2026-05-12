# Automated Configuration (PowerShell)

Source: https://kb.kentik.com/docs/automated-configuration-powershell

---

If you opted for **Help me configure my provider via Powershell** in the **Kentik Export Configuration**, follow these steps. You'll use both the Kentik and Azure portals to complete the automated setup of an Azure cloud export.

> **Notes**:
>
> * As of February 2025, Kentik’s Powershell script can only be used for configuring VNet flow logs.
> * For manual configuration guidance, see the Microsoft Azure document **Tutorial: Log network traffic**.

#### Generate PowerShell Script

To generate and customize a Powershell script from the **Azure Powershell** section of the Kentik wizard, follow these steps:

1. Verify Script Values:

   1. Check that the values for subscription ID, resource group, storage account, and location are correct.
   2. If any values are incorrect, click the gray left arrow until you return to the first wizard step and correct the entered information.
2. Click **Copy to Clipboard** at the top right of the textbox to copy the script.
3. Optional: To change the log retention duration from the default two days:

   1. Paste the script into a text editor.
   2. Locate the `RetentionInDays` argument in the declaration of the `$ret` variable (line 220).
   3. Modify the value to an integer representing whole days, between 1 and 365.
   4. Copy the edited script to the clipboard.

#### Configure Using PowerShell

To set up Azure log export using PowerShell in the Azure portal, follow these steps:

1. Navigate to your Azure portal and log in.
2. In the main Azure navbar, click the PowerShell icon (`>_`).
3. Click PowerShell in the popup and wait for PowerShell to initialize.
4. Once initialization is complete, type code at the prompt to open the code editor.
5. Paste and Save Script:

   1. Paste the script from **Generate PowerShell Script** into the code editor.
   2. Right-click anywhere in the editor and select **Save** to open the **Save new file** dialog.
   3. Enter a name for the script with a `.ps1` extension (e.g., `MyKentikScript.ps1`).
   4. Click **Save**.
6. Choose **Close** from the script editor menu to close the editor.
7. From the PowerShell prompt, enter the full path to the script (e.g., `/home/user_name/MyKentikScript.ps1`).

   > **Note**: The `user_name` for an Azure script file path is the first word of your full user name (e.g., "Sallie Mae" becomes `sallie`).
8. Execute the script (see **Azure Script Operations**). Upon completion, you'll see a confirming message with the details you entered for subscription ID, resource group, location, and storage account.

> **Note**: PowerShell instances are ephemeral. If you time out or lose connection, restart the process from step #2.

#### Azure Script Operations

The configuration script for Azure log export performs the following operations based on your settings in the Kentik wizard:

* **Confirmations**:

  + Confirms the existence of the provided subscription ID, location, and resource group.
  + Confirms that a service principal has been created for the Kentik VNet Flow Exporter application.
  + Confirms that VNet Flow Exporter has been granted "Reader" access to the specified resource group and to enrich subscriptions and resource groups.
* **Storage Account Setup**:

  + Checks if a storage account exists for the specified name, resource group, and location and creates one if not.
  + Grants VNet Flow Exporter "Contributor" access to the storage account.
* **Network Watcher and Insights Registration**:

  + Confirms that a Network Watcher exists for the specified resource group and location.
  + Confirms that the specified subscription is registered with Microsoft Insights, the resource provider namespace for Azure Monitor.
* **VNet Configuration**:

  + Builds a list of VNets in the specified resource group and location.
  + Enables flow logs for each found VNet.
