// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ec2clientvpn;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.ec2clientvpn.AuthorizationRuleArgs;
import com.pulumi.aws.ec2clientvpn.inputs.AuthorizationRuleState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides authorization rules for AWS Client VPN endpoints. For more information on usage, please see the
 * [AWS Client VPN Administrator&#39;s Guide](https://docs.aws.amazon.com/vpn/latest/clientvpn-admin/what-is.html).
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.ec2clientvpn.AuthorizationRule;
 * import com.pulumi.aws.ec2clientvpn.AuthorizationRuleArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var example = new AuthorizationRule(&#34;example&#34;, AuthorizationRuleArgs.builder()        
 *             .clientVpnEndpointId(aws_ec2_client_vpn_endpoint.example().id())
 *             .targetNetworkCidr(aws_subnet.example().cidr_block())
 *             .authorizeAllGroups(true)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * AWS Client VPN authorization rules can be imported using the endpoint ID and target network CIDR. If there is a specific group name that is included as well. All values are separated by a `,`.
 * 
 * ```sh
 *  $ pulumi import aws:ec2clientvpn/authorizationRule:AuthorizationRule example cvpn-endpoint-0ac3a1abbccddd666,10.1.0.0/24
 * ```
 * 
 * ```sh
 *  $ pulumi import aws:ec2clientvpn/authorizationRule:AuthorizationRule example cvpn-endpoint-0ac3a1abbccddd666,10.1.0.0/24,team-a
 * ```
 * 
 */
@ResourceType(type="aws:ec2clientvpn/authorizationRule:AuthorizationRule")
public class AuthorizationRule extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the group to which the authorization rule grants access. One of `access_group_id` or `authorize_all_groups` must be set.
     * 
     */
    @Export(name="accessGroupId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> accessGroupId;

    /**
     * @return The ID of the group to which the authorization rule grants access. One of `access_group_id` or `authorize_all_groups` must be set.
     * 
     */
    public Output<Optional<String>> accessGroupId() {
        return Codegen.optional(this.accessGroupId);
    }
    /**
     * Indicates whether the authorization rule grants access to all clients. One of `access_group_id` or `authorize_all_groups` must be set.
     * 
     */
    @Export(name="authorizeAllGroups", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> authorizeAllGroups;

    /**
     * @return Indicates whether the authorization rule grants access to all clients. One of `access_group_id` or `authorize_all_groups` must be set.
     * 
     */
    public Output<Optional<Boolean>> authorizeAllGroups() {
        return Codegen.optional(this.authorizeAllGroups);
    }
    /**
     * The ID of the Client VPN endpoint.
     * 
     */
    @Export(name="clientVpnEndpointId", refs={String.class}, tree="[0]")
    private Output<String> clientVpnEndpointId;

    /**
     * @return The ID of the Client VPN endpoint.
     * 
     */
    public Output<String> clientVpnEndpointId() {
        return this.clientVpnEndpointId;
    }
    /**
     * A brief description of the authorization rule.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A brief description of the authorization rule.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The IPv4 address range, in CIDR notation, of the network to which the authorization rule applies.
     * 
     */
    @Export(name="targetNetworkCidr", refs={String.class}, tree="[0]")
    private Output<String> targetNetworkCidr;

    /**
     * @return The IPv4 address range, in CIDR notation, of the network to which the authorization rule applies.
     * 
     */
    public Output<String> targetNetworkCidr() {
        return this.targetNetworkCidr;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AuthorizationRule(String name) {
        this(name, AuthorizationRuleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AuthorizationRule(String name, AuthorizationRuleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AuthorizationRule(String name, AuthorizationRuleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ec2clientvpn/authorizationRule:AuthorizationRule", name, args == null ? AuthorizationRuleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AuthorizationRule(String name, Output<String> id, @Nullable AuthorizationRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ec2clientvpn/authorizationRule:AuthorizationRule", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static AuthorizationRule get(String name, Output<String> id, @Nullable AuthorizationRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AuthorizationRule(name, id, state, options);
    }
}
