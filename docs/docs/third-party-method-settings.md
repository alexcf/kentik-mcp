# Third-Party Method Settings

Source: https://kb.kentik.com/docs/third-party-method-settings

---

This article covers how to define the mitigation method settings for third-party platforms like A10, Cloudflare, or Radware from Kentik’s **Manage Mitigations**page.

> **Note:** For information about third-party platforms, see **Mitigation Platforms** and **Third-Party Mitigation**.

## A10 Method Settings

For methods in thw **A10 Thunder TPS**platform, the controls on the **Details** tab of the settings dialog include Priority, A10 Mode, and Announce Via BGP Network Statement.

> **Note:** To ensure correct operation of your A10 mitigation, the above settings should be made in consultation with A10 customer support.

## Cloudflare Method Settings

For methods in the **Cloudflare Magic Transit** platform, the settings dialog has no **Details** tab (there are no Cloudflare-specific method settings).

> **Notes:**
>
> * Cloudflare applies Magic Transit mitigation only when traffic exceeds minimums (100K pps for TCP/UDP; 60K pps for ICMP/GRE). Assigning Cloudflare MT to a Kentik alert below these thresholds may show mitigation as active when it isn't. For lower volumes, use an alternative mitigation platform (RTBH, Flowspec, etc.).
> * Cloudflare throttles withdrawal operations to prevent “flapping.” Kentik recommends a minimum 30 minute grace period for Cloudflare methods.

## Radware Method Settings

For methods in the **Radware DefensePro** (v2.7 or greater) platform, the **Details** tab of the settings dialog includes settings such as Protocol, Use Protocol From Alert Dimension When Available, Ensure at least /24, and Protected Object Name.

> **Note:** To ensure correct operation of your Radware mitigation, the above settings should be made in consultation with Radware and with Kentik Product Support (see **Customer Care**).

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(678).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A33Z&se=2026-05-12T09%3A33%3A33Z&sr=c&sp=r&sig=fauuJ2C5JAd9H1w17uAGioUmXTkc1gN4eCuoMqKBjGY%3D)The **Details** tab also includes a set of **Mitigation Baselines** drop-downs (ICMP Bytes/second, etc.), which each list all of the policies in your organization.

You use these selectors to choose the policy in which your new method will be assigned to a policy threshold (see **About Alert Thresholds**).

This ensures that the metric and protocol of the baseline information passed to Radware when the threshold is triggered are the same as the primary metric and protocol used for the threshold's baseline evaluation.

### Policy Metric and Protocol

Unless you are creating a mitigation method for an "all traffic" policy, you'll choose a policy for only one of the metrics drop-downs. To choose the drop-down to set, you'll need to know the following information, which is set on the **Dataset** tab of the alert policy (see **Data Funneling**) and can also be found in the policy's **Policy Details Drawer**:

* **Primary metric**: The metric designated as primary in the tab's **Metrics** setting.
* **Protocol**: The protocol (if any) set in the tab's **Filter** setting.

The table below shows the situations in which the **Protocol** setting on the mitigation method's **Details** tab will be used.

| Use Protocol From Alert Dimension switch | Protocol dimension set in the policy Dataset | Protocol setting in Method details |
| --- | --- | --- |
| Off | N.A. | Used |
| On | No | Used |
| On | Yes | Not used |

### Create a Radware Mitigation Method

To create a Radware-based mitigation method and assign it to a policy threshold:

1. On the Policies page, create the policy to which you'd like to attach a Radware method. For more information on creating a policy, see **Adding a Policy**.
2. On the **Manage Mitigations**page, click the **Add New Method** button under the Radware platform to open the Add Mitigation Method dialog.

   > **Note:**Each policy that will use baseline evaluation to trigger Radware mitigation should have only one method.
3. On the dialog's **General** tab, specify the settings in **General Method Settings**.
4. On the dialog's **Details** tab, specify the settings discussed above (Protocol, etc.).
5. In the **Mitigation Baselines** pane, find the dropdown for the protocol (TCP, UDP, or ICMP) and metric (Bytes/second or Packets/second) set on the **Dataset** tab of the alert policy (see **Data Funneling**).

   > **Note:** If you didn't set protocol as a dimension in the policy or the method will use the protocol from the **Details** tab's **Protocol** selector (see table above) and you've set the selector to Other, please consult with Kentik Product Support (see **Customer Care**).
6. Once the settings of the Add Mitigation Method dialog are fully specified, click the **Add Mitigation Method** button to create the method and close the dialog.
7. Go to the **Alert Policies Page** (choose **Alerting** from the portal's main navbar menu, then click the **Manage Policies** button at the upper right).
8. In the **Policies** list, find the policy that you created in step 1, and choose **Edit Policy** from the Action menu at the right of the row.
9. On the **Thresholds** tab of the policy's **Edit Policy** page (see **Policy Threshold Settings**), use the list at the left to choose the threshold (Critical, Severe, etc.) to which you'd like to attach the mitigation.
10. Scroll down to the **Actions** pane and use the **Mitigations** drop-down to choose the mitigation created above. Once a mitigation is selected you can click on the **Add Mitigation** button to attach the mitigation to this threshold of the policy.

> **Notes:**
>
> * To view details about the mitigation method, go to the **Manage Mitigations** page and click the status icon in the controls at the right of the Radware platform pane heading. This will open a Platform Status popup with JSON-formatted platform information.
> * For general method settings like notification channels and automation grace periods, return to the **Mitigation Methods** article.
