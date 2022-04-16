package global

import (
	"gin-vue-admin/utils/timer"

	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"

	"gin-vue-admin/config"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GVA_DB     *gorm.DB
	GVA_PWD    string
	GVA_REDIS  *redis.Client
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
	//GVA_LOG    *oplogging.Logger
	GVA_LOG        *zap.Logger
	GVA_Timer      timer.Timer = timer.NewTimerTask()
	GVA_K8SCLIENTS []*kubernetes.Clientset
)
