package ste_crawler

type testcase struct {
	p Parser
}

func (tc *testcase) testParserClass(t *Testing.T) {
	p := parser.New()
	if (len(p.configs) != 1) {
		t.Errorf("wrong length of configs. expected %d, go %d\n", 1, len(p.configs))
	}
	tc.p = p
}

func TestParser(t *testing.T) {
	
}

