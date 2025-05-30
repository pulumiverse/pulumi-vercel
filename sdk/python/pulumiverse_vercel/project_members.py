# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ProjectMembersArgs', 'ProjectMembers']

@pulumi.input_type
class ProjectMembersArgs:
    def __init__(__self__, *,
                 members: pulumi.Input[Sequence[pulumi.Input['ProjectMembersMemberArgs']]],
                 project_id: pulumi.Input[str],
                 team_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ProjectMembers resource.
        :param pulumi.Input[Sequence[pulumi.Input['ProjectMembersMemberArgs']]] members: The set of members to manage for this project.
        :param pulumi.Input[str] project_id: The ID of the existing Vercel Project.
        :param pulumi.Input[str] team_id: The team ID to add the project to. Required when configuring a team resource if a default team has not been set in the provider.
        """
        pulumi.set(__self__, "members", members)
        pulumi.set(__self__, "project_id", project_id)
        if team_id is not None:
            pulumi.set(__self__, "team_id", team_id)

    @property
    @pulumi.getter
    def members(self) -> pulumi.Input[Sequence[pulumi.Input['ProjectMembersMemberArgs']]]:
        """
        The set of members to manage for this project.
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: pulumi.Input[Sequence[pulumi.Input['ProjectMembersMemberArgs']]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[str]:
        """
        The ID of the existing Vercel Project.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> Optional[pulumi.Input[str]]:
        """
        The team ID to add the project to. Required when configuring a team resource if a default team has not been set in the provider.
        """
        return pulumi.get(self, "team_id")

    @team_id.setter
    def team_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "team_id", value)


@pulumi.input_type
class _ProjectMembersState:
    def __init__(__self__, *,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input['ProjectMembersMemberArgs']]]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 team_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ProjectMembers resources.
        :param pulumi.Input[Sequence[pulumi.Input['ProjectMembersMemberArgs']]] members: The set of members to manage for this project.
        :param pulumi.Input[str] project_id: The ID of the existing Vercel Project.
        :param pulumi.Input[str] team_id: The team ID to add the project to. Required when configuring a team resource if a default team has not been set in the provider.
        """
        if members is not None:
            pulumi.set(__self__, "members", members)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if team_id is not None:
            pulumi.set(__self__, "team_id", team_id)

    @property
    @pulumi.getter
    def members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ProjectMembersMemberArgs']]]]:
        """
        The set of members to manage for this project.
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ProjectMembersMemberArgs']]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the existing Vercel Project.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> Optional[pulumi.Input[str]]:
        """
        The team ID to add the project to. Required when configuring a team resource if a default team has not been set in the provider.
        """
        return pulumi.get(self, "team_id")

    @team_id.setter
    def team_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "team_id", value)


class ProjectMembers(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ProjectMembersMemberArgs', 'ProjectMembersMemberArgsDict']]]]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 team_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_vercel as vercel

        example = vercel.Project("example", name="example-with-members")
        example_project_members = vercel.ProjectMembers("example",
            project_id=example.id,
            members=[
                {
                    "email": "user@example.com",
                    "role": "PROJECT_VIEWER",
                },
                {
                    "username": "some-example-user",
                    "role": "PROJECT_DEVELOPER",
                },
            ])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['ProjectMembersMemberArgs', 'ProjectMembersMemberArgsDict']]]] members: The set of members to manage for this project.
        :param pulumi.Input[str] project_id: The ID of the existing Vercel Project.
        :param pulumi.Input[str] team_id: The team ID to add the project to. Required when configuring a team resource if a default team has not been set in the provider.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProjectMembersArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_vercel as vercel

        example = vercel.Project("example", name="example-with-members")
        example_project_members = vercel.ProjectMembers("example",
            project_id=example.id,
            members=[
                {
                    "email": "user@example.com",
                    "role": "PROJECT_VIEWER",
                },
                {
                    "username": "some-example-user",
                    "role": "PROJECT_DEVELOPER",
                },
            ])
        ```

        :param str resource_name: The name of the resource.
        :param ProjectMembersArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProjectMembersArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ProjectMembersMemberArgs', 'ProjectMembersMemberArgsDict']]]]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 team_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProjectMembersArgs.__new__(ProjectMembersArgs)

            if members is None and not opts.urn:
                raise TypeError("Missing required property 'members'")
            __props__.__dict__["members"] = members
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["team_id"] = team_id
        super(ProjectMembers, __self__).__init__(
            'vercel:index/projectMembers:ProjectMembers',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            members: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ProjectMembersMemberArgs', 'ProjectMembersMemberArgsDict']]]]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            team_id: Optional[pulumi.Input[str]] = None) -> 'ProjectMembers':
        """
        Get an existing ProjectMembers resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['ProjectMembersMemberArgs', 'ProjectMembersMemberArgsDict']]]] members: The set of members to manage for this project.
        :param pulumi.Input[str] project_id: The ID of the existing Vercel Project.
        :param pulumi.Input[str] team_id: The team ID to add the project to. Required when configuring a team resource if a default team has not been set in the provider.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProjectMembersState.__new__(_ProjectMembersState)

        __props__.__dict__["members"] = members
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["team_id"] = team_id
        return ProjectMembers(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def members(self) -> pulumi.Output[Sequence['outputs.ProjectMembersMember']]:
        """
        The set of members to manage for this project.
        """
        return pulumi.get(self, "members")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The ID of the existing Vercel Project.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> pulumi.Output[str]:
        """
        The team ID to add the project to. Required when configuring a team resource if a default team has not been set in the provider.
        """
        return pulumi.get(self, "team_id")

