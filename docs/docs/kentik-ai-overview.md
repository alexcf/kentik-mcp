# Kentik AI Overview

Source: https://kb.kentik.com/docs/kentik-ai-overview

---

This article covers Kentik AI, including data privacy and security considerations.

## What is Kentik AI?

Kentik AI is an umbrella term for a wide range of different AI capabilities in Kentik products that deliver an improved, enriched, and streamlined user experience in network observability and operations use cases.

Kentik AI features include:

* **AI Advisor**: AI Advisor is our next evolution of AI services which provides an AI agent designed to perform troubleshooting of network problems using a conversational user interface. Users can ask high-level questions while the AI Advisor creates a plan, collects and analyzes data, and suggests conclusions and next steps. The AI Advisor is equipped with internally-defined tooling used to support tasks and investigations requested by the end user. This tooling includes the ability to check relevant network telemetry and metadata from Kentik Platform, use multi-step processes to work with data, and integrate responses with common network tooling (such as PeeringDB, RIPE, etc.).
* **Metrics Explorer Query Assistant**: Metrics Explorer Query Assistant (or just Query Assistant) translates natural language questions into a Kentik Metrics Explorer query for faster and more natural interactions with NMS metrics data.
* **Cause Analysis**: Cause Analysis – available on-demand or automatically triggered – finds the most likely contributing factors for sudden traffic changes, spikes, or drops.
* **AI-enhanced Insights**: Kentik Ai-enhanced Insights provides users with an AI-driven analysis to correlate data from multiple data points and sources.
* **Summary**: Currently integrated with Cloud Pathfinder, the Summary provides a simple AI-generated breakdown of complex network issues with handy solutions and tips.

## How does Kentik use AI?

### GenAI Models

Kentik AI uses Generative AI Large Language Model (LLM) services from companies that are leaders in this technology, including:

* OpenAI models hosted by OpenAI
* Foundational models from Anthropic, Google, or others hosted within Google Cloud (GCP Vertex AI service)
* Anthropic models hosted in AWS, part of Amazon Bedrock service
* Anthropic models hosted by Anthropic

### Security & Privacy Commitments

The following commitments apply to all of Kentik’s AI services:

* Kentik ensures AI partners are not allowed to retain system data for their own purposes (excluding short holds for fraud and abuse monitoring).
* Kentik ensures customer queries and data are not incorporated into AI-training data sets. Kentik uses a prompt-engineering based approach to deliver AI services, which avoids the need to retrain or fine tune genAI services with customer data.
* Kentik designs the system to limit prompt injection risks by ensuring outputs remain within expected formats, functions are bounded within internally defined limits, and AI-accessible data types are limited to a restricted set managed by Kentik’s AI engineers.
* Kentik executes AI features within a specific user or customer context, which helps ensure mistakes or prompt injection attacks can’t be used to compromise data outside of the intended environment.
* For certain AI features, Kentik uses AI systems to translate natural language questions into queries Kentik can execute against internal systems to limit the risk of hallucinations or inaccurate outputs.
* All AI queries are read-only and can’t be used to update or alter system data.
* Kentik uses reputable AI service providers from well-known generative AI companies that perform their own security, privacy, anti-abuse, and safety analyses prior to release, and conduct their own security-focused auditing. To read more about our provider’s security practices and privacy standards, please see:

  + **OpenAI’s Enterprise privacy policy**
  + **Google Cloud Vertex AI, Generative AI and data governance**
  + **Amazon Bedrock, User Guide, Data Protection**
  + **Anthropic Trust Center**

> Notes:
>
> * AI features can be used to help non-English speakers interface more naturally with certain elements of Kentik systems.
> * Kentik AI features are used to help reveal insights into complex network management issues, and can’t be used to make automated decisions affecting end users, process sensitive information types relating to end users (such as political affiliation or union membership), or result in discriminatory effects against your customers.

#### User Control

To keep customers in control when it comes to AI services, AI features are OFF by default and must be enabled at the company level by a super admin user.

* Enabling Kentik AI will enable the use of AI Advisor, Cause Analysis, Query Assistant, and other Kentik AI-assisted features for your company as described in this document.
* Disabling Kentik AI will disable the use of generative AI-related features for your company.

### AI Advisor

AI Advisor is our newest and most powerful AI offering, providing an AI agent equipped with Kentik-managed internal tooling. AI Advisor is designed to understand user’s questions/queries, plan out a sequence of steps to arrive at a result, access relevant system data surfaced to the AI system (described below), analyze results, and provide feedback and summaries to the user.

To accomplish this, AI Advisor analyzes the following user account-level data:

* **Telemetry**: Time-series and aggregated data of traffic flows, metrics, SNMP traps, Syslogs, BGP routes, DNS records, OTT/application services, from on-prem network devices, and the Cloud
* **Inventory**: Cloud and on-prem inventory, Devices, Interfaces, Sites, Kentik Agents and related metadata, and device configuration (if collected)
* **Kentik Configuration**: Interface Classification, Custom Dimensions, Custom Applications, Flow Tags, AS Groups, Custom Geo Locations, Site Markets, Labels, Saved Filters, Network Classification and other related information
* **Events & Alerts**: Alerts, Insights, and Cause Analysis results
* **Test Results**: Synthetics test results and Cloud Pathfinder queries
* **Workflows**: Capacity Planning, Connectivity Costs, DDoS, Dashboards, Saved Views, and other Kentik features
* **Data Enrichment Sources**: AI responses may include contextual data from common networking tools like RIPEStat, BGPview, PeeringDB, or customer-configured sources to help interpret and analyze network features.

### Query Assistant

For the services described above, generative AI is used to translate a user’s questions/queries into a structured format that is used to query telemetry data from Kentik Data Engine.

Data submitted to the GenAI service includes:

* **User’s question**: Any question typed in the NMS Query Assistant (QA) input prompt.
* **Query context**: Names of metrics, dimensions, and measurements used in the query schema for Kentik Data Explorer and NMS Metrics Explorer.
* **Network metadata**: A small set of relevant metadata about a user’s network, currently contains:

  + Site names
  + Site countries
  + Site markets
  + Provider names used in interface classification feature
  + Device labels
  + AS Groups
  + DDoS alerting policy names
* **Results**: Aggregated results of previous queries, including tables. This allows for extended troubleshooting sessions to reference previously-returned data.

> **Note**: Network metadata types may be expanded over time to ensure support for commonly-requested user queries. Kentik will document any changes made.

### AI-enhanced Insights, Cause Analysis, and Summary

For Ai-enhanced Insights, Cause Analysis, and Summary, the GenAI service is used to interpret query or analysis results and describe them in a more user-friendly format.

* Kentik performs an analysis of the collected network telemetry data using in-house algorithms, data mining, or machine learning techniques.
* Data typically relates to a specific problem, such as an anomaly in network traffic, out-of-threshold metric values, or connectivity paths in the public cloud infrastructure.
* Analyses yield aggregated data structures that are used as inputs to GenAI services that provide user-friendly explanations of the data.

## Resources

If you still have questions about AI features, please reach out to your dedicated Account Executive, Customer Success Manager, or Kentik’s Security team at **security@kentik.com****.** To learn more about Kentik’s overall security and privacy posture, see our **Trust Center** or **Security Overview**.

## Legal Disclaimer

Kentik's AI features involve sending user queries and associated contextual data as described here to LLM providers, including OpenAI and Google. We make sure these providers are NOT allowed to use this data to train or refine their AI systems; however, we encourage all customers to make sure any use of AI is consistent with their corporate policies prior to use. Please be aware that AI may contain mistakes and that the use of these features is provided on an as-is basis without any warranties. By using Kentik AI, you consent to the data processing described above.
