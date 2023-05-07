package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type business struct {
	ConfName  string `json:"conf_name" bson:"conf_name"`
	ConfValue string `json:"conf_value" bson:"conf_value"`
	//LastModifyBy string `json:"last_modify_by" bson:"last_modify_by"`
	//ModifyTime   string `json:"modify_time" bson:"modify_time"`
}
type DetailConfValue struct {
	EnvFlag string `json:"env_flag"`
	Value   string `json:"value"`
}

func main() {
	var apps = []string{"zeus-app"}
	mm := &mongodb{}
	mm.init("db_ares", "db_ares", "db_ares", "business-conf")
	for _, app := range apps {
		if err := mm.updateService(context.Background(), app, "PRE", "annualylbank"); err != nil {
			return
		}
	}

}

type mongodb struct {
	//client     *mongo.Client
	collection *mongo.Collection
}

func (m *mongodb) init(user, password, db, collection string) {
	option := options.Client().ApplyURI("mongodb://127.0.0.1:60017").SetConnectTimeout(2 * time.Second).SetAuth(options.Credential{Username: user, Password: password, AuthSource: db})
	client, err := mongo.Connect(context.Background(), option)
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.Background(), nil) //ping通才代表连接成功
	if err != nil {
		panic(err)
	}
	fmt.Println("连接集群成功", "db", db)
	m.collection = client.Database(db).Collection(collection)
}

func (m *mongodb) updateService(ctx context.Context, svc, src, dst string) error {
	documents, _ := m.collection.CountDocuments(ctx, bson.M{"template_name": svc})
	fmt.Printf("%s 总条数: %d\n", svc, documents)
	find, _ := m.collection.Find(ctx, bson.M{"template_name": svc})

	for find.Next(ctx) {
		var (
			b                  business
			dstDetailConfValue DetailConfValue
			d                  []DetailConfValue
			mid                string
		)
		var bo bool = true
		if err := find.Decode(&b); err != nil {
			fmt.Println("decode失败", err)
			return err
		}
		if err := json.Unmarshal([]byte(b.ConfValue), &d); err != nil {
			fmt.Println("业务配置反序列化失败,", err)
			return err
		}
		for _, value := range d {
			if value.EnvFlag == src {
				dstDetailConfValue.EnvFlag = dst
				mid = value.Value
			}
			if value.EnvFlag == dst {
				bo = false
			}
		}
		if bo {
			dstDetailConfValue.EnvFlag = dst
			dstDetailConfValue.Value = mid
			d = append(d, dstDetailConfValue)
		} else {
			fmt.Printf("更新%s的业务配置%s已经存在,跳过更新\n", svc, b.ConfName)
			continue
		}
		marshal, err := json.Marshal(d)
		if err != nil {
			fmt.Println("序列化失败,", err)
			return err
		}
		//fmt.Println(string(marshal))
		// 更新
		_, err = m.collection.UpdateOne(context.Background(), bson.D{
			{"template_name", svc},
		}, bson.D{
			{"$set", bson.D{
				{"conf_value", string(marshal)},
			}},
		})
		if err != nil {
			fmt.Printf("更新%s的业务配置%s失败\n", svc, b.ConfName)
		}
		fmt.Printf("更新%s的业务配置%s成功\n", svc, b.ConfName)
		break
	}
	return nil
}
