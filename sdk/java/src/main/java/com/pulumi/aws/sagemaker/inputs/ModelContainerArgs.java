// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.sagemaker.inputs;

import com.pulumi.aws.sagemaker.inputs.ModelContainerImageConfigArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ModelContainerArgs extends com.pulumi.resources.ResourceArgs {

    public static final ModelContainerArgs Empty = new ModelContainerArgs();

    /**
     * The DNS host name for the container.
     * 
     */
    @Import(name="containerHostname")
    private @Nullable Output<String> containerHostname;

    /**
     * @return The DNS host name for the container.
     * 
     */
    public Optional<Output<String>> containerHostname() {
        return Optional.ofNullable(this.containerHostname);
    }

    /**
     * Environment variables for the Docker container.
     * A list of key value pairs.
     * 
     */
    @Import(name="environment")
    private @Nullable Output<Map<String,String>> environment;

    /**
     * @return Environment variables for the Docker container.
     * A list of key value pairs.
     * 
     */
    public Optional<Output<Map<String,String>>> environment() {
        return Optional.ofNullable(this.environment);
    }

    /**
     * The registry path where the inference code image is stored in Amazon ECR.
     * 
     */
    @Import(name="image")
    private @Nullable Output<String> image;

    /**
     * @return The registry path where the inference code image is stored in Amazon ECR.
     * 
     */
    public Optional<Output<String>> image() {
        return Optional.ofNullable(this.image);
    }

    /**
     * Specifies whether the model container is in Amazon ECR or a private Docker registry accessible from your Amazon Virtual Private Cloud (VPC). For more information see [Using a Private Docker Registry for Real-Time Inference Containers](https://docs.aws.amazon.com/sagemaker/latest/dg/your-algorithms-containers-inference-private.html). see Image Config.
     * 
     */
    @Import(name="imageConfig")
    private @Nullable Output<ModelContainerImageConfigArgs> imageConfig;

    /**
     * @return Specifies whether the model container is in Amazon ECR or a private Docker registry accessible from your Amazon Virtual Private Cloud (VPC). For more information see [Using a Private Docker Registry for Real-Time Inference Containers](https://docs.aws.amazon.com/sagemaker/latest/dg/your-algorithms-containers-inference-private.html). see Image Config.
     * 
     */
    public Optional<Output<ModelContainerImageConfigArgs>> imageConfig() {
        return Optional.ofNullable(this.imageConfig);
    }

    /**
     * The container hosts value `SingleModel/MultiModel`. The default value is `SingleModel`.
     * 
     */
    @Import(name="mode")
    private @Nullable Output<String> mode;

    /**
     * @return The container hosts value `SingleModel/MultiModel`. The default value is `SingleModel`.
     * 
     */
    public Optional<Output<String>> mode() {
        return Optional.ofNullable(this.mode);
    }

    /**
     * The URL for the S3 location where model artifacts are stored.
     * 
     */
    @Import(name="modelDataUrl")
    private @Nullable Output<String> modelDataUrl;

    /**
     * @return The URL for the S3 location where model artifacts are stored.
     * 
     */
    public Optional<Output<String>> modelDataUrl() {
        return Optional.ofNullable(this.modelDataUrl);
    }

    /**
     * The Amazon Resource Name (ARN) of the model package to use to create the model.
     * 
     */
    @Import(name="modelPackageName")
    private @Nullable Output<String> modelPackageName;

    /**
     * @return The Amazon Resource Name (ARN) of the model package to use to create the model.
     * 
     */
    public Optional<Output<String>> modelPackageName() {
        return Optional.ofNullable(this.modelPackageName);
    }

    private ModelContainerArgs() {}

    private ModelContainerArgs(ModelContainerArgs $) {
        this.containerHostname = $.containerHostname;
        this.environment = $.environment;
        this.image = $.image;
        this.imageConfig = $.imageConfig;
        this.mode = $.mode;
        this.modelDataUrl = $.modelDataUrl;
        this.modelPackageName = $.modelPackageName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ModelContainerArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ModelContainerArgs $;

        public Builder() {
            $ = new ModelContainerArgs();
        }

        public Builder(ModelContainerArgs defaults) {
            $ = new ModelContainerArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param containerHostname The DNS host name for the container.
         * 
         * @return builder
         * 
         */
        public Builder containerHostname(@Nullable Output<String> containerHostname) {
            $.containerHostname = containerHostname;
            return this;
        }

        /**
         * @param containerHostname The DNS host name for the container.
         * 
         * @return builder
         * 
         */
        public Builder containerHostname(String containerHostname) {
            return containerHostname(Output.of(containerHostname));
        }

        /**
         * @param environment Environment variables for the Docker container.
         * A list of key value pairs.
         * 
         * @return builder
         * 
         */
        public Builder environment(@Nullable Output<Map<String,String>> environment) {
            $.environment = environment;
            return this;
        }

        /**
         * @param environment Environment variables for the Docker container.
         * A list of key value pairs.
         * 
         * @return builder
         * 
         */
        public Builder environment(Map<String,String> environment) {
            return environment(Output.of(environment));
        }

        /**
         * @param image The registry path where the inference code image is stored in Amazon ECR.
         * 
         * @return builder
         * 
         */
        public Builder image(@Nullable Output<String> image) {
            $.image = image;
            return this;
        }

        /**
         * @param image The registry path where the inference code image is stored in Amazon ECR.
         * 
         * @return builder
         * 
         */
        public Builder image(String image) {
            return image(Output.of(image));
        }

        /**
         * @param imageConfig Specifies whether the model container is in Amazon ECR or a private Docker registry accessible from your Amazon Virtual Private Cloud (VPC). For more information see [Using a Private Docker Registry for Real-Time Inference Containers](https://docs.aws.amazon.com/sagemaker/latest/dg/your-algorithms-containers-inference-private.html). see Image Config.
         * 
         * @return builder
         * 
         */
        public Builder imageConfig(@Nullable Output<ModelContainerImageConfigArgs> imageConfig) {
            $.imageConfig = imageConfig;
            return this;
        }

        /**
         * @param imageConfig Specifies whether the model container is in Amazon ECR or a private Docker registry accessible from your Amazon Virtual Private Cloud (VPC). For more information see [Using a Private Docker Registry for Real-Time Inference Containers](https://docs.aws.amazon.com/sagemaker/latest/dg/your-algorithms-containers-inference-private.html). see Image Config.
         * 
         * @return builder
         * 
         */
        public Builder imageConfig(ModelContainerImageConfigArgs imageConfig) {
            return imageConfig(Output.of(imageConfig));
        }

        /**
         * @param mode The container hosts value `SingleModel/MultiModel`. The default value is `SingleModel`.
         * 
         * @return builder
         * 
         */
        public Builder mode(@Nullable Output<String> mode) {
            $.mode = mode;
            return this;
        }

        /**
         * @param mode The container hosts value `SingleModel/MultiModel`. The default value is `SingleModel`.
         * 
         * @return builder
         * 
         */
        public Builder mode(String mode) {
            return mode(Output.of(mode));
        }

        /**
         * @param modelDataUrl The URL for the S3 location where model artifacts are stored.
         * 
         * @return builder
         * 
         */
        public Builder modelDataUrl(@Nullable Output<String> modelDataUrl) {
            $.modelDataUrl = modelDataUrl;
            return this;
        }

        /**
         * @param modelDataUrl The URL for the S3 location where model artifacts are stored.
         * 
         * @return builder
         * 
         */
        public Builder modelDataUrl(String modelDataUrl) {
            return modelDataUrl(Output.of(modelDataUrl));
        }

        /**
         * @param modelPackageName The Amazon Resource Name (ARN) of the model package to use to create the model.
         * 
         * @return builder
         * 
         */
        public Builder modelPackageName(@Nullable Output<String> modelPackageName) {
            $.modelPackageName = modelPackageName;
            return this;
        }

        /**
         * @param modelPackageName The Amazon Resource Name (ARN) of the model package to use to create the model.
         * 
         * @return builder
         * 
         */
        public Builder modelPackageName(String modelPackageName) {
            return modelPackageName(Output.of(modelPackageName));
        }

        public ModelContainerArgs build() {
            return $;
        }
    }

}
