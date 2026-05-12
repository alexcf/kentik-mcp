# Kentik for Azure

Source: https://kb.kentik.com/docs/kentik-for-azure

---

This guide provides instructions for integrating Kentik with Microsoft Azure.

![Diagram illustrating Azure subscriptions, roles, and data flow to Kentik account.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/Simple Azure Reference Architecture Diagram.jpg?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A01Z&se=2026-05-12T09%3A30%3A01Z&sr=c&sp=r&sig=alp1%2FArulOg1gbzRDdI5Z9qN5hxbOfeRN4ntnLJbagA%3D)

A reference architecture diagram showing Azure subscriptions, roles, and data flow to a Kentik account.

## Microsoft Flow Logging Recommendations

* Kentik supports processing both VNet (Virtual Network) flow logs and NSG (Network Security Group) flow logs from your Microsoft Azure storage account. However, Microsoft advises enabling only **one type of flow log at a time** to prevent duplicate traffic recording and additional costs (see **VNet Flow Logs Overview**).
* Kentik encourages all customers to use **Microsoft's migration scripts** to enable VNet flow logs. This ensures your Kentik account remains current and avoids potential issues.

## Process Overview

Integrating Azure with Kentik involves preparing your Azure environment and configuring the Kentik portal to enable the collection of metadata, **flow logs**, **firewall logs**, and metrics.

> **TIP:** See the **Cloud Overview** for an introduction to Kentik cloud setup.

We recommend following this process in order:

1. **Azure Prerequisites & Roles**: Verify you have the correct administrative permissions in Azure, choose your authentication method (Kentik Enterprise App vs. Custom App Registration), and gather your necessary Subscription and Resource Group IDs.
2. **Flow/Firewall Log Collection**: Configure your Azure environment to generate VNet/NSG flow logs and firewall logs, and export them to a designated Azure storage account.
3. **Kentik Export Configuration**: Use the Kentik portal UI to create a "cloud export" that authorizes Kentik to ingest the telemetry from your Azure storage account.

**Optional***:* If you prefer to deploy using infrastructure-as-code, see **Automated Configuration (Powershell)** to script the Azure storage and Kentik export setup simultaneously.

Once the setup is complete, you can use the Kentik portal to monitor your Azure network traffic, visualize resource utilization, and gain insights to optimize network performance and security.

Next, let’s move onto completing the tasks in **Azure Prerequisites & Roles**.
