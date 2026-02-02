package main

import (
	"gin-api/controllers"
	"gin-api/database"
	"gin-api/model"
	"io"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupDasRotasDeTest() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func CriaAlunoMock() {
	aluno := model.Aluno{Nome: "aluno1", CPF: "11111111111", RG: "000000000"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	var aluno model.Aluno
	database.DB.Delete(&aluno, ID)
}

func Test_VerificaStatusCodeDaSaudacaoComParametro(t *testing.T) {
	assert := assert.New(t)

	r := SetupDasRotasDeTest()
	r.GET("/:nome", controllers.Saudacao)

	req, _ := http.NewRequest("GET", "/gui", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	assert.Equal(http.StatusOK, resposta.Code, "Deveriam ser iguais")

	mockDaResposta := `{"API diz":"E ai gui, tudo beleza?"}`
	respostaBody, _ := io.ReadAll(resposta.Body)
	assert.Equal(mockDaResposta, string(respostaBody))
}

func Test_ListandoTodosOsAlunosHandler(t *testing.T) {
	assert := assert.New(t)

	database.ConectComBancoDeDados()

	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupDasRotasDeTest()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	assert.Equal(http.StatusOK, resposta.Code)
}

func Test_BuscaAlunoPorCPFHandler(t *testing.T) {
	assert := assert.New(t)

	database.ConectComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupDasRotasDeTest()
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCpf)
	req, _ := http.NewRequest("GET", "/alunos/cpf/12345678901", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	assert.Equal(http.StatusOK, resposta.Code)
}
