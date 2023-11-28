package controllers

import (
	"gin-api-rest/database"
	"gin-api-rest/models"

	"github.com/gin-gonic/gin"
)

func ExibeTodosAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "E ai " + nome + ", tudo beleza?",
	})
}

func CriaAluno(c *gin.Context) {
	var novoAluno models.Aluno
	if err := c.ShouldBindJSON(&novoAluno); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Create(&novoAluno)
	c.JSON(200, novoAluno)
}

func BuscaAlunoPorID(c *gin.Context) {
	id := c.Params.ByName("id")
	var aluno models.Aluno
	database.DB.First(&aluno, id)
	if aluno.ID == 0 {
		c.JSON(404, gin.H{
			"Not found": "Aluno não encontrado",
		})
		return
	}
	c.JSON(200, aluno)
}

func DeletaAluno(c *gin.Context) {
	id := c.Params.ByName("id")
	var aluno models.Aluno
	database.DB.Delete(&aluno, id)
	c.JSON(200, gin.H{
		"data": "Aluno deletado com sucesso",
	})
}

func EditaAluno(c *gin.Context) {
	id := c.Params.ByName("id")
	var aluno models.Aluno
	database.DB.First(&aluno, id)
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(200, aluno)
}

func BuscaAlunoPorCPF(c *gin.Context) {
	cpf := c.Param("cpf")
	var aluno models.Aluno
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)
	if aluno.ID == 0 {
		c.JSON(404, gin.H{
			"Not found": "Aluno não encontrado",
		})
		return
	}
	c.JSON(200, aluno)
}
