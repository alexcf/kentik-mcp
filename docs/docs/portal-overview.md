# Portal Overview

Source: https://kb.kentik.com/docs/portal-overview

---

This article provides a basic introduction to the Kentik portal.

![Network Explorer dashboard displaying bandwidth metrics and cloud service performance statistics.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/PO-portal-landing-main.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A12Z&se=2026-05-12T09%3A33%3A12Z&sr=c&sp=r&sig=ylR7bDFpq2l8B%2FUqOTx6L5xa7kj3FuP%2BjPUFacAdsB0%3D)

*The default Kentik portal landing page is configurable to Network Explorer or other modules.*

> **Note:** For browser support and login procedures for the portal, see **Kentik Portal**.

## Portal Sections

The Kentik portal consists of modules (sometimes called “workflows”) designed to help you efficiently access information for specific networking use cases.

In the **Portal Main Menu**, the most commonly used modules are featured at the top of the left column: **Alerting****,** **AI Advisor****,** **Kentik Map****,** **Library****,** **My Kentik Portal**, and **Observation Deck**.

The remaining modules are categorized into the following main sections:

> **Note**: Additional modules not grouped in the main menu sections are covered in **Portal Main Menu**.

| Portal Section | Included Modules | Description |
| --- | --- | --- |
| **Settings** | **Main Settings** | Administrative controls to tailor your Kentik setup for you and your organization. |
| **Core** | **Network Explorer****,** **Data Explorer****,** **Capacity Planning****,** **Insights****,** **Raw Flow Explorer** | Provides a comprehensive view of your day-to-day network operations. |
| **Synthetics** | **Synthetics Dashboard****,** **Test Control Center****,** **Agents****,** **BGP Route Viewer****,** **State of the Internet** | Continuously monitor network performance, enabling proactive testing with public and private agents. |
| **Network Monitoring System** | **NMS Dashboard****,** **Metrics Explorer****,** **Query Assistant****,** **Devices****,** **Interfaces** | Provides a base layer of visibility by discovering and monitoring network infrastructure. |
| **Edge** | **Traffic Engineering****,** **Connectivity Costs****,** **Traffic Costs****,** **Discover Peers** | Enables Ops/Engineering teams to understand network utilization and manage costs related to external traffic. |
| **Protect** | **DDos Defense****,** **Mitigations****,** **Botnet & Threat Analysis****,** **RPKI Analysis** | Defends your network by detecting anomalies and providing response mechanisms against DDoS and other threats\*. |
| **Service Provider** | **CDN Analytics****,** **OTT Service Tracking****,** **Kentik Market Intelligence** | Helps you understand service provider requirements and track performance for customers and subscribers. |
| **Cloud** | **Kentik Cloud****,** **Kentik Kube****,** **Cloud Traffic Overviews**, **Performance Monitor****,** **Cloud Pathfinder** | Manages public and hybrid cloud networking by evaluating  performance and minimizing costs. |

> **Notes:**
>
> * The **DDoS Defense** and **Mitigations** modules work with **Alerting**.
> * If you aren’t seeing Kentik Kube, it may not be enabled for your account. Please contact your **Kentik Account Team**.

## Portal Page Structure

The basic structure of portal pages is covered here.

### Portal Page Layout

The portal layout’s main elements are similar across most pages:

![Kentik interface showing Core and Network Explorer options with search functionality.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/PO-portal-navbar-subnav.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A12Z&se=2026-05-12T09%3A33%3A12Z&sr=c&sp=r&sig=ylR7bDFpq2l8B%2FUqOTx6L5xa7kj3FuP%2BjPUFacAdsB0%3D)

*The Kentik portal’s main navbar with Network Explorer subnav shown directly beneath.*

* **Navbar**: The primary horizontal navigation bar at the top of the screen. It includes links to global portal functions, such as search, notifications, and user settings (see **Portal Navbar**).
* **Subnav**: The secondary horizontal bar located immediately below the navbar. Its layout and available controls dynamically update based on the specific module or page you are viewing (see **Portal Subnav**).
* **Main menu**: The navigation overlay accessed via the hamburger menu icon in the upper-left corner. It provides categorized access to all portal modules, workflows, and frequently used pages (see **Portal Main Menu**).
* **Main content**: The central workspace of the screen. This area houses your data visualizations, tables, query controls, and sidebars, which vary depending on your active workflow.

### Portal Navbar

The portal's main navbar is located at the top of the screen and includes the following options from left to right:

* **Menu** (hamburger icon): Opens the main menu, allowing navigation to portal modules and workflows (see **Portal Main Menu**).
* **Kentik logo**: Links to your designated landing page (see **User-Specific Defaults**).
* **Search**: Opens the global **Search** modal to find portal text, favorites, and recently viewed pages (see **Portal Search**).
* **Ask**: Opens **AI Advisor** in overlay mode.![Discovery process complete, showing one device found in the interface notification.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/PO-portal-navbar-discovery.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A12Z&se=2026-05-12T09%3A33%3A12Z&sr=c&sp=r&sig=ylR7bDFpq2l8B%2FUqOTx6L5xa7kj3FuP%2BjPUFacAdsB0%3D)
* **Device Discovery**: Opens the **Device Discovery** page. Icon includes a badge for the number of devices discovered. Hover to get a summary popup.
* **Product Updates** (megaphone icon): Opens a popup containing the latest portal and Kentik product changes (see **Product Updates**).
* **Help and Support** (question mark icon): Opens the **Help and Support** popup to search the Knowledge Base or submit a ticket (see **Portal Help and Support**).
* **Organization settings** (building icon): Drops down a menu of organization-wide settings (see **Organization Settings**).
* **User settings** (user icon): Drops down a menu of user-specific settings, including theme preferences and the option to log out (see **User Menu**).

> **Note:** In addition to the **Organization Settings** and **User** menus, various administrative settings can be accessed via the central **Main Settings** page.

#### Product Updates

Accessed via the megaphone icon in the navbar, this popup keeps you informed about recent Kentik features and improvements:

* **Article List:** Displays cards for recent updates. Click anywhere on a card to open the full release note in a new tab.
* **Search & Filters:** Use the search bar to find specific updates, or filter by tags (e.g., Improvement, New feature).
* **Subscribe:** Click to receive product updates via email, Slack, or RSS.
* **Feedback:** Use the emoticons (happy, neutral, sad) at the bottom of any article card to submit quick feedback to the product team.

> **Note**: Click on any element except the feedback controls to open the full article in a new tab.

### Portal Subnav

The subnav is the secondary horizontal strip located directly below the navbar. Its contents adapt to the specific page you are viewing. It typically includes:

* **Breadcrumbs**: Indicates your current location within the portal and provides links to navigate up to a higher level.
* **Page-Specific Buttons**: Contextual controls for the active module, such as refreshing data, sharing content, creating a dashboard panel, or adjusting query parameters.

> **Note:** Several module landing pages, such as **Observation Deck** and **Alerting**, do not feature a subnav.

### Portal Main Menu

The portal’s main menu is accessed via the hamburger menu icon at the far left of the navbar. It acts as your primary navigation hub and includes the following areas:

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(29).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A12Z&se=2026-05-12T09%3A33%3A12Z&sr=c&sp=r&sig=ylR7bDFpq2l8B%2FUqOTx6L5xa7kj3FuP%2BjPUFacAdsB0%3D)

*The Kentik portal main menu in its default state.*

* **Featured** (top left): Quick links to commonly used workflows, including **Alerting****,** **Journeys****,** **Kentik Map****,** **Library****,** **My Kentik Portal**, and**Observation Deck**.
* **Portal sections** (center columns): Categorized links to all available portal modules (e.g., Core, Synthetics, Protect). See **Portal Sections**for complete breakdown of these modules.
* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(42).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A12Z&se=2026-05-12T09%3A33%3A12Z&sr=c&sp=r&sig=ylR7bDFpq2l8B%2FUqOTx6L5xa7kj3FuP%2BjPUFacAdsB0%3D)

  **Top Talkers & Settings** (bottom left): Hovering over either of these options displays a contextual submenu. Clicking directly on the links will route you to **Network Explorer**and**Main Settings****,** respectively.
* **Discover** (top right): Helpful internal and external resources, including the Knowledge Base, Product Blog, **Integrations**, a Demo Environment, and the API Tester.
* **Recently Viewed** (bottom right): A quick-access history of your most recently visited portal pages.

  > **Note:** The portal sections list is covered by a popup submenu when you hover on **Top Talkers** or **Settings**.

## Portal Search

The portal’s global search functionality allows you to instantly find network entities, dashboards, and settings. It is accessible from the **Search** field in the main navbar or via **keyboard shortcuts**.

![Search results for devices with navigation and saved filters displayed on the interface.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/PO-portal-search-results.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A12Z&se=2026-05-12T09%3A33%3A12Z&sr=c&sp=r&sig=ylR7bDFpq2l8B%2FUqOTx6L5xa7kj3FuP%2BjPUFacAdsB0%3D)

*The Search Results tab showing hits for Navigation (with Knowledge Base links), Saved Filters, etc, along with an Ask AI Advisor button.*

> **TIP**: The search is incremental, so you don’t need to press return after entering the term.

#### Portal Search Tabs

The top of the **Search** modal features three lozenge-style tabs to help you navigate your data:

| Tab | Description |
| --- | --- |
| **Favorites** | Bookmarked portal views, dashboards, and saved views. To manage your favorites, simply click the star icon located in page title bars, the Library list, or the Recents tab to toggle the designation. |
| **Recents** | A reverse-chronological list of your recently visited dashboards and saved views. Each row displays the view name, an optional description, and when it was last accessed. |
| **Search Results** | Displays automatically as you begin typing. Results are grouped by category (e.g., Interfaces, Dashboards). You can filter visible categories using the Type List on the right, or click See KB results to search the Knowledge Base. |

> **Notes:**
>
> * When you open the Search modal without any text entered, it will default to your **Favorites**. If you have no favorites, it defaults to **Recents**, and if you have no recent history, it defaults to the empty **Search Results** tab.
> * Changes to Favorites in the Search popup are not reflected until the next time it's opened.

## Portal Help and Support

The **Help and Support** popup opens from the question mark icon in the **Portal Navbar**. It has two main panes with information and support:

* **Check our Knowledge Base**: Links to the most helpful KB article on using or understanding the settings and features at your current location (URL).
* **Submit a Support Request**: Fill out a form to report bugs, feature requests, suggestions, or observations. Track your requests in our support portal at **https://support.kentik.com/hc/en-us**.

To submit a request from the popup:

1. In the **Summary** field, briefly state the reason for your support request.
2. Choose the type of support request (question, bug report, training) from the **Type** drop-down.
3. Provide more details in the **Additional Details** field, especially if it's a bug, and include steps to reproduce it.
4. Click **Submit**.

## Portal Keyboard Shortcuts

Listed below are the global Kentik portal shortcuts, which can be used anywhere **except** within the **Search** modal. For search-specific hotkeys, see **Search Keyboard Shortcuts** below.

| Command | Keys | Action Taken |
| --- | --- | --- |
| **Close Drawer** | `esc` | Closes any active side drawer, flyout menu, or open modal. |
| **Open Search** | `cmd` + `/` OR `shift` + `S` | Opens the global Search modal from anywhere in the portal. |
| **Data Explorer** | `shift` + `E` | Navigates directly to the Data Explorer module. |
| **Library** | `shift` + `L` | Navigates directly to the Library module. |
| **Observation Deck** | `shift` + `O` | Navigates directly to the Observation Deck module. |
| **Select Tab** | `Enter` | Activates the currently highlighted tab, option, or search result. |
| **Shortcut Legend** | `shift` + `?` | Opens an on-screen overlay displaying all global keyboard shortcuts. |

### Search Keyboard Shortcuts

The **Search** modal includes several dedicated keyboard shortcuts to help you navigate your results quickly.

| Command | Keys | Action Taken |
| --- | --- | --- |
| **Open Search** | `cmd + /` or `shift + S` | Opens **Search** from anywhere in the portal. |
| **Clear Search** | `esc` | Clears any text in the **Search** field. |
| **Close Search** | `esc` | Exits **Search** if the field is blank. |
| **Switch Tabs** | `Left/Right arrows` | Toggles between Favorites, Recents, and Results tabs. |
| **Navigate Lists** | `Up/Down arrows` | Moves your selection up or down within the active tab’s list. |
| **Go to Page** | `Enter` | Navigates to the page for the selected item. |

> **Legacy Portal**: As of **March 31, 2025**, Kentik’s v3 portal is fully retired and is no longer accessible. For help transitioning to our current portal, see **Customer Care**.
