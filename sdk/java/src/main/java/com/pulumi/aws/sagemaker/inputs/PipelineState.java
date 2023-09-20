// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.sagemaker.inputs;

import com.pulumi.aws.sagemaker.inputs.PipelineParallelismConfigurationArgs;
import com.pulumi.aws.sagemaker.inputs.PipelinePipelineDefinitionS3LocationArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PipelineState extends com.pulumi.resources.ResourceArgs {

    public static final PipelineState Empty = new PipelineState();

    /**
     * The Amazon Resource Name (ARN) assigned by AWS to this Pipeline.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return The Amazon Resource Name (ARN) assigned by AWS to this Pipeline.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * This is the configuration that controls the parallelism of the pipeline. If specified, it applies to all runs of this pipeline by default. see Parallelism Configuration details below.
     * 
     */
    @Import(name="parallelismConfiguration")
    private @Nullable Output<PipelineParallelismConfigurationArgs> parallelismConfiguration;

    /**
     * @return This is the configuration that controls the parallelism of the pipeline. If specified, it applies to all runs of this pipeline by default. see Parallelism Configuration details below.
     * 
     */
    public Optional<Output<PipelineParallelismConfigurationArgs>> parallelismConfiguration() {
        return Optional.ofNullable(this.parallelismConfiguration);
    }

    /**
     * The [JSON pipeline definition](https://aws-sagemaker-mlops.github.io/sagemaker-model-building-pipeline-definition-JSON-schema/) of the pipeline.
     * 
     */
    @Import(name="pipelineDefinition")
    private @Nullable Output<String> pipelineDefinition;

    /**
     * @return The [JSON pipeline definition](https://aws-sagemaker-mlops.github.io/sagemaker-model-building-pipeline-definition-JSON-schema/) of the pipeline.
     * 
     */
    public Optional<Output<String>> pipelineDefinition() {
        return Optional.ofNullable(this.pipelineDefinition);
    }

    /**
     * The location of the pipeline definition stored in Amazon S3. If specified, SageMaker will retrieve the pipeline definition from this location. see Pipeline Definition S3 Location details below.
     * 
     */
    @Import(name="pipelineDefinitionS3Location")
    private @Nullable Output<PipelinePipelineDefinitionS3LocationArgs> pipelineDefinitionS3Location;

    /**
     * @return The location of the pipeline definition stored in Amazon S3. If specified, SageMaker will retrieve the pipeline definition from this location. see Pipeline Definition S3 Location details below.
     * 
     */
    public Optional<Output<PipelinePipelineDefinitionS3LocationArgs>> pipelineDefinitionS3Location() {
        return Optional.ofNullable(this.pipelineDefinitionS3Location);
    }

    /**
     * A description of the pipeline.
     * 
     */
    @Import(name="pipelineDescription")
    private @Nullable Output<String> pipelineDescription;

    /**
     * @return A description of the pipeline.
     * 
     */
    public Optional<Output<String>> pipelineDescription() {
        return Optional.ofNullable(this.pipelineDescription);
    }

    /**
     * The display name of the pipeline.
     * 
     */
    @Import(name="pipelineDisplayName")
    private @Nullable Output<String> pipelineDisplayName;

    /**
     * @return The display name of the pipeline.
     * 
     */
    public Optional<Output<String>> pipelineDisplayName() {
        return Optional.ofNullable(this.pipelineDisplayName);
    }

    /**
     * The name of the pipeline.
     * 
     */
    @Import(name="pipelineName")
    private @Nullable Output<String> pipelineName;

    /**
     * @return The name of the pipeline.
     * 
     */
    public Optional<Output<String>> pipelineName() {
        return Optional.ofNullable(this.pipelineName);
    }

    /**
     * The name of the Pipeline (must be unique).
     * 
     */
    @Import(name="roleArn")
    private @Nullable Output<String> roleArn;

    /**
     * @return The name of the Pipeline (must be unique).
     * 
     */
    public Optional<Output<String>> roleArn() {
        return Optional.ofNullable(this.roleArn);
    }

    /**
     * A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

    private PipelineState() {}

    private PipelineState(PipelineState $) {
        this.arn = $.arn;
        this.parallelismConfiguration = $.parallelismConfiguration;
        this.pipelineDefinition = $.pipelineDefinition;
        this.pipelineDefinitionS3Location = $.pipelineDefinitionS3Location;
        this.pipelineDescription = $.pipelineDescription;
        this.pipelineDisplayName = $.pipelineDisplayName;
        this.pipelineName = $.pipelineName;
        this.roleArn = $.roleArn;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PipelineState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PipelineState $;

        public Builder() {
            $ = new PipelineState();
        }

        public Builder(PipelineState defaults) {
            $ = new PipelineState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn The Amazon Resource Name (ARN) assigned by AWS to this Pipeline.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn The Amazon Resource Name (ARN) assigned by AWS to this Pipeline.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param parallelismConfiguration This is the configuration that controls the parallelism of the pipeline. If specified, it applies to all runs of this pipeline by default. see Parallelism Configuration details below.
         * 
         * @return builder
         * 
         */
        public Builder parallelismConfiguration(@Nullable Output<PipelineParallelismConfigurationArgs> parallelismConfiguration) {
            $.parallelismConfiguration = parallelismConfiguration;
            return this;
        }

        /**
         * @param parallelismConfiguration This is the configuration that controls the parallelism of the pipeline. If specified, it applies to all runs of this pipeline by default. see Parallelism Configuration details below.
         * 
         * @return builder
         * 
         */
        public Builder parallelismConfiguration(PipelineParallelismConfigurationArgs parallelismConfiguration) {
            return parallelismConfiguration(Output.of(parallelismConfiguration));
        }

        /**
         * @param pipelineDefinition The [JSON pipeline definition](https://aws-sagemaker-mlops.github.io/sagemaker-model-building-pipeline-definition-JSON-schema/) of the pipeline.
         * 
         * @return builder
         * 
         */
        public Builder pipelineDefinition(@Nullable Output<String> pipelineDefinition) {
            $.pipelineDefinition = pipelineDefinition;
            return this;
        }

        /**
         * @param pipelineDefinition The [JSON pipeline definition](https://aws-sagemaker-mlops.github.io/sagemaker-model-building-pipeline-definition-JSON-schema/) of the pipeline.
         * 
         * @return builder
         * 
         */
        public Builder pipelineDefinition(String pipelineDefinition) {
            return pipelineDefinition(Output.of(pipelineDefinition));
        }

        /**
         * @param pipelineDefinitionS3Location The location of the pipeline definition stored in Amazon S3. If specified, SageMaker will retrieve the pipeline definition from this location. see Pipeline Definition S3 Location details below.
         * 
         * @return builder
         * 
         */
        public Builder pipelineDefinitionS3Location(@Nullable Output<PipelinePipelineDefinitionS3LocationArgs> pipelineDefinitionS3Location) {
            $.pipelineDefinitionS3Location = pipelineDefinitionS3Location;
            return this;
        }

        /**
         * @param pipelineDefinitionS3Location The location of the pipeline definition stored in Amazon S3. If specified, SageMaker will retrieve the pipeline definition from this location. see Pipeline Definition S3 Location details below.
         * 
         * @return builder
         * 
         */
        public Builder pipelineDefinitionS3Location(PipelinePipelineDefinitionS3LocationArgs pipelineDefinitionS3Location) {
            return pipelineDefinitionS3Location(Output.of(pipelineDefinitionS3Location));
        }

        /**
         * @param pipelineDescription A description of the pipeline.
         * 
         * @return builder
         * 
         */
        public Builder pipelineDescription(@Nullable Output<String> pipelineDescription) {
            $.pipelineDescription = pipelineDescription;
            return this;
        }

        /**
         * @param pipelineDescription A description of the pipeline.
         * 
         * @return builder
         * 
         */
        public Builder pipelineDescription(String pipelineDescription) {
            return pipelineDescription(Output.of(pipelineDescription));
        }

        /**
         * @param pipelineDisplayName The display name of the pipeline.
         * 
         * @return builder
         * 
         */
        public Builder pipelineDisplayName(@Nullable Output<String> pipelineDisplayName) {
            $.pipelineDisplayName = pipelineDisplayName;
            return this;
        }

        /**
         * @param pipelineDisplayName The display name of the pipeline.
         * 
         * @return builder
         * 
         */
        public Builder pipelineDisplayName(String pipelineDisplayName) {
            return pipelineDisplayName(Output.of(pipelineDisplayName));
        }

        /**
         * @param pipelineName The name of the pipeline.
         * 
         * @return builder
         * 
         */
        public Builder pipelineName(@Nullable Output<String> pipelineName) {
            $.pipelineName = pipelineName;
            return this;
        }

        /**
         * @param pipelineName The name of the pipeline.
         * 
         * @return builder
         * 
         */
        public Builder pipelineName(String pipelineName) {
            return pipelineName(Output.of(pipelineName));
        }

        /**
         * @param roleArn The name of the Pipeline (must be unique).
         * 
         * @return builder
         * 
         */
        public Builder roleArn(@Nullable Output<String> roleArn) {
            $.roleArn = roleArn;
            return this;
        }

        /**
         * @param roleArn The name of the Pipeline (must be unique).
         * 
         * @return builder
         * 
         */
        public Builder roleArn(String roleArn) {
            return roleArn(Output.of(roleArn));
        }

        /**
         * @param tags A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

        public PipelineState build() {
            return $;
        }
    }

}
