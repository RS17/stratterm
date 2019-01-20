package investmentstrategy

import (
	"../sqlcmdr"
	"../sthelper"
)

type investment struct{
	symbol string 
	description string
	startprice float64
	endprice float64
}

func Handle(args []string ) string{
	switch args[0]	{
		case "create" : 
			return create( args[1:] )
		case "show" :
			return show( args[1:] )
		default :
			return "ERROR: invalid argument"
	}
}

func create( args []string ) string{
	if( len(args) < 2) {return "ERROR: must declare investment and strategy"}
	
	icmd := sqlcmdr.InsertCmd{ Tablename: "investmentstrategy" }
	icmd.Add( "investment", args[0] )
	icmd.Add( "strategy", args[1] )
	
	sthelper.SimpleInsert( icmd )
		
	return "created investmentstrategy"
}

func show( args []string ) string{
	scmd := sqlcmdr.SelectCmd{ Tablename: "investmentstrategy" }
	result := sthelper.ExecShow( args, scmd, "investmentstrategy.csv" );
	return result
}

func Createtable(){ 
	sqlcmdr.JustRunIt( 	"CREATE TABLE IF NOT EXISTS investmentstrategy (" +
						"investment INT, " +
						"strategy VARCHAR(50) )"  )
}

