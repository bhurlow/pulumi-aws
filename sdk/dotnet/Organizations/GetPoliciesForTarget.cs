// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Organizations
{
    public static class GetPoliciesForTarget
    {
        /// <summary>
        /// Data source for managing an AWS Organizations Policies For Target.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetPoliciesForTargetResult> InvokeAsync(GetPoliciesForTargetArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPoliciesForTargetResult>("aws:organizations/getPoliciesForTarget:getPoliciesForTarget", args ?? new GetPoliciesForTargetArgs(), options.WithDefaults());

        /// <summary>
        /// Data source for managing an AWS Organizations Policies For Target.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetPoliciesForTargetResult> Invoke(GetPoliciesForTargetInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPoliciesForTargetResult>("aws:organizations/getPoliciesForTarget:getPoliciesForTarget", args ?? new GetPoliciesForTargetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPoliciesForTargetArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Must supply one of the 4 different policy filters for a target (SERVICE_CONTROL_POLICY | TAG_POLICY | BACKUP_POLICY | AISERVICES_OPT_OUT_POLICY)
        /// </summary>
        [Input("filter", required: true)]
        public string Filter { get; set; } = null!;

        /// <summary>
        /// The root (string that begins with "r-" followed by 4-32 lowercase letters or digits), account (12 digit string), or Organizational Unit (string starting with "ou-" followed by 4-32 lowercase letters or digits. This string is followed by a second "-" dash and from 8-32 additional lowercase letters or digits.)
        /// </summary>
        [Input("targetId", required: true)]
        public string TargetId { get; set; } = null!;

        public GetPoliciesForTargetArgs()
        {
        }
        public static new GetPoliciesForTargetArgs Empty => new GetPoliciesForTargetArgs();
    }

    public sealed class GetPoliciesForTargetInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Must supply one of the 4 different policy filters for a target (SERVICE_CONTROL_POLICY | TAG_POLICY | BACKUP_POLICY | AISERVICES_OPT_OUT_POLICY)
        /// </summary>
        [Input("filter", required: true)]
        public Input<string> Filter { get; set; } = null!;

        /// <summary>
        /// The root (string that begins with "r-" followed by 4-32 lowercase letters or digits), account (12 digit string), or Organizational Unit (string starting with "ou-" followed by 4-32 lowercase letters or digits. This string is followed by a second "-" dash and from 8-32 additional lowercase letters or digits.)
        /// </summary>
        [Input("targetId", required: true)]
        public Input<string> TargetId { get; set; } = null!;

        public GetPoliciesForTargetInvokeArgs()
        {
        }
        public static new GetPoliciesForTargetInvokeArgs Empty => new GetPoliciesForTargetInvokeArgs();
    }


    [OutputType]
    public sealed class GetPoliciesForTargetResult
    {
        public readonly string Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// List of all the policy ids found.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string TargetId;

        [OutputConstructor]
        private GetPoliciesForTargetResult(
            string filter,

            string id,

            ImmutableArray<string> ids,

            string targetId)
        {
            Filter = filter;
            Id = id;
            Ids = ids;
            TargetId = targetId;
        }
    }
}
