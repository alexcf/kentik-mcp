# Filters Settings

Source: https://kb.kentik.com/docs/filtering-settings

---

This article covers filters settings in queries for **Data Explorer** and **Metrics Explorer** dashboard panels in the Kentik portal.

> **TIP:** For best performance, whenever possible use the Device selector (see **Data Sources Pane**) to choose devices rather than filtering for devices.

## About Filters

In Kentik’s **Data Explorer****,** **Metrics Explorer****,** and dashboard panels, a filter allows you to narrow query results inclusively/exclusively based on a value (or range of values) for one or more dimensions (see **About Dimensions**). Individual filter conditions may be grouped, ANDed, ORed, and nested to create complex compound filters.

#### Filter Types

Kentik supports the following types of filters:

* **Ad hoc filters**: Settings are defined and applied using controls in the **Filters Dialog**. For a list of locations in which the dialog is used (e.g., the **Filters** pane of the Data Explorer sidebar) see **Filters Dialog Locations**.
* **Saved filters**: Settings are called up from previously stored values (see **Saved Filters**). Includes both of the following:

  + **Company filters**: Filters that are created and shared by users within your organization (see **Company Saved Filters**).
  + **Preset filters**: Built-in filters provided by Kentik (see **Preset Saved Filters**).

## Filters Pane

The **Filters** pane appears in the sidebar in **Data Explorer**, **Metrics Explorer**, Dashboards, and in the settings dialog for the individual panels on dashboards. The pane includes the following UI elements:

* **Filter Groups**: A list of the filter groups that are currently applied to the query:

  + When no filters are selected, the list is replaced with "No filters applied."
  + An individual filter group in the list can be removed by clicking the **X** at right.
* **Edit Filters**: A button that opens the **Filters Dialog**.
* **Add Dimension/Metric Filters** (Metrics Explorer): A button that opens the **Filters Dialog**.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/DE Filters Pane.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A12Z&se=2026-05-12T09%3A38%3A12Z&sr=c&sp=r&sig=VpQUULvNlOrCJ%2B3opGjuCmytifVy64KvzCq4%2Fgcogv8%3D)

*Data Explorer Filters Pane*

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/ME Filters Pane.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A12Z&se=2026-05-12T09%3A38%3A12Z&sr=c&sp=r&sig=VpQUULvNlOrCJ%2B3opGjuCmytifVy64KvzCq4%2Fgcogv8%3D)

*Metrics Explorer Filters Pane*

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/Library Dashboard Filters Pane.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A12Z&se=2026-05-12T09%3A38%3A12Z&sr=c&sp=r&sig=VpQUULvNlOrCJ%2B3opGjuCmytifVy64KvzCq4%2Fgcogv8%3D)

*Library Dashboard Filters Pane*

## Filters Dialog

Except when managing saved filters (see **Saved Filter Admin Dialogs**), filters are configured in the Filters dialog.

> **Note:**The filter groups controls in the Filters dialog are covered in **Filter Groups UI**.

### Filters Dialog UI

The Filters dialog UI includes the following elements:

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/Filters Dialog.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A12Z&se=2026-05-12T09%3A38%3A12Z&sr=c&sp=r&sig=VpQUULvNlOrCJ%2B3opGjuCmytifVy64KvzCq4%2Fgcogv8%3D)

*Filters Dialog via Data Explorer Filters Pane*

* **Close button**: Click the **X** in the upper right corner to close the dialog without saving changes to the filter settings.
* **Filter Groups pane**: The primary control area for configuring, saving, and applying filters:

  + The main controls of the pane are covered in **Filter Groups UI**.
  + The initial state of the pane is covered in **Filter Groups Initial State**.
* **Cancel button**: Close the dialog without saving changes.
* **Save button**: Save changes to filter settings and exit the dialog.

### Filters Dialog Locations

The Filters dialog is used in multiple portal locations to configure filters.

> **Note**: The table also indicates the absence of certain controls in certain portal sections.

| Portal Location | Access via… | Save Filters button | Add Saved Filter button | Add Nested Group button |
| --- | --- | --- | --- | --- |
| Library Dashboards | Filters pane | Yes | Yes | Yes |
| Library Dashboard Panels | Actions » Edit Dashboard | Yes | Yes | Yes |
| Data Explorer | Filters pane (Query sidebar) | Yes | Yes | Yes |
| Metrics Explorer | Filters pane (Query sidebar) | No | No | No |
| Alerting » **Manage** **Alert Policies** | Add/Edit Policy | Yes | Yes | Yes |
| Organization Settings » **Users** | Add/Edit User dialog | No | Yes | No |

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/Alert Policies Filters.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A12Z&se=2026-05-12T09%3A38%3A12Z&sr=c&sp=r&sig=VpQUULvNlOrCJ%2B3opGjuCmytifVy64KvzCq4%2Fgcogv8%3D)

*Filters via Alerting » Manage Alert Policies*

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/Users Filters(1).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A12Z&se=2026-05-12T09%3A38%3A12Z&sr=c&sp=r&sig=VpQUULvNlOrCJ%2B3opGjuCmytifVy64KvzCq4%2Fgcogv8%3D)

*Filters via Organization Settings* *»* *Users*

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/Library Edit Dashboard Filters.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A12Z&se=2026-05-12T09%3A38%3A12Z&sr=c&sp=r&sig=VpQUULvNlOrCJ%2B3opGjuCmytifVy64KvzCq4%2Fgcogv8%3D)

*Filters via Library* *»* *Edit Dashboard*

## Filter Groups UI

Filters are configured within **Filter Groups**, enabling flexible control over the ANDing and ORing of their component conditions. Filter groups can:

* Be configured either in the **Filters Dialog** or the **Saved Filter Admin Dialogs**.
* Contain up to 250 ad hoc or saved filters (company or preset, see **Filter Types**).
* Be nested within other filter groups.

### Filter Groups Initial State

The initial state of the **Filter Groups** pane occurs only if the Filtering Groups dialog is opened when there are no filters currently defined for the query. It does not occur in the **Saved Filter Admin Dialogs**.

In its initial state the **Filter Groups** pane contains the following two buttons:

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/Add Dimensions or Metrics Filter(1).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A12Z&se=2026-05-12T09%3A38%3A12Z&sr=c&sp=r&sig=VpQUULvNlOrCJ%2B3opGjuCmytifVy64KvzCq4%2Fgcogv8%3D)

*Filter Groups Initial State via the Add Dimension/Metric Filters button in the Metrics Explorer Filters Pane*

* **Add Ad Hoc Filter**: Adds a new filter group with the default dimension (e.g., Destination ASN) and unspecified value (see **Ad Hoc Filter Controls**).
* **Add Saved Filter**: Opens a drop-down menu to choose a filter from a list of all available saved filters (company and preset).

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/Edit Filters(1).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A12Z&se=2026-05-12T09%3A38%3A12Z&sr=c&sp=r&sig=VpQUULvNlOrCJ%2B3opGjuCmytifVy64KvzCq4%2Fgcogv8%3D)

*Filter Groups Initial State via the Edit Filters button in the Data Explorer Filters Pane*

### Filter Groups Controls

The following UI elements are high-level controls for the entire set of filter groups in the **Filter Groups** pane:

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/Filters Dialog.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A12Z&se=2026-05-12T09%3A38%3A12Z&sr=c&sp=r&sig=VpQUULvNlOrCJ%2B3opGjuCmytifVy64KvzCq4%2Fgcogv8%3D)

*Filters Dialog via Data Explorer Filters Pane*

* **Match selector** (shown only when there are multiple filter groups): The drop-down match selector determines the conjunctive operator used to join all groups in the filter set:

  + All of the following groups (default) is used to AND the groups.
  + Any of the following groups is used to OR the groups.
* **Save Filters**: A button at upper right that opens the **Save Filter Dialog** so that you can save your current filter configuration for later reuse.

  > **Note:** Not shown in all locations (see **Filters Dialog Locations**).
* **Remove All Groups**: Removes all filter groups.
* **Filter Groups**: One or more filter groups with ad hoc filters and/or saved filters; see **Group-level Controls**.
* **Add Filter Group buttons**: Adds a new filter group to the set:

  + **Add Ad Hoc Filter**: Adds a new filter group with the default dimension (e.g., Destination ASN) and unspecified value.
  + **Add Saved Filter**: Opens a drop-down menu to choose a filter from a list of all available saved filters (company and preset). See **Saved Filter Controls**.
  > **Note:** Not shown in all locations (see **Filters Dialog Locations**).

### Group-level Controls

A filter group is a container for one or more individual filters (conditions) or nested groups.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/Group-level Controls.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A12Z&se=2026-05-12T09%3A38%3A12Z&sr=c&sp=r&sig=VpQUULvNlOrCJ%2B3opGjuCmytifVy64KvzCq4%2Fgcogv8%3D)

*Group-level Controls in the Filters Dialog via the Data Explorer Filters Pane*

Each filter group may include the following controls:

* **Include/Exclude**: A drop-down selector that determines whether results that match the group are included or excluded from the results of queries to which the filter set is applied.
* **All/Any**: A drop-down selector that, if there are multiple filters in a filter group, determines the conjunctive operator used to join those filters:

  + **All** (default): Select to AND the filters.
  + **Any**: Select to OR the filters.
* **Remove Group**: Removes the entire filter group.
* **Remove All Groups**: Removes all filter groups for the query.
* **Filters**: One or more of either of the following (up to a maximum of 250 total):

  + **Ad Hoc Filter**: See **Ad Hoc Filter Controls**.
  + **Saved Filter**: See **Saved Filter Controls**.
* **Add Ad Hoc Filter** (button): Adds a new filter group with the default dimension (e.g., Destination ASN) and unspecified value. The UI of an individual ad hoc filter is covered in **Ad Hoc Filter Controls**.
* **Add Saved Filter** (button): Opens a drop-down menu to choose a filter from a list of all available saved filters (company and preset). See **Saved Filter Controls**.
* **Add Nested Group** (button): Add a new ad hoc filter group (rather than an individual filter condition) that is nested inside the current filter group (instead of at the same level as the current group). Nested filter groups enable you to create more complex filters.

  > **Note:** **Add Saved Filter** and **Add Nested Group** buttons are not shown in all locations (see **Filters Dialog Locations**).

### Ad Hoc Filter Controls

Each ad hoc filter in a filter group defines a condition and includes the following controls:

* **Dimension selector**: Opens a dialog (see **Dimension Selector Dialog**) enabling you to choose the dimensions (see **Dimensions Reference**) on which to filter.
* **Direction button**: Appearing just to the right of any directional dimension, this button toggles the perspective of the dimension between source and destination.![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/Ad Hoc Filter Controls.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A12Z&se=2026-05-12T09%3A38%3A12Z&sr=c&sp=r&sig=VpQUULvNlOrCJ%2B3opGjuCmytifVy64KvzCq4%2Fgcogv8%3D)
* **Operator**: The operator to apply in the filter. Options, which vary depending on whether the dimension takes numeric or alphanumeric values, may include:

  + Equals, does not equal
  + Greater than, less than
  + Contains, does not contain (case insensitive)
  + Matches regex/does not match regex
  + Equals the value of, does not equal the value of (applies only to certain dimensions, see **Source/Destination Comparisons**).
* **Value**: Opens a dropdown list to select the value to match.

  > **Note**: Leave the value undefined to include or exclude all records for which no value has been assigned to the filter dimension.
* **Remove** (**X**): Removes the filter from the group.

#### Source/Destination Comparisons

The following dimensions have the “equals the value of” and “does not equal the value of” operators available for comparing source and destination traffic:

| Dimension Family | Dimension Name  (Source or Destination) |
| --- | --- |
| IP & BGP Routing | Site by IP |
| IP & BGP Routing | Site Type by IP |
| Amazon Web Services | Zone |
| Amazon Web Services | Region |
| Google Cloud Platform | Zone |
| Google Cloud Platform | Region |
| Microsoft Azure | Zone |
| Microsoft Azure | Region |

> **TIP**: These two operators can be useful for investigating:
>
> * Traffic across various public cloud regions and availability zones in GCP, AWS, and Azure.
> * Inter-site traffic based on source or destination IP addresses as defined in the **Site Field Definitions**.

#### Dimension Selector Dialog

The dimension selector dialog allows you to choose a dimension on which to filter, and includes the following elements:

* **Filter field**: Filters the available dimensions to show only those whose name contains the entered text.
* **Dimension lists**: Lists of the available dimensions in each category, which are organized as described in **Dimension Selection Groups**.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/Dimension Selector Dialog.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A12Z&se=2026-05-12T09%3A38%3A12Z&sr=c&sp=r&sig=VpQUULvNlOrCJ%2B3opGjuCmytifVy64KvzCq4%2Fgcogv8%3D)

### Saved Filter Controls

Saved filters in filter groups have the following controls:

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/Saved Filter Controls(1).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A12Z&se=2026-05-12T09%3A38%3A12Z&sr=c&sp=r&sig=VpQUULvNlOrCJ%2B3opGjuCmytifVy64KvzCq4%2Fgcogv8%3D)

*Saved Filter Controls via the Data Explorer Filters Pane*

* **Expand/Collapse** (triangle icon): Shows or hides a saved filter detail, which is a read-only list of the filter groups and individual filters that make up the saved filter.
* **Name**: The name assigned to the filter when it was saved or edited.
* **Description**: The description text entered when the filter was saved or edited.
* **Include/Exclude Results**: A drop-down selector that determines whether results that match the saved filter are included or excluded from the results of the queries in which the filter is applied.
* **Customize**: Converts the saved filter into an ad hoc filter so that its filter groups and individual filters can be modified using the controls described in **Ad Hoc Filter Controls**.
* **Filter details**: A list of the individual filters (conditions) in the saved filter, as well as their associated logic for inclusion (include/exclude) and matching (any/all).

#### How to Add a Saved Filter

1. Click the **Add Saved Filter** button.
2. Type the name of a specific filter into the field at the top of the list.
3. Select the filter.

   * A new filter group that contains the newly chosen filter will be added.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image-1769725367380.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A12Z&se=2026-05-12T09%3A38%3A12Z&sr=c&sp=r&sig=VpQUULvNlOrCJ%2B3opGjuCmytifVy64KvzCq4%2Fgcogv8%3D)

## Save Filter Dialog

A collection of filter groups can be saved for later reuse in the Filtering Options dialog via the **Add Saved Filter** button in the **Filter Groups Controls** or the **Group-level Controls**. The saved filter will be available to all other users in the organization and will be labeled as a company filter (see **Company Saved Filters**) in the Saved Filters List (Settings » **Saved Filters**).

### How to Save a Filter

1. Click the **Save Filters** button in the **Filter Groups Controls**.
2. Add a **Filter Name** (required), which is a unique user-supplied name for the saved filter.
3. Add a **Filter Description** (optional), which is a user-supplied text that will display when the cursor hovers over the filter in the Saved Filters selection box.
4. Click the **Add Filter** button to save the filter.

> **Note:** To see a list of existing company filters, go to Settings » **Saved Filters** and click **Company** in the Share Levelselector (see **Saved Filters Page UI**).

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/Save Filter Dialog.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A12Z&se=2026-05-12T09%3A38%3A12Z&sr=c&sp=r&sig=VpQUULvNlOrCJ%2B3opGjuCmytifVy64KvzCq4%2Fgcogv8%3D)
