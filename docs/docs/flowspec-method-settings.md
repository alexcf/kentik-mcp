# Flowspec Method Settings

Source: https://kb.kentik.com/docs/flowspec-method-settings

---

This article covers the specific configuration fields and settings required when building a Flowspec mitigation method from Kentik’s **Manage Mitigations**page**.**

> **Notes:**
>
> * For more on Flowspec, see **Flowspec Mitigation**and **Mitigation Platforms**.
> * For information on mitigation settings on the **General** tab, see **General Method Settings**.

Configuring a Flowspec-based method involves settings in the following two panes on the **Details** tab of a mitigation method dialog:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(675).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A31Z&se=2026-05-12T09%3A38%3A31Z&sr=c&sp=r&sig=frBm0ocoWbs1uJBraOy6oaNAKS8ss7cnhcl6RxVJv9Y%3D)**Traffic Matching**: Settings (covered in **Flowspec Condition Controls**) that define a flow specification by identifying traffic with specific characteristics (values or ranges of values for one or more component types).
* **Traffic Filtering Actions**: Settings (covered in **Flowspec Traffic Actions**) that specify the actions to take on the traffic subset defined in the **Traffic Matching** pane.

## Flowspec Component Types

A flow specification is a filter that matches traffic based on the values of Network Layer Reachability Information (NLRI) "component types" defined in **RFC 8955**, each of which represents a property of a packet (IP, ports, etc.).  
  
The **Traffic Matching** pane includes controls (condition groups; see **Flowspec Condition Controls**) that set the conditions for matching traffic based on the component types covered in the topics below.

> **Note:** Except as noted below, all component types support multiple conditions and nesting of condition groups.

### IP/CIDR Components

| Name | Description | Selection Method |
| --- | --- | --- |
| **Destination IP/CIDR** | Matches on a range of destination IP addresses. | Enter a destination IP/CIDR |
| **Source IP/CIDR** | Matches on a range of source IP addresses. | Enter a source IP/CIDR |

**Considerations**

* **CIDR Scope**: A lower CIDR means broader Flowspec actions; `/0` matches all traffic.
* If the **Infer From Alarm** switch is On for these component types:

  + The IP/CIDR field is locked.
  + **Automated mitigation**: Kentik will derive the IP from an alarm. Be cautious, as the inferred values may lead to very broad mitigation (see **Infer From Alarm**).
  + **Manual mitigation**: The user must enter the IP in the Start Manual Mitigation dialog (see **Start a Manual Mitigation**).
  + The mitigation is only available for assignment to an alert policy threshold if the policy's key definition includes the corresponding dimension (e.g., source or destination IP/CIDR).
* Do not enable the **Infer From Alarm** switch for both Source and Destination.
* These component types don't support multiple conditions or nested groups.

### Protocol and Port Components

| Name | Description | Selection Method |
| --- | --- | --- |
| **Protocols** | Matches on IP protocol, e.g., “UDP (17)”. | Enter a protocol number into the field at right. |
| **Source Ports** | Matches on source port. | Enter a port number into the value field at right. |
| **Destination Ports** | Matches on destination port. | Enter a port number into the value field at right. |

If the **Infer From Alarm** switch is On for these component types:

* **Automated mitigation**: the protocol or port will be derived by Kentik from the data provided by the alert (see **Infer From Alarm**).
* **Manual mitigation**: The method can't be used for manual mitigation.
* The mitigation is only available for assignment to an alert policy threshold if the policy's key definition includes the corresponding dimension (e.g., protocol or source/destination port).

### Additional Components

| Name | Description | Selection Method |
| --- | --- | --- |
| **ICMP Types** | Matches on the type field of an ICMP packet. | Choose an ICMP type from the drop-down value list at right (use the filter field at top to narrow the list). |
| **ICMP Codes** | Matches on the code field of an ICMP packet. | Enter an ICMP code into the value field at right. |
| **TCP Flags** | Matches on flags in the TCP header. | Select a TCP flag from the drop-down value list at right (use the filter field at top to narrow the list). |
| **Packet Lengths** | Matches on total IP packet length (excluding Layer 2 but including IP header). | Enter a packet length in bytes into the input field at right. |
| **DSCPs** | Matches on 6-bit DSCP field (Diffserv Code Point). | Select a DSCP value from the drop-down list at right (use the filter field at top to narrow the list). |
| **Fragments** | Matches on fragment status header. | Select a status value from the drop-down list at right (use the filter field at top to narrow the list). |

### Infer From Alarm

When you create an alert policy, you configure the conditions under which the system will generate an alert. When those conditions have been satisfied and the alert is triggered, the system can then use the data from the alert to construct the advertisement to send to the router (rather than having the user enter that information manually). Enabling the**Infer From Alarm** switch is what allows the system to do this.

## Flowspec Condition Controls

The controls of the **Traffic Matching** pane are structured as a series of rows that each represents a group of one or more conditions for a specific "component type" (see **Flowspec Component Types**). Each of these condition groups may be individually included or excluded from the Flowspec.  
  
Except as indicated in **Flowspec Component Types**, condition groups can support multiple conditions and may contain nested condition groups. When traffic is evaluated, the individual conditions in each nested condition group will be ANDed, the conditions and nested groups in each group will be ORed, and the main groups will be ANDed. The result is that only traffic matching the Flowspec defined by all specified condition groups will be affected by the actions specified in **Flowspec Traffic Actions**.

The controls for condition groups are largely the same, with some variation between component types as indicated below:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(676).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A31Z&se=2026-05-12T09%3A38%3A31Z&sr=c&sp=r&sig=frBm0ocoWbs1uJBraOy6oaNAKS8ss7cnhcl6RxVJv9Y%3D)**Enable/Disable**: To enable or disable a condition group, use the switch beside the group’s title.
* **Infer from Alarm**: If this switch is On, the remaining fields in the condition group will be locked and ignored. See **Infer From Alarm**.

  + The Protocols, Source Ports, or Destination Ports condition groups can't be used for manual mitigation.
* **Operator**: If the condition group is enabled and the **Infer from Alarm** switch is Off, choose an operator (e.g., equals, greater than) from the drop-down list at left.

  + Not included in the condition groups for source and destination IP/CIDR.
* **Value**: If the condition group is enabled and the Infer from Alarm switch is Off, a condition group will include one of the following:

  + Value field: Enter a value into the value field. Applies to the following component types: IP/CIDR (source and destination), Protocols, Ports (source and destination), ICMP codes, and Packet Lengths.
  + Value selector: Choose a value from the drop-down list at right (use the filter field at top to narrow the list). Applies to the component types not listed immediately above.
* **Remove**: To remove an individual condition, click the red **X** at the right of the condition.
* **Add Condition**: Add an individual condition to a conditions group. The condition will be ORed with other conditions in the group (match any).
* **Add Group**: Nest a conditions group within the top-level conditions group. The conditions in the nested group will be ANDed (match all).

  > **Note:** Not available for IP/CIDR (source or destination).

> **Note:** As with any powerful technique, Flowspec-based mitigation requires attention to detail and carries with it the risk of unintended results and adverse consequences. Before attempting to configure and deploy Flowspec, be sure that you fully understand the component-specific considerations noted in **Flowspec Component Types**.

## Flowspec Traffic Actions

Traffic actions are applied by a Flowspec receiver (routing system) to traffic that has been matched to the Flowspec defined in the **Traffic Matching**pane. The **Traffic Filtering Actions** pane contains the controls covered in the topics below.

### Traffic Action Setting

The following primary actions are available from the drop-down **Traffic Action** menu:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(677).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A31Z&se=2026-05-12T09%3A38%3A31Z&sr=c&sp=r&sig=frBm0ocoWbs1uJBraOy6oaNAKS8ss7cnhcl6RxVJv9Y%3D)**Rate Limit**: The Flowspec receiver will rate-limit matching traffic to the bytes/sec value entered into the input field that appears when you choose this menu item. Corresponds to the BGP Extended community ID 0x8006 (traffic-rate).
* **Discard**: The Flowspec receiver will discard matching packets (same as setting rate limit to 0). Corresponds to the BGP Extended community ID 0x8006 (traffic-rate).
* **Mark DSCP**: The Flowspec receiver will set the DSCP header of the matching packets to the Differentiated Services Code Point value entered into the input field that appears when you choose this menu item. Corresponds to the BGP Extended community ID 0x8009 (traffic-marking).
* **Route-Target Redirect**: The Flowspec receiver will assign to matching packets the MP-BGP Route-Target value entered into the input field that appears when you choose this menu item. This allows packets to be redirected to another VRF, where a different action may be applied (useful, for example, for DDoS scrubbing VRFs). Corresponds to the BGP Extended community ID 0x8008 (rt-redirect).
* **Next-Hop Redirect**: The Flowspec receiver will redirect all matching packets to the IP address entered into the input field that appears when you choose this menu item. Corresponds to the BGP Extended community ID 0x0800 (from **RFC 7153**).

### Sample Setting

As described in **RFC 8955 section 7.3**, the **Sample** setting "enables traffic sampling and logging for this flow specification." The implementation of this logging feature is vendor-specific, both in terms of the type of logging (typically syslog or equivalent) and the location where the log is kept (e.g., the syslog file/server [Juniper], a separate log specified with a “sample-log” CLI syntax [Cisco]). Consult your router vendor documentation for information about configuring the log destination.  
  
In a typical implementation, the receiver will begin to create log entries for packets that match the Flowspec. The log entries can be used (via your own log-reading system, not within Kentik), to accomplish the following:

* Check that Flowspec rules are being correctly applied (the right traffic is being matched and the right actions are being taken).
* Examine traffic of interest (e.g., abnormal traffic pattern suggestive of DDoS attack) for diagnosis, troubleshooting, etc.

If the **Sample** switch is On, log entries will be created for matching packets as follows:

* **If the volume of packets is below a router-configured threshold**: Log every packet matching the Flowspec.
* **If the volume of packets is above the threshold**: Log a subset of matching packets, sampled at a router-configured rate.

### Terminal Setting

The **Terminal** setting (based on **RFC 8955** section 7.3) applies when a given packet is matched by the rules of more than one mitigation method. It tells a Flowspec receiver how to proceed after a given action has been applied:

* **Terminal ON**: Continue to the next rule that applies to this packet.
* **Terminal OFF**: Skip subsequent rules, if any, and move directly to the next packet.

The follow sequence provides a simple example of how the Terminal setting works in practice:

1. An event occurs on the network that triggers Flowspec mitigation methods A, B, and C.
2. The Flowspec rules (traffic matching + action) for each method are broadcast by Kentik and received by the Flowspec receiver (router).
3. The router prioritizes the rules based on **RFC 8955** section 5, **Traffic Filtering**, determining that the processing order is A, B, C.
4. The router receives and evaluates packet 1, finding that it matches the traffic filter of all three rules.
5. The action of rule A is applied to packet 1.
6. The Terminal value of rule A is ON, so the action of rule B is applied to packet 1.
7. The Terminal value of rule B is OFF, so the action of rule C is not applied to packet 1 and the filtering engine begins its evaluation of packet 2.
