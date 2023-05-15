// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ec2;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class NetworkInsightsAnalysisArgs extends com.pulumi.resources.ResourceArgs {

    public static final NetworkInsightsAnalysisArgs Empty = new NetworkInsightsAnalysisArgs();

    /**
     * A list of ARNs for resources the path must traverse.
     * 
     */
    @Import(name="filterInArns")
    private @Nullable Output<List<String>> filterInArns;

    /**
     * @return A list of ARNs for resources the path must traverse.
     * 
     */
    public Optional<Output<List<String>>> filterInArns() {
        return Optional.ofNullable(this.filterInArns);
    }

    /**
     * ID of the Network Insights Path to run an analysis on.
     * 
     */
    @Import(name="networkInsightsPathId", required=true)
    private Output<String> networkInsightsPathId;

    /**
     * @return ID of the Network Insights Path to run an analysis on.
     * 
     */
    public Output<String> networkInsightsPathId() {
        return this.networkInsightsPathId;
    }

    /**
     * Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * If enabled, the resource will wait for the Network Insights Analysis status to change to `succeeded` or `failed`. Setting this to `false` will skip the process. Default: `true`.
     * 
     */
    @Import(name="waitForCompletion")
    private @Nullable Output<Boolean> waitForCompletion;

    /**
     * @return If enabled, the resource will wait for the Network Insights Analysis status to change to `succeeded` or `failed`. Setting this to `false` will skip the process. Default: `true`.
     * 
     */
    public Optional<Output<Boolean>> waitForCompletion() {
        return Optional.ofNullable(this.waitForCompletion);
    }

    private NetworkInsightsAnalysisArgs() {}

    private NetworkInsightsAnalysisArgs(NetworkInsightsAnalysisArgs $) {
        this.filterInArns = $.filterInArns;
        this.networkInsightsPathId = $.networkInsightsPathId;
        this.tags = $.tags;
        this.waitForCompletion = $.waitForCompletion;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NetworkInsightsAnalysisArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NetworkInsightsAnalysisArgs $;

        public Builder() {
            $ = new NetworkInsightsAnalysisArgs();
        }

        public Builder(NetworkInsightsAnalysisArgs defaults) {
            $ = new NetworkInsightsAnalysisArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param filterInArns A list of ARNs for resources the path must traverse.
         * 
         * @return builder
         * 
         */
        public Builder filterInArns(@Nullable Output<List<String>> filterInArns) {
            $.filterInArns = filterInArns;
            return this;
        }

        /**
         * @param filterInArns A list of ARNs for resources the path must traverse.
         * 
         * @return builder
         * 
         */
        public Builder filterInArns(List<String> filterInArns) {
            return filterInArns(Output.of(filterInArns));
        }

        /**
         * @param filterInArns A list of ARNs for resources the path must traverse.
         * 
         * @return builder
         * 
         */
        public Builder filterInArns(String... filterInArns) {
            return filterInArns(List.of(filterInArns));
        }

        /**
         * @param networkInsightsPathId ID of the Network Insights Path to run an analysis on.
         * 
         * @return builder
         * 
         */
        public Builder networkInsightsPathId(Output<String> networkInsightsPathId) {
            $.networkInsightsPathId = networkInsightsPathId;
            return this;
        }

        /**
         * @param networkInsightsPathId ID of the Network Insights Path to run an analysis on.
         * 
         * @return builder
         * 
         */
        public Builder networkInsightsPathId(String networkInsightsPathId) {
            return networkInsightsPathId(Output.of(networkInsightsPathId));
        }

        /**
         * @param tags Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param waitForCompletion If enabled, the resource will wait for the Network Insights Analysis status to change to `succeeded` or `failed`. Setting this to `false` will skip the process. Default: `true`.
         * 
         * @return builder
         * 
         */
        public Builder waitForCompletion(@Nullable Output<Boolean> waitForCompletion) {
            $.waitForCompletion = waitForCompletion;
            return this;
        }

        /**
         * @param waitForCompletion If enabled, the resource will wait for the Network Insights Analysis status to change to `succeeded` or `failed`. Setting this to `false` will skip the process. Default: `true`.
         * 
         * @return builder
         * 
         */
        public Builder waitForCompletion(Boolean waitForCompletion) {
            return waitForCompletion(Output.of(waitForCompletion));
        }

        public NetworkInsightsAnalysisArgs build() {
            $.networkInsightsPathId = Objects.requireNonNull($.networkInsightsPathId, "expected parameter 'networkInsightsPathId' to be non-null");
            return $;
        }
    }

}
