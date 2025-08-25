package lazy_test

import (
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	. "github.com/onsi/gomega"

	"go.robertomontagna.dev/zapfluent/internal/functional/lazy"
)

func TestLazy_New_EvaluatesLazily(t *testing.T) {
	g := NewGomegaWithT(t)

	var counter int64
	lazyValue := lazy.New(func() int {
		atomic.AddInt64(&counter, 1)
		return 10
	})
	g.Expect(atomic.LoadInt64(&counter)).To(BeZero(), "The function should not be called upon creation")

	result := lazyValue.Get()

	g.Expect(result).To(Equal(10))
	g.Expect(atomic.LoadInt64(&counter)).To(Equal(int64(1)), "The function should be called once after Get")
}

func TestLazy_Get_IsMemoized(t *testing.T) {
	g := NewGomegaWithT(t)

	var counter int64
	lazyValue := lazy.New(func() int {
		atomic.AddInt64(&counter, 1)
		return 10
	})
	lazyValue.Get()
	g.Expect(atomic.LoadInt64(&counter)).To(Equal(int64(1)))

	result := lazyValue.Get()

	g.Expect(result).To(Equal(10))
	g.Expect(atomic.LoadInt64(&counter)).To(Equal(int64(1)), "The function should not be called again")
}

func TestLazy_Of_CreatesFromValue(t *testing.T) {
	g := NewGomegaWithT(t)

	lazyValue := lazy.Of(20)

	result := lazyValue.Get()

	g.Expect(result).To(Equal(20))
}

func TestLazy_Map_AppliesFunction(t *testing.T) {
	g := NewGomegaWithT(t)

	var counter int64
	lazyValue := lazy.New(func() int {
		atomic.AddInt64(&counter, 1)
		return 5
	})
	mappedLazy := lazy.Map(&lazyValue, func(i int) string {
		return "value-" + strconv.Itoa(i)
	})
	g.Expect(atomic.LoadInt64(&counter)).To(BeZero(), "The original function should not be called yet")

	result := mappedLazy.Get()

	g.Expect(result).To(Equal("value-5"))
	g.Expect(atomic.LoadInt64(&counter)).To(Equal(int64(1)), "The original function should be called once")
}

func TestLazy_FlatMap_AppliesFunction(t *testing.T) {
	g := NewGomegaWithT(t)

	var counter1 int64
	lazyValue1 := lazy.New(func() int {
		atomic.AddInt64(&counter1, 1)
		return 3
	})
	var counter2 int64
	flatMappedLazy := lazy.FlatMap(&lazyValue1, func(i int) lazy.Lazy[string] {
		return lazy.New(func() string {
			atomic.AddInt64(&counter2, 1)
			return "nested-" + strconv.Itoa(i)
		})
	})
	g.Expect(atomic.LoadInt64(&counter1)).To(BeZero())
	g.Expect(atomic.LoadInt64(&counter2)).To(BeZero())

	result := flatMappedLazy.Get()

	g.Expect(result).To(Equal("nested-3"))
	g.Expect(atomic.LoadInt64(&counter1)).To(Equal(int64(1)))
	g.Expect(atomic.LoadInt64(&counter2)).To(Equal(int64(1)))
}

func TestLazy_Get_IsConcurrentSafe(t *testing.T) {
	g := NewGomegaWithT(t)

	var counter int64
	lazyValue := lazy.New(func() int {
		time.Sleep(10 * time.Millisecond)
		atomic.AddInt64(&counter, 1)
		return 100
	})
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			g.Expect(lazyValue.Get()).To(Equal(100))
		}()
	}
	wg.Wait()

	g.Expect(atomic.LoadInt64(&counter)).To(Equal(int64(1)), "The function should be called exactly once")
}
