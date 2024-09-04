package commons

import "errors"

var UserNotFound = errors.New("User not found, cannot continue processing")
var RecipeNotFound = errors.New("Recipe not found, cannot continue processing")
