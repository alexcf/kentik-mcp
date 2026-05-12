# Authentication & SSO

Source: https://kb.kentik.com/docs/authentication-sso

---

This article covers the **Authentication & SSO** page in the Kentik portal.

> **Important:** You must be a Super Admin user (see **About Super Admin Users**) to administer Authentication and SSO. If your organization has no Super Admin users, contact **support@kentik.com** for assistance.

![Settings page for Authentication and SSO with Single Sign-On options and service provider details.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/SSO-main(2).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A13Z&se=2026-05-12T09%3A44%3A13Z&sr=c&sp=r&sig=rf4ijsjr7pKUfJNvratPY2CiST%2Bs%2B9bn3HSNQMOlzTg%3D)

*The settings page for configuration of password expiration and SSO.*

## About Authentication & SSO

![Organization settings menu highlighting Authentication and SSO option for user management.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/SSO-menu-path.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A13Z&se=2026-05-12T09%3A44%3A13Z&sr=c&sp=r&sig=rf4ijsjr7pKUfJNvratPY2CiST%2Bs%2B9bn3HSNQMOlzTg%3D)Accessible via Organization Settings » **Authentication & SSO** from the Kentik portal’s main nav bar, this page is where you manage your **Password Policy** and **Single Sign-On** (SSO) settings. Both are ways to enhance your Kentik account’s security.

Single sign-on (SSO) lets Kentik users access the portal with the same credentials they use for other SSO-enabled applications, providing seamless access with one login.  
  
Kentik’s SSO implementation uses the standard **SAML2** protocol, also known as “Federated Identity Management.” In this terminology:

* Kentik is the **service provider** (SP).
* Your SSO is the **identity provider** (IDP).

### SSO Identity Providers

Kentik’s SSO implementation has been successfully tested with the following identity providers (IDPs):

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/SSO-Logos-121h672w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A13Z&se=2026-05-12T09%3A44%3A13Z&sr=c&sp=r&sig=rf4ijsjr7pKUfJNvratPY2CiST%2Bs%2B9bn3HSNQMOlzTg%3D)

### SSO Process Overview

These steps and diagram describe the SSO process with Kentik:

1. Your user attempts to log into the Kentik portal.
2. Kentik redirects them to your organization's IDP for verification.
3. The IDP authenticates the user and sends a SAML2 response to Kentik.
4. If the response confirms a successful authentication, Kentik logs the user in to the Kentik portal with certain permissions (see **User Level Determination Flow**), otherwise, access to the Kentik portal is denied.

![SSO Authentication Process](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/SSO-Overview_diagram.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A13Z&se=2026-05-12T09%3A44%3A13Z&sr=c&sp=r&sig=rf4ijsjr7pKUfJNvratPY2CiST%2Bs%2B9bn3HSNQMOlzTg%3D)

*SSO Authentication Process*

> **Notes**:
>
> * To maintain security, user sessions for **Admin** and **Super Admin** accounts automatically expire after 8 hours of inactivity (excludes passive actions like auto-refresh).
> * For **Member** accounts, sessions do not expire due to inactivity. This allows for continuous use in scenarios such as Network Operations Center (NOC) dashboards, where a user might need to view real-time data without manual interaction.

## SSO Prerequisites

To configure your Kentik account for SSO, you must meet the following requirements:

1. **Identity Provider Account**: You must have one of the following:

   * **Existing IDP Account**: An account with an existing SAML2 identity provider (e.g., Okta, OneLogin, Ping Federate, Google G Suite, Duo) and a directory of users.
   * **In-House Identity Management**: A self-operated SAML2-compatible IDP or identity gateway, such as Shibboleth.

     > **Note:** For users who need LDAP or Active Directory as their authentication backend, Kentik recommends Shibboleth, which has open-source LDAP and AD extensions.
2. **Kentik Super Admin User**: At least one Kentik portal user in your organization must be a Super Admin (see **About Super Admin Users**).

## Kentik SSO Configuration

Configuration of single sign-on (SSO) in Kentik is covered in the following topics.

> **Important:** You must be a Super Admin user to configure SSO in Kentik (see **About Super Admin Users**).

### SSO Types

You can manage SSO for two distinct groups of Kentik portal users, as follows:

| SSO Type | User Group | Managed Via |
| --- | --- | --- |
| **Direct-User SSO** | Admin and Member users (see **About Users**) that sign directly into your organization | Kentik portal’s **Authentication & SSO** page |
| **Tenant SSO** | Users of one of your organization’s My Kentik Portal tenants (see **About Tenancy**) | My Kentik Portal Single Sign-On page:   * Click **Settings** on the **My Kentik Portal Page** (Main Menu » **My Kentik Portal**). * Select **Tenant SSO** from the left menu. |

> **Tip**: The Tenant SSO settings differ from the Direct-User settings as follows:
>
> * The **Profile Attribute for Tenant ID** field is present for Tenant SSO only
> * The **SSO Attribute Mapping** tab’s settings are not available for Tenant SSO

### SSO Configuration Overview

![Settings page for Authentication and SSO with options for Single Sign-On configuration.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/SSO-enable-sso.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A13Z&se=2026-05-12T09%3A44%3A13Z&sr=c&sp=r&sig=rf4ijsjr7pKUfJNvratPY2CiST%2Bs%2B9bn3HSNQMOlzTg%3D)Here’s an overview of the steps you’ll perform to setup Kentik SSO:

* In the Kentik portal (Authentication & SSO page, SSO Settings tab), the **SSO Disabled** (default) option is selected from the dropdown.
* Complete all steps in **Add Kentik to Your IDP** and **Configure IDP Settings in Kentik**.
* To begin using SSO: In the Kentik portal, select **SSO Required** from the dropdown and click **Save** to apply the changes.

> **Tips**:
>
> * For more on the migration process, see **Migrating to SSO**.
> * Selecting **SSO Disabled** can be done at any time without losing any settings you’ve made.

### Add Kentik to Your IDP

SSO involves two-way communication between Kentik and the IDP, so you’ll need to configure your IDP to recognize Kentik as a service provider (SP):

1. On the **SSO Settings** tab, copy the values of the following:

* ![Kentik Service Provider configuration details including Entity ID and ACS URL.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/SSO-SP-config-details(1).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A13Z&se=2026-05-12T09%3A44%3A13Z&sr=c&sp=r&sig=rf4ijsjr7pKUfJNvratPY2CiST%2Bs%2B9bn3HSNQMOlzTg%3D)**Service Provider (SP) Entity ID**: A unique identifier for Kentik.
* **Service Provider (SP) Assertion Consumer Service (ACS) URL**: Kentik’s Assertion Consumer Service URL, to which the IDP should post its responses.

2. Paste the values into the appropriate form in your IDP.

> **Notes:**
>
> * Some IDP solutions, including Shibboleth, can read the SP Entity ID and SP ACS URL from an XML configuration file. Download this file from the Authentication & SSO page via the **Download Kentik SP metadata** button.
> * For Tenant SSO, depending on the IDP you might also need the IDs listed in the **Tenant** **ID Name** column on the My Kentik Portal page.

### Configure IDP Settings in Kentik

Once you’ve added Kentik to your IDP, follow these steps to set IDP-related settings in the Kentik portal (e.g., setting the attribute names the IDP should use in responses to Kentik’s authentication requests) or configuring encrypted assertions.

1. Go to the Authentication & SSO page and select the **SSO Settings** tab.
2. Make the following settings:

* ![Configuration settings for Kentik Identity Provider including SSO URL and encryption options.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/SSO-IDP-config.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A13Z&se=2026-05-12T09%3A44%3A13Z&sr=c&sp=r&sig=rf4ijsjr7pKUfJNvratPY2CiST%2Bs%2B9bn3HSNQMOlzTg%3D)**Identity Provider (IDP) SSO Entry Point URL** (required): Enter the IDP entry-point URL to which Kentik should redirect the browser when a user initially attempts to log in (e.g., “https://sso.jumpcloud.com/saml2/kentikportalprod”).
* **Profile Attribute for User Email** (required): Enter the attribute name (e.g., “nameID”) in which the IDP should return the IDP’s email attribute key, enabling Kentik to receive the user’s email in the IDP’s response.
* **Profile Attribute for Tenant ID** (required, tenant SSO only): The field name in which the IDP should return the Kentik tenant ID.

  > **Note:** You can find the Tenant ID in the **Tenant Name** column on the My Kentik Portal page.
* **IDP requires encrypted assertions** (default = Off): Switch to On to encrypt SAML assertions (authentication, attribute, and/or authorization decision statements) in IDP responses.
* **Sign outgoing authentication requests** (default = Off):Switch to On if you want all outgoing authentication requests (AuthnRequests) from Kentik to your IDP to be signed with Kentik’s certificate.

  > **Notes**:
  >
  > + SAML2 SSO usually requires this to be option to be Off.
  > + If this setting is On, you must update your IDP's configuration annually. Kentik rotates the certificate used for signing requests on **October 15th** of every year (see **Signed Certificate Rotation Process**).
* **Omit requested** `authn` **context from outgoing authentication requests** (default = Off):

  + **If On**: Outgoing authentication requests do not contain “authnContext”.
  + **If Off**: Outgoing requests contain “authnContext”.
* **IDP Public Signing Key** (required): Paste in the IDP’s public signing key.

  > **Notes**:
  >
  > + Kentik recommends preparing the IDP Public Signing Key with this **web-based tool****.**
  > + The IDP Public Signing Key is **required** for SSO configuration as of January 27th, 2025.

### Configure User-level Behavior

![Settings page for Authentication and SSO with User Level Mapping details displayed.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/SSO-attribute-mapping-user-level.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A13Z&se=2026-05-12T09%3A44%3A13Z&sr=c&sp=r&sig=rf4ijsjr7pKUfJNvratPY2CiST%2Bs%2B9bn3HSNQMOlzTg%3D)For Direct-User SSO only (see **SSO Types**), enable your Kentik portal user levels to be managed from the IDP with the **User Level** and Role-Based Access Control (**RBAC**) permission models.

#### **User Level vs. RBAC Roles**

While the Kentik portal moves to a full RBAC model, both of these permission models coexist as follows:

* **User Level****:** Governs legacy Kentik portal areas using the basic permissions (Member, Admin, SuperAdmin).
* **RBAC Roles and Role Sets****:** Governs modern Kentik portal areas using granular RBAC permissions.

> **Tip**: For further explanation, see **User Level Determination Flow**and **Default Fallback Outcome Table**.

#### Configure User Level Mapping

1. Go to the Authentication & SSO page and select the **SSO Attribute Mapping** tab.
2. In **Profile attribute key for User Level**, optionally enter the field name you want Kentik to inspect (in SAML2 responses from your IDP) to determine the user’s User Level.

   1. Kentik accepts any string in the User Level attribute.
   2. If the IDP provides a valid level, Kentik overrides its own internal user-level value with the IDP’s value at login.

      > **Example**: If a Kentik Admin is identified instead as a Member by the IDP, Kentik resets their level to Member, removing Admin privileges.
   3. If the level is blank, Kentik uses its internal user-level value.
3. If you specified a **Profile attribute key for User Level**, optionally add entries to the **Mapping Table for User Level**.

   1. Click **Add Mapping** to add a new entry and specify its details as follows:

      1. **Attribute Value**: The value that Kentik, when encountering it in your IDP’s SAML2 responses, will map to the value specified in Mapped User Level.
      2. **Mapped User Level**: Choose from the dropdown the Kentik user role to which this Attribute Value should be mapped.
4. Set the fallback behavior when 1) no mapping is found for the entered **Profile attribute key for User Level** value or 2) the key is missing in the SAML2 response.

   1. Choose a radio button, either **Portal-Local User Attribute**, **Member**, or **Deny Access** (see **SSO Attribute Mapping Tab**) to set the default fallback for the User Level mapping.

**A****dditional Considerations:**

* If the IDP provides a valid user-level profile attribute, any user level changed directly in Kentik (via portal or API) will be reset to the IDP-provided level at the next SSO login.
* If a Super Admin is in an IDP group whose user level is collectively set via an IDP user-level profile attribute, that user will lose Super Admin privileges.
* If your IDP sends user-level values considered as “invalid” by Kentik (e.g., true and false rather than 1 and 0), they may be able to map them to Kentik-valid values in the SAML assertions sent to Kentik.

> **Tip**: See **User Level Determination Flow** for further explanation about Kentik SSO and User Levels.

#### Configure RBAC Role Mapping

1. Go to the Authentication & SSO page and select the **SSO Attribute Mapping** tab.
2. In **Profile attribute key for Role Set**, optionally enter the field name you want Kentik to inspect (in SAML2 responses from your IDP) to determine the RBAC role set.

   1. Kentik accepts any string in the Role Set attribute.
3. If you specified a **Profile attribute key for Role Set**, optionally add entries to the **Mapping Table for Role Set**.

   1. Click **Add Mapping** to add a new entry and specify its details as follows:

      1. **Attribute Value**: The value that Kentik, when encountering it in your IDP’s SAML2 responses, will map to the value specified in Mapped User Level.
      2. **Mapped Role Set**: Choose from the dropdown the Kentik RBAC role set to which this Attribute Value should be mapped.
4. Set the fallback behavior when either 1) no mapping is found for the entered **Profile attribute key for Role Set** value or 2) the key is missing in the SAML2 response.

   1. Choose a radio button, either **Portal-Local User Attribute** or **Assign this default Role Set** (see **SSO Attribute Mapping Tab**) to set the default fallback for the RBAC Role mapping.

> **Note**: The RBAC Role mapping is distinct from the **User Level** mapping (see **User Level vs. RBAC Roles**).

### Configure Additional Options

Use these additional settings to customize the configuration to your organization’s specific needs:

1. Go to the **Authentication & SSO** page and select the **SSO Settings** tab.
2. Make the following settings (as needed):

* **Disable 2FA when user has authenticated via SSO** (default = Off):

  + **If On**: Two-factor authentication (2FA) is disabled when SSO is enabled.
  + **If Off**: 2FA users who log in via SSO must input a one-time password from their 2FA source (e.g., Google Authenticator or any OTP client).
* **Allow auto-provisioning of new users** (default = Off) determines the action for Kentik to take when a user not already registered with Kentik attempts to log in (see **Auto-Provisioning Considerations**):

  + **If On**: Kentik allows the login and the user is automatically registered with Kentik.
  + **If Off**: Kentik denies the login.

#### **Auto-Provisioning Considerations**

When deciding to allow auto-provisioning of your organization’s SSO users, consider the following:

* If the **Profile Attribute for User** **Level** field is blank or a valid user-level key is not found in the IDP response,  auto-provisioned users will be assigned a user level of “Member”.
* Kentik can only set the Full Name of auto-provisioned users to the IDP-provided email address (see **Profile** **Attribute** **for User Email** in **Configure IDP Settings in Kentik**).
* Kentik cannot automatically “de-provision” a user, meaning that removing a user from the IDP does not also remove them from Kentik.

  > **Note**: You can update a user’s Full Name or remove them at any time via Settings » **Users** in the Kentik portal, or via Kentik’s **User APIs**.

### Enable SSO

The final step in configuring Kentik SSO is in the Kentik portal, as follows:

1. Go to the **Authentication & SSO** page and select the **SSO Settings** tab.
2. In the dropdown at the top, choose **SSO Required** (see **SSO Settings Tab**) and click **Save**.

   > **Tip**: For more on the migration process, see **Migrating to SSO**.

## SSO Login Process

Once SSO is enabled, your users log in with a newly-created URL that is specific to your organization, for example:

`https://portal.kentik.com/login/sso/company_shortname`

To construct your SSO login URL, replace `company_shortname` in the above example with the value found in last segment of either the Service Provider (SP) Entity ID or Service Provider (SP) Assertion Consumer Service (ACS) URL fields (see **Add Kentik to Your IDP**).

When users land on your Kentik SSO login page, what happens depends if the they have a valid active session as defined by the IDP:

* Valid active session exists: User is automatically logged into the Kentik portal.
* No active session exists: User is redirected to their IDP’s login screen and then back to Kentik Portal upon successful authentication.

## Migrating to SSO

The following procedure is the recommended way to transition from plain authentication to SSO:

1. Configure SSO while the dropdown is set to **SSO Disabled** (default). Perform all of the needed tests and validate the optional features with a single user, typically a Super Admin who is also in charge of SSO for your organization.
2. Send an announcement to your organization’s Kentik user base with the following:

   1. The cutoff date on which Kentik access will become SSO-only.
   2. Contact info for the Super Admin, so that users can ask for help before the cutoff date.
   3. The new login URL for Kentik access via SSO (see **SSO Login Process**).
3. On the cutoff date, select **SSO Required** from the dropdown, after which Kentik users will only be able to log in using the SSO URL.

   > **Important:** Once SSO is required, users attempting non-SSO access (https://portal.kentik.com/login) will be denied access.

## About Super Admin Users

Super Admin users have the same capabilities as Admin users, with the following additional privileges:

* Can configure SSO from the portal’s Authentication & SSO page (Settings » **Authentication & SSO**).
* When SSO is required, Super Admins are the only ones who can still use non-SSO login (e.g., username + password, or username + password + 2FA), allowing a Super Admin to disable SSO in case of an identity provider failure.
* Can turn other users into Super Admin users.

> **Tip:** To prevent a single point of failure, Kentik recommends you set up two Super Admins. More than two is not recommended, however, in favor of restricting how many users can log in with the traditional username/password approach.

Admin-level users can check who the Super Admin users as follows:  

* Check the **Level** column in the User List (Settings » **Users**).
* If no user is assigned a Super Admin role, contact Kentik support (**support@kentik.com**) to request that a Super Admin be designated for your organization.

> **Note:** If your organization signed up with Kentik prior to October 2017, the first user registered to your account will be automatically set as a Super Admin (to change, go to Settings » **Users**).

## Signed Certificate Rotation Process

To maintain security, Kentik rotates the certificate used to sign outgoing authentication requests on **October 15th** of every year. If you have the **Sign outgoing authentication requests** switch enabled, you must update your IDP to trust the new certificate to avoid an interruption of your SSO service.

The annual rotation process:

* **Download New SP Metadata**: On or after October 15th, click the **Download Kentik SP metadata** button on the Authentication & SSO page to download the updated XML file. This file contains the new signing certificate.
* **Update Your IDP**: You must then configure your IDP with the new Kentik SP metadata. While specific steps may vary, the general process involves:

  + Extracting the new x509 certificate from the downloaded XML file. The certificate value is the text found between the `<ds:X509Certificate>` tags.
  + Installing or uploading this new certificate into your IDP's configuration for the Kentik application.
* **Verify Access**: After updating your IDP, users will be able to log in via SSO again, as your IDP will now trust the signed requests from Kentik.

If you require assistance with this process, contact Kentik support at **support@kentik.com**.

## Authentication & SSO UI

The **Authentication & SSO** page has three tabs for managing settings:

* **Password Policy Tab**: Configure security requirements/constraints for user passwords when logging in directly via email/password.
* **SSO Settings Tab**: Enable SSO and set up communication details (endpoints, certificate, encryption) between Kentik and the IDP (see **About SSO**).
* **SSO Attribute Mapping Tab**: Define rules for translating user attributes sent from the IDP (e.g., Role Set or User Level) into the corresponding Kentik portal permissions.

> **Tip**: Click the page’s always-visible **Save** button to save all pending changes across all tabs.

### Password Policy Tab

The **Password Policy** tab contains the following UI elements: ![Settings page showing password expiration policy options for authentication and SSO.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/SSO-password-policy-settings.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A13Z&se=2026-05-12T09%3A44%3A13Z&sr=c&sp=r&sig=rf4ijsjr7pKUfJNvratPY2CiST%2Bs%2B9bn3HSNQMOlzTg%3D)

* **Password Expiration Duration**: A dropdown to set the duration before a password must be reset.
* **Disable Password Expiration**:

  + **If** **Off** (default): Password-only users must reset their password according to the configured expiration duration.
  + **If** **On**: Users won’t be required to reset their password.

#### Password Policy Considerations

* Kentik’s password policy adheres to SOC-2 compliance by enforcing specific rules.
* Kentik checks password complexity for all new passwords, whether created for new accounts or during password resets.
* Passwords can be set to expire after a specific number of days to minimize the risk of unauthorized access.
* Kentik prevents password reuse by keeping track of the five most recently used passwords.

### SSO Settings Tab

The **SSO Settings** tab has the following UI elements:

* **Single Sign-On Settings** (pane):

  + **SSO Status** (dropdown): Choose from the following states of Single Sign-On for your organization:

    - **SSO Enabled:** Users can choose either SSO or standard login.
    - **SSO Disabled**: All users must use standard login.
    - **SSO Required**: All users (except Super Admins) must log in via SSO.
  + **Disable 2FA when user has authenticated via SSO** (toggle): When On, bypasses two-factor authentication for users authenticated via the IDP.
  + **Allow auto-provisioning of new users** (toggle): When On, automatically creates new user accounts upon first successful SSO login.
* **Kentik Service Provider (SP) Config details** (pane):

  + **Service Provider (SP) Entity ID** (button): Click to copy Kentik's unique identifier for the IDP configuration.
  + **Service Provider (SP) Assertion Consumer Service (ACS) URL** (button): Click to copythe endpoint where the IDP sends the SAML assertion.
  + **Download Kentik SP metadata** (button): Click to download the Kentik SP metadata file for automated IDP setup.
* **Identity Provider (IDP) Configuration** (pane):

  + **Identity Provider (IDP) SSO Entry Point URL**: Enter the URL provided by the IdP where Kentik initiates the SAML login request.
  + **Profile attribute for User Email**: Enter the SAML attribute name (e.g., nameID) that carries the user's email address.
  + **IDP requires encrypted assertions** (toggle): When On, Kentik will expect the IDP to encrypt the SAML assertion before sending it.
  + **Sign outgoing authentication requests** (toggle): When On, Kentik will digitally sign the SAML authentication request it sends to the IDP, allowing them to verify that the request originated from Kentik.
  + **Omit requested authn context from outgoing authentication requests** (toggle): When On, Kentik will exclude the standard `RequestedAuthnContext` element from the SAML authentication request (sometimes required for compatibility with certain IDPs).
  + **IDP Public Signing Key (X.509 cert)**: Enter the IDP's X.509 public signing certificate, used by Kentik to verify the assertion signature.

![SSO Settings tab of the Authentication and SSO page with various configuration options displayed.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/SSO-sso-settings-tab-main(1).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A13Z&se=2026-05-12T09%3A44%3A13Z&sr=c&sp=r&sig=rf4ijsjr7pKUfJNvratPY2CiST%2Bs%2B9bn3HSNQMOlzTg%3D)

*The SSO Settings tab of the Authentication and SSO page.*

### SSO Attribute Mapping Tab

The **SSO Attribute Mapping** tab defines how SAML attributes passed during login are translated into Kentik user permissions using the User Level and RBAC Roles models (see **User Level vs. RBAC Roles**).

The tab has the following UI elements:

* **User Level Mapping** (pane):

  + **Profile attribute key for User Level:** Enter the SAML attribute key for Kentik to determine its User Level assignment.
  + **Mapping Table for User Level:** Define a custom user level mapping table:

    - **+ Add Mapping** (button): Click to add a table row. Only active once a profile attribute key is defined (see above).
    - **Attribute Value**: Enter the raw SAML attribute string for the mapping table entry.
    - **Mapped User Level** (dropdown): Choose the Kentik user level for the mapping table entry.
    - **Remove** (trash icon): Click to delete the mapping table entry.

    ![User level mapping configuration for the Kentik Portal with attribute key and values.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/SSO-mapping-table-for-user-level(2).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A13Z&se=2026-05-12T09%3A44%3A13Z&sr=c&sp=r&sig=rf4ijsjr7pKUfJNvratPY2CiST%2Bs%2B9bn3HSNQMOlzTg%3D)

    *Define a custom User Level mapping table.*
  + **Default Fallback** (radio buttons)**:** Choose from the following defaults when the User Level attribute is missing or unmapped (see **Default Fallback Outcome Table**):

    - **Portal-Local User Attribute**: Retain the user's existing permissions and User Level that were previously configured and stored directly within the Kentik portal.
    - **Member**: Grant the user Member-level access to the Kentik portal.
    - **Deny Access**: Deny the user’s attempt to access the Kentik portal via SSO.

> **Tip**: For further explanation, see **User Level Determination Flow**and **Default Fallback Outcome Table**.

**Role Mapping** (pane):

This pane maps SSO attributes to **Role Sets** for granular RBAC permissions. A Role Set is a collection of RBAC Roles, and as an administrative tool, it simplifies the process of mapping complex permissions to a single attribute value from your Identity Provider (IDP).

* **Profile attribute key for Role Set:** Enter the SAML attribute key for Kentik to use to determine its RBAC Role Set assignment.
* **Mapping Table for Role Set:** Define a custom role set mapping table:

  + **+ Add Mapping** (button): Click to add a table row. Only active once a profile attribute key is defined (see above).
  + **Attribute Value**: Enter the raw SAML attribute string for the mapping table entry.
  + **Mapped Role Set** (dropdown): Choose the Kentik RBAC role set for the mapping table entry.
  + **Remove** (trash icon): Click to delete the mapping table entry.

  ![Role mapping configuration for user attributes in the Kentik Portal interface.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/SSO-mapping-table-for-role-set.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A13Z&se=2026-05-12T09%3A44%3A13Z&sr=c&sp=r&sig=rf4ijsjr7pKUfJNvratPY2CiST%2Bs%2B9bn3HSNQMOlzTg%3D)

  *Define a custom Role Set mapping table.*
* **Default Fallback:** Choose from the following defaults when the Role Set attribute is missing or unmapped:

  + **Portal-Local User assigned Roles**: Retain the user's existing permissions and Role Set that were previously configured and stored directly within the Kentik portal.
  + **Assign this default Role Set**: Choose a default Role Set from the dropdown to be applied to users on fallback.

![The SSO Attribute Mapping tab of the Authentication & SSO page, with user and role mapping options.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/SSO-sso-attribute-mapping-tab-main.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A13Z&se=2026-05-12T09%3A44%3A13Z&sr=c&sp=r&sig=rf4ijsjr7pKUfJNvratPY2CiST%2Bs%2B9bn3HSNQMOlzTg%3D)

*The SSO Attribute Mapping tab of the Authentication & SSO page.*

## User Level Determination Flow

When a SSO user successfully authenticates with the Identity Provider (IDP), Kentik determines their final User Level by following a specific hierarchy of checks. This flow governs the user's ultimate access level, especially for users not previously registered (auto-provisioned) or those whose User Level attribute is missing or unmapped.

| Check | Condition | Resulting User Level |
| --- | --- | --- |
| 1. Mapped Attribute | The IDP provides a valid SAML attribute string that is found in the Mapping Table for User Level. | Mapped User Level (e.g., Admin, Member) |
| 2. Missing/Unmapped Attribute | The IDP either does not send the User Level attribute OR the attribute string is not found in the Mapping Table. | Default Fallback Setting is applied. |
| 3. Auto-Provisioning | New user logs in AND **Allow auto-provisioning** is On, AND User Level attribute is missing/blank. | Member |

### Default Fallback Outcome Table

If the user is Unmapped (Condition 2 above) and not subject to auto-provisioning, the access decision is based entirely on the **Default Fallback** radio button selection:

| Default Fallback Setting | User Status | Final Access Level |
| --- | --- | --- |
| Portal-Local User Attribute | Unmapped | LOCAL\_USER\_LEVEL (Retains existing permissions stored in Kentik) |
| Member | Unmapped | Member-level access |
| Deny Access | Unmapped | DENIED ACCESS |
