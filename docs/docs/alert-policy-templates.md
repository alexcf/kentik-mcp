# Alert Policy Templates

Source: https://kb.kentik.com/docs/alert-policy-templates

---

This article discusses the **Alert Policy Templates** page of the Kentik portal.

> **Notes:**
>
> * For general information about policy-based alerting, see **Policy Alerts Overview**.
> * For information on the Alerting page, see **Alerting**.
> * For information on alert notifications, see **Notification Channels**.
> * For information on alert mitigations, see **Mitigation Overview**.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(319).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A20Z&se=2026-05-12T09%3A32%3A20Z&sr=c&sp=r&sig=MMJ50vtQOVY1It5E3wUc2w4Kvs45zRZeO4GD8GovGF8%3D)

*Kentik-provided templates can be cloned and tailored to your organization's alerting needs.*

## About Alert Policy Templates

Kentik’s alerting system includes templates for common traffic anomaly policies, allowing you to respond with manual or automated mitigation. These templates are listed on the **Alert Policy Templates** page. You can’t add or modify policy templates, but you can clone them to your Alert Policies page. The cloned template remains unchanged, while your new policy is fully editable to meet your organization’s needs.

## Alert Policy Templates Page

The Alert Policy Templates page is a list of Kentik-provided policy templates to cover the most common network traffic anomalies.

> **Note**: Only Administrators can access the Alert Policy Templates page.

### Policy Templates Access

To access the Alert Policy Templates page:

1. Choose **Alerting** from the portal's main navbar to go to the Alerting page.
2. Click **Manage** **Policies** at the upper right to go to the Policies page.
3. Click **Alert Policy Templates** at upper right to go to the Alert Policy Templates page.

### Policy Templates Page UI

The Alert Policy Templates page includes the following UI elements:

* **Subnav**: Breadcrumbs for the page.
* **Show/hide filters** (filter icon): Toggles the expanded/collapsed Filters pane.
* **Search field**: Enter text to filter the policies in the **Policy Templates List**. The list will show only policies with matching text in at least one column (ID or Name). The field also displays lozenges for filters applied with the **Filters** pane, if any.
* **Filters pane**: Use to narrow the alert policies shown in the **Policy Templates List****,** (see **Policy Templates Filters**).
* **Policy templates list**: A table listing the available policy templates (see **Policy Templates List**).

### Policy Templates List

The Policy Templates list is a table showing all Kentik-provided policy templates. The table includes the following columns:

* **ID**: The template's unique ID.
* **Type**: The type of template, NMS, Traffic, Cloud, or Protect (see **Policy Types**).
* **Name**: The name and description of the template. Once cloned, you can change the name and description of the resulting policy.
* **Metrics**: The units (e.g., bits/s, packets/s) by which the template measures incoming flow data (see **Data Funneling**). The primary metric is listed first, followed by secondary metrics (if any).
* **Dimensions**: The dimensions in the template, combining to make a key for how traffic is grouped for evaluation (see **About Keys**and **Dimensions Reference**).
* **Create**: Clones the template and opens the resulting policy for editing (see **Clone a Policy Template**). Once saved, the new policy will appear in the **Policies List**, while the template remains unchanged.

> **Note**: To see more template details, click its row to open its **Template Details Drawer**. Click same row again to close the drawer.

### Policy Templates Filters

Filter the templates in the **Policy Templates List** using the controls in the **Filters** pane. The pane includes:

* **Reset To Default** (appears only when a filter is specified): Click to clear all filters.
* **Type**: Click a policy type box, NMS, Traffic, Cloud, and/or Protect to restrict the list (see **Policy Types**).
* **Template ID**: Search for a specific template by its ID number (no partial matches).

  > **Note**: The template ID differs from the policy ID assigned to a policy cloned from that template.

### Template Details Drawer

The Policy Template Details drawer, which provides a summary of a template’s settings, opens when you click a row in the **Policy Templates List****.** It includes the following:

* **ID**: The unique template ID.![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(320).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A20Z&se=2026-05-12T09%3A32%3A20Z&sr=c&sp=r&sig=MMJ50vtQOVY1It5E3wUc2w4Kvs45zRZeO4GD8GovGF8%3D)
* **Name**: The Kentik-provided template name.
* **Create** (button): Click to create a new policy from the template (see **Clone a Policy Template**).
* **Description**: The Kentik-provided template description, if any.
* **Dataset** (pane): Summary of settings that determine evaluated traffic for the template (see **Policy Dataset Settings**).
* **Thresholds** (pane): Summary of threshold settings for the template. The number of thresholds defined in the template is shown in parentheses (see **Policy Threshold Settings**).
* **Baseline** (pane): Summary of baseline settings for the template (see **Policy Baseline Settings**).

When you clone a template into a policy, you can modify the policy settings while the template remain unchanged.

## Clone a Policy Template

To create an alert policy from a template using the Alert Policy Templates page:

1. Navigate to the Alert Policy Templates page (see **Policy Templates Access**).
2. Click **Create** for the policy you'd like to clone.

   1. Alternatively, click **Create** in the template's **Template Details Drawer**.
3. Modify the policy as needed (see **Policy Settings**).
4. Use the indicators in the **Policy Summary Pane** to check for missing information or errors.
5. Click **Save.** The new policy will now appear in your **Policies List**.

> **Note**: To create an alert policy from a template on the **Policies** page, click the arrow on the **Add Policy** button and select **Add Alert Policy from Template** (see **Add Policy from Template**).
