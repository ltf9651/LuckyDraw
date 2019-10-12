package main

import (
	"fmt"
	"github.com/kataras/iris/httptest"
	"sync"
	"testing"
)

func TestMVC(t *testing.T) {
	e := httptest.New(t, newApp())

	var wg sync.WaitGroup
	e.GET("/").Except().Status(httptest.StatusOK).Body().Equal("当前参与抽奖人数：0\n")

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			e.POST("/import").WithFormField("users", fmt.Sprintf("test_u%d", i)).Except().Status(httptest.StatusOK)
		}(i)
	}

	wg.Wait()

	e.GET("/").Except().Status(httptest.StatusOK).Body().Equal("当前参与抽奖人数：100\n")

	e.GET("/lucky").Except().Status(httptest.StatusOK).Body().Equal("当前参与抽奖人数：99\n")
}
