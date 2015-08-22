package fma;

func buildParameters(params map[string] string) string{
	 p := ""
	 for key,value := range(params){
		 p += key+"="+value+"&"
	 }
	 return p;
}
