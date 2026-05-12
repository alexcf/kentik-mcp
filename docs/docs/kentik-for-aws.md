# Kentik for AWS

Source: https://kb.kentik.com/docs/kentik-for-aws

---

This guide provides instructions for integrating Kentik with Amazon Web Services (AWS).

![Flowchart illustrating data export from AWS accounts to Kentik for analysis and processing.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/AWS-reference-architecture.jpg?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A09Z&se=2026-05-12T09%3A30%3A09Z&sr=c&sp=r&sig=egEDW9%2BQGm0RczrzccmsZa80Yrzy7IriVEYlsY1vbGs%3D)

*A Kentik-recommended architecture for collecting flow and metadata from your AWS accounts.*

> **Notes:**
>
> * See the **Cloud Overview** for an introduction to Kentik cloud setup.
> * For setup assistance, email **support@kentik.com** (see **Customer Care**).

## Process Overview

Integrating AWS with Kentik involves configuring the collection of three key data types: **Metadata**, **Flow/Firewall Logs**, and **CloudWatch Metrics** (see **About AWS Cloud Visibility**).

Follow these phases to complete the setup:

* **Metadata Configuration**: Select and configure an IAM architecture that enables Kentik to access your AWS metadata.
* **Flow/Firewall Log Collection Configuration**

  + **Configure Logging**: Enable flow logging for your target VPCs, Transit Gateways, and Network Firewalls and direct them to a dedicated S3 bucket.
  + **Grant Permissions**: Configure IAM roles and policies to allow Kentik to securely query AWS APIs and read your S3 buckets. You can choose a **Standard Configuration** (for single accounts) or a **Nested Configuration** (hub-and-spoke for large organizations).
* **Kentik Export Configuration**:

  + **Create Cloud Export**: Configure a new "cloud export" in the Kentik portal to ingest data from AWS.

    - **Link Resources**: Input your AWS Role ARN and S3 bucket details to establish the connection.
* **Verify & Monitor**:

  + **Validate**: Confirm that the cloud export status changes to “Active” (green checkmark) and data begins populating on the **Public Clouds** page of the Kentik portal.
  + **Visualize**: Use the **Kentik Map**, **Data Explorer**, and**Cloud Pathfinder** to monitor traffic patterns, resource utilization immediately. Gain insights for optimizing network performance and enhancing security monitoring

> **TIP**: Kentik recommends starting with a **Metadata-Only** cloud export to verify connectivity before configuring high-volume flow logs.

## About AWS Cloud Visibility

Kentik ingests three core telemetry types to provide complete visibility into your AWS environment:

* **Metadata (Context)**:

  Collected via AWS APIs, metadata is used to build topology maps (**Kentik Map**), enrich flow records, and power connectivity analysis (**Cloud Pathfinder**).

  1. **Core Infrastructure**: VPCs, Subnets, Availability Zones (AZs), and ENIs.
  2. **Routing & Security**: Route tables, Network ACLs, Security Groups, and Network Firewalls
  3. **Gateways**: Internet Gateways, NAT Gateways, and Transit Gateways (including attachments).
* **Flow & Firewall Logs (Traffic)**

  Flow logs provide the raw traffic telemetry needed for analytics, alerting, and security forensics. Used for traffic analytics in Kentik modules like **Data Explorer**, **Alerting**, and **Kentik Map**.

  1. **Collection Methods**: Logs are enabled on AWS resources (VPCs, Transit Gateways, Network Firewalls, etc.) and exported to an S3 bucket for Kentik to ingest.
* **Metrics (Performance)**

  Kentik collects CloudWatch metrics to track historical performance and health trends. Supported namespaces include:

  1. Load Balancing: `AWS/ELB`,`AWS/ApplicationELB`,`AWS/NetworkELB`,`AWS/GatewayELB`
  2. Connectivity: `AWS/VPN`,`AWS/DX` (Direct Connect), `AWS/TransitGateway`,`AWS/PrivateLinkEndpoints`,`AWS/PrivateLinkServices`
  3. Compute & Storage: `AWS/EC2`,`AWS/S3`,`AWS/NATGateway`
  4. Network Services: `AWS/Route53`,`AWS/ApiGateway`,`AWS/NetworkFirewall`,`AWS/NetworkManager`

> **Note**: Metadata and metrics collection can be configured independently of flow logs. For optimal cost and performance, many customers consolidate flow logs from multiple VPCs into centralized S3 buckets (e.g., one per region).

## AWS Logging Overview

Kentik can ingest flow logs from VPCs, Transit Gateways, and **Network Firewalls**. Together, these logs capture detailed information about IP traffic flowing across your instances, centralized routing hubs, and security perimeters.

Similar to NetFlow or sFlow in physical networks (see **About Flow**), this raw telemetry essential for visualizing traffic paths, auditing security, and troubleshooting connectivity issues.

Use the table below to determine which logging strategy best fits your visibility needs.

| Feature | VPC Flow Logs | Transit Gateway (TGW) Logs | Network Firewall Logs |
| --- | --- | --- | --- |
| Scope | **Granular**: Captures IP traffic for specific ENIs, Subnets, or entire VPCs. | **Aggregated**: Captures all IP traffic traversing the central Transit Gateway hub. | **Security-Focused**: Captures traffic evaluated by stateful firewall rules. |
| Visibility Level | Detailed view of internal VPC traffic (East-West) and security group efficacy. | High-level view of Inter-VPC and Hybrid (VPN/Direct Connect) traffic. | Deep visibility into allowed/dropped packets and intrusion attempts. |
| Best For | Deep-dive troubleshooting, security forensics, and monitoring specific workloads. | Simplified monitoring for complex, multi-VPC architectures. | Compliance, threat detection, and auditing firewall rule efficacy. |

> **Note:** Transit Gateway logs do not replace VPC flow logs. For comprehensive observability, Kentik recommends enabling **both**: use TGW logs for the "big picture" of your network and VPC logs for granular analysis of critical resources.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/AWS-Log_excerpt-136h674w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A09Z&se=2026-05-12T09%3A30%3A09Z&sr=c&sp=r&sig=egEDW9%2BQGm0RczrzccmsZa80Yrzy7IriVEYlsY1vbGs%3D)

*An excerpt from an AWS VPC flow log file.*

### Network Firewall Visibility in Kentik

When you configure AWS Network Firewall logs, Kentik seamlessly integrates that security telemetry into your existing workflows, allowing you to instantly see what traffic was blocked and exactly *why*.

You can leverage this data across several Kentik modules:

* **Kentik Map****:** Firewall instances appear natively on your topology map. Selecting a firewall reveals dedicated region and VPC panels where you can analyze allowed vs. blocked connections over time, traffic volumes, and a complete breakdown of rule hits.
* **Cloud Pathfinder****:** When tracing traffic paths through your AWS architecture, Pathfinder automatically includes your firewalls in the hop-by-hop path. If a connection is dropped, Pathfinder highlights the exact firewall and blocking rule responsible, providing the context you need to remediate the issue.
* **Data Explorer****:** Run deep-dive queries to inspect your traffic by specific firewall, policy, and rule to audit your security posture and threat landscape.

## AWS Documentation Reference

For deeper details on AWS logging mechanics, refer to:

* **Amazon VPC FAQs**
* **Creating and Configuring an S3 Bucket**
* **Flow Log Limitations**

> Notes:
>
> * Cloud export setup can also be initiated from the Welcome Page during Kentik onboarding.
> * Each VPC, Transit Gateway, subnet, or interface sending flow logs to a bucket is represented in Kentik as a "cloud device” (see **Cloud Exports and Devices**).
