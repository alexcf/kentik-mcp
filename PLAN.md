# kentik-mcp: Full API Coverage Plan

## Current State (51 tools)

| Domain | Read | Write | Notes |
|--------|------|-------|-------|
| Devices | ✅ | ✅ | V5 endpoints |
| Interfaces | ✅ read-only | ❌ no write | V5 only |
| Labels | ✅ | ✅ | V5 endpoints |
| Sites | ✅ | ✅ | V5 endpoints |
| Users | ✅ read-only | ❌ no write | V5 only |
| Flow Tags | ✅ read-only | ❌ no write | V5 only |
| Query (flow data) | ✅ | N/A | |
| Alerting | ⚠️ active alarms only | ❌ no policy CRUD | Missing policy management entirely |
| Synthetics tests | ✅ | ✅ | V6 |
| Synthetics agents | ✅ read-only | ❌ no write | Missing agent alerts, update, delete |
| AI Advisor | ✅ | N/A | |
| BGP Monitoring | ❌ | ❌ | Completely missing |
| AS Groups | ❌ | ❌ | Completely missing |
| Notification Channels | ❌ | ❌ | Completely missing |
| Capacity Plan (API) | ❌ | N/A | Have flow-based approx only |
| Cloud Export | ❌ | ❌ | Completely missing |
| Network Classification | ❌ | ❌ | Completely missing |
| Connectivity Costs | ❌ | N/A | Completely missing |
| Credentials Vault | ❌ | N/A | Completely missing |
| Pathfinder | ❌ | N/A | Completely missing |
| KMI | ❌ | N/A | Completely missing |
| MKP (multi-tenant) | ❌ | ❌ | Completely missing |
| Site Markets | ❌ | ❌ | Part of site-apis, missing |
| NMS Metrics | ❌ | N/A | Not in public API schema |

---

## Implementation Batches

### Batch 1 — Complete the core admin CRUD
*Everything an ops team needs to manage Kentik day-to-day.*

| Tool | Method | Endpoint |
|------|--------|----------|
| `kentik_create_user` | POST | `/user/v202211/users` |
| `kentik_update_user` | PUT | `/user/v202211/users/{id}` |
| `kentik_delete_user` | DELETE | `/user/v202211/users/{id}` |
| `kentik_reset_user_sessions` | PUT | `/user/v202211/users/{id}/reset_active_sessions` |
| `kentik_reset_user_api_token` | PUT | `/user/v202211/users/{id}/reset_api_token` |
| `kentik_create_tag` | POST | `/flow_tag/v202404alpha1/tag` |
| `kentik_update_tag` | PUT | `/flow_tag/v202404alpha1/tag/{id}` |
| `kentik_delete_tag` | DELETE | `/flow_tag/v202404alpha1/tag/{id}` |
| `kentik_create_interface` | POST | `/interface/v202108alpha1/interfaces` |
| `kentik_update_interface` | PUT | `/interface/v202108alpha1/interfaces/{id}` |
| `kentik_delete_interface` | DELETE | `/interface/v202108alpha1/interfaces/{id}` |
| `kentik_classify_interfaces` | POST | `/interface/v202108alpha1/manual_classify` |

**Files:** `write_users.go`, `write_tags.go`, `write_interfaces.go`
**Estimated tools added:** 12

---

### Batch 2 — Alerting + Notifications (end-to-end)
*Create policies, set thresholds, configure where alerts go.*

| Tool | Method | Endpoint |
|------|--------|----------|
| `kentik_list_alert_policies` | GET | V5 `/alerts` |
| `kentik_get_alert_policy` | GET | V5 `/alert/{id}` |
| `kentik_create_alert_policy` | POST | V5 `/alert` |
| `kentik_update_alert_policy` | PUT | V5 `/alert/{id}` |
| `kentik_delete_alert_policy` | DELETE | V5 `/alert/{id}` |
| `kentik_list_notification_channels` | GET | `/notification_channel/v202210/notification_channels` |
| `kentik_get_notification_channel` | GET | `/notification_channel/v202210/notification_channels/{id}` |
| `kentik_search_notification_channels` | POST | `/notification_channel/v202210/notification_channels/search` |
| `kentik_start_manual_mitigation` | POST | V5 `/alerts/manual-mitigate` |

**Files:** `write_alerting.go`, `notifications.go`
**Estimated tools added:** 9

---

### Batch 3 — BGP Monitoring
*BGP route visibility, prefix monitoring, RPKI.*

| Tool | Method | Endpoint |
|------|--------|----------|
| `kentik_list_bgp_monitors` | GET | `/bgp_monitoring/v202210/monitors` |
| `kentik_get_bgp_monitor` | GET | `/bgp_monitoring/v202210/monitors/{id}` |
| `kentik_create_bgp_monitor` | POST | `/bgp_monitoring/v202210/monitors` |
| `kentik_update_bgp_monitor` | PUT | `/bgp_monitoring/v202210/monitors/{id}` |
| `kentik_delete_bgp_monitor` | DELETE | `/bgp_monitoring/v202210/monitors/{id}` |
| `kentik_set_bgp_monitor_status` | PUT | `/bgp_monitoring/v202210/monitors/{id}/status` |
| `kentik_get_bgp_metrics` | POST | `/bgp_monitoring/v202210/metrics` |
| `kentik_get_bgp_routes` | POST | `/bgp_monitoring/v202210/routes` |

**Files:** `bgp.go`
**Estimated tools added:** 8

---

### Batch 4 — Synthetics completeness
*Fill the gaps in synthetics: agent alerts, agent management.*

| Tool | Method | Endpoint |
|------|--------|----------|
| `kentik_list_agent_alerts` | GET | `/synthetics/v202309/agentAlerts` |
| `kentik_get_agent_alert` | GET | `/synthetics/v202309/agentAlerts/{id}` |
| `kentik_create_agent_alert` | POST | `/synthetics/v202309/agentAlerts` |
| `kentik_update_agent_alert` | PUT | `/synthetics/v202309/agentAlerts/{id}` |
| `kentik_delete_agent_alert` | DELETE | `/synthetics/v202309/agentAlerts/{id}` |
| `kentik_update_synthetic_agent` | PUT | `/synthetics/v202309/agents/{id}` |
| `kentik_delete_synthetic_agent` | DELETE | `/synthetics/v202309/agents/{id}` |

**Files:** `write_synthetics.go` (extend existing)
**Estimated tools added:** 7

---

### Batch 5 — AS Groups + Network Classification
*Traffic analysis grouping and custom network tagging.*

| Tool | Method | Endpoint |
|------|--------|----------|
| `kentik_list_as_groups` | GET | `/as_group/v202212/as_group` |
| `kentik_get_as_group` | GET | `/as_group/v202212/as_group/{id}` |
| `kentik_create_as_group` | POST | `/as_group/v202212/as_group` |
| `kentik_update_as_group` | PUT | `/as_group/v202212/as_group/{id}` |
| `kentik_delete_as_group` | DELETE | `/as_group/v202212/as_group/{id}` |
| `kentik_get_network_classification` | GET | `/network_class/v202109alpha1/network_class` |
| `kentik_update_network_classification` | POST | `/network_class/v202109alpha1/network_class` |

**Files:** `as_groups.go`, `network_classification.go`
**Estimated tools added:** 7

---

### Batch 6 — Capacity, Costs, Cloud
*Read-only intelligence: what's full, what costs what, cloud flows.*

| Tool | Method | Endpoint |
|------|--------|----------|
| `kentik_list_capacity_plans` | GET | `/capacity_plan/v202212/capacity_plan` |
| `kentik_get_capacity_plan` | GET | `/capacity_plan/v202212/capacity_plan/{id}` |
| `kentik_get_capacity_summary` | GET | `/capacity_plan/v202212/capacity_plan/summary` |
| `kentik_list_cost_providers` | GET | `/cost/v202308/cost/providers` |
| `kentik_get_cost_summary` | GET | `/cost/v202308/cost/summary` |
| `kentik_list_cloud_exports` | GET | `/cloud_export/v202210/exports` |
| `kentik_get_cloud_export` | GET | `/cloud_export/v202210/exports/{id}` |
| `kentik_create_cloud_export` | POST | `/cloud_export/v202210/exports` |
| `kentik_update_cloud_export` | PUT | `/cloud_export/v202210/exports/{id}` |
| `kentik_delete_cloud_export` | DELETE | `/cloud_export/v202210/exports/{id}` |

**Files:** `capacity_api.go`, `costs.go`, `cloud_export.go`
**Estimated tools added:** 10

---

### Batch 7 — Pathfinder + KMI
*Path analysis and internet intelligence.*

| Tool | Method | Endpoint |
|------|--------|----------|
| `kentik_create_pathfinder_report` | POST | `/pathfinder/v202505beta1/create` |
| `kentik_list_kmi_markets` | GET | `/kmi/v202212/markets` |
| `kentik_get_kmi_network` | POST | `/kmi/v202212/market/{marketId}/network/{asn}/{type}` |
| `kentik_get_kmi_rankings` | POST | `/kmi/v202212/market/{marketId}/rankings/{rankType}` |

**Files:** `pathfinder.go`, `kmi.go`
**Estimated tools added:** 4

---

### Batch 8 — MKP (Multi-Tenant) + Site Markets
*Service provider / multi-tenant management.*

| Tool | Method | Endpoint |
|------|--------|----------|
| `kentik_list_mkp_tenants` | GET | `/mkp/v202407/tenants` |
| `kentik_get_mkp_tenant` | GET | `/mkp/v202407/tenants/{id}` |
| `kentik_create_mkp_tenant` | POST | `/mkp/v202407/tenants` |
| `kentik_update_mkp_tenant` | PUT | `/mkp/v202407/tenants/{id}` |
| `kentik_delete_mkp_tenant` | DELETE | `/mkp/v202407/tenants/{id}` |
| `kentik_list_mkp_packages` | GET | `/mkp/v202407/packages` |
| `kentik_get_mkp_package` | GET | `/mkp/v202407/packages/{id}` |
| `kentik_create_mkp_package` | POST | `/mkp/v202407/packages` |
| `kentik_update_mkp_package` | PUT | `/mkp/v202407/packages/{id}` |
| `kentik_delete_mkp_package` | DELETE | `/mkp/v202407/packages/{id}` |
| `kentik_list_site_markets` | GET | `/site/v202211/site_markets` |
| `kentik_create_site_market` | POST | `/site/v202211/site_markets` |
| `kentik_update_site_market` | PUT | `/site/v202211/site_markets/{id}` |
| `kentik_delete_site_market` | DELETE | `/site/v202211/site_markets/{id}` |

**Files:** `mkp.go`, extend `write_sites.go`
**Estimated tools added:** 14

---

### Batch 9 — Credentials Vault
*Read-only credential group management for NMS.*

| Tool | Method | Endpoint |
|------|--------|----------|
| `kentik_list_credential_groups` | GET | `/credential/v202407alpha1/group` |
| `kentik_get_credential_group` | GET | `/credential/v202407alpha1/group/{id}` |

**Files:** `credentials.go`
**Estimated tools added:** 2

---

## V5 APIs Still to Add

These don't appear in the V6 schema but are documented in V5:

| Domain | Missing |
|--------|---------|
| Alert Policies | Full CRUD (`POST /alert`, `PUT /alert/{id}`, `DELETE /alert/{id}`) |
| Custom Dimensions | Full CRUD + populators |
| Saved Filters | Full CRUD |
| Plan API | `GET /plans` |
| Audit Log | `GET /audit-log` |

---

## Summary

| Batch | New Tools | Priority | Audience |
|-------|-----------|----------|----------|
| 1 — Core admin CRUD | 12 | 🔴 High | Everyone |
| 2 — Alerting + Notifications | 9 | 🔴 High | Everyone |
| 3 — BGP Monitoring | 8 | 🟠 Medium | Network ops |
| 4 — Synthetics completeness | 7 | 🟠 Medium | Everyone |
| 5 — AS Groups + Net Classification | 7 | 🟠 Medium | Network ops |
| 6 — Capacity, Costs, Cloud | 10 | 🟡 Normal | Network ops |
| 7 — Pathfinder + KMI | 4 | 🟡 Normal | Network ops |
| 8 — MKP + Site Markets | 14 | 🟢 Low | Service providers |
| 9 — Credentials Vault | 2 | 🟢 Low | NMS admins |
| V5 gaps | ~10 | 🟠 Medium | Everyone |

**Current: 51 tools → Target: ~134 tools**

---

## Implementation Notes

1. **V5 vs V6 endpoints**: Where both exist (labels, sites, users), prefer V6 — better typed, more consistent
2. **Naming convention**: `kentik_{verb}_{noun}` — list, get, create, update, delete, search
3. **Error handling**: All handlers already follow the pattern — return `ToolResultError` not Go errors
4. **Each batch = one PR** — keeps diffs reviewable and testable independently
5. **NMS metrics**: Not in public API schema. Needs separate investigation — likely behind a private gRPC endpoint

