# DDoS Defense

Source: https://kb.kentik.com/docs/ddos-defense

---

Kentik’s DDoS Defense page is covered in this article.

|  |  |
| --- | --- |
| ***Purpose:*** | Streamline the protection of your network from DDoS attacks with easily customizable Kentik-provided preset alert policies for the most common attack profiles. |
| ***Benefits:*** | * Eliminate false positives/negatives and decrease response time with automatic machine learning-based traffic profiling. * Visualize attack characteristics and network impact. * Trigger automatic mitigation actions including RTBH, Flowspec, and external mitigation hardware or services. |
| ***Use Cases:*** | DDoS detection, mitigation, and administration |
| ***Relevant Roles:*** | Network Engineers/Architects, Network Security Engineers, NOC Engineers, Carrier Product Managers |

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(261).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A16Z&se=2026-05-12T09%3A35%3A16Z&sr=c&sp=r&sig=ewtz%2FmNP47hMhiEGtjTl1HcHVpamDRijvpRzdHtvvOk%3D)

*Automate DDoS detection, investigation, and mitigation.* 

## About DDoS Defense

DDoS attacks employ various attack vectors, including volumetric, invalid protocols, UDP, TCP, ICMP, amplification and reflection, and DNS. Effective detection and alerting require a comprehensive understanding of  these possibilities. Kentik's DDoS Defense workflow (Protect » **DDoS Defense**) is part of our overall Alerting system (see **Protect**) focused on DDoS policies that require minimal customization to protect against common attacks.

### Access DDoS Defense

The DDoS protection features have two portal locations:

* **DDoS Defense Page**: Go to Protect » **DDoS Defense** to quickly view ongoing and historical attacks and mitigations.
* **Policy Templates**: Go to Settings » Alerting and click **Alert Policy Templates** for Kentik-created alert policies to help you quickly deploy protection (see **Policy Templates**).

  + Use the **Filters** pane to narrow the templates to DDoS-targeted ones (see **DDoS Defense Templates**), then clone a template to add it to your **Alert Policies Page**.
  + The template remains unchanged and the resulting policy is fully editable to meet the needs of your organization.

To access the Policy Templates page from the DDoS Defense page:

1. Click **Manage Alert Policies** to go the Policies page
2. Click **Alert Policy Templates** to go to the Policy Templates page.

> **Note:** For attack profiles not covered by DDoS templates, create a custom alert policy in Protect » **Alerting** (see **Alerting**).

## DDoS Defense Page

The DDoS Defense page provides a high-level view of DDoS attack activity that has generated alarms from DDoS-related **Alert Policies**.

### DDoS Defense Page UI

The DDoS Defense page includes:

* **Manage Alert Policies** (button): Click to go to the **Alert Policies** page, filtered by Protect policy type.
* **DDoS Defense Breakdowns**: Bar charts showing alerts categorized by status, severity, and policy (see **DDoS Defense Breakdowns**). Breakdowns cover the last 24 hours. Hover over a bar to open a popup with additional information about its alerts.
* **Last 24 Hour Attack Activity**: A chart showing attack activity in the last 24 hours (see **Attack Activity Chart**).
* **Last 24 Hour Top Dest IPs**: Chart showing the top destination IP addresses in the last 24 hours as measured in different metrics (see **Top Dest IPs**).
* **Attacks Within the Last 24 Hours**: A table listing attacks within the last 24 hours along with details on those attacks (see **Attack Table**).

### DDoS Defense Breakdowns

The DDoS Defense breakdowns are bar charts at the top of the page representing different breakdowns of alerts. The charts are similar to the **Alerting Breakdowns** and contain the following UI elements:

* **Status**: The bars represent alerts by their state (see **Alerting Breakdowns**).
* **Severity**: The bars represent alerts by the severity level of their policy thresholds (see **Alerting Breakdowns**).
* **Policies**: The bars represent individual policies triggered in the last 24 hours, in descending order by alert count. Hover on a bar to see the name, type, ID, description, and alert count for the policy.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(262).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A16Z&se=2026-05-12T09%3A35%3A16Z&sr=c&sp=r&sig=ewtz%2FmNP47hMhiEGtjTl1HcHVpamDRijvpRzdHtvvOk%3D)

*These cards show a breakdown of alerts over the last 24 hours.*

### Attack Activity Chart

This chart shows attack activity over the last 24 hours. The blue bar indicates new alarms, while the red line shows active alarms. Hovering over the chart reveals a popup with counts of new and active alarms.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(263).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A16Z&se=2026-05-12T09%3A35%3A16Z&sr=c&sp=r&sig=ewtz%2FmNP47hMhiEGtjTl1HcHVpamDRijvpRzdHtvvOk%3D)

*The timeline shows the number and times of attacks in the last 24 hours.*

### Top Dest IPs

The Last 24 Hour Top Dest IPs charts show traffic for each the top three Destination IP addresses as evaluated by DDoS Defense alert policies for the following:

* Bits/s
* Packets/s
* Unique Src IPs

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(264).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A16Z&se=2026-05-12T09%3A35%3A16Z&sr=c&sp=r&sig=ewtz%2FmNP47hMhiEGtjTl1HcHVpamDRijvpRzdHtvvOk%3D)

*These charts show traffic over the last 24 hours expressed in the metrics used to evaluate attacks.*

### Attack Table

The Attacks Within Last 24 Hours table shows alarms in the last day from Protect alert policies. Each row provides summary information about an alarm (see **Alerts List**). Click a row for more details (see **Attack Details Drawer**).  
  
The **View More Attacks** button at the top right takes you to the **Alerting** page, filtering the **Alerts List** to show only DDoS attacks (not limited to the last 24 hours).

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(265).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A16Z&se=2026-05-12T09%3A35%3A16Z&sr=c&sp=r&sig=ewtz%2FmNP47hMhiEGtjTl1HcHVpamDRijvpRzdHtvvOk%3D)

*The table lists the alarms generated by alert policies over the last 24 hours.*

#### Attack Details Drawer

Click the row for an attack to open a Details drawer identical to the **Alert Details Drawer**.

## DDoS Defense Templates

The **Policy Templates** page offers DDoS templates designed to address various common DDoS attack types. To utilize a template, clone it and then modify it to suit your specific situation. Kentik provides templates for the following attack types:

* **Amplification Reflection Attack**: Amplification attacks exploit connectionless services that respond to requests with large (amplified) responses. They use a spoofed source IP to reflect these responses to the target's destination IP. This policy focuses on high-amplification services and looks for sudden increases in bits/second of traffic to a single destination IP. It also identifies the source port of the involved services.
* **Non-reflective DNS Flood**: These attacks overwhelm target DNS servers with a high volume of spoofed DNS request packets, preventing timely responses to legitimate requests. This policy, which should be filtered to the IPs of your monitored servers, looks for sudden increases in the number of requests per second to individual DNS servers.
* **ICMP Flood Attack**: These attacks overwhelm target hosts with a high rate of ICMP echo-request packets (e.g., echo or echo-reply). This policy, which should be filtered to the IPs of your servers and appliances that listen to ICMP, looks for sudden increases in the number of packets per second to individual hosts.
* **TCP SYN Flood**: These attacks overwhelm infrastructure components (e.g., load balancers, firewalls, Intrusion Prevention Systems, and servers running TCP services) by bombarding them with a high volume of traffic from spoofed sources. This forces new connections that fill up the session table. This policy looks for changes in bits/second against individual destination IPs.
* **TCP ACK and ACK/PSH Floods**: These attacks overwhelm servers running TCP services by sending a high volume of ACK/PSH flags from different IP addresses (often spoofed). This forces the target to search matching sessions in its table. This policy looks for changes in packets/second against individual destination IPs.
* **Distributed TCP RST & FIN Flood**: These attacks overwhelm servers running TCP services by sending a high volume of RST and/or FIN packets. This forces the target to look for matching sessions in its table. If the attackers spoof addresses that already exist in the session table, these sessions are lost. This policy monitors changes in packets/second against individual destination IPs.
* **UDP Fragments Attack**: These attacks send fraudulent UDP packets that lack source or destination port information. This diverts server resources toward trying to reassemble the packets. This policy primarily focuses on high volume in bytes. However, it also includes unique source IPs as a secondary metric and uses multiple policy thresholds to escalate the response based on the volume of the attack packets.
* **UDP Flood**: These attacks overwhelm servers running UDP services by sending a high volume of UDP packets, typically max-MTU to generate a high bitrate. This policy measures UDP packets by Bits/s.

  > **Note:** Since this policy overlaps with policies such as UDP Fragments and Amplification and Reflection Attack, it is typically considered a "catchall" policy for UDP, with higher thresholds.
* **Total Traffic Volumetric Attack**: This catch-all policy covers attack traffic that might be missed by more specific policies when hosts are flooded with a wide variety of volumetric traffic from different protocols and ports. The policy analyzes all traffic, categorized by protocol.

---

© 2014-25 Kentik
