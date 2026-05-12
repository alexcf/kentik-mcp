# Device APIs

Source: https://kb.kentik.com/docs/device-apis

---

This article covers how to get started with the Device APIs:

> **Notes:**
>
> * The **API Tester** in the v4 Portal enables you to securely call the methods of these APIs using your organization's Kentik data.
> * For information about devices, start with **Network Devices**.
> * Protobuf and OpenAPI specifications for Kentik’s v6 APIs are available in our **api-schema-public** repository.

## Device Usage

The Device API provides programmatic access to configuration of devices

## Device RPCs

The tables below provide a quick reference to key information about each Remote Procedure Call in these APIs. Use the links in the tables to see schemas for the request and/or response body (if any).

> **Note:** To test methods using your own Kentik data, use the portal's **API Tester**.

### ListDevices

**API: DeviceService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| GET | /device/v202308beta1 /device | Returns list of configured devices (see [About Devices](https://kb.kentik.com/v4/Cb01.htm)). |
| |  |  | | --- | --- | | **Request body**: None  **Parameters**: None | **Response body**:  v202308beta1ListDevicesResponse | | | |

### CreateDevice

**API: DeviceService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| POST | /device/v202308beta1 /device | Create configuration for a new device. Returns the newly created configuration (see [About Devices](https://kb.kentik.com/v4/Cb01.htm)). |
| |  |  | | --- | --- | | **Request body**: v202308beta1CreateDeviceRequest  **Parameters**:  None | **Response body**:  v202308beta1CreateDeviceResponse | | | |

### DeleteDevices

**API: DeviceService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| POST | /device/v202308beta1 /device/batch | Deletes configuration of multiple devices with specific IDs (see [About Devices](https://kb.kentik.com/v4/Cb01.htm)). |
| |  |  | | --- | --- | | **Request body**: v202308beta1DeleteDevicesRequest  **Parameters**: None | **Response body**:  v202308beta1DeleteDevicesResponse | | | |

### UpdateDevices

**API: DeviceService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| PUT | /device/v202308beta1 /device/batch | Replaces configuration of multiple devices with attributes in the request. Returns the updated configurations (see [About Devices](https://kb.kentik.com/v4/Cb01.htm)). |
| |  |  | | --- | --- | | **Request body**: v202308beta1UpdateDevicesRequest **Parameters**: None | **Response body**:  v202308beta1UpdateDevicesResponse | | | |

### UpdateDevice

**API: DeviceService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| PUT | /device/v202308beta1 /device/{device.id} | Replaces configuration of a device with attributes in the request. Returns the updated configuration (see [About Devices](https://kb.kentik.com/v4/Cb01.htm)). |
| |  |  | | --- | --- | | **Request body**: v202308beta1UpdateDeviceRequest | **Response body**:  v202308beta1UpdateDeviceResponse |  **Parameters:**  | Name | Description | Required | Type | | --- | --- | --- | --- | | device.id | System generated unique identifier | true | string | | | |

### GetDevice

**API: DeviceService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| GET | /device/v202308beta1 /device/{device.id} | Returns configuration of a device specified by ID (see [About Devices](https://kb.kentik.com/v4/Cb01.htm)). |
| |  |  | | --- | --- | | **Request body**: None | **Response body**:  v202308beta1GetDeviceResponse |    **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | device.id | ID of the requested device | true | string | | | |

### DeleteDevice

**API: DeviceService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| DELETE | /device/v202308beta1 /device/{device.id} | Deletes configuration of a device with specific ID (see [About Devices](https://kb.kentik.com/v4/Cb01.htm)). |
| |  |  | | --- | --- | | **Request body**: None | **Response body**:  v202308beta1DeleteDeviceResponse |  **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | device.id | ID of the device to be deleted | true | string | | | |

### UpdateDeviceLabels

**API: DeviceService**

| REST Method | REST Endpoint | Description |
| --- | --- | --- |
| PUT | /device/v202308beta1 /device/{id}/labels | Removes all existing labels from the device and applies the device labels (see [About Device Labels](https://kb.kentik.com/v4/Cb16.htm)) specified by id. Returns the updated configuration. |
| |  |  | | --- | --- | | **Request body**: v202308beta1UpdateDeviceLabelsRequest | **Response body**:  v202308beta1UpdateDeviceLabelsResponse |  **Parameters**:  | Name | Description | Required | Type | | --- | --- | --- | --- | | id | ID of the device to be updated | true | string | | | |

## Device Schemas

This API uses the following schemas.

#### Label

|  |  |
| --- | --- |
| **Schema:** devicev202308beta1Label | **Type:** object |
| **Properties:**  | Name | Value | | --- | --- | | id | type:  string  description: Label ID | | name | type: string  description: Label name | | description | type: string  description: Label description | | edate | type: string  format: date-time  description: Label end date (UTC) | | cdate | type: string  format: date-time  description: Label creation date (UTC) | | userId | type: string  description: User ID | | companyId | type: string  description: Company ID | | color | type: string  description: Label color | | order | type: string  description: Label order | | pivotDeviceId | type: string  description: Pivot device ID | | pivotLabelId | type: string  description: Pivot label ID | | |

#### protobufAny

|  |  |
| --- | --- |
| **Schema:** protobufAny | **Type:** object |
| **Properties:**  | Name | Value | | --- | --- | | typeUrl | type: string | | value | type: string  format: byte | | |

#### rpcStatus

|  |  |
| --- | --- |
| **Schema:** rpcStatus | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | code | type: integer  format: int32 | | message | type: string | | details | type: array  items: $ref: protobufAny | | |

#### CreateDeviceRequest

|  |  |
| --- | --- |
| **Schema:** v202308beta1CreateDeviceRequest | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | device | $ref: v202308beta1DeviceConcise | | |

#### CreateDeviceResponse

|  |  |
| --- | --- |
| **Schema:** v202308beta1CreateDeviceResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | device | $ref: v202308beta1DeviceDetailed | | |

#### CreateDevicesRequest

|  |  |
| --- | --- |
| **Schema:** v202308beta1CreateDevicesRequest | **Type:** object |
| **Properties (\*  =  required)**  | Name | Value | | --- | --- | | devices \* | type: array  items: $ref: v202308beta1DeviceConcise  description: List of configurations of devices to be created | | |

#### CreateDevicesResponse

|  |  |
| --- | --- |
| **Schema:** v202308beta1CreateDevicesResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | devices | type: array  items: $ref: v202308beta1DeviceDetailed  description: List of configurations of newly created devices | | failedDevices | type: array  items: type: string  description: List of names of devices that failed to be created | | |

#### CustomColumnData

|  |  |
| --- | --- |
| **Schema:** v202308beta1CustomColumnData | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | deviceId | type: string  description: Device ID | | fieldId | type: string  description: Field ID | | colName | type: string  description: Column name | | description | type: string  description: Description | | colType | type: string  description: Column type | | deviceType | type: string  description: Device type | | |

#### DeleteDeviceResponse

|  |  |
| --- | --- |
| **Schema:** v202308beta1DeleteDeviceResponse  **Properties**: None. | **Type:** object |

#### DeleteDevicesRequest

|  |  |
| --- | --- |
| **Schema:** v202308beta1DeleteDevicesRequest | **Type:** object |
| **Properties** **(\*  =  required)**  | Name | Value | | --- | --- | | ids \* | type: array  items: type: string  description: List of IDs of devices to be deleted | | |

#### DeleteDevicesResponse

|  |  |
| --- | --- |
| **Schema:** v202308beta1DeleteDevicesResponse | **Type:** object |
| **Properties:**  | Name | Value | | --- | --- | | failedDevices | type: array  items: type: string  description: List of IDs of devices that failed to be deleted | | |

#### DeviceConcise

|  |  |
| --- | --- |
| **Schema:** v202308beta1DeviceConcise | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | id | type: string  description: System generated unique identifier  title: id | | deviceName | type: string  description: Device name (device\_name) - The name of the device. Valid characters: alphanumeric and underscores. Length: min=4, max=60. | | deviceSubtype | type: string  description: Device subtype (device\_subtype) - The device subtype. | | cdnAttr | type: string  description: CDN attributes (cdn\_attr) - If this is a DNS server, you can contribute its queries to Kentik's CDN attribution database. Valid values: "Y" or "N". \*\* cdn\_attr is required when the device subtype's parent type is "host-nprobe-dns-www" | | deviceDescription | type: string  description: Description (device\_description) - The device description. Valid characters: any. Length: max=128. | | sendingIps | type: array  items: type: string  description: Device ip (sending\_ips) - Array containing one or more IP address(es), from which the device is sending flow. \*\* sending\_ips is required when the device subtype's parent type is "router" | | deviceSampleRate | type: number  format: double  description: Sample rate (device\_sample\_rate) - The rate at which the device is sampling flows. Valid values: integer bigger than 1. Recommended rate varies depending on flow volume; see https://kb.kentik.com/Ab02.htm#Ab02-Flow\_Sampling. \*\* device\_sample\_rate is required when the device subtype's parent type is "router" | | planId | type: integer  format: int64  description: Plan (plan\_id) - The ID of the plan to which this device is assigned. Available plan(s) can be found via the Plans API. Valid value: integer. | | siteId | type: integer  format: int64  description: Site (site\_id) - The ID of the site (if any) to which this device is assigned. Site IDs are system generated when a site is created. Valid value: integer. | | minimizeSnmp | type: boolean  description: SNMP polling (minimize\_snmp) - The interval at which SNMP will be polled. If "false" (Standard), interface counter will be polled every 10 minutes and interface description every 30 minutes. If "true" (Minimum) (Minimum), interface counter won't be polled and interface description will be polled every 6 hours. \*\* minimize\_snmp is required when the device subtype's parent type is "router" | | deviceSnmpIp | type: string  description: Device SNMP IP (device\_snmp\_ip) - The SNMP IP to use when polling the device. device\_snmp\_ip is ignored unless the device subtype's parent type is "router" | | deviceSnmpCommunity | type: string  description: SNMP community (device\_snmp\_community) - The SNMP community to use when polling the device. device\_snmp\_community is ignored unless the device subtype's parent type is "router" | | deviceSnmpV3Conf | $ref: v202308beta1DeviceSnmpV3Conf | | deviceBgpType | type: string  description: BGP (device\_bgp\_type) - Device bgp type. Valid values: "none" (use generic IP/ASN mapping), "device" (peer with the device itself), "other\_device" (share routing table of existing peered device) | | deviceBgpNeighborIp | type: string  description: Your IPv4 peering address (device\_bgp\_neighbor\_ip) - A valid IPv4 address to use for peering with the device. \*\* An IPv4 and/or IPv6 peering address is required when device\_bgp\_type is set to "device" | | deviceBgpNeighborIp6 | type: string  description: Your IPv6 peering address (device\_bgp\_neighbor\_ip6) - A valid IPv6 address to use for peering with the device. \*\* An IPv4 and/or IPv6 peering address is required when device\_bgp\_type is set to "device" | | deviceBgpNeighborAsn | type: string  description: Your ASN (device\_bgp\_neighbor\_asn) - The valid AS number (ASN) of the autonomous system that this device belongs to. \*\* device\_bgp\_neighbor\_asn is required when device\_bgp\_type is set to "device" | | deviceBgpPassword | type: string  description: BGP md5 password (device\_bgp\_password) - Optional BGP MD5 password (shared authentication password for BGP peering). Valid characters: alphanumeric. Length: 32. device\_bgp\_password is optional when device\_bgp\_type is set to "device" | | useBgpDeviceId | type: integer  format: int64  description: Select master BGP device (use\_bgp\_device\_id) - The ID of the device whose BGP table should be shared with this device. \*\* use\_bgp\_device\_id is required when device\_bgp\_type is set to "other\_device"). Valid value: a system-generated device\_id | | deviceBgpFlowspec | type: boolean  description: BGP Flowspec Compatibility (device\_bgp\_flowspec) - Toggle BGP Flowspec Compatibility for device. | | nms | $ref: v202308beta1DeviceNmsConfig | | deviceBgpCredentialName | type: string  description: BGP Credential Name (device\_bgp\_credential\_name) - Optional Credential Name (Credential for BGP peering). Valid characters: alphanumeric. Length: 32. | | flowSnmpCredentialName | type: string  description: Snmp Credential Name (flow\_snmp\_credential\_name) - Optional Credential Name (Credential for Flow Snmp peering). Valid characters: alphanumeric. Length: 32. | | |

#### DeviceDetailed

|  |  |
| --- | --- |
| **Schema:** v202308beta1DeviceDetailed | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | id | type: string  description: System generated unique identifier  readOnly: true | | companyId | type: string  description: Company ID | | deviceName | type: string  description: Device name | | deviceAlias | type: string  description: Device alias | | deviceType | type: string  description: Device type | | deviceDescription | type: string  description: Device description | | site | $ref: v202308beta1Site | | plan | $ref: v202308beta1Plan | | labels | type: array  items: $ref: devicev202308beta1Label  description: List of labels | | allInterfaces | type: array  items: $ref: v202308beta1Interface  description: List of interfaces | | deviceFlowType | type: string  description: Device flow type | | deviceSampleRate | type: string  description: Device sample rate | | sendingIps | type: array  items: type: string  description: List of sending IPs | | deviceSnmpIp | type: string  description: Device SNMP IP | | deviceSnmpCommunity | type: string  description: Device SNMP community  title: Keeping these tokens so we can give the user a useful error. They are removed from DeviceRequest | | minimizeSnmp | type: boolean  description: Minimize SNMP | | deviceBgpType | type: string  description: Device BGP type | | deviceBgpNeighborIp | type: string  description: Device BGP neighbor IP | | deviceBgpNeighborIp6 | type: string  description: Device BGP neighbor IP6 | | deviceBgpNeighborAsn | type: string  description: Device BGP neighbor ASN | | deviceBgpFlowspec | type: boolean  description: Device BGP flowspec | | deviceBgpPassword | type: string  description: Device BGP password  title: Keeping these tokens so we can give the user a useful error. They are removed from DeviceRequest | | deviceBgpLabelUnicast | type: boolean  description: Device BGP label unicast | | bgpLookupStrategy | type: string  description: BGP lookup strategy | | deviceStatus | type: string  description: Device status | | useBgpDeviceId | type: string  description: Use BGP device ID | | customColumns | type: string  description: Custom columns | | customColumnData | type: array  items: $ref: v202308beta1CustomColumnData  description: Custom column data | | deviceChfClientPort | type: string  description: Device CHF client port | | deviceChfClientProtocol | type: string  description: Device CHF client protocol | | deviceChfInterface | type: string  description: Device CHF interface | | deviceAgentType | type: string  description: Device agent type | | maxFlowRate | type: integer  format: int64  description: Max flow rate | | maxBigFlowRate | type: integer  format: int64  description: Max big flow rate | | deviceProxyBgp | type: string  description: Device proxy BGP | | deviceProxyBgp6 | type: string  description: Device proxy BGP6 | | createdDate | type: string  format: date-time  description: Creation timestamp (UTC) | | updatedDate | type: string  format: date-time  description: Last modification timestamp (UTC) | | deviceSnmpV3ConfEnabled | type: boolean  description: Device SNMP v3 configuration enabled | | deviceSnmpV3Conf | $ref: v202308beta1DeviceSnmpV3Conf | | cdnAttr | type: string  description: CDN attributes | | bgpPeerIp4 | type: string  description: BGP peer IP4 | | bgpPeerIp6 | type: string  description: BGP peer IP6 | | deviceSubtype | type: string  description: Device subtype | | deviceVendorType | type: string  description: Device vendor type | | deviceModelType | type: string  description: Device model type | | cloudExportId | type: string  description: Cloud export ID | | deviceKproxy | type: string  description: Device KProxy | | snmpEnabled | type: string  description: SNMP enabled | | snmpDisabledReason | type: string  description: SNMP disabled reason | | snmpDisabledReasonOther | type: string  description: SNMP disabled reason other | | bgpDisabledReason | type: string  description: BGP disabled reason | | bgpDisabledReasonOther | type: string  description: BGP disabled reason other | | deviceManufacturer | type: string  description: Device manufacturer | | deviceAlert | type: string  description: Device alert | | role | type: string  description: Role | | deviceGnmiV1Conf | $ref: v202308beta1GnmiV1Conf | | useAsnFromFlow | type: boolean  description: Use ASN from flow | | maxInterface | type: integer  format: int64  description: Max interface | | maxInterfaceCheck | type: integer  format: int64  description: Max interface check | | nms | $ref: v202308beta1DeviceNmsConfig | | deviceBgpCredentialName | type: string  description: BGP Credential Name (device\_bgp\_credential\_name) - Optional Credential Name (Credential for BGP peering). Valid characters: alphanumeric. Length: 32. | | flowSnmpCredentialName | type: string  description: Snmp Credential Name (flow\_snmp\_credential\_name) - Optional Credential Name (Credential for Flow Snmp peering). Valid characters: alphanumeric. Length: 32. | | |

#### DeviceNmsConfig

|  |  |
| --- | --- |
| **Schema:** v202308beta1DeviceNmsConfig | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | agentId | type: string  description: ID of the agent that is monitoring this device. | | ipAddress | type: string  description: Local IP address of this device. | | snmp | $ref: v202308beta1DeviceNmsSnmpConfig | | st | $ref: v202308beta1DeviceNmsStConfig | | |

#### DeviceNmsSnmpConfig

|  |  |
| --- | --- |
| **Schema:** v202308beta1DeviceNmsSnmpConfig | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | credentialName | type: string  description: Name of the SNMP credentials from the credential vault. | | port | type: integer  format: int64  description: SNMP port, to override default of 161. | | timeout | type: string  description: Timeout, to override default of 2s. | | |

#### DeviceNmsStConfig

|  |  |
| --- | --- |
| **Schema:** v202308beta1DeviceNmsStConfig | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | credentialName | type: string  description: Name of the ST credentials from the credential vault. | | port | type: integer  format: int64  description: ST port, to override default of 6030. | | timeout | type: string  description: Timeout, to override default of 2s. | | secure | type: boolean  description: Use SSL to connect to this device. | | |

#### DeviceSnmpV3Conf

|  |  |
| --- | --- |
| **Schema:** v202308beta1DeviceSnmpV3Conf | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | username | type: string  description: UserName (username) - the user name to use to authenticate via SNMP v3. \*\* UserName is required when device\_snmp\_v3\_conf is not null | | authenticationProtocol | type: string  description: Authentication Protocol (authentication\_protocol) - the auth protocol to use via SNMP v3 | | authenticationPassphrase | type: string  description: Authentication Passphrase (authentication\_passphrase) - the passphrase to use for SNMP v3 authentication protocol (required when AuthenticationProtocol not NoAuth. | | privacyProtocol | type: string  description: Privacy Protocol (privacy\_protocol) - the privacy protocol to use to authenticate via SNMP v3 | | privacyPassphrase | type: string  description: Privacy Passphrase (privacy\_passphrase) - the passphrase to use for SNMP v3 privacy protocol (required when PrivacyProtocol not NoPriv) | | |

#### GetDeviceResponse

|  |  |
| --- | --- |
| **Schema:** v202308beta1GetDeviceResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | device | $ref: v202308beta1DeviceDetailed | | |

#### GnmiV1Conf

|  |  |
| --- | --- |
| **Schema:** v202308beta1GnmiV1Conf | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | dialoutServer | type: string  description: Dialout server | | |

#### Interface

|  |  |
| --- | --- |
| **Schema:** v202308beta1Interface | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | interfaceDescription | type: string  description: Interface description | | initialSnmpSpeed | type: string  description: Initial SNMP speed | | deviceId | type: string  description: Device ID | | snmpSpeed | type: string  description: SNMP speed | | snmpAlias | type: string  description: SNMP alias | | snmpId | type: string  description: SNMP ID | | connectivityType | type: string  description: Connectivity type | | networkBoundary | type: string  description: Network boundary | | provider | type: string  description: Provider | | |

#### LabelConcise

|  |  |
| --- | --- |
| **Schema:** v202308beta1LabelConcise | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | id | type: integer  format: int64  description: Label ID | | |

#### ListDevicesResponse

|  |  |
| --- | --- |
| **Schema:** v202308beta1ListDevicesResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | devices | type: array  items: $ref: v202308beta1DeviceDetailed  description: List of configurations of devices | | invalidCount | type: integer  format: int64  description: Number of invalid entries encountered while collecting dat | | |

#### Plan

|  |  |
| --- | --- |
| **Schema:** v202308beta1Plan | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | id | type: string  description: Plan ID | | name | type: string  description: Plan name | | |

#### Site

|  |  |
| --- | --- |
| **Schema:** v202308beta1Site | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | id | type: string  description: Site ID | | siteName | type: string  description: Site name | | lat | type: number  format: double  description: Site latitude | | lon | type: number  format: double  description: Site longitude | | companyId | type: string  description: Company ID | | |

#### UpdateDeviceLabelsRequest

|  |  |
| --- | --- |
| **Schema:** v202308beta1UpdateDeviceLabelsRequest | **Type:** object |
| **Properties** **(\*  =  required)**  | Name | Value | | --- | --- | | id \* | type: string  description: ID of the device to be updated | | labels \* | type: array  items: $ref: v202308beta1LabelConcise  description: List of labels to be added to the device | | |

#### UpdateDeviceLabelsResponse

|  |  |
| --- | --- |
| **Schema:** v202308beta1UpdateDeviceLabelsResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | device | $ref: v202308beta1DeviceDetailed | | |

#### UpdateDeviceRequest

|  |  |
| --- | --- |
| **Schema:** v202308beta1UpdateDeviceRequest | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | device | $ref: v202308beta1DeviceConcise | | |

#### UpdateDeviceResponse

|  |  |
| --- | --- |
| **Schema:** v202308beta1UpdateDeviceResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | device | $ref: v202308beta1DeviceDetailed | | |

#### UpdateDevicesRequest

|  |  |
| --- | --- |
| **Schema:** v202308beta1UpdateDevicesRequest | **Type:** object |
| **Properties (\*  =  required)**  | Name | Value | | --- | --- | | devices \* | type: array  items: $ref: v202308beta1DeviceConcise  description: List of configurations of devices to be updated | | |

#### UpdateDevicesResponse

|  |  |
| --- | --- |
| **Schema:** v202308beta1UpdateDevicesResponse | **Type:** object |
| **Properties**:  | Name | Value | | --- | --- | | devices | type: array  items: $ref: v202308beta1DeviceDetailed  description: List of configurations of updated devices | | failedDevices | type: array  items: type: string  description: List of IDs of devices that failed to be updated | | |
