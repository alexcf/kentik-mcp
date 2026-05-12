# Security Overview

Source: https://kb.kentik.com/docs/security-overview

---

The article discusses how Kentik ensures the security of its systems and processed data.

> **Note:** For a high level overview of Kentik’s security program, see **https://www.kentik.com/pdfs/KentikSecurityOverview.pdf**.

## Security Summary

When you partner with us, you trust us with metadata and device information related to your private networks. Our systems ensure this information’s safety. Key features include:

* Multiple ways to encrypt data in transit, including legacy data types.
* Encryption of all customer data stored by Kentik (public, private, and on-prem) using AES-256 (compatible with FIPS 140-2).
* Strict access protection to all Kentik systems.

  + Engineers connect over VPNs and use hardware security tokens (Yubikeys) for technical maintenance.
  + Support for single sign-on (SSO) integrations via SAML 2.0 and multi-factor authentication (MFA) enforcement.
  + MFA is limited to time-based one-time password (TOTP) tokens (e.g., Google Authenticator) and hardware-based tokens (e.g., Yubikey) to protect against phishing and takeover attempts.

    > **Note:** For more about SSO integrations, see **About SSO**.
* Kentik maintains annual SOC 2 audits, penetration testing, and continuous membership in a bug bounty program. Audit reports under non-disclosure agreement (NDA).
* Customers are segregated with customer-specific databases.
* Security logs, including authentication events, are in the **Audit Log**.
* Our default service level agreement (SLA) for uptime is 99.99%. Downtime, service degradations, or maintenance are reported on a public-facing status page at **https://status.kentik.com** (US) and **https://status.kentik.eu** (EU).

  > **Note:** Read more about our support and SLA commitments at **https://www.kentik.com/pdfs/KentikSupportSLA.pdf**.

Here’s how the system and data are protected:

|  |  |
| --- | --- |
| **Data in transit** | * All communication to and from Kentik.com is secured with TLS 1.2 or greater. * SNMP supports encrypted and unencrypted modes; we recommend encrypted modes for direct SNMP data. * UDP-based flow logs are unencrypted by design, but Kentik offers methods to gather and encrypt them. * TLS 1.2+ is used for customer data transmitted to Kentik from a host or Google VPC. |
| **Data on disk** | All customer data at rest is encrypted using AES-256, consistent with FIPS 140-2 standards. |
| **Access control** | * Each customer has a separate logical database. * Users can’t modify ingested data or overwrite database tables. * When using local user accounts, we use modern hashing techniques with unique salts to protect passwords, making them harder to crack if exposed. * We support mandatory use of MFA, including via smartphone app (TOTP) or hardware based tokens like Yubikey. |

## Security FAQ

Kentik’s Platform offers scalability, reliability, and security through layered protections. Most customers use our multi-tenant SaaS platform. While this FAQ addresses security concerns in the SaaS environment, many apply to on-prem deployments.

### General Security Questions

* **Does Kentik hold a SOC 2 Type 2, ISO27001, or similar third-party audit report?**  
  Yes. Kentik completes annual SOC 2 (type II) attestations and share audit reports under NDA. We ensure all data center locations with production systems or customer data have maintained strict audit certifications, including SOC 2, ISO 27001, and similar quality certifications.
* **Has Kentik undergone a vulnerability assessment or penetration test of its production environment within the past twelve months?**  
  Yes. Kentik performs vulnerability assessments covering both internal and public-facing systems several times per month. Kentik also engages with an independent third party firm for penetration tests annually, or when significant infrastructure changes occur.
* **How does Kentik monitor the system for security events?**  
  Platform events are logged and aggregated in a dedicated platform.

  + Notable events are aggregated into dashboards, with notifications for anomalous activity sent to internal communications or incident response channels.
  + Certain log types, including access logs, are provided directly to customer admins in-app.
  + Customers can report suspected or confirmed security incidents via email at **security@kentik.com**, or as otherwise specified in your signed master service agreement (MSA) and data processing addendum (DPA).
* **Does Kentik notify customers if their data is involved in a security incident?**  
  Yes, while our signed MSAs and DPAs are authoritative, we notify customers within 48-72 hours (or as specified in relevant data protection law) of confirmed security incidents affecting their data. Find our standard DPA here: **https://www.kentik.com/pdfs/KentikDataProtectionAddendum.pdf**.
* **Does Kentik maintain cyber-liability insurance?**  
  Yes. We maintain cyber liability coverage in addition to other business insurance. We’re happy to share a copy of our certificate of insurance (COI) upon request.
* **How does Kentik ensure that security concerns are addressed throughout the company?**  
  Kentik’s corporate security measures and security culture are addressed holistically in the **Kentik Security Overview PDF**.

  + We ensure technical staff are trained on our SDLC, including the OWASP top 10, Security by Design and Privacy by Design principles, supply chain security, and more.
  + We employ intelligent, driven staff who care about doing the right thing and have the resources to support them.
  + Our Security department includes a CISO, compliance, privacy, and assurance-focused staff, and technical security experts. Each team member maintains strong ties across the organization to ensure security is considered in various business initiatives.

### Compliance and Certification

* **Are the hosting facilities SOC 2 or ISO 27001 certified?**  
  Yes. All of our hosting centers are SOC 2 and/or ISO 27001 certified, and most include additional security-and-privacy-focused certifications. Find certification information from our co-location facilities at **https://www.equinix.com/data-centers/design/standards-compliance** and **https://deft.com/compliance**.
* **Is Kentik’s service PCI and HIPAA compliant?**  
  Kentik doesn’t process or store PCI or ePHI data from customer networks. If you plan to use Kentik services in a PCI or HIPAA environment, discuss your security and compliance needs with our sales or customer engineering staff.
* **Can Kentik work with customers subject to the GDPR?**  
  Yes. Kentik’s privacy program meets GDPR requirements.

  + We minimize personal information processing, provide options to redact IP addresses, use a Data Protection Agreement (DPA) with EU Standard Contractual Clauses (EU SCCs) by default
  + We offer hosting in Frankfurt, Germany.
  + We help customers meet specific GDPR requirements, such as information correction, deletion, or audits.
  + Kentik doesn’t process “sensitive” information under the GDPR. To review a copy of our DPA, please see **https://www.kentik.com/pdfs/KentikDataProtectionAddendum.pdf**.
* **Are customers allowed to conduct vulnerability assessments or penetration testing?**  
  Yes, we’re open to coordinating such activity, but security testing in multi-tenant environments requires careful consideration to minimize risks.

  + Coordinate any penetration testing with your point of contact in advance. Limit penetration testing to non-destructive, rate-limited tests that minimize disruptions and avoid data corruption.
  + Tests are at customer expense and must be completed against up-to-date Kentik platform copies.
  + For more aggressive testing, arrange a dedicated, segregated test environment from production.
  + Passive vulnerability testing can be performed without prior coordination, but limit it to once per month (preferably once per quarter) with reasonable rate limits.
* **How does Kentik manage risks?**  
  Kentik manages risks through annual or more frequent risk assessment exercises involving executives and security.

  + These assessments address underlying risks, evaluate compensating controls, and assess residual risks.
  + They help formally consider risks and identify those requiring treatment, such as transferring risks, pursuing mitigation plans, or avoiding risky activities.
  + Auditors review these formal assessments, which include risks to privacy, security, data integrity, uptime, business continuity, fraud, compliance, and other areas.

### Data Protection and Segregation

* **How does Kentik keep customer data separate in a multi-tenant system?**  
  In our public SaaS, customer data is stored with multiple internal system safeguards to prevent data breaches:

  + Kentik Data Engine (KDE) stores customer network data in separate logical databases and tables.
  + Customer access is strictly limited to their databases and tables.
  + Data access audit logs record all activity within the environment.
  + Kentik’s internal access to production systems and infrastructure is limited to users with compelling business needs.

    - They use dedicated VPNs and hardware-based security tokens to authenticate with the relevant infrastructure.
    - Any such access is subject to internal logging and access review protocols.
* **Is data encrypted at rest?**  
  All customer data across all production systems is encrypted at rest using AES-256, consistent with FIPS 140-2 standards.
* **Is data encrypted in transit?**  
  Customers can enable encryption for all data sent to Kentik, including by using a proxy agent that gathers and encrypts network information.

  + Private network interconnects (PNIs) keep data from public networks.
  + Kentik processes non-TCP protocols that were historically unencrypted.
  + User access to the SaaS web platform is encrypted using TLS 1.2+. To learn more about using our proxy agent, see **About the Proxy Agent**.
* **Can data be removed at customer’s request and per customer’s policies?**  
  Yes. Kentik retains telemetry data according to the customer’s retention period (typically 45 days for full resolution data and 120 days for summarized/derived metrics).

  + Data is automatically deleted upon expiration.
  + If a customer discontinues use of Kentik, we remove their data from our primary systems.
  + Data can also be removed at customer request.
  + Certain data, like device configurations, is retained for the duration of the customer relationship or until deleted by the customer.
* **Will customer data be shared with third party vendors?**

  + Kentik never sells or discloses customer data without permission. For certain relationships, Kentik may seek written consent to use selected or aggregated data for marketing or research purposes, as covered in customer contracts.
  + Kentik uses subprocessors (listed on our **Legal page**) with limited interactions with customer data. The vendor hosting our support ticketing system, for example, processes any information shared by customers in support tickets.
  + Third-party vendors and subprocessors undergo thorough security assessments are required to provide security commitments similar to Kentik’s. Vendor security controls are reassessed during the duration of services, consistent with SOC 2 commitments. Kentik ensures data is returned and/or deleted at the end of a vendor relationship.
* **Are there any interfaces or connections with third party systems?**  
  Kentik offers integrations to third party systems for data and insights management.

  + Customers manage integrations, configuring and authenticating against their accounts.
  + Kentik provides integrations for notification and alerting channels, DDoS mitigation systems, SSO providers, and more (see **Integrations**).

### Access and Identity Management

* **What forms of user management features and controls exist in Kentik?**

  + Access to the SaaS platform requires corporate VPN connections and hardware tokens for authentication. Infrastructure access rights are strictly controlled with regular internal and external, independent audits.
  + Customers can opt for MFA with TOTP/app-based devices or hardware tokens. Kentik supports SSO via SAML 2.0, giving customers control over authentication.

    - For more information on SSO, see **About SSO**
    - For information about configuration of MFA, see **Two-factor Authentication**.
  + Kentik offers tiered user roles (user, admin, super admin) for customer management. RBAC limits user access based on the least-privilege principle (see **Kentik RBAC**).
  + API access tokens are generated at the individual account level and can be reset at any time (see **API Token**).
* **Does Kentik offer single sign-on (SSO)?**  
  Yes. Kentik supports SSO integrations via SAML 2.0 (see **About SSO**).
* **Is multi-factor authentication deployed for "high-risk" environments?**  
  Yes. Kentik uses MFA to protect its corporate environment and prevent unauthorized access to its platform. We also offer MFA options for customer access to the Kentik SaaS.
* **Do you conduct background checks against your employees & contractors (pursuant to local laws)?**  
  Yes. Kentik performs criminal background checks for all employees and contractors, including international workers, as required by local laws. US checks include reviewing federal, state, and local records.
* **Are employees required to sign NDA or Confidentiality Agreements as a condition of employment to protect customer/tenant information?**  
  Yes.
* **What are Kentik’s user termination procedures and timelines?**

  + Kentik deletes customer users immediately from the system when they are removed via the User administration section of the portal.
  + Kentik staff accounts are audited via its access review program and subject to SOC 2 audit. Staff accounts are reviewed promptly upon role change, access is reassigned as necessary, and disabled immediately upon termination (24 hours maximum).
* **What is Kentik’s password policy for systems and infrastructure that is used to host customer data?**  
  Kentik employs secure-key, 2-factor authentication, and logs all access to multiple endpoints.

### Application Security Questions

* **Does Kentik employ a formal secure application methodology?**  
  Kentik maintains a formal SDLC and trains technical staff on security expectations, including reviewing the OWASP Top 10, Security by Design (SbD) and Privacy by Design (PbD) principles, supply chain security, vulnerability and dependency management, internal policy requirements, and other security concepts.
* **What are your processes for security code review and secure development lifecycle?**  
  Kentik’s formal SDLC includes security tools and automation to mitigate risks of insecure code in production.

  + All code undergoes peer review, automated and manual testing, CI/CD integration checks, and a staged migration from test to production environments.
  + Continuous monitoring and enhancement ensure alignment with industry best practices. More details are in our SOC 2 report.
* **Can engineers push code directly to Production?**  
  All code changes undergo peer review and tracking in a dedicated project and code management system with attribution, revert support, and protections to enforce SDLC steps and change control processes. Servers are managed via an Infrastructure as Code (IAC) system for secure, repeatable, and protected deployments.

### Privacy Questions

* **Does Kentik have a Data Protection Officer (DPO)?**  
  Kentik has a designated individual in security and compliance filling the DPO role. Questions to the DPO are sent to **privacy@kentik.com**.
* **Does Kentik have a breach notification policy?**  
  Breach notification is part of incident response. Kentik notifies customers of security incidents consistent with relevant data protection law and signed DPAs, typically within 48-72 hours of confirmation.
* **How long does Kentik store data?**  
  Our standard retention period for full resolution flow data is 45 days, with 120 days retention for derived statistics.

  + Customers can negotiate these retention periods in their signed contract.
  + For NMS users, data retention is 3 months or 13 months, depending on the plan tier.
  + Certain data, like device configurations, is stored as long as you’re an active customer.
  + Kentik automatically purges your data when you move on, or you can purge it at any time.
* **Where are your servers located?**  
  Our U.S. servers are in Ashburn, VA, and we offer EU hosting in Frankfurt, Germany. Specify EU hosting when signing.
* **Can customers control access to features in app?**  
  Yes. We’re moving to a full role-based access control structure (Kentik RBAC) with more granular permissions, helping customers align with the least-privilege principle.

## Security Details

The following tables contain frequently requested information by prospective customers to assess security concerns.

### Data Types

The categories of data that Kentik does and does not collect and store:

|  |  |
| --- | --- |
| **Business Critical Information** | * IP addresses are exported to Kentik either over HTTPS transport or via unencrypted UDP. * UDP flow can be encrypted before leaving your site. * SNMP supports unencrypted and encrypted formats. * Kentik provides tools to redact portions of IP addresses prior to ingestion. |
| **Intellectual Property** | No packet payload data is exported to Kentik. |
| **Other Sensitive Information** | None. Kentik does not process information typically classified as "sensitive" under data protection laws. Processing information is limited to IP address and high level network information. |
| **Payment Card Information** | No packet payload data is exported from customer environments into Kentik. |
| **Personally Identifiable Information (PII)** | No PII is exported (unless contained in IP address). No packet payload data is exported. |
| **Protected Health Information (PHI)** | No PHI is exported. |
| **Public Information** | No public information is exported. |
| **Sensitive Digital Research Data** | No packet payload data is exported to Kentik. |
| **Social Security Numbers (SSN)** | No packet payload data is exported to Kentik. |

### Data Handling

What Kentik does with collected data, and where:

|  |  |
| --- | --- |
| **Overview** | Data is sent via UDP or HTTPS to Kentik’s PaaS offering, and is stored, indexed, and made available via the API and portal. Kentik can be configured to generate alerts related to processed data, such as anomalous traffic spikes or poorly performing devices. Alerts can be sent over email, webhooks, or via customer-configured integrations as outlined in **Integrations**. |
| **Encryption** | * Kentik provides mechanisms to ensure that all data can be encrypted in transit to the Platform. * NetFlow may be sent to a Kentik-provided proxy agent that sits inside the customer’s network and encrypts data bound for Kentik using TLS 1.2+. * All user communication with Kentik’s platform and SaaS website is encrypted. * All on-disk customer data across all production systems is encrypted with AES-256, which meets FIPS 140-2 standards. |
| **Data segregation** | * Each company has an independent database fully partitioned from other customers. * No databases can be directly altered by customers. |
| **Sub-processors** | See **https://kentik.com/legal**for complete sub-processor information. |

### Compliance

Current compliance-related information:

|  |  |
| --- | --- |
| **Compliance reporting** | All customer data is stored in data centers that maintain robust certifications, including SOC 2 and ISO 27001. |
| **Internal/external audits** | Kentik performs annual SOC 2 audits and penetration testing, regular vulnerability assessments (several per month), and maintains continuous membership in a bug-bounty program. |

### Security Monitoring

Information on security-related monitoring:

|  |  |
| --- | --- |
| **Audit logging** | We maintain records of system, server, and application behavior that may be relevant to security investigations or forensic purposes. Logs are stored in a dedicated and centralized system and accessible to administrators. The logging system includes security-relevant alerts that are escalated to internal staff via company communications systems. |
| **Event logging and monitoring** | The following are some of the types of events that are logged and monitored:   * Failed authentication attempts * Unusual data-traffic patterns both internally and externally * Crashing/failed components |
| **Incident Response Plan** | Kentik has a formalized incident response program in place based on NIST 800-61 that includes annual tabletop testing. |

### Security Methodology

Information about the processes used to ensure the security of the application:

|  |  |
| --- | --- |
| **Change management program** | Changes are tested via the originating developer, peer reviewers, Continuous Integration (CI) checks, and evaluation in internal, non-production environments. |
| **Data used for development and/ or testing** | Raw production data is never used in dev or test environments without formal and explicit customer permission. |
| **Testing and approval of changes for Production** | All changes are tested prior to production roll out using a layered approach that includes active and passive checks for unexpected impacts. All changes are approved by an independent reviewer prior to production deployment. |
| **Deployment of changes to Production** | Software updates are deployed to production through an automated process and systems that include testing, security verification, logging, and monitoring. |
| **Separation between testing and Production environments** | Production environments are physically and logically segregated from non-production environments. |
| **OWASP Top 10 Application Security Risks** | Developers are trained on common security weaknesses, including the OWASP Top 10. Vulnerability assessments and bug bounties check for OWASP Top 10 style issues in the production system and infrastructure. |
| **Threat modeling** | Threat modeling or security architecture reviews are performed for new features and major changes. |
| **Automated source-code analysis** | In addition to peer review, we also leverage tooling to detect possible issues related to security, vulnerabilities, and errors in the codebase. |
| **Third-party developers** | The Kentik Platform is Kentik’s primary business asset, with development conducted and managed in-house. Most developers are Kentik employees, with a small number of Kentik-managed contractors in place to assist with specific projects and scopes of work, or to provide specific skill sets that are otherwise difficult to secure. Whenever contractors are in use, Kentik remains accountable for contractor activity and deliverables, and contractors are managed and subject to oversight under Kentik’s overall security program. |
| **SAML/SSO login** | Both 2FA and SAML2/SSO are supported for customers, and SAML2/SSO is utilized by all backend systems. |
| **Web application vulnerability assessment and penetration tests** | We conduct frequent (multiple per month) vulnerability assessments and annual penetration tests. |

### Infrastructure Security

Information related to security on production servers (application and database):

|  |  |
| --- | --- |
| **Infrastructure vulnerability assessment and penetration testing** | We conduct frequent infrastructure vulnerability assessment (several per month) and annual penetration testing. |
| **Patch management** | The Kentik patch management program is based around the severity rating of the flaw, taking into consideration the likelihood of the exploitation of the vulnerability, the distribution of the vulnerability across the production environment, and the impact to operations:   * Critical severity vulnerabilities must be remediated within 7 days of discovery. * High severity vulnerabilities must be remediated within 30 days of discovery. * Medium severity vulnerabilities must be remediated within 60 days of discovery. * Low severity vulnerabilities must be remediated within 90 days of discovery. |
| **User authentication to Production** | Users are authenticated via SSH RSA/DSS key access and 2FA through our SSH gateway servers. |
| **Testing, approval, and logging of system changes** | All changes are managed in a dedicated code and project management system that includes support for tracking and reversion of all changes. All changes undergo peer review and formal sign off by an independent party before migration to Production environments. |

### Disaster Recovery

Basic disaster recovery information:

|  |  |
| --- | --- |
| **Primary data center (Production)** | Located on Kentik's own equipment in Ashburn, VA at Equinix and Deft data centers. |
| **Disaster recovery data centers** | Within our datacenters, Kentik guarantees server-level redundancy and makes a best effort to provide rack- or site-level redundancy where possible. |
| **Recovery time objective (RTO)** | Less than 2 hours |
| **Recovery point objective (RPO)** | Kentik takes daily backups of platform configuration, yielding an RPO of 24 hours. Backups do not include copies of streaming telemetry, such as NetFlow and SNMP data, as this data is extremely high volume (tens of petabytes across Kentik customers) and is regenerated continuously. Note that streaming telemetry data is redundantly stored/replicated within the primary production environment. |
