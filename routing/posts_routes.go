package routing

import (
	"github.com/go-chi/chi/v5"
)

func (routing *Routing) configurePostsRouting(r chi.Router) {
	r.Route("/posts", func(r chi.Router) {
		r.Use(routing.middleware.JWTAuthMiddleware)
		r.Post("/", routing.handlers.PostHandler.CreatePostHandler)

		r.Route("/{postID}", func(r chi.Router) {
			r.Use(routing.middleware.PostsPathParamMiddleware)

			r.Get("/", routing.handlers.PostHandler.GetPostHandler)
			r.Patch("/", routing.middleware.CheckPostOwnershipMiddleware(moderatorRole, routing.handlers.PostHandler.UpdatePostHandler))
			r.Delete("/", routing.middleware.CheckPostOwnershipMiddleware(adminRole, routing.handlers.PostHandler.DeletePostHandler))
		})
	})
}
