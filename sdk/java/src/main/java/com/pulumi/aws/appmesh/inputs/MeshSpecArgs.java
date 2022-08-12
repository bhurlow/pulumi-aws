// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.appmesh.inputs;

import com.pulumi.aws.appmesh.inputs.MeshSpecEgressFilterArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class MeshSpecArgs extends com.pulumi.resources.ResourceArgs {

    public static final MeshSpecArgs Empty = new MeshSpecArgs();

    /**
     * The egress filter rules for the service mesh.
     * 
     */
    @Import(name="egressFilter")
    private @Nullable Output<MeshSpecEgressFilterArgs> egressFilter;

    /**
     * @return The egress filter rules for the service mesh.
     * 
     */
    public Optional<Output<MeshSpecEgressFilterArgs>> egressFilter() {
        return Optional.ofNullable(this.egressFilter);
    }

    private MeshSpecArgs() {}

    private MeshSpecArgs(MeshSpecArgs $) {
        this.egressFilter = $.egressFilter;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(MeshSpecArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private MeshSpecArgs $;

        public Builder() {
            $ = new MeshSpecArgs();
        }

        public Builder(MeshSpecArgs defaults) {
            $ = new MeshSpecArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param egressFilter The egress filter rules for the service mesh.
         * 
         * @return builder
         * 
         */
        public Builder egressFilter(@Nullable Output<MeshSpecEgressFilterArgs> egressFilter) {
            $.egressFilter = egressFilter;
            return this;
        }

        /**
         * @param egressFilter The egress filter rules for the service mesh.
         * 
         * @return builder
         * 
         */
        public Builder egressFilter(MeshSpecEgressFilterArgs egressFilter) {
            return egressFilter(Output.of(egressFilter));
        }

        public MeshSpecArgs build() {
            return $;
        }
    }

}