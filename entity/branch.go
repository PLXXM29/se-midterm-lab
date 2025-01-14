package entity
import(
	"gorm.io/gorm"
)

type Branch struct{
	gorm.Model
	Name string `valid:"required~name is require"`
	Email string `valid:"email~name is invalid"`
}