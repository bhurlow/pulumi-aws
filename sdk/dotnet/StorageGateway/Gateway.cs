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
    /// Manages an AWS Storage Gateway file, tape, or volume gateway in the provider region.
    /// 
    /// &gt; **NOTE:** The Storage Gateway API requires the gateway to be connected to properly return information after activation. If you are receiving `The specified gateway is not connected` errors during resource creation (gateway activation), ensure your gateway instance meets the [Storage Gateway requirements](https://docs.aws.amazon.com/storagegateway/latest/userguide/Requirements.html).
    /// 
    /// ## Example Usage
    /// ### Local Cache
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var testVolumeAttachment = new Aws.Ec2.VolumeAttachment("testVolumeAttachment", new()
    ///     {
    ///         DeviceName = "/dev/xvdb",
    ///         VolumeId = aws_ebs_volume.Test.Id,
    ///         InstanceId = aws_instance.Test.Id,
    ///     });
    /// 
    ///     var testLocalDisk = Aws.StorageGateway.GetLocalDisk.Invoke(new()
    ///     {
    ///         DiskNode = testVolumeAttachment.DeviceName,
    ///         GatewayArn = aws_storagegateway_gateway.Test.Arn,
    ///     });
    /// 
    ///     var testCache = new Aws.StorageGateway.Cache("testCache", new()
    ///     {
    ///         DiskId = testLocalDisk.Apply(getLocalDiskResult =&gt; getLocalDiskResult.DiskId),
    ///         GatewayArn = aws_storagegateway_gateway.Test.Arn,
    ///     });
    /// 
    /// });
    /// ```
    /// ### FSx File Gateway
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.StorageGateway.Gateway("example", new()
    ///     {
    ///         GatewayIpAddress = "1.2.3.4",
    ///         GatewayName = "example",
    ///         GatewayTimezone = "GMT",
    ///         GatewayType = "FILE_FSX_SMB",
    ///         SmbActiveDirectorySettings = new Aws.StorageGateway.Inputs.GatewaySmbActiveDirectorySettingsArgs
    ///         {
    ///             DomainName = "corp.example.com",
    ///             Password = "avoid-plaintext-passwords",
    ///             Username = "Admin",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### S3 File Gateway
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.StorageGateway.Gateway("example", new()
    ///     {
    ///         GatewayIpAddress = "1.2.3.4",
    ///         GatewayName = "example",
    ///         GatewayTimezone = "GMT",
    ///         GatewayType = "FILE_S3",
    ///     });
    /// 
    /// });
    /// ```
    /// ### Tape Gateway
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.StorageGateway.Gateway("example", new()
    ///     {
    ///         GatewayIpAddress = "1.2.3.4",
    ///         GatewayName = "example",
    ///         GatewayTimezone = "GMT",
    ///         GatewayType = "VTL",
    ///         MediumChangerType = "AWS-Gateway-VTL",
    ///         TapeDriveType = "IBM-ULT3580-TD5",
    ///     });
    /// 
    /// });
    /// ```
    /// ### Volume Gateway (Cached)
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.StorageGateway.Gateway("example", new()
    ///     {
    ///         GatewayIpAddress = "1.2.3.4",
    ///         GatewayName = "example",
    ///         GatewayTimezone = "GMT",
    ///         GatewayType = "CACHED",
    ///     });
    /// 
    /// });
    /// ```
    /// ### Volume Gateway (Stored)
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.StorageGateway.Gateway("example", new()
    ///     {
    ///         GatewayIpAddress = "1.2.3.4",
    ///         GatewayName = "example",
    ///         GatewayTimezone = "GMT",
    ///         GatewayType = "STORED",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Using `pulumi import`, import `aws_storagegateway_gateway` using the gateway Amazon Resource Name (ARN). For example:
    /// 
    /// ```sh
    ///  $ pulumi import aws:storagegateway/gateway:Gateway example arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678
    /// ```
    ///  Certain resource arguments, like `gateway_ip_address` do not have a Storage Gateway API method for reading the information after creation, either omit the argument from the Pulumi program or use `ignore_changes` to hide the difference. For example:
    /// </summary>
    [AwsResourceType("aws:storagegateway/gateway:Gateway")]
    public partial class Gateway : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Gateway activation key during resource creation. Conflicts with `gateway_ip_address`. Additional information is available in the [Storage Gateway User Guide](https://docs.aws.amazon.com/storagegateway/latest/userguide/get-activation-key.html).
        /// </summary>
        [Output("activationKey")]
        public Output<string> ActivationKey { get; private set; } = null!;

        /// <summary>
        /// Amazon Resource Name (ARN) of the gateway.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The average download bandwidth rate limit in bits per second. This is supported for the `CACHED`, `STORED`, and `VTL` gateway types.
        /// </summary>
        [Output("averageDownloadRateLimitInBitsPerSec")]
        public Output<int?> AverageDownloadRateLimitInBitsPerSec { get; private set; } = null!;

        /// <summary>
        /// The average upload bandwidth rate limit in bits per second. This is supported for the `CACHED`, `STORED`, and `VTL` gateway types.
        /// </summary>
        [Output("averageUploadRateLimitInBitsPerSec")]
        public Output<int?> AverageUploadRateLimitInBitsPerSec { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the Amazon CloudWatch log group to use to monitor and log events in the gateway.
        /// </summary>
        [Output("cloudwatchLogGroupArn")]
        public Output<string?> CloudwatchLogGroupArn { get; private set; } = null!;

        /// <summary>
        /// The ID of the Amazon EC2 instance that was used to launch the gateway.
        /// </summary>
        [Output("ec2InstanceId")]
        public Output<string> Ec2InstanceId { get; private set; } = null!;

        /// <summary>
        /// The type of endpoint for your gateway.
        /// </summary>
        [Output("endpointType")]
        public Output<string> EndpointType { get; private set; } = null!;

        /// <summary>
        /// Identifier of the gateway.
        /// </summary>
        [Output("gatewayId")]
        public Output<string> GatewayId { get; private set; } = null!;

        /// <summary>
        /// Gateway IP address to retrieve activation key during resource creation. Conflicts with `activation_key`. Gateway must be accessible on port 80 from where this provider is running. Additional information is available in the [Storage Gateway User Guide](https://docs.aws.amazon.com/storagegateway/latest/userguide/get-activation-key.html).
        /// </summary>
        [Output("gatewayIpAddress")]
        public Output<string> GatewayIpAddress { get; private set; } = null!;

        /// <summary>
        /// Name of the gateway.
        /// </summary>
        [Output("gatewayName")]
        public Output<string> GatewayName { get; private set; } = null!;

        /// <summary>
        /// An array that contains descriptions of the gateway network interfaces. See Gateway Network Interface.
        /// </summary>
        [Output("gatewayNetworkInterfaces")]
        public Output<ImmutableArray<Outputs.GatewayGatewayNetworkInterface>> GatewayNetworkInterfaces { get; private set; } = null!;

        /// <summary>
        /// Time zone for the gateway. The time zone is of the format "GMT", "GMT-hr:mm", or "GMT+hr:mm". For example, `GMT-4:00` indicates the time is 4 hours behind GMT. The time zone is used, for example, for scheduling snapshots and your gateway's maintenance schedule.
        /// </summary>
        [Output("gatewayTimezone")]
        public Output<string> GatewayTimezone { get; private set; } = null!;

        /// <summary>
        /// Type of the gateway. The default value is `STORED`. Valid values: `CACHED`, `FILE_FSX_SMB`, `FILE_S3`, `STORED`, `VTL`.
        /// </summary>
        [Output("gatewayType")]
        public Output<string?> GatewayType { get; private set; } = null!;

        /// <summary>
        /// VPC endpoint address to be used when activating your gateway. This should be used when your instance is in a private subnet. Requires HTTP access from client computer running this provider. More info on what ports are required by your VPC Endpoint Security group in [Activating a Gateway in a Virtual Private Cloud](https://docs.aws.amazon.com/storagegateway/latest/userguide/gateway-private-link.html).
        /// </summary>
        [Output("gatewayVpcEndpoint")]
        public Output<string?> GatewayVpcEndpoint { get; private set; } = null!;

        /// <summary>
        /// The type of hypervisor environment used by the host.
        /// </summary>
        [Output("hostEnvironment")]
        public Output<string> HostEnvironment { get; private set; } = null!;

        /// <summary>
        /// The gateway's weekly maintenance start time information, including day and time of the week. The maintenance time is the time in your gateway's time zone. More details below.
        /// </summary>
        [Output("maintenanceStartTime")]
        public Output<Outputs.GatewayMaintenanceStartTime> MaintenanceStartTime { get; private set; } = null!;

        /// <summary>
        /// Type of medium changer to use for tape gateway. This provider cannot detect drift of this argument. Valid values: `STK-L700`, `AWS-Gateway-VTL`, `IBM-03584L32-0402`.
        /// </summary>
        [Output("mediumChangerType")]
        public Output<string?> MediumChangerType { get; private set; } = null!;

        /// <summary>
        /// Nested argument with Active Directory domain join information for Server Message Block (SMB) file shares. Only valid for `FILE_S3` and `FILE_FSX_SMB` gateway types. Must be set before creating `ActiveDirectory` authentication SMB file shares. More details below.
        /// </summary>
        [Output("smbActiveDirectorySettings")]
        public Output<Outputs.GatewaySmbActiveDirectorySettings?> SmbActiveDirectorySettings { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the shares on this gateway appear when listing shares.
        /// </summary>
        [Output("smbFileShareVisibility")]
        public Output<bool?> SmbFileShareVisibility { get; private set; } = null!;

        /// <summary>
        /// Guest password for Server Message Block (SMB) file shares. Only valid for `FILE_S3` and `FILE_FSX_SMB` gateway types. Must be set before creating `GuestAccess` authentication SMB file shares. This provider can only detect drift of the existence of a guest password, not its actual value from the gateway. This provider can however update the password with changing the argument.
        /// </summary>
        [Output("smbGuestPassword")]
        public Output<string?> SmbGuestPassword { get; private set; } = null!;

        /// <summary>
        /// Specifies the type of security strategy. Valid values are: `ClientSpecified`, `MandatorySigning`, and `MandatoryEncryption`. See [Setting a Security Level for Your Gateway](https://docs.aws.amazon.com/storagegateway/latest/userguide/managing-gateway-file.html#security-strategy) for more information.
        /// </summary>
        [Output("smbSecurityStrategy")]
        public Output<string> SmbSecurityStrategy { get; private set; } = null!;

        /// <summary>
        /// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// Type of tape drive to use for tape gateway. This provider cannot detect drift of this argument. Valid values: `IBM-ULT3580-TD5`.
        /// </summary>
        [Output("tapeDriveType")]
        public Output<string?> TapeDriveType { get; private set; } = null!;


        /// <summary>
        /// Create a Gateway resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Gateway(string name, GatewayArgs args, CustomResourceOptions? options = null)
            : base("aws:storagegateway/gateway:Gateway", name, args ?? new GatewayArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Gateway(string name, Input<string> id, GatewayState? state = null, CustomResourceOptions? options = null)
            : base("aws:storagegateway/gateway:Gateway", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "smbGuestPassword",
                    "tagsAll",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Gateway resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Gateway Get(string name, Input<string> id, GatewayState? state = null, CustomResourceOptions? options = null)
        {
            return new Gateway(name, id, state, options);
        }
    }

    public sealed class GatewayArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gateway activation key during resource creation. Conflicts with `gateway_ip_address`. Additional information is available in the [Storage Gateway User Guide](https://docs.aws.amazon.com/storagegateway/latest/userguide/get-activation-key.html).
        /// </summary>
        [Input("activationKey")]
        public Input<string>? ActivationKey { get; set; }

        /// <summary>
        /// The average download bandwidth rate limit in bits per second. This is supported for the `CACHED`, `STORED`, and `VTL` gateway types.
        /// </summary>
        [Input("averageDownloadRateLimitInBitsPerSec")]
        public Input<int>? AverageDownloadRateLimitInBitsPerSec { get; set; }

        /// <summary>
        /// The average upload bandwidth rate limit in bits per second. This is supported for the `CACHED`, `STORED`, and `VTL` gateway types.
        /// </summary>
        [Input("averageUploadRateLimitInBitsPerSec")]
        public Input<int>? AverageUploadRateLimitInBitsPerSec { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the Amazon CloudWatch log group to use to monitor and log events in the gateway.
        /// </summary>
        [Input("cloudwatchLogGroupArn")]
        public Input<string>? CloudwatchLogGroupArn { get; set; }

        /// <summary>
        /// Gateway IP address to retrieve activation key during resource creation. Conflicts with `activation_key`. Gateway must be accessible on port 80 from where this provider is running. Additional information is available in the [Storage Gateway User Guide](https://docs.aws.amazon.com/storagegateway/latest/userguide/get-activation-key.html).
        /// </summary>
        [Input("gatewayIpAddress")]
        public Input<string>? GatewayIpAddress { get; set; }

        /// <summary>
        /// Name of the gateway.
        /// </summary>
        [Input("gatewayName", required: true)]
        public Input<string> GatewayName { get; set; } = null!;

        /// <summary>
        /// Time zone for the gateway. The time zone is of the format "GMT", "GMT-hr:mm", or "GMT+hr:mm". For example, `GMT-4:00` indicates the time is 4 hours behind GMT. The time zone is used, for example, for scheduling snapshots and your gateway's maintenance schedule.
        /// </summary>
        [Input("gatewayTimezone", required: true)]
        public Input<string> GatewayTimezone { get; set; } = null!;

        /// <summary>
        /// Type of the gateway. The default value is `STORED`. Valid values: `CACHED`, `FILE_FSX_SMB`, `FILE_S3`, `STORED`, `VTL`.
        /// </summary>
        [Input("gatewayType")]
        public Input<string>? GatewayType { get; set; }

        /// <summary>
        /// VPC endpoint address to be used when activating your gateway. This should be used when your instance is in a private subnet. Requires HTTP access from client computer running this provider. More info on what ports are required by your VPC Endpoint Security group in [Activating a Gateway in a Virtual Private Cloud](https://docs.aws.amazon.com/storagegateway/latest/userguide/gateway-private-link.html).
        /// </summary>
        [Input("gatewayVpcEndpoint")]
        public Input<string>? GatewayVpcEndpoint { get; set; }

        /// <summary>
        /// The gateway's weekly maintenance start time information, including day and time of the week. The maintenance time is the time in your gateway's time zone. More details below.
        /// </summary>
        [Input("maintenanceStartTime")]
        public Input<Inputs.GatewayMaintenanceStartTimeArgs>? MaintenanceStartTime { get; set; }

        /// <summary>
        /// Type of medium changer to use for tape gateway. This provider cannot detect drift of this argument. Valid values: `STK-L700`, `AWS-Gateway-VTL`, `IBM-03584L32-0402`.
        /// </summary>
        [Input("mediumChangerType")]
        public Input<string>? MediumChangerType { get; set; }

        /// <summary>
        /// Nested argument with Active Directory domain join information for Server Message Block (SMB) file shares. Only valid for `FILE_S3` and `FILE_FSX_SMB` gateway types. Must be set before creating `ActiveDirectory` authentication SMB file shares. More details below.
        /// </summary>
        [Input("smbActiveDirectorySettings")]
        public Input<Inputs.GatewaySmbActiveDirectorySettingsArgs>? SmbActiveDirectorySettings { get; set; }

        /// <summary>
        /// Specifies whether the shares on this gateway appear when listing shares.
        /// </summary>
        [Input("smbFileShareVisibility")]
        public Input<bool>? SmbFileShareVisibility { get; set; }

        [Input("smbGuestPassword")]
        private Input<string>? _smbGuestPassword;

        /// <summary>
        /// Guest password for Server Message Block (SMB) file shares. Only valid for `FILE_S3` and `FILE_FSX_SMB` gateway types. Must be set before creating `GuestAccess` authentication SMB file shares. This provider can only detect drift of the existence of a guest password, not its actual value from the gateway. This provider can however update the password with changing the argument.
        /// </summary>
        public Input<string>? SmbGuestPassword
        {
            get => _smbGuestPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _smbGuestPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Specifies the type of security strategy. Valid values are: `ClientSpecified`, `MandatorySigning`, and `MandatoryEncryption`. See [Setting a Security Level for Your Gateway](https://docs.aws.amazon.com/storagegateway/latest/userguide/managing-gateway-file.html#security-strategy) for more information.
        /// </summary>
        [Input("smbSecurityStrategy")]
        public Input<string>? SmbSecurityStrategy { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Type of tape drive to use for tape gateway. This provider cannot detect drift of this argument. Valid values: `IBM-ULT3580-TD5`.
        /// </summary>
        [Input("tapeDriveType")]
        public Input<string>? TapeDriveType { get; set; }

        public GatewayArgs()
        {
        }
        public static new GatewayArgs Empty => new GatewayArgs();
    }

    public sealed class GatewayState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gateway activation key during resource creation. Conflicts with `gateway_ip_address`. Additional information is available in the [Storage Gateway User Guide](https://docs.aws.amazon.com/storagegateway/latest/userguide/get-activation-key.html).
        /// </summary>
        [Input("activationKey")]
        public Input<string>? ActivationKey { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN) of the gateway.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The average download bandwidth rate limit in bits per second. This is supported for the `CACHED`, `STORED`, and `VTL` gateway types.
        /// </summary>
        [Input("averageDownloadRateLimitInBitsPerSec")]
        public Input<int>? AverageDownloadRateLimitInBitsPerSec { get; set; }

        /// <summary>
        /// The average upload bandwidth rate limit in bits per second. This is supported for the `CACHED`, `STORED`, and `VTL` gateway types.
        /// </summary>
        [Input("averageUploadRateLimitInBitsPerSec")]
        public Input<int>? AverageUploadRateLimitInBitsPerSec { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the Amazon CloudWatch log group to use to monitor and log events in the gateway.
        /// </summary>
        [Input("cloudwatchLogGroupArn")]
        public Input<string>? CloudwatchLogGroupArn { get; set; }

        /// <summary>
        /// The ID of the Amazon EC2 instance that was used to launch the gateway.
        /// </summary>
        [Input("ec2InstanceId")]
        public Input<string>? Ec2InstanceId { get; set; }

        /// <summary>
        /// The type of endpoint for your gateway.
        /// </summary>
        [Input("endpointType")]
        public Input<string>? EndpointType { get; set; }

        /// <summary>
        /// Identifier of the gateway.
        /// </summary>
        [Input("gatewayId")]
        public Input<string>? GatewayId { get; set; }

        /// <summary>
        /// Gateway IP address to retrieve activation key during resource creation. Conflicts with `activation_key`. Gateway must be accessible on port 80 from where this provider is running. Additional information is available in the [Storage Gateway User Guide](https://docs.aws.amazon.com/storagegateway/latest/userguide/get-activation-key.html).
        /// </summary>
        [Input("gatewayIpAddress")]
        public Input<string>? GatewayIpAddress { get; set; }

        /// <summary>
        /// Name of the gateway.
        /// </summary>
        [Input("gatewayName")]
        public Input<string>? GatewayName { get; set; }

        [Input("gatewayNetworkInterfaces")]
        private InputList<Inputs.GatewayGatewayNetworkInterfaceGetArgs>? _gatewayNetworkInterfaces;

        /// <summary>
        /// An array that contains descriptions of the gateway network interfaces. See Gateway Network Interface.
        /// </summary>
        public InputList<Inputs.GatewayGatewayNetworkInterfaceGetArgs> GatewayNetworkInterfaces
        {
            get => _gatewayNetworkInterfaces ?? (_gatewayNetworkInterfaces = new InputList<Inputs.GatewayGatewayNetworkInterfaceGetArgs>());
            set => _gatewayNetworkInterfaces = value;
        }

        /// <summary>
        /// Time zone for the gateway. The time zone is of the format "GMT", "GMT-hr:mm", or "GMT+hr:mm". For example, `GMT-4:00` indicates the time is 4 hours behind GMT. The time zone is used, for example, for scheduling snapshots and your gateway's maintenance schedule.
        /// </summary>
        [Input("gatewayTimezone")]
        public Input<string>? GatewayTimezone { get; set; }

        /// <summary>
        /// Type of the gateway. The default value is `STORED`. Valid values: `CACHED`, `FILE_FSX_SMB`, `FILE_S3`, `STORED`, `VTL`.
        /// </summary>
        [Input("gatewayType")]
        public Input<string>? GatewayType { get; set; }

        /// <summary>
        /// VPC endpoint address to be used when activating your gateway. This should be used when your instance is in a private subnet. Requires HTTP access from client computer running this provider. More info on what ports are required by your VPC Endpoint Security group in [Activating a Gateway in a Virtual Private Cloud](https://docs.aws.amazon.com/storagegateway/latest/userguide/gateway-private-link.html).
        /// </summary>
        [Input("gatewayVpcEndpoint")]
        public Input<string>? GatewayVpcEndpoint { get; set; }

        /// <summary>
        /// The type of hypervisor environment used by the host.
        /// </summary>
        [Input("hostEnvironment")]
        public Input<string>? HostEnvironment { get; set; }

        /// <summary>
        /// The gateway's weekly maintenance start time information, including day and time of the week. The maintenance time is the time in your gateway's time zone. More details below.
        /// </summary>
        [Input("maintenanceStartTime")]
        public Input<Inputs.GatewayMaintenanceStartTimeGetArgs>? MaintenanceStartTime { get; set; }

        /// <summary>
        /// Type of medium changer to use for tape gateway. This provider cannot detect drift of this argument. Valid values: `STK-L700`, `AWS-Gateway-VTL`, `IBM-03584L32-0402`.
        /// </summary>
        [Input("mediumChangerType")]
        public Input<string>? MediumChangerType { get; set; }

        /// <summary>
        /// Nested argument with Active Directory domain join information for Server Message Block (SMB) file shares. Only valid for `FILE_S3` and `FILE_FSX_SMB` gateway types. Must be set before creating `ActiveDirectory` authentication SMB file shares. More details below.
        /// </summary>
        [Input("smbActiveDirectorySettings")]
        public Input<Inputs.GatewaySmbActiveDirectorySettingsGetArgs>? SmbActiveDirectorySettings { get; set; }

        /// <summary>
        /// Specifies whether the shares on this gateway appear when listing shares.
        /// </summary>
        [Input("smbFileShareVisibility")]
        public Input<bool>? SmbFileShareVisibility { get; set; }

        [Input("smbGuestPassword")]
        private Input<string>? _smbGuestPassword;

        /// <summary>
        /// Guest password for Server Message Block (SMB) file shares. Only valid for `FILE_S3` and `FILE_FSX_SMB` gateway types. Must be set before creating `GuestAccess` authentication SMB file shares. This provider can only detect drift of the existence of a guest password, not its actual value from the gateway. This provider can however update the password with changing the argument.
        /// </summary>
        public Input<string>? SmbGuestPassword
        {
            get => _smbGuestPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _smbGuestPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Specifies the type of security strategy. Valid values are: `ClientSpecified`, `MandatorySigning`, and `MandatoryEncryption`. See [Setting a Security Level for Your Gateway](https://docs.aws.amazon.com/storagegateway/latest/userguide/managing-gateway-file.html#security-strategy) for more information.
        /// </summary>
        [Input("smbSecurityStrategy")]
        public Input<string>? SmbSecurityStrategy { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

        /// <summary>
        /// Type of tape drive to use for tape gateway. This provider cannot detect drift of this argument. Valid values: `IBM-ULT3580-TD5`.
        /// </summary>
        [Input("tapeDriveType")]
        public Input<string>? TapeDriveType { get; set; }

        public GatewayState()
        {
        }
        public static new GatewayState Empty => new GatewayState();
    }
}
