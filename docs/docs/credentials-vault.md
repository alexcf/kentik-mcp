# Credentials Vault

Source: https://kb.kentik.com/docs/credentials-vault

---

This article discusses the **Credentials Vault** page in the Kentik portal.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(134).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A51Z&se=2026-05-12T09%3A44%3A51Z&sr=c&sp=r&sig=vJW5i7SMr6yVKjL27OkXgvbmnpuBidm0bmblf2yA7cc%3D)

*The Credentials Vault page features a filterable list of credentials.*

Kentik’s **Credentials Vault** enables you to save credentials centrally and securely and use them in different configuration pages throughout the portal.

> **TIP**: Updating a credential in the Credentials Vault updates it everywhere else it is used.

|  |  |
| --- | --- |
| **Purpose:** | Provide a secure place to centralize credentials for Kentik Administrators. |
| **Benefits:** | * Central management of credentials. * Secure storage of your network credentials and access to them as needed on the Kentik portal. * Share your network credentials in the portal without revealing their values. * Credentials are double encrypted at rest using AES256 - one unique customer key and one global key. |
| **Use Cases:** | * Update your credential values in the Credentials Vault to automatically update your credential wherever in the portal it is used. * Keep your credentials secure if another user needs to employ them within the portal. |
| **Relevant Roles:** | Network Engineers, NetOps Engineers, Kentik Administrators |

## About Credentials

When discussing credentials within the Kentik portal, we refer to four primary terms:

* **Key**: The name for the key/value pair, which must be unique to that credential.
* **Value**: The key’s value, e.g., email address, password (see **Credential Values**).
* **Key/value pair**: A single key with a single corresponding value. A value is accessed by referencing the key.
* **Credential**: A collection of key/value pairs, which has an associated name, description, type (and optionally, labels).

### Credential Security

Credentials stored in the Credentials Vault undergo double encryption using AES256. Credential values are write-only: once a credential is saved, values can be only be reset, not viewed nor disclosed.

## Credential Portal Usage

Credentials are currently used in the portal for Kentik-registered devices and HTTP synthetic tests in the **Test Control Center**.

### Credentials for Devices

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(136).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A51Z&se=2026-05-12T09%3A44%3A51Z&sr=c&sp=r&sig=vJW5i7SMr6yVKjL27OkXgvbmnpuBidm0bmblf2yA7cc%3D)For Devices, credentials can be used as follows:

1. Check a box on one of these tabs of the device settings dialog:

   1. **SNMP**: Under **Collection method**, check the **For NMS via Universal Agent** checkbox (see **Device SNMP Settings**).
   2. **Streaming Telemetry**: Under **Collection** method, check **Dial-In from Universal Agent** (see **Device ST Settings**).
2. The **Credential** drop-down appears, which includes all your credentials for the selected collection method.
3. Select an existing credential or click **New Credential** to create a new credential for that device (see **Add SNMP Device Credential** and **Add ST Device Credential**).

### Credentials for HTTP Tests

To use your Credentials in Synthetics for HTTP tests, click **Credentials Vault** on one of the following tabs of the **Test Settings Page** (see **Set HTTP Test Credentials**):

* **Transaction**: **Target** and **Agents** tabs.
* **Page Load**: **HTTP** tab, when the **Headers** tab is selected in the Configure Request section.
* **HTTP(S) or API**: **HTTP** tab, when the **Headers** and/or **Body** tabs are selected in the Configure Request section.

  > **Notes**:
  >
  > + For a step-by-step walk-through of creating a credential and using it with an HTTP(S) or API test, see **API Test Credentials Tutorial**.
  > + To determine which synthetic tests use specific credentials, use the **Credentials** filter in the **Tests List Filters** on the Test Control Center page.

## Credentials Vault Page

The **Credentials Vault** page is home to the **Credential List**, which displays all of the credentials to which you have access in the Kentik portal. To access the page, choose **Credentials Vault** from the **Organization Settings** menu on the portal's main navbar.

> **IMPORTANT:** Only Administrator-level users can add or modify credentials (see **About Users**). Member-level users can only access the Credentials Vault page and use credentials (assuming adequate RBAC permissions).

### Credentials Vault Page UI

The Credentials Vault page includes the following UI elements:

* **Add Credential** (Administrator-level users only): Opens the Add Credential dialog to create a new credential (see **Credential Settings Dialogs**).
* **Show/Hide Filters**: A filter icon to expand/collapse the **Filters** pane.
* **Group by**: A drop-down to group the credentials on the list:

  + **None**: Credentials not grouped.
  + **Type**: Credentials grouped by type, e.g., Basic Authentication, SNMP etc.
  + **Usage**: Credentials grouped by current Kentik portal usage, e.g. used vs. unused.
* **Search**: Enter search terms to filter by matches in the **Type** or **Name** columns. Click the **X** at the right of the field to clear entered text.

  + The field also shows any filter lozenges set in the **Filters** pane.

    - Click the **X** in a lozenge to clear the corresponding filter.
* **Filters**: Pane of controls to filter the **Credential List** (see **Credential List Filters**).
* **Credential list**: A table listing the credentials in your organization to which you have access (see **Credential List**).

### Credential List Filters

The **Filters** pane of the **Credential List** narrows the list to credentials matching the specified criteria. When a filter is added with this pane, a lozenge for the filter appears in the **Search** field (see **Credentials Vault Page UI**). It includes the following controls:

* **Reset To Default** (appears only when you’ve specified one or more filters): Clears all current filters.
* **Type**: Checkboxes that filter the list of credentials by **Credential Type**.
* **Labels**: Displays a lozenge for each selected label, which is used to filter the Credentials list (see **Labels**).

  + **Select a Label**: Click in the field, then choose a label from the filterable list.
  + **Remove a Label**: Click the **X** at the right of that label's lozenge.

    > **Note**: Only labels currently applied to one or more credentials are displayed in the list.
* **Created by**: Checkboxes to filter the credentials list by the user who created them.

### Credential List

The **Credential** list shows all credentials for your organization. To sort the list, click a column header, then choose the sort direction: ascending or descending.

The list includes the following columns:

* **Name**: The credential’s name with an optional description, and optional applied labels.
* **Type**: The credential type (see **Credential Type**).
* **Keys**: The key names for this credential, displayed alphabetically.
* **Used By**: Lists the ways this credential is used in the portal (e.g., in a Synthetics test):

  + **When in Use**: Click the link to navigate to where in the portal it is used (e.g., Devices page or the Test Control Center), where the list is pre-filtered for only elements using that credential.
  + **When Not in Use**: This column reads “Not In Use”.
* **Created By**: The user name that created the credential.
* **Last Updated**: The last updated date and time for the credential.
* **Edit** (Administrator-level users only): A pencil icon that opens the Edit Credential dialog (see **Credential Settings Dialogs**).
* **Remove** (Administrator-level users only): A trash icon that removes the credential from your organization (requires confirmation).

## Credential Settings Dialogs

Add or edit a credential by specifying information in the fields of a **Add** **Credential** or **Edit Credential** settings dialog.

> **Note:** Only Administrator-level users can view the credential settings dialogs.

![Form for adding credentials with fields for name, description, type, key, and value.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/CV-add-credential-dialog.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A51Z&se=2026-05-12T09%3A44%3A51Z&sr=c&sp=r&sig=vJW5i7SMr6yVKjL27OkXgvbmnpuBidm0bmblf2yA7cc%3D)

### Credential Settings UI

The credential settings dialogs include:

* **Cancel**: Click the **X** at the upper right or **Cancel** at the lower right to exits without saving.
* **Add Credential** (Add Credential dialog only): Click **Add Credential** to save the credential and exit.
* **Save** (Edit Credential dialog only): Click **Save** to save the credential and exit.

### Credential Settings Definitions

The credential settings dialogs include the following:

| Element | Add Credential Dialog | Edit Credential Dialog | Description |
| --- | --- | --- | --- |
| Name | Editable field | Non-editable field | The credential's name. Can only use alpha-numeric characters, dashes, and underscores. |
| Description | Editable field | Editable field | A description of the credential. |
| Type | Field (drop-down list) | Field (drop-down list) | The type of credential selected. To select a type, click the field and choose from the drop-down list (see **Credential Type**). |
| Labels | Field | Field | Displays a lozenge for each selected label. To select a label, click the field and choose from the drop-down list. To remove a label, click the **X** at the right of the lozenge (see **Credential Labels**). |
| Add/Edit Label | Link | Link | Takes you to the **Labels** page to create or remove the labels available for credentials. |
| Key - Unstructured (Type is Basic Authentication, API Token, or Other) | Editable field | - Non-editable field (saved keys) - Editable (new keys) | The key of the key/value pair. Can only use alpha-numeric characters, dashes, and underscores (see **Credential Type**). |
| Key - System template (Type is SNMP v1 or v2 Community String, SNMP v3 Authentication, Streaming Telemetry Authentication, or BGP MD5) | Non-editable field | Non-editable field | The key of a system template’s key/value pair. Cannot change the key name or add key/value pairs (see **Credential Type**). |
| Value | Editable field | Editable field | The value of the key/value pair. Once saved, the field’s contents can be reset, but cannot be unmasked. Can only use alpha-numeric characters, dashes, and underscores. See **Credential Values**. |
| Eye icon | Button | Button | A button that appears in the Value field that unmasks the characters in the field. Only shown for unsaved values in editable fields. |
| Add | Button | Button | Adds a key/value pair to the credential. Only appears with unstructured keys. |
| Remove | Button | Button | Removes the key/value pair from the credential. Only appears when a set of key/value pairs is added and hasn’t been saved yet. |

### Credential Labels

Applying labels to your credentials provides a way to easily find and filter them within both the **Credential** list and the **Credentials Vault Popup** (when creating a Synthetics test). Here are some considerations:

* The **Labels** field in the credential settings dialogs displays a lozenge for each selected label.
* **To Apply a Label to a Credential**: Click the field and choose a label from the filterable drop-down list.
* **To Remove a Label from a Credential**: Click the **X** at the right of the label’s lozenge.

> **TIP:** Labels can be linked to RBAC roles to grant permissions for accessing and modifying data throughout the portal (see **Labels**).

### Credential Type

The credential type is assigned when it’s created or edited, and has an associated key type:

* **System template**: Pre-set system templates where key counts and names cannot be changed (i.e. they appear grayed out in the **Credential Settings Dialogs**).
* **Unstructured**: Customizable key counts and names.

The following table shows all available credential/key type combinations:

| Credential Type | Key Type | Key Name | Value | Additional Key/Value Pairs? |
| --- | --- | --- | --- | --- |
| Basic Authentication | Unstructured | User-defined | User-defined | Yes |
| SNMP v1 Community String | System template | `community` | User-defined | No |
| SNMP v2 Community String | System template | `community` | User-defined | No |
| SNMP v3 Authentication | System template | `username` | User-defined | No |
| SNMP v3 Authentication | System template | `authentication` | User-defined | No |
| SNMP v3 Authentication | System template | `authentication_protocol` | Drop-down menu:  - MD5  - SHA | No |
| SNMP v3 Authentication | System template | `privacy` | User-defined | No |
| SNMP v3 Authentication | System template | `privacy_protocol` | Drop-down menu:  - DES  - AES | No |
| Streaming Telemetry Authentication | System template | `username` | User-defined | No |
| Streaming Telemetry Authentication | System template | `password` | User-defined | No |
| BGP MD5 | System template | `peer_asn` | User-defined | No |
| BGP MD5 | System template | `md5_password` | User-defined | No |
| API Token | Unstructured | User-defined | User-defined | Yes |
| Other | Unstructured | User-defined | User-defined | Yes |

> **TIPS**:
>
> * A user-defined key name can contain only alphanumeric characters, underscores, and dashes.
> * If you change a credential’s type, recheck the value of each of its keys.

**Keys for Unstructured Credentials**

When naming keys/values for unstructured credentials, here are some considerations:

* Only alphanumeric characters, dashes, or hyphens are allowed. This allows the credentials to be used in HTTP Transaction tests within a Javascript/Puppeteer script.
* Each key name must be unique, and be sure to use names that are self-explanatory and descriptive.
* Once a credential is saved, the value of any **Key** is no longer editable, however you can still add key/value pairs.

> **Note:** The key/value pairs of a saved credential, when listed in the **Edit Credential** dialog or in a **Credentials Vault Popup** in a Synthetics test, are arranged alphabetically by key name, not in the order they were added.    

### Credential Values

A credential value is a value (e.g., email address, password) provided to the system when its paired key is requested on the portal.

* When you first enter text into a **Value** field of a credential settings dialog, you can click the eye icon to view the entered text.
* Each value is write-only: once a credential is saved, the Value field can no longer be viewed.
* Each saved **Value** field shows asterisks to represent the entered value.
* When a credential’s value is copied from the **Credentials Vault Popup**, a programmatic representation (or variable) is used to maintain the value’s security (see **Value Variables**).
* You can reset a **Value** (**Update a Credential Value**) or delete the credential entirely (**Remove a Credential**).

  + The same visibility conditions apply upon resetting a value: you can view it until saved, after which its value is masked.

> **Notes**:
>
> * The only exceptions to the value-masking described above are with the **authentication\_protocol** and **privacy\_protocol** keys for the credential type SNMP v3 Authentication. The options for these two fields are visible even after the credential is saved.
> * Administrator-level users can create, edit, or delete a credential, whether that user created the credential or not.

#### Value Variables

When creating a Synthetics test involving a credential, the credential's settings can be specified with strings that you've copied from the **Credentials Vault Popup**. When pasted into a **Key** field, the text will display as a programmatic representation (variable) of the actual value. The structure of the variable string will be `$vault(‘credential_name.key_name')`, where:

* **$vault**: A vault function that calls the set specified in the following parentheses.
* **()**: The contents of the parentheses show what is being called from the Credentials Vault.
* **credential\_name**: The credential’s name.
* **key\_name**: The credential’s key name.

## Manage Credentials

The management of credentials is covered in the following sections.

### Add a Credential

To add a new credential, visit one of these portal locations:

1. **Credentials Vault Page**
2. **SNMP** tab of a device settings dialog
3. **Streaming Telemetry** tab of a device settings dialog

> **Notes**:
>
> * The **Credential Name**field accepts only alpha-numeric characters, dashes, and underscores.
> * Credentials created in the settings dialog for a device can only be edited or removed on the Credentials Vault page.
> * To add labels or descriptions to credentials created in the settings dialog for a device, edit the credential from the Credentials Vault page (see **Edit a Credential**).
> * A newly created credential will immediately be listed in the **Credential List**.

#### Add Credential on CV Page

To add a new credential from the Credentials Vault (CV) page:

1. Navigate to **Credentials Vault** from the **Organization Settings** menu on the portal's main navbar.
2. Click **Add Credential** to open the Add Credential dialog.
3. Enter the necessary values (see **Credential Settings Definitions**).
4. Click **Add Credential** to save your changes. Your credential now appears in the **Credential** list.

> **Notes**:
>
> * When adding a key/value pair, the eye icon in the **Value** field reveals the entered value.
> * Once a credential is saved, the contents of its **Value** fields are masked and cannot be revealed.

#### Add SNMP Device Credential

To add a new credential for SNMP when adding or editing a device (see **Device SNMP Settings**):

1. Navigate to Settings » **Devices** and click **Add Device**.

   1. To edit, click the **Edit** (pencil icon) at the right of a row in the Device list.
2. Click the **SNMP** tab.
3. Check the **For NMS via Universal Agent** checkbox.
4. From the Credential drop-down, choose **New Credential** to open the **Add SNMP Credential Dialog**.
5. Use the **Type** radio buttons to choose an SNMP version (v1, v2, or v3).
6. Enter a value in the following fields:

   * Credential Name
   * Community (for v1 and v2)
   * User Name, Authentication type and value, and  Privacy type and value (for v3).
7. Click **Add Credential** to save the credential.

   > **Note**: To add a Streaming Telemetry credential, see **Add ST Device Credential**.
8. Click either **Add Device** (when adding) or **Save** (when editing) to save the device and exit.

#### Add ST Device Credential

To add a new credential for Streaming Telemetry (ST) when adding or editing a device (see **Device ST Settings**):

1. Navigate to Settings » **Devices** and click **Add Device**

   1. To edit, click the Edit button (pencil icon) at the right of a row in the Device list.
2. Click the **Streaming Telemetry** tab of the dialog.
3. Check the **Dial-In from Universal Agent** checkbox.
4. From the **Credential** drop-down, choose **New Credential** to open the Add Streaming Telemetry Credential dialog.
5. Enter values in the following fields:

   1. Credential Name
   2. User Name
   3. Password
6. Click **Add Credential**to save the credential and exit. The new credential will be selected in the Credential field.
7. Click either **Add Device** (when adding) or **Save** (when editing) to save the device and exit.

   > **Note**: To add a credential for SNMP, see **Add SNMP Device Credential**.

### Edit a Credential

To edit the settings for an existing credential:

1. Open the **Credentials Vault** page from the **Organization Settings** menu on the portal's main navbar.
2. Find the credential you wish to edit, and click **Edit** to open the Edit Credential dialog.
3. Edit the credential settings as desired (see **Credential Settings Definitions**).

   > **TIP**: If you change a credential’s type, recheck the value of each of its keys.
4. Click **Add** to add a new key/value pair to the credential (see **Add a Key/Value Pair**).
5. Click **Save** to save changes and exit.

> **Note**: Once a credential is saved, its key/value pairs cannot be removed and all **Key** fields are read-only. Key values, however, can still be edited (see **Update a Credential Value**).

#### Add a Key/Value Pair

Add a new key/value pair to an **Unstructured** credential (see **Keys by Credential Type**) by following these steps:

1. Click **Add** to create a new key/value pair in the **Key** list.
2. Enter a value in the **Key** field.

   > **IMPORTANT**: After saving, you cannot edit the **Key** value.
3. To delete the key/value pair before saving, click the Remove button (trash icon).

   > **IMPORTANT**: After saving, you cannot remove a key/value pair from a credential.
4. Click **Save** to save the new key/value pair.

> **Notes**:
>
> * After saving, when you open the **Edit Credential** dialog, the key/value pair list is auto-arranged alphabetically by **Key** name.
> * **System Template** credentials cannot be modified.

#### Update a Credential Value

To update a key's value for a credential:

1. On the portal's main navbar, navigate to Organization Settings » **Credentials Vault**.
2. Find the credential in the **Credential List**, and click **Edit** to open the Edit Credential dialog.
3. Enter a value in the **Value** field. Only alpha-numeric characters, dashes, and underscores are allowed.

   > **TIP**: Click the eye icon to check the contents of the **Value** field before saving.
4. Click **Save** to update the value.

   > **IMPORTANT**: Credential values are update everywhere in the portal that the credential is used.

### Use a Credential

In the Kentik portal, you can use credentials in two distinct contexts:

* **Devices**: Credentials for authorizing SNMP polling of a device and/or the publication of Streaming Telemetry data from a device (see **Credentials for Devices**):

  + For device SNMP polling, see **Add SNMP Device Credential**.
  + For Streaming Telemetry from a device, see **Add ST Device Credential**.
* **Synthetics Tests**: Credentials for authenticating HTTP tests (see **Credentials for HTTP Tests**).

  + To set a credential for an HTTP test, see **Set HTTP Test Credentials**.

    > **Note:** The minimum synthetics agent version for credentials support is 0.6.10.

#### Set HTTP Test Credentials

Credentials can be used in Synthetics for authentication in tests from the HTTP category:

1. On the Synthetics » **Test Control Center** page, click **Add Test** at top right.
2. In the **Application** section, click one of the following to open a **Test Settings Page** for the new test:

   1. HTTP(S) or API
   2. Page Load
   3. Transaction
3. On the settings page, click **Credentials Vault** to open the **Credentials Vault Popup**. The button’s presence depends on the test type:

   1. **HTTP(S)/API and Page Load**: On the **HTTP** tab, only when the tab of the **Configure Request** table is either **Headers** or **Body**.
   2. **Transaction**: On the **Target** and **Agents** tab.
4. Filter the list by credential name or any assigned labels. The matching credential(s) are then listed under **Credentials** in the Credentials Vault popup.

   > **Note**: Filtering by credential name is case sensitive.
5. Click the arrow by the credential’s name to expand it to show the key names (listed under **Values**).
6. Click the Copy button beside a key name (e.g., `auth_email_header`) to copy its value.
7. Paste it into the **Key** field of the Headers tab of the Configure Request table.

   > **Note**: The value will be displayed as a programmatic variable (see **Value Variables**).
8. Repeat steps 6 to 7 for any other values that you need to add to the test (e.g., `auth_email_value`, `auth_api_token_header`, and `auth_api_token_value` (see **API Test Credentials Tutorial**).
9. Click the **X** to close the **Credentials Vault** popup.
10. Complete any remaining settings on the other **Test Settings Tabs**.
11. Ensure the test is error-free. Preview it if needed (see **Test Preview Overview**).
12. Click **Create Test**.

**Edit Credentials for an HTTP Test**

1. Open the test’s Edit Test dialog (see **Edit a Test**).
2. Complete steps 3 to 11 from **Set HTTP Test Credentials**.
3. Click **Save**.

### Remove a Credential

To remove a credential from your organization’s list of credentials:

1. Open the Credentials Vault page from the **Organization Settings** menu on the portal's main navbar.
2. In the **Credential List**, find the row corresponding to the credential you’d like to remove, then click the **Remove** button at the right of the row.
3. Click **Remove** in the resulting confirmation dialog.

## API Test Credentials Tutorial

This section covers the use of credentials in Synthetics. This involves:

1. Creating an example credential with four key-value pairs.
2. Using the credential's values in an **HTTP(S)** or **API** test against our **Label APIs**.

### Finding Header Elements

We need these four things to create the credential to authenticate with the API:

* Header name for the user's email address
* Header value for the email address
* Header name for the user's API token
* Header value for the API token

To get the four values, follow these steps:

1. In the Kentik portal, navigate from the main navbar menu to the API Tester and click **Label Management API** in the API list.
2. On the Label Management API page, select the **Swagger** tab.
3. In the **LabelService** pane, click the row for the **GET** method to expand it.
4. In the **Parameters** section, click **Try it Out**.
5. Click **Execute** to execute the API call and open the **Responses** section.

   ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(139).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A51Z&se=2026-05-12T09%3A44%3A51Z&sr=c&sp=r&sig=vJW5i7SMr6yVKjL27OkXgvbmnpuBidm0bmblf2yA7cc%3D)
6. In the **Curl** section of **Responses,** note the following header names and values:

   1. **Email header name**: `X-CH-Auth-Email`
   2. **Email header value**: `r.song@tardis.time (placeholder for actual address)`
   3. **API token header name**: `X-CH-Auth-API-Token`
   4. **API token header value**: `718bc4cc02e10139b46a4e31efa2face (placeholder for actual token)`

### Creating a Credential

To create a new credential:

1. From the Credentials Vault page, click **Add Credential** to open the Add Credential dialog.
2. ![Credential setup interface displaying email and API token fields for user authentication.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/CV-add-credential-dialog-filled.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A51Z&se=2026-05-12T09%3A44%3A51Z&sr=c&sp=r&sig=vJW5i7SMr6yVKjL27OkXgvbmnpuBidm0bmblf2yA7cc%3D)In the **Name** field, enter a name using only alpha-numeric characters, dashes, and underscores, e.g., “TimeLord Credentials".
3. In the **Description** field, enter a description for the credential.
4. Select a **Type,** for example "Other".
5. Add the appropriate labels to easily find our credentials, e.g., when creating a synthetic test or filtering the **Credential List**.
6. Click the **Add** button three times to create a total of four **Key**/**Value** pairs for the credential.
7. In each **Value** field, paste the relevant element obtained in **Finding Header Elements**. Click the eye icon to see the entered values to confirm their accuracy.

   > **Note**: After saving, values are masked for all user types. While you can reset the value, you’ll no longer be able to view it.
8. In the **Key** field to the left of each **Value** field, enter the following names:

   1. `auth_email_header`
   2. `auth_email_value`
   3. `auth_api_token_header`
   4. `auth_api_token_value`

      > **Note**: After saving, your list of key-value pairs will auto-arrange alphabetically by **Key**, so be sure to use a self-explanatory and/or descriptive name.
9. Click **Add Credential** to save your credential.

### Creating a Test

To create a synthetic test against Kentik’s API using Kentik’s credential:

1. On the Synthetics » **Test Control Center** page, click the **Add Test** button at top right.
2. In the **Application** section of the page, click **HTTP(S)** or **API** to open a **Test Settings Page** for the new test.
3. On the settings page, fill in the settings of the **Test Information** and **Target and Agents** tabs.

   > **Note**: Credentials are only supported for agents running version 0.6.10 or higher.
4. ![Credentials vault displaying TimeLord_Credentials with various authentication values and notes.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/CV-vault-popup.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A51Z&se=2026-05-12T09%3A44%3A51Z&sr=c&sp=r&sig=vJW5i7SMr6yVKjL27OkXgvbmnpuBidm0bmblf2yA7cc%3D)On the **HTTP** tab, click the **Headers** tab in the **Configure Request** pane to show pairs of  **Key** and **Value** fields, the first of which we'll use for the `auth_email_header`.
5. Click **Add Headers** to add a second key/value pair for the `auth_api_token_header`.
6. Click **Credentials Vault** to open the **Credentials Vault Popup** to show the credentials to which you currently have access.
7. Use the **Filter by name** field (case sensitive) or the **Find by labels** drop-down to find the credential added in **Creating a Credential**.
8. Click the arrow next to the credential’s name to expand it. The names of the credential's four keys will be listed under **Values**.
9. Click the Copy button next to `auth_email_header` to copy the value to your clipboard.
10. On the **Headers** tab in the **Configure Request** pane, paste the value into the **Key** field of the first key/value pair. The field will display the following (see **Value Variables**): `$vault('TimeLord_Credentials.auth_email_header')`
11. Back in the **Credentials Vault** popup, click Copy beside `auth_email_value`.
12. On the **Headers** tab in the **Configure Request** pane, paste the value into the **Value** field of the first key/value pair.
13. Repeat steps 9 to 12 to specify the next key/value pair using `auth_api_token_header` (**Key** field) and `auth_api_token_value` (**Value** field).

    ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(142).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A51Z&se=2026-05-12T09%3A44%3A51Z&sr=c&sp=r&sig=vJW5i7SMr6yVKjL27OkXgvbmnpuBidm0bmblf2yA7cc%3D)
14. Click the **X** to close the **Credentials Vault** popup.
15. Complete the remaining settings for your test (**Test Settings Tabs**).
16. Click **Preview** to open the **Test Preview Page** to confirm the test is set up correctly. To check your credit usage estimate, go to either of these places:

    1. **Test Preview Page**, top right.
    2. **Test Settings Page**, bottom left.
17. Click **Create Test** to save the test.
