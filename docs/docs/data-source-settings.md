# Data Source Settings

Source: https://kb.kentik.com/docs/data-source-settings

---

This article covers data sources settings (e.g., devices or clouds) for queries in **Data Explorer** and dashboard panels in the Kentik portal.

## About Data Sources

Data sources are the devices and cloud resources from which Kentik receives network data (e.g., NetFlow, sFlow, BGP, SNMP, flow logs, etc.). Devices fall into two broad categories, routers and hosts, each of which includes many types of devices, such as routers, firewalls, and hosts. For further detail, see **Supported Device Types**.

## Data Sources Pane

The **Data Sources** pane is found in the **Query** sidebar in Data Explorer and on the **Query** tab of the settings dialogs for panels in dashboards (see **View Panel Query Tab**). The pane includes the following UI elements:

* **Data sources** list: A list of the data sources that are currently included in the query:

  + The list defaults to All Data Sources. To change this setting, click to open the Data Sources dialog.
  + An individual item in the list (except All Data Sources) can be removed by clicking the **X** at right.
* **Edit Data Sources**: A button that opens the **Data Sources Dialog**.

## Data Sources Dialog

The Data Sources Dialog includes the following settings and controls:

* **Filter field**: Filters the list of available data sources to show only data sources whose name, label, or site contains the entered string.
* **Available data sources**: A list of available data sources and logical data source groups from which you can choose data sources to add to the **Selected Data Sources** list (see **Available Data Sources List**).
* **Selected Data Sources list**: A list of the data sources and logical data source groups that have been chosen from the sidebar for inclusion in the query. An “X” at the right of each item allows you to remove an item from the list.
* **Cancel button**: Exit the dialog without saving changes.
* **Save button**: Save changes to the list of selected data sources and exit the dialog.

### Available Data Sources List

The list of available data sources and logical data source groups, which is displayed in the sidebar of the Data Sources dialog, is categorized as follows:

* **Use All Data Sources**: Select all of your organization's Kentik-registered data sources.
* **Types**: Select all data sources of a given type: routers, hosts, clouds, etc. (see **Supported Device Types**).
* **Labels**: Custom free-form groupings of data sources (see **About Labels**).
* **Sites**: Sites represent user-defined physical locations, such as a data center (see **About Sites**).
* **Data sources**: Individual data sources that are registered with Kentik.

  > **Note:** If an individual data source is part of a group (e.g., label or site) that has already been selected then the data source will not appear in the **Available Devices** list.

The following additional functionality is supported for the **Labels** and **Sites** categories of the **Available Devices** list:

* **Clear button** (shown only when the Filter field is empty): Removes all items of the category (**Label**s or **Sites**) from the **Selected Devices** list.
* **Add All button** (shown only when the **Filter** field is not empty): Adds all items of the category (**Labels** or **Sites**) that are currently shown in the **Available Devices** list to the **Selected Devices** list.

### Device Type Selector

The device type selector adds an entire category of devices to the **Selected Devices** list. The following categories can be added:

* **All**: Selects all Kentik-registered devices in your organization.
* **None**: Clears the selection of all devices in your organization.
* **Kentik Host Agent**: Selects host devices in your organization that represent Kentik host agents, which enable Kentik to derive DNS/WWW dimensions and NPM metrics (see **Universal Agents**).
* **Routers**: Selects all routers, switches, and other network infrastructure devices in your organization.

### Dialog Locations in Portal

The Data Sources dialog is available in the following Kentik portal locations:

* **Library » Dashboards**:

  + Add/Edit Dashboard dialog » Default Query Options » **Edit Data Sources** button
  + Add/Edit View Panel dialog » Query tab » Data Sources pane » **Edit Data Sources** button
  + Add/Edit View Panel dialog » Navigate To tab» Data Sources setting » Use Custom Data Source List » **Edit Data Sources** button
* **Data Explorer**: **Data Sources** pane in **Query** sidebar.
* **Insights & Alerting » Alert Policies**: Add/Edit Alert Policy dialog » Dataset tab » **Edit Devices** button (the dialog in this location is called Devices).
* **Settings » My Kentik Portal**: Add/Edit Tenant dialog » **Edit Devices** button (the dialog in this location is called Devices).

## Add Data Sources

The method used to specify data sources for a query depends on the dialog used for data source selection in your current location in the Kentik portal (see **Data Source Dialogs**).  
  
 To select data sources using the **Data Sources Dialog**:

1. In the **Data Sources** pane, click the **Edit Data Sources** button. The Data Sources dialog will open.
2. In the **Available Data Sources** list, do one of the following to add data sources to the **Selected Data Sources** list:

   * To add an individual data source, click on the data source in the Data Sources section.
   * To add all data sources with a given label, click on the label in the **Labels** section.
   * To add all data sources from a given site, click on the site in the **Sites** section.
3. Repeat the above to add more data sources either individually or by label or site.
4. Click the **Save** button. The data sources in the **Selected Data Sources** list will be listed in the **Data Sources** pane.

## Remove Data Sources

The method used to remove data sources from a query depends on the dialog used for data source selection in your current location in the Kentik portal (see **Data Source Dialogs**).  
  
 To remove data sources (individually or by label or site) in the **Data Sources Dialog**, do one of the following:

* In the **Selected Data Sources** list, click the gray **X** to the right of the name of a device, label, or site.
* In the **Available Data Sources** list, click a label or site (indicated with a highlight and checkmark).
* In the **Available Data Sources** list, if the **Filter** field is empty, click the **Clear** button at the right of the **Labels** or **Sites** heading to remove any items in the corresponding category.
