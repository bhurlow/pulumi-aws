// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.cloudformation.inputs;

import com.pulumi.aws.cloudformation.inputs.StackSetAutoDeploymentArgs;
import com.pulumi.aws.cloudformation.inputs.StackSetManagedExecutionArgs;
import com.pulumi.aws.cloudformation.inputs.StackSetOperationPreferencesArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class StackSetState extends com.pulumi.resources.ResourceArgs {

    public static final StackSetState Empty = new StackSetState();

    /**
     * Amazon Resource Number (ARN) of the IAM Role in the administrator account. This must be defined when using the `SELF_MANAGED` permission model.
     * 
     */
    @Import(name="administrationRoleArn")
    private @Nullable Output<String> administrationRoleArn;

    /**
     * @return Amazon Resource Number (ARN) of the IAM Role in the administrator account. This must be defined when using the `SELF_MANAGED` permission model.
     * 
     */
    public Optional<Output<String>> administrationRoleArn() {
        return Optional.ofNullable(this.administrationRoleArn);
    }

    /**
     * Amazon Resource Name (ARN) of the StackSet.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return Amazon Resource Name (ARN) of the StackSet.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * Configuration block containing the auto-deployment model for your StackSet. This can only be defined when using the `SERVICE_MANAGED` permission model.
     * 
     */
    @Import(name="autoDeployment")
    private @Nullable Output<StackSetAutoDeploymentArgs> autoDeployment;

    /**
     * @return Configuration block containing the auto-deployment model for your StackSet. This can only be defined when using the `SERVICE_MANAGED` permission model.
     * 
     */
    public Optional<Output<StackSetAutoDeploymentArgs>> autoDeployment() {
        return Optional.ofNullable(this.autoDeployment);
    }

    /**
     * Specifies whether you are acting as an account administrator in the organization&#39;s management account or as a delegated administrator in a member account. Valid values: `SELF` (default), `DELEGATED_ADMIN`.
     * 
     */
    @Import(name="callAs")
    private @Nullable Output<String> callAs;

    /**
     * @return Specifies whether you are acting as an account administrator in the organization&#39;s management account or as a delegated administrator in a member account. Valid values: `SELF` (default), `DELEGATED_ADMIN`.
     * 
     */
    public Optional<Output<String>> callAs() {
        return Optional.ofNullable(this.callAs);
    }

    /**
     * A list of capabilities. Valid values: `CAPABILITY_IAM`, `CAPABILITY_NAMED_IAM`, `CAPABILITY_AUTO_EXPAND`.
     * 
     */
    @Import(name="capabilities")
    private @Nullable Output<List<String>> capabilities;

    /**
     * @return A list of capabilities. Valid values: `CAPABILITY_IAM`, `CAPABILITY_NAMED_IAM`, `CAPABILITY_AUTO_EXPAND`.
     * 
     */
    public Optional<Output<List<String>>> capabilities() {
        return Optional.ofNullable(this.capabilities);
    }

    /**
     * Description of the StackSet.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description of the StackSet.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Name of the IAM Role in all target accounts for StackSet operations. Defaults to `AWSCloudFormationStackSetExecutionRole` when using the `SELF_MANAGED` permission model. This should not be defined when using the `SERVICE_MANAGED` permission model.
     * 
     */
    @Import(name="executionRoleName")
    private @Nullable Output<String> executionRoleName;

    /**
     * @return Name of the IAM Role in all target accounts for StackSet operations. Defaults to `AWSCloudFormationStackSetExecutionRole` when using the `SELF_MANAGED` permission model. This should not be defined when using the `SERVICE_MANAGED` permission model.
     * 
     */
    public Optional<Output<String>> executionRoleName() {
        return Optional.ofNullable(this.executionRoleName);
    }

    /**
     * Configuration block to allow StackSets to perform non-conflicting operations concurrently and queues conflicting operations.
     * 
     */
    @Import(name="managedExecution")
    private @Nullable Output<StackSetManagedExecutionArgs> managedExecution;

    /**
     * @return Configuration block to allow StackSets to perform non-conflicting operations concurrently and queues conflicting operations.
     * 
     */
    public Optional<Output<StackSetManagedExecutionArgs>> managedExecution() {
        return Optional.ofNullable(this.managedExecution);
    }

    /**
     * Name of the StackSet. The name must be unique in the region where you create your StackSet. The name can contain only alphanumeric characters (case-sensitive) and hyphens. It must start with an alphabetic character and cannot be longer than 128 characters.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the StackSet. The name must be unique in the region where you create your StackSet. The name can contain only alphanumeric characters (case-sensitive) and hyphens. It must start with an alphabetic character and cannot be longer than 128 characters.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Preferences for how AWS CloudFormation performs a stack set update.
     * 
     */
    @Import(name="operationPreferences")
    private @Nullable Output<StackSetOperationPreferencesArgs> operationPreferences;

    /**
     * @return Preferences for how AWS CloudFormation performs a stack set update.
     * 
     */
    public Optional<Output<StackSetOperationPreferencesArgs>> operationPreferences() {
        return Optional.ofNullable(this.operationPreferences);
    }

    /**
     * Key-value map of input parameters for the StackSet template. All template parameters, including those with a `Default`, must be configured or ignored with `lifecycle` configuration block `ignore_changes` argument. All `NoEcho` template parameters must be ignored with the `lifecycle` configuration block `ignore_changes` argument.
     * 
     */
    @Import(name="parameters")
    private @Nullable Output<Map<String,String>> parameters;

    /**
     * @return Key-value map of input parameters for the StackSet template. All template parameters, including those with a `Default`, must be configured or ignored with `lifecycle` configuration block `ignore_changes` argument. All `NoEcho` template parameters must be ignored with the `lifecycle` configuration block `ignore_changes` argument.
     * 
     */
    public Optional<Output<Map<String,String>>> parameters() {
        return Optional.ofNullable(this.parameters);
    }

    /**
     * Describes how the IAM roles required for your StackSet are created. Valid values: `SELF_MANAGED` (default), `SERVICE_MANAGED`.
     * 
     */
    @Import(name="permissionModel")
    private @Nullable Output<String> permissionModel;

    /**
     * @return Describes how the IAM roles required for your StackSet are created. Valid values: `SELF_MANAGED` (default), `SERVICE_MANAGED`.
     * 
     */
    public Optional<Output<String>> permissionModel() {
        return Optional.ofNullable(this.permissionModel);
    }

    /**
     * Unique identifier of the StackSet.
     * 
     */
    @Import(name="stackSetId")
    private @Nullable Output<String> stackSetId;

    /**
     * @return Unique identifier of the StackSet.
     * 
     */
    public Optional<Output<String>> stackSetId() {
        return Optional.ofNullable(this.stackSetId);
    }

    /**
     * Key-value map of tags to associate with this StackSet and the Stacks created from it. AWS CloudFormation also propagates these tags to supported resources that are created in the Stacks. A maximum number of 50 tags can be specified. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Key-value map of tags to associate with this StackSet and the Stacks created from it. AWS CloudFormation also propagates these tags to supported resources that are created in the Stacks. A maximum number of 50 tags can be specified. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    @Import(name="tagsAll")
    private @Nullable Output<Map<String,String>> tagsAll;

    /**
     * @return A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    public Optional<Output<Map<String,String>>> tagsAll() {
        return Optional.ofNullable(this.tagsAll);
    }

    /**
     * String containing the CloudFormation template body. Maximum size: 51,200 bytes. Conflicts with `template_url`.
     * 
     */
    @Import(name="templateBody")
    private @Nullable Output<String> templateBody;

    /**
     * @return String containing the CloudFormation template body. Maximum size: 51,200 bytes. Conflicts with `template_url`.
     * 
     */
    public Optional<Output<String>> templateBody() {
        return Optional.ofNullable(this.templateBody);
    }

    /**
     * String containing the location of a file containing the CloudFormation template body. The URL must point to a template that is located in an Amazon S3 bucket. Maximum location file size: 460,800 bytes. Conflicts with `template_body`.
     * 
     */
    @Import(name="templateUrl")
    private @Nullable Output<String> templateUrl;

    /**
     * @return String containing the location of a file containing the CloudFormation template body. The URL must point to a template that is located in an Amazon S3 bucket. Maximum location file size: 460,800 bytes. Conflicts with `template_body`.
     * 
     */
    public Optional<Output<String>> templateUrl() {
        return Optional.ofNullable(this.templateUrl);
    }

    private StackSetState() {}

    private StackSetState(StackSetState $) {
        this.administrationRoleArn = $.administrationRoleArn;
        this.arn = $.arn;
        this.autoDeployment = $.autoDeployment;
        this.callAs = $.callAs;
        this.capabilities = $.capabilities;
        this.description = $.description;
        this.executionRoleName = $.executionRoleName;
        this.managedExecution = $.managedExecution;
        this.name = $.name;
        this.operationPreferences = $.operationPreferences;
        this.parameters = $.parameters;
        this.permissionModel = $.permissionModel;
        this.stackSetId = $.stackSetId;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
        this.templateBody = $.templateBody;
        this.templateUrl = $.templateUrl;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(StackSetState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private StackSetState $;

        public Builder() {
            $ = new StackSetState();
        }

        public Builder(StackSetState defaults) {
            $ = new StackSetState(Objects.requireNonNull(defaults));
        }

        /**
         * @param administrationRoleArn Amazon Resource Number (ARN) of the IAM Role in the administrator account. This must be defined when using the `SELF_MANAGED` permission model.
         * 
         * @return builder
         * 
         */
        public Builder administrationRoleArn(@Nullable Output<String> administrationRoleArn) {
            $.administrationRoleArn = administrationRoleArn;
            return this;
        }

        /**
         * @param administrationRoleArn Amazon Resource Number (ARN) of the IAM Role in the administrator account. This must be defined when using the `SELF_MANAGED` permission model.
         * 
         * @return builder
         * 
         */
        public Builder administrationRoleArn(String administrationRoleArn) {
            return administrationRoleArn(Output.of(administrationRoleArn));
        }

        /**
         * @param arn Amazon Resource Name (ARN) of the StackSet.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn Amazon Resource Name (ARN) of the StackSet.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param autoDeployment Configuration block containing the auto-deployment model for your StackSet. This can only be defined when using the `SERVICE_MANAGED` permission model.
         * 
         * @return builder
         * 
         */
        public Builder autoDeployment(@Nullable Output<StackSetAutoDeploymentArgs> autoDeployment) {
            $.autoDeployment = autoDeployment;
            return this;
        }

        /**
         * @param autoDeployment Configuration block containing the auto-deployment model for your StackSet. This can only be defined when using the `SERVICE_MANAGED` permission model.
         * 
         * @return builder
         * 
         */
        public Builder autoDeployment(StackSetAutoDeploymentArgs autoDeployment) {
            return autoDeployment(Output.of(autoDeployment));
        }

        /**
         * @param callAs Specifies whether you are acting as an account administrator in the organization&#39;s management account or as a delegated administrator in a member account. Valid values: `SELF` (default), `DELEGATED_ADMIN`.
         * 
         * @return builder
         * 
         */
        public Builder callAs(@Nullable Output<String> callAs) {
            $.callAs = callAs;
            return this;
        }

        /**
         * @param callAs Specifies whether you are acting as an account administrator in the organization&#39;s management account or as a delegated administrator in a member account. Valid values: `SELF` (default), `DELEGATED_ADMIN`.
         * 
         * @return builder
         * 
         */
        public Builder callAs(String callAs) {
            return callAs(Output.of(callAs));
        }

        /**
         * @param capabilities A list of capabilities. Valid values: `CAPABILITY_IAM`, `CAPABILITY_NAMED_IAM`, `CAPABILITY_AUTO_EXPAND`.
         * 
         * @return builder
         * 
         */
        public Builder capabilities(@Nullable Output<List<String>> capabilities) {
            $.capabilities = capabilities;
            return this;
        }

        /**
         * @param capabilities A list of capabilities. Valid values: `CAPABILITY_IAM`, `CAPABILITY_NAMED_IAM`, `CAPABILITY_AUTO_EXPAND`.
         * 
         * @return builder
         * 
         */
        public Builder capabilities(List<String> capabilities) {
            return capabilities(Output.of(capabilities));
        }

        /**
         * @param capabilities A list of capabilities. Valid values: `CAPABILITY_IAM`, `CAPABILITY_NAMED_IAM`, `CAPABILITY_AUTO_EXPAND`.
         * 
         * @return builder
         * 
         */
        public Builder capabilities(String... capabilities) {
            return capabilities(List.of(capabilities));
        }

        /**
         * @param description Description of the StackSet.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description of the StackSet.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param executionRoleName Name of the IAM Role in all target accounts for StackSet operations. Defaults to `AWSCloudFormationStackSetExecutionRole` when using the `SELF_MANAGED` permission model. This should not be defined when using the `SERVICE_MANAGED` permission model.
         * 
         * @return builder
         * 
         */
        public Builder executionRoleName(@Nullable Output<String> executionRoleName) {
            $.executionRoleName = executionRoleName;
            return this;
        }

        /**
         * @param executionRoleName Name of the IAM Role in all target accounts for StackSet operations. Defaults to `AWSCloudFormationStackSetExecutionRole` when using the `SELF_MANAGED` permission model. This should not be defined when using the `SERVICE_MANAGED` permission model.
         * 
         * @return builder
         * 
         */
        public Builder executionRoleName(String executionRoleName) {
            return executionRoleName(Output.of(executionRoleName));
        }

        /**
         * @param managedExecution Configuration block to allow StackSets to perform non-conflicting operations concurrently and queues conflicting operations.
         * 
         * @return builder
         * 
         */
        public Builder managedExecution(@Nullable Output<StackSetManagedExecutionArgs> managedExecution) {
            $.managedExecution = managedExecution;
            return this;
        }

        /**
         * @param managedExecution Configuration block to allow StackSets to perform non-conflicting operations concurrently and queues conflicting operations.
         * 
         * @return builder
         * 
         */
        public Builder managedExecution(StackSetManagedExecutionArgs managedExecution) {
            return managedExecution(Output.of(managedExecution));
        }

        /**
         * @param name Name of the StackSet. The name must be unique in the region where you create your StackSet. The name can contain only alphanumeric characters (case-sensitive) and hyphens. It must start with an alphabetic character and cannot be longer than 128 characters.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the StackSet. The name must be unique in the region where you create your StackSet. The name can contain only alphanumeric characters (case-sensitive) and hyphens. It must start with an alphabetic character and cannot be longer than 128 characters.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param operationPreferences Preferences for how AWS CloudFormation performs a stack set update.
         * 
         * @return builder
         * 
         */
        public Builder operationPreferences(@Nullable Output<StackSetOperationPreferencesArgs> operationPreferences) {
            $.operationPreferences = operationPreferences;
            return this;
        }

        /**
         * @param operationPreferences Preferences for how AWS CloudFormation performs a stack set update.
         * 
         * @return builder
         * 
         */
        public Builder operationPreferences(StackSetOperationPreferencesArgs operationPreferences) {
            return operationPreferences(Output.of(operationPreferences));
        }

        /**
         * @param parameters Key-value map of input parameters for the StackSet template. All template parameters, including those with a `Default`, must be configured or ignored with `lifecycle` configuration block `ignore_changes` argument. All `NoEcho` template parameters must be ignored with the `lifecycle` configuration block `ignore_changes` argument.
         * 
         * @return builder
         * 
         */
        public Builder parameters(@Nullable Output<Map<String,String>> parameters) {
            $.parameters = parameters;
            return this;
        }

        /**
         * @param parameters Key-value map of input parameters for the StackSet template. All template parameters, including those with a `Default`, must be configured or ignored with `lifecycle` configuration block `ignore_changes` argument. All `NoEcho` template parameters must be ignored with the `lifecycle` configuration block `ignore_changes` argument.
         * 
         * @return builder
         * 
         */
        public Builder parameters(Map<String,String> parameters) {
            return parameters(Output.of(parameters));
        }

        /**
         * @param permissionModel Describes how the IAM roles required for your StackSet are created. Valid values: `SELF_MANAGED` (default), `SERVICE_MANAGED`.
         * 
         * @return builder
         * 
         */
        public Builder permissionModel(@Nullable Output<String> permissionModel) {
            $.permissionModel = permissionModel;
            return this;
        }

        /**
         * @param permissionModel Describes how the IAM roles required for your StackSet are created. Valid values: `SELF_MANAGED` (default), `SERVICE_MANAGED`.
         * 
         * @return builder
         * 
         */
        public Builder permissionModel(String permissionModel) {
            return permissionModel(Output.of(permissionModel));
        }

        /**
         * @param stackSetId Unique identifier of the StackSet.
         * 
         * @return builder
         * 
         */
        public Builder stackSetId(@Nullable Output<String> stackSetId) {
            $.stackSetId = stackSetId;
            return this;
        }

        /**
         * @param stackSetId Unique identifier of the StackSet.
         * 
         * @return builder
         * 
         */
        public Builder stackSetId(String stackSetId) {
            return stackSetId(Output.of(stackSetId));
        }

        /**
         * @param tags Key-value map of tags to associate with this StackSet and the Stacks created from it. AWS CloudFormation also propagates these tags to supported resources that are created in the Stacks. A maximum number of 50 tags can be specified. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Key-value map of tags to associate with this StackSet and the Stacks created from it. AWS CloudFormation also propagates these tags to supported resources that are created in the Stacks. A maximum number of 50 tags can be specified. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tagsAll A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
         * 
         * @return builder
         * 
         * @deprecated
         * Please use `tags` instead.
         * 
         */
        @Deprecated /* Please use `tags` instead. */
        public Builder tagsAll(@Nullable Output<Map<String,String>> tagsAll) {
            $.tagsAll = tagsAll;
            return this;
        }

        /**
         * @param tagsAll A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
         * 
         * @return builder
         * 
         * @deprecated
         * Please use `tags` instead.
         * 
         */
        @Deprecated /* Please use `tags` instead. */
        public Builder tagsAll(Map<String,String> tagsAll) {
            return tagsAll(Output.of(tagsAll));
        }

        /**
         * @param templateBody String containing the CloudFormation template body. Maximum size: 51,200 bytes. Conflicts with `template_url`.
         * 
         * @return builder
         * 
         */
        public Builder templateBody(@Nullable Output<String> templateBody) {
            $.templateBody = templateBody;
            return this;
        }

        /**
         * @param templateBody String containing the CloudFormation template body. Maximum size: 51,200 bytes. Conflicts with `template_url`.
         * 
         * @return builder
         * 
         */
        public Builder templateBody(String templateBody) {
            return templateBody(Output.of(templateBody));
        }

        /**
         * @param templateUrl String containing the location of a file containing the CloudFormation template body. The URL must point to a template that is located in an Amazon S3 bucket. Maximum location file size: 460,800 bytes. Conflicts with `template_body`.
         * 
         * @return builder
         * 
         */
        public Builder templateUrl(@Nullable Output<String> templateUrl) {
            $.templateUrl = templateUrl;
            return this;
        }

        /**
         * @param templateUrl String containing the location of a file containing the CloudFormation template body. The URL must point to a template that is located in an Amazon S3 bucket. Maximum location file size: 460,800 bytes. Conflicts with `template_body`.
         * 
         * @return builder
         * 
         */
        public Builder templateUrl(String templateUrl) {
            return templateUrl(Output.of(templateUrl));
        }

        public StackSetState build() {
            return $;
        }
    }

}
