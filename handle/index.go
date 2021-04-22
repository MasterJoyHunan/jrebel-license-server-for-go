package handle

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-jrebel-license/response"
	"go-jrebel-license/util"
	"strconv"
)

func Index(ctx *gin.Context) {
	ctx.Writer.Write([]byte(`
		<h1>Hello World</h1>
	`))
}

func JrebelLeasesHandler(ctx *gin.Context) {
	randomness := ctx.PostForm("randomness")
	username := ctx.PostForm("username")
	guid := ctx.PostForm("guid")
	offline := ctx.PostForm("offline")
	clientTime := ctx.PostForm("clientTime")
	validFrom, validUntil := "null", "null"

	if offline == "true" {
		ct, _ := strconv.Atoi(clientTime)
		validFrom = clientTime
		validUntil = strconv.Itoa(ct + 180*24*60*60*1000)
	}

	signature := util.JrebelSign(randomness, guid, validFrom, validUntil, offline)

	validFromInt, _ := strconv.Atoi(validFrom)
	validUntilInt, _ := strconv.Atoi(validUntil)

	res := response.JrebelResponse{
		ServerVersion:         "3.2.4",
		ServerProtocolVersion: "1.1",
		ServerGuid:            "a1b4aea8-b031-4302-b602-670a990272cb",
		GroupType:             "managed",
		Id:                    1,
		LicenseType:           1,
		EvaluationLicense:     false,
		ServerRandomness:      "H2ulzLlh7E0=",
		SeatPoolType:          "standalone",
		StatusCode:            "SUCCESS",
		Signature:             signature,
		Offline:               offline == "true",
		ValidFrom:             validFromInt,
		ValidUntil:            validUntilInt,
		Company:               username,
		OrderId:               "",
		ZeroIds:               []string{},
		LicenseValidFrom:      1490544001000,
		LicenseValidUntil:     1691839999000,
	}
	indent, _ := json.MarshalIndent(res, "", "    ")
	fmt.Println(string(indent))
	ctx.JSON(200, res)
}

func JrebelLeases1Handler(ctx *gin.Context) {
	username := ctx.Query("username")
	h := gin.H{
		"serverVersion":         "3.2.4",
		"serverProtocolVersion": "1.1",
		"serverGuid":            "a1b4aea8-b031-4302-b602-670a990272cb",
		"groupType":             "managed",
		"statusCode":            "SUCCESS",
		"msg":                   "",
		"statusMessage":         "",
	}
	if username != "" {
		h["company"] = username
	}
	ctx.JSON(200, h)
}

func JrebelValidateHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"serverVersion":         "3.2.4",
		"serverProtocolVersion": "1.1",
		"serverGuid":            "a1b4aea8-b031-4302-b602-670a990272cb",
		"groupType":             "managed",
		"statusCode":            "SUCCESS",
		"company":               "Administrator",
		"canGetLease":           true,
		"licenseType":           1,
		"evaluationLicense":     false,
		"seatPoolType":          "standalone",
	})
}
