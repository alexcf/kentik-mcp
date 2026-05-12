# User Profile

Source: https://kb.kentik.com/docs/user-profile

---

This article discusses the **User Profile** page in the Kentik portal.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(112).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A01Z&se=2026-05-12T09%3A38%3A01Z&sr=c&sp=r&sig=BuX3hXuQNBuhO7WAnDl65wr0AlUCT%2B%2B6iX9LP6Qx%2Fjw%3D)

*The Profile page enables you to customize user-specific display settings and preferences.*

## About User Profile

The portal's user profile serves two related purposes for Kentik users:

* Manage user-specific settings that control how information (e.g., time zone, theme etc.) is presented within the Kentik portal.
* Manage user-specific preferences in areas such as authentication, notifications, and landing pages.

Settings made in your profile affect only your experience of Kentik; they have no effect on other Kentik users in your organization.

## Profile Page

The Profile page is a central place to view and specify user-specific settings for the Kentik portal. These settings enable you to tailor aspects of Kentik to your own needs. To access your profile, click the user icon at the far right of the main navbar, then choose **User Profile** from the **User Menu**.  
  
The User Profile page is organized as a collection of tabs that each represent different categories of settings:

* **General**: Your role, name, and email address, as well as defaults and notification preferences. See **General Settings**.
* **Authentication**: SSO, password and API token resets, and two factor authentication settings. See **Authentication Settings**.
* **Chart Colors:** Personalization settings for how you view the portal. See **Chart Colors**.

> **Note:** Additional settings related to the administration of users are covered in **Users**.

## General Settings

The **General** tab of the Profile page is covered in the following sections.

> **Note:** Use the **Save** button to save changes made to any settings on the **General** tab.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(113).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A01Z&se=2026-05-12T09%3A38%3A01Z&sr=c&sp=r&sig=BuX3hXuQNBuhO7WAnDl65wr0AlUCT%2B%2B6iX9LP6Qx%2Fjw%3D)

*Settings such as user information, themes, and landing page are set on the General tab.*

### User Information

This pane contains basic settings related to the individual user:

* **Role** (read-only): Your user level in Kentik: Member, Administrator, or Super Administrator (see **About Users**).
* **Full name**: Your full name as registered with Kentik.
* **Email**: Your email address as registered with Kentik.

### User-Specific Theme

The theme determines the background colors used for the portal UI. Three choices are available:

* **Light**: Light background, often considered "standard" for daytime or brightly lit environments.
* **Dark**: Dark background, often used in low light environments.
* **Auto** (OS Setting): The theme follows your OS settings:

  + If your OS is set to Dark mode, the Kentik portal will also display in Dark mode.
  + If your OS is set to change with the time of day, the Light theme is used during the day and the Dark theme is used at night.

The color palette for Light and Dark themes is set on the **Chart Colors** tab (see **Chart Colors**).

### User-Specific Defaults

The following user profile settings enable customization of the user experience in the portal:

* **Landing page**: Radio buttons that determine which Kentik portal page you'll be taken to upon login (see **Default Landing Page**).
* **Time zone**: Radio buttons that determine the time zone in which times are expressed, either UTC or local.
* **Data Explorer**:

  + Historical Overlay: A switch that determines whether the Historical Overlay display setting in Data Explorer is enabled by default.
  + rDNS Lookup: A switch that determines whether reverse DNS (rDNS) lookups are performed by default when querying IPs. Disabling this option will speed up IP queries.
* **Network Explorer**:

  + Show Other Traffic: A switch that determines whether Network Explorer shows all of the remaining aggregated traffic outside of the Top N charted time series.

#### Default Landing Page

This setting determines which Kentik portal page you'll be taken when you first log into the portal. Use the radio buttons to choose one of the following:

* **Observation Deck** (default): The **Observation Deck** is a user-configurable high-level view of your organization’s network.
* **Network Explorer**: The **Network Explorer** page in the Core section of the portal provides you with a high-level overview of the traffic on your network.
* **Library**: The **Library** provides a single place to find, view, and manage views of your organization’s network data.
* **Synthetics**: The **Synthetics Dashboard** module of the Synthetics section provides visibility into the health, performance, and availability of your network.
* **Data Explorer**: **Data Explorer** is Kentik’s primary interface for manually exploring the network data stored in the main tables of the Kentik Data Engine. The query settings of the Data Explorer sidebar will be set to defaults.
* **Alerting**: The **Alerting** page provides a list of current and historical alerts generated by Kentik’s alerting system.
* **Insights**: The **Insights** page provides a centralized list of your organization’s insights (Kentik and custom).
* **Dashboard**: Displays a drop-down list from which you can choose a dashboard. Filter the list to a specific dashboard by starting to enter its name.
* **Saved View**: Displays a drop-down list from which you can choose a saved view. Filter the list to a specific saved view by starting to enter its name.
* **Market Intelligence**: The **Kentik Market Intelligence** page in the Service Provider section of the portal.

### Notification Preferences

Notification preferences relate both to notifications that appear in the portal and those that can be emailed to the user. The following toggle switches turn different types of notifications on/off:

* **Service Updates**: A switch that enables notifications about service status (e.g., scheduled maintenance, outage), which are emailed to the user and displayed in the portal during an active session.
* **Product Updates**: A switch that enables notifications about the product (e.g., updates to features, interface, capabilities), which are emailed to the user and displayed in the portal during an active session.
* **BGP Session Events**: A switch that enables emails sent when there is a non-service-affecting interruption (flap) in BGP.

> **Note:** Product and service updates will be displayed at login even if the corresponding switches (above) are off.

       ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(116).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A01Z&se=2026-05-12T09%3A38%3A01Z&sr=c&sp=r&sig=BuX3hXuQNBuhO7WAnDl65wr0AlUCT%2B%2B6iX9LP6Qx%2Fjw%3D)

## Authentication Settings

The Authentication tab is covered in the following topics.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(117).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A01Z&se=2026-05-12T09%3A38%3A01Z&sr=c&sp=r&sig=BuX3hXuQNBuhO7WAnDl65wr0AlUCT%2B%2B6iX9LP6Qx%2Fjw%3D)

*Settings on the Authentication tab control how you access Kentik.*

### Single Sign-on

If Single Sign-on (see **Authentication & SSO**) is enabled for your company, you will see a **Single Sign-on** pane at the top of the page. The pane contains information about logging in with SSO.

> **Note:** Users whose level is Super Administrator can access Single Sign-on settings by selecting **Authentication & SSO** from the**Organization Settings**menu**.**

### Password

The **Password** pane of the **Authentication** tab contains a **Reset****Password** button. When the button is clicked a password reset email is sent to the address that you provided on the General tab, and you’ll see a notification at the top of the page confirming that password reset has been initiated. Follow the instructions in the email to reset your password.

### Two-factor Authentication

The**Two-factor****Authentication** pane of the **Authentication** tab contains settings that support enhanced security by enabling the following types of two-factor authentication (2FA):

* **YubiKey**: YubiKey is a hardware authentication device manufactured by **Yubico** that supports one-time passwords, public key encryption and authentication, and the Universal 2nd Factor protocol developed by the FIDO Alliance.
* **TOTP**: TOTP uses authenticator applications from companies such as Cisco, Google, Microsoft, and Twilio to implement time-based one-time password authentication based on **IETF RFC 6238**.

The 2FA settings in this pane include:

* **2FA Authentication list**: A table listing of all of your currently registered 2FA methods (authenticators and YubiKeys) and their status (enabled/disabled); see **2FA Authentication List**.
* **Disable 2FA** (present if 2FA is enabled): A button that disables all of the two-factor authentication methods shown in the 2FA Authentication list.

  > **Note:** 2FA will be disabled automatically if you remove all items in the **Two-factor Authentication** list.
* **Add YubiKey**: A button that opens the **Add YubiKey** dialog so you can register a new YubiKey.
* **Add TOTP**: A button that opens the **Add TOTP** dialog so you can register a new TOTP authenticator.

> **Note:** When logging into Kentik, any valid 2FA method that is currently configured in the **Two-factor Authentication** list will be accepted in the **One-time Token** field (see **Two-factor Login**).

### 2FA Authentication List

The**Two-factor Authentication** list is a table listing all of your currently registered 2FA methods (authenticators and YubiKeys) and their status (enabled/disabled). The table includes the following columns:

* **Status**: Indicates the current status of the authentication method (checkmark if enabled, X if not).
* **Name**: The name given to the authentication method when it was added with the Add YubiKey or Add TOTP dialog, followed by the type of the method (YubiKey or TOTP).
* **Disable** (shield icon): A button that disables (but does not remove) this authentication method.
* **Remove** (trash icon): A button that opens a confirmation dialog allowing you to remove this authentication method for your organization's collection of methods. If you proceed the method will no longer enable you to access Kentik and it will no longer appear in the **Two-factor Authentication** list.

> **Note:** If you remove all the items from the **Two-factor Authentication** list, 2FA is automatically disabled.

### Add YubiKey

To register a YubiKey, use the **Add YubiKey** button in the Two-factor Authentication pane:

1. ![Form to add a YubiKey with fields for name and token entry.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/UP-add-yubikey-dialog.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A01Z&se=2026-05-12T09%3A38%3A01Z&sr=c&sp=r&sig=BuX3hXuQNBuhO7WAnDl65wr0AlUCT%2B%2B6iX9LP6Qx%2Fjw%3D)Click the **Add YubiKey** button.
2. Enter an appropriate name into the **Name** field.
3. With your YubiKey inserted into your machine, touch the key. The YubiKey will write a one-time password (OTP) into the **Yubi** **Token** field.
4. Click the **Save** button. A "registration successful" alert will appear briefly in Kentik, after which the dialog will close and the new authentication method will appear in the **Two-factor Authentication** list. When you next log in, you'll be able to use this YubiKey for two-factor authentication (see **Two-factor Login**).

### Add TOTP

To register a TOTP authenticator, use the **Add TOTP** button in the**Two-factor Authentication** pane. The following steps illustrate the process for the Google Authenticator app (steps may vary when using a different app):

1. Download the Google Authenticator app to your device.

   > **Note:** Always use a different device (e.g., a mobile phone) than the device from which you will access Kentik.
2. In the Authentication tile of your profile page, click the **Add TOTP** button.
3. In Google Authenticator, press the Add button (plus sign). A popup menu will appear.
4. Choose **Scan a QR code**. The camera application on your phone will open.
5. Point the phone toward the bar code in the **Register TOTP** dialog on your computer screen. Authenticator will display "secret saved" when the code is recognized, and a new item will appear in authenticator's list of authentication codes.
6. Enter the new authentication code from your phone into the **Token Validation** field in the **Register TOTP** dialog, then click the **Save** button. A "registration successful" alert will appear briefly in Kentik, after which the dialog will close and the new authentication method will appear in the **Two-factor Authentication** list. When you next log in, you'll be able to use TOTP for two-factor authentication (see **Two-factor Login**).

   ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(121).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A01Z&se=2026-05-12T09%3A38%3A01Z&sr=c&sp=r&sig=BuX3hXuQNBuhO7WAnDl65wr0AlUCT%2B%2B6iX9LP6Qx%2Fjw%3D)

### API Token

The **API Token** pane of the **Authentication** tab contains the following elements:

* **API Token**: A field containing a Kentik-generated string that is used to authenticate the user in an HTTP header (X-CH-Auth-API-Token).
* **Copy to Clipboard**: A button that copies the API token to your clipboard.
* **Reset API Token**: A button that resets your API token. A new token will be generated and written to the API Token field.

> **Note:** If a security note is displayed just under the Token field, the access setting for the API subsystem on the **Access Control** page (Organization Settings » **Portal Access Control**) is set to allow only specified APIs.

## Chart Colors

User-specific settings for visualizations are covered in the following topics.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(122).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A01Z&se=2026-05-12T09%3A38%3A01Z&sr=c&sp=r&sig=BuX3hXuQNBuhO7WAnDl65wr0AlUCT%2B%2B6iX9LP6Qx%2Fjw%3D)

*Use the Chart Colors tab to set custom colors for charts in light mode and dark mode.*

### About Chart Colors

User visualization settings, accessed via the **Chart Colors** tab on the **Profile** page, enable you to control the appearance of visualizations that are rendered in **Data Explorer** and in the dashboards and saved views of the portal **Library**. The tab is structured with a control area on the left and a preview area on the right that allows you to see the effect of your settings on various types of visualizations.

### Chart Colors Display Controls

The control area at left is used to set the colors and palettes used for visualizations when the portal is in the light theme and, independently, the dark theme. The control area includes the following settings and controls:

* **Theme**: A button group that determines whether you are currently previewing and setting colors and palettes for the light (standard) or dark theme, which are each set independently. This setting allows you to see how the colors that you select will look when the current theme is active in the portal.

  > **Note:** To set when each theme is active, see **User-specific Theme**.
* **Labels**: Controls that determine the color of the labels that are applied to the key of a chart. To set a label color, do one of the following:

  + Click the color swatch at left and choose from the popup color palette.
  + Enter a hex number in the field at right.
* **Overlays**: Controls that determine the color of the Total and Historical overlays that are used on time-series visualizations. The color controls for Overlays are the same as for Labels.
* **Quantitative** and **Qualitative**: Controls that determine the color palettes (see **Chart Colors Color Palettes**) that will be applied to quantitative and qualitative visualizations (see **Chart Colors Previews**).
* **Reset to Default**: A button that restores all settings of the current theme (light or dark) to default values.
* **Save**: A button that saves the current settings of both themes. The settings for each theme will be applied throughout the portal whenever that theme is active.

### Chart Colors Color Palettes

The controls in the lower part of the control area enable you to choose color palettes for visualizations in both of the categories (quantitative and qualitative) whose chart types are listed in **Chart Colors Previews**. You can either choose a preset palette (by clicking on it in the palette list) or build your own palette.  
  
To build a custom palette for either quantitative or qualitative visualizations:

1. Turn on the **Use custom values** switch. The preset palettes will be replaced with a set of individual color controls like those used for labels and overlays.
2. Use the controls to change any of the individual colors that you'd like to change.
3. Click **Save**.

   > **Notes:**
   >
   > * To go back to using a preset color palette, switch **Use Custom Values** off.
   > * Color palette settings will have no effect for the following chart types: Sankey, Horizon, Geo HeatMap, Gauge, and Table.

### Chart Colors Previews

The right side of the **Chart Colors** tab contains two preview panes that display charts like those in the Data Explorer chart display area (see **Explorer Chart Display**). The top preview pane is for quantitative view types and the bottom is for qualitative. The **Visualization** drop-down at the top right of each preview pane is used to choose a representative visualization of a given type so that you can see the effect of the color choices made with the **Chart Colors Display Controls**.  
  
 The following **Chart Visualization Types** are available to preview in the two panes:

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(123).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A01Z&se=2026-05-12T09%3A38%3A01Z&sr=c&sp=r&sig=BuX3hXuQNBuhO7WAnDl65wr0AlUCT%2B%2B6iX9LP6Qx%2Fjw%3D)

* **Preview Quantitative Views**: The **Visualization** drop-down at the top right of the **Quantitative Views** preview pane (top preview) includes the following chart types that display quantitative data:

  + Stacked Area Chart
  + 100% Stacked Area Chart
  + Stacked Column Chart
  + Horizon Chart
  + Bar Chart
  + Matrix
* **Preview Qualitative Views**: The **Visualization** drop-down at the top right of the **Qualitative Views** preview pane (bottom preview) includes the following chart types that display qualitative data:

  + Line Chart
  + Pie Chart
  + Sunburst

## Manage General Preferences

To set general preferences in the User Profile:

1. On the main navbar, click the **User** icon at the far right.
2. Choose **User Profile** from the drop-down menu.
3. On the **General** tab of the Profile page:

   1. Choose a theme (see **User-Specific Theme**).
   2. Choose a default landing page and time zone.
   3. Make default settings for Data Explorer and Network Explorer (see **User-Specific Defaults**).
   4. Specify what you'd like to receive notifications about (see **Notification Preferences**).
4. Click the **Save** button at the bottom right of the page.
