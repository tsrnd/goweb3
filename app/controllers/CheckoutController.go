package controller

import "net/http"
import "strconv"
import "github.com/goweb3/app/shared/view"
import "github.com/gorilla/csrf"
import "github.com/goweb3/app/models"
import "github.com/jianfengye/web-golang/web/session"
import "github.com/goweb3/app/shared/cookie"
import service "github.com/goweb3/app/services"

func Checkout(w http.ResponseWriter, r *http.Request) {
	sess, _ := session.SessionStart(r, w)

	userId, _ := strconv.ParseUint(sess.Get("id"), 10, 32)

	cart := models.Cart{}
	cart.FindByUserID(uint(userId))
	cart.LoadCartProducts()
	for i, _ := range cart.CartProducts {
		cart.CartProducts[i].LoadProducts()
		cart.CartProducts[i].Product.LoadProductImage()
	}
	v := view.New(r)
	v.Vars[csrf.TemplateTag] = csrf.TemplateField(r)
	if content := cookie.GetMessage(w, r, "ErrorCheckout"); content != "" {
		v.Vars["titleMessage"] = "Error"
		v.Vars["contentMessage"] = content
	} else if content := cookie.GetMessage(w, r, "SuccessCheckout"); content != "" {
		v.Vars["titleMessage"] = "Success"
		v.Vars["contentMessage"] = content
	}
	v.Vars["cart"] = cart
	v.Vars["totalPrice"] = cart.TotalPrice()
	v.Name = "checkout/index"
	v.Render(w)
}

func CheckoutPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if err := service.CheckoutOrder(w, r); err != nil {
		cookie.SetMessage(w, "Order failed!", "ErrorCheckout")		
	} else {
		cookie.SetMessage(w, "Order successful! Thank you!", "SuccessCheckout")
	}
	http.Redirect(w, r, "/checkout", http.StatusSeeOther)	
}
