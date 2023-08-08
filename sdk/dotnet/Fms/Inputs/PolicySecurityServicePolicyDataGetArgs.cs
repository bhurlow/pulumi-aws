// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Fms.Inputs
{

    public sealed class PolicySecurityServicePolicyDataGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Details about the service that are specific to the service type, in JSON format. For service type `SHIELD_ADVANCED`, this is an empty string. Examples depending on `type` can be found in the [AWS Firewall Manager SecurityServicePolicyData API Reference](https://docs.aws.amazon.com/fms/2018-01-01/APIReference/API_SecurityServicePolicyData.html).
        /// </summary>
        [Input("managedServiceData")]
        public Input<string>? ManagedServiceData { get; set; }

        /// <summary>
        /// Contains the Network Firewall firewall policy options to configure a centralized deployment model. Documented below.
        /// </summary>
        [Input("policyOption")]
        public Input<Inputs.PolicySecurityServicePolicyDataPolicyOptionGetArgs>? PolicyOption { get; set; }

        /// <summary>
        /// The service that the policy is using to protect the resources. For the current list of supported types, please refer to the [AWS Firewall Manager SecurityServicePolicyData API Type Reference](https://docs.aws.amazon.com/fms/2018-01-01/APIReference/API_SecurityServicePolicyData.html#fms-Type-SecurityServicePolicyData-Type).
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public PolicySecurityServicePolicyDataGetArgs()
        {
        }
        public static new PolicySecurityServicePolicyDataGetArgs Empty => new PolicySecurityServicePolicyDataGetArgs();
    }
}
