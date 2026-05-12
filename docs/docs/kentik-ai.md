# Kentik AI

Source: https://kb.kentik.com/docs/kentik-ai

---

This article discusses the settings page for **Kentik AI** in the Kentik portal.

![The Kentik AI settings page showing the Enable Kentik AI switch disabled by default.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/KAS-main-disabled.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A15Z&se=2026-05-12T09%3A31%3A15Z&sr=c&sp=r&sig=pB5Q4MlWmOUYUmvM7XSuLEyLC3JSy9rUCilbxwUJmdU%3D)

*The Kentik AI Settings page showing Kentik AI disabled by default.*

> **Note:** Before enabling Kentik AI, see **Kentik AI Overview** for a full understanding of what it is, how it works, and to learn about the precautions we take to ensure our customers' data privacy and security.

## About Kentik AI Settings

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(105).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A15Z&se=2026-05-12T09%3A31%3A15Z&sr=c&sp=r&sig=pB5Q4MlWmOUYUmvM7XSuLEyLC3JSy9rUCilbxwUJmdU%3D)

The Kentik AI settings page allows you to enable and configure Kentik AI features like **AI Advisor** and others. Key points include:

* Accessible only to **Super Administrator** users (see **About Users**) of the Kentik portal.
* Navigate to it by clicking the navbar's **Organization Settings** drop-down and selecting **Kentik AI**.
* This page allows you to turn On or Off aspects of Kentik AI for the entire organization (see **Customer Control of AI**).

## Manage Kentik AI

Follow these steps to manage Kentik AI settings in the Kentik portal.

### Overview Pane

The Overview pane is the default pane of the Kentik AI settings page, and contains the switch to set Kentik AI to either disabled (default) or enabled for your entire organization:

1. Choose **Kentik AI** from the Organization Settings menu to open the Kentik AI settings page.

   > **Important**: This menu option is only present for users with Super Administrator privileges.
2. Set the **Enable Kentik AI** switch On to enable Kentik AI. Set the switch to Off to disable Kentik AI.
3. When Kentik AI is enabled, the following additional settings panes display in the left sidebar. Follow the links for details:

   1. **Models & Credentials**(not enabled for most customers)
   2. **Custom Network Content**
   3. **Runbooks**

![Kentik AI settings page showing the Enable Kentik AI switch enabled and features like Query Assistant and AI Advisor highlighted.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/KAS-main-enabled.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A15Z&se=2026-05-12T09%3A31%3A15Z&sr=c&sp=r&sig=pB5Q4MlWmOUYUmvM7XSuLEyLC3JSy9rUCilbxwUJmdU%3D)

*Kentik AI is enabled*

### Models & Credentials Pane

> **Important**: The Models & Credentials pane is only visible/enabled by request, typically by customers that require a specific provider/region/model to be used by Kentik AI. Contact Kentik support for assistance with this feature.

The Models & Credentials pane of the Kentik AI settings page contains the settings for the Advanced and Basic models used by Kentik AI.

* ![Kentik AI settings page showing model and provider selection options for AI tasks.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/KAS-models-and-credentials-pane.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A15Z&se=2026-05-12T09%3A31%3A15Z&sr=c&sp=r&sig=pB5Q4MlWmOUYUmvM7XSuLEyLC3JSy9rUCilbxwUJmdU%3D)**Advanced Model**: Used for interactive tasks such as **AI Advisor** and other user-driven interactions.
* **Basic Model**: Optimized for lightweight tasks like summarizing data, detecting change points, and analyzing probably causes.

For each model, you can customize the provider and corresponding model, region, credentials, and API token as applicable. Click the **Save** button to save your changes.

* **Provider**: Choose a Large Language Model (LLM) provider and model for Kentik AI.

  + Kentik default
  + Open AI
  + Google Vertex AI
  + Google AI Studio (Google Gemini models)

> Notes:
>
> * If Kentik default is selected, the other settings will not display.
> * If additional support is required (Azure, AWS, Anthropic), open a feature request with Kentik.
> * To opt for “Bring-Your-Own-Key” mode, select **Use my own token** and follow the steps in **Create AI API Token**.

* **Model**: The \* symbol indicates this is a required field. Depending on the provider you choose, select one of the models in the drop-down selection.
* **Region**: The \* symbol indicates this is a required field. Depending on the model you choose, select on of the regions in the drop-down selection.
* **Credentials**: Select Use Kentik Access or Use my own token. If you select Use my own token, then the API Token section will display.
* **API Token**: When Use my own token credentials is selected, choose an API Token from the drop-down selection. You also have the option to Create an AI API Token (see **Create AI API Token**).

#### Create AI API Token

When Use my own token credentials is selected, you can either select an API Token or Create your own AI API Token. Follow the steps below to Create your own AI API Token.

1. Ensure **Use my own token** is selected in the Credentials section.
2. Select **Create AI API Token**.
3. Complete the required fields in the **Add Credential** popup window.

   1. To edit existing credentials, click the **Credentials Vault** link. This will navigate you to the “Settings » Credentials Vault” page in the portal.
4. Click **Add Credential** to save your changes for the API Token.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(754).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A15Z&se=2026-05-12T09%3A31%3A15Z&sr=c&sp=r&sig=pB5Q4MlWmOUYUmvM7XSuLEyLC3JSy9rUCilbxwUJmdU%3D)

*The Credentials and API Token settings for Kentik AI*

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(755).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A15Z&se=2026-05-12T09%3A31%3A15Z&sr=c&sp=r&sig=pB5Q4MlWmOUYUmvM7XSuLEyLC3JSy9rUCilbxwUJmdU%3D)

*The Add Credential popup window for creating an AI API token*

### Custom Network Context Pane

![The Kentik AI setting page is shown with Custom Network Context pane selected.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/KAS-custom-network-context-pane.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A15Z&se=2026-05-12T09%3A31%3A15Z&sr=c&sp=r&sig=pB5Q4MlWmOUYUmvM7XSuLEyLC3JSy9rUCilbxwUJmdU%3D)The Custom Network Context pane allows you to add additional text-based information about your network environment. This helps Kentik AI deliver more accurate and relevant insights as your organization’s needs evolve. Here’s how it works:

* Add up to 100K characters of Markdown-formatted text into the **Custom Network Context** input field.
* Click the eye icon to preview the rendered Markdown.
* The context you provide can be as specific (e.g., IP addressing schemes, network design info) or general (strategies for troubleshooting, personas to take on) as necessary to accomplish your organization’s goals.
* Kentik AI uses this additional context in all conversations in your organization.

### Runbooks Pane

The Runbooks pane of the Kentik AI settings page contains settings for creating Runbooks, or predefined, plain text or Markdown recipes to guide AI Advisor through the diagnostic steps for specific alerts. Runbooks can:

* Be assigned to one or more alert policies
* Start the investigation from any alert with all the relevant information already in place
* Help to minimize human and/or AI errors by ensuring a pre-defined, systematic troubleshooting approach for specific situations

![Kentik AI settings interface displaying several Runbooks listed.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/KAS-runbooks-pane(1).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A15Z&se=2026-05-12T09%3A31%3A15Z&sr=c&sp=r&sig=pB5Q4MlWmOUYUmvM7XSuLEyLC3JSy9rUCilbxwUJmdU%3D)

*The Kentik AI settings page showing a list of Runbooks.*

#### Add a Runbook

To add a new Runbook and assign it to an existing Kentik alert policy, follow these steps:

1. Click the Runbooks pane on the Kentik AI settings page.
2. Click **+ Add Runbook** to open the dialog.

   1. **Name** (required)**:** Enter a unique name for the Runbook.
   2. **Description**:Optionally, add a Runbook description.
   3. **Instructions**: Add your Markdown-formatted text with detailed instructions and context for Kentik AI to use in a particular situation.
3. Navigate to the **Alert Policy Assignments** tab to see the list of alert policies in your organization.
4. Optionally, check the box for each alert policy to which you want to assign this Runbook.
5. Click **Add Runbook** to create the new Runbook and return to Kentik AI settings page.

#### Revert a Runbook to a Prior Version

To view the list of saved versions of a Runbook and revert to a previous version, follow these steps:

1. Click the Runbooks pane on the Kentik AI settings page.
2. Click the Runbook Revision History button (the ‘rewind the clock’ icon) next to the Runbook in the list.
3. The current revision’s contents are shown in the left pane while the list of revisions is shown in the right pane. Click a revision in the right pane to view it.
4. When viewing a past revision, changes to the Runbook are marked in green (addition) and red (removal).
5. Click **Revert** to adopt the previous Runbook version and return to the Kentik AI settings page, or click the **X** to exit without saving.

#### Edit/Delete a Runbook

To edit or delete a Runbook, follow these steps.

1. Click the Runbooks pane on the Kentik AI settings page.
2. To edit a Runbook:

   1. Click the Edit button (pencil icon) next to the Runbook in the list.
   2. Make the changes (see **Add a Runbook**), then click **Save** to save changes and exit.
3. To delete a Runbook:

   1. Click the Delete button (trash can icon) next to the Runbook in the list.
   2. Click **Remove** to confirm deletion or **Cancel** to exit without saving.

### Privacy and Security

The Privacy & Security pane of the Kentik AI Settings page is informational only, and links to the **Kentik AI Overview** article in the Kentik Knowledge Base. Please read this article before enabling Kentik AI.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/Screenshot 2025-08-26 at 13.28.44.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A15Z&se=2026-05-12T09%3A31%3A15Z&sr=c&sp=r&sig=pB5Q4MlWmOUYUmvM7XSuLEyLC3JSy9rUCilbxwUJmdU%3D)
