// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ecs.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetTaskExecutionPlacementConstraint {
    /**
     * @return A cluster query language expression to apply to the constraint. The expression can have a maximum length of 2000 characters. You can&#39;t specify an expression if the constraint type is `distinctInstance`.
     * 
     */
    private @Nullable String expression;
    /**
     * @return The type of constraint. Valid values are `distinctInstance` or `memberOf`. Use `distinctInstance` to ensure that each task in a particular group is running on a different container instance. Use `memberOf` to restrict the selection to a group of valid candidates.
     * 
     */
    private String type;

    private GetTaskExecutionPlacementConstraint() {}
    /**
     * @return A cluster query language expression to apply to the constraint. The expression can have a maximum length of 2000 characters. You can&#39;t specify an expression if the constraint type is `distinctInstance`.
     * 
     */
    public Optional<String> expression() {
        return Optional.ofNullable(this.expression);
    }
    /**
     * @return The type of constraint. Valid values are `distinctInstance` or `memberOf`. Use `distinctInstance` to ensure that each task in a particular group is running on a different container instance. Use `memberOf` to restrict the selection to a group of valid candidates.
     * 
     */
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetTaskExecutionPlacementConstraint defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String expression;
        private String type;
        public Builder() {}
        public Builder(GetTaskExecutionPlacementConstraint defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.expression = defaults.expression;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder expression(@Nullable String expression) {
            this.expression = expression;
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }
        public GetTaskExecutionPlacementConstraint build() {
            final var o = new GetTaskExecutionPlacementConstraint();
            o.expression = expression;
            o.type = type;
            return o;
        }
    }
}
