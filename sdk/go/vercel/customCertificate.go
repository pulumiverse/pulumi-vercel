// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vercel

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-vercel/sdk/v3/go/vercel/internal"
)

// Provides a Custom Certificate Resource, allowing Custom Certificates to be uploaded to Vercel.
//
// By default, Vercel provides all domains with a custom SSL certificates. However, Enterprise teams can upload their own custom SSL certificate.
//
// For more detailed information, please see the [Vercel documentation](https://vercel.com/docs/domains/custom-SSL-certificate).
type CustomCertificate struct {
	pulumi.CustomResourceState

	// The certificate itself. Should be in PEM format.
	Certificate pulumi.StringOutput `pulumi:"certificate"`
	// The Certificate Authority root certificate such as one of Let's Encrypt's ISRG root certificates. This will be provided by your certificate issuer and is different to the core certificate. This may be included in their download process or available for download on their website. Should be in PEM format.
	CertificateAuthorityCertificate pulumi.StringOutput `pulumi:"certificateAuthorityCertificate"`
	// The private key of the Certificate. Should be in PEM format.
	PrivateKey pulumi.StringOutput `pulumi:"privateKey"`
	// The ID of the team the Custom Certificate should exist under. Required when configuring a team resource if a default team has not been set in the provider.
	TeamId pulumi.StringOutput `pulumi:"teamId"`
}

// NewCustomCertificate registers a new resource with the given unique name, arguments, and options.
func NewCustomCertificate(ctx *pulumi.Context,
	name string, args *CustomCertificateArgs, opts ...pulumi.ResourceOption) (*CustomCertificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Certificate == nil {
		return nil, errors.New("invalid value for required argument 'Certificate'")
	}
	if args.CertificateAuthorityCertificate == nil {
		return nil, errors.New("invalid value for required argument 'CertificateAuthorityCertificate'")
	}
	if args.PrivateKey == nil {
		return nil, errors.New("invalid value for required argument 'PrivateKey'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CustomCertificate
	err := ctx.RegisterResource("vercel:index/customCertificate:CustomCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomCertificate gets an existing CustomCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomCertificateState, opts ...pulumi.ResourceOption) (*CustomCertificate, error) {
	var resource CustomCertificate
	err := ctx.ReadResource("vercel:index/customCertificate:CustomCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomCertificate resources.
type customCertificateState struct {
	// The certificate itself. Should be in PEM format.
	Certificate *string `pulumi:"certificate"`
	// The Certificate Authority root certificate such as one of Let's Encrypt's ISRG root certificates. This will be provided by your certificate issuer and is different to the core certificate. This may be included in their download process or available for download on their website. Should be in PEM format.
	CertificateAuthorityCertificate *string `pulumi:"certificateAuthorityCertificate"`
	// The private key of the Certificate. Should be in PEM format.
	PrivateKey *string `pulumi:"privateKey"`
	// The ID of the team the Custom Certificate should exist under. Required when configuring a team resource if a default team has not been set in the provider.
	TeamId *string `pulumi:"teamId"`
}

type CustomCertificateState struct {
	// The certificate itself. Should be in PEM format.
	Certificate pulumi.StringPtrInput
	// The Certificate Authority root certificate such as one of Let's Encrypt's ISRG root certificates. This will be provided by your certificate issuer and is different to the core certificate. This may be included in their download process or available for download on their website. Should be in PEM format.
	CertificateAuthorityCertificate pulumi.StringPtrInput
	// The private key of the Certificate. Should be in PEM format.
	PrivateKey pulumi.StringPtrInput
	// The ID of the team the Custom Certificate should exist under. Required when configuring a team resource if a default team has not been set in the provider.
	TeamId pulumi.StringPtrInput
}

func (CustomCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*customCertificateState)(nil)).Elem()
}

type customCertificateArgs struct {
	// The certificate itself. Should be in PEM format.
	Certificate string `pulumi:"certificate"`
	// The Certificate Authority root certificate such as one of Let's Encrypt's ISRG root certificates. This will be provided by your certificate issuer and is different to the core certificate. This may be included in their download process or available for download on their website. Should be in PEM format.
	CertificateAuthorityCertificate string `pulumi:"certificateAuthorityCertificate"`
	// The private key of the Certificate. Should be in PEM format.
	PrivateKey string `pulumi:"privateKey"`
	// The ID of the team the Custom Certificate should exist under. Required when configuring a team resource if a default team has not been set in the provider.
	TeamId *string `pulumi:"teamId"`
}

// The set of arguments for constructing a CustomCertificate resource.
type CustomCertificateArgs struct {
	// The certificate itself. Should be in PEM format.
	Certificate pulumi.StringInput
	// The Certificate Authority root certificate such as one of Let's Encrypt's ISRG root certificates. This will be provided by your certificate issuer and is different to the core certificate. This may be included in their download process or available for download on their website. Should be in PEM format.
	CertificateAuthorityCertificate pulumi.StringInput
	// The private key of the Certificate. Should be in PEM format.
	PrivateKey pulumi.StringInput
	// The ID of the team the Custom Certificate should exist under. Required when configuring a team resource if a default team has not been set in the provider.
	TeamId pulumi.StringPtrInput
}

func (CustomCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customCertificateArgs)(nil)).Elem()
}

type CustomCertificateInput interface {
	pulumi.Input

	ToCustomCertificateOutput() CustomCertificateOutput
	ToCustomCertificateOutputWithContext(ctx context.Context) CustomCertificateOutput
}

func (*CustomCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomCertificate)(nil)).Elem()
}

func (i *CustomCertificate) ToCustomCertificateOutput() CustomCertificateOutput {
	return i.ToCustomCertificateOutputWithContext(context.Background())
}

func (i *CustomCertificate) ToCustomCertificateOutputWithContext(ctx context.Context) CustomCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomCertificateOutput)
}

// CustomCertificateArrayInput is an input type that accepts CustomCertificateArray and CustomCertificateArrayOutput values.
// You can construct a concrete instance of `CustomCertificateArrayInput` via:
//
//	CustomCertificateArray{ CustomCertificateArgs{...} }
type CustomCertificateArrayInput interface {
	pulumi.Input

	ToCustomCertificateArrayOutput() CustomCertificateArrayOutput
	ToCustomCertificateArrayOutputWithContext(context.Context) CustomCertificateArrayOutput
}

type CustomCertificateArray []CustomCertificateInput

func (CustomCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomCertificate)(nil)).Elem()
}

func (i CustomCertificateArray) ToCustomCertificateArrayOutput() CustomCertificateArrayOutput {
	return i.ToCustomCertificateArrayOutputWithContext(context.Background())
}

func (i CustomCertificateArray) ToCustomCertificateArrayOutputWithContext(ctx context.Context) CustomCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomCertificateArrayOutput)
}

// CustomCertificateMapInput is an input type that accepts CustomCertificateMap and CustomCertificateMapOutput values.
// You can construct a concrete instance of `CustomCertificateMapInput` via:
//
//	CustomCertificateMap{ "key": CustomCertificateArgs{...} }
type CustomCertificateMapInput interface {
	pulumi.Input

	ToCustomCertificateMapOutput() CustomCertificateMapOutput
	ToCustomCertificateMapOutputWithContext(context.Context) CustomCertificateMapOutput
}

type CustomCertificateMap map[string]CustomCertificateInput

func (CustomCertificateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomCertificate)(nil)).Elem()
}

func (i CustomCertificateMap) ToCustomCertificateMapOutput() CustomCertificateMapOutput {
	return i.ToCustomCertificateMapOutputWithContext(context.Background())
}

func (i CustomCertificateMap) ToCustomCertificateMapOutputWithContext(ctx context.Context) CustomCertificateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomCertificateMapOutput)
}

type CustomCertificateOutput struct{ *pulumi.OutputState }

func (CustomCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomCertificate)(nil)).Elem()
}

func (o CustomCertificateOutput) ToCustomCertificateOutput() CustomCertificateOutput {
	return o
}

func (o CustomCertificateOutput) ToCustomCertificateOutputWithContext(ctx context.Context) CustomCertificateOutput {
	return o
}

// The certificate itself. Should be in PEM format.
func (o CustomCertificateOutput) Certificate() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomCertificate) pulumi.StringOutput { return v.Certificate }).(pulumi.StringOutput)
}

// The Certificate Authority root certificate such as one of Let's Encrypt's ISRG root certificates. This will be provided by your certificate issuer and is different to the core certificate. This may be included in their download process or available for download on their website. Should be in PEM format.
func (o CustomCertificateOutput) CertificateAuthorityCertificate() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomCertificate) pulumi.StringOutput { return v.CertificateAuthorityCertificate }).(pulumi.StringOutput)
}

// The private key of the Certificate. Should be in PEM format.
func (o CustomCertificateOutput) PrivateKey() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomCertificate) pulumi.StringOutput { return v.PrivateKey }).(pulumi.StringOutput)
}

// The ID of the team the Custom Certificate should exist under. Required when configuring a team resource if a default team has not been set in the provider.
func (o CustomCertificateOutput) TeamId() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomCertificate) pulumi.StringOutput { return v.TeamId }).(pulumi.StringOutput)
}

type CustomCertificateArrayOutput struct{ *pulumi.OutputState }

func (CustomCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomCertificate)(nil)).Elem()
}

func (o CustomCertificateArrayOutput) ToCustomCertificateArrayOutput() CustomCertificateArrayOutput {
	return o
}

func (o CustomCertificateArrayOutput) ToCustomCertificateArrayOutputWithContext(ctx context.Context) CustomCertificateArrayOutput {
	return o
}

func (o CustomCertificateArrayOutput) Index(i pulumi.IntInput) CustomCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CustomCertificate {
		return vs[0].([]*CustomCertificate)[vs[1].(int)]
	}).(CustomCertificateOutput)
}

type CustomCertificateMapOutput struct{ *pulumi.OutputState }

func (CustomCertificateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomCertificate)(nil)).Elem()
}

func (o CustomCertificateMapOutput) ToCustomCertificateMapOutput() CustomCertificateMapOutput {
	return o
}

func (o CustomCertificateMapOutput) ToCustomCertificateMapOutputWithContext(ctx context.Context) CustomCertificateMapOutput {
	return o
}

func (o CustomCertificateMapOutput) MapIndex(k pulumi.StringInput) CustomCertificateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CustomCertificate {
		return vs[0].(map[string]*CustomCertificate)[vs[1].(string)]
	}).(CustomCertificateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CustomCertificateInput)(nil)).Elem(), &CustomCertificate{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomCertificateArrayInput)(nil)).Elem(), CustomCertificateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomCertificateMapInput)(nil)).Elem(), CustomCertificateMap{})
	pulumi.RegisterOutputType(CustomCertificateOutput{})
	pulumi.RegisterOutputType(CustomCertificateArrayOutput{})
	pulumi.RegisterOutputType(CustomCertificateMapOutput{})
}
