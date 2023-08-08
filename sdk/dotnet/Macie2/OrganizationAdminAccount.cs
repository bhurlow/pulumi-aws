// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Macie2
{
    /// <summary>
    /// Provides a resource to manage an [Amazon Macie Organization Admin Account](https://docs.aws.amazon.com/macie/latest/APIReference/admin.html).
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
    ///     var exampleAccount = new Aws.Macie2.Account("exampleAccount");
    /// 
    ///     var exampleOrganizationAdminAccount = new Aws.Macie2.OrganizationAdminAccount("exampleOrganizationAdminAccount", new()
    ///     {
    ///         AdminAccountId = "ID OF THE ADMIN ACCOUNT",
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             exampleAccount,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// terraform import {
    /// 
    ///  to = aws_macie2_organization_admin_account.example
    /// 
    ///  id = "abcd1" } Using `pulumi import`, import `aws_macie2_organization_admin_account` using the id. For exampleconsole % pulumi import aws_macie2_organization_admin_account.example abcd1
    /// </summary>
    [AwsResourceType("aws:macie2/organizationAdminAccount:OrganizationAdminAccount")]
    public partial class OrganizationAdminAccount : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The AWS account ID for the account to designate as the delegated Amazon Macie administrator account for the organization.
        /// </summary>
        [Output("adminAccountId")]
        public Output<string> AdminAccountId { get; private set; } = null!;


        /// <summary>
        /// Create a OrganizationAdminAccount resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OrganizationAdminAccount(string name, OrganizationAdminAccountArgs args, CustomResourceOptions? options = null)
            : base("aws:macie2/organizationAdminAccount:OrganizationAdminAccount", name, args ?? new OrganizationAdminAccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OrganizationAdminAccount(string name, Input<string> id, OrganizationAdminAccountState? state = null, CustomResourceOptions? options = null)
            : base("aws:macie2/organizationAdminAccount:OrganizationAdminAccount", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing OrganizationAdminAccount resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OrganizationAdminAccount Get(string name, Input<string> id, OrganizationAdminAccountState? state = null, CustomResourceOptions? options = null)
        {
            return new OrganizationAdminAccount(name, id, state, options);
        }
    }

    public sealed class OrganizationAdminAccountArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The AWS account ID for the account to designate as the delegated Amazon Macie administrator account for the organization.
        /// </summary>
        [Input("adminAccountId", required: true)]
        public Input<string> AdminAccountId { get; set; } = null!;

        public OrganizationAdminAccountArgs()
        {
        }
        public static new OrganizationAdminAccountArgs Empty => new OrganizationAdminAccountArgs();
    }

    public sealed class OrganizationAdminAccountState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The AWS account ID for the account to designate as the delegated Amazon Macie administrator account for the organization.
        /// </summary>
        [Input("adminAccountId")]
        public Input<string>? AdminAccountId { get; set; }

        public OrganizationAdminAccountState()
        {
        }
        public static new OrganizationAdminAccountState Empty => new OrganizationAdminAccountState();
    }
}
