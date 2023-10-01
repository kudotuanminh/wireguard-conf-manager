package main

import (
	"errors"
	"net/http"
	_ "wg-conf-manager/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initAPI() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		profiles := v1.Group("/profiles")
		{
			profiles.GET("/", listProfiles)
			profiles.GET("/:profileName", listProfileByName)
			profiles.POST("/new", createProfile)
			profiles.DELETE("/:profileName", deleteProfileByName)
			profiles.GET("/:profileName/getconf", getConfig)
		}
		v1.POST("/apply", applyServerConfig)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}

// @Summary		List all profiles
// @Description	Responds with the list of all client profiles as JSON
// @Tags		profiles
// @Produce		json
// @Success		200	{array} []Client
// @Failure		404	{object} error
// @Router		/profiles [get]
func listProfiles(c *gin.Context) {
	db := initDB()
	profiles, err := getProfiles(db)
	if err == nil {
		c.JSON(http.StatusOK, profiles)
	} else {
		handleErr(err)
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
}

// @Summary		List profile by name
// @Description	Responds with the client profile as JSON
// @Tags		profiles
// @Produce		json
// @Param		profileName path string true "Profile Name"
// @Success		200	{object} Client
// @Failure		404	{object} error
// @Router		/profiles/{profileName} [get]
func listProfileByName(c *gin.Context) {
	db := initDB()
	profileName := c.Param("profileName")
	profile, err := getProfile(db, profileName)
	if err == nil {
		c.JSON(http.StatusOK, profile)
	} else {
		handleErr(err)
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
}

// @Summary		Create new profile
// @Description	Responds with the newly created client profile as JSON
// @Tags		profiles
// @Produce		json
// @Param		profileName formData string true "Profile Name"
// @Param		clientIP formData string true "Client IP"
// @Param		allowedIPs formData string true "Allowed IPs"
// @Success		200	{object} Client
// @Failure		400	{object} error
// @Router		/profiles/new [post]
func createProfile(c *gin.Context) {
	db := initDB()
	profileName := c.PostForm("profileName")
	clientIP := c.PostForm("clientIP")
	allowedIPs := c.PostForm("allowedIPs")
	if profileName != "" && clientIP != "" && allowedIPs != "" {
		profile, err := newProfile(db, profileName, clientIP, allowedIPs)
		if err == nil {
			c.JSON(http.StatusOK, profile)
		} else {
			handleErr(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
		}
	} else {
		err := errors.New("invalid request")
		handleErr(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

// @Summary		Delete profile by name
// @Description	Responds with the deleted client profile as JSON
// @Tags		profiles
// @Produce 	json
// @Param		profileName path string true "Profile Name"
// @Success		200 {object} Client
// @Failure		404 {object} error
// @Router		/profiles/{profileName} [delete]
func deleteProfileByName(c *gin.Context) {
	db := initDB()
	profileName := c.Param("profileName")
	profile, err := getProfile(db, profileName)
	if err == nil {
		err := deleteProfile(db, profileName)
		if err == nil {
			c.JSON(http.StatusOK, profile)
		} else {
			handleErr(err)
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
	} else {
		handleErr(err)
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
}

// @Summary		Get profile config
// @Description	Responds with the client profile config as plain text
// @Tags		profiles
// @Produce		text/plain
// @Param		profileName path string true "Profile Name"
// @Success		200
// @Failure		404	{object} error
// @Router		/profiles/{profileName}/getconf [get]
func getConfig(c *gin.Context) {
	db := initDB()
	profileName := c.Param("profileName")
	configPath, err := genClientConf(db, profileName)
	if err == nil {
		c.FileAttachment(configPath, profileName+".conf")
	} else {
		handleErr(err)
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
}

// @Summary		Apply server config
// @Description	Responds with the request status
// @Tags		server
// @Success		200
// @Router		/apply [post]
func applyServerConfig(c *gin.Context) {
	db := initDB()
	configPath, err := genServerConf(db)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"filePath": configPath})
	} else {
		handleErr(err)
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
}
