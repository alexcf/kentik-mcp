# Metadata Configuration

Source: https://kb.kentik.com/docs/metadata-configuration-aws

---

Metadata collection allows Kentik to **discover your AWS infrastructure** (VPCs, Subnets, Gateways) and visualize your network topology on the **Kentik Map**. This guide covers how to select and configure an IAM architecture that enables Kentik to access your AWS metadata.

> **When Should I Use Metadata Collection Only?**
>
> * **Topology Only**: You want to see your network inventory and hierarchy without **ingesting traffic logs**.
> * **Split Architecture**: Your flow logs are centralized in a "Log Archive" account, but you need to gather metadata from the various "Member" accounts where the resources actually live.

## IAM Architecture: Standard vs. Nested

To authorize Kentik to fetch your AWS metadata, you must configure IAM roles. Choose the strategy that matches your AWS environment.

| Feature | **Standard Strategy** | **Nested Strategy** |
| --- | --- | --- |
| **Connection** | Kentik connects to every account directly. | Kentik connects to one "Primary" account. |
| **Trust Mechanism** | Individual Trust Policy per account. | Primary Role "assumes" Secondary Roles. |
| **Best For** | Small environments or single accounts. | Large Orgs, Control Tower, or MSPs. |

![Diagram illustrating AWS account roles and API interactions for Kentik integration.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/AWS-nested-structure.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A47Z&se=2026-05-12T09%3A32%3A47Z&sr=c&sp=r&sig=lGY0kjmEapdI6qC83fQ4uTQOyrt8FFBsjbhyzg1odyk%3D)

*A nested structure in which account A is primary and accounts B and C are secondary.*

## Standard Configuration

For a **standard** AWS metadata collection setup, repeat these steps for every account you want Kentik to monitor.

### **Create Policy**

Create a JSON policy in the AWS account using the permissions required for metadata (EC2 describe, CloudWatch, etc.).

To set up an AWS account policy for metadata export to Kentik:

1. Log into the AWS account you want to monitor.
2. Navigate to **IAM »** **Policies** and click **Create policy**.
3. Select the **JSON** tab.
4. Replace the editor’s content with the required metadata permissions.

   > **Note**: You can use the exact same JSON block as the Secondary Policy JSON here, as it contains the necessary `ec2:Describe*`, `cloudwatch:ListMetrics`, and `s3:Get*` actions.
5. Click **Next**.
6. Provide a name and description for the new policy under *Policy details*.
7. Click **Create policy** to save and exit.

### **Create IAM Role**

To create the role that Kentik will assume:

1. Navigate to **IAM »** **Roles** and click **Create role**.
2. Select **Custom trust policy**.
3. Replace the editor content with a trust policy that sets the Principal to Kentik’s AWS Account: `arn:aws:iam::834693425129:root` (see JSON directly below).

   > **CRITICAL**: In the `Condition` block for `sts:ExternalId`, replace the placeholder with your Kentik Company ID (found in the portal under **Settings »** **Licenses**).
4. Click **Next**.

#### Standard Role Trust Policy JSON

The following JSON assigns a trust policy to the role, enabling access by Kentik's AWS account.

> **IMPORTANT**: In the `sts:ExternalId` field, you must replace `<your_Company_ID>` with your specific Kentik company ID, which is the "Account #" on the portal's **Licenses** page (**Settings »** **Licenses**).

```
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "KentikTrust",
      "Effect": "Allow",
      "Principal": {
        "AWS": "arn:aws:iam::834693425129:root"
      },
      "Action": "sts:AssumeRole",
      "Condition": {
        "StringEquals": {
          "sts:ExternalId": "<your_Company_ID>"
        },
        "ArnEquals": {
          "aws:PrincipalArn": "arn:aws:iam::834693425129:role/eks-ingest-node"
        }
      }
    }
  ]
}
```

### **Attach Policy**

To link your permissions to the new role:

1. On the **Add permissions** screen, search for and select the policy you created in the first section.
2. Click **Next**.
3. Enter a role name and description.
4. Click **Create role** to save and return to the Roles page.

## Nested Configuration

For a **nested** AWS metadata collection setup, you configure one “**hub**” account and multiple “**spoke**” accounts.

### Provision the Hub Account

#### Create a Primary Policy

To set up a primary AWS account policy for metadata export to Kentik:

1. Log into the AWS account designated as the primary for metadata export.
2. Navigate to **IAM »** **Policies** and click **Create policy**.
3. Select the **JSON** tab.
4. Replace the editor’s content with the JSON specified in **Primary Policy JSON** below.
5. Click **Next** and ensure the policy includes STS as an action.
6. Provide a name and description for the new policy under **Policy details**.
7. Click **Create policy** to save and exit.

#### Primary Policy JSON

The following JSON defines a policy to enable access to the primary AWS account:

```
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "AssumeSecondaryRoles",
      "Effect": "Allow",
      "Action": "sts:AssumeRole",
      "Resource": "*"
    },
    {
      "Sid": "OrgListing",
      "Effect": "Allow",
      "Action": [
        "organizations:ListAccounts"
      ],
      "Resource": "*"
    }
  ]
}
```

#### Create a Primary Role

To assign the created policy to a role in the primary AWS account:

1. Navigate to **IAM »** **Roles** and click **Create role**.
2. Select **Custom trust policy**.
3. Replace editor content with the **Primary Role JSON** below.

   > **CRITICAL**: You must replace `<your_Company_ID>` with your specific Kentik company ID (found in the portal under **Settings »** **Licenses**).
4. Click **Next**. Find and select your policy and click **Next**.
5. Enter a role name and description.
6. Click **Create role** to save and return to the **Roles** page.

#### Primary Role JSON

The following JSON  assigns a trust policy to a role in the primary account, enabling access by Kentik's AWS account. In the `sts:ExternalId` field, use your Kentik company ID which is the "Account #" on the portal's **Licenses** page (**Settings »** **Licenses**).

```
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "KentikTrust",
      "Effect": "Allow",
      "Principal": {
        "AWS": "arn:aws:iam::834693425129:root"
      },
      "Action": "sts:AssumeRole",
      "Condition": {
        "StringEquals": {
          "sts:ExternalId": "<your Company or External ID here>"
        },
        "ArnEquals": {
          "aws:PrincipalArn": "arn:aws:iam::834693425129:role/eks-ingest-node"
        }
      }
    }
  ]
}
```

> **IMPORTANT**:
>
> * If you prefer to use a 16-digit randomized string as your `ExternalId` instead of your Kentik Company ID, email Kentik support at **support@kentik.com**.
> * For more on `ExternalId`, see **Automated Configuration Options**.

### Provision the Spoke Accounts

#### Create Secondary Policies

To provision secondary AWS accounts to enable access to their metadata by the primary account:

1. Log into the AWS console for the secondary account.
2. Go to the **Policy editor** page.
3. Replace editor content with the **Secondary Policy JSON** below.
4. Click **Next**, then enter a name and description for the new policy.
5. Click **Create policy** to save and return to the **Policies** page.
6. Repeat for each secondary account.

#### Secondary Policy JSON

```
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "cloudwatch:ListMetrics",
        "cloudwatch:GetMetricStatistics",
        "cloudwatch:GetMetricData",
        "organizations:ListAccounts",
        "cloudwatch:Describe*",
        "directconnect:List*",
        "directconnect:Describe*",
        "ec2:Describe*",
        "ec2:Search*",
        "ec2:GetManagedPrefixListEntries",
        "elasticloadbalancing:Describe*",
        "iam:ListAccountAliases",
        "network-firewall:Describe*",
        "network-firewall:List*",
        "networkmanager:ListCoreNetworks",
        "networkmanager:GetCoreNetwork",
        "networkmanager:GetCoreNetworkPolicy",
        "networkmanager:ListAttachments",
        "networkmanager:GetNetworkRoutes"
      ],
      "Resource": "*"
    },
    {
      "Effect": "Allow",
      "Action": [
        "s3:Get*",
        "s3:List*"
      ],
      "Resource": [
        "arn:aws:s3:::YOUR_S3_BUCKET_NAME",
        "arn:aws:s3:::YOUR_S3_BUCKET_NAME/*"
      ]
    }
  ]
}
```

#### Create Secondary Roles

To create a role in each secondary AWS account and assign the policy to the role:

1. Log into the AWS console for the secondary account.
2. Go to the **Select trusted entity page**.
3. Choose **Custom trust policy**.
4. Replace editor content with the **Secondary Role JSON** below.

   > **CRITICAL**: Replace `primary_account_id` with the account ID of the **primary** account.
5. Assign the created policy to the new role, enter a name and description, and click **Create role**.
6. Repeat for each secondary account.

#### Secondary Role JSON

The following JSON assigns a policy to a role in the secondary AWS account, enabling access by the primary account:

```
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "AWS": "arn:aws:iam::primary_account_id:root"
      },
      "Action": "sts:AssumeRole",
      "Condition": {}
    }
  ]
}
```
