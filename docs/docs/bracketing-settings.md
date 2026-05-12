# Bracketing Settings

Source: https://kb.kentik.com/docs/bracketing-settings

---

This article discusses bracketing in queries for **Data Explorer** and dashboard panels in the Kentik portal.

## About Bracketing

The bracketing feature allows you to define from two to five ranges (brackets) of values and assign various colors to the brackets so that you can see at a glance the range into which the current value falls. Bracketing is configured with the **Bracketing Options Dialog**.

> **Note:** Bracketing is required when view type is Gauge (see **Chart Visualization Types**).

## Bracketing Visualization Types

As shown in the following table, the effect of bracketing in charts and results tables (e.g. in Data Explorer or a dashboard panel) depends on the current visualization type.

| Visualization Type | Bracketing Required? | Chart Bracketed? | Table Bracketed? |
| --- | --- | --- | --- |
| Time Series Stacked Graph | No | No | Yes |
| Time Series 100% Stacked Graph | No | No | Yes |
| Time Series Bar Graph | No | No | Yes |
| Time Series Line Graph | No | Yes | Yes |
| Comparison Bar Chart | No | Yes | Yes |
| Pie Chart | No | No | Yes |
| Sankey Flow Diagram | No | No | Yes |
| Table | No | N.A. | Yes |
| Matrix Diagram | No | Yes | N.A. |
| Gauge | Yes | Yes | Yes |
| Geo HeatMap | No | Yes | Yes |

In supported view types, bracketing has the following effect:

* **Bracketing in charts**: Results displayed in the chart will be colored by bracket. In the Gauge chart at right, for example, the background color (red) corresponds to the bracket that includes the current bracketing value (see **Bracketing Value Control**).
* **Bracketing in tables**: The bracket of each row will be indicated with a colored bar at the left of the row (as shown below).

  > **Note**: If a series is excluded from bracketing because its value falls outside of the range defined with the **Trim Overall Range** controls (see **Bracketing Range Controls**) then the colored bar for the row corresponding to that series will be gray.

## Bracketing Pane

The **Bracketing** pane appears in the sidebar in Data Explorer and in the settings dialog for the individual panels on dashboards. The pane includes the following UI elements:

* **Range** indicator: Appears only when bracketing has been applied. The indicator shows the upper bound and the color of each range.
* **Edit Bracketing**: A button that opens the **Bracketing Options Dialog**.

## Bracketing Options Dialog

The Bracketing Options dialog defines the type of bracketing and the ranges of the brackets. The dialog consists of the following general UI elements as well as several panes described below:

* **Close button**: Click the **X** in the upper right corner to close the dialog without saving changes to the settings.
* **Cancel button**: Close the dialog without saving changes. All elements will be restored to their values at the time the dialog was opened.
* **Save button**: Save changes to settings and exit the dialog.
* **Remove button** (shown only when bracketing already exists on the query): Opens a confirming dialog that allows you to remove the bracketing. When bracketing is removed, the dialog is reset to defaults and closed. You must click the **Run Query** button in the sidebar to execute removal of the bracketing from the query.

### Bracketing Type Control

The **Bracketing Type** pane includes a single drop-down menu used to choose the basis on which the brackets will be defined in the Bracketing Ranges pane, which may be one of the following:

* **Static ranges**: Series are bracketed by comparing Bracketing Value to user-specified numeric ranges.
* **Percentages**: Series are evaluated by Bracketing Value and bracketed by percent of the highest value.
* **Percentiles**: Series are ranked by Bracketing Value and bracketed by percentile relative to all series.

As shown in the following set of screen shots, the **Bracketing Type** setting determines the key displayed in the **Bracketing** pane in the sidebar.

### Bracketing Value Control

The **Bracketing Value** pane includes a single switch, **Use Last Datapoint Value**, which is used to determine the value by which each series is evaluated for bracketing:

* If the switch is On: Series are evaluated by the value of their most recent data point.
* If the switch is Off: Series are evaluated by the value of the **Primary Display & Sort Metric** specified in the Metrics dialog (accessed via the **Customize Metrics** button in the sidebar’s **Query** pane).

### Bracketing Range Controls

The **Bracketing Ranges** pane includes the following controls:

* **Group Results**: A switch that, when on, sets the results table to include a collapsible header row for each bracket.
* **Trim Overall Range**: A switch that, when on, enables the following fields, which are used to narrow the range of values considered for bracketing:

  + **Exclude values under**: Set the minimum value for the lowest range.
  + **Exclude values over**: Set the maximum value for the highest range.
  + **Value type**: A drop-down menu used to choose the value type (fixed, percentage, or percentile) for the minimum and maximum values.
* **Range controls**: Set the color and value boundary associated with each range, as well as the number of ranges. By default there are two ranges, but you can have up to five ranges with different colors (e.g. orange for a range that’s slightly above normal and red for a critical peak).

The following controls are available for each range:

* **Select A Color:** Click the button to choose a color for the range.
* **Value field**: Enter the highest value that will be included in the range:

  + The lowest range will include values from the entered number down to the **Exclude Values Under** value from **Trim Overall Range**.
  + Intermediate ranges will include values above the previous range up to and including the entered number.
  + The field is not shown for the highest range.
* **Delete button**: Click the red **X** to delete the range.
