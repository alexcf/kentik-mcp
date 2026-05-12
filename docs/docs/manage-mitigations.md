# Manage Mitigations

Source: https://kb.kentik.com/docs/manage-mitigations

---

This article covers the general UI and page layout of the **Manage Mitigations** settings page in the Kentik portal.

![The Manage Mitigations page enables settings for mitigation platforms and methods.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/MM-manage-mitigations-main.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A34Z&se=2026-05-12T09%3A34%3A34Z&sr=c&sp=r&sig=9vj7oRnvwvxkltN6CmKt1Rfq92F7a6pNO8pSpTWYfwg%3D)

*The Manage Mitigations page enables settings for mitigation platforms and methods.*

> **Note:** For more on Kentik’s mitigations features, see **Mitigation Overview**, **Mitigations Page**, and **Manual Mitigation**.

## Getting Started

The **Manage Mitigations** settings page, accessed via **Settings** **»** **Manage Mitigations** in the portal's main nav menu, is your central hub for configuring how your network responds to DDoS attacks and other traffic anomalies.

Because configuring mitigations involves several distinct components, the documentation is divided into the following guides:

* **Manage Mitigations** (this article): Covers the general UI, page layout, and how to read platform status and method lists.
* **Mitigation** **Platforms**: Covers how to add, edit, or remove the infrastructure where mitigations run, including universal API and BGP settings.
* **Mitigation Methods**: Covers how to configure the universal rules, notification channels, and automation grace periods that trigger mitigations.
* **Method Configuration Deep-Dives**: For detailed, protocol-specific settings, see the individual guides for **RTBH**, **Flowspec**, **Adaptive Flowspec**, and **Third-Party** methods.

### Create Initial Platform

The first time you access the Manage Mitigations page, you'll see a splash screen inviting you to create your first mitigation.

Click **Create Your First Mitigation** to continue to the Add Mitigation Platform page. Your first platform need not be perfect right away; you can go back and edit it before you use it. For information about creating the platform, see **Platform Settings Pages** and **Add a Mitigation Platform**.

![Kentik Mitigations prevent disruptive traffic, ensuring network availability and security.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/MM-create-your-first-mitigation.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A34Z&se=2026-05-12T09%3A34%3A34Z&sr=c&sp=r&sig=9vj7oRnvwvxkltN6CmKt1Rfq92F7a6pNO8pSpTWYfwg%3D)

*The initial splash screen appears until a platform has been created in your organization.*

## Manage Mitigations Page UI

Once a mitigation platform exists in your system, you’ll see the **Manage Mitigations** page, which includes the following UI elements:

* ![Buttons for starting manual mitigation and adding a mitigation platform are displayed.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/MM-main-buttons.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A34Z&se=2026-05-12T09%3A34%3A34Z&sr=c&sp=r&sig=9vj7oRnvwvxkltN6CmKt1Rfq92F7a6pNO8pSpTWYfwg%3D)**Start a Manual Mitigation** (button): Opens the **Start Manual Mitigation Dialog** (see **Manual Mitigation**).
* **Add Mitigation Platform** (button): Opens the Add Mitigation Platform page.
* **Mitigation Platforms**: A list of platform panes where you can manage it and list associated methods (see **Platform Panes**).
* **Unassigned Methods**: A list of methods that are not currently assigned to a platform (see **Unassigned Methods List**).

### Platform Panes

Each mitigation platform pane contains information about both the platform and the mitigation methods that have been created and added to the platform. Each pane contains the following UI elements:

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(671).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A34Z&se=2026-05-12T09%3A34%3A34Z&sr=c&sp=r&sig=9vj7oRnvwvxkltN6CmKt1Rfq92F7a6pNO8pSpTWYfwg%3D)

* **Platform Header**: Common info and settings for this platform.

  + **Type**: The type of mitigation platform (e.g., RTBH, Radware, etc.) appears directly above the platform’s name and the logo for that type appears just to the left.
  + **Name**: The user-assigned name for the mitigation platform.
  + **Description**: Optional user-defined description for the mitigation platform.
  + **ID**: A system-assigned unique identification number for the mitigation platform.
  + **Platform Status**: A button that opens the **Platform Status Popup**:

    - **Green checkmark icon**: The platform is available for mitigation.
    - **Orange question mark icon**: There’s an error that is detailed in the popup (e.g., platform not currently active, authentication error).
  + **Edit Platform** (pencil icon): Opens the Edit Mitigation Platform page (see **Platform Settings Pages**).
  + **Remove** (button): Opens a popup where you click **Remove** to confirm the deletion of the mitigation platform from your organization.
* **Methods List**: A list of the mitigation methods assigned to this platform (see **Mitigation Methods List**).
* **Add New Method** (button): Opens the Add Mitigation Method dialog (see **Mitigation Method Dialogs**), enabling you to add a method to this platform.
* **Assign Existing Method**: A dropdown to choose to assign a method to this platform that was previously created for this platform.

### Platform Status Popup

The status popup for each platform contains information that can be used for troubleshooting by Kentik Product Support (see **Customer Care**). The popup contains the following elements:

* ![Platform status showing device connection and active announcements.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/MM-platform-status-popup-connected.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A34Z&se=2026-05-12T09%3A34%3A34Z&sr=c&sp=r&sig=9vj7oRnvwvxkltN6CmKt1Rfq92F7a6pNO8pSpTWYfwg%3D)**Lozenges**: Shows status for All, Auth, Options, Version.

  + Green means healthy, pink indicates an error.
  + “All” lozenge turns green when all other lozenges are green.
* **Device Count**: The total number of devices connected to the platform.
* **Copy Status JSON** (button): Copies the JSON to the clipboard, with identification and status information about the platform. You can paste the JSON into your correspondence with Product Support.
* **Device Table**: Lists each Device, along with a Connected status icon, and an Announcements summary with links to any active announcements.

  > **Note**: If no devices are found, a message that the platform is not currently active is shown instead.

### Mitigation Methods List

Each row of the **Mitigation Methods** list represents a mitigation method created for this platform by a user in your organization, and includes the following:

* **Name**: A user-assigned name for the mitigation method. ![List of A10 methods with IDs and options for editing or deleting entries.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/MM-methods-list.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A34Z&se=2026-05-12T09%3A34%3A34Z&sr=c&sp=r&sig=9vj7oRnvwvxkltN6CmKt1Rfq92F7a6pNO8pSpTWYfwg%3D)
* **Description**: An optional user-assigned description for the mitigation method.
* **ID**: A system-assigned identification number for the mitigation method.
* **Clone**: Click the copy icon to create a duplicate of the mitigation method.
* **Edit**: Click the pencil icon to open the Edit Mitigation Method dialog (see **Mitigation Method Dialogs**).
* **Unassign Method**: Click the **X** icon  to immediately remove the mitigation method from the platform to which it was assigned.

  > **TIP**: If a method is not assigned to any platforms, it appears on the **Unassigned Methods List**.

### Unassigned Methods List

Each row of the **Unassigned Methods** list represents a previously created mitigation method that is not currently assigned to a platform, and includes the following:

* **Type:** The platform type (e.g., RTBH, Flowspec, etc.) for which this method can be used.
* **Name**: A user-assigned name for the mitigation method.
* **Description**: A user-assigned description for the mitigation method.
* **ID**: A system-assigned identification number for the mitigation method.
* **Clone**: A button that creates a duplicate of the mitigation method, which is added to the bottom of the list.
* **Edit Method** (pencil icon): A button that opens the Edit Mitigation Method dialog (see **Mitigation Method Settings**).
* **Remove**: A button that opens a popup in which you can click **Remove** to remove the method from your organization's collection of methods (see **Remove a Mitigation Method**).

## Summary and Next Steps

You should now have a solid understanding of the Manage Mitigations interface, how to check the health of your integrations using the Platform Status popups, and the fundamental difference between your platforms and methods.

**Next Steps**:

1. **Build your infrastructure**: If you haven't already, your first step is to configure the routers or third-party services that will execute your mitigations. See **Mitigation Platforms** to get started.
2. **Define your rules**: Once your platform is active and healthy, you can begin defining the specific traffic thresholds and routing behaviors by creating **Mitigation Methods**.
