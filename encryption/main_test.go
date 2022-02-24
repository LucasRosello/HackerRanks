package main

import "testing"

//BenchMark
func BenchmarkEncryption(b *testing.B) {
	for n := 0; n < b.N; n++ {
		encryption("have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day")
	}
}

// encryption()
func TestMainEncryption(t *testing.T) {

	// Test 1
	s := ""
	expected := ""
	got := encryption(s)

	if got != expected {
		t.Errorf("Esperado: %v, obtenido: %v", expected, got)
	}

	// Test 2
	s = "feed the dog"
	expected = "fto ehg ee dd"
	got = encryption(s)

	if got != expected {
		t.Errorf("Esperado: %v, obtenido: %v", expected, got)
	}

	// Test 3
	s = "have a nice day"
	expected = "hae and via ecy"
	got = encryption(s)

	if got != expected {
		t.Errorf("Esperado: %v, obtenido: %v", expected, got)
	}

	// Test 4
	s = "have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day have a nice day"
	expected = "hihihihihihihihihihihihihihih acacacacacacacacacacacacacaca vevevevevevevevevevevevevevev edededededededededededededede aaaaaaaaaaaaaaaaaaaaaaaaaaaaa nynynynynynynynynynynynynynyn ihihihihihihihihihihihihihihi cacacacacacacacacacacacacacac eveveveveveveveveveveveveveve dedededededededededededededed aaaaaaaaaaaaaaaaaaaaaaaaaaaaa ynynynynynynynynynynynynynyny hihihihihihihihihihihihihihi acacacacacacacacacacacacacac veveveveveveveveveveveveveve edededededededededededededed aaaaaaaaaaaaaaaaaaaaaaaaaaaa nynynynynynynynynynynynynyny ihihihihihihihihihihihihihih cacacacacacacacacacacacacaca evevevevevevevevevevevevevev dededededededededededededede aaaaaaaaaaaaaaaaaaaaaaaaaaaa ynynynynynynynynynynynynynyn hihihihihihihihihihihihihihi acacacacacacacacacacacacacac veveveveveveveveveveveveveve edededededededededededededed aaaaaaaaaaaaaaaaaaaaaaaaaaaa nynynynynynynynynynynynynyny"
	got = encryption(s)

	if got != expected {
		t.Errorf("Esperado: %v, obtenido: %v", expected, got)
	}
}

// encrypt()

func TestEncrypt(t *testing.T) {
	s := "haveaniceday"
	upRoot := int32(4)
	downRoot := int32(3)
	expected := "hae and via ecy"

	got := encrypt(s, upRoot, downRoot)

	if got != expected {
		t.Errorf("Esperado: %v, obtenido: %v", expected, got)
	}
}

// generateDistribution()
func TestGenerateDistribution(t *testing.T) {
	// Test Square Root Of 54
	upExpected := int32(8)
	downExpected := int32(7)
	upRoot, downRoot := generateDistribution(float64(54))

	if upExpected != upRoot {
		t.Errorf("Esperado: %v, obtenido: %v", upExpected, upRoot)
	}

	if downExpected != downRoot {
		t.Errorf("Esperado: %v, obtenido: %v", downExpected, downRoot)
	}

	// Test Square Root Of 9
	upExpected = int32(3)
	downExpected = int32(3)
	upRoot, downRoot = generateDistribution(float64(9))

	if upExpected != upRoot {
		t.Errorf("Esperado: %v, obtenido: %v", upExpected, upRoot)
	}

	if downExpected != downRoot {
		t.Errorf("Esperado: %v, obtenido: %v", downExpected, downRoot)
	}
}

// clean()
func TestClean(t *testing.T) {
	s := "     Hi    fr om    the   test   "
	expected := "Hifromthetest"
	got := clean(s)

	if got != expected {
		t.Errorf("Esperado: %v, obtenido: %v", expected, got)
	}
}
