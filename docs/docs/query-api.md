# Query API

Source: https://kb.kentik.com/docs/query-api

---

The Kentik V5 Query API allows programmatic querying of Kentik-stored network data (flow records, BGP, selected SNMP OIDs) for your organization's devices.

> Notes:
>
> * For an overview of all Kentik APIs, see **Kentik APIs**.
> * For information on using the Kentik V5 Query API with cURL, see **API Access Via cURL**.
> * For help using any API version, contact **support@kentik.com**.
> * As of May 1st, 2025, the **Query SQL Method** has been deprecated and is no longer supported.

## About V5 Query API

The Kentik V5 Query API provides methods to retrieve information from the main tables of the Kentik Data Engine (see **KDE Tables**) as JSON data or image files for charts similar to those in Data Explorer.  
  
 Kentik V5 Query APIs can be accessed via:

* **Command line**: Use **cURL** in environments like Terminal. See **API Access Via cURL**.
* **Programmatically**: Send a request body to a V5 API endpoint using any application language that supports HTTP.

> **Note:**The V5 API tester was discontinued in January 2025.

## V5 Query Methods

The following table lists available methods in the V5 Query API. Click the topic link for more details on each method:

| Method | Endpoint | Description | Topic |
| --- | --- | --- | --- |
| POST | /topXchart | Returns a graph as image file data. | **Query Chart Method** |
| POST | /url | Returns a URL to open Data Explorer in the Kentik portal, with the sidebar configured according to the request JSON query parameters. | **Query URL Method** |
| POST | /topXdata | Returns JSON data for generating tables and time-series graphs similar to those in Data Explorer. | **Query Data Method** |

> **Note:** Endpoint URLs are region-specific, `https://api.kentik.com/api/v5` for US and `https://api.kentik.eu/api/v5` for EU.

## Query API JSON

The JSON for requests and responses in V5 Query API methods is detailed here.

### Query API Request JSON

The Query Chart, Query Data, and Query URL methods of the V5 Query API use JSON in the request body. The JSON array specifies what data Kentik should return and how. This JSON structure is largely consistent across all three methods.  
  
You can create the JSON array for a Query Chart, Query Data, or Query URL call by:

* Building the JSON from scratch using the information in the tables below.
* Retrieving the JSON from the Kentik portal:  
   1. In **Data Explorer**, define your query with the fields in the sidebar panes and click **Apply Changes**.  
   2. In the results, open the options menu at the upper right of the graph (see **Explorer Title Pane**), select **Show API Call**, then choose Data or Chart to open the Query API example dialog.  
   3. Click **JSON** to view the query JSON.  
   4. Click **Copy** to copy the JSON to your clipboard.  
   5. Paste the JSON in a Query Data call in cURL or another compatible setting for generating HTTP calls.

The example below demonstrates the objects utilized in the Query Chart, Query Data, and Query URL methods:

```
{
  "queries": [
    {
      "query": {
        "viz_type": "matrix",
        "show_overlay": false,
        "overlay_day": -7,
        "sync_axes": false,
        "metric": "bytes",
        "matrixBy": [
          "dst_geo_city"
        ],
        "cidr": 32,
        "cidr6": 128,
        "pps_threshold": 500,
        "topx": 8,
        "depth": 100,
        "fastData": "Auto",
        "outsort": "max_bits_per_sec",
        "lookback_seconds": 3600,
        "time_format": "UTC",
        "hostname_lookup": "false",
        "starting_time": null,
        "ending_time": null,
        "device_name": "device_name_1, device_name_2",
        "all_selected": true,
        "filters_obj": {
          "connector": "All",
          "filterGroups": [
            {
              "connector": "All",
              "filters": [
                {
                  "filterField": "protocol",
                  "operator": "=",
                  "filterValue": "17",
                }
              ],
              "not": false,
            }
          ],
        },
        "saved_filters": [
          {
            "filter_id": 363
            "is_not": true
          }
        ],
        "descriptor": "Total",
        "aggregates": [
          {
            "name": "avg_bits_per_sec",
            "column": "f_sum_both_bytes",
            "fn": "average",
            "sample_rate": 1
          },
          {
            "name": "p95th_bits_per_sec",
            "column": "f_sum_both_bytes",
            "fn": "percentile",
            "rank": 95,
            "sample_rate": 1
          },
          {
            "name": "max_bits_per_sec",
            "column": "f_sum_both_bytes",
            "fn": "max",
            "raw": true,
            "sample_rate": 1
          }
        ],
        "query_title": "Top Source City vs Top Dest City by Max Bits/s",
        "dimension": [
          "src_geo_city"
        ]
      },
      "bucket": "Matrix",
      "bucketIndex": 0,
      "isOverlay": false
    }
  ],
  "imageType": "image_file_format"
}
```

The following sections describe the objects in detail.

#### Query JSON Object

The highest-level object of the V5 Query JSON includes these two members:

| Field | Type | Description |
| --- | --- | --- |
| queries | array | One or more query objects, each element of which includes a query object and additional associated members. See **Queries Array**. |
| imageType | string | Required: The format of the image file for returning a chart: pdf, png, jpg, or svg.   * Used only in Query Chart method. |

#### Queries Array

The `queries` array usually includes a single `query` object and the associated objects shown in the following table. For multi-series data or charts (see **Compound Queries**), add multiple elements, each representing a desired chart axis.

| Field | Type | Description |
| --- | --- | --- |
| query | object | An object with members that specify a query for Kentik. See **Query Object**. |
| bucket | string | A name for a set of queries to run simultaneously for multi-series graphing. Queries can inherit values from those with a lower bucketIndex. |
| bucketIndex | number | The index of a query within its bucket, with valid values being integers greater than -1. |
| isOverlay | boolean | Determines if the query is treated as an overlay, applicable only in the Query Chart, Query URL methods. |

#### Query Object

The `query` object includes the objects and arrays that specify a query for execution on Kentik:

| Field | Type | Description |
| --- | --- | --- |
| viz\_type | string | Type of chart (see **Chart Visualization Types**).   * Used only in these methods: Query Chart, Query URL. * Valid values: stackedArea (default), line, stackedBar, bar, pie, sankey, table, matrix |
| show\_overlay | boolean | Determines if historical overlay is shown (see **Advanced Query Settings**).   * Ignored unless overlay is supported by specified viz\_type. * Used only in these methods: Query Chart, Query URL. |
| overlay\_day | number | Number of days to look back for historical overlay.   * Ignored unless show\_overlay is true and overlay is supported by specified viz\_type. * Used only in these methods: Query Chart, Query URL. |
| sync\_axes | boolean | Determines whether left and right axis scales are synced.   * Used only for **Compound Queries**. * Used only in these methods: Query Chart, Query URL. |
| metric | string | Required: Unit of measure, e.g., bytes, packets, etc. See **About Metrics**   * Valid value: bytes, in\_bytes, out\_bytes, packets, in\_packets, out\_packets, tcp\_retransmit, perc\_retransmit, retransmits\_in, perc\_retransmits\_in, out\_of\_order\_in, perc\_out\_of\_order\_in, fragments, perc\_fragments, client\_latency, server\_latency, appl\_latency, fps, unique\_src\_ip, or unique\_dst\_ip |
| matrixBy | array | One or more matrix-with dimensions (see **Matrix Visualization**), specified in a JSON array.   * Valid values: any combination of standard dimensions (see **Dimensions Reference**) and/or **Custom Dimensions**. |
| cidr | number | A CIDR filter to use on IP queries, such as when grouping by dest\_ip/24.   * Valid values: digit between 0 and 32. |
| cidr6 | number | A CIDR6 filter to use on IP queries.   * Valid values: digit between 0 and 128. |
| pps\_threshold | number | The minimum packets per second required to match flow (used in conjunction with metrics like retransmits, out of order, etc.).   * Valid values: digit greater than 0. |
| topx | number | The number of topX rows from the pool of returned results (size of which is determined by the depth), that will be returned as time-series data and (for topXchart) included in the chart.   * Valid values: digit between 1 and 40. |
| depth | number | The number of results in the pool from which topX will be determined.   * Valid values: digit between 25 and 250. |
| fastData | string | Selects the dataset on which to run the query, overriding the default. See **API Resolution Selection**.   * Valid values: Auto, Fast, Full |
| outsort | string | The name of an aggregate object by which to sort results for depth and topx.   * Required if the query includes more than one user-specified aggregate (see aggregates below) * If no user-specified aggregates are needed, optionally use to specify the Kentik-derived aggregate (see **Kentik Sort-by Options**) by which results will be sorted. |
| lookback\_seconds | number | The look back time in seconds, e.g., for last 2 hours, set to 7200   * Required unless starting\_time and ending\_time are set. * Unless set to 0, this overrides the starting\_time and ending\_time settings. |
| time\_format | string | The time zone, either UTC or Local.   * Used only in the Query URL method. |
| hostname\_lookup | boolean | When true (default), allow DNS lookup.  **Note:** Set to false to speed queries against IP dimensions. |
| starting\_time | string | The fixed start time of a query time-span, e.g. 'YYYY-MM-DD HH:mm:00'.   * Ignored unless lookback\_seconds is zero. |
| ending\_time | string | The fixed end time of a query time-span, e.g. 'YYYY-MM-DD HH:mm:00'.   * Ignored unless lookback\_seconds is zero. |
| device\_name | array/string | List of names of the devices to query.   * To set multiple devices, use a comma delimited string. |
| device\_labels | array/string | List of label IDs (integers) to query. See Labels.  **Note**: Label names are not accepted, only label IDs. |
| all\_selected | boolean | Query against all devices.   * If true, device\_name is ignored. |
| filters\_obj | array | See **Filters\_obj Object**. |
| saved\_filters | array | See **Saved\_filters Array**. |
| descriptor | string | The name for labeling results when the queries array has multiple elements (Multi-series Graphing mode).   * Ignored unless dimension is “traffic.” |
| aggregates | array | See **Aggregates Array**. |
| query\_title | string | The title of a query (appears in the returned chart). |
| dimension | string | Required: One or more group-by dimensions, specified in a JSON array.   * Valid values: any combination of standard dimensions (see **Dimensions Reference**) and/or **Custom Dimensions**. * Default: AS\_dst * Valid values: AS\_src, Geography\_src, InterfaceID\_src, Port\_src, src\_eth\_mac, VLAN\_src, IP\_src, AS\_dst, Geography\_dst, InterfaceID\_dst, Port\_dst, dst\_eth\_mac, VLAN\_dst, IP\_dst, TopFlow, Proto, Traffic, ASTopTalkers, InterfaceTopTalkers, PortPortTalkers, TopFlowsIP, src\_geo\_region, src\_geo\_city, dst\_geo\_region, dst\_geo\_city, RegionTopTalkers, i\_device\_id, i\_device\_site\_name, src\_route\_prefix\_len, src\_route\_length, src\_bgp\_community, src\_bgp\_aspath, src\_nexthop\_ip, src\_nexthop\_asn, src\_second\_asn, src\_third\_asn, src\_proto\_port, dst\_route\_prefix\_len, dst\_route\_length, dst\_bgp\_community, dst\_bgp\_aspath, dst\_nexthop\_ip, dst\_nexthop\_asn, dst\_second\_asn, dst\_third\_asn, dst\_proto\_port, inet\_family, TOS, tcp\_flags |
| minsPolling | number | The flow aggregation window, or size of the data points (in minutes) that data is aggregated into across the timeframe of the query.   * Ignored if forceMinsPolling is set to False. * Valid values: 1, 5, 10, 20 |
| forceMinsPolling | boolean | Determines whether to use the automatic flow aggregation window. The minsPolling value for the flow aggregation window is used if the value is set to True.   * Valid values: True, False |
| reAggInterval | number | The time interval (in seconds) used for the aggregate function.   * Valid values: 300, 600, 1200, 1800, 3600, 21600, 43200, 86400, 604800 |
| reAggFn | string | The aggregate function that will be applied to the query results.   * Valid values: avg, max, min, none |

#### Filters\_obj Object

The `filters_obj` object contains one of more filter groups for a Query Chart or Query Data query. For more information on query filters, see **About Filters**.

| Field | Type | Description |
| --- | --- | --- |
| connector | string | Sets the conjunctive operator ("and" or "or") that will join all filter groups in this filter set. |
| filterGroups | array | See **FilterGroups Array**. |

#### FilterGroups Array

The `filterGroups` array defines a group of one or more filters for a Query Chart or Query Data query. For more information on filters in a query see **About Filters**. Each member of the array contains the following objects:

| Field | Type | Description |
| --- | --- | --- |
| connector | string | Sets the "and" or "or" operator to join all filter groups in this set. |
| filters | string | See **Filters Array**. |
| not | boolean | Determines if traffic matching the filter is included (false) or excluded (true). |

#### Filters Array

The `filters` array defines individual filters in a filter group for a Query Chart or Query Data query (see **About Filters**). Each array member represents a filter and includes the following objects:

| Field | Type | Description |
| --- | --- | --- |
| filterField | string | The KDE name of a dimension to filter on is listed in the **Dimensions Reference**. Dimension are shown by portal name, with the corresponding KDE name provided in the KDE Names column of the same row. |
| operator | string | The operator to apply in the filter.   * Valid value: "=", "<>", "ILIKE", "NOT ILIKE", "˜", "!˜", ">", "<", "&", “field\_not\_equal” |
| filterValue | string | The filterField value to match. |

#### Saved\_filters Array

The `saved_filters` array contains IDs of saved filters (custom or preset) to apply to the query (see **Saved Filters**).

| Field | Type | Description |
| --- | --- | --- |
| filter\_id | number | The ID of a saved filter. |
| is\_not | boolean | Determines if traffic matching the saved filter set is included (false) or excluded (true). Defaults to false. |

#### Aggregates Array

The `aggregates` array specifies an aggregate for a table column from a Query Chart query when the display type (`viz_type`; see **Query Object**) is Table. An aggregate is the result of basic statistical analyses on a dataset.

| Field | Type | Description |
| --- | --- | --- |
| name | string | The name of the aggregate. |
| column | string | The raw SQL column (see **KDE Tables**) to use as the source for this aggregation. |
| fn | string | The aggregation function to run against the column: sum, average, percentile, max, composite, exponent, modulus, greaterThan, greaterThanEquals, lessThan, lessThanEquals, equals, or notEquals |
| rank | number | The percentile to apply in an aggregate function.   * Ignored unless fn is "percentile" * Valid value: integer from 5 to 99. |
| sample\_rate | number | Set a rate at which the values will be sampled, e.g., multiply by -1 (-1), or divide by 1,000,000 (1/1,000,000). |

#### Kentik Sort-by Options

If no aggregates are specified in the **Query Object**, Kentik automatically generates aggregates based on the metric field. These Kentik-generated aggregates can be used in the outsort field to sort results for depth and topx. The table below lists valid Kentik-generated aggregates for each metric:

| Metric | Aggregate name |
| --- | --- |
| bytes | avg\_bits\_per\_sec p95th\_bits\_per\_sec max\_bits\_per\_sec |
| in\_bytes | avg\_in\_bits\_per\_sec p95th\_in\_bits\_per\_sec max\_in\_bits\_per\_sec |
| out\_bytes | avg\_out\_bits\_per\_sec p95th\_out\_bits\_per\_sec max\_out\_bits\_per\_sec |
| packets | avg\_pkts\_per\_sec p95th\_pkts\_per\_sec max\_pkts\_per\_sec |
| in\_packets | avg\_in\_pkts\_per\_sec p95th\_in\_pkts\_per\_sec max\_in\_pkts\_per\_sec |
| out\_packets | avg\_out\_pkts\_per\_sec p95th\_out\_pkts\_per\_sec max\_out\_pkts\_per\_sec |
| tcp\_retransmit | sum\_rxmits avg\_rxmits\_per\_sec p98th\_rxmits\_per\_sec max\_rxmits\_per\_sec p98th\_perc\_rxmits max\_perc\_rxmits avg\_bits\_per\_sec sum\_pkts avg\_pkts\_per\_sec p98th\_pkts\_per\_sec |
| perc\_retransmit | sum\_rxmits avg\_rxmits\_per\_sec p98th\_rxmits\_per\_sec max\_rxmits\_per\_sec p98th\_perc\_rxmits max\_perc\_rxmits avg\_bits\_per\_sec sum\_pkts avg\_pkts\_per\_sec p98th\_pkts\_per\_sec |
| retransmits\_in | sum\_rxmits avg\_rxmits\_per\_sec p98th\_rxmits\_per\_sec max\_rxmits\_per\_sec rxmits\_to\_pkts\_ratio perc\_rxmits p98th\_perc\_rxmits max\_perc\_rxmits avg\_bits\_per\_sec sum\_pkts avg\_pkts\_per\_sec p98th\_pkts\_per\_sec |
| perc\_retransmits\_in | sum\_rxmits avg\_rxmits\_per\_sec p98th\_rxmits\_per\_sec max\_rxmits\_per\_sec rxmits\_to\_pkts\_ratio perc\_rxmits p98th\_perc\_rxmits max\_perc\_rxmits avg\_bits\_per\_sec sum\_pkts avg\_pkts\_per\_sec p98th\_pkts\_per\_sec |
| out\_of\_order\_in | sum\_ooorder avg\_ooorder\_per\_sec p98th\_ooorder\_per\_sec max\_ooorder\_per\_sec ooorder\_to\_pkts\_ratio perc\_ooorder\*p98th\_perc\_ooorder max\_perc\_ooorder avg\_bits\_per\_sec  sum\_pkts avg\_pkts\_per\_sec p98th\_pkts\_per\_sec |
| perc\_out\_of\_order\_in | sum\_ooorder avg\_ooorder\_per\_sec p98th\_ooorder\_per\_sec max\_ooorder\_per\_sec ooorder\_to\_pkts\_ratio perc\_ooorder p98th\_perc\_ooorder max\_perc\_ooorder avg\_bits\_per\_sec sum\_pkts avg\_pkts\_per\_sec p98th\_pkts\_per\_sec |
| fragments | sum\_fragments avg\_fragments\_per\_sec p98th\_fragments\_per\_sec max\_fragments\_per\_sec fragments\_to\_pkts\_ratio perc\_fragments p98th\_perc\_fragments max\_perc\_fragments avg\_bits\_per\_sec sum\_pkts avg\_pkts\_per\_sec p98th\_pkts\_per\_sec |
| perc\_fragments | sum\_fragments avg\_fragments\_per\_sec p98th\_fragments\_per\_sec max\_fragments\_per\_sec fragments\_to\_pkts\_ratio perc\_fragments p98th\_perc\_fragments max\_perc\_fragments avg\_bits\_per\_sec sum\_pkts avg\_pkts\_per\_sec p98th\_pkts\_per\_sec |
| client\_latency | sum\_client\_latency\_ms avg\_client\_latency\_ms max\_client\_latency\_ms sum\_flows p98th\_bits\_per\_sec p98th\_pkts\_per\_sec |
| server\_latency | sum\_server\_latency\_ms avg\_server\_latency\_ms max\_server\_latency\_ms sum\_flows p98th\_bits\_per\_sec p98th\_pkts\_per\_sec |
| appl\_latency | sum\_appl\_latency\_ms avg\_appl\_latency\_ms max\_appl\_latency\_ms sum\_flows p98th\_bits\_per\_sec p98th\_pkts\_per\_sec |
| fps | avg\_flows\_per\_sec p95th\_flows\_per\_sec max\_flows\_per\_sec p95th\_bits\_per\_sec p95th\_pkts\_per\_sec |
| unique\_src\_ip | max\_ips p95th\_bits\_per\_sec p95th\_pkts\_per\_sec |
| unique\_dst\_ip | max\_ips max p95th\_bits\_per\_sec p95th\_pkts\_per\_sec |

  

### Query JSON as cURL

Here’s a JSON example for a Query Chart call in the data (-d) parameter of a cURL HTTP request:

```
curl -X POST \
-H 'X-CH-Auth-Email: user@domain.suffix' \
-H 'X-CH-Auth-API-Token: user_api_token' \
-H 'Content-Type: application/json' \
-d '{ "queries": [ { "query": { "viz_type": "stackedArea", "show_overlay": true, "overlay_day": -7, "sync_axes": false, "metric": "bytes", "matrixBy": [], "cidr": 32, "cidr6": 128, "pps_threshold": 500, "topx": 8, "depth": 25, "fastData": "Auto", "outsort": "max_bits_per_sec", "lookback_seconds": 300, "time_format": "UTC", "starting_time": null, "ending_time": null, "device_name": "device_names", "all_selected": false, "filters_obj": {}, "saved_filters": [], "descriptor": "Total", "query_title": "Top Source Country by Max Bits/s", "time_type": "relative", "device_id": "2951", "num_selected": 0, "sql_string": "", "queryChanged": false, "granularity": "total", "ppsThreshold": "500", "last_run": null, "aggregates": [ { "name": "avg_bits_per_sec", "column": "f_sum_both_bytes", "fn": "average", "sample_rate": 1 }, { "name": "p95th_bits_per_sec", "column": "f_sum_both_bytes", "fn": "percentile", "rank": 95, "sample_rate": 1 }, { "name": "max_bits_per_sec", "column": "f_sum_both_bytes", "fn": "max", "raw": true, "sample_rate": 1 } ], "dimension": [ "Geography_src" ] }, "bucket": "Left +Y Axis", "bucketIndex": 0, "isOverlay": false } ], "imageType": "pdf"}' \
https://api.kentik.com/api/v5/query/topXchart | python -m json.tool
```

> **Notes:**
>
> * Replace placeholders (in italics) with actual values: your email address (`user@domain.suffix`), API token (`user_api_token`), and device names (`device_names`, comma-delimited list) for the query.
> * The JSON for a Query Data call is similar but omits some fields. See **Query API Request JSON** for details.
> * For more information on using Kentik APIs with cURL, see **API Access Via cURL**.

### Query API Response JSON

The Query Data method returns JSON data arrays in the HTTP response body, while the Query Chart method returns an image file. The array data comes from the main tables of the Kentik Data Engine (see **KDE Tables**), which store flow records and related network data (BGP, GeoIP, SNMP). The arrays from the two methods differ, even with the same query settings, because Query Data results are processed for improved accuracy, such as in averaging queries.

> **Note:** Example response JSON is provided in the topic for each method.

## Query Chart Method

The Query Chart `POST` method queries Kentik's KDE and returns image data for a graph like those in Data Explorer (see **Data Explorer Chart**). The image data is returned as a JSON object with a `dataURI` member containing a base64 encoded string.  
  
Supported image types are SVG, PDF, PNG, or JPG. SVG is the default and recommended type because it returns quickly, scales without data loss, and can be easily integrated by injecting it directly into the DOM.  
  
**HTTP Request**  
The following table shows the path and HTTP request for the Query Chart method:

|  |  |
| --- | --- |
| **URL** | https://api.kentik.com/api/v5/query/topXchart |
| **Request** | POST /api/v5/query/topXchart HTTP/1.1 Host: api.kentik.com X-CH-Auth-Email: user@domain.suffix X-CH-Auth-API-Token: user\_api\_token Content-Type: application/json |

> **Note:** If your organization is registered on Kentik's EU cluster, use api.kentik.eu in place of api.kentik.com in the URL.

The request body includes a JSON array with elements containing the information needed by Kentik to perform the query. For details on creating this array and its components, see **Query API Request JSON**.

**HTTP Response**  
A successful response from the Query Chart method includes the following elements:

* Response headers
* HTTP response code
* JSON response body with a `dataURI` member

The following example shows the structure of the returned JSON (placeholder highlighted):

```
{
"dataUri": "base64_encoded_string"
}
```

## Query URL Method

The Query URL `POST` method provides a URL to open **Data Explorer** in the Kentik portal, with the sidebar according to the request JSON query parameters and the query results visualized in the display area.  
  
**HTTP Request**  
The following table shows the path and HTTP request for the Query URL method:

|  |  |
| --- | --- |
| **URL** | https://api.kentik.com/api/v5/query/url |
| **Request** | POST /api/v5/query/url HTTP/1.1  Host: api.kentik.com  X-CH-Auth-Email: user@domain.suffix  X-CH-Auth-API-Token: user\_api\_token  Content-Type: application/json |

> **Note:** If your organization is registered on Kentik's EU cluster, use `api.kentik.eu` in place of `api.kentik.com` in the URL above.

The request body contains a JSON array with elements detailing the query settings for Data Explorer. For details on creating this array and its components, see **Query API Request JSON**.

> **Note:** This request body is identical to that of the **Query Chart Method**, except it excludes the `imageType` parameter.

**HTTP Response**

A successful response from the Query URL method includes the following elements:

* Response headers
* HTTP response code
* Response body with a short URL in quotes

The following example shows an example of a URL returned from the Query URL method:

`https://portal.kentik.com/portal/#Charts/shortUrl/d3f24034b386ebf9e8ad6de8f0f715a1`

## Query Data Method

The Query Data `POST` method executes a query on Kentik's KDE data (see **KDE Tables**) and returns `results` in a JSON results array.  
  
**HTTP Request**  
The following table shows the path and HTTP request for the Query Data method:

|  |  |
| --- | --- |
| **URL** | https://api.kentik.com/api/next/v5/query/topXdata |
| **Request** | POST /api/next/v5/query/topXdata HTTP/1.1  Host: api.kentik.com  X-CH-Auth-Email: user@domain.suffix  X-CH-Auth-API-Token: user\_api\_token  Content-Type: application/json |

> **Note:** If your organization is registered on Kentik's EU cluster, use `api.kentik.eu` in place of `api.kentik.com` in the URL above.

The request body contains a JSON array with elements specifying the necessary information for Kentik to execute the query. For details on creating this array and its components, see **Query API Request JSON**.  
  
**HTTP Response**  
A successful response from the Query Data method includes the following elements:

* Response headers
* HTTP response code
* JSON response body with a `results` array, which varies based on the query provided.

This example displays a `results` array from a query on top source country by max bits per second, showing two results (actual results range from 25 to 350, based on the `depth` field in the request JSON):

* The first result (US) includes five-minute time-series data (`lookback_seconds` set to 300) and represents top results based on the topx field in the request JSON (see **Query API Request JSON**).
* The second result (AU) illustrates non-topx results.

```
{
"results": [
    {
      "bucket": "Left +Y Axis",
      "data": [
        {
          "Geography_src": "US",
          "avg_bits_per_sec": 130548781.51111111,
          "key": "US",
          "max_bits_per_sec": 153616657.06666666,
          "p95th_bits_per_sec": 150816290.13333333,
          "timeSeries": {
            "both_bits_per_sec": {
              "flow": [
                [
                  1482186900000,
                  128568524.8,
                  60
                ],
                [
                  1482186960000,
                  133069482.66666667,
                  60
                ],
                 [
                   1482187020000,
                   115999675.73333333,
                   60
                 ],
                [
                  1482187080000,
                  153616657.06666666,
                  60
                ],
                 [
                   1482187140000,
                   104022425.6,
                   60
                 ],
                [
                  1482187258438,
                  148015923.2,
                  60
                ]
               ]
             }
           }
         },
         {
           "Geography_src": "AU",
           "avg_bits_per_sec": 732114.4888888889,
           "key": "AU",
           "max_bits_per_sec": 1145924.2666666666,
           "p95th_bits_per_sec": 1003315.2
         }
      ]
    }
  ]
}
```

## Query SQL Method

As of May 1st, 2025, the Query SQL Method (API endpoint) has been deprecated and is no longer supported.
