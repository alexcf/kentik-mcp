# Using Threat Feeds

Source: https://kb.kentik.com/docs/using-threat-feeds

---

This article introduces using **Threat Feeds** in Kentik to reveal compromised hosts.

## Threat Feed Dimensions

Kentik enhances flow records with IP reputation data from Spamhaus to help identify traffic from infected or compromised hosts. This enrichment allows for the use of two key dimensions for grouping or filtering:

* **Bot Net CC**: Identifies IPs associated with botnet command and control (CC) servers. Monitoring communication with these IPs helps detect potentially infected hosts or subscribers involved in botnet activities.
* **Threat List Host**: Covers a wider range of potentially malicious IPs, such as those used for malware distribution, phishing, and spam. This dimension helps identify problematic IPs within a network and client IPs interacting with known threats.

## Query for Bots

As a query example for Bot Net CC, let’s imagine a fictional enterprise called Pear, Inc. who want to detect if they have any compromised hosts on their network communicating with a C&C server, thereby participating in a botnet. ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(214).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A16Z&se=2026-05-12T09%3A30%3A16Z&sr=c&sp=r&sig=AlMxWvPl7I%2F9rnikjYtJCVX4%2F7YDY8UbXdKV1%2BcgncM%3D)

To check this, we can build a query in Data Explorer. Using the **Group By Dimensions** selector in the **Query** pane, we can look at two dimensions that, when combined, give the list of compromised hosts on our network: Source IP/CIDR and Destination Bot Net CC.

We’ll set the **Time** pane (see **Time Pane**) to look back at the last hour, and in the **Devices** pane (**Data Sources Pane**) we choose All to look at our entire network.

We’ll set the **Filters** pane (see **Filters Pane**) to ignore normal network traffic and look only at botnet traffic. We do this by filtering out traffic whose botnet value is blank:

* Click anywhere in the pane to open the **Filtering Options** dialog.![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(215).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A16Z&se=2026-05-12T09%3A30%3A16Z&sr=c&sp=r&sig=AlMxWvPl7I%2F9rnikjYtJCVX4%2F7YDY8UbXdKV1%2BcgncM%3D)
* In the Ad-Hoc Filter Groups pane, click **Add Filter Group**.
* Change Include to Exclude.
* Change the dimension to Destination Bot Net CC.
* Click **Save Changes**.

Once the query configuration is complete, click the Run Query button at the top of the sidebar.

This query returns a graph showing four infected hosts. Based on these results, the staff at Pear, Inc. would know which hosts to examine for botnet malware removal.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(216).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A16Z&se=2026-05-12T09%3A30%3A16Z&sr=c&sp=r&sig=AlMxWvPl7I%2F9rnikjYtJCVX4%2F7YDY8UbXdKV1%2BcgncM%3D)

## Set Up a Threat Report

Pear, Inc. could also automate this query as a scheduled report emailed to specified users. Follow these steps to save the query to a new dashboard and set up a weekly report:![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(217).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A16Z&se=2026-05-12T09%3A30%3A16Z&sr=c&sp=r&sig=AlMxWvPl7I%2F9rnikjYtJCVX4%2F7YDY8UbXdKV1%2BcgncM%3D)

* In Data Explorer, from the drop-down **View Options** menu, select **Add to Dashboard** **»** **New Dashboard**.
* Enter “Bot Traffic” in the **Name** field and click **Add Dashboard**.
* Keep the default panel settings and click **Add Panel** to create the “Bot Traffic” dashboard.
* Go to **Admin** in the main navbar, then select **Subscriptions** from the sidebar.
* Click **Add Subscription**, fill in the settings, and click **Add Subscription**.

By following these steps, Pear, Inc. will receive an email every Monday listing infected hosts from the past seven days.

> **Note**: Reports are based on dashboards (see **Subscriptions**).

## Alert on Compromised Hosts

To set up automatic notifications for compromised hosts, Pear, Inc. could set up an alert policy in Kentik’s alerting system and get a Slack notification (for example) any time a new host starts talking to a C&C server.

To do this, follow these steps:

1. **Navigate to Alerting**: From the main portal navbar, select **Alerting.**
2. **Access Policies**: Click on **Policies** in the sidebar.
3. **Add a New Policy**: On the **Policies** page, click **Add Policy**.
4. **Configure Alert Policy**: In the **Add Alert Policy** dialog, configure the necessary settings across the tabs:

   1. **General Settings Tab**
   2. **Dataset Tab**
   3. **Alert Thresholds Tab**
   4. **Historical Baseline Tab**

#### General Settings Tab

1. Enter a policy name and select the appropriate **Policy Dashboard** (e.g., Bot Traffic*)*.
2. Leave **Silent Mode** enabled to allow the alerting system to get a baseline before activating any alerts.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(218).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A16Z&se=2026-05-12T09%3A30%3A16Z&sr=c&sp=r&sig=AlMxWvPl7I%2F9rnikjYtJCVX4%2F7YDY8UbXdKV1%2BcgncM%3D)

#### Dataset Tab

1. In the **Data Funneling** pane, set **Devices** to All.
2. Click anywhere in the **Dimensions** selector to add the dimensions from your original query: Source IP/CIDR and Destination Bot Net CC.
3. Set the **Secondary Metrics** to Bits/second.
4. Create a filter to exclude traffic whose botnet value is blank.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(219).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A16Z&se=2026-05-12T09%3A30%3A16Z&sr=c&sp=r&sig=AlMxWvPl7I%2F9rnikjYtJCVX4%2F7YDY8UbXdKV1%2BcgncM%3D)

1. In the **Building Your Dataset** pane, set the **Maximum Number of Keys Analyzed Per Evaluation** to 100.
2. Leave all other settings as default so that the system auto-calculates the **Minimum Traffic Threshold**.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(220).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A16Z&se=2026-05-12T09%3A30%3A16Z&sr=c&sp=r&sig=AlMxWvPl7I%2F9rnikjYtJCVX4%2F7YDY8UbXdKV1%2BcgncM%3D)

#### Alert Thresholds Tab

1. Select a criticality level (e.g., Major), then set the switch to Enabled to enable the threshold.
2. In the **Threshold Configuration** pane, select **Check If Certain Conditions Exist** from the **This Threshold Will** menu. Leave the other settings at their defaults.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(221).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A16Z&se=2026-05-12T09%3A30%3A16Z&sr=c&sp=r&sig=AlMxWvPl7I%2F9rnikjYtJCVX4%2F7YDY8UbXdKV1%2BcgncM%3D)

1. In the **Conditions** pane, set a threshold of **Greater Than 1 kpps**. This triggers alerts for traffic greater than 1000 packets per second to a Botnet C&C server.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(222).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A16Z&se=2026-05-12T09%3A30%3A16Z&sr=c&sp=r&sig=AlMxWvPl7I%2F9rnikjYtJCVX4%2F7YDY8UbXdKV1%2BcgncM%3D)

1. From the **Notifications** pane, in the **Send Alert Status Updates To** selector, choose “Slack itsec’. Ensure the channel is pre-configured under **Alerting »** **Channels** (see **Notifications**).

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(223).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A16Z&se=2026-05-12T09%3A30%3A16Z&sr=c&sp=r&sig=AlMxWvPl7I%2F9rnikjYtJCVX4%2F7YDY8UbXdKV1%2BcgncM%3D)

#### Historical Baseline Tab

1. On the Historical Baseline tab, leave the default settings. Once all of the required fields have been specified, there will be a green check mark on all of the tabs.
2. Click **Add Alert Policy** to start the policy in Silent Mode, studying the normal patterns of traffic.

After the Silent Mode end date (see **General Policy Settings**), a Slack notification gets sent every time a new host starts talking to a botnet C&C server.

Kentik’s Threat Feeds feature provides powerful insights into compromised hosts on your network, offering visibility that may have been previously unattainable. This capability supports both proactive monitoring and data forensics, enhancing the protection of your company’s infrastructure.

For assistance setting up or using this feature, contact the Kentik support team at **support@kentik.com**.
