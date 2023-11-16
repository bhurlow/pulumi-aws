// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Budgets
{
    /// <summary>
    /// Provides a budgets budget resource. Budgets use the cost visualisation provided by Cost Explorer to show you the status of your budgets, to provide forecasts of your estimated costs, and to track your AWS usage, including your free tier usage.
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
    ///     var ec2 = new Aws.Budgets.Budget("ec2", new()
    ///     {
    ///         BudgetType = "COST",
    ///         CostFilters = new[]
    ///         {
    ///             new Aws.Budgets.Inputs.BudgetCostFilterArgs
    ///             {
    ///                 Name = "Service",
    ///                 Values = new[]
    ///                 {
    ///                     "Amazon Elastic Compute Cloud - Compute",
    ///                 },
    ///             },
    ///         },
    ///         LimitAmount = "1200",
    ///         LimitUnit = "USD",
    ///         Notifications = new[]
    ///         {
    ///             new Aws.Budgets.Inputs.BudgetNotificationArgs
    ///             {
    ///                 ComparisonOperator = "GREATER_THAN",
    ///                 NotificationType = "FORECASTED",
    ///                 SubscriberEmailAddresses = new[]
    ///                 {
    ///                     "test@example.com",
    ///                 },
    ///                 Threshold = 100,
    ///                 ThresholdType = "PERCENTAGE",
    ///             },
    ///         },
    ///         TimePeriodEnd = "2087-06-15_00:00",
    ///         TimePeriodStart = "2017-07-01_00:00",
    ///         TimeUnit = "MONTHLY",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// Create a budget for *$100*.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var cost = new Aws.Budgets.Budget("cost", new()
    ///     {
    ///         BudgetType = "COST",
    ///         LimitAmount = "100",
    ///         LimitUnit = "USD",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// Create a budget with planned budget limits.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var cost = new Aws.Budgets.Budget("cost", new()
    ///     {
    ///         PlannedLimits = new[]
    ///         {
    ///             new Aws.Budgets.Inputs.BudgetPlannedLimitArgs
    ///             {
    ///                 Amount = "100",
    ///                 StartTime = "2017-07-01_00:00",
    ///                 Unit = "USD",
    ///             },
    ///             new Aws.Budgets.Inputs.BudgetPlannedLimitArgs
    ///             {
    ///                 Amount = "200",
    ///                 StartTime = "2017-08-01_00:00",
    ///                 Unit = "USD",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// Create a budget for s3 with a limit of *3 GB* of storage.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var s3 = new Aws.Budgets.Budget("s3", new()
    ///     {
    ///         BudgetType = "USAGE",
    ///         LimitAmount = "3",
    ///         LimitUnit = "GB",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// Create a Savings Plan Utilization Budget
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var savingsPlanUtilization = new Aws.Budgets.Budget("savingsPlanUtilization", new()
    ///     {
    ///         BudgetType = "SAVINGS_PLANS_UTILIZATION",
    ///         CostTypes = new Aws.Budgets.Inputs.BudgetCostTypesArgs
    ///         {
    ///             IncludeCredit = false,
    ///             IncludeDiscount = false,
    ///             IncludeOtherSubscription = false,
    ///             IncludeRecurring = false,
    ///             IncludeRefund = false,
    ///             IncludeSubscription = true,
    ///             IncludeSupport = false,
    ///             IncludeTax = false,
    ///             IncludeUpfront = false,
    ///             UseBlended = false,
    ///         },
    ///         LimitAmount = "100.0",
    ///         LimitUnit = "PERCENTAGE",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// Create a RI Utilization Budget
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var riUtilization = new Aws.Budgets.Budget("riUtilization", new()
    ///     {
    ///         BudgetType = "RI_UTILIZATION",
    ///         CostFilters = new[]
    ///         {
    ///             new Aws.Budgets.Inputs.BudgetCostFilterArgs
    ///             {
    ///                 Name = "Service",
    ///                 Values = new[]
    ///                 {
    ///                     "Amazon Relational Database Service",
    ///                 },
    ///             },
    ///         },
    ///         CostTypes = new Aws.Budgets.Inputs.BudgetCostTypesArgs
    ///         {
    ///             IncludeCredit = false,
    ///             IncludeDiscount = false,
    ///             IncludeOtherSubscription = false,
    ///             IncludeRecurring = false,
    ///             IncludeRefund = false,
    ///             IncludeSubscription = true,
    ///             IncludeSupport = false,
    ///             IncludeTax = false,
    ///             IncludeUpfront = false,
    ///             UseBlended = false,
    ///         },
    ///         LimitAmount = "100.0",
    ///         LimitUnit = "PERCENTAGE",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// Create a Cost Filter using Resource Tags
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var cost = new Aws.Budgets.Budget("cost", new()
    ///     {
    ///         CostFilters = new[]
    ///         {
    ///             new Aws.Budgets.Inputs.BudgetCostFilterArgs
    ///             {
    ///                 Name = "TagKeyValue",
    ///                 Values = new[]
    ///                 {
    ///                     "TagKey$TagValue",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// Create a cost_filter using resource tags, obtaining the tag value from a variable
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var cost = new Aws.Budgets.Budget("cost", new()
    ///     {
    ///         CostFilters = new[]
    ///         {
    ///             new Aws.Budgets.Inputs.BudgetCostFilterArgs
    ///             {
    ///                 Name = "TagKeyValue",
    ///                 Values = new[]
    ///                 {
    ///                     "TagKey${var.TagValue}",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Using `pulumi import`, import budgets using `AccountID:BudgetName`. For example:
    /// 
    /// ```sh
    ///  $ pulumi import aws:budgets/budget:Budget myBudget 123456789012:myBudget
    /// ```
    /// </summary>
    [AwsResourceType("aws:budgets/budget:Budget")]
    public partial class Budget : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the target account for budget. Will use current user's account_id by default if omitted.
        /// </summary>
        [Output("accountId")]
        public Output<string> AccountId { get; private set; } = null!;

        /// <summary>
        /// The ARN of the budget.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Object containing [AutoAdjustData] which determines the budget amount for an auto-adjusting budget.
        /// </summary>
        [Output("autoAdjustData")]
        public Output<Outputs.BudgetAutoAdjustData?> AutoAdjustData { get; private set; } = null!;

        /// <summary>
        /// Whether this budget tracks monetary cost or usage.
        /// </summary>
        [Output("budgetType")]
        public Output<string> BudgetType { get; private set; } = null!;

        /// <summary>
        /// A list of CostFilter name/values pair to apply to budget.
        /// </summary>
        [Output("costFilters")]
        public Output<ImmutableArray<Outputs.BudgetCostFilter>> CostFilters { get; private set; } = null!;

        /// <summary>
        /// Object containing CostTypes The types of cost included in a budget, such as tax and subscriptions.
        /// </summary>
        [Output("costTypes")]
        public Output<Outputs.BudgetCostTypes> CostTypes { get; private set; } = null!;

        /// <summary>
        /// The amount of cost or usage being measured for a budget.
        /// </summary>
        [Output("limitAmount")]
        public Output<string> LimitAmount { get; private set; } = null!;

        /// <summary>
        /// The unit of measurement used for the budget forecast, actual spend, or budget threshold, such as dollars or GB. See [Spend](http://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/data-type-spend.html) documentation.
        /// </summary>
        [Output("limitUnit")]
        public Output<string> LimitUnit { get; private set; } = null!;

        /// <summary>
        /// The name of a budget. Unique within accounts.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The prefix of the name of a budget. Unique within accounts.
        /// </summary>
        [Output("namePrefix")]
        public Output<string> NamePrefix { get; private set; } = null!;

        /// <summary>
        /// Object containing Budget Notifications. Can be used multiple times to define more than one budget notification.
        /// </summary>
        [Output("notifications")]
        public Output<ImmutableArray<Outputs.BudgetNotification>> Notifications { get; private set; } = null!;

        /// <summary>
        /// Object containing Planned Budget Limits. Can be used multiple times to plan more than one budget limit. See [PlannedBudgetLimits](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_budgets_Budget.html#awscostmanagement-Type-budgets_Budget-PlannedBudgetLimits) documentation.
        /// </summary>
        [Output("plannedLimits")]
        public Output<ImmutableArray<Outputs.BudgetPlannedLimit>> PlannedLimits { get; private set; } = null!;

        /// <summary>
        /// The end of the time period covered by the budget. There are no restrictions on the end date. Format: `2017-01-01_12:00`.
        /// </summary>
        [Output("timePeriodEnd")]
        public Output<string?> TimePeriodEnd { get; private set; } = null!;

        /// <summary>
        /// The start of the time period covered by the budget. If you don't specify a start date, AWS defaults to the start of your chosen time period. The start date must come before the end date. Format: `2017-01-01_12:00`.
        /// </summary>
        [Output("timePeriodStart")]
        public Output<string> TimePeriodStart { get; private set; } = null!;

        /// <summary>
        /// The length of time until a budget resets the actual and forecasted spend. Valid values: `MONTHLY`, `QUARTERLY`, `ANNUALLY`, and `DAILY`.
        /// </summary>
        [Output("timeUnit")]
        public Output<string> TimeUnit { get; private set; } = null!;


        /// <summary>
        /// Create a Budget resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Budget(string name, BudgetArgs args, CustomResourceOptions? options = null)
            : base("aws:budgets/budget:Budget", name, args ?? new BudgetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Budget(string name, Input<string> id, BudgetState? state = null, CustomResourceOptions? options = null)
            : base("aws:budgets/budget:Budget", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Budget resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Budget Get(string name, Input<string> id, BudgetState? state = null, CustomResourceOptions? options = null)
        {
            return new Budget(name, id, state, options);
        }
    }

    public sealed class BudgetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the target account for budget. Will use current user's account_id by default if omitted.
        /// </summary>
        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        /// <summary>
        /// Object containing [AutoAdjustData] which determines the budget amount for an auto-adjusting budget.
        /// </summary>
        [Input("autoAdjustData")]
        public Input<Inputs.BudgetAutoAdjustDataArgs>? AutoAdjustData { get; set; }

        /// <summary>
        /// Whether this budget tracks monetary cost or usage.
        /// </summary>
        [Input("budgetType", required: true)]
        public Input<string> BudgetType { get; set; } = null!;

        [Input("costFilters")]
        private InputList<Inputs.BudgetCostFilterArgs>? _costFilters;

        /// <summary>
        /// A list of CostFilter name/values pair to apply to budget.
        /// </summary>
        public InputList<Inputs.BudgetCostFilterArgs> CostFilters
        {
            get => _costFilters ?? (_costFilters = new InputList<Inputs.BudgetCostFilterArgs>());
            set => _costFilters = value;
        }

        /// <summary>
        /// Object containing CostTypes The types of cost included in a budget, such as tax and subscriptions.
        /// </summary>
        [Input("costTypes")]
        public Input<Inputs.BudgetCostTypesArgs>? CostTypes { get; set; }

        /// <summary>
        /// The amount of cost or usage being measured for a budget.
        /// </summary>
        [Input("limitAmount")]
        public Input<string>? LimitAmount { get; set; }

        /// <summary>
        /// The unit of measurement used for the budget forecast, actual spend, or budget threshold, such as dollars or GB. See [Spend](http://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/data-type-spend.html) documentation.
        /// </summary>
        [Input("limitUnit")]
        public Input<string>? LimitUnit { get; set; }

        /// <summary>
        /// The name of a budget. Unique within accounts.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The prefix of the name of a budget. Unique within accounts.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        [Input("notifications")]
        private InputList<Inputs.BudgetNotificationArgs>? _notifications;

        /// <summary>
        /// Object containing Budget Notifications. Can be used multiple times to define more than one budget notification.
        /// </summary>
        public InputList<Inputs.BudgetNotificationArgs> Notifications
        {
            get => _notifications ?? (_notifications = new InputList<Inputs.BudgetNotificationArgs>());
            set => _notifications = value;
        }

        [Input("plannedLimits")]
        private InputList<Inputs.BudgetPlannedLimitArgs>? _plannedLimits;

        /// <summary>
        /// Object containing Planned Budget Limits. Can be used multiple times to plan more than one budget limit. See [PlannedBudgetLimits](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_budgets_Budget.html#awscostmanagement-Type-budgets_Budget-PlannedBudgetLimits) documentation.
        /// </summary>
        public InputList<Inputs.BudgetPlannedLimitArgs> PlannedLimits
        {
            get => _plannedLimits ?? (_plannedLimits = new InputList<Inputs.BudgetPlannedLimitArgs>());
            set => _plannedLimits = value;
        }

        /// <summary>
        /// The end of the time period covered by the budget. There are no restrictions on the end date. Format: `2017-01-01_12:00`.
        /// </summary>
        [Input("timePeriodEnd")]
        public Input<string>? TimePeriodEnd { get; set; }

        /// <summary>
        /// The start of the time period covered by the budget. If you don't specify a start date, AWS defaults to the start of your chosen time period. The start date must come before the end date. Format: `2017-01-01_12:00`.
        /// </summary>
        [Input("timePeriodStart")]
        public Input<string>? TimePeriodStart { get; set; }

        /// <summary>
        /// The length of time until a budget resets the actual and forecasted spend. Valid values: `MONTHLY`, `QUARTERLY`, `ANNUALLY`, and `DAILY`.
        /// </summary>
        [Input("timeUnit", required: true)]
        public Input<string> TimeUnit { get; set; } = null!;

        public BudgetArgs()
        {
        }
        public static new BudgetArgs Empty => new BudgetArgs();
    }

    public sealed class BudgetState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the target account for budget. Will use current user's account_id by default if omitted.
        /// </summary>
        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        /// <summary>
        /// The ARN of the budget.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Object containing [AutoAdjustData] which determines the budget amount for an auto-adjusting budget.
        /// </summary>
        [Input("autoAdjustData")]
        public Input<Inputs.BudgetAutoAdjustDataGetArgs>? AutoAdjustData { get; set; }

        /// <summary>
        /// Whether this budget tracks monetary cost or usage.
        /// </summary>
        [Input("budgetType")]
        public Input<string>? BudgetType { get; set; }

        [Input("costFilters")]
        private InputList<Inputs.BudgetCostFilterGetArgs>? _costFilters;

        /// <summary>
        /// A list of CostFilter name/values pair to apply to budget.
        /// </summary>
        public InputList<Inputs.BudgetCostFilterGetArgs> CostFilters
        {
            get => _costFilters ?? (_costFilters = new InputList<Inputs.BudgetCostFilterGetArgs>());
            set => _costFilters = value;
        }

        /// <summary>
        /// Object containing CostTypes The types of cost included in a budget, such as tax and subscriptions.
        /// </summary>
        [Input("costTypes")]
        public Input<Inputs.BudgetCostTypesGetArgs>? CostTypes { get; set; }

        /// <summary>
        /// The amount of cost or usage being measured for a budget.
        /// </summary>
        [Input("limitAmount")]
        public Input<string>? LimitAmount { get; set; }

        /// <summary>
        /// The unit of measurement used for the budget forecast, actual spend, or budget threshold, such as dollars or GB. See [Spend](http://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/data-type-spend.html) documentation.
        /// </summary>
        [Input("limitUnit")]
        public Input<string>? LimitUnit { get; set; }

        /// <summary>
        /// The name of a budget. Unique within accounts.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The prefix of the name of a budget. Unique within accounts.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        [Input("notifications")]
        private InputList<Inputs.BudgetNotificationGetArgs>? _notifications;

        /// <summary>
        /// Object containing Budget Notifications. Can be used multiple times to define more than one budget notification.
        /// </summary>
        public InputList<Inputs.BudgetNotificationGetArgs> Notifications
        {
            get => _notifications ?? (_notifications = new InputList<Inputs.BudgetNotificationGetArgs>());
            set => _notifications = value;
        }

        [Input("plannedLimits")]
        private InputList<Inputs.BudgetPlannedLimitGetArgs>? _plannedLimits;

        /// <summary>
        /// Object containing Planned Budget Limits. Can be used multiple times to plan more than one budget limit. See [PlannedBudgetLimits](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_budgets_Budget.html#awscostmanagement-Type-budgets_Budget-PlannedBudgetLimits) documentation.
        /// </summary>
        public InputList<Inputs.BudgetPlannedLimitGetArgs> PlannedLimits
        {
            get => _plannedLimits ?? (_plannedLimits = new InputList<Inputs.BudgetPlannedLimitGetArgs>());
            set => _plannedLimits = value;
        }

        /// <summary>
        /// The end of the time period covered by the budget. There are no restrictions on the end date. Format: `2017-01-01_12:00`.
        /// </summary>
        [Input("timePeriodEnd")]
        public Input<string>? TimePeriodEnd { get; set; }

        /// <summary>
        /// The start of the time period covered by the budget. If you don't specify a start date, AWS defaults to the start of your chosen time period. The start date must come before the end date. Format: `2017-01-01_12:00`.
        /// </summary>
        [Input("timePeriodStart")]
        public Input<string>? TimePeriodStart { get; set; }

        /// <summary>
        /// The length of time until a budget resets the actual and forecasted spend. Valid values: `MONTHLY`, `QUARTERLY`, `ANNUALLY`, and `DAILY`.
        /// </summary>
        [Input("timeUnit")]
        public Input<string>? TimeUnit { get; set; }

        public BudgetState()
        {
        }
        public static new BudgetState Empty => new BudgetState();
    }
}
