# Troubleshooting

Source: https://kb.kentik.com/docs/ssh-command-access-troubleshooting

---

Having trouble getting Kentik's**AI Advisor** to interact with your network devices? This article covers the most common issues that prevent **Config Context** (configuration backups) and **Command Access** (live diagnostics) from functioning properly.

## Why Configuration History May Be Unavailable

Configuration collection in Kentik typically requires several key components:

* **Device Configuration Backup Feature:** This must be enabled in your **Kentik plan** to allow for configuration history tracking.
* **Proper Device Credentials:** Ensure that read-only SSH/API credentials or NETCONF access is correctly configured.
* **Universal Agent Configuration:** Agents need the necessary permissions to collect configurations from the devices.
* **Supported Device Types:** Not all **device types** support configuration collection, which may limit the ability to track changes.

## Why Command Access May Fail

If AI Advisor is unable to execute diagnostic `show` commands, check the following common issues:

* **Authorization Restrictions:** The specific `show` command must be authorized in your local parser view (Cisco IOS-XE) or centralized AAA (TACACS+/RADIUS) policy. If the `kentik_ro` user is explicitly denied access to the command, the execution will fail.
* **SSH Credential Rejection:** Verify that the Universal Agent's SSH key or password is still valid and hasn't been rotated, expired, or removed from the device.
* **Privilege Level Issues:** Kentik does not interactively elevate privileges (e.g., typing `enable` and a secondary password). Ensure the service account drops directly into the correct privilege level or role required to run operational commands.
* **Feature Not Enabled:** **Double-check** that the **Enable Read-Only Diagnostic Commands** toggle is actually switched to **On** within the device's **SSH** settings tab.

## Alternative Ways to Track Device Changes

When configuration history isn't available via AI Advisor, you can monitor device changes through:

* **Configs Tab in Device Details**: View config versions and diffs directly from the **Device Details** page.
* **Syslog Monitoring:** Filter by device name to see configuration-related events. Look for messages such as "config commit" or "configuration changed".
* **SNMP Traps:** Many devices send traps on configuration changes. Check for `coldStart`, `warmStart`, or config change traps.
* **NMS Metrics:** Monitor device uptime, track component changes in hardware inventory, and watch for routing protocol changes (e.g., BGP, OSPF state changes).
