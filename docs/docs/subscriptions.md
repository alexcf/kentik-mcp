# Subscriptions

Source: https://kb.kentik.com/docs/subscriptions

---

This article discusses the **Subscriptions** page in the Kentik portal.

> **Note:** To create a subscription directly from a portal module, see **Portal View Sharing**.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(126).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A08Z&se=2026-05-12T09%3A36%3A08Z&sr=c&sp=r&sig=9dZi0hVUNv3pwQYtH9NJucPraI4Ez6O2%2FvUBxukkdHY%3D)

*Subscriptions, used to share Kentik data with others, are managed from the Subscriptions page.*

## About Subscriptions

Subscriptions enable you to share information from the Kentik portal by emailing it on a recurring schedule to a specified set of recipients. You can set up a subscription in a few different ways:

* **Subscription dialog**: On the Subscriptions page, open the Subscription dialog by clicking the **Add** **Subscription** button, then filling in the fields (see **Add a Subscription**).
* **Share dialog**: Click the **Share** button in the subnav of applicable portal pages (see **Portal Page Layout**). Next, fill in the **Subscription** tab (see **Sharing by Subscription**). The portal modules in which you can set up a subscription are listed in **Sharable Views**.
* **Subscribe dialog**: Click the **Actions** button in the subnav of applicable portal pages (see **Portal Page Layout**), then click **Subscribe**. Fill in the details as described in the above Share dialog section (see **Sharing by Subscription**).

## AI Advisor Subscriptions

### About AI Advisor Subscriptions

An AI Advisor subscription executes a user-defined **AI Advisor** prompt at a configured interval and sends the results to specified recipients. This is useful for teams that need recurring, automated AI-generated summaries, e.g., a NOC team that wants a nightly network health report delivered before the morning shift begins.

Each delivery includes:

* A PDF attachment containing the full conversation, including reasoning and tool call details.

### Setting Up an AI Advisor Subscription

To create an **AI Advisor** subscription in the Kentik portal: ![Subscription settings for report sharing, including email and schedule options.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/SUB-add-ai-advisor-sub.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A08Z&se=2026-05-12T09%3A36%3A08Z&sr=c&sp=r&sig=9dZi0hVUNv3pwQYtH9NJucPraI4Ez6O2%2FvUBxukkdHY%3D)

1. Open the Organization Settings from the main nav bar, and select **Report Subscriptions**.
2. Click **Add Subscription**.
3. From the Share dropdown menu, select **AI Advisor**.
4. In the **AI Adv****isor Input Question** field, enter the prompt or question you want AI Advisor to answer on each run. You can use detailed, **runbook-style** instructions.
5. Fill in any remaining info (file name, schedule, recipients, etc.)
6. Click **Share** to save and activate the subscription.

### Notes

* AI Advisor subscriptions run automatically without human review of each execution. Ensure your prompt is carefully crafted to avoid unintended outputs.
* The PDF attachment contains the full AI Advisor conversation, including intermediate reasoning steps.
* AI Advisor subscriptions follow the same access control rules as other subscriptions: Admin-level users can create, edit, and remove them; Member-level users can subscribe or unsubscribe.

## Subscriptions Page UI

The **Subscriptions** page is home to the **Subscriptions List**, which shows your organization’s active subscriptions. To access the page, choose **Report Subscriptions** from the **Organization Settings** menu on the portal's main navbar.

> **Note:** While member-level users can access the page, view subscriptions, subscribe, or unsubscribe, only Admin-level users can add, edit, or remove report subscriptions (see **About Users**).

The page includes the following UI elements:

* **Add Subscription**: A button (visible only to Admin-level users) that opens a **Subscription Dialog** where you can add a subscription.
* **Show/Hide Filters**: A button (filter icon) that toggles the **Filters** pane between expanded and collapsed.
* **Search**: A field that shows lozenges for the filters currently set in the **Filters** pane and enables you to enter text. The **Subscriptions** list will be filtered to show only rows in which there's a match between the entered text and the **Title** or **Recipients** columns. Click the **X** at the right of the field to clear entered text, and the **X** in a lozenge to clear the corresponding filter.
* **Filters**: A pane with controls that narrow the subscriptions shown in the **Subscriptions** list (see **Subscriptions Filters**).
* **Subscriptions list**: A table listing your organization’s current subscriptions (see **Subscriptions List**).

### Subscriptions Filters

The **Filters** pane at the left of the **Subscriptions List** narrows the list to subscriptions that match specified criteria. It includes the following controls:

* **Type**: Radio buttons that filter the list of subscriptions to show:

  + **All subscriptions**: All subscriptions within your organization will be displayed in the **Subscriptions** list.
  + **Only my subscribed reports**: Only subscriptions to which you subscribe will display.
  + **Subscriptions shared externally**: Only subscriptions shared outside of the organization will display.
* **Export Source**: Checkboxes that narrow the list to subscriptions that originate from the checked types of portal locations (see **Sharable Views**).

  > **Note**: “AI Advisor” is now an available source type.
* **Frequency**: Checkboxes that narrow the list to subscriptions that are sent to subscribers at the checked intervals.
* **Report File Type**: Checkboxes that narrow the list to subscriptions that are sent in the checked file format: PDF or CSV (see **Portal Export Options**).

### Subscriptions List

The **Subscriptions** list is a table that lists all active report subscriptions for your organization. To change the sort order of the list, click a heading to select a column, and click the resulting blue up or down arrow to choose the sort direction (ascending or descending).

#### Subscriptions List Columns

The **Subscriptions** list includes the following columns:

* **Title**: The name given to the subscription.
* **Export Source**: The kind of portal view (dashboard, saved view, etc.) from which the report was created (see **Sharable Views**).
* **Frequency**: The interval (Daily, Weekly, etc.) at which the report is sent to subscribers (see Schedule in **Subscription Tab UI**).
* **Recipients**: The subscribed email addresses. If there are too many addresses to fit, an ellipsis appears at the right of the field.
* **Last Report Sent On**: The date that the most recent report was sent to subscribers. Hover over the date for a tool tip with the full date-time (UTC).
* **Last Edited**: The date on which the subscription was last edited. Hover over the date for a tool tip with the full date-time (UTC).
* **Created**: The date on which the subscription was created. Hover over the date for a tool tip with the full date-time (UTC).
* **Actions**: A drop-down menu from which you can choose to edit, subscribe/unsubscribe, or remove the subscription (see **Subscriptions Actions**).

  > **Note**: Only Admin-level users can edit or remove a subscription.

#### Subscriptions Actions

The following actions are available from the **Action** menu that pops up from the vertical ellipsis at the right of each row:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(127).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A08Z&se=2026-05-12T09%3A36%3A08Z&sr=c&sp=r&sig=9dZi0hVUNv3pwQYtH9NJucPraI4Ez6O2%2FvUBxukkdHY%3D)**Edit** (Admin-level users only): Opens a **Subscription Dialog** that allows you to change the subscription information (see **Edit a Subscription**).
* **Subscribe**: Opens a dialog enabling you to confirm that you wish to be added to the subscription (using your Kentik-registered email address).
* **Unsubscribe**: Opens a dialog enabling you to confirm that you wish to be removed from the subscription. This option is inactive if you're not already subscribed.

  > **Note**: If you are the only recipient of a subscription and you unsubscribe, that subscription will not display in the Subscriptions list.
* **Remove** (Admin-level users only): Opens a confirmation dialog in which you can remove the subscription from your organization’s collection of subscriptions.

## Subscription Dialog

Adding or editing a subscription via the Subscriptions page involves specifying information in the fields of the Subscription dialog..

> **Note:** The Subscription dialog is visible only to Admin-level users.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(128).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A08Z&se=2026-05-12T09%3A36%3A08Z&sr=c&sp=r&sig=9dZi0hVUNv3pwQYtH9NJucPraI4Ez6O2%2FvUBxukkdHY%3D)

### Subscription Dialog UI

The Subscription dialog includes the following UI controls:

* **Cancel**: A control — either the **X** at the upper right or the **Cancel** button at the lower right — that exits the dialog while leaving all settings as they were when it was opened.
* **Share**: Saves the dialog’s current settings and exits the dialog.

### Subscription Dialog Fields

The fields of the Subscription dialog are the same as the fields on the **Subscription** tab of the Share dialog, covered in **Subscription Tab UI**, except that there is no **Manage Subscriptions** link at the top of the dialog.

## Manage Subscriptions

A subscription can be managed on the **Subscriptions Page** (accessed via the **Organization Settings** menu) or added/edited in the portal location whose information you'd like to share via a subscription (see **Sharable Views**).

### Add a Subscription

To add a new subscription:

1. Open the Subscriptions page from the **Organization Settings** menu on the portal's main navbar.
2. Click the **Add Subscription** button to open the Subscription dialog.
3. Specify the values of the fields in the dialog (see **Subscription Dialog Fields**).
4. Save the new subscription by clicking the **Share** button (lower right).

### Edit a Subscription

To edit the settings for an existing subscription:

1. Open the Subscriptions page from the **Organization Settings** menu on the portal's main navbar.
2. In the **Subscriptions List**, find the row corresponding to the subscription you’d like to edit, then click the vertical ellipsis at the right of the row to open the Action menu.
3. Choose **Edit**. The Subscription dialog will open.
4. Edit the subscription settings by changing any fields that you'd like to modify (see **Subscription Dialog Fields**).
5. To save changes, click the **Share** button (lower right).

### Remove a Subscription

To remove a subscription from your organization’s list of subscriptions:

1. Open the **Subscriptions** page from the **Organization Settings** menu on the portal's main navbar.
2. In the **Subscriptions List**, find the row corresponding to the subscription you’d like to edit, then click the vertical ellipsis at the right of the row to open the Action menu.
3. Choose **Remove**, then click **Remove** in the resulting confirmation dialog.

The subnav is the gray, horizontal bar or strip located below the main navigation bar or navbar in various pages and modules. It typically contains page-wide controls and buttons specific to the page's functionality.
