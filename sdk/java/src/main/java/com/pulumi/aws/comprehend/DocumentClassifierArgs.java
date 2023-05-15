// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.comprehend;

import com.pulumi.aws.comprehend.inputs.DocumentClassifierInputDataConfigArgs;
import com.pulumi.aws.comprehend.inputs.DocumentClassifierOutputDataConfigArgs;
import com.pulumi.aws.comprehend.inputs.DocumentClassifierVpcConfigArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DocumentClassifierArgs extends com.pulumi.resources.ResourceArgs {

    public static final DocumentClassifierArgs Empty = new DocumentClassifierArgs();

    /**
     * The ARN for an IAM Role which allows Comprehend to read the training and testing data.
     * 
     */
    @Import(name="dataAccessRoleArn", required=true)
    private Output<String> dataAccessRoleArn;

    /**
     * @return The ARN for an IAM Role which allows Comprehend to read the training and testing data.
     * 
     */
    public Output<String> dataAccessRoleArn() {
        return this.dataAccessRoleArn;
    }

    /**
     * Configuration for the training and testing data.
     * See the `input_data_config` Configuration Block section below.
     * 
     */
    @Import(name="inputDataConfig", required=true)
    private Output<DocumentClassifierInputDataConfigArgs> inputDataConfig;

    /**
     * @return Configuration for the training and testing data.
     * See the `input_data_config` Configuration Block section below.
     * 
     */
    public Output<DocumentClassifierInputDataConfigArgs> inputDataConfig() {
        return this.inputDataConfig;
    }

    /**
     * Two-letter language code for the language.
     * One of `en`, `es`, `fr`, `it`, `de`, or `pt`.
     * 
     */
    @Import(name="languageCode", required=true)
    private Output<String> languageCode;

    /**
     * @return Two-letter language code for the language.
     * One of `en`, `es`, `fr`, `it`, `de`, or `pt`.
     * 
     */
    public Output<String> languageCode() {
        return this.languageCode;
    }

    /**
     * The document classification mode.
     * One of `MULTI_CLASS` or `MULTI_LABEL`.
     * `MULTI_CLASS` is also known as &#34;Single Label&#34; in the AWS Console.
     * 
     */
    @Import(name="mode")
    private @Nullable Output<String> mode;

    /**
     * @return The document classification mode.
     * One of `MULTI_CLASS` or `MULTI_LABEL`.
     * `MULTI_CLASS` is also known as &#34;Single Label&#34; in the AWS Console.
     * 
     */
    public Optional<Output<String>> mode() {
        return Optional.ofNullable(this.mode);
    }

    /**
     * KMS Key used to encrypt trained Document Classifiers.
     * Can be a KMS Key ID or a KMS Key ARN.
     * 
     */
    @Import(name="modelKmsKeyId")
    private @Nullable Output<String> modelKmsKeyId;

    /**
     * @return KMS Key used to encrypt trained Document Classifiers.
     * Can be a KMS Key ID or a KMS Key ARN.
     * 
     */
    public Optional<Output<String>> modelKmsKeyId() {
        return Optional.ofNullable(this.modelKmsKeyId);
    }

    /**
     * Name for the Document Classifier.
     * Has a maximum length of 63 characters.
     * Can contain upper- and lower-case letters, numbers, and hypen (`-`).
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name for the Document Classifier.
     * Has a maximum length of 63 characters.
     * Can contain upper- and lower-case letters, numbers, and hypen (`-`).
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Configuration for the output results of training.
     * See the `output_data_config` Configuration Block section below.
     * 
     */
    @Import(name="outputDataConfig")
    private @Nullable Output<DocumentClassifierOutputDataConfigArgs> outputDataConfig;

    /**
     * @return Configuration for the output results of training.
     * See the `output_data_config` Configuration Block section below.
     * 
     */
    public Optional<Output<DocumentClassifierOutputDataConfigArgs>> outputDataConfig() {
        return Optional.ofNullable(this.outputDataConfig);
    }

    /**
     * A map of tags to assign to the resource. If configured with a provider `default_tags` Configuration Block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the resource. If configured with a provider `default_tags` Configuration Block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * Name for the version of the Document Classifier.
     * Each version must have a unique name within the Document Classifier.
     * If omitted, the provider will assign a random, unique version name.
     * If explicitly set to `&#34;&#34;`, no version name will be set.
     * Has a maximum length of 63 characters.
     * Can contain upper- and lower-case letters, numbers, and hypen (`-`).
     * Conflicts with `version_name_prefix`.
     * 
     */
    @Import(name="versionName")
    private @Nullable Output<String> versionName;

    /**
     * @return Name for the version of the Document Classifier.
     * Each version must have a unique name within the Document Classifier.
     * If omitted, the provider will assign a random, unique version name.
     * If explicitly set to `&#34;&#34;`, no version name will be set.
     * Has a maximum length of 63 characters.
     * Can contain upper- and lower-case letters, numbers, and hypen (`-`).
     * Conflicts with `version_name_prefix`.
     * 
     */
    public Optional<Output<String>> versionName() {
        return Optional.ofNullable(this.versionName);
    }

    /**
     * Creates a unique version name beginning with the specified prefix.
     * Has a maximum length of 37 characters.
     * Can contain upper- and lower-case letters, numbers, and hypen (`-`).
     * Conflicts with `version_name`.
     * 
     */
    @Import(name="versionNamePrefix")
    private @Nullable Output<String> versionNamePrefix;

    /**
     * @return Creates a unique version name beginning with the specified prefix.
     * Has a maximum length of 37 characters.
     * Can contain upper- and lower-case letters, numbers, and hypen (`-`).
     * Conflicts with `version_name`.
     * 
     */
    public Optional<Output<String>> versionNamePrefix() {
        return Optional.ofNullable(this.versionNamePrefix);
    }

    /**
     * KMS Key used to encrypt storage volumes during job processing.
     * Can be a KMS Key ID or a KMS Key ARN.
     * 
     */
    @Import(name="volumeKmsKeyId")
    private @Nullable Output<String> volumeKmsKeyId;

    /**
     * @return KMS Key used to encrypt storage volumes during job processing.
     * Can be a KMS Key ID or a KMS Key ARN.
     * 
     */
    public Optional<Output<String>> volumeKmsKeyId() {
        return Optional.ofNullable(this.volumeKmsKeyId);
    }

    /**
     * Configuration parameters for VPC to contain Document Classifier resources.
     * See the `vpc_config` Configuration Block section below.
     * 
     */
    @Import(name="vpcConfig")
    private @Nullable Output<DocumentClassifierVpcConfigArgs> vpcConfig;

    /**
     * @return Configuration parameters for VPC to contain Document Classifier resources.
     * See the `vpc_config` Configuration Block section below.
     * 
     */
    public Optional<Output<DocumentClassifierVpcConfigArgs>> vpcConfig() {
        return Optional.ofNullable(this.vpcConfig);
    }

    private DocumentClassifierArgs() {}

    private DocumentClassifierArgs(DocumentClassifierArgs $) {
        this.dataAccessRoleArn = $.dataAccessRoleArn;
        this.inputDataConfig = $.inputDataConfig;
        this.languageCode = $.languageCode;
        this.mode = $.mode;
        this.modelKmsKeyId = $.modelKmsKeyId;
        this.name = $.name;
        this.outputDataConfig = $.outputDataConfig;
        this.tags = $.tags;
        this.versionName = $.versionName;
        this.versionNamePrefix = $.versionNamePrefix;
        this.volumeKmsKeyId = $.volumeKmsKeyId;
        this.vpcConfig = $.vpcConfig;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DocumentClassifierArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DocumentClassifierArgs $;

        public Builder() {
            $ = new DocumentClassifierArgs();
        }

        public Builder(DocumentClassifierArgs defaults) {
            $ = new DocumentClassifierArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param dataAccessRoleArn The ARN for an IAM Role which allows Comprehend to read the training and testing data.
         * 
         * @return builder
         * 
         */
        public Builder dataAccessRoleArn(Output<String> dataAccessRoleArn) {
            $.dataAccessRoleArn = dataAccessRoleArn;
            return this;
        }

        /**
         * @param dataAccessRoleArn The ARN for an IAM Role which allows Comprehend to read the training and testing data.
         * 
         * @return builder
         * 
         */
        public Builder dataAccessRoleArn(String dataAccessRoleArn) {
            return dataAccessRoleArn(Output.of(dataAccessRoleArn));
        }

        /**
         * @param inputDataConfig Configuration for the training and testing data.
         * See the `input_data_config` Configuration Block section below.
         * 
         * @return builder
         * 
         */
        public Builder inputDataConfig(Output<DocumentClassifierInputDataConfigArgs> inputDataConfig) {
            $.inputDataConfig = inputDataConfig;
            return this;
        }

        /**
         * @param inputDataConfig Configuration for the training and testing data.
         * See the `input_data_config` Configuration Block section below.
         * 
         * @return builder
         * 
         */
        public Builder inputDataConfig(DocumentClassifierInputDataConfigArgs inputDataConfig) {
            return inputDataConfig(Output.of(inputDataConfig));
        }

        /**
         * @param languageCode Two-letter language code for the language.
         * One of `en`, `es`, `fr`, `it`, `de`, or `pt`.
         * 
         * @return builder
         * 
         */
        public Builder languageCode(Output<String> languageCode) {
            $.languageCode = languageCode;
            return this;
        }

        /**
         * @param languageCode Two-letter language code for the language.
         * One of `en`, `es`, `fr`, `it`, `de`, or `pt`.
         * 
         * @return builder
         * 
         */
        public Builder languageCode(String languageCode) {
            return languageCode(Output.of(languageCode));
        }

        /**
         * @param mode The document classification mode.
         * One of `MULTI_CLASS` or `MULTI_LABEL`.
         * `MULTI_CLASS` is also known as &#34;Single Label&#34; in the AWS Console.
         * 
         * @return builder
         * 
         */
        public Builder mode(@Nullable Output<String> mode) {
            $.mode = mode;
            return this;
        }

        /**
         * @param mode The document classification mode.
         * One of `MULTI_CLASS` or `MULTI_LABEL`.
         * `MULTI_CLASS` is also known as &#34;Single Label&#34; in the AWS Console.
         * 
         * @return builder
         * 
         */
        public Builder mode(String mode) {
            return mode(Output.of(mode));
        }

        /**
         * @param modelKmsKeyId KMS Key used to encrypt trained Document Classifiers.
         * Can be a KMS Key ID or a KMS Key ARN.
         * 
         * @return builder
         * 
         */
        public Builder modelKmsKeyId(@Nullable Output<String> modelKmsKeyId) {
            $.modelKmsKeyId = modelKmsKeyId;
            return this;
        }

        /**
         * @param modelKmsKeyId KMS Key used to encrypt trained Document Classifiers.
         * Can be a KMS Key ID or a KMS Key ARN.
         * 
         * @return builder
         * 
         */
        public Builder modelKmsKeyId(String modelKmsKeyId) {
            return modelKmsKeyId(Output.of(modelKmsKeyId));
        }

        /**
         * @param name Name for the Document Classifier.
         * Has a maximum length of 63 characters.
         * Can contain upper- and lower-case letters, numbers, and hypen (`-`).
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name for the Document Classifier.
         * Has a maximum length of 63 characters.
         * Can contain upper- and lower-case letters, numbers, and hypen (`-`).
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param outputDataConfig Configuration for the output results of training.
         * See the `output_data_config` Configuration Block section below.
         * 
         * @return builder
         * 
         */
        public Builder outputDataConfig(@Nullable Output<DocumentClassifierOutputDataConfigArgs> outputDataConfig) {
            $.outputDataConfig = outputDataConfig;
            return this;
        }

        /**
         * @param outputDataConfig Configuration for the output results of training.
         * See the `output_data_config` Configuration Block section below.
         * 
         * @return builder
         * 
         */
        public Builder outputDataConfig(DocumentClassifierOutputDataConfigArgs outputDataConfig) {
            return outputDataConfig(Output.of(outputDataConfig));
        }

        /**
         * @param tags A map of tags to assign to the resource. If configured with a provider `default_tags` Configuration Block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A map of tags to assign to the resource. If configured with a provider `default_tags` Configuration Block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param versionName Name for the version of the Document Classifier.
         * Each version must have a unique name within the Document Classifier.
         * If omitted, the provider will assign a random, unique version name.
         * If explicitly set to `&#34;&#34;`, no version name will be set.
         * Has a maximum length of 63 characters.
         * Can contain upper- and lower-case letters, numbers, and hypen (`-`).
         * Conflicts with `version_name_prefix`.
         * 
         * @return builder
         * 
         */
        public Builder versionName(@Nullable Output<String> versionName) {
            $.versionName = versionName;
            return this;
        }

        /**
         * @param versionName Name for the version of the Document Classifier.
         * Each version must have a unique name within the Document Classifier.
         * If omitted, the provider will assign a random, unique version name.
         * If explicitly set to `&#34;&#34;`, no version name will be set.
         * Has a maximum length of 63 characters.
         * Can contain upper- and lower-case letters, numbers, and hypen (`-`).
         * Conflicts with `version_name_prefix`.
         * 
         * @return builder
         * 
         */
        public Builder versionName(String versionName) {
            return versionName(Output.of(versionName));
        }

        /**
         * @param versionNamePrefix Creates a unique version name beginning with the specified prefix.
         * Has a maximum length of 37 characters.
         * Can contain upper- and lower-case letters, numbers, and hypen (`-`).
         * Conflicts with `version_name`.
         * 
         * @return builder
         * 
         */
        public Builder versionNamePrefix(@Nullable Output<String> versionNamePrefix) {
            $.versionNamePrefix = versionNamePrefix;
            return this;
        }

        /**
         * @param versionNamePrefix Creates a unique version name beginning with the specified prefix.
         * Has a maximum length of 37 characters.
         * Can contain upper- and lower-case letters, numbers, and hypen (`-`).
         * Conflicts with `version_name`.
         * 
         * @return builder
         * 
         */
        public Builder versionNamePrefix(String versionNamePrefix) {
            return versionNamePrefix(Output.of(versionNamePrefix));
        }

        /**
         * @param volumeKmsKeyId KMS Key used to encrypt storage volumes during job processing.
         * Can be a KMS Key ID or a KMS Key ARN.
         * 
         * @return builder
         * 
         */
        public Builder volumeKmsKeyId(@Nullable Output<String> volumeKmsKeyId) {
            $.volumeKmsKeyId = volumeKmsKeyId;
            return this;
        }

        /**
         * @param volumeKmsKeyId KMS Key used to encrypt storage volumes during job processing.
         * Can be a KMS Key ID or a KMS Key ARN.
         * 
         * @return builder
         * 
         */
        public Builder volumeKmsKeyId(String volumeKmsKeyId) {
            return volumeKmsKeyId(Output.of(volumeKmsKeyId));
        }

        /**
         * @param vpcConfig Configuration parameters for VPC to contain Document Classifier resources.
         * See the `vpc_config` Configuration Block section below.
         * 
         * @return builder
         * 
         */
        public Builder vpcConfig(@Nullable Output<DocumentClassifierVpcConfigArgs> vpcConfig) {
            $.vpcConfig = vpcConfig;
            return this;
        }

        /**
         * @param vpcConfig Configuration parameters for VPC to contain Document Classifier resources.
         * See the `vpc_config` Configuration Block section below.
         * 
         * @return builder
         * 
         */
        public Builder vpcConfig(DocumentClassifierVpcConfigArgs vpcConfig) {
            return vpcConfig(Output.of(vpcConfig));
        }

        public DocumentClassifierArgs build() {
            $.dataAccessRoleArn = Objects.requireNonNull($.dataAccessRoleArn, "expected parameter 'dataAccessRoleArn' to be non-null");
            $.inputDataConfig = Objects.requireNonNull($.inputDataConfig, "expected parameter 'inputDataConfig' to be non-null");
            $.languageCode = Objects.requireNonNull($.languageCode, "expected parameter 'languageCode' to be non-null");
            return $;
        }
    }

}
