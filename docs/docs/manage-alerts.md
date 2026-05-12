# Manage Alerts

Source: https://kb.kentik.com/docs/manage-alerts

---

Follow these steps to manage alerts in the Kentik portal.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(266).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A03Z&se=2026-05-12T09%3A37%3A03Z&sr=c&sp=r&sig=htKHIkjOV8vQxDkb3onVfSN47dHnGMU4U%2BH1pQQ5HTw%3D)

*Alerts can be managed from the Alerts list on the Alerting page.*

## About Alert Management

Alert management involves four main tasks:

* Acknowledging (acking) alerts, both manually and automatically (see **Acknowledging Alerts**).
* Clearing alerts (see **Alert State**).
* Commenting on alerts (see **Alert Comments**).
* Silencing alert notifications (see **Silence Notifications**).

## Alert Management Access

Access the following locations in the Kentik portal to manage alerts.

### Alert Management Locations

Alert management is done in the portal’s Alerting or DDoS Protect modules. The following table shows where to start the various tasks, while the topics afterward explain where to find those locations.

| Task | Alerts list (checkboxes) | Action menu | Details drawer | Details page sidebar |
| --- | --- | --- | --- | --- |
| **Acknowledge an Alert** | Yes | Yes | Yes | Yes |
| **Auto-acknowledge an Alert** | No | Yes | Yes | Yes |
| **Remove an Alert Ack** | No | Yes | Yes | Yes |
| **Clear an Alert** | Yes | Yes | Yes | Yes |
| **Add an Alert Comment** | No | Yes | Yes | Yes |
| **Edit an Alert Comment** | No | No | Yes | Yes |
| **Remove an Alert Comment** | No | No | Yes | Yes |
| **Silence Alert Notifications** | No | Yes | Yes | Yes |

### Access the Alerts List

Open the **Alerting Page** from the portal’s main menu to manage alerts directly from the **Alerts** list.

### Access the Action Menu

Open the Action menu for alerts in these portal locations:

* **Alerting page**: **Alerts List**.
* **DDoS Defense page**: Attacks Active Within the Last 24 Hours table (see **Attack Table**).

To access an alert’s Action menu:

1. Open either the DDoS Defense page or the Alerting page from the main nav menu.
2. Click the vertical dots icon for the alert to open its Action menu.

### Access the Details Drawer

Access the **Alert Details Drawer** on the following pages:

* **Alerting page**
* **DDoS Defense page**

To access an alert’s **Details** drawer:

1. Open either the DDoS Defense page or the **Alerting** page from the main nav menu.
2. Click an alert to open the **Details** drawer.

### Access the Details Page

To access an alert’s **Alert Details Page**:

1. Open either the DDoS Defense page or the Alerting page from the main nav menu.
2. Click the alert’s ID (**Alert ID** column) to open the alert’s Details page.

   > **Note:** If the **Alert ID** column isn’t shown, do one of the following:
   >
   > * Use the **Customize Columns Popup** to show the column, or;
   > * Click the vertical dots icon for the alert and select **View Alert Details**.

## Acknowledge an Alert

Acknowledge (ack) an alert from the portal locations shown in **Alert Management Locations**.

#### Ack Alert on Alerts List

To acknowledge alerts on the **Alerts** list:

1. Open the **Alerting** page from the main nav menu.
2. Check the boxes next to the alerts you want to ack. The **Alert Controls** will appear above the **Alerts** list.
3. Click **Acknowledge Alert**.

> **Note:** This method does not use the **Acknowledge Alert Dialog**, so there’s no opportunity to comment, auto-ack, or confirm acknowledgment of the alert.

#### Ack Alert via Action Menu

To acknowledge an alert from its Action menu (see **Alert-specific Actions**):

1. Open the alert’s Action menu (see **Access the Action Menu**) and select **Ack Alert**.
2. Click **Confirm** in the **Acknowledge Alert Dialog**.

#### Ack Alert in Details Drawer

To acknowledge an alert in its **Alert Details Drawer**:

1. Open the alert’s Details drawer (see **Access the Details Drawer**).
2. Click **Ack Alert** in the **Take Action** pane (see **Alert-specific Actions**).
3. Click **Confirm** in the **Acknowledge Alert Dialog**.

#### Ack Alert on Details Page

To acknowledge an alert on its **Alert Details Page**:

1. Go to the alert’s Details page (see **Access the Details Page**).
2. Click **Ack Alert** in the **Alert Details Page Sidebar**.
3. Click **Confirm** in the **Acknowledge Alert Dialog**.

> **Note**: If the History table is present in the **Threshold Data Pane** or **NMS Data Pane**of a Detail’s page, you can ack alerts with the **Ack Alert** button at the right of each row.

## Auto-acknowledge an Alert

When you acknowledge an alert using the **Acknowledge Alert Dialog**, you can auto-acknowledge (auto-ack) future instances of the same alert for a specified duration:

1. Open the dialog from the **Ack Alert** button in one of the following locations:

   1. Action menu: See **Access the Action Menu**.
   2. Details drawer: See **Access the Details Drawer**.
   3. Details page sidebar: See **Access the Details Page**.
2. Click the **Acknowledge additional occurrences (auto-ack)** checkbox.
3. Specify the duration of the auto-ack (see **Ack Alert Dialog UI**).
4. Click **Confirm**.

## Remove an Alert Ack

Remove your acknowledgement (ack) of an alert from the locations shown in **Alert Management Locations**.

> **Notes**:
>
> * You can only remove an ack for an alert that you acknowledged.
> * Removing an auto-ack is only possible from the **Auto-acknowledgements** page.

#### Remove Ack via Action Menu

To remove an ack using the alert’s Action menu (see **Alert-specific Actions**):

1. Open the alert’s Action menu (see **Access the Action Menu**).
2. Click **Remove Ack**.

#### Remove Ack in Details Drawer

To remove an ack in the alert’s **Alert Details Drawer**:

1. Open the alert’s **Details** drawer (see **Access the Details Drawer**).
2. Click **Remove Ack** in the drawer’s **Take Action** pane.

#### Remove Ack on Details Page

To remove an ack on an alert’s **Alert Details Page**:

1. Go to the alert’s **Details** page (see **Access the Details Page**).
2. Click **Remove Ack** in the page’s **Alert Details Page Sidebar**,

## Clear an Alert

Clear an active alert from the locations shown in **Alert Management Locations**.

#### Clear Alert via Alerts List

To clear alerts on the **Alerts List**:

1. Open the Alerting page from the portal’s main nav menu.
2. Check the boxes next to the alerts you’d like to clear. The **Alert Controls** will appear at the top left of the list.
3. Click **Clear Alert**.
4. Click **Confirm** in the confirmation popup.

#### Clear Alert in Details Drawer

To clear an alert from its **Alert Details Drawer**:

1. Open the alert’s **Details** drawer (see **Access the Details Drawer**).
2. Click **Clear Alert** in the **Take Action** pane.
3. Click **Confirm** in the confirmation popup.

#### Clear Alert on Details Page

To clear an alert from its **Alert Details Page**:

1. Go to the alert’s **Details** page (see **Access the Details Page**).
2. Click **Clear Alert** in the **Alert Details Page Sidebar**.
3. Click **Confirm** in the confirmation popup.

## Add an Alert Comment

Add an alert comment in the locations shown in **Alert Management Locations**.

> **Note:** Comments are displayed in the locations described in **Alert Comments**.

#### Add Comment from Action Menu

To add a comment from an alert’s Action menu (see **Alert-specific Actions**):

1. Open the alert’s Action menu (see **Access the Action Menu**).
2. Click **Add** **Comment**.
3. Enter your comment in the **Add a Comment** field and click **Save**.

   > **Note**: The field is expandable with the handle in its lower right corner.

#### Add Comment on Details Drawer

To comment on an alert in its **Alert Details Drawer**:

1. Open the alert’s **Details** drawer (see **Access the Details Drawer**).
2. Enter your comment in the **Comments** field.
3. Click the blue arrow at the right to add the comment to the alert.

#### Add Comment on Details Page

To comment on an alert on its **Alert Details Page**:

1. Go to the alert’s **Details** page (see **Access the Details Page**).
2. Enter your comment in the **Comments** field in the **Alert Details Page Sidebar**.
3. Click the blue arrow at the right to add the comment to the alert.

#### Add Comment in Ack Alert Dialog

To add a comment when acknowledging an alert in the **Acknowledge Alert Dialog**:

1. Open the dialog from the **Ack Alert** button in one of the following locations:

   1. Action menu: See **Access the Action Menu**.
   2. Details drawer: See **Access the Details Drawer**.
   3. Details page sidebar: See **Access the Details Page**.
2. Enter your comment in the **Comments** field and click **Confirm**.

## Edit an Alert Comment

Edit your comments on an alert in the locations shown in **Alert Management Locations**.

> **Note**: You cannot edit another user’s comments.

#### Edit Comment in Details Drawer

To edit a comment in an alert’s **Alert Details Drawer**:

1. Open the alert’s **Details** drawer (see **Access the Details Drawer**).
2. In the **Details** drawer, find the Comments card for the comment you’d like to change.
3. Click **Edit** and make the needed changes.
4. Click **Save**.

#### Edit Comment on Details Page

To edit a comment in an alert’s **Alert Details Page**:

1. Go to the alert’s **Details** page (see **Access the Details Page**).
2. In the **Alert Details Page Sidebar**, find the Comments card for the comment you’d like to change.
3. Click **Edit** and make the needed changes.
4. Click **Save**.

## Remove an Alert Comment

Remove your comment from an alert in the locations shown in **Alert Management Locations**.

> **Note**: You cannot remove another user’s comments.

#### Remove Comment on Details Drawer

To remove your comment from an alert on its **Alert Details Drawer**:

1. Open the alert’s **Details** drawer (see **Access the Details Drawer**).
2. In the **Details** drawer, find the Comments card for the comment you’d like to remove.
3. Click **Remove** at the top right.
4. Click **Remove** in the confirmation popup.

#### Remove Comment on Details Page

To remove your comment from an alert on its **Alert Details Page**:

1. Go to the alert’s **Details** page (see **Access the Details Page**).
2. In the **Alert Details Page Sidebar**, find the Comments card for the comment you’d like to remove.
3. Click **Remove**.
4. Click **Remove** in the confirmation popup.

## Silence Alert Notifications

Silence an alert’s **Notifications** for seven days from the locations shown in **Alert Management Locations**.

> **Note**: To set a custom silence duration when acknowledging an alert, see **Alert-specific Actions**.

#### Silence Alert using Action Menu

To silence an alert using the alert’s Action menu (see **Alert-specific Actions**):

1. Open the alert’s Action menu (see **Access the Action Menu**).
2. Click **Silence Notifications**.

#### Silence Alert in Details Drawer

To silence notifications for an alert in the alert’s **Alert Details Drawer**:

1. Open the alert’s **Details** drawer (see **Access the Details Drawer**).
2. Click **Silence Notifications** in the **Take Action** pane.

#### Silence Alert on Details Page

To silence an alert’s notifications on its **Alert Details Page**:

1. Go to the alert’s **Details** page (see **Access the Details Page**).
2. Click **Silence Notifications** in the **Alert Details Page Sidebar**, .

#### Custom Silence Alert

Silence an alert’s notifications for a custom duration when acknowledging an alert in the **Acknowledge Alert Dialog**:

1. Open the dialog from the **Ack Alert** button in one of the following locations:

   1. Action menu: See **Access the Action Menu**.
   2. Details drawer: See **Access the Details Drawer**.
   3. Details page sidebar: See **Access the Details Page**.
2. Check the box for **Silence notifications for this alert**.
3. Set the desired duration using the **Duration** controls (see **Ack Alert Dialog UI**).
4. Click **Confirm** to silence notifications, acknowledge the alert, and close the dialog.

---

© 2014-25 Kentik
