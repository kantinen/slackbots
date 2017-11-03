package main

import (
  "fmt"
  "strconv"
)

type Money struct {
  n int
  err int
}

func (m Money) String() string {
  n := m.n
  s := strconv.Itoa(n)

  prefix := ""
  if (n < 0) {
    prefix = "-"
    n = (-1) * m.n
  }

  if (n < 10) {
    prefix += "0.0"
  } else if (n < 100) {
    prefix += "0."
  } else {
    prefix += s[0:len(s)-2] + "."
    s = s[len(s)-2:len(s)]
  }

  return prefix + s
}

func (m Money) AddMoney(n Money) Money {
  return Money { m.n + n.n, m.err + n.err }
}

func (m Money) AddInt(n int) Money {
  return Money { m.n + n, m.err }
}

func (m Money) Div(v int) Money {
  n := m.n
  nv := n / v
  err := 0

  // Round to nearest even
  rn10 := n % 10
  rm2 := nv % 2
  if (rn10 > 5 || (rn10 == 5 && rm2 == 1)) {
    nv += 1
    err += 10 - rn10
  } else {
    err += rn10
  }

  return Money { nv, err }
}

func (m Money) Normalize() Money {
  n := m.n
  rn100 := n % 100
  if (rn100 < 25) {
    n = n - rn100
  } else if (rn100 > 75) {
    n = n + (100 -rn100)
  } else {
    n = n - rn100 + 50
  }

  return Money { n, m.err }
}

func (m Money) SagioPrice() Money {
  return m.AddMoney(m.Div(4)).Normalize()
}

func (m Money) NayaxPrice() Money {
  return m.SagioPrice().AddMoney(m.Div(17)).Normalize()
}

func (m Money) MobilePayPrice() Money {
  return m.SagioPrice().AddInt(75).Normalize()
}


func main() {
  m := Money { 1000, 0 }
  fmt.Printf("Sagio: %s\n", m.SagioPrice())
  fmt.Printf("Nayax: %s\n", m.NayaxPrice())
  fmt.Printf("MobilePay: %s\n", m.MobilePayPrice())
}
