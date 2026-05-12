# Policy Alerts Overview

Source: https://kb.kentik.com/docs/policy-alerts-overview

---

This article provides an overview of Kentik’s policy-based Alerting system.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(254).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A57Z&se=2026-05-12T09%3A30%3A57Z&sr=c&sp=r&sig=2z0bmStnVUJ2InM0wSSg33i%2F0Va5OzikLbGUUlpJG4Q%3D)

*Triggered based on policy settings, alerts provide details on anomalous network activity or state.*

## About Alerting

Kentik's powerful alerting system analyzes your network traffic and detects anomalous patterns that might represent threats to availability or performance.

Alert policiesare a set of conditions, that when met, cause the policy to enter "alarm state" and generate an alert. Policies can generate notifications in response to an alert, and initiate automatic mitigations (built-in or third-party).

> **Notes:**
>
> * For information on alert policy settings, see **Alert Policies**.
> * For information on active/historical alerts, see **Alerting**.
> * For information on alert notifications, see **Notification Channels**.
> * For information on alert mitigations, see **Mitigations**.

## Alerts and Policies

Kentik's alerting system is implemented via alert policies of various types.

### Policy Types

When a policy is created, either by a Kentik user or, in the case of **Policy Templates**, by Kentik, it’s assigned one of these types.

> **Note**: Trigger type refers to the two alert-generation mechanisms, **NMS** and **Classic Threshold**, both choices in the **Add Alert Policy** dialog.

| Policy Type | Where created | Trigger Type | On Policies page? | On DDoS Defense page? |
| --- | --- | --- | --- | --- |
| **NMS** | Add Policy page | NMS or Classic Threshold | Yes | No |
| **Traffic** | * Add Policy page * Data Explorer (query-based) | Classic Threshold | Yes | No |
| **Cloud** | Add Policy page | Classic Threshold | Yes | No |
| **Protect** | Add Policy page | Classic Threshold | Yes | Yes |

> **Note:** The filters on the **Alert Policies** and **Alert Policy Templates**pages determine which policy types appear in the lists:
>
> * Checking the **NMS** filter includes both NMS policies and NMS threshold policies in the list.

### About NMS Policies

NMS alert policies are metrics-based and built on Kentik NMS (see **Network Monitoring**). They alert you when a monitored entity, such as a device, interface, or BGP neighbor, is unhealthy (e.g., down).

Like threshold alert policies (see **About Threshold Policies**), NMS policies can be added, cloned, and edited via the **Alert Policies** page:

* **Add**: Create a policy from scratch via the **Add Policy** button. Select **NMS** and click **Continue** to go to the Add Policy page for NMS policies (see **NMS Settings Page**).
* **Clone**: Duplicate an existing policy by choosing **Clone Policy** from the Action menu (vertical dots) in the **Policies** list (see **Clone a Policy**).
* **Edit**: Modify the settings of an existing NMS policy by choosing **Edit Policy** from the Action menu in the **Policies** list (see **NMS Settings Page**).

### About Threshold Policies

Threshold policies are sets of comparative evaluations that trigger alerts when one or more comparisons match (see **About Matches**). Alerts results in actions like notifications, DDoS mitigations, and the policy entering an ALARM state (see **Alert State**).  
  
Threshold policies can be Custom, DDoS, or Query-based. This doesn't affect their functionality but affects their display in the portal (see **Policy Types**).  
  
Each policy defines the characteristics of your network traffic that trigger alerts and the alerting system’s response. The configuration covers:

* **Evaluated traffic**: “What traffic flow data should be evaluated upon ingest into Kentik?”

  + Data sources, policy dimensions, metrics, and filters in the policy’s Dataset tab define the traffic scope and time interval for evaluations.
* **Comparison mode**: “To what should the current traffic be compared?”

  + Current traffic can be compared to a static value, a historical baseline, or track its presence or absence.
* **Thresholds**: “What differences between the current traffic and the comparator should trigger an alarm?”

  + Each alert can have up to five thresholds, each with its own comparison mode and settings that trigger an alarm, determine the timing of entering and leaving an alarm state, and specify actions to take.
* **Actions**: “What actions should occur in response to an alarm?”

  + Each threshold has independent action settings for notification (see **Notification Channels**) and mitigation (see **Mitigation Overview**). As an alert enters an alarm state, it’s added to the Alerts list (see **Alerts List**).

Once a policy is defined and saved, it appears on the **Alert Policies Page**, where policies can be added, cloned, and edited.

## Threshold Alerting Concepts

This section covers important concepts for working with threshold policies.

### About Matches

If a threshold policy is enabled, Kentik evaluates flow data from your network devices (routers, hosts, etc.) at the specified frequency for a match between the evaluated traffic and the policy thresholds (see **Threshold Conditions**). If a specified number of matches are found within a given time span (see **Threshold Frequency**), an alarm triggers, and the system responds with the matched threshold’s actions.

> **Note**: Policies enable powerful control but can be challenging to configure. The Kentik support team encourages you to contact us at **support@kentik.com** for assistance with alert policy configuration.

### About Keys

At Kentik, a key is a unique identifier for a set of dimensions. For instance, if the dimensions are Destination IP/CIDR, Destination Port Number, and Protocol, each unique combination of values for these three dimensions constitutes a key.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(256).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A57Z&se=2026-05-12T09%3A30%3A57Z&sr=c&sp=r&sig=2z0bmStnVUJ2InM0wSSg33i%2F0Va5OzikLbGUUlpJG4Q%3D)

*In this top-x table, each row represents a unique combination of values for each columns’s dimensions, also known as a key.*

Threshold alerting dimensions are chosen on the Dataset tab (see**Policy Dataset Settings**). The top-X ranking of traffic is performed by evaluating the volume of the traffic (primary metric across the selected devices, filtered by specified filters) for each key.

### Additional Alerting Resources

For more information about alerting and related concepts:

* **About Alert Thresholds**
* **About Historical Baselines**
* **About Alert Policy Templates**
* **Notification Channels**
* **About Alert Suppressions & Silences**
* **About Mitigation**
* **About Manual Mitigation**

## Alerting Pages

The pages used to configure and manage alerting and mitigation are as follows:

| Page | Portal Location | Description |
| --- | --- | --- |
| **Alerting** | Main menu » **Alerting** | Info about current/previous alerts in your organization. |
| **Policies** | Alerting » **Manage Policies** | List of alert policies (see **Alerts and Policies**), from which policies can be added, duplicated, and edited. |
| **Policy Templates** | Alerting » Configure Alert Policies » **Policy Templates** | List of Kentik-provided alert policy templates for common notification situations (see **Policy Templates**). |
| **Mitigations** | Protect » **Mitigations** | Info about your current/past mitigations (see **Mitigations**). |
| **Manage Mitigations** | Settings » **Mitigations** | List of available mitigation platforms and their methods (see **Manage Mitigations**). |
| **Manual Mitigation** | Protect » **Mitigations** | Dialog for manually applying a mitigation in real time, avoiding an alert with an alarm state (see **Manual Mitigation**). |
| **Alert Suppressions & Silences** | Settings » **Alert Suppressions & Silences** | List of “patterns” representing a set of conditions under which alerts will not be triggered (see **About Alert Suppressions & Silences**). |
| **Notifications** | Settings » **Notifications** | List of notification channels (see **Notification Channels**), i.e. sets of notification types (e.g., email) and targets (e.g., a list of email addresses). |
