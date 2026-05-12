# Kentik Kubernetes Agents

Source: https://kb.kentik.com/docs/kentik-kubernetes-agents

---

> Kentik Kube
>
> If you don’t see Kentik Kube in your portal, it might not be enabled for your company’s account. Please contact your **Kentik Account Team**.

This article discusses the software agents that enable Kubernetes observability in Kentik portal modules such as **Kentik Kube****.**

![Deployment of Kentik agents in a Kubernetes cluster.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/KA-Kubernetes_cluster-586h1120w.svg?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A32Z&se=2026-05-12T09%3A32%3A32Z&sr=c&sp=r&sig=z%2Fso5GZejS6nfRzcnTIVHR%2F5NrEDWiJ%2Fpc3eyYer2qE%3D)

*Deployment of Kentik agents in a Kubernetes cluster.*

## About Kubernetes Agents

Kubernetes observability in the Kentik portal relies on multiple agents that are installed together on a Kubernetes cluster:

* **Kappa**: A lightweight eBPF agent that runs in a pod on each node in a given Kubernetes cluster. The data collected by this process-aware telemetry agent allows you to query, graph, and alert on conditions in that node. See **About Kappa**.
* **Kappa-agg**: Aggregates data on a cluster-wide basis from all of the kappa pods installed in a given cluster's individual nodes, and sends the aggregated data to Kentik.
* **Kubemeta**: A “watcher” object that collects Kubernetes object metadata from a cluster and sends it to Kentik. The metadata is used to populate the map on the portal's Kentik Kube page (If you don’t see Kentik Kube in your portal, it might not be enabled for your company’s account. Please contact your **Kentik Account Team**).

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(26).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A32Z&se=2026-05-12T09%3A32%3A32Z&sr=c&sp=r&sig=z%2Fso5GZejS6nfRzcnTIVHR%2F5NrEDWiJ%2Fpc3eyYer2qE%3D)

*Data from kappa and kubemeta enable detailed mapping of structure and traffic in Kentik Kube.*

## About Kappa

> Kentik Kube
>
> If you don’t see Kentik Kube in your portal, it might not be enabled for your company’s account. Please contact your **Kentik Account Team**.

The core agent for flow telemetry in Kentik Kube is kappa.

### Kappa Overview

Kappa is a host-based telemetry agent providing ultra-efficient observability across production settings including both on-premises data centers and cloud infrastructure. Kappa enhances your organization's ability to:

* Understand traffic flows from all vantage points, including containers and Kubernetes clusters.
* Find congestion and performance hotspots.
* Identify application dependencies.
* Perform network forensics.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(27).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A32Z&se=2026-05-12T09%3A32%3A32Z&sr=c&sp=r&sig=z%2Fso5GZejS6nfRzcnTIVHR%2F5NrEDWiJ%2Fpc3eyYer2qE%3D)

*Kappa enables Kentik portal modules to include traffic information involving containers and Kubernetes clusters.*

Kappa generates kernel flow data using **eBPF**, which allows Kentik to see the total traffic passing between any source and destination IP, port, and protocol across every conversation taking place within a host, cluster, or data center. Because this information is generated using the Linux kernel, kappa also reports performance characteristics, such as session latency and TCP retransmit statistics.  
  
Kappa also enriches these flow summaries with the application context. Kappa associates conversations with the process name, PID, and the command-line syntax used to launch the IP conversation. If the process is running inside of a container, the container name and ID are also associated. And if the container was scheduled by Kubernetes, kappa enriches the flow record with the name, namespace, workload, and node identifiers of the Kubernetes pod.  
  
Before exporting these records to Kentik, kappa looks for any records associated with other nodes in an environment and joins the source and destination traffic records together with the source and destination host, container, and process metadata, painting the complete picture of application communications within a data center.

> **Note:** Kappa does not “de-duplicate” records. If kappa produces a record for a given flow once at a source node and again at the destination node, the system will preserve and enrich both records. For an accurate volume count or data-rate metric when querying, include a filter whose dimension limits results to a single vantage point (source or destination).

### Kappa Advantages

Kappa's design and implementation enable significant advantages over existing approaches to host-based telemetry:

* **Easy, flexible deployment**: In contrast to application performance monitoring, kappa is flow-based instrumentation that runs with no changes to application code. That makes kappa easy to deploy, and enables you to start monitoring 100% of your on-premises and cloud infrastructure within minutes.
* **Versatile single agent**: Existing products designed to solve similar problems are typically narrow in scope. But kappa seamlessly covers — in a single agent — the full range of situations where you need to monitor traffic. Deploy kappa as a process or container onto bare-metal hosts, or directly into a Kubernetes cluster to monitor containerized cloud-native workloads.
* **Low resource utilization**: Kappa uses virtually no resources because it's built on eBPF, which allows sandboxed programs to run within the operating system kernel. Unlike other agent-based solutions, kappa can scale to persistent traffic throughput of 10 Gbps while consuming only a single core.
* **Built on the Kentik platform**: While open source tools typically rely upon difficult-to-scale backends, kappa sends its data to Kentik Data Engine (KDE), the high-availability cluster that powers the Kentik SaaS. At KDE ingest the kappa flow records are further enriched with data including Threat, Geolocation, and Internet context.
* **Advanced analytics included**: Kentik's portal and APIs offer a rich set of visualizations that help you understand your network's utilization, performance, vulnerabilities, and costs. Our proactive baselining also yields actionable insights and rapidly detects unusual or unexpected patterns in your traffic.

## Kubernetes Dimensions & Metrics

Kappa generates **Metrics & Dimensions** whose values are stored in the flow records of the Kentik Data Engine (KDE) and used for queries in the portal and in APIs, either as group-by and filtering parameters (for dimensions) or as the unit by which results are calculated, expressed, and ranked (for metrics).

![Dimensions whose values originate from kappa-collected data can be used in queries.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/KA-Portal_dimensions-354h845w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A32Z&se=2026-05-12T09%3A32%3A32Z&sr=c&sp=r&sig=z%2Fso5GZejS6nfRzcnTIVHR%2F5NrEDWiJ%2Fpc3eyYer2qE%3D)

*Dimensions whose values originate from kappa-collected data can be used in queries.*

#### Supported Dimensions

Kappa supports dimensions (see **About Dimensions**) that can be used in Kentik APIs or the Kentik portal for group-by and filtering. These dimensions are found in the portal under the category "Process-Aware Telemetry Agent" (see **PATA Dimensions** in the KB's Dimensions Reference).

#### Supported Metrics

Kappa supports the following metrics (see **About Metrics**), which can be used in Kentik APIs or the Kentik portal (see **Metrics in the Portal**):

* **Bits/s**: Bits per second.
* **Packet/s**: Packets per second.
* **Application Latency**: One-way network latency (in milliseconds) derived by examining request/response pairs at the application layer.

## Kubernetes Agents Deployment

Deployment of kappa for Kubernetes is covered here.

### Agents Node Environments

You can deploy Kentik's Kubernetes agents in all Linux nodes with a 5.6 kernel version. The agents may also work on some earlier kernels.

> **Note:** Any systems that have an XDP dataplane will not generate data.

### Agents Download and Install

The agents required for Kubernetes observability in Kentik are downloaded from the public kentik-kube-deploy repo in GitHub. For installation information, please see **https://github.com/kentik/kentik-kube-deploy#readme**.

---

© 2014-25 Kentik
