// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vercel

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-vercel/sdk/v3/go/vercel/internal"
)

// Provider a datasource for managing a team member.
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
//			_, err := vercel.LookupTeamMember(ctx, &vercel.LookupTeamMemberArgs{
//				UserId: "uuuuuuuuuuuuuuuuuuuuuuuuuu",
//				TeamId: "team_xxxxxxxxxxxxxxxxxxxxxxxx",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupTeamMember(ctx *pulumi.Context, args *LookupTeamMemberArgs, opts ...pulumi.InvokeOption) (*LookupTeamMemberResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTeamMemberResult
	err := ctx.Invoke("vercel:index/getTeamMember:getTeamMember", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTeamMember.
type LookupTeamMemberArgs struct {
	// The ID of the existing Vercel Team.
	TeamId string `pulumi:"teamId"`
	// The ID of the existing Vercel Team Member.
	UserId string `pulumi:"userId"`
}

// A collection of values returned by getTeamMember.
type LookupTeamMemberResult struct {
	// If access groups are enabled on the team, and the user is a CONTRIBUTOR, `projects`, `accessGroups` or both must be specified. A set of access groups IDs that the user should be granted access to.
	AccessGroups []string `pulumi:"accessGroups"`
	// The email address of the existing Vercel Team Member.
	Email string `pulumi:"email"`
	// The ID of this resource.
	Id string `pulumi:"id"`
	// If access groups are enabled on the team, and the user is a CONTRIBUTOR, `projects`, `accessGroups` or both must be specified. A set of projects that the user should be granted access to, along with their role in each project.
	Projects []GetTeamMemberProject `pulumi:"projects"`
	// The role that the user should have in the project. One of 'MEMBER', 'OWNER', 'VIEWER', 'DEVELOPER', 'BILLING' or 'CONTRIBUTOR'. Depending on your Team's plan, some of these roles may be unavailable.
	Role string `pulumi:"role"`
	// The ID of the existing Vercel Team.
	TeamId string `pulumi:"teamId"`
	// The ID of the existing Vercel Team Member.
	UserId string `pulumi:"userId"`
}

func LookupTeamMemberOutput(ctx *pulumi.Context, args LookupTeamMemberOutputArgs, opts ...pulumi.InvokeOption) LookupTeamMemberResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTeamMemberResultOutput, error) {
			args := v.(LookupTeamMemberArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupTeamMemberResult
			secret, err := ctx.InvokePackageRaw("vercel:index/getTeamMember:getTeamMember", args, &rv, "", opts...)
			if err != nil {
				return LookupTeamMemberResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupTeamMemberResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupTeamMemberResultOutput), nil
			}
			return output, nil
		}).(LookupTeamMemberResultOutput)
}

// A collection of arguments for invoking getTeamMember.
type LookupTeamMemberOutputArgs struct {
	// The ID of the existing Vercel Team.
	TeamId pulumi.StringInput `pulumi:"teamId"`
	// The ID of the existing Vercel Team Member.
	UserId pulumi.StringInput `pulumi:"userId"`
}

func (LookupTeamMemberOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTeamMemberArgs)(nil)).Elem()
}

// A collection of values returned by getTeamMember.
type LookupTeamMemberResultOutput struct{ *pulumi.OutputState }

func (LookupTeamMemberResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTeamMemberResult)(nil)).Elem()
}

func (o LookupTeamMemberResultOutput) ToLookupTeamMemberResultOutput() LookupTeamMemberResultOutput {
	return o
}

func (o LookupTeamMemberResultOutput) ToLookupTeamMemberResultOutputWithContext(ctx context.Context) LookupTeamMemberResultOutput {
	return o
}

// If access groups are enabled on the team, and the user is a CONTRIBUTOR, `projects`, `accessGroups` or both must be specified. A set of access groups IDs that the user should be granted access to.
func (o LookupTeamMemberResultOutput) AccessGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupTeamMemberResult) []string { return v.AccessGroups }).(pulumi.StringArrayOutput)
}

// The email address of the existing Vercel Team Member.
func (o LookupTeamMemberResultOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTeamMemberResult) string { return v.Email }).(pulumi.StringOutput)
}

// The ID of this resource.
func (o LookupTeamMemberResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTeamMemberResult) string { return v.Id }).(pulumi.StringOutput)
}

// If access groups are enabled on the team, and the user is a CONTRIBUTOR, `projects`, `accessGroups` or both must be specified. A set of projects that the user should be granted access to, along with their role in each project.
func (o LookupTeamMemberResultOutput) Projects() GetTeamMemberProjectArrayOutput {
	return o.ApplyT(func(v LookupTeamMemberResult) []GetTeamMemberProject { return v.Projects }).(GetTeamMemberProjectArrayOutput)
}

// The role that the user should have in the project. One of 'MEMBER', 'OWNER', 'VIEWER', 'DEVELOPER', 'BILLING' or 'CONTRIBUTOR'. Depending on your Team's plan, some of these roles may be unavailable.
func (o LookupTeamMemberResultOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTeamMemberResult) string { return v.Role }).(pulumi.StringOutput)
}

// The ID of the existing Vercel Team.
func (o LookupTeamMemberResultOutput) TeamId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTeamMemberResult) string { return v.TeamId }).(pulumi.StringOutput)
}

// The ID of the existing Vercel Team Member.
func (o LookupTeamMemberResultOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTeamMemberResult) string { return v.UserId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTeamMemberResultOutput{})
}
