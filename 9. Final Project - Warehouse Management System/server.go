package main

import (
	config "warehouse-management-system-eFishery/config/database"
<<<<<<< HEAD
=======
	// handler "warehouse-management-system-eFishery/handler"
	// repository "warehouse-management-system-eFishery/repository"
	// routes "warehouse-management-system-eFishery/routes"
	// services "warehouse-management-system-eFishery/services"
>>>>>>> 376269eea8d5fea2c9daa01c5d52d28f50c3a08f

	handlerSupplier "warehouse-management-system-eFishery/handler/supplier"
	repositorySupplier "warehouse-management-system-eFishery/repository/supplier"
	routesSupplier "warehouse-management-system-eFishery/routes/supplier"
	servicesSupplier "warehouse-management-system-eFishery/services/supplier"

	handlerWarehouse "warehouse-management-system-eFishery/handler/warehouse"
	repositoryWarehouse "warehouse-management-system-eFishery/repository/warehouse"
	routesWarehouse "warehouse-management-system-eFishery/routes/warehouse"
	servicesWarehouse "warehouse-management-system-eFishery/services/warehouse"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.Connect()
	config.Migrate()
	e := echo.New()

	e.Use((middleware.Logger()))
	e.Use(middleware.Recover())

	// Warehouse Section
	warehouseProductRepository := repositoryWarehouse.NewWarehouseProductRepository(config.DB)
	warehouseProductService := servicesWarehouse.NewWarehouseProductService(warehouseProductRepository)
	warehouseProductHandler := handlerWarehouse.NewWarehouseProductHandler(warehouseProductService)
	routesWarehouse.RoutesProducts(e, warehouseProductHandler)

	warehouseWorkerRepository := repositoryWarehouse.NewWarehouseWorkersRepository(config.DB)
	warehouseWorkerService := servicesWarehouse.NewWarehouseWorkerService(warehouseWorkerRepository)
	warehouseWorkerHandler := handlerWarehouse.NewWarehouseWorkerHandler(warehouseWorkerService)
	routesWarehouse.RoutesWorkers(e, warehouseWorkerHandler)

	warehouseOrderRepository := repositoryWarehouse.NewWarehouseOrderRepository(config.DB)
	warehouseOrderService := servicesWarehouse.NewWarehouseOrderService(warehouseOrderRepository)
	warehouseOrderHandler := handlerWarehouse.NewWarehouseOrdersHandler(warehouseOrderService)
	routesWarehouse.RouteOreders(e, warehouseOrderHandler)

	warehouseRepository := repositoryWarehouse.NewWarehouseRepository(config.DB)
	warehouseService := servicesWarehouse.NewWarehouseService(warehouseRepository)
	warehouseHandler := handlerWarehouse.NewWarehouseHandler(warehouseService)
	routesWarehouse.Routes(e, warehouseHandler)

	warehouseAddressRepository := repositoryWarehouse.NewWarehouseAddressRepository(config.DB)
	warehouseAddressService := servicesWarehouse.NewWarehouseAddressService(warehouseAddressRepository)
	warehouseAddressHandler := handlerWarehouse.NewWarehouseAddressHandler(warehouseAddressService)
	routesWarehouse.RoutesAddress(e, warehouseAddressHandler)

	warehouseCategoryRepository := repositoryWarehouse.NewWarehouseCategoriesRepository(config.DB)
	warehouseCategoryService := servicesWarehouse.NewWarehouseCategoryService(warehouseCategoryRepository)
	warehouseCategoryHandler := handlerWarehouse.NewWarehouseCategoryHandler(warehouseCategoryService)
	routesWarehouse.RoutesCategory(e, warehouseCategoryHandler)

	warehouseRolesRepository := repositoryWarehouse.NewWarehouseRolesRepository(config.DB)
	warehouseRolesService := servicesWarehouse.NewWarehouseRolesService(warehouseRolesRepository)
	warehouseRolesHandler := handlerWarehouse.NewWarehouseRolesHandler(warehouseRolesService)
	routesWarehouse.RoutesRoles(e, warehouseRolesHandler)

	warehouseSectionRepository := repositoryWarehouse.NewWarehouseSectionRepository(config.DB)
	warehouseSectionService := servicesWarehouse.NewWarehouseSectionService(warehouseSectionRepository)
	warehouseSectionHandler := handlerWarehouse.NewWarehouseSectionHandler(warehouseSectionService)
	routesWarehouse.RoutesSection(e, warehouseSectionHandler)

	supplierRepository := repositorySupplier.NewSupplierRepository(config.DB)
	supplierService := servicesSupplier.NewSupplierService(supplierRepository)
	supplierHandler := handlerSupplier.NewSupplierHandler(supplierService)
	routesSupplier.RoutesSupplier(e, supplierHandler)

	supplierAddressRepository := repositorySupplier.NewSupplierAddressRepository(config.DB)
	supplierAddressService := servicesSupplier.NewSupplierAddressService(supplierAddressRepository)
	supplierAddressHandler := handlerSupplier.NewSupplierAddressHandler(supplierAddressService)
	routesSupplier.RoutesSupplierAddress(e, supplierAddressHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
