package handlers

func (h Handlers) routes() {
	apiV1 := h.router.Group("/api/v1")

	emailServiceGroup := apiV1.Group("/emails")
	{
		emailServiceGroup.POST("/api/v1/send", h.send)
		emailServiceGroup.GET("/api/v1/lookup/:id", h.lookupStatus)
		emailServiceGroup.GET("/api/v1/stats", h.getStats)
	}
}
