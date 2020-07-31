# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class Disk(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    The location of the resource.
    """
    name: pulumi.Output[str]
    """
    The name of the resource.
    """
    properties: pulumi.Output[dict]
    """
    The properties of the resource.
      * `created_date` (`str`) - The creation date of the disk.
      * `disk_blob_name` (`str`) - When backed by a blob, the name of the VHD blob without extension.
      * `disk_size_gi_b` (`float`) - The size of the disk in GibiBytes.
      * `disk_type` (`str`) - The storage type for the disk (i.e. Standard, Premium).
      * `disk_uri` (`str`) - When backed by a blob, the URI of underlying blob.
      * `host_caching` (`str`) - The host caching policy of the disk (i.e. None, ReadOnly, ReadWrite).
      * `leased_by_lab_vm_id` (`str`) - The resource ID of the VM to which this disk is leased.
      * `managed_disk_id` (`str`) - When backed by managed disk, this is the ID of the compute disk resource.
      * `provisioning_state` (`str`) - The provisioning status of the resource.
      * `unique_identifier` (`str`) - The unique immutable identifier of a resource (Guid).
    """
    tags: pulumi.Output[dict]
    """
    The tags of the resource.
    """
    type: pulumi.Output[str]
    """
    The type of the resource.
    """
    def __init__(__self__, resource_name, opts=None, lab_name=None, location=None, name=None, properties=None, resource_group_name=None, tags=None, user_name=None, __props__=None, __name__=None, __opts__=None):
        """
        A Disk.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] lab_name: The name of the lab.
        :param pulumi.Input[str] location: The location of the resource.
        :param pulumi.Input[str] name: The name of the disk.
        :param pulumi.Input[dict] properties: The properties of the resource.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[dict] tags: The tags of the resource.
        :param pulumi.Input[str] user_name: The name of the user profile.

        The **properties** object supports the following:

          * `disk_blob_name` (`pulumi.Input[str]`) - When backed by a blob, the name of the VHD blob without extension.
          * `disk_size_gi_b` (`pulumi.Input[float]`) - The size of the disk in GibiBytes.
          * `disk_type` (`pulumi.Input[str]`) - The storage type for the disk (i.e. Standard, Premium).
          * `disk_uri` (`pulumi.Input[str]`) - When backed by a blob, the URI of underlying blob.
          * `host_caching` (`pulumi.Input[str]`) - The host caching policy of the disk (i.e. None, ReadOnly, ReadWrite).
          * `leased_by_lab_vm_id` (`pulumi.Input[str]`) - The resource ID of the VM to which this disk is leased.
          * `managed_disk_id` (`pulumi.Input[str]`) - When backed by managed disk, this is the ID of the compute disk resource.
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

            if lab_name is None:
                raise TypeError("Missing required property 'lab_name'")
            __props__['lab_name'] = lab_name
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if properties is None:
                raise TypeError("Missing required property 'properties'")
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            if user_name is None:
                raise TypeError("Missing required property 'user_name'")
            __props__['user_name'] = user_name
            __props__['type'] = None
        super(Disk, __self__).__init__(
            'azurerm:devtestlab/v20180915:Disk',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Disk resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Disk(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop