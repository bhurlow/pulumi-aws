// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Sagemaker.Inputs
{

    public sealed class DomainDefaultUserSettingsJupyterServerAppSettingsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The default instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance. see Default Resource Spec below.
        /// </summary>
        [Input("defaultResourceSpec")]
        public Input<Inputs.DomainDefaultUserSettingsJupyterServerAppSettingsDefaultResourceSpecGetArgs>? DefaultResourceSpec { get; set; }

        public DomainDefaultUserSettingsJupyterServerAppSettingsGetArgs()
        {
        }
    }
}
