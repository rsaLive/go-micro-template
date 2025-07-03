package socialconfig

import (
	configData "go-micro-template/config"
	"github.com/google/wire"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
	"os"
)

// ProviderSet 包含所有社交媒体平台配置的提供者
var ProviderSet = wire.NewSet(NewSocialConfigProvider)

type Provider struct {
	YouTubeConfig *oauth2.Config
}

// NewSocialConfigProvider 创建一个包含所有社交媒体平台配置的提供者
func NewSocialConfigProvider(confData *configData.AppConfig) *Provider {
	b, err := os.ReadFile("./conf/youtube_client_secret.json")
	if err != nil {
		panic(err)
	}

	youtubeConfig, err := google.ConfigFromJSON(b, youtube.YoutubeUploadScope)
	if err != nil {
		panic(err)
	}

	return &Provider{
		YouTubeConfig: youtubeConfig,
	}
}
