package Controllers

import (
	"bytes"
	"html"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseBody struct {
	Message string                 `json:"message"`
	Body    map[string]interface{} `json:"body"`
}

func RequestBodyLogger(c *gin.Context) string {
	requestBody, _ := ioutil.ReadAll(c.Request.Body)
	rdr1 := ioutil.NopCloser(bytes.NewBuffer(requestBody))
	rdr2 := ioutil.NopCloser(bytes.NewBuffer(requestBody))
	c.Request.Body = rdr2
	return readBody(rdr1)
}
func readBody(reader io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	s := buf.String()
	return s
}
func ValidationResponse(c *gin.Context, message string) {
	response := ResponseBody{
		Message: message,
		Body:    map[string]interface{}{},
	}
	c.JSON(http.StatusUnprocessableEntity, response)
}

func NoDataFoundResponse(c *gin.Context, message string) {
	response := ResponseBody{
		Message: message,
		Body:    map[string]interface{}{},
	}
	c.JSON(http.StatusNotFound, response)
}

func InternalServerErrorResponse(c *gin.Context, message string) {
	response := ResponseBody{
		Message: message,
		Body:    map[string]interface{}{},
	}
	c.JSON(http.StatusInternalServerError, response)
}
func SuccessResponseXml(c *gin.Context, output []byte) {
	c.Header("Content-Type", "application/xml")
	c.String(http.StatusOK, html.UnescapeString(string(output)))
}
