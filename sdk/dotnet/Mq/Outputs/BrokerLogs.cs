// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Mq.Outputs
{

    [OutputType]
    public sealed class BrokerLogs
    {
        /// <summary>
        /// Enables audit logging. Auditing is only possible for `engine_type` of `ActiveMQ`. User management action made using JMX or the ActiveMQ Web Console is logged. Defaults to `false`.
        /// </summary>
        public readonly bool? Audit;
        /// <summary>
        /// Enables general logging via CloudWatch. Defaults to `false`.
        /// </summary>
        public readonly bool? General;

        [OutputConstructor]
        private BrokerLogs(
            bool? audit,

            bool? general)
        {
            Audit = audit;
            General = general;
        }
    }
}
