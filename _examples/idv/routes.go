package main

func initializeRoutes() {

	// Use the setSessionStatus middleware for every route to set a flag
	// indicating whether the session has been created or not
	router.Use(setSessionStatus())

	// Handle the index route
	router.GET("/", showIndexPage)
	router.GET("/success", showSuccessPage)
	router.GET("/media", getMedia)
	router.GET("/privacy-policy", showPrivacyPolicyPage)
}
