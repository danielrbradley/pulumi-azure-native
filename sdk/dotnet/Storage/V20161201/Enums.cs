// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.Storage.V20161201
{
    /// <summary>
    /// Required for storage accounts where kind = BlobStorage. The access tier used for billing.
    /// </summary>
    [EnumType]
    public readonly struct AccessTier : IEquatable<AccessTier>
    {
        private readonly string _value;

        private AccessTier(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AccessTier Hot { get; } = new AccessTier("Hot");
        public static AccessTier Cool { get; } = new AccessTier("Cool");

        public static bool operator ==(AccessTier left, AccessTier right) => left.Equals(right);
        public static bool operator !=(AccessTier left, AccessTier right) => !left.Equals(right);

        public static explicit operator string(AccessTier value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AccessTier other && Equals(other);
        public bool Equals(AccessTier other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The protocol permitted for a request made with the account SAS.
    /// </summary>
    [EnumType]
    public readonly struct HttpProtocol : IEquatable<HttpProtocol>
    {
        private readonly string _value;

        private HttpProtocol(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static HttpProtocol Https_http { get; } = new HttpProtocol("https,http");
        public static HttpProtocol Https { get; } = new HttpProtocol("https");

        public static bool operator ==(HttpProtocol left, HttpProtocol right) => left.Equals(right);
        public static bool operator !=(HttpProtocol left, HttpProtocol right) => !left.Equals(right);

        public static explicit operator string(HttpProtocol value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is HttpProtocol other && Equals(other);
        public bool Equals(HttpProtocol other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Required. Indicates the type of storage account.
    /// </summary>
    [EnumType]
    public readonly struct Kind : IEquatable<Kind>
    {
        private readonly string _value;

        private Kind(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static Kind Storage { get; } = new Kind("Storage");
        public static Kind BlobStorage { get; } = new Kind("BlobStorage");

        public static bool operator ==(Kind left, Kind right) => left.Equals(right);
        public static bool operator !=(Kind left, Kind right) => !left.Equals(right);

        public static explicit operator string(Kind value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is Kind other && Equals(other);
        public bool Equals(Kind other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The signed permissions for the service SAS. Possible values include: Read (r), Write (w), Delete (d), List (l), Add (a), Create (c), Update (u) and Process (p).
    /// </summary>
    [EnumType]
    public readonly struct Permissions : IEquatable<Permissions>
    {
        private readonly string _value;

        private Permissions(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static Permissions R { get; } = new Permissions("r");
        public static Permissions D { get; } = new Permissions("d");
        public static Permissions W { get; } = new Permissions("w");
        public static Permissions L { get; } = new Permissions("l");
        public static Permissions A { get; } = new Permissions("a");
        public static Permissions C { get; } = new Permissions("c");
        public static Permissions U { get; } = new Permissions("u");
        public static Permissions P { get; } = new Permissions("p");

        public static bool operator ==(Permissions left, Permissions right) => left.Equals(right);
        public static bool operator !=(Permissions left, Permissions right) => !left.Equals(right);

        public static explicit operator string(Permissions value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is Permissions other && Equals(other);
        public bool Equals(Permissions other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The signed services accessible with the service SAS. Possible values include: Blob (b), Container (c), File (f), Share (s).
    /// </summary>
    [EnumType]
    public readonly struct SignedResource : IEquatable<SignedResource>
    {
        private readonly string _value;

        private SignedResource(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SignedResource B { get; } = new SignedResource("b");
        public static SignedResource C { get; } = new SignedResource("c");
        public static SignedResource F { get; } = new SignedResource("f");
        public static SignedResource S { get; } = new SignedResource("s");

        public static bool operator ==(SignedResource left, SignedResource right) => left.Equals(right);
        public static bool operator !=(SignedResource left, SignedResource right) => !left.Equals(right);

        public static explicit operator string(SignedResource value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SignedResource other && Equals(other);
        public bool Equals(SignedResource other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Gets or sets the sku name. Required for account creation; optional for update. Note that in older versions, sku name was called accountType.
    /// </summary>
    [EnumType]
    public readonly struct SkuName : IEquatable<SkuName>
    {
        private readonly string _value;

        private SkuName(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SkuName Standard_LRS { get; } = new SkuName("Standard_LRS");
        public static SkuName Standard_GRS { get; } = new SkuName("Standard_GRS");
        public static SkuName Standard_RAGRS { get; } = new SkuName("Standard_RAGRS");
        public static SkuName Standard_ZRS { get; } = new SkuName("Standard_ZRS");
        public static SkuName Premium_LRS { get; } = new SkuName("Premium_LRS");

        public static bool operator ==(SkuName left, SkuName right) => left.Equals(right);
        public static bool operator !=(SkuName left, SkuName right) => !left.Equals(right);

        public static explicit operator string(SkuName value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SkuName other && Equals(other);
        public bool Equals(SkuName other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
