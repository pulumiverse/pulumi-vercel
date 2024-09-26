// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vercel

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-vercel/sdk/go/vercel/internal"
)

// Provides a Project Deployment Retention resource.
//
// A Project Deployment Retention resource defines an Deployment Retention on a Vercel Project.
//
// For more detailed information, please see the [Vercel documentation](https://vercel.com/docs/security/deployment-retention).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-vercel/sdk/go/vercel"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := vercel.NewProject(ctx, "example", &vercel.ProjectArgs{
//				GitRepository: &vercel.ProjectGitRepositoryArgs{
//					Type: pulumi.String("github"),
//					Repo: pulumi.String("vercel/some-repo"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			// An unlimited deployment retention policy that will be created
//			// for this project for all deployments.
//			_, err = vercel.NewProjectDeploymentRetention(ctx, "exampleUnlimited", &vercel.ProjectDeploymentRetentionArgs{
//				ProjectId:            example.ID(),
//				TeamId:               example.TeamId,
//				ExpirationPreview:    pulumi.String("unlimited"),
//				ExpirationProduction: pulumi.String("unlimited"),
//				ExpirationCanceled:   pulumi.String("unlimited"),
//				ExpirationErrored:    pulumi.String("unlimited"),
//			})
//			if err != nil {
//				return err
//			}
//			// A customized deployment retention policy that will be created
//			// for this project for all deployments.
//			_, err = vercel.NewProjectDeploymentRetention(ctx, "exampleCustomized", &vercel.ProjectDeploymentRetentionArgs{
//				ProjectId:            example.ID(),
//				TeamId:               example.TeamId,
//				ExpirationPreview:    pulumi.String("3m"),
//				ExpirationProduction: pulumi.String("1y"),
//				ExpirationCanceled:   pulumi.String("1m"),
//				ExpirationErrored:    pulumi.String("2m"),
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
// You can import via the team_id and project_id.
//
// - team_id can be found in the team `settings` tab in the Vercel UI.
//
// - project_id can be found in the project `settings` tab in the Vercel UI.
//
// ```sh
// $ pulumi import vercel:index/projectDeploymentRetention:ProjectDeploymentRetention example team_xxxxxxxxxxxxxxxxxxxxxxxx/prj_xxxxxxxxxxxxxxxxxxxxxxxxxxxx
// ```
type ProjectDeploymentRetention struct {
	pulumi.CustomResourceState

	// The retention period for canceled deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
	ExpirationCanceled pulumi.StringOutput `pulumi:"expirationCanceled"`
	// The retention period for errored deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
	ExpirationErrored pulumi.StringOutput `pulumi:"expirationErrored"`
	// The retention period for preview deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
	ExpirationPreview pulumi.StringOutput `pulumi:"expirationPreview"`
	// The retention period for production deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
	ExpirationProduction pulumi.StringOutput `pulumi:"expirationProduction"`
	// The ID of the Project for the retention policy
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The ID of the Vercel team.
	TeamId pulumi.StringOutput `pulumi:"teamId"`
}

// NewProjectDeploymentRetention registers a new resource with the given unique name, arguments, and options.
func NewProjectDeploymentRetention(ctx *pulumi.Context,
	name string, args *ProjectDeploymentRetentionArgs, opts ...pulumi.ResourceOption) (*ProjectDeploymentRetention, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProjectDeploymentRetention
	err := ctx.RegisterResource("vercel:index/projectDeploymentRetention:ProjectDeploymentRetention", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectDeploymentRetention gets an existing ProjectDeploymentRetention resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectDeploymentRetention(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectDeploymentRetentionState, opts ...pulumi.ResourceOption) (*ProjectDeploymentRetention, error) {
	var resource ProjectDeploymentRetention
	err := ctx.ReadResource("vercel:index/projectDeploymentRetention:ProjectDeploymentRetention", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectDeploymentRetention resources.
type projectDeploymentRetentionState struct {
	// The retention period for canceled deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
	ExpirationCanceled *string `pulumi:"expirationCanceled"`
	// The retention period for errored deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
	ExpirationErrored *string `pulumi:"expirationErrored"`
	// The retention period for preview deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
	ExpirationPreview *string `pulumi:"expirationPreview"`
	// The retention period for production deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
	ExpirationProduction *string `pulumi:"expirationProduction"`
	// The ID of the Project for the retention policy
	ProjectId *string `pulumi:"projectId"`
	// The ID of the Vercel team.
	TeamId *string `pulumi:"teamId"`
}

type ProjectDeploymentRetentionState struct {
	// The retention period for canceled deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
	ExpirationCanceled pulumi.StringPtrInput
	// The retention period for errored deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
	ExpirationErrored pulumi.StringPtrInput
	// The retention period for preview deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
	ExpirationPreview pulumi.StringPtrInput
	// The retention period for production deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
	ExpirationProduction pulumi.StringPtrInput
	// The ID of the Project for the retention policy
	ProjectId pulumi.StringPtrInput
	// The ID of the Vercel team.
	TeamId pulumi.StringPtrInput
}

func (ProjectDeploymentRetentionState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectDeploymentRetentionState)(nil)).Elem()
}

type projectDeploymentRetentionArgs struct {
	// The retention period for canceled deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
	ExpirationCanceled *string `pulumi:"expirationCanceled"`
	// The retention period for errored deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
	ExpirationErrored *string `pulumi:"expirationErrored"`
	// The retention period for preview deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
	ExpirationPreview *string `pulumi:"expirationPreview"`
	// The retention period for production deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
	ExpirationProduction *string `pulumi:"expirationProduction"`
	// The ID of the Project for the retention policy
	ProjectId string `pulumi:"projectId"`
	// The ID of the Vercel team.
	TeamId *string `pulumi:"teamId"`
}

// The set of arguments for constructing a ProjectDeploymentRetention resource.
type ProjectDeploymentRetentionArgs struct {
	// The retention period for canceled deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
	ExpirationCanceled pulumi.StringPtrInput
	// The retention period for errored deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
	ExpirationErrored pulumi.StringPtrInput
	// The retention period for preview deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
	ExpirationPreview pulumi.StringPtrInput
	// The retention period for production deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
	ExpirationProduction pulumi.StringPtrInput
	// The ID of the Project for the retention policy
	ProjectId pulumi.StringInput
	// The ID of the Vercel team.
	TeamId pulumi.StringPtrInput
}

func (ProjectDeploymentRetentionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectDeploymentRetentionArgs)(nil)).Elem()
}

type ProjectDeploymentRetentionInput interface {
	pulumi.Input

	ToProjectDeploymentRetentionOutput() ProjectDeploymentRetentionOutput
	ToProjectDeploymentRetentionOutputWithContext(ctx context.Context) ProjectDeploymentRetentionOutput
}

func (*ProjectDeploymentRetention) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectDeploymentRetention)(nil)).Elem()
}

func (i *ProjectDeploymentRetention) ToProjectDeploymentRetentionOutput() ProjectDeploymentRetentionOutput {
	return i.ToProjectDeploymentRetentionOutputWithContext(context.Background())
}

func (i *ProjectDeploymentRetention) ToProjectDeploymentRetentionOutputWithContext(ctx context.Context) ProjectDeploymentRetentionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectDeploymentRetentionOutput)
}

// ProjectDeploymentRetentionArrayInput is an input type that accepts ProjectDeploymentRetentionArray and ProjectDeploymentRetentionArrayOutput values.
// You can construct a concrete instance of `ProjectDeploymentRetentionArrayInput` via:
//
//	ProjectDeploymentRetentionArray{ ProjectDeploymentRetentionArgs{...} }
type ProjectDeploymentRetentionArrayInput interface {
	pulumi.Input

	ToProjectDeploymentRetentionArrayOutput() ProjectDeploymentRetentionArrayOutput
	ToProjectDeploymentRetentionArrayOutputWithContext(context.Context) ProjectDeploymentRetentionArrayOutput
}

type ProjectDeploymentRetentionArray []ProjectDeploymentRetentionInput

func (ProjectDeploymentRetentionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectDeploymentRetention)(nil)).Elem()
}

func (i ProjectDeploymentRetentionArray) ToProjectDeploymentRetentionArrayOutput() ProjectDeploymentRetentionArrayOutput {
	return i.ToProjectDeploymentRetentionArrayOutputWithContext(context.Background())
}

func (i ProjectDeploymentRetentionArray) ToProjectDeploymentRetentionArrayOutputWithContext(ctx context.Context) ProjectDeploymentRetentionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectDeploymentRetentionArrayOutput)
}

// ProjectDeploymentRetentionMapInput is an input type that accepts ProjectDeploymentRetentionMap and ProjectDeploymentRetentionMapOutput values.
// You can construct a concrete instance of `ProjectDeploymentRetentionMapInput` via:
//
//	ProjectDeploymentRetentionMap{ "key": ProjectDeploymentRetentionArgs{...} }
type ProjectDeploymentRetentionMapInput interface {
	pulumi.Input

	ToProjectDeploymentRetentionMapOutput() ProjectDeploymentRetentionMapOutput
	ToProjectDeploymentRetentionMapOutputWithContext(context.Context) ProjectDeploymentRetentionMapOutput
}

type ProjectDeploymentRetentionMap map[string]ProjectDeploymentRetentionInput

func (ProjectDeploymentRetentionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectDeploymentRetention)(nil)).Elem()
}

func (i ProjectDeploymentRetentionMap) ToProjectDeploymentRetentionMapOutput() ProjectDeploymentRetentionMapOutput {
	return i.ToProjectDeploymentRetentionMapOutputWithContext(context.Background())
}

func (i ProjectDeploymentRetentionMap) ToProjectDeploymentRetentionMapOutputWithContext(ctx context.Context) ProjectDeploymentRetentionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectDeploymentRetentionMapOutput)
}

type ProjectDeploymentRetentionOutput struct{ *pulumi.OutputState }

func (ProjectDeploymentRetentionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectDeploymentRetention)(nil)).Elem()
}

func (o ProjectDeploymentRetentionOutput) ToProjectDeploymentRetentionOutput() ProjectDeploymentRetentionOutput {
	return o
}

func (o ProjectDeploymentRetentionOutput) ToProjectDeploymentRetentionOutputWithContext(ctx context.Context) ProjectDeploymentRetentionOutput {
	return o
}

// The retention period for canceled deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
func (o ProjectDeploymentRetentionOutput) ExpirationCanceled() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectDeploymentRetention) pulumi.StringOutput { return v.ExpirationCanceled }).(pulumi.StringOutput)
}

// The retention period for errored deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
func (o ProjectDeploymentRetentionOutput) ExpirationErrored() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectDeploymentRetention) pulumi.StringOutput { return v.ExpirationErrored }).(pulumi.StringOutput)
}

// The retention period for preview deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
func (o ProjectDeploymentRetentionOutput) ExpirationPreview() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectDeploymentRetention) pulumi.StringOutput { return v.ExpirationPreview }).(pulumi.StringOutput)
}

// The retention period for production deployments. Should be one of '1m', '2m', '3m', '6m', '1y', 'unlimited'.
func (o ProjectDeploymentRetentionOutput) ExpirationProduction() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectDeploymentRetention) pulumi.StringOutput { return v.ExpirationProduction }).(pulumi.StringOutput)
}

// The ID of the Project for the retention policy
func (o ProjectDeploymentRetentionOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectDeploymentRetention) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The ID of the Vercel team.
func (o ProjectDeploymentRetentionOutput) TeamId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectDeploymentRetention) pulumi.StringOutput { return v.TeamId }).(pulumi.StringOutput)
}

type ProjectDeploymentRetentionArrayOutput struct{ *pulumi.OutputState }

func (ProjectDeploymentRetentionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectDeploymentRetention)(nil)).Elem()
}

func (o ProjectDeploymentRetentionArrayOutput) ToProjectDeploymentRetentionArrayOutput() ProjectDeploymentRetentionArrayOutput {
	return o
}

func (o ProjectDeploymentRetentionArrayOutput) ToProjectDeploymentRetentionArrayOutputWithContext(ctx context.Context) ProjectDeploymentRetentionArrayOutput {
	return o
}

func (o ProjectDeploymentRetentionArrayOutput) Index(i pulumi.IntInput) ProjectDeploymentRetentionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProjectDeploymentRetention {
		return vs[0].([]*ProjectDeploymentRetention)[vs[1].(int)]
	}).(ProjectDeploymentRetentionOutput)
}

type ProjectDeploymentRetentionMapOutput struct{ *pulumi.OutputState }

func (ProjectDeploymentRetentionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectDeploymentRetention)(nil)).Elem()
}

func (o ProjectDeploymentRetentionMapOutput) ToProjectDeploymentRetentionMapOutput() ProjectDeploymentRetentionMapOutput {
	return o
}

func (o ProjectDeploymentRetentionMapOutput) ToProjectDeploymentRetentionMapOutputWithContext(ctx context.Context) ProjectDeploymentRetentionMapOutput {
	return o
}

func (o ProjectDeploymentRetentionMapOutput) MapIndex(k pulumi.StringInput) ProjectDeploymentRetentionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProjectDeploymentRetention {
		return vs[0].(map[string]*ProjectDeploymentRetention)[vs[1].(string)]
	}).(ProjectDeploymentRetentionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectDeploymentRetentionInput)(nil)).Elem(), &ProjectDeploymentRetention{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectDeploymentRetentionArrayInput)(nil)).Elem(), ProjectDeploymentRetentionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectDeploymentRetentionMapInput)(nil)).Elem(), ProjectDeploymentRetentionMap{})
	pulumi.RegisterOutputType(ProjectDeploymentRetentionOutput{})
	pulumi.RegisterOutputType(ProjectDeploymentRetentionArrayOutput{})
	pulumi.RegisterOutputType(ProjectDeploymentRetentionMapOutput{})
}