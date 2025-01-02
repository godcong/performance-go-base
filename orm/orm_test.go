package orm

import (
	"context"
	"fmt"
	"math/rand"
	"sync/atomic"
	"testing"

	"entgo.io/ent/dialect/sql"
	"github.com/bwmarrin/snowflake"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"xorm.io/core"
	"xorm.io/xorm"

	"performance-go-base/orm/ent"
)

const per = 10000

var idgen, _ = snowflake.NewNode(1)

// 准备测试数据
func setupTestData(b *testing.B) (*gorm.DB, *ent.Client, *xorm.Engine) {
	gdb := initGormDB()
	xeng := initXormEngine()
	ecli := initEntClient()
	ctx := context.Background()
	// 清理现有数据
	//gdb.Exec("TRUNCATE TABLE users, posts, tags CASCADE")

	b.Log("Start generating test data...")

	// 生成标签数据
	//tags := make([]*Tag, 10)
	//entTags := make([]*ent.Tag, 10)
	//for i := 0; i < 10; i++ {
	//	// GORM tag
	//	tag := &Tag{
	//		Name: fmt.Sprintf("tag_%d", i),
	//	}
	//	gdb.Create(tag)
	//	tags[i] = tag
	//
	//	// Ent tag
	//	entTag := ecli.Tag.Create().
	//		SetName(fmt.Sprintf("tag_%d", i)).
	//		SetDescription(fmt.Sprintf("description_%d", i)).
	//		SaveX(ctx)
	//	entTags[i] = entTag
	//}

	// 生成用户和帖子数据
	for i := 1; i <= per; i++ {
		// GORM user 和 posts
		guser := &User{
			ID:       int(idgen.Generate().Int64()),
			Username: fmt.Sprintf("user_%d", i),
			Email:    fmt.Sprintf("user_%d@example.com", i),
		}
		gdb.Create(guser)

		//// 为每个用户创建1-3个帖子
		//postCount := rand.Intn(3) + 1
		//for j := 0; j < postCount; j++ {
		//	post := &Post{
		//		Title:    fmt.Sprintf("post_%d_%d", i, j),
		//		Content:  fmt.Sprintf("content_%d_%d", i, j),
		//		AuthorID: user.ID,
		//	}
		//	gdb.Create(post)
		//
		//	// 随机添加1-3个标签
		//	tagCount := rand.Intn(3) + 1
		//	selectedTags := make([]*Tag, 0, tagCount)
		//	for k := 0; k < tagCount; k++ {
		//		selectedTags = append(selectedTags, tags[rand.Intn(len(tags))])
		//	}
		//	gdb.Model(post).Association("Tags").Append(selectedTags)
		//}
		xuser := &User{
			Username: "test user",
			Email:    "test@example.com",
			Password: "password",
		}
		xeng.ID(int(idgen.Generate().Int64())).Insert(xuser)

		// Ent user 和 posts
		ecli.User.Create().
			SetID(int(idgen.Generate().Int64())).
			SetUsername(fmt.Sprintf("user_%d", i)).
			SetEmail(fmt.Sprintf("user_%d@example.com", i)).
			SetPassword("password").
			SaveX(ctx)

		//// 为每个用户创建1-3个帖子
		//for j := 0; j < postCount; j++ {
		//	post := ecli.Post.Create().
		//		SetTitle(fmt.Sprintf("post_%d_%d", i, j)).
		//		SetContent(fmt.Sprintf("content_%d_%d", i, j)).
		//		SetAuthor(entUser).
		//		SaveX(ctx)
		//
		//	// 随机添加1-3个标签
		//	tagCount := rand.Intn(3) + 1
		//	selectedEntTags := make([]*ent.Tag, 0, tagCount)
		//	for k := 0; k < tagCount; k++ {
		//		selectedEntTags = append(selectedEntTags, entTags[rand.Intn(len(entTags))])
		//	}
		//	post.Update().AddTags(selectedEntTags...).SaveX(ctx)
		//}

	}

	b.Log("The test data is generated")
	return gdb, ecli, xeng
}

func BenchmarkORMCreate(b *testing.B) {
	// 初始化数据库连接
	db := initGormDB()
	var id int64
	b.Run("GORM-Create", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			user := &User{
				ID:       int(atomic.AddInt64(&id, 1)),
				Username: "test user",
				Email:    "test@example.com",
				Password: "password",
			}
			db.Create(user)
		}
	})

	xeng := initXormEngine()
	b.Run("XORM-Create", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			user := &User{
				ID:       int(atomic.AddInt64(&id, 1)),
				Username: "test user",
				Email:    "test@example.com",
				Password: "password",
			}
			xeng.Insert(user)
		}
	})

	client := initEntClient()
	ctx := context.Background()
	b.Run("Ent-Create", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			client.User.Create().
				SetID(int(atomic.AddInt64(&id, 1))).
				SetUsername("test user").
				SetEmail("test@example.com").
				SetPassword("password").
				SaveX(ctx)
		}
	})

}

func BenchmarkORMQuery(b *testing.B) {
	db, client, xeng := setupTestData(b)
	ctx := context.Background()
	b.ResetTimer()

	// 简单查询
	b.Run("GORM-SimpleQuery", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			id := rand.Intn(10000) + 1
			var user User
			db.First(&user, id)
		}
	})

	b.Run("XORM-SimpleQuery", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			id := rand.Intn(10000) + 1
			var user User
			_, _ = xeng.ID(id).Get(&user)
		}
	})

	b.Run("Ent-SimpleQuery", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			id := rand.Intn(10000) + 1
			user, _ := client.User.Get(ctx, id)
			_ = user
		}
	})

	// 关联查询
	//b.Run("GORM-QueryWithRelations", func(b *testing.B) {
	//	for i := 0; i < b.N; i++ {
	//		id := rand.Intn(10000) + 1
	//		var user User
	//		db.Preload("Posts.Tags").First(&user, id)
	//	}
	//})

	//b.Run("Ent-QueryWithRelations", func(b *testing.B) {
	//	for i := 0; i < b.N; i++ {
	//		id := rand.Intn(10000) + 1
	//		client.User.Query().
	//			Where(user.ID(id)).
	//			WithPosts(func(q *ent.PostQuery) {
	//				q.WithTags()
	//			}).
	//			FirstX(ctx)
	//	}
	//})
}

func BenchmarkORMDelete(b *testing.B) {
	db, client, xeng := setupTestData(b)
	ctx := context.Background()

	b.ResetTimer()

	// 简单删除
	b.Run("GORM-SimpleDelete", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			id := rand.Intn(10000) + 1
			db.Delete(&User{}, id)
		}
	})

	b.Run("XORM-SimpleDelete", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			id := rand.Intn(10000) + 1
			xeng.ID(id).Delete(&User{})
		}
	})

	b.Run("Ent-SimpleDelete", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			id := rand.Intn(10000) + 1
			_ = client.User.DeleteOneID(id).Exec(ctx)
		}
	})

	//// 级联删除
	//b.Run("GORM-DeleteWithRelations", func(b *testing.B) {
	//	for i := 0; i < b.N; i++ {
	//		id := rand.Intn(10000) + 1
	//		db.Select("Posts").Delete(&User{}, id)
	//	}
	//})
	//
	//b.Run("Ent-DeleteWithRelations", func(b *testing.B) {
	//	for i := 0; i < b.N; i++ {
	//		id := rand.Intn(10000) + 1
	//		client.User.DeleteOneID(id).
	//			ExecX(ctx)
	//	}
	//})
}

//func BenchmarkORMManyToMany(b *testing.B) {
//	//db := initGormDB()
//
//	//b.Run("GORM-ManyToMany", func(b *testing.B) {
//	//	for i := 0; i < b.N; i++ {
//	//		var user User
//	//		db.Model(&user).Association("Tags").Append(&Tag{Name: "test"})
//	//	}
//	//})
//	//
//	//b.Run("Ent-ManyToMany", func(b *testing.B) {
//	//	client := initEntClient()
//	//	ctx := context.Background()
//	//	for i := 0; i < b.N; i++ {
//	//		//tag := client.Tag.Create().
//	//		//	SetName("test").
//	//		//	SaveX(ctx)
//	//		//client.User.UpdateOneID(1).
//	//		//	AddTags(tag).
//	//		//	SaveX(ctx)
//	//	}
//	//})
//}

// 初始化辅助函数
func initGormDB() *gorm.DB {
	dsn := "file::memory:?cache=shared&mode=memory"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.Migrator().AutoMigrate(&User{})
	return db
}

func initEntClient() *ent.Client {
	drv, err := sql.Open("sqlite3", "file::memory:?cache=shared&mode=memory")
	if err != nil {
		panic(err)
	}

	c := ent.NewClient(ent.Driver(drv))
	c.Schema.Create(context.Background())
	return c
}

func initXormEngine() *xorm.Engine {
	engine, err := xorm.NewEngine("sqlite3", "file::memory:?cache=shared&mode=memory")
	if err != nil {
		panic(err)
	}
	engine.SetMapper(core.GonicMapper{})
	err = engine.Sync(new(User))
	if err != nil {
		panic(err)
	}
	engine.SetDisableGlobalCache(true)
	return engine
}
