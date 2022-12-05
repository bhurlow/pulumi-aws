// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.neptune;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GlobalClusterArgs extends com.pulumi.resources.ResourceArgs {

    public static final GlobalClusterArgs Empty = new GlobalClusterArgs();

    /**
     * If the Global Cluster should have deletion protection enabled. The database can&#39;t be deleted when this value is set to `true`. The default is `false`.
     * 
     */
    @Import(name="deletionProtection")
    private @Nullable Output<Boolean> deletionProtection;

    /**
     * @return If the Global Cluster should have deletion protection enabled. The database can&#39;t be deleted when this value is set to `true`. The default is `false`.
     * 
     */
    public Optional<Output<Boolean>> deletionProtection() {
        return Optional.ofNullable(this.deletionProtection);
    }

    /**
     * Name of the database engine to be used for this DB cluster. The provider will only perform drift detection if a configuration value is provided. Current Valid values: `neptune`. Conflicts with `source_db_cluster_identifier`.
     * 
     */
    @Import(name="engine")
    private @Nullable Output<String> engine;

    /**
     * @return Name of the database engine to be used for this DB cluster. The provider will only perform drift detection if a configuration value is provided. Current Valid values: `neptune`. Conflicts with `source_db_cluster_identifier`.
     * 
     */
    public Optional<Output<String>> engine() {
        return Optional.ofNullable(this.engine);
    }

    /**
     * Engine version of the global database. Upgrading the engine version will result in all cluster members being immediately updated and will.
     * * **NOTE:** Upgrading major versions is not supported.
     * 
     */
    @Import(name="engineVersion")
    private @Nullable Output<String> engineVersion;

    /**
     * @return Engine version of the global database. Upgrading the engine version will result in all cluster members being immediately updated and will.
     * * **NOTE:** Upgrading major versions is not supported.
     * 
     */
    public Optional<Output<String>> engineVersion() {
        return Optional.ofNullable(this.engineVersion);
    }

    /**
     * The global cluster identifier.
     * 
     */
    @Import(name="globalClusterIdentifier", required=true)
    private Output<String> globalClusterIdentifier;

    /**
     * @return The global cluster identifier.
     * 
     */
    public Output<String> globalClusterIdentifier() {
        return this.globalClusterIdentifier;
    }

    /**
     * Amazon Resource Name (ARN) to use as the primary DB Cluster of the Global Cluster on creation. The provider cannot perform drift detection of this value.
     * 
     */
    @Import(name="sourceDbClusterIdentifier")
    private @Nullable Output<String> sourceDbClusterIdentifier;

    /**
     * @return Amazon Resource Name (ARN) to use as the primary DB Cluster of the Global Cluster on creation. The provider cannot perform drift detection of this value.
     * 
     */
    public Optional<Output<String>> sourceDbClusterIdentifier() {
        return Optional.ofNullable(this.sourceDbClusterIdentifier);
    }

    /**
     * Specifies whether the DB cluster is encrypted. The default is `false` unless `source_db_cluster_identifier` is specified and encrypted. The provider will only perform drift detection if a configuration value is provided.
     * 
     */
    @Import(name="storageEncrypted")
    private @Nullable Output<Boolean> storageEncrypted;

    /**
     * @return Specifies whether the DB cluster is encrypted. The default is `false` unless `source_db_cluster_identifier` is specified and encrypted. The provider will only perform drift detection if a configuration value is provided.
     * 
     */
    public Optional<Output<Boolean>> storageEncrypted() {
        return Optional.ofNullable(this.storageEncrypted);
    }

    private GlobalClusterArgs() {}

    private GlobalClusterArgs(GlobalClusterArgs $) {
        this.deletionProtection = $.deletionProtection;
        this.engine = $.engine;
        this.engineVersion = $.engineVersion;
        this.globalClusterIdentifier = $.globalClusterIdentifier;
        this.sourceDbClusterIdentifier = $.sourceDbClusterIdentifier;
        this.storageEncrypted = $.storageEncrypted;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GlobalClusterArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GlobalClusterArgs $;

        public Builder() {
            $ = new GlobalClusterArgs();
        }

        public Builder(GlobalClusterArgs defaults) {
            $ = new GlobalClusterArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param deletionProtection If the Global Cluster should have deletion protection enabled. The database can&#39;t be deleted when this value is set to `true`. The default is `false`.
         * 
         * @return builder
         * 
         */
        public Builder deletionProtection(@Nullable Output<Boolean> deletionProtection) {
            $.deletionProtection = deletionProtection;
            return this;
        }

        /**
         * @param deletionProtection If the Global Cluster should have deletion protection enabled. The database can&#39;t be deleted when this value is set to `true`. The default is `false`.
         * 
         * @return builder
         * 
         */
        public Builder deletionProtection(Boolean deletionProtection) {
            return deletionProtection(Output.of(deletionProtection));
        }

        /**
         * @param engine Name of the database engine to be used for this DB cluster. The provider will only perform drift detection if a configuration value is provided. Current Valid values: `neptune`. Conflicts with `source_db_cluster_identifier`.
         * 
         * @return builder
         * 
         */
        public Builder engine(@Nullable Output<String> engine) {
            $.engine = engine;
            return this;
        }

        /**
         * @param engine Name of the database engine to be used for this DB cluster. The provider will only perform drift detection if a configuration value is provided. Current Valid values: `neptune`. Conflicts with `source_db_cluster_identifier`.
         * 
         * @return builder
         * 
         */
        public Builder engine(String engine) {
            return engine(Output.of(engine));
        }

        /**
         * @param engineVersion Engine version of the global database. Upgrading the engine version will result in all cluster members being immediately updated and will.
         * * **NOTE:** Upgrading major versions is not supported.
         * 
         * @return builder
         * 
         */
        public Builder engineVersion(@Nullable Output<String> engineVersion) {
            $.engineVersion = engineVersion;
            return this;
        }

        /**
         * @param engineVersion Engine version of the global database. Upgrading the engine version will result in all cluster members being immediately updated and will.
         * * **NOTE:** Upgrading major versions is not supported.
         * 
         * @return builder
         * 
         */
        public Builder engineVersion(String engineVersion) {
            return engineVersion(Output.of(engineVersion));
        }

        /**
         * @param globalClusterIdentifier The global cluster identifier.
         * 
         * @return builder
         * 
         */
        public Builder globalClusterIdentifier(Output<String> globalClusterIdentifier) {
            $.globalClusterIdentifier = globalClusterIdentifier;
            return this;
        }

        /**
         * @param globalClusterIdentifier The global cluster identifier.
         * 
         * @return builder
         * 
         */
        public Builder globalClusterIdentifier(String globalClusterIdentifier) {
            return globalClusterIdentifier(Output.of(globalClusterIdentifier));
        }

        /**
         * @param sourceDbClusterIdentifier Amazon Resource Name (ARN) to use as the primary DB Cluster of the Global Cluster on creation. The provider cannot perform drift detection of this value.
         * 
         * @return builder
         * 
         */
        public Builder sourceDbClusterIdentifier(@Nullable Output<String> sourceDbClusterIdentifier) {
            $.sourceDbClusterIdentifier = sourceDbClusterIdentifier;
            return this;
        }

        /**
         * @param sourceDbClusterIdentifier Amazon Resource Name (ARN) to use as the primary DB Cluster of the Global Cluster on creation. The provider cannot perform drift detection of this value.
         * 
         * @return builder
         * 
         */
        public Builder sourceDbClusterIdentifier(String sourceDbClusterIdentifier) {
            return sourceDbClusterIdentifier(Output.of(sourceDbClusterIdentifier));
        }

        /**
         * @param storageEncrypted Specifies whether the DB cluster is encrypted. The default is `false` unless `source_db_cluster_identifier` is specified and encrypted. The provider will only perform drift detection if a configuration value is provided.
         * 
         * @return builder
         * 
         */
        public Builder storageEncrypted(@Nullable Output<Boolean> storageEncrypted) {
            $.storageEncrypted = storageEncrypted;
            return this;
        }

        /**
         * @param storageEncrypted Specifies whether the DB cluster is encrypted. The default is `false` unless `source_db_cluster_identifier` is specified and encrypted. The provider will only perform drift detection if a configuration value is provided.
         * 
         * @return builder
         * 
         */
        public Builder storageEncrypted(Boolean storageEncrypted) {
            return storageEncrypted(Output.of(storageEncrypted));
        }

        public GlobalClusterArgs build() {
            $.globalClusterIdentifier = Objects.requireNonNull($.globalClusterIdentifier, "expected parameter 'globalClusterIdentifier' to be non-null");
            return $;
        }
    }

}
