# Using Ultimate Exit

Source: https://kb.kentik.com/docs/using-ultimate-exit

---

Kentik includes **Ultimate Exit Dimensions** that enable end-to-end visibility of your traffic for business intelligence analytics.

## Why Ultimate Exit

End-to-end visibility of your traffic means that you need an easy way to figure out what volumes of traffic are flowing in and out of your network, from any source to any destination network. You can then use that information to cut costs (e.g., peering) and to more accurately estimate the cost of carrying any set of traffic for any given customer.

A great use case for this is assessing potential peer or transit prospects. How many times have you had to toggle between multiple spreadsheets that contain only approximations of traffic to or from various ASNs? Instead of being able to accurately estimate ROI to make what should be a simple decision, you quickly get bogged down in hacked, convoluted Excel formulas.

What about trying to figure out how much traffic from a peer is being routed locally versus over more costly long-haul links? You need to able to figure out precisely at the site and device level — and at the interface level in the future — the traffic flowing between network entry and exit points. ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(200).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A55Z&se=2026-05-12T09%3A29%3A55Z&sr=c&sp=r&sig=79YyGKc9SRdrSxOFkUlcnhXpi2NL%2FiTM4Q6%2F3lUnsR4%3D)It turns out that the flow consolidation and analytics required for this task are sophisticated, and beyond the scope of the home-grown tools, data infrastructure, and software engineering capacities available to most network engineering teams. That's not surprising, because it’s a hard problem. So to make it less painful to get accurate information on utilization of your network, a large set of **Ultimate Exit Dimensions** are available in Kentik:

* Ultimate Exit Interface

  > **Note:** This single group-by dimension is represented in filtering by three Ultimate Exit Interface dimensions: ID, Name, and Description.
* Ultimate Exit Connectivity Type
* Ultimate Exit Network Boundary
* Ultimate Exit Provider
* Ultimate Exit Site
* Ultimate Exit Device
* Ultimate Exit Site Country

## Using the Dimensions

How do you use these **Ultimate Exit Dimensions**? Let’s take one example, using the Ultimate Exit Site dimension. If you're a transit provider, you move packets from content providers to eyeball ISPs, and carry them over a costly global backbone. You want to look at the traffic you’re exchanging with one of the major content providers, such as Google, and see where it comes into and out of your network.![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(201).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A55Z&se=2026-05-12T09%3A29%3A55Z&sr=c&sp=r&sig=79YyGKc9SRdrSxOFkUlcnhXpi2NL%2FiTM4Q6%2F3lUnsR4%3D)  
  
Let’s further assume that you run a well-organized network, so your interconnections with Google are indicated as part of your interface description nomenclature. This means that, as shown at right, you can easily use a simple filter to identify the interfaces involved in these interconnects.  
  
As a side note, if you know that you’re going to be looking at these same interfaces often, you can save the filter (see **Saved Filters**) so that you can apply it any time you need it in any Data Explorer query that you’re working on.  
  
Continuing with the Google example, here, in sequence, is what you're going to want to look at:![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(202).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A55Z&se=2026-05-12T09%3A29%3A55Z&sr=c&sp=r&sig=79YyGKc9SRdrSxOFkUlcnhXpi2NL%2FiTM4Q6%2F3lUnsR4%3D)

* The site where the traffic enters the network.
* The site where the traffic leaves the network.
* The next-hop Network.
* Which eyeball network it is terminating at, i.e. its Destination AS.

Using Kentik’s Ultimate Exit dimensions, you can now answer this question with a query that's set up in the sidebar as shown in the image at right. For the most useful visualization, set the display type in the Display pane (not shown) to Sankey Flow diagram.

Looking at the generated Sankey diagram below you can now instantly see what traffic is flowing between the entry site and the ultimate exit site, and which eyeball networks are reached. At this point you would typically look at where transport is the most expensive or least performant between your entry site and ultimate exit site and take appropriate steps to optimize based on that information.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(203).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A55Z&se=2026-05-12T09%3A29%3A55Z&sr=c&sp=r&sig=79YyGKc9SRdrSxOFkUlcnhXpi2NL%2FiTM4Q6%2F3lUnsR4%3D)

*A simple example of what you can do with Kentik's Ultimate Exit feature.*
