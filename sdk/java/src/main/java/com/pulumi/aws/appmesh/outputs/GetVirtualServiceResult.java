// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.appmesh.outputs;

import com.pulumi.aws.appmesh.outputs.GetVirtualServiceSpec;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetVirtualServiceResult {
    /**
     * @return ARN of the virtual service.
     * 
     */
    private String arn;
    /**
     * @return Creation date of the virtual service.
     * 
     */
    private String createdDate;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Last update date of the virtual service.
     * 
     */
    private String lastUpdatedDate;
    private String meshName;
    private String meshOwner;
    private String name;
    /**
     * @return Resource owner&#39;s AWS account ID.
     * 
     */
    private String resourceOwner;
    /**
     * @return Virtual service specification
     * 
     */
    private List<GetVirtualServiceSpec> specs;
    /**
     * @return Map of tags.
     * 
     */
    private Map<String,String> tags;

    private GetVirtualServiceResult() {}
    /**
     * @return ARN of the virtual service.
     * 
     */
    public String arn() {
        return this.arn;
    }
    /**
     * @return Creation date of the virtual service.
     * 
     */
    public String createdDate() {
        return this.createdDate;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Last update date of the virtual service.
     * 
     */
    public String lastUpdatedDate() {
        return this.lastUpdatedDate;
    }
    public String meshName() {
        return this.meshName;
    }
    public String meshOwner() {
        return this.meshOwner;
    }
    public String name() {
        return this.name;
    }
    /**
     * @return Resource owner&#39;s AWS account ID.
     * 
     */
    public String resourceOwner() {
        return this.resourceOwner;
    }
    /**
     * @return Virtual service specification
     * 
     */
    public List<GetVirtualServiceSpec> specs() {
        return this.specs;
    }
    /**
     * @return Map of tags.
     * 
     */
    public Map<String,String> tags() {
        return this.tags;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetVirtualServiceResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String arn;
        private String createdDate;
        private String id;
        private String lastUpdatedDate;
        private String meshName;
        private String meshOwner;
        private String name;
        private String resourceOwner;
        private List<GetVirtualServiceSpec> specs;
        private Map<String,String> tags;
        public Builder() {}
        public Builder(GetVirtualServiceResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.arn = defaults.arn;
    	      this.createdDate = defaults.createdDate;
    	      this.id = defaults.id;
    	      this.lastUpdatedDate = defaults.lastUpdatedDate;
    	      this.meshName = defaults.meshName;
    	      this.meshOwner = defaults.meshOwner;
    	      this.name = defaults.name;
    	      this.resourceOwner = defaults.resourceOwner;
    	      this.specs = defaults.specs;
    	      this.tags = defaults.tags;
        }

        @CustomType.Setter
        public Builder arn(String arn) {
            this.arn = Objects.requireNonNull(arn);
            return this;
        }
        @CustomType.Setter
        public Builder createdDate(String createdDate) {
            this.createdDate = Objects.requireNonNull(createdDate);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder lastUpdatedDate(String lastUpdatedDate) {
            this.lastUpdatedDate = Objects.requireNonNull(lastUpdatedDate);
            return this;
        }
        @CustomType.Setter
        public Builder meshName(String meshName) {
            this.meshName = Objects.requireNonNull(meshName);
            return this;
        }
        @CustomType.Setter
        public Builder meshOwner(String meshOwner) {
            this.meshOwner = Objects.requireNonNull(meshOwner);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder resourceOwner(String resourceOwner) {
            this.resourceOwner = Objects.requireNonNull(resourceOwner);
            return this;
        }
        @CustomType.Setter
        public Builder specs(List<GetVirtualServiceSpec> specs) {
            this.specs = Objects.requireNonNull(specs);
            return this;
        }
        public Builder specs(GetVirtualServiceSpec... specs) {
            return specs(List.of(specs));
        }
        @CustomType.Setter
        public Builder tags(Map<String,String> tags) {
            this.tags = Objects.requireNonNull(tags);
            return this;
        }
        public GetVirtualServiceResult build() {
            final var o = new GetVirtualServiceResult();
            o.arn = arn;
            o.createdDate = createdDate;
            o.id = id;
            o.lastUpdatedDate = lastUpdatedDate;
            o.meshName = meshName;
            o.meshOwner = meshOwner;
            o.name = name;
            o.resourceOwner = resourceOwner;
            o.specs = specs;
            o.tags = tags;
            return o;
        }
    }
}
