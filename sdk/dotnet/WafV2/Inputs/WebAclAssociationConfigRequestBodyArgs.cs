// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Inputs
{

    public sealed class WebAclAssociationConfigRequestBodyArgs : global::Pulumi.ResourceArgs
    {
        [Input("cloudfronts")]
        private InputList<Inputs.WebAclAssociationConfigRequestBodyCloudfrontArgs>? _cloudfronts;

        /// <summary>
        /// Customizes the request body that your protected CloudFront distributions forward to AWS WAF for inspection. See `cloudfront` below for details.
        /// </summary>
        public InputList<Inputs.WebAclAssociationConfigRequestBodyCloudfrontArgs> Cloudfronts
        {
            get => _cloudfronts ?? (_cloudfronts = new InputList<Inputs.WebAclAssociationConfigRequestBodyCloudfrontArgs>());
            set => _cloudfronts = value;
        }

        public WebAclAssociationConfigRequestBodyArgs()
        {
        }
        public static new WebAclAssociationConfigRequestBodyArgs Empty => new WebAclAssociationConfigRequestBodyArgs();
    }
}
