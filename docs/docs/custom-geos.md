# Custom Geos

Source: https://kb.kentik.com/docs/custom-geos

---

This article covers the **Custom Geos** feature of the Kentik portal.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(662).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A20Z&se=2026-05-12T09%3A36%3A20Z&sr=c&sp=r&sig=XR1Z35cnWFe39jX7a3vYwRGWv%2Bafbw7pdRFKGSsi0Cs%3D)

*Custom Geos enables organizations to group countries into custom-defined regions.*

## About Custom Geos

Custom Geos (geographies) enable Kentik portal administrators to assign countries to custom groups that they can then use in queries via group-by or filter dimensions from the Custom Geo dimension family.

A custom geo is simply a property with a text value. Apply this value to one or more countries to create a custom-tailored group of countries that can be referred to together rather than individually.

> **Notes:**
>
> * Custom geos are mutually exclusive, so a given country can only be included in one custom geo at a time.
> * Members can view the information displayed on the custom geos page but cannot make any changes.

#### Using Custom Geos

Once defined on the Custom Geos page (Settings » **Custom Geos**), custom geos can be used a couple different ways:

* **Custom geo group-by**: If Source or Destination Custom Geo is included in a dashboard, an alerting policy, or a query as a group-by dimension (see **Dimensions in the Portal**), the traffic from all countries assigned to a given custom geo will be evaluated together. For example, if you assign Norway, Sweden, Denmark, Finland, and Iceland to a custom geo for Scandinavia then when a Custom Geo dimension is included in a query the traffic from all five countries will be treated as a single group (Scandinavia) for the top-X evaluation (see **Choosing Query Dimensions**).
* **Custom geo filter**: If Source or Destination Custom Geo is chosen as a dimension for a filter, then all currently defined custom geos (such as Scandinavia from our example above) will be available to choose from the drop-down list of filter values (see **Ad Hoc Filter Controls**).

#### Custom Geos in Flow Records

Each flow record whose source or destination country is assigned to a custom geo is permanently tagged with that custom geo at time of ingest, which has the following important consequences:

* A flow's custom geo value won't change even if the countries making up that custom geo are later changed. As a result, if a query's time range includes a point at which there was a change in the custom geo assignment of a given country that change will be reflected in the time series data returned by the query.

  + **Example***:* Let's say that you have all of your custom geos defined by Jan 1, but on Jan 5 you realize that you had mistakenly assigned Namibia to Scandinavia instead of Africa, so you move it from one custom geo to the other. Then on Jan 10 you run a query looking back 10 days. The time series for Africa would show increased traffic at the time of change, and Scandinavian traffic would show a drop at the same time.
* Editing the name of a custom geo effectively creates a new custom geo even if the assigned countries are not changed. Traffic from the countries of the original custom geo will be evaluated separately from the traffic of the same countries that is ingested after the custom geo is renamed.

> **Notes:**
>
> * Unless changed by an administrator on the **Custom Geos Page**, your organization’s custom geos will default to continents.
> * Custom geo (source or destination) is among the dimensions supported for the **Geo HeatMap** visualization type.

## Custom Geos Page

Custom geos are managed on the Custom Geos page. To access the page, choose **Settings** from the Kentik navbar, then select **Custom Geos** under Data Enrichment.

### Custom Geos Page UI

The Custom Geos page includes the following UI elements:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(663).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A20Z&se=2026-05-12T09%3A36%3A20Z&sr=c&sp=r&sig=XR1Z35cnWFe39jX7a3vYwRGWv%2Bafbw7pdRFKGSsi0Cs%3D)**Custom Geo Map**: A world map that is color-coded to indicate which countries belong to your organization's various custom geos.
* **Custom Geo list** (at right): A list of your organization’s currently defined custom geos. Each row in the list includes the name of a custom geo against a background whose color corresponds to the color used for the custom geo on the map. A row also includes the following UI elements:

  + Edit button (pencil icon): Opens a form to edit the custom geo (see **Custom Geo Controls**).
  + Remove button (trash icon): Removes the custom geo.

    > **Note:** There is no confirmation dialog after you click the Remove button; the custom geo is deleted immediately.
* **Add Custom Geo button**: Opens the **Custom Geo Controls**.

### Custom Geo Controls

The custom geo controls enable you to manage your organization’s collection of custom geos. To access the controls:

* Click the **Add Custom Geo** button to add a new custom geo.
* Click the edit icon to modify an existing custom geo.

The controls include the following fields and buttons:

* **Custom Geo Name**: Enter a name for the new custom geo or edit the name of an existing custom geo.

  > **Note:** Because flow records are associated with custom geos at ingest, changing the name of an existing custom geo will effectively create a new custom geo even if the assigned countries are not changed.
* **Countries**: A field listing the countries included in an existing custom geo (if editing) or to be included in a new custom geo (if adding). See **Changing the Countries List**.![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(664).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A20Z&se=2026-05-12T09%3A36%3A20Z&sr=c&sp=r&sig=XR1Z35cnWFe39jX7a3vYwRGWv%2Bafbw7pdRFKGSsi0Cs%3D)
* **Cancel**: Cancels the operation without adding a new custom geo. If editing a custom geo, all elements will be restored to their values at the time the controls were opened.
* **Save (**button, active only when text is present in the **Custom Geo Name** and **Countries** fields): Adds a new custom geo or updates an existing one with the specified name. A new custom geo will appear in the custom geo list.

#### Changing the Countries List

The Countries field lists the countries currently assigned to a custom geo. With the custom geo controls displayed, countries may be added or removed from the list in the following ways:

* **Adding a country**:

  + **In the field**: Click to drop down a menu of countries, then scroll to and click on one country in the menu. The menu will close when the country is added. Repeat to add all desired countries.
  + **In the map**: Click a country that isn't already assigned to a custom geo (countries whose color is gray). The country will be added to the list. Repeat to add all desired countries.

    > **Note:** If you click a country that is already a part of another custom geo, it will be removed from the other custom geo and added to the current one.
* **Removing a country**:

  + In the field: Click on the **X** at the right of the country you'd like to remove. The country will be removed from the list.
  + In the map: Click the country that you want to remove from this custom geo. The country will be removed from the list. Repeat to remove additional countries.

After making any changes to a custom geo, click **Save** to confirm the changes.

## Add or Edit Custom Geos

The management of custom geos is covered in the following sections.

### Adding a Custom Geo

To add a new custom geo to your organization’s collection of custom geos:

1. Open the Custom Geos page (Settings » **Custom Geos**).
2. Click the **Add Custom Geo** button to open the **Custom Geo Controls**.
3. In the **Custom Geo Name** field, enter a name for the new custom geo.
4. Add countries to the **Countries** field by clicking either in the field or directly on the country on the map (see **Changing the Countries List**).
5. Click the **Save** button. The custom geo will be added to the bottom of the **Custom Geo** **List**.

### Editing a Custom Geo

To edit an existing custom geo:

1. Open the Custom Geo page (Settings » **Custom Geos**).
2. In the **Custom Geo List**, click the **Edit** button (pencil icon) beside the name of the custom geo that you want to edit. The **Custom Geo Controls** will open and the custom geo’s name will appear italic in the Custom Geo list.
3. Edit the custom geo:

   1. To change the name, enter a new name In the **Custom Geo Name** field.

      > **Note:** Because custom geo labeling is applied to flow records at ingest, changing the name of an existing custom geo will effectively create a new custom geo even if the assigned countries are not changed.
   2. To change the countries assigned to the custom geo, use the **Countries** field or the map (see **Changing the Countries List**).
4. Click the **Save** button to save changes to the custom geo.
