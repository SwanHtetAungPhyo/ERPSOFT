package main

import (
    "errors"
    "time"

    "github.com/SwanHtetAungPhyo/esfor/internal/database"
    "github.com/SwanHtetAungPhyo/esfor/internal/handlers"
    middlewares "github.com/SwanHtetAungPhyo/esfor/internal/middleware"
    "github.com/SwanHtetAungPhyo/esfor/internal/models"
    repositories "github.com/SwanHtetAungPhyo/esfor/internal/repo"
    "github.com/SwanHtetAungPhyo/esfor/internal/services"
    "github.com/SwanHtetAungPhyo/esfor/internal/utils"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cache"
    "github.com/gofiber/fiber/v2/middleware/compress"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/gofiber/fiber/v2/middleware/encryptcookie"
    "github.com/gofiber/fiber/v2/middleware/healthcheck"
    "github.com/gofiber/fiber/v2/middleware/limiter"
    "github.com/gofiber/fiber/v2/middleware/monitor"
    "github.com/gofiber/fiber/v2/middleware/pprof"
    "github.com/gofiber/fiber/v2/middleware/etag"
    _ "github.com/SwanHtetAungPhyo/esfor/docs" // Import generated docs
    fiberSwagger "github.com/swaggo/fiber-swagger" // fiber-swagger middleware
    "go.uber.org/zap"
)

// @title Fiber API
// @version 1.0
// @description This is a sample server for a Fiber API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8006
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
    appLog := initializeLogger()
    env := utils.NewEnvConfig()
    initializeDatabase(env.DSN, appLog)
    app := initializeFiberApp()
    app.Use(
        cors.New(cors.Config{
            AllowOrigins: "http://localhost:3001, http://localhost:3000",
            AllowHeaders: "Origin, Content-Type, Accept, Authorization",
            AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
        }),
        compress.New(
            compress.Config{
                Level: compress.LevelBestSpeed,
            }),
        cache.New(cache.Config{
            Expiration:   2 * 60 * 60,
            CacheControl: true,
        }),
        encryptcookie.New(encryptcookie.Config{
            Key: "secret-key",
        }),
        etag.New(),
        limiter.New(limiter.Config{
            Max:               20,
            Expiration:        30 * time.Second,
            LimiterMiddleware: limiter.SlidingWindow{},
        }),
    )
    app.Get("/health", healthcheck.New())
    app.Get("/monitor", monitor.New())
    app.Use(pprof.New(pprof.Config{Prefix: "/profile"}))
    app.Get("/swagger/*", fiberSwagger.WrapHandler) // Serve Swagger documentation
    registerRoutes(app)
    startServer(app, env.PORT, appLog)
}

func initializeLogger() *zap.Logger {
    return utils.GetLogger()
}

func initializeDatabase(dsn string, appLog *zap.Logger) {
    database.Init(dsn)
    if err := database.DB.AutoMigrate(&models.Role{},
        &models.Employee{},
        &models.Student{},
        &models.Section{},
        &models.Course{},
        &models.Announcement{},
        &models.CourseAnnouncement{},
    ); err != nil {
        appLog.Fatal("Error in schema migration", zap.Error(err))
    }
}

func initializeFiberApp() *fiber.App {
    return fiber.New(fiber.Config{
        ErrorHandler: func(c *fiber.Ctx, err error) error {
            code := fiber.StatusInternalServerError
            message := "internal server error"
            if fiberError, ok := err.(*fiber.Error); ok {
                code = fiberError.Code
                message = fiberError.Message
            } else if errors.Is(err, fiber.ErrUnauthorized) {
                code = fiber.StatusUnauthorized
                message = "Unauthorized"
            } else {
                message = err.Error()
            }
            utils.GetLogger().Error("Request error", zap.Error(err))
            return utils.JsonResp(c, code, message)
        },
    })
}

func registerRoutes(app *fiber.App) {
    app.Use(middlewares.LoggingMiddleware())

    roleRepo := repositories.NewRoleRepository(database.DB)
    roleService := services.NewRoleService(roleRepo)
    roleController := handlers.NewRoleController(roleService)
    roleController.RegisterRoutes(app)

    employeeRepo := repositories.NewEmployeeRepository(database.DB)
    employeeService := services.NewEmployeeService(employeeRepo)
    employeeController := handlers.NewEmployeeController(employeeService)
    employeeController.RegisterRoutes(app)

    studentRepo := repositories.NewStudentRepository(database.DB)
    studentService := services.NewStudentService(studentRepo)
    studentController := handlers.NewStudentController(studentService)
    studentController.RegisterRoutes(app)

    sectionRepo := repositories.NewSectionRepository(database.DB)
    sectionService := services.NewSectionService(sectionRepo)
    sectionController := handlers.NewSectionController(sectionService)
    sectionController.RegisterRoutes(app)

    courseRepo := repositories.NewCourseRepository(database.DB)
    courseService := services.NewCourseService(courseRepo, employeeRepo)
    courseController := handlers.NewCourseController(courseService)
    courseController.RegisterRoutes(app)

    courseAnnouncementRepo := repositories.NewCourseAnnouncementRepository(database.DB)
    announcementRepo := repositories.NewAnnouncementRepository(database.DB)
    announcementService := services.NewAnnouncementService(announcementRepo, courseAnnouncementRepo, employeeRepo, database.DB)
    announcementController := handlers.NewAnnouncementController(announcementService)
    announcementController.RegisterRoutes(app)

    app.Use(middlewares.JwtMiddleware())
}

func startServer(app *fiber.App, port string, appLog *zap.Logger) {
    appLog.Info("Listening at the port", zap.String("port", port))
    if err := app.ListenTLS(":" + port, "cert.pem", "key.pem"); err != nil {
        appLog.Fatal("Error in init server", zap.Error(err))
    }
}