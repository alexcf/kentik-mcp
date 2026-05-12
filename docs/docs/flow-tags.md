# Flow Tags

Source: https://kb.kentik.com/docs/flow-tags

---

> **Note:** The Flow Tags page is accessible only to Admin users and is hidden from Member users.

Kentik allows the use of user-defined flow tags to simplify querying in the Kentik Data Engine (KDE).

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(665).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A16Z&se=2026-05-12T09%3A41%3A16Z&sr=c&sp=r&sig=ZyN1eJkcs2q%2FW9xhYOqtAqjxpIXw9u0J1ah%2F6qMcE4c%3D)

*The Flow Tags page lists the tags you've set to match flow record fields at ingest.*

## About Flow Tags

Flow tags are labels applied to flow data based on user-defined criteria as Kentik ingests the data into the Kentik Data Engine (KDE). Tags must be created before they can be applied to  flow records, as they cannot be applied retroactively. However, since Kentik stores complete flow records, historical queries are not limited to pre-defined flow attributes. Tags enhance querying options without limiting or filtering the stored data.  
  
Flow tags can be created in the portal (see **Adding a Flow Tag**) or via Kentik V5 APIs (see **About the V5 APIs**):

* **Individual Tags**: Use the **Tag API**.
* **Batch Tags**: Use the **Batch API**.

#### Tag Application at Ingest

When a flow record is sent to Kentik, it is evaluated against existing tags as follows:

* **Source Fields**: If all values in the ANDed tag fields match SRC-related flow fields (e.g., SRC IP, SRC port), the tag's name is appended to the `src_flow_tags` column.
* **Destination Fields**: If all values in the ANDed tag fields match DST-related flow fields (e.g., DST IP, DST port), the tag's name is appended to the `dst_flow_tags` column.

> **Notes:**
>
> * Tags are applied only when all specified fields are matched.
> * The `src_flow_tags` and `dst_flow_tags` columns of each device's main table contain a delimited list of tags, which can be searched as part of a KDE query.

## Flow Tags Page

The Flow Tags page contains the **Flow Tag List**, a table displaying all existing tags for your organization. To access this page, navigate to **Settings** in the main navbar, then select **Flow Tags** under **Data Enrichment**.

### Flow Tags Page UI

The Flow Tags page has the following main UI elements:

* **Filter field**: Enter text to filter the Tag Name column of the **Flow Tag List** for a match on the string entered in this field.
* **Add Tag button**: Opens the **Add Tag** dialog (see **Adding a Flow Tag**).
* **Flow Tag List**: A table listing the flow tags currently set up in your organization (see **Flow Tag List**).

### Flow Tag List

The **Flow Tag List** is a table that lists all of your organization’s currently defined flow tags. The table's columns provide the following information and actions for each tag:

* **Tag Name**: The name of the flow tag as specified at the time the flow tag was created. Click anywhere on a tag’s row to open the **Edit Tag** dialog for that flow tag.
* **Last Edited**: The date on which the tag was last edited. Hover over the date for a tool tip with the full date-time (UTC).
* **Edited by**: The email address of the person who last modified the tag.
* **Created**: The date on which the tag was created. Hover over the date to see the full date-time (UTC).
* **ID**: The system-generated unique ID assigned to the tag when it was created.
* **View in Chart**: Opens the **Total Matching Traffic Dialog** for either the Source or the Destination.
* **Remove** (trash icon)*:* Opens a confirmation dialog that allows you to remove the tag from the Kentik portal.

> **Notes:**
>
> * Click on the column headings for **Tag Name** or **ID** to sort the list (ascending or descending).
> * To see additional information about a given flow tag, click anywhere in the row for that flow tag, which opens an **Edit Tag** dialog where you can review settings (see **Editing a Flow Tag**).

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(666).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A16Z&se=2026-05-12T09%3A41%3A16Z&sr=c&sp=r&sig=ZyN1eJkcs2q%2FW9xhYOqtAqjxpIXw9u0J1ah%2F6qMcE4c%3D)

#### Total Matching Traffic Dialog

The **Total Matching Traffic** dialog is opened via one of the **View in Chart** buttons, either **Source** or **Dest**. The dialog displays a chart showing the total traffic (expressed as max bits/second), both historically and for the last 24 hours, that had matches for this tag in either source flow (**Source** button) or destination flow (**Dest** button).

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(667).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A16Z&se=2026-05-12T09%3A41%3A16Z&sr=c&sp=r&sig=ZyN1eJkcs2q%2FW9xhYOqtAqjxpIXw9u0J1ah%2F6qMcE4c%3D)

The dialog includes the following UI elements:

* **Close buttons**: To close the dialog, click the **X** in the upper right corner or the **Close** button at lower right.
* **Applied Filters**: A button that opens a popup in which you can see which filters are applied to the traffic shown in the chart.
* **View Type**: A drop-down menu used to set the type of visualization used for the graph (defaults to Line Chart). For descriptions of the options, see **Chart Visualization Types**.
* **Chart**: The visualization of traffic (using the current view type). Hover over any part of the chart for more precise information.
* **View in Explorer button**: Opens Data Explorer for further exploration of the device's traffic. The query settings will be set to show the same traffic that is shown in the **Total Matching Traffic** dialog.

## Tag Admin Dialogs

Adding or editing a flow tag via the Kentik portal involves specifying information in the fields of the tag admin dialogs.

> **Note:** Tags can also be added and edited with either the **Tag API** or the **Batch API**.

### About Tag Dialogs

The Kentik portal uses tag admin dialogs to collect and display flow tag information. The required information is entered into the fields of either of the following dialogs:

* Add Tag when adding a new tag to Kentik.
* Edit Tag when editing an existing tag.

![Instructions for specifying a tag in a tagging system interface.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/FT-edit-tags-dialog.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A16Z&se=2026-05-12T09%3A41%3A16Z&sr=c&sp=r&sig=ZyN1eJkcs2q%2FW9xhYOqtAqjxpIXw9u0J1ah%2F6qMcE4c%3D)

### Tag Dialogs UI

The **Add Tag** and **Edit Tag** dialogs share the same layout and the following common UI elements:

* **Close button**: Click the **X** in the upper right corner to close the dialog. All elements will be restored to their values at the time the dialog was opened.
* **Tab selectors**: Choose the tab to display: **General**, **Device Matching**, **IP Matching**, **BGP Matching**, or **Other** (see tab-specific topics below).
* **Cancel**: A button that cancels the add tag or edit tag operation and exits the dialog. All elements will be restored to their values at the time the dialog was opened.
* **Add Tag** (**Add Tag** dialog only): A button that saves settings for the new tag and exits the dialog.
* **Save** (**Edit Tag** dialog only): A button that saves changes to tag settings and exits the dialog.

## Tag Field Definitions

The fields of the tag admin dialogs are used to specify and display flow tag settings.

### About Tag Fields

Tag admin dialogs are divided into tabs, each containing several fields. To create a flow tag, specify a name in the Tag Name field and additional tag attribute fields to match incoming flow  fields. The fields for defining tags are detailed below.  
  
These columns indicate applicable validations for each tag field:

* **Comma**: Comma-delimited list
* **Database**: Database patterns (e.g., % and \_)
* **Regex**: Regex

> **Note:** For more, see **Tag Field Validation**.

### General Tag Settings

The following table shows the settings on the **General** tab of tag admin dialogs:

| Field | Description | Comma | Database | Regex |
| --- | --- | --- | --- | --- |
| Tag Name | The string that will be added to the src\_flow\_tags and/or dst\_flow\_tags column in the main KDE table of the device sending flow when a match is found in the flow for the values in any of the following fields.  **Notes:**   * A tag name must be from 2 to 20 characters long: alphanumeric, hyphen, or underscore (no spaces). * A tag name must be unique, but tags whose names contain a common string can be ORed in a query (e.g. "tag1" and "tag2" both contain "tag"). | No | No | No |

### Tag Device Matching

The following table shows the settings on the **Device Matching** tab of tag admin dialogs:

| Field | Description | Comma | Database | Regex |
| --- | --- | --- | --- | --- |
| Device Name | Results in a match if this value appears within the name or equals the IP address of a device that has been configured to send flow records to Kentik. If there's a match, the tag is applied to both src\_flow\_tags and dst\_flow\_tags columns. | Yes | Yes | No |
| Device Type | Type of device to match (router, host, etc.; see **Supported Device Types**). | Yes | No | Yes |
| Site | Results in a match if this value appears within the name of a site to which the device sending the flow record to Kentik has been assigned (see **About Sites**). | Yes | No | Yes |
| Interface Name | Results in a match if this value appears within the name or description of a source or destination interface. If there's a match, the tag is applied to the src\_flow\_tags column if the received flow shows traffic entering on the interface, and a tag is applied to the dst\_flow\_tags column if the received flow shows traffic leaving on the interface. | Yes | Yes | Yes |

### Tag IP Matching

The following table shows the settings on the **IP Matching** tab of tag admin dialogs:

| Field | Description | Comma | Database | Regex |
| --- | --- | --- | --- | --- |
| IP address (IP/CIDR format) | Expressed in IPv4 or IPv6 CIDR notation (e.g., 38.12.34.0/24), this value will result in a match if it corresponds to a range of IP addresses in the flow, either source (SRC IP) or destination (DST IP). If there's a match, the tag is applied to both src\_flow\_tags and dst\_flow\_tags columns.  **Note**: This field can contain up to 249 IP/CIDR items in a comma-delimited list. | Yes | No | No |
| Port | Results in a match if this value appears within a port number in the flow, either source (SRC Port) or destination (DST Port). | Yes | No | No |
| TCP Flag | An integer number between 0 and 255 representing an 8-bit binary bit pattern. At ingest this pattern is used as a bitmask that is ANDed with the composite (ORed) bit pattern of the TCP flags set in the flow. A match will result if the value in both the flow bit pattern and the bitmask is 1 at any of the eight places. | No | No | No |
| Protocol Number | Results in a match if this value is the same as the protocol of the traffic represented by the flow. The protocol of TCP is 6, and of UDP is 17. | Yes | No | No |

### Tag BGP Matching

The following table shows the settings on the **BGP Matching** tab of tag admin dialogs:

| Field | Description | Comma | Database | Regex |
| --- | --- | --- | --- | --- |
| Last-hop (origin) ASN | Results in a match if this value is the same as the last ASN (16- or 32-bit) in the path in the routing table for either the source (SRC IP) or destination (DST IP). | Yes | No | No |
| Last-hop (origin) AS Name | Results in a match if this value represents the name corresponding to the last ASN in the path in the routing table for either the source (SRC IP) or destination (DST IP). | Yes | No | Yes |
| Next-hop ASN | Results in a match if this value is the same as the ASN (16- or 32-bit) of the next hop router based on AS path. | Yes | No | No |
| Next-hop AS Name | Results in a match if this value represents the name corresponding to the ASN of the next hop router based on AS path. | Yes | No | Yes |
| Next-hop IP | If a CIDR grouping (IPv4 or IPv6) is specified, a match can be on any address within that grouping. If no CIDR grouping is specified a match requires an exact IP.   * CIDRs may be expressed in "short form" (e.g., 1::2/127). | Yes | No | No |
| BGP AS Path | Results in a match if this value is the same as the BGP AS path in the route (see **Specifying BGP Tag Fields**). | Yes | No | Yes |
| BGP Community | Results in a match if this value is the same as the BGP community of BGP route. May be specified with a form of regex (see **Specifying BGP Tag Fields**). | Yes | No | Yes |

#### Specifying BGP Tag Fields

Kentik’s BGP data collection enables tagging of incoming flows based on matching communities and AS paths or partial paths. Tags are applied separately for source and destination during flow ingestion into KDE and can be used to refine query results in Data Explorer.

> **Note**: Tags apply only to flows arriving after their creation and may take up to 20 minutes to activate.

#### Matching BGP-related Tag Fields

The following considerations apply to matching BGP-related tag fields:

* **ASN and Next-hop ASN**: Use a comma-delimited list for matching strings.
* **BGP AS Path Tags**:

  + Enter "10" to match paths containing "10", "100", "010", etc.
  + Use regex "\_10\_" to match paths with ASN 10, such as "10 ", " 10", and " 10 ".
  + Example: "\_10 100\_" for specific path matching.
* **BGP Community Tags**:

  + Similar to AS path tags but support regex periods.
  + Example: "2000:1...." matches any flow with community 2000:1xxxx.

The following table shows the regex special characters that are supported when specifying the BGP AS Path and BGP Community:

| Special Character | Matches… |
| --- | --- |
| \_ (underscore) | * Start of string * End of string * " " (space) |
| . | Any single character, including white space. |
| [ ] | The characters, or a range of characters separated by a hyphen, contained within square brackets. |
| ' | The character or null string at the beginning of an input string. |
| ? | Zero or one occurrence of the pattern containing the question mark. |
| $ | End of string. |
| \* | Zero or more sequences of the preceding character. Also acts as a wildcard for matching any number of characters. |
| + | One or more sequences of the preceding character. |
| () | Used for nesting of expressions. |

> **Note:** For BGP community and AS path tags, any spaces at the beginning or end of the input field, and also before and after each comma will be removed.

### Other Tag Settings

The following table shows the settings on the **Other** tab of tag admin dialogs:

| Field | Description | Comma | Database | Regex |
| --- | --- | --- | --- | --- |
| MAC Address | Results in a match if this value matches source or destination Ethernet (L2) address. | Yes | No | No |
| Country | Results in a match if this value includes a two-letter country code associated with the source or destination IP of the flow. | Yes | No | No |
| VLAN(s) | Results in a match if this value includes a VLAN ID associated with the source or destination IP of the flow. | Yes | No | No |

> **Note:** The validation codes used in the table above are defined in **Tag Field Validation**.

### Tag Field Validation

The following considerations apply to the validation of the tag fields described in **Tag Field Definitions**:

* Fields may accept multiple values as a comma-delimited list (see tables above).
* Commas are allowed only as list delimiters, not part of values or regex.
* Some fields allow database patterns (see tables above).
* In regex-supported fields, a period (".") can replace a comma.
* BGP AS Path and BGP Community tags use PostgreSQL POSIX Extended Regex (see **Specifying BGP Tag Fields**).
* Other regex-supporting tags use PostgreSQL Advanced Regex.

> **Note:** **PostgreSQL Documentation** covers database patterns under "LIKE" and regex under "POSIX Regular Expressions."

## Add or Edit Tags

Flow tags are created and edited via the **Flow Tags Page** of the Kentik portal (from the main navbar, select Settings and then Flow Tags).

> **Note:** Tags can also be added and edited with either the **Tag API** or the **Batch API**.

### Adding a Flow Tag

To add a new tag:

1. Go to the Flow Tags page (Settings » **Flow Tags**).
2. Open the **Add Tag** dialog by clicking the **Add Tag** button.
3. On the **General** tab, name the tag in the Tag Name field.
4. Specify the values of the tag fields that will be evaluated for a match with the properties of the incoming flow (see **Tag Field Definitions**).
5. Click the **Add Tag** button (bottom right) to save the new tag with the currently specified values and close the dialog.

> **Note:** A tag name must be unique, but tags whose names contain a common string can be ORed in a query (e.g., "tag1" and "tag2" both contain "tag").

### Editing a Flow Tag

To edit an existing tag:

1. Go to the Flow Tags page (Settings » **Flow Tags**).
2. In the **Flow Tag List**, open an **Edit Tag** dialog by clicking anywhere in the row of the tag that you'd like to edit.
3. Edit the tag fields that you'd like to change (see **Tag Field Definitions**).
4. To save changes, click the **Save** button at bottom right. The dialog will close.
