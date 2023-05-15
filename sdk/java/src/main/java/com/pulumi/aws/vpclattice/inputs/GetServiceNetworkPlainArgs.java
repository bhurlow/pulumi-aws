// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.vpclattice.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetServiceNetworkPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetServiceNetworkPlainArgs Empty = new GetServiceNetworkPlainArgs();

    /**
     * Identifier of the network service.
     * 
     */
    @Import(name="serviceNetworkIdentifier", required=true)
    private String serviceNetworkIdentifier;

    /**
     * @return Identifier of the network service.
     * 
     */
    public String serviceNetworkIdentifier() {
        return this.serviceNetworkIdentifier;
    }

    @Import(name="tags")
    private @Nullable Map<String,String> tags;

    public Optional<Map<String,String>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private GetServiceNetworkPlainArgs() {}

    private GetServiceNetworkPlainArgs(GetServiceNetworkPlainArgs $) {
        this.serviceNetworkIdentifier = $.serviceNetworkIdentifier;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetServiceNetworkPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetServiceNetworkPlainArgs $;

        public Builder() {
            $ = new GetServiceNetworkPlainArgs();
        }

        public Builder(GetServiceNetworkPlainArgs defaults) {
            $ = new GetServiceNetworkPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param serviceNetworkIdentifier Identifier of the network service.
         * 
         * @return builder
         * 
         */
        public Builder serviceNetworkIdentifier(String serviceNetworkIdentifier) {
            $.serviceNetworkIdentifier = serviceNetworkIdentifier;
            return this;
        }

        public Builder tags(@Nullable Map<String,String> tags) {
            $.tags = tags;
            return this;
        }

        public GetServiceNetworkPlainArgs build() {
            $.serviceNetworkIdentifier = Objects.requireNonNull($.serviceNetworkIdentifier, "expected parameter 'serviceNetworkIdentifier' to be non-null");
            return $;
        }
    }

}
