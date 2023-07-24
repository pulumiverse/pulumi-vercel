// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vercel

import (
	"context"
	"reflect"

	"github.com/omercnet/pulumi-vercel/sdk/go/vercel/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetProjectDirectory(ctx *pulumi.Context, args *GetProjectDirectoryArgs, opts ...pulumi.InvokeOption) (*GetProjectDirectoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetProjectDirectoryResult
	err := ctx.Invoke("vercel:index/getProjectDirectory:getProjectDirectory", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProjectDirectory.
type GetProjectDirectoryArgs struct {
	Path string `pulumi:"path"`
}

// A collection of values returned by getProjectDirectory.
type GetProjectDirectoryResult struct {
	// A map of filename to metadata about the file. The metadata contains the file size and hash, and allows a deployment to be created if the file changes.
	Files map[string]string `pulumi:"files"`
	// The ID of this resource.
	Id   string `pulumi:"id"`
	Path string `pulumi:"path"`
}

func GetProjectDirectoryOutput(ctx *pulumi.Context, args GetProjectDirectoryOutputArgs, opts ...pulumi.InvokeOption) GetProjectDirectoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetProjectDirectoryResult, error) {
			args := v.(GetProjectDirectoryArgs)
			r, err := GetProjectDirectory(ctx, &args, opts...)
			var s GetProjectDirectoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetProjectDirectoryResultOutput)
}

// A collection of arguments for invoking getProjectDirectory.
type GetProjectDirectoryOutputArgs struct {
	Path pulumi.StringInput `pulumi:"path"`
}

func (GetProjectDirectoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectDirectoryArgs)(nil)).Elem()
}

// A collection of values returned by getProjectDirectory.
type GetProjectDirectoryResultOutput struct{ *pulumi.OutputState }

func (GetProjectDirectoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectDirectoryResult)(nil)).Elem()
}

func (o GetProjectDirectoryResultOutput) ToGetProjectDirectoryResultOutput() GetProjectDirectoryResultOutput {
	return o
}

func (o GetProjectDirectoryResultOutput) ToGetProjectDirectoryResultOutputWithContext(ctx context.Context) GetProjectDirectoryResultOutput {
	return o
}

// A map of filename to metadata about the file. The metadata contains the file size and hash, and allows a deployment to be created if the file changes.
func (o GetProjectDirectoryResultOutput) Files() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetProjectDirectoryResult) map[string]string { return v.Files }).(pulumi.StringMapOutput)
}

// The ID of this resource.
func (o GetProjectDirectoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectDirectoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetProjectDirectoryResultOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectDirectoryResult) string { return v.Path }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProjectDirectoryResultOutput{})
}
