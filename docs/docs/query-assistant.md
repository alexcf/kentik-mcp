# Query Assistant

Source: https://kb.kentik.com/docs/query-assistant

---

This article covers the Query Assistant page in Kentik's Network Monitoring System (NMS).

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(450).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A27Z&se=2026-05-12T09%3A29%3A27Z&sr=c&sp=r&sig=wPOmFsNejWrQe44YbJro7cHtCSCnIO1cZBFDBwGSvBI%3D)

*Query Assistant enables you to ask questions in natural language.*

> **Notes:**  
>  • To use Query Assistant, Kentik AI must be enabled (see **Manage Kentik AI**and **Kentik AI Overview**).  
>  • The following articles provide information about other aspects of Kentik's Network Monitoring System:
>
> * **Network Monitoring**
> * **Kentik NMS Configuration**
> * **NMS Setup**
> * **NMS Dashboard**
> * **Metrics Explorer**
> * **NMS Devices**
> * **NMS Interfaces**

## About the Query Assistant Page

The Query Assistant page in Kentik's Network Monitoring System (NMS) enables you to enter a query in your own words rather than specifying it with the controls of the **Query** sidebar in **Metrics Explorer**. This allows you to generate a query quickly without delving into the details of the fields in the **Query** sidebar in Metrics Explorer, after which you can refine the query by tweaking the settings as needed. While the Query Assistant has been developed primarily in English, it is able to understand entries in some other languages as well.  
  
Similar to the operation of the field described in **Metrics Query Assistant**, the text you enter into the field across the top of the Query Assistant page will be used by Kentik to formulate a query that matches our understanding of what you'd like to see. The Query Assistant page also includes four cards that each provide readymade natural-language queries in a given category (devices, interfaces, routing, and hardware components).

## Query Assistant Page Layout

The Query Assistant page includes the following main elements:

* **Query field**: An input field into which you can enter, in natural language, text describing what you'd like your query results to show you (e.g., "show me interface utilization using line chart with "PEER" in the description"). To submit the entered text for conversion into a query, press **Return/Enter** on your keyboard or click the **Send** button at the right of the field.
* **Example Queries**: A set of cards that each correspond to an area of your infrastructure about which you can query metrics: devices, interfaces, routing, and hardware components. Each card contains:

  + Example query chart: A chart showing the results from an example query.
  + Example queries list: A list of additional example queries. Click on a listed query to enter it to the **Query** field, then run it by pressing **Return/Enter** on your keyboard or clicking the **Send** button at the right of the field.

After the query runs you'll be taken to **Metrics Explorer**, where:

* The string from the Query Assistant page appears in the **Query** field.
* The main display shows the results of the query.
* The **Query** sidebar has settings reflecting Kentik's interpretation of the entered text. To refine the query, change these settings and click **Run** at the bottom of the sidebar.

> **Note:** Query Assistant is still a Beta feature. Use the **Query** sidebar in Metrics Explorer to make adjustments to better reflect your intent (see **Metrics Query Assistant**).
