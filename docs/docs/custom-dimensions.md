# Custom Dimensions

Source: https://kb.kentik.com/docs/custom-dimensions

---

This article discusses **Custom Dimensions** and their associated populators in the Kentik portal.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(656).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A44Z&se=2026-05-12T09%3A50%3A44Z&sr=c&sp=r&sig=%2BX5VoiGAfy%2FFCk5l02qEqEAbMPy4h4HbpxBtaKCSqN4%3D)

*The Custom Dimensions module lets you make your own dimensions for group-by and filtering.*

## Dimensions and Populators

Custom dimensions in Kentik allow customers to add custom columns to the main tables in the Kentik Data Engine (KDE, see **KDE Tables**). These dimensions can be used for group-by in a query (see **Query Sidebar Controls**) or as a filter in a filter group (see **Filter Groups UI**).  
  
 **Creation Options**:

* **Manual**: Users assign “populators” to the dimension.
* **Automatic**: Uses AWS tags to generate populators.

Kentik uses populators, like tags (see **About Flow Tags**), to evaluate flow for matches to a specific value. Tags and populators differ, however, in what happens when a match is found:

* **Tag**: A match places the tag name in the `src_flow_tags` or `dst_flow_tags` column.
* **Populator**: A match places the populator’s `value` in the custom column for the dimension.

> **Note**: Populators require specifying the flow direction (source, destination, either).

#### Auto-Generated Dimensions

Kentik can use AWS or Azure tags to automatically create custom dimension columns in the KDE tables. Each tag key becomes a custom dimension, and the associated values become populators (see **Automatic Custom Dimensions**).

> **Note:** If you update AWS or Azure tags, you must regenerate the populators in Kentik (see **Regenerate Populators**).

#### Dimension and Populator Limits

The table below shows the max dimensions and populators included with your Kentik subscription. The populator limits are customer-wide totals across all custom dimensions. There is no additional per-dimension limit.

| Platform Edition | Max Custom Dimensions | Max Populators | Max Populator Value Length |
| --- | --- | --- | --- |
| Pro | 4 | 5,000 | 128 characters |
| Premium | 10 | 25,000 | 128 characters |

> **Notes**:
>
> * Custom dimensions are not available in Kentik's Essentials edition (see **Kentik Editions**).
> * Additional custom dimensions and populators can be purchased by contacting **support@kentik.com**.
> * Only one populator value can be stored per flow record. If multiple populators match, the most recently added populator value is used.

#### Managing Dimensions and Populators

Custom dimensions and their populators can be managed through the portal (see **Add a Custom Dimension**) or via Kentik V5 APIs (for manually created dimensions, see **About the V5 APIs**):

* **Custom Dimension API**: Manage individual custom dimensions and populators.
* **Batch API**: Manage populators of an existing custom dimension in batches.

#### **Note on Dimension Directionality**

To ensure Kentik correctly detects a custom dimension as directional, the **Dimension Name** must include the pattern `_src_` or `_dst_` (with underscores on both sides).

* **Incorrect:** `c_dest_metro` (Detected as non-directional because it uses `dest` instead of `dst` and is missing the trailing underscore).
* **Correct (Directional):**

  + `c_dst_metro` (Destination)
  + `c_src_metro` (Source)
  + `c_metro_dst` (Pattern mid-name)

## Custom Dimensions Page

The Custom Dimensions page lists all of your organization’s custom dimensions. To view the Custom Dimensions page, choose **Settings** from the main Kentik navbar, then **Custom Dimensions** (under **Data Enrichment**). The Custom Dimensions page is documented in the following topics:

> **Note:** While Member-level users can view this page, only Admin-level users can manage (create, edit, or remove) custom dimensions or populators.

### Custom Dimensions Page UI

The Custom Dimensions page has the following main UI elements:

* **Filter**: A field into which you can enter text to filter the **Custom Dimensions List**. The Dimension Name, Data Format, Display Name, and ID columns of the list are searched for a match on the string entered in this field.
* **Add Custom Dimension**: A button that opens the **Create Custom Dimension** dialog (see **Add a Custom Dimension**).
* **Current Populator Utilization**: A card that indicates your organization’s current usage of populators:

  + **Used**: The number of populators currently in use by your organization.
  + **Total**: The number of populators your organization has purchased.
  + **Remaining**: The populators purchased by your organization but not currently in use.
* **Current Column Utilization**: A card that indicates your organization’s current usage of custom KDE columns:

  + **Used**: The number of columns currently in use by your organization.
  + **Total**: The number of columns your organization has purchased.
  + **Remaining**: The number of columns still available for use by your organization.
* **Custom Dimensions**: A table listing your organization’s currently defined custom dimensions (see **Custom Dimensions List**).

### Custom Dimensions List

The **Custom Dimensions List** is a table that lists all of your organization's existing custom dimensions. Click on a heading to sort the list (ascending or descending).  
  
Each row in the list represents one custom dimension, which corresponds to one column in your organization’s main table in the Kentik Data Engine. The list provides the following information and actions for each custom dimension:

* **Dimension Name**: The name of the custom dimension (starts with "c\_", no spaces allowed).

  > **Note**: To enable directionality detection, the name must contain the `_src_` or `_dst_` pattern (e.g., `c_dst_network`).
* **Populator Type**: The origin of the dimension's populators:

  + **Manually Created**: Created by a user in your organization.
  + **AWS Automatic**: Generated automatically using AWS tags.
* **Data Format**: The custom dimension type, either string or uint32 (unsigned 32-bit integer).
* **Display Name**: The name that will be displayed for the custom dimension in group by and filters UI. Unlike the Dimension Name, the Display Name may contain spaces and is editable.
* **Populators**: The number of populators in the custom dimension.
* **ID**: The system-generated unique ID assigned when the custom dimension was created.
* **Edit** (icon): Opens the **Custom Dimension** dialog where admin users can review and edit settings for the custom dimension.
* **Remove** (icon): Opens a confirmation dialog that allows admin users to remove the custom dimension from the Custom Dimensions List and from the KDE.

## Custom Dimension Dialog

Adding or editing a custom dimension via the Kentik portal involves specifying information in the fields of the custom dimension admin dialogs. The dialog used for these settings depends on what you're trying to do:

* **Create a manually populated dimension**:  
   1. Open the **Create Custom Dimension** dialog from the **Add Custom Dimension** button.  
   2. Choose to populate your custom dimension manually, then click **Next**.  
   3. Fill out the fields to specify settings for the dimension, then Click **Save**.
* **Create an automatically populated dimension** (AWS or Azure only):  
   1. Open the **Create Custom Dimension** dialog from the **Add Custom Dimension** button.  
   2. Choose the AWS or Azure radio button to populate your custom dimension automatically, then click **Next**.  
   3. Specify the AWS or Azure tags that should be used to populate the dimension.
* **Edit the properties of an existing dimension** (manually or automatically populated):  
   1. Click the edit icon in the dimension's row in the **Custom Dimensions List**.  
   2. Edit dimension properties in the **General** tab of the **Custom Dimension Dialog**.
* **Add/edit the populators of an existing dimension** (manually-populated dimensions only):  
   1. Click the edit icon in the dimension's row in the **Custom Dimensions List**.  
   2. In the **Populators** tab of the **Custom Dimension Dialogs**:

  + Add a populator: Click the **Add Populator** button, then use the resulting dialog to set values on which the new populator will match (see **Populator Settings Dialog**).
  + Edit a populator: Click the **Edit** icon for the populator to edit, then use the resulting dialog to change the values on which the new populator will match.

> **Note:** You can only add a populator after you’ve created a dimension.

## Dimension Dialog UI

The UI of the dialogs used for custom dimension settings is covered in the following topics.

> **Notes:**
>
> * Custom dimension admin dialogs are accessible only to users whose level is Administrator.
> * Manually created custom dimensions can also be added and edited with the **Custom Dimension API**.

### Common UI Elements

The following UI elements are common across all dimension dialogs:

* **Close**: The **X** in the upper right corner closes the dialog and restores all elements to their values at the time the dialog was opened.
* **Cancel**: A button that closes the dialog and restores all elements to their values at the time the dialog was opened.

### Create Custom Dimension

The **Create Custom Dimension** dialog is a wizard that appears when you click the **Add Custom Dimension** button on the Custom Dimensions page. The first step of the wizard is to select the type of dimension you'd like to create. Depending on the type, you'll then advance to additional steps in the wizard.

#### Select Dimension Type

The main content of this dialog is a set of radio buttons that you use to choose how you would like to populate the new custom dimension:

* ![Instructions for creating custom dimensions using AWS or Azure tags in Kentik.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/CD-create-custom-dimension-dialog.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A44Z&se=2026-05-12T09%3A50%3A44Z&sr=c&sp=r&sig=%2BX5VoiGAfy%2FFCk5l02qEqEAbMPy4h4HbpxBtaKCSqN4%3D)**Automatically using AWS tags**: Create a custom dimension automatically using AWS tags.
* **Automatically using Azure tags**: Create a custom dimension automatically using Azure tags.
* **Manually**: Manually specify populators for the custom dimension.

After choosing the dimension type, click **Next** to continue to the next step.

#### Create Dimension Automatically

![Selection of AWS entities for creating custom dimensions in a user interface.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/CD-create-custom-dimension-boxes.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A44Z&se=2026-05-12T09%3A50%3A44Z&sr=c&sp=r&sig=%2BX5VoiGAfy%2FFCk5l02qEqEAbMPy4h4HbpxBtaKCSqN4%3D)If you choose to create your custom dimension automatically using AWS or Azure tags (see **Auto-Generated Dimensions**) the next step in the wizard includes the following settings and controls:

* **Entities**: A list of all available AWS or Azure entities from which you can create a custom dimension. Check the checkboxes for the entities you’d like to use.
* **Back**: Return to the previous step.
* **Next**: Continue to the final step in the wizard.

The final step of the wizard includes the following UI elements:

* **Tags**: A list of all available AWS or Azure tag keys from which you can create custom dimension columns. Check the checkboxes for the tags you’d like to use.
* **Back**: Return to the previous step.
* **Save**: Save your settings for the new dimension and exit the dialog.

> **Notes:**
>
> * If your organization is already using the maximum number of custom dimensions (see **Dimension and Populator Limits**), the **Save** button will be disabled.
> * Kentik retrieves AWS metadata at an interval of 15 to 20 minutes, so if you just created the tag, you may need to wait a few minutes for it to appear as an available option in this dialog.
> * It takes approximately 30 minutes to generate the data for the dimensions’ populators after creating the dimension.

#### Create Dimension Manually

![Creating a custom dimension with specified name and display name fields in the interface.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/CD-create-custom-dimension-name.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A44Z&se=2026-05-12T09%3A50%3A44Z&sr=c&sp=r&sig=%2BX5VoiGAfy%2FFCk5l02qEqEAbMPy4h4HbpxBtaKCSqN4%3D)This step appears when you choose to create your custom dimension manually. In addition to the UI elements described in **Common UI Elements**, the **Create Manual Dimension** step contains the following settings and controls:

* **Dimension settings**: Specify the properties of the dimension (see **Dimension General Properties**).
* **Back**: Return to the previous step.
* **Save**: Save your settings for the new dimension and exit the dialog.

### Edit Custom Dimension

To edit an existing custom dimension, you'll use the **Custom Dimension** dialog, which you access from the **Custom Dimensions List** by clicking the **Edit** icon (pencil) on the row of the dimension you’d like to edit.

In addition to the **Common UI Elements**, the **Custom Dimension** dialog features the following settings and controls:

* **Tab selector**: Select which tab is displayed:

  + **General tab** (see **Dimension General Properties**): Contains fields for the properties of a custom dimension.
  + **Populators tab** (see **Dimension Populators Tab**): Contains a list of populators. You can edit or delete populators for manually created custom dimensions from this tab.

    > **Note:** These tabs are shown only after creating a custom dimension and clicking the **Add Custom Dimension** button. When you click to edit the dimension, the two tabs will be visible.
* **Cancel** (button): Click **Cancel** to exit the dialog without saving changes.
* **Save** (button): Click **Save** to save changes to the edited dimension and exit the dialog.

### Dimension General Properties

The general properties of a custom dimension may be specified in the following locations:

* **New dimension**: In the **Create Custom Dimension** dialog.
* **Existing dimension**: On the **General** tab of the **Custom Dimension** dialog.

General settings include the fields in the following table.

| Element | Create Custom Dimension Dialog | Custom Dimension Dialog | Description |
| --- | --- | --- | --- |
| Dimension name | Editable field | Read-only | The name of the custom dimension, which must start with “c\_” and contain only alphanumeric characters and underscores. This is the name that will be used for the column that corresponds to the new dimension in the **KDE Tables**. |
| Type | Editable field | Read-only | The type of the custom dimension, chosen from the drop-down menu: string (default) or uint32. |
| Display name | Editable field | Editable field | The name that will be shown for the dimension in the Kentik portal (e.g., in the **Query Sidebar Controls** or the **Filters Pane**). |

### Dimension Populators Tab

The **Populators** tab of the **Custom Dimension Dialogs** is used to view a dimension's populators. This tab contains the following elements:

* ![Table displaying custom dimension values, directions, and IDs for data management.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/CD-create-custom-dimension-populators.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A44Z&se=2026-05-12T09%3A50%3A44Z&sr=c&sp=r&sig=%2BX5VoiGAfy%2FFCk5l02qEqEAbMPy4h4HbpxBtaKCSqN4%3D)**Filter field**: Enter text to filter the **Populator** **List**. The Value, Direction, and ID columns of the list are searched for a match on the string entered in this field.
* **Add Populator** (present only for manually generated dimensions): Opens the Add Populator dialog (see **Populator Settings Dialog**).
* **Regenerate Populator** (present only for automatically generated dimensions): Regenerates any data gathered from AWS tags. Use this button to generate your populators if anyone in your organization has added or edited any AWS tags since you created that specific dimension.
* **Populator List**: A list of the dimension's populators (see **Populator List**).

### Populator List

The **Populator List** is a table on the **Dimension Populators Tab** that lists all of the populators defined for a given custom dimension. The table provides the following information and actions for each populator:

* **Value**: The value placed in the custom column when there's a match with the ANDed populator parameters (see**Populator Field Definitions**).
* **Direction**: The direction (source, destination, or either) of the flow fields within which the populator will look for a match.
* **ID**: The system-generated unique ID assigned when the populator was created.
* **Edit** (pencil icon; manually-created dimensions only): Opens the Populator dialog to review and edit settings (see **Edit a Populator**).
* **Remove** (trash icon; manually-created dimensions only): Opens a confirmation dialog to remove the populator.

  > **Note:** When a populator for a custom dimension is removed, the KDE main table rows that were previously matched by that populator will no longer be filterable by the populator's value.

## Populator Settings Dialog

> **Note:** This topic applies only to manually-created dimensions.

The Populator dialog is used to add and edit populators in custom dimensions, and is accessed from the **Custom Dimension Dialog**:

* To add a new populator, click the **Add Populator** button.
* To edit an existing populator, click the **Edit** icon in the populator's row of the **Populator List**.

The **Populator** dialog features the following UI elements:

* ![BGP Matching settings with fields for ASN, IP, and community configurations.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/CD-populator-dialog.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A44Z&se=2026-05-12T09%3A50%3A44Z&sr=c&sp=r&sig=%2BX5VoiGAfy%2FFCk5l02qEqEAbMPy4h4HbpxBtaKCSqN4%3D)**Close** (button): Click the **X** in the upper right corner to close the dialog without saving changes.
* **Tabs**: Select a tab to view/edit fields for each category of populators (see **Populator Field Definitions**).
* **Cancel** (button): Click **Cancel** to exit the dialog without saving changes.
* **Add Populator** (button, add only): Click **Add Populator** to save settings for the new populator and exit the dialog.
* **Save** (button, edit only): Click **Save** to save changes to populator settings and exit the dialog.

> **Notes:**
>
> * The Populator dialog is visible only to users whose level is Administrator.
> * Populators can also be added and edited with the **Custom Dimension API** (only applies to manually created dimensions).

## Populator Fields

> **Note:** This topic applies only to manually-created dimensions.

Manually-created populators are defined with a set of fields on the tabs of the Populator dialog.

### Populator Field Definitions

Manually-created populators for a given custom dimension are defined with the fields on the tabs of the Populator dialog. The validation columns in each table below indicate whether the following validations will be applied to a populator field:

* **Comma**: Comma-delimited list
* **Database**: Database patterns (e.g., % and \_)
* **Regex**: Regex

> **Note:** For additional information on validation of populator field values, see **Populator Field Validation**.

#### General Populator Settings

| Field | Description | Comma | Database | Regex |
| --- | --- | --- | --- | --- |
| Value | Required*:* The value placed in the custom column when there's a match with the ANDed populator parameters below: When the custom dimension type is "string":   * Valid characters: alphanumeric, spaces, dashes, underscores. * Length: min=1, max=128.   When the custom dimension type is "uint32":   * Valid values: min=0, max=4294967295. | No | No | No |
| Direction | Required: The direction (source, destination, or either) of the flow fields in which to look for a match with the populator parameters below. | N.A. | N.A. | N.A. |

#### Populator Device Matching

| Field | Description | Comma | Database | Regex |
| --- | --- | --- | --- | --- |
| Device name | Results in a match if this value appears within the name of a device or equals the IP address of a device that has been configured to send flow records to Kentik. | Yes | Yes | No |
| Device Type | Type of device to match (router, host, etc.,  see **Supported Device Types**). | Yes | No | Yes |
| Site | Results in a match if this value appears within the name of a site to which the device sending the flow record to Kentik has been assigned (see **About Sites**). | Yes | No | Yes |
| Interface name/description | Results in a match if this value appears within the name or description of a source or destination interface. | Yes | Yes | Yes |

#### Populator IP Matching

| Field | Description | Comma | Database | Regex |
| --- | --- | --- | --- | --- |
| IP address (IP/CIDR format) | Expressed in IPv4 or IPv6 CIDR notation (e.g., 38.12.34.0/24), this value will result in a match if it corresponds to a range of IP addresses in the flow, either source (SRC IP) or destination (DST IP). | Yes | No | No |
| Port | Results in a match if this value appears within a port number in the flow, either source (SRC Port) or destination (DST Port). | Yes | No | No |
| TCP flag | An integer number between 0 and 255 representing an 8-bit binary bit pattern. At ingest this pattern is used as a bitmask that is ANDed with the composite (ORed) bit pattern of the TCP flags set in the flow. A match will result if the value in both the flow bit pattern and the bitmask is 1 at any of the eight places. | No | No | No |
| Protocol Number | Results in a match if this value is the same as the protocol of the traffic represented by the flow. The protocol of TCP is 6, and of UDP is 17. | Yes | No | No |

#### Populator BGP Matching

| Field | Description | Comma | Database | Regex |
| --- | --- | --- | --- | --- |
| Last-hop (origin) ASN | Results in a match if this value is the same as the last ASN (16- or 32-bit) in the path in the routing table for either the source (SRC IP) or destination (DST IP). | Yes | No | No |
| Last-hop (origin) AS Name | Results in a match if this value represents the name corresponding to the last ASN in the path in the routing table for either the source (SRC IP) or destination (DST IP). | Yes | No | Yes |
| Next-hop ASN | Results in a match if this value is the same as the ASN (16- or 32-bit) of the next hop router based on AS path. | Yes | No | No |
| Next-hop AS Name | Results in a match if this value represents the name corresponding to the ASN of the next hop router based on AS path. | Yes | No | Yes |
| Next-hop IP | If a CIDR grouping (IPv4 or IPv6) is specified, a match can be on any address within that grouping. If no CIDR grouping is specified a match requires an exact IP.   * CIDRs may be expressed in "short form" (e.g., 1::2/127). | Yes | No | No |
| BGP AS Path | Results in a match if this value is the same as the BGP AS path in the route (see **Specifying BGP Fields**). | Yes | No | Yes |
| BGP Community | Results in a match if this value is the same as the BGP community of BGP route. May be specified with a form of regex (see **Specifying BGP Fields**). | Yes | No | Yes |

#### Populator UDR

| Field | Description | Comma | Database | Regex |
| --- | --- | --- | --- | --- |
| Device Type | A dropdown to select a device type for the matching criteria. | N.A. | N.A. | N.A. |
| Various | Fields that vary depending on the selected device type. Results in a match if the values in the fields match to a flow field. | N.A. | N.A. | N.A. |

#### Other Populator Settings

| Field | Description | Comma | Database | Regex |
| --- | --- | --- | --- | --- |
| MAC Address | Results in a match if this value matches source or destination Ethernet (L2) address. | Yes | No | No |
| Country | Results in a match if this value includes a two-letter country code associated with the source or destination IP of the flow. | Yes | No | No |
| VLAN(s) | Results in a match if this value includes a VLAN ID associated with the source or destination IP of the flow. | Yes | No | No |

### Populator Field Validation

The following general considerations apply to the validation of values in the populator fields described in **Populator Field Definitions**:

* Some fields support the entry of multiple values as a comma-delimited list (see tables above).
* Commas are supported only as list delimiters (not in actual values or regex).
* Some fields support the use of database patterns (see tables above).
* In fields where regex is supported (see tables above), a period (".") may be used in place of a comma.
* The BGP AS Path and BGP Community fields use PostgreSQL POSIX Extended Regular Expressions. For additional information, see **Specifying BGP Fields**.
* All other fields that support regex use PostgreSQL Advanced Regular Expressions.

> **Note:** Documentation for PostgreSQL regex and database patterns can be found at **PostgreSQL documentation**:
>
> * Database patterns are documented under "9.7.1. LIKE."
> * Regular expressions are documented under "9.7.3. POSIX Regular Expressions."

### Specifying BGP Fields

Matches on the BGP-related populator fields (see **Populator Field Definitions**) are made on substrings. For ASN and Next Hop ASN, the string(s) to match are specified in a simple comma-delimited list. For both the BGP AS Path and BGP Community fields, the specified values are also evaluated using a subset of standard regex (see table below):

* **BGP AS path**: Entering "10" in the as-path field will match any path that includes "10", "100", "010", etc. Using regex, a value of "\_10\_" will match only paths that include ASN 10, including "10 ", " 10", and " 10 ". Also allowed are tags where as-path is specified as, for example, "\_10 100\_".
* **BGP Community**: Matches on communities are similar to matches on AS paths except that they also support the use of regex periods. This allows you to specify, for example, "2000:1…." to find any flow with community 2000:1xxxx in it.

The following table shows the regex special characters that are supported when specifying the BGP AS Path and BGP Community:

| Special Character | Matches |
| --- | --- |
| \_ (underscore) | start of string end of string  " " (space) |
| . | Any single character, including white space. |
| [ ] | The characters, or a range of characters separated by a hyphen, contained within square brackets. |
| ' | The character or null string at the beginning of an input string. |
| ? | Zero or one occurrence of the pattern containing the question mark. |
| $ | End of string. |
| \* | Zero or more sequences of the preceding character. Also acts as a wildcard for matching any number of characters. |
| + | One or more sequences of the preceding character. |
| () | Used for nesting of expressions. |

> **Note:** For BGP community and AS path, any spaces at the beginning or end of the input field and also before and after each comma will be removed.

## Manage Custom Dimensions

Custom dimensions are added and edited via the Custom Dimensions page.

> **Note:** Manually added custom dimensions can also be added and edited with the **Custom Dimension API**.

### Add a Custom Dimension

A custom dimension may be added either manually or automatically (using previously defined AWS or Azure tags; see **Auto-Generated Dimensions**).

#### Manual Custom Dimensions

In the Kentik portal:

1. Choose **Settings** from the Kentik navbar, then Custom Dimensions under Data Enrichment.
2. Click the **Add Custom Dimension** button at the upper right, which opens the **Create Custom Dimension** dialog.

   > **Note:** If your organization has already reached its limit of custom dimensions (see **Dimension and Populator Limits**), the **Add Custom Dimension** button will be grayed out, and you will not be able to add any custom dimensions without first deleting one or more existing custom dimensions.
3. Select **Manual**, then click **Next**.
4. Enter the Dimension Name, Type, and Display Name for the new custom dimension.

   > **Notes**:
   >
   > * After the custom dimension is created, neither the Dimension Name nor the Type will be editable.
   > * If you intend for this dimension to be directional, follow the `_src_` or `_dst_` naming convention before saving.
5. Click **Save** to save the new dimension and exit. It will now appear as a row in the **Custom Dimensions List**.

The new custom dimension is saved by default without populators, enabling:

* You to create views, dashboards, and alerts around a dimension before defining its populators.
* The custom dimension to function like the standard “FULL:Total” (all traffic) dimension when used as a group-by dimension or filter.

> **Tip**: To add one or more populators to the custom dimension, see **Manage Populators**.

#### Automatic Custom Dimensions

In the Kentik portal:

1. Choose **Settings** from the Kentik navbar, then **Custom Dimensions** (under **Data Enrichment**).
2. Click the **Add Custom Dimension** button at the upper right, which opens the **Create Custom Dimension** dialog.

   > **Note:** If your organization has already reached its limit of custom dimensions (see **Dimension and Populator Limits**), the **Add Custom Dimension** button will be grayed-out, and you will not be able to add any custom dimensions without first deleting existing dimensions.
3. Select **Automatically using AWS tags** or **Automatically using Azure tags**, then click **Next**.
4. Select the AWS or Azure Entity keys from which you'd like to create a custom dimension, then click **Next**.
5. Select the AWS or Azure tags from which you would like to generate custom dimension columns that you can use for group-by and filtering in queries.
6. Click **Save** to close the dialog. The new dimension will appear as a row in the **Custom Dimensions List**. The Kentik system will then begin to generate the populators from your AWS or Azure tags, which may take up to 30 minutes.

### Edit a Custom Dimension

To edit a custom dimension:

1. Choose **Settings** from the Kentik portal navbar, then Data Enrichment » **Custom Dimensions**.
2. Click the Edit icon (pencil) for the custom dimension that you'd like to edit.
3. Edit the custom dimension as follows:

   1. **Change Name**: To change the custom dimension’s name, enter a new name in the **Display Name** field on the **General** tab.
   2. **Add Populator**: To add a populator to the custom dimension, see **Add a Populator**.
   3. **Edit Populator**: To edit an existing populator, see **Edit a Populator**.
   4. **Remove Populator**: To remove a populator from the custom dimension, click **Remove** (trash icon) for that populator in **Populator List**.

      > **Note:** When a populator for a custom dimension is removed, the KDE main table rows that were previously matched by that populator will no longer be filterable by the populator's value.
4. Click **Save** to save changes to the custom dimension and exit.

## Manage Populators

> **Notes:**
>
> * This topic applies only to manually-created dimensions.
> * Populators of manually created custom dimensions can also be added and edited with either the **Custom Dimension API** or the **Batch API**.

Populators of existing custom dimensions can be managed in the **Custom Dimension Dialog**.

### Add a Populator

To add a populator to a manually-created custom dimension using the **Custom Dimension Dialog**, follow these steps:

1. Go to the **Populators** tab of the dialog.
2. Click **Add Populator** at the upper right to open the Populator dialog.
3. Specify values for the populator’s settings (see **Populator Field Definitions**).
4. Click **Add Populator** to add the populator to the custom dimension and return to theCustom Dimension dialog. The added populator is now shown in the **Populator List**.
5. Click **Save** to save changes to the custom dimension and exit.

> **Note:** If you add two or more populators with the same Dimension value (see **Populator Field Definitions**), they’ll be ORed (e.g., if one protocol value is 6 and the other is 17, either 6 or 17 will match). This allows the creation of more complex populators.

### Edit a Populator

To edit an existing populator of a manually-created custom dimension using the **Custom Dimension Dialog**, follow these steps:

1. Go to the **Populators** tab of the dialog.
2. Click the Edit icon for the populator that you wish to edit.
3. Edit the populator fields as needed (see **Populator Field Definitions**).
4. Click **Save** to save changes to the populator and exit to the Custom Dimension dialog.
5. Click **Save** to save changes to the custom dimension and exit.

### Regenerate Populators

Automatically created custom dimensions use AWS or Azure tags to generate populators. Any time you update or create a tag in AWS or Azure, you’ll need to regenerate the populators of any custom dimension that includes populators based on that tag.

To regenerate a dimension's populators:

1. Go to the **Populators** tab of the Custom Dimension dialog.
2. Click **Regenerate Populator**. The dialog will automatically close, and the following confirmation message will appear: “Populators were successfully queued for regeneration. It might take approximately 30 minutes for data to refresh.”
