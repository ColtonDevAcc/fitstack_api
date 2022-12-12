package delivery

import (
	"fmt"
	"net/http"

	"firebase.google.com/go/v4/auth"
	"github.com/VooDooStack/FitStackAPI/api/middleware"
	"github.com/VooDooStack/FitStackAPI/domain/dto"
	healthLogs "github.com/VooDooStack/FitStackAPI/domain/health_logs"
	"github.com/VooDooStack/FitStackAPI/domain/user"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

// ArticleHandler  represent the httphandler for article
type UserHandler struct {
	UUsecase user.UserUsecase
}

// NewArticleHandler will initialize the articles/ resources endpoint
func NewUserHandler(g *gin.RouterGroup, us user.UserUsecase, client *auth.Client) {
	handler := &UserHandler{
		UUsecase: us,
	}
	g.GET("/get", middleware.AuthJWT(client), handler.FetchUser)
	g.POST("/refresh-token", handler.RefreshToken)
	g.POST("/signup", handler.SignUp)
	g.DELETE("/delete", middleware.AuthJWT(client), handler.DeleteUser)
	g.POST("/signin", middleware.AuthJWT(client), handler.SignInWithToken)
	g.POST("/signin-email-password", handler.SignInWithEmailAndPassword)
	g.POST("/update-avatar", middleware.AuthJWT(client), handler.UpdateUserAvatar)
	g.GET("/fetch-profile", handler.fetchUserProfile)
	g.GET("/profile", middleware.AuthJWT(client), handler.getUserProfile)
	g.POST("/statistics", middleware.AuthJWT(client), handler.UpdateUserStatistics)
	g.GET("/statistics", middleware.AuthJWT(client), handler.GetUserStatistics)
	g.GET("/statistics/snapshot", middleware.AuthJWT(client), handler.GetUserStatisticsSnapshot)
	g.POST("/statistic", middleware.AuthJWT(client), handler.GetUserHealthLog)
}

func (ur *UserHandler) FetchUser(c *gin.Context) {
	user, err := ur.UUsecase.GetByUuid(c.Param("uuid"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})

		return
	}

	c.JSON(http.StatusOK, user)
}

func (ur *UserHandler) SignUp(c *gin.Context) {
	requestedUser := dto.UserSignUp{}
	err := c.ShouldBindJSON(&requestedUser)
	if err != nil {
		logrus.Error(err)

		c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
		return
	}

	user, err := ur.UUsecase.SignUp(&requestedUser, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (ur *UserHandler) DeleteUser(c *gin.Context) {
	err := ur.UUsecase.Delete(c.Param("uuid"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	c.JSON(http.StatusOK, nil)
}

func (ur *UserHandler) SignInWithToken(c *gin.Context) {
	token := c.GetString("token")

	user, err := ur.UUsecase.SignInWithToken(c, token)
	if err != nil {
		logrus.Error(err)

		c.JSON(http.StatusInternalServerError, ResponseError{Message: "error" + err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (ur *UserHandler) SignInWithEmailAndPassword(c *gin.Context) {
	loginInEmailAndPassword := &dto.LoginInEmailAndPassword{}
	err := c.ShouldBindJSON(loginInEmailAndPassword)
	if err != nil {
		logrus.Error(err)

		c.JSON(http.StatusInternalServerError, ResponseError{Message: "error" + err.Error()})
		return
	}

	user, err := ur.UUsecase.SignInWithEmailAndPassword(c, loginInEmailAndPassword)
	if err != nil {
		logrus.Error(err)

		c.JSON(http.StatusInternalServerError, ResponseError{Message: "error" + err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (ur *UserHandler) RefreshToken(c *gin.Context) {
	var refreshToken struct {
		RefreshToken string `json:"refresh_token"`
	}

	err := c.ShouldBindJSON(&refreshToken)
	if err != nil {
		logrus.Error(err)

		c.JSON(http.StatusInternalServerError, ResponseError{Message: "error" + err.Error()})
		return
	}

	token, err := ur.UUsecase.RefreshToken(c, refreshToken.RefreshToken)
	if err != nil {
		logrus.Error(err)

		c.JSON(http.StatusInternalServerError, ResponseError{Message: "error" + err.Error()})
		return
	}

	c.JSON(http.StatusOK, token)
}

func (ur *UserHandler) UpdateUserAvatar(c *gin.Context) {
	uuid := c.GetString("uuid")
	src, file, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	url, err := ur.UUsecase.UpdateUserAvatar(c, uuid, file, src)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	c.String(http.StatusOK, url)
}

func (ur *UserHandler) fetchUserProfile(c *gin.Context) {
	type uuid struct {
		UUID string `json:"id"`
	}

	userUUID := uuid{}
	err := c.ShouldBindJSON(&userUUID)
	if err != nil {
		logrus.Error(err)

		c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
		return
	}

	profile, err := ur.UUsecase.GetUserProfile(userUUID.UUID)
	if err != nil {
		logrus.Error(err)

		c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}

func (ur *UserHandler) getUserProfile(c *gin.Context) {
	uuid := c.GetString("uuid")

	profile, err := ur.UUsecase.GetUserProfile(uuid)
	if err != nil {
		logrus.Error(err)

		c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}

func (ur *UserHandler) UpdateUserStatistics(c *gin.Context) {
	uuid := c.GetString("uuid")
	statistics := user.UserStatistic{}
	err := c.ShouldBindJSON(&statistics)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
		return
	}

	for i := range statistics.WeightLogs {
		statistics.WeightLogs[i].UserStatisticID = uuid
	}

	for i := range statistics.BodyFatPercentageLogs {
		statistics.BodyFatPercentageLogs[i].UserStatisticID = uuid
	}

	for i := range statistics.BodyMassIndexLogs {
		statistics.BodyMassIndexLogs[i].UserStatisticID = uuid
	}

	for i := range statistics.ActiveEnergyBurnedLogs {
		statistics.ActiveEnergyBurnedLogs[i].UserStatisticID = uuid
	}

	for i := range statistics.SleepAsleepLogs {
		statistics.SleepAsleepLogs[i].UserStatisticID = uuid
	}

	for i := range statistics.StepsLogs {
		statistics.StepsLogs[i].UserStatisticID = uuid
	}
	statistics.ID = uuid

	err = ur.UUsecase.UpdateUserStatistics(&statistics)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
		return
	}

	fmt.Println(statistics)

	c.JSON(http.StatusOK, nil)
}

func (ur *UserHandler) GetUserStatistics(c *gin.Context) {
	uuid := c.GetString("uuid")

	statistics, err := ur.UUsecase.GetUserStatistics(uuid)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, statistics)
}

func (ur *UserHandler) GetUserStatisticsSnapshot(c *gin.Context) {
	uuid := c.GetString("uuid")

	statistics, err := ur.UUsecase.GetUserStatisticsSnapshot(uuid)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, statistics)
}

func (ur *UserHandler) GetUserHealthLog(c *gin.Context) {
	type healthType struct {
		HealthType healthLogs.HealthDataType `json:"health_type"`
	}

	dataType := healthType{}
	err := c.ShouldBindJSON(&dataType)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
		return
	}
	uuid := c.GetString("uuid")

	statistics, err := ur.UUsecase.GetUserHealthLog(uuid, dataType.HealthType)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, statistics)
}
