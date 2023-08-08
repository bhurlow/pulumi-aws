// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Efs
{
    /// <summary>
    /// Creates a replica of an existing EFS file system in the same or another region. Creating this resource causes the source EFS file system to be replicated to a new read-only destination EFS file system. Deleting this resource will cause the replication from source to destination to stop and the destination file system will no longer be read only.
    /// 
    /// &gt; **NOTE:** Deleting this resource does **not** delete the destination file system that was created.
    /// 
    /// ## Example Usage
    /// 
    /// Will create a replica using regional storage in us-west-2 that will be encrypted by the default EFS KMS key `/aws/elasticfilesystem`.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleFileSystem = new Aws.Efs.FileSystem("exampleFileSystem");
    /// 
    ///     var exampleReplicationConfiguration = new Aws.Efs.ReplicationConfiguration("exampleReplicationConfiguration", new()
    ///     {
    ///         SourceFileSystemId = exampleFileSystem.Id,
    ///         Destination = new Aws.Efs.Inputs.ReplicationConfigurationDestinationArgs
    ///         {
    ///             Region = "us-west-2",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// Replica will be created as One Zone storage in the us-west-2b Availability Zone and encrypted with the specified KMS key.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleFileSystem = new Aws.Efs.FileSystem("exampleFileSystem");
    /// 
    ///     var exampleReplicationConfiguration = new Aws.Efs.ReplicationConfiguration("exampleReplicationConfiguration", new()
    ///     {
    ///         SourceFileSystemId = exampleFileSystem.Id,
    ///         Destination = new Aws.Efs.Inputs.ReplicationConfigurationDestinationArgs
    ///         {
    ///             AvailabilityZoneName = "us-west-2b",
    ///             KmsKeyId = "1234abcd-12ab-34cd-56ef-1234567890ab",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// terraform import {
    /// 
    ///  to = aws_efs_replication_configuration.example
    /// 
    ///  id = "fs-id" } Using `pulumi import`, import EFS Replication Configurations using the file system ID of either the source or destination file system. When importing, the `availability_zone_name` and `kms_key_id` attributes must **not** be set in the configuration. The AWS API does not return these values when querying the replication configuration and their presence will therefore show as a diff in a subsequent plan. For exampleconsole % pulumi import aws_efs_replication_configuration.example fs-id
    /// </summary>
    [AwsResourceType("aws:efs/replicationConfiguration:ReplicationConfiguration")]
    public partial class ReplicationConfiguration : global::Pulumi.CustomResource
    {
        /// <summary>
        /// When the replication configuration was created.
        /// </summary>
        [Output("creationTime")]
        public Output<string> CreationTime { get; private set; } = null!;

        /// <summary>
        /// A destination configuration block (documented below).
        /// </summary>
        [Output("destination")]
        public Output<Outputs.ReplicationConfigurationDestination> Destination { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the original source Amazon EFS file system in the replication configuration.
        /// </summary>
        [Output("originalSourceFileSystemArn")]
        public Output<string> OriginalSourceFileSystemArn { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the current source file system in the replication configuration.
        /// </summary>
        [Output("sourceFileSystemArn")]
        public Output<string> SourceFileSystemArn { get; private set; } = null!;

        /// <summary>
        /// The ID of the file system that is to be replicated.
        /// </summary>
        [Output("sourceFileSystemId")]
        public Output<string> SourceFileSystemId { get; private set; } = null!;

        /// <summary>
        /// The AWS Region in which the source Amazon EFS file system is located.
        /// * `destination[0].file_system_id` - The fs ID of the replica.
        /// * `destination[0].status` - The status of the replication.
        /// </summary>
        [Output("sourceFileSystemRegion")]
        public Output<string> SourceFileSystemRegion { get; private set; } = null!;


        /// <summary>
        /// Create a ReplicationConfiguration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReplicationConfiguration(string name, ReplicationConfigurationArgs args, CustomResourceOptions? options = null)
            : base("aws:efs/replicationConfiguration:ReplicationConfiguration", name, args ?? new ReplicationConfigurationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ReplicationConfiguration(string name, Input<string> id, ReplicationConfigurationState? state = null, CustomResourceOptions? options = null)
            : base("aws:efs/replicationConfiguration:ReplicationConfiguration", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ReplicationConfiguration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ReplicationConfiguration Get(string name, Input<string> id, ReplicationConfigurationState? state = null, CustomResourceOptions? options = null)
        {
            return new ReplicationConfiguration(name, id, state, options);
        }
    }

    public sealed class ReplicationConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A destination configuration block (documented below).
        /// </summary>
        [Input("destination", required: true)]
        public Input<Inputs.ReplicationConfigurationDestinationArgs> Destination { get; set; } = null!;

        /// <summary>
        /// The ID of the file system that is to be replicated.
        /// </summary>
        [Input("sourceFileSystemId", required: true)]
        public Input<string> SourceFileSystemId { get; set; } = null!;

        public ReplicationConfigurationArgs()
        {
        }
        public static new ReplicationConfigurationArgs Empty => new ReplicationConfigurationArgs();
    }

    public sealed class ReplicationConfigurationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// When the replication configuration was created.
        /// </summary>
        [Input("creationTime")]
        public Input<string>? CreationTime { get; set; }

        /// <summary>
        /// A destination configuration block (documented below).
        /// </summary>
        [Input("destination")]
        public Input<Inputs.ReplicationConfigurationDestinationGetArgs>? Destination { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the original source Amazon EFS file system in the replication configuration.
        /// </summary>
        [Input("originalSourceFileSystemArn")]
        public Input<string>? OriginalSourceFileSystemArn { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the current source file system in the replication configuration.
        /// </summary>
        [Input("sourceFileSystemArn")]
        public Input<string>? SourceFileSystemArn { get; set; }

        /// <summary>
        /// The ID of the file system that is to be replicated.
        /// </summary>
        [Input("sourceFileSystemId")]
        public Input<string>? SourceFileSystemId { get; set; }

        /// <summary>
        /// The AWS Region in which the source Amazon EFS file system is located.
        /// * `destination[0].file_system_id` - The fs ID of the replica.
        /// * `destination[0].status` - The status of the replication.
        /// </summary>
        [Input("sourceFileSystemRegion")]
        public Input<string>? SourceFileSystemRegion { get; set; }

        public ReplicationConfigurationState()
        {
        }
        public static new ReplicationConfigurationState Empty => new ReplicationConfigurationState();
    }
}
