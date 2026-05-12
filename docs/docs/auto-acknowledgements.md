# Auto-acknowledgements

Source: https://kb.kentik.com/docs/auto-acknowledgements

---

This article discusses **Auto-acknowledgements** in the Kentik portal.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(583).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A28Z&se=2026-05-12T09%3A35%3A28Z&sr=c&sp=r&sig=MtbBUH9wGFclVDr%2F3UmgaplyWrc3EugzyztguroPqu4%3D)

*The Auto-acknowledgments page.*

## About Auto-acknowledgement

Acknowledging an alert indicates to other users that you are aware of that alert. With auto-acknowledgement, you can set a duration for which all subsequent instances of a given alert (triggered by a given policy threshold and based on the same key) will be acknowledged automatically. During this duration, the **Ack State** of the alert will be indicated as "Auto-acked" in all locations where **Ack state** is displayed, and the alert will be listed on the **Auto-acknowledgements Page**.  
  
The duration for auto-acknowledgement is set in the **Acknowledge Alert Dialog** when you **Auto-acknowledge an Alert**. The minimum duration is one hour; the maximum is seven days for member-level users or one year for admin-level users. This duration can be changed on the Auto-acknowledgements page.

## Auto-acknowledgements Page

The Auto-acknowledgements (Auto-acks) page is covered in the following topics.

### Auto-acks Access

To access the Auto-acknowledgements (Auto-acks) page, from the portal’s main menu, either:

* Hover over **Settings** to view the **Settings** submenu, then select **Auto-acknowledgements**, or
* Click **Settings** to open the Settings page, then click **Auto-acknowledgements** (under Alerting).

### Auto-acks Page UI

The Auto-acknowledgements (Auto-acks) page has the following UI elements:

* **Breadcrumbs** (on subnav): An indicator of your current location within the Kentik portal.
* **Favorite**: A star to the left of the page title that allows you to add this page to the **Favorites Tab** of the **Portal Search**.
* **Auto-ack list**: A table listing all auto-acks (current and expired) within your organization (see **Auto-ack List**).

### Auto-ack List

The rows of the **Auto-ack** list each represent an alert for which a user in your organization has created an auto-acknowledgement (auto-ack). The list can be sorted (ascending or descending) by clicking the headings of the table's **Policy**, **Expires**, or **Created by** columns.  
  
The list provides the following information and actions for each auto-acknowledgement:

* **Matching**: The type and conditions (dimension/value pairs) that triggered the original alert from which the auto-ack was created (i.e. the policy dimensions for a threshold alert, or the status for an NMS alert).
* **Comment**: The comment, if any, entered for the auto-ack when it was created. The user who added the comment will see the following buttons below it:

  + Edit: A button that allows you to edit the comment.
  + Remove: A button that opens a confirmation dialog in which you can remove the comment. Once you remove a comment, you can't add another to the same auto-ack.

    > **Note**: Changing a comment in the Auto-ack list will not affect the alert's comments in the **Alert Details Drawer** or the **Alert Details Page Sidebar**.
* **Policy**: The policy of the alert from which the auto-ack was created. Click the link to go to the alert’s **Edit****Policy** page.
* **Expires**: The date-time of the end of the auto-acknowledgment duration for this alert, which displays in red if the date-time is in the past. Beneath the date-time is one of the following:

  + The remaining duration, in parentheses (if the auto-ack is active);
  + How long ago the auto-acked expired (if past).
* **Created by**: The user who created the auto-ack.
* **Edit** (pencil icon): A button that opens the Edit Auto-acknowledgment dialog (see **Edit Auto-ack Dialog**), where any user in your organization can specify a different ending date-time.
* **Remove** (trash icon): A button that opens a confirmation dialog that allows you to delete the auto-ack.

## Edit Auto-ack Dialog

The Edit Auto-acknowledgement dialog enables you to change the date-time at which an auto-acknowledgement (auto-ack) will expire. Accessed by clicking the **Edit** button at the right of each row in the **Auto-ack** list, the dialog contains the following UI elements:

* ![Editing auto-acknowledgement expiration date in a user interface calendar view.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/AAK-edit-auto-ack.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A28Z&se=2026-05-12T09%3A35%3A28Z&sr=c&sp=r&sig=MtbBUH9wGFclVDr%2F3UmgaplyWrc3EugzyztguroPqu4%3D)**Cancel**: A button, **X** at top right or **Cancel** at bottom,that closes the dialog without changing the expiration date-time of the auto-acknowledgement.
* **Auto-acknowledgement expires**: A field that shows the 24-hour date-time at which the auto-acknowledgement is set to expire. Clicking the field will pop up a calendar:

  + To change the expiration date, edit the field or select a new date from the calendar.
  + To change the time, edit the field.
* **Expiry notification** (only when the date-time is past): An orange box that tells you that the expiry date is in the past.
* **Save**: A button that saves your changes and exits the dialog.

> **Notes**:
>
> * Member-level users can set an auto-ack to a maximum of seven days in the future. Admin-level users can set an auto-ack of up to one year.
> * If you enter a date-time that's already passed the change will not be saved and an error message will be shown when the dialog is closed.

## Manage Auto-acks

Once an auto-acknowledgement is created (see **Auto-acknowledge an Alert**) you can perform the following tasks on the **Auto-acknowledgements Page****.**

### Edit an Auto-ack

You can edit or remove an auto-ack’s comments or adjust its expiration date-time from the **Auto-acknowledgements Page**, which you can access from the **Settings** submenu of the portal’s main menu (see **Portal Main Menu**).

> **Note:** You cannot edit the conditions that will trigger an auto-acknowledged alert.

#### Edit a Comment

To edit an auto-ack’s comment:

1. In the **Auto-ack** list on the Auto-acknowledgements page, find the auto-ack comment that you’d like to modify, then click **Edit**.
2. In the now-editable **Comment** field, make the necessary changes.
3. Click the **Save** button.

> **Note**: Only the user who added the auto-ack’s comment can edit it.

#### Remove a Comment

To remove an auto-ack’s comment:

1. In the **Auto-ack** list on the Auto-acknowledgements page, find the auto-ack comment that you’d like to remove, then click **Remove**.
2. In the confirmation dialog, click **Remove**. The Comment column for the auto-ack will now display two dashes in place of the removed comment.

> **Notes**:
>
> * Only the user who added the comment can remove it.
> * Once an auto-ack’s comment is removed, you cannot add another comment.

#### Reset Auto-ack End Date

To change the expiry of an auto-acknowledgement:

1. In the **Auto-ack** list on the Auto-acknowledgements page, find the auto-ack whose expiry date you’d like to modify, then click the **Edit** button (far right of the row). The **Edit Auto-ack Dialog** will open.
2. Click the **Auto-acknowledgement expires** field to display a popup calendar from which you can set the duration for the auto-ack.

   > **Note**: Member-level users can set the duration to a maximum of seven days in the future and admin-level users can set the duration for up to one year (365 days).
3. Click the **Save** button to save the auto-ack and close the dialog.

### Remove an Auto-ack

To remove an auto-acknowledgement (auto-ack) from the **Auto-Ack** list:

1. In the **Auto-ack** list on the Auto-acknowledgements page, find the auto-ack you’d like to remove.
2. Click the **Remove** button.
3. Click the **Remove** button in the resulting confirmation popup. The auto-ack will be permanently removed from your organization.
