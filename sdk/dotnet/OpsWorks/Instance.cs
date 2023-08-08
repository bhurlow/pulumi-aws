// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.OpsWorks
{
    /// <summary>
    /// Provides an OpsWorks instance resource.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var my_instance = new Aws.OpsWorks.Instance("my-instance", new()
    ///     {
    ///         StackId = aws_opsworks_stack.Main.Id,
    ///         LayerIds = new[]
    ///         {
    ///             aws_opsworks_custom_layer.My_layer.Id,
    ///         },
    ///         InstanceType = "t2.micro",
    ///         Os = "Amazon Linux 2015.09",
    ///         State = "stopped",
    ///     });
    /// 
    /// });
    /// ```
    /// ## Block devices
    /// 
    /// Each of the `*_block_device` attributes controls a portion of the AWS
    /// Instance's "Block Device Mapping". It's a good idea to familiarize yourself with [AWS's Block Device
    /// Mapping docs](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html)
    /// to understand the implications of using these attributes.
    /// 
    /// ### `ebs_block_device`
    /// 
    /// * `delete_on_termination` - (Optional) Whether the volume should be destroyed on instance termination. Default is `true`.
    /// * `device_name` - (Required) Name of the device to mount.
    /// * `iops` - (Optional) Amount of provisioned [IOPS](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-io-characteristics.html). This must be set with a `volume_type` of `io1`.
    /// * `snapshot_id` - (Optional) Snapshot ID to mount.
    /// * `volume_size` - (Optional) Size of the volume in gigabytes.
    /// * `volume_type` - (Optional) Type of volume. Valid values are `standard`, `gp2`, or `io1`. Default is `standard`.
    /// 
    /// Modifying any `ebs_block_device` currently requires resource replacement.
    /// 
    /// ### `ephemeral_block_device`
    /// 
    /// * `device_name` - Name of the block device to mount on the instance.
    /// * `virtual_name` - The [Instance Store Device Name](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html#InstanceStoreDeviceNames) (e.g., `ephemeral0`).
    /// 
    /// Each AWS Instance type has a different set of Instance Store block devices
    /// available for attachment. AWS [publishes a
    /// list](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html#StorageOnInstanceTypes)
    /// of which ephemeral devices are available on each type. The devices are always
    /// identified by the `virtual_name` in the format `ephemeral{0..N}`.
    /// 
    /// ### `root_block_device`
    /// 
    /// * `delete_on_termination` - (Optional) Whether the volume should be destroyed on instance termination. Default is `true`.
    /// * `iops` - (Optional) Amount of provisioned [IOPS](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-io-characteristics.html). This must be set with a `volume_type` of `io1`.
    /// * `volume_size` - (Optional) Size of the volume in gigabytes.
    /// * `volume_type` - (Optional) Type of volume. Valid values are `standard`, `gp2`, or `io1`. Default is `standard`.
    /// 
    /// Modifying any of the `root_block_device` settings requires resource
    /// replacement.
    /// 
    /// &gt; **NOTE:** Currently, changes to `*_block_device` configuration of _existing_
    /// resources cannot be automatically detected by this provider. After making updates
    /// to block device configuration, resource recreation can be manually triggered by
    /// using the [`up` command with the --replace argument](https://www.pulumi.com/docs/reference/cli/pulumi_up/).
    /// 
    /// ## Import
    /// 
    /// terraform import {
    /// 
    ///  to = aws_opsworks_instance.my_instance
    /// 
    ///  id = "4d6d1710-ded9-42a1-b08e-b043ad7af1e2" } Using `pulumi import`, import Opsworks Instances using the instance `id`. For exampleconsole % pulumi import aws_opsworks_instance.my_instance 4d6d1710-ded9-42a1-b08e-b043ad7af1e2
    /// </summary>
    [AwsResourceType("aws:opsworks/instance:Instance")]
    public partial class Instance : global::Pulumi.CustomResource
    {
        /// <summary>
        /// OpsWorks agent to install. Default is `INHERIT`.
        /// </summary>
        [Output("agentVersion")]
        public Output<string?> AgentVersion { get; private set; } = null!;

        /// <summary>
        /// AMI to use for the instance.  If an AMI is specified, `os` must be `Custom`.
        /// </summary>
        [Output("amiId")]
        public Output<string> AmiId { get; private set; } = null!;

        /// <summary>
        /// Machine architecture for created instances.  Valid values are `x86_64` or `i386`. The default is `x86_64`.
        /// </summary>
        [Output("architecture")]
        public Output<string?> Architecture { get; private set; } = null!;

        /// <summary>
        /// Creates load-based or time-based instances.  Valid values are `load`, `timer`.
        /// </summary>
        [Output("autoScalingType")]
        public Output<string?> AutoScalingType { get; private set; } = null!;

        /// <summary>
        /// Name of the availability zone where instances will be created by default.
        /// </summary>
        [Output("availabilityZone")]
        public Output<string> AvailabilityZone { get; private set; } = null!;

        /// <summary>
        /// Time that the instance was created.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Whether to delete EBS volume on deletion. Default is `true`.
        /// </summary>
        [Output("deleteEbs")]
        public Output<bool?> DeleteEbs { get; private set; } = null!;

        /// <summary>
        /// Whether to delete the Elastic IP on deletion.
        /// </summary>
        [Output("deleteEip")]
        public Output<bool?> DeleteEip { get; private set; } = null!;

        /// <summary>
        /// Configuration block for additional EBS block devices to attach to the instance. See Block Devices below.
        /// </summary>
        [Output("ebsBlockDevices")]
        public Output<ImmutableArray<Outputs.InstanceEbsBlockDevice>> EbsBlockDevices { get; private set; } = null!;

        /// <summary>
        /// Whether the launched EC2 instance will be EBS-optimized.
        /// </summary>
        [Output("ebsOptimized")]
        public Output<bool?> EbsOptimized { get; private set; } = null!;

        /// <summary>
        /// EC2 instance ID.
        /// </summary>
        [Output("ec2InstanceId")]
        public Output<string> Ec2InstanceId { get; private set; } = null!;

        /// <summary>
        /// ECS cluster's ARN for container instances.
        /// </summary>
        [Output("ecsClusterArn")]
        public Output<string> EcsClusterArn { get; private set; } = null!;

        /// <summary>
        /// Instance Elastic IP address.
        /// </summary>
        [Output("elasticIp")]
        public Output<string> ElasticIp { get; private set; } = null!;

        /// <summary>
        /// Configuration block for ephemeral (also known as "Instance Store") volumes on the instance. See Block Devices below.
        /// </summary>
        [Output("ephemeralBlockDevices")]
        public Output<ImmutableArray<Outputs.InstanceEphemeralBlockDevice>> EphemeralBlockDevices { get; private set; } = null!;

        /// <summary>
        /// Instance's host name.
        /// </summary>
        [Output("hostname")]
        public Output<string> Hostname { get; private set; } = null!;

        /// <summary>
        /// For registered instances, infrastructure class: ec2 or on-premises.
        /// </summary>
        [Output("infrastructureClass")]
        public Output<string> InfrastructureClass { get; private set; } = null!;

        /// <summary>
        /// Controls where to install OS and package updates when the instance boots.  Default is `true`.
        /// </summary>
        [Output("installUpdatesOnBoot")]
        public Output<bool?> InstallUpdatesOnBoot { get; private set; } = null!;

        /// <summary>
        /// ARN of the instance's IAM profile.
        /// </summary>
        [Output("instanceProfileArn")]
        public Output<string> InstanceProfileArn { get; private set; } = null!;

        /// <summary>
        /// Type of instance to start.
        /// </summary>
        [Output("instanceType")]
        public Output<string?> InstanceType { get; private set; } = null!;

        /// <summary>
        /// ID of the last service error.
        /// </summary>
        [Output("lastServiceErrorId")]
        public Output<string> LastServiceErrorId { get; private set; } = null!;

        /// <summary>
        /// List of the layers the instance will belong to.
        /// </summary>
        [Output("layerIds")]
        public Output<ImmutableArray<string>> LayerIds { get; private set; } = null!;

        /// <summary>
        /// Name of operating system that will be installed.
        /// </summary>
        [Output("os")]
        public Output<string> Os { get; private set; } = null!;

        /// <summary>
        /// Instance's platform.
        /// </summary>
        [Output("platform")]
        public Output<string> Platform { get; private set; } = null!;

        /// <summary>
        /// Private DNS name assigned to the instance. Can only be used inside the Amazon EC2, and only available if you've enabled DNS hostnames for your VPC.
        /// </summary>
        [Output("privateDns")]
        public Output<string> PrivateDns { get; private set; } = null!;

        /// <summary>
        /// Private IP address assigned to the instance.
        /// </summary>
        [Output("privateIp")]
        public Output<string> PrivateIp { get; private set; } = null!;

        /// <summary>
        /// Public DNS name assigned to the instance. For EC2-VPC, this is only available if you've enabled DNS hostnames for your VPC.
        /// </summary>
        [Output("publicDns")]
        public Output<string> PublicDns { get; private set; } = null!;

        /// <summary>
        /// Public IP address assigned to the instance, if applicable.
        /// </summary>
        [Output("publicIp")]
        public Output<string> PublicIp { get; private set; } = null!;

        /// <summary>
        /// For registered instances, who performed the registration.
        /// </summary>
        [Output("registeredBy")]
        public Output<string> RegisteredBy { get; private set; } = null!;

        /// <summary>
        /// Instance's reported AWS OpsWorks Stacks agent version.
        /// </summary>
        [Output("reportedAgentVersion")]
        public Output<string> ReportedAgentVersion { get; private set; } = null!;

        /// <summary>
        /// For registered instances, the reported operating system family.
        /// </summary>
        [Output("reportedOsFamily")]
        public Output<string> ReportedOsFamily { get; private set; } = null!;

        /// <summary>
        /// For registered instances, the reported operating system name.
        /// </summary>
        [Output("reportedOsName")]
        public Output<string> ReportedOsName { get; private set; } = null!;

        /// <summary>
        /// For registered instances, the reported operating system version.
        /// </summary>
        [Output("reportedOsVersion")]
        public Output<string> ReportedOsVersion { get; private set; } = null!;

        /// <summary>
        /// Configuration block for the root block device of the instance. See Block Devices below.
        /// </summary>
        [Output("rootBlockDevices")]
        public Output<ImmutableArray<Outputs.InstanceRootBlockDevice>> RootBlockDevices { get; private set; } = null!;

        /// <summary>
        /// Name of the type of root device instances will have by default. Valid values are `ebs` or `instance-store`.
        /// </summary>
        [Output("rootDeviceType")]
        public Output<string> RootDeviceType { get; private set; } = null!;

        /// <summary>
        /// Root device volume ID.
        /// </summary>
        [Output("rootDeviceVolumeId")]
        public Output<string> RootDeviceVolumeId { get; private set; } = null!;

        /// <summary>
        /// Associated security groups.
        /// </summary>
        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;

        /// <summary>
        /// SSH key's Deep Security Agent (DSA) fingerprint.
        /// </summary>
        [Output("sshHostDsaKeyFingerprint")]
        public Output<string> SshHostDsaKeyFingerprint { get; private set; } = null!;

        /// <summary>
        /// SSH key's RSA fingerprint.
        /// </summary>
        [Output("sshHostRsaKeyFingerprint")]
        public Output<string> SshHostRsaKeyFingerprint { get; private set; } = null!;

        /// <summary>
        /// Name of the SSH keypair that instances will have by default.
        /// </summary>
        [Output("sshKeyName")]
        public Output<string> SshKeyName { get; private set; } = null!;

        /// <summary>
        /// Identifier of the stack the instance will belong to.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Output("stackId")]
        public Output<string> StackId { get; private set; } = null!;

        /// <summary>
        /// Desired state of the instance. Valid values are `running` or `stopped`.
        /// </summary>
        [Output("state")]
        public Output<string?> State { get; private set; } = null!;

        /// <summary>
        /// Instance status. Will be one of `booting`, `connection_lost`, `online`, `pending`, `rebooting`, `requested`, `running_setup`, `setup_failed`, `shutting_down`, `start_failed`, `stop_failed`, `stopped`, `stopping`, `terminated`, or `terminating`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Subnet ID to attach to.
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;

        /// <summary>
        /// Instance tenancy to use. Valid values are `default`, `dedicated` or `host`.
        /// </summary>
        [Output("tenancy")]
        public Output<string> Tenancy { get; private set; } = null!;

        /// <summary>
        /// Keyword to choose what virtualization mode created instances will use. Valid values are `paravirtual` or `hvm`.
        /// </summary>
        [Output("virtualizationType")]
        public Output<string> VirtualizationType { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs args, CustomResourceOptions? options = null)
            : base("aws:opsworks/instance:Instance", name, args ?? new InstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
            : base("aws:opsworks/instance:Instance", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Instance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Instance Get(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new Instance(name, id, state, options);
        }
    }

    public sealed class InstanceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// OpsWorks agent to install. Default is `INHERIT`.
        /// </summary>
        [Input("agentVersion")]
        public Input<string>? AgentVersion { get; set; }

        /// <summary>
        /// AMI to use for the instance.  If an AMI is specified, `os` must be `Custom`.
        /// </summary>
        [Input("amiId")]
        public Input<string>? AmiId { get; set; }

        /// <summary>
        /// Machine architecture for created instances.  Valid values are `x86_64` or `i386`. The default is `x86_64`.
        /// </summary>
        [Input("architecture")]
        public Input<string>? Architecture { get; set; }

        /// <summary>
        /// Creates load-based or time-based instances.  Valid values are `load`, `timer`.
        /// </summary>
        [Input("autoScalingType")]
        public Input<string>? AutoScalingType { get; set; }

        /// <summary>
        /// Name of the availability zone where instances will be created by default.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// Time that the instance was created.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Whether to delete EBS volume on deletion. Default is `true`.
        /// </summary>
        [Input("deleteEbs")]
        public Input<bool>? DeleteEbs { get; set; }

        /// <summary>
        /// Whether to delete the Elastic IP on deletion.
        /// </summary>
        [Input("deleteEip")]
        public Input<bool>? DeleteEip { get; set; }

        [Input("ebsBlockDevices")]
        private InputList<Inputs.InstanceEbsBlockDeviceArgs>? _ebsBlockDevices;

        /// <summary>
        /// Configuration block for additional EBS block devices to attach to the instance. See Block Devices below.
        /// </summary>
        public InputList<Inputs.InstanceEbsBlockDeviceArgs> EbsBlockDevices
        {
            get => _ebsBlockDevices ?? (_ebsBlockDevices = new InputList<Inputs.InstanceEbsBlockDeviceArgs>());
            set => _ebsBlockDevices = value;
        }

        /// <summary>
        /// Whether the launched EC2 instance will be EBS-optimized.
        /// </summary>
        [Input("ebsOptimized")]
        public Input<bool>? EbsOptimized { get; set; }

        /// <summary>
        /// ECS cluster's ARN for container instances.
        /// </summary>
        [Input("ecsClusterArn")]
        public Input<string>? EcsClusterArn { get; set; }

        /// <summary>
        /// Instance Elastic IP address.
        /// </summary>
        [Input("elasticIp")]
        public Input<string>? ElasticIp { get; set; }

        [Input("ephemeralBlockDevices")]
        private InputList<Inputs.InstanceEphemeralBlockDeviceArgs>? _ephemeralBlockDevices;

        /// <summary>
        /// Configuration block for ephemeral (also known as "Instance Store") volumes on the instance. See Block Devices below.
        /// </summary>
        public InputList<Inputs.InstanceEphemeralBlockDeviceArgs> EphemeralBlockDevices
        {
            get => _ephemeralBlockDevices ?? (_ephemeralBlockDevices = new InputList<Inputs.InstanceEphemeralBlockDeviceArgs>());
            set => _ephemeralBlockDevices = value;
        }

        /// <summary>
        /// Instance's host name.
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// For registered instances, infrastructure class: ec2 or on-premises.
        /// </summary>
        [Input("infrastructureClass")]
        public Input<string>? InfrastructureClass { get; set; }

        /// <summary>
        /// Controls where to install OS and package updates when the instance boots.  Default is `true`.
        /// </summary>
        [Input("installUpdatesOnBoot")]
        public Input<bool>? InstallUpdatesOnBoot { get; set; }

        /// <summary>
        /// ARN of the instance's IAM profile.
        /// </summary>
        [Input("instanceProfileArn")]
        public Input<string>? InstanceProfileArn { get; set; }

        /// <summary>
        /// Type of instance to start.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        [Input("layerIds", required: true)]
        private InputList<string>? _layerIds;

        /// <summary>
        /// List of the layers the instance will belong to.
        /// </summary>
        public InputList<string> LayerIds
        {
            get => _layerIds ?? (_layerIds = new InputList<string>());
            set => _layerIds = value;
        }

        /// <summary>
        /// Name of operating system that will be installed.
        /// </summary>
        [Input("os")]
        public Input<string>? Os { get; set; }

        [Input("rootBlockDevices")]
        private InputList<Inputs.InstanceRootBlockDeviceArgs>? _rootBlockDevices;

        /// <summary>
        /// Configuration block for the root block device of the instance. See Block Devices below.
        /// </summary>
        public InputList<Inputs.InstanceRootBlockDeviceArgs> RootBlockDevices
        {
            get => _rootBlockDevices ?? (_rootBlockDevices = new InputList<Inputs.InstanceRootBlockDeviceArgs>());
            set => _rootBlockDevices = value;
        }

        /// <summary>
        /// Name of the type of root device instances will have by default. Valid values are `ebs` or `instance-store`.
        /// </summary>
        [Input("rootDeviceType")]
        public Input<string>? RootDeviceType { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// Associated security groups.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// Name of the SSH keypair that instances will have by default.
        /// </summary>
        [Input("sshKeyName")]
        public Input<string>? SshKeyName { get; set; }

        /// <summary>
        /// Identifier of the stack the instance will belong to.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Input("stackId", required: true)]
        public Input<string> StackId { get; set; } = null!;

        /// <summary>
        /// Desired state of the instance. Valid values are `running` or `stopped`.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// Instance status. Will be one of `booting`, `connection_lost`, `online`, `pending`, `rebooting`, `requested`, `running_setup`, `setup_failed`, `shutting_down`, `start_failed`, `stop_failed`, `stopped`, `stopping`, `terminated`, or `terminating`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Subnet ID to attach to.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// Instance tenancy to use. Valid values are `default`, `dedicated` or `host`.
        /// </summary>
        [Input("tenancy")]
        public Input<string>? Tenancy { get; set; }

        /// <summary>
        /// Keyword to choose what virtualization mode created instances will use. Valid values are `paravirtual` or `hvm`.
        /// </summary>
        [Input("virtualizationType")]
        public Input<string>? VirtualizationType { get; set; }

        public InstanceArgs()
        {
        }
        public static new InstanceArgs Empty => new InstanceArgs();
    }

    public sealed class InstanceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// OpsWorks agent to install. Default is `INHERIT`.
        /// </summary>
        [Input("agentVersion")]
        public Input<string>? AgentVersion { get; set; }

        /// <summary>
        /// AMI to use for the instance.  If an AMI is specified, `os` must be `Custom`.
        /// </summary>
        [Input("amiId")]
        public Input<string>? AmiId { get; set; }

        /// <summary>
        /// Machine architecture for created instances.  Valid values are `x86_64` or `i386`. The default is `x86_64`.
        /// </summary>
        [Input("architecture")]
        public Input<string>? Architecture { get; set; }

        /// <summary>
        /// Creates load-based or time-based instances.  Valid values are `load`, `timer`.
        /// </summary>
        [Input("autoScalingType")]
        public Input<string>? AutoScalingType { get; set; }

        /// <summary>
        /// Name of the availability zone where instances will be created by default.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// Time that the instance was created.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Whether to delete EBS volume on deletion. Default is `true`.
        /// </summary>
        [Input("deleteEbs")]
        public Input<bool>? DeleteEbs { get; set; }

        /// <summary>
        /// Whether to delete the Elastic IP on deletion.
        /// </summary>
        [Input("deleteEip")]
        public Input<bool>? DeleteEip { get; set; }

        [Input("ebsBlockDevices")]
        private InputList<Inputs.InstanceEbsBlockDeviceGetArgs>? _ebsBlockDevices;

        /// <summary>
        /// Configuration block for additional EBS block devices to attach to the instance. See Block Devices below.
        /// </summary>
        public InputList<Inputs.InstanceEbsBlockDeviceGetArgs> EbsBlockDevices
        {
            get => _ebsBlockDevices ?? (_ebsBlockDevices = new InputList<Inputs.InstanceEbsBlockDeviceGetArgs>());
            set => _ebsBlockDevices = value;
        }

        /// <summary>
        /// Whether the launched EC2 instance will be EBS-optimized.
        /// </summary>
        [Input("ebsOptimized")]
        public Input<bool>? EbsOptimized { get; set; }

        /// <summary>
        /// EC2 instance ID.
        /// </summary>
        [Input("ec2InstanceId")]
        public Input<string>? Ec2InstanceId { get; set; }

        /// <summary>
        /// ECS cluster's ARN for container instances.
        /// </summary>
        [Input("ecsClusterArn")]
        public Input<string>? EcsClusterArn { get; set; }

        /// <summary>
        /// Instance Elastic IP address.
        /// </summary>
        [Input("elasticIp")]
        public Input<string>? ElasticIp { get; set; }

        [Input("ephemeralBlockDevices")]
        private InputList<Inputs.InstanceEphemeralBlockDeviceGetArgs>? _ephemeralBlockDevices;

        /// <summary>
        /// Configuration block for ephemeral (also known as "Instance Store") volumes on the instance. See Block Devices below.
        /// </summary>
        public InputList<Inputs.InstanceEphemeralBlockDeviceGetArgs> EphemeralBlockDevices
        {
            get => _ephemeralBlockDevices ?? (_ephemeralBlockDevices = new InputList<Inputs.InstanceEphemeralBlockDeviceGetArgs>());
            set => _ephemeralBlockDevices = value;
        }

        /// <summary>
        /// Instance's host name.
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// For registered instances, infrastructure class: ec2 or on-premises.
        /// </summary>
        [Input("infrastructureClass")]
        public Input<string>? InfrastructureClass { get; set; }

        /// <summary>
        /// Controls where to install OS and package updates when the instance boots.  Default is `true`.
        /// </summary>
        [Input("installUpdatesOnBoot")]
        public Input<bool>? InstallUpdatesOnBoot { get; set; }

        /// <summary>
        /// ARN of the instance's IAM profile.
        /// </summary>
        [Input("instanceProfileArn")]
        public Input<string>? InstanceProfileArn { get; set; }

        /// <summary>
        /// Type of instance to start.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// ID of the last service error.
        /// </summary>
        [Input("lastServiceErrorId")]
        public Input<string>? LastServiceErrorId { get; set; }

        [Input("layerIds")]
        private InputList<string>? _layerIds;

        /// <summary>
        /// List of the layers the instance will belong to.
        /// </summary>
        public InputList<string> LayerIds
        {
            get => _layerIds ?? (_layerIds = new InputList<string>());
            set => _layerIds = value;
        }

        /// <summary>
        /// Name of operating system that will be installed.
        /// </summary>
        [Input("os")]
        public Input<string>? Os { get; set; }

        /// <summary>
        /// Instance's platform.
        /// </summary>
        [Input("platform")]
        public Input<string>? Platform { get; set; }

        /// <summary>
        /// Private DNS name assigned to the instance. Can only be used inside the Amazon EC2, and only available if you've enabled DNS hostnames for your VPC.
        /// </summary>
        [Input("privateDns")]
        public Input<string>? PrivateDns { get; set; }

        /// <summary>
        /// Private IP address assigned to the instance.
        /// </summary>
        [Input("privateIp")]
        public Input<string>? PrivateIp { get; set; }

        /// <summary>
        /// Public DNS name assigned to the instance. For EC2-VPC, this is only available if you've enabled DNS hostnames for your VPC.
        /// </summary>
        [Input("publicDns")]
        public Input<string>? PublicDns { get; set; }

        /// <summary>
        /// Public IP address assigned to the instance, if applicable.
        /// </summary>
        [Input("publicIp")]
        public Input<string>? PublicIp { get; set; }

        /// <summary>
        /// For registered instances, who performed the registration.
        /// </summary>
        [Input("registeredBy")]
        public Input<string>? RegisteredBy { get; set; }

        /// <summary>
        /// Instance's reported AWS OpsWorks Stacks agent version.
        /// </summary>
        [Input("reportedAgentVersion")]
        public Input<string>? ReportedAgentVersion { get; set; }

        /// <summary>
        /// For registered instances, the reported operating system family.
        /// </summary>
        [Input("reportedOsFamily")]
        public Input<string>? ReportedOsFamily { get; set; }

        /// <summary>
        /// For registered instances, the reported operating system name.
        /// </summary>
        [Input("reportedOsName")]
        public Input<string>? ReportedOsName { get; set; }

        /// <summary>
        /// For registered instances, the reported operating system version.
        /// </summary>
        [Input("reportedOsVersion")]
        public Input<string>? ReportedOsVersion { get; set; }

        [Input("rootBlockDevices")]
        private InputList<Inputs.InstanceRootBlockDeviceGetArgs>? _rootBlockDevices;

        /// <summary>
        /// Configuration block for the root block device of the instance. See Block Devices below.
        /// </summary>
        public InputList<Inputs.InstanceRootBlockDeviceGetArgs> RootBlockDevices
        {
            get => _rootBlockDevices ?? (_rootBlockDevices = new InputList<Inputs.InstanceRootBlockDeviceGetArgs>());
            set => _rootBlockDevices = value;
        }

        /// <summary>
        /// Name of the type of root device instances will have by default. Valid values are `ebs` or `instance-store`.
        /// </summary>
        [Input("rootDeviceType")]
        public Input<string>? RootDeviceType { get; set; }

        /// <summary>
        /// Root device volume ID.
        /// </summary>
        [Input("rootDeviceVolumeId")]
        public Input<string>? RootDeviceVolumeId { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// Associated security groups.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// SSH key's Deep Security Agent (DSA) fingerprint.
        /// </summary>
        [Input("sshHostDsaKeyFingerprint")]
        public Input<string>? SshHostDsaKeyFingerprint { get; set; }

        /// <summary>
        /// SSH key's RSA fingerprint.
        /// </summary>
        [Input("sshHostRsaKeyFingerprint")]
        public Input<string>? SshHostRsaKeyFingerprint { get; set; }

        /// <summary>
        /// Name of the SSH keypair that instances will have by default.
        /// </summary>
        [Input("sshKeyName")]
        public Input<string>? SshKeyName { get; set; }

        /// <summary>
        /// Identifier of the stack the instance will belong to.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Input("stackId")]
        public Input<string>? StackId { get; set; }

        /// <summary>
        /// Desired state of the instance. Valid values are `running` or `stopped`.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// Instance status. Will be one of `booting`, `connection_lost`, `online`, `pending`, `rebooting`, `requested`, `running_setup`, `setup_failed`, `shutting_down`, `start_failed`, `stop_failed`, `stopped`, `stopping`, `terminated`, or `terminating`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Subnet ID to attach to.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// Instance tenancy to use. Valid values are `default`, `dedicated` or `host`.
        /// </summary>
        [Input("tenancy")]
        public Input<string>? Tenancy { get; set; }

        /// <summary>
        /// Keyword to choose what virtualization mode created instances will use. Valid values are `paravirtual` or `hvm`.
        /// </summary>
        [Input("virtualizationType")]
        public Input<string>? VirtualizationType { get; set; }

        public InstanceState()
        {
        }
        public static new InstanceState Empty => new InstanceState();
    }
}
