// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.StorageGateway
{
    /// <summary>
    /// Manages an AWS Storage Gateway SMB File Share.
    /// 
    /// ## Example Usage
    /// ### Active Directory Authentication
    /// 
    /// &gt; **NOTE:** The gateway must have already joined the Active Directory domain prior to SMB file share creationE.g., via "SMB Settings" in the AWS Storage Gateway console or `smb_active_directory_settings` in the `aws.storagegateway.Gateway` resource.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.StorageGateway.SmbFileShare("example", new()
    ///     {
    ///         Authentication = "ActiveDirectory",
    ///         GatewayArn = aws_storagegateway_gateway.Example.Arn,
    ///         LocationArn = aws_s3_bucket.Example.Arn,
    ///         RoleArn = aws_iam_role.Example.Arn,
    ///     });
    /// 
    /// });
    /// ```
    /// ### Guest Authentication
    /// 
    /// &gt; **NOTE:** The gateway must have already had the SMB guest password set prior to SMB file share creationE.g., via "SMB Settings" in the AWS Storage Gateway console or `smb_guest_password` in the `aws.storagegateway.Gateway` resource.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.StorageGateway.SmbFileShare("example", new()
    ///     {
    ///         Authentication = "GuestAccess",
    ///         GatewayArn = aws_storagegateway_gateway.Example.Arn,
    ///         LocationArn = aws_s3_bucket.Example.Arn,
    ///         RoleArn = aws_iam_role.Example.Arn,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Using `pulumi import`, import `aws_storagegateway_smb_file_share` using the SMB File Share Amazon Resource Name (ARN). For example:
    /// 
    /// ```sh
    ///  $ pulumi import aws:storagegateway/smbFileShare:SmbFileShare example arn:aws:storagegateway:us-east-1:123456789012:share/share-12345678
    /// ```
    /// </summary>
    [AwsResourceType("aws:storagegateway/smbFileShare:SmbFileShare")]
    public partial class SmbFileShare : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The files and folders on this share will only be visible to users with read access. Default value is `false`.
        /// </summary>
        [Output("accessBasedEnumeration")]
        public Output<bool?> AccessBasedEnumeration { get; private set; } = null!;

        /// <summary>
        /// A list of users in the Active Directory that have admin access to the file share. Only valid if `authentication` is set to `ActiveDirectory`.
        /// </summary>
        [Output("adminUserLists")]
        public Output<ImmutableArray<string>> AdminUserLists { get; private set; } = null!;

        /// <summary>
        /// Amazon Resource Name (ARN) of the SMB File Share.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the CloudWatch Log Group used for the audit logs.
        /// </summary>
        [Output("auditDestinationArn")]
        public Output<string?> AuditDestinationArn { get; private set; } = null!;

        /// <summary>
        /// The authentication method that users use to access the file share. Defaults to `ActiveDirectory`. Valid values: `ActiveDirectory`, `GuestAccess`.
        /// </summary>
        [Output("authentication")]
        public Output<string?> Authentication { get; private set; } = null!;

        /// <summary>
        /// The region of the S3 buck used by the file share. Required when specifying a `vpc_endpoint_dns_name`.
        /// </summary>
        [Output("bucketRegion")]
        public Output<string?> BucketRegion { get; private set; } = null!;

        /// <summary>
        /// Refresh cache information. see Cache Attributes for more details.
        /// </summary>
        [Output("cacheAttributes")]
        public Output<Outputs.SmbFileShareCacheAttributes?> CacheAttributes { get; private set; } = null!;

        /// <summary>
        /// The case of an object name in an Amazon S3 bucket. For `ClientSpecified`, the client determines the case sensitivity. For `CaseSensitive`, the gateway determines the case sensitivity. The default value is `ClientSpecified`.
        /// </summary>
        [Output("caseSensitivity")]
        public Output<string?> CaseSensitivity { get; private set; } = null!;

        /// <summary>
        /// The default [storage class](https://docs.aws.amazon.com/storagegateway/latest/APIReference/API_CreateNFSFileShare.html#StorageGateway-CreateNFSFileShare-request-DefaultStorageClass) for objects put into an Amazon S3 bucket by the file gateway. Defaults to `S3_STANDARD`.
        /// </summary>
        [Output("defaultStorageClass")]
        public Output<string?> DefaultStorageClass { get; private set; } = null!;

        /// <summary>
        /// The name of the file share. Must be set if an S3 prefix name is set in `location_arn`.
        /// </summary>
        [Output("fileShareName")]
        public Output<string> FileShareName { get; private set; } = null!;

        /// <summary>
        /// ID of the SMB File Share.
        /// </summary>
        [Output("fileshareId")]
        public Output<string> FileshareId { get; private set; } = null!;

        /// <summary>
        /// Amazon Resource Name (ARN) of the file gateway.
        /// </summary>
        [Output("gatewayArn")]
        public Output<string> GatewayArn { get; private set; } = null!;

        /// <summary>
        /// Boolean value that enables guessing of the MIME type for uploaded objects based on file extensions. Defaults to `true`.
        /// </summary>
        [Output("guessMimeTypeEnabled")]
        public Output<bool?> GuessMimeTypeEnabled { get; private set; } = null!;

        /// <summary>
        /// A list of users in the Active Directory that are not allowed to access the file share. Only valid if `authentication` is set to `ActiveDirectory`.
        /// </summary>
        [Output("invalidUserLists")]
        public Output<ImmutableArray<string>> InvalidUserLists { get; private set; } = null!;

        /// <summary>
        /// Boolean value if `true` to use Amazon S3 server side encryption with your own AWS KMS key, or `false` to use a key managed by Amazon S3. Defaults to `false`.
        /// </summary>
        [Output("kmsEncrypted")]
        public Output<bool?> KmsEncrypted { get; private set; } = null!;

        /// <summary>
        /// Amazon Resource Name (ARN) for KMS key used for Amazon S3 server side encryption. This value can only be set when `kms_encrypted` is true.
        /// </summary>
        [Output("kmsKeyArn")]
        public Output<string?> KmsKeyArn { get; private set; } = null!;

        /// <summary>
        /// The ARN of the backed storage used for storing file data.
        /// </summary>
        [Output("locationArn")]
        public Output<string> LocationArn { get; private set; } = null!;

        /// <summary>
        /// The notification policy of the file share. For more information see the [AWS Documentation](https://docs.aws.amazon.com/storagegateway/latest/APIReference/API_CreateNFSFileShare.html#StorageGateway-CreateNFSFileShare-request-NotificationPolicy). Default value is `{}`.
        /// </summary>
        [Output("notificationPolicy")]
        public Output<string?> NotificationPolicy { get; private set; } = null!;

        /// <summary>
        /// Access Control List permission for S3 objects. Defaults to `private`.
        /// </summary>
        [Output("objectAcl")]
        public Output<string?> ObjectAcl { get; private set; } = null!;

        /// <summary>
        /// Boolean to indicate Opportunistic lock (oplock) status. Defaults to `true`.
        /// </summary>
        [Output("oplocksEnabled")]
        public Output<bool> OplocksEnabled { get; private set; } = null!;

        /// <summary>
        /// File share path used by the NFS client to identify the mount point.
        /// </summary>
        [Output("path")]
        public Output<string> Path { get; private set; } = null!;

        /// <summary>
        /// Boolean to indicate write status of file share. File share does not accept writes if `true`. Defaults to `false`.
        /// </summary>
        [Output("readOnly")]
        public Output<bool?> ReadOnly { get; private set; } = null!;

        /// <summary>
        /// Boolean who pays the cost of the request and the data download from the Amazon S3 bucket. Set this value to `true` if you want the requester to pay instead of the bucket owner. Defaults to `false`.
        /// </summary>
        [Output("requesterPays")]
        public Output<bool?> RequesterPays { get; private set; } = null!;

        /// <summary>
        /// The ARN of the AWS Identity and Access Management (IAM) role that a file gateway assumes when it accesses the underlying storage.
        /// </summary>
        [Output("roleArn")]
        public Output<string> RoleArn { get; private set; } = null!;

        /// <summary>
        /// Set this value to `true` to enable ACL (access control list) on the SMB fileshare. Set it to `false` to map file and directory permissions to the POSIX permissions. This setting applies only to `ActiveDirectory` authentication type.
        /// </summary>
        [Output("smbAclEnabled")]
        public Output<bool?> SmbAclEnabled { get; private set; } = null!;

        /// <summary>
        /// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// A list of users in the Active Directory that are allowed to access the file share. If you need to specify an Active directory group, add '@' before the name of the group. It will be set on Allowed group in AWS console. Only valid if `authentication` is set to `ActiveDirectory`.
        /// </summary>
        [Output("validUserLists")]
        public Output<ImmutableArray<string>> ValidUserLists { get; private set; } = null!;

        /// <summary>
        /// The DNS name of the VPC endpoint for S3 private link.
        /// </summary>
        [Output("vpcEndpointDnsName")]
        public Output<string?> VpcEndpointDnsName { get; private set; } = null!;


        /// <summary>
        /// Create a SmbFileShare resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SmbFileShare(string name, SmbFileShareArgs args, CustomResourceOptions? options = null)
            : base("aws:storagegateway/smbFileShare:SmbFileShare", name, args ?? new SmbFileShareArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SmbFileShare(string name, Input<string> id, SmbFileShareState? state = null, CustomResourceOptions? options = null)
            : base("aws:storagegateway/smbFileShare:SmbFileShare", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "tagsAll",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SmbFileShare resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SmbFileShare Get(string name, Input<string> id, SmbFileShareState? state = null, CustomResourceOptions? options = null)
        {
            return new SmbFileShare(name, id, state, options);
        }
    }

    public sealed class SmbFileShareArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The files and folders on this share will only be visible to users with read access. Default value is `false`.
        /// </summary>
        [Input("accessBasedEnumeration")]
        public Input<bool>? AccessBasedEnumeration { get; set; }

        [Input("adminUserLists")]
        private InputList<string>? _adminUserLists;

        /// <summary>
        /// A list of users in the Active Directory that have admin access to the file share. Only valid if `authentication` is set to `ActiveDirectory`.
        /// </summary>
        public InputList<string> AdminUserLists
        {
            get => _adminUserLists ?? (_adminUserLists = new InputList<string>());
            set => _adminUserLists = value;
        }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the CloudWatch Log Group used for the audit logs.
        /// </summary>
        [Input("auditDestinationArn")]
        public Input<string>? AuditDestinationArn { get; set; }

        /// <summary>
        /// The authentication method that users use to access the file share. Defaults to `ActiveDirectory`. Valid values: `ActiveDirectory`, `GuestAccess`.
        /// </summary>
        [Input("authentication")]
        public Input<string>? Authentication { get; set; }

        /// <summary>
        /// The region of the S3 buck used by the file share. Required when specifying a `vpc_endpoint_dns_name`.
        /// </summary>
        [Input("bucketRegion")]
        public Input<string>? BucketRegion { get; set; }

        /// <summary>
        /// Refresh cache information. see Cache Attributes for more details.
        /// </summary>
        [Input("cacheAttributes")]
        public Input<Inputs.SmbFileShareCacheAttributesArgs>? CacheAttributes { get; set; }

        /// <summary>
        /// The case of an object name in an Amazon S3 bucket. For `ClientSpecified`, the client determines the case sensitivity. For `CaseSensitive`, the gateway determines the case sensitivity. The default value is `ClientSpecified`.
        /// </summary>
        [Input("caseSensitivity")]
        public Input<string>? CaseSensitivity { get; set; }

        /// <summary>
        /// The default [storage class](https://docs.aws.amazon.com/storagegateway/latest/APIReference/API_CreateNFSFileShare.html#StorageGateway-CreateNFSFileShare-request-DefaultStorageClass) for objects put into an Amazon S3 bucket by the file gateway. Defaults to `S3_STANDARD`.
        /// </summary>
        [Input("defaultStorageClass")]
        public Input<string>? DefaultStorageClass { get; set; }

        /// <summary>
        /// The name of the file share. Must be set if an S3 prefix name is set in `location_arn`.
        /// </summary>
        [Input("fileShareName")]
        public Input<string>? FileShareName { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN) of the file gateway.
        /// </summary>
        [Input("gatewayArn", required: true)]
        public Input<string> GatewayArn { get; set; } = null!;

        /// <summary>
        /// Boolean value that enables guessing of the MIME type for uploaded objects based on file extensions. Defaults to `true`.
        /// </summary>
        [Input("guessMimeTypeEnabled")]
        public Input<bool>? GuessMimeTypeEnabled { get; set; }

        [Input("invalidUserLists")]
        private InputList<string>? _invalidUserLists;

        /// <summary>
        /// A list of users in the Active Directory that are not allowed to access the file share. Only valid if `authentication` is set to `ActiveDirectory`.
        /// </summary>
        public InputList<string> InvalidUserLists
        {
            get => _invalidUserLists ?? (_invalidUserLists = new InputList<string>());
            set => _invalidUserLists = value;
        }

        /// <summary>
        /// Boolean value if `true` to use Amazon S3 server side encryption with your own AWS KMS key, or `false` to use a key managed by Amazon S3. Defaults to `false`.
        /// </summary>
        [Input("kmsEncrypted")]
        public Input<bool>? KmsEncrypted { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN) for KMS key used for Amazon S3 server side encryption. This value can only be set when `kms_encrypted` is true.
        /// </summary>
        [Input("kmsKeyArn")]
        public Input<string>? KmsKeyArn { get; set; }

        /// <summary>
        /// The ARN of the backed storage used for storing file data.
        /// </summary>
        [Input("locationArn", required: true)]
        public Input<string> LocationArn { get; set; } = null!;

        /// <summary>
        /// The notification policy of the file share. For more information see the [AWS Documentation](https://docs.aws.amazon.com/storagegateway/latest/APIReference/API_CreateNFSFileShare.html#StorageGateway-CreateNFSFileShare-request-NotificationPolicy). Default value is `{}`.
        /// </summary>
        [Input("notificationPolicy")]
        public Input<string>? NotificationPolicy { get; set; }

        /// <summary>
        /// Access Control List permission for S3 objects. Defaults to `private`.
        /// </summary>
        [Input("objectAcl")]
        public Input<string>? ObjectAcl { get; set; }

        /// <summary>
        /// Boolean to indicate Opportunistic lock (oplock) status. Defaults to `true`.
        /// </summary>
        [Input("oplocksEnabled")]
        public Input<bool>? OplocksEnabled { get; set; }

        /// <summary>
        /// Boolean to indicate write status of file share. File share does not accept writes if `true`. Defaults to `false`.
        /// </summary>
        [Input("readOnly")]
        public Input<bool>? ReadOnly { get; set; }

        /// <summary>
        /// Boolean who pays the cost of the request and the data download from the Amazon S3 bucket. Set this value to `true` if you want the requester to pay instead of the bucket owner. Defaults to `false`.
        /// </summary>
        [Input("requesterPays")]
        public Input<bool>? RequesterPays { get; set; }

        /// <summary>
        /// The ARN of the AWS Identity and Access Management (IAM) role that a file gateway assumes when it accesses the underlying storage.
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        /// <summary>
        /// Set this value to `true` to enable ACL (access control list) on the SMB fileshare. Set it to `false` to map file and directory permissions to the POSIX permissions. This setting applies only to `ActiveDirectory` authentication type.
        /// </summary>
        [Input("smbAclEnabled")]
        public Input<bool>? SmbAclEnabled { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("validUserLists")]
        private InputList<string>? _validUserLists;

        /// <summary>
        /// A list of users in the Active Directory that are allowed to access the file share. If you need to specify an Active directory group, add '@' before the name of the group. It will be set on Allowed group in AWS console. Only valid if `authentication` is set to `ActiveDirectory`.
        /// </summary>
        public InputList<string> ValidUserLists
        {
            get => _validUserLists ?? (_validUserLists = new InputList<string>());
            set => _validUserLists = value;
        }

        /// <summary>
        /// The DNS name of the VPC endpoint for S3 private link.
        /// </summary>
        [Input("vpcEndpointDnsName")]
        public Input<string>? VpcEndpointDnsName { get; set; }

        public SmbFileShareArgs()
        {
        }
        public static new SmbFileShareArgs Empty => new SmbFileShareArgs();
    }

    public sealed class SmbFileShareState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The files and folders on this share will only be visible to users with read access. Default value is `false`.
        /// </summary>
        [Input("accessBasedEnumeration")]
        public Input<bool>? AccessBasedEnumeration { get; set; }

        [Input("adminUserLists")]
        private InputList<string>? _adminUserLists;

        /// <summary>
        /// A list of users in the Active Directory that have admin access to the file share. Only valid if `authentication` is set to `ActiveDirectory`.
        /// </summary>
        public InputList<string> AdminUserLists
        {
            get => _adminUserLists ?? (_adminUserLists = new InputList<string>());
            set => _adminUserLists = value;
        }

        /// <summary>
        /// Amazon Resource Name (ARN) of the SMB File Share.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the CloudWatch Log Group used for the audit logs.
        /// </summary>
        [Input("auditDestinationArn")]
        public Input<string>? AuditDestinationArn { get; set; }

        /// <summary>
        /// The authentication method that users use to access the file share. Defaults to `ActiveDirectory`. Valid values: `ActiveDirectory`, `GuestAccess`.
        /// </summary>
        [Input("authentication")]
        public Input<string>? Authentication { get; set; }

        /// <summary>
        /// The region of the S3 buck used by the file share. Required when specifying a `vpc_endpoint_dns_name`.
        /// </summary>
        [Input("bucketRegion")]
        public Input<string>? BucketRegion { get; set; }

        /// <summary>
        /// Refresh cache information. see Cache Attributes for more details.
        /// </summary>
        [Input("cacheAttributes")]
        public Input<Inputs.SmbFileShareCacheAttributesGetArgs>? CacheAttributes { get; set; }

        /// <summary>
        /// The case of an object name in an Amazon S3 bucket. For `ClientSpecified`, the client determines the case sensitivity. For `CaseSensitive`, the gateway determines the case sensitivity. The default value is `ClientSpecified`.
        /// </summary>
        [Input("caseSensitivity")]
        public Input<string>? CaseSensitivity { get; set; }

        /// <summary>
        /// The default [storage class](https://docs.aws.amazon.com/storagegateway/latest/APIReference/API_CreateNFSFileShare.html#StorageGateway-CreateNFSFileShare-request-DefaultStorageClass) for objects put into an Amazon S3 bucket by the file gateway. Defaults to `S3_STANDARD`.
        /// </summary>
        [Input("defaultStorageClass")]
        public Input<string>? DefaultStorageClass { get; set; }

        /// <summary>
        /// The name of the file share. Must be set if an S3 prefix name is set in `location_arn`.
        /// </summary>
        [Input("fileShareName")]
        public Input<string>? FileShareName { get; set; }

        /// <summary>
        /// ID of the SMB File Share.
        /// </summary>
        [Input("fileshareId")]
        public Input<string>? FileshareId { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN) of the file gateway.
        /// </summary>
        [Input("gatewayArn")]
        public Input<string>? GatewayArn { get; set; }

        /// <summary>
        /// Boolean value that enables guessing of the MIME type for uploaded objects based on file extensions. Defaults to `true`.
        /// </summary>
        [Input("guessMimeTypeEnabled")]
        public Input<bool>? GuessMimeTypeEnabled { get; set; }

        [Input("invalidUserLists")]
        private InputList<string>? _invalidUserLists;

        /// <summary>
        /// A list of users in the Active Directory that are not allowed to access the file share. Only valid if `authentication` is set to `ActiveDirectory`.
        /// </summary>
        public InputList<string> InvalidUserLists
        {
            get => _invalidUserLists ?? (_invalidUserLists = new InputList<string>());
            set => _invalidUserLists = value;
        }

        /// <summary>
        /// Boolean value if `true` to use Amazon S3 server side encryption with your own AWS KMS key, or `false` to use a key managed by Amazon S3. Defaults to `false`.
        /// </summary>
        [Input("kmsEncrypted")]
        public Input<bool>? KmsEncrypted { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN) for KMS key used for Amazon S3 server side encryption. This value can only be set when `kms_encrypted` is true.
        /// </summary>
        [Input("kmsKeyArn")]
        public Input<string>? KmsKeyArn { get; set; }

        /// <summary>
        /// The ARN of the backed storage used for storing file data.
        /// </summary>
        [Input("locationArn")]
        public Input<string>? LocationArn { get; set; }

        /// <summary>
        /// The notification policy of the file share. For more information see the [AWS Documentation](https://docs.aws.amazon.com/storagegateway/latest/APIReference/API_CreateNFSFileShare.html#StorageGateway-CreateNFSFileShare-request-NotificationPolicy). Default value is `{}`.
        /// </summary>
        [Input("notificationPolicy")]
        public Input<string>? NotificationPolicy { get; set; }

        /// <summary>
        /// Access Control List permission for S3 objects. Defaults to `private`.
        /// </summary>
        [Input("objectAcl")]
        public Input<string>? ObjectAcl { get; set; }

        /// <summary>
        /// Boolean to indicate Opportunistic lock (oplock) status. Defaults to `true`.
        /// </summary>
        [Input("oplocksEnabled")]
        public Input<bool>? OplocksEnabled { get; set; }

        /// <summary>
        /// File share path used by the NFS client to identify the mount point.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// Boolean to indicate write status of file share. File share does not accept writes if `true`. Defaults to `false`.
        /// </summary>
        [Input("readOnly")]
        public Input<bool>? ReadOnly { get; set; }

        /// <summary>
        /// Boolean who pays the cost of the request and the data download from the Amazon S3 bucket. Set this value to `true` if you want the requester to pay instead of the bucket owner. Defaults to `false`.
        /// </summary>
        [Input("requesterPays")]
        public Input<bool>? RequesterPays { get; set; }

        /// <summary>
        /// The ARN of the AWS Identity and Access Management (IAM) role that a file gateway assumes when it accesses the underlying storage.
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        /// <summary>
        /// Set this value to `true` to enable ACL (access control list) on the SMB fileshare. Set it to `false` to map file and directory permissions to the POSIX permissions. This setting applies only to `ActiveDirectory` authentication type.
        /// </summary>
        [Input("smbAclEnabled")]
        public Input<bool>? SmbAclEnabled { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Obsolete(@"Please use `tags` instead.")]
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableDictionary.Create<string, string>());
                _tagsAll = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        [Input("validUserLists")]
        private InputList<string>? _validUserLists;

        /// <summary>
        /// A list of users in the Active Directory that are allowed to access the file share. If you need to specify an Active directory group, add '@' before the name of the group. It will be set on Allowed group in AWS console. Only valid if `authentication` is set to `ActiveDirectory`.
        /// </summary>
        public InputList<string> ValidUserLists
        {
            get => _validUserLists ?? (_validUserLists = new InputList<string>());
            set => _validUserLists = value;
        }

        /// <summary>
        /// The DNS name of the VPC endpoint for S3 private link.
        /// </summary>
        [Input("vpcEndpointDnsName")]
        public Input<string>? VpcEndpointDnsName { get; set; }

        public SmbFileShareState()
        {
        }
        public static new SmbFileShareState Empty => new SmbFileShareState();
    }
}
