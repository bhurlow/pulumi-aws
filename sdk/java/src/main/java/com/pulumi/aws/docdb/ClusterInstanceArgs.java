// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.docdb;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ClusterInstanceArgs extends com.pulumi.resources.ResourceArgs {

    public static final ClusterInstanceArgs Empty = new ClusterInstanceArgs();

    /**
     * Specifies whether any database modifications
     * are applied immediately, or during the next maintenance window. Default is`false`.
     * 
     */
    @Import(name="applyImmediately")
    private @Nullable Output<Boolean> applyImmediately;

    /**
     * @return Specifies whether any database modifications
     * are applied immediately, or during the next maintenance window. Default is`false`.
     * 
     */
    public Optional<Output<Boolean>> applyImmediately() {
        return Optional.ofNullable(this.applyImmediately);
    }

    /**
     * This parameter does not apply to Amazon DocumentDB. Amazon DocumentDB does not perform minor version upgrades regardless of the value set (see [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBInstance.html)). Default `true`.
     * 
     */
    @Import(name="autoMinorVersionUpgrade")
    private @Nullable Output<Boolean> autoMinorVersionUpgrade;

    /**
     * @return This parameter does not apply to Amazon DocumentDB. Amazon DocumentDB does not perform minor version upgrades regardless of the value set (see [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBInstance.html)). Default `true`.
     * 
     */
    public Optional<Output<Boolean>> autoMinorVersionUpgrade() {
        return Optional.ofNullable(this.autoMinorVersionUpgrade);
    }

    /**
     * The EC2 Availability Zone that the DB instance is created in. See [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_CreateDBInstance.html) about the details.
     * 
     */
    @Import(name="availabilityZone")
    private @Nullable Output<String> availabilityZone;

    /**
     * @return The EC2 Availability Zone that the DB instance is created in. See [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_CreateDBInstance.html) about the details.
     * 
     */
    public Optional<Output<String>> availabilityZone() {
        return Optional.ofNullable(this.availabilityZone);
    }

    /**
     * (Optional) The identifier of the CA certificate for the DB instance.
     * 
     */
    @Import(name="caCertIdentifier")
    private @Nullable Output<String> caCertIdentifier;

    /**
     * @return (Optional) The identifier of the CA certificate for the DB instance.
     * 
     */
    public Optional<Output<String>> caCertIdentifier() {
        return Optional.ofNullable(this.caCertIdentifier);
    }

    /**
     * The identifier of the `aws.docdb.Cluster` in which to launch this instance.
     * 
     */
    @Import(name="clusterIdentifier", required=true)
    private Output<String> clusterIdentifier;

    /**
     * @return The identifier of the `aws.docdb.Cluster` in which to launch this instance.
     * 
     */
    public Output<String> clusterIdentifier() {
        return this.clusterIdentifier;
    }

    /**
     * Copy all DB instance `tags` to snapshots. Default is `false`.
     * 
     */
    @Import(name="copyTagsToSnapshot")
    private @Nullable Output<Boolean> copyTagsToSnapshot;

    /**
     * @return Copy all DB instance `tags` to snapshots. Default is `false`.
     * 
     */
    public Optional<Output<Boolean>> copyTagsToSnapshot() {
        return Optional.ofNullable(this.copyTagsToSnapshot);
    }

    /**
     * A value that indicates whether to enable Performance Insights for the DB Instance. Default `false`. See [docs] (https://docs.aws.amazon.com/documentdb/latest/developerguide/performance-insights.html) about the details.
     * 
     */
    @Import(name="enablePerformanceInsights")
    private @Nullable Output<Boolean> enablePerformanceInsights;

    /**
     * @return A value that indicates whether to enable Performance Insights for the DB Instance. Default `false`. See [docs] (https://docs.aws.amazon.com/documentdb/latest/developerguide/performance-insights.html) about the details.
     * 
     */
    public Optional<Output<Boolean>> enablePerformanceInsights() {
        return Optional.ofNullable(this.enablePerformanceInsights);
    }

    /**
     * The name of the database engine to be used for the DocumentDB instance. Defaults to `docdb`. Valid Values: `docdb`.
     * 
     */
    @Import(name="engine")
    private @Nullable Output<String> engine;

    /**
     * @return The name of the database engine to be used for the DocumentDB instance. Defaults to `docdb`. Valid Values: `docdb`.
     * 
     */
    public Optional<Output<String>> engine() {
        return Optional.ofNullable(this.engine);
    }

    /**
     * The identifier for the DocumentDB instance, if omitted, the provider will assign a random, unique identifier.
     * 
     */
    @Import(name="identifier")
    private @Nullable Output<String> identifier;

    /**
     * @return The identifier for the DocumentDB instance, if omitted, the provider will assign a random, unique identifier.
     * 
     */
    public Optional<Output<String>> identifier() {
        return Optional.ofNullable(this.identifier);
    }

    /**
     * Creates a unique identifier beginning with the specified prefix. Conflicts with `identifier`.
     * 
     */
    @Import(name="identifierPrefix")
    private @Nullable Output<String> identifierPrefix;

    /**
     * @return Creates a unique identifier beginning with the specified prefix. Conflicts with `identifier`.
     * 
     */
    public Optional<Output<String>> identifierPrefix() {
        return Optional.ofNullable(this.identifierPrefix);
    }

    /**
     * The instance class to use. For details on CPU and memory, see [Scaling for DocumentDB Instances](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-cluster-manage-performance.html#db-cluster-manage-scaling-instance).
     * DocumentDB currently supports the below instance classes.
     * Please see [AWS Documentation](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-instance-classes.html#db-instance-class-specs) for complete details.
     * - db.r6g.large
     * - db.r6g.xlarge
     * - db.r6g.2xlarge
     * - db.r6g.4xlarge
     * - db.r6g.8xlarge
     * - db.r6g.12xlarge
     * - db.r6g.16xlarge
     * - db.r5.large
     * - db.r5.xlarge
     * - db.r5.2xlarge
     * - db.r5.4xlarge
     * - db.r5.12xlarge
     * - db.r5.24xlarge
     * - db.r4.large
     * - db.r4.xlarge
     * - db.r4.2xlarge
     * - db.r4.4xlarge
     * - db.r4.8xlarge
     * - db.r4.16xlarge
     * - db.t4g.medium
     * - db.t3.medium
     * 
     */
    @Import(name="instanceClass", required=true)
    private Output<String> instanceClass;

    /**
     * @return The instance class to use. For details on CPU and memory, see [Scaling for DocumentDB Instances](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-cluster-manage-performance.html#db-cluster-manage-scaling-instance).
     * DocumentDB currently supports the below instance classes.
     * Please see [AWS Documentation](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-instance-classes.html#db-instance-class-specs) for complete details.
     * - db.r6g.large
     * - db.r6g.xlarge
     * - db.r6g.2xlarge
     * - db.r6g.4xlarge
     * - db.r6g.8xlarge
     * - db.r6g.12xlarge
     * - db.r6g.16xlarge
     * - db.r5.large
     * - db.r5.xlarge
     * - db.r5.2xlarge
     * - db.r5.4xlarge
     * - db.r5.12xlarge
     * - db.r5.24xlarge
     * - db.r4.large
     * - db.r4.xlarge
     * - db.r4.2xlarge
     * - db.r4.4xlarge
     * - db.r4.8xlarge
     * - db.r4.16xlarge
     * - db.t4g.medium
     * - db.t3.medium
     * 
     */
    public Output<String> instanceClass() {
        return this.instanceClass;
    }

    /**
     * The KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key. If you do not specify a value for PerformanceInsightsKMSKeyId, then Amazon DocumentDB uses your default KMS key.
     * 
     */
    @Import(name="performanceInsightsKmsKeyId")
    private @Nullable Output<String> performanceInsightsKmsKeyId;

    /**
     * @return The KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key. If you do not specify a value for PerformanceInsightsKMSKeyId, then Amazon DocumentDB uses your default KMS key.
     * 
     */
    public Optional<Output<String>> performanceInsightsKmsKeyId() {
        return Optional.ofNullable(this.performanceInsightsKmsKeyId);
    }

    /**
     * The window to perform maintenance in.
     * Syntax: &#34;ddd:hh24:mi-ddd:hh24:mi&#34;. Eg: &#34;Mon:00:00-Mon:03:00&#34;.
     * 
     */
    @Import(name="preferredMaintenanceWindow")
    private @Nullable Output<String> preferredMaintenanceWindow;

    /**
     * @return The window to perform maintenance in.
     * Syntax: &#34;ddd:hh24:mi-ddd:hh24:mi&#34;. Eg: &#34;Mon:00:00-Mon:03:00&#34;.
     * 
     */
    public Optional<Output<String>> preferredMaintenanceWindow() {
        return Optional.ofNullable(this.preferredMaintenanceWindow);
    }

    /**
     * Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
     * 
     */
    @Import(name="promotionTier")
    private @Nullable Output<Integer> promotionTier;

    /**
     * @return Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
     * 
     */
    public Optional<Output<Integer>> promotionTier() {
        return Optional.ofNullable(this.promotionTier);
    }

    /**
     * A map of tags to assign to the instance. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the instance. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private ClusterInstanceArgs() {}

    private ClusterInstanceArgs(ClusterInstanceArgs $) {
        this.applyImmediately = $.applyImmediately;
        this.autoMinorVersionUpgrade = $.autoMinorVersionUpgrade;
        this.availabilityZone = $.availabilityZone;
        this.caCertIdentifier = $.caCertIdentifier;
        this.clusterIdentifier = $.clusterIdentifier;
        this.copyTagsToSnapshot = $.copyTagsToSnapshot;
        this.enablePerformanceInsights = $.enablePerformanceInsights;
        this.engine = $.engine;
        this.identifier = $.identifier;
        this.identifierPrefix = $.identifierPrefix;
        this.instanceClass = $.instanceClass;
        this.performanceInsightsKmsKeyId = $.performanceInsightsKmsKeyId;
        this.preferredMaintenanceWindow = $.preferredMaintenanceWindow;
        this.promotionTier = $.promotionTier;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ClusterInstanceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ClusterInstanceArgs $;

        public Builder() {
            $ = new ClusterInstanceArgs();
        }

        public Builder(ClusterInstanceArgs defaults) {
            $ = new ClusterInstanceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param applyImmediately Specifies whether any database modifications
         * are applied immediately, or during the next maintenance window. Default is`false`.
         * 
         * @return builder
         * 
         */
        public Builder applyImmediately(@Nullable Output<Boolean> applyImmediately) {
            $.applyImmediately = applyImmediately;
            return this;
        }

        /**
         * @param applyImmediately Specifies whether any database modifications
         * are applied immediately, or during the next maintenance window. Default is`false`.
         * 
         * @return builder
         * 
         */
        public Builder applyImmediately(Boolean applyImmediately) {
            return applyImmediately(Output.of(applyImmediately));
        }

        /**
         * @param autoMinorVersionUpgrade This parameter does not apply to Amazon DocumentDB. Amazon DocumentDB does not perform minor version upgrades regardless of the value set (see [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBInstance.html)). Default `true`.
         * 
         * @return builder
         * 
         */
        public Builder autoMinorVersionUpgrade(@Nullable Output<Boolean> autoMinorVersionUpgrade) {
            $.autoMinorVersionUpgrade = autoMinorVersionUpgrade;
            return this;
        }

        /**
         * @param autoMinorVersionUpgrade This parameter does not apply to Amazon DocumentDB. Amazon DocumentDB does not perform minor version upgrades regardless of the value set (see [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBInstance.html)). Default `true`.
         * 
         * @return builder
         * 
         */
        public Builder autoMinorVersionUpgrade(Boolean autoMinorVersionUpgrade) {
            return autoMinorVersionUpgrade(Output.of(autoMinorVersionUpgrade));
        }

        /**
         * @param availabilityZone The EC2 Availability Zone that the DB instance is created in. See [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_CreateDBInstance.html) about the details.
         * 
         * @return builder
         * 
         */
        public Builder availabilityZone(@Nullable Output<String> availabilityZone) {
            $.availabilityZone = availabilityZone;
            return this;
        }

        /**
         * @param availabilityZone The EC2 Availability Zone that the DB instance is created in. See [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_CreateDBInstance.html) about the details.
         * 
         * @return builder
         * 
         */
        public Builder availabilityZone(String availabilityZone) {
            return availabilityZone(Output.of(availabilityZone));
        }

        /**
         * @param caCertIdentifier (Optional) The identifier of the CA certificate for the DB instance.
         * 
         * @return builder
         * 
         */
        public Builder caCertIdentifier(@Nullable Output<String> caCertIdentifier) {
            $.caCertIdentifier = caCertIdentifier;
            return this;
        }

        /**
         * @param caCertIdentifier (Optional) The identifier of the CA certificate for the DB instance.
         * 
         * @return builder
         * 
         */
        public Builder caCertIdentifier(String caCertIdentifier) {
            return caCertIdentifier(Output.of(caCertIdentifier));
        }

        /**
         * @param clusterIdentifier The identifier of the `aws.docdb.Cluster` in which to launch this instance.
         * 
         * @return builder
         * 
         */
        public Builder clusterIdentifier(Output<String> clusterIdentifier) {
            $.clusterIdentifier = clusterIdentifier;
            return this;
        }

        /**
         * @param clusterIdentifier The identifier of the `aws.docdb.Cluster` in which to launch this instance.
         * 
         * @return builder
         * 
         */
        public Builder clusterIdentifier(String clusterIdentifier) {
            return clusterIdentifier(Output.of(clusterIdentifier));
        }

        /**
         * @param copyTagsToSnapshot Copy all DB instance `tags` to snapshots. Default is `false`.
         * 
         * @return builder
         * 
         */
        public Builder copyTagsToSnapshot(@Nullable Output<Boolean> copyTagsToSnapshot) {
            $.copyTagsToSnapshot = copyTagsToSnapshot;
            return this;
        }

        /**
         * @param copyTagsToSnapshot Copy all DB instance `tags` to snapshots. Default is `false`.
         * 
         * @return builder
         * 
         */
        public Builder copyTagsToSnapshot(Boolean copyTagsToSnapshot) {
            return copyTagsToSnapshot(Output.of(copyTagsToSnapshot));
        }

        /**
         * @param enablePerformanceInsights A value that indicates whether to enable Performance Insights for the DB Instance. Default `false`. See [docs] (https://docs.aws.amazon.com/documentdb/latest/developerguide/performance-insights.html) about the details.
         * 
         * @return builder
         * 
         */
        public Builder enablePerformanceInsights(@Nullable Output<Boolean> enablePerformanceInsights) {
            $.enablePerformanceInsights = enablePerformanceInsights;
            return this;
        }

        /**
         * @param enablePerformanceInsights A value that indicates whether to enable Performance Insights for the DB Instance. Default `false`. See [docs] (https://docs.aws.amazon.com/documentdb/latest/developerguide/performance-insights.html) about the details.
         * 
         * @return builder
         * 
         */
        public Builder enablePerformanceInsights(Boolean enablePerformanceInsights) {
            return enablePerformanceInsights(Output.of(enablePerformanceInsights));
        }

        /**
         * @param engine The name of the database engine to be used for the DocumentDB instance. Defaults to `docdb`. Valid Values: `docdb`.
         * 
         * @return builder
         * 
         */
        public Builder engine(@Nullable Output<String> engine) {
            $.engine = engine;
            return this;
        }

        /**
         * @param engine The name of the database engine to be used for the DocumentDB instance. Defaults to `docdb`. Valid Values: `docdb`.
         * 
         * @return builder
         * 
         */
        public Builder engine(String engine) {
            return engine(Output.of(engine));
        }

        /**
         * @param identifier The identifier for the DocumentDB instance, if omitted, the provider will assign a random, unique identifier.
         * 
         * @return builder
         * 
         */
        public Builder identifier(@Nullable Output<String> identifier) {
            $.identifier = identifier;
            return this;
        }

        /**
         * @param identifier The identifier for the DocumentDB instance, if omitted, the provider will assign a random, unique identifier.
         * 
         * @return builder
         * 
         */
        public Builder identifier(String identifier) {
            return identifier(Output.of(identifier));
        }

        /**
         * @param identifierPrefix Creates a unique identifier beginning with the specified prefix. Conflicts with `identifier`.
         * 
         * @return builder
         * 
         */
        public Builder identifierPrefix(@Nullable Output<String> identifierPrefix) {
            $.identifierPrefix = identifierPrefix;
            return this;
        }

        /**
         * @param identifierPrefix Creates a unique identifier beginning with the specified prefix. Conflicts with `identifier`.
         * 
         * @return builder
         * 
         */
        public Builder identifierPrefix(String identifierPrefix) {
            return identifierPrefix(Output.of(identifierPrefix));
        }

        /**
         * @param instanceClass The instance class to use. For details on CPU and memory, see [Scaling for DocumentDB Instances](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-cluster-manage-performance.html#db-cluster-manage-scaling-instance).
         * DocumentDB currently supports the below instance classes.
         * Please see [AWS Documentation](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-instance-classes.html#db-instance-class-specs) for complete details.
         * - db.r6g.large
         * - db.r6g.xlarge
         * - db.r6g.2xlarge
         * - db.r6g.4xlarge
         * - db.r6g.8xlarge
         * - db.r6g.12xlarge
         * - db.r6g.16xlarge
         * - db.r5.large
         * - db.r5.xlarge
         * - db.r5.2xlarge
         * - db.r5.4xlarge
         * - db.r5.12xlarge
         * - db.r5.24xlarge
         * - db.r4.large
         * - db.r4.xlarge
         * - db.r4.2xlarge
         * - db.r4.4xlarge
         * - db.r4.8xlarge
         * - db.r4.16xlarge
         * - db.t4g.medium
         * - db.t3.medium
         * 
         * @return builder
         * 
         */
        public Builder instanceClass(Output<String> instanceClass) {
            $.instanceClass = instanceClass;
            return this;
        }

        /**
         * @param instanceClass The instance class to use. For details on CPU and memory, see [Scaling for DocumentDB Instances](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-cluster-manage-performance.html#db-cluster-manage-scaling-instance).
         * DocumentDB currently supports the below instance classes.
         * Please see [AWS Documentation](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-instance-classes.html#db-instance-class-specs) for complete details.
         * - db.r6g.large
         * - db.r6g.xlarge
         * - db.r6g.2xlarge
         * - db.r6g.4xlarge
         * - db.r6g.8xlarge
         * - db.r6g.12xlarge
         * - db.r6g.16xlarge
         * - db.r5.large
         * - db.r5.xlarge
         * - db.r5.2xlarge
         * - db.r5.4xlarge
         * - db.r5.12xlarge
         * - db.r5.24xlarge
         * - db.r4.large
         * - db.r4.xlarge
         * - db.r4.2xlarge
         * - db.r4.4xlarge
         * - db.r4.8xlarge
         * - db.r4.16xlarge
         * - db.t4g.medium
         * - db.t3.medium
         * 
         * @return builder
         * 
         */
        public Builder instanceClass(String instanceClass) {
            return instanceClass(Output.of(instanceClass));
        }

        /**
         * @param performanceInsightsKmsKeyId The KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key. If you do not specify a value for PerformanceInsightsKMSKeyId, then Amazon DocumentDB uses your default KMS key.
         * 
         * @return builder
         * 
         */
        public Builder performanceInsightsKmsKeyId(@Nullable Output<String> performanceInsightsKmsKeyId) {
            $.performanceInsightsKmsKeyId = performanceInsightsKmsKeyId;
            return this;
        }

        /**
         * @param performanceInsightsKmsKeyId The KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key. If you do not specify a value for PerformanceInsightsKMSKeyId, then Amazon DocumentDB uses your default KMS key.
         * 
         * @return builder
         * 
         */
        public Builder performanceInsightsKmsKeyId(String performanceInsightsKmsKeyId) {
            return performanceInsightsKmsKeyId(Output.of(performanceInsightsKmsKeyId));
        }

        /**
         * @param preferredMaintenanceWindow The window to perform maintenance in.
         * Syntax: &#34;ddd:hh24:mi-ddd:hh24:mi&#34;. Eg: &#34;Mon:00:00-Mon:03:00&#34;.
         * 
         * @return builder
         * 
         */
        public Builder preferredMaintenanceWindow(@Nullable Output<String> preferredMaintenanceWindow) {
            $.preferredMaintenanceWindow = preferredMaintenanceWindow;
            return this;
        }

        /**
         * @param preferredMaintenanceWindow The window to perform maintenance in.
         * Syntax: &#34;ddd:hh24:mi-ddd:hh24:mi&#34;. Eg: &#34;Mon:00:00-Mon:03:00&#34;.
         * 
         * @return builder
         * 
         */
        public Builder preferredMaintenanceWindow(String preferredMaintenanceWindow) {
            return preferredMaintenanceWindow(Output.of(preferredMaintenanceWindow));
        }

        /**
         * @param promotionTier Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
         * 
         * @return builder
         * 
         */
        public Builder promotionTier(@Nullable Output<Integer> promotionTier) {
            $.promotionTier = promotionTier;
            return this;
        }

        /**
         * @param promotionTier Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
         * 
         * @return builder
         * 
         */
        public Builder promotionTier(Integer promotionTier) {
            return promotionTier(Output.of(promotionTier));
        }

        /**
         * @param tags A map of tags to assign to the instance. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A map of tags to assign to the instance. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        public ClusterInstanceArgs build() {
            $.clusterIdentifier = Objects.requireNonNull($.clusterIdentifier, "expected parameter 'clusterIdentifier' to be non-null");
            $.instanceClass = Objects.requireNonNull($.instanceClass, "expected parameter 'instanceClass' to be non-null");
            return $;
        }
    }

}
