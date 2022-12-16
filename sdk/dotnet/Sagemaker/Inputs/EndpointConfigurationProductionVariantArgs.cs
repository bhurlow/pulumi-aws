// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Sagemaker.Inputs
{

    public sealed class EndpointConfigurationProductionVariantArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The size of the Elastic Inference (EI) instance to use for the production variant.
        /// </summary>
        [Input("acceleratorType")]
        public Input<string>? AcceleratorType { get; set; }

        /// <summary>
        /// The timeout value, in seconds, for your inference container to pass health check by SageMaker Hosting. For more information about health check, see [How Your Container Should Respond to Health Check (Ping) Requests](https://docs.aws.amazon.com/sagemaker/latest/dg/your-algorithms-inference-code.html#your-algorithms-inference-algo-ping-requests). Valid values between `60` and `3600`.
        /// </summary>
        [Input("containerStartupHealthCheckTimeoutInSeconds")]
        public Input<int>? ContainerStartupHealthCheckTimeoutInSeconds { get; set; }

        /// <summary>
        /// Specifies configuration for a core dump from the model container when the process crashes. Fields are documented below.
        /// </summary>
        [Input("coreDumpConfig")]
        public Input<Inputs.EndpointConfigurationProductionVariantCoreDumpConfigArgs>? CoreDumpConfig { get; set; }

        /// <summary>
        /// Initial number of instances used for auto-scaling.
        /// </summary>
        [Input("initialInstanceCount")]
        public Input<int>? InitialInstanceCount { get; set; }

        /// <summary>
        /// Determines initial traffic distribution among all of the models that you specify in the endpoint configuration. If unspecified, it defaults to `1.0`.
        /// </summary>
        [Input("initialVariantWeight")]
        public Input<double>? InitialVariantWeight { get; set; }

        /// <summary>
        /// The type of instance to start.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// The timeout value, in seconds, to download and extract the model that you want to host from Amazon S3 to the individual inference instance associated with this production variant. Valid values between `60` and `3600`.
        /// </summary>
        [Input("modelDataDownloadTimeoutInSeconds")]
        public Input<int>? ModelDataDownloadTimeoutInSeconds { get; set; }

        /// <summary>
        /// The name of the model to use.
        /// </summary>
        [Input("modelName", required: true)]
        public Input<string> ModelName { get; set; } = null!;

        /// <summary>
        /// Specifies configuration for how an endpoint performs asynchronous inference.
        /// </summary>
        [Input("serverlessConfig")]
        public Input<Inputs.EndpointConfigurationProductionVariantServerlessConfigArgs>? ServerlessConfig { get; set; }

        /// <summary>
        /// The name of the variant. If omitted, this provider will assign a random, unique name.
        /// </summary>
        [Input("variantName")]
        public Input<string>? VariantName { get; set; }

        /// <summary>
        /// The size, in GB, of the ML storage volume attached to individual inference instance associated with the production variant. Valid values between `1` and `512`.
        /// </summary>
        [Input("volumeSizeInGb")]
        public Input<int>? VolumeSizeInGb { get; set; }

        public EndpointConfigurationProductionVariantArgs()
        {
        }
        public static new EndpointConfigurationProductionVariantArgs Empty => new EndpointConfigurationProductionVariantArgs();
    }
}
