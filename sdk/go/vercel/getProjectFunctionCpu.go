// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vercel

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-vercel/sdk/go/vercel/internal"
)

// > This data source has been deprecated and no longer works. Please use the `Project` data source and its `resourceConfig` attribute instead.
//
// Provides information about a Project's Function CPU setting.
//
// This controls the maximum amount of CPU utilization your Serverless Functions can use while executing. Standard is optimal for most frontend workloads. You can override this per function using the vercel.json file.
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
//			exampleProject, err := vercel.LookupProject(ctx, &vercel.LookupProjectArgs{
//				Name: "example",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = vercel.LookupProjectFunctionCpu(ctx, &vercel.LookupProjectFunctionCpuArgs{
//				ProjectId: exampleProject.Id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupProjectFunctionCpu(ctx *pulumi.Context, args *LookupProjectFunctionCpuArgs, opts ...pulumi.InvokeOption) (*LookupProjectFunctionCpuResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupProjectFunctionCpuResult
	err := ctx.Invoke("vercel:index/getProjectFunctionCpu:getProjectFunctionCpu", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProjectFunctionCpu.
type LookupProjectFunctionCpuArgs struct {
	// The ID of the Project to read the Function CPU setting for.
	ProjectId string `pulumi:"projectId"`
	// The ID of the team the Project exists under. Required when configuring a team resource if a default team has not been set in the provider.
	TeamId *string `pulumi:"teamId"`
}

// A collection of values returned by getProjectFunctionCpu.
type LookupProjectFunctionCpuResult struct {
	// The amount of CPU available to your Serverless Functions. Should be one of 'basic' (0.6vCPU), 'standard' (1vCPU) or 'performance' (1.7vCPUs).
	Cpu string `pulumi:"cpu"`
	// The ID of the resource.
	Id string `pulumi:"id"`
	// The ID of the Project to read the Function CPU setting for.
	ProjectId string `pulumi:"projectId"`
	// The ID of the team the Project exists under. Required when configuring a team resource if a default team has not been set in the provider.
	TeamId string `pulumi:"teamId"`
}

func LookupProjectFunctionCpuOutput(ctx *pulumi.Context, args LookupProjectFunctionCpuOutputArgs, opts ...pulumi.InvokeOption) LookupProjectFunctionCpuResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProjectFunctionCpuResultOutput, error) {
			args := v.(LookupProjectFunctionCpuArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupProjectFunctionCpuResult
			secret, err := ctx.InvokePackageRaw("vercel:index/getProjectFunctionCpu:getProjectFunctionCpu", args, &rv, "", opts...)
			if err != nil {
				return LookupProjectFunctionCpuResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupProjectFunctionCpuResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupProjectFunctionCpuResultOutput), nil
			}
			return output, nil
		}).(LookupProjectFunctionCpuResultOutput)
}

// A collection of arguments for invoking getProjectFunctionCpu.
type LookupProjectFunctionCpuOutputArgs struct {
	// The ID of the Project to read the Function CPU setting for.
	ProjectId pulumi.StringInput `pulumi:"projectId"`
	// The ID of the team the Project exists under. Required when configuring a team resource if a default team has not been set in the provider.
	TeamId pulumi.StringPtrInput `pulumi:"teamId"`
}

func (LookupProjectFunctionCpuOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectFunctionCpuArgs)(nil)).Elem()
}

// A collection of values returned by getProjectFunctionCpu.
type LookupProjectFunctionCpuResultOutput struct{ *pulumi.OutputState }

func (LookupProjectFunctionCpuResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectFunctionCpuResult)(nil)).Elem()
}

func (o LookupProjectFunctionCpuResultOutput) ToLookupProjectFunctionCpuResultOutput() LookupProjectFunctionCpuResultOutput {
	return o
}

func (o LookupProjectFunctionCpuResultOutput) ToLookupProjectFunctionCpuResultOutputWithContext(ctx context.Context) LookupProjectFunctionCpuResultOutput {
	return o
}

// The amount of CPU available to your Serverless Functions. Should be one of 'basic' (0.6vCPU), 'standard' (1vCPU) or 'performance' (1.7vCPUs).
func (o LookupProjectFunctionCpuResultOutput) Cpu() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectFunctionCpuResult) string { return v.Cpu }).(pulumi.StringOutput)
}

// The ID of the resource.
func (o LookupProjectFunctionCpuResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectFunctionCpuResult) string { return v.Id }).(pulumi.StringOutput)
}

// The ID of the Project to read the Function CPU setting for.
func (o LookupProjectFunctionCpuResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectFunctionCpuResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// The ID of the team the Project exists under. Required when configuring a team resource if a default team has not been set in the provider.
func (o LookupProjectFunctionCpuResultOutput) TeamId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectFunctionCpuResult) string { return v.TeamId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProjectFunctionCpuResultOutput{})
}
