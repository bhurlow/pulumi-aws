// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ecs.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetTaskExecutionOverridesContainerOverrideResourceRequirementArgs extends com.pulumi.resources.ResourceArgs {

    public static final GetTaskExecutionOverridesContainerOverrideResourceRequirementArgs Empty = new GetTaskExecutionOverridesContainerOverrideResourceRequirementArgs();

    /**
     * The type of resource to assign to a container. Valid values are `GPU` or `InferenceAccelerator`.
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return The type of resource to assign to a container. Valid values are `GPU` or `InferenceAccelerator`.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     * The value for the specified resource type. If the `GPU` type is used, the value is the number of physical GPUs the Amazon ECS container agent reserves for the container. The number of GPUs that&#39;s reserved for all containers in a task can&#39;t exceed the number of available GPUs on the container instance that the task is launched on. If the `InferenceAccelerator` type is used, the value matches the `deviceName` for an InferenceAccelerator specified in a task definition.
     * 
     */
    @Import(name="value", required=true)
    private Output<String> value;

    /**
     * @return The value for the specified resource type. If the `GPU` type is used, the value is the number of physical GPUs the Amazon ECS container agent reserves for the container. The number of GPUs that&#39;s reserved for all containers in a task can&#39;t exceed the number of available GPUs on the container instance that the task is launched on. If the `InferenceAccelerator` type is used, the value matches the `deviceName` for an InferenceAccelerator specified in a task definition.
     * 
     */
    public Output<String> value() {
        return this.value;
    }

    private GetTaskExecutionOverridesContainerOverrideResourceRequirementArgs() {}

    private GetTaskExecutionOverridesContainerOverrideResourceRequirementArgs(GetTaskExecutionOverridesContainerOverrideResourceRequirementArgs $) {
        this.type = $.type;
        this.value = $.value;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetTaskExecutionOverridesContainerOverrideResourceRequirementArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetTaskExecutionOverridesContainerOverrideResourceRequirementArgs $;

        public Builder() {
            $ = new GetTaskExecutionOverridesContainerOverrideResourceRequirementArgs();
        }

        public Builder(GetTaskExecutionOverridesContainerOverrideResourceRequirementArgs defaults) {
            $ = new GetTaskExecutionOverridesContainerOverrideResourceRequirementArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param type The type of resource to assign to a container. Valid values are `GPU` or `InferenceAccelerator`.
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type The type of resource to assign to a container. Valid values are `GPU` or `InferenceAccelerator`.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        /**
         * @param value The value for the specified resource type. If the `GPU` type is used, the value is the number of physical GPUs the Amazon ECS container agent reserves for the container. The number of GPUs that&#39;s reserved for all containers in a task can&#39;t exceed the number of available GPUs on the container instance that the task is launched on. If the `InferenceAccelerator` type is used, the value matches the `deviceName` for an InferenceAccelerator specified in a task definition.
         * 
         * @return builder
         * 
         */
        public Builder value(Output<String> value) {
            $.value = value;
            return this;
        }

        /**
         * @param value The value for the specified resource type. If the `GPU` type is used, the value is the number of physical GPUs the Amazon ECS container agent reserves for the container. The number of GPUs that&#39;s reserved for all containers in a task can&#39;t exceed the number of available GPUs on the container instance that the task is launched on. If the `InferenceAccelerator` type is used, the value matches the `deviceName` for an InferenceAccelerator specified in a task definition.
         * 
         * @return builder
         * 
         */
        public Builder value(String value) {
            return value(Output.of(value));
        }

        public GetTaskExecutionOverridesContainerOverrideResourceRequirementArgs build() {
            $.type = Objects.requireNonNull($.type, "expected parameter 'type' to be non-null");
            $.value = Objects.requireNonNull($.value, "expected parameter 'value' to be non-null");
            return $;
        }
    }

}
