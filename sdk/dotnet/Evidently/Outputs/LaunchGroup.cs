// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Evidently.Outputs
{

    [OutputType]
    public sealed class LaunchGroup
    {
        /// <summary>
        /// Specifies the description of the launch group.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Specifies the name of the feature that the launch is using.
        /// </summary>
        public readonly string Feature;
        /// <summary>
        /// Specifies the name of the lahnch group.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Specifies the feature variation to use for this launch group.
        /// </summary>
        public readonly string Variation;

        [OutputConstructor]
        private LaunchGroup(
            string? description,

            string feature,

            string name,

            string variation)
        {
            Description = description;
            Feature = feature;
            Name = name;
            Variation = variation;
        }
    }
}
