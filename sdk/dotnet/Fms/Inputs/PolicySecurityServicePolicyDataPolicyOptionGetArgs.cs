// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Fms.Inputs
{

    public sealed class PolicySecurityServicePolicyDataPolicyOptionGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defines the deployment model to use for the firewall policy. Documented below.
        /// </summary>
        [Input("networkFirewallPolicy")]
        public Input<Inputs.PolicySecurityServicePolicyDataPolicyOptionNetworkFirewallPolicyGetArgs>? NetworkFirewallPolicy { get; set; }

        [Input("thirdPartyFirewallPolicy")]
        public Input<Inputs.PolicySecurityServicePolicyDataPolicyOptionThirdPartyFirewallPolicyGetArgs>? ThirdPartyFirewallPolicy { get; set; }

        public PolicySecurityServicePolicyDataPolicyOptionGetArgs()
        {
        }
        public static new PolicySecurityServicePolicyDataPolicyOptionGetArgs Empty => new PolicySecurityServicePolicyDataPolicyOptionGetArgs();
    }
}
