# Configuring Third-Party Mitigations

Source: https://kb.kentik.com/docs/configuring-third-party-mitigations

---

Kentik Protect allows you to seamlessly integrate with external, third-party mitigation providers—such as Radware, A10 TPS, and Cloudflare Magic Transit—using API-driven automated responses.

This guide covers the generic setup workflow for third-party platforms, as well as a specific configuration example for Cloudflare Magic Transit.

## Before You Begin

* **Review core concepts:** If you are new to Kentik Protect, please review the **Mitigation Overview** to understand how platforms, methods, and automated triggers work.
* **Allow Kentik IPs:** Ensure your third-party mitigation platform is configured to allow API requests from Kentik's infrastructure. You must allow the following IP ranges:

  + **US**:`209.50.158.0/23` (IPv4) and `2620:129:1::/48` (IPv6)
  + **EU**: `193.177.128.0/22` (IPv4)

> **Note:** The platform and method settings for third-party mitigations are made on the portal's **Manage Mitigations** page (**Settings »** **Mitigations**).

## Third-Party Mitigation Workflow

The following steps provide a generic guide to configuring a third-party mitigation:

1. On the portal's Manage Mitigations page (**Settings »** **Mitigations**), click the **Add Mitigation Platform** button.
2. On the Add Mitigation Platform page, under **Select Your Mitigation Type** (see **Common Platform Settings**) choose the type of third-party mitigation you'd like to use (Radware, Cloudflare Magic Transit, or A10 TPS).
3. Under Assign Mitigation Methods, click the **Add New Method** button, which opens the **Add Mitigation Method** dialog.
4. On the **General** tab of the dialog, make the settings described in **General Method Settings**.

   > **Note:** The settings dialog for a Cloudflare Magic Transit method includes only general settings and use no tabs.
5. If the mitigation type is A10 or Radware, the dialog will include a **Details** tab. Configure the settings on this tab, which vary depending on the mitigation type.

   > **Note:** Customers are advised to make these settings in consultation with a support representative of the third-party vendor.
6. Click **Add Mitigation Method** to save your method.
7. Back on the Add Mitigation Platform page, under **Configure API**, fill in the API settings particular to the third-party mitigation type of your platform (see **Configure Platform APIs**).
8. Under **Finish Up**, fill in the **Name** and **Description** fields.
9. Click **Create Mitigation Platform** to save your new platform.

## Cloudflare MT Setup

The following procedure is specific to Cloudflare Magic Transit (MT) but illustrates the more general mitigation setup process outlined in **Third-Party Mitigation Workflow**:

1. On the Manage Mitigations page (**Settings »** **Mitigations**) in the Kentik portal, click the **Add Mitigation** **Platform** button.
2. On the Add Mitigation Platform page, select **Cloudflare Magic Transit** as the Mitigation platform.

   > **Note:** Cloudflare applies Magic Transit mitigation only when traffic exceeds minimums (100K pps for TCP/UDP; 60K pps for ICMP/GRE). Assigning Cloudflare MT to a Kentik alert below these thresholds may show mitigation as active when it isn't. For lower volumes, use an alternative mitigation platform (RTBH, Flowspec, etc.).
3. Click the **Add New Method** button, which opens the Add Mitigation Method (Cloudflare Magic Transit) dialog.
4. Fill in the fields described in **General Method Settings**.
5. Click **Add Mitigation Method** to save your method.
6. Back on the Add Mitigation Platform page, fill in the required fields, including **API Login**, **Cloudflare Account ID**, and **API Key** (see **Configure Platform APIs**).
7. Fill in a name and description of the new platform (see **Common Platform Settings**).
8. If the mitigation method you created in steps 3 to 5 isn't listed yet under **Assign mitigation methods**, click **Add Existing Method** and choose the method from the drop-down list. It may take a few moments to display.
9. When the settings are complete, click the **Create Mitigation Platform** button to save your new platform and close the dialog.

Once you've created a Cloudflare MT mitigation for your organization, it may be deployed manually or automatically (e.g., in response to conditions defined in a Custom Insight threshold); see **Mitigation Deployment.**

> **Note:** This procedure only covers setup in the Kentik portal. You must also configure mitigation in the Cloudflare account. For further information, refer to Cloudflare’s documentation and support for Cloudflare Magic Transit.
