# Agent Clustering

Source: https://kb.kentik.com/docs/agent-clustering

---

This article discusses the clustering of Kentik's **Universal Agents**.

> **Important**:
>
> * This feature is in **Early Access** to Kentik customers. To enable it, contact your Kentik account team.
> * Only agents with the **SNMP/ST capability** installed are eligible for clustering.
> * For optimal performance and reliability, clusters should consist of **agents located within the same region/site**. This ensures low-latency communication between nodes and consistent data reporting.

## Clustering Overview

Kentik allows you to group multiple Universal Agents into a single, resilient cluster, ensuring continuous data collection and simplified workload distribution.

Agent clustering provides two key benefits:

* **High Availability** (HA): In any cluster, one or more agents are automatically designated as a **standby** node. If one of the active agents in the cluster fails, the standby agent will automatically take over its collection tasks, preventing any loss of visibility.
* **Load Balancing**: The distribution of workloads within a cluster depends on the specific capability of the agents:

  + **Polling Tasks (SNMP/ST)**: For agents performing active polling, workloads are **automatically distributed** across all active agents within the cluster. This prevents any single agent from becoming overloaded and ensures optimal performance.
  + **Inbound Data (Flow Proxy)**: For agents that receive data from your network, load balancing is **not handled automatically** by Kentik. To distribute this traffic, customers must provide their our Load Balancer and Virtual IP (VIP) to front the agent cluster.

A cluster requires at least **two functional agents** to be considered healthy and **three functional agents** to provide failover capabilities.

> **Note**: For more on Universal Agents, see **Universal Agents**.

## Clustering Setup (Early Access Process)

While this feature is in Early Access, Kentik handles the creation and management of the Universal Agent cluster.

To create a cluster:

1. Provide your Kentik account team with a list of the Universal Agents you wish to group, and they will configure the cluster for you.
2. Once an agent is assigned to a cluster, it will no longer be available for individual device assignment in the UI. The cluster itself will function as the assignable target.

   1. UI visibility into cluster configuration is limited during the Early Access period.

## Manage Device Assignments

Managing device assignments to Universal Agent clusters conveniently uses the same Kentik portal workflow as assigning a device to a single agent**.** Follow the steps below to assign, move, and bulk-assign devices to clusters:

### Assign a Single Device to a Cluster

To assign a single device to a cluster of Universal Agents, follow these steps:

1. Go to Settings » **Devices** and locate the device in the list.
2. Click the pencil icon to open the device’s settings dialog and select the **SNMP** tab.
3. From the **Collection Agent** dropdown, choose a cluster where you would normally select an individual agent. The cluster appears as a single option in the list with the following attributes:

   1. **Name**: The name of the cluster, assigned by Kentik during setup.
   2. **Agent count**: The count of agents in the cluster, e.g., “3 agents”.
   3. **Status**: The overall status of the cluster, e.g., “Healthy”.

![Editing a device's settings in the Kentik portal to assign it to a cluster of Universal Agents.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/UA-assign-device-to-cluster.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A24Z&se=2026-05-12T09%3A33%3A24Z&sr=c&sp=r&sig=h7Do0q5mp8zxYvC6wJM%2BuUiXFbN8KpoO00LeiqXDMrI%3D)

*Edit a device's settings in the Kentik portal to assign it to a cluster of Universal Agents.*

### Move a Device Between Clusters

You can easily move a device from one cluster to another using the standard Kentik portal workflow.

1. Go to Settings » **Devices** and locate the device in the list.
2. Edit the device's settings on the **SNMP** tab to select the cluster as its new agent assignment (see **Assign a Single Device to a Cluster**).

### Bulk Assign Devices to a Cluster

Devices can be bulk-assigned to a cluster just as they can be assigned to individual agents, and the process in the Kentik portal follows the same pattern:

1. Go to Settings » **Devices** and locate the devices to be bulk-assigned.
2. Select the devices using the checkboxes in the far left column of the list.
3. From the bulk actions bar that overlays the list, click the **Agent** dropdown and select the desired cluster.
4. A green message at the top of the page confirms the updated device assignments.

![Bulk assignment of devices to a Universal Agent cluster using the Devices page of the Kentik portal.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/UA-assign-devices-to-cluster-bulk.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A24Z&se=2026-05-12T09%3A33%3A24Z&sr=c&sp=r&sig=h7Do0q5mp8zxYvC6wJM%2BuUiXFbN8KpoO00LeiqXDMrI%3D)

*Bulk assignment of devices to a Universal Agent cluster on the Devices page.*

> **Important**: Due to a known limitation, performance may be slow when attempting to bulk-assign **more than 20 devices** at once to a Universal Agent cluster. Kentik is working to improve the capacity of this feature.
