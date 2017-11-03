package kantinepris

import (
  "testing"
)

func TestSagio1(t *testing.T) {
  m := Money { 900, 0 }
  actual := m.SagioPrice().String()
  expected := "11.50"
  if (actual != expected) {
    t.Errorf("Exepcted: %s, Got: %s.\n", expected, actual)
  }
}

func TestNayax1(t *testing.T) {
  m := Money { 900, 0 }
  actual := m.NayaxPrice().String()
  expected := "12.00"
  if (actual != expected) {
    t.Errorf("Exepcted: %s, Got: %s.\n", expected, actual)
  }
}

func TestMobilePay1(t *testing.T) {
  m := Money { 900, 0 }
  actual := m.MobilePayPrice().String()
  expected := "12.50"
  if (actual != expected) {
    t.Errorf("Exepcted: %s, Got: %s.\n", expected, actual)
  }
}
