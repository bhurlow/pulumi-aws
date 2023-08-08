// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.quicksight;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class IngestionArgs extends com.pulumi.resources.ResourceArgs {

    public static final IngestionArgs Empty = new IngestionArgs();

    /**
     * AWS account ID.
     * 
     */
    @Import(name="awsAccountId")
    private @Nullable Output<String> awsAccountId;

    /**
     * @return AWS account ID.
     * 
     */
    public Optional<Output<String>> awsAccountId() {
        return Optional.ofNullable(this.awsAccountId);
    }

    /**
     * ID of the dataset used in the ingestion.
     * 
     */
    @Import(name="dataSetId", required=true)
    private Output<String> dataSetId;

    /**
     * @return ID of the dataset used in the ingestion.
     * 
     */
    public Output<String> dataSetId() {
        return this.dataSetId;
    }

    /**
     * ID for the ingestion.
     * 
     */
    @Import(name="ingestionId", required=true)
    private Output<String> ingestionId;

    /**
     * @return ID for the ingestion.
     * 
     */
    public Output<String> ingestionId() {
        return this.ingestionId;
    }

    /**
     * Type of ingestion to be created. Valid values are `INCREMENTAL_REFRESH` and `FULL_REFRESH`.
     * 
     * The following arguments are optional:
     * 
     */
    @Import(name="ingestionType", required=true)
    private Output<String> ingestionType;

    /**
     * @return Type of ingestion to be created. Valid values are `INCREMENTAL_REFRESH` and `FULL_REFRESH`.
     * 
     * The following arguments are optional:
     * 
     */
    public Output<String> ingestionType() {
        return this.ingestionType;
    }

    private IngestionArgs() {}

    private IngestionArgs(IngestionArgs $) {
        this.awsAccountId = $.awsAccountId;
        this.dataSetId = $.dataSetId;
        this.ingestionId = $.ingestionId;
        this.ingestionType = $.ingestionType;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IngestionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IngestionArgs $;

        public Builder() {
            $ = new IngestionArgs();
        }

        public Builder(IngestionArgs defaults) {
            $ = new IngestionArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param awsAccountId AWS account ID.
         * 
         * @return builder
         * 
         */
        public Builder awsAccountId(@Nullable Output<String> awsAccountId) {
            $.awsAccountId = awsAccountId;
            return this;
        }

        /**
         * @param awsAccountId AWS account ID.
         * 
         * @return builder
         * 
         */
        public Builder awsAccountId(String awsAccountId) {
            return awsAccountId(Output.of(awsAccountId));
        }

        /**
         * @param dataSetId ID of the dataset used in the ingestion.
         * 
         * @return builder
         * 
         */
        public Builder dataSetId(Output<String> dataSetId) {
            $.dataSetId = dataSetId;
            return this;
        }

        /**
         * @param dataSetId ID of the dataset used in the ingestion.
         * 
         * @return builder
         * 
         */
        public Builder dataSetId(String dataSetId) {
            return dataSetId(Output.of(dataSetId));
        }

        /**
         * @param ingestionId ID for the ingestion.
         * 
         * @return builder
         * 
         */
        public Builder ingestionId(Output<String> ingestionId) {
            $.ingestionId = ingestionId;
            return this;
        }

        /**
         * @param ingestionId ID for the ingestion.
         * 
         * @return builder
         * 
         */
        public Builder ingestionId(String ingestionId) {
            return ingestionId(Output.of(ingestionId));
        }

        /**
         * @param ingestionType Type of ingestion to be created. Valid values are `INCREMENTAL_REFRESH` and `FULL_REFRESH`.
         * 
         * The following arguments are optional:
         * 
         * @return builder
         * 
         */
        public Builder ingestionType(Output<String> ingestionType) {
            $.ingestionType = ingestionType;
            return this;
        }

        /**
         * @param ingestionType Type of ingestion to be created. Valid values are `INCREMENTAL_REFRESH` and `FULL_REFRESH`.
         * 
         * The following arguments are optional:
         * 
         * @return builder
         * 
         */
        public Builder ingestionType(String ingestionType) {
            return ingestionType(Output.of(ingestionType));
        }

        public IngestionArgs build() {
            $.dataSetId = Objects.requireNonNull($.dataSetId, "expected parameter 'dataSetId' to be non-null");
            $.ingestionId = Objects.requireNonNull($.ingestionId, "expected parameter 'ingestionId' to be non-null");
            $.ingestionType = Objects.requireNonNull($.ingestionType, "expected parameter 'ingestionType' to be non-null");
            return $;
        }
    }

}
