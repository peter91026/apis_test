package purchase

import (
	"eirc.app/internal/v1/resolver/purchase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Presenter interface {
	Created(ctx *gin.Context)
	List(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Updated(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type presenter struct {
	PurchaseResolver purchase.Resolver
}

func New(db *gorm.DB) Presenter {
	return &presenter{
		PurchaseResolver: purchase.New(db),
	}
}
