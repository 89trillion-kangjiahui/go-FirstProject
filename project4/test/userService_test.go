package test

import (
	"fmt"
	"testing"

	"project3/response"

	"github.com/golang/protobuf/proto"
)

//测试proto编解码是否正确
func Test_protoDecode(t *testing.T) {
	// 这里定义一个临时的结构体来存储测试case的参数以及期望的返回值
	changes := make(map[uint32]uint64)
	balance := make(map[uint32]uint64)
	changes[1] = 3
	changes[2] = 5
	balance[1] = 3
	balance[2] = 5
	response1 := &response.GeneralReward{
		Code:    1,
		Msg:     "加油",
		Changes: changes,
		Balance: balance,
	}
	data, _ := proto.Marshal(response1)

	ret := response.GeneralReward{}
	proto.Unmarshal(data, &ret)
	if response1.Code == ret.Code && response1.Msg == ret.Msg {
		fmt.Println(ret)
		fmt.Println("proto编解码成功")
	} else {
		fmt.Println("proto编解码失败")
	}
}
