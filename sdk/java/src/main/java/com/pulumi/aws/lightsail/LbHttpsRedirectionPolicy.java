// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.lightsail;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.lightsail.LbHttpsRedirectionPolicyArgs;
import com.pulumi.aws.lightsail.inputs.LbHttpsRedirectionPolicyState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Configures Https Redirection for a Lightsail Load Balancer. A valid Certificate must be attached to the load balancer in order to enable https redirection.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.lightsail.Lb;
 * import com.pulumi.aws.lightsail.LbArgs;
 * import com.pulumi.aws.lightsail.LbCertificate;
 * import com.pulumi.aws.lightsail.LbCertificateArgs;
 * import com.pulumi.aws.lightsail.LbCertificateAttachment;
 * import com.pulumi.aws.lightsail.LbCertificateAttachmentArgs;
 * import com.pulumi.aws.lightsail.LbHttpsRedirectionPolicy;
 * import com.pulumi.aws.lightsail.LbHttpsRedirectionPolicyArgs;
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
 *         var testLb = new Lb(&#34;testLb&#34;, LbArgs.builder()        
 *             .healthCheckPath(&#34;/&#34;)
 *             .instancePort(&#34;80&#34;)
 *             .tags(Map.of(&#34;foo&#34;, &#34;bar&#34;))
 *             .build());
 * 
 *         var testLbCertificate = new LbCertificate(&#34;testLbCertificate&#34;, LbCertificateArgs.builder()        
 *             .lbName(testLb.id())
 *             .domainName(&#34;test.com&#34;)
 *             .build());
 * 
 *         var testLbCertificateAttachment = new LbCertificateAttachment(&#34;testLbCertificateAttachment&#34;, LbCertificateAttachmentArgs.builder()        
 *             .lbName(testLb.name())
 *             .certificateName(testLbCertificate.name())
 *             .build());
 * 
 *         var testLbHttpsRedirectionPolicy = new LbHttpsRedirectionPolicy(&#34;testLbHttpsRedirectionPolicy&#34;, LbHttpsRedirectionPolicyArgs.builder()        
 *             .lbName(testLb.name())
 *             .enabled(true)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * `aws_lightsail_lb_https_redirection_policy` can be imported by using the `lb_name` attribute, e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:lightsail/lbHttpsRedirectionPolicy:LbHttpsRedirectionPolicy test example-load-balancer
 * ```
 * 
 */
@ResourceType(type="aws:lightsail/lbHttpsRedirectionPolicy:LbHttpsRedirectionPolicy")
public class LbHttpsRedirectionPolicy extends com.pulumi.resources.CustomResource {
    /**
     * - The Https Redirection state of the load balancer. `true` to activate http to https redirection or `false` to deactivate http to https redirection.
     * 
     */
    @Export(name="enabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> enabled;

    /**
     * @return - The Https Redirection state of the load balancer. `true` to activate http to https redirection or `false` to deactivate http to https redirection.
     * 
     */
    public Output<Boolean> enabled() {
        return this.enabled;
    }
    /**
     * The name of the load balancer to which you want to enable http to https redirection.
     * 
     */
    @Export(name="lbName", refs={String.class}, tree="[0]")
    private Output<String> lbName;

    /**
     * @return The name of the load balancer to which you want to enable http to https redirection.
     * 
     */
    public Output<String> lbName() {
        return this.lbName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public LbHttpsRedirectionPolicy(String name) {
        this(name, LbHttpsRedirectionPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public LbHttpsRedirectionPolicy(String name, LbHttpsRedirectionPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public LbHttpsRedirectionPolicy(String name, LbHttpsRedirectionPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:lightsail/lbHttpsRedirectionPolicy:LbHttpsRedirectionPolicy", name, args == null ? LbHttpsRedirectionPolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private LbHttpsRedirectionPolicy(String name, Output<String> id, @Nullable LbHttpsRedirectionPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:lightsail/lbHttpsRedirectionPolicy:LbHttpsRedirectionPolicy", name, state, makeResourceOptions(options, id));
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
    public static LbHttpsRedirectionPolicy get(String name, Output<String> id, @Nullable LbHttpsRedirectionPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new LbHttpsRedirectionPolicy(name, id, state, options);
    }
}
