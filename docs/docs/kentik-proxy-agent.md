# Kentik Proxy Agent

Source: https://kb.kentik.com/docs/kentik-proxy-agent

---

This article discusses Kentik's NetFlow proxy agent, `kproxy`, which encrypts local flow from your organization's routers and switches.

> **Note****:** Because `kproxy` agents do not de-duplicate flow, if a given device sends flow to more than one instance of `kproxy`, the flows from that device will be overrepresented in the flow records of the Kentik Data Engine (KDE).

## Proxy Agent Overview

Kentik provides a free NetFlow proxy agent called `kproxy` that enables data collected from network devices to be sent securely to Kentik. The machine running `kproxy` isn't actually handling traffic directly but instead collects flow records (NetFlow v5/v9, IPFIX, and sFlow) and SNMP and encrypts it locally before forwarding it to Kentik. A single instance of the `kproxy` executable can redirect flow for multiple routers and switches. Multiple servers across the network can run `kproxy` to distribute traffic and load.  
  
In addition to the collection and encryption of device data, `kproxy` also performs rate limiting and resampling to keep maximum flows per second (FPS) within applicable license limits (see **About Licenses**).  
  
The steps involved in downloading, installation, and configuration of `kproxy` are described in the following topics below:

* The step-by-step routing of flow and SNMP to `kproxy` is covered in **kproxy Setup**.
* Downloading `kproxy` from Kentik is described in **kproxy Download and Install**.
* Configuration of `kproxy` for flow encryption is described in **kproxy Command Line**.

> **Notes*:***
>
> * sFlow data received by `kproxy` is parsed to ensure that only fields relevant to flow analytics are forwarded to the Kentik ingest layer. `kproxy` does not forward the entire sFlow packet to Kentik.
> * For assistance with any aspect of the setup process, see **Customer Care**.

### Kentik Traffic Flows

The diagram below shows an overview of traffic flows between the Kentik NetFlow proxy agent, installed in a customer backbone or IT facility, and Kentik:

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/Proxy_data_flows-674w(1).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A36Z&se=2026-05-12T09%3A50%3A36Z&sr=c&sp=r&sig=8785vVCNtvxEbB5AljU%2BKKKgTdQva4pUmm%2Ba5jf58F8%3D)

### Connections and Behavior

The following points describe agent behavior and the connection between customer devices, the Kentik proxy agent, and Kentik:

* `kproxy` transports all traffic received from customer devices securely to Kentik.
* The Kentik Data Engine and the Kentik portal handle data forwarded via `kproxy` identically to data received directly from customer devices.
* `kproxy` connects to any customer devices sending NetFlow telemetry to Kentik, and also to Kentik itself to send the encrypted data as well as to receive configuration information.
* In order to send traffic to Kentik, `kproxy` will build, for each NetFlow-sending device, two HTTPS sessions, one for flow traffic and the other for SNMP. All traffic is sent to Kentik using such HTTPS sessions, with a Kentik flow ingest server certificate.
* SNMP is converted into JSON format and NetFlow/sFlow/IPFIX is converted to a Kentik proprietary binary format.

### Example Proxy Deployment

The following diagram illustrates typical deployment of the Kentik proxy agent.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/Proxy_deployment-674w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A36Z&se=2026-05-12T09%3A50%3A36Z&sr=c&sp=r&sig=8785vVCNtvxEbB5AljU%2BKKKgTdQva4pUmm%2Ba5jf58F8%3D)

## `kproxy` Requirements

The following resources must be available (at minimum) to support the use of `kproxy` running on a VM:

* RAM allocation of 1GB per 1000 flow records per second (FPS).

  > **Note*:*** Actual requirements vary depending on how “bursty” the flow is.
* One CPU core per 5000 flow records per second (before rate-limiting).

When `kproxy` is used for non-flow data, the following equivalents may be applied to the above requirements:

* **Streaming Telemetry**: 1 interface metric = 1 flow record.

  > **Note*:*** Using Streaming Telemetry with `kproxy` is distinct from using Streaming Telemetry for Kentik NMS (see **NMS via Streaming Telemetry**).

> **Notes*:***
>
> * `kproxy` must be deployed on a separate server from any host agents sending flow to Kentik via `kproxy`.
> * When running `kproxy` on a VM, the CPU must be dedicated to the instance.

#### Connectivity Requirements

`kproxy` will listen for inbound UDP NetFlow/sflow on its configured port (default 9995) and relay this flow securely to flow.kentik.com in TLS-encrypted format. Upon receipt of flow from a configured device, the proxy will attempt (by default) to connect back to the IP address sending flow via SNMP to enumerate basic device information such as interface metadata and CPU.

| From | To | Protocol-Port | NAT Support |
| --- | --- | --- | --- |
| `kproxy` | flow.kentik.com | tcp-443 | Y |
| `kproxy` | monitored devices | udp-161  (optional but recommended) | Y |
| monitored devices | `kproxy` | udp-9995  (UDP port for receiving flow at the proxy) | N |

> **Note*:*** The `kproxy` UDP port is 9995 unless specified with the command line argument -port (see **kproxy CLI Reference**).

### Supported Distributions

`kproxy` v7.48.0 supports the following distributions:

* **Ubuntu**

  + 20.04 focal
  + 22.04 jammy
  + 24.04 noble
* **Debian**

  + 11 bullseye
  + 12 bookworm
* **Red Hat**

  + RHEL8
  + RHEL9

## `kproxy` Setup

When used for encryption, `kproxy` pulls information from the Kentik system to determine which routers it will talk to. The routing of flow and SNMP to `kproxy` is enabled with the following steps:

1. Create a device in the Kentik portal (Devices » Add Device; see **Network Devices**) for each of the routers from which you'd like to send flow:

   1. Set a device's Name and Description.
   2. Set the device type as "Router" (even though you are sending through the agent).
   3. Set the Device IP as the IP of the router that the agent will see the flow being received from. You may enter multiple IPs, comma separated, if there is the possibility that flow may source from multiple IPs for the said router. Private IPs are acceptable.
   4. Set Device SNMP IP to the router IP that the agent will poll for SNMP.
   5. Set Flow Type to the type of flow that the router is configured to export from.
   6. Set Sample Rate to the rate at which the router is set to sample.
   7. Save (Add) the device.
2. Check the system clock and time zone settings on the host server where `kproxy` will be deployed, which must be accurate to within a minute for `kproxy` to function correctly. Kentik recommends that hosts running `kproxy` use Network Time Protocol (NTP).
3. Determine the IP that `kproxy` will bind to on the host in order to receive flow.
4. Determine the port that the host will accept flow on (i.e., where you will point your routers too).
5. Download and install `kproxy` (see **kproxy Download and Install**).
6. Run `kproxy`, specifying the arguments described in **kproxy Command Line**. You'll likely want to test it initially from the command line and then place it into one of your startup scripts so that it begins on boot. If not placed in the background, the agent will run in the foreground and adhere to standard kill/end signals (e.g., run "nohup kproxy +options &" if running from the command line and exiting shell).
7. Configure your routers and switches to send flow to the `kproxy` instance (see **Router Configuration**).

## `kproxy` Download and Install

Download and installation of `kproxy`, deployed via a Debian/Ubuntu package, an RPM package, a Docker image, or via systemd, is covered in the following topics.

### `kproxy` Docker Install

To install `kproxy` via Docker:

1. In the Kentik portal, go to the **kproxy Agents** page (Settings » Proxy Agents), then click the **Add kproxy Agent** button to open the **Adding a Kproxy Agent** dialog.
2. In the dialog, click the **Docker card**, then copy the value of the `--api_email` and `--api_token` arguments shown in step 2 of the instructions.
3. In terminal, pull down the latest `kproxy` image from Kentik’s Docker hub repository:  
    `docker pull kentik/kproxy:latest`
4. To comply with security best practices, create an environment variable named `KENTIK_API_TOKEN` to store the authentication token, which is the `--api_token` value that you copied in step 2 above.
5. Run the `kproxy` docker image (use the `--api_email` value from step 2 above):  
    `docker run --name kproxy -d --net=host kentik/kproxy --api_email api-email-value@kentik.com`

> **Notes*:***
>
> * Storing the token as an environment variable prevents it from being exposed via a process (ps) command. If you prefer not to use an environment variable for the token, include an `--api_token` argument in the docker run command shown in step 5 above, using the value from step 2 above: `--api_token api-token-value`*\*
> * For additional options, including arguments for specifying a different IP address and/or port, see **kproxy CLI Arguments**.

### `kproxy` Package Install

Linux packages for `kproxy` are available for both Debian/Ubuntu and RPM (CentOS/RHEL) installations. To download and install:

1. In the Kentik portal, go to the **kproxy Agents** page (Settings » **Proxy Agents**), then click the **Add kproxy Agent** button to open the Adding a Kproxy Agentdialog.
2. In the dialog, click the **Debian/Ubuntu Package or RPM Package** card, then copy the value of the `--api_token` and `--api_email` argument shown in step 3 of the instructions.
3. In terminal, download and install the `kproxy` repository:

   1. **Debian/Ubuntu:**  
      `curl -s https://packagecloud.io/install/repositories/kentik/kproxy/script.deb.sh | sudo bash  
       sudo apt-get install kproxy`
   2. **RPM:**  
      `curl -s https://packagecloud.io/install/repositories/kentik/kproxy/script.rpm.sh | sudo bash  
      yum install kproxy`

      > **Note*:*** If the `kproxy` package is not found, your operating system may be unsupported. To try a manual installation, select a version from the **repository**.
4. To comply with security best practices, create an environment variable named `KENTIK_API_TOKEN` to store the authentication token, which is the `--api_token` value that you copied in step 2 above.
5. Run `kproxy` (use the `--api_email` value from step 2 above):  
    `kproxy  --api_email api-email-value@kentik.com`

> **Notes*:***
>
> * Storing the token as an environment variable prevents it from being exposed via a process (ps) command. If you prefer not to use an environment variable for the token, include an `--api_token` argument in the docker run command shown in step 5 above, using the value from step 2 above: `--api_token api-token-value`
> * To run `kproxy` via systemd, refer either to **Deploying kproxy via systemd** or to the Debian/Ubuntu or RPM instructions in the **Add kproxy Agent**dialog on the Kentik portal's **kproxy Agents** page.
> * For additional options, including arguments for specifying a different IP address and/or port, see **kproxy CLI Arguments**.
> * To upgrade a package to the current version of `kproxy`, see **Upgrading an Existing kproxy**.
> * Once `kproxy` has been downloaded and installed, continue with the steps outlined in **kproxy Setup**.

#### Upgrading an Existing `kproxy`

Once installed, a `kproxy` package must be kept up to date to ensure correct performance. For best results, Kentik recommends using your OS's package manager to upgrade existing instances of `kproxy` automatically. If using the package manager is not possible, then you can follow the steps below to perform a manual upgrade.  
  
 1. Check the version of the existing `kproxy` instance.

* **Debian/Ubuntu:**

```
# Command
dpkg -l | grep kproxy

# Response

ii  kproxy-latest-ubuntu-16.04 2.3 amd64 no description given
```

* **Red Hat/Centos/Fedora:**

```
# Command

rpm -qa | grep kproxy

# Response

kproxy-latest-rhel_7-1.0-1.x86_64

# Use that package name in this command:

rpm -qi kproxy-latest-rhel_7-1.0-1.x86_64

# Response

Name : kproxy-latest-rhel_7

Version : 3.3
```

2. Kill the running `kproxy` process (all OS versions; placeholders highlighted).

```
# Command

ps -ef | grep kproxy

# Response

root
```

```
  9979  9895  0 13:59 pts/0  00:00:00 /usr/bin/kproxy -api_pass password_string -api_email username@domain.suffix

# Command

sudo killall kproxy
```

3. Remove the existing `kproxy` package.

* **Debian/Ubuntu**

```
# Command

dpkg -l | grep kproxy

# Response
ii  kproxy-latest-ubuntu-16.04 2.3 amd64 no description given

# Command

sudo dpkg -r kproxy-latest-ubuntu-16.04

# Response

(Reading database... 60028 files and directories currently installed.)

Removing kproxy-latest-ubuntu-16.04 (2.3)
```

* **Red Hat/Centos/Fedora**

```
# Command

rpm -qa | grep kproxy

# Response

kproxy-latest-rhel_7-1.0-1.x86_64

# Command

sudo rpm -e kproxy-latest-rhel_7-1.0-1.x86_64
```

4. Download and install the latest version of `kproxy` as described in **kproxy Download and Install****.**  
 5. Check the version of the newly installed `kproxy` instance and compare it with the version from step 1 to confirm successful installation.  
 6. Restart `kproxy` with the same command options returned in step 2 in response to the `ps` command. See step 6 of **kproxy Setup**.

### Deploying `kproxy` via systemd

To enable faster, more efficient setup of `kproxy`, especially when you need to deploy many instances, the `kproxy` installation process includes a `kproxy` service definition that can be run with Linux's systemd system and service manager:

* To use `kproxy` for flow encryption, run the default version of the service definition (replace placeholders with actual authentication credentials).
* To take advantage of other `kproxy` features, also edit the service definition to specify the desired `kproxy` behavior using the arguments covered in **kproxy CLI Arguments**.

#### Running `kproxy` from systemd

The `kproxy` install process results in installation of two systemd-related files:

* `/lib/systemd/system/kproxy.service`
* `/etc/default/kentik.env.sample`

To run the service definition:

1. Copy the config file at `/etc/default/kentik.env.sample to /etc/default/kentik.env`
2. In the new `/etc/default/kentik.env` service definition file, replace the following highlighted placeholders with actual values associated with a Kentik-registered user:  
    - `KENTIK_API_TOKEN=XXX_FILL_ME_IN:`  The API token from the user's **Authentication Settings**.  
    - `KENTIK_API_EMAIL=foo@example.com`:  The user's email address.
3. To use this `kproxy` instance for tasks other than flow encryption, you can make additional edits in the service definition file, specifying the desired behavior using the arguments in **kproxy CLI Arguments**.
4. Choose the Linux command corresponding to how you want `kproxy` to start:  
    - To start manually, use `systemctl start kproxy`  
    - To start upon boot, use `systemctl enable kproxy`
5. To get logs (confirm that `kproxy` is running) and show the service state of the proxy, run the following Linux commands:  
   `journalctl -u kproxy  
   systemctl status kproxy -l`

## `kproxy` Command Line

The use of the `kproxy` command line is covered in the following topics.

### `kproxy` CLI Arguments

The following topics cover the most common command line arguments for `kproxy`. For a brief reference to additional arguments, see **kproxy CLI Reference**.

> **Notes*:***
>
> * If your organization is registered with Kentik in the EU you must set the `-region` argument to `eu`.
> * If `kproxy` fails to launch, add the `-verbose` flag and try again so that you can provide the output to **support@kentik.com** in order to facilitate troubleshooting.
> * Either " " or "=" can be used between an argument's name and value.
> * Use -h to return a list of arguments.

#### `kproxy` Required Arguments

The following arguments are required to authenticate access to `kproxy`:

* `-api_email` (required): The email address of a registered user as displayed on that user's API System page, which is accessed via the API button for that user on the Users page.
* `-host` (required): Set to one of the following interface IPs:

  + The IP of a single interface for `kproxy` to listen on;
  + `0.0.0.0` to listen on all interfaces.
* `-region` (required for EU customers): If your organization is registered with Kentik in the EU (rather than the US), you must include this argument and set the value to eu.

#### `kproxy` Proxy Agent Arguments

The command line arguments below are used when configuring `kproxy` as a NetFlow proxy agent.

* `-port` (optional): Set the port to listen on. If omitted, `kproxy` defaults to listening on port 9995.
* `-dns` (optional): The address, expressed as IP:Port (e.g. `127.0.0.1:5353`), of an alternate (private) DNS server that is to be used for reverse DNS lookups instead of Kentik's default server. For usage, see **Alternate DNS Lookup**.
* `-site_id` (optional): The ID of a Kentik-registered site (see **Manage Sites**). If a different site ID is specified for two or more instances of `kproxy`, then a device sending flow to Kentik via one such instance will be able to use the same Sending IP address (see **Device General Settings**) as a device sending flow via another instance (see **IP Overloading**).
* `-log_level` (optional): Sets the level of logging. Valid values are `debug`, `info` (default), `warn`, `error`.
* `-proxy_host` (optional): Sets the IP that `kproxy` will use for child processes. The `kprobe` parent process spawns a child process to handle flow from each individual combination of device and time-slice; default is `127.0.0.1`.
* `-healthcheck` (optional): The IP to use for the healthcheck service; default is `127.0.0.1`.
* `-bootstrap_devices` (optional): A comma-separated list of device ids (e.g. "12002,1221") for which this instance of `kproxy` should initiate a child process without first waiting to receive flow. Used in scenarios such as using `kproxy` on a given device for SNMP only (no flow collection).
* `-tee_kproxy` (optional): Reserved for Kentik internal use.
* `-sflow_agent_address` (optional): Set the agent address as the sending IP rather than the original packet source IP address. Use this argument when sFlow may originate from multiple IP addresses on a single device or may traverse a NAT.

If you choose not to store the `api_token` as an environment variable (step 2 above), despite the risk of exposure via a process (`ps`) command, your command line will need to include the token using the following additional parameter:

* `-api_token`: A Kentik-generated string that `kproxy` will use to authenticate a registered user (must be the same user as for `-api_email`).

### Alternate DNS Lookup

If an alternate DNS server is designated in the `kproxy` command line, then at ingest, the server at the provided IP: Port combination (e.g., `127.0.0.1:5353`) looks up IP addresses in the flow records from this `kproxy`. The returned host names will be stored in Kentik's KDE backend in Source Hostname and Destination Hostname dimensions (see **IP Info Dimensions**. These directional hostname dimensions can be used for group-by and filtering in queries.  
  
To avoid including public IP addresses from the Internet in the IP-to-hostname resolution,  the designation of an alternate DNS server involves specifying to `kproxy` which IPs should be considered as "internal." The following command line arguments are used:

* **-dns** (optional): The IP address and the port of the local DNS that will be used for IP-to-hostname resolution. The -dns command line parameter supports only one server. If an`internal` flag is included in the argument, `kproxy` will include the RFC1918 IP space and other common private ranges in the IP addresses on which it performs IP-to-hostname resolution:

  + With flag: `internal:ip-address:port`
  + Without flag: `ip-address:port`

    > **Note*:*** If the internal flag is not included, the `-my_network argument` is required.
* `-my_network`: The "internal" prefixes for IP-to-hostname resolution. The IPs to use may be specified as one of the following:  
   `- prefixes:` A comma-separated list of IP prefixes.  
   `- netclass:` A value that references the IPs listed in the **Internal IPs** field on the Kentik portal's**Network Classification Page**.  
   `- netclass, prefixes:` The `netclass` value included in the list of prefixes.

> **Notes*:***
>
> * If the `-my_network` argument is not provided but the -dns argument includes the internal flag, `kproxy` will perform IP-to-hostname resolution on only the RFC1918 IP space and other common private ranges.
> * If the `-my_network argument` includes netclass and IPs not specified in Network Classification, `kproxy` will fail with an error.

The table below shows how `kproxy` determines the IPs for IP-to-hostname resolution based on combinations of the internal flag and the `-my-network` parameter.

| Internal flag | my\_network | Included IPs |
| --- | --- | --- |
| N | none | None; `kproxy` start will fail with an error. |
| N | `netclass` | Internal IPs from Network Classification |
| N | `prefixes` | CIDRs specified with `my_network` |
| N | `netclass, prefixes` | Internal IPs from Network Classification plus CIDRs specified with `my_network` |
| Y | none | Private only (RFC1918, etc.) |
| Y | `netclass` | Private (RFC1918, etc.) plus Internal IPs from Network Classification |
| Y | `prefixes` | Private (RFC1918, etc.) plus CIDRs specified with `my_network` |
| Y | `netclass, prefixes` | Private (RFC1918, etc.), Internal IPs from Network Classification, and CIDRs specified with `my_network` |

The following example shows how to configure the proxy agent (version 7.43.0 or later) to derive the Hostname dimensions using an alternate DNS server that resolves the prefixes specified by `my_network`:

```
/usr/bin/kproxy -local_proxy http://10.10.2.2:3128 -proxy-http 192.168.2.16:3182 -bootstrap_devices 104884,88726 -dns 192.168.1.2:53 -my_network 192.168.1.0/24,192.168.2.0/24 -api_email=kproxy-agent+xxxx@kentik.com -log_level=info -host=0.0.0.0 -proxy_host=127.0.0.1 -healthcheck=0.0.0.0 -port=9995
```

> **Notes*:***
>
> * To generate directional hostname dimensions, your organization should be proxying flows through `kproxy` version 7.43.0 or newer. To enable alternate DNS resolution with earlier versions of `kproxy`, see **Customer Care****.**
> * When the IP addresses of an ingested flow record can't be resolved, the directional hostname dimensions in the corresponding KDE record will be left unspecified.
> * The `-dns` argument is distinct from **Custom DNS**, which affects only the display of hostnames in the portal and specifies an alternate DNS server for flows from all sources rather than just from an individual instance of `kproxy`.
> * If an alternate DNS server is specified with `-dns`, the hostname returned for a given IP address won't be checked again for 24 hours.
> * The `kproxy` can use the local resolvers configured on the hosts they are running on. For example, `-dns internal:127:0.0.1:53` will use the (one) local DNS server running on port 53.
> * Setting up load balancing on the DNS resolvers used by `kproxy` is not typically necessary, however, reach out to **Customer Care** for options if you think you have a special case.
> * Due to the volume of flow records from high-rate devices, Kentik may not be able to resolve all IP addresses to hostnames, even with alternate server lookups.

### kproxy for Streaming Telemetry

> **Note*:*** Using Streaming Telemetry with `kproxy` is distinct from using Streaming Telemetry for Kentik NMS (see **NMS via Streaming Telemetry**).

The following `kproxy` arguments are specified only when `kproxy` will be used to receive streaming telemetry data from devices and forward it to KDE (see **SNMP and Streaming Telemetry**). Leave these arguments unspecified in all other scenarios:

* `-st_dialout_listener` (required for ST on Junos or Cisco): Set to auto to enable receipt by `kproxy` of streaming telemetry data collected and sent either by devices running Junos or by Cisco IOS-XRv 9000 routers.
* `-st_udp_bind` (required for ST on Junos): Specifies the IP address and port on which `kproxy` should listen for streaming telemetry data collected and sent as UDP by devices running Junos (version 18.4R2.7 or later).
* `-st_tcp_bind` (required for ST on Cisco): Specifies the IP address and port on which `kproxy` should listen for streaming telemetry data collected and sent as TCP by devices such as Cisco IOS-XRv 9000 routers (version 6.2.3 or later).

The combination of above arguments used to enable Streaming Telemetry with Kentik varies depending on the data source:

* To enable ST via `kproxy` from devices running Junos:  
   - Always use `st_dialout_listener`, set to `auto.`  
   - Also use `st_udp_bind`, typically set to `0.0.0.0:9555`.
* To enable ST via `kproxy` from Cisco IOS-XRv 9000 routers:  
   - Always use `st_dialout_listener`, set to `auto`.  
   - Also use `st_tcp_bind`, typically set to `0.0.0.0:9555`.
* To enable ST from both Junos devices and Cisco IOS-XRv 9000 routers to the same instance of `kproxy`:  
   - Always use `st_dialout_listener`, set to `auto`.  
   - Also use both `st_tcp_bind` and `st_udp_bind`, both typically set to `0.0.0.0:9555`.

> **Note*:*** Setting the binding arguments above to the Kentik-recommended value of `0.0.0.0:9555` enables `kproxy` to listen on the specified port on all IPs on the server.

### `kproxy` CLI Reference

The table below shows the complete set of `kproxy` arguments, along with their descriptions, as returned from the `-h` and `--help` commands. Either " " or "=" can be used between an argument's name and value. Arguments of type “bool” don’t require a value to be specified.

| Argument | Type | Description |
| --- | --- | --- |
| `-active_listen` | string | Run Active measurement server on this host:port combo (0.0.0.0:8080) |
| `-active_max_speed` | float | Cap active measurement at this rate (bytes/sec). 0 is uncapped. |
| `-active_protocol` | string | Active measurement protocol. (udp|tcp|icmp) (default "tcp") |
| `-active_servers` | string | Active measurement servers to connect to. Host and port, comma separated. (127.0.0.1:8080,127.0.02:8080) |
| `-alert_fwd` | string | Obsolete |
| `-api` | string | API to use (default "https://api.kentik.com/api/internal") |
| `-api_email` | string | Email to use when connecting to the CH API |
| `-api_pass` | string | Password to use when connecting to the CH API |
| `-api_token` | string | Password to use when connecting to the CH API |
| `-base_port` | int | Port to begin sending flow to as more devices export flow (default 40010) |
| `-bgp_neighbor_asn` | int | [deprecated] ASN of BGP neighbor |
| `-bgp_neighbor_ip` | string | IP of IPv4 BGP neighbor |
| `-bgp_neighbor_ip6` | string | IP of IP6 BGP neighbor |
| `-bgp_proxy` | string | [deprecated] BGP proxy IP:PORT |
| `-bgp_proxy6` | string | [deprecated] BGP proxy IP:PORT |
| `-bgp_config_file` | string | BGPd JSON configuration file |
| `-bootstrap_devices` | string | Spin up a pre-sent client for these devices. Comma separated list of device ids (12002,1221) |
| `-client_id` | string | Client ID (company\_id:device\_name:device\_id) (default "100:default:110") |
| `-cloud_map` | string | File to use for mapping cloud ranges |
| `-custom_cols` | string | Map of custom column names to ids for this device |
| `-device_stat` | string | Obsolete |
| `-device_stat_ssl_url` | string | Obsolete |
| `-device_subtype` | string | Client Device Subtype |
| `-device_type` | string | Client Device Type |
| `-dns` | string | Try to resolve flow addresses against this host:port (UDP only) |
| `-dnstagd` | string | If set, the address:port to contact for a dnstag assignment |
| `-dump` | int | Dump Every MS (default 1000) |
| `-exit_low_traffic` | bool | Exit if traffic is 0 for 15 min |
| `-extensible_cols` | string | Map of extensible column names to ids for this device |
| `-fg` | bool | Run in foreground, not using a guarding process. |
| `-force_bgp` | bool | Obsolete |
| `-forked_pprof_host` | string | Bind PPROF for forked clients to this interface and the healthcheck port (optional) |
| `-geo` | string | GEO File Location (default "/usr/share/cloudhelix/country.csv") |
| `-geo_small` | bool | Use small geo file |
| `-healthcheck` | string | Bind to this interface to allow healthchecks (127.0.0.1) (default "127.0.0.1") |
| `-host` | string | Host to listen for flow on (default "127.0.0.1") |
| `-hostname` | string | Obsolete |
| `-http` | string | HTTP port to bind on (default "off") |
| `-http_udr` | string | Run an HTTP server here, listening for UDR metrics passed in. Write these into flow and send along. |
| `-interface_metadata_comparison` | string | Interface metadata comparison type |
| `-just_active` | bool | Only perform active measurements |
| `-leader_id` | string | Leader ID (company\_id:device\_name:device\_id) for sdmi |
| `-listen` | string | Obsolete |
| `-local_proxy` | string | Route HTTP via this proxy, e.g., “http://proxy:3128” |
| `-log_level` | string | Logging Level: access|debug|info|warn|error|panic|off (default "info") |
| `-logstash` | string | logstash endpoint (kproxy only), ‘DISABLE’ to disable |
| `-mask_file` | string | Truncates results of IP matches to the corresponding subnet based on netmask in CIDR format. List must contain netmask for v4,v6 and may contain CIDR blocks that are outside of default mask. Format is: first line: v4,v6 cidrs, After list of cidrs |
| `-max` | int | Max FPS (default 4000) |
| `-max_big` | int | Max Big FPS (default 30) |
| `-max_http_send` | int | Number of outstanding HTTP connections (default 10) |
| `-metrics` | string | Metrics Configuration. none|stderr|graphite:127.0.0.1:2003 |
| `-my_network` | string | Comma separated list of cidrs to treat as internal (10.10.0.0/16,23.1.3.0/24) |
| `-num_http_send` | int | Number of HTTP sending threads (default 5) |
| `-pcap_flow_period` | int | Obsolete |
| `-pcap_transmit_period` | int | Obsolete |
| `-port` | int | Port to listen for flow on (default 9995) |
| `-proxy-http` | string | Set up forward HTTP proxy at the port |
| `-proxy-http-allow` | string | If running a forward proxy, a single domain to which proxying should be allowed (default “.kentik.com”) |
| `-proxy_host` | string | Interface to use for proxied clients. Defaults to -host value if not set |
| `-proxy_pcap` | bool | Obsolete |
| `-proxy_snmp` | string | Interface to use for proxied clients on SNMP polling. Defaults to -host value if not set |
| `-rateLimit` | uint | Obsolete (default 12000) |
| `-rcvbuf_size` | int | number of bytes to give to the main UDP rcv buffer if in proxy mode |
| `-realtime` | bool | Obsolete (default true) |
| `-region` | string | Way to set common flags for a given region |
| `-remap_file` | string | File to use to remap IPs while in agent mode |
| `-rpki_bind` | string | Run RPKI server at this address (default off) |
| `-rpki_cache` | string | URL of the cached RPKI JSON data (default off) |
| `-rpki_local_cache` | string | Local place to cache rpki data (default "/data/blob-storage/rpki/local.cache") |
| `-sampleInterval` | int | Number of milliseconds across which any flow downsampling should occur, when running in proxy mode (default 15000) |
| `-sample_metric` | int | Sample Rate to use for metrics like SNMP counters (default 1) |
| `-sample_rate` | int | Sample rate - default of 0 means we'll get it from the server |
| `-servers` | string | Comma separated list of CHF Servers to talk to (default "tcp://127.0.0.1:9996") |
| `-sflow_agent_address` | bool | If set, use the agent address for sending ip (`kproxy` only) |
| `-sflow_drop_egress` | bool | If a sflow device, drop packet-samples that we know to be egress packets (datasource == dst-interface) |
| `-shard_num` | int | For a sharded device, what shard am I |
| `-site_id` | int | ID of Kentik site to look devices up in |
| `-sleep` | int | Sleep this long (MS) before proceeding to init Client |
| `-snmp` | string | Obsolete |
| `-snmp_community` | string | SNMP community |
| `-snmp_file` | string | File to use to look up SNMP data while in agent mode |
| `-snmp_ip` | string | IP Address to poll for SNMP data |
| `-snmp_low` | bool | Poll for SNMP at a low rate. Turn off SNMP counter polling. |
| `-snmp_max` | int | Number of device interfaces we import from SNMP, sorted by traffic (default 750) |
| `-snmp_max_cycle` | int | Number of device interfaces we import from SNMP, sorted by traffic, after initial run (default 500) |
| `-snmp_max_repetitions` | int | The max repetitions for snmp polling (default 50) |
| `-snmp_retries` | int | The retries setting for snmp polling |
| `-snmp_timeout_seconds` | int | The timeout seconds settings for snmp polling (default 5) |
| `-snmp_v3` | string | SNMP V3 Info |
| `-sql_proxy_listen` | string | Bind sql proxy for client on this host:port (default "127.0.0.1:20017") |
| `-st_dialout_listener` | string | Process dialout ST. Supported types: junos|ios|auto |
| `-st_tcp_bind` | string | Run TCP ST dialout server at this address (default off) |
| `-st_udp_bind` | string | Run UDP ST dialout server at this address (default off) |
| `-st_v1` | string | STv1 info (JSON) |
| `-stdout` | bool | Log using stdout |
| `-tag_update` | int | Obsolete (default 1200) |
| `-tag_url` | string | URL to pull tags from (default "https://api.kentik.com/api/internal/company/COMPANY\_ID/device/DEVICE\_ID/tags") |
| `-tee_kproxy` | string | If set, tee traffic to this location in the `kproxy` case. Format is (<src-ip>:)<dst-ip>:<dst-port> |
| `-test_snmp` | bool | Run using the passed in config to check SNMP and then exit. |
| `-tls_ca_file` | string | x509 CA file to use |
| `-tls_cert_file` | string | x509 cert file to use |
| `-tls_key_file` | string | x509 cert key file to use |
| `-total_shards` | int | For a sharded device, total number of shards |
| `-type` | string | Client Type: sflow|ipfix|v5|v9|proxy|kflow|snmp\_check (default "proxy") |
| `-update_on_restart` | bool | Update the code when starting. |
| `-use_asn_from_flow` | bool | If set, use asn from flow instead of clearing to default asn |
| `-use_ue_from_flow` | bool | If set, use ue from flow instead of doing UE enrichment |
| `-v` | bool | Version |
| `-verbose` | bool | Obsolete |

## SNMP Configuration File

> **Note*:*** A local config file should be used to specify SNMP settings only when customer information security policies prohibit the configuration of SNMP settings in the Kentik portal.

By default, the SNMP configuration (SNMP IP and SNMP Community) for a given device that sends flow to Kentik (e.g., router; see **Supported Device Types**) is learned by `kproxy` from that device's settings in the Kentik portal (see **Device SNMP Settings**). There may be circumstances, however, in which it is necessary (e.g., for security compliance) not to specify SNMP settings for a given router in the portal. In this case, it is possible instead to specify the settings through `kproxy` configuration, using the optional `-snmp_file` command line argument to direct `kproxy` to get that information from a local config file.  
  
When SNMP is configured with an external file, the required SNMP parameters are set from the values in that file. These values are described in the following table:

| Parameter | Description |
| --- | --- |
| `device_id` | Required*:* The Kentik assigned ID of the device. |
| `snmp_comm` | Required*:* The device's SNMP community. |
| `snmp_ip` | Required*:* The IP address that should be used to poll the router. |
| `minimize_snmp` | Required when device type is router:   * If false (standard), interface counter will be polled every 5 minutes and interface description every 30 minutes; * If true (minimized), interface counter won't be polled, and interface description will be polled every 6 hours. |
| `snmp_v3` | Optional*:* Configuration settings required for SNMP v3 (see SNMP V3 Config). |

#### SNMP V2 Config

The required settings are stored in the config file as JSON key/value pairs. The following example shows a local configuration file for two devices using SNMP V2, with `minimize_snmp` set to true for the second device:

```
{

  "devices": [
    {
      "device_id": 2466,
      "snmp_comm": "device 2466 community string",
      "snmp_ip": "polling.ip.of.first_device"
    },
    {
      "device_id": 2681,
      "snmp_comm": "device 2681 community string",
      "snmp_ip": "polling.ip.of.second_device",
      "minimize_snmp": true
    }

  ]

}
```

#### SNMP V3 Config

When using SNMP v3, the additional required settings are specified with a `snmp_v3` object. The object contains the following parameters:

| Parameter | Description |
| --- | --- |
| `UserName` | Required: The user name for SNMP v3 authentication. |
| `AuthenticationProtocol` | Required: The SNMP v3 authentication protocol:   * NoAuth (must be specified when no authentication is used) * MD5 * SHA |
| `AuthenticationPassphrase` | Required unless `AuthenticationProtocol` is NoAuth: Password for SNMP V3 authentication. |
| `PrivacyProtocol` | Required: The SNMP V3 privacy type:   * NoPriv (must be specified when no privacy is used) * DES (56-bit encryption) * AES-128 |
| `PrivacyPassphrase` | Required unless `PrivacyProtocol` is NoPriv: Password for SNMP V3 privacy. |

An example `snmp_v3` object is shown in the following sample code:

```
{

  "devices": [
    {
      "device_id": 3030,
      "snmp_comm": "device 3030 community string",
      "snmp_ip": "polling.ip.of.third_device",
      "minimize_snmp": false,
      "snmp_v3": "{\"UserName\":\"SNMPv3-Admin\",\"AuthenticationProtocol\":\"SHA\",\"AuthenticationPassphrase\":\"big_secret\",\"PrivacyProtocol\":\"AES\",\"PrivacyPassphrase\":\"unguessable\"}"
    }
  ]
}
```

## `kproxy` Debugging

The following tips may be useful in debugging issues related to the use of `kproxy`:

* Our article on **Router Configuration** will guide you through the general setup of routers to work with Kentik.
* The IPs allowed in the **Agent** tile on the Admin » **Access Control** page (see**Access Control**) must include the public IP of the server running `kproxy`. If the server is behind a NAT gateway, you can get its public IP by running `wget -qO- ifconfig.co` on the server.
* If the `kproxy` command line argument `-metrics` was set to `stderr`, you will receive a checkpoint every minute that indicates how much flow you are receiving from the router. If that count is not increasing, there is an issue between your router and `kproxy`, either router configuration or `kproxy` config of communication between them.
* It may take 2-3 minutes for the agent to download flow templates and begin to process flow. You can expect to receive errors ("`[ipfix_parse_msg] no template for 256, skip data set`") for the first few minutes, after which the errors should stop.
* Errors will be logged in `stdout` by default.
* A `kproxy` debug and health check port is opened by default on the loopback address (127.0.0.1), one port higher than the configured flow ingest port (e.g., for the default flow port of `udp/9995`, heathcheck is `tcp/9996`). The health check information fields are detailed in **Health Check Fields**.

#### Health Check Fields

The heath check output includes information about two classes of devices:

* **Connected Devices:** Devices that are already registered in Kentik and currently sending flow.
* **Unregistered Devices:** Devices that are not yet registered in Kentik but have sent flow in the last 30 minutes.

The health check output states the number of Connected Devices and returns the following fields for each:

* **In1:** Difference in incoming (from device) packet count over 1 minute.
* **Out1:** Difference in outgoing (to Kentik) packet count over 1 minute.
* **In15:** Difference in incoming (from device) packet count over 15 minutes.
* **Out15:** Difference in outgoing (to Kentik) packet count over 15 minutes.
* **Last seen:** Last time device sent a packet.
* **Sources:** IP addresses sending device packets.
* **Channel highwater:** Maximum lengths of internal message queues.

The health check output also states the number of Unregistered Devices and returns the following fields for each:

* **Source IP address:** Port sending flow packets.
* **Timestamp:** Time of first received flow packet.

## `kproxy` Syslog Parsing

The `kproxy` Syslog Parsing feature is no longer supported.
