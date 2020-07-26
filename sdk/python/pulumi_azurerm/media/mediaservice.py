# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class Mediaservice(pulumi.CustomResource):
    identity: pulumi.Output[dict]
    """
    The Managed Identity for the Media Services account.
      * `principal_id` (`str`) - The Principal ID of the identity.
      * `tenant_id` (`str`) - The Tenant ID of the identity.
      * `type` (`str`) - The identity type.
    """
    location: pulumi.Output[str]
    """
    The geo-location where the resource lives
    """
    name: pulumi.Output[str]
    """
    The name of the resource
    """
    properties: pulumi.Output[dict]
    """
    The resource properties.
      * `encryption` (`dict`) - The account encryption properties.
        * `key_vault_properties` (`dict`) - The properties of the key used to encrypt the account.
          * `current_key_identifier` (`str`) - The current key used to encrypt the Media Services account, including the key version.
          * `key_identifier` (`str`) - The URL of the Key Vault key used to encrypt the account. The key may either be versioned (for example https://vault/keys/mykey/version1) or reference a key without a version (for example https://vault/keys/mykey).

        * `type` (`str`) - The type of key used to encrypt the Account Key.

      * `media_service_id` (`str`) - The Media Services account ID.
      * `storage_accounts` (`list`) - The storage accounts for this resource.
        * `id` (`str`) - The ID of the storage account resource. Media Services relies on tables and queues as well as blobs, so the primary storage account must be a Standard Storage account (either Microsoft.ClassicStorage or Microsoft.Storage). Blob only storage accounts can be added as secondary storage accounts.
        * `type` (`str`) - The type of the storage account.

      * `storage_authentication` (`str`)
    """
    tags: pulumi.Output[dict]
    """
    Resource tags.
    """
    type: pulumi.Output[str]
    """
    The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
    """
    def __init__(__self__, resource_name, opts=None, identity=None, location=None, name=None, properties=None, resource_group_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        A Media Services account.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] identity: The Managed Identity for the Media Services account.
        :param pulumi.Input[str] location: The geo-location where the resource lives
        :param pulumi.Input[str] name: The Media Services account name.
        :param pulumi.Input[dict] properties: The resource properties.
        :param pulumi.Input[str] resource_group_name: The name of the resource group within the Azure subscription.
        :param pulumi.Input[dict] tags: Resource tags.

        The **identity** object supports the following:

          * `type` (`pulumi.Input[str]`) - The identity type.

        The **properties** object supports the following:

          * `encryption` (`pulumi.Input[dict]`) - The account encryption properties.
            * `key_vault_properties` (`pulumi.Input[dict]`) - The properties of the key used to encrypt the account.
              * `key_identifier` (`pulumi.Input[str]`) - The URL of the Key Vault key used to encrypt the account. The key may either be versioned (for example https://vault/keys/mykey/version1) or reference a key without a version (for example https://vault/keys/mykey).

            * `type` (`pulumi.Input[str]`) - The type of key used to encrypt the Account Key.

          * `storage_accounts` (`pulumi.Input[list]`) - The storage accounts for this resource.
            * `id` (`pulumi.Input[str]`) - The ID of the storage account resource. Media Services relies on tables and queues as well as blobs, so the primary storage account must be a Standard Storage account (either Microsoft.ClassicStorage or Microsoft.Storage). Blob only storage accounts can be added as secondary storage accounts.
            * `type` (`pulumi.Input[str]`) - The type of the storage account.

          * `storage_authentication` (`pulumi.Input[str]`)
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

            __props__['identity'] = identity
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
        super(Mediaservice, __self__).__init__(
            'azurerm:media:Mediaservice',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Mediaservice resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Mediaservice(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop