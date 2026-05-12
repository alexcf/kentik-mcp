# Mitigation Methods

Source: https://kb.kentik.com/docs/mitigation-methods

---

On Kentik’s **Manage Mitigations**page, a **Mitigation Method** defines the specific rules, traffic thresholds, and actions taken when a Kentik alert triggers an active mitigation. Methods act as the logical instructions that run on your assigned **Mitigation Platforms**.

This article covers the **General** tab settings (common to all mitigation methods) and the **step-by-step processes for managing methods**. For detailed configuration steps specific to the type of mitigation you are building, see the following guides:

* **RTBH Method Settings**
* **Flowspec Method Settings**
* **Adaptive Flowspec Method Settings**
* **Third-Party Method Settings**

## Mitigation Method Settings

Adding or editing a mitigation method via the Kentik portal involves specifying information in the fields of the mitigation method dialogs.

> **Note:** In addition to configuring a mitigation platform and method in Kentik, you must also allow the following IP ranges on third-party mitigation platforms (e.g., Radware, A10) and on any devices used for Flowspec, RTBH, or Adaptive Flowspec mitigations:
>
> * **US**:  `209.50.158.0/23 (IPv4)` and `2620:129:1::/48 (IPv6)`
> * **EU**: `193.177.128.0/22 (IPv4)`

### Mitigation Method Dialogs

The settings for mitigation methods are managed in the following dialogs:

* **Add Mitigation Method** when registering a new method with Kentik.
* **Edit Mitigation Method** when editing an already registered method.

### Method Dialogs UI

The Add Mitigation Method and Edit Mitigation Method dialogs share the same layout and the following common UI elements:

* ![Settings for automated mitigations including acknowledgment and grace period options.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/MM-edit-mitigation-method-dialog.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A14Z&se=2026-05-12T09%3A35%3A14Z&sr=c&sp=r&sig=Nknz%2BAMdTLGHAaJuNoCJGK%2FhIEzGGxNwJyg4%2BCjh410%3D)**Close** (button): Click the **X** in the upper right corner to close the dialog. All elements will be restored to their values at the time the dialog was opened.
* **General** (tab): Settings common to all mitigation platforms (see **General Method Settings**).
* **Details** (tab): Configurations specific to each platform (see **RTBH Method Settings**, **Flowspec Method Settings**, **Adaptive Flowspec Method Settings**, and **Third-Party Method Settings**.
* **Cancel** (button): Cancel the add method or edit method operation and exit the dialog. All elements will be restored to their values at the time the dialog was opened.
* **Add Mitigation Method** (button, Add Mitigation Method dialog only): Save settings for the new method and exit the dialog.
* **Update Mitigation Method** (button, Edit Mitigation Method dialog only): Save changes to method settings and exit the dialog.

  > **Note:** Cloudflare Magic Transit does not have separate tabs. The only options available are the **General Method Settings**.

### General Method Settings

The **General** tab of the mitigation method dialogs has settings and controls that are common to all mitigation method types.

* **Name**: User-specified name for the mitigation method.
* **Description**: Optional user-provided method description.
* **Notification Channels**: Choose one or more notification channels for the mitigation method from the dropdown. Notification channels are created on the Notifications page (see **Notifications**).
* **Settings for Automated Mitigations**: Settings that apply only to automated mitigations, and not to manual mitigations (see **Automated Mitigation Settings**).

#### Automated Mitigation Settings

The following settings are applicable only to automated mitigations, which are triggered in response to an alarm (see **Threshold Mitigations**):

* **Acknowledgement Required**: If this switch is on, a mitigation alarm from this method must be manually (rather than automatically) cleared from the Alerting page (select **Alerting** from the main menu) before the mitigation can proceed.
* **IPs/CIDRs Excluded From Mitigation**: IP addresses that should be excluded from being mitigated with this method, for example infrastructure addresses, point-to-point networks, or other addresses critical to the normal functioning of your network. Enter as a comma-separated list.
* **Grace period**: The grace period that Kentik should honor prior to ending mitigation (e.g. withdrawing a blackhole route). Default is 30 minutes.

> **Note:** Automated mitigation settings have no effect on manual mitigations.

## Manage Mitigation Methods

Methods are added, edited, and removed via the Manage Mitigations page (**Settings »** **Manage Mitigations**).

### Add a Mitigation Method

A mitigation method can only be added to a mitigation platform of the same type (see **Add a Mitigation Platform**). To add a new mitigation method to a platform:

1. Within a platform pane on the Manage Mitigations page (**Settings »** **Mitigations**), click **Add Mitigation Method**.
2. On the **General** tab, specify general properties of the method (see **General Method Settings**).
3. On the **Details** tab, fill in the details fields, which vary by method type:

   1. **RTBH Method Settings**
   2. **Flowspec Method Settings**
   3. **Adaptive Flowspec Method Settings**
   4. **Third-Party Method Settings**

      > **Note:** Cloudflare methods have no **Details** tab.
4. Save the new method by clicking the **Add Mitigation Method** button (lower right).

### Edit a Mitigation Method

The settings for an existing mitigation method may be edited from the following locations:

* On the Manage Mitigations page, from the **Mitigation Methods List** in the pane of a platform to which the method is assigned.
* On an Add Mitigation Platform or Edit Mitigation Platform page, from the list of methods under Assign Mitigation Methods (see **Common Platform Settings**).

In one of the lists above:

1. Click the edit icon (pencil) in the row of the method that you'd like to edit. The Edit Mitigation Method dialog will open.
2. To change general properties of the method, use the **General** tab (see **General Method Settings**).
3. To change settings that are specific to the method's mitigation type, use the **Details** tab:

   1. **RTBH Method Settings**
   2. **Flowspec Method Settings**
   3. **Adaptive Flowspec Method Settings**
   4. **Third-Party Method Settings**

      > **Note:** Cloudflare methods have no details to edit.
4. To save changes, click the **Update Mitigation Method** button (lower right).

   > **Note:** If you save the changes to your mitigation method, the method will be updated whether or not you saved any changes to the mitigation platform.

### Remove a Mitigation Method

**To unassign a mitigation method from an existing platform**:

1. On the Manage Mitigations page, click **Edit** at the upper right of the pane corresponding to a platform to which the method has been assigned.
2. On theEdit Mitigation Platform page, click the **X** at the left of the method under Assign mitigation Methods (see **Common Platform Settings**).

   > **Note:** The method is unassigned from that platform only and with no confirmation popup.
3. Click **Update Mitigation Platform**.

**To remove a method from your organization's collection of mitigation methods**:

1. In one of the following locations, click the Remove icon (red trash can) next to the mitigation method that you'd like to delete:

   1. **Manage Mitigations page**: In the **Mitigation Methods List** of a mitigation platform pane.
   2. **Add Mitigation Platform or Edit Mitigation Platform page**: In the list under Assign Mitigation Methods (see **Common Platform Settings**).
2. In the resulting confirmation popup, click **Remove**.

   > **Note:** The method will be removed from all of the platforms to which it is assigned.

### Summary and Next Steps

You now know how to configure the logical rules and routing behaviors for your mitigations.

**Next Steps:** Methods are most powerful when they are automated. Now that your methods are built and assigned to active platforms, your final step is to tie them to an alert policy so they trigger automatically when an attack occurs. See the **Alert Policies** article to learn how to assign your mitigation methods to your alerts.
