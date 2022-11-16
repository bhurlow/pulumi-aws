// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.location.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RouteCalculationState extends com.pulumi.resources.ResourceArgs {

    public static final RouteCalculationState Empty = new RouteCalculationState();

    /**
     * The Amazon Resource Name (ARN) for the Route calculator resource. Use the ARN when you specify a resource across AWS.
     * 
     */
    @Import(name="calculatorArn")
    private @Nullable Output<String> calculatorArn;

    /**
     * @return The Amazon Resource Name (ARN) for the Route calculator resource. Use the ARN when you specify a resource across AWS.
     * 
     */
    public Optional<Output<String>> calculatorArn() {
        return Optional.ofNullable(this.calculatorArn);
    }

    /**
     * The name of the route calculator resource.
     * 
     */
    @Import(name="calculatorName")
    private @Nullable Output<String> calculatorName;

    /**
     * @return The name of the route calculator resource.
     * 
     */
    public Optional<Output<String>> calculatorName() {
        return Optional.ofNullable(this.calculatorName);
    }

    /**
     * The timestamp for when the route calculator resource was created in ISO 8601 format.
     * 
     */
    @Import(name="createTime")
    private @Nullable Output<String> createTime;

    /**
     * @return The timestamp for when the route calculator resource was created in ISO 8601 format.
     * 
     */
    public Optional<Output<String>> createTime() {
        return Optional.ofNullable(this.createTime);
    }

    /**
     * Specifies the data provider of traffic and road network data.
     * 
     */
    @Import(name="dataSource")
    private @Nullable Output<String> dataSource;

    /**
     * @return Specifies the data provider of traffic and road network data.
     * 
     */
    public Optional<Output<String>> dataSource() {
        return Optional.ofNullable(this.dataSource);
    }

    /**
     * The optional description for the route calculator resource.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The optional description for the route calculator resource.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Key-value tags for the route calculator. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Key-value tags for the route calculator. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    @Import(name="tagsAll")
    private @Nullable Output<Map<String,String>> tagsAll;

    /**
     * @return A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    public Optional<Output<Map<String,String>>> tagsAll() {
        return Optional.ofNullable(this.tagsAll);
    }

    /**
     * The timestamp for when the route calculator resource was last update in ISO 8601.
     * 
     */
    @Import(name="updateTime")
    private @Nullable Output<String> updateTime;

    /**
     * @return The timestamp for when the route calculator resource was last update in ISO 8601.
     * 
     */
    public Optional<Output<String>> updateTime() {
        return Optional.ofNullable(this.updateTime);
    }

    private RouteCalculationState() {}

    private RouteCalculationState(RouteCalculationState $) {
        this.calculatorArn = $.calculatorArn;
        this.calculatorName = $.calculatorName;
        this.createTime = $.createTime;
        this.dataSource = $.dataSource;
        this.description = $.description;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
        this.updateTime = $.updateTime;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RouteCalculationState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RouteCalculationState $;

        public Builder() {
            $ = new RouteCalculationState();
        }

        public Builder(RouteCalculationState defaults) {
            $ = new RouteCalculationState(Objects.requireNonNull(defaults));
        }

        /**
         * @param calculatorArn The Amazon Resource Name (ARN) for the Route calculator resource. Use the ARN when you specify a resource across AWS.
         * 
         * @return builder
         * 
         */
        public Builder calculatorArn(@Nullable Output<String> calculatorArn) {
            $.calculatorArn = calculatorArn;
            return this;
        }

        /**
         * @param calculatorArn The Amazon Resource Name (ARN) for the Route calculator resource. Use the ARN when you specify a resource across AWS.
         * 
         * @return builder
         * 
         */
        public Builder calculatorArn(String calculatorArn) {
            return calculatorArn(Output.of(calculatorArn));
        }

        /**
         * @param calculatorName The name of the route calculator resource.
         * 
         * @return builder
         * 
         */
        public Builder calculatorName(@Nullable Output<String> calculatorName) {
            $.calculatorName = calculatorName;
            return this;
        }

        /**
         * @param calculatorName The name of the route calculator resource.
         * 
         * @return builder
         * 
         */
        public Builder calculatorName(String calculatorName) {
            return calculatorName(Output.of(calculatorName));
        }

        /**
         * @param createTime The timestamp for when the route calculator resource was created in ISO 8601 format.
         * 
         * @return builder
         * 
         */
        public Builder createTime(@Nullable Output<String> createTime) {
            $.createTime = createTime;
            return this;
        }

        /**
         * @param createTime The timestamp for when the route calculator resource was created in ISO 8601 format.
         * 
         * @return builder
         * 
         */
        public Builder createTime(String createTime) {
            return createTime(Output.of(createTime));
        }

        /**
         * @param dataSource Specifies the data provider of traffic and road network data.
         * 
         * @return builder
         * 
         */
        public Builder dataSource(@Nullable Output<String> dataSource) {
            $.dataSource = dataSource;
            return this;
        }

        /**
         * @param dataSource Specifies the data provider of traffic and road network data.
         * 
         * @return builder
         * 
         */
        public Builder dataSource(String dataSource) {
            return dataSource(Output.of(dataSource));
        }

        /**
         * @param description The optional description for the route calculator resource.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The optional description for the route calculator resource.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param tags Key-value tags for the route calculator. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Key-value tags for the route calculator. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
         */
        public Builder tagsAll(@Nullable Output<Map<String,String>> tagsAll) {
            $.tagsAll = tagsAll;
            return this;
        }

        /**
         * @param tagsAll A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
         * 
         * @return builder
         * 
         */
        public Builder tagsAll(Map<String,String> tagsAll) {
            return tagsAll(Output.of(tagsAll));
        }

        /**
         * @param updateTime The timestamp for when the route calculator resource was last update in ISO 8601.
         * 
         * @return builder
         * 
         */
        public Builder updateTime(@Nullable Output<String> updateTime) {
            $.updateTime = updateTime;
            return this;
        }

        /**
         * @param updateTime The timestamp for when the route calculator resource was last update in ISO 8601.
         * 
         * @return builder
         * 
         */
        public Builder updateTime(String updateTime) {
            return updateTime(Output.of(updateTime));
        }

        public RouteCalculationState build() {
            return $;
        }
    }

}
