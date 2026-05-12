# Release Notes 2021

Source: https://kb.kentik.com/docs/release-notes-2021

---

Kentik employs a continuous deployment methodology for constant extension and refinement of the Kentik v4 portal and the underlying Kentik platform. Release notes for each successive month of Kentik v4 updates are covered in the following topics:

* **December 2021**
* **November 2021**
* **October 2021**
* **September 2021**
* **August 2021**
* **July 2021**
* **June 2021**
* **May 2021**
* **April 2021**
* **March 2021**
* **February 2021**
* **January 2021**

> ***Note:*** For additional insight into what's new with Kentik, be sure to check the **Product Updates** page of our website (Platform » Product Updates).

## December 2021

**New Functionality**

* **Synthetics**

  + Ping+Trace in PageLoad Tests
  + BGP Monitor Aggregated Prefixes and Prefix Selector Filter
  + Support for ALL Notification Types in Synthetics
* **Cloud**

  + WeatherMap Layer Selector: Utilization layer, Health Layer
  + Map sidebar UX/UI Enhancements
* **DDoS Mitigation**

  + Support for multiple mitigations in DDoS policies
  + Support to apply / clear timers per policy
  + Acknowledgements required for thresholds
* **Core**

  + Complete Capacity Planning workflow overhaul
  + Link Sharing
  + Kentik Market Intelligence

**Issues Resolved**

* 10239 - Looking to add a data source
* 10234 - Subtest Threshold for Status doesn't work
* 10231 - Profile screen won't load
* 10226 - Sometimes tests appears as empty on the the view, BUT if you click on the Dest IP we can see that th
* 10209 - Complains that email and token are missing
* 10162 - Traffic not being balanced through ECMP for FLOW Counters
* 10151 - false-positive 100% packet loss
* 10135 - Creating a Saved View doesn't work for member access
* 10095 - why is the aggregate 20% packet loss?
* 10088 - Clouds page hangs when there are no clouds
* 10084 - [SYN] Test Description Overflow
* 10074 - choosing wholesale or backbone for customer type gives NaN for the percentage
* 10060 - Member users don't have access to KMI
* 10054 - Amplification and Reflection Attack is not working
* 10045 - 404 error - following ksynth agent docker hyperlink..
* 10019 - Search results return a 500 toast
* 10017 - Dark mode legibility issue for market selector
* 10006 - Unique Sources not available condition
* 9999 - no threshold numbers
* 9997 - No Subpolicy Alarms in v4
* 9996 - show path to is drawing an errant arrow
* 9995 - Missing a Mitigation Method from the /v0/mitigations/methods results
* 9987 - DDoS Defense: Cannot disable policies
* 9960 - [MKP] Subscribe Button Error
* 9957 - [MKP] 400 Changing Policy
* 9953 - AWS traffic trends and overview, shows no data for recently added accounts.
* 9948 - Different AS groups while just one expected
* 9946 - Strange alert
* 9936 - Subnet showing up in wrong availability zone
* 9935 - MKP Page Feedback Button While Spoofed (2)
* 9934 - MKP Page Feedback Button While Spoofed
* 9928 - Ping metrics missing in Page Load subtest screens
* 9901 - network boundary classification appears to be intermittently breaking for this device- br1\_sin1
* 9891 - Cannot navigate to Network Details on EU portal
* 9860 - While using the KENTIK-NO-FLOW bucket for metadata collection only, an Export Process Timed out error
* 9844 - ASN selection in the config screen throws a cryptic error message preventing to save
* 9817 - Contrast issue w/ Market selector in dark mode
* 9809 - AS0 should never show up
* 9804 - My Network barchart sometimes misses ASN
* 9796 - market selector menu has a layout issue
* 9760 - DE: Visualization depth not respected when using SNMP metrics
* 9747 - Extract from OTT regex not working for OTT fields
* 9709 - Different results from Web and API II
* 9683 - synthetic alarms are in insights and alerting
* 9608 - mitigation method created in v4, does not have DSCP value set when viewed in v3
* 9577 - BGP Monitor notification parameters not saving
* 9474 - Saving Alerting Policies 500
* 9439 - Azure Cloud Map - "Show Details" Pane Has No Data
* 9312 - Settings » Custom DNS: Click to remove a DNS server changes its name rather than removing it from th
* 9261 - traffic sidebar for weathermap looks a little messy
* 9250 - Test notification button not working for MS teams
* 9018 - Settings > Notifications: Clicking the column heading Daily doesn't sort the column. It leaves it un
* 8891 - Market filtering side panel: cannot scroll
* 8883 - instead of single selection from the Network Classification, we need to be able to do multiple in the config screen
* 8882 - Cannot select Region in Market Selector
* 8866 - Filter in the Market selection is broken
* 8493 - hide new notification channels until we have proper templates
* 8451 - Notifications on Sub-Policies at MKP is not working
* 8230 - [Show Path] Not Drawing Less Preferred Route
* 7934 - Issue With DDoS Mitigation selection in UI
* 7688 - flow logs are being deleted from the S3 bucket, even though this is disabled.
* 7576 - Lots of missing AWS data
* 7501 - Kproxy DNS entries are currently creating Dimensions under Custom Dimensions, as well as kt\_dimensio
* 7375 - data derived from AWS VPC flow logs is not showing up in Data Explorer
* 7374 - Cloud  - AWS Traffic Trends - Dependency Check Failure
* 7340 - Tableau host is missing metadata for some DB records
* 7243 - [Synthetics] Global Agents with the same Location and ASN
* 7201 - tenant policies not activating
* 6726 - AWS-Device-Label does not open Cloud Dimensions
* 6701 - Current value is 0bps but there is traffic
* 6690 - After enabling Show Historical - switching between Type no longer functions correctly.  Kentik Insig
* 6667 - Dashboard Results are inconsistant
* 6573 - I receive an unknown error when I select Kentik Map within the iNAP portal.
* 6550 - maps/interface metadata - lines are mixed up
* 6438 - Connectivity Cost, overcharging
* 6397 - Orangeflow config - restart of Docker shows unable to connect to Kentik API
* 6386 - Synthetic dashboard panels do not export correctly
* 6119 - DDoS Defense Backfill Jobs not running
* 6059 - IBM VPC map view does not show a chart
* 5803 - IBM traffic seems like its a ramp. is this a bug?
* 5485 - v4 shows 3 alarms needing an Ack but v3 does not show any active alarms.
* 5481 - AWS ksynth Agent IP Does Not Update After Reboot
* 5324 - another missing device from the hybrid map. i feel like there must be a bug that misconfigures this.
* 5286 - Ensure users can see historical data for paused tests
* 5155 - metadata in account is not appearing consistently
* 3964 - When a test is paused, it appears we are unable to access historical data by changing the time range
* 3838 - alerting is mixing up the keys of this CPU metric alarm
* 3640 - botnet comparison metrics is broken
* 2976 - bidirectional isn't producing the second query

## November 2021

**New Functionality**

* **Synthetics**

  + BGP Monitor in EU-SaaS (no notifications)
  + BGP Route Viewer UX improvements
* **Cloud**

  + Weathermap enhancements
  + Bidirectional support for Show Path Feature
  + Consolidated views for high-density Cloud Interconnects
  + SFlow VXLAN for the Data Center
  + Other backend enhancements
* **DDoS/Alerting**

  + Factors for DDoS alert details
  + Acknowledgement action button for alerts
  + Mitigations link in main nav
* **Core**

  + Auto-magic Dark / Light theme switching based on OS settings
  + Library UX improvements

**Issues Resolved**

* 9905 - [Synth] Ignore TLS not persisting in Edit
* 9877 - Uber alerting policies using static and baseline error out after coming out of silent mode
* 9862 - Demo Environment link missing?
* 9852 - IPv4 missing traffic for a specific interface - edge1\_scl1
* 9851 - All Queries keep erroring out
* 9840 - Synthetic BGP view error
* 9838 - Errant ASNs appear when window is narrowed
* 9836 - Market Intelligence : AS1299 looks to be a customer of AS3356
* 9819 - Powershell script for Azure failed near end
* 9787 - Unable to re-use activation code if user closes Activate modal
* 9783 - IPv4 missing traffic for a interface - edge1\_scl1
* 9766 - Public clouds listing - status warning shown when not necessary for NO-FLOW Azure devices
* 9765 - why did all these devices drop flow across the board
* 9762 - Incorrect Security Group Matching Logic
* 9759 - DDoS policy changes won't save
* 9754 - Capacity planning: differences between values seen in DE
* 9732 - Alerting: manually cleared mitigation keep showing up with failed to clear
* 9731 - Alerting: flowspec mitigation created in v4 rate limit value not visible in v3
* 9730 - we need to change these ip address to be something other than kentik ips because its messing with OT
* 9728 - only one direct link is visible
* 9727 - Only one DirectConnect visible in AWS Kentik Map
* 9716 - an insight should be show in the tab here as well https://portal.kentik.com/v4/operate/insights/x440
* 9715 - there should be an insight associated with this test https://portal.kentik.com/v4/core/insights/x438
* 9714 - this test should have an insight associated with it but its not listed in the insights tab https://p
* 9700 - Test Name Change Edit Not Working
* 9688 - 500 error on this cloud page
* 9682 - Line regression in performance monitor - health colors are gone
* 9674 - Network 500 error when trying to save a copied test
* 9673 - AWS Config Status often does not return without a page reload
* 9667 - DDoS Mitigation: Issue to select value of 0 in the mitigation method
* 9666 - Manual mitigation is not working
* 9663 - DDoS Mitigation: manual mitigation reports "remote error"
* 9650 - apiv6: application\_mesh test API validation must require 2 or more agents
* 9633 - BGP Route viewer: Total column in the table is not sorted as number but as string.
* 9632 - BGP Route viewer: dataset and AS Path length disappear from the table
* 9631 - Alert email notification contains alarms not associated with the notification channel
* 9625 - mitigation did not fire when alert was triggered.
* 9624 - Tenant Notification Channel issues
* 9607 - [Connectivity Cost] Connectivity Type does not match Manage Interfaces
* 9601 - maps not drawing correctly
* 9587 - EU SaaS Agent Install instructions incorrect
* 9582 - 404 error opening map
* 9577 - BGP Monitor notification parameters not saving
* 9574 - [Library » Dashboards] Add Synthetic Test Panel dialog allows Display Type of "Density Grid" when Panel Type is "Single"
* 9566 - [Library » Dashboards] Erroneous settings in Edit Synthetic Test Panel dialog, causing error
* 9562 - AWS / kappa - No Latency Data
* 9561 - Mitigation didn't trigger for alert
* 9557 - 30min synthetic test run, but 15min test display window
* 9552 - Inconsistent Breadcrumb for BGP Monitor
* 9551 - "No data" shown when there are no active alerts
* 9545 - AWS Config Status shows flow status for wrong VPC
* 9543 - seems like the weathermap traffic arrows are reversed from the chord digram
* 9532 - Update URLs for SSO icons on /v4/integrations page
* 9527 - Differences between API and Web results
* 9510 - error from notification channel
* 9494 - With "Aggregate Pulldown". Graph data appears to be identical when select any of the options Max, 95
* 9476 - It shows no result for Full Dataseries but it have result in Fast Dataseries
* 9469 - When creating a new dashboard, if you create a new category but don't hit "save" and then try to sub
* 9468 - Clicking on a vpc card in next throws an error occured
* 9466 - Update Integrations page (/v4/integrations)
* 9463 - not receiving flow from interface Eth-Trunk100
* 9462 - Manual Mitigation TTL Broken
* 9459 - While trying to add two new private agents to our synthetics mesh test I get a 500 error
* 9450 - [Library » clone a Saved View] Dialog name should be Clone Save View
* 9422 - Test Button in v2 Webhook Notification not Working
* 9410 - Device Config for SNMP Device has wrong URL for ""Need help configuring SNMP polling..."
* 9400 - DDoS Defense: mitigation platform-method keeps replicating in UI multiple times
* 9392 - Multiple error about notification channels in DDoS Defense Configuration when trying to add existing
* 9385 - Visualizations don't render if Charts per Row is set above 1
* 9384 - I can not choose "Ultimate exit interface" in the "Alert Polices"
* 9377 - this hostname test (and others) are failing because the target is not supported
* 9367 - Performance monitor is stepping on the legend
* 9366 - change in v3 policy which is allocated to the MKP tenants throws 500 Error
* 9359 - Path View Agent drop down shows way more agents than are configured in the test
* 9352 - Ultimate Exit Interface stop working
* 9319 - On remove popup with loading indicator disappears before loading completes
* 9306 - formatting issues when I expand the "Collapsed Path"
* 9305 - Agents positioned incorrectly on map
* 9272 - [COGENT] Kentik Map Path Choice
* 9263 - kentik map toolbar seems to obfuscating text
* 9248 - Notifications disappear from Attack Profile configuration
* 9247 - Error when trying to save Attack Profile config.
* 9245 - entry prompt for Notifications occassionally disappears.
* 9238 - obs deck - 'recent data explorer queries' display issue
* 9230 - Unable to directly switch from peer with this, to use another device
* 9212 - Tenants Receiving Other Tenant's Alerts
* 9187 - Map will not display
* 9174 - Time range displays in Local if you select a time range, but applies it in UTC if you modify it
* 9076 - 500 error on entering Subscription ID in Add Azure Cloud
* 9068 - v4 Community Backwards
* 9050 - Sankey colors not making sense
* 9049 - Test has results but TCC says failing
* 8910 - tests don't fully load for this account
* 8659 - webhook - DDOS notification not being saved
* 8533 - Query is no longer returning results
* 8331 - errors are happening constantly. when you select the right time range. but when you zoom out you don
* 8303 - timeline is in UTC but the datepicker provides local time.
* 8270 - v4 Mitigation 500 Error
* 7979 - No DNS decode dimension seen from kprobe
* 7904 - confusion between local tz/utc
* 7883 - Kentik Map » AWS Topology: Some "connection" links aren't being drawn
* 7863 - No Azure Zones listed
* 7856 - DNS fields not reporting correctly
* 7842 - DE displays Connectivity Type in interface Dimension column only when Connectivity Type is selected
* 7771 - Non Events Triggering PagerDuty
* 7767 - changing threshold value and the other threshold changes
* 7544 - On Prem Unknown Device on Kentik Map
* 7486 - Settings » Add Cloud workflow takes you to Welcome to Network Observability page
* 7444 - Clicking on any "Interface" from DE table (SNMP dimension) takes one to a blank screen that never re
* 7413 - Cloud-GCP Traffic Trends and Overview "no results"
* 7394 - Unassigned layer in Sites' Kentik Map Overflows
* 7343 - Leakage policies appear broken
* 7338 - cloud dimensions is not selectable for this saved view. it should be
* 7327 - Wrong GEO location for a Synthetic probe (flz-dcs7280sr2a-edge-1)
* 7298 - Getting RFC1918 Route Prefixes without BGP setup
* 7291 - Multiple Hostname dimensions shown in Filtering options
* 7111 - Network explorer throws a bug each time this account loads
* 7098 - saved filters not saved to policy
* 7087 - Device shows on proxy side, but no flow reaches Kentik
* 7053 - Kproxy -log\_dir option doesn't work
* 6959 - something is wrong with this lab traffic. this isn't really happening. not sure why UE device is the
* 6826 - The "Add to a dashboard" (New or Existing) option has disappeared from v3 to v4
* 6821 - Traceroute to Azure or AWS public agents do not complete with either UDP or TCP from private agents.
* 6818 - Error when trying to use Edit Architecture feature of "FAB4" site. Feature appears to work on other
* 6789 - Adding Panel-Filter Skews Results
* 6711 - In the test control center for synthetics if I change the time range to look back more then 1hr I ge
* 6689 - No Data being provide in the top × table when displaying a drilldown for an IP.
* 6686 - kproxy not able to upload DNS query results
* 6645 - custom alerts I have configured do not show up in alerting unless I toggle show historical.
* 6634 - Different results in a guided dashboard between v3 and v4
* 6628 - Mitigation is taking 1m to 1m45s to trigger once an alert becomes active
* 6601 - Panel will not accept filters
* 6548 - 'error occurred' from kentik map/show connections
* 6532 - not able to apply labels to devices
* 6491 - Kproxy not sending flows to Kentik for Nokia router
* 6487 - No traffic showing on connectivity cost site
* 6441 - number of views all 0 if package used
* 6425 - on some reloads we get stuck without the viz type chosen, or long delays and variation in total/hist
* 6417 - related insights is not working for this site. it should pick up https://portal.kentik.com/v4/operat
* 6290 - why does this synthetic insight not have any data on the drill down. seems like there is a bug in ei
* 5993 - this doesn't seem like an insight to me.
* 5830 - these interfaces shouldn't really be connected. we should change the cisco nbar side
* 5737 - application insights isn't generating a description for some reason.
* 5343 - i think this chart used to work on the insights.
* 4652 - current traffic value of 0 is not correct
* 3562 - inaccurate throughput on device
* 2921 - really big issue with IC
* 2737 - "CDN and OTT" activation - "go to explorer" buttons in CDN and OTT workflows
* 2238 - Interfaces not showing description

## October 2021

**New Functionality**

* **Core**

  + Library has gotten a major UI/UX update
  + Actions menu update on both Dashboards and Saved VIews
  + Connectivity Costs history capabilities
* **Automation & Guided Workflow**

  + In October ˜1,000 hostname patterns were added to Kentik’s OTT content detection catalog, with focus on providers from Vietnam, Brazil, and Germany.
* **Synthetics**

  + BGP RPKI Validation
  + BGP Monitor for Prefix Reachability (dev)
  + App Agent Private Agent (Limited Beta/ Internal)
* **Cloud**

  + Kentik Kube: Network Visibility for Kubernetes Clusters
  + In-bucket file retention for AWS Flow Logs
  + Support for VPC Endpoints and ENI Gateways
  + AWS Flowlog ingest via Kentik-owned S3 bucket
  + Map rendering improvements

**Issues Resolved**

* 966 - Customer Analytics incorrectly shows a 0.00 calculation
* 2519 - Device Labels: browser freeze when trying to edit existing label
* 2864 - AT&T is not recognized as a valid provider name.
* 3032 - UE Interface as dimension in alert triggers error
* 4303 - DID missing
* 4475 - When zooming into a graph, measured values change by exactly 10x.
* 4508 - V4 Cloud Setup Issue
* 5021 - Tenant Template Description Overflow
* 5095 - Destination Traffic Termination is reported as Inside, but Dest IP is external.
* 5130 - Device labels returns 404
* 5262 - If drop-down is selected in Customize Metrics, screen closes before any other changes can be made
* 5393 - No flow from device reported with v4 when clearly good valid flows are arriving. 5352 - health monitoring - incorrect capacity reported on an IRB interface
* 5394 - Interface classification rules don't match when using type=SilverPeak.
* 5396 - Unable to edit dashboard
* 5424 - Flow Type is empty on devices sending flow via kproxy
* 5425 - Devices list shows 0 v4 bgp prefixes
* 5439 - ibm popover seems broken
* 5450 - CDN Analytics not showing Embedded Cache traffic
* 5453 - click on peering inbound/outbound selector produces access forbidden error for member users
* 5579 - Traffic by Connectivity pie chart appears incorrect
* 5603 - Device 'More Info' screen indicates incorrect results
* 5612 - Resizing the parent panel in MKP causes all dashboard panels to become the same size
* 5623 - VRF Name filter not working properly
* 5663 - flow type is missing on these devices
* 5664 - Ampersands in interfaces seem to cause filtering issues
* 5685 - V4 DDoS - Baseline Backfill - Jobs not completing
* 5717 - incorrectly parsed ampersand in description - see &amp
* 5769 - missing formatting of bits/s
* 5783 - OTT Service Tracking should indicate whether OTT is enabled on account
* 5789 - date range isn't fully shown in this from-to query
* 5798 - MKP policy is going off at tenant levels in landlord
* 5810 - preview as tenant is flipping my inbound/outbound
* 5831 - v3 firewall dashboard doesn't work the same as the v4. see screenshots.
* 5867 - Updating Custom Dimension Populators Isn't working but still return '200' response
* 5877 - the View Interfaces button from a specific device doesn't filter to opening the interfaces from a sp
* 5884 - We're not displaying the internal/other traffic
* 5923 - Formatting bug on OTT export
* 5950 - streaming telemetry data doesn't have the correct IC values.
* 6010 - interface costs per interface looks like it had a regression
* 6223 - Cannot add Mitigation Method when configuring a Policy
* 6268 - peering IP not showing in device config.
* 6269 - ingress/egress calculation is backwards
* 6288 - The "Discover Peers" dashboard gives "Error 500" and the potential peers tab never loads.
* 6301 - Incorrect sso URL in profile on EU cluster
* 6354 - no dashboard when logging in
* 6372 - Weird behaviour with URLs obtained from the API
* 6413 - Graphs are not loading
* 6511 - dashboard not able to load
* 6517 - 500 on tenant save
* 6547 - Receiving 'Access Denied' when attempting to access the Audit Log.
* 6626 - Mitigations not firing
* 6659 - can't select aws dimensions
* 6695 - Missing Alarm
* 6823 - Filters not working
* 6897 - v3 Alerts Wrong AS Name
* 7034 - Cleared DDoS alerts show a state of cleared in the 1st column, but under the time column, still show
* 7102 - incorrectly quoted ampersand
* 7167 - Cannot Add Saved Filter
* 7212 - Turning on new Alert Policies
* 7304 - Customer gateways should be drawn inside the devices box, not outside
* 7376 - Impossible to add a new AWS Cloud into portal
* 7575 - Some global agents can ping while others cannot
* 7606 - Tenant policy search broken
* 7652 - Cloned Dashboard Panel Titles
* 7672 - no flow alert triggers but there appears to be flow
* 7807 - Manual Mitigation 500
* 7812 - HTTP metrics not being returned on HTTP(S) test
* 7914 - Cloud Performance - Stuck In "Loading health"
* 8040 - Interface Classification
* 8114 - Silver peak interface incorrectly showing no interfaces
* 8219 - Paused test continues to run
* 8256 - DDoS Attack Profiles present, but alert profile doesn't exist or trigger anything
* 8299 - tenant view details on alarm gives 404 page
* 8307 - Notification preview test for slack does not work
* 8360 - Can not delete a specific dashboard in Edit Tenant
* 8388 - Syn Details Unknown Error
* 8389 - Syn Status Mismatch
* 8456 - The Terraform resource file template does not use the most recent version of the kentik-cloudexport
* 8474 - Audit Log not loading
* 8485 - Flow Type field is not populated
* 8509 - v4 DDoS policy adding multiple notifications saving and returning back does not show all of them
* 8562 - Select a new name for a dashboard clone?
* 8567 - Issues with Data Explorer Report - Filter Ultimate Exit Interface
* 8589 - Ultimate Exit interface information missing
* 8643 - Deleting mitigation fails with 404 Not found: sql no rows in result set
* 8659 - webhook - DDOS notification not being saved
* 8686 - apiv6: Cannot modify existing test
* 8699 - drag-n-drop doesn't refresh view
* 8801 - ksynth installation instructions do not point to REGION dependancies
* 8805 - [MKP] Guided Interface Filter
* 8812 - Spoofed Tenant Bug Report
* 8823 - Category Creation bug
* 8827 - Category rename action not implemented
* 8829 - Favorites entries text seems centered
* 8830 - Incorrect "Remove from favorites" tooltip
* 8831 - debug text appended to descriptions
* 8832 - Add Guided Mode input field to Guided Mode dashboards
* 8834 - Bottom border of user content heading should be plain, not dotted
* 8838 - empty category matches filter string for no reason
* 8893 - BGP info not loading
* 8904 - DDoS Defense TCP SYN Flood policy has a wrong filter "TCP Flags is set SYN"
* 8906 - When not using Major revisiting profile turns Major back on
* 8907 - DDoS Policy Config drawer - enabled switch doesn't update
* 8937 - Adding a router to the mitigation platform results in 500 Error
* 8940 - Observation Deck shows "error occured"
* 8944 - Incident Log only shows 14 day history, even though there were Incidents further back.
* 8954 - apiv6: API allows to set `settings.health\_settings.http\_valid\_codes` to an unsupported value
* 8962 - Removing provider affects historic data
* 8972 - users observation deck returns blank page with "Error Occured"
* 8976 - Audit Log Doesn't Respect Time
* 8987 - Reported data does not match actual router data
* 8988 - categories without content still have the dropdown arrow
* 8989 - Categories to reorder after renaming
* 8991 - impossible to add notification channels
* 9014 - Not able to clone SaaS test and Save
* 9017 - Socket resubscribes can potentially go on indefinitely and cause nodes to crash
* 9020 - notification timeout errors in go from operations
* 9037 - Strange behavior with traffic classification
* 9047 - AWS Flow log metadata ignored by Kentik
* 9060 - Adding Mitigation Platform fails
* 9061 - "Enable" slider on some attack profiles set to "on" although policy is disabled
* 9067 - Orangeflow v1.3 -sqs\_messages argument is limited to 10
* 9096 - Kentki Preset Saved Views are globally writeable by anyone
* 9098 - do not display flowspec mitigations in v4 if company does not have them enabled
* 9103 - Can't add Azure Cloud
* 9104 - DNS perf results should point to "last hour" not filtered to last minute
* 9111 - Callout if a dash has been deleted since last viewed
* 9112 - Add subscription for Guided Mode dashboard
* 9113 - Getting an error when trying to add pager-duty as a notification channel to a syn test
* 9118 - strings w/ underscores break search
* 9131 - While adding Azure resource groups, Kentik portal sent back 500 errors
* 9140 - Webhook Template Text Box
* 9148 - Paused Synth Test and unable to resume as it indicates I will be over my available credits afterward
* 9162 - AWS view - Kentik map not displaying any data - mapping objects
* 9171 - Changing the "Include covered prefixes" option does not stick.
* 9176 - no way to email dashboard as a one off
* 9177 - reports need date/time information
* 9214 - Agent setup defaults to US even in the EU
* 9218 - DDoS - flowspec method created in v3 misses traffic action field in v4
* 9231 - Error Occurred on viewing syn test details
* 9240 - Unable to use universal search
* 9249 - 500 error when trying to create a mitigation platform
* 9251 - Origin network query broken on kentik weathermap
* 9253 - Can't delete certain folders from the Kentik preset library
* 9267 - Synth dynamic tests don't auto-refresh
* 9277 - DDoS Notification Channels
* 9281 - Changing a Dash/View name & then favoriting it
* 9292 - test notification not working with email.
* 9312 - Settings » Custom DNS: Click to remove a DNS server changes its name rather than removing it
* 9323 - [Library » splash screen] Continue to Library btn doesn't work
* 9325 - hostname view shows the break out by ip address but you aren't able to drill down per ip address
* 9331 - Library - searching breaks sort order
* 9369 - dark mode not legible enough
* 9370 - Network Error - Status 500 when trying to add RTBH mitigation method
* 9371 - RTBH Community created in v3 is being transposed
* 9384 - Ultimate exit interfacenot available in Alert Polices
* 9391 - "Add a subscription" bug
* 9396 - Not all agents show the result of test
* 9416 - No Flow detected on all devices
* 9435 - Only One Community Announced
* 9445 - UI issue with steps 4 & 5 not able to complete after saving.
* 9471 - Tenant Policy Doesn't Fire
* 9473 - DDoS Save Error

## September 2021

**New Functionality**

* **Core**

  + New UI for Kentik Library to better organize dashboards and saved views
* **Platform**

  + Improved GeoIP accuracy
  + ASN name cleanup for improved readability
* **Synthetics**

  + Page Load Asset Validation
  + Dashboard Widgets for Aggregated Test Results
  + Property Templated and Working notifications for BGP Monitor
  + BGP Monitor for Prefix Reachability
  + Added ALL Routeviews (BGP) Vantage Points
* **Cloud/Hybrid**

  + New Weather Map
  + Historical queries
  + Clickable Lines in the Kentik Map for AWS
  + NAT gateways and Transit Gateways
* **DDoS**

  + Mitigation method / platform configuration forms are now native v4 forms
  + Removed limitation on configuring both RTBH and flowspec on the same router
  + V4 DDoS policies now support ratio-based thresholds

**Issues Resolved**

* 8909 - Terraform URL seems wrong
* 8893 - BGP info not loading
* 8846 - Invalid Tests From API
* 8804 - Got error using custom dimension
* 8799 - Unable to add panels to dashboards.  Add Panel button is always grayed-out.
* 8730 - Can't add agent to mesh and save (errors out)
* 8725 - MKP - Error 500 Active Mitigations
* 8717 - BGP Test results: custom time range doesn't work
* 8661 - Cannot add Site visualisations to Observation Deck
* 8626 - apiv6: Cannot create dns\_grid test
* 8604 - Credit Exceeded value is incorrect
* 8590 - Misclassification - Ultimate Exit None
* 8573 - MKP 500 Errors
* 8558 - Clicking Save for Later expands/exposes the Amplification attacks
* 8549 - need to create presets (SaaS/ DNS) in eu-saas - shows blank until we do
* 8545 - apiv6: agent.name should be agent.site\_name
* 8543 - apiv6: Syn List Agent API returns all agents (including pending global agents)
* 8490 - BGP Monitor not generating alerts
* 8465 - Interface utilization spike insight is still visible
* 8430 - Synthetic alerts not sending clear to custom webhook
* 8416 - 500 adding policy to tenant
* 8404 - Saved views are not in the correct category
* 8342 - azure location list is missing values
* 8316 - Alerting Page 500
* 8234 - Add subnet CIDR block to the metadata header in the sidebar
* 8158 - View Alert Policy better readable values
* 8109 -  Issues with Flowspec Manual Mitigation
* 8012 - Traffic is not being displayed with message "No Results".
* 7743 - Route table searches should use bitwise contains comparator
* 7520 - wrong name for conn type
* 7054 - Cloud Missing Vertical Scroll
* 6815 - kproxy using local DNS not always picking up mapping.
* 5621 - Tenants show no Alarms
* 5501 - Tenant Alarms Missing
* 2442 - exclude UE provider = blank causes no results
* 2428 - UE Provider of "----" can not be totally excluded

## August 2021

**New Functionality**

* **Synthetics**

  + Test credit limits now enforced
  + More test frequency options
  + Ignore TLS option
  + BGP Monitor test credits update
  + Standard Deviation (baseline) based thresholds
  + Performance Dashboard presets updates, including option to request new entries in the preset tabs
  + New “DNS Performance” tab
  + “SaaS Performance” tab upgrade (to HTTP GETs)
* Kentik API

  + Site API
  + User Management API
  + Network & Interface Classification API
* Kentik Cloud

  + Network ACL/Security Group Visibility in Kentik Map
  + Support for AWS V5 Dimensions
  + ENI Tagging and Dimensions
  + Cloud Native views for Kentik Map
  + Site-IP dimensions in sidebar traffic queries
  + Better handling of large metadata sets in sidebar
  + Metadata-only exports for Azure
  + Azure stability improvements
  + New AWS Connectivity Troubleshooting tour in Demo Mode
* DDoS / Alerting / Notifications

  + Custom HTTP headers for v2 webhook notification
  + Additional v2 notification methods
  + v2 notification support for v4 DDoS policies

**Issues Resolved**

* 8482 - 500 Error when adding notification channel
* 8467 - azure export from alternative RG doesn't seem to tag metadata
* 8462 - VPC exists, but Kibana - Transfluo shows an error message in onboarding
* 8460 - Site by IP dimension not being populated, even though a match
* 8451 - Notifications on Sub-Policies at MKP is not working
* 8443 - Performance Monitor Screen Crash
* 8429 - AWS Config Status is broken
* 8423 - Path not rendered correctly in eu-saas
* 8411 - Tenant Packages Disappeared
* 8403 - Data aggregated to 60 min interval even when using Full dataseries
* 8378 - Metrics break DE (cause half screen to say "Error Occurred")
* 8373 - 500 on Tenant landing page
* 8372 - Valid HTTP and DNS code drop downs are empty
* 8367 - Can't Remove Dashboards
* 8365 - 404 on MKP Dashboard
* 8361 - The combination of Magic Tenant Custom Dimension and Filter-Based dimension does not work
* 8356 - Number of VMs shown in Kentik Map incorrect
* 8353 - UDP trace is default but says it's invalid
* 8346 - Can't save a synthetic test
* 8327 - Selected time range inconsistent
* 8326 - Cant create users.
* 8322 - full page error trying to save DDoS config
* 8317 - Docker Agent Failure
* 8294 - "Settings->Public Clouds" page show "Error Occurred"
* 8273 - this agent test is failing but we are showing no results.
* 8255 - BGP Route Viewer not displaying anything
* 8232 - HTTP test (rust) works but pageload (node) does not (to same endpoint)
* 8231 - Cisco Icon instead of Juniper
* 8224 - Alibaba Agents don't display provider name
* 8222 - Test stopped working until agents were removed and re-added
* 8220 - Internal error when trying to update Alerts for a tenant on MKP
* 8217 - SNMP Values not displayed
* 8214 - searching on a subnet ID doesn't return any results
* 8206 - filtering on logging status is throwing an error in data explorer
* 8196 - BGP Monitor Test Notifications Not Sent
* 8193 - Open sidebar header with metadata in a collapsed state
* 8183 - Test slack notifications not working
* 8171 - Agents are not selectable
* 8163 - Capacity Export Sort Wrong
* 8159 - kproxy -proxy\_http and -local\_proxy do not work together
* 8049 - API timing out
* 7881 - Entity Explorer traffic graphs are all reporting "No traffic data"
* 7777 - Edge doesn't redirect to v3 shim properly
* 7554 - Botnet Dashboard Export Error
* 7543 - Inter-Region Traffic Line Not Drawing
* 7520 - wrong name for conn type
* 7392 - Sites Hamburger DE Link Fails
* 7247 - discover peers page doesn't refresh once you exclude an ASN
* 7156 - Resizing window shuffles panels
* 7056 - Validating cloud causes 504's
* 6780 - Devices put in an architecture layer appear there, but duplicated as well in the Unassigned list.
* 6685 - Don't show geo/ ASN info for NAT'ed nodes
* 6409 - Wrong AS in Path
* 5301 - Can't filter based on AWS logging status without throwing errors

## July 2021

**New Functionality**

* **Synthetics**:

  + Major Path UX overhaul
  + Page Load Test
  + BGP Monitor for Hijack detection
  + More BGP Vantage Points
* **Cloud**:

  + Show Path Feature
  + AWS Configuration Status
  + Search feature for Map & Performance Monitor
  + Support for AWS External ID
  + Backend improvements for scalability, troubleshooting, and data completeness

**Issues Resolved**

* 8166 - Cap Group Export Fail
* 8157 - Exporter Status / Convert Flow Logs From v2 to v5
* 8151 - Configured Notifications do not show up in config screen.
* 8145 - Error occurred when accessing swagger tab of API tester
* 8144 - Selecting a notification does not show it selected
* 8143 - API does not retrieve data for data longer than few hours.
* 8136 - Can't Switch to demo environment
* 8124 - Routeviewer not loading
* 8123 - saved views are not saving the Category
* 8119 - onboarding wizard does not recognise manually added devices.
* 8106 - OTT Service Tracking graphs have conflicting data
* 8103 - $lookups.asNumbers fields clears on second item
* 8098 - [OTT] » Service » Subscribers tab: Subscribers chart doesn't resize.
* 8094 - Runout Abbreviations
* 8093 - BGP setup link in Settings/Devices/Edit should be changed
* 8092 - Azure data not showing flows from Australia
* 8091 - TenGigE0/3/0/0 "UNUSED" should be classified via IC
* 8089 - account members cannot access OTT workflow
* 8087 - Internal Error - Status 500
* 8085 - Setup Tasks 0% Complete, even with users and routers added
* 8081 - OTT Service Tracking » Service» Capacity tab: Percent reported as NaN
* 8071 - Overview/Connectivty/Subscriber links not working for a particular content source.
* 8063 - Public Cloud Settings
* 8062 - Guided Dashboard Time Zone Issue
* 8053 - Capacity planning report download and email never completes
* 8046 - [OTT] Provide details page > "Top Services" horizontal barchart is wrong
* 8044 - NE Map 404s
* 8042 - PDF exports are blank
* 8041 - Regions are incorrect in sub- South Africa E.g. England, budapest, beijing etc...
* 8033 - [OTT Capacity] Interface capacity = 0 for a specific interface
* 8006 - Health labels changing
* 8005 - bug in showing the threshold values in the graph for DDoS defense configuration
* 7993 - instance count is not responsive to text search in entity explorer
* 7975 - AWS Cloud Device Not Processing Flow
* 7963 - Botnet shows up as a saved filter in search, but it is not in the saved filter llist
* 7953 - Can't Update Tenant
* 7952 - Alert $current displayed
* 7945 - Plan FPS being enforced as a flowpak not per device - EU
* 7944 - Router with 100G interface shows more than 100G of traffic
* 7943 - Filtering on the agent management page causes filter to be applied in test config
* 7939 - Unspoofing as a Landlord from an MKP tenant should take the landlord-user back to the MKP tenants
* 7938 - Tenant account activation screen should use landlord configured logo, not kentik logo
* 7937 - Page Load UX inconsistencies
* 7935 - no response or no results when selecting timeranges
* 7930 - SNMP total bytes metric is shown as Bits/sec
* 7928 - Feedback Icon/Menu is missing on Configure Custom Insights / Alert Policies Page
* 7925 - lost ability to see next hops in transit gateway route table
* 7921 - Links in MKP-shared synthetics dashboards lead to a 404
* 7920 - Kentik Map » AWS Topology » Details sidebar: Performance pane is present but isn't supposed to be
* 7918 - 40 Devices Breaks Layer
* 7917 - Botnet & Threat-feed Analysis Workflow will not run for any timeframe other than 1 hour.
* 7916 - Changing "Draw Links Using Connected Interfaces" options doesn't take effect until a manual refresh
* 7912 - Cloud performance monitor page changes to NOT FOUND when selecting a test that is not found
* 7907 - Export errors
* 7901 - Kentik Map » Cloud Topology » Region Topology: Error on choosing Show Details from VNet popup
* 7897 - DDoS attack log shows "currently active" end time for all alerts
* 7895 - Performance monitor is showing green links after the agent VM is deleted/unavailable, and stays black after the agent and test config is deleted
* 7893 - Include 2Bytes ASN in AS Groups
* 7891 - DNS Application decodes very "flaky"
* 7874 - when as-name and descr fields are available, use as-name
* 7866 - BGP Settings not Refreshing
* 7859 - Active Attacks shown as zero, when two Alerts are currently active
* 7858 - [NodeJS/ PageLoad] Amazon Nodes don't seem to be able to communicate out
* 7857 - can't change device kproxy agent
* 7851 - Insights On/Off Access Denied
* 7837 - CDN filter
* 7836 - NaN Editing Cost Group
* 7822 - Devices drop back to unassigned the following day after assignment to role.
* 7808 - bottom of screen cut off
* 7790 - HTTP test showing NaN ms
* 7781 - code push caused manual entry of interfaces UE feature to stop working
* 7772 - JSon webhook notification not working
* 7762 - Exporting capacity report errors
* 7754 - Global search only displays 'Search in KB' option if there are results in the main portal search
* 7744 - Download as PDF or Email Report results in error for Capacity Planning
* 7726 - V4 Access Controls show Deny All Except for API, but admin4 Access Controls show Allow All
* 7722 - Add NAT Icon
* 7707 - Tasks completed but still shown as incomplete
* 7706 - Layout preferences saves changes but reverts to prior state
* 7702 - clicking from on profile to another keeps the dimensions of the first one visible.
* 7697 - Switching between attack profiles doesn't display properly unless you close the side display first
* 7690 - hostname monitor works only for IPv4, although "v4+v6" option is selected and hostname is resolved
* 7658 - Accidentally clicking in the greyed out area outside an edit box causes the box to close an any edit
* 7648 - Devices do not stick in their architecture
* 7647 - Data explorer query panel shown over Graphs even with Full width enabled
* 7646 - UI web page constantly reloading in Demo / VPN troubleshooting scenario
* 7628 - Expanding alert reverts focus
* 7607 - MKP Alert Dashboard
* 7579 - Disable V4 Amplification and Reflection Attacks (alltypes)
* 7569 - Alerts under Capacity Planning workflow are not triggering similar alerts under Insights and Alerting
* 7568 - Kentik Maps only shows one site
* 7537 - Time Series charts not showing all agents
* 7531 - Azure agents no longer sending data
* 7521 - MKP Packages Templates not save correctly the dashboards
* 7520 - wrong name for conn type
* 7514 - error occurred in performance monitor for new relic
* 7503 - krpoxy with -dns is using both populators and kt\_dimension to add hostnames. We should only be using
* 7493 - Maps Memory Needs Filter
* 7487 - No username being displayed in created by column
* 7483 - API tester doesn't look good in dark mode
* 7455 - response size says "NaN B" when unavailable
* 7446 - Site-to-Site VPNs / Wrong Info
* 7393 - User Profile defaults for disabled rDNS and Historical Overlay appear to not function.  7391 - Kentik map not showing data
* 7342 - test shows failed when it is not
* 7331 - Incorrect syn traffic graph - filter is srcORdst IP, should be src IP
* 7328 - private agent configuration - unknown state
* 7312 - Dashboard Errored and Never Loaded
* 7271 - Clicking to go to explorer from Capacity Planning results in a bug
* 7209 - Devices keep getting 'unassigned' in Kentik Map (Architecture)
* 7199 - Mitigation timeline timezone issue
* 7198 - Mitigation Field Validation 500
* 7189 - getting error toast in cloud only account
* 7185 - Connectivity Costs » Configure Cost Group » Interfaces tab: interface tables don't sort on columns
* 7163 - Active test goes to Pending on Time Out
* 7126 - can't readd saved filter after deleting
* 7081 - Tests between global Azure agents are flaky
* 7067 - Removing Dashboards Doesn't Stay
* 7060 - "Invalid saved filter (2659)" in standard DNS Analytics dashboard.
* 7023 - azure eastus not reporting any data
* 7002 - Kentik Dashboard Page is not loading. Taking too much time
* 6942 - Some panels don't report all traffic
* 6941 - Istio metrics available even though no Istio devices
* 6895 - Private IP Tests Do Not Work
* 6894 - Can't create architecture for site
* 6876 - Auto Classification 500
* 6861 - Sankey visualization cut off in both DE view and when exporting to PDF
* 6828 - High packet loss from us-east-2
* 6817 - Very high packet loss between ap-northeast-3 and ap-northeast-1
* 6791 - EU trial not enabling Synthetics
* 6732 - Cannot delete platform/method
* 6709 - Order of thresholds is scrambled
* 6679 - The alerts logs are not showing the endtime, instead showing Currently active even with the state as
* 6662 - SP - SYD is experiencing a 21.1% increase in Packet Loss to. should this say 100%?
* 6596 - Rename PDF name when we enable Hide Kentik Branding
* 6501 - Same problem as we've seen before when UE interface is repurposed and flow goes to zero but the boun
* 6353 - Possible Syn Onboarding deadlock situation
* 5951 - Can't create new filter using API
* 5653 - these synthetic tests have errors in the data but we aren't showing that in the UI
* 5129 - arista cpu names seem wrong
* 4331 - interface reporting 100X capacity via SNMP
* 4245 - v3 says there are active mitigations v4 doesn't show any
* 4238 - SNMP interface counters showing impossible values.

## June 2021

**New Functionality**

* **Synthetics**:

  + Density Grid Group option in Dashboards
  + NodeJS Public Agent Deployment
  + Full browser Page Load Test
  + 9 new LATAM agents (EdgeUno)
* **Cloud**:

  + AWS Entity Explorer
  + Configurable sharding for AWS flow logs
  + Configurable sample rate for cloud flow logs

**Issues Resolved**

* 7902 - Firefox: Kentik Map » Azure Topology » Region Topology: No Details drawer after choosing Show Detail
* 7862 - API - Query URL Generator creates invalid URL
* 7860 - Create site from ksynth add new agent creates null
* 7854 - Activate through new agent dialogue box not working
* 7849 - When adding new agent clicking Activate does nothing and the pop up screen remains.
* 7845 - Cloud Performance 500 Error
* 7844 - Failed Syn test
* 7829 - Azure resource group and zone are not showing up in queries. Are we still parsing the logs correctly
* 7818 - choosing 'menu' goes to blank screen with 'Error Occurred'.
* 7810 - All Capacity Plans Trending Down
* 7802 - tenant notification channels should not be an option for landlord notifications
* 7801 - MTK - Tenant ITANET is not displaying any result
* 7780 - opening a transit gateway in the map from the entity explorer results in no links being drawn betwee
* 7769 - agent activation not succeeding
* 7752 - Azure -> Create a Cloud does not add a new azure cloud
* 7749 - Inconsistent AWS Icons in Inventory Widget
* 7725 - Packet Loss in Unit ms
* 7719 - VPC queries should include the gateway ID from the cloud performance module
* 7718 - Clicking on view in Map takes users to the top of the hybrid map instead of the AWS map
* 7717 - Can't actually start a test from the performance monitor
* 7713 - page showing "pageload" test twice
* 7699 - Agent and Status Code column headers overlapping
* 7687 - Not being able to scroll when the number of options overflow in a field
* 7686 - IP Address Missing On Device
* 7666 - Direct Connect not appearing on Kentik map through TGW.
* 7657 - dxconns aren't showing up in appnexus map
* 7656 - configure provider gives error occurred
* 7655 - Connection line not being drawn between transit gateway and transit gateway attachment in AWS cloud.
* 7651 - all synthentic tests showing as failed
* 7650 - Why all tests are suddenly in warning reporting 50% packet loss ?
* 7643 - Syn API Does not seem to reporting packetloss numbers correctly
* 7640 - Transit gateway traffic queries are incomplete (they don't show traffic forwarded from other TGWS)
* 7636 - Missing AWS Instance NamesNo instance names despite having ec2:describeInstances allowed
* 7635 - Network classification broken for Uber's AWS
* 7634 - Switching Datasource Issue
* 7630 - Page does not update with new graph after applying change in lookback time
* 7629 - Unable to apply "Time" filter on a dashboard
* 7626 - Not all the panels are picking up the lookback period when changed from 1d to 2d - only 2 of the pan
* 7620 - Editing GCP cloud does not show all fields from Add GCP Cloud
* 7611 - Cloud Performance - Clicking on any line, expecting popup dialog for deploy agent - nope - error
* 7597 - Created by does not show user in some cases
* 7591 - After clicking launch ksynth agent deployment, backing out of the UI dialog does not work.
* 7590 - When clicking deploy agent, screen to launch shows Debian, RPM, and Docker. - none working
* 7589 - 404 when stepping through Azure setup
* 7587 - folder path unable to be validaed in workflow
* 7586 - configure provider gives "error occurred" page
* 7578 - Setup » Monitor your AWS Cloud: update Step 4 instruction
* 7573 - insight links don't work
* 7566 - Landing Menu broken yet again
* 7560 - Label Filtering Not working
* 7551 - Granularity of results when requesting one month ranges via API
* 7550 - Inaccurate data in Connectivity Costs
* 7545 - E-Mail adresses cannot be added to notifications
* 7541 - MKP - Number of Views (column) indicated on the MKP/Tenants page does not match number of Views actu
* 7538 - Time series values for some agents don't match table below
* 7536 - Connectivity Costs - currency (probably Bug)
* 7524 - Paused tests labeled as "pending" in the perf dashboard
* 7497 - nat-t-ike does not show up in the application filter list
* 7465 - intuit's direct connect gateways are not being drawn correctly
* 7464 - adding a new cloud takes you to the onboarding workflow
* 7447 - Site To Site VPN / Wrong Icon
* 7410 - Demo environment not available
* 7363 - Site Architecture does not Save
* 7263 - FPS SAMPLE is much below from Plan FPS - NKIA-SDR ID: 72075 - others have a similiar behaviour
* 7258 - Removed Ksynth agent. Can't add same agent back.
* 7256 - Showing high Inbound Utilization spikes randomly on different interfaces, far out of range of what S
* 7229 - Changing email address
* 7193 - Internet box doesn't seem to connect to anything
* 7188 - Received error when loading cloud landing page
* 7132 - MKP is very very slow - some panels are not opening or partilally showing data
* 7124 - Drawing TGW To TGW Link
* 7104 - interfaces unclassified?
* 7052 - tests are failing but test page is not updating
* 7004 - CDN Analytics no data
* 6611 - looks like query type 65 (HTTPS) not supported
* 6407 - Custom Dimension Disappearing/Reappearing
* 4712 - why is CDN77 showing up on the first page?

## May 2021

**New Functionality**

* **Cloud**:

  + Observation Deck
  + Kentik Map enhancements
  + Cloud Performance Monitor
  + Automated cloud flow log onboarding based on Terraform
  + New Data Explorer dimensions for cloud
* **Synthetics**:

  + Agent (ksynth) version 0.0.13 and 0.0.14
  + HTTP stage timings and bar charts
  + “Time Series” tab
  + ICMP traceroute
  + Inter-probe delay
  + IPv4/v6 Agent Selection
  + Test Health Improvements
  + Synthetics API Tester (BETA)
* **Connectivity Costs** (major update):

  + Incorporates many customer feature requests
  + Handles scenarios with multiple currencies
* **Setup** (onboarding):

  + Redesigned flow
  + Setup wizards for devices / cloud / synthetics
  + Sets initial Observation Deck widgets based on user interview
* **OTT Tracking**: Accuracy improvements
* **Pie charts**: Max and P95 metrics now available for
* **Main menu**: New features now flagged

**Issues Resolved**

* 7523 - I am not seeing any data in my connectivity costs
* 7506 - User account created by SSO has Default Landing Page of Library
* 7499 - Same test works when in mesh but not in agent-to-agent test
* 7489 - Old v4 DE query doesn't redirect properly to new URL location
* 7481 - AWS cloud configuration doesn't allow users to configure buckets with folders
* 7473 - User mail change failing verification
* 7471 - Alarm Values Wrong
* 7467 - Kentik API brings wrong links
* 7435 - RFC1918 check boxes not recognized
* 7414 - Error Occurred.
* 7409 - no flow data displayed on pppoe
* 7407 - kentik keep kicking me out in last 2 days, not sure what is happening,and the view generation is su
* 7406 - missing values
* 7405 - Configure Flow Log Bucket - can not validate, even though S3 bucket name is correct
* 7402 - AWS onboarding bug
* 7400 - User cannot be created with a dot in the name
* 7395 - Red "Unknown Error" popup whenever opening Network Explorer
* 7384 - Perf Monitor - Node/Link Highlights Not Restricted to Map Type
* 7382 - Packet Loss - Summary shows 0% Packet Loss however I see blips in the graph proving otherwise.
* 7377 - error
* 7362 - Synthetics section missing from new Menu
* 7357 - inaccurate data in Connectivity Cost
* 7355 - Access Denied for users in my company starting today approx 12pm est.
* 7353 - Cross-Panel filtering with SNMP Interface Dimension doesn't add correct filter
* 7352 - Firefox on Windows fonts are not pretty
* 7350 - errant transit gateway line being drawn
* 7349 - performance monitor current test results should color lines red/orange/green
* 7348 - traffic lines don't track with closed/open state
* 7347 - Spelling error in map
* 7339 - PPS Incorrect
* 7326 - Can't add Synthetic (Agent-to-Agent) test
* 7325 - Inconsistent results when going back and forth between time ranges
* 7317 - Can't load the cloud performance monitor
* 7314 - fix error occurred in carfax
* 7307 - why does the traffic line drawn on the map not match our data explorer query?
* 7303 - fix sidebar scrolling behavior
* 7302 - fix tgw lines from drawing through other tgws
* 7297 - some route table entries don't have routes
* 7295 - can't destroy deploy agent window
* 7289 - Route Traffic 500
* 7273 - Error when sorting by dimension
* 7272 - HTTPS test fails when Curl works
* 7266 - All Custom Alert policies in ERROR state
* 7264 - v3 Basic Visibility Dashboard Errors
* 7261 - Cross panel filtering does not work in v4 with table visualizations when clicking on a table row.  F
* 7257 - data table is shaking, cannot read data or get it to stop.
* 7255 - Registering Agent fails
* 7252 - No route viewer showing up for Tata even with flow present
* 7245 - Error Occurred message when trying to view a ksynth agent
* 7239 - 'View Details' of Reverse path generates incorrect URL
* 7238 - No BGP data for Syn only customers
* 7233 - Guided Dashboards Time Zone
* 7232 - region validation never returns if invalid IAM role is supplied
* 7225 - endless waiting
* 7224 - Doesn’t show up performance dashboards (eternal loading)
* 7223 - endless waiting
* 7222 - infinite wait
* 7221 - infinite circle
* 7211 - Receiving "Error Occurred" when viewing certain tests
* 7196 - Time Range Issue below 3 hours
* 7192 - api threw an error when creating a cloud export but still created the export
* 7191 - 500 when trying to create new aws export using admin>clouds
* 7182 - [View Interfaces] in the suggestion cards doesn't filter the interface list
* 7181 - Cost block diagram corner case
* 7177 - Connectivity Costs » Configure Provider page: mislabeled Type field
* 7175 - Connectivity Costs » Breakdown tabs » Traffic History table: column sorting
* 7173 - Can't add policies, nothing in list
* 7171 - TenantService API does not work correctly when the tenant has more than two users
* 7170 - Format in reports' exports
* 7169 - Neverending rotating circle -> showing cloud performance tab in the performance dashboard for syntht
* 7166 - No Slack for Tenants
* 7159 - multilple "500 Network Error" messages.
* 7155 - loss of tab context going back to home page
* 7153 - Download report generates the report for the previous cost UI
* 7148 - Cloud Performance Meshes won't load
* 7147 - 1-hour time range on Performance Dashboard only reports one value for one minute.
* 7142 - total overlay shows by default even though the toggle is set to off.
* 7138 - Test Device Deletion with Tenants
* 7133 - Cannot add Policiies to a tenant
* 7128 - Getting 500 error trying to remove a tenant
* 7127 - Multiple issues found with DNS server monitor tests
* 7120 - Since at least Friday detail pages no OTT service tracking have been crashing browser windows and no
* 7117 - Big difference between 1 day results here and 10 hours. Packet loss is.3 -.5 percent when you look
* 7107 - No hostname dimension available in filtering
* 7096 - 'unknown error' on Kentik landing page
* 7092 - Mesh Test 500 Error
* 7083 - UI weirdness with HTTP test results
* 7076 - Default landing page in profile indicates 'Landing Page has an invalid selection'
* 7071 - Extract from DNS Query does not work, only looks for Legacy DNS dimensions
* 7066 - Site filter in MKP tenant dropdown
* 7061 - Agent Map to wrong Location
* 7047 - observation deck + tour is buggy
* 7031 - Kentik logo on the NAV bar takes me to Library when it should take me to Observation Deck
* 7029 - Re-Map OD widget against spreadsheet
* 7028 - invite email template has <COMPANY> missing
* 7024 - Path Never loads for 68 agent mesh
* 7014 - verify device configuration broken with new onboarding
* 7013 - MKP Policy Error Occurred
* 6978 - Final state of device onboarding not prompting the user to go to Observation Deck.
* 6977 - completion rate doesn't match task list
* 6973 - Kentik Map interface Metadata is mis-aligned
* 6956 - In a saved view, if I change the vis type it just reverts when I click on options to close the DE co
* 6937 - Partial data showing up sometimes
* 6936 - Flat Rate Charge summary on a cost group card is incorrect
* 6934 - Issue w/ legends on provider details page
* 6932 - the tooltips are out of place...the packet loss tool tip shows up near the traffic chart and so on.
* 6924 - Test fails all the time but TCC status shows healthy and no alerts generated in Incident log
* 6906 - Syn widgets don't refresh
* 6901 - Issues after I had upgraded Synthetic agent
* 6898 - banner for no data sources leads to 404 error
* 6870 - Green Incident Log timeline with open incidents
* 6850 - Data agg issue with mesh test
* 6844 - Provider > Traffic history chart issues to be fixed
* 6843 - Cost Group breakdown data-table in Provider details cannot be sorted
* 6840 - Legend seem broken on the traffic chart in cost group details drawer
* 6839 - Why are we showing decimals in the Gbps thresholds of the cost block model ?
* 6838 - Utilization % bug on interfaces in the table on the drawer
* 6835 - can't delete excluded interface filters
* 6825 - Display error for decimal tier thresholds
* 6798 - The sites page is not loading correctly. I cannot view site: FAB4
* 6779 - Kentik Map Error
* 6777 - Sites / Kentik map is behaving weird
* 6746 - Potential Peer Horizon Chart - hover over values are offset to the left
* 6745 - v4 Guided Mode - Override Specific Filters not working in v4.  Works fine in v3
* 6724 - Data Source is 1 with Labels
* 6708 - Drill Down Filter continues to remain to dashboards in MKP
* 6665 - Unable to change email address
* 6657 - No results found when clicking on reverse path
* 6502 - SNMP counters not displaying properly for 2 devices in v4, but appear properly in v3 Interfaces
* 6489 - Link to 'AS Groups' is Broken
* 6437 - Provider suggestions are incorrect
* 6436 - [Add Provider] leads to a blank screen
* 6434 - Issues w/ Flat Rate model
* 6433 - Gbps vs. Mbps in Committed Information Rate
* 6431 - Changing the Committed Information Rate (GBps) doesn't update the diagram
* 6428 - Editing the Unit Price cost in Cost Group details left panel doesn't update the diagram
* 6416 - Selected view is not removed correctly in Edit Tenant
* 6415 - The selected view is sometimes not removed in Edit Tenant
* 6411 - Timeseries on graph seemingly random
* 6405 - Large mesh takes several mins to load
* 6321 - EU SaaS preset account not showing test user name in create/ update list
* 6302 - 67.43.92.0/24 origin-AS 46887 reported as invalid by Kentik
* 6263 - can't use type-ahead to pick dashboard from list
* 6153 - Traffic Origination is coming back with a blank key
* 6114 - are alarms supposed to be firing for this test? packet loss thresholds are set to 10%. all subtests
* 6056 - not sure why there is no data in this query when last 2weeks works.
* 5991 - not sure why this chart doesn't have data. but if you zoom out to the hour granularity there is data.
* 5988 - when you click on a View Details on this test, shouldn't you only see the results from the IP that y
* 5975 - last 2 week query doesn't show two weeks of data for this test
* 5926 - I got a critical alarm in slack with this link and I opened it and the status is healthy.
* 5104 - why is there latency values and 100% loss
* 4454 - Deploy Ksynth agent on RaspberryPi OS

## April 2021

**New Functionality**

* **Synthetics**:

  + New DNS server grid test
  + Aggregated metrics per agent
  + Identified issues table makes it easier to find problems in a large result set
  + Grid view switches to vertical orientation with a large number of agents for easier viewing
  + Clicking on map markers now shows a target list pop-up
* **Synthetics Agent** (ksynth 0.0.12 GA release):

  + Per-probe timeouts
  + Additional HTTP methods (for API testing)
  + Customizable request timers
  + Customizable request headers and body
* **Cloud**:

  + Reduced lag for ingesting AWS flows, now within 4 minutes of write to S3
  + API for automating configuration of new AWS accounts / devices
  + Terraform configurations for automating AWS, GCP, Azure, and Kentik configuration
  + Kentik Cloud onboarding greatly simplified using new Terraform configurations
  + AWS manual onboarding improvements

**Issues Resolved**

* 7070 Fixed traffic shown in tests to align with test timeslices
* 7062 Fixed empty http header field breaking agents
* 7025 Aligned health issues with timeslice not task in synth test table
* 7022 Fixed issue adding policies to MKP tenants
* 7016 Fixed cloud export API storing bucket names
* 6998 Removed path info from health data in synth timeline
* 6989 Showing tests as pending when they’re just created
* 6985, 6981 Fixed issues editing capacity plans
* 6967 Fixed errors due to deleted agents in DNS test results
* 6952 Fixed orangeflow metadata tagging issues
* 6930 Fixed legend/colors with OTT pie chart
* 6926 Fixed text around trace timeouts
* 6925 Fixed map tooltip when agents overlap
* 6908, 6747 Fixed reset api token
* 6902, 6885 Fixed issue with missing health info in http tests
* 6899 Fixed issue with view as tenant CD filters not getting applied correctly
* 6879 Fixed issue with creating site when adding a device
* 6860 Fixed DNS test credit calculation
* 6859, 6663 Fixed private agents tab info
* 6855 Improved map marker algorithm where agents overlap
* 6853 Fixed issues with configuring cost providers caused by bad data
* 6846 Fixed issue with role shown in user profile when tenant spoofing
* 6827, 6814 Fixed issue editing sites when no layers have devices
* 6813 Fixed sorting of v4 and v6 ip addresses together
* 6804 Fixed saved filter save errors caused by audit logging
* 6800 Fixed issue with site layers that have no devices
* 6784 Subtest path for agent to agent showing all targets
* 6781 Fixed issue with Kentik map site not found
* 6766 Added URL changes for tab clicks in Connectivity Costs
* 6758 Fixed IPs shown in agent to agent subtest views
* 6756 Fixed dashboard filters not being applied correctly
* 6744 Fixed bracketing color visibility in Data Explorer legend table
* 6742, 6719 Fixed issue with cross panel filtering in dashboard
* 6710 Prevent Chrome password not in a form warning
* 6702, 6416, 6415 Fixed issue editing/removing views in MKP tenants
* 6694 Removed decimal places for single digit fps in device settings
* 6680 Fixed missing hop details causing failures in path view
* 6677 Fixed issues with too many horizontal tabs in network grid test
* 6656 Fixed sorting agents by version
* 6648 Fixed exclude expired accounts from refresh scan in syn-back
* 6607 Hid dependency checks for MKP tenants
* 6593 Fixed issue with auto selecting plans in cloud onboarding
* 6530 Fixed View in Explorer for site to site traffic in Synth test results
* 6474 Fixed grey series in Kentik map chart
* 6427 Fixed cost group charges not updating diagram
* 6373 Fixed issue with changing packages in MKP
* 6346 Fixed map tooltip in site dialog
* 6287 Fixed issues with exclude interfaces in DDoS
* 6241 Fixed SSO for custom MKP domains
* 5458 Fixed issues with deleting sites tied to agents
* 4747 Fixed legend table horizontal overflow in Data Explorer

## March 2021

**New Functionality**

* **Traffic Engineering**:

  + Landing page performance enhancements.
  + Inbound and ad-hoc traffic engineering.
  + UI clarifications and enhancements.
* **Connectivity Costs**:

  + Monthly cost tracking history.
  + Multi-currency conversions.
  + Complete overhaul of cost group configuration and visualization.
  + Per-user access control.
* **Capacity Planning**:

  + New computation methodology for Capacity Runout calculations.
  + Now computed using double or triple exponential smoothing, leading to much more accurate forecasting vs. prior simple linear regression.
* **Cloud**: New AWS gateway dimensions.
* **Synthetics**:

  + New kSynth Agent version 0.0.12-rc2 to support advanced HTTPS and API tests.
  + Traceroute explorer improvements to better display path changes.

**Issues Resolved**

* 6766 editing a provider and going back to provider listing looses context
* 6758 Agent-to-agent test does not show name of the public agent
* 6751 Filtered Traffic panel defective
* 6749 Max KBits In not abbreviating to mbits
* 6730 Error 500 when trying to remove/delete tenant. San Carlos Apache Telecommunications Utility Inc.
* 6723 Ping and Trace Sliders are all messed up
* 6721 Traffic is not being classified either by inteface classification or Network classification
* 6719 cross panel filtering not working
* 6713 OTT not working for the last 3hours
* 6702 Wrong View Deleted From Package
* 6696 United-CS (UCS) dashboard data incorrect from Saturday 3pm till Sunday 6pm.
* 6691 IP address sorting not working correctly (see many other bugs)
* 6677 Adding more agents than the page width breaks the network-grid horizontal tabs view.
* 6648 Test Errors
* 6644 Graphs are not displayed correctly when traffic exceeds a certain amount
* 6640 test ignore
* 6639 Alarm Value Missing
* 6618 gateway id has lots of wrong values
* 6613 Grid test errors out for Square Enix
* 6603 can't edit device sending Ip
* 6602 selecting a target agent creates as error
* 6599 It just says error occurred with no feedback about why or what happened. I clicked the button on the
* 6595 Default Time Zone is set to UTC at Tenant SSO Login
* 6585 Need to clear and re-load the SNMP IF descriptions
* 6552 Exclusions filters have dns strings
* 6541 improve the formatting of the milliseconds in this string. ord4-kentik-b8d8cp2 is experiencing a 32
* 6537 popup menu is missing when clicking on a link.
* 6521 Can't chose private agent under the test (Global) Public DNS from DC
* 6520 horizon chart ain't workin on this query
* 6513 Guided Mode Filter does not work in MKP
* 6506 Velocity (Butler Electric) My Kentik Portal error Network 500 when trying to save changes.
* 6503 page won't scroll
* 6499 Statistics on this page are not accurate compared to the v3 version
* 6497 no vertical scroll bar when browser is short vertically
* 6488 providers show list doesn't make sense
* 6486 Some populators are not applied correctly periodically
* 6481 Insights says NaN% failure and shows no data
* 6452 Not possible to increase number of probes per hop
* 6410 Update API Tester Dimension List
* 6378 latency/jitter/loss color green when failing
* 6366 Auto-IC internal suggestion screen unusable as-is
* 6335 NE ASN Quickview Filter
* 6204 SSO information should not be displayed in the profile page of MKP bug customer
* 6128 Sankeys cropped in PDF Exports
* 5745 remove sdm\_leader devices in all of the UI.
* 5280 can't center or right click on links to open in new tab
* 5177 DDos Defense page shows 1 active mitigation, detail page shows 8
* 4497 CDN test not testing all the configured IPs
* 3725 SNMP metrics reporting non-existent interfaces, and traffic far beyond what is showing for device in backend
* 2330 Mitigations isn't active but still listed

## February 2021

**New Functionality**

* **Synthetics**:

  + New ksynth agent v0.0.11 supporting interface bind options.
  + New AWS Wavelength global agent locations in Verizon, KDDI and SK Telecom.
  + [BETA] API for managing agents, tests and results.
  + New “Network Grid” test with bulk IP upload (CSV/txt).
  + Grid view summary of test results.
  + Insights, starting with week-over-week comparisons.
  + Autonomous test improvements to set number of IPs per target and auto-prune un-pingable IPs.
  + Autonomous test result aggregation by Agent and Provider.
  + Autonomous test support for local ASN based on inbound or outbound traffic.
  + Path Enhancements (rDNS, user defined geo, interface traffic, IP prefix).
* **Insights**:

  + *AWS Gateway to Inside Comparison* (Hybrid Cloud): Detects week-over-week changes in gateway traffic to on-prem infrastructure.
  + *AWS Gateway to Outside Comparison* (Hybrid Cloud): Detects week-over-week changes in gateway traffic to outside destinations.
  + *AWS Internet Gateway for Cloud Internal Traffic* (Hybrid Cloud): Detects use of Internet gateway to reach cloud internal destinations.
  + *Through Site Comparison* (Edge): Detects week-over-week changes in traffic composition through each site.
* **DDoS**:

  + Alert and mitigation details now show in sidebar drawer rather than row expansion for easier access.
  + DDoS policies now support user-configured filters.
  + DDoS policies now show a notice when they’ve been modified from defaults.
  + Search/filtering now supported for active and historical mitigations.
  + Mitigation IDs are now linked to mitigation status page wherever they are shown.
  + Mitigation status pages now show complete state history.
  + Fixed mitigation controls (start/stop/pause).
  + Manual mitigation setup ported to v4.
* **Map**:

  + Map details moved from pop-up/overlay to sidebar drawer to provide more real estate and to prevent obscuring the map when open.
  + Map items now show an action menu when clicked, with options to View Topology, Show Details, or Show Connections.
  + Controls previously in the sidebar have moved to a toolbar at the top.
  + Map can now display connection lines based on Site IPs or Ultimate Exit IPs.
  + Links between On-Prem, Cloud and Internet blocks are now clickable with details.

**Issues Resolved**

* 6471 Fixed default landing page selection in new profile page
* 6083 Fixed geo lookups in comparison insights
* Fixed an issue discovered with Pendo configs for trials
* 5360 Handling 65535 dest port situation when using service dimension
* 6421 Cross-panel filtering does not work with bidirectional mode
* 6419 Fixed issues with dashboard drilldown when there’s no metric
* Fixed issues with selected count in drilldown button
* 6072 Bidirectional table in saved views is blank
* 6420 Navigate to “source” filter choice will not pull saved filters on panel
* 5611 Fixed issue with saved views where custom dimensions in tenant configs
* 6233 Fixed IC links in DDoS pre-checks
* 6360 Meshes shouldn’t show scrollbars unless absolutely necessary
* 6232 Fixed network explorer tabbed chart directions
* 6181 and 6257 Link DDoS insights now link to DDoS instead of v3 policy editor
* Added Open in DDoS button to load new Attack Log
* 6161 Through/Other bits not Greek prefixing
* 5634 Showing device creation date in drawer/More Info tab
* 4630 Fixes for editing site architecture from map
* 6221 Show DDoS charts in local or UTC based on user preference
* 6458 Members can’t update their profile
* 6239 Deleting tests used in dashboard panels breaks those panels
* 5989 Sorting of login date doesn’t handle never logged in users properly
* 6349 Allow removal of filters in capacity planning screen
* 6356 Health timeline popover obscures time input
* Improvements to top × results in “Internet box” in map
* 5715 Members cannot edit shared saved views that they created

## January 2021

**New Functionality**

* **Synthetics**:

  + *Syn widgets*: Now available in dashboards and MKP.
  + *Mesh improvements*: High density meshes, Usability and labeling improvements, Preset Cloud Performance meshes are now clone-able/customizable.
  + *Flow-based traffic data*: Now displayed in mesh test results based on Site IP data
  + *UX improvements*: Navigating to failing tests, configuring notification channels, tooltip about propagation delay for IP changes, timeline day-marker/time-of-day labels, private agent version now displayed, paused tests now show historical data, warning about potentially blocked test traffic.
* **Insights** (new  comparisons):

  + *Edge Traffic*: WoW comparison of traffic through ASNs in inbound and outbound paths.
  + *Cloud Traffic*: WoW comparisons of traffic to/from cloud services, regions, cloud providers and more.
  + *SP Traffic*: WoW comparison of OTT Service and Type categories.
* **DDoS usability improvements**:

  + Threshold line colors.
  + Cross-reference between alerts and mitigations.
  + Show all mitigations (including those triggered by v3 policies).
  + Removed baseline backfill card from status page.
  + Display configured thresholds while graph is still loading.
  + Enforce threshold order/values.
  + Searchable alert history.

**Issues Resolved**

* 5903 Browser title improvement for MKP Advanced Mode (IMF request)
* 5938 Added tooltip to explain synth agent public IP change lag
* 5628 Fixed a situational issue preventing dashboard drill down
* 4216 Fixed error message displayed to members when viewing user settings page
* 3988 Fixed Sankey getting chopped (including in PNG exports)
* 5974 Only allow MKP users to alert time window in saved views
* 5619 Added download metadata button to SSO settings page
* 6001 Fixed an issue with include/exclude in Data Explorer
* 5900 Fixed logout redirect URL in MKP
* 5962 Added at least one data source validation to MKP tenant configs
* 5936 Added redirect to subtest when user clicks on a map marker in synth test results
* 5391 Fixed logic when clicking “Add All” in Data Explorer Device selector if some options were already selected
* 4677 Added cloud region links for all cloud providers in Data Explorer
* 5930 Reassigning cloud exports owner by a user when deleting the user
* 5986 Added device manufacturer information to device details drawer/more info tab
* 5995 Fixed error fetching insights for certain companies
* 6040 Missing breadcrumb in test configuration page
* 5325 Improved how include/exclude get combined in Explorer
* 5524 Improvements to Synthetics time range selector
* 6088, 5875 Fixed missing insight descriptions (pending feedback)
* 6138 Improvements to password reset/activation emails
* 4151, 6133 Fixed status indicators when archiving a device
* Fixed subscription editing with guided dashboards
* 6180 Hiding dashboard clone menu item for subtenants

**Other tickets closed**

* 4955 Strange Path display
* 4023 Test showing odd packet loss
* 5459 Test is in critical state when it shouldn’t be
* 6164 Complaints about alert policy shim
* 4782 Can’t read agent names in mesh dashboard
* 5252 Changing primary metric in DDOS Attack Profiles
* 4545 Provide ability to group by site
* 6149 500 when trying to register an agent
* 3626 Star doesn’t work for certain insights
* 6155 Adding device for RTBH only
* 6170 Trial assistance requested
* 3881 Click an agent in map doesn’t do anything

---

© 2014-25 Kentik
