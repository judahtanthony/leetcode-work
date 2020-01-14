package main

import "fmt"

type StockSpanner struct {
    prices []int
    days   []int
    currDay int
}


func Constructor() StockSpanner {
    return StockSpanner{[]int{1000000}, []int{0}, 0}
}

func (this *StockSpanner) Next(price int) int {
    this.currDay++
    
    prevIndex := 0
    n := len(this.prices)

    l, r := 0, n
    for l < r {
        prevIndex = l + (r-l)/2
        if this.prices[prevIndex] > price {
            if prevIndex == n-1 || this.prices[prevIndex + 1] < price {
                break
            }
            if l == prevIndex {
                break
            }
            l = prevIndex
        } else {
            if r == prevIndex {
                break
            }
            r = prevIndex
        }
    }
    
    prevDay     := this.days[prevIndex]
    if prevIndex < n-1 {
        this.prices = this.prices[:prevIndex + 1]
        this.days   = this.days[:prevIndex + 1]
    }
    this.prices = append(this.prices, price)
    this.days   = append(this.days, this.currDay)
    
    return this.currDay - prevDay
}


/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */

func main() {
  prices := []int{}
  
	obj := Constructor();
  prices = append(prices, obj.Next(31))
  prices = append(prices, obj.Next(41))
  prices = append(prices, obj.Next(48))
  prices = append(prices, obj.Next(59))
  prices = append(prices, obj.Next(79))
  
  fmt.Println(prices)
}
