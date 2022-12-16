// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.elasticloadbalancing;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.elasticloadbalancing.LoadBalancerBackendServerPolicyArgs;
import com.pulumi.aws.elasticloadbalancing.inputs.LoadBalancerBackendServerPolicyState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Attaches a load balancer policy to an ELB backend server.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.elb.LoadBalancer;
 * import com.pulumi.aws.elb.LoadBalancerArgs;
 * import com.pulumi.aws.elb.inputs.LoadBalancerListenerArgs;
 * import com.pulumi.aws.elb.LoadBalancerPolicy;
 * import com.pulumi.aws.elb.LoadBalancerPolicyArgs;
 * import com.pulumi.aws.elb.inputs.LoadBalancerPolicyPolicyAttributeArgs;
 * import com.pulumi.aws.elb.LoadBalancerBackendServerPolicy;
 * import com.pulumi.aws.elb.LoadBalancerBackendServerPolicyArgs;
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
 *         var wu_tang = new LoadBalancer(&#34;wu-tang&#34;, LoadBalancerArgs.builder()        
 *             .availabilityZones(&#34;us-east-1a&#34;)
 *             .listeners(LoadBalancerListenerArgs.builder()
 *                 .instancePort(443)
 *                 .instanceProtocol(&#34;http&#34;)
 *                 .lbPort(443)
 *                 .lbProtocol(&#34;https&#34;)
 *                 .sslCertificateId(&#34;arn:aws:iam::000000000000:server-certificate/wu-tang.net&#34;)
 *                 .build())
 *             .tags(Map.of(&#34;Name&#34;, &#34;wu-tang&#34;))
 *             .build());
 * 
 *         var wu_tang_ca_pubkey_policy = new LoadBalancerPolicy(&#34;wu-tang-ca-pubkey-policy&#34;, LoadBalancerPolicyArgs.builder()        
 *             .loadBalancerName(wu_tang.name())
 *             .policyName(&#34;wu-tang-ca-pubkey-policy&#34;)
 *             .policyTypeName(&#34;PublicKeyPolicyType&#34;)
 *             .policyAttributes(LoadBalancerPolicyPolicyAttributeArgs.builder()
 *                 .name(&#34;PublicKey&#34;)
 *                 .value(Files.readString(Paths.get(&#34;wu-tang-pubkey&#34;)))
 *                 .build())
 *             .build());
 * 
 *         var wu_tang_root_ca_backend_auth_policy = new LoadBalancerPolicy(&#34;wu-tang-root-ca-backend-auth-policy&#34;, LoadBalancerPolicyArgs.builder()        
 *             .loadBalancerName(wu_tang.name())
 *             .policyName(&#34;wu-tang-root-ca-backend-auth-policy&#34;)
 *             .policyTypeName(&#34;BackendServerAuthenticationPolicyType&#34;)
 *             .policyAttributes(LoadBalancerPolicyPolicyAttributeArgs.builder()
 *                 .name(&#34;PublicKeyPolicyName&#34;)
 *                 .value(aws_load_balancer_policy.wu-tang-root-ca-pubkey-policy().policy_name())
 *                 .build())
 *             .build());
 * 
 *         var wu_tang_backend_auth_policies_443 = new LoadBalancerBackendServerPolicy(&#34;wu-tang-backend-auth-policies-443&#34;, LoadBalancerBackendServerPolicyArgs.builder()        
 *             .loadBalancerName(wu_tang.name())
 *             .instancePort(443)
 *             .policyNames(wu_tang_root_ca_backend_auth_policy.policyName())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * @deprecated
 * aws.elasticloadbalancing.LoadBalancerBackendServerPolicy has been deprecated in favor of aws.elb.LoadBalancerBackendServerPolicy
 * 
 */
@Deprecated /* aws.elasticloadbalancing.LoadBalancerBackendServerPolicy has been deprecated in favor of aws.elb.LoadBalancerBackendServerPolicy */
@ResourceType(type="aws:elasticloadbalancing/loadBalancerBackendServerPolicy:LoadBalancerBackendServerPolicy")
public class LoadBalancerBackendServerPolicy extends com.pulumi.resources.CustomResource {
    /**
     * The instance port to apply the policy to.
     * 
     */
    @Export(name="instancePort", refs={Integer.class}, tree="[0]")
    private Output<Integer> instancePort;

    /**
     * @return The instance port to apply the policy to.
     * 
     */
    public Output<Integer> instancePort() {
        return this.instancePort;
    }
    /**
     * The load balancer to attach the policy to.
     * 
     */
    @Export(name="loadBalancerName", refs={String.class}, tree="[0]")
    private Output<String> loadBalancerName;

    /**
     * @return The load balancer to attach the policy to.
     * 
     */
    public Output<String> loadBalancerName() {
        return this.loadBalancerName;
    }
    /**
     * List of Policy Names to apply to the backend server.
     * 
     */
    @Export(name="policyNames", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> policyNames;

    /**
     * @return List of Policy Names to apply to the backend server.
     * 
     */
    public Output<Optional<List<String>>> policyNames() {
        return Codegen.optional(this.policyNames);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public LoadBalancerBackendServerPolicy(String name) {
        this(name, LoadBalancerBackendServerPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public LoadBalancerBackendServerPolicy(String name, LoadBalancerBackendServerPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public LoadBalancerBackendServerPolicy(String name, LoadBalancerBackendServerPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:elasticloadbalancing/loadBalancerBackendServerPolicy:LoadBalancerBackendServerPolicy", name, args == null ? LoadBalancerBackendServerPolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private LoadBalancerBackendServerPolicy(String name, Output<String> id, @Nullable LoadBalancerBackendServerPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:elasticloadbalancing/loadBalancerBackendServerPolicy:LoadBalancerBackendServerPolicy", name, state, makeResourceOptions(options, id));
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
    public static LoadBalancerBackendServerPolicy get(String name, Output<String> id, @Nullable LoadBalancerBackendServerPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new LoadBalancerBackendServerPolicy(name, id, state, options);
    }
}
