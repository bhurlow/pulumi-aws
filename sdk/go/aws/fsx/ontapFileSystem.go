// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fsx

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Amazon FSx for NetApp ONTAP file system.
// See the [FSx ONTAP User Guide](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/what-is-fsx-ontap.html) for more information.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/fsx"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := fsx.NewOntapFileSystem(ctx, "test", &fsx.OntapFileSystemArgs{
//				StorageCapacity: pulumi.Int(1024),
//				SubnetIds: pulumi.StringArray{
//					aws_subnet.Test1.Id,
//					aws_subnet.Test2.Id,
//				},
//				DeploymentType:     pulumi.String("MULTI_AZ_1"),
//				ThroughputCapacity: pulumi.Int(512),
//				PreferredSubnetId:  pulumi.Any(aws_subnet.Test1.Id),
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
// FSx File Systems can be imported using the `id`, e.g.,
//
// ```sh
//
//	$ pulumi import aws:fsx/ontapFileSystem:OntapFileSystem example fs-543ab12b1ca672f33
//
// ```
//
//	Certain resource arguments, like `security_group_ids`, do not have a FSx API method for reading the information after creation. If the argument is set in the the provider configuration on an imported resource, the provider will always show a difference. To workaround this behavior, either omit the argument from the provider configuration or use `ignore_changes` to hide the difference, e.g., terraform resource "aws_fsx_ontap_file_system" "example" {
//
// # ... other configuration ...
//
//	security_group_ids = [aws_security_group.example.id]
//
// # There is no FSx API for reading security_group_ids
//
//	lifecycle {
//
//	ignore_changes = [security_group_ids]
//
//	} }
type OntapFileSystem struct {
	pulumi.CustomResourceState

	// Amazon Resource Name of the file system.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The number of days to retain automatic backups. Setting this to 0 disables automatic backups. You can retain automatic backups for a maximum of 90 days.
	AutomaticBackupRetentionDays pulumi.IntPtrOutput `pulumi:"automaticBackupRetentionDays"`
	// A recurring daily time, in the format HH:MM. HH is the zero-padded hour of the day (0-23), and MM is the zero-padded minute of the hour. For example, 05:00 specifies 5 AM daily. Requires `automaticBackupRetentionDays` to be set.
	DailyAutomaticBackupStartTime pulumi.StringOutput `pulumi:"dailyAutomaticBackupStartTime"`
	// The filesystem deployment type. Supports `MULTI_AZ_1` and `SINGLE_AZ_1`.
	DeploymentType pulumi.StringOutput `pulumi:"deploymentType"`
	// The SSD IOPS configuration for the Amazon FSx for NetApp ONTAP file system. See Disk Iops Configuration Below.
	DiskIopsConfiguration OntapFileSystemDiskIopsConfigurationOutput `pulumi:"diskIopsConfiguration"`
	// The Domain Name Service (DNS) name for the file system. You can mount your file system using its DNS name.
	DnsName pulumi.StringOutput `pulumi:"dnsName"`
	// Specifies the IP address range in which the endpoints to access your file system will be created. By default, Amazon FSx selects an unused IP address range for you from the 198.19.* range.
	EndpointIpAddressRange pulumi.StringOutput `pulumi:"endpointIpAddressRange"`
	// The endpoints that are used to access data or to manage the file system using the NetApp ONTAP CLI, REST API, or NetApp SnapMirror. See Endpoints below.
	Endpoints OntapFileSystemEndpointArrayOutput `pulumi:"endpoints"`
	// The ONTAP administrative password for the fsxadmin user that you can use to administer your file system using the ONTAP CLI and REST API.
	FsxAdminPassword pulumi.StringPtrOutput `pulumi:"fsxAdminPassword"`
	// ARN for the KMS Key to encrypt the file system at rest, Defaults to an AWS managed KMS Key.
	KmsKeyId pulumi.StringOutput `pulumi:"kmsKeyId"`
	// Set of Elastic Network Interface identifiers from which the file system is accessible The first network interface returned is the primary network interface.
	NetworkInterfaceIds pulumi.StringArrayOutput `pulumi:"networkInterfaceIds"`
	// AWS account identifier that created the file system.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// The ID for a subnet. A subnet is a range of IP addresses in your virtual private cloud (VPC).
	PreferredSubnetId pulumi.StringOutput `pulumi:"preferredSubnetId"`
	// Specifies the VPC route tables in which your file system's endpoints will be created. You should specify all VPC route tables associated with the subnets in which your clients are located. By default, Amazon FSx selects your VPC's default route table.
	RouteTableIds pulumi.StringArrayOutput `pulumi:"routeTableIds"`
	// A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// The storage capacity (GiB) of the file system. Valid values between `1024` and `196608`.
	StorageCapacity pulumi.IntPtrOutput `pulumi:"storageCapacity"`
	// The filesystem storage type. defaults to `SSD`.
	StorageType pulumi.StringPtrOutput `pulumi:"storageType"`
	// A list of IDs for the subnets that the file system will be accessible from. Upto 2 subnets can be provided.
	SubnetIds pulumi.StringArrayOutput `pulumi:"subnetIds"`
	// A map of tags to assign to the file system. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Sets the throughput capacity (in MBps) for the file system that you're creating. Valid values are `128`, `256`, `512`, `1024`, and `2048`.
	ThroughputCapacity pulumi.IntOutput `pulumi:"throughputCapacity"`
	// Identifier of the Virtual Private Cloud for the file system.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// The preferred start time (in `d:HH:MM` format) to perform weekly maintenance, in the UTC time zone.
	WeeklyMaintenanceStartTime pulumi.StringOutput `pulumi:"weeklyMaintenanceStartTime"`
}

// NewOntapFileSystem registers a new resource with the given unique name, arguments, and options.
func NewOntapFileSystem(ctx *pulumi.Context,
	name string, args *OntapFileSystemArgs, opts ...pulumi.ResourceOption) (*OntapFileSystem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeploymentType == nil {
		return nil, errors.New("invalid value for required argument 'DeploymentType'")
	}
	if args.PreferredSubnetId == nil {
		return nil, errors.New("invalid value for required argument 'PreferredSubnetId'")
	}
	if args.SubnetIds == nil {
		return nil, errors.New("invalid value for required argument 'SubnetIds'")
	}
	if args.ThroughputCapacity == nil {
		return nil, errors.New("invalid value for required argument 'ThroughputCapacity'")
	}
	if args.FsxAdminPassword != nil {
		args.FsxAdminPassword = pulumi.ToSecret(args.FsxAdminPassword).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"fsxAdminPassword",
	})
	opts = append(opts, secrets)
	var resource OntapFileSystem
	err := ctx.RegisterResource("aws:fsx/ontapFileSystem:OntapFileSystem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOntapFileSystem gets an existing OntapFileSystem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOntapFileSystem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OntapFileSystemState, opts ...pulumi.ResourceOption) (*OntapFileSystem, error) {
	var resource OntapFileSystem
	err := ctx.ReadResource("aws:fsx/ontapFileSystem:OntapFileSystem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OntapFileSystem resources.
type ontapFileSystemState struct {
	// Amazon Resource Name of the file system.
	Arn *string `pulumi:"arn"`
	// The number of days to retain automatic backups. Setting this to 0 disables automatic backups. You can retain automatic backups for a maximum of 90 days.
	AutomaticBackupRetentionDays *int `pulumi:"automaticBackupRetentionDays"`
	// A recurring daily time, in the format HH:MM. HH is the zero-padded hour of the day (0-23), and MM is the zero-padded minute of the hour. For example, 05:00 specifies 5 AM daily. Requires `automaticBackupRetentionDays` to be set.
	DailyAutomaticBackupStartTime *string `pulumi:"dailyAutomaticBackupStartTime"`
	// The filesystem deployment type. Supports `MULTI_AZ_1` and `SINGLE_AZ_1`.
	DeploymentType *string `pulumi:"deploymentType"`
	// The SSD IOPS configuration for the Amazon FSx for NetApp ONTAP file system. See Disk Iops Configuration Below.
	DiskIopsConfiguration *OntapFileSystemDiskIopsConfiguration `pulumi:"diskIopsConfiguration"`
	// The Domain Name Service (DNS) name for the file system. You can mount your file system using its DNS name.
	DnsName *string `pulumi:"dnsName"`
	// Specifies the IP address range in which the endpoints to access your file system will be created. By default, Amazon FSx selects an unused IP address range for you from the 198.19.* range.
	EndpointIpAddressRange *string `pulumi:"endpointIpAddressRange"`
	// The endpoints that are used to access data or to manage the file system using the NetApp ONTAP CLI, REST API, or NetApp SnapMirror. See Endpoints below.
	Endpoints []OntapFileSystemEndpoint `pulumi:"endpoints"`
	// The ONTAP administrative password for the fsxadmin user that you can use to administer your file system using the ONTAP CLI and REST API.
	FsxAdminPassword *string `pulumi:"fsxAdminPassword"`
	// ARN for the KMS Key to encrypt the file system at rest, Defaults to an AWS managed KMS Key.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// Set of Elastic Network Interface identifiers from which the file system is accessible The first network interface returned is the primary network interface.
	NetworkInterfaceIds []string `pulumi:"networkInterfaceIds"`
	// AWS account identifier that created the file system.
	OwnerId *string `pulumi:"ownerId"`
	// The ID for a subnet. A subnet is a range of IP addresses in your virtual private cloud (VPC).
	PreferredSubnetId *string `pulumi:"preferredSubnetId"`
	// Specifies the VPC route tables in which your file system's endpoints will be created. You should specify all VPC route tables associated with the subnets in which your clients are located. By default, Amazon FSx selects your VPC's default route table.
	RouteTableIds []string `pulumi:"routeTableIds"`
	// A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// The storage capacity (GiB) of the file system. Valid values between `1024` and `196608`.
	StorageCapacity *int `pulumi:"storageCapacity"`
	// The filesystem storage type. defaults to `SSD`.
	StorageType *string `pulumi:"storageType"`
	// A list of IDs for the subnets that the file system will be accessible from. Upto 2 subnets can be provided.
	SubnetIds []string `pulumi:"subnetIds"`
	// A map of tags to assign to the file system. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Sets the throughput capacity (in MBps) for the file system that you're creating. Valid values are `128`, `256`, `512`, `1024`, and `2048`.
	ThroughputCapacity *int `pulumi:"throughputCapacity"`
	// Identifier of the Virtual Private Cloud for the file system.
	VpcId *string `pulumi:"vpcId"`
	// The preferred start time (in `d:HH:MM` format) to perform weekly maintenance, in the UTC time zone.
	WeeklyMaintenanceStartTime *string `pulumi:"weeklyMaintenanceStartTime"`
}

type OntapFileSystemState struct {
	// Amazon Resource Name of the file system.
	Arn pulumi.StringPtrInput
	// The number of days to retain automatic backups. Setting this to 0 disables automatic backups. You can retain automatic backups for a maximum of 90 days.
	AutomaticBackupRetentionDays pulumi.IntPtrInput
	// A recurring daily time, in the format HH:MM. HH is the zero-padded hour of the day (0-23), and MM is the zero-padded minute of the hour. For example, 05:00 specifies 5 AM daily. Requires `automaticBackupRetentionDays` to be set.
	DailyAutomaticBackupStartTime pulumi.StringPtrInput
	// The filesystem deployment type. Supports `MULTI_AZ_1` and `SINGLE_AZ_1`.
	DeploymentType pulumi.StringPtrInput
	// The SSD IOPS configuration for the Amazon FSx for NetApp ONTAP file system. See Disk Iops Configuration Below.
	DiskIopsConfiguration OntapFileSystemDiskIopsConfigurationPtrInput
	// The Domain Name Service (DNS) name for the file system. You can mount your file system using its DNS name.
	DnsName pulumi.StringPtrInput
	// Specifies the IP address range in which the endpoints to access your file system will be created. By default, Amazon FSx selects an unused IP address range for you from the 198.19.* range.
	EndpointIpAddressRange pulumi.StringPtrInput
	// The endpoints that are used to access data or to manage the file system using the NetApp ONTAP CLI, REST API, or NetApp SnapMirror. See Endpoints below.
	Endpoints OntapFileSystemEndpointArrayInput
	// The ONTAP administrative password for the fsxadmin user that you can use to administer your file system using the ONTAP CLI and REST API.
	FsxAdminPassword pulumi.StringPtrInput
	// ARN for the KMS Key to encrypt the file system at rest, Defaults to an AWS managed KMS Key.
	KmsKeyId pulumi.StringPtrInput
	// Set of Elastic Network Interface identifiers from which the file system is accessible The first network interface returned is the primary network interface.
	NetworkInterfaceIds pulumi.StringArrayInput
	// AWS account identifier that created the file system.
	OwnerId pulumi.StringPtrInput
	// The ID for a subnet. A subnet is a range of IP addresses in your virtual private cloud (VPC).
	PreferredSubnetId pulumi.StringPtrInput
	// Specifies the VPC route tables in which your file system's endpoints will be created. You should specify all VPC route tables associated with the subnets in which your clients are located. By default, Amazon FSx selects your VPC's default route table.
	RouteTableIds pulumi.StringArrayInput
	// A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
	SecurityGroupIds pulumi.StringArrayInput
	// The storage capacity (GiB) of the file system. Valid values between `1024` and `196608`.
	StorageCapacity pulumi.IntPtrInput
	// The filesystem storage type. defaults to `SSD`.
	StorageType pulumi.StringPtrInput
	// A list of IDs for the subnets that the file system will be accessible from. Upto 2 subnets can be provided.
	SubnetIds pulumi.StringArrayInput
	// A map of tags to assign to the file system. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// Sets the throughput capacity (in MBps) for the file system that you're creating. Valid values are `128`, `256`, `512`, `1024`, and `2048`.
	ThroughputCapacity pulumi.IntPtrInput
	// Identifier of the Virtual Private Cloud for the file system.
	VpcId pulumi.StringPtrInput
	// The preferred start time (in `d:HH:MM` format) to perform weekly maintenance, in the UTC time zone.
	WeeklyMaintenanceStartTime pulumi.StringPtrInput
}

func (OntapFileSystemState) ElementType() reflect.Type {
	return reflect.TypeOf((*ontapFileSystemState)(nil)).Elem()
}

type ontapFileSystemArgs struct {
	// The number of days to retain automatic backups. Setting this to 0 disables automatic backups. You can retain automatic backups for a maximum of 90 days.
	AutomaticBackupRetentionDays *int `pulumi:"automaticBackupRetentionDays"`
	// A recurring daily time, in the format HH:MM. HH is the zero-padded hour of the day (0-23), and MM is the zero-padded minute of the hour. For example, 05:00 specifies 5 AM daily. Requires `automaticBackupRetentionDays` to be set.
	DailyAutomaticBackupStartTime *string `pulumi:"dailyAutomaticBackupStartTime"`
	// The filesystem deployment type. Supports `MULTI_AZ_1` and `SINGLE_AZ_1`.
	DeploymentType string `pulumi:"deploymentType"`
	// The SSD IOPS configuration for the Amazon FSx for NetApp ONTAP file system. See Disk Iops Configuration Below.
	DiskIopsConfiguration *OntapFileSystemDiskIopsConfiguration `pulumi:"diskIopsConfiguration"`
	// Specifies the IP address range in which the endpoints to access your file system will be created. By default, Amazon FSx selects an unused IP address range for you from the 198.19.* range.
	EndpointIpAddressRange *string `pulumi:"endpointIpAddressRange"`
	// The ONTAP administrative password for the fsxadmin user that you can use to administer your file system using the ONTAP CLI and REST API.
	FsxAdminPassword *string `pulumi:"fsxAdminPassword"`
	// ARN for the KMS Key to encrypt the file system at rest, Defaults to an AWS managed KMS Key.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// The ID for a subnet. A subnet is a range of IP addresses in your virtual private cloud (VPC).
	PreferredSubnetId string `pulumi:"preferredSubnetId"`
	// Specifies the VPC route tables in which your file system's endpoints will be created. You should specify all VPC route tables associated with the subnets in which your clients are located. By default, Amazon FSx selects your VPC's default route table.
	RouteTableIds []string `pulumi:"routeTableIds"`
	// A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// The storage capacity (GiB) of the file system. Valid values between `1024` and `196608`.
	StorageCapacity *int `pulumi:"storageCapacity"`
	// The filesystem storage type. defaults to `SSD`.
	StorageType *string `pulumi:"storageType"`
	// A list of IDs for the subnets that the file system will be accessible from. Upto 2 subnets can be provided.
	SubnetIds []string `pulumi:"subnetIds"`
	// A map of tags to assign to the file system. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Sets the throughput capacity (in MBps) for the file system that you're creating. Valid values are `128`, `256`, `512`, `1024`, and `2048`.
	ThroughputCapacity int `pulumi:"throughputCapacity"`
	// The preferred start time (in `d:HH:MM` format) to perform weekly maintenance, in the UTC time zone.
	WeeklyMaintenanceStartTime *string `pulumi:"weeklyMaintenanceStartTime"`
}

// The set of arguments for constructing a OntapFileSystem resource.
type OntapFileSystemArgs struct {
	// The number of days to retain automatic backups. Setting this to 0 disables automatic backups. You can retain automatic backups for a maximum of 90 days.
	AutomaticBackupRetentionDays pulumi.IntPtrInput
	// A recurring daily time, in the format HH:MM. HH is the zero-padded hour of the day (0-23), and MM is the zero-padded minute of the hour. For example, 05:00 specifies 5 AM daily. Requires `automaticBackupRetentionDays` to be set.
	DailyAutomaticBackupStartTime pulumi.StringPtrInput
	// The filesystem deployment type. Supports `MULTI_AZ_1` and `SINGLE_AZ_1`.
	DeploymentType pulumi.StringInput
	// The SSD IOPS configuration for the Amazon FSx for NetApp ONTAP file system. See Disk Iops Configuration Below.
	DiskIopsConfiguration OntapFileSystemDiskIopsConfigurationPtrInput
	// Specifies the IP address range in which the endpoints to access your file system will be created. By default, Amazon FSx selects an unused IP address range for you from the 198.19.* range.
	EndpointIpAddressRange pulumi.StringPtrInput
	// The ONTAP administrative password for the fsxadmin user that you can use to administer your file system using the ONTAP CLI and REST API.
	FsxAdminPassword pulumi.StringPtrInput
	// ARN for the KMS Key to encrypt the file system at rest, Defaults to an AWS managed KMS Key.
	KmsKeyId pulumi.StringPtrInput
	// The ID for a subnet. A subnet is a range of IP addresses in your virtual private cloud (VPC).
	PreferredSubnetId pulumi.StringInput
	// Specifies the VPC route tables in which your file system's endpoints will be created. You should specify all VPC route tables associated with the subnets in which your clients are located. By default, Amazon FSx selects your VPC's default route table.
	RouteTableIds pulumi.StringArrayInput
	// A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
	SecurityGroupIds pulumi.StringArrayInput
	// The storage capacity (GiB) of the file system. Valid values between `1024` and `196608`.
	StorageCapacity pulumi.IntPtrInput
	// The filesystem storage type. defaults to `SSD`.
	StorageType pulumi.StringPtrInput
	// A list of IDs for the subnets that the file system will be accessible from. Upto 2 subnets can be provided.
	SubnetIds pulumi.StringArrayInput
	// A map of tags to assign to the file system. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Sets the throughput capacity (in MBps) for the file system that you're creating. Valid values are `128`, `256`, `512`, `1024`, and `2048`.
	ThroughputCapacity pulumi.IntInput
	// The preferred start time (in `d:HH:MM` format) to perform weekly maintenance, in the UTC time zone.
	WeeklyMaintenanceStartTime pulumi.StringPtrInput
}

func (OntapFileSystemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ontapFileSystemArgs)(nil)).Elem()
}

type OntapFileSystemInput interface {
	pulumi.Input

	ToOntapFileSystemOutput() OntapFileSystemOutput
	ToOntapFileSystemOutputWithContext(ctx context.Context) OntapFileSystemOutput
}

func (*OntapFileSystem) ElementType() reflect.Type {
	return reflect.TypeOf((**OntapFileSystem)(nil)).Elem()
}

func (i *OntapFileSystem) ToOntapFileSystemOutput() OntapFileSystemOutput {
	return i.ToOntapFileSystemOutputWithContext(context.Background())
}

func (i *OntapFileSystem) ToOntapFileSystemOutputWithContext(ctx context.Context) OntapFileSystemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OntapFileSystemOutput)
}

// OntapFileSystemArrayInput is an input type that accepts OntapFileSystemArray and OntapFileSystemArrayOutput values.
// You can construct a concrete instance of `OntapFileSystemArrayInput` via:
//
//	OntapFileSystemArray{ OntapFileSystemArgs{...} }
type OntapFileSystemArrayInput interface {
	pulumi.Input

	ToOntapFileSystemArrayOutput() OntapFileSystemArrayOutput
	ToOntapFileSystemArrayOutputWithContext(context.Context) OntapFileSystemArrayOutput
}

type OntapFileSystemArray []OntapFileSystemInput

func (OntapFileSystemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OntapFileSystem)(nil)).Elem()
}

func (i OntapFileSystemArray) ToOntapFileSystemArrayOutput() OntapFileSystemArrayOutput {
	return i.ToOntapFileSystemArrayOutputWithContext(context.Background())
}

func (i OntapFileSystemArray) ToOntapFileSystemArrayOutputWithContext(ctx context.Context) OntapFileSystemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OntapFileSystemArrayOutput)
}

// OntapFileSystemMapInput is an input type that accepts OntapFileSystemMap and OntapFileSystemMapOutput values.
// You can construct a concrete instance of `OntapFileSystemMapInput` via:
//
//	OntapFileSystemMap{ "key": OntapFileSystemArgs{...} }
type OntapFileSystemMapInput interface {
	pulumi.Input

	ToOntapFileSystemMapOutput() OntapFileSystemMapOutput
	ToOntapFileSystemMapOutputWithContext(context.Context) OntapFileSystemMapOutput
}

type OntapFileSystemMap map[string]OntapFileSystemInput

func (OntapFileSystemMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OntapFileSystem)(nil)).Elem()
}

func (i OntapFileSystemMap) ToOntapFileSystemMapOutput() OntapFileSystemMapOutput {
	return i.ToOntapFileSystemMapOutputWithContext(context.Background())
}

func (i OntapFileSystemMap) ToOntapFileSystemMapOutputWithContext(ctx context.Context) OntapFileSystemMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OntapFileSystemMapOutput)
}

type OntapFileSystemOutput struct{ *pulumi.OutputState }

func (OntapFileSystemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OntapFileSystem)(nil)).Elem()
}

func (o OntapFileSystemOutput) ToOntapFileSystemOutput() OntapFileSystemOutput {
	return o
}

func (o OntapFileSystemOutput) ToOntapFileSystemOutputWithContext(ctx context.Context) OntapFileSystemOutput {
	return o
}

// Amazon Resource Name of the file system.
func (o OntapFileSystemOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *OntapFileSystem) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The number of days to retain automatic backups. Setting this to 0 disables automatic backups. You can retain automatic backups for a maximum of 90 days.
func (o OntapFileSystemOutput) AutomaticBackupRetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OntapFileSystem) pulumi.IntPtrOutput { return v.AutomaticBackupRetentionDays }).(pulumi.IntPtrOutput)
}

// A recurring daily time, in the format HH:MM. HH is the zero-padded hour of the day (0-23), and MM is the zero-padded minute of the hour. For example, 05:00 specifies 5 AM daily. Requires `automaticBackupRetentionDays` to be set.
func (o OntapFileSystemOutput) DailyAutomaticBackupStartTime() pulumi.StringOutput {
	return o.ApplyT(func(v *OntapFileSystem) pulumi.StringOutput { return v.DailyAutomaticBackupStartTime }).(pulumi.StringOutput)
}

// The filesystem deployment type. Supports `MULTI_AZ_1` and `SINGLE_AZ_1`.
func (o OntapFileSystemOutput) DeploymentType() pulumi.StringOutput {
	return o.ApplyT(func(v *OntapFileSystem) pulumi.StringOutput { return v.DeploymentType }).(pulumi.StringOutput)
}

// The SSD IOPS configuration for the Amazon FSx for NetApp ONTAP file system. See Disk Iops Configuration Below.
func (o OntapFileSystemOutput) DiskIopsConfiguration() OntapFileSystemDiskIopsConfigurationOutput {
	return o.ApplyT(func(v *OntapFileSystem) OntapFileSystemDiskIopsConfigurationOutput { return v.DiskIopsConfiguration }).(OntapFileSystemDiskIopsConfigurationOutput)
}

// The Domain Name Service (DNS) name for the file system. You can mount your file system using its DNS name.
func (o OntapFileSystemOutput) DnsName() pulumi.StringOutput {
	return o.ApplyT(func(v *OntapFileSystem) pulumi.StringOutput { return v.DnsName }).(pulumi.StringOutput)
}

// Specifies the IP address range in which the endpoints to access your file system will be created. By default, Amazon FSx selects an unused IP address range for you from the 198.19.* range.
func (o OntapFileSystemOutput) EndpointIpAddressRange() pulumi.StringOutput {
	return o.ApplyT(func(v *OntapFileSystem) pulumi.StringOutput { return v.EndpointIpAddressRange }).(pulumi.StringOutput)
}

// The endpoints that are used to access data or to manage the file system using the NetApp ONTAP CLI, REST API, or NetApp SnapMirror. See Endpoints below.
func (o OntapFileSystemOutput) Endpoints() OntapFileSystemEndpointArrayOutput {
	return o.ApplyT(func(v *OntapFileSystem) OntapFileSystemEndpointArrayOutput { return v.Endpoints }).(OntapFileSystemEndpointArrayOutput)
}

// The ONTAP administrative password for the fsxadmin user that you can use to administer your file system using the ONTAP CLI and REST API.
func (o OntapFileSystemOutput) FsxAdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OntapFileSystem) pulumi.StringPtrOutput { return v.FsxAdminPassword }).(pulumi.StringPtrOutput)
}

// ARN for the KMS Key to encrypt the file system at rest, Defaults to an AWS managed KMS Key.
func (o OntapFileSystemOutput) KmsKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *OntapFileSystem) pulumi.StringOutput { return v.KmsKeyId }).(pulumi.StringOutput)
}

// Set of Elastic Network Interface identifiers from which the file system is accessible The first network interface returned is the primary network interface.
func (o OntapFileSystemOutput) NetworkInterfaceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OntapFileSystem) pulumi.StringArrayOutput { return v.NetworkInterfaceIds }).(pulumi.StringArrayOutput)
}

// AWS account identifier that created the file system.
func (o OntapFileSystemOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *OntapFileSystem) pulumi.StringOutput { return v.OwnerId }).(pulumi.StringOutput)
}

// The ID for a subnet. A subnet is a range of IP addresses in your virtual private cloud (VPC).
func (o OntapFileSystemOutput) PreferredSubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *OntapFileSystem) pulumi.StringOutput { return v.PreferredSubnetId }).(pulumi.StringOutput)
}

// Specifies the VPC route tables in which your file system's endpoints will be created. You should specify all VPC route tables associated with the subnets in which your clients are located. By default, Amazon FSx selects your VPC's default route table.
func (o OntapFileSystemOutput) RouteTableIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OntapFileSystem) pulumi.StringArrayOutput { return v.RouteTableIds }).(pulumi.StringArrayOutput)
}

// A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
func (o OntapFileSystemOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OntapFileSystem) pulumi.StringArrayOutput { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// The storage capacity (GiB) of the file system. Valid values between `1024` and `196608`.
func (o OntapFileSystemOutput) StorageCapacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OntapFileSystem) pulumi.IntPtrOutput { return v.StorageCapacity }).(pulumi.IntPtrOutput)
}

// The filesystem storage type. defaults to `SSD`.
func (o OntapFileSystemOutput) StorageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OntapFileSystem) pulumi.StringPtrOutput { return v.StorageType }).(pulumi.StringPtrOutput)
}

// A list of IDs for the subnets that the file system will be accessible from. Upto 2 subnets can be provided.
func (o OntapFileSystemOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OntapFileSystem) pulumi.StringArrayOutput { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

// A map of tags to assign to the file system. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o OntapFileSystemOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *OntapFileSystem) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o OntapFileSystemOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *OntapFileSystem) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// Sets the throughput capacity (in MBps) for the file system that you're creating. Valid values are `128`, `256`, `512`, `1024`, and `2048`.
func (o OntapFileSystemOutput) ThroughputCapacity() pulumi.IntOutput {
	return o.ApplyT(func(v *OntapFileSystem) pulumi.IntOutput { return v.ThroughputCapacity }).(pulumi.IntOutput)
}

// Identifier of the Virtual Private Cloud for the file system.
func (o OntapFileSystemOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *OntapFileSystem) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

// The preferred start time (in `d:HH:MM` format) to perform weekly maintenance, in the UTC time zone.
func (o OntapFileSystemOutput) WeeklyMaintenanceStartTime() pulumi.StringOutput {
	return o.ApplyT(func(v *OntapFileSystem) pulumi.StringOutput { return v.WeeklyMaintenanceStartTime }).(pulumi.StringOutput)
}

type OntapFileSystemArrayOutput struct{ *pulumi.OutputState }

func (OntapFileSystemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OntapFileSystem)(nil)).Elem()
}

func (o OntapFileSystemArrayOutput) ToOntapFileSystemArrayOutput() OntapFileSystemArrayOutput {
	return o
}

func (o OntapFileSystemArrayOutput) ToOntapFileSystemArrayOutputWithContext(ctx context.Context) OntapFileSystemArrayOutput {
	return o
}

func (o OntapFileSystemArrayOutput) Index(i pulumi.IntInput) OntapFileSystemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OntapFileSystem {
		return vs[0].([]*OntapFileSystem)[vs[1].(int)]
	}).(OntapFileSystemOutput)
}

type OntapFileSystemMapOutput struct{ *pulumi.OutputState }

func (OntapFileSystemMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OntapFileSystem)(nil)).Elem()
}

func (o OntapFileSystemMapOutput) ToOntapFileSystemMapOutput() OntapFileSystemMapOutput {
	return o
}

func (o OntapFileSystemMapOutput) ToOntapFileSystemMapOutputWithContext(ctx context.Context) OntapFileSystemMapOutput {
	return o
}

func (o OntapFileSystemMapOutput) MapIndex(k pulumi.StringInput) OntapFileSystemOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OntapFileSystem {
		return vs[0].(map[string]*OntapFileSystem)[vs[1].(string)]
	}).(OntapFileSystemOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OntapFileSystemInput)(nil)).Elem(), &OntapFileSystem{})
	pulumi.RegisterInputType(reflect.TypeOf((*OntapFileSystemArrayInput)(nil)).Elem(), OntapFileSystemArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OntapFileSystemMapInput)(nil)).Elem(), OntapFileSystemMap{})
	pulumi.RegisterOutputType(OntapFileSystemOutput{})
	pulumi.RegisterOutputType(OntapFileSystemArrayOutput{})
	pulumi.RegisterOutputType(OntapFileSystemMapOutput{})
}
