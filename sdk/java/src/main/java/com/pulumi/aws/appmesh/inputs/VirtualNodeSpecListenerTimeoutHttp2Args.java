// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.appmesh.inputs;

import com.pulumi.aws.appmesh.inputs.VirtualNodeSpecListenerTimeoutHttp2IdleArgs;
import com.pulumi.aws.appmesh.inputs.VirtualNodeSpecListenerTimeoutHttp2PerRequestArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VirtualNodeSpecListenerTimeoutHttp2Args extends com.pulumi.resources.ResourceArgs {

    public static final VirtualNodeSpecListenerTimeoutHttp2Args Empty = new VirtualNodeSpecListenerTimeoutHttp2Args();

    /**
     * The idle timeout. An idle timeout bounds the amount of time that a connection may be idle.
     * 
     */
    @Import(name="idle")
    private @Nullable Output<VirtualNodeSpecListenerTimeoutHttp2IdleArgs> idle;

    /**
     * @return The idle timeout. An idle timeout bounds the amount of time that a connection may be idle.
     * 
     */
    public Optional<Output<VirtualNodeSpecListenerTimeoutHttp2IdleArgs>> idle() {
        return Optional.ofNullable(this.idle);
    }

    /**
     * The per request timeout.
     * 
     */
    @Import(name="perRequest")
    private @Nullable Output<VirtualNodeSpecListenerTimeoutHttp2PerRequestArgs> perRequest;

    /**
     * @return The per request timeout.
     * 
     */
    public Optional<Output<VirtualNodeSpecListenerTimeoutHttp2PerRequestArgs>> perRequest() {
        return Optional.ofNullable(this.perRequest);
    }

    private VirtualNodeSpecListenerTimeoutHttp2Args() {}

    private VirtualNodeSpecListenerTimeoutHttp2Args(VirtualNodeSpecListenerTimeoutHttp2Args $) {
        this.idle = $.idle;
        this.perRequest = $.perRequest;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VirtualNodeSpecListenerTimeoutHttp2Args defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VirtualNodeSpecListenerTimeoutHttp2Args $;

        public Builder() {
            $ = new VirtualNodeSpecListenerTimeoutHttp2Args();
        }

        public Builder(VirtualNodeSpecListenerTimeoutHttp2Args defaults) {
            $ = new VirtualNodeSpecListenerTimeoutHttp2Args(Objects.requireNonNull(defaults));
        }

        /**
         * @param idle The idle timeout. An idle timeout bounds the amount of time that a connection may be idle.
         * 
         * @return builder
         * 
         */
        public Builder idle(@Nullable Output<VirtualNodeSpecListenerTimeoutHttp2IdleArgs> idle) {
            $.idle = idle;
            return this;
        }

        /**
         * @param idle The idle timeout. An idle timeout bounds the amount of time that a connection may be idle.
         * 
         * @return builder
         * 
         */
        public Builder idle(VirtualNodeSpecListenerTimeoutHttp2IdleArgs idle) {
            return idle(Output.of(idle));
        }

        /**
         * @param perRequest The per request timeout.
         * 
         * @return builder
         * 
         */
        public Builder perRequest(@Nullable Output<VirtualNodeSpecListenerTimeoutHttp2PerRequestArgs> perRequest) {
            $.perRequest = perRequest;
            return this;
        }

        /**
         * @param perRequest The per request timeout.
         * 
         * @return builder
         * 
         */
        public Builder perRequest(VirtualNodeSpecListenerTimeoutHttp2PerRequestArgs perRequest) {
            return perRequest(Output.of(perRequest));
        }

        public VirtualNodeSpecListenerTimeoutHttp2Args build() {
            return $;
        }
    }

}