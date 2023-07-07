// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.cognito;

import com.pulumi.aws.cognito.inputs.ManagedUserPoolClientAnalyticsConfigurationArgs;
import com.pulumi.aws.cognito.inputs.ManagedUserPoolClientTokenValidityUnitsArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ManagedUserPoolClientArgs extends com.pulumi.resources.ResourceArgs {

    public static final ManagedUserPoolClientArgs Empty = new ManagedUserPoolClientArgs();

    /**
     * Time limit, between 5 minutes and 1 day, after which the access token is no longer valid and cannot be used.
     * By default, the unit is hours.
     * The unit can be overridden by a value in `token_validity_units.access_token`.
     * 
     */
    @Import(name="accessTokenValidity")
    private @Nullable Output<Integer> accessTokenValidity;

    /**
     * @return Time limit, between 5 minutes and 1 day, after which the access token is no longer valid and cannot be used.
     * By default, the unit is hours.
     * The unit can be overridden by a value in `token_validity_units.access_token`.
     * 
     */
    public Optional<Output<Integer>> accessTokenValidity() {
        return Optional.ofNullable(this.accessTokenValidity);
    }

    /**
     * List of allowed OAuth flows (code, implicit, client_credentials).
     * 
     */
    @Import(name="allowedOauthFlows")
    private @Nullable Output<List<String>> allowedOauthFlows;

    /**
     * @return List of allowed OAuth flows (code, implicit, client_credentials).
     * 
     */
    public Optional<Output<List<String>>> allowedOauthFlows() {
        return Optional.ofNullable(this.allowedOauthFlows);
    }

    /**
     * Whether the client is allowed to follow the OAuth protocol when interacting with Cognito user pools.
     * 
     */
    @Import(name="allowedOauthFlowsUserPoolClient")
    private @Nullable Output<Boolean> allowedOauthFlowsUserPoolClient;

    /**
     * @return Whether the client is allowed to follow the OAuth protocol when interacting with Cognito user pools.
     * 
     */
    public Optional<Output<Boolean>> allowedOauthFlowsUserPoolClient() {
        return Optional.ofNullable(this.allowedOauthFlowsUserPoolClient);
    }

    /**
     * List of allowed OAuth scopes (phone, email, openid, profile, and aws.cognito.signin.user.admin).
     * 
     */
    @Import(name="allowedOauthScopes")
    private @Nullable Output<List<String>> allowedOauthScopes;

    /**
     * @return List of allowed OAuth scopes (phone, email, openid, profile, and aws.cognito.signin.user.admin).
     * 
     */
    public Optional<Output<List<String>>> allowedOauthScopes() {
        return Optional.ofNullable(this.allowedOauthScopes);
    }

    /**
     * Configuration block for Amazon Pinpoint analytics for collecting metrics for this user pool. Detailed below.
     * 
     */
    @Import(name="analyticsConfiguration")
    private @Nullable Output<ManagedUserPoolClientAnalyticsConfigurationArgs> analyticsConfiguration;

    /**
     * @return Configuration block for Amazon Pinpoint analytics for collecting metrics for this user pool. Detailed below.
     * 
     */
    public Optional<Output<ManagedUserPoolClientAnalyticsConfigurationArgs>> analyticsConfiguration() {
        return Optional.ofNullable(this.analyticsConfiguration);
    }

    /**
     * Amazon Cognito creates a session token for each API request in an authentication flow. AuthSessionValidity is the duration, in minutes, of that session token. Your user pool native user must respond to each authentication challenge before the session expires. Valid values between `3` and `15`. Default value is `3`.
     * 
     */
    @Import(name="authSessionValidity")
    private @Nullable Output<Integer> authSessionValidity;

    /**
     * @return Amazon Cognito creates a session token for each API request in an authentication flow. AuthSessionValidity is the duration, in minutes, of that session token. Your user pool native user must respond to each authentication challenge before the session expires. Valid values between `3` and `15`. Default value is `3`.
     * 
     */
    public Optional<Output<Integer>> authSessionValidity() {
        return Optional.ofNullable(this.authSessionValidity);
    }

    /**
     * List of allowed callback URLs for the identity providers.
     * 
     */
    @Import(name="callbackUrls")
    private @Nullable Output<List<String>> callbackUrls;

    /**
     * @return List of allowed callback URLs for the identity providers.
     * 
     */
    public Optional<Output<List<String>>> callbackUrls() {
        return Optional.ofNullable(this.callbackUrls);
    }

    /**
     * Default redirect URI. Must be in the list of callback URLs.
     * 
     */
    @Import(name="defaultRedirectUri")
    private @Nullable Output<String> defaultRedirectUri;

    /**
     * @return Default redirect URI. Must be in the list of callback URLs.
     * 
     */
    public Optional<Output<String>> defaultRedirectUri() {
        return Optional.ofNullable(this.defaultRedirectUri);
    }

    /**
     * Activates the propagation of additional user context data.
     * 
     */
    @Import(name="enablePropagateAdditionalUserContextData")
    private @Nullable Output<Boolean> enablePropagateAdditionalUserContextData;

    /**
     * @return Activates the propagation of additional user context data.
     * 
     */
    public Optional<Output<Boolean>> enablePropagateAdditionalUserContextData() {
        return Optional.ofNullable(this.enablePropagateAdditionalUserContextData);
    }

    /**
     * Enables or disables token revocation.
     * 
     */
    @Import(name="enableTokenRevocation")
    private @Nullable Output<Boolean> enableTokenRevocation;

    /**
     * @return Enables or disables token revocation.
     * 
     */
    public Optional<Output<Boolean>> enableTokenRevocation() {
        return Optional.ofNullable(this.enableTokenRevocation);
    }

    /**
     * List of authentication flows (ADMIN_NO_SRP_AUTH, CUSTOM_AUTH_FLOW_ONLY, USER_PASSWORD_AUTH, ALLOW_ADMIN_USER_PASSWORD_AUTH, ALLOW_CUSTOM_AUTH, ALLOW_USER_PASSWORD_AUTH, ALLOW_USER_SRP_AUTH, ALLOW_REFRESH_TOKEN_AUTH).
     * 
     */
    @Import(name="explicitAuthFlows")
    private @Nullable Output<List<String>> explicitAuthFlows;

    /**
     * @return List of authentication flows (ADMIN_NO_SRP_AUTH, CUSTOM_AUTH_FLOW_ONLY, USER_PASSWORD_AUTH, ALLOW_ADMIN_USER_PASSWORD_AUTH, ALLOW_CUSTOM_AUTH, ALLOW_USER_PASSWORD_AUTH, ALLOW_USER_SRP_AUTH, ALLOW_REFRESH_TOKEN_AUTH).
     * 
     */
    public Optional<Output<List<String>>> explicitAuthFlows() {
        return Optional.ofNullable(this.explicitAuthFlows);
    }

    /**
     * Time limit, between 5 minutes and 1 day, after which the ID token is no longer valid and cannot be used.
     * By default, the unit is hours.
     * The unit can be overridden by a value in `token_validity_units.id_token`.
     * 
     */
    @Import(name="idTokenValidity")
    private @Nullable Output<Integer> idTokenValidity;

    /**
     * @return Time limit, between 5 minutes and 1 day, after which the ID token is no longer valid and cannot be used.
     * By default, the unit is hours.
     * The unit can be overridden by a value in `token_validity_units.id_token`.
     * 
     */
    public Optional<Output<Integer>> idTokenValidity() {
        return Optional.ofNullable(this.idTokenValidity);
    }

    /**
     * List of allowed logout URLs for the identity providers.
     * 
     */
    @Import(name="logoutUrls")
    private @Nullable Output<List<String>> logoutUrls;

    /**
     * @return List of allowed logout URLs for the identity providers.
     * 
     */
    public Optional<Output<List<String>>> logoutUrls() {
        return Optional.ofNullable(this.logoutUrls);
    }

    /**
     * Regular expression that matches the name of the desired User Pool Client.
     * Must match only one User Pool Client.
     * 
     */
    @Import(name="namePattern")
    private @Nullable Output<String> namePattern;

    /**
     * @return Regular expression that matches the name of the desired User Pool Client.
     * Must match only one User Pool Client.
     * 
     */
    public Optional<Output<String>> namePattern() {
        return Optional.ofNullable(this.namePattern);
    }

    /**
     * String that matches the beginning of the name of the desired User Pool Client.
     * Must match only one User Pool Client.
     * 
     * The following arguments are optional:
     * 
     */
    @Import(name="namePrefix")
    private @Nullable Output<String> namePrefix;

    /**
     * @return String that matches the beginning of the name of the desired User Pool Client.
     * Must match only one User Pool Client.
     * 
     * The following arguments are optional:
     * 
     */
    public Optional<Output<String>> namePrefix() {
        return Optional.ofNullable(this.namePrefix);
    }

    /**
     * Choose which errors and responses are returned by Cognito APIs during authentication, account confirmation, and password recovery when the user does not exist in the user pool. When set to `ENABLED` and the user does not exist, authentication returns an error indicating either the username or password was incorrect, and account confirmation and password recovery return a response indicating a code was sent to a simulated destination. When set to `LEGACY`, those APIs will return a `UserNotFoundException` exception if the user does not exist in the user pool.
     * 
     */
    @Import(name="preventUserExistenceErrors")
    private @Nullable Output<String> preventUserExistenceErrors;

    /**
     * @return Choose which errors and responses are returned by Cognito APIs during authentication, account confirmation, and password recovery when the user does not exist in the user pool. When set to `ENABLED` and the user does not exist, authentication returns an error indicating either the username or password was incorrect, and account confirmation and password recovery return a response indicating a code was sent to a simulated destination. When set to `LEGACY`, those APIs will return a `UserNotFoundException` exception if the user does not exist in the user pool.
     * 
     */
    public Optional<Output<String>> preventUserExistenceErrors() {
        return Optional.ofNullable(this.preventUserExistenceErrors);
    }

    /**
     * List of user pool attributes the application client can read from.
     * 
     */
    @Import(name="readAttributes")
    private @Nullable Output<List<String>> readAttributes;

    /**
     * @return List of user pool attributes the application client can read from.
     * 
     */
    public Optional<Output<List<String>>> readAttributes() {
        return Optional.ofNullable(this.readAttributes);
    }

    /**
     * Time limit, between 60 minutes and 10 years, after which the refresh token is no longer valid and cannot be used.
     * By default, the unit is days.
     * The unit can be overridden by a value in `token_validity_units.refresh_token`.
     * 
     */
    @Import(name="refreshTokenValidity")
    private @Nullable Output<Integer> refreshTokenValidity;

    /**
     * @return Time limit, between 60 minutes and 10 years, after which the refresh token is no longer valid and cannot be used.
     * By default, the unit is days.
     * The unit can be overridden by a value in `token_validity_units.refresh_token`.
     * 
     */
    public Optional<Output<Integer>> refreshTokenValidity() {
        return Optional.ofNullable(this.refreshTokenValidity);
    }

    /**
     * List of provider names for the identity providers that are supported on this client. Uses the `provider_name` attribute of `aws.cognito.IdentityProvider` resource(s), or the equivalent string(s).
     * 
     */
    @Import(name="supportedIdentityProviders")
    private @Nullable Output<List<String>> supportedIdentityProviders;

    /**
     * @return List of provider names for the identity providers that are supported on this client. Uses the `provider_name` attribute of `aws.cognito.IdentityProvider` resource(s), or the equivalent string(s).
     * 
     */
    public Optional<Output<List<String>>> supportedIdentityProviders() {
        return Optional.ofNullable(this.supportedIdentityProviders);
    }

    /**
     * Configuration block for units in which the validity times are represented in. Detailed below.
     * 
     */
    @Import(name="tokenValidityUnits")
    private @Nullable Output<ManagedUserPoolClientTokenValidityUnitsArgs> tokenValidityUnits;

    /**
     * @return Configuration block for units in which the validity times are represented in. Detailed below.
     * 
     */
    public Optional<Output<ManagedUserPoolClientTokenValidityUnitsArgs>> tokenValidityUnits() {
        return Optional.ofNullable(this.tokenValidityUnits);
    }

    /**
     * User pool the client belongs to.
     * 
     */
    @Import(name="userPoolId", required=true)
    private Output<String> userPoolId;

    /**
     * @return User pool the client belongs to.
     * 
     */
    public Output<String> userPoolId() {
        return this.userPoolId;
    }

    /**
     * List of user pool attributes the application client can write to.
     * 
     */
    @Import(name="writeAttributes")
    private @Nullable Output<List<String>> writeAttributes;

    /**
     * @return List of user pool attributes the application client can write to.
     * 
     */
    public Optional<Output<List<String>>> writeAttributes() {
        return Optional.ofNullable(this.writeAttributes);
    }

    private ManagedUserPoolClientArgs() {}

    private ManagedUserPoolClientArgs(ManagedUserPoolClientArgs $) {
        this.accessTokenValidity = $.accessTokenValidity;
        this.allowedOauthFlows = $.allowedOauthFlows;
        this.allowedOauthFlowsUserPoolClient = $.allowedOauthFlowsUserPoolClient;
        this.allowedOauthScopes = $.allowedOauthScopes;
        this.analyticsConfiguration = $.analyticsConfiguration;
        this.authSessionValidity = $.authSessionValidity;
        this.callbackUrls = $.callbackUrls;
        this.defaultRedirectUri = $.defaultRedirectUri;
        this.enablePropagateAdditionalUserContextData = $.enablePropagateAdditionalUserContextData;
        this.enableTokenRevocation = $.enableTokenRevocation;
        this.explicitAuthFlows = $.explicitAuthFlows;
        this.idTokenValidity = $.idTokenValidity;
        this.logoutUrls = $.logoutUrls;
        this.namePattern = $.namePattern;
        this.namePrefix = $.namePrefix;
        this.preventUserExistenceErrors = $.preventUserExistenceErrors;
        this.readAttributes = $.readAttributes;
        this.refreshTokenValidity = $.refreshTokenValidity;
        this.supportedIdentityProviders = $.supportedIdentityProviders;
        this.tokenValidityUnits = $.tokenValidityUnits;
        this.userPoolId = $.userPoolId;
        this.writeAttributes = $.writeAttributes;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ManagedUserPoolClientArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ManagedUserPoolClientArgs $;

        public Builder() {
            $ = new ManagedUserPoolClientArgs();
        }

        public Builder(ManagedUserPoolClientArgs defaults) {
            $ = new ManagedUserPoolClientArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param accessTokenValidity Time limit, between 5 minutes and 1 day, after which the access token is no longer valid and cannot be used.
         * By default, the unit is hours.
         * The unit can be overridden by a value in `token_validity_units.access_token`.
         * 
         * @return builder
         * 
         */
        public Builder accessTokenValidity(@Nullable Output<Integer> accessTokenValidity) {
            $.accessTokenValidity = accessTokenValidity;
            return this;
        }

        /**
         * @param accessTokenValidity Time limit, between 5 minutes and 1 day, after which the access token is no longer valid and cannot be used.
         * By default, the unit is hours.
         * The unit can be overridden by a value in `token_validity_units.access_token`.
         * 
         * @return builder
         * 
         */
        public Builder accessTokenValidity(Integer accessTokenValidity) {
            return accessTokenValidity(Output.of(accessTokenValidity));
        }

        /**
         * @param allowedOauthFlows List of allowed OAuth flows (code, implicit, client_credentials).
         * 
         * @return builder
         * 
         */
        public Builder allowedOauthFlows(@Nullable Output<List<String>> allowedOauthFlows) {
            $.allowedOauthFlows = allowedOauthFlows;
            return this;
        }

        /**
         * @param allowedOauthFlows List of allowed OAuth flows (code, implicit, client_credentials).
         * 
         * @return builder
         * 
         */
        public Builder allowedOauthFlows(List<String> allowedOauthFlows) {
            return allowedOauthFlows(Output.of(allowedOauthFlows));
        }

        /**
         * @param allowedOauthFlows List of allowed OAuth flows (code, implicit, client_credentials).
         * 
         * @return builder
         * 
         */
        public Builder allowedOauthFlows(String... allowedOauthFlows) {
            return allowedOauthFlows(List.of(allowedOauthFlows));
        }

        /**
         * @param allowedOauthFlowsUserPoolClient Whether the client is allowed to follow the OAuth protocol when interacting with Cognito user pools.
         * 
         * @return builder
         * 
         */
        public Builder allowedOauthFlowsUserPoolClient(@Nullable Output<Boolean> allowedOauthFlowsUserPoolClient) {
            $.allowedOauthFlowsUserPoolClient = allowedOauthFlowsUserPoolClient;
            return this;
        }

        /**
         * @param allowedOauthFlowsUserPoolClient Whether the client is allowed to follow the OAuth protocol when interacting with Cognito user pools.
         * 
         * @return builder
         * 
         */
        public Builder allowedOauthFlowsUserPoolClient(Boolean allowedOauthFlowsUserPoolClient) {
            return allowedOauthFlowsUserPoolClient(Output.of(allowedOauthFlowsUserPoolClient));
        }

        /**
         * @param allowedOauthScopes List of allowed OAuth scopes (phone, email, openid, profile, and aws.cognito.signin.user.admin).
         * 
         * @return builder
         * 
         */
        public Builder allowedOauthScopes(@Nullable Output<List<String>> allowedOauthScopes) {
            $.allowedOauthScopes = allowedOauthScopes;
            return this;
        }

        /**
         * @param allowedOauthScopes List of allowed OAuth scopes (phone, email, openid, profile, and aws.cognito.signin.user.admin).
         * 
         * @return builder
         * 
         */
        public Builder allowedOauthScopes(List<String> allowedOauthScopes) {
            return allowedOauthScopes(Output.of(allowedOauthScopes));
        }

        /**
         * @param allowedOauthScopes List of allowed OAuth scopes (phone, email, openid, profile, and aws.cognito.signin.user.admin).
         * 
         * @return builder
         * 
         */
        public Builder allowedOauthScopes(String... allowedOauthScopes) {
            return allowedOauthScopes(List.of(allowedOauthScopes));
        }

        /**
         * @param analyticsConfiguration Configuration block for Amazon Pinpoint analytics for collecting metrics for this user pool. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder analyticsConfiguration(@Nullable Output<ManagedUserPoolClientAnalyticsConfigurationArgs> analyticsConfiguration) {
            $.analyticsConfiguration = analyticsConfiguration;
            return this;
        }

        /**
         * @param analyticsConfiguration Configuration block for Amazon Pinpoint analytics for collecting metrics for this user pool. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder analyticsConfiguration(ManagedUserPoolClientAnalyticsConfigurationArgs analyticsConfiguration) {
            return analyticsConfiguration(Output.of(analyticsConfiguration));
        }

        /**
         * @param authSessionValidity Amazon Cognito creates a session token for each API request in an authentication flow. AuthSessionValidity is the duration, in minutes, of that session token. Your user pool native user must respond to each authentication challenge before the session expires. Valid values between `3` and `15`. Default value is `3`.
         * 
         * @return builder
         * 
         */
        public Builder authSessionValidity(@Nullable Output<Integer> authSessionValidity) {
            $.authSessionValidity = authSessionValidity;
            return this;
        }

        /**
         * @param authSessionValidity Amazon Cognito creates a session token for each API request in an authentication flow. AuthSessionValidity is the duration, in minutes, of that session token. Your user pool native user must respond to each authentication challenge before the session expires. Valid values between `3` and `15`. Default value is `3`.
         * 
         * @return builder
         * 
         */
        public Builder authSessionValidity(Integer authSessionValidity) {
            return authSessionValidity(Output.of(authSessionValidity));
        }

        /**
         * @param callbackUrls List of allowed callback URLs for the identity providers.
         * 
         * @return builder
         * 
         */
        public Builder callbackUrls(@Nullable Output<List<String>> callbackUrls) {
            $.callbackUrls = callbackUrls;
            return this;
        }

        /**
         * @param callbackUrls List of allowed callback URLs for the identity providers.
         * 
         * @return builder
         * 
         */
        public Builder callbackUrls(List<String> callbackUrls) {
            return callbackUrls(Output.of(callbackUrls));
        }

        /**
         * @param callbackUrls List of allowed callback URLs for the identity providers.
         * 
         * @return builder
         * 
         */
        public Builder callbackUrls(String... callbackUrls) {
            return callbackUrls(List.of(callbackUrls));
        }

        /**
         * @param defaultRedirectUri Default redirect URI. Must be in the list of callback URLs.
         * 
         * @return builder
         * 
         */
        public Builder defaultRedirectUri(@Nullable Output<String> defaultRedirectUri) {
            $.defaultRedirectUri = defaultRedirectUri;
            return this;
        }

        /**
         * @param defaultRedirectUri Default redirect URI. Must be in the list of callback URLs.
         * 
         * @return builder
         * 
         */
        public Builder defaultRedirectUri(String defaultRedirectUri) {
            return defaultRedirectUri(Output.of(defaultRedirectUri));
        }

        /**
         * @param enablePropagateAdditionalUserContextData Activates the propagation of additional user context data.
         * 
         * @return builder
         * 
         */
        public Builder enablePropagateAdditionalUserContextData(@Nullable Output<Boolean> enablePropagateAdditionalUserContextData) {
            $.enablePropagateAdditionalUserContextData = enablePropagateAdditionalUserContextData;
            return this;
        }

        /**
         * @param enablePropagateAdditionalUserContextData Activates the propagation of additional user context data.
         * 
         * @return builder
         * 
         */
        public Builder enablePropagateAdditionalUserContextData(Boolean enablePropagateAdditionalUserContextData) {
            return enablePropagateAdditionalUserContextData(Output.of(enablePropagateAdditionalUserContextData));
        }

        /**
         * @param enableTokenRevocation Enables or disables token revocation.
         * 
         * @return builder
         * 
         */
        public Builder enableTokenRevocation(@Nullable Output<Boolean> enableTokenRevocation) {
            $.enableTokenRevocation = enableTokenRevocation;
            return this;
        }

        /**
         * @param enableTokenRevocation Enables or disables token revocation.
         * 
         * @return builder
         * 
         */
        public Builder enableTokenRevocation(Boolean enableTokenRevocation) {
            return enableTokenRevocation(Output.of(enableTokenRevocation));
        }

        /**
         * @param explicitAuthFlows List of authentication flows (ADMIN_NO_SRP_AUTH, CUSTOM_AUTH_FLOW_ONLY, USER_PASSWORD_AUTH, ALLOW_ADMIN_USER_PASSWORD_AUTH, ALLOW_CUSTOM_AUTH, ALLOW_USER_PASSWORD_AUTH, ALLOW_USER_SRP_AUTH, ALLOW_REFRESH_TOKEN_AUTH).
         * 
         * @return builder
         * 
         */
        public Builder explicitAuthFlows(@Nullable Output<List<String>> explicitAuthFlows) {
            $.explicitAuthFlows = explicitAuthFlows;
            return this;
        }

        /**
         * @param explicitAuthFlows List of authentication flows (ADMIN_NO_SRP_AUTH, CUSTOM_AUTH_FLOW_ONLY, USER_PASSWORD_AUTH, ALLOW_ADMIN_USER_PASSWORD_AUTH, ALLOW_CUSTOM_AUTH, ALLOW_USER_PASSWORD_AUTH, ALLOW_USER_SRP_AUTH, ALLOW_REFRESH_TOKEN_AUTH).
         * 
         * @return builder
         * 
         */
        public Builder explicitAuthFlows(List<String> explicitAuthFlows) {
            return explicitAuthFlows(Output.of(explicitAuthFlows));
        }

        /**
         * @param explicitAuthFlows List of authentication flows (ADMIN_NO_SRP_AUTH, CUSTOM_AUTH_FLOW_ONLY, USER_PASSWORD_AUTH, ALLOW_ADMIN_USER_PASSWORD_AUTH, ALLOW_CUSTOM_AUTH, ALLOW_USER_PASSWORD_AUTH, ALLOW_USER_SRP_AUTH, ALLOW_REFRESH_TOKEN_AUTH).
         * 
         * @return builder
         * 
         */
        public Builder explicitAuthFlows(String... explicitAuthFlows) {
            return explicitAuthFlows(List.of(explicitAuthFlows));
        }

        /**
         * @param idTokenValidity Time limit, between 5 minutes and 1 day, after which the ID token is no longer valid and cannot be used.
         * By default, the unit is hours.
         * The unit can be overridden by a value in `token_validity_units.id_token`.
         * 
         * @return builder
         * 
         */
        public Builder idTokenValidity(@Nullable Output<Integer> idTokenValidity) {
            $.idTokenValidity = idTokenValidity;
            return this;
        }

        /**
         * @param idTokenValidity Time limit, between 5 minutes and 1 day, after which the ID token is no longer valid and cannot be used.
         * By default, the unit is hours.
         * The unit can be overridden by a value in `token_validity_units.id_token`.
         * 
         * @return builder
         * 
         */
        public Builder idTokenValidity(Integer idTokenValidity) {
            return idTokenValidity(Output.of(idTokenValidity));
        }

        /**
         * @param logoutUrls List of allowed logout URLs for the identity providers.
         * 
         * @return builder
         * 
         */
        public Builder logoutUrls(@Nullable Output<List<String>> logoutUrls) {
            $.logoutUrls = logoutUrls;
            return this;
        }

        /**
         * @param logoutUrls List of allowed logout URLs for the identity providers.
         * 
         * @return builder
         * 
         */
        public Builder logoutUrls(List<String> logoutUrls) {
            return logoutUrls(Output.of(logoutUrls));
        }

        /**
         * @param logoutUrls List of allowed logout URLs for the identity providers.
         * 
         * @return builder
         * 
         */
        public Builder logoutUrls(String... logoutUrls) {
            return logoutUrls(List.of(logoutUrls));
        }

        /**
         * @param namePattern Regular expression that matches the name of the desired User Pool Client.
         * Must match only one User Pool Client.
         * 
         * @return builder
         * 
         */
        public Builder namePattern(@Nullable Output<String> namePattern) {
            $.namePattern = namePattern;
            return this;
        }

        /**
         * @param namePattern Regular expression that matches the name of the desired User Pool Client.
         * Must match only one User Pool Client.
         * 
         * @return builder
         * 
         */
        public Builder namePattern(String namePattern) {
            return namePattern(Output.of(namePattern));
        }

        /**
         * @param namePrefix String that matches the beginning of the name of the desired User Pool Client.
         * Must match only one User Pool Client.
         * 
         * The following arguments are optional:
         * 
         * @return builder
         * 
         */
        public Builder namePrefix(@Nullable Output<String> namePrefix) {
            $.namePrefix = namePrefix;
            return this;
        }

        /**
         * @param namePrefix String that matches the beginning of the name of the desired User Pool Client.
         * Must match only one User Pool Client.
         * 
         * The following arguments are optional:
         * 
         * @return builder
         * 
         */
        public Builder namePrefix(String namePrefix) {
            return namePrefix(Output.of(namePrefix));
        }

        /**
         * @param preventUserExistenceErrors Choose which errors and responses are returned by Cognito APIs during authentication, account confirmation, and password recovery when the user does not exist in the user pool. When set to `ENABLED` and the user does not exist, authentication returns an error indicating either the username or password was incorrect, and account confirmation and password recovery return a response indicating a code was sent to a simulated destination. When set to `LEGACY`, those APIs will return a `UserNotFoundException` exception if the user does not exist in the user pool.
         * 
         * @return builder
         * 
         */
        public Builder preventUserExistenceErrors(@Nullable Output<String> preventUserExistenceErrors) {
            $.preventUserExistenceErrors = preventUserExistenceErrors;
            return this;
        }

        /**
         * @param preventUserExistenceErrors Choose which errors and responses are returned by Cognito APIs during authentication, account confirmation, and password recovery when the user does not exist in the user pool. When set to `ENABLED` and the user does not exist, authentication returns an error indicating either the username or password was incorrect, and account confirmation and password recovery return a response indicating a code was sent to a simulated destination. When set to `LEGACY`, those APIs will return a `UserNotFoundException` exception if the user does not exist in the user pool.
         * 
         * @return builder
         * 
         */
        public Builder preventUserExistenceErrors(String preventUserExistenceErrors) {
            return preventUserExistenceErrors(Output.of(preventUserExistenceErrors));
        }

        /**
         * @param readAttributes List of user pool attributes the application client can read from.
         * 
         * @return builder
         * 
         */
        public Builder readAttributes(@Nullable Output<List<String>> readAttributes) {
            $.readAttributes = readAttributes;
            return this;
        }

        /**
         * @param readAttributes List of user pool attributes the application client can read from.
         * 
         * @return builder
         * 
         */
        public Builder readAttributes(List<String> readAttributes) {
            return readAttributes(Output.of(readAttributes));
        }

        /**
         * @param readAttributes List of user pool attributes the application client can read from.
         * 
         * @return builder
         * 
         */
        public Builder readAttributes(String... readAttributes) {
            return readAttributes(List.of(readAttributes));
        }

        /**
         * @param refreshTokenValidity Time limit, between 60 minutes and 10 years, after which the refresh token is no longer valid and cannot be used.
         * By default, the unit is days.
         * The unit can be overridden by a value in `token_validity_units.refresh_token`.
         * 
         * @return builder
         * 
         */
        public Builder refreshTokenValidity(@Nullable Output<Integer> refreshTokenValidity) {
            $.refreshTokenValidity = refreshTokenValidity;
            return this;
        }

        /**
         * @param refreshTokenValidity Time limit, between 60 minutes and 10 years, after which the refresh token is no longer valid and cannot be used.
         * By default, the unit is days.
         * The unit can be overridden by a value in `token_validity_units.refresh_token`.
         * 
         * @return builder
         * 
         */
        public Builder refreshTokenValidity(Integer refreshTokenValidity) {
            return refreshTokenValidity(Output.of(refreshTokenValidity));
        }

        /**
         * @param supportedIdentityProviders List of provider names for the identity providers that are supported on this client. Uses the `provider_name` attribute of `aws.cognito.IdentityProvider` resource(s), or the equivalent string(s).
         * 
         * @return builder
         * 
         */
        public Builder supportedIdentityProviders(@Nullable Output<List<String>> supportedIdentityProviders) {
            $.supportedIdentityProviders = supportedIdentityProviders;
            return this;
        }

        /**
         * @param supportedIdentityProviders List of provider names for the identity providers that are supported on this client. Uses the `provider_name` attribute of `aws.cognito.IdentityProvider` resource(s), or the equivalent string(s).
         * 
         * @return builder
         * 
         */
        public Builder supportedIdentityProviders(List<String> supportedIdentityProviders) {
            return supportedIdentityProviders(Output.of(supportedIdentityProviders));
        }

        /**
         * @param supportedIdentityProviders List of provider names for the identity providers that are supported on this client. Uses the `provider_name` attribute of `aws.cognito.IdentityProvider` resource(s), or the equivalent string(s).
         * 
         * @return builder
         * 
         */
        public Builder supportedIdentityProviders(String... supportedIdentityProviders) {
            return supportedIdentityProviders(List.of(supportedIdentityProviders));
        }

        /**
         * @param tokenValidityUnits Configuration block for units in which the validity times are represented in. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder tokenValidityUnits(@Nullable Output<ManagedUserPoolClientTokenValidityUnitsArgs> tokenValidityUnits) {
            $.tokenValidityUnits = tokenValidityUnits;
            return this;
        }

        /**
         * @param tokenValidityUnits Configuration block for units in which the validity times are represented in. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder tokenValidityUnits(ManagedUserPoolClientTokenValidityUnitsArgs tokenValidityUnits) {
            return tokenValidityUnits(Output.of(tokenValidityUnits));
        }

        /**
         * @param userPoolId User pool the client belongs to.
         * 
         * @return builder
         * 
         */
        public Builder userPoolId(Output<String> userPoolId) {
            $.userPoolId = userPoolId;
            return this;
        }

        /**
         * @param userPoolId User pool the client belongs to.
         * 
         * @return builder
         * 
         */
        public Builder userPoolId(String userPoolId) {
            return userPoolId(Output.of(userPoolId));
        }

        /**
         * @param writeAttributes List of user pool attributes the application client can write to.
         * 
         * @return builder
         * 
         */
        public Builder writeAttributes(@Nullable Output<List<String>> writeAttributes) {
            $.writeAttributes = writeAttributes;
            return this;
        }

        /**
         * @param writeAttributes List of user pool attributes the application client can write to.
         * 
         * @return builder
         * 
         */
        public Builder writeAttributes(List<String> writeAttributes) {
            return writeAttributes(Output.of(writeAttributes));
        }

        /**
         * @param writeAttributes List of user pool attributes the application client can write to.
         * 
         * @return builder
         * 
         */
        public Builder writeAttributes(String... writeAttributes) {
            return writeAttributes(List.of(writeAttributes));
        }

        public ManagedUserPoolClientArgs build() {
            $.userPoolId = Objects.requireNonNull($.userPoolId, "expected parameter 'userPoolId' to be non-null");
            return $;
        }
    }

}
