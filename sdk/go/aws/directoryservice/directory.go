// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package directoryservice

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides a Simple or Managed Microsoft directory in AWS Directory Service.
//
// ## Example Usage
// ### SimpleAD
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/directoryservice"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			main, err := ec2.NewVpc(ctx, "main", &ec2.VpcArgs{
//				CidrBlock: pulumi.String("10.0.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			foo, err := ec2.NewSubnet(ctx, "foo", &ec2.SubnetArgs{
//				VpcId:            main.ID(),
//				AvailabilityZone: pulumi.String("us-west-2a"),
//				CidrBlock:        pulumi.String("10.0.1.0/24"),
//			})
//			if err != nil {
//				return err
//			}
//			barSubnet, err := ec2.NewSubnet(ctx, "barSubnet", &ec2.SubnetArgs{
//				VpcId:            main.ID(),
//				AvailabilityZone: pulumi.String("us-west-2b"),
//				CidrBlock:        pulumi.String("10.0.2.0/24"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = directoryservice.NewDirectory(ctx, "barDirectory", &directoryservice.DirectoryArgs{
//				Name:     pulumi.String("corp.notexample.com"),
//				Password: pulumi.String("SuperSecretPassw0rd"),
//				Size:     pulumi.String("Small"),
//				VpcSettings: &directoryservice.DirectoryVpcSettingsArgs{
//					VpcId: main.ID(),
//					SubnetIds: pulumi.StringArray{
//						foo.ID(),
//						barSubnet.ID(),
//					},
//				},
//				Tags: pulumi.StringMap{
//					"Project": pulumi.String("foo"),
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
// ### Microsoft Active Directory (MicrosoftAD)
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/directoryservice"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			main, err := ec2.NewVpc(ctx, "main", &ec2.VpcArgs{
//				CidrBlock: pulumi.String("10.0.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			foo, err := ec2.NewSubnet(ctx, "foo", &ec2.SubnetArgs{
//				VpcId:            main.ID(),
//				AvailabilityZone: pulumi.String("us-west-2a"),
//				CidrBlock:        pulumi.String("10.0.1.0/24"),
//			})
//			if err != nil {
//				return err
//			}
//			barSubnet, err := ec2.NewSubnet(ctx, "barSubnet", &ec2.SubnetArgs{
//				VpcId:            main.ID(),
//				AvailabilityZone: pulumi.String("us-west-2b"),
//				CidrBlock:        pulumi.String("10.0.2.0/24"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = directoryservice.NewDirectory(ctx, "barDirectory", &directoryservice.DirectoryArgs{
//				Name:     pulumi.String("corp.notexample.com"),
//				Password: pulumi.String("SuperSecretPassw0rd"),
//				Edition:  pulumi.String("Standard"),
//				Type:     pulumi.String("MicrosoftAD"),
//				VpcSettings: &directoryservice.DirectoryVpcSettingsArgs{
//					VpcId: main.ID(),
//					SubnetIds: pulumi.StringArray{
//						foo.ID(),
//						barSubnet.ID(),
//					},
//				},
//				Tags: pulumi.StringMap{
//					"Project": pulumi.String("foo"),
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
// ### Microsoft Active Directory Connector (ADConnector)
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/directoryservice"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			main, err := ec2.NewVpc(ctx, "main", &ec2.VpcArgs{
//				CidrBlock: pulumi.String("10.0.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			foo, err := ec2.NewSubnet(ctx, "foo", &ec2.SubnetArgs{
//				VpcId:            main.ID(),
//				AvailabilityZone: pulumi.String("us-west-2a"),
//				CidrBlock:        pulumi.String("10.0.1.0/24"),
//			})
//			if err != nil {
//				return err
//			}
//			bar, err := ec2.NewSubnet(ctx, "bar", &ec2.SubnetArgs{
//				VpcId:            main.ID(),
//				AvailabilityZone: pulumi.String("us-west-2b"),
//				CidrBlock:        pulumi.String("10.0.2.0/24"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = directoryservice.NewDirectory(ctx, "connector", &directoryservice.DirectoryArgs{
//				Name:     pulumi.String("corp.notexample.com"),
//				Password: pulumi.String("SuperSecretPassw0rd"),
//				Size:     pulumi.String("Small"),
//				Type:     pulumi.String("ADConnector"),
//				ConnectSettings: &directoryservice.DirectoryConnectSettingsArgs{
//					CustomerDnsIps: pulumi.StringArray{
//						pulumi.String("A.B.C.D"),
//					},
//					CustomerUsername: pulumi.String("Admin"),
//					SubnetIds: pulumi.StringArray{
//						foo.ID(),
//						bar.ID(),
//					},
//					VpcId: main.ID(),
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
//
// ## Import
//
// Using `pulumi import`, import DirectoryService directories using the directory `id`. For example:
//
// ```sh
//
//	$ pulumi import aws:directoryservice/directory:Directory sample d-926724cf57
//
// ```
type Directory struct {
	pulumi.CustomResourceState

	// The access URL for the directory, such as `http://alias.awsapps.com`.
	AccessUrl pulumi.StringOutput `pulumi:"accessUrl"`
	// The alias for the directory (must be unique amongst all aliases in AWS). Required for `enableSso`.
	Alias pulumi.StringOutput `pulumi:"alias"`
	// Connector related information about the directory. Fields documented below.
	ConnectSettings DirectoryConnectSettingsPtrOutput `pulumi:"connectSettings"`
	// A textual description for the directory.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The number of domain controllers desired in the directory. Minimum value of `2`. Scaling of domain controllers is only supported for `MicrosoftAD` directories.
	DesiredNumberOfDomainControllers pulumi.IntOutput `pulumi:"desiredNumberOfDomainControllers"`
	// A list of IP addresses of the DNS servers for the directory or connector.
	DnsIpAddresses pulumi.StringArrayOutput `pulumi:"dnsIpAddresses"`
	// The MicrosoftAD edition (`Standard` or `Enterprise`). Defaults to `Enterprise`.
	Edition pulumi.StringOutput `pulumi:"edition"`
	// Whether to enable single-sign on for the directory. Requires `alias`. Defaults to `false`.
	EnableSso pulumi.BoolPtrOutput `pulumi:"enableSso"`
	// The fully qualified name for the directory, such as `corp.example.com`
	Name pulumi.StringOutput `pulumi:"name"`
	// The password for the directory administrator or connector user.
	Password pulumi.StringOutput `pulumi:"password"`
	// The ID of the security group created by the directory.
	SecurityGroupId pulumi.StringOutput `pulumi:"securityGroupId"`
	// The short name of the directory, such as `CORP`.
	ShortName pulumi.StringOutput `pulumi:"shortName"`
	// (For `SimpleAD` and `ADConnector` types) The size of the directory (`Small` or `Large` are accepted values). `Large` by default.
	Size pulumi.StringOutput `pulumi:"size"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The directory type (`SimpleAD`, `ADConnector` or `MicrosoftAD` are accepted values). Defaults to `SimpleAD`.
	Type pulumi.StringPtrOutput `pulumi:"type"`
	// VPC related information about the directory. Fields documented below.
	VpcSettings DirectoryVpcSettingsPtrOutput `pulumi:"vpcSettings"`
}

// NewDirectory registers a new resource with the given unique name, arguments, and options.
func NewDirectory(ctx *pulumi.Context,
	name string, args *DirectoryArgs, opts ...pulumi.ResourceOption) (*Directory, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Directory
	err := ctx.RegisterResource("aws:directoryservice/directory:Directory", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDirectory gets an existing Directory resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDirectory(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DirectoryState, opts ...pulumi.ResourceOption) (*Directory, error) {
	var resource Directory
	err := ctx.ReadResource("aws:directoryservice/directory:Directory", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Directory resources.
type directoryState struct {
	// The access URL for the directory, such as `http://alias.awsapps.com`.
	AccessUrl *string `pulumi:"accessUrl"`
	// The alias for the directory (must be unique amongst all aliases in AWS). Required for `enableSso`.
	Alias *string `pulumi:"alias"`
	// Connector related information about the directory. Fields documented below.
	ConnectSettings *DirectoryConnectSettings `pulumi:"connectSettings"`
	// A textual description for the directory.
	Description *string `pulumi:"description"`
	// The number of domain controllers desired in the directory. Minimum value of `2`. Scaling of domain controllers is only supported for `MicrosoftAD` directories.
	DesiredNumberOfDomainControllers *int `pulumi:"desiredNumberOfDomainControllers"`
	// A list of IP addresses of the DNS servers for the directory or connector.
	DnsIpAddresses []string `pulumi:"dnsIpAddresses"`
	// The MicrosoftAD edition (`Standard` or `Enterprise`). Defaults to `Enterprise`.
	Edition *string `pulumi:"edition"`
	// Whether to enable single-sign on for the directory. Requires `alias`. Defaults to `false`.
	EnableSso *bool `pulumi:"enableSso"`
	// The fully qualified name for the directory, such as `corp.example.com`
	Name *string `pulumi:"name"`
	// The password for the directory administrator or connector user.
	Password *string `pulumi:"password"`
	// The ID of the security group created by the directory.
	SecurityGroupId *string `pulumi:"securityGroupId"`
	// The short name of the directory, such as `CORP`.
	ShortName *string `pulumi:"shortName"`
	// (For `SimpleAD` and `ADConnector` types) The size of the directory (`Small` or `Large` are accepted values). `Large` by default.
	Size *string `pulumi:"size"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The directory type (`SimpleAD`, `ADConnector` or `MicrosoftAD` are accepted values). Defaults to `SimpleAD`.
	Type *string `pulumi:"type"`
	// VPC related information about the directory. Fields documented below.
	VpcSettings *DirectoryVpcSettings `pulumi:"vpcSettings"`
}

type DirectoryState struct {
	// The access URL for the directory, such as `http://alias.awsapps.com`.
	AccessUrl pulumi.StringPtrInput
	// The alias for the directory (must be unique amongst all aliases in AWS). Required for `enableSso`.
	Alias pulumi.StringPtrInput
	// Connector related information about the directory. Fields documented below.
	ConnectSettings DirectoryConnectSettingsPtrInput
	// A textual description for the directory.
	Description pulumi.StringPtrInput
	// The number of domain controllers desired in the directory. Minimum value of `2`. Scaling of domain controllers is only supported for `MicrosoftAD` directories.
	DesiredNumberOfDomainControllers pulumi.IntPtrInput
	// A list of IP addresses of the DNS servers for the directory or connector.
	DnsIpAddresses pulumi.StringArrayInput
	// The MicrosoftAD edition (`Standard` or `Enterprise`). Defaults to `Enterprise`.
	Edition pulumi.StringPtrInput
	// Whether to enable single-sign on for the directory. Requires `alias`. Defaults to `false`.
	EnableSso pulumi.BoolPtrInput
	// The fully qualified name for the directory, such as `corp.example.com`
	Name pulumi.StringPtrInput
	// The password for the directory administrator or connector user.
	Password pulumi.StringPtrInput
	// The ID of the security group created by the directory.
	SecurityGroupId pulumi.StringPtrInput
	// The short name of the directory, such as `CORP`.
	ShortName pulumi.StringPtrInput
	// (For `SimpleAD` and `ADConnector` types) The size of the directory (`Small` or `Large` are accepted values). `Large` by default.
	Size pulumi.StringPtrInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
	// The directory type (`SimpleAD`, `ADConnector` or `MicrosoftAD` are accepted values). Defaults to `SimpleAD`.
	Type pulumi.StringPtrInput
	// VPC related information about the directory. Fields documented below.
	VpcSettings DirectoryVpcSettingsPtrInput
}

func (DirectoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*directoryState)(nil)).Elem()
}

type directoryArgs struct {
	// The alias for the directory (must be unique amongst all aliases in AWS). Required for `enableSso`.
	Alias *string `pulumi:"alias"`
	// Connector related information about the directory. Fields documented below.
	ConnectSettings *DirectoryConnectSettings `pulumi:"connectSettings"`
	// A textual description for the directory.
	Description *string `pulumi:"description"`
	// The number of domain controllers desired in the directory. Minimum value of `2`. Scaling of domain controllers is only supported for `MicrosoftAD` directories.
	DesiredNumberOfDomainControllers *int `pulumi:"desiredNumberOfDomainControllers"`
	// The MicrosoftAD edition (`Standard` or `Enterprise`). Defaults to `Enterprise`.
	Edition *string `pulumi:"edition"`
	// Whether to enable single-sign on for the directory. Requires `alias`. Defaults to `false`.
	EnableSso *bool `pulumi:"enableSso"`
	// The fully qualified name for the directory, such as `corp.example.com`
	Name string `pulumi:"name"`
	// The password for the directory administrator or connector user.
	Password string `pulumi:"password"`
	// The short name of the directory, such as `CORP`.
	ShortName *string `pulumi:"shortName"`
	// (For `SimpleAD` and `ADConnector` types) The size of the directory (`Small` or `Large` are accepted values). `Large` by default.
	Size *string `pulumi:"size"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// The directory type (`SimpleAD`, `ADConnector` or `MicrosoftAD` are accepted values). Defaults to `SimpleAD`.
	Type *string `pulumi:"type"`
	// VPC related information about the directory. Fields documented below.
	VpcSettings *DirectoryVpcSettings `pulumi:"vpcSettings"`
}

// The set of arguments for constructing a Directory resource.
type DirectoryArgs struct {
	// The alias for the directory (must be unique amongst all aliases in AWS). Required for `enableSso`.
	Alias pulumi.StringPtrInput
	// Connector related information about the directory. Fields documented below.
	ConnectSettings DirectoryConnectSettingsPtrInput
	// A textual description for the directory.
	Description pulumi.StringPtrInput
	// The number of domain controllers desired in the directory. Minimum value of `2`. Scaling of domain controllers is only supported for `MicrosoftAD` directories.
	DesiredNumberOfDomainControllers pulumi.IntPtrInput
	// The MicrosoftAD edition (`Standard` or `Enterprise`). Defaults to `Enterprise`.
	Edition pulumi.StringPtrInput
	// Whether to enable single-sign on for the directory. Requires `alias`. Defaults to `false`.
	EnableSso pulumi.BoolPtrInput
	// The fully qualified name for the directory, such as `corp.example.com`
	Name pulumi.StringInput
	// The password for the directory administrator or connector user.
	Password pulumi.StringInput
	// The short name of the directory, such as `CORP`.
	ShortName pulumi.StringPtrInput
	// (For `SimpleAD` and `ADConnector` types) The size of the directory (`Small` or `Large` are accepted values). `Large` by default.
	Size pulumi.StringPtrInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// The directory type (`SimpleAD`, `ADConnector` or `MicrosoftAD` are accepted values). Defaults to `SimpleAD`.
	Type pulumi.StringPtrInput
	// VPC related information about the directory. Fields documented below.
	VpcSettings DirectoryVpcSettingsPtrInput
}

func (DirectoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*directoryArgs)(nil)).Elem()
}

type DirectoryInput interface {
	pulumi.Input

	ToDirectoryOutput() DirectoryOutput
	ToDirectoryOutputWithContext(ctx context.Context) DirectoryOutput
}

func (*Directory) ElementType() reflect.Type {
	return reflect.TypeOf((**Directory)(nil)).Elem()
}

func (i *Directory) ToDirectoryOutput() DirectoryOutput {
	return i.ToDirectoryOutputWithContext(context.Background())
}

func (i *Directory) ToDirectoryOutputWithContext(ctx context.Context) DirectoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DirectoryOutput)
}

func (i *Directory) ToOutput(ctx context.Context) pulumix.Output[*Directory] {
	return pulumix.Output[*Directory]{
		OutputState: i.ToDirectoryOutputWithContext(ctx).OutputState,
	}
}

// DirectoryArrayInput is an input type that accepts DirectoryArray and DirectoryArrayOutput values.
// You can construct a concrete instance of `DirectoryArrayInput` via:
//
//	DirectoryArray{ DirectoryArgs{...} }
type DirectoryArrayInput interface {
	pulumi.Input

	ToDirectoryArrayOutput() DirectoryArrayOutput
	ToDirectoryArrayOutputWithContext(context.Context) DirectoryArrayOutput
}

type DirectoryArray []DirectoryInput

func (DirectoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Directory)(nil)).Elem()
}

func (i DirectoryArray) ToDirectoryArrayOutput() DirectoryArrayOutput {
	return i.ToDirectoryArrayOutputWithContext(context.Background())
}

func (i DirectoryArray) ToDirectoryArrayOutputWithContext(ctx context.Context) DirectoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DirectoryArrayOutput)
}

func (i DirectoryArray) ToOutput(ctx context.Context) pulumix.Output[[]*Directory] {
	return pulumix.Output[[]*Directory]{
		OutputState: i.ToDirectoryArrayOutputWithContext(ctx).OutputState,
	}
}

// DirectoryMapInput is an input type that accepts DirectoryMap and DirectoryMapOutput values.
// You can construct a concrete instance of `DirectoryMapInput` via:
//
//	DirectoryMap{ "key": DirectoryArgs{...} }
type DirectoryMapInput interface {
	pulumi.Input

	ToDirectoryMapOutput() DirectoryMapOutput
	ToDirectoryMapOutputWithContext(context.Context) DirectoryMapOutput
}

type DirectoryMap map[string]DirectoryInput

func (DirectoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Directory)(nil)).Elem()
}

func (i DirectoryMap) ToDirectoryMapOutput() DirectoryMapOutput {
	return i.ToDirectoryMapOutputWithContext(context.Background())
}

func (i DirectoryMap) ToDirectoryMapOutputWithContext(ctx context.Context) DirectoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DirectoryMapOutput)
}

func (i DirectoryMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Directory] {
	return pulumix.Output[map[string]*Directory]{
		OutputState: i.ToDirectoryMapOutputWithContext(ctx).OutputState,
	}
}

type DirectoryOutput struct{ *pulumi.OutputState }

func (DirectoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Directory)(nil)).Elem()
}

func (o DirectoryOutput) ToDirectoryOutput() DirectoryOutput {
	return o
}

func (o DirectoryOutput) ToDirectoryOutputWithContext(ctx context.Context) DirectoryOutput {
	return o
}

func (o DirectoryOutput) ToOutput(ctx context.Context) pulumix.Output[*Directory] {
	return pulumix.Output[*Directory]{
		OutputState: o.OutputState,
	}
}

// The access URL for the directory, such as `http://alias.awsapps.com`.
func (o DirectoryOutput) AccessUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Directory) pulumi.StringOutput { return v.AccessUrl }).(pulumi.StringOutput)
}

// The alias for the directory (must be unique amongst all aliases in AWS). Required for `enableSso`.
func (o DirectoryOutput) Alias() pulumi.StringOutput {
	return o.ApplyT(func(v *Directory) pulumi.StringOutput { return v.Alias }).(pulumi.StringOutput)
}

// Connector related information about the directory. Fields documented below.
func (o DirectoryOutput) ConnectSettings() DirectoryConnectSettingsPtrOutput {
	return o.ApplyT(func(v *Directory) DirectoryConnectSettingsPtrOutput { return v.ConnectSettings }).(DirectoryConnectSettingsPtrOutput)
}

// A textual description for the directory.
func (o DirectoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Directory) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The number of domain controllers desired in the directory. Minimum value of `2`. Scaling of domain controllers is only supported for `MicrosoftAD` directories.
func (o DirectoryOutput) DesiredNumberOfDomainControllers() pulumi.IntOutput {
	return o.ApplyT(func(v *Directory) pulumi.IntOutput { return v.DesiredNumberOfDomainControllers }).(pulumi.IntOutput)
}

// A list of IP addresses of the DNS servers for the directory or connector.
func (o DirectoryOutput) DnsIpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Directory) pulumi.StringArrayOutput { return v.DnsIpAddresses }).(pulumi.StringArrayOutput)
}

// The MicrosoftAD edition (`Standard` or `Enterprise`). Defaults to `Enterprise`.
func (o DirectoryOutput) Edition() pulumi.StringOutput {
	return o.ApplyT(func(v *Directory) pulumi.StringOutput { return v.Edition }).(pulumi.StringOutput)
}

// Whether to enable single-sign on for the directory. Requires `alias`. Defaults to `false`.
func (o DirectoryOutput) EnableSso() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Directory) pulumi.BoolPtrOutput { return v.EnableSso }).(pulumi.BoolPtrOutput)
}

// The fully qualified name for the directory, such as `corp.example.com`
func (o DirectoryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Directory) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The password for the directory administrator or connector user.
func (o DirectoryOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *Directory) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// The ID of the security group created by the directory.
func (o DirectoryOutput) SecurityGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Directory) pulumi.StringOutput { return v.SecurityGroupId }).(pulumi.StringOutput)
}

// The short name of the directory, such as `CORP`.
func (o DirectoryOutput) ShortName() pulumi.StringOutput {
	return o.ApplyT(func(v *Directory) pulumi.StringOutput { return v.ShortName }).(pulumi.StringOutput)
}

// (For `SimpleAD` and `ADConnector` types) The size of the directory (`Small` or `Large` are accepted values). `Large` by default.
func (o DirectoryOutput) Size() pulumi.StringOutput {
	return o.ApplyT(func(v *Directory) pulumi.StringOutput { return v.Size }).(pulumi.StringOutput)
}

// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o DirectoryOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Directory) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o DirectoryOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Directory) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// The directory type (`SimpleAD`, `ADConnector` or `MicrosoftAD` are accepted values). Defaults to `SimpleAD`.
func (o DirectoryOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Directory) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

// VPC related information about the directory. Fields documented below.
func (o DirectoryOutput) VpcSettings() DirectoryVpcSettingsPtrOutput {
	return o.ApplyT(func(v *Directory) DirectoryVpcSettingsPtrOutput { return v.VpcSettings }).(DirectoryVpcSettingsPtrOutput)
}

type DirectoryArrayOutput struct{ *pulumi.OutputState }

func (DirectoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Directory)(nil)).Elem()
}

func (o DirectoryArrayOutput) ToDirectoryArrayOutput() DirectoryArrayOutput {
	return o
}

func (o DirectoryArrayOutput) ToDirectoryArrayOutputWithContext(ctx context.Context) DirectoryArrayOutput {
	return o
}

func (o DirectoryArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Directory] {
	return pulumix.Output[[]*Directory]{
		OutputState: o.OutputState,
	}
}

func (o DirectoryArrayOutput) Index(i pulumi.IntInput) DirectoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Directory {
		return vs[0].([]*Directory)[vs[1].(int)]
	}).(DirectoryOutput)
}

type DirectoryMapOutput struct{ *pulumi.OutputState }

func (DirectoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Directory)(nil)).Elem()
}

func (o DirectoryMapOutput) ToDirectoryMapOutput() DirectoryMapOutput {
	return o
}

func (o DirectoryMapOutput) ToDirectoryMapOutputWithContext(ctx context.Context) DirectoryMapOutput {
	return o
}

func (o DirectoryMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Directory] {
	return pulumix.Output[map[string]*Directory]{
		OutputState: o.OutputState,
	}
}

func (o DirectoryMapOutput) MapIndex(k pulumi.StringInput) DirectoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Directory {
		return vs[0].(map[string]*Directory)[vs[1].(string)]
	}).(DirectoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DirectoryInput)(nil)).Elem(), &Directory{})
	pulumi.RegisterInputType(reflect.TypeOf((*DirectoryArrayInput)(nil)).Elem(), DirectoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DirectoryMapInput)(nil)).Elem(), DirectoryMap{})
	pulumi.RegisterOutputType(DirectoryOutput{})
	pulumi.RegisterOutputType(DirectoryArrayOutput{})
	pulumi.RegisterOutputType(DirectoryMapOutput{})
}
