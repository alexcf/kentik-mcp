# Configuration Examples

Source: https://kb.kentik.com/docs/configuration-examples

---

To allow Kentik’s **AI Advisor** to pull configuration state and run diagnostic commands, the Universal Agent requires read-only SSH access to your network devices.

Below are the standard configuration snippets for creating a local read-only user (e.g., `kentik_ro`) on our supported platforms.

## Juniper Junos

Junos has a built-in `read-only` login class that provides the exact permissions the Kentik Agent needs to view configurations and run operational `show` commands without the ability to alter the system.

```
configure
set system login user kentik_ro class read-only
set system login user kentik_ro authentication plain-text-password
! Enter password when prompted
set system services ssh
commit
```

## Cisco NX-OS (Nexus)

For Nexus devices, the built-in `network-operator` role restricts the user to read-only access, preventing any configuration changes while allowing full visibility into the device state.

```
configure terminal
feature ssh
username kentik_ro password 0 <PASSWORD> role network-operator
exit
copy running-config startup-config
```

## Arista EOS

Similar to NX-OS, Arista EOS utilizes a `network-operator` role. Assigning this to the Kentik service account ensures secure, read-only API and CLI access.

```
configure terminal
management ssh
   server enable
exit
username kentik_ro privilege 1 secret <PASSWORD> role network-operator
exit
write memory
```

## Cisco IOS / IOS-XE

Cisco IOS-XE requires slightly more care. Because the `show running-config` command natively requires Privilege Level 15, the safest local approach is to create a Level 15 user but restrict them to a read-only Parser View.

> **Note**: If you use TACACS+, simply authorize `show` commands and deny `configure terminal`).

```
configure terminal
aaa new-model
! Enable secret is required to use parser views
enable secret <ENABLE_PASSWORD>

! Create a read-only view
parser view kentik-view
 secret 5 <VIEW_PASSWORD>
 commands exec include all show
 ! specifically ensure show run is allowed
 commands exec include show running-config
 exit

! Create the user and assign them to this view
username kentik_ro view kentik-view secret <USER_PASSWORD>
exit
write memory
```
