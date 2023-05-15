// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.dataexchange;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DataSetArgs extends com.pulumi.resources.ResourceArgs {

    public static final DataSetArgs Empty = new DataSetArgs();

    /**
     * The type of asset that is added to a data set. Valid values are: `S3_SNAPSHOT`, `REDSHIFT_DATA_SHARE`, and `API_GATEWAY_API`.
     * 
     */
    @Import(name="assetType", required=true)
    private Output<String> assetType;

    /**
     * @return The type of asset that is added to a data set. Valid values are: `S3_SNAPSHOT`, `REDSHIFT_DATA_SHARE`, and `API_GATEWAY_API`.
     * 
     */
    public Output<String> assetType() {
        return this.assetType;
    }

    /**
     * A description for the data set.
     * 
     */
    @Import(name="description", required=true)
    private Output<String> description;

    /**
     * @return A description for the data set.
     * 
     */
    public Output<String> description() {
        return this.description;
    }

    /**
     * The name of the data set.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the data set.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
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

    private DataSetArgs() {}

    private DataSetArgs(DataSetArgs $) {
        this.assetType = $.assetType;
        this.description = $.description;
        this.name = $.name;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DataSetArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DataSetArgs $;

        public Builder() {
            $ = new DataSetArgs();
        }

        public Builder(DataSetArgs defaults) {
            $ = new DataSetArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param assetType The type of asset that is added to a data set. Valid values are: `S3_SNAPSHOT`, `REDSHIFT_DATA_SHARE`, and `API_GATEWAY_API`.
         * 
         * @return builder
         * 
         */
        public Builder assetType(Output<String> assetType) {
            $.assetType = assetType;
            return this;
        }

        /**
         * @param assetType The type of asset that is added to a data set. Valid values are: `S3_SNAPSHOT`, `REDSHIFT_DATA_SHARE`, and `API_GATEWAY_API`.
         * 
         * @return builder
         * 
         */
        public Builder assetType(String assetType) {
            return assetType(Output.of(assetType));
        }

        /**
         * @param description A description for the data set.
         * 
         * @return builder
         * 
         */
        public Builder description(Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description A description for the data set.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param name The name of the data set.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the data set.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
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

        public DataSetArgs build() {
            $.assetType = Objects.requireNonNull($.assetType, "expected parameter 'assetType' to be non-null");
            $.description = Objects.requireNonNull($.description, "expected parameter 'description' to be non-null");
            return $;
        }
    }

}
