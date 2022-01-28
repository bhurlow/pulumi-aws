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
    public sealed class ProjectServiceCatalogProvisioningDetailsProvisioningParameter
    {
        /// <summary>
        /// The key that identifies a provisioning parameter.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// The value of the provisioning parameter.
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private ProjectServiceCatalogProvisioningDetailsProvisioningParameter(
            string key,

            string? value)
        {
            Key = key;
            Value = value;
        }
    }
}
