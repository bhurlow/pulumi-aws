// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.sagemaker.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PipelinePipelineDefinitionS3LocationArgs extends com.pulumi.resources.ResourceArgs {

    public static final PipelinePipelineDefinitionS3LocationArgs Empty = new PipelinePipelineDefinitionS3LocationArgs();

    /**
     * Name of the S3 bucket.
     * 
     */
    @Import(name="bucket", required=true)
    private Output<String> bucket;

    /**
     * @return Name of the S3 bucket.
     * 
     */
    public Output<String> bucket() {
        return this.bucket;
    }

    /**
     * The object key (or key name) uniquely identifies the object in an S3 bucket.
     * 
     */
    @Import(name="objectKey", required=true)
    private Output<String> objectKey;

    /**
     * @return The object key (or key name) uniquely identifies the object in an S3 bucket.
     * 
     */
    public Output<String> objectKey() {
        return this.objectKey;
    }

    /**
     * Version Id of the pipeline definition file. If not specified, Amazon SageMaker will retrieve the latest version.
     * 
     */
    @Import(name="versionId")
    private @Nullable Output<String> versionId;

    /**
     * @return Version Id of the pipeline definition file. If not specified, Amazon SageMaker will retrieve the latest version.
     * 
     */
    public Optional<Output<String>> versionId() {
        return Optional.ofNullable(this.versionId);
    }

    private PipelinePipelineDefinitionS3LocationArgs() {}

    private PipelinePipelineDefinitionS3LocationArgs(PipelinePipelineDefinitionS3LocationArgs $) {
        this.bucket = $.bucket;
        this.objectKey = $.objectKey;
        this.versionId = $.versionId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PipelinePipelineDefinitionS3LocationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PipelinePipelineDefinitionS3LocationArgs $;

        public Builder() {
            $ = new PipelinePipelineDefinitionS3LocationArgs();
        }

        public Builder(PipelinePipelineDefinitionS3LocationArgs defaults) {
            $ = new PipelinePipelineDefinitionS3LocationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param bucket Name of the S3 bucket.
         * 
         * @return builder
         * 
         */
        public Builder bucket(Output<String> bucket) {
            $.bucket = bucket;
            return this;
        }

        /**
         * @param bucket Name of the S3 bucket.
         * 
         * @return builder
         * 
         */
        public Builder bucket(String bucket) {
            return bucket(Output.of(bucket));
        }

        /**
         * @param objectKey The object key (or key name) uniquely identifies the object in an S3 bucket.
         * 
         * @return builder
         * 
         */
        public Builder objectKey(Output<String> objectKey) {
            $.objectKey = objectKey;
            return this;
        }

        /**
         * @param objectKey The object key (or key name) uniquely identifies the object in an S3 bucket.
         * 
         * @return builder
         * 
         */
        public Builder objectKey(String objectKey) {
            return objectKey(Output.of(objectKey));
        }

        /**
         * @param versionId Version Id of the pipeline definition file. If not specified, Amazon SageMaker will retrieve the latest version.
         * 
         * @return builder
         * 
         */
        public Builder versionId(@Nullable Output<String> versionId) {
            $.versionId = versionId;
            return this;
        }

        /**
         * @param versionId Version Id of the pipeline definition file. If not specified, Amazon SageMaker will retrieve the latest version.
         * 
         * @return builder
         * 
         */
        public Builder versionId(String versionId) {
            return versionId(Output.of(versionId));
        }

        public PipelinePipelineDefinitionS3LocationArgs build() {
            $.bucket = Objects.requireNonNull($.bucket, "expected parameter 'bucket' to be non-null");
            $.objectKey = Objects.requireNonNull($.objectKey, "expected parameter 'objectKey' to be non-null");
            return $;
        }
    }

}
