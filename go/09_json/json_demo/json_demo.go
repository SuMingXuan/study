package json_demo
import (
	"fmt"
	"encoding/json"
)
// `` 里面的内容是tag， 序列化时替换掉原有的key
type Server struct {
	ServerName string `json:"name"`
	ServerIP   string `json:"ip"`
	ServerPort int    `json:"port"`
}

func SerializeStruct()  {
	server := new(Server)
	server.ServerIP   = "127.0.0.1"
	server.ServerName = "dev"
	server.ServerPort = 3000
	b, err := json.Marshal(server)
	if err != nil {
		fmt.Println("err: ", err.Error())
	}
	fmt.Println("b ", string(b))
}

func SerializeMap()  {
	server := make(map[string]interface{})
	server["ServerIP"] = "127.0.0.1"
	server["ServerName"] = "dev"
	server["ServerPort"] = 3000
	b, err := json.Marshal(server)
	if err != nil {
		fmt.Println("err: ", err.Error())
	}
	fmt.Println("b ", string(b))
}

func DeSerializeStruct()  {
	jsonString := `{"name":"dev","ip":"127.0.0.1","port":3000}`
	server 		 := new(Server)
	err      	 := json.Unmarshal([]byte(jsonString), &server)
	if err != nil {
		fmt.Println("err ", err.Error())
	}
	fmt.Println("server ", server)

}

func DeSerializeMap()  {
	jsonString := `{"ServerName":"dev","ServerIP":"127.0.0.1","ServerPort":3000}`
	server     := make(map[string]interface{})
	err      	 := json.Unmarshal([]byte(jsonString), &server)
	if err != nil {
		fmt.Println("err: ", err.Error())
	}
	fmt.Println("server ", server)
}