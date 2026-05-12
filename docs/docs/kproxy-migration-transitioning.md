# Transitioning from Legacy Kproxy

Source: https://kb.kentik.com/docs/kproxy-migration-transitioning

---

If you are currently running the standalone `kproxy` binary or a Docker container, follow these steps to transition to the **Universal Agent** with minimal data interruption.

> **Operational Strategy: Warm Failover**
>
> * To ensure no telemetry is lost, the Universal Agent should be configured and set to "Up" status **before** disabling the legacy process.
> * Network devices should be updated to point to the UA IP/Port, and once data flow is verified in the portal, the legacy service can be decommissioned.

### Linux Binary Transition

If `kproxy` is running as a `systemd` service:

```
# 1. Verify the Universal Agent is 'Up' in the Kentik portal

# 2. Stop and disable the legacy kproxy service
sudo systemctl stop kproxy
sudo systemctl disable kproxy

# 3. Confirm the UA is now binding to the telemetry port (e.g., 9995)
sudo ss -lupn | grep :9995
```

> **Note**: If you changed the default port in **Configure Flow Proxy**, replace `9995` with your custom port in the above command.

### **Docker Container Transition**

If `kproxy` is running as a container, follow these validation steps to ensure data continuity and a safe rollback path.

#### Step 1: Stop the Legacy Container

First, identify and stop the running container without removing it. This allows for an immediate rollback if the new Universal Agent configuration needs adjustment.

```
# 1. Identify the running kproxy container
docker ps | grep kproxy

# 2. Stop the legacy container
docker stop <container_id_or_name>
```

> **Note**: Do not remove the container yet. Leave it in a stopped state so you can quickly restore it if needed.

#### Step 2: Verify UA is Successfully Receiving Flow

Before permanently removing the container, confirm the Universal Agent is actively collecting data:

1. In the Kentik portal, navigate to **Settings »** **Universal Agents** and confirm the **Flow Proxy** capability shows a status of "Up".
2. Navigate to **Settings** **»** **Networking Devices** and confirm that flow data transmission from the relevant devices has resumed. Look for active traffic records from the interfaces previously handled by kproxy.
3. Allow **5–10 minutes** of monitoring to ensure flow continuity and rule out transient collection issues.

> **Note**: After deploying the Universal Agent and enabling the Flow Proxy capability, you may notice that the new agent also appears in the portal under **Settings »** **Kproxy Agents**. Because the UA's Flow Proxy utilizes the same core engine as the legacy agent, this dual-listing is expected behavior and does not indicate a migration failure.

#### Step 3: Finalize or Rollback

**If verification succeeds — Remove the legacy container**

Once you have confirmed the UA is actively collecting flow data, permanently remove the stopped container:

```
# Remove the legacy container
docker rm <container_id_or_name>
```

**If verification fails — Restore the legacy container**

If flow data is not appearing in the portal after 5–10 minutes, you can safely restore the legacy container to resume collection while you troubleshoot:

> **CRITICAL**: You may need to stop the Universal Agent service/container before restarting the legacy container to avoid a port binding conflict (e.g., both agents attempting to bind to UDP 9995).

```
# Restart the legacy container to restore previous state
docker start <container_id_or_name>
```

Once the legacy agent is running again, review the **Troubleshooting** section to diagnose the Universal Agent issue before reattempting the transition.
