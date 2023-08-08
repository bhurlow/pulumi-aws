// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ec2transitgateway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ConnectState extends com.pulumi.resources.ResourceArgs {

    public static final ConnectState Empty = new ConnectState();

    /**
     * The tunnel protocol. Valid values: `gre`. Default is `gre`.
     * 
     */
    @Import(name="protocol")
    private @Nullable Output<String> protocol;

    /**
     * @return The tunnel protocol. Valid values: `gre`. Default is `gre`.
     * 
     */
    public Optional<Output<String>> protocol() {
        return Optional.ofNullable(this.protocol);
    }

    /**
     * Key-value tags for the EC2 Transit Gateway Connect. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Key-value tags for the EC2 Transit Gateway Connect. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
     * Boolean whether the Connect should be associated with the EC2 Transit Gateway association default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
     * 
     */
    @Import(name="transitGatewayDefaultRouteTableAssociation")
    private @Nullable Output<Boolean> transitGatewayDefaultRouteTableAssociation;

    /**
     * @return Boolean whether the Connect should be associated with the EC2 Transit Gateway association default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
     * 
     */
    public Optional<Output<Boolean>> transitGatewayDefaultRouteTableAssociation() {
        return Optional.ofNullable(this.transitGatewayDefaultRouteTableAssociation);
    }

    /**
     * Boolean whether the Connect should propagate routes with the EC2 Transit Gateway propagation default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
     * 
     */
    @Import(name="transitGatewayDefaultRouteTablePropagation")
    private @Nullable Output<Boolean> transitGatewayDefaultRouteTablePropagation;

    /**
     * @return Boolean whether the Connect should propagate routes with the EC2 Transit Gateway propagation default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
     * 
     */
    public Optional<Output<Boolean>> transitGatewayDefaultRouteTablePropagation() {
        return Optional.ofNullable(this.transitGatewayDefaultRouteTablePropagation);
    }

    /**
     * Identifier of EC2 Transit Gateway.
     * 
     */
    @Import(name="transitGatewayId")
    private @Nullable Output<String> transitGatewayId;

    /**
     * @return Identifier of EC2 Transit Gateway.
     * 
     */
    public Optional<Output<String>> transitGatewayId() {
        return Optional.ofNullable(this.transitGatewayId);
    }

    /**
     * The underlaying VPC attachment
     * 
     */
    @Import(name="transportAttachmentId")
    private @Nullable Output<String> transportAttachmentId;

    /**
     * @return The underlaying VPC attachment
     * 
     */
    public Optional<Output<String>> transportAttachmentId() {
        return Optional.ofNullable(this.transportAttachmentId);
    }

    private ConnectState() {}

    private ConnectState(ConnectState $) {
        this.protocol = $.protocol;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
        this.transitGatewayDefaultRouteTableAssociation = $.transitGatewayDefaultRouteTableAssociation;
        this.transitGatewayDefaultRouteTablePropagation = $.transitGatewayDefaultRouteTablePropagation;
        this.transitGatewayId = $.transitGatewayId;
        this.transportAttachmentId = $.transportAttachmentId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ConnectState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ConnectState $;

        public Builder() {
            $ = new ConnectState();
        }

        public Builder(ConnectState defaults) {
            $ = new ConnectState(Objects.requireNonNull(defaults));
        }

        /**
         * @param protocol The tunnel protocol. Valid values: `gre`. Default is `gre`.
         * 
         * @return builder
         * 
         */
        public Builder protocol(@Nullable Output<String> protocol) {
            $.protocol = protocol;
            return this;
        }

        /**
         * @param protocol The tunnel protocol. Valid values: `gre`. Default is `gre`.
         * 
         * @return builder
         * 
         */
        public Builder protocol(String protocol) {
            return protocol(Output.of(protocol));
        }

        /**
         * @param tags Key-value tags for the EC2 Transit Gateway Connect. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Key-value tags for the EC2 Transit Gateway Connect. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
         * @param transitGatewayDefaultRouteTableAssociation Boolean whether the Connect should be associated with the EC2 Transit Gateway association default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
         * 
         * @return builder
         * 
         */
        public Builder transitGatewayDefaultRouteTableAssociation(@Nullable Output<Boolean> transitGatewayDefaultRouteTableAssociation) {
            $.transitGatewayDefaultRouteTableAssociation = transitGatewayDefaultRouteTableAssociation;
            return this;
        }

        /**
         * @param transitGatewayDefaultRouteTableAssociation Boolean whether the Connect should be associated with the EC2 Transit Gateway association default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
         * 
         * @return builder
         * 
         */
        public Builder transitGatewayDefaultRouteTableAssociation(Boolean transitGatewayDefaultRouteTableAssociation) {
            return transitGatewayDefaultRouteTableAssociation(Output.of(transitGatewayDefaultRouteTableAssociation));
        }

        /**
         * @param transitGatewayDefaultRouteTablePropagation Boolean whether the Connect should propagate routes with the EC2 Transit Gateway propagation default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
         * 
         * @return builder
         * 
         */
        public Builder transitGatewayDefaultRouteTablePropagation(@Nullable Output<Boolean> transitGatewayDefaultRouteTablePropagation) {
            $.transitGatewayDefaultRouteTablePropagation = transitGatewayDefaultRouteTablePropagation;
            return this;
        }

        /**
         * @param transitGatewayDefaultRouteTablePropagation Boolean whether the Connect should propagate routes with the EC2 Transit Gateway propagation default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
         * 
         * @return builder
         * 
         */
        public Builder transitGatewayDefaultRouteTablePropagation(Boolean transitGatewayDefaultRouteTablePropagation) {
            return transitGatewayDefaultRouteTablePropagation(Output.of(transitGatewayDefaultRouteTablePropagation));
        }

        /**
         * @param transitGatewayId Identifier of EC2 Transit Gateway.
         * 
         * @return builder
         * 
         */
        public Builder transitGatewayId(@Nullable Output<String> transitGatewayId) {
            $.transitGatewayId = transitGatewayId;
            return this;
        }

        /**
         * @param transitGatewayId Identifier of EC2 Transit Gateway.
         * 
         * @return builder
         * 
         */
        public Builder transitGatewayId(String transitGatewayId) {
            return transitGatewayId(Output.of(transitGatewayId));
        }

        /**
         * @param transportAttachmentId The underlaying VPC attachment
         * 
         * @return builder
         * 
         */
        public Builder transportAttachmentId(@Nullable Output<String> transportAttachmentId) {
            $.transportAttachmentId = transportAttachmentId;
            return this;
        }

        /**
         * @param transportAttachmentId The underlaying VPC attachment
         * 
         * @return builder
         * 
         */
        public Builder transportAttachmentId(String transportAttachmentId) {
            return transportAttachmentId(Output.of(transportAttachmentId));
        }

        public ConnectState build() {
            return $;
        }
    }

}
