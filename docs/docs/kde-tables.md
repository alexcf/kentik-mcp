# KDE Tables

Source: https://kb.kentik.com/docs/kde-tables

---

This article covers the tables of the **Kentik Data Engine** (KDE).

The **Kentik Data Engine**(KDE) is the distributed columnar database at the heart of the Kentik platform. It stores raw flow records and metadata in discrete tables, which are then exposed as dimensions for analysis in the Kentik portal.

> **CRITICAL:** As of May 1st, 2025, the **Query SQL Method**, which was the only way to access the data described here, has been **deprecated**.

## Architecture Overview

KDE maintains discrete databases for customer flow records. These include:

* **Main Tables:** Individual tables storing flow records and associated metadata for each specific device (router, host, etc.).
* **Supplemental Data:** Derived datasets that enrich raw flow with context.
* **All Devices Table:** A unified view that merges data from every device in your organization into a single queryable source (see **All Devices Table**).

The columns in these tables form the basis for Kentik **Dimensions**. Most are available in the **Dimension Selector** for filtering and group-by operations within the Portal.

## Universal Data Records (UDR)

To support a vast array of networking hardware without a rigid schema, Kentik uses Universal Data Records.

Rather than having thousands of empty columns for every possible vendor field, UDR dynamically maps fields based on the device type. This allows Kentik to ingest vendor-specific data (e.g., Cisco-specific flow fields) alongside standard NetFlow or sFlow without losing granularity.

**Key Benefits of UDR**:

* **Scalability**: Supports disparate sources like Cloud flow logs and traditional SNMP data.
* **Customization**: Powering **Custom Dimensions**to tailor data to your business logic.
* **Extensibility**: Rapidly adds support for new **Device-Specific Dimensions**.

## The All Devices Table

The `all_devices` table is the most common data source for global network visibility. It contains every field found in individual device tables, plus specialized metadata fields used for cross-device correlation.

### Core Identity Fields

These fields are essential for identifying the source of traffic across your entire infrastructure:

| Column | Type | Description |
| --- | --- | --- |
| `i_device_id` | text | The unique numerical ID assigned by Kentik. Best for precise API filtering. |
| `i_device_name` | text | The user-defined name of the device. |
| `i_device_site_name` | text | The **Site** assigned to the device. (Empty if unassigned). |
| `i_device_type` | text | The category of the source (e.g., router, host, or cloud gateway). |

> **Note**: Site assignments in the KDE may lag behind Admin settings by up to 10 minutes. The `i_device_type` field is optimized for filtering (WHERE clauses) rather than grouping.

---

## See Also

* **Kentik Data Engine Overview**
* **Dimensions Reference**
* **Supported Device Types**
