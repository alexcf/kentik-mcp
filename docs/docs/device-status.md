# Device Status

Source: https://kb.kentik.com/docs/device-status

---

This article discusses the Device Status page of the Kentik portal.

> **Note:** For ongoing management of devices, or to add additional devices, access the **Network Devices** page from the **Settings** link on the main menu.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(560).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A04Z&se=2026-05-12T09%3A27%3A04Z&sr=c&sp=r&sig=pnAjyn5r8Jk6vuO6ar5x007WNs4%2FP6xVySvmr9V26ew%3D)

*The Device Status page enables you to review and fix device settings.*

## About Device Status

The Device Status page is used to identify and correct issues with the settings or configuration of a device. The existence of these issues is indicated in a callout at the top of **Network Explorer**.

![A message about misconfigured devices from the Device Status page.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/NE-Review_devices_notice-85h672w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A04Z&se=2026-05-12T09%3A27%3A04Z&sr=c&sp=r&sig=pnAjyn5r8Jk6vuO6ar5x007WNs4%2FP6xVySvmr9V26ew%3D)

A message about misconfigured devices from the Device Status page.

## Device Status Page

The Device Status page of the Kentik portal helps you complete and correct issues related to device configuration. The page includes the following main sections:

* **Status list**: A list of cards that each show the number of devices in a given status. Click a card to view the list of devices with the corresponding status (see **Device Statuses**).
* **Device list**: A list of all devices in the currently selected status.

> **Note:** To access the page, see **Using Device Status**.

#### Device Statuses

A device can have one of four possible statuses:

* **Incomplete**: Devices for which the Flow and SNMP configurations made during the onboarding process have errors.

  > **Note**: This status applies only to devices added during your organization's onboarding process.
* **Never received flow**: Devices from which Kentik hasn’t received any flow data.
* **Stopped sending flow**: Devices from which Kentik previously received flow data but no longer does.
* **No interfaces**: Devices from which no interfaces have been registered with Kentik.

## Using Device Status

To view and correct issues with the setup or status of devices:

1. In the callout at the top of Network Explorer, click the **Review your devices** button, which takes you to the **Device Status Page**.
2. In the **Status List** at left, click a card with the status you’d like to view (see **Device Statuses**).
3. In the display area at right you'll see a list of cards representing the devices in the status you chose. The card will contain a button, e.g., **Complete Device Configuration** or **Verify Device Configuration**. Click the button to open the **Device Settings Dialog**and correct any issues for the device.

---

© 2014-25 Kentik
