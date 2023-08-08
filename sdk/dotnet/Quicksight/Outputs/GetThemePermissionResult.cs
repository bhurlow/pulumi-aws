// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Quicksight.Outputs
{

    [OutputType]
    public sealed class GetThemePermissionResult
    {
        /// <summary>
        /// List of IAM actions to grant or revoke permissions on.
        /// </summary>
        public readonly ImmutableArray<string> Actions;
        /// <summary>
        /// ARN of the principal. See the [ResourcePermission documentation](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_ResourcePermission.html) for the applicable ARN values.
        /// </summary>
        public readonly string Principal;

        [OutputConstructor]
        private GetThemePermissionResult(
            ImmutableArray<string> actions,

            string principal)
        {
            Actions = actions;
            Principal = principal;
        }
    }
}
