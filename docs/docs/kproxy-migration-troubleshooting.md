# Troubleshooting

Source: https://kb.kentik.com/docs/kproxy-migration-troubleshooting

---

While you **migrate from kproxy to the Universal Agent**, if you encounter issues, consult the table below for common symptoms and their resolutions (see also: **FAQ**).

| Symptom | Technical Root Cause | Resolution |
| --- | --- | --- |
| **Agent Status "Down" (Exit 10)** | Configuration validation failure | Ensure Site and **My Network** fields are assigned in the portal under **Settings »** **Universal Agents**. |
| **Agent Status "Up" but No Data** | Local firewall or routing path blockage | * Check local host firewalls (e.g.,  `iptables`, `ufw`, or `firewalld`). * Verify that flow traffic is actually reaching the host by running `sudo tcpdump -i any port 9995`. |
| **Missing Interface Names (Showing Index IDs)** | SNMP connectivity or polling failure | * Verify the **SNMP/ST** capability is installed and showing an "Up" status. * Ensure the SNMP tab in Device Settings is set to **Agent-based SNMP for Flow Enrichment**. * Ensure community strings are correct. |
| **SNMP Polling Fails Post-Migration** | Credential mismatch | Verify that the SNMP v2c/v3 strings and passwords in **Organization Settings »** **Credentials Vault** exactly match the device configurations. |
