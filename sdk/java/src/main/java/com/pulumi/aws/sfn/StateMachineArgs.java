// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.sfn;

import com.pulumi.aws.sfn.inputs.StateMachineLoggingConfigurationArgs;
import com.pulumi.aws.sfn.inputs.StateMachineTracingConfigurationArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class StateMachineArgs extends com.pulumi.resources.ResourceArgs {

    public static final StateMachineArgs Empty = new StateMachineArgs();

    /**
     * The [Amazon States Language](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-amazon-states-language.html) definition of the state machine.
     * 
     */
    @Import(name="definition", required=true)
    private Output<String> definition;

    /**
     * @return The [Amazon States Language](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-amazon-states-language.html) definition of the state machine.
     * 
     */
    public Output<String> definition() {
        return this.definition;
    }

    /**
     * Defines what execution history events are logged and where they are logged. The `logging_configuration` parameter is only valid when `type` is set to `EXPRESS`. Defaults to `OFF`. For more information see [Logging Express Workflows](https://docs.aws.amazon.com/step-functions/latest/dg/cw-logs.html) and [Log Levels](https://docs.aws.amazon.com/step-functions/latest/dg/cloudwatch-log-level.html) in the AWS Step Functions User Guide.
     * 
     */
    @Import(name="loggingConfiguration")
    private @Nullable Output<StateMachineLoggingConfigurationArgs> loggingConfiguration;

    /**
     * @return Defines what execution history events are logged and where they are logged. The `logging_configuration` parameter is only valid when `type` is set to `EXPRESS`. Defaults to `OFF`. For more information see [Logging Express Workflows](https://docs.aws.amazon.com/step-functions/latest/dg/cw-logs.html) and [Log Levels](https://docs.aws.amazon.com/step-functions/latest/dg/cloudwatch-log-level.html) in the AWS Step Functions User Guide.
     * 
     */
    public Optional<Output<StateMachineLoggingConfigurationArgs>> loggingConfiguration() {
        return Optional.ofNullable(this.loggingConfiguration);
    }

    /**
     * The name of the state machine. The name should only contain `0`-`9`, `A`-`Z`, `a`-`z`, `-` and `_`. If omitted, the provider will assign a random, unique name.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the state machine. The name should only contain `0`-`9`, `A`-`Z`, `a`-`z`, `-` and `_`. If omitted, the provider will assign a random, unique name.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     * 
     */
    @Import(name="namePrefix")
    private @Nullable Output<String> namePrefix;

    /**
     * @return Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     * 
     */
    public Optional<Output<String>> namePrefix() {
        return Optional.ofNullable(this.namePrefix);
    }

    /**
     * Set to true to publish a version of the state machine during creation. Default: false.
     * 
     */
    @Import(name="publish")
    private @Nullable Output<Boolean> publish;

    /**
     * @return Set to true to publish a version of the state machine during creation. Default: false.
     * 
     */
    public Optional<Output<Boolean>> publish() {
        return Optional.ofNullable(this.publish);
    }

    /**
     * The Amazon Resource Name (ARN) of the IAM role to use for this state machine.
     * 
     */
    @Import(name="roleArn", required=true)
    private Output<String> roleArn;

    /**
     * @return The Amazon Resource Name (ARN) of the IAM role to use for this state machine.
     * 
     */
    public Output<String> roleArn() {
        return this.roleArn;
    }

    /**
     * Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * Selects whether AWS X-Ray tracing is enabled.
     * 
     */
    @Import(name="tracingConfiguration")
    private @Nullable Output<StateMachineTracingConfigurationArgs> tracingConfiguration;

    /**
     * @return Selects whether AWS X-Ray tracing is enabled.
     * 
     */
    public Optional<Output<StateMachineTracingConfigurationArgs>> tracingConfiguration() {
        return Optional.ofNullable(this.tracingConfiguration);
    }

    /**
     * Determines whether a Standard or Express state machine is created. The default is `STANDARD`. You cannot update the type of a state machine once it has been created. Valid values: `STANDARD`, `EXPRESS`.
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return Determines whether a Standard or Express state machine is created. The default is `STANDARD`. You cannot update the type of a state machine once it has been created. Valid values: `STANDARD`, `EXPRESS`.
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    private StateMachineArgs() {}

    private StateMachineArgs(StateMachineArgs $) {
        this.definition = $.definition;
        this.loggingConfiguration = $.loggingConfiguration;
        this.name = $.name;
        this.namePrefix = $.namePrefix;
        this.publish = $.publish;
        this.roleArn = $.roleArn;
        this.tags = $.tags;
        this.tracingConfiguration = $.tracingConfiguration;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(StateMachineArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private StateMachineArgs $;

        public Builder() {
            $ = new StateMachineArgs();
        }

        public Builder(StateMachineArgs defaults) {
            $ = new StateMachineArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param definition The [Amazon States Language](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-amazon-states-language.html) definition of the state machine.
         * 
         * @return builder
         * 
         */
        public Builder definition(Output<String> definition) {
            $.definition = definition;
            return this;
        }

        /**
         * @param definition The [Amazon States Language](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-amazon-states-language.html) definition of the state machine.
         * 
         * @return builder
         * 
         */
        public Builder definition(String definition) {
            return definition(Output.of(definition));
        }

        /**
         * @param loggingConfiguration Defines what execution history events are logged and where they are logged. The `logging_configuration` parameter is only valid when `type` is set to `EXPRESS`. Defaults to `OFF`. For more information see [Logging Express Workflows](https://docs.aws.amazon.com/step-functions/latest/dg/cw-logs.html) and [Log Levels](https://docs.aws.amazon.com/step-functions/latest/dg/cloudwatch-log-level.html) in the AWS Step Functions User Guide.
         * 
         * @return builder
         * 
         */
        public Builder loggingConfiguration(@Nullable Output<StateMachineLoggingConfigurationArgs> loggingConfiguration) {
            $.loggingConfiguration = loggingConfiguration;
            return this;
        }

        /**
         * @param loggingConfiguration Defines what execution history events are logged and where they are logged. The `logging_configuration` parameter is only valid when `type` is set to `EXPRESS`. Defaults to `OFF`. For more information see [Logging Express Workflows](https://docs.aws.amazon.com/step-functions/latest/dg/cw-logs.html) and [Log Levels](https://docs.aws.amazon.com/step-functions/latest/dg/cloudwatch-log-level.html) in the AWS Step Functions User Guide.
         * 
         * @return builder
         * 
         */
        public Builder loggingConfiguration(StateMachineLoggingConfigurationArgs loggingConfiguration) {
            return loggingConfiguration(Output.of(loggingConfiguration));
        }

        /**
         * @param name The name of the state machine. The name should only contain `0`-`9`, `A`-`Z`, `a`-`z`, `-` and `_`. If omitted, the provider will assign a random, unique name.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the state machine. The name should only contain `0`-`9`, `A`-`Z`, `a`-`z`, `-` and `_`. If omitted, the provider will assign a random, unique name.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param namePrefix Creates a unique name beginning with the specified prefix. Conflicts with `name`.
         * 
         * @return builder
         * 
         */
        public Builder namePrefix(@Nullable Output<String> namePrefix) {
            $.namePrefix = namePrefix;
            return this;
        }

        /**
         * @param namePrefix Creates a unique name beginning with the specified prefix. Conflicts with `name`.
         * 
         * @return builder
         * 
         */
        public Builder namePrefix(String namePrefix) {
            return namePrefix(Output.of(namePrefix));
        }

        /**
         * @param publish Set to true to publish a version of the state machine during creation. Default: false.
         * 
         * @return builder
         * 
         */
        public Builder publish(@Nullable Output<Boolean> publish) {
            $.publish = publish;
            return this;
        }

        /**
         * @param publish Set to true to publish a version of the state machine during creation. Default: false.
         * 
         * @return builder
         * 
         */
        public Builder publish(Boolean publish) {
            return publish(Output.of(publish));
        }

        /**
         * @param roleArn The Amazon Resource Name (ARN) of the IAM role to use for this state machine.
         * 
         * @return builder
         * 
         */
        public Builder roleArn(Output<String> roleArn) {
            $.roleArn = roleArn;
            return this;
        }

        /**
         * @param roleArn The Amazon Resource Name (ARN) of the IAM role to use for this state machine.
         * 
         * @return builder
         * 
         */
        public Builder roleArn(String roleArn) {
            return roleArn(Output.of(roleArn));
        }

        /**
         * @param tags Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tracingConfiguration Selects whether AWS X-Ray tracing is enabled.
         * 
         * @return builder
         * 
         */
        public Builder tracingConfiguration(@Nullable Output<StateMachineTracingConfigurationArgs> tracingConfiguration) {
            $.tracingConfiguration = tracingConfiguration;
            return this;
        }

        /**
         * @param tracingConfiguration Selects whether AWS X-Ray tracing is enabled.
         * 
         * @return builder
         * 
         */
        public Builder tracingConfiguration(StateMachineTracingConfigurationArgs tracingConfiguration) {
            return tracingConfiguration(Output.of(tracingConfiguration));
        }

        /**
         * @param type Determines whether a Standard or Express state machine is created. The default is `STANDARD`. You cannot update the type of a state machine once it has been created. Valid values: `STANDARD`, `EXPRESS`.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Determines whether a Standard or Express state machine is created. The default is `STANDARD`. You cannot update the type of a state machine once it has been created. Valid values: `STANDARD`, `EXPRESS`.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public StateMachineArgs build() {
            $.definition = Objects.requireNonNull($.definition, "expected parameter 'definition' to be non-null");
            $.roleArn = Objects.requireNonNull($.roleArn, "expected parameter 'roleArn' to be non-null");
            return $;
        }
    }

}
