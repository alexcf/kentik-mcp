# Kentik NMS Configuration

Source: https://kb.kentik.com/docs/kentik-nms-configuration

---

This article covers the configuration of Kentik’s Universal Agent for Kentik NMS.

![Deployment options for the Universal Agent, highlighting SNMP and Syslog capabilities.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/UA-deploy-agent-dialog-capability-selection.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A37Z&se=2026-05-12T09%3A32%3A37Z&sr=c&sp=r&sig=XPzIzw%2FVHeT4z1iAbsEYumAZ%2B0q22pHbIZNj581sk34%3D)

> **TIP:** See the **Network Monitoring**category of articles for more information on Kentik NMS.

Configuring Kentik NMS requires Kentik's **Universal Agent**, a software agent that collects data from monitored devices on your network and delivers it to the Kentik platform. In this guide, you’ll learn how to install the agent on a host machine in your network and enable the **SNMP/ST** capability.

|  |  |
| --- | --- |
| **Purpose:** | Discover and collect NMS device data using Kentik's Universal Agent |
| **Benefits:** | Efficient installation and deployment through Docker or package managers |
| **Use Cases:** | Enable network analysis with data collection via SNMP and Streaming Telemetry |
| **Relevant Roles:** | Network Engineer |

> **Notes:**
>
> * SNMP/ST is one of several **Agent Capabilities** that can be enabled on Kentik’s Universal Agent.
> * For more on using Streaming Telemetry for Kentik NMS (see **NMS via Streaming Telemetry****)**.
> * Kentik's Universal Agent will eventually replace all Kentik agents (see **Agents**).

## SNMP/ST Capability Requirements

Requirements for using the **SNMP/ST** capability of Kentik’s Universal Agent are as follows.

#### Hardware Requirements

Hardware requirements for the host machine running the Universal Agent vary based on the number of devices from which metrics will be polled.

**Recommendation**: Conservatively, allocate 1 CPU core and 4 GB RAM for every 250 devices.

> **Note:** For optimal hardware allocation, be sure to:
>
> * Test in your specific environment.
> * Factor in competing resources (e.g., databases) on the host machine running the Universal Agent.

#### Connectivity Requirements

Deployed Universal Agent instances need to be able to initiate traffic to the following:

* `grpc.api.kentik.com` (or `kentik.eu` for Europe) using TLS on TCP 443
* Your monitored network devices using SNMP on UDP 161
* Your monitored network devices using ICMP (Echo)

## Universal Agent Installation

The Universal Agent is deployed via Docker or Linux on a host machine on your network, as initiated from one of the following Kentik portal locations:

* **NMS Devices** **page** (Network Monitoring System » **Devices**):

  + Click **Add Devices**.
  + Choose **Full monitoring** to open the **NMS Setup Wizard**, where you can either:

    - Select an existing agent from the Select an Agent list.
    - Deploy a new agent using the commands provided on the **Docker** and **Linux** tabs (see **Deploy the Universal Agent**).
    > **TIP:** This method is recommended when the agent will be used only for Kentik NMS, because the SNMP/ST capability is configured automatically.
* **Universal Agents** **page** (Settings » **Universal Agents**): ![Select capabilities for deploying the Universal Agent, including SNMP/ST option.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/UA-deploy-agent-dialog-snmp-st-capability.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A37Z&se=2026-05-12T09%3A32%3A37Z&sr=c&sp=r&sig=XPzIzw%2FVHeT4z1iAbsEYumAZ%2B0q22pHbIZNj581sk34%3D)

  + Click **Deploy Agent** at the top right, and then either:

    - Authorize an existing agent from the Select an Agent list.
    - Deploy a new agent using the commands provided on the **Docker** and **Linux** tabs, ensuring the **SNMP/ST** capability is selected.
    > **TIP:** This method is recommended when deploying and managing multiple Universal Agents with a variety of capabilities (not exclusively for Kentik NMS).

### Agent Authentication

The Universal Agent uses the same authentication as for other Kentik agents, with two custom headers sent by HTTP over TLS:

* **X-CH-Auth-Email**: The Kentik user’s email address (found in the **User** menu on the portal’s main navbar).
* **X-CH-Auth-API-Token**: The Kentik user’s API token (found in the **User** **Information** pane of the portal's **User Profile**page).

After installation, the agent starts up without a token and announces itself to the Kentik platform in a “pending-authorization” state. Admin-level users can view/authorize agents in this state from the **Universal Agents** page (see **NMS Download and Install**).

Once authorized, the new agent is provided with an email and token that it stores in `/opt/kentik/components/kagent/current/conf/instance-auth.yaml` on your host machine and uses for all future communications with the Kentik platform.

## SNMP/ST Capability Configuration

The **SNMP/ST** capability of the Universal Agent is configured upon deployment to:

* Discover devices within a given IP range for Kentik to monitor, as set in the **NMS Setup Wizard**.
* Choose which discovered devices to monitor (see **NMS Discovery**).

> **Note:** NMS data collection via Streaming Telemetry is set up after setting up monitoring via SNMP. For details see **NMS via Streaming Telemetry**.

#### Migrate Devices to New Agent

After deploying a new agent, you can migrate some or all devices from an existing agent to the new one. To do so, no action is required for the old agent. Instead, perform the following steps:

1. Deploy the new agent.
2. Run a regular discovery (see **NMS Discoveries Page**) of the target IP range with the new agent.
3. Check the checkbox of each discovered device that you'd like to monitor with the new agent.
4. Click **Add Devices** to migrate the selected devices to the new agent.

## Custom Device Profiles

Kentik uses device profiles to normalize metrics and metadata from a variety of common device vendors and models. You can also customize profiles for specific situations, usually one of the following:

* Your devices aren't supported.
* Your devices are supported, but you require capturing non-standard metrics for that make/model.

  > **Note:** For a detailed walkthrough of configuration for this use case, see our blog post **How to Configure Kentik NMS to Collect Custom SNMP Metrics**.

### Generating Custom Profiles

Custom profiles are based on the output of a common Linux utility called `snmpwalk`, and Kentik supports these two approaches to using it:

* **Remote SNMP walk** (**recommended**): Allow Kentik to perform remote SNMP walks, then request one for a specific device. See **Remote SNMP Walk**.
* **Manual SNMP walk**: Run an SNMP walk yourself on the device, then pass Kentik the information. See **Manual SNMP Walk**.

  > **Note:** Use this approach only if your organization would prefer not to grant Kentik permission to perform remote SNMP walks.

#### Remote SNMP Walk

To enable a custom device profile:

1. **Enable remote SNMP walks**: On the NMS Devices page, choose **Remote SNMP Walks** from the Actions menu (see **NMS Devices Actions**), then you can grant Kentik permission to perform `snmpwalk` operations remotely.

   > **Note:** Your organization needs to grant this permission only once. To remove the permission, open the dialog and switch it off.
2. **Request a custom profile**: Contact Kentik (see **Customer Care**) to request a custom profile for a specific device. Kentik will run `snmpwalk` against the device (which must be SNMP-enabled).

#### Manual SNMP Walk

If your organization would prefer not to grant Kentik permission to perform remote SNMP walks, you can run `snmpwalk` internally against any device for which you'd like a custom profile (the device must be SNMP-enabled). You can then provide the output file to Kentik (see **Customer Care**) so we can use it to create a custom profile for the device.  
  
The syntax to run `snmpwalk` depends on the SNMP version enabled on the device:

**SNMP v1 or v2c**:

```
snmpwalk --hexOutputLength=0 -ObentUx -v2c -c community X.X.X.X.1 > device-model.walk
```

**SNMP v3**:

```
snmpwalk --hexOutputLength=0 -ObentUx -v3 -l authPriv -u username -a SHA -A passphrase -x AES -X passphrase X.X.X.X.1 > device-model.walk
```

The above commands include the following arguments:

* `-v`: The specific SNMP version (e.g., 1, 2c, or 3).
* `-c` (versions 1 and 2c only): The SNMP community configured on the device.
* `X.X.X.X`: The IP address of the device you are polling (replace with the actual IP).
* `-A` (v3 only): The authentication protocol passphrase.
* `-X` (v3 only): The privacy protocol passphrase.
* `.1`: Following the IP address with ".1" tells `snmpwalk` to poll all available MIBs on the device (walk the full tree).

  > **Note:** If the device in question carries a full internet routing table, each prefix will show up in multiple MIB trees, in which case you may instead need to walk selected trees and then combine the output. For help, see **Customer Care**.
* `device-model.walk`: The name of the file to which the output from `snmpwalk` will be saved. Specify "device" as the vendor and "model" as the model of the device on which `snmpwalk` was run. For example, if the device is a Cisco ASR9001 then name the file “cisco-asr-9001.walk.”

## NMS Custom Metrics

Kentik NMS enables you to include arbitrary custom metrics in the data that will appear in the portal's **Metrics Explorer** module.

Submit custom metrics via API requests conforming to the **InfluxDB Line Protocol**used by Kentik's Universal Agent.

> **Notes**:
>
> * Client libraries are available for a wide variety of languages (see **InfluxDB docs**)
> * You can also output this format with basic string handling, but be careful with string quoting if you need to handle arbitrary data in string tags.

* **Kentik US cluste**r: `https://grpc.api.kentik.com/kmetrics/v202207/metrics`
* **Kentik EU cluster**: `https://grpc.api.kentik.eu/kmetrics/v202207/metrics`

Authentication uses the two custom headers described in **NMS Authentication**.

> **TIPS:**
>
> * A `User-Agent` header is not required, but putting something descriptive there can help with troubleshooting.
> * Tags or fields sent with a custom measurement do not age out of **Metrics Explorer**; they will stay listed indefinitely.
