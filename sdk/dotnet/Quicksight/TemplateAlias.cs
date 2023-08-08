// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Quicksight
{
    /// <summary>
    /// Resource for managing an AWS QuickSight Template Alias.
    /// 
    /// ## Example Usage
    /// ### Basic Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.Quicksight.TemplateAlias("example", new()
    ///     {
    ///         AliasName = "example-alias",
    ///         TemplateId = aws_quicksight_template.Test.Template_id,
    ///         TemplateVersionNumber = aws_quicksight_template.Test.Version_number,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// terraform import {
    /// 
    ///  to = aws_quicksight_template_alias.example
    /// 
    ///  id = "123456789012,example-id,example-alias" } Using `pulumi import`, import QuickSight Template Alias using the AWS account ID, template ID, and alias name separated by a comma (`,`). For exampleconsole % pulumi import aws_quicksight_template_alias.example 123456789012,example-id,example-alias
    /// </summary>
    [AwsResourceType("aws:quicksight/templateAlias:TemplateAlias")]
    public partial class TemplateAlias : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Display name of the template alias.
        /// </summary>
        [Output("aliasName")]
        public Output<string> AliasName { get; private set; } = null!;

        /// <summary>
        /// Amazon Resource Name (ARN) of the template alias.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// AWS account ID.
        /// </summary>
        [Output("awsAccountId")]
        public Output<string> AwsAccountId { get; private set; } = null!;

        /// <summary>
        /// ID of the template.
        /// </summary>
        [Output("templateId")]
        public Output<string> TemplateId { get; private set; } = null!;

        /// <summary>
        /// Version number of the template.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Output("templateVersionNumber")]
        public Output<int> TemplateVersionNumber { get; private set; } = null!;


        /// <summary>
        /// Create a TemplateAlias resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TemplateAlias(string name, TemplateAliasArgs args, CustomResourceOptions? options = null)
            : base("aws:quicksight/templateAlias:TemplateAlias", name, args ?? new TemplateAliasArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TemplateAlias(string name, Input<string> id, TemplateAliasState? state = null, CustomResourceOptions? options = null)
            : base("aws:quicksight/templateAlias:TemplateAlias", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TemplateAlias resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TemplateAlias Get(string name, Input<string> id, TemplateAliasState? state = null, CustomResourceOptions? options = null)
        {
            return new TemplateAlias(name, id, state, options);
        }
    }

    public sealed class TemplateAliasArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Display name of the template alias.
        /// </summary>
        [Input("aliasName", required: true)]
        public Input<string> AliasName { get; set; } = null!;

        /// <summary>
        /// AWS account ID.
        /// </summary>
        [Input("awsAccountId")]
        public Input<string>? AwsAccountId { get; set; }

        /// <summary>
        /// ID of the template.
        /// </summary>
        [Input("templateId", required: true)]
        public Input<string> TemplateId { get; set; } = null!;

        /// <summary>
        /// Version number of the template.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Input("templateVersionNumber", required: true)]
        public Input<int> TemplateVersionNumber { get; set; } = null!;

        public TemplateAliasArgs()
        {
        }
        public static new TemplateAliasArgs Empty => new TemplateAliasArgs();
    }

    public sealed class TemplateAliasState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Display name of the template alias.
        /// </summary>
        [Input("aliasName")]
        public Input<string>? AliasName { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN) of the template alias.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// AWS account ID.
        /// </summary>
        [Input("awsAccountId")]
        public Input<string>? AwsAccountId { get; set; }

        /// <summary>
        /// ID of the template.
        /// </summary>
        [Input("templateId")]
        public Input<string>? TemplateId { get; set; }

        /// <summary>
        /// Version number of the template.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Input("templateVersionNumber")]
        public Input<int>? TemplateVersionNumber { get; set; }

        public TemplateAliasState()
        {
        }
        public static new TemplateAliasState Empty => new TemplateAliasState();
    }
}
