// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.dynamodb.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TableImportTableInputFormatOptionsCsvArgs extends com.pulumi.resources.ResourceArgs {

    public static final TableImportTableInputFormatOptionsCsvArgs Empty = new TableImportTableInputFormatOptionsCsvArgs();

    /**
     * The delimiter used for separating items in the CSV file being imported.
     * 
     */
    @Import(name="delimiter")
    private @Nullable Output<String> delimiter;

    /**
     * @return The delimiter used for separating items in the CSV file being imported.
     * 
     */
    public Optional<Output<String>> delimiter() {
        return Optional.ofNullable(this.delimiter);
    }

    /**
     * List of the headers used to specify a common header for all source CSV files being imported.
     * 
     */
    @Import(name="headerLists")
    private @Nullable Output<List<String>> headerLists;

    /**
     * @return List of the headers used to specify a common header for all source CSV files being imported.
     * 
     */
    public Optional<Output<List<String>>> headerLists() {
        return Optional.ofNullable(this.headerLists);
    }

    private TableImportTableInputFormatOptionsCsvArgs() {}

    private TableImportTableInputFormatOptionsCsvArgs(TableImportTableInputFormatOptionsCsvArgs $) {
        this.delimiter = $.delimiter;
        this.headerLists = $.headerLists;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TableImportTableInputFormatOptionsCsvArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TableImportTableInputFormatOptionsCsvArgs $;

        public Builder() {
            $ = new TableImportTableInputFormatOptionsCsvArgs();
        }

        public Builder(TableImportTableInputFormatOptionsCsvArgs defaults) {
            $ = new TableImportTableInputFormatOptionsCsvArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param delimiter The delimiter used for separating items in the CSV file being imported.
         * 
         * @return builder
         * 
         */
        public Builder delimiter(@Nullable Output<String> delimiter) {
            $.delimiter = delimiter;
            return this;
        }

        /**
         * @param delimiter The delimiter used for separating items in the CSV file being imported.
         * 
         * @return builder
         * 
         */
        public Builder delimiter(String delimiter) {
            return delimiter(Output.of(delimiter));
        }

        /**
         * @param headerLists List of the headers used to specify a common header for all source CSV files being imported.
         * 
         * @return builder
         * 
         */
        public Builder headerLists(@Nullable Output<List<String>> headerLists) {
            $.headerLists = headerLists;
            return this;
        }

        /**
         * @param headerLists List of the headers used to specify a common header for all source CSV files being imported.
         * 
         * @return builder
         * 
         */
        public Builder headerLists(List<String> headerLists) {
            return headerLists(Output.of(headerLists));
        }

        /**
         * @param headerLists List of the headers used to specify a common header for all source CSV files being imported.
         * 
         * @return builder
         * 
         */
        public Builder headerLists(String... headerLists) {
            return headerLists(List.of(headerLists));
        }

        public TableImportTableInputFormatOptionsCsvArgs build() {
            return $;
        }
    }

}
