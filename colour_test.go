package colour

import "testing"

func TestSprintf(t *testing.T) {
	expected := []string{
		"\033[30mblack \033[31mred \033[32mgreen \033[33myellow \033[34mblue \033[35mmagenta \033[36mcyan \033[37mwhite\033[0m",
		"\033[30m\033[40mblack \033[41mred \033[42mgreen \033[43myellow \033[44mblue \033[45mmagenta \033[46mcyan \033[47mwhite\033[0m",
		"\033[1m\033[30mblack \033[31mred \033[32mgreen \033[33myellow \033[34mblue \033[35mmagenta \033[36mcyan \033[37mwhite\033[0m",
		"\033[1m\033[30m\033[40mblack \033[41mred \033[42mgreen \033[43myellow \033[44mblue \033[45mmagenta \033[46mcyan \033[47mwhite\033[0m",
		"\033[4m\033[30mblack \033[31mred \033[32mgreen \033[33myellow \033[34mblue \033[35mmagenta \033[36mcyan \033[37mwhite\033[0m",
		"\033[4m\033[30m\033[40mblack \033[41mred \033[42mgreen \033[43myellow \033[44mblue \033[45mmagenta \033[46mcyan \033[47mwhite\033[0m",
	}
	actual := []string{
		Sprintf("^0black ^1red ^2green ^3yellow ^4blue ^5magenta ^6cyan ^7white^R"),
		Sprintf("^0^8black ^9red ^agreen ^byellow ^cblue ^dmagenta ^ecyan ^fwhite^R"),
		Sprintf("^B^0black ^1red ^2green ^3yellow ^4blue ^5magenta ^6cyan ^7white^R"),
		Sprintf("^B^0^8black ^9red ^agreen ^byellow ^cblue ^dmagenta ^ecyan ^fwhite^R"),
		Sprintf("^U^0black ^1red ^2green ^3yellow ^4blue ^5magenta ^6cyan ^7white^R"),
		Sprintf("^U^0^8black ^9red ^agreen ^byellow ^cblue ^dmagenta ^ecyan ^fwhite^R"),
	}
	for i := 0; i < len(actual); i++ {
		if expected[i] != actual[i] {
			t.Errorf("'%s' did not format as expected", actual[i])
		}
	}
}
