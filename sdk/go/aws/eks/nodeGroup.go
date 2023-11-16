// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eks

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an EKS Node Group, which can provision and optionally update an Auto Scaling Group of Kubernetes worker nodes compatible with EKS. Additional documentation about this functionality can be found in the [EKS User Guide](https://docs.aws.amazon.com/eks/latest/userguide/managed-node-groups.html).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/eks"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
// func main() {
// pulumi.Run(func(ctx *pulumi.Context) error {
// var splat0 []interface{}
// for _, val0 := range aws_subnet.Example {
// splat0 = append(splat0, val0.Id)
// }
// _, err := eks.NewNodeGroup(ctx, "example", &eks.NodeGroupArgs{
// ClusterName: pulumi.Any(aws_eks_cluster.Example.Name),
// NodeRoleArn: pulumi.Any(aws_iam_role.Example.Arn),
// SubnetIds: toPulumiArray(splat0),
// ScalingConfig: &eks.NodeGroupScalingConfigArgs{
// DesiredSize: pulumi.Int(1),
// MaxSize: pulumi.Int(2),
// MinSize: pulumi.Int(1),
// },
// UpdateConfig: &eks.NodeGroupUpdateConfigArgs{
// MaxUnavailable: pulumi.Int(1),
// },
// }, pulumi.DependsOn([]pulumi.Resource{
// aws_iam_role_policy_attachment.ExampleAmazonEKSWorkerNodePolicy,
// aws_iam_role_policy_attachment.ExampleAmazonEKS_CNI_Policy,
// aws_iam_role_policy_attachment.ExampleAmazonEC2ContainerRegistryReadOnly,
// }))
// if err != nil {
// return err
// }
// return nil
// })
// }
// func toPulumiArray(arr []) pulumi.Array {
// var pulumiArr pulumi.Array
// for _, v := range arr {
// pulumiArr = append(pulumiArr, pulumi.(v))
// }
// return pulumiArr
// }
// ```
// ### Ignoring Changes to Desired Size
//
// You can utilize [ignoreChanges](https://www.pulumi.com/docs/intro/concepts/programming-model/#ignorechanges) create an EKS Node Group with an initial size of running instances, then ignore any changes to that count caused externally (e.g. Application Autoscaling).
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/eks"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := eks.NewNodeGroup(ctx, "example", &eks.NodeGroupArgs{
//				ScalingConfig: &eks.NodeGroupScalingConfigArgs{
//					DesiredSize: pulumi.Int(2),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Example IAM Role for EKS Node Group
//
// ```go
// package main
//
// import (
//
//	"encoding/json"
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			tmpJSON0, err := json.Marshal(map[string]interface{}{
//				"Statement": []map[string]interface{}{
//					map[string]interface{}{
//						"Action": "sts:AssumeRole",
//						"Effect": "Allow",
//						"Principal": map[string]interface{}{
//							"Service": "ec2.amazonaws.com",
//						},
//					},
//				},
//				"Version": "2012-10-17",
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			example, err := iam.NewRole(ctx, "example", &iam.RoleArgs{
//				AssumeRolePolicy: pulumi.String(json0),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = iam.NewRolePolicyAttachment(ctx, "example-AmazonEKSWorkerNodePolicy", &iam.RolePolicyAttachmentArgs{
//				PolicyArn: pulumi.String("arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy"),
//				Role:      example.Name,
//			})
//			if err != nil {
//				return err
//			}
//			_, err = iam.NewRolePolicyAttachment(ctx, "example-AmazonEKSCNIPolicy", &iam.RolePolicyAttachmentArgs{
//				PolicyArn: pulumi.String("arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy"),
//				Role:      example.Name,
//			})
//			if err != nil {
//				return err
//			}
//			_, err = iam.NewRolePolicyAttachment(ctx, "example-AmazonEC2ContainerRegistryReadOnly", &iam.RolePolicyAttachmentArgs{
//				PolicyArn: pulumi.String("arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly"),
//				Role:      example.Name,
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Using `pulumi import`, import EKS Node Groups using the `cluster_name` and `node_group_name` separated by a colon (`:`). For example:
//
// ```sh
//
//	$ pulumi import aws:eks/nodeGroup:NodeGroup my_node_group my_cluster:my_node_group
//
// ```
type NodeGroup struct {
	pulumi.CustomResourceState

	// Type of Amazon Machine Image (AMI) associated with the EKS Node Group. See the [AWS documentation](https://docs.aws.amazon.com/eks/latest/APIReference/API_Nodegroup.html#AmazonEKS-Type-Nodegroup-amiType) for valid values. This provider will only perform drift detection if a configuration value is provided.
	AmiType pulumi.StringOutput `pulumi:"amiType"`
	// Amazon Resource Name (ARN) of the EKS Node Group.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Type of capacity associated with the EKS Node Group. Valid values: `ON_DEMAND`, `SPOT`. This provider will only perform drift detection if a configuration value is provided.
	CapacityType pulumi.StringOutput `pulumi:"capacityType"`
	// Name of the EKS Cluster. Must be between 1-100 characters in length. Must begin with an alphanumeric character, and must only contain alphanumeric characters, dashes and underscores (`^[0-9A-Za-z][A-Za-z0-9\-_]+$`).
	ClusterName pulumi.StringOutput `pulumi:"clusterName"`
	// Disk size in GiB for worker nodes. Defaults to `50` for Windows, `20` all other node groups. The provider will only perform drift detection if a configuration value is provided.
	DiskSize pulumi.IntOutput `pulumi:"diskSize"`
	// Force version update if existing pods are unable to be drained due to a pod disruption budget issue.
	ForceUpdateVersion pulumi.BoolPtrOutput `pulumi:"forceUpdateVersion"`
	// List of instance types associated with the EKS Node Group. Defaults to `["t3.medium"]`. The provider will only perform drift detection if a configuration value is provided.
	InstanceTypes pulumi.StringArrayOutput `pulumi:"instanceTypes"`
	// Key-value map of Kubernetes labels. Only labels that are applied with the EKS API are managed by this argument. Other Kubernetes labels applied to the EKS Node Group will not be managed.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Configuration block with Launch Template settings. See `launchTemplate` below for details.
	LaunchTemplate NodeGroupLaunchTemplatePtrOutput `pulumi:"launchTemplate"`
	// Name of the EKS Node Group. If omitted, the provider will assign a random, unique name. Conflicts with `nodeGroupNamePrefix`. The node group name can't be longer than 63 characters. It must start with a letter or digit, but can also include hyphens and underscores for the remaining characters.
	NodeGroupName pulumi.StringOutput `pulumi:"nodeGroupName"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `nodeGroupName`.
	NodeGroupNamePrefix pulumi.StringOutput `pulumi:"nodeGroupNamePrefix"`
	// Amazon Resource Name (ARN) of the IAM Role that provides permissions for the EKS Node Group.
	NodeRoleArn pulumi.StringOutput `pulumi:"nodeRoleArn"`
	// AMI version of the EKS Node Group. Defaults to latest version for Kubernetes version.
	ReleaseVersion pulumi.StringOutput `pulumi:"releaseVersion"`
	// Configuration block with remote access settings. See `remoteAccess` below for details.
	RemoteAccess NodeGroupRemoteAccessPtrOutput `pulumi:"remoteAccess"`
	// List of objects containing information about underlying resources.
	Resources NodeGroupResourceArrayOutput `pulumi:"resources"`
	// Configuration block with scaling settings. See `scalingConfig` below for details.
	ScalingConfig NodeGroupScalingConfigOutput `pulumi:"scalingConfig"`
	// Status of the EKS Node Group.
	Status pulumi.StringOutput `pulumi:"status"`
	// Identifiers of EC2 Subnets to associate with the EKS Node Group.
	//
	// The following arguments are optional:
	SubnetIds pulumi.StringArrayOutput `pulumi:"subnetIds"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The Kubernetes taints to be applied to the nodes in the node group. Maximum of 50 taints per node group. See taint below for details.
	Taints NodeGroupTaintArrayOutput `pulumi:"taints"`
	// Configuration block with update settings. See `updateConfig` below for details.
	UpdateConfig NodeGroupUpdateConfigOutput `pulumi:"updateConfig"`
	// Kubernetes version. Defaults to EKS Cluster Kubernetes version. The provider will only perform drift detection if a configuration value is provided.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewNodeGroup registers a new resource with the given unique name, arguments, and options.
func NewNodeGroup(ctx *pulumi.Context,
	name string, args *NodeGroupArgs, opts ...pulumi.ResourceOption) (*NodeGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.NodeRoleArn == nil {
		return nil, errors.New("invalid value for required argument 'NodeRoleArn'")
	}
	if args.ScalingConfig == nil {
		return nil, errors.New("invalid value for required argument 'ScalingConfig'")
	}
	if args.SubnetIds == nil {
		return nil, errors.New("invalid value for required argument 'SubnetIds'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NodeGroup
	err := ctx.RegisterResource("aws:eks/nodeGroup:NodeGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNodeGroup gets an existing NodeGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNodeGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NodeGroupState, opts ...pulumi.ResourceOption) (*NodeGroup, error) {
	var resource NodeGroup
	err := ctx.ReadResource("aws:eks/nodeGroup:NodeGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NodeGroup resources.
type nodeGroupState struct {
	// Type of Amazon Machine Image (AMI) associated with the EKS Node Group. See the [AWS documentation](https://docs.aws.amazon.com/eks/latest/APIReference/API_Nodegroup.html#AmazonEKS-Type-Nodegroup-amiType) for valid values. This provider will only perform drift detection if a configuration value is provided.
	AmiType *string `pulumi:"amiType"`
	// Amazon Resource Name (ARN) of the EKS Node Group.
	Arn *string `pulumi:"arn"`
	// Type of capacity associated with the EKS Node Group. Valid values: `ON_DEMAND`, `SPOT`. This provider will only perform drift detection if a configuration value is provided.
	CapacityType *string `pulumi:"capacityType"`
	// Name of the EKS Cluster. Must be between 1-100 characters in length. Must begin with an alphanumeric character, and must only contain alphanumeric characters, dashes and underscores (`^[0-9A-Za-z][A-Za-z0-9\-_]+$`).
	ClusterName *string `pulumi:"clusterName"`
	// Disk size in GiB for worker nodes. Defaults to `50` for Windows, `20` all other node groups. The provider will only perform drift detection if a configuration value is provided.
	DiskSize *int `pulumi:"diskSize"`
	// Force version update if existing pods are unable to be drained due to a pod disruption budget issue.
	ForceUpdateVersion *bool `pulumi:"forceUpdateVersion"`
	// List of instance types associated with the EKS Node Group. Defaults to `["t3.medium"]`. The provider will only perform drift detection if a configuration value is provided.
	InstanceTypes []string `pulumi:"instanceTypes"`
	// Key-value map of Kubernetes labels. Only labels that are applied with the EKS API are managed by this argument. Other Kubernetes labels applied to the EKS Node Group will not be managed.
	Labels map[string]string `pulumi:"labels"`
	// Configuration block with Launch Template settings. See `launchTemplate` below for details.
	LaunchTemplate *NodeGroupLaunchTemplate `pulumi:"launchTemplate"`
	// Name of the EKS Node Group. If omitted, the provider will assign a random, unique name. Conflicts with `nodeGroupNamePrefix`. The node group name can't be longer than 63 characters. It must start with a letter or digit, but can also include hyphens and underscores for the remaining characters.
	NodeGroupName *string `pulumi:"nodeGroupName"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `nodeGroupName`.
	NodeGroupNamePrefix *string `pulumi:"nodeGroupNamePrefix"`
	// Amazon Resource Name (ARN) of the IAM Role that provides permissions for the EKS Node Group.
	NodeRoleArn *string `pulumi:"nodeRoleArn"`
	// AMI version of the EKS Node Group. Defaults to latest version for Kubernetes version.
	ReleaseVersion *string `pulumi:"releaseVersion"`
	// Configuration block with remote access settings. See `remoteAccess` below for details.
	RemoteAccess *NodeGroupRemoteAccess `pulumi:"remoteAccess"`
	// List of objects containing information about underlying resources.
	Resources []NodeGroupResource `pulumi:"resources"`
	// Configuration block with scaling settings. See `scalingConfig` below for details.
	ScalingConfig *NodeGroupScalingConfig `pulumi:"scalingConfig"`
	// Status of the EKS Node Group.
	Status *string `pulumi:"status"`
	// Identifiers of EC2 Subnets to associate with the EKS Node Group.
	//
	// The following arguments are optional:
	SubnetIds []string `pulumi:"subnetIds"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The Kubernetes taints to be applied to the nodes in the node group. Maximum of 50 taints per node group. See taint below for details.
	Taints []NodeGroupTaint `pulumi:"taints"`
	// Configuration block with update settings. See `updateConfig` below for details.
	UpdateConfig *NodeGroupUpdateConfig `pulumi:"updateConfig"`
	// Kubernetes version. Defaults to EKS Cluster Kubernetes version. The provider will only perform drift detection if a configuration value is provided.
	Version *string `pulumi:"version"`
}

type NodeGroupState struct {
	// Type of Amazon Machine Image (AMI) associated with the EKS Node Group. See the [AWS documentation](https://docs.aws.amazon.com/eks/latest/APIReference/API_Nodegroup.html#AmazonEKS-Type-Nodegroup-amiType) for valid values. This provider will only perform drift detection if a configuration value is provided.
	AmiType pulumi.StringPtrInput
	// Amazon Resource Name (ARN) of the EKS Node Group.
	Arn pulumi.StringPtrInput
	// Type of capacity associated with the EKS Node Group. Valid values: `ON_DEMAND`, `SPOT`. This provider will only perform drift detection if a configuration value is provided.
	CapacityType pulumi.StringPtrInput
	// Name of the EKS Cluster. Must be between 1-100 characters in length. Must begin with an alphanumeric character, and must only contain alphanumeric characters, dashes and underscores (`^[0-9A-Za-z][A-Za-z0-9\-_]+$`).
	ClusterName pulumi.StringPtrInput
	// Disk size in GiB for worker nodes. Defaults to `50` for Windows, `20` all other node groups. The provider will only perform drift detection if a configuration value is provided.
	DiskSize pulumi.IntPtrInput
	// Force version update if existing pods are unable to be drained due to a pod disruption budget issue.
	ForceUpdateVersion pulumi.BoolPtrInput
	// List of instance types associated with the EKS Node Group. Defaults to `["t3.medium"]`. The provider will only perform drift detection if a configuration value is provided.
	InstanceTypes pulumi.StringArrayInput
	// Key-value map of Kubernetes labels. Only labels that are applied with the EKS API are managed by this argument. Other Kubernetes labels applied to the EKS Node Group will not be managed.
	Labels pulumi.StringMapInput
	// Configuration block with Launch Template settings. See `launchTemplate` below for details.
	LaunchTemplate NodeGroupLaunchTemplatePtrInput
	// Name of the EKS Node Group. If omitted, the provider will assign a random, unique name. Conflicts with `nodeGroupNamePrefix`. The node group name can't be longer than 63 characters. It must start with a letter or digit, but can also include hyphens and underscores for the remaining characters.
	NodeGroupName pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `nodeGroupName`.
	NodeGroupNamePrefix pulumi.StringPtrInput
	// Amazon Resource Name (ARN) of the IAM Role that provides permissions for the EKS Node Group.
	NodeRoleArn pulumi.StringPtrInput
	// AMI version of the EKS Node Group. Defaults to latest version for Kubernetes version.
	ReleaseVersion pulumi.StringPtrInput
	// Configuration block with remote access settings. See `remoteAccess` below for details.
	RemoteAccess NodeGroupRemoteAccessPtrInput
	// List of objects containing information about underlying resources.
	Resources NodeGroupResourceArrayInput
	// Configuration block with scaling settings. See `scalingConfig` below for details.
	ScalingConfig NodeGroupScalingConfigPtrInput
	// Status of the EKS Node Group.
	Status pulumi.StringPtrInput
	// Identifiers of EC2 Subnets to associate with the EKS Node Group.
	//
	// The following arguments are optional:
	SubnetIds pulumi.StringArrayInput
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
	// The Kubernetes taints to be applied to the nodes in the node group. Maximum of 50 taints per node group. See taint below for details.
	Taints NodeGroupTaintArrayInput
	// Configuration block with update settings. See `updateConfig` below for details.
	UpdateConfig NodeGroupUpdateConfigPtrInput
	// Kubernetes version. Defaults to EKS Cluster Kubernetes version. The provider will only perform drift detection if a configuration value is provided.
	Version pulumi.StringPtrInput
}

func (NodeGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*nodeGroupState)(nil)).Elem()
}

type nodeGroupArgs struct {
	// Type of Amazon Machine Image (AMI) associated with the EKS Node Group. See the [AWS documentation](https://docs.aws.amazon.com/eks/latest/APIReference/API_Nodegroup.html#AmazonEKS-Type-Nodegroup-amiType) for valid values. This provider will only perform drift detection if a configuration value is provided.
	AmiType *string `pulumi:"amiType"`
	// Type of capacity associated with the EKS Node Group. Valid values: `ON_DEMAND`, `SPOT`. This provider will only perform drift detection if a configuration value is provided.
	CapacityType *string `pulumi:"capacityType"`
	// Name of the EKS Cluster. Must be between 1-100 characters in length. Must begin with an alphanumeric character, and must only contain alphanumeric characters, dashes and underscores (`^[0-9A-Za-z][A-Za-z0-9\-_]+$`).
	ClusterName string `pulumi:"clusterName"`
	// Disk size in GiB for worker nodes. Defaults to `50` for Windows, `20` all other node groups. The provider will only perform drift detection if a configuration value is provided.
	DiskSize *int `pulumi:"diskSize"`
	// Force version update if existing pods are unable to be drained due to a pod disruption budget issue.
	ForceUpdateVersion *bool `pulumi:"forceUpdateVersion"`
	// List of instance types associated with the EKS Node Group. Defaults to `["t3.medium"]`. The provider will only perform drift detection if a configuration value is provided.
	InstanceTypes []string `pulumi:"instanceTypes"`
	// Key-value map of Kubernetes labels. Only labels that are applied with the EKS API are managed by this argument. Other Kubernetes labels applied to the EKS Node Group will not be managed.
	Labels map[string]string `pulumi:"labels"`
	// Configuration block with Launch Template settings. See `launchTemplate` below for details.
	LaunchTemplate *NodeGroupLaunchTemplate `pulumi:"launchTemplate"`
	// Name of the EKS Node Group. If omitted, the provider will assign a random, unique name. Conflicts with `nodeGroupNamePrefix`. The node group name can't be longer than 63 characters. It must start with a letter or digit, but can also include hyphens and underscores for the remaining characters.
	NodeGroupName *string `pulumi:"nodeGroupName"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `nodeGroupName`.
	NodeGroupNamePrefix *string `pulumi:"nodeGroupNamePrefix"`
	// Amazon Resource Name (ARN) of the IAM Role that provides permissions for the EKS Node Group.
	NodeRoleArn string `pulumi:"nodeRoleArn"`
	// AMI version of the EKS Node Group. Defaults to latest version for Kubernetes version.
	ReleaseVersion *string `pulumi:"releaseVersion"`
	// Configuration block with remote access settings. See `remoteAccess` below for details.
	RemoteAccess *NodeGroupRemoteAccess `pulumi:"remoteAccess"`
	// Configuration block with scaling settings. See `scalingConfig` below for details.
	ScalingConfig NodeGroupScalingConfig `pulumi:"scalingConfig"`
	// Identifiers of EC2 Subnets to associate with the EKS Node Group.
	//
	// The following arguments are optional:
	SubnetIds []string `pulumi:"subnetIds"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// The Kubernetes taints to be applied to the nodes in the node group. Maximum of 50 taints per node group. See taint below for details.
	Taints []NodeGroupTaint `pulumi:"taints"`
	// Configuration block with update settings. See `updateConfig` below for details.
	UpdateConfig *NodeGroupUpdateConfig `pulumi:"updateConfig"`
	// Kubernetes version. Defaults to EKS Cluster Kubernetes version. The provider will only perform drift detection if a configuration value is provided.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a NodeGroup resource.
type NodeGroupArgs struct {
	// Type of Amazon Machine Image (AMI) associated with the EKS Node Group. See the [AWS documentation](https://docs.aws.amazon.com/eks/latest/APIReference/API_Nodegroup.html#AmazonEKS-Type-Nodegroup-amiType) for valid values. This provider will only perform drift detection if a configuration value is provided.
	AmiType pulumi.StringPtrInput
	// Type of capacity associated with the EKS Node Group. Valid values: `ON_DEMAND`, `SPOT`. This provider will only perform drift detection if a configuration value is provided.
	CapacityType pulumi.StringPtrInput
	// Name of the EKS Cluster. Must be between 1-100 characters in length. Must begin with an alphanumeric character, and must only contain alphanumeric characters, dashes and underscores (`^[0-9A-Za-z][A-Za-z0-9\-_]+$`).
	ClusterName pulumi.StringInput
	// Disk size in GiB for worker nodes. Defaults to `50` for Windows, `20` all other node groups. The provider will only perform drift detection if a configuration value is provided.
	DiskSize pulumi.IntPtrInput
	// Force version update if existing pods are unable to be drained due to a pod disruption budget issue.
	ForceUpdateVersion pulumi.BoolPtrInput
	// List of instance types associated with the EKS Node Group. Defaults to `["t3.medium"]`. The provider will only perform drift detection if a configuration value is provided.
	InstanceTypes pulumi.StringArrayInput
	// Key-value map of Kubernetes labels. Only labels that are applied with the EKS API are managed by this argument. Other Kubernetes labels applied to the EKS Node Group will not be managed.
	Labels pulumi.StringMapInput
	// Configuration block with Launch Template settings. See `launchTemplate` below for details.
	LaunchTemplate NodeGroupLaunchTemplatePtrInput
	// Name of the EKS Node Group. If omitted, the provider will assign a random, unique name. Conflicts with `nodeGroupNamePrefix`. The node group name can't be longer than 63 characters. It must start with a letter or digit, but can also include hyphens and underscores for the remaining characters.
	NodeGroupName pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `nodeGroupName`.
	NodeGroupNamePrefix pulumi.StringPtrInput
	// Amazon Resource Name (ARN) of the IAM Role that provides permissions for the EKS Node Group.
	NodeRoleArn pulumi.StringInput
	// AMI version of the EKS Node Group. Defaults to latest version for Kubernetes version.
	ReleaseVersion pulumi.StringPtrInput
	// Configuration block with remote access settings. See `remoteAccess` below for details.
	RemoteAccess NodeGroupRemoteAccessPtrInput
	// Configuration block with scaling settings. See `scalingConfig` below for details.
	ScalingConfig NodeGroupScalingConfigInput
	// Identifiers of EC2 Subnets to associate with the EKS Node Group.
	//
	// The following arguments are optional:
	SubnetIds pulumi.StringArrayInput
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// The Kubernetes taints to be applied to the nodes in the node group. Maximum of 50 taints per node group. See taint below for details.
	Taints NodeGroupTaintArrayInput
	// Configuration block with update settings. See `updateConfig` below for details.
	UpdateConfig NodeGroupUpdateConfigPtrInput
	// Kubernetes version. Defaults to EKS Cluster Kubernetes version. The provider will only perform drift detection if a configuration value is provided.
	Version pulumi.StringPtrInput
}

func (NodeGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nodeGroupArgs)(nil)).Elem()
}

type NodeGroupInput interface {
	pulumi.Input

	ToNodeGroupOutput() NodeGroupOutput
	ToNodeGroupOutputWithContext(ctx context.Context) NodeGroupOutput
}

func (*NodeGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**NodeGroup)(nil)).Elem()
}

func (i *NodeGroup) ToNodeGroupOutput() NodeGroupOutput {
	return i.ToNodeGroupOutputWithContext(context.Background())
}

func (i *NodeGroup) ToNodeGroupOutputWithContext(ctx context.Context) NodeGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeGroupOutput)
}

// NodeGroupArrayInput is an input type that accepts NodeGroupArray and NodeGroupArrayOutput values.
// You can construct a concrete instance of `NodeGroupArrayInput` via:
//
//	NodeGroupArray{ NodeGroupArgs{...} }
type NodeGroupArrayInput interface {
	pulumi.Input

	ToNodeGroupArrayOutput() NodeGroupArrayOutput
	ToNodeGroupArrayOutputWithContext(context.Context) NodeGroupArrayOutput
}

type NodeGroupArray []NodeGroupInput

func (NodeGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NodeGroup)(nil)).Elem()
}

func (i NodeGroupArray) ToNodeGroupArrayOutput() NodeGroupArrayOutput {
	return i.ToNodeGroupArrayOutputWithContext(context.Background())
}

func (i NodeGroupArray) ToNodeGroupArrayOutputWithContext(ctx context.Context) NodeGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeGroupArrayOutput)
}

// NodeGroupMapInput is an input type that accepts NodeGroupMap and NodeGroupMapOutput values.
// You can construct a concrete instance of `NodeGroupMapInput` via:
//
//	NodeGroupMap{ "key": NodeGroupArgs{...} }
type NodeGroupMapInput interface {
	pulumi.Input

	ToNodeGroupMapOutput() NodeGroupMapOutput
	ToNodeGroupMapOutputWithContext(context.Context) NodeGroupMapOutput
}

type NodeGroupMap map[string]NodeGroupInput

func (NodeGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NodeGroup)(nil)).Elem()
}

func (i NodeGroupMap) ToNodeGroupMapOutput() NodeGroupMapOutput {
	return i.ToNodeGroupMapOutputWithContext(context.Background())
}

func (i NodeGroupMap) ToNodeGroupMapOutputWithContext(ctx context.Context) NodeGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeGroupMapOutput)
}

type NodeGroupOutput struct{ *pulumi.OutputState }

func (NodeGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NodeGroup)(nil)).Elem()
}

func (o NodeGroupOutput) ToNodeGroupOutput() NodeGroupOutput {
	return o
}

func (o NodeGroupOutput) ToNodeGroupOutputWithContext(ctx context.Context) NodeGroupOutput {
	return o
}

// Type of Amazon Machine Image (AMI) associated with the EKS Node Group. See the [AWS documentation](https://docs.aws.amazon.com/eks/latest/APIReference/API_Nodegroup.html#AmazonEKS-Type-Nodegroup-amiType) for valid values. This provider will only perform drift detection if a configuration value is provided.
func (o NodeGroupOutput) AmiType() pulumi.StringOutput {
	return o.ApplyT(func(v *NodeGroup) pulumi.StringOutput { return v.AmiType }).(pulumi.StringOutput)
}

// Amazon Resource Name (ARN) of the EKS Node Group.
func (o NodeGroupOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *NodeGroup) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Type of capacity associated with the EKS Node Group. Valid values: `ON_DEMAND`, `SPOT`. This provider will only perform drift detection if a configuration value is provided.
func (o NodeGroupOutput) CapacityType() pulumi.StringOutput {
	return o.ApplyT(func(v *NodeGroup) pulumi.StringOutput { return v.CapacityType }).(pulumi.StringOutput)
}

// Name of the EKS Cluster. Must be between 1-100 characters in length. Must begin with an alphanumeric character, and must only contain alphanumeric characters, dashes and underscores (`^[0-9A-Za-z][A-Za-z0-9\-_]+$`).
func (o NodeGroupOutput) ClusterName() pulumi.StringOutput {
	return o.ApplyT(func(v *NodeGroup) pulumi.StringOutput { return v.ClusterName }).(pulumi.StringOutput)
}

// Disk size in GiB for worker nodes. Defaults to `50` for Windows, `20` all other node groups. The provider will only perform drift detection if a configuration value is provided.
func (o NodeGroupOutput) DiskSize() pulumi.IntOutput {
	return o.ApplyT(func(v *NodeGroup) pulumi.IntOutput { return v.DiskSize }).(pulumi.IntOutput)
}

// Force version update if existing pods are unable to be drained due to a pod disruption budget issue.
func (o NodeGroupOutput) ForceUpdateVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NodeGroup) pulumi.BoolPtrOutput { return v.ForceUpdateVersion }).(pulumi.BoolPtrOutput)
}

// List of instance types associated with the EKS Node Group. Defaults to `["t3.medium"]`. The provider will only perform drift detection if a configuration value is provided.
func (o NodeGroupOutput) InstanceTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NodeGroup) pulumi.StringArrayOutput { return v.InstanceTypes }).(pulumi.StringArrayOutput)
}

// Key-value map of Kubernetes labels. Only labels that are applied with the EKS API are managed by this argument. Other Kubernetes labels applied to the EKS Node Group will not be managed.
func (o NodeGroupOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NodeGroup) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Configuration block with Launch Template settings. See `launchTemplate` below for details.
func (o NodeGroupOutput) LaunchTemplate() NodeGroupLaunchTemplatePtrOutput {
	return o.ApplyT(func(v *NodeGroup) NodeGroupLaunchTemplatePtrOutput { return v.LaunchTemplate }).(NodeGroupLaunchTemplatePtrOutput)
}

// Name of the EKS Node Group. If omitted, the provider will assign a random, unique name. Conflicts with `nodeGroupNamePrefix`. The node group name can't be longer than 63 characters. It must start with a letter or digit, but can also include hyphens and underscores for the remaining characters.
func (o NodeGroupOutput) NodeGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *NodeGroup) pulumi.StringOutput { return v.NodeGroupName }).(pulumi.StringOutput)
}

// Creates a unique name beginning with the specified prefix. Conflicts with `nodeGroupName`.
func (o NodeGroupOutput) NodeGroupNamePrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *NodeGroup) pulumi.StringOutput { return v.NodeGroupNamePrefix }).(pulumi.StringOutput)
}

// Amazon Resource Name (ARN) of the IAM Role that provides permissions for the EKS Node Group.
func (o NodeGroupOutput) NodeRoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *NodeGroup) pulumi.StringOutput { return v.NodeRoleArn }).(pulumi.StringOutput)
}

// AMI version of the EKS Node Group. Defaults to latest version for Kubernetes version.
func (o NodeGroupOutput) ReleaseVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *NodeGroup) pulumi.StringOutput { return v.ReleaseVersion }).(pulumi.StringOutput)
}

// Configuration block with remote access settings. See `remoteAccess` below for details.
func (o NodeGroupOutput) RemoteAccess() NodeGroupRemoteAccessPtrOutput {
	return o.ApplyT(func(v *NodeGroup) NodeGroupRemoteAccessPtrOutput { return v.RemoteAccess }).(NodeGroupRemoteAccessPtrOutput)
}

// List of objects containing information about underlying resources.
func (o NodeGroupOutput) Resources() NodeGroupResourceArrayOutput {
	return o.ApplyT(func(v *NodeGroup) NodeGroupResourceArrayOutput { return v.Resources }).(NodeGroupResourceArrayOutput)
}

// Configuration block with scaling settings. See `scalingConfig` below for details.
func (o NodeGroupOutput) ScalingConfig() NodeGroupScalingConfigOutput {
	return o.ApplyT(func(v *NodeGroup) NodeGroupScalingConfigOutput { return v.ScalingConfig }).(NodeGroupScalingConfigOutput)
}

// Status of the EKS Node Group.
func (o NodeGroupOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *NodeGroup) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Identifiers of EC2 Subnets to associate with the EKS Node Group.
//
// The following arguments are optional:
func (o NodeGroupOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NodeGroup) pulumi.StringArrayOutput { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o NodeGroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NodeGroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o NodeGroupOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NodeGroup) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// The Kubernetes taints to be applied to the nodes in the node group. Maximum of 50 taints per node group. See taint below for details.
func (o NodeGroupOutput) Taints() NodeGroupTaintArrayOutput {
	return o.ApplyT(func(v *NodeGroup) NodeGroupTaintArrayOutput { return v.Taints }).(NodeGroupTaintArrayOutput)
}

// Configuration block with update settings. See `updateConfig` below for details.
func (o NodeGroupOutput) UpdateConfig() NodeGroupUpdateConfigOutput {
	return o.ApplyT(func(v *NodeGroup) NodeGroupUpdateConfigOutput { return v.UpdateConfig }).(NodeGroupUpdateConfigOutput)
}

// Kubernetes version. Defaults to EKS Cluster Kubernetes version. The provider will only perform drift detection if a configuration value is provided.
func (o NodeGroupOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *NodeGroup) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

type NodeGroupArrayOutput struct{ *pulumi.OutputState }

func (NodeGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NodeGroup)(nil)).Elem()
}

func (o NodeGroupArrayOutput) ToNodeGroupArrayOutput() NodeGroupArrayOutput {
	return o
}

func (o NodeGroupArrayOutput) ToNodeGroupArrayOutputWithContext(ctx context.Context) NodeGroupArrayOutput {
	return o
}

func (o NodeGroupArrayOutput) Index(i pulumi.IntInput) NodeGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NodeGroup {
		return vs[0].([]*NodeGroup)[vs[1].(int)]
	}).(NodeGroupOutput)
}

type NodeGroupMapOutput struct{ *pulumi.OutputState }

func (NodeGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NodeGroup)(nil)).Elem()
}

func (o NodeGroupMapOutput) ToNodeGroupMapOutput() NodeGroupMapOutput {
	return o
}

func (o NodeGroupMapOutput) ToNodeGroupMapOutputWithContext(ctx context.Context) NodeGroupMapOutput {
	return o
}

func (o NodeGroupMapOutput) MapIndex(k pulumi.StringInput) NodeGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NodeGroup {
		return vs[0].(map[string]*NodeGroup)[vs[1].(string)]
	}).(NodeGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NodeGroupInput)(nil)).Elem(), &NodeGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*NodeGroupArrayInput)(nil)).Elem(), NodeGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NodeGroupMapInput)(nil)).Elem(), NodeGroupMap{})
	pulumi.RegisterOutputType(NodeGroupOutput{})
	pulumi.RegisterOutputType(NodeGroupArrayOutput{})
	pulumi.RegisterOutputType(NodeGroupMapOutput{})
}
