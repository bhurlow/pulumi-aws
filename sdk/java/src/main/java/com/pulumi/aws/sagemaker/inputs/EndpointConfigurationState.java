// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.sagemaker.inputs;

import com.pulumi.aws.sagemaker.inputs.EndpointConfigurationAsyncInferenceConfigArgs;
import com.pulumi.aws.sagemaker.inputs.EndpointConfigurationDataCaptureConfigArgs;
import com.pulumi.aws.sagemaker.inputs.EndpointConfigurationProductionVariantArgs;
import com.pulumi.aws.sagemaker.inputs.EndpointConfigurationShadowProductionVariantArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class EndpointConfigurationState extends com.pulumi.resources.ResourceArgs {

    public static final EndpointConfigurationState Empty = new EndpointConfigurationState();

    /**
     * The Amazon Resource Name (ARN) assigned by AWS to this endpoint configuration.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return The Amazon Resource Name (ARN) assigned by AWS to this endpoint configuration.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * Specifies configuration for how an endpoint performs asynchronous inference.
     * 
     */
    @Import(name="asyncInferenceConfig")
    private @Nullable Output<EndpointConfigurationAsyncInferenceConfigArgs> asyncInferenceConfig;

    /**
     * @return Specifies configuration for how an endpoint performs asynchronous inference.
     * 
     */
    public Optional<Output<EndpointConfigurationAsyncInferenceConfigArgs>> asyncInferenceConfig() {
        return Optional.ofNullable(this.asyncInferenceConfig);
    }

    /**
     * Specifies the parameters to capture input/output of SageMaker models endpoints. Fields are documented below.
     * 
     */
    @Import(name="dataCaptureConfig")
    private @Nullable Output<EndpointConfigurationDataCaptureConfigArgs> dataCaptureConfig;

    /**
     * @return Specifies the parameters to capture input/output of SageMaker models endpoints. Fields are documented below.
     * 
     */
    public Optional<Output<EndpointConfigurationDataCaptureConfigArgs>> dataCaptureConfig() {
        return Optional.ofNullable(this.dataCaptureConfig);
    }

    /**
     * Amazon Resource Name (ARN) of a AWS Key Management Service key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.
     * 
     */
    @Import(name="kmsKeyArn")
    private @Nullable Output<String> kmsKeyArn;

    /**
     * @return Amazon Resource Name (ARN) of a AWS Key Management Service key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.
     * 
     */
    public Optional<Output<String>> kmsKeyArn() {
        return Optional.ofNullable(this.kmsKeyArn);
    }

    /**
     * The name of the endpoint configuration. If omitted, this provider will assign a random, unique name. Conflicts with `name_prefix`.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the endpoint configuration. If omitted, this provider will assign a random, unique name. Conflicts with `name_prefix`.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Creates a unique endpoint configuration name beginning with the specified prefix. Conflicts with `name`.
     * 
     */
    @Import(name="namePrefix")
    private @Nullable Output<String> namePrefix;

    /**
     * @return Creates a unique endpoint configuration name beginning with the specified prefix. Conflicts with `name`.
     * 
     */
    public Optional<Output<String>> namePrefix() {
        return Optional.ofNullable(this.namePrefix);
    }

    /**
     * An list of ProductionVariant objects, one for each model that you want to host at this endpoint. Fields are documented below.
     * 
     */
    @Import(name="productionVariants")
    private @Nullable Output<List<EndpointConfigurationProductionVariantArgs>> productionVariants;

    /**
     * @return An list of ProductionVariant objects, one for each model that you want to host at this endpoint. Fields are documented below.
     * 
     */
    public Optional<Output<List<EndpointConfigurationProductionVariantArgs>>> productionVariants() {
        return Optional.ofNullable(this.productionVariants);
    }

    /**
     * Array of ProductionVariant objects. There is one for each model that you want to host at this endpoint in shadow mode with production traffic replicated from the model specified on ProductionVariants.If you use this field, you can only specify one variant for ProductionVariants and one variant for ShadowProductionVariants. Fields are documented below.
     * 
     */
    @Import(name="shadowProductionVariants")
    private @Nullable Output<List<EndpointConfigurationShadowProductionVariantArgs>> shadowProductionVariants;

    /**
     * @return Array of ProductionVariant objects. There is one for each model that you want to host at this endpoint in shadow mode with production traffic replicated from the model specified on ProductionVariants.If you use this field, you can only specify one variant for ProductionVariants and one variant for ShadowProductionVariants. Fields are documented below.
     * 
     */
    public Optional<Output<List<EndpointConfigurationShadowProductionVariantArgs>>> shadowProductionVariants() {
        return Optional.ofNullable(this.shadowProductionVariants);
    }

    /**
     * A mapping of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A mapping of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

    private EndpointConfigurationState() {}

    private EndpointConfigurationState(EndpointConfigurationState $) {
        this.arn = $.arn;
        this.asyncInferenceConfig = $.asyncInferenceConfig;
        this.dataCaptureConfig = $.dataCaptureConfig;
        this.kmsKeyArn = $.kmsKeyArn;
        this.name = $.name;
        this.namePrefix = $.namePrefix;
        this.productionVariants = $.productionVariants;
        this.shadowProductionVariants = $.shadowProductionVariants;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EndpointConfigurationState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EndpointConfigurationState $;

        public Builder() {
            $ = new EndpointConfigurationState();
        }

        public Builder(EndpointConfigurationState defaults) {
            $ = new EndpointConfigurationState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn The Amazon Resource Name (ARN) assigned by AWS to this endpoint configuration.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn The Amazon Resource Name (ARN) assigned by AWS to this endpoint configuration.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param asyncInferenceConfig Specifies configuration for how an endpoint performs asynchronous inference.
         * 
         * @return builder
         * 
         */
        public Builder asyncInferenceConfig(@Nullable Output<EndpointConfigurationAsyncInferenceConfigArgs> asyncInferenceConfig) {
            $.asyncInferenceConfig = asyncInferenceConfig;
            return this;
        }

        /**
         * @param asyncInferenceConfig Specifies configuration for how an endpoint performs asynchronous inference.
         * 
         * @return builder
         * 
         */
        public Builder asyncInferenceConfig(EndpointConfigurationAsyncInferenceConfigArgs asyncInferenceConfig) {
            return asyncInferenceConfig(Output.of(asyncInferenceConfig));
        }

        /**
         * @param dataCaptureConfig Specifies the parameters to capture input/output of SageMaker models endpoints. Fields are documented below.
         * 
         * @return builder
         * 
         */
        public Builder dataCaptureConfig(@Nullable Output<EndpointConfigurationDataCaptureConfigArgs> dataCaptureConfig) {
            $.dataCaptureConfig = dataCaptureConfig;
            return this;
        }

        /**
         * @param dataCaptureConfig Specifies the parameters to capture input/output of SageMaker models endpoints. Fields are documented below.
         * 
         * @return builder
         * 
         */
        public Builder dataCaptureConfig(EndpointConfigurationDataCaptureConfigArgs dataCaptureConfig) {
            return dataCaptureConfig(Output.of(dataCaptureConfig));
        }

        /**
         * @param kmsKeyArn Amazon Resource Name (ARN) of a AWS Key Management Service key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.
         * 
         * @return builder
         * 
         */
        public Builder kmsKeyArn(@Nullable Output<String> kmsKeyArn) {
            $.kmsKeyArn = kmsKeyArn;
            return this;
        }

        /**
         * @param kmsKeyArn Amazon Resource Name (ARN) of a AWS Key Management Service key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.
         * 
         * @return builder
         * 
         */
        public Builder kmsKeyArn(String kmsKeyArn) {
            return kmsKeyArn(Output.of(kmsKeyArn));
        }

        /**
         * @param name The name of the endpoint configuration. If omitted, this provider will assign a random, unique name. Conflicts with `name_prefix`.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the endpoint configuration. If omitted, this provider will assign a random, unique name. Conflicts with `name_prefix`.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param namePrefix Creates a unique endpoint configuration name beginning with the specified prefix. Conflicts with `name`.
         * 
         * @return builder
         * 
         */
        public Builder namePrefix(@Nullable Output<String> namePrefix) {
            $.namePrefix = namePrefix;
            return this;
        }

        /**
         * @param namePrefix Creates a unique endpoint configuration name beginning with the specified prefix. Conflicts with `name`.
         * 
         * @return builder
         * 
         */
        public Builder namePrefix(String namePrefix) {
            return namePrefix(Output.of(namePrefix));
        }

        /**
         * @param productionVariants An list of ProductionVariant objects, one for each model that you want to host at this endpoint. Fields are documented below.
         * 
         * @return builder
         * 
         */
        public Builder productionVariants(@Nullable Output<List<EndpointConfigurationProductionVariantArgs>> productionVariants) {
            $.productionVariants = productionVariants;
            return this;
        }

        /**
         * @param productionVariants An list of ProductionVariant objects, one for each model that you want to host at this endpoint. Fields are documented below.
         * 
         * @return builder
         * 
         */
        public Builder productionVariants(List<EndpointConfigurationProductionVariantArgs> productionVariants) {
            return productionVariants(Output.of(productionVariants));
        }

        /**
         * @param productionVariants An list of ProductionVariant objects, one for each model that you want to host at this endpoint. Fields are documented below.
         * 
         * @return builder
         * 
         */
        public Builder productionVariants(EndpointConfigurationProductionVariantArgs... productionVariants) {
            return productionVariants(List.of(productionVariants));
        }

        /**
         * @param shadowProductionVariants Array of ProductionVariant objects. There is one for each model that you want to host at this endpoint in shadow mode with production traffic replicated from the model specified on ProductionVariants.If you use this field, you can only specify one variant for ProductionVariants and one variant for ShadowProductionVariants. Fields are documented below.
         * 
         * @return builder
         * 
         */
        public Builder shadowProductionVariants(@Nullable Output<List<EndpointConfigurationShadowProductionVariantArgs>> shadowProductionVariants) {
            $.shadowProductionVariants = shadowProductionVariants;
            return this;
        }

        /**
         * @param shadowProductionVariants Array of ProductionVariant objects. There is one for each model that you want to host at this endpoint in shadow mode with production traffic replicated from the model specified on ProductionVariants.If you use this field, you can only specify one variant for ProductionVariants and one variant for ShadowProductionVariants. Fields are documented below.
         * 
         * @return builder
         * 
         */
        public Builder shadowProductionVariants(List<EndpointConfigurationShadowProductionVariantArgs> shadowProductionVariants) {
            return shadowProductionVariants(Output.of(shadowProductionVariants));
        }

        /**
         * @param shadowProductionVariants Array of ProductionVariant objects. There is one for each model that you want to host at this endpoint in shadow mode with production traffic replicated from the model specified on ProductionVariants.If you use this field, you can only specify one variant for ProductionVariants and one variant for ShadowProductionVariants. Fields are documented below.
         * 
         * @return builder
         * 
         */
        public Builder shadowProductionVariants(EndpointConfigurationShadowProductionVariantArgs... shadowProductionVariants) {
            return shadowProductionVariants(List.of(shadowProductionVariants));
        }

        /**
         * @param tags A mapping of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A mapping of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

        public EndpointConfigurationState build() {
            return $;
        }
    }

}
