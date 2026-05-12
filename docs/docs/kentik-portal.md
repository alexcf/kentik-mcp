# Kentik Portal

Source: https://kb.kentik.com/docs/kentik-portal

---

This article provides an introduction to the Kentik portal. Also see the **Portal Overview**from the Portal section of the Knowledge Base.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(8).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A53Z&se=2026-05-12T09%3A29%3A53Z&sr=c&sp=r&sig=oVSDiiGFdrh2rXpip3u3QVUolNRmz%2Fh%2BcwRMmmT2pOM%3D)

*The default landing page of the Kentik v4 portal is the Observation Deck.*

## About the Kentik Portal

The v4 portal is the main user interface for Kentik (see **Platform vs. Portal**):

* Actively maintained on a continuous deployment model
* Offers access to the full capabilities of the Kentik platform.

To get started, read **Portal Overview**, then explore the modules, workflows, and settings of the v4 portal.

> **Note:** Kentik’s v3 portal was fully retired as of March 31, 2025 and is longer accessible to customers. For assistance in transitioning to the current portal, contact **Customer Care**.

## Browser Support

Kentik recommends the latest versions of thes browsers for use with its portal:

* **Google Chrome**
* **Mozilla Firefox**
* **Apple Safari**
* **Microsoft Edge**

## Portal Login

Login to the Kentik portal is covered in the following topics.

### About Portal Login

Kentik supports various portal login approaches, each balancing convenience and security. These approaches aren’t mutually exclusive; users can have multiple layers of authentication, and those without one method can still access via another.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/Portal-Basic_login-338w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A53Z&se=2026-05-12T09%3A29%3A53Z&sr=c&sp=r&sig=oVSDiiGFdrh2rXpip3u3QVUolNRmz%2Fh%2BcwRMmmT2pOM%3D)Portal login authentication has two layers: a required first layer and an optional second layer:

* **First layer** (required): Basic authentication using one of:

  + Email/password (see **Basic Login**).
  + Single sign-on (see **SSO Login**).
* **Second layer** (optional): More secure access using two-factor methods (see **Two-factor Login**):

  + Time-based One-Time Password (TOTP).
  + YubiKey, a hardware authentication device from Yubico.

> **Notes:**
>
> * The **One-time Token** field on the Two-factor Authentication page (see **Two-factor Login**) accepts any configured 2FA methods in your user profile (see **Authentication Settings**).
> * Register as a user to log into Kentik (see **Add a User**).
> * By default, the portal page upon login is the **Observation Deck**. Change this in your profile by clicking the User icon in the main nav bar, and selecting **User Profile** » **Default Settings** (see **User Profile**).

### Authentication Sequence

Portal login authentication follows this sequence:

* **First layer**:

  + If SSO is enabled, use SSO (see **SSO Login**).
  + If SSO is not enabled, log in with email/password on the main portal login page (see **Basic Login**).
* **Second layer**:

  + If TOTP and/or YubiKey is enabled, proceed to the Two-factor Authentication page after first-layer authentication. Enter a valid TOTP key or YubiKey key in the **One-Time Token** field (see **Two-factor Login**).

> **Notes:**
>
> * The **One-time Token** field accepts any configured 2FA methods in your user profile (multiple methods may be configured simultaneously; see 2FA Authentication list in **Two-factor Authentication**).
> * If **SSO Required** is enabled on the Admin » Single Sign-on page, only Super Admins can use **Basic Login**; all others must use SSO (see **Additional Configuration Options**and **About Super Admin Users**).
> * If **Disable 2FA** is enabled on the Admin » Single Sign-on page, the second-layer authentication enabled for the user is bypassed when signing in with SSO.

### Basic Login

Basic login requires an email address and password associated with a registered user:

1. Go to **https://portal.kentik.com/login**.
2. Enter your email and password and click **Login**.
3. The portal opens to your specified landing page.

> **Note*:*** If your organization is registered with Kentik in the EU, use **https://portal.kentik.eu/login**.

### SSO Login

SSO login requires your organization to be set up for SSO in Kentik (see **Single Sign-on**).

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/Portal-SSO_login-348h338w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A53Z&se=2026-05-12T09%3A29%3A53Z&sr=c&sp=r&sig=oVSDiiGFdrh2rXpip3u3QVUolNRmz%2Fh%2BcwRMmmT2pOM%3D)Once set up, SSO access is enabled via either of the following paths:

* **SSO landing page**:

  + Go to `https://portal.kentik.com/login/sso/company_shortname`.
  + Click **Login**.
* **Direct**: Go to `https://portal.kentik.com/sso/company_shortname`.
* > **Note*:*** If your organization is registered with Kentik in the EU, use **https://portal.kentik.eu**.

Both paths lead to an authentication check (see **How SSO Works**):

* If authenticated, you'll be automatically logged into the Kentik portal.
* If not, you'll be redirected to your identity provider’s login screen, then back to the Kentik portal upon successful authentication.

**Forgot your SSO URL?**![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(9).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A53Z&se=2026-05-12T09%3A29%3A53Z&sr=c&sp=r&sig=oVSDiiGFdrh2rXpip3u3QVUolNRmz%2Fh%2BcwRMmmT2pOM%3D)

If your organization has SSO enabled, but you don’t know the SSO URL.

Follow these steps to request the SSO URL:

1. Go to your **Basic Login** page
2. Click **Looking for Single sign-on?** below the Login button.
3. Enter a Kentik-registered email address to which to send the SSO login URL.

   > **Notes:**
   >
   > * If not authenticated, navigate to an SSO URL above to use SSO. Directly navigating to a portal page will redirect to the main portal login page (email/password) even with SSO enabled.
   > * Kentik does not support single sign-out.

### Two-factor Login

If TOTP (time-based one-time password) and/or YubiKey are enabled in the **Authentication** tile on your **User Profile** page (see **Authentication Settings**), an extra layer will be added to the login process:

1. ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/Portal-2FA_login-173h338w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A53Z&se=2026-05-12T09%3A29%3A53Z&sr=c&sp=r&sig=oVSDiiGFdrh2rXpip3u3QVUolNRmz%2Fh%2BcwRMmmT2pOM%3D)Follow the first-layer login process (**Basic Login** or **SSO Login**).
2. On the Two-factor Authentication page, enter the one-time token (key code) from either of the following:

   1. A TOTP-compliant application on your mobile phone or device.
   2. A YubiKey dongle plugged into your machine.
3. Click **Verify** to open the portal to your specified landing page.

> **Notes*:***
>
> * The **One-time Token** field accepts any configured 2FA methods in your **User Profile** (multiple methods may be configured simultaneously; see **Authentication Settings**).
> * If **Disable 2FA** is enabled in Admin » **Single Sign-on** (see **Additional Configuration Options**), the second layer of authentication (TOTP or YubiKey) won’t be required for SSO.
> * Common TOTP-compliant apps include Authy, Google Authenticator, Duo Mobile, LastPass, Microsoft Authenticator, FreeOTP Authenticator, and 1Password Authenticator.
> * Kentik recommends running TOTP apps on a separate device (e.g., mobile or tablet) to avoid defeating the purpose of 2-factor authentication.
