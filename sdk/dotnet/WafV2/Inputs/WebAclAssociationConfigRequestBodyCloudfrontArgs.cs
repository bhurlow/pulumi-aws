// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Inputs
{

    public sealed class WebAclAssociationConfigRequestBodyCloudfrontArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the maximum size of the web request body component that an associated CloudFront distribution should send to AWS WAF for inspection. This applies to statements in the web ACL that inspect the body or JSON body. Valid values are `KB_16`, `KB_32`, `KB_48` and `KB_64`.
        /// </summary>
        [Input("defaultSizeInspectionLimit", required: true)]
        public Input<string> DefaultSizeInspectionLimit { get; set; } = null!;

        public WebAclAssociationConfigRequestBodyCloudfrontArgs()
        {
        }
        public static new WebAclAssociationConfigRequestBodyCloudfrontArgs Empty => new WebAclAssociationConfigRequestBodyCloudfrontArgs();
    }
}
