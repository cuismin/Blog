package Tool

import (
	"github.com/satori/go.uuid"
	"Blog/models"
	"fmt"
)

//添加文章
func AddArticle() {
	Atype := []string{"Java","Python","Web","Go","Solidity","Truffle","HyperLedgerFabric","Mysql","Docker"};
	for data:= range Atype {
		uuid, _ := uuid.NewV4()
		var article models.Article
		article.Uuid  = uuid.String()     //序号
		article.Title = Atype[data]+"使用"       //标题
		article.Content = Atype[data]   //内容
		article.CreateDate = models.GetCurrentDate()     //创建时间
		article.Publisher  = "James"     //发布者
		article.Forward    = 0       //转发次数
		article.ArticleType = &models.ArticleType{Uuid:"4502a843-cb71-44fd-9a2e-0e9c87a6a040"}    //文章类型
		article.CreateUser = "James"  //创建人
		article.SetTop = false         //是否置顶
		article.PrioritySort = "11" //优先排序输出 在不置顶的时候有用 或者大家都在置顶的时候安装 当前输出
		article.User = &models.User{Uuid:"James"}       //设置一对多关系
		id := models.SaveData(&article)
		fmt.Println(id)
	}
}

//添加文章类型
func AddArticleType() {
	Atype := []string{"Java","Python","Web","Go","Solidity","Truffle","HyperLedgerFabric","Mysql","Docker"};
	for data:= range Atype {
		uuid, _ := uuid.NewV4()
		var articleType models.ArticleType
		articleType.Uuid = uuid.String()
		articleType.TypeName = Atype[data]
		articleType.CreateDate = models.GetCurrentDate()
		id := models.SaveData(&articleType)
		fmt.Println(id)
	}
}

//添加用户
func AddUser() {
	Users := []string{"Java","Python","Web","Go","Solidity","Truffle","James","Mysql","Docker"};
	for data:= range Users {
		var user models.User
		user.Uuid = Users[data]
		user.CreateDate = models.GetCurrentDate()
		user.Age = 25
		user.Sex = 0
		user.UserName = "James"
		user.HeadPortrait = "././"
		user.Nickname = "一花一世界"
		user.RealName = "cuism"
		id := models.SaveData(&user)
		fmt.Println(id)
	}
}

func InitDataBaseData(){
	AddUser()
	AddArticleType()
	AddArticle()
}


