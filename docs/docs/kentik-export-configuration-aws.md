# Kentik Export Configuration

Source: https://kb.kentik.com/docs/kentik-export-configuration-aws

---

Now that you have successfully prepared your AWS environment, by **configuring your IAM roles** (Standard or Nested) and **granting the necessary S3 permissions for flow logs**, the final step is to register that setup in Kentik.

You do this by creating a **Cloud Export** in the Kentik portal. A cloud export represents the active data pipeline between your AWS architecture and Kentik, and it can be configured to collect various types of telemetry:

* **Metadata Only**: Collects only AWS metadata. Use this when your flow logs are already being ingested by a different cloud export.
* **Flow/Firewall Logs and Metadata**: Collects telemetry logs and metadata for all entities (VPCs, subnets, interfaces, firewalls) publishing to a **specific S3 bucket**. Kentik automatically creates a **cloud device** for each entity.
* **Metrics**: Cloud metrics history for historical telemetry analysis, trending, and alerting.

![Cloud export configuration settings for AWS, GCP, Azure, and OCI with observability features.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/KFA-create-cloud-export-provider-and-features(1).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A07Z&se=2026-05-12T09%3A32%3A07Z&sr=c&sp=r&sig=sPwnM0djpNSYymekZX45CRHeUIFycfM4DOQ9lW7KQ7U%3D)

*Configuration settings for AWS Cloud Export with various observability features listed.*

## Initial Cloud Export Steps

To create a new AWS cloud export:

1. Navigate to **Settings »** **Public Clouds** and click **Create Cloud Export**.
2. Click **AWS Cloud**.
3. Under **Observability Features**, select the data types to collect:

   1. **Metadata collection (Required): Automatically selected.**
   2. **Flow log collection**: Select to collect AWS flow logs.

      1. **Help me configure my provider via Terraform**: Choose to automatically configure the cloud export using Terraform in the next step of the wizard (see **Automated Setup**).
   3. **Firewall log collection**: Select to collect AWS firewall logs.
   4. **Cloud metrics history**: Select to collect AWS CloudWatch metrics.
4. Click the green arrow to proceed.

The next steps depend on the export type:

* **Metadata-Only Export**
* **Flow/Firewall Logs and Metadata Export**

## Metadata-Only Export

To set up a new AWS metadata-only cloud export, follow these steps:

1. Complete the **Initial Cloud Export Steps** while leaving **Flow log collection** unselected. (Selecting **Cloud metrics history** is optional).
2. In the **AWS Role** field (Required):

   1. **Standard Configuration**: Enter the ARN of the **IAM role** you created in your single AWS account. Leave “This is an AWS organization role” unselected.
   2. **Nested Configuration**: Enter the ARN of the **Primary (Hub) Role**. If using AWS Organizations to automatically discover child accounts, select the “**This is an AWS organization role**” checkbox.
   3. Click **Verify Role**.
3. Select an **AWS Region** from the dropdown (Required):

   1. Choose the region of the primary account.
   2. Click **Verify Region** (fails if the **AWS Role** is blank or invalid).
4. Specify Additional Roles (Nested Configurations Only): Expand the **Optional: Additional Metadata Roles** pane to access these options:

   1. **Secondary AWS Accounts**: Enter a comma-separated list of the **secondary (spoke) account IDs** you configured earlier.
   2. **Regions Used Across Secondary Accounts**: Select from the filterable dropdown all regions where these secondary accounts exist.
   3. **Role suffix**: Enter the exact name of the Secondary Role you created in those accounts.
   4. Click **Verify** to validate all entered values.
5. Click the green arrow to proceed to the final step.
6. Enter the cloud export name/description:

   1. **Name** (Required): Specify or accept the default name for the cloud export.
   2. **Description**: Provide a description or accept the default.
7. Select the appropriate Kentik billing plan for the cloud export from the **Billing Plan** dropdown.
8. Click **Save** to finalize the cloud export and return to the **Public Clouds** page, where the new export will be listed.

## Flow/Firewall Logs and Metadata Export

To set up a new AWS flow logs and metadata export, follow these steps:

1. Complete the **Initial Cloud Export Steps** while selecting **Flow log collection** and/or **Firewall log collection**.
2. Complete the first three steps of **Metadata-Only Export**.
3. **S3 Bucket Name**: Provide the exact name of the S3 bucket where your flow logs are stored.

   > **Note**: This must be the same bucket name you added to your IAM policy’s S3 permissions in the previous step.
4. **S3 Bucket Prefix:** Specify a prefix for Kentik to add to the S3 bucket name when creating the cloud export.
5. Click **Verify S3 Bucket** to ensure the bucket is accessible and correctly configured.
6. Complete the “Optional: Additional Metadata Roles” section (described in Step 4 of **Metadata-Only Export**)
7. **Sampling**: Configure your desired flow sampling settings (see **Cloud Export Sampling**).
8. Click the green arrow to proceed.
9. Specify or accept the default name for the cloud export.
10. Optionally provide a description for the cloud export or accept the default.
11. Choose the appropriate Kentik billing plan for the cloud export from the dropdown.
12. Click **Save** to finalize the cloud export and return to the Public Clouds page, where the new export will be listed.

### Automated Setup

To automatically configure your AWS setup using Terraform, follow these steps.

![Options for flow log collection and AWS log management settings are displayed.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/KFA-create-cloud-export-auto-1.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A07Z&se=2026-05-12T09%3A32%3A07Z&sr=c&sp=r&sig=sPwnM0djpNSYymekZX45CRHeUIFycfM4DOQ9lW7KQ7U%3D)

1. Follow the **Initial Cloud Export Steps** and select the **Help me configure my provider via Terraform** box.
2. For AWS Provider Profile Name, the default is “default”. Enter a different name if needed.
3. Select the AWS region from the dropdown, which populates the `region` field in the generated configuration.
4. Configure settings in the **Select options** section (see **Automated Configuration Options**).
5. Copy the generated configuration and save it as `main.tf` in an empty directory where Terraform will be run.
6. Execute the commands provided in the wizard to apply the configuration.
7. Click **Finish** to return to the **Public Clouds** page, where the new cloud export will be listed under **Cloud Exports**.

![Configuration settings for AWS provider profile, including region and logging options.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/KFA-create-cloud-export-auto-2.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A07Z&se=2026-05-12T09%3A32%3A07Z&sr=c&sp=r&sig=sPwnM0djpNSYymekZX45CRHeUIFycfM4DOQ9lW7KQ7U%3D)

### Automated Configuration Options

When configuring AWS setup automatically via Terraform in Kentik, you can customize the following options:

* ![Automated configuration options for an AWS cloud export.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/KFA-create-cloud-export-automated.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A07Z&se=2026-05-12T09%3A32%3A07Z&sr=c&sp=r&sig=sPwnM0djpNSYymekZX45CRHeUIFycfM4DOQ9lW7KQ7U%3D)**Enable flow logs**:

  + **For all VPCs in the selected region(s)**: Automatically configures flow logs for all VPCs in the selected region.
  + **For selected VPCs in the selected region(s)**: Enter VPC IDs in the `vpc_id_list` parameter to configure only those VPCs.
* **Write logs to bucket**:

  + **Every minute (recommended)**: Provides a higher volume of logs at a consistent rate, ideal for traffic engineering, security, and real-time monitoring.
  + **Every 10 minutes** **(AWS default)**: Reduces log volume and AWS charges.
* **Automatically create necessary role in AWS account**: Decide whether to automatically create the AWS role or manage it manually according to your security protocols.
* **Use External ID**: Optionally, include an AWS external ID for Kentik to use to access your S3 bucket (see **AWS doc on External ID**):

  + This ID is known only to you and Kentik. Per AWS, its primary purpose is to avoid the **confused deputy problem**.
  + By default, your Kentik company ID is used.
  + This ID should also be used when creating the AWS role (see **Primary Role JSON**).

    > **Note**: If you prefer to use a 16-digit randomized string as your External ID, contact Kentik support at **support@kentik.com**.
* **Cloud Export Name Prefix**: Specify a prefix to add to the Kentik cloud export name for easy identification.
* **S3 Bucket Prefix**: Specify a prefix to add to Kentik-created S3 bucket name.
* **IAM Role Prefix**: Specify a prefix to add to the Kentik-created IAM role.
* **Billing Plan**: Select the appropriate Kentik billing plan for the cloud export.

> **Notes**:
>
> * Prefix fields help in identifying and managing your cloud exports more effectively.
> * Different values can be used for each prefix field to suit your organizational needs.

## Using Your Cloud Export

Once the setup process is complete, you can view and utilize your cloud export in Kentik:

* **Cloud Exports List**

  + Go to **Settings »** **Public Clouds** to see the updated list of cloud exports.
  + A new cloud export will be listed, representing the VPCs, transit gateways, subnets, or interfaces whose logs are pulled from the specified bucket.
* **Devices Column**

  + Each VPC, transit gateway, subnet, or interface sending flow logs is listed as a cloud device.
  + Devices are named after their respective VPC, transit gateway, subnet, or interface.
  + These names can be used as group-by and filter values in Kentik queries using the Device Name dimension.
* **Metadata and Mapping**

  + The collected metadata, such as routing tables, security groups, and ACLs, enables Kentik to automatically map and visualize the topology of your AWS resources in the **Kentik Map**.

![AWS service status overview with highlighted issues and device group details.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(45).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A07Z&se=2026-05-12T09%3A32%3A07Z&sr=c&sp=r&sig=sPwnM0djpNSYymekZX45CRHeUIFycfM4DOQ9lW7KQ7U%3D)

*The Public Clouds page lists your AWS resources as “cloud exports”, each with a service status overview, highlighted issues, and device group details.*
