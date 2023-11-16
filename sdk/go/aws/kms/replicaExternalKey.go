// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a KMS multi-Region replica key that uses external key material.
// See the [AWS KMS Developer Guide](https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-import.html) for more information on importing key material into multi-Region keys.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kms"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := aws.NewProvider(ctx, "primary", &aws.ProviderArgs{
//				Region: pulumi.String("us-east-1"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = kms.NewExternalKey(ctx, "primaryExternalKey", &kms.ExternalKeyArgs{
//				Description:          pulumi.String("Multi-Region primary key"),
//				DeletionWindowInDays: pulumi.Int(30),
//				MultiRegion:          pulumi.Bool(true),
//				Enabled:              pulumi.Bool(true),
//				KeyMaterialBase64:    pulumi.String("..."),
//			}, pulumi.Provider(aws.Primary))
//			if err != nil {
//				return err
//			}
//			_, err = kms.NewReplicaExternalKey(ctx, "replica", &kms.ReplicaExternalKeyArgs{
//				Description:          pulumi.String("Multi-Region replica key"),
//				DeletionWindowInDays: pulumi.Int(7),
//				PrimaryKeyArn:        pulumi.Any(aws_kms_external.Primary.Arn),
//				KeyMaterialBase64:    pulumi.String("..."),
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
// Using `pulumi import`, import KMS multi-Region replica keys using the `id`. For example:
//
// ```sh
//
//	$ pulumi import aws:kms/replicaExternalKey:ReplicaExternalKey example 1234abcd-12ab-34cd-56ef-1234567890ab
//
// ```
type ReplicaExternalKey struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the replica key. The key ARNs of related multi-Region keys differ only in the Region value.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A flag to indicate whether to bypass the key policy lockout safety check.
	// Setting this value to true increases the risk that the KMS key becomes unmanageable. Do not set this value to true indiscriminately.
	// For more information, refer to the scenario in the [Default Key Policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam) section in the _AWS Key Management Service Developer Guide_.
	// The default value is `false`.
	BypassPolicyLockoutSafetyCheck pulumi.BoolPtrOutput `pulumi:"bypassPolicyLockoutSafetyCheck"`
	// The waiting period, specified in number of days. After the waiting period ends, AWS KMS deletes the KMS key.
	// If you specify a value, it must be between `7` and `30`, inclusive. If you do not specify a value, it defaults to `30`.
	DeletionWindowInDays pulumi.IntPtrOutput `pulumi:"deletionWindowInDays"`
	// A description of the KMS key.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies whether the replica key is enabled. Disabled KMS keys cannot be used in cryptographic operations. Keys pending import can only be `false`. Imported keys default to `true` unless expired.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// Whether the key material expires. Empty when pending key material import, otherwise `KEY_MATERIAL_EXPIRES` or `KEY_MATERIAL_DOES_NOT_EXPIRE`.
	ExpirationModel pulumi.StringOutput `pulumi:"expirationModel"`
	// The key ID of the replica key. Related multi-Region keys have the same key ID.
	KeyId pulumi.StringOutput `pulumi:"keyId"`
	// Base64 encoded 256-bit symmetric encryption key material to import. The KMS key is permanently associated with this key material. The same key material can be [reimported](https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html#reimport-key-material), but you cannot import different key material.
	KeyMaterialBase64 pulumi.StringPtrOutput `pulumi:"keyMaterialBase64"`
	// The state of the replica key.
	KeyState pulumi.StringOutput `pulumi:"keyState"`
	// The [cryptographic operations](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations) for which you can use the KMS key. This is a shared property of multi-Region keys.
	KeyUsage pulumi.StringOutput `pulumi:"keyUsage"`
	// The key policy to attach to the KMS key. If you do not specify a key policy, AWS KMS attaches the [default key policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default) to the KMS key.
	Policy pulumi.StringOutput `pulumi:"policy"`
	// The ARN of the multi-Region primary key to replicate. The primary key must be in a different AWS Region of the same AWS Partition. You can create only one replica of a given primary key in each AWS Region.
	PrimaryKeyArn pulumi.StringOutput `pulumi:"primaryKeyArn"`
	// A map of tags to assign to the replica key. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Time at which the imported key material expires. When the key material expires, AWS KMS deletes the key material and the key becomes unusable. If not specified, key material does not expire. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
	ValidTo pulumi.StringPtrOutput `pulumi:"validTo"`
}

// NewReplicaExternalKey registers a new resource with the given unique name, arguments, and options.
func NewReplicaExternalKey(ctx *pulumi.Context,
	name string, args *ReplicaExternalKeyArgs, opts ...pulumi.ResourceOption) (*ReplicaExternalKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrimaryKeyArn == nil {
		return nil, errors.New("invalid value for required argument 'PrimaryKeyArn'")
	}
	if args.KeyMaterialBase64 != nil {
		args.KeyMaterialBase64 = pulumi.ToSecret(args.KeyMaterialBase64).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"keyMaterialBase64",
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ReplicaExternalKey
	err := ctx.RegisterResource("aws:kms/replicaExternalKey:ReplicaExternalKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicaExternalKey gets an existing ReplicaExternalKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicaExternalKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicaExternalKeyState, opts ...pulumi.ResourceOption) (*ReplicaExternalKey, error) {
	var resource ReplicaExternalKey
	err := ctx.ReadResource("aws:kms/replicaExternalKey:ReplicaExternalKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicaExternalKey resources.
type replicaExternalKeyState struct {
	// The Amazon Resource Name (ARN) of the replica key. The key ARNs of related multi-Region keys differ only in the Region value.
	Arn *string `pulumi:"arn"`
	// A flag to indicate whether to bypass the key policy lockout safety check.
	// Setting this value to true increases the risk that the KMS key becomes unmanageable. Do not set this value to true indiscriminately.
	// For more information, refer to the scenario in the [Default Key Policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam) section in the _AWS Key Management Service Developer Guide_.
	// The default value is `false`.
	BypassPolicyLockoutSafetyCheck *bool `pulumi:"bypassPolicyLockoutSafetyCheck"`
	// The waiting period, specified in number of days. After the waiting period ends, AWS KMS deletes the KMS key.
	// If you specify a value, it must be between `7` and `30`, inclusive. If you do not specify a value, it defaults to `30`.
	DeletionWindowInDays *int `pulumi:"deletionWindowInDays"`
	// A description of the KMS key.
	Description *string `pulumi:"description"`
	// Specifies whether the replica key is enabled. Disabled KMS keys cannot be used in cryptographic operations. Keys pending import can only be `false`. Imported keys default to `true` unless expired.
	Enabled *bool `pulumi:"enabled"`
	// Whether the key material expires. Empty when pending key material import, otherwise `KEY_MATERIAL_EXPIRES` or `KEY_MATERIAL_DOES_NOT_EXPIRE`.
	ExpirationModel *string `pulumi:"expirationModel"`
	// The key ID of the replica key. Related multi-Region keys have the same key ID.
	KeyId *string `pulumi:"keyId"`
	// Base64 encoded 256-bit symmetric encryption key material to import. The KMS key is permanently associated with this key material. The same key material can be [reimported](https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html#reimport-key-material), but you cannot import different key material.
	KeyMaterialBase64 *string `pulumi:"keyMaterialBase64"`
	// The state of the replica key.
	KeyState *string `pulumi:"keyState"`
	// The [cryptographic operations](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations) for which you can use the KMS key. This is a shared property of multi-Region keys.
	KeyUsage *string `pulumi:"keyUsage"`
	// The key policy to attach to the KMS key. If you do not specify a key policy, AWS KMS attaches the [default key policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default) to the KMS key.
	Policy *string `pulumi:"policy"`
	// The ARN of the multi-Region primary key to replicate. The primary key must be in a different AWS Region of the same AWS Partition. You can create only one replica of a given primary key in each AWS Region.
	PrimaryKeyArn *string `pulumi:"primaryKeyArn"`
	// A map of tags to assign to the replica key. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Time at which the imported key material expires. When the key material expires, AWS KMS deletes the key material and the key becomes unusable. If not specified, key material does not expire. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
	ValidTo *string `pulumi:"validTo"`
}

type ReplicaExternalKeyState struct {
	// The Amazon Resource Name (ARN) of the replica key. The key ARNs of related multi-Region keys differ only in the Region value.
	Arn pulumi.StringPtrInput
	// A flag to indicate whether to bypass the key policy lockout safety check.
	// Setting this value to true increases the risk that the KMS key becomes unmanageable. Do not set this value to true indiscriminately.
	// For more information, refer to the scenario in the [Default Key Policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam) section in the _AWS Key Management Service Developer Guide_.
	// The default value is `false`.
	BypassPolicyLockoutSafetyCheck pulumi.BoolPtrInput
	// The waiting period, specified in number of days. After the waiting period ends, AWS KMS deletes the KMS key.
	// If you specify a value, it must be between `7` and `30`, inclusive. If you do not specify a value, it defaults to `30`.
	DeletionWindowInDays pulumi.IntPtrInput
	// A description of the KMS key.
	Description pulumi.StringPtrInput
	// Specifies whether the replica key is enabled. Disabled KMS keys cannot be used in cryptographic operations. Keys pending import can only be `false`. Imported keys default to `true` unless expired.
	Enabled pulumi.BoolPtrInput
	// Whether the key material expires. Empty when pending key material import, otherwise `KEY_MATERIAL_EXPIRES` or `KEY_MATERIAL_DOES_NOT_EXPIRE`.
	ExpirationModel pulumi.StringPtrInput
	// The key ID of the replica key. Related multi-Region keys have the same key ID.
	KeyId pulumi.StringPtrInput
	// Base64 encoded 256-bit symmetric encryption key material to import. The KMS key is permanently associated with this key material. The same key material can be [reimported](https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html#reimport-key-material), but you cannot import different key material.
	KeyMaterialBase64 pulumi.StringPtrInput
	// The state of the replica key.
	KeyState pulumi.StringPtrInput
	// The [cryptographic operations](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations) for which you can use the KMS key. This is a shared property of multi-Region keys.
	KeyUsage pulumi.StringPtrInput
	// The key policy to attach to the KMS key. If you do not specify a key policy, AWS KMS attaches the [default key policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default) to the KMS key.
	Policy pulumi.StringPtrInput
	// The ARN of the multi-Region primary key to replicate. The primary key must be in a different AWS Region of the same AWS Partition. You can create only one replica of a given primary key in each AWS Region.
	PrimaryKeyArn pulumi.StringPtrInput
	// A map of tags to assign to the replica key. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
	// Time at which the imported key material expires. When the key material expires, AWS KMS deletes the key material and the key becomes unusable. If not specified, key material does not expire. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
	ValidTo pulumi.StringPtrInput
}

func (ReplicaExternalKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicaExternalKeyState)(nil)).Elem()
}

type replicaExternalKeyArgs struct {
	// A flag to indicate whether to bypass the key policy lockout safety check.
	// Setting this value to true increases the risk that the KMS key becomes unmanageable. Do not set this value to true indiscriminately.
	// For more information, refer to the scenario in the [Default Key Policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam) section in the _AWS Key Management Service Developer Guide_.
	// The default value is `false`.
	BypassPolicyLockoutSafetyCheck *bool `pulumi:"bypassPolicyLockoutSafetyCheck"`
	// The waiting period, specified in number of days. After the waiting period ends, AWS KMS deletes the KMS key.
	// If you specify a value, it must be between `7` and `30`, inclusive. If you do not specify a value, it defaults to `30`.
	DeletionWindowInDays *int `pulumi:"deletionWindowInDays"`
	// A description of the KMS key.
	Description *string `pulumi:"description"`
	// Specifies whether the replica key is enabled. Disabled KMS keys cannot be used in cryptographic operations. Keys pending import can only be `false`. Imported keys default to `true` unless expired.
	Enabled *bool `pulumi:"enabled"`
	// Base64 encoded 256-bit symmetric encryption key material to import. The KMS key is permanently associated with this key material. The same key material can be [reimported](https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html#reimport-key-material), but you cannot import different key material.
	KeyMaterialBase64 *string `pulumi:"keyMaterialBase64"`
	// The key policy to attach to the KMS key. If you do not specify a key policy, AWS KMS attaches the [default key policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default) to the KMS key.
	Policy *string `pulumi:"policy"`
	// The ARN of the multi-Region primary key to replicate. The primary key must be in a different AWS Region of the same AWS Partition. You can create only one replica of a given primary key in each AWS Region.
	PrimaryKeyArn string `pulumi:"primaryKeyArn"`
	// A map of tags to assign to the replica key. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Time at which the imported key material expires. When the key material expires, AWS KMS deletes the key material and the key becomes unusable. If not specified, key material does not expire. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
	ValidTo *string `pulumi:"validTo"`
}

// The set of arguments for constructing a ReplicaExternalKey resource.
type ReplicaExternalKeyArgs struct {
	// A flag to indicate whether to bypass the key policy lockout safety check.
	// Setting this value to true increases the risk that the KMS key becomes unmanageable. Do not set this value to true indiscriminately.
	// For more information, refer to the scenario in the [Default Key Policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam) section in the _AWS Key Management Service Developer Guide_.
	// The default value is `false`.
	BypassPolicyLockoutSafetyCheck pulumi.BoolPtrInput
	// The waiting period, specified in number of days. After the waiting period ends, AWS KMS deletes the KMS key.
	// If you specify a value, it must be between `7` and `30`, inclusive. If you do not specify a value, it defaults to `30`.
	DeletionWindowInDays pulumi.IntPtrInput
	// A description of the KMS key.
	Description pulumi.StringPtrInput
	// Specifies whether the replica key is enabled. Disabled KMS keys cannot be used in cryptographic operations. Keys pending import can only be `false`. Imported keys default to `true` unless expired.
	Enabled pulumi.BoolPtrInput
	// Base64 encoded 256-bit symmetric encryption key material to import. The KMS key is permanently associated with this key material. The same key material can be [reimported](https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html#reimport-key-material), but you cannot import different key material.
	KeyMaterialBase64 pulumi.StringPtrInput
	// The key policy to attach to the KMS key. If you do not specify a key policy, AWS KMS attaches the [default key policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default) to the KMS key.
	Policy pulumi.StringPtrInput
	// The ARN of the multi-Region primary key to replicate. The primary key must be in a different AWS Region of the same AWS Partition. You can create only one replica of a given primary key in each AWS Region.
	PrimaryKeyArn pulumi.StringInput
	// A map of tags to assign to the replica key. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Time at which the imported key material expires. When the key material expires, AWS KMS deletes the key material and the key becomes unusable. If not specified, key material does not expire. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
	ValidTo pulumi.StringPtrInput
}

func (ReplicaExternalKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicaExternalKeyArgs)(nil)).Elem()
}

type ReplicaExternalKeyInput interface {
	pulumi.Input

	ToReplicaExternalKeyOutput() ReplicaExternalKeyOutput
	ToReplicaExternalKeyOutputWithContext(ctx context.Context) ReplicaExternalKeyOutput
}

func (*ReplicaExternalKey) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicaExternalKey)(nil)).Elem()
}

func (i *ReplicaExternalKey) ToReplicaExternalKeyOutput() ReplicaExternalKeyOutput {
	return i.ToReplicaExternalKeyOutputWithContext(context.Background())
}

func (i *ReplicaExternalKey) ToReplicaExternalKeyOutputWithContext(ctx context.Context) ReplicaExternalKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicaExternalKeyOutput)
}

// ReplicaExternalKeyArrayInput is an input type that accepts ReplicaExternalKeyArray and ReplicaExternalKeyArrayOutput values.
// You can construct a concrete instance of `ReplicaExternalKeyArrayInput` via:
//
//	ReplicaExternalKeyArray{ ReplicaExternalKeyArgs{...} }
type ReplicaExternalKeyArrayInput interface {
	pulumi.Input

	ToReplicaExternalKeyArrayOutput() ReplicaExternalKeyArrayOutput
	ToReplicaExternalKeyArrayOutputWithContext(context.Context) ReplicaExternalKeyArrayOutput
}

type ReplicaExternalKeyArray []ReplicaExternalKeyInput

func (ReplicaExternalKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReplicaExternalKey)(nil)).Elem()
}

func (i ReplicaExternalKeyArray) ToReplicaExternalKeyArrayOutput() ReplicaExternalKeyArrayOutput {
	return i.ToReplicaExternalKeyArrayOutputWithContext(context.Background())
}

func (i ReplicaExternalKeyArray) ToReplicaExternalKeyArrayOutputWithContext(ctx context.Context) ReplicaExternalKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicaExternalKeyArrayOutput)
}

// ReplicaExternalKeyMapInput is an input type that accepts ReplicaExternalKeyMap and ReplicaExternalKeyMapOutput values.
// You can construct a concrete instance of `ReplicaExternalKeyMapInput` via:
//
//	ReplicaExternalKeyMap{ "key": ReplicaExternalKeyArgs{...} }
type ReplicaExternalKeyMapInput interface {
	pulumi.Input

	ToReplicaExternalKeyMapOutput() ReplicaExternalKeyMapOutput
	ToReplicaExternalKeyMapOutputWithContext(context.Context) ReplicaExternalKeyMapOutput
}

type ReplicaExternalKeyMap map[string]ReplicaExternalKeyInput

func (ReplicaExternalKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReplicaExternalKey)(nil)).Elem()
}

func (i ReplicaExternalKeyMap) ToReplicaExternalKeyMapOutput() ReplicaExternalKeyMapOutput {
	return i.ToReplicaExternalKeyMapOutputWithContext(context.Background())
}

func (i ReplicaExternalKeyMap) ToReplicaExternalKeyMapOutputWithContext(ctx context.Context) ReplicaExternalKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicaExternalKeyMapOutput)
}

type ReplicaExternalKeyOutput struct{ *pulumi.OutputState }

func (ReplicaExternalKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicaExternalKey)(nil)).Elem()
}

func (o ReplicaExternalKeyOutput) ToReplicaExternalKeyOutput() ReplicaExternalKeyOutput {
	return o
}

func (o ReplicaExternalKeyOutput) ToReplicaExternalKeyOutputWithContext(ctx context.Context) ReplicaExternalKeyOutput {
	return o
}

// The Amazon Resource Name (ARN) of the replica key. The key ARNs of related multi-Region keys differ only in the Region value.
func (o ReplicaExternalKeyOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicaExternalKey) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// A flag to indicate whether to bypass the key policy lockout safety check.
// Setting this value to true increases the risk that the KMS key becomes unmanageable. Do not set this value to true indiscriminately.
// For more information, refer to the scenario in the [Default Key Policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam) section in the _AWS Key Management Service Developer Guide_.
// The default value is `false`.
func (o ReplicaExternalKeyOutput) BypassPolicyLockoutSafetyCheck() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ReplicaExternalKey) pulumi.BoolPtrOutput { return v.BypassPolicyLockoutSafetyCheck }).(pulumi.BoolPtrOutput)
}

// The waiting period, specified in number of days. After the waiting period ends, AWS KMS deletes the KMS key.
// If you specify a value, it must be between `7` and `30`, inclusive. If you do not specify a value, it defaults to `30`.
func (o ReplicaExternalKeyOutput) DeletionWindowInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ReplicaExternalKey) pulumi.IntPtrOutput { return v.DeletionWindowInDays }).(pulumi.IntPtrOutput)
}

// A description of the KMS key.
func (o ReplicaExternalKeyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicaExternalKey) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Specifies whether the replica key is enabled. Disabled KMS keys cannot be used in cryptographic operations. Keys pending import can only be `false`. Imported keys default to `true` unless expired.
func (o ReplicaExternalKeyOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ReplicaExternalKey) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// Whether the key material expires. Empty when pending key material import, otherwise `KEY_MATERIAL_EXPIRES` or `KEY_MATERIAL_DOES_NOT_EXPIRE`.
func (o ReplicaExternalKeyOutput) ExpirationModel() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicaExternalKey) pulumi.StringOutput { return v.ExpirationModel }).(pulumi.StringOutput)
}

// The key ID of the replica key. Related multi-Region keys have the same key ID.
func (o ReplicaExternalKeyOutput) KeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicaExternalKey) pulumi.StringOutput { return v.KeyId }).(pulumi.StringOutput)
}

// Base64 encoded 256-bit symmetric encryption key material to import. The KMS key is permanently associated with this key material. The same key material can be [reimported](https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html#reimport-key-material), but you cannot import different key material.
func (o ReplicaExternalKeyOutput) KeyMaterialBase64() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicaExternalKey) pulumi.StringPtrOutput { return v.KeyMaterialBase64 }).(pulumi.StringPtrOutput)
}

// The state of the replica key.
func (o ReplicaExternalKeyOutput) KeyState() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicaExternalKey) pulumi.StringOutput { return v.KeyState }).(pulumi.StringOutput)
}

// The [cryptographic operations](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations) for which you can use the KMS key. This is a shared property of multi-Region keys.
func (o ReplicaExternalKeyOutput) KeyUsage() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicaExternalKey) pulumi.StringOutput { return v.KeyUsage }).(pulumi.StringOutput)
}

// The key policy to attach to the KMS key. If you do not specify a key policy, AWS KMS attaches the [default key policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default) to the KMS key.
func (o ReplicaExternalKeyOutput) Policy() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicaExternalKey) pulumi.StringOutput { return v.Policy }).(pulumi.StringOutput)
}

// The ARN of the multi-Region primary key to replicate. The primary key must be in a different AWS Region of the same AWS Partition. You can create only one replica of a given primary key in each AWS Region.
func (o ReplicaExternalKeyOutput) PrimaryKeyArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicaExternalKey) pulumi.StringOutput { return v.PrimaryKeyArn }).(pulumi.StringOutput)
}

// A map of tags to assign to the replica key. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o ReplicaExternalKeyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ReplicaExternalKey) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o ReplicaExternalKeyOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ReplicaExternalKey) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// Time at which the imported key material expires. When the key material expires, AWS KMS deletes the key material and the key becomes unusable. If not specified, key material does not expire. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
func (o ReplicaExternalKeyOutput) ValidTo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicaExternalKey) pulumi.StringPtrOutput { return v.ValidTo }).(pulumi.StringPtrOutput)
}

type ReplicaExternalKeyArrayOutput struct{ *pulumi.OutputState }

func (ReplicaExternalKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReplicaExternalKey)(nil)).Elem()
}

func (o ReplicaExternalKeyArrayOutput) ToReplicaExternalKeyArrayOutput() ReplicaExternalKeyArrayOutput {
	return o
}

func (o ReplicaExternalKeyArrayOutput) ToReplicaExternalKeyArrayOutputWithContext(ctx context.Context) ReplicaExternalKeyArrayOutput {
	return o
}

func (o ReplicaExternalKeyArrayOutput) Index(i pulumi.IntInput) ReplicaExternalKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ReplicaExternalKey {
		return vs[0].([]*ReplicaExternalKey)[vs[1].(int)]
	}).(ReplicaExternalKeyOutput)
}

type ReplicaExternalKeyMapOutput struct{ *pulumi.OutputState }

func (ReplicaExternalKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReplicaExternalKey)(nil)).Elem()
}

func (o ReplicaExternalKeyMapOutput) ToReplicaExternalKeyMapOutput() ReplicaExternalKeyMapOutput {
	return o
}

func (o ReplicaExternalKeyMapOutput) ToReplicaExternalKeyMapOutputWithContext(ctx context.Context) ReplicaExternalKeyMapOutput {
	return o
}

func (o ReplicaExternalKeyMapOutput) MapIndex(k pulumi.StringInput) ReplicaExternalKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ReplicaExternalKey {
		return vs[0].(map[string]*ReplicaExternalKey)[vs[1].(string)]
	}).(ReplicaExternalKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicaExternalKeyInput)(nil)).Elem(), &ReplicaExternalKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicaExternalKeyArrayInput)(nil)).Elem(), ReplicaExternalKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicaExternalKeyMapInput)(nil)).Elem(), ReplicaExternalKeyMap{})
	pulumi.RegisterOutputType(ReplicaExternalKeyOutput{})
	pulumi.RegisterOutputType(ReplicaExternalKeyArrayOutput{})
	pulumi.RegisterOutputType(ReplicaExternalKeyMapOutput{})
}
