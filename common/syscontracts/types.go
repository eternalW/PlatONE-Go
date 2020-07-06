package syscontracts

import "github.com/PlatONEnetwork/PlatONE-Go/common"

// the system contract addr  table
var (
	UserManagementAddress      = common.HexToAddress("0x1000000000000000000000000000000000000001") // The PlatONE Precompiled contract addr for user management
	NodeManagementAddress      = common.HexToAddress("0x1000000000000000000000000000000000000002") // The PlatONE Precompiled contract addr for node management
	CnsManagementAddress       = common.HexToAddress("0x0000000000000000000000000000000000000011") // The PlatONE Precompiled contract addr for CNS
	ParameterManagementAddress = common.HexToAddress("0x1000000000000000000000000000000000000004") // The PlatONE Precompiled contract addr for parameter management
	FirewallManagementAddress  = common.HexToAddress("0x1000000000000000000000000000000000000005") // The PlatONE Precompiled contract addr for fire wall management
	GroupManagementAddress     = common.HexToAddress("0x1000000000000000000000000000000000000006") // The PlatONE Precompiled contract addr for group management
	ContractDataProcessorAddress     = common.HexToAddress("0x1000000000000000000000000000000000000007") // The PlatONE Precompiled contract addr for group management
	CnsInvokeAddress           = common.HexToAddress("0x0000000000000000000000000000000000000000") // The PlatONE Precompiled contract addr for group management

)

type UpdateNode struct {
	Desc *string `json:"desc,omitempty"`
	Typ  *int32  `json:"type,omitempty"` // 0:观察者节点；1:共识节点
	// status 1为正常节点, 2为删除节点
	Status *int32 `json:"status,omitempty,required"`
	// delay set validatorSet
	DelayNum *uint64 `json:"delayNum,omitempty"` //共识节点延迟设置的区块高度 (可选, 默认实时设置)
}

func (un *UpdateNode) SetStatus(status int32) {
	un.Status = &status
}

func (un *UpdateNode) SetTyp(typ int32) {
	un.Typ = &typ
}

type NodeInfo struct {
	Name  string `json:"name,omitempty,required"` //全网唯一，不能重复。所有接口均以此为主键。 这个名称意义是？
	Owner string `json:"owner,omitempty"`
	Desc  string `json:"desc,omitempty"`
	Typ   int32  `json:"type,omitempty"` // 0:观察者节点；1:共识节点
	// status 1为正常节点, 2为删除节点
	Status     int32  `json:"status,omitempty,required"`
	ExternalIP string `json:"externalIP,omitempty"`
	InternalIP string `json:"internalIP,omitempty,required"`
	PublicKey  string `json:"publicKey,omitempty,required"` //节点公钥，全网唯一，不能重复
	RpcPort    int32  `json:"rpcPort,omitempty"`
	P2pPort    int32  `json:"p2pPort,omitempty,required"`
	// delay set validatorSet
	DelayNum uint64 `json:"delayNum,omitempty"` //共识节点延迟设置的区块高度 (可选, 默认实时设置)
}
