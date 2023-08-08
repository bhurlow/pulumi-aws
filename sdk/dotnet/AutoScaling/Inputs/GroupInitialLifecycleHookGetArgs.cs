// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AutoScaling.Inputs
{

    public sealed class GroupInitialLifecycleHookGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("defaultResult")]
        public Input<string>? DefaultResult { get; set; }

        [Input("heartbeatTimeout")]
        public Input<int>? HeartbeatTimeout { get; set; }

        [Input("lifecycleTransition", required: true)]
        public Input<string> LifecycleTransition { get; set; } = null!;

        /// <summary>
        /// Name of the Auto Scaling Group. By default generated by Pulumi. Conflicts with `name_prefix`.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("notificationMetadata")]
        public Input<string>? NotificationMetadata { get; set; }

        [Input("notificationTargetArn")]
        public Input<string>? NotificationTargetArn { get; set; }

        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        public GroupInitialLifecycleHookGetArgs()
        {
        }
        public static new GroupInitialLifecycleHookGetArgs Empty => new GroupInitialLifecycleHookGetArgs();
    }
}
