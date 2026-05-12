# BGP Proxy Agent

Source: https://kb.kentik.com/docs/bgp-proxy-agent

---

Kentik’s `kbgp` is a software agent that proxies BGP peering sessions between Kentik and your devices.

![Traffic flows between Kentik and a kbgp proxy agent installed in a customer network.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/BPA-Traffic_flows_diagram-454h1120w(1).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A31Z&se=2026-05-12T09%3A35%3A31Z&sr=c&sp=r&sig=RI7MwW7xys1wq36sMCZe9I0hQo5KwHpZ7rs%2BkdGXE3I%3D)

*Traffic flows between Kentik and a* `kbgp` *proxy agent installed in a customer network.*

## About BGP Proxy

Kentik's BGP proxy agent, `kbgp`, enables devices within a customer's private network to establish secure BGP peering sessions with Kentik without the need to assign those devices a public IP address. Multiple customer devices can establish BGP peering with a single proxy agent, which will multiplex and relay BGP packets to Kentik in real time over a secure gRPC session.  
  
The key benefits of deploying `kbgp` in a Kentik-monitored network include:

* BGP data will be secured and encrypted while it transits from your network to Kentik.
* Kentik will be able to add BGP context to the flow data from internal network devices that don't have global IP connectivity.
* Because `kbgp` does not store BGP state or BGP routes, the agent is lightweight and requires very little resources.

The steps involved in downloading, installation, and configuration of `kbgp` are described the topics below.

## `kbgp` Requirements

The following resources must be available (at minimum) to support the use of `kbgp` running on a VM:

* Linux operating system
* RAM allocation of 1 GB
* One CPU core
* Establishment of an HTTP/2 connection on TCP port 443 to Kentik:

  + **US cluster**: `grpc.api.kentik.com`
  + **EU cluster**: `grpc.api.kentik.eu`

> **Note**: Actual requirements may vary depending on how many BGP sessions are handled by `kbgp`.

## `kbgp` Download and Install

This article discusses the download and installation of `kbgp`, which may be deployed via a Debian/Ubuntu package, an RPM package, or a Docker image.

### Gather Required Information

To install `kbgp`, you'll need to gather your organization's "system account" email address and API token from the Kentik portal, which are used for authentication. Here's how to find them:

1. In the Kentik portal, go to the Settings » **Proxy Agents** page. Click the **Add kproxy Agent** button at the upper right.
2. In the resulting dialog, click the Debian/Ubuntu card to expand the instructions.

   > **Note:** You need not be planning to install kbgp on a Debian/Ubuntu system.
3. On the **Manually** tab of the **Run proxy** step:

   1. Copy the value of the `-api_email` argument.
   2. Copy the value of the `-api_token` argument.

You will need the above system account values when specifying install arguments for `kbgp`.

### `kbgp` Docker Install

To install `kbgp` via Docker:

1. In your terminal, pull down the latest `kbgp` image from Kentik’s Docker Hub repository:

   ```
   $ docker pull kentik/kbgp:latest
   ```
2. To comply with security best practices, create an environment variable named `KENTIK_API_TOKEN` to store the authentication token, which is the `-api_email` value that you copied in step 3 of **Gather Required Information**.
3. Run the `kbgp` docker image:

   ```
   $ docker run --rm -d --network host --name kbgp -e KENTIK_API_EMAIL=system_account@company.com -e KENTIK_ENABLE_MD5=true kentik/kbgp:latest
   ```

   > **TIP**: Storing the token as an environment variable prevents exposure via process (ps) command. If you prefer not to use an environment variable, include the  `-e KENTIK_API_TOKEN=System_Account_Token` argument in the `docker run` command.

> **Notes:**
>
> * If your company is registered with Kentik in the EU, the docker run command in step 3 must also include the following argument:  
>    `-e KENTIK_REGION=fra1`
> * For additional options, including arguments for specifying a different IP address and/or port, see **kbgp CLI Arguments**.

### `kbgp` Package Install

Kentik’s `kbgp` repository offers both .rpm and .deb packages compatible with all major distributions and versions of Linux. You'll first download and install the package, then configure and run `kbgp` via `systemd`:

1. In terminal, download and install the `kbgp` repository:

   1. **Debian/Ubuntu**:

      ```
      $ curl -s https://packagecloud.io/install/repositories/kentik/kbgp/script.deb.sh | sudo bash 
      $ sudo apt-get install kbgp
      ```
   2. **RPM**:

      ```
      $ curl -s https://packagecloud.io/install/repositories/kentik/kbgp/script.rpm.sh | sudo bash 
      $ yum install kbgp
      ```

      > **Note:** If the `kbgp` package is not found, your operating system may be unsupported. To try a manual installation, select a version from the repository at **https://packagecloud.io/kentik/kbgp**.
2. Make a copy of the supplied example service configuration file:

   ```
   $ sudo cp /etc/default/kbgp.env.sample /etc/default/kbgp.env
   ```
3. Update the following variables in the new `/etc/default/kbgp.env` service configuration file:

   1. `KENTIK_API_TOKEN`(required): The system account `-api_token` value that you copied in step 3 of **Gather Required Information** above.
   2. `KENTIK_API_EMAIL`(required): The system account `-api_email` value that you copied in step 3 of **Gather Required Information** above.
   3. `KENTIK_REGION` (EU customers only): The region in which your organization is registered with Kentik. If you're registered on our EU cluster, specify as `fra1`. If you're registered on our US cluster, you need not include this variable, but if you do, specify it as `iad1`.
   4. `KENTIK_ENABLE_MD5`: Set to “true” if any of your BGP sessions use the MD5 key.
   5. `KENTIK_SITE`: The Kentik site ID of the site where the `kbgp` agent is deployed.

      > **Note:** Configure the `KENTIK_SITE` argument only if your BGP peering network devices have overlapping IP addresses. Devices with overlapping IP addresses must BGP peer with a different `kbgp` agent and each device must be assigned to a different site in the Kentik portal.
4. Configure use of web proxy in the service configuration file (optional):

   ```
   HTTP_PROXY=http://<user>:<pass>@<proxy-ip>:<proxy-port>
   HTTPS_PROXY=http://<user>:<pass>@<proxy-ip>:<proxy-port>
   ```

   > **Note:** Configure the web proxy if your network only allows the use of web proxy.
5. Configure `kbgp` to start automatically at boot:

   ```
   $ sudo systemctl enable kbgp
   ```
6. Restart `kbgp` with the new updated configuration:

   ```
   $ sudo systemctl restart kbgp
   ```
7. Check the `kbgp` logs to verify that there are no unexpected errors:

   ```
   $ journalctl -fu kbgp
   ```

## Configure Devices for `kbgp`

Use the `kbgp` agent to establish a BGP session with a given device, requiring proper configuration of that device on the device itself and also in the Kentik portal.

### Portal `kbgp` Settings

To configure the device in Kentik:

1. Navigate to the Devices page (Settings » **Network Devices**).
2. In the **Device List**, find the row corresponding to the device with which you'd like to establish a session.
3. Click the edit icon at the right to open the **Device Settings Dialog**.
4. Click the BGP tab (see **Device BGP Settings**).
5. From the **BGP Type** drop-down, select **Peer with device**.
6. In the **Peering Address** field for your preferred IP version, enter the source IPv4 or IPv6 address of the BGP session for the `kbgp` agent.
7. Click **Save** to save changes and exit the dialog.

### Device `kbgp` Configuration

To configure the device itself:

1. Set the device's BGP peer address to the IP address of the `kbgp` agent.
2. In the terminal, check the logs to verify successful connection to the `kbgp` agent:

```
$ journalctl -fu kbgp

May 05 20:01:43 kbgp01 kbgp[17986]: 2022-05-05T20:01:43Z INF opening Data stream client-metadata={

  "backend-pin":["42b0e9c9"],
  "counter-down":["0"],
  "counter-up":["0"],

  "device-ip":["10.250.15.1"],
  "device-port":["65411"],
  "instance":["64bae291-66ea-4910-a2a2-92eb99224a28"],

  "proto":["BGP"],
  "proxy-ip":["10.250.12.10"],
  "proxy-port":["179"],

  "session-id":["cbce1788-bb34-4c2c-8056-5b3224d4e893"],
  "version":["dirty-ebc215b30f4b2f4e73f0915ca4cf4436e1623cf5"],
  "x-ch-auth-api-token":["myKentikApiToken"],

  "x-ch-auth-email":["myEmail@myCompany.com"]

} 

device=1289:a6r_kx1_gobgpd:127353-ipv4 local-addr=10.250.12.10:179 module=proxy/client-data peer-addr=10.250.15.1:65411 service=ktbgp-agent session-id=cbce1788-bb34-4c2c-8056-5b3224d4e893
```

## `kbgp` Command Line

The use of the `kbgp` command line interface is covered here.

### `kbgp` CLI Arguments

The following topics cover the most common command line arguments for `kbgp`. For a brief reference to additional arguments, see **kbgp CLI Reference**.

> **Notes:**
>
> * Either " " or "=" can be used between an argument's name and value.
> * Use `-h` or `--help` to return a list of arguments and its description.

#### `kbgp` Required Arguments

The following arguments are required to authenticate access to `kbgp`:

* `--email:` The system account `-api_email` value that you copied in step 3 of Gather Required Information above.
* `--token:` The system account `-api_token` value that you copied in step 3 of Gather Required Information above.
* `--region` (EU customers only): If your organization is registered with Kentik in the EU (rather than the US), you must include this argument and set the value to `fra1`.

### `kbgp` CLI Reference

The following table shows the complete set of `kbgp` arguments as returned from the `-h` and `--help` commands. Either " " or "=" can be used between an argument's name and value.

| Argument | Type | Description |
| --- | --- | --- |
| `--disable-log-export` | bool | Disable log export to Kentik |
| `--email` | string | Kentik API email ($KENTIK\_API\_EMAIL) |
| `--enable-md5-support` | float | Enable MD5 password support (requires root) ($KENTIK\_ENABLE\_MD5) |
| `--human-logs` | bool | Enable human-readable logs |
| `--insecure` | bool | Connect over insecure HTTP to the target |
| `--listen` | string | Listen address(es) to bind to |
| `--log-export-url` | string | URL to export logs to. Default value is "https://grpc.api.kentik.com/agents/v202204/logs" |
| `--log-level` | string | Log level. Default value is “info” |
| `--md5-listen-interface` | string | Network interface to capture connection attempts on |
| `--metrics-addr` | string | Metrics endpoint bind address. Default value is "127.0.0.1" |
| `--metrics-export-url` | string | URL to export metrics to. Default value is "https://grpc.api.kentik.com/agents/v202204/metrics" |
| `--metrics-port` | string | Metrics endpoint port. Default value is 20920 |
| `--multi` | int | Multiplex N client connections over 1 gRPC connection. Default value is 10 |
| `--region` | string | Kentik environment/region ($KENTIK\_REGION) |
| `--site-id` | int | Kentik site ID to look devices up in (default means any site) ($KENTIK\_SITE) |
| `--target` | string | Target for proxied BGP sessions (overrides region). Default value is "grpc.api.kentik.com:443" |
| `--token` | string | Kentik API token ($KENTIK\_API\_TOKEN) |
| `--version` | bool | Version |

## `kbgp` Status APIs

The `kbgp` agent can be queried via an API endpoint to report status information (in JSON) that can help Kentik Product Support troubleshoot any issues (see **Customer Care**). Request (cURL) and response (JSON) examples for various uses of the interface are provided.

### `kbgp` Overall Status

Request the overall status of the agent:

```
$ curl -s http://127.0.0.1:20920/hc | jq.
```

**Example Response**:

```
{
  "timestamp": "2023-07-25T17:01:25.409183151Z",
  "status": "ok"
}
```

The following status values may be returned:

* `ok`: All checks are passing.
* `initializing`: The initial health check isn't yet finished.
* `timeout`: The timeout expired before the checks were complete.
* `old_result`: The last Result is older than the acceptable Result delay.
* `failing`: One or more checks returned an error.

### `kbgp` Agent Details

Request details about the agent:

```
$ curl http://127.0.0.1:20920/admin/status\?pretty
```

**Example Response**:

```
{
  "InstanceID": "9ef12a66-47ff-45b5-b56b-23a76941c655",
  "ControlStream": {
      "Connected": true, "Since": "2023-07-24T16:37:51.158790049Z", "DurationMs": 74925903, "MessagesRx": 1311
  },
  "Heartbeats": {
    "LastTx": "2023-07-25T13:25:51.159072268Z", "MsSinceLastTx": 45902, "LastTxRTTMs": 0, "LastRx": "2023-07-25T13:25:51.194488212Z", "MsSinceLastRx": 45867
  },
  "ActiveSessions": 1,
  "Config": {
    "Listeners": [ "0.0.0.0:179", "[::]:179" ],
    "MD5Enabled": true,
    "MD5Interface": "any",
    "ProxyTarget": "grpc.api.our1.kentik.com:443",
    "APIEmail": "alistair+bgpproxy@kentik.com",
    "SiteID": 0
  }
}
```

### `kbgp` Sessions Summary

Request a summary of current sessions on the agent:

```
$ curl http://127.0.0.1:20920/admin/sessions\?pretty
```

**Example Response**:

```
{
  "NumSessions": 1,
  "Sessions": [{
    "Peer": "10.250.15.2:42875", "Time": "2023-07-25T14:00:22.696499347Z", "DurationMs": 16831, "Connected": true
  }]
}
```

### `kbgp` Session Details

Request details about an individual session on the agent:

```
$ curl http://127.0.0.1:20920/admin/sessions/10.250.15.2\?pretty
```

**Example Response:**

```
{
  "Peer": "10.250.15.2:44507",
  "ConnectedSince": "2023-07-25T16:48:01.132557342Z",
  "ConnectedDurationMs": 21881,
  "SessionID": "abc9a042-61b1-4114-8f65-6b4db10d9121",
  "ServerMetadata": {
    "content-type": [
      "application/grpc"
    ],
    "counter-down": [
      "0"
    ],
    "counter-up": [
      "0"
    ],
    "date": [
      "Tue, 25 Jul 2023 16:48:02 GMT"
    ],
    "referrer-policy": [
      "strict-origin-when-cross-origin"
    ],
    "server": [
      "envoy"
    ],
    "strict-transport-security": [
      "max-age=31536000; includeSubDomains; preload;"
    ],
    "x-content-type-options": [
      "nosniff"
    ],
    "x-frame-options": [
      "DENY"
    ],
    "x-ratelimit-details": [
      "eyJkZXNjcmlwdG9yX3N0YXR1c2VzIjpbeyJjb2RlIjoxLCJjdXJyZW50X2xpbWl0Ijp7InJlcXVlc3RzX3Blcl91bml0IjoxODAwLCJ1bml0IjoyfSwibGltaXRfcmVtYWluaW5nIjoxNzk4LCJkdXJhdGlvbl91bnRpbF9yZXNldCI6eyJzZWNvbmRzIjo1OX19XX0"
    ],
    "x-request-id": [
      "4792f21a-58a8-969a-9f45-df5c09a43b24"
    ],
    "x-xss-protection": [
      "1; mode=block"
    ]
  }
}
```

## `kbgp` Troubleshooting Notes

The following may help with troubleshooting `kbgp`-related issues:

* Kentik BGP ingest cluster maintenance may cause BGP session restarts.
* Peering the same BGP router with two `kbgp` agents simultaneously is unsupported.
* In high availability deployments, only one session is established with one `kbgp` agent at a time. For high availability failover, use `Keepalived` with a shared Virtual IP (VIP) address between multiple servers, where only one will be active.
* Incorrectly configured devices are logged once and temporarily blocked for 20 minutes:

  ```
  2023-07-25 17:01:49.87 INF Deny-listed an INVALID / UNKNOWN device 
    component=proxy/client-data 
    device_ip=::1 
    local-addr=[::1]:179 
    peer-addr=[::1]:49285 
    service=ktbgp-agent 
    sessid=904e3b00-d6c5-4b2c-b40a-67b0b4f128a8 
    ttl_counter=0 
    ttl_seconds=1200
  ```
