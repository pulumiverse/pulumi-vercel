// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vercel

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-vercel/sdk/v3/go/vercel/internal"
)

// Provides an Access Group Resource.
//
// Access Groups provide a way to manage groups of Vercel users across projects on your team. They are a set of project role assignations, a combination of Vercel users and the projects they work on.
//
// For more detailed information, please see the [Vercel documentation](https://vercel.com/docs/accounts/team-members-and-roles/access-groups).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-vercel/sdk/v3/go/vercel"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vercel.NewAccessGroup(ctx, "example", &vercel.AccessGroupArgs{
//				Name: pulumi.String("example-access-group"),
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
// # If importing into a personal account, or with a team configured on
//
// the provider, simply use the access_group_id.
//
// ```sh
// $ pulumi import vercel:index/accessGroup:AccessGroup example ag_xxxxxxxxxxxxxxxxxxxxxxxxxxxx
// ```
//
// If importing to a team, use the team_id and access_group_id.
//
// ```sh
// $ pulumi import vercel:index/accessGroup:AccessGroup example team_xxxxxxxxxxxxxxxxxxxxxxxx/ag_xxxxxxxxxxxxxxxxxxxxxxxxxxxx
// ```
type AccessGroup struct {
	pulumi.CustomResourceState

	// The name of the Access Group
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the team the Access Group should exist under. Required when configuring a team resource if a default team has not been set in the provider.
	TeamId pulumi.StringOutput `pulumi:"teamId"`
}

// NewAccessGroup registers a new resource with the given unique name, arguments, and options.
func NewAccessGroup(ctx *pulumi.Context,
	name string, args *AccessGroupArgs, opts ...pulumi.ResourceOption) (*AccessGroup, error) {
	if args == nil {
		args = &AccessGroupArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AccessGroup
	err := ctx.RegisterResource("vercel:index/accessGroup:AccessGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessGroup gets an existing AccessGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessGroupState, opts ...pulumi.ResourceOption) (*AccessGroup, error) {
	var resource AccessGroup
	err := ctx.ReadResource("vercel:index/accessGroup:AccessGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessGroup resources.
type accessGroupState struct {
	// The name of the Access Group
	Name *string `pulumi:"name"`
	// The ID of the team the Access Group should exist under. Required when configuring a team resource if a default team has not been set in the provider.
	TeamId *string `pulumi:"teamId"`
}

type AccessGroupState struct {
	// The name of the Access Group
	Name pulumi.StringPtrInput
	// The ID of the team the Access Group should exist under. Required when configuring a team resource if a default team has not been set in the provider.
	TeamId pulumi.StringPtrInput
}

func (AccessGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessGroupState)(nil)).Elem()
}

type accessGroupArgs struct {
	// The name of the Access Group
	Name *string `pulumi:"name"`
	// The ID of the team the Access Group should exist under. Required when configuring a team resource if a default team has not been set in the provider.
	TeamId *string `pulumi:"teamId"`
}

// The set of arguments for constructing a AccessGroup resource.
type AccessGroupArgs struct {
	// The name of the Access Group
	Name pulumi.StringPtrInput
	// The ID of the team the Access Group should exist under. Required when configuring a team resource if a default team has not been set in the provider.
	TeamId pulumi.StringPtrInput
}

func (AccessGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessGroupArgs)(nil)).Elem()
}

type AccessGroupInput interface {
	pulumi.Input

	ToAccessGroupOutput() AccessGroupOutput
	ToAccessGroupOutputWithContext(ctx context.Context) AccessGroupOutput
}

func (*AccessGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessGroup)(nil)).Elem()
}

func (i *AccessGroup) ToAccessGroupOutput() AccessGroupOutput {
	return i.ToAccessGroupOutputWithContext(context.Background())
}

func (i *AccessGroup) ToAccessGroupOutputWithContext(ctx context.Context) AccessGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessGroupOutput)
}

// AccessGroupArrayInput is an input type that accepts AccessGroupArray and AccessGroupArrayOutput values.
// You can construct a concrete instance of `AccessGroupArrayInput` via:
//
//	AccessGroupArray{ AccessGroupArgs{...} }
type AccessGroupArrayInput interface {
	pulumi.Input

	ToAccessGroupArrayOutput() AccessGroupArrayOutput
	ToAccessGroupArrayOutputWithContext(context.Context) AccessGroupArrayOutput
}

type AccessGroupArray []AccessGroupInput

func (AccessGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccessGroup)(nil)).Elem()
}

func (i AccessGroupArray) ToAccessGroupArrayOutput() AccessGroupArrayOutput {
	return i.ToAccessGroupArrayOutputWithContext(context.Background())
}

func (i AccessGroupArray) ToAccessGroupArrayOutputWithContext(ctx context.Context) AccessGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessGroupArrayOutput)
}

// AccessGroupMapInput is an input type that accepts AccessGroupMap and AccessGroupMapOutput values.
// You can construct a concrete instance of `AccessGroupMapInput` via:
//
//	AccessGroupMap{ "key": AccessGroupArgs{...} }
type AccessGroupMapInput interface {
	pulumi.Input

	ToAccessGroupMapOutput() AccessGroupMapOutput
	ToAccessGroupMapOutputWithContext(context.Context) AccessGroupMapOutput
}

type AccessGroupMap map[string]AccessGroupInput

func (AccessGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccessGroup)(nil)).Elem()
}

func (i AccessGroupMap) ToAccessGroupMapOutput() AccessGroupMapOutput {
	return i.ToAccessGroupMapOutputWithContext(context.Background())
}

func (i AccessGroupMap) ToAccessGroupMapOutputWithContext(ctx context.Context) AccessGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessGroupMapOutput)
}

type AccessGroupOutput struct{ *pulumi.OutputState }

func (AccessGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessGroup)(nil)).Elem()
}

func (o AccessGroupOutput) ToAccessGroupOutput() AccessGroupOutput {
	return o
}

func (o AccessGroupOutput) ToAccessGroupOutputWithContext(ctx context.Context) AccessGroupOutput {
	return o
}

// The name of the Access Group
func (o AccessGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the team the Access Group should exist under. Required when configuring a team resource if a default team has not been set in the provider.
func (o AccessGroupOutput) TeamId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessGroup) pulumi.StringOutput { return v.TeamId }).(pulumi.StringOutput)
}

type AccessGroupArrayOutput struct{ *pulumi.OutputState }

func (AccessGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccessGroup)(nil)).Elem()
}

func (o AccessGroupArrayOutput) ToAccessGroupArrayOutput() AccessGroupArrayOutput {
	return o
}

func (o AccessGroupArrayOutput) ToAccessGroupArrayOutputWithContext(ctx context.Context) AccessGroupArrayOutput {
	return o
}

func (o AccessGroupArrayOutput) Index(i pulumi.IntInput) AccessGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AccessGroup {
		return vs[0].([]*AccessGroup)[vs[1].(int)]
	}).(AccessGroupOutput)
}

type AccessGroupMapOutput struct{ *pulumi.OutputState }

func (AccessGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccessGroup)(nil)).Elem()
}

func (o AccessGroupMapOutput) ToAccessGroupMapOutput() AccessGroupMapOutput {
	return o
}

func (o AccessGroupMapOutput) ToAccessGroupMapOutputWithContext(ctx context.Context) AccessGroupMapOutput {
	return o
}

func (o AccessGroupMapOutput) MapIndex(k pulumi.StringInput) AccessGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AccessGroup {
		return vs[0].(map[string]*AccessGroup)[vs[1].(string)]
	}).(AccessGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccessGroupInput)(nil)).Elem(), &AccessGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessGroupArrayInput)(nil)).Elem(), AccessGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessGroupMapInput)(nil)).Elem(), AccessGroupMap{})
	pulumi.RegisterOutputType(AccessGroupOutput{})
	pulumi.RegisterOutputType(AccessGroupArrayOutput{})
	pulumi.RegisterOutputType(AccessGroupMapOutput{})
}
