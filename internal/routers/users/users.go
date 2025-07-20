package users

import (
	"auth-rest/internal/dal"
	authm "auth-rest/internal/middlewares/auth"
	"auth-rest/internal/modules/auth"
	"auth-rest/internal/modules/users"
	"auth-rest/internal/routers/common"
	"auth-rest/internal/utils"
	"errors"
	"github.com/gofiber/fiber/v3"
	"strconv"
)

// HandleSingleUpdateSelf is handler for updating current user
//
// @Summary      Update Current User
// @Description  Updates info for current user
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerToken
// @Param        request body SelfUpdateRequest true "Self Update Request"
// @Success      200  {object}  common.SuccessResponse
// @Failure      401  {object}  common.ErrorResponse
// @Failure      500  {object}  common.ErrorResponse
// @Router       /user [post]
func HandleSingleUpdateSelf(ctx fiber.Ctx) error {
	currentUser := authm.UserFromContext(ctx)

	var request SelfUpdateRequest
	err := utils.DecodeJSONBody(ctx, &request)
	if err != nil {
		return err
	}

	repos := dal.FromContext(ctx)
	model := users.UserModel{
		ID:   currentUser.ID,
		Name: request.Name,
	}

	err = users.UpdateUser(repos.Users(), model)
	if err != nil {
		return err
	}
	return ctx.JSON(&common.SuccessResponse{Success: true})
}

// HandleSingleGetSelf is handler for fetching current user's info
//
// @Summary      Get Current User
// @Description  Gets info for current user
// @Tags         users
// @Produce      json
// @Security     BearerToken
// @Success      200  {object}  UserResponse
// @Failure      401  {object}  common.ErrorResponse
// @Failure      500  {object}  common.ErrorResponse
// @Router       /user [get]
func HandleSingleGetSelf(ctx fiber.Ctx) error {
	user := authm.UserFromContext(ctx)
	return ctx.JSON(ToResponse(user))
}

// HandleSingleDelete is handler for fetching current user's info
//
// @Summary      Delete User
// @Description  Deletes given user from database
// @Tags         users
// @Produce      json
// @Security     BearerToken
// @Param		 id   path      int      true   "User ID"
// @Success      200  {object}  common.SuccessResponse
// @Failure      401  {object}  common.ErrorResponse
// @Failure      500  {object}  common.ErrorResponse
// @Router       /users/{id} [delete]
func HandleSingleDelete(ctx fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return fiber.ErrNotFound
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return fiber.ErrNotFound
	}

	currentUser := authm.UserFromContext(ctx)
	// user cannot delete itself
	if idInt == currentUser.ID {
		return fiber.ErrBadRequest
	}

	repos := dal.FromContext(ctx)
	user, err := users.UserByID(repos.Users(), idInt)
	if err != nil {
		if !errors.Is(err, dal.ErrRecordNotFound) {
			return fiber.ErrInternalServerError
		}
		return fiber.ErrNotFound
	}
	// user cannot delete a user with higher role or equal role to him
	if !auth.HasHigherRole(currentUser.Role, user.Role) {
		return fiber.ErrForbidden
	}
	err = users.DeleteUser(repos.Users(), idInt)
	if err != nil {
		return err
	}
	return ctx.JSON(&common.SuccessResponse{Success: true})
}

// HandleSingleUpdate is handler for updating given user's info
//
// @Summary      Update User
// @Description  Updates given user info
// @Tags         users
// @Produce      json
// @Security     BearerToken
// @Param		 id      path      int               true   "User ID"
// @Param        request body      UserUpdateRequest true   "User Update Request"
// @Success      200  {object}  common.SuccessResponse
// @Failure      401  {object}  common.ErrorResponse
// @Failure      500  {object}  common.ErrorResponse
// @Router       /users/{id} [post]
func HandleSingleUpdate(ctx fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return fiber.ErrNotFound
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return fiber.ErrNotFound
	}

	var request UserUpdateRequest
	err = utils.DecodeJSONBody(ctx, &request)
	if err != nil {
		return err
	}

	currentUser := authm.UserFromContext(ctx)

	repos := dal.FromContext(ctx)
	model := users.UserModel{
		ID:   idInt,
		Name: request.Name,
	}
	// role update only by superadmin and only for others
	if currentUser.Role == auth.RoleSuper && request.Role != "" && currentUser.ID != idInt {
		// validate role
		if !auth.IsValidRole(request.Role) {
			return fiber.ErrBadRequest
		}
		if auth.HasHigherRole(currentUser.Role, request.Role) {
			model.Role = request.Role
		} else {
			return fiber.ErrForbidden
		}
	} else {
		if request.Role != "" {
			return fiber.ErrForbidden
		}
	}

	err = users.UpdateUser(repos.Users(), model)
	if err != nil {
		return err
	}
	return ctx.JSON(&common.SuccessResponse{Success: true})
}

// HandleSingleGet is handler for fetching given user's info
//
// @Summary      Get User
// @Description  Fetches given user's info
// @Tags         users
// @Produce      json
// @Security     BearerToken
// @Param		 id      path      int               true   "User ID"
// @Success      200  {object}  UserResponse
// @Failure      401  {object}  common.ErrorResponse
// @Failure      500  {object}  common.ErrorResponse
// @Router       /users/{id} [get]
func HandleSingleGet(ctx fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return fiber.ErrNotFound
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return fiber.ErrNotFound
	}
	repos := dal.FromContext(ctx)
	user, err := users.UserByID(repos.Users(), idInt)
	if err != nil {
		return err
	}
	return ctx.JSON(ToResponse(user))
}

// HandleCreate is handler for creating a new user
//
// @Summary      Create User
// @Description  Creates a new user
// @Tags         users
// @Produce      json
// @Security     BearerToken
// @Param        request body      UserCreateRequest true   "User Create Request"
// @Success      200  {object}  UserCreateResponse
// @Failure      401  {object}  common.ErrorResponse
// @Failure      500  {object}  common.ErrorResponse
// @Router       /users [post]
func HandleCreate(ctx fiber.Ctx) error {
	var request UserCreateRequest
	err := utils.DecodeJSONBody(ctx, &request)
	if err != nil {
		return err
	}
	if !utils.IsValidPhone(request.Phone) {
		return fiber.ErrBadRequest
	}
	repos := dal.FromContext(ctx)
	model := users.UserModel{
		Name:  request.Name,
		Phone: request.Phone,
		Role:  auth.RoleUser,
	}

	currentUser := authm.UserFromContext(ctx)

	// only superuser can create users with roles different from "user"
	if currentUser.Role == auth.RoleSuper && request.Role != "" {
		if !auth.IsValidRole(request.Role) {
			return fiber.ErrBadRequest
		}
		if auth.HasHigherRole(currentUser.Role, request.Role) {
			model.Role = request.Role
		} else {
			return fiber.ErrForbidden
		}
	} else {
		if request.Role != "" {
			return fiber.ErrForbidden
		}
	}

	model, err = users.CreateUser(repos.Users(), model)
	if err != nil {
		return err
	}
	return ctx.JSON(&UserCreateResponse{ID: model.ID, Success: true})
}

// HandleList is handler for listing all users
//
// @Summary      List Users
// @Description  Lists all users from database
// @Tags         users
// @Produce      json
// @Security     BearerToken
// @Success      200  {object}  UsersResponse
// @Failure      401  {object}  common.ErrorResponse
// @Failure      500  {object}  common.ErrorResponse
// @Router       /users [get]
func HandleList(ctx fiber.Ctx) error {
	repos := dal.FromContext(ctx)
	usersList, err := users.ListUsers(repos.Users())
	if err != nil {
		return err
	}
	result := make([]UserResponse, len(usersList))
	for i, user := range usersList {
		result[i] = ToResponse(user)
	}
	return ctx.JSON(result)
}
