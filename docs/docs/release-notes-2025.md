# Release Notes 2025

Source: https://kb.kentik.com/docs/release-notes-2025

---

Kentik employs a continuous deployment methodology for constant extension and refinement of the Kentik v4 portal and the underlying Kentik platform. Release notes for each successive month of Kentik v4 updates are covered in the following topics:

> **Note*:*** For additional insight into what’s new with Kentik, check the **Product Updates** page of our website.

## December

**AI / Insights**

Fixed

* Kentik AI editable default model field

Added

* A “+” button to open the dialog in a full screen
* The ability to ask for diagrams (e.g., mermaid diagrams)
* Export to Mermaid diagrams to AI Advisor
* New Runbooks tool that allows AI Advisor to access and retrieve runbook instructions for answering questions and troubleshooting alerts
* Data around Network Classification and Custom Applications tool
* More Policy details to Alert context

Updated

* Kentik AI settings to allow admin users to view with certain fields disabled
* AI Advisor Public API
* The Sitemap Tool for contextual page information for AI Advisor
* And enabled the new Synthetics tools

Started

* Initial clean-up of Journeys

**Alerting**

Fixed

* Notification OAuth 2.0 credentials key max length to 100
* MKP tenant and template policy creation functionality
* Missing edit conditions button
* Alerting pagination and improved export functionality
* Select all checkbox for the alert table
* Alerting pagination for initial load

Added

* Sudo view to list and delete MKP policies that have been created without being properly associated with a tenant
* Support for canonical application policies, alerts, and notifications

Updated

* Notification channel chips
* UI and functionality of MKP Parent Policies
* Notification channel preview and test
* Alert Config Store utilities and wrappers

**Cloud**

Fixed

* Logic for cloud task status

Added

* New CloudPak card to the Licenses page to replace Cloud-specific cards for AWS, GCP, and Azure
* Added navigation and side bar opening from Entity Explorer, and scroll for AWS, OCI, and Azure maps
* Added timestamp for Azure to Data Explorer
* New flow metrics chart to the Public Clouds page and updated measurement names used for the Licenses Page and Public Cloud Page queries
* A new OCI region

Updated

* Enabling autogeneration of custom dimensions for OCI and GCP-only customers

**Core**

Fixed

* And/Or buttons for Interface Classification settings
* Device flow fetch
* Run report with WebSocket request looping
* IP Details page
* Universal Agent Capability config options
* Sankey in the new AS path collapsed dimension
* Missing Metrics Explorer page to Site Map
* DNS Device Creation process
* Syslog Device sidebar filter
* Creating Public Dashboard shares
* Missing hidden populator API methods
* Interface queries to load faster

Added

* A low-resolution login background
* OTT Category icons
* New API endpoint for retrieving network classification data by company ID
* IP validation for DNS Device Create
* Universal Agents in-page context for AI Advisor
* Added missing page to the UI Site Map and removed unnecessary “Name” column
* Support for bulk device config setup
* Config information to the devices table
* Metrics for `ktraffic` capability
* MFA filter sidebar to Users Page
* New Syslogs tab in Device Details
* New Credentials Overview Collection for the Devices Page

Updated

* Device Configs fetch and display
* Universal Agent to allow UI to configure environment variables for capability configuration
* Metadata description as help text for fields
* Efficiency and reliability of the dynamic interface group population process
* Legacy Map to exclude DNS server devices
* Default that WebAuthn is enabled
* Device Details performance
* Device List page to prevent erroring after adding a device

**Network Monitoring System (NMS)**

Fixed

* Device deletion and archiving workflows
* Letter casing in Metric Filters dialog

Added

* AI context to the Overview and Traffic tabs
* New device filtering capabilities to the Devices page, specifically allowing users to filter devices by Universal Agent capabilities and UA names
* Ranger interface metrics support for device SNMP detection and ranger capacity count
* `kmeta` (Metadata Collector) capability to the Universal Agent page
* Context name for SNMPv3 credentials
* A download button in the Device Details Configs tab
* Search feature to the device configuration view

Updated

* Hostname dispaly of IS-IS neighbor if no matching device is found
* Device warning logic in the Capacity component to treat “ranger” as an enabled SNMP method
* Configs tab to be restricted to NMS devices only

**Synthetics**

Fixed

* DNS Target Field to show DNS results that occurred prior to updating the target field

Added

* A public-facing API for users to export the syngest test results as CSV
* `ksynth` capability to Universal Agent UI
* Alert Status filtering to the Test Control Center page and made the Test Status Summary Status buttons link to TCC filtered by status

Updated

* Synth Path View

## November

**AI / Insights**

Fixed

* AI Advisor (AIA) logic for the “thinking” indicator and scrolling function
* Chat date time to align with user profile date settings
* Trap tool device dimension

Added

* Timestamps in each user message

Updated

* AIA to restore prompts after a network/server error occurs
* AIA device tool to optimize results
* Removed Ask Eddy (Kentik KB’s AI agent) references
* AIA Connectivity Test tool
* SNMP Trap tool to accommodate for NMS changes

**Alerting**

Fixed

* Issue when failing to find relevant tenant policy
* Alert Page result count and filter interactions
* UI components on the Policy Table and Form
* MKP API 500 errors
* Transform mitigation apply and clear timeout unit to seconds

Added

* Enable/disable policy to the audit log
* HA support for Radware platforms, toggle switch to enable/disable, and input for a secondary IP

Updated

* Checks to determine when to enable/disable tenant policies
* Threshold UI for the Add Traffic Alert Policy dialog
* Alerting public API for alert listing and retrieval by ID with mitigation ID
* Filter options to only allow filtering by tenants with subpolicies
* Templates for all policy types

Removed

* Outdated policy limits

**Cloud**

Added

* AI Advisor context to the Azure, OCI, GCP, and AWS Cloud Map
* New OCI regions

**Core**

Fixed

* WebAuthn issue when trying to authenticate
* Spoofing with WebAuthn
* HTTP URL metrics
* Daily stats
* Cleanup of long-running worker processes when sockets disconnect or when users disable live updates
* Loss of context from incorrect unspoofing

Added

* New non-prepending AS Path/Hop, IPv6 flow label support, and no overriding of generic ASN lookups with device BGP data
* Skip button on 2FA verification selection screen
* Kentik Site Navigation to the existing Global Search functionality
* New Interface Classification types
* New warning popup when a user changes OTP state, a warning dialog with CTA on the Authentication page when a user is in this state

Updated

* Interface Classification KMR to remove “New!” tag
* Usage of TOTP in SudoDialog
* Universal Agent to prevent OTT Tap DNS Failure due to flag issue
* Traffic Costs to use the correct filter for OTT related queries
* Monitoring and metrics for WebSocket connections and RealtimeBus query subscriptions to improve observability
* UI Site Search
* WebAuthn corner cases
* Limit to NMS interface groups to devices that are assigned to the monitoring templates
* Dynamic Interface Groups functionality

**Network Monitoring System (NMS)**

Fixed

* How associated devices are displayed for the Ranger capability in the Kentik Agent settings

Added

* A new combined utilization chart for network interfaces to enhance the visibility of both inbound and outbound utilization metrics in the interface details view
* AI context to Device and Device Details pages
* Fetch endpoint that returns the respective columns given a specific trap type, OID, and template name

Updated

* Logic and UI for displaying device usage warnings and progress bars in the NMS Device Usage Card component to ensure warnings are shown for both full and non-full licenses when the device limit is reached
* Device name link to render if it is a dimension field for an NMS alert
* Utilization chart Y-axis range to be fixed between 0 and 100, instead of dynamically adjusting based on the data

**Synthetics**

Fixed

* Synth Test CRUD API
* DNS Ping Port update

Added

* Enabling audit log entries for toggling of synthetics test status

Updated

* Agent test metric page to support CSV export
* Synth error task handing to track error states for individual tasks
* Labeling from Agent errors to Task errors, placement for the Task Error column, and alarm timeline tooltip to show a breakdown of Task Errors when applicable

## October

**AI / Insights**

Fixed

* AI Advisor page title from the previous view sticking and no page/bookmark title

Added

* Spoof user column to AI Advisor

Updated

* Prompt for Data Explorer CIDR filters
* Related AI permissions to turn off when enabling Disable Kentik AI
* AI Advisor tool security

**Alerting**

Fixed

* Mitigation subscription filter and initial load for mitigations table with All Time lookback
* Context search pagination interaction
* Required validation for zero (0) value in NMS native policy form
* Missing package policy bug, Data Views tab submit that was disabled if “all devices” was selected, and population of Configure Policy link

Added

* Alerting RBAC permissions and RBAC permissions to the public alerting API
* Default roles for non-KMR users
* Auto-refresh support for the admin table
* Device label information for NMS policies in the notification enrichment layer
* Note for users indicating default Content-Type for custom webhooks
* Mitigation method edit for A10 and Radware

Updated

* To include all types of NMS alerts when requesting native NMS alerts and removed the NMS-native only button and its logic
* Sorting and grouping with pagination and removed infinite scroll for alerting
* Support for “All time” filter for active mitigations
* Adaptive Flowspec preview to load alert details first, before sending a request to endpoint

Prevented

* Mouse wheel from incrementing numeric inputs while scrolling

**Cloud**

Fixed

* Property name in the cloud export edit screen to show GCP sampling rate
* Resolved null error when showing neighbor Transit Gateway ID
* Stuck load topology window for AI Advisor
* Format for NMS device labels notification
* “Invalid Date” message on mitigation filters datepicker

Added

* Null checks to OCI Entity Explorer Widget
* New Canada West region to AWS region list
* Azure inventory tool, resource graph query tool, and updated sidebar details to display new tool results
* Protection for missing gateway attachment in map sidebar

Updated

* Handling of AWS accounts in the AWS Inventory tool
* Sampling property to be false for AWS default unsampled cloud export
* OCI SDK, fixed summary generation, and excluded non-existing regions

Prevented

* Alerting test from failing inconsistently

**Core**

Fixed

* Universal Agent to allow spoofers to install `ktraffic`
* AWS topology fetching
* Incorrect field type for the new version of site API
* Various Kentik Market Intelligence (KMI) elements
* Missing unit in MetricsResult
* Spoof for region-restricted users
* Interface Classification rule evaluation
* Incorrect targeting of 2FA while spoofing

Added

* Missing device flow routes
* `ktraffic` as a capability
* Metric renderer for Kentik Agent ID column in Data Explorer
* `ktraffic`, licensing refresh, and legacy `kprobe`
* KMR to create system user
* DNS Analytics Plan type throughout the application and Create Device Route
* Display for `kprobe` device subtype in the device settings form for users with appropriate permission
* Synth user to defaults
* New WebAuthn page for all users
* SSH private key and password credential types

Updated

* Filter for AdminStatus and OperStatus for the Device » Interfaces tab
* Page title for Data Explorer browser tab
* To populate `kflow` map during unexpired flow
* Agents unification efforts

Prevented

* 500 error when grouping interfaces by “Interface Type”, “Connectivity Type”, and “Network Boundary”
* Device Interface tab filters from being removed when changing group by
* KMI page from erroring is invalid ASN is provided
* Double focus style in AI Advisor input field when in dark mode
* Lookup Kentik Agent user for DNS device create

**Network Monitoring System (NMS)**

Fixed

* Test connection button by removing outline prop

Removed

* Extraneous dimension filter in MNS capability slide-over

Added

* Scaffolding page for the License refresh effort
* The ability to edit/save an Npak type plan

Updated

* Enterprise/model data for NMS Discovery
* Test Connection button style
* Device page header to show ICMP devices
* NMS Device Usage card UI and links

**Synthetics**

Fixed

* DNS error on 100% packet loss (ping)
* Partial/first timestamp results
* Synth UX to ensure credits utilization and projection boxes maintain a consistent height and loading state in Test Control Center, and enabled scrollability of content in parent test results

Added

* Support for grouped alerting by label and site
* Host as an additional property in DNS alert details sidebars

Updated

* Credentials Vault popover
* Shading if validation is present for agents without local IP
* Legacy Alerting Types detection
* Network Error Status 500  when creating synthetic tests
* Showing the resolver as the destination in a DNS test’s incident log destination
* Synth translate and Synth subtest results alarms

## September

**AI / Insights**

Fixed:

* Active alert count for AI Runbooks
* Logic to convert resourceID to number
* MKP 404 - Runbook fetching on alerting page
* AI Advisor alerts empty results
* Cause Analysis to not display options if Kentik AI is disabled
* Sankey overflow

Added:

* A thinking indicator to AI Advisor when a step is in processing status, but no response has been generated yet

Updated:

* Runbooks dialog heights and improved row calculation for Markdown preview
* Syslog tool to increase default top-x value and convert timestamps to UTC

Removed:

* aiModels permission dependency for custom context

**Alerting**

Fixed:

* Disabling flow policy via public API
* Baseline percent descriptions on the alert details view
* Schema and removed application from public filters
* Available actions for mitigations
* Notification channel OAuth form validation errors
* Server-side table sort
* Alerting Table date range filters
* Notification channel OAuth form secret handling
* Alerting Public API CRUD operations
* v5 API manual mitigation prop names
* Alert subscriptions filter error

Added:

* Missing NMS public API work
* New Alerting Public APIs to API tester
* Expiry date to policies
* Proper error handling to public API v6
* Validations and fallbacks for Public API service parameter defaults
* Navigation bar to the Alerting and Alerting Overview pages

Updated:

* Context search public API
* Alerting public API baseline data and value calculation for conditions
* API error message and page interactions
* API interaction
* Policy link component to send user to filter table by ID with drawer expanded if the user does not have edit permissions
* API headings from “token” to “key”
* Support in types and notify requests for context and wired a pass-thru synthetics notify translator
* Alerting table context search integration with search bar
* Suppressions page to extend to see, filter, and list silences
* Subpolicy status to match tenant status on save
* Test notification button to be disabled if a channel is disabled

**Cloud**

Fixed:

* AWS Map Account selector
* Selecting the correct previous Azure SPN record when editing Azure cloud configuration

Added:

* DRG traffic queries for OCI sidebar details

Updated:

* AWS Cloud default exporter sampling type to ‘unsampled’
* Inventory error handling
* Cloud Performance Monitor page by removing “No Direct Connects Found” banner if there are no Direct Connect Entities but there ARE Direct Connect Gateways
* UI logic for showing sidebar details when clicking on the grid of the AWS Cloud Map
* OCI onboarding UI to make Cross Tenancy Policy the default, and OCI edit cloud export screen

**Core**

Fixed:

* Saving credentials with webhook notification
* SSO setting save on disable toggle
* Support for app protocol columns
* DeviceTelemetry
* Onboarding import
* Error when country lookup fails
* Issue where User without Roles could not be viewed
* Persistent console warning messages
* Typo in SearchResults, dark theme appearance issue, and added error boundary to GlobalSearch
* Admin login reason to populate correctly
* Column filter count
* UA Clustering API calls
* Labels page crash when rendering company with package labels applied

Added:

* Icon next to agent name that belongs to the cluster
* Device count column to site table
* Audit log mapping for AI Advisor chat session deletion
* Missing MKP RoleSet route
* WebAuthn to possible authentication methods during login, and is automatically triggered if desirable
* Yubikey login as a fallback if no 2FA method can be found
* Filtered device links on the Settings » Labels page
* Universal Agent dynamic capability configuration support
* Classname for pendo to the table search bar

Updated:

* Device and Details pages to support agents clustering
* Device page group by agent with proper cluster name
* Device Add/Edit screen for UA Flow Proxy
* Interface classification speed
* User Admin page bulk action for Role, RoleSets, and default landing page
* Display of the device toolbar menus until the device is loaded
* Expand results to fully expand when filtering/searching in one category
* Agent clustering to show and filter clusters and provide a simple dialog to see ALL the Cluster Data
* Kentik Agent to properly allow deleting OTT configs
* Display of clusters
* Handling for admin 2FA
* Maps API reliability
* Handling for device deletion
* Users table to add Role Sets column and filtering

Prevented:

* Unnecessary agent lookups when grouped
* Missing Site Traffic for Connectivity Costs
* Removing one data source from removing all in Library Dashboards
* Device name updates from causing malformed dashboard saved query
* Duplicate saved filters for Presets Company
* Error in Labels page relating to pluralizing

**Network Monitoring System (NMS)**

Fixed:

* Removed duplicate form field for NMS Template
* Prevented incorrect monitoring template assignment
* Transceiver metrics in the interface table
* “Out Errors” labeled as “In Errors” in NMS interfaces

Added:

* An option for APIs that lookup devices by sending IP
* Test Connection button to the SNMP tab, which can only be selected/run if all required fields are provided

Updated:

* UX of the new SNMP collection assignment workflow
* Use of metrics plans that are active when adding discovered devices
* UA tooltip language
* Enterprise/model data for NMS discovery

**Synthetics**

Fixed:

* Cloning Synth panel to existing dashboards
* Synth test trace API

Updated:

* Agent health timestamp
* Support for grouped alerting by agent and updated incident log to reflect new alerting type
* Grouped alerting navigation improvements

## August

**AI / Insights**

Fixed:

* Time series to use AS number filter instead of AS Name
* AWS traffic queries for customers without cloud devices
* Enabled model config in demo box
* Journeys permissions

Added:

* Ability for Sudo users to access INTERNAL LLM models in Kentik AI settings

**Alerting**

Fixed:

* Alerting public API v6
* Regex matching for notification channel custom headers

Added:

* Support for TLS client certificates to custom-webhook notifications
* Support for time unit dropdown for manual and threshold mitigation TTL and timeout
* Device session status in BGP announcements table

Updated:

* Mitigation platform status table to be scrollable
* To allow spaces and simplify http header regex for notification channels
* “Delete” to “Remove” in Alert Policies UI for consistency
* Alert status to update in all instances on the Alert Detail page
* Keeping IP CIDR value when switching methods
* Alerting policies by disabling Route Prefix+Length filters

**Cloud**

Fixed:

* Missing AWS API status results

Added:

* Policy example when creating exporter
* Azure UX to support Azure SPN configuration

Updated:

* Exporter to make OCI user field optional, list of necessary permissions for OCI exporter
* AWS Entity Explorer “Open in Map” to navigate to map and display entity traffic in sidebar
* Azure regions list to ensure up-to-date region processing

**Core**

Fixed:

* Time format of inspect traffic queries
* Test import issue
* Select tag for filter and date picker
* My Kentik Portal (MKP) brand settings functionality
* Dashboard panels stuck in the loading state
* Missing `E` override markers on the UDD `InterfacesTable`
* Device creation date showing current date instead of actual creation date
* Interfaces created by v6 testing
* MKP support form validation
* Flow tag API v6 and migration
* RBAC Label Permission with multiple roles
* Timing issue when logging some spoofing events
* Live update for SNMP queries
* SSO setting save on disable toggle

Added:

* Missing interface type conversion logic, API layer updates, and updated data model
* Check for valid CSV input before parsing into array
* Missing CID metrics to api/next/v5 routes
* Device Handler API v5 to v6
* x-v5-migration header to existing API v5 requesters
* Vendor logos
* Rejection for API requests to create or update devices with pseudo-device subtypes
* Renderer for JSON KVs values in Data Explorer to show the new “Parsed KVs” column for SNMP Traps
* Predefined RBAC roles to users with custom roles only
* Universal Agent Capability edge case
* Host metric to Universal Agent Capability Details queries
* Expanded rows in grouped tables

Updated:

* Devices page to show spark lines when loading page with existing filters
* Query observability with context
* `kproxy` capability instructions
* Universal Agent
* Click labels in the header of the Devices page to see a list of devices with that label
* Live update query observability fields
* Traffic costs calculation to use full cost information instead of just traffic
* Kentik Market Intelligence subscription, export, and APIs
* Maps ‘null’ to ‘None’
* Legacy KB links

Prevented:

* New MKP errors
* Devices export report from cutting off
* Traffic Costs from using the wrong ASN when ASN is in the description of another ASN
* Page navigation from showing up where it doesn’t

**Network Monitoring System (NMS)**

Fixed:

* Time established in BGP Neighbor metrics
* Inefficiency in NMS BGP query

Added:

* Ranger Discovery Service that proxies call to Ranger GRPC testConnections endpoint
* gNMI page that opens from devices to all for two gNMI walks, capabilities and subscribeOnce
* Link to policy from Metrics Explorer

Updated:

* SNMP status pill and added a tooltip on the Devices » Admin page to reflect the status of a new metric
* Flow monitoring templates
* Nulls to not be treated as nulls when averaging enums

**Synthetics**

Fixed:

* Synth Agent maintenance mode
* Configuration for private IP tests to block users from making invalid tests

Added:

* Warning for users when selecting ping and trace to global agents using ICMP or UDP and bidirectional option warning

Updated:

* Audit log
* Protected undefined config
* To show generic SynthUser when creating/updating tests while spoofed

## July

**AI / Insights**

* Added a “Test API Token” button to the AI Model Settings page, which verifies credentials
* Updated NMS metric and dimensions from the LLM context
* Triggered AI Gateway model refresh after updating AI models

**Alerting**

Fixed:

* Undefined labels property on policies page
* Various things for NMS policies and wizard adjustments
* Sort by start date on the BGP announcements page
* Errors in Alert Why baseline values
* MKP notification links
* Default serializer import

Added:

* Start manual mitigation button to alert actions
* Show platform status details for Adaptive Flowspec information (e.g., goal, traffic, rules) and support for announcements in BGP Announcements table
* Initial handling for Alerting Public API v6 with policies
* Source IP graph to the Alert Details view
* Guard for duplicates for alerting table entries

Updated:

* NMS Policy Form condition states error
* Suppressions to use multi-attribute filters
* Name from “standard” to “flow” for alert policies
* NMS Active Alerts widget title for clarity
* Device data with BGP status for DDoS configure pre-check
* Alert chart axis scales to be consistent with DE
* To show fallbacks for missing selections in MKP device selector
* To show seconds for alert start and end times
* Custom columns to support backfilling

**Cloud**

Fixed:

* GCP metadata fetch bugs
* Cloud Pathfinder Error Alert

Added:

* Additional route to bypass ACL for Transfluo
* Checkmarks to Azure cloud export edit screen for NSG and VNET properties of azure access check
* Logic and tests for the next version of the Cloud Export API
* Cloud Pathfinder to API Tester
* Null check and better error handling to AWS topology fetch

Updated:

* The alarm processing logic and filtered to only include critical and warning severity levels
* “cloud metadata” to “AWS Metadata”
* AWS Firewall UI and AWS Cloud WAN pathfinder support
* Cloud Export Config Wizard
* Error handling for AWS access check

**Core**

* Fixed:

  + `kdns` capability charts
  + Shared link cache and infinite logout loop
  + TrafficGrouping for missing key
  + Web hook custom headers
  + Currency Rate fetching
  + Interfaces with no IP to show in global search
  + Scrollbars for table with tabs in DE
  + View in DE
  + Legacy v5 queries with old device name and API v5
  + Saving companies with banners
  + Units check for old queries
  + Value in OTT widget to report the correct classification number
  + Required fields to display properly with asterisks
  + Issue where view query in DE was not updating in traffic costs
  + Showing Add Label for RBAC label selector
* Added:

  + APIv6 metrics to get visibility on v5 forwards
  + API routes for frequent pattern analysis and Journeys management
  + Missing MKP routes
  + Missing validations for API queries
  + Uninstall button for Universal Agent capabilities
  + New WebAuthn API routes
* Updated:

  + MKP RBAC and label filtering
  + Landing page for user personal preference page, user admin page, RBAC RoleSets
  + APIs for the landing page
  + To allow non-exact filter values in API to allow bidirectional `ktsubtype` filters in API
  + DateRange to be VPAT accessible, VPAT for menus, and Calendar VPAT and select focus
  + Traffic Costs enhancements
  + The performance of Device Details page
  + DE event controls
  + Search and filter for Configure Provider
  + Generator charts when Chart-level Group By Dimensions changes and allow saved filter usage
* Prevented

  + Double interface results in query for Traffic Costs by ASN
  + Flat-cost interfaces from spiking traffic cost values
* Enabled device RBAC for all users

**Network Monitoring System (NMS)**

* Added

  + Support for security appliance device class
  + Optional transceiver columns to device details > interfaces table
  + Transceiver charts to interface drawer
  + LastSeen and Usage to metrics dictionary
  + OID matching for device class
* Updated

  + Filtering available dimensions based on domain
  + Vendor NMS icons
  + Enterprise/model data for NMS Discovery
  + Alarm clear events to show in audit log
* Removed

  + Unused components
  + SNMP timeout
  + i\_device\_site\_name from syslog/trap capability sidebar queries
* Prevented error in NMS chart export and for unsupported queries

**Synthetics**

* Fixed:

  + Update Agent API Test
  + SoTi not displaying correct tests
  + Synth Test Results Map
  + Synth Network Explorer card
* Added:

  + Switch to ping/traceroute test settings to request collection of path MTU
  + A set to dedupe metric interfaces before sending query
  + Hover text to details column in Incident log UI
  + New option to subnav component to display an arrow on the last breadcrumb trail
* Updated:

  + Default health value to undefined in higher density tables to accurately assign dots the gray value seen in regular density
  + Synth Agent Maintenance Mode improvements
  + Agent management UI
  + Classification of global agents in agent model
  + Error message when unable to preview tests
  + Map dot markers and added test results legend
* Improved readability on tags
* Provided custom empty state for mesh panels with aggregated data

## June

**AI / Insights**

* Journeys:

  + Fixed Synth Test Results
  + Fixed Tracing implementation
  + Returned collection from findDevices
* Journeys API: Fixed DB insert error and AI permission setting check
* AI Settings: Updated UI
* Cause Analysis:

  + Updated wording and UI treatment
* Safe guard window.Notification if not supported
* Fixed incorrect and missing filter operator serialization mappings to correctly save the policy

**Alerting**

* Fixed:

  + The Protect landing page alert counts
  + Policy wizard "remove" button for `kevent` policies
  + Numeric input math related to key entry
* Added:

  + Support for header and collapsed wizard components
  + “Copy Status JSON” to mitigation Platform Status popup to support other platform-type statuses
  + New severity icon
* Updated:

  + Policy filter serializer
  + ackRequired to be derived from rule, not policy
  + Mitigation type to be derived from record, not associated method
  + NMS notifications to differentiate the name and label
  + Log manual alarm clear requests with alarm IDs
  + Wizard window sizing
* Supported skipValidationIf on complexArray fields
* Simplified `kevent` policy management and serialization
* Removed restriction on “All Devices” selection for tenant data partition in MKP

**Cloud**

* Added:

  + VNET peering logic for Azure connectivity report
* Updated:

  + Default sort order in Azure subscriptions filter to be by subscription name
  + OCI Map compartment selector
  + Pathfinder API

**Core**

* Fixed:

  + Graphical visual interface for tabbed chart
  + Credential Vault scrolling issue
  + Broken PopulatorFormDialog
  + LabelList overflows on Devices page
  + Cause Analysis tabs in Data Explorer
  + Interface metrics in Kentik map sidebar
  + Browser logging error
  + Error during device save
  + Empty columns in results table
  + Missing `href` from LinkButton
  + Performance issues on /core when rendering a large number of sites
  + Custom Dimension filter direction function representation
* Added:

  + User and deviceLabels APIv5 translation service
  + Default member and error handling for API create user by
  + Notify test channel `GET` call to audit logs as requested by security report
  + Validators to BYO DNS to prevent RFC1918 or public IP addresses degrading query performance
  + Sudo categories and ID display to Audit Logs
  + "Show Devices" popover to Metrics Plan card and "View All" link to Plan popover
  + More space for Cost Group name in the Cost table
  + IP port rule for valid IPv4 and moved the rules for `-dns` and `-dns_resolv` in Kproxy to this new rule
  + Search limits to interface lookup
  + Site market collection/model API for MKP
  + KMI Only flag to Company sudo to skip flow checks
  + Switch for intelligent classification
  + Site market lookup support for MKP
  + AWS Zone ID fields to dictionary
  + Floating bulk actions bar for devices and device interfaces
* Updated

  + Menu and Tabs VPAT accessibility
  + Query API to properly handle authentication error to specify that the user is not logged in (not a server issue)
  + “View in Data Explorer” button on the Devices » Traffic chart to be a right-clickable link
  + NMS IP addresses to be searched for on the devices page
  + logLevel options in Kproxy (replaced “Off” with “Disabled” and removed “Access” entirely)
  + Traffic Cost export
  + Universal Agent KProxy capability device linking
  + Sorting for Monthly Traffic Cost Snapshot History
  + Kproxy Flows In metric
  + Persisting Interface Classification results only if device is present
  + Name of device restricted API
  + Select “choosing” indicators
* Removed

  + Kubernetes from the Upcoming section of the Universal Agents page
  + Webinar from Login page
  + Period over Period selector from MKP and Dashboards
* Matched Interface Traffic & Capacity timeline with bandwidth level for OTT
* Prevented potential errors when saving traffic costs snapshot
* Defaulted Cause Analysis to On
* Allowed create/update of certain cost provider names
* Cached interface flow between fetches
* Converted ResultsTable underlying implementation

**Network Monitoring System (NMS)**

* Fixed:

  + Dictionary initialization fail on recon measurements
  + Active alert count
  + Device table queries
  + Metrics Explorer sidebar that was stuck in an invalid state
* Added:

  + In bit rate and out bit rate graphs to the sidebar of NMS interface
  + SNMP polling options in Add NMS Monitoring Template form
  + Class as a Filter and Group By options in the Devices table
* Updated:

  + Device class display on Device List page, Device Details page, and Device Configuration modal
  + Monitoring Template SNMP polling configuration defaults
  + Average and percent Metrics Explorer rollup aggregations
  + Devices table and ensured consistency with use of “Flow Device Type” and “Flow Type”
  + NMS alerts to show device alerts that don't contain device ID
  + NMS device classification
  + /api/ui/recon/metrics (listReconMeasurements) to use kmetrics-api
  + Format of NMS discovery models

**Synthetics**

* Fixed:

  + Agent filter malfunction
* Added:

  + Subtest rollup permissions
* Updated:

  + ‘Agent Management’ to ‘Agents’
  + Label lists on Synthetics pages to be in alphabetical order
  + Target column rendering with agent details when available

## May

**AI / Insights**

* Journeys:

  + Removed maxHeight on CC PromptResponse so that CC results do not overflow vertically and just scroll with the page
  + Fixed KB queries and references
  + Explicitly passed all attributes manually in every DB call
  + Added Investigations, Policy Forecasting, and UX clean-up
  + Fixed show all alert types
  + Fixed regression in messages for LLM
* Fixed bug with NMS result with no time series for Journeys API
* Added more AWS Dimensions for Data Explorer NLQ
* Enabled AI Gateway on all environments, making it the default
* Multiple AI Gateway updates
* Refactored settings to populate fields using data from AI Gateway API endpoint

**Alerting**

* Fixed

  + MKP update package
  + Issues with false positives when checking alerts for matching suppressions
  + Support for searching by tenant and dimension details to the alerting table
* Added:

  + Support for DateTime, Dashboards and hidden sections for wizard
  + A new piece that stores everything that’s common between all policy types with possibility for overrides on individual models
  + Capability to skip field validations for FormState
  + New alert severity colors and indicator
  + Ability to select type (Flowspec, RTBH, etc.) to filter by on the Mitigations table
  + `kmetrics` EVM notification test
  + User timezone to the alert table column, and related times in the drawer and detail page
  + Filter on mitigations table for manual control
  + Description getter for Policy Model
  + Read-only mode and URL by ID handling for mitigation methods and platforms
* Updated

  + Mitigation announcements table, collection, and model and reworked status indicators for mitigation platforms to show device statuses and link to the announcements table
  + Alerting constants
  + Mitigation drawer
* Allowed configuration of threshold conditions for enum-type metrics in NMS policies
* Enforced unique rule alias when creating policies
* Made sure policy Activation Frequency fits within the Evaluation Frequency
* Showed alert count on the device detail page
* Improved mitigation notification template
* Cleared mitigation table filters when creating a manual mitigation
* Disabled V5 alarm tests

**Cloud**

* Fixed:

  + 500 error when trying to load Azure map with no subscriptionTopologies
  + AWS cloud to ignore blackhole routes when generating route table summary for core networks
  + Azure no-metadata status bug
  + Public Clouds export to stay enabled but shows warning when access check fails
  + Missing Cloud Pathfinder summaries
  + AWS map error due to empty links array
* Added:

  + mexicocentral Azure region to dictionary
  + GCP and AWS regions in dictionary
* Updated:

  + Azure Subscription search
  + Error when direct connect does not have any attachments and added support for direct connect gateways connected to core network through TGW
* Prevented duplicate exporters in cloud config result and updated the title to show exported name
* Handled null Azure access check and implemented better icon status if location is bad
* Resolved text overlap for select all when adding custom dimensions for AWS and Azure tags
* Made S3 status checks in API results when relevant

**Core**

* Fixed:

  + Voluntary Product Accessibility Template (VPAT) blocker tabs and menu
  + customer port egress results
  + Capacity Plan Severity Filters
  + SSO saml default user\_level
  + Display of date on chart x-axis
  + Display of IP addresses in interface sidebar
  + AWS VPC and subnet name filters
  + Multiple lookups in aggregate series
  + Traffic Costs currency related calculations
  + Custom Hex Values in Bracket Options + ColorPicker Enhancements
  + Site select error
  + Device name change when there are empty saved filters
  + Handling of cancelled queries
  + Legacy redirects
  + Trending View populate job
  + Bug in Data Explorer and Dashboards with table visualizations, in which overlay rows are being incorrectly included in the Total Row calculations and count of Top X records
  + Performance issues in Safari with new login background
  + Name lookup for traffic cost snapshot
  + Interface APIv6
* Added:

  + Fallback for site\_id search
  + Plan APIv5 translation service and updated plan public API fields
  + Webinar login card
  + RBAC tests around getUniqueDeviceNames
  + SUDO tag to raw Plan FPS on licenses page
  + Process level handling for unhandled exceptions
  + Quick view selectors and extra context on OTT detail pages
  + Device vendor icons
  + More specific application name to OTT queries instead of ui-app
  + Missing query filter operators for Cloud Dimensions
  + Missing dimensions when saving customer port traffic costs snapshot
  + New label filtering component with and/or ability
* Updated:

  + Device Support request type and placeholder text in the support modal
  + To not require authorization when service visibility is private
  + UI for Cloud Pak and Synthetics Edit Plan windows
  + Device table after saving device configuration
  + User Level (legacy) changes to only reflect in RBAC Roles and Permissions if the user is not assigned any custom roles
* Prevented:

  + "Payload Too Large" errors when updating devices
  + Device dialog from erroring if flow IPs are not defined
  + Error on Archived NMS Device Details
* Interface Classification:

  + Fixed re-queueing of failed jobs
  + IC safety changes
  + Added performance improvements for Interface Classification
* Wrapped tabs as a flex instead of just a div
* Allowed apostrophes in site form address and added node sanitary validation to address sub-fields
* Removed sending IPs restriction for metrics plans
* Upgraded and tested latest select VPAT
* Allowed NMS panels in dashboard even if flow hasn't been received
* Various Traffic Cost export and miscellaneous bug fixes
* Supported exporting matrix to CSV
* Supported modal updates
* Showed page title on traffic costs for ASN in Dest Path and Customer Port
* Sent back appropriate SSO login error message to handle IDP login errors
* Re-enabled company-specific dictionary cache
* Redesigned Agent Capability with new config edit option for `kproxy` and OTT/kdns
* Disabled snapshot tracking to set the end date properly

**Network Monitoring System (NMS)**

* Fixed:

  + BGP neighbors table prefixes for sent/received
  + Device Details page and Metrics Explorer dashboards
  + UI in "Add Visualizations" dialog
  + Trim list of loading messages
  + Metrics Explorer 'run query' to execute query even if the query has not changed
  + Fixed disappearing charts on device details page
* Added:

  + Clarity to dimension filter summary display
  + SUDO tag to NMS and adjacent sudo-only actions
* Updated:

  + Sudo UX for auto updates in Metrics Explorer
  + Enterprise/model data for NMS discovery
  + NMS devices that do not have a monitoring template set to show the default “Everything” template in the device list page
  + Metrics Explorer to display mac address dimension column
  + Alert drawer layout
* Deleted Devices showing up in Availability Dashboard
* Improved NMS queries
* Shortened NMS interface name
* Allowed Flow and ICMP together
* Sorted NMS alarm dimensions and metrics when enriching EventViewModel
* Showed device name in device\_id column
* Ensured monitoring templates list is updated
* Ensured appropriate dimension filter operators are shown based on dimension type

**Synthetics**

* Fixed:

  + Path View Empty Targets
  + Multiple components of the Synth Map
  + Message for the fail to load incident log widget
* Added:

  + IBM Cloud Classic regions to cloud metadata
* Ensured:

  + freeCredits and paidCredits are numbers
  + `ksynth` agent select options get passed on
* Provided a subtest results link in th ereverse direction (target → agent) for bi-directional agent and mesh tests
* Increased default synthetics test alarm limit from 25 to 100
* Grouped multi-select component
* Included “lan” as a valid TLD for synth test validation

## April

**AI / Insights**

* Journeys API: Fixed error if no `kappa` and error on KB result
* Journeys UI: Fixed multiple UI components (missing feedback bar, missing heights for alerting and syn results, scroll function on alerting admin table, NMS alertMan application filter updates, hidden results in DE table visualization type, auto-closing panel, renaming a journey, filters not displaying in ME, etc.)
* Journeys: Updated delete and restore actions for Kentik AI
* Journeys: Updated AWS Zone ID few-shot examples to match dimension changes
* Journeys: Cache dimension enum query result

**Alerting**

* Fixed clone policy for regular threshold and copied from template flows
* Fixed MKP policies and tenant configuration
* Fixed group by primary dimension for NMS native alerts
* Fixed IPv6 wrap in alert table
* Fixed adding comments for alert on DDoS landing page
* Fixed broken Alerting backfill dimensions and covered with UTs
* Fixed alerting API test
* Fixed broken insight and NMS charts
* Updated alert triggered event details for more accuracy
* Updated notification channel association list
* Used policy metric order to identify baseline metric
* Used proper query property for KDE aggregate minutes
* Added extensions to core fields and filter components
* Added device site name to notifications EVM
* Added the capability for users to save a baseline-only condition policy without having to specify a static threshold condition metric value
* Added a refresh results button to the table to show last fetched time in tooltip
* Removed extra right padding on the nav and subnav
* Removed Last 24 Hour Top Dest IPs from DDOS landing page
* Removed “Total” dimension when constructing alert DE queries
* Removed unsupported flowspec fields
* Improved Alert CSV export native NMS
* Streamlined alert table filtering for exact alert ID search
* Validated device type selection against device availability when creating policies
* Showed time in either local or UTC format on the Alert debug chart
* Honored user timezone setting for suppression details and simplified start time validation
* Scope required condition rules only to NMS policy types
* Supported `kevent` policies and alerts
* Showed “Clear Alert” as an action in the alert table kebab
* Handled `kevent` null threshold filters
* Accounted for escalations in alert charts and explanations
* Used DE-required dimensions as Syslog and SNMP Trap policy defaults
* Serialized missing threshold filters as null
* Stripped commas from numeric policy filter values for standard alerting policies
* Handled null case for Top Keys alert chart threshold condition value
* Improved mitigations table and drawer, and filter alerts
* Renamed mitigation model and collection
* Supported navigating to mitigation by ID in URL
* Optimized single alert fetch

**Cloud**

* Fixed the AWS API Status Cloud Config section timing out
* Fixed reporting API Status results when scraping AWS topology
* Updated OCI DRG connection queries/filters
* Updated AWS map performance and map account selector
* Added support for security group rule with security group references in AWS connectivity checker
* Added extra null check when drg route table is not found and showed CalloutOutline error when route table not found in topology
* Added more logging to aid in debugging future AWS and OCI metadata fetch and API result caching
* Added Riyadh region for the OCI dropdown
* Added ‘Use OCI Timestamp (beta)’ toggle for OCI datasource in DE
* Optimized OCI scraper to prevent duplicate scrapes and created custom Kentik client configuration to prevent too many delays
* Disabled on-demand OCI API calls for config status
* Allowed for creation of cloud export that will not fetch metadata for AWS, GCP, and OCI, and fixed outdated KB links
* Removed processing route to service tags for Azure

**Core**

* Fixed subdomain URL parsing tenant for SSO login
* Fixed company trial expiration date display
* Fixed corner case where SSO lookup button is disabled
* Fixed possible redirect bug during 2FA
* Fixed login password functionality
* Fixed ASN details view error
* Fixed alignment of fields without helper text
* Fixed broken Virtual Column Filter for Prefix+LEN
* Fixed right filter fields that are subtype or appproto
* Fixed showing Account Team Card on license page
* Fixed unique device fetch with all datasource option selected
* Fixed error when ASN information doesn’t exist
* Fixed accessibility for login
* Fixed the Add New Saved View
* Updated SSO lookups to be taken directly from email address rather than TLD
* Updated UX for all login pages
* Updated style for Yubikey verification
* Updated KB links in the portal
* Added manual run option for report subscription
* Added pending tab to OTT
* Added a new API endpoint to fetch device by IP and/or Site
* Prevented possibilities of bad redirects during login
* Prevented the verify DNS dialog from erroring on focus
* Prevented interface table from showing duplicate results
* Prevented crash with invalid view ID
* Prevented device labels from being removed by bulk labels edit
* Improved log message by using a consistent unit (milliseconds) and displaying the unit
* Correctly applied default permissions to company details
* Truncated dimensions separately in tooltips
* Made minimal tags consistent across the platform
* Linked sentry issue to error report
* Made OTT bar label readable
* Included all relevant fields in update device response
* Adjusted device column widths due to new tags
* Used new OTT-classification column
* Only showed “Add Cost Group’ button after provider has been saved
* Reworked login background animation for improved CPU-only performance
* Moved Credential Vault API to restricted internal service
* Released Traffic Costs

**Network Monitoring System (NMS)**

* Fixed the charts that are rendered with NMS alerts that occur when the polling cycle is less than the default
* Fixed filters for custom up/down charts
* Fixed loading of saved views
* Made default condition group connector ‘all’ instead of ‘any’ for Add NMS Alert Policy
* Restricted dimension filter operators based on type of selected dimension
* Prevented error when device availability fails
* Ensured device list and detail views show same CPU utilization
* Ensured device details CPU sparkline is identical to list page and list drawer
* Allowed cloud dimensions with spaces
* Expanded SNMPv3 auth options

**Synthetics**

* Fixed Synth subtest chart hover and avg metrics to fetch results only for a given agent/target pairing (if available)
* Fixed Incident Log widget configure panel to not show by default
* Fixed Alarm Table DNS Test links
* Fixed mesh toolbar (ensured checkboxes in the metric selector are selectable, displayed down caret icon in agent selectors, started agent numbering at 1, ensured consistent ordering of metrics display)
* Fixed the dot colors for high/max density aggregate mesh
* Fixed Synth IPv6 agent support validator
* Fixed Synth Agent IP search
* Fixed selected timestamp synth test export
* Fixed alarm timeline selected state
* Fixed Synth test public shared defined in a date range calculated by look-back seconds
* Updated performance to trace query and allowed filtering traces by target
* Removed agent IP mapping
* Fixed Synth Trace Hop Count by enabling filtering the trace data by agent ID and target ID
* Addressed Synth path panel feedback and showed a tooltip indicating selection restrictions and hint for filtering by selected items
* Included paused tests in agent details response
* Unlinked edit link text from add test

## March

**AI / Insights**

* DE NLQ: Fixed IP range filtering
* DE NLQ: Updated filtering on inexact device name
* DE NLQ: Added flow tags to filtering
* Journeys: Added more instructions for Journeys to only respond to relevant parts of a question / user queries with multiple questions
* Journeys: Updated v6 API
* Journeys: Updated UX and UI, and added live synth NLQ

**Alerting**

* Fixed an issue to now show `kevent` records in the alerting table
* Fixed the Policies page performance for customers with high device counts
* Fixed a typo in the Start Manual Mitigation dialog error message
* Fixed up/down chart safety boundaries when constructing query
* Fixed bulk add notification channel button for NMS policies
* Syslog: Extended core notification translator to support `kevent` alerts
* Syslog: Implemented a translator function for turning NMS-shaped form values into event-shaped policy values
* Syslog: Charts and filter transforms for `kevent` alerts
* Syslog: Fixed issues with existing Query-to-Policy (Q2P), and extended methods to prepare for syslog Q2P
* Updated to show full alert count for the past 7 days, and showed a link to view alerts in the table
* Centralized managing company policy application requests and permissions
* Removed 1px border from the drawer drop shadow
* Handled clear override for `kmetrics` policies

**Cloud**

* Fixed AWS Cloud Map Transit Gateway Route Table sidewar widget to include an accurate number of route tables in the summary, group-by table ID, and display table name in the group header. Table ID and name will show as a group header with an empty row containing text “No Routes Found”.
* Fixed bug in Azure entity explorer and fixed link from performance monitor to cloud export config wizard
* Fixed direct connect gateways bug in AWS metadata fetch
* Added AWS Firewalls and Firewall Policies into topology
* Added AWS account summary generation, added menu in the search map section, and updated topology refresh to load only the selected account
* Metadata boxes do not allow unchecking for non-Azure cloud exports
* Improved AWS account selector
* Updated cloud AWS and OCI map
* Changed references to connectivity checker to be Cloud Pathfinder

**Core**

* Deprecated TSDB
* Fixed Query URL
* Fixed error thrown during device details overview export
* Fixed the “none” option for network boundary
* Fixed nested query saved filter
* Fixed phrasing on column selector
* Fixed sudo lookup
* Fixed Regression from APIv5 migration
* Fixed device details redirect
* Updated Devices Export CSV to include NMS attributes
* Updated system user provisioning, and added `kagent` and AI users
* Updated running queries to ignore invalid saved filters
* Updated KB links in Device Settings dialog
* Updated APIv5 translation service
* Added charts for syslog and traps capabilities
* Added Copy to the Clipboard button for Audit Log payloads
* Added OTT Classification Type column
* Added V3 redirects and fixed company delete
* Cached hash results for KentikAgentLink to prevent unnecessary fetches
* Applied CPU, memory, temperature, and traffic data values for NMS metrics
* Verified interface sidebar numbers match the table
* Eliminated partial 2FA bypass method
* Made new APIv6 routes available to API tester
* Ensured Add Device buttons work on all device list tabs
* Set the default filter value for ‘Show Archived Devices” to true
* Allowed query editor even when there’s no flow
* Showed Devices correct ST status on load
* Corrected typos on Configure New Provider page
* Removed Return to V3 links
* Prevented what field from being overwritten when undefined in the update settings payload

**Network Monitoring System (NMS)**

* Corrected the display status for ICMP-only devices
* Updated the config and data modeling structure for the current NMS policy form
* Fixed wizard arrow navigation between sections
* Fixed NMS policy form validation
* Added Kentik Agent as a target type for NMS alerts
* Added NMS sudo options for filtering alerts/policies
* Formatted metrics in the policy drawer
* Updated the NMS policy form to only show entity types that the UI knows how to support
* Updated policy table dimensions for custom target NMS policies
* Fixed the refresh of metrics dashboards
* Corrected policy state display
* Adjusted layout of NMS policy conditions display
* Used average instead of last to sort device details interfaces queries
* Added support for nested condition groups in NMS policy form
* Added help text to NMS policy form
* Fixed type errors in MetricsResult
* Returned `kmetrics` device properties with device detail requests
* Fixed alerting pages
* Fixed manual clear toggle for NMS policies
* Ensured all series are within y-axis bounds of device details charts
* Added active state styling to condition groups when clicked
* Used consistent NMS dimension display logic across components
* Updated the default Universal Agent Problem policy to an NMS-native policy
* Limited `kagent` docker logging on general principles
* Fixed active condition group select action for portal drop-downs

**Synthetics**

* Synth Alarm first trigger updates include display and functionality of test results like the incident log and timeline by adding new features, renaming existing ones, and addressing how data is shown and linked
* Fixed various issues related to alarms timeline accuracy, particularly in subtest results
* Fixed Synth Agent activation with Alerting
* Fixed waterfall view height to expand with parent’s flex

## February

**AI / Insights**

* Journeys NMS: Allowed grouped NLQ filters, port metric filters to standard UI, and various re-factors
* Journeys: Added AWS CCC subnet name and instance name lookup
* Journeys: Added Syn Test results and fixed alert prompt delete
* Journeys: Added ability to undo delete
* Journeys: Added missing function export
* DE NLQ: Added unique & total metrics
* DE NLQ: Updated filtering on IPs individually in a range

**Alerting**

* Added new event application type for `kevent` type policies
* Added "is greater than" option to policy condition operators
* Mapped `kevent` alerts to the Core notification translator
* Removed ‘require’ marker from the Custom Webhook OAuth fields
* Deprecated the Alert Policy Minimum Traffic Field
* Changed Flowspec rate limit label back to bytes and added a tooltip warning users that their device may interpret this value as bits
* Forced refetching notifications when loading tenants
* Fixed the condition error that appears on policies with Top Key conditions
* Fixed the short-circuit infinite load when using local search
* Fixed MKP Alerting fetching and filtering functionality
* Fixed tests and serialization for policy silent mode
* Fixed a display bug where dimensions overlapped with controls, and added an ellipsis if the dimension is too long

**Cloud**

* Updated the frequency of fetching AWS metadata for Principal to every hour
* Updated documentation link in Empty Traffic Overview
* Updated UI to add AWS Availability Zone ID and AZ name
* Updated query times to have a 5-minute floor
* Added an ENI cloud icon and display in the connectivity checker report
* Added a confirmation modal before deleting a cloud export
* Dashboards that use all data sources for queries no longer require cloud exports
* Refactored AWS scrape and API call reporting
* Established null filters for TransitGateway and TransitGateway Attachment
* Refreshed access token when previous one expires
* Enlarged Data Explorer filter options modal viewport
* Azure: Removed NSG logs and added VNet logs support
* Fixed UI bugs in the Azure cloud export edit experience

**Core**

* Saved and planned filter APIv6
* Devices List Page - Agents - Combined lookups, showed filter on initial load, and allowed grouping
* Devices List Page - URL parameter for tab
* Prevented the Devices page from showing no devices by default
* Prevented no columns if linked directly to Devices custom view mode
* Switched to library for extracting SSO domain
* Modified the 'Show Archived Devices' filter so it does not hide non-archived devices
* Unified Devices List: Stored last view mode to user and filter sidebar group open states to hash
* Updated NMS status details in telemetry
* Allowed monitoring template updates even if the device has multiple SNMP collection methods
* Established the correct KMR for auto provision user via SSO login
* Updated sentry.io DSN fetch request
* Removed the leave empty IDP certificate description
* Significant updates to Syslog collect functionality, alerts, constraints, and UI display (new page for NSM Logs and new widget for observation deck)
* Updated device landing page and details page
* Core: Added metrics suppression on top of existing dimension and filter suppression
* Increased the default limit of the metric device collection
* Showed correct actions column in devices custom columns
* Used the same unknown icon in the devices listing page as the details page when relevant
* Loaded the kproxyAgents collection when visiting the Telemetry tab
* Reverted sentry for appwrapper router
* Added Bulk Classify and PDB IX Assignment to Interfaces
* Added a bitrate/utilization switch for interface charts
* Added site market column/grouping to Unified Devices
* Added node-side validations for last super admin modifications
* Added space between lower sublayer interfaces
* Added sentry.io integration to react router paths and error capture
* Added manage menu to the unified devices page and added ellipsis to the credentials column
* Added id and cdate to the device details page
* Added NMS device parameters to list devices call
* Added NMS device parameters to the main device call
* Added ability to bulk update device monitoring templates
* Added kproxy 7.50
* Added custom dimensions API to v6
* Fixed the logic for recently added parameter for /api/internal/customdimensions
* Fixed filters between interfaces and connection on device details
* Fixed the locked up devices page due to agent lookup
* Fixed the interface filter sidebar for handling nulls
* Fixed loading of metrics device on the details page
* Fixed device filter and made other improvements on the interfaces table
* Fixed entity name in tab selector
* Fixed PDF exporting certain device pages
* Fixed fetching query parameter
* Fixed incorrect filtered devices link that led to incorrect filter hashes
* Fixed the “Auto-Greek prefixing for Traffic-In/Traffic-Out columns in the interface page” in work log
* Fixed chart colors profile settings
* Traffic Costs links now only show where it is applicable
* Fixed label suggestions to automatically populate for RBAC-restricted users creating new devices
* Added Nokia SAP Dimensions
* V5 Custom Dimension Translator

**Network Monitoring System (NMS)**

* Enabled NMS SNMP Walks by default
* Alert Manager rule for the NMS policy is now always enabled
* Fixed overlapping text in the BGP search bar with peer AS filter
* Fixed error on NMS query widgets with sparklines
* Updated enterprise/model data for NMS discovery
* Switched the default policy type selection from NMS to Threshold
* Added monitoring template type to audit log
* Adjusted threshold condition value label in NMS policy form
* Autofilled target and condition measurements in NMS policy form
* Linked to active NMS policy alerts from policy form and drawer
* Fixed NMS policy validation

**Synthetics**

* Updated agent throughput

  + Updated Agent-to-Agent test creation and Agent install modal TCP (allowed to be 0)
* Added protection against empty subtest trend results to eliminate fatal errors
* Added a switch to the alerting section of test config wizard that allows defining targeted notification conjunctions based on test rule ID and critical severity
* Sorted mesh rows and columns by alphabetical order instead of by agent ID
* Fixed Path View latency threshold logic to ensure a consistent value is used when determining latency threshold violations
* Fixed alarm timeline show health button
* Fixed Path View success rate and link popover metrics

## January

**AI / Insights**

* Journeys: Added query and NMS results to context
* Journeys, new models: gpt-4o-11-20, Gemini 2.0 Flash, Codestral, Mistral Large
* Updated Journeys UI
* NMS NLQ: Ensured that if a filter is used, that metric is also included in the main query

**Alerting**

* Recon 1535/BGP display in alert table
* Performance improvements for alert detail requests
* Fixed "invalid date" on policy suppression date picker
* Fixed silent mode date serialization
* Fixed Alerting Overview loading bug
* Fixed metrics up/down charts for NMS component entity alerts
* Fixed BGP alert dimension labels
* Fixed alerting table infinite scroll

**Cloud**

* AWS configuration status reporting missing API calls
* AWS: Added new version and API support for cloud export API to allow/block list updates
* Removed “beta” from connectivity checker page and menu, and added other copy changes
* Improved AWS Metrics Sidebar Link
* Fixed cloud export wiz unsampled bug
* Fixed bug in Availability Zone metadata fetch

**Core**

* Applied device RBAC to subscriptions
* Increased device name max length to 256
* Enhanced highcharts tooltips
* PR Label Validation
* Required IDP Certification and added Pendo Account information
* Added EdgeNext and Netskrt CDNs
* Added Contains and Regex to Cloud Service Filter
* Added kproxy 7.49 commit hash
* Added SSO URL back to the login page
* Added node-jobs service to regular deployment workflow
* Removed install button for agents not yet authorized
* Removed V5 API Tester Link
* API schemas auto-generate typescript definitions
* Updated certification location from docker volume mount
* Trimmed IDP certification to pass validation
* Various Traffic Costs improvements and fixes
* Prevented create site error during onboard when missing address
* Initial pass at storage referrer on SSO login
* Exports: Allowed png report, fix NMS report
* DE: Made device name filter case insensitive
* Global Search: Pressing ENTER now goes to the correct search result
* Traffic Costs: handled record missing better
* Prevented rows in cost csv with unknown sites from being excluded
* Fixed empty UDR columns, DNS traffic nav link, and various tooltip issues
* Fixed slow rendering of saved filters edit dialog
* Audit Log: Fixed user column for generic events
* Fixed clearPlotBandsAndZones highcharts issue
* Fixed Sample Rate in device sidebar
* Fixed saved filter showing loading spinner
* Fixed PDF exports
* Fixed node-job-update-traffic-cost-history
* Fixed Connectivity Costs logic when ensuring each timestamp slice is accounted for

**Network Monitoring System (NMS)**

* Made NMS Alerting on the default for all non-on-prems
* Preserved notifications when enabling/disabling NMS policies
* Fixed association of warning/severe notifications with NMS policies
* Fixed charts and ME for BGP entity NMS alerts
* Fixed NMS notifications for warning and severe levels

**Synthetics**

* Aligned test type names between the “Add Test” page and the TCC filters
* Global agent:target pair no longer displays traffic charts
* Updated Synth threshold validation
* Improved Synth scheduled tests
* Updated http latency in test results to not display zero (0)
* Allowed local domains in Synth URL Fields
* Ensured Synth plan is active when fetching Synth plan
* Ensured reciprocal agent-to-agent credits calculation
* Shortened timeframe for clicking from the incident log or notification details
* Fixed Synth Library Health
* Fixed grace period/test frequency validation for preview tests
* Fixed Synth Library Health
* Fixed Synth Path View agent filter
