# Protect

Source: https://kb.kentik.com/docs/protect-overview

---

The article covers the **Protect** section of the Kentik portal.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(253).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A56Z&se=2026-05-12T09%3A27%3A56Z&sr=c&sp=r&sig=28EJsx5D%2FlPu%2FTZdMfZUEH8jjuUeSP8DhJ21SdlVc3Y%3D)

*Protect modules like Alerting to see and respond to network attacks and availability/performance threats*

## About Protect

The Kentik portal’s **Protect** section provides modules/workflows that inform you about network availability and security threats, enabling you to defend your network.

## Protect Workflows

The **Protect** section of the portal includes the following modules/workflows:

* **DDoS Defense**: Automates the entire DDoS attack lifecycle, including detection, investigation, and mitigation  (see **DDoS Defense**).

  + Machine learning-based traffic profiling identifies attacks faster and more accurately, eliminating false positives/negatives.
  + Visualizations show attack characteristics and network impact, and built-in triggers respond automatically with mitigations like RTBH, Flowspec, and external hardware/services.
* **Alerting**: Analyzes your network traffic and detects anomalous patterns that threaten availability or performance (see **Alerting**).

  + Alert policies define conditions that trigger "alarm states" and generate alerts.

    Current and historical alerts are listed on the Protect » **Alerting** page, showing the time, severity, status, and triggered conditions.
  + Policies can generate notifications and initiate automatic mitigations in response to alerts .
* **Mitigations**: Mitigations are protective actions that prevent network interruptions due to undesirable traffic like DDoS attacks (see **Mitigations**).

  + Can be triggered automatically or manually.
  + Kentik offers built-in options like Flowspec and RTBH, and integrates with third-party providers like Cloudflare, Radware, and A10.
* **Threat/Botnet Analysis**: Identifies threats such as botnet command and control (CC) servers, malware distribution points, phishing websites, and spam sources (see **Threats & Botnets**).

  + Kentik enriches flow records in KDE with IP reputation data from Spamhaus.
  + Accessible via the **Botnet & Threat-feed Analysis** dashboard.
* **RPKI Analysis**: Identifies would-be invalid and/or dropped traffic sites if strict RPKI validation is enforced on the routers (see **RPKI Analysis**).

  + Kentik's RPKI implementation (see **About RPKI**), based on Cloudflare’s GoRTR, validates BGP routes and assigns RPKI values to KDE flow records.
  + Accessible via the **RPKI Analysis** dashboard.
