package services

import (
	"fmt"
	"test-flow-code-gen/testflow"
	"testing"
)

type MathServiceTest interface {
	testflow.TestProcess
	TestCalcSquareArea(t *testing.T)
	TestCalcTriangleArea(t *testing.T)
}

type mathServiceTest struct {
	mathServiceTestTitle mathServiceTestTitle
	mathService          *MathService
	// add common parameters between test cases
}

type mathServiceTestTitle struct {
	titleTestCalcSquareArea   string
	titleTestCalcTriangleArea string
}

func NewMathServiceTestObject(
	testTitles mathServiceTestTitle,
) MathServiceTest {
	return &mathServiceTest{
		mathService: NewMathService(),
		mathServiceTestTitle: mathServiceTestTitle{
			titleTestCalcSquareArea:   testTitles.titleTestCalcSquareArea,
			titleTestCalcTriangleArea: testTitles.titleTestCalcTriangleArea,
		},
		// add initialize common parameters between test cases
	}
}

func NewMathServiceTestTitle() mathServiceTestTitle {
	return mathServiceTestTitle{
		titleTestCalcSquareArea:   "TestCalcSquareArea",
		titleTestCalcTriangleArea: "TestCalcTriangleArea",
	}
}

func TestMathService(t *testing.T) {
	testTitles := NewMathServiceTestTitle()
	testInterface := NewMathServiceTestObject(testTitles)
	testDescribes := []testflow.Describe{
		{Title: testTitles.titleTestCalcSquareArea, Test: testInterface.TestCalcSquareArea},
		{Title: testTitles.titleTestCalcTriangleArea, Test: testInterface.TestCalcTriangleArea},
	}
	testflow.TestSuite(t, testInterface, testDescribes)
}

func (ut *mathServiceTest) BeforeAll(t *testing.T) {}

func (ut *mathServiceTest) BeforeTest(t *testing.T, testTitle string) {
	// setup process by test method case
	switch testTitle {
	case ut.mathServiceTestTitle.titleTestCalcSquareArea:
	case ut.mathServiceTestTitle.titleTestCalcTriangleArea:
	}
}

func (ut *mathServiceTest) BeforeCase(t *testing.T, testTitle string) {
	// setup process by test case
	switch testTitle {
	case ut.mathServiceTestTitle.titleTestCalcSquareArea:
	case ut.mathServiceTestTitle.titleTestCalcTriangleArea:
	}
}

func (ut *mathServiceTest) AfterAll(t *testing.T) {}

func (ut *mathServiceTest) AfterTest(t *testing.T, testTitle string) {
	// teardown process by test method case
	switch testTitle {
	case ut.mathServiceTestTitle.titleTestCalcSquareArea:
	case ut.mathServiceTestTitle.titleTestCalcTriangleArea:
	}
}

func (ut *mathServiceTest) AfterCase(t *testing.T, testTitle string) {
	// teardown process by test case
	switch testTitle {
	case ut.mathServiceTestTitle.titleTestCalcSquareArea:
	case ut.mathServiceTestTitle.titleTestCalcTriangleArea:
	}
}

func (ut *mathServiceTest) TestCalcSquareArea(t *testing.T) {

	type Argument struct {
		width  int
		height int
	}

	type TestCase struct {
		testflow.It
		input       Argument
		expect      int
		expectError error
	}
	testCases := []TestCase{
		// TODO design test case
		// {
		// 	It: testflow.It{
		// 		SerialNum:   1,
		// 		Title:       "",
		// 		TestClass:   testflow.TestClassNormal, // testflow.TestClassSemiNormal | testflow.TestClassError
		// 		ExpectError: nil,
		// 	},
		// 	input: Argument{},
		// 	expect: nil,
		// },
		// add test case
	}
	for _, testCase := range testCases {
		t.Log(testCase.SerialNum, testCase.TestClass.String(), testCase.Title, testCase.input)
		ut.BeforeCase(t, ut.mathServiceTestTitle.titleTestCalcSquareArea)

		err := ut.mathService.CalcSquareArea(testCase.input.width, testCase.input.height)
		fmt.Println(err)
		//assert.Equal(t, testCase.expectError, err)

		ut.AfterCase(t, ut.mathServiceTestTitle.titleTestCalcSquareArea)
	}
}

func (ut *mathServiceTest) TestCalcTriangleArea(t *testing.T) {

	type Argument struct {
		width  int
		height int
	}

	type TestCase struct {
		testflow.It
		input       Argument
		expect      int
		expectError error
	}
	testCases := []TestCase{
		// TODO design test case
		// {
		// 	It: testflow.It{
		// 		SerialNum:   1,
		// 		Title:       "",
		// 		TestClass:   testflow.TestClassNormal, // testflow.TestClassSemiNormal | testflow.TestClassError
		// 		ExpectError: nil,
		// 	},
		// 	input: Argument{},
		// 	expect: nil,
		// },
		// add test case
	}
	for _, testCase := range testCases {
		t.Log(testCase.SerialNum, testCase.TestClass.String(), testCase.Title, testCase.input)
		ut.BeforeCase(t, ut.mathServiceTestTitle.titleTestCalcTriangleArea)

		err := ut.mathService.CalcTriangleArea(testCase.input.width, testCase.input.height)
		fmt.Println(err)
		//assert.Equal(t, testCase.expectError, err)

		ut.AfterCase(t, ut.mathServiceTestTitle.titleTestCalcTriangleArea)
	}
}
