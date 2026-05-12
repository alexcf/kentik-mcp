# General Dimensions

Source: https://kb.kentik.com/docs/general-dimensions

---

This article discusses General Dimensions in Kentik.

> **Note:** As of May 1st, 2025, the **Query SQL Method** has been deprecated and is no longer supported.

## About General Dimensions

General Dimensions in Kentik allow you to filter and group data based on device metrics that are non-flow-related, such as SNMP or Streaming Telemetry. Unlike **Routing Dimensions**, these fields provide context regarding device identity, interface properties, and cloud-specific metadata.

### Where to Find Dimensions

The dimensions available in the **Query** sidebar of **Data Explorer** vary based on the device category (e.g., Router vs. Host) and the specific device type (see **Supported Device Types**).

If you cannot find a specific dimension in this article, refer to Kentik’s specialized references:

* **Traffic Routing**: See **Routing Dimensions** (ASNs, BGP paths, etc.).
* **Device Health**: See **Device Metrics Dimensions** (CPU, Memory, etc.).
* **Specific Hardware**: See **Device-Specific Dimensions** (Vendor-specific attributes).

### Understanding Column Types

To optimize query performance, it is important to understand how Kentik stores and retrieves these dimensions:

| Column Type | Description | Performance Impact |
| --- | --- | --- |
| Native | Stored directly in the **Kentik Data Engine (KDE)**. | **Fastest**: Best for high-volume filtering. |
| Virtual | Derived from other KDE-stored data at query-time. | **Standard**: May impact speed on very large time ranges. |
| UDR | Part of our **Universal Data Records** framework. | **Flexible**: Used for modern host and cloud decodes. |

> **TIP: Naming Note:** Throughout these tables, the Portal Name refers to the label used in the UI, while the KDE Name is the underlying column name required when using Kentik APIs.

## Network and Traffic Topology

Dimensions in this category are used to filter or group-by on device properties, including interface names, descriptions, and port IDs.

### Dimension Mapping in the Portal

In the Kentik Portal, some dimensions are represented differently in the **Dimension Selector** (used for Group By and Matrix) compared to how they appear in **Filtering**.

> **Note:** When using the **Dimension Selector Dialog**, keep the following mappings in mind:
>
> * **Device:** Used to select **Device Name**.
> * **Interface:** Used to select **Interface Name** and **Interface Description**.
> * **Traffic Origination / Termination:** Represented as two distinct dimensions for filtering.
> * **Ultimate Exit Interface:** Used to select **Ultimate Exit Interface Name** and **Description**.

#### Device Info Dimensions

These dimensions provide metadata related to physical/logical devices (see **About Devices**):

| Dimension Name (Portal) | Description | Type: value column | Direction  KDE name(s) |
| --- | --- | --- | --- |
| Device ID | Kentik-assigned unique numerical ID of the device (see **Device General Settings**). | string Virtual | Non-directional: `i_device_id` |
| Device Name | User-defined name for the device (see **Device General Settings**). | string Virtual | Non-directional: `i_device_name` |
| Device Type | Type of device: router, host, etc. (see **Supported Device Types**).  **Note:** Used only for selection (filtering with WHERE clause), not for display or GROUP\_BY. | string Virtual | Non-directional: `i_device_type` |
| Device Sample Rate | The ratio of total flows per sampled flow (see **Device Flow Settings**). | string Virtual | Non-directional:  `device_sample_rate` |
| Site | Name of the site to which the device has been assigned (see **About Sites**). If the device hasn't been assigned to a site, returns an empty string.  **Notes:**   * Supported operators for WHERE clause: case-insensitive equality, LIKE, IN, and regex matching. * Site assignments in the table may lag Admin settings by up to 10 minutes. | string Virtual | Non-directional: `i_device_site_name` |
| Site Market | Name of the site market to which the device has been assigned (see **About Site Markets**). If the device hasn't been assigned to a site market, returns an empty string. | string Virtual | Non-directional: `i_site_market` |
| Device Labels | A label assigned to a collection of devices (see **About Labels**). | string Virtual | Non-directional: `i_device_label` |

#### Interface Info Dimensions

Dimensions related to physical/logical interfaces through which traffic enters or exits your devices (see **About Interfaces**):

| Dimension Name (Portal) | Description | Type: value column | Direction  KDE name(s) |
| --- | --- | --- | --- |
| Interface ID | ID of the receiving/sending host or router interface (see **Interface Field Definitions**). | int Native | Src/Dst: `input_port`, `output_port` |
| Interface Name | The Name (e.g. “GigabitEthernet0/1”) of the device interface (physical or logical) through which flow ingressed/egressed (see **Interface Field Definitions**). | string Virtual | Src/Dst: `i_input_interface_description`, `i_output_interface_description` |
| Interface Description | The Description (e.g. “Connected to upstream ISP”) of the device interface (physical or logical) through which flow ingressed/egressed (see **Interface Field Definitions**). | string Virtual | Src/Dst: `i_input_snmp_alias`, `i_output_snmp_alias` |
| Interface Capacity | The speed of the device interface through which flow ingressed/egressed (see **Interface Field Definitions**). | bigint Virtual | Src/Dst: `i_input_interface_speed`, `i_output_interface_speed` |

#### Interface Classification Dimensions

Classification dimensions allow you to categorize traffic by business logic, such as identifying whether traffic is traversing a peering link versus a transit link. For setup details, see **Interface Classification**.

| Dimension Name (Portal) | Description | Type: value column | Direction  KDE name(s) |
| --- | --- | --- | --- |
| Connectivity Type | The connectivity type, such as transit, IX peering, etc., of the source/destination interface of this flow (see **Connectivity Type Attribute**). | string Virtual | Src/Dst:  `i_src_connect_type_name`, `i_dst_connect_type_name` |
| Network Boundary | The network boundary value (internal or external) of the source/destination interface of this flow (see **Network Boundary Attribute**). | string Virtual | Src/Dst: `i_src_network_bndry_name`, `i_dst_network_bndry_name` |
| Provider | A string representing the provider via which source/destination traffic over a given interface reaches the Internet (see **About Provider Classification**). | string Virtual | Src/Dst: `i_src_provider_classification` `i_dst_provider_classification` |

#### Network Classification Dimensions

Dimensions in this category help you identify the flow’s relationship to your network boundaries, specifically whether traffic is staying internal, entering from the outside, or leaving your network. For a deeper dive into the logic, see **Network Classification**.

| Dimension Name (Portal) | Description | Type: value column | Direction  KDE name(s) |
| --- | --- | --- | --- |
| Traffic Orig/Term | Indicates the location (inside or outside) of the source/destination of the flow | string Virtual | Src/Dst: `i_trf_origination`, `i_trf_termination` |
| Host Direction | If flow record is from host, indicates whether the direction of traffic is into or out of that host. | string Virtual | Non-directional: `i_host_direction` |
| Traffic Profile | The origination and termination of the flow. | string Virtual | Non-directional: `i_trf_profle` |
| Simple Traffic Profile | Alternate dimension for origination and termination of the flow. | string Virtual | Non-directional: `simple_trf_prof` |

> **Note:** For more technical details on how these are calculated, see **Network Classification Dimensions**.

#### Ultimate Exit Dimensions

Dimensions in this category identify the specific interface and device through which traffic leaves your network to reach an external destination (e.g., a peering partner or transit provider). For a detailed explanation of the underlying logic, see **Using Ultimate Exit****.**

| Dimension Name (Portal) | Description | Type:  value  column | Direction  KDE name(s) |
| --- | --- | --- | --- |
| Ultimate Exit Interface ID | Number of ports through which the flow leaves (see **Network Classification Dimensions**). | bigint Native | Non-directional: `ult_exit_port` |
| Ultimate Exit Interface Name | The SNMP description (portal name) of the interface through which the flow leaves (see **Network Classification Dimensions**). | string Virtual | Non-directional: `i_ult_exit_interface_description` |
| Ultimate Exit Interface Description | The SNMP alias (portal description) of the interface through which the flow leaves (see **Network Classification Dimensions**). | string Virtual | Non-directional: `i_ult_exit_snmp_alias` |
| Ultimate Exit Connectivity Type | The connectivity type value of the interface through which traffic left the network for another AS (see **Network Classification Dimensions**). | string Virtual | Non-directional: `i_ult_exit_connect_type_name` |
| Ultimate Exit Network Boundary | The network boundary value of the interface through which traffic left the network for another AS (see **Network Classification Dimensions**). | string Virtual | Non-directional: `i_ult_exit_network_bndry_name` |
| Ultimate Exit Provider | A string representing the ultimate exit provider (see **Why Ultimate Exit**). | string Virtual | Non-directional: `i_ult_provider_classifcation` |
| Ultimate Exit Site | The name of the site through which the flow leaves (see **Why Ultimate Exit**). | string Virtual | Non-directional: `i_ult_exit_site` |
| Ultimate Exit Site Market | The name of the site market through which the flow leaves (see **Why Ultimate Exit**). | string Virtual | Non-directional:  `i_ult_exit_site_market` |
| Ultimate Exit Device ID | The numerical ID of the device through which the flow leaves (see **Why Ultimate Exit**). | string Virtual | Non-directional: `i_ult_exit_device_id` |
| Ultimate Exit Device | The name of the device through which the flow leaves (see **Why Ultimate Exit**). | string Virtual | Non-directional: `i_ult_exit_device_name` |

#### LAN Dimensions

Dimensions in this category provide visibility into the Layer 2 (Data Link) properties of your network traffic, allowing you to filter by VLAN IDs and hardware-specific MAC addresses.

| Dimension name (portal) | Description | Type: value column | Direction  KDE name(s) |
| --- | --- | --- | --- |
| VLAN | ID of receiving/sending VLAN. | int Native | Src/Dst: `vlan_in`, `vlan_out` |
| MAC Address | Ethernet (L2) address of source/destination. | string Native | Src/Dst: `src_eth_mac`, `dst_eth_mac` |

## Cloud Dimensions

These dimensions are used to filter or group-by on fields within VPC flow logs from cloud providers. They allow you to treat cloud resources (like Instances or Subnets) with the same granularity as physical hardware.

### General Cloud Dimensions

The following dimensions are applicable across all cloud providers supported by Kentik (e.g., AWS, GCP, or Azure).

| Dimension Name (Portal) | Description | Type: value column | Direction  KDE name(s) |
| --- | --- | --- | --- |
| Cloud Provider | The cloud provider (e.g., AWS, GCP, or Azure) from which Kentik retrieved the flow log containing the data in this flow record. | string native | Non-directional: `kt_cloud_provider` |

### AWS Dimensions

The following dimensions represent metadata extracted from Amazon VPC Flow Logs (see **Kentik for AWS**). For a complete list of field definitions, refer to the **Amazon VPC User Guide**.

#### Directional AWS Dimensions

These fields allow you to filter based on whether the AWS resource is the **Source** or **Destination** of the traffic.

| Dimension Name (Portal) | Description | Type: value column | Direction  KDE name(s) |
| --- | --- | --- | --- |
| Account ID | Source/destination AWS account identifier | int Virtual | Src/Dst: `kt_aws_src_acc_id`, `kt_aws_dst_acc_id` |
| Account Alias | Source/destination AWS account alias | string  Virtual | Src/Dst:  `kt_aws_src_acc_alias`,  `kt_aws_dst_acc_alias` |
| Account Name | Source/destination AWS account name | string  Virtual | Src/Dst:  `kt_aws_src_acc_name`,  `kt_aws_dst_acc_name` |
| Instance Name | Source/destination AWS instance name | string Virtual | Src/Dst: `kt_aws_src_vm_name`, `kt_aws_dst_vm_name` |
| Instance | Source/destination AWS instance | string Virtual | Src/Dst: `kt_aws_src_vm_id`, `kt_aws_dst_vm_id` |
| Region | Source/destination AWS Region | string Virtual | Src/Dst: `kt_aws_src_region`, `kt_aws_dst_region` |
| Zone ID | Source/destination AWS Availability Zone identifier | string  Virtual | Src/Dst:  `kt_aws_src_zone_id`,  `kt_aws_dst_zone_id` |
| Zone Name | Source/destination AWS Availability Zone name | string Virtual | Src/Dst: `kt_aws_src_zone`, `kt_aws_dst_zone` |
| Instance Type | Source/destination AWS Instance Type | string Virtual | Src/Dst:  `kt_aws_src_vm_type`,  `kt_aws_dst_vm_type` |
| Image ID | Source/destination AWS Image ID | string Virtual | Src/Dst: `kt_aws_src_image_id`, `kt_aws_dst_image_id` |
| Security Group | Source/destination security group | string Virtual | Src/Dst: `kt_aws_src_sg`, `kt_aws_dst_sg` |
| Auto Scaling Group | Source/destination auto scaling group | string Virtual | Src/Dst: `kt_aws_src_asg`, `kt_aws_dst_asg` |
| Public DNS Name | Source/destination public DNS name | string Virtual | Src/Dst: `kt_aws_src_pub_dns`, `kt_aws_dst_pub_dns` |
| Private DNS Name | Source/destination private DNS name | string Virtual | Src/Dst: `kt_aws_src_priv_dns`, `kt_aws_dst_priv_dns` |
| VPC ID | Source/destination VPC ID | string Virtual | Src/Dst: `kt_aws_src_vpc_id`, `kt_aws_dst_vpc_id` |
| Subnet ID | Source/destination subnet ID | string Virtual | Src/Dst: `kt_aws_src_subnet_id`, `kt_aws_dst_subnet_id` |
| Instance Tags | Tags applied to VMs by users | string Virtual | Src/Dst: `kt_aws_src_vm_tags`, `kt_aws_dst_vm_tags` |
| Packet Address | The packet-level (original) source/destination IP address of the traffic. See pkt-srcaddr/pkt-dstaddr in AWS documentation. | bytes/IP Address Native | Src/Dst: `ktsubtype__aws_subnet__INET_00`, `ktsubtype__aws_subnet__INET_01` |
| Gateway ID | The ID of the gateway through which the flow entered your AWS resources. | string Virtual | Dst only: `ktsubtype__aws_subnet__STR17` |
| Gateway Type | The type of the gateway through which the flow entered your AWS resources | string Virtual | Dst only: `ktsubtype__aws_subnet__STR19` |
| Forwarding State | The route state of the destination prefix:   * “active” if traffic is flowing towards an active route; * “blackholed” if traffic is flowing towards a blackhole route. | string Virtual | Dst only: `ktsubtype__aws_subnet__STR04`, `ktsubtype__aws_subnet__STR05` |
| Interface ID | The ID (inferred from IP address) of the first (for source) or last (for destination) Elastic Network Interface on which the flow was recorded. | string Virtual | Src only: `ktsubtype__aws_subnet__STR20` |
| Interface Type | The type of the network interface that recorded the flow:   * 0 - No value provided * 1 - Unknown * 2 - interface * 3 - nat\_gateway * 4 - lambda * 5 - transit\_gateway * 6 - vpc\_endpoint * 7 -network\_load\_balancer * 8 - gateway\_load\_balancer\_endpoint * 9 - trunk * 10 - global\_accelerator\_managed | int Virtual | Src/Dst: `ktsubtype__aws_subnet__INT02`, `ktsubtype__aws_subnet__INT03` |
| AWS Service | The name of the subset of IP address ranges for the pkt-srcaddr field, if the source IP address is for an AWS service. For possible values see pkt-src-aws-service in AWS documentation. | int Virtual | Src/Dst: `ktsubtype__aws_subnet__INT04`, `ktsubtype__aws_subnet__INT05` |
| ENI Description | The description field of the Elastic Network Interface that recorded the flow. | string Virtual | Src/Dst: `ktsubtype__aws_subnet__STR21`, `ktsubtype__aws_subnet__STR22` |
| ENI Entity Name | The name of the entity based on the Elastic Network Interface that recorded the flow. | string Virtual | Src/Dst: `ktsubtype__aws_subnet__STR23`, `ktsubtype__aws_subnet__STR24` |
| AWS TGW AZ ID | The ID of the transit gateway availability zone. | string Virtual | Src/Dst: `ktsubtype__aws_subnet__STR35`, `ktsubtype__aws_subnet__STR29` |
| AWS TGW ENI | The Elastic Network Interface of the transit gateway. | string  Virtual | Src/Dst: `ktsubtype__aws_subnet__STR36`, `ktsubtype__aws_subnet__STR30` |
| AWS TGW Subnet ID | The subnet ID of the transit gateway. | string  Virtual | Src/Dst: `ktsubtype__aws_subnet__STR37`, `ktsubtype__aws_subnet__STR31` |
| AWS TGW VPC ID | The VPC ID of the transit gateway. | string  Virtual | Src/Dst: `ktsubtype__aws_subnet__STR38`, `ktsubtype__aws_subnet__STR32` |
| AWS TGW VPC Account ID | The VPC account ID of the transit gateway. | bigint  Virtual | Src/Dst: `ktsubtype__aws_subnet__INT64_05`, `ktsubtype__aws_subnet__INT64_06` |

#### Non-directional AWS Dimensions

These fields describe the status or path of the flow as recorded by AWS.

| Dimension Name (Portal) | Description | Type: value column | KDE name(s) |
| --- | --- | --- | --- |
| Firewall Action | The action associated with the traffic:   * ACCEPT: The recorded traffic was permitted by the security groups or network ACLs. * REJECT: The recorded traffic was not permitted by the security groups or network ACLs. | string Native | `kt_aws_action` |
| Logging Status | The logging status of the flow log:   * OK: Data is logging normally to the chosen destinations. * NODATA: There was no network traffic to or from the network interface during the capture window. * SKIPDATA: Some flow log records were skipped during the capture window. This may be because of an internal capacity constraint, or an internal error. | string Native | `kt_aws_status` |
| Start Time | The time, in Unix seconds, when the first packet of the flow was received within the aggregation interval. This might be up to 60 seconds after the packet was transmitted or received on the network interface. | bigint Native | `ktsubtype__aws_subnet__INT00` |
| End Time | The time, in Unix seconds, when the last packet of the flow was received within the aggregation interval. This might be up to 60 seconds after the packet was transmitted or received on the network interface. | bigint Native | `ktsubtype__aws_subnet__INT01` |
| Observing Interface ID | The ID of the network interface that recorded the flow. | string Native | `ktsubtype__aws_subnet__STR03` |
| Flow Log Account ID | The AWS account ID of the owner of the source network interface that recorded the flow.  **Note:** This value may be unknown when the interface is created by an AWS service, e.g., when creating a VPC endpoint or Network Load Balancer. | string Native | `ktsubtype__aws_subnet__INT64_00` |
| Flow Direction | The direction of the flow with respect to the interface where traffic is captured:   * ingress * egress | string Native | `ktsubtype__aws_subnet__INT06` |
| Traffic Path | The path that egress traffic (see Flow Direction) takes to the destination:  1 - Through another resource in the same VPC  2 - Through an internet gateway or a gateway VPC endpoint  3 - Through a virtual private gateway  4 - Through an intra-region VPC peering connection  5 - Through an inter-region VPC peering connection  6 - Through a local gateway  7 - Through a gateway VPC endpoint (Nitro-based instances only)  8 - Through an internet gateway (Nitro-based instances only)  **Note:** If none of the above values apply, the field is set to "-". | int Native | `ktsubtype__aws_subnet__INT07` |
| Cloud Ultimate Exit | The last gateway the flow will traverse on its way to the destination IP address. | string Virtual | `ktsubtype__aws_subnet__STR25` |
| Cloud Ultimate Exit Type | The type of the last gateway the flow will traverse on its way to the destination IP address:   * Virtual Gateway * Customer Gateway * Transit * Internet * VPC Peering Gateway * Egress Only Internet Gateway * NAT Gateway * Carrier Gateway | string Virtual | `ktsubtype__aws_subnet__STR26` |
| Observing VPC ID | The ID of the observing VPC. | string Native | `ktsubtype__aws_subnet__STR10` |
| Observing VPC Subnet ID | The ID of the observing VPC subnet. | string Native | `ktsubtype__aws_subnet__STR11` |
| Observing Instance ID | The ID of the observing instance. | string Native | `ktsubtype__aws_subnet__STR12` |
| Observing Region | The ID of the observing region. | string Native | `ktsubtype__aws_subnet__STR14` |
| Observing Availability Zone ID | The ID of the observing availability zone. | string Native | `ktsubtype__aws_subnet__STR15` |
| Observing Sublocation ID | The ID of the observing subscription. | string Native | `ktsubtype__aws_subnet__STR16` |
| AWS Resource Type | The type of AWS resource being observed such as EC2, RDS, etc. | string Virtual | `ktsubtype__aws_subnet__STR27` |
| AWS TGW Attachment ID | The ID of the transit gateway attachment. | string Virtual | `ktsubtype__aws_subnet__STR28` |
| AWS TGW ID | The ID of the transit gateway. | string Virtual | `ktsubtype__aws_subnet__STR33` |
| AWS TGW Pair Attachment ID | The ID of the paired transit gateway attachment. | string Virtual | `ktsubtype__aws_subnet__STR34` |
| AWS Packets Lost - Blackhole | The number of packets lost due to blackhole routing. | int Virtual | `ktsubtype__aws_subnet__INT64_01` |
| AWS Packets Lost - MTU Exceeded | The number of packets lost because the packet size exceeded the MTU. | int Virtual | `ktsubtype__aws_subnet__INT64_02` |
| AWS Packets Lost - No Route | The number of packets lost due to the absence of a valid route. | int Virtual | `ktsubtype__aws_subnet__INT64_03` |
| AWS Packets Lost - TTL Expired | The number of packets lost because the TTL expired. | int Virtual | `ktsubtype__aws_subnet__INT64_04` |

### Azure Dimensions

Dimensions in this category allow you to analyze traffic within your Azure Virtual Networks (VNets), providing visibility into VMs, Resource Groups, and security rules (see **Kentik for Azure**).

#### Directional Azure Dimensions

These fields allow you to filter based on whether the Azure resource is the **Source** or **Destination** of the traffic.

| Dimension Name (Portal) | Description | Type:  value  column | Direction  KDE name(s) |
| --- | --- | --- | --- |
| Instance Name | The name of the Azure instance (VM) that generated the flow log. | string  Virtual | Src/Dst: `kt_az_src_inst_name`, `kt_az_dst_inst_name` |
| Instance | The raw ID of the log-generating instance, which is useful for programmatic management of compute resources. | string  Virtual | Src/Dst: `kt_az_src_inst_id`, `kt_az_dst_inst_id` |
| Region | The geographical region of the Azure instance, which corresponds to a specific set of Azure data centers in which the instance may run. | string  Virtual | Src/Dst: `kt_az_src_region`, `kt_az_dst_region` |
| Zone | The High Availability Zone where the instance is currently deployed, which corresponds to a specific data center within a region. | int  Virtual | Src/Dst: `kt_az_src_zone`, `kt_az_dst_zone` |
| Instance Type | The kind of instance-generated flow logs, which may be Azure-provided or custom-built. These values do not follow a standard naming nomenclature. | string  Virtual | Src/Dst: `kt_az_src_inst_type`, `kt_az_dst_inst_type` |
| Public DNS Name | The publicly resolvable DNS name for an instance. | string  Virtual | Src/Dst: `kt_az_src_fqdn`, `kt_az_dst_fqdn` |
| VNet ID | An identifier for the virtual network object in which an instance resides. A virtual network is a collection of subnets within a given region. | string  Virtual | Src/Dst: `kt_az_src_vnet`, `kt_az_dst_vnet` |
| Subnet Name | The name of a subnet resource assigned to a virtual network. | string  Virtual | Src/Dst: `kt_az_src_subnet`, `kt_az_dst_subnet` |
| Resource Group | A set of related technical resources (disk, storage, VMs, APIs, services, etc.) that can be accessed as a group for bulk operations. | string  Virtual | Src/Dst: `kt_az_src_resource_group`, `kt_az_dst_resource_group` |
| Public IP Address | The public IP address assigned to an Azure instance. Public IP addresses are not assigned by default. | string  Virtual | Src/Dst: `kt_az_src_public_ip`, `kt_az_dst_public_ip` |
| Subscription | A top-level administrative object representing a set of resources that will be billed together in a monthly cycle. All Azure resources are tied to a subscription, which may contain multiple resource groups. | string  Virtual | Src/Dst: `kt_az_src_sub_id`, `kt_az_dst_sub_id` |
| Security Rule | The name of the security rule by which this flow was allowed or denied as it passed through a security group (see below) on its way to or from an Azure instance. | string  Virtual | Src/Dst: `ktsubtype__azure_subnet__STR01`, `ktsubtype__azure_subnet__STR00` |
| Firewall Action | The actions (allow or deny) taken on this flow by the security rules by which it was evaluated on the way to or from an Azure instance. | string  Virtual | Src/Dst: `ktsubtype__azure_subnet__STR03`, `ktsubtype__azure_subnet__STR02` |
| Security Group | A collection of enforced security policies (each a collection of rules) at the edge of a virtual network and/or applied to a network interface attached to an instance. Traffic to an instance from the internet must pass through at least one security group at the edge of the virtual network and may also pass through an additional security group attached to the interface of an instance. | string  Virtual | Src/Dst: `kt_az_src_nsg_name`, `kt_az_dst_nsg_name` |
| Interface Name | The name of the network interface. | string  Virtual | Src/Dst: `ktsubtype__azure_subnet__STR06`, `ktsubtype__azure_subnet__STR07` |
| Gateway Name | The name of the gateway. | string  Virtual | Src/Dst: `ktsubtype__azure_subnet__STR09`, `ktsubtype__azure_subnet__STR10` |
| Gateway Type | The type of gateway such as VPN, ExpressRoute, Application Gateway, etc. | string  Virtual | Src/Dst: `ktsubtype__azure_subnet__STR11`, `ktsubtype__azure_subnet__STR12` |
| FQDN | The fully qualified domain name. | string  Virtual | Dst: `ktsubtype__azure_subnet__STR19` |
| Load Balancer | The Azure load balancer resource that distributes traffic. | string  Virtual | Src/Dst: `ktsubtype__azure_subnet__STR23`, `ktsubtype__azure_subnet__STR24` |

#### Non-directional Azure Dimensions

These fields provide context on how and where the log was observed within the Azure infrastructure.

| Dimension Name (Portal) | Description | Type: value column | KDE name(s) |
| --- | --- | --- | --- |
| Observing MAC Address | The MAC address of the observing device. | string  Virtual | `ktsubtype__azure_subnet__STR08` |
| Ingress Virtual Hub | The name of the Azure Virtual Hub where the traffic enters. | string Virtual | `ktsubtype__azure_subnet__STR13` |
| Egress Virtual Hub | The name of the Azure Virtual Hub where the traffic exits. | string  Virtual | `ktsubtype__azure_subnet__STR14` |
| ExpressRoute Circuit Name | The name of the Azure ExpressRoute circuit. | string Virtual | `ktsubtype__azure_subnet__STR15` |
| ExpressRoute Peering Type | The name of peering type of the Azure ExpressRoute. | string Virtual | `ktsubtype__azure_subnet__STR16` |
| Firewall Policy | The policy that governs that handling of traffic. | string Virtual | `ktsubtype__azure_subnet__STR20` |
| Firewall Rule | The rule within the policy that allows or denies traffic. | string Virtual | `ktsubtype__azure_subnet__STR21` |
| Application Protocol | The application layer protocol used in the communication process. | string Virtual | `ktsubtype__azure_subnet__STR22` |
| Logging Resource Category | The classification of the logging resource. | string Virtual | `ktsubtype__azure_subnet__STR17` |
| Logging Resource Name | The name of the logging resource. | string Virtual | `ktsubtype__azure_subnet__STR18` |

### GCP Dimensions

Dimensions in this category allow you to analyze traffic within your Google Cloud environments, providing visibility into resources like Compute Engine (GCE) instances and Virtual Private Clouds (VPC); see **Kentik for GCP**.

#### Directional GCP Dimensions

These fields allow you to filter based on whether the GCP resource is the **Source** or **Destination** of the traffic.

| Dimension name (portal) | Description | Type: value column | Direction  KDE name(s) |
| --- | --- | --- | --- |
| Project ID | The Google Compute Engine (GCE) Project ID. | string Native | Src/Dst: `kt_gce_src_proj_id`, `kt_gce_dst_proj_id` |
| VM Name | The GCE VM name. | string Native | Src/Dst: `kt_gce_src_vm_name`, `kt_gce_dst_vm_name` |
| Region | The GCE region. | string Native | Src/Dst: `kt_gce_src_region`, `kt_gce_dst_region` |
| Zone | The GCE zone. | string Native | Src/Dst: `kt_gce_src_zone`, `kt_gce_dst_zone` |
| Subnet Name | The GCE subnet name. | string Native | Src/Dst: `kt_gce_src_vpc_snn`, `kt_gce_dst_vpc_snn` |
| VM Type | The GCE VM type. | string Virtual | Src/Dst: `kt_gce_src_vm_type`, `kt_gce_dst_vm_type` |
| Image ID | The GCE image ID. | string Virtual | Src/Dst: `kt_gce_src_vm_image`, `kt_gce_dst_vm_image` |
| Instance Group ID or Name | The GCE instance group ID or name. | string Virtual | Src/Dst: `kt_gce_src_vm_group`, `kt_gce_dst_vm_group` |
| VPC Name | The name of the VPC. | string Native | Src/Dst: `ktsubtype__gcp_subnet__STR11`, `ktsubtype__gcp_subnet__STR12` |
| GKE Pod Name | The name of the individual pod in GKE (Google Kubernetes Engine). | string Native | Src/Dst: `kt_k8s_src_pod_name`, `kt_k8s_dst_pod_name` |
| GKE Pod Namespace | The namespace that the pod is a part of in GKE. | string Native | Src/Dst: `kt_k8s_src_pod_ns`, `kt_k8s_dst_pod_ns` |
| GKE Cluster Name | The name of the GKE cluster. | string Native | Src/Dst: `kt_k8s_src_load_name`, `kt_k8s_src_load_name` |
| GKE Cluster Location | The location of the GKE cluster. | string Native | Src/Dst: `kt_k8s_src_load_ns`, `kt_k8s_src_load_ns` |
| GKE Service Name | The name of the GKE service. | string Native | Src/Dst: `kt_k8s_src_svc_name`, `kt_k8s_dst_svc_name` |
| GKE Service Namespace | The namespace that the GKE is a part of. | string Native | Src/Dst: `kt_k8s_src_svc_ns`, `kt_k8s_src_svc_ns` |
| Entity Gateway Name | The name of the gateway entity. | string Virtual | Src/Dst: `ktsubtype__gcp_subnet__STR27`, `ktsubtype__gcp_subnet__STR28` |
| Entity Gateway Type | The type of gateway. | string Virtual | Src/Dst: `ktsubtype__gcp_subnet__STR29`, `ktsubtype__gcp_subnet__STR30` |

#### Non-directional GCP Dimensions

| Dimension Name (Portal) | Description | Type: value column | KDE name(s) |
| --- | --- | --- | --- |
| Reporter | Indicates where the flow was collected/reported:   * By the source VM/instance if value is SRC; * By the destination VM/instance if value is DEST. | string Native | `kt_gce_reporter` |
| Dedicated Interconnect Name | The name of the GCP Dedicated Interconnect. | string Virtual | `ktsubtype__gcp_subnet__STR25` |
| Dedicated Interconnect Type | The type of interconnection established. | string Virtual | `ktsubtype__gcp_subnet__STR26` |

#### Non-directional Google Cloud Run (GCR) Dimensions

Specialized dimensions for serverless workloads running on Google Cloud Run.

| Dimension Name (Portal) | Description | Type: value column | KDE name(s) |
| --- | --- | --- | --- |
| Project ID | The GCP Cloud Run project ID. | string Virtual | `ktsubtype__gcp_subnet__STR04` |
| Resource Type | The GCP Cloud Run resource type. | string Virtual | `ktsubtype__gcp_subnet__STR00` |
| Service Name | The GCP Cloud Run service name. | string  Virtual | `ktsubtype__gcp_subnet__STR01` |
| Location | The region where the GCP Cloud Run service is deployed. | string Virtual | `ktsubtype__gcp_subnet__STR02` |
| Service Revision Name | The name of the GCP Cloud Run service revision. | string Virtual | `ktsubtype__gcp_subnet__STR03` |
| HTTP Status Code | The HTTP response status code indicating the result of the request processed by the GCP Cloud Run service. | int Virtual | `ktsubtype__gcp_subnet__INT00` |

### OCI Dimensions

Dimensions in this category provide visibility into your Oracle Cloud Virtual Cloud Networks (VCN), covering instances, subnets, and specialized gateways (see **Kentik for OCI**).

> **Notes:**
>
> * All KDE columns for OCI dimensions are UDR (see **Universal Data Records**).
> * Additional details about OCI flow logs may be found in the OCI documentation topic **Details for VCN Flow Logs**.

#### Directional OCI Dimensions

These fields allow you to filter based on whether the OCI resource is the **Source** or **Destination** of the traffic.

| Dimension name (portal) | Description | Type: value column | KDE name(s) |
| --- | --- | --- | --- |
| Subnet Name | The name of the subnet. | string Virtual | Src/Dst: `ktsubtype__oci_subnet__STR06`, `ktsubtype__oci_subnet__STR07` |
| VCN Name | The name of the virtual cloud network. | string Virtual | Src/Dst: `ktsubtype__oci_subnet__STR08`, `ktsubtype__oci_subnet__STR09` |
| Region | The region where the resource is located. | string Virtual | Src/Dst: `ktsubtype__oci_subnet__STR10`, `ktsubtype__oci_subnet__STR11` |
| VNIC Name | The name of the virtual network interface card. | string Virtual | Src/Dst: `ktsubtype__oci_subnet__STR12`, `ktsubtype__oci_subnet__STR13` |
| Instance Name | The name of the instance. | string Virtual | Src/Dst: `ktsubtype__oci_subnet__STR14`, `ktsubtype__oci_subnet__STR15` |
| Dynamic Routing Gateway Name | The name of the dynamic routing gateway. | string Virtual | Src/Dst: `ktsubtype__oci_subnet__STR16`, `ktsubtype__oci_subnet__STR17` |
| Local Peering Gateway Name | The name of the local peering gateway. | string Virtual | Src/Dst: `ktsubtype__oci_subnet__STR18`, `ktsubtype__oci_subnet__STR19` |

#### Non-directional OCI Dimensions

These fields provide administrative context and connectivity details for OCI flows.

| Dimension Name (Portal) | Description | Type: value column | KDE name(s) |
| --- | --- | --- | --- |
| Action | Indicates whether the record’s traffic was accepted or rejected by the security lists. | string Native | `ktsubtype__oci_subnet__STR00` |
| Status | The status of the data capture window for the flow log. | string Native | `ktsubtype__oci_subnet__STR01` |
| Tenant OCID | The Oracle Cloud ID of the tenant. | string Native | `ktsubtype__oci_subnet__STR02` |
| VNIC OCID | The Oracle Cloud ID of the virtual network interface card. | string Native | `ktsubtype__oci_subnet__STR03` |
| VNIC Compartment OCID | The Oracle Cloud ID of the compartment to which the VNIC belongs. | string Native | `ktsubtype__oci_subnet__STR04` |
| VNIC Subnet OCID | The Oracle Cloud ID of the subnet to which the VNIC belongs. | string Native | `ktsubtype__oci_subnet__STR05` |
| Internet Gateway | The OCI Internet Gateway that provides a path for the network traffic between the VCN and the internet. | string Virtual | `ktsubtype__oci_subnet__STR20` |
| NAT Gateway | The OCI Network Address Translation gateway that enables instances to initiate connections to the internet. | string Virtual | `ktsubtype__oci_subnet__STR21` |
| Service Gateway | The OCI Service Gateway that provides a path for the network traffic between the VCN and the other OCI resources. | string Virtual | `ktsubtype__oci_subnet__STR22` |
| IPsec Connection Name | The name of the IPsec VPN connection. | string Virtual | `ktsubtype__oci_subnet__STR23` |
| Virtual Circuit Name | The name of the OCI Virtual Circuit. | string Virtual | `ktsubtype__oci_subnet__STR24` |

## Geolocation Dimensions

These dimensions allow you to filter or group-by the geographical properties of the source and destination IP addresses. Kentik uses a high-accuracy provider to map these coordinates.

| Dimension name (portal) | Description | Type: value column | Direction KDE name(s) |
| --- | --- | --- | --- |
| Custom Geo | A collection of countries that have been assigned a common geographical label (see **About Custom Geos**). | string Native | Src/Dst: kt\_src\_market, kt\_dst\_market |
| Country | Two-letter country code associated with the source/destination IP of the flow. | string Native | Src/Dst: src\_geo, dst\_geo |
| Region | Full-string English name of the region (state or province, e.g. "California") associated with the source IP of the flow. | string Native | Src/Dst: src\_geo\_region, dst\_geo\_region |
| City | Full-string English name of the city (e.g. "San Francisco") associated with the source IP of the flow. | string Native | Src/Dst: src\_geo\_city, dst\_geo\_city |
| Site Country | A country in which your organization has sites; enables the grouping, with a single dimension, of traffic from all sites in a given country. | string Virtual | Non-directional: i\_device\_site\_country |
| Ultimate Exit Site Country | The name of the country containing the site through which flow leaves. | string Virtual | Non-directional: i\_ult\_exit\_site\_country |

## Application Context and Security

Dimensions in this category provide high-level context about the nature of the traffic, such as identifying commercial **CDNs** and **Cloud Services**, or matching traffic against known **Threat Feeds**.

These dimensions are used to filter or group-by based on various factors related to context — whether a flow originated or terminated with a commercial CDN, for example, or what "service" (port and protocol) it represents — as well as whether the value of certain flow fields match those of known security threats.

| Dimension Name (Portal) | Description | Type: value column | Direction  KDE name(s) |
| --- | --- | --- | --- |
| Cloud | The name of the vendor (e.g., AWS, GCP, Azure, etc.) operating the cloud computing service in which this flow originated (src) or terminated (dst). The value is derived by checking the IP address (src or dst) in the flow against the cloud provider's list of IPs. | string Native | Src/Dst: `kt_src_cloud kt_dst_cloud` |
| Cloud Service | The name that a cloud computing vendor assigns to the service in which a flow originated (src) or terminated (dst). The value is derived by checking the IP address (src or dst) in the flow against the cloud provider's list of IPs. | string Native | Src/Dst: `kt_src_cloud_service`,`kt_dst_cloud_service` |
| CDN | Commercial CDN (if any) with which the flow originated/terminated (see **CDN Attribution Dimensions**).  **Note:** This dimension is available only for organizations with CDN Attribution enabled. | string Native | Src/Dst: `src_cdn`, `dst_cdn` |
| Service (Port + Proto) | The combination of the port and protocol of the source/destination flow.  **Note:** This dimension is available only for group-by. For filtering, use Port Number and Protocol. | string Virtual | Src/Dst: N.A. |
| Bot Net CC | A source/destination IP for the flow that has been identified as a botnet command and control (CC) servers (see **Threat Feed Dimensions**). | string Native | Src/Dst: `src_threat_bnetcc`, `dst_threat_bnetcc` |
| Threat List Host | A source/destination IP for the flow that has been identified as a threat (see **Threat Feed Dimensions**). | string Native | Src/Dst: `src_threat_host`, `dst_threat_host` |
| Application | An identifying string for the application associated with a flow, which is either derived by evaluating flow data or provided in the flow data itself (see **About Applications**). | string Native | Non-directional: application |
| TCP Flags | TCP flags that were set on the flow using a flow mask. | int Native | Non-directional: `tcp_flags` |
| OTT Service | An individual OTT content service whose hostname is looked up via DNS. | string Native | Non-directional: `ott_service` |
| OTT Service Type | The nature of the content provided by an OTT content service. Values include Adult, Ads, Antivirus, Audio, Cloud, Conferencing, Dating, Developer Tools, Documents, Ecommerce, File Sharing, Gaming, IoT, Mail, Maps, Media, Messaging, Network, Newsgroups, Photo Sharing, Social, Software Download, Software Updates, Storage, Video, VPN, Web. | string Virtual | Non-directional: N.A. |
| OTT Service Provider | An entity that offers an OTT content service. For example, Google is the provider for Google Drive, Gmail, Google Maps, etc. | string Virtual | Non-directional: N.A. |

## Application Decodes

Application decodes dimensions allow you to filter or group-by traffic based on host-related data (e.g., HTTP, DNS, TLS) captured and enriched by Kentik (see **Universal Agent**).

### Dimensional Evolution: Native vs. UDR

* **Current Dimensions:** Modern versions of the host agent use **Universal Data Records** columns. These are found under the **Application Decodes** category in the Portal.
* **Legacy Dimensions:** Older agent versions used a fixed set of **Native KDE columns**. These are now categorized as **Legacy Application Decodes**.

> **IMPORTANT**: All dimensions listed in the tables below are **Non-directional**. For associated performance data, see **Application Decodes Metrics**.

### Application Decodes Dimensions

The following dimensions correspond to application decode fields from current host agent versions, which use UDR columns in KDE (see **Universal Data Records**).

#### DNS Dimensions

Dimensions related to DNS properties (see **Host Traffic Dimensions**):

| Dimension Name (Portal) | Description | Type: value column | Direction |
| --- | --- | --- | --- |
| DNS Query Name | Query from a DNS resolver to a DNS name server. | string UDR | Non-directional |
| DNS Query Type | The resource record type requested by the DNS query. | bigint UDR | Non-directional |
| DNS Reply Code | DNS return code (see **IANA DNS Parameters**). | bigint UDR | Non-directional |
| DNS Reply Data | The response from a DNS server to a DNS query. | string UDR | Non-directional |

#### HTTP Dimensions

Dimensions related to HTTP properties (see **Host Traffic Dimensions**):

| Dimension Name (Portal) | Description | Type: value column | Direction |
| --- | --- | --- | --- |
| HTTP URL | Filename portion of path, with query string (if any). | string UDR | Non-directional |
| HTTP Host | Domain name of the server. | string UDR | Non-directional |
| HTTP Referrer | The address from which a destination webpage is requested. | string UDR | Non-directional |
| HTTP URL | Filename portion of path, with query string (if any). | string UDR | Non-directional |
| HTTP Host | Domain name of the server. | string UDR | Non-directional |

#### TLS Dimensions

Dimensions related to Transport Layer Security (see **IETF RFC8446**):

| Dimension Name (Portal) | Description | Type: value column | Direction |
| --- | --- | --- | --- |
| TLS Server Name | The Server Name Indication (SNI), which is a TLS extension by which a client indicates which hostname it is attempting to connect to at the start of the handshaking process. | string UDR | Non-directional |
| TLS Server Version | The TLS server version. | int UDR | Non-directional |
| TLS Cipher Suite | A set of cryptographic algorithms used to create keys and encrypt information for TLS. | int UDR | Non-directional |

#### DHCP Dimensions

Dimensions related to Dynamic Host Configuration Protocol (see **IETF RFC2131**):

| Dimension Name (Portal) | Description | Type: value column | Direction |
| --- | --- | --- | --- |
| DHCP OP | Message op code / message type: 1 = BOOTREQUEST, 2 = BOOTREPLY. | int UDR | Non-directional |
| DHCP Message Type | The type of the DHCP message, e.g., DHCPDISCOVER, DHCPOFFER (see **DHCP Message Type**). | int UDR | Non-directional |
| DHCP CI Address | A client IP address (`ciaddr`) that has already been allocated and accepted; only filled in if client is in BOUND, RENEW or REBINDING state and can respond to ARP requests. | string UDR | Non-directional |
| DHCP YI Address | The IP address of the client (`yiaddr`) as allocated by the server and accepted by the client. | string UDR | Non-directional |
| DHCP SI Address | The IP address of next server to use in bootstrap (`siaddr`). | string UDR | Non-directional |
| DHCP Lease | In a client request (DHCPDISCOVER or DHCPREQUEST), the requested lease time for the IP address; in a server reply, the lease time offered by the server (see **IP Address Lease Time**). | int UDR | Non-directional |
| DHCP CH Address | The client hardware address (`chaddr`). | string UDR | Non-directional |
| DHCP Hostname | The name of the client (see **Host Name Option**). | string UDR | Non-directional |
| DHCP Domain | The domain name that client should use when resolving hostnames via the Domain Name System (see **Domain Name Option**). | string UDR | Non-directional |

#### Radius Dimensions

Dimensions related to RADIUS (see **FreeRADIUS attributes**):

| Dimension Name (Portal) | Description | Type: value column | Direction |
| --- | --- | --- | --- |
| Radius Code | The RADIUS Packet type: Access-Request, Access-Accept, Access-Reject, or Access-Challenge (see **IETF RFC2865**). | int UDR | Non-directional |
| Radius User Name | The name of the user to be authenticated. | string UDR | Non-directional |
| Radius Service Type | The type of service the user has requested, or the type of service to be provided. | int UDR | Non-directional |
| Radius Framed IP Address | The address to be configured for the user. | string UDR | Non-directional |
| Radius Framed IP Mask | The IP netmask to be configured for the user when the user is a router to a network. | string UDR | Non-directional |
| Radius Framed Protocol | The framing to be used for framed access. | string UDR | Non-directional |
| Radius Accounting Status | Indicates whether this Accounting-Request marks the beginning of the user service (Start) or the end (Stop). | int UDR | Non-directional |
| Radius Accounting Session ID | A unique Accounting ID that enables the matching of start and stop records in a log file. | string UDR | Non-directional |

### Legacy Application Decodes

The dimensions in this category correspond to application-layer data captured by older host agent versions. While still supported for historical data, we recommend using the modern **Application Decodes** for all new queries and dashboards.

> **Note:** The dimensions below require Kentik's software host agent (see **About the Universal Agent**).

#### Legacy DNS Dimensions

These dimensions represent DNS query and response data stored in Native KDE columns (see **Host Traffic Dimensions**):

| Dimension Name (Portal) | Description | Type: value column | Direction KDE name(s) |
| --- | --- | --- | --- |
| DNS Query | Query from a DNS resolver to a DNS name server.  **Note:** Superseded by DNS Query Name. | string Native | Src/Dst: `kflow_dns_query`, N.A. |
| DNS Query Type | The resource record type requested by the DNS query. | bigint Native | Src/Dst: `kflow_dns_query_type`, N.A. |
| DNS Return Code | DNS return code (see **IANA DNS Parameters**)  **Note:** Superseded by DNS Reply Code. | bigint Native | Src/Dst: `kflow_dns_ret_code`, N.A. |
| DNS Response | The response from a DNS server to a DNS query.  **Note:** Superseded by DNS Reply Data. | string Native | Src/Dst: `kflow_dns_response`, N.A. |

#### Legacy HTTP Dimensions

These dimensions capture basic web traffic metadata using legacy indexing (see **Host Traffic Dimensions**):

| Dimension Name (Portal) | Description | Type:  value  column | Direction KDE name(s) |
| --- | --- | --- | --- |
| HTTP URL | Filename portion of path, with query string (if any). | string Native | Src/Dst: N.A., `kflow_http_url` |
| HTTP Host Header | Domain name of the server.  **Note:** Superseded by HTTP Host. | string Native | Src/Dst: N.A., `kflow_http_host` |
| HTTP Return Code | HTTP status code.  **Note:** Superseded by HTTP Status. | bigint Native | Src/Dst: N.A., `kflow_http_status` |
| HTTP Referrer | The address from which a destination webpage is requested. | string Native | Src/Dst: N.A., `kflow_http_referer` |
| HTTP User Agent | User agent information identifying the client that submitted a request. | string Native | Src/Dst: N.A., `kflow_http_ua` |

## Container Networking Dimensions

Kentik dimensions for containerized environments allow you to bridge the gap between traditional network flow and logical orchestration metadata.

> **Note:** Use of Kubernetes with Kentik requires a special software agent; for further information contact Kentik (see **Customer Care**).

### Kubernetes Dimensions

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/logo-kubernetes-043h255w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A22Z&se=2026-05-12T09%3A53%3A22Z&sr=c&sp=r&sig=1FzhJ3fIPXVuSt1NVjpDhbXVmNAAqQ6kV3GJJHHgrEY%3D)These dimensions represent the logical structure of your Kubernetes cluster. For more technical background, see **What is Kubernetes?**.

| Dimension Name (portal) | Description | Type: value column | Direction |
| --- | --- | --- | --- |
| Pod Name | The name of a pod, which represents a set of running containers on your cluster. | string | Src/Dst |
| Pod Namespace | The scope within which the pod name is valid and unique. | string | Src/Dst |
| Workload Name | The name of a workload, which is a system of services or applications that can run to fulfill a task or carry out a business process. | string | Src/Dst |
| Workload Namespace | The scope within which the workload name is valid and unique. | string | Src/Dst |
| Container Name | The name of an executable image that contains software and all of its dependencies. | string | Dst only |
| Service Name | The name of a network application that is running as one or more pods in your cluster. | string | Src/Dst |
| Service Namespace | The namespace in which the service is running. | string | Src/Dst |

### PATA Dimensions

The **Process-Aware Telemetry Agent (PATA)** provides granular visibility into the specific system processes generating network traffic. This is vital for security auditing and identifying "noisy" applications.

> **TIP**: PATA dimensions require the kappa host agent. If you are seeing "N/A" in these fields, verify that the agent is correctly deployed on your nodes, or contact Kentik support for assistance.

| Dimension Name (Portal) | Description | Type | Direction |
| --- | --- | --- | --- |
| Process PID | The ID of the source or destination process of the flow. | integer | Src & Dst |
| Process Name | The name of the source or destination process of the flow | string | Src & Dst |
| Process Cmdline | The command entered at the CLI, related to the process ID of the source or destination process of the flow. | string | Src & Dst |
| Process Container ID | The string (guid format) that uniquely identifies the container. | string | Src & Dst |
| Node | The source or destination Kubernetes node that originated or terminated the flow. | string | Src & Dst |
| Object Name | The Kubernetes object (pod or service name) for the source or destination pod in the traffic flow. | string | Src & Dst |
| Object Namespace | The Kubernetes namespace of the source or destination pod in the traffic flow. | string | Src & Dst |
| Object Type | The Kubernetes object type for the traffic flow (pod or service). | string | Src & Dst |
| Container Name | The container name(s) related to the source or destination process ID. | string | Src & Dst |
| Workload Name | The name of the workload that a pod or service was deployed as. | string | Src & Dst |
| Workload Namespace | The namespace name of a source or destination traffic flow. | string | Src & Dst |
| Object Labels | The object labels associated with source or destination traffic. | string | Src & Dst |
| Cluster ID | A unique identifying integer that the cluster assigns to itself. | integer | Other |
