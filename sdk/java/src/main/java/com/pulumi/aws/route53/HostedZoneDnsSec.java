// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.route53;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.route53.HostedZoneDnsSecArgs;
import com.pulumi.aws.route53.inputs.HostedZoneDnsSecState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages Route 53 Hosted Zone Domain Name System Security Extensions (DNSSEC). For more information about managing DNSSEC in Route 53, see the [Route 53 Developer Guide](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-configuring-dnssec.html).
 * 
 * !&gt; **WARNING:** If you disable DNSSEC signing for your hosted zone before the DNS changes have propagated, your domain could become unavailable on the internet. When you remove the DS records, you must wait until the longest TTL for the DS records that you remove has expired before you complete the step to disable DNSSEC signing. Please refer to the [Route 53 Developer Guide - Disable DNSSEC](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-configuring-dnssec-disable.html) for a detailed breakdown on the steps required to disable DNSSEC safely for a hosted zone.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.AwsFunctions;
 * import com.pulumi.aws.kms.Key;
 * import com.pulumi.aws.kms.KeyArgs;
 * import com.pulumi.aws.route53.Zone;
 * import com.pulumi.aws.route53.KeySigningKey;
 * import com.pulumi.aws.route53.KeySigningKeyArgs;
 * import com.pulumi.aws.route53.HostedZoneDnsSec;
 * import com.pulumi.aws.route53.HostedZoneDnsSecArgs;
 * import static com.pulumi.codegen.internal.Serialization.*;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         final var current = AwsFunctions.getCallerIdentity();
 * 
 *         var exampleKey = new Key(&#34;exampleKey&#34;, KeyArgs.builder()        
 *             .customerMasterKeySpec(&#34;ECC_NIST_P256&#34;)
 *             .deletionWindowInDays(7)
 *             .keyUsage(&#34;SIGN_VERIFY&#34;)
 *             .policy(serializeJson(
 *                 jsonObject(
 *                     jsonProperty(&#34;Statement&#34;, jsonArray(
 *                         jsonObject(
 *                             jsonProperty(&#34;Action&#34;, jsonArray(
 *                                 &#34;kms:DescribeKey&#34;, 
 *                                 &#34;kms:GetPublicKey&#34;, 
 *                                 &#34;kms:Sign&#34;, 
 *                                 &#34;kms:Verify&#34;
 *                             )),
 *                             jsonProperty(&#34;Effect&#34;, &#34;Allow&#34;),
 *                             jsonProperty(&#34;Principal&#34;, jsonObject(
 *                                 jsonProperty(&#34;Service&#34;, &#34;dnssec-route53.amazonaws.com&#34;)
 *                             )),
 *                             jsonProperty(&#34;Resource&#34;, &#34;*&#34;),
 *                             jsonProperty(&#34;Sid&#34;, &#34;Allow Route 53 DNSSEC Service&#34;)
 *                         ), 
 *                         jsonObject(
 *                             jsonProperty(&#34;Action&#34;, &#34;kms:*&#34;),
 *                             jsonProperty(&#34;Effect&#34;, &#34;Allow&#34;),
 *                             jsonProperty(&#34;Principal&#34;, jsonObject(
 *                                 jsonProperty(&#34;AWS&#34;, String.format(&#34;arn:aws:iam::%s:root&#34;, current.applyValue(getCallerIdentityResult -&gt; getCallerIdentityResult.accountId())))
 *                             )),
 *                             jsonProperty(&#34;Resource&#34;, &#34;*&#34;),
 *                             jsonProperty(&#34;Sid&#34;, &#34;Enable IAM User Permissions&#34;)
 *                         )
 *                     )),
 *                     jsonProperty(&#34;Version&#34;, &#34;2012-10-17&#34;)
 *                 )))
 *             .build());
 * 
 *         var exampleZone = new Zone(&#34;exampleZone&#34;);
 * 
 *         var exampleKeySigningKey = new KeySigningKey(&#34;exampleKeySigningKey&#34;, KeySigningKeyArgs.builder()        
 *             .hostedZoneId(exampleZone.id())
 *             .keyManagementServiceArn(exampleKey.arn())
 *             .build());
 * 
 *         var exampleHostedZoneDnsSec = new HostedZoneDnsSec(&#34;exampleHostedZoneDnsSec&#34;, HostedZoneDnsSecArgs.builder()        
 *             .hostedZoneId(exampleKeySigningKey.hostedZoneId())
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(exampleKeySigningKey)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * `aws_route53_hosted_zone_dnssec` resources can be imported by using the Route 53 Hosted Zone identifier, e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:route53/hostedZoneDnsSec:HostedZoneDnsSec example Z1D633PJN98FT9
 * ```
 * 
 */
@ResourceType(type="aws:route53/hostedZoneDnsSec:HostedZoneDnsSec")
public class HostedZoneDnsSec extends com.pulumi.resources.CustomResource {
    /**
     * Identifier of the Route 53 Hosted Zone.
     * 
     */
    @Export(name="hostedZoneId", refs={String.class}, tree="[0]")
    private Output<String> hostedZoneId;

    /**
     * @return Identifier of the Route 53 Hosted Zone.
     * 
     */
    public Output<String> hostedZoneId() {
        return this.hostedZoneId;
    }
    /**
     * Hosted Zone signing status. Valid values: `SIGNING`, `NOT_SIGNING`. Defaults to `SIGNING`.
     * 
     */
    @Export(name="signingStatus", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> signingStatus;

    /**
     * @return Hosted Zone signing status. Valid values: `SIGNING`, `NOT_SIGNING`. Defaults to `SIGNING`.
     * 
     */
    public Output<Optional<String>> signingStatus() {
        return Codegen.optional(this.signingStatus);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public HostedZoneDnsSec(String name) {
        this(name, HostedZoneDnsSecArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public HostedZoneDnsSec(String name, HostedZoneDnsSecArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public HostedZoneDnsSec(String name, HostedZoneDnsSecArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:route53/hostedZoneDnsSec:HostedZoneDnsSec", name, args == null ? HostedZoneDnsSecArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private HostedZoneDnsSec(String name, Output<String> id, @Nullable HostedZoneDnsSecState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:route53/hostedZoneDnsSec:HostedZoneDnsSec", name, state, makeResourceOptions(options, id));
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
    public static HostedZoneDnsSec get(String name, Output<String> id, @Nullable HostedZoneDnsSecState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new HostedZoneDnsSec(name, id, state, options);
    }
}
