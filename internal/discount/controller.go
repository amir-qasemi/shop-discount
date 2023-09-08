package discount

import (
	"net/http"
	"time"

	"github.com/amir-qasemi/shop-discount/internal/cart"
	"github.com/amir-qasemi/shop-discount/internal/product"
	"github.com/amir-qasemi/shop-discount/internal/server"
	"github.com/amir-qasemi/shop-discount/internal/user"
	"github.com/labstack/echo/v4"
)

// Codes related to REST API for discount service(Could be easily replaced by gRPC endpoints or ...)
// Some structures like context(which should be absolutly present) are ommitted by choice to keep the code more simple and readable

// Controller for discount
type Controller struct {
	service     Service
	cartService cart.Service
	userService user.Service
}

func NewController(service Service, cartService cart.Service, userService user.Service) server.Controller {
	return &Controller{
		service:     service,
		cartService: cartService,
		userService: userService,
	}
}

// Setup  does all server related configs like routing and nessecary middleware
func (c *Controller) Setup(srv *server.Server) {
	g := srv.Echo.Group("/discount")

	// Creating discounts
	server.GroupWrapper(g, "/", server.MethodPut, c.createDiscount)
	server.GroupWrapper(g, "/new-user", server.MethodPut, c.createNewUserDiscount)
	server.GroupWrapper(g, "/product", server.MethodPut, c.createProductDiscount)
	server.GroupWrapper(g, "/min-product", server.MethodPut, c.createMinProductDiscount)

	// Applying and evaluting discounts
	server.GroupWrapper(g, "/apply", server.MethodPost, c.applyDiscount)
	server.GroupWrapper(g, "/is-eligibile", server.MethodGet, c.isEligilble)
	server.GroupWrapper(g, "/rollback", server.MethodPost, c.rollback)
}

type createNewUserDiscountRequest struct {
	Code     string `validate:"required"`
	Value    int    `validate:"required"`
	Unit     Unit   `validate:"required"`
	MaxVal   int    `validate:"required"`
	ValidNum int    `validate:"required"`
}

func (c *Controller) createNewUserDiscount(request createNewUserDiscountRequest, ctx echo.Context) error {
	discount := &NewUserDiscount{
		generalAdhocDiscount: generalAdhocDiscount{
			DiscountCode: request.Code,
			CreationTs:   time.Now(),
			ValidNum:     request.ValidNum,
			XUsages:      make(map[string]Usage),
			XUnit:        request.Unit,
			XValue:       request.Value,
			XMaxVal:      request.MaxVal,
		},
	}

	err := c.service.Save(discount)

	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	} else {
		return ctx.JSON(http.StatusOK, "discount created")
	}
}

type createProductDiscountRequest struct {
	Code      string `validate:"required"`
	Value     int    `validate:"required"`
	Unit      Unit   `validate:"required"`
	MaxVal    int    `validate:"required"`
	ValidNum  int    `validate:"required"`
	ProductId string `validate:"required"`
}

func (c *Controller) createProductDiscount(request createProductDiscountRequest, ctx echo.Context) error {
	discount := &ProductDiscount{
		generalAdhocDiscount: generalAdhocDiscount{
			DiscountCode: request.Code,
			CreationTs:   time.Now(),
			ValidNum:     request.ValidNum,
			XUsages:      make(map[string]Usage),
			XUnit:        request.Unit,
			XValue:       request.Value,
			XMaxVal:      request.MaxVal,
		},
		Product: product.Product{Id: request.ProductId},
	}

	err := c.service.Save(discount)

	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	} else {
		return ctx.JSON(http.StatusOK, "discount created")
	}
}

type createMinProductDiscountRequest struct {
	Code      string `validate:"required"`
	Value     int    `validate:"required"`
	Unit      Unit   `validate:"required"`
	MaxVal    int    `validate:"required"`
	ValidNum  int    `validate:"required"`
	ProductId string `validate:"required"`
	MinNumber int    `validate:"required"`
}

func (c *Controller) createMinProductDiscount(request createMinProductDiscountRequest, ctx echo.Context) error {
	discount := &MinProductDiscount{
		generalAdhocDiscount: generalAdhocDiscount{
			DiscountCode: request.Code,
			CreationTs:   time.Now(),
			ValidNum:     request.ValidNum,
			XUsages:      make(map[string]Usage),
			XUnit:        request.Unit,
			XValue:       request.Value,
			XMaxVal:      request.MaxVal,
		},
		Product:   product.Product{Id: request.ProductId},
		MinNumber: request.MinNumber,
	}

	err := c.service.Save(discount)

	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	} else {
		return ctx.JSON(http.StatusOK, "discount created")
	}
}

type createDiscountRequest struct {
	RuleDef string
}

// createDiscount a function to create a general discount(may be adhoc or rule based). Needs a DSL for discounts.
func (c *Controller) createDiscount(request createDiscountRequest, ctx echo.Context) error {
	return ctx.String(http.StatusNotImplemented, "Not implemented")
}

type applyDiscountRequest struct {
	DiscountCode string `validate:"required"`
	CartId       string `validate:"required"`
	Username     string `validate:"required,alphanumunicode"`
}

type applyDiscountResponse struct {
	cart *cart.Cart
}

func (c *Controller) applyDiscount(request applyDiscountRequest, ctx echo.Context) error {
	user, err := c.userService.GetUser(request.Username)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	cart, err := c.cartService.GetCartById(request.CartId)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	err = c.service.Apply(cart, request.DiscountCode, user)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	res := applyDiscountResponse{cart: cart}
	return ctx.JSON(http.StatusOK, res)
}

type isEligilbleRequest struct {
	DiscountCode string `validate:"required"`
	CartId       string `validate:"required"`
	Username     string `validate:"required,alphanumunicode"`
}

type isEligilbleResponse struct {
	result bool
	cart   *cart.Cart
}

func (c *Controller) isEligilble(request isEligilbleRequest, ctx echo.Context) error {
	user, err := c.userService.GetUser(request.Username)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	cart, err := c.cartService.GetCartById(request.CartId)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	isElig := c.service.IsEligible(cart, request.DiscountCode, user)

	res := isEligilbleResponse{cart: cart, result: isElig}
	return ctx.JSON(http.StatusOK, res)
}

type rollbackRequest struct {
	UsageId      string `validate:"required"`
	DiscountCode string `validate:"required"`
}

func (c *Controller) rollback(request rollbackRequest, ctx echo.Context) error {
	err := c.service.RollbackUsage(request.UsageId, request.DiscountCode)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, "rollback successfull")
}

// TODO: getdiscount controller
