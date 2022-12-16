// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.budgets;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.budgets.BudgetArgs;
import com.pulumi.aws.budgets.inputs.BudgetState;
import com.pulumi.aws.budgets.outputs.BudgetAutoAdjustData;
import com.pulumi.aws.budgets.outputs.BudgetCostFilter;
import com.pulumi.aws.budgets.outputs.BudgetCostTypes;
import com.pulumi.aws.budgets.outputs.BudgetNotification;
import com.pulumi.aws.budgets.outputs.BudgetPlannedLimit;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a budgets budget resource. Budgets use the cost visualisation provided by Cost Explorer to show you the status of your budgets, to provide forecasts of your estimated costs, and to track your AWS usage, including your free tier usage.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.budgets.Budget;
 * import com.pulumi.aws.budgets.BudgetArgs;
 * import com.pulumi.aws.budgets.inputs.BudgetCostFilterArgs;
 * import com.pulumi.aws.budgets.inputs.BudgetNotificationArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var ec2 = new Budget(&#34;ec2&#34;, BudgetArgs.builder()        
 *             .budgetType(&#34;COST&#34;)
 *             .costFilters(BudgetCostFilterArgs.builder()
 *                 .name(&#34;Service&#34;)
 *                 .values(&#34;Amazon Elastic Compute Cloud - Compute&#34;)
 *                 .build())
 *             .limitAmount(&#34;1200&#34;)
 *             .limitUnit(&#34;USD&#34;)
 *             .notifications(BudgetNotificationArgs.builder()
 *                 .comparisonOperator(&#34;GREATER_THAN&#34;)
 *                 .notificationType(&#34;FORECASTED&#34;)
 *                 .subscriberEmailAddresses(&#34;test@example.com&#34;)
 *                 .threshold(100)
 *                 .thresholdType(&#34;PERCENTAGE&#34;)
 *                 .build())
 *             .timePeriodEnd(&#34;2087-06-15_00:00&#34;)
 *             .timePeriodStart(&#34;2017-07-01_00:00&#34;)
 *             .timeUnit(&#34;MONTHLY&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * Create a budget for *$100*.
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.budgets.Budget;
 * import com.pulumi.aws.budgets.BudgetArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var cost = new Budget(&#34;cost&#34;, BudgetArgs.builder()        
 *             .budgetType(&#34;COST&#34;)
 *             .limitAmount(&#34;100&#34;)
 *             .limitUnit(&#34;USD&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * Create a budget with planned budget limits.
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.budgets.Budget;
 * import com.pulumi.aws.budgets.BudgetArgs;
 * import com.pulumi.aws.budgets.inputs.BudgetPlannedLimitArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var cost = new Budget(&#34;cost&#34;, BudgetArgs.builder()        
 *             .plannedLimits(            
 *                 BudgetPlannedLimitArgs.builder()
 *                     .amount(&#34;100&#34;)
 *                     .startTime(&#34;2017-07-01_00:00&#34;)
 *                     .unit(&#34;USD&#34;)
 *                     .build(),
 *                 BudgetPlannedLimitArgs.builder()
 *                     .amount(&#34;200&#34;)
 *                     .startTime(&#34;2017-08-01_00:00&#34;)
 *                     .unit(&#34;USD&#34;)
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * Create a budget for s3 with a limit of *3 GB* of storage.
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.budgets.Budget;
 * import com.pulumi.aws.budgets.BudgetArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var s3 = new Budget(&#34;s3&#34;, BudgetArgs.builder()        
 *             .budgetType(&#34;USAGE&#34;)
 *             .limitAmount(&#34;3&#34;)
 *             .limitUnit(&#34;GB&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * Create a Savings Plan Utilization Budget
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.budgets.Budget;
 * import com.pulumi.aws.budgets.BudgetArgs;
 * import com.pulumi.aws.budgets.inputs.BudgetCostTypesArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var savingsPlanUtilization = new Budget(&#34;savingsPlanUtilization&#34;, BudgetArgs.builder()        
 *             .budgetType(&#34;SAVINGS_PLANS_UTILIZATION&#34;)
 *             .costTypes(BudgetCostTypesArgs.builder()
 *                 .includeCredit(false)
 *                 .includeDiscount(false)
 *                 .includeOtherSubscription(false)
 *                 .includeRecurring(false)
 *                 .includeRefund(false)
 *                 .includeSubscription(true)
 *                 .includeSupport(false)
 *                 .includeTax(false)
 *                 .includeUpfront(false)
 *                 .useBlended(false)
 *                 .build())
 *             .limitAmount(&#34;100.0&#34;)
 *             .limitUnit(&#34;PERCENTAGE&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * Create a RI Utilization Budget
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.budgets.Budget;
 * import com.pulumi.aws.budgets.BudgetArgs;
 * import com.pulumi.aws.budgets.inputs.BudgetCostFilterArgs;
 * import com.pulumi.aws.budgets.inputs.BudgetCostTypesArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var riUtilization = new Budget(&#34;riUtilization&#34;, BudgetArgs.builder()        
 *             .budgetType(&#34;RI_UTILIZATION&#34;)
 *             .costFilters(BudgetCostFilterArgs.builder()
 *                 .name(&#34;Service&#34;)
 *                 .values(&#34;Amazon Relational Database Service&#34;)
 *                 .build())
 *             .costTypes(BudgetCostTypesArgs.builder()
 *                 .includeCredit(false)
 *                 .includeDiscount(false)
 *                 .includeOtherSubscription(false)
 *                 .includeRecurring(false)
 *                 .includeRefund(false)
 *                 .includeSubscription(true)
 *                 .includeSupport(false)
 *                 .includeTax(false)
 *                 .includeUpfront(false)
 *                 .useBlended(false)
 *                 .build())
 *             .limitAmount(&#34;100.0&#34;)
 *             .limitUnit(&#34;PERCENTAGE&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Budgets can be imported using `AccountID:BudgetName`, e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:budgets/budget:Budget myBudget 123456789012:myBudget`
 * ```
 * 
 */
@ResourceType(type="aws:budgets/budget:Budget")
public class Budget extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the target account for budget. Will use current user&#39;s account_id by default if omitted.
     * 
     */
    @Export(name="accountId", refs={String.class}, tree="[0]")
    private Output<String> accountId;

    /**
     * @return The ID of the target account for budget. Will use current user&#39;s account_id by default if omitted.
     * 
     */
    public Output<String> accountId() {
        return this.accountId;
    }
    /**
     * The ARN of the budget.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The ARN of the budget.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * Object containing [AutoAdjustData] which determines the budget amount for an auto-adjusting budget.
     * 
     */
    @Export(name="autoAdjustData", refs={BudgetAutoAdjustData.class}, tree="[0]")
    private Output</* @Nullable */ BudgetAutoAdjustData> autoAdjustData;

    /**
     * @return Object containing [AutoAdjustData] which determines the budget amount for an auto-adjusting budget.
     * 
     */
    public Output<Optional<BudgetAutoAdjustData>> autoAdjustData() {
        return Codegen.optional(this.autoAdjustData);
    }
    /**
     * Whether this budget tracks monetary cost or usage.
     * 
     */
    @Export(name="budgetType", refs={String.class}, tree="[0]")
    private Output<String> budgetType;

    /**
     * @return Whether this budget tracks monetary cost or usage.
     * 
     */
    public Output<String> budgetType() {
        return this.budgetType;
    }
    /**
     * Map of CostFilters key/value pairs to apply to the budget.
     * 
     * @deprecated
     * Use the attribute &#34;cost_filter&#34; instead.
     * 
     */
    @Deprecated /* Use the attribute ""cost_filter"" instead. */
    @Export(name="costFilterLegacy", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> costFilterLegacy;

    /**
     * @return Map of CostFilters key/value pairs to apply to the budget.
     * 
     */
    public Output<Map<String,String>> costFilterLegacy() {
        return this.costFilterLegacy;
    }
    /**
     * A list of CostFilter name/values pair to apply to budget.
     * 
     */
    @Export(name="costFilters", refs={List.class,BudgetCostFilter.class}, tree="[0,1]")
    private Output<List<BudgetCostFilter>> costFilters;

    /**
     * @return A list of CostFilter name/values pair to apply to budget.
     * 
     */
    public Output<List<BudgetCostFilter>> costFilters() {
        return this.costFilters;
    }
    /**
     * Object containing CostTypes The types of cost included in a budget, such as tax and subscriptions.
     * 
     */
    @Export(name="costTypes", refs={BudgetCostTypes.class}, tree="[0]")
    private Output<BudgetCostTypes> costTypes;

    /**
     * @return Object containing CostTypes The types of cost included in a budget, such as tax and subscriptions.
     * 
     */
    public Output<BudgetCostTypes> costTypes() {
        return this.costTypes;
    }
    /**
     * The amount of cost or usage being measured for a budget.
     * 
     */
    @Export(name="limitAmount", refs={String.class}, tree="[0]")
    private Output<String> limitAmount;

    /**
     * @return The amount of cost or usage being measured for a budget.
     * 
     */
    public Output<String> limitAmount() {
        return this.limitAmount;
    }
    /**
     * The unit of measurement used for the budget forecast, actual spend, or budget threshold, such as dollars or GB. See [Spend](http://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/data-type-spend.html) documentation.
     * 
     */
    @Export(name="limitUnit", refs={String.class}, tree="[0]")
    private Output<String> limitUnit;

    /**
     * @return The unit of measurement used for the budget forecast, actual spend, or budget threshold, such as dollars or GB. See [Spend](http://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/data-type-spend.html) documentation.
     * 
     */
    public Output<String> limitUnit() {
        return this.limitUnit;
    }
    /**
     * The name of a budget. Unique within accounts.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of a budget. Unique within accounts.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The prefix of the name of a budget. Unique within accounts.
     * 
     */
    @Export(name="namePrefix", refs={String.class}, tree="[0]")
    private Output<String> namePrefix;

    /**
     * @return The prefix of the name of a budget. Unique within accounts.
     * 
     */
    public Output<String> namePrefix() {
        return this.namePrefix;
    }
    /**
     * Object containing Budget Notifications. Can be used multiple times to define more than one budget notification.
     * 
     */
    @Export(name="notifications", refs={List.class,BudgetNotification.class}, tree="[0,1]")
    private Output</* @Nullable */ List<BudgetNotification>> notifications;

    /**
     * @return Object containing Budget Notifications. Can be used multiple times to define more than one budget notification.
     * 
     */
    public Output<Optional<List<BudgetNotification>>> notifications() {
        return Codegen.optional(this.notifications);
    }
    /**
     * Object containing Planned Budget Limits. Can be used multiple times to plan more than one budget limit. See [PlannedBudgetLimits](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_budgets_Budget.html#awscostmanagement-Type-budgets_Budget-PlannedBudgetLimits) documentation.
     * 
     */
    @Export(name="plannedLimits", refs={List.class,BudgetPlannedLimit.class}, tree="[0,1]")
    private Output</* @Nullable */ List<BudgetPlannedLimit>> plannedLimits;

    /**
     * @return Object containing Planned Budget Limits. Can be used multiple times to plan more than one budget limit. See [PlannedBudgetLimits](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_budgets_Budget.html#awscostmanagement-Type-budgets_Budget-PlannedBudgetLimits) documentation.
     * 
     */
    public Output<Optional<List<BudgetPlannedLimit>>> plannedLimits() {
        return Codegen.optional(this.plannedLimits);
    }
    /**
     * The end of the time period covered by the budget. There are no restrictions on the end date. Format: `2017-01-01_12:00`.
     * 
     */
    @Export(name="timePeriodEnd", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> timePeriodEnd;

    /**
     * @return The end of the time period covered by the budget. There are no restrictions on the end date. Format: `2017-01-01_12:00`.
     * 
     */
    public Output<Optional<String>> timePeriodEnd() {
        return Codegen.optional(this.timePeriodEnd);
    }
    /**
     * The start of the time period covered by the budget. If you don&#39;t specify a start date, AWS defaults to the start of your chosen time period. The start date must come before the end date. Format: `2017-01-01_12:00`.
     * 
     */
    @Export(name="timePeriodStart", refs={String.class}, tree="[0]")
    private Output<String> timePeriodStart;

    /**
     * @return The start of the time period covered by the budget. If you don&#39;t specify a start date, AWS defaults to the start of your chosen time period. The start date must come before the end date. Format: `2017-01-01_12:00`.
     * 
     */
    public Output<String> timePeriodStart() {
        return this.timePeriodStart;
    }
    /**
     * The length of time until a budget resets the actual and forecasted spend. Valid values: `MONTHLY`, `QUARTERLY`, `ANNUALLY`, and `DAILY`.
     * 
     */
    @Export(name="timeUnit", refs={String.class}, tree="[0]")
    private Output<String> timeUnit;

    /**
     * @return The length of time until a budget resets the actual and forecasted spend. Valid values: `MONTHLY`, `QUARTERLY`, `ANNUALLY`, and `DAILY`.
     * 
     */
    public Output<String> timeUnit() {
        return this.timeUnit;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Budget(String name) {
        this(name, BudgetArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Budget(String name, BudgetArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Budget(String name, BudgetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:budgets/budget:Budget", name, args == null ? BudgetArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Budget(String name, Output<String> id, @Nullable BudgetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:budgets/budget:Budget", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Budget get(String name, Output<String> id, @Nullable BudgetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Budget(name, id, state, options);
    }
}
