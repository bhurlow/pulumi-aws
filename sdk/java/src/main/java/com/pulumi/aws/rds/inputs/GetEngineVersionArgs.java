// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.rds.inputs;

import com.pulumi.aws.rds.inputs.GetEngineVersionFilterArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetEngineVersionArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetEngineVersionArgs Empty = new GetEngineVersionArgs();

    /**
     * When set to `true`, the default version for the specified `engine` or combination of `engine` and major `version` will be returned. Can be used to limit responses to a single version when they would otherwise fail for returning multiple versions.
     * 
     */
    @Import(name="defaultOnly")
    private @Nullable Output<Boolean> defaultOnly;

    /**
     * @return When set to `true`, the default version for the specified `engine` or combination of `engine` and major `version` will be returned. Can be used to limit responses to a single version when they would otherwise fail for returning multiple versions.
     * 
     */
    public Optional<Output<Boolean>> defaultOnly() {
        return Optional.ofNullable(this.defaultOnly);
    }

    /**
     * Database engine. Engine values include `aurora`, `aurora-mysql`, `aurora-postgresql`, `docdb`, `mariadb`, `mysql`, `neptune`, `oracle-ee`, `oracle-se`, `oracle-se1`, `oracle-se2`, `postgres`, `sqlserver-ee`, `sqlserver-ex`, `sqlserver-se`, and `sqlserver-web`.
     * 
     * The following arguments are optional:
     * 
     */
    @Import(name="engine", required=true)
    private Output<String> engine;

    /**
     * @return Database engine. Engine values include `aurora`, `aurora-mysql`, `aurora-postgresql`, `docdb`, `mariadb`, `mysql`, `neptune`, `oracle-ee`, `oracle-se`, `oracle-se1`, `oracle-se2`, `postgres`, `sqlserver-ee`, `sqlserver-ex`, `sqlserver-se`, and `sqlserver-web`.
     * 
     * The following arguments are optional:
     * 
     */
    public Output<String> engine() {
        return this.engine;
    }

    /**
     * One or more name/value pairs to filter off of. There are several valid keys; for a full reference, check out [describe-db-engine-versions in the AWS CLI reference](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/rds/describe-db-engine-versions.html).
     * 
     */
    @Import(name="filters")
    private @Nullable Output<List<GetEngineVersionFilterArgs>> filters;

    /**
     * @return One or more name/value pairs to filter off of. There are several valid keys; for a full reference, check out [describe-db-engine-versions in the AWS CLI reference](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/rds/describe-db-engine-versions.html).
     * 
     */
    public Optional<Output<List<GetEngineVersionFilterArgs>>> filters() {
        return Optional.ofNullable(this.filters);
    }

    /**
     * When set to `true`, the specified `version` or member of `preferred_versions` will be returned even if it is `deprecated`. Otherwise, only `available` versions will be returned.
     * 
     */
    @Import(name="includeAll")
    private @Nullable Output<Boolean> includeAll;

    /**
     * @return When set to `true`, the specified `version` or member of `preferred_versions` will be returned even if it is `deprecated`. Otherwise, only `available` versions will be returned.
     * 
     */
    public Optional<Output<Boolean>> includeAll() {
        return Optional.ofNullable(this.includeAll);
    }

    /**
     * When set to `true`, the data source attempts to return the most recent version matching the other criteria you provide. This differs from `default_only`. For example, the latest version is not always the default. In addition, AWS may return multiple defaults depending on the criteria. Using `latest` will avoid `multiple RDS engine versions` errors. **Note:** The data source uses a best-effort approach at selecting the latest version but due to the complexity of version identifiers across engines and incomplete version date information provided by AWS, using `latest` may _not_ return the latest version in every situation.
     * 
     */
    @Import(name="latest")
    private @Nullable Output<Boolean> latest;

    /**
     * @return When set to `true`, the data source attempts to return the most recent version matching the other criteria you provide. This differs from `default_only`. For example, the latest version is not always the default. In addition, AWS may return multiple defaults depending on the criteria. Using `latest` will avoid `multiple RDS engine versions` errors. **Note:** The data source uses a best-effort approach at selecting the latest version but due to the complexity of version identifiers across engines and incomplete version date information provided by AWS, using `latest` may _not_ return the latest version in every situation.
     * 
     */
    public Optional<Output<Boolean>> latest() {
        return Optional.ofNullable(this.latest);
    }

    /**
     * Name of a specific database parameter group family. Examples of parameter group families are `mysql8.0`, `mariadb10.4`, and `postgres12`.
     * 
     */
    @Import(name="parameterGroupFamily")
    private @Nullable Output<String> parameterGroupFamily;

    /**
     * @return Name of a specific database parameter group family. Examples of parameter group families are `mysql8.0`, `mariadb10.4`, and `postgres12`.
     * 
     */
    public Optional<Output<String>> parameterGroupFamily() {
        return Optional.ofNullable(this.parameterGroupFamily);
    }

    /**
     * Ordered list of preferred major version upgrade targets. The version corresponding to the first match in this list will be returned unless the `latest` parameter is set to `true`. If you don&#39;t configure `version`, `preferred_major_targets`, `preferred_upgrade_targets`, and `preferred_versions`, the data source will return the default version for the engine. You can use this with other version criteria.
     * 
     */
    @Import(name="preferredMajorTargets")
    private @Nullable Output<List<String>> preferredMajorTargets;

    /**
     * @return Ordered list of preferred major version upgrade targets. The version corresponding to the first match in this list will be returned unless the `latest` parameter is set to `true`. If you don&#39;t configure `version`, `preferred_major_targets`, `preferred_upgrade_targets`, and `preferred_versions`, the data source will return the default version for the engine. You can use this with other version criteria.
     * 
     */
    public Optional<Output<List<String>>> preferredMajorTargets() {
        return Optional.ofNullable(this.preferredMajorTargets);
    }

    /**
     * Ordered list of preferred version upgrade targets. The version corresponding to the first match in this list will be returned unless the `latest` parameter is set to `true`. If you don&#39;t configure `version`, `preferred_major_targets`, `preferred_upgrade_targets`, and `preferred_versions`, the data source will return the default version for the engine. You can use this with other version criteria.
     * 
     */
    @Import(name="preferredUpgradeTargets")
    private @Nullable Output<List<String>> preferredUpgradeTargets;

    /**
     * @return Ordered list of preferred version upgrade targets. The version corresponding to the first match in this list will be returned unless the `latest` parameter is set to `true`. If you don&#39;t configure `version`, `preferred_major_targets`, `preferred_upgrade_targets`, and `preferred_versions`, the data source will return the default version for the engine. You can use this with other version criteria.
     * 
     */
    public Optional<Output<List<String>>> preferredUpgradeTargets() {
        return Optional.ofNullable(this.preferredUpgradeTargets);
    }

    /**
     * Ordered list of preferred versions. The first match in this list that matches any other criteria will be returned unless the `latest` parameter is set to `true`. If you don&#39;t configure `version`, `preferred_major_targets`, `preferred_upgrade_targets`, and `preferred_versions`, the data source will return the default version for the engine. You can use this with other version criteria.
     * 
     */
    @Import(name="preferredVersions")
    private @Nullable Output<List<String>> preferredVersions;

    /**
     * @return Ordered list of preferred versions. The first match in this list that matches any other criteria will be returned unless the `latest` parameter is set to `true`. If you don&#39;t configure `version`, `preferred_major_targets`, `preferred_upgrade_targets`, and `preferred_versions`, the data source will return the default version for the engine. You can use this with other version criteria.
     * 
     */
    public Optional<Output<List<String>>> preferredVersions() {
        return Optional.ofNullable(this.preferredVersions);
    }

    @Import(name="version")
    private @Nullable Output<String> version;

    public Optional<Output<String>> version() {
        return Optional.ofNullable(this.version);
    }

    private GetEngineVersionArgs() {}

    private GetEngineVersionArgs(GetEngineVersionArgs $) {
        this.defaultOnly = $.defaultOnly;
        this.engine = $.engine;
        this.filters = $.filters;
        this.includeAll = $.includeAll;
        this.latest = $.latest;
        this.parameterGroupFamily = $.parameterGroupFamily;
        this.preferredMajorTargets = $.preferredMajorTargets;
        this.preferredUpgradeTargets = $.preferredUpgradeTargets;
        this.preferredVersions = $.preferredVersions;
        this.version = $.version;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetEngineVersionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetEngineVersionArgs $;

        public Builder() {
            $ = new GetEngineVersionArgs();
        }

        public Builder(GetEngineVersionArgs defaults) {
            $ = new GetEngineVersionArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param defaultOnly When set to `true`, the default version for the specified `engine` or combination of `engine` and major `version` will be returned. Can be used to limit responses to a single version when they would otherwise fail for returning multiple versions.
         * 
         * @return builder
         * 
         */
        public Builder defaultOnly(@Nullable Output<Boolean> defaultOnly) {
            $.defaultOnly = defaultOnly;
            return this;
        }

        /**
         * @param defaultOnly When set to `true`, the default version for the specified `engine` or combination of `engine` and major `version` will be returned. Can be used to limit responses to a single version when they would otherwise fail for returning multiple versions.
         * 
         * @return builder
         * 
         */
        public Builder defaultOnly(Boolean defaultOnly) {
            return defaultOnly(Output.of(defaultOnly));
        }

        /**
         * @param engine Database engine. Engine values include `aurora`, `aurora-mysql`, `aurora-postgresql`, `docdb`, `mariadb`, `mysql`, `neptune`, `oracle-ee`, `oracle-se`, `oracle-se1`, `oracle-se2`, `postgres`, `sqlserver-ee`, `sqlserver-ex`, `sqlserver-se`, and `sqlserver-web`.
         * 
         * The following arguments are optional:
         * 
         * @return builder
         * 
         */
        public Builder engine(Output<String> engine) {
            $.engine = engine;
            return this;
        }

        /**
         * @param engine Database engine. Engine values include `aurora`, `aurora-mysql`, `aurora-postgresql`, `docdb`, `mariadb`, `mysql`, `neptune`, `oracle-ee`, `oracle-se`, `oracle-se1`, `oracle-se2`, `postgres`, `sqlserver-ee`, `sqlserver-ex`, `sqlserver-se`, and `sqlserver-web`.
         * 
         * The following arguments are optional:
         * 
         * @return builder
         * 
         */
        public Builder engine(String engine) {
            return engine(Output.of(engine));
        }

        /**
         * @param filters One or more name/value pairs to filter off of. There are several valid keys; for a full reference, check out [describe-db-engine-versions in the AWS CLI reference](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/rds/describe-db-engine-versions.html).
         * 
         * @return builder
         * 
         */
        public Builder filters(@Nullable Output<List<GetEngineVersionFilterArgs>> filters) {
            $.filters = filters;
            return this;
        }

        /**
         * @param filters One or more name/value pairs to filter off of. There are several valid keys; for a full reference, check out [describe-db-engine-versions in the AWS CLI reference](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/rds/describe-db-engine-versions.html).
         * 
         * @return builder
         * 
         */
        public Builder filters(List<GetEngineVersionFilterArgs> filters) {
            return filters(Output.of(filters));
        }

        /**
         * @param filters One or more name/value pairs to filter off of. There are several valid keys; for a full reference, check out [describe-db-engine-versions in the AWS CLI reference](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/rds/describe-db-engine-versions.html).
         * 
         * @return builder
         * 
         */
        public Builder filters(GetEngineVersionFilterArgs... filters) {
            return filters(List.of(filters));
        }

        /**
         * @param includeAll When set to `true`, the specified `version` or member of `preferred_versions` will be returned even if it is `deprecated`. Otherwise, only `available` versions will be returned.
         * 
         * @return builder
         * 
         */
        public Builder includeAll(@Nullable Output<Boolean> includeAll) {
            $.includeAll = includeAll;
            return this;
        }

        /**
         * @param includeAll When set to `true`, the specified `version` or member of `preferred_versions` will be returned even if it is `deprecated`. Otherwise, only `available` versions will be returned.
         * 
         * @return builder
         * 
         */
        public Builder includeAll(Boolean includeAll) {
            return includeAll(Output.of(includeAll));
        }

        /**
         * @param latest When set to `true`, the data source attempts to return the most recent version matching the other criteria you provide. This differs from `default_only`. For example, the latest version is not always the default. In addition, AWS may return multiple defaults depending on the criteria. Using `latest` will avoid `multiple RDS engine versions` errors. **Note:** The data source uses a best-effort approach at selecting the latest version but due to the complexity of version identifiers across engines and incomplete version date information provided by AWS, using `latest` may _not_ return the latest version in every situation.
         * 
         * @return builder
         * 
         */
        public Builder latest(@Nullable Output<Boolean> latest) {
            $.latest = latest;
            return this;
        }

        /**
         * @param latest When set to `true`, the data source attempts to return the most recent version matching the other criteria you provide. This differs from `default_only`. For example, the latest version is not always the default. In addition, AWS may return multiple defaults depending on the criteria. Using `latest` will avoid `multiple RDS engine versions` errors. **Note:** The data source uses a best-effort approach at selecting the latest version but due to the complexity of version identifiers across engines and incomplete version date information provided by AWS, using `latest` may _not_ return the latest version in every situation.
         * 
         * @return builder
         * 
         */
        public Builder latest(Boolean latest) {
            return latest(Output.of(latest));
        }

        /**
         * @param parameterGroupFamily Name of a specific database parameter group family. Examples of parameter group families are `mysql8.0`, `mariadb10.4`, and `postgres12`.
         * 
         * @return builder
         * 
         */
        public Builder parameterGroupFamily(@Nullable Output<String> parameterGroupFamily) {
            $.parameterGroupFamily = parameterGroupFamily;
            return this;
        }

        /**
         * @param parameterGroupFamily Name of a specific database parameter group family. Examples of parameter group families are `mysql8.0`, `mariadb10.4`, and `postgres12`.
         * 
         * @return builder
         * 
         */
        public Builder parameterGroupFamily(String parameterGroupFamily) {
            return parameterGroupFamily(Output.of(parameterGroupFamily));
        }

        /**
         * @param preferredMajorTargets Ordered list of preferred major version upgrade targets. The version corresponding to the first match in this list will be returned unless the `latest` parameter is set to `true`. If you don&#39;t configure `version`, `preferred_major_targets`, `preferred_upgrade_targets`, and `preferred_versions`, the data source will return the default version for the engine. You can use this with other version criteria.
         * 
         * @return builder
         * 
         */
        public Builder preferredMajorTargets(@Nullable Output<List<String>> preferredMajorTargets) {
            $.preferredMajorTargets = preferredMajorTargets;
            return this;
        }

        /**
         * @param preferredMajorTargets Ordered list of preferred major version upgrade targets. The version corresponding to the first match in this list will be returned unless the `latest` parameter is set to `true`. If you don&#39;t configure `version`, `preferred_major_targets`, `preferred_upgrade_targets`, and `preferred_versions`, the data source will return the default version for the engine. You can use this with other version criteria.
         * 
         * @return builder
         * 
         */
        public Builder preferredMajorTargets(List<String> preferredMajorTargets) {
            return preferredMajorTargets(Output.of(preferredMajorTargets));
        }

        /**
         * @param preferredMajorTargets Ordered list of preferred major version upgrade targets. The version corresponding to the first match in this list will be returned unless the `latest` parameter is set to `true`. If you don&#39;t configure `version`, `preferred_major_targets`, `preferred_upgrade_targets`, and `preferred_versions`, the data source will return the default version for the engine. You can use this with other version criteria.
         * 
         * @return builder
         * 
         */
        public Builder preferredMajorTargets(String... preferredMajorTargets) {
            return preferredMajorTargets(List.of(preferredMajorTargets));
        }

        /**
         * @param preferredUpgradeTargets Ordered list of preferred version upgrade targets. The version corresponding to the first match in this list will be returned unless the `latest` parameter is set to `true`. If you don&#39;t configure `version`, `preferred_major_targets`, `preferred_upgrade_targets`, and `preferred_versions`, the data source will return the default version for the engine. You can use this with other version criteria.
         * 
         * @return builder
         * 
         */
        public Builder preferredUpgradeTargets(@Nullable Output<List<String>> preferredUpgradeTargets) {
            $.preferredUpgradeTargets = preferredUpgradeTargets;
            return this;
        }

        /**
         * @param preferredUpgradeTargets Ordered list of preferred version upgrade targets. The version corresponding to the first match in this list will be returned unless the `latest` parameter is set to `true`. If you don&#39;t configure `version`, `preferred_major_targets`, `preferred_upgrade_targets`, and `preferred_versions`, the data source will return the default version for the engine. You can use this with other version criteria.
         * 
         * @return builder
         * 
         */
        public Builder preferredUpgradeTargets(List<String> preferredUpgradeTargets) {
            return preferredUpgradeTargets(Output.of(preferredUpgradeTargets));
        }

        /**
         * @param preferredUpgradeTargets Ordered list of preferred version upgrade targets. The version corresponding to the first match in this list will be returned unless the `latest` parameter is set to `true`. If you don&#39;t configure `version`, `preferred_major_targets`, `preferred_upgrade_targets`, and `preferred_versions`, the data source will return the default version for the engine. You can use this with other version criteria.
         * 
         * @return builder
         * 
         */
        public Builder preferredUpgradeTargets(String... preferredUpgradeTargets) {
            return preferredUpgradeTargets(List.of(preferredUpgradeTargets));
        }

        /**
         * @param preferredVersions Ordered list of preferred versions. The first match in this list that matches any other criteria will be returned unless the `latest` parameter is set to `true`. If you don&#39;t configure `version`, `preferred_major_targets`, `preferred_upgrade_targets`, and `preferred_versions`, the data source will return the default version for the engine. You can use this with other version criteria.
         * 
         * @return builder
         * 
         */
        public Builder preferredVersions(@Nullable Output<List<String>> preferredVersions) {
            $.preferredVersions = preferredVersions;
            return this;
        }

        /**
         * @param preferredVersions Ordered list of preferred versions. The first match in this list that matches any other criteria will be returned unless the `latest` parameter is set to `true`. If you don&#39;t configure `version`, `preferred_major_targets`, `preferred_upgrade_targets`, and `preferred_versions`, the data source will return the default version for the engine. You can use this with other version criteria.
         * 
         * @return builder
         * 
         */
        public Builder preferredVersions(List<String> preferredVersions) {
            return preferredVersions(Output.of(preferredVersions));
        }

        /**
         * @param preferredVersions Ordered list of preferred versions. The first match in this list that matches any other criteria will be returned unless the `latest` parameter is set to `true`. If you don&#39;t configure `version`, `preferred_major_targets`, `preferred_upgrade_targets`, and `preferred_versions`, the data source will return the default version for the engine. You can use this with other version criteria.
         * 
         * @return builder
         * 
         */
        public Builder preferredVersions(String... preferredVersions) {
            return preferredVersions(List.of(preferredVersions));
        }

        public Builder version(@Nullable Output<String> version) {
            $.version = version;
            return this;
        }

        public Builder version(String version) {
            return version(Output.of(version));
        }

        public GetEngineVersionArgs build() {
            if ($.engine == null) {
                throw new MissingRequiredPropertyException("GetEngineVersionArgs", "engine");
            }
            return $;
        }
    }

}
