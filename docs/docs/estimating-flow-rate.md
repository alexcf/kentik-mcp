# Estimating Flow Rate

Source: https://kb.kentik.com/docs/estimating-flow-rate

---

This article covers how to measure and calculate the baseline Flows Per Second (FPS) on your network devices, and explains the relationship between FPS and the pricing of flow-related **Kentik Plans**.

> **TIP**: Follow this guide and the **Flow Overview** (which has Kentik’s sampling rate recommendations) before discussing with Kentik support about a **Kentik Plan** that best fits your needs.

## About FPS Calculation

In networking terms, a “flow” is a unidirectional set of packets with shared metadata (e.g., source and destination IP, ports, and protocol) traversing your network devices (routers, switches, or hosts). Flow protocols for **physical infrastructure** (e.g. sFlow, IPFIX**,** NetFlow) and **cloud infrastructure** (AWS VPC logs) are ways to collect a **user-configured sample** of that metadata and send it to Kentik (see **Flow Sampling**).

The sampling rate in **Flows Per Second (FPS)** at which flow is sent to Kentik is the primary driver of plan capacity/cost, so an accurate FPS estimate helps as follows:

* **Right-Sized Deployment**: Avoid both overpayment for unused capacity and unexpected budget overages from license spikes.
* **Ensure Operational Performance**: Guarantee that Kentik can process your data in as close to real-time as possible, maintaining the integrity and accuracy of your network analytics.

### Background: Plans

Every device sending flow data to Kentik is assigned to a **Kentik Plan**, helping to right-size the resources required to process and store your network telemetry in the **Kentik Data Engine** (KDE).

When choosing a plan, consider how these variables affect your capacity:

| Attribute | Description |
| --- | --- |
| Flows Per Second | The provisioned maximum bandwidth. measured differently by protocol. |
| Device Quantity | The total number of exporting devices allowed per plan. |
| Device Types | Supports various hardware like routers, switches, and hosts (see **Supported Device Types**). |
| BGP Integration | Inclusion of BGP routing tables for advanced peering analytics. |
| Data Retention | The number of days the data is stored at both full and optimized resolutions. |

> **IMPORTANT:**
>
> * FPS represents **flow records** for NetFlow/IPFIX (aggregated metadata) but represents **sampled packets** (datagrams) for sFlow.
> * AWS VPC flow logs are **one flow per record**, but Azure VNet flow logs **may be split into two** (given that sent and received are represented in a single line).

## Before You Begin

Before starting your FPS calculations, gather the following information for each device that will export flow to Kentik:

* **Device Inventory**: A complete list of all routers, switches, or hosts intended for monitoring.
* **Configuration Status**: Ensure flow data export is already enabled (see **Device Configs Directory**).
* **Inbound Traffic Volume**: The approximate Gb/s of **ingress-only** traffic for each device. Measuring both directions will result in an inaccurate FPS estimate in **Step 3**.
* **Target Sampling Rate**: For comparison purposes, determine the device’s ideal sampling rate based on its role (e.g., Edge vs. Data Center) using the **Recommended Sampling Rates** table. The sampling rate represents a different unit of measure depending on flow protocol:

  + **NetFlow/IPFIX**: The ratio of total flows to exported records to Kentik.
  + **sFlow**: The ratio of total raw packets to sampled datagrams sent to Kentik.

> **Notes**:
>
> * For NetFlow v9 and IPFIX, ensure your device is configured to send **Option Templates**. These are necessary for Kentik to decode metadata like SNMP interface names and VLAN IDs that aren't in the standard data stream.
> * Your sampling rate is the most effective way to balance the depth of your visibility with your Kentik Plan capacity.

## FPS Calculations

To measure your current FPS levels and estimate your requirements for peak events and future growth, follow these steps.

> **Notes:**
>
> * To avoid double-counting, only measure traffic (Gb/s) at either ingress or egress, never both.
> * You can calculate a single "worst-case" device to extrapolate your needs or repeat these steps for each router for a precise total.

### **Step 1: Capture Real-Time Flow Data**

To calculate your current flow rate, log into your router's CLI and capture two flow-count measurements exactly **10 seconds apart**.

1. **Run the Export Command**

   Log into your router’s CLI and run the command corresponding to your manufacturer. Locate the **flow count** within the resulting report (see **Sample Code for Common Devices**).

   | Manufacturer | OS | Command to Export Flow Count |
   | --- | --- | --- |
   | **Cisco** | iOS/NX-OS | `show ip flow export` or `show flow export` |
   | **Cisco** | iOS XR | `show flow exporter [name] location [loc]` |
   | **Juniper** | EX, MX, T Series | `show services accounting flow` |
   | **Arista/Brocade/Foundry** | Arista EOS/Network OS | `show sflow`  **Note**: For Arista IPFIX, use `show flow tracking hardware counter`. |
2. **Record Two Data Points**

   Capture two consecutive measurements to establish a baseline. For the most accurate estimate, perform these steps during your **peak traffic hours**:

   1. **First Measurement**: Run the command and record the current flow count (F1).
   2. **Second Measurement**: After **10 seconds**, run the command again and record the new flow count (F2).
   > **What value should I look for?**
   >
   > * **NetFlow/IPFIX:** Look for `Flows exported` or `Number of Flow Records Exported`.
   > * **sFlow:** Look for `Number of Samples` or `sFlow samples collected`.

### Step 2: Calculate Your Baseline FPS

Use the difference between your two captured flow counts to determine your baseline **Flows Per Second (FPS)**.

* **F2**: The value from your second measurement.
* **F1**: The value from your first measurement.
* **10**: The exact number of seconds between your two captures.

**Example**: If your F1 was 50,000 and increased to an F2 of 61,000 (an increase of 11,000 flows over 10 seconds), your baseline is 1,100 FPS.

### **Step 3: Forecast for Peak and Growth**

If your baseline measurements weren't taken during a high-traffic period, use your known traffic volume (**Gb/s**) to estimate your requirements for peak events and future growth. Calculate your peak FPS by multiplying your baseline by the ratio of peak-to-baseline traffic:

#### Proportional Growth Table

Use this table to map out your expected capacity needs:

| Metric | Baseline Measurement | Peak-Time Estimate | 6-Month Forecast |
| --- | --- | --- | --- |
| **On-Prem Traffic** | 5.5 Gbps | 9.5 Gbps | 20 Gbps |
| **On-Prem FPS** | 1,100 FPS | 1,900 FPS | 4,000 FPS |
| **Cloud Egress** | 2 Gbps | 4 Gbps | 10 Gbps |
| **Cloud FPS (VPC)** | 200 FPS | 400 FPS | 1,000 FPS |

## Sample Code For Common Devices

Click to expand an example of how to export flow count from common devices.

### Cisco IOS

Locate the `flows exported` value to use in your **Step 2** calculation.

```
Router# show ip flow export
Flow export v5 is enabled for main cache
  Exporting flows to 10.51.12.4 (9991) 10.1.97.50 (9111)
  Exporting using source IP address 10.1.97.17
  Version 5 flow records
  11 flows exported in 8 udp datagrams
  0 flows failed due to lack of export packet
  0 export packets were sent up to process level
  0 export packets were dropped due to no fib
  0 export packets were dropped due to adjacency issues
  0 export packets were dropped due to fragmentation failures
  0 export packets were dropped due to encapsulation fixup failures
  0 export packets were dropped enqueuing for the RP
  0 export packets were dropped due to IPC rate limiting
  0 export packets were dropped due to output drops
```

### Cisco IOS XR

For IOS XR, you can use the `1 second` flow rate directly, or calculate the baseline using the `Flows exported` value if more precision is needed over a 10-second window.

```
RP/0/RP0/CPU0:router# show flow exporter fem1 location 0/0/CPU0
Flow Exporter: NFC
Used by flow monitors: fmm4
Status: Normal
Transport   UDP
Destination 12.24.39.0    (50001)
Source    12.25.54.3    (5956)
Flows exported:                  0 (0 bytes)
Flows dropped:                  0 (0 bytes)
Templates exported:                1 (88 bytes)
Templates dropped:                0 (0 bytes)
Option data exported:              0 (0 bytes)
Option data dropped:              0 (0 bytes)
Option templates exported:            2 (56 bytes)
Option templates dropped:            0 (0 bytes)
Packets exported:                3 (144 bytes)
Packets dropped:                  0 (0 bytes)
Total export over last interval of:
  1 hour:                      0 pkts
                          0 bytes
                          0 flows
  1 minute:                    3 pkts
                         144 bytes
                          0 flows
  1 second:                    0 pkts
                          0 bytes
                          0 flows
```

### Cisco Nexus NX-OS

Identify the `Number of Flow Records Exported` value under 'Exporter Statistics' to use as your flow count.

```
N7K1# show flow export
Flow exporter KENTIK:
  Description: ships flows to Kentik cloud
  Destination: 10.255.255.100
  VRF: default (1)
  Destination UDP Port 2055
  Source Interface Vlan10 (10.10.10.5)
  Export Version 9
  Exporter Statistics
    Number of Flow Records Exported 726
    Number of Templates Exported 1
    Number of Export Packets Sent 37
    Number of Export Bytes Sent 38712
    Number of Destination Unreachable Events 0
    Number of No Buffer Events 0
    Number of Packets Dropped (No Route to Host) 0
    Number of Packets Dropped (other) 0
    Number of Packets Dropped (LC to RP Error) 0
    Number of Packets Dropped (Output Drops) 0
    Time statistics were last cleared: Tue Jul  8 21:12:06 2014
```

### Juniper EX, MX, and T Series

Locate the `Flows exported` value under the 'Flow information' section for your specific Service Accounting interface.

> **TIP**: If you are using **J-Flow** on newer Junos versions (like the MX series), adding the command `show services ipflow collector statistics` is a helpful alternative if you don't see the "Flows exported" counter in the first command.

```
user@host> show services accounting flow
  Flow information
    Service Accounting interface: sp-2/0/0, Local interface index: 215
    Flow packets: 9867, Flow bytes: 631488
    Flow packets 10-second rate: 0, Flow bytes 10-second rate: 628
    Active flows: 2, Total flows: 10
    Flows exported: 4028, Flows packets exported: 6150
    Flows inactive timed out: 8, Flows active timed out: 4026
    Service Accounting interface: sp-2/1/0, Local interface index: 223
    Flow packets: 0, Flow bytes: 0
    Flow packets 10-second rate: 0, Flow bytes 10-second rate: 0
    Active flows: 0, Total flows: 0
    Flows exported: 0, Flows packets exported: 1
    Flows inactive timed out: 0, Flows active timed out: 0
```

### Arista

Arista routers use sFlow; identify the `Number of Samples` below to use as your flow count for the **Step 2** calculation.

> **Note**: For Arista IPFIX, use `show flow tracking hardware counter`.

```
vEOS1__20:41:53#show sflow
sFlow Configuration
-------------------
Destination(s):
  172.16.0.1:6343 ( VRF: default)   <-- should be pointing at the external sFlow collector
Source(s):
  172.16.0.41 ( VRF: default)
  :: ( default) ( VRF: default)
Sample Rate: 16384            
Polling Interval (sec): 30.0    
Rewrite DSCP value: No
Status
------
Running: Yes
Polling On: Yes                  
Sampling On: Yes ( default)
Send Datagrams: Yes ( VRF: default)
Statistics
----------
Total Packets: 487924
Number of Samples: 847
Sample Pool: 0
Hardware Trigger: 0
Number of Datagrams: 27
```

### Brocade and Foundry

Brocade/Foundry routers utilize sFlow; identify the `sFlow samples collected` value at the bottom of the report to determine your baseline flow rate.

```
device#show sflow
sFlow version:5
sFlow services are enabled.
sFlow agent IP address: 10.123.123.1
sFlow source IP address: 5.5.5.5
sFlow source IPv6 address: 4545::2
4 collector destinations configured:
Collector IP 192.168.4.204, UDP 6343
Configured UDP source port: 33333
Polling interval is 0 seconds.
Configured default sampling rate: 1 per 512 packets
Actual default sampling rate: 1 per 512 packets
Sample mode: Non-dropped packets
The maximum sFlow sample size:512
exporting cpu-traffic is enabled
exporting cpu-traffic sample rate:16
exporting system-info is enabled
exporting system-info polling interval:20 seconds
10552 UDP packets exported
24127 sFlow samples collected.
sFlow ports: ethe 1/2 to 1/12 ethe 1/15 ethe 1/25 to 1/26 ethe 4/1 ethe 5/10 to
5/20 ethe 8/1 ethe 8/4
Module Sampling Rates
---------------------
Slot  1 configured rate=512, actual rate=512
Slot  3 configured rate=0, actual rate=0
Slot  4 configured rate=10000, actual rate=32768
Slot  5 configured rate=512, actual rate=512
Slot  7 configured rate=0, actual rate=0
Slot  8 configured rate=512, actual rate=512
Port Sampling Rates
-------------------
Port 8/4, configured rate=512, actual rate=512, Subsampling factor=1
Port 8/1, configured rate=512, actual rate=512, Subsampling factor=1
```

## **Summary and Next Steps**

With your current and forecasted FPS calculated, you’re ready to review your configuration and Kentik Plan capacity:

* **Review Configuration**: Check the sampling rate settings for your devices in the Kentik portal or cloud provider console (see **Device Settings** and **Public Clouds**) and make updates as necessary.
* **Verify Plan Capacity**: Navigate to the **Licenses** page in the Kentik portal (under Organization Settings) to compare your new FPS estimates against your current plan limits. Pay attention to any downsampling indicators (yellow/red warnings).
* **Refine and Optimize**: If your FPS exceeds your plan or your visibility needs change, contact Kentik support for assistance in balancing data depth with cost (see **Customer Care**).

FPS stands for flows per second, which is the rate at which flow metadata is sent to Kentik from a given customer. It directly correlates with the resources required to provide that customer with the Kentik service.

The process of reducing the number of flow records by sampling, where only one flow record is exported for every X flows.
