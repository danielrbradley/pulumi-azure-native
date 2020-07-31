# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class Workspace(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    The location of the resource. This cannot be changed after the resource is created.
    """
    name: pulumi.Output[str]
    """
    The name of the resource.
    """
    properties: pulumi.Output[dict]
    """
    The properties of the machine learning workspace.
      * `creation_time` (`str`) - The creation time for this workspace resource.
      * `key_vault_identifier_id` (`str`) - The key vault identifier used for encrypted workspaces.
      * `owner_email` (`str`) - The email id of the owner for this workspace.
      * `studio_endpoint` (`str`) - The regional endpoint for the machine learning studio service which hosts this workspace.
      * `user_storage_account_id` (`str`) - The fully qualified arm id of the storage account associated with this workspace.
      * `workspace_id` (`str`) - The immutable id associated with this workspace.
      * `workspace_state` (`str`) - The current state of workspace resource.
      * `workspace_type` (`str`) - The type of this workspace.
    """
    tags: pulumi.Output[dict]
    """
    The tags of the resource.
    """
    type: pulumi.Output[str]
    """
    The type of the resource.
    """
    def __init__(__self__, resource_name, opts=None, location=None, name=None, properties=None, resource_group_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        An object that represents a machine learning workspace.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: The location of the resource. This cannot be changed after the resource is created.
        :param pulumi.Input[str] name: The name of the machine learning workspace.
        :param pulumi.Input[dict] properties: The properties of the machine learning workspace.
        :param pulumi.Input[str] resource_group_name: The name of the resource group to which the machine learning workspace belongs.
        :param pulumi.Input[dict] tags: The tags of the resource.

        The **properties** object supports the following:

          * `key_vault_identifier_id` (`pulumi.Input[str]`) - The key vault identifier used for encrypted workspaces.
          * `owner_email` (`pulumi.Input[str]`) - The email id of the owner for this workspace.
          * `user_storage_account_id` (`pulumi.Input[str]`) - The fully qualified arm id of the storage account associated with this workspace.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['type'] = None
        super(Workspace, __self__).__init__(
            'azurerm:machinelearning/v20160401:Workspace',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Workspace resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Workspace(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop