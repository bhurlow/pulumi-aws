// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Inspector
{
    /// <summary>
    /// Provides a Inspector assessment template
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/inspector_assessment_template.html.markdown.
    /// </summary>
    public partial class AssessmentTemplate : Pulumi.CustomResource
    {
        /// <summary>
        /// The template assessment ARN.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The duration of the inspector run.
        /// </summary>
        [Output("duration")]
        public Output<int> Duration { get; private set; } = null!;

        /// <summary>
        /// The name of the assessment template.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The rules to be used during the run.
        /// </summary>
        [Output("rulesPackageArns")]
        public Output<ImmutableArray<string>> RulesPackageArns { get; private set; } = null!;

        /// <summary>
        /// The assessment target ARN to attach the template to.
        /// </summary>
        [Output("targetArn")]
        public Output<string> TargetArn { get; private set; } = null!;


        /// <summary>
        /// Create a AssessmentTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AssessmentTemplate(string name, AssessmentTemplateArgs args, CustomResourceOptions? options = null)
            : base("aws:inspector/assessmentTemplate:AssessmentTemplate", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private AssessmentTemplate(string name, Input<string> id, AssessmentTemplateState? state = null, CustomResourceOptions? options = null)
            : base("aws:inspector/assessmentTemplate:AssessmentTemplate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AssessmentTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AssessmentTemplate Get(string name, Input<string> id, AssessmentTemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new AssessmentTemplate(name, id, state, options);
        }
    }

    public sealed class AssessmentTemplateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The duration of the inspector run.
        /// </summary>
        [Input("duration", required: true)]
        public Input<int> Duration { get; set; } = null!;

        /// <summary>
        /// The name of the assessment template.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("rulesPackageArns", required: true)]
        private InputList<string>? _rulesPackageArns;

        /// <summary>
        /// The rules to be used during the run.
        /// </summary>
        public InputList<string> RulesPackageArns
        {
            get => _rulesPackageArns ?? (_rulesPackageArns = new InputList<string>());
            set => _rulesPackageArns = value;
        }

        /// <summary>
        /// The assessment target ARN to attach the template to.
        /// </summary>
        [Input("targetArn", required: true)]
        public Input<string> TargetArn { get; set; } = null!;

        public AssessmentTemplateArgs()
        {
        }
    }

    public sealed class AssessmentTemplateState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The template assessment ARN.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The duration of the inspector run.
        /// </summary>
        [Input("duration")]
        public Input<int>? Duration { get; set; }

        /// <summary>
        /// The name of the assessment template.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("rulesPackageArns")]
        private InputList<string>? _rulesPackageArns;

        /// <summary>
        /// The rules to be used during the run.
        /// </summary>
        public InputList<string> RulesPackageArns
        {
            get => _rulesPackageArns ?? (_rulesPackageArns = new InputList<string>());
            set => _rulesPackageArns = value;
        }

        /// <summary>
        /// The assessment target ARN to attach the template to.
        /// </summary>
        [Input("targetArn")]
        public Input<string>? TargetArn { get; set; }

        public AssessmentTemplateState()
        {
        }
    }
}
