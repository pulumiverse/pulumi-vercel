// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vercel

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-vercel/sdk/go/vercel/internal"
)

// Provides information about an existing project within Vercel.
//
// A Project groups deployments and custom domains. To deploy on Vercel, you need a Project.
//
// For more detailed information, please see the [Vercel documentation](https://vercel.com/docs/concepts/projects/overview).
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
//			foo, err := vercel.LookupProject(ctx, &vercel.LookupProjectArgs{
//				Name: "my-existing-project",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("projectId", foo.Id)
//			return nil
//		})
//	}
//
// ```
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
	// The name of the project.
	Name string `pulumi:"name"`
	// The team ID the project exists beneath. Required when configuring a team resource if a default team has not been set in the provider.
	TeamId *string `pulumi:"teamId"`
}

// A collection of values returned by getProject.
type LookupProjectResult struct {
	// Automatically assign custom production domains after each Production deployment via merge to the production branch or Vercel CLI deploy with --prod. Defaults to `true`
	AutoAssignCustomDomains bool `pulumi:"autoAssignCustomDomains"`
	// Vercel provides a set of Environment Variables that are automatically populated by the System, such as the URL of the Deployment or the name of the Git branch deployed. To expose them to your Deployments, enable this field
	AutomaticallyExposeSystemEnvironmentVariables bool `pulumi:"automaticallyExposeSystemEnvironmentVariables"`
	// The build command for this project. If omitted, this value will be automatically detected.
	BuildCommand string `pulumi:"buildCommand"`
	// Allows Vercel Customer Support to inspect all Deployments' source code in this project to assist with debugging.
	CustomerSuccessCodeVisibility bool `pulumi:"customerSuccessCodeVisibility"`
	// The dev command for this project. If omitted, this value will be automatically detected.
	DevCommand string `pulumi:"devCommand"`
	// If no index file is present within a directory, the directory contents will be displayed.
	DirectoryListing bool `pulumi:"directoryListing"`
	// A list of environment variables that should be configured for the project.
	Environments []GetProjectEnvironment `pulumi:"environments"`
	// The framework that is being used for this project. If omitted, no framework is selected.
	Framework string `pulumi:"framework"`
	// Automatically failover Serverless Functions to the nearest region. You can customize regions through vercel.json. A new Deployment is required for your changes to take effect.
	FunctionFailover bool `pulumi:"functionFailover"`
	// Configuration for Git Comments.
	GitComments GetProjectGitComments `pulumi:"gitComments"`
	// Ensures that pull requests targeting your Git repository must be authorized by a member of your Team before deploying if your Project has Environment Variables or if the pull request includes a change to vercel.json.
	GitForkProtection bool `pulumi:"gitForkProtection"`
	// Enables Git LFS support. Git LFS replaces large files such as audio samples, videos, datasets, and graphics with text pointers inside Git, while storing the file contents on a remote server like GitHub.com or GitHub Enterprise.
	GitLfs bool `pulumi:"gitLfs"`
	// The Git Repository that will be connected to the project. When this is defined, any pushes to the specified connected Git Repository will be automatically deployed. This requires the corresponding Vercel for [Github](https://vercel.com/docs/concepts/git/vercel-for-github), [Gitlab](https://vercel.com/docs/concepts/git/vercel-for-gitlab) or [Bitbucket](https://vercel.com/docs/concepts/git/vercel-for-bitbucket) plugins to be installed.
	GitRepository GetProjectGitRepository `pulumi:"gitRepository"`
	// The ID of this resource.
	Id string `pulumi:"id"`
	// When a commit is pushed to the Git repository that is connected with your Project, its SHA will determine if a new Build has to be issued. If the SHA was deployed before, no new Build will be issued. You can customize this behavior with a command that exits with code 1 (new Build needed) or code 0.
	IgnoreCommand string `pulumi:"ignoreCommand"`
	// The install command for this project. If omitted, this value will be automatically detected.
	InstallCommand string `pulumi:"installCommand"`
	// The name of the project.
	Name string `pulumi:"name"`
	// Configuration for OpenID Connect (OIDC) tokens.
	OidcTokenConfig GetProjectOidcTokenConfig `pulumi:"oidcTokenConfig"`
	// Disable Deployment Protection for CORS preflight `OPTIONS` requests for a list of paths.
	OptionsAllowlist GetProjectOptionsAllowlist `pulumi:"optionsAllowlist"`
	// The output directory of the project. When null is used this value will be automatically detected.
	OutputDirectory string `pulumi:"outputDirectory"`
	// Ensures visitors of your Preview Deployments must enter a password in order to gain access.
	PasswordProtection GetProjectPasswordProtection `pulumi:"passwordProtection"`
	// Whether comments are enabled on your Preview Deployments.
	PreviewComments bool `pulumi:"previewComments"`
	// If enabled, builds for the Production environment will be prioritized over Preview environments.
	PrioritiseProductionBuilds bool `pulumi:"prioritiseProductionBuilds"`
	// Allows automation services to bypass Vercel Authentication and Password Protection for both Preview and Production Deployments on this project when using an HTTP header named `x-vercel-protection-bypass`.
	ProtectionBypassForAutomation bool `pulumi:"protectionBypassForAutomation"`
	// Specifies whether the source code and logs of the deployments for this project should be public or not.
	PublicSource bool `pulumi:"publicSource"`
	// Resource Configuration for the project.
	ResourceConfig GetProjectResourceConfig `pulumi:"resourceConfig"`
	// The name of a directory or relative path to the source code of your project. When null is used it will default to the project root.
	RootDirectory string `pulumi:"rootDirectory"`
	// The region on Vercel's network to which your Serverless Functions are deployed. It should be close to any data source your Serverless Function might depend on. A new Deployment is required for your changes to take effect. Please see [Vercel's documentation](https://vercel.com/docs/concepts/edge-network/regions) for a full list of regions.
	ServerlessFunctionRegion string `pulumi:"serverlessFunctionRegion"`
	// Ensures that outdated clients always fetch the correct version for a given deployment. This value defines how long Vercel keeps Skew Protection active.
	SkewProtection string `pulumi:"skewProtection"`
	// The team ID the project exists beneath. Required when configuring a team resource if a default team has not been set in the provider.
	TeamId string `pulumi:"teamId"`
	// Ensures only visitors from an allowed IP address can access your deployment.
	TrustedIps GetProjectTrustedIps `pulumi:"trustedIps"`
	// Ensures visitors to your Preview Deployments are logged into Vercel and have a minimum of Viewer access on your team.
	VercelAuthentication GetProjectVercelAuthentication `pulumi:"vercelAuthentication"`
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
	// The name of the project.
	Name pulumi.StringInput `pulumi:"name"`
	// The team ID the project exists beneath. Required when configuring a team resource if a default team has not been set in the provider.
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

// Automatically assign custom production domains after each Production deployment via merge to the production branch or Vercel CLI deploy with --prod. Defaults to `true`
func (o LookupProjectResultOutput) AutoAssignCustomDomains() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.AutoAssignCustomDomains }).(pulumi.BoolOutput)
}

// Vercel provides a set of Environment Variables that are automatically populated by the System, such as the URL of the Deployment or the name of the Git branch deployed. To expose them to your Deployments, enable this field
func (o LookupProjectResultOutput) AutomaticallyExposeSystemEnvironmentVariables() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.AutomaticallyExposeSystemEnvironmentVariables }).(pulumi.BoolOutput)
}

// The build command for this project. If omitted, this value will be automatically detected.
func (o LookupProjectResultOutput) BuildCommand() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.BuildCommand }).(pulumi.StringOutput)
}

// Allows Vercel Customer Support to inspect all Deployments' source code in this project to assist with debugging.
func (o LookupProjectResultOutput) CustomerSuccessCodeVisibility() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.CustomerSuccessCodeVisibility }).(pulumi.BoolOutput)
}

// The dev command for this project. If omitted, this value will be automatically detected.
func (o LookupProjectResultOutput) DevCommand() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.DevCommand }).(pulumi.StringOutput)
}

// If no index file is present within a directory, the directory contents will be displayed.
func (o LookupProjectResultOutput) DirectoryListing() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.DirectoryListing }).(pulumi.BoolOutput)
}

// A list of environment variables that should be configured for the project.
func (o LookupProjectResultOutput) Environments() GetProjectEnvironmentArrayOutput {
	return o.ApplyT(func(v LookupProjectResult) []GetProjectEnvironment { return v.Environments }).(GetProjectEnvironmentArrayOutput)
}

// The framework that is being used for this project. If omitted, no framework is selected.
func (o LookupProjectResultOutput) Framework() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Framework }).(pulumi.StringOutput)
}

// Automatically failover Serverless Functions to the nearest region. You can customize regions through vercel.json. A new Deployment is required for your changes to take effect.
func (o LookupProjectResultOutput) FunctionFailover() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.FunctionFailover }).(pulumi.BoolOutput)
}

// Configuration for Git Comments.
func (o LookupProjectResultOutput) GitComments() GetProjectGitCommentsOutput {
	return o.ApplyT(func(v LookupProjectResult) GetProjectGitComments { return v.GitComments }).(GetProjectGitCommentsOutput)
}

// Ensures that pull requests targeting your Git repository must be authorized by a member of your Team before deploying if your Project has Environment Variables or if the pull request includes a change to vercel.json.
func (o LookupProjectResultOutput) GitForkProtection() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.GitForkProtection }).(pulumi.BoolOutput)
}

// Enables Git LFS support. Git LFS replaces large files such as audio samples, videos, datasets, and graphics with text pointers inside Git, while storing the file contents on a remote server like GitHub.com or GitHub Enterprise.
func (o LookupProjectResultOutput) GitLfs() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.GitLfs }).(pulumi.BoolOutput)
}

// The Git Repository that will be connected to the project. When this is defined, any pushes to the specified connected Git Repository will be automatically deployed. This requires the corresponding Vercel for [Github](https://vercel.com/docs/concepts/git/vercel-for-github), [Gitlab](https://vercel.com/docs/concepts/git/vercel-for-gitlab) or [Bitbucket](https://vercel.com/docs/concepts/git/vercel-for-bitbucket) plugins to be installed.
func (o LookupProjectResultOutput) GitRepository() GetProjectGitRepositoryOutput {
	return o.ApplyT(func(v LookupProjectResult) GetProjectGitRepository { return v.GitRepository }).(GetProjectGitRepositoryOutput)
}

// The ID of this resource.
func (o LookupProjectResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Id }).(pulumi.StringOutput)
}

// When a commit is pushed to the Git repository that is connected with your Project, its SHA will determine if a new Build has to be issued. If the SHA was deployed before, no new Build will be issued. You can customize this behavior with a command that exits with code 1 (new Build needed) or code 0.
func (o LookupProjectResultOutput) IgnoreCommand() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.IgnoreCommand }).(pulumi.StringOutput)
}

// The install command for this project. If omitted, this value will be automatically detected.
func (o LookupProjectResultOutput) InstallCommand() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.InstallCommand }).(pulumi.StringOutput)
}

// The name of the project.
func (o LookupProjectResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Name }).(pulumi.StringOutput)
}

// Configuration for OpenID Connect (OIDC) tokens.
func (o LookupProjectResultOutput) OidcTokenConfig() GetProjectOidcTokenConfigOutput {
	return o.ApplyT(func(v LookupProjectResult) GetProjectOidcTokenConfig { return v.OidcTokenConfig }).(GetProjectOidcTokenConfigOutput)
}

// Disable Deployment Protection for CORS preflight `OPTIONS` requests for a list of paths.
func (o LookupProjectResultOutput) OptionsAllowlist() GetProjectOptionsAllowlistOutput {
	return o.ApplyT(func(v LookupProjectResult) GetProjectOptionsAllowlist { return v.OptionsAllowlist }).(GetProjectOptionsAllowlistOutput)
}

// The output directory of the project. When null is used this value will be automatically detected.
func (o LookupProjectResultOutput) OutputDirectory() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.OutputDirectory }).(pulumi.StringOutput)
}

// Ensures visitors of your Preview Deployments must enter a password in order to gain access.
func (o LookupProjectResultOutput) PasswordProtection() GetProjectPasswordProtectionOutput {
	return o.ApplyT(func(v LookupProjectResult) GetProjectPasswordProtection { return v.PasswordProtection }).(GetProjectPasswordProtectionOutput)
}

// Whether comments are enabled on your Preview Deployments.
func (o LookupProjectResultOutput) PreviewComments() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.PreviewComments }).(pulumi.BoolOutput)
}

// If enabled, builds for the Production environment will be prioritized over Preview environments.
func (o LookupProjectResultOutput) PrioritiseProductionBuilds() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.PrioritiseProductionBuilds }).(pulumi.BoolOutput)
}

// Allows automation services to bypass Vercel Authentication and Password Protection for both Preview and Production Deployments on this project when using an HTTP header named `x-vercel-protection-bypass`.
func (o LookupProjectResultOutput) ProtectionBypassForAutomation() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.ProtectionBypassForAutomation }).(pulumi.BoolOutput)
}

// Specifies whether the source code and logs of the deployments for this project should be public or not.
func (o LookupProjectResultOutput) PublicSource() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.PublicSource }).(pulumi.BoolOutput)
}

// Resource Configuration for the project.
func (o LookupProjectResultOutput) ResourceConfig() GetProjectResourceConfigOutput {
	return o.ApplyT(func(v LookupProjectResult) GetProjectResourceConfig { return v.ResourceConfig }).(GetProjectResourceConfigOutput)
}

// The name of a directory or relative path to the source code of your project. When null is used it will default to the project root.
func (o LookupProjectResultOutput) RootDirectory() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.RootDirectory }).(pulumi.StringOutput)
}

// The region on Vercel's network to which your Serverless Functions are deployed. It should be close to any data source your Serverless Function might depend on. A new Deployment is required for your changes to take effect. Please see [Vercel's documentation](https://vercel.com/docs/concepts/edge-network/regions) for a full list of regions.
func (o LookupProjectResultOutput) ServerlessFunctionRegion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.ServerlessFunctionRegion }).(pulumi.StringOutput)
}

// Ensures that outdated clients always fetch the correct version for a given deployment. This value defines how long Vercel keeps Skew Protection active.
func (o LookupProjectResultOutput) SkewProtection() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.SkewProtection }).(pulumi.StringOutput)
}

// The team ID the project exists beneath. Required when configuring a team resource if a default team has not been set in the provider.
func (o LookupProjectResultOutput) TeamId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.TeamId }).(pulumi.StringOutput)
}

// Ensures only visitors from an allowed IP address can access your deployment.
func (o LookupProjectResultOutput) TrustedIps() GetProjectTrustedIpsOutput {
	return o.ApplyT(func(v LookupProjectResult) GetProjectTrustedIps { return v.TrustedIps }).(GetProjectTrustedIpsOutput)
}

// Ensures visitors to your Preview Deployments are logged into Vercel and have a minimum of Viewer access on your team.
func (o LookupProjectResultOutput) VercelAuthentication() GetProjectVercelAuthenticationOutput {
	return o.ApplyT(func(v LookupProjectResult) GetProjectVercelAuthentication { return v.VercelAuthentication }).(GetProjectVercelAuthenticationOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProjectResultOutput{})
}
