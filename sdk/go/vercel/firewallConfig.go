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

type FirewallConfig struct {
	pulumi.CustomResourceState

	// Whether firewall is enabled or not.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// IP rules to apply to the project.
	IpRules FirewallConfigIpRulesPtrOutput `pulumi:"ipRules"`
	// The managed rulesets that are enabled.
	ManagedRulesets FirewallConfigManagedRulesetsPtrOutput `pulumi:"managedRulesets"`
	// The ID of the project this configuration belongs to.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Custom rules to apply to the project
	Rules FirewallConfigRulesPtrOutput `pulumi:"rules"`
	// The ID of the team this project belongs to.
	TeamId pulumi.StringOutput `pulumi:"teamId"`
}

// NewFirewallConfig registers a new resource with the given unique name, arguments, and options.
func NewFirewallConfig(ctx *pulumi.Context,
	name string, args *FirewallConfigArgs, opts ...pulumi.ResourceOption) (*FirewallConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FirewallConfig
	err := ctx.RegisterResource("vercel:index/firewallConfig:FirewallConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallConfig gets an existing FirewallConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallConfigState, opts ...pulumi.ResourceOption) (*FirewallConfig, error) {
	var resource FirewallConfig
	err := ctx.ReadResource("vercel:index/firewallConfig:FirewallConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallConfig resources.
type firewallConfigState struct {
	// Whether firewall is enabled or not.
	Enabled *bool `pulumi:"enabled"`
	// IP rules to apply to the project.
	IpRules *FirewallConfigIpRules `pulumi:"ipRules"`
	// The managed rulesets that are enabled.
	ManagedRulesets *FirewallConfigManagedRulesets `pulumi:"managedRulesets"`
	// The ID of the project this configuration belongs to.
	ProjectId *string `pulumi:"projectId"`
	// Custom rules to apply to the project
	Rules *FirewallConfigRules `pulumi:"rules"`
	// The ID of the team this project belongs to.
	TeamId *string `pulumi:"teamId"`
}

type FirewallConfigState struct {
	// Whether firewall is enabled or not.
	Enabled pulumi.BoolPtrInput
	// IP rules to apply to the project.
	IpRules FirewallConfigIpRulesPtrInput
	// The managed rulesets that are enabled.
	ManagedRulesets FirewallConfigManagedRulesetsPtrInput
	// The ID of the project this configuration belongs to.
	ProjectId pulumi.StringPtrInput
	// Custom rules to apply to the project
	Rules FirewallConfigRulesPtrInput
	// The ID of the team this project belongs to.
	TeamId pulumi.StringPtrInput
}

func (FirewallConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallConfigState)(nil)).Elem()
}

type firewallConfigArgs struct {
	// Whether firewall is enabled or not.
	Enabled *bool `pulumi:"enabled"`
	// IP rules to apply to the project.
	IpRules *FirewallConfigIpRules `pulumi:"ipRules"`
	// The managed rulesets that are enabled.
	ManagedRulesets *FirewallConfigManagedRulesets `pulumi:"managedRulesets"`
	// The ID of the project this configuration belongs to.
	ProjectId string `pulumi:"projectId"`
	// Custom rules to apply to the project
	Rules *FirewallConfigRules `pulumi:"rules"`
	// The ID of the team this project belongs to.
	TeamId *string `pulumi:"teamId"`
}

// The set of arguments for constructing a FirewallConfig resource.
type FirewallConfigArgs struct {
	// Whether firewall is enabled or not.
	Enabled pulumi.BoolPtrInput
	// IP rules to apply to the project.
	IpRules FirewallConfigIpRulesPtrInput
	// The managed rulesets that are enabled.
	ManagedRulesets FirewallConfigManagedRulesetsPtrInput
	// The ID of the project this configuration belongs to.
	ProjectId pulumi.StringInput
	// Custom rules to apply to the project
	Rules FirewallConfigRulesPtrInput
	// The ID of the team this project belongs to.
	TeamId pulumi.StringPtrInput
}

func (FirewallConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallConfigArgs)(nil)).Elem()
}

type FirewallConfigInput interface {
	pulumi.Input

	ToFirewallConfigOutput() FirewallConfigOutput
	ToFirewallConfigOutputWithContext(ctx context.Context) FirewallConfigOutput
}

func (*FirewallConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallConfig)(nil)).Elem()
}

func (i *FirewallConfig) ToFirewallConfigOutput() FirewallConfigOutput {
	return i.ToFirewallConfigOutputWithContext(context.Background())
}

func (i *FirewallConfig) ToFirewallConfigOutputWithContext(ctx context.Context) FirewallConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallConfigOutput)
}

// FirewallConfigArrayInput is an input type that accepts FirewallConfigArray and FirewallConfigArrayOutput values.
// You can construct a concrete instance of `FirewallConfigArrayInput` via:
//
//	FirewallConfigArray{ FirewallConfigArgs{...} }
type FirewallConfigArrayInput interface {
	pulumi.Input

	ToFirewallConfigArrayOutput() FirewallConfigArrayOutput
	ToFirewallConfigArrayOutputWithContext(context.Context) FirewallConfigArrayOutput
}

type FirewallConfigArray []FirewallConfigInput

func (FirewallConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallConfig)(nil)).Elem()
}

func (i FirewallConfigArray) ToFirewallConfigArrayOutput() FirewallConfigArrayOutput {
	return i.ToFirewallConfigArrayOutputWithContext(context.Background())
}

func (i FirewallConfigArray) ToFirewallConfigArrayOutputWithContext(ctx context.Context) FirewallConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallConfigArrayOutput)
}

// FirewallConfigMapInput is an input type that accepts FirewallConfigMap and FirewallConfigMapOutput values.
// You can construct a concrete instance of `FirewallConfigMapInput` via:
//
//	FirewallConfigMap{ "key": FirewallConfigArgs{...} }
type FirewallConfigMapInput interface {
	pulumi.Input

	ToFirewallConfigMapOutput() FirewallConfigMapOutput
	ToFirewallConfigMapOutputWithContext(context.Context) FirewallConfigMapOutput
}

type FirewallConfigMap map[string]FirewallConfigInput

func (FirewallConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallConfig)(nil)).Elem()
}

func (i FirewallConfigMap) ToFirewallConfigMapOutput() FirewallConfigMapOutput {
	return i.ToFirewallConfigMapOutputWithContext(context.Background())
}

func (i FirewallConfigMap) ToFirewallConfigMapOutputWithContext(ctx context.Context) FirewallConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallConfigMapOutput)
}

type FirewallConfigOutput struct{ *pulumi.OutputState }

func (FirewallConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallConfig)(nil)).Elem()
}

func (o FirewallConfigOutput) ToFirewallConfigOutput() FirewallConfigOutput {
	return o
}

func (o FirewallConfigOutput) ToFirewallConfigOutputWithContext(ctx context.Context) FirewallConfigOutput {
	return o
}

// Whether firewall is enabled or not.
func (o FirewallConfigOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *FirewallConfig) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// IP rules to apply to the project.
func (o FirewallConfigOutput) IpRules() FirewallConfigIpRulesPtrOutput {
	return o.ApplyT(func(v *FirewallConfig) FirewallConfigIpRulesPtrOutput { return v.IpRules }).(FirewallConfigIpRulesPtrOutput)
}

// The managed rulesets that are enabled.
func (o FirewallConfigOutput) ManagedRulesets() FirewallConfigManagedRulesetsPtrOutput {
	return o.ApplyT(func(v *FirewallConfig) FirewallConfigManagedRulesetsPtrOutput { return v.ManagedRulesets }).(FirewallConfigManagedRulesetsPtrOutput)
}

// The ID of the project this configuration belongs to.
func (o FirewallConfigOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallConfig) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Custom rules to apply to the project
func (o FirewallConfigOutput) Rules() FirewallConfigRulesPtrOutput {
	return o.ApplyT(func(v *FirewallConfig) FirewallConfigRulesPtrOutput { return v.Rules }).(FirewallConfigRulesPtrOutput)
}

// The ID of the team this project belongs to.
func (o FirewallConfigOutput) TeamId() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallConfig) pulumi.StringOutput { return v.TeamId }).(pulumi.StringOutput)
}

type FirewallConfigArrayOutput struct{ *pulumi.OutputState }

func (FirewallConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallConfig)(nil)).Elem()
}

func (o FirewallConfigArrayOutput) ToFirewallConfigArrayOutput() FirewallConfigArrayOutput {
	return o
}

func (o FirewallConfigArrayOutput) ToFirewallConfigArrayOutputWithContext(ctx context.Context) FirewallConfigArrayOutput {
	return o
}

func (o FirewallConfigArrayOutput) Index(i pulumi.IntInput) FirewallConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallConfig {
		return vs[0].([]*FirewallConfig)[vs[1].(int)]
	}).(FirewallConfigOutput)
}

type FirewallConfigMapOutput struct{ *pulumi.OutputState }

func (FirewallConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallConfig)(nil)).Elem()
}

func (o FirewallConfigMapOutput) ToFirewallConfigMapOutput() FirewallConfigMapOutput {
	return o
}

func (o FirewallConfigMapOutput) ToFirewallConfigMapOutputWithContext(ctx context.Context) FirewallConfigMapOutput {
	return o
}

func (o FirewallConfigMapOutput) MapIndex(k pulumi.StringInput) FirewallConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallConfig {
		return vs[0].(map[string]*FirewallConfig)[vs[1].(string)]
	}).(FirewallConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallConfigInput)(nil)).Elem(), &FirewallConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallConfigArrayInput)(nil)).Elem(), FirewallConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallConfigMapInput)(nil)).Elem(), FirewallConfigMap{})
	pulumi.RegisterOutputType(FirewallConfigOutput{})
	pulumi.RegisterOutputType(FirewallConfigArrayOutput{})
	pulumi.RegisterOutputType(FirewallConfigMapOutput{})
}
