# Migration Scenarios

Source: https://kb.kentik.com/docs/kproxy-migration-scenarios

---

Select the migration path from `kproxy` to **Universal Agent** that matches your observability needs:

### **Scenario 1: Flow Collection Only**

Use this path if you require NetFlow, sFlow, or IPFIX ingestion without immediate SNMP-based enrichment.

1. Install and register the Universal Agent on your host.
2. In the Kentik portal, navigate to **Settings »** **Universal Agents**.
3. Enable the **Flow Proxy** capability for the agent.
4. Configure the local listening port (default 9995) and map the agent to the appropriate Site and My Network settings.
5. In **Settings »** **Networking Devices**, edit the target device(s) to set their SNMP option to **Disable SNMP Collection from this device**.

### Scenario 2: Enriched Flow without Full Monitoring

Use this path for enriched flow collection. SNMP data is used to map traffic data to interface names, descriptions, and speeds but the device will not be monitored by Kentik NMS or have access to related features.

1. Enable the **Flow Proxy** capability as described in **Scenario 1**.
2. Enable the **SNMP/ST** capability on the same agent instance.
3. In **Settings »** **Networking Devices**, edit the target device(s) to set their SNMP option to to **Agent-based SNMP for Flow Enrichment**.
4. Configure your SNMP v2c or v3 credentials and specify the desired polling intervals.

### Scenario 3: Enriched Flow and Full Monitoring

> **Note**: This is the **Kentik-recommended** **path** for full visibility and requires a **DevicePak** license. Collected flow traffic data is enriched with interface names, descriptions, and speeds. The device’s performance, health, and operational state are also monitored by NMS.

1. Enable the **Flow Proxy** capability as described in **Scenario 1**.
2. Enable the **SNMP/ST** capability on the same agent instance.
3. In **Settings »** **Networking Devices**, edit the target device(s) to set their SNMP option to to **Agent-based SNMP for Flow Enrichment**.
4. Configure your SNMP v2c or v3 credentials and specify the desired polling intervals.
