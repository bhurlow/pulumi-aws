// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Connect.Outputs
{

    [OutputType]
    public sealed class GetQuickConnectQuickConnectConfigPhoneConfigResult
    {
        /// <summary>
        /// Specifies the phone number in in E.164 format.
        /// </summary>
        public readonly string PhoneNumber;

        [OutputConstructor]
        private GetQuickConnectQuickConnectConfigPhoneConfigResult(string phoneNumber)
        {
            PhoneNumber = phoneNumber;
        }
    }
}
