// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Sagemaker.Outputs
{

    [OutputType]
    public sealed class EndpointDeploymentConfigAutoRollbackConfiguration
    {
        /// <summary>
        /// List of CloudWatch alarms in your account that are configured to monitor metrics on an endpoint. If any alarms are tripped during a deployment, SageMaker rolls back the deployment. See Alarms.
        /// </summary>
        public readonly ImmutableArray<Outputs.EndpointDeploymentConfigAutoRollbackConfigurationAlarm> Alarms;

        [OutputConstructor]
        private EndpointDeploymentConfigAutoRollbackConfiguration(ImmutableArray<Outputs.EndpointDeploymentConfigAutoRollbackConfigurationAlarm> alarms)
        {
            Alarms = alarms;
        }
    }
}