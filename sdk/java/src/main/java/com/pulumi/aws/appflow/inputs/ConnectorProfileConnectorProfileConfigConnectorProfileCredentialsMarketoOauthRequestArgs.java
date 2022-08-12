// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.appflow.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoOauthRequestArgs extends com.pulumi.resources.ResourceArgs {

    public static final ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoOauthRequestArgs Empty = new ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoOauthRequestArgs();

    /**
     * The code provided by the connector when it has been authenticated via the connected app.
     * 
     */
    @Import(name="authCode")
    private @Nullable Output<String> authCode;

    /**
     * @return The code provided by the connector when it has been authenticated via the connected app.
     * 
     */
    public Optional<Output<String>> authCode() {
        return Optional.ofNullable(this.authCode);
    }

    /**
     * The URL to which the authentication server redirects the browser after authorization has been granted.
     * 
     */
    @Import(name="redirectUri")
    private @Nullable Output<String> redirectUri;

    /**
     * @return The URL to which the authentication server redirects the browser after authorization has been granted.
     * 
     */
    public Optional<Output<String>> redirectUri() {
        return Optional.ofNullable(this.redirectUri);
    }

    private ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoOauthRequestArgs() {}

    private ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoOauthRequestArgs(ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoOauthRequestArgs $) {
        this.authCode = $.authCode;
        this.redirectUri = $.redirectUri;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoOauthRequestArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoOauthRequestArgs $;

        public Builder() {
            $ = new ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoOauthRequestArgs();
        }

        public Builder(ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoOauthRequestArgs defaults) {
            $ = new ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoOauthRequestArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param authCode The code provided by the connector when it has been authenticated via the connected app.
         * 
         * @return builder
         * 
         */
        public Builder authCode(@Nullable Output<String> authCode) {
            $.authCode = authCode;
            return this;
        }

        /**
         * @param authCode The code provided by the connector when it has been authenticated via the connected app.
         * 
         * @return builder
         * 
         */
        public Builder authCode(String authCode) {
            return authCode(Output.of(authCode));
        }

        /**
         * @param redirectUri The URL to which the authentication server redirects the browser after authorization has been granted.
         * 
         * @return builder
         * 
         */
        public Builder redirectUri(@Nullable Output<String> redirectUri) {
            $.redirectUri = redirectUri;
            return this;
        }

        /**
         * @param redirectUri The URL to which the authentication server redirects the browser after authorization has been granted.
         * 
         * @return builder
         * 
         */
        public Builder redirectUri(String redirectUri) {
            return redirectUri(Output.of(redirectUri));
        }

        public ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoOauthRequestArgs build() {
            return $;
        }
    }

}