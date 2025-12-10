package handlers

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mustafa91ameen/prjalgo/backend/internal/dtos"
	"github.com/mustafa91ameen/prjalgo/backend/internal/response"
	"github.com/mustafa91ameen/prjalgo/backend/internal/services"
)

type TeamMemberHandler struct {
	teamMemberService *services.TeamMemberService
}

func NewTeamMemberHandler(teamMemberService *services.TeamMemberService) *TeamMemberHandler {
	return &TeamMemberHandler{
		teamMemberService: teamMemberService,
	}
}

// GetAll handles GET /team-members
func (h *TeamMemberHandler) GetAll(c *gin.Context) {
	var pagination dtos.PaginationQuery
	if err := c.ShouldBindQuery(&pagination); err != nil {
		response.ValidationError(c, err)
		return
	}
	pagination.Normalize()

	teamMembers, err := h.teamMemberService.GetAll(c.Request.Context(), pagination.Limit, pagination.Offset())
	if err != nil {
		response.InternalError(c, "failed to fetch team members")
		return
	}

	response.Success(c, teamMembers)
}

// GetByID handles GET /team-members/:id
func (h *TeamMemberHandler) GetByID(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid team member id")
		return
	}

	teamMember, err := h.teamMemberService.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrTeamMemberNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to fetch team member")
		return
	}

	response.Success(c, teamMember)
}

// GetByProjectID handles GET /projects/:id/team-members
func (h *TeamMemberHandler) GetByProjectID(c *gin.Context) {
	projectID, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid project id")
		return
	}

	teamMembers, err := h.teamMemberService.GetByProjectID(c.Request.Context(), projectID)
	if err != nil {
		response.InternalError(c, "failed to fetch team members")
		return
	}

	response.Success(c, teamMembers)
}

// GetByUserID handles GET /users/:id/team-members
func (h *TeamMemberHandler) GetByUserID(c *gin.Context) {
	userID, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid user id")
		return
	}

	teamMembers, err := h.teamMemberService.GetByUserID(c.Request.Context(), userID)
	if err != nil {
		response.InternalError(c, "failed to fetch team members")
		return
	}

	response.Success(c, teamMembers)
}

// Create handles POST /team-members
func (h *TeamMemberHandler) Create(c *gin.Context) {
	var req dtos.CreateTeamMember
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}

	teamMember, err := h.teamMemberService.Create(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, "failed to create team member")
		return
	}

	response.Created(c, teamMember)
}

// Delete handles DELETE /team-members/:id
func (h *TeamMemberHandler) Delete(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid team member id")
		return
	}

	err = h.teamMemberService.Delete(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrTeamMemberNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to delete team member")
		return
	}

	response.Success(c, nil)
}

// parseID extracts an int64 ID from the URL path
func (h *TeamMemberHandler) parseID(c *gin.Context, param string) (int64, error) {
	idStr := c.Param(param)
	return strconv.ParseInt(idStr, 10, 64)
}
