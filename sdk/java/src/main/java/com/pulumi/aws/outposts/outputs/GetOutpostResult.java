// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.outposts.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetOutpostResult {
    private final String arn;
    /**
     * @return Availability Zone name.
     * 
     */
    private final String availabilityZone;
    /**
     * @return Availability Zone identifier.
     * 
     */
    private final String availabilityZoneId;
    /**
     * @return Description.
     * 
     */
    private final String description;
    private final String id;
    private final String name;
    private final String ownerId;
    /**
     * @return Site identifier.
     * 
     */
    private final String siteId;

    @CustomType.Constructor
    private GetOutpostResult(
        @CustomType.Parameter("arn") String arn,
        @CustomType.Parameter("availabilityZone") String availabilityZone,
        @CustomType.Parameter("availabilityZoneId") String availabilityZoneId,
        @CustomType.Parameter("description") String description,
        @CustomType.Parameter("id") String id,
        @CustomType.Parameter("name") String name,
        @CustomType.Parameter("ownerId") String ownerId,
        @CustomType.Parameter("siteId") String siteId) {
        this.arn = arn;
        this.availabilityZone = availabilityZone;
        this.availabilityZoneId = availabilityZoneId;
        this.description = description;
        this.id = id;
        this.name = name;
        this.ownerId = ownerId;
        this.siteId = siteId;
    }

    public String arn() {
        return this.arn;
    }
    /**
     * @return Availability Zone name.
     * 
     */
    public String availabilityZone() {
        return this.availabilityZone;
    }
    /**
     * @return Availability Zone identifier.
     * 
     */
    public String availabilityZoneId() {
        return this.availabilityZoneId;
    }
    /**
     * @return Description.
     * 
     */
    public String description() {
        return this.description;
    }
    public String id() {
        return this.id;
    }
    public String name() {
        return this.name;
    }
    public String ownerId() {
        return this.ownerId;
    }
    /**
     * @return Site identifier.
     * 
     */
    public String siteId() {
        return this.siteId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetOutpostResult defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private String arn;
        private String availabilityZone;
        private String availabilityZoneId;
        private String description;
        private String id;
        private String name;
        private String ownerId;
        private String siteId;

        public Builder() {
    	      // Empty
        }

        public Builder(GetOutpostResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.arn = defaults.arn;
    	      this.availabilityZone = defaults.availabilityZone;
    	      this.availabilityZoneId = defaults.availabilityZoneId;
    	      this.description = defaults.description;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.ownerId = defaults.ownerId;
    	      this.siteId = defaults.siteId;
        }

        public Builder arn(String arn) {
            this.arn = Objects.requireNonNull(arn);
            return this;
        }
        public Builder availabilityZone(String availabilityZone) {
            this.availabilityZone = Objects.requireNonNull(availabilityZone);
            return this;
        }
        public Builder availabilityZoneId(String availabilityZoneId) {
            this.availabilityZoneId = Objects.requireNonNull(availabilityZoneId);
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
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        public Builder ownerId(String ownerId) {
            this.ownerId = Objects.requireNonNull(ownerId);
            return this;
        }
        public Builder siteId(String siteId) {
            this.siteId = Objects.requireNonNull(siteId);
            return this;
        }        public GetOutpostResult build() {
            return new GetOutpostResult(arn, availabilityZone, availabilityZoneId, description, id, name, ownerId, siteId);
        }
    }
}