// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Cfg
{
    /// <summary>
    /// Manages a Config Organization Managed Rule. More information about these rules can be found in the [Enabling AWS Config Rules Across all Accounts in Your Organization](https://docs.aws.amazon.com/config/latest/developerguide/config-rule-multi-account-deployment.html) and [AWS Config Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_use-managed-rules.html) documentation. For working with Organization Custom Rules (those invoking a custom Lambda Function), see the `aws.cfg.OrganizationCustomRule` resource.
    /// 
    /// &gt; **NOTE:** This resource must be created in the Organization master account and rules will include the master account unless its ID is added to the `excluded_accounts` argument.
    /// 
    /// &gt; **NOTE:** Every Organization account except those configured in the `excluded_accounts` argument must have a Configuration Recorder with proper IAM permissions before the rule will successfully create or update. See also the `aws.cfg.Recorder` resource.
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
    ///     var exampleOrganization = new Aws.Organizations.Organization("exampleOrganization", new()
    ///     {
    ///         AwsServiceAccessPrincipals = new[]
    ///         {
    ///             "config-multiaccountsetup.amazonaws.com",
    ///         },
    ///         FeatureSet = "ALL",
    ///     });
    /// 
    ///     var exampleOrganizationManagedRule = new Aws.Cfg.OrganizationManagedRule("exampleOrganizationManagedRule", new()
    ///     {
    ///         RuleIdentifier = "IAM_PASSWORD_POLICY",
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             exampleOrganization,
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
    ///  to = aws_config_organization_managed_rule.example
    /// 
    ///  id = "example" } Using `pulumi import`, import Config Organization Managed Rules using the name. For exampleconsole % pulumi import aws_config_organization_managed_rule.example example
    /// </summary>
    [AwsResourceType("aws:cfg/organizationManagedRule:OrganizationManagedRule")]
    public partial class OrganizationManagedRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the rule
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Description of the rule
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// List of AWS account identifiers to exclude from the rule
        /// </summary>
        [Output("excludedAccounts")]
        public Output<ImmutableArray<string>> ExcludedAccounts { get; private set; } = null!;

        /// <summary>
        /// A string in JSON format that is passed to the AWS Config Rule Lambda Function
        /// </summary>
        [Output("inputParameters")]
        public Output<string?> InputParameters { get; private set; } = null!;

        /// <summary>
        /// The maximum frequency with which AWS Config runs evaluations for a rule, if the rule is triggered at a periodic frequency. Defaults to `TwentyFour_Hours` for periodic frequency triggered rules. Valid values: `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, or `TwentyFour_Hours`.
        /// </summary>
        [Output("maximumExecutionFrequency")]
        public Output<string?> MaximumExecutionFrequency { get; private set; } = null!;

        /// <summary>
        /// The name of the rule
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Identifier of the AWS resource to evaluate
        /// </summary>
        [Output("resourceIdScope")]
        public Output<string?> ResourceIdScope { get; private set; } = null!;

        /// <summary>
        /// List of types of AWS resources to evaluate
        /// </summary>
        [Output("resourceTypesScopes")]
        public Output<ImmutableArray<string>> ResourceTypesScopes { get; private set; } = null!;

        /// <summary>
        /// Identifier of an available AWS Config Managed Rule to call. For available values, see the [List of AWS Config Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/managed-rules-by-aws-config.html) documentation
        /// </summary>
        [Output("ruleIdentifier")]
        public Output<string> RuleIdentifier { get; private set; } = null!;

        /// <summary>
        /// Tag key of AWS resources to evaluate
        /// </summary>
        [Output("tagKeyScope")]
        public Output<string?> TagKeyScope { get; private set; } = null!;

        /// <summary>
        /// Tag value of AWS resources to evaluate
        /// </summary>
        [Output("tagValueScope")]
        public Output<string?> TagValueScope { get; private set; } = null!;


        /// <summary>
        /// Create a OrganizationManagedRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OrganizationManagedRule(string name, OrganizationManagedRuleArgs args, CustomResourceOptions? options = null)
            : base("aws:cfg/organizationManagedRule:OrganizationManagedRule", name, args ?? new OrganizationManagedRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OrganizationManagedRule(string name, Input<string> id, OrganizationManagedRuleState? state = null, CustomResourceOptions? options = null)
            : base("aws:cfg/organizationManagedRule:OrganizationManagedRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing OrganizationManagedRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OrganizationManagedRule Get(string name, Input<string> id, OrganizationManagedRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new OrganizationManagedRule(name, id, state, options);
        }
    }

    public sealed class OrganizationManagedRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the rule
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("excludedAccounts")]
        private InputList<string>? _excludedAccounts;

        /// <summary>
        /// List of AWS account identifiers to exclude from the rule
        /// </summary>
        public InputList<string> ExcludedAccounts
        {
            get => _excludedAccounts ?? (_excludedAccounts = new InputList<string>());
            set => _excludedAccounts = value;
        }

        /// <summary>
        /// A string in JSON format that is passed to the AWS Config Rule Lambda Function
        /// </summary>
        [Input("inputParameters")]
        public Input<string>? InputParameters { get; set; }

        /// <summary>
        /// The maximum frequency with which AWS Config runs evaluations for a rule, if the rule is triggered at a periodic frequency. Defaults to `TwentyFour_Hours` for periodic frequency triggered rules. Valid values: `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, or `TwentyFour_Hours`.
        /// </summary>
        [Input("maximumExecutionFrequency")]
        public Input<string>? MaximumExecutionFrequency { get; set; }

        /// <summary>
        /// The name of the rule
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Identifier of the AWS resource to evaluate
        /// </summary>
        [Input("resourceIdScope")]
        public Input<string>? ResourceIdScope { get; set; }

        [Input("resourceTypesScopes")]
        private InputList<string>? _resourceTypesScopes;

        /// <summary>
        /// List of types of AWS resources to evaluate
        /// </summary>
        public InputList<string> ResourceTypesScopes
        {
            get => _resourceTypesScopes ?? (_resourceTypesScopes = new InputList<string>());
            set => _resourceTypesScopes = value;
        }

        /// <summary>
        /// Identifier of an available AWS Config Managed Rule to call. For available values, see the [List of AWS Config Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/managed-rules-by-aws-config.html) documentation
        /// </summary>
        [Input("ruleIdentifier", required: true)]
        public Input<string> RuleIdentifier { get; set; } = null!;

        /// <summary>
        /// Tag key of AWS resources to evaluate
        /// </summary>
        [Input("tagKeyScope")]
        public Input<string>? TagKeyScope { get; set; }

        /// <summary>
        /// Tag value of AWS resources to evaluate
        /// </summary>
        [Input("tagValueScope")]
        public Input<string>? TagValueScope { get; set; }

        public OrganizationManagedRuleArgs()
        {
        }
        public static new OrganizationManagedRuleArgs Empty => new OrganizationManagedRuleArgs();
    }

    public sealed class OrganizationManagedRuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the rule
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Description of the rule
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("excludedAccounts")]
        private InputList<string>? _excludedAccounts;

        /// <summary>
        /// List of AWS account identifiers to exclude from the rule
        /// </summary>
        public InputList<string> ExcludedAccounts
        {
            get => _excludedAccounts ?? (_excludedAccounts = new InputList<string>());
            set => _excludedAccounts = value;
        }

        /// <summary>
        /// A string in JSON format that is passed to the AWS Config Rule Lambda Function
        /// </summary>
        [Input("inputParameters")]
        public Input<string>? InputParameters { get; set; }

        /// <summary>
        /// The maximum frequency with which AWS Config runs evaluations for a rule, if the rule is triggered at a periodic frequency. Defaults to `TwentyFour_Hours` for periodic frequency triggered rules. Valid values: `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, or `TwentyFour_Hours`.
        /// </summary>
        [Input("maximumExecutionFrequency")]
        public Input<string>? MaximumExecutionFrequency { get; set; }

        /// <summary>
        /// The name of the rule
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Identifier of the AWS resource to evaluate
        /// </summary>
        [Input("resourceIdScope")]
        public Input<string>? ResourceIdScope { get; set; }

        [Input("resourceTypesScopes")]
        private InputList<string>? _resourceTypesScopes;

        /// <summary>
        /// List of types of AWS resources to evaluate
        /// </summary>
        public InputList<string> ResourceTypesScopes
        {
            get => _resourceTypesScopes ?? (_resourceTypesScopes = new InputList<string>());
            set => _resourceTypesScopes = value;
        }

        /// <summary>
        /// Identifier of an available AWS Config Managed Rule to call. For available values, see the [List of AWS Config Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/managed-rules-by-aws-config.html) documentation
        /// </summary>
        [Input("ruleIdentifier")]
        public Input<string>? RuleIdentifier { get; set; }

        /// <summary>
        /// Tag key of AWS resources to evaluate
        /// </summary>
        [Input("tagKeyScope")]
        public Input<string>? TagKeyScope { get; set; }

        /// <summary>
        /// Tag value of AWS resources to evaluate
        /// </summary>
        [Input("tagValueScope")]
        public Input<string>? TagValueScope { get; set; }

        public OrganizationManagedRuleState()
        {
        }
        public static new OrganizationManagedRuleState Empty => new OrganizationManagedRuleState();
    }
}
