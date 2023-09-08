package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/amir-qasemi/shop-discount/internal/cart"
	"github.com/amir-qasemi/shop-discount/internal/config"
	"github.com/amir-qasemi/shop-discount/internal/db"
	"github.com/amir-qasemi/shop-discount/internal/discount"
	"github.com/amir-qasemi/shop-discount/internal/lock"
	"github.com/amir-qasemi/shop-discount/internal/order"
	"github.com/amir-qasemi/shop-discount/internal/server"
	"github.com/amir-qasemi/shop-discount/internal/user"
	"github.com/amir-qasemi/shop-discount/internal/util"
)

func main() {
	// Get command-line arguments
	var configPath string
	flag.StringVar(&configPath, "config", "../config/config.yml", "path to config file")
	flag.Parse()

	// Load config
	config, err := config.New(configPath)
	if err != nil {
		log.Fatalln("Cannot load config file")
	}

	// Connect to DB
	db, err := db.Connect(config.DbConfig)
	if err != nil {
		log.Fatalln("Cannot load config file")
	}

	// setup required services
	ls := lock.NewInMemLockStore()

	// Setup server
	srv := server.New(config.ServerConfig)

	// Setup Cart
	var cartService cart.Service = &cart.DummyService{}

	// Setup User
	var userService user.Service = &user.DummyService{}

	// Setup discount
	discountRep := util.Ptr(discount.NewDummyRepository(db))
	discountService := &discount.AdHocDiscountService{DiscountRepository: discountRep, LockStore: ls, OrderService: &order.DummyService{}}
	discountController := discount.NewController(discountService, cartService, userService)
	discountController.Setup(srv)

	// Start server
	srv.Echo.Logger.Fatal(srv.Echo.Start(fmt.Sprintf("%v:%v", config.ServerConfig.Host, config.ServerConfig.Port)))
}
