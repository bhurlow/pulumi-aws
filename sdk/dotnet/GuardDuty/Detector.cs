// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.GuardDuty
{
    /// <summary>
    /// Provides a resource to manage a GuardDuty detector.
    /// 
    /// &gt; **NOTE:** Deleting this resource is equivalent to "disabling" GuardDuty for an AWS region, which removes all existing findings. You can set the `enable` attribute to `false` to instead "suspend" monitoring and feedback reporting while keeping existing data. See the [Suspending or Disabling Amazon GuardDuty documentation](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_suspend-disable.html) for more information.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var myDetector = new Aws.GuardDuty.Detector("myDetector", new Aws.GuardDuty.DetectorArgs
    ///         {
    ///             Datasources = new Aws.GuardDuty.Inputs.DetectorDatasourcesArgs
    ///             {
    ///                 Kubernetes = new Aws.GuardDuty.Inputs.DetectorDatasourcesKubernetesArgs
    ///                 {
    ///                     AuditLogs = new Aws.GuardDuty.Inputs.DetectorDatasourcesKubernetesAuditLogsArgs
    ///                     {
    ///                         Enable = false,
    ///                     },
    ///                 },
    ///                 S3Logs = new Aws.GuardDuty.Inputs.DetectorDatasourcesS3LogsArgs
    ///                 {
    ///                     Enable = true,
    ///                 },
    ///             },
    ///             Enable = true,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// GuardDuty detectors can be imported using the detector ID, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:guardduty/detector:Detector MyDetector 00b00fd5aecc0ab60a708659477e9617
    /// ```
    /// </summary>
    [AwsResourceType("aws:guardduty/detector:Detector")]
    public partial class Detector : Pulumi.CustomResource
    {
        /// <summary>
        /// The AWS account ID of the GuardDuty detector
        /// </summary>
        [Output("accountId")]
        public Output<string> AccountId { get; private set; } = null!;

        /// <summary>
        /// Amazon Resource Name (ARN) of the GuardDuty detector
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Describes which data sources will be enabled for the detector. See Data Sources below for more details.
        /// </summary>
        [Output("datasources")]
        public Output<Outputs.DetectorDatasources> Datasources { get; private set; } = null!;

        /// <summary>
        /// If true, enables [S3 protection](https://docs.aws.amazon.com/guardduty/latest/ug/s3-protection.html).
        /// Defaults to `true`.
        /// </summary>
        [Output("enable")]
        public Output<bool?> Enable { get; private set; } = null!;

        /// <summary>
        /// Specifies the frequency of notifications sent for subsequent finding occurrences. If the detector is a GuardDuty member account, the value is determined by the GuardDuty primary account and cannot be modified, otherwise defaults to `SIX_HOURS`. For standalone and GuardDuty primary accounts, it must be configured in this provider to enable drift detection. Valid values for standalone and primary accounts: `FIFTEEN_MINUTES`, `ONE_HOUR`, `SIX_HOURS`. See [AWS Documentation](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_findings_cloudwatch.html#guardduty_findings_cloudwatch_notification_frequency) for more information.
        /// </summary>
        [Output("findingPublishingFrequency")]
        public Output<string> FindingPublishingFrequency { get; private set; } = null!;

        /// <summary>
        /// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider .
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;


        /// <summary>
        /// Create a Detector resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Detector(string name, DetectorArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:guardduty/detector:Detector", name, args ?? new DetectorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Detector(string name, Input<string> id, DetectorState? state = null, CustomResourceOptions? options = null)
            : base("aws:guardduty/detector:Detector", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Detector resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Detector Get(string name, Input<string> id, DetectorState? state = null, CustomResourceOptions? options = null)
        {
            return new Detector(name, id, state, options);
        }
    }

    public sealed class DetectorArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Describes which data sources will be enabled for the detector. See Data Sources below for more details.
        /// </summary>
        [Input("datasources")]
        public Input<Inputs.DetectorDatasourcesArgs>? Datasources { get; set; }

        /// <summary>
        /// If true, enables [S3 protection](https://docs.aws.amazon.com/guardduty/latest/ug/s3-protection.html).
        /// Defaults to `true`.
        /// </summary>
        [Input("enable")]
        public Input<bool>? Enable { get; set; }

        /// <summary>
        /// Specifies the frequency of notifications sent for subsequent finding occurrences. If the detector is a GuardDuty member account, the value is determined by the GuardDuty primary account and cannot be modified, otherwise defaults to `SIX_HOURS`. For standalone and GuardDuty primary accounts, it must be configured in this provider to enable drift detection. Valid values for standalone and primary accounts: `FIFTEEN_MINUTES`, `ONE_HOUR`, `SIX_HOURS`. See [AWS Documentation](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_findings_cloudwatch.html#guardduty_findings_cloudwatch_notification_frequency) for more information.
        /// </summary>
        [Input("findingPublishingFrequency")]
        public Input<string>? FindingPublishingFrequency { get; set; }

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

        public DetectorArgs()
        {
        }
    }

    public sealed class DetectorState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The AWS account ID of the GuardDuty detector
        /// </summary>
        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN) of the GuardDuty detector
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Describes which data sources will be enabled for the detector. See Data Sources below for more details.
        /// </summary>
        [Input("datasources")]
        public Input<Inputs.DetectorDatasourcesGetArgs>? Datasources { get; set; }

        /// <summary>
        /// If true, enables [S3 protection](https://docs.aws.amazon.com/guardduty/latest/ug/s3-protection.html).
        /// Defaults to `true`.
        /// </summary>
        [Input("enable")]
        public Input<bool>? Enable { get; set; }

        /// <summary>
        /// Specifies the frequency of notifications sent for subsequent finding occurrences. If the detector is a GuardDuty member account, the value is determined by the GuardDuty primary account and cannot be modified, otherwise defaults to `SIX_HOURS`. For standalone and GuardDuty primary accounts, it must be configured in this provider to enable drift detection. Valid values for standalone and primary accounts: `FIFTEEN_MINUTES`, `ONE_HOUR`, `SIX_HOURS`. See [AWS Documentation](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_findings_cloudwatch.html#guardduty_findings_cloudwatch_notification_frequency) for more information.
        /// </summary>
        [Input("findingPublishingFrequency")]
        public Input<string>? FindingPublishingFrequency { get; set; }

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
        /// A map of tags assigned to the resource, including those inherited from the provider .
        /// </summary>
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        public DetectorState()
        {
        }
    }
}
