package main

import (
	"fmt"
	"math/rand"
	"time"
)

func estimate_pi(n int) (float64, int64) {
  var ans float64
  var num_point_total float64
  var num_point_circle float64
  var time_start time.Time
  var time_end time.Time
  var time_diff int64
  var start int64
  var end int64
  time_start = time.Now()
  for i:=0; i<n; i++ {
    x := rand.Float64()
    y := rand.Float64()
    distance := x*x + y*y
    if distance <= 1 {
      num_point_circle = num_point_circle + 1
    }
    num_point_total = num_point_total + 1
  }
  ans = 4 * num_point_circle / num_point_total
  time_end = time.Now()
  start = time_start.Unix()
  end = time_end.Unix()
  time_diff = end - start
  return ans, time_diff
}

func main() {
  var i int
  fmt.Scan(&i)
  ans, diff := estimate_pi(i)
  fmt.Println(ans)
  fmt.Printf("time taken %ds", diff)
}
