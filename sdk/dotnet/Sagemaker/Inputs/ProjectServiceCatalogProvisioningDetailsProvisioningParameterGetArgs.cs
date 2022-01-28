// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Sagemaker.Inputs
{

    public sealed class ProjectServiceCatalogProvisioningDetailsProvisioningParameterGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The key that identifies a provisioning parameter.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// The value of the provisioning parameter.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public ProjectServiceCatalogProvisioningDetailsProvisioningParameterGetArgs()
        {
        }
    }
}
