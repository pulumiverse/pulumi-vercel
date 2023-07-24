// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vercel

import (
	"context"
	"reflect"

	"errors"
	"github.com/omercnet/pulumi-vercel/sdk/go/vercel/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DnsRecord struct {
	pulumi.CustomResourceState

	// The domain name, or zone, that the DNS record should be created beneath.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// The priority of the MX record. The priority specifies the sequence that an email server receives emails. A smaller value
	// indicates a higher priority.
	MxPriority pulumi.IntPtrOutput `pulumi:"mxPriority"`
	// The subdomain name of the record. This should be an empty string if the rercord is for the root domain.
	Name pulumi.StringOutput `pulumi:"name"`
	// Settings for an SRV record.
	Srv DnsRecordSrvPtrOutput `pulumi:"srv"`
	// The team ID that the domain and DNS records belong to. Required when configuring a team resource if a default team has
	// not been set in the provider.
	TeamId pulumi.StringOutput `pulumi:"teamId"`
	// The TTL value in seconds. Must be a number between 60 and 2147483647. If unspecified, it will default to 60 seconds.
	Ttl pulumi.IntOutput `pulumi:"ttl"`
	// The type of DNS record. Available types: `A`, `AAAA`, `ALIAS`, `CAA`, `CNAME`, `MX`, `NS`, `SRV`, `TXT`.
	Type pulumi.StringOutput `pulumi:"type"`
	// The value of the DNS record. The format depends on the 'type' property. For an 'A' record, this should be a valid IPv4
	// address. For an 'AAAA' record, this should be an IPv6 address. For 'ALIAS' records, this should be a hostname. For 'CAA'
	// records, this should specify specify which Certificate Authorities (CAs) are allowed to issue certificates for the
	// domain. For 'CNAME' records, this should be a different domain name. For 'MX' records, this should specify the mail
	// server responsible for accepting messages on behalf of the domain name. For 'TXT' records, this can contain arbitrary
	// text.
	Value pulumi.StringPtrOutput `pulumi:"value"`
}

// NewDnsRecord registers a new resource with the given unique name, arguments, and options.
func NewDnsRecord(ctx *pulumi.Context,
	name string, args *DnsRecordArgs, opts ...pulumi.ResourceOption) (*DnsRecord, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DnsRecord
	err := ctx.RegisterResource("vercel:index/dnsRecord:DnsRecord", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDnsRecord gets an existing DnsRecord resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDnsRecord(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DnsRecordState, opts ...pulumi.ResourceOption) (*DnsRecord, error) {
	var resource DnsRecord
	err := ctx.ReadResource("vercel:index/dnsRecord:DnsRecord", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DnsRecord resources.
type dnsRecordState struct {
	// The domain name, or zone, that the DNS record should be created beneath.
	Domain *string `pulumi:"domain"`
	// The priority of the MX record. The priority specifies the sequence that an email server receives emails. A smaller value
	// indicates a higher priority.
	MxPriority *int `pulumi:"mxPriority"`
	// The subdomain name of the record. This should be an empty string if the rercord is for the root domain.
	Name *string `pulumi:"name"`
	// Settings for an SRV record.
	Srv *DnsRecordSrv `pulumi:"srv"`
	// The team ID that the domain and DNS records belong to. Required when configuring a team resource if a default team has
	// not been set in the provider.
	TeamId *string `pulumi:"teamId"`
	// The TTL value in seconds. Must be a number between 60 and 2147483647. If unspecified, it will default to 60 seconds.
	Ttl *int `pulumi:"ttl"`
	// The type of DNS record. Available types: `A`, `AAAA`, `ALIAS`, `CAA`, `CNAME`, `MX`, `NS`, `SRV`, `TXT`.
	Type *string `pulumi:"type"`
	// The value of the DNS record. The format depends on the 'type' property. For an 'A' record, this should be a valid IPv4
	// address. For an 'AAAA' record, this should be an IPv6 address. For 'ALIAS' records, this should be a hostname. For 'CAA'
	// records, this should specify specify which Certificate Authorities (CAs) are allowed to issue certificates for the
	// domain. For 'CNAME' records, this should be a different domain name. For 'MX' records, this should specify the mail
	// server responsible for accepting messages on behalf of the domain name. For 'TXT' records, this can contain arbitrary
	// text.
	Value *string `pulumi:"value"`
}

type DnsRecordState struct {
	// The domain name, or zone, that the DNS record should be created beneath.
	Domain pulumi.StringPtrInput
	// The priority of the MX record. The priority specifies the sequence that an email server receives emails. A smaller value
	// indicates a higher priority.
	MxPriority pulumi.IntPtrInput
	// The subdomain name of the record. This should be an empty string if the rercord is for the root domain.
	Name pulumi.StringPtrInput
	// Settings for an SRV record.
	Srv DnsRecordSrvPtrInput
	// The team ID that the domain and DNS records belong to. Required when configuring a team resource if a default team has
	// not been set in the provider.
	TeamId pulumi.StringPtrInput
	// The TTL value in seconds. Must be a number between 60 and 2147483647. If unspecified, it will default to 60 seconds.
	Ttl pulumi.IntPtrInput
	// The type of DNS record. Available types: `A`, `AAAA`, `ALIAS`, `CAA`, `CNAME`, `MX`, `NS`, `SRV`, `TXT`.
	Type pulumi.StringPtrInput
	// The value of the DNS record. The format depends on the 'type' property. For an 'A' record, this should be a valid IPv4
	// address. For an 'AAAA' record, this should be an IPv6 address. For 'ALIAS' records, this should be a hostname. For 'CAA'
	// records, this should specify specify which Certificate Authorities (CAs) are allowed to issue certificates for the
	// domain. For 'CNAME' records, this should be a different domain name. For 'MX' records, this should specify the mail
	// server responsible for accepting messages on behalf of the domain name. For 'TXT' records, this can contain arbitrary
	// text.
	Value pulumi.StringPtrInput
}

func (DnsRecordState) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsRecordState)(nil)).Elem()
}

type dnsRecordArgs struct {
	// The domain name, or zone, that the DNS record should be created beneath.
	Domain string `pulumi:"domain"`
	// The priority of the MX record. The priority specifies the sequence that an email server receives emails. A smaller value
	// indicates a higher priority.
	MxPriority *int `pulumi:"mxPriority"`
	// The subdomain name of the record. This should be an empty string if the rercord is for the root domain.
	Name *string `pulumi:"name"`
	// Settings for an SRV record.
	Srv *DnsRecordSrv `pulumi:"srv"`
	// The team ID that the domain and DNS records belong to. Required when configuring a team resource if a default team has
	// not been set in the provider.
	TeamId *string `pulumi:"teamId"`
	// The TTL value in seconds. Must be a number between 60 and 2147483647. If unspecified, it will default to 60 seconds.
	Ttl *int `pulumi:"ttl"`
	// The type of DNS record. Available types: `A`, `AAAA`, `ALIAS`, `CAA`, `CNAME`, `MX`, `NS`, `SRV`, `TXT`.
	Type string `pulumi:"type"`
	// The value of the DNS record. The format depends on the 'type' property. For an 'A' record, this should be a valid IPv4
	// address. For an 'AAAA' record, this should be an IPv6 address. For 'ALIAS' records, this should be a hostname. For 'CAA'
	// records, this should specify specify which Certificate Authorities (CAs) are allowed to issue certificates for the
	// domain. For 'CNAME' records, this should be a different domain name. For 'MX' records, this should specify the mail
	// server responsible for accepting messages on behalf of the domain name. For 'TXT' records, this can contain arbitrary
	// text.
	Value *string `pulumi:"value"`
}

// The set of arguments for constructing a DnsRecord resource.
type DnsRecordArgs struct {
	// The domain name, or zone, that the DNS record should be created beneath.
	Domain pulumi.StringInput
	// The priority of the MX record. The priority specifies the sequence that an email server receives emails. A smaller value
	// indicates a higher priority.
	MxPriority pulumi.IntPtrInput
	// The subdomain name of the record. This should be an empty string if the rercord is for the root domain.
	Name pulumi.StringPtrInput
	// Settings for an SRV record.
	Srv DnsRecordSrvPtrInput
	// The team ID that the domain and DNS records belong to. Required when configuring a team resource if a default team has
	// not been set in the provider.
	TeamId pulumi.StringPtrInput
	// The TTL value in seconds. Must be a number between 60 and 2147483647. If unspecified, it will default to 60 seconds.
	Ttl pulumi.IntPtrInput
	// The type of DNS record. Available types: `A`, `AAAA`, `ALIAS`, `CAA`, `CNAME`, `MX`, `NS`, `SRV`, `TXT`.
	Type pulumi.StringInput
	// The value of the DNS record. The format depends on the 'type' property. For an 'A' record, this should be a valid IPv4
	// address. For an 'AAAA' record, this should be an IPv6 address. For 'ALIAS' records, this should be a hostname. For 'CAA'
	// records, this should specify specify which Certificate Authorities (CAs) are allowed to issue certificates for the
	// domain. For 'CNAME' records, this should be a different domain name. For 'MX' records, this should specify the mail
	// server responsible for accepting messages on behalf of the domain name. For 'TXT' records, this can contain arbitrary
	// text.
	Value pulumi.StringPtrInput
}

func (DnsRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsRecordArgs)(nil)).Elem()
}

type DnsRecordInput interface {
	pulumi.Input

	ToDnsRecordOutput() DnsRecordOutput
	ToDnsRecordOutputWithContext(ctx context.Context) DnsRecordOutput
}

func (*DnsRecord) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsRecord)(nil)).Elem()
}

func (i *DnsRecord) ToDnsRecordOutput() DnsRecordOutput {
	return i.ToDnsRecordOutputWithContext(context.Background())
}

func (i *DnsRecord) ToDnsRecordOutputWithContext(ctx context.Context) DnsRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsRecordOutput)
}

// DnsRecordArrayInput is an input type that accepts DnsRecordArray and DnsRecordArrayOutput values.
// You can construct a concrete instance of `DnsRecordArrayInput` via:
//
//	DnsRecordArray{ DnsRecordArgs{...} }
type DnsRecordArrayInput interface {
	pulumi.Input

	ToDnsRecordArrayOutput() DnsRecordArrayOutput
	ToDnsRecordArrayOutputWithContext(context.Context) DnsRecordArrayOutput
}

type DnsRecordArray []DnsRecordInput

func (DnsRecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DnsRecord)(nil)).Elem()
}

func (i DnsRecordArray) ToDnsRecordArrayOutput() DnsRecordArrayOutput {
	return i.ToDnsRecordArrayOutputWithContext(context.Background())
}

func (i DnsRecordArray) ToDnsRecordArrayOutputWithContext(ctx context.Context) DnsRecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsRecordArrayOutput)
}

// DnsRecordMapInput is an input type that accepts DnsRecordMap and DnsRecordMapOutput values.
// You can construct a concrete instance of `DnsRecordMapInput` via:
//
//	DnsRecordMap{ "key": DnsRecordArgs{...} }
type DnsRecordMapInput interface {
	pulumi.Input

	ToDnsRecordMapOutput() DnsRecordMapOutput
	ToDnsRecordMapOutputWithContext(context.Context) DnsRecordMapOutput
}

type DnsRecordMap map[string]DnsRecordInput

func (DnsRecordMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DnsRecord)(nil)).Elem()
}

func (i DnsRecordMap) ToDnsRecordMapOutput() DnsRecordMapOutput {
	return i.ToDnsRecordMapOutputWithContext(context.Background())
}

func (i DnsRecordMap) ToDnsRecordMapOutputWithContext(ctx context.Context) DnsRecordMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsRecordMapOutput)
}

type DnsRecordOutput struct{ *pulumi.OutputState }

func (DnsRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsRecord)(nil)).Elem()
}

func (o DnsRecordOutput) ToDnsRecordOutput() DnsRecordOutput {
	return o
}

func (o DnsRecordOutput) ToDnsRecordOutputWithContext(ctx context.Context) DnsRecordOutput {
	return o
}

// The domain name, or zone, that the DNS record should be created beneath.
func (o DnsRecordOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsRecord) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

// The priority of the MX record. The priority specifies the sequence that an email server receives emails. A smaller value
// indicates a higher priority.
func (o DnsRecordOutput) MxPriority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DnsRecord) pulumi.IntPtrOutput { return v.MxPriority }).(pulumi.IntPtrOutput)
}

// The subdomain name of the record. This should be an empty string if the rercord is for the root domain.
func (o DnsRecordOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsRecord) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Settings for an SRV record.
func (o DnsRecordOutput) Srv() DnsRecordSrvPtrOutput {
	return o.ApplyT(func(v *DnsRecord) DnsRecordSrvPtrOutput { return v.Srv }).(DnsRecordSrvPtrOutput)
}

// The team ID that the domain and DNS records belong to. Required when configuring a team resource if a default team has
// not been set in the provider.
func (o DnsRecordOutput) TeamId() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsRecord) pulumi.StringOutput { return v.TeamId }).(pulumi.StringOutput)
}

// The TTL value in seconds. Must be a number between 60 and 2147483647. If unspecified, it will default to 60 seconds.
func (o DnsRecordOutput) Ttl() pulumi.IntOutput {
	return o.ApplyT(func(v *DnsRecord) pulumi.IntOutput { return v.Ttl }).(pulumi.IntOutput)
}

// The type of DNS record. Available types: `A`, `AAAA`, `ALIAS`, `CAA`, `CNAME`, `MX`, `NS`, `SRV`, `TXT`.
func (o DnsRecordOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsRecord) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The value of the DNS record. The format depends on the 'type' property. For an 'A' record, this should be a valid IPv4
// address. For an 'AAAA' record, this should be an IPv6 address. For 'ALIAS' records, this should be a hostname. For 'CAA'
// records, this should specify specify which Certificate Authorities (CAs) are allowed to issue certificates for the
// domain. For 'CNAME' records, this should be a different domain name. For 'MX' records, this should specify the mail
// server responsible for accepting messages on behalf of the domain name. For 'TXT' records, this can contain arbitrary
// text.
func (o DnsRecordOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsRecord) pulumi.StringPtrOutput { return v.Value }).(pulumi.StringPtrOutput)
}

type DnsRecordArrayOutput struct{ *pulumi.OutputState }

func (DnsRecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DnsRecord)(nil)).Elem()
}

func (o DnsRecordArrayOutput) ToDnsRecordArrayOutput() DnsRecordArrayOutput {
	return o
}

func (o DnsRecordArrayOutput) ToDnsRecordArrayOutputWithContext(ctx context.Context) DnsRecordArrayOutput {
	return o
}

func (o DnsRecordArrayOutput) Index(i pulumi.IntInput) DnsRecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DnsRecord {
		return vs[0].([]*DnsRecord)[vs[1].(int)]
	}).(DnsRecordOutput)
}

type DnsRecordMapOutput struct{ *pulumi.OutputState }

func (DnsRecordMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DnsRecord)(nil)).Elem()
}

func (o DnsRecordMapOutput) ToDnsRecordMapOutput() DnsRecordMapOutput {
	return o
}

func (o DnsRecordMapOutput) ToDnsRecordMapOutputWithContext(ctx context.Context) DnsRecordMapOutput {
	return o
}

func (o DnsRecordMapOutput) MapIndex(k pulumi.StringInput) DnsRecordOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DnsRecord {
		return vs[0].(map[string]*DnsRecord)[vs[1].(string)]
	}).(DnsRecordOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DnsRecordInput)(nil)).Elem(), &DnsRecord{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsRecordArrayInput)(nil)).Elem(), DnsRecordArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsRecordMapInput)(nil)).Elem(), DnsRecordMap{})
	pulumi.RegisterOutputType(DnsRecordOutput{})
	pulumi.RegisterOutputType(DnsRecordArrayOutput{})
	pulumi.RegisterOutputType(DnsRecordMapOutput{})
}