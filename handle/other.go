package handle

import (
	"github.com/gin-gonic/gin"
	"go-jrebel-license/util"
	"net/http"
)

func PingHandler(ctx *gin.Context) {
	salt := ctx.Query("salt")
	if salt == "" {
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}
	xmlContent := "<PingResponse><message></message><responseCode>OK</responseCode><salt>" + salt + "</salt></PingResponse>"
	xmlSignature := util.JrebelRsaSign([]byte(xmlContent))
	ctx.String(200, "<!-- "+xmlSignature+" -->\n"+xmlContent)
}

func ObtainTicketHandler(ctx *gin.Context) {
	salt := ctx.Query("salt")
	username := ctx.Query("userName")
	prolongationPeriod := "607875500"
	if salt == "" || username == "" {
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}
	xmlContent := "<ObtainTicketResponse><message></message><prolongationPeriod>" + prolongationPeriod + "</prolongationPeriod><responseCode>OK</responseCode><salt>" + salt + "</salt><ticketId>1</ticketId><ticketProperties>licensee=" + username + "\tlicenseType=0\t</ticketProperties></ObtainTicketResponse>"
	xmlSignature := util.JrebelRsaSign([]byte(xmlContent))
	ctx.String(200, "<!-- "+xmlSignature+" -->\n"+xmlContent)
}

func ReleaseTicketHandler(ctx *gin.Context) {
	salt := ctx.Query("salt")
	if salt == "" {
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}
	xmlContent := "<ReleaseTicketResponse><message></message><responseCode>OK</responseCode><salt>" + salt + "</salt></ReleaseTicketResponse>"
	xmlSignature := util.JrebelRsaSign([]byte(xmlContent))
	ctx.String(200, "<!-- "+xmlSignature+" -->\n"+xmlContent)
}
