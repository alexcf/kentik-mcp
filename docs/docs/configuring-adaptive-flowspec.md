# Configuring Adaptive Flowspec

Source: https://kb.kentik.com/docs/configuring-adaptive-flowspec

---

This guide walks you through the complete process of creating, configuring, and deploying an Adaptive Flowspec mitigation within the Kentik portal—from preparing your edge routers to triggering automated filtering actions.

Adaptive Flowspec is a dynamic mitigation technique that automatically adjusts its generated BGP Flowspec rules based on changing attack vectors and real-time network conditions.

## Before You Begin

* **Review core concepts**: If you are new to Kentik Protect, review the **Mitigation Overview** to understand how platforms, methods, and automated triggers work.
* **Allow Kentik IPs**: Ensure your edge routers are configured to accept Flowspec BGP announcements from Kentik's infrastructure. You must allow the following IP ranges:

  + US: 209.50.158.0/23 (IPv4) and 2620:129:1::/48 (IPv6)
  + EU: 193.177.128.0/22 (IPv4)

## Configure Devices for Adaptive Flowspec

Because Adaptive Flowspec propagates rules using standard BGP Flowspec, you must enable Flowspec in the configuration of any routing system that supports RFC 8955.

* To find configuration snippets specific to your device's brand and operating system, navigate to Kentik's **config-snippets repository** on GitHub and locate the `flowspec.conf` file for your hardware.

## Enable Flowspec in Kentik Device Settings

Each router you configure must also be enabled for Flowspec communication within the Kentik portal:

1. Navigate to the Device List (**Settings »** **Networking Devices**) and click the Edit icon for your router.
2. On the **BGP** tab, choose Peer with device from the BGP Type drop-down menu.
3. Set the **BGP Flowspec Compatible** switch to On.
4. Click **Save**.

## Create an Adaptive Flowspec Mitigation Platform

Once your devices are configured, you must create the platform that will execute the mitigations:

1. Go to the Manage Mitigations page (**Settings »** **Manage Mitigations**).
2. Click Add Mitigation Platform and select Adaptive Flowspec.
3. Fill in the common settings and select the devices you prepared in the previous steps.

> **Note**: For more details on platform UI settings, see **Mitigation Platforms**.

## Create an Adaptive Flowspec Mitigation Method

With your platform built, you can define the specific traffic analysis algorithms and matching rules:

1. In your new Adaptive Flowspec platform pane, click **Add Mitigation Method**.
2. On the **General** tab, provide a name, description, and assign your notification channels.
3. Use the **Traffic Analysis**, **Flowspec**, and **Preview** tabs to define how the mitigation will dynamically match and filter traffic.

> **Note**: For a complete explanation of the algorithms and UI fields on these tabs, see **Adaptive Flowspec Method Settings**.

4. Click **Add Mitigation Method** to save your rules.

## Assign the Mitigation to an Alert Policy

To deploy your new Adaptive Flowspec mitigation automatically when an attack occurs, you must assign it to an alert threshold:

1. Go to the Policies page (**Alerting »** **Configure Alert Policies**).
2. Edit the policy you want to trigger the mitigation, and navigate to the **Thresholds** tab.
3. Under the Actions section of your desired threshold (e.g., Critical), select your new Adaptive Flowspec mitigation from the drop-down list and click **Add Mitigation**.
4. Specify when the mitigation should take effect (e.g., Immediately) and click **Save**.

## Next Steps

Once your Adaptive Flowspec actions are applied to actual traffic, you can monitor their status and routing impact:

* View the overall status of the mitigation on the **Mitigations Page**.
* Track the specific dynamically generated routing actions on the **BGP Announcements** page.
