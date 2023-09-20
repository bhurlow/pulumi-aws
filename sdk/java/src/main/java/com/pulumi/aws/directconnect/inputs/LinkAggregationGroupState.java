// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.directconnect.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class LinkAggregationGroupState extends com.pulumi.resources.ResourceArgs {

    public static final LinkAggregationGroupState Empty = new LinkAggregationGroupState();

    /**
     * The ARN of the LAG.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return The ARN of the LAG.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * The ID of an existing dedicated connection to migrate to the LAG.
     * 
     */
    @Import(name="connectionId")
    private @Nullable Output<String> connectionId;

    /**
     * @return The ID of an existing dedicated connection to migrate to the LAG.
     * 
     */
    public Optional<Output<String>> connectionId() {
        return Optional.ofNullable(this.connectionId);
    }

    /**
     * The bandwidth of the individual physical connections bundled by the LAG. Valid values: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps and 100Gbps. Case sensitive.
     * 
     */
    @Import(name="connectionsBandwidth")
    private @Nullable Output<String> connectionsBandwidth;

    /**
     * @return The bandwidth of the individual physical connections bundled by the LAG. Valid values: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps and 100Gbps. Case sensitive.
     * 
     */
    public Optional<Output<String>> connectionsBandwidth() {
        return Optional.ofNullable(this.connectionsBandwidth);
    }

    /**
     * A boolean that indicates all connections associated with the LAG should be deleted so that the LAG can be destroyed without error. These objects are *not* recoverable.
     * 
     */
    @Import(name="forceDestroy")
    private @Nullable Output<Boolean> forceDestroy;

    /**
     * @return A boolean that indicates all connections associated with the LAG should be deleted so that the LAG can be destroyed without error. These objects are *not* recoverable.
     * 
     */
    public Optional<Output<Boolean>> forceDestroy() {
        return Optional.ofNullable(this.forceDestroy);
    }

    /**
     * Indicates whether the LAG supports a secondary BGP peer in the same address family (IPv4/IPv6).
     * 
     */
    @Import(name="hasLogicalRedundancy")
    private @Nullable Output<String> hasLogicalRedundancy;

    /**
     * @return Indicates whether the LAG supports a secondary BGP peer in the same address family (IPv4/IPv6).
     * 
     */
    public Optional<Output<String>> hasLogicalRedundancy() {
        return Optional.ofNullable(this.hasLogicalRedundancy);
    }

    /**
     * Indicates whether jumbo frames (9001 MTU) are supported.
     * 
     */
    @Import(name="jumboFrameCapable")
    private @Nullable Output<Boolean> jumboFrameCapable;

    /**
     * @return Indicates whether jumbo frames (9001 MTU) are supported.
     * 
     */
    public Optional<Output<Boolean>> jumboFrameCapable() {
        return Optional.ofNullable(this.jumboFrameCapable);
    }

    /**
     * The AWS Direct Connect location in which the LAG should be allocated. See [DescribeLocations](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLocations.html) for the list of AWS Direct Connect locations. Use `locationCode`.
     * 
     */
    @Import(name="location")
    private @Nullable Output<String> location;

    /**
     * @return The AWS Direct Connect location in which the LAG should be allocated. See [DescribeLocations](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLocations.html) for the list of AWS Direct Connect locations. Use `locationCode`.
     * 
     */
    public Optional<Output<String>> location() {
        return Optional.ofNullable(this.location);
    }

    /**
     * The name of the LAG.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the LAG.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The ID of the AWS account that owns the LAG.
     * 
     */
    @Import(name="ownerAccountId")
    private @Nullable Output<String> ownerAccountId;

    /**
     * @return The ID of the AWS account that owns the LAG.
     * 
     */
    public Optional<Output<String>> ownerAccountId() {
        return Optional.ofNullable(this.ownerAccountId);
    }

    /**
     * The name of the service provider associated with the LAG.
     * 
     */
    @Import(name="providerName")
    private @Nullable Output<String> providerName;

    /**
     * @return The name of the service provider associated with the LAG.
     * 
     */
    public Optional<Output<String>> providerName() {
        return Optional.ofNullable(this.providerName);
    }

    /**
     * A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

    private LinkAggregationGroupState() {}

    private LinkAggregationGroupState(LinkAggregationGroupState $) {
        this.arn = $.arn;
        this.connectionId = $.connectionId;
        this.connectionsBandwidth = $.connectionsBandwidth;
        this.forceDestroy = $.forceDestroy;
        this.hasLogicalRedundancy = $.hasLogicalRedundancy;
        this.jumboFrameCapable = $.jumboFrameCapable;
        this.location = $.location;
        this.name = $.name;
        this.ownerAccountId = $.ownerAccountId;
        this.providerName = $.providerName;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LinkAggregationGroupState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LinkAggregationGroupState $;

        public Builder() {
            $ = new LinkAggregationGroupState();
        }

        public Builder(LinkAggregationGroupState defaults) {
            $ = new LinkAggregationGroupState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn The ARN of the LAG.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn The ARN of the LAG.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param connectionId The ID of an existing dedicated connection to migrate to the LAG.
         * 
         * @return builder
         * 
         */
        public Builder connectionId(@Nullable Output<String> connectionId) {
            $.connectionId = connectionId;
            return this;
        }

        /**
         * @param connectionId The ID of an existing dedicated connection to migrate to the LAG.
         * 
         * @return builder
         * 
         */
        public Builder connectionId(String connectionId) {
            return connectionId(Output.of(connectionId));
        }

        /**
         * @param connectionsBandwidth The bandwidth of the individual physical connections bundled by the LAG. Valid values: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps and 100Gbps. Case sensitive.
         * 
         * @return builder
         * 
         */
        public Builder connectionsBandwidth(@Nullable Output<String> connectionsBandwidth) {
            $.connectionsBandwidth = connectionsBandwidth;
            return this;
        }

        /**
         * @param connectionsBandwidth The bandwidth of the individual physical connections bundled by the LAG. Valid values: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps and 100Gbps. Case sensitive.
         * 
         * @return builder
         * 
         */
        public Builder connectionsBandwidth(String connectionsBandwidth) {
            return connectionsBandwidth(Output.of(connectionsBandwidth));
        }

        /**
         * @param forceDestroy A boolean that indicates all connections associated with the LAG should be deleted so that the LAG can be destroyed without error. These objects are *not* recoverable.
         * 
         * @return builder
         * 
         */
        public Builder forceDestroy(@Nullable Output<Boolean> forceDestroy) {
            $.forceDestroy = forceDestroy;
            return this;
        }

        /**
         * @param forceDestroy A boolean that indicates all connections associated with the LAG should be deleted so that the LAG can be destroyed without error. These objects are *not* recoverable.
         * 
         * @return builder
         * 
         */
        public Builder forceDestroy(Boolean forceDestroy) {
            return forceDestroy(Output.of(forceDestroy));
        }

        /**
         * @param hasLogicalRedundancy Indicates whether the LAG supports a secondary BGP peer in the same address family (IPv4/IPv6).
         * 
         * @return builder
         * 
         */
        public Builder hasLogicalRedundancy(@Nullable Output<String> hasLogicalRedundancy) {
            $.hasLogicalRedundancy = hasLogicalRedundancy;
            return this;
        }

        /**
         * @param hasLogicalRedundancy Indicates whether the LAG supports a secondary BGP peer in the same address family (IPv4/IPv6).
         * 
         * @return builder
         * 
         */
        public Builder hasLogicalRedundancy(String hasLogicalRedundancy) {
            return hasLogicalRedundancy(Output.of(hasLogicalRedundancy));
        }

        /**
         * @param jumboFrameCapable Indicates whether jumbo frames (9001 MTU) are supported.
         * 
         * @return builder
         * 
         */
        public Builder jumboFrameCapable(@Nullable Output<Boolean> jumboFrameCapable) {
            $.jumboFrameCapable = jumboFrameCapable;
            return this;
        }

        /**
         * @param jumboFrameCapable Indicates whether jumbo frames (9001 MTU) are supported.
         * 
         * @return builder
         * 
         */
        public Builder jumboFrameCapable(Boolean jumboFrameCapable) {
            return jumboFrameCapable(Output.of(jumboFrameCapable));
        }

        /**
         * @param location The AWS Direct Connect location in which the LAG should be allocated. See [DescribeLocations](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLocations.html) for the list of AWS Direct Connect locations. Use `locationCode`.
         * 
         * @return builder
         * 
         */
        public Builder location(@Nullable Output<String> location) {
            $.location = location;
            return this;
        }

        /**
         * @param location The AWS Direct Connect location in which the LAG should be allocated. See [DescribeLocations](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLocations.html) for the list of AWS Direct Connect locations. Use `locationCode`.
         * 
         * @return builder
         * 
         */
        public Builder location(String location) {
            return location(Output.of(location));
        }

        /**
         * @param name The name of the LAG.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the LAG.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param ownerAccountId The ID of the AWS account that owns the LAG.
         * 
         * @return builder
         * 
         */
        public Builder ownerAccountId(@Nullable Output<String> ownerAccountId) {
            $.ownerAccountId = ownerAccountId;
            return this;
        }

        /**
         * @param ownerAccountId The ID of the AWS account that owns the LAG.
         * 
         * @return builder
         * 
         */
        public Builder ownerAccountId(String ownerAccountId) {
            return ownerAccountId(Output.of(ownerAccountId));
        }

        /**
         * @param providerName The name of the service provider associated with the LAG.
         * 
         * @return builder
         * 
         */
        public Builder providerName(@Nullable Output<String> providerName) {
            $.providerName = providerName;
            return this;
        }

        /**
         * @param providerName The name of the service provider associated with the LAG.
         * 
         * @return builder
         * 
         */
        public Builder providerName(String providerName) {
            return providerName(Output.of(providerName));
        }

        /**
         * @param tags A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

        public LinkAggregationGroupState build() {
            return $;
        }
    }

}
