# Kentik RBAC

Source: https://kb.kentik.com/docs/kentik-rbac

---

This article covers Role-Based Access Control (RBAC) in the Kentik portal.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(73).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A31Z&se=2026-05-12T09%3A43%3A31Z&sr=c&sp=r&sig=stt7iug29jqVtsHbLjAIMHpxVylS6sHWZ9UYWvJuyLY%3D)

*The settings page for RBAC enables you to manage both Kentik and Custom roles.*

## About Kentik RBAC

Kentik's rollout of Role-Based Access Control (RBAC) is a big step forward in enabling more granular management of user permissions in the Kentik portal.

### What is Kentik RBAC?

Our RBAC system uses roles and permissions to enable your organization to control access to Kentik features and settings:![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(78).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A31Z&se=2026-05-12T09%3A43%3A31Z&sr=c&sp=r&sig=stt7iug29jqVtsHbLjAIMHpxVylS6sHWZ9UYWvJuyLY%3D)

* **Permission**: The right to perform a specific action, e.g., “Edit specific Synthetic Monitoring Tests.” Permissions are organized into categories in relation to the portal area they   cover.
* **Roles**: A collection of permissions.

Permissions are assigned to roles, and roles are assigned to users, resulting in each user’s permissions including the permissions of all of their assigned roles. These assignments are managed on the **Manage RBAC Roles Page**, which is accessed via the **Organization Settings** menu that drops down from the main portal navbar.

> **Note**: Because permissions always enable — rather than disable — the ability to perform a permission-controlled action, such actions can be performed only by users whose roles explicitly grant the corresponding permission.

### Kentik-managed Roles

Kentik’s RBAC includes a mechanism called "Managed Roles" that simplifies the onboarding of new users to your Kentik account. With these roles — which are managed by Kentik — you don't need to set fine-grained permissions for every user. Instead, you can choose a ready-made role that grants permissions that are roughly equivalent to one of Kentik's existing user levels: Member, Administrator, and Super Administrator (see **About Users**). These roles will always contain reasonable defaults for all permissions available in Kentik Portal.

> **Note:** Kentik-managed roles are not editable by users.

### User Levels vs Permissions

Kentik's current system of user levels grants implicit permissions that apply throughout the portal. RBAC will eventually replace this model. As we gradually add permissions for other areas of the portal into RBAC, we will remove them from the user-level framework, meaning that access to an RBAC-enabled feature will no longer be regulated by a user's level.  
  
The above means that user levels and RBAC roles are going to coexist for as long as necessary, and that the assignment of a level will remain mandatory when adding a new user. This user level will be leveraged to dynamically assign the relevant Kentik-managed role.

### From User Levels to Roles

The transition to full RBAC implementation will be gradual, with user levels and RBAC co-existing for a considerable time. User levels will be active for portal areas that aren’t covered by RBAC while new RBAC-compatible permissions will be governed by the new RBAC engine.  
  
In our initial rollout, we are leveraging RBAC roles and permissions to replace a soon-to-be-deprecated functionality called Permission Overrides. To do so, we are porting these permissions into the new RBAC model, after which users’ permission overrides will be phased out.

### Migration for Existing Users

We’ve designed the migration to RBAC to be transparent, meaning that without exception existing users will keep the same level of permissions they had before RBAC. The specifics of this migration depend on the permissions currently granted to a user:

* Most users without permission overrides will be assigned to the Kentik-managed role that corresponds to their current user level.
* Users with permission overrides (only for Synthetics and Connectivity Costs) will be assigned a set of dynamically generated custom roles instead of a single Kentik-managed role. As we roll out additional RBAC permissions, you'll need to keep these custom roles updated within your organization.

### RBAC for Connectivity Costs

One additional aspect of the RBAC transition is a change related to **Connectivity Costs**. In the user-level framework, the Connectivity Cost workflow is available for viewing by the Members, and for both viewing and configuration by Administrators and Super Administrators. In our new RBAC-based Kentik-managed roles, however, we are — in response to popular demand — removing Connectivity Costs access for the Member role. To keep this migration transparent, existing level-based Members will be assigned a special RBAC role ("Connectivity Costs Viewer") that will maintain their existing access. However, for new users assigned the Member role, a custom role is needed to see Connectivity Costs.

### Initial RBAC Permissions

As noted earlier, RBAC permissions aren't portal-wide like user levels. Instead, they govern access in specific portal areas, the number of which will expand over time. At initial rollout, RBAC covered the following three areas (the second two areas replace permission overrides):

* RBAC Management
* Synthetic Tests and Agents
* Connectivity Costs

The following areas have been transitioned to RBAC subsequent to rollout:

* Dashboards
* Saved Views

All other portal areas are still regulated by user levels.

> **Note:** The above information will change as we actively expand RBAC implementation. For a current list of areas covered by RBAC, see the **Roles List**.

### What’s Coming Next?

Our initial RBAC release establishes the foundation of the RBAC permission engine. In the coming quarters we'll build on this foundation by expanding the RBAC model into new areas:

* We are working to extend **Labels** to multiple areas of the portal beyond **Devices**, **Synthetics Agents**, and Synthetics tests (see **Test Control Center**), and we want to extend RBAC permissions to apply to content grouped together by users using labels.
* We will extend label-based RBAC permissions to Synthetics agents and tests, and will shortly follow with **Dashboards** and **Saved Views**, allowing  the permissions for content created by users to be managed both centrally and granularly.
* Our **Credentials Vault** will be upgraded shortly with the ability to manage secrets based on labels.

More than anything, we would love to hear your thoughts on which area of the Kentik portal you would like us to work on implementing next, so please do let us know!

## Manage RBAC Roles Page

The Manage RBAC Roles page is used to manage your organization’s RBAC roles, and the permissions and users assigned to each role.

> **Notes**:
>
> * To access the Manage RBAC Roles page, choose Manage RBAC Roles from the **Organization Settings** menu on the main navbar.
>
>   Access to this page is governed by the RBAC Viewer permission.

### Manage RBAC Roles Page UI

The Manage RBAC Roles page displays the RBAC Roles list, which is a table listing the currently available Kentik-managed and custom roles. Users whose roles grant them the appropriate RBAC permissions can add, edit, or delete custom RBAC roles using this list.  
  
 The Manage RBAC Roles page includes the following main areas and controls:

* **Create a Role**: A button that opens the Add Role dialog (see **RBAC Role Dialogs**).

  > **Note:** This button is present only for users with permissions equivalent to those granted by the Kentik-managed roles of Administrator or Super Administrator, including users with the role RBAC Administrator.
* **Search**: A field that filters the **Roles List** to rows whose name or description contains a match for the entered text.
* **Roles List**: A table listing your organization’s Kentik-managed and custom roles (see **Roles List**).
* **Results Listed**: An indicator at the bottom right of the Roles list that gives the current (filtered) and total number of roles in the list.

### Roles List

The **Roles** list is a table that lists all of your organization's roles. The table includes two groups that each start with a heading row (see **Role Heading Rows**), one for Kentik-managed roles and the other for custom roles.

#### Roles List Columns

The **Roles** list includes the following columns:

* **Role Name**: The name of the role.
* **Description**: The description of the role.
* **# Users**: The number of users to which this role is assigned.
* **Created**: The date and time when the role was first created.
* **Last Modified**: The date and time when the role was last edited.
* **View** (eye icon): A button that opens the View Role dialog (see **RBAC Role Dialogs**).
* **Edit** (pencil icon): A button that opens the Edit Role dialog.
* **Remove** (trash icon): A button that opens a confirming dialog from which you can remove the role from your organization.

> **Note**: Kentik-managed roles cannot be removed or modified, including their permissions, names, and descriptions.

#### Role Heading Rows

The rows of the table are divided by heading rows into the following two groups:

* **Kentik-managed Roles**: Three roles that are editable only by Kentik and are modeled on the corresponding Kentik user levels:

  + Super Administrator: Can perform the most sensitive operations (including setting RBAC permissions for all users).
  + Administrator: Can perform some sensitive operations.
  + Member: Regular, non-administrator users.
* **Custom Roles**: Roles that can be created and managed by users who are assigned the Kentik-managed role of Super Administrator or Administrator, or custom role (e.g., RBAC Permissions) that includes sufficient permissions to configure RBAC roles.

The heading row for each of the above groups includes a caret at the left that you can click to expand or collapse the group.

## RBAC Role Dialogs

Viewing, adding, or editing an RBAC role in the Kentik portal involves interacting with one of the three Role dialogs: **View Role**, **Add Role**, and **Edit Role**.

### About Role Dialogs

The Kentik portal uses Role dialogs to collect and display information about RBAC roles. The dialogs that are available to you depend not only on the type of the role for which you'd like to see information but also on the permissions associated with your own roles. The following examples apply to Kentik-managed roles (which may be assigned to you either directly or because of your user level):

* **Super Administrator or Administrator**:

  + Add Role is used when creating a new custom role. Enter the required information into the dialog's fields (see **Role Dialogs UI**).
  + Edit Role is used to edit an existing custom role, including managing the users assigned to that role (see **Role Users Tab**).
  + View Role is used to see details about an existing Kentik-managed role, as well as to manage the users assigned to that role.
* **Member**: View Role is used to see details about an existing role, whether Kentik-managed or custom.

### Role Dialogs UI

The **Add Role**, **Edit Role**, and **View Role** dialogs share the same layout and the following common UI elements:

* **Name**: A field for the role’s name.
* **Description**: A field for a description of the role.
* **Tabs**: The dialog includes two tabs:

  + Permissions: Used to see/assign the permissions of the role (see **Role Permissions Tab**).
  + Users: Used to see/assign the users of the role (see **Role Users Tab**).
* **Cancel**: A control — either the **X** at the upper right or the **Cancel** button at the lower right — that exits the dialog while leaving all **Permissions** tab settings as they were when it opened.

  > **Note:** Changes to the Role Users list are saved at the time they are made, even if you subsequently exit the Role dialog with **Cancel**.

**Save** (Add/Edit Role dialogs only): Saves the dialog’s current settings and exits the dialog.

![Role editing interface showing permissions for Connectivity Costs Administrator and RBAC settings.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/RBAC-edit role.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A31Z&se=2026-05-12T09%3A43%3A31Z&sr=c&sp=r&sig=stt7iug29jqVtsHbLjAIMHpxVylS6sHWZ9UYWvJuyLY%3D)

*The Role dialog enables assignment of permissions to roles and roles to users.*

### Role Permissions Tab

The **Permissions** tab of the Role dialogs enables users to view and (except for Members) set the permissions of a role.  
  
 The tab includes the following UI elements:

* **Search**: A field that filters the **Permissions List** to show only permissions containing the entered text.
* **Enabled only**: A switch that, when on, filters the **Permissions List** to show only permissions that are already enabled for the role.
* **Portal sections**: A set of buttons to the left of the **Permissions List**. Each button corresponds to one section of the portal for which permissions can be set for this role. For each section, the number of permissions currently set in this role is indicated by the number at the right of the section's button. The buttons toggle between an on state and an off state:

  + On: The existing content of the **Search** field (if any) is replaced with button's name, filtering the list to show only the set of permissions available for the corresponding section of the portal.
  + Off: The button's name is cleared from the **Search** field, leaving the list unfiltered.

    > **Note:** If the text in the Search field is changed while a button is on, the button reverts to the off state.
* **Permissions pane**: Controls that enable you to set permissions for the role. See **Permissions Pane**.

> **Note**: While theView Roledialog includes all of the elements described above, you cannot modify the role’s permissions.

#### Permissions Pane

The **Permissions** pane is the area of the **Permissions** tab that displays (View Role dialog) or allows you to set (Add Role and Edit Role dialogs) the permissions for the role.  
  
The**Permissions** pane includes the following UI elements:

* **Select All**: A checkbox at the top left of the pane checks the checkboxes of all permissions in the Permissions list (visible or not). When checked, click the **Select All** checkbox to clear the checkboxes of all permissions.
* **Permissions list**: The list of permissions available to assign for the selected role (see **Permissions List**).
* **Results**: The number of permissions listed in the Permissions list is displayed at the bottom right of the **Permissions** pane.

#### Permissions List

The Permissions list is a table listing all available permissions for the role. The list is filtered by any text entered in the Search field (see Role Permissions Tab). The Permissions list includes the following UI elements:

* **Heading rows**: A row indicating a group of permission for a given area of the portal. An expand/collapse control at the left of the row toggles the visibility of the group's permissions.
* **Permissions**: Each row within a group represents a permission that grants access and/or allows certain actions. You can assign/unassign an individual permission to the role with the checkbox at the row's left.

  > **Note**: The Select All checkbox above the list toggles On/Off all of the individual checkboxes.

### Role Users Tab

The Users tab of each Role dialog displays the users that have that role applied to them. Depending on your permissions, you can view and/or modify the users listed.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(93).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A31Z&se=2026-05-12T09%3A43%3A31Z&sr=c&sp=r&sig=stt7iug29jqVtsHbLjAIMHpxVylS6sHWZ9UYWvJuyLY%3D)

*The Users tab of the Role dialog.*

The Users tab includes the following UI elements:

* **Search**: A field that filters the User list to show only users containing the entered text in one of the list's three columns.
* **Add User**: A button that drops down a list of all users that aren’t already assigned the role. See **Add User Menu**.
* **Role Users List**: A table that displays all users that have been assigned this role (see **Role Users List**).

#### Role Users List

The list features the following columns and elements:

* **User Name**: The user's Kentik username.
* **User Full Name**: The user’s name as it’s been entered in the **General Settings** tab of the Edit User dialog.
* **User Email**: The user’s email address.
* **Remove** (trash icon): A button that opens a confirmation dialog in which you can remove the user from this role.
* **Results**: An indicator at the bottom right of the list that shows the number of users currently displayed in the list.

#### Add User Menu

The Add User menu lists all users that haven’t already been assigned the role. The menu, which drops down from the**Add User** button on the **Users** tab, includes the following elements:

* **Search**: A field that filters the list of users to show only those whose name contains the entered text.
* **Users**: A set of rows, one for each user. Each row shows the user’s full name as entered in the **General** **Settings** tab of the Edit User dialog. A checkbox to the left of the name enables you to select the user.
* **Apply**: A button that adds the selected users to the **Role Users** list on the **Users** tab.  
   The button is inactive unless one or more users are selected.
* **Cancel**: A button that closes the menu without adding any users to the **Role Users** list.

> **Note**: Changes to the Role Users list are saved immediately, regardless of whether you subsequently exit the Role dialog with **Cancel** or **Save****.**

## Manage RBAC Roles

The following procedures cover basic operations related to managing RBAC roles and permissions in the Kentik portal.

> **Note**: With the exception of **View a Role**, this section’s procedures are only possible for users whose role is Member if they've been assigned a custom role (e.g., RBAC Permissions) that includes the corresponding RBAC permissions.

### Add a Role

To create a custom RBAC role in your organization:

1. From the **Organization Settings** menu on the main portal navbar, choose Manage RBAC Roles to open the Manage RBAC Roles page.
2. Click the **Create a Role** button at top right.
3. In the resulting **Add Role** dialog, specify the settings covered in **Role Dialogs UI**.
4. Click **Save**. The dialog will close, and your role will appear in the **Custom Roles** group.

> **Note**: When you add a custom role, the Users tab will not appear until you save the role and open the Edit Role dialog.

### View a Role

The View Role dialog is visible to users who are assigned a role that allows them to view RBAC roles, but not edit them. To view a role’s permissions and the users to whom it is assigned:

1. From the **Organization Settings** menu on the main portal navbar, choose Manage RBAC Roles to open the Manage RBAC Roles page.
2. In the Roles list, click the view button (eye) in the row for the role you’d like to view. The View Role dialog will open.
3. When done, click Cancel to close the dialog.

### Remove a Role

To remove a custom role from your organization’s list of roles:

1. From the **Organization Settings** menu on the main portal navbar, choose Manage RBAC Roles to open the Manage RBAC Roles page.
2. In the **Roles** list, find the row for the role you’d like to remove.
3. Click the **Remove** button (trash can) at the right of the row.
4. Click **Remove** in the resulting confirmation pop-up. The role will be permanently removed from your organization.

> **Note**: Kentik-managed roles cannot be removed from the RBAC Roles list.

### Manage Role Permissions

To assign or remove permissions to a custom role:

1. From the **Organization Settings** menu on the main portal navbar, choose Manage RBAC Roles to open the Manage RBAC Roles page.
2. In the **Roles** list, find the row for the role whose permissions you'd like to manage.
3. Click the **Edit** button (pencil) at the right of the row to open the Role dialog.
4. On the **Permissions** tab:

   1. Check the checkboxes for any permissions you’d like to enable for this role.
   2. Uncheck the checkboxes for any permissions you’d like to remove from this role.
5. Click **Save** to save changes to the permission, exit the dialog, and return to the Manage RBAC Roles page.

> **Note**: You cannot modify permissions for Kentik-managed roles.

### Assign Roles to Users

Users whose roles include sufficient RBAC permissions can assign roles to users in two separate areas of the portal: the Users page (User dialogs) and the RBAC Roles page (Roles dialogs).

#### Assign Roles on Users Page

On the Users page, roles can be assigned to individual users or to a selected set of users:

* To assign roles to a user in a User dialog (on the Users page), see **Assign Roles via User Dialog**.
* To assign one or more roles to multiple users simultaneously, see **Assign Roles via User List**.

#### Assign a Custom Role

To assign an existing custom role to a user:

1. From the **Organization Settings** menu on the main portal navbar, choose **Manage RBAC Roles** to open the Manage RBAC Roles page.
2. In the **Roles** list, find the row for a role that you'd like to assign to assign to the user.
3. Click the **Edit** icon (pencil) at the right of the row to open the Edit Role dialog.
4. On the **Users** tab, click the **Add User** button, which drops down a list of available users.
5. Use the **Search** field (or scroll) to find the user, then select the user’s checkbox.
6. Click **Apply** to assign the role to the user.
7. Click **Save** to close the Edit Role dialog and return to the Manage RBAC Roles page.

#### Assign a Kentik-Managed Role

To assign a Kentik-managed role to a user:

1. From the **Organization Settings** menu on the main portal navbar, choose Manage RBAC Roles to open the Manage RBAC Roles page.
2. In the **Roles** list, find the row for the Kentik-managed role that you'd like to assign to the user.
3. Click the **View** icon (eye) at the right of the row to open the View Role dialog.
4. On the **Users** tab, click the **Add User** button, which drops down a list of available users.
5. Use the **Search** field (or scroll) to find the user, then select the user’s checkbox.
6. If needed, repeat step 5 to add more users.
7. Click **Apply** to apply the role to the user(s).
8. Click **Cancel** to close the View Role dialog and return to the Manage RBAC Roles page.

### Remove Roles from Users

In general, you can unassign a role from a given user in either of the following locations:

* From the user's settings on the Users page.
* From the role's settings on the Manage RBAC Roles page.

> **Note**: Kentik-managed role can only be removed via the Users page.

#### Remove Role from User

To remove a custom or Kentik-managed role from a user using the Edit User dialog (from the Users page), see **Remove Roles from a User**.

#### Unassign User from Role

To remove a custom role from a user via the Manage RBAC Roles page:

1. From the **Organization Settings** menu on the main portal navbar, choose Manage RBAC Roles to open the Manage RBAC Roles page.
2. In the **Roles** list, find the row for the role from which you’d like to remove the user.
3. Click the **Edit** button (pencil) at the right of the row to open the **Edit Role** dialog.
4. On the **Users** tab, use the **Search** field or scroll in the **Role Users List** to find the user that you'd like to remove from the role.
5. Click **Remove** (trash can) at the right of the row for the user.
6. Click the **Remove** button in the confirmation pop-up.
7. Click **Save** or **Cancel** to close the **Edit Role** dialog and return to the Manage RBAC Roles page.

> **Note:** Changes to the Role Users list are saved at the time they are made, even if you subsequently exit the Role dialog with **Cancel**.
