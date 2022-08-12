// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.kinesis.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class AnalyticsApplicationInputsSchemaRecordFormatMappingParametersCsvArgs extends com.pulumi.resources.ResourceArgs {

    public static final AnalyticsApplicationInputsSchemaRecordFormatMappingParametersCsvArgs Empty = new AnalyticsApplicationInputsSchemaRecordFormatMappingParametersCsvArgs();

    /**
     * The Column Delimiter.
     * 
     */
    @Import(name="recordColumnDelimiter", required=true)
    private Output<String> recordColumnDelimiter;

    /**
     * @return The Column Delimiter.
     * 
     */
    public Output<String> recordColumnDelimiter() {
        return this.recordColumnDelimiter;
    }

    /**
     * The Row Delimiter.
     * 
     */
    @Import(name="recordRowDelimiter", required=true)
    private Output<String> recordRowDelimiter;

    /**
     * @return The Row Delimiter.
     * 
     */
    public Output<String> recordRowDelimiter() {
        return this.recordRowDelimiter;
    }

    private AnalyticsApplicationInputsSchemaRecordFormatMappingParametersCsvArgs() {}

    private AnalyticsApplicationInputsSchemaRecordFormatMappingParametersCsvArgs(AnalyticsApplicationInputsSchemaRecordFormatMappingParametersCsvArgs $) {
        this.recordColumnDelimiter = $.recordColumnDelimiter;
        this.recordRowDelimiter = $.recordRowDelimiter;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AnalyticsApplicationInputsSchemaRecordFormatMappingParametersCsvArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AnalyticsApplicationInputsSchemaRecordFormatMappingParametersCsvArgs $;

        public Builder() {
            $ = new AnalyticsApplicationInputsSchemaRecordFormatMappingParametersCsvArgs();
        }

        public Builder(AnalyticsApplicationInputsSchemaRecordFormatMappingParametersCsvArgs defaults) {
            $ = new AnalyticsApplicationInputsSchemaRecordFormatMappingParametersCsvArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param recordColumnDelimiter The Column Delimiter.
         * 
         * @return builder
         * 
         */
        public Builder recordColumnDelimiter(Output<String> recordColumnDelimiter) {
            $.recordColumnDelimiter = recordColumnDelimiter;
            return this;
        }

        /**
         * @param recordColumnDelimiter The Column Delimiter.
         * 
         * @return builder
         * 
         */
        public Builder recordColumnDelimiter(String recordColumnDelimiter) {
            return recordColumnDelimiter(Output.of(recordColumnDelimiter));
        }

        /**
         * @param recordRowDelimiter The Row Delimiter.
         * 
         * @return builder
         * 
         */
        public Builder recordRowDelimiter(Output<String> recordRowDelimiter) {
            $.recordRowDelimiter = recordRowDelimiter;
            return this;
        }

        /**
         * @param recordRowDelimiter The Row Delimiter.
         * 
         * @return builder
         * 
         */
        public Builder recordRowDelimiter(String recordRowDelimiter) {
            return recordRowDelimiter(Output.of(recordRowDelimiter));
        }

        public AnalyticsApplicationInputsSchemaRecordFormatMappingParametersCsvArgs build() {
            $.recordColumnDelimiter = Objects.requireNonNull($.recordColumnDelimiter, "expected parameter 'recordColumnDelimiter' to be non-null");
            $.recordRowDelimiter = Objects.requireNonNull($.recordRowDelimiter, "expected parameter 'recordRowDelimiter' to be non-null");
            return $;
        }
    }

}