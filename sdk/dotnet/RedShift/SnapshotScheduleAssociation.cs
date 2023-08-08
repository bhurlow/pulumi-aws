// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.RedShift
{
    /// <summary>
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
    ///     var defaultCluster = new Aws.RedShift.Cluster("defaultCluster", new()
    ///     {
    ///         ClusterIdentifier = "tf-redshift-cluster",
    ///         DatabaseName = "mydb",
    ///         MasterUsername = "foo",
    ///         MasterPassword = "Mustbe8characters",
    ///         NodeType = "dc1.large",
    ///         ClusterType = "single-node",
    ///     });
    /// 
    ///     var defaultSnapshotSchedule = new Aws.RedShift.SnapshotSchedule("defaultSnapshotSchedule", new()
    ///     {
    ///         Identifier = "tf-redshift-snapshot-schedule",
    ///         Definitions = new[]
    ///         {
    ///             "rate(12 hours)",
    ///         },
    ///     });
    /// 
    ///     var defaultSnapshotScheduleAssociation = new Aws.RedShift.SnapshotScheduleAssociation("defaultSnapshotScheduleAssociation", new()
    ///     {
    ///         ClusterIdentifier = defaultCluster.Id,
    ///         ScheduleIdentifier = defaultSnapshotSchedule.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// terraform import {
    /// 
    ///  to = aws_redshift_snapshot_schedule_association.default
    /// 
    ///  id = "tf-redshift-cluster/tf-redshift-snapshot-schedule" } Using `pulumi import`, import Redshift Snapshot Schedule Association using the `&lt;cluster-identifier&gt;/&lt;schedule-identifier&gt;`. For exampleconsole % pulumi import aws_redshift_snapshot_schedule_association.default tf-redshift-cluster/tf-redshift-snapshot-schedule
    /// </summary>
    [AwsResourceType("aws:redshift/snapshotScheduleAssociation:SnapshotScheduleAssociation")]
    public partial class SnapshotScheduleAssociation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The cluster identifier.
        /// </summary>
        [Output("clusterIdentifier")]
        public Output<string> ClusterIdentifier { get; private set; } = null!;

        /// <summary>
        /// The snapshot schedule identifier.
        /// </summary>
        [Output("scheduleIdentifier")]
        public Output<string> ScheduleIdentifier { get; private set; } = null!;


        /// <summary>
        /// Create a SnapshotScheduleAssociation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SnapshotScheduleAssociation(string name, SnapshotScheduleAssociationArgs args, CustomResourceOptions? options = null)
            : base("aws:redshift/snapshotScheduleAssociation:SnapshotScheduleAssociation", name, args ?? new SnapshotScheduleAssociationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SnapshotScheduleAssociation(string name, Input<string> id, SnapshotScheduleAssociationState? state = null, CustomResourceOptions? options = null)
            : base("aws:redshift/snapshotScheduleAssociation:SnapshotScheduleAssociation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SnapshotScheduleAssociation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SnapshotScheduleAssociation Get(string name, Input<string> id, SnapshotScheduleAssociationState? state = null, CustomResourceOptions? options = null)
        {
            return new SnapshotScheduleAssociation(name, id, state, options);
        }
    }

    public sealed class SnapshotScheduleAssociationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The cluster identifier.
        /// </summary>
        [Input("clusterIdentifier", required: true)]
        public Input<string> ClusterIdentifier { get; set; } = null!;

        /// <summary>
        /// The snapshot schedule identifier.
        /// </summary>
        [Input("scheduleIdentifier", required: true)]
        public Input<string> ScheduleIdentifier { get; set; } = null!;

        public SnapshotScheduleAssociationArgs()
        {
        }
        public static new SnapshotScheduleAssociationArgs Empty => new SnapshotScheduleAssociationArgs();
    }

    public sealed class SnapshotScheduleAssociationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The cluster identifier.
        /// </summary>
        [Input("clusterIdentifier")]
        public Input<string>? ClusterIdentifier { get; set; }

        /// <summary>
        /// The snapshot schedule identifier.
        /// </summary>
        [Input("scheduleIdentifier")]
        public Input<string>? ScheduleIdentifier { get; set; }

        public SnapshotScheduleAssociationState()
        {
        }
        public static new SnapshotScheduleAssociationState Empty => new SnapshotScheduleAssociationState();
    }
}
