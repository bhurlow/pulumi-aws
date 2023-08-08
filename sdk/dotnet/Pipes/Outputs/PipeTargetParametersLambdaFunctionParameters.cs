// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Pipes.Outputs
{

    [OutputType]
    public sealed class PipeTargetParametersLambdaFunctionParameters
    {
        /// <summary>
        /// Specify whether to invoke the function synchronously or asynchronously. Valid Values: REQUEST_RESPONSE, FIRE_AND_FORGET.
        /// </summary>
        public readonly string InvocationType;

        [OutputConstructor]
        private PipeTargetParametersLambdaFunctionParameters(string invocationType)
        {
            InvocationType = invocationType;
        }
    }
}
