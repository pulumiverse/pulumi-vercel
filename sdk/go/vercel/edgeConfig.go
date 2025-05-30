// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vercel

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-vercel/sdk/v3/go/vercel/internal"
)

// Provides an Edge Config resource.
//
// An Edge Config is a global data store that enables experimentation with feature flags, A/B testing, critical redirects, and more.
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
//			example, err := vercel.NewEdgeConfig(ctx, "example", &vercel.EdgeConfigArgs{
//				Name: pulumi.String("example"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleProject, err := vercel.NewProject(ctx, "example", &vercel.ProjectArgs{
//				Name: pulumi.String("edge-config-example"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleEdgeConfigToken, err := vercel.NewEdgeConfigToken(ctx, "example", &vercel.EdgeConfigTokenArgs{
//				EdgeConfigId: example.ID(),
//				Label:        pulumi.String("example token"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vercel.NewProjectEnvironmentVariable(ctx, "example", &vercel.ProjectEnvironmentVariableArgs{
//				ProjectId: exampleProject.ID(),
//				Targets: pulumi.StringArray{
//					pulumi.String("production"),
//					pulumi.String("preview"),
//					pulumi.String("development"),
//				},
//				Key:   pulumi.String("EDGE_CONFIG"),
//				Value: exampleEdgeConfigToken.ConnectionString,
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
// the provider, simply use the edge config id.
//
// - edge_config_id can be found by navigating to the Edge Config in the Vercel UI. It should begin with `ecfg_`.
//
// ```sh
// $ pulumi import vercel:index/edgeConfig:EdgeConfig example ecfg_xxxxxxxxxxxxxxxxxxxxxxxxxxxx
// ```
//
// Alternatively, you can import via the team_id and edge_config_id.
//
// - team_id can be found in the team `settings` tab in the Vercel UI.
//
// - edge_config_id can be found by navigating to the Edge Config in the Vercel UI. It should begin with `ecfg_`.
//
// ```sh
// $ pulumi import vercel:index/edgeConfig:EdgeConfig example team_xxxxxxxxxxxxxxxxxxxxxxxx/ecfg_xxxxxxxxxxxxxxxxxxxxxxxxxxxx
// ```
type EdgeConfig struct {
	pulumi.CustomResourceState

	// The name/slug of the Edge Config.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the team the Edge Config should exist under. Required when configuring a team resource if a default team has not been set in the provider.
	TeamId pulumi.StringOutput `pulumi:"teamId"`
}

// NewEdgeConfig registers a new resource with the given unique name, arguments, and options.
func NewEdgeConfig(ctx *pulumi.Context,
	name string, args *EdgeConfigArgs, opts ...pulumi.ResourceOption) (*EdgeConfig, error) {
	if args == nil {
		args = &EdgeConfigArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EdgeConfig
	err := ctx.RegisterResource("vercel:index/edgeConfig:EdgeConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEdgeConfig gets an existing EdgeConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEdgeConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EdgeConfigState, opts ...pulumi.ResourceOption) (*EdgeConfig, error) {
	var resource EdgeConfig
	err := ctx.ReadResource("vercel:index/edgeConfig:EdgeConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EdgeConfig resources.
type edgeConfigState struct {
	// The name/slug of the Edge Config.
	Name *string `pulumi:"name"`
	// The ID of the team the Edge Config should exist under. Required when configuring a team resource if a default team has not been set in the provider.
	TeamId *string `pulumi:"teamId"`
}

type EdgeConfigState struct {
	// The name/slug of the Edge Config.
	Name pulumi.StringPtrInput
	// The ID of the team the Edge Config should exist under. Required when configuring a team resource if a default team has not been set in the provider.
	TeamId pulumi.StringPtrInput
}

func (EdgeConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*edgeConfigState)(nil)).Elem()
}

type edgeConfigArgs struct {
	// The name/slug of the Edge Config.
	Name *string `pulumi:"name"`
	// The ID of the team the Edge Config should exist under. Required when configuring a team resource if a default team has not been set in the provider.
	TeamId *string `pulumi:"teamId"`
}

// The set of arguments for constructing a EdgeConfig resource.
type EdgeConfigArgs struct {
	// The name/slug of the Edge Config.
	Name pulumi.StringPtrInput
	// The ID of the team the Edge Config should exist under. Required when configuring a team resource if a default team has not been set in the provider.
	TeamId pulumi.StringPtrInput
}

func (EdgeConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*edgeConfigArgs)(nil)).Elem()
}

type EdgeConfigInput interface {
	pulumi.Input

	ToEdgeConfigOutput() EdgeConfigOutput
	ToEdgeConfigOutputWithContext(ctx context.Context) EdgeConfigOutput
}

func (*EdgeConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**EdgeConfig)(nil)).Elem()
}

func (i *EdgeConfig) ToEdgeConfigOutput() EdgeConfigOutput {
	return i.ToEdgeConfigOutputWithContext(context.Background())
}

func (i *EdgeConfig) ToEdgeConfigOutputWithContext(ctx context.Context) EdgeConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeConfigOutput)
}

// EdgeConfigArrayInput is an input type that accepts EdgeConfigArray and EdgeConfigArrayOutput values.
// You can construct a concrete instance of `EdgeConfigArrayInput` via:
//
//	EdgeConfigArray{ EdgeConfigArgs{...} }
type EdgeConfigArrayInput interface {
	pulumi.Input

	ToEdgeConfigArrayOutput() EdgeConfigArrayOutput
	ToEdgeConfigArrayOutputWithContext(context.Context) EdgeConfigArrayOutput
}

type EdgeConfigArray []EdgeConfigInput

func (EdgeConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EdgeConfig)(nil)).Elem()
}

func (i EdgeConfigArray) ToEdgeConfigArrayOutput() EdgeConfigArrayOutput {
	return i.ToEdgeConfigArrayOutputWithContext(context.Background())
}

func (i EdgeConfigArray) ToEdgeConfigArrayOutputWithContext(ctx context.Context) EdgeConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeConfigArrayOutput)
}

// EdgeConfigMapInput is an input type that accepts EdgeConfigMap and EdgeConfigMapOutput values.
// You can construct a concrete instance of `EdgeConfigMapInput` via:
//
//	EdgeConfigMap{ "key": EdgeConfigArgs{...} }
type EdgeConfigMapInput interface {
	pulumi.Input

	ToEdgeConfigMapOutput() EdgeConfigMapOutput
	ToEdgeConfigMapOutputWithContext(context.Context) EdgeConfigMapOutput
}

type EdgeConfigMap map[string]EdgeConfigInput

func (EdgeConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EdgeConfig)(nil)).Elem()
}

func (i EdgeConfigMap) ToEdgeConfigMapOutput() EdgeConfigMapOutput {
	return i.ToEdgeConfigMapOutputWithContext(context.Background())
}

func (i EdgeConfigMap) ToEdgeConfigMapOutputWithContext(ctx context.Context) EdgeConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeConfigMapOutput)
}

type EdgeConfigOutput struct{ *pulumi.OutputState }

func (EdgeConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdgeConfig)(nil)).Elem()
}

func (o EdgeConfigOutput) ToEdgeConfigOutput() EdgeConfigOutput {
	return o
}

func (o EdgeConfigOutput) ToEdgeConfigOutputWithContext(ctx context.Context) EdgeConfigOutput {
	return o
}

// The name/slug of the Edge Config.
func (o EdgeConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EdgeConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the team the Edge Config should exist under. Required when configuring a team resource if a default team has not been set in the provider.
func (o EdgeConfigOutput) TeamId() pulumi.StringOutput {
	return o.ApplyT(func(v *EdgeConfig) pulumi.StringOutput { return v.TeamId }).(pulumi.StringOutput)
}

type EdgeConfigArrayOutput struct{ *pulumi.OutputState }

func (EdgeConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EdgeConfig)(nil)).Elem()
}

func (o EdgeConfigArrayOutput) ToEdgeConfigArrayOutput() EdgeConfigArrayOutput {
	return o
}

func (o EdgeConfigArrayOutput) ToEdgeConfigArrayOutputWithContext(ctx context.Context) EdgeConfigArrayOutput {
	return o
}

func (o EdgeConfigArrayOutput) Index(i pulumi.IntInput) EdgeConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EdgeConfig {
		return vs[0].([]*EdgeConfig)[vs[1].(int)]
	}).(EdgeConfigOutput)
}

type EdgeConfigMapOutput struct{ *pulumi.OutputState }

func (EdgeConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EdgeConfig)(nil)).Elem()
}

func (o EdgeConfigMapOutput) ToEdgeConfigMapOutput() EdgeConfigMapOutput {
	return o
}

func (o EdgeConfigMapOutput) ToEdgeConfigMapOutputWithContext(ctx context.Context) EdgeConfigMapOutput {
	return o
}

func (o EdgeConfigMapOutput) MapIndex(k pulumi.StringInput) EdgeConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EdgeConfig {
		return vs[0].(map[string]*EdgeConfig)[vs[1].(string)]
	}).(EdgeConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EdgeConfigInput)(nil)).Elem(), &EdgeConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*EdgeConfigArrayInput)(nil)).Elem(), EdgeConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EdgeConfigMapInput)(nil)).Elem(), EdgeConfigMap{})
	pulumi.RegisterOutputType(EdgeConfigOutput{})
	pulumi.RegisterOutputType(EdgeConfigArrayOutput{})
	pulumi.RegisterOutputType(EdgeConfigMapOutput{})
}
