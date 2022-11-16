// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.transcribe;

import com.pulumi.aws.transcribe.inputs.LanguageModelInputDataConfigArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class LanguageModelArgs extends com.pulumi.resources.ResourceArgs {

    public static final LanguageModelArgs Empty = new LanguageModelArgs();

    /**
     * Name of reference base model.
     * 
     */
    @Import(name="baseModelName", required=true)
    private Output<String> baseModelName;

    /**
     * @return Name of reference base model.
     * 
     */
    public Output<String> baseModelName() {
        return this.baseModelName;
    }

    /**
     * The input data config for the LanguageModel. See Input Data Config for more details.
     * 
     */
    @Import(name="inputDataConfig", required=true)
    private Output<LanguageModelInputDataConfigArgs> inputDataConfig;

    /**
     * @return The input data config for the LanguageModel. See Input Data Config for more details.
     * 
     */
    public Output<LanguageModelInputDataConfigArgs> inputDataConfig() {
        return this.inputDataConfig;
    }

    /**
     * The language code you selected for your language model. Refer to the [supported languages](https://docs.aws.amazon.com/transcribe/latest/dg/supported-languages.html) page for accepted codes.
     * 
     */
    @Import(name="languageCode", required=true)
    private Output<String> languageCode;

    /**
     * @return The language code you selected for your language model. Refer to the [supported languages](https://docs.aws.amazon.com/transcribe/latest/dg/supported-languages.html) page for accepted codes.
     * 
     */
    public Output<String> languageCode() {
        return this.languageCode;
    }

    /**
     * The model name.
     * 
     */
    @Import(name="modelName", required=true)
    private Output<String> modelName;

    /**
     * @return The model name.
     * 
     */
    public Output<String> modelName() {
        return this.modelName;
    }

    /**
     * A map of tags to assign to the LanguageModel. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the LanguageModel. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private LanguageModelArgs() {}

    private LanguageModelArgs(LanguageModelArgs $) {
        this.baseModelName = $.baseModelName;
        this.inputDataConfig = $.inputDataConfig;
        this.languageCode = $.languageCode;
        this.modelName = $.modelName;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LanguageModelArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LanguageModelArgs $;

        public Builder() {
            $ = new LanguageModelArgs();
        }

        public Builder(LanguageModelArgs defaults) {
            $ = new LanguageModelArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param baseModelName Name of reference base model.
         * 
         * @return builder
         * 
         */
        public Builder baseModelName(Output<String> baseModelName) {
            $.baseModelName = baseModelName;
            return this;
        }

        /**
         * @param baseModelName Name of reference base model.
         * 
         * @return builder
         * 
         */
        public Builder baseModelName(String baseModelName) {
            return baseModelName(Output.of(baseModelName));
        }

        /**
         * @param inputDataConfig The input data config for the LanguageModel. See Input Data Config for more details.
         * 
         * @return builder
         * 
         */
        public Builder inputDataConfig(Output<LanguageModelInputDataConfigArgs> inputDataConfig) {
            $.inputDataConfig = inputDataConfig;
            return this;
        }

        /**
         * @param inputDataConfig The input data config for the LanguageModel. See Input Data Config for more details.
         * 
         * @return builder
         * 
         */
        public Builder inputDataConfig(LanguageModelInputDataConfigArgs inputDataConfig) {
            return inputDataConfig(Output.of(inputDataConfig));
        }

        /**
         * @param languageCode The language code you selected for your language model. Refer to the [supported languages](https://docs.aws.amazon.com/transcribe/latest/dg/supported-languages.html) page for accepted codes.
         * 
         * @return builder
         * 
         */
        public Builder languageCode(Output<String> languageCode) {
            $.languageCode = languageCode;
            return this;
        }

        /**
         * @param languageCode The language code you selected for your language model. Refer to the [supported languages](https://docs.aws.amazon.com/transcribe/latest/dg/supported-languages.html) page for accepted codes.
         * 
         * @return builder
         * 
         */
        public Builder languageCode(String languageCode) {
            return languageCode(Output.of(languageCode));
        }

        /**
         * @param modelName The model name.
         * 
         * @return builder
         * 
         */
        public Builder modelName(Output<String> modelName) {
            $.modelName = modelName;
            return this;
        }

        /**
         * @param modelName The model name.
         * 
         * @return builder
         * 
         */
        public Builder modelName(String modelName) {
            return modelName(Output.of(modelName));
        }

        /**
         * @param tags A map of tags to assign to the LanguageModel. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A map of tags to assign to the LanguageModel. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        public LanguageModelArgs build() {
            $.baseModelName = Objects.requireNonNull($.baseModelName, "expected parameter 'baseModelName' to be non-null");
            $.inputDataConfig = Objects.requireNonNull($.inputDataConfig, "expected parameter 'inputDataConfig' to be non-null");
            $.languageCode = Objects.requireNonNull($.languageCode, "expected parameter 'languageCode' to be non-null");
            $.modelName = Objects.requireNonNull($.modelName, "expected parameter 'modelName' to be non-null");
            return $;
        }
    }

}
