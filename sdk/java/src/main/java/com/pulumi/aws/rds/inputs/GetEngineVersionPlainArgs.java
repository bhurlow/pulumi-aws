// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.rds.inputs;

import com.pulumi.aws.rds.inputs.GetEngineVersionFilter;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetEngineVersionPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetEngineVersionPlainArgs Empty = new GetEngineVersionPlainArgs();

    /**
     * When set to `true`, the default version for the specified `engine` or combination of `engine` and major `version` will be returned. Can be used to limit responses to a single version when they would otherwise fail for returning multiple versions.
     * 
     */
    @Import(name="defaultOnly")
    private @Nullable Boolean defaultOnly;

    /**
     * @return When set to `true`, the default version for the specified `engine` or combination of `engine` and major `version` will be returned. Can be used to limit responses to a single version when they would otherwise fail for returning multiple versions.
     * 
     */
    public Optional<Boolean> defaultOnly() {
        return Optional.ofNullable(this.defaultOnly);
    }

    /**
     * DB engine. Engine values include `aurora`, `aurora-mysql`, `aurora-postgresql`, `docdb`, `mariadb`, `mysql`, `neptune`, `oracle-ee`, `oracle-se`, `oracle-se1`, `oracle-se2`, `postgres`, `sqlserver-ee`, `sqlserver-ex`, `sqlserver-se`, and `sqlserver-web`.
     * 
     */
    @Import(name="engine", required=true)
    private String engine;

    /**
     * @return DB engine. Engine values include `aurora`, `aurora-mysql`, `aurora-postgresql`, `docdb`, `mariadb`, `mysql`, `neptune`, `oracle-ee`, `oracle-se`, `oracle-se1`, `oracle-se2`, `postgres`, `sqlserver-ee`, `sqlserver-ex`, `sqlserver-se`, and `sqlserver-web`.
     * 
     */
    public String engine() {
        return this.engine;
    }

    /**
     * One or more name/value pairs to filter off of. There are several valid keys; for a full reference, check out [describe-db-engine-versions in the AWS CLI reference](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/rds/describe-db-engine-versions.html).
     * 
     */
    @Import(name="filters")
    private @Nullable List<GetEngineVersionFilter> filters;

    /**
     * @return One or more name/value pairs to filter off of. There are several valid keys; for a full reference, check out [describe-db-engine-versions in the AWS CLI reference](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/rds/describe-db-engine-versions.html).
     * 
     */
    public Optional<List<GetEngineVersionFilter>> filters() {
        return Optional.ofNullable(this.filters);
    }

    /**
     * When set to `true`, the specified `version` or member of `preferred_versions` will be returned even if it is `deprecated`. Otherwise, only `available` versions will be returned.
     * 
     */
    @Import(name="includeAll")
    private @Nullable Boolean includeAll;

    /**
     * @return When set to `true`, the specified `version` or member of `preferred_versions` will be returned even if it is `deprecated`. Otherwise, only `available` versions will be returned.
     * 
     */
    public Optional<Boolean> includeAll() {
        return Optional.ofNullable(this.includeAll);
    }

    /**
     * Name of a specific DB parameter group family. Examples of parameter group families are `mysql8.0`, `mariadb10.4`, and `postgres12`.
     * 
     */
    @Import(name="parameterGroupFamily")
    private @Nullable String parameterGroupFamily;

    /**
     * @return Name of a specific DB parameter group family. Examples of parameter group families are `mysql8.0`, `mariadb10.4`, and `postgres12`.
     * 
     */
    public Optional<String> parameterGroupFamily() {
        return Optional.ofNullable(this.parameterGroupFamily);
    }

    /**
     * Ordered list of preferred engine versions. The first match in this list will be returned. If no preferred matches are found and the original search returned more than one result, an error is returned. If both the `version` and `preferred_versions` arguments are not configured, the data source will return the default version for the engine.
     * 
     */
    @Import(name="preferredVersions")
    private @Nullable List<String> preferredVersions;

    /**
     * @return Ordered list of preferred engine versions. The first match in this list will be returned. If no preferred matches are found and the original search returned more than one result, an error is returned. If both the `version` and `preferred_versions` arguments are not configured, the data source will return the default version for the engine.
     * 
     */
    public Optional<List<String>> preferredVersions() {
        return Optional.ofNullable(this.preferredVersions);
    }

    /**
     * Version of the DB engine. For example, `5.7.22`, `10.1.34`, and `12.3`. If both the `version` and `preferred_versions` arguments are not configured, the data source will return the default version for the engine.
     * 
     */
    @Import(name="version")
    private @Nullable String version;

    /**
     * @return Version of the DB engine. For example, `5.7.22`, `10.1.34`, and `12.3`. If both the `version` and `preferred_versions` arguments are not configured, the data source will return the default version for the engine.
     * 
     */
    public Optional<String> version() {
        return Optional.ofNullable(this.version);
    }

    private GetEngineVersionPlainArgs() {}

    private GetEngineVersionPlainArgs(GetEngineVersionPlainArgs $) {
        this.defaultOnly = $.defaultOnly;
        this.engine = $.engine;
        this.filters = $.filters;
        this.includeAll = $.includeAll;
        this.parameterGroupFamily = $.parameterGroupFamily;
        this.preferredVersions = $.preferredVersions;
        this.version = $.version;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetEngineVersionPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetEngineVersionPlainArgs $;

        public Builder() {
            $ = new GetEngineVersionPlainArgs();
        }

        public Builder(GetEngineVersionPlainArgs defaults) {
            $ = new GetEngineVersionPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param defaultOnly When set to `true`, the default version for the specified `engine` or combination of `engine` and major `version` will be returned. Can be used to limit responses to a single version when they would otherwise fail for returning multiple versions.
         * 
         * @return builder
         * 
         */
        public Builder defaultOnly(@Nullable Boolean defaultOnly) {
            $.defaultOnly = defaultOnly;
            return this;
        }

        /**
         * @param engine DB engine. Engine values include `aurora`, `aurora-mysql`, `aurora-postgresql`, `docdb`, `mariadb`, `mysql`, `neptune`, `oracle-ee`, `oracle-se`, `oracle-se1`, `oracle-se2`, `postgres`, `sqlserver-ee`, `sqlserver-ex`, `sqlserver-se`, and `sqlserver-web`.
         * 
         * @return builder
         * 
         */
        public Builder engine(String engine) {
            $.engine = engine;
            return this;
        }

        /**
         * @param filters One or more name/value pairs to filter off of. There are several valid keys; for a full reference, check out [describe-db-engine-versions in the AWS CLI reference](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/rds/describe-db-engine-versions.html).
         * 
         * @return builder
         * 
         */
        public Builder filters(@Nullable List<GetEngineVersionFilter> filters) {
            $.filters = filters;
            return this;
        }

        /**
         * @param filters One or more name/value pairs to filter off of. There are several valid keys; for a full reference, check out [describe-db-engine-versions in the AWS CLI reference](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/rds/describe-db-engine-versions.html).
         * 
         * @return builder
         * 
         */
        public Builder filters(GetEngineVersionFilter... filters) {
            return filters(List.of(filters));
        }

        /**
         * @param includeAll When set to `true`, the specified `version` or member of `preferred_versions` will be returned even if it is `deprecated`. Otherwise, only `available` versions will be returned.
         * 
         * @return builder
         * 
         */
        public Builder includeAll(@Nullable Boolean includeAll) {
            $.includeAll = includeAll;
            return this;
        }

        /**
         * @param parameterGroupFamily Name of a specific DB parameter group family. Examples of parameter group families are `mysql8.0`, `mariadb10.4`, and `postgres12`.
         * 
         * @return builder
         * 
         */
        public Builder parameterGroupFamily(@Nullable String parameterGroupFamily) {
            $.parameterGroupFamily = parameterGroupFamily;
            return this;
        }

        /**
         * @param preferredVersions Ordered list of preferred engine versions. The first match in this list will be returned. If no preferred matches are found and the original search returned more than one result, an error is returned. If both the `version` and `preferred_versions` arguments are not configured, the data source will return the default version for the engine.
         * 
         * @return builder
         * 
         */
        public Builder preferredVersions(@Nullable List<String> preferredVersions) {
            $.preferredVersions = preferredVersions;
            return this;
        }

        /**
         * @param preferredVersions Ordered list of preferred engine versions. The first match in this list will be returned. If no preferred matches are found and the original search returned more than one result, an error is returned. If both the `version` and `preferred_versions` arguments are not configured, the data source will return the default version for the engine.
         * 
         * @return builder
         * 
         */
        public Builder preferredVersions(String... preferredVersions) {
            return preferredVersions(List.of(preferredVersions));
        }

        /**
         * @param version Version of the DB engine. For example, `5.7.22`, `10.1.34`, and `12.3`. If both the `version` and `preferred_versions` arguments are not configured, the data source will return the default version for the engine.
         * 
         * @return builder
         * 
         */
        public Builder version(@Nullable String version) {
            $.version = version;
            return this;
        }

        public GetEngineVersionPlainArgs build() {
            $.engine = Objects.requireNonNull($.engine, "expected parameter 'engine' to be non-null");
            return $;
        }
    }

}
