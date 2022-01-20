package main

import (
	"github.com/labstack/echo/v4"
	"github.com/qwerty22121998/slink/cmd/api/configs"
	"github.com/qwerty22121998/slink/cmd/api/controllers"
	"github.com/qwerty22121998/slink/pkg/adapter/firebase"
	"github.com/qwerty22121998/slink/pkg/repositories"
	"github.com/qwerty22121998/slink/pkg/services"
)

func main() {
	config := configs.Global
	e := echo.New()
	route := e.Group("/api")

	fb := firebase.NewFireBaseAdapter(firebase.Config{
		CertPath: config.GoogleAppCertPath,
	})

	repo := repositories.NewShortLinkRepository(fb.DB)

	//c, err := repo.Find(context.Background(), "ok")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%#v", c)

	//now := time.Now()
	//if _, err := repo.Set(context.Background(), "facebook", &models.ShortLink{
	//	Name:       "facebook",
	//	URL:        "https://www.facebook.com/",
	//	Short:      "facebook",
	//	ExpiredAt:  &now,
	//	ClickCount: 0,
	//}); err != nil {
	//	panic(err)
	//}
	//
	//return

	service := services.NewShortLinkService(repo)

	slController := controllers.NewShortLinkController(service)

	slController.Register(route)

	e.Logger.Fatal(e.Start(":" + config.Port))
}
