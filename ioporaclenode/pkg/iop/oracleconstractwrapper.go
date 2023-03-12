package iop

import (
	"fmt"
	"math/big"

)

type OracleContractWrapper struct {
	*OracleContract
}

// 开销考虑，报名的时候，只登记地址，然后通过该函数，找到报名的节点的地址就返回
func (n *OracleContractWrapper) FindEnrollNodes() ([]OracleContractEnrollNode , error){
	count, err := n.CountEnrollNodes(nil)
	if err != nil {
		return nil, fmt.Errorf("count oracle nodes: %w", err)
	}
	nodeEnrolls := make([]OracleContractEnrollNode, count.Int64())
	for i := int64(0); i < count.Int64(); i++ {
		enrollNode, err := n.FindEnrollNodeByIndex(nil, big.NewInt(i))
		if err != nil {
			return nil, fmt.Errorf("find enrolloracle node by index %d: %w", i, err)
		}
		
		nodeEnrolls[i] = enrollNode
	}
	return nodeEnrolls, nil
}