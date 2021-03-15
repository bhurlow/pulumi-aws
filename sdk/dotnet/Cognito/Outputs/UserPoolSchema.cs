// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Cognito.Outputs
{

    [OutputType]
    public sealed class UserPoolSchema
    {
        /// <summary>
        /// Attribute data type. Must be one of `Boolean`, `Number`, `String`, `DateTime`.
        /// </summary>
        public readonly string AttributeDataType;
        /// <summary>
        /// Whether the attribute type is developer only.
        /// </summary>
        public readonly bool? DeveloperOnlyAttribute;
        /// <summary>
        /// Whether the attribute can be changed once it has been created.
        /// </summary>
        public readonly bool? Mutable;
        /// <summary>
        /// Name of the attribute.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Configuration block for the constraints for an attribute of the number type. Detailed below.
        /// </summary>
        public readonly Outputs.UserPoolSchemaNumberAttributeConstraints? NumberAttributeConstraints;
        /// <summary>
        /// Whether a user pool attribute is required. If the attribute is required and the user does not provide a value, registration or sign-in will fail.
        /// </summary>
        public readonly bool? Required;
        /// <summary>
        /// Constraints for an attribute of the string type. Detailed below.
        /// </summary>
        public readonly Outputs.UserPoolSchemaStringAttributeConstraints? StringAttributeConstraints;

        [OutputConstructor]
        private UserPoolSchema(
            string attributeDataType,

            bool? developerOnlyAttribute,

            bool? mutable,

            string name,

            Outputs.UserPoolSchemaNumberAttributeConstraints? numberAttributeConstraints,

            bool? required,

            Outputs.UserPoolSchemaStringAttributeConstraints? stringAttributeConstraints)
        {
            AttributeDataType = attributeDataType;
            DeveloperOnlyAttribute = developerOnlyAttribute;
            Mutable = mutable;
            Name = name;
            NumberAttributeConstraints = numberAttributeConstraints;
            Required = required;
            StringAttributeConstraints = stringAttributeConstraints;
        }
    }
}
