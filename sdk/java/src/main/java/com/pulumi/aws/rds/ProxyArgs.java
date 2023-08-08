// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.rds;

import com.pulumi.aws.rds.inputs.ProxyAuthArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProxyArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProxyArgs Empty = new ProxyArgs();

    /**
     * Configuration block(s) with authorization mechanisms to connect to the associated instances or clusters. Described below.
     * 
     */
    @Import(name="auths", required=true)
    private Output<List<ProxyAuthArgs>> auths;

    /**
     * @return Configuration block(s) with authorization mechanisms to connect to the associated instances or clusters. Described below.
     * 
     */
    public Output<List<ProxyAuthArgs>> auths() {
        return this.auths;
    }

    /**
     * Whether the proxy includes detailed information about SQL statements in its logs. This information helps you to debug issues involving SQL behavior or the performance and scalability of the proxy connections. The debug information includes the text of SQL statements that you submit through the proxy. Thus, only enable this setting when needed for debugging, and only when you have security measures in place to safeguard any sensitive information that appears in the logs.
     * 
     */
    @Import(name="debugLogging")
    private @Nullable Output<Boolean> debugLogging;

    /**
     * @return Whether the proxy includes detailed information about SQL statements in its logs. This information helps you to debug issues involving SQL behavior or the performance and scalability of the proxy connections. The debug information includes the text of SQL statements that you submit through the proxy. Thus, only enable this setting when needed for debugging, and only when you have security measures in place to safeguard any sensitive information that appears in the logs.
     * 
     */
    public Optional<Output<Boolean>> debugLogging() {
        return Optional.ofNullable(this.debugLogging);
    }

    /**
     * The kinds of databases that the proxy can connect to. This value determines which database network protocol the proxy recognizes when it interprets network traffic to and from the database. For Aurora MySQL, RDS for MariaDB, and RDS for MySQL databases, specify `MYSQL`. For Aurora PostgreSQL and RDS for PostgreSQL databases, specify `POSTGRESQL`. For RDS for Microsoft SQL Server, specify `SQLSERVER`. Valid values are `MYSQL`, `POSTGRESQL`, and `SQLSERVER`.
     * 
     */
    @Import(name="engineFamily", required=true)
    private Output<String> engineFamily;

    /**
     * @return The kinds of databases that the proxy can connect to. This value determines which database network protocol the proxy recognizes when it interprets network traffic to and from the database. For Aurora MySQL, RDS for MariaDB, and RDS for MySQL databases, specify `MYSQL`. For Aurora PostgreSQL and RDS for PostgreSQL databases, specify `POSTGRESQL`. For RDS for Microsoft SQL Server, specify `SQLSERVER`. Valid values are `MYSQL`, `POSTGRESQL`, and `SQLSERVER`.
     * 
     */
    public Output<String> engineFamily() {
        return this.engineFamily;
    }

    /**
     * The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it. You can set this value higher or lower than the connection timeout limit for the associated database.
     * 
     */
    @Import(name="idleClientTimeout")
    private @Nullable Output<Integer> idleClientTimeout;

    /**
     * @return The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it. You can set this value higher or lower than the connection timeout limit for the associated database.
     * 
     */
    public Optional<Output<Integer>> idleClientTimeout() {
        return Optional.ofNullable(this.idleClientTimeout);
    }

    /**
     * The identifier for the proxy. This name must be unique for all proxies owned by your AWS account in the specified AWS Region. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens; it can&#39;t end with a hyphen or contain two consecutive hyphens.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The identifier for the proxy. This name must be unique for all proxies owned by your AWS account in the specified AWS Region. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens; it can&#39;t end with a hyphen or contain two consecutive hyphens.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * A Boolean parameter that specifies whether Transport Layer Security (TLS) encryption is required for connections to the proxy. By enabling this setting, you can enforce encrypted TLS connections to the proxy.
     * 
     */
    @Import(name="requireTls")
    private @Nullable Output<Boolean> requireTls;

    /**
     * @return A Boolean parameter that specifies whether Transport Layer Security (TLS) encryption is required for connections to the proxy. By enabling this setting, you can enforce encrypted TLS connections to the proxy.
     * 
     */
    public Optional<Output<Boolean>> requireTls() {
        return Optional.ofNullable(this.requireTls);
    }

    /**
     * The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access secrets in AWS Secrets Manager.
     * 
     */
    @Import(name="roleArn", required=true)
    private Output<String> roleArn;

    /**
     * @return The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access secrets in AWS Secrets Manager.
     * 
     */
    public Output<String> roleArn() {
        return this.roleArn;
    }

    /**
     * A mapping of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A mapping of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * One or more VPC security group IDs to associate with the new proxy.
     * 
     */
    @Import(name="vpcSecurityGroupIds")
    private @Nullable Output<List<String>> vpcSecurityGroupIds;

    /**
     * @return One or more VPC security group IDs to associate with the new proxy.
     * 
     */
    public Optional<Output<List<String>>> vpcSecurityGroupIds() {
        return Optional.ofNullable(this.vpcSecurityGroupIds);
    }

    /**
     * One or more VPC subnet IDs to associate with the new proxy.
     * 
     */
    @Import(name="vpcSubnetIds", required=true)
    private Output<List<String>> vpcSubnetIds;

    /**
     * @return One or more VPC subnet IDs to associate with the new proxy.
     * 
     */
    public Output<List<String>> vpcSubnetIds() {
        return this.vpcSubnetIds;
    }

    private ProxyArgs() {}

    private ProxyArgs(ProxyArgs $) {
        this.auths = $.auths;
        this.debugLogging = $.debugLogging;
        this.engineFamily = $.engineFamily;
        this.idleClientTimeout = $.idleClientTimeout;
        this.name = $.name;
        this.requireTls = $.requireTls;
        this.roleArn = $.roleArn;
        this.tags = $.tags;
        this.vpcSecurityGroupIds = $.vpcSecurityGroupIds;
        this.vpcSubnetIds = $.vpcSubnetIds;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProxyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProxyArgs $;

        public Builder() {
            $ = new ProxyArgs();
        }

        public Builder(ProxyArgs defaults) {
            $ = new ProxyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param auths Configuration block(s) with authorization mechanisms to connect to the associated instances or clusters. Described below.
         * 
         * @return builder
         * 
         */
        public Builder auths(Output<List<ProxyAuthArgs>> auths) {
            $.auths = auths;
            return this;
        }

        /**
         * @param auths Configuration block(s) with authorization mechanisms to connect to the associated instances or clusters. Described below.
         * 
         * @return builder
         * 
         */
        public Builder auths(List<ProxyAuthArgs> auths) {
            return auths(Output.of(auths));
        }

        /**
         * @param auths Configuration block(s) with authorization mechanisms to connect to the associated instances or clusters. Described below.
         * 
         * @return builder
         * 
         */
        public Builder auths(ProxyAuthArgs... auths) {
            return auths(List.of(auths));
        }

        /**
         * @param debugLogging Whether the proxy includes detailed information about SQL statements in its logs. This information helps you to debug issues involving SQL behavior or the performance and scalability of the proxy connections. The debug information includes the text of SQL statements that you submit through the proxy. Thus, only enable this setting when needed for debugging, and only when you have security measures in place to safeguard any sensitive information that appears in the logs.
         * 
         * @return builder
         * 
         */
        public Builder debugLogging(@Nullable Output<Boolean> debugLogging) {
            $.debugLogging = debugLogging;
            return this;
        }

        /**
         * @param debugLogging Whether the proxy includes detailed information about SQL statements in its logs. This information helps you to debug issues involving SQL behavior or the performance and scalability of the proxy connections. The debug information includes the text of SQL statements that you submit through the proxy. Thus, only enable this setting when needed for debugging, and only when you have security measures in place to safeguard any sensitive information that appears in the logs.
         * 
         * @return builder
         * 
         */
        public Builder debugLogging(Boolean debugLogging) {
            return debugLogging(Output.of(debugLogging));
        }

        /**
         * @param engineFamily The kinds of databases that the proxy can connect to. This value determines which database network protocol the proxy recognizes when it interprets network traffic to and from the database. For Aurora MySQL, RDS for MariaDB, and RDS for MySQL databases, specify `MYSQL`. For Aurora PostgreSQL and RDS for PostgreSQL databases, specify `POSTGRESQL`. For RDS for Microsoft SQL Server, specify `SQLSERVER`. Valid values are `MYSQL`, `POSTGRESQL`, and `SQLSERVER`.
         * 
         * @return builder
         * 
         */
        public Builder engineFamily(Output<String> engineFamily) {
            $.engineFamily = engineFamily;
            return this;
        }

        /**
         * @param engineFamily The kinds of databases that the proxy can connect to. This value determines which database network protocol the proxy recognizes when it interprets network traffic to and from the database. For Aurora MySQL, RDS for MariaDB, and RDS for MySQL databases, specify `MYSQL`. For Aurora PostgreSQL and RDS for PostgreSQL databases, specify `POSTGRESQL`. For RDS for Microsoft SQL Server, specify `SQLSERVER`. Valid values are `MYSQL`, `POSTGRESQL`, and `SQLSERVER`.
         * 
         * @return builder
         * 
         */
        public Builder engineFamily(String engineFamily) {
            return engineFamily(Output.of(engineFamily));
        }

        /**
         * @param idleClientTimeout The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it. You can set this value higher or lower than the connection timeout limit for the associated database.
         * 
         * @return builder
         * 
         */
        public Builder idleClientTimeout(@Nullable Output<Integer> idleClientTimeout) {
            $.idleClientTimeout = idleClientTimeout;
            return this;
        }

        /**
         * @param idleClientTimeout The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it. You can set this value higher or lower than the connection timeout limit for the associated database.
         * 
         * @return builder
         * 
         */
        public Builder idleClientTimeout(Integer idleClientTimeout) {
            return idleClientTimeout(Output.of(idleClientTimeout));
        }

        /**
         * @param name The identifier for the proxy. This name must be unique for all proxies owned by your AWS account in the specified AWS Region. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens; it can&#39;t end with a hyphen or contain two consecutive hyphens.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The identifier for the proxy. This name must be unique for all proxies owned by your AWS account in the specified AWS Region. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens; it can&#39;t end with a hyphen or contain two consecutive hyphens.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param requireTls A Boolean parameter that specifies whether Transport Layer Security (TLS) encryption is required for connections to the proxy. By enabling this setting, you can enforce encrypted TLS connections to the proxy.
         * 
         * @return builder
         * 
         */
        public Builder requireTls(@Nullable Output<Boolean> requireTls) {
            $.requireTls = requireTls;
            return this;
        }

        /**
         * @param requireTls A Boolean parameter that specifies whether Transport Layer Security (TLS) encryption is required for connections to the proxy. By enabling this setting, you can enforce encrypted TLS connections to the proxy.
         * 
         * @return builder
         * 
         */
        public Builder requireTls(Boolean requireTls) {
            return requireTls(Output.of(requireTls));
        }

        /**
         * @param roleArn The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access secrets in AWS Secrets Manager.
         * 
         * @return builder
         * 
         */
        public Builder roleArn(Output<String> roleArn) {
            $.roleArn = roleArn;
            return this;
        }

        /**
         * @param roleArn The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access secrets in AWS Secrets Manager.
         * 
         * @return builder
         * 
         */
        public Builder roleArn(String roleArn) {
            return roleArn(Output.of(roleArn));
        }

        /**
         * @param tags A mapping of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A mapping of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param vpcSecurityGroupIds One or more VPC security group IDs to associate with the new proxy.
         * 
         * @return builder
         * 
         */
        public Builder vpcSecurityGroupIds(@Nullable Output<List<String>> vpcSecurityGroupIds) {
            $.vpcSecurityGroupIds = vpcSecurityGroupIds;
            return this;
        }

        /**
         * @param vpcSecurityGroupIds One or more VPC security group IDs to associate with the new proxy.
         * 
         * @return builder
         * 
         */
        public Builder vpcSecurityGroupIds(List<String> vpcSecurityGroupIds) {
            return vpcSecurityGroupIds(Output.of(vpcSecurityGroupIds));
        }

        /**
         * @param vpcSecurityGroupIds One or more VPC security group IDs to associate with the new proxy.
         * 
         * @return builder
         * 
         */
        public Builder vpcSecurityGroupIds(String... vpcSecurityGroupIds) {
            return vpcSecurityGroupIds(List.of(vpcSecurityGroupIds));
        }

        /**
         * @param vpcSubnetIds One or more VPC subnet IDs to associate with the new proxy.
         * 
         * @return builder
         * 
         */
        public Builder vpcSubnetIds(Output<List<String>> vpcSubnetIds) {
            $.vpcSubnetIds = vpcSubnetIds;
            return this;
        }

        /**
         * @param vpcSubnetIds One or more VPC subnet IDs to associate with the new proxy.
         * 
         * @return builder
         * 
         */
        public Builder vpcSubnetIds(List<String> vpcSubnetIds) {
            return vpcSubnetIds(Output.of(vpcSubnetIds));
        }

        /**
         * @param vpcSubnetIds One or more VPC subnet IDs to associate with the new proxy.
         * 
         * @return builder
         * 
         */
        public Builder vpcSubnetIds(String... vpcSubnetIds) {
            return vpcSubnetIds(List.of(vpcSubnetIds));
        }

        public ProxyArgs build() {
            $.auths = Objects.requireNonNull($.auths, "expected parameter 'auths' to be non-null");
            $.engineFamily = Objects.requireNonNull($.engineFamily, "expected parameter 'engineFamily' to be non-null");
            $.roleArn = Objects.requireNonNull($.roleArn, "expected parameter 'roleArn' to be non-null");
            $.vpcSubnetIds = Objects.requireNonNull($.vpcSubnetIds, "expected parameter 'vpcSubnetIds' to be non-null");
            return $;
        }
    }

}
