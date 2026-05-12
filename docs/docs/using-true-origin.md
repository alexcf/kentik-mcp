# Using True Origin

Source: https://kb.kentik.com/docs/using-true-origin

---

This article discusses the use of **True Origin** in Kentik.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/TO-OTT_pie-424h674w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A54Z&se=2026-05-12T09%3A43%3A54Z&sr=c&sp=r&sig=oxl8wha8RpTIybmOiJyEqqgDSsgqqTtPmT9otc8XVtY%3D)

*A pie chart showing streaming service usage percentages, highlighting YouTube and Hulu.*

## Why True Origin?

In the topics below we'll explain what True Origin is, what operational needs it addresses, how it works, and the steps involved in putting it to use. By the end of this article you should have a pretty good idea of how True Origin helps network operations and engineering, particularly for Internet Service Providers. The following capabilities provide a taste of how useful True Origin can be:

* Get an accurate picture of what content is valuable to your subscribers, how that content is delivered, and the impact on your network.
* Identify hurdles in delivering content to subscribers and resolve them for a better subscriber experience.
* Dive deeper into embedded cache setups and better support their deployment and operation.

To understand how True Origin helps us with the above scenarios, let's start by putting ourselves in the place of an (eyeball) ISP with a large number of subscribers that consume content from the Internet. The content comes into the ISP network from various sources via various paths with varying cost and performance. For the ISP's executive and engineering teams, these variations increase the complexity of meeting two main operational goals:

* Ensure the high throughput required for a good user experience.
* Maximize cost-efficiency.

To succeed at these goals, an ISP needs to define a robust, performant architecture and also a set of best practices. To do that, you need real-time, high quality information about two factors: the traffic over your network and the cost to you of carrying that traffic. Until now, there hasn't been a single tool providing the required visibility — in the same place at the same time — into both of these factors. That's why True Origin is a major step forward for every ISP carrying a significant volume of Internet-sourced content.

## What is True Origin?

By correlating an ever-wider set of traffic data into a single, instantly queryable dataset (the Kentik Data Engine, a.k.a. KDE), Kentik is able to generate technical and business insights with direct, powerful relevance to your network operations. True Origin falls right in line with this mission, building on a set of key Kentik features that are geared towards eyeball ISPs:

* **CDN tagging**: Most Internet-based content is sourced from one of approximately 45 commercial and purpose-built CDNs. Kentik has built an engine that maintains a curated list of the source and destination IPs associated with these CDNs, so our flow records from their traffic include not just source or destination IP address but also the actual CDN names.
* **OTT tagging**: Because commercial CDNs deliver traffic for multiple content providers, OTT services can't be identified based on the server IP in a flow record. Instead, we need to be able to match the content provider from which a subscriber has requested content with the subsequent download to that subscriber. To do that we developed OTT tagging, a powerful DNS-based methodology that enables us to identify the OTT service behind a given flow regardless of the CDN, hosting provider, or connectivity type utilized for the delivery of that service's content.
* **Subscriber tagging**: A subscriber tag is a reference to the individual subscriber that is at the generating or receiving end of a flow. By incorporating subscriber information into each KDE flow record we can provide ISPs with the next level of flow visibility.

#### Questions True Origin Can Answer

The above features are each powerful by themselves; working together they uniquely address the questions faced daily by our ISP partners. The main goal is to provide actionable, human-readable answers about the relationship between specific content and the methods by which that content is delivered to subscribers. Here’s a look at some ISP questions that True Origin can help answer:

* Is there an ongoing content-related traffic event? Is it linked to a specific content provider?
* Does the root cause of a given issue lie with the content provider, the CDN, or my own network?
* What entities does specific content pass through on its way to my subscriber? What are the economics of that route?
* How does my throughput ranking from content providers compare to competing ISPs?
* To what extent are my subscribers choosing to get content directly from content providers rather than via my own IPTV offering?
* Are my subscribers getting content from the nearest possible servers?
* What content am I serving from embedded appliances, and how efficiently? Are more needed?
* Is content "shedding" from my embedded appliances to my transit or peers?

#### Enhanced Data in Flow Records

Flow data alone (source and destination ASNs and IPs) lacks the information needed to answer the questions above. True Origin addresses that limitation by enhancing KDE flow records with additional data that enables Kentik to provide richer context. The following factors are now either included in or derived from these enhanced records:

* **Connectivity types**: Content may be delivered from outside the network via IX peering, private peering, or transit. Or it may come internally from embedded cache appliances. Each method has different performance and cost criteria.
* **Source and destination CDNs**: Which CDN delivers what, how does each work, and how does that affect network utilization and cost?
* **Multi-service providers**: A single content provider may operate multiple services distributing multiple forms of content. Facebook, for example, delivers not only a social web application, but also live video, two messaging applications, Instagram, and a gaming application. It’s important to be able to evaluate the resource demands of each such service not only individually but by content provider, largely because interconnection and troubleshooting discussions for any of the underlying OTT services will involve the parent provider.
* **OTT service categorization**: As subscribers become cord-cutters, you need to be able to group OTT services into business categories in order to see the impact of individual types of content (e.g. video) on network utilization and costs regardless of the individual content providers involved.

## How True Origin Works

We've already implemented OTT tagging and CDN tagging as dimensions that you can filter and group-by in Kentik, so to get a feel for how True Origin works we'll start with an example that uses both features. Consider a simple (conceptually) transaction in which a client requests OTT content from a server, which we'll assume in this case is managed by a CDN:

1. The transaction typically begins with a DNS query from a given IP (a subscriber). The DNS server looks up the hostname, which corresponds to the content provider, and returns the IP address of the server that will be the actual source of the content. KDE receives this information via our software host agent.
2. The flow data from the transaction itself contains a source IP (server pushing the content), and a destination IP (a subscriber downloading the content).
3. As DNS responses and flow data are ingested, KDE correlates flows and DNS queries in which:

   1. The source IP in flow data is matched to the IP returned by the DNS query.
   2. The destination IP in flow data is matched to the IP from which the DNS query originated.
   3. The timestamp on both the flow data and the DNS query are approximately the same.
4. When the above conditions are met then the KDE flow record will be enriched with the CDN name (CDN tagging) and the content provider name (OTT tagging).

#### True Origin Challenges

Based on the transaction described above, True Origin may seem relatively simple in theory. In practice, though, successful execution requires resolving the following challenges:

* CDN Tagging by itself isn’t sufficient to obtain a clear picture:

  + CDNs carry traffic for multiple content providers on the same servers.
  + Content providers typically rely on a mix of CDNs to deliver their content.
  + Content providers may deliver their content via both their own infrastructure (including a variety of upstream connectivity methods and providers) and commercial CDNs.
* Correlating DNS queries and network generated flows into a single a time-series database is hard, because you can’t do it offline and send results later. DNS queries/responses must be correlated with flows in real time as they are ingested into the flow-record datastore.
* CDNs come in many different shapes and forms, each involving its own mix of infrastructure and connectivity, which brings up the following considerations that can affect performance and cost:

  + To what extent does the CDN embed servers in the network of the last mile ISP?
  + To what extent does it host content-serving cache servers on its own infrastructure?
  + To what extent does it peer directly with ISPs or leverage diverse transit?

#### True Origin Dimensions

With True Origin, Kentik has solved the challenges above, enabling ISPs — as well as the operators of enterprise and campus LANs — to get a clear picture of how your network utilization and costs are impacted by Internet-sourced content. To do so, you'll use the following source and destination dimensions, whose values are stored in each flow record at ingest:

* **CDN**: A content delivery network, derived as described in CDN tagging above.
* **OTT Service**: An individual content service whose hostname is looked up via DNS.
* **OTT Service Provider**: An entity that offers a service. For example Facebook is the provider for the services Facebook Messenger, WhatsApp, Instagram, etc., Netflix is the provider of the service Netflix, and Google is the provider for Google Drive, GMail, Google Maps, etc.
* **OTT Service Type**: The nature of the content provided by the service. Values include Adult, Ads, Antivirus, Audio, Cloud, Conferencing, Dating, Developer Tools, Documents, Ecommerce, File Sharing, Gaming, IoT, Mail, Maps, Media, Messaging, Network, Newsgroups, Photo Sharing, Social, Software Download, Software Updates, Storage, Video, VPN, Web.

The key factor in OTT tagging is the hostname that, as described earlier, is looked up in a DNS query. That's the source from which we derive the OTT Service value; the other OTT dimension values are associated with the service via our curated list of providers and service types, which continually evolves as we discover and analyze new sources of traffic.

## Classifying Content Traffic

Now that we understand a bit about how True Origin works, let's start thinking about how to put it to use in Kentik. Once again, we'll be approaching this from the point of view of an ISP with eyeball subscribers. Because our subscribers pull far more content from the Internet than they push to it, our ingress traffic has a much greater effect on volume and costs than our egress traffic. So we'll mainly focus on traffic that enters our network towards our subscribers.  
  
With most flow visibility tools it wouldn't be easy to segment our traffic this way, but Kentik's **Interface Classification** feature is designed to handle exactly this kind of challenge. So to get the results we need when using True Origin, we'll start with a quick detour into Interface Classification.

#### Interface Classification Dimensions

Interface classification works in multiple stages. First we apply rules against which interfaces are evaluated and classified, then the flow records from traffic across those interfaces is enriched with information about the interfaces, and then queries on those flow records can group-by or filter based on any or all of the following Interface Classification dimensions:

* **Network Boundary** (source and destination): Internal or External depending on whether the interface connects inside or outside of the network (see Network Boundary Attribute).

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/TO-Network_boundary-375h674w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A54Z&se=2026-05-12T09%3A43%3A54Z&sr=c&sp=r&sig=oxl8wha8RpTIybmOiJyEqqgDSsgqqTtPmT9otc8XVtY%3D)

* **Provider**: The provider via which traffic from a given externally facing interface reaches the Internet (see Provider Classification). Note that in this context “provider” is an umbrella term that could, depending on your specific situation, refer to a transit provider, an Internet Exchange supplier, a customer, or even a CDN (in the case of an embedded cache).
* **Connectivity Type** (source and destination): The type of interface, including types like Transit, Peering, and Backbone (see **Understanding Connectivity Types**).

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/TO-Connectivity_type-219h674w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A54Z&se=2026-05-12T09%3A43%3A54Z&sr=c&sp=r&sig=oxl8wha8RpTIybmOiJyEqqgDSsgqqTtPmT9otc8XVtY%3D)

#### Accounting for Embedded Caches

Among the types of connectivity covered by Interface Classification, **Embedded Cache** is particularly relevant to traffic delivered by some of the major CDNs. As an ISP, I may have interfaces that are connected to a cache appliance server that’s been provided by a CDN or content provider, such as a Facebook appliance (FNA), Google Global Cache (GGC), Netflix Open Connect (OCA), or Akamai cache. If so, I want to be sure that those interfaces are classified as such. To do that, I'll set an Interface Classification rule as shown below.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(186).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A54Z&se=2026-05-12T09%3A43%3A54Z&sr=c&sp=r&sig=oxl8wha8RpTIybmOiJyEqqgDSsgqqTtPmT9otc8XVtY%3D)

Using regex, the above rule will evaluate all SNMP-polled interface descriptions. Matching interfaces will be labeled with these three dimensions:

* Network Boundary = INTERNAL
* Connectivity Type = Embedded Cache
* Provider: The $1 value indicates that the provider value will be determined by performing a group capture, using regex matching, on the first group (AKAMAI|NETFLIX|FACEBOOK|GOOGLE) in the interface description (see Provider Classification with Regex).

Once Interface Classification is applied we can take advantage of it to better understand how traffic from content providers (via CDNs) is affecting our network utilization and costs. In Data Explorer, for example, we can set filters on Interface Classification dimensions so our visualizations and tables are focused on the traffic that's of greatest interest. To see only traffic coming from the outside, we'd set a filter to `SOURCE Network Boundary EQUALS External`; to see only traffic coming out of the network the filter would instead be `DESTINATION Network Boundary EQUALS External`.

## Visualizing Video Traffic

Now that we see how Interface Classification helps tailor our queries we can combine it with True Origin dimensions to really narrow down to a specific type of content. We'll focus on video traffic because these days it represents the lion’s share of traffic volume.  
  
We'll start in the **Filters** pane of the Data Explorer sidebar, where we'll create a new ad hoc filter made up of two filter groups (see **Filter Groups UI**). One group will contain two ORed conditions that cover traffic from two possible sources, either outside the network (via externally facing interfaces) or from caches embedded inside the network. A second filter group, ANDed with the first, will use the OTT Service Type dimension with a value of Video. The resulting filter will look as shown below.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(187).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A54Z&se=2026-05-12T09%3A43%3A54Z&sr=c&sp=r&sig=oxl8wha8RpTIybmOiJyEqqgDSsgqqTtPmT9otc8XVtY%3D)

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(188).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A54Z&se=2026-05-12T09%3A43%3A54Z&sr=c&sp=r&sig=oxl8wha8RpTIybmOiJyEqqgDSsgqqTtPmT9otc8XVtY%3D)

Next, still in the sidebar, we'll choose our group-by dimensions in the Query pane as shown at right (see **Query Sidebar Controls**), so that the query will show us top-X traffic by a combination of OTT dimensions, connectivity type, and CDN. We also specify traffic from all devices (see **Data Sources Pane**) and a time range of Last 1 Day (see **Time Pane**). Finally, from the View Type menu at the top right of the chart display area, we choose Sankey diagram (see **Chart Visualization Types**).  
  
When we run the query with these settings, the resulting Sankey diagram (below) shows the video traffic coming from external sources or an embedded cache, ranked (top-X) according to a key built from our five group-by dimensions.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(189).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A54Z&se=2026-05-12T09%3A43%3A54Z&sr=c&sp=r&sig=oxl8wha8RpTIybmOiJyEqqgDSsgqqTtPmT9otc8XVtY%3D)

## Interpreting True Origin Results

So what do we learn from the above Sankey? Let’s first look at relative traffic volume to confirm that what we see represents reality. Based on the volume of video-only traffic, our top OTT service providers (3rd column) are Netflix, Google's YouTube, and Hulu. Nothing surprising there, but it’s good that our query results are consistent with the usual suspects and our general knowledge about their traffic.  
  
Now let's dig into a few more interesting things. First, Netflix’s OpenConnect CDN is known to have a nearly perfect offload score. Based on these results that appears to be warranted, because we can see that the Netflix-supplied traffic is coming exclusively from the embedded caches in our network. If we ever see Netflix “shedding” traffic into connectivity types like Transit or Peering, we'll know that this is an anomaly that needs to be investigated.  
  
Second, it's worth noting (though it may be assumed) that the vast majority of our video traffic is being supplied via CDNs, either commercial (Akamai, Level3, etc.) or in-house (Netflix, Google CDN, Amazon CloudFront, etc.). This underscores the value to ISPs of being able to see beyond CDNs to identify the true providers of their bandwidth-heavy traffic.  
  
Next, by hovering over any listed OTT service (2nd column) we can easily identify both the connectivity type (4th column) and the CDN (5th column). This is important because, as previously noted, it's extremely important for ISPs to deliver OTT services with the best possible performance at the best possible cost. Content whose connectivity type is transit makes those goals more challenging for two reasons:

* Transit connectivity adds an unknown number of hops, which may impact performance.
* Transit costs are higher than embed or peering (IX or PNI) costs, which impacts our bottom line.

#### One Provider, Many Paths

Despite the potential drawbacks above, the delivery policies of some video content providers leverage the full spectrum of available means. In our Sankey, for example, we can hover over Hulu to see the following delivery methods in use:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(190).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A54Z&se=2026-05-12T09%3A43%3A54Z&sr=c&sp=r&sig=oxl8wha8RpTIybmOiJyEqqgDSsgqqTtPmT9otc8XVtY%3D)The majority is delivered by transit (Level3) which translates into higher cost.
* A good chunk is delivered via PNI (e.g. Edgecast and Fastly).
* A tiny slice of the traffic is delivered via embedded cache (Akamai).

Information like the above enables you to plan intelligently for the demands that CDN-delivered content puts on your infrastructure. And it also alerts you to opportunities to either lower costs or boost performance (e.g. deliver video at a higher bitrate) by trying to find alternatives to transit — more embedded caches, perhaps? — for traffic from content providers.
