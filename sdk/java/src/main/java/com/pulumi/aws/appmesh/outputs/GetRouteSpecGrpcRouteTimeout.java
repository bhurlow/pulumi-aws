// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.appmesh.outputs;

import com.pulumi.aws.appmesh.outputs.GetRouteSpecGrpcRouteTimeoutIdle;
import com.pulumi.aws.appmesh.outputs.GetRouteSpecGrpcRouteTimeoutPerRequest;
import com.pulumi.core.annotations.CustomType;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetRouteSpecGrpcRouteTimeout {
    private List<GetRouteSpecGrpcRouteTimeoutIdle> idles;
    private List<GetRouteSpecGrpcRouteTimeoutPerRequest> perRequests;

    private GetRouteSpecGrpcRouteTimeout() {}
    public List<GetRouteSpecGrpcRouteTimeoutIdle> idles() {
        return this.idles;
    }
    public List<GetRouteSpecGrpcRouteTimeoutPerRequest> perRequests() {
        return this.perRequests;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRouteSpecGrpcRouteTimeout defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetRouteSpecGrpcRouteTimeoutIdle> idles;
        private List<GetRouteSpecGrpcRouteTimeoutPerRequest> perRequests;
        public Builder() {}
        public Builder(GetRouteSpecGrpcRouteTimeout defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.idles = defaults.idles;
    	      this.perRequests = defaults.perRequests;
        }

        @CustomType.Setter
        public Builder idles(List<GetRouteSpecGrpcRouteTimeoutIdle> idles) {
            this.idles = Objects.requireNonNull(idles);
            return this;
        }
        public Builder idles(GetRouteSpecGrpcRouteTimeoutIdle... idles) {
            return idles(List.of(idles));
        }
        @CustomType.Setter
        public Builder perRequests(List<GetRouteSpecGrpcRouteTimeoutPerRequest> perRequests) {
            this.perRequests = Objects.requireNonNull(perRequests);
            return this;
        }
        public Builder perRequests(GetRouteSpecGrpcRouteTimeoutPerRequest... perRequests) {
            return perRequests(List.of(perRequests));
        }
        public GetRouteSpecGrpcRouteTimeout build() {
            final var o = new GetRouteSpecGrpcRouteTimeout();
            o.idles = idles;
            o.perRequests = perRequests;
            return o;
        }
    }
}
