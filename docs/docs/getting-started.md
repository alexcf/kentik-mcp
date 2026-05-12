# Getting Started

Source: https://kb.kentik.com/docs/getting-started

---

This article provides an introduction to the initial setup of Kentik for your organization.

> **Note:** To manage settings on an ongoing basis, use the Settings link from the main nav menu (see **Main Settings**).

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(30).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A48Z&se=2026-05-12T09%3A31%3A48Z&sr=c&sp=r&sig=sDORjPW4KllQyIk0unOeH6bI9sorrCwkp6uySKw50J0%3D)

*The Welcome page is the hub for completing your initial Kentik setup tasks.*

## About Setup Tasks

Setup — a.k.a. "onboarding" — is the process of registering for the Kentik service so you can check out our flow telemetry (by sending flow data from at least one data source) or set up a test in Synthetics. While the setup UI is designed to efficiently get new customers (including trial customers) up and running, it's not designed for ongoing management of your organization's overall Kentik experience, which should be handled via the Settings section of the portal (see **Main Settings**).

#### Setup Tasks Workflow

Because trial signup is a required prerequisite to onboarding, if you don't already have an account then your setup process will begin on the main Kentik website (see **Starting Your Trial**) rather than on the Kentik portal. After your organization's Kentik account is activated your first login takes you to the setup workflow in the portal. After gathering information about the types of data — flow telemetry, cloud metrics, synthetic testing — for which you'll initially be using Kentik, the rest of the workflow consists of a number of "setup tasks" that are initiated from the Welcome to Network Observability page.

#### Setup Tasks on the Main Menu

You don't need to complete these tasks in one sitting; until the tasks are complete all Admin users in your organization will see a **Setup Tasks** section at the lower right of the main portal menu. The area will include a chart indicating the percent of tasks that have been completed, a list of incomplete tasks — each of which links to a setup page for that task — and a **See all** link that takes you to the Welcome page, where you can also see your remaining tasks. Once all the tasks are complete the menu will no longer show the **Setup Tasks** pane and all subsequent management of your portal settings will be handled via the Settings pages (click **Settings** in the list of featured modules at the upper left of the main menu).

## Starting Your Trial

Signup for a trial Kentik account is covered in the following topics.

### Sign Up for a Trial

Before using Kentik, you'll need to complete the trial signup process on the main Kentik website at **https://www.kentik.com**:

1. Click **Try Kentik** at the upper right.
2. Enter the following information:

   1. **Name**: Your first and last name.
   2. **Email**: The email address to use for your Kentik account.
   3. **Company**: The organization name for which you are opening an account.
   4. **Country**: Select your location from the dropdown.
   5. **Domain**: Select either the US or EU Kentik portal.
3. Click the **Privacy Policy** and **Terms of Use** links at the lower left to read about the terms of the trial.
4. Click **Get Started** at the bottom.
5. When the success notification appears, check your email for an activation link.

### Account Activation

Once you've signed up for a trial on the Kentik website, you'll be sent a confirmation email:

1. In the email, click the activation link.
2. Enter a password for your account on the **Account Activation** page (see **Account Activation UI**).
3. When the success notification appears, click the link to go the Kentik portal login page.
4. Enter the email address and password for your new account.
5. Upon successful log in, you'll be taken to the portal.

#### Account Activation UI

Set a password using the Account Activation page:

* **Password**: Enter the password here.

  > **Note:** Depending on your organization's **Password Expiration Policy**, passwords expire and must be reset every 90 days unless logging in via SSO or using any form of 2FA.
* **Password Strength**: A gauge showing the password strength and advice how to increase it. Strength is evaluated as one of the following levels:

  + **Too Short** (gray): Password is too short.
  + **Weak** (red): Password is not strong enough to use.
  + **Fair** (red): Password is better than weak, but still not strong enough to use.
  + **Good** (orange): Password could be stronger but Kentik will accept it.
  + **Strong** (green): Password is strong.

    > **Note:** If your password is not accepted, use the provided information to improve it.
* **Confirm Password**: Reenter your password here.
* **Show Password** (eye icon): Click to show/hide the entered password.
* **Set Password**: Click to complete password setup (disabled until the password meets strength standards).

## Initial Setup Login

If you're the first user from your organization to log into the portal after **Account Activation****,** you'll see this notification:

![Instructions to add a traffic source with a prominent 'Get Started' button.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/SO-get-started.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A48Z&se=2026-05-12T09%3A31%3A48Z&sr=c&sp=r&sig=sDORjPW4KllQyIk0unOeH6bI9sorrCwkp6uySKw50J0%3D)

To complete the setup process:

1. Click **Get started**.
2. Check all checkboxes that apply to the reasons your organization might like to use Kentik (these choices will influence the initial selection of visualizations on your **Observation Deck**):

   1. Improve User Experience
   2. Troubleshoot Network Problems
   3. Optimize Peering & Interconnection
   4. Manage Cloud Networking
   5. Stop DDoS Attacks
   6. Manage WAN, SD-WAN

   ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(32).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A48Z&se=2026-05-12T09%3A31%3A48Z&sr=c&sp=r&sig=sDORjPW4KllQyIk0unOeH6bI9sorrCwkp6uySKw50J0%3D)

   *Tell us your priorities for network observability with Kentik.*
3. Click the **Next** button at the bottom of the screen.
4. Check all checkboxes that apply to the data you want Kentik to collect:

   1. Public Cloud (check the boxes for all cloud providers that you use)
   2. Synthetic Performance Tests
   3. Networking Infrastructure

   ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(34).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A48Z&se=2026-05-12T09%3A31%3A48Z&sr=c&sp=r&sig=sDORjPW4KllQyIk0unOeH6bI9sorrCwkp6uySKw50J0%3D)

   *Choose the types of data you intend to use in Kentik on the "Which Data..." page.*
5. Click **Next** at the bottom.
6. The **Setup Tasks** list at left now shows a set of tasks based on your selections.

## Setup Tasks

The **Setup Tasks** list on the Welcome screen includes the following:

* **Start Performance Monitoring**: Use Kentik's Synthetics module (see **Synthetics**) for one of the following (see **Performance Monitoring Setup**):

  + Initiate continuous testing for a hostname, URL, or IP address (consumes test credits, 5M are free to each trial account).
  + See current performance for common SaaS services (e.g., Zoom, GSuite, Office365) on our **State of the Internet** (free).
* **Connect a router**: Start sending flow records to Kentik and performing analytics on actual telemetry data (see **Device Setup**).

  > **Note:** Connecting a router involves configuring it to send flow data (e.g., NetFlow, sFlow) to Kentik.
* **Define Network Boundaries**: Enter the CIDRs used to and from your network (public or private) and optionally provide your ASNs.

  + Enables Kentik to determine which traffic comes into (ingress) or goes out of (egress) your network (see **Define Network Boundaries**).
* **Monitor Cloud Resources**: This task’s name varies based on the cloud providers checked on the Which Data page.
* **Invite your team**: Click **Invite Your Team** to invite others in your organization to experience Kentik.

  + Kentik doesn't charge your organization per seat.
  + We encourage as many users as possible to use whatever aspects of Kentik they find helpful (see **Invite Users**).

## Performance Monitoring Setup

To complete the Start Performance Monitoring setup task by monitoring a hostname or IP:

1. On the Welcome page, click **Start Performance Monitoring**.
2. Enter a hostname or IP address to monitor.
3. Click **Start Monitoring**.
4. Choose the Kentik global agents to use to monitor the host/IP’s performance. As you include agents, pay attention to the **Expected monthly credit usage** indicator at the bottom.
5. Click **Start Monitoring** to exit and go to Kentik's **Synthetics Dashboard****.**

   1. The test is listed as "Pending" in the **Recently Added Tests** pane at the right of the page.
   2. Once the test is active, view results on the test’s details page in the **Test Control Center**.

   ![Monitoring setup for IP address with agent selection and credit usage details displayed.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/SO-start-monitoring.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A48Z&se=2026-05-12T09%3A31%3A48Z&sr=c&sp=r&sig=sDORjPW4KllQyIk0unOeH6bI9sorrCwkp6uySKw50J0%3D)    

## Define Network Boundaries

The Define Network Boundaries setup task is an implementation of **Network Classification**, which uses source and destination information related to IPs and ASes to determine the direction of network traffic with respect to your network. To complete the task:

1. On the Welcome page, click **Define Network Boundaries** in the **Setup Tasks** list.
2. In the IPs field, enter the CIDRs used to and from your network (public or private). Optionally, in the ASNs field, provide your internal ASNs.
3. Click **Save** to return to the Welcome page and see a notification that the network classification was successful.

     ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(36).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A48Z&se=2026-05-12T09%3A31%3A48Z&sr=c&sp=r&sig=sDORjPW4KllQyIk0unOeH6bI9sorrCwkp6uySKw50J0%3D)

## Invite Users

The invite users task involves giving colleagues the opportunity to experience Kentik so everyone can work from a common understanding of the network. Kentik does not charge by the seat.

To invite others to try Kentik:

1. ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(37).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A48Z&se=2026-05-12T09%3A31%3A48Z&sr=c&sp=r&sig=sDORjPW4KllQyIk0unOeH6bI9sorrCwkp6uySKw50J0%3D)On the Welcome page, click **Invite Your Team** in the **Setup Tasks** list.
2. Enter the invitee’s email in the **Email** field.
3. Enter the invitee’s name in the **Name** field.
4. Click **User Level** to set a level for the user. Admin users can both see data and make settings, while Member users are generally not able to make settings.
5. To add additional users, click the **Add User** button and enter the required information. Click the trash icon to remove a user.
6. Click **Send Invites** to finish and return to the Welcome page and see a notification that the invitation emails were sent.
