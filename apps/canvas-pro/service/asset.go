package service

type AssetService struct{}

var assetService *AssetService

func init() {
	assetService = &AssetService{}

}

func GetAssetService() *AssetService {
	return assetService
}
