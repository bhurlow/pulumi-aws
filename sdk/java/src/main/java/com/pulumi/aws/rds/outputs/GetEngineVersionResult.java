// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.rds.outputs;

import com.pulumi.aws.rds.outputs.GetEngineVersionFilter;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetEngineVersionResult {
    /**
     * @return The default character set for new instances of this engine version.
     * 
     */
    private String defaultCharacterSet;
    private @Nullable Boolean defaultOnly;
    private String engine;
    /**
     * @return Description of the database engine.
     * 
     */
    private String engineDescription;
    /**
     * @return Set of log types that the database engine has available for export to CloudWatch Logs.
     * 
     */
    private List<String> exportableLogTypes;
    private @Nullable List<GetEngineVersionFilter> filters;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable Boolean includeAll;
    private String parameterGroupFamily;
    private @Nullable List<String> preferredVersions;
    /**
     * @return Status of the DB engine version, either available or deprecated.
     * 
     */
    private String status;
    /**
     * @return Set of the character sets supported by this engine.
     * 
     */
    private List<String> supportedCharacterSets;
    /**
     * @return Set of features supported by the DB engine.
     * 
     */
    private List<String> supportedFeatureNames;
    /**
     * @return Set of the supported DB engine modes.
     * 
     */
    private List<String> supportedModes;
    /**
     * @return Set of the time zones supported by this engine.
     * 
     */
    private List<String> supportedTimezones;
    /**
     * @return Indicates whether you can use Aurora global databases with a specific DB engine version.
     * 
     */
    private Boolean supportsGlobalDatabases;
    /**
     * @return Indicates whether the engine version supports exporting the log types specified by `exportable_log_types` to CloudWatch Logs.
     * 
     */
    private Boolean supportsLogExportsToCloudwatch;
    /**
     * @return Indicates whether you can use Aurora parallel query with a specific DB engine version.
     * 
     */
    private Boolean supportsParallelQuery;
    /**
     * @return Indicates whether the database engine version supports read replicas.
     * 
     */
    private Boolean supportsReadReplica;
    /**
     * @return Set of engine versions that this database engine version can be upgraded to.
     * 
     */
    private List<String> validUpgradeTargets;
    private String version;
    /**
     * @return Description of the database engine version.
     * 
     */
    private String versionDescription;

    private GetEngineVersionResult() {}
    /**
     * @return The default character set for new instances of this engine version.
     * 
     */
    public String defaultCharacterSet() {
        return this.defaultCharacterSet;
    }
    public Optional<Boolean> defaultOnly() {
        return Optional.ofNullable(this.defaultOnly);
    }
    public String engine() {
        return this.engine;
    }
    /**
     * @return Description of the database engine.
     * 
     */
    public String engineDescription() {
        return this.engineDescription;
    }
    /**
     * @return Set of log types that the database engine has available for export to CloudWatch Logs.
     * 
     */
    public List<String> exportableLogTypes() {
        return this.exportableLogTypes;
    }
    public List<GetEngineVersionFilter> filters() {
        return this.filters == null ? List.of() : this.filters;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Optional<Boolean> includeAll() {
        return Optional.ofNullable(this.includeAll);
    }
    public String parameterGroupFamily() {
        return this.parameterGroupFamily;
    }
    public List<String> preferredVersions() {
        return this.preferredVersions == null ? List.of() : this.preferredVersions;
    }
    /**
     * @return Status of the DB engine version, either available or deprecated.
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return Set of the character sets supported by this engine.
     * 
     */
    public List<String> supportedCharacterSets() {
        return this.supportedCharacterSets;
    }
    /**
     * @return Set of features supported by the DB engine.
     * 
     */
    public List<String> supportedFeatureNames() {
        return this.supportedFeatureNames;
    }
    /**
     * @return Set of the supported DB engine modes.
     * 
     */
    public List<String> supportedModes() {
        return this.supportedModes;
    }
    /**
     * @return Set of the time zones supported by this engine.
     * 
     */
    public List<String> supportedTimezones() {
        return this.supportedTimezones;
    }
    /**
     * @return Indicates whether you can use Aurora global databases with a specific DB engine version.
     * 
     */
    public Boolean supportsGlobalDatabases() {
        return this.supportsGlobalDatabases;
    }
    /**
     * @return Indicates whether the engine version supports exporting the log types specified by `exportable_log_types` to CloudWatch Logs.
     * 
     */
    public Boolean supportsLogExportsToCloudwatch() {
        return this.supportsLogExportsToCloudwatch;
    }
    /**
     * @return Indicates whether you can use Aurora parallel query with a specific DB engine version.
     * 
     */
    public Boolean supportsParallelQuery() {
        return this.supportsParallelQuery;
    }
    /**
     * @return Indicates whether the database engine version supports read replicas.
     * 
     */
    public Boolean supportsReadReplica() {
        return this.supportsReadReplica;
    }
    /**
     * @return Set of engine versions that this database engine version can be upgraded to.
     * 
     */
    public List<String> validUpgradeTargets() {
        return this.validUpgradeTargets;
    }
    public String version() {
        return this.version;
    }
    /**
     * @return Description of the database engine version.
     * 
     */
    public String versionDescription() {
        return this.versionDescription;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetEngineVersionResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String defaultCharacterSet;
        private @Nullable Boolean defaultOnly;
        private String engine;
        private String engineDescription;
        private List<String> exportableLogTypes;
        private @Nullable List<GetEngineVersionFilter> filters;
        private String id;
        private @Nullable Boolean includeAll;
        private String parameterGroupFamily;
        private @Nullable List<String> preferredVersions;
        private String status;
        private List<String> supportedCharacterSets;
        private List<String> supportedFeatureNames;
        private List<String> supportedModes;
        private List<String> supportedTimezones;
        private Boolean supportsGlobalDatabases;
        private Boolean supportsLogExportsToCloudwatch;
        private Boolean supportsParallelQuery;
        private Boolean supportsReadReplica;
        private List<String> validUpgradeTargets;
        private String version;
        private String versionDescription;
        public Builder() {}
        public Builder(GetEngineVersionResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.defaultCharacterSet = defaults.defaultCharacterSet;
    	      this.defaultOnly = defaults.defaultOnly;
    	      this.engine = defaults.engine;
    	      this.engineDescription = defaults.engineDescription;
    	      this.exportableLogTypes = defaults.exportableLogTypes;
    	      this.filters = defaults.filters;
    	      this.id = defaults.id;
    	      this.includeAll = defaults.includeAll;
    	      this.parameterGroupFamily = defaults.parameterGroupFamily;
    	      this.preferredVersions = defaults.preferredVersions;
    	      this.status = defaults.status;
    	      this.supportedCharacterSets = defaults.supportedCharacterSets;
    	      this.supportedFeatureNames = defaults.supportedFeatureNames;
    	      this.supportedModes = defaults.supportedModes;
    	      this.supportedTimezones = defaults.supportedTimezones;
    	      this.supportsGlobalDatabases = defaults.supportsGlobalDatabases;
    	      this.supportsLogExportsToCloudwatch = defaults.supportsLogExportsToCloudwatch;
    	      this.supportsParallelQuery = defaults.supportsParallelQuery;
    	      this.supportsReadReplica = defaults.supportsReadReplica;
    	      this.validUpgradeTargets = defaults.validUpgradeTargets;
    	      this.version = defaults.version;
    	      this.versionDescription = defaults.versionDescription;
        }

        @CustomType.Setter
        public Builder defaultCharacterSet(String defaultCharacterSet) {
            this.defaultCharacterSet = Objects.requireNonNull(defaultCharacterSet);
            return this;
        }
        @CustomType.Setter
        public Builder defaultOnly(@Nullable Boolean defaultOnly) {
            this.defaultOnly = defaultOnly;
            return this;
        }
        @CustomType.Setter
        public Builder engine(String engine) {
            this.engine = Objects.requireNonNull(engine);
            return this;
        }
        @CustomType.Setter
        public Builder engineDescription(String engineDescription) {
            this.engineDescription = Objects.requireNonNull(engineDescription);
            return this;
        }
        @CustomType.Setter
        public Builder exportableLogTypes(List<String> exportableLogTypes) {
            this.exportableLogTypes = Objects.requireNonNull(exportableLogTypes);
            return this;
        }
        public Builder exportableLogTypes(String... exportableLogTypes) {
            return exportableLogTypes(List.of(exportableLogTypes));
        }
        @CustomType.Setter
        public Builder filters(@Nullable List<GetEngineVersionFilter> filters) {
            this.filters = filters;
            return this;
        }
        public Builder filters(GetEngineVersionFilter... filters) {
            return filters(List.of(filters));
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder includeAll(@Nullable Boolean includeAll) {
            this.includeAll = includeAll;
            return this;
        }
        @CustomType.Setter
        public Builder parameterGroupFamily(String parameterGroupFamily) {
            this.parameterGroupFamily = Objects.requireNonNull(parameterGroupFamily);
            return this;
        }
        @CustomType.Setter
        public Builder preferredVersions(@Nullable List<String> preferredVersions) {
            this.preferredVersions = preferredVersions;
            return this;
        }
        public Builder preferredVersions(String... preferredVersions) {
            return preferredVersions(List.of(preferredVersions));
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        @CustomType.Setter
        public Builder supportedCharacterSets(List<String> supportedCharacterSets) {
            this.supportedCharacterSets = Objects.requireNonNull(supportedCharacterSets);
            return this;
        }
        public Builder supportedCharacterSets(String... supportedCharacterSets) {
            return supportedCharacterSets(List.of(supportedCharacterSets));
        }
        @CustomType.Setter
        public Builder supportedFeatureNames(List<String> supportedFeatureNames) {
            this.supportedFeatureNames = Objects.requireNonNull(supportedFeatureNames);
            return this;
        }
        public Builder supportedFeatureNames(String... supportedFeatureNames) {
            return supportedFeatureNames(List.of(supportedFeatureNames));
        }
        @CustomType.Setter
        public Builder supportedModes(List<String> supportedModes) {
            this.supportedModes = Objects.requireNonNull(supportedModes);
            return this;
        }
        public Builder supportedModes(String... supportedModes) {
            return supportedModes(List.of(supportedModes));
        }
        @CustomType.Setter
        public Builder supportedTimezones(List<String> supportedTimezones) {
            this.supportedTimezones = Objects.requireNonNull(supportedTimezones);
            return this;
        }
        public Builder supportedTimezones(String... supportedTimezones) {
            return supportedTimezones(List.of(supportedTimezones));
        }
        @CustomType.Setter
        public Builder supportsGlobalDatabases(Boolean supportsGlobalDatabases) {
            this.supportsGlobalDatabases = Objects.requireNonNull(supportsGlobalDatabases);
            return this;
        }
        @CustomType.Setter
        public Builder supportsLogExportsToCloudwatch(Boolean supportsLogExportsToCloudwatch) {
            this.supportsLogExportsToCloudwatch = Objects.requireNonNull(supportsLogExportsToCloudwatch);
            return this;
        }
        @CustomType.Setter
        public Builder supportsParallelQuery(Boolean supportsParallelQuery) {
            this.supportsParallelQuery = Objects.requireNonNull(supportsParallelQuery);
            return this;
        }
        @CustomType.Setter
        public Builder supportsReadReplica(Boolean supportsReadReplica) {
            this.supportsReadReplica = Objects.requireNonNull(supportsReadReplica);
            return this;
        }
        @CustomType.Setter
        public Builder validUpgradeTargets(List<String> validUpgradeTargets) {
            this.validUpgradeTargets = Objects.requireNonNull(validUpgradeTargets);
            return this;
        }
        public Builder validUpgradeTargets(String... validUpgradeTargets) {
            return validUpgradeTargets(List.of(validUpgradeTargets));
        }
        @CustomType.Setter
        public Builder version(String version) {
            this.version = Objects.requireNonNull(version);
            return this;
        }
        @CustomType.Setter
        public Builder versionDescription(String versionDescription) {
            this.versionDescription = Objects.requireNonNull(versionDescription);
            return this;
        }
        public GetEngineVersionResult build() {
            final var o = new GetEngineVersionResult();
            o.defaultCharacterSet = defaultCharacterSet;
            o.defaultOnly = defaultOnly;
            o.engine = engine;
            o.engineDescription = engineDescription;
            o.exportableLogTypes = exportableLogTypes;
            o.filters = filters;
            o.id = id;
            o.includeAll = includeAll;
            o.parameterGroupFamily = parameterGroupFamily;
            o.preferredVersions = preferredVersions;
            o.status = status;
            o.supportedCharacterSets = supportedCharacterSets;
            o.supportedFeatureNames = supportedFeatureNames;
            o.supportedModes = supportedModes;
            o.supportedTimezones = supportedTimezones;
            o.supportsGlobalDatabases = supportsGlobalDatabases;
            o.supportsLogExportsToCloudwatch = supportsLogExportsToCloudwatch;
            o.supportsParallelQuery = supportsParallelQuery;
            o.supportsReadReplica = supportsReadReplica;
            o.validUpgradeTargets = validUpgradeTargets;
            o.version = version;
            o.versionDescription = versionDescription;
            return o;
        }
    }
}
