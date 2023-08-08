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
    /// Provides a resource to manage an [AWS Macie Account](https://docs.aws.amazon.com/macie/latest/APIReference/macie.html).
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
    ///     var test = new Aws.Macie2.Account("test", new()
    ///     {
    ///         FindingPublishingFrequency = "FIFTEEN_MINUTES",
    ///         Status = "ENABLED",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// terraform import {
    /// 
    ///  to = aws_macie2_account.example
    /// 
    ///  id = "abcd1" } Using `pulumi import`, import `aws_macie2_account` using the id. For exampleconsole % pulumi import aws_macie2_account.example abcd1
    /// </summary>
    [AwsResourceType("aws:macie2/account:Account")]
    public partial class Account : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The date and time, in UTC and extended RFC 3339 format, when the Amazon Macie account was created.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Specifies how often to publish updates to policy findings for the account. This includes publishing updates to AWS Security Hub and Amazon EventBridge (formerly called Amazon CloudWatch Events). Valid values are `FIFTEEN_MINUTES`, `ONE_HOUR` or `SIX_HOURS`.
        /// </summary>
        [Output("findingPublishingFrequency")]
        public Output<string> FindingPublishingFrequency { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the service-linked role that allows Macie to monitor and analyze data in AWS resources for the account.
        /// </summary>
        [Output("serviceRole")]
        public Output<string> ServiceRole { get; private set; } = null!;

        /// <summary>
        /// Specifies the status for the account. To enable Amazon Macie and start all Macie activities for the account, set this value to `ENABLED`. Valid values are `ENABLED` or `PAUSED`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The date and time, in UTC and extended RFC 3339 format, of the most recent change to the status of the Macie account.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a Account resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Account(string name, AccountArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:macie2/account:Account", name, args ?? new AccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Account(string name, Input<string> id, AccountState? state = null, CustomResourceOptions? options = null)
            : base("aws:macie2/account:Account", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Account resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Account Get(string name, Input<string> id, AccountState? state = null, CustomResourceOptions? options = null)
        {
            return new Account(name, id, state, options);
        }
    }

    public sealed class AccountArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies how often to publish updates to policy findings for the account. This includes publishing updates to AWS Security Hub and Amazon EventBridge (formerly called Amazon CloudWatch Events). Valid values are `FIFTEEN_MINUTES`, `ONE_HOUR` or `SIX_HOURS`.
        /// </summary>
        [Input("findingPublishingFrequency")]
        public Input<string>? FindingPublishingFrequency { get; set; }

        /// <summary>
        /// Specifies the status for the account. To enable Amazon Macie and start all Macie activities for the account, set this value to `ENABLED`. Valid values are `ENABLED` or `PAUSED`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public AccountArgs()
        {
        }
        public static new AccountArgs Empty => new AccountArgs();
    }

    public sealed class AccountState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The date and time, in UTC and extended RFC 3339 format, when the Amazon Macie account was created.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Specifies how often to publish updates to policy findings for the account. This includes publishing updates to AWS Security Hub and Amazon EventBridge (formerly called Amazon CloudWatch Events). Valid values are `FIFTEEN_MINUTES`, `ONE_HOUR` or `SIX_HOURS`.
        /// </summary>
        [Input("findingPublishingFrequency")]
        public Input<string>? FindingPublishingFrequency { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the service-linked role that allows Macie to monitor and analyze data in AWS resources for the account.
        /// </summary>
        [Input("serviceRole")]
        public Input<string>? ServiceRole { get; set; }

        /// <summary>
        /// Specifies the status for the account. To enable Amazon Macie and start all Macie activities for the account, set this value to `ENABLED`. Valid values are `ENABLED` or `PAUSED`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The date and time, in UTC and extended RFC 3339 format, of the most recent change to the status of the Macie account.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public AccountState()
        {
        }
        public static new AccountState Empty => new AccountState();
    }
}
