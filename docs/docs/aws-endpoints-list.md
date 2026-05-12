# AWS Endpoints List

Source: https://kb.kentik.com/docs/aws-endpoints-list

---

As part of AWS **Metadata** and **Flow/Firewall Log Collection**, Kentik needs permission to access selected endpoints on your behalf, as detailed in the following lists.

### AWS Metadata Endpoints

#### **EC2**

* `DescribeAvailabilityZones`
* `DescribeCustomerGateways`
* `DescribeFlowLogs`
* `DescribeInternetGateways`
* `DescribeInstances`
* `DescribeNatGateways`
* `DescribeNetworkAcls`
* `DescribeNetworkInterfaces`
* `DescribeManagedPrefixLists`
* `DescribePrefixLists`
* `DescribeRouteTables`
* `DescribeSecurityGroups`
* `DescribeSubnets`
* `DescribeTransitGateways`
* `DescribeTransitGatewayAttachments`
* `DescribeTransitGatewayVpcAttachments`
* `DescribeTransitGatewayRouteTables`
* `DescribeTransitGatewayConnects`
* `DescribeTransitGatewayConnectPeers`
* `DescribeVpcs`
* `DescribeVpcEndpoints`
* `DescribeVpcPeeringConnections`
* `DescribeVpnConnections`
* `DescribeVpnGateways`
* `DescribeManagedPrefixLists`
* `DescribeTransitGatewayRouteTables`
* `SearchTransitGatewayRoutes`
* `GetManagedPrefixListEntries`

#### **Direct Connect**

* `DescribeDirectConnectGateways`
* `DescribeVirtualInterfaces`
* `DescribeLags`
* `DescribeConnections`

#### **ELB**

* `DescribeLoadBalancers`

#### **IAM**

* `ListAccountAliases`

#### **Network Manager** (core network metadata)

* `ListCoreNetworks`
* `GetCoreNetwork`
* `GetCoreNetworkPolicy`
* `ListAttachments`
* `GetNetworkRoutes`

#### **Network Firewall**

* `ListFirewalls`
* `DescribeFirewall`
* `ListFirewallPolicies`
* `DescribeFirewallPolicy`
* `DescribeRuleGroup`

### Optional AWS Endpoints

To optionally get a list of accounts in an AWS organization, Kentik may need to access the following additional endpoints:

#### **Organizations**

* `ListAccounts`

#### **CloudWatch**

* `ListMetrics`
* `GetMetricStatistics`
* `GetMetricData`

#### **STS**

* `AssumeRole`
