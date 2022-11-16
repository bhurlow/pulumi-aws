// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.inspector.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ResourceGroupState extends com.pulumi.resources.ResourceArgs {

    public static final ResourceGroupState Empty = new ResourceGroupState();

    /**
     * The resource group ARN.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return The resource group ARN.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * Key-value map of tags that are used to select the EC2 instances to be included in an Amazon Inspector assessment target.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Key-value map of tags that are used to select the EC2 instances to be included in an Amazon Inspector assessment target.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private ResourceGroupState() {}

    private ResourceGroupState(ResourceGroupState $) {
        this.arn = $.arn;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ResourceGroupState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ResourceGroupState $;

        public Builder() {
            $ = new ResourceGroupState();
        }

        public Builder(ResourceGroupState defaults) {
            $ = new ResourceGroupState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn The resource group ARN.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn The resource group ARN.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param tags Key-value map of tags that are used to select the EC2 instances to be included in an Amazon Inspector assessment target.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Key-value map of tags that are used to select the EC2 instances to be included in an Amazon Inspector assessment target.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        public ResourceGroupState build() {
            return $;
        }
    }

}
