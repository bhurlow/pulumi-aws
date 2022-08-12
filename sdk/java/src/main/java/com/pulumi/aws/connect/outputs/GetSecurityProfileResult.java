// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.connect.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetSecurityProfileResult {
    /**
     * @return The Amazon Resource Name (ARN) of the Security Profile.
     * 
     */
    private final String arn;
    /**
     * @return Specifies the description of the Security Profile.
     * 
     */
    private final String description;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private final String id;
    private final String instanceId;
    private final String name;
    /**
     * @return The organization resource identifier for the security profile.
     * 
     */
    private final String organizationResourceId;
    /**
     * @return Specifies a list of permissions assigned to the security profile.
     * 
     */
    private final List<String> permissions;
    private final String securityProfileId;
    /**
     * @return A map of tags to assign to the Security Profile.
     * 
     */
    private final Map<String,String> tags;

    @CustomType.Constructor
    private GetSecurityProfileResult(
        @CustomType.Parameter("arn") String arn,
        @CustomType.Parameter("description") String description,
        @CustomType.Parameter("id") String id,
        @CustomType.Parameter("instanceId") String instanceId,
        @CustomType.Parameter("name") String name,
        @CustomType.Parameter("organizationResourceId") String organizationResourceId,
        @CustomType.Parameter("permissions") List<String> permissions,
        @CustomType.Parameter("securityProfileId") String securityProfileId,
        @CustomType.Parameter("tags") Map<String,String> tags) {
        this.arn = arn;
        this.description = description;
        this.id = id;
        this.instanceId = instanceId;
        this.name = name;
        this.organizationResourceId = organizationResourceId;
        this.permissions = permissions;
        this.securityProfileId = securityProfileId;
        this.tags = tags;
    }

    /**
     * @return The Amazon Resource Name (ARN) of the Security Profile.
     * 
     */
    public String arn() {
        return this.arn;
    }
    /**
     * @return Specifies the description of the Security Profile.
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
    public String instanceId() {
        return this.instanceId;
    }
    public String name() {
        return this.name;
    }
    /**
     * @return The organization resource identifier for the security profile.
     * 
     */
    public String organizationResourceId() {
        return this.organizationResourceId;
    }
    /**
     * @return Specifies a list of permissions assigned to the security profile.
     * 
     */
    public List<String> permissions() {
        return this.permissions;
    }
    public String securityProfileId() {
        return this.securityProfileId;
    }
    /**
     * @return A map of tags to assign to the Security Profile.
     * 
     */
    public Map<String,String> tags() {
        return this.tags;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetSecurityProfileResult defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private String arn;
        private String description;
        private String id;
        private String instanceId;
        private String name;
        private String organizationResourceId;
        private List<String> permissions;
        private String securityProfileId;
        private Map<String,String> tags;

        public Builder() {
    	      // Empty
        }

        public Builder(GetSecurityProfileResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.arn = defaults.arn;
    	      this.description = defaults.description;
    	      this.id = defaults.id;
    	      this.instanceId = defaults.instanceId;
    	      this.name = defaults.name;
    	      this.organizationResourceId = defaults.organizationResourceId;
    	      this.permissions = defaults.permissions;
    	      this.securityProfileId = defaults.securityProfileId;
    	      this.tags = defaults.tags;
        }

        public Builder arn(String arn) {
            this.arn = Objects.requireNonNull(arn);
            return this;
        }
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        public Builder instanceId(String instanceId) {
            this.instanceId = Objects.requireNonNull(instanceId);
            return this;
        }
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        public Builder organizationResourceId(String organizationResourceId) {
            this.organizationResourceId = Objects.requireNonNull(organizationResourceId);
            return this;
        }
        public Builder permissions(List<String> permissions) {
            this.permissions = Objects.requireNonNull(permissions);
            return this;
        }
        public Builder permissions(String... permissions) {
            return permissions(List.of(permissions));
        }
        public Builder securityProfileId(String securityProfileId) {
            this.securityProfileId = Objects.requireNonNull(securityProfileId);
            return this;
        }
        public Builder tags(Map<String,String> tags) {
            this.tags = Objects.requireNonNull(tags);
            return this;
        }        public GetSecurityProfileResult build() {
            return new GetSecurityProfileResult(arn, description, id, instanceId, name, organizationResourceId, permissions, securityProfileId, tags);
        }
    }
}