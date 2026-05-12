# Alerting

Source: https://kb.kentik.com/docs/portal-alerting

---

Kentik’s alerting systems notify you of DDoS attacks and various traffic and performance anomalies.

## Policy-based Alerting

Policy-based alerting uses alert policies to define conditions for alerting. Policies are managed with the portal’s **Policy Settings** pages, accessed via the **Alerting** page in the portal’s main nav menu. Each policy-based alert includes thresholds for triggering an alarm and actions (notifications and/or mitigations) when “alarm state” is entered.  
  
For more information, refer to the following KB articles:

* **Policy Alerts Overview**: A high-level explanation of the policy-based alerting system.
* **Alert Policies**: Managing policies and setting a policy.
* **Alerting**: Viewing active or historical alerts.
* **Notifications**: Specifying alert notification recipients and methods.
* **Mitigations Overview**: Defining mitigation platforms, methods, and application rules.

> **Note:** Please contact **support@kentik.com** to get help with policy-based alerting.

## Synthetics Alerting

Kentik’s synthetic testing (see **About Synthetics**) includes an alerting system that generates alerts based on test results:

* Set thresholds for a tested entity’s health status on the Health tab of the **Test Control Center** when adding or editing a test (see **Health Settings**).
* The health status metrics depend on the test type (see **Health Options by Test Type**).
* Specify alert conditions on the **Alerting and Notifications** tab (see **Alerting Conditions**).

> **Note:** For more information on Synthetics alerting, see **Alerting and Notifications**.
