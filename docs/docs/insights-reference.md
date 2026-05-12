# Insights Reference

Source: https://kb.kentik.com/docs/insights-reference

---

This article discusses Kentik's Insights engine, which analyzes traffic to extract useful information that can be presented to users.

> **Note:** For more information about insights and the Insights module of the Kentik portal, see **Insights**.

## Edge

For comparison insights in the Edge family, we evaluate traffic volume through your infrastructure and notify you of changes related to top-X rankings. The primary dimensions we evaluate are ASN and sites. The conditions evaluated for a given dimension include:

* Changes in individual or overall traffic volume.
* Changes in the composition of the top 10 set.
* Changes in top 10 ranking.

> **Notes:** For the insights in this family, the metrics displayed in the **Statistics Pane** (see **Insight Details Page**) include bitrate over the most recent 7-day period and percent change from the prior period.

This family includes the following insights:

| **Insight** | Summary |
| --- | --- |
| **Peering Egress Traffic Today** | Significant change in volume of peering traffic. |
| **Through Destination ASN Comparison** | Week-over-week changes to the volume and/or ranking of the top 10 destination ASNs of the traffic through the edge of your infrastructure. |
| **Through Source ASN Comparison** | Week-over-week changes to the volume and/or ranking of the top 10 source ASNs of the traffic through the edge of your infrastructure. |
| **Through Site Comparison** | Week-over-week changes to the through traffic in your site. |

#### Peering Egress Traffic Today

Identifies significant bit rate changes to your peering traffic.

* **Evaluation**: Most recent 24-hour period vs. prior 24-hour period.
* **Dimensions**: None
* **Filters**: Connectivity Type = IX, Free PNI, Paid PNI
* **Metric**: Bits per second
* **Data sources**: All traffic

#### Through Destination ASN Comparison

Identifies significant week-over-week changes to the top 10 Destination ASNs in the traffic through the edge of your network. The conditions evaluated for major changes include traffic volume for one or more ASNs, the makeup of the top 10 ASN set, and changes in top 10 ranking. Metrics are provided for bitrate over the most recent 7-day period and percent change from the prior period.

* **Evaluation**: Most recent 7-day period vs. prior 7-day period.
* **Dimensions**: Destination ASN
* **Filters**: Simple Traffic Profile = Through
* **Metric**: Bits per second
* **Data sources**: All traffic

#### Through Source ASN Comparison

Identifies significant week-over-week changes to the top 10 Source ASNs in the traffic through the edge of your network. The conditions evaluated for major changes include traffic volume for one or more ASNs, the makeup of the top 10 ASN set, and changes in top 10 ranking. Metrics are provided for bitrate over the most recent 7-day period and percent change from the prior period.

* **Evaluation**: Most recent 7-day period vs. prior 7-day period.
* **Dimensions**: Source ASN
* **Filters**: Simple Traffic Profile = Through
* **Metric**: Bits per second
* **Data sources**: All traffic

#### Through Site Comparison

Identifies significant week-over-week changes to traffic through a specific site.

* **Evaluation**: Most recent 7-day period vs. prior 7-day period.
* **Dimensions**: Site
* **Filters**: Simple Traffic Profile = Through
* **Metric**: Bits per second
* **Data sources**: All traffic

## Geolocation

For insights in the Geolocation family, we evaluate traffic volume to or from your infrastructure (on-prem and/or cloud) and notify you of changes related to top-X rankings. The dimensions we evaluate include city, region, and country (e.g., a given country moving into or out of the top-10 countries). The conditions evaluated for a given dimension include:

* Changes in individual or overall traffic volume.
* Changes in the composition of the top 10 set.
* Changes in top 10 ranking.

> **Note:** For the insights in this family, the metrics displayed in the **Statistics Pane** (see **Insight Details Page**) include bitrate over the most recent 7-day period and percent change from the prior period.

This family includes the following insights:

| **Insight** | Summary |
| --- | --- |
| **Inbound City Comparison** | Week-over-week changes to the volume and/or ranking of the top 10 cities in the traffic inbound to your infrastructure. |
| **Outbound City Comparison** | Week-over-week changes to the volume and/or ranking of the top 10 cities in the traffic outbound from your infrastructure. |
| **Inbound Country Comparison** | Week-over-week changes to the volume and/or ranking of the top 10 countries in the traffic inbound to your infrastructure. |
| **Outbound Country Comparison** | Week-over-week changes to the volume and/or ranking of the top 10 countries in the traffic outbound from your infrastructure. |
| **Inbound Region Comparison** | Week-over-week changes to the volume and/or ranking of the top 10 regions in the traffic inbound to your infrastructure. |
| **Outbound Region Comparison** | Week-over-week changes to the volume and/or ranking of the top 10 regions in the traffic outbound from your infrastructure. |

#### Inbound City Comparison

Identifies significant week-over-week changes to the top 10 cities in the traffic inbound to your infrastructure (on-prem and/or cloud). The conditions evaluated for major changes include traffic volume for one or more cities, the makeup of the top 10 cities set, and changes in top 10 ranking.

* **Evaluation**: Most recent 7-day period vs. prior 7-day period.
* **Dimensions**: Source City
* **Filters**: Simple Traffic Profile = Inbound
* **Metric**: Bits per second
* **Data sources**: All traffic

#### Outbound City Comparison

Identifies significant week-over-week changes to the top 10 cities in the traffic outbound from your infrastructure (on-prem and/or cloud). The conditions evaluated for major changes include traffic volume for one or more cities, the makeup of the top 10 cities set, and changes in top 10 ranking.

* **Evaluation**: Most recent 7-day period vs. prior 7-day period.
* **Dimensions**: Destination City
* **Filters**: Simple Traffic Profile = Outbound
* **Metric**: Bits per second
* **Data sources**: All traffic

#### Inbound Country Comparison

Identifies significant week-over-week changes to the top 10 countries in the traffic inbound to your infrastructure (on-prem and/or cloud). The conditions evaluated for major changes include traffic volume for one or more countries, the makeup of the top 10 countries set, and changes in top 10 ranking.

* **Evaluation**: Most recent 7-day period vs. prior 7-day period.
* **Dimensions**: Source Country
* **Filters**: Simple Traffic Profile = Inbound
* **Metric**: Bits per second
* **Data sources**: All traffic

#### Outbound Country Comparison

Identifies significant week-over-week changes to the top 10 countries in the traffic outbound from your infrastructure (on-prem and/or cloud). The conditions evaluated for major changes include traffic volume for one or more countries, the makeup of the top 10 countries set, and changes in top 10 ranking.

* **Evaluation**: Most recent 7-day period vs. prior 7-day period.
* **Dimensions**: Destination Country
* **Filters**: Simple Traffic Profile = Outbound
* **Metric**: Bits per second
* **Data sources**: All traffic

#### Inbound Region Comparison

Identifies significant week-over-week changes to the top 10 regions in the traffic inbound to your infrastructure (on-prem and/or cloud). The conditions evaluated for major changes include traffic volume for one or more regions, the makeup of the top 10 regions set, and changes in top 10 ranking.

* **Evaluation**: Most recent 7-day period vs. prior 7-day period.
* **Dimensions**: Source Region
* **Filters**: Simple Traffic Profile = Inbound
* **Metric**: Bits per second
* **Data sources**: All traffic

#### Outbound Region Comparison

Identifies significant week-over-week changes to the top 10 regions in the traffic outbound from your infrastructure (on-prem and/or cloud). The conditions evaluated for major changes include traffic volume for one or more regions, the makeup of the top 10 regions set, and changes in top 10 ranking.

* **Evaluation**: Most recent 7-day period vs. prior 7-day period.
* **Dimensions**: Destination Region
* **Filters**: Simple Traffic Profile = Outbound
* **Metric**: Bits per second
* **Data sources**: All traffic

## Hybrid Cloud

For insights in the Hybrid Cloud family, we evaluate traffic volume to, from, or in your cloud infrastructure and notify you of changes related to top-X rankings. The dimensions we evaluate include applications, cloud services, and regions (e.g., a given region moving into or out of the top-10 regions). The conditions evaluated for a given dimension include:

* Changes in individual or overall traffic volume.
* Changes in the composition of the top 10 set.
* Changes in top 10 ranking.

> **Note:** For the insights in this family, the metrics displayed in the **Statistics Pane** (see **Insight Details Page**) include bitrate over the most recent 7-day period and percent change from the prior period.

This family includes the following insights:

| **Insight** | Summary |
| --- | --- |
| **Outbound Cloud Service Comparison** | Week-over-week changes to the volume and/or ranking of the top 10 cloud service traffic outbound from your on-prem infrastructure. |
| **Inbound Cloud Service Comparison** | Week-over-week changes to the volume and/or ranking of the top 10 cloud service traffic inbound to your on-prem infrastructure. |
| **Multi-Cloud Comparison** | Week-over-week changes to the volume and/or ranking of the top 10 cloud provider traffic to and from your cloud infrastructure. |
| **AWS Region to On-Prem Comparison** | Week-over-week changes to the volume and/or ranking of the top 10 AWS Region traffic inbound to your on-prem infrastructure. |
| **Inbound AWS Region Comparison** | Week-over-week changes to the volume and/or ranking of the top 10 AWS Region traffic inbound to your AWS infrastructure. |
| **On-Prem to AWS Region Comparison** | Week-over-week changes to the volume and/or ranking of the top 10 AWS Region traffic outbound to your on-prem infrastructure. |
| **Outbound AWS Region Comparison** | Week-over-week changes to the volume and/or ranking of the top 10 AWS Region traffic outbound to your AWS infrastructure. |
| **Outbound AWS Region and Subnet Comparison** | Week-over-week changes to the volume and/or ranking of the top 10 AWS Region and AWS Subnet traffic outbound to your AWS infrastructure. |
| **Internal AWS Region to Region Comparison** | Week-over-week changes to the volume and/or ranking of the top 10 AWS Region Pairs traffic outbound to your AWS infrastructure. |
| **Hybrid Cloud Application Comparison** | Week-over-week changes to the volume and/or ranking of the top 10 application traffic to and from your on-prem infrastructure. |
| **AWS Gateway to Inside Comparison** | Week-over-week changes to the volume and/or ranking of the top 10 AWS Gateway traffic to your on-prem infrastructure. |
| **AWS Gateway to Outside Comparison** | Week-over-week changes to the volume and/or ranking of the top 10 AWS Gateway traffic to outside from your infrastructure. |
| **AWS Internet Gateway for Cloud Internal Traffic** | Identifies any use of an Internet Gateway for Cloud Internal traffic. |

#### Outbound Cloud Service Comparison

Identifies significant week-over-week changes to the top 10 cloud services in the traffic outbound from your infrastructure (on-prem and/or cloud). The conditions evaluated for major changes include traffic volume for one or more cloud services, the makeup of the top 10 cloud service set, and changes in top 10 ranking. Metrics are provided for bitrate over the most recent 7-day period and percent change from the prior period.

* **Evaluation**: Most recent 7-day period vs. prior 7-day period.
* **Dimensions**: Destination Cloud Service
* **Filters**: Simple Traffic Profile = Outbound
* **Metric**: Bits per second
* **Data sources**: All traffic

#### Inbound Cloud Service Comparison

Identifies significant week-over-week changes to the top 10 cloud services in the traffic inbound to your infrastructure (on-prem and/or cloud). The conditions evaluated for major changes include traffic volume for one or more cloud services, the makeup of the top 10 cloud service set, and changes in top 10 ranking. Metrics are provided for bitrate over the most recent 7-day period and percent change from the prior period.

* **Evaluation**: Most recent 7-day period vs. prior 7-day period.
* **Dimensions**: Inbound Cloud Service
* **Filters**: Simple Traffic Profile = Inbound
* **Metric**: Bits per second
* **Data sources**: All traffic

#### Multi-Cloud Comparison

Identifies significant week-over-week changes to the top cloud service traffic. The conditions evaluated for major changes include traffic volume for one or more cloud services, the makeup of the top cloud service set, and changes in rankings. Metrics are provided for bitrate over the most recent 7-day period and percent change from the prior period.

* **Evaluation**: Most recent 7-day period vs. prior 7-day period.
* **Dimensions**: Traffic Origination and Traffic Termination
* **Filters**: Traffic Profile = [Inside to Cloud, Cloud to Inside, Multi Cloud]
* **Metric**: Bits per second
* **Data sources**: All traffic

#### AWS Region to On-Prem Comparison

Identifies significant week-over-week changes to the top AWS Regions in the traffic inbound to your on-prem infrastructure. The conditions evaluated for major changes include traffic volume for one or more AWS Regions, the makeup of the top AWS Region set, and changes in AWS Region ranking. Metrics are provided for bitrate over the most recent 7-day period and percent change from the prior period.

* **Evaluation**: Most recent 7-day period vs. prior 7-day period.
* **Dimensions**: Source AWS Region
* **Filters**: Cloud Provider = AWS. Traffic Profile = Cloud to Inside
* **Metric**: Bits per second
* **Data sources**: All traffic

#### Inbound AWS Region Comparison

Identifies significant week-over-week changes to the top AWS Regions in the traffic inbound to your AWS infrastructure. The conditions evaluated for major changes include traffic volume for one or more AWS Regions, the makeup of the top AWS Region set, and changes in AWS Region ranking. Metrics are provided for bitrate over the most recent 7-day period and percent change from the prior period.

* **Evaluation**: Most recent 7-day period vs. prior 7-day period.
* **Dimensions**: Destination AWS Region
* **Filters**: Simple Traffic Profile = Inbound. Cloud Provider = AWS.
* **Metric**: Bits per second
* **Data sources**: All traffic

#### On-Prem to AWS Region Comparison

Identifies significant week-over-week changes to the top AWS Regions in the traffic outbound from your on-prem infrastructure. The conditions evaluated for major changes include traffic volume for one or more AWS Regions, the makeup of the top AWS Region set, and changes in AWS Region ranking. Metrics are provided for bitrate over the most recent 7-day period and percent change from the prior period.

* **Evaluation**: Most recent 7-day period vs. prior 7-day period.
* **Dimensions**: Destination AWS Region
* **Filters**: Cloud Provider = AWS. Traffic Profile = Inside to Cloud
* **Metric**: Bits per second
* **Data sources**: All traffic

#### Outbound AWS Region Comparison

Identifies significant week-over-week changes to the top AWS Regions in the traffic outbound from your AWS infrastructure. The conditions evaluated for major changes include traffic volume for one or more AWS Regions, the makeup of the top AWS Region set, and changes in AWS Region ranking. Metrics are provided for bitrate over the most recent 7-day period and percent change from the prior period.

* **Evaluation**: Most recent 7-day period vs. prior 7-day period.
* **Dimensions**: Source AWS Region
* **Filters**: Simple Traffic Profile = Outbound. Cloud Provider = AWS.
* **Metric**: Bits per second
* **Data sources**: All traffic

#### Outbound AWS Region and Subnet Comparison

Identifies significant week-over-week changes to the top AWS Regions and AWS Subnets in the traffic outbound from your AWS infrastructure. The conditions evaluated for major changes include traffic volume for one or more AWS Regions and AWS Subnets, the makeup of the top AWS Regions and AWS Subnets set, and changes in AWS Regions and AWS Subnets ranking. Metrics are provided for bitrate over the most recent 7-day period and percent change from the prior period.

* **Evaluation**: Most recent 7-day period vs. prior 7-day period.
* **Dimensions**: Source AWS Region, Source AWS Subnet
* **Filters**: Simple Traffic Profile = Outbound. Cloud Provider = AWS.
* **Metric**: Bits per second
* **Data sources**: All traffic

#### Internal AWS Region to Region Comparison

Identifies significant week-over-week changes to the top AWS Regions pairs in the traffic outbound from your AWS infrastructure. The conditions evaluated for major changes include traffic volume for one or more AWS Regions pairs, the makeup of the top AWS Regions pairs set, and changes in AWS Regions pairs ranking. Metrics are provided for bitrate over the most recent 7-day period and percent change from the prior period.

* **Evaluation**: Most recent 7-day period vs. prior 7-day period.
* **Dimensions**: Source AWS Region, Destination AWS Region
* **Filters**: Simple Traffic Profile = Internal. Cloud Provider = AWS.
* **Metric**: Bits per second
* **Data sources**: All traffic

#### Hybrid Cloud Application Comparison

Identifies significant week-over-week changes to the top 10 applications in the traffic to and from your on-prem and cloud infrastructure. The conditions evaluated for major changes include traffic volume for one or more applications, the makeup of the top 10 application set, and changes in top 10 ranking. Metrics are provided for bitrate over the most recent 7-day period and percent change from the prior period.

* **Evaluation**: Most recent 7-day period vs. prior 7-day period.
* **Dimensions**: Application
* **Filters**: Traffic Profile = [Inside to Cloud, Cloud to Inside]
* **Metric**: Bits per second
* **Data sources**: All traffic

#### AWS Gateway to Inside Comparison

Identifies significant week-over-week changes to the top AWS Gateways in the traffic inbound to your on-prem (inside) infrastructure. The conditions evaluated for major changes include traffic volume for one or more AWS Gateways, the makeup of the top AWS Gateway set, and changes in AWS Gateway ranking. Metrics are provided for bitrate over the most recent 7-day period and percent change from the prior period.

* **Evaluation**: Most recent 7-day period vs. prior 7-day period.
* **Dimensions**: Gateway Type
* **Filters**: Traffic Profile = Cloud to Inside. Cloud Provider = AWS
* **Metric**: Bits per second
* **Data sources**: All traffic

#### AWS Gateway to Outside Comparison

Identifies significant week-over-week changes to the top AWS Gateways in the traffic inbound to your on-prem infrastructure. The conditions evaluated for major changes include traffic volume for one or more AWS Gateways, the makeup of the top AWS Gateway set, and changes in AWS Gateway ranking. Metrics are provided for bitrate over the most recent 7-day period and percent change from the prior period.

* **Evaluation**: Most recent 7-day period vs. prior 7-day period.
* **Dimensions**: Gateway Type
* **Filters**: Traffic Profile = Cloud to Outside. Cloud Provider = AWS
* **Metric**: Bits per second
* **Data sources**: All traffic

#### AWS Internet Gateway for Cloud Internal Traffic

Identifies any use of an Internet Gateway for Cloud Internal traffic.

* **Evaluation**: Most recent 7-day period vs. prior 7-day period.
* **Dimensions**: Gateway Type
* **Filters**: Traffic Profile = Cloud Internal. Cloud Provider = AWS
* **Metric**: Bits per second
* **Data sources**: All traffic

## Network Health

Insights that call attention to possible problems on your network such as unexpectedly low or high traffic bitrates or flows per second.

This family includes the following insights:

| **Insight** | Summary |
| --- | --- |
| **Device Traffic Increase** | The traffic volume on a device is significantly higher than what's normal for that device. |
| **Flows Comparison** | Big change to the week-over-week flow count in the traffic data you are sending to Kentik. |
| **No Flow From Device** | A device that typically sends Kentik more than 5 flows per second is currently sending less than or equal to 5 fps. |
| **TCP Resets from Source IP** | Traffic outbound from an individual source IP address contains a high number of TCP RST packets. |
| **TCP Resets Total** | Traffic outbound from all source IP addresses contains a high number of TCP RST packets. |
| **Interface Utilization Spike** | There's a sudden large increase in the SNMP input bit rate on an interface. |
| **Interface Utilization Drop** | There's a sudden large decrease in the SNMP input bit rate on an interface. |

#### Device Traffic Increase

Identifies a device whose current traffic volume is significantly higher than what's normal for that device. This insight will provide a Kentik-AI generated summary of the possible causes for the traffic increase (see **Device Traffic Increase Insight**).

> **Note**: Kentik AI must be enabled for Cause Analysis to be engaged with Device Traffic Increase (see **Manage Kentik AI**).

* **Evaluation**: Most recent hour vs. same hour of the day for the last two weeks.
* **Dimensions**: Device
* **Metric**: Bits per second
* **Data sources**: All traffic

#### Flows Comparison

Identifies significant week-over-week changes to the flow count in the traffic data you are sending to Kentik.

* **Evaluation**: Most recent 7-day period vs. prior 7-day period.
* **Dimensions**: None
* **Metric**: Flow Count
* **Data sources**: All traffic

#### No Flow From Device

Identifies a device that typically sends Kentik more than 5 flows per second but is currently sending less than or equal to 5 fps.

* **Evaluation**: None
* **Dimensions**: None
* **Metric**: Flow Count
* **Data sources**: All traffic

#### TCP Resets from Source IP

Identifies a large number, significantly above normal, of TCP RST packets headed outbound from an individual Source IP address.

* **Evaluation**: Most recent hour vs. same hour of the day for the last two weeks.
* **Dimensions**: Source IP Address
* **Filters**: Source Network Boundary = Internal, Destination Network Boundary = External, Protocol = TCP, TCP Flags = RST.
* **Metric**: Flow count
* **Data sources**: All traffic

#### TCP Resets Total

Identifies a large number, significantly above normal, of TCP RST packets headed outbound from all Source IP addresses.

* **Evaluation**: Most recent hour vs. same hour of the day for the last two weeks.
* **Dimensions**: Source IP Address
* **Filters**: Source Network Boundary = Internal, Destination Network Boundary = External, Protocol = TCP, TCP Flags = RST.
* **Metric**: Flow count
* **Data sources**: All traffic

#### Interface Utilization Spike

Evaluates, once daily, the SNMP interface input bit rate, identifies sudden and sustained increases in interface utilization, and correlates the SNMP data with flow data for the same interface to determine the most likely cause of the increase.

* **Evaluation**: Interface utilization at any point during the last day vs. a percent of interface capacity.
* **Dimensions**: Interface
* **Metric**: Bits per second
* **Data sources**: All traffic

#### Interface Utilization Drop

Evaluates, once daily, the SNMP interface input bit rate, identifies sudden and sustained decreases in interface utilization, and correlates the SNMP data with flow data for the same interface to determine the most likely cause of the decrease.

* **Evaluation**: Interface utilization at any point during the last day vs. a percent of interface capacity.
* **Dimensions**: Interface
* **Metric**: Bits per second
* **Data sources**: All traffic

## Protect

Insights in the Protect family inform you about threats to network availability and security so that you can take action to defend your network.

This family includes the following insights:

| Insight | Summary |
| --- | --- |
| **Botnets Yesterday** | Changes to the volume and/or ranking of the top 10 conversations to botnets in the traffic outbound from your infrastructure. |

#### Botnets Yesterday

Identifies significant day-over-day changes to the top 10 conversations to botnets in the traffic outbound from your infrastructure (on-prem and/or cloud). The conditions evaluated for major changes include traffic volume for one or more botnets, the makeup of the top 10 botnets set, and changes in top 10 ranking.

* **Evaluation**: Most recent 24-hour period vs. prior 24-hour period.
* **Dimensions**: Destination Botnet
* **Filters**: Simple Traffic Profile = Outbound
* **Metric**: Bits per second
* **Data sources**: All traffic

## Security

Insights in the Security family are based on detecting security risks and vulnerabilities such as traffic transmitted to embargoed countries.

This family includes the following insights:

| **Insight** | Summary |
| --- | --- |
| **Device firmware version check** | A device with deprecated firmware was detected. |
| **Embargoed Country** | Traffic was detected to an embargoed country (e.g., Iran, Islamic Republic of (IR), Sudan (SD), Cuba (CU), Syrian Arab Republic (SY)). |

#### Device firmware version check

Identifies the traffic rate through a device with deprecated/discontinued firmware installed.

* **Evaluation**: Most recent 1-day period
* **Dimensions**: Device
* **Metric**: Bits per second
* **Data sources**: All traffic

#### Embargoed Country

Identifies significant week-over-week changes in traffic rates for each device that sent traffic to an embargoed country.

* **Evaluation**: Most recent 1-day period vs. same period 7 days back
* **Dimensions**: Device, Destination Country
* **Metric**: Bits per second
* **Data sources**: All traffic

## Traffic

For insights in the Traffic family, we evaluate traffic volume to or from your infrastructure (on-prem and/or cloud) and notify you of changes related to top-X rankings. The dimensions we evaluate include applications, protocols, and ASNs (e.g. a given AS moving into or out of the top-10 ASes). The conditions evaluated for a given dimension include:

* Changes in individual or overall traffic volume.
* Changes in the composition of the top 10 set.
* Changes in top 10 ranking.

> **Note:** For the insights in this family, the metrics displayed in the Statistics Pane (see **Insight Details Page**) include bitrate over the most recent 7-day period and percent change from the prior period.

This family includes the following insights:

| **Insight** | Summary |
| --- | --- |
| **Inbound Application Comparison** | Week-over-week changes to the volume and/or ranking of the top 10 applications in the traffic inbound to your infrastructure. |
| **Outbound Application Comparison** | Week-over-week changes to the volume and/or ranking of the top 10 applications in the traffic outbound from your infrastructure. |
| **Inbound ASN Comparison** | Week-over-week changes to the volume and/or ranking of the top 10 source ASNs of the traffic inbound to your infrastructure. |
| **Outbound ASN Comparison** | Week-over-week changes to the volume and/or ranking of the top 10 destination ASNs of the traffic outbound from your infrastructure. |
| **Total Traffic Today Comparison** | Significant change in volume of total traffic for the last 24 hours. |
| **Inbound Site Comparison** | Significant change in volume of traffic inbound to a site. |
| **Outbound Site Comparison** | Significant change in volume of traffic outbound from a site. |

#### Inbound Application Comparison

Identifies significant week-over-week changes to the top 10 applications in the traffic inbound to your infrastructure (on-prem and/or cloud). The conditions evaluated for major changes include traffic volume for one or more applications, the makeup of the top 10 application set, and changes in top 10 ranking.

* **Evaluation**: Most recent 7-day period vs. prior 7-day period.
* **Dimensions**: Source Application
* **Filters**: Simple Traffic Profile = Inbound
* **Metric**: Bits per second
* **Data sources**: All traffic

#### Outbound Application Comparison

Identifies significant week-over-week changes to the top 10 applications in the traffic outbound from your infrastructure (on-prem and/or cloud). The conditions evaluated for major changes include traffic volume for one or more applications, the makeup of the top 10 application set, and changes in top 10 ranking.

* **Evaluation**: Most recent 7-day period vs. prior 7-day period.
* **Dimensions**: Destination Application
* **Filters**: Simple Traffic Profile = Outbound
* **Metric**: Bits per second
* **Data sources**: All traffic

#### Inbound ASN Comparison

Identifies significant week-over-week changes to the top 10 ASNs in the traffic inbound to your infrastructure (on-prem and/or cloud). The conditions evaluated for major changes include traffic volume for one or more ASNs, the makeup of the top 10 ASN set, and changes in top 10 ranking.

* **Evaluation**: Most recent 7-day period vs. prior 7-day period.
* **Dimensions**: Source AS Number
* **Filters**: Simple Traffic Profile = Inbound
* **Metric**: Bits per second
* **Data sources**: All traffic

#### Outbound ASN Comparison

Identifies significant week-over-week changes to the top 10 ASNs in the traffic outbound from your infrastructure (on-prem and/or cloud). The conditions evaluated for major changes include traffic volume for one or more ASNs, the makeup of the top 10 ASN set, and changes in top 10 ranking.

* **Evaluation**: Most recent 7-day period vs. prior 7-day period.
* **Dimensions**: Destination AS Number
* **Filters**: Simple Traffic Profile = Outbound
* **Metric**: Bits per second
* **Data sources**: All traffic

#### Total Traffic Today Comparison

Identifies significant bit rate changes to your total traffic.

* **Evaluation**: Most recent 24-hour period vs. prior 24-hour period.
* **Dimensions**: None
* **Metric**: Bits per second
* **Data sources**: All traffic

#### Inbound Site Comparison

Identifies significant week-over-week changes to traffic inbound to a specific site.

* **Evaluation**: Most recent 7-day period vs. prior 7-day period.
* **Dimensions**: Site
* **Filters**: Simple Traffic Profile = Inbound
* **Metric**: Bits per second
* **Data sources**: All traffic

#### Outbound Site Comparison

Identifies significant week-over-week changes to traffic outbound from a specific site.

* **Evaluation**: Most recent 7-day period vs. prior 7-day period.
* **Dimensions**: Site
* **Filters**: Simple Traffic Profile = Outbound
* **Metric**: Bits per second
* **Data sources**: All traffic

## Traffic Comparisons

Insights in the Traffic Comparisons family involve differences in traffic observed by NetFlow vs. SNMP.

This family includes the following insights:

| **Insight** | Summary |
| --- | --- |
| **Flow and SNMP Difference Detection** | A difference in traffic was detected between NetFlow and SNMP. |

#### Flow and SNMP Difference Detection

Identifies differences in traffic detected between NetFlow data and SNMP data, per interface and per direction.

* **Evaluation**: 24 hours
* **Data sources**: All traffic

Flow:

* **Dimensions**: Device, Src Interface, Dst Interface
* **Filters**: Device, Src/Dst Interface
* **Metric**: Bits per second

SNMP:

* **Dimensions**: Device, Interface
* **Filters**: Device, Interface
* **Metric**: Bits/s In, Bits/s Out

---

© 2014-25 Kentik
