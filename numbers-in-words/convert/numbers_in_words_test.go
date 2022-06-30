package convert

import "testing"

func TestSevenFourFive(t *testing.T) {
	want := "seven hundred and fourty five dollars"
	got := ConvertNumberInWords(745)

	if got != want {
		t.Errorf("TestSevenFourFive() = %s, want %s", got, want)
	}
}

func TestZero(t *testing.T) {
	want := "Zero dollars"
	got := ConvertNumberInWords(0)

	if got != want {
		t.Errorf("TestZero() = %s, want %s", got, want)
	}
}

func Test1(t *testing.T) {
	want := "One dollar"
	got := ConvertNumberInWords(1)

	if got != want {
		t.Errorf("Test1() = %s, want %s", got, want)
	}
}

func Test2(t *testing.T) {
	want := "Two dollars"
	got := ConvertNumberInWords(2)

	if got != want {
		t.Errorf("Test2() = %s, want %s", got, want)
	}
}

func Test3(t *testing.T) {
	want := "Three dollars"
	got := ConvertNumberInWords(3)

	if got != want {
		t.Errorf("Test3() = %s, want %s", got, want)
	}
}

func Test4(t *testing.T) {
	want := "Four dollars"
	got := ConvertNumberInWords(4)

	if got != want {
		t.Errorf("Test4() = %s, want %s", got, want)
	}
}

func TestMax(t *testing.T) {
	want := "nine hundred and ninety nine thousand nine hundred and ninety nine dollars"
	got := ConvertNumberInWords(999999)

	if got != want {
		t.Errorf("TestMax() = %s, want %s", got, want)
	}
}
