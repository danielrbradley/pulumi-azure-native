# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class MediaServiceContentKeyPolicy(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    The name of the resource
    """
    properties: pulumi.Output[dict]
    """
    The properties of the Content Key Policy.
      * `created` (`str`) - The creation date of the Policy
      * `description` (`str`) - A description for the Policy.
      * `last_modified` (`str`) - The last modified date of the Policy
      * `options` (`list`) - The Key Policy options.
        * `configuration` (`dict`) - The key delivery configuration.
        * `name` (`str`) - The Policy Option description.
        * `policy_option_id` (`str`) - The legacy Policy Option ID.
        * `restriction` (`dict`) - The requirements that must be met to deliver keys with this configuration

      * `policy_id` (`str`) - The legacy Policy ID.
    """
    type: pulumi.Output[str]
    """
    The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
    """
    def __init__(__self__, resource_name, opts=None, account_name=None, name=None, properties=None, resource_group_name=None, __props__=None, __name__=None, __opts__=None):
        """
        A Content Key Policy resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: The Media Services account name.
        :param pulumi.Input[str] name: The Content Key Policy name.
        :param pulumi.Input[dict] properties: The properties of the Content Key Policy.
        :param pulumi.Input[str] resource_group_name: The name of the resource group within the Azure subscription.

        The **properties** object supports the following:

          * `description` (`pulumi.Input[str]`) - A description for the Policy.
          * `options` (`pulumi.Input[list]`) - The Key Policy options.
            * `configuration` (`pulumi.Input[dict]`) - The key delivery configuration.
            * `name` (`pulumi.Input[str]`) - The Policy Option description.
            * `restriction` (`pulumi.Input[dict]`) - The requirements that must be met to deliver keys with this configuration
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

            if account_name is None:
                raise TypeError("Missing required property 'account_name'")
            __props__['account_name'] = account_name
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['type'] = None
        super(MediaServiceContentKeyPolicy, __self__).__init__(
            'azurerm:media:MediaServiceContentKeyPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing MediaServiceContentKeyPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return MediaServiceContentKeyPolicy(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop