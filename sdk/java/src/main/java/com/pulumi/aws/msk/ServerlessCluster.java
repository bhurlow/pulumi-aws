// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.msk;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.msk.ServerlessClusterArgs;
import com.pulumi.aws.msk.inputs.ServerlessClusterState;
import com.pulumi.aws.msk.outputs.ServerlessClusterClientAuthentication;
import com.pulumi.aws.msk.outputs.ServerlessClusterVpcConfig;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages an Amazon MSK Serverless cluster.
 * 
 * &gt; **Note:** To manage a _provisioned_ Amazon MSK cluster, use the `aws.msk.Cluster` resource.
 * 
 * ## Import
 * 
 * Using `pulumi import`, import MSK serverless clusters using the cluster `arn`. For example:
 * 
 * ```sh
 *  $ pulumi import aws:msk/serverlessCluster:ServerlessCluster example arn:aws:kafka:us-west-2:123456789012:cluster/example/279c0212-d057-4dba-9aa9-1c4e5a25bfc7-3
 * ```
 * 
 */
@ResourceType(type="aws:msk/serverlessCluster:ServerlessCluster")
public class ServerlessCluster extends com.pulumi.resources.CustomResource {
    /**
     * The ARN of the serverless cluster.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The ARN of the serverless cluster.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * Specifies client authentication information for the serverless cluster. See below.
     * 
     */
    @Export(name="clientAuthentication", refs={ServerlessClusterClientAuthentication.class}, tree="[0]")
    private Output<ServerlessClusterClientAuthentication> clientAuthentication;

    /**
     * @return Specifies client authentication information for the serverless cluster. See below.
     * 
     */
    public Output<ServerlessClusterClientAuthentication> clientAuthentication() {
        return this.clientAuthentication;
    }
    /**
     * The name of the serverless cluster.
     * 
     */
    @Export(name="clusterName", refs={String.class}, tree="[0]")
    private Output<String> clusterName;

    /**
     * @return The name of the serverless cluster.
     * 
     */
    public Output<String> clusterName() {
        return this.clusterName;
    }
    /**
     * A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    @Export(name="tagsAll", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> tagsAll;

    /**
     * @return A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    public Output<Map<String,String>> tagsAll() {
        return this.tagsAll;
    }
    /**
     * VPC configuration information. See below.
     * 
     */
    @Export(name="vpcConfigs", refs={List.class,ServerlessClusterVpcConfig.class}, tree="[0,1]")
    private Output<List<ServerlessClusterVpcConfig>> vpcConfigs;

    /**
     * @return VPC configuration information. See below.
     * 
     */
    public Output<List<ServerlessClusterVpcConfig>> vpcConfigs() {
        return this.vpcConfigs;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ServerlessCluster(String name) {
        this(name, ServerlessClusterArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServerlessCluster(String name, ServerlessClusterArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServerlessCluster(String name, ServerlessClusterArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:msk/serverlessCluster:ServerlessCluster", name, args == null ? ServerlessClusterArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ServerlessCluster(String name, Output<String> id, @Nullable ServerlessClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:msk/serverlessCluster:ServerlessCluster", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "tagsAll"
            ))
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
    public static ServerlessCluster get(String name, Output<String> id, @Nullable ServerlessClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServerlessCluster(name, id, state, options);
    }
}
