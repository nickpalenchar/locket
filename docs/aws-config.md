

## Setting up backup sync on aws

### 1. Create a new S3 Bucket

It should have the following config:

- ACLs disabled (default)
- Block All public access (default)
- Bucket Versioning: choose the option best for you
    - `Disable` means lower cost/storage, but no "time machine" like backups (each backup overwrites the last)
    - `Enable` effectively preserves all backups, so costs might be slightly higher
- Server-Side encryption Disable (locket does the encryption)
- Advanced Setting > Object Lock: Enable for stronger backups

### 2. Create an IAM User for locket

- Create a new user, using `Access key - Programmatic access` for **Aws credential type**

#### Sample Policy config

```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "VisualEditor0",
            "Effect": "Allow",
            "Action": [
                "s3:PutObject",
                "s3:GetObject",
                "s3:ListBucket"
            ],
            "Resource": "arn:aws:s3:::name-of-your-bucket"
        }
    ]
}
```
