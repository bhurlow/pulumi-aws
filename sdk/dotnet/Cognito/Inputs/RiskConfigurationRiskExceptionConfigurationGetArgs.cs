// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Cognito.Inputs
{

    public sealed class RiskConfigurationRiskExceptionConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("blockedIpRangeLists")]
        private InputList<string>? _blockedIpRangeLists;

        /// <summary>
        /// Overrides the risk decision to always block the pre-authentication requests.
        /// The IP range is in CIDR notation, a compact representation of an IP address and its routing prefix.
        /// Can contain a maximum of 200 items.
        /// </summary>
        public InputList<string> BlockedIpRangeLists
        {
            get => _blockedIpRangeLists ?? (_blockedIpRangeLists = new InputList<string>());
            set => _blockedIpRangeLists = value;
        }

        [Input("skippedIpRangeLists")]
        private InputList<string>? _skippedIpRangeLists;

        /// <summary>
        /// Risk detection isn't performed on the IP addresses in this range list.
        /// The IP range is in CIDR notation.
        /// Can contain a maximum of 200 items.
        /// </summary>
        public InputList<string> SkippedIpRangeLists
        {
            get => _skippedIpRangeLists ?? (_skippedIpRangeLists = new InputList<string>());
            set => _skippedIpRangeLists = value;
        }

        public RiskConfigurationRiskExceptionConfigurationGetArgs()
        {
        }
        public static new RiskConfigurationRiskExceptionConfigurationGetArgs Empty => new RiskConfigurationRiskExceptionConfigurationGetArgs();
    }
}
