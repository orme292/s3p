package conf

type ProviderConfig interface {
    Validate() error
}

type GetProviderConfigFunc func(p YamlProvider, b YamlBucket) (ProviderConfig, error)

const EmptyStr string = ""
