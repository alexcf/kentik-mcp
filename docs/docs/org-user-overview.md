# Org & User Settings

Source: https://kb.kentik.com/docs/org-user-overview

---

This article discusses organization-wide and user-specific settings for the Kentik portal.

> **Notes:**
>
> * New/trial Kentik customers should refer to the initial Setup Tasks described in **Getting Started**.
> * For information the portal's Settings page, see **Main Settings**.

## Organization Settings

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/org settings(1).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A47Z&se=2026-05-12T09%3A27%3A47Z&sr=c&sp=r&sig=1mTNj3GTIP2ZDg8hCK%2BGjKWDGSHU%2BAz5ngs0zL4b2Bw%3D)The **Organization Settings** menu shows your organization’s name and Kentik Company ID (CID) at top, followed by links to a variety of pages where you can — depending on your assigned roles and permissions — view or edit company-wide settings.

The following settings pages may be accessed via this menu:

* **Licenses**: Information about your Kentik licenses (referred to as "plans" in earlier versions of Kentik), each representing a level of service specified in an agreement between you and Kentik.

  + Values are displayed for attributes such as maximum flows per second (FPS), BGP data, and data retention period.
  + Licenses are configured by Kentik and are not editable by customers.
* **Users**: Add, remove, or manage settings for your organization’s users (see **Users**).
* **RBAC Management**: View, add, remove, or edit settings for your organization’s RBAC roles (see **Kentik RBAC**).

  > **Note**: Administrators and Super Administrators (see **Kentik-managed Roles**) determine which other users in their organization can access the RBAC Management page.
* **Authentication & SSO** (Super Administrator users only): Access your organization’s password expiration and single Sign-on (SSO) settings. See **Authentication & SSO**.
* **Audit Log** (Administrator users only): Access your organization’s **Audit Log** to view the list of changes made in your Kentik environment and who made them.
* **Access Control**: Set your organization’s **Access Control** to prevent unauthorized access to Kentik.
* **Credentials Vault**: Securely store your multi-field credentials in a central location to easily use them throughout the portal. See **Credentials Vault**.
* **Kentik AI**: Set all aspects of **Kentik AI**, including AI Advisor and Query Assistant, to either disabled (default) or enabled for your entire organization.

  > **Note**: This setting is present only for users with Super Administrator privileges.
* **Public Shares**: Your organization’s Kentik-generated views shared with people (internal and external to the organization) who aren’t Kentik users (see **Public Share Link**).

  > **Note**: In the menu, the number displayed to the right of Public Shares is the number of public shares currently active for your organization.
* **Report Subscriptions**: View and manage your organization’s **Subscriptions**.

  > **Note:** In the menu, the number displayed to the right of Report Subscriptions is the number of subscriptions your organization is currently using.
* **Custom DNS**: Takes you to the **Custom DNS Page**, where you can run a **Verify Reverse DNS Lookup**(see **About Custom DNS**).

## User Menu

The **User** menu shows your full name, email (as defined in your **User Profile**), and Kentik user ID (UID) at top, followed by links to user-specific portal settings.   
  
The following user-specific settings and controls are accessible from this menu:

* ![User profile menu displaying options like authentication, theme, and log out.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/OUO-user-menu.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A47Z&se=2026-05-12T09%3A27%3A47Z&sr=c&sp=r&sig=1mTNj3GTIP2ZDg8hCK%2BGjKWDGSHU%2BAz5ngs0zL4b2Bw%3D)**User Profile**: Your Kentik **User Profile** page is where you can set your portal preferences for layout, display, notifications, landing page, time zone, and theme. You can also reset your password and manage two-factor authentication and API tokens.
* **Authentication**: Takes you directly to the **Authentication Settings** of your **User Profile** where you can reset your password or manage your API token and two-factor authentication settings.
* **Recent** **Exports**: Takes you to the **Recent Exports Page**, where you'll find a table listing the PDFs, CSVs, or PNGs that you've recently exported from the Kentik portal using **Actions** » **Export.**
* **Theme**: Choose Auto (OS), Light, or Dark as your **User-specific Theme**.
* **Chart Colors**: Takes you directly to the **Chart Colors** tab of your User Profile, where you can set color schemes for your portal visualizations.
* **UI Layout Preferences**: Opens the **Layout Preferences** page where you can specify how you’d like drawers and docks to display in the portal.
* **Log Out**: Log out of your session on the Kentik portal.
