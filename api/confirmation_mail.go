package api

import (
	"net/http"
	"github.com/Ignitus/ignitus-mailing-microservice/mailer"
	"github.com/Ignitus/ignitus-mailing-microservice/structure"
	"github.com/Ignitus/ignitus-mailing-microservice/utils"
	"github.com/gin-gonic/gin"
)

/* Fucntion to send confirmation mail to user. */

func ConfirmationMailAPI(c *gin.Context) {
	var userData structure.RequestBody
	err := userData.BindBody(c)
	if err != nil {
		utils.LogError("Error binding json body.", err)
		c.JSON(http.StatusBadRequest, structure.ErrorResponse("Invalid body."))
		return
	}
	sendTo := userData.ToEmail 
	emailLink := userData.VerificationLink
	htmlBody := mailer.GenerateConfirmationHTMLTemplate(emailLink)

	err = mailer.Mail(sendTo, mailer.ConfirmationMailSubject, htmlBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structure.ErrorResponse("Email was not sent."))
		return
	}

	c.JSON(http.StatusOK, structure.SuccessResponse("Successfully sent mail."))
}
