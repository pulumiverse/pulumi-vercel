// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vercel

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-vercel/sdk/go/vercel/internal"
)

func LookupProject(ctx *pulumi.Context, args *LookupProjectArgs, opts ...pulumi.InvokeOption) (*LookupProjectResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupProjectResult
	err := ctx.Invoke("vercel:index/getProject:getProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProject.
type LookupProjectArgs struct {
	Name   string  `pulumi:"name"`
	TeamId *string `pulumi:"teamId"`
}

// A collection of values returned by getProject.
type LookupProjectResult struct {
	AutoAssignCustomDomains                       bool                           `pulumi:"autoAssignCustomDomains"`
	AutomaticallyExposeSystemEnvironmentVariables bool                           `pulumi:"automaticallyExposeSystemEnvironmentVariables"`
	BuildCommand                                  string                         `pulumi:"buildCommand"`
	CustomerSuccessCodeVisibility                 bool                           `pulumi:"customerSuccessCodeVisibility"`
	DevCommand                                    string                         `pulumi:"devCommand"`
	DirectoryListing                              bool                           `pulumi:"directoryListing"`
	Environments                                  []GetProjectEnvironment        `pulumi:"environments"`
	Framework                                     string                         `pulumi:"framework"`
	FunctionFailover                              bool                           `pulumi:"functionFailover"`
	GitComments                                   GetProjectGitComments          `pulumi:"gitComments"`
	GitForkProtection                             bool                           `pulumi:"gitForkProtection"`
	GitLfs                                        bool                           `pulumi:"gitLfs"`
	GitRepository                                 GetProjectGitRepository        `pulumi:"gitRepository"`
	Id                                            string                         `pulumi:"id"`
	IgnoreCommand                                 string                         `pulumi:"ignoreCommand"`
	InstallCommand                                string                         `pulumi:"installCommand"`
	Name                                          string                         `pulumi:"name"`
	OidcTokenConfig                               GetProjectOidcTokenConfig      `pulumi:"oidcTokenConfig"`
	OptionsAllowlist                              GetProjectOptionsAllowlist     `pulumi:"optionsAllowlist"`
	OutputDirectory                               string                         `pulumi:"outputDirectory"`
	PasswordProtection                            GetProjectPasswordProtection   `pulumi:"passwordProtection"`
	PreviewComments                               bool                           `pulumi:"previewComments"`
	PrioritiseProductionBuilds                    bool                           `pulumi:"prioritiseProductionBuilds"`
	ProtectionBypassForAutomation                 bool                           `pulumi:"protectionBypassForAutomation"`
	ProtectionBypassForAutomationSecret           string                         `pulumi:"protectionBypassForAutomationSecret"`
	PublicSource                                  bool                           `pulumi:"publicSource"`
	ResourceConfig                                GetProjectResourceConfig       `pulumi:"resourceConfig"`
	RootDirectory                                 string                         `pulumi:"rootDirectory"`
	ServerlessFunctionRegion                      string                         `pulumi:"serverlessFunctionRegion"`
	SkewProtection                                string                         `pulumi:"skewProtection"`
	TeamId                                        string                         `pulumi:"teamId"`
	TrustedIps                                    GetProjectTrustedIps           `pulumi:"trustedIps"`
	VercelAuthentication                          GetProjectVercelAuthentication `pulumi:"vercelAuthentication"`
}

func LookupProjectOutput(ctx *pulumi.Context, args LookupProjectOutputArgs, opts ...pulumi.InvokeOption) LookupProjectResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProjectResultOutput, error) {
			args := v.(LookupProjectArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupProjectResult
			secret, err := ctx.InvokePackageRaw("vercel:index/getProject:getProject", args, &rv, "", opts...)
			if err != nil {
				return LookupProjectResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupProjectResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupProjectResultOutput), nil
			}
			return output, nil
		}).(LookupProjectResultOutput)
}

// A collection of arguments for invoking getProject.
type LookupProjectOutputArgs struct {
	Name   pulumi.StringInput    `pulumi:"name"`
	TeamId pulumi.StringPtrInput `pulumi:"teamId"`
}

func (LookupProjectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectArgs)(nil)).Elem()
}

// A collection of values returned by getProject.
type LookupProjectResultOutput struct{ *pulumi.OutputState }

func (LookupProjectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectResult)(nil)).Elem()
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutput() LookupProjectResultOutput {
	return o
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutputWithContext(ctx context.Context) LookupProjectResultOutput {
	return o
}

func (o LookupProjectResultOutput) AutoAssignCustomDomains() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.AutoAssignCustomDomains }).(pulumi.BoolOutput)
}

func (o LookupProjectResultOutput) AutomaticallyExposeSystemEnvironmentVariables() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.AutomaticallyExposeSystemEnvironmentVariables }).(pulumi.BoolOutput)
}

func (o LookupProjectResultOutput) BuildCommand() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.BuildCommand }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) CustomerSuccessCodeVisibility() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.CustomerSuccessCodeVisibility }).(pulumi.BoolOutput)
}

func (o LookupProjectResultOutput) DevCommand() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.DevCommand }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) DirectoryListing() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.DirectoryListing }).(pulumi.BoolOutput)
}

func (o LookupProjectResultOutput) Environments() GetProjectEnvironmentArrayOutput {
	return o.ApplyT(func(v LookupProjectResult) []GetProjectEnvironment { return v.Environments }).(GetProjectEnvironmentArrayOutput)
}

func (o LookupProjectResultOutput) Framework() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Framework }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) FunctionFailover() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.FunctionFailover }).(pulumi.BoolOutput)
}

func (o LookupProjectResultOutput) GitComments() GetProjectGitCommentsOutput {
	return o.ApplyT(func(v LookupProjectResult) GetProjectGitComments { return v.GitComments }).(GetProjectGitCommentsOutput)
}

func (o LookupProjectResultOutput) GitForkProtection() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.GitForkProtection }).(pulumi.BoolOutput)
}

func (o LookupProjectResultOutput) GitLfs() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.GitLfs }).(pulumi.BoolOutput)
}

func (o LookupProjectResultOutput) GitRepository() GetProjectGitRepositoryOutput {
	return o.ApplyT(func(v LookupProjectResult) GetProjectGitRepository { return v.GitRepository }).(GetProjectGitRepositoryOutput)
}

func (o LookupProjectResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) IgnoreCommand() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.IgnoreCommand }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) InstallCommand() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.InstallCommand }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) OidcTokenConfig() GetProjectOidcTokenConfigOutput {
	return o.ApplyT(func(v LookupProjectResult) GetProjectOidcTokenConfig { return v.OidcTokenConfig }).(GetProjectOidcTokenConfigOutput)
}

func (o LookupProjectResultOutput) OptionsAllowlist() GetProjectOptionsAllowlistOutput {
	return o.ApplyT(func(v LookupProjectResult) GetProjectOptionsAllowlist { return v.OptionsAllowlist }).(GetProjectOptionsAllowlistOutput)
}

func (o LookupProjectResultOutput) OutputDirectory() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.OutputDirectory }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) PasswordProtection() GetProjectPasswordProtectionOutput {
	return o.ApplyT(func(v LookupProjectResult) GetProjectPasswordProtection { return v.PasswordProtection }).(GetProjectPasswordProtectionOutput)
}

func (o LookupProjectResultOutput) PreviewComments() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.PreviewComments }).(pulumi.BoolOutput)
}

func (o LookupProjectResultOutput) PrioritiseProductionBuilds() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.PrioritiseProductionBuilds }).(pulumi.BoolOutput)
}

func (o LookupProjectResultOutput) ProtectionBypassForAutomation() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.ProtectionBypassForAutomation }).(pulumi.BoolOutput)
}

func (o LookupProjectResultOutput) ProtectionBypassForAutomationSecret() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.ProtectionBypassForAutomationSecret }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) PublicSource() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.PublicSource }).(pulumi.BoolOutput)
}

func (o LookupProjectResultOutput) ResourceConfig() GetProjectResourceConfigOutput {
	return o.ApplyT(func(v LookupProjectResult) GetProjectResourceConfig { return v.ResourceConfig }).(GetProjectResourceConfigOutput)
}

func (o LookupProjectResultOutput) RootDirectory() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.RootDirectory }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) ServerlessFunctionRegion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.ServerlessFunctionRegion }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) SkewProtection() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.SkewProtection }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) TeamId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.TeamId }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) TrustedIps() GetProjectTrustedIpsOutput {
	return o.ApplyT(func(v LookupProjectResult) GetProjectTrustedIps { return v.TrustedIps }).(GetProjectTrustedIpsOutput)
}

func (o LookupProjectResultOutput) VercelAuthentication() GetProjectVercelAuthenticationOutput {
	return o.ApplyT(func(v LookupProjectResult) GetProjectVercelAuthentication { return v.VercelAuthentication }).(GetProjectVercelAuthenticationOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProjectResultOutput{})
}
