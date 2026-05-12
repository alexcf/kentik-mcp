# Configuration Procedures (UI)

Source: https://kb.kentik.com/docs/kproxy-migration-configuration

---

Follow these steps in the Kentik portal to configure Flow and SNMP collection via the Universal Agent.

1. **Access Agent Configuration for Capability Install**
2. Configure Flow Proxy Capability
3. Configure SNMP/ST Capability

## Install Agent Capabilities

1. Log in to the Kentik portal and navigate to **Settings »** **Universal Agents**.
2. Identify the target agent in the list, then click it to open the management panel.
3. Click **Install**next to the capability you wish to install.

## Configure Flow Proxy Capability

1. In **Settings »** **Universal Agents**, select **Flow Proxy** from the capability list.
2. Click the **Edit** button to configure the following **parameters**:

   * Host: The local IP address (default `0.0.0.0`).
   * Port: The UDP port for incoming telemetry (e.g., `9995`).
   * Site: The site associated with this data.
   * My Network: The internal network boundary definition.
3. Click **Save**.

## Configure SNMP/ST Capability

1. Navigate to **Settings »****Networking****Devices**.
2. Click the device name from the list you wish to configure.
3. In the top right of the device page, select the gear icon.
4. Select the **SNMP** tab and configure the enrichment settings as described in the menu:

   1. **Credentials:** Create a new **credential** or use an existing credential.
   2. **Collection Agent:** Select the Kentik agent to collect ST data.
   3. **Monitoring Template:** Choose a template to be applied for SNMP monitoring.
   4. **Device SNMP IP:** Enter the IP that Kentik should use to connect to your SNMP Device.
   5. **Port:** Enter a port number for SNMP polling of your device.
5. Click **Save**.
