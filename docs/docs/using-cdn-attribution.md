# Using CDN Attribution

Source: https://kb.kentik.com/docs/using-cdn-attribution

---

The use of CDN attribution in Kentik is covered here.

> **Note:** To use CDN attribution, at least one of your DNS servers must be running Kentik's Universal Agent with the **DNS OTT Tap Capability** enabled (see **Universal Agents**).

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(195).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A13Z&se=2026-05-12T09%3A32%3A13Z&sr=c&sp=r&sig=vfw51hqjPcD4dD1r9Gdi7zHr4bL%2Bnd9R4CxEB%2B27oPU%3D)

*Running a CDN query in Data Explorer.*

## About CDN Attribution

Mapping IPs to CDNs is challenging in practice due to these two infrastructure scenarios:

* **CDN-owned Infrastructure**: Each CDN runs its own ASN(s), with their own PoPs running their own cache servers
* **ISP-embedded infrastructure**: Many CDNs rely on ISP-embedded caching servers for better last mile performance, typically by directing users of a given ISP to CDN nodes co-located at the local ISP, closer to the end-users.

Different CDNs blend these scenarios, e.g., commercial and multi-tenant CDNs vs. single-purpose CDNs. Kentik’s dynamic algorithms allow IP mapping across the spectrum: a base mapping lists ASNs by CDN and their originated IP ranges, while a more dynamic layer uses DNS traffic (upon ISP agreement) to deduce the remaining mappings.

Kentik’s CDN attribution system is self-learning and updates daily. It continuously discovers new CDNs and IPs (caching servers) as they emerge.

![A simplified visualization of Kentik's CDN attribution.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/CDN-Overview_diagram-605h1124w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A13Z&se=2026-05-12T09%3A32%3A13Z&sr=c&sp=r&sig=vfw51hqjPcD4dD1r9Gdi7zHr4bL%2Bnd9R4CxEB%2B27oPU%3D)

*A simplified visualization of Kentik's CDN attribution.*

## Enabling CDN Attribution

Preparing your Kentik setup for CDN attribution involves the following tasks:

* **Deploy the Universal Agent** software on at least one of your DNS servers (see **Universal Agents**).
* Enable the **DNS OTT Tap Capability** on each agent.
* Configure the corresponding device (in the Kentik portal or via **Device API**) to send DNS data to Kentik.

To configure CDN attribution in the Kentik portal:

1. Choose Settings » **Networking Devices** from the main nav menu to open the Devices page.
2. Open the Device settings dialog:

   1. To register a new DNS server with Kentik, click the **Add Device** button.
   2. To change the CDN attribution settings of an existing DNS server, click the Edit button (pencil icon) at the right of the device’s row in the **Device List**.
3. If you're registering a new device, on the **General** tab of the dialog set the **Flow Device Type** field to **Kentik****Host Agent**.
4. On the **General** tab, turn on the **Contribute to CDN Attribution** switch.
5. Set the remaining settings on the dialog's tabs as needed (see **Device Settings**).
6. Click **Save** button to save changes and return to the Devices page. If you added a new device, it will now be shown in the Device list.

> **Note:** DNS servers covering different geographical zones usually result in distinct, non-overlapping, IP-to-CDN mappings. Kentik recommends deploying **Universal Agents** on as many DNS servers as possible to export the best IP-to-CDN mapping data to KDE. This approach also provides granular NPM metrics for those devices.

## Applying CDN Attribution

Once you've registered one or more DNS servers with Kentik and configured them for CDN attribution, the flow records stored for those devices in KDE (see **KDE Tables**) will include **CDN Attribution Dimensions**. You will now be able to use those dimensions in Kentik queries:

* To use source or destination CDN as a group-by dimension, see **Dimension Selectors**.
* To use source or destination CDN as a filter (as shown in the screenshot below), see **Filters Pane**.

![Filtering options for selecting results based on destination and conditions in a user interface.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/UK-CDN_Attribution.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A13Z&se=2026-05-12T09%3A32%3A13Z&sr=c&sp=r&sig=vfw51hqjPcD4dD1r9Gdi7zHr4bL%2Bnd9R4CxEB%2B27oPU%3D)

Using CDN for filters or group-by dimensions in Data Explorer or Dashboard queries can reveal (among other things) how much of your traffic is coming from (shown in screenshot below) or going to various CDNs.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(199).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A13Z&se=2026-05-12T09%3A32%3A13Z&sr=c&sp=r&sig=vfw51hqjPcD4dD1r9Gdi7zHr4bL%2Bnd9R4CxEB%2B27oPU%3D)

### CDN Attribution Dimensions

CDN attribution makes it possible for Kentik to determine whether a given flow record originated or terminated with a commercial CDN, and to store that information for each record using the following two columns of the Kentik Data Engine (KDE):![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(196).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A13Z&se=2026-05-12T09%3A32%3A13Z&sr=c&sp=r&sig=vfw51hqjPcD4dD1r9Gdi7zHr4bL%2Bnd9R4CxEB%2B27oPU%3D)

* `src_cdn`: The commercial name of the CDN derived from the source IP (`inet_src_addr`) of an ingested flow.
* `dst_cdn`: The commercial name of the CDN derived from the destination IP (`inet_dst_addr`) of an ingested flow.

  > **Note:** This dimension enables you to track "fill traffic" that is pointed toward a CDN server to fill a local cache.

Once stored in KDE (see **KDE Tables**), the columns can be used for both group-by dimensions and filters in Kentik queries (e.g., in Data Explorer, Dashboards) as described in **Applying CDN Attribution**.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(197).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A13Z&se=2026-05-12T09%3A32%3A13Z&sr=c&sp=r&sig=vfw51hqjPcD4dD1r9Gdi7zHr4bL%2Bnd9R4CxEB%2B27oPU%3D)

##
