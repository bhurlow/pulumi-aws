// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Kendra.Outputs
{

    [OutputType]
    public sealed class DataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfiguration
    {
        /// <summary>
        /// A block that specifies the condition used for when a Lambda function should be invoked. For example, you can specify a condition that if there are empty date-time values, then Amazon Kendra should invoke a function that inserts the current date-time. See invocation_condition.
        /// </summary>
        public readonly Outputs.DataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationCondition? InvocationCondition;
        /// <summary>
        /// The Amazon Resource Name (ARN) of a Lambda Function that can manipulate your document metadata fields or attributes and content.
        /// </summary>
        public readonly string LambdaArn;
        /// <summary>
        /// Stores the original, raw documents or the structured, parsed documents before and after altering them. For more information, see [Data contracts for Lambda functions](https://docs.aws.amazon.com/kendra/latest/dg/custom-document-enrichment.html#cde-data-contracts-lambda).
        /// </summary>
        public readonly string S3Bucket;

        [OutputConstructor]
        private DataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfiguration(
            Outputs.DataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationCondition? invocationCondition,

            string lambdaArn,

            string s3Bucket)
        {
            InvocationCondition = invocationCondition;
            LambdaArn = lambdaArn;
            S3Bucket = s3Bucket;
        }
    }
}
