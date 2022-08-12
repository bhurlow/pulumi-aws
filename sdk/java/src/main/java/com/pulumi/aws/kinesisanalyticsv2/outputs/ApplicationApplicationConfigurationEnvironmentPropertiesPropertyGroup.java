// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.kinesisanalyticsv2.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class ApplicationApplicationConfigurationEnvironmentPropertiesPropertyGroup {
    /**
     * @return The key of the application execution property key-value map.
     * 
     */
    private final String propertyGroupId;
    /**
     * @return Application execution property key-value map.
     * 
     */
    private final Map<String,String> propertyMap;

    @CustomType.Constructor
    private ApplicationApplicationConfigurationEnvironmentPropertiesPropertyGroup(
        @CustomType.Parameter("propertyGroupId") String propertyGroupId,
        @CustomType.Parameter("propertyMap") Map<String,String> propertyMap) {
        this.propertyGroupId = propertyGroupId;
        this.propertyMap = propertyMap;
    }

    /**
     * @return The key of the application execution property key-value map.
     * 
     */
    public String propertyGroupId() {
        return this.propertyGroupId;
    }
    /**
     * @return Application execution property key-value map.
     * 
     */
    public Map<String,String> propertyMap() {
        return this.propertyMap;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ApplicationApplicationConfigurationEnvironmentPropertiesPropertyGroup defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private String propertyGroupId;
        private Map<String,String> propertyMap;

        public Builder() {
    	      // Empty
        }

        public Builder(ApplicationApplicationConfigurationEnvironmentPropertiesPropertyGroup defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.propertyGroupId = defaults.propertyGroupId;
    	      this.propertyMap = defaults.propertyMap;
        }

        public Builder propertyGroupId(String propertyGroupId) {
            this.propertyGroupId = Objects.requireNonNull(propertyGroupId);
            return this;
        }
        public Builder propertyMap(Map<String,String> propertyMap) {
            this.propertyMap = Objects.requireNonNull(propertyMap);
            return this;
        }        public ApplicationApplicationConfigurationEnvironmentPropertiesPropertyGroup build() {
            return new ApplicationApplicationConfigurationEnvironmentPropertiesPropertyGroup(propertyGroupId, propertyMap);
        }
    }
}