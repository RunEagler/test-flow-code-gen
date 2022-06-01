package testflow

import "testing"

// TestClass test classify
type TestClass int

const (
	TestClassNormal     TestClass = iota // normal
	TestClassSemiNormal                  // semi-normal
	TestClassError                       // error
)

func (c TestClass) String() string {
	switch c {
	case TestClassNormal:
		return "normal"
	case TestClassSemiNormal:
		return "semi-normal"
	case TestClassError:
		return "error"
	}
	return ""
}

// Describe メソッドごとのテストケース(1メソッド1Describe)
type Describe struct {
	Title string
	Test  func(t *testing.T) // 各テストメソッド
}

// It test case by 1 test method
type It struct {
	SerialNum int
	Title     string
	TestClass TestClass
}

// TestProcess test process methods
type TestProcess interface {
	BeforeAll(t *testing.T)                    // setup process by 1 test file
	AfterAll(t *testing.T)                     // teardown process by 1 test file
	BeforeTest(t *testing.T, testTitle string) // setup process by 1 test method
	AfterTest(t *testing.T, testTitle string)  // teardown process by 1 test method
	BeforeCase(t *testing.T, testTitle string) // setup process by 1 test case
	AfterCase(t *testing.T, testTitle string)  // teardown process by 1 test case
}

// TestSuite is test life cycle.
func TestSuite(t *testing.T, p TestProcess, testDescribes []Describe) {
	p.BeforeAll(t)
	for _, testDescribe := range testDescribes {
		p.BeforeTest(t, testDescribe.Title)
		t.Run(testDescribe.Title, testDescribe.Test)
		p.AfterTest(t, testDescribe.Title)
	}
	p.AfterAll(t)
}
