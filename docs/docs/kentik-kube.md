# Kentik Kube

Source: https://kb.kentik.com/docs/kentik-kube

---

> **Note**: If you don’t see Kentik Kube in your portal, it might not be enabled for your company’s account. Please contact your **Kentik Account Team**.

This article discusses the **Kentik Kube** page of the Kentik portal.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(159).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A51Z&se=2026-05-12T09%3A43%3A51Z&sr=c&sp=r&sig=dTlsVZYgIqNlkRVTbauy5ke6cVzDBPapOQs13y2KWH8%3D)

*The Kentik Kube landing page.*

## About Kentik Kube

> **Note**: If you don’t see Kentik Kube in your portal, it might not be enabled for your company’s account. Please contact your **Kentik Account Team**.

The **Kentik Kube** module provides detailed network traffic and performance observability that enables cloud and infrastructure engineers to quickly detect and solve network problems in the following areas:

* Inside and between Kubernetes clusters;
* Between Kubernetes clusters and network other resources (e.g. data center or cloud);
* Between Kubernetes clusters and the Internet in general.

Kentik Kube supports cloud-managed Kubernetes clusters (AKS, EKS and GKS) as well as self-managed on premises clusters using the most widely implemented network models.

> **Note:** For more information about Kubernetes, see **https://kubernetes.io**, including the following resources:
>
> * For a high-level explanation, start with the **Overview**.
> * For Kubernetes terminology, see **Glossary**.

### Kubernetes Entities

Kentik Kube enables you to drill down into the details of your Kubernetes resources at multiple levels:

* **Cluster**: A set of one or more nodes that run containerized applications.
* **Namespaces**: Groups of resources within clusters.
* **Node**: A machine, virtual or physical, that contains the services necessary to run pods.
* **Pod**: A group of one or more application-specific co-located containers that share storage and network resources.
* **Services and Deployments**: Applications and services which run as deployed groups of pods.

Kube provides detailed information about the above entities in two main locations:

* **Kube map**: A visual representation of the relationship between entities including clusters, namespaces, services, and deployments, as well as between those entities and the rest of your network (see **Kube Network Map**). The entities shown in the Kubernetes block of the Kube map (see **About Map Blocks**) vary depending on the type of page (see **Kube Module Page Types**).
* **Details drawer**: A drawer that slides out from the right of the page (see **Kentik Kube Details**). The drawer includes a collection of panes that give various types of information about each of the entities above, including the pods and nodes used for services and deployments.

### Kube Network Map

At core, Kentik Kube is similar to the **Kentik Map**, in that it features an interactive map that diagrams the structure of your organization's network. In this case the focus is on your Kubernetes infrastructure. The map illustrates the relationship between your Kubernetes infrastructure and the other main aspects of your network, each of which is represented as a "block".

### Kentik Kube Agents

Kentik Kube relies on multiple agents (kappa, kappa-agg, kubemeta) that are installed together on a Kubernetes cluster. For further information about these agents, including deployment, see **Kentik Kubernetes Agents**.

## Kentik Kube Module

> **Note**: If you don’t see Kentik Kube in your portal, it might not be enabled for your company’s account. Please contact your Kentik Account Team.

### Kube Module Page Types

The Kube module includes the following types of pages:

* **Kentik Kube**: The module's landing page, on which the Kubernetes block shows your Kubernetes clusters.
* **Kube Cluster**: A page for an individual cluster, on which the Kubernetes block shows the cluster's namespaces.
* **Kube Namespace**: A page for an individual namespace within a cluster, on which the Kubernetes block shows the namespace's services and deployments.

The UI of these various page types is very similar, including the fact that all of them include the three main block types described in **About Map Blocks**. The elements of the Kentik Kube page are described in the following topic, followed by topics detailing the differences for Clusters and Namespaces pages.

### Kentik Kube Page UI

The Kentik Kube page contains the following main elements:

* **Details** (in the subnav): Toggle visibility of the right-side Details drawer, which contains details about the currently selected Kubernetes element (see **About Kube Details**).

  > **Note:** This button is inactive when no individual elements are selected. To select an element, click it and choose **Show Details** from the popup.
* **Filters**: A button that opens a popup showing the filters applied to the data displayed on this page:

  + If no filters are currently applied you can add a filter by clicking the **Edit Filters** button, which opens the **Filtering Options Dialog**.
  + If any filters are currently applied, each will be represented as a card in the popup. You can remove a filter by clicking the red X at the right of its card, or you can modify or add filters by clicking the **Edit** **Filters** button.
* **Time range**: A control that indicates the current time range of the data displayed on this page and pops up a calendar form enabling you to specify the time range. This control is the same as the **Time Range Control** in the Cloud Performance Monitor.
* **Kube indicators**: A set of indicators found along the top of the diagram to the right of the title. The indicators provide current information about the number and status of clusters and their component parts (see **Kube Indicators**).
* **Kube map**: The main display area of the page, which shows a map of the main blocks and their contents (see **Kube Map**).

#### Kube Indicators

The following indicators are displayed above the map on the Kentik Kube page:

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(160).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A51Z&se=2026-05-12T09%3A43%3A51Z&sr=c&sp=r&sig=dTlsVZYgIqNlkRVTbauy5ke6cVzDBPapOQs13y2KWH8%3D)

* **Issues indicator** (warning icon): A button indicating the number (if any) of high latency pods and nodes. Click to open the **Kubernetes Problems** dialog for a list.
* **Pods:** The total number of pods in your organization's Kentik-connected Kubernetes clusters.
* **Nodes:** The total number of nodes in your organization's Kentik-connected Kubernetes clusters.
* **Clusters:** The total number of your organization's Kentik-connected Kubernetes clusters.

#### Kubernetes Problems

This dialog features a table listing pods and nodes with high latency. The table includes the following columns:

* **Name**: The name of the pod or node experiencing high latency. The icon for a pod or node is shown at the left.
* **Latency**: Delay in milliseconds between when a TCP connect was issued (measured in-kernel) to when a response packet was received for this connection.
* **Namespace**: A link to the namespace page in Kentik Kube for the pod or node.

  ![Kubernetes dashboard showing high latency issues for various pods requiring immediate investigation.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/KK-kubernetes-problems-dialog.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A51Z&se=2026-05-12T09%3A43%3A51Z&sr=c&sp=r&sig=dTlsVZYgIqNlkRVTbauy5ke6cVzDBPapOQs13y2KWH8%3D)

  *The Kubernetes Problems dialog alerts you to urgent issues with pods and nodes.*

### Kube Cluster Page UI

A Cluster page in Kentik Kube includes the same basic UI as the Kentik Kube page but with the following differences:

* **Kube Indicators**: These indicators are not included on a Cluster page.
* **Cluster indicator**: A box between the **Other Infrastructure** block and the **Internet** block that indicates the cluster that the page is about. Click the **Change** link to return to the Kentik Kube page, where you can choose a different cluster.
* **Related Clusters**: An inset at the top left of the Kubernetes block that lists clusters that have traffic going to and/or from any namespace listed in the **Kubernetes** block. When connections are shown (see **Kube Entity Popup**) the circle at the right of any listed cluster that has a connection will be blue.

### Kube Namespace Page UI

A Namespace page in Kentik Kube includes the same basic UI as the Kentik Kube page but with the following differences:

* **Kube Indicators**: These indicators are not included on a Namespace page.
* **Location indicator**: Two boxes between the **Other Infrastructure** block and the **Internet** block:

  + Cluster indicator: Indicates the cluster that the page is about. Click the **Change** link to return to the Kentik Kube page, where you can choose a different cluster.
  + Namespace indicator: Indicates the namespace that the page is about. Click the **Change** link to return to the namespace's Cluster page, where you can choose a different namespace.
* **Related Clusters**: An inset at the top left of the Kubernetes block that lists clusters that have traffic going to and/or from any service or deployment shown in the **Kubernetes** block. When connections are shown (see **Kube Entity Popup**) the circle at the right of any listed cluster that has a connection will be blue.

## Kube Map

> **Note:** If you don’t see Kentik Kube in your portal, it might not be enabled for your company’s account. Please contact your **Kentik Account Team**.

### Kube Map Parts

The Kube map is made up of the following main types of parts:

* **Blocks**: Gray rectangles that each enclose different buckets of infrastructure (see **Kube Map Blocks**).
* **Connections**: Blue lines that represent the connections between blocks and between network entities (e.g., clusters) within blocks. Each connection is made up of two segments, which each have an arrow at one end representing the direction of the traffic. When you hover over a connection the volume of traffic in each direction is displayed over the corresponding segment.

  ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(162).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A51Z&se=2026-05-12T09%3A43%3A51Z&sr=c&sp=r&sig=dTlsVZYgIqNlkRVTbauy5ke6cVzDBPapOQs13y2KWH8%3D)

### Kube Entity Popup

When you click an entity in the Kube map, a popup opens that gives information and lists possible actions (see **Kube Entity Actions**).

The popup includes the following types of information:![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(163).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A51Z&se=2026-05-12T09%3A43%3A51Z&sr=c&sp=r&sig=dTlsVZYgIqNlkRVTbauy5ke6cVzDBPapOQs13y2KWH8%3D)

* **Entity type**: The type of entity, e.g., cluster, namespace, service, or deployment.
* **Entity name**: The name of the entity.
* **Health status**: The health of the entity, either Healthy (green lozenge), Critical (red lozenge), or empty (no lozenge). An entity is deemed healthy when its latency is below 300ms.

### Kube Entity Actions

The following entity actions, which vary depending on the entity type, are available from the **Kube Entity Popup**.

* **Show Details**: Open the Kentik Kube Details drawer (see **Kentik Kube Details**) with details for the entity.
* **Show Connections**: Draw connections (see **Kube Map Parts**) between the entity and other entities on the map.
* **View Namespaces** (clusters only): Go to the Clusters page for this cluster, where the **Kubernetes** block will show the cluster's namespaces.

  > **Note:** This item is not active when the cluster has no namespaces.
* **View Topology** (namespaces only): Go to the Namespaces page for this namespace, where the Kubernetes block will show the namespace's services and deployments.

  > **Note:** This item is not active when the namespace has no services or deployments.

## Kube Map Blocks

The Kube map includes three blocks (rectangles with gray background) that illustrate the relationship between your Kubernetes cluster and the other main aspects of your network:

* **Other Infrastructure**: Includes any of your organization's non-Kubernetes resources (data center or cloud) with which your Kubernetes clusters exchange network traffic. See **Other Infrastructure Block**.
* **Internet**: The external sources and destinations of traffic to and from your network, broken down by Origin Network. See **Internet Block**.
* **Kubernetes**: Includes your Kubernetes infrastructure (see **Kubernetes Block**). The specific contents vary depending on the current level:

  + **Kentik Kube page**: The block shows your Kubernetes clusters.
  + **Kube Cluster page**: The block shows the namespaces within an individual cluster.
  + **Kube Namespace page**: The block shows the services and deployments within an individual namespace in a cluster.

The above blocks contain individual elements that you can click on to access further details about the network's structure and traffic. As an infrastructure engineer, this enables you to better understand what's happening in your Kubernetes cluster in real time. You can compare traffic for specific network entities and view data quickly without having to run queries.  
  
Traffic between the **Kubernetes** block and the other blocks is shown with gray lines that indicate the volume of traffic in each direction. Click on a line to open a Traffic Details drawer (see **Traffic Details**).

### Other Infrastructure Block

The **Other Infrastructure** block includes any of your organization's non-Kubernetes resources with which your Kubernetes clusters exchange network traffic. This includes physical infrastructure such as data centers, as well as resources in a public cloud. It contains the following tabs:

* **Services**: A list of network-based services or applications that are monitored within your organization. Each item in the list gives the service's name followed by its portnum/protocol (in parentheses).
* **IPs**: A list of IP addresses for resources that exchange network traffic with your Kubernetes clusters. Each IP is followed by the reverse DNS name of the IP address (in parentheses).

Clicking on an item in the **Services** or **IPs** list pops up a menu from which you can choose **Show Details** or **Show Connections** (see **Kube Entity Popup**).

> **Note:** This block can be toggled between expanded (default) and collapsed with the icon in the upper right corner.

### Internet Block

The **Internet** block lists the external sources and destinations of traffic to and from your network broken down by origin network. Each item in the list represents traffic to and from one origin AS. Clicking on an item in the list pops up a menu from which you can choose **Show Details** or **Show Connections** (see **Kube Entity Popup**).

> **Note:** This block can be toggled between expanded (default) and collapsed with the icon in the upper right corner.

### Kubernetes Block

The **Kubernetes** block provides visibility into your Kubernetes clusters and their respective namespaces, services, and deployments. As described in the topics below, the contents of the block vary depending on the type of the current page (see **Kube Module Page Types**).

#### Kubernetes on Kentik Kube

On the module's landing page, the Kubernetes block shows your Kubernetes clusters on a map with clusters grouped in circles by region. A cluster that's experiencing latency over 300ms will be rendered in red; all others will be rendered in blue. Click on a cluster to open the **Kube Entity Popup**, which gives you information about the cluster and lists possible actions.

#### Kubernetes on Cluster Page

The page for an individual cluster is reached by choosing **View Namespaces** (see **Kube Entity Actions**) in the popup that opens from a cluster on the Kentik Kube page. The Kubernetes block on this page shows the cluster's namespaces. A namespace that's experiencing latency over 300ms will be rendered in red; all others will be rendered in blue. Click on a namespace to open the **Kube Entity Popup**, which gives you information about the namespace and lists possible actions.

#### Kubernetes on Namespace Page

The page for an individual namespace within a cluster is reached by choosing **View Topology** (see **Kube Entity Actions**) in the popup that opens from a namespace on a Cluster page. The Kubernetes block on this page shows the namespace's services and deployments. A service or deployment that's experiencing latency over 300ms will be rendered in red; all others will be rendered in blue. Click on a service or deployment to open the **Kube Entity Popup**, which gives you information about the entity and lists possible actions.

## Kentik Kube Details

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(164).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A51Z&se=2026-05-12T09%3A43%3A51Z&sr=c&sp=r&sig=dTlsVZYgIqNlkRVTbauy5ke6cVzDBPapOQs13y2KWH8%3D)

The Details drawer opens from the right of the page when you choose **Show Details** (see **Kube Entity Actions**) from the **Kube Entity Popup**. The drawer includes a set of panes that  display information that you clicked on to open the popup.  
  
The information displayed in a Details drawer varies depending on the type of the element from which it is opened:

* **Traffic details**: Information about traffic between two entities, opened from a connection (link line); see **Traffic Details**.
* **Entity details**: Information pertaining to an individual entity; see **Entity Details**.

Details are available for the following entity types:

* **Other infrastructure block**: Services and IPs
* **Internet block**: Origin networks
* **Kubernetes block**:

  + Clusters: On the Kentik Kube page.
  + Namespaces: On a Cluster page.
  + Services and Deployments: On a Namespace page.

> **Note:** To close the drawer, click anywhere outside of it, on the **X**at the upper right, or on the **Details** button in the subnav.

### Traffic Details

The Details drawer for connections includes the following main parts:

* **Details type**: Located at the top left of the drawer, this "Traffic Details" label indicates that the details displayed in the drawer are for a connection.
* **Link entities**: Also at top left, this "Traffic between…" field indicates the two entities at either end of the link.
* **Details panes**: A set of one or more panes with graphs and visualizations showing the results returned from queries that are automatically run on the link being detailed (see **Details Panes**).

  > **Note:** The results reflect Kentik Kube’s current time range and filters settings.

### Entity Details

The Details drawer for entities includes the following main parts:

* **Entity type**: Located at the top left of the drawer, this field indicates the type of entity for which details are displayed in the drawer (e.g., service, cluster, namespace, etc.; see the list in **About Kube Details**).
* **Name**: The name of the entity for which the drawer is showing details.
* **Entity-specific details**: Additional information that varies depending on the type of the entity (see **Entity-specific Details**).
* **Charts**: A tab containing details panes for the entity.
* **YAML** (namespace, service, and deployment details only): A tab giving configuration YAML for the entity.
* **Details panes**: A set of one or more panes with graphs and visualizations showing the results returned from queries that are automatically run on the entity being detailed (see **Details Panes**).

  > **Note:** The results reflect Kentik Kube’s current time range and filters settings.

#### Entity-specific Details

The Details drawer for the following entities includes additional elements that are exclusive to those entities:

* **Origin networks** (ASN Details): A **View Peering Analysis** link that opens the portal's **Potential Peer Page** for this AS.
* **Cluster**: A **Region** field giving the region in which the cluster is located.

## Details Panes

Every Details drawer includes at least one pane that displays the results of queries that Kentik runs regarding the entity or link that is the subject of the drawer. The panes included for a given entity depend on the entity type. The collection of available panes currently includes the following:

* **Traffic pane**: Ingress and egress for traffic on the selected network element or link (see **Traffic Pane**).
* **Pods pane** (service or deployment entities only): A table that lists all pods and the nodes on which they run (see **Pods Pane**).

  > **Note:** The Details drawer for a pod (see **Pod Details**) or a node (see **Node Details**) opens from the table in the Pods pane.
* **Performance pane**: Latency for pods or external traffic (see **Performance Pane**).

### Traffic Pane

The **Traffic** pane includes the following UI elements:

* **Expand/Collapse**: Toggles visibility between the title bar only and the full pane.
* **Open in modal** (diagonal arrows icon): Opens the pane in a modal.
* **Traffic selector**: A drop-down menu from which to choose the set of traffic that will be evaluated for the query whose results will be displayed in the pane. Options (e.g., All Traffic, pods, etc.) will vary depending on the entity type of the details drawer.
* **Dimension selector**: A drop-down menu from which to choose the **Kubernetes Dimensions & Metrics** that will be evaluated for the query whose results will be displayed in the pane.
* **Metric selector**: The metric used to quantify the query results shown in the pane.
* **Traffic charts**: Two time-series stacked charts (or one if traffic is set to All Traffic), the top for ingress and the bottom for egress, showing traffic over the time period specified by the time range control. The charts include the following elements:

  + **Heading**: Shows the direction and volume of the traffic to or from this entity.
  + **Time scale** (bottom chart only): A time scale representing the time period specified by the time range control.
  + **Time-point details**: A popup, which opens when hovering over either chart, that shows values for the Total and Historical Total (7 days prior) at that point in the time range.
  + **View in Data Explorer** (on hover only): Opens a new browser tab to show the query in the portal's Data Explorer module.
* **Traffic table** (not shown for total traffic queries): A list of the top-X results returned from the current query. This table is similar to the traffic table in Data Explorer (see **Explorer Table Overview**).

> **Note:** The query results returned in this pane are affected by the settings in the Kentik Kube’s Time and Filter controls.    

### Pods Pane

The **Pods** pane is a table with the following columns:

* **Pod Name**: The name of the resource object in the cluster. Click the name to display the **Pod Details**.
* **Node**: The name of the node on which the pod is running. Click the name to display the **Node Details**.
* **Latency**: The average latency for the pod.
* **Average bit/s:** The average bit/s for the pod.

### Performance Pane

The **Performance** pane includes the following UI elements:

* **Expand/Collapse**: Toggles visibility between the title bar only and the full pane.
* **Open in modal** (diagonal arrows icon): Opens the pane in a modal.
* **Traffic selector**: A drop-down menu from which to select pods or external traffic.
* **Dimension selector**: A drop-down menu from which to choose the Kappa Dimensions and Metrics that will be evaluated for the query whose results will be displayed in the pane (see **Kubernetes Dimensions & Metrics**).
* **Metric selector**: A drop-down menu from which to choose the metric used to quantify the query results shown in the pane: Average, p98th, or Max latency. Latency is the delay in milliseconds between when a TCP connect was issued (measured in-kernel) to when a response packet was received for this connection.
* **Traffic chart**: A chart showing traffic over the time period specified by the time range control. The chart include the following elements:

  + **Heading**: Shows the volume of the traffic for this entity.
  + **Time scale**: A time scale representing the time period specified by the time range control.
  + **Time-point details**: A popup, which opens when hovering over the chart, that shows values for the Total and Historical Total (7 days prior) at that point in the time range.
  + **View in Data Explorer** (icon; on hover only): Opens a new browser tab to show the query in the portal's Data Explorer module.

### Node Details

Node details are displayed in the Details drawer when you click the name of a node listed in the table in the **Pods Pane** for a given service or deployment. To return to the sidebar contents for the service or deployment, click the back arrow at the upper left of the drawer.  
  
 When showing node details the drawer includes the following elements on the **Charts** tab:

* **Node information**: Fields with information about the node (see **Node Information Fields**).
* **Traffic pane**: See **Traffic Pane**.
* **Pods**: See **Pods Pane**.
* **Performance pane**: See **Performance Pane**.

Node details also include a **YAML** tab giving configuration YAML for the node.

#### Node Information Fields

The following information is provided in a set of fields at the top of the node details:

* **Node name**: The name of the node.
* **Service name**: The service running on the node.
* **Latency**: The average latency on the node.
* **Traffic**: The average traffic volume on the node.

### Pod Details

Pod details are displayed in the Details drawer when you click the name of a pod listed in the table in the **Pods Pane** for a given service or deployment. To return to the sidebar contents for the service or deployment, click the back arrow at the upper left of the drawer.  
  
 When showing pod details the drawer includes the following elements on the **Charts** tab:

* **Pod information**: Fields with information about the pod (see **Pod Information Fields**).
* **Traffic pane**: See **Traffic Pane**.
* **Performance pane**: See **Performance Pane**.

Pod details also include a **YAML** tab giving configuration YAML for the pod.

#### Pod Information Fields

The following information is provided in a set of fields at the top of the pod details:

* **Pod name**: The name of the pod.
* **Node name**: The node on which the pod is running.
* **Service name**: The service running on the pod.
* **Latency**: The average latency on the pod.
* **Traffic**: The average traffic volume on the pod.
