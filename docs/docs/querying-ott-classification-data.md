# Querying OTT Classification Data

Source: https://kb.kentik.com/docs/querying-ott-classification-data

---

While the **OTT Service Tracking** dashboards group classification data into broad UI categories (such as **Fully Classified** or **Pending**), Kentik's **Data Explorer** allows for much more granular querying.

By using the **OTT Classification Type** dimension in Data Explorer, advanced users and API developers can filter or group traffic based on the exact deterministic or probabilistic matching behavior the True Origin engine used to classify the flow.

## OTT Classification Type Dimension Values

When building a query in **Data Explorer**, the **OTT Classification Type** dimension returns the specific method used to identify the traffic. The table below maps the underlying dimension values to their corresponding high-level dashboard categories:

| UI Category | Dimension Value | Description |
| --- | --- | --- |
| **Fully Classified (Deterministic)** | `Fixed (Full Match)` | Traffic matched via fixed, known IP addresses associated with a specific service. |
|  | `Direct (Full Match)` | Traffic matched via a direct DNS query to a known service hostname. |
|  | `Indirect (Full Match)` | Traffic matched via an indirect DNS query (e.g., CNAME or alias) to a known service hostname. |
| **Intelligently Classified (Probabilistic)** | `Projected (Full Match)` | Traffic matched probabilistically using Kentik's Intelligent Classification statistical model, based on known multi-service IP distributions. |
| **Provider-Only** | `Indirect (Provider-only Match)` | Traffic matched to an OTT Provider, but the specific OTT Service cannot be determined (typically because multiple services from the same provider share the hostname). |
| **Pending Classification** | `Direct (Pending Classification)` | Traffic matched to a direct DNS query for a high-volume hostname that True Origin is actively evaluating for a future service mapping. |
|  | `Indirect (Pending Classification)` | Traffic matched to an indirect DNS query for a high-volume hostname that True Origin is actively evaluating for a future service mapping. |
| **Unclassified** | `Unknown` | Traffic that lacks matching DNS, IP patterns, or statistical models, resulting in no identification. |

## Example Use Cases

Querying by the exact **OTT Classification Type** is particularly useful when auditing traffic classifications or building custom reports. For example:

* **Auditing Statistical Data**: You can build a **Data Explorer** query filtered strictly by the `statistical_full` dimension value to isolate and evaluate only the traffic inferred by Intelligent Classification, excluding all deterministic DNS traffic.
* **Identifying Future Mappings**: Grouping traffic by the `direct_ad_hoc` and `indirect_ad_hoc` dimension values allows you to see exactly which high-volume, unmapped hostnames on your network are currently pending classification by the True Origin engine.
