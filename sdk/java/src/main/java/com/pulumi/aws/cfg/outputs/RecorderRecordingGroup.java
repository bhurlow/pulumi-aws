// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.cfg.outputs;

import com.pulumi.aws.cfg.outputs.RecorderRecordingGroupExclusionByResourceType;
import com.pulumi.aws.cfg.outputs.RecorderRecordingGroupRecordingStrategy;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class RecorderRecordingGroup {
    /**
     * @return Specifies whether AWS Config records configuration changes for every supported type of regional resource (which includes any new type that will become supported in the future). Conflicts with `resource_types`. Defaults to `true`.
     * 
     */
    private @Nullable Boolean allSupported;
    /**
     * @return An object that specifies how AWS Config excludes resource types from being recorded by the configuration recorder.To use this option, you must set the useOnly field of RecordingStrategy to `EXCLUSION_BY_RESOURCE_TYPES` Requires `all_supported = false`. Conflicts with `resource_types`.
     * 
     */
    private @Nullable List<RecorderRecordingGroupExclusionByResourceType> exclusionByResourceTypes;
    /**
     * @return Specifies whether AWS Config includes all supported types of _global resources_ with the resources that it records. Requires `all_supported = true`. Conflicts with `resource_types`.
     * 
     */
    private @Nullable Boolean includeGlobalResourceTypes;
    /**
     * @return Recording Strategy - see below..
     * 
     */
    private @Nullable List<RecorderRecordingGroupRecordingStrategy> recordingStrategies;
    /**
     * @return A list that specifies the types of AWS resources for which AWS Config records configuration changes (for example, `AWS::EC2::Instance` or `AWS::CloudTrail::Trail`). See [relevant part of AWS Docs](http://docs.aws.amazon.com/config/latest/APIReference/API_ResourceIdentifier.html#config-Type-ResourceIdentifier-resourceType) for available types. In order to use this attribute, `all_supported` must be set to false.
     * 
     */
    private @Nullable List<String> resourceTypes;

    private RecorderRecordingGroup() {}
    /**
     * @return Specifies whether AWS Config records configuration changes for every supported type of regional resource (which includes any new type that will become supported in the future). Conflicts with `resource_types`. Defaults to `true`.
     * 
     */
    public Optional<Boolean> allSupported() {
        return Optional.ofNullable(this.allSupported);
    }
    /**
     * @return An object that specifies how AWS Config excludes resource types from being recorded by the configuration recorder.To use this option, you must set the useOnly field of RecordingStrategy to `EXCLUSION_BY_RESOURCE_TYPES` Requires `all_supported = false`. Conflicts with `resource_types`.
     * 
     */
    public List<RecorderRecordingGroupExclusionByResourceType> exclusionByResourceTypes() {
        return this.exclusionByResourceTypes == null ? List.of() : this.exclusionByResourceTypes;
    }
    /**
     * @return Specifies whether AWS Config includes all supported types of _global resources_ with the resources that it records. Requires `all_supported = true`. Conflicts with `resource_types`.
     * 
     */
    public Optional<Boolean> includeGlobalResourceTypes() {
        return Optional.ofNullable(this.includeGlobalResourceTypes);
    }
    /**
     * @return Recording Strategy - see below..
     * 
     */
    public List<RecorderRecordingGroupRecordingStrategy> recordingStrategies() {
        return this.recordingStrategies == null ? List.of() : this.recordingStrategies;
    }
    /**
     * @return A list that specifies the types of AWS resources for which AWS Config records configuration changes (for example, `AWS::EC2::Instance` or `AWS::CloudTrail::Trail`). See [relevant part of AWS Docs](http://docs.aws.amazon.com/config/latest/APIReference/API_ResourceIdentifier.html#config-Type-ResourceIdentifier-resourceType) for available types. In order to use this attribute, `all_supported` must be set to false.
     * 
     */
    public List<String> resourceTypes() {
        return this.resourceTypes == null ? List.of() : this.resourceTypes;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RecorderRecordingGroup defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean allSupported;
        private @Nullable List<RecorderRecordingGroupExclusionByResourceType> exclusionByResourceTypes;
        private @Nullable Boolean includeGlobalResourceTypes;
        private @Nullable List<RecorderRecordingGroupRecordingStrategy> recordingStrategies;
        private @Nullable List<String> resourceTypes;
        public Builder() {}
        public Builder(RecorderRecordingGroup defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.allSupported = defaults.allSupported;
    	      this.exclusionByResourceTypes = defaults.exclusionByResourceTypes;
    	      this.includeGlobalResourceTypes = defaults.includeGlobalResourceTypes;
    	      this.recordingStrategies = defaults.recordingStrategies;
    	      this.resourceTypes = defaults.resourceTypes;
        }

        @CustomType.Setter
        public Builder allSupported(@Nullable Boolean allSupported) {
            this.allSupported = allSupported;
            return this;
        }
        @CustomType.Setter
        public Builder exclusionByResourceTypes(@Nullable List<RecorderRecordingGroupExclusionByResourceType> exclusionByResourceTypes) {
            this.exclusionByResourceTypes = exclusionByResourceTypes;
            return this;
        }
        public Builder exclusionByResourceTypes(RecorderRecordingGroupExclusionByResourceType... exclusionByResourceTypes) {
            return exclusionByResourceTypes(List.of(exclusionByResourceTypes));
        }
        @CustomType.Setter
        public Builder includeGlobalResourceTypes(@Nullable Boolean includeGlobalResourceTypes) {
            this.includeGlobalResourceTypes = includeGlobalResourceTypes;
            return this;
        }
        @CustomType.Setter
        public Builder recordingStrategies(@Nullable List<RecorderRecordingGroupRecordingStrategy> recordingStrategies) {
            this.recordingStrategies = recordingStrategies;
            return this;
        }
        public Builder recordingStrategies(RecorderRecordingGroupRecordingStrategy... recordingStrategies) {
            return recordingStrategies(List.of(recordingStrategies));
        }
        @CustomType.Setter
        public Builder resourceTypes(@Nullable List<String> resourceTypes) {
            this.resourceTypes = resourceTypes;
            return this;
        }
        public Builder resourceTypes(String... resourceTypes) {
            return resourceTypes(List.of(resourceTypes));
        }
        public RecorderRecordingGroup build() {
            final var o = new RecorderRecordingGroup();
            o.allSupported = allSupported;
            o.exclusionByResourceTypes = exclusionByResourceTypes;
            o.includeGlobalResourceTypes = includeGlobalResourceTypes;
            o.recordingStrategies = recordingStrategies;
            o.resourceTypes = resourceTypes;
            return o;
        }
    }
}
