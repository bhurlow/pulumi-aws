// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.cfg;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.cfg.OrganizationCustomRuleArgs;
import com.pulumi.aws.cfg.inputs.OrganizationCustomRuleState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a Config Organization Custom Rule. More information about these rules can be found in the [Enabling AWS Config Rules Across all Accounts in Your Organization](https://docs.aws.amazon.com/config/latest/developerguide/config-rule-multi-account-deployment.html) and [AWS Config Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_use-managed-rules.html) documentation. For working with Organization Managed Rules (those invoking an AWS managed rule), see the `aws_config_organization_managed__rule` resource.
 * 
 * &gt; **NOTE:** This resource must be created in the Organization master account and rules will include the master account unless its ID is added to the `excluded_accounts` argument.
 * 
 * &gt; **NOTE:** The proper Lambda permission to allow the AWS Config service invoke the Lambda Function must be in place before the rule will successfully create or update. See also the `aws.lambda.Permission` resource.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.lambda.Permission;
 * import com.pulumi.aws.lambda.PermissionArgs;
 * import com.pulumi.aws.organizations.Organization;
 * import com.pulumi.aws.organizations.OrganizationArgs;
 * import com.pulumi.aws.cfg.OrganizationCustomRule;
 * import com.pulumi.aws.cfg.OrganizationCustomRuleArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         var examplePermission = new Permission(&#34;examplePermission&#34;, PermissionArgs.builder()        
 *             .action(&#34;lambda:InvokeFunction&#34;)
 *             .function(aws_lambda_function.example().arn())
 *             .principal(&#34;config.amazonaws.com&#34;)
 *             .build());
 * 
 *         var exampleOrganization = new Organization(&#34;exampleOrganization&#34;, OrganizationArgs.builder()        
 *             .awsServiceAccessPrincipals(&#34;config-multiaccountsetup.amazonaws.com&#34;)
 *             .featureSet(&#34;ALL&#34;)
 *             .build());
 * 
 *         var exampleOrganizationCustomRule = new OrganizationCustomRule(&#34;exampleOrganizationCustomRule&#34;, OrganizationCustomRuleArgs.builder()        
 *             .lambdaFunctionArn(aws_lambda_function.example().arn())
 *             .triggerTypes(&#34;ConfigurationItemChangeNotification&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(                
 *                     examplePermission,
 *                     exampleOrganization)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Config Organization Custom Rules can be imported using the name, e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:cfg/organizationCustomRule:OrganizationCustomRule example example
 * ```
 * 
 */
@ResourceType(type="aws:cfg/organizationCustomRule:OrganizationCustomRule")
public class OrganizationCustomRule extends com.pulumi.resources.CustomResource {
    /**
     * Amazon Resource Name (ARN) of the rule
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return Amazon Resource Name (ARN) of the rule
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * Description of the rule
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of the rule
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * List of AWS account identifiers to exclude from the rule
     * 
     */
    @Export(name="excludedAccounts", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> excludedAccounts;

    /**
     * @return List of AWS account identifiers to exclude from the rule
     * 
     */
    public Output<Optional<List<String>>> excludedAccounts() {
        return Codegen.optional(this.excludedAccounts);
    }
    /**
     * A string in JSON format that is passed to the AWS Config Rule Lambda Function
     * 
     */
    @Export(name="inputParameters", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> inputParameters;

    /**
     * @return A string in JSON format that is passed to the AWS Config Rule Lambda Function
     * 
     */
    public Output<Optional<String>> inputParameters() {
        return Codegen.optional(this.inputParameters);
    }
    /**
     * Amazon Resource Name (ARN) of the rule Lambda Function
     * 
     */
    @Export(name="lambdaFunctionArn", refs={String.class}, tree="[0]")
    private Output<String> lambdaFunctionArn;

    /**
     * @return Amazon Resource Name (ARN) of the rule Lambda Function
     * 
     */
    public Output<String> lambdaFunctionArn() {
        return this.lambdaFunctionArn;
    }
    /**
     * The maximum frequency with which AWS Config runs evaluations for a rule, if the rule is triggered at a periodic frequency. Defaults to `TwentyFour_Hours` for periodic frequency triggered rules. Valid values: `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, or `TwentyFour_Hours`.
     * 
     */
    @Export(name="maximumExecutionFrequency", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> maximumExecutionFrequency;

    /**
     * @return The maximum frequency with which AWS Config runs evaluations for a rule, if the rule is triggered at a periodic frequency. Defaults to `TwentyFour_Hours` for periodic frequency triggered rules. Valid values: `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, or `TwentyFour_Hours`.
     * 
     */
    public Output<Optional<String>> maximumExecutionFrequency() {
        return Codegen.optional(this.maximumExecutionFrequency);
    }
    /**
     * The name of the rule
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the rule
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Identifier of the AWS resource to evaluate
     * 
     */
    @Export(name="resourceIdScope", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> resourceIdScope;

    /**
     * @return Identifier of the AWS resource to evaluate
     * 
     */
    public Output<Optional<String>> resourceIdScope() {
        return Codegen.optional(this.resourceIdScope);
    }
    /**
     * List of types of AWS resources to evaluate
     * 
     */
    @Export(name="resourceTypesScopes", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> resourceTypesScopes;

    /**
     * @return List of types of AWS resources to evaluate
     * 
     */
    public Output<Optional<List<String>>> resourceTypesScopes() {
        return Codegen.optional(this.resourceTypesScopes);
    }
    /**
     * Tag key of AWS resources to evaluate
     * 
     */
    @Export(name="tagKeyScope", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> tagKeyScope;

    /**
     * @return Tag key of AWS resources to evaluate
     * 
     */
    public Output<Optional<String>> tagKeyScope() {
        return Codegen.optional(this.tagKeyScope);
    }
    /**
     * Tag value of AWS resources to evaluate
     * 
     */
    @Export(name="tagValueScope", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> tagValueScope;

    /**
     * @return Tag value of AWS resources to evaluate
     * 
     */
    public Output<Optional<String>> tagValueScope() {
        return Codegen.optional(this.tagValueScope);
    }
    /**
     * List of notification types that trigger AWS Config to run an evaluation for the rule. Valid values: `ConfigurationItemChangeNotification`, `OversizedConfigurationItemChangeNotification`, and `ScheduledNotification`
     * 
     */
    @Export(name="triggerTypes", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> triggerTypes;

    /**
     * @return List of notification types that trigger AWS Config to run an evaluation for the rule. Valid values: `ConfigurationItemChangeNotification`, `OversizedConfigurationItemChangeNotification`, and `ScheduledNotification`
     * 
     */
    public Output<List<String>> triggerTypes() {
        return this.triggerTypes;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public OrganizationCustomRule(String name) {
        this(name, OrganizationCustomRuleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public OrganizationCustomRule(String name, OrganizationCustomRuleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public OrganizationCustomRule(String name, OrganizationCustomRuleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:cfg/organizationCustomRule:OrganizationCustomRule", name, args == null ? OrganizationCustomRuleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private OrganizationCustomRule(String name, Output<String> id, @Nullable OrganizationCustomRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:cfg/organizationCustomRule:OrganizationCustomRule", name, state, makeResourceOptions(options, id));
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
    public static OrganizationCustomRule get(String name, Output<String> id, @Nullable OrganizationCustomRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new OrganizationCustomRule(name, id, state, options);
    }
}
