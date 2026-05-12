# Public Shares

Source: https://kb.kentik.com/docs/public-shares

---

This article discusses the **Public Shares** feature of the Kentik portal.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(129).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A51Z&se=2026-05-12T09%3A32%3A51Z&sr=c&sp=r&sig=7m2nhLPaVz%2BIVgakST8Gqn9Gb1ZbQi51qpzcWaTBDxE%3D)

*A public share, like this one from Data Explorer, lets you show Kentik-generated views to non-Kentik users.*

## About Public Shares

A public share is a page within Kentik that contains data and visualizations that you can share securely with people who aren't registered Kentik users or aren't in your Kentik organization. Public sharing is available for a several different types of views in the Kentik portal (see **Sharable Views**). Each public share page has a link that you can provide to a visitor who can then navigate to the page to see the information you'd like to share.  
  
The destination page for a public share is not the original page from which the public share was created; instead, it is based on the original page while being designed for sharing outside the context of the Kentik portal. The visitor will be able to see the results from the original page — e.g., the chart and table showing query results in Data Explorer — and make some changes to how that information is displayed on the share (e.g., change the visualization type, zoom in on the chart, show and hide plots on the chart). Changes made on a share have no impact on source page settings or any other aspect of the original view.

> **Note:** A public share gives visitors (non-users) access to information and visualizations that may be proprietary or otherwise sensitive. Before sharing, be sure that the shared information is appropriate for sharing via a public link (no authentication required).

#### Public Share User Levels

The ability to create and manage public shares varies depending on your Kentik user level:

* Both members and administrators can create shares.
* Members can only modify/delete the shares they created.
* Administrators can modify/delete shares created by anyone in the organization.

> **Note:** Super Administrators have the same sharing privileges as Administrators.

#### Sharing Public Shares

The link to a public share can be shared with an individual or group via email or posted to a social media site such as LinkedIn, Twitter, or Slack. A public share link will unfurl when posted to social media, showing a preview that provides additional context.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(130).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A51Z&se=2026-05-12T09%3A32%3A51Z&sr=c&sp=r&sig=7m2nhLPaVz%2BIVgakST8Gqn9Gb1ZbQi51qpzcWaTBDxE%3D)

*An unfurled public share link in Slack.*

## Public Shares Page

The **Public Shares** page, which is accessed from the **Organization Settings** menu, is used to manage your organization's public shares.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(131).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A51Z&se=2026-05-12T09%3A32%3A51Z&sr=c&sp=r&sig=7m2nhLPaVz%2BIVgakST8Gqn9Gb1ZbQi51qpzcWaTBDxE%3D)

*Manage your public shares on the Public Shares page.*

### Public Shares List

The **Public Shares** page features a table that lists the shares that have been created by all Kentik users in your organization. Shares may be edited (title, description, and/or share link) or removed from this list.  
  
By default, the shares are listed in reverse chronological order. To change the sort order of the list, click a heading to select a column, and click the resulting blue up or down arrow to choose the sort direction (ascending or descending).

### Public Shares Columns

The**Public Shares** list includes the following columns (left to right):

* **Title**: The name given to the share when it was created or last edited.
* **Type**: The type of view at the shared link. The view type reflects the location at which the shared view was created (see **Sharable Views**).
* **URL**: A link to the public share page corresponding to the shared view. Two options are available:

  + Copy *(*icon): Copies the link to the Clipboard.
  + Open: Navigates to the public share version of the page from which the share was created.
* **Created At**: The date-time at which the public share was created.
* **Shared by User**: The Kentik user who created the public share.
* **Edit**: A button (pencil icon) that opens the **Link Dialog**, allowing you to review and edit the settings of the public share.

  > **Note**: While admin-level users can modify any shares, members can only modify the shares they’ve created.
* **Remove**: A button (trash icon) that opens a confirmation popup in which you can remove the share from your organization's collection of public shares.

  > **Note**: While admin-level users can remove any shares, members can only remove the shares they’ve created.

## Public Shares Dialogs

The dialogs used when working with public shares are covered in the following topics.

> **Note:** Before creating a share with the Share or Link dialog, be sure that the shared information is appropriate for sharing via a public link (no authentication)

### Share Dialog

The Share dialog is used to create a public share from a view in one of the modules listed in **Sharable Views**. The dialog opens from the **Share** button in the subnav of the view. Use the settings in the **Public Share** pane of the **Link** tab to create the public share (see **Public Share Link**). Kentik will automatically convert lookback settings (e.g., one day) to a fixed range (e.g., from 24 hours ago to the current time) and creates a public share using that time range.

> **Note:** Data will be returned only for the portion of the Kentik-generated time range that falls within the data retention period of your organization's Kentik plan (see **Device and Cloud Information**).

### Link Dialog

The Link dialog opens from the **Edit** icon at the right of a public share's row in the **Public Shares List**. The dialog is used to get an **Internal Share Link** for the share or to modify the share. The settings in this dialog are essentially the same as those of the **Link** tab of the Share dialog (see **Sharing by Link**), but the Link dialog doesn’t include the **Recipients** or **Message** fields.

![Public sharing options with a link and description for external users.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/PS-link-dialog.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A51Z&se=2026-05-12T09%3A32%3A51Z&sr=c&sp=r&sig=7m2nhLPaVz%2BIVgakST8Gqn9Gb1ZbQi51qpzcWaTBDxE%3D)

## Public Share Pages

A public share page is a page that contains data and visualizations from a shareable view. Because a public share page can be accessed without authentication, it enables you to share Kentik-generated information with people who aren't registered Kentik users or aren't in your Kentik organization.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(133).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A51Z&se=2026-05-12T09%3A32%3A51Z&sr=c&sp=r&sig=7m2nhLPaVz%2BIVgakST8Gqn9Gb1ZbQi51qpzcWaTBDxE%3D)

*A public share of a test results page from the Test Control Center in Synthetics.*

The destination page for a public share is not the original page from which the public share was created; instead, both its content and its UI are based on the original page, but its design is optimized for sharing outside the context of the Kentik portal. Visitors to a public share page will be able to see information from the source page and change how information is displayed on the share (e.g., change the visualization type, zoom in on the chart, show and hide plots on the chart). Changes made on a share have no impact on source page settings or any other aspect of the original view.  
  
The layout and UI elements on a public share page vary depending on the type of the share (see Type in **Public Shares Columns**), and, within synthetics, the test type (**Test Details by Test Type**).

## Managing Public Shares

The following topics cover basic operations involving public shares.

## Add a Public Share

To create a public share:

1. Navigate to the page that you'd like to share (see **Sharable Views** to see which types of pages are sharable).
2. Click the **Share** button in the subnav. The **Share Dialog** will open.
3. In the **Public Share** pane of the **Link** tab (see **Public Share Link**), set the **Title,** **Description,** and **Share Link**. You may also optionally specify recipients and text for an email containing the link.
4. Click **Share** to save the new share and close the dialog.
5. A notification will confirm that the share has been created. To see the public share, click the link in the URL column.

> **Note:** A query will return data only for the portion of the time range that falls within the data retention period of your Kentik plan (see **Device and Cloud Information**).

### Edit a Public Share

To edit an existing public share:

1. Open the Public Shares page from the **Organization Settings** menu.
2. In the **Public Shares List**, find the row corresponding to the share you’d like to edit, then click the Edit button (pencil icon) at the right of that row. The **Link Dialog** will open.
3. In the **Public Share** pane of the Link dialog (see **Public Share Link**), modify the **Title,** **Description,** or **Share** **Link** fields as needed.
4. Click the **Share** button to save the changes and close the dialog.

### Remove a Public Share

To remove an existing public share:

1. Open the Public Shares page from the **Organization Settings** menu.
2. In the **Public Shares List**, find the row corresponding to the share you’d like to remove.
3. Click the **Remove** button (trash icon) at the right of that row.
4. In the resulting confirmation dialog, click **Remove** to proceed with removal of the share.
5. A notification will appear, confirming that the share has been removed. The share will disappear immediately from the Public Shares list.

---

© 2014-25 Kentik
