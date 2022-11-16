// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ec2transitgateway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class MulticastDomainState extends com.pulumi.resources.ResourceArgs {

    public static final MulticastDomainState Empty = new MulticastDomainState();

    /**
     * EC2 Transit Gateway Multicast Domain Amazon Resource Name (ARN).
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return EC2 Transit Gateway Multicast Domain Amazon Resource Name (ARN).
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * Whether to automatically accept cross-account subnet associations that are associated with the EC2 Transit Gateway Multicast Domain. Valid values: `disable`, `enable`. Default value: `disable`.
     * 
     */
    @Import(name="autoAcceptSharedAssociations")
    private @Nullable Output<String> autoAcceptSharedAssociations;

    /**
     * @return Whether to automatically accept cross-account subnet associations that are associated with the EC2 Transit Gateway Multicast Domain. Valid values: `disable`, `enable`. Default value: `disable`.
     * 
     */
    public Optional<Output<String>> autoAcceptSharedAssociations() {
        return Optional.ofNullable(this.autoAcceptSharedAssociations);
    }

    /**
     * Whether to enable Internet Group Management Protocol (IGMP) version 2 for the EC2 Transit Gateway Multicast Domain. Valid values: `disable`, `enable`. Default value: `disable`.
     * 
     */
    @Import(name="igmpv2Support")
    private @Nullable Output<String> igmpv2Support;

    /**
     * @return Whether to enable Internet Group Management Protocol (IGMP) version 2 for the EC2 Transit Gateway Multicast Domain. Valid values: `disable`, `enable`. Default value: `disable`.
     * 
     */
    public Optional<Output<String>> igmpv2Support() {
        return Optional.ofNullable(this.igmpv2Support);
    }

    /**
     * Identifier of the AWS account that owns the EC2 Transit Gateway Multicast Domain.
     * 
     */
    @Import(name="ownerId")
    private @Nullable Output<String> ownerId;

    /**
     * @return Identifier of the AWS account that owns the EC2 Transit Gateway Multicast Domain.
     * 
     */
    public Optional<Output<String>> ownerId() {
        return Optional.ofNullable(this.ownerId);
    }

    /**
     * Whether to enable support for statically configuring multicast group sources for the EC2 Transit Gateway Multicast Domain. Valid values: `disable`, `enable`. Default value: `disable`.
     * 
     */
    @Import(name="staticSourcesSupport")
    private @Nullable Output<String> staticSourcesSupport;

    /**
     * @return Whether to enable support for statically configuring multicast group sources for the EC2 Transit Gateway Multicast Domain. Valid values: `disable`, `enable`. Default value: `disable`.
     * 
     */
    public Optional<Output<String>> staticSourcesSupport() {
        return Optional.ofNullable(this.staticSourcesSupport);
    }

    /**
     * Key-value tags for the EC2 Transit Gateway Multicast Domain. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Key-value tags for the EC2 Transit Gateway Multicast Domain. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
     * EC2 Transit Gateway identifier. The EC2 Transit Gateway must have `multicast_support` enabled.
     * 
     */
    @Import(name="transitGatewayId")
    private @Nullable Output<String> transitGatewayId;

    /**
     * @return EC2 Transit Gateway identifier. The EC2 Transit Gateway must have `multicast_support` enabled.
     * 
     */
    public Optional<Output<String>> transitGatewayId() {
        return Optional.ofNullable(this.transitGatewayId);
    }

    private MulticastDomainState() {}

    private MulticastDomainState(MulticastDomainState $) {
        this.arn = $.arn;
        this.autoAcceptSharedAssociations = $.autoAcceptSharedAssociations;
        this.igmpv2Support = $.igmpv2Support;
        this.ownerId = $.ownerId;
        this.staticSourcesSupport = $.staticSourcesSupport;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
        this.transitGatewayId = $.transitGatewayId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(MulticastDomainState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private MulticastDomainState $;

        public Builder() {
            $ = new MulticastDomainState();
        }

        public Builder(MulticastDomainState defaults) {
            $ = new MulticastDomainState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn EC2 Transit Gateway Multicast Domain Amazon Resource Name (ARN).
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn EC2 Transit Gateway Multicast Domain Amazon Resource Name (ARN).
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param autoAcceptSharedAssociations Whether to automatically accept cross-account subnet associations that are associated with the EC2 Transit Gateway Multicast Domain. Valid values: `disable`, `enable`. Default value: `disable`.
         * 
         * @return builder
         * 
         */
        public Builder autoAcceptSharedAssociations(@Nullable Output<String> autoAcceptSharedAssociations) {
            $.autoAcceptSharedAssociations = autoAcceptSharedAssociations;
            return this;
        }

        /**
         * @param autoAcceptSharedAssociations Whether to automatically accept cross-account subnet associations that are associated with the EC2 Transit Gateway Multicast Domain. Valid values: `disable`, `enable`. Default value: `disable`.
         * 
         * @return builder
         * 
         */
        public Builder autoAcceptSharedAssociations(String autoAcceptSharedAssociations) {
            return autoAcceptSharedAssociations(Output.of(autoAcceptSharedAssociations));
        }

        /**
         * @param igmpv2Support Whether to enable Internet Group Management Protocol (IGMP) version 2 for the EC2 Transit Gateway Multicast Domain. Valid values: `disable`, `enable`. Default value: `disable`.
         * 
         * @return builder
         * 
         */
        public Builder igmpv2Support(@Nullable Output<String> igmpv2Support) {
            $.igmpv2Support = igmpv2Support;
            return this;
        }

        /**
         * @param igmpv2Support Whether to enable Internet Group Management Protocol (IGMP) version 2 for the EC2 Transit Gateway Multicast Domain. Valid values: `disable`, `enable`. Default value: `disable`.
         * 
         * @return builder
         * 
         */
        public Builder igmpv2Support(String igmpv2Support) {
            return igmpv2Support(Output.of(igmpv2Support));
        }

        /**
         * @param ownerId Identifier of the AWS account that owns the EC2 Transit Gateway Multicast Domain.
         * 
         * @return builder
         * 
         */
        public Builder ownerId(@Nullable Output<String> ownerId) {
            $.ownerId = ownerId;
            return this;
        }

        /**
         * @param ownerId Identifier of the AWS account that owns the EC2 Transit Gateway Multicast Domain.
         * 
         * @return builder
         * 
         */
        public Builder ownerId(String ownerId) {
            return ownerId(Output.of(ownerId));
        }

        /**
         * @param staticSourcesSupport Whether to enable support for statically configuring multicast group sources for the EC2 Transit Gateway Multicast Domain. Valid values: `disable`, `enable`. Default value: `disable`.
         * 
         * @return builder
         * 
         */
        public Builder staticSourcesSupport(@Nullable Output<String> staticSourcesSupport) {
            $.staticSourcesSupport = staticSourcesSupport;
            return this;
        }

        /**
         * @param staticSourcesSupport Whether to enable support for statically configuring multicast group sources for the EC2 Transit Gateway Multicast Domain. Valid values: `disable`, `enable`. Default value: `disable`.
         * 
         * @return builder
         * 
         */
        public Builder staticSourcesSupport(String staticSourcesSupport) {
            return staticSourcesSupport(Output.of(staticSourcesSupport));
        }

        /**
         * @param tags Key-value tags for the EC2 Transit Gateway Multicast Domain. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Key-value tags for the EC2 Transit Gateway Multicast Domain. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
         * @param transitGatewayId EC2 Transit Gateway identifier. The EC2 Transit Gateway must have `multicast_support` enabled.
         * 
         * @return builder
         * 
         */
        public Builder transitGatewayId(@Nullable Output<String> transitGatewayId) {
            $.transitGatewayId = transitGatewayId;
            return this;
        }

        /**
         * @param transitGatewayId EC2 Transit Gateway identifier. The EC2 Transit Gateway must have `multicast_support` enabled.
         * 
         * @return builder
         * 
         */
        public Builder transitGatewayId(String transitGatewayId) {
            return transitGatewayId(Output.of(transitGatewayId));
        }

        public MulticastDomainState build() {
            return $;
        }
    }

}
