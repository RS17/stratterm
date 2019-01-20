package strategy

import (
	"../sqlcmdr"
	"../sthelper"
)

type strategy struct{
	name string 
	description string
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
	if( len(args) < 2) {return "ERROR: must declare name and description"}
	name := args[0]
	description := args[1]
	icmd := sqlcmdr.InsertCmd{ Tablename: "strategy" }
	icmd.Add( "name", name )
	icmd.Add( "description", description )
	
	sthelper.SimpleInsert( icmd )
	
	return "created strategy"
}

func show( args []string ) string{
	scmd := sqlcmdr.SelectCmd{ Tablename: "strategy" }
	addjoins( &scmd, args )
	result := sthelper.ExecShow( args, scmd, "strategy.csv" )
	return result
}

func Createtable(){ 
	sqlcmdr.JustRunIt( 	"CREATE TABLE IF NOT EXISTS strategy (" + 
						" name VARCHAR( 50 ) NOT NULL PRIMARY KEY," +
						"description VARCHAR( 1000 ) )"  )
}

func addjoins( scmd *sqlcmdr.SelectCmd, columns []string ){
	for _, col := range columns {
        switch ( col ){
			case "profit", "symbol" : 
				scmd.AddJoin( "strategy-name", "investmentstrategy-strategy", "LEFT" )
				scmd.AddJoin( "investmentstrategy-investment", "investment-rowid", "LEFT" )
			default :  
		}
	}
} 

