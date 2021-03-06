package controller

import (
	"MareWood/sql"
	"MareWood/models"
	"MareWood/service/serviceRepository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func RepositoryFindAll(c *gin.Context) {

	var result []models.Repository

	isNormal := c.Query("isNormal")

	query := sql.DB.Order("created_at desc").Select("id,name")

	if isNormal == "1" {
		query = query.Where("status = ?", models.RepoStatusSuccess)
	}

	if query.Find(&result).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    sql.DB.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   result,
		"msg":    "查询成功",
	})

}

func RepositoryFind(c *gin.Context)  {

	pageNum, err := strconv.Atoi(c.Query("pageNum"))
	if err != nil {
		pageNum = 1
	}
	pageSize, err2 := strconv.Atoi(c.Query("pageSize"))
	if err2 != nil {
		pageSize = 8
	}

	result, err := models.FindRepository(pageNum, pageSize, map[string]interface{}{})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   result,
		"msg":    "查询成功",
	})
}

func RepositoryCreate(c *gin.Context) {

	var repository models.Repository

	if err := c.ShouldBindJSON(&repository); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    err.Error(),
		})
		return
	}

	err := repository.Create()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    err.Error(),
		})
		return
	}
	go func() {
		//克隆并更新记录
		out, err := serviceRepository.GitClone(strconv.Itoa(int(repository.ID)), repository.Url, repository.UserName, repository.Password)
		if err != nil {
			sql.DB.Model(&repository).
				Where("id = ?", repository.ID).
				Update("status", models.RepoStatusFail).
				Update("terminal_info", out)
			return
		}

		sql.DB.Model(&repository).
			Where("id = ?", repository.ID).
			Update("status", models.RepoStatusSuccess).
			Update("terminal_info", out)

	}()

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   "",
		"msg":    "创建成功,后台正在执行仓库克隆",
	})
}

func RepositoryDestroy(c *gin.Context) {

	id := c.Query("id")

	var jobCount int

	if sql.DB.Model(&models.Job{}).Where("repository_id = ?", id).Count(&jobCount).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   id,
			"msg":    sql.DB.Error.Error(),
		})
	}

	if jobCount > 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   jobCount,
			"msg":    "无法删除，还有" + strconv.Itoa(jobCount) + "个任务在使用此仓库",
		})
		return
	}

	var repository models.Repository

	if sql.DB.First(&repository, id).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    sql.DB.Error.Error(),
		})
		return
	}

	if repository.Status == models.RepoStatusProcessing {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    "无法删除正在克隆的仓库，请稍后再试",
		})
		return
	}

	//if repository.Status == models.RepoStatusSuccess {
	//
	//	err := serviceRepository.DeleteRepository(strconv.Itoa(int(repository.ID)))
	//
	//	if err != nil {
	//		c.JSON(http.StatusOK, gin.H{
	//			"status": false,
	//			"data":   "",
	//			"msg":    err.Error(),
	//		})
	//		return
	//	}
	//}

	if sql.DB.Delete(&repository).Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    sql.DB.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   id,
		"msg":    "删除成功",
	})

}

func RepositoryGitPull(c *gin.Context) {

	id := c.Query("id")
	out, err := serviceRepository.GitPull(id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   out,
		"msg":    "执行成功",
	})
}

func RepositoryDiscardChange(c *gin.Context)  {

	id := c.Query("id")
	out, err := serviceRepository.DiscardChange(id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   out,
		"msg":    "执行成功",
	})
}

func RepositoryDeleteDepend(c *gin.Context) {

	id := c.Query("id")

	err := serviceRepository.DeleteDepend(id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   "仓库依赖删除成功，相关任务执行的时候会重新安装",
		"msg":    "执行成功",
	})
}

func RepositoryBranch(c *gin.Context) {

	id := c.Query("id")

	branch, err := serviceRepository.GetBranch(id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   branch,
		"msg":    "执行成功",
	})

}

func RepositoryPruneBranch(c *gin.Context) {

	id := c.Query("id")

	out, err := serviceRepository.PruneBranch(id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    out,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   "",
		"msg":    "分支已经裁剪",
	})

}

func RepositoryScript(c *gin.Context) {
	id := c.Query("id")
	packageJson, err := serviceRepository.ParsePackageJson(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   packageJson.Scripts,
		"msg":    "",
	})
}

func RepositoryUpdateField(c *gin.Context)  {

	id := c.Query("id")
	field := c.Query("field")
	fieldContent := c.Query("fieldContent")

	err := new(models.Repository).UpdateFieldContent(id, field, fieldContent)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   id,
		"msg":    "更新成功",
	})
}