// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.cloudformation;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.cloudformation.StackSetInstanceArgs;
import com.pulumi.aws.cloudformation.inputs.StackSetInstanceState;
import com.pulumi.aws.cloudformation.outputs.StackSetInstanceDeploymentTargets;
import com.pulumi.aws.cloudformation.outputs.StackSetInstanceOperationPreferences;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a CloudFormation StackSet Instance. Instances are managed in the account and region of the StackSet after the target account permissions have been configured. Additional information about StackSets can be found in the [AWS CloudFormation User Guide](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/what-is-cfnstacksets.html).
 * 
 * &gt; **NOTE:** All target accounts must have an IAM Role created that matches the name of the execution role configured in the StackSet (the `execution_role_name` argument in the `aws.cloudformation.StackSet` resource) in a trust relationship with the administrative account or administration IAM Role. The execution role must have appropriate permissions to manage resources defined in the template along with those required for StackSets to operate. See the [AWS CloudFormation User Guide](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs.html) for more details.
 * 
 * &gt; **NOTE:** To retain the Stack during resource destroy, ensure `retain_stack` has been set to `true` in the state first. This must be completed _before_ a deployment that would destroy the resource.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.cloudformation.StackSetInstance;
 * import com.pulumi.aws.cloudformation.StackSetInstanceArgs;
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
 *         var example = new StackSetInstance(&#34;example&#34;, StackSetInstanceArgs.builder()        
 *             .accountId(&#34;123456789012&#34;)
 *             .region(&#34;us-east-1&#34;)
 *             .stackSetName(aws_cloudformation_stack_set.example().name())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Example IAM Setup in Target Account
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.iam.IamFunctions;
 * import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
 * import com.pulumi.aws.iam.Role;
 * import com.pulumi.aws.iam.RoleArgs;
 * import com.pulumi.aws.iam.RolePolicy;
 * import com.pulumi.aws.iam.RolePolicyArgs;
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
 *         final var aWSCloudFormationStackSetExecutionRoleAssumeRolePolicy = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
 *             .statements(GetPolicyDocumentStatementArgs.builder()
 *                 .actions(&#34;sts:AssumeRole&#34;)
 *                 .effect(&#34;Allow&#34;)
 *                 .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
 *                     .identifiers(aws_iam_role.AWSCloudFormationStackSetAdministrationRole().arn())
 *                     .type(&#34;AWS&#34;)
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *         var aWSCloudFormationStackSetExecutionRole = new Role(&#34;aWSCloudFormationStackSetExecutionRole&#34;, RoleArgs.builder()        
 *             .assumeRolePolicy(aWSCloudFormationStackSetExecutionRoleAssumeRolePolicy.applyValue(getPolicyDocumentResult -&gt; getPolicyDocumentResult.json()))
 *             .build());
 * 
 *         final var aWSCloudFormationStackSetExecutionRoleMinimumExecutionPolicyPolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
 *             .statements(GetPolicyDocumentStatementArgs.builder()
 *                 .actions(                
 *                     &#34;cloudformation:*&#34;,
 *                     &#34;s3:*&#34;,
 *                     &#34;sns:*&#34;)
 *                 .effect(&#34;Allow&#34;)
 *                 .resources(&#34;*&#34;)
 *                 .build())
 *             .build());
 * 
 *         var aWSCloudFormationStackSetExecutionRoleMinimumExecutionPolicyRolePolicy = new RolePolicy(&#34;aWSCloudFormationStackSetExecutionRoleMinimumExecutionPolicyRolePolicy&#34;, RolePolicyArgs.builder()        
 *             .policy(aWSCloudFormationStackSetExecutionRoleMinimumExecutionPolicyPolicyDocument.applyValue(getPolicyDocumentResult -&gt; getPolicyDocumentResult.json()))
 *             .role(aWSCloudFormationStackSetExecutionRole.name())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Example Deployment across Organizations account
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.cloudformation.StackSetInstance;
 * import com.pulumi.aws.cloudformation.StackSetInstanceArgs;
 * import com.pulumi.aws.cloudformation.inputs.StackSetInstanceDeploymentTargetsArgs;
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
 *         var example = new StackSetInstance(&#34;example&#34;, StackSetInstanceArgs.builder()        
 *             .deploymentTargets(StackSetInstanceDeploymentTargetsArgs.builder()
 *                 .organizationalUnitIds(aws_organizations_organization.example().roots()[0].id())
 *                 .build())
 *             .region(&#34;us-east-1&#34;)
 *             .stackSetName(aws_cloudformation_stack_set.example().name())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * CloudFormation StackSet Instances that target an AWS Account ID can be imported using the StackSet name, target AWS account ID, and target AWS region separated by commas (`,`) e.g.
 * 
 * ```sh
 *  $ pulumi import aws:cloudformation/stackSetInstance:StackSetInstance example example,123456789012,us-east-1
 * ```
 * 
 *  CloudFormation StackSet Instances that target AWS Organizational Units can be imported using the StackSet name, a slash (`/`) separated list of organizational unit IDs, and target AWS region separated by commas (`,`) e.g.
 * 
 * ```sh
 *  $ pulumi import aws:cloudformation/stackSetInstance:StackSetInstance example example,ou-sdas-123123123/ou-sdas-789789789,us-east-1
 * ```
 * 
 */
@ResourceType(type="aws:cloudformation/stackSetInstance:StackSetInstance")
public class StackSetInstance extends com.pulumi.resources.CustomResource {
    /**
     * Target AWS Account ID to create a Stack based on the StackSet. Defaults to current account.
     * 
     */
    @Export(name="accountId", refs={String.class}, tree="[0]")
    private Output<String> accountId;

    /**
     * @return Target AWS Account ID to create a Stack based on the StackSet. Defaults to current account.
     * 
     */
    public Output<String> accountId() {
        return this.accountId;
    }
    /**
     * Specifies whether you are acting as an account administrator in the organization&#39;s management account or as a delegated administrator in a member account. Valid values: `SELF` (default), `DELEGATED_ADMIN`.
     * 
     */
    @Export(name="callAs", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> callAs;

    /**
     * @return Specifies whether you are acting as an account administrator in the organization&#39;s management account or as a delegated administrator in a member account. Valid values: `SELF` (default), `DELEGATED_ADMIN`.
     * 
     */
    public Output<Optional<String>> callAs() {
        return Codegen.optional(this.callAs);
    }
    /**
     * The AWS Organizations accounts to which StackSets deploys. StackSets doesn&#39;t deploy stack instances to the organization management account, even if the organization management account is in your organization or in an OU in your organization. Drift detection is not possible for this argument. See deployment_targets below.
     * 
     */
    @Export(name="deploymentTargets", refs={StackSetInstanceDeploymentTargets.class}, tree="[0]")
    private Output</* @Nullable */ StackSetInstanceDeploymentTargets> deploymentTargets;

    /**
     * @return The AWS Organizations accounts to which StackSets deploys. StackSets doesn&#39;t deploy stack instances to the organization management account, even if the organization management account is in your organization or in an OU in your organization. Drift detection is not possible for this argument. See deployment_targets below.
     * 
     */
    public Output<Optional<StackSetInstanceDeploymentTargets>> deploymentTargets() {
        return Codegen.optional(this.deploymentTargets);
    }
    /**
     * Preferences for how AWS CloudFormation performs a stack set operation.
     * 
     */
    @Export(name="operationPreferences", refs={StackSetInstanceOperationPreferences.class}, tree="[0]")
    private Output</* @Nullable */ StackSetInstanceOperationPreferences> operationPreferences;

    /**
     * @return Preferences for how AWS CloudFormation performs a stack set operation.
     * 
     */
    public Output<Optional<StackSetInstanceOperationPreferences>> operationPreferences() {
        return Codegen.optional(this.operationPreferences);
    }
    /**
     * The organization root ID or organizational unit (OU) IDs specified for `deployment_targets`.
     * 
     */
    @Export(name="organizationalUnitId", refs={String.class}, tree="[0]")
    private Output<String> organizationalUnitId;

    /**
     * @return The organization root ID or organizational unit (OU) IDs specified for `deployment_targets`.
     * 
     */
    public Output<String> organizationalUnitId() {
        return this.organizationalUnitId;
    }
    /**
     * Key-value map of input parameters to override from the StackSet for this Instance.
     * 
     */
    @Export(name="parameterOverrides", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> parameterOverrides;

    /**
     * @return Key-value map of input parameters to override from the StackSet for this Instance.
     * 
     */
    public Output<Optional<Map<String,String>>> parameterOverrides() {
        return Codegen.optional(this.parameterOverrides);
    }
    /**
     * Target AWS Region to create a Stack based on the StackSet. Defaults to current region.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return Target AWS Region to create a Stack based on the StackSet. Defaults to current region.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * During resource destroy, remove Instance from StackSet while keeping the Stack and its associated resources. Must be enabled in the state _before_ destroy operation to take effect. You cannot reassociate a retained Stack or add an existing, saved Stack to a new StackSet. Defaults to `false`.
     * 
     */
    @Export(name="retainStack", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> retainStack;

    /**
     * @return During resource destroy, remove Instance from StackSet while keeping the Stack and its associated resources. Must be enabled in the state _before_ destroy operation to take effect. You cannot reassociate a retained Stack or add an existing, saved Stack to a new StackSet. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> retainStack() {
        return Codegen.optional(this.retainStack);
    }
    /**
     * Stack identifier
     * 
     */
    @Export(name="stackId", refs={String.class}, tree="[0]")
    private Output<String> stackId;

    /**
     * @return Stack identifier
     * 
     */
    public Output<String> stackId() {
        return this.stackId;
    }
    /**
     * Name of the StackSet.
     * 
     */
    @Export(name="stackSetName", refs={String.class}, tree="[0]")
    private Output<String> stackSetName;

    /**
     * @return Name of the StackSet.
     * 
     */
    public Output<String> stackSetName() {
        return this.stackSetName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public StackSetInstance(String name) {
        this(name, StackSetInstanceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public StackSetInstance(String name, StackSetInstanceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public StackSetInstance(String name, StackSetInstanceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:cloudformation/stackSetInstance:StackSetInstance", name, args == null ? StackSetInstanceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private StackSetInstance(String name, Output<String> id, @Nullable StackSetInstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:cloudformation/stackSetInstance:StackSetInstance", name, state, makeResourceOptions(options, id));
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
    public static StackSetInstance get(String name, Output<String> id, @Nullable StackSetInstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new StackSetInstance(name, id, state, options);
    }
}
