# BGP Announcements

Source: https://kb.kentik.com/docs/bgp-announcements

---

This article covers the **BGP Announcements** page in the Kentik portal.

![Overview of BGP announcements with filters for type, session status, and devices.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/BGPA-main.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A37Z&se=2026-05-12T09%3A31%3A37Z&sr=c&sp=r&sig=h5txom%2BpYlHTw1dW5to1srZ4AWE79%2FQcKsziPxzohfE%3D)

*Overview of BGP announcements with filters for type, session status, and devices.*

The **BGP Announcements** page is a dedicated dashboard that displays all BGP announcements tied to your Kentik Protect **mitigations**. It provides a centralized view for monitoring active mitigation sessions, their corresponding rules, and the actions being applied to your network traffic.

> **Note**: For broader information on configuring and managing mitigations, start with the Mitigations Overview.

## BGP Announcements Access

You can reach the **BGP Announcements** dashboard from within the Kentik portal:

1. From the portal’s main nav menu, in the **Protect** section, select **Mitigations**.
2. Click the **BGP Announcements** link located in the top right corner of the page.

![Manage mitigations and BGP announcements with available action options displayed.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/BGPA-announcements-button.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A37Z&se=2026-05-12T09%3A31%3A37Z&sr=c&sp=r&sig=h5txom%2BpYlHTw1dW5to1srZ4AWE79%2FQcKsziPxzohfE%3D)

## Filter the Announcements List

At the top of the page, you can use the **Search** bar or the **Filters** menu to narrow down the list of announcements. You can filter by:

* **Type:** Adaptive Flowspec, Flowspec, or RTBH.
* **Session Status:** Up or Down.
* **Devices:** Filter by specific peering devices.
* **Mitigations:** Filter by the specific mitigation triggering the announcement.

> **Note**: You can also use the **Group By** dropdown to organize the list view.

## Announcements List UI

The main view consists of a table detailing your BGP announcements. Each row represents a specific announcement and includes the following columns:

* **Type:** The classification of the mitigation (e.g., RTBH, Flowspec, Adaptive Flowspec).
* **Session Status:** The current operational state of the BGP session (e.g., Up, Down).
* **Device:** The specific network device with which the BGP session is peered. Click to open its **Device Details** page.
* **Mitigations:** The ID or name of the mitigation triggering the announcement. Click to open it in the **Mitigations**page.
* **Match:** The specific conditions or routing prefixes (e.g., Destination IP/CIDR, Source Prefix, BGP Communities) that triggered the mitigation.
* **Action:** The resulting routing action applied to the matching traffic (e.g., Discard).
* **Unapplied Mitigations:** A count or view of mitigations configured for this session but not currently applied.
* **Start Time (UTC):** The exact date and time the announcement was initiated.

## Announcement Details and Actions

![Details of a network announcement including source prefix, actions, and mitigations.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/BGPA-announcements-detail-drawer.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A37Z&se=2026-05-12T09%3A31%3A37Z&sr=c&sp=r&sig=h5txom%2BpYlHTw1dW5to1srZ4AWE79%2FQcKsziPxzohfE%3D)Click on any row in the Announcements list to open a side panel with further details about that specific session.

From this side panel, you can review the full breakdown of a particular announcement, including the details covered in  **Announcements List UI**, plus the following:

* **Withdraw announcement** (button)**:** Click to immediately halt the BGP announcement (i.e. if you need to manually intervene).
