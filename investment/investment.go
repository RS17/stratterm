package investment

import (
	"../sqlcmdr"
	"../sthelper"
)

type investment struct{
	symbol string 
	description string
	profit float64
	notes string
}

func Handle(args []string ) string{
	switch args[0]	{
		case "create" : 
			return create( sthelper.SliceShift( args ) )
		case "show" :
			return show( sthelper.SliceShift( args ) )
 		default :
			return "ERROR: invalid argument"
	}
}

func create( args []string ) string{
	// creates row entry in investment table
	if( len(args) < 1) {return "ERROR: must declare symbol"}
	
	// build command
	icmd := sqlcmdr.InsertCmd{ Tablename: "investment" }
	symbol := args[0]
	icmd.Add( "symbol", symbol )
	if( len(args) > 1){ icmd.Add( "description", args[1] )}
	if( len(args) > 2){ icmd.Add( "profit", args[2] )}
	
	// run command
	sthelper.SimpleInsert( icmd )
	
	return "created investment"
}

func show( args []string ) string{
	scmd := sqlcmdr.SelectCmd{ Tablename: "investment",RowID: true }
	result := sthelper.ExecShow( args, scmd, "investment.csv" );
	return result
}

func Createtable(){ 
	sqlcmdr.JustRunIt( 	"CREATE TABLE IF NOT EXISTS investment (" +
						"symbol VARCHAR( 50 )," +
						"description VARCHAR( 1000 )," +
						"profit DECIMAL(19, 4) )"  )
}

