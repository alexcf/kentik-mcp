# Release Notes 2024

Source: https://kb.kentik.com/docs/release-notes-2024

---

Kentik employs a continuous deployment methodology for constant extension and refinement of the Kentik v4 portal and the underlying Kentik platform. Release notes for each successive month of Kentik v4 updates are covered in the following topics:

* **December 2024**
* **November 2024**
* **October 2024**
* **September 2024**
* **August 2024**
* **July 2024**
* **June 2024**
* **May 2024**
* **April 2024**
* **March 2024**
* **February 2024**
* **January 2024**

> ***Note:*** For additional insight into what’s new with Kentik, be sure to check the **Product Updates** page of our website.

## December 2024

**Alerting**

* Notifications: interface names instead of IDs and Protocols
* Filter alerts against current/final severity instead of historical severities
* Updated Export Alerting Data dialog letter casing
* Policy Form - Min Traffic configuration greek-it support
* Mitigations - Flowspec match protocol enhancement
* Prevent undefined metricValueToLabel on Up/Down Chart
* Increase max baseline keys to 10k for kmetrics threshold policies
* Add Overview Link to Alerting Page
* NMS Notifications - remove dupe dimension keys
* Notifications policy association fix
* NMS Policies - summary panel in wizard for both edit and create
* Native NMS Alert conditions summary component
* Support Native NMS alerts in Site Health Map, Alerts by Site results
* Clearer toggle actions for binary policy settings
* Native NMS Alert Dimension, Chart & Debug Improvements
* NMS Policy Form - Finalize validations

**Cloud**

* Updated menus and headers for Cloud => Performance Monitor
* Pass apiurl to terraform script to fix 401 auth err
* Fix AWS terraform page reload bug
* Weather Map - OCI Connection To On-Prem (Missing Line)
* Canva transit gateway route table missing
* Cloud metrics explorer button missing queries
* Select/Deselect Child Accounts when Using Org Role
* AWS Modifications to 'Show Path' details

**Core**

* Changed 'Proxy Agent' to 'Kproxy Agent' in the settings menus
* Changing the link to match the page heading and breadcrumbs
* In Search > Recents, if there are no saved views hide the header
* Top Talkers and Settings sub-menus no longer remain open when NavbarM…
* Changes the heading in MKP>Add Tenant>Data Sources>Advanced Filtering…
* Auto deploy test runner with RC updates
* EnableCapabilities: rightWorkingCapabilitiesList should be state
* New criteria for unsupported distro
* Removes (NMS) from the Universal Agent links
* Allow rightFilter operator to pass validation
* Add readnews to device rbac
* Added 'excluding' to apply filter pop-up
* Updated word "Master" to "Primary" in Add Device dialog
* Fix for OTT Capability Edit
* All vs Any
* More KDNS Improvements
* Add to api internal whitelist
* Add "Create public branch" gh workflow
* Fix RC deploy action
* Connectivity Costs - Show Remove Provider button appropriately
* Kagent: two new cmd args in KDNS
* Enable Capabilities fix
* Move API schema into its own subpackage
* Gives option to choose between account vs product question
* Script for importing certain cost groups from CSV
* show not found for APIs instead of throwing error
* Kproxy UI v1.0
* Add device selector to quick-view subnav
* DeviceSelector verbiage update
* Show missing timestamps for cost SNMP export
* Fix onboarding device create
* KProxy Capability feature flag
* Skip BGP tests via comment
* Run traffic cost job daily

**NMS**

* Fix error when closing device edit dialog
* Fix margin around nlq bar
* Don't show device configure button until device is loaded
* Fix presence and label of site name dimension
* Remove all references to table mn\_kmetrics\_lldpneighbor
* NMS - ICMP discovery should not convert existing SNMP devices
* Allow Metrics Explorer if any measurements exist
* NMS policy form verbiage change
* Fix EVM service logging for NMS alerts
* NMS test for EventViewModelService
* Add request info to logs for grpc request errors
* Add edit button/dialog to device details page

**Synthetics**

* Update Synth Grace Period Bin Script
* De-dup synth test targets in service on create/update, adjust validation to fix legacy dups
* Alerting conditions text change

## November 2024

**AI/Insights**

* Journeys: Cloud CC view report button and scrollX
* Journeys: Cloud Check by IP, DNS, or CIDR
* Journeys: NMS queries will still show devices, time range, and filters with no results.
* Core: Kentik Probable Cause in Data Explorer
* Journeys: add rightFilter, cleanup few-shot examples, cleanup tests
* Flag FPA RC
* Add new opt-in company setting for Kentik AI
* AI-Opt in should be Super Admins only
* If a user has no Journeys, submitting a new prompt will create one
* More Kentik AI enablement tweaks
* User Journeys to display `share\_level` more prominently
* Journeys: Add Cloud Connectivity Checker
* Cloud Connectivity Checker with Kentik AI summaries
* Properly show Journeys based on share\_level
* Journeys: Ensure default DE dimension
* Journeys usability improvements

**Alerting**

* NMS Native Alert Debug Charts
* Fixes to NMS Policy Collection Filter & Removal behavior
* Clarify Suppression Start/End Times
* NMS notifications update - minor adjustment
* Clarify Alert Silence Start/End Times
* Native NMS Notifications EventViewModel Fix
* NMS Notifications - remove dupe dimension keys (#25400) Port to master
* Fix: Add missing startTime to suppression serialization
* NMS Policy - Level condition threshold value input improvements
* Alert detail wording + form improvements
* Show policy validation errors in the policy list
* Wizard - query and location state filters support
* Add activation clearance delays for toggle mode
* Fix query builder for alert baseline backfill
* Improved filtering for MKP policy alerts (landlord view)
* Clone policy from template fix
* Support for Custom Native NMS Suppressions
* Policy drawer for NMS policies
* Wizard - on save handling
* Interim policy selector for Native NMS
* Fixes for NMS notifications
* Various NMS Policy fixes and adjustments
* Fix policy clone for NMS policies
* Fixes Auto Ack Edit/Remove Comment dialog typo
* Fix Alerting PDF export
* Only update BGP proxy if current matches expected
* Change the maximum time window and grace period for alerting policies to be 360 minutes
* Wizard scrolling + NMS Conditions Improvements
* New alerting policy url scheme
* Bug fix: NMS policy measurement onchange condition
* Add Sudo + Feature Flag option to Create NMS Beta alert poilcy
* Sort and group by application NMS alerts
* Charts for Native NMS Alerts
* NMS Policy notifications handling
* Fix presets sync

**Cloud**

* GCP firewall rules redesign
* AWS Connectivity Checker improvements
* Traffic queries in cloud map sidebar should use cloud devices
* Azure support with subnets with addressPrefixes
* Azure firewall log export needs storage acct
* Connectivity checker api
* Error in connectivity checker for AWS asymmetric routing scenario

**Core**

* No Longer email support on mkp advanced mode issues
* Add src or dest site by ip filter
* Fix `DataViewHeader` applied data sources alignment
* Add metric to capture redis list size
* Set sdm start date to the start of next hour
* Cache previous period data correctly
* OTT: don't show comparison until query is fully loaded
* Remove internal ASN filter from Kentik Map queries
* OTT Config Tab - OTT tap statuses and a call to action to add more
* Core: Add support for right side filter field in Data Explorer and Query API
* Hide interface classification from users with restricted device permissions
* DE: Device Type filter now works + removed unused functions from $devices
* Turn BGP e2e tests off
* Fix botnet threat phrasing 2
* Label changes for 2 UX Papercuts
* Update `jsconfig.json` `excludes`
* Add safety for scenarios where dv is loading (FPA)
* Fix race conditions when rendering alerting table
* KDNS: part 3 - better errors and handling of errors
* Update execute cost endpoint to allow updating specific cost groups
* Add missing PDB Custom Dimensions to Signup
* Minor cleanup to finalize device RBAC
* Prevent possible render issue when showing warnings about incomplete setup
* Correctly display geo market stack when selecting new one
* Possible fix for error when rendering metrics explorer
* Traffic Costs
* RBAC SSO: Prevent copying of existing role/role set
* Connectivity Costs - Backbone
* Fix role save toast
* Standardizing the phrasing for all Botnet & Threat Analysis headings …
* Fix device create with monitoring template preset
* Add Populate Dynamic Interface Groups Route to internal API
* Fix New User Create with Rolesets
* RBAC SSO UI: Select in radio now matches behavior
* RBAC SSO: Updated defaultValue to match db
* Refactor: remove barrel file from mesh utils
* Remove delay when showing RBAC label selector
* Added an analyze button to Metrics Explorer, allowing easy access to Data Explorer queries
* fix for transferring ownership of multiple subscriptions on user deletion
* Minor improvement to subscription table text
* Updated Profile verbiage from "Visualization Color Settings" to "Chart Colors"
* Patch tving mlb logos
* Optimize chunking in OTT queries
* Make ME->DE time range always consistent
* Make metrics explorer->data explorer queries have closer timestamps
* Fix common ix filter
* Add loading skeleton to global Search so the UX is more obvious
* RBAC SSO UI
* Enable device RBAC feature flag for zayocom
* Kprobe release 2 - mostly UI
* Fix User Update
* make sure asns is wrapped in an array
* Fix breadcrumbs around Discover Peers workflow
* Prevent MKP edit package from rendering infinitely if it doesn't exist
* Conn Costs - Export SNMP in csv with relevant timestamps only
* kagent: add `-net=host` to docker install command
* Fix login error on express middleware
* Removes 'Manage' from settings page for site, site markets and interf…
* Prevent Kentik Map sidebar from infinitely updating component
* Remove Database ACL
* APIv6 refactor metrics with prom metrics
* RBAC SSO Integration Milestone 1
* KDNS initial release
* Prevent cost details graphs from having invalid x-axis labels when changing months
* Fix feature flag environment name detection
* Add AWS1 Onprem env

**NMS**

* Allow selecting labels that are not currently applied to active devices in the device selector
* Allow grouping by `Vendor` and `Model` on NMS Devices
* Try to help someone debug node-interface-group.
* Fix NMS alert charts for BGP session-state metric
* Show all device names and labels in device selector
* Policy table updates for NMS policies
* More specific labels for NMS policy condition values
* Fix default monitoring template breaking device update via API
* NMS Policy activation and clearance delay integration for Alert Manager
* NMS Policies - integrate ack required
* Fix NMS policy editing
* Fix NMS policy cloning
* Fix alerting table dimension column rendering
* Fix serialization of NMS policies with custom targets
* Adjust alert table dimensions rendering
* Allow users to filter credential dropdown
* Show ack state for un-acked alerts
* NMS discovery agent table updates
* Select default monitoring template in forms
* Fix error in percent renderer
* Device monitoring templates
* Policy form should render if list request throws
* Policies table dimension column rendering for NMS policies
* NMS policy device selector
* CSV export of NMS alerts
* Fix admin\_status interface group filter
* Rollout bundle manger service
* Docs: make it clear that some properties on nmsAlarmContext are optional
* Alert table dimension rendering for NMS alerts
* Metric column rendering for NMS alerts
* Alert drawer part 1
* Native NMS Alert Actions: Silence, Suppress, Auto-Acknowledge
* NMS alert detail drawer
* NMS alert entity details
* NMS alert detail page
* Update condition groups when entity type changes

**Synthetics**

* Enable Synth Test Results Redesign
* Synth Test Results Feedback
* Synth Summary Panel Date Display
* Synth Grace Period Initialization/Validation
* Fix Synth Test Alert Suppressions
* Prevent BGP Alarms Fetches of More Than 1 Day
* Fix Duplicate Summary Component in BGP Combo Test Results
* Synth Test Results Default Selected Tab
* Remove Seconds Unit from Synth Date Picker
* Fix Synth Subtest Results Empty State
* Synth Flow Test Results Empty State
* Synth Test Results Tab Rendering
* Remove Synth Subtest Results Incident Log
* Add name label to syn test dialog
* Reorder dropdown list of metrics on Toolbar for mesh results
* Update text on Check RPKI label
* Show 'New Test' on creating new tests and test name if editing it
* Increases contrast of the Syn private agent icons
* Fix Synth Min Lookback Calculation
* Rollback BGP Alerting Field Validations
* Fix Sync Master Into Develop
* Synth BGP Feature

## October 2024

**AI/Insights**

* Fix Gemini text responses. Add Gemini Pro 1.5 002, Llama 3.2, Claude 3.5 v2
* Journeys sidebar improvements
* Fixes Insights flow count, rounded to two decimals
* DE NLQ: fix geo region filters

**Alerting**

* Feature flagged NMS policy collection, with all policy collections wrapper
* Remove obsolete pendo v3 alert policy config
* Native NMS policy notification associations & serialization
* Omit Native NMS policies from MKP, mark todos
* Fix unit tests
* Support for NMS alerts in node alerting routes
* NMS Policy form and small Wizard improvements
* Node layer NMS feature flag based permission checks
* Fix presets policy sync
* Fix preset policy sync
* Fix the grace period calculations for MKP policies
* NMS Policy Form continuation
* NMS Policy CRUD operations, notification handling and enable/disable
* Add activation conditions to NMS policy form
* Test & Validator clean-up
* Various NMS policy model mapping fixes
* Fix to properly assign/remove notification channels on MKP subpolicies
* Change "Core" policy label to "Traffic"
* NMS Policy prep work - listing nms policies and some constants
* Add an option to disable TLS verification for mitigation platforms
* Ignore legacy notification channel conjunctions
* Wizard small clean-up and adjustments
* Basic re-usable form wizard

**Cloud**

* Cloud export wiz Refactored
* GCP route table missing some targets
* Make custom dimensions admin again b/c of security vuln
* Merge master fix conflicts
* AWS Scrape Metadata Improvement

**Core**

* Expand device RBAC feature flag
* Fix up feature flags
* Add iij1 onprem env
* Add iij1 portal hcl
* Fix error when deleting user
* Fix end\_date attribute when saving cost group from previous cycle
* Fix cost groups not being found for any company during cost job
* Add device RBAC to APIv6 routes
* Enhancements for OTT landing page top categories and providers
* New ASN logos
* Peacock logo fix
* Migrate "/api/restricted" from node-portal to node-api-internal service
* Resolved minor gap in Device RBAC
* Integration Test Runner
* Device RBAC: minor fix for findSite
* Users without sufficient RBAC rights are not allowed to subscribe to reports
* Device RBAC feature flag
* Fix flow tag search
* Added device RBAC to listInterfaceGroup
* HackerOne: prevent dashboardItems being altered via the wrong dashboard ID
* Prevent underscores in subdomain for MKP
* Form state updates
* Add error handling when loading malformed dashboards
* Fix db mock to return null when appropriate
* Update nav menu item Discover Peers
* Add company\_id to http metrics with 5xx code
* Fixing the scaling of the navbar on narrower windows
* Small UX tweaks around Library and Saved View detail page
* Fix error in dataview header when query isn't available
* Fix ott service capacity interface chart lookback
* Improve universal agent status tags
* Update new LGU1 Sending Ips
* Fix sorting in Universal Agent Page
* Prevent Members from Creating Site Market
* Fixed initialize order for scheduled reports

**NMS**

* Display percentage with enum metrics in average rollups
* Prevent error when device availability widget data is empty
* NMS alerting: list NMS alerts in alert table
* Fix NMS policy edit navigation
* Fix NMS policy serialization of state in conditions
* Expose NMS alarm context from AlertModel
* Update enterprise/model data for NMS discovery
* Use non-prefixed dimensions in NMS policy target
* Update enterprise/model data for nms discovery
* NMS alerting form config setup
* Set up entity type data for use in NMS Native policy form
* Initialize model correctly on nms policy page
* Add util to remove ui-specific properties during serialization
* Wizard custom field types
* Add new paths to codeowners
* NMS Policy active alert count and wizard warnings
* Handle null values when parsing query results
* Fix snmp status of devices that have disabled nms
* Route scaffolding for NMS policies
* Fix alerting routes
* NMS policy model

**Synthetics**

* Adding '+' to certificate expiry threshold maximum label
* Syngest-Driven Preview Test Results
* Final Migration of v1/v2 Synth Views
* Syngest-Driven Preview Test Results Final Migration
* Delete tests from inactive companies
* Fix Synth Subtest Metric Charts Offset
* Final Fix Agent Renderer
* Fix Synth Run Preview Again Functionality
* Remove Per-Agent Alerting Logic
* Fix State of the Internet NaN Results
* Summary Panel Component
* Synth Subtest Avg/Min/Max Aggregates
* Handle Unknown/Removed Agents in High Density Grids
* Support Direct Alarm Navigation in Alarms
* Fix Unknown Agent Logic in AgentRenderer

## September 2024

**AI/Insights**

* Fix regression in Journeys API setup for model ID
* Journeys: Add unique IPs for DE
* Journeys: Render math/LaTeX in LLM responses
* Journeys: Use optional company or env model ID
* Journeys: Convert VPC name to ID
* Journeys: Add safety for location.state
* Journeys: Fix location.state bug if logging in with redirect
* Journeys: Compare models
* Journeys: Fix NMS check

**Alerting**

* Fix MKP Alerting/Mitigation links for tenants
* Add HoC annotation for subscriptions
* Prevent errors/discrepancies on unassigning/duplicating mitigations
* Alerting export to CSV fixes
* Add NMS native policy API services & handlers
* EVM Notifications clean-up
* Alerting Overview
* Notifications fix - allow port 0 too
* Order severity groups from most to least severe on alerting table
* Fix dupe Alerting breadcrumbs on Not Found state
* Permit custom CIDR subpolicies on tenants with CIDR filters
* Spelling & syntax fixes for "acknowledgement" related words and labels
* Group by Ack State, Ack & kebab menu improvements
* Test configuration notification selector fix
* Mitigation ID filter parameter should get cleared properly when cleared from table search bar
* Alert Severity History and Chart Improvements
* DDoS attack activity chart - fix cleared alert detection
* Use alarmIDs for mitigation notifications
* Hide alert detail debug button for MKP tenants
* Policy update rollback fixes, improve error messaging
* Include BaseDomain on subpolicy notifications
* DDoS landing page 24h attack chart fix
* Syslog notification improvements
* Subscriptions improvements and fixes
* grpc says that in the NMS subobject, ports are ints and timeouts are durations
* Mitigations partial target search
* Fix policy enable for unexpired company
* Add source and dest ports to manual mitigations
* Adjust mitigation rate limit action label to reflect actual metric used
* MKP: More accurate CIDR compatibility check for subpolicies
* Show correct activation settings for subpolicy mitigations
* Fix subdomain check for notifications, store tenant ID on subpolicy
* Fix mitigation data CSV export for alerting
* Fix subscription report type empty options for alerting
* Point legacy alert insight links at insight details page
* Disable custom CIDR subpolicies & explain why policies are disabled in MKP
* Remove subminute interval case from alert query chart builder
* Show alerting charts by policy metric order
* Improvements to policy form UI and validation

**Cloud**

* GCP config status missing some api calls
* Add null checks to cloud config status filter
* Connectivity checker | using Security Group that is not linked to ENI interface
* No results for Nat Gateways IP address in entity explorer
* AWS add account aliases into metadata
* AWS Core Network Traffic Data
* Fix missing tags from AWS RAM TGWs, missing labels for obs deck TGW widget
* Restrict AWS metadata fetch to only ACTIVE accounts
* Enforce pagination for some aws cli calls
* Azure Topology Scraper Optimization
* AWS Direct Connect Metrics
* No metadata for aws-cn nested secondary accounts
* Prevent caching of empty kube metadata
* OCI Add flowlog sample rate into metadata
* Generate AWS china access token keys from secrets
* Kube Map Performance Improvements
* Empty Azure and GCP clusters

**Core**

* Check if MapboxGL JS is instantiated before checking style is loaded
* Fix raw flow
* Fix for failure to render cloud account status sidebar
* Menu: 'Recently Viewed' correctly shows last 3 recently viewed items
* Parse query result arrays as json
* HackerOne: Shared Links able to update across companies
* Fix for failure to render metrics explorer options menu
* jsdoc improvements for /api
* Fix Filtering on UDR Fields with commas (no-split)
* Make sure device rebuild is done after as group dimensions are created
* Consistent costs in Cost API
* Update redis queue params to test connection behavior
* ASN detail page should render for AS group with null matchers
* Core: Fix ktrac queries post device RBAC
* Add request ID info shortcut for Firefox users
* Fix host in tagApi test
* Fix alerting application name
* Allow site ID to be passed to findDeviceHandler
* Throw clear error if spoof restrictions and user is not in yaml
* Added 7.48 to KproxyAgentModel.js
* lgu1: add hcl file
* Fix test failure in deleteDevice
* Fix Live Update for Time Series
* Remove Demo Guides
* Re-changing device selection where clause syntax
* kmetrics device RBAC
* Basic Device RBAC for Journeys
* Added `force\_huawei\_16bit\_fix` to deviceInternalApiSerializer
* Small JSDoc improvement
* queryBuilder dynamically applies RBAC for Devices
* Allow View as Tenant for MKP with 2FA
* Remove device RBAC from MKP
* Ensure all attachments are included with subscriptions
* Fix inability to edit dashboard items set to ignore guided mode
* Add LGU1 Onprem env
* Fix invite users, informing account team, audit log exclude and more
* More metrics and status for job scheduler
* Mass testing for node layer moves coverage up 0.75%
* Polish and testing on SVG Grid element
* Minor JSDoc improvement
* Fix most existing JSDoc declarations
* Mocking gRPC communication in test
* Remove deprecated $query functions
* Add job id to refresh jobs
* `pt-intent-primary` now passes AA color standards
* Fix Manual Interface Classification
* Refactor to resolve circular dependency issue
* Add exclusion for sdm-leader
* Correctly use transactions for metrics devices
* Make FormErrorsButton Observable Again
* Device label user\_id might be null in db, but proto requires string
* Associate new devices with the `metrics\_device` for SDM AP 5 and AP 16
* Fix up Application Names

**NMS**

* Update enterprise/model data for NMS discovery
* Process NMS query results faster
* Throw if attempting to delete a kmetrics device
* Add columns to NMS device connections table
* Hide auto added interface filters with capacity
* NMS ME: Fix filter count
* NMS Metric Filters now use available Enum values when appropriate
* Fix interface bitrate not showing up when capacity is 0
* Show AS groups in device BGP neighbors table
* NMS up/down policy save fix
* Fix Site column and Group By on NMS Devices page
* Alert Debug: handle missing / undefined debug values
* Small optimizations for device availability query
* Update enterprise/model data for NMS discovery

**Synthetics**

* Fix import of getASNDescriptionByID function in preprocessPathGraph
* Do not enact `alertsDurationGreaterThanPeriod` for BGP Monitor Tests
* Return all consumers of a credential, regardless of status, to agent
* Convert to per-agent alerting
* Allow saving/creating when relevant sections have errors
* Synth Test Results Redesign Updates
* Trace packet loss label update
* Fix Synth Test Redesign Dataview Render and Incident Log Links
* Allow theme switching while on map
* Add Validator.js rules to accommodate slider input on TestConfig Range sliders
* Allow test authors to provide gRPC manual mocks as functions
* Synth Test Results Redesign - Time Picker
* Test Results Redesign - Parent and Subtest
* Proto `useLocalIp` -> `use\_local\_ip` mapping for ip-address tests
* Update Universal Agents > Edit Agent KB link
* Add City to Synth Mesh Cloud Agent Labels
* Modifications to support synth agent clustering
* Agent Cluster CRUD
* Fix Incident Log Warning Filter
* Fix MKP Synthetics 404 calls

## August 2024

**AI/Insights**

* Journeys: Re-org source
* Move Insight voting to header for more visibility
* Journeys: Update DDoS to AlertManager
* Fix dimension value getter for insights
* Journeys: Improve filter operator handling
* Journeys: per company model selection
* Journeys: model comparison tool

**Alerting**

* Fixed error on manual mitigation when no method is selected
* Add permission based checks around mitigations / notifications / policies
* Fix `$query.addFilters` errors on DDoS alert details tabs
* When filtering by a date range, current time should refer to actual current time, not page load time
* Fix MKP Subpolicy and Mitigations Save Errors
* Add query-to-policy type mapping for notifications
* Fix: Prevent subpolicy rule alias collisions
* Hide mitigation selection for MKP package policies & exclude A10 methods for MKP tenant policies
* Warn but allow old end dates when editing Auto-Ack expiration
* Show sites as a clickable link in Alerting
* Show dimension labels when exporting to alerts to csv
* Fix notifications missing device name and type
* Checking existing silences should look at expiration too
* Safeguard bidirectional check for existence of condition on ratio edit dialog
* Fix alerting dimension names and group by
* Fix: NMS alert detail dimension wrapping
* Hotfix: Use subdomain in MKP tenant alert notification links, if configured
* Hotfix: Evaluation period subminute message
* Hotfix: Prevent undefined property on storage error
* Hotfix: Don't filter suppressions by current user
* Hotfix: Silent mode expiration date format

**Cloud**

* Fix missing subnets
* Update AWS topology cron job frequency updates
* Handle GCP project permission err better, add null check
* AWS Glacier buckets error message for cloud export validation
* OCI Cloud Summary Widget
* OCI Config Status API displays a green check when some APIs fail
* AWS Cloud WAN: Show path to destinations route should be Propagated
* Incorrect line style for AWS core network lines
* GCP metadata fail
* Kentik Map logo bug
* AWS Core network null check when generating topology
* Add Core Network metrics to topology map
* Error when showing details for region summary
* Fix bugs in Azure subscription enrichment for China
* Treats UDR fields starting with 'str4' as a kentik dimension for filters
* AWS Map - Display AWS Core Network.
* AWS core network existence check when computing routes

**Core**

* Fix error when changing flow aggregation settings when using this / last month lookback
* NMS interfaces RBAC
* Re-enabled eslint import rules with webpack config stub
* Pure refactor: split out Metrics Result columns
* Extract some reusable functionality from $query
* Better handling of missing dictionary entry for cloud provider
* Fix GeneratorView erroring out in corner case
* Fix Visualizations error when attempting to redraw an empty dataview
* BGP devices: Make the peering IP number uniqueness across companies
* Hide Dimension component if it is missing required props
* Fix api internal error handle and logging
* [automated testing] Some low hanging fruit
* Fix: pass user\_id to findPackages
* Rollback device permissions yml entry
* Fix interface group dynamic interfaces preview
* Additional metrics and logging around schedule jobs
* Extract all queries into separate functions
* Fix trending view node-job
* Created generic Resizer element
* Code coverage
* deviceModel enforces fetches for fps
* Fix for not validating roles in SSO admin routes
* "Controlled by Dashboard" fields now showing correct initial values
* Fix circuit route tables node-job
* Assign correct role when user created using API
* UI Interface RBAC
* NON-alpha release of MKP public api
* Update [deploy.md](http://deploy.md/)
* remove react dependencies from $sites
* Public API version of credential vault
* return no results for null kube queries
* Adding a key null check in v4/profile/viz
* Safer computation of metric prefixes
* Fix restore of archived devices
* Fix failing test on develop
* Fix env typo for HRD1
* Fix Consul watch initialization error handling
* Fix MKP 404s
* For a Grafana link in Agents
* Remove incorrect RBAC check from updateBgpProxy
* Prevent error on month change when cost drawer is open
* Update query engine error logging
* Connectivity Costs CSV exports show data from the correct month
* Managing Single Role was titles as Roles
* Basic Recon device RBAC
* Update RBAC option in Company menu
* Interface classification exec base on consul leader status.
* Remove schedule report cron job
* Granular MKP Features
* Remove incorrect RBAC check from updateBgpProxy again
* Enable populate trending view job
* PR ephemeral deploy prod env
* Migrate ui/deviceRoutes to RBAC
* Fix auto deploy for develop branch
* Delete unused files which reference nonexistent code
* Export utils was missing a param
* AWS Core network null check when generating topology
* Fix spoof reason dialog not being shown when not logged in
* All everyone to set subscription run time
* Add safety checks for String.prototype.toLowerCase in maps files
* Fix device params and update jsdocs
* Prevent filtering on undefined object in CoreNetworkSummary

**NMS**

* Fix disappearing connections on return to device overview
* Color BGP idle state as red
* MPS query should use avg aggregation
* Add.csv and.pdf export
* Update enterprise/model data for nms discovery
* Fix health monitoring widget
* Fix missing/disappearing alerts in up down history
* Sort results by chart metric
* Don't show link if remote device/interface details are unavailable
* Fetch interface IPs correctly
* Make sure hashes are persisted in saved views and dashboards
* Change NMS views and APIs to use mn\_device.site\_id instead of mn\_kmetrics\_device.site\_id
* Update display value formatting in metrics explorer chart tooltip
* NMS in MKP
* Make `status=2` the default device-down alert instead of `available=0`
* Reduce sublimit to speed up queries
* Don't show internal links/pages to demo users
* Update enterprise/model data for NMS discovery
* Ensure label for attributes matches input id

**Synthetics**

* Do not show ERROR when trying to configure a non-private Agent via Agent Management
* Fix MeshPopover render errors when missing props
* Fix "Agent Not Found" in Test Results Table Target Column
* Trx Screenshot error handling
* Fix Alert Manager Alarms Fetch to Support Context
* Update Synth Agent Maintenance Mode Filter
* Flex Health Timeout Error
* Remove Dead Test Results Code Paths
* Fix Historical Synth Tasks
* Fix Cross-Region BGP Routes
* Fix Incident Log Start Date
* Adds Data Retention to Syn panel on Licenses page
* Clean Up Delete Agent Alerting Rules
* Remove Cross-Region BGP Policy Code
* Updating query relay app name for ksynth services
* Synth Test Suppressions Clean Up
* Library filter dropdown was too narrow.
* Disable 'Start Monitor' if user does not have synthetic test create
* Re-add fix where target filter can be either resolver or target
* Notification synth test used by association removal fix
* Migrate Synth Store - Final Removal
* Fix Synth Test API Notification Channels
* Fix Notification Channels Usage Stats
* Merge Notification Channel Usage Fix
* Reset Rollup Level When Using Per-Agent Alerting
* Remove some of the alerting 1.0 synth code

## July 2024

**AI/Insights**

* Journeys: Update “Site by IP”, OTT filters, and map viz
* Add NMS chart API and hide CoverPage for DE chart export
* Journey input now focuses more reliably, and has a button to submit prompts
* Journeys API: Fix NMS query output
* Hide "New Journeys" that are not created by the active user
* Journeys: AS groups, Sankey, KB reference note, Peering few-shot
* Journeys API: new version and add DDoS
* Add KDE miner and KDE debug mode, update CLI modes

**Alerting**

* Hotfix - NMS Notification Device Lookups
* Hotfix - Adjust policy edit button height
* Hotfix - Subscriptions need applications as an array
* Hotfix - Alerting dashboard selector, policy deserialize was stripping dashboard
* Fix: Handle “perc\_retransmits\_in” & “retransmits\_in” for alert metric to DE agg conversion
* Don't pad the alert chart end time if the alert is still active
* Hotfix - DDoS Tabs height on DDoS alert detail page
* Alerting - Integrate Core & Protect Policies with Alerting 2.0
* Alerting Milestone 5 Notifications - Alert Manager integration
* Suppressions UI Improvements
* Limit what platforms/methods tenants can see
* Custom Webhook Notifications OAuth integration
* Add a flag to NMS policies for when device settings change

**Cloud**

* Prevent updating GCP cloud exports when values don't change
* Azure Cloud Config sidebar does not count partial requests
* Increase max payload size to 256MB
* Azure Incorrect Sidebar Details label for application gateways
* Verification for Azure Subscription Enrichment Scope is not working
* Azure missing information about the application gateway
* Azure Explorer Widget additional data for application gateways and load balancers
* Azure Cloud check Invalid Date 500 error
* Azure Map shows duplicate gateways
* Azure BGP peer status 500 error
* Alerting Migration Settings updates for rule and policy services
* OCI Exporter - Allow configuring multiple compartment ids in a single exporter

**Core**

* SNMP community audit typo fix
* Override for bp4 dialog header size change
* Fix lastdayOfMonth schedule report run
* Audit Log censorship hotfix
* Updated puppeteer function calls for csv
* Exposing GCP tradingviews secret in Nomad
* Prevent agent status cache from desync
* Github action to auto deploy PR build (part 2)
* CDN Analytics: Added G-core & Globo to list of valid CDNs
* Upgrade to Blueprint 4
* Bug Fix: add override for bp4 dialog theming issue
* Fix DialogFooter margin
* Update run schedule report timezone to UTC
* Update User Filter Script
* Fix API auth error
* Tidy up kentik/ui-shared package
* No longer scroll to element which has ceased to exist
* auditLogger no longer sensitive to capitalization
* Bin script to decode creds
* Added Ultimate Exit VRF Name and RD Dimensions
* Split up rbacMiddleware to allow for dynamic RBAC calls
* RBAC device script update
* Add Restricted GET Device API
* Report Subscriptions 2.0
* Fix infinite loop when checking missing session key
* Require reasoning and restrictions for spoofing and log them to audit logger
* Release/fix auditlog merge
* Add Subscriber ID Fields
* Remove 8 hour offset for report subscriptions
* Fix page error on settings page from warning callout
* Hotfix: more coverage for auditLog exclusions
* Unified apiV6 auditLogger safeties with NodeJS app
* Add Restricted GET Device API
* Make mobx stop complaining about ErrorBoundary
* CDN Analytics: Config link rendered only if admin
* Fix error toasts when visiting mkp
* Use recent views for subtenants
* Refactored dashboard query to prevent private dashboard access
* Docs updates
* Improve CreateDevice API error messages
* Fix Assigning Roles in User Page
* Audit logging again targeting v4.93
* Fix error toasts when visiting MKP
* Raise protocolTimeout value
* Basic structure of RBAC for devices
* Remove SNMP data from audit logs
* Trigger a form onChange when calling setValue on ArrayFieldState
* kagent docker instructions: add type=volume to mount
* User Settings: fixed typo in updatePermissions
* Standardize error-handling and improve user messaging
* MKP: Tenant library unlabeled views fix
* Add prometheus metrics proto file
* Updated cdn migrate script
* Update node 20 master image tag
* Fix device update API validation
* Made $store.initializeApp() fail more gracefully
* Kill device summaries API
* Copy users.yaml file from operations
* Small design team tweaks - June
* Add yarn resolutions config for node-gyp to avoid python version issues
* Use km\_measurement\_id ILO km\_measurement\_name in KDE queries

**NMS**

* Improve Flow(Core)+NMS device license usage
* Pass saved query id to refresh for metrics dashboard items
* Fix NMS device page 500 due to null IS-IS adjacencies
* Update enterprise/model data for NMS discovery
* Remove sudo check from Connections widget on NMS device overview
* Add Manage Discoveries button to NMS > Devices
* Add ip sorting and disable metric sorting when table is limited
* Fix page error on device interfaces table
* Base Use of NMS Up/Down Policy Form on ActivationMode
* Fix device status discrepancies between Availability widget and Device Details page

**Synthetics**

* DNS results for target filter
* Revert of target filter fix
* Removing space to fix merge conflict
* Migrate Synth Store - Setup and Network Explorer Final Cleanup
* Migrate Synth Store - Shared Test Results
* Migrate Synth Store - Performance Monitor Services
* Unit Tests for Aggregation Logic
* Remove "Show Health" from Test Results View
* Migrate Synthetics Store in Licensing Views
* Migrate Synth Store in Notification Settings
* Migrate Synth Store - Protect Views
* Remove Synth v1 Test Control Center
* Move Synth Test Control Center Out of v2 Namespace
* Remove Synthetics Insights References
* Migrate Synth Store - CDN
* Migrate Synth Store - Setup
* Migrate Synth Store - Observation Decks
* Migrate Synth Store - Test Config Wizard
* Fix Alarm Timeline Color
* Focus Time Range of Test Results
* Fixing SOTI agents not rendering
* Fix Performance Monitor Interconnects Create Test
* Fix for the getWorstHealth is not a function error
* Fix Agent Mesh in State of the Internet Clouds
* Alarm Timeline Redesign
* Colored Raw Test Results Data
* Using other\_healths to determine the status code health for DNS, HTTP, PageLoad tests
* Synth GetTest API Bug Fix
* Dedupe agent selector list

## June 2024

**AI/Insights**

* Journeys: AS groups, Sankey, KB reference note, Peering few-shot
* Hide "New Journeys" that are not created by the active user
* Journeys API: new version and add DDoS
* Add KDE miner and KDE debug mode, update CLI modes
* Journeys: Include optional company prompt
* Journeys: Update RAG embeddings
* Journeys/DE: display filter value labels correctly in tooltip
* Journeys: Add custom dimensions to system prompt
* Journeys: PeeringDB filters, bi-directional, gpt-4o (omni)
* Journeys: respect user Local/UTC setting
* Journeys: Add language prompt to Ask KB
* Journeys: Capture query\_engine errors in NLQ table
* Journeys input autoGrows to accommodate long prompts
* Journeys: enable API in tester
* Journeys: DE Device Label Filter

**Alerting**

* Limit what platforms/methods tenants can see
* Custom Webhook Notifications OAuth integration
* Add a flag to NMS policies for when device settings change
* Fix Legacy Suppressions
* Include Custom policies as options in MKP tenant config
* Alert Auto Acknowledgments
* Show NMS Suppressions
* Silence Alert Notifications
* Rename "Silence Alerts" to "Silence Notifications"
* Fix X axis labels on the threshold chart
* Fix notification usages to correctly report synthetic tests
* Change “Custom” to “Core” Policy Application Type

**Cloud**

* Azure Cloud check Invalid Date 500 error
* Azure Map shows duplicate gateways
* Azure BGP peer status 500 error
* Alerting Migration Settings updates for rule and policy services
* OCI Exporter - Allow configuring multiple compartment ids in a single exporter
* Azure Transfluo Region field for azure exporters
* Azure cloud check api error
* Azure Topology for China
* AWS List Org Account Error fix
* OCI Topology error handler for promise.map
* Azure Connectivity Checker "internal error"
* Cloud map link to k8s broken
* OCI Metrics Collection
* Disabled interface streaming api call to generate aws summary
* NextToken handling for AWS Org listAccount call
* Allow specifying S3 sub-folder for flowlogs in AWS export config

**Core**

* Basic structure of RBAC for devices
* Remove SNMP data from audit logs
* Trigger a form onChange when calling setValue on ArrayFieldState
* kagent docker instructions: add type=volume to mount
* User Settings: fixed typo in updatePermissions
* Standardize error-handling and improve user messaging
* MKP: Tenant library unlabeled views fix
* Add prometheus metrics proto file
* Updated cdn migrate script
* Update node 20 master image tag
* Fix device update API validation
* Made $store.initializeApp() fail more gracefully
* Kill device summaries API
* Copy users.yaml file from operations
* Small design team tweaks - June
* Add yarn resolutions config for node-gyp to avoid python version issues
* Use km\_measurement\_id ILO km\_measurement\_name in KDE queries
* Fix ip address quick view for ipv6 more tab
* Fix trending views tag formatting
* Fix empty RBAC roles on user create
* Add Raw Flow overlay for long queries
* RBAC filter needs a scrollbar
* Ensure simpleFilterUtils.js functions are calling Array.prototype.map on an array
* DE: Fixed custom time range not applied after zoom
* Site Settings: Added large subnet validation
* MKP: Fixed tenant showing no results for label + site combination
* Added kproxy 7.47 commit
* Add an empty users.yaml file for incoming GitHub Actions
* Fixed device flow indicators
* Hotfix excessively strict user API validation
* Fix protect flow check
* Refactor: moved internal/restricted device handlers into subdirectories
* Frontend changes to support two separate definitions of device flowcheck
* Default CDN Providers
* Menu error fix
* Small fixes for CDN dashboard link
* Fix universal agent policy defaults
* Fix bug in TextArea autoGrow when border is non-integer
* Fix API request to topxchart timeout for Sankey viz
* Fix kt teams and add alerting paths
* Minor bugfix for getDeviceStatus API
* Reduce min pool size to 0.
* MKP: Added device\_type filters
* Split out flow checks for pseudo-devices, backend only
* Reduce friction for trial signups
* Optimize connectivity costs history API
* Fix error report team assignment
* Fix device table export
* Added commit hash for kproxy 7.46
* DE: Fixed column capitalization mismatch in SQL query
* Fix Live Update with Reagg
* Fix script name
* Query hotfix
* More generic fix for excluding pseudo devices
* Ensure that Dashboard Form is always created and configured
* Update NMS specific query application names
* DE: Fix extracted DNS information not rendering when using "Total" metrics
* Rewrite OpenTelemetry UI instrumentation
* Better usage of MapFallback
* Batch processing for expensive FPS requests on /settings/devices page
* Less aggressive refetch of Library data
* Interface classification page cleanup
* Debounce filter requests on large collections of unclassified interfaces
* Kentik Map: Don't show private ASN
* CDN Config: Added new collection to override fetch
* Fix kentik map query error
* handling for devices with no flow and speedy clickers
* Enable API Tester Flow Tag API
* Nomad auto rollback on failed deploy

**NMS**

* Remove sudo check from Connections widget on NMS device overview
* Add Manage Discoveries button to NMS > Devices
* Add ip sorting and disable metric sorting when table is limited
* Fix page error on device interfaces table
* Base Use of NMS Up/Down Policy Form on ActivationMode
* Fix device status discrepancies between Availability widget and Device Details page
* Fix incorrect 0s in cpu/mem stats
* Update enterprise/model data for NMS discovery
* Fix missing ICMP host in ranger config
* Update enterprise/model data for NMS discovery
* Disable refetch during queries because it adds too much time
* Fix KentikAgentLink target path
* Return errors from snmp walk as normal responses, not 500s
* NMS Device Availability widget now performs proper lookups on device names and statuses
* When assigning a new agent to a device, ignore error if old agent doesn't exist
* Update enterprise/model data for nms discovery
* Prevent refetch when there's only 1 result

**Synthetics**

* Fix for the getWorstHealth is not a function error
* Fix Agent Mesh in State of the Internet Clouds
* Alarm Timeline Redesign
* Colored Raw Test Results Data
* Using other\_healths to determine the status code health for DNS, HTTP, PageLoad tests
* Synth GetTest API Bug Fix
* Dedupe agent selector list
* Modified fetchAgents service to omit/rollup agent cluster workers
* Migrate synth to use the same Alerting Conditions form as Add Policy page
* Itemize Synth Credit Calculations
* Fixing issue with dns\_status\_health not showing up
* Call setTasksForTests when an agent IP address changes happens
* Fixing the syn internal share link query
* Display Timestamps Up to the Second in Health Timeline
* Fix Incorrect Health Settings Bin Script
* Update Synth Migration Scripts
* Fix Synth Dashboard Form Truncated Fields
* Paginate results when there are >10 tasks or target ips

## May 2024

**AI/Insights**

* Expand LLM query to be valid with next/v5 API
* Add DE Metrics (SNMP) to Journeys
* Remove IBM Cloud from Journeys
* API V6 NLQ Service for Journeys

**Alerting**

* Improved alert dashboard resolver
* Enhanced Ack Workflow - Manual Ack/Unack
* Fix Alerting Site Health Map counts
* Remove check for synth from wrapper service
* Make sure comment box doesn't grow tall enough to reach the moon
* Fix Dst/Src Route Len alerting dashboard dimensions
* Hotfix: Alerting Table State Filters & Summary Chart
* Allow user to filter by dimension value labels, sort by primary metric value
* Fix dashboard device id filter passed from alerting
* Add generic comments module

**Cloud**

* Remove IBM references, incl top talkers and licenses
* Mismatch between path drawn in maps and path in conn checker
* [k8s] main map health indicator
* Add jest snapshot for cloudExportDetails
* OCI traffic charts for DRG Gateways
* Add metric collection option to GCP cloud export
* Connectivity checker cloud WAN infinite loop
* Missing AWS Tag for Network interface
* Weathermap - No Connection To On-Prem
* Refresh Topology populators autogeneration fix
* Remove IBM from cloud exports
* Kentik Kube default device name

**Core**

* Fix erroring on tables with group rows
* Prevent error on interface selector dialog
* Upgrade mobx-react to v6.3.1
* Add new src/dst BGP Large Community dimensions
* [Chore] Upgrade Open Telemetry packages
* FlowTag APIv6
* getDeviceStatus optimization
* Nodejs 20 version upgrade
* Upgrade packages to play nicer with node 20
* Fix error when rendering aws icon
* Added commit hash for kproxy 7.45
* Update [prod-setup.md](http://prod-setup.md/): set password after “add\_user”
* Fix build
* Global Search bug fixes and improvements
* Log spoof url as part of structure field
* Fix Aggressive FlowCheck
* Fix user create
* Credential Vault - Support editing of templated credentials
* Improved webpack dev server config for significantly faster hot reloads
* Fix webpack dev server proxy cache
* Proper fix for webpack this time
* Fix merge conflict resolution from before
* Minor include bug
* Add new src/dst BGP ORIGIN dimensions
* Cache device warnings fetch
* Fixing label names to match grafana's configs
* Fix up logging key explosions in kibana (take 2)
* New Universal Search
* Final Agent Management Polishes
* Optimized setAllUnclassifiedInterfaces
* Renamed "Kentik Agent" to "Universal Agent"
* Fix Button component not showing active status
* Flatten log object to compile with ELK index rule
* Hotfix: get rid of NaN for devices with no interfaces
* Global Search - Feedback while on next
* Fix credential filter
* Fix Misc MKP Issues
* Add new src/dst BGP Extended Community dimensions
* Fix error report logs for ELK
* Update Browser Error Logging
* Populate email addr in ZD (Hopefully)
* Fix calculations of interface count classification percentage
* Eslint upgrade in advance of airbnb-config
* Disable filter based rbac for create credential
* Update “Checkbox” indeterminate state for better clarity
* Lazy-load password set form as it contains zxcbn (very heavy dependency)
* Switch to user swagger URL to load api spec
* Fix IBM UDR removal from breaking DE
* Move SDM leader device create to restricted
* Make all mapbox-gl components async
* Add Delete Universal Agent functionality
* Interface upsert API: persist multiple RTs for VRFs
* Do not check ACLs for shared links
* Fixes Library's Favorites column from disappearing
* Hotfix: improved serve method for JS
* Fix null pointer error in “$silentMode”
* Log Browser Errors to Kibana
* Enable babel caching
* Show (company-wide) last viewed date for Library items
* Create Kentik Agent alerting policy on first deploy
* Check for KMR role IDs in user roles in migrate script
* Add new kt\_team label to browserErrorReports #deploy
* Revert webpack changes which shrunk chunk size
* Revert "Enable babel caching (#22500)"
* Build every branch on Jenkins
* Trending Views
* Exclude pseudo-devices from flow heartbeat
* Move AWS china secrets to saas/our1 only

**NMS**

* In MetricsPlanCard, display count of devices actually using plan
* NMS basic topo
* Add labels to nms credential dialog
* Fix page error when status reason is not given
* Add button to remove ISIS adjacencies
* Show, filter by, and group by device labels
* Devices use “status” instead of “available” for status tags
* Export device and interface tables on server
* Make sure table only query widgets have sparklines
* Add window to nlq cpu query
* Fix issues when changing chart types
* NMS Device Details - tweak logo size, make up/down chart larger by default, wrap better
* Assign a universal agent a Closest Network Device
* “MetricsUpDownChart” properly shows Unknown status
* Add MTU column to interfaces tables
* Fix title on dashboard panel dialog
* Lookup devices from “deviceSummaries” for NMS walk/inspect
* Improve kmetrics notification translation
* Update enterprise/model data for nms discovery

**Synthetics**

* Fix IBM Cloud Synth Agent Details
* Synth Health Timeline
* Fixing Credit Calc Widget
* Update import path
* Refresh Dashboard Widgets
* Flex Health: UI changes
* Migrate Agent Management to the $syn Store
* Removing sudo checks from web/bgp combo tests
* Node/UI for agent settings
* Map Dot and Hover fixes
* Updating synth health description
* Fix Synth Health Timeline
* Fix Test Toggle Activating Policies
* Move agent session validation from redis to syn-back
* Update import path
* Fix Inconsistent Test Alert Rule Rolling Window
* Fix Test Control Center Health
* Fix Test Results Table Render
* Refresh library widgets when query time updates
* Fix Synthetics Test Dashboard Health Timeline
* Type Ahead Search Filter for Availability Widget
* Fix Alert Rule Notification Channel Logic
* Fix Synth API Update Test Frequency
* Fix Notification Channels Post-Migration
* Update synth agent alert tests
* Synth Alerting Post-Migration Fixes
* Fix Migrated Test Policies
* Fix Alerting 2.0 Migration Alarms
* Migrate Synthetics to Alerting 2.0 - Plan A
* Fix mesh drill-down trace having incorrect target
* Create agent alerting public API
* Fix usage query joins, added map\_status criteria
* Fix Path View with Non-Existent Agents
* Fix Agent Downtime Alert Notifications
* Fix Synth Agent Invalid Downtime Minutes

## April 2024

**Alerting**

* Remove invalid alerting column mapping for query builder
* Allow conflicting uniques for Policies
* Fix open in dashboard button saved filters
* Fix alerting constants importing
* Pad Endtime by 30 minutes for alerting insight charts, when applicable
* Fix issue preventing edit of mitigation method
* Support NMS Alert & Custom Suppressions
* Fix Missing Test & Broken Match Suppressions
* Alerting chart and detail fixes
* Add how-to-test comment to baseline backfill
* Fix inaccurate activation details by raising rule response limit
* Update baseline backfill to use global var
* Make alarm start time the minimum event end time

**Cloud**

* Add beta flag for cloud export metrics collection
* GCP metrics - sidebar and backend
* Azure VNET Peering path computing.
* Kentik map redirect from gce to gcp
* AWS China Support
* Disabled AWS Traffic KDE debugging links
* Update gcp dallas, columbus, madrid latLng
* Weather map error
* Specify project id in cloud export config
* Kube - Use new naming scheme for device\_name prop in metadata, revert ui change from #22309
* Fix Cloud Interconnects Link Click
* GCP - Fix missing cloud router status/nat gw bug
* OCI Show Connections functionality
* Fix kube device name with new [env]\*[region]\*[device\_id] scheme
* AWS show path to should prioritize associated route tables over main ones.
* AWS/Kube Error when clicking on showPath to link
* Change checkbox label for default networks
* Undo sudo restriction
* GCP NAT Gateways
* AWS remove unnecessary eks cluster service
* Kentik Kube 1.1 UI upgrade
* OCI metadata fetching expanded
* Kube: Traffic details error
* Subnets Menu connectivity checker link launches cloud dashboard

**Core**

* Revert api internal deploy action
* Fix navbar wrap issue
* Network Explorer defer queries for tabs
* Performance improvements around API Response (Success/Error) Popups
* Update puppeteer version
* Do not check for down agents on companies with no agents
* Credential Vault RBAC permissions
* Journeys: Fix NMS Check
* Get company\_id from $auth.activeUser.company\_id, not $auth.company\_id
* Add the ability to deploy to iad1 only
* Exclude.cache directory from git
* Prevent collapsed rows on exports
* Update deploy deadline param
* Fix Rate Limiting Status Code for query api
* Agent down count targeting wrong method
* Added a dashboards search sidebar
* Connectivity Cost: Made provider filter more exact
* Adding script file name to job complete log message
* New MKP groupings simplified
* Optimize $companySettings fetch
* DeviceStats Performance improvement - switch augmentDeviceStats to use mGet
* Device Stats optimized augmentDeviceMetadata
* Prevent NMS Metrics Plan on Flow Devices
* Upgrade eslint version
* Device stats type mismatch
* Kentik Agent deploy hide next button on first step
* Store entire dashboard config instead of ID on public share
* Journeys: Add DDoS Alerts
* Reduce permissions exposure for shared links
* Exclude Ksynth and Metrics\_device from device queries
* “onboardingWarning” styling slightly off
* Fix up logging key explosions in kibana
* Fix Device Stats showing SNMP
* Remove verifyDeviceReady check in updateDevice
* Use a more robust method for determining if a deviceID is an IP
* Agent down count targeting wrong method
* Add server sorting to paged device and interface tables #deploy
* Update “Edit Dashboard” dialog title
* Support: Updated to fix for/from customer
* Speed up device collection fetching.
* Pin jobs to portal-nodes only
* Cleaning legacy mysql audit
* Check for ENABLE\_INTERFACE\_CLASSIFICATION\_ENGINE true value
* Field shouldComponentUpdate now needs “unsafe” prop
* resolve circular dependency for auditlog
* LabelList shows filtered labels first
* Make better use of webpack chunking
* Warning when Universal Agents are down
* Update app name category for /v4/export
* Credentials - Templated Credential types receive UI tags
* Credentials - Move credential key/value errors to specific area
* Backend credential name validation loosened
* Fix version info
* FieldState should be updating value earlier
* Credential Vault Internal API
* Fix getting devices ready for flow
* Add SDM Leader device create to internal
* Network Explorer export now corresponds with aggregate/time/metric settings
* Add spoofer debug info to SQL Query Editor
* Auto-create metrics\_device on company creation
* Fix metrics device creation
* Migrate to earthly kbt
* Circular dependancy error in query audit log
* Fix wrappers, by prefixing arguments with dashes
* Enable caching for webpack
* Fix alert notifications
* Add jobs to nomad
* Update add company status check and 0x00 cleanup
* Fixing typo in project.yml
* Update [README.md](http://readme.md/)
* Update redis version
* Snmp hide community with v3
* [Credentials] Frontend polish
* [Credentials] Reusable and improved validation rules
* Scroll fix to Library
* API Internal Migration
* Fix horizontal overflow caused by Site Role selection on Device form dialog
* Prevent quotes in kvs input from breaking updateInterface
* Use the new labels style in Dashboards View
* CRUD for NMS devices via API
* Fix Internal API json marshaling for interfaces
* add node-job-cron to the pipeline
* Wrong label style for credentials filter
* Agents filter requests
* Add chfproxy filter to plan fps in licenses page
* Add Namespace for custom dimension internal serializer
* Internal cd serializer pt2
* Audit Log: 503 Error Fix
* Revert to old version of puppeteer
* Exports hotfix
* Fix DeviceDetails header layout when Labels are present
* Fixed Share URL Mixup
* LabelList fixedWidth wound up in the wrong place
* Device more info sidebar should only show recent kProxyAgents
* Credential Vault Duplicate key handling validation (backend)
* Use AgentCapabilityServiceClient.UpsertAgentCapability instead
* Credential Vault Duplicate key handling validation (frontend)
* Port batch route from chnod\_v2
* Show a warning to Admin users when no Super Admin exists
* Allow docking in DE like the others pages that use queries
* Increase limit for insight utilization search and optimize relevant UI
* Basic test coverage for auditLogger
* Default metrics - w/ preview popover
* Journeys: Enhance KB prompt and add HTTP(s) few-shot
* fix errors from device api
* Refresh SNMP queries properly
* [Default Metrics] fix shadow and disable button / popover on empty defaults
* Journeys: Add few-shot example for two stage filter
* Journeys: Prevent UI from executing queries on everything
* Exports failed fix
* Create new job
* Temporarily disable api-internal deploy

**NMS**

* Properly sort discoveries by creation date
* Update enterprise/model data for nms discovery
* Make Device status Tag sizing consistent on Devices screen
* Fix SNMP status handling
* NMS clean up the Device Details header to accomodate longer names
* NMS: Fix interface ip address display (v4/v6)
* Fix errors on device/interface tables when server sorting
* Hide NMS Policies & UI if Metrics Unavailable
* Update enterprise/model data for nms discovery
* Fix fetchUpdatedOnly sending local time
* Safety check on lastFetched
* Add NMS device menu item for ranger inspect
* Add link to device logs; add agent id; ungate inspect link
* Make adding NMS devices faster and enable parallelization
* Update enterprise/model data for nms discovery
* Make Site/i\_device\_site\_name available as a dimension in ME
* Better handling of existing devices found by discovery
* Make sure single IP discoveries display results
* Update enterprise/model data for nms discovery

**Synthetics**

* BGP Route Viewer - Store Migration
* Fix Synth API Test Frequency
* Upping trx limit
* Fix Agent Results Path View Time Range Selection
* Fixes Shares With Synful Panels
* Fix Incident Log Duplicate Alarms
* v2 Alarms-Compatible BGP Test Results
* Remove Synth Application Constant
* Integrate v2 Alarms into Alarm-Driven Synthetics Features
* Legacy Alarm Responses Can Be Empty
* Incident Log - Fix Test Detail Dates for v2 Alerts
* Fix Synth Error Reports
* Aggregate Incident Log Alarms
* Suppressions UI v2
* UpdateSynthTests Updates
* UpdateSynthTests Bin Script Updates
* UpdateSynthTests Bin Script Improvements
* Apply "test" mode to alerting 2.0 rules in synth tests
* BGP Monitor Undefined Prefix - Take 2
* Add too many requests http code

## March 2024

**Alerting**

* Support any measurement on up/down policies
* Change misleading Clear All to Reset To Default on filter sidebar
* Small Active Alarm Callout handling fixes
* Adjust logic around Alert Detail Insight Chart so end time better represents alert state and trigger cause
* Support cidr dimensions in baseline backfill
* Add filter alerts interaction via summary charts
* Add currently active alarms callout for Policy Edit
* Always include Device ID & Measurement as NMS Policy Dimensions
* Small null-check fix for Alerting Sidebar Drawer
* Implement Manually Clearing Alarms, both Bulk and One-off
* Implement Up/Down Admin Status Check as a Condition instead of a Filter
* Hotfix: Prevent error when changing threshold policy activation mode
* Fix: Restore Default Thresholds in Policy Form for Query-to-Policy
* Fix broken MKP alert drawer in tenant view
* Fix Alerting CSV export mitigationId mapping
* Small bugfixes for Alerting

**Cloud**

* Azure Topo Cronjob Error
* Remove cloud run from default manual\_setup\_type options, add conditionally for gcp only
* Fix export type config disabled
* OCI topology enhancements
* AWS Map hover performance
* Allow selecting/unselecting all for custom dimensions checkboxes
* Refactor cluster metadata fetch for GCP
* Fetch service gateways during OCI metadata collection
* Azure path computing across subscriptions
* GCP Labels as Custom Dimensions

**Core**

* Update IG memory and stop IC background process running on next
* “node-api-internal” as separate service
* Deploy Kentik Agent dialog tweaks
* Store initialize cleanup
* Dashboard filtering failed when applied from multiple graphs
* OTT: Fixed subscriber count calculation
* Allow new site in agents
* Handle classify device error
* Fixed Access Control Saving Issue
* Credentials Vault - Improve readability of multi-key table entries
* Credential Vault - Should not reset labels when type changes
* Credentials table UX improvements
* Add Interface Classification rule at any rank position
* Fix company settings to load on init again
* Prevent public shares from failing to load if no recon permissions
* Update UserMenu snapshot
* Fixes Public Shares for Synthetics and Metrics Explorer Dashboards
* Remove unused createObject function from old api
* Credential Vault Consumer Status field and preview test fixes
* Bug: Cant filter library by users (Saved Views ignored)
* Kentik Agent Management Demo Review
* Kentik Agent Merge Install and Deploy
* Respond with latest device state after updating labels
* Fix missing recents in Saved Views Sidebar
* Fix join in getDeviceStatus.js
* Fix device archive/delete button
* Setup plumbing for node-api-internal service
* IG test and workflow cleanup script
* Add vault key rotation bin script
* MKP Fixes for Synth Dashboards and Recently Views
* Make filter menu auto size
* Support: Adding enhancements to ticket submission
* Revert "MKP Fixes for Synth Dashboards and Recently Views"
* Swagger from bct1 (Telxius) will include x-CH-forward-v6 header
* LabelMultiSelect component to DRY up Label Select component code
* Label API consistency overhaul
* Removing unused parts of old Agent Management API
* “safelyParseJSON” safely parses passed JSON
* Don't need label filtering on RBAC Create Roles
* Remove Add Site button from SiteSelector
* Show correct value in IC Add Rule dialog's Provider field
* Fix issues with exporting on UI page and exporting with email
* MKP Recently View and Syn Fixes
* Move populate dynamic interface from http to function call
* Bugfix: Vault Audit Log
* Update kentik agent column sizing
* Add browser error report metrics
* Moving device Objects to Device Bundles
* Bugfix: Clean up LabelList, applying ReadOnly in most places
* Add GCP Cloud Run & OCI device subtype for enabled dimensions
* Allow adding saved filters even if there are filter groups already
* Increase IC timeout depending on device's number of interfaces
* Revert IC maxForks testing
* Move UI IC trigger to node-portal
* IC: Use concurrency for classifyCompany
* Turn on label filtering for RBAC Dash and Saved Views
* Bugfix: Dashboard graph selections should not survive external updates
* IC: Reduce maxForks to 1 for testing purposes

**NMS**

* Update enterprise/model data for nms discovery
* Persist and read NMS Discoveries from postgres
* Save discovery jobs and discovered devices to db
* Consider IP address in addition to sys\_name when matching discovered devices
* Add ifindex to Interface UpDown Preset
* Remote SNMP walks
* Add update ranger bundles script
* Ability to create ICMP-only NMS devices in bulk
* Fix check if device uses nms in createDevice
* Fix memory utilization meter with multiple series
* “connector\_present” column was replaced with logical a long time ago
* Fix call to createDevice from saveDiscoveredDevices
* Render metric with unit “duration:microseconds” in milliseconds
* Add new "use this config dir but don't whine if it's missing" arg for ranger
* Query api for nms
* Delete NMS devices from Collection after Settings > Devices are deleted
* Handle bundle generation errors correctly
* Show last value in devices table metrics
* Another attempt to prevent db pool filling up while creating ranger bundles
* Add new config dir arg for local ranger config
* Optimize assignConsumer to keep db pool from filling up
* Top-level interface filter group should use connector:all
* Exclude virtual metrics\_device when checking plan device limit
* Rename Components to Hardware in NMS device details
* Do not append gnmi config block if ST is disabled
* Prevent selection of internal device types
* Fix timestamp mismatches and prevent log spam
* Preserve lookback so live queries work
* NMS threshold policies use a fixed evaluation period based on measurement name

**Synthetics**

* Search bar in synth library availability widget
* Fixing api for update test
* Hundreds Of Undefined Notifications
* Change "Status" in TCC to "Alert Status" to avoid confusion.
* Fixing screenshot counts on error
* Fix (maybe) for dns test edit bug
* Decouple Agent Alert Suppression
* Script to create rules for old synth test policies
* Updated the TCC filter text to be more consistent ("Filter by")
* Pass updated after timestamp to ensure tasks are generated for most recent test
* Mesh shows agent not on test
* Ensure original “cluster\_id” is migrated
* BGP Monitor Undefined Prefix
* Set up test config alerting rules for v2 alerting
* Last item in health timeline always green
* Hide network agents based on form config
* Test doesn't show notification channel until refresh
* Invalid Notification Validation
* Fix has ping results
* Dashboards struggle with some highcharts in safari

**Other Changes**

* Insights: Disables Individual Checkboxes if Company is Disabled

## February 2024

**Alerting**

* Suppressions for NMS policies
* Prevent labels from prepopulating on new policy
* Fix alerting subscription PDF export column collapse
* Prevent conflicting uniques for Policies
* Use synthetic dimension value as fallback for test ID
* Fix “baselinepct” in drawer
* Fix pdf export timeouts
* Fix alerting table group labels
* Hotfix: Infer Alert Limit, Fix Device Links
* Fix NMS site health map count

**Cloud**

* AWS Metrics collection config should only be visible for sudo users
* OCI Custom dimensions autogeneration
* Bin script to set sampling rate already exists, just add OCI param
* Add sampling rate to OCI cloud export config
* Number of pods not visualized correctly
* OCI logo inconsistence
* Increase max payload size to 64MB
* Related Cluster radio buttons should be selectable
* Azure topology minor error fix
* Skip cloud run metadata
* Update GCP Cloud Export Type config options
* AWS CloudWatch metrics on Kentik AWS Map
* Azure refresh topology script update.
* Azure Refresh Topology Summary generation

**Core**

* Make sure IC rule ranks are correct in the database when sorting by dragging
* Hotfix: Ping health data was not consistently deserialized from API
* Custom dimensions must only contain alpha-numeric characters and underscore
* Start IC rule index correctly
* Grid Layout fixes
* Fix exporting when there are url params
* Fix onScroll Error on dashboard drill down
* Bytedance CDN add
* Make upcoming favorite changes backwards compatible
* Agent Management API v2 integration (capability integrations)
* Move recent views to use chwww
* Bugfix: DateRange is no longer cut off at the end
* Better verbiage for customer & providers tables
* Bugfix: Pie chart labels no longer truncated as aggressively
* Bugfix: Firefox top talkers modal had incorrect grid layout
* Update DetectedCaches.js - Facebook\_FNA
* Send subscription emails even if recipients is empty but has bcc/cc
* Fix MKP subtenant view for next.\* and other hostname lookup
* Bugfix: no longer displaying outdated kproxyAgent entries in devices
* Send source info to socket data fetch
* Bugfix: Credentials table tags should be readonly
* Bugfix: Providers state should be coupled to URL changes
* Bugfix: Library checkboxes should not show in customize columns menu
* Fix email sharing report when using urlParams
* Device settings menu now supports Labels column
* PeeringDB related filters in Data Explorer
* Fix undefined error on Saved Views Drawer
* Quick agent updates: null check bug and run state
* Bugfix: Core Grid element changes were too aggressive
* Enable PeeringDB rule in Interface Classification
* Tests for new Agents API, part 1
* Remove sudo check for Embargoed Country Insight
* Turn on metrics for APIv6
* RBAC: simplifying a test
* Added kproxy 7.43 to KproxyAgentModel.js
* Credentials: Added groupby
* DE: Fixed unnecessary shifting of sankey labels
* Capacity Planning runout column sorts properly
* Library: Share menu copy button copies the wrong link
* Agent Management - Leverage new APIs
* Costs: Global charges no longer removed when adding/removing an interface
* Bugfix: Allow users to display credentials while editing
* Library polish
* Agent Management tweaks
* Saved View Shortcut Navigation Drawer
* Make sure label selects don't show presets
* Removed duplicate Label components and alternative invocations of LabelList
* Handle errors from incomplete form when adding credential
* Change CPU column on devices page to avg over cpu\_index instead of sum
* $auth.isSudoer works using app but fails test
* Move creating credential keys to signup
* Added kproxy 7.43 & 7.44 KproxyAgentModel.js
* Core/Select - Use “active” instead of “Intent.PRIMARY” on selected items
* RBAC bugfix
* Only show relevant credential types in dialogs
* Kentik Agent: quick UI feedback part 1
* Agents: just sort capabilities
* Include source prop in query object
* Do not show no flow warning for link share
* RBAC tests w updates
* Clear “device\_snmp\_ip” when turning off SNMP for flow
* Add discrete SNMP IP Address field for NMS
* Prevent KMI for Demo Mode
* Prevent disabled options from getting selected
* Small RBAC Tweaks
* Update Users when updating tenant
* Exports Fail Fix - Added check for export
* Turn RBAC label filters back on. labelFilter: true
* Update add KMR PDB script to use correct fields
* Library: Save As Error
* Persist user page after session timeout
* New User Edit screen to support RBAC m3
* Dashboards/Saved View: Show description button & dialog
* Catch validation errors as 400s in restricted api
* Remove obsolete portal log envs
* Make sure kproxy user exists for kagent to function

**NMS**

* Fix funky sorting in metrics table when “id” dimension is used
* Hotfix: Handle missing measurement model during policy deserialization
* Fix metric selector on NMS charts added as widgets
* Hotfix: NMS Interface Policy Filter Serializations (+ Tests)
* NMS Alerting CSV exports
* Fix IS-IS neighbor matching logic (again)
* Fix error with missing IS-IS neighbor
* Add default timeout to ST settings, remove Interval, refine New Credentials dialog
* Add site, device, and interface filters to metrics explorer
* Support multi-credential Discovery
* Use new agent statuses in nms flows
* Combine merges correctly to prevent truncated results
* NMS Alerting Fixes for Table & Policy Form
* Fix metrics license query
* Return per device errors for saving discovered devices
* Ignore NMS POSTs that are really GETs in audit log
* Adds In Errors and Out Errors charts to the Interface detail sidebar
* Hide merge series options in Policy form
* Add NMS dashboard as a default landing page option
* Update wording in DeviceForm SNMP and ST collection methods
* Disable the start discovery button when the form is not valid.
* Allow editing devices with down agents
* Disable NMS for onprems
* Default sample rate of 1 for nms devices
* Update defaultColumns for Interfaces tables to include Errors/Discards again
* Add “-name=kagent” to Docker KAgent install command
* Fix sparklines not showing for values of 0.
* Hotfix: Handle rendering for unexpected BGP session-state values
* Add “-pull=always” to Docker KAgent install command

**Synthetics**

* State of the Internet DNS Services loading fix
* Error when cloning test with non-zero port for ICMP ping
* Hide waterfall footer again
* Allow Override of Agent Alert Start/End Time On Load
* TRX bug fixes
* DNS test agent detail null checks
* More Trx Bugs
* New chrome recorder launch option
* Fixes undefined metricsToDisplay
* Align date state between test results and agent test results
* Remove task status on agent deletion setTask calls
* Pass task status to set tasks for tests
* Fixing typo
* Clean up use of results store following fork for performance improvements
* Fix capitalization for "DNS Resolution" on DNS test results page
* Call loadUsageData after getting tests
* Update references to test.results -> test.agentResults
* Decode test URIs in a try/catch
* Agent Results Page Updates
* Missing table columns
* Credit usage query
* Updating the credit usage sql
* Add 191873 cid for mesh/a2a agent alerting
* Agent Details Direct Link
* Notification channels no longer applying to test
* Fix agent config modal
* Fix Agent Downtime Alarms
* Adding optional chaining on agent config modal error
* Clean up cloud logos

## January 2024

**Alerting**

* Expose Policy ActivationMode
* Hotfix policy editing
* Alerting Manager Integration for Policies
* Alerting Manager: List Alerts
* Small fixes and cleanup for alert manager integration
* Alerting NMS fixes and improvements
* Update Site Health Map to use AlertManager for NMS Alerts
* Add Policy Application and Type to NMS Notifications
* Improve NMS Notify Dimensions Names for JSON Webhook
* Activation Settings Improvements to Policy Form (Fetch Rules with Policies)
* Include Context in Alarms List
* Disable memoization in EVM service
* Only show tenant notification channels where appropriate
* Merge $attacks into $alerting store

**Cloud**

* Increase GRPC payload limit
* Fetches instances as part of OCI metadata topology
* Use us-east-1 region when generating list of accounts in AWS organization.
* Change azure express route sidebar query
* Enable cloud run log collection
* OCI Monterrey Mexico cloud region
* AWS fetch managed prefix lists and their cidrs
* Add GCP VPN Gateway Traffic sidebar widget
* Auto add metadata collection for all accounts in an org using principal role

**Core**

* Prevent DVW from issuing query when no flow
* Added filter for labels
* Fix mock data url cache/hash
* Show Sflow Interfaces like regular interfaces in DE queries
* Correcting dialog text for removing a crendential
* Correct breadcrumb for Credentials Vault
* Correct spelling of “policies” on label settings page
* Node job scheduler service
* Site Settings: Address component overrides
* RBAC M3 UI Implementation
* Site Settings - Aligned fields on one line
* Kentik Agent: make deploy consistent with edit screen
* Add AWS Instance Tags to nosplit
* Shut off RBAC label filters: Just for temp!
* Attempted fix for cost job having some zero values
* In Kentik Agents path: Show Hostname when Agent name is missing and other improvements
* Journeys Storage, Feedback, and LLM Update
* Fix bug in NLQ user ratings
* Fixed CDNs List table overflow
* Moved Classification Types field to the right
* Add site ip to csv export
* Move CDN analytics time range and configure buttons to the right side
* Credentials vault templates for SNMP, BGP
* Enable Credentials Vault for all
* Fix mock tests
* PeeringDB IC - Use correct Provider name
* Add special application name to query editor queries
* Fix issue where exporting preset items might fail
* Fixes Interface Classification issue if company has no rules
* Spread out forkChain promise array
* Update KproxyAgentModel.js w/ kproxy 7.42 commit hash
* Adding docs about adding new IC things
* Add chat history to Journey & re-write DE NLQ
* Added CDNs list title
* GH action deploy node service in parallel
* Update branch checkout version
* Fix Interfaces page locking up
* Only get what we need for OnboardingWarning
* Move Public Shares under settings
* Update edate on custom application update
* Add debug info to DataViewWrapper for Sudo-ers
* Add script to create PDB IC Rules for everyone
* Fix OTT service type to category
* Fix for the 1 second we wait for “refreshCapabilitiesByAgentID”
* Fix Device model deserialize
* Use LabelSelector for Devices
* Pulling from the updated collection
* Bin script for classifying device manually
* Add structured logging for live update in query engine
* Fix interface connectivity type api
* Fix incorrect logic in adding new IC rule
* Update release notes for NMS label
* Limit IG worker count to 1
* RBAC: Changed button “Add role” to “Add Role”
* Update connect-redis from v6 to v7
* Proposed git ignore
* Color popover for add Label improvements
* Use Active User companyId for label filtering
* Tests workflow: try using a bigger machine to avoid running out of memory
* RBAC Label filter dashboards and saved views
* Increase number of decimal places shown for costs depending on configuration
* Migrate internalApi from v3 to v4
* PeeringDB - Interface Classification rule
* BGP Community Add Pipe character. Fixes #20256
* Kentik Agent UI polish
* Setup IG env\_var for deploy env
* Only doing label filter for Saved Views and Dashboards

**NMS**

* Fix "setup NMS" completion checks
* NMS Credential cleanup: add validation, add v1 type, improve optionRenderer
* Allow creation of devices with no sending\_ips or sample\_rate if kentik agent collection SNMP is enabled
* Propagate credential changes to agent
* Fix bad redirect after no devices returned after filtering, fix form validation
* Fix status badge misrepresenting the actual status
* Clear starting and ending time when using lookback
* ranger 191: Update relative paths in configs we send to ranger.
* Add Docker deploy steps for Kentik Agent
* Change default ranger snmp workers to 3
* Fix broken “device\_snmp\_ip” validator
* Show NMS credentials as in use
* Update default sort field in MetricDeviceCollection, remove cpu\_sort
* Return no results when measurements haven't been initialized yet
* Hide the capabilities section on ME device details
* Add start and completed times to a Discovery result
* Convert NMS Alert History & Active Widgets to AlertManager, UI Improvements
* Refactor “snmp\_collection\_method” logic in deviceSerializer
* Fix: NMS Alert History Widget dataMax undefined
* Hide the total row on select NMS Dashboard widgets
* Fix filter in components queries
* Recover streaming updates when sockets die
* Add capabilities details for Ranger
* Keep metrics devices collection in sync when devices are added and removed
* Support for NMS/Kmetrics AlertManager Notifications
* Prevent undefined on y-axis
* Remove endpoint address from ranger config
* Fix validation issues on mps queries
* Better handling of discoveries with expired keys
* Don't fetch agents when not allowed
* Agent Capability Descriptions
* Create NMS policies for all
* “fetchExistingDevice” fails with a sql error in iad1
* NMS device form improvements
* Change wording on NMS devices callout
* Add "Device ID" column to Discovery Details page
* Fix device creation and modification bugs
* Fix syntax for gnmi/tls/skipVerify in ranger device config
* Generate proper query for NMS alert baseline backfill
* Allow excluding addresses from discovery with a '-' prefix
* Remove broken filter to get cpu/mem stats to show up in devices table
* Add better check around /interface/counters
* NMS Disco Ranger Config
* Fix bad redirect for when filtering device collection returns no results
* Filter out empty values that could have snuck in
* Make ranger config files mode 600
* Loosen up sudo checks around discovery/agent routes
* Allow NMS customer to access NMS setup page
* Fix redirect to NMS setup page
* Use correct field names for snmpv3 “authentication\_type” and “privacy\_type”
* runtime-config.yaml needs to have the binary specified
* Fix resolution of setup merge conflict
* Handle mismatched device names
* Properly handle api address for next envs
* Prevent creating a redis connection on every discovery socket request
* Prevent interface table query from filling up query queue
* Fix crash when attempting to hide NMS task
* Fix missing filters field from modified filters object
* Make sure filters exists before modifying it
* Look up vendor and model for sysobject id from discovered devices
* Break NMS queries into 3 day chunks
* Device Discovery
* Workaround for no results on NMS landing page
* Filter out “sub\_type: 'kagent'“ in relevant spots
* Metrics Schema - delete buttons for measurements, metrics, & dimensions
* ranger 124: Copy agent, company ID fields from DiscoveryResultRequests to redis
* Copy errors and devices.credential\_name from discovery result to redis

**Synthetics**

* Fix capitalization on Credentials Vault
* State of Internet - No Results
* Hotfix TCC health status w/activity check
* Fix Agent Suppression Edge Case
* Fix mesh crash on dashboard
* Fix agent downtime time changer
* Revert setTimeout bump back to 1 sec for setTasksForTests
* Fix for traceroute paths
* Expect test results can be absent, provide meaningful view state in such case
* Improve performance of parent test screens
* Filter subtest alarms appropriately on agent details
* Update Agent Suppressions During a2a and mesh Test Lifecycle
* Hotfix timeline
* Incident log and test results improvements for subtest alarms
* Force kentik prod to use subtest alerting
* Update prefix in form
* Export list of synthetic tests from Test Control Center to CSV
* TCC status should be driven by alarm status rather than health
* Change Health Bar to include alerts in addition to health status
* Agent Alert Suppression UI
* Synth health status summary and availability widget should be driven by alarm health
* Match collect syn query
* Remove IBM export creation from cloud settings
* Pagination of trace results, no dns lookup on private ips
* Quick patch to validate target as ip before setting
* Add reciprocal property to AgentTest
* Add use\_local\_ip bool to IP test for network\_grid
* Add Target Name and Subtest URL to Synthetics API Response
* Synth subtest updates
* Fixing lint unused var
* Availability widget doesn't seem to work
* Bump setTasksForTests timeout
* Mesh tasks might be missing valid targets
* ThresholdSeconds stored as string
* Update to fixes for IBM mesh policies
* Fix for duplicated browser agents
* IBM activate hotfix
* Per-Agent Test Alert List
* Script that will fix period value for affected tests
* Changes to ui-app where customers can configure agent alerting rules
* Better url validation
* Regex refactor2
* Syn: "Traffic with" DE button is using incorrect filters
* BGP route viewer ASN fix
* Update v2 model per v1 updates
* Fix error when adding a new agent
* Using target\_agent\_id to replace agent-ip map, bugfix
* config.test\_type is not a thing
* Not all tests with ping tasks will have target agents
* Agent Alerting Rules Errors in Non-Private Agent Editor
* dfw1 mesh alerting v1 updates
* Removing all mvp stopgaps for cred vault
* Fix missing onClose fn w/ no op fn
* Alert rules can be disabled
* Fix: Set task status to be the same as test status

**Other**

* DE NLQ: Add filterValue enums, post-processing, and bug fixes
* Add device subtype descriptions for LLM
* NLQ name change, polish
* Query Assistant: Include measurement description in schema

---

© 2014-25 Kentik
