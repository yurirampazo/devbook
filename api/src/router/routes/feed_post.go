package routes

import (
	"api/src/controllers"
	"net/http"
)

var feedPostRoutes = []Route {
	{
		URI: "/posts",
		Method: http.MethodPost,
		Function: controllers.CreatePost,
		RequireAuth: true,
	},
	{
		URI: "/posts",
		Method: http.MethodGet,
		Function: controllers.FindPosts,
		RequireAuth: true,
	},
	{
		URI: "/posts/{id}",
		Method: http.MethodGet,
		Function: controllers.GetPostById,
		RequireAuth: true,
	},
	{
		URI: "/posts/{id}",
		Method: http.MethodPut,
		Function: controllers.EditPost,
		RequireAuth: true,
	},	
	{
		URI: "/posts/{id}",
		Method: http.MethodDelete,
		Function: controllers.DeletePost,
		RequireAuth: true,
	},

	{
		URI: "/users/{userId}/posts",
		Method: http.MethodGet,
		Function: controllers.GetUserPosts,
		RequireAuth: true,
	},
	
	{
		URI: "/posts/{postId}/like",
		Method: http.MethodPost,
		Function: controllers.LikePost,
		RequireAuth: true,
	},
	
	
}