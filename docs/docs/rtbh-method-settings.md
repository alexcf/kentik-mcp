# RTBH Method Settings

Source: https://kb.kentik.com/docs/rtbh-method-settings

---

This article covers the specific configuration fields and settings required when building an RTBH (Remotely Triggered Black Hole) mitigation method from Kentik’s **Manage Mitigations**page.

## RTBH Method Details

The **Details** tab of the settings dialog for an RTBH mitigation includes the following fields:

* ![Configuration details for Carpet Bomb RTBH method including community and next hop information.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/MM-edit-rtbh-method-details-tab.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A33Z&se=2026-05-12T09%3A31%3A33Z&sr=c&sp=r&sig=FBPdF3XuygHUNbfK0ghbtHJP8Y%2BxOlOEUZKqIz8wT9M%3D)**Commonly Used Communities** (table): Provided for reference only, a list of communities commonly used in RTBH (you are not restricted to using one of these communities).
* **Community to Advertise**: The community that has been programmed on the customer’s router to induce a black hole next hop for the IPv4 address attached to the community.
* **Next Hop**: Next-hop IP addresses (IPv4 and Ipv6). In some environments these will be the destination IP to blackhole. These numbers have traditionally been selected from the `192.0.2.0/24` CIDR block, but may be any IP address.
* **Local Preference**: Set the priority for the RTBH route. A high setting helps ensure that when there is more than one route the RTBH route will be preferred by the BGP best path selection process.
* **Ensure at least/24**: A switch that tells Kentik to convert the provided next-hop IP address to CIDR notation. Use if you plan to withdraw blocks from certain routers and re-advertise in other locations (otherwise, leave unchecked).

> **Notes:**
>
> * For more on RTBH, see**Configuring RTBH Mitigation**, **Mitigation Platforms**, and **RTBH Mitigation**.
> * For general method settings like notification channels and automation grace periods, return to the **Mitigation Methods** article.
