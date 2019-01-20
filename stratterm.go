package main

import (
	"./strategy"
	"./investment"
	"./investmentstrategy"
	"os"
	"fmt"
	"./sthelper"
)

func main() {
	args := os.Args
	if( len( args ) <= 1 ){ 
		fmt.Println( "Error: need args [imagine help screen stuff here]")
	}	else if( args[1] == "setup" ) {	
		setup()
	}	else if( !sthelper.CheckExists( "data.db" ) ){
		fmt.Println("No database found!  Please type 'stratterm setup' to build database" );
	}	else if( args[1] == "strategy" )	{
		fmt.Println( strategy.Handle( args[2:] ) )
	}	else if( args[1] == "investment" )	{
		fmt.Println( investment.Handle( args[2:] ) )
	}	else if( args[1] == "investmentstrategy" )	{
		fmt.Println( investmentstrategy.Handle( args[2:] ) )
	}	else if( args[1] == "csv" ) {
		csv()
	}	else {
		panic("ERROR: invalid argument " + args[1] )
	}
}

func csv() {
	baseargs := []string{ "show", "csv" }
	strategy.Handle( baseargs );
	investment.Handle( baseargs );
	investmentstrategy.Handle( baseargs );
}

func setup(){
	// this only creates tables, if you want to re-create tables you'll need to delete the dara.db file manually to avoid accidental deletes
	strategy.Createtable()
	investment.Createtable()
	investmentstrategy.Createtable()
	fmt.Println("data tables created" );
}
	
