# Configuring Flowspec

Source: https://kb.kentik.com/docs/configuring-flowspec

---

BGP Flow Specification (**Flowspec**), defined in **RFC 8955**, is a highly granular mitigation technique that allows you to distribute traffic filtering rules directly to your routing infrastructure via BGP extended communities.

This guide walks you through the complete process of creating, configuring, and deploying a **Flowspec Mitigation** within the Kentik portal, from preparing your edge routers to triggering automated filtering actions.

## Before You Begin

* **Review core concepts:** If you are new to Kentik Protect, review the **Mitigation Overview** to understand how platforms, methods, and automated triggers work
* **Allow Kentik IPs:** Ensure your edge routers are configured to accept Flowspec BGP announcements from Kentik's infrastructure. You must allow the following IP ranges:

  + **US**: `209.50.158.0/23` (IPv4) and `2620:129:1::/48` (IPv6)
  + **EU**: `193.177.128.0/22` (IPv4)

> **Notes:**
>
> * The settings for Flowspec platforms and methods are made on the portal's **Manage Mitigations** page (**Settings »** **Manage Mitigations**).

## Configure Devices for Flowspec

You should be able to enable Flowspec in the configuration of any routing system that supports **RFC 8955**. The specifics of how to achieve this vary depending on the brand, operating system, model, and version of the device. To find configuration information for your device:

1. Navigate to Kentik's **config-snippets repository** on GitHub.
2. In the list of router brands, find and click on the directory corresponding to the manufacturer of your device. If the manufacturer makes multiple models that are configured differently, continue clicking until you reach the directory corresponding to the device that you want to configure.
3. Inside that directory, look for the file named flowspec.conf.
4. Open the file and configure your device as shown in the code snippet.

The following example is a config snippet for Flowspec on Juniper MX routers:

```
# this configuration on JunOS assumes you already have a BGP session configured
Protocols {
  bgp {
    group route-consumers_v4 {
      # Kentik-provided peering IP.
      neighbor {{kentik_UI_bgp_peering_ipv4}} {
        family inet {
          flow;
        }
      }
    }
  }
}
# Use the RFC 8955 defined ordering of the terms instead of the earlier draft ersion.
routing-options {
  flow {
    term-order standard;
  }
}
```

### Enable Flowspec in Device Settings

Each device that you configure to enable Flowspec (see part A above) must also be enabled for Flowspec in the device settings in the Kentik portal:

1. In the **Device List** on the Devices page (**Settings »** **Networking Devices**), click the Edit icon of the device’s row to open the **Edit Device** dialog.

   ![BGP configuration options for Kentik, including type and compatibility settings.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(346).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A32Z&se=2026-05-12T09%3A37%3A32Z&sr=c&sp=r&sig=DaliSXvfL1DEjZz5%2FfvqWorIvEHPIN9unnAnsPU2T2w%3D)
2. On the **BGP** tab, choose *Peer with device* from the BGP Type drop-down menu.
3. Set the **BGP Flowspec Compatible** switch to on.
4. Click the **Save** button to save changes and exit the dialog.

### Create a Flowspec Mitigation Platform

Once you've enabled devices for Flowspec (both on the device and in the BGP settings for the device in the Kentik portal) you can create the platform which, when combined with a method, make a mitigation:

1. Go to the **Manage Mitigations** page (**Settings »** **Mitigations**).
2. Click the **Add Mitigation Platform** button.
3. On the Add Mitigation Platform page, select **Flowspec** for your mitigation type.
4. Fill in the common settings for your new platform (see **Mitigation Platform Settings**).
5. Click **Select Devices** to open the **Data Sources** dialog, then choose the routers that you enabled for Flowspec in parts A and B of this workflow.
6. Click **Create Mitigation Platform** to save your new platform.

### Create a Flowspec Mitigation Method

To create a mitigation method for Flowspec:

1. Go to the Manage Mitigations page (**Settings »** **Mitigations**).
2. In your new Flowspec platform pane (created in step C), click the **Add Mitigation Method** button.
3. On the **General** tab of the resulting dialog (see **Mitigation Method Dialogs**), give your method an informative name and description.
4. From the drop-down **Notification Channels** menu, assign a notification channel so that you can be notified when your alert triggers a mitigation.

   > **Note:** If you have not already configured a notification channel, go to the Notifications page to add yourself to an existing channel or create a new channel for this mitigation method. See **Notifications**.
5. Turn on the **Acknowledgement** Required switch, which means that after a mitigation from this method has been triggered, the corresponding alarm must be cleared manually (rather than automatically) from the **Alerts List** on the Alerting page. Once you have acknowledged the alert, the mitigation can then proceed.
6. Use the **IP/CIDRs Excluded** field to enter any IP address(es) that you’d like to exclude from the actions taken by this method (see **Flowspec Traffic Actions**). Good candidates might be infrastructure addresses, point-to-point networks, or other address critical to the normal functioning of your network.
7. Now select the **Grace Period** that Kentik should honor prior to withdrawing the Flowspec action. Many operators are happy with the 30-minute default because it provides enough cushion to discourage repeat attacks while not being excessively punitive to the IP that was the attack destination.

### Specify Flowspec Conditions and Actions

To configure the new mitigation method for Flowspec:

1. On the **Details** tab of the **Add Mitigation Method** dialog, in the **Traffic Matching** pane, set the conditions that will be used to match traffic.

   > **Note:** The components that are used for conditions are described in **Flowspec Component Types**, and the controls for setting conditions are explained in **Flowspec Condition Controls**. To avoid unintended consequences, familiarize yourself with the notes in those topics.
2. In the **Traffic Filtering Actions** pane, specify the actions that Flowspec receivers are to take on traffic that matches the traffic matching conditions (see **Flowspec Traffic Actions**).

   > **Note:** To save the method you must select one of the five actions available on the Traffic Filtering Actions drop-down. Setting only the Sample and/or Terminal switch is not sufficient to define a Flowspec method.
3. Click the **Add Mitigation Method** button to save the new method and exit the dialog. The new method will appear in the Flowspec mitigation platform pane.

### Link Existing Methods to a Platform

If your system has existing Flowspec mitigation methods, you can also add them to your platform:

1. Click **Edit** in the mitigation platform pane (top right).
2. Under Assign mitigation methods, click to open a drop-down list of Flowspec-based mitigation methods. Select any methods you would like to add to the platform. Already linked methods appear in blue.

   > **Note:** By linking the method to the platform, you create a mitigation that can be automatically or manually applied.
3. Click the **Update Mitigation Platform** button to save the platform and return to the Manage Mitigations page.

### Assign an Automated Flowspec Mitigation to an Alert Policy

Once configured, a mitigation can be deployed for either automated mitigation (initiated when an alarm is triggered by a threshold in an alert policy) or manual mitigation. To deploy your new Flowspec mitigation automatically, assign the mitigation to an alert threshold (see **Threshold Mitigations**):

1. Go to the Policies page (choose **Alerting** from the navbar menu, then click the **Configure Alert Policies** button).
2. In the **Alert Policies List**, find the policy to which you’d like to assign a Flowspec mitigation.
3. Click to display the Action menu and select **Configure Policy** from the drop-down menu.
4. On the **Edit Alert Policy** page, go to the **Thresholds** tab. Click one of the tabs to select the threshold (Critical, Major, etc.) to which you’d like to assign mitigation.
5. Under **Actions,** click to display the drop-down below **Mitigations** to show a list of mitigations (each displayed as "platform - method"). Choose the mitigation that you created in tasks C to F above, then click the **Add** **Mitigation** button.

   > **Note:** The Mitigations section is only visible if the policy’s dimension list (Dimensions setting in **Data Funneling**) includes source or destination IP/CIDR.
6. In the individual mitigation pane that appears under **Mitigations,** click to open the **Apply Mitigation** drop-down and choose when you’d like to have the mitigation take effect, which could be:

   1. Immediately when Kentik raises the alarm.
   2. After a user manually acknowledges the alarm,
   3. After a user manually acknowledges the alarm or after a timer expires where no user has acknowledged the alarm.
7. Click the **Clear Mitigation** drop-down to choose when the alarm about the mitigation should be cleared from the **Mitigations List**.
8. Click the **Save** button (right top).

   > **Note:** The creation and settings of policies to which you can assign mitigations is covered in **Alert Policies**.

> **Notes:**
>
> * Flowspec-based mitigations should only be assigned to a threshold in an alert policy whose key definition includes the dimension IP/CIDR (see Dimensions in **Data Funneling**).
> * If the **Infer From Alarm** switch is on for any conditions in a Flowspec-based method (see **Flowspec Condition Controls**) then a mitigation using that method should only be assigned to a threshold in an alert policy whose key definition includes the corresponding dimensions.
>
>   **Example***:* If a given method has Infer from Alarm on for Destination IP/CIDR, Protocols, and Source Ports, then a mitigation using that method will be available to assign to a threshold only in policies whose key definition includes the dimensions Destination IP/CIDR, Protocol, and Source Port Number.

### Deploy a Manual Flowspec Mitigation

Flowspec mitigations can be deployed manually as described in **Start a Manual Mitigation**. Before applying a Flowspec-based manual mitigation, familiarize yourself with the special considerations discussed in **Manual Mitigation with Flowspec**.

## Next Steps

As with any mitigation, once your Flowspec actions are applied to actual traffic (manually or automatically), it is important to monitor their status and routing impact:

* View the overall status of the mitigation on the **Mitigations** page.
* Track the specific routing actions and matches on the **BGP Announcements** page.
