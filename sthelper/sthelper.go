package sthelper

import(
		"../sqlcmdr"
		"os"
)


func SliceShift( slice []string ) []string {
	if( len( slice ) == 0 ){
		return slice;
	}
	return slice[1:]
}

func ExecShow( args []string, scmd sqlcmdr.SelectCmd, csvout string ) string {
	
	// handle csv
	iscsv := false;
	if( len( args ) > 0 && args[0] == "csv" ){
		iscsv = true
		args = SliceShift( args );
	}
	
	// handle columns
	if( len( args ) > 0 )	{
		scmd.Columns = args[0]
	}	else	{
		scmd.Columns = "*"
	}
	
	conn := sqlcmdr.InitDB() 
	defer conn.Close();
	retvals := sqlcmdr.Select( conn, scmd )
	result := "";
	if( iscsv )	{
		result = sqlcmdr.ResultCSV( retvals, csvout );
	}else{
		result = sqlcmdr.ResultString( retvals );
	}
	
	return result
}

func CheckExists( filepath string ) bool {
	//based on Sridhar Ratnakumar's answer here: https://stackoverflow.com/questions/12518876/how-to-check-if-a-file-exists-in-go
	if _, err := os.Stat(filepath); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	}
	panic( "Unexpected result" )
	return true
}

func SimpleInsert( icmd sqlcmdr.InsertCmd ){
	conn := sqlcmdr.InitDB() 
	defer conn.Close()
	sqlcmdr.Insert( conn, icmd )
}

