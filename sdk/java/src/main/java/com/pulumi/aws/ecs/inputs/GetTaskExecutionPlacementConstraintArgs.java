// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ecs.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetTaskExecutionPlacementConstraintArgs extends com.pulumi.resources.ResourceArgs {

    public static final GetTaskExecutionPlacementConstraintArgs Empty = new GetTaskExecutionPlacementConstraintArgs();

    /**
     * A cluster query language expression to apply to the constraint. The expression can have a maximum length of 2000 characters. You can&#39;t specify an expression if the constraint type is `distinctInstance`.
     * 
     */
    @Import(name="expression")
    private @Nullable Output<String> expression;

    /**
     * @return A cluster query language expression to apply to the constraint. The expression can have a maximum length of 2000 characters. You can&#39;t specify an expression if the constraint type is `distinctInstance`.
     * 
     */
    public Optional<Output<String>> expression() {
        return Optional.ofNullable(this.expression);
    }

    /**
     * The type of constraint. Valid values are `distinctInstance` or `memberOf`. Use `distinctInstance` to ensure that each task in a particular group is running on a different container instance. Use `memberOf` to restrict the selection to a group of valid candidates.
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return The type of constraint. Valid values are `distinctInstance` or `memberOf`. Use `distinctInstance` to ensure that each task in a particular group is running on a different container instance. Use `memberOf` to restrict the selection to a group of valid candidates.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    private GetTaskExecutionPlacementConstraintArgs() {}

    private GetTaskExecutionPlacementConstraintArgs(GetTaskExecutionPlacementConstraintArgs $) {
        this.expression = $.expression;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetTaskExecutionPlacementConstraintArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetTaskExecutionPlacementConstraintArgs $;

        public Builder() {
            $ = new GetTaskExecutionPlacementConstraintArgs();
        }

        public Builder(GetTaskExecutionPlacementConstraintArgs defaults) {
            $ = new GetTaskExecutionPlacementConstraintArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param expression A cluster query language expression to apply to the constraint. The expression can have a maximum length of 2000 characters. You can&#39;t specify an expression if the constraint type is `distinctInstance`.
         * 
         * @return builder
         * 
         */
        public Builder expression(@Nullable Output<String> expression) {
            $.expression = expression;
            return this;
        }

        /**
         * @param expression A cluster query language expression to apply to the constraint. The expression can have a maximum length of 2000 characters. You can&#39;t specify an expression if the constraint type is `distinctInstance`.
         * 
         * @return builder
         * 
         */
        public Builder expression(String expression) {
            return expression(Output.of(expression));
        }

        /**
         * @param type The type of constraint. Valid values are `distinctInstance` or `memberOf`. Use `distinctInstance` to ensure that each task in a particular group is running on a different container instance. Use `memberOf` to restrict the selection to a group of valid candidates.
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type The type of constraint. Valid values are `distinctInstance` or `memberOf`. Use `distinctInstance` to ensure that each task in a particular group is running on a different container instance. Use `memberOf` to restrict the selection to a group of valid candidates.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public GetTaskExecutionPlacementConstraintArgs build() {
            $.type = Objects.requireNonNull($.type, "expected parameter 'type' to be non-null");
            return $;
        }
    }

}
