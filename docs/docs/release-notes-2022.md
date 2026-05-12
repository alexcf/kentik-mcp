# Release Notes 2022

Source: https://kb.kentik.com/docs/release-notes-2022

---

Kentik employs a continuous deployment methodology for constant extension and refinement of the Kentik v4 portal and the underlying Kentik platform. Release notes for each successive month of Kentik v4 updates are covered in the following topics:

* **December 2022**
* **November 2022**
* **October 2022**
* **September 2022**
* **August 2022**
* **July 2022**
* **June 2022**
* **May 2022**
* **April 2022**
* **March 2022**
* **February 2022**
* **January 2022**

> ***Note:*** For additional insight into what's new with Kentik, be sure to check the **Product Updates** page of our website (Platform » Product Updates).

## December 2022

**Synthetics**

***Features*:**

* Added details in synthetic BGP monitoring alarms
* Changed bgp alarm row group and sorting
* API: BGP Monitoring: Added support for allowed\_upstreams, a.k.a “Upstream leak detection”

***Bug Fixes*:**

* Test retention display issues
* Safety around health data in SOTI DNS Tab
* Don’t return Synth Tests in global search with test\_status: “D”
* Fixed some of the easier traceroute bugs
* Show agent errors in DNS test types
* Refresh Site select options list after creating a new site
* Update getFlowBasedTestCreditsPerMin to account for max\_tasks
* Alarms count is a string
* Omitting deleted agents when re-saving test
* Test retention issues
* Ensure credit usage refs task\_status
* Agent migration should retain latlng
* Logging for sporadic SignatureDoesNotMatch s3 issue
* API: BGP monitoring, handle absence of origin ASNs

**Cloud**

***Features*:**

* Azure Tenant ID Bin Script
* Set Subscription ID Context in Powershell Script
* Cloud Export Task Status Tooltips
* Azure Resource Graph API
* Guided Demo Allow Azure Cloud Map Access
* Nav Menu Cloud Link is now based on configured clouds
* Added Azure subscription limit for topology fetch
* vhub routeTable in sidebar
* Update “No Flow Logs” Warning Message
* Handle and Report on Request Timeouts
* Azure Metadata Troubleshooting Improvements
* Azure Entity Explorer will have subscriptions selection
* More azure topo stuff

***Bug Fixes*:**

* Weathermap error for companies with no clouds configured
* Azure Entity Explorer doesnt refresh topology and subscriptions change
* Supporting src|dst lookups for AWS VPC Name/Subnet
* VPN Line VGW / TGW fix to remove unnecessary lines between transit gateways.
* VPN Metadata shows incorrect ip for connected device
* Weathermap doesnt show AWS Regions
* AWS Metadata doesn’t load when there are too many links generated.
* Azure-duplicate-topology-fetch call
* Show cloud regions for exports without devices configured in Map
* Fix Empty Azure Resource Groups
* AzureRefreshTopologies cron update
* Only save entity data to azure SA
* Store metadata twice but reset security groups for transfluo

**Core**

***Features*:**

* Updated PagerDuty and OpsGenie logos
* Support error display when field is prepopulated but showPristineErrors is not defined
* Move Guided Mode to not being ‘freeform’ except where necessary
* Audit Log improvements
* Supporting horizontal scrolling in tables
* Send trial support emails to their account team
* Capacity Planning Interface Menu updates
* Enable/Disable Streaming Telemetry for device
* Include ASN in IP Quickviews
* Lightstep end to end tracing
* Caching dictionary in Redis
* Capacity Planning Updates ServiceNow
* Update Report Issue Template for Zendesk -> Github sync

***Bug Fixes*:**

* Fixed form field dirty state with empty array default values
* Opening guided up to freeform interface names, descriptions
* Accessing option group safely now when showing Dimension tags
* Short flow agg window fixes
* Fixed chunking issue for SNMP queries
* Fixed validation of uniqueness of site market name
* Fixed Raw Flow permissioning by default
* Fixed warning-sign and error icons to use color over intent
* Resolve (cloud-related) populator search returning dupes in filter dropdowns
* Fixed site market name validation
* Server-side filtering on interfaces page plus “Null” search issue
* Populator Dialog Fixes for Device Types and UDR rollback
* Only mark as dirty in defaultValue situations
* Fixed Audit Log Date Range filter fix
* Fixed Traffic Engineering for devices with dashes.
* Audit Log Fixed MKP, KMI, Totp and CD things
* Horizontal Scrolling Fixes
* Fixed Query Engine OOM
* More Audit Log fixes
* Fix Filtering with Site By IP with Commas
* Add Support for UDR Populators
* Improve capacity configuration error insights
* Fix how we chunk reagg queries
* Safety check for supported-dimensions call for MKP
* Moved time\_format selector out to main window
* Fix when there’s a query polling rate but no re-aggregation
* Allow long-running queries to return
* Aggressive fastdata chunking
* Use i\_orderby for bytes and packets
* Enforcing SSO required in v4
* Populator Dialog Bug Fix + Enhancement
* Tightened time restrictions on Raw Flow
* Fix up ReAgg Windows when no flow agg window

**Alerting**

***Features*:**

* Amp & Reflection migration script improvements
* Alerting Summary, Insight & Alerting Table Improvements
* Added time range filters to attacks, insights and mitigations
* Adjusted new alerting summary for tenants
* Linking to dashboards in MKP tenant alarm views when they’re accessible
* Cleans up the variable names in the transit shift insight summary table
* DDoS onboarding rework and old DDoS workflow clean-up
* Change button label “Configure Alert Policies” to “Manage Alert Policies”
* Add dry run + company flag to tenant migration script

***Bug Fixes*:**

* Fixed minor silent mode form issues
* Baseline backfill fixes
* Fixed notification channels not being persisted when threshold is disabled before update
* Fixed insights duration on fetch
* Add alarms sortBy and pageToken to alerting API
* Fixed insight rendering when no icon name is specified

## November 2022

**Synthetics**

***Features*:**

* Update supported ksynth-agent DNS types
* Node / UI implementation for ping and trace DSCP config
* DNSSEC
* Ad Hoc / Preview Tests
* Added scamper version filter to agent update logic
* Add more properties for synthetic notification details
* Display Test Task Error
* Buffer for reachability activation window
* Conditionally hide network rust agent install instructions

***Bug Fixes*:**

* Fix BGP test advanced options error
* Resolve error isPreview undefined in AgentTestResults
* Another upstream pruning bug
* Update Denver wavelength site name
* Disable policies where tests don't exist and enable BGP policies on fra1
* BGP Monitoring API missing defaults
* ksynth modal error occurred if closed/reopened after step 3
* Bug pruning as path viz when upstream is first as
* Lowercase env var in azure storage path
* Update TCC default lookback to ensure data exists
* Use results\_first\_valid instead of flow\_first\_seen
* Fix to proper render and refresh of certificate expiry date
* Prune IncidentLogWidget saved config of delinquent values
* Fix trimming of decimal values causing NaN on number with commas
* Fixed error occurred on country fetch

**Cloud**

***Features*:**

* Expand custom dimensions for AWS tags from multiple objects
* Map Search Updates
* Improve Azure Error Detection and Reporting
* Add s3 validation and improve AWS validation and s3 helpers
* Save azure metadata to storage acct
* Rename cloud export type option
* Redis Fetch bypass when skipCache passed as true for azure metadata api
* AWS VPN metadata (Part 1) -> Direct Connections
* AWS VPN metadata for VGW and TGW (Part 2)
* Add more container info to azure storage validation service
* AzureMap vnet/subnet sidebar details traffic tab

***Bug Fixes*:**

* Create test results in error message regarding DSCP
* Resolve error Cannot read properties of undefined (reading 'locations')
* Make Azure Access Checks Optional
* Fix onprem router error occurred
* Config Status Improvements
* Disable Azure Check Cached Responses
* Kentik Maps time display in sync with profile
* Fix Cloud Connectivity Report Create
* Fix env var in path for azure sa
* Hybrid map localstorage saved settings retrieval bug;
* Maps: On-Prem box showing "Unknown Device" - rename to be CPE
* Fix Internet Box Item Selectable Behavior
* Make Enrichment Scope Additive
* Match configured Routers on AWS map to Azure map routers naming convention
* Use 24 hour format for Azure store path
* Fix Cloud Export Update
* Add plan id validation for IDOR fix
* Maps: Topology sites do not offer pop up modal on click
* Only reset sidebar details with config
* Cloud map route table not refreshing

**Core**

***Features*:**

* GA version of user API
* Raw Flow Explorer port to v4
* Add period-over-period in Data Explorer
* Connectivity Cost Updates
* Add Support for Time Series manipulation (minsPolling, reAgg). Support Max of Max queries
* Device SNMP Field Enhancements
* Add hasKMI enabled flag to Pendo account
* KMI reciprocal peering edges
* Re-implement Timeseries Re-aggregation
* GA version of the site API
* Refine layout pref language, add icons to describe behavior
* Support Updates
* Capacity Planning: Hide healthy interfaces by default on load
* Connectivity Costs Label Change
* Add Percent Change to KMI Prefix Origination Insight
* Adding UI indication that profile settings impact spoofed users
* Prevent Pendo init on shared links and add Kube beta status
* Capacity Planning performance improvements
* Change db query metric map id
* Temporary validation for Devices that have dots in the name with BGP

***Bug Fixes*:**

* Connectivity Cost: Display calendar sidebar after fully loading
* Fix saving password fields in device
* Fix older queries that used re-aggregation
* API: make api-node continue if proto file is not found
* Fixes related to Period over Period support
* Capacity Planning fix thresholds changes and remove duplicate add new plan box
* Period over Period safety check
* Capacity Planning: Fix #15230
* Capacity Planning fix thresholds changes and remove duplicate add new plan box
* KMI - real country numbers on Overview
* Removing redundant history change on logout
* Eliminating error messages seen on the "Interfaces" table
* Added missing click handler, and fixed issue with timezones in DE query
* Adding logic to prevent duplicate form error messages from being displayed
* Fix zendesk ticket link
* Env flag to run query on local instance
* Fix interface quick views for devices that have dashes
* Set edate on manual classification
* Change device match check to only if matches are less than the number requested
* Adding missing onprems
* Costs: Column sort reverts to default on state change
* Fix sorting on Interface description in Interfaces page
* Maintaining unknown device\_kvs fields
* Hiding the "kproxy" option in the UI when creating/editing a device
* Rerefix Safari scrolling issues

**Alerting**

***Features*:**

* Migration Script for Legacy Amplification Reflection Policies
* Cleanup of policy dataset editability
* Alerting Table Column Additions & Updates
* Initialize Alerting metadata (supported dimensions) without blocking app load
* Change < 1 min policy baseline conditions check to be < 30 sec instead
* Allow users to edit older notification channels

***Bug Fixes*:**

* Prevent baseline percent values lower than 100%
* Fix baseline lookback unit calculation
* Omit reading synth test on cross-region events
* Hide internal & device-based policy metric options
* Fix to allow subpolicy threshold updates
* Fix restricted baseline backfill
* Fix "invalid date" message for alert event end time
* Silent Mode - Hide Remove button for consistency with the rest of the platform
* Support mitigation and insight event within EVM
* Debug alert not working from alert details
* Show counts for all alert types on MKP landing
* Load missing notification channels onto mitigation method
* Fix insights not silencing properly

## October 2022

**Synthetics**

***Features*:**

* Add synth test cleanup util
* Update Agent Install instruction
* Upstream Route Leak Detection
* Add Edit Test button to test subpage
* Keep old agent\_family on migration
* Handle display of transaction errors statusMessage
* Add Use Local IP config to Grid and IP tests
* BGP: add covering prefix search
* Allow agent site creation without geo location info
* Update App Agent instructions based on new systemd instructions
* Make Customize this test Button a copy config button
* Reset req promise for non-ratelimiting errors on API calls
* Agent migration should have option to retain country, region and city

***Bug Fixes*:**

* TCP port validation
* HTTP test's ping metric health issues not visible on subtest page
* Fix error occurred; do not display floating cloud item
* Test Results fetched twice on load
* Migrate agent route handler fix
* Avoid synth test config port values of 0
* Revert no zero ports change on synth tests
* Validate task list
* Show only tests from company
* Reset req promise for non-ratelimiting errors on API calls

**Cloud**

***Features*:**

* Propagate s3 region to s3 constructor
* Comment out xr cross connections until we can use them
* First launch of connectivity checker
* Complete "Show Path To" (add lines)
* Added missing ENI support for connectivity checker
* Enable Azure v2 Onboarding for Everyone
* Enrichment Scope Editor Help Text
* GA version of `cloud\_export` public API
* Add ability to navigate to Kube map from AWS

***Bug Fixes*:**

* Fix Enrichment Scope Drag/Drop Zone
* Fix inaccuracies and bugs in Kube map
* Fix label on pod to pod traffic charts
* Minified-icons-fix
* Fix error occurred; do not display floating cloud item #2
* Fix very high latency values in Kube map
* Kentik Map CIDR filter bug fix
* Fix error on initial load of Kube map
* Fix for creating custom cloud dimensions
* Connectivity Checker Transit Gateway Peering not working for AWS
* Azure Map Add "Show Path To" -> Part 1 -> Show Path to table calculations
* Performance page tests crash.

**Core**

***Features*:**

* Log redis reconnect exception as warning
* Move report issue modal to popover and send to zendesk
* Expose query helper methods for developer sanity
* Add spoof user id for queries and propagate to queryrelay if present
* KMI Insights
* Relax device name rules
* Supporting AS Groups within aggregate filter queries
* Change zendesk email policy
* GA version of v6 Label API.
* Short-term performance improvements for Capacity Planning
* Enable UI query engine service connection
* Update how we collect metrics around HTTP Requests
* Widget re-organization
* Dynamic Interface Selector filtering by physical/logical
* Zendesk SSO redirects
* Cache Shared Links in S3
* Improve KMI Data fetching performance
* Truncate SFDC Support Case Subject Line to match GitHub Issue
* KMI historical chart bahavior + btn text
* Added kproxy 7.36
* ui-query-engine service [part 1]
* Add ability to force full data for longer queries

***Bug Fixes*:**

* Fix MKP default view
* Fix Split Queries in Shared Links
* Fix up Country Include/Exclude Filtering from DE
* Add "advanced" cookie naming logic
* Fixed for when there are no capacity plans
* Added sudo override for negative permission
* Fix Pub/Sub Redis error handling
* Remove cost field value limit
* Fix error on load of CDN Config
* Fix startenv with cluster mode to with app.js only
* Add a default maxVisibleColumns
* KMI Fix mismatching labels
* KMI Slider label popover bug
* Exclude sdm leader devices from showing in licenses page
* Fix cost groups from showing deleted interfaces
* Fixing shared synth tests
* MKP Fix blank error message when opening views
* Only adjust bucket durations when reaggregating with sum
* Fixing shared synth tests
* MKP Fix blank error message when opening views
* Add ability to force full data for longer queries

**Alerting**

***Features*:**

* GA version of the v6 Notification Channel API
* Change default mitigation time range to "All Time"
* [INSIGHTS] ASN Transit Shift InsightDetails improvements
* [INSIGHTS] Capacity Planning Insights
* Group by severity order

***Bug Fixes*:**

* Use policy CIDR settings in explorer query
* Only show DDoS Attacks on DDoS Landing Page
* Fix breadcrumbs for Add/Edit Mitigation Platform
* Only show tenant policies to tenants
* Fix getPoliciesList
* Fix mitigation links in the Alerting Drawer

## September 2022

**Synthetics**

***Features*:**

* Test results cleanup
* syn-back owns task status
* Add Notes field to all tests. Include validation and popover view
* Set maxScale of 3 to HTTP Stages Breakdown ms
* Add more metrics to synthetic notifications
* Move agent\_impl into shared const and add NETWORK type
* Add BGP monitoring to API tester
* Refactored calculate test credits
* Add light/dark mode legend label colors
* Migrate global agent changes for node global rollout
* Syn form config refactor

***Bug Fixes*:**

* SOTI bug with fetching preset test result details
* Old agent needs to be active for migration
* Fix notes updates and display for bgp-monitor tests
* Fix BGP device APIs
* Missing a reset value
* Perf monitor create test fix
* Fix shared results bug
* Fix update network viewer
* Agent mesh widget options

**Cloud**

***Features*:**

* Azure Map - Complete the vHub story; VPN Gateways, VPN Sites, ExpressRoute Gateways, P2S VPN Gateways
* Azure Map - Add Load Balancers / Application Gateways
* Improvements to Terraform based onboarding to AWS
* Kube: add high density layout and high latency stuff
* Azure Map: Add Vnet Routers
* Change default value for CloudExport.enable attribute
* Azure firewalls
* More Azure topo
* Azure Map vHub - Sidebar VNet Metadata

***Bug Fixes*:**

* Resolve errors in Cloud Map/Show connections - cannot read properties of undefined
* Topology Links not showing detail
* Issue with AWS mapping capability using "show path to" function.

**Core**

***Features*:**

* Further performance optimizations for Discover Peers
* New Audit log service
* Add BGP VRF Route Selection
* Speed up initial cost fetching
* Do not send ids anymore to pendo for shared links
* Prom metric grouping and new metrics
* Removed most native promises in favor of bluebird from portal endpoints
* Site Markets
* Add cross-[within-cluster-]domain session cookie support

***Bug Fixes*:**

* Exclude kproxy from user count
* Allow Members to reset API Token. Fixes #13664
* Fixed 'access denied' error when visiting Settings page for companies with Link Sharing disabled and hiding Audit Log for members
* Session cookies improvements
* Removing prom metrics from GRPC calls
* Fix Promise.resolve typo.
* Add safety check to prom metrics
* Metric count http at top level
* Unchain audit log promise
* Fix link from cost landing to provider detail
* Fix $cost fetching, safely destructuring for when no arguments are used
* Fixed logger initialization, which was breaking exports
* Optimize perf of discover peers
* Settings: Custom Dimensions Text inconsistencies
* Extra validation of LinkSharing metadata
* Fix Nintendo logo
* Properly convert showNetworkExplorerOther to boolean
* Live update device labels + prop error
* Add missing CDN icons
* Add logic to fix dev env session cookie-ing
* Fix mix of allowsCut and realias dimensions
* Prevent redirect to login when user refreshes
* Library collapse lag fix
* Remove upper limit for chunking full data and limit parallel queries
* Added safety and log message when queryRelay returns null
* Fix query chunking ddos
* Change bitmap for site market
* Fix session cookie key in socket.io server handler
* Fixes issues with Guided Mode
* Remove time reset to last hour if query is 'lookback'
* Fixes KMI rankings where user's company doesn't have their own ASN

**Alerting**

***Features*:**

* Policies and mitigations table improvements
* Support baselineBackfillImmediately in policy serialization
* Change default mitigations lookback from all time to 30 days
* Remove useless MKP mitigations page
* Improve Mitigations UI, allow method assign/unassign
* Add Query-Based to Alerting Filters
* [Insights] Network Explorer insights dropdown should show only insights
* [Insights] Rename Flow vs SNMP insight
* Add more data about devices for applicable alerting policies

***Bug Fixes*:**

* Fix template filtering, policy sorting
* Fix Faulty Dimension Validation
* Fixed Policies page drawer not showing notification channels
* Insights page grouping by default and not showing confusing key details
* Fix manual mitigation button text
* Send core/explorer/alarm/:id to Alert Details
* Let the mitigation platform status popover scroll vertically
* Support old dashboard route
* Fix MKP tenant subpolicy / condition errors
* Fix insight request on tenant alert details
* Mitigations table status text cosmetic fix
* Navigating to Mitigations page for specific id shows all statuses by default
* Saving new policy v2 notifications correctly
* Fix attack summaries and filter url
* Don't truncate mitigation comments
* Resolving alarm links to dashboard and explorer

## August 2022

**Synthetics**

***Features*:**

* Roll Private and Global Agent filter fields into one to not be treate…
* Hide allowed dns results
* Align test status look between agent test results and test results
* Display DNS Response Code
* Update critical health icon
* Update synthetics mega menu items and dashboard title
* Allow AS Group selection in BGP Monitor test config
* Add aggregate param to calculate aggregated test metrics
* Add orWhereRaw query to find agent-to-agent tests that targets itself

***Bug Fixes*:**

* Mesh hide removed agents
* mesh safety check
* getAgentTests needs agent-to-agent logic change
* INET Family issue with IP Address tests
* "Configure Notifications" works some of the time take 2
* Look at all policies for notification channels
* Performance deck bugfixes
* Fix DNS results processing

**Cloud**

***Features*:**

* Add links for xr circuits that haven't been processed yet
* Add Virtual WANs to Azure Map
* Add more XR detail to azure map
* VHUB -> VNET Links
* Add Route tables in Azure Map Details Sidebar
* Add Azure topo refresh job
* Add Azure expressroutes metadata
* Add Network Security Group sidebar in Azure map
* Add Azure vwan, app gateways, load balancers, vms
* Add Express Route
* Azure express routes improvements
* AWS Cloud Connectivity Checker Improvements
* Weather Map: logo treatment for cloud regions is inconsistent between providers

***Bug Fixes*:**

* Weathermap Error when showing sidebar details for the link
* Safety checks for source/target link values being -1
* Intermediate VPC not shown
* Fix issues in Kube and remove/replace hard coded components
* Fix AWS Export Save
* When hovering over elements on the Azure map, "Error Occurred" appears
* Kentik Map Error on drill down to site
* Fix when kubernetes dimensions/filters are disabled
* Mcgrawhill AWS topology error

**Core**

***Features*:**

* Revise GRPC metrics grouping
* Creating social media "unfurling" images for public shares
* KMI rank export query improvment
* Add Company Account Team to licenses page
* set isPak on default trial cloudpaks
* Add kproxy 7.35
* Apply the same dependency check safety as in v3
* CDN and OTT workflow improvements
* Capacity Fixes / Feedback
* Show full UTC time on flow tag table's time created/edited columns
* Add support for portal enabled/disabled
* Remove Syn/Topo tabs in NE pages
* Add new kproxy 7.34

***Bug Fixes*:**

* Fix agent API validation
* Making sure we don't clear out user settings
* Upgrade redis and dependency
* Fixed CDN dimension filtering
* Api tester icon
* PR Feedback for Show Account Team
* Fix couple kmi bugs + rename var
* Safety check around settings
* Fix #13830 - Failing Capacity Report Subscriptions
* Fix how we treat falsies in company permission
* Prevent tenants seeing share dialog
* Fix redis upgrade
* BGP passwords cannot contain #
* Fixed loading of AS groups page when company has no AS Groups
* Make sure we persist DE url hashes
* Fixed flow tag port magic no-render
* Fix CDN Aggregate View Header - Fixes #13775
* KMI pre insights bugs
* Fix Shared Links because of Portal Enabled Permission - Fixes #13712
* Update Device SNMP IP Help Text
* Fix depth visualization breaks with empty series for pie chart
* Filter MKP Tenant Policy Options
* Geo HeatMap Cities
* DE Matrix Color Axis
* kproxy Select Control Labeling
* Saved View Tooltips
* Disable Tabbed Chart Zoom in Network Explorer
* Incorrect Azure Enrichment Feature Flag Permissions Path
* Prevent kproxy user from showing up in user list

**Alerting**

***Features*:**

* Require mitigation permissions
* Add device labels to notifications
* Change "View Insight" to "View Attack" on DDoS landing
* Allow 15s, 20s, 30s policy windows for authorized users
* Use alert-api for mitigations requests
* Load mitigation platforms when loading the filter sidebar
* Split up insight and alerting details components
* Prevent users from deleting platforms/methods with active mitigations
* Fix button state colors for dark mode
* Support custom CIDR prefixes for policy dimensions
* Make the new Policy Management live for all
* Flowspec preview on mitigations and alerts
* SilentMode Add Pattern dialog improvements

***Bug Fixes*:**

* Resolve radware platform password length error
* Fix mitigation toolbar action error
* Fix policy type filter + minor fixes
* Fix DDoS filter view URL
* Fix broken alert detail redirects

## July 2022

**Synthetics**

***Features*:**

* Combo Tests - add back 1a / start 1b
* Agent site location part 2
* Hide Credentials in Test Settings
* Force ksynth user for traces
* XR requests have already been formatted
* Control to show BGP reachability by upstream ASN
* BGP Related Copy Changes
* Alert DNS mapping UI
* Patched trx test validation
* More trx test validation
* Load loadUsageData
* New performance dashboard
* Disable/reactivate policies when a test is paused/resumed
* Synth auto update
* New getSynthPolicies
* Phase 2 tls cert UI
* Restrict mimetype on POST and GET of syn objects
* Add notifications to trx tests
* Allow trx tests when not sudo
* Test config.policies refactor part 2
* App agents available on all test types for all users
* Onprem bgp proxy config UI
* apiv6: bgp config check for in prod data
* apiv6: use policy ids to fetch policy model
* apiv6: update BGP monitor test result origin\_asn and as\_path value

***Bug Fixes*:**

* Fix missing/incorrect test.config.policies
* Fix BGP synth state query for newly-routed pfxs
* Fix test update without notification
* Fix handling for BGP monitor tests for EventViewModel
* Fix Performance Dashboard showing wrong test results
* Maintain notification channels bug
* don't try to show a tooltip if points are null
* Dashboard Incorrect Aggregation Values
* Fix agent location filters site incongruities
* Fix notification fetch for syn test
* ignore\_tls field should be editable
* Safety check test config policies
* Update testValidation.js
* config.policies should be an array of id, companyID, userID, synTestType objects
* Fix mesh error
* BGP policy fallback typo
* Bug with adding a policies to test config after update

**Cloud**

***Features*:**

* Azure Onboarding Updates
* Showing confusing text for metadata only option
* Azure Map - Relocate VNet Gateway Connections
* Hiding CloudBGP field on CloudFormDialog
* Show cloud vpc input when cloud provider selected
* Fix - AWS Traffic Link Sidebars with VPN Connections

***Bug Fixes*:**

* Cloud deck InsightsOverviewWidget fixes
* AWS VPN Connection Links Draw All Over the Map
* Lost "details" when moving mouse over the devices in Kentik Map - On Premium - Topology
* VPC view UI has overlapping subnets
* Fix - All AWS On Prem Bridge Entity Sidebars
* Remove unnessesary change
* Kube map - no pods drawn

**Core**

***Features*:**

* Minor EventViewModel updates
* Improved audit log url ignore list
* Add failsafe for scheduled reports if knex connection fails
* Fine-tune the error handling for EventViewModel
* Rename Options to Query to be consistent, fixes #13456
* Capacity - Dynamic Interface Selection Enhancements + design changes
* Added feature request link
* Add IPV6 support to Device IPs
* Removed annoying always-on-top-percentage circle hiding data
* Add call to action to activate account
* Add Copy To Clipboard Button to API Token
* User Menu updates
* Make Other in NE be on by default and add info text
* Link sharing available on dashboards
* Add all ASNs in AS path to asns set
* Link Sharing (public view): Only display the DE viz types that are compatible with the current query
* add parameter to route user to specific demo use case
* Enable api tester link in UI
* Prometheus metrics

***Bug Fixes*:**

* Eliminating bad titles from the table (Link Sharing admin page)
* Fix DB transaction deadlock
* Query changed to handle start and end rank
* Capacity Planning: Firefox filter width fix
* Allowing access to query editor when spoofed
* IBM Dimensions do not appear in DE, when data source is IBM
* Disable kubernetes dimensions/filters when not supported
* Geo market breadcrumb trail gets confused
* Sorting the countries by name for geo market
* Fix up errors when accessing MKP and prevent subscriptions + more
* Fix export png height
* Fix incorrect internal share urls on library and shared links page
* Cleaned up API Tester spacing and simplified markup
* Fix up session sharedUser stuff
* Better address matching for european address
* Change forensic raw flow path
* Add safety around req baseUrl
* Search Bar fixed for KMI
* Fix Export: Transit Traffic Report Missing Graphs
* Avoiding sso fetch for companies that aren't allowed to do sso
* Search for existing account on contact and link that if it exists
* Protecting rules with bad regex from blowing up IC results
* Fixes for 11359
* Tab colors matches the chart line color for Interface Line Chart
* Hide kproxy devices
* Create agent plan for all newly-created companies
* Fix country names in quick views
* Fix interface capacity rendering issue when data is missing

**Alerting**

***Features*:**

* Policy Application Cleanup for categorizing policies better
* [INSIGHTS] Flow vs SNMP aggregate insight
* Show metrics & dimensions on policy template list (#13248)
* Policy status & template library drawer cleanup
* Mitigations Filtering and Grouping and match-checking for Manual Mitigations
* Allow users to create DDoS policies
* Enable policy action
* Minor EventViewModel touches
* Manual Mitigations: Validate the IP/CIDR after the platform-method is chosen
* Insights allow site filter
* Use seconds for Policy Baseline neighborhood radius
* Change Remove to Delete in Policy RemoveMenuItem
* Display ratio conditions on InsightWhy component
* Add synthetic labels to Event View Model response
* Policies table drawer and tooltips
* enable policy management on next
* Add incomplete state to policy form
* Update policy page & dialog titles for consistency
* Only allow mitigations on policy if it has source or dest IP/CIDR dimensions
* Mitigation withdrawal time
* Feature mitigations visualization with bgp monitoring
* Update MKP tenant alerting to use Alert API

***Bug Fixes*:**

* Fix clearing silent mode when editing policy
* Fix RTBH platform device selection
* Fixed evaluation period deserialize in alert policies
* Audit log Policies and Mitigations fixes
* Fix console warnings
* Display mitigation state label in tooltip
* Fixed evaluation period deserialize in alert policies
* Resolve unit test naming conflict
* Fix ratio chart so that it appears
* Fix policy management Evaluation Period and Dimension Grouping field names
* Fix alert-api filter root properties

## June 2022

**New Functionality**

* **Transaction tests in Synthetics**: Transaction tests allow you to simulate a user's experience as they interact with a Web application (see **Synthetics: Transaction Test** on our website's Product Updates page). A "transaction" is a series of actions, running in Headless Chromium, that are driven by a Google Puppeteer script created in the Recorder tab of Chrome Developer Tools. Results, which appear on the Results tab of the details page for each subtest in the test (see **Subtest Results Tab**), include the health status and total transaction time for each agent you test from, and show any screenshots that you specified as actions in the script. For documentation of test setup, see **About Transaction Test Scripts**.
* **Synthetics Health Timelines in the Library**: The **Dashboard Panels** and **Saved Views** in the Library that display results from a single synthetic test now include a timeline showing health status over a time range. Hover over a segment in the timeline to display results from the corresponding point in time.
* **Synthetics App agents now support all test types**: Kentik's "app agent" for Synthetics testing (see **Synthetics Agent Types**) now supports all **Synthetics Test Types**. In addition to the previously supported Web tests (HTTP, Page Load, and Transaction), app agents can now execute all network-based test types (including Network Mesh & Grid and Autonomous) and DNS-based test types. For more information about the change, see **Application Agents Now Support All Test Types** on our website's Product Updates page.
* **Redesigned portal sharing UX**: We've significantly updated the way users can share Kentik-derived information with others both in and out of customer organizations. Sharing features, including visualizations (public shares), reports, and subscriptions are now grouped into a consistent interface that's available across all relevant areas of the portal (see **About View Sharing**). You'll find more on the thinking behind this change at **Completely redesigned Kentik Sharing UX** on our website's Product Updates page.

**Issues Resolved**

* 13053 - AWS Link Summary Sidebar Traffic Queries Are Incomplete
* 13035 - Error 500 on public share
* 13009 - Lost "details" when moving mouse over the devices in Kentik Map - On Premium - Topology
* 12979 - Unable to use public share link for synthetics; 500 error.
* 12967 - VPC view UI has overlapping subnets
* 12964 - NE: Italy appears twice and one link is for GR
* 12963 - Feature Request or Bug. Kentik AWS Maps - Intermediate VPC not being displayed in Connectivity Check
* 12954 - Bug: Showing confusing text for metadata only option
* 12950 - uploadObjects API should not allow HTML uploads
* 12942 - Cloud: Sidebar becomes half page error occurred
* 12941 - AWS VPN Connection Links Draw All Over the Map
* 12918 - Incorrect Destination ASN displayed.
* 12894 - Switching from light mode to dark mode on Kentik map causes the Map to crash
* 12865 - updating BGP monitor check, bad gateway
* 12817 - CORE: CD Populator Search Broken
* 12783 - Syn BGP test IDs are not globally unique
* 12771 - [Alerting]Notification Delivery Unsuccessful
* 12757 - Settings: kProxy "git version" links are dead to customers and point at our code updates
* 12745 - Adjust Session Timeout
* 12736 - Can't create synthetics page load
* 12724 - Sankey graph "Top Peering Site, Dest Connectivity Type, Dest BGP AS Path by 95th Percentile bits/s"
* 12713 - Mitigation: can't edit method
* 12689 - Dashboard: Add panel from DE, panel will not load until page refresh
* 12674 - Unable to create pageload synthetics
* 12661 - Clicking lines on the map, opens panel -- shows "error occurred"
* 12660 - Trx and PageLoad tests health data missing HAR and Screenshot data
* 12656 - Internal error when trying to apply recently added Custom Dimension to query
* 12653 - NE: Show in DE erroneous filter, multiple copies, and broken query
* 12640 - [SYN/BUG] Incident Log missing events
* 12626 - Alerting: v3 history partial data return, page does not load
* 12618 - Mitigation table Target column text overflows
* 12613 - Quick-views interface not showing utilization graph anymore
* 12594 - Customer is not able to add new users with error message
* 12586 - Blaise found that adding a panel to a dashboard, will not populate data
* 12585 - Panel not loading when adding to dashboard
* 12584 - 500 Error when editing HTTPS synth test
* 12537 - SYN: BGP Erroneous Criticality
* 12522 - NE: Graph Hides Itself
* 12484 - SYN: When clicking waterfall shown a blank page
* 12457 - COST: Incorrect Currency Conversion
* 12448 - COST: $10000 Limit
* 12446 - NE: Other Traffic Profile Behavior
* 12435 - CPM: Conversations tab 'Create Test' button issue
* 12364 - DE: Bar Chart Spacing
* 12343 - MKP: Landing Page View Only View
* 12325 - Core - Interface quick-view cutting off / collapsing top visualization
* 12322 - SYN: Clear Specific Prefix On Prefix Change
* 12297 - Mitigations: During 'Add Platform' newly created methods aren't correctly
* 12139 - NE: Chart is missing from interfaces detail page
* 12078 - kproxy: installation guides fail when running debian-bullseye distribution
* 12057 - MKP: Policy Threshold Settings Wrong
* 11987 - Mitigation: RTBH Mitigation Requires both IPv4 and IPv6 Next Hop
* 11972 - NE: Ingress line color doesn't match legend
* 11931 - Kentik Map » Weather Map: Every zoom in/out opens the Details drawer
* 11900 - Alerting: Unsanitary error on details page and incorrect route prefix format in filter
* 11586 - Hide "Cloud VPC" field when "Cloud Provider" is "none" Under Agent Management > Configure
* 11452 - KMI: 500 Error when querying for PCCW (AS3491)
* 11359 - Cost: Graph colors switch when toggled
* 11279 - MKP: Tenant Alerting Visibility
* 10942 - Cloud: 504 Errors on exports page
* 10931 - SYN: Aggregate shows healthy even when nearly all reporting 100% loss
* 10715 - Refresh page after submitting manual mitigation
* 10682 - Observation deck: insights overview panel config not respected
* 10715 - Mitigation: Refresh page after submitting manual mit
* 10451 - Mitigations: Code check on configured IP's
* 10368 - Remove "View Details" on Syn public shares for multiple types of tests
* 10364 - Maps: "kt\_src\_site\_title" not found
* 10157 - Cloud: Missing Azure Traffic
* 10102 - Mitigation: Not able to start flowspec manually with infer in v4
* 10070 - While using the KENTIK-NO-FLOW bucket for metadata collection only, an Export Process Timed out error
* 10063 - Zooming on a graph logs the user out
* 9918 - apiv6: synthetics - internal error on test creation request (KnexTimeoutError)
* 9782 - Create date for a dashboard in the library is not correct.
* 9736 - Group-by dropdown has misleading text
* 9733 - 500 Error updating mitigation method
* 9691 - AWS config status can not display: CPG, after AWS admins modified AWS IAM permissions
* 9664 - DDoS: UDP fragments attack has two filters for source network boundary eq external
* 9613 - Visualization preview for Sankey is broken in profile color palette selector
* 9518 - DE: expired saved view hash
* 9504 - Library: inconsistent name and non-default settings for saved views
* 9462 - Manual Mitigation TTL Broken
* 9222 - KMI: Alphabetical order within Market selector
* 8726 - MKP: 500 when adding policy to tenant
* 8560 - SYN: network mesh test fails on time period > 16h
* 8525 - ASN naming in the Synth Path diagram needs to take into account the overridden name.
* 8315 - Test Control Center » Test Details (results) page » main table: inconsistent columns for ASN tests
* 8283 - yelp has cloud traffic. but they don't have any clouds configured.
* 7908 - Cloud PERF - dotted lines
* 7890 - Got Error when use Custom Dimension IPTP Clients equals Stormwall
* 7627 - No VPC Info
* 7500 - Filter by src or dst in Sankey not working
* 7365 - Cloud:  change name of set up pages to match buttons
* 7146 - Settings: Device Type Dimension Filter Wrong
* 7073 - DE: Sorting by Country does not work properly
* 6864 - TE: "Manual Prefix/Len Filtering" does not do anything
* 6647 - DE: Hamburger Filter Cancel Doesn't
* 6357 - Capacity: Sorting by %utilization doesn't seem to work
* 6313 - Cloud: "Cloud\_service" should read "Cloud Service"
* 5827 - No "Detach" / "Copy Assets on Detach" option in tenant view assignments.
* 5824 - Requested changes to MKP domain/branding settings
* 5815 - NE: Export to PDF too verbose
* 5491 - Cloud: Missing Cloud to On-Prem traffic on map
* 4908 - Can't click on the link between site and cloud
* 4808 - Counters view and link between devices do not agree on BW
* 4610 - BGP Peered from another device is not easily apparent
* 4579 - DE: Clicking apply filters does not apply
* 4572 - Flow metric + sample rate queries are broken in v4
* 4571 - Scroll bar in query panel
* 4518 - Cloud: 404 error when no AWS data
* 4472 - MyKentik: cannot remove Default View dashboard
* 4406 - Filtering on source or dest int descriptions returns no results
* 4154 - Settings: Cannot make changes to kprobe devices
* 4064 - Notifications: Digest settings issue
* 3939 - Issue with Kentik provided site selection and peering shared facilities
* 3020 - Strange candlestick where p5th = 0
* 2933 - Unsupported Insight Configuration

## May 2022

**New Functionality**

* **Synthetics widgets for Observation Deck**: The range of synthetics widgets available for Observation Deck has been expanded, enabling information and visualizations related to both synthetic testing and traffic data to be seen side by side. The new widgets include Synthetics Agents, Recently Added Tests, Credit Utilization, and Synthetic Meshes.
* **Cloud configuration status for Azure**: Already available for AWS, our status checker for cloud configuration is now available for Azure as well. The checker helps with cloud troubleshooting by revealing issues with the configuration of your cloud exports. A summary is provided on the Settings » Public Clouds page, where you can click through to a details page listing the resources covered by your cloud exports and showing the status of the connections from those resources to Kentik.
* **Azure in Observation Deck and Kentik Cloud**: New widgets (for Observation Deck) and views (for the Kentik Cloud page) have been added to provide faster access to information about what’s happening with your Azure cloud resources. The Show control on the Kentik Cloud page lets you choose whether to see views from all cloud providers at once or only from a specific provider.
* **Azure regions on the Weather Map**: The Weather Map in the Kentik Map module now shows the location not only of AWS regions but of Azure regions as well.
* **Exports and density chart in KMI**: Kentik Market Intelligence now features a “density chart” at the top of each tab that illustrates how the relative ranking of providers has changed over time. You can also now export CSV files of provider rankings by Customer Base, Customer Growth, or Peering.

**Issues Resolved**

* 12552 - [BUG] Synthetics IP Address Test
* 12535 - Incorrect ASN displayed
* 12529 - Bar chart for specific customers seems to have a bug
* 12523 - DE: FBD Error Occurred
* 12517 - [Product Updates] NavBar » Product Updates » Filter field: not returning results
* 12509 - Config status table overflows, making lower rows and tables unusable
* 12500 - AS Groups in Web GUI results but not in API results
* 12485 - OTT Classification pending
* 12450 - Maps:  Weather map does not show all sites, but topology does.
* 12442 - Syn: Private agent test itself during mesh tests
* 12441 - Unable to update test. Policy creation fails.
* 12430 - Mitigation Status JSON overflows dialog box
* 12393 - AUTH: Default v3
* 12384 - Cloud Performance Monitor - Inaccurate Agent Count on Load
* 12332 - OTT: kProbes erroneously excluded
* 12326 - Companies can enable/create more policies than their limits allow
* 12324 - Bulk enable/disable policies is broken
* 12319 - SYN - Waterfall Dark Theme
* 12283 - Unable to edit Mesh Test
* 12266 - for capone a lot of this traffic should be aws internal but since they don't have any metadata (and
* 12261 - When opening AWS quick view from observation deck the URL contains "undefined"
* 12257 - Customer configured some Transit providers Like Cogent, Sparkle, Transtelco but the cost associated
* 12248 - Settings: Change link from "Mitigations" to "Configure Mitigations"
* 12236 - Unexpected Origin for 81.218.219.0/24
* 12226 - Alerting: Various Route Prefix Alerting Issues
* 12223 - Cost:  Negative cost data and inconsistent
* 12219 - Settings: IP address column multiplies
* 12213 - map lines are not drawing properly
* 12182 - US West traffic receiving in flow logs, but metadata not collected
* 12177 - NE: Top Talker Options Missing
* 12176 - AWS site defined as type "Cloud" shows up on the "On Prem" map.
* 12169 - DE: Query never loads chart, produces error, or 'no results' message
* 12166 - bug
* 12147 - Mitigation:  Policy not firing attached flowspec method
* 12136 - SYN - Hover repeatedly closing/reopening
* 12133 - BGP Monitor Edit Page reports 503 Error when clicking Save & Start Testing button after adding a not
* 12130 - NE: No Export Button
* 12123 - Major Traffic Discrepancy Across Many Sites
* 12096 - Kproxy:  Issue with effective sample rate
* 12083 - Raw Flow: Query form seems totally broken
* 12076 - Credit Consumption is not doubled when testing bidirectionally in A2A test
* 12040 - [Subscriptions] Email subject in Email PDF (or CSV) does not do well w special chars
* 11985 - Handling units for bits/s in policy threshold conditions
* 11984 - Interface Capacity Inconsistency
* 11944 - MAPS: Color Blocks And Auto-Scaling
* 11887 - [NEXT] KMI: landing page data table needs a better layout
* 11831 - Cap: capacity plans show wrong number of interfaces
* 11830 - Cap: Plans do not appear until you toggle a severity
* 11818 - Syn:  500 on Dashboard Synthetics Panel
* 11807 - DDoS: flowspec mitigation removed from policy when changed
* 11784 - Page Load Waterfall view does not format correctly in Dark Mode
* 11692 - Spoof links cause yubi issues
* 11671 - CD: Populators not applying unless updated without change
* 11574 - [Subscriptions] Export Dashboard is not workingClick on export and the message that the report is being processing
* 11548 - Dataset is running too slow
* 11501 - SYN: BGP v1 Notifications
* 11472 - VPN Demo Situation is a bit of a stretch
* 11471 - "Reduce Internet Traffic Costs" demo situation broken
* 11469 - Connectivity Costs Demo situation is incorrect
* 11351 - Settings:  Multiple provider selections yield single result
* 11333 - Data Explorer:  Include Filter Logic Incorrect
* 11254 - Dashboard export fails
* 11164 - Type on step 4 of the App Agent Install instructions for RPM and Debian
* 11122 - Cap: Initial data is still processing
* 11046 - Service Providers » KMI: Inconsistent order of controls in Filter Bar
* 11001 - Do not show -Reserved AS- in any ranking
* 10875 - AS65430 (Private Use ASN) appears in the Customer Growth table
* 10838 - Cap: Hover Text shows "object Object"
* 10389 - DE:  Simple Traffic Profile occasionally blank
* 10047 - [DDOS] Amplification and Reflection Attack profile is misconfigured
* 9808 - Display table for Customers is locked to a very short height
* 9654 - S3 bucket name not being populated in Kentik config by Terraform
* 9579 - kprobe crashes
* 9486 - Guided Use Case - AWS Tshoot fails
* 9423 - v2 Custom Webhook Notification not working?
* 9355 - KMI Dark Mode theme issues
* 9288 - When selecting a link and choosing metrics "odometer" counters are shown
* 9043 - Getting Internal error500 on Cloud Performance Tab
* 9042 - Issue with reports exported to a csv file: Groups missing
* 8992 - Market Selector panel is only semi-sorted
* 8925 - Baseline which uses "the same hour of every day" does not use the yesterday's baseline value
* 8924 - Baseline that "goes back 1 day from now" with the "comparison data from the same hour of every day" fails
* 8903 - Alert debug does not show baseline graphs for some policies
* 7998 - Observation Deck graphs out of sync on time
* 7797 - Address ping and http latency values getting averaged together
* 7527 - Validate Health Behavior
* 7484 - Discrepancies in AS Names between IP Quick View and ASN Quick View
* 7419 - Insights Alert Graph shows my local time as UTC-5 (correctly), however the actual time is probably d
* 7390 - Recurring syn alerts for recovered tests
* 6642 - $current is displayed instead of the current value
* 6610 - Auto Classification Description
* 6575 - from time to time there seems to be a bug that isn't tagging simple traffic profile correctly. perha
* 6534 - Cannot remove alert after acknowledgement
* 6472 - DDOS Defense threshold changes are not being saved
* 6347 - Capacity Runout value bug
* 6298 - Unacceptable performance for the Traffic Engineering workflow depending on customer
* 6283 - Interfaces page unresponsive for any customer w/ a large number of interfaces
* 6267 - ksynth 0.0.11-rc1 (--bind) option
* 6012 - MKP v4 Alert policy importing incorrect thresholds from DDoS policies
* 6003 - the interface details sidebar spacing for the traffic in/out can be improved when there is the outsi
* 5742 - the ordering on the labels makes it hard to read. i would want to see all the last week and this wee
* 5558 - this plan doesn't seem to be working: "Initial capacity data is processing and will be available soo
* 5544 - Certain lookbacks w/ SNMP metrics return No Results
* 5442 - MKP User Creation Issue
* 5372 - Traffic stats from hovering on link sometimes are not shown
* 5357 - this top threshold is for any of the keys in the top 25 (which is a strange thing). so in this case
* 5347 - Admin Device Status Incorrect (Flows=Yes; Reports=No)
* 5342 - star insight doesn't work on this insight
* 5329 - Palo Alto fields not working
* 5319 - Tables not rendering properly (filed on behalf of a customer).
* 5289 - Hard to read alert emails in dark mode due to coloring scheme
* 5270 - kProbe crashes on v4 devices
* 5221 - Custom Dimension pop's with apostrophes or pound signs break filter logic
* 5209 - guided mode dashboard isn't properly setting the device name filter
* 5196 - Ability to set "BGP: Enabled (Full)" for FlowPaks
* 5176 - Botnet/Threatfeed data corrupted
* 5168 - why do all these kgen devices have No Flow Detected?
* 5148 - Links in alert emails go to v3 rather than v4 portal
* 5088 - the snmp data seems to be too low to trigger this insight.
* 5082 - Clouds and Sites - Total traffic do not match sum total profile traffic
* 5076 - When using "Source and Destination" filter setting - drill down dashboard does not drill down
* 5030 - CapPlan numbers seem wrong
* 5016 - Interfaces device filter persists after leaving page
* 4989 - adding interfaces in capacity planning group - something odd
* 4967 - synthetic insights is not running every hour as configured.
* 4368 - Missing user ID under manage users
* 3792 - DE link sets the right metrics for latency but v4 Data Explorer doesn't support it in the ui
* 3483 - Drill down doesn't work from bi-directional panels
* 3481 - why does adding/removing Destination AS change the UE field values?
* 3053 - flowspec manual mitigations doesn't work

## April 2022

**New Functionality**

Public-facing changes released in April were concentrated in Synthetics:

* **New test frequency options**: All synthetics tests can now be run at frequencies of 30 or 60 mins, allowing customers to include targets that require less aggressive monitoring and that would be cost-prohibitive to test more frequently.
* **TLS Certificate Expiry Check**: Kentik now checks the TLS certificates of pages tested with HTTP(S) and Page Load tests. The certificate expiry date of these pages can be displayed in a new column in the results table on the details page for those tests, and a test that fails due to a TLS certificate error can be flagged in the title pane of the page.
* **BGP Monitor Updates**: The Details page for a BGP Monitor test now includes three types of visualizations, including Reachability (the percent of Kentik's BGP vantage points with routes to the monitored prefixes), AS Path Changes (the average number of AS path changes observed), and BGP Events (BGP announcements and withdrawals).
* **Global Agent Updates**: With the addition of nearly 20 additional global agents in Lumen's network, Kentik now has approximately 250 geo-distributed global agents.

**Issues Resolved**

* 12054 - Kentik requests for Bytes Out or Egress Bytes return no data
* 12006 - Palo Alto LLDP/CDP not working -  cannot identify connected switch ports
* 11958 - Map: Links are purple in Firefox, gray in Safari and Chrome
* 11955 - unable to save changes to a networking device
* 11930 - ICMP port reporting
* 11924 - IC:  Rule suggests interfaces are classified but nothing in DB or in Settings
* 11920 - there was a network classification change around this time. is it working correctly?
* 11918 - No peer, use generic IP/ASN mapping under BGP device doen't skip BGP check
* 11912 - Interface Capacity condition showing incorrectly in alert details side panel, but shows properly in the actual alert details.
* 11895 - Page Load Waterfall "Started at NaN"
* 11873 - Certificate Expiry Dates incorrect or missing
* 11869 - Customer Gateway and VPN Connection Name Vanished
* 11864 - [SYN] HTTP Tests Error / Fail to Load
* 11861 - [COST] Interfaces N/A
* 11860 - AWS Map Selected Node Errors When Attempting to Change Internet Tabs
* 11841 - Syn: BGP monitor sending too many emails
* 11840 - Trying to install synthetic node agent
* 11836 - when editing a device, save button greyed out until bgp session is selected then deselected
* 11824 - [SYN] HTTP Test Agent Details DE Query Filter
* 11823 - Looks like we have another sub interface that is not being picked up
* 11820 - Inconsistent Packet Loss data
* 11816 - Error Attempting to Expand NAT Gateways in AWS Entity Explorer Widget
* 11811 - config.notifications.channels error in configure notifications prevents test from being saved.
* 11810 - CDN: blank page when using kprobe manually classified interfaces
* 11808 - Free tests and route viewer "clone" buttons don't carry over context anymore
* 11794 - Saving a JSON notification gives an error
* 11793 - State of the internet > SaaS Application > Customize this template is blank
* 11790 - how does this get to 10.54?
* 11789 - need alarm set and clear explained
* 11779 - IP Address column appears twice
* 11778 - Cloud: Perf monitor DE query seems to expect contains behavior, which isn't supported
* 11770 - NAT Trace Route Bug
* 11767 - Alerting: Notifications and Mitigation fields do not show values in V4 but do in V3.
* 11761 - AWS Cloud Exporter property aws\_transfluo\_region being blown away when editing exporter
* 11747 - NE:  Clicking AS Groups in DE - full page error
* 11745 - Unable to modify and save any devices which do not have BGP configured, or are configured to use ano
* 11744 - Kentik AWS Map: filtering based on Account ID, does not show VPN, LAGs, DCGs, DCs
* 11743 - when is "Available soon"?
* 11736 - Alerting: Interface Descriptions Missing
* 11735 - connections by site IP doesn't seem like its working
* 11734 - Remove VRF mappings on backend
* 11732 - Safari - Performance Dashboard - Test numbers are outside of the box
* 11716 - Cloud Performance Monitor - Service Name Changes
* 11688 - DE: Custom dimension "kt\_aws\_dst\_vm\_name" not found error
* 11684 - res1,res2 tyo not registering correctly
* 11683 - Settings: BGP Validation Error
* 11677 - seems like the traffic bandwidths are entirely broken on the aws map. see screenshots
* 11672 - synthetics performance dashboard reporting all tests critical since the 26th
* 11671 - CD: Populators not applying unless updated without change
* 11667 - Some populators are not applied correctly (recurrence)
* 11638 - alert shown twice
* 11607 - Settings:  BGP tab on edit existing device has field validation preventing save
* 11591 - Settings: Devices list customize columns changes everything to IP
* 11570 - Group By: Site seems broken in the device page
* 11560 - Public Shares Blank
* 11410 - Data not displayed correctly when source/destination interface is used in filter
* 11404 - [EXPORT] Dashboard 500 error
* 11402 - Source IP not appearing
* 11370 - NE: indicates incorrect number of active, historical, attacks in side panel
* 11369 - Sharing: View has no devices
* 11193 - Notifications:  Custom Headers required for Custom Webhooks
* 11149 - DDoS Defense Query errors out
* 11139 - View Portal Changes
* 11053 - Time Series Chart for Page Load and HTTP test with ping and traceroute does not chart "Avg HTTP Late
* 11052 - Time Series Chart labeled wrong for Page Load and HTTP Tests
* 11051 - In Page Load and HTTP Test (with ping and trace) the Time Series and Table  values do not match.
* 11049 - Cloud Perf Agent Deployed But Not Showing Up
* 11020 - 'Insight definition not found' danger toast when voting 'No' on 'Was insight helpful'
* 10672 - sflow ingress data missing
* 10653 - Data Explorer: 'netflow router' no longer appears in data sources
* 10352 - RPKI status
* 10283 - Webpage Error Message: https://portal.kentik.eu/v4/core
* 10223 - Payload too large error when trying to remove interface groups from DDoS Defense interface exclusion
* 10197 - [SYN] v3 Notification 503's - BGP
* 10194 - V4 Ack\_Req Alerts Hidden
* 10124 - Kentik Insights: "Botnets Yesterday" and "Interface Utilization Spike" are still shown in active lis
* 10048 - [DDOS] Let user control min\_traffic value in attack profiles
* 10022 - Interface Classification Tagging Intermittent
* 9933 - Alarms are not showing in tenant
* 9925 - Notifications v2 receiving empty payload
* 9872 - Access Control: user cannot access to the portal when his IP is included in the allowed IP ranges
* 9849 - Spoof link with alarm page for tenant loads the wrong event
* 9784 - Getting "Error Occurred" when attempting to access the dropdown menu to the right of a dashboard.
* 9752 - Mitigation method excluded IPs shown in v3 form, but not in v4
* 9726 - Performance: Cost takes a long time to load
* 9724 - star insight is throwing an error
* 9686 - the regions are missing here.
* 9556 - Kentik Insights shown as current, but ended 2 days ago
* 9523 - Mattermost support for notifications
* 9509 - Cannot zoom out from empty results view in Insights & Alerting (Insights details)
* 9348 - no flow triggers a lot of zoom. but really its just their normal traffic pattern.
* 9180 - Insight Graph is showing incorrect data
* 9144 - AWS Gateway Traffic shows inbound and outbound. i think this is broken because we removed the source
* 8570 - Time range selection doesn't work
* 8413 - Subpolicy Issues
* 7923 - Agent Deployed In Wrong Path
* 7788 - MegaNetRJ requires that their DDoS Policies exclude the PNI interfaces. We created a Saved List with
* 7642 - Can't Delete v4 DDoS Policy Copy
* 7558 - Interface filter in MKP tenant dropdown
* 6501 - Same problem as we've seen before when UE interface is repurposed and flow goes to zero but the boun
* 6486 - Some populators are not applied correctly periodically
* 6414 - v3 Alerting Responses
* 6365 - kentik insights show a zero on the row expander. i think its because of an alarm id that is shown fo
* 5373 - Full dataseries on Cisco ASA only shows null results

## March 2022

New Functionality

* **Synthetics**

  + *State of the Internet page*: SaaS Performance, Cloud Performance and DNS Performance, which were formerly tabs on the Performance Dashboard, are now found on a separate **State of the Internet** page.
  + *BGP Route Viewer page*: Formerly a tab on the Performance Dashboard, this feature now lives on its own separate **BGP Route Viewer** page.
  + *Performance Dashboard redesign*: Display was revised for existing widgets in the **Performance Dashboard**, including the Test Status Summary pane, the Incident Log, the Recently Added Tests pane, the Agents pane, and the Credit Utilization pane.
  + *New Waterfall view for Page Load tests*: The **Subtest Details Page** for a Page Load test in the Test Control Center now includes a Waterfall view that shows the load order and load duration of each element in the DOM of a tested page.
* **Kentik Map**

  + *Topology Map for Azure*: A completely revised **Azure Topology** map debuted in our Kentik Map module. Similar to our existing map for AWS topology, the Azure map now shows each Azure region and VNet, and allows users to dig into the details of subnets and connection entities such as gateways.
  + *VPC Endpoint support for AWS*: We now display VPC endpoints in the **AWS Topology** view of the Kentik Map. These endpoints, which enable connectivity to 3rd-party or AWS services over private networking, are rendered as network gateways so users can easily view to/from traffic.
* **Protect**

  + *New Alerting page*: Alerts that were formerly included on the Insights & Alerting page are now shown on their own separate **Alerting** page.
  + *New Mitigations page*: A new page has been added specifically to view/manage **Mitigations**.
* **Core**

  + *Revised Insights page*: The former Insights & Alerting page in Core is now exclusively for **Insights**.
* **Kentik APIs**

  + *Synthetics API goes GA*: The Synthetics API is now generally available in the API Tester.

**Issues Resolved**

* 11673 - Cannot create page load test without opening and closing Advanced Options
* 11670 - Receiving Internal 500 error message when loading pages
* 11668 - Network Error Status: 500 - Internal Server Error on Synthetics->Test Control Center
* 11666 - Could not add an agent to a test - Erro 500
* 11664 - [SYN] Credit Usage 500
* 11660 - When clicking search while menu is open, search isn't possible because the menu stays on top
* 11645 - [MKP] All tenants disappered
* 11641 - [DE] Auto-refresh "No Result" Persistence
* 11606 - Alerting: Tenant alarm is visible causing confusion
* 11596 - Synthetic policy thresholds getting deleted on alert PATCH requests
* 11590 - Receiving a 503 Error when trying to configure Notfications
* 11577 - Can't save pageload test with default values
* 11566 - Capacity Planning plan edition is broken
* 11559 - why is the performance dashboard showing alarms but the tests themselves are healthy
* 11547 - fra very slow - bgp test does not load
* 11546 - test stuck in critical state
* 11532 - BGP monitor: shows critical state all the time, while unexpected origin is just in certain point in
* 11530 - UMass - Consistent False Positive Packet Loss
* 11528 - Cap: 500 on saving dynamic group
* 11520 - clear all link doesn't do anything useful
* 11501 - Failed to save BGP monitor test when adding notifications
* 11499 - Error Attempting to Open Kube Node Sidebar
* 11496 - Digging deeper in Kentik Map - AWS, Selecting a vpc, then View details
* 11482 - hongkong v4 tests are failing
* 11479 - synth\_test: validation of alarm activation condition is (still) incorrect
* 11468 - Cloud Performance Monitor error "Network 500 Error"
* 11446 - Editing a cloud exporter results in enable setting changing in front of my eyes
* 11444 - Regex rule test failing
* 11440 - broken mesh because of the ipv4 problem
* 11434 - Cap: 500 Error on Dynamic Group
* 11426 - Unable to modify "Configure Notifications" section for remote-sha4-edge wan synthetic test. Receivin
* 11425 - Alerting: Sidebar Error Occurred "Insight not Found" 404
* 11424 - Settings: Show user IDs in the Portal user list
* 11417 - agent label won't apply when using the config modal
* 11406 - Public Clouds » Configuration Checker: broken in Firefox
* 11401 - [Synth] Healthy Test Shows Failing
* 11398 - Help intrepreting syn test results.  Not sure if data is being mis-read or a bug is being hit.
* 11395 - synth\_test: UI shows different alarm activation parameters than are stored in the DB
* 11392 - Settings:  Device "View Interfaces" does not filter list to device
* 11380 - Alerting:  Metric column does not sort & Information always in percentages
* 11375 - not able to group by key
* 11374 - names of keys gone?
* 11373 - group by key???
* 11368 - Sharing:  Hovering over save button has wrong popups
* 11366 - [SYN] Unable to create BGP test
* 11361 - Settings » AS Groups: Select All button missing from Add Group dialog
* 11352 - Profile:  Found a way around the fake data in the profile section
* 11349 - Settings » AS Groups: Group Pane disappears after closing Add Group dialog
* 11342 - Unable to create synthetic test
* 11336 - Settings » AS Groups: Cannot re-order columns in Group Pane
* 11325 - Sites: cannot delete site
* 11318 - path view hover stats incorrect
* 11304 - Private clusters - Push Private App Agent Instructions
* 11297 - Settings > Interfaces: Toggle Filter Function
* 11280 - MKP: Tenant does not see alerts from its subpolicy
* 11273 - Non-US cluster - Private App Agent Install  - Typo in Region Argument
* 11263 - User profile is not showing the option to restrict the Connectivity Cost View
* 11255 - wrong interface speed in capacity planning
* 11251 - apiv6:synthetics:v202101beta1: Agent Patch method always fails with internal error
* 11247 - [DE] Image Exports Broken
* 11234 - Settings » AS Groups: Cannot scroll on page when list extends below the range of the visible screen
* 11220 - IP address shown in alarm and data explorer does not appear in subsequent data explorer queries
* 11216 - Error Occurred
* 11210 - Settings: Flow tag view in chart and DE are incorrect
* 11205 - Top panel in Botnet & Threat-feed Analysis showing "Error Occurred"
* 11202 - Flow Tags: creating Flow tags in UI v4 does not work
* 11201 - Synth: 15 min test period shows error toast
* 11192 - Public Shares Broken
* 11190 - Notifications not sent when status code check fails
* 11177 - Synth:  "Bad Duration" when creating or editing a Synthetic Test
* 11175 - test
* 11172 - Linkshare Title
* 11166 - Mitigation: Flowspec mitigation displays the same IP two times
* 11147 - FLOW TAGS NOT FUNCTIONING CORRECTLY
* 11142 - Max Custom Dimensions reached although license shows "4 of 12"
* 11125 - synth test: failure to create alerting policy for tests with period other than 60, 120 or 300s
* 11123 - Maps:  Controls must be clicked twice for correct details to load
* 11118 - DDoS: Device Selection Does Not Save
* 11114 - error occurred when attempting to group map health problems by device tag
* 11111 - newreclic's weather map takes a very long time to load.
* 11105 - v4 mesh includes v6 addresses in mesh.
* 11104 - If I click "interface classification" and try to save it there is an error
* 11078 - Demo:  Click to change to demo env doesn't work
* 11077 - Weird wrapping of tables on the Ranking screen
* 11065 - Clicking on 'View Details' from a network grid view does not show traffic statistics alongside the s
* 10999 - interface graph link leads to empty DE graph
* 10962 - Mitigation: not able to take manual control on mitigation or to remove it
* 10943 - Settings:  Duplicate Users
* 10938 - Synthetic agent -bind flag does not work in case of DNS tests
* 10908 - Changing time range from the Kentik Map interface details page doesn't modify the time
* 10886 - device error warning appears when user hovers over link
* 10839 - Synthetics tests using some Kentik public agents were failing,
* 10779 - IC: classified interfaces by the rule are not shown in the panel
* 10706 - Path Visualization Consistency throughout  AWS Map
* 10602 - cloned tests do not include notifications
* 10600 - Capacity planning: Difference in numerical values and graphs
* 10599 - Capacity planning: In NE interface quick view, capacity planning module section does not load after
* 10589 - [SYN] Mesh Payload Error
* 10489 - Synth - Endless looping requests
* 10430 - Cloud Weathermap Zoom Out Globe Limit
* 10086 - [SYN] Time Selector Box Bounds
* 9778 - Filtering options for configured tests
* 9539 - In a new trial, if no network flow data is set up, visiting the Map page shows the "intro" screen bu
* 9479 - Requesting the API call gives incorrect filter operator "˜"
* 9452 - The target (queried DNS record) is not displayed in DNS grid results
* 9437 - Azure Cloud Map - Displays 0 VM
* 9086 - Unable to add a comma separated list of ASN's in BGP Monitor test
* 8391 - Cloud export requires a device type & billing plan when a user configures it from quick view
* 8227 - Cloud route table viewer isn't selecting the most specific route
* 6388 - Create site in add device doesn't work
* 5992 - the screen shots below has the same data but at different KDE resolutions. it seems at the 5 minute
* 5399 - click interfaces, shows next-hops.   click next-hops will not show interfaces

## February 2022

**New Functionality**

* **Synthetics**

  + *Private app agent*: The ksynth app agent is now available as a private agent (Beta) for deployment in customer infrastructure, enabling all synthetic tests supported by the network agent as well as web layer tests such as Page Load.
* **Cloud**

  + *Kentik Map Health Digest*: An indicator in the toolbar gives a count of issues we’ve discovered on your network and opens a popover that provides a high-level rundown. Click the View Problems button to see a list detailing the issues, then drill down to the map to see an impacted component in the context of its surrounding infrastructure.
  + *Kentik Map Layer Selector*: A new layer selector for the Kentik Map lets you choose which categories of overlays to display, including link traffic types, traffic layer types, cloud regions and backbones, traffic utilization, health, and clustering.
  + *Kentik Map legends*: The legends that identify the value ranges represented by link colors are now persistent, so it's always easy to see what the colors mean.
  + *Kentik Map AWS Regions and Backbones*: A layer for AWS regions and backbones has been added to the Kentik Map.
  + *Google Kubernetes Engine dimensions*: The Kentik Data Engine (KDE) now supports Google's extended flow logs for Google Kubernetes Engine (GKE) environments.
  + *AWS Agentless Ingest*: Kentik Cloud customers can now choose to send AWS metadata to Kentik via a REST API while sending flow logs via a Kentik-hosted S3 bucket.
  + *Improved Metadata-only onboarding/settings*: It's now easier to configure an AWS cloud export to collect only metadata from a given account/region.
  + *Enhanced kappa agent for Kubernetes*: Kentik has released version 1.0 of the eBPF-based kappa agent for Kubernetes network performance and telemetry. Improvements in this version include easier deployment, critical performance telemetry (% Retransmit and % Out of Order Packets), and host metadata reporting.
* **Alerting / Insights**

  + *Refactored flow/SNMP insight*: The Flow and SNMP Difference Detection insight has been updated to provide independent analysis of the ingress and egress directions on a per-interface basis and to respond to abnormal differences with a context-dependent message that offers both a diagnosis and specific remedies.
* **Core**

  + *Public Shares*: Publicly accessible Public Share pages can now be created from Data Explorer visualizations and Synthetics test results and shared as links with people who aren't registered Kentik users.
  + *Kentik Market Intelligence*: KMI is a new Service Provider workflow that uses the global routing table to classify the peering and transit relationships between ASes and to identify the providers, peers, and customers for any AS in any geography. KMI estimates the volume of IP space transited by ASes in different geographies and produces rankings based on that volume, thereby enabling users to compare ASes in various markets.
  + *My Kentik Portal landing page*: MKP landlords can now choose a landing page different from the default (see the Views tab of the Add Tenant and Edit Tenant pages).

**Issues Resolved**

* 11081 - Public Shares url does not work
* 11067 - Alert policy disabled due to system error
* 11045 - Upon selecting "Embedded Cache" in a rule action, the CDN list doesn't appear anymore in the typeahe
* 11008 - updateCloudExportTask.js writes sampling\_rate as a string rather than int
* 10997 - 404 on https://portal.kentik.eu/v4/kentik-map
* 10977 - customer is reporting 81 notifications for ddos policy
* 10967 - Kube Map URL 404 error
* 10961 - Mitigation: v4 missing devices in the Edit mitigation dialog
* 10954 - cost queries are causing heavy load on ourselves.
* 10887 - I don't know what OID this is pulling from
* 10881 - Cannot clone a kentik Preset dashboard
* 10876 - Show All/None data-series in DE data-table has big perf issues
* 10832 - When entering an IP address into "Global Exclusions" in dDoS Defense, an error appears.
* 10831 - Strange amount of warnings in the Capacity Planning Overview and wrong alerts.
* 10810 - Cloning dashboard gives 500 error
* 10789 - Clicking on a rule with matches, the initial popover list of device with interface matches count say
* 10780 - IC: new rule randomly changes position when IC page is refreshed
* 10778 - New IC page does not word wrap long rules, causing the screen to not show the right menu
* 10757 - Settings » Interface Classification: Unclassified Interfaces List extends below its dialog
* 10752 - Public Shares » Share Visualization dialog: Share Link field is not suggesting a last segment for the link
* 10747 - Inconsistent AWS tag keys are returned when user enters filter
* 10746 - [CLOUD] Connectivity Test Failure Notice
* 10742 - Unable to get all tenants lists via TenantService API
* 10741 - Add user button disappears
* 10736 - filters are persisting
* 10729 - Flowspec prefixes remain announced, even though Kentik mitigation and alert have cleared
* 10711 - Start Manual mitigation not listed when another manual mitigation is already exits.
* 10710 - Check agent validation for page load tests
* 10700 - Synthetics » Performance Dashboard: Export menu is mislabeled "Share"
* 10697 - browser throws an error when creating aws custom dimensions
* 10689 - Traffic costs
* 10613 - Agents Labels list needs a scroll bar or autocomplete
* 10612 - changing any agent config (even just saving with no change) causes the agent to immediately show off
* 10565 - MKP: company logo in MKP not set and then the Kentik logo is not working
* 10554 - Disappearing Criticals/Warnings when inside Test Details
* 10550 - [CLOUD] Terraform plan is incomplete, and so are manual steps
* 10512 - Settings » Access Control: First Exception Group added is named Migrated and first edit step seems s
* 10490 - numbers on busy link sets add to more than 100%
* 10483 - User Profile » Visualizations tab » Preview panes: Constrain View Type drop-downs to match Preview label
* 10453 - No Ulitimate Exit information seen for a prefix that is being sent to Kentik.
* 10387 - Test results page formatting needs to be fixed
* 10378 - DE: K8s Dimension Error
* 10341 - Artificial packet loss being logged against target using TCP pings.
* 10308 - Inter-packet delay is not functioning under Advanced Options / Ping/TCP
* 10242 - [SYN] RPKI Valid But Test Critical
* 10169 - Syn:  Notification sent in error
* 10106 - A customer cannot be both "Mutual" and "Single Homed"
* 10099 - [MKP] Unable to match complete Inbound/Outbound traffic to a tenant due to limited filtering
* 10068 - fix/improve Customer-Provider filtering options
* 10028 - Settings » Manage Interfaces » Customize Columns popup: dragging columns to change order doesn't work
* 9949 - Market Intelligence catJam
* 9854 - No Panels When Cloning Dashboard from Library
* 9830 - Settings » Sites: Architecture information overlaps Address information when there are two CORE entr
* 9816 - Ranking comparative barchart in Overview seems incorrect
* 9781 - Trial plans provision a flowpak that does not end in 5 so they are not being enforced properly
* 9745 - Library: Kentik preset dashboard cloning does not work
* 9704 - some asn tests doesn't seem to be returning data in the detail view
* 9679 - Main results show 0% packetloss but details show 100% packetloss
* 9616 - broken link in portal
* 9559 - Settings » Interfaces: On the Customize Columns dialog, drag-and-drop "appears" to be an option, but
* 9520 - [Library » Dashboard] Clone Dashboard on contextual menu creates empty dashboard
* 9491 - [esc] doesn't dismiss the search screen
* 9489 - Dashboard Clone From Library Doesn't
* 9484 - Spoof links fail when user ID no longer exists
* 9310 - Dashboard cloning on /v4/library/dashboards/:id does not navigate to cloned
* 9221 - Settings » Devices: Apply Labels drop-down menu is not navigable; cannot scroll or search through en
* 9204 - Sudo-Only Cost Export Button
* 8953 - apiv6: PATCH request on test succeeds,  but has no effect (synthetics beta2)
* 8952 - apiv6: PATCH request does not allow to set http\_valid\_codes to empty list (synthetics beta2)
* 5231 - [GREG in SUDO] wider IC test result popover
* 5224 - Raw flow view - columns resize back and forth a few times per second
* 5192 - Segment Routing checkbox doubled
* 5060 - Trying to submit feedback when the search overlay is on is clunky
* 4925 - Issue with interfaces page and search
* 4261 - hybrid maps - popup window does not occur when clicking on link from a site to clouds

## January 2022

**New Functionality**

* **Synthetics**

  + *Network agent*: Release candidate for version 1.0.0.
  + *App agent*: Added support for rpm and docker (in addition to existing deb).
* **Cloud**

  + *Connectivity Checker*: New workflow analyzes AWS metadata; inventories subnets, instances, and VPCs; and identifies issues.
  + *Services tab in Cloud Performance Monitor*: Integrates traffic and synthetic data to discover which services are used by which VPCs, and enables performance testing for real-time understanding of network-impacting issues.
  + *Custom Dimensions for AWS tags*: Enhanced integration of customer-configured AWS tagging keys and values into Kentik's custom dimension functionality.
  + *Cloud-scale flow log ingest*: All exports to Kentik Cloud now run in individual Kubernetes pods.
  + *Multi-Region AWS support*: Eliminated inter-region data transfer for multi-region exports.
  + *Parquet flow logs*: Kentik now ingests Parquet-formatted logs.
* **Alerting / Insights**

  + *Synthetics Latency/Hop Count Insight*: Detects latency threshold violations and correlates with increases in hop counts in the testing path.
* **Core**

  + *Improved AS Name mapping*: Streamlined mapping of ASNs to AS Names, which speeds overall portal performance.
  + *Data-Explorer Select All/None*: New control enables all or none display of time series in Data Explorer charts.

**Issues Resolved**

* 10674 - Some alarms are not showing the site. Example: Alarm ID:146920913
* 10651 - [MKP] Superfluous metric selection
* 10635 - Settings » Custom Dimensions: Display name of custom dimension does not update until page refreshed
* 10630 - [DDoS] Activation Fields Display Issue
* 10607 - Map clusters aren't unclustering
* 10594 - Can not add populators via post API
* 10593 - Interface reports 130+ tbits of ingress traffic when circuit is 100G.  this is a repeat problem
* 10588 - [CLOUD] Error Loading Status Data
* 10587 - Kentik Map - Show Connections - Error
* 10578 - Traffic not being balanced through ECMP for FLOW Records
* 10577 - Cannot initiate manual Flowspec mitigation
* 10573 - OTT Unclassified
* 10564 - Interface details won't load when I click on an interface
* 10560 - View Dashboard button takes to a dashboard with a broken filter
* 10549 - On the summary of the alerts, once the Plan is set to monitor runout WoW, it should shows the letter
* 10548 - DE: filtering per dst IP address with metric unique dst IP no allowed
* 10493 - "Error 500" when update Saved filter
* 10482 - In DE, cloud dimension are greyed out when cloud devices are selected.
* 10477 - Kentikdemonew does not allow selection of cloud dimensions
* 10476 - Error message when change the Plan Thresholds on Capacity Plan
* 10474 -  Tenant Filter is not applied to the graph in the Preview screen in MKP
* 10473 - JavaScript Error occurs when selecting Profile in MKP
* 10457 - Cap: Payload Too Large Toast
* 10422 - AWS Map Search broken
* 10416 - Can't seem to apply labels in bulk from the Device management table
* 10402 - Synthetics Details page is not graphing where packet loss is 100%
* 10397 - invalid mitigation error, when there is no mitigation configured.
* 10395 - Syn: Route View hangs for 2000::/12
* 10392 - Fix "view insight" page for Synthetics
* 10374 - [DDoS] Mitigation Selection Issue
* 10354 - Create Restricted User - Critical
* 10353 - VTS own prefixes:: BGP mon test
* 10350 - Getting false positive high BW utilization alerts on a 2x10G LACP bundle
* 10333 - Clicking a link in the Capacity Group interface table does not generate a proper link
* 10324 - synth:api incorrect default for settings.trace.count
* 10321 - synth:api incomplete validation of settings.ping.count
* 10303 - [capacity plan details] (was: Status filter toggles seem to fail sometimes) incorrect interface overall status
* 10291 - Major discrepancy in traffic reported (flow)
* 10287 - device health panel is broken.
* 10284 - "Invalid data to Invalid date" when selecting panel time override in synth dashboard panel
* 10281 - i cannot save this test. i get a 500.
* 10267 - UE interface is incorrect
* 10262 - Notification Channel Validation Errors
* 10260 - When I "Email report" for the "Transit monthly report (2022 update)" I am receiving an email with a
* 10238 - Capacity Summary Card alignment
* 10235 - Capacity Planning » Landing page » Capacity Plan list: Bug in Runout config display on capacity plan cards
* 10205 - Hostname dims appear twice per direction in filtering
* 10137 - Dragging a favorite into the Library causes an error
* 10130 - RPKI ROA list stale again?
* 10129 - weathermap traffic is missing src port. we should change this to be https traffic.
* 10073 - color code at top of map
* 10008 - Got error Payload Too Large while saving Capacity Planning
* 9968 - Custom Dimension does not work
* 9839 - v4DDoS Config
* 9757 - interface SNMP queries are erroring here
* 9741 - Displaying Network Error- 500 while editing the saved filter
* 9706 - Library: all dashboards and saved view have today's date
* 9667 - DDoS Mitigation: Issue to select value of 0 in the mitigation method
* 9622 - Clicking Cancel on a Configure Plan screen takes you back to the Capacity Landing
* 9610 - Settings » Interfaces: When you click Clear All to clear filters, any text entered in the Search fie
* 9583 - map opens to a view that is 3x global map
* 9546 - De-pluralize "Days" in Subscription
* 9403 - Email dashboard only sending first widget and not using configured time span.
* 9205 - Settings » Labels: Clicking the Clear all button in the Filters pane doesn't clear the search field
* 9150 - Syn Credit Estimate Inaccurate
* 8935 - OTT logos inconsistent
* 8235 - Are saved searches (in the map) stored locally?
* 8148 - Invalid Country Filter
* 7610 - Cloud: Clicking 'Enter Agent challenge code' does not work
* 6922 - Azure manual configuration has typo, prevents correct instruction
* 6761 - connections lines don't remap properly when clicking a site. only if the side 'details' view is open
* 6527 - Error Occurred when trying to transition the metadata sidebar section into a popover
* 5797 - DNS - PTR Test Result
* 5796 - DNS - SOA Test Result
* 5793 - DNS - MX Test Result
* 5710 - New data points are not graphed when automatic 120s refresh is enabled for data explorer visualization

---

© 2014-25 Kentik
