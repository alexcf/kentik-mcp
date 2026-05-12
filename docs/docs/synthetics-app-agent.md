# Synthetics Agent

Source: https://kb.kentik.com/docs/synthetics-app-agent

---

This article discusses `ksynth-agent`, Kentik's software agent for synthetic testing.

> **Notes:**
>
> * For a high-level overview of agent types and deployment types for synthetics, see **Kentik Synthetics Agent**.
> * For a general explanation of synthetic testing in Kentik, see **Synthetics**.
> * For assistance with any aspect of the agent setup process, contact Kentik (see **Customer Care**).

## Agent Versions

Installation packages for Kentik’s synthetics agent, `ksynth-agent`, are available for the following common Linux distributions:

| Debian | Ubuntu | RedHat Enterprise Linux | CentOS |
| --- | --- | --- | --- |
| Buster (10.0) | Focal Fossa (20.04 LTS) | 8.0 | 8 (Stream) |

> **Note:** For versions not listed above, contact Kentik (see **Customer Care**).

Downloads are available from the following locations:

* **PackageCloud**:

  + Debian-based distributions: APT package manager.
  + RPM based distributions: YUM package manager.
* **Docker hub**

## Hardware Requirements

The following resources must be available (at minimum) to each instance of the `ksynth-agent` agent:

* **RAM allocation**: 2GB
* **CPU cores**: 2

> **Note:** For additional considerations for deploying `ksynth-agent`, see **Deployment Considerations**.

## Agent Command Line

The `ksynth-agent` agent uses a standard UNIX command line interface (CLI), as follows:

```
ksynth-agent [options]
```

### Agent CLI Arguments

To get a list of all CLI options and their descriptions, execute this command on any `ksynth-agent` host machine.

```
ksynth-agent --help
```

#### Listener CLI Arguments

These three CLI arguments are required for `ksynth-agent` when creating **agent-to-agent** synthetics tests. They configure listeners on the target agent to respond to ping requests:

* `--tcp-ping-port`

  + Opens TCP listener(s) to respond to TCP SYN ping requests (tcp-syn-sport)
  + Accepts multiple port values

    **Examples**:

    - `--tcp-ping-port 8877`
    - `--tcp-ping-port 8877 8878`
* `--udp-ping-port`

  + Opens UDP listener(s) to respond to UDP ping requests (udp-dport)
  + Accepts port value(s)

    **Example**: `--udp-ping-port 9977`
* `--udp-listen-ipv6`

  + Enables UDP listener on IPv6 addresses for the UDP ping port
  + Flag only (no value needed)
  + Required if you want to support IPv6 UDP pings

**Example 1: Run with listener arguments**

```
ksynth-agent --company-id 175913 --name my-agent \
  --tcp-ping-port 8877 \
  --udp-ping-port 9977 \
  --udp-listen-ipv6
```

**Example 2: Run with listener arguments (Docker)**

```
docker run -d --name ksynth-agent \
  --network host \
  -v /var/lib/ksynth-agent:/var/lib/ksynth-agent \
  kentik/ksynth-agent:latest \
  ksynth-agent --company-id <id> --name <name> \
  --tcp-ping-port 8877 --udp-ping-port 9977 --udp-listen-ipv6
```

> **Tip**: Common port values used are TCP 8877 and UDP 9977, but you can specify any available ports.

## About Agent Setup

Setup of the `ksynth-agent` agent involves two main stages:

1. **Installation**: Install the agent in your network infrastructure or public cloud instance using either the **Agent Package Installation**or **Agent Docker Installation**.
2. **Final Step: Register and Activate the Agent**: Register and activate the instance in the Kentik portal (see **Activate Agent**). Instructions depend on which of the **Registration Methods** you used.

> **Notes:**
>
> * If `ksynth-agent` has previously been installed on the same machine, you must delete the existing `ksynth.id` file before reinstalling.
> * Before installing, refer to **Hardware Requirements** and **ksynth Deployment Considerations**.

### Registration Methods

Kentik uses your organization's Company ID to associate the `ksynth-agent` with your account. You can provide this ID using one of two methods, which determines the registration workflow in the portal:

* **Challenge Code Method**: During installation, the agent starts without a Company ID and generates a unique challenge code in its logs. You must manually enter this code in the Kentik portal’s **Agents** page to both register and activate the agent.
* **Company ID Method**: Pass your Company ID (**Account#** from the **Licenses Page**) to the agent during installation in a command-line parameter. The agent automatically registers with your account but must still be manually activated on the Kentik portal **Agents** page.

  > **Tip**: This method is recommended for bulk or automated deployments because it eliminates the need to capture and enter individual challenge codes.

## Agent Package Installation

Installation for both the Debian-based and RPM-based `ksynth-agent` packages are covered in the following topics.

> **Note:** In case of errors, see **Linux Installation Troubleshooting**.

To install `ksynth-agent` via package:

1. **Install Package Repository**: On the host machine, run the following command to install the `ksynth-agent` package repo:

   1. **Debian**:

      ```
      curl -s https://packagecloud.io/install/repositories/kentik/ksynth-agent/script.deb.sh | sudo bash
      ```
   2. **RPM**:

      ```
      curl -s https://packagecloud.io/install/repositories/kentik/ksynth-agent/script.rpm.sh | sudo bash
      ```
2. **Install Agent Service**: Use the package manager to install the `ksynth-agent` service:

   1. **Debian**:

      ```
      apt-get install ksynth-agent
      ```
   2. **RPM**:

      ```
      yum install ksynth-agent
      ```

      > **Notes:**
      >
      > * Depending on the Linux distro/version, the agent may notify you of additional libraries that you need to install.
      > * By default, the agent runs in Network Agent mode. To optionally add web test functionality, see **Add Web Test Functionality**.
3. **Configure Company ID and Optional Features**: Before starting the agent, configure required and optional parameters in the agent's configuration file on the host machine (`/etc/default/ksynth-agent`).

   * **Required: Choose an Activation Method**: Choose one of the following methods:

     + **Challenge Code Method:** Leave `KENTIK_COMPANY=` empty (the default) in the config file. The agent will generate a challenge code when started.
     + **Company ID Method:**

       - From the Kentik portal’s main nav bar, go to Organization Settings » **Licenses****.**
       - Copy your Kentik **Account #**, which is yourCompany ID.
       - Paste the Company ID after `KENTIK_COMPANY=` in the `/etc/default/ksynth-agent` file on the host machine.
   * **Optional/Conditional Features**: Update the `/etc/default/ksynth-agent` file on the host machine as necessary/desired:

     + **EU Accounts**: Required if your account is on the EU cluster:

       - Change `KENTIK_REGION=US` to `KENTIK_REGION=EU`.
     + **SSL-DPI**: Required if you use SSL DPI on your network:

       - Add the `NODE_EXTRA_CA_CERTS` environment variable to the file as shown below. This allows the agent to trust the device performing the SSL description:

         ```
         export NODE_EXTRA_CA_CERTS=[your CA certificate file path]
         ```
     + **Web Tests** (optional): To enable web tests, follow the instructions in **Add Web Test Functionality**.

4. **Start the** `ksynth-agent` **Service**: Execute the following command to start the service:

   ```
   systemctl start ksynth-agent
   ```
5. **Register/Activate the Agent**: Follow the procedures in **Final Step: Register and Activate the Agent****.**

> **Note**: If you update environment variables while `ksynth-agent` is running, you need to restart the agent for the changes to become active.

#### Add Web Test Functionality

To enable web test functionality when installing `ksynth-agent`, include the following procedure as step 4 of the process outlined above:

1. Open the service configuration file in your editor of choice (nano is used in the following example):

   ```
   nano /etc/default/ksynth-agent
   ```
2. Remove or comment-out `K_AGENT_DISABLE=browser` and save the file.
3. Navigate to the scripts folder in the agent's working directory:

   ```
   cd /var/lib/ksynth-agent/scripts
   ```
4. Run the appropriate configuration script and ensure successful completion:

   1. **Debian-based packages**:

      ```
      ./install-browser-deps_debian
      ```
   2. **RPM-based packages**:

      ```
      ./install-browser-deps_centos
      ```

### Linux Installation Troubleshooting

During package installation, if you receive an error message about missing dependencies, run the commands below to install the commonly required dependencies.

* **Debian-based packages**:

  ```
  sudo apt-get install ca-certificates fonts-liberation libappindicator3-1 libasound2 libatk-bridge2.0-0 libatk1.0-0 libc6 libcairo2 libcups2 libdbus-1-3 libexpat1 libfontconfig1 libgbm1 libgcc1 libglib2.0-0 libgtk-3-0 libnspr4 libnss3 libpango-1.0-0 libpangocairo-1.0-0 libstdc++6 libx11-6 libx11-xcb1 libxcb1 libxcomposite1 libxcursor1 libxdamage1 libxext6 libxfixes3 libxi6 libxrandr2 libxrender1 libxss1 libxtst6 lsb-release wget xdg-utils
  ```
* **RPM-based packages**:

  ```
  sudo yum install alsa-lib.x86_64 atk.x86_64 cups-libs.x86_64 gtk3.x86_64 ipa-gothic-fonts libXcomposite.x86_64 libXcursor.x86_64 libXdamage.x86_64 libXext.x86_64 libXi.x86_64 libXrandr.x86_64 libXScrnSaver.x86_64 libXtst.x86_64 pango.x86_64 xorg-x11-fonts-100dpi xorg-x11-fonts-75dpi xorg-x11-fonts-cyrillic xorg-x11-fonts-misc xorg-x11-fonts-Type1 xorg-x11-utils
  ```
* **RHEL8 RPM-based packages:**

  ```
  sudo yum install -y alsa-lib.x86_64 atk.x86_64 cups-libs.x86_64 gtk3.x86_64 libdrm.x86_64 mesa-libgbm.x86_64 libXcomposite.x86_64 libXcursor.x86_64 libXdamage.x86_64 libXext.x86_64 libXi.x86_64 libXrandr.x86_64 libXScrnSaver.x86_64 libXtst.x86_64 pango.x86_64 xorg-x11-fonts-100dpi xorg-x11-fonts-75dpi xorg-x11-fonts-cyrillic xorg-x11-fonts-misc xorg-x11-fonts-Type1 xorg-x11-utils
  ```

After installing dependencies, update the `nss` library using this command (for Debian, use `apt` instead of `yum`):

```
yum update nss -y
```

> **Notes:**
>
> * The dependencies referenced above are subject to change. For the latest information, check the lists at **Chrome doesn't launch on UNIX**.
> * If you continue to experience installation errors after trying the above, contact Kentik (see **Customer Care**).

## Agent Docker Installation

Installation of `ksynth-agent` from a downloaded Docker image provides a convenient and easy deployment mechanism for containerized environments.

To install and run ksynth-agent via Docker, complete the following steps:

1. **Prepare a Volume**: Create a local directory on the host machine to mount as a persistent volume. This is required so the agent's identity (`ksynth.id`) persists across container upgrades or restarts.
2. **Choose Your Activation Command**: Run the appropriate docker run command below based on your chosen **Registration Method**:

   1. **Challenge Code** (for manual, one-off activation):

      ```
      docker run -it --rm -v /path/to/local/directory:/var/lib/ksynth-agent kentik/ksynth-agent:latest ksynth-agent
      ```
   2. **Company ID** (for bulk/scripted deployment):

      ```
      docker run -it --rm -v /path/to/local/directory:/var/lib/ksynth-agent kentik/ksynth-agent:latest ksynth-agent -c <your_company_id>
      ```

> **Tips:**
>
> * Kentik recommends the default installation method unless you need to do bulk or automated installations.
> * In the above commands, update the placeholders `/path/to/local/directory` and `<your_company_id>` with your actual values.

3. **Registration/Activation:** Follow the procedures in **Final Step: Register and Activate the Agent** based on the method you chose above.

### **Notes on** `docker run`

* The Docker run `-it` argument is a combination of two flags:

  + `-i` stands for interactive and allows you to interact with the container's command line.
  + `-t` stands for terminal which tells Docker to allocate a pseudo terminal for the container.
* The `-rm` argument (remove) tells Docker to automatically remove the container and its file system when the container exits.
* The `-v` argument (volume) will cause a `ksynth` directory to be created on the server if it doesn't already exist and mounted in the Docker container at `/var/lib/ksynth`. When `ksynth-agent` runs, the `ksynth.id` ID file will, by default, be written to this location, where it will persist between container upgrades (otherwise every Docker image upgrade will prompt an attempt to register a new agent).
* The `kentik/ksynth-agent:latest` statement will run `ksynth-agent`, passing to it the agent subcommand arguments that follow (`-c`, etc.).
* The value of the `-c` argument is your organization's Kentik company ID, which is the number next to **Account #** on the portal's Licenses page.

## Final Step: Register and Activate the Agent

Registering an instance of `ksynth-agent` involves securely associating the agent with your Kentik account, which you'll do on the **Agents** page in the Synthetics section of the portal. As noted in **Registration Methods**, your choice of either the Challenge Code or Company ID method determines your next action:

| **Registration Method** | Installation Status | Next Action |
| --- | --- | --- |
| Challenge Code | Pending Registration | **Register Agent** (Requires challenge code from logs) |
| Company ID | Pending Activation | **Activate Agent** (Manual process in portal) |

### Register Agent

If you installed `ksynth-agent` using the Challenge Code method (see **Registration Methods**), complete the following steps to register it:

1. On the `ksynth-agent` host, run one or the other of the following commands:

   1. **Debian/RPM**:

      ```
      systemctl status ksynth-agent -f or journalctl -u ksynth-agent
      ```
   2. **Docker**:

      ```
      docker logs ksynth-agent
      ```
2. `ksynth-agent` returns a challenge code. Copy your challenge code to the clipboard.

   **Example**:

   ```
   [INFO] authenticate - Activation challenge code: f4dd8632fbdb09714a38c17a7b536a4fe84d973beb4172d01abbba1c77exxxxx
   ```
3. Go to the **Agents** page (Synthetics » **Agents**) in the Kentik portal and click **Enter your Agent challenge code** to open the Register Agent dialog.
4. Paste the code from the earlier step and click **Register** to proceed to the Activate Agent dialog.
5. Review/update settings as necessary and click **Activate**.
6. The new agent instance is now shown with a status of Active in the **Private Agents** tab of the **Agents** page.

> **Note:** To register multiple installed `ksynth-agent` instances, repeat the above steps with the challenge code corresponding to each instance.

### Activate Agent

If you installed `ksynth-agent` using the Company ID method (see **Registration Methods**), complete the following steps to activate it:

1. Go to the **Agents** page (Synthetics » **Agents**) in the Kentik portal and find the newly installed agent on the **Private Agents** tab. The agent's status is Pending at this point.
2. Click the agent to open the Agent Details drawer.
3. Click **Activate** to open the Activate Agent dialog.
4. Review and update settings as necessary and click **Activate** to return to the Agents List.
5. The agent's status is now indicated as Active on the Private Agents tab.
