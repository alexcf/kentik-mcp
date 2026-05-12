# Install Multiple Agents on a Host

Source: https://kb.kentik.com/docs/linux-install-multiple-agents-per-host

---

This article describes how to deploy multiple Kentik Universal Agents on a single Linux host.

> **Note**: For more information on deploying Universal Agents, see **Universal Agents**.

While the standard deployment model for Kentik is one Universal Agent (`kagent`) per host, you can install multiple agents on the same Linux system. This approach can be useful for logically isolating agents by function or team, or for testing different agent configurations simultaneously on one machine.

**Process Summary**: You’ll modify the standard Linux shell script installer to specify unique installation directories and `systemd` service names for each agent instance (see **Install Multiple Agents (Linux)**).

> **Warning**: If you run multiple Universal Agents (or other network services) on a single Linux host, to avoid conflicts be sure to review the default ports used by Kentik shown in **Port Usage and Potential Conflicts**.

## Port Usage and Potential Conflicts

Some of Kentik’s Universal Agent capabilities can open network ports on the Linux host to, for example, listen for inbound data or for internal communication, or make outbound connections for polling/scraping activities. If unmitigated, port conflicts could occur when two services attempt to listen on the same port and protocol, resulting in one of them failing.

To avoid this, review the default ports used by Kentik capabilities that listen for traffic, and take the necessary steps to mitigate conflicts.

> **Tip**: When running multiple flow collection capabilities on a Linux host (either within an agent or across agents), ensure each capability is configured to listen on a unique port.

| Capability | Default Port(s) | Protocol | Direction / Scope | Purpose |
| --- | --- | --- | --- | --- |
| **Flow Collector (sFlow)** | 6343 | UDP | Inbound | Receives sFlow datagrams from network devices. |
| **Flow Collector (NetFlow)** | 9995, 2055 | UDP | Inbound | Receives NetFlow v5/v9 packets from network devices. |
| **Flow Collector (IPFIX)** | 4739 | UDP | Inbound | Receives IPFIX packets from network devices. |
| **Synthetics (LiveSynth)** | 19020 | TCP | Localhost | Used for internal gRPC communication by the Scamper service. |
| **DNS Monitoring (kdns)** | 53 (Monitored) | TCP/UDP | Capture | Does not open a port but captures traffic destined for port 53. |
| **Discovery (Ranger)** | (None) | UDP | Outbound | Initiates outbound SNMP connections to devices on port 161. |

## Installation Parameters

To install more than one Universal Agent on a Linux host, you’ll append these two query parameters to each installation URL obtained from the Kentik portal, as described in **Install Multiple Agents (Linux)**:

| **Parameter** | Description | Default Value | Required |
| --- | --- | --- | --- |
| `install_root` | The absolute path to the agent's installation root directory. Must be unique per agent. | /opt/kentik | Yes |
| `systemd_service` | The name for the agent's `systemd` unit file. Must be unique per agent. | kagent | Yes |

> **Note**: This method only applies to the `curl | sh` installation method for Linux (see **Install Multiple Agents (Linux)**. For Docker-based deployments, multiple agents are run as separate containers, each with a unique name (`--name`) and a dedicated host volume mount (`--mount`).

## Install Multiple Agents (Linux)

1. **Generate the Base URL**: In the Kentik portal, navigate to the **Universal Agents** page (Settings » **Universal Agents**).

   * Click **+ Deploy Agent** and select the Linux installation option.
   * Copy the URL from the generated one-line installation command.
   * The base URL will look similar to `https://grpc.api.kentik.com/install/COMPANY_ID`
2. **Construct the Modified URL**: For each agent you intend to install:

   * Append the `install_root` and `systemd_service` parameters to the base URL (see **Installation Parameters**).
   * Ensure the values for each parameter are unique for each agent.
   * **Example**: Using a base URL for company ID 220277 and targeting an installation for a logical agent named “agentA”, the fully constructed URL would be:  
     `https://grpc.api.kentik.com/install/220277?install_root=/opt/kentik/agentA&systemd_service=kagent-agentA`
3. **Execute the Installation Command**: Run the `curl | sh` command for each agent using its unique URL (see **Example: Install Two Agents**).

### Example: Install Two Agents

This example demonstrates installing two logically separate Kentik Universal Agents, Agent A and Agent B, on the same Linux host.

#### Agent A Installation

Append `install_root=/opt/kentik/agentA` and `systemd_service=kagent-agentA` to the URL:

```
# Command for Agent A
curl -fsSL \  "https://grpc.api.kentik.com/install/220277?install_root=/opt/kentik/agentA&systemd_service=kagent-agentA" | sudo sh
```

#### Agent B Installation

Append `install_root=/opt/kentik/agentB` and `systemd_service=kagent-agentB` to the URL:

```
# Command for Agent B
curl -fsSL \  "https://grpc.api.kentik.com/install/220277?install_root=/opt/kentik/agentB&systemd_service=kagent-agentB" | sudo sh
```

## Manage Multiple Agents (Linux)

Managing agents installed in custom directories requires targeting the specific `kagent` binary within the installation path or using the custom `systemd` service name.

### Repair or Uninstall an Agent

To repair or uninstall an agent, execute the `kagent` binary located in that agent's `/bin` directory.

```
# Repair Agent B
/opt/kentik/agentB/bin/kagent repair

# Uninstall Agent B
/opt/kentik/agentB/bin/kagent uninstall
```

### Manage the Agent Service

Use `systemctl` with the custom service name to manage the agent service.

```
# Check the status of Agent A
sudo systemctl status kagent-agentA

# Stop Agent B
sudo systemctl stop kagent-agentB
```

#### Example `systemd` Unit Files

The installer script creates a service unit file for each agent in `/etc/systemd/system/`. The critical differences are the `ExecStart` path and the `Environment` variables, which must point to the unique `install_root` directory for each agent.

> **Note**: Running multiple agents on a single host will increase its resource consumption (CPU, memory, network I/O). Monitor the host to ensure it has sufficient capacity to handle the combined load.

**File: /etc/systemd/system/kagent-agentA.service**

```
[Unit]
Description=Kentik Agent (agentA)
Wants=network-online.target
After=network-online.target

[Service]
OOMScoreAdjust=-500
LimitNOFILE=524288
ExecStart=/opt/kentik/agentA/bin/kagent agent run
User=kentik
Environment=K_ROOT=/opt/kentik/agentA
Environment=K_SYSTEMD_SERVICE=kagent-agentA
AmbientCapabilities=CAP_NET_ADMIN CAP_NET_BIND_SERVICE CAP_NET_RAW CAP_SETGID CAP_SETUID CAP_SYS_CHROOT CAP_SYS_RESOURCE
Restart=always
RestartSec=1s

[Install]
WantedBy=multi-user.target
```

**File: /etc/systemd/system/kagent-agentB.service**

```
[Unit]
Description=Kentik Agent (agentB)
Wants=network-online.target
After=network-online.target

[Service]
OOMScoreAdjust=-500
LimitNOFILE=524288
ExecStart=/opt/kentik/agentB/bin/kagent agent run
User=kentik
Environment=K_ROOT=/opt/kentik/agentB
Environment=K_SYSTEMD_SERVICE=kagent-agentB
AmbientCapabilities=CAP_NET_ADMIN CAP_NET_BIND_SERVICE CAP_NET_RAW CAP_SETGID CAP_SETUID CAP_SYS_CHROOT CAP_SYS_RESOURCE
Restart=always
RestartSec=1s

[Install]
WantedBy=multi-user.target
```
