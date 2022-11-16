// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.msk;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.msk.ScramSecretAssociationArgs;
import com.pulumi.aws.msk.inputs.ScramSecretAssociationState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Associates SCRAM secrets stored in the Secrets Manager service with a Managed Streaming for Kafka (MSK) cluster.
 * 
 * &gt; **Note:** The following assumes the MSK cluster has SASL/SCRAM authentication enabled. See below for example usage or refer to the [Username/Password Authentication](https://docs.aws.amazon.com/msk/latest/developerguide/msk-password.html) section of the MSK Developer Guide for more details.
 * 
 * To set up username and password authentication for a cluster, create an `aws.secretsmanager.Secret` resource and associate
 * a username and password with the secret with an `aws.secretsmanager.SecretVersion` resource. When creating a secret for the cluster,
 * the `name` must have the prefix `AmazonMSK_` and you must either use an existing custom AWS KMS key or create a new
 * custom AWS KMS key for your secret with the `aws.kms.Key` resource. It is important to note that a policy is required for the `aws.secretsmanager.Secret`
 * resource in order for Kafka to be able to read it. This policy is attached automatically when the `aws.msk.ScramSecretAssociation` is used,
 * however, this policy will not be in the state and as such, will present a diff on plan/apply. For that reason, you must use the `aws.secretsmanager.SecretPolicy`
 * resource](/docs/providers/aws/r/secretsmanager_secret_policy.html) as shown below in order to ensure that the state is in a clean state after the creation of secret and the association to the cluster.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.msk.Cluster;
 * import com.pulumi.aws.msk.ClusterArgs;
 * import com.pulumi.aws.msk.inputs.ClusterClientAuthenticationArgs;
 * import com.pulumi.aws.msk.inputs.ClusterClientAuthenticationSaslArgs;
 * import com.pulumi.aws.kms.Key;
 * import com.pulumi.aws.kms.KeyArgs;
 * import com.pulumi.aws.secretsmanager.Secret;
 * import com.pulumi.aws.secretsmanager.SecretArgs;
 * import com.pulumi.aws.secretsmanager.SecretVersion;
 * import com.pulumi.aws.secretsmanager.SecretVersionArgs;
 * import com.pulumi.aws.msk.ScramSecretAssociation;
 * import com.pulumi.aws.msk.ScramSecretAssociationArgs;
 * import com.pulumi.aws.secretsmanager.SecretPolicy;
 * import com.pulumi.aws.secretsmanager.SecretPolicyArgs;
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
 *         var exampleCluster = new Cluster(&#34;exampleCluster&#34;, ClusterArgs.builder()        
 *             .clientAuthentication(ClusterClientAuthenticationArgs.builder()
 *                 .sasl(ClusterClientAuthenticationSaslArgs.builder()
 *                     .scram(true)
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *         var exampleKey = new Key(&#34;exampleKey&#34;, KeyArgs.builder()        
 *             .description(&#34;Example Key for MSK Cluster Scram Secret Association&#34;)
 *             .build());
 * 
 *         var exampleSecret = new Secret(&#34;exampleSecret&#34;, SecretArgs.builder()        
 *             .kmsKeyId(exampleKey.keyId())
 *             .build());
 * 
 *         var exampleSecretVersion = new SecretVersion(&#34;exampleSecretVersion&#34;, SecretVersionArgs.builder()        
 *             .secretId(exampleSecret.id())
 *             .secretString(serializeJson(
 *                 jsonObject(
 *                     jsonProperty(&#34;username&#34;, &#34;user&#34;),
 *                     jsonProperty(&#34;password&#34;, &#34;pass&#34;)
 *                 )))
 *             .build());
 * 
 *         var exampleScramSecretAssociation = new ScramSecretAssociation(&#34;exampleScramSecretAssociation&#34;, ScramSecretAssociationArgs.builder()        
 *             .clusterArn(exampleCluster.arn())
 *             .secretArnLists(exampleSecret.arn())
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(exampleSecretVersion)
 *                 .build());
 * 
 *         var exampleSecretPolicy = new SecretPolicy(&#34;exampleSecretPolicy&#34;, SecretPolicyArgs.builder()        
 *             .secretArn(exampleSecret.arn())
 *             .policy(exampleSecret.arn().applyValue(arn -&gt; &#34;&#34;&#34;
 * {
 *   &#34;Version&#34; : &#34;2012-10-17&#34;,
 *   &#34;Statement&#34; : [ {
 *     &#34;Sid&#34;: &#34;AWSKafkaResourcePolicy&#34;,
 *     &#34;Effect&#34; : &#34;Allow&#34;,
 *     &#34;Principal&#34; : {
 *       &#34;Service&#34; : &#34;kafka.amazonaws.com&#34;
 *     },
 *     &#34;Action&#34; : &#34;secretsmanager:getSecretValue&#34;,
 *     &#34;Resource&#34; : &#34;%s&#34;
 *   } ]
 * }
 * &#34;, arn)))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * MSK SCRAM Secret Associations can be imported using the `id` e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:msk/scramSecretAssociation:ScramSecretAssociation example arn:aws:kafka:us-west-2:123456789012:cluster/example/279c0212-d057-4dba-9aa9-1c4e5a25bfc7-3
 * ```
 * 
 */
@ResourceType(type="aws:msk/scramSecretAssociation:ScramSecretAssociation")
public class ScramSecretAssociation extends com.pulumi.resources.CustomResource {
    /**
     * Amazon Resource Name (ARN) of the MSK cluster.
     * 
     */
    @Export(name="clusterArn", type=String.class, parameters={})
    private Output<String> clusterArn;

    /**
     * @return Amazon Resource Name (ARN) of the MSK cluster.
     * 
     */
    public Output<String> clusterArn() {
        return this.clusterArn;
    }
    /**
     * List of AWS Secrets Manager secret ARNs.
     * 
     */
    @Export(name="secretArnLists", type=List.class, parameters={String.class})
    private Output<List<String>> secretArnLists;

    /**
     * @return List of AWS Secrets Manager secret ARNs.
     * 
     */
    public Output<List<String>> secretArnLists() {
        return this.secretArnLists;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ScramSecretAssociation(String name) {
        this(name, ScramSecretAssociationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ScramSecretAssociation(String name, ScramSecretAssociationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ScramSecretAssociation(String name, ScramSecretAssociationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:msk/scramSecretAssociation:ScramSecretAssociation", name, args == null ? ScramSecretAssociationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ScramSecretAssociation(String name, Output<String> id, @Nullable ScramSecretAssociationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:msk/scramSecretAssociation:ScramSecretAssociation", name, state, makeResourceOptions(options, id));
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
    public static ScramSecretAssociation get(String name, Output<String> id, @Nullable ScramSecretAssociationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ScramSecretAssociation(name, id, state, options);
    }
}
