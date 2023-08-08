// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Fis.Inputs
{

    public sealed class ExperimentTemplateTargetGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("filters")]
        private InputList<Inputs.ExperimentTemplateTargetFilterGetArgs>? _filters;

        /// <summary>
        /// Filter(s) for the target. Filters can be used to select resources based on specific attributes returned by the respective describe action of the resource type. For more information, see [Targets for AWS FIS](https://docs.aws.amazon.com/fis/latest/userguide/targets.html#target-filters). See below.
        /// </summary>
        public InputList<Inputs.ExperimentTemplateTargetFilterGetArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.ExperimentTemplateTargetFilterGetArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// Friendly name given to the target.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("parameters")]
        private InputMap<string>? _parameters;

        /// <summary>
        /// The resource type parameters.
        /// 
        /// &gt; **NOTE:** The `target` configuration block requires either `resource_arns` or `resource_tag`.
        /// </summary>
        public InputMap<string> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<string>());
            set => _parameters = value;
        }

        [Input("resourceArns")]
        private InputList<string>? _resourceArns;

        /// <summary>
        /// Set of ARNs of the resources to target with an action. Conflicts with `resource_tag`.
        /// </summary>
        public InputList<string> ResourceArns
        {
            get => _resourceArns ?? (_resourceArns = new InputList<string>());
            set => _resourceArns = value;
        }

        [Input("resourceTags")]
        private InputList<Inputs.ExperimentTemplateTargetResourceTagGetArgs>? _resourceTags;

        /// <summary>
        /// Tag(s) the resources need to have to be considered a valid target for an action. Conflicts with `resource_arns`. See below.
        /// </summary>
        public InputList<Inputs.ExperimentTemplateTargetResourceTagGetArgs> ResourceTags
        {
            get => _resourceTags ?? (_resourceTags = new InputList<Inputs.ExperimentTemplateTargetResourceTagGetArgs>());
            set => _resourceTags = value;
        }

        /// <summary>
        /// AWS resource type. The resource type must be supported for the specified action. To find out what resource types are supported, see [Targets for AWS FIS](https://docs.aws.amazon.com/fis/latest/userguide/targets.html#resource-types).
        /// </summary>
        [Input("resourceType", required: true)]
        public Input<string> ResourceType { get; set; } = null!;

        /// <summary>
        /// Scopes the identified resources. Valid values are `ALL` (all identified resources), `COUNT(n)` (randomly select `n` of the identified resources), `PERCENT(n)` (randomly select `n` percent of the identified resources).
        /// </summary>
        [Input("selectionMode", required: true)]
        public Input<string> SelectionMode { get; set; } = null!;

        public ExperimentTemplateTargetGetArgs()
        {
        }
        public static new ExperimentTemplateTargetGetArgs Empty => new ExperimentTemplateTargetGetArgs();
    }
}
