// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ec2;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VpcIpamScopeArgs extends com.pulumi.resources.ResourceArgs {

    public static final VpcIpamScopeArgs Empty = new VpcIpamScopeArgs();

    /**
     * A description for the scope you&#39;re creating.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return A description for the scope you&#39;re creating.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The ID of the IPAM for which you&#39;re creating this scope.
     * 
     */
    @Import(name="ipamId", required=true)
    private Output<String> ipamId;

    /**
     * @return The ID of the IPAM for which you&#39;re creating this scope.
     * 
     */
    public Output<String> ipamId() {
        return this.ipamId;
    }

    /**
     * Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private VpcIpamScopeArgs() {}

    private VpcIpamScopeArgs(VpcIpamScopeArgs $) {
        this.description = $.description;
        this.ipamId = $.ipamId;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VpcIpamScopeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VpcIpamScopeArgs $;

        public Builder() {
            $ = new VpcIpamScopeArgs();
        }

        public Builder(VpcIpamScopeArgs defaults) {
            $ = new VpcIpamScopeArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description A description for the scope you&#39;re creating.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description A description for the scope you&#39;re creating.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param ipamId The ID of the IPAM for which you&#39;re creating this scope.
         * 
         * @return builder
         * 
         */
        public Builder ipamId(Output<String> ipamId) {
            $.ipamId = ipamId;
            return this;
        }

        /**
         * @param ipamId The ID of the IPAM for which you&#39;re creating this scope.
         * 
         * @return builder
         * 
         */
        public Builder ipamId(String ipamId) {
            return ipamId(Output.of(ipamId));
        }

        /**
         * @param tags Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        public VpcIpamScopeArgs build() {
            $.ipamId = Objects.requireNonNull($.ipamId, "expected parameter 'ipamId' to be non-null");
            return $;
        }
    }

}
