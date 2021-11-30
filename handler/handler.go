package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/ekokurniadi/terbilang"
	"github.com/ekokurniadi/terbilang_demo/helper"
	"github.com/ekokurniadi/terbilang_demo/input"
	"github.com/gin-gonic/gin"
)

type routerHandler struct {
}

func NewRouterHandler() *routerHandler {
	return &routerHandler{}
}

func (h *routerHandler) Process(c *gin.Context) {
	var input input.NominalInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.ApiResponse("Invalid", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatRupiah := terbilang.Init()
	angka, err := strconv.Atoi(input.Angka)
	if err != nil {
		fmt.Println(err)
	}
	hasil := formatRupiah.Convert(int64(angka))

	response := helper.ApiResponse("Success", http.StatusOK, "success", strings.Title(hasil))
	c.JSON(http.StatusOK, response)

}

func (h *routerHandler) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index", nil)
}
