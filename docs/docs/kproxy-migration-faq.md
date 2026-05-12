# FAQ

Source: https://kb.kentik.com/docs/kproxy-migration-faq

---

The following questions relate to the migration from `kproxy` to the Universal Agent (see also: **Troubleshooting**).

---

### FAQ: Customer, Process, and Risk

What is the risk of downtime during this migration?

#### 

By following the "Warm Failover" strategy, the risk is minimal. Ensure the Universal Agent (UA) is "Up" and registered in the Kentik Portal before stopping the legacy kproxy service. The moment the legacy service stops, the UA can bind to the telemetry port, resulting in a near-instant cutover.

#### What is actually changing in my environment?

You are moving from a single-purpose, monolithic binary (`kproxy`) to a modular framework. This simplifies the security audit surface and reduces the number of disparate installation procedures you need to manage.

#### What stays the same?

* **Data Integrity**: Your historical data in the Kentik portal remains untouched.
* **Port Logic**: You can continue using your standard UDP ports (e.g., 9995 for Flow, 161/162 for SNMP).
* **Licensing**: Your existing Flow and SNMP entitlements carry over seamlessly to the UA.

#### What action is required on my network devices?

If the UA is installed on a new host IP, you must update the destination IP for NetFlow/sFlow/IPFIX exports on your routers and switches. If the UA is replacing kproxy on the same host, no device-side changes are required.

---

### FAQ: Technical & Architectural

#### Can one Universal Agent handle both Flow and SNMP simultaneously?

Yes. This is the **Scenario 2 (Enriched Flow)** deployment. The UA is designed to be multi-threaded and modular, allowing it to act as both a Flow Proxy and an SNMP Poller within a single engine. Both the Flow Proxy and SNMP Poller workloads need to be accounted for in the scaling for the system running the agent.

#### Why move to UA if my legacy `kproxy` is working fine?

Legacy `kproxy` is entering a maintenance phase. All future innovations, such as advanced integrations and some exciting DNS capabilities, are being built exclusively for the Universal Agent. Moving now ensures you aren't "locked out" of the roadmap.

#### What happens if the SNMP capability is "Down" but Flow is "Up"?

You will still receive flow records, but they will lack Enrichment. In the Kentik UI, you will see interface IDs (Index numbers) instead of human-readable names (e.g., `xe-0/0/1`) until the SNMP capability is restored.

#### Why does my new Universal Agent still appear on the legacy Kproxy Agents page?

Under the hood, the Universal Agent's Flow Proxy capability utilizes the same core engine as the legacy `kproxy` agent. Because of this shared architecture, your active UA flow deployments will automatically populate under **Settings »** **Kproxy Agents** as well as the **Universal Agents** page. This is completely normal and expected behavior.

---

### FAQ: Scalability & Performance

#### What are the hardware requirements for the UA as my primary agent?

Hardware requirements scale with throughput. For a standard deployment handling up to 50k Flows Per Second (FPS), we recommend 4 vCPUs and 8GB of RAM. For massive global deployments, the UA can be scaled horizontally.

#### How does the UA improve host-level resource efficiency?

Instead of running separate processes for flow collection and SNMP polling, which causes redundant metadata requests, the UA coordinates these tasks through a single engine. This reduces CPU context switching and memory footprint on the host.

New capabilities are also designed with a modern simple and largely stateless approach allowing for higher throughput as heavy tasks such as enrichment are performed later in the enrichment pipeline.

#### Can I run the Universal Agent in Kubernetes (K8s)?

Yes. The UA is fully container-ready. When deploying in K8s, ensure you use a Service of type LoadBalancer or NodePort to expose the UDP ports required for flow ingestion.

---

### FAQ: Security & Compliance

#### How does the UA communicate with the Kentik SaaS platform?

The agent establishes an encrypted outbound connection (TLS) to the Kentik cloud. No inbound connections from the internet to your data center are required, significantly reducing your attack surface.

#### Does the UA support SNMPv3 for secure polling?

Yes. The UA fully supports SNMPv3, including USM (User-based Security Model) with AES encryption and SHA authentication, ensuring that your network management traffic remains private.

#### How are agent updates handled?

The UA supports remote management via the Kentik Portal. You can trigger capability updates and agent version increments directly from the UI, ensuring your security patches are always current without requiring manual SSH access to every host.

---

### FAQ: Future-Proofing & Capabilities

#### Can the UA ingest more than just Flow and SNMP?

Yes. The architecture is designed for **Capabilities**, which are modular extensions to the core functionality for Universal Agent. Future additions will include support for synthetic agent monitoring and `ktranslate`/**Firehose** and a new integration-focused service, making UA the only agent you need for Network Telemetry, Automation/Enrichment and Operations.

FPS stands for flows per second, which is the rate at which flow metadata is sent to Kentik from a given customer. It directly correlates with the resources required to provide that customer with the Kentik service.
