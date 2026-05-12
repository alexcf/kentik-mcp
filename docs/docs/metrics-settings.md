# Metrics Settings

Source: https://kb.kentik.com/docs/metrics-settings

---

This article covers metrics settings for queries in **Data Explorer** and in dashboard panels in the Kentik portal.

> **Note:** For more information about metrics, see **About Metrics**.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/Metrics Dialog.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A57Z&se=2026-05-12T09%3A33%3A57Z&sr=c&sp=r&sig=bZ6ch5CNLfvlsK6VgMLkIGKufmdUb43KV8vwiDaCYxs%3D)

*Metrics settings pane in Data Explorer*

## Metrics Pane

The Metrics pane of the **Query** sidebar in Data Explorer includes the following UI elements:

* **Custom** (drop-down): Choose the primary metric for query results (see **About Metrics**).
* **Edit Metrics**: A button that opens the Metrics dialog (see **About the Metrics Dialog**).

> **Note:** For each metric in the drop-down menu Kentik has preselected a set of additional metrics that will be shown in the results table when the query is run. To customize the additional metrics, click the **Customize****Metrics** button.

## About the Metrics Dialog

A metric is a combination of a unit (e.g., a bit) with a method of calculation (e.g., average per second) to create a quantifiable measurements (average bits/second). In **Data Explorer**, the Metrics dialog determines how metrics are displayed in the chart and table:

* **Primary metric**: The metric that is plotted on the positive Y axis in the chart, and by which the table will initially be sorted.
* **Secondary metric**: If configured, this metric is plotted on the negative Y axis in a compound query (see **Compound Queries**).
* **Additional metrics**: Multiple additional metrics may be chosen to appear as columns in the results table (see **Data Explorer Table**).

## Metrics Dialog UI

The following UI elements are used to set primary, secondary, and additional metrics:

* **Close button**: Click the **X** in the upper right corner to close the dialog. All elements will be restored to their values at the time the dialog was opened.
* **Use Preset Selections For**: A drop-down menu that lists primary metrics. When you choose a primary metric any existing metrics in the **Selected Metrics** list are removed and the chosen metric is added to the list along with a set of additional metrics chosen by Kentik to supplement the primary metric.
* **Clear Selections button**: Clears all metrics from the **Selected Metrics** list.
* **Filter field**: Filters the list of available metrics to those whose name contains the entered text.
* **Selected Metrics list**: A list of the metrics currently selected from the **Available Metrics** list. To change the order in which the metrics are displayed in the results table, drag the handle at the left of each selected metric.
* **Available metrics list**: A list of all of the metrics that are currently supported by Kentik, grouped into unit-based categories (see **Metrics Reference**).

  + To select all metrics in a given category (e.g., Bits/s or Unique IPs), check the box at the right of the category name.
  + To select all calculations (e.g., Average, 95th percentile, and Max) of a given metric, check the box at the right of the metric name.
  + To select an individual metric (e.g., Average Flows/s) check the box to the left of the individual metric in the list.
  > **Notes:**
  >
  > + TCP metrics are available only when the currently selected data sources (see **Data Sources Pane**) include one or more hosts.
  > + The categories of Bits/s and Packets/s are broken into subcategories based on where flow is sampled: Ingress, Egress, or both.
* **Primary Display & Sort Metric**: A drop-down menu used to choose the metric that is plotted on the positive Y axis in the chart, and by which the table will initially be sorted.

  > **Note:** For TCP (host) metrics related to percents and latencies (e.g., % Retransmits, Server latency, etc.) the menu items include the option to sort using Kentik Intellisort, in which the rows of the result table are sorted based on assumptions about which results would logically be the highest priority for users to know.
* **Secondary Display Metric** (shown only when the **Selected Metrics** list includes metrics from more than one category): A metric that will be plotted on the negative Y axis in a compound query (see **Compound Queries**).
* **Sort Secondary Metric Separately** (shown only if **Secondary Display Metric** is set to something other than None): If the switch is On, the results table will include a second tab in which the rows are sorted by the secondary metric.
* **Cancel button**: Cancel changes to the selected metrics and exit the dialog. The current query’s metrics will be restored to what they were when the dialog was opened.
* **Save button**: Save changes to the selected metrics and exit the dialog.

## Default Metrics

Decide which metrics are used by default for your Data Explorer queries. By default, Data Explorer always uses the same three default metrics when users build queries from scratch:

* Average
* 95th Percentile
* Max

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(750).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A57Z&se=2026-05-12T09%3A33%3A57Z&sr=c&sp=r&sig=bZ6ch5CNLfvlsK6VgMLkIGKufmdUb43KV8vwiDaCYxs%3D)

*Default Metrics options to set and load defaults*

The Default Metrics panel allows you to:

* Click the **Set as Default** button, which will set the currently selected Metrics as default.

  + The next new Data Explorer query you create will have these metrics selected by default.
  + Saving currently selected Metrics to default also includes the settings in the Primary and Secondary Display and Sort Metrics.
* Hover over **Load Defaults** to show the defaults that have been set.

  + Before loading these to replace the current set of selected metrics, a tooltip will show the metrics in your default set, and it will differentiate the primary from the secondary metrics.
  + Both metrics will be highlighted with a colored dot and the number to show primary versus secondary.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(751).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A57Z&se=2026-05-12T09%3A33%3A57Z&sr=c&sp=r&sig=bZ6ch5CNLfvlsK6VgMLkIGKufmdUb43KV8vwiDaCYxs%3D)

*Tooltip that displays the primary and secondary metrics when hovering over Load Defaults*

* Click on **Load Defaults** to load queries to replace the current selection.

> **Note:** Defaults are set per user, and they will differ between different users within the same company.
