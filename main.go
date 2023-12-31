package main

import (
	"fmt"
	"oracle-test/plugins"
	"plugin"
)

func main() {
	var res interface{}
	var err error

	var pm = plugins.PluginManager{}
	pm.InitializePlugins()

	res, _ = pm.CallMethod("Add_Numbers", uint64(123), uint64(321))
	fmt.Println(res)

	return

	// mathapi api
	mathapi, err := loadPlugin("build/plugins/mathapi.so")
	if err != nil {
		fmt.Println(err)
		return
	}

	mathapi.Initialize()
	fmt.Println(plugins.DefaultGetMethods(mathapi))
	res, err = plugins.DefaultCallMethod(mathapi, "Add_Numbers", uint64(123), uint64(321))
	fmt.Println("Add_Numbers:", res.(uint64))

	// combined api
	exchange, err := loadPlugin("build/plugins/combinedexhcnagerate.so")
	if err != nil {
		fmt.Println(err)
		return
	}

	exchange.Initialize()
	fmt.Println(plugins.DefaultGetMethods(exchange))
	// res, err = exchange.CallMethod("Exhange_Rate", "USD", "EUR")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("Exhange_Rate:", res)

	// livescore api
	sportsdb, err := loadPlugin("build/plugins/livescore.so")
	if err != nil {
		fmt.Println(err)
		return
	}

	sportsdb.Initialize()
	fmt.Println(plugins.DefaultGetMethods(sportsdb))
	// res, err = sportsdb.CallMethod("Get_All_Leagues")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("Get_All_Leagues:", res)

	// bet365 api
	bet365, err := loadPlugin("build/plugins/bet365.so")
	if err != nil {
		fmt.Println(err)
		return
	}

	bet365.Initialize()
	fmt.Println(plugins.DefaultGetMethods(bet365))
	// res, err = bet365.CallMethod("Get_Premier_League_Quotas")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("Get_Premiere_League_Quotas:", res)

	// goerli
	goerli, err := loadPlugin("build/plugins/goerli.so")
	if err != nil {
		fmt.Println(err)
		return
	}

	goerli.Initialize()
	fmt.Println(plugins.DefaultGetMethods(goerli))
	// res, err = goerli.CallMethod("Eth_blockNumber")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("Eth_blockNumber: ", res)

	// blockNumberBytes := big.NewInt(1000000).Bytes()
	// res, err = goerli.CallMethod("Eth_getBlockByNumber", blockNumberBytes)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("Eth_getBlockByNumber: ", res)
}

func loadPlugin(path string) (plugins.IPlugin, error) {
	p, err := plugin.Open(path)
	if err != nil {
		return nil, err
	}

	sym, err := p.Lookup("ExportPlugin")
	if err != nil {
		return nil, err
	}

	return sym.(plugins.IPlugin), nil
}
