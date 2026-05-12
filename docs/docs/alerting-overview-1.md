# Alerting Overview

Source: https://kb.kentik.com/docs/alerting-overview-1

---

This article covers the Alerting Overview page of the Kentik portal.

> **Notes:**
>
> * For an overview of the Alerting page, see **Alerting**.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(772).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A05Z&se=2026-05-12T09%3A32%3A05Z&sr=c&sp=r&sig=kAnbchBRJqsXZcYY3DaftCMCg8fankiFtZSoilvQAF8%3D)

*Alerting Overview page*

The Alerting Overview page provides an interactive view of how the shape of alert volume changes over time. This helps identify and prioritize problems from a macro view.

* View highlights of:

  + Alerts by Type
  + Most Triggered Policies
  + Monthly Alert Trends by Severity
  + Alerts by Site
* Subscribe to Alerting Overview reports.
* Export reports to PDF.
* Filter the report by past months (3, 6, 9, and 12) and source alert type.

## Alerting Overview UI

The Alerting Overview page includes the following UI elements:

* **Actions** (button): A drop-down to choose **Export** or **Subscribe**.

  + **Export**: Export visual reports as PDFs (see **Portal Export Options**).
  + **Subscribe**:Opens the Subscribe dialog to create an Alerting Overview report subscription (see **Subscription Dialog UI**and**Sharing by Subscription**).
* **Time period** (button): A drop-down to choose a time period (Past 3 months, Past 6 months, Past 9 months, or Past 12 months) for the visualizations.
* **Sources**:Select one or more sources to update the visualizations on the page.

  + To customize visualizations, deselect each source accordingly.

    ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(773).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A05Z&se=2026-05-12T09%3A32%3A05Z&sr=c&sp=r&sig=kAnbchBRJqsXZcYY3DaftCMCg8fankiFtZSoilvQAF8%3D)

    *Traffic and NMS are enabled sources for visualization. Cloud and Protect sources are not selected and do not display in the Alerting Overview visualizations.*
* **Zoom In/Out** (+/- buttons):Buttons on the Alerts by Site map that narrows/widens the view of the map.
* **Sites**: Click a link under the Site column of the Alerts by Site pane to navigate to the Alerting page of that particular site.

## Alerts by Type Pane

This pane displays a pie chart of the alerts selected from Sources, for the time period selected from the time period drop-down.

Hover over one of the colors to isolate the alerts by a single source.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(776).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A05Z&se=2026-05-12T09%3A32%3A05Z&sr=c&sp=r&sig=kAnbchBRJqsXZcYY3DaftCMCg8fankiFtZSoilvQAF8%3D)

## Most Triggered Policies Pane

This pane displays a list of the 10 most triggered policies based on the Sources selected.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(777).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A05Z&se=2026-05-12T09%3A32%3A05Z&sr=c&sp=r&sig=kAnbchBRJqsXZcYY3DaftCMCg8fankiFtZSoilvQAF8%3D)

## Monthly Alert Trends by Severity Pane

This pane displays a bar chart visualization for the number of alert trends (y-axis) and corresponding severity percentage (x-axis). The severity levels are:

* Critical
* Severe
* Major
* Warning
* Minor

Hover over one of the colored bars to isolate the monthly alert trends by a single severity level.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(775).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A05Z&se=2026-05-12T09%3A32%3A05Z&sr=c&sp=r&sig=kAnbchBRJqsXZcYY3DaftCMCg8fankiFtZSoilvQAF8%3D)

## Alerts by Site Pane

This pane includes two visualization components:

1. A map that shows severity levels by location.
2. A table of listing the number of alerts per level of severity by site.

### Map

The Alerts by Site pane map displays locations and their corresponding severity levels. If the location dot is green it means there are no alerts for that location. All other location dot colors indicate that there is at least one alert for the location.

* Hover over the location dot to display the list of severity levels followed by the number of alerts for each severity level.
* Use the +/- buttons to zoom in/out on a specific area.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(782).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A05Z&se=2026-05-12T09%3A32%3A05Z&sr=c&sp=r&sig=kAnbchBRJqsXZcYY3DaftCMCg8fankiFtZSoilvQAF8%3D)

### Table

The table includes site links in the far-left column with the corresponding number of alerts per level of severity in the following columns.

* Click a site link to navigate to the Alerting page for that particular site.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(779).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A05Z&se=2026-05-12T09%3A32%3A05Z&sr=c&sp=r&sig=kAnbchBRJqsXZcYY3DaftCMCg8fankiFtZSoilvQAF8%3D)
