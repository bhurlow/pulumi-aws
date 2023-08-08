// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.LicenseManager
{
    /// <summary>
    /// Provides a License Manager association.
    /// 
    /// &gt; **Note:** License configurations can also be associated with launch templates by specifying the `license_specifications` block for an `aws.ec2.LaunchTemplate`.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleAmi = Aws.Ec2.GetAmi.Invoke(new()
    ///     {
    ///         MostRecent = true,
    ///         Owners = new[]
    ///         {
    ///             "amazon",
    ///         },
    ///         Filters = new[]
    ///         {
    ///             new Aws.Ec2.Inputs.GetAmiFilterInputArgs
    ///             {
    ///                 Name = "name",
    ///                 Values = new[]
    ///                 {
    ///                     "amzn-ami-vpc-nat*",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var exampleInstance = new Aws.Ec2.Instance("exampleInstance", new()
    ///     {
    ///         Ami = exampleAmi.Apply(getAmiResult =&gt; getAmiResult.Id),
    ///         InstanceType = "t2.micro",
    ///     });
    /// 
    ///     var exampleLicenseConfiguration = new Aws.LicenseManager.LicenseConfiguration("exampleLicenseConfiguration", new()
    ///     {
    ///         LicenseCountingType = "Instance",
    ///     });
    /// 
    ///     var exampleAssociation = new Aws.LicenseManager.Association("exampleAssociation", new()
    ///     {
    ///         LicenseConfigurationArn = exampleLicenseConfiguration.Arn,
    ///         ResourceArn = exampleInstance.Arn,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// terraform import {
    /// 
    ///  to = aws_licensemanager_association.example
    /// 
    ///  id = "arn:aws:ec2:eu-west-1:123456789012:image/ami-123456789abcdef01,arn:aws:license-manager:eu-west-1:123456789012:license-configuration:lic-0123456789abcdef0123456789abcdef" } Using `pulumi import`, import license configurations using `resource_arn,license_configuration_arn`. For exampleconsole % pulumi import aws_licensemanager_association.example arn:aws:ec2:eu-west-1:123456789012:image/ami-123456789abcdef01,arn:aws:license-manager:eu-west-1:123456789012:license-configuration:lic-0123456789abcdef0123456789abcdef
    /// </summary>
    [AwsResourceType("aws:licensemanager/association:Association")]
    public partial class Association : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ARN of the license configuration.
        /// </summary>
        [Output("licenseConfigurationArn")]
        public Output<string> LicenseConfigurationArn { get; private set; } = null!;

        /// <summary>
        /// ARN of the resource associated with the license configuration.
        /// </summary>
        [Output("resourceArn")]
        public Output<string> ResourceArn { get; private set; } = null!;


        /// <summary>
        /// Create a Association resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Association(string name, AssociationArgs args, CustomResourceOptions? options = null)
            : base("aws:licensemanager/association:Association", name, args ?? new AssociationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Association(string name, Input<string> id, AssociationState? state = null, CustomResourceOptions? options = null)
            : base("aws:licensemanager/association:Association", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Association resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Association Get(string name, Input<string> id, AssociationState? state = null, CustomResourceOptions? options = null)
        {
            return new Association(name, id, state, options);
        }
    }

    public sealed class AssociationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ARN of the license configuration.
        /// </summary>
        [Input("licenseConfigurationArn", required: true)]
        public Input<string> LicenseConfigurationArn { get; set; } = null!;

        /// <summary>
        /// ARN of the resource associated with the license configuration.
        /// </summary>
        [Input("resourceArn", required: true)]
        public Input<string> ResourceArn { get; set; } = null!;

        public AssociationArgs()
        {
        }
        public static new AssociationArgs Empty => new AssociationArgs();
    }

    public sealed class AssociationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ARN of the license configuration.
        /// </summary>
        [Input("licenseConfigurationArn")]
        public Input<string>? LicenseConfigurationArn { get; set; }

        /// <summary>
        /// ARN of the resource associated with the license configuration.
        /// </summary>
        [Input("resourceArn")]
        public Input<string>? ResourceArn { get; set; }

        public AssociationState()
        {
        }
        public static new AssociationState Empty => new AssociationState();
    }
}
