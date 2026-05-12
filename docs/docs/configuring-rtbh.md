# Configuring RTBH

Source: https://kb.kentik.com/docs/configuring-rtbh

---

**Remotely Triggered Black-Hole** (RTBH) routing is a powerful technique used to protect network infrastructure against DDoS attacks by dropping undesirable traffic at the edge.

This guide walks you through the end-to-end process of configuring an RTBH mitigation within the Kentik portal, from setting up your initial BGP communities to deploying the automated response.

> **Note:** The settings for RTBH platforms and methods are made on the portal's **Manage Mitigations** page (**Settings »** **Manage Mitigations**).

## Before You Begin

* **Review core concepts:** If you are new to Kentik Protect, review the **Mitigation Overview** to understand how platforms, methods, and automated triggers work.
* **Allow Kentik IPs:** Ensure your edge routers are configured to accept BGP announcements from Kentik's infrastructure. You must allow the following IP ranges:

  + **US**: `209.50.158.0/23` (IPv4) and `2620:129:1::/48` (IPv6)
  + **EU**: `193.177.128.0/22` (IPv4)

## IP Families in RTBH Mitigation

An **RTBH Mitigation** in Kentik involves using BGP to instruct devices that are part of a mitigation platform to redirect traffic destined for a given “target” IP/CIDR. To mitigate traffic to a given target, both of the following must be true:

* The BGP session between Kentik and each device assigned to the mitigation platform (see **Configure Platform Devices**), which is set on the **BGP** tab of **Edit Device** dialog in **Settings »** **Networking** **Devices** (see **Device BGP Settings**), must be established in the same address family (IPv4 or IPv6) as the target IP/CIDR of the traffic that will be mitigated.
* The address family (IPv4 or IPv6) of the **Next Hop** setting of the RTBH mitigation method (see **RTBH Method Details**) must match that of the target IP/CIDR.

Kentik applies the above requirements as follows:

* **Manual mitigation**: In the **Start Manual Mitigation** dialog, the family of the address you enter in the **IP/CIDR to Mitigate** field is checked against the family of the **Next Hop** setting of the RTBH mitigation method you chose with the **Mitigation Platform** **and** **Method** drop-down (see **Manage Mitigations**). You will only be allowed to start the mitigation if the address families match.
* **Automated mitigation**: When a user associates a mitigation with an alert policy threshold (see **Threshold Mitigations**) it can’t be known in advance what addresses will trigger the mitigation. It is therefore up to the user to do one of the following to ensure that the mitigation is properly configured (and thus won’t fail when triggered):

  + Create family-specific policies using associated filters and mitigation options.
  + Associate multiple RTBH mitigations with the threshold to cover both IPv4 and IPv6.

    > **Note:** This may result in attempted announcements on the wrong BGP sessions. Ask Kentik for assistance (see **Customer Care**).

## RTBH Mitigation Workflow

The following workflow outlines the general process of creating, configuring, and deploying an **RTBH Mitigation**. If this is your first time working through the process, we recommend that you contact Kentik before starting so that we can assist you (see **Customer Care**).

> **Note:** The following IP ranges must be allowed on devices that will be used for RTBH mitigations:
>
> * **US**:  `209.50.158.0/23 (IPv4)` and `2620:129:1::/48 (IPv6)`
> * **EU**: `193.177.128.0/22 (IPv4)`

### Identify and Enable Routers

To prepare routers for use in RTBH mitigation:

1. In the Kentik portal, go to **Settings »** **Networking Devices**.
2. In the **Device List**, check the **Status** column (BGP V4 and/or V6; see screenshot below) for each device on which you wish to implement RTBH mitigation (if the icon[s] are present, skip to step 7). If a given device doesn’t have a check mark in that column, click the Edit icon of the device’s row to open the **Edit Device** dialog.
3. In the resulting **Edit Device** dialog, go to the **BGP** tab.
4. Choose Peer with Device from the **BGP Type** drop-down menu.
5. Click the **Save** button to save changes and exit the dialog.
6. Repeat steps 2 through 5 to enable direct peering with the devices that you’d like to use for RTBH mitigation (if they aren’t already BGP-enabled).
7. Make a note of the names of the RTBH-ready devices.

![Table displaying network device statuses, IP addresses, and BGP prefix information.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(340).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A38Z&se=2026-05-12T09%3A37%3A38Z&sr=c&sp=r&sig=LDLYgDzGqxd13ZEi10%2FrjEWl%2FC6lY1v%2Fxo5Q3r9%2FI5k%3D)

*The Devices list in Settings » Network Devices shows the BGP status of devices.*

### Create an RTBH Mitigation Platform

Once you’ve enabled devices for RTBH, you can create a platform, with which the method will combine to make a mitigation:

1. Go to the Manage Mitigations page (**Settings »** **Mitigations).**
2. Click the **Add Mitigation Platform** button.
3. On the Add Mitigation Platform page, select **RTBH** for your mitigation type.
4. Fill in the common settings for your new platform (see **Mitigation Platform Settings**).
5. Click **Select Devices** to open the **Data Sources** dialog, then choose the routers that you enabled for RTBH in part A of this workflow.
6. Click **Create Mitigation Platform** to save your new platform.

### Create an RTBH Mitigation Method

To create a mitigation method for RTBH:

1. Go to the Manage Mitigations page (**Settings »** **Mitigations**).
2. On your new RTBH platform pane (created in step B), click the **Add Mitigation Method** button.
3. On the **General** tab of the resulting dialog (see **Mitigation Method Dialogs**), give your method an informative name and description.
4. From the drop-down **Notification Channels** menu, assign a notification channel so that you can be notified when your alert triggers a mitigation.

   > **Note:** If you have not already configured a notification channel, go to the Notifications page to add yourself to an existing channel or create a new channel for this mitigation method.
5. Turn on the **Acknowledgement Required** switch, which means that after a mitigation from this method has been triggered, the corresponding alarm must be cleared manually (rather than automatically) from the **Alerts List** on the Alerting page. Once you have acknowledged the alert, the mitigation can then proceed.
6. Use the **IP/CIDRs Excluded** field to enter any IP address(es) that you’d like to exclude from being blackholed with this method. Good candidates might be infrastructure addresses, point-to-point networks, or other addresses critical to the normal functioning of your network.
7. Select the **Grace Period** that Kentik should honor prior to withdrawing the blackhole route. Many operators are happy with the 30-minute default because it provides enough cushion to discourage repeat attacks while not being excessively punitive to the IP that was the attack destination.

### Specify RTBH Details

To configure the new mitigation method for RTBH:

1. Switch to the dialog's **Details** pane and in the **Community to Advertise** field, enter the community that has been programmed on the customer's router to induce a black hole next hop for the IPv4 address attached to the community.
2. For each **Next Hop** field (IPv4 and IPv6), enter a next-hop IP address. In some environments this will be the destination IP to blackhole. This number has traditionally been selected from the 192.0.2.0/24 CIDR block but may be any IP address.
3. Use the **Local Preference** field (see **RTBH Method Details**) to set the priority for the RTBH route.
4. If you plan to withdraw blocks from certain routers and re-advertise in other locations, you may want to turn on the **Ensure at least /24** checkbox. Otherwise, leave it off.
5. When you’re finished, click the **Add Mitigation Method** button to save the new method and exit the dialog.

### Link Existing Methods to an RTBH Platform

If your system has existing RTBH mitigation methods, you can also add them to your platform:

1. Click **Edit** in your RTBH mitigation platform pane (top right).
2. Under Assign mitigation methods, click to open a drop-down list of RTBH-based mitigation methods. Select any methods you would like to add to the platform. Already linked methods appear in blue.

   > **Note:** By linking the method to the platform, you create a mitigation that can be automatically or manually applied.
3. Click the **Update Mitigation Platform** button to save the platform and return to the Manage Mitigations page.

### Assign RTBH Mitigation to an Alert Policy

Once configured, a mitigation can be deployed for either automated mitigation (initiated when an alarm is triggered by a threshold in an alert policy) or manual mitigation (described in **Start a Manual Mitigation**). To deploy your new RTBH mitigation automatically, assign the mitigation to an alert threshold (see **Threshold Mitigations**):

![Settings for activating mitigations to protect network availability from undesirable traffic.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(341).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A38Z&se=2026-05-12T09%3A37%3A38Z&sr=c&sp=r&sig=LDLYgDzGqxd13ZEi10%2FrjEWl%2FC6lY1v%2Fxo5Q3r9%2FI5k%3D)

1. Go to the Policies page (from the main menu, go to the **Alerting** page and click the **Configure Alert Policies** button).
2. In the **Alert Policies List**, find the policy to which you’d like to assign an RTBH mitigation.
3. Click the Action menu and select **Configure Policy** from the drop-down menu.
4. On the **Edit Alert Policy** page, go to the **Thresholds** tab. Click one of the tabs to select the threshold (Critical, Major, etc.) to which you’d like to assign mitigation.
5. Under **Actions,** click to display the drop-down below **Mitigations** to show a list of mitigations (each displayed as "platform - method"). Choose the mitigation that you created in tasks B to E above, then click the **Add Mitigation** button.
6. In the individual mitigation pane that appears under **Mitigations,** click to open the **Apply Mitigation** drop-down and choose when you’d like to have the mitigation take effect, which could be:

   1. Immediately when Kentik raises the alarm.
   2. After a user manually acknowledges the alarm,
   3. After a user manually acknowledges the alarm or after a timer expires where no user has acknowledged the alarm.
7. Click the **Clear Mitigation** drop-down to choose when the alarm about the mitigation should be cleared from the **Mitigations List**.
8. Click the **Save** button (top right).

   > **Note:** The creation and settings of policies to which you can assign mitigations is covered in **Alert Policies**.

## Next Steps

Once your RTBH mitigation is active, you can monitor the specific routes being advertised to your routers on the **BGP Announcements** dashboard.
