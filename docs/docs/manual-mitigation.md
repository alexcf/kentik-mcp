# Manual Mitigation

Source: https://kb.kentik.com/docs/manual-mitigation

---

This article discusses manual mitigation in Kentik.

> **Notes:**
>
> * For information on the v4 portal's Mitigations page, see **Mitigations**.
> * For a high-level overview of mitigation, see **About Mitigation**.
> * For mitigation methods and platforms, see **Manage Mitigations**.

## About Manual Mitigation

Kentik's mitigation system lets you start and stop mitigations not only as automated responses triggered by alert policies (see **Threshold Mitigations**), but also manually (independent of an alert). A mitigation can be set to stop automatically based on TTL (see **Manual Mitigation Settings**) or to continue until stopped manually (see **Manual Mitigation Actions**) from the Mitigations page, where it will be listed once it's started.

> **Note:** Manual mitigation life cycle states are covered in **Mitigation Status**.

#### Accessing Manual Mitigation

Manual mitigation is exposed in the portal via the **Start Manual Mitigation Dialog**, which is accessed via the Start Manual Mitigation button on the following pages:

* **Mitigations** page at Protect » **Mitigations**
* **Manage Mitigations** page at Settings » **Mitigations**

## Start Manual Mitigation Dialog

The **Start Manual Mitigation** dialog is covered here.

### Manual Mitigation Dialog UI

The **Start Manual Mitigation** dialog includes these general UI elements:

* **Cancel**: A button — **X** at upper right or **Cancel** at bottom right — that exits the dialog without starting a mitigation.
* **Add Manual Mitigation**: A button that starts the mitigation.

> **Note:** Kentik requires 1-2 minutes for provisioning after creating or editing a mitigation method or platform. A mitigation using that platform or method cannot be applied during that time.

### Manual Mitigation Settings

The **Start Manual Mitigation** dialog includes the following settings and controls for manual mitigations:

* **Mitigation Method**: A drop-down from which you can select an existing platform-method combination from a filterable list.

  > **Note:** The list includes only methods that are linked to a platform in the **Mitigation Methods** field of the **Edit Mitigation Platform** page (see **Common Platform Settings**).
* **View Method Details** (Flowspec methods only): A button (zoom-in icon) that opens the **Mitigation Method Details** dialog.
* **IP/CIDR to Mitigate**: A field in which to enter the IP range for mitigation.
* **Source Port to Mitigate** (Flowspec methods only): A field in which to enter the number of the source port.
* **Destination Port to Mitigate** (Flowspec methods only):  A field in which to enter the number of the destination port.
* **Comment**: A field in which to enter an optional comment string (reserved for future use).
* **Minutes Before Auto-Stop (TTL)**: A field in which to enter, in minutes, the duration for which the mitigation will run. If left at 0 (default), the mitigation must be stopped manually on the Mitigations page (see **Stop a Manual Mitigation**).
* **Matching mitigations**: A field whose value is "No matching mitigations found" unless the current value of the **IP/CIDR to Mitigate** field is valid and a mitigation exists on that IP/CIDR, in which case the field shows the matching mitigation's status, ID, and target.

  > **Note:** Only one mitigation is allowed per IP/CIDR.

> **Notes**:
>
> * For Flowspec mitigations, see **Manual Mitigation with Flowspec**.
> * To add a platform or method, see **Add a Mitigation Platform** or **Add a Mitigation Method**).

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(377).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A21Z&se=2026-05-12T09%3A34%3A21Z&sr=c&sp=r&sig=98Co8670d23I2LGs7P6tBKdCqwNBt1lKW6OPw5nsjWE%3D)

#### Manual Mitigation with Flowspec

Flowspec-based manual mitigations are based on Flowspec mitigation methods, which are configured in the **Mitigation Method Dialogs**. Flowspec configuration includes settings in the Traffic Matching pane of the dialog's Details tab. As shown in the table below, several of the **Flowspec Component Types** for which settings are made in that pane correspond to fields in the **Start Manual Mitigation** dialog.

| Flowspec Component types | Start Manual Mitigation fields |
| --- | --- |
| * Source IP/CIDR * Destination IP/CIDR  (only one or the other) | IP/CIDR to Mitigate |
| Source Ports | Source Port to Mitigate |
| Destination Ports | Destination Port to Mitigate |

When the **Infer From Alarm** switch is on for the component types listed above and a mitigation is triggered automatically by an alert policy threshold, the values of those components are derived from the conditions that triggered the threshold. In a manual mitigation, however, there is no threshold information, so those values must be set manually in the fields of the **Start Manual Mitigation** dialog. The table below shows how the **Infer from Alarm** switches in the method's **Traffic Matching** pane affect the settings in the **Start Manual Mitigation** dialog.

| Condition group settings | Start Manual Mitigation dialog |
| --- | --- |
| Infer from Alarm switch is on. | Field is present; enter a value. |
| * Infer from Alarm switch is off. * IP or port is specified. | Field is locked or absent. |

> **Notes:**
>
> * Only one Flowspec mitigation with a static IP/CIDR can be active at a time.
> * A Flowspec mitigation won’t appear in the **Mitigation Method** menu if the **Infer From Alarm** switch is on for a Protocol condition group (see **Protocol and Port Components**) in the method configuration.

#### Details of mitigation methods including inferred destination IP and source ports.Mitigation Method Details

The **Mitigation Method Details** dialog appears when you select a Flowspec method in the **Start Manual Mitigation Dialog** and then click **View Method Details**.

* The dialog shows the method’s name and possible default values that can be associated with it.
* While it cannot be edited, the dialog shows if any of the 11 dimensions listed are inferred, not considered, or specified with certain default values. This information is included in the mitigation announcement.
* The dialog also includes a **Close** button (**X**) in the upper right corner.
* This dialog is also accessible via the Thresholds tab of an Add/Edit Policy page after adding a Flowspec method to the policy (see **Threshold Mitigations**).

## Start a Manual Mitigation

To start a manual mitigation:

1. As described in **Accessing Manual Mitigation**, navigate to one of the following pages:

   1. Protect » **Mitigations**
   2. Settings » **Mitigations**
2. Click the **Start Manual Mitigation** button to open the dialog.
3. Specify settings as per **Manual Mitigation Settings**.
4. Click the **Add Manual Mitigation** button. The manual mitigation starts immediately and appears in the **Mitigations List**.

> **Note:** Settings in the dialog are not saved for later reuse because manual mitigation is intended for one-off use only.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(379).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A21Z&se=2026-05-12T09%3A34%3A21Z&sr=c&sp=r&sig=98Co8670d23I2LGs7P6tBKdCqwNBt1lKW6OPw5nsjWE%3D)

*Once started, a manual mitigation appears in the Mitigations list.*

## Stop a Manual Mitigation

To stop a manual mitigation:

1. Go to the Protect » **Mitigations** page as described in **Accessing Manual Mitigation**.
2. In the **Mitigations List**, locate the row for the manual mitigation.
3. Click the **Action** menu (vertical ellipsis) at the right of the row.
4. Click the **Stop** button (gray square) in the popup menu.

---

© 2014-25 Kentik
