// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppFlow.Outputs
{

    [OutputType]
    public sealed class FlowDestinationFlowConfigDestinationConnectorPropertiesCustomerProfiles
    {
        /// <summary>
        /// The unique name of the Amazon Connect Customer Profiles domain.
        /// </summary>
        public readonly string DomainName;
        /// <summary>
        /// The object specified in the Amazon Connect Customer Profiles flow destination.
        /// </summary>
        public readonly string? ObjectTypeName;

        [OutputConstructor]
        private FlowDestinationFlowConfigDestinationConnectorPropertiesCustomerProfiles(
            string domainName,

            string? objectTypeName)
        {
            DomainName = domainName;
            ObjectTypeName = objectTypeName;
        }
    }
}
