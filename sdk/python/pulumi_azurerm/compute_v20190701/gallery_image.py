# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class GalleryImage(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    Resource location
    """
    name: pulumi.Output[str]
    """
    Resource name
    """
    properties: pulumi.Output[dict]
    """
    Describes the properties of a gallery Image Definition.
      * `description` (`str`) - The description of this gallery Image Definition resource. This property is updatable.
      * `disallowed` (`dict`) - Describes the disallowed disk types.
        * `disk_types` (`list`) - A list of disk types.

      * `end_of_life_date` (`str`) - The end of life date of the gallery Image Definition. This property can be used for decommissioning purposes. This property is updatable.
      * `eula` (`str`) - The Eula agreement for the gallery Image Definition.
      * `hyper_v_generation` (`str`) - The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
      * `identifier` (`dict`) - This is the gallery Image Definition identifier.
        * `offer` (`str`) - The name of the gallery Image Definition offer.
        * `publisher` (`str`) - The name of the gallery Image Definition publisher.
        * `sku` (`str`) - The name of the gallery Image Definition SKU.

      * `os_state` (`str`) - This property allows the user to specify whether the virtual machines created under this image are 'Generalized' or 'Specialized'.
      * `os_type` (`str`) - This property allows you to specify the type of the OS that is included in the disk when creating a VM from a managed image. <br><br> Possible values are: <br><br> **Windows** <br><br> **Linux**
      * `privacy_statement_uri` (`str`) - The privacy statement uri.
      * `provisioning_state` (`str`) - The provisioning state, which only appears in the response.
      * `purchase_plan` (`dict`) - Describes the gallery Image Definition purchase plan. This is used by marketplace images.
        * `name` (`str`) - The plan ID.
        * `product` (`str`) - The product ID.
        * `publisher` (`str`) - The publisher ID.

      * `recommended` (`dict`) - The properties describe the recommended machine configuration for this Image Definition. These properties are updatable.
        * `memory` (`dict`) - Describes the resource range.
          * `max` (`float`) - The maximum number of the resource.
          * `min` (`float`) - The minimum number of the resource.

        * `v_cp_us` (`dict`) - Describes the resource range.

      * `release_note_uri` (`str`) - The release note uri.
    """
    tags: pulumi.Output[dict]
    """
    Resource tags
    """
    type: pulumi.Output[str]
    """
    Resource type
    """
    def __init__(__self__, resource_name, opts=None, gallery_name=None, location=None, name=None, properties=None, resource_group_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Specifies information about the gallery Image Definition that you want to create or update.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] gallery_name: The name of the Shared Image Gallery in which the Image Definition is to be created.
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[str] name: The name of the gallery Image Definition to be created or updated. The allowed characters are alphabets and numbers with dots, dashes, and periods allowed in the middle. The maximum length is 80 characters.
        :param pulumi.Input[dict] properties: Describes the properties of a gallery Image Definition.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[dict] tags: Resource tags

        The **properties** object supports the following:

          * `description` (`pulumi.Input[str]`) - The description of this gallery Image Definition resource. This property is updatable.
          * `disallowed` (`pulumi.Input[dict]`) - Describes the disallowed disk types.
            * `disk_types` (`pulumi.Input[list]`) - A list of disk types.

          * `end_of_life_date` (`pulumi.Input[str]`) - The end of life date of the gallery Image Definition. This property can be used for decommissioning purposes. This property is updatable.
          * `eula` (`pulumi.Input[str]`) - The Eula agreement for the gallery Image Definition.
          * `hyper_v_generation` (`pulumi.Input[str]`) - The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
          * `identifier` (`pulumi.Input[dict]`) - This is the gallery Image Definition identifier.
            * `offer` (`pulumi.Input[str]`) - The name of the gallery Image Definition offer.
            * `publisher` (`pulumi.Input[str]`) - The name of the gallery Image Definition publisher.
            * `sku` (`pulumi.Input[str]`) - The name of the gallery Image Definition SKU.

          * `os_state` (`pulumi.Input[str]`) - This property allows the user to specify whether the virtual machines created under this image are 'Generalized' or 'Specialized'.
          * `os_type` (`pulumi.Input[str]`) - This property allows you to specify the type of the OS that is included in the disk when creating a VM from a managed image. <br><br> Possible values are: <br><br> **Windows** <br><br> **Linux**
          * `privacy_statement_uri` (`pulumi.Input[str]`) - The privacy statement uri.
          * `purchase_plan` (`pulumi.Input[dict]`) - Describes the gallery Image Definition purchase plan. This is used by marketplace images.
            * `name` (`pulumi.Input[str]`) - The plan ID.
            * `product` (`pulumi.Input[str]`) - The product ID.
            * `publisher` (`pulumi.Input[str]`) - The publisher ID.

          * `recommended` (`pulumi.Input[dict]`) - The properties describe the recommended machine configuration for this Image Definition. These properties are updatable.
            * `memory` (`pulumi.Input[dict]`) - Describes the resource range.
              * `max` (`pulumi.Input[float]`) - The maximum number of the resource.
              * `min` (`pulumi.Input[float]`) - The minimum number of the resource.

            * `v_cp_us` (`pulumi.Input[dict]`) - Describes the resource range.

          * `release_note_uri` (`pulumi.Input[str]`) - The release note uri.
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

            if gallery_name is None:
                raise TypeError("Missing required property 'gallery_name'")
            __props__['gallery_name'] = gallery_name
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
        super(GalleryImage, __self__).__init__(
            'azurerm:compute/v20190701:GalleryImage',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing GalleryImage resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return GalleryImage(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop