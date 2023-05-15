// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datasync

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Object Storage Location within AWS DataSync.
//
// > **NOTE:** The DataSync Agents must be available before creating this resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/datasync"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := datasync.NewLocationObjectStorage(ctx, "example", &datasync.LocationObjectStorageArgs{
//				AgentArns: pulumi.StringArray{
//					aws_datasync_agent.Example.Arn,
//				},
//				ServerHostname: pulumi.String("example"),
//				BucketName:     pulumi.String("example"),
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
// `aws_datasync_location_object_storage` can be imported by using the Amazon Resource Name (ARN), e.g.,
//
// ```sh
//
//	$ pulumi import aws:datasync/locationObjectStorage:LocationObjectStorage example arn:aws:datasync:us-east-1:123456789012:location/loc-12345678901234567
//
// ```
type LocationObjectStorage struct {
	pulumi.CustomResourceState

	// The access key is used if credentials are required to access the self-managed object storage server. If your object storage requires a user name and password to authenticate, use `accessKey` and `secretKey` to provide the user name and password, respectively.
	AccessKey pulumi.StringPtrOutput `pulumi:"accessKey"`
	// A list of DataSync Agent ARNs with which this location will be associated.
	AgentArns pulumi.StringArrayOutput `pulumi:"agentArns"`
	// Amazon Resource Name (ARN) of the DataSync Location.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The bucket on the self-managed object storage server that is used to read data from.
	BucketName pulumi.StringOutput `pulumi:"bucketName"`
	// The secret key is used if credentials are required to access the self-managed object storage server. If your object storage requires a user name and password to authenticate, use `accessKey` and `secretKey` to provide the user name and password, respectively.
	SecretKey pulumi.StringPtrOutput `pulumi:"secretKey"`
	// Specifies a certificate to authenticate with an object storage system that uses a private or self-signed certificate authority (CA). You must specify a Base64-encoded .pem file (for example, file:///home/user/.ssh/storage_sys_certificate.pem). The certificate can be up to 32768 bytes (before Base64 encoding).
	ServerCertificate pulumi.StringPtrOutput `pulumi:"serverCertificate"`
	// The name of the self-managed object storage server. This value is the IP address or Domain Name Service (DNS) name of the object storage server. An agent uses this host name to mount the object storage server in a network.
	ServerHostname pulumi.StringOutput `pulumi:"serverHostname"`
	// The port that your self-managed object storage server accepts inbound network traffic on. The server port is set by default to TCP 80 (`HTTP`) or TCP 443 (`HTTPS`). You can specify a custom port if your self-managed object storage server requires one.
	ServerPort pulumi.IntPtrOutput `pulumi:"serverPort"`
	// The protocol that the object storage server uses to communicate. Valid values are `HTTP` or `HTTPS`.
	ServerProtocol pulumi.StringPtrOutput `pulumi:"serverProtocol"`
	// A subdirectory in the HDFS cluster. This subdirectory is used to read data from or write data to the HDFS cluster. If the subdirectory isn't specified, it will default to /.
	Subdirectory pulumi.StringOutput `pulumi:"subdirectory"`
	// Key-value pairs of resource tags to assign to the DataSync Location. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The URL of the Object Storage location that was described.
	Uri pulumi.StringOutput `pulumi:"uri"`
}

// NewLocationObjectStorage registers a new resource with the given unique name, arguments, and options.
func NewLocationObjectStorage(ctx *pulumi.Context,
	name string, args *LocationObjectStorageArgs, opts ...pulumi.ResourceOption) (*LocationObjectStorage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AgentArns == nil {
		return nil, errors.New("invalid value for required argument 'AgentArns'")
	}
	if args.BucketName == nil {
		return nil, errors.New("invalid value for required argument 'BucketName'")
	}
	if args.ServerHostname == nil {
		return nil, errors.New("invalid value for required argument 'ServerHostname'")
	}
	if args.SecretKey != nil {
		args.SecretKey = pulumi.ToSecret(args.SecretKey).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"secretKey",
	})
	opts = append(opts, secrets)
	var resource LocationObjectStorage
	err := ctx.RegisterResource("aws:datasync/locationObjectStorage:LocationObjectStorage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocationObjectStorage gets an existing LocationObjectStorage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocationObjectStorage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocationObjectStorageState, opts ...pulumi.ResourceOption) (*LocationObjectStorage, error) {
	var resource LocationObjectStorage
	err := ctx.ReadResource("aws:datasync/locationObjectStorage:LocationObjectStorage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocationObjectStorage resources.
type locationObjectStorageState struct {
	// The access key is used if credentials are required to access the self-managed object storage server. If your object storage requires a user name and password to authenticate, use `accessKey` and `secretKey` to provide the user name and password, respectively.
	AccessKey *string `pulumi:"accessKey"`
	// A list of DataSync Agent ARNs with which this location will be associated.
	AgentArns []string `pulumi:"agentArns"`
	// Amazon Resource Name (ARN) of the DataSync Location.
	Arn *string `pulumi:"arn"`
	// The bucket on the self-managed object storage server that is used to read data from.
	BucketName *string `pulumi:"bucketName"`
	// The secret key is used if credentials are required to access the self-managed object storage server. If your object storage requires a user name and password to authenticate, use `accessKey` and `secretKey` to provide the user name and password, respectively.
	SecretKey *string `pulumi:"secretKey"`
	// Specifies a certificate to authenticate with an object storage system that uses a private or self-signed certificate authority (CA). You must specify a Base64-encoded .pem file (for example, file:///home/user/.ssh/storage_sys_certificate.pem). The certificate can be up to 32768 bytes (before Base64 encoding).
	ServerCertificate *string `pulumi:"serverCertificate"`
	// The name of the self-managed object storage server. This value is the IP address or Domain Name Service (DNS) name of the object storage server. An agent uses this host name to mount the object storage server in a network.
	ServerHostname *string `pulumi:"serverHostname"`
	// The port that your self-managed object storage server accepts inbound network traffic on. The server port is set by default to TCP 80 (`HTTP`) or TCP 443 (`HTTPS`). You can specify a custom port if your self-managed object storage server requires one.
	ServerPort *int `pulumi:"serverPort"`
	// The protocol that the object storage server uses to communicate. Valid values are `HTTP` or `HTTPS`.
	ServerProtocol *string `pulumi:"serverProtocol"`
	// A subdirectory in the HDFS cluster. This subdirectory is used to read data from or write data to the HDFS cluster. If the subdirectory isn't specified, it will default to /.
	Subdirectory *string `pulumi:"subdirectory"`
	// Key-value pairs of resource tags to assign to the DataSync Location. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The URL of the Object Storage location that was described.
	Uri *string `pulumi:"uri"`
}

type LocationObjectStorageState struct {
	// The access key is used if credentials are required to access the self-managed object storage server. If your object storage requires a user name and password to authenticate, use `accessKey` and `secretKey` to provide the user name and password, respectively.
	AccessKey pulumi.StringPtrInput
	// A list of DataSync Agent ARNs with which this location will be associated.
	AgentArns pulumi.StringArrayInput
	// Amazon Resource Name (ARN) of the DataSync Location.
	Arn pulumi.StringPtrInput
	// The bucket on the self-managed object storage server that is used to read data from.
	BucketName pulumi.StringPtrInput
	// The secret key is used if credentials are required to access the self-managed object storage server. If your object storage requires a user name and password to authenticate, use `accessKey` and `secretKey` to provide the user name and password, respectively.
	SecretKey pulumi.StringPtrInput
	// Specifies a certificate to authenticate with an object storage system that uses a private or self-signed certificate authority (CA). You must specify a Base64-encoded .pem file (for example, file:///home/user/.ssh/storage_sys_certificate.pem). The certificate can be up to 32768 bytes (before Base64 encoding).
	ServerCertificate pulumi.StringPtrInput
	// The name of the self-managed object storage server. This value is the IP address or Domain Name Service (DNS) name of the object storage server. An agent uses this host name to mount the object storage server in a network.
	ServerHostname pulumi.StringPtrInput
	// The port that your self-managed object storage server accepts inbound network traffic on. The server port is set by default to TCP 80 (`HTTP`) or TCP 443 (`HTTPS`). You can specify a custom port if your self-managed object storage server requires one.
	ServerPort pulumi.IntPtrInput
	// The protocol that the object storage server uses to communicate. Valid values are `HTTP` or `HTTPS`.
	ServerProtocol pulumi.StringPtrInput
	// A subdirectory in the HDFS cluster. This subdirectory is used to read data from or write data to the HDFS cluster. If the subdirectory isn't specified, it will default to /.
	Subdirectory pulumi.StringPtrInput
	// Key-value pairs of resource tags to assign to the DataSync Location. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// The URL of the Object Storage location that was described.
	Uri pulumi.StringPtrInput
}

func (LocationObjectStorageState) ElementType() reflect.Type {
	return reflect.TypeOf((*locationObjectStorageState)(nil)).Elem()
}

type locationObjectStorageArgs struct {
	// The access key is used if credentials are required to access the self-managed object storage server. If your object storage requires a user name and password to authenticate, use `accessKey` and `secretKey` to provide the user name and password, respectively.
	AccessKey *string `pulumi:"accessKey"`
	// A list of DataSync Agent ARNs with which this location will be associated.
	AgentArns []string `pulumi:"agentArns"`
	// The bucket on the self-managed object storage server that is used to read data from.
	BucketName string `pulumi:"bucketName"`
	// The secret key is used if credentials are required to access the self-managed object storage server. If your object storage requires a user name and password to authenticate, use `accessKey` and `secretKey` to provide the user name and password, respectively.
	SecretKey *string `pulumi:"secretKey"`
	// Specifies a certificate to authenticate with an object storage system that uses a private or self-signed certificate authority (CA). You must specify a Base64-encoded .pem file (for example, file:///home/user/.ssh/storage_sys_certificate.pem). The certificate can be up to 32768 bytes (before Base64 encoding).
	ServerCertificate *string `pulumi:"serverCertificate"`
	// The name of the self-managed object storage server. This value is the IP address or Domain Name Service (DNS) name of the object storage server. An agent uses this host name to mount the object storage server in a network.
	ServerHostname string `pulumi:"serverHostname"`
	// The port that your self-managed object storage server accepts inbound network traffic on. The server port is set by default to TCP 80 (`HTTP`) or TCP 443 (`HTTPS`). You can specify a custom port if your self-managed object storage server requires one.
	ServerPort *int `pulumi:"serverPort"`
	// The protocol that the object storage server uses to communicate. Valid values are `HTTP` or `HTTPS`.
	ServerProtocol *string `pulumi:"serverProtocol"`
	// A subdirectory in the HDFS cluster. This subdirectory is used to read data from or write data to the HDFS cluster. If the subdirectory isn't specified, it will default to /.
	Subdirectory *string `pulumi:"subdirectory"`
	// Key-value pairs of resource tags to assign to the DataSync Location. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a LocationObjectStorage resource.
type LocationObjectStorageArgs struct {
	// The access key is used if credentials are required to access the self-managed object storage server. If your object storage requires a user name and password to authenticate, use `accessKey` and `secretKey` to provide the user name and password, respectively.
	AccessKey pulumi.StringPtrInput
	// A list of DataSync Agent ARNs with which this location will be associated.
	AgentArns pulumi.StringArrayInput
	// The bucket on the self-managed object storage server that is used to read data from.
	BucketName pulumi.StringInput
	// The secret key is used if credentials are required to access the self-managed object storage server. If your object storage requires a user name and password to authenticate, use `accessKey` and `secretKey` to provide the user name and password, respectively.
	SecretKey pulumi.StringPtrInput
	// Specifies a certificate to authenticate with an object storage system that uses a private or self-signed certificate authority (CA). You must specify a Base64-encoded .pem file (for example, file:///home/user/.ssh/storage_sys_certificate.pem). The certificate can be up to 32768 bytes (before Base64 encoding).
	ServerCertificate pulumi.StringPtrInput
	// The name of the self-managed object storage server. This value is the IP address or Domain Name Service (DNS) name of the object storage server. An agent uses this host name to mount the object storage server in a network.
	ServerHostname pulumi.StringInput
	// The port that your self-managed object storage server accepts inbound network traffic on. The server port is set by default to TCP 80 (`HTTP`) or TCP 443 (`HTTPS`). You can specify a custom port if your self-managed object storage server requires one.
	ServerPort pulumi.IntPtrInput
	// The protocol that the object storage server uses to communicate. Valid values are `HTTP` or `HTTPS`.
	ServerProtocol pulumi.StringPtrInput
	// A subdirectory in the HDFS cluster. This subdirectory is used to read data from or write data to the HDFS cluster. If the subdirectory isn't specified, it will default to /.
	Subdirectory pulumi.StringPtrInput
	// Key-value pairs of resource tags to assign to the DataSync Location. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (LocationObjectStorageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*locationObjectStorageArgs)(nil)).Elem()
}

type LocationObjectStorageInput interface {
	pulumi.Input

	ToLocationObjectStorageOutput() LocationObjectStorageOutput
	ToLocationObjectStorageOutputWithContext(ctx context.Context) LocationObjectStorageOutput
}

func (*LocationObjectStorage) ElementType() reflect.Type {
	return reflect.TypeOf((**LocationObjectStorage)(nil)).Elem()
}

func (i *LocationObjectStorage) ToLocationObjectStorageOutput() LocationObjectStorageOutput {
	return i.ToLocationObjectStorageOutputWithContext(context.Background())
}

func (i *LocationObjectStorage) ToLocationObjectStorageOutputWithContext(ctx context.Context) LocationObjectStorageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationObjectStorageOutput)
}

// LocationObjectStorageArrayInput is an input type that accepts LocationObjectStorageArray and LocationObjectStorageArrayOutput values.
// You can construct a concrete instance of `LocationObjectStorageArrayInput` via:
//
//	LocationObjectStorageArray{ LocationObjectStorageArgs{...} }
type LocationObjectStorageArrayInput interface {
	pulumi.Input

	ToLocationObjectStorageArrayOutput() LocationObjectStorageArrayOutput
	ToLocationObjectStorageArrayOutputWithContext(context.Context) LocationObjectStorageArrayOutput
}

type LocationObjectStorageArray []LocationObjectStorageInput

func (LocationObjectStorageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocationObjectStorage)(nil)).Elem()
}

func (i LocationObjectStorageArray) ToLocationObjectStorageArrayOutput() LocationObjectStorageArrayOutput {
	return i.ToLocationObjectStorageArrayOutputWithContext(context.Background())
}

func (i LocationObjectStorageArray) ToLocationObjectStorageArrayOutputWithContext(ctx context.Context) LocationObjectStorageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationObjectStorageArrayOutput)
}

// LocationObjectStorageMapInput is an input type that accepts LocationObjectStorageMap and LocationObjectStorageMapOutput values.
// You can construct a concrete instance of `LocationObjectStorageMapInput` via:
//
//	LocationObjectStorageMap{ "key": LocationObjectStorageArgs{...} }
type LocationObjectStorageMapInput interface {
	pulumi.Input

	ToLocationObjectStorageMapOutput() LocationObjectStorageMapOutput
	ToLocationObjectStorageMapOutputWithContext(context.Context) LocationObjectStorageMapOutput
}

type LocationObjectStorageMap map[string]LocationObjectStorageInput

func (LocationObjectStorageMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocationObjectStorage)(nil)).Elem()
}

func (i LocationObjectStorageMap) ToLocationObjectStorageMapOutput() LocationObjectStorageMapOutput {
	return i.ToLocationObjectStorageMapOutputWithContext(context.Background())
}

func (i LocationObjectStorageMap) ToLocationObjectStorageMapOutputWithContext(ctx context.Context) LocationObjectStorageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationObjectStorageMapOutput)
}

type LocationObjectStorageOutput struct{ *pulumi.OutputState }

func (LocationObjectStorageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocationObjectStorage)(nil)).Elem()
}

func (o LocationObjectStorageOutput) ToLocationObjectStorageOutput() LocationObjectStorageOutput {
	return o
}

func (o LocationObjectStorageOutput) ToLocationObjectStorageOutputWithContext(ctx context.Context) LocationObjectStorageOutput {
	return o
}

// The access key is used if credentials are required to access the self-managed object storage server. If your object storage requires a user name and password to authenticate, use `accessKey` and `secretKey` to provide the user name and password, respectively.
func (o LocationObjectStorageOutput) AccessKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocationObjectStorage) pulumi.StringPtrOutput { return v.AccessKey }).(pulumi.StringPtrOutput)
}

// A list of DataSync Agent ARNs with which this location will be associated.
func (o LocationObjectStorageOutput) AgentArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocationObjectStorage) pulumi.StringArrayOutput { return v.AgentArns }).(pulumi.StringArrayOutput)
}

// Amazon Resource Name (ARN) of the DataSync Location.
func (o LocationObjectStorageOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *LocationObjectStorage) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The bucket on the self-managed object storage server that is used to read data from.
func (o LocationObjectStorageOutput) BucketName() pulumi.StringOutput {
	return o.ApplyT(func(v *LocationObjectStorage) pulumi.StringOutput { return v.BucketName }).(pulumi.StringOutput)
}

// The secret key is used if credentials are required to access the self-managed object storage server. If your object storage requires a user name and password to authenticate, use `accessKey` and `secretKey` to provide the user name and password, respectively.
func (o LocationObjectStorageOutput) SecretKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocationObjectStorage) pulumi.StringPtrOutput { return v.SecretKey }).(pulumi.StringPtrOutput)
}

// Specifies a certificate to authenticate with an object storage system that uses a private or self-signed certificate authority (CA). You must specify a Base64-encoded .pem file (for example, file:///home/user/.ssh/storage_sys_certificate.pem). The certificate can be up to 32768 bytes (before Base64 encoding).
func (o LocationObjectStorageOutput) ServerCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocationObjectStorage) pulumi.StringPtrOutput { return v.ServerCertificate }).(pulumi.StringPtrOutput)
}

// The name of the self-managed object storage server. This value is the IP address or Domain Name Service (DNS) name of the object storage server. An agent uses this host name to mount the object storage server in a network.
func (o LocationObjectStorageOutput) ServerHostname() pulumi.StringOutput {
	return o.ApplyT(func(v *LocationObjectStorage) pulumi.StringOutput { return v.ServerHostname }).(pulumi.StringOutput)
}

// The port that your self-managed object storage server accepts inbound network traffic on. The server port is set by default to TCP 80 (`HTTP`) or TCP 443 (`HTTPS`). You can specify a custom port if your self-managed object storage server requires one.
func (o LocationObjectStorageOutput) ServerPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LocationObjectStorage) pulumi.IntPtrOutput { return v.ServerPort }).(pulumi.IntPtrOutput)
}

// The protocol that the object storage server uses to communicate. Valid values are `HTTP` or `HTTPS`.
func (o LocationObjectStorageOutput) ServerProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocationObjectStorage) pulumi.StringPtrOutput { return v.ServerProtocol }).(pulumi.StringPtrOutput)
}

// A subdirectory in the HDFS cluster. This subdirectory is used to read data from or write data to the HDFS cluster. If the subdirectory isn't specified, it will default to /.
func (o LocationObjectStorageOutput) Subdirectory() pulumi.StringOutput {
	return o.ApplyT(func(v *LocationObjectStorage) pulumi.StringOutput { return v.Subdirectory }).(pulumi.StringOutput)
}

// Key-value pairs of resource tags to assign to the DataSync Location. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o LocationObjectStorageOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LocationObjectStorage) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o LocationObjectStorageOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LocationObjectStorage) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// The URL of the Object Storage location that was described.
func (o LocationObjectStorageOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v *LocationObjectStorage) pulumi.StringOutput { return v.Uri }).(pulumi.StringOutput)
}

type LocationObjectStorageArrayOutput struct{ *pulumi.OutputState }

func (LocationObjectStorageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocationObjectStorage)(nil)).Elem()
}

func (o LocationObjectStorageArrayOutput) ToLocationObjectStorageArrayOutput() LocationObjectStorageArrayOutput {
	return o
}

func (o LocationObjectStorageArrayOutput) ToLocationObjectStorageArrayOutputWithContext(ctx context.Context) LocationObjectStorageArrayOutput {
	return o
}

func (o LocationObjectStorageArrayOutput) Index(i pulumi.IntInput) LocationObjectStorageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LocationObjectStorage {
		return vs[0].([]*LocationObjectStorage)[vs[1].(int)]
	}).(LocationObjectStorageOutput)
}

type LocationObjectStorageMapOutput struct{ *pulumi.OutputState }

func (LocationObjectStorageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocationObjectStorage)(nil)).Elem()
}

func (o LocationObjectStorageMapOutput) ToLocationObjectStorageMapOutput() LocationObjectStorageMapOutput {
	return o
}

func (o LocationObjectStorageMapOutput) ToLocationObjectStorageMapOutputWithContext(ctx context.Context) LocationObjectStorageMapOutput {
	return o
}

func (o LocationObjectStorageMapOutput) MapIndex(k pulumi.StringInput) LocationObjectStorageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LocationObjectStorage {
		return vs[0].(map[string]*LocationObjectStorage)[vs[1].(string)]
	}).(LocationObjectStorageOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocationObjectStorageInput)(nil)).Elem(), &LocationObjectStorage{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocationObjectStorageArrayInput)(nil)).Elem(), LocationObjectStorageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocationObjectStorageMapInput)(nil)).Elem(), LocationObjectStorageMap{})
	pulumi.RegisterOutputType(LocationObjectStorageOutput{})
	pulumi.RegisterOutputType(LocationObjectStorageArrayOutput{})
	pulumi.RegisterOutputType(LocationObjectStorageMapOutput{})
}
