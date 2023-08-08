// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.wafv2.outputs;

import com.pulumi.aws.wafv2.outputs.WebAclAssociationConfigRequestBody;
import com.pulumi.core.annotations.CustomType;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class WebAclAssociationConfig {
    /**
     * @return Customizes the request body that your protected resource forward to AWS WAF for inspection. See `request_body` below for details.
     * 
     */
    private @Nullable List<WebAclAssociationConfigRequestBody> requestBodies;

    private WebAclAssociationConfig() {}
    /**
     * @return Customizes the request body that your protected resource forward to AWS WAF for inspection. See `request_body` below for details.
     * 
     */
    public List<WebAclAssociationConfigRequestBody> requestBodies() {
        return this.requestBodies == null ? List.of() : this.requestBodies;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(WebAclAssociationConfig defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<WebAclAssociationConfigRequestBody> requestBodies;
        public Builder() {}
        public Builder(WebAclAssociationConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.requestBodies = defaults.requestBodies;
        }

        @CustomType.Setter
        public Builder requestBodies(@Nullable List<WebAclAssociationConfigRequestBody> requestBodies) {
            this.requestBodies = requestBodies;
            return this;
        }
        public Builder requestBodies(WebAclAssociationConfigRequestBody... requestBodies) {
            return requestBodies(List.of(requestBodies));
        }
        public WebAclAssociationConfig build() {
            final var o = new WebAclAssociationConfig();
            o.requestBodies = requestBodies;
            return o;
        }
    }
}
