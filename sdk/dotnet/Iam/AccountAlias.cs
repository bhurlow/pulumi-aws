// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Iam
{
    /// <summary>
    /// &gt; **Note:** There is only a single account alias per AWS account.
    /// 
    /// Manages the account alias for the AWS Account.
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
    ///     var @alias = new Aws.Iam.AccountAlias("alias", new()
    ///     {
    ///         Alias = "my-account-alias",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// terraform import {
    /// 
    ///  to = aws_iam_account_alias.alias
    /// 
    ///  id = "my-account-alias" } Using `pulumi import`, import the current Account Alias using the `account_alias`. For exampleconsole % pulumi import aws_iam_account_alias.alias my-account-alias
    /// </summary>
    [AwsResourceType("aws:iam/accountAlias:AccountAlias")]
    public partial class AccountAlias : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The account alias
        /// </summary>
        [Output("accountAlias")]
        public Output<string> Alias { get; private set; } = null!;


        /// <summary>
        /// Create a AccountAlias resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AccountAlias(string name, AccountAliasArgs args, CustomResourceOptions? options = null)
            : base("aws:iam/accountAlias:AccountAlias", name, args ?? new AccountAliasArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AccountAlias(string name, Input<string> id, AccountAliasState? state = null, CustomResourceOptions? options = null)
            : base("aws:iam/accountAlias:AccountAlias", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AccountAlias resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AccountAlias Get(string name, Input<string> id, AccountAliasState? state = null, CustomResourceOptions? options = null)
        {
            return new AccountAlias(name, id, state, options);
        }
    }

    public sealed class AccountAliasArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The account alias
        /// </summary>
        [Input("accountAlias", required: true)]
        public Input<string> Alias { get; set; } = null!;

        public AccountAliasArgs()
        {
        }
        public static new AccountAliasArgs Empty => new AccountAliasArgs();
    }

    public sealed class AccountAliasState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The account alias
        /// </summary>
        [Input("accountAlias")]
        public Input<string>? Alias { get; set; }

        public AccountAliasState()
        {
        }
        public static new AccountAliasState Empty => new AccountAliasState();
    }
}
