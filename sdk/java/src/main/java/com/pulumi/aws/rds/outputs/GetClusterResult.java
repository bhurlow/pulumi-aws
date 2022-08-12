// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.rds.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetClusterResult {
    private final String arn;
    private final List<String> availabilityZones;
    private final Integer backtrackWindow;
    private final Integer backupRetentionPeriod;
    private final String clusterIdentifier;
    private final List<String> clusterMembers;
    private final String clusterResourceId;
    private final String databaseName;
    private final String dbClusterParameterGroupName;
    private final String dbSubnetGroupName;
    private final List<String> enabledCloudwatchLogsExports;
    private final String endpoint;
    private final String engine;
    private final String engineVersion;
    private final String finalSnapshotIdentifier;
    private final String hostedZoneId;
    private final Boolean iamDatabaseAuthenticationEnabled;
    private final List<String> iamRoles;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private final String id;
    private final String kmsKeyId;
    private final String masterUsername;
    private final Integer port;
    private final String preferredBackupWindow;
    private final String preferredMaintenanceWindow;
    private final String readerEndpoint;
    private final String replicationSourceIdentifier;
    private final Boolean storageEncrypted;
    private final Map<String,String> tags;
    private final List<String> vpcSecurityGroupIds;

    @CustomType.Constructor
    private GetClusterResult(
        @CustomType.Parameter("arn") String arn,
        @CustomType.Parameter("availabilityZones") List<String> availabilityZones,
        @CustomType.Parameter("backtrackWindow") Integer backtrackWindow,
        @CustomType.Parameter("backupRetentionPeriod") Integer backupRetentionPeriod,
        @CustomType.Parameter("clusterIdentifier") String clusterIdentifier,
        @CustomType.Parameter("clusterMembers") List<String> clusterMembers,
        @CustomType.Parameter("clusterResourceId") String clusterResourceId,
        @CustomType.Parameter("databaseName") String databaseName,
        @CustomType.Parameter("dbClusterParameterGroupName") String dbClusterParameterGroupName,
        @CustomType.Parameter("dbSubnetGroupName") String dbSubnetGroupName,
        @CustomType.Parameter("enabledCloudwatchLogsExports") List<String> enabledCloudwatchLogsExports,
        @CustomType.Parameter("endpoint") String endpoint,
        @CustomType.Parameter("engine") String engine,
        @CustomType.Parameter("engineVersion") String engineVersion,
        @CustomType.Parameter("finalSnapshotIdentifier") String finalSnapshotIdentifier,
        @CustomType.Parameter("hostedZoneId") String hostedZoneId,
        @CustomType.Parameter("iamDatabaseAuthenticationEnabled") Boolean iamDatabaseAuthenticationEnabled,
        @CustomType.Parameter("iamRoles") List<String> iamRoles,
        @CustomType.Parameter("id") String id,
        @CustomType.Parameter("kmsKeyId") String kmsKeyId,
        @CustomType.Parameter("masterUsername") String masterUsername,
        @CustomType.Parameter("port") Integer port,
        @CustomType.Parameter("preferredBackupWindow") String preferredBackupWindow,
        @CustomType.Parameter("preferredMaintenanceWindow") String preferredMaintenanceWindow,
        @CustomType.Parameter("readerEndpoint") String readerEndpoint,
        @CustomType.Parameter("replicationSourceIdentifier") String replicationSourceIdentifier,
        @CustomType.Parameter("storageEncrypted") Boolean storageEncrypted,
        @CustomType.Parameter("tags") Map<String,String> tags,
        @CustomType.Parameter("vpcSecurityGroupIds") List<String> vpcSecurityGroupIds) {
        this.arn = arn;
        this.availabilityZones = availabilityZones;
        this.backtrackWindow = backtrackWindow;
        this.backupRetentionPeriod = backupRetentionPeriod;
        this.clusterIdentifier = clusterIdentifier;
        this.clusterMembers = clusterMembers;
        this.clusterResourceId = clusterResourceId;
        this.databaseName = databaseName;
        this.dbClusterParameterGroupName = dbClusterParameterGroupName;
        this.dbSubnetGroupName = dbSubnetGroupName;
        this.enabledCloudwatchLogsExports = enabledCloudwatchLogsExports;
        this.endpoint = endpoint;
        this.engine = engine;
        this.engineVersion = engineVersion;
        this.finalSnapshotIdentifier = finalSnapshotIdentifier;
        this.hostedZoneId = hostedZoneId;
        this.iamDatabaseAuthenticationEnabled = iamDatabaseAuthenticationEnabled;
        this.iamRoles = iamRoles;
        this.id = id;
        this.kmsKeyId = kmsKeyId;
        this.masterUsername = masterUsername;
        this.port = port;
        this.preferredBackupWindow = preferredBackupWindow;
        this.preferredMaintenanceWindow = preferredMaintenanceWindow;
        this.readerEndpoint = readerEndpoint;
        this.replicationSourceIdentifier = replicationSourceIdentifier;
        this.storageEncrypted = storageEncrypted;
        this.tags = tags;
        this.vpcSecurityGroupIds = vpcSecurityGroupIds;
    }

    public String arn() {
        return this.arn;
    }
    public List<String> availabilityZones() {
        return this.availabilityZones;
    }
    public Integer backtrackWindow() {
        return this.backtrackWindow;
    }
    public Integer backupRetentionPeriod() {
        return this.backupRetentionPeriod;
    }
    public String clusterIdentifier() {
        return this.clusterIdentifier;
    }
    public List<String> clusterMembers() {
        return this.clusterMembers;
    }
    public String clusterResourceId() {
        return this.clusterResourceId;
    }
    public String databaseName() {
        return this.databaseName;
    }
    public String dbClusterParameterGroupName() {
        return this.dbClusterParameterGroupName;
    }
    public String dbSubnetGroupName() {
        return this.dbSubnetGroupName;
    }
    public List<String> enabledCloudwatchLogsExports() {
        return this.enabledCloudwatchLogsExports;
    }
    public String endpoint() {
        return this.endpoint;
    }
    public String engine() {
        return this.engine;
    }
    public String engineVersion() {
        return this.engineVersion;
    }
    public String finalSnapshotIdentifier() {
        return this.finalSnapshotIdentifier;
    }
    public String hostedZoneId() {
        return this.hostedZoneId;
    }
    public Boolean iamDatabaseAuthenticationEnabled() {
        return this.iamDatabaseAuthenticationEnabled;
    }
    public List<String> iamRoles() {
        return this.iamRoles;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String kmsKeyId() {
        return this.kmsKeyId;
    }
    public String masterUsername() {
        return this.masterUsername;
    }
    public Integer port() {
        return this.port;
    }
    public String preferredBackupWindow() {
        return this.preferredBackupWindow;
    }
    public String preferredMaintenanceWindow() {
        return this.preferredMaintenanceWindow;
    }
    public String readerEndpoint() {
        return this.readerEndpoint;
    }
    public String replicationSourceIdentifier() {
        return this.replicationSourceIdentifier;
    }
    public Boolean storageEncrypted() {
        return this.storageEncrypted;
    }
    public Map<String,String> tags() {
        return this.tags;
    }
    public List<String> vpcSecurityGroupIds() {
        return this.vpcSecurityGroupIds;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetClusterResult defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private String arn;
        private List<String> availabilityZones;
        private Integer backtrackWindow;
        private Integer backupRetentionPeriod;
        private String clusterIdentifier;
        private List<String> clusterMembers;
        private String clusterResourceId;
        private String databaseName;
        private String dbClusterParameterGroupName;
        private String dbSubnetGroupName;
        private List<String> enabledCloudwatchLogsExports;
        private String endpoint;
        private String engine;
        private String engineVersion;
        private String finalSnapshotIdentifier;
        private String hostedZoneId;
        private Boolean iamDatabaseAuthenticationEnabled;
        private List<String> iamRoles;
        private String id;
        private String kmsKeyId;
        private String masterUsername;
        private Integer port;
        private String preferredBackupWindow;
        private String preferredMaintenanceWindow;
        private String readerEndpoint;
        private String replicationSourceIdentifier;
        private Boolean storageEncrypted;
        private Map<String,String> tags;
        private List<String> vpcSecurityGroupIds;

        public Builder() {
    	      // Empty
        }

        public Builder(GetClusterResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.arn = defaults.arn;
    	      this.availabilityZones = defaults.availabilityZones;
    	      this.backtrackWindow = defaults.backtrackWindow;
    	      this.backupRetentionPeriod = defaults.backupRetentionPeriod;
    	      this.clusterIdentifier = defaults.clusterIdentifier;
    	      this.clusterMembers = defaults.clusterMembers;
    	      this.clusterResourceId = defaults.clusterResourceId;
    	      this.databaseName = defaults.databaseName;
    	      this.dbClusterParameterGroupName = defaults.dbClusterParameterGroupName;
    	      this.dbSubnetGroupName = defaults.dbSubnetGroupName;
    	      this.enabledCloudwatchLogsExports = defaults.enabledCloudwatchLogsExports;
    	      this.endpoint = defaults.endpoint;
    	      this.engine = defaults.engine;
    	      this.engineVersion = defaults.engineVersion;
    	      this.finalSnapshotIdentifier = defaults.finalSnapshotIdentifier;
    	      this.hostedZoneId = defaults.hostedZoneId;
    	      this.iamDatabaseAuthenticationEnabled = defaults.iamDatabaseAuthenticationEnabled;
    	      this.iamRoles = defaults.iamRoles;
    	      this.id = defaults.id;
    	      this.kmsKeyId = defaults.kmsKeyId;
    	      this.masterUsername = defaults.masterUsername;
    	      this.port = defaults.port;
    	      this.preferredBackupWindow = defaults.preferredBackupWindow;
    	      this.preferredMaintenanceWindow = defaults.preferredMaintenanceWindow;
    	      this.readerEndpoint = defaults.readerEndpoint;
    	      this.replicationSourceIdentifier = defaults.replicationSourceIdentifier;
    	      this.storageEncrypted = defaults.storageEncrypted;
    	      this.tags = defaults.tags;
    	      this.vpcSecurityGroupIds = defaults.vpcSecurityGroupIds;
        }

        public Builder arn(String arn) {
            this.arn = Objects.requireNonNull(arn);
            return this;
        }
        public Builder availabilityZones(List<String> availabilityZones) {
            this.availabilityZones = Objects.requireNonNull(availabilityZones);
            return this;
        }
        public Builder availabilityZones(String... availabilityZones) {
            return availabilityZones(List.of(availabilityZones));
        }
        public Builder backtrackWindow(Integer backtrackWindow) {
            this.backtrackWindow = Objects.requireNonNull(backtrackWindow);
            return this;
        }
        public Builder backupRetentionPeriod(Integer backupRetentionPeriod) {
            this.backupRetentionPeriod = Objects.requireNonNull(backupRetentionPeriod);
            return this;
        }
        public Builder clusterIdentifier(String clusterIdentifier) {
            this.clusterIdentifier = Objects.requireNonNull(clusterIdentifier);
            return this;
        }
        public Builder clusterMembers(List<String> clusterMembers) {
            this.clusterMembers = Objects.requireNonNull(clusterMembers);
            return this;
        }
        public Builder clusterMembers(String... clusterMembers) {
            return clusterMembers(List.of(clusterMembers));
        }
        public Builder clusterResourceId(String clusterResourceId) {
            this.clusterResourceId = Objects.requireNonNull(clusterResourceId);
            return this;
        }
        public Builder databaseName(String databaseName) {
            this.databaseName = Objects.requireNonNull(databaseName);
            return this;
        }
        public Builder dbClusterParameterGroupName(String dbClusterParameterGroupName) {
            this.dbClusterParameterGroupName = Objects.requireNonNull(dbClusterParameterGroupName);
            return this;
        }
        public Builder dbSubnetGroupName(String dbSubnetGroupName) {
            this.dbSubnetGroupName = Objects.requireNonNull(dbSubnetGroupName);
            return this;
        }
        public Builder enabledCloudwatchLogsExports(List<String> enabledCloudwatchLogsExports) {
            this.enabledCloudwatchLogsExports = Objects.requireNonNull(enabledCloudwatchLogsExports);
            return this;
        }
        public Builder enabledCloudwatchLogsExports(String... enabledCloudwatchLogsExports) {
            return enabledCloudwatchLogsExports(List.of(enabledCloudwatchLogsExports));
        }
        public Builder endpoint(String endpoint) {
            this.endpoint = Objects.requireNonNull(endpoint);
            return this;
        }
        public Builder engine(String engine) {
            this.engine = Objects.requireNonNull(engine);
            return this;
        }
        public Builder engineVersion(String engineVersion) {
            this.engineVersion = Objects.requireNonNull(engineVersion);
            return this;
        }
        public Builder finalSnapshotIdentifier(String finalSnapshotIdentifier) {
            this.finalSnapshotIdentifier = Objects.requireNonNull(finalSnapshotIdentifier);
            return this;
        }
        public Builder hostedZoneId(String hostedZoneId) {
            this.hostedZoneId = Objects.requireNonNull(hostedZoneId);
            return this;
        }
        public Builder iamDatabaseAuthenticationEnabled(Boolean iamDatabaseAuthenticationEnabled) {
            this.iamDatabaseAuthenticationEnabled = Objects.requireNonNull(iamDatabaseAuthenticationEnabled);
            return this;
        }
        public Builder iamRoles(List<String> iamRoles) {
            this.iamRoles = Objects.requireNonNull(iamRoles);
            return this;
        }
        public Builder iamRoles(String... iamRoles) {
            return iamRoles(List.of(iamRoles));
        }
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        public Builder kmsKeyId(String kmsKeyId) {
            this.kmsKeyId = Objects.requireNonNull(kmsKeyId);
            return this;
        }
        public Builder masterUsername(String masterUsername) {
            this.masterUsername = Objects.requireNonNull(masterUsername);
            return this;
        }
        public Builder port(Integer port) {
            this.port = Objects.requireNonNull(port);
            return this;
        }
        public Builder preferredBackupWindow(String preferredBackupWindow) {
            this.preferredBackupWindow = Objects.requireNonNull(preferredBackupWindow);
            return this;
        }
        public Builder preferredMaintenanceWindow(String preferredMaintenanceWindow) {
            this.preferredMaintenanceWindow = Objects.requireNonNull(preferredMaintenanceWindow);
            return this;
        }
        public Builder readerEndpoint(String readerEndpoint) {
            this.readerEndpoint = Objects.requireNonNull(readerEndpoint);
            return this;
        }
        public Builder replicationSourceIdentifier(String replicationSourceIdentifier) {
            this.replicationSourceIdentifier = Objects.requireNonNull(replicationSourceIdentifier);
            return this;
        }
        public Builder storageEncrypted(Boolean storageEncrypted) {
            this.storageEncrypted = Objects.requireNonNull(storageEncrypted);
            return this;
        }
        public Builder tags(Map<String,String> tags) {
            this.tags = Objects.requireNonNull(tags);
            return this;
        }
        public Builder vpcSecurityGroupIds(List<String> vpcSecurityGroupIds) {
            this.vpcSecurityGroupIds = Objects.requireNonNull(vpcSecurityGroupIds);
            return this;
        }
        public Builder vpcSecurityGroupIds(String... vpcSecurityGroupIds) {
            return vpcSecurityGroupIds(List.of(vpcSecurityGroupIds));
        }        public GetClusterResult build() {
            return new GetClusterResult(arn, availabilityZones, backtrackWindow, backupRetentionPeriod, clusterIdentifier, clusterMembers, clusterResourceId, databaseName, dbClusterParameterGroupName, dbSubnetGroupName, enabledCloudwatchLogsExports, endpoint, engine, engineVersion, finalSnapshotIdentifier, hostedZoneId, iamDatabaseAuthenticationEnabled, iamRoles, id, kmsKeyId, masterUsername, port, preferredBackupWindow, preferredMaintenanceWindow, readerEndpoint, replicationSourceIdentifier, storageEncrypted, tags, vpcSecurityGroupIds);
        }
    }
}