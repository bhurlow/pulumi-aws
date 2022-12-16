// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.eks;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.eks.NodeGroupArgs;
import com.pulumi.aws.eks.inputs.NodeGroupState;
import com.pulumi.aws.eks.outputs.NodeGroupLaunchTemplate;
import com.pulumi.aws.eks.outputs.NodeGroupRemoteAccess;
import com.pulumi.aws.eks.outputs.NodeGroupResource;
import com.pulumi.aws.eks.outputs.NodeGroupScalingConfig;
import com.pulumi.aws.eks.outputs.NodeGroupTaint;
import com.pulumi.aws.eks.outputs.NodeGroupUpdateConfig;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages an EKS Node Group, which can provision and optionally update an Auto Scaling Group of Kubernetes worker nodes compatible with EKS. Additional documentation about this functionality can be found in the [EKS User Guide](https://docs.aws.amazon.com/eks/latest/userguide/managed-node-groups.html).
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.eks.NodeGroup;
 * import com.pulumi.aws.eks.NodeGroupArgs;
 * import com.pulumi.aws.eks.inputs.NodeGroupScalingConfigArgs;
 * import com.pulumi.aws.eks.inputs.NodeGroupUpdateConfigArgs;
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
 *         var example = new NodeGroup(&#34;example&#34;, NodeGroupArgs.builder()        
 *             .clusterName(aws_eks_cluster.example().name())
 *             .nodeRoleArn(aws_iam_role.example().arn())
 *             .subnetIds(aws_subnet.example().stream().map(element -&gt; element.id()).collect(toList()))
 *             .scalingConfig(NodeGroupScalingConfigArgs.builder()
 *                 .desiredSize(1)
 *                 .maxSize(2)
 *                 .minSize(1)
 *                 .build())
 *             .updateConfig(NodeGroupUpdateConfigArgs.builder()
 *                 .maxUnavailable(1)
 *                 .build())
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(                
 *                     aws_iam_role_policy_attachment.example-AmazonEKSWorkerNodePolicy(),
 *                     aws_iam_role_policy_attachment.example-AmazonEKS_CNI_Policy(),
 *                     aws_iam_role_policy_attachment.example-AmazonEC2ContainerRegistryReadOnly())
 *                 .build());
 * 
 *     }
 * }
 * ```
 * ### Ignoring Changes to Desired Size
 * 
 * You can utilize [ignoreChanges](https://www.pulumi.com/docs/intro/concepts/programming-model/#ignorechanges) create an EKS Node Group with an initial size of running instances, then ignore any changes to that count caused externally (e.g. Application Autoscaling).
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.eks.NodeGroup;
 * import com.pulumi.aws.eks.NodeGroupArgs;
 * import com.pulumi.aws.eks.inputs.NodeGroupScalingConfigArgs;
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
 *         var example = new NodeGroup(&#34;example&#34;, NodeGroupArgs.builder()        
 *             .scalingConfig(NodeGroupScalingConfigArgs.builder()
 *                 .desiredSize(2)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Example IAM Role for EKS Node Group
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.iam.Role;
 * import com.pulumi.aws.iam.RoleArgs;
 * import com.pulumi.aws.iam.RolePolicyAttachment;
 * import com.pulumi.aws.iam.RolePolicyAttachmentArgs;
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
 *         var example = new Role(&#34;example&#34;, RoleArgs.builder()        
 *             .assumeRolePolicy(serializeJson(
 *                 jsonObject(
 *                     jsonProperty(&#34;Statement&#34;, jsonArray(jsonObject(
 *                         jsonProperty(&#34;Action&#34;, &#34;sts:AssumeRole&#34;),
 *                         jsonProperty(&#34;Effect&#34;, &#34;Allow&#34;),
 *                         jsonProperty(&#34;Principal&#34;, jsonObject(
 *                             jsonProperty(&#34;Service&#34;, &#34;ec2.amazonaws.com&#34;)
 *                         ))
 *                     ))),
 *                     jsonProperty(&#34;Version&#34;, &#34;2012-10-17&#34;)
 *                 )))
 *             .build());
 * 
 *         var example_AmazonEKSWorkerNodePolicy = new RolePolicyAttachment(&#34;example-AmazonEKSWorkerNodePolicy&#34;, RolePolicyAttachmentArgs.builder()        
 *             .policyArn(&#34;arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy&#34;)
 *             .role(example.name())
 *             .build());
 * 
 *         var example_AmazonEKSCNIPolicy = new RolePolicyAttachment(&#34;example-AmazonEKSCNIPolicy&#34;, RolePolicyAttachmentArgs.builder()        
 *             .policyArn(&#34;arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy&#34;)
 *             .role(example.name())
 *             .build());
 * 
 *         var example_AmazonEC2ContainerRegistryReadOnly = new RolePolicyAttachment(&#34;example-AmazonEC2ContainerRegistryReadOnly&#34;, RolePolicyAttachmentArgs.builder()        
 *             .policyArn(&#34;arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly&#34;)
 *             .role(example.name())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * EKS Node Groups can be imported using the `cluster_name` and `node_group_name` separated by a colon (`:`), e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:eks/nodeGroup:NodeGroup my_node_group my_cluster:my_node_group
 * ```
 * 
 */
@ResourceType(type="aws:eks/nodeGroup:NodeGroup")
public class NodeGroup extends com.pulumi.resources.CustomResource {
    /**
     * Type of Amazon Machine Image (AMI) associated with the EKS Node Group. See the [AWS documentation](https://docs.aws.amazon.com/eks/latest/APIReference/API_Nodegroup.html#AmazonEKS-Type-Nodegroup-amiType) for valid values. This provider will only perform drift detection if a configuration value is provided.
     * 
     */
    @Export(name="amiType", refs={String.class}, tree="[0]")
    private Output<String> amiType;

    /**
     * @return Type of Amazon Machine Image (AMI) associated with the EKS Node Group. See the [AWS documentation](https://docs.aws.amazon.com/eks/latest/APIReference/API_Nodegroup.html#AmazonEKS-Type-Nodegroup-amiType) for valid values. This provider will only perform drift detection if a configuration value is provided.
     * 
     */
    public Output<String> amiType() {
        return this.amiType;
    }
    /**
     * Amazon Resource Name (ARN) of the EKS Node Group.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return Amazon Resource Name (ARN) of the EKS Node Group.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * Type of capacity associated with the EKS Node Group. Valid values: `ON_DEMAND`, `SPOT`. This provider will only perform drift detection if a configuration value is provided.
     * 
     */
    @Export(name="capacityType", refs={String.class}, tree="[0]")
    private Output<String> capacityType;

    /**
     * @return Type of capacity associated with the EKS Node Group. Valid values: `ON_DEMAND`, `SPOT`. This provider will only perform drift detection if a configuration value is provided.
     * 
     */
    public Output<String> capacityType() {
        return this.capacityType;
    }
    /**
     * Name of the EKS Cluster. Must be between 1-100 characters in length. Must begin with an alphanumeric character, and must only contain alphanumeric characters, dashes and underscores (`^[0-9A-Za-z][A-Za-z0-9\-_]+$`).
     * 
     */
    @Export(name="clusterName", refs={String.class}, tree="[0]")
    private Output<String> clusterName;

    /**
     * @return Name of the EKS Cluster. Must be between 1-100 characters in length. Must begin with an alphanumeric character, and must only contain alphanumeric characters, dashes and underscores (`^[0-9A-Za-z][A-Za-z0-9\-_]+$`).
     * 
     */
    public Output<String> clusterName() {
        return this.clusterName;
    }
    /**
     * Disk size in GiB for worker nodes. Defaults to `20`. The provider will only perform drift detection if a configuration value is provided.
     * 
     */
    @Export(name="diskSize", refs={Integer.class}, tree="[0]")
    private Output<Integer> diskSize;

    /**
     * @return Disk size in GiB for worker nodes. Defaults to `20`. The provider will only perform drift detection if a configuration value is provided.
     * 
     */
    public Output<Integer> diskSize() {
        return this.diskSize;
    }
    /**
     * Force version update if existing pods are unable to be drained due to a pod disruption budget issue.
     * 
     */
    @Export(name="forceUpdateVersion", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> forceUpdateVersion;

    /**
     * @return Force version update if existing pods are unable to be drained due to a pod disruption budget issue.
     * 
     */
    public Output<Optional<Boolean>> forceUpdateVersion() {
        return Codegen.optional(this.forceUpdateVersion);
    }
    /**
     * List of instance types associated with the EKS Node Group. Defaults to `[&#34;t3.medium&#34;]`. The provider will only perform drift detection if a configuration value is provided.
     * 
     */
    @Export(name="instanceTypes", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> instanceTypes;

    /**
     * @return List of instance types associated with the EKS Node Group. Defaults to `[&#34;t3.medium&#34;]`. The provider will only perform drift detection if a configuration value is provided.
     * 
     */
    public Output<List<String>> instanceTypes() {
        return this.instanceTypes;
    }
    /**
     * Key-value map of Kubernetes labels. Only labels that are applied with the EKS API are managed by this argument. Other Kubernetes labels applied to the EKS Node Group will not be managed.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return Key-value map of Kubernetes labels. Only labels that are applied with the EKS API are managed by this argument. Other Kubernetes labels applied to the EKS Node Group will not be managed.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * Configuration block with Launch Template settings. Detailed below.
     * 
     */
    @Export(name="launchTemplate", refs={NodeGroupLaunchTemplate.class}, tree="[0]")
    private Output</* @Nullable */ NodeGroupLaunchTemplate> launchTemplate;

    /**
     * @return Configuration block with Launch Template settings. Detailed below.
     * 
     */
    public Output<Optional<NodeGroupLaunchTemplate>> launchTemplate() {
        return Codegen.optional(this.launchTemplate);
    }
    /**
     * Name of the EKS Node Group. If omitted, the provider will assign a random, unique name. Conflicts with `node_group_name_prefix`.
     * 
     */
    @Export(name="nodeGroupName", refs={String.class}, tree="[0]")
    private Output<String> nodeGroupName;

    /**
     * @return Name of the EKS Node Group. If omitted, the provider will assign a random, unique name. Conflicts with `node_group_name_prefix`.
     * 
     */
    public Output<String> nodeGroupName() {
        return this.nodeGroupName;
    }
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `node_group_name`.
     * 
     */
    @Export(name="nodeGroupNamePrefix", refs={String.class}, tree="[0]")
    private Output<String> nodeGroupNamePrefix;

    /**
     * @return Creates a unique name beginning with the specified prefix. Conflicts with `node_group_name`.
     * 
     */
    public Output<String> nodeGroupNamePrefix() {
        return this.nodeGroupNamePrefix;
    }
    /**
     * Amazon Resource Name (ARN) of the IAM Role that provides permissions for the EKS Node Group.
     * 
     */
    @Export(name="nodeRoleArn", refs={String.class}, tree="[0]")
    private Output<String> nodeRoleArn;

    /**
     * @return Amazon Resource Name (ARN) of the IAM Role that provides permissions for the EKS Node Group.
     * 
     */
    public Output<String> nodeRoleArn() {
        return this.nodeRoleArn;
    }
    /**
     * AMI version of the EKS Node Group. Defaults to latest version for Kubernetes version.
     * 
     */
    @Export(name="releaseVersion", refs={String.class}, tree="[0]")
    private Output<String> releaseVersion;

    /**
     * @return AMI version of the EKS Node Group. Defaults to latest version for Kubernetes version.
     * 
     */
    public Output<String> releaseVersion() {
        return this.releaseVersion;
    }
    /**
     * Configuration block with remote access settings. Detailed below.
     * 
     */
    @Export(name="remoteAccess", refs={NodeGroupRemoteAccess.class}, tree="[0]")
    private Output</* @Nullable */ NodeGroupRemoteAccess> remoteAccess;

    /**
     * @return Configuration block with remote access settings. Detailed below.
     * 
     */
    public Output<Optional<NodeGroupRemoteAccess>> remoteAccess() {
        return Codegen.optional(this.remoteAccess);
    }
    /**
     * List of objects containing information about underlying resources.
     * 
     */
    @Export(name="resources", refs={List.class,NodeGroupResource.class}, tree="[0,1]")
    private Output<List<NodeGroupResource>> resources;

    /**
     * @return List of objects containing information about underlying resources.
     * 
     */
    public Output<List<NodeGroupResource>> resources() {
        return this.resources;
    }
    /**
     * Configuration block with scaling settings. Detailed below.
     * 
     */
    @Export(name="scalingConfig", refs={NodeGroupScalingConfig.class}, tree="[0]")
    private Output<NodeGroupScalingConfig> scalingConfig;

    /**
     * @return Configuration block with scaling settings. Detailed below.
     * 
     */
    public Output<NodeGroupScalingConfig> scalingConfig() {
        return this.scalingConfig;
    }
    /**
     * Status of the EKS Node Group.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return Status of the EKS Node Group.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * Identifiers of EC2 Subnets to associate with the EKS Node Group. These subnets must have the following resource tag: `kubernetes.io/cluster/CLUSTER_NAME` (where `CLUSTER_NAME` is replaced with the name of the EKS Cluster).
     * 
     */
    @Export(name="subnetIds", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> subnetIds;

    /**
     * @return Identifiers of EC2 Subnets to associate with the EKS Node Group. These subnets must have the following resource tag: `kubernetes.io/cluster/CLUSTER_NAME` (where `CLUSTER_NAME` is replaced with the name of the EKS Cluster).
     * 
     */
    public Output<List<String>> subnetIds() {
        return this.subnetIds;
    }
    /**
     * Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
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
     * The Kubernetes taints to be applied to the nodes in the node group. Maximum of 50 taints per node group. Detailed below.
     * 
     */
    @Export(name="taints", refs={List.class,NodeGroupTaint.class}, tree="[0,1]")
    private Output</* @Nullable */ List<NodeGroupTaint>> taints;

    /**
     * @return The Kubernetes taints to be applied to the nodes in the node group. Maximum of 50 taints per node group. Detailed below.
     * 
     */
    public Output<Optional<List<NodeGroupTaint>>> taints() {
        return Codegen.optional(this.taints);
    }
    @Export(name="updateConfig", refs={NodeGroupUpdateConfig.class}, tree="[0]")
    private Output<NodeGroupUpdateConfig> updateConfig;

    public Output<NodeGroupUpdateConfig> updateConfig() {
        return this.updateConfig;
    }
    /**
     * EC2 Launch Template version number. While the API accepts values like `$Default` and `$Latest`, the API will convert the value to the associated version number (e.g., `1`) on read and the provider will show a difference on next plan. Using the `default_version` or `latest_version` attribute of the `aws.ec2.LaunchTemplate` resource or data source is recommended for this argument.
     * 
     */
    @Export(name="version", refs={String.class}, tree="[0]")
    private Output<String> version;

    /**
     * @return EC2 Launch Template version number. While the API accepts values like `$Default` and `$Latest`, the API will convert the value to the associated version number (e.g., `1`) on read and the provider will show a difference on next plan. Using the `default_version` or `latest_version` attribute of the `aws.ec2.LaunchTemplate` resource or data source is recommended for this argument.
     * 
     */
    public Output<String> version() {
        return this.version;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public NodeGroup(String name) {
        this(name, NodeGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public NodeGroup(String name, NodeGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public NodeGroup(String name, NodeGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:eks/nodeGroup:NodeGroup", name, args == null ? NodeGroupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private NodeGroup(String name, Output<String> id, @Nullable NodeGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:eks/nodeGroup:NodeGroup", name, state, makeResourceOptions(options, id));
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
    public static NodeGroup get(String name, Output<String> id, @Nullable NodeGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new NodeGroup(name, id, state, options);
    }
}
