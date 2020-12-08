package handler

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/lz1998/ecust_library/config"
	"github.com/lz1998/ecust_library/dto"
	"github.com/lz1998/ecust_library/model/admin"
)

func CreateAdmin(c *gin.Context) {
	req := &dto.CreateAdminReq{}

	if err := c.Bind(req); err != nil {
		c.String(http.StatusBadRequest, "bad request, not protobuf")
		return
	}
	admins := req.GetAdmins()
	if admins == nil {
		c.String(http.StatusBadRequest, "bad request, book is nil")
		return
	}

	for _, a := range admins {
		if err := admin.CreateAdmin(a.Username, a.Password); err != nil {
			c.String(http.StatusInternalServerError, "create error")
			return
		}
	}
	resp := &dto.CreateBookResp{}
	Return(c, resp)
}

func ListAdmin(c *gin.Context) {
	req := &dto.ListAdminReq{}

	if err := c.Bind(req); err != nil {
		c.String(http.StatusBadRequest, "bad request, not protobuf")
		return
	}

	usernames := req.GetUsernames()
	if usernames == nil {
		usernames = make([]string, 0)
	}

	admins, err := admin.ListAdmin(req.Usernames)
	if err != nil {
		c.String(http.StatusInternalServerError, "db error")
		return
	}

	resp := &dto.ListAdminResp{
		Admins: convertAdminsModelToProto(admins),
	}
	Return(c, resp)
}

func UpdateAdmin(c *gin.Context) {
	req := &dto.UpdateAdminReq{}

	if err := c.Bind(req); err != nil {
		c.String(http.StatusBadRequest, "bad request, not protobuf")
		return
	}

	for _, a := range req.Admins {
		if err := admin.UpdateAdmin(a.Username, a.Password, a.Status); err != nil {
			c.String(http.StatusInternalServerError, "db error")
			return
		}
	}
	resp := &dto.UpdateAdminResp{}
	Return(c, resp)
}

func Login(c *gin.Context) {
	req := &dto.LoginReq{}

	if err := c.Bind(req); err != nil {
		c.String(http.StatusBadRequest, "bad request, not protobuf")
		return
	}

	ecustAdmin, err := admin.GetAdminByUsername(req.GetUsername())
	if err != nil {
		resp := &dto.LoginResp{
			Success: false,
			Msg:     "get admin by username error",
			Token:   "",
		}
		Return(c, resp)
		return
	}

	// TODO 加密
	if ecustAdmin.Password == req.GetPassword() {
		token, err := GenerateJwtTokenString(ecustAdmin)
		if err != nil {
			resp := &dto.LoginResp{
				Success: false,
				Msg:     "generate token error",
				Token:   "",
			}
			Return(c, resp)
			return
		}

		resp := &dto.LoginResp{
			Success: true,
			Msg:     "ok",
			Token:   token,
		}
		Return(c, resp)
		return
	}

	resp := &dto.LoginResp{
		Success: false,
		Msg:     "password error",
		Token:   "",
	}
	Return(c, resp)
}

func convertAdminModelToProto(modelAdmin *admin.EcustAdmin) *dto.EcustAdmin {
	return &dto.EcustAdmin{
		Id:        modelAdmin.ID,
		Username:  modelAdmin.Username,
		Password:  modelAdmin.Password,
		Status:    modelAdmin.Status,
		CreatedAt: modelAdmin.CreatedAt.Unix(),
		UpdatedAt: modelAdmin.UpdatedAt.Unix(),
	}
}

func convertAdminsModelToProto(modelAdmins []*admin.EcustAdmin) []*dto.EcustAdmin {
	admins := make([]*dto.EcustAdmin, 0)
	for _, modelAdmin := range modelAdmins {
		admins = append(admins, convertAdminModelToProto(modelAdmin))
	}
	return admins
}

func GenerateJwtTokenString(admin *admin.EcustAdmin) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["expire"] = time.Now().Add(time.Hour * time.Duration(24)).Unix()
	claims["id"] = admin.ID
	claims["username"] = admin.Username
	claims["password"] = admin.Password
	claims["status"] = admin.Status
	claims["createdAt"] = admin.CreatedAt.Unix()
	claims["updatedAt"] = admin.UpdatedAt.Unix()
	token.Claims = claims
	return token.SignedString(config.JwtSecret)
}
