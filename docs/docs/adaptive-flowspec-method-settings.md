# Adaptive Flowspec Method Settings

Source: https://kb.kentik.com/docs/adaptive-flowspec-method-settings

---

This article covers the specific configuration fields and settings required when building an Adaptive Flowspec mitigation method from Kentik’s **Manage Mitigations**page.

## Adaptive Flowspec Method Details

In addition to the settings on the **General** tab of the mitigation method dialog, configuring an **Adaptive Flowspec**-based method involves settings specific to dynamically matching and filtering traffic across the following tabs:

### Traffic Analysis Tab

This tab contains three sections that control how Adaptive Flowspec responds to alerts:

**Mitigation Goal**: Defines the overall mitigation effect you want to achieve:

* **Max Flowspec Rules**: Enter the maximum number of Flowspec rules a single mitigation may generate and propagate to a single BGP device.
* **Traffic Volume**: Enter the minimum percentage of policy-matched traffic to block.
* **Traffic Metric**: The metric used for the Traffic Volume setting (Bits, Packets, or Flows Per Second).

**Algorithms**: Defines which traffic analysis algorithms are run, and in which order, to satisfy the Mitigation Goal. If an algorithm doesn't meet the goal, the next one is tried:

* **Top Source Port Numbers**: When switched On, blocks the top source port numbers as measured by the Traffic Metric.

  + **Ignore Protocol**: When switched On, treats the same port across different protocols (e.g., TCP/80 and UDP/80) as one combined port.
  + **Minimum Traffic Volume**: Enter the percentage of traffic below which the source ports should be skipped (i.e. not blocked).
  + **Exclude Ports and Ranges**: Enter one source port/range per line to prevent them from being blocked, e.g., 0-1024, 0-1024/TCP, 80/UDP.
* **Top Source CIDRs**: Aggregates traffic by the top source IP addresses and finds an optimal set of covering IP prefixes to block.

  + **IPv4 Budget**: Enter the maximum total number of IPv4 addresses this mitigation can block, expressed as an IP prefix length, e.g., /4.
  + **IPv6 Budget**: Enter the maximum total number of IPv6 addresses this mitigation can block, e.g., /8.
  + **Exclude CIDRs**: Enter a comma-separated list of IPv4 and/or IPv6 CIDRs to exclude them from blocking.
* **Top Destination Port Numbers**: Functions identically to the Top Source Port Numbers algorithm, but applied to the destination traffic.
* **Top Destination CIDRs**: Functions identically to the Top Source CIDRs algorithm, but applied to the destination prefixes.

**Failsafe Mode**

* **Deploy failsafe rules if the goal is not reached**:

  + **When switched On**: If none of the algorithms can satisfy the mitigation goal (e.g., reverting to blackholing the attacked service), failsale rules will be deployed.
  + **When switched Off**: The mitigation will fail to start and no rules will be propagated.

### Flowspec Tab

This tab configures the Flowspec rule template and matching behaviors, and has the following panes and controls.

#### Traffic Matching

This pane has settings to define how traffic is matched by inferring values (e.g., Source Ports, TCP Flags) from the alarm that triggered the mitigation.

* **Default Infer from Alarm**: Choose from the dropdown whether to infer matches from the alert that triggered the mitigation:

  > **Note**: This sets the default for all Traffic Matching controls. You can override this setting using Infer from Alarm on a per-

  + **Enabled** (default): Inferring is turned on. Your value is overwritten with the matching alert dimension value.
  + **Disabled**: Inferring is turned off. No attempt is made to overwrite your value with the alert dimension value.
  + **Merge**: Your value is added to the matching alert dimension value.
* **Infer from Alarm**: Overrides the global setting to enable or disable inference for a specific match.

#### **Traffic Filtering Actions**

This pane allows you to override, for the **matching traffic**, the default Flowspec action (Discard) and trigger arbitrary actions by adding specific BGP community attributes:

* Turn one or more switches to override the default action for the matching traffic (e.g., Rate Limit, Large Community).
* For any enabled actions, enter any additional required information (e.g., IP address), noting the formatting hints in the UI help text.

### Preview Tab

This tab allows you to safely preview the Flowspec rules that would be generated on past alerts without propagating any changes to your network. The tab includes the following panes and controls:

#### Choose an Event

* **Alert ID**: Enter the ID of a **past Kentik alert** to test against in UUID format, e.g., 019d033b-5593-7c07-b56f-a222b45ddaa3.
* **Load Alert Event** (button): Click to retrieve the entered alert ID and generate the preview.

#### **Preview Results**

For the entered alert ID, a summary is generated of what would be achieved and prints the top 100 Flowspec rules in text form based on your current configuration.

> **Note:** For general method settings like notification channels and automation grace periods, return to the **Mitigation Methods** article.

FPS stands for flows per second, which is the rate at which flow metadata is sent to Kentik from a given customer. It directly correlates with the resources required to provide that customer with the Kentik service.
