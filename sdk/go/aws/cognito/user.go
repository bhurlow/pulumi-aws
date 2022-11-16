// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cognito

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Cognito User Resource.
//
// ## Example Usage
// ### Basic configuration
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/cognito"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleUserPool, err := cognito.NewUserPool(ctx, "exampleUserPool", nil)
//			if err != nil {
//				return err
//			}
//			_, err = cognito.NewUser(ctx, "exampleUser", &cognito.UserArgs{
//				UserPoolId: exampleUserPool.ID(),
//				Username:   pulumi.String("example"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Setting user attributes
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/cognito"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleUserPool, err := cognito.NewUserPool(ctx, "exampleUserPool", &cognito.UserPoolArgs{
//				Schemas: cognito.UserPoolSchemaArray{
//					&cognito.UserPoolSchemaArgs{
//						Name:                   pulumi.String("example"),
//						AttributeDataType:      pulumi.String("Boolean"),
//						Mutable:                pulumi.Bool(false),
//						Required:               pulumi.Bool(false),
//						DeveloperOnlyAttribute: pulumi.Bool(false),
//					},
//					&cognito.UserPoolSchemaArgs{
//						Name:                       pulumi.String("foo"),
//						AttributeDataType:          pulumi.String("String"),
//						Mutable:                    pulumi.Bool(false),
//						Required:                   pulumi.Bool(false),
//						DeveloperOnlyAttribute:     pulumi.Bool(false),
//						StringAttributeConstraints: nil,
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cognito.NewUser(ctx, "exampleUser", &cognito.UserArgs{
//				UserPoolId: exampleUserPool.ID(),
//				Username:   pulumi.String("example"),
//				Attributes: pulumi.StringMap{
//					"example":        pulumi.String("true"),
//					"foo":            pulumi.String("bar"),
//					"email":          pulumi.String("no-reply@domain.example"),
//					"email_verified": pulumi.String("true"),
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
// Cognito User can be imported using the `user_pool_id`/`name` attributes concatenated, e.g.,
//
// ```sh
//
//	$ pulumi import aws:cognito/user:User user us-east-1_vG78M4goG/user
//
// ```
type User struct {
	pulumi.CustomResourceState

	// A map that contains user attributes and attribute values to be set for the user.
	Attributes pulumi.StringMapOutput `pulumi:"attributes"`
	// A map of custom key-value pairs that you can provide as input for any custom workflows that user creation triggers. Amazon Cognito does not store the `clientMetadata` value. This data is available only to Lambda triggers that are assigned to a user pool to support custom workflows. If your user pool configuration does not include triggers, the ClientMetadata parameter serves no purpose. For more information, see [Customizing User Pool Workflows with Lambda Triggers](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html).
	ClientMetadata pulumi.StringMapOutput `pulumi:"clientMetadata"`
	CreationDate   pulumi.StringOutput    `pulumi:"creationDate"`
	// A list of mediums to the welcome message will be sent through. Allowed values are `EMAIL` and `SMS`. If it's provided, make sure you have also specified `email` attribute for the `EMAIL` medium and `phoneNumber` for the `SMS`. More than one value can be specified. Amazon Cognito does not store the `desiredDeliveryMediums` value. Defaults to `["SMS"]`.
	DesiredDeliveryMediums pulumi.StringArrayOutput `pulumi:"desiredDeliveryMediums"`
	// Specifies whether the user should be enabled after creation. The welcome message will be sent regardless of the `enabled` value. The behavior can be changed with `messageAction` argument. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// If this parameter is set to True and the `phoneNumber` or `email` address specified in the `attributes` parameter already exists as an alias with a different user, Amazon Cognito will migrate the alias from the previous user to the newly created user. The previous user will no longer be able to log in using that alias. Amazon Cognito does not store the `forceAliasCreation` value. Defaults to `false`.
	ForceAliasCreation pulumi.BoolPtrOutput `pulumi:"forceAliasCreation"`
	LastModifiedDate   pulumi.StringOutput  `pulumi:"lastModifiedDate"`
	// Set to `RESEND` to resend the invitation message to a user that already exists and reset the expiration limit on the user's account. Set to `SUPPRESS` to suppress sending the message. Only one value can be specified. Amazon Cognito does not store the `messageAction` value.
	MessageAction   pulumi.StringPtrOutput   `pulumi:"messageAction"`
	MfaSettingLists pulumi.StringArrayOutput `pulumi:"mfaSettingLists"`
	// The user's permanent password. This password must conform to the password policy specified by user pool the user belongs to. The welcome message always contains only `temporaryPassword` value. You can suppress sending the welcome message with the `messageAction` argument. Amazon Cognito does not store the `password` value. Conflicts with `temporaryPassword`.
	Password            pulumi.StringPtrOutput `pulumi:"password"`
	PreferredMfaSetting pulumi.StringOutput    `pulumi:"preferredMfaSetting"`
	// current user status.
	Status pulumi.StringOutput `pulumi:"status"`
	// unique user id that is never reassignable to another user.
	Sub pulumi.StringOutput `pulumi:"sub"`
	// The user's temporary password. Conflicts with `password`.
	TemporaryPassword pulumi.StringPtrOutput `pulumi:"temporaryPassword"`
	// The user pool ID for the user pool where the user will be created.
	UserPoolId pulumi.StringOutput `pulumi:"userPoolId"`
	Username   pulumi.StringOutput `pulumi:"username"`
	// The user's validation data. This is an array of name-value pairs that contain user attributes and attribute values that you can use for custom validation, such as restricting the types of user accounts that can be registered. Amazon Cognito does not store the `validationData` value. For more information, see [Customizing User Pool Workflows with Lambda Triggers](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html).
	ValidationData pulumi.StringMapOutput `pulumi:"validationData"`
}

// NewUser registers a new resource with the given unique name, arguments, and options.
func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOption) (*User, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.UserPoolId == nil {
		return nil, errors.New("invalid value for required argument 'UserPoolId'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	var resource User
	err := ctx.RegisterResource("aws:cognito/user:User", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUser gets an existing User resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserState, opts ...pulumi.ResourceOption) (*User, error) {
	var resource User
	err := ctx.ReadResource("aws:cognito/user:User", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering User resources.
type userState struct {
	// A map that contains user attributes and attribute values to be set for the user.
	Attributes map[string]string `pulumi:"attributes"`
	// A map of custom key-value pairs that you can provide as input for any custom workflows that user creation triggers. Amazon Cognito does not store the `clientMetadata` value. This data is available only to Lambda triggers that are assigned to a user pool to support custom workflows. If your user pool configuration does not include triggers, the ClientMetadata parameter serves no purpose. For more information, see [Customizing User Pool Workflows with Lambda Triggers](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html).
	ClientMetadata map[string]string `pulumi:"clientMetadata"`
	CreationDate   *string           `pulumi:"creationDate"`
	// A list of mediums to the welcome message will be sent through. Allowed values are `EMAIL` and `SMS`. If it's provided, make sure you have also specified `email` attribute for the `EMAIL` medium and `phoneNumber` for the `SMS`. More than one value can be specified. Amazon Cognito does not store the `desiredDeliveryMediums` value. Defaults to `["SMS"]`.
	DesiredDeliveryMediums []string `pulumi:"desiredDeliveryMediums"`
	// Specifies whether the user should be enabled after creation. The welcome message will be sent regardless of the `enabled` value. The behavior can be changed with `messageAction` argument. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// If this parameter is set to True and the `phoneNumber` or `email` address specified in the `attributes` parameter already exists as an alias with a different user, Amazon Cognito will migrate the alias from the previous user to the newly created user. The previous user will no longer be able to log in using that alias. Amazon Cognito does not store the `forceAliasCreation` value. Defaults to `false`.
	ForceAliasCreation *bool   `pulumi:"forceAliasCreation"`
	LastModifiedDate   *string `pulumi:"lastModifiedDate"`
	// Set to `RESEND` to resend the invitation message to a user that already exists and reset the expiration limit on the user's account. Set to `SUPPRESS` to suppress sending the message. Only one value can be specified. Amazon Cognito does not store the `messageAction` value.
	MessageAction   *string  `pulumi:"messageAction"`
	MfaSettingLists []string `pulumi:"mfaSettingLists"`
	// The user's permanent password. This password must conform to the password policy specified by user pool the user belongs to. The welcome message always contains only `temporaryPassword` value. You can suppress sending the welcome message with the `messageAction` argument. Amazon Cognito does not store the `password` value. Conflicts with `temporaryPassword`.
	Password            *string `pulumi:"password"`
	PreferredMfaSetting *string `pulumi:"preferredMfaSetting"`
	// current user status.
	Status *string `pulumi:"status"`
	// unique user id that is never reassignable to another user.
	Sub *string `pulumi:"sub"`
	// The user's temporary password. Conflicts with `password`.
	TemporaryPassword *string `pulumi:"temporaryPassword"`
	// The user pool ID for the user pool where the user will be created.
	UserPoolId *string `pulumi:"userPoolId"`
	Username   *string `pulumi:"username"`
	// The user's validation data. This is an array of name-value pairs that contain user attributes and attribute values that you can use for custom validation, such as restricting the types of user accounts that can be registered. Amazon Cognito does not store the `validationData` value. For more information, see [Customizing User Pool Workflows with Lambda Triggers](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html).
	ValidationData map[string]string `pulumi:"validationData"`
}

type UserState struct {
	// A map that contains user attributes and attribute values to be set for the user.
	Attributes pulumi.StringMapInput
	// A map of custom key-value pairs that you can provide as input for any custom workflows that user creation triggers. Amazon Cognito does not store the `clientMetadata` value. This data is available only to Lambda triggers that are assigned to a user pool to support custom workflows. If your user pool configuration does not include triggers, the ClientMetadata parameter serves no purpose. For more information, see [Customizing User Pool Workflows with Lambda Triggers](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html).
	ClientMetadata pulumi.StringMapInput
	CreationDate   pulumi.StringPtrInput
	// A list of mediums to the welcome message will be sent through. Allowed values are `EMAIL` and `SMS`. If it's provided, make sure you have also specified `email` attribute for the `EMAIL` medium and `phoneNumber` for the `SMS`. More than one value can be specified. Amazon Cognito does not store the `desiredDeliveryMediums` value. Defaults to `["SMS"]`.
	DesiredDeliveryMediums pulumi.StringArrayInput
	// Specifies whether the user should be enabled after creation. The welcome message will be sent regardless of the `enabled` value. The behavior can be changed with `messageAction` argument. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// If this parameter is set to True and the `phoneNumber` or `email` address specified in the `attributes` parameter already exists as an alias with a different user, Amazon Cognito will migrate the alias from the previous user to the newly created user. The previous user will no longer be able to log in using that alias. Amazon Cognito does not store the `forceAliasCreation` value. Defaults to `false`.
	ForceAliasCreation pulumi.BoolPtrInput
	LastModifiedDate   pulumi.StringPtrInput
	// Set to `RESEND` to resend the invitation message to a user that already exists and reset the expiration limit on the user's account. Set to `SUPPRESS` to suppress sending the message. Only one value can be specified. Amazon Cognito does not store the `messageAction` value.
	MessageAction   pulumi.StringPtrInput
	MfaSettingLists pulumi.StringArrayInput
	// The user's permanent password. This password must conform to the password policy specified by user pool the user belongs to. The welcome message always contains only `temporaryPassword` value. You can suppress sending the welcome message with the `messageAction` argument. Amazon Cognito does not store the `password` value. Conflicts with `temporaryPassword`.
	Password            pulumi.StringPtrInput
	PreferredMfaSetting pulumi.StringPtrInput
	// current user status.
	Status pulumi.StringPtrInput
	// unique user id that is never reassignable to another user.
	Sub pulumi.StringPtrInput
	// The user's temporary password. Conflicts with `password`.
	TemporaryPassword pulumi.StringPtrInput
	// The user pool ID for the user pool where the user will be created.
	UserPoolId pulumi.StringPtrInput
	Username   pulumi.StringPtrInput
	// The user's validation data. This is an array of name-value pairs that contain user attributes and attribute values that you can use for custom validation, such as restricting the types of user accounts that can be registered. Amazon Cognito does not store the `validationData` value. For more information, see [Customizing User Pool Workflows with Lambda Triggers](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html).
	ValidationData pulumi.StringMapInput
}

func (UserState) ElementType() reflect.Type {
	return reflect.TypeOf((*userState)(nil)).Elem()
}

type userArgs struct {
	// A map that contains user attributes and attribute values to be set for the user.
	Attributes map[string]string `pulumi:"attributes"`
	// A map of custom key-value pairs that you can provide as input for any custom workflows that user creation triggers. Amazon Cognito does not store the `clientMetadata` value. This data is available only to Lambda triggers that are assigned to a user pool to support custom workflows. If your user pool configuration does not include triggers, the ClientMetadata parameter serves no purpose. For more information, see [Customizing User Pool Workflows with Lambda Triggers](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html).
	ClientMetadata map[string]string `pulumi:"clientMetadata"`
	// A list of mediums to the welcome message will be sent through. Allowed values are `EMAIL` and `SMS`. If it's provided, make sure you have also specified `email` attribute for the `EMAIL` medium and `phoneNumber` for the `SMS`. More than one value can be specified. Amazon Cognito does not store the `desiredDeliveryMediums` value. Defaults to `["SMS"]`.
	DesiredDeliveryMediums []string `pulumi:"desiredDeliveryMediums"`
	// Specifies whether the user should be enabled after creation. The welcome message will be sent regardless of the `enabled` value. The behavior can be changed with `messageAction` argument. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// If this parameter is set to True and the `phoneNumber` or `email` address specified in the `attributes` parameter already exists as an alias with a different user, Amazon Cognito will migrate the alias from the previous user to the newly created user. The previous user will no longer be able to log in using that alias. Amazon Cognito does not store the `forceAliasCreation` value. Defaults to `false`.
	ForceAliasCreation *bool `pulumi:"forceAliasCreation"`
	// Set to `RESEND` to resend the invitation message to a user that already exists and reset the expiration limit on the user's account. Set to `SUPPRESS` to suppress sending the message. Only one value can be specified. Amazon Cognito does not store the `messageAction` value.
	MessageAction *string `pulumi:"messageAction"`
	// The user's permanent password. This password must conform to the password policy specified by user pool the user belongs to. The welcome message always contains only `temporaryPassword` value. You can suppress sending the welcome message with the `messageAction` argument. Amazon Cognito does not store the `password` value. Conflicts with `temporaryPassword`.
	Password *string `pulumi:"password"`
	// The user's temporary password. Conflicts with `password`.
	TemporaryPassword *string `pulumi:"temporaryPassword"`
	// The user pool ID for the user pool where the user will be created.
	UserPoolId string `pulumi:"userPoolId"`
	Username   string `pulumi:"username"`
	// The user's validation data. This is an array of name-value pairs that contain user attributes and attribute values that you can use for custom validation, such as restricting the types of user accounts that can be registered. Amazon Cognito does not store the `validationData` value. For more information, see [Customizing User Pool Workflows with Lambda Triggers](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html).
	ValidationData map[string]string `pulumi:"validationData"`
}

// The set of arguments for constructing a User resource.
type UserArgs struct {
	// A map that contains user attributes and attribute values to be set for the user.
	Attributes pulumi.StringMapInput
	// A map of custom key-value pairs that you can provide as input for any custom workflows that user creation triggers. Amazon Cognito does not store the `clientMetadata` value. This data is available only to Lambda triggers that are assigned to a user pool to support custom workflows. If your user pool configuration does not include triggers, the ClientMetadata parameter serves no purpose. For more information, see [Customizing User Pool Workflows with Lambda Triggers](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html).
	ClientMetadata pulumi.StringMapInput
	// A list of mediums to the welcome message will be sent through. Allowed values are `EMAIL` and `SMS`. If it's provided, make sure you have also specified `email` attribute for the `EMAIL` medium and `phoneNumber` for the `SMS`. More than one value can be specified. Amazon Cognito does not store the `desiredDeliveryMediums` value. Defaults to `["SMS"]`.
	DesiredDeliveryMediums pulumi.StringArrayInput
	// Specifies whether the user should be enabled after creation. The welcome message will be sent regardless of the `enabled` value. The behavior can be changed with `messageAction` argument. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// If this parameter is set to True and the `phoneNumber` or `email` address specified in the `attributes` parameter already exists as an alias with a different user, Amazon Cognito will migrate the alias from the previous user to the newly created user. The previous user will no longer be able to log in using that alias. Amazon Cognito does not store the `forceAliasCreation` value. Defaults to `false`.
	ForceAliasCreation pulumi.BoolPtrInput
	// Set to `RESEND` to resend the invitation message to a user that already exists and reset the expiration limit on the user's account. Set to `SUPPRESS` to suppress sending the message. Only one value can be specified. Amazon Cognito does not store the `messageAction` value.
	MessageAction pulumi.StringPtrInput
	// The user's permanent password. This password must conform to the password policy specified by user pool the user belongs to. The welcome message always contains only `temporaryPassword` value. You can suppress sending the welcome message with the `messageAction` argument. Amazon Cognito does not store the `password` value. Conflicts with `temporaryPassword`.
	Password pulumi.StringPtrInput
	// The user's temporary password. Conflicts with `password`.
	TemporaryPassword pulumi.StringPtrInput
	// The user pool ID for the user pool where the user will be created.
	UserPoolId pulumi.StringInput
	Username   pulumi.StringInput
	// The user's validation data. This is an array of name-value pairs that contain user attributes and attribute values that you can use for custom validation, such as restricting the types of user accounts that can be registered. Amazon Cognito does not store the `validationData` value. For more information, see [Customizing User Pool Workflows with Lambda Triggers](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html).
	ValidationData pulumi.StringMapInput
}

func (UserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userArgs)(nil)).Elem()
}

type UserInput interface {
	pulumi.Input

	ToUserOutput() UserOutput
	ToUserOutputWithContext(ctx context.Context) UserOutput
}

func (*User) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil)).Elem()
}

func (i *User) ToUserOutput() UserOutput {
	return i.ToUserOutputWithContext(context.Background())
}

func (i *User) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserOutput)
}

// UserArrayInput is an input type that accepts UserArray and UserArrayOutput values.
// You can construct a concrete instance of `UserArrayInput` via:
//
//	UserArray{ UserArgs{...} }
type UserArrayInput interface {
	pulumi.Input

	ToUserArrayOutput() UserArrayOutput
	ToUserArrayOutputWithContext(context.Context) UserArrayOutput
}

type UserArray []UserInput

func (UserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*User)(nil)).Elem()
}

func (i UserArray) ToUserArrayOutput() UserArrayOutput {
	return i.ToUserArrayOutputWithContext(context.Background())
}

func (i UserArray) ToUserArrayOutputWithContext(ctx context.Context) UserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserArrayOutput)
}

// UserMapInput is an input type that accepts UserMap and UserMapOutput values.
// You can construct a concrete instance of `UserMapInput` via:
//
//	UserMap{ "key": UserArgs{...} }
type UserMapInput interface {
	pulumi.Input

	ToUserMapOutput() UserMapOutput
	ToUserMapOutputWithContext(context.Context) UserMapOutput
}

type UserMap map[string]UserInput

func (UserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*User)(nil)).Elem()
}

func (i UserMap) ToUserMapOutput() UserMapOutput {
	return i.ToUserMapOutputWithContext(context.Background())
}

func (i UserMap) ToUserMapOutputWithContext(ctx context.Context) UserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserMapOutput)
}

type UserOutput struct{ *pulumi.OutputState }

func (UserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil)).Elem()
}

func (o UserOutput) ToUserOutput() UserOutput {
	return o
}

func (o UserOutput) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return o
}

// A map that contains user attributes and attribute values to be set for the user.
func (o UserOutput) Attributes() pulumi.StringMapOutput {
	return o.ApplyT(func(v *User) pulumi.StringMapOutput { return v.Attributes }).(pulumi.StringMapOutput)
}

// A map of custom key-value pairs that you can provide as input for any custom workflows that user creation triggers. Amazon Cognito does not store the `clientMetadata` value. This data is available only to Lambda triggers that are assigned to a user pool to support custom workflows. If your user pool configuration does not include triggers, the ClientMetadata parameter serves no purpose. For more information, see [Customizing User Pool Workflows with Lambda Triggers](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html).
func (o UserOutput) ClientMetadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *User) pulumi.StringMapOutput { return v.ClientMetadata }).(pulumi.StringMapOutput)
}

func (o UserOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

// A list of mediums to the welcome message will be sent through. Allowed values are `EMAIL` and `SMS`. If it's provided, make sure you have also specified `email` attribute for the `EMAIL` medium and `phoneNumber` for the `SMS`. More than one value can be specified. Amazon Cognito does not store the `desiredDeliveryMediums` value. Defaults to `["SMS"]`.
func (o UserOutput) DesiredDeliveryMediums() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *User) pulumi.StringArrayOutput { return v.DesiredDeliveryMediums }).(pulumi.StringArrayOutput)
}

// Specifies whether the user should be enabled after creation. The welcome message will be sent regardless of the `enabled` value. The behavior can be changed with `messageAction` argument. Defaults to `true`.
func (o UserOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *User) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// If this parameter is set to True and the `phoneNumber` or `email` address specified in the `attributes` parameter already exists as an alias with a different user, Amazon Cognito will migrate the alias from the previous user to the newly created user. The previous user will no longer be able to log in using that alias. Amazon Cognito does not store the `forceAliasCreation` value. Defaults to `false`.
func (o UserOutput) ForceAliasCreation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *User) pulumi.BoolPtrOutput { return v.ForceAliasCreation }).(pulumi.BoolPtrOutput)
}

func (o UserOutput) LastModifiedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.LastModifiedDate }).(pulumi.StringOutput)
}

// Set to `RESEND` to resend the invitation message to a user that already exists and reset the expiration limit on the user's account. Set to `SUPPRESS` to suppress sending the message. Only one value can be specified. Amazon Cognito does not store the `messageAction` value.
func (o UserOutput) MessageAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.MessageAction }).(pulumi.StringPtrOutput)
}

func (o UserOutput) MfaSettingLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *User) pulumi.StringArrayOutput { return v.MfaSettingLists }).(pulumi.StringArrayOutput)
}

// The user's permanent password. This password must conform to the password policy specified by user pool the user belongs to. The welcome message always contains only `temporaryPassword` value. You can suppress sending the welcome message with the `messageAction` argument. Amazon Cognito does not store the `password` value. Conflicts with `temporaryPassword`.
func (o UserOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

func (o UserOutput) PreferredMfaSetting() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.PreferredMfaSetting }).(pulumi.StringOutput)
}

// current user status.
func (o UserOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// unique user id that is never reassignable to another user.
func (o UserOutput) Sub() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Sub }).(pulumi.StringOutput)
}

// The user's temporary password. Conflicts with `password`.
func (o UserOutput) TemporaryPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.TemporaryPassword }).(pulumi.StringPtrOutput)
}

// The user pool ID for the user pool where the user will be created.
func (o UserOutput) UserPoolId() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.UserPoolId }).(pulumi.StringOutput)
}

func (o UserOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

// The user's validation data. This is an array of name-value pairs that contain user attributes and attribute values that you can use for custom validation, such as restricting the types of user accounts that can be registered. Amazon Cognito does not store the `validationData` value. For more information, see [Customizing User Pool Workflows with Lambda Triggers](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html).
func (o UserOutput) ValidationData() pulumi.StringMapOutput {
	return o.ApplyT(func(v *User) pulumi.StringMapOutput { return v.ValidationData }).(pulumi.StringMapOutput)
}

type UserArrayOutput struct{ *pulumi.OutputState }

func (UserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*User)(nil)).Elem()
}

func (o UserArrayOutput) ToUserArrayOutput() UserArrayOutput {
	return o
}

func (o UserArrayOutput) ToUserArrayOutputWithContext(ctx context.Context) UserArrayOutput {
	return o
}

func (o UserArrayOutput) Index(i pulumi.IntInput) UserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *User {
		return vs[0].([]*User)[vs[1].(int)]
	}).(UserOutput)
}

type UserMapOutput struct{ *pulumi.OutputState }

func (UserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*User)(nil)).Elem()
}

func (o UserMapOutput) ToUserMapOutput() UserMapOutput {
	return o
}

func (o UserMapOutput) ToUserMapOutputWithContext(ctx context.Context) UserMapOutput {
	return o
}

func (o UserMapOutput) MapIndex(k pulumi.StringInput) UserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *User {
		return vs[0].(map[string]*User)[vs[1].(string)]
	}).(UserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserInput)(nil)).Elem(), &User{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserArrayInput)(nil)).Elem(), UserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserMapInput)(nil)).Elem(), UserMap{})
	pulumi.RegisterOutputType(UserOutput{})
	pulumi.RegisterOutputType(UserArrayOutput{})
	pulumi.RegisterOutputType(UserMapOutput{})
}
