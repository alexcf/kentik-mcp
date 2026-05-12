# Command Access

Source: https://kb.kentik.com/docs/ai-assisted-netops-command-access

---

When investigating device performance issues, logging in and running CLI commands to understand what’s happening is cumbersome, error-prone, and hindered by knowledge gaps.

The **Command Access** feature empowers Kentik’s **AI Advisor** to identify, execute, and analyze the results of read-only `show` commands in response to user prompts. This provides immediate, pre-analyzed diagnostic context, often without needing to know CLI syntax.

> **IMPORTANT:** Ensure your environment meets the Prerequisites (including Universal Agent deployment and TCP/22 reachability) before attempting to configure device synchronization.

### Setup: Device Diagnostics Commands

To establish the secure tunnel required for Kentik to interact with the device CLI, navigate to the **SSH** tab in **Device Settings** in the Kentik portal.

![Configuration settings for SSH access and device diagnostics.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/ANMS-edit-device-ssh-tab.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A07Z&se=2026-05-12T09%3A31%3A07Z&sr=c&sp=r&sig=jSz522%2FxArrVYy6eTTwVKjXCZcayMT6jGohZkAMpt2Q%3D)

1. **Enable SSH command access**: Toggle this **On** to establish the main connection.
2. **Credential**: Select your pre-configured **SSH key or password**.
3. **Collection Agent**: Select the **agent** with network reachability to the device.
4. **Device Platform**: Recommended setting is **Autodetect** to ensure correct syntax usage.
5. **SSH Hostname (or IP):** The Fully Qualified Domain Name (FQDN) or IP address the agent will use to connect.
6. **Enable Read-Only Diagnostic Commands**: Toggle this **On** to empower AI Advisor to securely run `show` commands (e.g., `show ip bgp summary`) for immediate context.

   > **Note**: Results are not shared between users.

> **TIP**: If the connection fails initially, verify that the Universal Agent has accepted the device's SSH host key or that your firewall isn't blocking the initial handshake on TCP/22.

### Example Prompts for AI Advisor

* **Routing:** "Check the BGP summary on [Device Name] and tell me if any peers are currently down."
* **Interfaces:** "Show me the interface counters for `eth0` on [Device Name] and highlight any recent input errors."
* **Troubleshooting:** "Run a routing table check for [IP Address] on [Device Name] and explain the next hop."
* **Security:** "Review the current active SSH sessions on [Device Name] and list the connected IP addresses."
