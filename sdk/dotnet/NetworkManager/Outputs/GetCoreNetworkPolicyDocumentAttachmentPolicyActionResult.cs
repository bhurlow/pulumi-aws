// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.NetworkManager.Outputs
{

    [OutputType]
    public sealed class GetCoreNetworkPolicyDocumentAttachmentPolicyActionResult
    {
        /// <summary>
        /// Defines how a segment is mapped. Values can be `constant` or `tag`. `constant` statically defines the segment to associate the attachment to. `tag` uses the value of a tag to dynamically try to map to a segment.reference_policies_elements_condition_operators.html) to evaluate.
        /// </summary>
        public readonly string AssociationMethod;
        /// <summary>
        /// Determines if this mapping should override the segment value for `require_attachment_acceptance`. You can only set this to `true`, indicating that this setting applies only to segments that have `require_attachment_acceptance` set to `false`. If the segment already has the default `require_attachment_acceptance`, you can set this to inherit segment’s acceptance value.
        /// </summary>
        public readonly bool? RequireAcceptance;
        /// <summary>
        /// The name of the segment.
        /// </summary>
        public readonly string? Segment;
        /// <summary>
        /// Maps the attachment to the value of a known key. This is used with the `association_method` is `tag`. For example a `tag` of `stage = “test”`, will map to a segment named `test`. The value must exactly match the name of a segment. This allows you to have many segments, but use only a single rule without having to define multiple nearly identical conditions. This prevents creating many similar conditions that all use the same keys to map to segments.
        /// </summary>
        public readonly string? TagValueOfKey;

        [OutputConstructor]
        private GetCoreNetworkPolicyDocumentAttachmentPolicyActionResult(
            string associationMethod,

            bool? requireAcceptance,

            string? segment,

            string? tagValueOfKey)
        {
            AssociationMethod = associationMethod;
            RequireAcceptance = requireAcceptance;
            Segment = segment;
            TagValueOfKey = tagValueOfKey;
        }
    }
}
