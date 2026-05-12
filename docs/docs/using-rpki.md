# Using RPKI

Source: https://kb.kentik.com/docs/using-rpki

---

This article covers the use of RPKI in Kentik.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/RPKI-With-483h674w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A01Z&se=2026-05-12T09%3A40%3A01Z&sr=c&sp=r&sig=8aO8DjVjftWWvbxiHBohWn7V3spHq9kDAJYFcPpnodQ%3D)

*An image from the Cloudflare blog post* **RPKI and BGP**.

## About RPKI

Resource Public Key Infrastructure is a system that's used to validate the BGP routes announced by an Autonomous System by verifying that the AS is authorized to originate the prefixes in an announced route. The overarching goal is to improve the reliability of the information that networks rely on when filtering routes, enabling them to identify and drop BGP routes when the announced prefixes aren't part of the address space assigned to the AS making the announcement. In this way RPKI effectively does for BGP what DNSSEC does for DNS.  
  
RPKI validation addresses serious vulnerabilities in the current implementation of BGP, in which any BGP-speaking network can announce itself (deliberately or not) as the origin of any prefix. This results, not infrequently, in “BGP prefix hijack,” in which invalid prefixes are announced throughout the Internet, often making part of the Internet temporarily unreachable. Without RPKI, the only option to protect your network against the propagation of BGP hijacks and other erroneous routing information is to parse the Internet Routing Registry and hope that the IRRs are correct. As shown in the diagram below from Cloudflare, when that doesn't work traffic can be sent to an incorrect or invalid destination.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/RPKI-Without-324h674w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A01Z&se=2026-05-12T09%3A40%3A01Z&sr=c&sp=r&sig=8aO8DjVjftWWvbxiHBohWn7V3spHq9kDAJYFcPpnodQ%3D)RPKI is built on a set of **IETF RFCs** that define Public Key Infrastructure (RFC5280) and specify various elements of RPKI implementation (including RFC6410, RFC6411, RFC6480, RFC6482, and RFC8481). The underlying idea is to employ the PKI mechanism, which involves public-key encryption, to certify the validity of announced routes. Since the same entities that assign resources (specifically IP addresses) to ASes are best positioned to determine if a given AS is the authorized "holder" of a given resource, the Regional Internet Registries serve as certificate authorities (referred to as "Trust Anchors" or TAs) for RPKI encryption. The infrastructure via which RPKI is supported involves the main building blocks described in the topics below.

### RPKI Trust Anchor

As shown below (image from Cloudflare) the five RIRs — AFRINIC, APNIC, ARIN, LACNIC, and RIPE — act as TAs for RPKI.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/RPKI-Trust_anchors-327h674w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A01Z&se=2026-05-12T09%3A40%3A01Z&sr=c&sp=r&sig=8aO8DjVjftWWvbxiHBohWn7V3spHq9kDAJYFcPpnodQ%3D)TAs are responsible for several elements of the RPKI process:

* They provide a mechanism, called a **Route Origin Authorization** or ROA, by which ASes can sign their routes. Declaring that it is the authorized assignee of a given block of IP prefixes means that an AS has permission to advertise routes to one or more prefixes within the block.
* They maintain a "collection" of ROAs, which is effectively a database that maps signed prefixes to authorized origin ASes.
* They operate servers that host the ROA databases. "Relying parties" (e.g. network operators) regularly connect to these servers to download the latest database versions to their own local caches. The idea is that routers within each AS will query the local cache to obtain the information needed to validate BGP announcements by matching a given prefix with an origin ASN. Each ROA in a local cache is referred to as a "Validated ROA Payload" (VRP).

### Route Origin Authorization

A Route Origin Authorization (ROA) is a record that associates a route with an originating AS number. The core rationale for ROAs is summarized in this passage from RFC6480: "BGP is based on the assumption that the AS that originates routes for a particular prefix is authorized to do so by the holder of that prefix (or an address block encompassing the prefix)... A Route Origination Authorization (ROA) makes such authorization explicit, allowing a holder of IP address space to create an object that explicitly and verifiably asserts that an AS is authorized to originate routes to a given set of prefixes."  
  
The five RIRs provide a method by which ASes can create an ROA for a given IP/ASN pair and encrypt it by signing with a private key. As shown in this example from Kentik, an ROA includes a set of attributes in the form of key:value pairs:

```
{
  "prefix": "141.193.36.0/22",
  "maxLength": 24,
  "asn": "AS6169",
  "ta": "Cloudflare - ARIN"
}
```

The key:value pairs in the ROA pass the following information:

* **prefix**: The prefix that is being associated with an ASN in this ROA, which is the set of IP address prefixes to which the AS is authorized to originate routes. This field may contain multiple authorized prefixes.
* **maxLength**: The highest mask that will be valid for the prefix. In the example above, for instance, all prefixes contained within `141.193.36.0/22` with netmasks equals to 22, 23, and 24 will be considered valid, while a prefix contained within `141.193.36.0/22` that has a /25 mask would be considered invalid, as would a less specific prefix (e.g., /21).
* **asn**: A single certified origin ASN for the AS that is authorized to originate routes to the IP addresses in the prefix field (AS6169 is Kentik’s ASN). This allows a route to be validated by comparing the origin ASN of a received prefix to the authorized holder of that resource. A discrepancy indicates that the route has been hijacked.
* **ta**: The Trust Anchor from which the ROA comes. In the above example, the ROA was pulled by Cloudflare from the ROA list maintained by ARIN (which is why the TA name is prefixed with "Cloudflare").

### RPKI Validation State

An RPKI validation state is the result of an attempt to confirm a valid match between a prefix and an ASN. Three possible values are defined in the specification (RFC6411):

* **Valid**: At least one VRP (Validated ROA Payload) matches the route prefix, meaning that the route prefix length is less than or equal to the VRP maximum length, and the route origin ASN is equal to the VRP ASN.
* **Invalid**: At least one VRP "covers" the route prefix, meaning that the route prefix is either identical to or more specific than the VRP prefix, but no VRP matches the route prefix.
* **Not Found** (unknown): No VRP covers the route prefix (the route prefix is neither identical to the VRP prefix nor more specific than the VRP prefix).

## RPKI in Kentik

As described in **Kentik BGP Collection**, Kentik is able to collect BGP routing information and associate that data with individual flows as they are ingested into the Kentik Data Engine (KDE). With the addition of RPKI support, Kentik is now also able to enrich flow records with the RPKI validation state, which enables you to determine the following:

* Is traffic is being originated by ASNs that don’t own them (i.e. hijacks)?
* Is traffic being announced by the correct origin AS but with an invalid (too specific) netmask?
* What flows would be dropped if your Kentik-connected routers were configured for strict route validation?

### RPKI at Ingest

Kentik currently supports RPKI via Cloudflare's GoRTR, which uses a centralized RPKI validator that consolidates the five RIR ROA databases. A list of valid routes from this central validator is securely distributed via Cloudflare's own CDN using a lightweight local RTR server. At ingest we consult our local cache of the GoRTR data, determine the RPKI validation state associated with the router sending a flow, and use that state to derive the RPKI values that we assign to RPKI columns in the KDE flow record.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(192).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A01Z&se=2026-05-12T09%3A40%3A01Z&sr=c&sp=r&sig=8aO8DjVjftWWvbxiHBohWn7V3spHq9kDAJYFcPpnodQ%3D)

### RPKI Dimensions

Kentik uses two dimensions for RPKI values, both of which relate to the validation state of a prefix in a BGP-advertised route:

* **RPKI Quick Status**: A simple status based on the **RPKI Validation State**, which indicates whether the route would be used or dropped if the router were configured to enforce strict route validation.
* **RPKI Validation Status**: The validation state plus, if the state is Invalid, an explanation of the cause of the invalid state, e.g., "prefix length out of bounds."

The RPKI Quick Status dimension enables filtering on the following values:

* **RPKI Unknown**: No ROA found. Since there's no way to be sure that the AS announcing the route is not the valid holder of the corresponding prefixes, a route validator would not force the router to drop the route.
* **RPKI Valid**: An ROA was found, certifying that the AS announcing the route is authorized to do so. The legitimacy of the origin ASN is not in doubt and a route validator would not force the router to drop the route.
* **RPKI Invalid - covering Valid/Unknown**: Route is invalid but would not be dropped; see **RPKI Invalid Values**.
* **RPKI Invalid - will be dropped**: Route is invalid and would be dropped.

Using these two dimensions, you can filter (e.g., in Data Explorer) on RPKI Quick Status to return a table of the top-X flows marked with either of the "invalid" values, then check the RPKI Validation Status column of the table to see the specific issue.

#### RPKI Dimension Mapping

The table below shows how the values for the RPKI Validation Status dimension are mapped to values for the RPKI Quick Status dimension, and what the result would be with strict validation enabled.

| RPKI Validation Status | RPKI Quick Status | Strict Validation Behavior |
| --- | --- | --- |
| RPKI Unknown | RPKI Unknown | Will not be dropped |
| RPKI Valid | RPKI Valid | Will not be dropped |
| RPKI Invalid - Valid Covering Prefix RPKI Invalid - Unknown Covering Prefix | RPKI Invalid - Covering Valid/Unknown | Will not be dropped |
| RPKI Invalid - Prefix Length Out of Bounds RPKI Invalid - Incorrect Origin ASN RPKI Invalid - Explicit ASN 0 | RPKI Invalid - Will be dropped | Will be dropped |
| No value | No value | Undetermined behavior:  - The prefix may be in a static route  - The prefix may be a /32 or /31  - No AS\_Path info available |

### RPKI Invalid Values

As described above, when the value of the RPKI Validation Status dimension is RPKI Invalid the state can be determined more specifically by checking the RPKI Validation Status values, which include the following:

* **RPKI Invalid: valid covering prefix**: Kentik has determined that there is a valid secondary BGP announcement covering the route. With strict route validation enabled the initial prefix would be dropped but the covering RPKI-valid prefix would prevent interruption of traffic.
* **RPKI Invalid: unknown covering prefix**: Kentik has determined that there is a secondary BGP announcement in which the validation state of the route is unknown. With strict route validation enabled the route will not be dropped.

  > **Note:** Until RPKI is widely adopted and everyone signs all of their prefixes, the vast majority of traffic over the Internet will have an RPKI state of unknown (ROA not found).
* **RPKI Invalid: prefix length out of bounds**: Kentik has determined that the prefix length is outside of the range specified in the corresponding ROA.  
  **Example**: The ROA specifies a `prefix` of `141.193.36.0/22` and a `maxLength` of `24`. A flow whose prefix is `141.193.26.0/25` would be out of bounds.
* **RPKI Invalid: incorrect Origin ASN (should be AS<AS\_Number>)**: The origin ASN specified in the route doesn’t match the ASN specified by the ROA, indicating an apparent route hijack. No secondary announcement has been found, so if strict route validation is enabled then traffic corresponding to this prefix will be dropped.
* **RPKI Invalid: explicit ASN 0**: The RPKI standard allows TAs, and anyone else who leverages RPKI, to inject artificial ROAs that contain an origin ASN value of 0, enabling you to statically define prefixes that shouldn’t be trusted at all. An ROA with ASN = 0 means that any traffic coming from that prefix, and all prefixes contained within it, as defined by `maxLength`, will be considered invalid.

## Applying RPKI

As network operators consider adopting strict route validation it's critical to determine what traffic would end up being dropped. As mentioned earlier, Kentik's two RPKI dimensions can be used to figure that out:

1. In the **Devices** pane inData Explorer, check that the selected devices include routers.
2. In the **Filters** pane, set an ad hoc filter (see **Ad Hoc Filter Controls**) in which RPKI Quick Status equals "RPKI Invalid - Will be dropped". Query results will show only traffic that would be dropped by strict validation.

   ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(193).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A01Z&se=2026-05-12T09%3A40%3A01Z&sr=c&sp=r&sig=8aO8DjVjftWWvbxiHBohWn7V3spHq9kDAJYFcPpnodQ%3D)
3. In the **Query** pane, open the **Group by Dimensions** selector (see **Choosing Query Dimensions**) and select the following dimensions:![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(712).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A01Z&se=2026-05-12T09%3A40%3A01Z&sr=c&sp=r&sig=8aO8DjVjftWWvbxiHBohWn7V3spHq9kDAJYFcPpnodQ%3D)

   1. **Site**: Group results by site (see **About Sites**).
   2. **RPKI Validation Status**: The **RPKI Invalid Values** of the dropped traffic.
   3. **Destination Route Prefix/LEN**: The prefixes that will be dropped, so you can further investigate individual prefixes.
   4. **Destination AS Number**: The destination ASN in the BGP announcement, which may differ from the RPKI-certified origin ASN for this prefix.
4. Run the query.

   ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(194).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A01Z&se=2026-05-12T09%3A40%3A01Z&sr=c&sp=r&sig=8aO8DjVjftWWvbxiHBohWn7V3spHq9kDAJYFcPpnodQ%3D)

If a route is invalid because of an incorrect origin ASN, you'll be able to see this in the query results table (eventually you'll also be able to create an alert policy for this situation). As shown above, for example, when the query was run it resulted in a mismatch between the ASN information in two of the columns in the first row of the results table:

* The expected origin ASN (as based on ROA) in the RPKI Validation Status column is 266852.
* The announced origin ASN that Kentik correlated with the flow at ingest, which is shown in the Destination ASN column, is 61503.

Based on the above, if strict route validation were enabled the routers at the site ATL1 would conclude that AS 61503 has hijacked (either intentionally or by misconfiguration) prefixes whose actual resource holder is AS 266852, and traffic on the corresponding route would be dropped.  
  
If you wanted to dig deeper into the issue, perhaps to gather information to pass along to the involved parties (the official origin ASN and/or the Transit Provider conveying the hijack), you could click the **Action** menu at the right of the table row, choose **Show by**, and choose AS PATH in the resulting Show By Dimensions dialog. As shown below, the results table would then indicate the ASNs in the path.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(711).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A01Z&se=2026-05-12T09%3A40%3A01Z&sr=c&sp=r&sig=8aO8DjVjftWWvbxiHBohWn7V3spHq9kDAJYFcPpnodQ%3D)
