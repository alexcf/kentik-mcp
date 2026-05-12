# Saved Filters

Source: https://kb.kentik.com/docs/saved-filters

---

Saved filters include both customer-created “company” filters and Kentik-provided “preset” filters.

> **Notes**:
>
> * Saved filters can be applied in multiple locations in the portal, typically via the **Filters Dialog**. For a list of locations, see **Filters Dialog Locations**.
> * For general information about the structure and application of filters, see **About Filters**.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(635).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A12Z&se=2026-05-12T09%3A37%3A12Z&sr=c&sp=r&sig=7dvfrvFlJ%2BQ7f98sY1QsTN%2FPMQ4Gq6kWZ3xnIA5%2B%2Fxw%3D)

*The Saved Filters page lists saved filters that you can apply via the Filtering Options dialog.*

## About Saved Filters

Two types of saved filters are supported in Kentik and managed from the Saved Filters page (Settings » **Saved Filters**):

* **Company Saved Filters**
* **Preset Saved Filters**

Saved filters enable you to apply a stored set of filter values to a query so that you can narrow results, either inclusively or exclusively, based on a value or range of values for one of the dozens of dimensions that represent the fields of the KDE main table (see **About Dimensions**). In the underlying SQL for a given query, the specified filters are included in the WHERE clause.

### Company Saved Filters

A company filter is a filter created by a Kentik user and available to other users in the same organization. Company filters are:![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(636).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A12Z&se=2026-05-12T09%3A37%3A12Z&sr=c&sp=r&sig=7dvfrvFlJ%2BQ7f98sY1QsTN%2FPMQ4Gq6kWZ3xnIA5%2B%2Fxw%3D)

* Created in and saved from one of the following:

  + The **Filters Settings** on the **Filters** pane of the **Query** sidebar in Data Explorer or in dashboards (Library » **Dashboards);**
  + The Saved Filters page (Settings » **Saved Filters**), either by duplicating an existing filter (preset or company) or by creating a filter from scratch (see **Add or Edit Saved Filter**).
* Applied to queries via the **Add Saved Filter** button in the **Filter Groups UI**, which may be accessed in various locations including the **Filters** pane of the **Query** sidebar in Data Explorer and in dashboards in the portal Library (see **Filters Dialog Locations**).
* Managed via the **Saved Filters Page**.

### Preset Saved Filters

A preset filter is a ready-made filter provided as part of Kentik. Preset filters are:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(637).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A12Z&se=2026-05-12T09%3A37%3A12Z&sr=c&sp=r&sig=7dvfrvFlJ%2BQ7f98sY1QsTN%2FPMQ4Gq6kWZ3xnIA5%2B%2Fxw%3D)Applied to queries via a drop-down menu from the **Add Saved Filter** button in the **Filters Dialog**. For a list of locations in which the dialog is used (e.g., the **Filters** pane of the Data Explorer sidebar) see **Filters Dialog Locations**.
* Managed via the **Saved Filters Page**.

## Saved Filters Page

To access the Saved Filters page, from the main menu, choose **Settings** and then select **Saved Filters** from the **Customizations** section.

### Saved Filters Page UI

The Saved Filters page includes the following UI elements:

* **Share Level selector**: Allows you to narrow the filters listed in the **Saved Filter List** by share level:

  + All: List all saved filters.
  + Company: List only company saved filters (see **Company Saved Filters**).
  + Preset: List only preset saved filters (see **Preset Saved Filters**).
* **Filter field**: Filters the **Saved Filter List** to show only filters whose name or description contains the entered text.
* **Add Saved Filter button**: Opens an Add Saved Filter dialog where you can add a saved filter (see **Saved Filter Admin Dialogs**).
* **Saved Filter List**: A table listing your organization’s preset and company filters (see **Saved Filter List**).

### Saved Filter List

The **Saved Filter List** is a table that shows information about the listed saved filters as well as available actions. By default, the list is ordered alphabetically by name. To change the sort order of the list, click on a heading to choose a column to sort, and on the resulting blue up or down arrow to choose the sort direction (ascending or descending).

> **Note:** To see additional information about a saved filter, click anywhere in the filter’s row to open an Edit Saved Filter dialog where you can review and sometimes edit settings (see **Saved Filter Admin Dialogs**).

The Saved Filter List includes the following columns (left to right):

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(638).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A12Z&se=2026-05-12T09%3A37%3A12Z&sr=c&sp=r&sig=7dvfrvFlJ%2BQ7f98sY1QsTN%2FPMQ4Gq6kWZ3xnIA5%2B%2Fxw%3D)**Share Level**: Either preset or company.
* **Name**: The name of the saved filter, as well as a description (i.e. what it does and how it’s used).
* **Copy**: Duplicates the saved filter, appending “Copy” to the name, and adds it to the list.
* **Remove** (company filters only): Opens a confirming dialog that enables you to remove the saved filter from your organization.

## Saved Filter Admin Dialogs

Adding or editing a saved filter via the Kentik portal involves specifying information in the fields of the saved filter admin dialogs, which are covered in the following topics.

> **Notes:**
>
> * Saved filter admin dialogs are visible only to users whose level is Administrator.
> * Saved filters can also be added and edited with the **Saved Filter API**.

### About Saved Filter Dialogs

The Kentik portal uses saved filter admin dialogs to collect and display saved filter information. The required information is entered into or displayed in the fields of the following dialogs:

* Add Saved Filter when creating a new saved filter.
* Edit Saved Filter when editing an existing company filter.
* Preset Filter Details for viewing (but not editing) details of a preset filter (see **Preset Filter Details Dialog**).

### Saved Filter Dialogs UI

The Add Saved Filter and Edit Saved Filter dialogs share the same layout and have the following UI elements:

* ![Configuration screen for AWS filter settings, focusing on source and destination zones.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/SF-edit-saved-filter-dialog.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A12Z&se=2026-05-12T09%3A37%3A12Z&sr=c&sp=r&sig=7dvfrvFlJ%2BQ7f98sY1QsTN%2FPMQ4Gq6kWZ3xnIA5%2B%2Fxw%3D)**Close**: An **X** in the upper right corner that you can click to close the dialog. All elements will be restored to their values at the time the dialog was opened.
* **Filter Name**: A field where you can enter a name for the saved filter.
* **Description**: A field where you can enter a description string, which should include what the filter does and how it is used.
* **Filter Groups pane**: The interface for managing the filter groups that make up the saved filter, as well as for configuring an individual filter group.

  > **Notes:**
  >
  > + The UI of the **Filter Groups** pane is mostly the same in the saved filter admin dialogs as it is in the **Filtering Options** dialog (see **Filter Groups UI**). For a description of differences, see **Filter UI Variations**.
  > + When you first open the Add Saved Filter dialog, the **Filter Groups** pane is populated with one filter group whose default dimension is Destination AS Number.
* **Cancel**: A button that cancels the add/edit operation and exits the dialog. All elements will be restored to their values at the time the dialog was opened.
* **Add Saved Filter** (Add Saved Filter dialog only): A button that saves changes to saved filter settings and exits the dialog.
* **Save** (Edit Saved Filter dialog only): A button that saves changes to saved filter settings and exits the dialog.

#### Filter UI Variations

The **Filter Groups** pane is used in multiple locations in the Kentik portal. In most cases, the pane appears in the Filtering Options dialog (see **Filter Groups UI**). The interface of the pane as it appears in the Add Saved Filter and Edit Saved Filter dialogs is mostly the same as elsewhere, but in the saved filter dialogs the **Save Filters** button is not needed (the **Add Saved Filter** or **Save** button at the bottom right of the dialog is used instead).

### Preset Filter Details Dialog

The Preset Filter Details dialog opens when you click on a row in the **Saved Filter List** that represents a preset filter (see **Preset Saved Filters**). The information in this dialog is read-only (not editable). The dialog contains the following UI elements:

* ![List of commonly scanned TCP ports with specific destination port numbers and protocols.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/SF-preset-filter-details-dialog.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A12Z&se=2026-05-12T09%3A37%3A12Z&sr=c&sp=r&sig=7dvfrvFlJ%2BQ7f98sY1QsTN%2FPMQ4Gq6kWZ3xnIA5%2B%2Fxw%3D)**Close button**: Click the **X** in the upper right corner to close the dialog.
* **Filter Name**: The user-specified name of the filter.
* **Description**: A user-specified description string including (ideally) what the filter does and how it is used.
* **Filter Groups**: One or more filter group tiles, each of which contains:

  + Any/All indicator: Indicates whether the filters in the filter group are matched on an Any or All basis.
  + Filters: A list of the individual filters in the filter group, including the dimension, operator, and value.
* **Match indicator**: Indicates whether the filter groups in the saved filter are matched on an AND or OR basis.

## Add or Edit Saved Filter

Saved filters are added and edited via the Saved Filters page of the Kentik portal (Settings » **Saved Filters**).

> **Note:** The Saved Filters page is visible only to users whose level is Administrator.

### Add a Saved Filter

To add a new saved filter:

1. Open the Saved Filters page (Settings » **Saved Filters**).
2. Click the **Add Filter** button to open the Add Saved Filter dialog.
3. Specify the values of the general fields in the dialog (see **Saved Filter Dialogs UI**).
4. Under **Filter Groups**, configure one or more filters (see **Filter Groups UI**).
5. Save the new saved filter by clicking the **Add Saved Filter** button (lower right).

### Edit a Saved Filter

Only saved filters with a share level of “Company” can be edited. To edit the settings for an existing company filter:

1. Open the Saved Filters page (Settings » **Saved Filters**).
2. In the **Saved Filter List**, click in the row of the company filter that you'd like to edit. The Edit Saved Filter dialog will open.
3. Edit the saved filter settings:

   1. Change any general fields that you'd like to modify (see **Saved Filter Dialogs UI**).
   2. Add, remove, or modify any filter groups or individual filters that you’d like to change (see **Filter Groups Interface**).
4. To save changes, click the **Save** button (lower right).
