// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a single-Region or multi-Region primary KMS key.
//
// > **NOTE on KMS Key Policy:** KMS Key Policy can be configured in either the standalone resource `kms.KeyPolicy`
// or with the parameter `policy` in this resource.
// Configuring with both will cause inconsistencies and may overwrite configuration.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/kms"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := kms.NewKey(ctx, "key", &kms.KeyArgs{
//				DeletionWindowInDays: pulumi.Int(10),
//				Description:          pulumi.String("KMS key 1"),
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
// KMS Keys can be imported using the `id`, e.g.,
//
// ```sh
//
//	$ pulumi import aws:kms/key:Key a 1234abcd-12ab-34cd-56ef-1234567890ab
//
// ```
type Key struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the key.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A flag to indicate whether to bypass the key policy lockout safety check.
	// Setting this value to true increases the risk that the KMS key becomes unmanageable. Do not set this value to true indiscriminately.
	// For more information, refer to the scenario in the [Default Key Policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam) section in the _AWS Key Management Service Developer Guide_.
	// The default value is `false`.
	BypassPolicyLockoutSafetyCheck pulumi.BoolPtrOutput `pulumi:"bypassPolicyLockoutSafetyCheck"`
	// ID of the KMS [Custom Key Store](https://docs.aws.amazon.com/kms/latest/developerguide/create-cmk-keystore.html) where the key will be stored instead of KMS (eg CloudHSM).
	CustomKeyStoreId pulumi.StringPtrOutput `pulumi:"customKeyStoreId"`
	// Specifies whether the key contains a symmetric key or an asymmetric key pair and the encryption algorithms or signing algorithms that the key supports.
	// Valid values: `SYMMETRIC_DEFAULT`,  `RSA_2048`, `RSA_3072`, `RSA_4096`, `HMAC_256`, `ECC_NIST_P256`, `ECC_NIST_P384`, `ECC_NIST_P521`, or `ECC_SECG_P256K1`. Defaults to `SYMMETRIC_DEFAULT`. For help with choosing a key spec, see the [AWS KMS Developer Guide](https://docs.aws.amazon.com/kms/latest/developerguide/symm-asymm-choose.html).
	CustomerMasterKeySpec pulumi.StringPtrOutput `pulumi:"customerMasterKeySpec"`
	// The waiting period, specified in number of days. After the waiting period ends, AWS KMS deletes the KMS key.
	// If you specify a value, it must be between `7` and `30`, inclusive. If you do not specify a value, it defaults to `30`.
	// If the KMS key is a multi-Region primary key with replicas, the waiting period begins when the last of its replica keys is deleted. Otherwise, the waiting period begins immediately.
	DeletionWindowInDays pulumi.IntPtrOutput `pulumi:"deletionWindowInDays"`
	// The description of the key as viewed in AWS console.
	Description pulumi.StringOutput `pulumi:"description"`
	// Specifies whether [key rotation](http://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html) is enabled. Defaults to `false`.
	EnableKeyRotation pulumi.BoolPtrOutput `pulumi:"enableKeyRotation"`
	// Specifies whether the key is enabled. Defaults to `true`.
	IsEnabled pulumi.BoolPtrOutput `pulumi:"isEnabled"`
	// The globally unique identifier for the key.
	KeyId pulumi.StringOutput `pulumi:"keyId"`
	// Specifies the intended use of the key. Valid values: `ENCRYPT_DECRYPT`, `SIGN_VERIFY`, or `GENERATE_VERIFY_MAC`.
	// Defaults to `ENCRYPT_DECRYPT`.
	KeyUsage pulumi.StringPtrOutput `pulumi:"keyUsage"`
	// Indicates whether the KMS key is a multi-Region (`true`) or regional (`false`) key. Defaults to `false`.
	MultiRegion pulumi.BoolOutput `pulumi:"multiRegion"`
	// A valid policy JSON document. Although this is a key policy, not an IAM policy, an `iam.getPolicyDocument`, in the form that designates a principal, can be used.
	Policy pulumi.StringOutput `pulumi:"policy"`
	// A map of tags to assign to the object. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewKey registers a new resource with the given unique name, arguments, and options.
func NewKey(ctx *pulumi.Context,
	name string, args *KeyArgs, opts ...pulumi.ResourceOption) (*Key, error) {
	if args == nil {
		args = &KeyArgs{}
	}

	var resource Key
	err := ctx.RegisterResource("aws:kms/key:Key", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKey gets an existing Key resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeyState, opts ...pulumi.ResourceOption) (*Key, error) {
	var resource Key
	err := ctx.ReadResource("aws:kms/key:Key", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Key resources.
type keyState struct {
	// The Amazon Resource Name (ARN) of the key.
	Arn *string `pulumi:"arn"`
	// A flag to indicate whether to bypass the key policy lockout safety check.
	// Setting this value to true increases the risk that the KMS key becomes unmanageable. Do not set this value to true indiscriminately.
	// For more information, refer to the scenario in the [Default Key Policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam) section in the _AWS Key Management Service Developer Guide_.
	// The default value is `false`.
	BypassPolicyLockoutSafetyCheck *bool `pulumi:"bypassPolicyLockoutSafetyCheck"`
	// ID of the KMS [Custom Key Store](https://docs.aws.amazon.com/kms/latest/developerguide/create-cmk-keystore.html) where the key will be stored instead of KMS (eg CloudHSM).
	CustomKeyStoreId *string `pulumi:"customKeyStoreId"`
	// Specifies whether the key contains a symmetric key or an asymmetric key pair and the encryption algorithms or signing algorithms that the key supports.
	// Valid values: `SYMMETRIC_DEFAULT`,  `RSA_2048`, `RSA_3072`, `RSA_4096`, `HMAC_256`, `ECC_NIST_P256`, `ECC_NIST_P384`, `ECC_NIST_P521`, or `ECC_SECG_P256K1`. Defaults to `SYMMETRIC_DEFAULT`. For help with choosing a key spec, see the [AWS KMS Developer Guide](https://docs.aws.amazon.com/kms/latest/developerguide/symm-asymm-choose.html).
	CustomerMasterKeySpec *string `pulumi:"customerMasterKeySpec"`
	// The waiting period, specified in number of days. After the waiting period ends, AWS KMS deletes the KMS key.
	// If you specify a value, it must be between `7` and `30`, inclusive. If you do not specify a value, it defaults to `30`.
	// If the KMS key is a multi-Region primary key with replicas, the waiting period begins when the last of its replica keys is deleted. Otherwise, the waiting period begins immediately.
	DeletionWindowInDays *int `pulumi:"deletionWindowInDays"`
	// The description of the key as viewed in AWS console.
	Description *string `pulumi:"description"`
	// Specifies whether [key rotation](http://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html) is enabled. Defaults to `false`.
	EnableKeyRotation *bool `pulumi:"enableKeyRotation"`
	// Specifies whether the key is enabled. Defaults to `true`.
	IsEnabled *bool `pulumi:"isEnabled"`
	// The globally unique identifier for the key.
	KeyId *string `pulumi:"keyId"`
	// Specifies the intended use of the key. Valid values: `ENCRYPT_DECRYPT`, `SIGN_VERIFY`, or `GENERATE_VERIFY_MAC`.
	// Defaults to `ENCRYPT_DECRYPT`.
	KeyUsage *string `pulumi:"keyUsage"`
	// Indicates whether the KMS key is a multi-Region (`true`) or regional (`false`) key. Defaults to `false`.
	MultiRegion *bool `pulumi:"multiRegion"`
	// A valid policy JSON document. Although this is a key policy, not an IAM policy, an `iam.getPolicyDocument`, in the form that designates a principal, can be used.
	Policy *string `pulumi:"policy"`
	// A map of tags to assign to the object. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type KeyState struct {
	// The Amazon Resource Name (ARN) of the key.
	Arn pulumi.StringPtrInput
	// A flag to indicate whether to bypass the key policy lockout safety check.
	// Setting this value to true increases the risk that the KMS key becomes unmanageable. Do not set this value to true indiscriminately.
	// For more information, refer to the scenario in the [Default Key Policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam) section in the _AWS Key Management Service Developer Guide_.
	// The default value is `false`.
	BypassPolicyLockoutSafetyCheck pulumi.BoolPtrInput
	// ID of the KMS [Custom Key Store](https://docs.aws.amazon.com/kms/latest/developerguide/create-cmk-keystore.html) where the key will be stored instead of KMS (eg CloudHSM).
	CustomKeyStoreId pulumi.StringPtrInput
	// Specifies whether the key contains a symmetric key or an asymmetric key pair and the encryption algorithms or signing algorithms that the key supports.
	// Valid values: `SYMMETRIC_DEFAULT`,  `RSA_2048`, `RSA_3072`, `RSA_4096`, `HMAC_256`, `ECC_NIST_P256`, `ECC_NIST_P384`, `ECC_NIST_P521`, or `ECC_SECG_P256K1`. Defaults to `SYMMETRIC_DEFAULT`. For help with choosing a key spec, see the [AWS KMS Developer Guide](https://docs.aws.amazon.com/kms/latest/developerguide/symm-asymm-choose.html).
	CustomerMasterKeySpec pulumi.StringPtrInput
	// The waiting period, specified in number of days. After the waiting period ends, AWS KMS deletes the KMS key.
	// If you specify a value, it must be between `7` and `30`, inclusive. If you do not specify a value, it defaults to `30`.
	// If the KMS key is a multi-Region primary key with replicas, the waiting period begins when the last of its replica keys is deleted. Otherwise, the waiting period begins immediately.
	DeletionWindowInDays pulumi.IntPtrInput
	// The description of the key as viewed in AWS console.
	Description pulumi.StringPtrInput
	// Specifies whether [key rotation](http://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html) is enabled. Defaults to `false`.
	EnableKeyRotation pulumi.BoolPtrInput
	// Specifies whether the key is enabled. Defaults to `true`.
	IsEnabled pulumi.BoolPtrInput
	// The globally unique identifier for the key.
	KeyId pulumi.StringPtrInput
	// Specifies the intended use of the key. Valid values: `ENCRYPT_DECRYPT`, `SIGN_VERIFY`, or `GENERATE_VERIFY_MAC`.
	// Defaults to `ENCRYPT_DECRYPT`.
	KeyUsage pulumi.StringPtrInput
	// Indicates whether the KMS key is a multi-Region (`true`) or regional (`false`) key. Defaults to `false`.
	MultiRegion pulumi.BoolPtrInput
	// A valid policy JSON document. Although this is a key policy, not an IAM policy, an `iam.getPolicyDocument`, in the form that designates a principal, can be used.
	Policy pulumi.StringPtrInput
	// A map of tags to assign to the object. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
}

func (KeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*keyState)(nil)).Elem()
}

type keyArgs struct {
	// A flag to indicate whether to bypass the key policy lockout safety check.
	// Setting this value to true increases the risk that the KMS key becomes unmanageable. Do not set this value to true indiscriminately.
	// For more information, refer to the scenario in the [Default Key Policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam) section in the _AWS Key Management Service Developer Guide_.
	// The default value is `false`.
	BypassPolicyLockoutSafetyCheck *bool `pulumi:"bypassPolicyLockoutSafetyCheck"`
	// ID of the KMS [Custom Key Store](https://docs.aws.amazon.com/kms/latest/developerguide/create-cmk-keystore.html) where the key will be stored instead of KMS (eg CloudHSM).
	CustomKeyStoreId *string `pulumi:"customKeyStoreId"`
	// Specifies whether the key contains a symmetric key or an asymmetric key pair and the encryption algorithms or signing algorithms that the key supports.
	// Valid values: `SYMMETRIC_DEFAULT`,  `RSA_2048`, `RSA_3072`, `RSA_4096`, `HMAC_256`, `ECC_NIST_P256`, `ECC_NIST_P384`, `ECC_NIST_P521`, or `ECC_SECG_P256K1`. Defaults to `SYMMETRIC_DEFAULT`. For help with choosing a key spec, see the [AWS KMS Developer Guide](https://docs.aws.amazon.com/kms/latest/developerguide/symm-asymm-choose.html).
	CustomerMasterKeySpec *string `pulumi:"customerMasterKeySpec"`
	// The waiting period, specified in number of days. After the waiting period ends, AWS KMS deletes the KMS key.
	// If you specify a value, it must be between `7` and `30`, inclusive. If you do not specify a value, it defaults to `30`.
	// If the KMS key is a multi-Region primary key with replicas, the waiting period begins when the last of its replica keys is deleted. Otherwise, the waiting period begins immediately.
	DeletionWindowInDays *int `pulumi:"deletionWindowInDays"`
	// The description of the key as viewed in AWS console.
	Description *string `pulumi:"description"`
	// Specifies whether [key rotation](http://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html) is enabled. Defaults to `false`.
	EnableKeyRotation *bool `pulumi:"enableKeyRotation"`
	// Specifies whether the key is enabled. Defaults to `true`.
	IsEnabled *bool `pulumi:"isEnabled"`
	// Specifies the intended use of the key. Valid values: `ENCRYPT_DECRYPT`, `SIGN_VERIFY`, or `GENERATE_VERIFY_MAC`.
	// Defaults to `ENCRYPT_DECRYPT`.
	KeyUsage *string `pulumi:"keyUsage"`
	// Indicates whether the KMS key is a multi-Region (`true`) or regional (`false`) key. Defaults to `false`.
	MultiRegion *bool `pulumi:"multiRegion"`
	// A valid policy JSON document. Although this is a key policy, not an IAM policy, an `iam.getPolicyDocument`, in the form that designates a principal, can be used.
	Policy *string `pulumi:"policy"`
	// A map of tags to assign to the object. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Key resource.
type KeyArgs struct {
	// A flag to indicate whether to bypass the key policy lockout safety check.
	// Setting this value to true increases the risk that the KMS key becomes unmanageable. Do not set this value to true indiscriminately.
	// For more information, refer to the scenario in the [Default Key Policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam) section in the _AWS Key Management Service Developer Guide_.
	// The default value is `false`.
	BypassPolicyLockoutSafetyCheck pulumi.BoolPtrInput
	// ID of the KMS [Custom Key Store](https://docs.aws.amazon.com/kms/latest/developerguide/create-cmk-keystore.html) where the key will be stored instead of KMS (eg CloudHSM).
	CustomKeyStoreId pulumi.StringPtrInput
	// Specifies whether the key contains a symmetric key or an asymmetric key pair and the encryption algorithms or signing algorithms that the key supports.
	// Valid values: `SYMMETRIC_DEFAULT`,  `RSA_2048`, `RSA_3072`, `RSA_4096`, `HMAC_256`, `ECC_NIST_P256`, `ECC_NIST_P384`, `ECC_NIST_P521`, or `ECC_SECG_P256K1`. Defaults to `SYMMETRIC_DEFAULT`. For help with choosing a key spec, see the [AWS KMS Developer Guide](https://docs.aws.amazon.com/kms/latest/developerguide/symm-asymm-choose.html).
	CustomerMasterKeySpec pulumi.StringPtrInput
	// The waiting period, specified in number of days. After the waiting period ends, AWS KMS deletes the KMS key.
	// If you specify a value, it must be between `7` and `30`, inclusive. If you do not specify a value, it defaults to `30`.
	// If the KMS key is a multi-Region primary key with replicas, the waiting period begins when the last of its replica keys is deleted. Otherwise, the waiting period begins immediately.
	DeletionWindowInDays pulumi.IntPtrInput
	// The description of the key as viewed in AWS console.
	Description pulumi.StringPtrInput
	// Specifies whether [key rotation](http://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html) is enabled. Defaults to `false`.
	EnableKeyRotation pulumi.BoolPtrInput
	// Specifies whether the key is enabled. Defaults to `true`.
	IsEnabled pulumi.BoolPtrInput
	// Specifies the intended use of the key. Valid values: `ENCRYPT_DECRYPT`, `SIGN_VERIFY`, or `GENERATE_VERIFY_MAC`.
	// Defaults to `ENCRYPT_DECRYPT`.
	KeyUsage pulumi.StringPtrInput
	// Indicates whether the KMS key is a multi-Region (`true`) or regional (`false`) key. Defaults to `false`.
	MultiRegion pulumi.BoolPtrInput
	// A valid policy JSON document. Although this is a key policy, not an IAM policy, an `iam.getPolicyDocument`, in the form that designates a principal, can be used.
	Policy pulumi.StringPtrInput
	// A map of tags to assign to the object. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (KeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keyArgs)(nil)).Elem()
}

type KeyInput interface {
	pulumi.Input

	ToKeyOutput() KeyOutput
	ToKeyOutputWithContext(ctx context.Context) KeyOutput
}

func (*Key) ElementType() reflect.Type {
	return reflect.TypeOf((**Key)(nil)).Elem()
}

func (i *Key) ToKeyOutput() KeyOutput {
	return i.ToKeyOutputWithContext(context.Background())
}

func (i *Key) ToKeyOutputWithContext(ctx context.Context) KeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyOutput)
}

// KeyArrayInput is an input type that accepts KeyArray and KeyArrayOutput values.
// You can construct a concrete instance of `KeyArrayInput` via:
//
//	KeyArray{ KeyArgs{...} }
type KeyArrayInput interface {
	pulumi.Input

	ToKeyArrayOutput() KeyArrayOutput
	ToKeyArrayOutputWithContext(context.Context) KeyArrayOutput
}

type KeyArray []KeyInput

func (KeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Key)(nil)).Elem()
}

func (i KeyArray) ToKeyArrayOutput() KeyArrayOutput {
	return i.ToKeyArrayOutputWithContext(context.Background())
}

func (i KeyArray) ToKeyArrayOutputWithContext(ctx context.Context) KeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyArrayOutput)
}

// KeyMapInput is an input type that accepts KeyMap and KeyMapOutput values.
// You can construct a concrete instance of `KeyMapInput` via:
//
//	KeyMap{ "key": KeyArgs{...} }
type KeyMapInput interface {
	pulumi.Input

	ToKeyMapOutput() KeyMapOutput
	ToKeyMapOutputWithContext(context.Context) KeyMapOutput
}

type KeyMap map[string]KeyInput

func (KeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Key)(nil)).Elem()
}

func (i KeyMap) ToKeyMapOutput() KeyMapOutput {
	return i.ToKeyMapOutputWithContext(context.Background())
}

func (i KeyMap) ToKeyMapOutputWithContext(ctx context.Context) KeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyMapOutput)
}

type KeyOutput struct{ *pulumi.OutputState }

func (KeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Key)(nil)).Elem()
}

func (o KeyOutput) ToKeyOutput() KeyOutput {
	return o
}

func (o KeyOutput) ToKeyOutputWithContext(ctx context.Context) KeyOutput {
	return o
}

// The Amazon Resource Name (ARN) of the key.
func (o KeyOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Key) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// A flag to indicate whether to bypass the key policy lockout safety check.
// Setting this value to true increases the risk that the KMS key becomes unmanageable. Do not set this value to true indiscriminately.
// For more information, refer to the scenario in the [Default Key Policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam) section in the _AWS Key Management Service Developer Guide_.
// The default value is `false`.
func (o KeyOutput) BypassPolicyLockoutSafetyCheck() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Key) pulumi.BoolPtrOutput { return v.BypassPolicyLockoutSafetyCheck }).(pulumi.BoolPtrOutput)
}

// ID of the KMS [Custom Key Store](https://docs.aws.amazon.com/kms/latest/developerguide/create-cmk-keystore.html) where the key will be stored instead of KMS (eg CloudHSM).
func (o KeyOutput) CustomKeyStoreId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Key) pulumi.StringPtrOutput { return v.CustomKeyStoreId }).(pulumi.StringPtrOutput)
}

// Specifies whether the key contains a symmetric key or an asymmetric key pair and the encryption algorithms or signing algorithms that the key supports.
// Valid values: `SYMMETRIC_DEFAULT`,  `RSA_2048`, `RSA_3072`, `RSA_4096`, `HMAC_256`, `ECC_NIST_P256`, `ECC_NIST_P384`, `ECC_NIST_P521`, or `ECC_SECG_P256K1`. Defaults to `SYMMETRIC_DEFAULT`. For help with choosing a key spec, see the [AWS KMS Developer Guide](https://docs.aws.amazon.com/kms/latest/developerguide/symm-asymm-choose.html).
func (o KeyOutput) CustomerMasterKeySpec() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Key) pulumi.StringPtrOutput { return v.CustomerMasterKeySpec }).(pulumi.StringPtrOutput)
}

// The waiting period, specified in number of days. After the waiting period ends, AWS KMS deletes the KMS key.
// If you specify a value, it must be between `7` and `30`, inclusive. If you do not specify a value, it defaults to `30`.
// If the KMS key is a multi-Region primary key with replicas, the waiting period begins when the last of its replica keys is deleted. Otherwise, the waiting period begins immediately.
func (o KeyOutput) DeletionWindowInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Key) pulumi.IntPtrOutput { return v.DeletionWindowInDays }).(pulumi.IntPtrOutput)
}

// The description of the key as viewed in AWS console.
func (o KeyOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Key) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Specifies whether [key rotation](http://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html) is enabled. Defaults to `false`.
func (o KeyOutput) EnableKeyRotation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Key) pulumi.BoolPtrOutput { return v.EnableKeyRotation }).(pulumi.BoolPtrOutput)
}

// Specifies whether the key is enabled. Defaults to `true`.
func (o KeyOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Key) pulumi.BoolPtrOutput { return v.IsEnabled }).(pulumi.BoolPtrOutput)
}

// The globally unique identifier for the key.
func (o KeyOutput) KeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *Key) pulumi.StringOutput { return v.KeyId }).(pulumi.StringOutput)
}

// Specifies the intended use of the key. Valid values: `ENCRYPT_DECRYPT`, `SIGN_VERIFY`, or `GENERATE_VERIFY_MAC`.
// Defaults to `ENCRYPT_DECRYPT`.
func (o KeyOutput) KeyUsage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Key) pulumi.StringPtrOutput { return v.KeyUsage }).(pulumi.StringPtrOutput)
}

// Indicates whether the KMS key is a multi-Region (`true`) or regional (`false`) key. Defaults to `false`.
func (o KeyOutput) MultiRegion() pulumi.BoolOutput {
	return o.ApplyT(func(v *Key) pulumi.BoolOutput { return v.MultiRegion }).(pulumi.BoolOutput)
}

// A valid policy JSON document. Although this is a key policy, not an IAM policy, an `iam.getPolicyDocument`, in the form that designates a principal, can be used.
func (o KeyOutput) Policy() pulumi.StringOutput {
	return o.ApplyT(func(v *Key) pulumi.StringOutput { return v.Policy }).(pulumi.StringOutput)
}

// A map of tags to assign to the object. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o KeyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Key) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o KeyOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Key) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type KeyArrayOutput struct{ *pulumi.OutputState }

func (KeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Key)(nil)).Elem()
}

func (o KeyArrayOutput) ToKeyArrayOutput() KeyArrayOutput {
	return o
}

func (o KeyArrayOutput) ToKeyArrayOutputWithContext(ctx context.Context) KeyArrayOutput {
	return o
}

func (o KeyArrayOutput) Index(i pulumi.IntInput) KeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Key {
		return vs[0].([]*Key)[vs[1].(int)]
	}).(KeyOutput)
}

type KeyMapOutput struct{ *pulumi.OutputState }

func (KeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Key)(nil)).Elem()
}

func (o KeyMapOutput) ToKeyMapOutput() KeyMapOutput {
	return o
}

func (o KeyMapOutput) ToKeyMapOutputWithContext(ctx context.Context) KeyMapOutput {
	return o
}

func (o KeyMapOutput) MapIndex(k pulumi.StringInput) KeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Key {
		return vs[0].(map[string]*Key)[vs[1].(string)]
	}).(KeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KeyInput)(nil)).Elem(), &Key{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyArrayInput)(nil)).Elem(), KeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyMapInput)(nil)).Elem(), KeyMap{})
	pulumi.RegisterOutputType(KeyOutput{})
	pulumi.RegisterOutputType(KeyArrayOutput{})
	pulumi.RegisterOutputType(KeyMapOutput{})
}
