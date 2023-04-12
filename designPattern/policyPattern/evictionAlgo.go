package policypattern
type EvictionAlgo interface {
    evict(c *Cache)
}