# Release Notes 2020

Source: https://kb.kentik.com/docs/release-notes-2020

---

Kentik employs a continuous deployment methodology for constant extension and refinement of the Kentik v4 portal and the underlying Kentik platform. Release notes for each successive month of Kentik v4 updates are covered in the following topics:

* **December 2020**
* **November 2020**
* **October 2020**
* **September 2020**

> ***Note:*** For additional insight into what's new with Kentik, be sure to check the **Product Updates** page of our website (Platform » Product Updates).

## December 2020

Product Delivery Summary

* **Synthetics**:

  + New ksynth agent releases 0.0.10-rc1 and 0.0.10-rc-2.
  + Added cloud performance mesh presets -- pre-built zone-to-zone performance stats for major public cloud providers.
  + Traceroute now includes text output display.
  + Added features to prune path display and improve readability.
  + Added basic RBAC (view / edit tests, results) controls.
  + Added tagging capability for tests and agents (currently feature flagged, not generally available).
* **Kproxy**:

  + List of deployed kproxy agents, versions, status in UI.
  + Integrated into onboarding flow.
  + Per-company API credentials not tied to individual user.
* **MyKentik Portal v4**:

  + Ported existing functionality from v3.
  + Opened access to all companies, includes 50 tenants.
  + Improvements to landlord tenant configuration and view management UI.
  + Redesigned tenant UI.
* **PDF and email export**: Now available for selected workflows:

  + Network Explorer screens.
  + Capacity Planning.
  + Connectivity Costs.
  + CDN Analytics.
  + OTT Service Tracking.
  + Synthetics (feature flagged for now).
* Custom / local DNS lookup via kproxy now puts resolved hostnames into native columns instead of custom dimensions.
* Added details sidebar to improve UX for device and interface lists.
* Added the ability to override SNMP-learned values for interface IP addresses.

**Issues Addressed**

* Multiple fixes for exports in Synthetics
* 4673 Fixed inherit BGP to require a master BGP device
* 2135 Restricting percentile outsorts in pie and sunburst visualizations
* 5573 Fix for Synthetics policy status getting incorrectly set on updates
* 4857 Selecting provider tab in Hybrid Map with site selected fixed
* Fixed validation messages in HTTP and DNS Codes fields in Synthetics form
* Added internal debugging URL on API errors in portal
* Flow Tag, AS Group and Custom Dimension lookups now return up to 100 matches instead of 20
* 5574 Fixed errors in Synthetics Path visualization caused by undefined interface or geo info
* 5316 Interface IP/netmask overrides: we’re incorporating this into Q4 IC work so there’s a solution that encompasses all we’re trying to do with interface overrides
* 4622 Filtering devices by status: this is still in review
* 4503 Fixes for bidirectional generator views with filters metrics that match the group by dimension: this is still in review
* 4591 Better organization of device details: this was partially completed on bug day and has been continued to accommodate bulk device label editing in v4
* 5056 Performance meshes customize button disabled when there’s no meshes
* 5480 Fixed cloning of saas app tests
* 5584 Fixed exporting of generator and bidirectional charts
* 5521 Removing saas apps tabs in onprems
* 5703 Made dashboard query panel scrollable
* Fixed DDoS max minutes client-side validation
* 5689 Synth test activate settings couldn’t be updated from their initial values
* 5559 Fixed sorting by utilization in capacity plans
* 5597 Fixed “Show Top 50 Results” button functionality
* 5486 Synthetics Insight support added
* Fixed missing customize button in NE pages

## November 2020

**Product Delivery Summary**

* **Synthetics:**

  + Several ksynth dot releases to fix bugs and support new test functionality.
  + Support alerting of test health changes.
  + Support notifications from test health alerts.
  + Support customization of test health thresholds.
  + Support customization of test health status for HTTP & DNS response codes.
  + Display SNMP metrics for devices in the test result path view.
  + Support tags for global and private agents to make it easier to select groups of agents during test configuration.
  + Support PDF export of reports within Synthetics via UI download or email.
  + Public release of DNS & HTTP tests.
  + Support TCP-based traceroute.
* Support fixup to prevent double counting of sflow data when ingress + egress sampling is enabled on both upstream and downstream links.
* Support manual override of interface IPs learned from SNMP.

**Issues Addressed**

* 5212 Supporting IBM Cloud in Onboarding
* 5213 Added Plan ID to Cloud Detail page
* 3956 Added clone and delete to LIbrary page
* 4618 Fixed last datapoint for unique ASNs is always 0
* 4487 Improved exported tables with many columns
* 4882 rDNS lookup in Data Explorer respects preference in user profile
* 5003 Fixed issues with pausing and resume synthetic tests
* 3540, 3386, 4455, 5039, 4859, 3538 Hybrid Map link connection improvements
* 5200 Fixed connection between device to other sites in Hybrid Map
* 5164 Allowing BGP MD5 Password to be empty
* Column label change in DNS test results (Resolved IPs -> DNS Results)
* 4914 Improved copy in Kentik Firehose
* 5054 Fixed issues with selections in Interface Groups
* 5230 No BGP device IP is shown
* 4743 IC filters do not stay when changing filter type to source/dest
* 2111 MPLS Router dimensions not available
* 5299 Fixed disappearing device site role box
* 5339 Fixed tooltip show/hide loop in synthetic path viz
* 4918 Fixed issues with legacy saved filter format in user filters
* 5314 Overhauled device stat reporting in device settings page
* 4741 Fixed dashboard clone losing some settings
* Rolled back a fix that caused library clone/delete to not work
* Fixed issue with GCP onboarding
* 4736 Fixed cloud save to render updated plan usage info
* Fixed margin in synthetic timeline
* Fixed issues caused by AS prepending in AS Path tooltip
* Fixed issues caused by rDNS lookups in IP Address quick view links

## October 2020

**New Functionality**

The following new features and capabilities were added this month:

* **Demo Mode**: The Kentik portal's new Demo Mode enables you to test all of Kentik's functionality without first configuring data sources to send us your organization's own traffic data. Using network devices and data that are pre-generated by Kentik, Demo Mode also includes a collection of Situations, which walk you step-by-step through real-world troubleshooting scenarios. New customers can enter Demo Mode from an option on the onboarding intro page. Existing customers can enter Demo Mode from a new link in the main navigation menu.
* **New sFlow field**s (Advanced sFlow devices): Kentik now supports collection of TTL and VLAN ID fields from devices that send sFlow data. These fields appear in Data Explorer and elsewhere in the Kentik platform as new dimensions that can be grouped or filtered. To enable collection of these fields from a given device, set the Type field in the Device dialog (Settings » Networking Devices) to “Advanced sFlow.”
* **Nokia device interface metadata**: Kentik has changed how we map flow record interface IDs from Nokia (formerly Alcatel) routers to other interface metadata (including names and descriptions) in the SNMP MIB. By default these routers populate flow records with interface IDs from a Nokia-specific “global IF index” table, but they may alternatively be set, via the `use-vrtr-if-index` configuration command, to populate these IDs from the standard IF-MIB. Kentik previously required the alternate setting, but we now support only the default mode (global IF index). If flow data from these routers is sent via Kentik's kproxy agent, the agent must be v7.26 or higher to support this change.
* **Operate » Hybrid Map, aggregate interface views**: The Hybrid Map view now allows you to choose whether an aggregate interface (a.k.a. “LAG” or “bundle” interface) should be shown as a single interface or as separate component interfaces. When you click on the interface's link in the map, the resulting popover will now include a drop-down selector from which you can choose the aggregate or individual member interfaces. The graphs below will update accordingly.
* **Operate » Hybrid Map, traffic path visualization**: The Hybrid Map view can now overlay traffic path visualizations onto the topology view. A Filtering pane has been added to the Options sidebar to enable traffic to be filtered. Bold dashed lines on the topology map indicate the links and devices traversed by the matching traffic.
* **Synthetics » Test Control Center**: A new How does this test work link on each Test Settings Page pops up a detailed description and diagram of each test.
* **Synthetics » Test Control Center, Autonomous tests:** For tests in the Autonomous category, the UI now clearly indicates that the direction of the test will be from agents to the selected target (ASN, CND, City, Region, or Country), and the target selection UI now gives you more ways to specify a target:

  + Using the tabbed target table, you can (as before) enter the target manually (Manual tab) or pick the target from a list on the new Inbound or Outbound tabs. The lists are auto-generated based on top-X traffic volume.
  + Any targets in the Inbound or Outbound lists that are already the target of another test are indicated as "Monitored."
  + When you select a target in the Inbound or Outbound list a tabbed agents table will appear in the Select agents field of the UI.
  + The Inbound and Outbound tabs of the agents table will each contain an auto-generated list of agents that have traffic toward the target. The agents in each list will be grouped by site; if no agents are deployed in a given site then the Add an Agent button will be displayed for that site.
  + The Manual tab of the agents table will contain the Edit Agents button (in Common Test Settings), which opens the Select Agents Dialog.
* **Synthetics » Test Control Center, Advanced options**: The Advanced Options pane, used to change test settings from default values, has been redesigned for improved readability and in anticipation of upcoming additions and options.
* **Synthetics » Test Control Center, traffic charts in subtest details**: The Metrics tab of the Subtest Details Page for all test types except Site Mesh and Custom Mesh now include dynamic traffic charts based on flow records from your real traffic, both inbound and outbound, along the same path as the subtest. To ease correlation of synthetic test results with real traffic impact, the timeline of the traffic chart is aligned with the timelines for synthetic test metrics.
* **Synthetics » Test Control Center, test details page**: Test Details pages now include the following enhancements:

  + Info icons pop up explanations related to the displayed results, such as why status is indicated in gray (test details table) or why some agent traces are incomplete (Paths tab).
  + The heading block now includes a field giving the test type and the name of the target being tested (i.e. “IP address 1.1.1.1” or “Hostname sockshop.aws.kentik.io”).

**Issues Addressed**

The following issues were addressed this month:

* 677 Queries that have no devices will now show an error in dataview + cleaned up loading indicators, warnings and empty states in dataviews
* 2021 Improved category selection logic in global search
* 3050 Utilization reported as 0 across circuits in a cost group
* 3322 IP Address Quick View Network Classification incorrect
* 3359 Fixed issues with removing last interface in a cost group
* 3446 VPN Identifier values don’t render in Sankey
* 3637 Flagged by Company filter added to insights page
* 3652 VLAN values don’t render in Sankey
* 3789 Comparison insight highlighting things it shouldn’t
* 3940 Firefox specific saved view error
* 4006/2757 Added Synth Tests to global search
* 4011 Disable 2fa in user form needs UX work
* 4012 Autocomplete dropdowns don’t close when selecting a value
* 4127 CD Populators dialog remembers filter and doesn’t show all populators
* 4132 Updated mesh popover button text
* 4229 Keeping plan device counts in sync
* 4365 Fixed sample rate shown for devices
* 4562 Fixed classification filtering OTT landing page
* 4569 Long interface names not viewable
* 4672 Fixed links drawn through cloud assets
* 4703 Device Selector is laggy at times
* 4714 Hybrid popover has super short table
* 4731 Fixed issues with multiple cross-panel filters
* 4745 Fixed AWS subnet links from Hybrid Map
* 4760 Prevent duplicate aggregates in insight queries
* 4860 Allow reset of time range when chart is zoomed
* 4891 Fixed sorting by traffic in Interfaces page
* 4916 Device Labels page doesn’t scroll
* 4928 Fixed mismatched timezones in Synth timeline
* 4933 device\_metadata failing to update for dropboxdc
* 4984 Allow populator fields to be empty
* 4990 Don’t show capacity lines in insights with no capacity threshold
* 5002 Prevent changing protocols when editing Synth tests
* 5008 Fixed 500 returned on excluded interface group in DDoS
* Fixed enter key behavior in Select control
* Make Export ID more clear/visible in Cloud Export Detail page
* Added device name to interface search results

## September 2020

**New Functionality**

The following new features and capabilities were added this month:

* **Synthetics » Test Control Center**: The Path tab for test and subtest details (Test Details Tab Views and Subtest Details UI) now provides hop-by-hop traceroute visualization of the path between the test agent (source) and target (destination). Paths are color-coded to highlight BGP ASNs and links/nodes with high latency or loss. The Explorer view highlights AS Path changes.
* **Synthetics » Test Control Center**: Test settings pages now include the following additional Advanced Options to provide finer control over testing:

  + Test Frequency
  + Alerting conditions, including the number of unhealthy subtests required to report status as unhealthy.
  + Ping options, including number of probes, protocol, and timeout.
  + Traceroute options, including number of probes, protocol, port, timeout, and hops.
* **Synthetics » Test Control Center**: On the Mesh tab of a Test Mesh Details page the popup that appears when you hover over the cell for an individual subtest at the intersection of two agents now includes test results and path metrics in both directions (forward and reverse paths).
* **Operate » Hybrid Map**: Improved mapping for sFlow/VLAN interfaces now associates the physical and logical interfaces of sFlow devices, resulting in improved traffic visibility.
* **Operate » Hybrid Map**: The new Links pane on the Options sidebar lets you choose whether connections between elements in the topology should be drawn based on Layer 2, Layer 3, or both.

**Issues Addressed**

The following issues were addressed this month:

* 2292 - Show interface details instead of insight details for capacity insights
* 2312 - Fixed issue with viz depth control not updating when changing to/from table component in DE
* 2327 - DDoS reporting incorrect IC percentage
* 2445 - Allow flipping to/from UE dimension filters
* 2452 - Total Bytes metric not showing reverse DNS values
* 2457 - SNMP filters missing dropdowns (also fixed ST filters and SNMP/ST renderers)
* 2503 - Sudo shows archived devices against plan device limits
* 2895 - Prevent information leakage in Slack notification channels
* 2910 - Fixed greek prefixing when there is only overlay data
* 2973 - Let users know when cost data doesn’t reflect recent changes
* 2988 - Updated RPKI Analysis preset to exclude null results
* 3742 - Showing ack required color and fixing start time in insights table
* 4035 - Viz depth in dashboard panels set different than when you drill in
* 4326 - Fixed pak fps query to better line up with real metrics
* 4460 - Fixed Blizzard logo in OTT
* 4491 - Showing OnPrem Connected Interfaces in Hybrid Map
* 4492 - Insight chart shows no results
* 4504 - Active alarms shouldn’t show end time equal to start time
* 4529 - v4 interfaces page missing data on many interfaces
* 4582 - Fix DE legend table menu dismissing
* 4585 - Strange drop and final hour results (forcing minsPolling on chunked queries)
* 4588 - Fix error toasts in dashboard (due to using fixed time window)
* 4589 - Can’t clone existing panel into new dashboard
* 4621 - Fixed ddos config overflow
* 4657 - Fixed sort order in Network Explorer tables
* 4685 - Fixed dashboard overflow in Safari
* 4711, 4386 - Fixing directional issues in Hybrid Map popover queries
* 4718 - Filtering out connections with no traffic in Hybrid Map
* 4737 - Disallow moving devices to a Synth plan
* 4739 - Showing Plan FPS in Device row expander
* 4754 - Updated SNMP polling text
* 4767 - Fixed ddos landing card value for attacks in last 24h
* Fixed guided mode presentation in dashboard panel applied filters popover
* Fixed no flow warnings in settings pages when you have devices
* Fixed time range input overflow in date range component in Synthetics
* Preserving fastData settings properly throughout dashboard panels

---

© 2014-25 Kentik
