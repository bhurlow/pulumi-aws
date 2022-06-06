// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redshift

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides redshift subnet group.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/redshift"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := redshift.GetClusterCredentials(ctx, &redshift.GetClusterCredentialsArgs{
// 			Name: aws_redshift_cluster_credentials.Example.Name,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetClusterCredentials(ctx *pulumi.Context, args *GetClusterCredentialsArgs, opts ...pulumi.InvokeOption) (*GetClusterCredentialsResult, error) {
	var rv GetClusterCredentialsResult
	err := ctx.Invoke("aws:redshift/getClusterCredentials:getClusterCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getClusterCredentials.
type GetClusterCredentialsArgs struct {
	// Create a database user with the name specified for the user named in `dbUser` if one does not exist.
	AutoCreate *bool `pulumi:"autoCreate"`
	// The unique identifier of the cluster that contains the database for which your are requesting credentials.
	ClusterIdentifier string `pulumi:"clusterIdentifier"`
	// A list of the names of existing database groups that the user named in `dbUser` will join for the current session, in addition to any group memberships for an existing user. If not specified, a new user is added only to `PUBLIC`.
	DbGroups []string `pulumi:"dbGroups"`
	// The name of a database that DbUser is authorized to log on to. If `dbName` is not specified, `dbUser` can log on to any existing database.
	DbName *string `pulumi:"dbName"`
	// The name of a database user. If a user name matching `dbUser` exists in the database, the temporary user credentials have the same permissions as the  existing user. If `dbUser` doesn't exist in the database and `autoCreate` is `True`, a new user is created using the value for `dbUser` with `PUBLIC` permissions.  If a database user matching the value for `dbUser` doesn't exist and `not` is `False`, then the command succeeds but the connection attempt will fail because the user doesn't exist in the database.
	DbUser string `pulumi:"dbUser"`
	// The number of seconds until the returned temporary password expires. Valid values are between `900` and `3600`. Default value is `900`.
	DurationSeconds *int `pulumi:"durationSeconds"`
}

// A collection of values returned by getClusterCredentials.
type GetClusterCredentialsResult struct {
	AutoCreate        *bool    `pulumi:"autoCreate"`
	ClusterIdentifier string   `pulumi:"clusterIdentifier"`
	DbGroups          []string `pulumi:"dbGroups"`
	DbName            *string  `pulumi:"dbName"`
	// A temporary password that authorizes the user name returned by `dbUser` to log on to the database `dbName`.
	DbPassword      string `pulumi:"dbPassword"`
	DbUser          string `pulumi:"dbUser"`
	DurationSeconds *int   `pulumi:"durationSeconds"`
	// The date and time the password in `dbPassword` expires.
	Expiration string `pulumi:"expiration"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}

func GetClusterCredentialsOutput(ctx *pulumi.Context, args GetClusterCredentialsOutputArgs, opts ...pulumi.InvokeOption) GetClusterCredentialsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetClusterCredentialsResult, error) {
			args := v.(GetClusterCredentialsArgs)
			r, err := GetClusterCredentials(ctx, &args, opts...)
			var s GetClusterCredentialsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetClusterCredentialsResultOutput)
}

// A collection of arguments for invoking getClusterCredentials.
type GetClusterCredentialsOutputArgs struct {
	// Create a database user with the name specified for the user named in `dbUser` if one does not exist.
	AutoCreate pulumi.BoolPtrInput `pulumi:"autoCreate"`
	// The unique identifier of the cluster that contains the database for which your are requesting credentials.
	ClusterIdentifier pulumi.StringInput `pulumi:"clusterIdentifier"`
	// A list of the names of existing database groups that the user named in `dbUser` will join for the current session, in addition to any group memberships for an existing user. If not specified, a new user is added only to `PUBLIC`.
	DbGroups pulumi.StringArrayInput `pulumi:"dbGroups"`
	// The name of a database that DbUser is authorized to log on to. If `dbName` is not specified, `dbUser` can log on to any existing database.
	DbName pulumi.StringPtrInput `pulumi:"dbName"`
	// The name of a database user. If a user name matching `dbUser` exists in the database, the temporary user credentials have the same permissions as the  existing user. If `dbUser` doesn't exist in the database and `autoCreate` is `True`, a new user is created using the value for `dbUser` with `PUBLIC` permissions.  If a database user matching the value for `dbUser` doesn't exist and `not` is `False`, then the command succeeds but the connection attempt will fail because the user doesn't exist in the database.
	DbUser pulumi.StringInput `pulumi:"dbUser"`
	// The number of seconds until the returned temporary password expires. Valid values are between `900` and `3600`. Default value is `900`.
	DurationSeconds pulumi.IntPtrInput `pulumi:"durationSeconds"`
}

func (GetClusterCredentialsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClusterCredentialsArgs)(nil)).Elem()
}

// A collection of values returned by getClusterCredentials.
type GetClusterCredentialsResultOutput struct{ *pulumi.OutputState }

func (GetClusterCredentialsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClusterCredentialsResult)(nil)).Elem()
}

func (o GetClusterCredentialsResultOutput) ToGetClusterCredentialsResultOutput() GetClusterCredentialsResultOutput {
	return o
}

func (o GetClusterCredentialsResultOutput) ToGetClusterCredentialsResultOutputWithContext(ctx context.Context) GetClusterCredentialsResultOutput {
	return o
}

func (o GetClusterCredentialsResultOutput) AutoCreate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetClusterCredentialsResult) *bool { return v.AutoCreate }).(pulumi.BoolPtrOutput)
}

func (o GetClusterCredentialsResultOutput) ClusterIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v GetClusterCredentialsResult) string { return v.ClusterIdentifier }).(pulumi.StringOutput)
}

func (o GetClusterCredentialsResultOutput) DbGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetClusterCredentialsResult) []string { return v.DbGroups }).(pulumi.StringArrayOutput)
}

func (o GetClusterCredentialsResultOutput) DbName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetClusterCredentialsResult) *string { return v.DbName }).(pulumi.StringPtrOutput)
}

// A temporary password that authorizes the user name returned by `dbUser` to log on to the database `dbName`.
func (o GetClusterCredentialsResultOutput) DbPassword() pulumi.StringOutput {
	return o.ApplyT(func(v GetClusterCredentialsResult) string { return v.DbPassword }).(pulumi.StringOutput)
}

func (o GetClusterCredentialsResultOutput) DbUser() pulumi.StringOutput {
	return o.ApplyT(func(v GetClusterCredentialsResult) string { return v.DbUser }).(pulumi.StringOutput)
}

func (o GetClusterCredentialsResultOutput) DurationSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetClusterCredentialsResult) *int { return v.DurationSeconds }).(pulumi.IntPtrOutput)
}

// The date and time the password in `dbPassword` expires.
func (o GetClusterCredentialsResultOutput) Expiration() pulumi.StringOutput {
	return o.ApplyT(func(v GetClusterCredentialsResult) string { return v.Expiration }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetClusterCredentialsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetClusterCredentialsResult) string { return v.Id }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetClusterCredentialsResultOutput{})
}
