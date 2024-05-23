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

// Provides a Function CPU resource for a Project.
//
// This controls the maximum amount of CPU utilization your Serverless Functions can use while executing. Standard is optimal for most frontend workloads. You can override this per function using the vercel.json file.
//
// A new Deployment is required for your changes to take effect.
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
//			exampleProject, err := vercel.NewProject(ctx, "exampleProject", nil)
//			if err != nil {
//				return err
//			}
//			_, err = vercel.NewProjectFunctionCpu(ctx, "exampleProjectFunctionCpu", &vercel.ProjectFunctionCpuArgs{
//				ProjectId: exampleProject.ID(),
//				Cpu:       pulumi.String("performance"),
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
// $ pulumi import vercel:index/projectFunctionCpu:ProjectFunctionCpu example team_xxxxxxxxxxxxxxxxxxxxxxxx/prj_xxxxxxxxxxxxxxxxxxxxxxxxxxxx
// ```
type ProjectFunctionCpu struct {
	pulumi.CustomResourceState

	// The amount of CPU available to your Serverless Functions. Should be one of 'basic' (0.6vCPU), 'standard' (1vCPU) or 'performance' (1.7vCPUs).
	Cpu pulumi.StringOutput `pulumi:"cpu"`
	// The ID of the Project to adjust the CPU for.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The ID of the team the Project exists under. Required when configuring a team resource if a default team has not been set in the provider.
	TeamId pulumi.StringOutput `pulumi:"teamId"`
}

// NewProjectFunctionCpu registers a new resource with the given unique name, arguments, and options.
func NewProjectFunctionCpu(ctx *pulumi.Context,
	name string, args *ProjectFunctionCpuArgs, opts ...pulumi.ResourceOption) (*ProjectFunctionCpu, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Cpu == nil {
		return nil, errors.New("invalid value for required argument 'Cpu'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProjectFunctionCpu
	err := ctx.RegisterResource("vercel:index/projectFunctionCpu:ProjectFunctionCpu", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectFunctionCpu gets an existing ProjectFunctionCpu resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectFunctionCpu(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectFunctionCpuState, opts ...pulumi.ResourceOption) (*ProjectFunctionCpu, error) {
	var resource ProjectFunctionCpu
	err := ctx.ReadResource("vercel:index/projectFunctionCpu:ProjectFunctionCpu", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectFunctionCpu resources.
type projectFunctionCpuState struct {
	// The amount of CPU available to your Serverless Functions. Should be one of 'basic' (0.6vCPU), 'standard' (1vCPU) or 'performance' (1.7vCPUs).
	Cpu *string `pulumi:"cpu"`
	// The ID of the Project to adjust the CPU for.
	ProjectId *string `pulumi:"projectId"`
	// The ID of the team the Project exists under. Required when configuring a team resource if a default team has not been set in the provider.
	TeamId *string `pulumi:"teamId"`
}

type ProjectFunctionCpuState struct {
	// The amount of CPU available to your Serverless Functions. Should be one of 'basic' (0.6vCPU), 'standard' (1vCPU) or 'performance' (1.7vCPUs).
	Cpu pulumi.StringPtrInput
	// The ID of the Project to adjust the CPU for.
	ProjectId pulumi.StringPtrInput
	// The ID of the team the Project exists under. Required when configuring a team resource if a default team has not been set in the provider.
	TeamId pulumi.StringPtrInput
}

func (ProjectFunctionCpuState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectFunctionCpuState)(nil)).Elem()
}

type projectFunctionCpuArgs struct {
	// The amount of CPU available to your Serverless Functions. Should be one of 'basic' (0.6vCPU), 'standard' (1vCPU) or 'performance' (1.7vCPUs).
	Cpu string `pulumi:"cpu"`
	// The ID of the Project to adjust the CPU for.
	ProjectId string `pulumi:"projectId"`
	// The ID of the team the Project exists under. Required when configuring a team resource if a default team has not been set in the provider.
	TeamId *string `pulumi:"teamId"`
}

// The set of arguments for constructing a ProjectFunctionCpu resource.
type ProjectFunctionCpuArgs struct {
	// The amount of CPU available to your Serverless Functions. Should be one of 'basic' (0.6vCPU), 'standard' (1vCPU) or 'performance' (1.7vCPUs).
	Cpu pulumi.StringInput
	// The ID of the Project to adjust the CPU for.
	ProjectId pulumi.StringInput
	// The ID of the team the Project exists under. Required when configuring a team resource if a default team has not been set in the provider.
	TeamId pulumi.StringPtrInput
}

func (ProjectFunctionCpuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectFunctionCpuArgs)(nil)).Elem()
}

type ProjectFunctionCpuInput interface {
	pulumi.Input

	ToProjectFunctionCpuOutput() ProjectFunctionCpuOutput
	ToProjectFunctionCpuOutputWithContext(ctx context.Context) ProjectFunctionCpuOutput
}

func (*ProjectFunctionCpu) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectFunctionCpu)(nil)).Elem()
}

func (i *ProjectFunctionCpu) ToProjectFunctionCpuOutput() ProjectFunctionCpuOutput {
	return i.ToProjectFunctionCpuOutputWithContext(context.Background())
}

func (i *ProjectFunctionCpu) ToProjectFunctionCpuOutputWithContext(ctx context.Context) ProjectFunctionCpuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectFunctionCpuOutput)
}

// ProjectFunctionCpuArrayInput is an input type that accepts ProjectFunctionCpuArray and ProjectFunctionCpuArrayOutput values.
// You can construct a concrete instance of `ProjectFunctionCpuArrayInput` via:
//
//	ProjectFunctionCpuArray{ ProjectFunctionCpuArgs{...} }
type ProjectFunctionCpuArrayInput interface {
	pulumi.Input

	ToProjectFunctionCpuArrayOutput() ProjectFunctionCpuArrayOutput
	ToProjectFunctionCpuArrayOutputWithContext(context.Context) ProjectFunctionCpuArrayOutput
}

type ProjectFunctionCpuArray []ProjectFunctionCpuInput

func (ProjectFunctionCpuArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectFunctionCpu)(nil)).Elem()
}

func (i ProjectFunctionCpuArray) ToProjectFunctionCpuArrayOutput() ProjectFunctionCpuArrayOutput {
	return i.ToProjectFunctionCpuArrayOutputWithContext(context.Background())
}

func (i ProjectFunctionCpuArray) ToProjectFunctionCpuArrayOutputWithContext(ctx context.Context) ProjectFunctionCpuArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectFunctionCpuArrayOutput)
}

// ProjectFunctionCpuMapInput is an input type that accepts ProjectFunctionCpuMap and ProjectFunctionCpuMapOutput values.
// You can construct a concrete instance of `ProjectFunctionCpuMapInput` via:
//
//	ProjectFunctionCpuMap{ "key": ProjectFunctionCpuArgs{...} }
type ProjectFunctionCpuMapInput interface {
	pulumi.Input

	ToProjectFunctionCpuMapOutput() ProjectFunctionCpuMapOutput
	ToProjectFunctionCpuMapOutputWithContext(context.Context) ProjectFunctionCpuMapOutput
}

type ProjectFunctionCpuMap map[string]ProjectFunctionCpuInput

func (ProjectFunctionCpuMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectFunctionCpu)(nil)).Elem()
}

func (i ProjectFunctionCpuMap) ToProjectFunctionCpuMapOutput() ProjectFunctionCpuMapOutput {
	return i.ToProjectFunctionCpuMapOutputWithContext(context.Background())
}

func (i ProjectFunctionCpuMap) ToProjectFunctionCpuMapOutputWithContext(ctx context.Context) ProjectFunctionCpuMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectFunctionCpuMapOutput)
}

type ProjectFunctionCpuOutput struct{ *pulumi.OutputState }

func (ProjectFunctionCpuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectFunctionCpu)(nil)).Elem()
}

func (o ProjectFunctionCpuOutput) ToProjectFunctionCpuOutput() ProjectFunctionCpuOutput {
	return o
}

func (o ProjectFunctionCpuOutput) ToProjectFunctionCpuOutputWithContext(ctx context.Context) ProjectFunctionCpuOutput {
	return o
}

// The amount of CPU available to your Serverless Functions. Should be one of 'basic' (0.6vCPU), 'standard' (1vCPU) or 'performance' (1.7vCPUs).
func (o ProjectFunctionCpuOutput) Cpu() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectFunctionCpu) pulumi.StringOutput { return v.Cpu }).(pulumi.StringOutput)
}

// The ID of the Project to adjust the CPU for.
func (o ProjectFunctionCpuOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectFunctionCpu) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The ID of the team the Project exists under. Required when configuring a team resource if a default team has not been set in the provider.
func (o ProjectFunctionCpuOutput) TeamId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectFunctionCpu) pulumi.StringOutput { return v.TeamId }).(pulumi.StringOutput)
}

type ProjectFunctionCpuArrayOutput struct{ *pulumi.OutputState }

func (ProjectFunctionCpuArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectFunctionCpu)(nil)).Elem()
}

func (o ProjectFunctionCpuArrayOutput) ToProjectFunctionCpuArrayOutput() ProjectFunctionCpuArrayOutput {
	return o
}

func (o ProjectFunctionCpuArrayOutput) ToProjectFunctionCpuArrayOutputWithContext(ctx context.Context) ProjectFunctionCpuArrayOutput {
	return o
}

func (o ProjectFunctionCpuArrayOutput) Index(i pulumi.IntInput) ProjectFunctionCpuOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProjectFunctionCpu {
		return vs[0].([]*ProjectFunctionCpu)[vs[1].(int)]
	}).(ProjectFunctionCpuOutput)
}

type ProjectFunctionCpuMapOutput struct{ *pulumi.OutputState }

func (ProjectFunctionCpuMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectFunctionCpu)(nil)).Elem()
}

func (o ProjectFunctionCpuMapOutput) ToProjectFunctionCpuMapOutput() ProjectFunctionCpuMapOutput {
	return o
}

func (o ProjectFunctionCpuMapOutput) ToProjectFunctionCpuMapOutputWithContext(ctx context.Context) ProjectFunctionCpuMapOutput {
	return o
}

func (o ProjectFunctionCpuMapOutput) MapIndex(k pulumi.StringInput) ProjectFunctionCpuOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProjectFunctionCpu {
		return vs[0].(map[string]*ProjectFunctionCpu)[vs[1].(string)]
	}).(ProjectFunctionCpuOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectFunctionCpuInput)(nil)).Elem(), &ProjectFunctionCpu{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectFunctionCpuArrayInput)(nil)).Elem(), ProjectFunctionCpuArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectFunctionCpuMapInput)(nil)).Elem(), ProjectFunctionCpuMap{})
	pulumi.RegisterOutputType(ProjectFunctionCpuOutput{})
	pulumi.RegisterOutputType(ProjectFunctionCpuArrayOutput{})
	pulumi.RegisterOutputType(ProjectFunctionCpuMapOutput{})
}