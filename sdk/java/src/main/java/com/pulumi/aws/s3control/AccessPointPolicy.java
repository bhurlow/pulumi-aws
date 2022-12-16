// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.s3control;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.s3control.AccessPointPolicyArgs;
import com.pulumi.aws.s3control.inputs.AccessPointPolicyState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a resource to manage an S3 Access Point resource policy.
 * 
 * &gt; **NOTE on Access Points and Access Point Policies:** The provider provides both a standalone Access Point Policy resource and an Access Point resource with a resource policy defined in-line. You cannot use an Access Point with in-line resource policy in conjunction with an Access Point Policy resource. Doing so will cause a conflict of policies and will overwrite the access point&#39;s resource policy.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.s3.BucketV2;
 * import com.pulumi.aws.s3.AccessPoint;
 * import com.pulumi.aws.s3.AccessPointArgs;
 * import com.pulumi.aws.s3.inputs.AccessPointPublicAccessBlockConfigurationArgs;
 * import com.pulumi.aws.s3control.AccessPointPolicy;
 * import com.pulumi.aws.s3control.AccessPointPolicyArgs;
 * import static com.pulumi.codegen.internal.Serialization.*;
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
 *         var exampleBucketV2 = new BucketV2(&#34;exampleBucketV2&#34;);
 * 
 *         var exampleAccessPoint = new AccessPoint(&#34;exampleAccessPoint&#34;, AccessPointArgs.builder()        
 *             .bucket(exampleBucketV2.id())
 *             .publicAccessBlockConfiguration(AccessPointPublicAccessBlockConfigurationArgs.builder()
 *                 .blockPublicAcls(true)
 *                 .blockPublicPolicy(false)
 *                 .ignorePublicAcls(true)
 *                 .restrictPublicBuckets(false)
 *                 .build())
 *             .build());
 * 
 *         var exampleAccessPointPolicy = new AccessPointPolicy(&#34;exampleAccessPointPolicy&#34;, AccessPointPolicyArgs.builder()        
 *             .accessPointArn(exampleAccessPoint.arn())
 *             .policy(exampleAccessPoint.arn().applyValue(arn -&gt; serializeJson(
 *                 jsonObject(
 *                     jsonProperty(&#34;Version&#34;, &#34;2008-10-17&#34;),
 *                     jsonProperty(&#34;Statement&#34;, jsonArray(jsonObject(
 *                         jsonProperty(&#34;Effect&#34;, &#34;Allow&#34;),
 *                         jsonProperty(&#34;Action&#34;, &#34;s3:GetObjectTagging&#34;),
 *                         jsonProperty(&#34;Principal&#34;, jsonObject(
 *                             jsonProperty(&#34;AWS&#34;, &#34;*&#34;)
 *                         )),
 *                         jsonProperty(&#34;Resource&#34;, String.format(&#34;%s/object/*&#34;, arn))
 *                     )))
 *                 ))))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Access Point policies can be imported using the `access_point_arn`, e.g.
 * 
 * ```sh
 *  $ pulumi import aws:s3control/accessPointPolicy:AccessPointPolicy example arn:aws:s3:us-west-2:123456789012:accesspoint/example
 * ```
 * 
 */
@ResourceType(type="aws:s3control/accessPointPolicy:AccessPointPolicy")
public class AccessPointPolicy extends com.pulumi.resources.CustomResource {
    /**
     * The ARN of the access point that you want to associate with the specified policy.
     * 
     */
    @Export(name="accessPointArn", refs={String.class}, tree="[0]")
    private Output<String> accessPointArn;

    /**
     * @return The ARN of the access point that you want to associate with the specified policy.
     * 
     */
    public Output<String> accessPointArn() {
        return this.accessPointArn;
    }
    /**
     * Indicates whether this access point currently has a policy that allows public access.
     * 
     */
    @Export(name="hasPublicAccessPolicy", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> hasPublicAccessPolicy;

    /**
     * @return Indicates whether this access point currently has a policy that allows public access.
     * 
     */
    public Output<Boolean> hasPublicAccessPolicy() {
        return this.hasPublicAccessPolicy;
    }
    /**
     * The policy that you want to apply to the specified access point.
     * 
     */
    @Export(name="policy", refs={String.class}, tree="[0]")
    private Output<String> policy;

    /**
     * @return The policy that you want to apply to the specified access point.
     * 
     */
    public Output<String> policy() {
        return this.policy;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AccessPointPolicy(String name) {
        this(name, AccessPointPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AccessPointPolicy(String name, AccessPointPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AccessPointPolicy(String name, AccessPointPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:s3control/accessPointPolicy:AccessPointPolicy", name, args == null ? AccessPointPolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AccessPointPolicy(String name, Output<String> id, @Nullable AccessPointPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:s3control/accessPointPolicy:AccessPointPolicy", name, state, makeResourceOptions(options, id));
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
    public static AccessPointPolicy get(String name, Output<String> id, @Nullable AccessPointPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AccessPointPolicy(name, id, state, options);
    }
}
