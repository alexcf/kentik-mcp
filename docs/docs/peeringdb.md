# PeeringDB

Source: https://kb.kentik.com/docs/peeringdb

---

This article covers Kentik’s PeeringDB integration.

![PeeringDB interface showing facilities and exchanges for network mapping.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/PD-integrate-pd.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A59Z&se=2026-05-12T09%3A48%3A59Z&sr=c&sp=r&sig=cYmyfYQfZLCHfkv62XUCIEJB%2F5VfRJNwoE8S8EsvyPk%3D)

*Kentik's PeeringDB integration is configured in a dialog opened from the Integrations page.*

## About PeeringDB Integration

**PeeringDB** is a global database of networks (ASes) that are available for interconnection at Internet Exchange Points (IXPs), data centers, and other interconnection facilities. Originally set up to facilitate peering between networks and peering coordinators, this free, open, user-maintained service now includes all types of interconnection data for networks, clouds, services, enterprise, and interconnection facilities.  
  
 An AS’s footprint encompasses its physical data center locations, which we refer to in Kentik as "sites" (PeeringDB calls them "facilities"), as well as any Internet Exchanges (IXes) in which the AS's interfaces are present. When you declare your AS's footprint in PeeringDB, you’re publicly stating the points at which you’re willing to interconnect (peer) with other networks. In each AS’s PeeringDB record you can see the following information:

* General information about the network (traffic levels, ratio of inbound to outbound, geographic scope, protocols, etc.).
* The network's interconnection policy, which is a general statement of the conditions under which they are open to peering.
* Public peering exchange points and Interconnection facilities, which are the locations where you can peer.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(559).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A59Z&se=2026-05-12T09%3A48%3A59Z&sr=c&sp=r&sig=cYmyfYQfZLCHfkv62XUCIEJB%2F5VfRJNwoE8S8EsvyPk%3D)

*A typical record for an AS in PeeringDB.*

The information in PeeringDB enables Kentik to help you find other networks whose footprints overlap with your own. You can map your own network footprint, check another organization’s network footprint, and determine points of overlap in your footprint with those of other ASes, allowing you to evaluate potential candidates for peering with the goal of reducing your transit costs.  
  
As a part of Kentik’s PeeringDB integration, for each AS’s Details page in Network Explorer we have added a **PeeringDB Info Tab** that provides a wealth of information about the AS's footprint and how (or if) it intersects with yours.

> **Note:** Kentik's PeeringDB data is updated daily.

## Integrate PeeringDB Dialog

The Integrate PeeringDB dialog is covered in the following topics.

### Integrate PeeringDB Overview

The Integrate PeeringDB dialog opens from the **PeeringDB** tile under **Network Platforms** on the Integrations page (Discover »**Integrations** in the portal's main navbar menu). The dialog shows the mapping between sites and interfaces that you've configured in Kentik and PeeringDB-registered interconnection facilities and peering exchanges. It also enables you to manually change those mappings.

> **Note:** You can also map to facilities and exchanges from a site or interface in an AS that your organization hasn't registered with PeeringDB, but you do that from the settings for the site or interface (see **Map from Sites List** or **Map from Interfaces List**).

### Integrate PeeringDB UI

TheIntegrate PeeringDB dialog includes the following settings and controls:

* **Close**: Click the **X** in the upper right corner to close the dialog without saving changes to the PeeringDB settings.
* **PeeringDB ASN lookup**: A field that is populated with lozenges for the ASNs to look up in PeeringDB (see **PeeringDB ASN Lookup**).
* **PeeringDB mapping table**: A list correlating PeeringDB facilities and exchanges (listed on the left) with your own sites and interfaces (see **PeeringDB Mapping Table**).
* **Remove All Mappings**: A button with which you can remove all of the mappings in the **Assigned** column of the **Custom-Facilities** and **Custom-Exchanges** sections of the mappings table (see **Remove All Mappings**).
* **Cancel**: Close the dialog without saving changes.
* **Save**: Save changes to the PeeringDB settings and exit the dialog.

### PeeringDB ASN Lookup

When you open theIntegrate PeeringDB dialog, the **PeeringDB ASN Lookup** field is populated with lozenges for any ASNs entered in the **Internal ASNs** field on your **Network Classification** page. PeeringDB is checked to see if any records are associated with these ASNs. If such records exist they will be listed in the **To be mapped** column of the **PeeringDB Mapping Table**:

* The facilities shown in the **Interconnection Facilities** section of the PeeringDB record will be listed under the **Facilities** heading row.
* The exchanges shown in the **Public Peering Exchange** Points section of the PeeringDB record will be listed under the **Exchanges** heading row.

The **PeeringDB ASN Lookup** field may be blank if:

* The **Internal ASNs** field is blank on your Network Classification page.
* None of your internal ASNs are associated with a record in PeeringDB.

If the **Internal ASNs** field on your Network Classification page is blank:

* Instead of the mapping table, theIntegrate PeeringDB dialog will display a message asking you to set up Network Classification.
* You will be unable to map any sites or interfaces to PeeringDB facilities or exchanges.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(561).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A59Z&se=2026-05-12T09%3A48%3A59Z&sr=c&sp=r&sig=cYmyfYQfZLCHfkv62XUCIEJB%2F5VfRJNwoE8S8EsvyPk%3D)

*An alert appears in place of the table when you haven't specified any internal ASNs in Network Classification.*

## PeeringDB Mapping Table

The**PeeringDB Mapping**table is covered in the following topics.

### Mapping Table Overview

The **PeeringDB Mapping** table in the Integrate PeeringDB dialog details the facilities and exchanges that have been mapped — either automatically or manually — to your own sites and interfaces. The manual creation of "custom" mappings in the dialog makes it possible for you to identify footprint overlaps with networks that have a PeeringDB record when your own network does not.

### Mapping Columns

The PeeringDB Mapping table has two main columns:

* **To be mapped**: A list of the facilities and exchanges found on PeeringDB that can be mapped to your Kentik-monitored sites and Interfaces.
* **Assigned**: The sites (for facilities) and interfaces (for exchanges) that are mapped to the facility or exchange in the **To be Mapped** column. If there's no mapping, the column will display a blank field (for sites) or a **Select Interfaces** button (for interfaces).

  > **Note**: The parenthetic information beside the Assigned column head tells you how many Kentik sites and interfaces are assigned to the listed PeeringDB facilities and exchanges (“9 out of 15”, for example, means that 9 of the possible 15 mappings are complete).

### Group Heading Rows

The rows of the table are divided by heading rows into the following groups:

* **Facilities**: The PeeringDB facilities at which your organization has a footprint (see **PeeringDB Facilities**), and the sites to which they are mapped.
* **Exchanges**: The PeeringDB-registered exchanges at which your organization has interfaces (see **PeeringDB Exchanges**). An exchange can span over several facilities.
* **Custom - Facilities**: The Peering DB facilities to which you have mapped any Kentik sites (see **Custom Facilities**).
* **Custom - Exchanges**: The Peering DB exchanges to which you have mapped any Kentik interfaces (see **Custom Exchanges**).

The heading row for each of the above groups includes the following:

* A triangle at the left that you can click to expand or collapse the group below.
* A number showing the count of facilities or exchanges in that group.
* A heading in the Assigned column indicating whether the assigned elements are sites or interfaces.

> **Note:** For more information about each group, see **Mapping Table Groups**.

### Facilities and Exchanges Rows

The individual (non-heading) rows within the Facilities and Exchanges groups include:

* **Name**: The name of the facility or exchange.
* **Location**: The physical location (address) of the facility or exchange.
* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(562).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A59Z&se=2026-05-12T09%3A48%3A59Z&sr=c&sp=r&sig=cYmyfYQfZLCHfkv62XUCIEJB%2F5VfRJNwoE8S8EsvyPk%3D)**Sites field** (facilities rows only): A field containing a lozenge for each site that has been mapped to the facility. If no sites are mapped, the field will be empty. To add a site, click in the field to open a filterable drop-down from which you can choose sites.
* **Interfaces button** (exchanges rows only): ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(564).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A59Z&se=2026-05-12T09%3A48%3A59Z&sr=c&sp=r&sig=cYmyfYQfZLCHfkv62XUCIEJB%2F5VfRJNwoE8S8EsvyPk%3D)A button that opens the **Select Exchange Interfaces** dialog:

  + If no interfaces have been mapped to this exchange, the button is labeled "Select Interfaces."
  + When one or more interfaces have been mapped, the button is labeled with a count of the mapped interfaces. Hover over the button for details about the interfaces, including name, description, capacity, connectivity type, and network boundary.

The individual (non-heading) rows within the **Custom - Facilities** and**Custom - Exchanges** groups are identical to the rows of the **Facilities** and **Exchanges** rows except for the following differences:

* **Custom - Facilities rows**: The lozenges for sites are not within a field and there's no drop-down from which to select sites.
* **Custom - Exchanges rows**: The interfaces are presented as a lozenge (labeled with the number of interfaces) and there's no button from which to open the **Select Exchange Interfaces** dialog. Hover over the button for details about the interfaces.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(566).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A59Z&se=2026-05-12T09%3A48%3A59Z&sr=c&sp=r&sig=cYmyfYQfZLCHfkv62XUCIEJB%2F5VfRJNwoE8S8EsvyPk%3D)

*A Custom - Exchange lozenge with informational popup.*

The above differences stem from the fact that the sites and interfaces for the table's custom groups are defined on the **Sites** and **Interfaces** pages respectively, and are not alterable within the Integrate PeeringDB dialog. For information on PeeringDB settings for sites and interfaces, see **Map Site to Facility** and **Map Interface to Exchange**.

## Mapping Table Groups

The following topics explain how the groups of the **PeeringDB Mapping** table in theIntegrate PeeringDB dialog are populated.

### PeeringDB Facilities

Kentik queries PeeringDB for Interconnection Facilities that are associated with the ASes operated by your organization, as specified in the **Internal ASNs** field on the **Network Classification** page. Each found facility is listed in a row of the **Facilities** group of the **PeeringDB Mapping Table**. Each row shows the Kentik-monitored site(s) that map to the facility, as well as its name and location. If none of your sites in Kentik are mapped to a PeeringDB-registered facility, the Site field in the facility's table row will be blank.

> **Note:** The sites mapped to a facility can be manually added or changed as described in **Map Site to Facility**.

### PeeringDB Exchanges

Kentik queries PeeringDB for Internet Exchanges that are associated with the ASes operated by your organization, as specified in the **Internal ASNs** field on the **Network Classification** page. Each found exchange is listed in a row of the **Exchanges** group of the **PeeringDB Mapping Table**. Each row includes an **Interfaces** button (see **Facilities and Exchanges Rows**) that is labeled either "Select interfaces" (if no interfaces are mapped) or with the number of mapped interfaces (hover for details).

> **Note:** The interfaces mapped to an exchange can be manually added or changed by clicking the **Interfaces** button to open the Select Exchange Interfaces dialog as described in **Map Interface In Mapping Table**.

### Custom Facilities

If your organization operates an AS that isn't registered in PeeringDB then the facilities used for that AS won't appear in your **PeeringDB Mapping Table**. You can still include the AS in your organization's network footprint by manually creating a custom mapping to a PeeringDB facility from a Kentik site (see **Map Site to Facility**). The manually mapped facility will be listed in the **Custom - Facilities**group as a row that shows its physical address and the Kentik sites mapped to it.

> **Note**: The mappings listed in the**Custom - Facilities**group are not editable in theIntegrate PeeringDB dialog. To edit or remove them, you need to modify the site settings (see **Remove a Site Mapping**).

### Custom Exchanges

If your organization operates an AS that isn't registered in PeeringDB then the exchanges in which you use interfaces for that AS won't appear in your **PeeringDB Mapping Table**. You can still include the AS in your organization's network footprint by manually creating a custom mapping to a PeeringDB exchange from a Kentik-monitored interface (see **Map Interface to Exchange**). The manually mapped exchange will be listed in the **Custom - Exchanges** group as a row that shows its location and an **Interfaces** button (see **Facilities and Exchanges Rows**) that is labeled either "Select interfaces" (if no interfaces are mapped) or with the number of mapped interfaces (hover for details).

> **Note**: The mappings listed in the**Custom - Exchanges** group are not editable on the **Integrate PeeringDB Dialog**. To edit or remove them, you need to modify the interface settings (see **Remove an Interface Mapping**).

## Select Exchange Interfaces

The **Select Exchange Interfaces**dialog is used to map Kentik-monitored interfaces to a PeeringDB-registered exchange. The dialog opens when you click an **Interfaces** button in the **Facilities and Exchanges Rows** of the mapping table in theIntegrate PeeringDB dialog. It is covered in the following topics:

![Interface selection screen displaying various connectivity options and selected interfaces for review.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/PD-select-exchange-interfaces.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A59Z&se=2026-05-12T09%3A48%3A59Z&sr=c&sp=r&sig=cYmyfYQfZLCHfkv62XUCIEJB%2F5VfRJNwoE8S8EsvyPk%3D)

*A dialog used to map interfaces to exchanges.*

### Select Exchange Interfaces UI

The Select Exchange Interfaces dialog includes the following main UI elements:

* **Close button**: Click the **X** in the upper right corner to exit the dialog leaving the **Selected Interface** list as it when the dialog was opened.
* **Information pane**: This pane provides you with steps for how to use the **Select Exchange Interfaces** dialog. To close it, click the **X** on the right side of the pane.
* **Filters**: A set of fields, drop-downs, radio buttons, and checkboxes that define filter criteria and are used to narrow the set of interfaces that are available to select in the **Matching Interfaces** list (see **Interfaces Filters Pane**).
* **Group By**: A set of buttons that determine how the interfaces in the Matching Interface list will be organized into groups, either by Device, Connectivity Type, or Network Boundary (see **Matching Interfaces List**).
* **Matching Interfaces** (center pane): A list of the interfaces that match all criteria currently defined in the **Filters** pane. If no filters are set, the list will include the interfaces on all of your Kentik-registered devices. Select interfaces from this list with the + button to populate the **Selected Interfaces** list (see **Matching Interfaces List**).
* **Selected Interfaces** (right pane): A list of the interfaces you’ve added from the Matching Interface list (see **Selected Interfaces List**).
* **Cancel**: Exit the dialog, leaving the **Selected Interface** list as it was when the dialog was opened.
* **Save**: Save changes to the **Selected Interfaces** list, exit the dialog, and return to the **Integrate PeeringDB Dialog**.

### Interfaces Filters Pane

The **Filters** pane of the Select Exchange Interfaces dialog includes the following types of controls:

* **Reset To Default**: Click to clear all current filters.
* **Text box**: Enter a full or partial text string by which you’d like to filter your interfaces.
* **Selection fields**: Click to drop down a list from which you can choose an item as a filter criteria, which will appear as a lozenge in the field. Repeat to add more items. To remove an item, click the X at the right of its lozenge.
* **Checkboxes**: Click checkboxes in a given category to choose one or more filter criteria in that category. To remove, uncheck the box.
* **Radio buttons**: Click only one of the available options.

> **Notes:**
>
> * The **Filters** pane shows only controls (checkboxes and selectors) for which there is a match with the interfaces shown in your **Interfaces List**. As the list is filtered, the number and type of controls displayed in the **Filters** pane may change.
> * The easiest way to find IX interfaces in order to map them to a PeeringDB exchange is to set the **Connectivity Type** filter to IX.

#### Filter Categories

Filter criteria for interfaces fall into the following categories:

* **Interface Name** (text field): Match interfaces whose name contains the entered.  Regular are expressions allowed.
* **Interface Description** (text field): Match interfaces whose description contains the entered text (case insensitive). Regular are expressions allowed.
* **IP** (text field): Match interfaces with the entered IP/CIDR address range.
* **Device** (selection field): Match interfaces to the device selected from the drop-down menu.
* **Device Name** (text field): Match interfaces whose device name contains the entered string. Regular are expressions allowed.
* **Device Label** (selection field): Match interfaces to the device label selected from the drop-down menu.
* **Site** (selection field): Match interfaces to the site selected from the drop-down menu.
* **Capacity** (selection field): Match interfaces whose capacities match the entered value.
* **Interface Type** (radio buttons): Select whether interfaces are to match **All Interfaces**, only **Logical Interfaces** (interfaces that have been assigned an IP address), or only **Physical Interfaces**.
* **Provider** (selection field): Match interfaces whose source/destination traffic reaches the Internet via the entered provider.
* **Connectivity Type** (checkboxes): Match interfaces whose Connectivity Type is checked (see **Understanding Connectivity Types**).
* **Network Boundary** (checkboxes): Match interfaces whose Network Boundary is checked (see **Network Boundary Attribute**).

You can use the above filter controls to narrow the set of available interfaces. The **Matching Interfaces** list will update to show only the interfaces that match the current filters.

#### Filter Application Rules

Kentik applies the following rules to filter categories and criteria:

* Criteria are ORed (match any) within categories and ANDed (match all) between categories.
* When criteria are selected in more than one category, an interface is displayed only if it matches at least one of the selected criteria in all of the categories with criteria selected.
* When no criteria are selected for a category, the interface may match any of the category’s listed criteria.

### Filtered Interfaces Pane

The**Filtered Interfaces** pane (center pane) includes the following elements:

* **Group By**: A set of buttons from which you can choose how to group the interfaces in the **Matching Interfaces List**:

  + Device: Group the interfaces by the device on which they exist**.**
  + Connectivity Type: Group the interfaces by their connectivity type (see **Understanding Connectivity Types**).
  + Network Boundary: Group the interfaces by their network boundary attribute (see **Network Boundary Attribute**).
* **Matching Interfaces**: A list that is populated based on the current filters in the **Interfaces Filters Pane** and grouped according to the Group By settings (see **Matching Interfaces List**).
* **Interface count**: Below the Matching Interfaces list, a statement will inform you how many interfaces are listed (to a maximum of 1000) and how many interfaces match the criteria you’ve selected in the Filters pane.

#### Matching Interfaces List

The **Matching Interfaces** list is populated based on the current filters in the **Interfaces Filters Pane** and grouped according to the Group By setting.  
  
The list includes the following columns and controls:

* **Interface**: In each row, the top line is the interface’s SNMP alias and the bottom line (if present) is the interface’s description (see **Information SNMP OIDs**).
* **Information**: Gray indicators showing the interface’s capacity, connectivity type, and network boundary (see **Interface Classification Dimensions**).
* **Add** (+): An icon at the right of each individual row. Click to add the corresponding interface to the **Selected Interfaces List**. The icon will change to a green circled checkmark.

### Selected Interfaces List

The**Selected Interfaces** list (right pane) displays any interfaces you’ve selected from the **Matching Interfaces List** (center pane). These are the interfaces that will be mapped to the PeeringDB exchange whose Interfaces button you clicked in the mapping table.  
  
To remove an interface from the **Selected Interfaces** list, click the red**X** displayed to its right. Click the**Save** button to map the selected interfaces to the PeeringDB exchange and return to the **Integrate PeeringDB Dialog**.

## Manage PeeringDB Mapping

The Kentik portal offers multiple ways to manage the mappings between the sites and interfaces in your network and the facilities and exchanges registered in PeeringDB.

### Map Site to Facility

There are three ways to map a Kentik site to a PeeringDB facility:

* Using the **Integrate PeeringDB** dialog (see**Map Site in Mapping Table**).

  > **Note:** This method is available only for facilities that appear in the record of ASes that you have registered in PeeringDB, which will be listed in the **PeeringDB Facilities** group of the **PeeringDB Mapping Table**.
* Using multi-row selection in the **Sites List** (see **Map from Sites List**).
* Using the settings dialog for a new or existing site on the **Sites Page** (see **Map In Site Settings**).

> **Note**: The group in which a facility that you manually map via the Sites page (the second and third methods above) will appear in the **PeeringDB Mapping Table** depends on whether your ASes are registered with PeeringDB: Facilities group if yes, **Custom - Facilities** if no.

#### Map Site in Mapping Table

To map a site to a PeeringDB facility in the **Integrate PeeringDB Dialog**:

1. Choose Integrations from the main menu (under Discover).
2. Click **PeeringDB** (under **Network Platforms**) to open the **Integrate PeeringDB** dialog.
3. In the **Facilities** group of the **PeeringDB Mapping Table**, find the row for the facility to which you'd like to map sites.
4. Click the **Site** field at the right of the row.
5. In the resulting drop-down, enter the name of the site in the filter field.
6. In the filtered list, click to select the desired site. The drop-down will close.
7. Repeat steps 4 to 6 as needed to add more sites.
8. Click **Save** to save your site mappings, close the dialog, and return to the **Integrations** page.

#### Map from Sites List

To map one or more sites to a single PeeringDB facility using the **Sites List**:

1. On the main portal menu, choose **Settings** to open the Settings page.
2. Click **Manage Sites** (under **Network Metadata**).
3. In the **Sites** list, click the checkbox for each site that you’d like to map to a specific PeeringDB facility.
4. From the controls that appear above the **Sites** list (see **Selected Site Actions**), click the **PeeringDB Facility** button.  ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(568).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A59Z&se=2026-05-12T09%3A48%3A59Z&sr=c&sp=r&sig=cYmyfYQfZLCHfkv62XUCIEJB%2F5VfRJNwoE8S8EsvyPk%3D)
5. In the resulting drop-down, enter the name of the facility in the filter field.
6. In the filtered list, click to select the desired facility.
7. In the **Confirm PeeringDB Facility Mapping** dialog, click **Confirm**. The dialog will close. In the **Sites** list, the facility should now be listed for the selected sites in the **PeeringDB Facility** column.

#### Map In Site Settings

You can map an individual site to one or more facilities using a **Site Settings** dialog. The following example illustrates making the setting for an existing site in theEdit Site dialog, but you can also set it when adding a new site with the Add Site dialog:

1. On the main portal menu, choose **Settings** to open the Settings page.
2. Click **Manage Sites** (under **Network Metadata**).
3. In the **Sites** list, click the Edit icon in the row for the site you’d like to map. This will open the **Edit Site** dialog.
4. Click in the **PeeringDB Facility Mapping** field.
5. In the resulting drop-down, enter the name of the facility in the filter field.
6. In the filtered list, click to select the desired facility.
7. Repeat steps 4 to 6 as needed to select more facilities.
8. Click **Save** to save your mappings. You’ll be returned to the **Sites** page.

### Map Interface to Exchange

There are three ways to map a Kentik interface to a PeeringDB exchange:

* Using the **Select Exchange Interfaces** dialog from the mapping table in the Integrate PeeringDB dialog (see **Map Interface In Mapping Table**).

  > **Note:** This method is available only for exchanges that appear in the record of ASes that you have registered in PeeringDB, which will be listed in the **PeeringDB Exchanges** group of the **PeeringDB Mapping Table**.
* Using multi-row selection in the **Interfaces List** (see **Map from Interfaces List**).
* Using the settings dialog for a new or existing interface on the **Interfaces Page** (see **Map In Interface Settings**).

> **Note**: The group in which a manually mapped exchange will appear in the **PeeringDB Mapping Table** depends on whether your ASes are registered with **PeeringDB:** **Exchanges** group if yes, **Custom - Exchanges**if no.

#### Map Interface In Mapping Table

To map an interface to a PeeringDB exchange in the **Integrate PeeringDB Dialog**:

1. Choose **Integrations** from the main menu (under Discover).
2. Click **PeeringDB** (under Network Platforms) to open the Integrate PeeringDB dialog.
3. In the **Exchanges** group of the **PeeringDB Mapping Table**, find the row for the exchange to which you'd like to map interfaces.
4. Click the **Interfaces** button at the right of the row.
5. In the resulting **Select Exchange Interfaces** dialog, populate the Selected Interfaces list with the interfaces that you'd like to map to the exchange.

   > **Note**: The easiest way to find IX interfaces to map a PeeringDB exchange is to set the **Connectivity Type** filter to IX.
6. Click **Save** to exit the Select Exchange Interfaces dialog and return to the Integrate PeeringDB dialog.
7. Click **Save** to exit the dialog and return to the **Integrations** page.

#### Map from Interfaces List

To map one or more interfaces to a single PeeringDB exchange using the **Interfaces List**:

1. On the main portal menu, choose **Settings** to open the Settings page.
2. Click **Manage Interfaces** (under **Network Metadata**).
3. In the **Interfaces List**, click the checkbox for each interface that you’d like to map to a specific PeeringDB exchange.
4. From the controls that appear above the **Interfaces List** (see **Interfaces Page UI**), click the PeeringDB IX button.
5. In the resulting drop-down, enter the name of the exchange in the filter field.
6. In the filtered list, click to select the desired exchange.
7. In the Confirm PeeringDB Exchange Mapping dialog, click **Confirm**. The dialog will close. In the Interfaces list, the exchange should now be listed for the selected interfaces in the **PeeringDB IX** column.

   > **Note**: If the PeeringDB IX column isn’t shown in the list, configure the table's columns via the Customize button.

#### Map In Interface Settings

You can map an individual interface to an exchange using an **Interface Settings Dialog**. The following example illustrates making the setting for an existing interface in the Edit Interface dialog, but you can also set it when adding a new interface with the Add Interfacedialog:

1. On the main portal menu, choose **Settings** to open the Settings page.![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(570).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A59Z&se=2026-05-12T09%3A48%3A59Z&sr=c&sp=r&sig=cYmyfYQfZLCHfkv62XUCIEJB%2F5VfRJNwoE8S8EsvyPk%3D)
2. Click **Manage Interfaces** (under **Network Metadata**).
3. In the **Interfaces** list, click the **Edit** icon in the row for the interface you’d like to map. This will open the **Edit interface** dialog.
4. Click the **PeeringDB Internet Exchange** button.
5. In the resulting drop-down, enter the name of the exchange in the filter field.
6. In the filtered list, click to select the desired exchange.
7. Click **Save** to exit the dialog and return to the **Interfaces** page.

> **Note**: The group in which a manually mapped exchange will appear in the **PeeringDB Mapping Table** depends on whether your ASes are registered with **PeeringDB: Exchanges** group if yes, **Custom - Exchanges** if no.

### Remove a Site Mapping

There are three different ways to remove a site mapping to a PeeringDB facility.

#### Remove from Facilities Group

To remove a site mapping from a facility listed in the **Facilities** group of the **PeeringDB Mapping Table**:

1. Choose **Integrations** from the main menu (under **Discover**).
2. Click **PeeringDB** (under **Network Platforms**) to open the Integrate PeeringDB dialog.
3. In the **Facilities** group on the dialog's mapping table, each site will be listed in a lozenge with an **X** on the right. Click the **X**s for the site(s) you’d like to remove.
4. Click **Save** to exit the dialog and return to the **Integrations** page.

#### Remove from Custom - Facilities

To remove a site mapping from a facility listed in the **Custom - Facilities** group:

1. Choose Settings from the main menu, then **Manage Sites** (under Network Metadata).
2. In the **Sites List**, click the edit button in the row for the site you’d like to modify. This will open the **Edit Site** dialog.
3. In the **PeeringDB Facility Mapping** field, each site will be listed in a lozenge with an **X** on the right. Click the **X**s for the site(s) you’d like to remove.
4. Click **Save** to exit the dialog and return to the **Sites** page.

#### Remove Sites List Facilities

To remove site mappings using the **Sites List**:

1. On the main portal menu, choose **Settings** to open the Settings page.
2. Click **Manage Sites** (under **Network Metadata**) to open the Sites page.
3. In the **Filters** pane to the left of the Sites list, use the **PeeringDB Facility** checkboxes to filter the **Sites** list to only sites that map to the checked facilities.
4. In the **Sites** list, click the checkbox for each site whose mapping you'd like to remove.
5. From the buttons that appear above the **Sites** list (see **Selected Site Actions**), click **PeeringDB Facility**.
6. Click **None**. The **PeeringDB Facility** column for the selected sites will now be empty.

### Remove an Interface Mapping

There are three different ways to remove an interface mapping to a PeeringDB exchange.

#### Remove from Exchanges Group

To remove a site mapping from a facility listed in the **Exchanges** group of the **PeeringDB Mapping Table**:

1. Choose **Integrations** from the main menu (under **Discover**).
2. Click **PeeringDB** (under **Network Platforms**) to open theIntegrate PeeringDB dialog.
3. In the **Exchanges** group on the dialog's mapping table, each interface will be listed in a lozenge with an **X** on the right. Click the **X**s for the interface(s) you’d like to remove.
4. Click **Save** to exit the dialog and return to the **Integrations** page.

#### Remove from Custom - Exchanges

To remove an interface mapping from a facility listed in the**Custom - Exchanges** group:

1. On the main portal menu, choose **Settings** to open the Settings page.
2. Click **Manage Interfaces** (under **Network Metadata**).
3. In the **Interfaces List**, click the Edit icon in the row for the interface you’d like to modify. This will open the Edit Interface dialog.
4. Click the **PeeringDB Internet Exchange** button.
5. In the resulting drop-down, click **None**.
6. Click **Save** to exit the dialog and return to the **Interfaces** page.

#### Remove Interfaces List Exchanges

To remove interface mappings using the **Interfaces List**:

1. On the main portal menu, choose **Settings** to open the Settings page.
2. Click **Manage Interfaces** (under **Network Metadata**) to open the Interfaces page.
3. In the **Filters** pane to the left of the Interfaces list, use the **PeeringDB IX** checkboxes to filter the **Interfaces** list to only interfaces that map to the checked exchanges.
4. In the **Interfaces** list, click the checkbox for each interface whose mapping you'd like to remove.
5. In the buttons that appear above the **Interfaces** list, click **PeeringDB IX**.
6. Click **None**. The **PeeringDB IX** column for the selected interfaces will now be empty.

### Remove All Mappings

The **Remove All Mappings** button in the bottom left of theIntegrate PeeringDB dialog opens a popup with Remove and Cancel buttons. Click **Remove** to confirm that you wish to delete all of the mappings in the **Assigned** column of the **Custom-Facilities** and **Custom-Exchanges** groups of the **PeeringDB Mapping Table**. After clicking the button, the Integrate PeeringDB dialog will close.

> **Note:** This button has no effect on the **Facilities** and **Exchanges** groups of the mapping table.

---

© 2014-25 Kentik
