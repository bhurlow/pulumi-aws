// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AutoScaling.Inputs
{

    public sealed class GroupMixedInstancesPolicyLaunchTemplateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Override the instance launch template specification in the Launch Template.
        /// </summary>
        [Input("launchTemplateSpecification", required: true)]
        public Input<Inputs.GroupMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecificationArgs> LaunchTemplateSpecification { get; set; } = null!;

        [Input("overrides")]
        private InputList<Inputs.GroupMixedInstancesPolicyLaunchTemplateOverrideArgs>? _overrides;

        /// <summary>
        /// List of nested arguments provides the ability to specify multiple instance types. This will override the same parameter in the launch template. For on-demand instances, Auto Scaling considers the order of preference of instance types to launch based on the order specified in the overrides list. Defined below.
        /// </summary>
        public InputList<Inputs.GroupMixedInstancesPolicyLaunchTemplateOverrideArgs> Overrides
        {
            get => _overrides ?? (_overrides = new InputList<Inputs.GroupMixedInstancesPolicyLaunchTemplateOverrideArgs>());
            set => _overrides = value;
        }

        public GroupMixedInstancesPolicyLaunchTemplateArgs()
        {
        }
    }
}
