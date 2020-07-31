# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class ManagerExtendedInfo(pulumi.CustomResource):
    etag: pulumi.Output[str]
    """
    The etag of the resource.
    """
    kind: pulumi.Output[str]
    """
    The Kind of the object. Currently only Series8000 is supported
    """
    name: pulumi.Output[str]
    """
    The name of the object.
    """
    properties: pulumi.Output[dict]
    """
    The extended info properties.
      * `algorithm` (`str`) - Represents the encryption algorithm used to encrypt the keys. None - if Key is saved in plain text format. Algorithm name - if key is encrypted
      * `encryption_key` (`str`) - Represents the CEK of the resource.
      * `encryption_key_thumbprint` (`str`) - Represents the Cert thumbprint that was used to encrypt the CEK.
      * `integrity_key` (`str`) - Represents the CIK of the resource.
      * `portal_certificate_thumbprint` (`str`) - Represents the portal thumbprint which can be used optionally to encrypt the entire data before storing it.
      * `version` (`str`) - The version of the extended info being persisted.
    """
    type: pulumi.Output[str]
    """
    The hierarchical type of the object.
    """
    def __init__(__self__, resource_name, opts=None, etag=None, kind=None, name=None, properties=None, resource_group_name=None, __props__=None, __name__=None, __opts__=None):
        """
        The extended info of the manager.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: The etag of the resource.
        :param pulumi.Input[str] kind: The Kind of the object. Currently only Series8000 is supported
        :param pulumi.Input[str] name: The manager name
        :param pulumi.Input[dict] properties: The extended info properties.
        :param pulumi.Input[str] resource_group_name: The resource group name

        The **properties** object supports the following:

          * `algorithm` (`pulumi.Input[str]`) - Represents the encryption algorithm used to encrypt the keys. None - if Key is saved in plain text format. Algorithm name - if key is encrypted
          * `encryption_key` (`pulumi.Input[str]`) - Represents the CEK of the resource.
          * `encryption_key_thumbprint` (`pulumi.Input[str]`) - Represents the Cert thumbprint that was used to encrypt the CEK.
          * `integrity_key` (`pulumi.Input[str]`) - Represents the CIK of the resource.
          * `portal_certificate_thumbprint` (`pulumi.Input[str]`) - Represents the portal thumbprint which can be used optionally to encrypt the entire data before storing it.
          * `version` (`pulumi.Input[str]`) - The version of the extended info being persisted.
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

            __props__['etag'] = etag
            __props__['kind'] = kind
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['type'] = None
        super(ManagerExtendedInfo, __self__).__init__(
            'azurerm:storsimple/v20170601:ManagerExtendedInfo',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing ManagerExtendedInfo resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ManagerExtendedInfo(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop