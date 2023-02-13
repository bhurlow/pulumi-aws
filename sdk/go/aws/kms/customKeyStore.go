// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource for managing an AWS KMS (Key Management) Custom Key Store.
//
// ## Example Usage
// ### Basic Usage
//
// ```go
// package main
//
// import (
//
//	"os"
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/kms"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func readFileOrPanic(path string) pulumi.StringPtrInput {
//		data, err := os.ReadFile(path)
//		if err != nil {
//			panic(err.Error())
//		}
//		return pulumi.String(string(data))
//	}
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := kms.NewCustomKeyStore(ctx, "test", &kms.CustomKeyStoreArgs{
//				CloudHsmClusterId:      pulumi.Any(_var.Cloud_hsm_cluster_id),
//				CustomKeyStoreName:     pulumi.String("kms-custom-key-store-test"),
//				KeyStorePassword:       pulumi.String("noplaintextpasswords1"),
//				TrustAnchorCertificate: readFileOrPanic("anchor-certificate.crt"),
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
// KMS (Key Management) Custom Key Store can be imported using the `id`, e.g.,
//
// ```sh
//
//	$ pulumi import aws:kms/customKeyStore:CustomKeyStore example cks-5ebd4ef395a96288e
//
// ```
type CustomKeyStore struct {
	pulumi.CustomResourceState

	// Cluster ID of CloudHSM.
	CloudHsmClusterId pulumi.StringOutput `pulumi:"cloudHsmClusterId"`
	// Unique name for Custom Key Store.
	CustomKeyStoreName pulumi.StringOutput `pulumi:"customKeyStoreName"`
	// Password for `kmsuser` on CloudHSM.
	KeyStorePassword pulumi.StringOutput `pulumi:"keyStorePassword"`
	// Customer certificate used for signing on CloudHSM.
	TrustAnchorCertificate pulumi.StringOutput `pulumi:"trustAnchorCertificate"`
}

// NewCustomKeyStore registers a new resource with the given unique name, arguments, and options.
func NewCustomKeyStore(ctx *pulumi.Context,
	name string, args *CustomKeyStoreArgs, opts ...pulumi.ResourceOption) (*CustomKeyStore, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CloudHsmClusterId == nil {
		return nil, errors.New("invalid value for required argument 'CloudHsmClusterId'")
	}
	if args.CustomKeyStoreName == nil {
		return nil, errors.New("invalid value for required argument 'CustomKeyStoreName'")
	}
	if args.KeyStorePassword == nil {
		return nil, errors.New("invalid value for required argument 'KeyStorePassword'")
	}
	if args.TrustAnchorCertificate == nil {
		return nil, errors.New("invalid value for required argument 'TrustAnchorCertificate'")
	}
	var resource CustomKeyStore
	err := ctx.RegisterResource("aws:kms/customKeyStore:CustomKeyStore", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomKeyStore gets an existing CustomKeyStore resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomKeyStore(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomKeyStoreState, opts ...pulumi.ResourceOption) (*CustomKeyStore, error) {
	var resource CustomKeyStore
	err := ctx.ReadResource("aws:kms/customKeyStore:CustomKeyStore", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomKeyStore resources.
type customKeyStoreState struct {
	// Cluster ID of CloudHSM.
	CloudHsmClusterId *string `pulumi:"cloudHsmClusterId"`
	// Unique name for Custom Key Store.
	CustomKeyStoreName *string `pulumi:"customKeyStoreName"`
	// Password for `kmsuser` on CloudHSM.
	KeyStorePassword *string `pulumi:"keyStorePassword"`
	// Customer certificate used for signing on CloudHSM.
	TrustAnchorCertificate *string `pulumi:"trustAnchorCertificate"`
}

type CustomKeyStoreState struct {
	// Cluster ID of CloudHSM.
	CloudHsmClusterId pulumi.StringPtrInput
	// Unique name for Custom Key Store.
	CustomKeyStoreName pulumi.StringPtrInput
	// Password for `kmsuser` on CloudHSM.
	KeyStorePassword pulumi.StringPtrInput
	// Customer certificate used for signing on CloudHSM.
	TrustAnchorCertificate pulumi.StringPtrInput
}

func (CustomKeyStoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*customKeyStoreState)(nil)).Elem()
}

type customKeyStoreArgs struct {
	// Cluster ID of CloudHSM.
	CloudHsmClusterId string `pulumi:"cloudHsmClusterId"`
	// Unique name for Custom Key Store.
	CustomKeyStoreName string `pulumi:"customKeyStoreName"`
	// Password for `kmsuser` on CloudHSM.
	KeyStorePassword string `pulumi:"keyStorePassword"`
	// Customer certificate used for signing on CloudHSM.
	TrustAnchorCertificate string `pulumi:"trustAnchorCertificate"`
}

// The set of arguments for constructing a CustomKeyStore resource.
type CustomKeyStoreArgs struct {
	// Cluster ID of CloudHSM.
	CloudHsmClusterId pulumi.StringInput
	// Unique name for Custom Key Store.
	CustomKeyStoreName pulumi.StringInput
	// Password for `kmsuser` on CloudHSM.
	KeyStorePassword pulumi.StringInput
	// Customer certificate used for signing on CloudHSM.
	TrustAnchorCertificate pulumi.StringInput
}

func (CustomKeyStoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customKeyStoreArgs)(nil)).Elem()
}

type CustomKeyStoreInput interface {
	pulumi.Input

	ToCustomKeyStoreOutput() CustomKeyStoreOutput
	ToCustomKeyStoreOutputWithContext(ctx context.Context) CustomKeyStoreOutput
}

func (*CustomKeyStore) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomKeyStore)(nil)).Elem()
}

func (i *CustomKeyStore) ToCustomKeyStoreOutput() CustomKeyStoreOutput {
	return i.ToCustomKeyStoreOutputWithContext(context.Background())
}

func (i *CustomKeyStore) ToCustomKeyStoreOutputWithContext(ctx context.Context) CustomKeyStoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomKeyStoreOutput)
}

// CustomKeyStoreArrayInput is an input type that accepts CustomKeyStoreArray and CustomKeyStoreArrayOutput values.
// You can construct a concrete instance of `CustomKeyStoreArrayInput` via:
//
//	CustomKeyStoreArray{ CustomKeyStoreArgs{...} }
type CustomKeyStoreArrayInput interface {
	pulumi.Input

	ToCustomKeyStoreArrayOutput() CustomKeyStoreArrayOutput
	ToCustomKeyStoreArrayOutputWithContext(context.Context) CustomKeyStoreArrayOutput
}

type CustomKeyStoreArray []CustomKeyStoreInput

func (CustomKeyStoreArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomKeyStore)(nil)).Elem()
}

func (i CustomKeyStoreArray) ToCustomKeyStoreArrayOutput() CustomKeyStoreArrayOutput {
	return i.ToCustomKeyStoreArrayOutputWithContext(context.Background())
}

func (i CustomKeyStoreArray) ToCustomKeyStoreArrayOutputWithContext(ctx context.Context) CustomKeyStoreArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomKeyStoreArrayOutput)
}

// CustomKeyStoreMapInput is an input type that accepts CustomKeyStoreMap and CustomKeyStoreMapOutput values.
// You can construct a concrete instance of `CustomKeyStoreMapInput` via:
//
//	CustomKeyStoreMap{ "key": CustomKeyStoreArgs{...} }
type CustomKeyStoreMapInput interface {
	pulumi.Input

	ToCustomKeyStoreMapOutput() CustomKeyStoreMapOutput
	ToCustomKeyStoreMapOutputWithContext(context.Context) CustomKeyStoreMapOutput
}

type CustomKeyStoreMap map[string]CustomKeyStoreInput

func (CustomKeyStoreMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomKeyStore)(nil)).Elem()
}

func (i CustomKeyStoreMap) ToCustomKeyStoreMapOutput() CustomKeyStoreMapOutput {
	return i.ToCustomKeyStoreMapOutputWithContext(context.Background())
}

func (i CustomKeyStoreMap) ToCustomKeyStoreMapOutputWithContext(ctx context.Context) CustomKeyStoreMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomKeyStoreMapOutput)
}

type CustomKeyStoreOutput struct{ *pulumi.OutputState }

func (CustomKeyStoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomKeyStore)(nil)).Elem()
}

func (o CustomKeyStoreOutput) ToCustomKeyStoreOutput() CustomKeyStoreOutput {
	return o
}

func (o CustomKeyStoreOutput) ToCustomKeyStoreOutputWithContext(ctx context.Context) CustomKeyStoreOutput {
	return o
}

// Cluster ID of CloudHSM.
func (o CustomKeyStoreOutput) CloudHsmClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomKeyStore) pulumi.StringOutput { return v.CloudHsmClusterId }).(pulumi.StringOutput)
}

// Unique name for Custom Key Store.
func (o CustomKeyStoreOutput) CustomKeyStoreName() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomKeyStore) pulumi.StringOutput { return v.CustomKeyStoreName }).(pulumi.StringOutput)
}

// Password for `kmsuser` on CloudHSM.
func (o CustomKeyStoreOutput) KeyStorePassword() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomKeyStore) pulumi.StringOutput { return v.KeyStorePassword }).(pulumi.StringOutput)
}

// Customer certificate used for signing on CloudHSM.
func (o CustomKeyStoreOutput) TrustAnchorCertificate() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomKeyStore) pulumi.StringOutput { return v.TrustAnchorCertificate }).(pulumi.StringOutput)
}

type CustomKeyStoreArrayOutput struct{ *pulumi.OutputState }

func (CustomKeyStoreArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomKeyStore)(nil)).Elem()
}

func (o CustomKeyStoreArrayOutput) ToCustomKeyStoreArrayOutput() CustomKeyStoreArrayOutput {
	return o
}

func (o CustomKeyStoreArrayOutput) ToCustomKeyStoreArrayOutputWithContext(ctx context.Context) CustomKeyStoreArrayOutput {
	return o
}

func (o CustomKeyStoreArrayOutput) Index(i pulumi.IntInput) CustomKeyStoreOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CustomKeyStore {
		return vs[0].([]*CustomKeyStore)[vs[1].(int)]
	}).(CustomKeyStoreOutput)
}

type CustomKeyStoreMapOutput struct{ *pulumi.OutputState }

func (CustomKeyStoreMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomKeyStore)(nil)).Elem()
}

func (o CustomKeyStoreMapOutput) ToCustomKeyStoreMapOutput() CustomKeyStoreMapOutput {
	return o
}

func (o CustomKeyStoreMapOutput) ToCustomKeyStoreMapOutputWithContext(ctx context.Context) CustomKeyStoreMapOutput {
	return o
}

func (o CustomKeyStoreMapOutput) MapIndex(k pulumi.StringInput) CustomKeyStoreOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CustomKeyStore {
		return vs[0].(map[string]*CustomKeyStore)[vs[1].(string)]
	}).(CustomKeyStoreOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CustomKeyStoreInput)(nil)).Elem(), &CustomKeyStore{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomKeyStoreArrayInput)(nil)).Elem(), CustomKeyStoreArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomKeyStoreMapInput)(nil)).Elem(), CustomKeyStoreMap{})
	pulumi.RegisterOutputType(CustomKeyStoreOutput{})
	pulumi.RegisterOutputType(CustomKeyStoreArrayOutput{})
	pulumi.RegisterOutputType(CustomKeyStoreMapOutput{})
}
