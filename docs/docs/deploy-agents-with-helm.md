# Deploy Universal Agents with Helm

Source: https://kb.kentik.com/docs/deploy-agents-with-helm

---

This article covers how to use Helm (the package manager for Kubernetes) to deploy and manage the lifecycle of of Kentik’s Universal Agent within your cluster.

## Overview

The provided **Helm chart** supports two primary deployment patterns, **StatefulSet** and **DaemonSet**, to accommodate different architectural requirements and telemetry use cases.

> **TIP**: See the **Kubernetes Workload Management**documentation for background information on choosing between these workload patterns.

### Prerequisites

* **Kubernetes**: v1.20+
* **Helm**: v3.x
* **Kentik Account**: Valid **Company ID**
* **Network Access**: Egress access to `grpc.api.kentik.com:443`

## Install the Chart

Following the official repository documentation, you can install the chart either by cloning the repository or using the OCI registry.

### Option A: Clone the Repository

```
# Clone the repository
git clone https://github.com/kentik/kagent-helm.git
cd kagent-helm

# Install with StatefulSet pattern
helm install kagent . \
  --set-string kagent.companyId=<YOUR_COMPANY_ID>
```

### Option B: Direct OCI Install (Recommended)

```
# Install directly via OCI
helm install kagent oci://ghcr.io/kentik/kagent-helm \
  --set-string kagent.companyId=<YOUR_COMPANY_ID>
```

## **Deployment Patterns**

### **Pattern 1: StatefulSet (Persistent Storage)**

* **Identity Persistence**: The agent's identity is tied to an **ed25519** keypair. This pattern uses PVCs or Secrets to ensure that if a pod restarts, it retains the same ID in the Kentik Portal.
* **Storage**: Creates two persistent volumes by default:

  + `/opt/kentik` (10Gi) for application data, buffers, and logs
  + `/opt/ua/keys` (100Mi) for identity keypair

> **IMPORTANT**: **Keypair Management (Agent Identity)**
>
> The agent's ed25519 keypair MUST persist. If the keypair is lost, the agent will appear as a "New" device in Kentik. This forces the customer to manually reauthorize the agent in the Portal, causing telemetry gaps and orphaned records.

### **Pattern 2: DaemonSet (Node-Level)**

* **Coverage**: Exactly one pod per node.
* **Storage Requirement**: Requires `persistence.type: hostPath` or `persistence.type: emptyDir`. You must ensure the host directories exist and have correct permissions:

```
sudo mkdir -p /var/lib/kagent/data /var/lib/kagent/keys
sudo chown -R 500:500 /var/lib/kagent
```

## **Container Security & Linux Capabilities**

To adhere to the principle of least privilege, the agent drops all capabilities by default. Specific Linux capabilities are added back only for the Agent Features you enable.

> **Note**: By default, only `NET_RAW` is added. If you require `NET_ADMIN` or `NET_BIND_SERVICE` for features like `kdns`, `kbgp`, or `kproxy`, you must explicitly add them to `securityContext.capabilities.add` in your `values.yaml` file.

| **Linux Capability** | **Associated Agent Feature** | **Purpose** |
| --- | --- | --- |
| `NET_ADMIN` | `kdns` | Network administration and DNS telemetry. |
| `NET_BIND_SERVICE` | `kbgp`, `kproxy` | Binding to privileged ports (<1024). |
| `NET_RAW` | NMS (ranger), `ksynth`, `livesynth` | Raw socket access for packet capture and Synthetics. |

### **Storage Options (persistence.keypair.type)**

| **Type** | **Description** |
| --- | --- |
| `secret` (Recommended) | Best for production and GitOps. Replicas expect secrets named `{{ .Release.Name }}-{{ index }}-secret`. |
| `pvc` | Traditional persistent volume approach (StatefulSet only). |
| `hostPath` | Uses local node storage (required for DaemonSet with persistence). |
| `emptyDir` | Temporary storage (not recommended for production). |

### **Generating Secrets**

If using the secret storage type, generate them before installation:

```
# Generate secrets for 3 replicas
./generate-secrets.sh 3
kubectl apply -f generated_secrets/generated_secrets.yaml
```

## **Configuration Reference**

| Parameter | **Description** | **Default** |
| --- | --- | --- |
| `kagent.companyId` | Kentik Company ID (**Required**) | "" |
| `kagent.apiEndpoint` | Kentik API endpoint | `grpc.api.kentik.com:443` |
| `deploymentType` | Pattern: StatefulSet or DaemonSet | `statefulset` |
| `persistence.enabled` | Enable persistent storage | `true` |
| `persistence.type` | pvc, hostPath, or emptyDir | `pvc` |
| `persistence.pvc.size` | Data volume size | `10Gi` |
| `persistence.keypair.type` | secret, pvc, hostPath, or emptyDir | `secret` |
| `resources.requests.cpu` | CPU request | `1` |
| `resources.requests.memory` | Memory request | `1024Mi` |
| `resources.limits.cpu` | CPU limit | `2` |
| `resources.limits.memory` | Memory limit | `4096Mi` |

## **Troubleshooting**

### **Verify Connectivity**

To troubleshoot connection issues, first check the logs. If connectivity is suspected, verify that the agent can reach the Kentik gRPC endpoint on port 443.

#### Option A: Using Netcat (if available)

```
kubectl exec -it <pod-name> -- nc -zv grpc.api.kentik.com 443
```

#### Option B: Using Bash (for minimal images without tools)

If the container does not have nc or telnet, use the built-in Bash TCP handler:

```
kubectl exec -it <pod-name> -- bash -c "timeout 1 bash -c '</dev/tcp/grpc.api.kentik.com/443' && echo 'Port is open' || echo 'Connection failed'"
```

### **Check Identity Secrets**

If using secret persistence, verify the keypair is mounted:

```
# Should contain private_key.pem and public_key.pem
kubectl exec -it kagent-0 -- ls -la /opt/ua/keys/
```
