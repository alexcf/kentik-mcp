# Resolution Overview

Source: https://kb.kentik.com/docs/resolution-overview

---

Kentik generates two independent dataseries at ingest into the **Kentik Data Engine** (KDE), enabling efficient querying over extended periods without losing detail for shorter timespans.

## About Dataseries Resolution

Kentik builds both the Full and Fast dataseries at data ingest, maintaining them as independent entities. The default dataseries for a query varies by timespan:

* **Full Dataseries**: Used for timespans < 24 hours.
* **Fast Dataseries**: Used for timespans ≥ 24 hours.

> **Note*:*** The 24-hour criterion is based solely on the duration of the query's timespan, not how far back in time the query is looking.

The following table compares the storage, performance, and data retention characteristics of the Full and Fast dataseries:

| Feature | Full Dataseries | Fast Dataseries |
| --- | --- | --- |
| Retention | Standard (Detailed) | Extended (Historical) |
| Query Speed | Standard | High-Speed |
| Granularity | Every record | Downsampled (30 fps) |
| Port Detail | All ports preserved | Ports > 32767 grouped to **65535** |
| Aggregation | None | 7-tuple aggregation |

#### Resolution Intervals

The dataseries resolution determines the interval (granularity) used for reporting packets and bytes:

* **Fast Dataseries**: Query intervals "snap" to the nearest hour. For example, for a query from 1:05 PM on 5/12 to 10:31 PM on 5/15, the results will cover 1:00 PM on 5/12 to 11:00 PM on 5/15.

  > **Note**: 'Snapping' ensures query performance but may result in slight data rounding (flat-lining) at the start and end of your selected time range.
* **Full Dataseries**: Query intervals vary depending on query width.

> **TIP:** You can manually select timespans as short as **3 hours** for **Fast Dataseries**, and up to **72 hours** for **Full Dataseries** (see **Query Resolution Selection**).

## Query Resolution Selection

The default dataseries selection may be overridden in both the portal and the API.

### Portal Resolution Selection

In the Kentik portal, the dataseries selection is based on the timespan set in Dashboards and Data Explorer:

* **Dashboards**:

  + **< 24 hours**: Automatically selects **Full** (see **About Dataseries Resolution**).
  + **≥ 24 hours**: Automatically selects **Fast**.
* **Data Explorer**:

  + **≤ 3 hours**: Always uses **Full**.
  + **3 - 24 hours**: Defaults to **Full** (user can manually toggle to Fast).
  + **24 - 72 hours**: Defaults to **Fast** (user can manually toggle to Full).
  + **> 72 hours***:* Always uses **Fast**.

### API Resolution Selection

The Kentik API uses the same dataseries defaults as the portal, with an option to override using the `i_fast_dataset` query parameter:

* **Without** `i_fast_dataset`:

  + Automatically selects Full for timespans less than 24 hours (see **About Dataseries Resolution**).
  + Selects Fast for timespans of 24 hours or more.
* **With** `i_fast_dataset`:

  + `false:` Uses Full dataseries for timespans up to 72 hours.
  + `true:` Uses Fast dataseries regardless of timespan.

> **Note*:*** Avoid using the Fast dataseries for timespans under 3 hours and the Full dataseries for timespans over 72 hours.
