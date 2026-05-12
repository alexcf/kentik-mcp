# Mitigation Overview

Source: https://kb.kentik.com/docs/mitigation-overview

---

This article provides an introduction to Kentik's **Mitigation** capabilities.

> **Notes:**
>
> * For configuring mitigation platforms and methods, see **Manage Mitigations**.
> * For more about active mitigations, see **Mitigations Page**.

![Table displaying alert details including status, ID, platform, and remaining time.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(334).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A35Z&se=2026-05-12T09%3A32%3A35Z&sr=c&sp=r&sig=s5yhUrhiNBiZDKPYloPcQnqBw4rkYQ5KPeqnPLteUiA%3D)

*An active mitigation on the Kentik portal's Mitigations page.*

Kentik’s mitigation system provides a powerful, flexible framework for responding to DDoS attacks and other network traffic anomalies. By integrating with your existing edge routers or third-party scrubbing services, Kentik can actively filter, drop, or reroute malicious traffic before it impacts your network performance.

### Core Mitigation Concepts

Mitigations in Kentik are built using two foundational, interconnected components:

* **Mitigation Platforms****:** The physical infrastructure or third-party service that executes the mitigation (e.g., your BGP-enabled routers, Radware, or Cloudflare).
* **Mitigation Methods****:** The logical rules, traffic thresholds, and BGP actions that dictate **how** the platform should handle the attack traffic.

#### The Platform/Method Relationship

Methods of the same type can be used in more than one platform of the same type. To better understand how method configuration differs from platform configuration, consider a scenario where your RTBH policies must be differentiated based on transit providers, interface capacities, or available peers. Being able to create multiple distinct methods and assign them to the same underlying platform enables you to precisely tailor your defenses for different deployment scenarios.

### Supported Protocols

Kentik supports several different mitigation protocols, allowing you to tailor your defense strategy to your specific network architecture. These include:

* **RTBH (Remotely Triggered Black Hole)**
* **Flowspec**
* **Adaptive Flowspec**
* **Third-Party Integrations**

  > **Note**: For a detailed explanation of how each protocol works, see **Supported Mitigation Types**.

### Mitigation Deployment

Once your platforms and methods are configured, mitigations can be deployed to your network in two ways:

1. **Automated Mitigation (Recommended):** The mitigation is assigned to a threshold in an alert policy and triggered automatically when network conditions match that threshold. Automated mitigations will dynamically escalate and de-escalate as changing conditions trigger different thresholds within the underlying policy.

   > **Note**: Automatic mitigations are available only in alert policies whose Dimensions setting includes source or destination IP/CIDR (see **Data Funneling**).
2. **Manual Mitigation****:** If you notice an anomaly that hasn't triggered an alarm, you can manually force a mitigation to start (unassociated with an alert policy threshold).

## Navigating Mitigation in Kentik

To build, deploy, and monitor mitigations in the Kentik portal, you will primarily use the following pages:

* **Manage Mitigations** (**Settings »****Mitigations**): The configuration hub where you define your mitigation platforms and methods.
* **Mitigations** (**Protect »** **Mitigations**): The monitoring dashboard where you view current and past mitigations, or start a manual mitigation.
* **BGP Announcements** (**Protect » Mitigations »** **BGP Announcements**): The routing ledger where you can view and manage the specific BGP routing announcements triggered by your active mitigations.

## Next Steps

For step-by-step instructions on configuring your preferred mitigation platforms and methods, select the appropriate guide from the **Configuration Guides** section in the sidebar.
