# Release Notes 2023

Source: https://kb.kentik.com/docs/release-notes-2023

---

Kentik employs a continuous deployment methodology for constant extension and refinement of the Kentik v4 portal and the underlying Kentik platform. Release notes for each successive month of Kentik v4 updates are covered in the following topics:

* **December 2023**
* **November 2023**
* **October 2023**
* **September 2023**
* **August 2023**
* **July 2023**
* **June 2023**
* **May 2023**
* **April 2023**
* **March 2023**
* **February 2023**
* **January 2023**

> ***Note:*** For additional insight into what's new with Kentik, be sure to check the **Product Updates** page of our website (Platform » Product Updates).

## December 2023

**Alerting**

* Persist policy & threshold IDs when saving MKP Packages
* Fix legacy notification usage
* Migrate ddos exclusions into saved filters
* Scaffold EVM for Alert Manager
* Hotfix: Provide meaningful alarm severity in notifications
* Improve render time on Notification List
* Hotfix: Prevent serialization error that stops policies from loading
* Hotfix: MKP tenant view notifications
* Hotfix: Show Custom Policies in Templates
* Hotfix: Accept updated EVM schema
* Allow landlords to save subpolicy notification channels

**Cloud**

* GCP Cloud Regions
* Kentik Kube support for OCI cloud
* Fix extreme slowness when rendering connections for AWS map
* Fix force flag for AWS map topology.
* Adding alibaba cloud regoins
* Deprecate usage of historical AWS data from interface streaming
* On prem kube hierarchy error due to not existing metadata field.
* Circle indicators for healthy/empty Kubernetes symbols in Kube UI

**Core**

* Hide site links on topology map
* Fix rbac migration kmr role check
* Kentik Agent Dummy Capability + Multi Arch Capability Install
* Add Restricted Fetch for KMR Roles
* New Component: Real date when hovering over fromNow date
* Do not allow disabling capability if kentik agent is not running
* Set resource limits for ephemeral envs
* Update icon on Kproxy Agents
* Do not let capability details open unless Kentik Agent is installed
* Kentik Agent Management - Better refresh logic
* Prevent Library table from overlapping rows
* Use correct device reference for cost calculations
* Fix missing site ip links
* Adjust unit tests to match Connectivity Cost fixes
* Fix Misc RBAC issues
* Fix Misc. Hackerone findings
* Fix interface inbound/outbound flow queries
* Fixes Generator View in Dashboards from not updating
* Avoid a disconnect Error on Enable Capabilities within deploy
* Kentik Agent UI Polish
* RBAC M3 label item filtering synthetics
* Fix cloning in Library
* Enable kube default for Pro and Premier plans
* Adjusting Kentik Agent Headers / API
* Library: Fix Kentik Preset labels being shown in wrong places and label saving issues
* Remove legacy user permission checks
* Kentik Agent Management - Auto Refresh
* Fix Streaming Telemetry status indicator
* Library Rework - More library changes
* Fix selecting multiple view types in Library Sidebar
* Kentik Agents use old data when agent is "down"
* Fix kagent ports
* Use static timestamp for baseline backfill test

**Synthetics**

* Fix Agent Results Tooltips
* Use private ip instead of public ip for private agent mesh tests
* Script to update policy status
* Make Filtering agents work across tabs
* UI/Node changes for full credential vault
* Fix Agent Results Performance Timeline Data
* Harden Traffic Flow Query Builder
* Fix Test Results Table Render
* Fix test frequency text in credit/min sidebar
* Synthetic alert URLs not pointing to the right error time
* Adding lookup for targetIp address in agent collection
* Remove ping & trace obj from config if not in tasks
* Saving a paused test shouldn't unpause it
* Update Test Results Lookback Options
* Editing existing test shows no global agents
* Enable Ping and Trace fix
* Regression for onLoad, DnsSec
* dfw1 specific trace limit increase
* Only show Credentials that are used within Tests in credential filter options
* Harden Mesh UIs
* Backfill Incomplete Meshes
* Fixed no agents but select all is checked
* TCW - update form state when mobx form changes
* Fix Agent Mesh Regressions
* Fix switch fieldname in bgp routeviewer config sidebar
* Patch to prevent 500 on mesh results pages
* Update HOSTNAME\_REGEX to allow a single trailing dot
* Bug with active policies when creating 2k paused tests
* Improve Performance Timeline Responsiveness

## November 2023

**Alerting**

* Fix missing policy ID for baselineBackfill response
* MKP - fix tenant threshold condition values
* Fix comparison chart color issue
* Alerting & mitigation subscription fixes
* Extend alerting export sleep and add synth bgp proxy flag to hrd1 nomad config
* Notifications V1 to V2 migration clean-up and prep
* Policy Manager Threshold Vertical Tabs, with Severity Colors
* Stores cleanup: add $mitigations; remove $ddos and $alarms
* baselineBackfill API update for Geography\_src
* Fix default policy notifications
* Indicate "greater than or equals" on policy conditions
* Fix baseline backfill/queryBuilder so that it doesn't coalesce columns for alerting

**Cloud**

* Module circular dependency issue
* Kube Map UI feedback improvements
* Kubernetes map
* Add next to regions to load env correctly for topology
* Azure Subscriptions Summary and CIDR search
* Direction filters fix for link generated to view traffic in data explorer from AWS security group tabs
* Allow demo user to update settings to show default azure networks
* "painted door" page for Kube

**Core**

* Add Syngest calls to MKP
* Kentik Agent - more feature additions / polish
* Fix Profile Password Reset in MKP
* Do lookup properly for Aggregate series
* Kentik agent install instructions
* Improvements to capability status in Kentik Agent Management
* Better handling of errors with agent capabilities management
* Kentik Agent Management alignment tweaks
* Fix logout not redirecting to login page
* Allow previous period exports to complete
* Remove some lightstep oldness
* Remove old, unused dd-trace config
* Only show Kentik Agent Management to Sudo-ers
* Fix Adding/Removing users from Rbac Role Dialog
* Update KproxyAgentModel.js -- add v7.41.0
* Super Agent part 2 (a continuing saga)
* Fix Time Range issues when using 'this month'/'last month'
* ASN Quick-View should be ASN Traffic View. Fixes #19861
* Cleans filter fields properly when using One Chart Per Series in Data Explorer
* KMI Network Detail Markets tab bug fixes
* Use Dimension tag labels if available
* Ignore empty arrays in interface filters/search
* Fix Interface Type filter being hidden in some interface selectors
* Migrate apiv6 nomad
* Update Access Control link text in Org Settings menu
* Library Rework - AdminTable and Labels
* Use Promise.map in credential key migration script
* Library Rework - Miscellaneous subscription and filter related bugs fixed
* Show Super Administrator KMR RBAC roles on some onprems
* Fix Tenant Package Error, with “updateTenant” refactor
* MKP Usability Improvements
* Ensure demo users have Connectivity Cost Viewer RBAC role
* Fix ability to enter demo mode for new trials
* Add RBAC routes for MKP
* Fix adding an SNMP Interface dimension using "Include/Exclude this in a new query"
* Update saved query and filter when site or device name update
* Fix node-jobs
* RBAC - Milestone 2
* Show Raw FPS in licenses page
* Fix Company Create and RBAC Handler Refactor
* Store obfuscated credential keys
* Make sure ASNs are ints for boundary ASN override
* Missing on prem env
* Fix error reports for Show API Dialog

**Synthetics**

* Add RBAC Permission checks to Syngest Pages
* High Density Mesh
* Synthetics Dashboard Widget Migration
* Fix Target IP Search
* dfw global agents can support tcp and udp-echo
* Link to tcc from creds table
* TCC add test id and credential filters
* Fix Checks for Targetless Agent Test Results
* Let all users use v1 test links
* Show cloud region instead of IP in test agent results
* Prevent fetching credentials on app bootstrap
* Final Fix for Agent Test Results with Optional Targets
* Increase timeout for syngest results req
* Fix Test Results with Optional Targets
* Fix DNS Health Data Array
* Fix issue with dnssec setting not displaying properly on test edit
* Syngest UI Refactor - Test Control Center & Test Results
* Align counts in syn status summary with active/paused tests in TCC - changes for syngest
* Fix Test Control Center for Companies with No Tests
* Fix Test Results Map Markers
* Putting credential vault behind feature flag

## October 2023

**Alerting**

* Validate against duplicate policy names
* baselineBackfill use policy dimension name by default
* Support 5min SNMP and <60sec Flow chart intervals
* Fix PDF export for alert details page
* Fix Notification Channel usage list
* Fix baseline backfill
* Use red 'danger' color as alarm color
* Filter Alerts upon summary chart interaction
* baselineBackfill return map of dimensions to kde fields
* Remove dimensions sort for backfill info
* Mitigations: add support for subscriptions
* Add a restricted post route for baseline backfill
* Cloud metrics Alert detail page fix
* Threshold condition improvements: simplify baseline & top keys conditions, comparison direction field

**Cloud**

* Azure Metrics Collection Opt-In Form
* Minor GCP topo map ui improvements
* GCP metadata only exporter
* Alphabetize subnets
* OCI Traffic Charts
* Add sidebar details for Interconnect entities in GCP topology
* GCP vpn gateways
* Don't remove Target VPN Gateways from GCP topology
* Fix build, case mis-match on VPNTunnels import
* OCI vNIC metadata
* Update cloud path computing from subnet to customer Routers through Direct Connect Gateway
* Connection links to Routers should no go through Customer Gateways
* GCP Topology Job Concurrency
* Azure Entity Explorer Query Update
* GCP Topology sidebar traffic widget
* addUSER role to new cloud services POST endpoints
* Enable OCI for everyone
* Add null check to getInterconnectAttachmentMetadata in GCP Topology
* Deprecate interface-streaming Cloud API
* Add OCI as a traffic origination/termination option
* Azure Entity Explorer Widget and Open in Map
* Use names instead of ids for VPCs & TGWs
* OCI exporter object storage flow configuration
* Perf Monitor Services - Only Support AWS VPC Agent Mgmt
* Route Tables - hide route tables at region level
* Connectivity checker - Convert ENI to Instance when generating connectivity checker report
* Azure Connectivity Report Denied Traffic Tests
* Update kube metadata api to generate separate index files for every cluster
* Dedicated interconnects in GCP topology
* GCP Index file points to incorrect latest metadata file location
* Deprecate Performance Monitor Conversations
* Dot not show regions for clouds that are not set up on weathermap
* Added clusters metadata generation for AWS / Azure / GCP clouds
* Fixes for AWS Terraform Onboarding
* Update alert actions to match
* Revert aws tf onboarding region changes
* Connectivity Checker Path missing destination interface on forward path

**Core**

* Use the same buckets for db acquire and queries
* Add support for FBD to Query API
* Graceful Rollout for Ephemeral deployments
* Update credential vault api
* Allow dashboards using synthetics tests results availability and generator views in public shares
* Deck widgets configure button no longer captures drag events
* Add link to kb in support dialog
* Allow 5 min SNMP interval queries when we really need it (Capacity Planning and Alerting)
* Add ability to graph at 15s granularity in Data Explorer
* Add support for bulk create/update/delete for Devices V6 API
* Fix tenant logout authenticationType unset for landlord
* Synthetics Credential Vault Usage
* Add node-jobs service to make it deployable via ktools
* Update boundary asn override number
* Credential Vault UI Tweaks
* Various: Policies typos in various places
* Kentik Agent Management Initial Milestone
* Add extra safety for constructing dynamic interface query
* Prevent going to sudo page when spoofed
* Fix kibana log proc.name
* Update regions map with custom baseURL
* Ability to manually set Boundary ASN
* Send trial issues to support
* Fix Device API v6 to not require unnecessary fields on update
* Display correct currency symbol in cost group drawer
* Fix landing card for when no capacity plans
* Change minimum length constraint for device\_name to 1
* Fix Loaded Detection for Insight Reports
* Fix embargoed country insight generation when there are zero results
* Fix API Tester to be more specific for cost/device swagger
* Increase concurrency for fetching capacity plans
* Fix watch mode
* Credential Vault API
* Add transaction support and removing of consumers for Credential Vault
* Revert capacity planning cache but leave optimization to summary card
* Restore the ability to export subset of devices as CSV
* Fix subscribing to existing report
* Fix/insight report loading
* Add Caching for Capacity Insights
* Fix Capitalization of Data Explorer in DNS blurb
* Settings: Limiting device sample rate (avg) to two decimal places
* Shared DE view site market dimension support
* Add device\_sample\_rate to Data Explorer
* Add Guided mode query param to share dialog
* Loosen up IPv6 Validation
* Fix UI Tests and editing dash item edge case
* Open up Audit Log for spoofers even on essentials
* Update from User menu - 'Logout' to 'Log Out'.
* Fixes sankey label overlaps
* UserMenu and CompanyMenu: Key adjustments for consistency!
* Connectivity Cost API
* Open up PDF exports to all users
* Remove Company Menu for MKP
* Cost fix to prevent incorrect interfaces being used when dynamic groups
* Auto deploy node-services to portal.our1 on merge
* Use node watch in startdev
* Fix usermenu snapshot test

**Synthetics**

* Align counts in syn status summary with active/paused tests in TCC
* Fix Erroneous 'Add Label' warning
* Paused tests were being counted as failing
* Include app agents in global agents
* BGP Monitor Test Keeps Enabling Alerting despite Disabling Alert
* Show incident log when bgp call fails
* Include a2a tests in agent test count

## September 2023

**Alerting**

* Policies - Fix policy drawer missing mitigation details intermittently
* Fix: Alerting table filter errors
* Show if policy is used by tenants, prevent editing metrics
* Make the wording for exact match on alerting filters
* Initialize alerting chart on DDoS landing
* Fix Alerting CSV export not working
* Improve Alerting type subscription time range selection
* Alerts: Changed reference of individual alert from 'Alerting' to 'Alert'
* Fix ascending threshold condition values error calculation
* Add support for labels on policies
* Fix: show notification channels in the Policy table drawer
* Re-adding button for migrating tests to active global agent

**Cloud**

* GCP load balancer api work
* GCP load balancer ui
* GCP fix missing project in load balancer ui
* Azure Storage Account Check Logging Improvements
* Format datetime in firewall policy sidebar
* Expose More Azure Access Check Messaging
* Expand Connectivity Checker to work with azure
* Fix AWS Customer Gateway Discovery
* Add Box Link Traffic Calculation Info Message
* Fix Route Table Target Sorting
* Fix Cloud Gateway Traffic Viz Depth
* Azure Performance Monitor Services
* API endpoint to store kube metadata

**Core**

* Silence noisy error logs about aws-sdk v2 deprecation
* Add support for tab linking to User Profile screen
* Update auto deploy RC to next
* Add auto deploy for develop branch and support #deploy for PR
* Fix hcl deploy script
* Connectivity Cost - Prevents skipping certain providers during calculation due to dynamic config
* MKP: Tenant height and better column widths
* Subscriptions: fix source link for alert type subscriptions in settings
* Simplify auto deploy
* Nomad: Migrate to Shared Manifest
* Nomad: add job suffix
* Nomad: add constraints
* Support Dialog - add Dropdown and update ticket Metadata
* Add Site Market to Guided Mode
* Dashboards: Changing time range does not reload panels if they have not finished initial load
* Subscriptions: Subscribing to an existing report removes and replaces original subscriber
* New CDN logo for i3D's recently added CDN
* Fix share dialog links
* Fetch cost interfaces without locking up node
* Added request type back to support dialog
* Fix Payload Too Large for Device CSV Export
* Prevent MKP spoofing from triggering password reset
* Update to User Menu and new Company Menu on NavBar
* Rewrote fetching of cost interface groups to prevent several individual queries
* Bump app\_count for onprem portal
* Interface group service error handling
* Nomad: Config for ephemeral deployments
* Fix raw fps on devices page
* Empty global agent proxy to fix yubikey in onprem nomad containers
* Add Device API
* Fix device links that were sometimes reversed or doubled
* Interface Group Not Equal Operator fixed in DE
* Parse certain environment variables as booleans
* Switch bootstrap to nomad exec
* Fix Device Editing on onprems

**Insights**

* Report embargo list
* Rework embargo insight filters
* Conditionally add embargo insight to insights report

**Synthetics**

* Create a script to get synth test results from synback
* Brought back sudoMigrateAgent Script and tests

## August 2023

**Alerting**

* "Alerting Health Map by Sites" Dashboard Widget & Endpoint
* Show alert & policy debug button for all users
* Prepend special device names to alerting metric labels
* Move Alerting to top level route
* Omit Synthetics policies from tenant subpolicy request
* Alerting table dimension list-out style fixes
* Subscription UI and workflow fixes
* Misc alerting fixes
* Fix alerting site health widget url
* Remove "&!" and replace with "bitwiseNone" as filter operator
* support legacy protect/alerting routes
* Support all notification channel versions in MKP
* Improvements to Alerting table, filter sidebar, and summary widget
* Rework EventViewModel changes for synth origin details
* Improvements to Notification Channels Settings & Filters
* Fix subscriptions API error for MKP alerting page
* Fix notification channel form error
* Alerting Summary Graphs
* Fix threshold range bug

**Cloud**

* Exclude Internal ASNs from Internet Queries
* Connectivity Checker - update source and target destination links
* Fix "Missing "key" prop for element in iterator" in ConnectivityReport
* Fix Kentik Managed S3 Bucket Creation
* AWS CloudWatch Metrics Panel
* Fix AWS Entity Explorer Widget Search
* Cloud Config Status - Use Device FPS to Determine Flow State
* Show GCP Route Table and FW Rules Sidebar Panels
* AWS Map - Show Path to modal should include destinations from TGW Route tables only for tgw attachment within same VPC
* Fix GCP Cloud Network Sidebar
* Additional Performance Monitor Service Branding
* Support AWS Async Lookup for VPC and Subnet Names in Aggregate Series
* OCI topology - Default parameters for OCI metadata scraper to upload topology to storage bucket
* AWS Map - Update route table destination routes IP addresses validation
* Remove bgp status from cloud export details
* OCI Map - Store historic OCI topology into storage bucket
* GCP firewall policies
* OCI production account bucket secrets
* Update AWS Perf Monitor Service Urls
* Enable Firewall Log Collection on Cloud Exports
* Connectivity Checker for multiple clouds, initial refactor to extract aws logic out

**Core**

* Nodejs 18 update and more
* Update node versions
* Fix Geo Heatmap View when using City Dimension
* Update embargo insight api country filter
* Fix Site Country Name sorting
* Fix sending metrics at login screen cause logout
* Connectivity Costs - Redirect if invalid tab, access w/o permissions, and fix data sometimes failing to update
* Fix missing SVG option in the API Call JSON
* Retry query on query relay application error (503)
* Remove node\_config\_env from common.hcl
* Adds share/exports for Settings->Users, Sites, Devices
* Added kproxy 7.39
* Fix tenant user collection spoof bug
* ktrac: remove pointless aggregate from KDE queries
* Fixed Gauge Preview Formatting Issue while adding a Data Explorer Panel
* Move default preset filter to UserCollection
* Use new version of ip-address lib
* Prep for deploying to BCT1 and DFW1 using nomad
* Remove debug log
* Fix agent api metadata
* Add PeeringDB IX info to Interface Details drawer
* Custom Dashboard Page now has a consistent Button structure with the Synthetics Test page
* Fix Internal Class Name for ViewActions
* Prevent ProviderInvoiceCard from being remounted due to changing key
* Fix zeros in high cardinality SNMP queries
* Core Aggregate page: Time Range - Capitalize This Month
* Link old API Tester to Latest API Tester
* Update to latest name of provider in cost graphs
* Inconsistent placement of parameter controls
* Fixing subscribe button for Capacity and Capacity Detail
* Subscribe Is Greyed Out: Library
* Subscribe Is Greyed Out: all but Library and Capacity
* Fix Saved View Edit Validation
* Subscription Details Dialog prevent a console error on unmount
* OCI Nomad environment variables
* NEXT deploy, onprem envs, and separate service for device interface group and classification
* Auto build on develop
* Users CSV export to use YYYY-MM-DD HH:mm format and devices export available only after loading
* Added kproxy 7.40
* Fix RC deploy
* Move manual cost job to /restricted

**Insights**

* Address insights report errors (Release Fix)
* Insights Report Improvements
* Insights report errors
* Summary page for insight votes, improved feedback flow

**Synthetics**

* Temp remove xregion suppression call to buy more time to find the issue
* Add missing xregion suppression routes
* Fixing bug when icmp port is non 0 on update
* Changed isAdmin to canEditTests
* Add Agent Details to Synthetic Alarms Response Payload
* Rename ‘Performance Dashboard’ to ‘Synthetics Dashboard’
* BGP Synth: don’t filter paths with ASN 0
* BGP Synth: filter private ASNs, improve query performance
* Show alerting status in TCC
* Reachability and leak detection policy updates
* Fresh notifications collections after test update

## July 2023

**Alerting**

* Add notifications to multiple policies
* Custom baseline settings should preserve values
* Fix Policy form - baseline window and condition maximum %
* Allow Baseline Conditions for <60 sec policies
* Export the Alerting list as a PDF
* Subpolicy dimension & form validation improvements
* MKP: Fix to alert policy defaultValue
* Disable bgp type / flowspec toggle for devices with active mitigations
* Fix mitigations checks for policies integration and vice-versa
* Max top keys validation fix: less than *or equal to* max keys
* Update Adaptive Flowspec Traffic Volume text
* Omit missing device from filter
* Support for Adaptive Flowspec Mitigations

**Cloud**

* GCP Map - Link Connections update for densely populated containers
* Azure Map - Update subscriptions selection menu.
* AWS Map - Subnet - TGW routing path drawing update
* Azure Map - Update VPN Link Connector
* Weathermap - Add GCP regions into map.
* Azure Map - Link Connections drawing updates.
* Azure Firewall Sidebar Traffic
* Use consistent wording between breadcrumbs and Cloud Landing page
* Readme updates to new secrets generation
* OCI Cloud integration
* Copy error text from cloud config status
* Azure scroll bar missing on edit view
* Fit link from vHub to Express Route Connection
* Update sidebar details based on latest component state
* GCP firewall rules
* Connect VNET Peering to Virtual Hub and expand its destinations
* GCP firewall rules aka nacls
* Update Cron Job Time Field Labels
* Fix Cloud to OnPrem Sidebar Traffic Panel
* Fix GCP Sidebar in Kentik Map
* Update Cloud Topology Job Logging
* Complex path computing and drawing across multiple VPCs connected through TGWs and VPC Endpoints
* OCI environment variables
* Adds src and dst Azure subscription name dimensions to UI
* Azure Firewall Denied Traffic Panel
* Azure Cron Job Structured Logging
* Check for Azure Unsupported Version 1 Flow Logs
* Fix “Show Connections” When Sidebar is Open
* Fix cloud status config in GCP Topology
* “Show Connections” Additional Fixes
* Fix GCP Map Subnet Connector

**Core**

* Insight service setup
* Correctly sanitise insights IDs in paths
* Email and Entry point fields are no longer required in the SSO page under settings
* Remove dead code querying chf\_realtime
* Remove broken sudo-only ingest host overview
* Preparation for nodejs 18 upgrade
* Revert base image to earlier version
* Fix Cirion logo for Connectivity Costs
* Add missing nomad env config and fra1
* Update fra1 redis secret per env
* Fix and cleanup helpUrls
* Remove 24hr device stats implementation
* Set NODE\_CONFIG\_ENV to next for nomad deploy
* Fix ASN Prefix distribution KMI route for Members
* Fix for generator view having header links in MKP
* Fix for Extract DNS Query Name
* KMI Historical Charts + Prefix Breakdown
* Increasing timeout for dashboards export/subscriptions with several panels
* KMI Ranking Over Time capitalization
* Fix 'Share -> Email Export' having improper default filename
* Comparison Bar label change from greek to standard abbreviation
* Track Search queries and clicks with a custom event
* Fix showing staging branches in hc-full
* KMI Network Detail Rankings tab missing default bug
* Site Settings error on load with row selected
* Connectivity Costs: Add 'Site Market' breakdown on certain detail pages
* Connectivity Costs - Move Cost Group Interfaces to DIG Structure
* Fixed CP raw label\_id and site\_id
* MKP - Fix the policies tab on the Package form
* Updated title on PeeringDB Exchange dialog
* Fix top level connector logic for include exclude
* Add more logging for querie
* Add Redis client name to all redis connections for better debugging
* Connectivity Costs: Fix all monthly charts to use changed data structure
* lightstep to otel transition
* Put lightstep config back so that app can start
* When saving a Filter Group, remove any empty nested filter groups
* Field names from previous device types are no longer being concatenated after a device change.
* Updated verbiage for removing Peering DB mappings to "remove"
* Adding a horizontal scrollbar to Settings > Authentication and SSO
* Added filtering based on country set in Traffic Profile on the PeeringDB ASN Details page
* Remove reference to 7 day deletion after archive.
* Connectivity Costs - Ability to hide suggested providers
* Removed row display options if row was not found in topx
* Subscriptions - Custom filenames and ability to CC/BCC
* Add error trap for Pendo not initializing correctly due
* More fixes around include/exclude in DE
* Capacity Planning: Use bookshelf transactions on CRUD
* Added feature to create sites with address autofilled based on selected PeeringDB records
* Changed the popup label "IX Mapping" to "PeeringDB IX"
* Update metrics aggregate registry setup
* Adding Connectivity Provider CIRION and its logo to the list in Connectivity Costs
* Connectivity Costs - Add service/endpoint to support getting cost snmp data
* Connectivity Costs - Allow contract dates to be in the past up to 10 years
* Log Query JSON objects when queries are issued
* Prevent structured log conflicts with query json debug logging
* Search bar now shows the correct label name after hitting browser refresh.
* Longer dimension names are no longer being cut off in Dimension Selector
* Safely escape content in CSV export worker
* Added a safety check for undefined markers in MapGL
* Connectivity Costs - Reflect changes in chart bar stripes and days remaining when selecting past months
* Connectivity Costs - Add safety default billing date and safety checks

**Synthetics**

* Adding udp as ping option
* Fix for date picker localization issue when timezone is set to local
* Fix confusing messaging around rpki status when alarming
* Fix test type validation
* Fix required rule
* Syn tests silence alarms
* Dynatrace migration script v0
* Enable search filter in labels dropdown
* Synthetics availability library widget
* Default State of the Internet tab to SaaS
* Fixed synthetics tooltip specify typo
* Normalize agent id to string
* Fix synth mesh crash on hover for tooltip
* Fix syn mesh missing columns
* Filter out v3 notification channels if has bgp task
* Assign a default config at higher level
* Adding aws and gcp apac regions to dictionary
* Separate Alerting and Notification
* Made agent uploadObject handler async
* Make AllowedUpstreams optional

## June 2023

**Synthetics**

* Script to update whitespace

***Bug Fixes:***

* Agent listing table should take the entire remaining height
* Issues table cuts off at sub 1800px
* Private IP doesn't sort
* Configuring HTTP test with Request header hidden results with two headers
* Loading Combo test results BGP tab fails with "Network Error, 500"
* Prevent mutating original field default value

**Cloud**

* Azure VNet Gateway BGP Peer Status and Learned Routes
* Azure Load Balancer Backend Pool Load Distribution
* Connectivity Checker and Perf Monitor Static Pages
* Subscription object saved with Azure topology
* Azure Load Balancer Sidebar Panels
* GCP Route Tables
* Azure Load Balancer Traffic
* Azure Load Balancer Updates
* Azure Load Balancers - Demo Feedback
* Reduce Logging Generated by Azure Topology Refresh Job

***Bug Fixes:***

* Fix AS Groups Connections
* Fix VNet Denied Traffic Table
* Fix API status errors in GCP Topology
* Fix Azure Metrics Sum Aggregates
* Fix Azure Cron Job Failure
* Add null check for gcpSampling
* buf1 fallback

**Core**

* Applying release version label to PR
* Workflow trigger labeling PRs
* Correct GH action labels
* PeeringDB Related Fixes with Sites and Interface Drawer
* Rename metrics path to match standard
* Connectivity Cost - Billing/Contract Date moved from Providers to Cost Groups
* Added metrics for first render pass and appReady render
* Bootstrap for Next Production deployments using Nomad
* Added % sign at the end of necessary metrics to differentiate them from non-percentage metrics
* Use postgres read replicas for a couple of requests
* Added Quick View Default Option on User Profile
* Connectivity Cost enhancements/fixes
* Tabs changing URL path parameters for better usage tracking
* Connectivity Costs - Add Site Markets Grouping
* Query API v5 Automated Testing and Structural Fixes
* Nomad migration for portal, QE, and rbac
* Bring back Next url for Show API Dialog
* Unified ASN Rendering Component

***Bug Fixes:***

* Made time selection less picky in custom date and time selector
* Honor advanced interface filtering for Tenant Guided Mode
* Fixed Sankey diagram when results are all zero
* Interface Quick View - Fixing incorrect display values in Metrics Tab
* Remove ASN renderer links for MKP
* Clamp Fast Data to 7min index period
* Added a safety check for undefined markers in MapGL
* Fix broken 'auto' value on Selects
* Subdomain on MKP Settings page has a. in front
* Connectivity Cost - Billing Date changes bug fixes
* Fixed unnecessary scrollbar on Dashboard
* DE: Include row query filter logic
* Fix Performance Analytics card height
* Fix Kubernetes filter that use contains
* Add Missing SSO Field
* Fixed bidirectional filter for Azure devices & disabled appropriate fields
* Fix intermittent errors occurring in non-WebGL maps
* Prevent Connectivity Costs from crashing when a site is deleted
* Capacity Plan: Proceed with capacity plan creation when insights are not found
* Fix incorrect max FPS on licenses page
* Fix using Replica DB for Synthetics Credit Usage
* Allow querying SNMP data beyond the full data retention limit
* Fixed UTC DST issue
* Fixed Filtering Issue in Audit Log
* Fixed Direct Flow Detected indicator
* Reduced size of DE tooltips with multiple ASNs
* Settings users no longer saves previous input
* Adjust UI test
* Connectivity Cost: Improved legibility of <Month><Year> pills
* Prevent Spoofers from seeing password reset dialog
* Fix for multiple site object mappings
* ASN Renderer Fix for onclick propagation
* Fix Table rendering for PeeringDB Facilities and Exchanges in ASN Details

**Alerting**

* Allow URL params as filter overrides for Alerting table
* Show disabled notifications in the notifications selector
* Map Component for Alerting Sites Dashboard
* Alert Table & Wrapper Component for Alerting Site Dashboard
* Load more alerts on the DDoS landing page
* Add summary panel to top of alerting table
* Alerting page subscriptions integration

***Bug Fixes:***

* Fix the dimensions on the Policies table
* Fix the notification channel's delete "used by" functionality
* Fix the initial state for the Policy editor's Baseline Bucket width
* Fix: Do not apply top key validations to disabled thresholds
* Change notification URLs for NRT to use IP address instead of hostname
* Fix baseline condition limitation of 9999
* Thresholds - Mitigations pop-up box too small.
* Alerts: Updated Help button to link to Alert Policies page
* Fix max top keys per evaluation validation

**Insights**

* Expand insights search
* Comparison Insights Report

***Bug Fixes:***

* Fix ui.unique\_device\_names.hit metric

## May 2023

**Synthetics**

* Hide metric tooltip before showing next tooltip
* Added alerting graph to test config
* Hide Insights tab
* Fix URL\_SINGLE\_DOMAIN\_REGEX to must include protocol
* Fixing IP address white space
* Refactor augmentation of mesh data to include offline agents
* Web / BGP Combo Test Improvements
* Synthetics Onboarding support
* New external BGP endpoint for getting active tests
* Validation on uploadObjects agent route
* Limit ping probes to 1 on test freq less than 1 min
* Update BGP Route Viewer Config UI
* Increase transaction test timeout for Emburse

***Bug Fixes*:**

* Fixed missing usage data
* Fixed form section states
* Replace old credit warning message with new banner
* Fix agent selector modal scrollbar
* v2 form updates
* Fix remove agent array find bug in single agent selector

**Cloud**

* Show flow sampling rate in cloud export config
* Cloud metadata caching refactor
* Reword "Case Sampling" copy in Cloud Export config
* Azure Historic Metadata download from storage account
* Allow Multiple CID Arguments to Azure Refresh Topology Script
* Azure topology generation cronjob index file
* Load topology from prev hour from cache if current hour cache is not available
* Azure - added express route circuit traffic on a sidebar
* Better API error handling
* Pull azure vm scale set metadata as well as vmss vms and interfaces
* Toggle Default Network Visibility
* GCP historical metadata
* GCP Region, Subnet Sidebar Traffic Queries
* Change default color scheme from blue to green
* Add env key to regions in configUtils
* Azure Express Circuit Route Tables cronjob
* GCP Map
* Traffic chart for VHubs on a sidebar
* Dead tooltip code for form field

***Bug Fixes*:**

* Fix Missing Azure Filter Fields
* Fix GCP Topology bug where Entity defaults are overwritten if missing
* Demo: 'Access Denied' error for Connectivity Checker reports
* Interface Metadata truncates IP addresses
* Fix Link Connector Regression
* Update Data Explorer Filter Autocomplete to use cached AWS topology
* Fix Box Link Traffic Calculations
* AWS Map with enormous amount of links crashes browser
* Fix AWS Internet Box Paging Error
* Show all pod namespaces and fix incorrect no pods found message

**Core**

* Switch audit log to audit-api by default
* Skeleton plumbing for RBAC permission authorization
* Slightly better DSCP selection UX
* Clean up of Service Provider
* Increasing lookup limits to 250
* Simplify drawer/dock positioning
* Modify start dev script to work with tailscale
* Make build process faster
* Device Dialog Streaming Telemetry tweak
* Ability to manually kick off cost job for company\_id and dates
* Device current stats default values
* Add ASN Lookup to Peering Integrations
* Add address info to site and facilities for PeeringDB
* Auto IC updates
* Fix InputGroup rightElement when it's passed in
* Update fix unit test
* MKP License Card tweaks
* PeeringDB Integration Beta
* More logging improvements
* PeeringDB Interfaces Page improvements
* Show peering details button in Subnav
* Multi-Assign PDB mappings
* Actually enforcing RBAC
* Improve Capacity Planning CRUD
* Add RBAC deploy to GH action
* Switch query api url for next
* Add logging around DNS lookups
* Add new connectivity types for Interface Classification
* Add Password Expiration settings page
* RBAC (Milestone 1)
* Support AS Groups and Kubernetes queries in API
* Metrics: tweak bucketing for a few metrics. Reword descriptions
* Bring back UDR Custom Dimensions
* Adding a "close" button to the Page's Drawer component
* Add API for Creating KB Issues
* Query Cancellation
* Rename Auth Settings Page

***Bug Fixes*:**

* Prevent DNS Lookups from blowing up
* Fix TCP Flags filter logic/options
* Fix red dot rendering for service in LegendTable
* Fix breadcrumbs in API tester
* Capacity Plan export PDF shows all interfaces despite severity
* Fix functional tests
* Fix creating an initial capacity plan
* Fix forward earthly
* Fix Earthly builds to show version
* More earthly fixes for versioning
* KnowledgeBase Issue API Fixes
* Fix error when creating dashboard panels
* Fix submenus that are taller than visible screen height
* Fix impossible SNMP Memory Utilization values
* Fix interface dialog after adding ix stuff
* Select Fix
* Cosmetic user settings fixes and AdminTable cleanup
* Fix Logging conflicts and perf improve
* Fix blank audit and access settings pages
* Remove extra comma from sql query causing error
* Fixing "Other" total overlay on DE
* Updating the Shared Links table to use consistent time formatting
* Properly destroying HTTP-ONLY cookies on logout
* Fixing UX display of icon in the Library row
* Fix SNMP Device Metrics queries over more than 1 day

**Alerting**

* Summary totals improvements for alerts
* Better baseline window/completeness validations
* Show tenant label for alerts based on tenant package subpolicies
* Show readable labels in notifications for kmetric dimensions
* Tenant parent policy preview, required tenant dimensions
* Add support for ingress & egress metrics to View In Data Explorer Button
* Add warning if policy was edited since alert triggered
* Export alerting data

***Bug Fixes*:**

* Alert details fixes
* Fix the display of mitigation names with dashes
* Make sure policy collection is loaded when calculating totals for attacks
* Fix MKP tenant notifications
* Fix NaN static condition in attack drawer
* Alerting table group by dimension and infinite scroll load fixes
* Fix overlapping plotline labels on insight chart

## April 2023

**Synthetics**

* Remove separate page load section
* Allow preview test status
* Update query to include all companies with private rust agents
* Adding response header section to the health settings of HTTP and Page Load tests
* Hack the hack to send BGP queries to the ktrac cid
* Expanded Allowed DNS Results
* Flip v2 form config routes as default
* Query which customers have active app agents AND rust agents both running tests
* Add TEST\_STATUS\_PREVIEW to StatusMap
* New logs for aws services

***Bug Fixes*:**

* Filter out duplicate alerts from results by default
* Display --- for NaN values
* Add missing Transaction Time Threshold
* Add missing Response Time Threshold Health setting for DNS Tasks
* Label fix
* Fix syn display in notification settings
* Include missing HTTP related Health fields for PageLoad tasks
* Dark mode legend item text is invisible
* 404 Test 'Run Preview' (Mesh) doesn't load results pages
* Fix rendering of invalid test page
* Remove whitespaces before splitting values to array

**Cloud**

* Allow ap-northeast-2 region to be used for transfluo
* Connectivity Checker Nav bar
* Add support for nested AWS metadata retrieval roles
* Update cloud status page to support AWS metadata retrieval roles
* More AWS logging changes
* Historical Metadata Migration from interface-streaming to node
* Direct Connect Public-VIF on-prem Router
* Load Azure Historic Topology from storage account
* Use https proxy config for AWS
* Update AWS node metadata retrieval to use workers

***Bug Fixes*:**

* AWS Exporter default regions/roles to assume while scraping
* Cross site traffic not generating lines in maps
* AwsVirtualInterfacesPopover Sidebar Error
* Links generation taking too long
* Azure Map -> Show Connections error when connecting horizontally aligned subnets
* Error when trying to open sidebar when clicking on traffic line
* Show connections from internet box doesn't draw links
* Sidebar error on Transit Gateway Peering Attachments details
* Use application/json as Content-Type for topology gzip files uploaded to s3
* Slightly Loading Performance Improvement
* Improving the presentation of flowing line of Site to Site VPN
* AWS Map Error - Failed to execute 'getPointAtLength' on 'SVGGeometryElement': The element's path is empty
* Demo Env - Getting an Access Denied error trying to execute azure resource graph query
* S3 bucket env key should be lowercased
* Fix Azure Refresh topology saving to storage account

**Core**

* GH action to deploy portal services
* GH action deploy RC updates to next
* Deploy portal via GH action
* Enable manual deploy for auto deploys
* Ignore metrics in audit log
* Add User-Agent header for queryrleay client
* Use Query V5 3/4 for Next
* Add query relay count metrics
* Add support for subLabel property
* Destroy transit advisor idle connections
* Add authentication method to pendo
* Add GH action to sync master to RC branch
* Update GH action to fix warnings
* Query apiv5 middleware ports
* Remove test rc branch check
* Update annotation options in API node GRPC server
* Support when DB is in read only mode

***Bug Fixes*:**

* More logging safety
* Fix color picker when no color on label
* Use correct intf id attributes for selection
* Properly support interface group not equals filtering
* Modified link share hover-over
* Open up Raw Flow to members
* Fix Average metric calculations
* Fixing AdminTable column customization
* Persist hash when opening shared view
* Fixed settings getting washed out
* Adding auto-scrolling to the tabs wrapper for small screen sizes
* Fix timestamp issues when AS Groups was used with live update
* Fix SDM issues post PG15
* Add check for filterField and value
* Streaming Telemetry healthcheck fix
* Revert object validation and check for fields only
* Fix sorting of FPS metrics
* Use nonCloudCollection to exclude cloud devices from SNMP checks
* Add mkp logging route and fix metrics route
* Update GH action version
* CSS changes to allow more UX-friendly scrolling of code snippets
* Add milliseconds to raw flow timestamp formatting
* Fix Fastly Potential Peers page
* Disable incompatible options in metrics dropdown
* Fix selecting and clearing in cross panel filtering
* Disable dashboard save button when the query has errors

**Alerting**

* Create & View Tenant Mitigations
* Improve local text search for alerts
* Make flowspec details dialog consistent
* Allow true clear of mitigation status filter
* Set non-Archived default states for ActiveMitigationCollection
* Move /policies from views/protect to views/alerting
* Add metrics routes for MKP
* Infinite scroll on admin table

***Bug Fixes*:**

* Fix policy delete and policy baseline presets
* Fix policy comparison direction fallback values
* Fix policy auto-calc behavior
* Fix: support big numbers for flowspec rate limit traffic action
* Fix dimension cidr defaults
* Fix Alerting and Mitigation sidebars if passed id
* Don't error on unrecognized mitigation types

## March 2023

**Synthetics**

* Prevent refreshing entire page when agents list loads
* Enable test preview for all users
* When cloning a config with a default value, set defaultValue to clone
* Move labels value hydration responsibility to TestLabels component
* Refactor implementation to use metadata.capabilities
* Form unit tests
* Average health by provider group
* Adding noFloats validator and editing transform of ms for specific values
* Adding filter for AWS WL agents if it's a dns test
* Introducing capabilities array on agent.metadata
* Add v2 Add Test page
* Update agent selection for DNS tests based on metadata.capabilities
* Preview button updates

***Bug Fixes*:**

* v2 test config bugfixes
* Fix link to IP address details
* More v2 test config bugfixes

**Cloud**

* AWS Direct Connect Gateway - show account ID
* Update Azure Manual Onboarding text
* GCP Metadata API Private Key
* Added subscription group handlings
* Remove sudo flag for AWS/Azure autogenerated CDs
* Update Subscription ID Reuse Rules
* Enrichment Scope Task Status Refresh
* Update Kentik Kube query for namespaces

***Bug Fixes*:**

* Azure Subscriptions panel is opened on page load
* Fix merge conflicts after update to Azure populator changes
* Fix Cloud Perf Monitor update tests
* Fix error loading pod topology
* Fix for multi tenant Azure topologies

**Core**

* Cron Job README
* Add new Device Fields to serializer
* Allowing dots in device names for BGP devices in non-onprems now
* Added restricted route for dictionary cache busting
* Update GH pre-release action version
* Kick the hydras and use\_pg\_query
* Add inactive company middleware for node api calls
* Update changed devices by ID api to not include existing devices in notfound
* Subscriptions - Send emails for uncaughtException, unhandledRejection errors, invalid subscriptions
* Add Support for Bug Bounties
* Add Support for contains to Hostname filter
* Reinstate generation of swagger file for API tester
* [APIv6] Fix log level labels
* Added kproxy 7.38
* KMI API - Remove unnecessary complexities/fields
* Adding telemetry hooks around Dataviews for future DevTools plugin
* Add Capacity Planning API
* Improved logging for OTT, costs, and reports jobs
* Browser metrics
* Open up Auto Populating Cloud Custom Dimensions for everyone
* Add UI Events for Slow Queries
* Add v4 endpoint for changed devices by last since and IDs
* Interface classifier error metrics
* Cost group monthly traffic chart bounds and lag improvements
* Rename 'Network Map' to 'Logical Map' on provider quick-view
* UI metrics improvements
* Clean up references to long running forks

***Bug Fixes*:**

* Ability to change visualization types during live update
* Add Flow Start/End and Duration to Raw Flow
* Remove unnecessary filters in older queries
* Metrics - Interface Classifier Percentage
* Dynamic Interface Selector: fix comma usage in regex; increase debounce to 1s
* Make API Tester available for customers with no flow
* Show parametric options for misconfigured subscription form and sanitize filename
* Fix public shares not initializing app
* Fix interface capacity chart
* Add box-shadow on both left/right drawers for consistency
* Fix KMI Export mutual customer labeling
* Fix createDevice based on hydra changes
* Fix device summary queries for no SNMP devices
* Fix Entity Explorer widget button getting squashed
* Fix node bin script help
* APIv6 Fix/update api test
* Fix excluding DHCP dimension values in Data Explorer
* Do not limit query results over API for table visualization type
* Prevent ELK from blowing up with a ton of deep objects
* Open browser logging to member roles
* Fix SiteMarket Site count
* Fix setting category on preset dashboards
* Fix validation of dates in with positive UTC offsets
* Fix spiky Express Route Metrics chart
* Device form fixes
* Allow wrapping of content in profile's settings
* Improve Observation Deck widget header colors for dark mode
* Fix Kubernetes queries with missing service data
* Fix JSON circular error in logging
* Fix bin/bootstrap
* Remove unused tabs in DeviceDetails
* Replace util.inspect JSON stringify circular fix with modified version of JSON.stringify

**Alerting**

* Lock mitigation method and platform edit in case of active mitigations
* Improved policy condition values and labels
* Support smaller screens on the DDoS landing page
* EventViewModel minor tweaks for sub-policy sourced alarms
* Modified Mitigation Splash Page
* Remove unnecessary breadcrumb

***Bug Fixes*:**

* Fix bugs with flowspec discard
* Date picker not restoring hashed state properly
* Tenant config bug fixes & improvements
* Fix alert drawer headings
* Hide capacity insights from family select & config
* Date range filter and query param fix for mitigations
* Fix DDoS Configuration step 2
* Fix insights widget on the observation deck
* [Insights] Fix the time range on the Flow and SNMP Difference chart
* Fix active mitigations summary count

## February 2023

**Synthetics**

* Updating BGP test description
* Test Config Redesign
* Disable edit http target field for page load tests
* Remove whitespaces from comma separated allowed DNS list
* Route Viewer Changes
* Remove BGP events timeline and table
* Add target-correlated synth failure insight

***Bug Fixes*:**

* Fixes error after "Hide Timelines" && bgpEvents from 'showAll'
* Fix BGP Monitor Test typo
* Fix to sort by date on Last Updated column
* BGP Preview Tests - Labels Not Added when Created
* Cannot close BGP route viewer window
* Single url validation
* Update URL\_SINGLE\_DOMAIN\_REGEX to support paths, ports, and querystrings
* Show url on hostname subtest screen
* Web tests should only allow a single target
* API: Improve reporting of "out of credits" errors
* Fix bgp queries, bug with target table rendering
* Use updated route viewer for BGP monitor test preview
* Error message when trying to view Insights tab on a Network Grid Test

**Cloud**

* Update Subscription ID Reuse Rules
* Enrichment Scope Task Status Refresh
* Add global proxy config support for AWS SDK
* Autogenerate custom dimension for AWS
* Weather map: Azure Region Summary Sidebar Section
* Use Resource Graphs for Azure Entity Explorer widget
* Add pop update to AZ pull

***Bug Fixes*:**

* Resource Graph query execution should work for multi tenant customers.
* azureRefreshTopology cronjob causing 500 errors when uploading data to storage account.
* Azure topology errors
* azureRefreshTopology cronjob error when autogenerating populators
* Fix AWS List Bucket Validation
* awsBucketTest only looks for AWS exporters
* Trim cd desc to 50 char max
* Pods are not shown correctly under Clusters
* In / out utilization not showing for express route circuit peering
* Fix azure custom dimensions not autogenerating populators in prod
* Azure metadata saved to storage account using numerical day of month
* Connectivity Report Error for non existing entities
* Fix selecting various Kentik Map nodes
* Update Azure topology subscription filter UI
* Fix bug in 15833
* Reload and Cache most common azure resource graph queries on azure topology refresh
* Show connections on cloud provider menu doesn't do anything

**Core**

* Add safety and logging to quicksort
* Add company id to realtimebus metrics
* Add metrics for query count and chunk sizes
* Preliminary logging route
* Add CID to http request
* Convert query engine to structured logging
* KMI api and unit tests
* Add device column to Detected Embedded Caches table in configuration
* Connectivity Cost - Change 'Configure Providers' label
* Project Recon (behind feature flag)
* Add BGP device getter to restricted API
* Admin table column width fine-tuning
* Device BGP Additions
* Password Expiration flow
* Adding extra data to errors and report issues
* Disable and collapse unsupported metrics
* Disable salesforce sync - rely on Marketo
* Send browser trace as protobuf and fix starttime span
* Add AS group CRUD API and tests
* Add support for invisible captcha in signup
* Interface Redesign
* Adding tenant IDs to the MKP table
* Add Azure instance ID as guided mode dimension
* Better formatting in custom applications table
* Add as group API to tester whitelist
* Save last selected severity filter per capacity plan in localStorage
* Add safety and logging to Quicksort
* More password design tweaks

***Bug Fixes*:**

* Fix JSON circular error in logging
* Fix query errors with AWS dimensions
* Fix structured log
* Fix Live Update Queries
* Raw Flow fix columns showing empty values
* Dashboards: Mesh widget not responsive to resizing window
* Connectivity Costs - Fix #15822
* Fix use of Hostname dimensions
* Capacity Planning - Fix capacity summary filters may be incorrect on load
* Query fixes for 4.27
* Misc query fixes
* Fix Password Reset
* Expand Prometheus Metrics Bucket Range
* Dark mode device status indicator colors
* Fix for some DE Fields showing empty values
* Disable wide query runner forking
* Fix empty filter dropdown labels
* Fix enabling and disabling 2FA items
* Fix TOTP QR code scanning in dark mode
* Add dark theme to matrix chart
* Fix time selection for synthetics dashboard items
* Fix broken scrolling in dynamic/static interface group selector
* Fix Lightstep handle old trace data
* Prevent unnecessary dimension error messages
* Fix excluding empty country from results
* Fix raw flow metrics
* Stop filtering on flipped app protocol values
* Fix for onboarding layout
* Fix AWS Subnet name queries
* NaN displayed in Capacity Plans
* Network Explorer - fix handling incomplete interfaces
* Fix missing SNMP site name in results
* Fixes for password reset flow
* Fix Safari regex
* Fix More Info tab in AS Group details page
* Fix issues in period-over-period exports

**Alerting**

* Remove support for tenant/duplicate severity thresholds
* Notification improvements from a UX and technical standpoint
* Sort policy template dropdown alphabetically, show type filter
* Omit details on inconsistent synthetics notification
* Show filters on Policy Summary Panel
* DDoS code cleanup
* Preserve dashboard and data sources when cloning policy
* DDoS landing page improvements
* Clean up legacy tenant thresholds
* Gracefully handle legacy tenant config alerts in attack and policy requests

***Bug Fixes*:**

* Fix baseline backfill queryBuilder
* Fix TCP Flag dictionary entries
* Update baseline settings wording to reflect KB
* Fix for custom dimensions breaking dashboards, 15515
* Fix Threshold Top Keys Evaluation Validation
* Don't apply custom dimension filter overrides in tenant insight queries
* Update API eventViewModel Alert links
* Fix heading color
* Policy debug fixes

## January 2023

**Synthetics**

* Add details in synthetic BGP monitoring alarms
* Change bgp alarm row group and sorting
* API: BGP Monitoring: Add support for `allowed\_upstreams`, a.k.a "Upstream leak detection"
* Show agent errors in DNS test types
* API: synthetics `dscp` attribute
* More AWS logging changes

***Bug Fixes*:**

* Error message when trying to view Insights tab on a Network Grid Test
* Portal not showing data for an agent, while data are available in API
* Cant Save a SYN script after adding a Puppeteer Script
* Transaction test frequency UI error
* Unbreak service, and style crazy AS names
* Dashboard messing with TCC health
* Page load ping/trace expected credit usage bug
* Incorrectly labeled settings in Configuration pane
* After Run Preview, back button goes to another preview
* Add hover text of a test's status to icons on tabs
* Changing selected test results in error
* Transaction Test Unstable [internal]
* Alarms count is a string
* Don't return Synth Tests in global search with `test\_status: "D"`
* Fix some of the easier traceroute bugs
* Syn show dns test agent error
* Refresh Site select options list after creating a new site
* Update getFlowBasedTestCreditsPerMin to account for max\_tasks
* Omit deleted agents when re-saving test
* Test retention issues
* Ensure credit usage refs task\_status
* Test retention display issues
* Quick fix for getting active vps
* Shared display bugs
* Test "Save" button incorrectly disabled
* Fix Certificate Expiry formatting

**Cloud**

* Set Subscription ID Context in Powershell Script
* Azure Tenant ID Bin Script
* Cloud Export Task Status Tooltips
* Azure Resource Graph API
* Update Cloud Export Disabled Status Tooltip
* Nav Menu Cloud Link is now based on configured clouds
* Add Azure subscription limit for topology fetch
* Add Sampling and Sampling Rate to UI Cloud Config form
* Azure custom dimensions
* Add support for Azure api CRP rate limit handling

***Bug Fixes*:**

* Cloud: performance Services tab will not load
* Supporting src|dst lookups for AWS VPC Name/Subnet
* VPN Line VGW / TGW fix to remove unnecessary lines between transit gateways.
* VPN Metadata shows incorrect ip for connected device
* Incorrect Traffic path for AWS map for DC
* Time Range control: validate empty dates and fit long dates without going to new line.
* Guided Demo - Allow Azure Cloud Map Access
* Weathermap Optimization
* Weathermap error for companies with no clouds configured
* Azure Entity Explorer doesnt refresh topology and subscriptions change
* Fix Onboarding Security
* Kentik Map - Azure - Errors
* Fix missing azure subscription ids on custom dimension selection modal

**Core**

* Device Form - Allowing kappa against cloudpaks
* Update PagerDuty and OpsGenie logos
* Move Guided Mode to not being 'freeform' except where necessary
* Supporting horizontal scrolling in tables
* Send trial support emails to their account team
* Capacity Planning Interface Menu updates
* Add trace\_id to query logging
* KDE query post-processing performance improvements
* Create zendesk for spoofers in trials
* Update pg calls in prep for PG 15

***Bug Fixes*:**

* Add safety and logging to quicksort
* Fix vpc name bug
* Handle unexpected websocket disconnect
* Fix live update offset timestamps
* Don't set device\_alias anymore
* DeviceSelector - better handling of duplicate device names
* Query normalization fix
* Live update time shifting
* Front time shifting in live updates
* Fix query interval for when using SNMP queries
* Fixed http prometheus metric uri props
* Fixed sorting of query results
* Fixed Showing OTT Matching Percentages
* Handling archived/incomplete devices in NE device details
* Improved adding saved views to Observation Deck
* Misc Query Engine Work
* Fixed dashboard links for queries with cidr filters
* Fix creating dashboard with Azure guided dimensions
* Fix city, region, and country links
* Small Query Improvements
* Query Builder Fixes
* Fix for old reAggFn
* Fix results in previous period queries
* Fix validation of uniqueness of site market name
* Making raw flow permission exist by default
* Fixed warning-sign and error icons to use color over intent
* Support error display when field is prepopulated but showPristineErrors is not defined
* Fix site market name validation
* Audit Log improvements
* Resolve (cloud-related) populator search returning dupes in filter dropdowns
* Server-side filtering on interfaces page plus "Null" search issue
* Only mark as dirty in defaultValue situations
* Fix: Audit Log Date Range filter fix
* Fix Traffic Engineering for devices with dashes.
* Enable/Disable Streaming Telemetry for device
* Audit Log Fixed MKP, KMI, Totp and CD things
* Horizontal Scrolling Fixes
* Fix Query Engine OOM
* Table - Horizontal Scrolling fixes
* Audit log fixes
* Fix potential chunking issue for large queries
* Accessing option group safely now when showing Dimension tags
* Time Series Bucket Aggregation tweaks
* Prevent rendering issues with interfaces in bar charts
* Preventing 401 errors in shared synthetics
* Fix site country links in Data Explorer
* Remove external id from zendesk sso payload
* Improve query aggregation option handling
* Put chart in comparison mode when using Compare Over option on row
* Safety check to prevent bracketing error reports
* Fix Kube Query Errors
* Short flow agg window fixes
* Adding "overriden" indicators per #10002
* Exclude kmi urls from auditlog + better name

**Alerting**

* Amp & Reflection migration script improvements
* Alerting Summary, Insight & Alerting Table Improvements
* Add time range filters to attacks, insights and mitigations
* Adjust new alerting summary for tenants
* Linking to dashboards in MKP tenant alarm views when they're accessible
* Add alarms sortBy and pageToken to alerting API
* Updates to Amplification & Reflection Policy Migration Script

***Bug Fixes*:**

* Fixed minor silent mode form issues
* Baselinebackfil fixes
* Notification channels not being persisted when threshold is disabled before update
* Fix insights duration on fetch
* Insights - Safety checks for startsWith() computeds
* Fix "default policy to silent mode" bug
* Fix hashed table filter time range selections
* Show overlays with links when no flow devices setup
* Fix DDoS config step 2

---

© 2014-25 Kentik
