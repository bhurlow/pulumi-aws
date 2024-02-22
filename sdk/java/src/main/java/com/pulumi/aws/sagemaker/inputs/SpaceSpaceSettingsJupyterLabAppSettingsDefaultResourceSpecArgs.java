// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.sagemaker.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecArgs extends com.pulumi.resources.ResourceArgs {

    public static final SpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecArgs Empty = new SpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecArgs();

    /**
     * The instance type.
     * 
     */
    @Import(name="instanceType")
    private @Nullable Output<String> instanceType;

    /**
     * @return The instance type.
     * 
     */
    public Optional<Output<String>> instanceType() {
        return Optional.ofNullable(this.instanceType);
    }

    /**
     * The Amazon Resource Name (ARN) of the Lifecycle Configuration attached to the Resource.
     * 
     */
    @Import(name="lifecycleConfigArn")
    private @Nullable Output<String> lifecycleConfigArn;

    /**
     * @return The Amazon Resource Name (ARN) of the Lifecycle Configuration attached to the Resource.
     * 
     */
    public Optional<Output<String>> lifecycleConfigArn() {
        return Optional.ofNullable(this.lifecycleConfigArn);
    }

    /**
     * The Amazon Resource Name (ARN) of the SageMaker image created on the instance.
     * 
     */
    @Import(name="sagemakerImageArn")
    private @Nullable Output<String> sagemakerImageArn;

    /**
     * @return The Amazon Resource Name (ARN) of the SageMaker image created on the instance.
     * 
     */
    public Optional<Output<String>> sagemakerImageArn() {
        return Optional.ofNullable(this.sagemakerImageArn);
    }

    /**
     * The SageMaker Image Version Alias.
     * 
     */
    @Import(name="sagemakerImageVersionAlias")
    private @Nullable Output<String> sagemakerImageVersionAlias;

    /**
     * @return The SageMaker Image Version Alias.
     * 
     */
    public Optional<Output<String>> sagemakerImageVersionAlias() {
        return Optional.ofNullable(this.sagemakerImageVersionAlias);
    }

    /**
     * The ARN of the image version created on the instance.
     * 
     */
    @Import(name="sagemakerImageVersionArn")
    private @Nullable Output<String> sagemakerImageVersionArn;

    /**
     * @return The ARN of the image version created on the instance.
     * 
     */
    public Optional<Output<String>> sagemakerImageVersionArn() {
        return Optional.ofNullable(this.sagemakerImageVersionArn);
    }

    private SpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecArgs() {}

    private SpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecArgs(SpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecArgs $) {
        this.instanceType = $.instanceType;
        this.lifecycleConfigArn = $.lifecycleConfigArn;
        this.sagemakerImageArn = $.sagemakerImageArn;
        this.sagemakerImageVersionAlias = $.sagemakerImageVersionAlias;
        this.sagemakerImageVersionArn = $.sagemakerImageVersionArn;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecArgs $;

        public Builder() {
            $ = new SpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecArgs();
        }

        public Builder(SpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecArgs defaults) {
            $ = new SpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param instanceType The instance type.
         * 
         * @return builder
         * 
         */
        public Builder instanceType(@Nullable Output<String> instanceType) {
            $.instanceType = instanceType;
            return this;
        }

        /**
         * @param instanceType The instance type.
         * 
         * @return builder
         * 
         */
        public Builder instanceType(String instanceType) {
            return instanceType(Output.of(instanceType));
        }

        /**
         * @param lifecycleConfigArn The Amazon Resource Name (ARN) of the Lifecycle Configuration attached to the Resource.
         * 
         * @return builder
         * 
         */
        public Builder lifecycleConfigArn(@Nullable Output<String> lifecycleConfigArn) {
            $.lifecycleConfigArn = lifecycleConfigArn;
            return this;
        }

        /**
         * @param lifecycleConfigArn The Amazon Resource Name (ARN) of the Lifecycle Configuration attached to the Resource.
         * 
         * @return builder
         * 
         */
        public Builder lifecycleConfigArn(String lifecycleConfigArn) {
            return lifecycleConfigArn(Output.of(lifecycleConfigArn));
        }

        /**
         * @param sagemakerImageArn The Amazon Resource Name (ARN) of the SageMaker image created on the instance.
         * 
         * @return builder
         * 
         */
        public Builder sagemakerImageArn(@Nullable Output<String> sagemakerImageArn) {
            $.sagemakerImageArn = sagemakerImageArn;
            return this;
        }

        /**
         * @param sagemakerImageArn The Amazon Resource Name (ARN) of the SageMaker image created on the instance.
         * 
         * @return builder
         * 
         */
        public Builder sagemakerImageArn(String sagemakerImageArn) {
            return sagemakerImageArn(Output.of(sagemakerImageArn));
        }

        /**
         * @param sagemakerImageVersionAlias The SageMaker Image Version Alias.
         * 
         * @return builder
         * 
         */
        public Builder sagemakerImageVersionAlias(@Nullable Output<String> sagemakerImageVersionAlias) {
            $.sagemakerImageVersionAlias = sagemakerImageVersionAlias;
            return this;
        }

        /**
         * @param sagemakerImageVersionAlias The SageMaker Image Version Alias.
         * 
         * @return builder
         * 
         */
        public Builder sagemakerImageVersionAlias(String sagemakerImageVersionAlias) {
            return sagemakerImageVersionAlias(Output.of(sagemakerImageVersionAlias));
        }

        /**
         * @param sagemakerImageVersionArn The ARN of the image version created on the instance.
         * 
         * @return builder
         * 
         */
        public Builder sagemakerImageVersionArn(@Nullable Output<String> sagemakerImageVersionArn) {
            $.sagemakerImageVersionArn = sagemakerImageVersionArn;
            return this;
        }

        /**
         * @param sagemakerImageVersionArn The ARN of the image version created on the instance.
         * 
         * @return builder
         * 
         */
        public Builder sagemakerImageVersionArn(String sagemakerImageVersionArn) {
            return sagemakerImageVersionArn(Output.of(sagemakerImageVersionArn));
        }

        public SpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpecArgs build() {
            return $;
        }
    }

}
