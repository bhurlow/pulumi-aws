// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.networkmanager;

import com.pulumi.aws.networkmanager.inputs.VpcAttachmentOptionsArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VpcAttachmentArgs extends com.pulumi.resources.ResourceArgs {

    public static final VpcAttachmentArgs Empty = new VpcAttachmentArgs();

    /**
     * The ID of a core network for the VPC attachment.
     * 
     */
    @Import(name="coreNetworkId", required=true)
    private Output<String> coreNetworkId;

    /**
     * @return The ID of a core network for the VPC attachment.
     * 
     */
    public Output<String> coreNetworkId() {
        return this.coreNetworkId;
    }

    /**
     * Options for the VPC attachment.
     * 
     */
    @Import(name="options")
    private @Nullable Output<VpcAttachmentOptionsArgs> options;

    /**
     * @return Options for the VPC attachment.
     * 
     */
    public Optional<Output<VpcAttachmentOptionsArgs>> options() {
        return Optional.ofNullable(this.options);
    }

    /**
     * The subnet ARN of the VPC attachment.
     * 
     */
    @Import(name="subnetArns", required=true)
    private Output<List<String>> subnetArns;

    /**
     * @return The subnet ARN of the VPC attachment.
     * 
     */
    public Output<List<String>> subnetArns() {
        return this.subnetArns;
    }

    /**
     * Key-value tags for the attachment. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Key-value tags for the attachment. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The ARN of the VPC.
     * 
     */
    @Import(name="vpcArn", required=true)
    private Output<String> vpcArn;

    /**
     * @return The ARN of the VPC.
     * 
     */
    public Output<String> vpcArn() {
        return this.vpcArn;
    }

    private VpcAttachmentArgs() {}

    private VpcAttachmentArgs(VpcAttachmentArgs $) {
        this.coreNetworkId = $.coreNetworkId;
        this.options = $.options;
        this.subnetArns = $.subnetArns;
        this.tags = $.tags;
        this.vpcArn = $.vpcArn;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VpcAttachmentArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VpcAttachmentArgs $;

        public Builder() {
            $ = new VpcAttachmentArgs();
        }

        public Builder(VpcAttachmentArgs defaults) {
            $ = new VpcAttachmentArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param coreNetworkId The ID of a core network for the VPC attachment.
         * 
         * @return builder
         * 
         */
        public Builder coreNetworkId(Output<String> coreNetworkId) {
            $.coreNetworkId = coreNetworkId;
            return this;
        }

        /**
         * @param coreNetworkId The ID of a core network for the VPC attachment.
         * 
         * @return builder
         * 
         */
        public Builder coreNetworkId(String coreNetworkId) {
            return coreNetworkId(Output.of(coreNetworkId));
        }

        /**
         * @param options Options for the VPC attachment.
         * 
         * @return builder
         * 
         */
        public Builder options(@Nullable Output<VpcAttachmentOptionsArgs> options) {
            $.options = options;
            return this;
        }

        /**
         * @param options Options for the VPC attachment.
         * 
         * @return builder
         * 
         */
        public Builder options(VpcAttachmentOptionsArgs options) {
            return options(Output.of(options));
        }

        /**
         * @param subnetArns The subnet ARN of the VPC attachment.
         * 
         * @return builder
         * 
         */
        public Builder subnetArns(Output<List<String>> subnetArns) {
            $.subnetArns = subnetArns;
            return this;
        }

        /**
         * @param subnetArns The subnet ARN of the VPC attachment.
         * 
         * @return builder
         * 
         */
        public Builder subnetArns(List<String> subnetArns) {
            return subnetArns(Output.of(subnetArns));
        }

        /**
         * @param subnetArns The subnet ARN of the VPC attachment.
         * 
         * @return builder
         * 
         */
        public Builder subnetArns(String... subnetArns) {
            return subnetArns(List.of(subnetArns));
        }

        /**
         * @param tags Key-value tags for the attachment. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Key-value tags for the attachment. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param vpcArn The ARN of the VPC.
         * 
         * @return builder
         * 
         */
        public Builder vpcArn(Output<String> vpcArn) {
            $.vpcArn = vpcArn;
            return this;
        }

        /**
         * @param vpcArn The ARN of the VPC.
         * 
         * @return builder
         * 
         */
        public Builder vpcArn(String vpcArn) {
            return vpcArn(Output.of(vpcArn));
        }

        public VpcAttachmentArgs build() {
            $.coreNetworkId = Objects.requireNonNull($.coreNetworkId, "expected parameter 'coreNetworkId' to be non-null");
            $.subnetArns = Objects.requireNonNull($.subnetArns, "expected parameter 'subnetArns' to be non-null");
            $.vpcArn = Objects.requireNonNull($.vpcArn, "expected parameter 'vpcArn' to be non-null");
            return $;
        }
    }

}
