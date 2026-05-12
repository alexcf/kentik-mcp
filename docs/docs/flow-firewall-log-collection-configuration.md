# Flow/Firewall Log Collection Configuration

Source: https://kb.kentik.com/docs/flow-firewall-log-collection-configuration

---

After **configuring IAM permissions for metadata collection**, the next step is to set up flow and/or firewall log publishing from your VPCs to an S3 bucket.

> **Note:** The setup can be applied to VPCs, subnets, or interfaces. Refer to Amazon’s **Creating a Flow Log that Publishes to Amazon S3** for specific details on creating flow logs for different resources.

## S3 Bucket Strategy

Before creating a bucket, decide on your storage strategy. This choice impacts both management overhead and AWS data transfer costs.

| Strategy | Description | Best For |
| --- | --- | --- |
| **Local Buckets** | Create a separate S3 bucket in each region where you have resources (e.g., `logs-us-east-1`). | **Cost & Latency**: Eliminates cross-region data transfer fees.  **Compliance**: Keeps data resident in specific geographic regions (e.g., GDPR). |
| **Centralized Bucket** | Send logs from all regions to a single, central S3 bucket. | **Simplicity**: Easier to manage if you have low traffic volume or need a central archive.  **Note**: You will incur AWS cross-region data transfer costs. |

## S3 Bucket Management

When configuring the S3 bucket that will hold your flow and firewall logs, Kentik recommends the following best practices to ensure optimal visibility and cost efficiency.

### Log Record Formats

While AWS supports both default and custom flow log formats, Kentik strongly recommends using a **Custom Format** when setting up your log exports.

* **Recommended Configuration:** Select Custom format and enable all available fields (v2, v3, v4, and v5).
* **Why?** This ensures Kentik receives not just basic 5-tuple traffic data, but also critical context like TCP flags, packet drop reasons, and region/zone information.
* **References:** **VPC Flow Log Records** | **Transit Gateway Flow Log Records**| **Network Firewall Log Records**

## Create an S3 Bucket

Begin by creating an S3 bucket to store flow logs, which Kentik will access later to collect metadata and logs.

1. Navigate to the **Amazon S3 console**and click **Create Bucket**.
2. Enter a descriptive name (e.g., `kentik-flow-logs-prod`) and select your target region (see **S3 Bucket Strategy**).
3. Scroll to "Default encryption".

   * **Recommended:** Select **SSE-S3** (Server-side encryption with Amazon S3 managed keys). This requires no extra configuration.
   * **Advanced:** If you must use **SSE-KMS**, you will need to add specific `kms:Decrypt` and `kms:GenerateDataKey` permissions to your IAM policy (see **Advanced: S3 Bucket Encryption (KMS)**).
4. Leave all other settings as default and click **Create**. The new bucket will appear in the S3 console bucket list.

   > **Note:** Refer to **Bucket Naming Rules** for conventions.

### Advanced: S3 Bucket Encryption (KMS)

If your S3 bucket uses **SSE-KMS** (AWS Key Management Service) for encryption, you must grant the Kentik IAM role permission to use your specific key. Without this, Kentik will be able to "see" the files but will be "Access Denied" when trying to read/decrypt them.

Add the following block to the `Statement` array of the IAM policy attached to the role Kentik assumes:

```
{
  "Sid": "KMSAccess",
  "Effect": "Allow",
  "Action": [
    "kms:Decrypt",
    "kms:DescribeKey"
  ],
  "Resource": "arn:aws:kms:region:account_id:key/key_id"
}
```

> **Notes:**
>
> * Replace the `Resource` ARN with the actual ARN of your KMS key. You can find this in the AWS Console under **KMS > Customer managed keys**.
> * Kentik does not support the SSE-C encryption option for S3 buckets.

## Enable VPC Flow Logs

To start sending traffic telemetry to Kentik, you must enable flow logs on your target resources:

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/AWS-Your_vpc_page-391h674w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A43Z&se=2026-05-12T09%3A31%3A43Z&sr=c&sp=r&sig=2w67Fb8sDzWzOsnRlufp9IwUG7UYHIV2%2BihTG2PSvrE%3D)

1. Go to the **AWS VPC Dashboard**.
2. Click **Your VPCs**, select the desired VPC(s), and choose **Actions >** **Create flow log**.
3. Configure Settings:

   1. **Filter**: Select **All** (recommended for best visibility)
   2. **Destination**: Select **Send to an S3 bucket**
   3. **S3 bucket ARN**: Paste the ARN of the bucket you created in the previous steps (e.g., `arn:aws:s3:::kentik-flow-logs`).
4. **Log Record Format (Critical):**

   1. Select **Custom format**
   2. Check the box for “Select all” (or manually select all AWS v2, v3, v4, and v5 fields).

      > **Why?** Default logs miss critical context like TCP flags and packet size, which limits Kentik's ability to visualize traffic types.
5. Click **Create flow log****.** The resulting page confirms log creation of the AWS-assigned log ID. Click **Close** to return to **Your VPCs**.

> **IMPORTANT: Rate Limits**
>
> Kentik treats each S3 Bucket as a single "Device." If you have hundreds of VPCs, we recommend consolidating them into regional buckets (e.g., `logs-us-east-1`, `logs-eu-west-1`) rather than one global bucket. This prevents AWS S3 throttling and ensures smooth ingestion.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/AWS-Create_flow_log-499h674w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A43Z&se=2026-05-12T09%3A31%3A43Z&sr=c&sp=r&sig=2w67Fb8sDzWzOsnRlufp9IwUG7UYHIV2%2BihTG2PSvrE%3D)

## Enable Network Firewall Logs

In addition to standard VPC flow logs, Kentik can ingest traffic telemetry directly from AWS Network Firewall. You can send these logs to the exact same S3 bucket you configured above.

To enable logging for your AWS Network Firewall:

1. Go to the **AWS VPC Dashboard**.
2. In the left navigation pane, under **AWS Network Firewall**, select **Firewalls**.
3. Click the name of the firewall you want to monitor.
4. Scroll down to the **Logging details** section and click **Edit**.
5. Under **Log type**, select the types of logs you want to send to Kentik:

   * **Flow logs:** Standard network traffic flow details.
   * **Alert logs:** Logs for traffic that matches your stateful rules.
6. Under **Destination type**, select **S3 bucket**.
7. In the **S3 bucket** field, select or paste the name of the S3 bucket you created earlier (e.g., `kentik-flow-logs`).
8. Optionally, provide a prefix to separate these logs from your standard VPC flow logs (e.g., `firewall-logs/`).
9. Click **Save**.

> **Note:** Just like VPC flow logs, it may take 5–10 minutes for the first Network Firewall logs to appear in your S3 bucket. You can verify them by checking the bucket for a newly created `AWSLogs/` folder path containing your firewall data.

## Verification: Check Log Status

To confirm the AWS flow/firewall log setup:

1. **Wait:** Allow 10 minutes for the first flow logs to generate. Flow logs are generally published every 5–10 minutes, only if there’s active VPC traffic.
2. **Inspect S3:** Go to your S3 bucket and look for the folder path: `AWSLogs / <Account_ID> / vpcflowlogs / ...` .
3. **Check Files:** If you see `.log.gz` files appearing, data is successfully flowing to the bucket.

> **Note**: Firewall logs appear in a slightly different subfolder structure but the verification concept is identical.

## Update IAM Permissions for S3

Now that you’ve created an IAM role and policy during the **Metadata Configuration** phase, you need to grant that existing role permission to read your new S3 flow/firewall log bucket.

### Locate Your Existing Policy

Log into your AWS account and navigate to **IAM »** **Policies**. Find the policy you created during the metadata setup:

* **Standard Configuration:** Locate your standard metadata policy.
* **Nested Configuration:** Locate your **Secondary (Spoke)** policy.

### Add S3 Permissions

Edit the policy to include the following statement in the `Statement` array, replacing `<your-logs-bucket>` with the exact name of the S3 bucket you created earlier:

```
{
  "Effect": "Allow",
  "Action": [
    "s3:Get*",
    "s3:List*"
  ],
  "Resource": [
    "arn:aws:s3:::<your-logs-bucket>",
    "arn:aws:s3:::<your-logs-bucket>/*"
  ]
}
```

**Important Notes:**

* **Flow Log Deletion:** If you want Kentik to automatically manage and delete old flow logs to save on AWS storage costs, change the `Action` array to `"Action": "s3:*"`.
* **KMS Encryption:** If your bucket uses SSE-KMS encryption, remember that you also need to add the `kms:Decrypt` and `kms:DescribeKey` permissions as outlined in the **Advanced: S3 Bucket Encryption** section above.

### Save Changes

Save the updated policy. Because this policy is already attached to the IAM role Kentik uses, the new S3 permissions will take effect automatically. You are now ready to **register this configuration in the Kentik portal**.
