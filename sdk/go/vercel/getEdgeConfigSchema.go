// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vercel

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-vercel/sdk/go/vercel/internal"
)

// An Edge Config Schema provides an existing Edge Config with a JSON schema. Use schema protection to prevent unexpected updates that may cause bugs or downtime.
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
//			_, err := vercel.LookupEdgeConfigSchema(ctx, &vercel.LookupEdgeConfigSchemaArgs{
//				Id: "ecfg_xxxxxxxxxxxxxxxxxxxxxxxxxxxx",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupEdgeConfigSchema(ctx *pulumi.Context, args *LookupEdgeConfigSchemaArgs, opts ...pulumi.InvokeOption) (*LookupEdgeConfigSchemaResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupEdgeConfigSchemaResult
	err := ctx.Invoke("vercel:index/getEdgeConfigSchema:getEdgeConfigSchema", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEdgeConfigSchema.
type LookupEdgeConfigSchemaArgs struct {
	// The ID of the Edge Config that the schema should be for.
	Id string `pulumi:"id"`
	// The ID of the team the Edge Config should exist under. Required when configuring a team resource if a default team has not been set in the provider.
	TeamId *string `pulumi:"teamId"`
}

// A collection of values returned by getEdgeConfigSchema.
type LookupEdgeConfigSchemaResult struct {
	// A JSON schema that will be used to validate data in the Edge Config.
	Definition string `pulumi:"definition"`
	// The ID of the Edge Config that the schema should be for.
	Id string `pulumi:"id"`
	// The ID of the team the Edge Config should exist under. Required when configuring a team resource if a default team has not been set in the provider.
	TeamId string `pulumi:"teamId"`
}

func LookupEdgeConfigSchemaOutput(ctx *pulumi.Context, args LookupEdgeConfigSchemaOutputArgs, opts ...pulumi.InvokeOption) LookupEdgeConfigSchemaResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEdgeConfigSchemaResult, error) {
			args := v.(LookupEdgeConfigSchemaArgs)
			r, err := LookupEdgeConfigSchema(ctx, &args, opts...)
			var s LookupEdgeConfigSchemaResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEdgeConfigSchemaResultOutput)
}

// A collection of arguments for invoking getEdgeConfigSchema.
type LookupEdgeConfigSchemaOutputArgs struct {
	// The ID of the Edge Config that the schema should be for.
	Id pulumi.StringInput `pulumi:"id"`
	// The ID of the team the Edge Config should exist under. Required when configuring a team resource if a default team has not been set in the provider.
	TeamId pulumi.StringPtrInput `pulumi:"teamId"`
}

func (LookupEdgeConfigSchemaOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEdgeConfigSchemaArgs)(nil)).Elem()
}

// A collection of values returned by getEdgeConfigSchema.
type LookupEdgeConfigSchemaResultOutput struct{ *pulumi.OutputState }

func (LookupEdgeConfigSchemaResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEdgeConfigSchemaResult)(nil)).Elem()
}

func (o LookupEdgeConfigSchemaResultOutput) ToLookupEdgeConfigSchemaResultOutput() LookupEdgeConfigSchemaResultOutput {
	return o
}

func (o LookupEdgeConfigSchemaResultOutput) ToLookupEdgeConfigSchemaResultOutputWithContext(ctx context.Context) LookupEdgeConfigSchemaResultOutput {
	return o
}

// A JSON schema that will be used to validate data in the Edge Config.
func (o LookupEdgeConfigSchemaResultOutput) Definition() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEdgeConfigSchemaResult) string { return v.Definition }).(pulumi.StringOutput)
}

// The ID of the Edge Config that the schema should be for.
func (o LookupEdgeConfigSchemaResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEdgeConfigSchemaResult) string { return v.Id }).(pulumi.StringOutput)
}

// The ID of the team the Edge Config should exist under. Required when configuring a team resource if a default team has not been set in the provider.
func (o LookupEdgeConfigSchemaResultOutput) TeamId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEdgeConfigSchemaResult) string { return v.TeamId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEdgeConfigSchemaResultOutput{})
}