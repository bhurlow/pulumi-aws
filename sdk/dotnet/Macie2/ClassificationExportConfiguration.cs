// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Macie2
{
    /// <summary>
    /// Provides a resource to manage an [Amazon Macie Classification Export Configuration](https://docs.aws.amazon.com/macie/latest/APIReference/classification-export-configuration.html).
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
    ///     var exampleAccount = new Aws.Macie2.Account("exampleAccount");
    /// 
    ///     var exampleClassificationExportConfiguration = new Aws.Macie2.ClassificationExportConfiguration("exampleClassificationExportConfiguration", new()
    ///     {
    ///         S3Destination = new Aws.Macie2.Inputs.ClassificationExportConfigurationS3DestinationArgs
    ///         {
    ///             BucketName = aws_s3_bucket.Example.Bucket,
    ///             KeyPrefix = "exampleprefix/",
    ///             KmsKeyArn = aws_kms_key.Example.Arn,
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             exampleAccount,
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
    ///  to = aws_macie2_classification_export_configuration.example
    /// 
    ///  id = "123456789012:us-west-2" } Using `pulumi import`, import `aws_macie2_classification_export_configuration` using the account ID and region. For exampleconsole % pulumi import aws_macie2_classification_export_configuration.example 123456789012:us-west-2
    /// </summary>
    [AwsResourceType("aws:macie2/classificationExportConfiguration:ClassificationExportConfiguration")]
    public partial class ClassificationExportConfiguration : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Configuration block for a S3 Destination. Defined below
        /// </summary>
        [Output("s3Destination")]
        public Output<Outputs.ClassificationExportConfigurationS3Destination?> S3Destination { get; private set; } = null!;


        /// <summary>
        /// Create a ClassificationExportConfiguration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ClassificationExportConfiguration(string name, ClassificationExportConfigurationArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:macie2/classificationExportConfiguration:ClassificationExportConfiguration", name, args ?? new ClassificationExportConfigurationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ClassificationExportConfiguration(string name, Input<string> id, ClassificationExportConfigurationState? state = null, CustomResourceOptions? options = null)
            : base("aws:macie2/classificationExportConfiguration:ClassificationExportConfiguration", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ClassificationExportConfiguration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ClassificationExportConfiguration Get(string name, Input<string> id, ClassificationExportConfigurationState? state = null, CustomResourceOptions? options = null)
        {
            return new ClassificationExportConfiguration(name, id, state, options);
        }
    }

    public sealed class ClassificationExportConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration block for a S3 Destination. Defined below
        /// </summary>
        [Input("s3Destination")]
        public Input<Inputs.ClassificationExportConfigurationS3DestinationArgs>? S3Destination { get; set; }

        public ClassificationExportConfigurationArgs()
        {
        }
        public static new ClassificationExportConfigurationArgs Empty => new ClassificationExportConfigurationArgs();
    }

    public sealed class ClassificationExportConfigurationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration block for a S3 Destination. Defined below
        /// </summary>
        [Input("s3Destination")]
        public Input<Inputs.ClassificationExportConfigurationS3DestinationGetArgs>? S3Destination { get; set; }

        public ClassificationExportConfigurationState()
        {
        }
        public static new ClassificationExportConfigurationState Empty => new ClassificationExportConfigurationState();
    }
}
