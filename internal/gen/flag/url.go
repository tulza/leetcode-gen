package gencmd_flag

import "lcgen/internal/utils"

func GetSlugFromURL(url string) (string, error) {
	// fetch problem title from route
	_, routes, _ := utils.ParseURL(url)
	slugName := routes[1]

	return slugName, nil
}