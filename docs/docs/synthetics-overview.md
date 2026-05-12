# Synthetics

Source: https://kb.kentik.com/docs/synthetics-overview

---

This article covers Synthetics section of the Kentik portal.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(490).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A05Z&se=2026-05-12T09%3A43%3A05Z&sr=c&sp=r&sig=c99eQfszMi1ihOjSSqXkcaDCG9lTLYMfqHieItzydSM%3D)

*Synthetics Dashboard provides a high-level overview of your organization's Synthetics testing.*

## About Synthetics

Kentik Synthetics enables you to continuously test network performance so you can uncover and correct issues before they impact customer experience. This proactive testing is enabled by Kentik software agents (see **About Synthetics Agents**). The public agents (available to all Kentik customers) that make up our Kentik Global Agent Network are located in all major Internet hubs and cloud regions (AWS, GCP, Azure, OCI, etc.). You can also deploy private agents (available only to your organization) in your physical infrastructure or your cloud resources.  
  
Ping and traceroute tests performed continuously with public and/or private agents generate key metrics (latency, jitter, and loss) that are evaluated for network health and performance. Additional types of testing available includes performance tests focused on HTTP, DNS, and full page load, as well as monitoring of BGP announcements and withdrawals.  
  
The **Test Control Center** displays test results graphically and allows you to configure each of the listed tests. Additionally, some metrics derived from synthetic monitoring will be displayed in other areas of the portal, such as **Kentik Map**, and integrated into **Insights**. The metrics from synthetic tests are in many cases correlated with metrics from real traffic data that comes from Kentik's flow and SNMP monitoring.

## Tests and Subtests

Kentik's Synthetics capabilities are designed to make it as fast and easy as possible to configure testing: you tell us at a high level what you are trying to test and then we largely handle the details of configuration automatically. With that in mind, we use the following terminology:

* **Subtest**: An individual recurring ping/traceroute combination in a single direction, meaning from one agent toward one other network entity such as an agent, IP address, or ASN.
* **Test**: A collection of subtests that are configured, run, and displayed as a single unit.
* **Test Preview**: A test that lets you preview the test results to confirm you’ve configured your settings successfully. A test preview does not consume any of your monthly test credits (see **Test Preview Overview**).

> **Note:** The consumption of credits for a given test is calculated based on a count of that test's subtests (see **Synthetics Test Credits**).

## Synthetics Test Types

Kentik's synthetic monitoring tests, which are managed in the **Test Control Center** in the Kentik portal (Synthetics » **Test Control Center**), fall into following main categories:

* **Routing Tests**
* **Network Tests**
* **Application Tests**

### Routing Tests

A BGP routing test monitors route announcements for potential hijacking. The following routing test is available:

* **BGP Monitor**: Based on user-specified sets of tested prefixes and "allowed" ASNs, check whether any prefix is originating from a non-allowed AS, which would possibly indicate route hijacking. If so, the route announcement that includes the hijack will be shown as a critical event in the **Incident Log** and as an Invalid Origin event (indicated with a red bar in the timeline) on the **Test Details Page** for the test. The BGP Monitor test also allows you to configure upstream leak detection for any listed ASNs (see Upstream Leak Detection under **BGP Monitoring**).

### Network Tests

Network tests monitor network connectivity performance between public or private Kentik agents, external servers, or specific test targets (e.g., ASN, CDN, Country, Region, or City). Kentik will automatically configure all of the required testing based on your from and to settings. Network tests fall into the following sub-categories:

* **Agent-to-Agent:** Test connectivity between public or private Kentik agents. The following Agent-to-Agent tests are available:

  + Agent-to-Agent: From one or more chosen agents to an individual agent.
  + Network Mesh: Bidirectional tests between agents in sites within your organization's infrastructure (private agents) and/or in Kentik's worldwide network of global agents (hosted and in public cloud).
* **Agent-to-Server**: Test connectivity between Kentik agents and external servers. The following test types are available for Agent-to-Server:

  + Server IP Address: From one or more chosen agents toward an individual IP address.
  + Server Hostname: From one or more chosen agents toward an individual hostname.
  + Network Grid: Ping and trace tests from selected private or global agents towards targeted IP addresses, which may be entered in a comma-separated list or added in bulk from a `.txt` or `.csv` file.
* **Autonomous Tests:** Performance monitoring that is guided by the actual traffic patterns on your network, which Kentik is able to do by analyzing your flow records in KDE along with correlated network traffic data (SNMP, BGP, GeoIP, etc.). You choose which of the below test types you'd like to use and choose the individual instance that will be the subject of the test, for example the ASN for an ASN test. We then analyze traffic patterns to or from your network that involve the test subject, and based on that analysis we intelligently select the IPs to test toward (typically the IPs your traffic is using most), so your test credits will be applied most efficiently. The following test types are currently available for autonomous testing:

  + ASN: Between one or more agents and a set of autonomously selected IPs in the specified AS.
  + CDN: Between one or more agents and a set of autonomously selected IPs associated with the specified CDN (see **CDN Attribution Dimensions**).
  + Country: Between one or more agents and a set of autonomously selected IPs in the specified Country.
  + Region: Between one or more agents and a set of autonomously selected IPs in the specified Region.
  + City: Between one or more agents and a set of autonomously selected IPs in the specified City.

### Application Tests

Application tests allow you to monitor reachability and performance for DNS servers, URLs, and HTTP resources. Application tests fall into the following sub-categories:

* **DNS Tests**: Test reachability and performance to DNS servers. The following DNS tests are available:

  + DNS Server Monitor: Test the performance of one or more DNS servers associated with a hostname, showing the resolution time and the resulting IP address.
  + DNS Server Grid: Test a hostname lookup simultaneously on multiple user-specified DNS servers, showing the resulting IP addresses and the resolution times, which enables you to identify servers that are non-performant.
* **HTTP**: Test reachability and performance to HTTPs. The following HTTP tests are available:

  + HTTP(S) or API: Test a specified web server's response to an HTTP GET request, or API endpoint to any HTTP Request Method. It shows the status code, the average time to last byte, and the response size (which enables you to spot anomalies in the amount of data returned). You can also optionally run ping tests towards the resolved IP address.
  + Page Load: Test a full browser page load using Headless Chromium run by Kentik app agents, showing the status code and the times for performance indicators such as navigation, domain lookup, connect, and response.
  + Transaction: Test a "transaction," which is a series of actions, running in Headless Chromium, that are driven by a **Google Puppeteer** script created in the Recorder tab of Chrome Developer Tools. Results include the health status and total transaction time for each agent you test from, and show any screenshots that you specified as actions in the script.

> **Notes:**
>
> * For information on configuring the tests listed above, see **Test Settings Page**.
> * The Synthetics module also reports BGP announcements and withdrawals (see **BGP Route Viewer Tab**) which, while not based on synthetic testing, enables you to correlate issues revealed by synthetic testing at the network (ping, trace) and web layers (HTTP, page load) with issues at the routing layer (BGP).

#### Supported Testing Configurations

Testing is available for configurations such as the following:

* From any Kentik global agents to external SaaS, application, and API endpoints.
* From any private agents to external SaaS, application, and API endpoints.
* From site (data center, remote office, campus, etc.) to site using private agents.
* Between sites (private agents) and public cloud (global agents).
* Between public cloud region instances using private agents.

## Synthetics Test Credits

The test credits used for synthetic testing in Kentik are covered in the following topics.

![Consumption of credits by your tests is indicated for both current and following month (projected).](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/TCC-Credit_utilization-91h1120w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A05Z&se=2026-05-12T09%3A43%3A05Z&sr=c&sp=r&sig=c99eQfszMi1ihOjSSqXkcaDCG9lTLYMfqHieItzydSM%3D)

*Consumption of credits by your tests is indicated for both current and following month (projected).*

### About Test Credits

Synthetic testing in Kentik is enabled by "test credits," which are purchased from Kentik in SyntheticPaks. The purchased credits add to your organization's test credit balance. Every time a test is run this balance is reduced by an amount that is calculated based on the factors explained in **Test Credit Consumption**. The results of these calculations are continually updated and displayed in the portal UI (see **Test Credit Display**) so you always know how many credits you have to work with and what the impact of any added testing would be on your test credit balance.

#### Adding Test Credits

If your organization's test credit balance drops below zero a notification will appear on some of the pages in the Synthetics section of the Kentik portal. Your currently running tests will not be paused or deleted, and your organization will not be automatically charged. You will, however, be unable to resume paused tests, add new tests, or edit any existing tests in a way that increases credit consumption. To remove these limitations, click the **submit a request** link in the notification, which will open the Add Credits dialog, whose fields will be populated with the information that Kentik will need to process your request (to contact us, see **Customer Care**).

### Test Credit Consumption

While the per-credit pricing of a SyntheticPak may vary depending on factors such as your Kentik platform (Pro or Premiere) and the number of credits in the SyntheticPak, the consumption of test credits in your organization varies by the type of test, as covered in the topics below.

> **Notes:**
>
> * Test Previews (see **Test Preview Overview**) do not consume any monthly synthetic credits until the test is saved. The Credit usage statement at the top right corner of the Preview screen informs you of how many credits the test would consume if saved with those settings.
> * The topics below apply to currently supported Synthetics tests. Additional test types added in the future may consume test credits in different ways and amounts.
> * For more a detailed breakdown of a test’s point usage, use the **Credit Consumption Dialog** on the test's Test Settings Page.

#### BGP Monitor Consumption

The BGP monitor test is not run using agents and its frequency is fixed at one minute. Kentik charges 0.53 credits for every minute that the test is configured and running.

#### Network Tests Consumption

Test credit consumption by a given network test (ping and traceroute) is always based on the following factors:

* The number of subtests (see **Tests and Subtests**).
* The type of the agent from which each subtest originates (see **About Synthetics Agents**):

  + Global: 5 credits per subtest (one direction).
  + Private: 2.5 credits per subtest (one direction).
* The testing interval (see **Test Information**).
* The testing duration. Once turned on, tests typically run continuously, so in a full 30-day month a test that is executed every minute will run 43,200 times (60 minutes × 24 hours × 30 days).

#### HTTP Tests Consumption

![Overview of credit consumption for HTTP and network testing with associated costs.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/SY-credit-consumption-dialog.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A05Z&se=2026-05-12T09%3A43%3A05Z&sr=c&sp=r&sig=c99eQfszMi1ihOjSSqXkcaDCG9lTLYMfqHieItzydSM%3D)Credits consumption in some HTTP tests (Page Load tests and HTTP(S) or API tests only) involve a timeout factor in which the timeout in seconds (see **HTTP Settings**) is multiplied by 0.5 for private agents or 1 for global agents to arrive at the credits per subtest. With the default timeout of 5 seconds, for example, the credits for a private agent subtest would be 2.5 (5 × 0.5), which is equivalent to the "network" tests described above.  
  
When used in conjunction with network testing (ping and traceroute), the credit consumption for network testing is added to that of the HTTP test. For example, if you’re testing two private agents against one target with HTTP and network testing, the breakdown would be as shown in the dialog pictured below.

> **Note**: The point usage breakdown shown above is available for each test in the **Credit Consumption Dialog** on the test's Test Settings Page.

### Credit Consumption Scenarios

To better understand how testing impacts your organization's test credit balance, let's apply the factors described in **Test Credit Consumption** to a couple of example scenarios. The first example is a simple mesh test (see **Synthetics Test Types**) executed every minute between four private agents within your network. In this mesh, the four “from” agents would each test toward the other three target agents, for a total of 12 unidirectional subtests. The following table shows the test credit consumption over one month for those subtests.

| Source agent | Target | Credits/subtest | Subtests/test | Credits/test | Tests/month | Credits/month |
| --- | --- | --- | --- | --- | --- | --- |
| Private | Private | 2.5 | 12 | 30 | -- | -- |
| Private | Global | 2.5 | 0 | 0 | -- | -- |
| Global | Private | 5 | 0 | 0 | -- | -- |
| Global | Global | 5 | 0 | 0 | -- | -- |
| **Total:** |  |  | 12 | 30 | 43,200 | 1,296,000 |

The calculation gets more complicated when considering tests involving different kinds of source agents. In the table below, for example, we see a test involving a mesh of five agents (two private and three global) which results in a total of 20 unidirectional subtests, executed every minute.

| Source agent | Target | Credits/subtest | Subtests/test | Credits/test | Tests/month | Credits/month |
| --- | --- | --- | --- | --- | --- | --- |
| Private | Private | 2.5 | 2 | 5 | -- | -- |
| Private | Global | 2.5 | 6 | 15 | -- | -- |
| Global | Private | 5 | 6 | 30 | -- | -- |
| Global | Global | 5 | 6 | 30 | -- | -- |
| **Total:** |  |  | 20 | 80 | 43,200 | 3,456,000 |

> **Note:** The above calculations are for illustration only. In actual practice all calculations are performed by the portal automatically and results are displayed as described in **Test Credit Display**.

### Test Credit Display

Your organization's test credit balance is displayed in various places in the portal’s Synthetics section, as is the impact on your credit balance of consumption by individual tests:

* **Licenses Page**: The Synthetics pane shows credits used so far this month as a percent of available credits. It also shows a count of credits used and credits available.
* **Synthetics Dashboard**: The Credit Utilization pane shows counts of credits available, credits used so far this month, and credits remaining for the month.
* **Test Control Center Page**: Cards across the top show current and projected Credit utilization. The **% Monthly Credits** column of the **Tests List** gives consumption this month as a percent of your organization’s total available credits.
* **Test Settings Page**: The expected monthly credits chart in the bottom left corner of the page shows the credits the test will consume per month and how many will remain. For more a detailed breakdown of a test’s point usage, click **[#] credit(s) per min** to open the **Credit Consumption Dialog**.

## Synthetics Test Status

The status assigned to synthetic tests is covered in the following topics.

### Test Status Levels

The results from most synthetic tests are evaluated in three areas (latency, packet loss, and jitter) on a continuous basis to assign the test one of the following statuses for each of those areas:

* **Healthy**: No issues have been found that would significantly impact performance.
* **Warning**: Issues have been found that could significantly impact performance if not addressed.
* **Critical**: Issues have been found that are significantly impacting performance and need immediate attention.
* **Pending**: The test exists but no results are available yet.

> **Notes:**
>
> * Each health status (latency, packet loss, and jitter) can be enabled/disabled individually, with a minimum of one enabled status per test (see **Health Threshold Controls**).
> * A test’s health status can be evaluated as a whole or per agent (see Status Calculations under **Health Settings**).
> * For more about how test data is evaluated for status, see **Latency and Jitter Status** and **Packet Loss Status**.
> * BGP Monitoring, DNS, and Transaction tests do not include the test status levels described above.

### Status Time and Scope

Each test defined in the Test Control Center runs at the configured testing frequency (available frequencies vary depending on test type). This testing frequency is independent of the time range (lookback duration) for which test results may be viewed, which may range from last 5 minutes to the last 30 days. The longer the time range, the wider the time over which data from testing will be aggregated to create a set of "time slice" values. The following examples illustrate how this works for various time ranges:

* **Last 1 hour**: If the time slice is one minute, the latency and jitter values for the current one-minute slice will be compared against the average of the corresponding values for the 15 preceding minutes.
* **Last 6 hours**: If the time slice is 5 minutes, the latency and jitter values for the current five-minute slice will be compared against the average of the corresponding values for the 15 preceding five-minute slices (last 75 minutes).
* **Last 1 day**: If the time slice is 60 minutes (one hour), the latency and jitter values for the current 60-minute slice will be compared against the average of the corresponding values for the 15 preceding 1-hour slices (last 15 hours).

> **Notes:**
>
> * The time range for an individual test may be set on its **Test Details Page**. On the Test Control Center page, the time range is always calculated as described for Last 1 Hour above.
> * The frequency of an individual test is set in the Test Information section of the **Test Settings Page** for that test.

### Latency and Jitter Status

For latency or jitter, the status calculation involves comparing the metric’s value in the current time slice to a baseline value representing the average for the prior 15 time slices. The basis of comparison (milliseconds, percentage, standard deviation, etc.) is chosen with the **Use** setting in the **Thresholds** control set of the test’s **Health Settings**.  
  
Assuming the default standard deviation settings, status is indicated as Healthy unless one of the following applies:

* **Warning**: The current value exceeds the baseline by a Kentik-calculated amount, and the difference is between 1.5 and 3 standard deviations.
* **Critical**: The current value exceeds the baseline by a Kentik-calculated amount, and the difference exceeds 3 standard deviations.

> **Notes:**
>
> * The default deviation values above can be changed.
> * Kentik uses a sliding scale to calculate the amounts above so that a large change in terms of standard deviation doesn't trigger a status of Warning or Critical when the amount is not large enough to realistically impact performance.

### Packet Loss Status

For packet loss, the status is characterized as Healthy unless one of the following applies:

* **Warning**: The current packet loss value is between 0 and 50 percent.
* **Critical**: The current packet loss value is greater than 50 percent.

> **Note**: The default percentage values above can be changed in the test’s **Health Settings**.

### Overall Status

In many places within our Synthetics UI, the status values for latency, packet loss, and jitter are presented individually; in other places, an overall status is presented. The overall status represents the worst status across the three areas. In other words, if the status of latency and packet loss are Healthy but the status of jitter is Warning, then the overall status will be shown as Warning.

## Synthetics Modules

The Synthetics section of the portal includes the following modules/workflows:

* **Synthetics Dashboard**: The **Synthetics Dashboard** provides a summary overview of the health status of your tests, as well as information on the agents deployed and the consumption of test credits.
* **Test Control Center**: The **Test Control Center** is used to manage your organization's synthetic tests. Cards across the top of the page show your current (month to date) and projected (full month, at current trends) consumption of test credits (purchased from Kentik in SyntheticPaks). A **Tests List** (filterable by name, type, label, or agent) provides information about your currently configured tests and enables you to drill down into the details of subtest results. You can also access an edit page for any existing test or use the **Add Test** button to go to the **Add Test** page, where you can choose and configure a new test.
* **Agent Management**: The **Agent Management** page lets you see and manage all currently deployed agents, global (including app agents and public cloud agents) and private (see **About Synthetics Agents**). The location of agents is shown on a map, below which is a filterable table providing information about each agent (including status) and allowing you to activate an agent, edit its configuration, or remove it. You can also add private agents on this page using the **Deploy Private Agent** pane across the top.
* **BGP Route Viewer**: The **BGP Route Viewer** shows BGP announcements and withdrawals, which, while not based on synthetic testing, enables you to correlate issues revealed by synthetic testing at the network (ping, trace) and web layers (HTTP, page load) with issues at the routing layer (BGP).
* **State of the Internet**: The **State of the Internet** module provides at-a-glance visibility — without consuming test credits — into the health, performance, and availability of common public applications, services, clouds, and networks that may impact your applications, networks, and services.

> **Note:** A user’s ability to view the above pages and to add/edit synthetic tests or agents depends on the permissions granted by their assigned Kentik roles (see **Kentik RBAC**).
