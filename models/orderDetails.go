
package models
import 	"go.mongodb.org/mongo-driver/bson/primitive"


//OrderDetails contain product information as well as quantity
type OrderDetails struct {
	UnitPrice   float64
	Quantity    int
	Discount    float64
	ProductInfo primitive.ObjectID
}