//From Espen Haug, The Complete Guide To Option Pricing Formulas (2007)
package main

import (
  "fmt"
  "math"
)

//Cumulative normal distribution function
func CND(X float64) float64{
  a1 := 0.31938153
  a2 := -0.356563782
  a3 := 1.781477937
  a4 := -1.821255978
  a5 := 1.330274429

  L := math.Abs(X)

  K := 1 / (1 + 0.2316419 * L)

  w := 1 - 1 / math.Sqrt(2 * math.Pi) * math.Exp(-L * L / 2) * ((a1 * K) + a2 *
   math.Pow(K, 2) + a3 * math.Pow(K, 3) + a4 * math.Pow(K, 4) + a5 * math.Pow(K, 5))

  if X < 0 {
    w = 1.0 - w
  }
  return w
}

//The Black and Scholes (1973) Stock option formula
func BlackScholes(CallPutFlag string, S, X, T, r, v float64) float64 {

  d1 := (math.Log(S / X) + (r + math.Pow(v, 2) / 2) * T) / (v * math.Sqrt(T))
  d2 := d1 - (v * math.Sqrt(T))

  if CallPutFlag == "c"{
    return S * CND(d1) - X * (math.Exp(-r * T) * CND(d2))
  } else {
    return X * math.Exp(-r * T) * CND(-d2) - S * CND(-d1)
  }
}
