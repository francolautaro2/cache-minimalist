package main


import (
	"fmt"

)

func main(){
	// create a new app
	c := newCache(2)
	// set, get, delete and list of cache app
	c.Set(10, "item1")
	c.Set(11, "item2")
	c.Set(12, "item3")
	c.Set(13, "item4")
	c.Set(20, "item5")
	c.Get(10)
	c.Delete(10)
	fmt.Println(c.List())

	
}

// cache app 
type Cache struct {
	prev *Cache // WIP
	next *Cache // WIP
	k int
	ma map[int]caching
}

// caching format of data
type caching struct {
	value string
}

// create a new Cache app
func newCache(key int) *Cache {
	return &Cache{
		prev : nil,
		next : nil,
		k : key,
		ma : make(map[int]caching),
	}
}

// Set data to Cache app
func (c *Cache) Set(key int, value string) {
	c.ma[key] = caching{value : value}	 
}

// Get data of Cache app
func (c *Cache) Get(key int){
	for v := range(c.ma){
		if v == key {
			fmt.Println(c.ma[key])
		}
	}
}

// Delete data of Cache app
func (c *Cache) Delete(key int){
	delete(c.ma, key)
}

// get a list of caching data in Cache app
func (c *Cache) List() []string {
	entry := c.ma
	var values []string
	for _, v := range entry {
		values = append(values, v.value)
	}
	return values
}





