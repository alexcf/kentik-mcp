# Supported Mitigation Types

Source: https://kb.kentik.com/docs/supported-mitigation-types

---

The Kentik portal supports several industry-standard and proprietary protocols for mitigating DDoS attacks and traffic anomalies. The best protocol for your network depends on your hardware capabilities, your vendor relationships, and the specific nature of the attack traffic.

This article provides a high-level explanation of how each supported mitigation protocol functions within Kentik. For step-by-step configuration instructions, see the articles in the **Configuration Guides** folder.

## RTBH Mitigation

Remotely Triggered Black-Hole routing (RTBH) is an extremely powerful technique that network operators can use to protect their network infrastructure and their customers against Distributed Denial of Service (DDoS) attacks. By automating the redirection of undesirable traffic at discrete points in a network, RTBH gives operators the ability to mitigate DDoS and to enforce denylist/bogon routing policy from a centralized station.

While there are a variety of methods to implement RTBH, including RFC 5634 and source-based RTBH, most network engineers consider **destination-based RTBH** to be their best first-line defense against DDoS attacks. In a traditional, non-automated RTBH setup, a customer might call and say, “Help! I think I’m under attack.” An operator would then log into a route server and configure the mitigation, which is a /32 host “black hole” route. The route is then redistributed via BGP, along with a ‘no-export’ community and a next-hop address, to the routers where the attack traffic is entering the network. These routers then route the traffic to a destination that doesn’t exist (the black hole), for example a null interface.

Destination-based RTBH is not only incredibly effective at protecting network infrastructure, it’s also relatively simple to implement and to automate with Kentik alert policies that define any number of conditions that will trigger an alarm, and define as well as how the system should respond to an alarm depending on the specific situation. That response may be immediate, delayed, or manual; it may include notification and/or mitigation; and it may involve just a /32 host route or include an aggregate /24. The details of the mitigation are defined via the **Manage Mitigations** page. Once active, you can monitor the specific routes advertised by your RTBH mitigations on the the **BGP Announcements** page.

## Flowspec Mitigation

Kentik’s Flowspec-based mitigation is effectively an implementation of **IETF RFC 8955**. As a mitigation tool, Flowspec has the advantage of being far more precise than RTBH in two respects:

* Greater precision in defining which traffic is affected by a mitigation action.
* Greater range of possible mitigation actions.

RFC 8955 defines two complementary aspects of what is commonly referred to as a “Flowspec,” which are defined in the details tab of the settings dialog for a Flowspec mitigation method (see **Flowspec Method Details**).

* **Traffic matching**: A BGP Network Layer Reachability Information (BGP NLRI) encoding format is used to distribute a "flow specification" to a Flowspec receiver, which can be any routing system that supports MP-BGP. The flow specification is a filter that matches traffic based on the values of certain "component types", which are properties of packets (IP, ports, etc.; see **Flowspec Component Types**).
* **Traffic actions**: An Extended Community value is used to convey actions that a Flowspec receiver can take on packets that match the flow specification (see **Flowspec Traffic Actions**).

> **Note**: Active Flowspec announcements and their associated traffic actions can be monitored on the **BGP Announcements** page.

## Adaptive Flowspec Mitigation

Adaptive Flowspec is a dynamic extension of Kentik’s standard **Flowspec** mitigation capabilities. While traditional Flowspec relies on static traffic matching criteria, Adaptive Flowspec automatically analyzes real-time attack traffic and dynamically generates precise BGP Flowspec rules to neutralize threats as they evolve.

By utilizing customizable traffic analysis algorithms, such as identifying the top source or destination ports and IP prefixes (CIDRs), Adaptive Flowspec continuously adjusts its filtering rules to meet your predefined mitigation goals. For example, you can configure the system to block a specific percentage of attack traffic volume while strictly adhering to a maximum number of generated rules. If an attack vector shifts, the mitigation adapts alongside it. If the dynamic algorithms cannot safely satisfy the mitigation goal, the system can fall back to predefined failsafe rules to ensure your network remains protected.

To use Adaptive Flowspec with Kentik:

* Add an Adaptive Flowspec mitigation platform on the **Manage Mitigations** page.
* Add an Adaptive Flowspec mitigation method that defines your mitigation goals, traffic analysis algorithms, and failsafe modes (see **Adaptive Flowspec Method Settings**).

The step-by-step procedure for preparing your routers and deploying these dynamic defenses is outlined in **Configuring Adaptive Flowspec**.

> **Note:** Just like standard Flowspec, active Adaptive Flowspec announcements and their associated traffic actions can be monitored on the **BGP Announcements** page.

## Third-Party Mitigation

Kentik offers third-party integrations with cloud or data-center mitigation systems offered by leading vendors such as Radware, Cloudflare Magic Transit, and A10. Third-party mitigation is implemented via APIs that support third-party orchestration servers like Radware DefenseFlow as well as hardware scrubbing appliances like Radware DefensePro or A10 Thunder TPS. Third-party mitigations can be automated or activated manually.

To use a third-party mitigation system with Kentik:

* Add a new mitigation platform specifying the third-party system as the mitigation type (see **Add a Mitigation Platform**).
* Add a mitigation method to run on the platform (see **Add a Mitigation Method**).

The step-by-step procedure for adding these platforms is outlined in the **Configuring Third-Party Mitigations** guide.

> **Notes:**
>
> * In addition to configuring a mitigation platform and method in Kentik, you must also allow the following IP ranges on third-party mitigation platforms (e.g., Radware or A10) and on any devices used for Flowspec or RTBH mitigations:
>
>   + **US**:  `209.50.158.0/23 (IPv4)` and `2620:129:1::/48 (IPv6)`
>   + **EU**: `193.177.128.0/22 (IPv4)`
> * Cloudflare applies Magic Transit mitigation only when traffic exceeds minimums (100K pps for TCP/UDP; 60K pps for ICMP/GRE). Assigning Cloudflare MT to a Kentik alert below these thresholds may show mitigation as active when it isn't. For lower volumes, use an alternative mitigation platform (RTBH, Flowspec, etc.).
