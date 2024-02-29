// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.elasticache.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class GetSubnetGroupResult {
    /**
     * @return ARN of the subnet group.
     * 
     */
    private String arn;
    /**
     * @return Description of the subnet group.
     * 
     */
    private String description;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String name;
    /**
     * @return Set of VPC Subnet ID-s of the subnet group.
     * 
     */
    private List<String> subnetIds;
    /**
     * @return Map of tags assigned to the subnet group.
     * 
     */
    private @Nullable Map<String,String> tags;
    /**
     * @return The Amazon Virtual Private Cloud identifier (VPC ID) of the cache subnet group.
     * 
     */
    private String vpcId;

    private GetSubnetGroupResult() {}
    /**
     * @return ARN of the subnet group.
     * 
     */
    public String arn() {
        return this.arn;
    }
    /**
     * @return Description of the subnet group.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String name() {
        return this.name;
    }
    /**
     * @return Set of VPC Subnet ID-s of the subnet group.
     * 
     */
    public List<String> subnetIds() {
        return this.subnetIds;
    }
    /**
     * @return Map of tags assigned to the subnet group.
     * 
     */
    public Map<String,String> tags() {
        return this.tags == null ? Map.of() : this.tags;
    }
    /**
     * @return The Amazon Virtual Private Cloud identifier (VPC ID) of the cache subnet group.
     * 
     */
    public String vpcId() {
        return this.vpcId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetSubnetGroupResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String arn;
        private String description;
        private String id;
        private String name;
        private List<String> subnetIds;
        private @Nullable Map<String,String> tags;
        private String vpcId;
        public Builder() {}
        public Builder(GetSubnetGroupResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.arn = defaults.arn;
    	      this.description = defaults.description;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.subnetIds = defaults.subnetIds;
    	      this.tags = defaults.tags;
    	      this.vpcId = defaults.vpcId;
        }

        @CustomType.Setter
        public Builder arn(String arn) {
            if (arn == null) {
              throw new MissingRequiredPropertyException("GetSubnetGroupResult", "arn");
            }
            this.arn = arn;
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            if (description == null) {
              throw new MissingRequiredPropertyException("GetSubnetGroupResult", "description");
            }
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetSubnetGroupResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetSubnetGroupResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder subnetIds(List<String> subnetIds) {
            if (subnetIds == null) {
              throw new MissingRequiredPropertyException("GetSubnetGroupResult", "subnetIds");
            }
            this.subnetIds = subnetIds;
            return this;
        }
        public Builder subnetIds(String... subnetIds) {
            return subnetIds(List.of(subnetIds));
        }
        @CustomType.Setter
        public Builder tags(@Nullable Map<String,String> tags) {

            this.tags = tags;
            return this;
        }
        @CustomType.Setter
        public Builder vpcId(String vpcId) {
            if (vpcId == null) {
              throw new MissingRequiredPropertyException("GetSubnetGroupResult", "vpcId");
            }
            this.vpcId = vpcId;
            return this;
        }
        public GetSubnetGroupResult build() {
            final var _resultValue = new GetSubnetGroupResult();
            _resultValue.arn = arn;
            _resultValue.description = description;
            _resultValue.id = id;
            _resultValue.name = name;
            _resultValue.subnetIds = subnetIds;
            _resultValue.tags = tags;
            _resultValue.vpcId = vpcId;
            return _resultValue;
        }
    }
}
