# My Kentik Portal

Source: https://kb.kentik.com/docs/my-kentik-portal

---

This article discusses the configuration of My Kentik Portal (MKP), which is Kentik's implementation of tenancy.

> **Note:** For information on how your tenant users see and interact with your implementation of MKP, see **MKP Tenant Portal**.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(303).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A48Z&se=2026-05-12T10%3A02%3A48Z&sr=c&sp=r&sig=KS%2Fc%2Bvp6WcwNEP7unJtgajVTK%2B74vz4yhr35ENk01Bk%3D)

*Manage tenants and packages from the tabs of the My Kentik Portal page.*

## About My Kentik Portal

My Kentik Portal (MKP) is a mechanism by which you can make a limited, customer-branded version of the Kentik portal available to your external or internal customers (“tenants”). You can assign users to a tenant, enabling them to see a curated set of visualizations and metrics that are filtered to show only their own traffic.

> **Note:** As used in the context of tenancy, the term “user” does not refer to a registered Kentik user, but rather to a user of an MKP instance (tenant portal).

### Tenancy Structure

Kentik's implementation of tenancy is structured as follows:

* **My Kentik Portal**: Each Kentik customer may set up a single MKP tenant portal through which their customers (i.e. tenants) can enable access to a simplified version of the Kentik portal (see **MKP Summary**).

  + **Tenants**: A Kentik customer's MKP can have multiple tenants (see **Tenant Summary**).

    - **Users**: Each MKP tenant can have an unlimited number of users.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/MKP-Tenancy_diagram-353h672w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A48Z&se=2026-05-12T10%3A02%3A48Z&sr=c&sp=r&sig=KS%2Fc%2Bvp6WcwNEP7unJtgajVTK%2B74vz4yhr35ENk01Bk%3D)

*A sample MKP implementation structure.*

#### MKP Summary

As a Kentik customer, you can set up a single tenant portal, called My Kentik Portal (MKP), where you are effectively the "landlord" for your "tenants" (your customers):

* The tenant portal is a much-simplified version of the Kentik portal, including:

  + An Explore Your Network page with a **Library** of **Dashboards** and **Saved Views** that you specify on the **Views** tab of the tenant's settings page (see **Tenant View Settings**).
  + Information related to alerting, notifications, and mitigation.
  > **Notes:**
  >
  > + For more information on what tenants can see in their portal, see **MKP Tenant Portal**.
  > + Dashboards and Saved Views from both **Data Explorer** and **Metrics Explorer** are supported in MKP.
* The portal URL is the same for all of your tenants.
* The portal is branded with your logo and may optionally include your support information.

#### Tenant Summary

As a Kentik customer, your My Kentik Portal (MKP) exists to provide your tenants with timely information on the state of your network:

* If you’re a provider, each tenant typically corresponds to one of your subscribers.
* The maximum number of tenants per customer depends on your Kentik edition (see **Kentik Editions**).
* The types of reports available to a given tenant, and the characteristics (e.g., devices, CIDRs, etc.) of the shown traffic, are specified with the settings on the **Tenant Settings Pages**.
* Common settings for multiple tenants can be implemented with "packages" (see **Packages Tab**).

### Tenancy Use Cases

Tenancy is designed for two main use cases:

* The Kentik customer is a service provider, and each tenant represents a customer of the provider. Each tenant has multiple users who can access the tenant portal.
* The Kentik customer is an enterprise, and each tenant represents an internal user group with varying network traffic data needs and presentation/view preferences.

> **Note**: User views in the tenant portal can differ based on their assigned tenant.

### Enabling Tenancy

The tenancy feature is enabled by default, but the number of tenants supported by your Kentik subscription is determined by your organization's edition of Kentik (see **About Licenses**).

> **Note**: To obtain unlimited tenant support, you can purchase Kentik's Advanced MKP option. For details, contact Kentik (see **Customer Care**).

### Tenants and Packages

MKP tenancy configuration on the **My Kentik Portal Page**of the Kentik portal is structured into two parts:

* **Tenants**: Each tenant is an external or internal entity to whom you, as a Kentik customer, make available a self-branded version of the Kentik portal so that the tenant's users can access a limited subset of information about your network's status and performance.
* **Packages**: Packages enable you to group and assign views and policies to multiple tenants rather than having to repeatedly assign the individual items. Each package is a collection of the following:

  + **Views**: One or more individual views (charts, tables, etc.) that provide the tenant's users with information about the network. The MKP seen by your tenant's users will be a dashboard containing these views.
  + **Alert policies**: A set of alert policies that trigger alarms (notifications) to your tenant's users when a set of custom-defined conditions are met.

General guidelines for the relationship between tenants and packages:

* When a package is assigned to a given tenant, all of that tenant's users see all of the views (and get the alarms) associated with that package.
* Each tenant can currently be assigned only one package at a time.
* A given tenant may also be assigned as many individual views as you (the Kentik customer) choose to assign them.

## My Kentik Portal Page

The My Kentik Portal (MKP) page is home to the settings needed to configure and manage tenancy. To view the page, choose **My Kentik Portal** from the Kentik portal’s main navbar.  
  
 The My Kentik Portal page has the following main UI elements:

* **Settings** (button): Opens the **MKP Settings Page**.
* **Tenants tab**: Contains the Tenants list and other controls for configuring tenants of your organization's My Kentik Portal (see **Tenants Tab**).
* **Packages tab**: Contains the Packages list and other controls for configuring the views and alert policies for your organization's MKP tenants (see **Packages Tab**).

### Tenants Tab

The **Tenants** tab of the My Kentik Portal page, used to create and manage tenants (see **Tenancy Structure**), contains the following main UI elements:

* **Search field**: Filters the Tenants list to show only the tenants whose name or ID contains the entered text.
* **Tenant usage**: A statement showing how many tenants you are currently using and how many are allowed under your organization's edition of Kentik (see **About Licenses**).
* **Add Tenant**: A button that opens the Add Tenant page (see **Tenant Settings Pages**).
* **Tenants List**: A list of the tenants currently set up in your organization (see **Tenants List**).

#### Tenants List

The Tenants list shows all of your organization’s tenants (see **Tenancy Structure**). The table includes the following columns:

* **Status**: Indicates whether the tenant is currently enabled or disabled.
* **ID**: The tenant’s unique, Kentik-generated ID.
* **Name**: The name of the tenant and the tenant's description (if provided).
* **Package**: The name of the package, if any, assigned to this tenant (see **Tenants and Packages**).
* **Users**: The number of users assigned to this tenant (see **Tenant User Settings**).
* **Views**: The number of views assigned to this tenant, either individually or as part of a package (see **Tenant View Settings**).
* **Policies**: The number of alert policies assigned to this tenant, either individually or as part of a package (see **Tenant Policy Settings**).
* **Mitigations**: The number of mitigations assigned to the tenant’s policies, either individually or as part of a package (see **Threshold Mitigations**).
* **Last login**: The date-time of the last login by any user assigned to this tenant.
* **Actions**: A vertical dots icon that pops up the **Tenant Action Menu**.

#### Tenant Action Menu

The tenant Action menu is accessed via the vertical dots icon at the right of each row in the Tenants list. The menu includes the following options:

* **View as Tenant** (users icon): Takes you to the tenant's implementation of MKP (see **MKP Tenant Portal**).
* **Edit Tenant** (pencil icon): Takes you to the tenant’s Edit Tenant page (see **Tenant Settings Pages**).
* **Delete** (trash icon)*:* Pops up a confirmation dialog that allows you to delete the tenant from your collection of tenants.

### Packages Tab

The **Packages** tab of the My Kentik Portal page, used to create and manage packages (see **Tenants and Packages**), contains the following main UI elements:

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(314).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A48Z&se=2026-05-12T10%3A02%3A48Z&sr=c&sp=r&sig=KS%2Fc%2Bvp6WcwNEP7unJtgajVTK%2B74vz4yhr35ENk01Bk%3D)

*The Packages tab lists the packages in your organization*

* **Search packages**: A field into which you can enter text to filter the listed packages to those with matching name or description.
* **Add Package**: A button that opens the Add Package page (see **Package Settings Pages**).
* **Packages list**: A list of the packages currently set up in your organization (see **Packages List**).
* **Package Overview**: An information pane at the right of the page that shows information about the packages in use in your organization:

  + **Total**: The total number of packages defined in your organization.
  + **Unused**: The number of defined packages that haven't yet been assigned to any tenant.
  + **With/Without Package*s***: The number of tenants with packages/without packages.
  + **Tenants per Package**: A ring chart showing the number of tenants to which each of the in-use packages have been assigned (hover over any segment to see a pop-up giving the exact number for the corresponding package).

#### Packages List

The Packages list shows all your organization’s packages. Each row includes the following elements:

* **Icon**: The icon assigned to the package.
* **Package Name**: The name of the package.
* **Description**: A description of the package (if provided).
* **Tenants**: The number of tenants to which this package has been assigned.
* **Views**: The number of views assigned to this package.
* **Policies**: The number of policies assigned to this package.
* **Edit**: A link to the Edit Package page for this package (see **Package Settings Pages**).
* **Remove***:* A link that opens a confirmation popup that allows you to remove the package from your organization's collection of packages.

## MKP Settings Page

The MKP Settings page is where you specify the shared settings for all tenants in your organization. Use the tabs in the left sidebar to accomplish this, as covered here.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(315).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A48Z&se=2026-05-12T10%3A02%3A48Z&sr=c&sp=r&sig=KS%2Fc%2Bvp6WcwNEP7unJtgajVTK%2B74vz4yhr35ENk01Bk%3D)

*Specify shared settings for all of your tenants using the MKP Settings page.*

### Domain Tab

Use the **Domain** tab of the MKP Settings page to set the subdomain to include in the URL that tenant users will use to navigate to your tenant portal:

* Specify the subdomain name in the **Subdomain** field.
* The rest of the URL is predefined by Kentik as shown in the following example:  
  `https://customer_subdomain.my.kentik.com`

> **Note:** If your organization is registered with Kentik in the EU, the above URL will instead end with `my.kentik.eu`.

### Support Links Tab

The **Support Links** tab of the MKP Settings page is used to set the email address and URL that users added by your tenants should use when they need support.

* **Email Address**: The email address at which your organization will provide support for tenant users, and from which emails may be sent to tenant users. Emails auto-generated by MKP may include alerts sent via an email notifications channel and invitations sent when the user is added to a tenant.

  > **Important**: If no tenant support email address is configured, **support@kentik.com** will be the sender for MKP user activation emails.
* **URL**: The URL at which your organization will provide support for tenant users. This support link will be displayed in various locations in MKP and also in emails to tenant users.

### Branding Tab

The settings on the **Branding** tab of the MKP Settings page determine the branding elements that are displayed to tenant users of your MKP implementation:

* **Company Logo**: A field with which you can set an image that will be displayed in your My Kentik Portal in place of the Kentik logo. Click the **Browse** button to open a dialog from which you can select a file from your computer. Images can be:

  + **Max file size**: 1MB.
  + **Supported file types**: JPG, PNG, GIF, BMP.
  > **Notes:**
  >
  > + For best results, use a logo that is legible at a height of 21px.
  > + Once an image file has been uploaded, a thumbnail of the image will appear to the right of the **Logo** field.
* **Hide Kentik Branding**: Indicates whether Kentik will be mentioned in the UI of your MKP tenant portal. By default, the setting is off (mentions of Kentik are not hidden).

  > **Note:** This setting is actually made by Kentik in response to a request from your organization (see **Customer Care**) to license the Advanced MKP Package, which allows unlimited MKP tenants. Once it is turned on, it can only be reverted by dropping that package.

### Tenant SSO Tab

You have the option of enabling tenant users to access your My Kentik Portal (MKP) via single sign-on (SSO, see **About SSO**). You do this from the **Tenant SSO** tab of the MKP Settings page.

> **Notes**:
>
> * Only Super Admin users (see **About Super Admin Users**) can administer SSO. If you don’t yet have a Super Admin user in your organization, please contact **support@kentik.com**.
> * See **Kentik SSO Configuration** for details about completing this tab, since the non-tenant SSO settings are nearly identical.

#### SSO Login URL

In order for MKP users to login via SSO, they’ll need the login URL, which will follow this pattern:

`https://[custom_mkp_domain].my.kentik.com/sso/[landlord_company_name]`

**Example URL**: `https://uits.my.kentik.com/sso/uofa`

## Tenant Settings Pages

The settings pages for tenants are covered here.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(317).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A48Z&se=2026-05-12T10%3A02%3A48Z&sr=c&sp=r&sig=KS%2Fc%2Bvp6WcwNEP7unJtgajVTK%2B74vz4yhr35ENk01Bk%3D)

*Specify settings that are specific to an individual tenant.*

### About Tenant Settings

The Add Tenant and the Edit Tenant pages, collectively referred to as the "tenant settings pages," both contain the same structure, layout, and settings. These pages effectively function as wizards on which settings are specified in sequence on a series of tabs. You can navigate the tabs out of order using the tab names in the left sidebar, but you can't save your changes unless all required settings are specified.

Tenant settings are used to configure and manage MKP for an individual tenant, thus determining what Kentik-derived status and performance information will be presented to that tenant's users (see **Tenancy Structure**).

Settings for a given tenant may be accessed in the following ways:

* **Add Tenant**: If the tenant doesn't already exist in your organization, add the tenant via the **Add Tenant** button on the **Tenants Tab** of the My Kentik Portal page, which takes you to an Add Tenant page.
* **Edit Tenant**: If the tenant already exists in your organization, you can edit the tenant's settings on the Edit Tenant page for that tenant, which you reach by clicking on the tenant in the Name column of the **Tenants List** on the My Kentik Portal page.

### General Tab (Tenant)

The **General** tab of the tenant settings pages contains the following settings and controls:

* **Name** (required): The name by which the tenant will be listed in the **Tenants List**.
* **Description**: A description for the tenant.
* **Tenant Status**: A switch that lets you enable or disable MKP for this tenant. Disabling a tenant allows you to keep that tenant configured while blocking their access to the MKP tenant portal.
* **Cancel** (Add Tenant page only): This button takes you back to the My Kentik Portal page without adding a tenant.
* **Next** (Add Tenant page only): This button advances you to the next tab of the page. The Next button appears only when all required settings on the tab have been specified.
* **Save** (Edit Tenant page only): This button saves the settings you’ve made to the tenant and returns you to the **Tenants** tab on the MKP page.

### Data Sources Tab

The **Data Sources** tab of the tenant settings pages allows you to specify which network parts a tenant can access data from. This defines a partition of your flow data (device or cloud) and grants the tenant's users access only to that partition. You can specify access in two ways:

* **Data sources only**: Enable tenant access based on the data sources specified in the **Choose Data Sources** pane.

  > **Note:** Typically used only when each data source handles flow records from a single tenant.
* **Data sources plus filtering**: Enable tenant access based on the data sources specified in the **Choose Data Sources** pane and also based on applying additional filters in the **Apply Filtering Criteria** pane. The additional filters can include interfaces, source and/or destination IP ranges, ASNs, and **Custom Dimensions**.

> **Note:** Kentik recommends assigning a unique label to each data source associated with a tenant. If you use a tenant’s label to specify data sources in the MKP Data Sources dialog, the MKP data sources for that tenant will automatically change as you add or remove data sources to which the label is applied (see **Labels**).

#### Data Sources Tab UI

The **Data Sources** tab contains the following settings and controls:

* **Choose Data Sources**: A pane whose controls enable you to choose the data sources for this tenant:

  + **Data sources**: Lists the network data sources that this tenant can see. Each source is represented by a lozenge. Click the **X** in a source's lozenge to remove that source.
  + **Edit Data Sources** (button): Opens the **Data Sources Dialog****,** from which you can choose data sources (devices and clouds).
* **Advanced Filtering Options** (button): Expands the **Apply Filtering Criteria** pane.
* **Apply Filtering Criteria**: A pane, hidden by default, for applying filters to narrow the network data this tenant can see. Also used to set the tenant’s data retention period (see **Advanced Filtering Criteria**).
* **Cancel** (button, Add Tenant only): Returns you to the My Kentik Portal page without adding a tenant.
* **Next** (button, Add Tenant only): Advances you to the next tab of the page. Appears only when all required settings on the tab have been specified.
* **Save** (button, Edit Tenant only): Saves the tenant’s settings and returns you to the Tenants tab on the MKP page.

#### Advanced Filtering Criteria

Use the **Advanced Filtering Criteria** pane to filter the network data this tenant can see from the data sources selected with the **Data Sources Tab UI**. The filter criteria are applied as follows:

* ANDed (match all) between fields.
* ORed (match any) within fields that support multiple values.

> **Notes**:
>
> * Display the Advanced Filtering Criteria pane by clicking the **Advanced Filtering Options** button.
> * Any dimension (data source or filtering criteria) that you add to the tenant must also be used by the tenant’s policies. For example, if you populate the **ASN** field in the Advanced Filtering Criteria pane, each policy you select to apply to the tenant must also use the ASN dimension.

This pane includes the following:

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(329).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A48Z&se=2026-05-12T10%3A02%3A48Z&sr=c&sp=r&sig=KS%2Fc%2Bvp6WcwNEP7unJtgajVTK%2B74vz4yhr35ENk01Bk%3D)**Show Tenant Filters** (button): Opens a popup showing all filters currently applied in the **Advanced Filtering Criteria** pane.

* **Interface**: Specify regular expressions (regex) to match against the following fields of interfaces on the devices selected with the Data Sources dialog.

  + **Interface Name**: Traffic will be matched based on the vendor-defined name (e.g., “GigabitEthernet0/1”) of the device interface (physical or logical) through which flow ingressed/egressed.
  + **Interface Description**: Traffic will be matched based on the user-provided description (e.g., “Connected to upstream ISP”) of the device interface (physical or logical) through which flow ingressed/egressed.
* **CIDR(s)**: Enter a comma-separated list of one or more IP CIDRs. Traffic will be matched based on interfaces with the entered IP/CIDR address range.
* **ASN(s)**: Enter a comma-separated list of one or more ASNs. Traffic will be matched based on the origin ASN associated with the source/destination IP of the flow.
* **Custom Dimensions**: A card that expands from the **Add** button and contains a control set that is used for **Filtering with Custom Dimensions**.
* **Tenant Retention**: A drop-down menu from which you can choose how long the tenant’s data remains available for queries. Options include: Full Retention; 1, 2, or 3 weeks; and 30, 60, or 90 days.

  ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(331).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A48Z&se=2026-05-12T10%3A02%3A48Z&sr=c&sp=r&sig=KS%2Fc%2Bvp6WcwNEP7unJtgajVTK%2B74vz4yhr35ENk01Bk%3D)

  *The Advanced Filtering Criteria pane of the tenant settings pages.*

#### Filtering with Custom Dimensions

One or more custom dimensions can be applied as filters for tenants in the cards that open from the **Add** button. Each card includes the following searchable lists:

* **Dimension**: Choose a custom dimension name.
* **Populator** (available after a dimension is chosen): Choose a populator from selected dimension.

> **Notes:**
>
> * The **Add** button always appears at the end of the most recently specified card, enabling you to add another set.
> * To remove a Dimension/Populator set, click the **X** at the top right of its card.

#### Custom Dimension Considerations

The following considerations apply when using custom dimensions for MKP tenant filtering:

* Multiple dimension/populator pairs may be used simultaneously.
* The custom dimension(s) to match on must already be defined in your organization (see **Custom Dimensions**).
* The specific attributes used by a custom dimension to attribute flows to different tenants will vary depending how your network is organized.
* The value that is inserted into the custom dimension’s flow record field at ingest should allow you to uniquely identify each tenant (e.g., tenant name), so that you can tell which traffic is associated with each tenant.
* If a custom dimension is specified for MKP data source filtering (e.g., Tenant\_DST) and that same dimension is also used as a filter dimension in a dashboard panel that is included on a view that you make available to tenants, then the populator value specified in each tenant’s MKP filtering settings will override the value of the same populator in the panel filtering. This enables you to use a given panel for multiple tenants while ensuring that the traffic shown in the panel is specific to each tenant.

### Users Tab

The **Users** tab of the tenant settings pages contains the following settings and controls:

* **Add User**: A button that pops up the Add User dialog (see **Tenant User Dialogs**), where you can enter a name and email address to initiate **MKP User Activation**.
* **User list**: A table listing the users of this tenant. The table includes the following columns:

  + **Name**: The name of the user.
  + **Email**: The user’s email address.
  + **Last Login**: The date-time of the most recent login by this user to the tenant’s MKP.
  + **Remove** (trash icon): A button that pops up a confirmation dialog that enables you to remove the user.
  + **Edit** (pencil icon): A button that pops up the Edit User dialog which enables you to change the tenant’s name (see **Tenant User Dialogs**).
  > **Note**: You cannot change a tenant’s email address in the Edit User dialog.
* **Save**: A button that enables you to save the current tenant settings. If the button is not active, hover over it to pop up a message indicating which required fields are still incomplete.

#### Tenant User Dialogs

The Add User and Edit User dialogs allow you to add or edit a user on the tenant account. The dialogs include the following UI elements:

* **Name**: A field to enter the user’s name.
* **Email**: A field to enter the user’s email address.

  > **Note:** This field is not editable in the Edit User dialog.
* **Cancel**: A control — either the **X** at upper right or the **Cancel** button at bottom — that exits the dialog while leaving all settings as they were when it was opened.
* **Add User** (Add User dialog only): A button that adds the user to the tenant and exits the dialog.
* **Save**: (Edit User dialog only): A button that saves changes for the user and exits the dialog.

#### MKP User Activation

When a user is added to an MKP tenant, the account activation process involves the following steps:

1. Add the user with the **Add Users** button on the **Users** tab of a tenant settings page (see **Tenant User Settings**).
2. An invitation email will be sent to the email address that was entered for the user. Click the **Activate your account** button in the email.
3. On the resulting Account Activation page, enter (and confirm) a password.
4. When the account has been activated, click the Return to Kentik link, which goes to the login page for the MKP tenant portal.
5. Enter the credentials and log in.

> **Important**: If no tenant support email address is configured on the **MKP Settings Page**, then **support@kentik.com**will be the sender of the MKP user activation emails.

### Views Tab (Tenant)

The **Views** tab of the tenant settings pages enables you to determine what views — Dashboards and Saved Views (see **About the Library**) — in your organization will be visible to this tenant's users. You can make your organization's views available to your tenants in the following ways:

* **By package**: Choose a package (see **Tenants and Packages**) from the **Select a Package** pane at the top of the tab.
* **Individually**: Choose a view from the categorized Available list at the left of the tab.

When a view is chosen, either individually or as part of a package, it is removed from the Available list and added to the Selected Views & Dashboards list at the right of the tab. In the screenshot below, for example, the Selected Views & Dashboards list shows three views selected individually and eight views selected because they are part of a package named "Unicorn." The Available list shows (in collapsible categories) the views that are still available to be selected.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(332).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A48Z&se=2026-05-12T10%3A02%3A48Z&sr=c&sp=r&sig=KS%2Fc%2Bvp6WcwNEP7unJtgajVTK%2B74vz4yhr35ENk01Bk%3D)

*Views may be selected either individually or as part of a package.*

#### Tenant Views Tab UI

The **Views** tab contains the following settings and controls:

* **Select a Package**: A list showing the packages available for your organization, plus No Package (default). Click a package to add that package's views to the Selected Views & Dashboards list.

  > **Note:** Each tenant can be assigned only one package at a time.
* **Manage Packages**: A button that takes you to the **Packages Tab** of the My Kentik Portal page.
* **Available**: A categorized list of your organization's views (Dashboards and Saved Views) that aren't yet selected for this tenant. The categories match the categories used in the Library list in the **Library**. Click the Plus icon (+) at the right of a given view to select it. The view will be moved from the Available list to the Selected Views & Dashboards list.
* **Selected Views & Dashboards**: A categorized list of the views assigned to this tenant, either individually or via a package. Click the Minus icon (-) at the right of a given view to unselect it; it will return to the Available list.
* **Default Landing Page**: A pane with radio buttons that determine the landing page that the users of this tenant will see when they first log in:

  + **Explore Your Network**: A page whose **Library** tab includes **Dashboards** and **Saved Views** that you specify on the **Views** tab of the tenant's settings page (see **Tenant View Settings**).
  + **Selected View**: Expands the pane to reveal a filterable drop-down listing the views in Selected Views & Dashboards (see above). To choose a landing page, click an item in the list.
* **Save**: A button that enables you to save the current tenant settings. If the button is not active, hover over it to pop up a message indicating which required fields are still incomplete.

> **Note:**To see an individual view in the Available list or the Selected Views & Dashboards list as it would appear to a given MKP tenant, click the name of that view to open it in a new tab (see **Preview as Tenant**).

### Policies Tab (Tenant)

The **Policies** tab of the tenant settings pages enables you to do the following:

* Configure alert policies that will notify you (the landlord) and/or the tenant's users of anomalous conditions affecting their network traffic. These tenant policies are actually "subpolicies" of your organization's **Alert Policies** because they are dependent on the settings of those policies.
* Choose a mitigation (if available) to activate when a policy’s thresholds are met.

To ensure that a subpolicy evaluates only data from the individual tenant's part of your overall dataset, the dimensions for a subpolicy must include all of the data source and filtering dimensions set on the **Data Sources** tab (see **Tenant Data Sources**). For example, if you populate the **ASN** field in the **Advanced Filtering** Criteria pane, ASN must be included as a dimension in each of the tenant's subpolicies.

> **Note:**Tenant-specific subpolicies are in addition to any policies that have been assigned to a tenant via a package (see **Package Policy Settings**).

The **Policies** tab includes the following UI elements:

* **Required dimensions**: A card listing the dimensions, set on the **Data Sources** tab, that must be used in this subpolicy.

  > **Note**: This card is not visible until dimensions have been saved for this tenant on the **Data Sources** tab.
* **Policy list**: A set of vertically arrayed cards, one for each of this tenant's subpolicies (see **Alert Policy Card**.).

  > **Note:** This list is not visible until it contains at least one policy card.
* **Add Alert Policy**: A button at the end of the Policy list that adds a new Policy card.
* **Save**: A button that saves any changes made to the tenant alert policies.

#### Alert Policy Card

When a policy card is first opened from the **Add Alert Policy** button it contains only the following elements:

* **Select a policy**: A drop-down that is populated with all alert policies currently existing in your organization (the same policies listed on the Settings » **Alert Policies** page). Click the drop-down to open the filterable list, then click on a policy to select it.

  > **Note**: Policies that are tagged as "Incompatible" cannot be added to the tenant. If none of the listed policies are compatible, modify the dimensions in the Data Sources tab.
* **Cancel**: A button that closes the card and removes the corresponding subpolicy from the tenant.

Once a policy has been selected from the Select a policy drop-down, the policy card will be expanded and populated with many of the same UI elements that appear on the Thresholds tab of the policy's Edit Policy page (see **Policy Threshold Settings**), including:

* **Error status**: An indicator of the policy's status, which can be one of the following:

  + **Valid** (green checkmark): No issues.
  + **Error** (red stop-sign): Hover on the indicator for a popup stating the issue(s).
* **Severity selector**: A set of tabs (in the left sidebar) containing settings for each threshold of the policy: Critical, Severe, Major, Warning, or Minor.
* **Conditions**: A pane whose settings determine what constitutes a match for the threshold (see **Threshold Conditions**). On each threshold tab there is a list of one or more **Conditions,** with operator and value settings for each condition (e.g., "> 33 Mpackets/s"). These settings default to the settings defined in the **Conditions** pane of the **Thresholds** tab of the parent policy. While you can customize the values of the conditions in a subpolicy, the types of those conditions can only be changed by editing the parent policy.
* **Frequency**: A pane whose settings determine how many matches must occur within a given duration of time in order to trigger an alert (see **Threshold Frequency**).
* **Actions**: A pane whose settings determine how the alerting system responds to an alert generated by this threshold (see **Threshold Actions**).

  > **Notes:**
  >
  > + Parent policy notification channels are not carried over to a subpolicy.
  > + To change tenant notification settings, a Kentik user (landlord) must spoof the tenant in the MKP tenant portal (see **Tenant Notifications Access**).
  > + A channel created on the **Tenant Notifications Page** of the MKP Tenant Portal by a tenant user (actual or spoofed) will appear in the **Notification** **Channels** field of any of the tenant's subpolicies that exist when the channel is created.
  > + Mitigations are only available for subpolicies whose dimensions include source or destination IP/CIDR.
* **Threshold Configuration** (under **Advanced Options**): A pane whose settings vary depending on the **Conditions** settings (see **Threshold Configuration**).
* **Policy Summary**: A pane (on the right) showing a summary of the parent policy’s dataset. While the dataset cannot be modified in the subpolicy, you can click **Configure** **Policy** to modify the dataset in the parent policy.
* **Configure Policy**: A direct link back to the Edit Policy page for the parent policy (see **Policy Settings Pages**).

  > **Note**: Be sure to save the subpolicy before clicking the **Configure Policy** button.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(336).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A48Z&se=2026-05-12T10%3A02%3A48Z&sr=c&sp=r&sig=KS%2Fc%2Bvp6WcwNEP7unJtgajVTK%2B74vz4yhr35ENk01Bk%3D)

*The Policy list with two cards, one saved and collapsed, the other not yet saved.*

#### Saved Alert Policy Card

When a tenant with one or more subpolicies has been saved the following changes are made to the policy card(s) for those subpolicies:

* The **Select a Policy** drop-down is replaced with a **Subpolicy** **summary** at the top of the policy card.
* The **Cancel** button is replaced with a **Remove** button (trash icon) that enables you to remove the subpolicy from this tenant’s alert policies list.

The**Subpolicy summary** includes the following elements:

* **Expand Policy**: A button (single angle quotation mark) that determines which elements — in addition to the **Subpolicy summary** itself — are visible in the policy card:

  + **Expanded**: All of the elements described in **Alert Policy Card** are visible.
  + **Collapsed**: Only the **Error status** and **Remove** elements are visible.
* **Name**: The name given to the policy in the **General** tab of the policy itself (see **General Policy Settings**).
* **ID**: The Kentik-assigned number for the policy.
* **Subpolicy ID**: The Kentik-assigned number for the subpolicy.
* **Description**: The description assigned in the **General** tab of the parent policy (see **General Policy Settings**).

### Preview as Tenant

The **Preview as Tenant** feature enables you to see an individual view (Dashboard or Saved View) as it would appear to a MKP tenant, where only the traffic data that has been configured for that tenant is shown.

You can access a tenant preview in the following ways:

* **Tenant Settings Page**: On the **Views** tab, click a view in the Available list or the Selected Views & Dashboards list to open the tenant-filtered view in a new tab (see **Tenant View Settings**).
* **Data Explorer:** From any query, choose **Preview as Tenant** from the Actions menu (see **Data Explorer Actions**), then choose the tenant for the preview.
* **Metrics Explorer**: From any query, choose **Preview as Tenant** from the Actions menu (see **Metrics Explorer Actions**), then choose the tenant for the preview.
* **Library**: On a Dashboard or Saved View, choose **Preview as Tenant** from the Actions menu, then choose the tenant for the preview (see **Dashboard Subnav Controls**).

> **Notes:**
>
> * When using “Preview as Tenant”, you can still access your own menu options on the portal’s main navbar.
> * To simulate a tenant's experience in the MKP tenant portal (see **MKP Tenant Portal**), in the **Tenants List**, open the Action menu for that tenant and select **View as Tenant**.

## Package Settings Pages

The Add Package and Edit Package pages, referred to collectively as the "package settings pages," both contain the same structure, layout, and settings. These settings are specified in **General**,**Views**, and **Policies** tabs accessed via the left sidebar.

Package settings are used to configure and manage packages, which are named collections of views (Dashboards and Saved Views) and Alert Policies (which define traffic conditions in which tenant users may receive notifications).

Package settings can be accessed in the following ways:

* **Add Package**: Click **Add Package** on the **Packages Tab** of the My Kentik Portal page to open the Add Package page, where you configure the new package.
* **Edit Package**: Click **Edit** next to a package on the My Kentik Portal page to open the Edit Package page, where you can modify its settings.

  > **Note:**Changes to package settings are applied to all tenants assigned to that package. Kentik recommends notifying tenants of any significant changes.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(339).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A48Z&se=2026-05-12T10%3A02%3A48Z&sr=c&sp=r&sig=KS%2Fc%2Bvp6WcwNEP7unJtgajVTK%2B74vz4yhr35ENk01Bk%3D)

*Defining a package includes choosing an icon and color.*

### General Tab (Package)

The **General** tab of a package settings page contains the following settings and controls:

* **Default**: A switch that, if enabled, will result in this package being assigned automatically to new tenants.
* **Name** (required): The name by which this package will be listed on the **Packages Tab**.
* **Description**: A description of the package.
* **Icon**: A drop-down containing icons that you can use to represent this package.
* **Color**: A drop-down from which you can choose a color to associate with this package, either by clicking a preset swatch or entering a hex color code.
* **Preview**: A rendering of the icon/color combination resulting from the current settings.
* **Save** (Edit Package only): A button that enables you to save the current package settings. If the button is not active, hover over it to pop up a message indicating which required fields are still incomplete.
* **Cancel** (Add Package only): A button that returns you to the **Packages** tab on the My Kentik Portal page without adding the package.
* **Next** (Add Package only): A button that saves changes to the **General** tab and takes you to the **Views** tab.

### Views Tab (Package)

The **Views** tab of the package settings pages enables you to determine what views — Dashboards and/or Saved Views (see **About the Library**) — in your organization will be visible to the users of the tenants to which this package is assigned. You make a given view available in a package by choosing it from the categorized Available list at the left of the tab, which shows the views that are available to be selected. Once selected, the view is removed from the Available list and added to the Views in Package list at the right of the tab.

#### Package Views Tab UI

The **Views** tab contains the following settings and controls:

* **Available**: A categorized list of your organization's views (Dashboards and Saved Views) that aren't yet selected for this package. The categories match the categories used in the Views list in the **Library**. Click the Plus icon (**+**) at the right of a given view to select it. The view will be moved from the Available list to the Views in Package list.
* **Views in Package**: A categorized list of the views that your organization currently has assigned to this package. Click the Minus icon (-) at the right of a given view to unselect it; it will return to the Available list.
* **Save** (Edit Package only): A button that enables you to save the current package settings. If the button is not active, hover over it to pop up a message indicating which required fields are still incomplete.
* **Next** (Add Package only): A button that saves changes to the **Views** tab (creating the package) and takes you to the **Policies** tab.

### Policies Tab (Package)

The **Policies** tab of the package settings pages enables you to configure alerts that will notify about anomalous conditions affecting network traffic. When a threshold is triggered in a package-level alert policy, the specified notifications will go to the MKP users of all tenants to which the policy is assigned, whether directly or via a package. The settings and controls for package-level alert policies are identical to those for tenant-level policies, which are covered in **Tenant Policy Settings**.

> **Note**: Policies assigned to a package can only be edited from within the package. The Policies tab on the tenant settings pages displays the name, ID, and subpolicy ID of an applied package’s policy, as well as an **Edit** **Package** button which takes you to that package’s settings page.

## Configuring Tenancy

Follow these steps to configure tenancy for your organization’s My Kentik Portal:

### Configure Tenant Portal

To configure a tenant portal:

1. On the My Kentik Portal page, click **Settings** at upper right to open the **MKP Settings Page**.
2. On the **Domain** tab, set the subdomain for your organization’s tenant portal (see **MKP Domain Settings**).
3. On the **Support Links** tab, set the email address and URL that tenant users should contact when they need support (see **MKP Support Settings**).
4. On the **Branding** tab, set the logo that will appear on the tenant portal (see **MKP Branding Settings**).
5. To enable SSO, fill in the SSO settings on the **Tenant SSO** tab (see **MKP Tenant SSO Settings**).
6. Click **Save** to save the tenant portal configuration.

### Add a Tenant

To add a tenant:

1. On the **Tenants Tab** of the My Kentik Portal page, click **Add Tenant** (see **Tenant Settings Pages**).
2. On the **General** tab, provide a name and description for the tenant (see **Tenant General Settings**).
3. On the **Data Sources** tab, specify the parts of your network about which this tenant will be able to see network data (see **Tenant Data Sources**).
4. On the **Users** tab, specify the users that will have access to this tenant's MKP (see **Tenant User Settings**).
5. On the **Views** tab (see **Tenant View Settings**):

   1. Choose which package, if any, to assign to this tenant.
   2. Choose any additional views (Dashboards or Saved Views) to make available to the users of this tenant.
6. On the **Policies** tab, specify the alert policies for users of this tenant (see **Add a Tenant Policy** and **Tenant Policy Settings**).
7. Click **Add Tenant** to save the settings and exit. The new tenant will appear in the **Tenants List**.

### Edit a Tenant

To edit a tenant:

1. On the **Tenants Tab** of the My Kentik Portal page, find the tenant you wish to edit, using the **Filter** field as necessary.
2. Click the tenant’s name in the **Name** column.
3. Modify the settings as described in **Add a Tenant**.
4. Click **Save** to save the settings and exit.

### Add a Tenant Policy

To add a subpolicy for a tenant:

1. On the **Tenants Tab** of the My Kentik Portal page, find the tenant to which to add an alert policy, using the **Search** field as necessary.
2. Click the tenant name (under the **Name** column) or click the Actions menu at the right of the tenant’s row and select **Edit Tenant**.
3. On the **Policies** tab, click the **Add Alert Policy** button at the bottom of the **Policy** list (see **Tenant Policy Settings**).
4. Click the **Select a policy** drop-down to show a filterable drop-down list of policies.
5. Click a listed policy to select it and open the threshold settings for the new subpolicy (the instance of the policy that is specific to this tenant).
6. Modify the subpolicy’s thresholds as needed (see **Alert Policy Card**).

   > **Note**: Changes made to the tenant subpolicy do not affect the parent policy.
7. Click **Save** to add the new subpolicy to the tenant.

### Add a Package

To add a package:

1. On the **Packages Tab** of the My Kentik Portal page, click **Add** **Package** (see **Package Settings Pages**).
2. Click the **General** tab and provide a name and description for the package. If desired specify a new color and/or shape for the package icon (see **Package General Settings**).
3. Click the **Views** tab and choose the views (Dashboards or Saved Views) to make available to the users of this package (see **Package View Settings**).
4. Click the **Policies** tab and specify the alert policies for users of this package (see **Package Policy Settings**).
5. Click **Add Package** to save the settings and exit. The new package will appear in the **Packages List**.

### Edit a Package

To edit a package:

1. On the **Packages Tab** of the My Kentik Portal page, find the package you wish to edit, using the **Filter** field as necessary.
2. Click **Edit** to the right of the package's tile to open the Edit Package page.
3. Modify the settings tabs as described in steps 2 to 4 of **Add a Package**.
4. Click **Save** to save the settings and exit.
