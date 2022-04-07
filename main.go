package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3" //直接的な記述が無いが、インポートしたいものに対しては"_"を頭につける決まり
)

type Bss struct {
	gorm.Model
	Name    string
	Content string `form:"content" binding:"required"`
}

func gormConnect() *gorm.DB {
	db, err := gorm.Open("sqlite3", "/tmp/bss.db")
	if err != nil {
		panic(err.Error())
	}
	return db
}

//DB初期化
func dbInit() {
	db := gormConnect()

	defer db.Close()
	db.AutoMigrate(&Bss{}) //構造体に基づいてテーブル作成
}

//追加
//new
func dbInsert(name, content string) {
	db := gormConnect()

	defer db.Close()
	db.Create(&Bss{Name: name, Content: content})
}

//DB更新
func dbUpdate(id int, Text string) {
	db := gormConnect()
	var bss Bss
	db.First(&bss, id)
	bss.Content = Text
	db.Save(&bss)
	db.Close()
}

//全部取得
func dbGetAll() []Bss {
	db := gormConnect()

	defer db.Close()
	var bss []Bss
	//FindでDB名を指定、取得した後、orderとして登録順に並び替え
	db.Order("created_at desc").Find(&bss)
	return bss
}

//ひとつだけ
func dbGetOne(id int) Bss {
	db := gormConnect()
	var bss Bss
	db.First(&bss, id)
	db.Close()
	return bss
}

func dbDelete(id int) {
	db := gormConnect()
	var bss Bss
	db.First(&bss, id)
	db.Delete(&bss)
	db.Close()
}

func main() {
	router := gin.Default()
	//router.LoadHTMLGlob("views/*.html")

	dbInit()

	router.GET("/", func(c *gin.Context) {
		bsss := dbGetAll()
		c.JSON(200, gin.H{
			"bss": bsss,
		})
	})
	//new ?name= & content =
	router.POST("/new", func(c *gin.Context) {
		//var form Bss
		//content?content=432 この場合4323
		//id := c.Query("test")
		name := c.Query("name")
		content := c.Query("content")
		fmt.Println(content)
		dbInsert(name, content)
		c.Redirect(302, "/")
	})

	router.GET("/detail/:id", func(c *gin.Context) {
		n := c.Param("id")
		//文字列を数値型に変換
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		bss := dbGetOne(id)
		c.JSON(200, gin.H{"bss": bss})
	})
	//更新
	router.POST("/update/:id", func(c *gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("Error")
		}
		content := c.Query("content")
		dbUpdate(id, content)
		c.Redirect(302, "/")
	})

	//削除 確認
	router.GET("/delete_check/:id", func(c *gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("Error")
		}
		bss := dbGetOne(id)
		c.JSON(200, gin.H{
			"bss-one": bss,
		})
	})

	//削除
	router.POST("/delete/:id", func(c *gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("Error")
		}
		dbDelete(id)
		c.Redirect(302, "/")
	})

	router.Run(":3000")
}
