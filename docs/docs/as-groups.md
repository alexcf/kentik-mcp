# AS Groups

Source: https://kb.kentik.com/docs/as-groups

---

AS groups enable users to designate a set of Autonomous Systems (ASes) whose traffic will be evaluated as a group (summed) for the purpose of top-X evaluations and filtering in queries.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(590).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A30Z&se=2026-05-12T09%3A39%3A30Z&sr=c&sp=r&sig=2g5zXMGSggHwcEDW3UXTr%2FRr%2B%2FyqiFkpBs9zse0GQf8%3D)

*The AS Groups page, where you can assign multiple ASes to a single group for evaluation and filtering in queries.*

## About AS Groups

Autonomous systems (ASes) represent the networks that collectively make up the Internet (see **IETF RFC 1930**), enabling the use of BGP to compute routes for the delivery of traffic. The AS Groups feature in Kentik lets you refer to a set of ASes collectively by assigning one or more ASes to a group that you name.

#### AS Groups in Queries

The use of AS groups has the following effect on Kentik queries:

* **Group-by**: If the **Use AS Groups** switch is on (see **Advanced Query Settings**), the results from all ASes in each AS group will be summed for top-X evaluation, graph plotting, and display in the results table. As shown below, if a table row represents a group, it will include a group icon at the right of the group name. Hover over the icon to display a list of the ASes in the group.

  > ***Notes:***
  >
  > + The **Use AS Groups** switch is only visible if at least one AS group exists in your organization (see **Add an AS Group**). When visible, the switch is enabled by default.
  > + When a Data Explorer view is saved to a dashboard panel or saved view, the setting of the switch at the time the view is saved will determine whether the panel or saved view will use AS Groups.

  ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(591).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A30Z&se=2026-05-12T09%3A39%3A30Z&sr=c&sp=r&sig=2g5zXMGSggHwcEDW3UXTr%2FRr%2B%2FyqiFkpBs9zse0GQf8%3D)
* **Filtering**: Inclusion or exclusion of traffic from a given AS group may be achieved by setting that group as the value of a filter on the AS Group dimension, as shown below.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(592).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A30Z&se=2026-05-12T09%3A39%3A30Z&sr=c&sp=r&sig=2g5zXMGSggHwcEDW3UXTr%2FRr%2B%2FyqiFkpBs9zse0GQf8%3D)

#### AS Groups Use Cases

There are several scenarios in which the ability to control how ASes are represented in query results may be useful to the operator of a network:

* **Logical grouping of ASes**: ISPs that are part of a broader consortium may have multiple ASes spanning multiple countries. AS grouping gives you the ability to assess the traffic from a multi-AS organization as a single entity, which is particularly important for interconnection/capacity planners in the context of peering.
* **Mapping of existing AS labels**: Many organizations refer to ASes with labels that are distinct from the official name of the AS. An AS group may be used to associate an existing label with one or more ASes, so that the AS naming by which ASes are represented in Kentik can be made consistent with the AS naming that you already use in other systems.
* **Naming by internal topology**: It’s now common in modern, large-scale data centers to deploy Clos architecture, which relies on BGP to the rack, with each rack corresponding to its own private ASN. AS groups enable you to reference and visualize groups of private ASes in a way that reflects this Clos-based datacenter topology.

## AS Groups Page

The AS Groups page (Settings » **AS Groups**), is used to build and manage AS groups. The page is covered in the following topics:

> **Note:** While members can view the AS Groups page, only admin users can add or modify AS Groups.

### AS Groups Page UI

The AS Groups page is made up of the following UI elements:

* **Filter field**: Filters the **Groups List** to show only groups whose name contains the entered text.
* **Add Group button**: Opens the **Add Group Dialog**, where you can define a new group.
* **Groups List**: A list of your organization's AS groups (see **Groups List**).
* **Group Pane**: Information and controls for the group that is currently selected in the Groups List (see **Group Pane**), including a list of the ASes that belong to that group.

### Groups List

The **Groups List** is a table on the left side of the screen that lists all AS groups in your organization. When you click on a row in the list, the **Group Pane** will display information about the corresponding AS group, including the individual ASes in the group.  
  
![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(593).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A30Z&se=2026-05-12T09%3A39%3A30Z&sr=c&sp=r&sig=2g5zXMGSggHwcEDW3UXTr%2FRr%2B%2FyqiFkpBs9zse0GQf8%3D)The table provides the following information for each AS group:

* **Group**: The name specified for the group.
* **AS Count**: The number of ASes assigned to the group. Hover over the number to display a popup listing the ASNs assigned to the group.

To sort the table (ascending or descending), click on a column heading.

### Group Pane

The **Group Pane** includes information and controls for the group that is currently selected in the **Groups List**, including a table that lists the individual ASes that your organization has assigned to that group.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(594).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A30Z&se=2026-05-12T09%3A39%3A30Z&sr=c&sp=r&sig=2g5zXMGSggHwcEDW3UXTr%2FRr%2B%2FyqiFkpBs9zse0GQf8%3D)

#### Group Pane UI

The Group Pane includes the following UI elements:

* **Group name**: A read-only field showing the name of group. To edit the name, click the pencil icon.
* **Edit name** (pencil icon): A button that makes the **Group name** field editable:

  + To rename the group, enter a different name in the field, then click the Save button.
  + To keep the existing name, click the **Cancel** button.
* **Add AS**: A button that opens the **Add AS** dialog to add one or more ASes to the current AS group.
* **Remove Group**: A button that opens a confirmation dialog that enables you to remove the current group from your organization’s collection of AS groups.
* **AS List**: A table containing information about the ASes in the current AS group (see **AS List**).

#### AS List

The **AS List** is a table providing the following information for each autonomous system (AS) in the current AS group:

* **ASN**: The number of the AS.
* **Name**: The name of the AS.
* **Type**: One of the following:

  + Public: A public AS whose number has been assigned by IANA.
  + Private: An AS used only internally within a single organization (enabling the use of BGP to compute routes between internal systems).
  + Reserved: An AS whose number is reserved by IANA (not available for assignment).
* **Remove** (trash icon): A button that removes the AS from the current AS group.

  > **Note:** The group will be removed immediately, without a confirmation dialog.

## Add Group Dialog

The **Add Group** dialog is used to add an AS group and assign ASes to that group.

![A dialog box for adding a group with selected options and a group name field.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/ASG-add-group-dialog.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A30Z&se=2026-05-12T09%3A39%3A30Z&sr=c&sp=r&sig=2g5zXMGSggHwcEDW3UXTr%2FRr%2B%2FyqiFkpBs9zse0GQf8%3D)

The dialog is made up of the following UI elements:

* **Close button**: Click the **X** in the upper right corner to close the dialog.
* **Group Name**: Enter a unique name for the new group.
* **AS lookup**: A field in which to enter part or all of an AS number or AS name. Matches for the entered text will appear in the **AS list** below.
* **AS list**: A list of all ASes whose name or number matches the characters entered in the **AS lookup** field. To select ASes from the list to add to the group, do one of the following:

  + Click the checkbox to the left of an individual AS.
  + Click **Select All**.
* **Select All** (visible only when **AS list** is populated and contains more than one AS):

  + If no checkboxes are checked in the **AS list**, click this to check all of the checkboxes.
  + If any checkboxes are checked in the **AS list**, click this to uncheck all of the checkboxes.
* **Cancel**: A button that cancels the operation and exits the **Add Group** dialog.
* **Add Group**: A button that adds the new group to your organization’s collection of AS groups and closes the dialog.

## Add AS Dialog

The **Add AS** dialog, which opens from the Add AS button (see **Group Pane UI**), is used to add an AS to the current AS group (the group selected in the **Groups List** and displayed in the **Group Pane**).

![Interface for adding AS with options and a warning to select at least one.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/ASG-add-as-dialog.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A30Z&se=2026-05-12T09%3A39%3A30Z&sr=c&sp=r&sig=2g5zXMGSggHwcEDW3UXTr%2FRr%2B%2FyqiFkpBs9zse0GQf8%3D)

The dialog is made up of the following UI elements:

* **Close button**: Click the **X** in the upper right corner to close the dialog.
* **AS lookup**: A field in which to enter part or all of an AS number or AS name. Matches for the entered text will appear in the **AS list** below.
* **AS list**: A list of all ASes whose name or number matches the characters entered in the **AS lookup** field. To select ASes from the list to add to the group, do one of the following:

  + Click the checkbox to the left of an individual AS.
  + Click **Select All**.
* **Select All** (visible only when **AS list** is populated and contains more than one AS):

  + If no checkboxes are checked in the **AS list**, click this to check all of the checkboxes.
  + If any checkboxes are checked in the **AS list**, click this to uncheck all of the checkboxes.
* **Cancel**: A button that cancels the operation and exits the Add Group dialog.
* **Save button**: A button that adds the selected ASes to the current AS group and closes the dialog.

## Manage AS Groups

Configuration of AS groups on the AS Groups page (Settings » **AS Groups**) is covered in the following topics.

### Add an AS Group

To add an AS group:

1. On the AS Groups page (Settings » **AS Groups**), click the **Add Group** button to open the **Add Group Dialog**.
2. In the **Group Name** field, enter a name for the new group.
3. Click in the **ASes field** and start entering an AS number or name to show a list of all matching ASes.
4. Click to add a checkmark next to the AS you want to add to the group or click **Select All** to select all ASes displayed in the AS List.

   > **Note:** If you need to use more than one set of search criteria to find all of the ASes you want to add to the group, you’ll need to create the group first and then edit it once it’s created (see **Edit an AS Group**).
5. Click the **Add Group** button. The dialog will close and the new group will appear in the **Groups List**.

### Edit an AS Group

AS groups are edited directly on the AS Groups page; there is no **Edit Group** dialog. To choose a group to edit, click in the **Groups List** on the row of the group that you want to edit. The group will appear in the **Group Pane**.  
  
Once a group is listed the **Group Pane** you can perform the following actions:

#### Change the name of the AS group

1. Click the **Edit** **name** button (pencil icon) to make the **Group name** field editable.
2. Enter a different name in the field.
3. Click the **Save** button.

#### Add a new AS to the group

1. Click the **Add AS** button to open the **Add AS Dialog**.
2. In the dialog’s ASes field, start entering an AS number or name to show a list of all matching ASes.
3. In the list, click to add a checkmark for the AS that you want to add to the group or click **Select All** to select all ASes displayed in the AS List.
4. When you’ve selected all of the ASes that you wish to add to the group, click the **Save** button to close the dialog and add the selected ASes to the current group.

#### Remove an AS from the group

1. Click the **Remove** button (trash icon) at the right side of the row corresponding to the AS that you want to remove.
2. The AS will be removed from the current group (without a confirmation dialog).

### Remove an AS Group

To remove a group from your organization’s collection of AS groups:

1. In the **Groups List**, click in the row of the group that you want to remove. The group will appear in the **Group Pane**.
2. Click **Remove Group** at the top right of the Group Pane, which will open a confirmation dialog that enables you to remove the group.
3. In the dialog, click the **Remove** button. The dialog will close, the group will be removed, and a notification will confirm the removal.

---

© 2014-25 Kentik
