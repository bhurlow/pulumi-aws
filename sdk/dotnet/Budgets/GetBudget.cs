// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Budgets
{
    public static class GetBudget
    {
        /// <summary>
        /// Data source for managing an AWS Web Services Budgets Budget.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### Basic Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var test = Aws.Budgets.GetBudget.Invoke(new()
        ///     {
        ///         Name = aws_budgets_budget.Test.Name,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetBudgetResult> InvokeAsync(GetBudgetArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBudgetResult>("aws:budgets/getBudget:getBudget", args ?? new GetBudgetArgs(), options.WithDefaults());

        /// <summary>
        /// Data source for managing an AWS Web Services Budgets Budget.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### Basic Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var test = Aws.Budgets.GetBudget.Invoke(new()
        ///     {
        ///         Name = aws_budgets_budget.Test.Name,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetBudgetResult> Invoke(GetBudgetInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBudgetResult>("aws:budgets/getBudget:getBudget", args ?? new GetBudgetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBudgetArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the target account for budget. Will use current user's account_id by default if omitted.
        /// </summary>
        [Input("accountId")]
        public string? AccountId { get; set; }

        /// <summary>
        /// The name of a budget. Unique within accounts.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The prefix of the name of a budget. Unique within accounts.
        /// </summary>
        [Input("namePrefix")]
        public string? NamePrefix { get; set; }

        public GetBudgetArgs()
        {
        }
        public static new GetBudgetArgs Empty => new GetBudgetArgs();
    }

    public sealed class GetBudgetInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the target account for budget. Will use current user's account_id by default if omitted.
        /// </summary>
        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        /// <summary>
        /// The name of a budget. Unique within accounts.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The prefix of the name of a budget. Unique within accounts.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        public GetBudgetInvokeArgs()
        {
        }
        public static new GetBudgetInvokeArgs Empty => new GetBudgetInvokeArgs();
    }


    [OutputType]
    public sealed class GetBudgetResult
    {
        public readonly string AccountId;
        public readonly string Arn;
        /// <summary>
        /// Object containing [AutoAdjustData] which determines the budget amount for an auto-adjusting budget.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBudgetAutoAdjustDataResult> AutoAdjustDatas;
        /// <summary>
        /// Boolean indicating whether this budget has been exceeded.
        /// </summary>
        public readonly bool BudgetExceeded;
        /// <summary>
        /// The total amount of cost, usage, RI utilization, RI coverage, Savings Plans utilization, or Savings Plans coverage that you want to track with your budget. Contains object Spend.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBudgetBudgetLimitResult> BudgetLimits;
        /// <summary>
        /// Whether this budget tracks monetary cost or usage.
        /// </summary>
        public readonly string BudgetType;
        /// <summary>
        /// The spend objects that are associated with this budget. The actualSpend tracks how much you've used, cost, usage, RI units, or Savings Plans units and the forecastedSpend tracks how much that you're predicted to spend based on your historical usage profile.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBudgetCalculatedSpendResult> CalculatedSpends;
        /// <summary>
        /// A list of CostFilter name/values pair to apply to budget.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBudgetCostFilterResult> CostFilters;
        /// <summary>
        /// Object containing CostTypes The types of cost included in a budget, such as tax and subscriptions.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBudgetCostTypeResult> CostTypes;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string? NamePrefix;
        /// <summary>
        /// Object containing Budget Notifications. Can be used multiple times to define more than one budget notification.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBudgetNotificationResult> Notifications;
        /// <summary>
        /// Object containing Planned Budget Limits. Can be used multiple times to plan more than one budget limit. See [PlannedBudgetLimits](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_budgets_Budget.html#awscostmanagement-Type-budgets_Budget-PlannedBudgetLimits) documentation.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBudgetPlannedLimitResult> PlannedLimits;
        /// <summary>
        /// The end of the time period covered by the budget. There are no restrictions on the end date. Format: `2017-01-01_12:00`.
        /// </summary>
        public readonly string TimePeriodEnd;
        /// <summary>
        /// The start of the time period covered by the budget. If you don't specify a start date, AWS defaults to the start of your chosen time period. The start date must come before the end date. Format: `2017-01-01_12:00`.
        /// </summary>
        public readonly string TimePeriodStart;
        /// <summary>
        /// The length of time until a budget resets the actual and forecasted spend. Valid values: `MONTHLY`, `QUARTERLY`, `ANNUALLY`, and `DAILY`.
        /// </summary>
        public readonly string TimeUnit;

        [OutputConstructor]
        private GetBudgetResult(
            string accountId,

            string arn,

            ImmutableArray<Outputs.GetBudgetAutoAdjustDataResult> autoAdjustDatas,

            bool budgetExceeded,

            ImmutableArray<Outputs.GetBudgetBudgetLimitResult> budgetLimits,

            string budgetType,

            ImmutableArray<Outputs.GetBudgetCalculatedSpendResult> calculatedSpends,

            ImmutableArray<Outputs.GetBudgetCostFilterResult> costFilters,

            ImmutableArray<Outputs.GetBudgetCostTypeResult> costTypes,

            string id,

            string name,

            string? namePrefix,

            ImmutableArray<Outputs.GetBudgetNotificationResult> notifications,

            ImmutableArray<Outputs.GetBudgetPlannedLimitResult> plannedLimits,

            string timePeriodEnd,

            string timePeriodStart,

            string timeUnit)
        {
            AccountId = accountId;
            Arn = arn;
            AutoAdjustDatas = autoAdjustDatas;
            BudgetExceeded = budgetExceeded;
            BudgetLimits = budgetLimits;
            BudgetType = budgetType;
            CalculatedSpends = calculatedSpends;
            CostFilters = costFilters;
            CostTypes = costTypes;
            Id = id;
            Name = name;
            NamePrefix = namePrefix;
            Notifications = notifications;
            PlannedLimits = plannedLimits;
            TimePeriodEnd = timePeriodEnd;
            TimePeriodStart = timePeriodStart;
            TimeUnit = timeUnit;
        }
    }
}
