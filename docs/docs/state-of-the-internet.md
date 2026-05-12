# State of the Internet

Source: https://kb.kentik.com/docs/state-of-the-internet

---

This article covers the **State of the Internet** dashboard of the Kentik portal.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(500).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A21Z&se=2026-05-12T09%3A31%3A21Z&sr=c&sp=r&sig=cZDE6RrjRzG9fd%2FzouGAxhuMMmOWSoGlyk%2BeSGEQHaU%3D)

*The State of the Internet provides a high-level overview of synthetic test status.*

## State of the Internet Tabs

The **State of the Internet**module provides at-a-glance visibility — without consuming test credits —  into the health, performance, and availability of common public applications, services, clouds, and networks that may impact your applications, networks, and services. The module includes the following main tabs, which are provided to all Kentik customers without consuming test credits:

* **SaaS Applications**: This tab presents results from tests using agents in the Kentik Global Agent Network. See **SaaS Applications Tab**.
* **Public Clouds**: This tab uses Kentik-preset agent-to-agent meshes to give you a sense of where there may currently be issues affecting traffic to, from, or within the major cloud services (see **Public Clouds Tab**). These tests.
* **DNS Services**: Kentik global agents continuously make DNS queries to the most commonly used DNS service providers and to measure their availability and uptime, with the resulting health metrics being displayed on this tab (see **DNS Services Tab**). Additional providers can be added by customer request.

## SaaS Applications Tab

The **SaaS Applications** tab of the State of the Internet is covered in the following topics.

### About SaaS Apps Tests

The **SaaS Applications** tab presents status and performance information about common SaaS application services (Zoom, Office, Slack, etc.). The tests toward each SaaS service require no setup from your organization and do not consume any of your test credits.  
  
SaaS Application tests are performed by Kentik using a selected set of agents in our Kentik Global Agents Network. The results from a given test, which represent an aggregation of that test's subtests, are the same for all Kentik customers, meaning that they give a global sense of the health of the service. For insight into the specific performance of the service from a given location you can clone the test to make a new SaaS Application test for which you can specify the agents (e.g., your own private agents) with which to run tests toward that same service (see the Clone button in  **SaaS Applications List**).

### SaaS Applications List

The **SaaS Applications** list is a sortable table providing current and historical status information for the SaaS applications for which Kentik is currently running global performance tests. The table includes the following columns:

* **Status**: The current status of the test (see **Synthetics Test Status**).
* **Service**: The logo and name of the tested service.
* **Status Code**: The HTTP Return Code (see **Host Traffic Dimensions**) returned during the currently displayed time slice.

  > **Note:** A subtest for which multiple status codes were returned during the time slice will be represented in the **Test Details** table by multiple rows.
* **Response Size**: The average size, over the currently displayed time slice, of the HTTP responses returned to the source agent.
* **Domain Lookup Time**: From when the domain lookup starts to when the domain lookup is finished.
* **Connect Time**: From when the request to open a connection is sent to the network to when the connection is opened.
* **Response Time**: From when the first byte of the response is received to when the last byte of the response is received.
* **Avg HTTP Latency**: The average time, for HTTP requests made during the currently displayed time slice, from making the request to receiving the last byte of the response. This figure is the sum the columns Domain Lookup Time, Connect Time, and Response Time.
* **Avg Latency**: Latency on the tested path, averaged over the last hour (see **Status Time and Scope**).
* **Avg Jitter**: Jitter on the tested path averaged over the last hour.
* **Packet Loss**: The percent of packets lost on the tested path averaged over the last hour.
* **Customize** (Copy icon): Click to create a duplicate that will test toward the same service as this global test but from agents that you specify. You will be taken to a **Test Settings Page** that is already filled in with settings from the global test (e.g., agents to test from and name of test), which you can then customize. Once you save the revised settings the test will appear in the portal's lists of your organization's tests (e.g., in the Test Control Center's **Tests List**), will be functionally identical to your other internally created tests, and will consume your organization's test credits.

## Public Clouds Tab

The **Public Clouds** tab of the State of the Internet displays the current status of network performance between different regions of major public cloud providers. The status information — based on continuous (every minute) bi-directional testing between Kentik global agents installed in those public cloud regions — indicates where there may currently be issues affecting traffic to, from, or within the major cloud services. By default all of the available preset meshes are displayed in the **Public Clouds Meshes** pane. To change which meshes are displayed, click the **Show/Hide** button and use the checkboxes in the popup list.  
  
Each mesh is a matrix in which every agent is tested to and from every other agent. The display for each mesh is a grid showing the status of its component agents (see **Synthetics Test Status**). Inside each cell at the intersection of two agents there are three dots, which represent the latency, packet loss, and jitter status, respectively, for those two agents over the last 15 minutes.    
  
Hovering over a cell opens a popup with information about the subtest results in both directions (forward and reverse), including agent location, agent AS (name and number), and the latency, packet loss and jitter metrics for the current time slice. Clicking the **View Details** button takes you to the corresponding **Subtest Details Page**.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(501).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A21Z&se=2026-05-12T09%3A31%3A21Z&sr=c&sp=r&sig=cZDE6RrjRzG9fd%2FzouGAxhuMMmOWSoGlyk%2BeSGEQHaU%3D)

*Kentik provides built-in tests that show the performance of major cloud services.*

## DNS Services Tab

The DNS Services tab is based on information gathered by Kentik using continuous DNS queries from global agents to the most commonly used DNS service providers. These queries enable Kentik to measure the performance, availability, and uptime of these servers and to display the resulting metrics, giving Kentik customers real-time insight into the health status of the DNS services.  
  
The results are displayed as a grid in which each row represents an individual global agent that is querying multiple DNS servers, which are grouped into columns by DNS service provider. Each server within a service is represented as a square whose health status (Healthy, Warning, or Critical) is shown by color. Hovering over a server's square opens a popup with information about the subtest (see **Tests and Subtests**) toward this server from this row's agent.  
  
The informational popup for a given server includes the following UI elements:

* **Status**: The current health status of the DNS lookup from this agent using this server.
* **IP address**: The IP address of the DNS server.
* **Resolution time**: The time it currently takes to complete the response to the DNS query.
* **View Details**: A button that takes you to the **Subtest Details Page** for the subtest involving this agent and this DNS server.

> **Note:** The DNS Services tab includes a link enabling users to request that the tab include additional DNS servers. Click the link to open the portal's Give Feedback dialog, where you can enter the IP(s) and Hostname(s) you'd like to see added.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(502).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A21Z&se=2026-05-12T09%3A31%3A21Z&sr=c&sp=r&sig=cZDE6RrjRzG9fd%2FzouGAxhuMMmOWSoGlyk%2BeSGEQHaU%3D)

*A grid shows global agents as rows and DNS servers grouped into columns by DNS service.*

---

© 2014-25 Kentik
