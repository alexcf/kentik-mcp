# Mitigation Platforms

Source: https://kb.kentik.com/docs/mitigation-platforms

---

This article covers the settings required to manage migration platforms on Kentik’s **Manage Mitigations**page.

## Platform Settings Pages

A **Mitigation Platform** represents the physical infrastructure or third-party cloud service that Kentik interacts with to execute an active mitigation. Before you can define specific routing rules or thresholds (see **Mitigation Methods**), you must first configure the platform where those rules will be deployed.

The platform's settings are specified in one of two platform settings pages:

* **Add Mitigation Platform**: Use this page to register a new platform with Kentik. To access, click **Add Mitigation Platform** on the **Manage Mitigations** page.
* **Edit Mitigation Platform**: Use this page to edit an existing platform. To access, click **Edit Platform** (pencil icon) at the top of any existing platform’s pane on the **Manage Mitigations** page.

![Interface for adding a mitigation platform with various configuration options displayed.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/MM-add-mitigation-platform-dialog-radware.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A15Z&se=2026-05-12T09%3A36%3A15Z&sr=c&sp=r&sig=U2AM1ytFdLYjEgzhEiVKqC3fTRVUc3xbVOo1q3cD0hc%3D)

*The Add Mitigation Platform dialog showing Radware selected as a third-party integration.*

> **Note:** In addition to configuring a mitigation platform and method in Kentik, you must also allow the following IP ranges on third-party mitigation platforms (e.g., Radware, A10) and on any devices used for Flowspec, RTBH, or Adaptive Flowspec mitigations:
>
> * **US**:  `209.50.158.0/23 (IPv4)` and `2620:129:1::/48 (IPv6)`
> * **EU**: `193.177.128.0/22 (IPv4)`

### Common Platform Settings

The following controls on the platform settings pages are common to all types of mitigation platforms:

* **Cancel** (button): Click to exit without saving changes.
* **Create Mitigation Platform** (button, Add Mitigation Platform only): Click to save the new platform’s settings and exit.
* **Update Mitigation Platform** (Edit Mitigation Platform only): A button that saves the changes you've made to the platform settings and returns you to the Manage Mitigations page.
* **Select Your Mitigation Type** (Add Mitigation Platform page only): A list of the mitigation platform types supported by Kentik. Click on a type to select it.

  > **Notes:**
  >
  > + This list includes supported types which may not be accessible to your organization (e.g., if you lack a Radware, Cloudflare, or A10 system). Kentik does not verify your mitigation type choice.
  > + Cloudflare applies Magic Transit mitigation only when traffic exceeds minimums (100K pps for TCP/UDP; 60K pps for ICMP/GRE). Assigning Cloudflare MT to a Kentik alert below these thresholds may show mitigation as active when it isn't. For lower volumes, use an alternative mitigation platform (RTBH, Flowspec, etc.).
* **Assign Mitigation Methods**: Controls for assigning mitigation methods to this platform:

  + **Assigned methods**: A list of the methods assigned to the mitigation platform. A placeholder appears until methods are assigned. Once methods are added they will be listed by name, and each will have the **Copy**, **Edit**, and **Unassign Method** buttons described in **Mitigation Methods List**.
  + **Add New Method**: A button that opens the Add Mitigation Method dialog for the selected mitigation type.
  + **Assign Existing Method**: A drop-down from which you can choose to add a method that was previously created for this platform.
* **Configure Devices** (RTBH, Flowspec, or Adaptive Flowspec): Select one or more devices to assign to the mitigation platform, (see **Configure Platform Devices**).
* **Configure API** (Cloudflare, A10 or Radware): API settings varying by mitigation platform (see **Configure Platform APIs**).
* **Finish Up**:

  + **Name**: User-specified name for the mitigation platform.
  + **Description**: Optional user-provided description text.

### Configure Platform Devices

When the mitigation platform type is RTBH, Flowspec, or Adaptive Flowspec, the following Configure Devices controls are used to set the devices to which mitigations run on this platform will apply:

* **Devices list**: A list of the devices assigned to the platform. A placeholder will appear unless at least one device is assigned.
* **Select Devices**: A button that opens a **Data Sources Dialog**.

  > **Note:** The dialog will show only devices with the following settings in the **BGP** tab of theAdd Devices or Edit Devices dialog in **Settings »** **Network Devices** (see **Device BGP Settings**):
  >
  > + For RTBH, Flowspec, and Adaptive Flowspec: The drop-down BGP Type setting is **Peer with Device**.
  > + For Flowspec and Adaptive Flowspec: The **BGP Flowspec Compatible** switch is On. The devices will receive Flowspec rules via MP-BGP.

To choose the devices on which mitigation will be implemented with this platform, click the button and select devices from the dialog. Each selected device will be added to the list as a lozenge that shows the device's name. To remove a device from the list, use the **X** at the right of its lozenge.

> **Note:** For method configuration details and general information regarding these platform types, see the following resources:
>
> * **RTBH**: **RTBH Mitigation**| **RTBH Method Configuration**
> * **Flowspec:** **Flowspec Mitigation**| **Flowspec Method Configuration**
> * **Adaptive Flowspec**: **Adaptive Flowspec Method Configuration**

### Configure Platform APIs

![Kentik's Configure Mitigations settings for the Radware platform.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/MM-edit-platform-radware.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A15Z&se=2026-05-12T09%3A36%3A15Z&sr=c&sp=r&sig=U2AM1ytFdLYjEgzhEiVKqC3fTRVUc3xbVOo1q3cD0hc%3D)When the **Mitigation Type** is set to a third-party mitigation system (e.g., Cloudflare Magic Transit, A10, or Radware), the following Configure API controls are shown with the following API-related settings, depending on platform:

* **IP Address** (A10) or **Vision IP Address** (Radware): The IP address or URL (`https://ip` or `ip` or `https://name` or `name`) of the management interface of the third-party mitigation device.
* **API Login** (Cloudflare or A10) or **Vision API Login** (Radware): User name from the credentials for the third-party mitigation system.
* **API Key** (Cloudflare), **API Token** (A10) or **Vision API Token** (Radware): Key, token, or password for the third-party mitigation system.
* **Cloudflare Account ID** (Cloudflare): The ID of your Magic Transit account with Cloudflare.

  > **Note**: Cloudflare applies Magic Transit mitigation only when traffic exceeds minimums (100K pps for TCP/UDP; 60K pps for ICMP/GRE). Assigning Cloudflare MT to a Kentik alert below these thresholds may show mitigation as active when it isn't. For lower volumes, use an alternative mitigation platform (RTBH, Flowspec, etc.).
* **Delete IP** (Radware or A10): Kentik continually compares its internal list of mitigations with the third-party mitigation system’s list of resources utilized by Kentik-defined mitigations. This switch determines what happens when Kentik finds resources on the third-party system for mitigations that have been deleted from Kentik:

  + **If the switch is On**: Kentik will relay to the third-party mitigation system a list of these resources so that they can be deleted.
  + **If the switch is Off**: Kentik will not notify the third-party system about the resources.
* **Disable TLS Verification** (Radware or A10): An optional switch that when On, disables TLS certificate verification for API connections to the mitigation platform.

  + Allows Kentik to connect to platforms with self-signed certificates or non-standard certificate configurations.

> **Notes:**
>
> * Kentik does not verify the login username or password you provide. Incorrect login information will cause mitigations for the platform to fail.
> * See **Third-Party Method Configuration**for details.
> * For an overview of third-party mitigation, see **Third-Party Mitigation**.

## Manage Mitigation Platforms

Adding and editing platforms via the Manage Mitigations page (**Settings »** **Mitigations**) is covered in the following sections.

### Add a Mitigation Platform

To add a new mitigation platform:

1. On the Manage Mitigations page, click the **Add Mitigation Platform** button to open theAdd Mitigation Platform page.
2. Select your mitigation type and specify the necessary values for that platform (see **Platform Settings Pages**).
3. Save the new platform by clicking the **Create Mitigation Platform** button at upper right.

### Edit a Mitigation Platform

To edit the settings of an existing mitigation platform: ![Editing a Cloudflare mitigation platform with assigned methods and device selection prompt.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/MM-edit-mitigation-platform.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A15Z&se=2026-05-12T09%3A36%3A15Z&sr=c&sp=r&sig=U2AM1ytFdLYjEgzhEiVKqC3fTRVUc3xbVOo1q3cD0hc%3D)

1. On the Manage Mitigations page, within a platform pane, click **Edit** beside the platform you’d like to edit.
2. Edit the platform’s settings (see **Common Platform Settings**).
3. To save changes, click the **Update Mitigation Platform** button at upper right.

### Remove a Mitigation Platform

To remove a platform from your organization's collection of mitigation platforms:

1. On the Manage Mitigations page, click **Remove** in the top right corner of the pane of the platform you'd like to remove.
2. In the resulting confirmation popup, click the **Remove** button.

## Summary and Next Steps

You have now successfully configured the infrastructure and integrations required to execute active mitigations on your network.

**Next Steps**: With your platforms active and your BGP or API settings in place, you are ready to define the specific rules, thresholds, and actions that will trigger when an attack occurs. See the **Mitigation Methods** article to learn how to create and assign these rules to your newly configured platforms.
