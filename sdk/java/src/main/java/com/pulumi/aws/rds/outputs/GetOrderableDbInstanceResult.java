// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.rds.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Double;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetOrderableDbInstanceResult {
    private String availabilityZoneGroup;
    /**
     * @return Availability zones where the instance is available.
     * 
     */
    private List<String> availabilityZones;
    private String engine;
    private @Nullable Boolean engineLatestVersion;
    private String engineVersion;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String instanceClass;
    private String licenseModel;
    /**
     * @return Maximum total provisioned IOPS for a DB instance.
     * 
     */
    private Integer maxIopsPerDbInstance;
    /**
     * @return Maximum provisioned IOPS per GiB for a DB instance.
     * 
     */
    private Double maxIopsPerGib;
    /**
     * @return Maximum storage size for a DB instance.
     * 
     */
    private Integer maxStorageSize;
    /**
     * @return Minimum total provisioned IOPS for a DB instance.
     * 
     */
    private Integer minIopsPerDbInstance;
    /**
     * @return Minimum provisioned IOPS per GiB for a DB instance.
     * 
     */
    private Double minIopsPerGib;
    /**
     * @return Minimum storage size for a DB instance.
     * 
     */
    private Integer minStorageSize;
    /**
     * @return Whether a DB instance is Multi-AZ capable.
     * 
     */
    private Boolean multiAzCapable;
    /**
     * @return Whether a DB instance supports RDS on Outposts.
     * 
     */
    private Boolean outpostCapable;
    private @Nullable List<String> preferredEngineVersions;
    private @Nullable List<String> preferredInstanceClasses;
    private Boolean readReplicaCapable;
    private String storageType;
    private List<String> supportedEngineModes;
    private List<String> supportedNetworkTypes;
    private Boolean supportsClusters;
    private Boolean supportsEnhancedMonitoring;
    private Boolean supportsGlobalDatabases;
    private Boolean supportsIamDatabaseAuthentication;
    private Boolean supportsIops;
    private Boolean supportsKerberosAuthentication;
    private Boolean supportsMultiAz;
    private Boolean supportsPerformanceInsights;
    private Boolean supportsStorageAutoscaling;
    private Boolean supportsStorageEncryption;
    private Boolean vpc;

    private GetOrderableDbInstanceResult() {}
    public String availabilityZoneGroup() {
        return this.availabilityZoneGroup;
    }
    /**
     * @return Availability zones where the instance is available.
     * 
     */
    public List<String> availabilityZones() {
        return this.availabilityZones;
    }
    public String engine() {
        return this.engine;
    }
    public Optional<Boolean> engineLatestVersion() {
        return Optional.ofNullable(this.engineLatestVersion);
    }
    public String engineVersion() {
        return this.engineVersion;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String instanceClass() {
        return this.instanceClass;
    }
    public String licenseModel() {
        return this.licenseModel;
    }
    /**
     * @return Maximum total provisioned IOPS for a DB instance.
     * 
     */
    public Integer maxIopsPerDbInstance() {
        return this.maxIopsPerDbInstance;
    }
    /**
     * @return Maximum provisioned IOPS per GiB for a DB instance.
     * 
     */
    public Double maxIopsPerGib() {
        return this.maxIopsPerGib;
    }
    /**
     * @return Maximum storage size for a DB instance.
     * 
     */
    public Integer maxStorageSize() {
        return this.maxStorageSize;
    }
    /**
     * @return Minimum total provisioned IOPS for a DB instance.
     * 
     */
    public Integer minIopsPerDbInstance() {
        return this.minIopsPerDbInstance;
    }
    /**
     * @return Minimum provisioned IOPS per GiB for a DB instance.
     * 
     */
    public Double minIopsPerGib() {
        return this.minIopsPerGib;
    }
    /**
     * @return Minimum storage size for a DB instance.
     * 
     */
    public Integer minStorageSize() {
        return this.minStorageSize;
    }
    /**
     * @return Whether a DB instance is Multi-AZ capable.
     * 
     */
    public Boolean multiAzCapable() {
        return this.multiAzCapable;
    }
    /**
     * @return Whether a DB instance supports RDS on Outposts.
     * 
     */
    public Boolean outpostCapable() {
        return this.outpostCapable;
    }
    public List<String> preferredEngineVersions() {
        return this.preferredEngineVersions == null ? List.of() : this.preferredEngineVersions;
    }
    public List<String> preferredInstanceClasses() {
        return this.preferredInstanceClasses == null ? List.of() : this.preferredInstanceClasses;
    }
    public Boolean readReplicaCapable() {
        return this.readReplicaCapable;
    }
    public String storageType() {
        return this.storageType;
    }
    public List<String> supportedEngineModes() {
        return this.supportedEngineModes;
    }
    public List<String> supportedNetworkTypes() {
        return this.supportedNetworkTypes;
    }
    public Boolean supportsClusters() {
        return this.supportsClusters;
    }
    public Boolean supportsEnhancedMonitoring() {
        return this.supportsEnhancedMonitoring;
    }
    public Boolean supportsGlobalDatabases() {
        return this.supportsGlobalDatabases;
    }
    public Boolean supportsIamDatabaseAuthentication() {
        return this.supportsIamDatabaseAuthentication;
    }
    public Boolean supportsIops() {
        return this.supportsIops;
    }
    public Boolean supportsKerberosAuthentication() {
        return this.supportsKerberosAuthentication;
    }
    public Boolean supportsMultiAz() {
        return this.supportsMultiAz;
    }
    public Boolean supportsPerformanceInsights() {
        return this.supportsPerformanceInsights;
    }
    public Boolean supportsStorageAutoscaling() {
        return this.supportsStorageAutoscaling;
    }
    public Boolean supportsStorageEncryption() {
        return this.supportsStorageEncryption;
    }
    public Boolean vpc() {
        return this.vpc;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetOrderableDbInstanceResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String availabilityZoneGroup;
        private List<String> availabilityZones;
        private String engine;
        private @Nullable Boolean engineLatestVersion;
        private String engineVersion;
        private String id;
        private String instanceClass;
        private String licenseModel;
        private Integer maxIopsPerDbInstance;
        private Double maxIopsPerGib;
        private Integer maxStorageSize;
        private Integer minIopsPerDbInstance;
        private Double minIopsPerGib;
        private Integer minStorageSize;
        private Boolean multiAzCapable;
        private Boolean outpostCapable;
        private @Nullable List<String> preferredEngineVersions;
        private @Nullable List<String> preferredInstanceClasses;
        private Boolean readReplicaCapable;
        private String storageType;
        private List<String> supportedEngineModes;
        private List<String> supportedNetworkTypes;
        private Boolean supportsClusters;
        private Boolean supportsEnhancedMonitoring;
        private Boolean supportsGlobalDatabases;
        private Boolean supportsIamDatabaseAuthentication;
        private Boolean supportsIops;
        private Boolean supportsKerberosAuthentication;
        private Boolean supportsMultiAz;
        private Boolean supportsPerformanceInsights;
        private Boolean supportsStorageAutoscaling;
        private Boolean supportsStorageEncryption;
        private Boolean vpc;
        public Builder() {}
        public Builder(GetOrderableDbInstanceResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.availabilityZoneGroup = defaults.availabilityZoneGroup;
    	      this.availabilityZones = defaults.availabilityZones;
    	      this.engine = defaults.engine;
    	      this.engineLatestVersion = defaults.engineLatestVersion;
    	      this.engineVersion = defaults.engineVersion;
    	      this.id = defaults.id;
    	      this.instanceClass = defaults.instanceClass;
    	      this.licenseModel = defaults.licenseModel;
    	      this.maxIopsPerDbInstance = defaults.maxIopsPerDbInstance;
    	      this.maxIopsPerGib = defaults.maxIopsPerGib;
    	      this.maxStorageSize = defaults.maxStorageSize;
    	      this.minIopsPerDbInstance = defaults.minIopsPerDbInstance;
    	      this.minIopsPerGib = defaults.minIopsPerGib;
    	      this.minStorageSize = defaults.minStorageSize;
    	      this.multiAzCapable = defaults.multiAzCapable;
    	      this.outpostCapable = defaults.outpostCapable;
    	      this.preferredEngineVersions = defaults.preferredEngineVersions;
    	      this.preferredInstanceClasses = defaults.preferredInstanceClasses;
    	      this.readReplicaCapable = defaults.readReplicaCapable;
    	      this.storageType = defaults.storageType;
    	      this.supportedEngineModes = defaults.supportedEngineModes;
    	      this.supportedNetworkTypes = defaults.supportedNetworkTypes;
    	      this.supportsClusters = defaults.supportsClusters;
    	      this.supportsEnhancedMonitoring = defaults.supportsEnhancedMonitoring;
    	      this.supportsGlobalDatabases = defaults.supportsGlobalDatabases;
    	      this.supportsIamDatabaseAuthentication = defaults.supportsIamDatabaseAuthentication;
    	      this.supportsIops = defaults.supportsIops;
    	      this.supportsKerberosAuthentication = defaults.supportsKerberosAuthentication;
    	      this.supportsMultiAz = defaults.supportsMultiAz;
    	      this.supportsPerformanceInsights = defaults.supportsPerformanceInsights;
    	      this.supportsStorageAutoscaling = defaults.supportsStorageAutoscaling;
    	      this.supportsStorageEncryption = defaults.supportsStorageEncryption;
    	      this.vpc = defaults.vpc;
        }

        @CustomType.Setter
        public Builder availabilityZoneGroup(String availabilityZoneGroup) {
            if (availabilityZoneGroup == null) {
              throw new MissingRequiredPropertyException("GetOrderableDbInstanceResult", "availabilityZoneGroup");
            }
            this.availabilityZoneGroup = availabilityZoneGroup;
            return this;
        }
        @CustomType.Setter
        public Builder availabilityZones(List<String> availabilityZones) {
            if (availabilityZones == null) {
              throw new MissingRequiredPropertyException("GetOrderableDbInstanceResult", "availabilityZones");
            }
            this.availabilityZones = availabilityZones;
            return this;
        }
        public Builder availabilityZones(String... availabilityZones) {
            return availabilityZones(List.of(availabilityZones));
        }
        @CustomType.Setter
        public Builder engine(String engine) {
            if (engine == null) {
              throw new MissingRequiredPropertyException("GetOrderableDbInstanceResult", "engine");
            }
            this.engine = engine;
            return this;
        }
        @CustomType.Setter
        public Builder engineLatestVersion(@Nullable Boolean engineLatestVersion) {

            this.engineLatestVersion = engineLatestVersion;
            return this;
        }
        @CustomType.Setter
        public Builder engineVersion(String engineVersion) {
            if (engineVersion == null) {
              throw new MissingRequiredPropertyException("GetOrderableDbInstanceResult", "engineVersion");
            }
            this.engineVersion = engineVersion;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetOrderableDbInstanceResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder instanceClass(String instanceClass) {
            if (instanceClass == null) {
              throw new MissingRequiredPropertyException("GetOrderableDbInstanceResult", "instanceClass");
            }
            this.instanceClass = instanceClass;
            return this;
        }
        @CustomType.Setter
        public Builder licenseModel(String licenseModel) {
            if (licenseModel == null) {
              throw new MissingRequiredPropertyException("GetOrderableDbInstanceResult", "licenseModel");
            }
            this.licenseModel = licenseModel;
            return this;
        }
        @CustomType.Setter
        public Builder maxIopsPerDbInstance(Integer maxIopsPerDbInstance) {
            if (maxIopsPerDbInstance == null) {
              throw new MissingRequiredPropertyException("GetOrderableDbInstanceResult", "maxIopsPerDbInstance");
            }
            this.maxIopsPerDbInstance = maxIopsPerDbInstance;
            return this;
        }
        @CustomType.Setter
        public Builder maxIopsPerGib(Double maxIopsPerGib) {
            if (maxIopsPerGib == null) {
              throw new MissingRequiredPropertyException("GetOrderableDbInstanceResult", "maxIopsPerGib");
            }
            this.maxIopsPerGib = maxIopsPerGib;
            return this;
        }
        @CustomType.Setter
        public Builder maxStorageSize(Integer maxStorageSize) {
            if (maxStorageSize == null) {
              throw new MissingRequiredPropertyException("GetOrderableDbInstanceResult", "maxStorageSize");
            }
            this.maxStorageSize = maxStorageSize;
            return this;
        }
        @CustomType.Setter
        public Builder minIopsPerDbInstance(Integer minIopsPerDbInstance) {
            if (minIopsPerDbInstance == null) {
              throw new MissingRequiredPropertyException("GetOrderableDbInstanceResult", "minIopsPerDbInstance");
            }
            this.minIopsPerDbInstance = minIopsPerDbInstance;
            return this;
        }
        @CustomType.Setter
        public Builder minIopsPerGib(Double minIopsPerGib) {
            if (minIopsPerGib == null) {
              throw new MissingRequiredPropertyException("GetOrderableDbInstanceResult", "minIopsPerGib");
            }
            this.minIopsPerGib = minIopsPerGib;
            return this;
        }
        @CustomType.Setter
        public Builder minStorageSize(Integer minStorageSize) {
            if (minStorageSize == null) {
              throw new MissingRequiredPropertyException("GetOrderableDbInstanceResult", "minStorageSize");
            }
            this.minStorageSize = minStorageSize;
            return this;
        }
        @CustomType.Setter
        public Builder multiAzCapable(Boolean multiAzCapable) {
            if (multiAzCapable == null) {
              throw new MissingRequiredPropertyException("GetOrderableDbInstanceResult", "multiAzCapable");
            }
            this.multiAzCapable = multiAzCapable;
            return this;
        }
        @CustomType.Setter
        public Builder outpostCapable(Boolean outpostCapable) {
            if (outpostCapable == null) {
              throw new MissingRequiredPropertyException("GetOrderableDbInstanceResult", "outpostCapable");
            }
            this.outpostCapable = outpostCapable;
            return this;
        }
        @CustomType.Setter
        public Builder preferredEngineVersions(@Nullable List<String> preferredEngineVersions) {

            this.preferredEngineVersions = preferredEngineVersions;
            return this;
        }
        public Builder preferredEngineVersions(String... preferredEngineVersions) {
            return preferredEngineVersions(List.of(preferredEngineVersions));
        }
        @CustomType.Setter
        public Builder preferredInstanceClasses(@Nullable List<String> preferredInstanceClasses) {

            this.preferredInstanceClasses = preferredInstanceClasses;
            return this;
        }
        public Builder preferredInstanceClasses(String... preferredInstanceClasses) {
            return preferredInstanceClasses(List.of(preferredInstanceClasses));
        }
        @CustomType.Setter
        public Builder readReplicaCapable(Boolean readReplicaCapable) {
            if (readReplicaCapable == null) {
              throw new MissingRequiredPropertyException("GetOrderableDbInstanceResult", "readReplicaCapable");
            }
            this.readReplicaCapable = readReplicaCapable;
            return this;
        }
        @CustomType.Setter
        public Builder storageType(String storageType) {
            if (storageType == null) {
              throw new MissingRequiredPropertyException("GetOrderableDbInstanceResult", "storageType");
            }
            this.storageType = storageType;
            return this;
        }
        @CustomType.Setter
        public Builder supportedEngineModes(List<String> supportedEngineModes) {
            if (supportedEngineModes == null) {
              throw new MissingRequiredPropertyException("GetOrderableDbInstanceResult", "supportedEngineModes");
            }
            this.supportedEngineModes = supportedEngineModes;
            return this;
        }
        public Builder supportedEngineModes(String... supportedEngineModes) {
            return supportedEngineModes(List.of(supportedEngineModes));
        }
        @CustomType.Setter
        public Builder supportedNetworkTypes(List<String> supportedNetworkTypes) {
            if (supportedNetworkTypes == null) {
              throw new MissingRequiredPropertyException("GetOrderableDbInstanceResult", "supportedNetworkTypes");
            }
            this.supportedNetworkTypes = supportedNetworkTypes;
            return this;
        }
        public Builder supportedNetworkTypes(String... supportedNetworkTypes) {
            return supportedNetworkTypes(List.of(supportedNetworkTypes));
        }
        @CustomType.Setter
        public Builder supportsClusters(Boolean supportsClusters) {
            if (supportsClusters == null) {
              throw new MissingRequiredPropertyException("GetOrderableDbInstanceResult", "supportsClusters");
            }
            this.supportsClusters = supportsClusters;
            return this;
        }
        @CustomType.Setter
        public Builder supportsEnhancedMonitoring(Boolean supportsEnhancedMonitoring) {
            if (supportsEnhancedMonitoring == null) {
              throw new MissingRequiredPropertyException("GetOrderableDbInstanceResult", "supportsEnhancedMonitoring");
            }
            this.supportsEnhancedMonitoring = supportsEnhancedMonitoring;
            return this;
        }
        @CustomType.Setter
        public Builder supportsGlobalDatabases(Boolean supportsGlobalDatabases) {
            if (supportsGlobalDatabases == null) {
              throw new MissingRequiredPropertyException("GetOrderableDbInstanceResult", "supportsGlobalDatabases");
            }
            this.supportsGlobalDatabases = supportsGlobalDatabases;
            return this;
        }
        @CustomType.Setter
        public Builder supportsIamDatabaseAuthentication(Boolean supportsIamDatabaseAuthentication) {
            if (supportsIamDatabaseAuthentication == null) {
              throw new MissingRequiredPropertyException("GetOrderableDbInstanceResult", "supportsIamDatabaseAuthentication");
            }
            this.supportsIamDatabaseAuthentication = supportsIamDatabaseAuthentication;
            return this;
        }
        @CustomType.Setter
        public Builder supportsIops(Boolean supportsIops) {
            if (supportsIops == null) {
              throw new MissingRequiredPropertyException("GetOrderableDbInstanceResult", "supportsIops");
            }
            this.supportsIops = supportsIops;
            return this;
        }
        @CustomType.Setter
        public Builder supportsKerberosAuthentication(Boolean supportsKerberosAuthentication) {
            if (supportsKerberosAuthentication == null) {
              throw new MissingRequiredPropertyException("GetOrderableDbInstanceResult", "supportsKerberosAuthentication");
            }
            this.supportsKerberosAuthentication = supportsKerberosAuthentication;
            return this;
        }
        @CustomType.Setter
        public Builder supportsMultiAz(Boolean supportsMultiAz) {
            if (supportsMultiAz == null) {
              throw new MissingRequiredPropertyException("GetOrderableDbInstanceResult", "supportsMultiAz");
            }
            this.supportsMultiAz = supportsMultiAz;
            return this;
        }
        @CustomType.Setter
        public Builder supportsPerformanceInsights(Boolean supportsPerformanceInsights) {
            if (supportsPerformanceInsights == null) {
              throw new MissingRequiredPropertyException("GetOrderableDbInstanceResult", "supportsPerformanceInsights");
            }
            this.supportsPerformanceInsights = supportsPerformanceInsights;
            return this;
        }
        @CustomType.Setter
        public Builder supportsStorageAutoscaling(Boolean supportsStorageAutoscaling) {
            if (supportsStorageAutoscaling == null) {
              throw new MissingRequiredPropertyException("GetOrderableDbInstanceResult", "supportsStorageAutoscaling");
            }
            this.supportsStorageAutoscaling = supportsStorageAutoscaling;
            return this;
        }
        @CustomType.Setter
        public Builder supportsStorageEncryption(Boolean supportsStorageEncryption) {
            if (supportsStorageEncryption == null) {
              throw new MissingRequiredPropertyException("GetOrderableDbInstanceResult", "supportsStorageEncryption");
            }
            this.supportsStorageEncryption = supportsStorageEncryption;
            return this;
        }
        @CustomType.Setter
        public Builder vpc(Boolean vpc) {
            if (vpc == null) {
              throw new MissingRequiredPropertyException("GetOrderableDbInstanceResult", "vpc");
            }
            this.vpc = vpc;
            return this;
        }
        public GetOrderableDbInstanceResult build() {
            final var _resultValue = new GetOrderableDbInstanceResult();
            _resultValue.availabilityZoneGroup = availabilityZoneGroup;
            _resultValue.availabilityZones = availabilityZones;
            _resultValue.engine = engine;
            _resultValue.engineLatestVersion = engineLatestVersion;
            _resultValue.engineVersion = engineVersion;
            _resultValue.id = id;
            _resultValue.instanceClass = instanceClass;
            _resultValue.licenseModel = licenseModel;
            _resultValue.maxIopsPerDbInstance = maxIopsPerDbInstance;
            _resultValue.maxIopsPerGib = maxIopsPerGib;
            _resultValue.maxStorageSize = maxStorageSize;
            _resultValue.minIopsPerDbInstance = minIopsPerDbInstance;
            _resultValue.minIopsPerGib = minIopsPerGib;
            _resultValue.minStorageSize = minStorageSize;
            _resultValue.multiAzCapable = multiAzCapable;
            _resultValue.outpostCapable = outpostCapable;
            _resultValue.preferredEngineVersions = preferredEngineVersions;
            _resultValue.preferredInstanceClasses = preferredInstanceClasses;
            _resultValue.readReplicaCapable = readReplicaCapable;
            _resultValue.storageType = storageType;
            _resultValue.supportedEngineModes = supportedEngineModes;
            _resultValue.supportedNetworkTypes = supportedNetworkTypes;
            _resultValue.supportsClusters = supportsClusters;
            _resultValue.supportsEnhancedMonitoring = supportsEnhancedMonitoring;
            _resultValue.supportsGlobalDatabases = supportsGlobalDatabases;
            _resultValue.supportsIamDatabaseAuthentication = supportsIamDatabaseAuthentication;
            _resultValue.supportsIops = supportsIops;
            _resultValue.supportsKerberosAuthentication = supportsKerberosAuthentication;
            _resultValue.supportsMultiAz = supportsMultiAz;
            _resultValue.supportsPerformanceInsights = supportsPerformanceInsights;
            _resultValue.supportsStorageAutoscaling = supportsStorageAutoscaling;
            _resultValue.supportsStorageEncryption = supportsStorageEncryption;
            _resultValue.vpc = vpc;
            return _resultValue;
        }
    }
}
