// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.pipes.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PipeTargetParametersRedshiftDataParametersArgs extends com.pulumi.resources.ResourceArgs {

    public static final PipeTargetParametersRedshiftDataParametersArgs Empty = new PipeTargetParametersRedshiftDataParametersArgs();

    /**
     * The name of the database. Required when authenticating using temporary credentials.
     * 
     */
    @Import(name="database", required=true)
    private Output<String> database;

    /**
     * @return The name of the database. Required when authenticating using temporary credentials.
     * 
     */
    public Output<String> database() {
        return this.database;
    }

    /**
     * The database user name. Required when authenticating using temporary credentials.
     * 
     */
    @Import(name="dbUser")
    private @Nullable Output<String> dbUser;

    /**
     * @return The database user name. Required when authenticating using temporary credentials.
     * 
     */
    public Optional<Output<String>> dbUser() {
        return Optional.ofNullable(this.dbUser);
    }

    /**
     * The name or ARN of the secret that enables access to the database. Required when authenticating using Secrets Manager.
     * 
     */
    @Import(name="secretManagerArn")
    private @Nullable Output<String> secretManagerArn;

    /**
     * @return The name or ARN of the secret that enables access to the database. Required when authenticating using Secrets Manager.
     * 
     */
    public Optional<Output<String>> secretManagerArn() {
        return Optional.ofNullable(this.secretManagerArn);
    }

    /**
     * List of SQL statements text to run, each of maximum length of 100,000.
     * 
     */
    @Import(name="sqls", required=true)
    private Output<List<String>> sqls;

    /**
     * @return List of SQL statements text to run, each of maximum length of 100,000.
     * 
     */
    public Output<List<String>> sqls() {
        return this.sqls;
    }

    /**
     * The name of the SQL statement. You can name the SQL statement when you create it to identify the query.
     * 
     */
    @Import(name="statementName")
    private @Nullable Output<String> statementName;

    /**
     * @return The name of the SQL statement. You can name the SQL statement when you create it to identify the query.
     * 
     */
    public Optional<Output<String>> statementName() {
        return Optional.ofNullable(this.statementName);
    }

    /**
     * Indicates whether to send an event back to EventBridge after the SQL statement runs.
     * 
     */
    @Import(name="withEvent")
    private @Nullable Output<Boolean> withEvent;

    /**
     * @return Indicates whether to send an event back to EventBridge after the SQL statement runs.
     * 
     */
    public Optional<Output<Boolean>> withEvent() {
        return Optional.ofNullable(this.withEvent);
    }

    private PipeTargetParametersRedshiftDataParametersArgs() {}

    private PipeTargetParametersRedshiftDataParametersArgs(PipeTargetParametersRedshiftDataParametersArgs $) {
        this.database = $.database;
        this.dbUser = $.dbUser;
        this.secretManagerArn = $.secretManagerArn;
        this.sqls = $.sqls;
        this.statementName = $.statementName;
        this.withEvent = $.withEvent;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PipeTargetParametersRedshiftDataParametersArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PipeTargetParametersRedshiftDataParametersArgs $;

        public Builder() {
            $ = new PipeTargetParametersRedshiftDataParametersArgs();
        }

        public Builder(PipeTargetParametersRedshiftDataParametersArgs defaults) {
            $ = new PipeTargetParametersRedshiftDataParametersArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param database The name of the database. Required when authenticating using temporary credentials.
         * 
         * @return builder
         * 
         */
        public Builder database(Output<String> database) {
            $.database = database;
            return this;
        }

        /**
         * @param database The name of the database. Required when authenticating using temporary credentials.
         * 
         * @return builder
         * 
         */
        public Builder database(String database) {
            return database(Output.of(database));
        }

        /**
         * @param dbUser The database user name. Required when authenticating using temporary credentials.
         * 
         * @return builder
         * 
         */
        public Builder dbUser(@Nullable Output<String> dbUser) {
            $.dbUser = dbUser;
            return this;
        }

        /**
         * @param dbUser The database user name. Required when authenticating using temporary credentials.
         * 
         * @return builder
         * 
         */
        public Builder dbUser(String dbUser) {
            return dbUser(Output.of(dbUser));
        }

        /**
         * @param secretManagerArn The name or ARN of the secret that enables access to the database. Required when authenticating using Secrets Manager.
         * 
         * @return builder
         * 
         */
        public Builder secretManagerArn(@Nullable Output<String> secretManagerArn) {
            $.secretManagerArn = secretManagerArn;
            return this;
        }

        /**
         * @param secretManagerArn The name or ARN of the secret that enables access to the database. Required when authenticating using Secrets Manager.
         * 
         * @return builder
         * 
         */
        public Builder secretManagerArn(String secretManagerArn) {
            return secretManagerArn(Output.of(secretManagerArn));
        }

        /**
         * @param sqls List of SQL statements text to run, each of maximum length of 100,000.
         * 
         * @return builder
         * 
         */
        public Builder sqls(Output<List<String>> sqls) {
            $.sqls = sqls;
            return this;
        }

        /**
         * @param sqls List of SQL statements text to run, each of maximum length of 100,000.
         * 
         * @return builder
         * 
         */
        public Builder sqls(List<String> sqls) {
            return sqls(Output.of(sqls));
        }

        /**
         * @param sqls List of SQL statements text to run, each of maximum length of 100,000.
         * 
         * @return builder
         * 
         */
        public Builder sqls(String... sqls) {
            return sqls(List.of(sqls));
        }

        /**
         * @param statementName The name of the SQL statement. You can name the SQL statement when you create it to identify the query.
         * 
         * @return builder
         * 
         */
        public Builder statementName(@Nullable Output<String> statementName) {
            $.statementName = statementName;
            return this;
        }

        /**
         * @param statementName The name of the SQL statement. You can name the SQL statement when you create it to identify the query.
         * 
         * @return builder
         * 
         */
        public Builder statementName(String statementName) {
            return statementName(Output.of(statementName));
        }

        /**
         * @param withEvent Indicates whether to send an event back to EventBridge after the SQL statement runs.
         * 
         * @return builder
         * 
         */
        public Builder withEvent(@Nullable Output<Boolean> withEvent) {
            $.withEvent = withEvent;
            return this;
        }

        /**
         * @param withEvent Indicates whether to send an event back to EventBridge after the SQL statement runs.
         * 
         * @return builder
         * 
         */
        public Builder withEvent(Boolean withEvent) {
            return withEvent(Output.of(withEvent));
        }

        public PipeTargetParametersRedshiftDataParametersArgs build() {
            $.database = Objects.requireNonNull($.database, "expected parameter 'database' to be non-null");
            $.sqls = Objects.requireNonNull($.sqls, "expected parameter 'sqls' to be non-null");
            return $;
        }
    }

}
