# Audit Log

Source: https://kb.kentik.com/docs/audit-log

---

The Kentik portal’s Audit Log is discussed here.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(708).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A49Z&se=2026-05-12T09%3A30%3A49Z&sr=c&sp=r&sig=RtEiWgg6hP5%2BPkHJeozu6mT%2BgLP2PZuj7wUvfr5N3mw%3D)

*The Audit Log lists changes made to your Kentik environment.*

## About Audit Log

The Audit Log is a filterable list of the changes made in your organization’s Kentik environment, so you can see when each change was made and who made it. Only admin-level users can view the data on the Audit Log page.

## Audit Log Page

The Audit Log page (accessed from the **Organization Settings** menu) is home to the Audit Log list, which lists the actions performed in your organization’s Kentik environment, including when and by which user.

> **Note:** Only admin-level users can access the Audit Log page.

### Audit Log Page UI

The Audit Log page includes the following main elements:

* **Actions**: A drop-down menu with the following option:

  + Export to CSV: Export the currently visible log entries to a.csv file.
* **Filters pane**: Controls that enable you to narrow the entries displayed in the list (see **Audit Log Filters**).
* **Audit Log**: A table that lists all actions that occurred within the specified time range.

### Audit Log List

The **Audit Log** list is a table that lists the actions taken by your organization's Kentik users. Information about the listed actions is provided in the following columns:

* **Source**: The API or the UI used to perform the specific configuration change (e.g., Kentik portal).
* **Action**: The specific action performed in the system (e.g. Create, Update, Delete, etc.).
* **Type**: The type of the entity affected by the action (e.g. User Preferences).
* **User**: The name of the user who performed the action.
* **Date**: The date-time of the action.

> **Notes:**
>
> * For additional details about a given action, click anywhere in the action's row in the list, which will open the **Audit Log Details** for that action.
> * To sort the list (ascending or descending) by a given column, click on the column's heading.

### Audit Log Filters

The **Filters** pane at the left of the Audit Log page enables you to narrow the Audit Log list to actions that match specified criteria. It includes the following controls:

![Audit log filters for time range, source, and action selection options displayed.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/AL-filters-sidebar.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A49Z&se=2026-05-12T09%3A30%3A49Z&sr=c&sp=r&sig=RtEiWgg6hP5%2BPkHJeozu6mT%2BgLP2PZuj7wUvfr5N3mw%3D)

* **Time Range**: Fields that show the start and end time of the time range during which the actions occurred. To change the time range, click a field to open the **Time Range Popup**. To close the control, click outside of the popup.
* **Filter criteria**: A set of criteria that an action must match to be displayed in the list (see **Filter Criteria**).
* **Reset To Default** (present at top right when one or more criteria are selected): A button that removes all currently set filters. The list will show all actions in the current time range.

#### Time Range Popup

The **Time Range** popup includes the following controls:

* **Lookback** **list**: A list of durations. Click in the list to set the duration (shown in the Time Range fields) for which the time range will look back from the present. Options include Today, This week, This month, Last month, Last 7 days, and Last 30 days.
* **Calendars**: Side-by-side monthly calendars that show the time range (if it spans more than one day) and enable you to change the range (shown in the Time Range fields) by clicking on a date. The start and end days of the time range are indicated in blue, and the intervening days are shaded. The calendars include the following controls:

  + Forward/Back: Buttons at the top left and right change the pair of displayed months.
  + Month/Year: Drop-down selectors enable you to choose the month and year.
  + Time fields: A field below each calendar enables you to enter specific start and stop times (24h clock).

    ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(100).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A49Z&se=2026-05-12T09%3A30%3A49Z&sr=c&sp=r&sig=RtEiWgg6hP5%2BPkHJeozu6mT%2BgLP2PZuj7wUvfr5N3mw%3D)

#### Filter Criteria

The following categories of criteria are available to apply in the **Filters** pane:

* **Source**: Checkboxes that select the portal location in which, or API with which, the action was performed.
* **Action**: Checkboxes that select the kind of action performed.
* **User**: A field that, when clicked, pops up a list from which you can choose a user by name:

  + The popup includes a filter field into which you can enter text to narrow the listed users, who can then be selected individually (by clicking) or as a group (with the **Select All** button).
  + Each selected user will appear as a lozenge in the **User** field, and the **Audit Log** list will be filtered to actions by those users.
  + To add more users, click the field again.
  + To remove a user, click the **X** at the right of the corresponding lozenge.
* **Type**: Checkboxes that select the type of the entity affected by the action.

The following rules apply to filter categories and criteria:

* Criteria are ORed (match any) within categories and ANDed (match all) between categories.
* When criteria are selected in more than one category a row is only displayed if it matches at least one selected criteria in all of the categories with criteria selected.
* When no criteria are selected for a category, the row may match any of the category's listed criteria.

#### Audit Log Filter Example

Given the following settings in the **Filters** pane, the **Audit Log** will list all Delete actions (regardless of source or type) performed during the specified time range by either of the two specified users:

* **Source**: No criteria selected
* **Action**: Delete
* **User**: Tea Ardis and River Song
* **Type**: No criteria selected

### Audit Log Details

Details about an individual action are shown in a sidebar that slides open from the right when you click on that action's row in the Audit Log list. The details sidebar includes the following elements:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(145).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A49Z&se=2026-05-12T09%3A30%3A49Z&sr=c&sp=r&sig=RtEiWgg6hP5%2BPkHJeozu6mT%2BgLP2PZuj7wUvfr5N3mw%3D)**Title**: A heading consisting of the action's action and type (see **Audit Log List**).
* **Close**: The X closes the sidebar.
* **User**: The user who performed the action.
* **Date/time**: The date and time at which the action was performed.
* **API**: The method and endpoint of the underlying API call that executed the action.
* **Payload**: The request body (JSON) of the underlying API call that executed the action.

---

© 2014-25 Kentik
