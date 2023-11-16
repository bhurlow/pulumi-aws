// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Inspector2.Outputs
{

    [OutputType]
    public sealed class OrganizationConfigurationAutoEnable
    {
        /// <summary>
        /// Whether Amazon EC2 scans are automatically enabled for new members of your Amazon Inspector organization.
        /// </summary>
        public readonly bool Ec2;
        /// <summary>
        /// Whether Amazon ECR scans are automatically enabled for new members of your Amazon Inspector organization.
        /// </summary>
        public readonly bool Ecr;
        /// <summary>
        /// Whether Lambda Function scans are automatically enabled for new members of your Amazon Inspector organization.
        /// </summary>
        public readonly bool? Lambda;
        /// <summary>
        /// Whether AWS Lambda code scans are automatically enabled for new members of your Amazon Inspector organization. **Note:** Lambda code scanning requires Lambda standard scanning to be activated. Consequently, if you are setting this argument to `true`, you must also set the `lambda` argument to `true`. See [Scanning AWS Lambda functions with Amazon Inspector](https://docs.aws.amazon.com/inspector/latest/user/scanning-lambda.html#lambda-code-scans) for more information.
        /// </summary>
        public readonly bool? LambdaCode;

        [OutputConstructor]
        private OrganizationConfigurationAutoEnable(
            bool ec2,

            bool ecr,

            bool? lambda,

            bool? lambdaCode)
        {
            Ec2 = ec2;
            Ecr = ecr;
            Lambda = lambda;
            LambdaCode = lambdaCode;
        }
    }
}
