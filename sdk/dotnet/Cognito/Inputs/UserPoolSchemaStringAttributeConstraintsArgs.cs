// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Cognito.Inputs
{

    public sealed class UserPoolSchemaStringAttributeConstraintsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maximum length of an attribute value of the string type.
        /// </summary>
        [Input("maxLength")]
        public Input<string>? MaxLength { get; set; }

        /// <summary>
        /// Minimum length of an attribute value of the string type.
        /// </summary>
        [Input("minLength")]
        public Input<string>? MinLength { get; set; }

        public UserPoolSchemaStringAttributeConstraintsArgs()
        {
        }
    }
}
