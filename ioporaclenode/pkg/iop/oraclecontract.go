// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iop

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// OracleContractEnrollNode is an auto generated low-level Go binding around an user-defined struct.
type OracleContractEnrollNode struct {
	Addr  common.Address
	Index *big.Int
}

// OracleContractMetaData contains all meta data concerning the OracleContract contract.
var OracleContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumOracleContract.ValidationType\",\"name\":\"typ\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"ValidationRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumOracleContract.ValidationType\",\"name\":\"typ\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"ValidationResponse\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"enrollOracleNode\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"BASE_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CHANLLENGE_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"COMPENSATION_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EnrollOracleNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOTAL_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chanllenge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"countEnrollNodes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"findBlockValidationResult\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"findEnrollNodeByIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structOracleContract.EnrollNode\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"findTransactionValidationResult\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fine\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"oracleNodeIsEnroll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_result\",\"type\":\"bool\"},{\"internalType\":\"uint256[2]\",\"name\":\"_signature\",\"type\":\"uint256[2]\"}],\"name\":\"submitBlockValidationResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_result\",\"type\":\"bool\"},{\"internalType\":\"uint256[2]\",\"name\":\"_signature\",\"type\":\"uint256[2]\"}],\"name\":\"submitTransactionValidationResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"validateBlock\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"validateTransaction\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061278d806100206000396000f3fe6080604052600436106101025760003560e01c806363db7eae11610095578063ab3be84011610064578063ab3be840146102fd578063cf21c29314610328578063dd5e22df14610332578063e92c04d11461034e578063fc769ee91461038b57610109565b806363db7eae1461023f5780637da83e2b1461026a5780637e985fd014610295578063a2ff2ad7146102d257610109565b80633d27ef97116100d15780633d27ef97146101925780633dd14279146101cf57806343434590146101f8578063460e99f41461023557610109565b806312f1e8f41461010b57806325f0854914610122578063344829c81461014b5780633d18651e1461016757610109565b3661010957005b005b34801561011757600080fd5b506101206103b6565b005b34801561012e57600080fd5b50610149600480360381019061014491906118bb565b6105f8565b005b6101656004803603810190610160919061190e565b61060a565b005b34801561017357600080fd5b5061017c610756565b6040516101899190611954565b60405180910390f35b34801561019e57600080fd5b506101b960048036038101906101b4919061190e565b610761565b6040516101c6919061197e565b60405180910390f35b3480156101db57600080fd5b506101f660048036038101906101f191906118bb565b61078b565b005b34801561020457600080fd5b5061021f600480360381019061021a919061190e565b61079d565b60405161022c919061197e565b60405180910390f35b61023d6107c7565b005b34801561024b57600080fd5b506102546108cf565b6040516102619190611954565b60405180910390f35b34801561027657600080fd5b5061027f6108f7565b60405161028c9190611954565b60405180910390f35b3480156102a157600080fd5b506102bc60048036038101906102b791906119c5565b610901565b6040516102c99190611a71565b60405180910390f35b3480156102de57600080fd5b506102e7610a4c565b6040516102f49190611954565b60405180910390f35b34801561030957600080fd5b50610312610a56565b60405161031f9190611954565b60405180910390f35b610330610a61565b005b61034c6004803603810190610347919061190e565b610bc4565b005b34801561035a57600080fd5b5061037560048036038101906103709190611ab8565b610d10565b604051610382919061197e565b60405180910390f35b34801561039757600080fd5b506103a0610dde565b6040516103ad9190611954565b60405180910390f35b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663140f3daa336040518263ffffffff1660e01b81526004016104119190611af4565b602060405180830381865afa15801561042e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104529190611b24565b610491576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161048890611bae565b60405180910390fd5b61049a33610d10565b156104da576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104d190611c1a565b60405180910390fd5b6000600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050338160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600580549050816001018190555060058160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6106056002848484610deb565b505050565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663836f187a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610677573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061069b9190611c4f565b655af3107a40006106ac9190611cab565b66038d7ea4c680006106be9190611ced565b80341015610701576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106f890611d6d565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff167f3706933cbfd265e74e347f4c40263753cecc292080a1bfd0e9fd6ce994c0839660028460405161074a929190611e13565b60405180910390a25050565b66038d7ea4c6800081565b60006002600083815260200190815260200160002060009054906101000a900460ff169050919050565b6107986001848484610deb565b505050565b60006003600083815260200190815260200160002060009054906101000a900460ff169050919050565b6000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633ad59dbc6040518163ffffffff1660e01b8152600401600060405180830381865afa158015610836573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061085f9190612121565b9050655af3107a40008160600151116108ad576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108a4906121b6565b60405180910390fd5b655af3107a400081606001516108c391906121d6565b81606001818152505050565b6064655af3107a40006108e29190611cab565b66038d7ea4c680006108f49190611ced565b81565b655af3107a400081565b61090961179e565b6000821015801561091e575060058054905082105b61095d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161095490612256565b60405180910390fd5b600460006005848154811061097557610974612276565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060400160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820154815250509050919050565b655af3107a400081565b66038d7ea4c6800081565b6000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663655a102f336040518263ffffffff1660e01b8152600401610abe9190611af4565b600060405180830381865afa158015610adb573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190610b049190612121565b905066038d7ea4c68000816060015111610b53576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b4a906121b6565b60405180910390fd5b66038d7ea4c680008160600151610b6a91906121d6565b8160600181815250503373ffffffffffffffffffffffffffffffffffffffff166108fc66038d7ea4c680009081150290604051600060405180830381858888f19350505050158015610bc0573d6000803e3d6000fd5b5050565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663836f187a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c31573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c559190611c4f565b655af3107a4000610c669190611cab565b66038d7ea4c68000610c789190611ced565b80341015610cbb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cb290611d6d565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff167f3706933cbfd265e74e347f4c40263753cecc292080a1bfd0e9fd6ce994c08396600184604051610d04929190611e13565b60405180910390a25050565b60008060058054905003610d275760009050610dd9565b8173ffffffffffffffffffffffffffffffffffffffff166005600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001015481548110610d9457610d93612276565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161490505b919050565b6000600580549050905090565b60006002811115610dff57610dfe611d8d565b5b846002811115610e1257610e11611d8d565b5b03610e52576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e49906122f1565b60405180910390fd5b6000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633ad59dbc6040518163ffffffff1660e01b8152600401600060405180830381865afa158015610ec1573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190610eea9190612121565b90503373ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff1614610f5e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f559061235d565b60405180910390fd5b6000610f8c858588604051602001610f789392919061237d565b604051602081830303815290604052611444565b90506000600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632e3344526040518163ffffffff1660e01b8152600401608060405180830381865afa158015610ffd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110219190612465565b905060006040518061018001604052808460006002811061104557611044612276565b5b602002015181526020018460016002811061106357611062612276565b5b602002015181526020018360006004811061108157611080612276565b5b602002015181526020018360016004811061109f5761109e612276565b5b60200201518152602001836002600481106110bd576110bc612276565b5b60200201518152602001836003600481106110db576110da612276565b5b60200201518152602001866000600281106110f9576110f8612276565b5b602002013581526020018660016002811061111757611116612276565b5b602002013581526020017f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c281526020017f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed81526020017f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec81526020017f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d81525090506111c2816114ab565b611201576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111f8906124de565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff166108fc66038d7ea4c680009081150290604051600060405180830381858888f1935050505015801561124e573d6000803e3d6000fd5b5060005b6005805490508163ffffffff1610156113095760058163ffffffff168154811061127f5761127e612276565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc655af3107a40009081150290604051600060405180830381858888f193505050501580156112f5573d6000803e3d6000fd5b5080806113019061250e565b915050611252565b506001600281111561131e5761131d611d8d565b5b88600281111561133157611330611d8d565b5b0361136657856002600089815260200190815260200160002060006101000a81548160ff0219169083151502179055506113be565b60028081111561137957611378611d8d565b5b88600281111561138c5761138b611d8d565b5b036113bd57856003600089815260200190815260200160002060006101000a81548160ff0219169083151502179055505b5b3373ffffffffffffffffffffffffffffffffffffffff167f9739d27192db56a131c86c297677afccac89aa4f39cbc5bcd62a8d10ce559675898989600580549050655af3107a40006114109190611cab565b66038d7ea4c680006114229190611ced565b604051611432949392919061253a565b60405180910390a25050505050505050565b61144c6117ce565b6114a460028360405161145f91906125c6565b602060405180830381855afa15801561147c573d6000803e3d6000fd5b5050506040513d601f19601f8201168201806040525081019061149f91906125f2565b61152f565b9050919050565b60006114b56117f0565b600060208261018086600060086107d05a03f190508061150a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115019061266b565b60405180910390fd5b6001826000600181106115205761151f612276565b5b60200201511492505050919050565b6115376117ce565b60007f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478360001c61156891906126ba565b9050600080600090505b6001156116a8577f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47806115a8576115a761268b565b5b83840991507f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47806115dc576115db61268b565b5b83830991507f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47806116105761160f61268b565b5b60038308915061161f826116b0565b8092508193505050801561166e57828460006002811061164257611641612276565b5b60200201818152505081846001600281106116605761165f612276565b5b6020020181815250506116a8565b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478061169d5761169c61268b565b5b600184089250611572565b505050919050565b600080600060405160208152602080820152602060408201528460608201527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808201527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4760a082015260208160c08360056107d05a03fa9150805193507f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47848509851492505080611798576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161178f90612737565b60405180910390fd5b50915091565b6040518060400160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b6040518060400160405280600290602082028036833780820191505090505090565b6040518060200160405280600190602082028036833780820191505090505090565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b61183981611826565b811461184457600080fd5b50565b60008135905061185681611830565b92915050565b60008115159050919050565b6118718161185c565b811461187c57600080fd5b50565b60008135905061188e81611868565b92915050565b600080fd5b6000819050826020600202820111156118b5576118b4611894565b5b92915050565b6000806000608084860312156118d4576118d361181c565b5b60006118e286828701611847565b93505060206118f38682870161187f565b925050604061190486828701611899565b9150509250925092565b6000602082840312156119245761192361181c565b5b600061193284828501611847565b91505092915050565b6000819050919050565b61194e8161193b565b82525050565b60006020820190506119696000830184611945565b92915050565b6119788161185c565b82525050565b6000602082019050611993600083018461196f565b92915050565b6119a28161193b565b81146119ad57600080fd5b50565b6000813590506119bf81611999565b92915050565b6000602082840312156119db576119da61181c565b5b60006119e9848285016119b0565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000611a1d826119f2565b9050919050565b611a2d81611a12565b82525050565b611a3c8161193b565b82525050565b604082016000820151611a586000850182611a24565b506020820151611a6b6020850182611a33565b50505050565b6000604082019050611a866000830184611a42565b92915050565b611a9581611a12565b8114611aa057600080fd5b50565b600081359050611ab281611a8c565b92915050565b600060208284031215611ace57611acd61181c565b5b6000611adc84828501611aa3565b91505092915050565b611aee81611a12565b82525050565b6000602082019050611b096000830184611ae5565b92915050565b600081519050611b1e81611868565b92915050565b600060208284031215611b3a57611b3961181c565b5b6000611b4884828501611b0f565b91505092915050565b600082825260208201905092915050565b7f546865204f7261636c6520646f65736e27742072656769737465726564000000600082015250565b6000611b98601d83611b51565b9150611ba382611b62565b602082019050919050565b60006020820190508181036000830152611bc781611b8b565b9050919050565b7f616c726561647920656e726f6c6c656400000000000000000000000000000000600082015250565b6000611c04601083611b51565b9150611c0f82611bce565b602082019050919050565b60006020820190508181036000830152611c3381611bf7565b9050919050565b600081519050611c4981611999565b92915050565b600060208284031215611c6557611c6461181c565b5b6000611c7384828501611c3a565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611cb68261193b565b9150611cc18361193b565b9250828202611ccf8161193b565b91508282048414831517611ce657611ce5611c7c565b5b5092915050565b6000611cf88261193b565b9150611d038361193b565b9250828201905080821115611d1b57611d1a611c7c565b5b92915050565b7f746f6f206665772066656520616d6f756e740000000000000000000000000000600082015250565b6000611d57601283611b51565b9150611d6282611d21565b602082019050919050565b60006020820190508181036000830152611d8681611d4a565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60038110611dcd57611dcc611d8d565b5b50565b6000819050611dde82611dbc565b919050565b6000611dee82611dd0565b9050919050565b611dfe81611de3565b82525050565b611e0d81611826565b82525050565b6000604082019050611e286000830185611df5565b611e356020830184611e04565b9392505050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b611e8a82611e41565b810181811067ffffffffffffffff82111715611ea957611ea8611e52565b5b80604052505050565b6000611ebc611812565b9050611ec88282611e81565b919050565b600080fd5b600081519050611ee181611a8c565b92915050565b600080fd5b600080fd5b600067ffffffffffffffff821115611f0c57611f0b611e52565b5b611f1582611e41565b9050602081019050919050565b60005b83811015611f40578082015181840152602081019050611f25565b60008484015250505050565b6000611f5f611f5a84611ef1565b611eb2565b905082815260208101848484011115611f7b57611f7a611eec565b5b611f86848285611f22565b509392505050565b600082601f830112611fa357611fa2611ee7565b5b8151611fb3848260208601611f4c565b91505092915050565b600067ffffffffffffffff821115611fd757611fd6611e52565b5b611fe082611e41565b9050602081019050919050565b6000612000611ffb84611fbc565b611eb2565b90508281526020810184848401111561201c5761201b611eec565b5b612027848285611f22565b509392505050565b600082601f83011261204457612043611ee7565b5b8151612054848260208601611fed565b91505092915050565b600060a0828403121561207357612072611e3c565b5b61207d60a0611eb2565b9050600061208d84828501611ed2565b600083015250602082015167ffffffffffffffff8111156120b1576120b0611ecd565b5b6120bd84828501611f8e565b602083015250604082015167ffffffffffffffff8111156120e1576120e0611ecd565b5b6120ed8482850161202f565b604083015250606061210184828501611c3a565b606083015250608061211584828501611c3a565b60808301525092915050565b6000602082840312156121375761213661181c565b5b600082015167ffffffffffffffff81111561215557612154611821565b5b6121618482850161205d565b91505092915050565b7f7374616b6520697320746f6f206c6f7721000000000000000000000000000000600082015250565b60006121a0601183611b51565b91506121ab8261216a565b602082019050919050565b600060208201905081810360008301526121cf81612193565b9050919050565b60006121e18261193b565b91506121ec8361193b565b925082820390508181111561220457612203611c7c565b5b92915050565b7f6e6f7420666f756e640000000000000000000000000000000000000000000000600082015250565b6000612240600983611b51565b915061224b8261220a565b602082019050919050565b6000602082019050818103600083015261226f81612233565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f756e6b6e6f776e2076616c69646174696f6e2074797065000000000000000000600082015250565b60006122db601783611b51565b91506122e6826122a5565b602082019050919050565b6000602082019050818103600083015261230a816122ce565b9050919050565b7f6e6f74207468652061676772656761746f720000000000000000000000000000600082015250565b6000612347601283611b51565b915061235282612311565b602082019050919050565b600060208201905081810360008301526123768161233a565b9050919050565b60006060820190506123926000830186611e04565b61239f602083018561196f565b6123ac6040830184611df5565b949350505050565b600067ffffffffffffffff8211156123cf576123ce611e52565b5b602082029050919050565b60006123ed6123e8846123b4565b611eb2565b9050806020840283018581111561240757612406611894565b5b835b81811015612430578061241c8882611c3a565b845260208401935050602081019050612409565b5050509392505050565b600082601f83011261244f5761244e611ee7565b5b600461245c8482856123da565b91505092915050565b60006080828403121561247b5761247a61181c565b5b60006124898482850161243a565b91505092915050565b7f696e76616c6964207369676e6174757265000000000000000000000000000000600082015250565b60006124c8601183611b51565b91506124d382612492565b602082019050919050565b600060208201905081810360008301526124f7816124bb565b9050919050565b600063ffffffff82169050919050565b6000612519826124fe565b915063ffffffff820361252f5761252e611c7c565b5b600182019050919050565b600060808201905061254f6000830187611df5565b61255c6020830186611e04565b612569604083018561196f565b6125766060830184611945565b95945050505050565b600081519050919050565b600081905092915050565b60006125a08261257f565b6125aa818561258a565b93506125ba818560208601611f22565b80840191505092915050565b60006125d28284612595565b915081905092915050565b6000815190506125ec81611830565b92915050565b6000602082840312156126085761260761181c565b5b6000612616848285016125dd565b91505092915050565b7f65632070616972696e67206661696c6564000000000000000000000000000000600082015250565b6000612655601183611b51565b91506126608261261f565b602082019050919050565b6000602082019050818103600083015261268481612648565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60006126c58261193b565b91506126d08361193b565b9250826126e0576126df61268b565b5b828206905092915050565b7f73717274206d6f646578702063616c6c206661696c6564000000000000000000600082015250565b6000612721601783611b51565b915061272c826126eb565b602082019050919050565b6000602082019050818103600083015261275081612714565b905091905056fea2646970667358221220051edce9e2c0263052aa4c01162a644b4c8c8207fa6d5292399905ae64cdf18a64736f6c63430008130033",
}

// OracleContractABI is the input ABI used to generate the binding from.
// Deprecated: Use OracleContractMetaData.ABI instead.
var OracleContractABI = OracleContractMetaData.ABI

// OracleContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OracleContractMetaData.Bin instead.
var OracleContractBin = OracleContractMetaData.Bin

// DeployOracleContract deploys a new Ethereum contract, binding an instance of OracleContract to it.
func DeployOracleContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OracleContract, error) {
	parsed, err := OracleContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OracleContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OracleContract{OracleContractCaller: OracleContractCaller{contract: contract}, OracleContractTransactor: OracleContractTransactor{contract: contract}, OracleContractFilterer: OracleContractFilterer{contract: contract}}, nil
}

// OracleContract is an auto generated Go binding around an Ethereum contract.
type OracleContract struct {
	OracleContractCaller     // Read-only binding to the contract
	OracleContractTransactor // Write-only binding to the contract
	OracleContractFilterer   // Log filterer for contract events
}

// OracleContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleContractSession struct {
	Contract     *OracleContract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleContractCallerSession struct {
	Contract *OracleContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// OracleContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleContractTransactorSession struct {
	Contract     *OracleContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// OracleContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleContractRaw struct {
	Contract *OracleContract // Generic contract binding to access the raw methods on
}

// OracleContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleContractCallerRaw struct {
	Contract *OracleContractCaller // Generic read-only contract binding to access the raw methods on
}

// OracleContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleContractTransactorRaw struct {
	Contract *OracleContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracleContract creates a new instance of OracleContract, bound to a specific deployed contract.
func NewOracleContract(address common.Address, backend bind.ContractBackend) (*OracleContract, error) {
	contract, err := bindOracleContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OracleContract{OracleContractCaller: OracleContractCaller{contract: contract}, OracleContractTransactor: OracleContractTransactor{contract: contract}, OracleContractFilterer: OracleContractFilterer{contract: contract}}, nil
}

// NewOracleContractCaller creates a new read-only instance of OracleContract, bound to a specific deployed contract.
func NewOracleContractCaller(address common.Address, caller bind.ContractCaller) (*OracleContractCaller, error) {
	contract, err := bindOracleContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleContractCaller{contract: contract}, nil
}

// NewOracleContractTransactor creates a new write-only instance of OracleContract, bound to a specific deployed contract.
func NewOracleContractTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleContractTransactor, error) {
	contract, err := bindOracleContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleContractTransactor{contract: contract}, nil
}

// NewOracleContractFilterer creates a new log filterer instance of OracleContract, bound to a specific deployed contract.
func NewOracleContractFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleContractFilterer, error) {
	contract, err := bindOracleContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleContractFilterer{contract: contract}, nil
}

// bindOracleContract binds a generic wrapper to an already deployed contract.
func bindOracleContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OracleContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleContract *OracleContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OracleContract.Contract.OracleContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleContract *OracleContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleContract.Contract.OracleContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleContract *OracleContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleContract.Contract.OracleContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleContract *OracleContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OracleContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleContract *OracleContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleContract *OracleContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleContract.Contract.contract.Transact(opts, method, params...)
}

// BASEFEE is a free data retrieval call binding the contract method 0x3d18651e.
//
// Solidity: function BASE_FEE() view returns(uint256)
func (_OracleContract *OracleContractCaller) BASEFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "BASE_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASEFEE is a free data retrieval call binding the contract method 0x3d18651e.
//
// Solidity: function BASE_FEE() view returns(uint256)
func (_OracleContract *OracleContractSession) BASEFEE() (*big.Int, error) {
	return _OracleContract.Contract.BASEFEE(&_OracleContract.CallOpts)
}

// BASEFEE is a free data retrieval call binding the contract method 0x3d18651e.
//
// Solidity: function BASE_FEE() view returns(uint256)
func (_OracleContract *OracleContractCallerSession) BASEFEE() (*big.Int, error) {
	return _OracleContract.Contract.BASEFEE(&_OracleContract.CallOpts)
}

// CHANLLENGEFEE is a free data retrieval call binding the contract method 0xa2ff2ad7.
//
// Solidity: function CHANLLENGE_FEE() view returns(uint256)
func (_OracleContract *OracleContractCaller) CHANLLENGEFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "CHANLLENGE_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CHANLLENGEFEE is a free data retrieval call binding the contract method 0xa2ff2ad7.
//
// Solidity: function CHANLLENGE_FEE() view returns(uint256)
func (_OracleContract *OracleContractSession) CHANLLENGEFEE() (*big.Int, error) {
	return _OracleContract.Contract.CHANLLENGEFEE(&_OracleContract.CallOpts)
}

// CHANLLENGEFEE is a free data retrieval call binding the contract method 0xa2ff2ad7.
//
// Solidity: function CHANLLENGE_FEE() view returns(uint256)
func (_OracleContract *OracleContractCallerSession) CHANLLENGEFEE() (*big.Int, error) {
	return _OracleContract.Contract.CHANLLENGEFEE(&_OracleContract.CallOpts)
}

// COMPENSATIONFEE is a free data retrieval call binding the contract method 0xab3be840.
//
// Solidity: function COMPENSATION_FEE() view returns(uint256)
func (_OracleContract *OracleContractCaller) COMPENSATIONFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "COMPENSATION_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// COMPENSATIONFEE is a free data retrieval call binding the contract method 0xab3be840.
//
// Solidity: function COMPENSATION_FEE() view returns(uint256)
func (_OracleContract *OracleContractSession) COMPENSATIONFEE() (*big.Int, error) {
	return _OracleContract.Contract.COMPENSATIONFEE(&_OracleContract.CallOpts)
}

// COMPENSATIONFEE is a free data retrieval call binding the contract method 0xab3be840.
//
// Solidity: function COMPENSATION_FEE() view returns(uint256)
func (_OracleContract *OracleContractCallerSession) COMPENSATIONFEE() (*big.Int, error) {
	return _OracleContract.Contract.COMPENSATIONFEE(&_OracleContract.CallOpts)
}

// TOTALFEE is a free data retrieval call binding the contract method 0x63db7eae.
//
// Solidity: function TOTAL_FEE() view returns(uint256)
func (_OracleContract *OracleContractCaller) TOTALFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "TOTAL_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TOTALFEE is a free data retrieval call binding the contract method 0x63db7eae.
//
// Solidity: function TOTAL_FEE() view returns(uint256)
func (_OracleContract *OracleContractSession) TOTALFEE() (*big.Int, error) {
	return _OracleContract.Contract.TOTALFEE(&_OracleContract.CallOpts)
}

// TOTALFEE is a free data retrieval call binding the contract method 0x63db7eae.
//
// Solidity: function TOTAL_FEE() view returns(uint256)
func (_OracleContract *OracleContractCallerSession) TOTALFEE() (*big.Int, error) {
	return _OracleContract.Contract.TOTALFEE(&_OracleContract.CallOpts)
}

// VALIDATORFEE is a free data retrieval call binding the contract method 0x7da83e2b.
//
// Solidity: function VALIDATOR_FEE() view returns(uint256)
func (_OracleContract *OracleContractCaller) VALIDATORFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "VALIDATOR_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VALIDATORFEE is a free data retrieval call binding the contract method 0x7da83e2b.
//
// Solidity: function VALIDATOR_FEE() view returns(uint256)
func (_OracleContract *OracleContractSession) VALIDATORFEE() (*big.Int, error) {
	return _OracleContract.Contract.VALIDATORFEE(&_OracleContract.CallOpts)
}

// VALIDATORFEE is a free data retrieval call binding the contract method 0x7da83e2b.
//
// Solidity: function VALIDATOR_FEE() view returns(uint256)
func (_OracleContract *OracleContractCallerSession) VALIDATORFEE() (*big.Int, error) {
	return _OracleContract.Contract.VALIDATORFEE(&_OracleContract.CallOpts)
}

// CountEnrollNodes is a free data retrieval call binding the contract method 0xfc769ee9.
//
// Solidity: function countEnrollNodes() view returns(uint256)
func (_OracleContract *OracleContractCaller) CountEnrollNodes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "countEnrollNodes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CountEnrollNodes is a free data retrieval call binding the contract method 0xfc769ee9.
//
// Solidity: function countEnrollNodes() view returns(uint256)
func (_OracleContract *OracleContractSession) CountEnrollNodes() (*big.Int, error) {
	return _OracleContract.Contract.CountEnrollNodes(&_OracleContract.CallOpts)
}

// CountEnrollNodes is a free data retrieval call binding the contract method 0xfc769ee9.
//
// Solidity: function countEnrollNodes() view returns(uint256)
func (_OracleContract *OracleContractCallerSession) CountEnrollNodes() (*big.Int, error) {
	return _OracleContract.Contract.CountEnrollNodes(&_OracleContract.CallOpts)
}

// FindBlockValidationResult is a free data retrieval call binding the contract method 0x3d27ef97.
//
// Solidity: function findBlockValidationResult(bytes32 _hash) view returns(bool)
func (_OracleContract *OracleContractCaller) FindBlockValidationResult(opts *bind.CallOpts, _hash [32]byte) (bool, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "findBlockValidationResult", _hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FindBlockValidationResult is a free data retrieval call binding the contract method 0x3d27ef97.
//
// Solidity: function findBlockValidationResult(bytes32 _hash) view returns(bool)
func (_OracleContract *OracleContractSession) FindBlockValidationResult(_hash [32]byte) (bool, error) {
	return _OracleContract.Contract.FindBlockValidationResult(&_OracleContract.CallOpts, _hash)
}

// FindBlockValidationResult is a free data retrieval call binding the contract method 0x3d27ef97.
//
// Solidity: function findBlockValidationResult(bytes32 _hash) view returns(bool)
func (_OracleContract *OracleContractCallerSession) FindBlockValidationResult(_hash [32]byte) (bool, error) {
	return _OracleContract.Contract.FindBlockValidationResult(&_OracleContract.CallOpts, _hash)
}

// FindEnrollNodeByIndex is a free data retrieval call binding the contract method 0x7e985fd0.
//
// Solidity: function findEnrollNodeByIndex(uint256 _index) view returns((address,uint256))
func (_OracleContract *OracleContractCaller) FindEnrollNodeByIndex(opts *bind.CallOpts, _index *big.Int) (OracleContractEnrollNode, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "findEnrollNodeByIndex", _index)

	if err != nil {
		return *new(OracleContractEnrollNode), err
	}

	out0 := *abi.ConvertType(out[0], new(OracleContractEnrollNode)).(*OracleContractEnrollNode)

	return out0, err

}

// FindEnrollNodeByIndex is a free data retrieval call binding the contract method 0x7e985fd0.
//
// Solidity: function findEnrollNodeByIndex(uint256 _index) view returns((address,uint256))
func (_OracleContract *OracleContractSession) FindEnrollNodeByIndex(_index *big.Int) (OracleContractEnrollNode, error) {
	return _OracleContract.Contract.FindEnrollNodeByIndex(&_OracleContract.CallOpts, _index)
}

// FindEnrollNodeByIndex is a free data retrieval call binding the contract method 0x7e985fd0.
//
// Solidity: function findEnrollNodeByIndex(uint256 _index) view returns((address,uint256))
func (_OracleContract *OracleContractCallerSession) FindEnrollNodeByIndex(_index *big.Int) (OracleContractEnrollNode, error) {
	return _OracleContract.Contract.FindEnrollNodeByIndex(&_OracleContract.CallOpts, _index)
}

// FindTransactionValidationResult is a free data retrieval call binding the contract method 0x43434590.
//
// Solidity: function findTransactionValidationResult(bytes32 _hash) view returns(bool)
func (_OracleContract *OracleContractCaller) FindTransactionValidationResult(opts *bind.CallOpts, _hash [32]byte) (bool, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "findTransactionValidationResult", _hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FindTransactionValidationResult is a free data retrieval call binding the contract method 0x43434590.
//
// Solidity: function findTransactionValidationResult(bytes32 _hash) view returns(bool)
func (_OracleContract *OracleContractSession) FindTransactionValidationResult(_hash [32]byte) (bool, error) {
	return _OracleContract.Contract.FindTransactionValidationResult(&_OracleContract.CallOpts, _hash)
}

// FindTransactionValidationResult is a free data retrieval call binding the contract method 0x43434590.
//
// Solidity: function findTransactionValidationResult(bytes32 _hash) view returns(bool)
func (_OracleContract *OracleContractCallerSession) FindTransactionValidationResult(_hash [32]byte) (bool, error) {
	return _OracleContract.Contract.FindTransactionValidationResult(&_OracleContract.CallOpts, _hash)
}

// OracleNodeIsEnroll is a free data retrieval call binding the contract method 0xe92c04d1.
//
// Solidity: function oracleNodeIsEnroll(address _addr) view returns(bool)
func (_OracleContract *OracleContractCaller) OracleNodeIsEnroll(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "oracleNodeIsEnroll", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OracleNodeIsEnroll is a free data retrieval call binding the contract method 0xe92c04d1.
//
// Solidity: function oracleNodeIsEnroll(address _addr) view returns(bool)
func (_OracleContract *OracleContractSession) OracleNodeIsEnroll(_addr common.Address) (bool, error) {
	return _OracleContract.Contract.OracleNodeIsEnroll(&_OracleContract.CallOpts, _addr)
}

// OracleNodeIsEnroll is a free data retrieval call binding the contract method 0xe92c04d1.
//
// Solidity: function oracleNodeIsEnroll(address _addr) view returns(bool)
func (_OracleContract *OracleContractCallerSession) OracleNodeIsEnroll(_addr common.Address) (bool, error) {
	return _OracleContract.Contract.OracleNodeIsEnroll(&_OracleContract.CallOpts, _addr)
}

// EnrollOracleNode is a paid mutator transaction binding the contract method 0x12f1e8f4.
//
// Solidity: function EnrollOracleNode() returns()
func (_OracleContract *OracleContractTransactor) EnrollOracleNode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "EnrollOracleNode")
}

// EnrollOracleNode is a paid mutator transaction binding the contract method 0x12f1e8f4.
//
// Solidity: function EnrollOracleNode() returns()
func (_OracleContract *OracleContractSession) EnrollOracleNode() (*types.Transaction, error) {
	return _OracleContract.Contract.EnrollOracleNode(&_OracleContract.TransactOpts)
}

// EnrollOracleNode is a paid mutator transaction binding the contract method 0x12f1e8f4.
//
// Solidity: function EnrollOracleNode() returns()
func (_OracleContract *OracleContractTransactorSession) EnrollOracleNode() (*types.Transaction, error) {
	return _OracleContract.Contract.EnrollOracleNode(&_OracleContract.TransactOpts)
}

// Chanllenge is a paid mutator transaction binding the contract method 0x460e99f4.
//
// Solidity: function chanllenge() payable returns()
func (_OracleContract *OracleContractTransactor) Chanllenge(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "chanllenge")
}

// Chanllenge is a paid mutator transaction binding the contract method 0x460e99f4.
//
// Solidity: function chanllenge() payable returns()
func (_OracleContract *OracleContractSession) Chanllenge() (*types.Transaction, error) {
	return _OracleContract.Contract.Chanllenge(&_OracleContract.TransactOpts)
}

// Chanllenge is a paid mutator transaction binding the contract method 0x460e99f4.
//
// Solidity: function chanllenge() payable returns()
func (_OracleContract *OracleContractTransactorSession) Chanllenge() (*types.Transaction, error) {
	return _OracleContract.Contract.Chanllenge(&_OracleContract.TransactOpts)
}

// Fine is a paid mutator transaction binding the contract method 0xcf21c293.
//
// Solidity: function fine() payable returns()
func (_OracleContract *OracleContractTransactor) Fine(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "fine")
}

// Fine is a paid mutator transaction binding the contract method 0xcf21c293.
//
// Solidity: function fine() payable returns()
func (_OracleContract *OracleContractSession) Fine() (*types.Transaction, error) {
	return _OracleContract.Contract.Fine(&_OracleContract.TransactOpts)
}

// Fine is a paid mutator transaction binding the contract method 0xcf21c293.
//
// Solidity: function fine() payable returns()
func (_OracleContract *OracleContractTransactorSession) Fine() (*types.Transaction, error) {
	return _OracleContract.Contract.Fine(&_OracleContract.TransactOpts)
}

// SubmitBlockValidationResult is a paid mutator transaction binding the contract method 0x3dd14279.
//
// Solidity: function submitBlockValidationResult(bytes32 _hash, bool _result, uint256[2] _signature) returns()
func (_OracleContract *OracleContractTransactor) SubmitBlockValidationResult(opts *bind.TransactOpts, _hash [32]byte, _result bool, _signature [2]*big.Int) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "submitBlockValidationResult", _hash, _result, _signature)
}

// SubmitBlockValidationResult is a paid mutator transaction binding the contract method 0x3dd14279.
//
// Solidity: function submitBlockValidationResult(bytes32 _hash, bool _result, uint256[2] _signature) returns()
func (_OracleContract *OracleContractSession) SubmitBlockValidationResult(_hash [32]byte, _result bool, _signature [2]*big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.SubmitBlockValidationResult(&_OracleContract.TransactOpts, _hash, _result, _signature)
}

// SubmitBlockValidationResult is a paid mutator transaction binding the contract method 0x3dd14279.
//
// Solidity: function submitBlockValidationResult(bytes32 _hash, bool _result, uint256[2] _signature) returns()
func (_OracleContract *OracleContractTransactorSession) SubmitBlockValidationResult(_hash [32]byte, _result bool, _signature [2]*big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.SubmitBlockValidationResult(&_OracleContract.TransactOpts, _hash, _result, _signature)
}

// SubmitTransactionValidationResult is a paid mutator transaction binding the contract method 0x25f08549.
//
// Solidity: function submitTransactionValidationResult(bytes32 _hash, bool _result, uint256[2] _signature) returns()
func (_OracleContract *OracleContractTransactor) SubmitTransactionValidationResult(opts *bind.TransactOpts, _hash [32]byte, _result bool, _signature [2]*big.Int) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "submitTransactionValidationResult", _hash, _result, _signature)
}

// SubmitTransactionValidationResult is a paid mutator transaction binding the contract method 0x25f08549.
//
// Solidity: function submitTransactionValidationResult(bytes32 _hash, bool _result, uint256[2] _signature) returns()
func (_OracleContract *OracleContractSession) SubmitTransactionValidationResult(_hash [32]byte, _result bool, _signature [2]*big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.SubmitTransactionValidationResult(&_OracleContract.TransactOpts, _hash, _result, _signature)
}

// SubmitTransactionValidationResult is a paid mutator transaction binding the contract method 0x25f08549.
//
// Solidity: function submitTransactionValidationResult(bytes32 _hash, bool _result, uint256[2] _signature) returns()
func (_OracleContract *OracleContractTransactorSession) SubmitTransactionValidationResult(_hash [32]byte, _result bool, _signature [2]*big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.SubmitTransactionValidationResult(&_OracleContract.TransactOpts, _hash, _result, _signature)
}

// ValidateBlock is a paid mutator transaction binding the contract method 0xdd5e22df.
//
// Solidity: function validateBlock(bytes32 _hash) payable returns()
func (_OracleContract *OracleContractTransactor) ValidateBlock(opts *bind.TransactOpts, _hash [32]byte) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "validateBlock", _hash)
}

// ValidateBlock is a paid mutator transaction binding the contract method 0xdd5e22df.
//
// Solidity: function validateBlock(bytes32 _hash) payable returns()
func (_OracleContract *OracleContractSession) ValidateBlock(_hash [32]byte) (*types.Transaction, error) {
	return _OracleContract.Contract.ValidateBlock(&_OracleContract.TransactOpts, _hash)
}

// ValidateBlock is a paid mutator transaction binding the contract method 0xdd5e22df.
//
// Solidity: function validateBlock(bytes32 _hash) payable returns()
func (_OracleContract *OracleContractTransactorSession) ValidateBlock(_hash [32]byte) (*types.Transaction, error) {
	return _OracleContract.Contract.ValidateBlock(&_OracleContract.TransactOpts, _hash)
}

// ValidateTransaction is a paid mutator transaction binding the contract method 0x344829c8.
//
// Solidity: function validateTransaction(bytes32 _hash) payable returns()
func (_OracleContract *OracleContractTransactor) ValidateTransaction(opts *bind.TransactOpts, _hash [32]byte) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "validateTransaction", _hash)
}

// ValidateTransaction is a paid mutator transaction binding the contract method 0x344829c8.
//
// Solidity: function validateTransaction(bytes32 _hash) payable returns()
func (_OracleContract *OracleContractSession) ValidateTransaction(_hash [32]byte) (*types.Transaction, error) {
	return _OracleContract.Contract.ValidateTransaction(&_OracleContract.TransactOpts, _hash)
}

// ValidateTransaction is a paid mutator transaction binding the contract method 0x344829c8.
//
// Solidity: function validateTransaction(bytes32 _hash) payable returns()
func (_OracleContract *OracleContractTransactorSession) ValidateTransaction(_hash [32]byte) (*types.Transaction, error) {
	return _OracleContract.Contract.ValidateTransaction(&_OracleContract.TransactOpts, _hash)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_OracleContract *OracleContractTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _OracleContract.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_OracleContract *OracleContractSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _OracleContract.Contract.Fallback(&_OracleContract.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_OracleContract *OracleContractTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _OracleContract.Contract.Fallback(&_OracleContract.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OracleContract *OracleContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleContract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OracleContract *OracleContractSession) Receive() (*types.Transaction, error) {
	return _OracleContract.Contract.Receive(&_OracleContract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OracleContract *OracleContractTransactorSession) Receive() (*types.Transaction, error) {
	return _OracleContract.Contract.Receive(&_OracleContract.TransactOpts)
}

// OracleContractValidationRequestIterator is returned from FilterValidationRequest and is used to iterate over the raw logs and unpacked data for ValidationRequest events raised by the OracleContract contract.
type OracleContractValidationRequestIterator struct {
	Event *OracleContractValidationRequest // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OracleContractValidationRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleContractValidationRequest)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OracleContractValidationRequest)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OracleContractValidationRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleContractValidationRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleContractValidationRequest represents a ValidationRequest event raised by the OracleContract contract.
type OracleContractValidationRequest struct {
	Typ  uint8
	From common.Address
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterValidationRequest is a free log retrieval operation binding the contract event 0x3706933cbfd265e74e347f4c40263753cecc292080a1bfd0e9fd6ce994c08396.
//
// Solidity: event ValidationRequest(uint8 typ, address indexed from, bytes32 hash)
func (_OracleContract *OracleContractFilterer) FilterValidationRequest(opts *bind.FilterOpts, from []common.Address) (*OracleContractValidationRequestIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _OracleContract.contract.FilterLogs(opts, "ValidationRequest", fromRule)
	if err != nil {
		return nil, err
	}
	return &OracleContractValidationRequestIterator{contract: _OracleContract.contract, event: "ValidationRequest", logs: logs, sub: sub}, nil
}

// WatchValidationRequest is a free log subscription operation binding the contract event 0x3706933cbfd265e74e347f4c40263753cecc292080a1bfd0e9fd6ce994c08396.
//
// Solidity: event ValidationRequest(uint8 typ, address indexed from, bytes32 hash)
func (_OracleContract *OracleContractFilterer) WatchValidationRequest(opts *bind.WatchOpts, sink chan<- *OracleContractValidationRequest, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _OracleContract.contract.WatchLogs(opts, "ValidationRequest", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleContractValidationRequest)
				if err := _OracleContract.contract.UnpackLog(event, "ValidationRequest", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidationRequest is a log parse operation binding the contract event 0x3706933cbfd265e74e347f4c40263753cecc292080a1bfd0e9fd6ce994c08396.
//
// Solidity: event ValidationRequest(uint8 typ, address indexed from, bytes32 hash)
func (_OracleContract *OracleContractFilterer) ParseValidationRequest(log types.Log) (*OracleContractValidationRequest, error) {
	event := new(OracleContractValidationRequest)
	if err := _OracleContract.contract.UnpackLog(event, "ValidationRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleContractValidationResponseIterator is returned from FilterValidationResponse and is used to iterate over the raw logs and unpacked data for ValidationResponse events raised by the OracleContract contract.
type OracleContractValidationResponseIterator struct {
	Event *OracleContractValidationResponse // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OracleContractValidationResponseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleContractValidationResponse)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OracleContractValidationResponse)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OracleContractValidationResponseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleContractValidationResponseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleContractValidationResponse represents a ValidationResponse event raised by the OracleContract contract.
type OracleContractValidationResponse struct {
	Typ        uint8
	Aggregator common.Address
	Hash       [32]byte
	Valid      bool
	Fee        *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterValidationResponse is a free log retrieval operation binding the contract event 0x9739d27192db56a131c86c297677afccac89aa4f39cbc5bcd62a8d10ce559675.
//
// Solidity: event ValidationResponse(uint8 typ, address indexed aggregator, bytes32 hash, bool valid, uint256 fee)
func (_OracleContract *OracleContractFilterer) FilterValidationResponse(opts *bind.FilterOpts, aggregator []common.Address) (*OracleContractValidationResponseIterator, error) {

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _OracleContract.contract.FilterLogs(opts, "ValidationResponse", aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &OracleContractValidationResponseIterator{contract: _OracleContract.contract, event: "ValidationResponse", logs: logs, sub: sub}, nil
}

// WatchValidationResponse is a free log subscription operation binding the contract event 0x9739d27192db56a131c86c297677afccac89aa4f39cbc5bcd62a8d10ce559675.
//
// Solidity: event ValidationResponse(uint8 typ, address indexed aggregator, bytes32 hash, bool valid, uint256 fee)
func (_OracleContract *OracleContractFilterer) WatchValidationResponse(opts *bind.WatchOpts, sink chan<- *OracleContractValidationResponse, aggregator []common.Address) (event.Subscription, error) {

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _OracleContract.contract.WatchLogs(opts, "ValidationResponse", aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleContractValidationResponse)
				if err := _OracleContract.contract.UnpackLog(event, "ValidationResponse", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidationResponse is a log parse operation binding the contract event 0x9739d27192db56a131c86c297677afccac89aa4f39cbc5bcd62a8d10ce559675.
//
// Solidity: event ValidationResponse(uint8 typ, address indexed aggregator, bytes32 hash, bool valid, uint256 fee)
func (_OracleContract *OracleContractFilterer) ParseValidationResponse(log types.Log) (*OracleContractValidationResponse, error) {
	event := new(OracleContractValidationResponse)
	if err := _OracleContract.contract.UnpackLog(event, "ValidationResponse", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleContractEnrollOracleNodeIterator is returned from FilterEnrollOracleNode and is used to iterate over the raw logs and unpacked data for EnrollOracleNode events raised by the OracleContract contract.
type OracleContractEnrollOracleNodeIterator struct {
	Event *OracleContractEnrollOracleNode // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OracleContractEnrollOracleNodeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleContractEnrollOracleNode)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OracleContractEnrollOracleNode)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OracleContractEnrollOracleNodeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleContractEnrollOracleNodeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleContractEnrollOracleNode represents a EnrollOracleNode event raised by the OracleContract contract.
type OracleContractEnrollOracleNode struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEnrollOracleNode is a free log retrieval operation binding the contract event 0x317409294d66ea14b34f1d49eba1d450e59340b3ca893dd2131ba2989a0efa22.
//
// Solidity: event enrollOracleNode(address indexed sender)
func (_OracleContract *OracleContractFilterer) FilterEnrollOracleNode(opts *bind.FilterOpts, sender []common.Address) (*OracleContractEnrollOracleNodeIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _OracleContract.contract.FilterLogs(opts, "enrollOracleNode", senderRule)
	if err != nil {
		return nil, err
	}
	return &OracleContractEnrollOracleNodeIterator{contract: _OracleContract.contract, event: "enrollOracleNode", logs: logs, sub: sub}, nil
}

// WatchEnrollOracleNode is a free log subscription operation binding the contract event 0x317409294d66ea14b34f1d49eba1d450e59340b3ca893dd2131ba2989a0efa22.
//
// Solidity: event enrollOracleNode(address indexed sender)
func (_OracleContract *OracleContractFilterer) WatchEnrollOracleNode(opts *bind.WatchOpts, sink chan<- *OracleContractEnrollOracleNode, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _OracleContract.contract.WatchLogs(opts, "enrollOracleNode", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleContractEnrollOracleNode)
				if err := _OracleContract.contract.UnpackLog(event, "enrollOracleNode", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEnrollOracleNode is a log parse operation binding the contract event 0x317409294d66ea14b34f1d49eba1d450e59340b3ca893dd2131ba2989a0efa22.
//
// Solidity: event enrollOracleNode(address indexed sender)
func (_OracleContract *OracleContractFilterer) ParseEnrollOracleNode(log types.Log) (*OracleContractEnrollOracleNode, error) {
	event := new(OracleContractEnrollOracleNode)
	if err := _OracleContract.contract.UnpackLog(event, "enrollOracleNode", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
