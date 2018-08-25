// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package native

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// PrivateBankABI is the input ABI used to generate the binding from.
const PrivateBankABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pk\",\"type\":\"bytes32\"},{\"name\":\"rho\",\"type\":\"bytes32\"},{\"name\":\"value\",\"type\":\"uint64\"}],\"name\":\"broadcastNote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"shieldedTransferCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"cm\",\"type\":\"bytes32\"}],\"name\":\"getWitness\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"unshieldingCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\"},{\"name\":\"anchor\",\"type\":\"bytes32\"},{\"name\":\"spend_nf_1\",\"type\":\"bytes32\"},{\"name\":\"spend_nf_2\",\"type\":\"bytes32\"},{\"name\":\"send_nf_1\",\"type\":\"bytes32\"},{\"name\":\"send_nf_2\",\"type\":\"bytes32\"},{\"name\":\"cm_1\",\"type\":\"bytes32\"},{\"name\":\"cm_2\",\"type\":\"bytes32\"}],\"name\":\"shieldedTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"shieldingCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\"},{\"name\":\"spend_nf\",\"type\":\"bytes32\"},{\"name\":\"cm\",\"type\":\"bytes32\"},{\"name\":\"rt\",\"type\":\"bytes32\"},{\"name\":\"value\",\"type\":\"uint64\"}],\"name\":\"unshield\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\"},{\"name\":\"send_nf\",\"type\":\"bytes32\"},{\"name\":\"cm\",\"type\":\"bytes32\"}],\"name\":\"shield\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ZTOKEN_TREE_DEPTH\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"LogShielding\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"LogUnshielding\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"cm1\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"cm2\",\"type\":\"bytes32\"}],\"name\":\"LogShieldedTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"spendNullifier\",\"type\":\"bytes32\"}],\"name\":\"LogSpendNullifier\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"pk\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"rho\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint64\"}],\"name\":\"LogNewNote\",\"type\":\"event\"}]"

// PrivateBankBin is the compiled bytecode used for deploying new contracts.
const PrivateBankBin = `0x608060405234801561001057600080fd5b506000805561001d6100ad565b604051809103906000f080158015610039573d6000803e3d6000fd5b5060058054600160a060020a031916600160a060020a0392909216919091179055601d6100646100bd565b90815260405190819003602001906000f080158015610087573d6000803e3d6000fd5b5060048054600160a060020a031916600160a060020a03929092169190911790556100cd565b6040516103f98061113283390190565b604051610fc58061152b83390190565b611056806100dc6000396000f3006080604052600436106100a35763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166318160ddd81146100b5578063263cce1b146100dc57806341a42bea1461010657806347e722461461011b5780634e3cb2e9146101935780635e2aca56146101a85780637a14db1614610224578063835278b714610239578063b84e7724146102ae578063ce0a857214610303575b3480156100af57600080fd5b50600080fd5b3480156100c157600080fd5b506100ca610318565b60408051918252519081900360200190f35b3480156100e857600080fd5b5061010460043560243567ffffffffffffffff6044351661031e565b005b34801561011257600080fd5b506100ca61036c565b34801561012757600080fd5b50610133600435610372565b604080518481529081018290526060602080830182815285519284019290925284516080840191868101910280838360005b8381101561017d578181015183820152602001610165565b5050505090500194505050505060405180910390f35b34801561019f57600080fd5b506100ca610484565b3480156101b457600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261010494369492936024939284019190819084018382808284375094975050843595505050602083013592604081013592506060810135915060808101359060a08101359060c0013561048a565b34801561023057600080fd5b506100ca610921565b34801561024557600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261010494369492936024939284019190819084018382808284375094975050843595505050602083013592604081013592506060013567ffffffffffffffff169050610927565b6040805160206004803580820135601f8101849004840285018401909552848452610104943694929360249392840191908190840183828082843750949750508435955050506020909201359150610be09050565b34801561030f57600080fd5b506100ca610fd5565b60005481565b604080518481526020810184905267ffffffffffffffff83168183015290517f9ffdcb2deca08add84691e31fcaf8d45a7c62cd186a75416fd10619602e2f2ee9181900360600190a1505050565b60035481565b60048054604080517f47e72246000000000000000000000000000000000000000000000000000000008152928301849052516000926060928492600160a060020a03909116916347e72246916024808301928692919082900301818387803b1580156103dd57600080fd5b505af11580156103f1573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052606081101561041a57600080fd5b81516020830180519193928301929164010000000081111561043b57600080fd5b8201602081018481111561044e57600080fd5b815185602082028301116401000000008211171561046b57600080fd5b5050602090910151939650945091925050509193909250565b60025481565b600084815260066020526040902054156104a357600080fd5b600083815260066020526040902054156104bc57600080fd5b600086815260066020526040902054156104d557600080fd5b600085815260066020526040902054156104ee57600080fd5b600480546040805160008051602061100b833981519152815292830185905251600160a060020a039091169163863078659160248083019260209291908290030181600087803b15801561054157600080fd5b505af1158015610555573d6000803e3d6000fd5b505050506040513d602081101561056b57600080fd5b50511561057757600080fd5b600480546040805160008051602061100b833981519152815292830184905251600160a060020a039091169163863078659160248083019260209291908290030181600087803b1580156105ca57600080fd5b505af11580156105de573d6000803e3d6000fd5b505050506040513d60208110156105f457600080fd5b50511561060057600080fd5b6005546040517f794f69e30000000000000000000000000000000000000000000000000000000081526024810189905260448101889052606481018790526084810186905260a4810185905260c4810184905260e48101839052610100600482019081528a516101048301528a51600160a060020a039093169263794f69e3928c928c928c928c928c928c928c928c92909182916101249091019060208c019080838360005b838110156106be5781810151838201526020016106a6565b50505050905090810190601f1680156106eb5780820380516001836020036101000a031916815260200191505b509950505050505050505050602060405180830381600087803b15801561071157600080fd5b505af1158015610725573d6000803e3d6000fd5b505050506040513d602081101561073b57600080fd5b5051151561074557fe5b60048054604080517f9f7ad1d100000000000000000000000000000000000000000000000000000000815292830185905251600160a060020a0390911691639f7ad1d191602480830192600092919082900301818387803b1580156107a957600080fd5b505af11580156107bd573d6000803e3d6000fd5b505060048054604080517f9f7ad1d100000000000000000000000000000000000000000000000000000000815292830186905251600160a060020a039091169350639f7ad1d19250602480830192600092919082900301818387803b15801561082557600080fd5b505af1158015610839573d6000803e3d6000fd5b50505060008581526006602090815260408083206001908190558784528184208190558a84528184208190558984529281902092909255815189815291517fc7d4658ff28174cdb265eb970264466b3a14c413a2e2f3a5bbcab44427eb90729350918290030190a16040805186815290517fc7d4658ff28174cdb265eb970264466b3a14c413a2e2f3a5bbcab44427eb90729181900360200190a16040805183815260208101839052815133927f0d3211cecea8fbffc05221205fa1aefe539b636c658a996a6cbc8729fe1ad001928290030190a25050600380546001019055505050505050565b60015481565b6000848152600660205260408120541561094057600080fd5b600480546040805160008051602061100b833981519152815292830187905251600160a060020a039091169163863078659160248083019260209291908290030181600087803b15801561099357600080fd5b505af11580156109a7573d6000803e3d6000fd5b505050506040513d60208110156109bd57600080fd5b505115156109ca57600080fd5b6005546040517fdafca47a000000000000000000000000000000000000000000000000000000008152602481018790526044810185905267ffffffffffffffff84166064820152608060048201908152885160848301528851600160a060020a039093169263dafca47a928a928a928992899291829160a40190602088019080838360005b83811015610a67578181015183820152602001610a4f565b50505050905090810190601f168015610a945780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b158015610ab657600080fd5b505af1158015610aca573d6000803e3d6000fd5b505050506040513d6020811015610ae057600080fd5b50511515610aea57fe5b506000848152600660205260408082206001905551670de0b6b3a764000067ffffffffffffffff84160291339183156108fc0291849190818181858888f19350505050158015610b3e573d6000803e3d6000fd5b50610b556000548367ffffffffffffffff16610fda565b6000556040805186815290517fc7d4658ff28174cdb265eb970264466b3a14c413a2e2f3a5bbcab44427eb90729181900360200190a16040805167ffffffffffffffff8416815260208101869052815133927f44ed7e2197490271743d80ef2940b5ef22d886361e9634e9696842ccd2acfacf928290030190a2505060028054600101905550505050565b600082815260066020526040902054670de0b6b3a764000034049015610c6757604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f73656e64206e756c6c6966696572206578697374730000000000000000000000604482015290519081900360640190fd5b600480546040805160008051602061100b833981519152815292830185905251600160a060020a039091169163863078659160248083019260209291908290030181600087803b158015610cba57600080fd5b505af1158015610cce573d6000803e3d6000fd5b505050506040513d6020811015610ce457600080fd5b505115610d5257604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f636f6d6d69746d656e7420657869737473000000000000000000000000000000604482015290519081900360640190fd5b6005546040517fe088659e000000000000000000000000000000000000000000000000000000008152602481018590526044810184905267ffffffffffffffff83166064820152608060048201908152865160848301528651600160a060020a039093169263e088659e92889288928892889291829160a40190602088019080838360005b83811015610def578181015183820152602001610dd7565b50505050905090810190601f168015610e1c5780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b158015610e3e57600080fd5b505af1158015610e52573d6000803e3d6000fd5b505050506040513d6020811015610e6857600080fd5b50511515610ed757604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f636f756c64206e6f742076616c696461746520736869656c64696e6700000000604482015290519081900360640190fd5b60048054604080517f9f7ad1d100000000000000000000000000000000000000000000000000000000815292830185905251600160a060020a0390911691639f7ad1d191602480830192600092919082900301818387803b158015610f3b57600080fd5b505af1158015610f4f573d6000803e3d6000fd5b50505060008481526006602090815260409182902060019055815167ffffffffffffffff8516815290810185905281513393507f85b7bd30efc09a22e16ff1baa680581ec4bc1dd7a889ac239355ff19dc34f5e3929181900390910190a26001805481019055600054610fcc9067ffffffffffffffff8316610ff1565b60005550505050565b601d81565b60008083831115610fea57600080fd5b5050900390565b60008282018381101561100357600080fd5b939250505056008630786500000000000000000000000000000000000000000000000000000000a165627a7a72305820ce8f5433597cc7cd767c9caa9b9a96572b9405d91ee748fb9a07d2c3f3bc1ef30029608060405234801561001057600080fd5b506103d9806100206000396000f3006080604052600436106100615763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633b300c688114610073578063794f69e3146100de578063dafca47a1461016e578063e088659e146101de575b34801561006d57600080fd5b50600080fd5b34801561007f57600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526100cc94369492936024939284019190819084018382808284375094975061024e9650505050505050565b60408051918252519081900360200190f35b3480156100ea57600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261015a94369492936024939284019190819084018382808284375094975050843595505050602083013592604081013592506060810135915060808101359060a08101359060c0013561028a565b604080519115158252519081900360200190f35b34801561017a57600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261015a943694929360249392840191908190840183828082843750949750508435955050506020830135926040013567ffffffffffffffff16915061030c9050565b3480156101ea57600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261015a943694929360249392840191908190840183828082843750949750508435955050506020830135926040013567ffffffffffffffff16915061038a9050565b805160009060401461025f57600080fd5b60405160208160808560006188015af1151561027a57600080fd5b8051602490910160405292915050565b600080604051368060048337602082828460006188025af115156102ad57600080fd5b8151910160405260001a7f01000000000000000000000000000000000000000000000000000000000000009081027fff0000000000000000000000000000000000000000000000000000000000000016149a9950505050505050505050565b600080604051368060048337602082828460006188045af1151561032f57600080fd5b8151910160405260001a7f01000000000000000000000000000000000000000000000000000000000000009081027fff0000000000000000000000000000000000000000000000000000000000000016149695505050505050565b600080604051368060048337602082828460006188035af1151561032f57600080fd00a165627a7a72305820923b26ed546e528b1a1407267c1697ecf6460025021dc366696fe95995e4106000296080604052600060045534801561001557600080fd5b5060405160208062000fc58339810160405251610030610302565b604051809103906000f08015801561004c573d6000803e3d6000fd5b5060008054600160a060020a031916600160a060020a03929092169190911790556001819055600281810a905561008b81640100000000610091810204565b50610313565b60038054600181018255600091825260008051602062000fa583398151915201819055805b60018303811015610102576100d48280640100000000610107810204565b600380546001818101835560009290925260008051602062000fa583398151915201829055909250016100b6565b505050565b6040805181815260608082018352600092909183916020820161080080388339019050509150600090505b60208110156101985784816020811061014757fe5b1a7f010000000000000000000000000000000000000000000000000000000000000002828281518110151561017857fe5b906020010190600160f860020a031916908160001a905350600101610132565b5060005b6020811015610205578381602081106101b157fe5b1a7f01000000000000000000000000000000000000000000000000000000000000000282826020018151811015156101e557fe5b906020010190600160f860020a031916908160001a90535060010161019c565b600080546040517f3b300c68000000000000000000000000000000000000000000000000000000008152602060048201818152865160248401528651600160a060020a0390941694633b300c68948894929384936044019290860191908190849084905b83811015610281578181015183820152602001610269565b50505050905090810190601f1680156102ae5780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b1580156102cd57600080fd5b505af11580156102e1573d6000803e3d6000fd5b505050506040513d60208110156102f757600080fd5b505195945050505050565b6040516103f98062000bac83390190565b61088980620003236000396000f3006080604052600436106100da5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630af0b13281146100ec5780631d48f5e91461011957806347e722461461013457806348a0d754146101ac5780635ae60499146101c15780635cfc1a51146101dc578063631c56ef146101f15780636def0cbe14610206578063863078651461026b578063949d225d146102975780639f7ad1d1146102ac578063a06f0977146102c6578063c8c50515146102de578063d0316c50146102f6578063ebf0c7171461030e575b3480156100e657600080fd5b50600080fd5b3480156100f857600080fd5b50610107600435602435610323565b60408051918252519081900360200190f35b34801561012557600080fd5b5061010760043560243561033b565b34801561014057600080fd5b5061014c600435610573565b604080518481529081018290526060602080830182815285519284019290925284516080840191868101910280838360005b8381101561019657818101518382015260200161017e565b5050505090500194505050505060405180910390f35b3480156101b857600080fd5b50610107610640565b3480156101cd57600080fd5b5061010760043560243561064a565b3480156101e857600080fd5b50610107610651565b3480156101fd57600080fd5b50610107610657565b34801561021257600080fd5b5061021b61065d565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561025757818101518382015260200161023f565b505050509050019250505060405180910390f35b34801561027757600080fd5b506102836004356106b6565b604080519115158252519081900360200190f35b3480156102a357600080fd5b506101076106ca565b3480156102b857600080fd5b506102c46004356106d0565b005b3480156102d257600080fd5b50610107600435610724565b3480156102ea57600080fd5b50610107600435610755565b34801561030257600080fd5b50610107600435610779565b34801561031a57600080fd5b506101076107a3565b60008160020a8381151561033357fe5b049392505050565b6040805181815260608082018352600092909183916020820161080080388339019050509150600090505b60208110156103e45784816020811061037b57fe5b1a7f01000000000000000000000000000000000000000000000000000000000000000282828151811015156103ac57fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600101610366565b5060005b6020811015610469578381602081106103fd57fe5b1a7f010000000000000000000000000000000000000000000000000000000000000002828260200181518110151561043157fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506001016103e8565b600080546040517f3b300c6800000000000000000000000000000000000000000000000000000000815260206004820181815286516024840152865173ffffffffffffffffffffffffffffffffffffffff90941694633b300c68948894929384936044019290860191908190849084905b838110156104f25781810151838201526020016104da565b50505050905090810190601f16801561051f5780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b15801561053e57600080fd5b505af1158015610552573d6000803e3d6000fd5b505050506040513d602081101561056857600080fd5b505195945050505050565b6000818152600560205260408120546060908290818381808080861161059857600080fd5b6001860394506001546040519080825280602002602001820160405280156105ca578160200160208202803883390190505b50935060009250849150600090505b600154831015610623576105f682600118848060010195506107b7565b845160018301928691811061060757fe5b6020908102909101015261061c826001610323565b91506105d9565b848461062d6107a3565b9850985098505050505050509193909250565b6004546002540390565b60020a0290565b60025490565b60015490565b606060038054806020026020016040519081016040528092919081815260200182805480156106ac57602002820191906000526020600020905b81548152600190910190602001808311610697575b5050505050905090565b600090815260056020526040902054151590565b60045490565b600081815260056020526040812054156106e957600080fd5b600254600454106106f957600080fd5b5060048054600101908190556000828152600560209081526040808320849055928252600690522055565b600354600090821061073557600080fd5b600380548390811061074357fe5b90600052602060002001549050919050565b60008181526005602052604081205481811161077057600080fd5b60010192915050565b6000806004548310151561078c57600080fd5b505060010160009081526006602052604090205490565b60006107b260006001546107b7565b905090565b6000806000806107c7868661064a565b600454116107ef5760038054869081106107dd57fe5b90600052602060002001549350610854565b841515610812576001860160008181526006602052604090205494509250610854565b61082961082087600161064a565b600187036107b7565b915061084561083987600161064a565b600101600187036107b7565b9050610851828261033b565b93505b505050929150505600a165627a7a72305820f7db65d632a6ab6572a7427e1b49b5fbe2f7810c1129ecb437692df4d32167b20029608060405234801561001057600080fd5b506103d9806100206000396000f3006080604052600436106100615763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633b300c688114610073578063794f69e3146100de578063dafca47a1461016e578063e088659e146101de575b34801561006d57600080fd5b50600080fd5b34801561007f57600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526100cc94369492936024939284019190819084018382808284375094975061024e9650505050505050565b60408051918252519081900360200190f35b3480156100ea57600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261015a94369492936024939284019190819084018382808284375094975050843595505050602083013592604081013592506060810135915060808101359060a08101359060c0013561028a565b604080519115158252519081900360200190f35b34801561017a57600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261015a943694929360249392840191908190840183828082843750949750508435955050506020830135926040013567ffffffffffffffff16915061030c9050565b3480156101ea57600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261015a943694929360249392840191908190840183828082843750949750508435955050506020830135926040013567ffffffffffffffff16915061038a9050565b805160009060401461025f57600080fd5b60405160208160808560006188015af1151561027a57600080fd5b8051602490910160405292915050565b600080604051368060048337602082828460006188025af115156102ad57600080fd5b8151910160405260001a7f01000000000000000000000000000000000000000000000000000000000000009081027fff0000000000000000000000000000000000000000000000000000000000000016149a9950505050505050505050565b600080604051368060048337602082828460006188045af1151561032f57600080fd5b8151910160405260001a7f01000000000000000000000000000000000000000000000000000000000000009081027fff0000000000000000000000000000000000000000000000000000000000000016149695505050505050565b600080604051368060048337602082828460006188035af1151561032f57600080fd00a165627a7a72305820923b26ed546e528b1a1407267c1697ecf6460025021dc366696fe95995e410600029c2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b`

// DeployPrivateBank deploys a new Ethereum contract, binding an instance of PrivateBank to it.
func DeployPrivateBank(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PrivateBank, error) {
	parsed, err := abi.JSON(strings.NewReader(PrivateBankABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PrivateBankBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PrivateBank{PrivateBankCaller: PrivateBankCaller{contract: contract}, PrivateBankTransactor: PrivateBankTransactor{contract: contract}, PrivateBankFilterer: PrivateBankFilterer{contract: contract}}, nil
}

// PrivateBank is an auto generated Go binding around an Ethereum contract.
type PrivateBank struct {
	PrivateBankCaller     // Read-only binding to the contract
	PrivateBankTransactor // Write-only binding to the contract
	PrivateBankFilterer   // Log filterer for contract events
}

// PrivateBankCaller is an auto generated read-only Go binding around an Ethereum contract.
type PrivateBankCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrivateBankTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PrivateBankTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrivateBankFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PrivateBankFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrivateBankSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PrivateBankSession struct {
	Contract     *PrivateBank      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PrivateBankCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PrivateBankCallerSession struct {
	Contract *PrivateBankCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PrivateBankTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PrivateBankTransactorSession struct {
	Contract     *PrivateBankTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PrivateBankRaw is an auto generated low-level Go binding around an Ethereum contract.
type PrivateBankRaw struct {
	Contract *PrivateBank // Generic contract binding to access the raw methods on
}

// PrivateBankCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PrivateBankCallerRaw struct {
	Contract *PrivateBankCaller // Generic read-only contract binding to access the raw methods on
}

// PrivateBankTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PrivateBankTransactorRaw struct {
	Contract *PrivateBankTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPrivateBank creates a new instance of PrivateBank, bound to a specific deployed contract.
func NewPrivateBank(address common.Address, backend bind.ContractBackend) (*PrivateBank, error) {
	contract, err := bindPrivateBank(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PrivateBank{PrivateBankCaller: PrivateBankCaller{contract: contract}, PrivateBankTransactor: PrivateBankTransactor{contract: contract}, PrivateBankFilterer: PrivateBankFilterer{contract: contract}}, nil
}

// NewPrivateBankCaller creates a new read-only instance of PrivateBank, bound to a specific deployed contract.
func NewPrivateBankCaller(address common.Address, caller bind.ContractCaller) (*PrivateBankCaller, error) {
	contract, err := bindPrivateBank(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PrivateBankCaller{contract: contract}, nil
}

// NewPrivateBankTransactor creates a new write-only instance of PrivateBank, bound to a specific deployed contract.
func NewPrivateBankTransactor(address common.Address, transactor bind.ContractTransactor) (*PrivateBankTransactor, error) {
	contract, err := bindPrivateBank(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PrivateBankTransactor{contract: contract}, nil
}

// NewPrivateBankFilterer creates a new log filterer instance of PrivateBank, bound to a specific deployed contract.
func NewPrivateBankFilterer(address common.Address, filterer bind.ContractFilterer) (*PrivateBankFilterer, error) {
	contract, err := bindPrivateBank(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PrivateBankFilterer{contract: contract}, nil
}

// bindPrivateBank binds a generic wrapper to an already deployed contract.
func bindPrivateBank(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PrivateBankABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PrivateBank *PrivateBankRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PrivateBank.Contract.PrivateBankCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PrivateBank *PrivateBankRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrivateBank.Contract.PrivateBankTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PrivateBank *PrivateBankRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PrivateBank.Contract.PrivateBankTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PrivateBank *PrivateBankCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PrivateBank.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PrivateBank *PrivateBankTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrivateBank.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PrivateBank *PrivateBankTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PrivateBank.Contract.contract.Transact(opts, method, params...)
}

// ZTOKENTREEDEPTH is a free data retrieval call binding the contract method 0xce0a8572.
//
// Solidity: function ZTOKEN_TREE_DEPTH() constant returns(uint256)
func (_PrivateBank *PrivateBankCaller) ZTOKENTREEDEPTH(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PrivateBank.contract.Call(opts, out, "ZTOKEN_TREE_DEPTH")
	return *ret0, err
}

// ZTOKENTREEDEPTH is a free data retrieval call binding the contract method 0xce0a8572.
//
// Solidity: function ZTOKEN_TREE_DEPTH() constant returns(uint256)
func (_PrivateBank *PrivateBankSession) ZTOKENTREEDEPTH() (*big.Int, error) {
	return _PrivateBank.Contract.ZTOKENTREEDEPTH(&_PrivateBank.CallOpts)
}

// ZTOKENTREEDEPTH is a free data retrieval call binding the contract method 0xce0a8572.
//
// Solidity: function ZTOKEN_TREE_DEPTH() constant returns(uint256)
func (_PrivateBank *PrivateBankCallerSession) ZTOKENTREEDEPTH() (*big.Int, error) {
	return _PrivateBank.Contract.ZTOKENTREEDEPTH(&_PrivateBank.CallOpts)
}

// GetWitness is a free data retrieval call binding the contract method 0x47e72246.
//
// Solidity: function getWitness(cm bytes32) constant returns(uint256, bytes32[], bytes32)
func (_PrivateBank *PrivateBankCaller) GetWitness(opts *bind.CallOpts, cm [32]byte) (*big.Int, [][32]byte, [32]byte, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new([][32]byte)
		ret2 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _PrivateBank.contract.Call(opts, out, "getWitness", cm)
	return *ret0, *ret1, *ret2, err
}

// GetWitness is a free data retrieval call binding the contract method 0x47e72246.
//
// Solidity: function getWitness(cm bytes32) constant returns(uint256, bytes32[], bytes32)
func (_PrivateBank *PrivateBankSession) GetWitness(cm [32]byte) (*big.Int, [][32]byte, [32]byte, error) {
	return _PrivateBank.Contract.GetWitness(&_PrivateBank.CallOpts, cm)
}

// GetWitness is a free data retrieval call binding the contract method 0x47e72246.
//
// Solidity: function getWitness(cm bytes32) constant returns(uint256, bytes32[], bytes32)
func (_PrivateBank *PrivateBankCallerSession) GetWitness(cm [32]byte) (*big.Int, [][32]byte, [32]byte, error) {
	return _PrivateBank.Contract.GetWitness(&_PrivateBank.CallOpts, cm)
}

// ShieldedTransferCount is a free data retrieval call binding the contract method 0x41a42bea.
//
// Solidity: function shieldedTransferCount() constant returns(uint256)
func (_PrivateBank *PrivateBankCaller) ShieldedTransferCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PrivateBank.contract.Call(opts, out, "shieldedTransferCount")
	return *ret0, err
}

// ShieldedTransferCount is a free data retrieval call binding the contract method 0x41a42bea.
//
// Solidity: function shieldedTransferCount() constant returns(uint256)
func (_PrivateBank *PrivateBankSession) ShieldedTransferCount() (*big.Int, error) {
	return _PrivateBank.Contract.ShieldedTransferCount(&_PrivateBank.CallOpts)
}

// ShieldedTransferCount is a free data retrieval call binding the contract method 0x41a42bea.
//
// Solidity: function shieldedTransferCount() constant returns(uint256)
func (_PrivateBank *PrivateBankCallerSession) ShieldedTransferCount() (*big.Int, error) {
	return _PrivateBank.Contract.ShieldedTransferCount(&_PrivateBank.CallOpts)
}

// ShieldingCount is a free data retrieval call binding the contract method 0x7a14db16.
//
// Solidity: function shieldingCount() constant returns(uint256)
func (_PrivateBank *PrivateBankCaller) ShieldingCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PrivateBank.contract.Call(opts, out, "shieldingCount")
	return *ret0, err
}

// ShieldingCount is a free data retrieval call binding the contract method 0x7a14db16.
//
// Solidity: function shieldingCount() constant returns(uint256)
func (_PrivateBank *PrivateBankSession) ShieldingCount() (*big.Int, error) {
	return _PrivateBank.Contract.ShieldingCount(&_PrivateBank.CallOpts)
}

// ShieldingCount is a free data retrieval call binding the contract method 0x7a14db16.
//
// Solidity: function shieldingCount() constant returns(uint256)
func (_PrivateBank *PrivateBankCallerSession) ShieldingCount() (*big.Int, error) {
	return _PrivateBank.Contract.ShieldingCount(&_PrivateBank.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_PrivateBank *PrivateBankCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PrivateBank.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_PrivateBank *PrivateBankSession) TotalSupply() (*big.Int, error) {
	return _PrivateBank.Contract.TotalSupply(&_PrivateBank.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_PrivateBank *PrivateBankCallerSession) TotalSupply() (*big.Int, error) {
	return _PrivateBank.Contract.TotalSupply(&_PrivateBank.CallOpts)
}

// UnshieldingCount is a free data retrieval call binding the contract method 0x4e3cb2e9.
//
// Solidity: function unshieldingCount() constant returns(uint256)
func (_PrivateBank *PrivateBankCaller) UnshieldingCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PrivateBank.contract.Call(opts, out, "unshieldingCount")
	return *ret0, err
}

// UnshieldingCount is a free data retrieval call binding the contract method 0x4e3cb2e9.
//
// Solidity: function unshieldingCount() constant returns(uint256)
func (_PrivateBank *PrivateBankSession) UnshieldingCount() (*big.Int, error) {
	return _PrivateBank.Contract.UnshieldingCount(&_PrivateBank.CallOpts)
}

// UnshieldingCount is a free data retrieval call binding the contract method 0x4e3cb2e9.
//
// Solidity: function unshieldingCount() constant returns(uint256)
func (_PrivateBank *PrivateBankCallerSession) UnshieldingCount() (*big.Int, error) {
	return _PrivateBank.Contract.UnshieldingCount(&_PrivateBank.CallOpts)
}

// BroadcastNote is a paid mutator transaction binding the contract method 0x263cce1b.
//
// Solidity: function broadcastNote(pk bytes32, rho bytes32, value uint64) returns()
func (_PrivateBank *PrivateBankTransactor) BroadcastNote(opts *bind.TransactOpts, pk [32]byte, rho [32]byte, value uint64) (*types.Transaction, error) {
	return _PrivateBank.contract.Transact(opts, "broadcastNote", pk, rho, value)
}

// BroadcastNote is a paid mutator transaction binding the contract method 0x263cce1b.
//
// Solidity: function broadcastNote(pk bytes32, rho bytes32, value uint64) returns()
func (_PrivateBank *PrivateBankSession) BroadcastNote(pk [32]byte, rho [32]byte, value uint64) (*types.Transaction, error) {
	return _PrivateBank.Contract.BroadcastNote(&_PrivateBank.TransactOpts, pk, rho, value)
}

// BroadcastNote is a paid mutator transaction binding the contract method 0x263cce1b.
//
// Solidity: function broadcastNote(pk bytes32, rho bytes32, value uint64) returns()
func (_PrivateBank *PrivateBankTransactorSession) BroadcastNote(pk [32]byte, rho [32]byte, value uint64) (*types.Transaction, error) {
	return _PrivateBank.Contract.BroadcastNote(&_PrivateBank.TransactOpts, pk, rho, value)
}

// Shield is a paid mutator transaction binding the contract method 0xb84e7724.
//
// Solidity: function shield(proof bytes, send_nf bytes32, cm bytes32) returns()
func (_PrivateBank *PrivateBankTransactor) Shield(opts *bind.TransactOpts, proof []byte, send_nf [32]byte, cm [32]byte) (*types.Transaction, error) {
	return _PrivateBank.contract.Transact(opts, "shield", proof, send_nf, cm)
}

// Shield is a paid mutator transaction binding the contract method 0xb84e7724.
//
// Solidity: function shield(proof bytes, send_nf bytes32, cm bytes32) returns()
func (_PrivateBank *PrivateBankSession) Shield(proof []byte, send_nf [32]byte, cm [32]byte) (*types.Transaction, error) {
	return _PrivateBank.Contract.Shield(&_PrivateBank.TransactOpts, proof, send_nf, cm)
}

// Shield is a paid mutator transaction binding the contract method 0xb84e7724.
//
// Solidity: function shield(proof bytes, send_nf bytes32, cm bytes32) returns()
func (_PrivateBank *PrivateBankTransactorSession) Shield(proof []byte, send_nf [32]byte, cm [32]byte) (*types.Transaction, error) {
	return _PrivateBank.Contract.Shield(&_PrivateBank.TransactOpts, proof, send_nf, cm)
}

// ShieldedTransfer is a paid mutator transaction binding the contract method 0x5e2aca56.
//
// Solidity: function shieldedTransfer(proof bytes, anchor bytes32, spend_nf_1 bytes32, spend_nf_2 bytes32, send_nf_1 bytes32, send_nf_2 bytes32, cm_1 bytes32, cm_2 bytes32) returns()
func (_PrivateBank *PrivateBankTransactor) ShieldedTransfer(opts *bind.TransactOpts, proof []byte, anchor [32]byte, spend_nf_1 [32]byte, spend_nf_2 [32]byte, send_nf_1 [32]byte, send_nf_2 [32]byte, cm_1 [32]byte, cm_2 [32]byte) (*types.Transaction, error) {
	return _PrivateBank.contract.Transact(opts, "shieldedTransfer", proof, anchor, spend_nf_1, spend_nf_2, send_nf_1, send_nf_2, cm_1, cm_2)
}

// ShieldedTransfer is a paid mutator transaction binding the contract method 0x5e2aca56.
//
// Solidity: function shieldedTransfer(proof bytes, anchor bytes32, spend_nf_1 bytes32, spend_nf_2 bytes32, send_nf_1 bytes32, send_nf_2 bytes32, cm_1 bytes32, cm_2 bytes32) returns()
func (_PrivateBank *PrivateBankSession) ShieldedTransfer(proof []byte, anchor [32]byte, spend_nf_1 [32]byte, spend_nf_2 [32]byte, send_nf_1 [32]byte, send_nf_2 [32]byte, cm_1 [32]byte, cm_2 [32]byte) (*types.Transaction, error) {
	return _PrivateBank.Contract.ShieldedTransfer(&_PrivateBank.TransactOpts, proof, anchor, spend_nf_1, spend_nf_2, send_nf_1, send_nf_2, cm_1, cm_2)
}

// ShieldedTransfer is a paid mutator transaction binding the contract method 0x5e2aca56.
//
// Solidity: function shieldedTransfer(proof bytes, anchor bytes32, spend_nf_1 bytes32, spend_nf_2 bytes32, send_nf_1 bytes32, send_nf_2 bytes32, cm_1 bytes32, cm_2 bytes32) returns()
func (_PrivateBank *PrivateBankTransactorSession) ShieldedTransfer(proof []byte, anchor [32]byte, spend_nf_1 [32]byte, spend_nf_2 [32]byte, send_nf_1 [32]byte, send_nf_2 [32]byte, cm_1 [32]byte, cm_2 [32]byte) (*types.Transaction, error) {
	return _PrivateBank.Contract.ShieldedTransfer(&_PrivateBank.TransactOpts, proof, anchor, spend_nf_1, spend_nf_2, send_nf_1, send_nf_2, cm_1, cm_2)
}

// Unshield is a paid mutator transaction binding the contract method 0x835278b7.
//
// Solidity: function unshield(proof bytes, spend_nf bytes32, cm bytes32, rt bytes32, value uint64) returns()
func (_PrivateBank *PrivateBankTransactor) Unshield(opts *bind.TransactOpts, proof []byte, spend_nf [32]byte, cm [32]byte, rt [32]byte, value uint64) (*types.Transaction, error) {
	return _PrivateBank.contract.Transact(opts, "unshield", proof, spend_nf, cm, rt, value)
}

// Unshield is a paid mutator transaction binding the contract method 0x835278b7.
//
// Solidity: function unshield(proof bytes, spend_nf bytes32, cm bytes32, rt bytes32, value uint64) returns()
func (_PrivateBank *PrivateBankSession) Unshield(proof []byte, spend_nf [32]byte, cm [32]byte, rt [32]byte, value uint64) (*types.Transaction, error) {
	return _PrivateBank.Contract.Unshield(&_PrivateBank.TransactOpts, proof, spend_nf, cm, rt, value)
}

// Unshield is a paid mutator transaction binding the contract method 0x835278b7.
//
// Solidity: function unshield(proof bytes, spend_nf bytes32, cm bytes32, rt bytes32, value uint64) returns()
func (_PrivateBank *PrivateBankTransactorSession) Unshield(proof []byte, spend_nf [32]byte, cm [32]byte, rt [32]byte, value uint64) (*types.Transaction, error) {
	return _PrivateBank.Contract.Unshield(&_PrivateBank.TransactOpts, proof, spend_nf, cm, rt, value)
}

// PrivateBankLogNewNoteIterator is returned from FilterLogNewNote and is used to iterate over the raw logs and unpacked data for LogNewNote events raised by the PrivateBank contract.
type PrivateBankLogNewNoteIterator struct {
	Event *PrivateBankLogNewNote // Event containing the contract specifics and raw log

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
func (it *PrivateBankLogNewNoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivateBankLogNewNote)
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
		it.Event = new(PrivateBankLogNewNote)
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
func (it *PrivateBankLogNewNoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivateBankLogNewNoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivateBankLogNewNote represents a LogNewNote event raised by the PrivateBank contract.
type PrivateBankLogNewNote struct {
	Pk    [32]byte
	Rho   [32]byte
	Value uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogNewNote is a free log retrieval operation binding the contract event 0x9ffdcb2deca08add84691e31fcaf8d45a7c62cd186a75416fd10619602e2f2ee.
//
// Solidity: e LogNewNote(pk bytes32, rho bytes32, value uint64)
func (_PrivateBank *PrivateBankFilterer) FilterLogNewNote(opts *bind.FilterOpts) (*PrivateBankLogNewNoteIterator, error) {

	logs, sub, err := _PrivateBank.contract.FilterLogs(opts, "LogNewNote")
	if err != nil {
		return nil, err
	}
	return &PrivateBankLogNewNoteIterator{contract: _PrivateBank.contract, event: "LogNewNote", logs: logs, sub: sub}, nil
}

// WatchLogNewNote is a free log subscription operation binding the contract event 0x9ffdcb2deca08add84691e31fcaf8d45a7c62cd186a75416fd10619602e2f2ee.
//
// Solidity: e LogNewNote(pk bytes32, rho bytes32, value uint64)
func (_PrivateBank *PrivateBankFilterer) WatchLogNewNote(opts *bind.WatchOpts, sink chan<- *PrivateBankLogNewNote) (event.Subscription, error) {

	logs, sub, err := _PrivateBank.contract.WatchLogs(opts, "LogNewNote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivateBankLogNewNote)
				if err := _PrivateBank.contract.UnpackLog(event, "LogNewNote", log); err != nil {
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

// PrivateBankLogShieldedTransferIterator is returned from FilterLogShieldedTransfer and is used to iterate over the raw logs and unpacked data for LogShieldedTransfer events raised by the PrivateBank contract.
type PrivateBankLogShieldedTransferIterator struct {
	Event *PrivateBankLogShieldedTransfer // Event containing the contract specifics and raw log

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
func (it *PrivateBankLogShieldedTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivateBankLogShieldedTransfer)
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
		it.Event = new(PrivateBankLogShieldedTransfer)
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
func (it *PrivateBankLogShieldedTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivateBankLogShieldedTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivateBankLogShieldedTransfer represents a LogShieldedTransfer event raised by the PrivateBank contract.
type PrivateBankLogShieldedTransfer struct {
	From common.Address
	Cm1  [32]byte
	Cm2  [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogShieldedTransfer is a free log retrieval operation binding the contract event 0x0d3211cecea8fbffc05221205fa1aefe539b636c658a996a6cbc8729fe1ad001.
//
// Solidity: e LogShieldedTransfer(from indexed address, cm1 bytes32, cm2 bytes32)
func (_PrivateBank *PrivateBankFilterer) FilterLogShieldedTransfer(opts *bind.FilterOpts, from []common.Address) (*PrivateBankLogShieldedTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _PrivateBank.contract.FilterLogs(opts, "LogShieldedTransfer", fromRule)
	if err != nil {
		return nil, err
	}
	return &PrivateBankLogShieldedTransferIterator{contract: _PrivateBank.contract, event: "LogShieldedTransfer", logs: logs, sub: sub}, nil
}

// WatchLogShieldedTransfer is a free log subscription operation binding the contract event 0x0d3211cecea8fbffc05221205fa1aefe539b636c658a996a6cbc8729fe1ad001.
//
// Solidity: e LogShieldedTransfer(from indexed address, cm1 bytes32, cm2 bytes32)
func (_PrivateBank *PrivateBankFilterer) WatchLogShieldedTransfer(opts *bind.WatchOpts, sink chan<- *PrivateBankLogShieldedTransfer, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _PrivateBank.contract.WatchLogs(opts, "LogShieldedTransfer", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivateBankLogShieldedTransfer)
				if err := _PrivateBank.contract.UnpackLog(event, "LogShieldedTransfer", log); err != nil {
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

// PrivateBankLogShieldingIterator is returned from FilterLogShielding and is used to iterate over the raw logs and unpacked data for LogShielding events raised by the PrivateBank contract.
type PrivateBankLogShieldingIterator struct {
	Event *PrivateBankLogShielding // Event containing the contract specifics and raw log

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
func (it *PrivateBankLogShieldingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivateBankLogShielding)
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
		it.Event = new(PrivateBankLogShielding)
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
func (it *PrivateBankLogShieldingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivateBankLogShieldingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivateBankLogShielding represents a LogShielding event raised by the PrivateBank contract.
type PrivateBankLogShielding struct {
	From       common.Address
	Value      uint64
	Commitment [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogShielding is a free log retrieval operation binding the contract event 0x85b7bd30efc09a22e16ff1baa680581ec4bc1dd7a889ac239355ff19dc34f5e3.
//
// Solidity: e LogShielding(from indexed address, value uint64, commitment bytes32)
func (_PrivateBank *PrivateBankFilterer) FilterLogShielding(opts *bind.FilterOpts, from []common.Address) (*PrivateBankLogShieldingIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _PrivateBank.contract.FilterLogs(opts, "LogShielding", fromRule)
	if err != nil {
		return nil, err
	}
	return &PrivateBankLogShieldingIterator{contract: _PrivateBank.contract, event: "LogShielding", logs: logs, sub: sub}, nil
}

// WatchLogShielding is a free log subscription operation binding the contract event 0x85b7bd30efc09a22e16ff1baa680581ec4bc1dd7a889ac239355ff19dc34f5e3.
//
// Solidity: e LogShielding(from indexed address, value uint64, commitment bytes32)
func (_PrivateBank *PrivateBankFilterer) WatchLogShielding(opts *bind.WatchOpts, sink chan<- *PrivateBankLogShielding, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _PrivateBank.contract.WatchLogs(opts, "LogShielding", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivateBankLogShielding)
				if err := _PrivateBank.contract.UnpackLog(event, "LogShielding", log); err != nil {
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

// PrivateBankLogSpendNullifierIterator is returned from FilterLogSpendNullifier and is used to iterate over the raw logs and unpacked data for LogSpendNullifier events raised by the PrivateBank contract.
type PrivateBankLogSpendNullifierIterator struct {
	Event *PrivateBankLogSpendNullifier // Event containing the contract specifics and raw log

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
func (it *PrivateBankLogSpendNullifierIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivateBankLogSpendNullifier)
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
		it.Event = new(PrivateBankLogSpendNullifier)
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
func (it *PrivateBankLogSpendNullifierIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivateBankLogSpendNullifierIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivateBankLogSpendNullifier represents a LogSpendNullifier event raised by the PrivateBank contract.
type PrivateBankLogSpendNullifier struct {
	SpendNullifier [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogSpendNullifier is a free log retrieval operation binding the contract event 0xc7d4658ff28174cdb265eb970264466b3a14c413a2e2f3a5bbcab44427eb9072.
//
// Solidity: e LogSpendNullifier(spendNullifier bytes32)
func (_PrivateBank *PrivateBankFilterer) FilterLogSpendNullifier(opts *bind.FilterOpts) (*PrivateBankLogSpendNullifierIterator, error) {

	logs, sub, err := _PrivateBank.contract.FilterLogs(opts, "LogSpendNullifier")
	if err != nil {
		return nil, err
	}
	return &PrivateBankLogSpendNullifierIterator{contract: _PrivateBank.contract, event: "LogSpendNullifier", logs: logs, sub: sub}, nil
}

// WatchLogSpendNullifier is a free log subscription operation binding the contract event 0xc7d4658ff28174cdb265eb970264466b3a14c413a2e2f3a5bbcab44427eb9072.
//
// Solidity: e LogSpendNullifier(spendNullifier bytes32)
func (_PrivateBank *PrivateBankFilterer) WatchLogSpendNullifier(opts *bind.WatchOpts, sink chan<- *PrivateBankLogSpendNullifier) (event.Subscription, error) {

	logs, sub, err := _PrivateBank.contract.WatchLogs(opts, "LogSpendNullifier")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivateBankLogSpendNullifier)
				if err := _PrivateBank.contract.UnpackLog(event, "LogSpendNullifier", log); err != nil {
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

// PrivateBankLogUnshieldingIterator is returned from FilterLogUnshielding and is used to iterate over the raw logs and unpacked data for LogUnshielding events raised by the PrivateBank contract.
type PrivateBankLogUnshieldingIterator struct {
	Event *PrivateBankLogUnshielding // Event containing the contract specifics and raw log

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
func (it *PrivateBankLogUnshieldingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivateBankLogUnshielding)
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
		it.Event = new(PrivateBankLogUnshielding)
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
func (it *PrivateBankLogUnshieldingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivateBankLogUnshieldingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivateBankLogUnshielding represents a LogUnshielding event raised by the PrivateBank contract.
type PrivateBankLogUnshielding struct {
	To         common.Address
	Value      uint64
	Commitment [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogUnshielding is a free log retrieval operation binding the contract event 0x44ed7e2197490271743d80ef2940b5ef22d886361e9634e9696842ccd2acfacf.
//
// Solidity: e LogUnshielding(to indexed address, value uint64, commitment bytes32)
func (_PrivateBank *PrivateBankFilterer) FilterLogUnshielding(opts *bind.FilterOpts, to []common.Address) (*PrivateBankLogUnshieldingIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PrivateBank.contract.FilterLogs(opts, "LogUnshielding", toRule)
	if err != nil {
		return nil, err
	}
	return &PrivateBankLogUnshieldingIterator{contract: _PrivateBank.contract, event: "LogUnshielding", logs: logs, sub: sub}, nil
}

// WatchLogUnshielding is a free log subscription operation binding the contract event 0x44ed7e2197490271743d80ef2940b5ef22d886361e9634e9696842ccd2acfacf.
//
// Solidity: e LogUnshielding(to indexed address, value uint64, commitment bytes32)
func (_PrivateBank *PrivateBankFilterer) WatchLogUnshielding(opts *bind.WatchOpts, sink chan<- *PrivateBankLogUnshielding, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PrivateBank.contract.WatchLogs(opts, "LogUnshielding", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivateBankLogUnshielding)
				if err := _PrivateBank.contract.UnpackLog(event, "LogUnshielding", log); err != nil {
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

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"}]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x6080604052348015600f57600080fd5b50604180601d6000396000f3006080604052348015600f57600080fd5b50600080fd00a165627a7a723058209e0f78b724550862f183c2a608847daa5856f2704d68b7fb41f8534d338f399d0029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// ZSLMerkleTreeABI is the input ABI used to generate the binding from.
const ZSLMerkleTreeABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"v\",\"type\":\"uint256\"},{\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"rightShift\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"left\",\"type\":\"bytes32\"},{\"name\":\"right\",\"type\":\"bytes32\"}],\"name\":\"combine\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"cm\",\"type\":\"bytes32\"}],\"name\":\"getWitness\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"available\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"v\",\"type\":\"uint256\"},{\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"leftShift\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"capacity\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"depth\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEmptyRoots\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"cm\",\"type\":\"bytes32\"}],\"name\":\"commitmentExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"size\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cm\",\"type\":\"bytes32\"}],\"name\":\"addCommitment\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"depth\",\"type\":\"uint256\"}],\"name\":\"getEmptyRoot\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"cm\",\"type\":\"bytes32\"}],\"name\":\"getLeafIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getCommitmentAtLeafIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"root\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"depth\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"}]"

// ZSLMerkleTreeBin is the compiled bytecode used for deploying new contracts.
const ZSLMerkleTreeBin = `0x6080604052600060045534801561001557600080fd5b5060405160208062000fc58339810160405251610030610302565b604051809103906000f08015801561004c573d6000803e3d6000fd5b5060008054600160a060020a031916600160a060020a03929092169190911790556001819055600281810a905561008b81640100000000610091810204565b50610313565b60038054600181018255600091825260008051602062000fa583398151915201819055805b60018303811015610102576100d48280640100000000610107810204565b600380546001818101835560009290925260008051602062000fa583398151915201829055909250016100b6565b505050565b6040805181815260608082018352600092909183916020820161080080388339019050509150600090505b60208110156101985784816020811061014757fe5b1a7f010000000000000000000000000000000000000000000000000000000000000002828281518110151561017857fe5b906020010190600160f860020a031916908160001a905350600101610132565b5060005b6020811015610205578381602081106101b157fe5b1a7f01000000000000000000000000000000000000000000000000000000000000000282826020018151811015156101e557fe5b906020010190600160f860020a031916908160001a90535060010161019c565b600080546040517f3b300c68000000000000000000000000000000000000000000000000000000008152602060048201818152865160248401528651600160a060020a0390941694633b300c68948894929384936044019290860191908190849084905b83811015610281578181015183820152602001610269565b50505050905090810190601f1680156102ae5780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b1580156102cd57600080fd5b505af11580156102e1573d6000803e3d6000fd5b505050506040513d60208110156102f757600080fd5b505195945050505050565b6040516103f98062000bac83390190565b61088980620003236000396000f3006080604052600436106100da5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630af0b13281146100ec5780631d48f5e91461011957806347e722461461013457806348a0d754146101ac5780635ae60499146101c15780635cfc1a51146101dc578063631c56ef146101f15780636def0cbe14610206578063863078651461026b578063949d225d146102975780639f7ad1d1146102ac578063a06f0977146102c6578063c8c50515146102de578063d0316c50146102f6578063ebf0c7171461030e575b3480156100e657600080fd5b50600080fd5b3480156100f857600080fd5b50610107600435602435610323565b60408051918252519081900360200190f35b34801561012557600080fd5b5061010760043560243561033b565b34801561014057600080fd5b5061014c600435610573565b604080518481529081018290526060602080830182815285519284019290925284516080840191868101910280838360005b8381101561019657818101518382015260200161017e565b5050505090500194505050505060405180910390f35b3480156101b857600080fd5b50610107610640565b3480156101cd57600080fd5b5061010760043560243561064a565b3480156101e857600080fd5b50610107610651565b3480156101fd57600080fd5b50610107610657565b34801561021257600080fd5b5061021b61065d565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561025757818101518382015260200161023f565b505050509050019250505060405180910390f35b34801561027757600080fd5b506102836004356106b6565b604080519115158252519081900360200190f35b3480156102a357600080fd5b506101076106ca565b3480156102b857600080fd5b506102c46004356106d0565b005b3480156102d257600080fd5b50610107600435610724565b3480156102ea57600080fd5b50610107600435610755565b34801561030257600080fd5b50610107600435610779565b34801561031a57600080fd5b506101076107a3565b60008160020a8381151561033357fe5b049392505050565b6040805181815260608082018352600092909183916020820161080080388339019050509150600090505b60208110156103e45784816020811061037b57fe5b1a7f01000000000000000000000000000000000000000000000000000000000000000282828151811015156103ac57fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600101610366565b5060005b6020811015610469578381602081106103fd57fe5b1a7f010000000000000000000000000000000000000000000000000000000000000002828260200181518110151561043157fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506001016103e8565b600080546040517f3b300c6800000000000000000000000000000000000000000000000000000000815260206004820181815286516024840152865173ffffffffffffffffffffffffffffffffffffffff90941694633b300c68948894929384936044019290860191908190849084905b838110156104f25781810151838201526020016104da565b50505050905090810190601f16801561051f5780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b15801561053e57600080fd5b505af1158015610552573d6000803e3d6000fd5b505050506040513d602081101561056857600080fd5b505195945050505050565b6000818152600560205260408120546060908290818381808080861161059857600080fd5b6001860394506001546040519080825280602002602001820160405280156105ca578160200160208202803883390190505b50935060009250849150600090505b600154831015610623576105f682600118848060010195506107b7565b845160018301928691811061060757fe5b6020908102909101015261061c826001610323565b91506105d9565b848461062d6107a3565b9850985098505050505050509193909250565b6004546002540390565b60020a0290565b60025490565b60015490565b606060038054806020026020016040519081016040528092919081815260200182805480156106ac57602002820191906000526020600020905b81548152600190910190602001808311610697575b5050505050905090565b600090815260056020526040902054151590565b60045490565b600081815260056020526040812054156106e957600080fd5b600254600454106106f957600080fd5b5060048054600101908190556000828152600560209081526040808320849055928252600690522055565b600354600090821061073557600080fd5b600380548390811061074357fe5b90600052602060002001549050919050565b60008181526005602052604081205481811161077057600080fd5b60010192915050565b6000806004548310151561078c57600080fd5b505060010160009081526006602052604090205490565b60006107b260006001546107b7565b905090565b6000806000806107c7868661064a565b600454116107ef5760038054869081106107dd57fe5b90600052602060002001549350610854565b841515610812576001860160008181526006602052604090205494509250610854565b61082961082087600161064a565b600187036107b7565b915061084561083987600161064a565b600101600187036107b7565b9050610851828261033b565b93505b505050929150505600a165627a7a72305820f7db65d632a6ab6572a7427e1b49b5fbe2f7810c1129ecb437692df4d32167b20029608060405234801561001057600080fd5b506103d9806100206000396000f3006080604052600436106100615763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633b300c688114610073578063794f69e3146100de578063dafca47a1461016e578063e088659e146101de575b34801561006d57600080fd5b50600080fd5b34801561007f57600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526100cc94369492936024939284019190819084018382808284375094975061024e9650505050505050565b60408051918252519081900360200190f35b3480156100ea57600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261015a94369492936024939284019190819084018382808284375094975050843595505050602083013592604081013592506060810135915060808101359060a08101359060c0013561028a565b604080519115158252519081900360200190f35b34801561017a57600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261015a943694929360249392840191908190840183828082843750949750508435955050506020830135926040013567ffffffffffffffff16915061030c9050565b3480156101ea57600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261015a943694929360249392840191908190840183828082843750949750508435955050506020830135926040013567ffffffffffffffff16915061038a9050565b805160009060401461025f57600080fd5b60405160208160808560006188015af1151561027a57600080fd5b8051602490910160405292915050565b600080604051368060048337602082828460006188025af115156102ad57600080fd5b8151910160405260001a7f01000000000000000000000000000000000000000000000000000000000000009081027fff0000000000000000000000000000000000000000000000000000000000000016149a9950505050505050505050565b600080604051368060048337602082828460006188045af1151561032f57600080fd5b8151910160405260001a7f01000000000000000000000000000000000000000000000000000000000000009081027fff0000000000000000000000000000000000000000000000000000000000000016149695505050505050565b600080604051368060048337602082828460006188035af1151561032f57600080fd00a165627a7a72305820923b26ed546e528b1a1407267c1697ecf6460025021dc366696fe95995e410600029c2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b`

// DeployZSLMerkleTree deploys a new Ethereum contract, binding an instance of ZSLMerkleTree to it.
func DeployZSLMerkleTree(auth *bind.TransactOpts, backend bind.ContractBackend, depth *big.Int) (common.Address, *types.Transaction, *ZSLMerkleTree, error) {
	parsed, err := abi.JSON(strings.NewReader(ZSLMerkleTreeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZSLMerkleTreeBin), backend, depth)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZSLMerkleTree{ZSLMerkleTreeCaller: ZSLMerkleTreeCaller{contract: contract}, ZSLMerkleTreeTransactor: ZSLMerkleTreeTransactor{contract: contract}, ZSLMerkleTreeFilterer: ZSLMerkleTreeFilterer{contract: contract}}, nil
}

// ZSLMerkleTree is an auto generated Go binding around an Ethereum contract.
type ZSLMerkleTree struct {
	ZSLMerkleTreeCaller     // Read-only binding to the contract
	ZSLMerkleTreeTransactor // Write-only binding to the contract
	ZSLMerkleTreeFilterer   // Log filterer for contract events
}

// ZSLMerkleTreeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZSLMerkleTreeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZSLMerkleTreeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZSLMerkleTreeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZSLMerkleTreeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZSLMerkleTreeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZSLMerkleTreeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZSLMerkleTreeSession struct {
	Contract     *ZSLMerkleTree    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZSLMerkleTreeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZSLMerkleTreeCallerSession struct {
	Contract *ZSLMerkleTreeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ZSLMerkleTreeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZSLMerkleTreeTransactorSession struct {
	Contract     *ZSLMerkleTreeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ZSLMerkleTreeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZSLMerkleTreeRaw struct {
	Contract *ZSLMerkleTree // Generic contract binding to access the raw methods on
}

// ZSLMerkleTreeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZSLMerkleTreeCallerRaw struct {
	Contract *ZSLMerkleTreeCaller // Generic read-only contract binding to access the raw methods on
}

// ZSLMerkleTreeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZSLMerkleTreeTransactorRaw struct {
	Contract *ZSLMerkleTreeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZSLMerkleTree creates a new instance of ZSLMerkleTree, bound to a specific deployed contract.
func NewZSLMerkleTree(address common.Address, backend bind.ContractBackend) (*ZSLMerkleTree, error) {
	contract, err := bindZSLMerkleTree(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZSLMerkleTree{ZSLMerkleTreeCaller: ZSLMerkleTreeCaller{contract: contract}, ZSLMerkleTreeTransactor: ZSLMerkleTreeTransactor{contract: contract}, ZSLMerkleTreeFilterer: ZSLMerkleTreeFilterer{contract: contract}}, nil
}

// NewZSLMerkleTreeCaller creates a new read-only instance of ZSLMerkleTree, bound to a specific deployed contract.
func NewZSLMerkleTreeCaller(address common.Address, caller bind.ContractCaller) (*ZSLMerkleTreeCaller, error) {
	contract, err := bindZSLMerkleTree(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZSLMerkleTreeCaller{contract: contract}, nil
}

// NewZSLMerkleTreeTransactor creates a new write-only instance of ZSLMerkleTree, bound to a specific deployed contract.
func NewZSLMerkleTreeTransactor(address common.Address, transactor bind.ContractTransactor) (*ZSLMerkleTreeTransactor, error) {
	contract, err := bindZSLMerkleTree(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZSLMerkleTreeTransactor{contract: contract}, nil
}

// NewZSLMerkleTreeFilterer creates a new log filterer instance of ZSLMerkleTree, bound to a specific deployed contract.
func NewZSLMerkleTreeFilterer(address common.Address, filterer bind.ContractFilterer) (*ZSLMerkleTreeFilterer, error) {
	contract, err := bindZSLMerkleTree(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZSLMerkleTreeFilterer{contract: contract}, nil
}

// bindZSLMerkleTree binds a generic wrapper to an already deployed contract.
func bindZSLMerkleTree(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZSLMerkleTreeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZSLMerkleTree *ZSLMerkleTreeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZSLMerkleTree.Contract.ZSLMerkleTreeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZSLMerkleTree *ZSLMerkleTreeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZSLMerkleTree.Contract.ZSLMerkleTreeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZSLMerkleTree *ZSLMerkleTreeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZSLMerkleTree.Contract.ZSLMerkleTreeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZSLMerkleTree *ZSLMerkleTreeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZSLMerkleTree.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZSLMerkleTree *ZSLMerkleTreeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZSLMerkleTree.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZSLMerkleTree *ZSLMerkleTreeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZSLMerkleTree.Contract.contract.Transact(opts, method, params...)
}

// Available is a free data retrieval call binding the contract method 0x48a0d754.
//
// Solidity: function available() constant returns(uint256)
func (_ZSLMerkleTree *ZSLMerkleTreeCaller) Available(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZSLMerkleTree.contract.Call(opts, out, "available")
	return *ret0, err
}

// Available is a free data retrieval call binding the contract method 0x48a0d754.
//
// Solidity: function available() constant returns(uint256)
func (_ZSLMerkleTree *ZSLMerkleTreeSession) Available() (*big.Int, error) {
	return _ZSLMerkleTree.Contract.Available(&_ZSLMerkleTree.CallOpts)
}

// Available is a free data retrieval call binding the contract method 0x48a0d754.
//
// Solidity: function available() constant returns(uint256)
func (_ZSLMerkleTree *ZSLMerkleTreeCallerSession) Available() (*big.Int, error) {
	return _ZSLMerkleTree.Contract.Available(&_ZSLMerkleTree.CallOpts)
}

// Capacity is a free data retrieval call binding the contract method 0x5cfc1a51.
//
// Solidity: function capacity() constant returns(uint256)
func (_ZSLMerkleTree *ZSLMerkleTreeCaller) Capacity(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZSLMerkleTree.contract.Call(opts, out, "capacity")
	return *ret0, err
}

// Capacity is a free data retrieval call binding the contract method 0x5cfc1a51.
//
// Solidity: function capacity() constant returns(uint256)
func (_ZSLMerkleTree *ZSLMerkleTreeSession) Capacity() (*big.Int, error) {
	return _ZSLMerkleTree.Contract.Capacity(&_ZSLMerkleTree.CallOpts)
}

// Capacity is a free data retrieval call binding the contract method 0x5cfc1a51.
//
// Solidity: function capacity() constant returns(uint256)
func (_ZSLMerkleTree *ZSLMerkleTreeCallerSession) Capacity() (*big.Int, error) {
	return _ZSLMerkleTree.Contract.Capacity(&_ZSLMerkleTree.CallOpts)
}

// Combine is a free data retrieval call binding the contract method 0x1d48f5e9.
//
// Solidity: function combine(left bytes32, right bytes32) constant returns(bytes32)
func (_ZSLMerkleTree *ZSLMerkleTreeCaller) Combine(opts *bind.CallOpts, left [32]byte, right [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ZSLMerkleTree.contract.Call(opts, out, "combine", left, right)
	return *ret0, err
}

// Combine is a free data retrieval call binding the contract method 0x1d48f5e9.
//
// Solidity: function combine(left bytes32, right bytes32) constant returns(bytes32)
func (_ZSLMerkleTree *ZSLMerkleTreeSession) Combine(left [32]byte, right [32]byte) ([32]byte, error) {
	return _ZSLMerkleTree.Contract.Combine(&_ZSLMerkleTree.CallOpts, left, right)
}

// Combine is a free data retrieval call binding the contract method 0x1d48f5e9.
//
// Solidity: function combine(left bytes32, right bytes32) constant returns(bytes32)
func (_ZSLMerkleTree *ZSLMerkleTreeCallerSession) Combine(left [32]byte, right [32]byte) ([32]byte, error) {
	return _ZSLMerkleTree.Contract.Combine(&_ZSLMerkleTree.CallOpts, left, right)
}

// CommitmentExists is a free data retrieval call binding the contract method 0x86307865.
//
// Solidity: function commitmentExists(cm bytes32) constant returns(bool)
func (_ZSLMerkleTree *ZSLMerkleTreeCaller) CommitmentExists(opts *bind.CallOpts, cm [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ZSLMerkleTree.contract.Call(opts, out, "commitmentExists", cm)
	return *ret0, err
}

// CommitmentExists is a free data retrieval call binding the contract method 0x86307865.
//
// Solidity: function commitmentExists(cm bytes32) constant returns(bool)
func (_ZSLMerkleTree *ZSLMerkleTreeSession) CommitmentExists(cm [32]byte) (bool, error) {
	return _ZSLMerkleTree.Contract.CommitmentExists(&_ZSLMerkleTree.CallOpts, cm)
}

// CommitmentExists is a free data retrieval call binding the contract method 0x86307865.
//
// Solidity: function commitmentExists(cm bytes32) constant returns(bool)
func (_ZSLMerkleTree *ZSLMerkleTreeCallerSession) CommitmentExists(cm [32]byte) (bool, error) {
	return _ZSLMerkleTree.Contract.CommitmentExists(&_ZSLMerkleTree.CallOpts, cm)
}

// Depth is a free data retrieval call binding the contract method 0x631c56ef.
//
// Solidity: function depth() constant returns(uint256)
func (_ZSLMerkleTree *ZSLMerkleTreeCaller) Depth(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZSLMerkleTree.contract.Call(opts, out, "depth")
	return *ret0, err
}

// Depth is a free data retrieval call binding the contract method 0x631c56ef.
//
// Solidity: function depth() constant returns(uint256)
func (_ZSLMerkleTree *ZSLMerkleTreeSession) Depth() (*big.Int, error) {
	return _ZSLMerkleTree.Contract.Depth(&_ZSLMerkleTree.CallOpts)
}

// Depth is a free data retrieval call binding the contract method 0x631c56ef.
//
// Solidity: function depth() constant returns(uint256)
func (_ZSLMerkleTree *ZSLMerkleTreeCallerSession) Depth() (*big.Int, error) {
	return _ZSLMerkleTree.Contract.Depth(&_ZSLMerkleTree.CallOpts)
}

// GetCommitmentAtLeafIndex is a free data retrieval call binding the contract method 0xd0316c50.
//
// Solidity: function getCommitmentAtLeafIndex(index uint256) constant returns(bytes32)
func (_ZSLMerkleTree *ZSLMerkleTreeCaller) GetCommitmentAtLeafIndex(opts *bind.CallOpts, index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ZSLMerkleTree.contract.Call(opts, out, "getCommitmentAtLeafIndex", index)
	return *ret0, err
}

// GetCommitmentAtLeafIndex is a free data retrieval call binding the contract method 0xd0316c50.
//
// Solidity: function getCommitmentAtLeafIndex(index uint256) constant returns(bytes32)
func (_ZSLMerkleTree *ZSLMerkleTreeSession) GetCommitmentAtLeafIndex(index *big.Int) ([32]byte, error) {
	return _ZSLMerkleTree.Contract.GetCommitmentAtLeafIndex(&_ZSLMerkleTree.CallOpts, index)
}

// GetCommitmentAtLeafIndex is a free data retrieval call binding the contract method 0xd0316c50.
//
// Solidity: function getCommitmentAtLeafIndex(index uint256) constant returns(bytes32)
func (_ZSLMerkleTree *ZSLMerkleTreeCallerSession) GetCommitmentAtLeafIndex(index *big.Int) ([32]byte, error) {
	return _ZSLMerkleTree.Contract.GetCommitmentAtLeafIndex(&_ZSLMerkleTree.CallOpts, index)
}

// GetEmptyRoot is a free data retrieval call binding the contract method 0xa06f0977.
//
// Solidity: function getEmptyRoot(depth uint256) constant returns(bytes32)
func (_ZSLMerkleTree *ZSLMerkleTreeCaller) GetEmptyRoot(opts *bind.CallOpts, depth *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ZSLMerkleTree.contract.Call(opts, out, "getEmptyRoot", depth)
	return *ret0, err
}

// GetEmptyRoot is a free data retrieval call binding the contract method 0xa06f0977.
//
// Solidity: function getEmptyRoot(depth uint256) constant returns(bytes32)
func (_ZSLMerkleTree *ZSLMerkleTreeSession) GetEmptyRoot(depth *big.Int) ([32]byte, error) {
	return _ZSLMerkleTree.Contract.GetEmptyRoot(&_ZSLMerkleTree.CallOpts, depth)
}

// GetEmptyRoot is a free data retrieval call binding the contract method 0xa06f0977.
//
// Solidity: function getEmptyRoot(depth uint256) constant returns(bytes32)
func (_ZSLMerkleTree *ZSLMerkleTreeCallerSession) GetEmptyRoot(depth *big.Int) ([32]byte, error) {
	return _ZSLMerkleTree.Contract.GetEmptyRoot(&_ZSLMerkleTree.CallOpts, depth)
}

// GetEmptyRoots is a free data retrieval call binding the contract method 0x6def0cbe.
//
// Solidity: function getEmptyRoots() constant returns(bytes32[])
func (_ZSLMerkleTree *ZSLMerkleTreeCaller) GetEmptyRoots(opts *bind.CallOpts) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _ZSLMerkleTree.contract.Call(opts, out, "getEmptyRoots")
	return *ret0, err
}

// GetEmptyRoots is a free data retrieval call binding the contract method 0x6def0cbe.
//
// Solidity: function getEmptyRoots() constant returns(bytes32[])
func (_ZSLMerkleTree *ZSLMerkleTreeSession) GetEmptyRoots() ([][32]byte, error) {
	return _ZSLMerkleTree.Contract.GetEmptyRoots(&_ZSLMerkleTree.CallOpts)
}

// GetEmptyRoots is a free data retrieval call binding the contract method 0x6def0cbe.
//
// Solidity: function getEmptyRoots() constant returns(bytes32[])
func (_ZSLMerkleTree *ZSLMerkleTreeCallerSession) GetEmptyRoots() ([][32]byte, error) {
	return _ZSLMerkleTree.Contract.GetEmptyRoots(&_ZSLMerkleTree.CallOpts)
}

// GetLeafIndex is a free data retrieval call binding the contract method 0xc8c50515.
//
// Solidity: function getLeafIndex(cm bytes32) constant returns(uint256)
func (_ZSLMerkleTree *ZSLMerkleTreeCaller) GetLeafIndex(opts *bind.CallOpts, cm [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZSLMerkleTree.contract.Call(opts, out, "getLeafIndex", cm)
	return *ret0, err
}

// GetLeafIndex is a free data retrieval call binding the contract method 0xc8c50515.
//
// Solidity: function getLeafIndex(cm bytes32) constant returns(uint256)
func (_ZSLMerkleTree *ZSLMerkleTreeSession) GetLeafIndex(cm [32]byte) (*big.Int, error) {
	return _ZSLMerkleTree.Contract.GetLeafIndex(&_ZSLMerkleTree.CallOpts, cm)
}

// GetLeafIndex is a free data retrieval call binding the contract method 0xc8c50515.
//
// Solidity: function getLeafIndex(cm bytes32) constant returns(uint256)
func (_ZSLMerkleTree *ZSLMerkleTreeCallerSession) GetLeafIndex(cm [32]byte) (*big.Int, error) {
	return _ZSLMerkleTree.Contract.GetLeafIndex(&_ZSLMerkleTree.CallOpts, cm)
}

// GetWitness is a free data retrieval call binding the contract method 0x47e72246.
//
// Solidity: function getWitness(cm bytes32) constant returns(uint256, bytes32[], bytes32)
func (_ZSLMerkleTree *ZSLMerkleTreeCaller) GetWitness(opts *bind.CallOpts, cm [32]byte) (*big.Int, [][32]byte, [32]byte, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new([][32]byte)
		ret2 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _ZSLMerkleTree.contract.Call(opts, out, "getWitness", cm)
	return *ret0, *ret1, *ret2, err
}

// GetWitness is a free data retrieval call binding the contract method 0x47e72246.
//
// Solidity: function getWitness(cm bytes32) constant returns(uint256, bytes32[], bytes32)
func (_ZSLMerkleTree *ZSLMerkleTreeSession) GetWitness(cm [32]byte) (*big.Int, [][32]byte, [32]byte, error) {
	return _ZSLMerkleTree.Contract.GetWitness(&_ZSLMerkleTree.CallOpts, cm)
}

// GetWitness is a free data retrieval call binding the contract method 0x47e72246.
//
// Solidity: function getWitness(cm bytes32) constant returns(uint256, bytes32[], bytes32)
func (_ZSLMerkleTree *ZSLMerkleTreeCallerSession) GetWitness(cm [32]byte) (*big.Int, [][32]byte, [32]byte, error) {
	return _ZSLMerkleTree.Contract.GetWitness(&_ZSLMerkleTree.CallOpts, cm)
}

// LeftShift is a free data retrieval call binding the contract method 0x5ae60499.
//
// Solidity: function leftShift(v uint256, n uint256) constant returns(uint256)
func (_ZSLMerkleTree *ZSLMerkleTreeCaller) LeftShift(opts *bind.CallOpts, v *big.Int, n *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZSLMerkleTree.contract.Call(opts, out, "leftShift", v, n)
	return *ret0, err
}

// LeftShift is a free data retrieval call binding the contract method 0x5ae60499.
//
// Solidity: function leftShift(v uint256, n uint256) constant returns(uint256)
func (_ZSLMerkleTree *ZSLMerkleTreeSession) LeftShift(v *big.Int, n *big.Int) (*big.Int, error) {
	return _ZSLMerkleTree.Contract.LeftShift(&_ZSLMerkleTree.CallOpts, v, n)
}

// LeftShift is a free data retrieval call binding the contract method 0x5ae60499.
//
// Solidity: function leftShift(v uint256, n uint256) constant returns(uint256)
func (_ZSLMerkleTree *ZSLMerkleTreeCallerSession) LeftShift(v *big.Int, n *big.Int) (*big.Int, error) {
	return _ZSLMerkleTree.Contract.LeftShift(&_ZSLMerkleTree.CallOpts, v, n)
}

// RightShift is a free data retrieval call binding the contract method 0x0af0b132.
//
// Solidity: function rightShift(v uint256, n uint256) constant returns(uint256)
func (_ZSLMerkleTree *ZSLMerkleTreeCaller) RightShift(opts *bind.CallOpts, v *big.Int, n *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZSLMerkleTree.contract.Call(opts, out, "rightShift", v, n)
	return *ret0, err
}

// RightShift is a free data retrieval call binding the contract method 0x0af0b132.
//
// Solidity: function rightShift(v uint256, n uint256) constant returns(uint256)
func (_ZSLMerkleTree *ZSLMerkleTreeSession) RightShift(v *big.Int, n *big.Int) (*big.Int, error) {
	return _ZSLMerkleTree.Contract.RightShift(&_ZSLMerkleTree.CallOpts, v, n)
}

// RightShift is a free data retrieval call binding the contract method 0x0af0b132.
//
// Solidity: function rightShift(v uint256, n uint256) constant returns(uint256)
func (_ZSLMerkleTree *ZSLMerkleTreeCallerSession) RightShift(v *big.Int, n *big.Int) (*big.Int, error) {
	return _ZSLMerkleTree.Contract.RightShift(&_ZSLMerkleTree.CallOpts, v, n)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() constant returns(bytes32)
func (_ZSLMerkleTree *ZSLMerkleTreeCaller) Root(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ZSLMerkleTree.contract.Call(opts, out, "root")
	return *ret0, err
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() constant returns(bytes32)
func (_ZSLMerkleTree *ZSLMerkleTreeSession) Root() ([32]byte, error) {
	return _ZSLMerkleTree.Contract.Root(&_ZSLMerkleTree.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() constant returns(bytes32)
func (_ZSLMerkleTree *ZSLMerkleTreeCallerSession) Root() ([32]byte, error) {
	return _ZSLMerkleTree.Contract.Root(&_ZSLMerkleTree.CallOpts)
}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() constant returns(uint256)
func (_ZSLMerkleTree *ZSLMerkleTreeCaller) Size(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZSLMerkleTree.contract.Call(opts, out, "size")
	return *ret0, err
}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() constant returns(uint256)
func (_ZSLMerkleTree *ZSLMerkleTreeSession) Size() (*big.Int, error) {
	return _ZSLMerkleTree.Contract.Size(&_ZSLMerkleTree.CallOpts)
}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() constant returns(uint256)
func (_ZSLMerkleTree *ZSLMerkleTreeCallerSession) Size() (*big.Int, error) {
	return _ZSLMerkleTree.Contract.Size(&_ZSLMerkleTree.CallOpts)
}

// AddCommitment is a paid mutator transaction binding the contract method 0x9f7ad1d1.
//
// Solidity: function addCommitment(cm bytes32) returns()
func (_ZSLMerkleTree *ZSLMerkleTreeTransactor) AddCommitment(opts *bind.TransactOpts, cm [32]byte) (*types.Transaction, error) {
	return _ZSLMerkleTree.contract.Transact(opts, "addCommitment", cm)
}

// AddCommitment is a paid mutator transaction binding the contract method 0x9f7ad1d1.
//
// Solidity: function addCommitment(cm bytes32) returns()
func (_ZSLMerkleTree *ZSLMerkleTreeSession) AddCommitment(cm [32]byte) (*types.Transaction, error) {
	return _ZSLMerkleTree.Contract.AddCommitment(&_ZSLMerkleTree.TransactOpts, cm)
}

// AddCommitment is a paid mutator transaction binding the contract method 0x9f7ad1d1.
//
// Solidity: function addCommitment(cm bytes32) returns()
func (_ZSLMerkleTree *ZSLMerkleTreeTransactorSession) AddCommitment(cm [32]byte) (*types.Transaction, error) {
	return _ZSLMerkleTree.Contract.AddCommitment(&_ZSLMerkleTree.TransactOpts, cm)
}

// ZSLPrecompileABI is the input ABI used to generate the binding from.
const ZSLPrecompileABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"sha256Compress\",\"outputs\":[{\"name\":\"result\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\"},{\"name\":\"treeRoot\",\"type\":\"bytes32\"},{\"name\":\"spendNullifier1\",\"type\":\"bytes32\"},{\"name\":\"spendNullifier2\",\"type\":\"bytes32\"},{\"name\":\"sendNullifier1\",\"type\":\"bytes32\"},{\"name\":\"sendNullifier2\",\"type\":\"bytes32\"},{\"name\":\"commitment1\",\"type\":\"bytes32\"},{\"name\":\"commitment2\",\"type\":\"bytes32\"}],\"name\":\"verifyShieldedTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\"},{\"name\":\"spendNullifier\",\"type\":\"bytes32\"},{\"name\":\"treeRoot\",\"type\":\"bytes32\"},{\"name\":\"value\",\"type\":\"uint64\"}],\"name\":\"verifyUnshielding\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\"},{\"name\":\"sendNullifier\",\"type\":\"bytes32\"},{\"name\":\"commitment\",\"type\":\"bytes32\"},{\"name\":\"value\",\"type\":\"uint64\"}],\"name\":\"verifyShielding\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"}]"

// ZSLPrecompileBin is the compiled bytecode used for deploying new contracts.
const ZSLPrecompileBin = `0x608060405234801561001057600080fd5b506103d9806100206000396000f3006080604052600436106100615763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633b300c688114610073578063794f69e3146100de578063dafca47a1461016e578063e088659e146101de575b34801561006d57600080fd5b50600080fd5b34801561007f57600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526100cc94369492936024939284019190819084018382808284375094975061024e9650505050505050565b60408051918252519081900360200190f35b3480156100ea57600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261015a94369492936024939284019190819084018382808284375094975050843595505050602083013592604081013592506060810135915060808101359060a08101359060c0013561028a565b604080519115158252519081900360200190f35b34801561017a57600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261015a943694929360249392840191908190840183828082843750949750508435955050506020830135926040013567ffffffffffffffff16915061030c9050565b3480156101ea57600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261015a943694929360249392840191908190840183828082843750949750508435955050506020830135926040013567ffffffffffffffff16915061038a9050565b805160009060401461025f57600080fd5b60405160208160808560006188015af1151561027a57600080fd5b8051602490910160405292915050565b600080604051368060048337602082828460006188025af115156102ad57600080fd5b8151910160405260001a7f01000000000000000000000000000000000000000000000000000000000000009081027fff0000000000000000000000000000000000000000000000000000000000000016149a9950505050505050505050565b600080604051368060048337602082828460006188045af1151561032f57600080fd5b8151910160405260001a7f01000000000000000000000000000000000000000000000000000000000000009081027fff0000000000000000000000000000000000000000000000000000000000000016149695505050505050565b600080604051368060048337602082828460006188035af1151561032f57600080fd00a165627a7a72305820923b26ed546e528b1a1407267c1697ecf6460025021dc366696fe95995e410600029`

// DeployZSLPrecompile deploys a new Ethereum contract, binding an instance of ZSLPrecompile to it.
func DeployZSLPrecompile(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZSLPrecompile, error) {
	parsed, err := abi.JSON(strings.NewReader(ZSLPrecompileABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZSLPrecompileBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZSLPrecompile{ZSLPrecompileCaller: ZSLPrecompileCaller{contract: contract}, ZSLPrecompileTransactor: ZSLPrecompileTransactor{contract: contract}, ZSLPrecompileFilterer: ZSLPrecompileFilterer{contract: contract}}, nil
}

// ZSLPrecompile is an auto generated Go binding around an Ethereum contract.
type ZSLPrecompile struct {
	ZSLPrecompileCaller     // Read-only binding to the contract
	ZSLPrecompileTransactor // Write-only binding to the contract
	ZSLPrecompileFilterer   // Log filterer for contract events
}

// ZSLPrecompileCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZSLPrecompileCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZSLPrecompileTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZSLPrecompileTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZSLPrecompileFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZSLPrecompileFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZSLPrecompileSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZSLPrecompileSession struct {
	Contract     *ZSLPrecompile    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZSLPrecompileCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZSLPrecompileCallerSession struct {
	Contract *ZSLPrecompileCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ZSLPrecompileTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZSLPrecompileTransactorSession struct {
	Contract     *ZSLPrecompileTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ZSLPrecompileRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZSLPrecompileRaw struct {
	Contract *ZSLPrecompile // Generic contract binding to access the raw methods on
}

// ZSLPrecompileCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZSLPrecompileCallerRaw struct {
	Contract *ZSLPrecompileCaller // Generic read-only contract binding to access the raw methods on
}

// ZSLPrecompileTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZSLPrecompileTransactorRaw struct {
	Contract *ZSLPrecompileTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZSLPrecompile creates a new instance of ZSLPrecompile, bound to a specific deployed contract.
func NewZSLPrecompile(address common.Address, backend bind.ContractBackend) (*ZSLPrecompile, error) {
	contract, err := bindZSLPrecompile(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZSLPrecompile{ZSLPrecompileCaller: ZSLPrecompileCaller{contract: contract}, ZSLPrecompileTransactor: ZSLPrecompileTransactor{contract: contract}, ZSLPrecompileFilterer: ZSLPrecompileFilterer{contract: contract}}, nil
}

// NewZSLPrecompileCaller creates a new read-only instance of ZSLPrecompile, bound to a specific deployed contract.
func NewZSLPrecompileCaller(address common.Address, caller bind.ContractCaller) (*ZSLPrecompileCaller, error) {
	contract, err := bindZSLPrecompile(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZSLPrecompileCaller{contract: contract}, nil
}

// NewZSLPrecompileTransactor creates a new write-only instance of ZSLPrecompile, bound to a specific deployed contract.
func NewZSLPrecompileTransactor(address common.Address, transactor bind.ContractTransactor) (*ZSLPrecompileTransactor, error) {
	contract, err := bindZSLPrecompile(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZSLPrecompileTransactor{contract: contract}, nil
}

// NewZSLPrecompileFilterer creates a new log filterer instance of ZSLPrecompile, bound to a specific deployed contract.
func NewZSLPrecompileFilterer(address common.Address, filterer bind.ContractFilterer) (*ZSLPrecompileFilterer, error) {
	contract, err := bindZSLPrecompile(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZSLPrecompileFilterer{contract: contract}, nil
}

// bindZSLPrecompile binds a generic wrapper to an already deployed contract.
func bindZSLPrecompile(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZSLPrecompileABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZSLPrecompile *ZSLPrecompileRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZSLPrecompile.Contract.ZSLPrecompileCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZSLPrecompile *ZSLPrecompileRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZSLPrecompile.Contract.ZSLPrecompileTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZSLPrecompile *ZSLPrecompileRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZSLPrecompile.Contract.ZSLPrecompileTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZSLPrecompile *ZSLPrecompileCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZSLPrecompile.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZSLPrecompile *ZSLPrecompileTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZSLPrecompile.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZSLPrecompile *ZSLPrecompileTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZSLPrecompile.Contract.contract.Transact(opts, method, params...)
}

// Sha256Compress is a free data retrieval call binding the contract method 0x3b300c68.
//
// Solidity: function sha256Compress(input bytes) constant returns(result bytes32)
func (_ZSLPrecompile *ZSLPrecompileCaller) Sha256Compress(opts *bind.CallOpts, input []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ZSLPrecompile.contract.Call(opts, out, "sha256Compress", input)
	return *ret0, err
}

// Sha256Compress is a free data retrieval call binding the contract method 0x3b300c68.
//
// Solidity: function sha256Compress(input bytes) constant returns(result bytes32)
func (_ZSLPrecompile *ZSLPrecompileSession) Sha256Compress(input []byte) ([32]byte, error) {
	return _ZSLPrecompile.Contract.Sha256Compress(&_ZSLPrecompile.CallOpts, input)
}

// Sha256Compress is a free data retrieval call binding the contract method 0x3b300c68.
//
// Solidity: function sha256Compress(input bytes) constant returns(result bytes32)
func (_ZSLPrecompile *ZSLPrecompileCallerSession) Sha256Compress(input []byte) ([32]byte, error) {
	return _ZSLPrecompile.Contract.Sha256Compress(&_ZSLPrecompile.CallOpts, input)
}

// VerifyShieldedTransfer is a free data retrieval call binding the contract method 0x794f69e3.
//
// Solidity: function verifyShieldedTransfer(proof bytes, treeRoot bytes32, spendNullifier1 bytes32, spendNullifier2 bytes32, sendNullifier1 bytes32, sendNullifier2 bytes32, commitment1 bytes32, commitment2 bytes32) constant returns(bool)
func (_ZSLPrecompile *ZSLPrecompileCaller) VerifyShieldedTransfer(opts *bind.CallOpts, proof []byte, treeRoot [32]byte, spendNullifier1 [32]byte, spendNullifier2 [32]byte, sendNullifier1 [32]byte, sendNullifier2 [32]byte, commitment1 [32]byte, commitment2 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ZSLPrecompile.contract.Call(opts, out, "verifyShieldedTransfer", proof, treeRoot, spendNullifier1, spendNullifier2, sendNullifier1, sendNullifier2, commitment1, commitment2)
	return *ret0, err
}

// VerifyShieldedTransfer is a free data retrieval call binding the contract method 0x794f69e3.
//
// Solidity: function verifyShieldedTransfer(proof bytes, treeRoot bytes32, spendNullifier1 bytes32, spendNullifier2 bytes32, sendNullifier1 bytes32, sendNullifier2 bytes32, commitment1 bytes32, commitment2 bytes32) constant returns(bool)
func (_ZSLPrecompile *ZSLPrecompileSession) VerifyShieldedTransfer(proof []byte, treeRoot [32]byte, spendNullifier1 [32]byte, spendNullifier2 [32]byte, sendNullifier1 [32]byte, sendNullifier2 [32]byte, commitment1 [32]byte, commitment2 [32]byte) (bool, error) {
	return _ZSLPrecompile.Contract.VerifyShieldedTransfer(&_ZSLPrecompile.CallOpts, proof, treeRoot, spendNullifier1, spendNullifier2, sendNullifier1, sendNullifier2, commitment1, commitment2)
}

// VerifyShieldedTransfer is a free data retrieval call binding the contract method 0x794f69e3.
//
// Solidity: function verifyShieldedTransfer(proof bytes, treeRoot bytes32, spendNullifier1 bytes32, spendNullifier2 bytes32, sendNullifier1 bytes32, sendNullifier2 bytes32, commitment1 bytes32, commitment2 bytes32) constant returns(bool)
func (_ZSLPrecompile *ZSLPrecompileCallerSession) VerifyShieldedTransfer(proof []byte, treeRoot [32]byte, spendNullifier1 [32]byte, spendNullifier2 [32]byte, sendNullifier1 [32]byte, sendNullifier2 [32]byte, commitment1 [32]byte, commitment2 [32]byte) (bool, error) {
	return _ZSLPrecompile.Contract.VerifyShieldedTransfer(&_ZSLPrecompile.CallOpts, proof, treeRoot, spendNullifier1, spendNullifier2, sendNullifier1, sendNullifier2, commitment1, commitment2)
}

// VerifyShielding is a free data retrieval call binding the contract method 0xe088659e.
//
// Solidity: function verifyShielding(proof bytes, sendNullifier bytes32, commitment bytes32, value uint64) constant returns(bool)
func (_ZSLPrecompile *ZSLPrecompileCaller) VerifyShielding(opts *bind.CallOpts, proof []byte, sendNullifier [32]byte, commitment [32]byte, value uint64) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ZSLPrecompile.contract.Call(opts, out, "verifyShielding", proof, sendNullifier, commitment, value)
	return *ret0, err
}

// VerifyShielding is a free data retrieval call binding the contract method 0xe088659e.
//
// Solidity: function verifyShielding(proof bytes, sendNullifier bytes32, commitment bytes32, value uint64) constant returns(bool)
func (_ZSLPrecompile *ZSLPrecompileSession) VerifyShielding(proof []byte, sendNullifier [32]byte, commitment [32]byte, value uint64) (bool, error) {
	return _ZSLPrecompile.Contract.VerifyShielding(&_ZSLPrecompile.CallOpts, proof, sendNullifier, commitment, value)
}

// VerifyShielding is a free data retrieval call binding the contract method 0xe088659e.
//
// Solidity: function verifyShielding(proof bytes, sendNullifier bytes32, commitment bytes32, value uint64) constant returns(bool)
func (_ZSLPrecompile *ZSLPrecompileCallerSession) VerifyShielding(proof []byte, sendNullifier [32]byte, commitment [32]byte, value uint64) (bool, error) {
	return _ZSLPrecompile.Contract.VerifyShielding(&_ZSLPrecompile.CallOpts, proof, sendNullifier, commitment, value)
}

// VerifyUnshielding is a free data retrieval call binding the contract method 0xdafca47a.
//
// Solidity: function verifyUnshielding(proof bytes, spendNullifier bytes32, treeRoot bytes32, value uint64) constant returns(bool)
func (_ZSLPrecompile *ZSLPrecompileCaller) VerifyUnshielding(opts *bind.CallOpts, proof []byte, spendNullifier [32]byte, treeRoot [32]byte, value uint64) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ZSLPrecompile.contract.Call(opts, out, "verifyUnshielding", proof, spendNullifier, treeRoot, value)
	return *ret0, err
}

// VerifyUnshielding is a free data retrieval call binding the contract method 0xdafca47a.
//
// Solidity: function verifyUnshielding(proof bytes, spendNullifier bytes32, treeRoot bytes32, value uint64) constant returns(bool)
func (_ZSLPrecompile *ZSLPrecompileSession) VerifyUnshielding(proof []byte, spendNullifier [32]byte, treeRoot [32]byte, value uint64) (bool, error) {
	return _ZSLPrecompile.Contract.VerifyUnshielding(&_ZSLPrecompile.CallOpts, proof, spendNullifier, treeRoot, value)
}

// VerifyUnshielding is a free data retrieval call binding the contract method 0xdafca47a.
//
// Solidity: function verifyUnshielding(proof bytes, spendNullifier bytes32, treeRoot bytes32, value uint64) constant returns(bool)
func (_ZSLPrecompile *ZSLPrecompileCallerSession) VerifyUnshielding(proof []byte, spendNullifier [32]byte, treeRoot [32]byte, value uint64) (bool, error) {
	return _ZSLPrecompile.Contract.VerifyUnshielding(&_ZSLPrecompile.CallOpts, proof, spendNullifier, treeRoot, value)
}
