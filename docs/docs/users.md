# Users

Source: https://kb.kentik.com/docs/users

---

This article discusses the **Users** page in the Kentik portal.

> **Note:** If you would like assistance with any aspect of registering a user, email us at **support@kentik.com**.     

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(71).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A02Z&se=2026-05-12T09%3A41%3A02Z&sr=c&sp=r&sig=KPLdbmb%2BjZ7IxzXicZRyeRDuNOvTXdcgbo1gJGXPFKI%3D)

*Manage settings for all users in your organization.*

## About Users

A user is anyone at your company that has authorization to access Kentik. Users access various parts of the portal via either level-based permissions or our recently introduced Role-Based Access Control (RBAC; see **What is Kentik RBAC?)**. The areas of the portal where we currently use RBAC are listed at **Initial RBAC Permissions**. All other areas are accessed based on user level (Member, Administrator, or Super Administrator; see **User Level**). When the rollout of RBAC is complete, access to all parts of the portal will be controlled with RBAC permissions and user levels will be deprecated.

## Users Page

The **Users** page is covered in the following topics.

> Notes:
>
> * To access the Users page, choose **Users** from the **Organization Settings** menu on the main navbar.
> * Members can view the Users page but not change any settings.

### Users Page UI

The Users page includes the following main areas and controls:

* **Share** (on subnav): A button that opens the Share dialog so you can share the current view (see **Sharing via the Share Dialog**).
* **Actions** (on subnav): A button that pops up a menu from which you can choose the following actions:

  + **Export**: Export the page’s content as a visual report (PDF) or data table (CSV). A notification appears when the export is ready to download.
  + **Subscribe**: Opens the **Subscription Dialog**, which enables you to subscribe to the Users page, either by choosing an existing subscription (combination of email address and schedule) or specifying a new one.
  + **Unsubscribe**: Opens the Unsubscribe dialog, which enables you to unsubscribe from the Users page. Click the **Subscription** drop-down, select the subscription from which you’d like to unsubscribe, and click **Unsubscribe**.

    > **Note**: **Unsubscribe** will only appear if you have an existing subscription for this page.
* **Add User**: A button that opens the Add User dialog (see **User Settings Dialogs**).
* **Show/Hide Filters** (filter icon): A button that toggles the **Filters** pane between expanded and collapsed.
* **Search**: A field that shows lozenges for the filters currently set in the **Filters** pane and also enables you to enter text. The **User** list will be filtered to show only rows that contain the entered text. Click the **X** at the right of the field to clear entered text, and the X in a lozenge to clear the corresponding filter.
* **Filters Pane**: A set of controls that enable you to filter the **User** list (see **User List Filters**).
* **Action controls**: A set of controls that appear at the top left of the **Users** list and are present only when one or more checkboxes are checked in the **Users** list, enabling you to apply an action to all selected alerts:

  + **Role** (button): Opens the **Role Selection Popup** where you can assign roles to multiple users simultaneously.
  + **Selection indicator**: Indicates how many users are currently selected.
* **User List**: A table listing your organization’s currently registered users (see **User List**).

### User List Filters

The **User** list can be filtered using the controls in the **Filters** pane on the left. The pane includes the following controls:

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(150).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A02Z&se=2026-05-12T09%3A41%3A02Z&sr=c&sp=r&sig=KPLdbmb%2BjZ7IxzXicZRyeRDuNOvTXdcgbo1gJGXPFKI%3D)

* **Clear all** (appears only when you’ve specified one or more filters): A button that removes all currently set filters.
* **Kentik-Managed Role**: Checkboxes that filter the users to those who are assigned a Kentik-managed role:

  + **Super Administrator**: Can perform the most sensitive operations.
  + **Administrator**: Can perform some sensitive operations.
  + **Member**: Can typically view information but not change settings.
* **Custom Role**: A field that, when clicked, drops down a list from which you can choose one or more custom roles by name. The drop-down includes a filter field into which you can enter text to narrow the listed roles, which can then be selected individually (by clicking) or as a group (with the **Select All** button).

  + Each selected role will appear as a lozenge in the **Custom Role** and **Search** fields, and the **User** list will be filtered to users assigned these roles.
  + To add more roles, click in the field again.
  + To remove a role, click the **X** at the right of its lozenge.

### User List

The **User** list is a table that lists all of your organization's existing users. Click on a column heading to sort the list (ascending or descending). The information and actions displayed in the table depend on your user level.

#### User Control Columns

The following control columns are displayed to Administrators and Super Administrators:

* **Select all** (in heading row): A checkbox for toggling the state of all checkboxes in the Select column:

  + If either no checkboxes in the list itself are checked or only some are checked, then clicking this checkbox will select all listed users.
  + If all checkboxes in the list are checked, clicking this checkbox will deselect all users.
* **Select** (in user rows): Checkboxes at the left of each row that select individual users.

If the permissions granted by your Kentik roles include "Can assign users to roles" and you select users with either of the above controls, the **Role** button will appear at the top of the Users list (see **Role Selection Popup**).

#### User Information Columns

The following information columns are displayed to users of all levels:

* **Full name**: The user's full name.
* **Email**: The user's email address.
* **Level**: The user's user level (e.g., Admin; see **User Level**).
* **Roles**: The number of RBAC roles assigned to the user.
* **Last login**: The date of the user's most recent login.

#### User Admin Actions

The following actions (far-right column) are available only to admin-level users:

* **Reset active sessions**: Opens a confirming dialog that allows you to reset any active user sessions by logging the user out of the Kentik system.
* **Edit user**: Opens an Edit User dialog that allows you to change the user’s settings (see **User Settings Dialogs**).
* **Remove**: Opens a confirmation dialog that allows you to remove the user from your organization's Kentik account.

  ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(151).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A02Z&se=2026-05-12T09%3A41%3A02Z&sr=c&sp=r&sig=KPLdbmb%2BjZ7IxzXicZRyeRDuNOvTXdcgbo1gJGXPFKI%3D)

  *A list of all users in your organization*

#### Role Selection Popup

The Role Selection popup opens from the **Role** button at the top of the **User** list, which is present only for users whose RBAC permissions include "Can assign users to roles." This permission can be assigned to you in one of the following ways:![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(152).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A02Z&se=2026-05-12T09%3A41%3A02Z&sr=c&sp=r&sig=KPLdbmb%2BjZ7IxzXicZRyeRDuNOvTXdcgbo1gJGXPFKI%3D)

* You are assigned the Kentik Managed role of Administrator or Super Administrator:

  + Directly by a user in your organization;
  + Indirectly because your user level is Administrator or Super Administrator.
* You are assigned a Custom role that includes the permission, which may be:

  + The Kentik-provided role RBAC Administrator;
  + A different Custom role, created within your organization, that includes the permission.

The Role Selection popup enables you to assign one or more roles to all of the users that are checked in the User list. It includes the following elements:

* **Search**: A field into which you can enter text to narrow the list of available roles. You can clear entered text by clicking the **X** at the right.
* **Role list**: The complete list of roles available for your organization. Each element of the list includes:

  + **Checkbox**: Click to select this role.
  + **Role name**: The name of a Kentik Managed or Custom role.
  + **Role description**: A description of the role.
* **Apply**: A button that closes the popup and applies the selected roles to the checked users.

## User Settings Dialogs

Adding or editing a user via the Kentik portal involves specifying information in the fields of the User Settings dialogs, which are covered in the following topics.

> Notes:
>
> * The User Settings dialogs are accessible only to users whose level is Administrator or above.
> * Users can also be added and edited with the **User API**.

### About User Dialogs

The settings dialogs are used to collect and display user information. The required information is entered into the fields of either of the following dialogs:

* Add User when registering a new user with Kentik.
* Edit User when editing an already registered user.

### User Dialogs UI

The Add User and Edit Userdialogs share the same layout and the following common UI elements:

* **General Settings**: A tab that enables you to add or modify a user’s email address, name, and API token (see **User Field Definitions**).
* **User Level (legacy)**: A tab that enables you to set a user’s level of access for the Kentik portal (see **User Level**).
* **Filters**: A tab that enables you to add filters to limit the data seen by this user (see **User Filters**).
* **Roles & Permissions**: A tab that enables you to set a user’s RBAC roles and view their effective permissions (see **User Roles & Permissions**).
* **Cancel**: A button — either the **X** at upper right or the **Cancel** button at lower right — that exits the dialog without making changes to the user settings (all elements will be restored to their values at the time the dialog was opened).
* **Add User** (Add Userdialog only): A button that saves the settings for the new user and exits the dialog.
* **Save** (Edit User dialog only): A button that saves the changes to settings for an existing user and exits the dialog.

![User settings page displaying email, full name, and API token fields.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/MU-edit-user-dialog.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A02Z&se=2026-05-12T09%3A41%3A02Z&sr=c&sp=r&sig=KPLdbmb%2BjZ7IxzXicZRyeRDuNOvTXdcgbo1gJGXPFKI%3D)

*Specify settings for an individual user.*

### User Field Definitions

User settings dialogs contain the elements shown in the following table.

| Tab | Element | Type | Description |
| --- | --- | --- | --- |
| General Settings | Email address | Editable field | The user's email address. |
| General Settings | Full name | Editable field | The user's full name. |
| General Settings | API token (Edit User only) | Text field   (non-editable) | A Kentik-generated string that is used to authenticate the user in an HTTP header (X-CH-Auth-API-Token). |
| General Settings | Reset API Token | Button | Generates a new API token, which is displayed in the API Token field. |
| General Settings | Disable 2-Factor Authentication (Edit User only) | Button or text | If 2-factor authentication is enabled (see **Two-factor Authentication**), the button disables it.  If it’s not enabled, the text indicates that it's disabled. |
| User Level (legacy) | User Level (legacy) | Drop-down menu | Sets the user level (Member, Administrator, or Super Administrator; see **User Level**) for portal areas whose access is determined by user level. |
| Filters | Add Filters | Button | Opens **Filtering Options Dialog** to apply filters that limit the traffic seen by the user (see **User Filters**). |
| Roles & Permissions | Assigned Roles (Kentik-Managed & Custom) | Lozenges | A list of lozenges, one for each role applied to the user. Click X in a lozenge to remove that role from the user. |
| Roles & Permissions | Add Role | Button | Opens a popup listing RBAC roles that can be applied to the user (see **Role Selection Popup**). |
| Roles & Permissions | Enabled Only | Switch | Shows all possible permissions (disabled) or only the permissions assigned to the user (enabled). |
| Roles & Permissions | Effective Permissions | Table (non-editable) | Lists the permissions granted by the user's assigned roles (see **Effective Permissions**). |

### User Level

User levels are Kentik’s legacy system for determining which users can access or configure specific areas of the portal. User levels currently co-exist with our recently introduced Role-Based Access Control system (see **About Kentik RBAC**). When the RBAC rollout covers all areas of the portal, user levels will be fully deprecated and the **User Level** tab of the user settings dialogs will be removed.  
  
 For areas of the portal not yet covered by RBAC, Kentik supports three levels of users:

* **Member**: Can view settings and data, including traffic analytics and synthetic test results, in most portal modules, but can't add tests in Synthetics, edit settings for devices, cloud exports, users, etc., or edit properties in modules such as Interface Classification, Insights, Alerting, Mitigation, etc.
* **Administrator**: Can perform all actions except resetting SSO options or turning other users into Super Administrators.
* **Super Administrator**: Can perform all actions (see **About Super Admin Users**).

When you initially choose a user's level on the**User Level** tab, the user will also be assigned a corresponding Kentik Managed role, which will be shown on the **Roles & Permissions** tab. Once a user has been saved, however, changes to that user’s level have no effect on the user's RBAC roles.

> **Notes:**
>
> * Any user whose level is Administrator can determine whether a user is set as an Administrator or a Member; only a Super Administrator can designate another user as a Super Administrator.
> * If SSO is required, only a Super Administrator can log on with just their password (see **Additional Configuration Options**).
> * Any changes made to a user’s settings will take effect at the user’s next login.

### User Filters

User-based filtering enables Administrators to restrict the data that can be accessed by a given user's queries. The filters (ad hoc or saved) are set with a **Filtering Options Dialog** that is accessed from the Filters tab of the user settings dialogs. Depending on the current context, the filtering dialog opens from one of the following buttons:

* Add User dialog: **Add Filters** button.
* Edit User dialog: **Edit Filters** button if filters already exist for the user, otherwise **Add Filters** button.

Once a user filter is set for a given user, it is systematically appended (ANDed) with any query run by that user, including:

* **Data Explorer** or **Dashboard** queries in the Kentik portal.
* Queries via Kentik APIs (see **Query API**).

One use case for a user filter would be to allow certain members to query data only from routers with a description that contains “backbone.” Another use case would be to allow certain members to query data only for interfaces whose description contains CUSTOMER and whose site is Ashburn DC3 or Ashburn DC4 (see **About Sites**).

### User Roles & Permissions

The **Roles & Permissions**tab of the user settings dialogs enables you to assign/unassign RBAC roles to a user and to view the user's "effective" permissions. The tab is divided into two main areas:

* The roles assigned to the user appear as a list of lozenges at the left (see **Assigned Roles**).
* The permissions associated with the assigned roles are displayed in a table on the right (see **Effective Permissions**).

> **Note:** The permissions associated with a given role are set on the **RBAC Management page**.

![User role management interface displaying effective permissions and options for administrators.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/MU-edit-user-roles-permissions.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A02Z&se=2026-05-12T09%3A41%3A02Z&sr=c&sp=r&sig=KPLdbmb%2BjZ7IxzXicZRyeRDuNOvTXdcgbo1gJGXPFKI%3D)

*Assign roles to users and see their effective permissions.*

#### Assigned Roles

The pane at the left of the **Roles & Permissions** tab enables you to manage role assignment for the user and shows the user's assigned roles. This pane includes the following elements:

* **Kentik-Managed Roles**: A list of lozenges, one for each Kentik-managed role assigned to the user. To unassign a role, click the **X** at the right of its lozenge (see **Remove Roles from a User**).
* **Custom Roles**: A list of lozenges, one for each Custom role assigned to the user. To unassign a role, click the **X** at the right of its lozenge.
* **Add Role**: A button that opens a popup listing RBAC roles that can be applied to the user. Click on a role to assign it to the user and add its lozenge to the list (see **Assign Roles to a User**).

  > **Notes:**
  >
  > + To see a role's description, hover over its lozenge.
  > + When a user is first added, the Kentik-managed role (and its associated permissions) is “inherited” from the user level chosen on the **User Level (legacy)** tab (see **User Levels vs Permissions**).
  > + To assign the RBAC Administrator role to a user, you must have that role yourself or be a Super Administrator.

#### Effective Permissions

The **Effective Permissions** pane includes the following elements:

* **Permissions table**: A table listing the permissions associated with the roles assigned to the user:

  + The permissions are grouped by portal area, with each group having a heading row that includes the area's name and an expand/collapse arrow that controls the visibility of the permissions in that group.
  + Each permission has a checkbox at the left of its row. The checkbox is checked for enabled permissions and grayed out for other permissions.
  + Hover over any individual permission to highlight the lozenge(s) for the roles to which that permission is assigned.
* **Enabled Only**: A switch that shows (when off) or hides (when on) the permissions that aren't enabled

> **Notes:**
>
> * the list shows all permissions for the roles assigned to a user, it doesn't allow you to change the permissions for any given role (including Custom roles), which you can only do from the **RBAC Management page**.
> * Permissions always enable access and/or actions rather than preventing it.

## Add or Edit Users

Users are added and edited on the **Users Page** of the Kentik portal.

> Notes:
>
> * To access the Users page, choose **Users** from the **Organization Settings** menu on the main navbar.
> * Users whose user level or Kentik-Managed role is Member do not have the permissions required to add or edit a user

### Add a User

To register a new user from the Users page:

1. Click the **Add User**button to open the Add User dialog.
2. Specify the fields on the **General Settings** tab of the dialog (see **User Field Definitions**).
3. Set the level on the **User Level** tab.
4. Set any desired **User Filters**.
5. Assign one or more RBAC roles to the user (see **User Roles & Permissions**).
6. Click the **Add User** button (lower right) to save the new user and exit the dialog.

### Edit a User

To edit the settings for an existing user from the Users page:

1. In the **User List**, find the row corresponding to the user you'd like to edit.
2. Click the **Edit** button at the right of the row.
3. In the **Edit User** dialog:

   1. Change any settings that you'd like to modify (see **User Field Definitions**), including any **User Filters** or **User Roles & Permissions**.
   2. To reset the API token, click the **Reset API Token** button. A new token will appear in the field.
4. Click the **Save** button (lower right) to save the changes and exit the dialog.

### Remove a User

To remove a user from your organization's collection of Kentik users:

1. In the **User List**, find the row corresponding to the user you'd like to remove.
2. Click the **Remove** button (red trash icon) at the right of the row.
3. Click **Remove** in the resulting confirmation dialog.

### Assign Roles to a User

You can assign RBAC roles to multiple users simultaneously in the Users list, or in the settings dialog for an individual user (on the **Roles & Permissions** tab).

> **Notes**:
>
> * To modify another user’s RBAC roles, you must be an Administrator or Super Administrator, or have been assigned a custom role (e.g., RBAC Permissions) with the required permissions.
> * Roles can also be assigned to users on the **RBAC Management page**.

#### Assign Roles via User List

To assign a role via the User list:

1. In the **User List**, click the checkbox at the left of the row for each of the users to whom you'd like to apply an RBAC role. A **Role** button will appear at the top of the **Users** list.
2. Click **Role** to open the **Role Selection Popup**.
3. Click the checkboxes for the role(s) that you’d like to assign to the user(s). If needed, use the **Search** field to filter the listed roles.
4. Click **Apply**.

#### Assign Roles via User Dialog

To assign roles to an individual user via a settings dialog:

1. In the **User List**, find the row for the user to whom you’d like to assign roles.
2. Click the **Edit** button at the right of the row.
3. In the **Roles & Permissions** tab, click **Add Role** (bottom left) to open a popup listing roles that can be applied to the user.
4. Click a role to assign it to the user. The popup will close, a lozenge for the role will appear in the list at the left of the dialog, and the role's permissions will be added to the table in the **Effective Permissions** pane.
5. Repeat steps 3 and 4 as needed to assign more roles to the user.
6. Click the **Save** button to save the changes and exit the dialog.

### Remove Roles from a User

To remove a role from a user:

1. In the **User List**, find the row for the user whose roles you'd like to modify.
2. Click the **Edit** button (pencil) at the right of the row. The Edit User dialog will open.
3. In the list at the left of the **Roles & Permissions** tab, find the lozenge for the role you’d like to remove.
4. Click the **X** at the right of the lozenge.
5. Click **Save** to save the change and exit the dialog.

> **Note:** Roles can also be removed from users on the **Manage RBAC Roles Page**.
