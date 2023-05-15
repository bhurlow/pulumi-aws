// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Rds
{
    /// <summary>
    /// Manages an RDS DB Reserved Instance.
    /// 
    /// &gt; **NOTE:** Once created, a reservation is valid for the `duration` of the provided `offering_id` and cannot be deleted. Performing a `destroy` will only remove the resource from state. For more information see [RDS Reserved Instances Documentation](https://aws.amazon.com/rds/reserved-instances/) and [PurchaseReservedDBInstancesOffering](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_PurchaseReservedDBInstancesOffering.html).
    /// 
    /// &gt; **NOTE:** Due to the expense of testing this resource, we provide it as best effort. If you find it useful, and have the ability to help test or notice issues, consider reaching out to us on GitHub.
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
    ///     var test = Aws.Rds.GetReservedInstanceOffering.Invoke(new()
    ///     {
    ///         DbInstanceClass = "db.t2.micro",
    ///         Duration = 31536000,
    ///         MultiAz = false,
    ///         OfferingType = "All Upfront",
    ///         ProductDescription = "mysql",
    ///     });
    /// 
    ///     var example = new Aws.Rds.ReservedInstance("example", new()
    ///     {
    ///         OfferingId = test.Apply(getReservedInstanceOfferingResult =&gt; getReservedInstanceOfferingResult.OfferingId),
    ///         ReservationId = "optionalCustomReservationID",
    ///         InstanceCount = 3,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// RDS DB Instance Reservations can be imported using the `instance_id`, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:rds/reservedInstance:ReservedInstance reservation_instance CustomReservationID
    /// ```
    /// </summary>
    [AwsResourceType("aws:rds/reservedInstance:ReservedInstance")]
    public partial class ReservedInstance : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ARN for the reserved DB instance.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Currency code for the reserved DB instance.
        /// </summary>
        [Output("currencyCode")]
        public Output<string> CurrencyCode { get; private set; } = null!;

        /// <summary>
        /// DB instance class for the reserved DB instance.
        /// </summary>
        [Output("dbInstanceClass")]
        public Output<string> DbInstanceClass { get; private set; } = null!;

        /// <summary>
        /// Duration of the reservation in seconds.
        /// </summary>
        [Output("duration")]
        public Output<int> Duration { get; private set; } = null!;

        /// <summary>
        /// Fixed price charged for this reserved DB instance.
        /// </summary>
        [Output("fixedPrice")]
        public Output<double> FixedPrice { get; private set; } = null!;

        /// <summary>
        /// Number of instances to reserve. Default value is `1`.
        /// </summary>
        [Output("instanceCount")]
        public Output<int?> InstanceCount { get; private set; } = null!;

        /// <summary>
        /// Unique identifier for the lease associated with the reserved DB instance. Amazon Web Services Support might request the lease ID for an issue related to a reserved DB instance.
        /// </summary>
        [Output("leaseId")]
        public Output<string> LeaseId { get; private set; } = null!;

        /// <summary>
        /// Whether the reservation applies to Multi-AZ deployments.
        /// </summary>
        [Output("multiAz")]
        public Output<bool> MultiAz { get; private set; } = null!;

        /// <summary>
        /// ID of the Reserved DB instance offering to purchase. To determine an `offering_id`, see the `aws.rds.getReservedInstanceOffering` data source.
        /// </summary>
        [Output("offeringId")]
        public Output<string> OfferingId { get; private set; } = null!;

        /// <summary>
        /// Offering type of this reserved DB instance.
        /// </summary>
        [Output("offeringType")]
        public Output<string> OfferingType { get; private set; } = null!;

        /// <summary>
        /// Description of the reserved DB instance.
        /// </summary>
        [Output("productDescription")]
        public Output<string> ProductDescription { get; private set; } = null!;

        /// <summary>
        /// Recurring price charged to run this reserved DB instance.
        /// </summary>
        [Output("recurringCharges")]
        public Output<ImmutableArray<Outputs.ReservedInstanceRecurringCharge>> RecurringCharges { get; private set; } = null!;

        /// <summary>
        /// Customer-specified identifier to track this reservation.
        /// </summary>
        [Output("reservationId")]
        public Output<string?> ReservationId { get; private set; } = null!;

        /// <summary>
        /// Time the reservation started.
        /// </summary>
        [Output("startTime")]
        public Output<string> StartTime { get; private set; } = null!;

        /// <summary>
        /// State of the reserved DB instance.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Map of tags to assign to the DB reservation. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// Hourly price charged for this reserved DB instance.
        /// </summary>
        [Output("usagePrice")]
        public Output<double> UsagePrice { get; private set; } = null!;


        /// <summary>
        /// Create a ReservedInstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReservedInstance(string name, ReservedInstanceArgs args, CustomResourceOptions? options = null)
            : base("aws:rds/reservedInstance:ReservedInstance", name, args ?? new ReservedInstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ReservedInstance(string name, Input<string> id, ReservedInstanceState? state = null, CustomResourceOptions? options = null)
            : base("aws:rds/reservedInstance:ReservedInstance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ReservedInstance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ReservedInstance Get(string name, Input<string> id, ReservedInstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new ReservedInstance(name, id, state, options);
        }
    }

    public sealed class ReservedInstanceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of instances to reserve. Default value is `1`.
        /// </summary>
        [Input("instanceCount")]
        public Input<int>? InstanceCount { get; set; }

        /// <summary>
        /// ID of the Reserved DB instance offering to purchase. To determine an `offering_id`, see the `aws.rds.getReservedInstanceOffering` data source.
        /// </summary>
        [Input("offeringId", required: true)]
        public Input<string> OfferingId { get; set; } = null!;

        /// <summary>
        /// Customer-specified identifier to track this reservation.
        /// </summary>
        [Input("reservationId")]
        public Input<string>? ReservationId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of tags to assign to the DB reservation. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ReservedInstanceArgs()
        {
        }
        public static new ReservedInstanceArgs Empty => new ReservedInstanceArgs();
    }

    public sealed class ReservedInstanceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ARN for the reserved DB instance.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Currency code for the reserved DB instance.
        /// </summary>
        [Input("currencyCode")]
        public Input<string>? CurrencyCode { get; set; }

        /// <summary>
        /// DB instance class for the reserved DB instance.
        /// </summary>
        [Input("dbInstanceClass")]
        public Input<string>? DbInstanceClass { get; set; }

        /// <summary>
        /// Duration of the reservation in seconds.
        /// </summary>
        [Input("duration")]
        public Input<int>? Duration { get; set; }

        /// <summary>
        /// Fixed price charged for this reserved DB instance.
        /// </summary>
        [Input("fixedPrice")]
        public Input<double>? FixedPrice { get; set; }

        /// <summary>
        /// Number of instances to reserve. Default value is `1`.
        /// </summary>
        [Input("instanceCount")]
        public Input<int>? InstanceCount { get; set; }

        /// <summary>
        /// Unique identifier for the lease associated with the reserved DB instance. Amazon Web Services Support might request the lease ID for an issue related to a reserved DB instance.
        /// </summary>
        [Input("leaseId")]
        public Input<string>? LeaseId { get; set; }

        /// <summary>
        /// Whether the reservation applies to Multi-AZ deployments.
        /// </summary>
        [Input("multiAz")]
        public Input<bool>? MultiAz { get; set; }

        /// <summary>
        /// ID of the Reserved DB instance offering to purchase. To determine an `offering_id`, see the `aws.rds.getReservedInstanceOffering` data source.
        /// </summary>
        [Input("offeringId")]
        public Input<string>? OfferingId { get; set; }

        /// <summary>
        /// Offering type of this reserved DB instance.
        /// </summary>
        [Input("offeringType")]
        public Input<string>? OfferingType { get; set; }

        /// <summary>
        /// Description of the reserved DB instance.
        /// </summary>
        [Input("productDescription")]
        public Input<string>? ProductDescription { get; set; }

        [Input("recurringCharges")]
        private InputList<Inputs.ReservedInstanceRecurringChargeGetArgs>? _recurringCharges;

        /// <summary>
        /// Recurring price charged to run this reserved DB instance.
        /// </summary>
        public InputList<Inputs.ReservedInstanceRecurringChargeGetArgs> RecurringCharges
        {
            get => _recurringCharges ?? (_recurringCharges = new InputList<Inputs.ReservedInstanceRecurringChargeGetArgs>());
            set => _recurringCharges = value;
        }

        /// <summary>
        /// Customer-specified identifier to track this reservation.
        /// </summary>
        [Input("reservationId")]
        public Input<string>? ReservationId { get; set; }

        /// <summary>
        /// Time the reservation started.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        /// <summary>
        /// State of the reserved DB instance.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of tags to assign to the DB reservation. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;

        /// <summary>
        /// Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        /// <summary>
        /// Hourly price charged for this reserved DB instance.
        /// </summary>
        [Input("usagePrice")]
        public Input<double>? UsagePrice { get; set; }

        public ReservedInstanceState()
        {
        }
        public static new ReservedInstanceState Empty => new ReservedInstanceState();
    }
}
