# OTT Service Tracking

Source: https://kb.kentik.com/docs/ott-service-tracking

---

This article introduces the **OTT Service Tracking** workflow in the Kentik portal.

![The page provides an overview of OTT traffic as well as ranked lists of  Categories, Providers, and Services.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/OTT-main.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A52Z&se=2026-05-12T09%3A31%3A52Z&sr=c&sp=r&sig=7PuscrYwONEH5sdJ5TNBSrYt2D4ImN803Yc%2F09E8pZQ%3D)

*The page provides an overview of OTT traffic as well as ranked lists of*   
*Categories, Providers, and Services.*

## About OTT Service Tracking

|  |  |
| --- | --- |
| **Purpose:** | Associate traffic with OTT services and providers to understand the services used on your organization's network, correlate performance issues to services, and optimize pricing to reflect subscriber usage. |
| **Benefits:** | * Identify network traffic origins and growth; assess in-house offering competitiveness. * Analyze customer usage patterns to refine plans and pricing. * Enhance customer service by providing better information. |
| **Use Cases:** | * Detect and analyze “content events” to provide timely and efficient guidance for network operations teams. * Evaluate OTT utilization metrics for specific categories of users and delivery methods (e.g., CDNs, interconnection types, or PoPs). * Assess Mbps-per-subscriber for OTT service impact.Evaluate the ramifications of zero-rating specific content providers. * Consider zero-rating content provider implications. * Improve customer retention by identifying delivery issues. * Analyze suspicious traffic for legal liability. * Get alerts on OTT service performance issues. |
| **Relevant Roles:** | Network Engineers, Network Strategists, Marketing Leaders, Security Leaders, Executives |

The **OTT Service Tracking** workflow allows you to monitor traffic across different OTT service categories (e.g., video, gaming, social media), view top-X provider breakdowns, and explore detailed traffic for specific services and providers.

This is crucial for engineering leaders and network strategists at eyeball ISPs to ensure reliable, efficient, and cost-effective content delivery to subscribers.

## Kentik OTT Engine

OTT Service Tracking is powered by Kentik's True Origin engine, which identifies OTT traffic sources by evaluating and correlating traffic data. It categorizes traffic into three main dimensions: **OTT Service Name**, **OTT Service Category** (i.e., service type), **OTT Provider Name**, and **OTT Classification Type**.

While True Origin's deterministic classification (via DNS mapping) is highly accurate, a portion of traffic (often 10-20%) can remain unclassified due to factors like external DNS resolvers, simplistic devices, or DNS caching. To address this, Kentik utilizes **Intelligent Classification (IQC)**. This feature applies a statistical and probabilistic model based on known multi-service IP traffic distributions to infer and classify the remaining traffic, providing a more comprehensive view of your network's OTT utilization.

## OTT Classification Values

The **Classification** column indicates how well Kentik has identified OTT Services. The categories align with the **OTT Classification Type** dimension values available in Data Explorer:

* **Fully Classified:** OTT Service Name, OTT Service Category, and OTT Provider Name are identified directly (e.g., via a known IP or matching DNS lookup).
* **Provider-Only:** Only the OTT Provider is identified (Indirect provider-only), typically because multiple services share a hostname.
* **Pending Classification:** High-traffic hostnames without a matching OTT service pattern (Direct ad-hoc and Indirect ad-hoc), allowing True Origin to learn and classify more services over time.
* **Projected Fully Classified:** Traffic inferred via Intelligent Classification's probabilistic model. A service is statistically chosen for this flow based on known distributions, and that chosen service matches a known service pattern in Kentik's True Origin database.
* **Projected Ad-hoc:** Traffic inferred via Intelligent Classification's probabilistic model. A service is statistically chosen for this flow, but the domain name does not currently match a known service pattern in Kentik's True Origin database.
* **Unclassified:** Traffic meeting none of the above conditions, resulting in no identification of OTT Service Name, Type, or Provider due to lack of matching DNS or IP patterns.

## OTT Service Tracking Metrics

Metrics in the **OTT Service Tracking** module vary by context:

* Capacity tab values (e.g., traffic bitrates) are based on **SNMP OID Polling**, showing maximum bitrates during the highest traffic time range (see **Table Time-Slicing**).
* Other traffic values are derived from flow data (e.g., NetFlow), classified via DNS matching based on IPs (see **OTT Classification Values** and **Using True Origin**).

Next, let’s move onto **Getting Started**with OTT Service Tracking.
