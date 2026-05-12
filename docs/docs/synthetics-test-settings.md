# Synthetics Test Settings

Source: https://kb.kentik.com/docs/synthetics-test-settings

---

This article covers the settings for **Synthetics** tests in the Kentik portal.

> **Notes:**
>
> * For information on viewing tests, see **Test Control Center**.
> * For information about the test types you can manage in Test Control Center, see **Synthetics Test Types**.
> * For information about how test status (Healthy, Warning, Critical) is derived, see **Synthetics Test Status**.
> * For information about test credit consumption, see **Synthetics Test Credits**.

## About Test Settings

A Kentik synthetic test is a collection of subtests that generally involve pings and traceroutes from a set of agents (global or private) toward a target. The test type is typically determined by the target’s nature, falling into one of the categories described in **Synthetics Test Types**.

> **TIP**: When editing existing tests, there are additional **Test Management Controls**available.

#### Test Settings Access

Test settings are accessed in the following contexts:

* **Add a Test**: To configure settings for a new test:

  + From **Test Control Center**, click **Add Test** to open the **Add Test Page**.
  + Click the test type to open a **Test Settings Page** for the new test.
* **Copy a Test**: To duplicate an existing test and modify its settings:

  + From **Test Control Center**, find the test you wish to copy in the **Tests List**.
  + Click the **Copy** icon at the right to open a **Test Settings Page** for the duplicated test.
* **Edit a Test**: To modify a test’s settings, do one of the following:

  + From **Test Control Center**, click the **Edit** icon for the test you’d like to modify.
  + On the **Test Details Page**of the test you’d like to modify, click **Edit Test** in the subnav.

#### Credits Balance Notification

If your organization's test credits balance (see **About Test Credits**) is below zero, a notification appears on the **Test Control Center**page and any **Add Test Page** or **Test Settings Page**.

To add more credits, click **Submit a Request** to open the Contact Support popup (see **Support Request**), pre-populated with a request for a quote for additional credits (see **Adding Test Credits**).

## Add Test Page

The Add Test page enables you to create a new test by choosing from the list of **Synthetics Test Types**, organized into two columns:

* **Category**: The category name, description, and a link to find out how the tests in this category work.
* **Test Type**: Within each category, one or more test types are shown with a name and description.

  + Click a tile to start a new test of that type  Test Settings page.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(529).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A13Z&se=2026-05-12T10%3A15%3A13Z&sr=c&sp=r&sig=%2Bje1%2B89jPBifDyV%2BHMQ8PpAvTYctzC02eei4YuqZfmQ%3D)

*Choose the type of new test that you'd like to add.*

At the top of the page, you’ll see four tabs just below the title:

* **All Tests**: This tab (default) lists all of the test types that you can create for your organization (see **Synthetics Test Types**).
* **Routing**: This tab lists all test types available in the Routing category (see **Routing Tests**).
* **Network**: This tab lists all test types available in the Network category (see **Network Tests**), which includes the following subcategories:

  + **Agent-to-Agent**: Agent-to-Agent and Network Mesh tests
  + **Agent-to-Server**: Server IP Addresses, Server Hostname, and Network Grid tests
  + **Autonomous Tests**: ASN, CDN, Country, Region, and City tests

    > **Note**: If your system does not yet include flow data, Autonomous Tests will be unavailable (grayed out) and you’ll see a notification that includes a **Get Started** button. To start the setup for up flow data, click the button, which will take you to the Which Data setup screen (step 4 in the setup tasks described in **Initial Setup Login**).
* **Application**: This tab lists all test types available in the Application category, which includes the following subcategories:

  + **DNS**: DNS Server Monitor and DNS Server Grid tests
  + **HTTP**: HTTP(S) or API, Page Load, and Transaction tests

> **Note:** As described in **Settings by Test Type**, test settings vary depending on the type of the test. - When your organization has no remaining **Synthetics Test Credits** the page will display a **Credits Balance Notification**.

## Test Settings Page

The Test Settings page (accessed as described in **Test Settings Access**) enables you to configure a new test or to modify an existing test or a duplicated test. The settings (e.g., name, frequency, notifications, etc.) are mostly the same whether you are adding or modifying a test. Controls for the settings are distributed across a set of tabs that you navigate via the sidebar at the left of the page. The page also allows you to initiate a Test Preview (see **Test Preview Overview**) to validate your settings and see the impact of this test on your test credit consumption.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(530).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A13Z&se=2026-05-12T10%3A15%3A13Z&sr=c&sp=r&sig=%2Bje1%2B89jPBifDyV%2BHMQ8PpAvTYctzC02eei4YuqZfmQ%3D)

*Configure the specific settings of an individual test on the Test Settings page.*

#### Test Settings Layout

A Test Settings page includes the following main areas:

* **Title and controls**: General controls for the test (see **Test Management Controls**).
* **Sidebar**: A list of tabs containing required/optional settings and a credit consumption estimator (see **Test Settings Sidebar**).
* **Credits Balance Notification**: Displays when your organization has no remaining **Synthetics Test Credits**.
* **Tabs**: The settings tabs for the test (see **Test Settings Tabs**).

### Test Management Controls

The following test management and control elements are distributed across the top of the Test Settings page (under the subnav):

* **Test Type**: The type of test (see **Synthetics Test Types**).
* **Info** (icon): Click to get an explanation of how this test type works (when available).
* **Cancel** (button): Click to exit without saving changes (requires confirmation).
* **Pause/Resume** (buttons, existing tests only): Pauses or resumes testing related to this test.
* **Preview** (button): Opens the **Test Preview Page** for the test, or if editing a BGP test, the **BGP Route Viewer** (see **BGP Route Viewer**). Button only active when current settings are sufficient to create a test.

  > **Notes**:
  >
  > + Once the Test Preview page is opened, the **Preview** button on the Test Settings page will be greyed out. To run the preview again, do so from the Test Preview page.
  > + If the test type is BGP Monitor, the button opens the **BGP Route Viewer**.
  > + This button is not currently available for Autonomous tests.
* **Create Test** (new tests only): A button that saves the configuration and starts the test, which will be added to the **Tests List** on the TCC landing page.
* **Save** (existing tests only): A button that saves changes to the test. The button is active only when the current settings are sufficient to create a test.
* **Delete** (existing tests only): A button that opens a confirmation popup in which you can permanently remove the test from your organization's collection of tests.

### Test Settings Sidebar

Use the left sidebar of the Test Settings page primarily to navigate tabs, each having a group of settings for the test:![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(531).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A13Z&se=2026-05-12T10%3A15%3A13Z&sr=c&sp=r&sig=%2Bje1%2B89jPBifDyV%2BHMQ8PpAvTYctzC02eei4YuqZfmQ%3D)

* **Settings tabs list**: A list of the **Test Settings Tabs**, each of which corresponds to one group of settings that are available to be configured for this test. The number of tabs in the list (up to nine) will vary depending on the type of the test (see **Settings by Test Type**). Tabs with required settings are listed first, followed by tabs with optional settings. Click on a tab name to display the group of settings in the corresponding tab.
* **Status indicators**: Indicators that show the status of the settings group on each tab (see **Settings Status Indicators**).
* **Monthly Credit Usage Estimate**: A card at the bottom of the sidebar that shows the impact of the test, as currently configured, on your organization’s consumption of credits available for the current month (see **About Test Credits**). The card includes:![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(532).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A13Z&se=2026-05-12T10%3A15%3A13Z&sr=c&sp=r&sig=%2Bje1%2B89jPBifDyV%2BHMQ8PpAvTYctzC02eei4YuqZfmQ%3D)

  + Number of credits the test will use in the current month.
  + Credits that will remain available in the current month if this test is active.
  + Bar graph illustrating the above.
  + How many credits the test uses per minute.
  + Number of targets and agents used for the test.

    > **Note**: On an Edit Test page (but not when adding a test), you can see a breakdown of the test’s points usage by clicking **credit(s) per min** to open the **Credit Consumption Dialog**.

#### Settings by Test Type

The table below shows which Test Settings tabs are available for each of the test types.

| Section | BGP Tests | Agent-to-Agent | Agent-to-Server | Autonomous Tests | DNS Tests | HTTP Tests |
| --- | --- | --- | --- | --- | --- | --- |
| Test Information | Yes | Yes | Yes | Yes | Yes | Yes |
| Target and Agents | No | Yes | Yes | Yes | Yes | Yes |
| BGP Monitoring | Yes | No | No | No | No | Yes (except Transaction) |
| Flow-Based Targeting | No | No | No | Yes | No | No |
| HTTP | No | No | No | No | No | Yes |
| Ping and Traceroute | No | Yes | Yes | Yes | Yes | Yes (except Transaction) |
| DNS | No | No | No | No | Yes | No |
| Health | Yes | Yes | Yes | Yes | Yes | Yes |
| Alerting and Notifications | Yes | Yes | Yes | Yes | Yes | Yes |

#### Credit Consumption Dialog

The Credit Consumption dialog is available when you are editing an existing test. The dialog, which opens from the **credit(s) per minute** estimate in the **Monthly Credit Usage Estimate** card in the **Test Settings Sidebar**, contains the following elements:

* ![Overview of credit consumption for HTTP and network testing with associated costs.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/STS-credit-consumption-dialog.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A13Z&se=2026-05-12T10%3A15%3A13Z&sr=c&sp=r&sig=%2Bje1%2B89jPBifDyV%2BHMQ8PpAvTYctzC02eei4YuqZfmQ%3D)**Close**: An **X** in the upper right corner that closes the dialog.
* **Summary**: The number of agents and targets selected for the test.
* **Breakdown**: The types of subtests and the number of credits consumed for each.
* **Total**: The total number of credits consumed per minute for the test.

### Settings Status Indicators

The indicators listed in the table below, which appear to the left of the tabs in the sidebar's tabs list, are used to show the status of the settings on each tab.

| Color | Icon | Description |
| --- | --- | --- |
| Gray | Circle (outline) | The tab has not yet been viewed or modified. The Save and Create Test buttons are disabled when this indicator is shown for any tab. |
| Gray | White checkmark on disk | The tab's settings are complete:   * If this is a new test, the settings haven't been modified from their default values. * If this is an existing test, the settings either use default values or have been previously modified and saved. |
| Gray | Circle (outline) with diagonal strikethrough | The tab is disabled. Enable it with the switch at the upper right on the tab itself. |
| Green | Half-circle (outline) | The tab's settings are only partially complete. The Save and Create Test buttons are disabled when this indicator is shown for any tab. |
| Green | White checkmark on disk | Settings have been modified since the test was last saved, and all required settings on this tab are complete. |
| Red | Octagon with white X | The tab's settings are incomplete or contain errors. A red lozenge to the right of the section name states the number of errors. The Save and Create Test buttons are disabled when this indicator is shown for any tab. |

### Settings Tabs

The settings in the settings tabs appear in the page’s main display area to the right of the sidebar, below the **Test Management Controls**. Each tab contains the following UI elements:

* **Title**: The name of the group of settings, which corresponds to the tab you’ve selected from the **Test Settings Sidebar**.
* **On/Off**: A switch at the upper right that enables/disables the settings on the tab; present only in the situations listed in **Enable or Disable Settings**.
* **Previous/Next**: Navigation buttons at the upper and lower right of the tab that change the displayed tab to the one before or after the current tab.

  > **Note**: The **Next** button will be disabled until you've completed the required settings on the current tab.
* **Tab Settings**: The test settings available for the current tab (see **Test Settings Tabs**).

#### Enable or Disable Settings

The table below shows tabs whose test settings may be disabled depending on the test type:

| Settings tab | Test type |
| --- | --- |
| Alerting and Notifications | All tests |
| BGP Monitoring | HTTP(S) or API, Page Load |
| Ping and Traceroute | DNS Server Monitor, DNS Server Grid, HTTP(S) or API, Page Load |

For tabs that can be disabled/enabled:

* When a tab is enabled, its settings are fully editable and will be applied when the test is run.
* When a tab is disabled, its settings are grayed out, they can’t be edited, and they won’t be applied.
* The enabled/disabled state determines the status indicated for the tab (see **Settings Status Indicators**).

### Required and Optional Settings

For each test type, some settings are required and others are optional, as described here.

#### Required Test Settings

The table below shows the available test types and the required settings for tests in each of the categories described in **Synthetics Test Types**. Every test requires that you enter a name as well as the settings listed below.

| Test Category | Test Sub-Category | Test Type | Description | Required settings |
| --- | --- | --- | --- | --- |
| Routing | BGP | BGP Monitoring | Track whether specified prefixes originate from allowed ASNs. | Prefixes to monitor, ASNs |
| Network | Agent-to-Agent | Agent-to-Agent | Measure, ping, latency, and packet loss to a single target agent. | Test frequency, target agent, agent to test from |
| Network | Agent-to-Agent | Network Mesh | Measure, ping, latency, and packet loss for a grid of multiple agents. | Test frequency, agents (minimum 2) |
| Network | Agent-to-Server | Server IP Addresses | Measure agent connectivity to one or more target IP addresses. | Test frequency, agent(s), IP address(es) to target |
| Network | Agent-to-Server | Server Hostname | Measure agent connectivity to a single target hostname. | Test frequency, agent(s), Target hostname |
| Network | Agent-to-Server | Network Grid | Measure agent connectivity to multiple target IP addresses. | Test frequency, agent(s), IP address(es) to target |
| Network | Autonomous Tests | ASN, CDN, Country, Region, City | Measure agent connectivity and metrics to high-usage ASNs, CDNs, Countries, Regions, and Cities. | Test frequency, one or more ASN, CDN, Country, Region, or City to target; agent(s) to test from |
| Application | DNS | DNS Server Monitor | Measure responses for a DNS server and specified hostname. | DNS lookup frequency, DNS server IP(s) to query, hostname to look up, agents to test from |
| Application | DNS | DNS Server Grid | Measure responses for multiple DNS servers for a hostname. | DNS lookup frequency, DNS server IP(s) to query, hostname to look up, agents to test from |
| Application | HTTP | HTTP(s) or API | Measure a specific web server’s response to a HTTP request. | Method, URL or FQDN, agents to test from; if BGP enabled, prefixes, ASNs |
| Application | HTTP | Page Load | Measure performance metrics for a full browser page load. | URL or IP address  to target, agents to test from; if BGP enabled, prefixes, ASNs |
| Application | HTTP | Transaction | Test a web page with a series of actions driven by a custom script. | Agents to test from, Puppeteer script |

> **Note**: The test types in each of the categories listed above are also listed on the **Add Test Page**.

#### Optional Test Settings

The tabs listed under **Optional Settings** in the **Test Settings Sidebar** enable you to customize settings that are otherwise set by default, giving you finer control over the details of how a test is conducted. Depending on test type, the following tabs may be found under Optional Settings:

* **Flow-Based Targeting**: Available only for Autonomous tests (see **Flow-Based Targeting**).
* **HTTP Settings**: Available only for HTTP or API tests (see **HTTP Settings**).
* **Ping and Traceroute**: Available for all tests, except BGP Monitor and Transaction tests (see **Ping and Traceroute**).
* **DNS Settings**: Available only for DNS Server Monitor and DNS Server Grid tests (see **DNS Settings**).
* **Health Settings**: Available for all tests (see **Health Settings**).
* **Alerting and Notifications**: Available for all tests (see **Alerting and Notifications**).

> **Note:** The settings on each of the above tabs vary depending on test type.

### Select Agents Dialog

Open the Select Agents dialog using buttons on the Target and Agents tab of a Test Settings page. The dialog contains a list from which to choose one or more agents to be used for a test. The dialog varies slightly as follows:

* **Single-select version**: In an Agent-to-Agent test, the target agent is set with the dialog that opens from a **Choose Agent** button. Radio buttons to the left of each available agent enable selection of a single agent.
* **Multi-select version**: In most test types, the agents to test from are selected from the dialog that opens from a **Choose Agent(s)** button (if no agents are already chosen) or an **Edit Agents** button (if agents have previously been selected). Agents are selected with checkboxes, which operate only on agents that are shown based on current filter settings:

  + A checkbox at the left of the list's column headings enables you to select/unselect all shown agents.
  + Checkboxes in the heading row for each agent category (Private Network Agents, Private App Agents, Global Network Agents, and Global App Agents) enable you to select/unselect all shown agents in one or more categories.
  + Checkboxes to the left of each available agent enable you to select/unselect one or more agents individually.

![Selection interface showing various network agents and their locations for configuration.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/STS-select-agents-dialog.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A13Z&se=2026-05-12T10%3A15%3A13Z&sr=c&sp=r&sig=%2Bje1%2B89jPBifDyV%2BHMQ8PpAvTYctzC02eei4YuqZfmQ%3D)

*Use the Select agents dialog to choose agents from which to test.*

#### Select Agents UI

The Select Agents dialog includes the following fields and controls:

* **Filter IP Version**: Displays a lozenge for each selected IP version and filters the list to agents of those versions.

  + **To select a version**: Click the field and choose **v4 + v6** or **v4 only** from the dropdown.
  + **To remove a version**: Click the **X** in the version’s lozenge.
* **Filter agents by label**: Displays a lozenge for each selected label and filters the list to agents with those labels.

  + **To select a label**: Click the field and choose the label from the dropdown list.
  + **To remove a label**: Click the **X** at the right of that label's lozenge.
* **Search**: Enter text to filter the list of agents by matching name or location.
* **Close**: An **X** that exits the dialog without saving changes.
* **Selections pane** (multi-select version only): Lists the agents selected from the Agents list. Each agent is shown in a card with its name, location, assigned labels, and an **X** to deselect the agent.

  > **Note**: A red icon beside an agent name indicates an error (hover over the icon for more information).
* **Select All** (multi-select version only): A checkbox to select all agents in the list (excluding agents hidden by filtering). Individual checkboxes can still be checked/unchecked even when Select All is checked.
* **Agents List**: The list of agents available for this test. If the test allows selection of both Private and Global agents, the list will include a heading row for each of four collapsible categories (Private Network Agents, Private App Agents, Global Network Agents, and Global App Agents). The list includes the following columns:

  + **Select**: A radio button (single-select) or checkbox (multi-select) at the left of each row that selects individual agents.
  + **Agent**: The name of the agent.
  + **Location**: The name of the site or cloud region in which an agent is deployed.
  + **IP Version**: The IP version of the agent (v4 + v6 or v4 only).
* **Cancel**: Click to exit the dialog without saving changes.
* **Save**: Click to save changes to the test’s selected agents and exit.

> **Note**: When Global agents are selected for Agent-to-Agent or Network Mesh tests, the Protocol setting on the **Ping and Traceroute** tab must be changed from TCP to either ICMP or UDP.

### Credentials Vault Popup

The Credentials Vault popup, opened via the **Credentials Vault** button, enables you to apply credentials from the **Credentials Vault** to Synthetics tests for purposes such as authentication. Credentials are supported for the following types of tests:

* **Page Load** and **HTTP(S) or API**: The button appears in the HTTP tab of the settings page. It is present when the current tab of the table in the **Configure Request** pane is either **Headers or Body**.
* **Transaction**: The button appears in the **Targets and Agents** tab of the settings page, just above the **Puppeteer Script** field.

> **Note:** A step-by-step procedure for using a credential in an HTTP test is covered in **Set HTTP Test Credentials**.

#### Credentials Vault Popup UI

The Credentials Vault popup appears beside the **Credentials Vault** button when clicked, and includes the following:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(535).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A13Z&se=2026-05-12T10%3A15%3A13Z&sr=c&sp=r&sig=%2Bje1%2B89jPBifDyV%2BHMQ8PpAvTYctzC02eei4YuqZfmQ%3D)**Close**: An **X** in the upper right corner closes the popup, as does clicking the **Credentials Vault** button again.
* **Filter by name**: Enter text to filters the list to only matching credential names.

  + This field is case sensitive.
* **Filter by labels**: Displays a lozenge for each selected label and filters the list to only credentials with those labels.

  + **To select a label**: Click the field and choose the label from the searchable dropdown.
  + **To remove a label**: Click the **X** at the right of that label’s lozenge.
* **Credentials list**: A list of all credentials to which you have access, filtered by the credential’s name or label fields (see **Credentials Popup List**).

#### Credentials Popup List

The Credentials Vault popup has cards the credentials available for application to the current test.  
  
Expanding a credential card reveals the following additional elements:

* **Values list**: A list of the names of each key/value pair saved to the credential.

  > **Note**: The number of key/value pairs saved to the credential is shown in parentheses beside the Values header.
* **Copy** (icon): A button to the right of each key name that copies to your clipboard a programmatic representation of the key’s value (see **Value Variables**).
* **Notes**: Any text entered in the **Description** field when the credential was created or edited (see **Credential Settings Definitions**).
* **Updated**: The last time this credential was modified.

## Test Settings Tabs

Synthetic test settings are grouped into tabs in the page’s left sidebar.

> **Notes**:
>
> * To see which test settings tabs are available for each test type, see **Settings by Test Type**.
> * When your organization has no remaining **Synthetics Test Credits** a notification appears just below the heading on the settings tabs. The notification includes a **Submit a Request** button, which opens a version of the **Help & Support** form (see **Support Request**) that is pre-populated with a request for a quote for additional credits.

### Test Information

This tab includes the following general test settings:

* **Name** (required): The test name that appears in the Kentik portal.
* **Description**: Notes or comments that will appear when you hover over the Notes icon in the top right corner of a **Test Details Page**.
* **Labels**: Displays a lozenge for each selected label to apply to the test.

  + **To select a label**: Click the field and choose the label from the searchable dropdown.
  + **To remove a label**: Click the **X** at the right of that label's lozenge.

    > **Note:** Labels for tests can also be managed on the TCC landing page (see **Label Controls**).
* **Add Label**: A link to open the **New Label Dialog**, where you can add a new label.
* **Test Frequency**/**DNS Lookup Frequency**: The interval at which the test should be run. The type of the test determines which intervals are available (see **Intervals by Test Type**) and whether the field is required (see **Required Test Settings**).

> **Note:** The **Name**, **Description**, and **Labels** settings are common to all test types.

#### Intervals by Test Type

The table below shows how interval settings vary by test type.

| Test Information Setting | BGP Tests | Agent-to-Agent | Agent-to-Server | Autonomous Tests | DNS Tests | HTTP Tests |
| --- | --- | --- | --- | --- | --- | --- |
| Test Frequency or  DNS Lookup Frequency: 1s, 15s | No | Yes | Yes | Yes | No | No |
| Frequency: 1m, 2m | No | Yes | Yes | Yes | Yes | Yes (except Transaction) |
| Frequency: 5m, 10m,15m, 30m, 60m | No | Yes | Yes | Yes | Yes | Yes |
| Frequency: 90m | No | No | No | No | No | Transaction only |

> **Notes:**
>
> * BGP Monitor tests don’t have this setting.
> * The default Test Frequency interval is 5 minutes for Page Load tests, 30 minutes for Transaction tests, and one minute for all other tests.
> * All probes for each subtest in a test (see **Ping and Traceroute**) will be sent at the start of the test interval.

#### New Label Dialog

The New Label dialog allows you to create a new label to apply to the test, and includes the following:

* **Close**: An **X** at the upper right that exits the dialog without creating a label.
* **Swatch**: A popup color selector to choose the color of the label.
* **Name**: A text field in which to enter a name for the new
* **Cancel**: A button that exits the dialog without creating a label.
* **Save**: A button that creates a label with the specified color and name.

> **Note:** A new label is not applied automatically to the test. To apply it, click the **Labels** field to open the dropdown.

### Target and Agents

This tab includes settings for a test’s target and source agents, and other settings that vary by test type.

#### Specify Test Target

The test type determines which of the following target fields are included in the settings:

* **Target Agent**: Controls for selecting the agent to test toward, including:

  + **Agent list**: Shows the currently selected agent.
  + **Choose Agent**: Click to open the **Select Agents Dialog** to select an agent.
* **Target IP Addresses**: Enter a comma-separated list of IP address(es) to target.
* **Target Hostname**: Enter a hostname to target.
* **ASN, CSN, Country, Region, or City to Target**: See **Autonomous Target Control**.
* **URL or IP Address to Target**: Enter the URL/IP to be the target of the selected HTTP request.

#### General Target and Agent Settings

The remaining settings on the tab vary depending on the test type (see **Target and Agents by Test Type**).

* **Agent(s) to Test From**: Controls for selecting agents from which to run test, including:

  + **Agents list**: Lists the currently selected agents.
  + **Choose/Edit Agent(s)**: Click to open the **Select Agents Dialog** to select agents.

    > **Note**: For Autonomous tests, see **Autonomous Agent Control**.
* **IP Version**: Choose the version of IP addresses targeted by the test from the dropdown: IPv4 only, IPv6 only, or both (v4 + v6). The **v4 + v6** setting is particularly useful for Hostname or ASN tests, where there could be a mix of IPv4 and IPv6 addresses.
* **Test Bidirectionally**: A switch that enables/disables testing in both directions when testing between agents.
* **Use Agent Private IPs**: A switch that enables/disables private IPs for mesh testing between your private agents. Kentik will test using the local IP that can be assigned to each private agent via the **Private IP** field in the **Configure Agent Dialog**, resulting in a more direct path than testing using the public IP addresses of the private agents.
* **DNS Server IP(s) to Query**: Enter a comma-separated list of the IP addresses of a DNS server to query.
* **DNSSEC Validation**: A switch that enables/disables validation of the authenticity of each signing entity in the chain, from the authoritative name server up to the root server. The result is a healthy or critical result displayed in the DNSSEC column of the **Test Details Table** on the Results tab of the Test Details page.
* **Hostname to Look Up**: Enter the hostname you want this test to resolve via the specified DNS server(s).
* **Method**: Choose the method of HTTP request from the dropdown: GET, HEAD, PATCH, POST, or PUT.
* **URL or FQDN**: Enter the URL or Fully Qualified Domain Name (FQDN) that will be the target of the selected HTTP request.
* **Enable Ping and Traceroute Testing**: A checkbox that, when On (default), enables the following on the Subtest Details pages for this test (see **Subtest Details Tabs**):

  + Display of timeline charts for latency, packet loss, and jitter on the Metrics tab.
  + Display of traceroute results on the **Path View** tab.
* **Puppeteer script**: A textbox that, by default, contains an example Puppeteer script that you can use to experiment with Transaction tests. To create a test, overwrite the example script with a script that executes the steps that you'd like to test (see **About Transaction Test Scripts**).

> **Notes:**
>
> * The subtests in a test are always run from an agent to an IP. For a private agent in your AS, a subtest runs from your AS to the target IP. For a global agent, a subtest runs from the global agent to the target IP.
> * To test your own AS, select it as a target, then select the global agent from which to test.

#### Target and Agents by Test Type

The table below shows which **Target and Agents** settings are available for each type of test.

| Target and Agents Settings | Agent-to-Agent Tests | Agent-to-Server Tests | Autonomous Tests | DNS Tests | HTTP Tests |
| --- | --- | --- | --- | --- | --- |
| Target Agent | Agent-to-Agent only | No | No | No | No |
| Agent(s) to Test From | Yes | Yes | Yes | Yes | Yes |
| IP Version | Hostname and Agent-to-Agent only | Server Hostname only | Yes | No | Yes (except Transaction) |
| Test Bidirectionally | Agent-to-Agent only | No | No | No | No |
| Use Agent Private IPs | Yes | Yes (except Server Hostname) | No | No | No |
| Target IP Addresses | No | Yes (except Server Hostname) | No | No | No |
| Target Hostname | No | Server Hostname only | No | No | No |
| ASN, CSN, Country, Region, or City to Target | No | No | Yes | No | No |
| DNS Server IP(s) to Query | No | No | No | Yes | No |
| DNSSEC Validation | No | No | No | Yes | No |
| Hostname to Look Up | No | No | No | Yes | No |
| Method | No | No | No | No | HTTP(s) or API only |
| URL or FQDN | No | No | No | No | HTTP(s) or API only |
| Enable Ping and Traceroute Testing | No | No | No | Yes | Yes (except Transaction) |
| URL or IP Address to Target | No | No | No | No | Page Load only |
| Puppeteer Script | No | No | No | No | Transaction only |

> **Note**: Settings pages for BGP tests don’t include a Target and Agents tab.

#### Autonomous Target Control

The settings pages for tests in the Autonomous category include controls to identify a specific instance of a dimension whose IPs will be tested toward. The dimension corresponds to the type of test (ASN, CDN, Region, Country, or City) and the control’s label varies accordingly (e.g., for an ASN test the control is labeled “ASN to target”).  
  
The target control for Autonomous tests includes three tabs to choose the target instance. Using an ASN test as an example, these tabs would be as follows:

* **By Inbound Traffic**: A table listing the top-X ASNs sending traffic to your network (based on the flow data your organization sends to Kentik), including volume and status information. Click **Select** to choose one ASN.
* **By Outbound Traffic**: A table that is the same as the **By Inbound Traffic** tab but lists the top-X ASNs to which your network sends traffic.
* **Manual**:

  + **Target**: A drop-down with which you can choose a target from a filterable list of instances of the dimension corresponding to this test type (e.g., ASNs, CDNs, etc.). The label of the drop-down varies depending on the test type.
  + **Target is the**: A drop-down with which you specify whether the target is tested as the Destination of outbound traffic (default) or the Source of inbound traffic.

> **Note**: You must select a target (either inbound or outbound) before you can select any agent(s) to test from.

#### Autonomous Agent Control

Select an agent from which to run the test. The agent control includes the following three tabs from which you can choose agent(s):

* **By Inbound Traffic**: A table listing the agents available for the test.

  + Kentik looks at inbound traffic from the selected target (e.g., ASN) and identifies inbound Sites used by live traffic, offering them for testing.
  + If the none of these sites have an agent, Kentik will offer to set one up.
* **By Outbound Traffic**: A table that is the same as the By Inbound Traffic tab, but with available agents chosen from sites involved in delivering traffic to the configured target (e.g., ASN).
* **Manual**: Any selected agents from the other two tabs appear here. Add agent(s) manually by clicking the Choose/Edit Agents buttons to open the **Select Agents Dialog**.

> **Notes**:
>
> * At least one agent must be selected in order to create/save the test.
> * If no agents are available, add agents by clicking **Add an agent** (opens the **Agent Management** page in a new tab).

#### About Transaction Test Scripts

In Transaction tests, you use the **Target and Agents** tab to specify a script to perform the actions that you'd like to test. To do this, our **Synthetics Agent** embeds Chrome’s **Puppeteer**, which is a Node runtime that gives JavaScript programmatic access to a Chrome browser.  
  
You'll typically create a script in the **Recorder** tab of Chrome Developer Tools and paste it into the **Puppeteer Script** field on the **Target and Agents** tab. You'll then be able to insert into the script, at whichever stages of the transaction you wish, a command to take a screenshot, which will be displayed in the **Screensh****ots** pane of the **Results** tab on any of the test's Subtest Details pages (see **Create a Transaction Script**, **Validate a Transaction Script**, and **View Transaction Results**).  
  
To insert a credential into the script, use the**Credentials Vault** button (located directly above the **Puppeteer Script** field) to open the **Credentials Vault Popup**, from which you can copy a credential value. When you paste the value into the script, it will be represented programmatically (see **Value Variables**).

> **TIP:** To ensure that a value variable is correctly interpreted, always enclose `$vault(key)` in double quotes, for example:  
> `const vaultRef = "${vault('yourGroup yourKey') }"`

### BGP Monitoring

The **BGP Monitoring** tab is enabled by default for a BGP Monitor test (see **Routing Tests**) and is also available for some types of **Application Tests** (Page Load and HTTP(S) or API). In HTTP tests, the tab’s settings and the inclusion of BGP data in the test, are enabled with the On/Off switch (see **Settings Tabs**).  
  
The tab includes the following settings:

* **Prefixes to Monitor**: Enter a comma-separated list of up to 10 prefixes to monitor.
* **Add prefixes from.txt or.csv file**: Click to open an OS file selection dialog to select a .txt or .csv file with the prefixes to monitor.
* **Include more specific prefixes**: A switch, that when On, Kentik will automatically discover and monitor more specific prefixes within the listed prefixes. Switch only available when all prefixes listed in the Prefixes to Monitor field are /15 or narrower, e.g., /16, /21 (see **Visualizing Alert Prefixes**).
* **Origin Hijack Detection**: Displays a lozenge for each ASN selected for route announcement and withdrawal monitoring.

  + **To select an ASN**: Click the field and choose the ASN from the searchable dropdown.
  + **To remove an ASN**: Click the **X** at the right of that ASN’s lozenge.
* **Add ASN(s) from.txt or.csv file**: Click to open an OS file selection dialog to select a .txt or .csv file with the ASNs to monitor.
* **Check RPKI**: A switch, that when On, ASNs not explicitly listed in the **Allowed ASNs** field, but that are RPKI-valid, will be treated as valid by Kentik. When Off, those ASNs will be treated as invalid by Kentik.
* **Upstream Leak Detection**: Enter the ASNs that you expect to be upstream of originating ASNs for the prefixes provided.

  + **To add an ASN to the list**: Click the field and choose the ASN from the searchable dropdown; the ASN will appear as a lozenge in the field.
  + **To remove an ASN**: Click the **X** at the right of that ASN’s lozenge.

    > **Note**: An upstream leak for any ASN listed in this field can be seen on the **BGP Monitor Details Page** when you choose **Upstream** from the Options menu (vertical dots icon) on the **Reachability/Visibility** pane.
* **Add ASN(s) from .txt or .csv file**: Click to open an OS file selection dialog to select a .txt or .csv file listing the expected upstream ASNs.

#### Visualizing Alert Prefixes

The**Include more specific prefixes** switch has the following effect on the **Reachability Time Series** visualization (Test Details page) when a test generates an alert (see **Alerting and Notifications**):

* **If the switch is Off**: The visualization won't indicate which of the prefixes in the **Prefixes to Monitor** field generated the alert.
* **If the switch is On**: The visualization will show the individual prefix that generated an alert, but only down to the level specified in the **Prefixes to Monitor** field ("more specific prefixes" are not displayed individually).

### Flow-Based Targeting

The **Flow-Based Targeting** tab is available only for Autonomous tests, and includes the following:

* **Top Source or Destination**: Choose from the dropdown whether the tested IPs of the entity to test toward (see **Target and Agents**) are the top source IPs or top destination IPs (default).
* **Target**: Select from the dropdown whether Kentik will test toward IPs within the target entity that are either sending traffic to your AS (Source of inbound traffic) or receiving traffic from your AS (Destination of outbound traffic; default).
* **Max Number of Providers to Track**: Use the slider to set the maximum number of providers to track autonomously (see **Provider Classification**).
* **Max Number of IP Targets to Track**: Use the slider to set the maximum number of IPs toward which Kentik should create subtests for this test. The actual number of subtests will depend on the number of IPs that meet the qualifications determined by other settings; this setting caps that number (and the consumption of test credits).
* **Frequency to Scan for New Targets**: Choose from the dropdown the interval at which your organization's flow data will be evaluated by Kentik to identify new targets (IP addresses). The default is “Every 12 hours”.

  > **Note:** This setting does not affect test frequency.

### HTTP Settings

The **HTTP** settings tab is available only for HTTP(S) or API, Page Load, and Transaction tests, and includes the following settings:

* **HTTP Timeout**: Enter the duration in ms for Kentik to wait for a response to a test HTTP request. If no response is received before the timeout, no results will be shown for the HTTP portion of the test, and health status will be determined by ping, if enabled.

  > **Note:** The default timeout is 5 seconds (5000 ms) for Transaction and HTTP(S) or API tests and 20 seconds (20000 ms) for Page Load tests.
* **Ignore TLS Errors**: A switch that determines whether a test in the HTTP category will ignore errors related to Transport Layer Security (encryption of data in transit). See **Ignore TLS Errors**.
* **Configure Request** (HTTP(S) or API test only): A set of tabs used to specify headers, parameters, and body for the HTTP request (see **Configure Request Settings**).
* **Configure Request and CSS Selectors Validation** (Page Load test only): A set of tabs used to specify headers, parameters, and CSS selectors for the HTTP request (see **Configure Request Settings**).

#### Ignore TLS Errors

The **Ignore TLS Errors** switch determines whether an HTTP-category test will ignore errors related to Transport Layer Security (encryption of data in transit). The effect of the switch varies depending on the specific test type:

* **HTTP(S) or API** and **Page Load**:

  + **If the switch is On**: The Certificate Expiry column does not display in the test results.
  + **If the switch is Off and the server certificate is expired**: A Certificate Expiry column is shown in the test results and a red lozenge will read “SSL Certificate Expired.”
  + **If the switch is Off and the certificate has not expired**: The Certificate Expiry column shows a lozenge that displays the certificate’s expiry date (see **Health Settings**).
* **Transaction**: If the switch is On, the test proceeds even if there are certificate issues.

#### Configure Request Settings

The optional **Configure Request** settings for Page Load tests and HTTP(S) or API tests enable additional control over how the test is structured and what it looks for:

* **Test structure**: Use the **Headers**, **Params**, and **Body** tabs to control HTTP configuration for the test (see **HTTP Test Structure**).

> **Note**: The Body tab is available only for HTTP(S) or API tests (see **Configure Request Body**).

* **CSS Selectors** (Page Load test only): Use the **CSS Selectors** tab to specify one or more selectors for the test to look for in a web page’s HTML (see **CSS Selectors**).

> **Notes**:
>
> * The tab heads forHeaders, Params, and CSS Selectors include a count (in parentheses) of the number of each used for the test.
> * The **Key**/**Value** rows each include a Remove button (trash icon) that you can click to remove a key/value pair.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(536).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A13Z&se=2026-05-12T10%3A15%3A13Z&sr=c&sp=r&sig=%2Bje1%2B89jPBifDyV%2BHMQ8PpAvTYctzC02eei4YuqZfmQ%3D)

#### HTTP Test Structure

The following tabs enable fine control of the HTTP configuration for a test:

* **Headers**: A tab where you can specify key/value pairs for HTTP request headers. The **Add Headers** button adds a row for another key/value pair.
* **Params**: A tab where you can specify key/value pairs to be passed as arguments in the query string of the test's URL. The **Add Params** button adds a row for another key/value pair.
* **Body**(HTTP(S) or API tests only): A tab where you can specify the payload of the HTTP request (see **Configure Request Body**).

  > **Note:** This tab is inactive when the test's **Method** is set to GET (see **General Target and Agent Settings**).

> **Notes:**
>
> * The Remove button (trash icon) at the right of each key/value row removes that key/value pair.
> * On the Headers and CSS Selectors tabs, the **Credentials Vault** button opens the **Credentials Vault Popup**, where you can copy credential values to paste into the **Key** and/or **Value** fields.

#### Configure Request Body

The **Body** tab of the **Configure Request** pane allows you to set the payload for the HTTP Request of an HTTP(S) or API test. The tab is active unless the test's Method is set to GET (see **General Target and Agent Settings**).  
  
Choose a request body format using the radio buttons at the top:

* **none**: The request will have an empty payload.
* **x-www-form-urlencoded**: The request body will be derived from strings entered in the **Key** and **Value** fields. The **Add Form Values** button adds a key/value row. The **Remove** button (trash icon) at the right of each key/value row removes that row.
* **raw**: The request body will be whatever is entered into the text field, with the format chosen from the dropdown (Text, JavaScript, JSON, HTML, or XML).

> **TIP:** The **Credentials Vault** button opens the **Credentials Vault Popup**, where you can copy credential values to paste into the text field.

#### CSS Selectors

The **CSS Selectors** tab (included only for Page Load tests) is where you can specify one or moreCSS selectors to  check for in the HTML of the tested web page, resulting in the following:

* A **CSS Validation** column will appear on the test's details page, indicating PASS only if all CSS selectors listed in the CSS Selectors tab are found, otherwise indicating FAIL.

  > **Note:** This result is not used in computing test health (which drives alerts and notifications).
* The **Metrics** tab of the **Subtest Details Page** for this test will include a CSS Selectors chart (see **Additional Metrics Charts**).

Each selector includes the following:

* **Name**: The user-supplied selector name, e.g., `submit button`.
* **CSS Selector**: The actual selector to check for in the tested page’s CSS, e.g., `#submit`.

The **Add CSS Selectors** button adds a row with another pair of fields. The **Remove** button (trash icon) at the right of each row removes that row.

### Ping and Traceroute

The **Ping and Traceroute** settings tab is available for all Network tests, as well as for HTTP tests except for Transaction. The tab includes two main settings panes: **Ping Options** and **Trace Options**.

#### Ping and Traceroute by Test Type

The table below shows the categories of test types for which Ping and Traceroute settings are available.

|  | BGP Tests | Agent-to-Agent | Agent-to-Server | Autonomous Tests | DNS Tests | HTTP Tests |
| --- | --- | --- | --- | --- | --- | --- |
| Ping Options | No | Yes | Yes | Yes | Yes | Yes (except Transaction) |
| Trace Options | No | Yes | Yes | Yes | Yes | Yes (except Transaction) |

#### Ping Options

The following ping-related settings, which vary by test type (see **Ping and Traceroute by Test Type**) appear in the Ping Options pane of the **Ping and Traceroute** tab:

* **Number of Probes per Ping**: Use the slider to set the number of individual packets to be sent per ping (default is 5).
* **Overall Ping Timeout**: Enter the duration in ms for Kentik to wait for the full ping test to complete (default is 3000). A “per-probe” timeout is derived by dividing this number by the **Number of** **Probes per Ping** setting above.
* **Inter-Probe Delay**: Enter the delay in ms to be added between consecutive ping probes. Default is 0 ms (no delay).
* **Target Port**: Enter the port to which the packets will be sent. The default, which depends on the **Protocol** setting, is 433 for TCP.

  > **Note**: This field is inactive when Protocol is set to ICMP.
* **Protocol**: Choose the protocol from the dropdown to use for sending packets. Protocols vary by test type, but can include UDP-ICMP, UDP-ECHO, ICMP, or TCP (default).

  > **Notes**:
  >
  > + Tests involving Global Agents must be set to a protocol other than TCP.
  > + Once a test is saved, the dropdown will be inactive and protocol can't be changed.
* **DSCP**: Choose the Differentiated Services Code Point (DSCP) value from the dropdown to test along all your classes of services. The default is “Best Effort (DSCP 0)” (see **DSCP Options**).

#### DSCP Options

The table below shows the available DSCP options.

| Option | DSCP Value | Description |
| --- | --- | --- |
| Best Effort (DSCP 0) | 0 | Best Effort |
| CS1 (DSCP 8) | 8 | Class 1 (CS1); Class Selector |
| AF11 (DSCP 10) | 10 | Class 1, Gold (AF11); Assured Forwarding |
| AF12 (DSCP 12) | 12 | Class 1, Silver (AF12); Assured Forwarding |
| AF13 (DSCP 14) | 14 | Class 1, Bronze (AF13); Assured Forwarding |
| CS2 (DSCP 16) | 16 | Class 2 (CS2); Class Selector |
| AF21 (DSCP 18) | 18 | Class 2, Gold (AF21); Assured Forwarding |
| AF22 (DSCP 20) | 20 | Class 2, Silver (AF22); Assured Forwarding |
| AF23 (DSCP 22) | 22 | Class 2, Bronze (AF23); Assured Forwarding |
| CS3 (DSCP 24) | 24 | Class 3 (CS3); Class Selector |
| AF31 (DSCP 26) | 26 | Class 3, Gold (AF31); Assured Forwarding |
| AF32 (DSCP 28) | 28 | Class 3, Silver (AF32); Assured Forwarding |
| AF33 (DSCP 30) | 30 | Class 3, Bronze (AF33); Assured Forwarding |
| CS4 (DSCP 32) | 32 | Class 4 (CS4); Class Selector |
| AF41 (DSCP 34) | 34 | Class 4, Gold (AF41); Assured Forwarding |
| AF42 (DSCP 36) | 36 | Class 4, Silver (AF42); Assured Forwarding |
| AF43 (DSCP 38) | 38 | Class 4, Bronze (AF43); Assured Forwarding |
| CS5 (DSCP 40) | 40 | Class 5 (CS5); Class Selector |
| Voice Admit (DSCP 44) | 44 | Voice Admit PHB |
| EF (DSCP 46) | 46 | Expedited Forwarding (EF) |
| CS6 (DSCP 48) | 48 | Control (CS6); Class Selector |
| CS7 (DSCP 56) | 56 | Control (CS7); Class Selector |

#### Trace Options

The Trace Options pane of the **Ping and Traceroute** tab contains the following:

* **Number of Probes per Hop**: Set the slider to the number of individual packets to be sent for each router in the path (default is 3).
* **Overall Trace Timeout**: Enter the duration in ms for Kentik to wait for the full ping traceroute to complete (default is 22500). A “per-probe” timeout is derived by dividing this number by the **Number of Probes per Hop** setting above.

  > **Note:** To minimize incomplete traces in the diagram on the **Path View** tab (see **Test Path View**), Kentik recommends a value of 7500 ms for each probe per hop.
* **Inter-Probe Delay**: Enter the delay in ms to be added between consecutive traceroute probes. Default is 0 ms (no delay).
* **Target Port**: Enter the port to which the packets will be sent (default is 443).

  > **Note**: If ICMP is selected as the traceroute protocol, you cannot select a target port and this field will be grayed out.
* **Protocol**: Choose the protocol from the dropdown to use for sending packets. Protocols vary by test type, but can include UDP-ICMP, UDP-ECHO, ICMP, or TCP (default).

  > **Note**: Tests involving Global Agents must be set to a protocol other than TCP.
* **Max Number of Hops (Max TTL)**: Enter the maximum number of hops to trace (default is 30).

  > **Note:** If the hops in the route being traced exceeds the max, the traces in the diagram on the Path View tab will be truncated.
* **DSCP**: Choose the Differentiated Services Code Point (DSCP) value from the dropdown to test along all your classes of services. The default is “Best Effort (DSCP 0)” (see **DSCP Options**).

### DNS Settings

The **DNS** settings tab is available only for tests in the DNS category (DNS Server Monitor and DNS Server Grid), and includes the following:

* **Target Port**: Enter the port on the DNS server being tested (default is 53).
* **DNS Record Type**: Choose the record type that the test will request from the server (see **List of DNS Record Types**).

  > **Note:** When the selected type is A or AAAA, the **Allowed DNS Results** field will be visible in **Health Settings**.

### Health Settings

The settings on the **Health** tab vary depending on the type of test, and may include the following:

* **Thresholds**: Set the threshold for reporting subtest health status as Warning and/or Critical (see **Health Threshold Controls**). The types of thresholds that can be set vary by test type (see **Health Threshold Metrics**).
* **Expected Responses**: Specify one or more key/value pairs for the headers required for Kentik to consider the response to be healthy.

  + **To add a key/value row**: Click **Add Headers**.
  + **To removes a row**: Click the Remove button (trash icon) at the right of the row.
* **Valid HTTP Codes**: Displays a lozenge for each selected HTTP status code that Kentik will consider normal vs. warning or critical.

  + **To select an HTTP code**: Click the field and choose one or more codes from the searchable dropdown.
  + **To remove a code**: Click the **X** in the code’s lozenge.

    > **Note:** When no codes are specified, all codes above 400 will result in a health status of critical.
* **Valid DNS Codes**: Displays a lozenge for each selected DNS code that Kentik will consider normal vs. warning or critical.

  + **To select a DNS code**: Click the field and choose one or more codes from the searchable dropdown.
  + **To remove a code**: Click the **X** in the code’s lozenge.

    > **Note:** When no codes are specified, all codes above 400 will result in a health status of critical.
* **Allowed DNS Results**: Specify (using a comma-separated list) the IPs that Kentik should consider healthy when found in the test's DNS results.

  > **Note**: This field is only visible if the DNS Record Type (under **DNS Settings**) is set to A or AAAA.
* **Test time-outs and failures**: Turn the switch On to receive notifications when the test fails to receive a response (e.g. test times out, or a TLS Cert Validation prevents the test from getting valid data).

  > **Note**: This does not include successful tests that returned an undesirable result (e.g., 100% packet loss).

#### Health Threshold Controls

Each threshold control determines how the subtest’s health status for a given metric (latency, jitter, etc.) is reported in the system. The following controls are included:

* **Enabled/Disabled**: A switch that adds (enabled) or removes (disabled) this metric from health calculations. At least one metric must remain enabled in order to save the test.

  > **TIP**: While a disabled switch removes this metric from health calculations, and disables alerts, you can still see the data in the UI.
* **Use**: ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(537).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A13Z&se=2026-05-12T10%3A15%3A13Z&sr=c&sp=r&sig=%2Bje1%2B89jPBifDyV%2BHMQ8PpAvTYctzC02eei4YuqZfmQ%3D)Choose the threshold type/unit (e.g. milliseconds, percentage, standard deviation, or days) from the dropdown. May not be present for all metrics.
* **Warning**: Enter a precise value (greater than 0) for the Warning threshold.
* **Critical**: Enter a precise value (greater than 0) for the Critical threshold. This value can be greater than the maximum shown on the slider, except for percentages where it must be between 1 and 100.

  > **Note**: The threshold slider automatically updates based on the values entered in the Warning and Critical fields.
* **Thresholds**: Two sliders that you can move along the colored bar to set the following thresholds for a metric:

  + **Warning**: The left slider sets the threshold below which the metric's health status will be considered Normal (green) and above which it will be Warning (orange).
  + **Critical**: The right slider sets the threshold below which the metric's health status will be considered Warning (orange) and above which it will be Critical (red).

#### Health Threshold Metrics

**Health Threshold Controls** are available for the following metrics, though their presence on the **Health** tab for a given test depends on the test's type:

* **Ping Latency**: Latency in standard deviation or in milliseconds (choose the metric from the Use drop-down).
* **Ping Jitter**: Jitter in standard deviation or in milliseconds.
* **Ping Packet Loss**: Packet loss in percentages:

  + The value in the **Warning** field must be 0 or greater.
  + The value in the **Critical** field must be between 1 and 100.
* **BGP Reachability**: The number of "vantage points" (public BGP routing data collectors) that are able to reach the tested prefixes as a percentage of all such vantage points. If reachability drops below the specified thresholds for Warning and/or Critical, an alert will be triggered in the Synthetics **Incident Log**, and a notification will also be triggered if one is set for the test.
* **HTTP Latency**: HTTP latency in standard deviation or in milliseconds.
* **Certificate Expiry**: The number of days that you'll be notified in advance of the expiration of a certificate, as well as what color the lozenge should be for each state (appears under the **Certificate Expiry** column of the Test Details page).
* **DNS Response Time**: DNS response time in standard deviation or in milliseconds.
* **Transaction time**: Elapsed time of transaction.

> **Notes:**
>
> * For thresholds using standard deviation, default values are 1.5 for **Warning** and 3 for **Critical**.
> * For thresholds using milliseconds, default values are 25 ms for **Warning** and 75 ms for **Critical**.

#### Health Options by Test Type

The table below shows which options are included in the **Health** settings for each test category.

| Health Options | BGP Tests | Agent-to-Agent | Agent-to-Server | Autonomous Tests | DNS Tests | HTTP Tests |
| --- | --- | --- | --- | --- | --- | --- |
| BGP Reachability | Yes | No | No | No | No | No |
| Ping Latency Thresholds | No | Yes | Yes | Yes | No | Yes (except Transaction) |
| Ping Jitter Thresholds | No | Yes | Yes | Yes | No | Yes (except Transaction) |
| Ping Packet Loss Thresholds | No | Yes | Yes | Yes | No | Yes (except Transaction) |
| HTTP Latency Thresholds | No | No | No | No | No | Yes (except Transaction) |
| Certificate Expiry Thresholds | No | No | No | No | No | Yes (except Transaction) |
| Response Time Thresholds | No | No | No | No | Yes | No |
| Transaction Time Thresholds | No | No | No | No | No | Transaction only |
| Valid HTTP Codes | No | No | No | No | No | Yes (except Transaction) |
| Valid DNS Codes | No | No | No | No | Yes | No |
| Allowed DNS Results | No | No | No | No | Yes | No |
| Expected Responses | No | No | No | No | No | Yes (except Transaction) |
| Alerts on test and agent failure | No | Yes | Yes | Yes | Yes | Yes |

### Alerting and Notifications

Use the **Alerting and Notifications** tab to configure alerts and notifications for all test types. The settings are enabled/disabled with the **On/Off** switch (see **Settings Tabs**). Disabling the tab turns off the alerts and notifications but maintains existing settings.

#### Alerting Conditions

The **Alerting Conditions** control set, which defines the conditions that will trigger an unhealthy status alert for the test agent, includes the following controls:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(538).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A13Z&se=2026-05-12T10%3A15%3A13Z&sr=c&sp=r&sig=%2Bje1%2B89jPBifDyV%2BHMQ8PpAvTYctzC02eei4YuqZfmQ%3D)**Times**: Enter the number of unhealthy status incidents (a metric’s health threshold reports as Warning or Critical) that will trigger an alert (default is 3).
* **Duration value**: Enter the number of time units in the duration (default is 5).
* **Duration units**: Select the time unit of the duration, either minutes (default) or hours from the dropdown.
* **Reset period**: Enter the number of minutes after which the counter will reset to 0 if the health threshold hasn’t been met within that span of time.

> **Note:** An unhealthy status condition refers to any **Health Threshold Metrics** threshold (warning or critical) or binary threshold (e.g., Valid HTTP Codes; see **Health Options by Test Type**) reached.

#### Notification Options

This pane enables you to send notifications of a test’s status to one or more notification channels defined in your organization (see **Notifications**). The following controls are included:

* **Notification Channels**: A field that shows a lozenge for each of the channels to which notifications are currently assigned for this test. To add a channel, click in the field to drop down a filterable list of channels in your organization, then click on a channel. To remove a channel, click the X in the channel’s lozenge.
* **Test Notification Channels**: A button, present only if at least one channel is selected in the Notification Channel field, that sends a test notification to the recipients in all of the currently selected channels.
* **Add New**: A button that opens the Add Notification dialog (see **Add or Edit Notification**), where you can create a new notification channel in your organization and automatically add it to the test.

#### Alert Suppressions

This pane enables you to mute alert notifications for a specified period. The pane has several distinct modes, each with their own UI elements:

* **Initial**: If no alert suppressions are currently applied to the test, the pane contains only a **Create Suppression** button. Click the button to show the control set used to specify an alert suppression.
* **Specify suppression**: The pane contains the control set used to specify an alert suppression (see **Specify Alert Suppression**).
* **Existing suppressions**: The pane includes a table showing the alert suppressions currently in place for this test (see **Existing Test Suppressions**).

#### Specify Alert Suppression

The specify suppression mode of the Alert Suppressions pane includes the following fields and controls:

* **Expires in**: States the time (e.g., number of days) until the test will resume sending alerts and notifications. If **Never Expire** is On, the indicator says "Expires never".
* **Silence window** (field, only when **Never Expire** is Off): Shows the date-time at which suppression will start and stop. Click to edit the following:

  + **Start**: Edit the start date-time.
  + **End**: Edit the end date-time.
  + **Calendar**: A popup calendar from which you can choose start and end times.
* **Start time** (field, only when **Never Expire** is On): Shows the date-time at which suppression will start. Click in the field to edit it, or choose a date from the calendar popup.
* **Never Expire** (switch): Determines if the alert suppression will expire after a set period of time (Off, default) or never (On).
* **Comment** (field): Enter a text comment, visible in **Existing Test Suppressions**.
* **Test(s)**: Specify one or more tests to be suppressed (see **Suppressions Test Selector**).
* **Cancel** (button): Closes the Alert Suppression controls without saving changes.
* **Save** (button): Saves the alert suppression based on the current settings and closes the controls.

#### Suppressions Test Selector

Use the **Tests** field to specify tests to be suppressed, each represented by a lozenge. When first opened from the **Create Suppression** button, the field contains a lozenge for the current test (the test being added/edited):

* **To choose additional tests to suppress**: Click the field and select a test from the searchable dropdown.
* **To stop suppressing a test**: Click **X** on the test’s lozenge or choose **None** to remove all tests, and click **Save**. Suppression of the selected tests is not applied until you click **Save**.

> **Note**: When you apply suppressions to any tests other than the current test, the **Alerting and Notifications** tab on that test’s settings page shows them in **Existing Test Suppressions**.

#### Existing Test Suppressions

The **Alert Suppressions** pane lists all suppressions applied to the current test, including Comment, Created, Started, and Expires values for the related settings (see **Specify Alert Suppression**). The Started and Expires dates are colored green if active, black if not yet active, and red if expired. The pane has the following controls:

* **Edit** (pencil icon): Opens the test suppression controls (see **Specify Alert Suppression**).
* **Remove** (trash icon): Opens a confirmation dialog to remove the alert suppression.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(539).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A13Z&se=2026-05-12T10%3A15%3A13Z&sr=c&sp=r&sig=%2Bje1%2B89jPBifDyV%2BHMQ8PpAvTYctzC02eei4YuqZfmQ%3D)       

## Manage Synthetics Tests

The following topics walk you through common tasks for Kentik Synthetics tests.

### Add a Test

To add a new test in the Test Control Center:

1. On the TCC landing page, click **Add Test** at the upper right to open the **Add Test Page**.
2. Select the test type you’d like to create.
3. On the **Test Settings Page**, choose **Test Information**.
4. Enter a test name and select a **Test Frequency** from the dropdown.
5. Click **Target and Agents** in the sidebar and complete all required fields.

   > **Note:** This step doesn’t apply to BGP Monitor tests.
6. Continue through any other tabs (e.g., **Optional Settings**) to customize additional settings for the test as desired.
7. Check the **Settings Status Indicators** for any errors.
8. Click **Preview** to open a **Test Preview Page** to check that the current settings are yielding your desired results.

   > **Note:** Test previews are not currently available for Autonomous tests.
9. Click **Create Test** to save the test and return to the TCC landing page.

### Edit a Test

To edit an existing test in the Test Control Center:

1. In the **Tests List**, find the test that you want to edit.
2. Click the **Edit** icon at the right to open the **Test Settings Page**.
3. Navigate via the sidebar to the **Test Settings Tabs** and modify the test as desired.
4. Click **Preview** to open a **Test Preview Page** to check that the current settings are yielding your desired results.

   > **Note:** Test previews are not available for Autonomous tests.
5. Click **Save** to save the test and return to the TCC landing page.

> **Note**: You can also reach a test’s settings page by clicking **Edit Test** in the subnav of the test’s **Test Details Page**.

### Copy a Test

To copy a test in the Test Control Center:

1. In the **Tests List**, find the test that you’d like to copy.
2. Click the **Copy** icon at the right to open the Test Settings page for the duplicated test.
3. In the **Test Information** tab, enter a new name for the test.
4. Navigate via the sidebar to the **Test Settings Tabs** and modify the test as desired.
5. Click **Preview** to open a **Test Preview Page** to check that the current settings are yielding your desired results.

   > **Note:** Test previews are not available for Autonomous tests.
6. Click **Create Test** to save the test and return to the TCC landing page.

### Remove a Test

There are two ways to remove a test from your organization:  
  
From the **Test Control Center** page:

1. In the Tests list, find the row of the test you’d like to remove.
2. Click the **Remove** icon (red trash can).
3. On the popup, click **Remove** to remove the test from the system.

From a **Test Settings Page**:

1. On the test’s Test Settings page, at the right of the Test Management Controls, click **Remove.**
2. On the popup, click **Remove** to remove the test from the system and return to the Test Control Center page.

### Create a Transaction Script

> **TIP:** Kentik continually updates the **Synthetics Agent** to optimize the creation of Transaction scripts with Google Chrome. For the supported Chrome versions for script creation, check with Product Support (see **Customer Care**).

To create a Transaction test from a script:

1. Open Chrome Developer tools and go to the **Recorder** tab.
2. Record the sequence of actions that you'd like to test.

   > **Note:** A test may not include more than five separate pages.
3. ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(540).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A13Z&se=2026-05-12T10%3A15%3A13Z&sr=c&sp=r&sig=%2Bje1%2B89jPBifDyV%2BHMQ8PpAvTYctzC02eei4YuqZfmQ%3D)Export the recording to a Puppeteer script (see description in this **Google article**).
4. In the Kentik portal, choose **Test Control Center** from the navbar menu (under Synthetics).
5. On the TCC landing page, click the **Add Test** button at upper right.
6. On the Add Test page, in the Application category, under HTTP, click **Transaction**.
7. On the **Test Information** tab of the resulting Test Settings page, enter a test name.
8. On the **Target and Agents** tab, choose the agents from which to test.  
   In the **Puppeteer Script** textbox, select the entire contents and paste in the script created above.

   > **IMPORTANT:** The script must meet the requirements listed in **Validate a Transaction Script**.
9. At each point in the series of actions where you'd like a screenshot taken of the page state, insert the following line at the corresponding point in the script:

   ```
   await page.screenshot({ path: "screenshot##.jpg" });
   ```

   > **Note:** Replace `##` with the number of the screenshot (`01` for the first, `02` for the second, etc.). A test may include no more than 10 screenshots.
10. If desired, under **Optional Settings**, change default settings for **Health Settings** and/or **Alerting and Notifications**.
11. Click the **Create Test** button to save and start the test.

    > **TIP:** If the button is not active and you've specified the test’s name/agents, make sure your script meets the requirements listed in **Validate a Transaction Script**.

### Validate a Transaction Script

To check that the script for a transaction test is valid, check that it meets the following constraints:

* The declarations for the `puppeteer`, `browser`, and `page` constants match the default example script:

  ```
  const puppeteer = require('puppeteer');
    const browser = await puppeteer.launch();
    const page = await browser.newPage();
  ```
* Five `page.goto` calls maximum.
* Ten `page.screenshot` calls maximum.
* 10,000 characters maximum

### View Transaction Results

Once a Transaction test is configured and running, here's how to see the results:

1. On the **Test Control Center** landing page, click the test name to open its details page.
2. At the bottom of the **Results** tab, click **Details** at the right of a subtest to open its details page.
3. The **Results** tab charts the transaction completion time for each time slice of the time range. Also shown are screenshots for the currently selected time slice.

The subnav is the gray, horizontal bar or strip located below the main navigation bar or navbar in various pages and modules. It typically contains page-wide controls and buttons specific to the page's functionality.
