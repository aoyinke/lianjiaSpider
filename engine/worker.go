package engine

func  Worker(r Request) (ParseResults,error)  {
	return r.Parser.Parse(r.Url),nil
}