# Sites

Source: https://kb.kentik.com/docs/sites

---

This article discusses the Sites page in the Kentik portal.

> **Note:** If you need assistance with any aspect of managing your sites, contact Kentik (see **Customer Care**).

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(624).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A09Z&se=2026-05-12T09%3A38%3A09Z&sr=c&sp=r&sig=wpYstNuznYthdPMMdXyPQSN%2BsEl7sbxgoKIdXROb0wo%3D)

*The Sites page enables you to manage your organization's sites.*

## About Sites

A site is a specific user-defined physical location (address), e.g., a data center, with which you can associate one or more devices. Once a site is created in Kentik, you can assign one or more devices to the site via the settings for those devices. Sites with common characteristics, such as all PoPs in a particular metropolitan area, can be grouped into a "site market."  
  
The settings for sites and site markets are accessed via the Settings page (choose **Settings** from the main menu):

* Sites are managed on the **Sites** page: choose Sites under **Network Metadata** on the Settings page.
* Sites markets are managed on the Site Markets page: choose **Site Markets** under **Network Metadata**.

> **Notes:**
>
> * Kentik works best when all of your data sources are assigned to sites. Set the site property for each device via the **General** tab of the **Add Device** or **Edit Device** dialog (see **Device General Settings**).
> * To view the devices assigned to a given site, see **View Site Devices**.

## Sites Page

Sites are managed from the Sites page of the Kentik portal (Settings » **Sites**).

### Sites Page UI

The Sites page includes the following UI elements:

* **Share** (on subnav): Opens the **Share** dialog to enable you to share the current view (see **Sharing via the Share Dialog**).
* **Actions** (on subnav): A drop-down menu from which you can choose to export or subscribe to the current view:

  + Export: Export the page’s content as a visual report (PDF) or data table (CSV). A notification appears when the export is ready to download.
  + Subscribe: Opens the **Subscription Dialog**, which enables you to subscribe to regular reports from the page, either by choosing an existing subscription (combination of email address and schedule) or specifying a new one.
  + Unsubscribe: Opens the **Unsubscribe** dialog. Click the Subscription drop-down, select the subscription from which you’d like to unsubscribe, and click **Unsubscribe**.

    > **Note**: **Unsubscribe** will only appear if you have an existing subscription for this page.
* **Filters** (filter icon): A button that toggles the Filters pane between expanded and collapsed.
* **Group By**: A drop-down from which you can choose how to group the sites in the list:

  + If the selected value in this control is Country, Type, or Market then the table will include a heading row for each instance of that value (e.g., each country with a site), followed by the rows of the sites that share that value (e.g., the sites in a given country).
  + If the selected value is None then sites are listed, without grouping, in the sort order determined by the currently selected column heading.
* **Search field**: Filters the **Sites List** to show only rows that contain the entered text, either in one of the columns or in the **Site Details** card.
* **Add Site button**: Opens an **Add Site** dialog where you can add a new site (see **Add a Site**).
* **Sites List**: A table listing your organization’s sites (see **Sites List**).
* **Filters pane**: A set of checkboxes (in a sidebar at left) that enable you to filter the **Sites List** by Type, Site Country, or Site Market (see **Site Filters**).

> **Note:** Sites can also be added and edited with the **Site API**.

### Sites List

The **Sites** list is a table that shows information about the listed sites. By default, the list is ordered alphabetically by name. To change the sort order of the list, click a heading to choose a column, and click the resulting blue up or down arrow to choose the sort direction (ascending or descending).

#### Sites List Columns

The **Sites** list includes the following columns (left to right):

* **Select** (checkbox): Select one or more sites to perform one of the actions described in **Selected Site Actions**. Member-level users will not see this column.

  > **Note:** The checkbox in the heading row of the table toggles selection of the checkboxes in the individual site rows.
* **Site**: The site’s customer-assigned name string.
* **Type**: The type of the site (see Type in **Site Field Definitions**).
* **City**: The city in which the site is located.
* **Region**: The region in which the site is located.
* **Country**: The country in which the site is located.
* **Site Market**: The site market (if any) to which the site has been assigned.
* **PeeringDB Facility**: The PeeringDB Facility that is mapped to this Kentik site.

  > **Note**: Mapping Kentik sites to PeeringDB facilities allows users to visualize a common site/facility footprint with other networks regardless of whether or not they have their own PeeringDB record.

#### Sites List Actions

The far-right column of the **Sites** list includes icons that function as buttons for the following actions:

* **Edit site**: Opens an **Edit Site** dialog that allows you to review and edit the site’s settings (e.g., name, address, type, architecture, and site IP classification). See **Site Settings**.
* **Remove**: Opens a confirmation dialog that allows you to remove the site from your organization's collection of sites.

> **Note:** The above actions are not available to Member-level users.

#### Selected Site Actions

The controls listed below appear above the **Sites** list when the **Select** checkbox is checked for one or more rows. An action initiated with one of these controls will apply to all currently selected sites:

* **Type**: A drop-down from which you can choose a type to apply to the selected sites (see Type in **Site Field Definitions**).
* **Site Market**: A drop-down from which you can associate the selected sites with a new or existing site market (see **Site Market Selector**).
* **PeeringDB Facility**: A drop-down from which you can map the selected sites to one or more PeeringDB facilities (see **PeeringDB**).
* **Remove**: A button that opens a confirmation dialog enabling you to remove the selected sites from your organization's collection of sites.
* **Selection indicator**: A count of the currently selected sites in the Sites list.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(625).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A09Z&se=2026-05-12T09%3A38%3A09Z&sr=c&sp=r&sig=wpYstNuznYthdPMMdXyPQSN%2BsEl7sbxgoKIdXROb0wo%3D)

#### Site Market Selector

This drop-down selector offers the following options related to associating a site with a site market:

* **Manage Options**: Takes you to the Site Markets page.
* **Create New**: Click to create a new site market. The drop-down will be replaced with a name field, an inactive **Add** button, and a **Cancel** button. When you enter a name, the **Add** button will become active. Click it to save the new site market.
* **Existing site markets**: Click a site market in the drop-down to associate the site with that market.

### Site Details

The **Details** drawer for a given site opens from the right of the page when you click the site's row in the **Sites** list. The drawer includes the following information and controls:

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(626).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A09Z&se=2026-05-12T09%3A38%3A09Z&sr=c&sp=r&sig=wpYstNuznYthdPMMdXyPQSN%2BsEl7sbxgoKIdXROb0wo%3D)

* **Site name**: The user-defined name of the site.
* **Architecture**: The user-defined architecture of the site. For more information about defining a site's architecture, see **Architecture Edit UI**.
* **PeeringDB Facility**: The PeeringDB Facility to which this site is mapped.
* **Address**: The physical location of the site.
* **View in Network Explorer**: A button that opens the site’s Network Explorer details page (see **Core Details Pages**).

### Site Filters

The **Filters** pane to the left of the Sites list allows you to filter the list to show only sites that match the criteria that you specify with the pane's checkboxes. If the pane is collapsed you can expand it with the **Filters** toggle (see **Sites Page UI**).

* **Collapse** (left arrow at upper right): Hide the pane. The Sites list will expand to the left to fill in the resulting space.

  > **Note:** To expand the Filters pane, click the Filters icon at the left of the Group By button.

  ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(627).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A09Z&se=2026-05-12T09%3A38%3A09Z&sr=c&sp=r&sig=wpYstNuznYthdPMMdXyPQSN%2BsEl7sbxgoKIdXROb0wo%3D)
* **Site Country**: The country in which the site is located.
* **Site Market**: The site market (if any) to which the site has been assigned.
* **PeeringDB Facility**: The PeeringDB facility to which this Kentik site is mapped.

The following rules apply to filter categories and criteria:

* Criteria are ORed (match any) within categories and ANDed (match all) between categories.
* When criteria are selected in more than one category, a row is displayed only if it matches at least one of the selected criteria in all of the categories with criteria selected.
* When no criteria are selected for a category, the row may match any of the category's listed criteria.

### View Site Devices

To see the devices assigned to a given site:

1. On the Sites page, click the site’s row in the **Sites List**, which shows a **Site Details** card.
2. In the card, click the **View in Network Explorer** button, which takes you to the site's details page (see **Core Details Pages**).
3. In the **Traffic Tab Table** at the bottom of the page, click the **Devices** tab on the table, which lists this site's devices.

## Site Settings

The admin dialogs for sites are used to collect and display site information that is not only displayed in the UI but also determines how the traffic handled by the sites' devices is characterized for the analytics presented in various Kentik portal modules. The required information is entered into the fields of either of the following dialogs:

* **Add Site**: To register a new site with Kentik, open this dialog from the **Add Site** button on the Sites page (Settings » **Sites**).
* **Edit Site**: To edit an already registered site, open this dialog from the **Edit** icon at the right of that site's row in the Sites list.

> **Note:**
>
> * Sites can also be added and edited with the **Site API**.
> * Both of the above dialogs include a button that opens the **Edit Site Architecture** dialog (see **Architecture Edit UI**), which is used to provide a meaningful on-premises topology view for the site.

![Site editing interface showing AWS location and facility mapping options.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/ST-edit-sites-dialog(1).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A09Z&se=2026-05-12T09%3A38%3A09Z&sr=c&sp=r&sig=wpYstNuznYthdPMMdXyPQSN%2BsEl7sbxgoKIdXROb0wo%3D)

*A dialog for setting the properties of a site.*

### Site Dialogs UI

The **Add Site** and **Edit Site** dialogs share the same layout and the following common UI elements:

* **Close button**: Click the **X** in the upper right corner to close the dialog. All elements will be restored to their values at the time the dialog was opened.
* **Cancel button**: Cancel the add site or edit site operation and exit the dialog. All elements will be restored to their values at the time the dialog was opened.
* **Add Site button** (**Add Site** dialog only): Save settings for the new site and exit the dialog.
* **Save button** (**Edit Site** dialog only): Save changes to site settings and exit the dialog.

### Site Field Definitions

The site admin dialogs contain the fields and controls shown in the table below.

| Element | Type | Description |
| --- | --- | --- |
| Name | Editable field | Customer-assigned name string (up to 40 characters). |
| PeeringDB Facility Mapping | Drop-down with editable options | Map your site to a PeeringDB facility to allow our PeeringDB integration to find the common footprint between your network and any other network with a PeeringDB record. |
| Address | Editable field with drop-down options | The site’s street address. |
| Site Market | Drop-down with editable options | Choose or create a site market with which to associate this site; see **Site Market Selector**. |
| Type | Selectable options (choose one) | Data Center, Cloud, Branch/Building, Connectivity PoP, Customer/Partner, or Other. |
| Architecture | Button | Opens the Edit Site Architecture dialog, which is  explained in **Site Architecture**. |
| Site IP Classification | Editable field | Comma-separated lists of IP addresses and/or subnets for Infrastructure Networks, User Access Networks, and Other IPs. |

## Site Markets Page

The Site Markets page is covered in the following topics.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(628).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A09Z&se=2026-05-12T09%3A38%3A09Z&sr=c&sp=r&sig=wpYstNuznYthdPMMdXyPQSN%2BsEl7sbxgoKIdXROb0wo%3D)

### About Site Markets

A site market is a logical grouping to which you can optionally assign sites with any common characteristic of your choosing, such as all PoPs in a particular metropolitan area. Site markets are created on the Site Markets page, but sites are assigned to markets in the **Add Site** or **Edit Site** dialog of the Sites page (see **Site Settings**).  
  
You can get to the Site Markets page from the following locations:

* From the Settings page, choose **Site Markets** under **Network Metadata**.
* From the **Add Site** or **Edit Site** dialog, click the **Site Market Selector** and choose **Manage Options**.

### Site Markets Page UI

The main elements of this page include the following:

* **Add Site Market**: A button that opens the **Add Site Market** dialog (see **Site Market Settings**).
* **Search field**: Filters the **Site Markets** list to show only rows that contain the entered text.
* **Site Markets list**: A list of existing Site Markets.

### Site Market Settings

The settings for site markets are made in the following dialogs:

![Form to add a site market with fields for name and description.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/ST-add-site-market-dialog.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A09Z&se=2026-05-12T09%3A38%3A09Z&sr=c&sp=r&sig=wpYstNuznYthdPMMdXyPQSN%2BsEl7sbxgoKIdXROb0wo%3D)

* **Add Site Market**: Create a new site market and specify its properties.
* **Edit Site Market**: Change the properties of an existing site market.

These dialogs both contain the following fields:

* **Name** (required): A unique name for the site market.
* **Description**: User-provided details about the site market, for example an explanation of what its sites have in common.

### Site Markets List

The **Site Markets** list is a table that shows information about the site markets in your organization. By default, the list is ordered alphabetically by the Name column. To sort by a different column, click that column's heading, then click the associated up or down arrow to choose the sort direction (ascending or descending).  
  
The list includes the following columns:

* **Name**: The name of the site market.
* **# of Sites**: The total number of sites assigned to the site market.
* **Description**: A brief description of the site market, if provided.
* **Created**: The date that the site market was added.

#### Site Markets List Actions

The far-right column of the **Site Markets** list includes icons that function as buttons for the following actions:

* **Edit:** Opens the **Edit Site Market** dialog, where you can change the site market's properties (see **Site Market Settings**).
* **Remove**: Opens a confirming popup that enables you to remove the market from your organization's collection of site markets.

## Add or Edit a Site

The adding and editing of sites is covered in the following topics.

### Add a Site

To add a new site:

1. Choose Settings from the main menu, then **Sites** (under Network Metadata).
2. Open the **Add Site** dialog by clicking the **Add Site** button at upper right.
3. Enter a name in the **Name** field.
4. Specify the values of the remaining fields (see **Site Field Definitions**).
5. Click the **Add Site** button to add the site.

> **Note:** To add the site’s architecture, see **Site Architecture**.

### Edit a Site

To edit an existing site:

1. Choose Settings from the main menu, then **Sites** (under Network Metadata).
2. In the **Sites List**, click the edit site button in the row for the site you’d like to edit, which will open the **Edit Site** dialog.
3. Edit the fields that you want to change (see **Site Field Definitions**).
4. Click the **Save** button to save your changes.

> **Note:** To edit site architecture, see **Site Architecture**.

---

© 2014-25 Kentik
