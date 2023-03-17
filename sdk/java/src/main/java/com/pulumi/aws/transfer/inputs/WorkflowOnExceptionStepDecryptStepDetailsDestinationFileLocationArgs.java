// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.transfer.inputs;

import com.pulumi.aws.transfer.inputs.WorkflowOnExceptionStepDecryptStepDetailsDestinationFileLocationEfsFileLocationArgs;
import com.pulumi.aws.transfer.inputs.WorkflowOnExceptionStepDecryptStepDetailsDestinationFileLocationS3FileLocationArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class WorkflowOnExceptionStepDecryptStepDetailsDestinationFileLocationArgs extends com.pulumi.resources.ResourceArgs {

    public static final WorkflowOnExceptionStepDecryptStepDetailsDestinationFileLocationArgs Empty = new WorkflowOnExceptionStepDecryptStepDetailsDestinationFileLocationArgs();

    /**
     * Specifies the details for the EFS file being copied.
     * 
     */
    @Import(name="efsFileLocation")
    private @Nullable Output<WorkflowOnExceptionStepDecryptStepDetailsDestinationFileLocationEfsFileLocationArgs> efsFileLocation;

    /**
     * @return Specifies the details for the EFS file being copied.
     * 
     */
    public Optional<Output<WorkflowOnExceptionStepDecryptStepDetailsDestinationFileLocationEfsFileLocationArgs>> efsFileLocation() {
        return Optional.ofNullable(this.efsFileLocation);
    }

    /**
     * Specifies the details for the S3 file being copied.
     * 
     */
    @Import(name="s3FileLocation")
    private @Nullable Output<WorkflowOnExceptionStepDecryptStepDetailsDestinationFileLocationS3FileLocationArgs> s3FileLocation;

    /**
     * @return Specifies the details for the S3 file being copied.
     * 
     */
    public Optional<Output<WorkflowOnExceptionStepDecryptStepDetailsDestinationFileLocationS3FileLocationArgs>> s3FileLocation() {
        return Optional.ofNullable(this.s3FileLocation);
    }

    private WorkflowOnExceptionStepDecryptStepDetailsDestinationFileLocationArgs() {}

    private WorkflowOnExceptionStepDecryptStepDetailsDestinationFileLocationArgs(WorkflowOnExceptionStepDecryptStepDetailsDestinationFileLocationArgs $) {
        this.efsFileLocation = $.efsFileLocation;
        this.s3FileLocation = $.s3FileLocation;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(WorkflowOnExceptionStepDecryptStepDetailsDestinationFileLocationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private WorkflowOnExceptionStepDecryptStepDetailsDestinationFileLocationArgs $;

        public Builder() {
            $ = new WorkflowOnExceptionStepDecryptStepDetailsDestinationFileLocationArgs();
        }

        public Builder(WorkflowOnExceptionStepDecryptStepDetailsDestinationFileLocationArgs defaults) {
            $ = new WorkflowOnExceptionStepDecryptStepDetailsDestinationFileLocationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param efsFileLocation Specifies the details for the EFS file being copied.
         * 
         * @return builder
         * 
         */
        public Builder efsFileLocation(@Nullable Output<WorkflowOnExceptionStepDecryptStepDetailsDestinationFileLocationEfsFileLocationArgs> efsFileLocation) {
            $.efsFileLocation = efsFileLocation;
            return this;
        }

        /**
         * @param efsFileLocation Specifies the details for the EFS file being copied.
         * 
         * @return builder
         * 
         */
        public Builder efsFileLocation(WorkflowOnExceptionStepDecryptStepDetailsDestinationFileLocationEfsFileLocationArgs efsFileLocation) {
            return efsFileLocation(Output.of(efsFileLocation));
        }

        /**
         * @param s3FileLocation Specifies the details for the S3 file being copied.
         * 
         * @return builder
         * 
         */
        public Builder s3FileLocation(@Nullable Output<WorkflowOnExceptionStepDecryptStepDetailsDestinationFileLocationS3FileLocationArgs> s3FileLocation) {
            $.s3FileLocation = s3FileLocation;
            return this;
        }

        /**
         * @param s3FileLocation Specifies the details for the S3 file being copied.
         * 
         * @return builder
         * 
         */
        public Builder s3FileLocation(WorkflowOnExceptionStepDecryptStepDetailsDestinationFileLocationS3FileLocationArgs s3FileLocation) {
            return s3FileLocation(Output.of(s3FileLocation));
        }

        public WorkflowOnExceptionStepDecryptStepDetailsDestinationFileLocationArgs build() {
            return $;
        }
    }

}
